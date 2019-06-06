# Sun 公司的 Java 跟 Android 使用的 Java 库有什么关系？ #

**全文重点罗列**

* Java 是 Sun 公司开发的一门语言，同时 Java 也是一个开放平台。
* Sun 公司为 JVM 发布了 JVM 规范，任何公司都可以按照此规范开发 JVM 语言，如现在的 Kotlin、Scala 等。
* JVM 语言必须要通过 JCP 组织(Java Community Process)对其拥有的 TCK(Technology Compatibility Kit)测试。
* Harmony 是 Apache 2005 年开发的一个开源、免费的 Java 实现，只是 **没有完全通过 Sun 公司的 TCK 测试** ，按照 **Apache 协议** 发布。
* Open JDK 是 Sun 公司在 2009 发布的完全自由，开放源码的 Java平台标准版（Java SE）免费开源实现，按照 **GPL 协议** 发布。
* Android 一开始使用的 Java 实现是基于 Apache 协议发布的 Harmony，后来由于 Harmony 本身的限制以及 Oracle 公司的起诉，从 Android N 开始, Google 开始用 Open JDK来替换 Harmony。

如题， **Sun 公司的 Java 跟 Android 使用的 Java 库有什么关系？** 做 Android 开发不少时间了，但是相信有很多人有同样的疑问，尽管这个疑问对做 Android 开发本身并没有多大关系，但我总觉得这里面一定有一些有意思的东西。

最开始有这个疑问是源自 Oracle 与 Google 2011 年左右诉讼。当时 Oracle 状告 Google 在未经授权就使用了 Java 的 API，那时对知识产权、开源协议等缺少了解，只是觉得奇怪，Java 是一门开放的技术，任何公司都可以拿来使用，为什么到 Google 这里怎么就变卦了呢，接下来就记录一下自己的研究笔记，以问答形式展开。

### Sun 公司与 Oracle 公司的关系 ###

> 
> 
> 
> [Sun Microsystems](
> https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2FSun%2520Microsystems%2F6064586
> ) 是IT及互联网技术服务公司（已被 [Oracle(甲骨文)](
> https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2F%25E7%2594%25B2%25E9%25AA%25A8%25E6%2596%2587%2F471435
> ) 收购）Sun Microsystems 创建于1982年。主要产品是 [工作站](
> https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2F%25E5%25B7%25A5%25E4%25BD%259C%25E7%25AB%2599%2F217955
> ) 及 [服务器](
> https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2F%25E6%259C%258D%25E5%258A%25A1%25E5%2599%25A8%2F100571
> ) 。1986年在美国成功上市。1992年 Sun 推出了市场上第一台多处理器台式机 SPARCstation 10
> system，并于1993年进入财富500强。
> 
> 

2009 年 Sun 公司被 Oracle 收购。

### 什么是 Java？ [wiki]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FJava ) ###

Sun 公司的 [詹姆斯·高斯林]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2F%25E8%25A9%25B9%25E5%25A7%2586%25E6%2596%25AF%25C2%25B7%25E9%25AB%2598%25E6%2596%25AF%25E6%259E%2597 ) 等人在 1990 年年初开发了 Java 语言的雏形（好吧，那年我刚出生，没想到自己跟 Java 同龄了），最初命名为 Oak，主要适用于家用电器等小型系统的开发。后来随着互联网的快速发展，于是改造 Oak，在 1995 年 5 月以 Java 的名称正式发布，所以更多时候，大家都在讲 Java 是 1995-96才有的语言，也没问题。

Java 从最初设计就有一个非常优秀的思想，即跨平台，它不受平台限制，任何平台都可以运行 Java 程序，比如我们常见的 Windows、Linux 等等系统都可以运行，Java 的跨平台具体是通过 JVM 做到的。

### JVM 是什么东西？ ###

首先， Java 语言不论在任何平台，想要运行，首先都需要先编译为字节码，这里的字节码是平台无关的，是完全独立的一个东西，不受平台特性限制。而最终在电脑上运行时，为了在各个平台上可以执行字节码，Sun 公司在开发 Java 的过程中同时也开发了虚拟机这个东西， **它专门用来执行字节码，它负责把字节码转换为机器可以认识的机器码** 然后运行。

于此同时，Sun 不仅仅是为各个平台开发了虚拟机。作为程序员，我们大都有这样的经验，当我们要为多个平台开发一个作用一致的东西时，为了后续更好的扩展维护，我们肯定会想先定义一个接口规范出来，这样所有平台的实现都通过这个规范去控制约束，后续的升级维护就会变得更简单。

作为开发 Java 的老一辈们当然也知道这个道理，所以他们也做了同样的事， **Sun 定义了 Java 虚拟机的规范** ，这样相当于发布了 Java 规范，即你只要按照这个规范，就可以开发自己的 JVM 语言了，前提是你的语言编译出的字节码要符合 JVM 规范，否则虚拟机就没法执行。现在这样的语言已经有不少了，比如我们知道的 Kotlin，Groovy、Scala等。

### JVM 语言是不是只要符合规范即可？ ###

原则上任何团体或者公司都可以开发 JVM 语言，因为这个规范是公开的，但是这里要知道的是，当你按照规范开发了语言，首先要做的一件事便是要 **通过 Sun 的 TCK (Technology Compatibility Kit)测试** 。这个也很好理解，这是官方的测试，通过这个测试才能成为官方认可的 JVM 语言。具体 TCK 的详细介绍可查看 [Wiki。]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FTechnology_Compatibility_Kit )

### Sun 公司如何通过 Java 赚钱 ###

Sun通过销售Java Enterprise System等 **专用产品的许可证** 从Java中 **获得收入** 。

### Apache 开发的 Harmony 是一个什么项目 ###

该项目是 Apache 开发的一个开源的，免费的 Java 实现，不过 2011 年时已经被 Apache 关闭了该项目。Harmony 项目实现了 J2SE 5.0的 99％完整性和 Java SE 6的 97％。

Harmony 的目标：

* 一个兼容的、独立的Java SE 5 JDK的实现，并根据 **Apache License v2** 发布。
* 一个由社区开发的模块化的运行时（包括 [Java虚拟机]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FJava%25E8%2599%259A%25E6%258B%259F%25E6%259C%25BA ) 和 [类库]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2F%25E7%25B1%25BB%25E5%25BA%2593 ) ）体系结构。

### Harmony 与 TCK 的纷争 ###

> 
> 
> 
> 如果需要成为一个带有Java logo标志的，可以声称自己兼容Sun公司实现的JDK，需要通过 [JCP](
> https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FJCP
> ) （Java Community Process）对其拥有的 [TCK](
> https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fw%2Findex.php%3Ftitle%3DTCK%26amp%3Baction%3Dedit%26amp%3Bredlink%3D1
> ) （Technology Compatibility Kit）的测试。Apache Harmony项目一直在努力争取获得 [JCP](
> https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FJCP
> ) 的授权。
> 
> 
> 
> 但是，由于Sun公司的态度， **JCP并没有给Harmony授予TCK许可** ，而且 SUN 发布 OpenJDK 之后，还规定只有衍生自
> OpenJDK 的采用GPL协议的开源实现才能运行OpenJDK的TCK[ [2]](
> https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FApache_Harmony%23cite_note-2
> ) ，
> 
> 
> 
> **但 Apache 的 Harmony 是 Apache 协议的，与 OpenJDK 的 GPLv2 协议不兼容，Apache
> 董事会和Harmony项目工作人员坚决反对这种带有条件的授权，认为这种是在开源社区里不可接受的** 。因此，两者谈判破裂。直到现在，Harmony一直没有获得TCK的授权。有批评称，Sun无视它签署的JCP法律协定，这摧毁了全部的信任。
> 
> 
> 

> 
> 
> 
> 摘自 [Apache Harmony - 维基百科，自由的百科全书](
> https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FApache_Harmony
> )
> 
> 

### GPL 协议与 Apache 的区别 ###

下面是常见开源协议之间的关系区别

![](https://user-gold-cdn.xitu.io/2019/5/8/16a946cd7a70892e?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到

* 用 GPL 协议发布的项目，当别人修改代码后必须采用相同的协议并且也要开源。
* Apache 协议则特别自由，当你修改该协议下发布的项目后，可以选择闭源，只需要在每一个文件头都加上版权说明。

### Harmony 与 Android 的关系 ###

从上面的开源协议对比可以看出，Apache 协议是一个特别自由的协议，而 Android 一开始正是因为 Harmony 项目采用了 Apache 协议，所以才更有动力使用 Harmony 作为自己的 Java 类库。

所以 Android 中的 Java 实现并不是官方的实现，而是 Harmony 对 Java 的实现。由于 Harmony 并没有通过 TCK 认证，这也就为日后 Oracle 起诉谷歌侵权埋下了伏笔。

到后来 Android 鉴于此，从 Android 7.0 开始，用 Sun 的 OpenJDK 替换了 Harmony。

### Open JDK 是什么 ###

OpenJDK（Open Java Development Kit）原是 Sun 公司为 Java 平台构建的 Java 开发环境（JDK）的开源版本，完全自由，开放源码。 Sun公司在 2006 年的 JavaOne 大会上称将对 Java 开放源代码， **于2009年4月15日正式发布OpenJDK** 。甲骨文在 2010 年收购 Sun之后接管了这个专案。

OpenJDK 是 Java平台标准版（Java SE）的免费开源实现，该实现是根据 **GNU通用公共许可证（GNU GPL）发布** 。OpenJDK 最初只基于 Java 平台的 JDK 7版本。

### Apache Harmony 与 Open JDK、原生 Java 以及 Android 之间的区别对比。 ###

+-----------------+--------------------------------+------------------------+------------------------+
|     ANDROID     |              JAVA              |     APACHE HARMONY     |        OPENJDK         |
+-----------------+--------------------------------+------------------------+------------------------+
| Android Runtime | Java Runtime Environment (JRE) | 同左                   | 同左                   |
| Core Libraries  | Java Core Libraries            | Harmony Core Libraries | OpenJDK Core Libraries |
| Dalvik VM       | Java VM                        | DRLVM                  | HotSpot VM             |
+-----------------+--------------------------------+------------------------+------------------------+

由图可看到

* Java 与 Android 有不同的运行时，但是 Open JDK 跟 Harmony 都有相同的运行时，这就保证了这两者编写出的 Java 代码均能正常跑在 Java 平台上。
* 而 Android 有自己的运行时。
* 大家都定义了自己的 VM，Android 定义的最彻底。
* Java 的 VM 基于堆栈结构，而 Android 的 VM 基于寄存器结构，后者效率更高，更适合移动端。

## 总结 ##

一开始只是想了解 Harmony ，但是真的展开了，其实有很多东西需要去了解，包括语言本身、语言特性、开源协议，项目历史，一些关键事件等等。而花时间了解后，却发现不知道的东西更多了，疑问也更多了，比如 Android 具体选用 Harmony 是如何思考的，Harmony 与 Open JDK 的实现过程中有什么异同等等，也是无奈，这里只能把自己看到的想到的记录下来。

无论如何，去了解一些东西背后的东西或者宏观背景，能给自己更多的视角去看待眼前的东西，也会让自己变得更理性，同时举一反三，任何事情都应该如此。

原文地址： [Sun 公司的 Java 跟 Android 使用的 Java 库有什么关系？ | 咕咚]( https://link.juejin.im?target=https%3A%2F%2Fgudong.name%2F2019%2F04%2F05%2Fandroid-why-java-harmony.html )

## 参考链接 ##

* [Apache Harmony - 维基百科，自由的百科全书]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2FApache_Harmony )