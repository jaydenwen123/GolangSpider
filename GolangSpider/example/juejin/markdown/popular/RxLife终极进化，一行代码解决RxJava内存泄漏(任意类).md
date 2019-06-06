# RxLife终极进化，一行代码解决RxJava内存泄漏(任意类) #

### 前言 ###

距离 [RxLife]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxLife ) 上个版本的开发已经过去一个多月了，这段时间一直在忙着 [RxHttp库]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxHttp ) 的更新及推广，如果你对RxHttp还不了解，强烈推荐你阅读

[RxHttp 一条链发送请求，新一代Http请求神器]( https://juejin.im/post/5cbd267fe51d456e2b15f623 )

[Android 史上最优雅的实现文件上传、下载及进度的监听]( https://juejin.im/post/5cd0568a51882569bf52153c )

上面两篇文章分别得到「玉刚说」及「刘望舒」微信公众号独家原创发布，看完上面两篇文章，相信你会爱上 [RxHttp]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxHttp ) 。

曾经有读者向我反馈，基于OkHttp开发的请求框架太多了，眼花缭乱，根本学不动，为此，我想说，RxHttp库学习成本极低，只要学会请求三部曲，就可以基本掌握，所有的请求都基于这3个步骤，如下：

` RxHttp.get( "http://..." ) //第一步，确定请求方式.fromSimpleParser(String.class) //第二步，确定解析器.subscribe(s -> { //第三部，订阅观察者 //成功回调 }, throwable -> { //失败回调 }); 复制代码`

然而，尽管框架太多， [RxHttp]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxHttp ) 还是收获了众多粉丝，推广6周，截止本文发表在Github上就已有了287个star，成绩虽然不是很优越，但也足以证明RxHttp有存在的必要，个人会一直维护下去(即使我不维护了，也会找人维护)，欢迎大家star，你的star就是我坚持的动力。

呃呃呃，貌似有点跑题了，我们回到 [RxLife]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxLife ) 的话题上，早在RxLife1.0.4版本，在Activity/Fragment上，我们就已经实现一行代码解决RxJava内存泄漏问题，如下：

` Observable.interval( 1 , 1 , TimeUnit.SECONDS) //隔一秒发送一条消息.as(RxLife.as( this )) //这里this 为LifecycleOwner接口对象.subscribe(aLong -> { Log.e( "LJX" , "accept=" + aLong); }); 复制代码`

然后这段代码只能在Activity/Fragment上书写，因为这里的 ` this` 为LifecycleOwner接口对象，而目前只有Activity/Fragment实现了这个接口，如果我们想在View、ViewModel或者任意上使用 ` RxLife.as(this)` ，就会报错，导致没法使用，为此，本篇文章就代理大家，如何在任意类上使用 ` RxLife.as(this)` 代码。

**gradle依赖**

` dependencies { implementation 'com.rxjava.rxlife:rxlife:1.0.6' //if you use AndroidX implementation 'com.rxjava.rxlife:rxlife-x:1.0.6' } 复制代码`

### Scope作用域 ###

开始之前，先给大家讲解作用域的概念，在这，我要感谢 [却把清梅嗅]( https://juejin.im/user/588555ff1b69e600591e8462 ) 这位大神，是他在我之前的文章中留言，让我首次了解到了作用域的概念。那么什么是作用域，简单来说，就是一个类从创建到回收，这就是它的作用域，比如：Activity/Fragment的作用域就是从 ` onCreate` 到 ` onDestroy` ；View的作用域就是从 ` onAttachedToWindow` 到 ` onDetachedFromWindow` ；ViewModel的作用域就是从 ` 构造方法` 到 ` onCleared` 方法；其它任意类的作用域就是从创建到销毁，当然，你也可以自己指定一些类的作用域。到这，相信你已经充分了解了作用域的概念，下面，我们正式开始。

### Activity/Fragment ###

首先，我们来回顾下，在Activity/Fragment上如何使用 ` RxLife.as` 操作符，如下：

` //在Activity/Fragment上 Observable.interval( 1 , 1 , TimeUnit.SECONDS) //隔一秒发送一条消息.as(RxLife.as( this )) //这里this 为LifecycleOwner接口对象.subscribe(aLong -> { Log.e( "LJX" , "accept=" + aLong); }); 复制代码`

此时Activity/Fragment销毁，就会自动关闭RxJava管道，避免内存泄漏，此时的 ` RxLife.as(LifecycleOwner owner)` 接受的是一个LifecycleOwner接口对象，因为Activity/Fragment实现了这个接口，故这里我们可以直接传 ` this` ，然而，如果我们想在View上，如何实现呢？

### View ###

我们在 [RxLife 1.0.5版本]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxLife ) 上，加入了两个重要方法，如下：

` public static <T> RxConverter<T> as (View view) { return as(ViewScope.from(view), false ); } public static <T> RxConverter<T> as (Scope scope) { return as(scope, false ); } 复制代码`

可以看到，新增的两个方法，支持传View对象及Scope对象，Scope对象我们稍后再做讲解，我们直接来看看在View上如何使用：

` //在View上 Observable.interval( 1 , 1 , TimeUnit.SECONDS) //隔一秒发送一条消息.as(RxLife.as( this )) //这里this 为View对象.subscribe(aLong -> { Log.e( "LJX" , "accept=" + aLong); }); 复制代码`

可以看到，在代码上，跟Activity/Fragment没有任何区别，都是 ` RxLife.as(this)` 一行代码，要说区别的话，那就是作用域的不同，传入一个View对象时，那么RxJava就会在View的 ` onDetachedFromWindow` 方法回调时中断管道，避免内存泄漏，这一点在RecyclerView/ListView列表上非常实用，item被滑出，对应的管道就会被中断。

### ViewModel ###

ViewModel是Google Jetpack里面的组件之一，如果还不了解的请查看 [却把清梅嗅对ViewModel的介绍]( https://juejin.im/post/5c047fd3e51d45666017ff86 ) ，在这不做过多介绍，上面我们说过ViewModel的作用域是从 ` 构造方法` 到 ` onCleared` 方法，而这个 ` onCleared` 方法会在Activity/Fragment销毁时自动执行，故ViewModel能感知Activity/Fragment的销毁，RxLife正是利用了这一点，并结合 ` RxLife.as(Scope scope)` 方法，做到中断RxJava的目的，我们先来看看如何实现的：

` public class MyViewModel extends ScopeViewModel { public MyViewModel () { Observable.interval( 1 , 1 , TimeUnit.SECONDS) .as(RxLife.as( this )) //这里的this 为Scope接口对象.subscribe(aLong -> { Log.e( "LJX" , "MyViewModel aLong=" + aLong); }); } } 复制代码`

然后在Activity/Fragment上就可以这样调用

` //在Activity/Fragment上 MyViewModel viewModel = ViewModelProviders.of( this ).get(MyViewModel.class) 复制代码`

**注：要想ViewModel对象感知Activity/Fragment销毁事件，不能使用new 关键字创建对象，必须要通过ViewModelProviders类获取ViewModel对象**

到这，如果你是看标题进来的，我要给你说声抱歉，在ViewModel上不是真正意义上一行代码实现的，而是要准一些准备工作，比如上面，我们继承了 ` ScopeViewModel` 类，我们来看看ScopeViewModel类源码：

` public class ScopeViewModel extends ViewModel implements Scope { private CompositeDisposable mDisposables; @Override public void onScopeStart (Disposable d) { addDisposable(d); //订阅事件时回调 } @Override public void onScopeEnd () { //事件正常结束时回调 } private void addDisposable (Disposable disposable) { CompositeDisposable disposables = mDisposables; if (disposables == null ) { disposables = mDisposables = new CompositeDisposable(); } disposables.add(disposable); } private void dispose () { final CompositeDisposable disposables = mDisposables; if (disposables == null ) return ; disposables.dispose(); } @Override protected void onCleared () { super.onCleared(); //Activity/Fragment 销毁时回调 dispose(); //中断RxJava管道 } } 复制代码`

可以看到ScopeViewModel继承了ViewModel类并实现类Scope接口，所以我们能在MyViewModel中直接使用 ` RxLife.as(Scope scope)` 。

到这，也许有人会说，这种继承的方式侵入性太强了，不友好，是，我也是这么认为的，所以我并没有将 ` ScopeViewModel` 类封装到RxLife内部中，有需要自取。 在这我想问，除了这种方式，还有更优雅的方式吗？我想应该没有（如果你有更好方式，请留言），在非Activity/Fragment/View类中，无论是 [trello/RxLifecycle]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftrello%2FRxLifecycle ) 还是 [uber/AutoDispose]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fuber%2FAutoDispose ) ，都需要做很多的准备工作，而 [RxLife]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxLife ) 相对于前面两者，更加的简单。

### 任意类 ###

相信大家对MVP都有了一定的了解，而在P层，我们一般都有发送Http请求的需求，而想在Activity/Fragment 销毁时，自动关闭P层发送的请求，我们来看看RxLife如何实现

` public class Presenter extends BaseScope { public Presenter (LifecycleOwner owner) { super (owner); Observable.interval( 1 , 1 , TimeUnit.SECONDS) .as(RxLife.as( this )) //这里的this 为Scope接口对象.subscribe(aLong -> { Log.e( "LJX" , "accept aLong=" + aLong); }); } } 复制代码`

可以看到，只需要继承 ` BaseScope` 类，就可以直接使用 ` RxLife.as(this)` ，我们来看看BaseScope类源码

` public class BaseScope implements Scope , GenericLifecycleObserver { private CompositeDisposable mDisposables; public BaseScope (LifecycleOwner owner) { owner.getLifecycle().addObserver( this ); } @Override public void onScopeStart (Disposable d) { addDisposable(d); } @Override public void onScopeEnd () { } private void addDisposable (Disposable disposable) { CompositeDisposable disposables = mDisposables; if (disposables == null ) { disposables = mDisposables = new CompositeDisposable(); } disposables.add(disposable); } private void dispose () { final CompositeDisposable disposables = mDisposables; if (disposables == null ) return ; disposables.dispose(); } @Override public void onStateChanged (LifecycleOwner source, Event event) { //Activity/Fragment 生命周期回调 if (event == Event.ON_DESTROY) { //Activity/Fragment 销毁 source.getLifecycle().removeObserver( this ); dispose(); //中断RxJava管道 } } } 复制代码`

可以看到， ` BaseScope` 类跟 ` ScopeViewModel` 类源码差不多，都实现了 ` Scope` 接口，并通过 ` CompositeDisposable` 管理 ` Disposable` 对象。这两个类都未封装进 [RxHttp]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxHttp ) 库中，有需要自取。

### 原理 ###

说起原理，其实 [trello/RxLifecycle]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftrello%2FRxLifecycle ) 、 [uber/AutoDispose]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fuber%2FAutoDispose ) 、 [RxLife]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxLife ) 三者的原理都是一样的，都是拿到最下层观察者的 ` Disposable` 对象，然后在某个时机，调用该对象的 ` Disposable.dispose()` 方法中断管道，以达到目的。 原理都一样，然而实现都大不相同，

* 

trello/RxLifecycle (3.0.0版本) 内部只有一个管道，但却有两个事件源，一个发送生命周期状态变化，一个发送正常业务逻辑，最终通过takeUntil操作符对事件进行过滤，当监听到符合条件的事件时，就会将管道中断，从而到达目的

* 

uber/AutoDispose（1.2.0版本） 内部维护了两个管道，一个是发送生命周期状态变化的管道，我们称之为A管道，另一个是业务逻辑的管道，我们称至为B管道，B管道持有A管道的观察者引用，故能监听A管道的事件，当监听到符合条件的事件时，就会将A、B管道同时中断，从而到达目的

* 

RxHttp 内部只有一个业务逻辑的管道，通过自定义观察者，拿到Disposable对象，暴露给Scope接口，Scope的实现者就可以在合适的时机调用 ` Disposable.dispose()` 方法中断管道，从而到达目的

### 问题暴露 ###

我们知道，任意类想要监听Activity/Fragment生命周期，都必须要实现 ` LifecycleObserver` 接口，然后通过以下代码添加进观察者队列

` owner.getLifecycle().addObserver( this ); 复制代码`

这行代码的内部是通过 ` FastSafeIterableMap` 类来管理观察者的，而这个类是非线程安全的，如下：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/2/16b18a7822d7e82d?imageView2/0/w/1280/h/960/ignore-error/1) 我们来看看上面 [trello/RxLifecycle]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftrello%2FRxLifecycle ) 、 [uber/AutoDispose]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fuber%2FAutoDispose ) 、 [RxLife]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fliujingxing%2FRxLife ) 三者是如何处理这个问题的。

**trello/RxLifecycle** RxLifecycle库是 ` AndroidLifecycle` 类感知生命周期，简单看看源码：

` public final class AndroidLifecycle implements LifecycleProvider < Lifecycle. Event >, LifecycleObserver { public static LifecycleProvider<Lifecycle.Event> createLifecycleProvider(LifecycleOwner owner) { return new AndroidLifecycle(owner); } private final BehaviorSubject<Lifecycle.Event> lifecycleSubject = BehaviorSubject.create(); private AndroidLifecycle (LifecycleOwner owner) { owner.getLifecycle().addObserver( this ); } //中间省略部分代码 @OnLifecycleEvent (Lifecycle.Event.ON_ANY) void onEvent (LifecycleOwner owner, Lifecycle.Event event) { lifecycleSubject.onNext(event); if (event == Lifecycle.Event.ON_DESTROY) { owner.getLifecycle().removeObserver( this ); } } } 复制代码`

可以看到，RxLifecycle是在对象创建时添加观察者，且它没有做任何处理，如果你在子线程使用，就需要额外注意了，而且它只有在页面销毁时，才会移除观察者，试想，我们在首页一般都会有非常多的请求，而这每一个请求都会有一个AndroidLifecycle对象，我们想请求结束就要回收这个对象，然而，这个对象还是观察者队列里，就导致了没办法回收，如果我们不停下拉刷新、上拉加载更多，对内存就是一个挑战。

RxLifecycle还有一个弊端时，当Activity/Fragment销毁时，始终会往下游发送一个onComplete事件，这对于在onComplete事件中有业务逻辑的同学来说，无疑是致命的打击。

**uber/AutoDispose** AutoDispose库我们看LifecycleEventsObservable类，如下

` class LifecycleEventsObservable extends Observable < Event > { //省略部分代码 @Override protected void subscribeActual (Observer<? super Event> observer) { ArchLifecycleObserver archObserver = new ArchLifecycleObserver(lifecycle, observer, eventsObservable); observer.onSubscribe(archObserver); if (!isMainThread()) { //非主线程，直接抛出异常 observer.onError( new IllegalStateException( "Lifecycles can only be bound to on the main thread!" )); return ; } lifecycle.addObserver(archObserver); //添加观察者 if (archObserver.isDisposed()) { lifecycle.removeObserver(archObserver); } } //省略部分代码 复制代码`

可以看到，AutoDispose是在事件订阅时添加观察者，并且当前非主线程时，直接抛出异常，也就说明使用AutoDispose不能在子线程订阅事件。在移除观察者方面，AutoDispose会在事件结束或者页面销毁时移除观察者，这一点要优于RxLifecycle。

**RxLife**

RxLife库我们看AbstractLifecycle类，如下：

` public abstract class AbstractLifecycle < T > extends AtomicReference < T > implements Disposable { //省略部分代码 //事件订阅时调用此方法 protected final void addObserver () throws Exception { //Lifecycle添加监听器需要在主线程执行 if (isMainThread() || !(scope instanceof LifecycleScope)) { addObserverOnMain(); } else { final Object object = mObject; AndroidSchedulers.mainThread().scheduleDirect(() -> { addObserverOnMain(); synchronized (object) { object.notifyAll(); } }); synchronized (object) { object.wait(); } } } //省略部分代码 } 复制代码`

可以看到，RxLife对子线程做了额外的操作，在子线程通过同步锁，添加完观察者后再往下走，且RxLife同样会在事件结束或者页面销毁时移除观察者。

**我的疑问** 我们知道对View添加OnAttachStateChangeListener监听器是线程安全的，如下：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/2/16b18a782d654426?imageView2/0/w/1280/h/960/ignore-error/1) 那为何AutoDispose库中的DetachEventCompletable依然会线程做判断？代码如下 ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/2/16b18a782ec62e26?imageView2/0/w/1280/h/960/ignore-error/1) 请大神为我解答。

### 小彩蛋 ###

RxLife类里面的as系列方法，皆适用于Observable、Flowable、ParallelFlowable、Single、Maybe、Completable这6个被观察者对象，道理都一样，这里不在一一讲解。 另外，在Activity/Fragment上，如果你想在某个生命周期方法中断管道，可使用 ` as` 操作符的重载方法，如下：

` //在Activity/Fragment上 Observable.interval( 1 , 1 , TimeUnit.SECONDS) //隔一秒发送一条消息..as(RxLife.as( this , Event.ON_STOP)) //在onStop方法中断管道.subscribe(aLong -> { Log.e( "LJX" , "accept=" + aLong); }); 复制代码`

此时如果你还想在主线程回调观察者，使用 ` asOnMain` 方法即可，如下：

` //在Activity/Fragment上 Observable.interval( 1 , 1 , TimeUnit.SECONDS) //隔一秒发送一条消息.as(RxLife.asOnMain( this , Event.ON_STOP)) //在onStop方法中断管道,并在主线程回调观察者.subscribe(aLong -> { Log.e( "LJX" , "accept=" + aLong); }); //等同于 //在Activity/Fragment上 Observable.interval( 1 , 1 , TimeUnit.SECONDS) //隔一秒发送一条消息.observeOn(AndroidSchedulers.mainThread()) .as(RxLife.as( this , Event.ON_STOP)) //在onStop方法中断管道,并在主线程回调观察者.subscribe(aLong -> { Log.e( "LJX" , "accept=" + aLong); }); 复制代码`

### 小结 ###

在Activity/Fragment/View中，我们可以直接使用 ` RxLife.as(this)` ， 而在ViewModel及任意类我们需要做些准备工作，ViewModel继承ScopeViewModel，任意类继承BaseScope类就可以使用 ` RxLife.as(this)` 。

**注:一定要使用ViewModelProviders获取ViewModel对象，如下**

` //在Activity/Fragment上 MyViewModel viewModel = ViewModelProviders.of( this ).get(MyViewModel.class) 复制代码`

本人水平有限，如文章中有见解不到之处，请广大读者指正，RxLife刚出来不久，使用过程中如有遇到问题，请在github上留言，当然欢迎您加群讨论 ` RxHttp&RxLife 交流群: 378530627` 如果你觉得RxLife不错，请给我点赞，好东西不应该被埋没，请让更多的人知道它。