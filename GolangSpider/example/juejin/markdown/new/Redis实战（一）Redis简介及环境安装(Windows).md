# Redis实战（一）Redis简介及环境安装(Windows) #

提到Redis，大家肯定都听过，并且应该都在项目中或多或少的使用过，也许你觉得Redis用起来挺简单的呀，但如果有人问你下面的几个问题(比如同事或者面试官)，你能回答的上来吗？

* 什么是Redis？
* Redis能存储哪几种数据结构？
* Redis有几种持久化机制？它们的优缺点分别是什么？
* 哪些场景需要使用Redis？
* 什么是缓存雪崩，如何避免？
* 什么是缓存穿透，如何避免？

如果你都能回答的上来，恭喜你，说明你对Redis有一定的了解，如果回答不上来，也没关系，本系列博客会对Redis进行一系列的讲解，欢迎关注！

所谓工欲善其事，必先利其器，既然要学习Redis，首先我们至少得知道什么是Redis以及如何安装Redis环境，这也是本篇博客的主要内容。

## 1. Redis简介 ##

什么是Redis呢？

Redis是一个开源(BSD许可)的内存数据结构存储，用作数据库、缓存和消息代理。它支持诸如字符串、散列、列表、集合、有序集合等数据结构。 **-- [Redis官网]( https://link.juejin.im?target=https%3A%2F%2Fredis.io%2F )**

Redis是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的高性能的key-value数据库。 **-- 百度百科**

Redis是一款依据BSD开源协议发行的高性能key-value存储系统，通常被称为数据结构服务器。 **-- 其它网友**

Redis是一个远程内存数据库，它不仅性能强劲，而且还具有复制特性以及为解决问题而生的独一无二的数据模型。Redis提供了5种不同类型的数据结构，各式各样的问题都可以很自然地映射到这些数据结构上。 **-- 《Redis实战》**

Redis是一个速度非常快的非关系型数据库，它可以存储键(key)与5种不同类型值(value)之间的映射(mapping)，可以将存储在内存的键值对数据持久化到硬盘，可以使用复制特性来扩展读性能，还可以使用客户端分片来扩展写性能。 **-- 《Redis实战》**

## 2. Redis环境安装(Windows) ##

**说明：Redis官方并没有提供Windows版本的Redis，也不建议在生产环境使用Windows版本的Redis，我目前所在的公司生产环境Redis是部署在Linux服务器的。**

虽然Redis官方不支持Windows版本，但是微软 Microsoft Open Tech Group 提供了1个Windows版本的Redis，下载地址： [github.com/microsoftar…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmicrosoftarchive%2Fredis%2Freleases )

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b5418a4295c8?imageView2/0/w/1280/h/960/ignore-error/1)

将下载好的文件解压到你喜欢的目录，我这里是E:\Tools\Redis-x64-3.0.504，如下所示：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b54188a14058?imageView2/0/w/1280/h/960/ignore-error/1)

双击上图中红色标记的redis-server.exe即可启动Redis服务：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b5418892a548?imageView2/0/w/1280/h/960/ignore-error/1)

也可以打开一个cmd窗口，切换到Redis所在目录，然后执行如下命令启动：

` redis-server.exe redis.windows.conf 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b5418a276ee5?imageView2/0/w/1280/h/960/ignore-error/1)

通过这2种方式打开，需要保证cmd窗口一直保持打开状态，关闭后客户端就无法连接，如果服务器重启了，需要再次打开Redis服务端，为了解决该问题，我们可以把Redis安装成Windows服务：

` cd E:\Tools\Redis-x64- 3.0. 504 redis-server --service-install redis.windows.conf 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b5418c018973?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b5418c1d59b5?imageView2/0/w/1280/h/960/ignore-error/1)

你可以直接在界面上启动/停止该服务，也可以执行cmd命令来启动/停止/卸载该服务：

卸载服务：

` redis-server --service-uninstall 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541ae712397?imageView2/0/w/1280/h/960/ignore-error/1)

启动服务：

` redis-server --service-start 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541be0199ad?imageView2/0/w/1280/h/960/ignore-error/1)

停止服务：

` redis-server --service-stop 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541c6477166?imageView2/0/w/1280/h/960/ignore-error/1)

## 3. Redis Hello World示例 ##

打开cmd窗口，打开一个客户端来简单使用下Redis：

` redis-cli.exe -h 127.0. 0.1 -p 6379 复制代码`

设置一个key-value缓存，其中key为hello,value为hello world!：

` set hello "hello world!" 复制代码`

获取key为hello的值：

` get hello 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541bacfb0c0?imageView2/0/w/1280/h/960/ignore-error/1)

## 4. Redis Desktop Manager使用 ##

虽然我们可以通过命令的方式来查看Redis存储的数据，但毕竟不太友好，这里推荐个比较流行的工具：Redis Desktop Manager。

官网地址： [redisdesktop.com/]( https://link.juejin.im?target=https%3A%2F%2Fredisdesktop.com%2F ) 。

官网现在的版本2019.1需要先赞助付费才能使用。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541bb5cfdfa?imageView2/0/w/1280/h/960/ignore-error/1)

不过我们仍然可以下载之前不付费的版本，下载地址： [github.com/uglide/Redi…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fuglide%2FRedisDesktopManager%2Freleases%2Ftag%2F0.8.8 ) 。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541bc53a79c?imageView2/0/w/1280/h/960/ignore-error/1)

安装过程比较简单，这里不再赘述，安装完成后，连接本机Redis服务端：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541e0436a35?imageView2/0/w/1280/h/960/ignore-error/1)

连接成功后，可以看到之前设置的值：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b541d3280d08?imageView2/0/w/1280/h/960/ignore-error/1)

后续文章会讲解Linux环境安装Redis的方式，Redis的5种数据结构，持久化机制等，敬请期待……

## 5. 参考 ##

[Redis的安装和部署（windows ）]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhellogt%2Fp%2F6954263.html )

[Windows下使用Redis（一）安装使用]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fpanchunting%2Fp%2FRedis_On_Windows_Install.html )