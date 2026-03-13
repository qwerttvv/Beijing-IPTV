// iptvscanner by qwerttvv is marked with CC0 1.0 Universal.
// https://creativecommons.org/publicdomain/zero/1.0/
// https://github.com/qwerttvv/Beijing-IPTV

package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/ipv4"
)

var (
	globalHandle   *pcap.Handle
	globalIPv4Conn *ipv4.PacketConn
)

func ipToUint32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func uint32ToIP(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

func readUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("读取输入失败: %v", err)
	}
	return strings.TrimSpace(text)
}

func chooseInterface() *net.Interface {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("获取网络接口失败: %v", err)
	}
	var validIfaces []net.Interface
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagMulticast == 0 {
			continue
		}
		validIfaces = append(validIfaces, iface)
	}
	if len(validIfaces) == 0 {
		log.Fatalf("未找到可用的多播网络接口")
	}
	fmt.Println("\n选择网络接口:\n")
	for i, iface := range validIfaces {
		fmt.Printf("%d. 名称: %s\n", i+1, iface.Name)
		if addrs, err := iface.Addrs(); err == nil {
			for _, addr := range addrs {
				fmt.Printf("   %s\n", addr.String())
			}
		}
	}
	choice, err := strconv.Atoi(readUserInput("\n输入接口编号: "))
	if err != nil || choice < 1 || choice > len(validIfaces) {
		log.Fatalf("接口编号无效或超出范围")
	}
	return &validIfaces[choice-1]
}

func extractIPv4s(iface *net.Interface) []net.IP {
	var ips []net.IP
	addrs, err := iface.Addrs()
	if err != nil {
		return ips
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok {
			if ip := ipNet.IP.To4(); ip != nil {
				ips = append(ips, ip)
			}
		}
	}
	return ips
}

func choosePcapDevice(iface *net.Interface, ifaceIPs []net.IP) string {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("查找pcap设备失败: %v", err)
	}
	if len(devices) == 0 {
		log.Fatalf("未找到pcap设备！安装 Npcap，并确保安装时勾选 WinPcap 兼容")
	}
	var chosenDev string
	for _, dev := range devices {
		if dev.Name == iface.Name || strings.Contains(dev.Description, iface.Name) {
			chosenDev = dev.Name
			break
		}
		for _, pcapAddr := range dev.Addresses {
			if pcapIP := pcapAddr.IP.To4(); pcapIP != nil {
				for _, ip := range ifaceIPs {
					if pcapIP.Equal(ip) {
						chosenDev = dev.Name
						break
					}
				}
			}
			if chosenDev != "" {
				break
			}
		}
	}
	if chosenDev == "" {
		fmt.Println("未找到与接口匹配的pcap设备，手动选择:")
		for i, dev := range devices {
			fmt.Printf("%d. 名称: %s, 描述: %s\n", i+1, dev.Name, dev.Description)
		}
		choice, err := strconv.Atoi(readUserInput("\n输入设备编号: "))
		if err != nil || choice < 1 || choice > len(devices) {
			log.Fatalf("设备编号无效或超出范围")
		}
		chosenDev = devices[choice-1].Name
	}
	return chosenDev
}

func scanIP(ip net.IP, iface *net.Interface, scanDuration time.Duration) (map[uint16]bool, error) {
	openPorts := make(map[uint16]bool)
	multicastAddr := &net.UDPAddr{IP: ip}
	if err := globalIPv4Conn.JoinGroup(iface, multicastAddr); err != nil {
		return openPorts, fmt.Errorf("加入多播组 %v 失败: %v", ip, err)
	}
	defer func() {
		if err := globalIPv4Conn.LeaveGroup(iface, multicastAddr); err != nil {
			log.Printf("警告: 离开多播组失败: %v", err)
		}
	}()
	filter := fmt.Sprintf("udp and dst host %s", ip.String())
	if err := globalHandle.SetBPFFilter(filter); err != nil {
		return openPorts, fmt.Errorf("设置BPF过滤器失败: %v", err)
	}
	packetSource := gopacket.NewPacketSource(globalHandle, globalHandle.LinkType())
	packetSource.Lazy = true
	packetSource.NoCopy = true
	timeout := time.After(scanDuration)
	for {
		select {
		case <-timeout:
			return openPorts, nil
		case packet, ok := <-packetSource.Packets():
			if !ok {
				return openPorts, nil
			}
			if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
				if udp, ok := udpLayer.(*layers.UDP); ok {
					openPorts[uint16(udp.DstPort)] = true
				}
			}
		}
	}
}

func main() {
	flag.Usage = func() {
		fmt.Printf("用法: %s <起始IP> <结束IP> [扫描时长(毫秒, 默认2222, 小于666则恢复为默认值)]\n", os.Args[0])
		fmt.Printf("示例: %s 239.3.1.1 239.3.1.254\n", os.Args[0])
	}
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	startIP := net.ParseIP(args[0])
	endIP := net.ParseIP(args[1])
	if startIP == nil || endIP == nil || !startIP.IsMulticast() || !endIP.IsMulticast() {
		log.Fatal("起始IP和结束IP必须是有效的IPv4多播地址（224.0.0.0到239.255.255.255）")
	}
	startUint := ipToUint32(startIP)
	endUint := ipToUint32(endIP)
	if startUint > endUint {
		log.Fatalf("起始IP必须小于结束IP")
	}
	scanDurationMs := 2222
	if len(args) >= 3 {
		if dur, err := strconv.Atoi(args[2]); err == nil && dur >= 666 {
			scanDurationMs = dur
		}
	}
	scanDuration := time.Duration(scanDurationMs) * time.Millisecond
	fmt.Println("\nhttps://github.com/qwerttvv/Beijing-IPTV")
	iface := chooseInterface()
	ifaceIPs := extractIPv4s(iface)
	chosenDev := choosePcapDevice(iface, ifaceIPs)
	udpConn, err := net.ListenPacket("udp4", "0.0.0.0:0")
	if err != nil {
		log.Fatalf("打开UDP连接失败: %v", err)
	}
	defer udpConn.Close()
	globalIPv4Conn = ipv4.NewPacketConn(udpConn)
	defer globalIPv4Conn.Close()
	globalHandle, err = pcap.OpenLive(chosenDev, 65535, true, scanDuration)
	if err != nil {
		log.Fatalf("pcap打开设备失败: %v", err)
	}
	defer globalHandle.Close()

	fmt.Printf("\n#EXTM3U name=\"Beijing-IPTV\"\n")
	var results []string
	for ipUint := startUint; ipUint <= endUint; ipUint++ {
		curIP := uint32ToIP(ipUint)
		ports, err := scanIP(curIP, iface, scanDuration)
		if err != nil {
			log.Printf("扫描 %s 失败: %v", curIP, err)
			continue
		}
		for port := range ports {
			entry := fmt.Sprintf("#EXTINF:-1,%s:%d\nrtp://%s:%d\n", curIP.String(), port, curIP.String(), port)
			fmt.Print(entry)
			results = append(results, entry)
		}
	}
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("无法获取可执行文件路径: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	now := time.Now()
	filename := filepath.Join(exeDir, now.Format("IPTV-2006-01-02_15-04-05")+".m3u")
	fmt.Printf("\nIPTV列表已经保存到 %s\n", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("创建保存文件失败: %v", err)
	}
	defer file.Close()
	if _, err := file.WriteString("#EXTM3U name=\"Beijing-IPTV\"\n"); err != nil {
		log.Fatalf("写入文件失败: %v", err)
	}
	for _, line := range results {
		if _, err := file.WriteString(line); err != nil {
			log.Fatalf("写入文件失败: %v", err)
		}
	}
}
