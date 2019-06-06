# java并发之计数器CountDownLatch原理 #

## ##

### ###

### ###

### ###

### ###

### ###

### ###

### ###

### ###

### ###

### CountDownLatch简介 ###

CountDownLatch顾名思义，count + down + latch ＝ 计数 ＋ 减 ＋ 门闩（这么拆分也是便于记忆=_=） 可以理解这个东西就是个计数器，只能减不能加，同时它还有个门闩的作用，当计数器不为0时，门闩是锁着的；当计数器减到0时，门闩就打开了。
如果你感到懵比的话，可以类比考生考试交卷，考生交一份试卷，计数器就减一。直到考生都交了试卷（计数器为0），监考老师（一个或多个）才能离开考场。至于考生是否做完试卷，监考老师并不关注。只要都交了试卷，他就可以做接下来的工作了。

### CountDownLatch实现原理 ###

下面从构造方法开始，一步步解释实现的原理：

### 构造方法 ###

下面是实现的源码，非常简短，主要是创建了一个Sync对象。

` public CountDownLatch(int count) { if (count < 0) throw new IllegalArgumentException( "count < 0" ); this.sync = new Sync(count); } 复制代码`

## ##

### Sync对象 ###

` private static final class Sync extends AbstractQueuedSynchronizer { private static final long serialVersionUID = 4982264981922014374L; Sync(int count) { set State(count); } int getCount () { return getState(); } protected int tryAcquireShared(int acquires) { return (getState() == 0) ? 1 : -1; } protected boolean tryReleaseShared(int releases) { // Decrement count; signal when transition to zero for (;;) { int c = getState(); if (c == 0) return false ; int nextc = c-1; if (compareAndSetState(c, nextc)) return nextc == 0; } } } 复制代码`

假设我们是这样创建的：new CountDownLatch(5)。其实也就相当于new Sync(5)，相当于setState(5)。setState其实就是共享锁资源总数,我们可以暂时理解为设置一个计数器，当前计数器初始值为5。

tryAcquireShared方法其实就是判断一下当前计数器的值，是否为0了，如果为0的话返回1（ **返回1的时候，就表示获取锁成功,awit()方法就不再阻塞** ）。

tryReleaseShared方法就是利用CAS的方式，对计数器进行减一的操作，而我们实际上每次调用countDownLatch.countDown()方法的时候，最终都会调到这个方法，对计数器进行减一操作，一直减到0为止。

### countDownLatch.await() ###

` public void await() throws InterruptedException { sync.acquireSharedInterruptibly(1); } 复制代码`

代码很简单，就一句话（注意acquireSharedInterruptibly（）方法是抽象类：AbstractQueuedSynchronizer的一个方法，我们上面提到的Sync继承了它），我们跟踪源码，继续往下看：

### acquireSharedInterruptibly(int arg) ###

` public final void acquireSharedInterruptibly(int arg) throws InterruptedException { if (Thread.interrupted()) throw new InterruptedException(); if (tryAcquireShared(arg) < 0) do AcquireSharedInterruptibly(arg); } 复制代码`

源码也是非常简单的，首先判断了一下，当前线程是否有被中断，如果没有的话，就调用tryAcquireShared(int acquires)方法，判断一下当前线程是否还需要“阻塞”。其实这里调用的tryAcquireShared方法，就是我们上面提到的java.util.concurrent.CountDownLatch.Sync.tryAcquireShared(int)这个方法。 当然，在一开始我们没有调用过countDownLatch.countDown()方法时，这里tryAcquireShared方法肯定是会返回-1的，因为会进入到doAcquireSharedInterruptibly方法。

### doAcquireSharedInterruptibly(int arg) ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25a692fb263fb?imageView2/0/w/1280/h/960/ignore-error/1)

### countDown()方法 ###

` // 计数器减1 public void countDown () { sync.releaseShared(1); } //调用AQS的releaseShared方法 public final boolean releaseShared(int arg) { if (tryReleaseShared(arg)) {//计数器减一 do ReleaseShared();//唤醒后继结点,这个时候队列中可能只有调用过await()的线程节点,也可能队列为空 return true ; } return false ; } 复制代码`

这个时候，我们应该对于countDownLatch.await()方法是怎么“阻塞”当前线程的，已经非常明白了。其实说白了，就是当你调用了countDownLatch.await()方法后，你当前线程就会进入了一个死循环当中，在这个死循环里面，会不断的进行判断，通过调用tryAcquireShared方法，不断判断我们上面说的那个计数器，看看它的值是否为0了（为0的时候，其实就是我们调用了足够多 countDownLatch.countDown()方法的时候），如果是为0的话，tryAcquireShared就会返回1，代码也会进入到图中的红框部分，然后跳出了循环，也就不再“阻塞”当前线程了。需要注意的是，说是在不停的循环，其实也并非在不停的执行for循环里面的内容，因为在后面调用parkAndCheckInterrupt（）方法时，在这个方法里面是会调用 LockSupport.park(this);来挂起当前线程。

### CountDownLatch 使用的注意点： ###

1、只有当count为0时， **await之后的程序才够执行** 。

****

2、countDown必须写在finally中，防止发生异程常时，导致程序死锁。

### 使用场景： ###

比如对于马拉松比赛，进行排名计算，参赛者的排名，肯定是跑完比赛之后，进行计算得出的，翻译成Java识别的预发，就是N个线程执行操作，主线程等到N个子线程执行完毕之后，在继续往下执行。

****

` public static void testCountDownLatch (){ int threadCount = 10; final CountDownLatch latch = new CountDownLatch(threadCount); for (int i=0; i< threadCount; i++){ new Thread(new Runnable () { @Override public void run () { System.out.println( "线程" + Thread.currentThread().getId() + "开始出发" ); try { Thread.sleep(1000); System.out.println( "线程" + Thread.currentThread().getId() + "已到达终点" ); } catch (InterruptedException e) { e.printStackTrace(); } fianlly { latch.countDown(); } } }).start(); } try { latch.await(); } catch (InterruptedException e) { e.printStackTrace(); } System.out.println( "10个线程已经执行完毕！开始计算排名" ); } 复制代码`

结果:

` 线程10开始出发 线程13开始出发 线程12开始出发 线程11开始出发 线程14开始出发 线程15开始出发 线程16开始出发 线程17开始出发 线程18开始出发 线程19开始出发 线程14已到达终点 线程15已到达终点 线程13已到达终点 线程12已到达终点 线程10已到达终点 线程11已到达终点 线程16已到达终点 线程17已到达终点 线程18已到达终点 线程19已到达终点 10个线程已经执行完毕！开始计算排名 复制代码`