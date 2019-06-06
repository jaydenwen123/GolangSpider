# java并发编程系列：wait/notify机制 #

标签： 「我们都是小青蛙」公众号文章

如果一个线程从头到尾执行完也不和别的线程打交道的话，那就不会有各种安全性问题了。但是协作越来越成为社会发展的大势，一个大任务拆成若干个小任务之后，各个小任务之间可能也需要相互协作最终才能执行完整个大任务。所以各个线程在执行过程中可以相互 ` 通信` ，所谓 ` 通信` 就是指相互交换一些数据或者发送一些控制指令，比如一个线程给另一个暂停执行的线程发送一个恢复执行的指令，下边详细看都有哪些通信方式。

### volatile和synchronized ###

可变共享变量是天然的通信媒介，也就是说一个线程如果想和另一个线程通信的话，可以修改某个在多线程间共享的变量，另一个线程通过读取这个共享变量来获取通信的内容。

由于原子性操作、内存可见性和指令重排序的存在，java提供了 ` volatile` 和 ` synchronized` 的同步手段来保证通信内容的正确性，假如没有这些同步手段，一个线程的写入不能被另一个线程立即观测到，那这种通信就是不靠谱的～

### ` wait/notify` 机制 ###

#### 故事背景 ####

也不知道是那个遭天杀的给我们学校厕所的坑里塞了个塑料瓶，导致楼道里如黄河泛滥一般，臭味熏天。更加悲催的是整个楼只有这么一个厕所，比这个更悲催的是这个厕所里只有一个坑！！！！！好吧，让我们用java来描述一下这个厕所：

` public class Washroom { private volatile boolean isAvailable = false ; //表示厕所是否是可用的状态 private Object lock = new Object(); //厕所门的锁 public boolean isAvailable () { return isAvailable; } public void setAvailable ( boolean available) { this.isAvailable = available; } public Object getLock () { return lock; } } 复制代码`

` isAvailable` 字段代表厕所是否可用，由于厕所损坏，默认是 ` false` 的， ` lock` 字段代表这个厕所门的锁。需要注意的是 ` isAvailable` 字段被 ` volatile` 修饰，也就是说有一个线程修改了它的值，它可以立即对别的线程可见～

由于厕所资源宝贵，英明的学校领导立即拟定了一个修复任务：

` public class RepairTask implements Runnable { private Washroom washroom; public RepairTask (Washroom washroom) { this.washroom = washroom; } @Override public void run () { synchronized (washroom.getLock()) { System.out.println( "维修工 获取了厕所的锁" ); System.out.println( "厕所维修中，维修厕所是一件辛苦活，需要很长时间。。。" ); try { Thread.sleep( 5000L ); //用线程sleep表示维修的过程 } catch (InterruptedException e) { throw new RuntimeException(e); } washroom.setAvailable( true ); //维修结束把厕所置为可用状态 System.out.println( "维修工把厕所修好了，准备释放锁了" ); } } } 复制代码`

这个维修计划的内容就是当维修工进入厕所之后，先把门锁上，然后开始维修，维修结束之后把 ` Washroom` 的 ` isAvailable` 字段设置为 ` true` ，以表示厕所可用。

与此同时，一群急得像热锅上的蚂蚁的家伙在厕所门前打转转，他们想做神马不用我明说了吧😏😏：

` public class ShitTask implements Runnable { private Washroom washroom; private String name; public ShitTask (Washroom washroom, String name) { this.washroom = washroom; this.name = name; } @Override public void run () { synchronized (washroom.getLock()) { System.out.println(name + " 获取了厕所的锁" ); while (!washroom.isAvailable()) { // 一直等 } System.out.println(name + " 上完了厕所" ); } } } 复制代码`

这个 ` ShitTask` 描述了上厕所的一个流程，先获取到厕所的锁，然后判断厕所是否可用，如果不可用，则在一个死循环里不断的判断厕所是否可用，直到厕所可用为止，然后上完厕所释放锁走人。

然后我们看看现实世界都发生了什么吧：

` public class Test { public static void main (String[] args) { Washroom washroom = new Washroom(); new Thread( new RepairTask(washroom), "REPAIR-THREAD" ).start(); try { Thread.sleep( 1000L ); } catch (InterruptedException e) { throw new RuntimeException(e); } new Thread( new ShitTask(washroom, "狗哥" ), "BROTHER-DOG-THREAD" ).start(); new Thread( new ShitTask(washroom, "猫爷" ), "GRANDPA-CAT-THREAD" ).start(); new Thread( new ShitTask(washroom, "王尼妹" ), "WANG-NI-MEI-THREAD" ).start(); } } 复制代码`

学校先让维修工进入厕所维修，然后包括狗哥、猫爷、王尼妹在内的上厕所大军就开始围着厕所打转转的旅程，我们看一下执行结果：

` 维修工 获取了厕所的锁 厕所维修中，维修厕所是一件辛苦活，需要很长时间。。。 维修工把厕所修好了，准备释放锁了 王尼妹 获取了厕所的锁 王尼妹 上完了厕所 猫爷 获取了厕所的锁 猫爷 上完了厕所 狗哥 获取了厕所的锁 狗哥 上完了厕所 复制代码`

看起来没有神马问题，但是再回头看看代码，发现有两处特别别扭的地方：

* 

在main线程开启 ` REPAIR-THREAD` 线程后，必须调用 ` sleep` 方法等待一段时间才允许上厕所线程开启。

如果 ` REPAIR-THREAD` 线程和其他上厕所线程一块儿开启的话，就有可能上厕所的人，比如狗哥先获取到厕所的锁，然后维修工压根儿连厕所也进不去。但是真实情况可能真的这样的，狗哥先到了厕所，然后维修工才到。不过狗哥的处理应该不是一直待在厕所里，而是先出来等着，啥时候维修工说修好了他再进去。所以这点有些别扭～

* 

在一个上厕所的人获取到厕所的锁的时候，必须不断判断 ` Washroom` 的 ` isAvailable` 字段是否为 ` true` 。

如果一个人进入到厕所发现厕所仍然处在不可用状态的话，那它应该在某个地方休息，啥时候维修工把厕所修好了，再叫一下等着上厕所的人就好了嘛，没必要自己不停的去检查厕所是否被修好了。

总结一下，就是 一个线程在获取到锁之后，如果指定条件不满足的话，应该主动让出锁，然后到专门的等待区 ` 等待` ，直到某个线程完成了指定的条件，再 ` 通知` 一下在等待这个条件完成的线程，让它们继续执行 。

> 
> 
> 
> 如果你觉得上边这句话比较绕的话，我来给你翻译一下：
> 当上狗哥获取到厕所门锁之后，如果厕所处于不可用状态，那就主动让出锁，然后到等待上厕所的队伍里排队`等待`，直到维修工把厕所修理好，把厕所的状态置为可用后，维修工再通知需要上厕所的人，然他们正常上厕所。
> 
> 
> 

#### 具体使用方式 ####

为了实现这个构想，java里提出了一套叫 ` wait/notify` 的机制。当一个线程获取到锁之后，如果发现条件不满足，那就主动让出锁，然后把这个线程放到一个 ` 等待队列` 里 ` 等待` 去，等到某个线程把这个条件完成后，就 ` 通知` 等待队列里的线程他们等待的条件满足了，可以继续运行啦！

如果不同线程有不同的等待条件肿么办，总不能都塞到同一个 ` 等待队列` 里吧？是的，java里规定了 每一个锁都对应了一个 ` 等待队列` ，也就是说如果一个线程在获取到锁之后发现某个条件不满足，就主动让出锁然后把这个线程放到与它获取到的锁对应的那个等待队列里，另一个线程在完成对应条件时需要获取同一个锁，在条件完成后通知它获取的锁对应的 ` 等待队列` 。这个过程意味着 锁和等待队列建立了一对一关联 。

怎么 让出锁并且把线程放到与锁关联的等待队列中 以及怎么 通知等待队列中的线程相关条件已经完成 java已经为我们规定好了。我们知道， ` 锁` 其实就是个对象而已，在所有对象的老祖宗类 ` Object` 中定义了这么几个方法：

` public final void wait () throws InterruptedException public final void wait ( long timeout) throws InterruptedException public final void wait ( long timeout, int nanos) throws InterruptedException public final void notify () ; public final void notifyAll () ; 复制代码`

+--------------------------------+------------------------------------------------------------------------------------------+
|             方法名             |                                           说明                                           |
+--------------------------------+------------------------------------------------------------------------------------------+
| ` wait()`                      | 在线程获取到锁后，调用锁对象的本方法，线程释放锁并且把该线程放置到与锁对象关联的等待队列 |
| ` wait(long timeout)`          | 与 ` wait()`                                                                             |
|                                | 方法相似，只不过等待指定的毫秒数，如果超过指定时间则自动把该线程从等待队列中移出         |
| ` wait(long timeout, int       | 与上边的一样，只不过超时时间粒度更小，即指定的毫秒数加纳秒数                             |
| nanos)`                        |                                                                                          |
| ` notify()`                    | 通知一个在与该锁对象关联的等待队列的线程，使它从wait()方法中返回继续往下执行             |
| ` notifyAll()`                 | 与上边的类似，只不过通知该等待队列中的所有线程                                           |
+--------------------------------+------------------------------------------------------------------------------------------+

了解了这些方法的意思以后我们再来改写一下 ` ShitTask` ：

` public class ShitTask implements Runnable { // ... 为节省篇幅，省略相关字段和构造方法 @Override public void run () { synchronized (washroom.getLock()) { System.out.println(name + " 获取了厕所的锁" ); while (!washroom.isAvailable()) { try { washroom.getLock().wait(); //调用锁对象的wait()方法，让出锁，并把当前线程放到与锁关联的等待队列 } catch (InterruptedException e) { throw new RuntimeException(e); } } System.out.println(name + " 上完了厕所" ); } } } 复制代码`

看，原来我们在判断厕所是否可用的死循环里加了这么一段代码：

` washroom.getLock().wait(); 复制代码`

这段代码的意思就是让出厕所的锁，并且把当前线程放到与厕所的锁相关联的等待队列里。

然后我们也需要修改一下维修任务：

` public class RepairTask implements Runnable { // ... 为节省篇幅，省略相关字段和构造方法 @Override public void run () { synchronized (washroom.getLock()) { System.out.println( "维修工 获取了厕所的锁" ); System.out.println( "厕所维修中，维修厕所是一件辛苦活，需要很长时间。。。" ); try { Thread.sleep( 5000L ); //用线程sleep表示维修的过程 } catch (InterruptedException e) { throw new RuntimeException(e); } washroom.setAvailable( true ); //维修结束把厕所置为可用状态 washroom.getLock().notifyAll(); //通知所有在与锁对象关联的等待队列里的线程，它们可以继续执行了 System.out.println( "维修工把厕所修好了，准备释放锁了" ); } } } 复制代码`

大家可以看出来，我们在维修结束后加了这么一行代码：

` washroom.getLock().notifyAll(); 复制代码`

这个代码表示将通知所有在与锁对象关联的等待队列里的线程，它们可以继续执行了。

在使用java的 ` wait/notify` 机制修改了 ` ShitTask` 和 ` RepairTask` 后，我们在复原一下整个现实场景：

` public class Test { public static void main (String[] args) { Washroom washroom = new Washroom(); new Thread( new ShitTask(washroom, "狗哥" ), "BROTHER-DOG-THREAD" ).start(); new Thread( new ShitTask(washroom, "猫爷" ), "GRANDPA-CAT-THREAD" ).start(); new Thread( new ShitTask(washroom, "王尼妹" ), "WANG-NI-MEI-THREAD" ).start(); try { Thread.sleep( 1000 ); } catch (InterruptedException e) { throw new RuntimeException(e); } new Thread( new RepairTask(washroom), "REPAIR-THREAD" ).start(); } } 复制代码`

在这个场景中，我们可以刻意让着急上厕所的先到达了厕所，维修工最后抵达厕所，来看一下加了 ` wait/notify` 机制的代码的执行结果是：

` 狗哥 获取了厕所的锁 猫爷 获取了厕所的锁 王尼妹 获取了厕所的锁 维修工 获取了厕所的锁 厕所维修中，维修厕所是一件辛苦活，需要很长时间。。。 维修工把厕所修好了，准备释放锁了 王尼妹 上完了厕所 猫爷 上完了厕所 狗哥 上完了厕所 复制代码`

从执行结果可以看出来，狗哥、猫爷、王尼妹虽然先到达了厕所并且获取到锁，但是由于厕所处于不可用状态，所以都先调用 ` wait()` 方法让出了自己获得的锁，然后躲到与这个锁关联的等待队列里，直到维修工修完了厕所，通知了在等待队列中的狗哥、猫爷、王尼妹，他们才又开始继续执行上厕所的程序～

#### 通用模式 ####

经过上边的厕所案例，大家应该对 ` wait/notify` 机制有了大致了解，下边我们总结一下这个机制的通用模式。首先看一下等待线程的通用模式：

* 获取对象锁。
* 如果某个条件不满足的话，调用锁对象的 ` wait` 方法，被通知后仍要检查条件是否满足。
* 条件满足则继续执行代码。

通用的代码如下：

` synchronized (对象) { 处理逻辑（可选） while (条件不满足) { 对象.wait(); } 处理逻辑（可选） } 复制代码`

除了判断条件是否满足和调用 ` wait` 方法以外的代码，其他的处理逻辑是可选的。

下边再来看 ` 通知` 线程的通用模式：

* 获得对象的锁。
* 完成条件。
* 通知在等待队列中的等待线程。
` synchronized (对象) { 完成条件 对象.notifyAll();、 } 复制代码`
> 
> 
> 
> 小贴士：
> 别忘了同步方法也是使用锁的喔，静态同步方法的锁对象是该类的`Class对象`，成员同步方法的锁对象是`this对象`。所以如果没有刻意强调，下边所说的同步代码块也包含同步方法。
> 
> 
> 

了解了 ` wait/notify` 的通用模式之后，使用的时候需要特别小心，需要注意下边这些方面：

* 

必须在同步代码块中调用 ` wait` 、 ` notify` 或者 ` notifyAll` 方法 。

有的童鞋会有疑问，为啥 ` wait/notify` 机制的这些方法必须都放在同步代码块中才能调用呢？ ` wait` 方法的意思只是让当前线程停止执行，把当前线程放在等待队列里， ` notify` 方法的意思只是从等待队列里移除一个线程而已，跟加锁有什么关系？

答：因为 ` wait` 方法是运行在等待线程里的， ` notify` 或者 ` notifyAll` 是运行在通知线程里的。而执行 ` wait` 方法前需要判断一下某个条件是否满足，如果不满足才会执行 ` wait` 方法，这是一个 ` 先检查后执行` 的操作，不是一个 ` 原子性操作` ，所以如果不加锁的话，在多线程环境下等待线程和通知线程的执行顺序可能是这样的：

![image_1c1ufvde41bas7g2184e1ma4peg16.png-40.3kB](https://user-gold-cdn.xitu.io/2019/4/18/16a2e70791c3d448?imageView2/0/w/1280/h/960/ignore-error/1)

也就是说当等待线程已经判断条件不满足，正要执行 ` wait` 方法，此时通知线程抢先把条件完成并且调用了 ` notify` 方法，之后等待线程才执行到 ` wait` 方法，这会导致等待线程永远停留在等待队列而没有人再去 ` notify` 它。所以等待线程中的 判断条件是否满足、调用wait方法 和通知线程中 完成条件、调用notify方法 都应该是原子性操作，彼此之间是互斥的，所以用 同一个锁 来对这两个原子性操作进行同步，从而避免出现等待线程永久等待的尴尬局面。

如果不在同步代码块中调用 ` wait` 、 ` notify` 或者 ` notifyAll` 方法，也就是说没有获取锁就调用 ` wait` 方法，就像这样：

` 对象.wait(); 复制代码`

是会抛出 ` IllegalMonitorStateException` 异常的。

* 

在同步代码块中，必须调用获取的锁对象的 ` wait` 、 ` notify` 或者 ` notifyAll` 方法 。

也就是说不能随便调用一个对象的 ` wait` 、 ` notify` 或者 ` notifyAll` 方法。比如等待线程中的代码是这样的：

` synchronized (对象 1 ) { while (条件不满足) { 对象 2.wait(); //随便调用一个对象的wait方法 } } 复制代码`

通知线程中的代码是这样的：

` synchronized (对象 1 ) { 完成条件 对象 2.notifyAll(); } 复制代码`

对于代码 ` 对象2.wait()` ，表示让出当前线程持有的 ` 对象2` 的锁，而当前线程持有的是 ` 对象1` 的锁，所以这么写是错误的，也会抛出 ` IllegalMonitorStateException` 异常的。意思就是 如果当前线程不持有某个对象的锁，那它就不能调用该对象的 ` wait` 方法来让出该锁 。所以如果想让等待线程让出当前持有的锁，只能调用 ` 对象1.wait()` 。然后这个线程就被放置到与 ` 对象1` 相关联的等待队列中，在通知线程中只能调用 ` 对象1.notifyAll()` 来通知这些等待的线程了。

* 

在等待线程判断条件是否满足时，应该使用 ` while` ，而不是 ` if` 。

也就是说在判断条件是否满足的时候要使用 ` while` ：

` while (条件不满足) { //正确✅ 对象.wait(); } 复制代码`

而不是使用 ` if` ：

` if (条件不满足) { //错误❌ 对象.wait(); } 复制代码`

这个是因为在多线程条件下，可能在一个线程调用 ` notify` 之后立即又有一个线程把条件改成了不满足的状态，比如在维修工把厕所修好之后通知大家上厕所吧的瞬间，有一个小屁孩以迅雷不及掩耳之势又给厕所坑里塞了个瓶子，厕所又被置为不可用状态，等待上厕所的还是需要再判断一下条件是否满足才能继续执行。

* 

在调用完锁对象的 ` notify` 或者 ` notifyAll` 方法后，等待线程并不会立即从wait()方法返回，需要调用notify()或者notifyAll()的线程释放锁之后，等待线程才从wait()返回继续执行。 。

也就是说如果通知线程在调用完锁对象的 ` notify` 或者 ` notifyAll` 方法后还有需要执行的代码，就像这样：

` synchronized (对象) { 完成条件 对象.notifyAll(); ... 通知后的处理逻辑 } 复制代码`

需要把通知后的处理逻辑执行完成后，把锁释放掉，其他线程才可以从 ` wait` 状态恢复过来，重新竞争锁来执行代码。比方说在维修工修好厕所并通知了等待上厕所的人们之后，他还没有从厕所出来，而是在厕所的墙上写了 "XXX到此一游"之类的话之后才从厕所出来，从厕所出来才代表着释放了锁，狗哥、猫爷、王尼妹才开始争抢进入厕所的机会。

* 

` notify` 方法只会将等待队列中的一个线程移出，而 ` notifyAll` 方法会将等待队列中的所有线程移出 。

大家可以把上边代码中的 ` notifyAll` 方法替换称 ` notify` 方法，看看执行结果～

#### ` wait` 和 ` sleep` 的区别 ####

眼尖的小伙伴肯定发现， ` wait` 和 ` sleep` 这两个方法都可以让线程暂停执行，而且都有 ` InterruptedException` 的异常说明，那么它们的区别是啥呢？

* 

` wait` 是 ` Object` 的成员方法，而 ` sleep` 是 ` Thread` 的静态方法 。

只要是作为锁的对象都可以在同步代码块中调用自己的 ` wait` 方法， ` sleep` 是 ` Thread` 的静态方法，表示的是让当前线程休眠指定的时间。

* 

调用 ` wait` 方法需要先获得锁，而调用 ` sleep` 方法是不需要的 。

在一次强调，一定要在同步代码块中调用锁对象的 ` wait` 方法，前提是要获得锁！前提是要获得锁！前提是要获得锁！而 ` sleep` 方法随时调用～

* 

调用 ` wait` 方法的线程需要用 ` notify` 来唤醒，而 ` sleep` 必须设置超时值 。

* 

线程在调用 ` wait` 方法之后会先释放锁，而 ` sleep` 不会释放锁 。

这一点可能是最重要的一点不同点了吧，狗哥、猫爷、王尼妹这些线程一开始是获取到厕所的锁了，但是调用了 ` wait` 方法之后主动把锁让出，从而让维修工得以进入厕所维修。如果狗哥在发现厕所是不可用的条件时选择调用 ` sleep` 方法的话，线程是不会释放锁的，也就是说维修工无法获得厕所的锁，也就修不了厕所了～ 大家一定要谨记这一点啊！

### 总结 ###

* 

线程间需要通过通信才能协作解决某个复杂的问题。

* 

可变共享变量是天然的通信媒介，但是使用的时候一定要保证线程安全性，通常使用 ` volatile` 变量或 ` synchronized` 来保证线程安全性。

* 

一个线程在获取到锁之后，如果指定条件不满足的话，应该主动让出锁，然后到专门的等待区等待，直到某个线程完成了指定的条件，再通知一下在等待这个条件完成的线程，让它们继续执行。这个机制就是 ` wait/notify` 机制。

* 

等待线程的通用模式：

` synchronized (对象) { 处理逻辑（可选） while (条件不满足) { 对象.wait(); } 处理逻辑（可选） } 复制代码`

可以分为下边几个步骤：

* 获取对象锁。
* 如果某个条件不满足的话，调用锁对象的wait方法，被通知后仍要检查条件是否满足。
* 条件满足则继续执行代码。

* 

通知线程的通用模式：

` synchronized (对象) { 完成条件 对象.notifyAll();、 } 复制代码`

可以分为下边几个步骤：

* 获得对象的锁。
* 完成条件。
* 通知在等待队列中的等待线程。

* 

` wait` 和 ` sleep` 的区别

* 

wait是Object的成员方法，而sleep是Thread的静态方法。

* 

调用wait方法需要先获得锁，而调用sleep方法是不需要的。

* 

调用wait方法的线程需要用notify来唤醒，而sleep必须设置超时值。

* 

线程在调用wait方法之后会先释放锁，而sleep不会释放锁。

### 题外话 ###

写文章挺累的，有时候你觉得阅读挺流畅的，那其实是背后无数次修改的结果。如果你觉得不错请帮忙转发一下，万分感谢～ 这里是我的公众号，里边有更多技术干货，时不时扯一下犊子，欢迎关注：

![](https://user-gold-cdn.xitu.io/2019/4/18/16a2e7131c8643e4?imageView2/0/w/1280/h/960/ignore-error/1)

### 小册 ###

另外，作者还写了一本MySQL小册： [《MySQL是怎样运行的：从根儿上理解MySQL》的链接]( https://juejin.im/book/5bffcbc9f265da614b11b731?referrer=5bff96c6e51d45452f2d6f95 ) 。小册的内容主要是从小白的角度出发，用比较通俗的语言讲解关于MySQL进阶的一些核心概念，比如记录、索引、页面、表空间、查询优化、事务和锁等，总共的字数大约是三四十万字，配有上百幅原创插图。主要是想降低普通程序员学习MySQL进阶的难度，让学习曲线更平滑一点～ 有在MySQL进阶方面有疑惑的同学可以看一下：

![](https://user-gold-cdn.xitu.io/2019/4/17/16a2907eb8386a9b?imageView2/0/w/1280/h/960/ignore-error/1)