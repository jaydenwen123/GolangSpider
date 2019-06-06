# 详解Java线程安全 #

## 一、内存模型 ##

### **高速缓存** ###

因为CPU执行速度和内存数据读写速度差距很大，因此CPU往往包含 ` 高速缓存` 结构。

![此处输入图片的描述](https://user-gold-cdn.xitu.io/2019/4/9/16a024f01184343c?imageView2/0/w/1280/h/960/ignore-error/1) 当程序在运行过程中，会将运算需要的数据从主存复制一份到CPU的高速缓存当中，那么CPU进行计算时就可以直接从它的高速缓存读取数据和向其中写入数据，当运算结束之后，再将高速缓存中的数据刷新到主存当中。

### **缓存不一致问题** ###

执行下面的代码：

` int i = 0 ; i = i + 1 ; 复制代码`

当线程执行这个语句时，会先从主存当中读取i的值 ` i = 0` ，然后复制一份到 ` 高速缓存` 当中，然后CPU执行指令对 ` i` 进行加1操作，然后将数据写入高速缓存，最后将高速缓存中 ` i` 最新的值刷新到 ` 主存` 当中。

可能存在情况：初始时，两个线程分别读取i的值存入各自所在的CPU的高速缓存当中，然后 ` 线程1` 进行加1操作，然后把i的最新值1写入到内存。此时线程2的高速缓存当中i的值还是0，进行加1操作之后，i的值为1，然后 ` 线程2` 把i的值写入内存。

也就是说，如果一个变量在多个CPU中都存在缓存（多线程情况），那么就可能存在 **缓存不一致** 的问题。

### **缓存不一致的解决** ###

一般有两种解决办法：

* **总线加锁**

> 
> 
> 
> 因为CPU和其他部件进行通信都是通过总线来进行的，如果对总线加锁的话，也就是说阻塞了其他CPU对其他部件访问（如内存），从而使得只能有一个CPU能使用这个变量的内存。
> 
> 
> 

* **缓存一致性协议**

> 
> 
> 
> 由于在锁住总线期间，其他CPU无法访问内存，导致效率低下。所以就出现了缓存一致性协议。最出名的就是Intel的 ` MESI协议` ， `
> MESI协议` 保证了每个缓存中使用的共享变量的副本是一致的。 ` MESI协议` 核心思想是：当CPU写数据时，如果发现操作的变量是共享变量，即在其他CPU中也存在该变量的副本，会发出信号通知其他CPU将该变量的缓存行置为无效状态，因此当其他CPU需要读取这个变量时，发现自己缓存中缓存该变量的缓存行是无效的，那么它就会从内存重新读取。
> 
> 
> 

## 二、线程安全问题 ##

### **产生原因** ###

从前面的分析，在并发编程（多线程编程）中，可能出现线程安全的问题：

* 

多个线程在操作共享的数据。

* 

操作共享数据的线程代码有多条。

* 

当一个线程在执行操作共享数据的多条代码过程中，其他线程参与了运算。

### **并发的核心概念** ###

**三个核心概念：原子性、可见性、顺序性。**

* 原子性：跟数据库事务的原子性概念差不多，即一个操作（有可能包含有多个子操作）要么全部执行（生效），要么全部都不执行（都不生效）。

> 
> 
> 
> ` 锁和同步` （同步方法和同步代码块）、 ` CAS` （CPU级别的CAS指令 ` cmpxchg` ）。
> 
> 

* 可见性：当多个线程并发访问共享变量时，一个线程对共享变量的修改，其它线程能够立即看到。

> 
> 
> 
> ` volatile` 关键字来保证可见性。
> 
> 

* 顺序性：程序执行的顺序按照代码的先后顺序执行。因为处理器为了提高程序运行效率，可能会对输入代码进行优化，它不保证程序中各个语句的执行先后顺序同代码中的顺序一致，但是它会保证程序最终执行结果和代码顺序执行的结果是一致的-即 ` 指令重排序` 。

> 
> 
> 
> ` volatile` 在一定程序上保证顺序性，另外还可以通过 ` synchronized` 和 ` 锁` 来保证顺序性。
> 
> 

## 三、Java对象头的结构 ##

Java对象可以作为并发编程中的锁。而锁实际上存在于Java对象头里。如果对象是数组类型，则虚拟机用 3 个 Word（字宽）存储对象头，如果对象是非数组类型，则用 2 字宽存储对象头。在 64 位虚拟机中，一字宽等于八字节，即 ` 64bit` 。

Java 对象头里的 ` Mark Word` 里默认存储对象的 HashCode，分代年龄和锁标记位。32 位 JVM 的 Mark Word 的默认存储结构如下：

|-|25 bit|4bit|偏向锁标志位(1bit)|锁标志位(2bit)| |::|::|::|::|::| |无锁状态|对象的hashCode|对象分代年龄| |01|

64 位JVM的存储结构如下：

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------+-------+------+------+------+------+
|                                                                                                                                                                                                                                                                                               |       |       |      |      |      |      |
| **锁状态**                                                                                                                                                                                                                                                                                    | 25bit | 31bit | 1bit | 4bit | 1bit | 2bit |
|                                                                                                                                                                                                                                                                                               |       |       |      |      |      |      |
|  ` </td> <td> </td> <td> <p><span lang="EN-US">cms_free</span></p> </td> <td> <p><span>分代年龄<span lang="EN-US"></span></span></p> </td> <td colspan="2"> <p><span>偏向锁<span lang="EN-US"></span></span></p> </td> <td> <p><span>锁标志位<span lang="EN-US"></span></span></p>            |
| </td> </tr><tr><td> <p><span>无锁<span lang="EN-US"></span></span></p> </td> <td> <p><span lang="EN-US">unused</span></p> </td> <td> <p><span lang="EN-US">hashCode</span></p> </td> <td> </td> <td> </td> <td> </td> <td colspan="2"> <p><span lang="EN-US">01</span></p>                    |
| </td> </tr><tr><td> <p><span>偏向锁<span lang="EN-US"></span></span></p> </td> <td colspan="2"> <p><span lang="EN-US">ThreadID(54bit) Epoch(2bit)</span></p> </td> <td> </td> <td> </td> <td> <p><span lang="EN-US">1</span></p> </td> <td colspan="2"> <p><span                              |
| lang="EN-US">01</span></p> </td> </tr></tbody></table> 复制代码`  在运行期间 ` Mark Word` 里存储的数据会随着锁标志位的变化而变化。    **在了解了相关概念后，接下来介绍Java是如何保证并发编程中的安全的。**    ## **四、synchronized** ##  ### **用法** ###  * 修饰同步代码块                  |
|  >  >  >  > 将多条操作共享数据的线程代码封装起来，当有线程在执行这些代码的时候，其他线程时不可以参与运算的。必须要当前线程把这些代码都执行完毕后，其他线程才可以参与运算。 >  >  >  ` synchronized (对象) { 需要被同步的代码 ； } 复制代码` * 修饰同步函数(方法)  `                           |
| 修饰符 synchronized 返回值 方法名(){ } 复制代码` *   修饰一个静态的方法，其作用的范围是整个静态方法，作用的对象是这个类的所有对象；  *   修饰一个类，其作用的范围是 ` synchronized` 后面括号括起来的部分，作用主的对象是这个类的所有对象。  >  >  >  > ` synchronized`                        |
| 的作用主要有三个： （1）确保线程互斥的访问同步代码 （2）保证共享变量的修改能够及时可见 （3）有效解决重排序问题。 >  >  >  ### **锁对象** ###  * 对于同步方法，锁是当前 ` 实例对象` 。 * 对于静态同步方法，锁是当前对象的 ` Class 对象` 。 * 对于同步方法块，锁是 `                            |
| synchonized` 括号里配置的对象。  ### **实现原理** ###  在编译的字节码中加入了两条指令来进行代码的同步。  #### **monitorenter ：** ####  每个对象有一个 ` 监视器锁（monitor）` 。当 ` monitor` 被占用时就会处于锁定状态，线程执行 ` monitorenter` 指令时尝试获取 ` monitor`                    |
| 的所有权，过程如下：  * 如果 ` monitor` 的进入数为0，则该线程进入 ` monitor` ，然后将进入数设置为1，该线程即为 ` monitor` 的所有者。 * 如果线程已经占有该 ` monitor` ，只是重新进入，则进入 ` monitor` 的进入数加1. * 如果其他线程已经占用了 ` monitor` ，则该线程进入阻塞状态，直到          |
| ` monitor` 的进入数为0，再重新尝试获取 ` monitor` 的所有权。  #### **monitorexit：** ####  执行 ` monitorexit` 的线程必须是 ` objectref` 所对应的 ` monitor` 的所有者。 指令执行时， ` monitor` 的进入数减1，如果减1后进入数为0，那线程退出 ` monitor` ，不再是这个 `                         |
| monitor` 的所有者。其他被这个 ` monitor` 阻塞的线程可以尝试去获取这个 ` monitor` 的所有权。  >  >  >  > ` synchronized` 的语义底层是通过一个 ` monitor` 的对象来完成，其实 ` wait/notify` 等方法也依赖于 ` > monitor` 对象，这就是为什么只有在同步的块或者方法中才能调用                      |
| ` wait/notify` 等方法，否则会抛出 ` > java.lang.IllegalMonitorStateException` 的异常的原因。 >  >  ### **好处和弊端** ###  **好处** ：解决了线程的安全问题。  **弊端** ：相对降低了效率，因为同步外的线程的都会判断同步锁。获得锁和释放锁带来性能消耗。                                       |
| ### **编译器对synchronized优化** ###  Java6 为了减少获得锁和释放锁所带来的性能消耗，引入了“偏向锁”和“轻量级锁”，所以在Java6 里锁一共有四种状态： **无锁状态，偏向锁状态，轻量级锁状态和重量级锁状态** ，它会随着竞争情况逐渐升级。锁可以升级但不能降级。                                  |
|  *   **偏向锁** ：大多数情况下锁不仅不存在多线程竞争，而且总是由同一线程多次获得。偏向锁的目的是在某个线程获得锁之后（线程的id会记录在对象的 ` Mark Wod` 中），消除这个线程锁重入（CAS）的开销，看起来让这个线程得到了偏护。  *   **轻量级锁（CAS）**                                         |
| ：轻量级锁是由偏向锁升级来的，偏向锁运行在一个线程进入同步块的情况下，当第二个线程加入锁争用的时候，偏向锁就会升级为轻量级锁；轻量级锁的意图是在没有多线程竞争的情况下，通过CAS操作尝试将MarkWord更新为指向LockRecord的指针，减少了使用重量级锁的系统互斥量产生的性能消耗。  *                |
| **重量级锁** ：虚拟机使用CAS操作尝试将MarkWord更新为指向LockRecord的指针，如果更新成功表示线程就拥有该对象的锁；如果失败，会检查MarkWord是否指向当前线程的栈帧，如果是，表示当前线程已经拥有这个锁；如果不是，说明这个锁被其他线程抢占，此时膨胀为重量级锁。  ### **锁状态对应的Mark          |
| Word** ###  以32位JVM为例：  +------------+------------------------------+--------------+--------------+------+----+ |            |                           25 |              |              |      | | **锁状态** | bit                          | 4bit         | 1bit                     |
|       | 2bit | |            |                              |              |              |      | |            |                              |              |              | | 23bit      | 2bit                         | 是否是偏向锁 | 锁标志位     | |            |                      |
|                           |              |              | |            |                              |              | | 轻量级锁   | 指向栈中锁记录的指针         |           00 | |            |                              |              | |            |                               |
|                 |              | | 重量级锁   | 指向互斥量（重量级锁）的指针 |           10 | |            |                              |              | |  GC        |                              |              | | 标记       | 空                           |                         |
|    11 | |            |                              |              | |            |  线程                        |              |              |      |    | | 偏向锁     | ID                           | Epoch        | 对象分代年龄 |    1 | 01 | |            |                           |
|                     |              |              |      |    | +------------+------------------------------+--------------+--------------+------+----+  ## **五、volatile** ##  **` volatile` 是Java中的一个关键字，用来修饰共享变量（类的成员变量、类的静态成员变量）。**                   |
| **被修饰的变量包含两层语义：**  * **保证可见性**  >  >  >  > 线程写入变量时不会把变量写入缓存，而是直接把值刷新回主存。同时，其他线程在读取该共享变量的时候，会从主内存重新获取值，而不是使用当前缓存中的值。（因此会带来一部分性能损失）。 > **注意：往主内存中写入的操作不能保证原子性。**  |
| >  >  * **禁止指令重排**  >  >  >  > 禁止指令重排序有两层意思：  1）当程序执行到 ` volatile` 变量的读操作或者写操作时，在其前面的操作的更改肯定全部已经进行，且结果已经对后面的操作可见；在其后面的操作肯定还没有进行； > 2）在进行指令优化时，不能将在对 ` volatile`                         |
| 变量访问的语句放在其后面执行，也不能把 ` volatile` 变量后面的语句放到其前面执行。 >  >  >  **底层实现：**观察加入 ` volatile` 关键字和没有加入 ` volatile` 关键字时所生成的汇编代码发现，加入 ` volatile` 关键字时，会多出一个 ` lock前缀指令` 。  ## **六、Lock** ##                         |
|  ### **应用场景** ###  如果一个代码块被 ` synchronized` 修饰了，当一个线程获取了对应的锁，并执行该代码块时，其他线程便只能一直等待，等待获取锁的线程释放锁，而这里获取锁的线程释放锁只会有 **两种情况** ：  * 获取锁的线程执行完了该代码块，然后线程释放对锁的占有； *                        |
| 线程执行发生异常，此时JVM会让线程自动释放锁。  如果这个获取锁的线程由于要等待IO或者其他原因（比如调用sleep方法）被阻塞了，但是又没有释放锁，会让程序效率很差。  **因此就需要有一种机制可以不让等待的线程一直无期限地等待下去（比如只等待一定的时间或者能够响应中断），通过                    |
| ` Lock` 就可以办到。**  ### **源码分析** ###  与Lock相关的接口和类位于 ` J.U.C` 的 ` java.util.concurrent.locks` 包下。 ![此处输入图片的描述](https://user-gold-cdn.xitu.io/2019/4/9/16a024f012365b32?imageView2/0/w/1280/h/960/ignore-error/1)   #### (1) **Lock接口**                       |
| ####  ` public interface Lock { void lock () ; void lockInterruptibly () throws InterruptedException ; boolean tryLock () ; boolean tryLock ( long time, TimeUnit unit) throws InterruptedException ; void unlock () ; Condition newCondition () ; } 复制代码` * **获取锁**                   |
| **lock()** ：获取锁，如果锁被暂用则一直等待。 **tryLock()** : 有返回值的获取锁。注意返回类型是 ` boolean` ，如果获取锁的时候锁被占用就返回 ` false` ，否则返回 ` true` 。 **tryLock(long time, TimeUnit unit)** ：比起tryLock()就是给了一个时间期限，保证等待参数时间。                       |
| **lockInterruptibly()** ：当通过这个方法去获取锁时，如果线程正在等待获取锁，则这个线程能够响应中断，即中断线程的等待状态。也就使说，当两个线程同时通过 ` lock.lockInterruptibly()` 想获取某个锁时，假若此时线程A获取到了锁，而线程B只有在等待，那么对线程B调用 ` threadB.interrupt()`         |
| 方法能够中断线程B的等待过程。  >  >  >  > **注意** ：当一个线程获取了锁之后，是不会被 ` interrupt()` 方法中断的。因为本身在前面的文章中讲过单独调用 ` > interrupt()` 方法不能中断正在运行过程中的线程，只能中断阻塞过程中的线程。因此当通过 ` lockInterruptibly()`                            |
| 方法获取某个锁时，如果不能获取到，只有进行等待的情况下，是可以响应中断的。用 > ` synchronized` 修饰的话，当一个线程处于等待某个锁的状态，是无法被中断的，只有一直等待下去。 >  >  * **释放锁** **unlock()** :释放锁。  #### (2) **ReentrantLock类** ####  ` ReentrantLock`                    |
| ，意思是“可重入锁”。 ` ReentrantLock` 是唯一实现了 ` Lock` 接口的类，并且 ` ReentrantLock` 提供了更多的方法，基于 ` AQS(AbstractQueuedSynchronizer)` 来实现的。  >  >  >  > 并且， ` ConcurrentHashMap` 并没有采用 ` synchronized` 进行控制，而是使用了 ` ReentrantLock` > 。 >  >          |
| * **构造方法** ` ReentrantLock` 分为 **公平锁** 和 **非公平锁** ，可以通过构造方法来指定具体类型：  ` public ReentrantLock () { sync = new NonfairSync(); } public ReentrantLock ( boolean fair) { sync = fair ? new FairSync() : new NonfairSync(); } 复制代码` * **获取锁**  ` public       |
| void lock () { sync.lock(); } 复制代码` 而 ` sync` 是一个 ` abstract` 内部类：  ` abstract static class Sync extends AbstractQueuedSynchronizer { private static final long serialVersionUID = -5179523762034025860L; abstract void lock(); 复制代码` 其 ` lock()` 方法用的是构造得到的       |
| ` FairSync` 对象，即 ` sync` 的实现类。  ` public ReentrantLock () { sync = new NonfairSync(); } //删去一些方法 static final class NonfairSync extends Sync { final void lock () { if (compareAndSetState( 0 , 1 )) setExclusiveOwnerThread(Thread.currentThread()); else acquire( 1 ); }     |
| protected final boolean tryAcquire ( int acquires) { return nonfairTryAcquire(acquires); } } 复制代码` 而 ` compareAndSetState` 是 ` AQS` 的一个方法，也就是基于 ` CAS` 操作。  ` public final void acquire ( int arg) { if (!tryAcquire(arg) && acquireQueued(addWaiter(Node.EXCLUSIVE),     |
| arg)) selfInterrupt(); } 复制代码` 尝试进一步获取锁（调用继承自父类 ` sync` 的 ` final` 方法）：  ` final boolean nonfairTryAcquire ( int acquires) { final Thread current = Thread.currentThread(); int c = getState(); if (c == 0 ) { if (compareAndSetState( 0 , acquires))                |
| { setExclusiveOwnerThread(current); return true ; } } else if (current == getExclusiveOwnerThread()) { int nextc = c + acquires; if (nextc < 0 ) // overflow throw new Error( "Maximum lock count exceeded" ); setState(nextc); return true ; } return false ; } 复制代码`                    |
| 首先会判断 ` AQS` 中的 ` state` 是否等于 0，0表示目前没有其他线程获得锁，当前线程就可以尝试获取锁。如果 ` state` 大于 0 时，说明锁已经被获取了，则需要判断获取锁的线程是否为当前线程( ` ReentrantLock` 支持重入)，是则需要将 state + 1，并将值更新。  如果 ` tryAcquire(arg)`                 |
| 获取锁失败，则需要用 ` addWaiter(Node.EXCLUSIVE)` 将当前线程写入队列中。写入之前需要将当前线程包装为一个 ` Node` 对象 ` (addWaiter(Node.EXCLUSIVE))` 。  即回到：  ` public final void acquire ( int arg) { if (!tryAcquire(arg) && acquireQueued(addWaiter(Node.EXCLUSIVE),                  |
| arg)) selfInterrupt(); } 复制代码` * **释放锁**  ` 公平锁和非公平锁的释放流程都是一样的： public void unlock () { sync.release( 1 ); } public final boolean release ( int arg) { if (tryRelease(arg)) { Node h = head; if (h != null && h.waitStatus != 0 ) //唤醒被挂起的线程                |
| unparkSuccessor(h); return true ; } return false ; } //尝试释放锁 protected final boolean tryRelease ( int releases) { int c = getState() - releases; if (Thread.currentThread() != getExclusiveOwnerThread()) throw new IllegalMonitorStateException(); boolean free = false ;               |
| if (c == 0 ) { free = true ; setExclusiveOwnerThread( null ); } setState(c); return free; } 复制代码` #### (3) **ReadWriteLock接口** 与 **ReentrantReadWriteLock类** ####  * **定义**  ` public interface ReadWriteLock { Lock readLock () ; Lock writeLock () ; } 复制代码` 在 `             |
| ReentrantLock` 中，线程之间的同步都是互斥的，不管是读操作还是写操作，但是在一些场景中读操作是可以并行进行的，只有写操作才是互斥的，这种情况虽然也可以使用 ` ReentrantLock` 来解决，但是在性能上也会损失， ` ReadWriteLock` 就是用来解决这个问题的。  * **实现-ReentrantReadWriteLock类**      |
|  在 ` ReentrantReadWriteLock` 中分别定义了读锁和写锁，与 ` ReentrantLock` 类似，读锁和写锁的功能也是通过 ` Sync` 实现的， ` Sync` 存在 **公平和非公平** 两种实现方式，不同的是表示锁状态的 ` state` 的定义，在 ` ReentrantReadWriteLock` 中具体定义如下：  ` static final int                 |
| SHARED_SHIFT = 16 ; static final int SHARED_UNIT = ( 1 << SHARED_SHIFT); static final int MAX_COUNT = ( 1 << SHARED_SHIFT) - 1 ; static final int EXCLUSIVE_MASK = ( 1 << SHARED_SHIFT) - 1 ; //获取读锁的占有次数 static int sharedCount ( int c) { return c >>> SHARED_SHIFT; }             |
| //获取写锁的占有次数 static int exclusiveCount ( int c) { return c & EXCLUSIVE_MASK; } //线程的id和对应线程获取的读锁的数量 static final class HoldCounter { int count = 0 ; // Use id, not reference, to avoid garbage retention final long tid = Thread.currentThread().getId();            |
| } //线程变量保存线程和线程中获取的读写的数量 static final class ThreadLocalHoldCounter extends ThreadLocal < HoldCounter > { public HoldCounter initialValue () { return new HoldCounter(); } } private transient ThreadLocalHoldCounter readHolds; //缓存最后一个获取读锁的线程              |
| private transient HoldCounter cachedHoldCounter; //保存第一个获取读锁的线程 private transient Thread firstReader = null ; private transient int firstReaderHoldCount; 复制代码` 其中，包含两个静态内部类： ` ReadLock()` 与 ` WriteLock()` ,都实现了 ` Lock接口` 。  **获取读锁** ：          |
| * 如果不存在线程持有写锁，则获取读锁成功。 * 如果其他线程持有写锁，则获取读锁失败。 * 如本线程持有写锁，并且不存在等待写锁的其他线程，则获取读锁成功。 * 如本线程持有写锁，并且存在等待写锁的其他线程，则如果本线程已经持有读锁，则获取读锁成功，如果不能存在读锁，则此次获取读锁失败。       |
|  **获取写锁** ：  * 判断是否有线程持有锁，包括读锁和写锁，如果有，则执行步骤2，否则步骤3 * 如果写锁为空(此时由于1步骤判断存在锁，则存在持有读锁的线程)，或者持有写锁的不是本线程,直接返回失败，如果写锁数量大于MAX_COUNT，返回失败，否则更新state，并且返回true *                             |
| 如果需要写锁堵塞判断，或者CAS失败直接返回false，否则设置持有写锁的线程为本线程，并且返回true * 通过writerShouldBlock写锁堵塞判断  ` final boolean writerShouldBlock () { return hasQueuedPredecessors(); } //判断是否堵塞 public final boolean hasQueuedPredecessors                          |
| () { Node t = tail; Node h = head; Node s; return h != t && ((s = h.next) == null || s.thread != Thread.currentThread()); } 复制代码` ## **七、比较** ##  ### **Lock和synchronized** ###  ` synchronized` 是基于JVM层面实现的，而Lock是基于JDK层面实现的。                                    |
| ` Lock` 需要 ` lock` 和 ` release` ，比 ` synchronized` 复杂，但 ` Lock` 可以做更细粒度的锁，支持获取超时、获取中断，这是 ` synchronized` 所不具备的。Lock的实现主要有 ` ReentrantLock` 、 ` ReadLock` 和 ` WriteLock` ,读读共享，写写互斥，读写互斥。  *                                     |
| Lock是一个 **接口** ，而synchronized是Java中的 **关键字** ，synchronized是内置的语言实现；  *   synchronized在发生异常时，会 **自动释放** 线程占有的锁，因此不会导致死锁现象发生；而Lock在发生异常时，如果没有主动通过unLock()去释放锁，则很可能造成 **死锁现象**                             |
| ，因此使用Lock时需要在finally块中释放锁；  *   Lock可以让等待锁的线程 **响应中断** ，而synchronized却不行，使用synchronized时，等待的线程会一直等待下去，不能够响应中断；  *   通过Lock可以知道有没有 **成功获取锁** ，而synchronized却无法办到。  *   Lock可以提高多个线程进行读操作的       |
| **效率** 。  *   Lock实现和synchronized不一样，后者是一种 **悲观锁** ，它胆子很小，它很怕有人和它抢吃的，所以它每次吃东西前都把自己关起来。而Lock底层其实是CAS **乐观锁** 的体现，它无所谓，别人抢了它吃的，它重新去拿吃的就好啦，所以它很乐观。底层主要靠                                    |
| ` volatile` 和 ` CAS` 操作实现的。  ### **synchronized和volatile** ###  *   volatile本质是在告诉jvm当前变量在寄存器（工作内存）中的值是不确定的，需要从主存中读取；  *   synchronized则是锁定当前变量，只有当前线程可以访问该变量，其他线程被阻塞住。  *                                      |
| volatile仅能使用在变量级别；synchronized则可以使用在变量、方法、和类级别的  *   volatile仅能实现变量的修改可见性，不能保证原子性；而synchronized则可以保证变量的修改可见性和原子性  *   volatile不会造成线程的阻塞；synchronized可能会造成线程的阻塞。  *                                     |
| volatile标记的变量不会被编译器优化；synchronized标记的变量可以被编译器优化  ## **七、死锁问题** ##  死锁有四个必要条件，打破一个即可去除死锁。  ### **四个必要条件：** ###  * 互斥条件  >  >  >  > 一个资源每次只能被一个进程使用。 >  >  * 请求与保持条件  >                                 |
|  >  >  > 一个线程因请求资源而阻塞时，对已获得的资源保持不放。 >  >  * 不剥夺条件  >  >  >  > 线程已获得的资源，在末使用完之前，不能强行剥夺。 >  >  * 循环等待条件  >  >  >  > 若干线程之间形成一种头尾相接的循环等待资源关系。 >  >  ### **死锁的例子** ###                                  |
| 同步嵌套时，两个线程互相锁住，都不释放，造成死锁。 **举例：** 创建两个字符串a和b，再创建两个线程A和B，让每个线程都用synchronized锁住字符串（A先锁a，再去锁b；B先锁b，再锁a），如果A锁住a，B锁住b，A就没办法锁住b，B也没办法锁住a，这时就陷入了死锁。  ` public                                |
| class DeadLock { public static String obj1 = "obj1" ; public static String obj2 = "obj2" ; public static void main (String[] args) { Thread a = new Thread( new Lock1()); Thread b = new Thread( new Lock2()); a.start(); b.start(); } } class Lock1 implements                               |
| Runnable { @Override public void run () { try { System.out.println( "Lock1 running" ); while ( true ){ synchronized (DeadLock.obj1){ System.out.println( "Lock1 lock obj1" ); Thread.sleep( 3000 ); //获取obj1后先等一会儿，让Lock2有足够的时间锁住obj2 synchronized                          |
| (DeadLock.obj2){ System.out.println( "Lock1 lock obj2" ); } } } } catch (Exception e){ e.printStackTrace(); } } } class Lock2 implements Runnable { @Override public void run () { try { System.out.println( "Lock2 running" ); while ( true ){ synchronized                                  |
| (DeadLock.obj2){ System.out.println( "Lock2 lock obj2" ); Thread.sleep( 3000 ); synchronized (DeadLock.obj1){ System.out.println( "Lock2 lock obj1" ); } } } } catch (Exception e){ e.printStackTrace(); } } } 复制代码` ## **八、锁的概念** ##  **在 java                                    |
| 中锁的实现主要有两类：内部锁 ` synchronized` （对象内置的monitor锁）和显示锁 ` java.util.concurrent.locks.Lock` 。**  * **可重入锁**  >  >  >  > 指的是同一线程外层函数获得锁之后 ，内层递归函数仍然有获取该锁的代码，但不受影响，执行对象中所有同步方法不用再次获得锁。                      |
| ` > synchronized` 和 ` Lock` 都具备可重入性。 >  >  * **可中断锁**  >  >  >  > ` synchronized` 就不是可中断锁，而Lock是可中断锁。 >  >  * **公平锁**  >  >  >  > 按等待获取锁的线程的等待时间进行获取， **等待时间长** 的具有优先获取锁权利。 `                                               |
| synchronized` 就是非公平锁；对于 ` > ReentrantLock` 和 ` ReentrantReadWriteLock` ，它默认情况下是非公平锁，但是可以设置为公平锁。 >  >  * **读写锁**  >  >  >  > 对资源读取和写入的时候拆分为2部分处理，读的时候可以多线程一起读，写的时候必须同步地写。                                      |
| ` ReadWriteLock` 就是读写锁，它是一个接口， > ` ReentrantReadWriteLock` 实现了这个接口。 >  >  * **自旋锁**  >  >  >  > 让线程去执行一个无意义的循环，循环结束后再去重新竞争锁，如果竞争不到继续循环，循环过程中线程会一直处于 ` running`                                                     |
| 状态，但是基于JVM的线程调度，会让出时间片，所以其他线程依旧有申请锁和释放锁的机会。自旋锁省去了阻塞锁的时间空间（队列的维护等）开销，但是长时间自旋就变成了“忙式等待”，忙式等待显然还不如阻塞锁。所以自旋的次数一般控制在一个范围内，例如10,100等，在超出这个范围后，自旋锁会升级为阻塞锁。 |
| >  >  >  * **独占锁**  >  >  >  > 是一种 **悲观锁** ， ` synchronized` 就是一种独占锁，会导致其它所有需要锁的线程挂起，等待持有锁的线程释放锁。 >  >  * **乐观锁**  >  >  >  > 每次不加锁，假设没有冲突去完成某项操作，如果因为冲突失败就重试，直到成功为止。                                 |
| >  >  * **悲观锁**  >  >  >  > 导致其它所有需要锁的线程挂起，等待持有锁的线程释放锁。 >  >  ### **关于JUC** ###  包含了两个子包：atomic以及lock，另外在concurrent下的阻塞队列以及executors，以后再深入学习吧，下面这个图很是经典：                                                            |
| ![此处输入图片的描述](https://user-gold-cdn.xitu.io/2018/5/3/163260cff7cb847c?imageView2/0/w/1280/h/960/ignore-error/1)     **参考链接**  *   [Java并发编程：volatile关键字解析]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3920373.html )  *           |
| [Java SE1.6 中的 Synchronized]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3923167.html )  *   [Java并发编程：Lock]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3923167.html )  *   [Java并发机制的底层实现](       |
| https://juejin.im/post/5a067fa251882578d84eee8f )  *   [Java死锁]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fmudao%2Fp%2F5867107.html )  *   [深入理解JVM]( https://link.juejin.im?target=https%3A%2F%2Fbook.douban.com%2Fsubject%2F6522893%2F )  *                       |
| [初识Lock与AbstractQueuedSynchronizer(AQS)]( https://juejin.im/post/5aeb055b6fb9a07abf725c8c )                                                                                                                                                                                                |
+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------+-------+------+------+------+------+