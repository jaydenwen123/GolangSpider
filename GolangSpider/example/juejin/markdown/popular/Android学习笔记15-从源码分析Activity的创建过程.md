# Android学习笔记15-从源码分析Activity的创建过程 #

## 问题 ##

本节要思考地问题 :

* 

系统内部是如何启动一个Acitivity的 ?

* 

新的Activity对象是何时创建的？

* 

Acitivity的onCreate()方法何时被系统回调的？

让我们带着这些问题来学习Activity的创建启动过程.

## 一 , Activty概述: ##

* 

一种展示性组件，用来向用户展示页面，接受用户的输入与之交互。、

* 

` Activity` 是由 Intent启动，而 Intent 分为 **显示Intent** 和 **隐式Intent** ，显示Intent是直接明确指定我想要启动的另一个活动，隐式Intent则可能是指向多个其他 ` Activity` ,比如我们在使用qq聊天时，打开输入框中的拍照功能，这时系统会弹出你手机中的好几个相机应用让你选择。

* 

` Activity` 有四种启动模式，不同 模式的启动方式有不同的效果。分别是 :

* Standard
* SingleTop
* SingleTask
* SingleInstance

## 二 , Activity的工作过程 ##

启动一个Activity我们最常使用的就是显示调用Intent :

` Intent intent = new Intent( this ,TestActivity.class); startActivity(intent); 复制代码`

` startActivity()` 方法在 Activity 中有其他几个重载方法，但最后都会调用 ` startActivityForResult()` 方法:

` public void startActivityForResult (Intent intent, int requestCode, @Nullable Bundle options) { if (mParent == null ) { Instrumentation.ActivityResult ar = mInstrumentation.execStartActivity( this , mMainThread.getApplicationThread(), mToken, this , intent, requestCode, options); if (ar != null ) { mMainThread.sendActivityResult( mToken, mEmbeddedID, requestCode, ar.getResultCode(), ar.getResultData()); } ... ... } 复制代码`

注意到第4行中，又调用了 ` Instrumentation.execStarAtActivity()` 方法。

从这个方法开始我们一层一层地查看，原来启动一个Activity需要这么深的方法栈。

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b2619544ca1c99?imageView2/0/w/1280/h/960/ignore-error/1)

方法栈的最后由 ` ApplicationThread.scheduleLaunchActivity()` 调用，这个类属于 ` ActivityThread` 的内部类，不太好找。贴一张 ApplicationThread 的类图

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b2619549531511?imageView2/0/w/1280/h/960/ignore-error/1)
结构是不是清晰了许多，我们再打开 这个 ` scheduleLaunchActivity()` 方法一探究竟 ，
哦对了,这个ApplicationThread是定义在ActivityThread.java中的，没想到吧。

ApplicationThread.scheduleLaunchAcitivity():

` // we use token to identify this activity without having to send the // activity itself back to the activity manager. (matters more with ipc) @Override public final void scheduleLaunchActivity (Intent intent, IBinder token, int ident, ActivityInfo info, Configuration curConfig, Configuration overrideConfig, CompatibilityInfo compatInfo, String referrer, IVoiceInteractor voiceInteractor, int procState, Bundle state, PersistableBundle persistentState, List<ResultInfo> pendingResults, List<ReferrerIntent> pendingNewIntents, boolean notResumed, boolean isForward, ProfilerInfo profilerInfo) { updateProcessState(procState, false ); ActivityClientRecord r = new ActivityClientRecord(); r.token = token; r.ident = ident; r.intent = intent; r.referrer = referrer; r.voiceInteractor = voiceInteractor; r.activityInfo = info; r.compatInfo = compatInfo; r.state = state; r.persistentState = persistentState; r.pendingResults = pendingResults; r.pendingIntents = pendingNewIntents; r.startsNotResumed = notResumed; r.isForward = isForward; r.profilerInfo = profilerInfo; r.overrideConfig = overrideConfig; updatePendingConfiguration(curConfig); sendMessage(H.LAUNCH_ACTIVITY, r); } 复制代码`

看到最后一行通过 ` sendMessage()` 发送了一条启动Activity的消息交给Handler处理,这个 Handler也是ActivityThread的一个内部类，名叫H，好简洁...

ActivityThread.H.sendMessage():

` private void sendMessage ( int what, Object obj, int arg1, int arg2, boolean async) { if (DEBUG_MESSAGES) Slog.v( TAG, "SCHEDULE " + what + " " + mH.codeToString(what) + ": " + arg1 + " / " + obj); Message msg = Message.obtain(); msg.what = what; msg.obj = obj; msg.arg1 = arg1; msg.arg2 = arg2; if (async) { msg.setAsynchronous( true ); } mH.sendMessage(msg); } 复制代码`

我们跳到名为 H 的Handler中，看看它对这个消息是怎么处理的 :

ActivityThread.H.handleMessage():

` private class H extends Handler { public static final int LAUNCH_ACTIVITY = 100 ; public static final int PAUSE_ACTIVITY = 101 ; public static final int PAUSE_ACTIVITY_FINISHING= 102 ; ... ... public void handleMessage (Message msg) { if (DEBUG_MESSAGES) Slog.v(TAG, ">>> handling: " + codeToString(msg.what)); switch (msg.what) { case LAUNCH_ACTIVITY: { Trace.traceBegin(Trace.TRACE_TAG_ACTIVITY_MANAGER, "activityStart" ); final ActivityClientRecord r = (ActivityClientRecord) msg.obj; r.packageInfo = getPackageInfoNoCheck( r.activityInfo.applicationInfo, r.compatInfo); handleLaunchActivity(r, null ); Trace.traceEnd(Trace.TRACE_TAG_ACTIVITY_MANAGER); } break ; case RELAUNCH_ACTIVITY: ... break ; case PAUSE_ACTIVITY: ... break ; case PAUSE_ACTIVITY_FINISHING: ... break ; } } } 复制代码`

在 case LAUNCH_ACTIVITY :中,调用了 **handleLaunchActivity()**

ActivityThread.handleLaunchActivity() :

` private void handleLaunchActivity (ActivityClientRecord r, Intent customIntent) { ... ... //省略部分代码 if (localLOGV) Slog.v( TAG, "Handling launch of " + r); // Initialize before creating the activity WindowManagerGlobal.initialize(); Activity a = performLaunchActivity(r, customIntent); if (a != null ) { r.createdConfig = new Configuration(mConfiguration); Bundle oldState = r.state; handleResumeActivity(r.token, false , r.isForward, !r.activity.mFinished && !r.startsNotResumed); ... ... //省略部分代码 } else { ... } } 复制代码`

第10行,由 ` performLaunchActivity(r, customIntent)` 返回一个Activity对象，完成 Activity 对象的创建和启动过程.
现在也就回答了我们开头提出的前两个问题 :

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b2619544ec6fe3?imageView2/0/w/1280/h/960/ignore-error/1)

第15行,由 ` handleResumeActivity()` 调用 Activity生命周期中的 ` onResume()` 方法。

ActivityThread.performLaunchActivity() 已经来到了最关键的步骤，这个函数主要完成了3件事情，我们拆开一步一步看。

#### 1.从参数 ` AcitivityClientRecord` 对象中获取待启动的 Activity 的信息: ####

` private Activity performLaunchActivity (ActivityClientRecord r, Intent customIntent) { //第一部分:从参数 AcitivityClientRecord 对象中获取待启动的 Activity 的信息 ActivityInfo aInfo = r.activityInfo; if (r.packageInfo == null ) { r.packageInfo = getPackageInfo(aInfo.applicationInfo, r.compatInfo, Context.CONTEXT_INCLUDE_CODE); } ComponentName component = r.intent.getComponent(); if (component == null ) { component = r.intent.resolveActivity( mInitialApplication.getPackageManager()); r.intent.setComponent(component); } if (r.activityInfo.targetActivity != null ) { component = new ComponentName(r.activityInfo.packageName, r.activityInfo.targetActivity); } } 复制代码`

#### 2.通过 Instrumentation.newActivity() 方法使用类加载器创建Activity对象 : ####

` private Activity performLaunchActivity (ActivityClientRecord r, Intent customIntent) { //第一部分:从参数 AcitivityClientRecord 对象中获取待启动的 Activity 的信息... ... //第二部分:通过Instrumentation.newActivity()方法使用类加载器创建Activity对象 Activity activity = null ; try { java.lang.ClassLoader cl = r.packageInfo.getClassLoader(); activity = mInstrumentation.newActivity( cl, component.getClassName(), r.intent); StrictMode.incrementExpectedActivityCount(activity.getClass()); r.intent.setExtrasClassLoader(cl); r.intent.prepareToEnterProcess(); if (r.state != null ) { r.state.setClassLoader(cl); } } catch (Exception e) { if (!mInstrumentation.onException(activity, e)) { throw new RuntimeException( "Unable to instantiate activity " + component + ": " + e.toString(), e); } } } 复制代码`

#### 3.通过 ` LoadedApk` 的 ` makeApplication()` 方法来尝试创建Application对象 : ####

` private Activity performLaunchActivity (ActivityClientRecord r, Intent customIntent) { //第一部分:从参数 AcitivityClientRecord 对象中获取待启动的 Activity 的信息... ... //第二部分:通过Instrumentation.newActivity()方法使用类加载器创建Activity对象... //第三部分:通过LoadedApk的makeApplication() 方法来尝试创建Application对象 try { Application app = r.packageInfo.makeApplication( false , mInstrumentation); } ... } 复制代码`

r.packeageInfo 是一个 LoadApk对象,用来加载 apk文件,进来看一下 makeApplication()方法 :

LoadApk.makeApplication():

` /** * Local state maintained about a currently loaded .apk. * @hide */ public final class LoadedApk { ... public Application makeApplication ( boolean forceDefaultAppClass, Instrumentation instrumentation) { if (mApplication != null ) { return mApplication; } Application app = null ; ... try { java.lang.ClassLoader cl = getClassLoader(); if (!mPackageName.equals( "android" )) { initializeJavaContextClassLoader(); } ContextImpl appContext = ContextImpl.createAppContext(mActivityThread, this ); app = mActivityThread.mInstrumentation.newApplication( cl, appClass, appContext); appContext.setOuterContext(app); } catch (Exception e) { if (!mActivityThread.mInstrumentation.onException(app, e)) { throw new RuntimeException( "Unable to instantiate application " + appClass + ": " + e.toString(), e); } } mActivityThread.mAllApplications.add(app); mApplication = app; ... ... return app; } } 复制代码`

#### 4.创建Context对象，调用activity的attach方法来注入一些重要数据，进行activity的初始化 ####

` private Activity performLaunchActivity (ActivityClientRecord r, Intent customIntent) { //第一部分:从参数 AcitivityClientRecord 对象中获取待启动的 Activity 的信息... ... //第二部分:通过Instrumentation.newActivity()方法使用类加载器创建Activity对象... //第三部分:通过LoadedApk的makeApplication() 方法来尝试创建Application对象... //第四部分:创建Context对象，调用activity的attach方法来注入一些重要数据，进行activity的初始化 if (activity != null ) { Context appContext = createBaseContextForActivity(r, activity); CharSequence title = r.activityInfo.loadLabel(appContext.getPackageManager()); Configuration config = new Configuration(mCompatConfiguration); if (DEBUG_CONFIGURATION) Slog.v(TAG, "Launching activity " + r.activityInfo.name + " with config " + config); activity.attach(appContext, this , getInstrumentation(), r.token, r.ident, app, r.intent, r.activityInfo, title, r.parent, r.embeddedID, r.lastNonConfigurationInstances, config, r.referrer, r.voiceInteractor); } } 复制代码`

相信大家也看出来了，我们平常使用的Context(App环境的上下文)就是通过 activity.attach() 方法与 Activity组件建立起联系的。

进入 attach()方法中,我们还看到了 window也是在这里进行初始化，并与 Activity 建立关联,这样当Window接收外部的输入事件之后就可以把事件传给Activity.

Activity.attach() :

` final void attach (Context context, ActivityThread aThread, Instrumentation instr, IBinder token, int ident, Application application, Intent intent, ActivityInfo info, CharSequence title, Activity parent, String id, NonConfigurationInstances lastNonConfigurationInstances, Configuration config, String referrer, IVoiceInteractor voiceInteractor) { attachBaseContext(context); mFragments.attachHost( null /*parent*/ ); mWindow = new PhoneWindow( this ); mWindow.setCallback( this ); mWindow.setOnWindowDismissedCallback( this ); mWindow.getLayoutInflater().setPrivateFactory( this ); if (info.softInputMode != WindowManager.LayoutParams.SOFT_INPUT_STATE_UNSPECIFIED) { mWindow.setSoftInputMode(info.softInputMode); } if (info.uiOptions != 0 ) { mWindow.setUiOptions(info.uiOptions); } mUiThread = Thread.currentThread(); mMainThread = aThread; mInstrumentation = instr; mToken = token; mIdent = ident; mApplication = application; mIntent = intent; mReferrer = referrer; mComponent = intent.getComponent(); mActivityInfo = info; mTitle = title; mParent = parent; mEmbeddedID = id; mLastNonConfigurationInstances = lastNonConfigurationInstances; if (voiceInteractor != null ) { if (lastNonConfigurationInstances != null ) { mVoiceInteractor = lastNonConfigurationInstances.voiceInteractor; } else { mVoiceInteractor = new VoiceInteractor(voiceInteractor, this , this , Looper.myLooper()); } } mWindow.setWindowManager( (WindowManager)context.getSystemService(Context.WINDOW_SERVICE), mToken, mComponent.flattenToString(), (info.flags & ActivityInfo.FLAG_HARDWARE_ACCELERATED) != 0 ); if (mParent != null ) { mWindow.setContainer(mParent.getWindow()); } mWindowManager = mWindow.getWindowManager(); mCurrentConfig = config; } 复制代码`

#### 5.调用Activity的onCreate(),onStart()等方法: ####

` private Activity performLaunchActivity (ActivityClientRecord r, Intent customIntent) { //第一部分:从参数 AcitivityClientRecord 对象中获取待启动的 Activity 的信息... ... //第二部分:通过Instrumentation.newActivity()方法使用类加载器创建Activity对象... //第三部分:通过LoadedApk的makeApplication() 方法来尝试创建Application对象... //第四部分:创建Context对象，调用activity的attach方法来注入一些重要数据，进行activity的初始化... //第五部分:调用Activity的onCreate()方法 activity.mCalled = false ; if (r.isPersistable()) { mInstrumentation.callActivityOnCreate(activity, r.state, r.persistentState); } else { mInstrumentation.callActivityOnCreate(activity, r.state); } if (!activity.mCalled) { throw new SuperNotCalledException( "Activity " + r.intent.getComponent().toShortString() + " did not call through to super.onCreate()" ); } r.activity = activity; r.stopped = true ; if (!r.activity.mFinished) { //调用activity的onStart()方法 activity.performStart(); r.stopped = false ; } if (!r.activity.mFinished) { if (r.isPersistable()) { if (r.state != null || r.persistentState != null ) { //调用activity的onRestoreInstanceState()方法 mInstrumentation.callActivityOnRestoreInstanceState(activity, r.state, r.persistentState); } } else if (r.state != null ) { mInstrumentation.callActivityOnRestoreInstanceState(activity, r.state); } } } 复制代码`

至此，Activity的启动与创建，再到onCreate()方法的调用，我们都已经分析完。

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b261953aaaa028?imageView2/0/w/1280/h/960/ignore-error/1)

#### Reference : ####

[《Android艺术探索》]( https://link.juejin.im?target=https%3A%2F%2Fbook.douban.com%2Fsubject%2F26599538%2F )

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b261953f65dd7a?imageView2/0/w/1280/h/960/ignore-error/1)

(完~) 谢谢大家浏览

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2619546c0e9b3?imageView2/0/w/1280/h/960/ignore-error/1)