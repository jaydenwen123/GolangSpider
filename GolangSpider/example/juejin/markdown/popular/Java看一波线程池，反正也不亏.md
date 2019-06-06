# Java:看一波线程池，反正也不亏 #

## 前言 ##

线程池在Java并发编程中，有着举足轻重的位置，学习和掌握它是学习Java的重中之重。反正有空看看，学点知识，又不亏。

在开发中，合理使用 **线程池** 能带来什么好处呢？

* **提高响应速度** 。当任务到达时，线程已经建立好，立即执行。
* **降低资源消耗** 。通过重复使用线程降低新建和销毁线程带来的开销。
* **提高线程的可管理性** 。线程池可以进行统一分配、调优和监控线程的情况，对资源的管控。

## 队列 ##

线程池内部持有一个用于存储工作任务的队列，在核心线程满了时候，会将任务存储到队列中。常用队列类型：

* **LinkedBlockingQueue** ：是基于链表结构的有界阻塞队列 。
* **ArraryBlockingQueue** ：是基于数组结构的有界阻塞队列。
* **SynchronousQueue** ：一个不存储元素的阻塞队列。每个插入操作必须等到一个移除操作，下一个插入操作才能进行。
* **PriorityBlocking** ：具有优先级的无界阻塞队列。
* **DelayQueue** ：延时的无界阻塞队列。

## 线程池的使用 ##

* **创建**
` ThreadPoolExecutor( int corePoolSize, int maximumPoolSize, long keepAliveTime, TimeUnit unit, BlockingQueue<Runnable> workQueue, ThreadFactory threadFactory, RejectedExecutionHandler handler) 复制代码`

在线程的构造方法中，共有5个参数：

* corePoolSize：线程池核心线程的数量。当线程池的线程数量未达到核心线程数量，每提交一个任务都会创建一个线程，直到线程数量等于核心线程数量。核心线程指一直存活在线程池中，不会被会销毁（可设置超时销毁），直到线程池关闭。而超过corePoolSize数量创建的线程的就是非核心线程，在空闲keepAliveTime时间后被销毁。
* maximumPoolSize：线程池最大线程数量。核线程最大数量=核心线程数量+非核心检查数量。使用无限的任务队列（如PriorityBlocking），会导致该参数失效。
* keepAliveTime：非核心线程闲置时长。闲置超过该时长，非核心线程会被回收。
* unit：keepAliveTime的时间单位。
* workQueue：工作队列，如前文讲到的四种队列。
* threadFactory：用于创建线程的工厂。可以通过 ` Executors.defaultThreadFactory();` 获得默认的线程工厂。主要作用就是给线程起个名字而已。
* handler：饱和策略。当队列和线程池都满了，此时该怎么处理新任务？ThreadPoolExecutor内有四个内部类实现的策略供选择。AbortPolicy：直接抛出异常；DiscardPolicy：不处理，抛弃新任务；DiscardOldestPolicy：抛弃队列总最近的一个队列；CallerRunsPolicy：使用调用者所在线程执行新任务。默认AbortPolicy策略。

* **提交任务到线程池**
` //execute()方法执行任务 executor.execute(new Runnable () { @Override public void run () { // TODO Auto-generated method stub } }); //submit()方法提交任务 executor.execute(new Runnable () { @Override public void run () { // TODO Auto-generated method stub } }); 复制代码`

提交到线程池中有种方式：execute()方法和submit()方法。

execute()方法提交的任务无法获得返回值，无法判断提交状态。 submit()方法可以获得返回Future类型的对象，根据Future对象可以判断任务是否执行成功和通过get()方法获得返回值,但get()方法会阻塞当前线程一段时间。

* **线程池的关闭**

* shutdown() 将线程池的状态设置为 ` SHUTDOWN` 状态。然后中断所有没有正在执行任务的线程。
* shutdownNow() 将线程池的状态设置为 ` STOP` 状态。尝试停止所有正在执行或者暂停任务的线程，并返回等待执行任务的列表。

由于两者都是遍历线程池中的工作线程，然后中断线程，所以无法响应中断的线程可能永远无法终止。

## 线程池的实现原理 ##

当我们提交一个新任务到线程池时，任务在线程池的流程是怎样的呢？

* 核心线程是否已满。核心线程大于等于corePoolSize，表示已满。不是的话，则会创建新的核心线程执行新任务（只要核心线程未达到数量corePoolSize，都会新建）。不然执行下一步骤。
* 工作队列是否已经满。在有界队列中，存储任务的数量有限。如果任务未满了，新提交任务存储到工作队列。否则进入下个流程。
* 非核心线程是否已满。虽然核心线程队列已满，但非核心线程不一定满了。非核心线程大于maximumPoolSize，表示已满，线程池拒绝新任务，交给饱和策略处理。否则新建非核心线程处理新任务。

任务队列的任务什么时候被处理？

线程池中线程处理任务有两种方式：一种就是新建线程处理任务，另一种就是循环从阻塞队列获取任务来执行。

## 四种线程池 ##

有时我们仅仅是使用一下线程池，不会自己定制线程池，毕竟线程池的构造方法参数那么多，我的妈耶。那看看类Executors提供的四个线程池。

### FixedThreadPool ###

` //创建FixedThreadPool线程池 Executors.newFixedThreadPool(nThreads); //newFixedThreadPool方法的实现 public static ExecutorService newFixedThreadPool(int nThreads) { return new ThreadPoolExecutor(nThreads, nThreads, 0L, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>()); } 复制代码`

创建一个FixedThreadPool线程池，线程的核心线程coreThreadSize和线程池线程容量maximumPoolSize都为nThreads。由于使用的无界的LinkedBlockingQueue队列，将导致maximumPoolSize参数失效，队列对任务将来者不拒。这里将保活时间设为0，意味着空线程会被立即终止。

### CachedThreadPool ###

` //创建CachedThreadPool线程池 Executors.newCachedThreadPool(); //newCachedThreadPool方法的实现 public static ExecutorService newCachedThreadPool () { return new ThreadPoolExecutor(0, Integer.MAX_VALUE, 60L, TimeUnit.SECONDS, new SynchronousQueue<Runnable>()); } 复制代码`

CachedThreadPool的核心线程为0，线程池容量为Integer.MAX_VALUE，意味着maximumPoolSize是无界的。保活时间keepAliveTime设为60秒，空闲线程闲置60秒后被终止。由于使用了没有容量的SynchronousQueue队列，意味当提交一个新任务到线程池中，没有空闲线程来对接，就会新建新的线程来处理新任务。

### SingleThreadExecutor ###

` //创建SingleThreadExecutor线程池 Executors.newSingleThreadExecutor(); //SingleThreadExecutor方法的实现 public static ExecutorService newSingleThreadExecutor () { return new FinalizableDelegatedExecutorService (new ThreadPoolExecutor(1, 1, 0L, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>())); } 复制代码`

可以看到SingleThreadExecutor可以看做定制化的FixedThreadPool，将nThreads置为1，将核心线程和线程池容量设为1，以保证只有一个线程在执行。

### ScheduledThreadPool ###

` //创建ScheduledThreadPool线程池 Executors.newScheduledThreadPool(corePoolSize); //newScheduledThreadPool方法的实现 public static ScheduledExecutorService newScheduledThreadPool(int corePoolSize) { return new ScheduledThreadPoolExecutor(corePoolSize); } //ScheduledThreadPoolExecutor方法的实现 public ScheduledThreadPoolExecutor(int corePoolSize) { super(corePoolSize, Integer.MAX_VALUE, 0, NANOSECONDS, new DelayedWorkQueue()); } 复制代码`

newScheduledThreadPool方法最终会调用ThreadPoolExecutor的构造来创建线程池。将定制化DelayQueue后的DelayedWorkQueue作为工作队列。DelayedWorkQueue队列会把执行时间小的任务排在前面优先执行。如果执行时间相同，就会优先执行提交时间早的任务。

## FutureTask ##

在通过 ` submit()` 方法提交任务到线程池，会返回有结果的Future类型的对象。Future是一个接口，FutureTask继承它,所以FutureTask也可以作为 ` submit()` 方法得返回值。同时，FutureTask继承Runnable接口，这样又可以作为任务提交到线程池中，由调用线程直接执行。

当调用FutureTask的 ` get()` 方法，如果FutureTask处于已完成状态（执行完毕），则会导致调用线程立即放回或者抛出异常。否则会使调用线程阻塞。

当调用FutureTask的 ` cancel()` 方法，如果FutureTask处于未启动状态（为执行run方法），则该任务不会被执行；如果处于启动状态，会尝试以中断来尝试停止任务。如果已经完成状态，则返回false。

## 总结 ##

通过对线程池知识点理解，可以清晰掌握创建线程池骚姿势和内涵。以为实战提供必要的理论知识。

![点个赞，老铁](https://user-gold-cdn.xitu.io/2019/5/30/16b06891c7562a34?imageView2/0/w/1280/h/960/format/png/ignore-error/1)

如果觉得文章有用，给文章点个赞，铁子

本文是个人学习总结和知识备忘，如知识有误或片面，请多加指正，谢谢

知识来源《Java并发编程的艺术》& 互联网 & 《Java并发编程实战》