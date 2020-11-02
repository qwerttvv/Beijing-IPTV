# Beijing-IPTV

------------

## 北京联通的IPTV节目列表

### 自用单播列表永久地址

**你可以更改udpxy地址为192.168.123.1:23234，进而直接使用该列表——强烈推荐**

源

- [https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom.m3u](https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom.m3u "https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom.m3u")

镜像

- [https://qwerttvv.github.io/Beijing-IPTV/IPTV-Unicom.m3u](https://qwerttvv.github.io/Beijing-IPTV/IPTV-Unicom.m3u "https://qwerttvv.github.io/Beijing-IPTV/IPTV-Unicom.m3u")
- [https://beijing-iptv.vercel.app/IPTV-Unicom.m3u](https://beijing-iptv.vercel.app/IPTV-Unicom.m3u "https://beijing-iptv.vercel.app/IPTV-Unicom.m3u")
- [https://raw.fastgit.org/qwerttvv/Beijing-IPTV/master/IPTV-Unicom.m3u](https://raw.fastgit.org/qwerttvv/Beijing-IPTV/master/IPTV-Unicom.m3u "https://raw.fastgit.org/qwerttvv/Beijing-IPTV/master/IPTV-Unicom.m3u")


### 通用组播列表永久地址

源

- [https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Multicast.m3u](https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Multicast.m3u "https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Multicast.m3u")

镜像

- [https://qwerttvv.github.io/Beijing-IPTV/IPTV-Unicom-Multicast.m3u](https://qwerttvv.github.io/Beijing-IPTV/IPTV-Unicom-Multicast.m3u "https://qwerttvv.github.io/Beijing-IPTV/IPTV-Unicom-Multicast.m3u")
- [https://beijing-iptv.vercel.app/IPTV-Unicom-Multicast.m3u](https://beijing-iptv.vercel.app/IPTV-Unicom-Multicast.m3u "https://beijing-iptv.vercel.app/IPTV-Unicom-Multicast.m3u")
- [https://raw.fastgit.org/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Multicast.m3u](https://raw.fastgit.org/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Multicast.m3u "https://raw.fastgit.org/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Multicast.m3u")

### 节目列表来源
1. 电视盒子抓包官方列表
2. 组播地址[扫描](https://github.com/sdhzdmzzl/iptv_channel_scanner_windows "扫描")
3. 交叉[参考1](https://gist.github.com/sdhzdmzzl/93cf74947770066743fff7c7f4fc5820 "参考1")
4. 交叉[参考2](https://github.com/islercn/BeiJing-Unicom-IPTV-List "参考2")
5. 交叉[参考3](https://github.com/friskit-china/beijing-unicom-iptv-tweaker "参考3")

### 其他说明
1. 增加非北京联通IPTV来源的高清频道，比如CCTV13HD，这些频道会单独标记【非官方】，其中CCTV4K为超清HDR频道，支持HDR的播放设备可以有不错的体验（组播文件没有添加非官方来源频道），个别非官方频道可能会用到IPv6网络
2. 那些原本单独收费的频道，全部免费解锁直接看
3. 有一些类似上网课一类的台，还有无意义的测试频道，还有一大堆购物频道，一并都删了
4. 说实话，列表里地方的高清台再往后的那些台，平时基本不看…不能说基本吧，应该是根本不看……
5. EPG节目单有时候儿不准，看个大概吧
6. 有的频道有高清和非高清，节目信号一样的。只保留了高清，避免重复，非高清的全部删了
7. 原始扫描到的列表文件[IPTV-Unicom-Scan-A.m3u](https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Scan-A.m3u "IPTV-Unicom-Scan-A.m3u")是239.3.1.X，[IPTV-Unicom-Scan-B.m3u](https://raw.githubusercontent.com/qwerttvv/Beijing-IPTV/master/IPTV-Unicom-Scan-B.m3u "IPTV-Unicom-Scan-B.m3u")是239.2.1.X，什么台都有，未经整理，仅作参考
8. 恕无法提供抓包的原始文件，文件里边有私人的信息，包括时移地址什么的。如果脱敏，那和上述已提供的列表就没啥区别了
9. 本项目的issues我关闭了。如果你想反馈问题或者提出宝贵的意见建议，总之你真心想联系我的话，你总会想到办法的：）

------------

## EPG节目单 [主页](http://epg.51zmt.top:8000/ "主页") [GitHub](https://github.com/supzhang/epg "GitHub")

**M3U文件已经内置EPG节目单列表，支持x-tvg-url标签的播放软件无需手动加载EPG文件**

源

- [http://epg.51zmt.top:8000/e.xml](http://epg.51zmt.top:8000/e.xml "http://epg.51zmt.top:8000/e.xml")
- [http://epg.51zmt.top:8000/e.xml.gz](http://epg.51zmt.top:8000/e.xml.gz "http://epg.51zmt.top:8000/e.xml.gz")

镜像

- [https://raw.githubusercontent.com/supzhang/epg/master/e.xml](https://raw.githubusercontent.com/supzhang/epg/master/e.xml "https://raw.githubusercontent.com/supzhang/epg/master/e.xml")
- [https://raw.githubusercontent.com/supzhang/epg/master/e.xml.gz](https://raw.githubusercontent.com/supzhang/epg/master/e.xml.gz "https://raw.githubusercontent.com/supzhang/epg/master/e.xml.gz")
- [https://raw.fastgit.org/supzhang/epg/master/e.xml](https://raw.fastgit.org/supzhang/epg/master/e.xml "https://raw.fastgit.org/supzhang/epg/master/e.xml")
- [https://raw.fastgit.org/supzhang/epg/master/e.xml.gz](https://raw.fastgit.org/supzhang/epg/master/e.xml.gz "https://raw.fastgit.org/supzhang/epg/master/e.xml.gz")

------------

**感谢外部引用的作者，以及过程中所有帮助过我的人**
