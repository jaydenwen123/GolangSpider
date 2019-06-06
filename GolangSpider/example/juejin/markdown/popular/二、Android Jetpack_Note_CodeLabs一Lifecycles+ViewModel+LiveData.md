# 二、Android Jetpack_Note_CodeLabs一Lifecycles+ViewModel+LiveData #

# 一、简介 #

本篇文章向您介绍了用于构建Android应用程序的以下生命周期感知架构组件：

* **ViewModel** - 提供了一种创建和检索绑定到特定生命周期的对象的方法。A ViewModel通常存储视图数据的状态，并与其他组件通信，例如数据存储库或处理业务逻辑的域层。要阅读本主题的介绍性指南，请参阅 [ViewModel]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Ftopic%2Flibraries%2Farchitecture%2Fviewmodel.html ) 。
* **LifecycleOwner** / **LifecycleRegistryOwner** -无论是LifecycleOwner和LifecycleRegistryOwner是在实现的接口AppCompatActivity和Support Fragment类。您可以将其他组件订阅到实现这些接口的所有者对象，以观察对所有者生命周期的更改。要阅读本主题的介绍性指南，请参阅 [Hanlde Lifecycles]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Ftopic%2Flibraries%2Farchitecture%2Flifecycle.html ) 。
* **LiveData** - 允许您观察应用程序的多个组件之间的数据更改，而无需在它们之间创建明确，严格的依赖关系路径。LiveData尊重应用程序组件的复杂生命周期，包括活动，片段，服务或LifecycleOwner应用程序中定义的任何内容。LiveData通过暂停对已停止LifecycleOwner对象的订阅以及取消对LifecycleOwner已完成对象的订阅来管理观察者订阅。要阅读本主题的介绍性指南，请参阅 [LiveData]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Ftopic%2Flibraries%2Farchitecture%2Flivedata.html ) 。

## 1.1 你要建造什么 ##

在此代码框中，您可以实现上述每个组件的示例。首先是示例应用程序，然后通过一系列步骤添加代码，在进度时集成各种体系结构组件。

## 1.2 你需要什么 ##

* Android Studio 3.3或更高版本
* 熟悉Android的生命周期

# 二、环境搭建 #

下载代码

` git clone git@github.com:googlecodelabs / android-lifecycles.git 复制代码`

运行程序：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a60f5b7cb60?imageView2/0/w/1280/h/960/ignore-error/1)

旋转屏幕，注意计时器重置

> 
> 
> 
> 如果我们按照之前的逻辑需要保存状态，然后恢复状态，并恢复计时器。现在您可以用 [ViewModel](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Freference%2Fandroid%2Farch%2Flifecycle%2FViewModel.html
> ) ，因为因为此类的实例在配置更改后仍然存在，例如屏幕旋转。
> 
> 

# 三、添加ViewModel #

在此步骤中，您使用ViewModel来跨屏幕旋转保持状态并解决您在上一步中观察到的行为。在上一步中，您运行了一个显示计时器的活动。当配置更改（例如屏幕旋转）破坏活动时，将重置此计时器。

可以使用ViewModel在活动或片段的整个生命周期中保留数据。如前一步所示，活动是管理应用数据的不良选择。活动和片段是短暂的对象，当用户与应用程序交互时，这些对象会频繁创建和销毁。 ViewModel还更适合管理与网络通信相关的任务，以及数据操作和持久性。

## 3.1 使用ViewModel保持计时器的状态 ##

打开ChronoActivity2并检查类如何检索和使用 **ViewModel** ：

` ChronometerViewModel chronometerViewModel = ViewModelProviders.of( this ). get (ChronometerViewModel. class ); 复制代码`
> 
> 
> 
> 
> this指的是一个实例LifecycleOwner。ViewModel只要活动范围存在，框架就会保持LifecycleOwner活力。ViewModel如果其所有者因配置更改（例如屏幕旋转）而被销毁，则不会销毁A.
> 所有者的新实例重新连接到现有实例，ViewModel如下图所示：
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a9b7854bbec?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> Activity或Fragment的范围从创建到完成（或终止），您不得与销毁混淆。请记住，当旋转设备时，活动会被销毁，但ViewModel与之关联的任何实例都不会被销毁。
> 
> 
> 

## 3.1 试一下 ##

运行应用程序（在“运行配置”下拉列表中选择“ 步骤2 ”），并在执行以下任一操作时确认计时器未重置：

* 旋转屏幕。 2。 导航到另一个应用程序，然后返回。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18aad36e3d292?imageView2/0/w/1280/h/960/ignore-error/1)

但是，如果您或系统退出应用程序，则计时器将重置。

> 
> 
> 
> 系统会ViewModel在生命周期所有者的整个生命周期中保留内存中的实例，例如片段或活动。系统不会持久ViewModel存储长期存储。
> 
> 

# 四、使用LiveData包装数据 #

在此步骤中，您将使用a的自定义计时器替换前面步骤中使用的计时器Timer，并每秒更新UI。 Timer是一个java.util可用于在将来循环安排任务的类。您将此逻辑添加到LiveDataTimerViewModel类中，并使活动专注于管理用户和UI之间的交互。

当计时器通知时，活动会更新UI。为了避免内存泄漏，ViewModel不包括对活动的引用。例如，配置更改（例如屏幕旋转）可能会导致对ViewModel应该进行垃圾回收的活动的引用。系统将保留实例，ViewModel直到相应的活动或生命周期所有者不再存在。

> 
> 
> 
> **注意** ：存储到一个参考上下文或视图中ViewModel可能会导致内存泄漏。避免使用引用Context或View类实例的字段。所述onCleared（）方法是用于清除引用有用退订或明确的引用与长周期其它的目的，但不Context或View对象。
> 
> 
> 

ViewModel您可以将活动或片段配置为观察数据源，而不是直接从其中修改视图，并在数据更改时接收数据。这种安排称为观察者模式。

> 
> 
> 
> **注意** ：要将数据公开为可观察对象，请将类型包装在LiveData类中。
> 
> 

如果您使用了数据绑定库或其他反应库（如RxJava），您可能熟悉观察者模式。LiveData是一个特殊的可观察类，它是生命周期感知的，只通知活跃的观察者。

## 4.1 LifecycleOwner ##

**ChronoActivity3** 是一个实例LifecycleActivity，它可以提供生命周期的状态。这是类声明：

` public class LifecycleActivity extends FragmentActivity implements LifecycleRegistryOwner {...} 复制代码`

将LifecycleRegistryOwner用于的实例的生命周期结合ViewModel并LiveData与活动或片段。片段的等价类是LifecycleFragment。

## 4.2 更新ChronoActivity ##

1.将以下代码添加到方法中的ChronoActivity3类中subscribe()以创建订阅：

` mLiveDataTimerViewModel.getElapsedTime().observe( this , elapsedTimeObserver); 复制代码`

2.接下来，在LiveDataTimerViewModel类中设置新的经过时间值。找到以下注释

` //TODO set the new value 复制代码`

用以下语句替换注释:

` mElapsedTime.postValue(newValue); 复制代码`

3.运行应用程序并在Android Studio中打开Android Monitor。请注意，除非您导航到另一个应用程序，否则日志会每秒更新。如果您的设备支持多窗口模式，您可能想尝试使用它。旋转屏幕不会影响应用的行为方式。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18ae8af7f3822?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> **注意** ：LiveData对象仅在活动时发送更新，或者处于LifecycleOwner活动状态。如果您导航到其他应用程序，日志消息会暂停，直到您返回。LiveData对象仅在其各自的生命周期所有者为STARTED或时将订阅视为活动的RESUMED。
> 
> 
> 

# 五、订阅生命周期事件 #

许多Android组件和库要求您：

* 订阅或初始化组件或库。
* 取消订阅或停止组件或库。

未能完成上述步骤可能会导致内存泄漏和细微错误。

可以将生命周期所有者对象传递给生命周期感知组件的新实例，以确保它们了解生命周期的当前状态。

您可以使用以下语句查询生命周期的当前状态：

` lifecycleOwner.getLifecycle().getCurrentState() 复制代码`

上面的语句返回一个状态，例如Lifecycle.State.RESUMED，或Lifecycle.State.DESTROYED。

实现的生命周期感知对象LifecycleObserver还可以观察生命周期所有者状态的变化：

` lifecycleOwner.getLifecycle().addObserver( this ); 复制代码`

您可以注释对象以指示它在需要时调用适当的方法：

` @OnLifecycleEvent(Lifecycle.EVENT.ON_RESUME) void addLocationListener() { ... } 复制代码`

## 5.1 创建一个支持生命周期的组件 ##

在此步骤中，您将创建一个对活动生命周期所有者作出反应的组件。使用片段作为生命周期所有者时，类似的原则和步骤适用。

您使用Android框架LocationManager获取当前的纬度和经度并将其显示给用户。此添加允许您：

* 订阅更改并使用自动更新UI LiveData。
* LocationManager根据活动状态的更改，创建注册和取消注册的包装器。
* 

您通常会订阅活动或方法中的LocationManager更改，并删除或方法中的侦听器： **onStart()** **onResume()** **onStop()** **onPause()**

` // Typical use, within an activity. @Override protected void onResume() { mLocationManager.requestLocationUpdates(LocationManager.GPS_PROVIDER, 0 , 0 , mListener); } @Override protected void onPause() { mLocationManager.removeUpdates(mListener); } 复制代码`

在此步骤中，您将在LifecycleOwner名为LifecycleRegistryOwner的类中使用被调用的实现BoundLocationManager。BoundLocationManager类的名称指的是类的实例绑定到活动的生命周期。

要让类观察活动的生命周期，必须将其添加为观察者。要实现此BoundLocationManager目的，请通过将以下代码添加到其构造函数来指示对象观察生命周期：

` lifecycleOwner.getLifecycle().addObserver( this ); 复制代码`

要在发生生命周期更改时调用方法，可以使用@OnLifecycleEvent注释。使用类中的以下注释更新addLocationListener()和removeLocationListener()方法BoundLocationListener：

` @OnLifecycleEvent(Lifecycle.Event.ON_RESUME) void addLocationListener() { ... } @OnLifecycleEvent(Lifecycle.Event.ON_PAUSE) void removeLocationListener() { ... } 复制代码`
> 
> 
> 
> 
> **注意** ：观察者被带到提供者的当前状态，因此不需要addLocationListener()从构造函数调用。当观察者被添加到生命周期所有者时，它会被调用。
> 
> 
> 

旋转设备时，运行应用程序并验证日志监视器是否显示以下操作：

` D / BoundLocationMgr：添加了监听器 D / BoundLocationMgr：已删除侦听器 D / BoundLocationMgr：添加了监听器 D / BoundLocationMgr：已删除侦听器 复制代码`

使用Android模拟器模拟更改设备的位置（单击三个点以显示扩展控件）。在TextView当它改变时被更新：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b2a2bc1eeb1?imageView2/0/w/1280/h/960/ignore-error/1)

# 六、在Fragment之间共享ViewModel #

## 6.1 在Fragment之间共享ViewModel ##

使用ViewModel完成以下附加步骤在Fragment之间的通信：

* 一项活动。
* 一个片段的两个实例，每个都有一个SeekBar。
* 单个ViewModel带LiveData字段。
* 

运行此步骤并注意其中两个实例SeekBar彼此独立：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b463e106c6a?imageView2/0/w/1280/h/960/ignore-error/1)

使用ViewModel以便在SeekBar更改Fragment时SeekBar更新另一个Fragment：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b4eee975110?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 注意：您应该将活动用作生命周期所有者，因为每个Fragment的生命周期是独立的。
> 
> 

# 七、跨进程保持ViewModel状态（实验版） #

> 
> 
> 
> 本节中使用的模块处于alpha阶段。这意味着API不是最终的，将来可能会发生变化。
> 
> 

从内存管理概述：

> 
> 
> 
> 当用户在应用程序之间切换时，Android会保留非前景的应用程序 -
> 即用户不可见或在最近最少使用（LRU）缓存中运行音乐播放等前台服务。例如，当用户首次启动应用程序时，会为其创建一个进程;
> 但是当用户离开应用程序时，该进程不会退出。系统会保持进程缓存。如果用户稍后返回应用程序，系统将重新使用该过程，从而使应用程序切换更快。
> 
> 

由于系统内存不足，它会从最近最少使用的进程开始杀死缓存中的进程。当用户导航回应用程序时，系统将在新进程中重新启动应用程序。

由于只有在用户暂时没有与应用程序交互时才会发生这种情况，因此可能允许他们返回应用程序并在初始状态下找到它。但是，在某些情况下，您可能希望保存应用程序的状态或部分应用程序的状态，以便在进程被杀死时不会丢失该信息。

该 lifecycle-viewmodel-savedstate模块提供对ViewModel中已保存状态的访问。

该模块的gradle依赖是

` "androidx.lifecycle:lifecycle-viewmodel-savedstate: $savedStateVersion " 复制代码`

ViewModels需要两次更改才能访问已保存的状态：

* 传递SavedStateVMFactory到ViewModelProvider
* 添加一个接收a的构造函数 SavedStateHandle

首先，让我们在没有这些变化的情况下尝试第6步：

1.打开运行配置“步骤6”

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b6bf98b167a?imageView2/0/w/1280/h/960/ignore-error/1)

你会看到一个简单的表单：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b70b49081fa?imageView2/0/w/1280/h/960/ignore-error/1)

2.更改名称，然后单击“保存”。这将把它存储在ViewModel内的LiveData中。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b730ea2e3ff?imageView2/0/w/1280/h/960/ignore-error/1)

3.模拟系统终止进程（需要运行P +的仿真器）。首先输入以下命令确保进程正在运行：

` $ adb shell ps -A | grep lifecycle 复制代码`

这应该输出带有名称的运行进程 com.example.android.codelabs.lifecycle

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b795f2ccd4a?imageView2/0/w/1280/h/960/ignore-error/1)

在设备或模拟器上按Home，然后运行

` $ adb shell am kill com.example.android.codelabs.lifecycle 复制代码`

如果你再输入一次

` $ adb shell ps -A | grep lifecycle 复制代码`

你应该什么也得不到，表明这个过程已被正确杀死。

4.再次打开应用程序（在应用程序启动器中查找LC Step6）。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18b83cc16a5f6?imageView2/0/w/1280/h/960/ignore-error/1)

ViewModel中的值未保留，但已EditText恢复其状态。这怎么可能？

> 
> 
> 
> 一些UI元素，包括EditText使用自己的onSaveInstanceState实现保存其状态。在进程被终止后，此状态将在配置更改后恢复时恢复。阅读ViewModels：Persistence，onSaveInstanceState（），恢复UI状态和加载器以获取更多信息。
> 
> 
> 

实际上，该lifecycle-viewmodel-savedstate模块还使用onSaveInstanceState和onRestoreInstanceState保持ViewModel状态，但它使这些操作更方便。

**为ViewModel实现保存的状态** 在SavedStateActivity.java文件中，替换

` mSavedStateViewModel = ViewModelProviders.of（ this ）. get （SavedStateViewModel. class ）; 复制代码`

为：

` mSavedStateViewModel = ViewModelProviders.of( this , new SavedStateVMFactory( this )) . get (SavedStateViewModel. class ); 复制代码`

在SavedStateViewModel.java文件中，您需要添加一个新的构造函数，该构造函数SavedStateHandle将状态存储在私有字段中：

` private SavedStateHandle mState; public SavedStateViewModel（SavedStateHandle savedStateHandle）{ mState = savedStateHandle; } 复制代码`

现在您将使用该模块的LiveData支持，因此您不再需要存储它：

` private static final String NAME_KEY =“name”; //公开不可变的LiveData LiveData <String> getName（）{ return mState.getLiveData（NAME_KEY）; } void saveNewName（String newName）{ mState. set （NAME_KEY，newName）; } 复制代码`

现在您正在使用LiveData mState，MutableLiveData名称不再使用，可以删除。

现在您可以再次尝试相同的过程。打开应用程序，更改名称并保存。然后，按Home并使用以下命令终止该过程：

` $ adb shell am kill com.example.android.codelabs.lifecycle 复制代码`

如果您重新打开该应用程序，您将看到ViewModel中的状态此时已保存。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18ba3b1907549?imageView2/0/w/1280/h/960/ignore-error/1)