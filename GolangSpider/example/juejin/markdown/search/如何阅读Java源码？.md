# 如何阅读Java源码？ #

阅读本文大概需要 3.6 分钟。

## 阅读Java源码的前提条件： ##

# #

### 1、技术基础 ###

在阅读源码之前，我们要有一定程度的技术基础的支持。

假如你从来都没有学过Java，也没有其它编程语言的基础，上来就啃《Core Java》，那样是很难有收获的，尤其是《深入Java虚拟机》这类书，或许别人觉得好，但是未必适合现在的你。

比如设计模式，许多Java源码当中都会涉及到。再比如阅读Spring源码的时候，势必要先对IOC，AOP，Java动态代理等知识点有所了解。

# #

### 2、强烈的求知欲 ###

强烈的求知欲是阅读源码的核心动力！

大多数程序员的学习态度分为如下几个层次：

* 

完成自己的项目就可以了，遇到不懂的地方就百度一下。

* 

不仅做好项目，还会去阅读一些和项目有关的书籍。

* 

除了阅读和项目相关的书籍之外，还会阅读一些IT行业相关的书籍。

* 

平时会经常逛逛GitHub，找一些开源项目看看。

* 

阅读基础框架、J2EE规范、源码。

大多数程序员的层次都是在第一层，到第五层的人就需要有强烈的求知欲了。

### 3、足够的耐心 ###

通过阅读源码我们可以学习大佬的设计思路，技巧。还可以把我们一些零碎的知识点整合起来，从而融会贯通。总之阅读源码的好处多多，想必大家也清楚。

但是真的把那么庞大复杂的代码放到你的眼前时，肯定会在阅读的过程中卡住，就如同陷入了一个巨大的迷宫，如果想要在这个巨大的迷宫中找到一条出路，那就需要把整个迷宫的整体结构弄清楚，比如：API结构、框架的设计图。而且还有理解它的核心思想，确实很不容易。

刚开始阅读源码的时候肯定会很痛苦，所以，没有足够的耐心是万万不行的。

## 如何读Java源码： ##

团长也是经历过阅读源码种种痛苦的人，算是有一些成功的经验吧，今天来给大家分享一下。

如果你已经有了一年左右的Java开发经验的话，那么你就有阅读Java源码的技术基础了。

### 1、建议从JDK源码开始读起，这个直接和eclipse集成，不需要任何配置。 ###

可以从JDK的工具包开始，也就是我们学的《数据结构和算法》Java版，如List接口和ArrayList、LinkedList实现，HashMap和TreeMap等。这些数据结构里也涉及到排序等算法，一举两得。

面试时，考官总喜欢问ArrayList和Vector的区别，你花10分钟读读源码，估计一辈子都忘不了。

然后是core包，也就是String、StringBuffer等。 如果你有一定的Java IO基础，那么不妨读读FileReader等类。

建议大家看看《Java In A Nutshell》，里面有整个Java IO的架构图。Java IO类库，如果不理解其各接口和继承关系，则阅读始终是一头雾水。

Java IO 包，我认为是对继承和接口运用得最优雅的案例。如果你将来做架构师，你一定会经常和它打交道，如项目中部署和配置相关的核心类开发。

读这些源码时，只需要读懂一些核心类即可，如和ArrayList类似的二三十个类，对于每一个类，也不一定要每个方法都读懂。像String有些方法已经到虚拟机层了(native方法)，如hashCode方法。

当然，如果有兴趣，可以对照看看JRockit的源码，同一套API，两种实现，很有意思的。

如果你再想钻的话，不妨看看针对虚拟机的那套代码，如System ClassLoader的原理，它不在JDK包里，JDK是基于它的。JDK的源码Zip包只有10来M，它像是有50来M，Sun公司有下载的，不过很隐秘。我曾经为自己找到、读过它很兴奋了一阵。

### 2、Java Web项目源码阅读 ###

步骤：表结构 → web.xml → mvc → db → spring ioc → log→ 代码

① 先了解项目数据库的表结构，这个方面是最容易忘记的，有时候我们只顾着看每一个方法是怎么进行的，却没有去了解数据库之间的主外键关联。其实如果先了解数据库表结构，再去看一个方法的实现会更加容易。

② 然后需要过一遍web.xml，知道项目中用到了什么拦截器，监听器，过滤器，拥有哪些配置文件。如果是拦截器，一般负责过滤请求，进行AOP等；如果是监听器，可能是定时任务，初始化任务；配置文件有如 使用了spring后的读取mvc相关，db相关，service相关，aop相关的文件。

③ 查看拦截器，监听器代码，知道拦截了什么请求，这个类完成了怎样的工作。有的人就是因为缺少了这一步，自己写了一个action，配置文件也没有写错，但是却怎么调试也无法进入这个action，直到别人告诉他，请求被拦截了。

④ 接下来，看配置文件，首先一定是mvc相关的，如springmvc中，要请求哪些请求是静态资源，使用了哪些view策略，controller注解放在哪个包下等。然后是db相关配置文件，看使用了什么数据库，使用了什么orm框架，是否开启了二级缓存，使用哪种产品作为二级缓存，事务管理的处理，需要扫描的实体类放在什么位置。最后是spring核心的ioc功能相关的配置文件，知道接口与具体类的注入大致是怎样的。当然还有一些如apectj等的配置文件，也是在这个步骤中完成。

⑤ log相关文件，日志的各个级别是如何处理的，在哪些地方使用了log记录日志。

⑥ 从上面几点后知道了整个开源项目的整体框架，阅读每个方法就不再那么难了。

⑦ 当然如果有项目配套的开发文档也是要阅读的。

### 3、Java框架源码阅读 ###

当然了，就是Spring、MyBatis这类框架。

在读Spring源码前，一定要先看看《J2EE Design and Development》这本书，它是Spring的设计思路。注意，不是中文版，中文版完全被糟蹋了。

想要阅读MyBatis的源码就要先了解它的一些概念，否则云里来雾里去的什么也不懂。有很多人会选择去买一些书籍来帮助阅读，当然这是可取的。那么如果不想的话，就可以去官网查看它的介绍（MyBatis网站：http://www.mybatis.org/mybatis-3/zh/getting-started.html），团长也是按照官网上面的介绍来进行源码阅读的。团长认为MyBatis的亮点就是管理SQL语句。

**总结**

没有人一开始就可以看得懂那些源码，我们都是从0开始的，而且没有什么捷径可寻，无非就是看我们谁愿意花时间去研究，谁的求知欲更强烈，谁更有耐心。阅读源码的过程中我们的能力肯定会提升，可以从中学到很多东西。在我们做项目的时候就会体现出来了，的确会比以前顺手很多。

**·END·**

**程序员的成长之路**

路虽远，行则必至

本文原发于 同名微信公众号「程序员的成长之路」，回复「1024」你懂得，给个赞呗。

微信ID：cxydczzl

**往期精彩回顾**

[程序员接私活的7大平台利器]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247485759%26amp%3Bidx%3D1%26amp%3Bsn%3D6189c69534c0f51b0a3a6d1b26f689ff%26amp%3Bchksm%3Dfa24f657cd537f41df3b688e63cfc88f6e68cb9262ec07e1418089aa6cdc5d794aeb681e4786%26amp%3Bscene%3D21%23wechat_redirect )

[Java程序员的成长之路]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247486053%26amp%3Bidx%3D1%26amp%3Bsn%3Dd3ab081c851786f4c9dbf1b323f187ea%26amp%3Bchksm%3Dfa24f50dcd537c1b34253c8e5ad027dfb317c0e7001b0ef48e336c27a17841717a197bb0b0c2%26amp%3Bscene%3D21%23wechat_redirect )

[白话TCP为什么需要进行三次握手]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247486048%26amp%3Bidx%3D1%26amp%3Bsn%3Dfe9d125ee560875239deeafc801a04c8%26amp%3Bchksm%3Dfa24f508cd537c1e576e88b1a179fd870a517c40b6526e9ca2cffb02d6991f75439c579e5b2c%26amp%3Bscene%3D21%23wechat_redirect )

[Java性能优化的50个细节（珍藏版）]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247486015%26amp%3Bidx%3D1%26amp%3Bsn%3Dc865a942d05aa13e19208ec9deb1f8ef%26amp%3Bchksm%3Dfa24f557cd537c418e605a112a08a4125105a80d36ac2b2b29e2ef6f456fb8b7ee5c91e11986%26amp%3Bscene%3D21%23wechat_redirect )

[设计电商平台优惠券系统]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247485996%26amp%3Bidx%3D1%26amp%3Bsn%3D70a1c8de8ee03d91130011494997c6a6%26amp%3Bchksm%3Dfa24f544cd537c524e83b4218067241cfc1eb8a5bf9b90fb64d3f13f699d57492fdbc96d8b7f%26amp%3Bscene%3D21%23wechat_redirect )

[一个对话让你明白架构师是做什么的？]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247485982%26amp%3Bidx%3D1%26amp%3Bsn%3D25a51a789918462326bc3d3d0c39a110%26amp%3Bchksm%3Dfa24f576cd537c6038ba1a6f903d38f9aa4c7695e7bee136b649fd2d22d4371e58e1db2fe83b%26amp%3Bscene%3D21%23wechat_redirect )

[教你一招用 IDE 编程提升效率的骚操作！]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247485965%26amp%3Bidx%3D1%26amp%3Bsn%3Db09bd68122a7a610cc6bcbad0b765ed2%26amp%3Bchksm%3Dfa24f565cd537c7393b174170c39fb2d7ba810b4cac1404aefb2360bb6552133e88a32016e44%26amp%3Bscene%3D21%23wechat_redirect )

[送给程序员们的经典电子书大礼包]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUyNDkzNzczNQ%3D%3D%26amp%3Bmid%3D2247485875%26amp%3Bidx%3D1%26amp%3Bsn%3D0801f98619e13eec9db30f218cb4d9af%26amp%3Bchksm%3Dfa24f6dbcd537fcd2cf3c36698702c4aecbb2ce18f935a7dfdbf1ab39cc1514f53b546be6ca7%26amp%3Bscene%3D21%23wechat_redirect )