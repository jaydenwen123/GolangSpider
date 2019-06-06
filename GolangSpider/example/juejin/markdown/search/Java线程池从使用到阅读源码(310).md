# Java线程池从使用到阅读源码(3/10) #

我们一般不会选择直接使用线程类 ` Thread` 进行多线程编程，而是使用更方便的线程池来进行任务的调度和管理。线程池就像共享单车，我们只要在我们有需要的时候去获取就可以了。甚至可以说线程池更棒，我们只需要把任务提交给它，它就会在合适的时候运行了。但是如果直接使用 ` Thread` 类，我们就需要在每次执行任务时自己创建、运行、等待线程了，而且很难对线程进行整体的管理，这可不是一件轻松的事情。既然我们已经有了线程池，那还是把这些麻烦事交给线程池来处理吧。

这篇文章将会从线程池的概念与一般使用入手，首先让大家可以了解线程池的基本使用方法，之后会介绍实践中最常用的四种线程池。最后，我们会通过对JDK源代码的剖析深入了解线程池的运行过程和具体设计，真正达到知其然而知其所以然的水平。虽然只要了解了API就可以满足一般的日常使用了，但是只有当我们真正厘清了多线程相关的知识点，才能在面对多线程的实践与面试问题时做到游刃有余、成竹在胸。

本文是一系列多线程文章中的第三篇，主要讲解了线程池相关的知识，这个系列总共有十篇文章，前五篇暂定结构如下，感兴趣的读者可以关注一下：

* 并发基本概念—— [当我们在说“并发、多线程”，说的是什么？]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F58316557 )
* 多线程入门—— [这一次，让我们完全掌握Java多线程(2/10)]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F58316557 )
* 线程池使用与原理剖析——本文
* 线程同步机制
* 并发常见问题

# 线程池的使用方法 #

一般我们最常用的线程池实现类是 ` ThreadPoolExecutor` ，我们接下来会介绍这个类的基本使用方法。JDK已经对线程池做了比较好的封装，相信这个过程会非常轻松。

### 创建线程池 ###

既然线程池是一个Java类，那么最直接的使用方法一定是new一个 ` ThreadPoolExecutor` 类的对象，例如 ` ThreadPoolExecutor threadPool = new ThreadPoolExecutor(1, 1, 0L, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>() )` 。那么这个构造器的里每个参数是什么意思呢？

下面就是这个构造器的方法签名：

` public ThreadPoolExecutor(int corePoolSize, int maximumPoolSize, long keepAliveTime, TimeUnit unit, BlockingQueue<Runnable> workQueue) 复制代码`

各个参数分别表示下面的含义：

* corePoolSize，核心线程池大小，一般线程池会至少保持这么多的线程数量；
* maximumPoolSize，最大线程池大小，也就是线程池最大的线程数量；
* keepAliveTime和unit共同组成了一个超时间， ` keepAliveTime` 是时间数量， ` unit` 是时间单位，单位加数量组成了最终的超时时间。这个超时时间表示如果线程池中包含了超过 ` corePoolSize` 数量的线程，则在有线程空闲的时间超过了超时时间时该线程就会被销毁；
* workQueue是任务的阻塞队列，在没有线程池中没有足够的线程可用的情况下会将任务先放入到这个阻塞队列中等待执行。这里传入的队列类型就决定了线程池在处理这些任务时的策略。

线程池中的阻塞队列专门用于存放待执行的任务，在 ` ThreadPoolExecutor` 中一个任务可以通过两种方式被执行：第一种是直接在创建一个新的Worker时被作为第一个任务传入，由这个新创建的线程来执行；第二种就是把任务放入一个阻塞队列，等待线程池中的工作线程捞取任务进行执行。

上面提到的 **阻塞队列** 是这样的一种数据结构，它是一个队列（类似于一个List），可以存放0到N个元素。我们可以对这个队列进行插入和弹出元素的操作，弹出操作可以理解为是一个获取并从队列中删除一个元素的操作。当队列中没有元素时，对这个队列的获取操作将会被阻塞，直到有元素被插入时才会被唤醒；当队列已满时，对这个队列的插入操作将会被阻塞，直到有元素被弹出后才会被唤醒。这样的一种数据结构非常适合于线程池的场景，当一个工作线程没有任务可处理时就会进入阻塞状态，直到有新任务提交后才被唤醒。

### 提交任务 ###

当创建了一个线程池之后我们就可以将任务提交到线程池中执行了。提交任务到线程池中相当简单，我们只要把原来传入 ` Thread` 类构造器的 ` Runnable` 对象传入线程池的 ` execute` 方法或者 ` submit` 方法就可以了。 ` execute` 方法和 ` submit` 方法基本没有区别，两者的区别只是 ` submit` 方法会返回一个 ` Future` 对象，用于检查异步任务的执行情况和获取执行结果（异步任务完成后）。

我们可以先试试如何使用比较简单的 ` execute` 方法，代码例子如下：

` public class ThreadPoolTest { private static int count = 0; public static void main(String[] args) throws Exception { Runnable task = new Runnable () { public void run () { for (int i = 0; i < 1000000; ++i) { synchronized (ThreadPoolTest.class) { count += 1; } } } }; // 重要：创建线程池 ThreadPoolExecutor threadPool = new ThreadPoolExecutor(1, 1, 0L, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>()); // 重要：向线程池提交两个任务 threadPool.execute(task); threadPool.execute(task); // 等待线程池中的所有任务完成 threadPool.shutdown(); while (!threadPool.awaitTermination(1L, TimeUnit.MINUTES)) { System.out.println( "Not yet. Still waiting for termination" ); } System.out.println( "count = " + count); } } 复制代码`

### 关闭线程池 ###

上面的代码中为了等待线程池中的所有任务执行完已经使用了 ` shutdown()` 方法，关闭线程池的方法主要有两个：

* ` shutdown()` ，有序关闭线程池，调用后线程池会让已经提交的任务完成执行，但是不会再接受新任务。
* ` shutdownNow()` ，直接关闭线程池，线程池中正在运行的任务会被中断，正在等待执行的任务不会再被执行，但是这些还在阻塞队列中等待的任务会被作为返回值返回。

### 监控线程池运行状态 ###

我们可以通过调用线程池对象上的一些方法来获取线程池当前的运行信息，常用的方法有：

* getTaskCount，线程池中已完成、执行中、等待执行的任务总数估计值。因为在统计过程中任务会发生动态变化，所以最后的结果并不是一个准确值；
* getCompletedTaskCount，线程池中已完成的任务总数，这同样是一个估计值；
* getLargestPoolSize，线程池曾经创建过的最大线程数量。通过这个数据可以知道线程池是否充满过，也就是达到过maximumPoolSize；
* getPoolSize，线程池当前的线程数量；
* getActiveCount，当前线程池中正在执行任务的线程数量估计值。

# 四种常用线程池 #

很多情况下我们也不会直接创建 ` ThreadPoolExecutor` 类的对象，而是根据需要通过 ` Executors` 的几个静态方法来创建特定用途的线程池。目前常用的线程池有四种：

* 可缓存线程池，使用 ` Executors.newCachedThreadPool` 方法创建
* 定长线程池，使用 ` Executors.newFixedThreadPool` 方法创建
* 延时任务线程池，使用 ` Executors.newScheduledThreadPool` 方法创建
* 单线程线程池，使用 ` Executors.newSingleThreadExecutor` 方法创建

下面通过这些静态方法的源码来具体了解一下不同类型线程池的特性与适用场景。

### 可缓存线程池 ###

JDK中的源码我们通过在IDE中进行跳转可以很方便地进行查看，下面就是 ` Executors.newCachedThreadPool` 方法中的源代码。从代码中我们可以看到，可缓存线程池其实也是通过直接创建 ` ThreadPoolExecutor` 类的构造器创建的，只是其中的参数都已经被设置好了，我们可以不用做具体的设置。所以我们要观察的重点就是在这个方法中具体产生了一个怎样配置的 ` ThreadPoolExecutor` 对象，以及这样的线程池适用于怎样的场景。

从下面的代码中，我们可以看到，传入 ` ThreadPoolExecutor` 构造器的值有： - corePoolSize核心线程数为0，代表线程池中的线程数可以为0 - maximumPoolSize最大线程数为Integer.MAX_VALUE，代表线程池中最多可以有无限多个线程 - 超时时间设置为60秒，表示线程池中的线程在空闲60秒后会被回收 - 最后传入的是一个 ` SynchronousQueue` 类型的阻塞队列，代表每一个新添加的任务都要马上有一个工作线程进行处理

` public static ExecutorService newCachedThreadPool () { return new ThreadPoolExecutor(0, Integer.MAX_VALUE, 60L, TimeUnit.SECONDS, new SynchronousQueue<Runnable>()); } 复制代码`

所以可缓存线程池在添加任务时会优先使用空闲的线程，如果没有就创建一个新线程，线程数没有上限，所以每一个任务都会马上被分配到一个工作线程进行执行，不需要在阻塞队列中等待；如果线程池长期闲置，那么其中的所有线程都会被销毁，节约系统资源。

* 优点

* 任务在添加后可以马上执行，不需要进入阻塞队列等待
* 在闲置时不会保留线程，可以节约系统资源

* 缺点

* 对线程数没有限制，可能会过量消耗系统资源

* 适用场景

* 适用于大量短耗时任务和对响应时间要求较高的场景

### 定长线程池 ###

传入 ` ThreadPoolExecutor` 构造器的值有:

* corePoolSize核心线程数和maximumPoolSize最大线程数都为固定值 ` nThreads` ，即线程池中的线程数量会保持在 ` nThreads` ，所以被称为“定长线程池”
* 超时时间被设置为0毫秒，因为线程池中只有核心线程，所以不需要考虑超时释放
* 最后一个参数使用了无界队列，所以在所有线程都在处理任务的情况下，可以无限添加任务到阻塞队列中等待执行

` public static ExecutorService newFixedThreadPool(int nThreads) { return new ThreadPoolExecutor(nThreads, nThreads, 0L, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>()); } 复制代码`

定长线程池中的线程数会逐步增长到nThreads个，并且在之后空闲线程不会被释放，线程数会一直保持在 ` nThreads` 个。如果添加任务时所有线程都处于忙碌状态，那么就会把任务添加到阻塞队列中等待执行，阻塞队列中任务的总数没有上限。

* 优点

* 线程数固定，对系统资源的消耗可控

* 缺点

* 在任务量暴增的情况下线程池不会弹性增长，会导致任务完成时间延迟
* 使用了无界队列，在线程数设置过小的情况下可能会导致过多的任务积压，引起任务完成时间过晚和资源被过度消耗的问题

* 适用场景

* 任务量峰值不会过高，且任务对响应时间要求不高的场景

### 延时任务线程池 ###

与之前的两个方法不同， ` Executors.newScheduledThreadPool` 返回的是 ` ScheduledExecutorService` 接口对象，可以提供延时执行、定时执行等功能。在线程池配置上有如下特点：

* maximumPoolSize最大线程数为无限，在任务量较大时可以创建大量新线程执行任务
* 超时时间为0，线程空闲后会被立即销毁
* 使用了延时工作队列，延时工作队列中的元素都有对应的过期时间，只有过期的元素才会被弹出

` public static ScheduledExecutorService newScheduledThreadPool(int corePoolSize) { return new ScheduledThreadPoolExecutor(corePoolSize); } public ScheduledThreadPoolExecutor(int corePoolSize) { super(corePoolSize, Integer.MAX_VALUE, 0, NANOSECONDS, new DelayedWorkQueue()); } 复制代码`

延时任务线程池实现了 ` ScheduledExecutorService` 接口，主要用于需要延时执行和定时执行的情况。

### 单线程线程池 ###

单线程线程池中只有一个工作线程，可以保证添加的任务都以指定顺序执行（先进先出、后进先出、优先级）。但是如果线程池里只有一个线程，为什么我们还要用线程池而不直接用 ` Thread` 呢？这种情况下主要有两种优点：一是我们可以通过共享的线程池很方便地提交任务进行异步执行，而不用自己管理线程的生命周期；二是我们可以使用任务队列并指定任务的执行顺序，很容易做到任务管理的功能。

` public static ExecutorService newSingleThreadExecutor () { return new FinalizableDelegatedExecutorService (new ThreadPoolExecutor(1, 1, 0L, TimeUnit.MILLISECONDS, new LinkedBlockingQueue<Runnable>())); } 复制代码`

# 线程池的内部实现 #

通过前面的内容我们其实已经可以在代码中使用线程池了，但是我们为什么还要去深究线程池的内部实现呢？首先，可能有一个很功利性的目的就是为了面试，在面试时如果能准确地说出一些底层的运行机制与原理那一定可以成为过程中一个重要的亮点。

但是我认为学习探究线程池的内部实现的作用绝对不仅是如此，只有深入了解并厘清了线程池的具体实现，我们才能解决实践中需要考虑的各种边界条件。因为多线程编程所代表的并发编程并不是一个固定的知识点，而是实践中不断在发展和完善的一个知识门类。我们也许会需要同时考虑多个维度，最后得到一个特定于应用场景的解决方案，这就要求我们具备从细节着手构建出解决方案并做好各个考虑维度之间的取舍的能力。

而且我相信只要在某一个点上能突破到相当的深度，那么以后从这个点上向外扩展就会容易得多。也许在刚开始我们的探究会碰到非常大的阻力，但是我们要相信，最后我们可以得到的将不止是一个知识点而是一整个知识面。

### 查看JDK源码的方式 ###

在IDE中，例如IDEA里，我们可以点击我们样例代码里的 ` ThreadPoolExecutor` 类跳转到JDK中 ` ThreadPoolExecutor` 类的源代码。在源代码中我们可以看到很多 ` java.util.concurrent` 包的缔造者大牛“Doug Lea”所留下的各种注释，下面的图片就是该类源代码的一个截图。

![](https://user-gold-cdn.xitu.io/2019/3/18/1699026a0e59e47f?imageView2/0/w/1280/h/960/ignore-error/1)

这些注释的内容非常有参考价值，建议有能力的读者朋友可以自己阅读一遍。下面，我们就一步步地抽丝剥茧，来揭开线程池类 ` ThreadPoolExecutor` 源代码的神秘面纱。

### 控制变量与线程池生命周期 ###

在 ` ThreadPoolExecutor` 类定义的开头，我们可以看到如下的几行代码：

` // 控制变量，前3位表示状态，剩下的数据位表示有效的线程数 private final AtomicInteger ctl = new AtomicInteger(ctlOf(RUNNING, 0)); // Integer的位数减去3位状态位就是线程数的位数 private static final int COUNT_BITS = Integer.SIZE - 3; // CAPACITY就是线程数的上限（含），即2^COUNT_BITS - 1个 private static final int CAPACITY = (1 << COUNT_BITS) - 1; 复制代码`

第一行是一个用来作为控制变量的整型值，即一个Integer。之所以要用 ` AtomicInteger` 类是因为要保证多线程安全，在本系列之后的文章中会对 ` AtomicInteger` 进行具体介绍。一个整型一般是32位，但是这里的代码为了保险起见，还是使用了 ` Integer.SIZE` 来表示整型的总位数。这里的“位”指的是数据位(bit)，在计算机中，8bit = 1字节，1024字节 = 1KB，1024KB = 1MB。每一位都是一个0或1的数字，我们如果把整型想象成一个二进制(0或1)的数组，那么一个Integer就是32个数字的数组。其中，前三个被用来表示状态，那么我们就可以表示2^3 = 8个不同的状态了。剩下的29位二进制数字都会被用于表示当前线程池中有效线程的数量，上限就是(2^29 - 1)个，即常量 ` CAPACITY` 。

之后的部分列出了线程池的所有状态：

` private static final int RUNNING = -1 << COUNT_BITS; private static final int SHUTDOWN = 0 << COUNT_BITS; private static final int STOP = 1 << COUNT_BITS; private static final int TIDYING = 2 << COUNT_BITS; private static final int TERMINATED = 3 << COUNT_BITS; 复制代码`

在这里可以忽略数字后面的 ` << COUNT_BITS` ，可以把状态简单地理解为前面的数字部分，这样的简化基本不影响结论。

各个状态的解释如下：

* RUNNING，正常运行状态，可以接受新的任务和处理队列中的任务
* SHUTDOWN，关闭中状态，不能接受新任务，但是可以处理队列中的任务
* STOP，停止中状态，不能接受新任务，也不处理队列中的任务，会中断进行中的任务
* TIDYING，待结束状态，所有任务已经结束，线程数归0，进入TIDYING状态后将会运行 ` terminated()` 方法
* TERMINATED，结束状态， ` terminated()` 方法调用完成后进入

这几个状态所对应的数字值是按照顺序排列的，也就是说线程池的状态只能从小到大变化，这也方便了通过数字比较来判断状态所在的阶段，这种通过数字大小来比较状态值的方法在 ` ThreadPoolExecutor` 的源码中会有大量的使用。

下图是这五个状态之间的变化过程：

![](https://user-gold-cdn.xitu.io/2019/3/18/169902818a24b206?imageView2/0/w/1280/h/960/ignore-error/1)

* 当线程池被创建时会处于 **RUNNING** 状态，正常接受和处理任务；
* 当 ` shutdown()` 方法被直接调用，或者在线程池对象被GC回收时通过 ` finalize()` 方法隐式调用了 ` shutdown()` 方法时，线程池会进入 **SHUTDOWN** 状态。该状态下线程池仍然会继续执行完阻塞队列中的任务，只是不再接受新的任务了。当队列中的任务被执行完后，线程池中的线程也会被回收。当队列和线程都被清空后，线程池将进入 **TIDYING** 状态；
* 在线程池处于 **RUNNING** 或者 **SHUTDOWN** 状态时，如果有代码调用了 ` shutdownNow()` 方法，则线程池会进入 **STOP** 状态。在 **STOP** 状态下，线程池会直接清空阻塞队列中待执行的任务，然后中断所有正在进行中的任务并回收线程。当线程都被清空以后，线程池就会进入 **TIDYING** 状态；
* 当线程池进入 **TIDYING** 状态时，将会运行 ` terminated()` 方法，该方法执行完后，线程池就会进入最终的 **TERMINATED** 状态，彻底结束。

到这里我们就已经清楚地了解了线程从刚被创建时的 **RUNNING** 状态一直到最终的 **TERMINATED** 状态的整个生命周期了。那么当我们要向一个 **RUNNING** 状态的线程池提交任务时会发生些什么呢？

### execute方法的实现 ###

我们一般会使用 ` execute` 方法提交我们的任务，那么线程池在这个过程中做了什么呢？在 ` ThreadPoolExecutor` 类的 ` execute()` 方法的源代码中，我们主要做了四件事：

* 如果当前线程池中的线程数小于核心线程数corePoolSize，则创建一个新的Worker代表一个线程，并把入参中的任务作为第一个任务传入Worker。 ` addWorker` 方法中的第一个参数是该线程的第一个任务，而第二个参数就是代表是否创建的是核心线程，在 ` execute` 方法中 ` addWorker` 总共被调用了三次，其中第一次传入的是true，后两次传入的都是false；
* 如果当前线程池中的线程数已经满足了核心线程数corePoolSize，那么就会通过 ` workQueue.offer()` 方法将任务添加到阻塞队列中等待执行；
* 如果线程数已经达到了corePoolSize且阻塞队列中无法插入该任务（比如已满），那么线程池就会再增加一个线程来执行该任务，除非线程数已经达到了最大线程数maximumPoolSize；
* 如果确实已经达到了最大线程数，那么就拒绝这个任务。

总体上的执行流程如下，下方的黑色同心圆代表流程结束：

![](https://user-gold-cdn.xitu.io/2019/3/18/16990288bb279ea5?imageView2/0/w/1280/h/960/ignore-error/1) 这里再重复一次阻塞队列的定义，方便大家阅读：

> 
> 
> 
> 线程池中的阻塞队列专门用于存放待执行的任务，在 ` ThreadPoolExecutor` 中一个任务可以通过两种方式被执行：第一种是直接在创建一个新的Worker时被作为第一个任务传入，由这个新创建的线程来执行；第二种就是把任务放入一个阻塞队列，等待线程池中的工作线程捞取任务进行执行。
> 
> 
> 

> 
> 
> 
> 上面提到的 **阻塞队列** 是这样的一种数据结构，它是一个队列（类似于一个List），可以存放0到N个元素。我们可以对这个队列进行插入和弹出元素的操作，弹出操作可以理解为是一个获取并从队列中删除一个元素的操作。当队列中没有元素时，对这个队列的获取操作将会被阻塞，直到有元素被插入时才会被唤醒；当队列已满时，对这个队列的插入操作将会被阻塞，直到有元素被弹出后才会被唤醒。这样的一种数据结构非常适合于线程池的场景，当一个工作线程没有任务可处理时就会进入阻塞状态，直到有新任务提交后才被唤醒。
> 
> 
> 

下面是带有注释的源代码，大家可以和上面的流程对照起来参考一下：

` public void execute(Runnable command ) { // 检查提交的任务是否为空 if ( command == null) throw new NullPointerException(); // 获取控制变量值 int c = ctl.get(); // 检查当前线程数是否达到了核心线程数 if (workerCountOf(c) < corePoolSize) { // 未达到核心线程数，则创建新线程 // 并将传入的任务作为该线程的第一个任务 if (addWorker( command , true )) // 添加线程成功则直接返回，否则继续执行 return ; // 因为前面调用了耗时操作addWorker方法 // 所以线程池状态有可能发生了改变，重新获取状态值 c = ctl.get(); } // 判断线程池当前状态是否是运行中 // 如果是则调用workQueue.offer方法将任务放入阻塞队列 if (isRunning(c) && workQueue.offer( command )) { // 因为执行了耗时操作“放入阻塞队列”，所以重新获取状态值 int recheck = ctl.get(); // 如果当前状态不是运行中，则将刚才放入阻塞队列的任务拿出，如果拿出成功，则直接拒绝这个任务 if (! isRunning(recheck) && remove( command )) reject( command ); else if (workerCountOf(recheck) == 0) // 如果线程池中没有线程了，那就创建一个 addWorker(null, false ); } // 如果放入阻塞队列失败（如队列已满），则添加一个线程 else if (!addWorker( command , false )) // 如果添加线程失败（如已经达到了最大线程数），则拒绝任务 reject( command ); } 复制代码`

### addWorker方法 ###

在前面 ` execute` 方法的代码中我们可以看到线程池是通过 ` addWorker` 方法来向线程池中添加新线程的，那么新的线程又是如何运行起来的呢？

这里我们暂时跳过 ` addWorker` 方法的详细源代码，因为虽然这个方法的代码行数较多，但是功能相对比较直接，只是创建一个代表线程的 ` Worker` 类对象，并调用这个对象所对应线程对象的 ` start()` 方法。我们知道一旦调用了 ` Thread` 类的 ` start()` 方法，则这个线程就会开始调用创建线程时传入的 ` Runnable` 对象。从下面的 ` Worker` 类构造器源代码可以看出， ` Worker` 类正是把自己(this指针)传入了线程的构造器当中，那么这个线程就会运行 ` Worker` 类的 ` run()` 方法了，这个 ` run()` 方法只执行了一行很简单的代码 ` runWorker(this);` 。

` Worker(Runnable firstTask) { set State(-1); // inhibit interrupts until runWorker this.firstTask = firstTask; this.thread = getThreadFactory().newThread(this); } public void run () { runWorker(this); } 复制代码`

### runWorker方法的实现 ###

我们看到线程池中的线程在启动时会调用对应的 ` Worker` 类的 ` runWorker` 方法，而这里就是整个线程池任务执行的核心所在了。 ` runWorker` 方法中包含有一个类似无限循环的while语句，让worker对象可以不断执行提交到线程池中的新任务。

大家可以配合代码上带有的注释来理解该方法的具体实现：

` final void runWorker(Worker w) { Thread wt = Thread.currentThread(); Runnable task = w.firstTask; w.firstTask = null; // 将worker的状态重置为正常状态，因为state状态值在构造器中被初始化为-1 w.unlock(); // 通过completedAbruptly变量的值判断任务是否正常执行完成 boolean completedAbruptly = true ; try { // 如果task为null就通过getTask方法获取阻塞队列中的下一个任务 // getTask方法一般不会返回null，所以这个 while 类似于一个无限循环 // worker对象就通过这个方法的持续运行来不断处理新的任务 while (task != null || (task = getTask()) != null) { // 每一次任务的执行都必须获取锁来保证下方临界区代码的线程安全 w.lock(); // 如果状态值大于等于STOP（状态值是有序的，即STOP、TIDYING、TERMINATED） // 且当前线程还没有被中断，则主动中断线程 if ((runStateAtLeast(ctl.get(), STOP) || (Thread.interrupted() && runStateAtLeast(ctl.get(), STOP))) && !wt.isInterrupted()) wt.interrupt(); // 开始 try { // 执行任务前处理操作，默认是一个空实现 // 在子类中可以通过重写来改变任务执行前的处理行为 beforeExecute(wt, task); // 通过thrown变量保存任务执行过程中抛出的异常 // 提供给下面finally块中的afterExecute方法使用 Throwable thrown = null; try { // *** 重要：实际执行任务的代码 task.run(); } catch (RuntimeException x) { thrown = x; throw x; } catch (Error x) { thrown = x; throw x; } catch (Throwable x) { // 因为Runnable接口的run方法中不能抛出Throwable对象 // 所以要包装成Error对象抛出 thrown = x; throw new Error(x); } finally { // 执行任务后处理操作，默认是一个空实现 // 在子类中可以通过重写来改变任务执行后的处理行为 afterExecute(task, thrown); } } finally { // 将循环变量task设置为null，表示已处理完成 task = null; // 累加当前worker已经完成的任务数 w.completedTasks++; // 释放 while 体中第一行获取的锁 w.unlock(); } } // 将completedAbruptly变量设置为 false ，表示任务正常处理完成 completedAbruptly = false ; } finally { // 销毁当前的worker对象，并完成一些诸如完成任务数量统计之类的辅助性工作 // 在线程池当前状态小于STOP的情况下会创建一个新的worker来替换被销毁的worker processWorkerExit(w, completedAbruptly); } } 复制代码`

在 ` runWorker` 方法的源代码中有两个比较重要的方法调用，一个是while条件中对 ` getTask` 方法的调用，一个是在方法的最后对 ` processWorkerExit` 方法的调用。下面是对这两个方法更详细的解释。

` getTask` 方法在阻塞队列中有待执行的任务时会从队列中弹出一个任务并返回，如果阻塞队列为空，那么就会阻塞等待新的任务提交到队列中直到超时（在一些配置下会一直等待而不超时），如果在超时之前获取到了新的任务，那么就会将这个任务作为返回值返回。

当 ` getTask` 方法返回null时会导致当前Worker退出，当前线程被销毁。在以下情况下 ` getTask` 方法才会返回null：

* 当前线程池中的线程数超过了最大线程数。这是因为运行时通过调用 ` setMaximumPoolSize` 修改了最大线程数而导致的结果；
* 线程池处于STOP状态。这种情况下所有线程都应该被立即回收销毁；
* 线程池处于SHUTDOWN状态，且阻塞队列为空。这种情况下已经不会有新的任务被提交到阻塞队列中了，所以线程应该被销毁；
* 线程可以被超时回收的情况下等待新任务超时。线程被超时回收一般有以下两种情况：

* 超出核心线程数部分的线程等待任务超时
* 允许核心线程超时（线程池配置）的情况下线程等待任务超时

` processWorkerExit` 方法会销毁当前线程对应的Worker对象，并执行一些累加总处理任务数等辅助操作。但在线程池当前状态小于STOP的情况下会创建一个新的Worker来替换被销毁的Worker，有兴趣的读者可以自行参考 ` processWorkerExit` 方法源代码。

# 总结 #

到这里我们的线程池源代码之旅就结束了，希望大家在看完这篇文章之后能对线程池的使用和运行都有一个大概的印象。为什么说只是有了一个大概的印象呢？因为我觉得很多没有相关基础的读者读到这里可能还只是对线程池有了一个自己的认识，对其中的一些细节可能还没有完全捕捉到。所以我建议大家在看完下面的总结之后不妨再返回到文章的开头多读几遍，相信第二遍的阅读能给大家带来不一样的体验，因为我自己也是在第三次读 ` ThreadPoolExecutor` 类的源代码时才真正打通了其中的一些重要关节的。

在这篇文章中我们从线程池的概念和基本使用方法说起，然后介绍了 ` ThreadPoolExecutor` 的构造器参数和常用的四种具体配置。最后的一大半篇幅我们一起在 ` TheadPoolExecutor` 类的源代码中畅游了一番，了解了从线程池的创建到任务执行的完整执行模型。

# 引子 #

在浏览 ` ThreadPoolExexutor` 源码的过程中，有几个点我们其实并没有完全说清楚，比如对锁的加锁操作、对控制变量的多次获取、控制变量的AtomicInteger类型。在下一篇文章中，我将会介绍这些以锁、volatile变量、CAS操作、AQS抽象类为代表的一系列线程同步方法，欢迎感兴趣的读者继续关注我后续发布的文章~