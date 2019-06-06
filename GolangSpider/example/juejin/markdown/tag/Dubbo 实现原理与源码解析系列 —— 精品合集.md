# Dubbo 实现原理与源码解析系列 —— 精品合集 #

摘要: 原创出处 http://www.iocoder.cn/Dubbo/good-collection/ 「芋道源码」欢迎转载，保留摘要，谢谢！

* [1.【芋艿】精尽 Dubbo 原理与源码专栏]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F )
* [2.【老徐】RPC 专栏]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F )
* [3.【肥朝】Dubbo 源码解析]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F )
* [4.【MR_QI】Dubbo 源码解析]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F )
* [5.【杨武兵】Dubbo 源码解析]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F )
* [666. 彩蛋]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F )

![](https://user-gold-cdn.xitu.io/2018/6/17/1640a9d749914316?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 🙂🙂🙂关注**微信公众号：【芋道源码】**有福利：
> 
> * RocketMQ / MyCAT / Sharding-JDBC **所有** 源码分析文章列表
> * RocketMQ / MyCAT / Sharding-JDBC **中文注释源码 GitHub 地址**
> * 您对于源码的疑问每条留言 **都** 将得到 **认真** 回复。 **甚至不知道如何读源码也可以请教噢** 。
> * **新的** 源码解析文章 **实时** 收到通知。 **每周更新一篇左右** 。
> * **认真的** 源码交流微信群。
> 

# 1. 【芋艿】精尽 Dubbo 原理与源码专栏 #

* 作者：芋艿
* **只更新在笔者的知识星球，欢迎加入一起讨论 Dubbo 源码与实现** 。 ![](https://user-gold-cdn.xitu.io/2018/6/23/1642a041b825721e?imageView2/0/w/1280/h/960/ignore-error/1)

* 目前已经有 **600+** 位球友加入...
* 进度：已经完成 **69+** 篇，预计总共 75+ 篇，完成度 **92%** 。

* 对应 Dubbo 版本号： **2.6.X**

## 1.1 目录 ##

* 调试环境搭建

* [《精尽 Dubbo 源码分析 —— 调试环境搭建》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 项目结构一览

* [《精尽 Dubbo 源码分析 —— 项目结构一览》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 配置 Configuration

* [《精尽 Dubbo 源码解析 —— API 配置（一）之应用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— API 配置（二）之服务提供者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— API 配置（三）之服务消费者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 属性配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精近 Dubbo 源码解析 —— XML 配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [ ] [《精尽 Dubbo 源码解析 —— 注解配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [ ] [《精尽 Dubbo 源码解析 —— 外部化配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [ ] [《精尽 Dubbo 源码解析 —— 自动装配》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 核心流程一览

* [《精近 Dubbo 源码解析 —— 核心流程一览》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 拓展机制 SPI

* [《精尽 Dubbo 源码分析 —— 拓展机制 SPI》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 线程池

* [《精尽 Dubbo 源码分析 —— 线程池》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 服务暴露 Export

* [《精尽 Dubbo 源码分析 —— 服务暴露（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务暴露（二）之远程暴露（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 服务引用 Refer

* [《精尽 Dubbo 源码分析 —— 服务引用（一）之本地引用（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（二）之远程引用（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 注册中心 Registry

* [《精尽 Dubbo 源码分析 —— Zookeeper 客户端》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（二）之 Zookeeper》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（三）之 Redis》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（四）之 Simple 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（五）之 Multicast 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 动态编译 Compile

* [《精尽 Dubbo 源码分析 —— 动态编译（一）之 Javassist》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 动态编译（二）之 JDK》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 动态代理 Proxy

* [《精尽 Dubbo 源码分析 —— 动态代理（一）之 Javassist》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 动态代理（二）之 JDK》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 动态代理（三）之本地存根 Stub》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 服务调用 Invoke

* [《精尽 Dubbo 源码分析 —— 服务调用（一）之本地调用（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【1】通信实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【2】同步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【3】异步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【4】参数回调》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（三）之远程调用（HTTP）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（四）之远程调用（Hessian）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（五）之远程调用（Webservice）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（六）之远程调用（REST）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（七）之远程调用（RMI）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（八）之远程调用（Redis）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（九）之远程调用（Memcached）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（十）之远程调用（Thrift）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 调用特性

* [《精尽 Dubbo 源码解析 —— 调用特性（一）之回音测试》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 调用特性（二）之泛化引用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 调用特性（三）之泛化实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 过滤器 Filter

* [《精尽 Dubbo 源码解析 —— 过滤器（一）之 ClassLoaderFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 过滤器（二）之 ContextFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 过滤器（三）之 AccessLogFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（四）之 ActiveLimitFilter && ExecuteLimitFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（五）之 TimeoutFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（六）之 DeprecatedFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（七）之 ExceptionFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（八）之 TokenFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（九）之 TpsLimitFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（十）之 CacheFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（十一）之 ValidationFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* NIO 服务器

* [《精尽 Dubbo 源码分析 —— NIO 服务器（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（二）之 Transport 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（三）之 Telnet 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（四）之 Exchange 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（五）之 Buffer 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（六）之 Netty4 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（七）之 Netty3 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（八）之 Mina1 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（九）之 Grizzly 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* P2P 服务器

* [《精尽 Dubbo 源码分析 —— P2P 服务器》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* HTTP 服务器

* [《精尽 Dubbo 源码分析 —— HTTP 服务器》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 序列化 Serialization

* [《精尽 Dubbo 源码分析 —— 序列化（一）之 总体实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 序列化（二）之 Dubbo 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 序列化（三）之 Kryo 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 服务容器 Container

* [《精尽 Dubbo 源码解析 —— 服务容器》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 集群容错 Cluster

* [《精尽 Dubbo 源码解析 —— 集群容错（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（二）之 Cluster 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（三）之 Directory 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（四）之 Loadbalance 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（五）之 Merger 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（六）之 Configurator 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（七）之 Router 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（八）之 Mock 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 优雅停机

* [《精尽 Dubbo 源码解析 —— 优雅停机》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 日志适配

* [《精尽 Dubbo 源码解析 —— 日志适配》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 状态检查

* [《精尽 Dubbo 源码解析 —— 状态检查》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

* 监控中心 Monitor

* TODO 等待新版本

* 管理中心 Admin

* TODO 等待新版本

* 运维命令 QOS

* TODO 比较简单，未来写一下

* 链路追踪 Tracing

* [《SkyWalking 源码分析 —— Agent 插件（二）之 Dubbo》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FSkyWalking%2Fagent-plugin-dubbo%2F%3Fself )

> 
> 
> 
> ps：
> 
> 

> 
> 
> 
> * 小圆点代表已完成；
> * 小方块代表未来计划完成
> * 钟华贤代表不考虑写
> 
> 
> 

## 1.2 Dubbo 用户指南 ##

本小节，我们将 [《精尽 Dubbo 源码解析》]( # ) 和 [《Dubbo 用户指南》]( https://link.juejin.im?target=https%3A%2F%2Fwww.gitbook.com%2Fbook%2Fdubbo%2Fdubbo-user-book ) 做一次映射，方便大家直接找到感兴趣的功能的具体源码实现。当然，如果有整理不到位的地方，欢迎大家给我留言斧正。

### 【<2> 快速启动】 ###

> 
> 
> 
> Dubbo 采用全 Spring 配置方式，透明化接入应用，对应用没有任何 API 侵入，只需用 Spring 加载 Dubbo
> 的配置即可，Dubbo 基于 Spring 的 Schema 扩展进行加载。
> 
> 
> 
> 如果不想使用 Spring 配置，可以通过 [API 的方式](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fconfiguration%2Fapi.md
> ) 进行调用。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 调试环境搭建》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.1> XML 配置】 ###

> 
> 
> 
> 有关 XML 的详细配置项，请参见： [配置参考手册](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fxml%2Fintroduction.html
> ) 。如果不想使用 Spring 配置，而希望通过 API 的方式进行调用，请参见： [API配置](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fconfiguration%2Fapi.html
> ) 。想知道如何使用配置，请参见： [快速启动](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fquick-start.html
> ) 。
> 
> 

对应源码解析文章：

* [《精近 Dubbo 源码解析 —— XML 配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.2> 属性配置】 ###

> 
> 
> 
> 如果公共配置很简单，没有多注册中心，多协议等情况，或者想多个 Spring 容器想共享配置，可以使用 dubbo.properties
> 作为缺省配置。
> 
> 
> 
> Dubbo 将自动加载 classpath 根目录下的 dubbo.properties，可以通过JVM启动参数 `
> -Ddubbo.properties.file=xxx.properties` 改变缺省配置位置。 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fconfiguration%2Fproperties.html%23fn_1
> )
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 属性配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.3> API 配置】 ###

> 
> 
> 
> API 属性与配置项一对一，各属性含义，请参见： [配置参考手册](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fxml%2Fintroduction.html
> ) ，比如： ` ApplicationConfig.setName("xxx")` 对应 ` <dubbo:application
> name="xxx" />` [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fconfiguration%2Fapi.html%23fn_1
> )
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— API 配置（一）之应用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— API 配置（二）之服务提供者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— API 配置（三）之服务消费者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.4> 注解配置】 ###

> 
> 
> 
> 需要 ` 2.5.7` 及以上版本支持
> 
> 

对应源码解析文章：

* [ ] [《精尽 Dubbo 源码解析 —— 注解配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.1> 启动时检查】 ###

> 
> 
> 
> Dubbo 缺省会在启动时检查依赖的服务是否可用，不可用时会抛出异常，阻止 Spring 初始化完成，以便上线时，能及早发现问题，默认 `
> check="true"` 。
> 
> 
> 
> 可以通过 ` check="false"` 关闭检查，比如，测试时，有些服务不关心，或者出现了循环依赖，必须有一方先启动。
> 
> 
> 
> 另外，如果你的 Spring 容器是懒加载的，或者通过 API 编程延迟引用服务，请关闭 check，否则服务临时不可用时，会抛出异常，拿到
> null 引用，如果 ` check="false"` ，总是会返回引用，当服务恢复时，能自动连上。
> 
> 

* **注意** ，服务提供者注册失败，是没有启动检查功能的。这一点，笔者一直理解错了。

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 注册中心（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.2> 集群容错】 ###

> 
> 
> 
> 在集群调用失败时，Dubbo 提供了多种容错方案，缺省为 failover 重试。
> 
> 
> 
> 
> 
> ![cluster](https://user-gold-cdn.xitu.io/2018/6/23/1642a0435681db67?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（二）之 Cluster 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（三）之 Directory 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.3> 负载均衡】 ###

> 
> 
> 
> 在集群负载均衡时，Dubbo 提供了多种均衡策略，缺省为 ` random` 随机调用。
> 
> 
> 
> 可以自行扩展负载均衡策略，参见： [负载均衡扩展](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-dev-book%2Fcontent%2Fimpls%2Fload-balance.html
> )
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（四）之 Loadbalance 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.4> 线程模型】 ###

> 
> 
> 
> 如果事件处理的逻辑能迅速完成，并且不会发起新的 IO 请求，比如只是在内存中记个标识，则直接在 IO 线程上处理更快，因为减少了线程池调度。
> 
> 
> 
> 但如果事件处理逻辑较慢，或者需要发起新的 IO 请求，比如需要查询数据库，则必须派发到线程池，否则 IO 线程阻塞，将导致不能接收其它请求。
> 
> 
> 
> 如果用 IO 线程处理事件，又在事件处理过程中发起新的 IO 请求，比如在连接事件中发起登录请求，会报“可能引发死锁”异常，但不会真死锁。
> 
> 
> 
> 
> 
> ![dubbo-protocol](https://user-gold-cdn.xitu.io/2018/6/23/1642a0436fe65f56?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 线程池》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.5> 直连提供者】 ###

> 
> 
> 
> 在开发及测试环境下，经常需要绕过注册中心，只测试指定服务提供者，这时候可能需要点对点直连，点对点直联方式，将以服务接口为单位，忽略注册中心的提供者列表，A
> 接口配置点对点，不影响 B 接口从注册中心获取列表。
> 
> 
> 
> 
> 
> ![/user-guide/images/dubbo-directly.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0436fcee79c?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务引用（二）之远程引用（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.6> 只订阅】 ###

> 
> 
> 
> 为方便开发测试，经常会在线下共用一个所有服务可用的注册中心，这时，如果一个正在开发中的服务提供者注册，可能会影响消费者不能正常运行。
> 
> 
> 
> 可以让服务提供者开发方，只订阅服务(开发的服务可能依赖其它服务)，而不注册正在开发的服务，通过直连测试正在开发的服务。
> 
> 
> 
> 
> 
> ![/user-guide/images/subscribe-only.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a04370201ac6?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务引用（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.7> 只注册】 ###

> 
> 
> 
> 如果有两个镜像环境，两个注册中心，有一个服务只在其中一个注册中心有部署，另一个注册中心还没来得及部署，而两个注册中心的其它应用都需要依赖此服务。这个时候，可以让服务提供者方只注册服务到另一注册中心，而不从另一注册中心订阅服务。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务引用（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.8> 静态服务】 ###

> 
> 
> 
> 有时候希望人工管理服务提供者的上线和下线，此时需将注册中心标识为非动态管理模式。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 注册中心（二）之 Zookeeper》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（三）之 Redis》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.9> 多协议】 ###

> 
> 
> 
> Dubbo 允许配置多协议，在不同服务上支持不同协议或者同一服务上同时支持多种协议。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务暴露（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务暴露（二）之远程暴露（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.10> 多注册中心】 ###

> 
> 
> 
> Dubbo
> 支持同一服务向多注册中心同时注册，或者不同服务分别注册到不同的注册中心上去，甚至可以同时引用注册在不同注册中心上的同名服务。另外，注册中心是支持自定义扩展的。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务暴露（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务暴露（二）之远程暴露（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（一）之本地引用（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（二）之远程引用（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.11> 服务分组】 ###

> 
> 
> 
> 当一个接口有多种实现时，可以用 group 区分。
> 
> 

* [《精尽 Dubbo 源码分析 —— 服务暴露（二）之远程暴露（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（二）之远程引用（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.12> 多版本】 ###

> 
> 
> 
> 当一个接口实现，出现不兼容升级时，可以用版本号过渡，版本号不同的服务相互间不引用。
> 
> 
> 
> 可以按照以下的步骤进行版本迁移：
> 
> * 在低压力时间段，先升级一半提供者为新版本
> * 再将所有消费者升级为新版本
> * 然后将剩下的一半提供者升级为新版本
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务暴露（二）之远程暴露（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（二）之远程引用（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【2】同步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.13> 分组聚合】 ###

> 
> 
> 
> 按组合并返回结果 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fgroup-merger.html%23fn_1
> ) ，比如菜单服务，接口一样，但有多种实现，用group区分，现在消费方需从每种group中调用一次返回结果，合并结果返回，这样就可以实现聚合菜单项。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（五）之 Merger 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.14> 参数验证】 ###

> 
> 
> 
> 参数验证功能 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fparameter-validation.html%23fn_1
> ) 是基于 [JSR303](
> https://link.juejin.im?target=https%3A%2F%2Fjcp.org%2Fen%2Fjsr%2Fdetail%3Fid%3D303
> ) 实现的，用户只需标识 JSR303 标准的验证 annotation，并通过声明 filter 来实现验证 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fparameter-validation.html%23fn_2
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（十一）之 ValidationFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.15> 结果缓存】 ###

> 
> 
> 
> 结果缓存 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fresult-cache.html%23fn_1
> ) ，用于加速热门数据的访问速度，Dubbo 提供声明式缓存，以减少用户加缓存的工作量 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fresult-cache.html%23fn_2
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（十）之 CacheFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.16> 泛化引用】 ###

> 
> 
> 
> 泛化接口调用方式主要用于客户端没有 API 接口及模型类元的情况，参数及返回值中的所有 POJO 均用 ` Map` 表示，通常用于框架集成，比如：实现一个通用的服务测试框架，可通过
> ` GenericService` 调用所有服务实现。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 调用特性（二）之泛化引用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.17> 泛化实现】 ###

> 
> 
> 
> 泛接口实现方式主要用于服务器端没有API接口及模型类元的情况，参数及返回值中的所有POJO均用Map表示，通常用于框架集成，比如：实现一个通用的远程服务Mock框架，可通过实现GenericService接口处理所有服务请求。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 调用特性（三）之泛化实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.18> 回声测试】 ###

> 
> 
> 
> 回声测试用于检测服务是否可用，回声测试按照正常请求流程执行，能够测试整个调用是否通畅，可用于监控。
> 
> 
> 
> 所有服务自动实现 ` EchoService` 接口，只需将任意服务引用强制转型为 ` EchoService` ，即可使用。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 调用特性（一）之回音测试》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.19> 上下文信息】 ###

> 
> 
> 
> 上下文中存放的是当前调用过程中所需的环境信息。所有配置信息都将转换为 URL 的参数，参见 [schema 配置参考手册](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fxml%2Fintroduction.html
> ) 中的 **对应URL参数** 一列。
> 
> 
> 
> RpcContext 是一个 ThreadLocal 的临时状态记录器，当接收到 RPC 请求，或发起 RPC 请求时，RpcContext
> 的状态都会变化。比如：A 调 B，B 再调 C，则 B 机器上，在 B 调 C 之前，RpcContext 记录的是 A 调 B 的信息，在 B 调
> C 之后，RpcContext 记录的是 B 调 C 的信息。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 过滤器（二）之 ContextFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.20> 隐式参数】 ###

> 
> 
> 
> 可以通过 ` RpcContext` 上的 ` setAttachment` 和 ` getAttachment` 在服务消费方和提供方之间进行参数的隐式传递。
> [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fattachment.html%23fn_1
> )
> 
> 
> 
> 
> 
> ![/user-guide/images/context.png](https://user-gold-cdn.xitu.io/2018/6/23/1642a0437171bc42?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 过滤器（二）之 ContextFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.21> 异步调用】 ###

> 
> 
> 
> 基于 NIO 的非阻塞实现并行调用，客户端不需要启动多线程即可完成并行调用多个远程服务，相对多线程开销较小。 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fasync-call.html%23fn_1
> )
> 
> 
> 
> 
> 
> ![/user-guide/images/future.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0440adb05da?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【3】异步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.22> 本地调用】 ###

> 
> 
> 
> 本地调用使用了 injvm 协议，是一个伪协议，它不开启端口，不发起远程调用，只在 JVM 内直接关联，但执行 Dubbo 的 Filter 链。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务调用（一）之本地调用（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.23> 参数回调】 ###

> 
> 
> 
> 参数回调方式与调用本地 callback 或 listener 相同，只需要在 Spring 的配置文件中声明哪个参数是 callback
> 类型即可。Dubbo 将基于长连接生成反向代理，这样就可以从服务器端调用客户端逻辑 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fcallback-parameter.html%23fn_1
> ) 。可以参考 [dubbo 项目中的示例代码](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Falibaba%2Fdubbo%2Ftree%2Fmaster%2Fdubbo-test%2Fdubbo-test-examples%2Fsrc%2Fmain%2Fjava%2Fcom%2Falibaba%2Fdubbo%2Fexamples%2Fcallback
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【4】参数回调》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.24> 事件通知】 ###

> 
> 
> 
> 在调用之前、调用之后、出现异常时，会触发 ` oninvoke` 、 ` onreturn` 、 ` onthrow` 三个事件，可以配置当事件发生时，通知哪个类的哪个方法
> [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fevents-notify.html%23fn_1
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【3】异步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.25> 本地存根】 ###

> 
> 
> 
> 远程服务后，客户端通常只剩下接口，而实现全在服务器端，但提供方有些时候想在客户端也执行部分逻辑，比如：做 ThreadLocal
> 缓存，提前验证参数，调用失败后伪造容错数据等等，此时就需要在 API 中带上 Stub，客户端生成 Proxy 实例，会把 Proxy
> 通过构造函数传给 Stub [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Flocal-stub.html%23fn_1
> ) ，然后把 Stub 暴露给用户，Stub 可以决定要不要去调 Proxy。
> 
> 
> 
> 
> 
> ![/user-guide/images/stub.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0460af26609?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 动态代理（三）之本地存根 Stub》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.26> 本地伪装】 ###

> 
> 
> 
> 本地伪装 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Flocal-mock.html%23fn_1
> ) 通常用于服务降级，比如某验权服务，当服务提供方全部挂掉后，客户端不抛出异常，而是通过 Mock 数据返回授权失败。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（八）之 Mock 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.27> 延迟暴露】 ###

> 
> 
> 
> 如果你的服务需要预热时间，比如初始化缓存，等待相关资源就位等，可以使用 delay 进行延迟暴露。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— API 配置（二）之服务提供者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.28> 并发控制】 ###

> 
> 
> 
> 限制 ` com.foo.BarService` 的每个方法，服务器端并发执行（或占用线程池线程数）不能超过 10 个
> 
> 
> 
> 限制 ` com.foo.BarService` 的 ` sayHello` 方法，服务器端并发执行（或占用线程池线程数）不能超过 10 个
> 
> 
> 
> 限制 ` com.foo.BarService` 的每个方法，每客户端并发执行（或占用连接的请求数）不能超过 10 个
> 
> 
> 
> 限制 ` com.foo.BarService` 的 ` sayHello` 方法，每客户端并发执行（或占用连接的请求数）不能超过 10 个
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（四）之 ActiveLimitFilter && ExecuteLimitFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（四）之 Loadbalance 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.29> 连接控制】 ###

> 
> 
> 
> 限制服务器端接受的连接不能超过 10 个 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fconfig-connections.html%23fn_1
> ) ：
> 
> 
> 
> 限制客户端服务使用连接不能超过 10 个 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fconfig-connections.html%23fn_2
> ) ：
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（四）之 ActiveLimitFilter && ExecuteLimitFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.30> 延迟连接】 ###

> 
> 
> 
> 延迟连接用于减少长连接数。当有调用发起时，再创建长连接。 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Flazy-connect.html%23fn_1
> )
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【1】通信实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.31> 粘滞连接】 ###

> 
> 
> 
> 粘滞连接用于有状态服务，尽可能让客户端总是向同一提供者发起调用，除非该提供者挂了，再连另一台。
> 
> 
> 
> 粘滞连接将自动开启 [延迟连接](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Flazy-connect.html
> ) ，以减少长连接数。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.32> 令牌验证】 ###

> 
> 
> 
> 通过令牌验证在注册中心控制权限，以决定要不要下发令牌给消费者，可以防止消费者绕过注册中心访问提供者，另外通过注册中心可灵活改变授权方式，而不需修改或升级提供者
> 
> 
> 
> 
> 
> 
> ![/user-guide/images/dubbo-token.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0460958e98f?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（八）之 TokenFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.33> 路由规则】 ###

> 
> 
> 
> 路由规则 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Frouting-rule.html%23fn_1
> ) 决定一次 dubbo 服务调用的目标服务器，分为条件路由规则和脚本路由规则，并且支持可扩展 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Frouting-rule.html%23fn_2
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（七）之 Router 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.34> 配置规则】 ###

> 
> 
> 
> 向注册中心写入动态配置覆盖规则 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fconfig-rule.html%23fn_1
> ) 。该功能通常由监控中心或治理中心的页面完成。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（六）之 Configurator 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.35> 服务降级】 ###

> 
> 
> 
> 可以通过服务降级功能 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fservice-downgrade.html%23fn_1
> ) 临时屏蔽某个出错的非关键服务，并定义降级后的返回策略。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（八）之 Mock 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.36> 优雅停机】 ###

> 
> 
> 
> Dubbo 是通过 JDK 的 ShutdownHook 来完成优雅停机的，所以如果用户使用 ` kill -9 PID` 等强制关闭指令，是不会执行优雅停机的，只有通过
> ` kill PID` 时，才会执行。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 优雅停机》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.37> 主机绑定】 ###

> 
> 
> 
> 缺省主机 IP 查找顺序：
> 
> 
> 
> * 通过 ` LocalHost.getLocalHost()` 获取本机地址。
> * 如果是 ` 127.*` 等 loopback 地址，则扫描各网卡，获取网卡 IP。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— API 配置（二）之服务提供者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.38> 日志适配】 ###

> 
> 
> 
> 自 ` 2.2.1` 开始，dubbo 开始内置 log4j、slf4j、jcl、jdk 这些日志框架的适配 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Flogger-strategy.html%23fn_1
> ) ，也可以通过以下方式显示配置日志输出策略：
> 
> * 命令行
> * 在 ` dubbo.properties` 中指定
> * 在 ` dubbo.xml` 中配置
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 日志适配》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.39> 访问日志】 ###

> 
> 
> 
> 如果你想记录每一次请求信息，可开启访问日志，类似于apache的访问日志。 **注意** ：此日志量比较大，请注意磁盘容量。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 过滤器（三）之 AccessLogFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.40> 服务容器】 ###

> 
> 
> 
> 服务容器是一个 standalone 的启动程序，因为后台服务不需要 Tomcat 或 JBoss 等 Web 容器的功能，如果硬要用 Web
> 容器去加载服务提供方，增加复杂性，也浪费资源。
> 
> 
> 
> 服务容器只是一个简单的 Main 方法，并加载一个简单的 Spring 容器，用于暴露服务。
> 
> 
> 
> 服务容器的加载内容可以扩展，内置了 spring, jetty, log4j 等加载，可通过 [容器扩展点](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-dev-book%2Fcontent%2Fimpls%2Fcontainer.html
> ) 进行扩展。配置配在 java 命令的 -D 参数或者 ` dubbo.properties` 中。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务容器》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.41> ReferenceConfig 缓存】 ###

> 
> 
> 
> ` ReferenceConfig` 实例很重，封装了与注册中心的连接以及与提供者的连接，需要缓存。否则重复生成 ` ReferenceConfig`
> 可能造成性能问题并且会有内存和连接泄漏。在 API 方式编程时，容易忽略此问题。
> 
> 
> 
> 因此，自 ` 2.4.0` 版本开始， dubbo 提供了简单的工具类 ` ReferenceConfigCache` 用于缓存 `
> ReferenceConfig` 实例。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— API 配置（三）之服务消费者》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.42> 分布式事务】 ###

> 
> 
> 
> 分布式事务基于 JTA/XA 规范实现 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fdemos%2Fdistributed-transaction.html%23fn_1
> ) 。
> 
> 
> 
> 两阶段提交：
> 
> 
> 
> 
> 
> ![/user-guide/images/jta-xa.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0460bddd35d?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

目前该功能暂未实现，如下是推荐阅读的内容：

* [《RocketMQ 源码分析 —— 事务消息》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRocketMQ%2Fmessage-transaction%2F%3Fself )
* [《Sharding-JDBC 源码分析 —— 分布式事务（一）之最大努力型》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FSharding-JDBC%2Ftransaction-bed%2F%3Fself )
* [《MyCAT 源码分析 —— XA分布式事务》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FMyCAT%2Fxa-distributed-transaction%2F%3Fself )
* [《TCC-Transaction 源码解析 —— 精品合集》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2Fcategories%2FTCC-Transaction%2F%3Fself )
* [《Happylifeplat-TCC 源码解析 —— 精品合集》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FHappylifeplat-TCC%2Fgood-collection%2F%3Fself )
* [《Myth 源码解析 —— 精品合集》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FMyth%2Fgood-collection%2F%3Fself )

### 【<6.43> 线程栈自动dump】 ###

> 
> 
> 
> 当业务线程池满时，我们需要知道线程都在等待哪些资源、条件，以找到系统的瓶颈点或异常点。dubbo通过Jstack自动导出线程堆栈来保留现场，方便排查问题
> 
> 
> 
> 
> 默认策略:
> 
> 
> 
> * 导出路径，user.home标识的用户主目录
> * 导出间隔，最短间隔允许每隔10分钟导出一次
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 线程池》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<6.44> Netty4】 ###

> 
> 
> 
> dubbo 2.5.6版本新增了对netty4通信模块的支持
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— NIO 服务器（六）之 Netty4 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<8> schema 配置参考手册】 ###

> 
> 
> 
> 这里以 XML Config [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fxml%2Fintroduction.html%23fn_1
> ) 为准，列举所有配置项 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fxml%2Fintroduction.html%23fn_2
> ) 。其它配置方式，请参见相应转换关系： [属性配置](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fconfiguration%2Fproperties.html
> ) ， [注解配置](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fconfiguration%2Fannotation.html
> ) ， [API 配置](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Fconfiguration%2Fapi.html
> ) 。
> 
> 

对应源码解析文章：

* [《精近 Dubbo 源码解析 —— XML 配置》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.1> dubbo://】 ###

> 
> 
> 
> Dubbo 缺省协议采用单一长连接和 NIO 异步通讯，适合于小数据量大并发的服务调用，以及服务消费者机器数远大于服务提供者机器数的情况。
> 
> 
> 
> 反之，Dubbo 缺省协议不适合传送大数据量的服务，比如传文件，传视频等，除非请求量很低。
> 
> 
> 
> 
> 
> ![dubbo-protocol.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0436fe65f56?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【1】通信实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【2】同步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【3】异步调用》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 服务调用（二）之远程调用（Dubbo）【4】参数回调》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.2> rmi://】 ###

> 
> 
> 
> RMI 协议采用 JDK 标准的 ` java.rmi.*` 实现，采用阻塞式短连接和 JDK 标准序列化方式。
> 
> 
> 
> 注意：如果正在使用 RMI 提供服务给外部访问 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Frmi.html%23fn_1
> ) ，同时应用里依赖了老的 common-collections 包 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Frmi.html%23fn_2
> ) 的情况下，存在反序列化安全风险 [3](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Frmi.html%23fn_3
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（七）之远程调用（RMI）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.3> hessian://】 ###

> 
> 
> 
> Hessian [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fhessian.html%23fn_1
> ) 协议用于集成 Hessian 的服务，Hessian 底层采用 Http 通讯，采用 Servlet 暴露服务，Dubbo 缺省内嵌 Jetty
> 作为服务器实现。
> 
> 
> 
> Dubbo 的 Hessian 协议可以和原生 Hessian 服务互操作，即：
> 
> 
> 
> * 提供者用 Dubbo 的 Hessian 协议暴露服务，消费者直接用标准 Hessian 接口调用
> * 或者提供方用标准 Hessian 暴露服务，消费方用 Dubbo 的 Hessian 协议调用。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（四）之远程调用（Hessian）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.4> http://】 ###

> 
> 
> 
> 基于 HTTP 表单的远程调用协议，采用 Spring 的 HttpInvoker 实现 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fhttp.html%23fn_1
> )
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（三）之远程调用（HTTP）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.5> webservice://】 ###

> 
> 
> 
> 基于 WebService 的远程调用协议，基于 [Apache CXF](
> https://link.juejin.im?target=http%3A%2F%2Fcxf.apache.org%2F ) [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fwebservice.html%23fn_1
> ) 的 ` frontend-simple` 和 ` transports-http` 实现 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fwebservice.html%23fn_2
> ) 。
> 
> 
> 
> 可以和原生 WebService 服务互操作，即：
> 
> 
> 
> * 提供者用 Dubbo 的 WebService 协议暴露服务，消费者直接用标准 WebService 接口调用，
> * 或者提供方用标准 WebService 暴露服务，消费方用 Dubbo 的 WebService 协议调用。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（五）之远程调用（Webservice）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.6> thrift://】 ###

> 
> 
> 
> 当前 dubbo 支持 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fthrift.html%23fn_1
> ) 的 thrift 协议是对 thrift 原生协议 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fthrift.html%23fn_2
> ) 的扩展，在原生协议的基础上添加了一些额外的头信息，比如 service name，magic number 等。
> 
> 
> 
> 使用 dubbo thrift 协议同样需要使用 thrift 的 idl compiler 编译生成相应的 java
> 代码，后续版本中会在这方面做一些增强。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（十）之远程调用（Thrift）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.7> memcached://】 ###

> 
> 
> 
> 基于 memcached [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fmemcached.html%23fn_1
> ) 实现的 RPC 协议 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fmemcached.html%23fn_2
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（九）之远程调用（Memcached）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.8> redis://】 ###

> 
> 
> 
> 基于 Redis [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fredis.html%23fn_1
> ) 实现的 RPC 协议 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fprotocol%2Fredis.html%23fn_2
> ) 。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（八）之远程调用（Redis）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<9.9> rest://】 ###

对应文档为

* [《在Dubbo中开发REST风格的远程调用（RESTful Remoting）》]( https://link.juejin.im?target=https%3A%2F%2Fdangdangdotcom.github.io%2Fdubbox%2Frest.html ) 。

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务调用（六）之远程调用（REST）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<10.1> Multicast 注册中心】 ###

> 
> 
> 
> Multicast 注册中心不需要启动任何中心节点，只要广播地址一样，就可以互相发现。
> 
> 
> 
> 
> 
> ![/user-guide/images/multicast.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a04604cbd0dc?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> * 提供方启动时广播自己的地址
> * 消费方启动时广播订阅请求
> * 提供方收到订阅请求时，单播自己的地址给订阅者，如果设置了 ` unicast=false` ，则广播给订阅者
> * 消费方收到提供方地址时，连接该地址进行 RPC 调用。
> 
> 
> 组播受网络结构限制，只适合小规模应用或开发阶段使用。组播地址段: 224.0.0.0 - 239.255.255.255
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 注册中心（五）之 Multicast 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<10.2> zookeeper 注册中心】 ###

> 
> 
> 
> [Zookeeper](
> https://link.juejin.im?target=http%3A%2F%2Fzookeeper.apache.org%2F ) 是
> Apacahe Hadoop 的子项目，是一个树型的目录服务，支持变更推送，适合作为 Dubbo
> 服务的注册中心，工业强度较高，可用于生产环境，并推荐使用 [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fregistry%2Fzookeeper.html%23fn_1
> ) 。
> 
> 
> 
> 
> 
> ![/user-guide/images/zookeeper.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0460126952f?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— Zookeeper 客户端》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（二）之 Zookeeper》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<10.3> Redis 注册中心】 ###

> 
> 
> 
> 基于 Redis [1](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fregistry%2Fredis.html%23fn_1
> ) 实现的注册中心 [2](
> https://link.juejin.im?target=https%3A%2F%2Fdubbo.gitbooks.io%2Fdubbo-user-book%2Freferences%2Fregistry%2Fredis.html%23fn_2
> ) 。
> 
> 
> 
> 
> 
> ![/user-guide/images/dubbo-redis-registry.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0467aaf2e8f?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 注册中心（三）之 Redis》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<10.4> Simple 注册中心】 ###

> 
> 
> 
> Simple 注册中心本身就是一个普通的 Dubbo 服务，可以减少第三方依赖，使整体通讯方式一致。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 注册中心（四）之 Simple 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<11> Telnet 命令参考手册】 ###

> 
> 
> 
> 从 ` 2.0.5` 版本开始，dubbo 开始支持通过 telnet 命令来镜像服务治理。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— NIO 服务器（三）之 Telnet 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<12> 在线运维命令-QOS】 ###

> 
> 
> 
> dubbo 2.5.8 新版本重构了 telnet 模块，提供了新的 telnet 命令支持。
> 
> 

对应源码解析文章：

TODO

` dubbo-ops`

## 1.3 Dubbo 开发指南 ##

本小节，我们将 [《精尽 Dubbo 源码解析》]( # ) 和 [《Dubbo 开发指南》]( https://link.juejin.im?target=https%3A%2F%2Fwww.gitbook.com%2Fbook%2Fdubbo%2Fdubbo-user-book ) 做一次映射，方便大家直接找到感兴趣的功能的具体源码实现。当然，如果有整理不到位的地方，欢迎大家给我留言斧正。

### 【<1> 源码构建】 ###

> 
> 
> 
> 代码签出
> 分支
> 构建
> 构建源代码 jar 包
> IDE 支持
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 调试环境搭建》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<2> 框架设计】 ###

> 
> 
> 
> 
> 
> ![/dev-guide/images/dubbo-framework.jpg](https://user-gold-cdn.xitu.io/2018/6/23/1642a0467ca8aa9f?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 项目结构一览》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精近 Dubbo 源码解析 —— 核心流程一览》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<3> 扩展点加载】 ###

> 
> 
> 
> Dubbo 的扩展点加载从 JDK 标准的 SPI (Service Provider Interface) 扩展点发现机制加强而来。
> 
> 
> 
> Dubbo 改进了 JDK 标准的 SPI 的以下问题：
> 
> 
> 
> * JDK 标准的 SPI 会一次性实例化扩展点所有实现，如果有扩展实现初始化很耗时，但如果没用上也加载，会很浪费资源。
> * 如果扩展点加载失败，连扩展点的名称都拿不到了。比如：JDK 标准的 ScriptEngine，通过 ` getName()` 获取脚本类型的名称，但如果
> RubyScriptEngine 因为所依赖的 jruby.jar 不存在，导致 RubyScriptEngine
> 类加载失败，这个失败原因被吃掉了，和 ruby 对应不起来，当用户执行 ruby 脚本时，会报不支持 ruby，而不是真正失败的原因。
> * 增加了对扩展点 IoC 和 AOP 的支持，一个扩展点可以直接 setter 注入其它扩展点。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 拓展机制 SPI》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.1> 协议扩展】 ###

> 
> 
> 
> RPC 协议扩展，封装远程调用细节。
> 
> 
> 
> 契约：
> 
> 
> 
> * 当用户调用 ` refer()` 所返回的 ` Invoker` 对象的 ` invoke()` 方法时，协议需相应执行同 URL 远端 `
> export()` 传入的 ` Invoker` 对象的 ` invoke()` 方法。
> * 其中， ` refer()` 返回的 ` Invoker` 由协议实现，协议通常需要在此 ` Invoker` 中发送远程请求， `
> export()` 传入的 ` Invoker` 由框架实现并传入，协议不需要关心。
> 
> 
> 
> 
> 注意：
> 
> 
> 
> * 协议不关心业务接口的透明代理，以 ` Invoker` 为中心，由外层将 ` Invoker` 转换为业务接口。
> * 协议不一定要是 TCP 网络通讯，比如通过共享文件，IPC 进程间通讯等。
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务暴露（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务暴露（二）之远程暴露（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（一）之本地引用（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 服务引用（二）之远程引用（Dubbo）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.2> 调用拦截扩展】 ###

> 
> 
> 
> 服务提供方和服务消费方调用过程拦截，Dubbo 本身的大多功能均基于此扩展点实现，每次远程方法执行，该拦截都会被执行，请注意对性能的影响。
> 
> 
> 
> 约定：
> 
> 
> 
> * 用户自定义 filter 默认在内置 filter 之后。
> * 特殊值 ` default` ，表示缺省扩展点插入的位置。比如： ` filter="xxx,default,yyy"` ，表示 ` xxx` 在缺省
> filter 之前， ` yyy` 在缺省 filter 之后。
> * 特殊符号 ` -` ，表示剔除。比如： ` filter="-foo1"` ，剔除添加缺省扩展点 ` foo1` 。比如： `
> filter="-default"` ，剔除添加所有缺省扩展点。
> * provider 和 service 同时配置的 filter 时，累加所有 filter，而不是覆盖。比如： ` <dubbo:provider
> filter="xxx,yyy"/>` 和 ` <dubbo:service filter="aaa,bbb" />` ，则 ` xxx` , `
> yyy` , ` aaa` , ` bbb` 均会生效。如果要覆盖，需配置： ` <dubbo:service
> filter="-xxx,-yyy,aaa,bbb" />`
> 
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 过滤器（一）之 ClassLoaderFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 过滤器（二）之 ContextFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 过滤器（三）之 AccessLogFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（四）之 ActiveLimitFilter && ExecuteLimitFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（五）之 TimeoutFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（六）之 DeprecatedFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（七）之 ExceptionFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（八）之 TokenFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（九）之 TpsLimitFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（十）之 CacheFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 过滤器（十一）之 ValidationFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.3> 引用监听扩展】 ###

> 
> 
> 
> 当有服务引用时，触发该事件。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务引用（一）之本地引用（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.4> 暴露监听扩展】 ###

> 
> 
> 
> 当有服务暴露时，触发该事件。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 服务暴露（一）之本地暴露（Injvm）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.5> 集群扩展】 ###

> 
> 
> 
> 当有多个服务提供方时，将多个服务提供方组织成一个集群，并伪装成一个提供方。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码解析 —— 集群容错（二）之 Cluster 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.6> 路由扩展】 ###

> 
> 
> 
> 从多个服务提者方中选择一个进行调用。
> 
> 

* 官方文档写的有问题，从多个服务提者方中筛选出所有匹配的结果集。

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（七）之 Router 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.7> 负载均衡扩展】 ###

> 
> 
> 
> 从多个服务提者方中选择一个进行调用
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（四）之 Loadbalance 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.8> 合并结果扩展】 ###

> 
> 
> 
> 合并返回结果，用于分组聚合。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 集群容错（五）之 Merger 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.9> 注册中心扩展】 ###

> 
> 
> 
> 负责服务的注册与发现。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— Zookeeper 客户端》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（二）之 Zookeeper》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（三）之 Redis》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（四）之 Simple 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 注册中心（五）之 Multicast 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.10> 监控中心扩展】 ###

> 
> 
> 
> 负责服务调用次和调用时间的监控。
> 
> 

对应源码解析文章：

TODO

### 【<5.11> 扩展点加载扩展】 ###

> 
> 
> 
> 扩展点本身的加载容器，可从不同容器加载扩展点。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 拓展机制 SPI》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.12> 动态代理扩展】 ###

> 
> 
> 
> 将 ` Invoker` 接口转换成业务接口。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 动态代理（一）之 Javassist》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 动态代理（二）之 JDK》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.13> 编译器扩展】 ###

> 
> 
> 
> Java 代码编译器，用于动态生成字节码，加速调用。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 动态编译（一）之 Javassist》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 动态编译（二）之 JDK》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.14> 消息派发扩展】 ###

> 
> 
> 
> 通道信息派发器，用于指定线程池模型。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— NIO 服务器（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.15> 线程池扩展】 ###

> 
> 
> 
> 服务提供方线程程实现策略，当服务器收到一个请求时，需要在线程池中创建一个线程去执行服务提供方业务逻辑。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 线程池》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.16> 序列化扩展】 ###

> 
> 
> 
> 将对象转成字节流，用于网络传输，以及将字节流转为对象，用于在收到字节流数据后还原成对象。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 序列化（一）之 总体实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 序列化（二）之 Dubbo 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— 序列化（三）之 Kryo 实现》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.17> 网络传输扩展】 ###

> 
> 
> 
> 远程通讯的服务器及客户端传输实现。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— NIO 服务器（一）之抽象 API》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )
* [《精尽 Dubbo 源码分析 —— NIO 服务器（二）之 Transport 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.18> 信息交换扩展】 ###

> 
> 
> 
> 基于传输层之上，实现 Request-Response 信息交换语义。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— NIO 服务器（四）之 Exchange 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.19> 组网扩展】 ###

> 
> 
> 
> 对等网络节点组网器。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 注册中心（五）之 Multicast 》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.20> Telnet 命令扩展】 ###

> 
> 
> 
> 所有服务器均支持 telnet 访问，用于人工干预。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— NIO 服务器（三）之 Telnet 层》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.21> 状态检查扩展】 ###

> 
> 
> 
> 检查服务依赖各种资源的状态，此状态检查可同时用于 telnet 的 status 命令和 hosting 的 status 页面。
> 
> 

对应源码解析文章：

TODO

### 【<5.22> 容器扩展】 ###

> 
> 
> 
> 服务容器扩展，用于自定义加载内容。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 服务容器》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.23> 页面扩展】 ###

> 
> 
> 
> 对等网络节点组网器。
> 
> 

* 官方文档写的有问题，应该是请求处理器，有点类似 Servlet 。

对应源码解析文章：

TODO

### 【<5.24> 缓存扩展】 ###

> 
> 
> 
> 用请求参数作为 key，缓存返回结果。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（十）之 CacheFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.25> 验证扩展】 ###

> 
> 
> 
> 参数验证扩展点。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码分析 —— 过滤器（十一）之 ValidationFilter》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

### 【<5.26> 日志适配扩展】 ###

> 
> 
> 
> 日志输出适配扩展点。
> 
> 

对应源码解析文章：

* [《精尽 Dubbo 源码解析 —— 日志适配》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FDubbo%2Fgood-collection%2F%3Fjuejin )

## 1.4 Dubbo 运维指南 ##

[《Dubbo 运维指南》]( https://link.juejin.im?target=http%3A%2F%2Fdubbo.apache.org%2Fbooks%2Fdubbo-admin-book%2F ) ，暂时无映射的需要。

# 2. 【老徐】RPC 专栏 #

* 作者 ：老徐
* 博客 ：https://www.cnkirito.moe/categories/RPC/
* 目录 ：

* [《简单了解 RPC 实现原理》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Feasy-know-rpc%2F )
* [《深入理解 RPC 之序列化篇 -- Kryo》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-serialize-1%2F )
* [《深入理解 RPC 之序列化篇 -- 总结篇》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-serialize-2%2F )
* [《深入理解 RPC 之动态代理篇》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-dynamic-proxy%2F )
* [《深入理解 RPC 之传输篇》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-transport%2F )
* [《Motan 中使用异步 RPC 接口》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Fmotan-async%2F )
* [《深入理解 RPC 之协议篇》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-protocol%2F )
* [《深入理解RPC之服务注册与发现篇》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-registry%2F )
* [《深入理解 RPC 之集群篇》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Frpc-cluster%2F )
* [《【千米网】从跨语言调用到dubbo2.js》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2Fdubbojs-in-qianmi%2F )
* [《天池中间件大赛 Dubbo Mesh 优化总结（QPS 从 1000 到 6850）》]( https://link.juejin.im?target=http%3A%2F%2Fwww.iocoder.cn%2FRPC%2Flaoxu%2FdubboMesh%2F )

# 3. 【肥朝】Dubbo 源码解析 #

* 作者 ：肥朝
* 博客 ：http://www.jianshu.com/u/f7daa458b874
* 目录 ：

* [《Dubbo 源码解析 —— 集群容错架构设计》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483767%26amp%3Bidx%3D1%26amp%3Bsn%3Dfaf031cdc362599276d3cc58598dd51d%26amp%3Bchksm%3Dfa497ec6cd3ef7d0729f6dff9baa116b91dfa6624ffac618620d46558d1d23afeb4b9cf789d2%23rd )
* [《Dubbo 源码解析 —— Directory》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483776%26amp%3Bidx%3D1%26amp%3Bsn%3D0410235af44b2991c163cfdfefeb26e4%26amp%3Bchksm%3Dfa497e31cd3ef727d30b1000a6e805c1c69fe65aef5b771d93b5283be4141850d234d5d765c0%23rd )
* [《Dubbo 源码解析 —— Router》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483785%26amp%3Bidx%3D1%26amp%3Bsn%3Da858a8cef7ecd86ac966138bfc28e6e0%26amp%3Bchksm%3Dfa497e38cd3ef72e3a16ef1c7294c379c73785de6e64388ec3e82986eda9bb4dfffbbb29b81e%23rd )
* [《Dubbo 源码解析 —— Cluster》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483794%26amp%3Bidx%3D1%26amp%3Bsn%3D02f1685fc1b0d32e3490d4d7536d6a6e%26amp%3Bchksm%3Dfa497e23cd3ef7351e30893cc79205fd684f69d1643056342b16dcdd20ac0293d1bcbaae60ab%23rd )
* [《Dubbo 源码解析 —— LoadBalance》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483815%26amp%3Bidx%3D1%26amp%3Bsn%3D9829fd2fe1eb03b1266f03415d1d305f%26amp%3Bchksm%3Dfa497e16cd3ef7005df4713fcaa7394022b268752ab0ac66fc15fff5c49e7c920fa6d5cebd33%23rd )
* [《Dubbo 源码解析 —— 服务暴露原理》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483828%26amp%3Bidx%3D1%26amp%3Bsn%3Dcf831fe49cad554b82e2add9b08dcf97%26amp%3Bchksm%3Dfa497e05cd3ef71365d1633fe3b3a2b81a929dafbdab510e4ed9a03bb43e4937a403fa25cac6%23rd )
* [《Dubbo 源码解析 —— 本地暴露》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483861%26amp%3Bidx%3D1%26amp%3Bsn%3D045bb16117e6175d9e6310905aa04405%26amp%3Bchksm%3Dfa497e64cd3ef7725e351efa3e35934caf46f109569fc97ce1cab77ba6460ca06519d3d6b8c7%23rd )
* [《Dubbo 源码解析 —— 远程暴露》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483887%26amp%3Bidx%3D1%26amp%3Bsn%3Df63868d98907959a2aea17f55d4fe0aa%26amp%3Bchksm%3Dfa497e5ecd3ef748cecd9641c034136b2e4502ba746d5d2dd6c95547bc73a1410377bb81eae2%23rd )
* [《Dubbo源码解析 —— Zookeeper 连接》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483909%26amp%3Bidx%3D1%26amp%3Bsn%3D7274f1e44d180f138a9e57c9ebd712e7%26amp%3Bchksm%3Dfa497db4cd3ef4a215b5109dec6a1269eae7207543640c997ead58fb15427c33c7e216ea152f%23rd )
* [《Dubbo源码解析 —— Zookeeper 创建节点》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483934%26amp%3Bidx%3D1%26amp%3Bsn%3Df22159486d50ee20f1d5400c3e70e51a%26amp%3Bchksm%3Dfa497dafcd3ef4b9aca5e3608ba7cfcd6dc28220e62030c97c1b091225ba84184df296ce5f09%23rd )
* [《Dubbo源码解析 —— 服务暴露总结》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247483961%26amp%3Bidx%3D1%26amp%3Bsn%3D1c26d1ee5280175ad1663e0c90f71cb5%26amp%3Bchksm%3Dfa497d88cd3ef49e310bcd3ac0bdb8a408504d3643975238816328066d88e787eac358adfc03%23rd )
* [《Dubbo源码解析 —— Zookeeper 订阅》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247484014%26amp%3Bidx%3D1%26amp%3Bsn%3D0b2e2efec6668d33166aca571add19da%26amp%3Bchksm%3Dfa497ddfcd3ef4c96d8a002ecd83b28660dddb7675b2177106aea1966e3c8ae78c04e01acc69%23rd )
* [《Dubbo源码解析 —— 逻辑层设计之服务降级》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247484096%26amp%3Bidx%3D1%26amp%3Bsn%3Dc6bf3c48ca3c95949fea5816dbdfa50c%26amp%3Bchksm%3Dfa497d71cd3ef467fba2578dab66a965b9e7c4f82bf7072f7179a3424e08eb993ddc5660a2d7%23rd )
* [《Dubbo源码解析 —— 简单原理、与spring融合》]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247484261%26amp%3Bidx%3D1%26amp%3Bsn%3De9526ff0b6e2b127ed7dd5dbfb694190%26amp%3Bchksm%3Dfa497cd4cd3ef5c28f771127273d536cb244d24aa118278dfe52e46fd4db5e99bccf59a0c6d5%23rd )
* [《Dubbo源码解析 —— 服务引用原理》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUzMTA2NTU2Ng%3D%3D%26amp%3Bmid%3D2247484325%26amp%3Bidx%3D1%26amp%3Bsn%3D1c686d832f5e88aacda2c089be8d831f%26amp%3Bchksm%3Dfa497c14cd3ef502a82eb7f8f4a57703762912911dff01968aabe9d42a9c695f9ce8ad13ed76%23rd )

# 4. 【MR_QI】Dubbo 源码解析 #

* 作者 ：MR_QI
* 博客 ：https://my.oschina.net/qixiaobo025/blog 或 http://qixiaobo.site/tags/dubbo/
* 目录 ：

* [《dubbo源码系列之filter的前生》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F995254 )
* [《dubbo源码系列之filter的今世》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F995281 )
* [《为什么说dubbo的声明式缓存不好用！！！》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F995772 )
* [《Dubbo优雅服务降级之Stub和回声服务》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1014845 )
* [《Dubbo序列化之hessian2》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1073902 )
* [《Dubbo异常处理》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1142720 )
* [《Dubbo自定义异常message过长解决》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1153876 )
* [《Dubbo客户端返回自定义异常》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1154492 )
* [《Dubbo超时控制源码分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1186779 )
* [《Dubbo之telnet实现》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1417321 )
* [《Fst反序列化失败》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1519566 )
* [《Dubbo配置直连》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fqixiaobo025%2Fblog%2F1527009 )

# 5. 【杨武兵】Dubbo 源码解析 #

* 作者 ：杨武兵
* 博客 ：https://my.oschina.net/ywbrj042/blog
* 目录 ：

* [《dubbo源码分析系列——项目工程结构介绍》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F683515 )
* [《dubbo源码分析系列——dubbo-rpc-api模块源码分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F683719 )
* [《dubbo源码分析系列——dubbo-rpc-default模块源码分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F684718 )
* [《dubbo源码分析系列——dubbo的SPI机制源码分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F687443 )
* [《dubbo源码分析系列——dubbo的SPI机制自适应Adpative类分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F688042 )
* [《dubbo源码分析系列——dubbo-cluster模块源码分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F689818 )
* [《dubbo源码分析系列——dubbo-register-api模块源码分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F690342 )
* [《dubbo典型协议、传输组件、序列化方式组合性能对比测试》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F690691 )
* [《dubbo核心流程分析》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F702521 )
* [《dubbo服务中的hessian序列化工厂使用hashmap加锁在高并发场景下的问题》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F466151 )
* [《《分布式服务框架原理与实践》——第5章协议栈阅读笔记》]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fywbrj042%2Fblog%2F713403 )

# 666. 欢迎投稿 #

![](https://user-gold-cdn.xitu.io/2018/6/23/1642a041b825721e?imageView2/0/w/1280/h/960/ignore-error/1)