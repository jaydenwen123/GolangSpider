# 七张图彻底讲清楚ZooKeeper分布式锁的实现原理【石杉的架构笔记】 #

欢迎关注个人公众号：石杉的架构笔记（ID:shishan100）

周一至周五早8点半！精品技术文章准时送上！

# 一、写在前面 #

之前写过一篇文章（《 [拜托，面试请不要再问我Redis分布式锁的实现原理]( https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Fjuejin.im%252525252Fpost%252525252F5bf3f15851882526a643e207 ) 》），给大家说了一下Redisson这个开源框架是如何实现Redis分布式锁原理的，这篇文章再给大家聊一下ZooKeeper实现分布式锁的原理。

同理，我是直接基于比较常用的 **Curator** 这个开源框架，聊一下这个框架对ZooKeeper（以下简称zk）分布式锁的实现。

一般除了大公司是自行封装分布式锁框架之外，建议大家用这些开源框架封装好的分布式锁实现，这是一个比较快捷省事儿的方式。

# 二、ZooKeeper分布式锁机制 #

接下来我们一起来看看， **多客户端获取及释放zk分布式锁的整个流程及背后的原理。**

首先大家看看下面的图，如果现在有两个客户端一起要争抢zk上的一把分布式锁，会是个什么场景？

![](https://user-gold-cdn.xitu.io/2018/11/30/167653167e579f54?imageView2/0/w/1280/h/960/ignore-error/1)

如果大家对zk还不太了解的话，建议先自行百度一下，简单了解点基本概念，比如zk有哪些节点类型等等。

参见上图。zk里有一把锁，这个锁就是zk上的一个节点。然后呢，两个客户端都要来获取这个锁，具体是怎么来获取呢？

咱们就假设客户端A抢先一步，对zk发起了加分布式锁的请求，这个加锁请求是用到了zk中的一个特殊的概念，叫做 **“临时顺序节点”。**

简单来说，就是直接在"my_lock"这个锁节点下，创建一个顺序节点，这个顺序节点有zk内部自行维护的一个节点序号。

比如说，第一个客户端来搞一个顺序节点，zk内部会给起个名字叫做：xxx-000001。然后第二个客户端来搞一个顺序节点，zk可能会起个名字叫做：xxx-000002。大家注意一下， **最后一个数字都是依次递增的** ，从1开始逐次递增。zk会维护这个顺序。

所以这个时候，假如说客户端A先发起请求，就会搞出来一个顺序节点，大家看下面的图，Curator框架大概会弄成如下的样子：

![](https://user-gold-cdn.xitu.io/2018/11/30/1676531853f4719a?imageView2/0/w/1280/h/960/ignore-error/1)

大家看，客户端A发起一个加锁请求，先会在你要加锁的node下搞一个临时顺序节点，这一大坨长长的名字都是Curator框架自己生成出来的。

然后，那个最后一个数字是"1"。大家注意一下，因为客户端A是第一个发起请求的，所以给他搞出来的顺序节点的序号是"1"。

接着客户端A创建完一个顺序节点。还没完，他会查一下" **my_lock** "这个锁节点下的所有子节点，并且这些子节点是按照序号排序的，这个时候他大概会拿到这么一个集合：

![](https://user-gold-cdn.xitu.io/2018/11/30/16765323dc0f42fc?imageView2/0/w/1280/h/960/ignore-error/1)

接着客户端A会走一个关键性的判断，就是说：唉！兄弟，这个集合里，我创建的那个顺序节点，是不是排在第一个啊？

如果是的话，那我就可以加锁了啊！因为明明我就是第一个来创建顺序节点的人，所以我就是第一个尝试加分布式锁的人啊！

bingo！ **加锁成功** ！大家看下面的图，再来直观的感受一下整个过程。

![](https://user-gold-cdn.xitu.io/2018/11/30/16765319dbc531f6?imageView2/0/w/1280/h/960/ignore-error/1)

接着假如说，客户端A都加完锁了，客户端B过来想要加锁了，这个时候他会干一样的事儿：先是在" **my_lock** "这个锁节点下创建一个 **临时顺序节点** ，此时名字会变成类似于：

![](https://user-gold-cdn.xitu.io/2018/11/30/16765326ccbc25ac?imageView2/0/w/1280/h/960/ignore-error/1)

**大家看看下面的图：**

![](https://user-gold-cdn.xitu.io/2018/11/30/1676531b1a2a4a8a?imageView2/0/w/1280/h/960/ignore-error/1)

客户端B因为是第二个来创建顺序节点的，所以zk内部会维护序号为"2"。

接着客户端B会走加锁判断逻辑，查询" **my_lock** "锁节点下的所有子节点，按序号顺序排列，此时他看到的类似于：

![](https://user-gold-cdn.xitu.io/2018/11/30/16765325567d7f4e?imageView2/0/w/1280/h/960/ignore-error/1)

同时检查自己创建的顺序节点，是不是集合中的第一个？

明显不是啊，此时第一个是客户端A创建的那个顺序节点，序号为"01"的那个。 **所以加锁失败** ！

加锁失败了以后，客户端B就会通过ZK的API对他的顺序节点的 **上一个顺序节点加一个监听器。** zk天然就可以实现对某个节点的监听。

如果大家还不知道zk的基本用法，可以百度查阅，非常的简单。客户端B的顺序节点是：

![](https://user-gold-cdn.xitu.io/2018/11/30/16765327d85f86fe?imageView2/0/w/1280/h/960/ignore-error/1)

他的上一个顺序节点，不就是下面这个吗？

![](https://user-gold-cdn.xitu.io/2018/11/30/16765328d82ccbed?imageView2/0/w/1280/h/960/ignore-error/1)

即客户端A创建的那个顺序节点！

所以，客户端B会对：

![](https://user-gold-cdn.xitu.io/2018/11/30/1676532a29a80f96?imageView2/0/w/1280/h/960/ignore-error/1)

这个节点加一个监听器，监听这个节点是否被删除等变化！大家看下面的图。

![](https://user-gold-cdn.xitu.io/2018/11/30/1676531c7057d0a2?imageView2/0/w/1280/h/960/ignore-error/1)

接着，客户端A加锁之后，可能处理了一些代码逻辑，然后就会释放锁。那么，释放锁是个什么过程呢？

其实很简单，就是把自己在zk里创建的那个顺序节点，也就是：

![](https://user-gold-cdn.xitu.io/2018/11/30/16765485aa6a55b6?imageView2/0/w/1280/h/960/ignore-error/1)

这个节点给删除。

删除了那个节点之后，zk会负责通知监听这个节点的监听器，也就是客户端B之前加的那个监听器，说：兄弟，你监听的那个节点被删除了，有人释放了锁。

![](https://user-gold-cdn.xitu.io/2018/11/30/1676531dd1f142d2?imageView2/0/w/1280/h/960/ignore-error/1)

此时客户端B的监听器感知到了上一个顺序节点被删除，也就是排在他之前的某个客户端释放了锁。

此时，就会通知客户端B重新尝试去获取锁，也就是获取" **my_lock** "节点下的子节点集合，此时为：

![](https://user-gold-cdn.xitu.io/2018/11/30/16765349dbf13775?imageView2/0/w/1280/h/960/ignore-error/1)

集合里此时只有客户端B创建的唯一的一个顺序节点了！

然后呢，客户端B判断自己居然是集合中的第一个顺序节点，bingo！可以加锁了！ **直接完成加锁** ，运行后续的业务代码即可，运行完了之后再次释放锁。

![](https://user-gold-cdn.xitu.io/2018/11/30/1676531f71973f37?imageView2/0/w/1280/h/960/ignore-error/1)

# 三、总结 #

其实如果有客户端C、客户端D等N个客户端争抢一个zk分布式锁，原理都是类似的。

* 大家都是上来直接创建一个锁节点下的一个接一个的临时顺序节点
* 如果自己不是第一个节点，就对自己上一个节点加监听器
* 只要上一个节点释放锁，自己就排到前面去了，相当于是一个排队机制。

而且用临时顺序节点的另外一个用意就是，如果某个客户端创建临时顺序节点之后，不小心自己宕机了也没关系，zk感知到那个客户端宕机，会自动删除对应的临时顺序节点，相当于自动释放锁，或者是自动取消自己的排队。

最后，咱们来看下用Curator框架进行加锁和释放锁的一个过程：

![](https://user-gold-cdn.xitu.io/2018/11/30/167653218b18739b?imageView2/0/w/1280/h/960/ignore-error/1)

其实用开源框架就是这点好，方便。这个Curator框架的zk分布式锁的加锁和释放锁的实现原理，其实就是上面我们说的那样子。

但是如果你要手动实现一套那个代码的话。还是有点麻烦的，要考虑到各种细节，异常处理等等。所以大家如果考虑用zk分布式锁，可以参考下本文的思路。

**END**

**如有收获，请帮忙转发，您的鼓励是作者最大的动力，谢谢！**

**一大波微服务、分布式、高并发、高可用的原创系列文章正在路上**

****欢迎扫描下方二维码** ，持续关注：**

![](https://user-gold-cdn.xitu.io/2018/11/14/167127ebdf8586b2?imageView2/0/w/1280/h/960/ignore-error/1)

**石杉的架构笔记（id:shishan100）**

**十余年BAT架构经验倾囊相授**

**
**

**
> 
> **推荐阅读：**
> 
> 1、 拜托！面试请不要再问我Spring Cloud底层原理
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5be13b83f265da6116393fc7
> )
> 
> 
> 
> 2、 【双11狂欢的背后】微服务注册中心如何承载大型系统的千万级访问？
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5be3f8dcf265da613a5382ca
> )
> 
> 
> 
> 3、 【性能优化之道】每秒上万并发下的Spring Cloud参数优化实战
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5be83e166fb9a049a7115580
> )
> 
> 
> 
> 4、 微服务架构如何保障双11狂欢下的99.99%高可用
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5be99a68e51d4511a8090440
> )
> 
> 
> 
> 5、 兄弟，用大白话告诉你小白都能听懂的Hadoop架构原理
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5beaf02ce51d457e90196069
> )
> 
> 
> 
> 6、 大规模集群下Hadoop NameNode如何承载每秒上千次的高并发访问
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5bec278c5188253e64332c76
> )
> 
> 
> 
> 7、【 性能优化的秘密】Hadoop如何将TB级大文件的上传性能优化上百倍
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Flink.juejin.im%25252525252F%25252525253Ftarget%25252525253Dhttps%2525252525253A%2525252525252F%2525252525252Fjuejin.im%2525252525252Fpost%2525252525252F5bed82a9e51d450f9461cfc7
> )
> 
> 
> 
> 8、 [拜托，面试请不要再问我TCC分布式事务的实现原理](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Flink.juejin.im%252525252F%252525253Ftarget%252525253Dhttps%25252525253A%25252525252F%25252525252Fjuejin.im%25252525252Fpost%25252525252F5bf201f7f265da610f63528a
> ) 坑爹呀！
> (
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Fjuejin.im%252525252Fpost%252525252F5bf2c6b6e51d456693549af4
> )
> 
> 
> 
> 9、 [【坑爹呀！】最终一致性分布式事务如何保障实际生产中99.99%高可用？](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Fjuejin.im%252525252Fpost%252525252F5bf2c6b6e51d456693549af4
> )
> 
> 
> 
> 10、 [拜托，面试请不要再问我Redis分布式锁的实现原理！](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Fjuejin.im%252525252Fpost%252525252F5bf3f15851882526a643e207
> )
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> **11、** **[【眼前一亮！】看Hadoop底层算法如何优雅的将大规模集群性能提升10倍以上？](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Flink.juejin.im%2525253Ftarget%2525253Dhttps%252525253A%252525252F%252525252Fjuejin.im%252525252Fpost%252525252F5bf5396f51882509a768067e
> )**
> 
> 
> 
> ****
> 
> 
> 
> **12、** **[亿级流量系统架构之如何支撑百亿级数据的存储与计算](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Flink.juejin.im%25252F%25253Ftarget%25253Dhttps%2525253A%2525252F%2525252Fjuejin.im%2525252Fpost%2525252F5bfab59fe51d4551584c7bcf
> )**
> 
> 
> 
> 13、 [亿级流量系统架构之如何设计高容错分布式计算系统](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Flink.juejin.im%253Ftarget%253Dhttps%25253A%25252F%25252Fjuejin.im%25252Fpost%25252F5bfbeeb9f265da61407e9679
> )
> 
> 
> 
> 14、 [亿级流量系统架构之如何设计承载百亿流量的高性能架构](
> https://link.juejin.im/?target=https%3A%2F%2Flink.juejin.im%3Ftarget%3Dhttps%253A%252F%252Fjuejin.im%252Fpost%252F5bfd2df1e51d4574b133dd3a
> )
> 
> 
> 
> 15、 [亿级流量系统架构之如何设计每秒十万查询的高并发架构](
> https://link.juejin.im/?target=https%3A%2F%2Fjuejin.im%2Fpost%2F5bfe771251882509a7681b3a
> )
> 
> 
> 
> 16、 [亿级流量系统架构之如何设计全链路99.99%高可用架构](
> https://juejin.im/post/5bffab686fb9a04a102f0022 )
> 
> 
> 

**