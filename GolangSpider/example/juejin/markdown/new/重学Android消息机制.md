# 重学Android消息机制 #

### Android消息机制 ###

> 
> 
> 
> 本文源码Api28
> 
> 

在Android在主线程的创建时，会自动创建一个looper，不需要我们自己来创建。

那么Android应用启动流程中，会由AMS调用的 ` ActivityThread` 类，它的main入口方法里：

` public static void main (String[] args) { ... Looper.prepareMainLooper(); // Find the value for {@link #PROC_START_SEQ_IDENT} if provided on the command line. // It will be in the format "seq=114" long startSeq = 0 ; if (args != null ) { for ( int i = args.length - 1 ; i >= 0 ; --i) { if (args[i] != null && args[i].startsWith(PROC_START_SEQ_IDENT)) { startSeq = Long.parseLong( args[i].substring(PROC_START_SEQ_IDENT.length())); } } } ActivityThread thread = new ActivityThread(); //注意： thread.attach( false , startSeq); if (sMainThreadHandler == null ) { sMainThreadHandler = thread.getHandler(); } if ( false ) { Looper.myLooper().setMessageLogging( new LogPrinter(Log.DEBUG, "ActivityThread" )); } // End of event ActivityThreadMain. Trace.traceEnd(Trace.TRACE_TAG_ACTIVITY_MANAGER); Looper.loop(); throw new RuntimeException( "Main thread loop unexpectedly exited" ); } 复制代码`

里面调用了到Looper的两个方法，Looper.prepareMainLooper以及Looper.loop()。

### Looper ###

先看prepareMainLooper

` public static void prepareMainLooper () { prepare( false ); synchronized (Looper.class) { if (sMainLooper != null ) { throw new IllegalStateException( "The main Looper has already been prepared." ); } sMainLooper = myLooper(); } } 复制代码`

在这里调用到了Looper主要有两个方法之一：prepare，而后面的判断sMainLooper是否为空以及赋值，都是为了保证只prepareMainLooper一次。

` public static @Nullable Looper myLooper () { return sThreadLocal.get(); } 复制代码`

那么我们可以只看prepare：

` private static void prepare ( boolean quitAllowed) { if (sThreadLocal.get() != null ) { throw new RuntimeException( "Only one Looper may be created per thread" ); } sThreadLocal.set( new Looper(quitAllowed)); } 复制代码`

可以看到，其实就是在prepare就做了两件事

* 判断当前对象是否为空，不为空就抛出异常----这是一个相对来说比较奇葩的判定，一般我们写程序，都是判断为空抛异常或者new，再仔细看这一段英文“一个线程只能有一个looper”，这就代表一个线程不能多次调用prepare来创建looper。
* 如果线程中没有looper，那么那new一个looper将其set到ThreadLocal中去。

#### ThreadLocal ####

顺便说一嘴ThreadLocal

ThreadLoacl是一个线程内部的数据存储类，通过它可以在指定的线程中存储数据，数据存储以扣，只有在这个线程中可以获取到存储的数据，对其它线程是无法获取到的——by《Android开发艺术探索》

特点在于使用set来存储数据，get来取出数据。

我们可以在代码中做个试验：

` final ThreadLocal<Integer> threadLocal = new ThreadLocal<>(); threadLocal.set( 1 ); new Thread( new Runnable() { @Override public void run () { Log.d(TAG, "run: " +threadLocal.get()); } }); 复制代码`

得到的结果果然是2019-06-05 22:05:07.933 3076-4709/com.apkcore.studdy D/MainActivity: run: null，可见在哪个线程放在数据，就必做在哪个线程取出。

关于ThreadLocal的内部代码，篇幅有限，下次再一起详细看。

继续来看Looper，我们已经看到在prepare中，如果线程中没有Looper对象时，会new一个looper，并把它加入到ThreadLocal中，那看一下它的构造函数

` private Looper ( boolean quitAllowed) { mQueue = new MessageQueue(quitAllowed); mThread = Thread.currentThread(); } 复制代码`

在构造方法中，实例化了一个消息队列MessageQueue，并还会获取到当前的线程，把它赋值给mThread。也就是说 **消息队列此时已经和当前线程绑定，其作用的区域为当前实例化looper的线程** 。

那么我们接下来可以看Main方法调用的Looper.loop()方法

` public static @Nullable Looper myLooper () { return sThreadLocal.get(); } public static void loop () { //获取一个looper对象，可以看到prepare一定在loop()调用之前 final Looper me = myLooper(); if (me == null ) { throw new RuntimeException( "No Looper; Looper.prepare() wasn't called on this thread." ); } ... for (;;) { //从消息队列MessageQueue中取消息，如果消息为空，跳出循环 Message msg = queue.next(); // might block if (msg == null ) { // No message indicates that the message queue is quitting. return ; } ... //如果消息不为空，那么msg交给msg.target.dispatchMessage(msg)来处理 try { msg.target.dispatchMessage(msg); dispatchEnd = needEndTime ? SystemClock.uptimeMillis() : 0 ; } finally { if (traceTag != 0 ) { Trace.traceEnd(traceTag); } } ... //方法回收 msg.recycleUnchecked(); } } 复制代码`

去掉了一些对我们分析不是特别重要的代码，对关键代码做了注释，其实就是获取当前的looper对象，并进入一个死循环中，一直从消息队列中取消息，如果消息为null就跳出循环，如果不为空，那么把消息交给 ` msg.target.dispatchMessage(msg)` 来处理，msg.target在后面会看到，它其实就是Handler，过会分析。最后调用recyleUnchecked来方法回收。

这里有一个知识点 ` queue.next()` ，我们到下面分析messageQueue的时候一起来分析。

### Handler ###

我们在使用Handler的时候

` private Handler mHandler = new Handler() { @Override public void handleMessage (Message msg) { super.handleMessage(msg); if (msg.what == 0 ) { Log.d(TAG, "handleMessage: 0" ); } } }; 复制代码`

一般是这么使用，当然在Activity直接这么使用，是有可能内存泄露的，但这不是我们这一节要讲的重点。

那么我们先从handler的构造函数开始看

` public Handler () { this ( null , false ); } public Handler (Callback callback, boolean async) { ... mLooper = Looper.myLooper(); if (mLooper == null ) { throw new RuntimeException( "Can't create handler inside thread " + Thread.currentThread() + " that has not called Looper.prepare()" ); } mQueue = mLooper.mQueue; mCallback = callback; mAsynchronous = async; } 复制代码`

从构造方法来看，还是先要获取到Looper的myLooper方法，从ThreadLocal中取出looper对象，之后判断对象是否为空，为空则抛出异常。不为空则获取到了looper的消息队列，这样，Handler中就持有了messageQueue的引用。

` Asynchronous messages represent interrupts or events that do not require global ordering with respect to synchronous messages. Asynchronous messages are not subject to the synchronization barriers introduced by {@link MessageQueue#enqueueSyncBarrier(long)}. 复制代码`

注释中详细讲了mAsynchronous这个参数的作用，是一个 **同步障碍，异步消息不受引入的同步障碍的限制，这一点也是同步消息和异步消息的区别了** 。在后面会讲到。

我们在使用handler发送消息时，一般使用sendMessage来发送，如

` new Thread( new Runnable() { @Override public void run () { try { Thread.sleep( 5000 ); mHandler.sendEmptyMessage( 0 ); } catch (InterruptedException e) { e.printStackTrace(); } } }).start(); 复制代码`

还是先忽略代码中直接new Thread的问题，继续跟踪源码

` public final boolean sendEmptyMessage ( int what) { return sendEmptyMessageDelayed(what, 0 ); } public final boolean sendEmptyMessageDelayed ( int what, long delayMillis) { Message msg = Message.obtain(); msg.what = what; return sendMessageDelayed(msg, delayMillis); } public final boolean sendMessageDelayed (Message msg, long delayMillis) { if (delayMillis < 0 ) { delayMillis = 0 ; } return sendMessageAtTime(msg, SystemClock.uptimeMillis() + delayMillis); } public boolean sendMessageAtTime (Message msg, long uptimeMillis) { MessageQueue queue = mQueue; if (queue == null ) { RuntimeException e = new RuntimeException( this + " sendMessageAtTime() called with no mQueue" ); Log.w( "Looper" , e.getMessage(), e); return false ; } return enqueueMessage(queue, msg, uptimeMillis); } 复制代码`

可以看到，都是层层调用，最终调到sendMessageAtTime方法里。先获取了队列的messageQueue，判断queue不能为空，然后调用到了enqueueMessage方法，把它自己的queue，msg和uptimeMillis都一并传了过去。

` private boolean enqueueMessage (MessageQueue queue, Message msg, long uptimeMillis) { msg.target = this ; if (mAsynchronous) { msg.setAsynchronous( true ); } return queue.enqueueMessage(msg, uptimeMillis); } 复制代码`

这enqueueMessage方法里，注意了！我们看到了熟悉的字眼 **msg.target** ，在Handler的构造方法中， ` mQueue = mLooper.mQueue;` ，传入到enqueueMessage方法中的就是这个值，则msg.target赋值为了Handler时，在looper中， **把消息交给 ` msg.target.dispatchMessage(msg)` 来处理** ，就是交给了Handler来处理。

也就是说 **Handler发出来的消息，全部发送给了MessageQueue中，并调用enqueueMessage方法，而Looper的loop循环中，又调用了MessageQueue的next方法，把这个消息给dispatchMessage来处理**

那么再来看dispatchMessage方法

` public void dispatchMessage (Message msg) { if (msg.callback != null ) { handleCallback(msg); } else { if (mCallback != null ) { if (mCallback.handleMessage(msg)) { return ; } } handleMessage(msg); } } 复制代码`

在这里，就能看到handler设置不同的回调的优先级了，如果msg调用了callback，那么只调用此callback，如果没有设置，那么构造方法中有传入callback的话，回调此callback，如果这二者都没有设置，才会调用覆写的Handler里的handleMessage方法

我们继承来看 ` enqueueMessage` 方法

` private boolean enqueueMessage (MessageQueue queue, Message msg, long uptimeMillis) { msg.target = this ; if (mAsynchronous) { msg.setAsynchronous( true ); } return queue.enqueueMessage(msg, uptimeMillis); } 复制代码`

在这里可以看到，最终会调用到 ` queue.enqueueMessage(msg, uptimeMillis);` 中，那我们接下来看MessageQueue

### MessageQueue ###

messageQueue是在Looper的构造方法中生成的

` private Looper ( boolean quitAllowed) { mQueue = new MessageQueue(quitAllowed); mThread = Thread.currentThread(); } 复制代码`

我们继续看它的 ` enqueueMessage` 方法

` boolean enqueueMessage (Message msg, long when) { ... msg.when = when; Message p = mMessages; boolean needWake; if (p == null || when == 0 || when < p.when) { // New head, wake up the event queue if blocked. //新消息会插入到链表的表头，这意味着队列需要调整唤醒时间 msg.next = p; mMessages = msg; needWake = mBlocked; } else { // Inserted within the middle of the queue. Usually we don't have to wake // up the event queue unless there is a barrier at the head of the queue // and the message is the earliest asynchronous message in the queue. //新消息插入到链表的内部，一般情况下，这不需要调整唤醒时间 //但还是要考虑到当表头是“同步分割栏的情况” needWake = mBlocked && p.target == null && msg.isAsynchronous(); Message prev; for (;;) { prev = p; p = p.next; if (p == null || when < p.when) { break ; } //注意isAsynchronous是在上面if (mAsynchronous) { msg.setAsynchronous(true); }这里有设置，默认false if (needWake && p.isAsynchronous()) { //当msg是异步的，也不是链表的第一个异步消息，所以就不用唤醒了 needWake = false ; } } msg.next = p; // invariant: p == prev.next prev.next = msg; } // We can assume mPtr != 0 because mQuitting is false. if (needWake) { nativeWake(mPtr); } } return true ; } 复制代码`

源码的关键地方都给了注释，只是对同步分割栏不理解，这个马上就讲。上述源码，就是在消息链表中找到合适的位置，插入message节点。因为消息链是按时间进行排序的，所以主要是比对message携带的when信息，首个节点对应着最先处理的消息，如果message被插入到表头了，就意味着最近唤醒时间也要做调整，把以needwake就设为了true，以便走到nativeWake(mPtr);

#### 同步分割栏 ####

[Android 中的 Handler 同步分割栏]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqinxue24%2Farticle%2Fdetails%2F80396315 )

[Handler之同步屏障机制(sync barrier)]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fasdgbc%2Farticle%2Fdetails%2F79148180 )

所谓的同步分割栏，可以理解为一个特殊的Message，它的target域为null，它不能通过sengMessageAtTime等方法传入队列中，只能通过调用Looper的postSyncBarrier()

##### 作用 #####

它卡在消息链表的某个位置，当异步线程可以在Handler中设置异步卡子，设置好了以后， **当前Handler的同步Message都不再执行** ，直到异步线程将卡子去掉。在Android的消息机制里，同步的message和异步的message也就是这点区别，也就是说，如果消息列表中没有设置同步分割栏的话，那么其实它俩处理是一样的。

继续看上面的nativeWake()方法，这是一个native方法，对应在C++里，有兴趣的盆友，可以找 ` framworks/base/core/jni/android_os_MessageQueue.cpp` 来查看它的源码，这里我就不分析了，主要动作就是身一个管道的写入端写入了W。

回到最上面，我们在讲Looper.loop的源码中，我们还留下了一个msg.next()没有分析

` Message next () { ... int pendingIdleHandlerCount = - 1 ; // -1 only during first iteration int nextPollTimeoutMillis = 0 ; for (;;) { if (nextPollTimeoutMillis != 0 ) { Binder.flushPendingCommands(); } //阻塞在此 nativePollOnce(ptr, nextPollTimeoutMillis); synchronized ( this ) { // Try to retrieve the next message. Return if found. //获取next消息，如能得到就返回 final long now = SystemClock.uptimeMillis(); Message prevMsg = null ; //尝试拿消息队列里当前第一个消息 Message msg = mMessages; if (msg != null && msg.target == null ) { // Stalled by a barrier. Find the next asynchronous message in the queue. //如果队列里拿的msg是那个同步分割栏，那么就寻找后面的第一个异步消息 do { prevMsg = msg; msg = msg.next; } while (msg != null && !msg.isAsynchronous()); } if (msg != null ) { if (now < msg.when) { // Next message is not ready. Set a timeout to wake up when it is ready. //下一条消息未就绪。设置超时以在准备就绪时唤醒。 nextPollTimeoutMillis = ( int ) Math.min(msg.when - now, Integer.MAX_VALUE); } else { // Got a message. mBlocked = false ; if (prevMsg != null ) { prevMsg.next = msg.next; } else { //重新设置一下消息队列的头部 mMessages = msg.next; } msg.next = null ; if (DEBUG) Log.v(TAG, "Returning message: " + msg); msg.markInUse(); return msg; //返回 } } else { // No more messages. nextPollTimeoutMillis = - 1 ; } // Process the quit message now that all pending messages have been handled. if (mQuitting) { dispose(); return null ; } // If first time idle, then get the number of idlers to run. // Idle handles only run if the queue is empty or if the first message // in the queue (possibly a barrier) is due to be handled in the future. if (pendingIdleHandlerCount < 0 && (mMessages == null || now < mMessages.when)) { pendingIdleHandlerCount = mIdleHandlers.size(); } if (pendingIdleHandlerCount <= 0 ) { // No idle handlers to run. Loop and wait some more. mBlocked = true ; continue ; } if (mPendingIdleHandlers == null ) { mPendingIdleHandlers = new IdleHandler[Math.max(pendingIdleHandlerCount, 4 )]; } mPendingIdleHandlers = mIdleHandlers.toArray(mPendingIdleHandlers); } // Run the idle handlers. // We only ever reach this code block during the first iteration. //处理idleHandlers部分，空闲时handler for ( int i = 0 ; i < pendingIdleHandlerCount; i++) { final IdleHandler idler = mPendingIdleHandlers[i]; mPendingIdleHandlers[i] = null ; // release the reference to the handler boolean keep = false ; try { keep = idler.queueIdle(); } catch (Throwable t) { Log.wtf(TAG, "IdleHandler threw exception" , t); } if (!keep) { synchronized ( this ) { mIdleHandlers.remove(idler); } } } // Reset the idle handler count to 0 so we do not run them again. pendingIdleHandlerCount = 0 ; // While calling an idle handler, a new message could have been delivered // so go back and look again for a pending message without waiting. nextPollTimeoutMillis = 0 ; } } 复制代码`

其实源码中英文注释已经讲得比较直白了，我不过是随便翻译了一下，这个函数里的for循环并不是起循环摘取消息节点的作用，而是为了连贯事件队列的首条消息是否真的到时间了，如果到了，就直接返回这个msg，如果时间还没有到，就计算一个比较精确的等待时间（nextPollTimeoutMillis），计算完后，for循环会再次调用nativePollOnce(mPtr, nextPollTimeoutMillis)进入阻塞，等待合适的时长。

上面代码中也处理了“同步分割栏”，如果队列中有它的话，是千万不能返回的，要尝试寻找其后第一个异步消息。

next里面另一个比较重要的就是IdleHandler，当消息队列处于空闲时，会判断用户是否设置了IdleHandler，如果有的话，则会尝试处理，依次调用IdleHandler的queueIdle()函数。这个特性，我们可以用在性能优化里，比如延迟加载：我们经常为了优化一个Activity的显示速度，可能把一些非必要第一时间启动的耗时任务，放在界面加载完成后进行，但是就算界面加载完成了，耗时任务一样的会占用cpu，当我们正好在操作时，一样可能会造成卡顿的现象，那么我们完全可以利用Idle Handler来做处理。

#### nativePollOnce ####

前面说了next中调用nativePollOnce起到了阻塞作用，保证消息循环不会在无消息处理时一直在那循环，它的实现还是在android_os_MessageQueue.cpp文件中，同样在这里不做过多的深入，有兴趣的童鞋可以自行在frameworks/base/core/jni/android_os_MessageQueue.cpp中查看。

这里提一点，在c++中，pollonce调用了c++层的looper对象，在这个looper和我们java中的Looper是不一样的，其内部除了创建一个管道外，还创建了一个epoll来监听管道的读取端。

### 扩展 ###

[Android中为什么主线程不会因为Looper.loop()方法造成阻塞]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fchenxibobo%2Fp%2F9640472.html ) ，在此文中详细讲了为什么我们在主线程loop死循环而不会卡死。

对于线程既然是一段可执行的代码，当可执行代码执行完成后，线程生命周期便该终止了，线程退出。而对于主线程，我们是绝不希望会被运行一段时间，自己就退出，那么如何保证能一直存活呢？ **简单做法就是可执行代码是能一直执行下去的，死循环便能保证不会被退出** 例如，binder线程也是采用死循环的方法，通过循环方式不同与Binder驱动进行读写操作，当然并非简单地死循环，无消息时会休眠。但这里可能又引发了另一个问题，既然是死循环又如何去处理其他事务呢？通过创建新线程的方式。

真正会卡死主线程的操作是在回调方法onCreate/onStart/onResume等操作时间过长，会导致掉帧，甚至发生ANR，looper.loop本身不会导致应用卡死。

消息循环数据通信采用的是epoll机制，它能显著的提高CPU的利用率，另外Android应用程序的主线程在进入消息循环前，会在内部创建一个Linux管道（Pipe），这个管道的作用是使得Android应用主线程在消息队列唯恐时可以进入空闲等待状态，并且使得当应用程序的消息队列有消息需要处理是唤醒应用程序的主线程。也就是说在无消息时，循环处于睡眠状态，并不会出现卡死情况。

[CSDN]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2FApkCore )

下面是我的公众号，欢迎大家关注我

![](https://user-gold-cdn.xitu.io/2019/6/6/16b286e4442df45b?imageView2/0/w/1280/h/960/ignore-error/1)