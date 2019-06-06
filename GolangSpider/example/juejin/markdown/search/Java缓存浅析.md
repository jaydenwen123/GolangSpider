# Java缓存浅析 #

拿破仑说：胜利属于坚持到最后的人。

而正巧，咱们今天就是要聊一个，关于怎么让系统在狂轰乱炸甚至泰山压顶的情况下，都屹立不倒并坚持到最后的话题——缓存。

![拿破仑](https://ipictures.github.io/category/people/%E6%8B%BF%E7%A0%B4%E4%BB%91.gif)

> 
> 
> 
> Victory belongs to the most persevering. — Napoleon Bonaparte, French
> military and political leader
> 
> 

## **目录体系** ##

下面我们先简单浏览一下这个分享的目录体系。

今天我会分五个方面给大家介绍关于缓存使用的问题，包括原理、实践、技术选型和常见问题。

这个目录体系就是一副人体骨骼，只有把各种内脏、器官和血肉都填充进去，缓存之美才能跃然纸上。接下来，我就邀请大家跟我一起来做这件事情.

让我们不止步于Hello World，一起来聊聊缓存。

![聊聊缓存-目录体系](https://user-gold-cdn.xitu.io/2019/3/10/169659b154bcce1e?imageView2/0/w/1280/h/960/ignore-error/1)

## 关于缓存 ##

### What ###

缓存是什么？

缓存是实际工作中非常常用的一种提高性能的方法。

而在java中，所谓缓存，就是将程序或系统经常要调用的对象存在内存中，再次调用时可以快速从内存中获取对象，不必再去创建新的重复的实例。

这样做可以减少系统开销，提高系统效率。

目前缓存的做法分为两种模式:

* 

内存缓存：缓存数据存放在服务器的内存空间中。

` 优点：速度快。 缺点：资源有限。 复制代码`
* 

文件缓存：缓存数据存放在服务器的硬盘空间中。

` 优点：容量大。 缺点：速度偏慢，尤其在缓存数量巨大时。 复制代码`

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b154ae7f1e?imageView2/0/w/1280/h/960/ignore-error/1)

### why ###

为什么要使用缓存？

对于为什么要使用缓存，我见过的最精炼的回答是：来源一个梦想，那就是多快好省的构建社会主义社会。

但这是一种很矛盾的说法，就好像你不是高富帅还想迎娶白富美，好像是痴人说梦啊。

因为多就不可能快，好就不能省，怎么做到多又快，好而且省呢？

答案就是用缓存！

下面我们就聊聊怎么用缓存实现这个梦想。

首先我想先声明一下，我什么会想到做这样一个分享。

其实，从第一次使用 Java整型的缓存，到了解CDN的代理缓存，从初次接触 MySQL内置的查询缓存，到使用 Redis缓存Session，我越来越发现使用缓存的重要性和普遍性。

因此我觉得自己有必要把自己的所学所用梳理出来，用于工作，并造福大家，因此才有了这样一个技术分享。

聊缓存之前我们先聊聊数据库。

在增删改查中，数据库查询占据了数据库操作的80%以上， 非常频繁的磁盘I/O读取操作，会导致数据库性能极度低下。

而数据库的重要性就不言而喻了：

* 数据库通常是企业应用系统最核心的部分
* 数据库保存的数据量通常非常庞大
* 数据库查询操作通常很频繁，有时还很复杂

我们知道，对于多数Web应用，整个系统的瓶颈在于数据库。

原因很简单，Web应用中的其他因素，例如网络带宽、负载均衡节点、应用服务器（包括CPU、内存、硬盘灯、连接数等）、缓存，都很容易通过水平的扩展（俗称加机器）来实现性能的提高。

而对于MySQL，由于数据一致性的要求，无法通过简单的增加机器来分散向数据库 写数据 带来的压力。虽然可以通过前置缓存（Redis等）、读写分离、分库分表来减轻压力，但是与系统其它组件的水平扩展相比，受到了太多的限制，而切会大大增加系统的复杂性。

因此数据库的连接和读写要十分珍惜。

可能你会想到那就直接用缓存呗，但大量的用、不分场景的用缓存显然是不科学的。我们不能手里有了一把锤子，看什么都是钉子。

但缓存也不是万能的，要慎用缓存，想要用好缓存并不容易。因此我花了点时间整理了一下关于缓存的实现以及常见的一些问题。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1569e8aa7?imageView2/0/w/1280/h/960/ignore-error/1)

### when ###

首先简单梳理一下Web请求的过程，以及不同节点缓存的作用。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1567f9f51?imageView2/0/w/1280/h/960/ignore-error/1)

### how ###

先不讲代码，对于缓存是如何工作的，简单的缓存数据请求流程就如下图。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b156b482f5?imageView2/0/w/1280/h/960/ignore-error/1)

设计缓存的时候需要考虑的最关键的两个缓存策略。

- TTL（Time To Live ） 存活期， 即从缓存中创建时间点开始直到它到期的一个时间段（不管在这个时间段内有没有访问都将过期）

* TTI（Time To Idle） 空闲期， 即一个数据多久没被访问将从缓存中移除的时间

后面讲到缓存雪崩的时候，会讲到，如果缓存策略设置不当，将会造成如何的灾难性后果，以及如何避免，这里先按下不表。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1826b9aff?imageView2/0/w/1280/h/960/ignore-error/1)

## 自定义缓存 ##

### 如何实现 ###

前面介绍了关于缓存的一些概念，那么实现缓存，或者确切的说实现存储的前置缓存很难吗？

答案是：不难。

JVM本身就是一个高速的缓存存储场所，同时Java为我们提供了线程安全的ConcurrentMap，可以非常方便的实现一个完全由你自定义的缓存实例。

后面你会发现，Spring Cache的缺省实现SimpleCacheManager，也是这样设计自己的缓存的。

这里放上简单的实现代码，不过36行，就实现了对缓存的存储、更新、读取和删除等基本操作。 再结合实际的业务代码，就能不依赖任何三方的实现，在JVM中轻松玩转缓存了。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b182c93c4c?imageView2/0/w/1280/h/960/ignore-error/1)

但是，我想作为有追求的技术人，各位是绝对不会止步于此的。

那么我们思考一下，我们自定义的缓存实现，有哪些优缺点呢？

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b186e44b5b?imageView2/0/w/1280/h/960/ignore-error/1)

同与自定义的缓存相比，就能更深刻的理解Spring Cache的原理，以及优点。

这里先把Spring Cache的特性列举出来，下面还会介绍它的原理和具体用法。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b187317c7f?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b18a14c179?imageView2/0/w/1280/h/960/ignore-error/1)

## Spring Cache ##

Spring Cache是Spring提供的对缓存功能的抽象：即允许绑定不同的缓存解决方案（如Ehcache、Redis、Memcache、Map等等），但本身不直接提供缓存功能的实现。

它支持注解方式使用缓存，非常方便。

Spring Cache的实现本质上依赖了Spring AOP对切面的支持。

知道了Spring Cache的原理，你会对Spring Cache的注解的使用有更深入的认识。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1a957dce3?imageView2/0/w/1280/h/960/ignore-error/1)

Spring Cache主要用到的注解有4个。

@CacheEvict对于保证缓存一致性非常重要，后面会专门讲一下这个问题。

同时，Spring还支持自定义的缓存Key以及SpringEL，这里不详细讲了，感兴趣的同学可以参考Spring Cache的文档。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1ad7e9b45?imageView2/0/w/1280/h/960/ignore-error/1)

## 缓存三高音 ##

正如写得再好的乐谱，都需要歌唱家演唱出来才能美妙动听一样。

上面讲到Spring Cache是对缓存的抽象，那么常用的缓存的实现有哪些呢？

歌唱界有世界三大男高音，那么缓存界如果来评选一下话，三大高音会是谁呢？

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1ac5b11ba?imageView2/0/w/1280/h/960/ignore-error/1)

### Redis ###

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1b1518722?imageView2/0/w/1280/h/960/ignore-error/1)

redis是一个key-value存储系统，这点和Memcached类似。

不同的是它支持存储的value类型相对更多，包括string(字符串)、list(链表)、set(集合)、zset(sorted set --有序集合)和hash（哈希类型）。这些数据类型都支持push/pop、add/remove及取交集并集和差集。

和Memcached一样，为了保证效率，数据都是缓存在内存中。

区别的是redis会周期性的把更新的数据写入磁盘或者把修改操作写入追加的记录文件，并且在此基础上实现了master-slave(主从)同步。 Redis支持主从同步。数据可以从主服务器向任意数量的从服务器上同步，从服务器可以是关联其他从服务器的主服务器。这使得Redis可执行单层树复制。

存盘可以有意无意的对数据进行写操作。由于完全实现了发布/订阅机制，使得从数据库在任何地方同步树时，可订阅一个频道并接收主服务器完整的消息发布记录。

同步对读取操作的可扩展性和数据冗余很有帮助。

Redis有哪些适合的场景？

* 会话缓存（Session Cache）:用Redis缓存会话比其他存储（如memcached）的优势在于，redis提供持久化。
* 全页缓存（FPC）:除基本的会话token之外，Redis还提供很简便的FPC平台。
* 队列：Redis在内存存储引擎领域的一大优点是提供list和set操作，这使得Redis能作为一个很好的消息队列平台来使用。
* 排行榜/计数器：Redis在内存中对数据进行递增递减的操作实现的非常好。
* 订阅/发布

缺点：

* 

持久化。Redis直接将数据存储到内存中，要将数据保存到磁盘上，Redis可以使用两种方式实现持久化过程。

定时快照（snapshot）：每隔一段时间将整个数据库写到磁盘上，每次均是写全部数据，代价非常高。 基于语句追加（aof）：只追踪变化的数据，但是追加的log可能过大，同时所有的操作均重新执行一遍，回复速度慢。

* 

耗内存，占用内存过高。

### Ehcache ###

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1c67de7db?imageView2/0/w/1280/h/960/ignore-error/1)

Ehcache 是一个成熟的缓存框架，你可以直接使用它来管理你的缓存。

Java缓存框架 EhCache EhCache 是一个纯Java的进程内缓存框架，具有快速、精干等特点，是Hibernate中默认的CacheProvider。

特性：可以配置内存不足时，启用磁盘缓存(maxEntriesLoverflowToDiskocalDisk配置当内存中对象数量达到maxElementsInMemory时，Ehcache将会对象写到磁盘中)。

### Memcached ###

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1d144229c?imageView2/0/w/1280/h/960/ignore-error/1)

Memcached 是一个高性能的分布式内存对象缓存系统，用于动态Web应用以减轻数据库负载。它基于一个存储键/值对的hashmap。

其守护进程（daemon ）是用C写的，但是客户端可以用任何语言来编写，并通过memcached协议与守护进程通信。

Memcached通过在内存中缓存数据和对象来减少读取数据库的次数，从而提高动态、数据库驱动网站的速度。

同属于个key-value存储系统，Memcached与Redis常常一起比:

* Memcached的数据结构和操作较为简单，不如Redis支持的结构丰富。
* 使用简单的key-value存储的话，Memcached的内存利用率更高， 而如果Redis采用hash结构来做key-value存储，由于其组合式的压缩，其内存利用率会高于Memcached。
* 由于Redis只使用单核，而Memcached可以使用多核，所以平均每一个核上Redis在存储小数据时比Memcached性能更高。 而在100k以上的数据中，Memcached性能要高于Redis，虽然Redis最近也在存储大数据的性能上进行优化，但是比起Memcached，还是稍有逊色。
* Redis虽然是基于内存的存储系统，但是它本身是支持内存数据的持久化的，而且提供两种主要的持久化策略：RDB快照和AOF日志。而memcached是不支持数据持久化操作的。 Memcached是全内存的数据缓冲系统，Redis虽然支持数据的持久化，但是全内存毕竟才是其高性能的本质。
* 作为基于内存的存储系统来说，机器物理内存的大小就是系统能够容纳的最大数据量。如果需要处理的数据量超过了单台机器的物理内存大小，就需要构建分布式集群来扩展存储能力。

Memcached本身并不支持分布式，因此只能在客户端通过像一致性哈希这样的分布式算法来实现Memcached的分布式存储。

相较于Memcached只能采用客户端实现分布式存储，Redis更偏向于在服务器端构建分布式存储。最新版本的Redis已经支持了分布式存储功能。

![缓存三高音比较](https://user-gold-cdn.xitu.io/2019/3/10/169659b1d1aa65b1?imageView2/0/w/1280/h/960/ignore-error/1) 缓存三高音比较

## 缓存进阶 ##

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1d6a2bf5a?imageView2/0/w/1280/h/960/ignore-error/1)

缓存由于其高并发和高性能的特性，已经在项目中被广泛使用。尤其是在高并发、分布式和微服务的业务场景和架构下。

无论是高并发、分布式还是微服务都依赖于高性能的服务器。而谈到高性能服务器，就必谈缓存。

所谓高性能主要体现在高可用情况下，业务处理时间短，数据正确。

数据处理及时就是个“空间换时间”的问题，利用分布式内存或者闪存等可以快速存取的设备，来替代部署在一般服务器上的数据库，机械硬盘上存储的文件，这是缓存提升服务器性能的本质。

高并发（High Concurrency）： 是互联网分布式系统架构设计中必须考虑的因素之一，它通常是指，通过设计保证系统能够同时并行处理很多请求。

分布式： 是以缩短单个任务的执行时间来提升效率的。 比如一个任务由10个子任务组成，每个子任务单独执行需1小时，则在一台服务器上执行改任务需10小时。 采用分布式方案，提供10台服务器，每台服务器只负责处理一个子任务，不考虑子任务间的依赖关系，执行完这个任务只需一个小时。

微服务： 架构强调的第一个重点就是业务系统需要彻底的组件化和服务化，原有的单个业务系统会拆分为多个可以独立开发，设计，运行和运维的小应用。这些小应用之间通过服务完成交互和集成。

### 缓存一致性问题 ###

缓存一致性是如何发生的：先写数据库，再淘汰缓存:

` 第一步写数据库成功，第二步淘汰缓存失败，则会引发一次严重的缓存不一致问题。 复制代码`

如何避免缓存不一致的问题：先淘汰缓存，再写数据库：

` 第一步淘汰缓存成功，第二步写数据库失败，则只会引发一次Cache miss。 复制代码`

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1e98a3e51?imageView2/0/w/1280/h/960/ignore-error/1)

#### 分布式缓存一致性 ####

我们使用zookeeper来协调各个缓存实例节点，zookeeper是一个分布式协调服务，包含一个原语集，可以通知所有watch节点的client端，并保证事件发生顺序和client收到消息的顺序一致；使用zookeeper集群可非常容易的实现这场景。

一致性Hash算法通过一个叫做一致性Hash环的数据结构，实现KEY到缓存服务器的Hash映射。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1eb110ade?imageView2/0/w/1280/h/960/ignore-error/1)

### 缓存雪崩 ###

产生原因1. a. 由于Cache层承载着大量请求，有效的保护了Storage层(通常认为此层抗压能力稍弱)，所以Storage的调用量实际很低，所以它很爽。 b. 但是，如果Cache层由于某些原因(宕机、cache服务挂了或者不响应了)整体crash掉了，也就意味着所有的请求都会达到Storage层，所有Storage的调用量会暴增，所以它有点扛不住了，甚至也会挂掉

产生原因2. 我们设置缓存时采用了相同的过期时间，导致缓存在某一时刻同时失效，请求全部转发到DB，DB瞬时压力过重雪崩。

雪崩问题在国外叫做：stampeding herd(奔逃的野牛)，指的的cache crash后，流量会像奔逃的野牛一样，打向后端。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1f565d313?imageView2/0/w/1280/h/960/ignore-error/1)

解决方案

* 加锁/队列 保证缓存单线程的写

失效时的雪崩效应对底层系统的冲击非常可怕。

大多数系统设计者考虑用加锁或者队列的方式保证缓存的单线 程（进程）写，从而避免失效时大量的并发请求落到底层存储系统上。

加锁排队只是为了减轻数据库的压力，并没有提高系统吞吐量。

假设在高并发下，缓存重建期间key是锁着的，这是过来1000个请求999个都在阻塞的。同样会导致用户等待超时，这是个治标不治本的方法！

加锁排队的解决方式分布式环境的并发问题，有可能还要解决分布式锁的问题；线程还会被阻塞，用户体验很差！因此，在真正的高并发场景下很少使用！

* 避免缓存同时失效

将缓存失效时间分散开，比如我们可以在原有的失效时间基础上，末尾增加一个随机值。

* 缓存降级

当访问量剧增、服务出现问题（如响应时间慢或不响应）或非核心服务影响到核心流程的性能时，仍然需要保证服务还是可用的，即使是有损服务。

系统可以根据一些关键数据进行自动降级，也可以配置开关实现人工降级。

降级的最终目的是保证核心服务可用，即使是有损的。而且有些服务是无法降级的（如加入购物车、结算）。

在进行降级之前要对系统进行梳理，看看系统是不是可以丢卒保帅；从而梳理出哪些必须誓死保护，哪些可降级。

比如可以参考日志级别设置预案：

（1）一般：比如有些服务偶尔因为网络抖动或者服务正在上线而超时，可以自动降级；

（2）警告：有些服务在一段时间内成功率有波动（如在95~100%之间），可以自动降级或人工降级，并发送告警；

（3）错误：比如可用率低于90%，或者数据库连接池被打爆了，或者访问量突然猛增到系统能承受的最大阀值，此时可以根据情况自动降级或者人工降级；

（4）严重错误：比如因为特殊原因数据错误了，此时需要紧急人工降级。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1fdf4d51d?imageView2/0/w/1280/h/960/ignore-error/1)

### 缓存击穿/缓存穿透 ###

缓存穿透是指查询一个一定不存在的数据，由于缓存是不命中时被动写的，并且出于容错考虑，如果从存储层查不到数据则不写入缓存，这将导致这个不存在的数据每次请求都要到存储层去查询，失去了缓存的意义。在流量大时，可能DB就挂掉了，要是有人利用不存在的key频繁攻击我们的应用，这就是漏洞。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b1f8f3268c?imageView2/0/w/1280/h/960/ignore-error/1)

缓存穿透-解决方案1

一个简单粗暴的方法，如果一个查询返回的数据为空（不管是数 据不存在，还是系统故障），我们仍然把这个空结果进行缓存，

但它的过期时间会很短，最长不超过五分钟。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b212616f2f?imageView2/0/w/1280/h/960/ignore-error/1)

缓存穿透-解决方案2

最常见的则是采用布隆过滤器，将所有可能存在的数据哈希到一个足够大的bitmap中，一个一定不存在的数据会被 这个bitmap拦截掉，从而避免了对底层存储系统的查询压力。

例如，商城有100万用户数据，将所有用户id刷入一个Map。

当请求过来以后，先判断Map中是否包含该用户id，不包含直接返回，包含的话先去缓存中查是否有这条数据，有的话返回，没有的话再去查数据库。

这样不仅减轻了数据库的压力，缓存系统的压力也将大大降低。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b2168ead90?imageView2/0/w/1280/h/960/ignore-error/1)

## 寄语 ##

古人云：纸上得来终觉浅，绝知此事要躬行。

别人的经验和智慧，需要经过你亲自验证才知道是不是真理，要经过亲手实践才能为我所用。

别人的知识只是一些树枝，需要你把它们编织成一架梯子，才能助你高升。

![](https://user-gold-cdn.xitu.io/2019/3/10/169659b218403a98?imageView2/0/w/1280/h/960/ignore-error/1)

## 参考链接 ##

* [百度百科 - 缓存]( https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2F%25E7%25BC%2593%25E5%25AD%2598 )
* [Spring思维导图，让Spring不再难懂（cache篇）]( https://link.juejin.im?target=http%3A%2F%2Fwww.iteye.com%2Fnews%2F32626 )
* [importnew : Spring Cache]( https://link.juejin.im?target=http%3A%2F%2Fwww.importnew.com%2F22757.html )
* [图解分布式架构的演进]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FnJhY4ug3iOfISfiq2K-KtA )
* [EHCACHE]( https://link.juejin.im?target=http%3A%2F%2Fwww.ehcache.org%2F )
* [ehcache官方文档]( https://link.juejin.im?target=http%3A%2F%2Fwww.ehcache.org%2Fdocumentation%2F )
* [ehcache入门基础示例]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fvbirdbest%2Farticle%2Fdetails%2F72763048 )
* [ehcache详细解读]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu014209975%2Farticle%2Fdetails%2F53320395 )
* [ehcache memcache redis 三大缓存男高音]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fjationxiaozi%2Farticle%2Fdetails%2F8509732 )
* [网站缓存技术 ehcache memcache redis 的比较]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fzq602316498%2Farticle%2Fdetails%2F49449083 )
* [缓存击穿、失效及热点key问题]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FRcWgyE28-3ao3M2r137HhQ )
* [Cache 应用中的服务过载案例研究]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FzmT2QABxENvXDpXsz-CJog )
* [Bloom Filter布隆过滤器]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FuQ0tQgpD_wwaeNg-XPh5XA )
* [缓存在高并发场景下的常见问题]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FZayqRozgXY-P5CiZU5nK7A )
* [缓存雪崩问题]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F8WIscH3SZDaABFHoI_EFYQ )
* [再聊缓存技术]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FMUtjej6FvX50JZ6qoGyoxA )
* [缓存穿透问题]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FcgnREJoMvM2dVLmmcDL_Bw )
* [微服务化之缓存的设计]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FwrUqF0bxDWnjP-oAvUTY4g )
* [缓存与数据库一致性保证]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2Fz364m8iWgk2tSag_q7VUiQ )
* [分布式之缓存击穿]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FYj38EiuWkURB4APMfrZu7Q )
* [CDN缓存小结]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fpedrojuliet%2Farticle%2Fdetails%2F78394732 )
* [ava 中整型的缓存机制]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F0xGA_b2blLjAAwxqj5eN6A )
* [mysql的查询缓存]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fking_kgh%2Farticle%2Fdetails%2F74855217 )
* [使用Spring Session和Redis解决分布式Session跨域共享问题]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fxlgen157387%2Farticle%2Fdetails%2F57406162 )
* [学习Spring-Session+Redis实现session共享]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fandyfengzp%2Fp%2F6434287.html )
* [详解 MySQL 基准测试和 sysbench 工具]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FsIXK7RKCsF5IJ6IcuQmuhg )
* [分布式之缓存击穿]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FD46nyEwjpx7T_5TtpdPSzQ )
* [分布式之数据库和缓存双写一致性方案解析]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FvOjDG8aH45o5vVXXm0v78g )