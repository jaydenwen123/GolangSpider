# 非常硬核的技术知识-CopyOnWrite思想 #

> 
> 
> 
> “
> 今天聊一个非常硬核的技术知识，给大家分析一下CopyOnWrite思想是什么，以及在Java并发包中的具体体现，包括在Kafka内核源码中是如何运用这个思想来优化并发性能的。
> 
> 
> 

> 
> 
> 
> 这个CopyOnWrite在面试的时候，很可能成为面试官的一个杀手锏把候选人给一击必杀，也很有可能成为候选人拿下Offer的独门秘籍，是相对高级的一个知识。
> 
> 
> 

## 1、读多写少的场景下引发的问题？ ##

大家可以设想一下现在我们的内存里有一个ArrayList，这个ArrayList默认情况下肯定是线程不安全的，要是多个线程并发读和写这个ArrayList可能会有问题。

好，问题来了，我们应该怎么让这个ArrayList变成线程安全的呢？

有一个非常简单的办法，对这个ArrayList的访问都加上线程同步的控制。

比如说一定要在synchronized代码段来对这个ArrayList进行访问，这样的话，就能同一时间就让一个线程来操作它了，或者是用ReadWriteLock读写锁的方式来控制，都可以。

我们假设就是用ReadWriteLock读写锁的方式来控制对这个ArrayList的访问。

这样多个读请求可以同时执行从ArrayList里读取数据，但是读请求和写请求之间互斥，写请求和写请求也是互斥的。

大家看看，代码大概就是类似下面这样：

` public Object read () { lock.readLock().lock(); // 对ArrayList读取 lock.readLock().unlock(); } public void write () { lock.writeLock().lock(); // 对ArrayList写 lock.writeLock().unlock(); } 复制代码`

大家想想，类似上面的代码有什么问题呢？

最大的问题，其实就在于写锁和读锁的互斥。假设写操作频率很低，读操作频率很高，是写少读多的场景。

那么偶尔执行一个写操作的时候，是不是会加上写锁，此时大量的读操作过来是不是就会被阻塞住，无法执行？

这个就是读写锁可能遇到的最大的问题。

## 2、引入 CopyOnWrite 思想解决问题 ##

这个时候就要引入CopyOnWrite思想来解决问题了。

他的思想就是，不用加什么读写锁，锁统统给我去掉，有锁就有问题，有锁就有互斥，有锁就可能导致性能低下，你阻塞我的请求，导致我的请求都卡着不能执行。

那么他怎么保证多线程并发的安全性呢？

很简单，顾名思义，利用“CopyOnWrite”的方式，这个英语翻译成中文，大概就是**“写数据的时候利用拷贝的副本来执行”。**

你在读数据的时候，其实不加锁也没关系，大家左右都是一个读罢了，互相没影响。

问题主要是在写的时候，写的时候你既然不能加锁了，那么就得采用一个策略。

假如说你的ArrayList底层是一个数组来存放你的列表数据，那么这时比如你要修改这个数组里的数据，你就必须先拷贝这个数组的一个副本。

然后你可以在这个数组的副本里写入你要修改的数据，但是在这个过程中实际上你都是在操作一个副本而已。

这样的话，读操作是不是可以同时正常的执行？这个写操作对读操作是没有任何的影响的吧！

大家看下面的图，一起来体会一下这个过程：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b217e4484cc9b4?imageView2/0/w/1280/h/960/ignore-error/1) 关键问题来了，那那个写线程现在把副本数组给修改完了，现在怎么才能让读线程感知到这个变化呢？

**关键点来了，划重点！这里要配合上volatile关键字的使用。**

笔者之前写过文章，给大家解释过volatile关键字的使用，核心就是让一个变量被写线程给修改之后，立马让其他线程可以读到这个变量引用的最近的值，这就是volatile最核心的作用。

所以一旦写线程搞定了副本数组的修改之后，那么就可以用volatile写的方式，把这个副本数组赋值给volatile修饰的那个数组的引用变量了。

只要一赋值给那个volatile修饰的变量，立马就会对读线程可见，大家都能看到最新的数组了。

下面是JDK里的 CopyOnWriteArrayList 的源码。

大家看看写数据的时候，他是怎么拷贝一个数组副本，然后修改副本，接着通过volatile变量赋值的方式，把修改好的数组副本给更新回去，立马让其他线程可见的。

` // 这个数组是核心的，因为用volatile修饰了 // 只要把最新的数组对他赋值，其他线程立马可以看到最新的数组 private transient volatile Object[] array; public boolean add (E e) { final ReentrantLock lock = this.lock; lock.lock(); try { Object[] elements = getArray(); int len = elements.length; // 对数组拷贝一个副本出来 Object[] newElements = Arrays.copyOf(elements, len + 1 ); // 对副本数组进行修改，比如在里面加入一个元素 newElements[len] = e; // 然后把副本数组赋值给volatile修饰的变量 setArray(newElements); return true ; } finally { lock.unlock(); } } 复制代码`

然后大家想，因为是通过副本来进行更新的，万一要是多个线程都要同时更新呢？那搞出来多个副本会不会有问题？

当然不能多个线程同时更新了，这个时候就是看上面源码里，加入了lock锁的机制，也就是同一时间只有一个线程可以更新。

那么更新的时候，会对读操作有任何的影响吗？

绝对不会，因为读操作就是非常简单的对那个数组进行读而已，不涉及任何的锁。而且只要他更新完毕对volatile修饰的变量赋值，那么读线程立马可以看到最新修改后的数组，这是volatile保证的。

这样就完美解决了我们之前说的读多写少的问题。

如果用读写锁互斥的话，会导致写锁阻塞大量读操作，影响并发性能。

但是如果用了CopyOnWriteArrayList，就是用空间换时间，更新的时候基于副本更新，避免锁，然后最后用volatile变量来赋值保证可见性，更新的时候对读线程没有任何的影响！

## 3、CopyOnWrite 思想在Kafka源码中的运用 ##

在Kafka的内核源码中，有这么一个场景，客户端在向Kafka写数据的时候，会把消息先写入客户端本地的内存缓冲，然后在内存缓冲里形成一个Batch之后再一次性发送到Kafka服务器上去，这样有助于提升吞吐量。

话不多说，大家看下图：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b217eb1ad2ca13?imageView2/0/w/1280/h/960/ignore-error/1) 这个时候Kafka的内存缓冲用的是什么数据结构呢？大家看源码：

` private final ConcurrentMap<topicpartition, deque<= "" span= "" > batches = new CopyOnWriteMap<TopicPartition, Deque>(); 复制代码`

这个数据结构就是核心的用来存放写入内存缓冲中的消息的数据结构，要看懂这个数据结构需要对很多Kafka内核源码里的概念进行解释，这里先不展开。

但是大家关注一点，他是自己实现了一个CopyOnWriteMap，这个CopyOnWriteMap采用的就是CopyOnWrite思想。

我们来看一下这个CopyOnWriteMap的源码实现：

` // 典型的volatile修饰普通Map private volatile Mapmap; @Override public synchronized V put (K k, V v) { // 更新的时候先创建副本，更新副本，然后对volatile变量赋值写回去 Mapcopy= new HashMap( this.map); V prev = copy.put(k, v); this.map = Collections.unmodifiableMap(copy); return prev; } @Override public V get (Object k) { // 读取的时候直接读volatile变量引用的map数据结构，无需锁 return map.get(k); } 复制代码`

所以Kafka这个核心数据结构在这里之所以采用CopyOnWriteMap思想来实现，就是因为这个Map的key-value对，其实没那么频繁更新。

也就是TopicPartition-Deque这个key-value对，更新频率很低。

但是他的get操作却是高频的读取请求，因为会高频的读取出来一个TopicPartition对应的Deque数据结构，来对这个队列进行入队出队等操作，所以对于这个map而言，高频的是其get操作。

这个时候，Kafka就采用了CopyOnWrite思想来实现这个Map，避免更新key-value的时候阻塞住高频的读操作，实现无锁的效果，优化线程并发的性能。

相信大家看完这个文章，对于CopyOnWrite思想以及适用场景，包括JDK中的实现，以及在Kafka源码中的运用，都有了一个切身的体会了。

如果你能在面试时说清楚这个思想以及他在JDK中的体现，并且还能结合知名的开源项目 Kafka 的底层源码进一步向面试官进行阐述，面试官对你的印象肯定大大的加分。

> 
> 
> 
> 原文自：石杉的架构笔记微信公众号
> 
>