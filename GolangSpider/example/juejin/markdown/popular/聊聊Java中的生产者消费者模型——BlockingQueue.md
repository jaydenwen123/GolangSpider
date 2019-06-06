# 聊聊Java中的生产者消费者模型——BlockingQueue #

## 前言 ##

**生产者/消费者模型** 相信各位都不陌生，是一种很常见的分布式资源调度模型。在这个模型中，至少有两个对象：生产者和消费者。生产者只负责创建资源，消费者只负责使用资源。如果自己实现一个简单的生产者/消费者模型也很容易，无非就是通过一个队列来做，但是这种方式有很多隐藏的缺陷：

* 需要保证资源的线程可见性，同时要手动实现线程同步
* 需要考虑各种临界情况和拒绝策略
* 需要在吞吐量与线程安全之间保持平衡

所以Java已经提前为我们封装好了接口和实现，接下来我们就要针对BlockingQueue接口和它的常用的实现类LinkedBlockingQueue进行简要的分析

## 阻塞队列 ##

### 概念 ###

BlockingQueue，含义为阻塞队列，我们可以从类定义看出，其继承了Queue接口，所以可以当作队列来使用：

![图1](https://user-gold-cdn.xitu.io/2019/6/5/16b255543bb0498b?imageView2/0/w/1280/h/960/ignore-error/1)

既然叫做阻塞队列，也就是说这个队列的操作是以阻塞方式进行的，体现在如下两个方面：

* 插入元素的操作是阻塞的：当队列满时，执行插入操作的线程被阻塞
* 移除元素的操作时阻塞的：当队列空时，执行移除操作的线程被阻塞

通过这种方式，可以方便地协调生产者和消费者之间的关系

### 接口方法 ###

在BlockingQueue中，定义了以下6个接口：

` public interface BlockingQueue < E > extends Queue < E > { boolean add (E e) ; boolean offer (E e) ; void put (E e) throws InterruptedException ; boolean offer (E e, long timeout, TimeUnit unit) throws InterruptedException ; E take () throws InterruptedException ; E poll ( long timeout, TimeUnit unit) throws InterruptedException ; int remainingCapacity () ; boolean remove (Object o) ; public boolean contains (Object o) ; int drainTo (Collection<? super E> c) ; int drainTo (Collection<? super E> c, int maxElements) ; } 复制代码`

这些接口方法按功能可以分为三类：

* 添加元素：包括add、offer、put
* 移除元素：包括remove、poll、take、drainTo
* 获取/检查元素：包括contains、remainingCapacity

一般地，我们也将添加元素叫做 ` put` 操作（即使使用的是 ` offer` 方法而不是 ` put` 方法），移除元素的叫做 ` take` 操作

对于前两类，可以按照异常处理方式再次分为以下几类：

* 抛出异常：add、remove
* 返回特殊值：offer(e)、poll
* 阻塞：put(e)、take
* 超时退出：offer(e, time, unit)、poll(time, unit)

这几种处理方式我就不多解释了，字面意义已经很显然了

## 阻塞队列的实现 ##

JDK8提供了以下BlockQueue的实现类：

![图2](https://user-gold-cdn.xitu.io/2019/6/5/16b257cdbc230dca?imageView2/0/w/1280/h/960/ignore-error/1)

我们常用的基本有以下几种：

* ArrayBlockingQueue：基于ArrayList实现的阻塞队列，有界
* LinkedBlockingQueue：基于LinkedList实现的阻塞队列，有界
* PriorityBlockingQueue：优先队列，无界
* DelayQueue：支持延时获取元素的优先队列，无界

其余的实现感兴趣的可以自行了解，我们这里就以LinkedBlockingQueue为例，介绍一下Java是如何实现阻塞队列的

### 接口方法 ###

除了BlockingQueue提供的接口方法之外，LinkedBlockingQueue还提供了一个方法 ` peek` ，用于获取队首节点

至此，我们常用的阻塞队列方法都已说明完毕，这里用一张表来总结一下 [[1]]( #fn1 ) ：

+---------------+-----------+------------+--------+-------------------------+
| 方法/处理方式 | 抛出异常  | 返回特殊值 |  阻塞  |        超时退出         |
+---------------+-----------+------------+--------+-------------------------+
| 插入元素      | add(e)    | offer(e)   | put(e) | offer(e, timeout, unit) |
| 移除元素      | remove()  | poll()     | take() | poll(timeout, unit)     |
| 获取元素      | element() | peek()     | /      | /                       |
+---------------+-----------+------------+--------+-------------------------+

其中 ` element` 方法和 ` peek` 方法功能是相同的

### 属性 ###

BlockingQueue仅仅定义了接口规范，真正的实现还是由具体的实现类来完成，我们暂且略过中间的AbstractQueue，直接来研究LinkedBlockingQueue，其中定义了几个重要的域对象：

` /** 元素个数 */ private final AtomicInteger count = new AtomicInteger(); /** 队首节点 */ transient Node<E> head; /** 队尾节点 */ private transient Node<E> last; /** take、poll等方法持有的锁，这里叫做take锁或出锁 */ private final ReentrantLock takeLock = new ReentrantLock(); /** take方法的等待队列 */ private final Condition notEmpty = takeLock.newCondition(); /** put、offer等方法持有的锁，这里叫做put锁或入锁 */ private final ReentrantLock putLock = new ReentrantLock(); /** put方法的等待队列 */ private final Condition notFull = putLock.newCondition(); 复制代码`

Node节点就是普通的队列节点，和LinkedList一样，我们主要关注后面的4个域对象，可以分为两类：用于插入元素的，和用于移除元素的。其中每类都有两个属性： ` ReentranLock` 和 ` Condition` 。其中 ` ReentranLock` 是基于AQS [[2]]( #fn2 ) 实现的一个可重入锁（不理解可重入概念的可以当作普通的锁）， ` Condition` 是一个等待/通知模式的具体实现（可以理解为一种提供了功能更强大的 ` wait` 和 ` notify` 的类）

` count` 属性自然不用说， ` head` 和 ` last` 很显然是用于维护存储元素的队列，相信也不用细说了。阻塞队列和普通队列的区分点是在于后面的 ` ReentrantLock` 和 ` Condition` 类型的四个属性，关于这四个属性的意义，在接下来的几个模块会进行深入的分析

不过我们为了接下来讲解方便，先来简单介绍一下 ` Condition` 这个类。实际上， ` Condition` 是一个接口，具体的实现类是在AQS中。对于本篇文章来说，你只需要清楚3个方法： ` await()` 、 ` signal()` ，还有 ` singalAll()` 。这三个方法完全就可以类比 ` wait()` 、 ` notify()` 和 ` notifyAll()` ，它们之间的区别可以模糊地理解为， ` wait/notify` 这些方法管理的是 **对象锁和类锁** ，它们操控的是等待这些锁的线程队列；而 ` await/signal` 这些方法管理的是 **基于AQS的锁** ，操控的自然也是AQS中的线程等待队列

所以这里的 ` notEmpty` 维护了等待 ` take锁` 的线程队列， ` notFull` 维护了等待 ` put锁` 的线程队列。从字面意义上也很好理解， ` notEmpty` 表示“队列还没空”，所以可以取元素，同理， ` notFull` 就表示“队列还没满”，可以往里插入元素

### 插入元素 ###

#### offer(e) ####

先来看 ` offer(e)` 方法，源码如下：

` public boolean offer (E e) { if (e == null ) throw new NullPointerException(); final AtomicInteger count = this.count; // 如果容量达到上限会返回false if (count.get() == capacity) return false ; int c = - 1 ; Node<E> node = new Node<E>(e); final ReentrantLock putLock = this.putLock; // 获取put锁 putLock.lock(); try { if (count.get() < capacity) { // 入队并自增元素个数 enqueue(node); // 注意，这里c返回的是增加前的值 c = count.getAndIncrement(); // 如果容量没到上限，就唤醒一个put操作 if (c + 1 < capacity) notFull.signal(); } } finally { // 解锁 putLock.unlock(); } if (c == 0 ) // 如果队列之前为空，会唤醒一个take操作 signalNotEmpty(); return c >= 0 ; } 复制代码`

这个方法大部分操作都很好理解，当添加元素的操作不允许时， ` offer` 方法会给用户返回 ` false` ，类似于非阻塞的通信方式。 ` offer` 方法的线程安全性是通过 ` put锁` 来保证的

这里有一个很有意思的地方，我们看最后判断如果 ` c == 0` ，那么就会唤醒一个 ` take` 操作。可能很多人疑惑这里为什么要加一条判断，是这样的，整个方法中， ` c` 的初值是 ` -1` ，修改其值的唯一地方就是 ` c = count.getAndIncrement()` 这条语句。也就是说，如果判定 ` c == 0` ，那么这条语句的返回值就是 ` 0` ，即在插入元素之前，队列是空的。所以，如果一开始队列为空，当插入第一个元素之后，会立刻唤醒一个 ` take` 操作 [[3]]( #fn3 )

至此，整个方法流程可以归纳为：

* 获取 ` put锁`
* 元素入队，并自增 ` count` 值
* 如果容量未到上限，则唤醒一个 ` put` 操作
* 如果在插入元素之前队列为空，则在最后唤醒一个 ` take` 操作

#### offer(e, timeout, unit) ####

趁热打铁，我们接着来看带有超时机制的 ` offer` 方法：

` public boolean offer (E e, long timeout, TimeUnit unit) throws InterruptedException { if (e == null ) throw new NullPointerException(); long nanos = unit.toNanos(timeout); int c = - 1 ; final ReentrantLock putLock = this.putLock; final AtomicInteger count = this.count; // 可被中断地获取put锁 putLock.lockInterruptibly(); try { // 重复执行while循环体，直到队列不满，或到了超时时间 while (count.get() == capacity) { // 到了超时时间后就返回false if (nanos <= 0 ) return false ; // 会将当前线程添加到notFull等待队列中， // 返回的是剩余可用的等待时间 nanos = notFull.awaitNanos(nanos); } enqueue( new Node<E>(e)); c = count.getAndIncrement(); if (c + 1 < capacity) notFull.signal(); } finally { putLock.unlock(); } if (c == 0 ) signalNotEmpty(); return true ; } 复制代码`

整个方法大体上和 ` offer(e)` 方法相同，不同的点有两处：

* 获取锁采用的是可中断的形式，即 ` putLock.lockInterruptibly()`
* 如果队列一直是满的，则会循环执行 ` notFull.awaitNanos(nanos)` 操作来将当前线程添加到 ` notFull` 等待队列中（等待 ` put` 操作的执行）

其余部分和 ` offer(e)` 完全一致，在这里就不赘述了

#### add(e) ####

` add` 方法与 ` offer` 方法相比，当操作不允许时，会抛出异常而不是返回一个特殊值，如下：

` public boolean add (E e) { if (offer(e)) return true ; else throw new IllegalStateException( "Queue full" ); } 复制代码`

单纯地就是对 ` offer(e)` 做了二次封装，没什么好说的，需要提一点的就是这个方法的实现是在 ` AbstractQueue` 中

#### put(e) ####

` put(e)` 方法当操作不允许时会阻塞线程，我们来看其是如何实现的：

` public void put (E e) throws InterruptedException { if (e == null ) throw new NullPointerException(); int c = - 1 ; Node<E> node = new Node<E>(e); final ReentrantLock putLock = this.putLock; final AtomicInteger count = this.count; // 以可中断的形式获取put锁 putLock.lockInterruptibly(); try { // 与offer(e, timeout, unit)相比，采用了无限等待的方式 while (count.get() == capacity) { // 当执行了移除元素操作后，会通过signal操作来唤醒notFull队列中的一个线程 notFull.await(); } enqueue(node); c = count.getAndIncrement(); if (c + 1 < capacity) notFull.signal(); } finally { putLock.unlock(); } if (c == 0 ) signalNotEmpty(); } 复制代码`

果然方法之间都是大同小异的， ` put(e)` 操作可以类比我们之前讲的 ` offer(e, timeout, unit)` ，只有一个不同的地方，就是当队列满时， ` await` 操作不再有超时时间，也就是说，只能等待 ` take` 操作 [[4]]( #fn4 ) 来调用 ` signal` 方法唤醒该线程

### 移除元素 ###

#### poll() ####

` poll()` 方法用于移除并返回队首节点，下面是方法的具体实现：

` public E poll () { final AtomicInteger count = this.count; if (count.get() == 0 ) return null ; E x = null ; int c = - 1 ; final ReentrantLock takeLock = this.takeLock; // 获取take锁 takeLock.lock(); try { if (count.get() > 0 ) { // 出队，并自减 x = dequeue(); c = count.getAndDecrement(); if (c > 1 ) // 只要队列还有元素，就唤醒一个take操作 notEmpty.signal(); } } finally { takeLock.unlock(); } // 如果在队列满的情况下移除一个元素，会唤醒一个put操作 if (c == capacity) signalNotFull(); return x; } 复制代码`

如果你认真看了 ` offer(e)` 方法之后， ` poll()` 方法就没什么好讲的了，完全就是 ` offer(e)` 的翻版（我也想讲点东西，但是 ` poll()` 方法完全和 ` offer(e)` 流程一模一样...）

#### 其他 ####

` poll(timeout, unit)/take()/remove()` 方法分别是 ` offer(e, timeout, unit)/put()/add()` 方法的翻版，没有什么特殊的地方，这里就一笔略过了

### 获取元素 ###

#### peek() ####

` peek()` 方法是用于获取队首元素，其实现如下：

` public E peek () { if (count.get() == 0 ) return null ; final ReentrantLock takeLock = this.takeLock; // 获取take锁 takeLock.lock(); try { Node<E> first = head.next; if (first == null ) return null ; else return first.item; } finally { takeLock.unlock(); } } 复制代码`

流程没什么好说的，需要注意的是该方法需要获取 ` take锁` ，也就是说在 ` peek()` 方法执行时，是不能执行移除元素的操作的

#### element() ####

` element()` 方法的实现是在 ` AbstractQueue` 中：

` public E element () { E x = peek(); if (x != null ) return x; else throw new NoSuchElementException(); } 复制代码`

还是同样的二次封装操作

## 总结 ##

本来说的是 ` BlockingQueue` ，结果说了半天 ` LinkedBlockingQueue` 。不过作为阻塞队列的一种经典实现， ` LinkedBlockingQueue` 中的方法实现思路也是对于理解阻塞队列来说也是很重要的。想要理解阻塞队列的理念，最重要的就是理解锁的概念，比如 ` LinkedBlockingQueue` 通过 ` 生产者锁/put锁` 和 ` 消费者锁/take锁` ，以及锁对应的 ` Condition` 对象来实现线程安全。理解了这一点，才能理解整个 ` 生产者/消费者模型`

* 

这里参考了《Java并发编程的艺术》 [↩︎]( #fnref1 )

* 

参见 [浅谈AQS（抽象队列同步器）]( https://juejin.im/post/5cf5cee0f265da1ba647d7f4 ) 一文 [↩︎]( #fnref2 )

* 

这里描述为“唤醒一个 ` take` 操作”有些不准确，实际应描述为“唤醒一个等待 ` take锁` 的线程”，不过我认为前者更有助于读者理解，所以沿用前者的描述方式 [↩︎]( #fnref3 )

* 

指的是和 ` take` 功能类似的一组方法，包含有 ` take/poll/remove` ， ` put` 操作同理 [↩︎]( #fnref4 )