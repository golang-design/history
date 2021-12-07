# Go 语言历史


> 英文原版：[https://golang.design/history/](https://golang.design/history/)，来自 Changkun Ou 及其他贡献者

> 此翻译文档可能更新不及时，您可以查看英文原版获得最新更新。


此文档收集了Go语言开发过程中公开的讨论及提案，以全面地展现Go语言发展历史。

<a id="top"></a>

**Table of Contents**

* [Go 语言历史](#Go+%E8%AF%AD%E8%A8%80%E5%8E%86%E5%8F%B2)
  * [免责声明](#%E5%85%8D%E8%B4%A3%E5%A3%B0%E6%98%8E)
  * [资料来源](#%E8%B5%84%E6%96%99%E6%9D%A5%E6%BA%90)
  * [起源](#%E8%B5%B7%E6%BA%90)
  * [贡献者](#%E8%B4%A1%E7%8C%AE%E8%80%85)
    * [核心作者](#%E6%A0%B8%E5%BF%83%E4%BD%9C%E8%80%85)
    * [编译运行时团队](#%E7%BC%96%E8%AF%91%E8%BF%90%E8%A1%8C%E6%97%B6%E5%9B%A2%E9%98%9F)
    * [小组会谈](#%E5%B0%8F%E7%BB%84%E4%BC%9A%E8%B0%88)
  * [时间线](#%E6%97%B6%E9%97%B4%E7%BA%BF)
  * [语言设计](#%E8%AF%AD%E8%A8%80%E8%AE%BE%E8%AE%A1)
    * [Misc](#Misc)
    * [切片 (1.2)](#%E5%88%87%E7%89%87+%281.2%29)
    * [包管理 (1.4, 1.5, 1.7)](#%E5%8C%85%E7%AE%A1%E7%90%86+%281.4%2C+1.5%2C+1.7%29)
    * [类型别名 (1.9)](#%E7%B1%BB%E5%9E%8B%E5%88%AB%E5%90%8D+%281.9%29)
    * [Defer (1.14)](#Defer+%281.14%29)
    * [错误值 (1.13)](#%E9%94%99%E8%AF%AF%E5%80%BC+%281.13%29)
    * [通道、选择器](#%E9%80%9A%E9%81%93%E3%80%81%E9%80%89%E6%8B%A9%E5%99%A8)
    * [泛型](#%E6%B3%9B%E5%9E%8B)
  * [编译工具链](#%E7%BC%96%E8%AF%91%E5%B7%A5%E5%85%B7%E9%93%BE)
    * [编译](#%E7%BC%96%E8%AF%91)
    * [连接器](#%E8%BF%9E%E6%8E%A5%E5%99%A8)
    * [调试器](#%E8%B0%83%E8%AF%95%E5%99%A8)
    * [竞争检测器](#%E7%AB%9E%E4%BA%89%E6%A3%80%E6%B5%8B%E5%99%A8)
    * [跟踪器](#%E8%B7%9F%E8%B8%AA%E5%99%A8)
    * [锁分析](#%E9%94%81%E5%88%86%E6%9E%90)
    * [构建](#%E6%9E%84%E5%BB%BA)
    * [模块](#%E6%A8%A1%E5%9D%97)
    * [gopls](#gopls)
    * [测试，x/perf](#%E6%B5%8B%E8%AF%95%EF%BC%8Cx%2Fperf)
  * [运行时核心](#%E8%BF%90%E8%A1%8C%E6%97%B6%E6%A0%B8%E5%BF%83)
    * [调度器](#%E8%B0%83%E5%BA%A6%E5%99%A8)
    * [执行栈](#%E6%89%A7%E8%A1%8C%E6%A0%88)
    * [内存分配器](#%E5%86%85%E5%AD%98%E5%88%86%E9%85%8D%E5%99%A8)
    * [垃圾收集器](#%E5%9E%83%E5%9C%BE%E6%94%B6%E9%9B%86%E5%99%A8)
    * [统计](#%E7%BB%9F%E8%AE%A1)
    * [内存模型](#%E5%86%85%E5%AD%98%E6%A8%A1%E5%9E%8B)
    * [ABI](#ABI)
    * [Misc](#Misc)
  * [标准库](#%E6%A0%87%E5%87%86%E5%BA%93)
    * [syscall](#syscall)
    * [os, io, io/fs, embed](#os%2C+io%2C+io%2Ffs%2C+embed)
    * [go/*](#go%2F%2A)
    * [sync](#sync)
      * [Map](#Map)
      * [Pool](#Pool)
      * [Mutex, RWMutex](#Mutex%2C+RWMutex)
      * [Groups](#Groups)
      * [atomic](#atomic)
    * [time](#time)
    * [context](#context)
    * [encoding](#encoding)
    * [image, x/image](#image%2C+x%2Fimage)
    * [Mobile](#Mobile)
    * [misc](#misc)
  * [未分类但相关的链接](#%E6%9C%AA%E5%88%86%E7%B1%BB%E4%BD%86%E7%9B%B8%E5%85%B3%E7%9A%84%E9%93%BE%E6%8E%A5)
  * [有趣的事实](#%E6%9C%89%E8%B6%A3%E7%9A%84%E4%BA%8B%E5%AE%9E)
  * [鸣谢](#%E9%B8%A3%E8%B0%A2)
  * [凭证](#%E5%87%AD%E8%AF%81)

---

## 免责声明

- 很多内容是对于公开信息的主观理解
- 可能出现实际错误或拼写错误。如果某些内容和您的观点相违背，请参考原始链接
- 非常欢迎关于新内容、错误修复、拼写修复的合并请求
- 使用 Github Issue 进行讨论

[Back To Top](#top)


## 资料来源
有很多探索Go历史设计的平台，以下是这些平台的官方网站：

- [blog.golang.org](https://blog.golang.org/)
- [dev.golang.org](https://dev.golang.org/)
- [talks.golang.org](https://talks.golang.org/)
- [golang.org/doc](https://golang.org/doc)
- [golang.org/pkg](https://golang.org/pkg)
- [github.com/golang/go](https://github.com/golang/go)
- [github.com/golang/proposal](https://github.com/golang/proposal)
- [github.com/golang/go/wiki](https://github.com/golang/go/wiki)
- [go-review.googlesource.com](https://go-review.googlesource.com/)
- [groups.google.com/g/golang-nuts](https://groups.google.com/g/golang-nuts)
- [groups.google.com/g/golang-dev](https://groups.google.com/g/golang-dev)
- [groups.google.com/g/golang-tools](https://groups.google.com/g/golang-tools)

[Back To Top](#top)


## 起源
Go 是由一个小组织以及语言社区的用户推动的。以下是一些核心贡献者，你可能会对他们杰出的工作感兴趣。


Go 的起源很吸引人。通过这些人举办的讲座，您可以了解到相关的口述历史和有趣的故事。


例如，您可以选择这些有关 Go 历史壮举的讲座作为起点（仅建议）：talk/rob2007, talk/rob2009, talk/rob2010b, talk/rob2010d, talk/rob2011a, talk/rob2013a, talk/rob2019, talk/robert2015, talk/russ2014, steve2019b, etc.
> YouTube 视频。

[Back To Top](#top)


## 贡献者
### 核心作者
Go 最初由 Rob, Robert 和 Ken 创建。Rob 跟 Robert 谈论 C++ 缓慢的编译时间， 而 Ken 刚好在隔壁办公室。之后 Ian 因为兴趣加入了这个项目，编写了 [gccgo](https://github.com/golang/gofrontend)。Rob 和 Ken 退休后，Robert 和 Ian 目前为 Go 添加通用方法。Russ 也是早期的核心贡献者，那时候他是谷歌的新人，Rob 从 [Plan 9](https://9p.io/plan9/) 项目回来的路上认识了他，邀请他加入 Go 团队。Russ 做了很多基础工作，包括早期的 Go 编译器、运行时以及 Go 1.5 启动时的巨大提升。Russ 现在是 Go 团队的技术领导。



- Rob Pike. (Robert C. Pike 科学硕士) [Website](http://herpolhode.com/rob/), [Blog](https://commandcenter.blogspot.com/), [GitHub](https://github.com/robpike), [Twitter](https://twitter.com/rob_pike), [Reddit](https://www.reddit.com/user/robpike).（退休）
   - 毕业院校：多伦多大学（理学学士），加州理工
   - [talk/rob2007](https://www.youtube.com/watch?v=hB05UFqOtFA) 编程语言的高级主题：并发/通过 Newsqueak 发送消息。2007年3月9日。
   - [talk/rob2009](https://changelog.com/podcast/3) Go 编程语言。2009年11月27日。
   - [talk/rob2010a](https://www.youtube.com/watch?v=jgVhBThJdXc) Go 编程。谷歌2010年网络开发者年会。2010年3月20日。
   - [talk/rob2010b](https://www.youtube.com/watch?v=-_DKfAn4pFA) Go 并发风格的起源。新兴语言集会2010。2010年7月21日。
   - [talk/rob2010c](https://www.youtube.com/watch?v=5kj5ApnhPAE) `public static void`，2010年 OSCON 大会。2010年7月22日。
   - [talk/rob2010d](https://www.youtube.com/watch?v=7VcArS4Wpqk) 另一种 Go 语言设计。2010年8月27日。
   - [talk/rob2011a](https://www.infoq.com/interviews/pike-concurrency/) 编程语言中的并行性和并发性。2011年2月17日。
   - [talk/rob2011b](https://www.infoq.com/interviews/pike-google-go/) Go 语言：并发、类型系统、内存管理和垃圾回收。2011年2月25日。
   - [talk/rob2011c](https://www.youtube.com/watch?v=HxaD_trXwRE) Go 词法分析。2011年8月30日。
   - [talk/rob2012a](https://www.youtube.com/watch?v=f6kdp27TYZs) Go 并发模式。谷歌2012年网络开发者年会。2012年7月2日。
   - [talk/rob2012b](https://www.youtube.com/watch?v=FTl0tl9BGdc) 什么学习Go？2012年9月12日。
   - [talk/rob2013a](https://www.youtube.com/watch?v=bj9T2c2Xk_s) Go 1 的路径图。2013年3月14日。
   - [talk/rob2013b](https://www.infoq.com/presentations/Go-Google/) Go 在谷歌。2013年4月13日。
   - [talk/rob2013c](https://changelog.com/podcast/100) Go 开发（Rob Pike 和 Andrew Gerrand)。2013年8月14日。
   - [talk/rob2013d](https://www.youtube.com/watch?v=cN_DpYBzKso) 并发不是并行。2013年10月20日。
   - [talk/rob2014a](https://www.youtube.com/watch?v=VoS7DsT1rdM) Hello Gophers! Gophercon 开幕演讲。2014年 Gophercon。2014年4月24日。
   - [talk/rob2014b](https://www.youtube.com/watch?v=PXoG0WX0r_E) 大数计算器实现。悉尼 Go 集会。2014年11月19日。
   - [talk/rob2015a](https://www.youtube.com/watch?v=cF1zJYkBW4A) 从 C 到 Go 的工具链变化。2015年 GopherFest。2015年6月10日。
   - [talk/rob2015b](https://www.youtube.com/watch?v=PAAkCSZUG1c) Go 箴言。GopherFest。2015年11月18日。
   - [talk/rob2015c](https://www.youtube.com/watch?v=rFejpH_tAHM) 简单是复杂的。2015年 dotGo。2015年12月2日。
   - [talk/rob2016a](https://www.youtube.com/watch?v=KINIAgRpkDA) Go 汇编的设计。2016年 GopherCon。2016年8月18日。
   - [talk/rob2016b](https://www.youtube.com/watch?v=sDTGhIqyMjo) 标识堆栈：接口学习。悉尼 Go 集会。2016年9月。
   - [talk/rob2017](https://www.youtube.com/watch?v=ENLWEfi0Tkg) Upspin 项目。2017 年 Gopherfest。2017年1月22日。
   - [talk/rob2018a](https://www.youtube.com/watch?v=_2NI6t2r_Hs) Unix 历史。2018年11月7日。
   - [talk/rob2018b](https://www.youtube.com/watch?v=RIvL2ONhFBI) Go 2 规划草案。悉尼 Go 集会。2018年11月13日。
   - [talk/rob2019a](https://changelog.com/gotime/100) Rob Pike 和 Robert Griesemer 创建 Go 语言。2019年9月10日。
   - [talk/rob2019b](https://www.youtube.com/watch?v=oU9cfQCxjpM) Rob Pike，Go 语言简史。2019年12月12日。
   - [talk/rob2020](https://evrone.com/rob-pike-interview) 一次 Rob Pike 采访。2020年。
   - [talk/rob2021](https://www.youtube.com/watch?v=YXV7sa4oM4I) Go 编程语言和环境。John Lions 卓越的演讲，新南威尔士大学。2021年3月27日。




- Robert Griesemer（博士）[GitHub](https://github.com/griesemer), [Twitter](https://twitter.com/robertgriesemer?lang=en)
   - 毕业院校：苏黎世联邦理工学院
   - [paper/robert1993](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.127.5290&rep=rep1&type=pdf) 矢量计算机编程语言。博士论文。
   - [talk/robert2012a](https://www.youtube.com/watch?v=on5DeUyWDqI) E2E: Erik Meijer 和 Robert Griesemer。2012年 Lang.NEXT Going Go。
   - [talk/robert2012b](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2012/Go-In-Three-Easy-Pieces) Go 的三个简单部分。2012年3月19日。
   - [talk/robert2012c](https://channel9.msdn.com/Events/Lang-NEXT/Lang-NEXT-2012/Panel-Native-Languages) 2012年 Lang.NEXT 专家组：原生语言。2012年4月11日。
   - [talk/robert2015](https://www.youtube.com/watch?v=0ReKdcpNyQg) Go 的演变。2015年 GopherCon，2015年7月28日。
   - [talk/robert2016a](https://www.youtube.com/watch?v=t-w6MyI2qlU) 和 Gom 的简单对话：声明别名（建议）。2016年 GopherCon，2016年10月9日。
   - [talk/robert2016b](https://www.youtube.com/watch?v=vLxX3yZmw5Q) 为你的设计做原型！2016年 dotGo，2016年11月29日。
   - [talk/robert2017](https://www.youtube.com/watch?v=KPk1UPihWtY) 开幕演讲：导出 Go。2017年 GopherCon SG，2017年3月29日。
   - [talk/robert2017](https://www.youtube.com/watch?v=UmwJsQTSEP8) Go 的概述。2017年4月28日。
   - [talk/robert2019](https://www.youtube.com/watch?v=i0zzChzk8KE) Go 10岁了！现在咋样？2019年 Gopherpalooza。2019年12月2日。
   - [talk/robert2020a](https://changelog.com/gotime/140) 泛型最终版：Robert Griesemer 和 Ian Lance Taylor。2020年1月21日。
   - [talk/robert2020b](https://www.gophercon.com/agenda/session/233094) 编写 Go（通用）。2020年11月11日。




- Ken Thompson（Kenneth Lane Thompson 科学硕士）（退休）
   - 毕业院校: 加州大学伯克利分校
   - [talk/ken1982a](https://www.youtube.com/watch?v=tc4ROCJYbm0) UNIX 系统：提高计算机生产效率。1982年
   - [talk/ken1982b](https://www.youtube.com/watch?v=XvDZLjaCJuw) UNIX 系统：提高计算机易用性。
   - [talk/ken1982c](https://www.youtube.com/watch?v=JoVQTPbD6UY) Ken Thompson 和 Dennis Ritchie 解释 UNIX（贝尔实验室）
   - [talk/ken1983](https://www.youtube.com/watch?v=LXZ1OL2U3lY) Ken Thompson 和 Dennis Ritchie：国家技术奖章。1996年。
   - [talk/ken2013](https://www.youtube.com/watch?v=dsMKJKTOte0) 系统架构、设计、工程和验证。2013年1月17日。
   - [talk/ken2019a](https://www.youtube.com/watch?v=g3jOJfrOknA) Thompson 和 Ritchie 的故事。2019年2月18日。
   - [talk/ken2019b](https://www.youtube.com/watch?v=EY6q5dv_B-o) Brian Kernighan 采访 Ken Thompson。2019年VCF East，2019年3月4日。




- Ian Taylor（Ian Lance Taylor 学士） [Website](https://www.airs.com/ian/), [GitHub](https://github.com/ianlancetaylor), [Quora](https://www.quora.com/profile/Ian-Lance-Taylor)
   - 毕业院校：耶鲁大学
   - [talk/ian2007](https://www.youtube.com/watch?v=gc78olyguqA) GCC：当前主题和未来方向。2007年2月27日。
   - [talk/ian2018](https://www.youtube.com/watch?v=LqKOY_pH8u0) 向 Go 2 转变。2018年 Gopherpalooza，2018年10月24日。
   - [talk/ian2019a](https://www.youtube.com/watch?v=WzgLqE-3IhY) Go 的泛型。2019年 GopherCon。2019年8月27日。
   - [talk/ian2019b](https://changelog.com/gotime/98) Go 的泛型。2019年8月27日。
   - [talk/ian2020](https://www.youtube.com/watch?v=yoZ05GG8aLs) Go 语言与 Ian Lance Taylor。CppCast，2020年8月9日。




- Russ Cox（Russell Stensby Cox 博士） [Website](https://swtch.com/~rsc/), [Blog](https://research.swtch.com/), [GitHub](https://github.com/rsc), [Twitter](https://twitter.com/_rsc), [Reddit](https://old.reddit.com/user/rsc)
   - 毕业院校：麻省理工学院
   - [paper/russ2008](https://pdos.csail.mit.edu/~rsc/) 面向扩展的编译器。博士论文。
   - [talk/russ2009](https://www.youtube.com/watch?v=wwoWei-GAPo) Go 语言广告。2009年11月10日。
   - [talk/russ2012](https://www.youtube.com/watch?v=MzYZhh6gpI0) Go 语言之旅。2012年1月24日。
   - [talk/russ2012b](https://www.youtube.com/watch?v=dP1xVpMPn8M) Acme 编辑器之旅。2012年9月17日。
   - [talk/russ2014](https://www.youtube.com/watch?v=QIE5nV5fDwA) 从 C 到 Go。2014年 GopherCon。2014年3月18日。
   - [talk/russ2015](https://www.youtube.com/watch?v=XvZOdpd_9tc) Go、开源、社区。2015年 GopherCon。2015年1月28日。
   - [talk/russ2016](https://www.youtube.com/watch?v=h6Cw9iCDVcU) 使用 Go 重构代码库。2016年12月5日。
   - [talk/russ2017](https://www.youtube.com/watch?v=0Zbh_vmAKvk) Go 的特性。2017年 GopherCon.2017年7月24日。
   - [talk/russ2018a](https://changelog.com/gotime/77) Go 的依赖和特性。2018年4月19日。
   - [talk/russ2018b](https://changelog.com/gotime/bonus-77) Go 和 WebAssembly (Wasm). 2018年4月19日。
   - [talk/russ2018c](https://www.youtube.com/watch?v=F8nrpe0XWRg) 开幕演讲：Go 和 版本。2018年 GopherConSG。2018年3月5日。
   - [talk/russ2018d](https://www.youtube.com/watch?v=6wIP3rO6On8) Go 2 草稿宣布。2018年8月28日。
   - [talk/russ2019](https://www.youtube.com/watch?v=kNHo788oO5Y) Go 2 路径图。2019年 GopherCon。2019年8月27日。
   - [talk/russ2020a](https://golang.org/s/go-build-video) Go 编译设计草稿。2020年6月30日。
   - [talk/russ2020b](https://golang.org/s/draft-iofs-video) os/fs 包设计草稿。2020年7月21日。
   - [talk/russ2020c](https://golang.org/s/draft-embed-video) //go:embed 包设计草稿。2020年7月21日。

[Back To Top](#top)


### 编译运行时团队
Dmitry 来自谷歌动态测试工具团队而不是 Go 团队。他编写了可伸缩的 goroutine 调度器，很多的性能提升，早期的同步，race 检测，Go 运行时阻塞分析器。Austin 在攻读博士期间是谷歌的实习生，在 Go 项目早期工作。后来，在他结束学生生涯后加入了 Go 团队，和 Rick 一起从事 Go 并发 GC 工作。他也在做抢占式调度器和链接器的工作。他现在是 Go 语言编译运行时团队的领导人。Keith 和 David 一起专注于 Go 编译器后端，尤其是当前的 SSA 后端。Michael 是 Go 团队的新成员，他的主要工作是运行时内存系统，例如内存分配器的重构和运行时指标。



- Dmitry Vyukov（Дмитрий Вьюков 理学硕士）[Website](http://www.1024cores.net/), [GitHub](https://github.com/dvyukov), [Twitter](https://twitter.com/dvyukov)
   - 毕业院校：莫斯科国立鲍曼技术大学
   - [talk/dmitry2014](https://www.youtube.com/watch?v=QEhpLb0UCfE) 用 Go 编写一个函数式的可靠快速的 Web 应用。2014年9月25日。
   - [talk/dmitry2015a](https://www.youtube.com/watch?v=Ef7TtSZlmlk) 和 Dmitry Vyukov 的谈论：Go 语言和信息安全。
   - [talk/dmitry2015b](https://www.youtube.com/watch?v=a9xrxRsIbSU) Go 动态工具。2015年 GopherCon。2015年7月28日。
   - [talk/dmitry2016](https://www.youtube.com/watch?v=9cpN2r22sLE) Dmitry Vyukov 的采访。2016年6月1日。
   - [talk/dmitry2017](https://www.youtube.com/watch?v=FD30Qzd6ylk) 模糊测试：新单元测试。2017年俄罗斯 C++。2017年3月10日。
   - [talk/dmitry2018a](https://www.youtube.com/watch?v=EJVp13f_aIs) 模糊测试：新单元测试。俄罗斯 GopherCon。2018年4月2日。
   - [talk/dmitry2018b](https://www.youtube.com/watch?v=qrBVXxZDVQY) Syzbot 和数千个内核 Bug 的故事。2018年9月1日。
   - [talk/dmitry2019](https://www.youtube.com/watch?v=-K11rY57K7k) Go 调度器：以轻量级并发实现编程语言。2019年10月14日。
   - [talk/dmitry2020](https://www.youtube.com/watch?v=YwX4UyXnhz0) syzkaller：连续覆盖引导内核模糊器。2020年BlueHat IL。2020年2月13日。




- Austin Clements（Austin T. Clements 博士）[GitHub](https://github.com/aclements), [Scholar](https://scholar.google.com/citations?user=MKDtxN4AAAAJ)
   - 毕业院校：麻省理工学院
   - [paper/autin2014](https://pdos.csail.mit.edu/papers/aclements-phd.pdf) 可伸缩交换规则：为多核处理器设计可伸缩软件。博士论文。
   - [talk/austin2020](https://www.gophercon.com/agenda/session/233441) 原谅中断：Go 1.14 中的循环抢占。2020年11月12日。




- Richard Hudson（Richard L. Hudson 理学硕士）[GitHub](https://github.com/RLH)（退休）
   - 毕业院校：马萨诸塞大学阿默斯特分校
   - [paper/rick](https://dl.acm.org/profile/81100566849/publications?Role=author) List 研究
   - [talk/rick2015](https://www.youtube.com/watch?v=aiv1JOfMjm0) Go 垃圾回收：解决延迟问题。2015年 GopherCon，2015年7月8日。
   - [talk/rick2015b](https://www.infoq.com/interviews/hudson-go-gc/) Go 垃圾回收：Rick Hudson。2015年12月21日




- Keith Randall（Keith H. Randall 博士）[GitHub](https://github.com/randall77)
   - 毕业院校：麻省理工学院
   - [paper/1998keith](http://supertech.csail.mit.edu/papers/PPoPP95.pdf) Cilk 语言：高校多线生程计算。博士论文。
   - [talk/keith2016](https://www.youtube.com/watch?v=Tl7mi9QmLns) Map 内部实现。2016年 GopherCon，2016年7月12日。
   - [talk/keith2017](https://www.youtube.com/watch?v=uTMvKVma5ms) 用 SSA 生成更好的机器码。2017年 GopherCon，2017年7月24日。




- David Chase（David Chase 博士）[Website](http://chasewoerner.org/resume.html), [Block](https://dr2chase.wordpress.com/), [GitHub](https://github.com/dr2chase), [Twitter](https://twitter.com/dr2chase), [Scholar](https://dblp.org/pid/51/3488.html)
   - 毕业院校：莱斯大学
   - [paper/1987david](http://www.chasewoerner.org/dissertation.pdf) 垃圾回收和其他优化。博士论文。
   - [talk/david2017](https://changelog.com/gotime/52) Go 编译器详述。2017年7月20日。



- Dan Scales.
   - [talk/dan2020](https://www.gophercon.com/agenda/session/233397) 实现更快的 defer。2020年11月11日。




- Michael Knyszek. [Website](https://www.ocf.berkeley.edu/~mknyszek/), [GitHub](https://github.com/mknyszek)
   - [talk/michael2020](https://www.gophercon.com/agenda/session/233086) 提升 Go 内存管理器的 RAM 和 CPU 效率。




- Than McIntosh. [GitHub](https://github.com/thanm)



- Cherry Zhang. [GitHub](https://github.com/cherrymui)



仓库、工具、安全、社区

- Andrew Gerrand [GitHub](https://github.com/adg), [Twitter](https://twitter.com/enneff)
- Brad Fitzpatrick [Website](https://bradfitz.com/), [GitHub](https://github.com/bradfitz), [Twitter](https://twitter.com/bradfitz)
   - [talk/brad2010](https://www.youtube.com/watch?v=c4znvD-7VDA) 编写 zippy 安卓应用。2010年谷歌网络开发者大会。2010年3月27日。
   - [talk/brad2011](https://www.youtube.com/watch?v=OExxMUVVjRM) Palestra Brad FitzPatrick：Perl 工厂。第1部分，2011年3月12日。
   - [talk/brad2012](https://www.youtube.com/watch?v=zHXoDB07Iwg) LiveJournal 幕后：缩放的故事时间。2012年3月21日。
   - [talk/brad2013a](https://www.youtube.com/watch?v=sYukPc0y_Ro) 令我高兴的软件。2013年 dotScale。2013年7月30日。
   - [talk/brad2013b](https://www.youtube.com/watch?v=kJ68OWnlycQ) LiveJournal 莫斯科办公室。Brad Fitzpatrick。2013年12月15日。
   - [talk/brad2014a](https://www.youtube.com/watch?v=Ds0KXj8ohR8) Camlistore。Go 集会。2014年1月
   - [talk/brad2014b](https://www.youtube.com/watch?v=D6okO8Qzusk) Camlistore 和标准仓库。2014年 GopherCon。2014年3月18日。
   - [talk/brad2014c](https://www.youtube.com/watch?v=4KFTacxqkcQ) Gopher 的状态。2014年 dotGo，2014年10月10日。
   - [talk/brad2014d](https://www.youtube.com/watch?v=VLciNxKAKyc) 采访。2014年 dotGo，2014年。
   - [talk/brad2015a](https://www.youtube.com/watch?v=yG-UaBJXZ80) Andrew 和 Brad 的工作：一个 HTTP/2 客户端。2015年2月25日。
   - [talk/brad2015b](https://www.youtube.com/watch?v=gukAZO1fqZQ) HTTP/2 和 Go。2015年伦敦 Go 集会。2015年3月4日。
   - [talk/brad2015c](https://www.youtube.com/watch?v=gAfoLwog_MA) Go 里的 HTTP/2。2015年欧洲开源开发者会议 Go 开发室。2015年3月4日。
   - [talk/brad2015d](https://www.youtube.com/watch?v=rHBbqjWCGq8) Gopher 的状态。慕尼黑 Gopher。2015年4月18日。
   - [talk/brad2015e](https://www.youtube.com/watch?v=stram5J144s) Go 语言。2015年波兰 devoxx。2015年6月。
   - [talk/brad2015f](https://www.youtube.com/watch?v=1rZ-JorHJEY) Andrew 和 Brad 的工作：tiop.golang.org。2015年1月8日。
   - [talk/brad2015g](https://www.youtube.com/watch?v=mj-1wscEQO8) 快速采访：Go 的 HTTP/2。2015年 GopherCon。2015年8月19日。
   - [talk/brad2015h](https://www.youtube.com/watch?v=xxDZuPEgbBU) Go 的分析及优化。2015年8月27日。
   - [talk/brad2016a](https://www.youtube.com/watch?v=4Dr8FXs9aJM) Go 1.6 介绍：逐渐接近单调。2016年印度 GopherCon。2016年3月29日。
   - [talk/brad2016b](https://www.youtube.com/watch?v=4yFb-b5GYWc) 快速谈话：我的视频和安全系统。2016年 GopherCon。2016年7月。
   - [talk/brad2016c](https://www.youtube.com/watch?v=8cQcPnzfkLk) Go 1.7。西雅图 Go 集会。2016年8月25日。
   - [talk/brad2016d](https://www.youtube.com/watch?v=18kmlJvR6Bk) Go 开源项目的内部结构。2016年 GothamGo。2016年12月15日。
   - [talk/brad2017](https://www.youtube.com/watch?v=4fWqcOubYQ0) 我的前半生在开源项目。2017年3月19日。
   - [talk/brad2018a](https://www.youtube.com/watch?v=ZCB-g2B4Y5A) Go：回顾和展望。2018年4月2日。
   - [talk/brad2018b](https://www.youtube.com/watch?v=rWJHbh6qO_Y) Go 1.11 和未来。2018年8月26日。
   - [talk/brad2018c](https://www.youtube.com/watch?v=69Zy77O-BUM) 快速谈话：核选项，`go test -run=InQemu`。2018年 GopherCon，2018年11月11日。
   - [talk/brad2019](https://www.youtube.com/watch?v=BRSam0xQJKY) Brad Fitzpatrick 更喜欢Go，而非 C, C++, Rust, Perl, Python, Ruby, JavaScript 和 Java。2019年11月28日。




- Bryan C. Mills. [GitHub](https://github.com/bcmills)
   - [talk/bryan2017](https://www.youtube.com/watch?v=C1EtfDnsdDs) 快速谈话：sync.Map 的概述。2017年 GopherCon，2017年7月24日。
   - [talk/bryan2018](https://www.youtube.com/watch?v=5zXAHh5tJqQ) 重新思考经典的并发模式。2018年 GopherCon，2018年9月14日。




- Steve Francia. [Website](https://spf13.com/), [GitHub](https://github.com/spf13), [Twitter](https://twitter.com/spf13).
   - [talk/steve2019a](https://spf13.com/presentation/what-should-a-modern-practical-programming-language-look-like/) 现代实用编程语言应该是什么样子。Langding Festival，2019年4月4日
   - [talk/steve2019b](https://spf13.com/presentation/the-legacy-of-go/) Go 的遗产。Go Lab，2019年10月22日。




- Jonathan Amsterdam.
   - [talk/jonathan2020](https://www.gophercon.com/agenda/session/233432) Errors 方面。2020年11月13日。




- Daniel Martí. [Website](https://mvdan.cc/), [GitHub](https://github.com/mvdan), [Twitter](https://twitter.com/mvdan_)




- Nigel Tao. [GitHub](https://github.com/nigeltao), Twitter



- Filippo Valsorda. GitHub,



- Michael Matloob. [GitHub](https://github.com/matloob), [Twitter](https://twitter.com/matloob)



- Jay Conrod. [Website](https://jayconrod.com/), [Twitter](https://twitter.com/jayconrod)



- Dave Cheney. [Website](https://dave.cheney.net/), [GitHub](https://github.com/davecheney), [Twitter](https://twitter.com/davecheney)



- Sam Boyer. [GitHub](https://github.com/sdboyer), [Twitter](https://twitter.com/sdboyer)



- Fillippo Valsorda. [Website](https://filippo.io/), [GitHub](https://github.com/FiloSottile), [Twitter](https://twitter.com/FiloSottile)
   - [talk/filo2016](https://www.youtube.com/watch?v=lhMhApWQp2E) 从 cgo 回到 Go。2016年7月12日。
   - [talk/filo2017](https://speakerdeck.com/filosottile/calling-rust-from-go-without-cgo-at-gothamgo-2017) 不使用 cgo，从 Go 到 Rust。
   - [talk/filo2018](https://speakerdeck.com/filosottile/why-cgo-is-slow-at-capitalgo-2018) 为什么 cgo 太慢。2018年 CapitalGo。
   - [talk/speakerdeck](https://speakerdeck.com/filosottile?page=1)
   

> 更多人和谈话应该被添加……

[Back To Top](#top)


### 小组会谈

- [talk/goteam2012](https://www.youtube.com/watch?v=sln-gJaURzk) Go 团队集会。2012年谷歌网络开发者大会。2012年7月2日。
- [talk/goteam2013](https://www.youtube.com/watch?v=p9VUCp98ay4) 和 Go 团队的炉边谈话。2013年谷歌网络开发者大会。2013年3月18日。
- [talk/goteam2014](https://www.youtube.com/watch?v=u-kkf76TDHE) 和 Blake Mizerany 在 Gophers 工作室。2014年 GopherCon。2014年3月21日。
- [talk/goteam2019](https://www.youtube.com/watch?v=3yghHvvZQmA) 认识作者：Go 语言。2019年 Cloud Next 大会，2019年4月10日。
- [talk/goteam2020a](https://www.youtube.com/watch?v=gJxvkOHpTSM) 2020年 GoLab：Go 团队问答。2020年10月22日。
- [talk/goteam2020b](https://www.youtube.com/watch?v=BNHwHLNLjLs) 2020年 GopherCon：Go 团队问答。2020年11月16日。

[Back To Top](#top)


## 时间线
时间线能帮助你确定关联到 Go 版本文档。

| Version | Release Date |
| --- | --- |
| Go 1 | 2012.03.28 |
| Go 1.1 | 2013.05.13 |
| Go 1.2 | 2013.12.01 |
| Go 1.3 | 2014.06.18 |
| Go 1.4 | 2014.12.10 |
| Go 1.5 | 2015.08.19 |
| Go 1.6 | 2016.02.17 |
| Go 1.7 | 2016.08.15 |
| Go 1.8 | 2017.02.16 |
| Go 1.9 | 2017.08.24 |
| Go 1.10 | 2018.02.16 |
| Go 1.11 | 2018.08.24 |
| Go 1.12 | 2019.02.25 |
| Go 1.13 | 2019.09.03 |
| Go 1.14 | 2020.02.25 |
| Go 1.15 | 2020.08.11 |
| Go 1.16 | 2021.02.16 |

历史版本说明可以提供一些大概信息：

- [doc/go1release](https://golang.org/doc/devel/release.html) Go 发布历史
- [doc/go1prerelease](https://golang.org/doc/devel/pre_go1.html) Go 1 预发布历史
- [doc/go0release](https://golang.org/doc/devel/weekly.html) 每周发布历史（Go 1 之前）

[Back To Top](#top)


## 语言设计
### Misc

- [design/go0initial](https://github.com/golang/go/blob/18c5b488a3b2e218c0e0cf2a7d4820d9da93a554/doc/go_spec) Rob Pike、Robert Griesemer、Ken Thompson，Go 语言注释规范，2008年3月3日。
- [design/go0spec0](https://github.com/golang/go/blob/e6626dafa8de8a0efae351e85cf96f0c683e0a4f/doc/go_lang.txt) Go 编程语言，语言规范。2008年3月7日。
- [design/go0semicolon](https://golang.org/s/semicolon-proposal) Rob Pike，Go 中的分号。2009年12月10日。
- [design/go11func](https://golang.org/s/go11func) Russ Cox，Go 1.11 函数调用。2013年2月。
- [design/go11return](https://golang.org/s/go11return) Russ Cox，Go 函数最后的 Return 要求。2013年3月。
- [design/go12nil](https://golang.org/s/go12nil) Russ Cox，Go 1.2 字段选择和空检查。2013年7月。
- [doc/go13todo](https://golang.org/s/go13todo) Go 1.3 待完成列表。
- [doc/goatgoogle](https://talks.golang.org/2012/splash.article#TOC_12.) Rob Pike，谷歌 Go 使用 - 语义。2012年10月25日。
- [doc/makego](https://talks.golang.org/2015/how-go-was-made.slide) Andrew Gerrand，Go 怎样创建的。2015年7月9日。
- [discuss/go1preview](https://docs.google.com/document/pub?id=1ny8uI-_BHrDCZv_zNBSthNKAMX_fR_0dc6epA6lztRE) Russ Cox。Go 1 预览版。
- [design/overlapping-interfaces](https://golang.org/design/6977-overlapping-interfaces) Robert Griesemer 提案：允许将接口嵌入折叠方法集。2019年10月16日。
   - [issue/6977](https://golang.org/issue/6977) 规则：允许嵌入重叠接口
- [design/struct-conversion](https://golang.org/design/16085-conversions-ignore-tags) Robert Griesemer 提案：struct 类型转换忽略 tag
   - [issue/16085](https://golang.org/issue/16085) 提案：struct 类型转换忽略 tag
- [design/go2trans](https://golang.org/design/28221-go2-transitions) Ian Lance Taylor 提案：Go 2 变迁。2018年10月15日。
- [design/signed-int-shift](https://golang.org/design/19113-signed-shift-counts) Robert Griesemer 提案：在 Go 2 中允许有符号整数进行移位计算。2019年1月17日。
   - [issue/19113](https://golang.org/issue/19113) 提案：规范：允许有符号整数进行移位计算。
- [design/number-literal](https://golang.org/design/19308-number-literals) Russ Cox、Robert Griesemer 提案：Go 2 数学字面量
   - [issue/12711](https://golang.org/issue/12711) 提案：不同的八进制字面量表示
   - [issue/19308](https://golang.org/issue/19308) 提案：规范：二进制字面量
   - [issue/28493](https://golang.org/issue/28493) 提案：允许在数字中使用 `_`
   - [issue/29008](https://golang.org/issue/29008) 提案：Go 2 十六进制浮点数
- [issue/33502](https://golang.org/issue/33502) 提案：审查会议记录
- [issue/33892](https://golang.org/issue/33892) 提案：审查 Go 2 会议记录
- [issue/19623](https://golang.org/issue/19623) 提案：规范：更改 int 为任意精度
- [issue/19367](https://golang.org/issue/19367) unsafe：添加 `Slice(ptr *T, len anyIntegerType) []T`
- [issue/40481](https://golang.org/issue/40481) unsafe：添加函数

[Back To Top](#top)


### 切片 (1.2)

- [design/read-only-slices](https://docs.google.com/document/d/1UKu_do3FRvfeN5Bb1RxLohV-zBOJWTzX0E8ZU1bkqX0/edit#heading=h.2wzvdd6vdi83) Brad Fitzpatrick，只读切片。2013年3月13日。
- [design/read-only-slices-russ](https://docs.google.com/document/d/1-NzIYu0qnnsshMBpMPmuO21qd8unlimHgKjRD9qwp2A/edit) Russ Cox，只读切片的评价。2013年3月。
- [design/go12slice](https://golang.org/s/go12slice) Russ Cox，Go 切片的容量。2013年6月。
- [design/multidim-slice](https://golang.org/design/6282-table-data) Brendan Tracey，提案：多维切片。2016年11月17日。
- [issue/41239](https://golang.org/issue/41239) 运行时：`append`的增长让人意外。
- [discuss/why-slice-grow](https://groups.google.com/g/golang-nuts/c/UaVlMQ8Nz3o) 为什么切片长度会在 1024 后以25%的速度增长？
- [cl/347917](https://golang.org/cl/347917) 运行时：让切片增长公式更平滑。
   - [doc/cl-347917-graph](https://docs.google.com/document/d/1JQvV6vyAYdHhIboY-zAwK06OXZjxHrUhOFeG38MuJ94/edit?resourcekey=0-L5OsHqwZZBxvjfK0dwsyVQ) CL 347917 图表

[Back To Top](#top)


### 包管理 (1.4, 1.5, 1.7)

- [design/go14internal](https://golang.org/s/go14internal) Russ Cox，Go 1.4 `Internal` 包，2014年6月。
- [design/go14nopkg](https://golang.org/s/go14nopkg) Russ Cox，Go 1.4 `src/pkg` → `src`，2014年6月。
- [design/go14customimport](https://golang.org/s/go14customimport) Russ Cox. Go 1.4 自定义路径导入检查。2014年6月。
- [design/go15vendor](https://golang.org/s/go15vendor) Russ Cox，Go 1.5 Vendor 实验，2015年7月。
- [design/go17binarypkg](https://golang.org/design/2775-binary-only-packages) Russ Cox，提案：纯二进制包，2016年4月24日。
   - [issue/2775](https://golang.org/issue/2775) cmd/go：当二进制可用但源文件丢失。

[Back To Top](#top)


### 类型别名 (1.9)

- [design/type-alias](https://golang.org/design/18130-type-alias) Russ Cox、Robert Griesemer 提案：类型别名。2016年12月16日。
   - [talk/type-alias](https://www.youtube.com/watch?v=t-w6MyI2qlU) 2016年 GopherCon —— 快速谈话：Robert Griesemer - Go 别名说明，建议。2016年10月9日。
   - [issue/16339](https://golang.org/issue/16339) 提案：Go 别名说明。
   - [issue/18130](https://golang.org/issue/18130) 所有：支持渐进式代码修复，同时在包之间移动类型。
- [talk/refactor-video](https://www.youtube.com/watch?v=h6Cw9iCDVcU) Russ Cox，重构代码库（在 Go 的帮助下）。2016年 GothamGo，2016年11月18日。
   - [doc/refactor](https://talks.golang.org/2016/refactor.article) Russ Cox，重构代码库（在 Go 的帮助下）。

[Back To Top](#top)


### Defer (1.14)

- [design/open-defer](https://golang.org/design/34481-opencoded-defers) Dan Scales、Keith Randall、and Austin Clements 提案：通过内联代码进行低成本 defer，以及额外管理 panic。2019年9月23日。
   - [issue/6980](https://golang.org/issue/6980) cmd/compile：在栈帧中分配一些 defer。
   - [issue/14939](https://golang.org/issue/14939) 运行时：defer 很慢
- 未解决的 `defer recover()`边缘案例
   - [issue/10458](https://golang.org/issue/10458) 规范：panic 在极端情况下的语义
   - [issue/23531](https://golang.org/issue/23531) 规范：在嵌套 defer 中使用 `recover()`
   - [issue/26275](https://golang.org/issue/26275) 规范：在使用 defer 的函数中记录调用者的行为
   - [issue/34530](https://golang.org/issue/34530) 规范：澄清何时调用 recover 停止 panic 
   - [cl/189377](https://golang.org/cl/189377) 规范：列举 recover 行为和递归 panic 的详情。 

[Back To Top](#top)


### 错误值 (1.13)

- [doc/err2011](https://blog.golang.org/error-handling-and-go) Andrew Gerrand，Go 错误处理。2011年7月。
- [doc/err-values](https://blog.golang.org/errors-are-values) Rob Pike。错误是值，2015年1月。
- [doc/err-philosophy](https://dave.cheney.net/paste/gocon-spring-2016.pdf) Dave Cheney，我处理错误的理念。2016年4月。
- [doc/err-gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) Dave Cheney，不要只检查错误，而是优雅处理。2016年4月。
- [doc/err-stacktrace](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package) Dave Cheney，堆栈跟踪和错误包，2016年6月12日。
- [doc/err-upspin](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html) Rob Pike，Upspin 的错误处理。2017年12月6日。
- [doc/err-work](https://blog.golang.org/go1.13-errors) Damien Neil 和 Jonathan Amsterdam，Go 1.13 中处理错误，2019年10月17日。
- [design/err-handling-overview](https://golang.org/design/go2draft-error-handling-overview) Russ Cox，错误处理——问题综述。
- [doc/err-value-faq](https://github.com/golang/go/wiki/ErrorValueFAQ) Jonathan Amsterdam 和 Bryan C. Mills 错误值：常见问题。2019年8月。
- [design/err-handle-check](https://golang.org/design/go2draft-error-handling) Marcel van Lohuizen，错误处理——设计草稿。2018年8月27日。
- [design/err-try](https://golang.org/design/32437-try-builtin) Robert Griesemer 提案：Go 内置错误检查函数 `try`
   - [issue/32437](https://golang.org/issue/32437#issuecomment-512035919) 提案：Go 内置错误检查函数 `try`。判定回复。
- [design/err-inspect](https://golang.org/design/go2draft-error-inspection) Jonathan Amsterdam、Damien Neil，错误检查——设计草稿，2018年8月27日。
- [design/err-values-overview](https://golang.org/design/go2draft-error-values-overview) Russ Cox，错误处理——问题综述。2018年8月27日。
- [design/error-values](https://golang.org/design/29934-error-values) Jonathan Amsterdam、Russ Cox、Marcel van Lohuizen、Damien Neil 提案：Go 2 错误检查。2019年1月25日。
   - [issue/29934](https://golang.org/issue/29934) Jonathan Amsterdam 提案：Go 2 错误检查。2019年1月25日。
   - [issue/29934-decision](https://golang.org/issue/29934#issuecomment-489682919) Damien Neil。Go 1.13 关于错误值的决定。2019年3月6日。
   - [issue/29934-russ](https://golang.org/issue/29934#issuecomment-490087200) Russ Cox，关于“提案：Go 2 错误检查”的回复。2019年3月7日。
- [design/err-print](https://golang.org/design/go2draft-error-printing) Marcel van Lohuizen，错误打印——设计草稿。2018年8月27日。
   - [issue/30468](https://golang.org/issue/30468) errors: 新版本的性能衰退。
- [issue/40432](https://golang.org/issue/40432) 语言：Go 2: 错误处理的讨论
- [issue/40776](https://golang.org/issue/40776) 提案：动态错误忽略的检测器
- [issue/41198](https://golang.org/issue/41198) 提案：errors：添加 ErrUnimplemented 作为失败接口方法的标准方式。

[Back To Top](#top)


### 通道、选择器

- [design/lockfree-channels](https://docs.google.com/a/google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub) Dmitry Vyukov，Go 通道。2014年1月28日。
   - [issue/8899](https://golang.org/issue/8899) 运行时：无锁通道。
   - [discuss/lockfree-channels](https://groups.google.com/g/golang-dev/c/0IElw_BbTrk/m/cGHMdNoHGQEJ) 无锁通道的更新。
   - [cl/112990043](https://codereview.appspot.com/112990043/) 运行时：选择器中的锁粒度
   - [cl/110580043](https://codereview.appspot.com/110580043/) 运行时：在非阻塞通道操作中添加快速路径
- [issue/8898](https://golang.org/issue/8898) 运行时：特殊案例，计时器通道。
- [issue/37196](https://golang.org/issue/37196) 时间：在停止或重置返回后，让 Timer/Ticker 通道不接受旧值。
- [issue/8903](https://golang.org/issue/8903) 运行时：让通道更快生成。
- [issue/21806](https://golang.org/issue/21806) 选择语句的公平性。
- [issue/40410](https://golang.org/issue/40410) 从 scase 中去除不必要的字段
- [issue/40641](https://golang.org/issue/40641) 运行时，在堆栈收缩和通道发送/接受之间的竞争导致错误的 sudog 值
- [issue/37350](https://golang.org/issue/37350) 反射：如果数组长度大于 1<< 16，请 panic

[Back To Top](#top)


### 泛型

- [doc/generics-discuss](https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/edit#heading=h.vuko0u3txoew) Go 泛型讨论摘要
- [doc/generics-dilemma](https://research.swtch.com/generic) Russ Cox，对泛型很犹豫。2009年12月3日。
- [design/type-functions](https://golang.org/design/15292/2010-06-type-functions) Ian Lance Taylor，类型函数。2010年6月。
- [design/generalized-types](https://golang.org/design/15292/2011-03-gen) Ian Lance Taylor，通用类型。2011年3月。
- [design/code-gen](https://docs.google.com/document/pub?id=1IXHI5Jr9k4zDdmUhcZImH59bOUK0G325J1FY6hdelcM) Russ Cox，Go 中动态代码生成的替代。2012年9月。
- [design/generalized-types2](https://golang.org/design/15292/2013-10-gen) Ian Lance Taylor，Go 中的通用类型，2013年10月。
- [design/type-parameters](https://golang.org/design/15292/2013-12-type-params) Ian Lance Taylor，Go 中的类型参数，2013年12月。
- [design/compile-time-function](https://golang.org/design/15292/2016-09-compile-time-functions) Bryan C. Mills，编译时函数和一等类型。2016年12月。
- [design/should-generics](https://github.com/golang/proposal/blob/b571c3273d2c6988d24a22dd1c529387ff05962a/design/15292-generics) Ian Lance Taylor，Go 应该有泛型。2011年1月。
- [design/should-generics2](https://golang.org/design/15292-generics) Ian Lance Taylor，Go 应该有泛型。更新：2016年4月。
   - [issue/15292](https://golang.org/issue/15292) 提案：规范：通用编程工具。
- [design/generics-overview](https://golang.org/design/go2draft-generics-overview) Russ Cox，泛型——问题综述。2018年8月27日。
- [design/contracts](https://golang.org/design/go2draft-contracts) Ian Lance Taylor, Robert Griesemer，合同——设计草案。2018年8月27日，更新：2019年7月31日。
   - [cl/187317](https://golang.org/cl/187317/) 原型实现
   - [paper/featherweight-go](https://arxiv.org/abs/2005.11710) Griesemer, Robert 等，轻量级 Go。arXiv 预印本 arXiv:2005.11710 (2020).
   - [talk/featherweight-go](https://www.youtube.com/watch?v=Dq0WFigax_c) Phil Wadler：轻量级 Go，2020年6月8日。
- [design/type-parameters2](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md) Ian Lance Taylor, Robert Griesemer。类型参数——设计草稿。2016年6月16日，更新：2020年8月28日。
   - [cl/dev.go2go](https://github.com/golang/go/blob/dev.go2go/README.go2go.md) dev.go2go 分支
   - [doc/type-check-readme](https://github.com/golang/go/tree/dev.go2go/src/go/types) 类型检查
   - [doc/type-check-notes](https://github.com/golang/go/blob/dev.go2go/src/go/types/NOTES) 这个文件服务器作为笔记或实现日志
- [discuss/generics-parenthesis](https://groups.google.com/g/golang-nuts/c/7t-Q2vt60J8) Robert，泛型和括号
- [issue/33232](https://golang.org/issue/33232) 提案：Go 2：为 interface{} 添加任意别名
- [discuss/generics-move](https://groups.google.com/g/golang-nuts/c/iAD0NBz3DYw) Ian，继续泛型设计。
- [discuss/generics-implementation](https://groups.google.com/g/golang-dev/c/OcW0ATRS4oM) 实现泛型
- [design/generics-implementation-stenciling](https://golang.org/design/generics-implementation-stenciling) 泛型实现——模板
- [design/generics-implementation-dictionaries](https://golang.org/design/generics-implementation-dictionaries) 泛型实现——字典
- [issue/43651](https://golang.org/issue/43651) 提案：规范：使用类型参数添加泛型编程
- [design/type-parameters3](https://golang.org/design/43651-type-parameters) 类型参数建议
- [issue/45346](https://golang.org/issue/45346) 提案：规范：泛型：使用类型集合删除类型中的类型关键字
- [issue/46477](https://golang.org/issue/46477) 提案：规范：泛型：别名中的类型参数
- 使用泛型的标准库
   - [issue/45458](https://golang.org/issue/45458) 提案：约束：定义标准类型参数约束的新包
   - [discuss/47319](https://github.com/golang/go/discussions/47319) 提案：约束：定义标准类型参数约束的新包（讨论）
   - [issue/45955](https://golang.org/issue/45955) 提案：切片：提供泛型切片函数的新包
   - [discuss/47203](https://github.com/golang/go/discussions/47203) 提案：切片：提供泛型切片函数的新包（讨论）
   - [discuss/47331](https://github.com/golang/go/discussions/47331) 提案：容器/集合：提供泛型集合类型的新包（讨论）
   - [discuss/47330](https://github.com/golang/go/discussions/47330) 提案：maps：提供泛型 map 函数的新包（讨论）
   - [issue/47657](https://golang.org/issue/47657) 提案：sync，sync/atomic：添加 PoolOf, MapOf, ValueOf

[Back To Top](#top)


## 编译工具链
### 编译

- [code/gc0initial](https://github.com/golang/go/tree/cb87526ce3531557ccf69969de4c8018956b10b5/src/c) Ken Thompson，Go 0 编译器初始版本，2008年3月28日。
- [code/6g](https://github.com/golang/go/commit/0cafb9ea3d3d34627e8f492ccafa6ba9b633a213) Rob Pike，Go 0 第一个完整的编译器，2008年6月4日。
- [design/go12symtab](https://golang.org/s/go12symtab) Russ Cox，Go 1.2 运行时符号信息，2013年7月。
- [design/go13compiler](https://golang.org/s/go13compiler) Russ Cox，Go 1.3+ 编译器大修，2013年12月。
- [design/go14generate](https://golang.org/s/go1.4-generate) Rob Pike，Go 生成：提案。
- [design/dev.cc](https://golang.org/s/dev.cc) Russ Cox，dev.cc 分支计划，2014年11月
- [design/go15bootstrap](https://golang.org/s/go15bootstrap) Russ Cox，Go 1.5 启动计划，2015年1月。
- [doc/escape-analysis](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/edit#) Dmitry Vyukov，Go 逃逸分析缺陷，2015年2月10日。
- [design/execmodes](https://golang.org/s/execmodes) Ian Lance Taylor，Go 执行模式，2015年8月（2016年1月更新）
- [design/go17ssa](https://golang.org/s/go17ssa) Keith Randall，Go 编译器新的 SSA 后端。2015年2月10日。
- [doc/compiler-optimization](https://github.com/golang/go/wiki/CompilerOptimizations) 编译器和运行时优化。
- [issue/6853](https://golang.org/issue/6853) 所有：二进制文件太大且不断增长。
- [design/go19inlining](https://golang.org/design/19348-midstack-inlining) David Lazar, Austin Clements，提案：Go 编译器中内联中层堆栈。
   - [issue/19348](https://golang.org/issue/19348) cmd/compile：允许中层堆栈内联
   - [talk/go19inliningtalk](https://golang.org/s/go19inliningtalk) David Lazar，Go 编译器中层堆栈内联
- [design/dwarf-inlining](https://golang.org/design/22080-dwarf-inlining) Than McIntosh，提案：在 Go 编译器中发出 DWARF 内联信息
   - [issue/22080](https://golang.org/issue/22080) cmd/compile：生成 DWARF 内联信息
- [issue/23109](https://golang.org/issue/23109) cmd/compile：重写逃逸分析
- [issue/27167](https://golang.org/issue/27167) cmd/compile：重命名很多东西
   - [doc/renames](https://docs.google.com/document/d/19_ExiylD9MRfeAjKIfEsMU1_RGhuxB9sA0b5Zv7byVI/edit) 提案：Go 1.12 工具链重命名
- GOEXPERIMENT=checkptr
   - [issue/22218](https://golang.org/issue/22218) 提案：添加 GOEXPERIMENT=checkptr
   - [issue/34964](https://golang.org/issue/34964) cmd/compile: 允许 -d=checkptr 作为 -race and/or -msan 的一部分?
   - [issue/34972](https://golang.org/issue/34972) 所有：通过 -d=checkptr 获取标准库构建
   - [discuss/checkptr](https://groups.google.com/forum/#!msg/golang-dev/SzwDoqoRVJA/Iozu8vWdDwAJ)
- [issue/37121](https://golang.org/issue/37121) 运行时：在 Intel 中，未对齐的跳转造成性能下降
- [issue/16798](https://golang.org/issue/16798) 提案：cmd/compile：仅为自递归添加尾部调用优化（已拒绝）
- [issue/22624](https://golang.org/issue/22624) 提案：Go 2：添加 become 语句来支持尾部调用（已拒绝）
- [design/64align](https://golang.org/design/36606-64-bit-field-alignment) Dan Scales，提案：在32位系统上使64位字段作为64位对齐，添加 //go:packed，//go:align 指令，2020年6月8日。
   - [issue/599](https://golang.org/issue/599) cmd/compile: 在32位系统上使64位字段作为64位对齐
   - [issue/36606](https://golang.org/issue/36606) 提案：cmd/compile: 在32位系统上使64位字段作为64位对齐，在结构体上添加 //go:packed，//go:align 指令
- [talk/gccgo](https://www.youtube.com/watch?v=U0w9eFunkX4) gccgo 的简要概述，“另一个” Go 编译器。2015年8月6日。
- [issue/28262](https://golang.org/issue/28262) cmd/compile：反馈导向优化

[Back To Top](#top)


### 连接器
Ken Thompson 编写了 Go 连接器，Russ 在 Go 1.3 中进行了一些大修。Austin 和 keith, Than, Cheery 和其他人一起对连接器返工，将在 Go 1.15 和 Go 1.16 中发布
[design/go13linker](https://golang.org/s/go13linker) Russ Cox，Go 1.3 连接器大修，2013年11月。

- [design/go116linker](https://golang.org/s/better-linker) Austin Clements，构建更好的 Go 连接器，2019年9月12日。

[Back To Top](#top)


### 调试器

- [doc/go13heapdump](https://golang.org/s/go13heapdump) 堆转储 Go 1.3
- [doc/go14heapdump](https://github.com/golang/go/wiki/heapdump14) 堆转储 Go 1.4
- [doc/go15heapdump](https://github.com/golang/go/wiki/heapdump15-through-heapdump17) 堆转储 Go 1.5 到 堆转储 Go 1.7
- [design/heap-viewer](https://golang.org/design/16410-heap-viewer) Michael Matloob，提案：Go 堆转储查看器，2016年7月20日。
   - [issue/16410](https://golang.org/issue/16410) x/tools/cmd/heapdump：创建一个堆转储查看器
- [design/profiler-labels](https://golang.org/design/17280-profile-labels) Michael Matloob，提案：支持 pprof 探查器标签，2017年3月15日。
   - [issue/17280](https://golang.org/issue/17280) pprof：添加探查器标签支持

[Back To Top](#top)


### 竞争检测器

- [issue/42598](https://golang.org/issue/42598) runtime: apparent false-positive race report for a buffered channel after CL 220419

[Back To Top](#top)


### 跟踪器

- [design/go15trace](https://golang.org/s/go15trace) Dmitry Vyukov，Go 执行跟踪器，2014年10月。
- [design/tracefmt](https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/preview#heading=h.yr4qxyxotyw) nduca@, dsinclair@，跟踪事件格式，2016年10月。

[Back To Top](#top)


### 锁分析

- [issue/38029](https://golang.org/issue/38029) x/build: linux-amd64-staticlockranking 持续失败。
   - [cl/192704](https://golang.org/cl/192704) 运行时：锁操作日志支持
   - [cl/207348](https://golang.org/cl/207348) 运行时：当运行 timer 时释放 timersLock
   - [cl/207846](https://golang.org/cl/207846) 运行时，在跟踪内存代码时避免潜在的死锁
   - [cl/207619](https://golang.org/cl/207619) 运行时：运行时的静态锁等级（通过 GOEXPERIMENT 开启）
   - [cl/222925](https://golang.org/cl/222925) cmd/go: 为任何已启用的 GOEXPERIMENT 定义构建标志
   - [cl/228417](https://golang.org/cl/228417) 运行时：将 Gscan 的合并/释放合并到锁排名顺序中
   - [cl/229480](https://golang.org/cl/229480) 运行时：添加几个新的 lock-rank partial order edges
   - [cl/231463](https://golang.org/cl/231463) 运行时：添加一个额外的 lock ranking partial edge
   - [cl/233599](https://golang.org/cl/233599) 运行时：添加 partial order edge 的锁（assistQueue -> mspanSpecial）
   - [cl/236137](https://golang.org/cl/236137) 运行时：为锁排名3个新的 partial orders
- [issue/40677](https://golang.org/issue/40677) 运行时：锁保持检查

[Back To Top](#top)


### 构建

- [design/go13nacl](https://go.dev/s/go13nacl) Russ Cox，Go 1.3 本地客户端支持。2013年10月。
- [design/gobuilder](http://go.dev/s/builderplan) Brad Fitzpatrick，Go 构建计划。2014年9月3日。
   - [discuss/gobuilder](https://groups.google.com/g/golang-dev/c/sdFD0-2Ed8k) 构建器运行的变更
- [design/go14android](https://go.dev/s/go14android) David Crawshaw，Go 支持安卓。2014年6月。
- [design/go-generate](https://go.dev/s/go1.4-generate) Rob Pike，Go 生成，2014年1月。
- [issue/13560](https://go.dev/issue/13560) 提案：构建：定义识别机器生成的文件的标准方式。
- [discuss/generatedcode](https://go.dev/s/generatedcode) Rob Pike 的关于 Issue 13560 的最终意见。
- [design/goenv](https://go.dev/design/30411-env) Russ Cox，提案：Go 命令行配置文件。2019年3月1日。
   - [issue/30411](https://go.dev/issue/30411) proposal: cmd/go: 添加 `go env -w`设置默认环境变量
- [design/go116build](https://go.dev/s/go-build-design%E2%80%8B) Russ Cox，防 Bug 构建约束——设计草稿。2020年6月30日。
   - [issue/41184](https://go.dev/issue/41184) cmd/go: 持续交流防 Bug //go:build
- Windows
   - [issue/41191](https://go.dev/issue/41191#issuecomment-690887303) 工具链指令
   - [discuss/win2000-golang-nuts](https://go.dev/s/win2000-golang-nuts) 反对删除对 Windows 2000(x86-32)的支持
- [design/wasm](https://docs.google.com/document/d/131vjr4DH6JFnb-blm_uRdaC0_Nv3OUwjEY5qVCxCup4/edit#heading=h.mjo1bish3xni) Richard Musiol，Go WebAssembly 的架构。2018年2月28日。
- [design/wasm2](https://docs.google.com/document/d/1GRmy3rA4DiYtBlX-I1Jr_iHykbX8EixC3Mq0TCYqbKc/edit#heading=h.q4c21ihutzk0) WebAssembly 程序文件

[Back To Top](#top)


### 模块

- [design/go-dep](https://docs.google.com/document/d/18tNd8r5DV0yluCR7tPvkMTsWD_lYcRO7NhpNSDymRr8) Go 打包的建议流程
- [design/go-dep2](https://docs.google.com/document/d/1qnmjwfMmvSCDaY4jxPmLAccaaUI5FfySNE90gB0pTKQ/edit) 依赖管理工具
   - [doc/go-dep](https://golang.design/history/6https://blog.gopheracademy.com/advent-2016/saga-go-dependency-management/) Sam Boyer，Go 依赖管理事迹。2016年12月13日。
   - [talk/go-dep](https://www.youtube.com/watch?v=5LtMb090AZI) GopherCon 2017：Sam Boyer，Go 依赖管理的新时代。2017年7月24日。
   - [talk/go-dep-design](https://www.youtube.com/watch?v=wBTGd_dvnO8) dotGo 2017：sam boyer， 依赖的功能设计。2017年12月8日。
   - [discuss/go-dep](https://www.youtube.com/watch?v=sbrZfPgNmfw) Russ Cox，Jess Frazelle, Sam Boyer, Pete Garcin，构建时添加可预测性。2018年2月22日。
- [design/vgo](https://go.dev/design/24301-versioned-go) Russ Cox，提案：Go 模块版本化。2018年3月20日。
   - [issue/24301](https://go.dev/issue/24301) cmd/go: 像 Go 工具链添加包版本支持。
   - [doc/deps](https://research.swtch.com/deps) Russ Cox，我们的软件依赖问题。2019年1月23日。
   - [doc/vgo](https://research.swtch.com/vgo) Russ Cox，Go 和版本化。
   - [discuss/groups-gomod](https://groups.google.com/g/golang-dev/c/a5PqQuBljF4) Go 模块化来了。
   - [discuss/go-dep-response](https://www.reddit.com/r/golang/comments/92f3q1/peter_bourgon_a_response_about_dep_and_vgo/) Reddit 讨论。
   - [doc/go-dep-response](https://peter.bourgon.org/blog/2018/07/27/a-response-about-dep-and-vgo.html) Peter Bourgon，关于 dep 和 vgo 的回应。2018年7月27日。
   - [discuss/go-dep-response2](https://news.ycombinator.com/item?id=17628311) 黑客新闻：关于 dep 和 vgo 的回应。
   - [discuss/go-dep-twitter](https://twitter.com/_rsc/status/1022588240501661696) Russ Cox 的推特风暴。
- [design/sumdb](https://go.dev/design/25530-sumdb) Russ Cox, Filippo Valsorda，提案：公共 Go 模块生态的安全性。2019年4月24日。
   - [issue/25530](https://go.dev/issue/25530) 提案：cmd/go: 透明日志的安全版本。
- [issue/23966](https://go.dev/issue/23966#issuecomment-377997161) 为什么 go.mod 要定制语法？
- [design/lazy-gomod](https://go.dev/design/36460-lazy-module-loading) Bryan C. Mills，提案：模块的延迟加载。2020年2月20日。
- [issue/45713](https://go.dev/issue/45713) 提案：cmd/go: 添加工作区模式
   - [design/workspace](https://go.dev/design/45713-workspace) 提案：cmd/go 多模块工作区

[Back To Top](#top)


### gopls

- [design/gopls-workspace](https://go.dev/design/37720-gopls-workspaces) Heschi Kreinick, Rebecca Stambler，提案：多项目 gopls 工作区，2020年3月9日。

[Back To Top](#top)


### 测试，x/perf

- [design/subtests](https://go.dev/design/12166-subtests) Marcel van Lohuizen，测试：计划中的子测试和子基准测试支持。2015年9月2日。
   - [issue/12166](https://go.dev/issue/12166) 测试：计划中的子测试和子基准测试支持。
- [design/gotest-bench](https://go.dev/design/14313-benchmark-format) Russ Cox, Austin Clements，提案：Go 基准测试数据格式。2016年2月。
   - [issue/14313](https://go.dev/issue/14313) cmd/go: 决定，标准基准数据格式的文档。
- [issue/20875](https://go.dev/issue/20875) 测试：考虑在基准测试时减少调用 ReadmemStats
- [issue/27217](https://go.dev/issue/27217) 测试：带有 StopTimer 的小型基准测试永久运行。
- [issue/41637](https://go.dev/issue/41637) 测试：基准测试迭代报告错误。
- [issue/41641](https://go.dev/issue/41641) 测出：中断计时器时基准测试不一致。
- [design/gotest-json](https://go.dev/design/2981-go-test-json) Nodir Turakulov，提案：go test 的 -json 标记。
- [design/testing-helper](https://go.dev/design/4899-testing-helper) Caleb Spare，提案：测试：用 TB.Helper 更好地支持测试帮助程序。2016年12月27日。
   - [issue/4899](https://go.dev/issue/4899) 测试：添加 t.Helper 让 file:line 的结果更好用。
- [design/fuzzing](https://go.dev/s/draft-fuzzing-design) Katie Hockman，设计草案：模糊测试的起步。
- [issue/43744](https://go.dev/issue/43744) 测试：基准测试单元属性
- [issue/48803](https://go.dev/issue/48803) 所有：Go 编译/运行时性能监控系统。
- [issue/49121](https://go.dev/issue/49121) x/perf/storage：支持 postgresql。

[Back To Top](#top)


## 运行时核心
### 调度器

- [paper/work-steal](https://dl.acm.org/citation.cfm?id=324234) Robert D. Blumofe 和 Charles E. Leiserson，通过工作窃取调度多线程计算。J. ACM 46, 5 (September 1999), 720-748.
- [cl/sched-m-1](https://github.com/golang/go/commit/96824000ed89d13665f6f24ddc10b3bf812e7f47#diff-1fe527a413d9f1c2e5e22e08e605a192) Russ Cox，清理调度器，2008年8月5日。
- [cl/sched-m-n](https://github.com/golang/go/commit/fe1e49241c04c748d0e3f4762925241adcb8d7da) 现在情况好多了。2009年11月11日。
- [design/go11sched](https://go.dev/s/go11sched) Dmitry Vyukov，可伸缩的 Go 调度器设计文档。2012年
   - [cl/7314062](https://github.com/golang/go/commit/779c45a50700bda0f6ec98429720802e6c1624e8) 运行时：提升的调度器。
- [design/sched-preempt-dmitry](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#heading=h.3pilqarbrc9h) Dmitry Vyukov。Go 优先权调度器设计文档。2013年
- [design/sched-numa](https://docs.google.com/document/u/0/d/1d3iI2QWURgDIsSR6G2275vMeQ_X7w-qxM2Vp7iGwwuM/pub) Dmitry Vyukov，Go 的 NUMA 感知调度器。2014年9月。
- [design/go15gomaxprocs](https://go.dev/s/go15gomaxprocs) Russ Cox，Go 1.5 的 GOMAXPROCS 默认值。2015年5月。
- [doc/go17sched](https://www.quora.com/How-does-the-golang-scheduler-work/answer/Ian-Lance-Taylor) Ian Lance Taylor，Go 调度器怎样工作？2016年7月16日。
- [design/sched-preempt-austin](https://go.dev/design/24543-non-cooperative-preemption) Austin Clements，提案：goroutine 抢占竞争。
   - [cl/43050](https://go.dev/cl/43050) cmd/compile：在 amd64 上使用错误分支的循环抢占。
   - [issue/10958](https://go.dev/issue/10958) 运行时：紧密循环应该是可抢占的。
   - [issue/24543](https://go.dev/issue/24543) 运行时：goroutine 抢占竞争。
   - [issue/36365](https://go.dev/issue/36365) 运行时：清理松散结尾上的异步抢占
- [issue/14592](https://go.dev/issue/14592) 运行时：让空闲的系统线程退出。
- [issue/18237](https://go.dev/issue/18237) 运行时：当 goroutine 被经常唤醒时调度很慢。
- [issue/20395](https://go.dev/issue/20395) 运行时：当所属的 goroutine 退出时中断锁住的系统线程。
- [issue/20458](https://go.dev/issue/20458) 提案：运行时：一起调用 LockOSThread、UnlockOSThread
- [issue/21827](https://go.dev/issue/21827) 运行时：使用 runtime.LockOSThread 会造成很大的性能影响。
- [issue/27345](https://go.dev/issue/27345) 运行时：将父 goroutine 的堆栈用于新的 goroutine。
- [issue/28808](https://go.dev/issue/28808) 运行时：当 GOMAXPROCS 很高时调度器工作窃取很慢。
- [issue/32113](https://go.dev/issue/32113) 运行时：减少P流失的优化。
- [issue/44313](https://go.dev/issue/44313) 运行时：已停止的Ms不能成为专用或小数的GC工作者。


[Back To Top](#top)

### 执行栈

- [design/contigstack](https://go.dev/s/contigstacks) 连续堆栈
- [issue/17007](https://go.dev/issue/17007) 运行时：致命错误：就绪状态糟糕的 g->status
- [issue/18138](https://go.dev/issue/18138) 运行时：新的 goroutine 在更多堆栈上花费更多的时间
   - [design/predict-stack-size](https://docs.google.com/document/d/1YDlGIdVTPnmUiTAavlZxBI1d9pwGQgZT7IKFKlIXohQ/edit#) Keith Randall，确定一个好的堆栈起始大小，2021年8月18日。
   - [cl/341990](https://go.dev/cl/341990) 运行时：预测堆栈大小
   - [cl/345889](https://go.dev/cl/345889) 运行时：测量堆栈使用情况；需要的话就启动更大的堆栈
- [issue/26061](https://go.dev/issue/26061) 运行时：g0 stack.lo 有时太慢


[Back To Top](#top)

### 内存分配器
Go 内存分配的实力简述：Russ Cox 在 Go 1 上首次实现了基于 tcmallor 的内存分配器，mcache 缓存在 M 上。然后他修改了分配器，允许使用 16GB 的内存，后来到了128GB。然而，分配器（包括收集器）有很严重的锁竞争，且不可伸缩。在 Dmitry 的可伸缩运行时调度器后，分配器可以直接从P锁竞争中分配（更少）。同时，收集器从独立线程迁移到系统监视器线程。现在，Michael 正致力于提高内存分配器的可伸缩性，例如将收集器迁移到用户线程，基于位图的页分配器，可伸缩的 mcentral。

- [doc/tcmalloc](http://goog-perftools.sourceforge.net/doc/tcmalloc.html) Sanjay Ghemawat, Paul Menage。TCMalloc : 线程缓存 Malloc，谷歌，2009。
- [design/go113scavenge](https://go.googlesource.com/proposal/+/aa701aae530695d32916b779e048a3e18311a2e3/design/30333-smarter-scavenging.md) Michael Knyszek，提案：智能收集。2019年5月9日。
   - [issue/30333](https://go.dev/issue/30333) 运行时：智能收集
   - [issue/32012](https://go.dev/issue/32012) 运行时：后台收集对小堆过于频繁
   - [issue/31966](https://go.dev/issue/31966) 运行时：后台收集堆栈会显著延迟死锁检测
   - [issue/34047](https://go.dev/issue/34047) 运行时：由 scavenge.lock 引起的潜在死循环
   - [issue/34048](https://go.dev/issue/34048) 运行时：收集器步长无法解释碎片
   - [issue/35788](https://go.dev/issue/35788) 运行时：收集器的效率低于之前
   - [issue/36521](https://go.dev/issue/36521) 运行时：Go 1.12 性能下降
   - [issue/36603](https://go.dev/issue/36603) 运行时：系统经常在不收集的内存上调用
- [design/go114pagealloc](https://go.googlesource.com/proposal/+/a078ea9d72b99dc88fdfd2cb6ee150a8ce202ea2/design/35112-scaling-the-page-allocator.md) Michael Knyszek, Austin Clements，提案：缩放Go 页分配器。2019年10月18日。
   - [issue/35112](https://go.dev/issue/35112) 运行时：页分配器缩放
   - [cl/200439](https://go.dev/cl/200439) 运行时：对触发比例设置下限
- [issue/8885](https://go.dev/issue/8885) 运行时：考虑添加24字节大小的类
- [issue/37487](https://go.dev/issue/37487) 运行时：提高 mcentral 的可扩展性
   - [cl/221182](https://go.dev/cl/221182) 运行时：添加新的 mcentral 实现
- [issue/18155](https://go.dev/issue/18155) 运行时：当没有垃圾时延迟扫描辅助
- [issue/19112](https://go.dev/issue/19112) 运行时：gcControllerState.enlistWorker 死锁
- [issue/23687](https://go.dev/issue/23687) 运行时：在 Linux 上也使用 MADV_FREE
   - [cl/135395](https://go.dev/cl/135395) 运行时：在 Linux 上使用 MADV_FREE（如果可用）
- [issue/29707](https://go.dev/issue/29707) cmd/trace: 无法解析轨迹：无法对事件进行一致的排序
- [issue/35954](https://go.dev/issue/35954) 运行时：句柄更优雅地分配到分配器地址空间的顶部
- [issue/37927](https://go.dev/issue/37927) 运行时：GC 步长在低 GOGC 表现出奇怪行为。
- [issue/38130](https://go.dev/issue/38130) 运行时：页分配器的健全性检查不正确。
- [issue/38404](https://go.dev/issue/38404) 运行时：在 arm64/mips64le 上，STW GC 不工作。
- [issue/38605](https://go.dev/issue/38605) 运行时：pageAlloc.allocToCache 不正确地更新 pageAlloc.searchAddr
- [issue/38617](https://go.dev/issue/38617) 运行时：Go 1.14 中收集器在 Windows 里冻结，因为时间粒度太粗略。
- [issue/38966](https://go.dev/issue/38966) 运行时：aix-ppc64 构建器出现致命错误：“修剪错误”，“地址范围的基准和限制不在同一个内存片段中”
- [issue/39128](https://go.dev/issue/39128) 运行时：从 CL 233497 开始，linux-mips-rtrk 构建器一直因为“修剪错误”而失败。
- [issue/40191](https://go.dev/issue/40191) 运行时：pageAlloc.searchAddr 可能指向不连续堆中的未映射的内存，违反了其不变性
- [issue/40457](https://go.dev/issue/40457) 运行时：runqputbatch 不保护它对 globrunqputbatch 的调用。
- [issue/40641](https://go.dev/issue/40641) 运行时：堆栈收缩和通道收发之前的竞争导致错误的 sudog 值
- [issue/42330](https://go.dev/issue/42330) 运行时：在 Linux 上默认为 MADV_DONTNEED
   - [cl/267100](https://go.dev/cl/267100) 运行时：在 Linux 上默认为 MADV_DONTNEED


[Back To Top](#top)

### 垃圾收集器

- [paper/on-the-fly-gc](https://doi.org/10.1145/359642.359655) Edsger W. Dijkstra, Leslie Lamport, A. J. Martin, C. S. Scholten 和 E. F. M. Steffens。动态垃圾收集：合作练习。Commun.ACM 21，11（1978年11月），966-975。
- [paper/yuasa-barrier](https://doi.org/10.1016/0164-1212(90)90084-Y) T. Yuasa。通用机器上的实时垃圾收集。J. Syst. Softw. 11,3 （1990年3月），181-198.
- [design/go13gc](https://docs.google.com/document/d/1v4Oqa0WwHunqlb8C3ObL_uNQw3DfSY-ztoA-4wWbKcg/pub) Dmitry Vyukov，更简单和更快速的 Go GC。2014年7月16日。
   - [cl/106260045](https://codereview.appspot.com/106260045) 运行时：更简单和更快速的 Go GC
- [design/go14gc](https://go.dev/s/go14gc) Richard L. Hudson。Go 1.4+ 的垃圾收集计划和路线图。2014年8月6日。
- [design/go15gcpacing](https://go.dev/s/go15gcpacing) Austin Clements。Go 1.5 并发垃圾收集步长。2015年3月10日。
- [discuss/gcpacing](https://groups.google.com/forum/#!topic/golang-dev/YjoG9yJktg4) Austin Clements 等人。关于“提案：垃圾收集器步长”的讨论。2015年3月10日。
- [issue/11970](https://go.dev/issue/11970) 运行时：将 GC 协调器替换为状态机。
- [design/sweep-free-alloc](https://go.dev/design/12800-sweep-free-alloc) Austin Clements。提案：密集标记位和无扫描分配。2015年9月30日。
- [issue/12800](https://go.dev/issue/12800) 运行时：将空间的列表替换为直接的位图分配。
- [design/decentralized-gc](https://go.dev/design/11970-decentralized-gc) Austin Clements。提案：分散的GC协调器。2015年10月25日。
- [issue/12967](https://go.dev/issue/12967#issuecomment-171466238) 运行时：标记终止时的收缩堆栈显著增加了 GC 的 STW 时间
- [issue/14951](https://go.dev/issue/14951) 运行时：突变的辅助过于激进，尤其是在高 GC 时。
- [design/eliminate-rescan](https://go.dev/design/17503-eliminate-rescan) Austin Clements, Rick Hudson。消除 STW 堆栈的再次扫描。2016年10月21日。
   - [issue/17503](https://go.dev/issue/17503) 运行时：消除堆栈再扫描。
- [design/concurrent-rescan](https://go.dev/design/17505-concurrent-rescan) Austin Clements, Rick Hudson。提案：堆栈的并发再扫描。2016年10月18日。
   - [issue/17505](https://go.dev/issue/17505) 运行时：堆栈并发再扫描。
- [design/soft-heap-limit](https://go.dev/design/14951-soft-heap-limit) Austin Clements。提案：分开软硬堆大小目标。2018年7月12日。
- [issue/22460](https://go.dev/issue/22460) 运行时：优化写入屏障。
- [design/roc](https://go.dev/s/gctoc) 面向请求的收集器算法（ROC）
   - [cl/roc](https://go.dev/cl/25058) 运行时：用 ROC 编写屏障代码
   - [cl/generational-gc](https://go.dev/cl/137482) 运行时：触发器生成GC
- [doc/ismm-gc](https://go.dev/blog/ismmkeynote) Rick Hudson，Go 垃圾收集器之旅。
- [discuss/ismm-gc](https://groups.google.com/forum/#!topic/golang-dev/UuDv7W1Hsns) 垃圾回收 Slides 和 Transcritp 已可用。
- [design/simplify-mark-termination](https://go.dev/design/26903-simplify-mark-termination) Austin Clements，提案：简略标记终止和消除标记2。2018年8月9日。
   - [issue/26903](https://go.dev/issue/26903) 运行时：简略标记终止和消除标记2
- [design/gcscan](https://docs.google.com/document/d/1un-Jn47yByHL7I0aVIP_uVCMxjdM5mpelJhiKlIqxkE/edit#) 提案：堆栈的GC扫描
   - [issue/22350](https://go.dev/issue/22350) cmd/compile: 编译器莫名其妙地保留内存。
- [issue/27993](https://go.dev/issue/27993) 运行时：错误消息：在标记终止时 P 会缓存 GC 内容。
- [issue/37116](https://go.dev/issue/37116) 运行时：在 Go1.14rc1 中 GC 会延迟 10ms-26ms，可能因为 GC 空闲。
- [issue/42430](https://go.dev/issue/42430) 运行时：GC 里 pacer 问题讨论
   - [issue/39983](https://go.dev/issue/39983) 运行时：空闲的GC妨碍自动缩放
   - [issue/17969](https://go.dev/issue/17969) 运行时：侵略型的GC实现会干扰共同租户。
   - [issue/14812](https://go.dev/issue/14812) 运行时：GC 导致延迟峰值。
   - [issue/40460](https://go.dev/issue/40460) 运行时：goroutine 可能分配超出硬堆目标。
   - [issue/29696](https://go.dev/issue/29696) 提案：运行时：为应用增加回应GC反压力的方式。
   - [issue/23044](https://go.dev/issue/23044) 提案：运行时：添加一种机制来指定最小目标堆大小
- [issue/42642](https://go.dev/issue/42642) 运行时：多毫秒扫描终止暂停（第二版）
- [issue/44163](https://go.dev/issue/44163) 运行时：删除空闲的GC工作者。
- [issue/44167](https://go.dev/issue/44167) 运行时：重新设计 GC Pacer
   - [design/gc-pacer-redesign](https://go.dev/design/44167-gc-pacer-redesign) 重新设计 GC Pacer
- [issue/44309](https://go.dev/issue/44309) 提案：runtime/debug: 用户可配置目标内存
- [issue/48409](https://go.dev/issue/48409) 提案：runtime/debug: 软内存限制
   - [design/user-configurable-memory-target](https://go.dev/design/44309-user-configurable-memory-target) 用户可配置目标内存
   - [design/soft-memory-limit](https://go.dev/design/48409-soft-memory-limit) 软内存限制
- [issue/45894](https://go.dev/issue/45894) 运行时：标记终止重启突变太慢。
- [issue/45315](https://go.dev/issue/45315) 运行时：runtime.GC 无需完成扫描就能返回。


[Back To Top](#top)

### 统计

- [issue/16843](https://go.dev/issue/16843) 运行时：用于监控堆大小的机制
   - [cl/setmaxheap](https://go-review.googlesource.com/c/go/+/46751/) Austin Clements，runtime/debug：添加 SetMaxHeap 应用接口。2017年6月26日。
- [design/go116runtime-metric](https://github.com/golang/proposal/blob/44d4d942c03cd8642cef3eb2f6c153f2e9883a77/design/37112-unstable-runtime-metrics.md) Michael Knyszek，提案：用于不稳定运行时的指标应用接口。2020年3月18日。
- [issue/19812](https://go.dev/issue/19812) 运行时：在GC期间无法 ReadMemStats。
- [issue/38712](https://go.dev/issue/38712) 运行时：TestMemStats 很奇怪
- [issue/40459](https://go.dev/issue/40459) 运行时：在循环中调用 ReadMemStats 可能阻止GC


[Back To Top](#top)

### 内存模型
Go内存模型包含以下方面：

- 关于原子操作的内存顺序
- 关于 sync 包应用接口的内存顺序
- 关于运行时机制的内存顺序（即对象终结器）
- [doc/refmem](https://go.dev/ref/mem) Rob Pike 和 Russ Cox。Go 内存模型。2009年2月21日。
- [issue/4947](https://go.dev/issue/4947) cmd/cc: 原子的实质
- [issue/5045](https://go.dev/issue/5045) 文档：定义 sync/atomic 和内存模型怎样交互
- [issue/7948](https://go.dev/issue/7948) 文档：定义 sync 和内存模型怎样交互
- [issue/9442](https://go.dev/issue/9442) 文档：定义 finalizers 和内存模型怎样交互
- [issue/33815](https://go.dev/issue/33815) doc/go_mem: "hello, world" 并不总是打印两次
- [cl/75130045](https://codereview.appspot.com/75130045) 文档：允许缓冲通道在未初始化时作为信号量
- [doc/gomem](http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf) Russ Cox。Go 内存模型。2016年2月25日。
- [doc/go2017russ](https://research.swtch.com/go2017#memory) Russ Cox，我的Go决议2017：内存模型。2017年1月18日。
- [doc/atomic-bug](https://go.dev/pkg/sync/atomic/#pkg-note-BUG) atomic 包
- [discuss/atomic-mem-order](https://groups.google.com/d/msg/golang-dev/vVkH_9fl1D8/azJa10lkAwAJ) 通过 atomic 的 Load/Store 来提供内存顺序保证。
- [issue/37355](https://go.dev/issue/37355) runtime/race: 使用 -race 运行但没有 race（与内存模型不匹配）
   - [cl/220419](https://go.dev/cl/220419) 运行时：交换 raceacquire() 和 racerelease 的顺序
   - [issue/42598](https://go.dev/issue/42598) 运行时：在 CL 220419 后缓冲通道的明显错误的竞争报告。
   - [cl/271987](https://go.dev/cl/271987) 运行时：在调用竞争检测器前检查通道的 elemsize
   - [paper/fava2020fix] Fava, Daniel Schnetzer，“查找并修复Go内存模型和数据竞争检测器间的不匹配。”软件工程和格式化方法的国际化会议。2020年 Springer, Cham。
- [doc/mm](https://research.swtch.com/mm) Russ Cox，内存模型，2021年6月
   - [doc/hwmm](https://research.swtch.com/hwmm) Russ Cox，硬件内存模型。2021年6月29日。
   - [doc/plmm](https://research.swtch.com/plmm) Russ Cox，编程语言内存模型。2021年7月6日。
   - [doc/gomm](https://research.swtch.com/gomm) Russ Cox，更新Go内存模型。2021年7月12日。
   - [discuss/47141](https://github.com/golang/go/discussions/47141) 更新Go内存模型。


[Back To Top](#top)

### ABI

- [design/cgo-pointers](https://go.dev/design/12416-cgo-pointers) Ian Lance Taylor，提案：在Go和C之间传递指针的规则。2015年10月。
   - [issue/12416](https://go.dev/issue/12416) cmd/cgo: 指定Go和C之间传递指针的规则。
- [design/internal-abi](https://go.dev/design/27539-internal-abi) Austin Clements. 创建未定义的内部调用规定。2019年1月14日。
   - [issue/27539](https://go.dev/issue/27539) 提案：运行时：使 ABI 未定义，作为更改的步骤。
   - [issue/31193](https://go.dev/issue/31193) cmd/compile，运行时：减少函数开头的开销。
- [design/register-call](https://go.dev/design/40724-register-calling) Austin Clements, 另有 Cherry Zhang, Michael Knyszek, Martin Möhrmann, Michael Pratt, David Chase, Keith Randall, Dan Scales, and Ian Lance Taylor 提供建议。提案：基于注册的Go调用协议。2020年8月10日。
- [issue/18597](https://go.dev/issue/18597) 提案: cmd/compile: 定义基于注册的调用协议。
- [issue/40724](https://go.dev/issue/40724) 提案：Go 函数切换到基于注册的调用协议。
   - [cl/266638](https://go.dev/cl/266638) 反射、运行时：为 ASM 例程使用内部 ABI，第二次尝试。
   - [cl/259445](https://go.dev/cl/259445) cmd/compile,cmd/link: 为ABI包装器的初始支持
- [design/internal-abi-spec](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md) Go 内部ABI规范。

[Back To Top](#top)


### Misc

- [issue/20135](https://go.dev/issue/20135) 运行时：删除元素时map收缩。


[Back To Top](#top)

## 标准库
### syscall

- [design/go14syscall](https://go.dev/s/go1.4-syscall) syscal 包。


[Back To Top](#top)

### os, io, io/fs, embed
Go 1.16中，对新的 os/fs 包进行了大量返工和改进。
[design/draft-iofs](https://go.dev/design/draft-iofs) Russ Cox, Rob Pike，Go文件系统接口——设计草案。2020年7月。

   - [issue/13473](https://go.dev/issue/13473) 提案：os: Stdin, Stdout 和 Stderr 应该是接口
   - [issue/14106](https://go.dev/issue/14106) 提案：os：File 应该是接口
   - [issue/19660](https://go.dev/issue/19660) 提案: 将 io/ioutil 重命名为 io/fileio 或类似名字
   - [issue/40025](https://go.dev/issue/40025) 提案: io/ioutil: 移动 Discard, NopCloser, ReadAll 到 io
   - [issue/42027](https://go.dev/issue/40027) 提案: path/filepath: 添加 WalkDir (使用 DirEntry 浏览)
   - [issue/41190](https://go.dev/issue/41190) io/fs: 添加文件系统接口
   - [issue/41467](https://go.dev/issue/41467) os: 为轻量级目录查询添加 ReadDir 方法。
   - [issue/41974](https://go.dev/issue/41974) 提案: io/fs, filepath: 添加高效的浏览替代方案。
   - [issue/42026](https://go.dev/issue/42026) 提案: os: 添加 ReadDir, ReadFile, WriteFile, CreateTemp, MkdirTemp，以及反对 io/ioutil
   - [issue/43223](https://go.dev/issue/43223) 提案: io/fs, net/http: 定义自动 ETag 服务的接口
- [design/go116embed](https://go.dev/s/draft-embed-design) Russ 和 Braid，嵌入文件。
   - [issue/41191](https://go.dev/issue/41191) embed, cmd/go: 添加嵌入文件的支持
   - [issue/42321](https://go.dev/issue/42321) embed: 在 embed.FS 文档中应该加入警告，说明包含点文件
   - [issue/42328](https://go.dev/issue/42328) 提案: cmd/go: 在使用 `//go:embed`时避免包含隐藏文件
   - [issue/43216](https://go.dev/issue/43216) 提案: embed: 删除在本地变量上对嵌入指令的支持
   - [issue/43217](https://go.dev/issue/43217) 提案: embed: 定义 String 和 Bytes 类型别名时必须使用 `//go:embed`
   - [issue/43218](https://go.dev/issue/43218) embed: 对 string, []byte 讨论的决定
   - [issue/44166](https://go.dev/issue/44166) io/fs,os: 使用 os.DirFS 的 fs.ReadDir 会造成错误的路径
   - [issue/42322](https://go.dev/issue/42322) io/fs: 添加函数 `Sub(fsys FS, dir string) FS`


[Back To Top](#top)

### go/*

- [doc/gotypes](https://go.dev/s/types-tutorial) Alan Donovan. go/types: Go 类型检查器
- [talk/gotypes](https://docs.google.com/presentation/d/13OvHYozAUBeISPRoLgG7kMBuja1NsU1D_mMlmbaYojk/view) Alan Donovan. 使用 go/types 作为代码理解和重构工具。2015年12月2日。
- [design/modular-interface](https://docs.google.com/document/d/1-azPLXaLgTCKeKDNg0HVMq2ovMlD-e7n1ZHzZVzOlJk/edit) Alan Donovan. 提案：用于Go模块化静态分析的通用接口。2018年9月9日。
   - [cl/134935](https://go.dev/cl/134935) go/analysis: 分析工具的新的应用接口


[Back To Top](#top)

### sync

- [design/percpu-sharded](https://go.dev/design/18802-percpu-sharded) Sanjay Menakuru. 提案：percpu.Sharded，一个用来减少缓存竞争的应用接口。2018年6月12日。
   - [issue/18802](https://go.dev/issue/18802) 提案：sync: 支持分片值
- [issue/37142](https://go.dev/issue/37142) sync: sync 包里的收缩类型


[Back To Top](#top)

#### Map

- [issue/21031](https://go.dev/issue/21031) sync: 减少 map 里的指针开销
- [issue/21032](https://go.dev/issue/21032) sync: 用存储新键值来减少 (*map).Load 的损失
- [issue/21035](https://go.dev/issue/21035) sync: 通过新的不相交的键来减少Map操作间的竞争
- [issue/37033](https://go.dev/issue/37033) 运行时：提供集中的工具来管理 go/cgo 的指针句柄


[Back To Top](#top)

#### Pool

- [discuss/add-pool](https://groups.google.com/d/msg/golang-dev/kJ_R6vYVYHU/LjoGriFTYxMJ) gc-aware 池耗尽策略
- [issue/4720](https://go.dev/issue/4720) sync：添加池类型，2013年1月28日。
- [cl/46010043](https://github.com/golang/go/commit/f8e0057bb71cded5bb2d0b09c6292b13c59b5748#diff-2e9fc106a7387ca4c32ecf856a91f82a) sync: 缩放池，2014年1月24日。
- [cl/86020043](https://github.com/golang/go/commit/8fc6ed4c8901d13fe1a5aa176b0ba808e2855af5#diff-2e9fc106a7387ca4c32ecf856a91f82a) sync: 池里减少激进的本地缓存，2014年4月14日。
- [cl/162980043](https://github.com/golang/go/commit/af3868f1879c7f8bef1a4ac43cfe1ab1304ad6a4#diff-491b0013c82345bf6cfa937bd78b690d) sync: 在第二次和之后的GC期间释放池内存，2014年10月22日。
- [issue/8979](https://go.dev/issue/8979) sync: 在GC期间池没有释放内存。
- [issue/22331](https://go.dev/issue/22331) 运行时：池清理导致 STW1 时间过长。
- [issue/22950](https://go.dev/issue/22950) sync: 避免每次GC都清理全部池。
   - [cl/166960](https://github.com/golang/go/commit/d5fd2dd6a17a816b7dfd99d4df70a85f1bf0de31) sync: 使用无锁结构进行池窃取。
   - [cl/166961](https://github.com/golang/go/commit/2dcbf8b3691e72d1b04e9376488cef3b6f93b286) sync: 使用受害者缓存在GC上进行平缓的池操作。
- [issue/24479](https://go.dev/issue/24479) sync: 在池操作里去除全部 Mutex。


[Back To Top](#top)

#### Mutex, RWMutex

- [cl/4631059](https://go.dev/cl/4631059) 运行时：替换 Semacquire/Semrelease 实现。
- [issue/9201](https://go.dev/issue/9201) 提案：同步：禁止解锁不同 goroutine 的 mutex。
- [issue/13086](https://go.dev/issue/13086) 运行时：在重复睡眠获取失败后回退到公平锁。
   - [cl/34310](https://go.dev/cl/34310) sync: 让 Mutex 更公平。
- [issue/17973](https://go.dev/issue/17973) sync: RWMutex 无法随 CPU 数量扩展。
   - [cl/215361](https://go.dev/cl/215361) sync: 实现能够避免缓存竞争的 RWMutex 版本。


[Back To Top](#top)

#### Groups

- [cl/134395](https://go.dev/cl/134395) errgroup: 重新考虑并发模式。
   - [cl/131815](https://go.dev/cl/131815) errgroup: 句柄运行时。从子 goroutine 进行 Goexit。
   - [issue/15758](https://go.dev/issue/15758) testing: 同时使用 T.FatalX 和 T.SkipX 时应该提示错误。
   - [issue/25448](https://go.dev/issue/25448) 提案：将`panic(nil)`改为非空的panic值。


[Back To Top](#top)

#### atomic

- [issue/8739](https://go.dev/issue/8739) runtime,sync/atomic: 统一 runtime/internal/atomic 和 sync/atomic 的应用接口。
- [issue/20164](https://go.dev/issue/20164) 提案: atomic: 添加 `(*Value).Swap`
- [discuss/atomic-value](https://groups.google.com/g/golang-dev/c/SBmIen68ys0/m/WGfYQQSO4nAJ)
- [issue/36606](https://go.dev/issue/36606) 提案: cmd/compile: 让64位字段在32位系统上对齐为64位，在结构上添加 `//go:packed`指令
- [issue/37262](https://go.dev/issue/37262) 运行时：在32位架构上不能原子性地访问 tiny-allocated 结构的第一个字。


[Back To Top](#top)

### time

- [design/monotonic-time](https://go.dev/design/12914-monotonic) Russ Cox，提案：Go 中单调经过的时间跨度。2017年1月26日。
   - [issue/12914](https://go.dev/issue/12914) time: 使用单调时钟测量经过的时间
- 可扩展的计时器
   - [cl/34784](https://go.dev/cl/34784) 运行时：在多 CPU 系统上的计时器可扩展性。
   - [issue/6239](https://go.dev/issue/6239) 运行时：让计时器更快。
   - [issue/15133](https://go.dev/issue/15133) 运行时：计时器无法在有大量计时器的多CPU系统上扩展。
   - [issue/27707](https://go.dev/issue/27707) time: 	当使用 Ticker 和 Sleep 时CPU使用率过高。
- 后续的延迟讨论
   - [issue/18023](https://go.dev/issue/18023) 运行时：使用 runtime.LockOSThread 意外会大幅减慢。
   - [issue/25471](https://go.dev/issue/25471) time: 需要大约7次 syscall 才会睡眠。
   - [issue/38070](https://go.dev/issue/38070) 运行时：因为抢占点导致计时器自己死锁。
   - [issue/36298](https://go.dev/issue/36298) net: mac 上的 1.14 版本性能减弱。
   - [issue/38860](https://go.dev/issue/38860) 运行时：CPU 限制 goroutine 导致不必要的计时器延迟。
      - [cl/216198](https://go.dev/cl/216198) 运行时：将 poller 返回的 goroutine 添加进本地运行队列。
      - [cl/232199](https://go.dev/cl/232199) 运行时：从运行中的 P's 中窃取计时器。
      - [cl/232298](https://go.dev/cl/232298) 运行时：减少计时器延迟。
   - [issue/44343](https://go.dev/issue/44343) 运行时：time.Sleep 比预计时间更长
   - [issue/44868](https://go.dev/issue/44868) time, runtime: 零间隔的计时器需要2分钟才能启动。


[Back To Top](#top)

### context

- [issue/8082](https://go.dev/issue/8082) 提案：规范：通过接口的定义而不是包和名字来表示接口。
- [issue/14660](https://go.dev/issue/14660) 提案：context: 标准库的新包。
- [issue/16209](https://go.dev/issue/16209) 提案：放宽名称不同但实际相同的接口的分配规则。
- [issue/20280](https://go.dev/issue/20280) 提案: io: 向 Reader 等添加 Context 参数。
- [issue/21355](https://go.dev/issue/21355) 提案：让 Context 使用本地 goroutine 存储
- [issue/24050](https://go.dev/issue/24050) testing: 在1.10上使用子进程测试有时会挂住。
- [issue/27982](https://go.dev/issue/27982) 提案: Go 2: 加入协作式 goroutine 取消机制
- [issue/28342](https://go.dev/issue/28342) 提案: Go 2: 更新 context 包
- [issue/29011](https://go.dev/issue/29011) 提案: Go 2: 使用结构化的并发
- [doc/context-go-away](https://faiface.github.io/post/context-should-go-away-go2/) Michal Štrba，Go 2 应该取消 Context，2017年8月6日。
- [doc/context](https://go.dev/blog/context) Go 并发模式: Context。
- [doc/context-isnt-for-cancellation](https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation) Dave Cheney，Context 不适用取消。
- [issue/42564](https://go.dev/issue/42564) context: cancelCtx 独占锁导致极端的竞争。


[Back To Top](#top)

### encoding

- [design/go12encoding](https://go.dev/s/go12encoding) Russ Cox，Go 1.2里的 encoding，TextMarshaler 和 TextUnmarshaler，2013年7月。
- [design/go12xml](https://docs.google.com/document/d/1G-AMeUPoyTnbZDkuCJA89DjJwTjdWFctIZ_N9NgA9RM/pub) Russ Cox，Go 1.2 里的 xml.Marshaler 和 Unmarshaler，2013年7月。
- [design/natural-xml](https://go.dev/design/13504-natural-xml) Sam Whited，提案: 原生 XML. 2016年9月27日。
   - [issue/13504](https://go.dev/issue/13504) encoding/xml: 添加 XML 数据的通用表现形式。
- [design/zip](https://go.dev/design/14386-zip-package-archives) Russ Cox，提案：基于 Zip 的 Go 包存档，2016年2月。
   - [issue/14386](https://go.dev/issue/14386) 提案：在 .a 和 .o 文件中使用 zip 格式。
- [design/xmlstream](https://go.dev/design/19480-xml-stream) Sam Whited，建议：XML 流。2017年3月9日。
   - [issue/19480](https://go.dev/issue/19480) encoding/xml: 从 TokenReader 添加解码，启用流转换器。
- [design/raw-xml](https://go.dev/design/26756-rawxml-token) Sam Whited，提案：原始 XML 令牌，2018年9月1日。
   - [issue/26756](https://go.dev/issue/26756) 提案: encoding/xml: 添加 RawXML 令牌。


[Back To Top](#top)

### image, x/image
以下的讨论关于颜色管理和图片标准库。现在，Go图片库不会在编码和解码期间读取或写入元信息。因此，在处理图片时（例如在非线性 sRGB 空间中缩放），颜色可能出错。通用解决方案是设计图片元数据API，来识别编码图片文件中的颜色配置文件。

- [issue/11420](https://go.dev/issue/11420) x/image/draw: 颜色空间校正插值
   - [issue/20613](https://go.dev/issue/20613) image/png: 不要忽略 PNG gAMA块
   - [issue/27830](https://go.dev/issue/27830) 提案: image: 解码选项
   - [cl/253497](https://go.dev/cl/253497) x/image/draw: 伽玛校正非线性插值
   - [issue/37188](https://go.dev/issue/37188) image/color: 文档没有包含指向颜色理论资源的链接
- [issue/33457](https://go.dev/issue/33457) 提案: image: 添加对 jpeg、gif、png 的通用元数据支持
   - [discuss/image-metadata](https://groups.google.com/g/golang-dev/c/aRvnYIcaIaA/m/9GVKL7mIAgAJ) image.Metadata 接口类型
   - [issue/18365](https://go.dev/issue/18365) image/png: 不支持设计和检索 PPI/DPI
   - [cl/208559](https://go.dev/cl/208559) image: 新的元数据感知读写API
   - [cl/216799](https://go.dev/cl/216799) image: 元数据API草图
- [issue/44808](https://go.dev/issue/44808) image, image/draw: 添加直接使用 RGBA64 的接口
- [issue/46395](https://go.dev/issue/46395) image/draw: 如果掩码是 *image.Alpha，通过使用特殊情况来提升性能。


[Back To Top](#top)

### Mobile

- [design/go14android](https://go.dev/s/go14android) David Crawshaw，Go 支持安卓，2014年6月。
- [design/gobind](https://go.dev/s/gobind) David Crawshaw，绑定Go和Java，2014年7月。


[Back To Top](#top)

### misc

- [design/mobile-audio](https://go.dev/design/13432-mobile-audio) Jaana Burcu Dogan，提案：手机音频。2015年11月30日。
   - [issue/13432](https://go.dev/issue/13432) 提案: x/mobile 音频
- [design/localization](https://go.dev/design/12750-localization) Marcel van Lohuizen，提案：Go 本地化支持。2016年1月28日。
   - [issue/12750](https://go.dev/issue/12750) x/text: 本地化支持
- [design/unbounded-queue](https://go.dev/design/27935-unbounded-queue-package) Christian Petrin，提案：内置高性能无界队列。2018年10月2日。
   - [issue/27935](https://go.dev/issue/27935) 提案: 添加 container/queue
- [design/lockfile](https://go.dev/design/33974-add-public-lockedfile-pkg) Adrien Delorme，公开内部的锁定文件包
   - [issue/33974](https://go.dev/issue/33974)
- [design/cidr](https://go.dev/design/16704-cidr-notation-no-proxy) Rudi Kramer, James Forrest，在 no_proxy 变量中添加对CIDR表示法的支持。2017年7月10日。
   - [issue/16704](https://go.dev/issue/16704) net/http: 考虑在 no_proxy 环境变量添加 CIDR 表示法
- [design/dns](https://go.dev/design/26160-dns-based-vanity-imports) Sam Whited，提案：基于DNS的虚导入。2018年6月30日。
   - [issue/26160](https://go.dev/issue/26160) 提案：将 DNS 文本记录用于虚导入路径


[Back To Top](#top)

## 未分类但相关的链接

- [使用Go Guru：用于导航Go代码的集成编辑器工具，Alan Donovan。](https://go.dev/s/using-guru)
- [Go Guru 计划 (née Oracle) ，Alan Donovan。](https://docs.google.com/document/d/1UErU12vR7jTedYvKHVNRzGPmXqdMASZ6PfE7B-p6sIg/edit)
- [为什么 X 在标准库中？（已放弃、未完成）](https://go.dev/s/stdwhy)
- [go oracle: 设计，Alan Donovan。](https://go.dev/s/oracle-design)
- [go oracle: 用户手册，Alan Donovan。](https://go.dev/s/oracle-user-manual)
- [cgihttpproxy](https://go.dev/s/cgihttpproxy)
- [sqldrivers](https://go.dev/s/sqldrivers)
- [Go 2 设计](https://go.dev/s/go2designs)
- [C编程语言笔记，Rob Pike。](http://doc.cat-v.org/bell_labs/pikestyle)
- [”我得到的最好的编程建议“，Rob Pike。](https://www.informit.com/articles/article.aspx?p=1941206)
- [Mihai Budiu 对 Brian Kernighan 的采访](http://www.cs.cmu.edu/~mihaib/kernighan-interview/index.html)
- [语言设计，Brian Kernighan Holiday Repeat By SE Daily](https://softwareengineeringdaily.com/2017/12/28/language-design-with-brian-kernighan-holiday-repeat/)
- [C 编程语言: Brian Kernighan - Computerphile](https://youtu.be/de2Hsvxaf8M)
- [调试Go程序中的性能问题，Dmitriy Vyukov。](https://software.intel.com/content/www/us/en/develop/blogs/debugging-performance-issues-in-go-programs.html)
- [我想实现的有趣的论文（或者有部分实现）](https://github.com/dgryski/interesting-papers)
- [golang/wiki/ResearchPapers](https://github.com/golang/go/wiki/ResearchPapers)


[Back To Top](#top)

## 有趣的事实

- [cl/1](https://github.com/golang/go/commit/7d7c6a97f815e9279d08cfaea7d5efb5e90695a8) Brian Kernighan，Go的首次提交。1972年7月19日。
- [issue/9](https://go.dev/issue/9) 我使用了 MY 作为编程语言的名字，2009年11月11日。
- [issue/2870](https://go.dev/issue/2870) 9 不是素数，2012年2月3日。
- [doc/gophercount](https://research.swtch.com/gophercount) 有多少Go开发者？2019年11月1日。
- [discuss/google-owns-go](https://groups.google.com/forum/#!msg/golang-nuts/6dKNSN0M_kg/Y1yDJRwQBgAJ) Russ Cox 关于“Go是谷歌的语言，不是我们的语言”的回应。


[Back To Top](#top)

## 鸣谢
首先感谢 [TalkGo](https://github.com/talkgo) 社区创建者 [Mai Yang](https://github.com/yangwenmai) 对 [golang.design](https://golang.design/) 倡议的大力推动。他创建的 TalkGo 极大地改变了中国的 Go 社区。他也是一个伟大的人，积极为各种与 Go 相关的项目作贡献。


同样重要的是感谢 TalkGo 社区核心成员 [qcrao](https://github.com/qcrao) 和 [eddycjy](https://github.com/eddycjy) 的持续而鼓舞人心的讨论和分享。


如果没有他们的支持，此文档无法完成。

[Back To Top](#top)

## 凭证
[golang.design/history](https://github.com/golang-design/history) | CC-BY-NC-ND 4.0 © [changkun](https://changkun.de/)