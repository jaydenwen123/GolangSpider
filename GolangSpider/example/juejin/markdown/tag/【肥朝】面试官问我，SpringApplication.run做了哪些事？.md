# 【肥朝】面试官问我，SpringApplication.run做了哪些事？ #

## 前言 ##

本篇题材仍然是源于肥朝粉丝在面试中遇到的问题

![](https://user-gold-cdn.xitu.io/2019/4/19/16a3158ac4e1d318?imageView2/0/w/1280/h/960/ignore-error/1)

坦白说,每天的消息挺多的,经常看不过来.正当我肥手即将要把聊天窗口划走时,他用简短的几句话, ` 彻底打动了我!`

![](https://user-gold-cdn.xitu.io/2019/4/19/16a3158e522a4e4b?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/4/19/16a31590520a75d3?imageView2/0/w/1280/h/960/ignore-error/1)

## 直入主题 ##

该问题,我们可以采用小学语文老师教给我们写作文的常用套路, ` 总分总`

> 
> 
> 
> 总
> 
> 

` SpringApplication.run` 一共做了两件事,分别是

* 

创建 ` SpringApplication` 对象

* 

利用创建好的 ` SpringApplication` 对象,调用 ` run` 方法

![](https://user-gold-cdn.xitu.io/2019/4/19/16a3159c8b79770c?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 分
> 
> 

1.创建 ` SpringApplication` 对象

![](https://user-gold-cdn.xitu.io/2019/4/19/16a315a3b9a10fa4?imageView2/0/w/1280/h/960/ignore-error/1)

2.调用 ` run` 方法

![](https://user-gold-cdn.xitu.io/2019/4/19/16a315a6df6a79a6?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 总
> 
> 

太多类名什么的记不住?没关系.上述内容都给你总结好了:

> 
> 
> 
> 面试官: 我看到你简历上写着熟悉SpringBoot,那你讲一下,SpringApplication.run都做了些什么？
> 
> 

> 
> 
> 
> 肥朝公众号粉丝:
> SpringApplication.run一共做了两件事,一件是创建SpringApplication对象,在该对象初始化时,找到配置的事件监听器,并保存起来.第二件事就是运行run方法,此时会将刚才保存的事件监听器根据当前时机触发不同的事件,比如容器初始化,容器创建完成等.同时也会刷新IoC容器,进行组件的扫描、创建、加载等工作.这两件事我都看过源码,我分别给你画个图细致讲一讲.
> 
> 
> 

> 
> 
> 
> 面试官:
> 
> 

![](https://user-gold-cdn.xitu.io/2019/4/19/16a315ac8f1d83aa?imageView2/0/w/1280/h/960/ignore-error/1)

**肥朝 是一个专注于 原理、源码、开发技巧的技术公众号，号内原创专题式源码解析、真实场景源码原理实战（重点）。 扫描下面二维码 关注肥朝，让本该造火箭的你，不再拧螺丝！**

![](https://user-gold-cdn.xitu.io/2019/4/19/16a315bd3c6fba8b?imageslim)