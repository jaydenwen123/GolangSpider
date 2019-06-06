# java并发编程系列：牛逼的AQS（下） #

标签： 「我们都是小青蛙」公众号文章

## Condition ##

### ReentrantLock的内部实现 ###

看完了 ` AQS` 中的底层同步机制，我们来简单分析一下之前介绍过的 ` ReentrantLock` 的实现原理。先回顾一下这个显式锁的典型使用方式：

` Lock lock = new ReentrantLock(); lock.lock(); try { 加锁后的代码 } finally { lock.unlock(); } 复制代码`

` ReentrantLock` 首先是一个显式锁，它实现了 ` Lock` 接口。可能你已经忘记了 ` Lock` 接口长啥样了，我们再回顾一遍：

` public interface Lock { void lock () ; void lockInterruptibly () throws InterruptedException ; boolean tryLock () ; boolean tryLock ( long time, TimeUnit unit) throws InterruptedException ; void unlock () ; Condition newCondition () ; } 复制代码`

其实 ` ReentrantLock` 内部定义了一个 ` AQS` 的子类来辅助它实现锁的功能，由于 ` ReentrantLock` 是工作在 ` 独占模式` 下的，所以它的 ` lock` 方法其实是调用 ` AQS` 对象的 ` aquire` 方法去获取同步状态， ` unlock` 方法其实是调用 ` AQS` 对象的 ` release` 方法去释放同步状态，这些大家已经很熟了，就不再赘述了，我们大致看一下 ` ReentrantLock` 的代码：

` public class ReentrantLock implements Lock { private final Sync sync; //AQS子类对象 abstract static class Sync extends AbstractQueuedSynchronizer { // ... 为节省篇幅，省略其他内容 } // ... 为节省篇幅，省略其他内容 } 复制代码`

所以如果我们简简单单写下下边这行代码：

` Lock lock = new ReentrantLock(); 复制代码`

就意味着在内存里创建了一个 ` ReentrantLock` 对象，一个 ` AQS` 对象，在 ` AQS` 对象里维护着 ` 同步队列` 的 ` head` 节点和 ` tail` 节点，不过初始状态下由于没有线程去竞争锁，所以 ` 同步队列` 是空的，画成图就是这样：

![image_1c3hf30h3bmodidrvogfe11oh2q.png-16.3kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab037f6fda14e0?imageView2/0/w/1280/h/960/ignore-error/1)

### Condition的提出 ###

我们前边唠叨线程间通信的时候提到过内置锁的 ` wait/notify` 机制， ` 等待线程` 的典型的代码如下：

` synchronized (对象) { 处理逻辑（可选） while (条件不满足) { 对象.wait(); } 处理逻辑（可选） } 复制代码`

通知线程的典型的代码如下：

` synchronized (对象) { 完成条件 对象.notifyAll();、 } 复制代码`

也就是当一个线程因为某个条件不能满足时就可以在持有锁的情况下调用该锁对象的 ` wait` 方法，之后该线程会释放锁并进入到与该锁对象关联的等待队列中等待；如果某个线程完成了该等待条件，那么在持有相同锁的情况下调用该锁的 ` notify` 或者 ` notifyAll` 方法唤醒在与该锁对象关联的等待队列中等待的线程。

显式锁的本质其实是通过 ` AQS` 对象获取和释放同步状态，而内置锁的实现是被封装在java虚拟机里的，我们并没有讲过，这两者的实现是不一样的 。而 ` wait/notify` 机制只适用于内置锁，在 ` 显式锁` 里需要另外定义一套类似的机制，在我们定义这个机制的时候需要整清楚： 在获取锁的线程因为某个条件不满足时，应该进入哪个等待队列，在什么时候释放锁，如果某个线程完成了该等待条件，那么在持有相同锁的情况下怎么从相应的等待队列中将等待的线程从队列中移出 。

为了定义这个等待队列，设计java的大叔们在 ` AQS` 中添加了一个名叫 ` ConditionObject` 的成员内部类：

` public abstract class AbstractQueuedSynchronizer { public class ConditionObject implements Condition , java. io. Serializable { private transient Node firstWaiter; private transient Node lastWaiter; // ... 为省略篇幅，省略其他方法 } } 复制代码`

很显然，这个 ` ConditionObject` 维护了一个队列， ` firstWaiter` 是队列的头节点引用， ` lastWaiter` 是队列的尾节点引用。但是节点类是 ` Node` ？对，你没看错，就是我们前边分析的 ` 同步队列` 里用到的 ` AQS` 的静态内部类 ` Node` ，怕你忘了，再把这个 ` Node` 节点类的主要内容写一遍：

` static final class Node { volatile int waitStatus; volatile Node prev; volatile Node next; volatile Thread thread; Node nextWaiter; static final int CANCELLED = 1 ; static final int SIGNAL = - 1 ; static final int CONDITION = - 2 ; static final int PROPAGATE = - 3 ; } 复制代码`

也就是说： ` AQS` 中的同步队列和自定义的等待队列使用的节点类是同一个 。

又由于 在等待队列中的线程被唤醒的时候需要重新获取锁，也就是重新获取同步状态，所以该等待队列必须知道线程是在持有哪个锁的时候开始等待的 。设计java的大叔们在 ` Lock` 接口中提供了这么一个通过锁来获取等待队列的方法：

` Condition newCondition () ; 复制代码`

我们上边介绍的 ` ConditionObject` 就实现了 ` Condition` 接口，看一下 ` ReentrantLock` 锁是怎么获取与它相关的等待队列的：

` public class ReentrantLock implements Lock { private final Sync sync; abstract static class Sync extends AbstractQueuedSynchronizer { final ConditionObject newCondition () { return new ConditionObject(); } // ... 为节省篇幅，省略其他方法 } public Condition newCondition () { return sync.newCondition(); } // ... 为节省篇幅，省略其他方法 } 复制代码`

可以看到，其实就是简单创建了一个 ` ConditionObject` 对象而已～ 由于 ConditionObject 是AQS 的成员内部类，所以在创建的 ConditionObject 对象中持有 AQS 对象的引用，所以通过 ConditionObject 对象访问到 同步队列，也就是可以重新获取同步状态，也就是重新获取锁 。用文字描述还是有些绕，我们先通过锁来创建一个 ` Condition` 对象：

` Lock lock = new ReentrantLock(); Condition condition = lock.newCondition(); 复制代码`

由于在初始状态下，没有线程去竞争锁，所以 ` 同步队列` 是空的，也没有线程因某个条件不成立而进入等待队列，所以 ` 等待队列` 也是空的， ` ReentrantLock` 对象、 ` AQS` 对象以及等待队列在内存中的表示就如图：

![image_1c3hji5a4uvn1sdj1ro33hm87m61.png-26.7kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab037f6cf5c004?imageView2/0/w/1280/h/960/ignore-error/1)

当然，这个 ` newCondition` 方法可以反复调用，从而可以通过一个锁来生成多个 ` 等待队列` ：

` Lock lock = new ReentrantLock(); Condition condition1 = lock.newCondition(); Condition condition2 = lock.newCondition(); 复制代码`

那接下来需要考虑怎么把线程包装成 ` Node` 节点放到等待队列的以及怎么从等待队列中移出了。 ` ConditionObject` 成员内部类实现了一个 ` Condition` 的接口，这个接口提供了下边这些方法：

` public interface Condition { void await () throws InterruptedException ; long awaitNanos ( long nanosTimeout) throws InterruptedException ; boolean await ( long time, TimeUnit unit) throws InterruptedException ; boolean awaitUntil (Date deadline) throws InterruptedException ; void awaitUninterruptibly () ; void signal () ; void signalAll () ; } 复制代码`

来看一下这些方法的具体意思：

+--------------------------------+--------------------------------------------------------------------------------------+
|             方法名             |                                         描述                                         |
+--------------------------------+--------------------------------------------------------------------------------------+
| ` void await()`                | 当前线程进入等待状态，直到被通知(调用signal或者signalAll方法)或中断                  |
| ` boolean await(long time,     | 当前线程在指定时间内进入等待状态，如果超出指定时间或者在等待状态中被通知或中断则返回 |
| TimeUnit unit)`                |                                                                                      |
| ` long awaitNanos(long         | 与上个方法相同，只不过默认使用的时间单位为纳秒                                       |
| nanosTimeout)`                 |                                                                                      |
| ` boolean awaitUntil(Date      | 当前线程进入等待状态，如果到达最后期限或者在等待状态中被通知或中断则返回             |
| deadline)`                     |                                                                                      |
| ` void awaitUninterruptibly()` | 当前线程进入等待状态，直到在等待状态中被通知，需要注意的时，本方法并不相应中断       |
| ` void signal()`               | 唤醒一个等待线程。                                                                   |
| ` void signalAll()`            | 唤醒所有等待线程。                                                                   |
+--------------------------------+--------------------------------------------------------------------------------------+

可以看到， ` Condition` 中的 ` await` 方法和内置锁对象的 ` wait` 方法的作用是一样的，都会使当前线程进入等待状态， ` signal` 方法和内置锁对象的 ` notify` 方法的作用是一样的，都会唤醒在等待队列中的线程。

像调用内置锁的 ` wait/notify` 方法时，线程需要首先获取该锁一样，调用 ` Condition` 对象的 ` await/siganl` 方法的线程需要首先获得产生该 ` Condition` 对象的显式锁 。它的基本使用方式就是： 通过显式锁的 newCondition 方法产生 ` Condition` 对象，线程在持有该显式锁的情况下可以调用生成的 ` Condition` 对象的 await/signal 方法 ，一般用法如下：

` Lock lock = new ReentrantLock(); Condition condition = lock.newCondition(); //等待线程的典型模式 public void conditionAWait () throws InterruptedException { lock.lock(); //获取锁 try { while (条件不满足) { condition.await(); //使线程处于等待状态 } 条件满足后执行的代码; } finally { lock.unlock(); //释放锁 } } //通知线程的典型模式 public void conditionSignal () throws InterruptedException { lock.lock(); //获取锁 try { 完成条件; condition.signalAll(); //唤醒处于等待状态的线程 } finally { lock.unlock(); //释放锁 } } 复制代码`

假设现在有一个锁和两个等待队列：

` Lock lock = new ReentrantLock(); Condition condition1 = lock.newCondition(); Condition condition2 = lock.newCondition(); 复制代码`

画图表示出来就是：

![image_1c3hjlt3n14rl1i9g1epg3371ce16e.png-39.7kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab037f73f8a1fb?imageView2/0/w/1280/h/960/ignore-error/1)

有3个线程 ` main` 、 ` t1` 、 ` t2` 同时调用 ` ReentrantLock` 对象的 ` lock` 方法去竞争锁的话，只有线程 ` main` 获取到了锁，所以会把线程 ` t1` 、 ` t2` 包装成 ` Node` 节点插入 ` 同步队列` ，所以 ` ReentrantLock` 对象、 ` AQS` 对象和 ` 同步队列` 的示意图就是这样的：

![image_1c3hjmnj819nl57f11l11p7f9196r.png-94.5kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab037f74b59eea?imageView2/0/w/1280/h/960/ignore-error/1)

因为此时 ` main` 线程是获取到锁处于运行中状态，但是因为某个条件不满足，所以它选择执行下边的代码来进入 ` condition1` 等待队列：

` lock.lock(); try { contition1.await(); } finally { lock.unlock(); } 复制代码`

具体的 ` await` 代码我们就不分析了，太长了，我怕你看的发困，这里只看这个 ` await` 方法做了什么事情：

* 

在 ` condition1` 等待队列中创建一个 ` Node` 节点，这个节点的 ` thread` 值就是 ` main` 线程，而且 ` waitStatus` 为 ` -2` ，也就是静态变量 ` Node.CONDITION` ，表示表示节点在等待队列中，由于这个节点是代表线程 ` main` 的，所以就把它叫做 ` main节点` 把，新创建的节点长这样：

![image_1c3hs5hvfu3g186m1g8m9tjdg9cg.png-14.1kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab037f6ccb6a98?imageView2/0/w/1280/h/960/ignore-error/1)

* 

将该节点插入 ` condition1` 等待队列中：

![image_1c3hs74l64m11n6eedc357acvct.png-118.9kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab037f7389296e?imageView2/0/w/1280/h/960/ignore-error/1)

* 

因为 ` main` 线程还持有者锁，所以需要释放锁之后通知后边等待获取锁的线程 ` t` ，所以 ` 同步队列` 里的0号节点被删除，线程 ` t` 获取锁， ` 节点1` 称为 ` head` 节点，并且把 ` thread` 字段设置为null：

![image_1c3hs8phe1r1smrl12q41231hfuda.png-103.4kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab03805f9597a0?imageView2/0/w/1280/h/960/ignore-error/1)

至此， ` main` 线程的等待操作就做完了，假如现在获得锁的 ` t1` 线程也执行下边的代码：

` lock.lock(); try { contition1.await(); } finally { lock.unlock(); } 复制代码`

还是会执行上边的过程，把 ` t1` 线程包装成 ` Node` 节点插入到 ` condition1` 等待队列中去，由于原来在等待队列中的 ` 节点1` 会被删除，我们把这个新插入等待队列代表线程 ` t1` 的节点称为 ` 新节点1` 吧：

![image_1c3hshhsik77531ribb6kv57e4.png-112.2kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab03806895a6ff?imageView2/0/w/1280/h/960/ignore-error/1)

这里需要特别注意的是： 同步队列是一个双向链表，prev表示前一个节点，next表示后一个节点，而等待队列是一个单向链表，使用nextWaiter表示下一个节点，这是它们不同的地方 。

现在获取到锁的线程是 ` t2` ，大家一起出来混的，前两个都进去，只剩下 ` t2` 多不好呀，不过这次不放在 ` condition1` 队列后头了，换成 ` condition2` 队列吧：

` lock.lock(); try { contition2.await(); } finally { lock.unlock(); } 复制代码`

效果就是：

![image_1c3hsjumr5jb3c5cqhdk57tieh.png-127.6kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab0380960f5c95?imageView2/0/w/1280/h/960/ignore-error/1)

大家发现，虽然现在没有线程获取锁，也没有线程在锁上等待，但是 ` 同步队列` 里仍旧有一个节点，是的， 同步队列只有初始时无任何线程因为锁而阻塞的时候才为空，只要曾经有线程因为获取不到锁而阻塞，这个队列就不为空了 。

至此， ` main` 、 ` t1` 和 ` t2` 这三个线程都进入到等待状态了，都进去了谁把它们弄出来呢？？？额～ 好吧，再弄一个别的线程去获取同一个锁，比方说线程 ` t3` 去把 ` condition2` 条件队列的线程去唤醒，可以调用这个 ` signal` 方法：

` lock.lock(); try { contition2.signal(); } finally { lock.unlock(); } 复制代码`

因为在 ` condition2` 等待队列的线程只有 ` t2` ，所以 ` t2` 会被唤醒，这个过程分两步进行：

* 

将在 ` condition2` 等待队列的代表线程 ` t2` 的 ` 新节点2` ，从等待队列中移出。

* 

将移出的 ` 节点2` 放在同步队列中等待获取锁，同时更改该节点的 ` waitStauts` 为 ` 0` 。

这个过程的图示如下：

![image_1c3hsv52u8i64rgsdo2t1lngeu.png-119.3kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab03808d298efc?imageView2/0/w/1280/h/960/ignore-error/1)

如果线程 ` t3` 继续调用 ` signalAll` 把 ` condition1` 等待队列中的线程给唤醒也是差不多的意思，只不过会把 ` condition1` 上的两个节点同时都移动到同步队列里：

` lock.lock(); try { contition1.signalAll(); } finally { lock.unlock(); } 复制代码`

效果如图：

![image_1c3hthb14i21a58i5168p1a1bfb.png-98.9kB](https://user-gold-cdn.xitu.io/2019/5/13/16ab03809c8de83d?imageView2/0/w/1280/h/960/ignore-error/1)

这样全部线程都从 ` 等待` 状态中恢复了过来，可以重新竞争锁进行下一步操作了。

以上就是 ` Condition` 机制的原理和用法，它其实是内置锁的 ` wait/notify` 机制在显式锁中的另一种实现，不过 原来的一个内置锁对象只能对应一个等待队列，现在一个显式锁可以产生若干个等待队列，我们可以根据线程的不同等待条件来把线程放到不同的等待队列上去 。 ` Condition` 机制的用途可以参考 ` wait/notify` 机制，我们接下来把之前用内置锁和 ` wait/notify` 机制编写的同步队列 ` BlockedQueue` 用 ` 显式锁 + Condition` 的方式来该写一下：

` import java.util.LinkedList; import java.util.Queue; import java.util.concurrent.locks.Condition; import java.util.concurrent.locks.Lock; import java.util.concurrent.locks.ReentrantLock; public class ConditionBlockedQueue < E > { private Lock lock = new ReentrantLock(); private Condition notEmptyCondition = lock.newCondition(); private Condition notFullCondition = lock.newCondition(); private Queue<E> queue = new LinkedList<>(); private int limit; public ConditionBlockedQueue ( int limit) { this.limit = limit; } public int size () { lock.lock(); try { return queue.size(); } finally { lock.unlock(); } } public boolean add (E e) throws InterruptedException { lock.lock(); try { while (size() >= limit) { notFullCondition.await(); } boolean result = queue.add(e); notEmptyCondition.signal(); return result; } finally { lock.unlock(); } } public E remove () throws InterruptedException { lock.lock(); try { while (size() == 0 ) { notEmptyCondition.await(); } E e = queue.remove(); notFullCondition.signalAll(); return e; } finally { lock.unlock(); } } } 复制代码`

在这个队列里边我们用了一个 ` ReentrantLock` 锁，通过这个锁生成了两个 ` Condition` 对象， ` notFullCondition` 表示队列未满的条件， ` notEmptyCondition` 表示队列未空的条件。当队列已满的时候，线程会在 ` notFullCondition` 上等待，每插入一个元素，会通知在 ` notEmptyCondition` 条件上等待的线程；当队列已空的时候，线程会在 ` notEmptyCondition` 上等待，每移除一个元素，会通知在 ` notFullCondition` 条件上等待的线程。这样语义就变得很明显了。 如果你有更多的等待条件，你可以通过显式锁生成更多的 ` Condition` 对象 。而 每个内置锁对象都只能有一个相关联的等待队列，这也是显式锁对内置锁的优势之一 。

我们总结一下上边的用法： 每个显式锁对象又可以产生若干个 ` Condition` 对象，每个 ` Condition` 对象都会对应一个等待队列，所以就起到了一个显式锁对应多个等待队列的效果 。

### ` AQS` 中其他针对等待队列的重要方法 ###

除了 ` Condition` 对象的 ` await` 和 ` signal` 方法， ` AQS` 还提供了许多直接访问这个队列的方法，它们由都是 ` public final` 修饰的：

` public abstract class AbstractQueuedSynchronizer { public final boolean owns (ConditionObject condition) public final boolean hasWaiters (ConditionObject condition) {} public final int getWaitQueueLength (ConditionObject condition) {} public final Collection<Thread> getWaitingThreads (ConditionObject condition) {} } 复制代码`

+-----------------------+----------------------------------------------------------------------------------------------------+
|        方法名         |                                                描述                                                |
+-----------------------+----------------------------------------------------------------------------------------------------+
| ` owns`               | 查询是否通过本 `                                                                                   |
|                       | AQS` 对象生成的指定的                                                                              |
|                       | ConditionObject对象                                                                                |
| ` hasWaiters`         | 指定的等待队列里是否有等待线程                                                                     |
| ` getWaitQueueLength` | 返回正在等待此条件的线程数估计值。因为在构造该结果时，多线程环境下实际线程集合可能发生大的变化     |
| ` getWaitingThreads`  | 返回正在等待此条件的线程集合的估计值。因为在构造该结果时，多线程环境下实际线程集合可能发生大的变化 |
+-----------------------+----------------------------------------------------------------------------------------------------+

如果有需要的话，可以在我们自定义的同步工具中使用它们。

### 题外话 ###

写文章挺累的，有时候你觉得阅读挺流畅的，那其实是背后无数次修改的结果。如果你觉得不错请帮忙转发一下，万分感谢～ 这里是我的公众号「我们都是小青蛙」，里边有更多技术干货，时不时扯一下犊子，欢迎关注：

![](https://user-gold-cdn.xitu.io/2019/4/18/16a2e7131c8643e4?imageView2/0/w/1280/h/960/ignore-error/1)

### 小册 ###

另外，作者还写了一本MySQL小册： [《MySQL是怎样运行的：从根儿上理解MySQL》的链接]( https://juejin.im/book/5bffcbc9f265da614b11b731?referrer=5bff96c6e51d45452f2d6f95 ) 。小册的内容主要是从小白的角度出发，用比较通俗的语言讲解关于MySQL进阶的一些核心概念，比如记录、索引、页面、表空间、查询优化、事务和锁等，总共的字数大约是三四十万字，配有上百幅原创插图。主要是想降低普通程序员学习MySQL进阶的难度，让学习曲线更平滑一点～ 有在MySQL进阶方面有疑惑的同学可以看一下：

![](https://user-gold-cdn.xitu.io/2019/4/17/16a2907eb8386a9b?imageView2/0/w/1280/h/960/ignore-error/1)