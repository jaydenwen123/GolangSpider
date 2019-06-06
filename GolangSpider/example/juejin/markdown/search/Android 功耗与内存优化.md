# Android 功耗与内存优化 #

## 主要内容 ##

* 功耗优化：关于一些对功耗的检测及优化
* 内存优化：关于一些对内存的常见优化手段及检测工具

## 功耗优化 ##

### 耗电原因 ###

* CPU：wakelocks
* 网络、无线、蓝牙等
* 屏幕

CPU 与网络等是属于 Background Process，若要优化，需要设法减少、延迟、合并 Background Process。

### Doze 与 App Standby ###

从 API 23 开始， Android 提供了这两种模式以延长电池寿命，是 Android 系统的大方针。Doze 就是将一些 Wakelocks、网络、Jobs/Syncs、Alarms、GPS/Wifi 禁止/延迟，仅在需要唤醒这些东西时即到维护窗口时才唤醒。App Standby 是应用级别的，细分到每一个应用中的，应用未使用一段时间后(即没有 foreground process 时)就会进入 App Standby，进入此状态后 App 就不能访问网络并且 Jobs/Syncs 都会受到延迟。

当然，启用 Foreground Service 就不会受到 Doze 和 App Standby 的影响。同时，在 AOSP 中这两个模式是关闭的，具体需要看看手机厂商有没有开启，提起这个只是给到大家一个官方的方案来引导我们做自己的方案。

[developer.android.google.cn/training/mo…]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.google.cn%2Ftraining%2Fmonitoring-device-state%2Fdoze-standby.html )

### 常见优化方案 ###

若要进行功耗优化，则需要设法减少、延迟、合并 Background Process。通过 Doze 则是对 Background Process 进行合并，即合并到维护窗口时；通过 App Standby 则是对 Backgroud Process 进行延迟，即延迟到"使用"状态下或充电状态下。这些是 系统级别的优化 ，而应用级别的优化则是在平时的积累中进行的，比如对耗电功能进行优化，原则上与上面一致，即 减少、延迟、合并 这些功耗大的功能：

* 屏幕：屏幕亮度对功耗影响是比较大的，在一些场景(如二维码展示、视频播放、游戏)中会对屏幕亮度进行调整或是保持屏幕常亮，在这种情况下是比较耗电的，注意对此方面的控制能够对功耗做到一定的优化。
* 充电状态下进行耗电操作：通过监听电量广播(Intent.ACTION_BATTERY_CHANGED)即可获取当前充电状态，在充电状态下进行耗电操作(比如自动备份等高网络请求的功能)。获取充电状态的代码如下：

` /** * This method checks for power by comparing the current battery state against all possible * plugged in states. In this case, a device may be considered plugged in either by USB, AC, or * wireless charge. (Wireless charge was introduced in API Level 17.) */ private boolean checkForPower () { // It is very easy to subscribe to changes to the battery state, but you can get the current // state by simply passing null in as your receiver. Nifty, isn't that? IntentFilter filter = new IntentFilter(Intent.ACTION_BATTERY_CHANGED); Intent batteryStatus = this.registerReceiver( null , filter); // There are currently three ways a device can be plugged in. We should check them all. int chargePlug = batteryStatus.getIntExtra(BatteryManager.EXTRA_PLUGGED, - 1 ); boolean usbCharge = (chargePlug == BatteryManager.BATTERY_PLUGGED_USB); boolean acCharge = (chargePlug == BatteryManager.BATTERY_PLUGGED_AC); boolean wirelessCharge = false ; if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN_MR1) { wirelessCharge = (chargePlug == BatteryManager.BATTERY_PLUGGED_WIRELESS); } return (usbCharge || acCharge || wirelessCharge); } 复制代码`

* 延迟、合并高功耗操作：在网络请求中，频繁请求会比集中请求功耗大，蜂窝数据请求会比无线网络请求功耗大。这种情况下可通过 **JobScheduler** 将一些操作延迟、合并在一些合适的情况下进行。
* 传感器、GPS：尽量复用、及时注销、选择正确的参数
* 及时释放 WakeLock

### 检测工具 Battery Historian ###

Battery Historian 是 Android 5.0 之后引入的一个获取设备电量消耗信息的图形化工具，能够直观展示手机电量消耗。

![img](https://user-gold-cdn.xitu.io/2019/6/5/16b26e4728d8bb9d?imageView2/0/w/1280/h/960/format/png/ignore-error/1)

[Battery_Historian_Tool使用说明]( https://link.juejin.im?target=https%3A%2F%2Fwwmmyy.github.io%2F2016%2F12%2F14%2FBattery_Historian_Tool%25E4%25BD%25BF%25E7%2594%25A8%25E8%25AF%25B4%25E6%2598%258E%2F )

### 推荐使用 API ###

JobScheduler：系统利用这些触发的设置，合并相同的 background process，从而优化内存和电池性能。

## 内存优化 ##

### 内存优化原因 ###

* 内存泄漏 -> 内存占用高 -> OOM

### ART GC 原因 ###

* Concurrent： 并发GC，不会使App的线程暂停，该GC是在后台线程运行的，并不会阻止内存分配。
* Alloc：当堆内存已满时，App尝试分配内存而引起的GC，这个GC会发生在正在分配内存的线程。
* Explicit：App显示的请求垃圾收集，例如调用System.gc()。与DVM一样，最佳做法是应该信任GC并避免显示的请求GC，显示的请求GC会阻止分配线程并不必要的浪费 CPU 周期。如果显式的请求GC导致其他线程被抢占，那么有可能会导致 jank（App同一帧画了多次)。
* NativeAlloc：Native内存分配时，比如为Bitmaps或者RenderScript分配对象， 这会导致Native内存压力，从而触发GC。
* CollectorTransition：由堆转换引起的回收，这是运行时切换GC而引起的。收集器转换包括将所有对象从空闲列表空间复制到碰撞指针空间（反之亦然）。当前，收集器转换仅在以下情况下出现：在内存较小的设备上，App将进程状态从可察觉的暂停状态变更为可察觉的非暂停状态（反之亦然）。
* HomogeneousSpaceCompact：齐性空间压缩是指空闲列表到压缩的空闲列表空间，通常发生在当App已经移动到可察觉的暂停进程状态。这样做的主要原因是减少了内存使用并对堆内存进行碎片整理。
* DisableMovingGc：不是真正的触发GC原因，发生并发堆压缩时，由于使用了 GetPrimitiveArrayCritical，收集会被阻塞。一般情况下，强烈建议不要使用 GetPrimitiveArrayCritical，因为它在移动收集器方面具有限制。
* HeapTrim：不是触发GC原因，但是请注意，收集会一直被阻塞，直到堆内存整理完毕。

### 检测工具 ###

* LeakCanary： [github.com/square/leak…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsquare%2Fleakcanary%25EF%25BC%258C%25E6%25AF%2594%25E8%25BE%2583%25E7%25AE%2580%25E5%258D%2595%25EF%25BC%258C%25E4%25BD%25BF%25E7%2594%25A8%25E6%2597%25B6%25E5%2585%25B3%25E9%2597%25AD%25E4%25BB%25A3%25E7%25A0%2581%25E6%25B7%25B7%25E6%25B7%2586 )
* Memory Profiler：Android Studio 自带工具https://developer.android.com/studio/profile/memory-profiler.html?hl=zh-cn

### 常见的内存泄漏场景 ###

* 非静态内部类的静态实例：非静态内部类会持有外部类的引用。
* 匿名内部类的静态实例：匿名内部类会持有外部类的引用。
* Handler：直接 new Handler(){} 是匿名内部类，也会持有外部类引用，同时 Message Queue 中的部分消息不会及时消费掉，因此：1、在销毁 Activity 时及时清除 MessageQueue 中的 Message；2、使用静态的 Handler 内部类；3、持有弱引用的 Activity。
* Context：避免在非必要的情况下使用 Activity Context，使用 Application Context 替代。
* 对有可能持有 Activity 引用的静态类(如静态 View)需要在销毁时置空。
* 未关闭资源对象(File、Cursor、MediaPlayer等)。
* Bitmap：Bitmap 的管理是非常麻烦的，使用其都需非常小心，临时创建需及时回收，避免静态变量持有 Bitmap 对象。
* 监听器及时 remove/unregister（传感器、定位等）。

### 其他内存优化手段 ###

除避免内存泄漏外还有其他内存优化手段：

* 

用 JobScheduler 替代 Service。若必须使用 Service，最好用 IntentService 限制服务寿命，所有请求完成后会自动停止。

* 

使用 SparseArray、SparseBooleanArray、LongSparseArray 代替 HashMap 等数据结构。

* 

必要时释放内存：

` import android.content.ComponentCallbacks2; public class MainActivity extends AppCompatActivity implements ComponentCallbacks2 { /** * 当UI不可见时或系统资源不足时释放内存 * @param level 引发的与内存相关的事件 */ public void onTrimMemory ( int level) { // 确定引发了哪个生命周期或系统事件 switch (level) { case ComponentCallbacks2.TRIM_MEMORY_UI_HIDDEN: /* 释放一些没必要的内存资源。现在用户界面已移至后台 */ break ; case ComponentCallbacks2.TRIM_MEMORY_RUNNING_MODERATE: case ComponentCallbacks2.TRIM_MEMORY_RUNNING_LOW: case ComponentCallbacks2.TRIM_MEMORY_RUNNING_CRITICAL: /* 释放应用不需要的内存。 应用程序运行时，设备内存不足。 引发的事件表示与内存相关的事件的严重性。 如果事件是TRIM_MEMORY_RUNNING_CRITICAL，那么系统将开始杀死后台进程。 */ break ; case ComponentCallbacks2.TRIM_MEMORY_BACKGROUND: case ComponentCallbacks2.TRIM_MEMORY_MODERATE: case ComponentCallbacks2.TRIM_MEMORY_COMPLETE: /* 释放尽可能多的内存。 该应用程序位于LRU列表中并且系统内存不足。 引发的事件表明应用程序位于LRU列表中的位置。 如果引发的事件是 TRIM_MEMORY_COMPLETE, 这一进程将是最先终止的进程之一。 */ break ; default : /* 释放任何非关键数据结构。 应用程序从系统接收到一个无法识别的内存级别值。 将其视为一般的低内存消息。 */ break ; } } } 复制代码`
* 

检查当前使用了多少内存，做好保护去释放一些次要资源以防 OOM：

` public void doSomethingMemoryIntensive () { // 在做一些需要大量内存的事情之前检查设备是否处于低内存状态 ActivityManager.MemoryInfo memoryInfo = getAvailableMemory(); if (!memoryInfo.lowMemory) { // 当前处于低内存状态 ... } } // 获取 MemoryInfo private ActivityManager. MemoryInfo getAvailableMemory () { ActivityManager activityManager = (ActivityManager) this.getSystemService(ACTIVITY_SERVICE); ActivityManager.MemoryInfo memoryInfo = new ActivityManager.MemoryInfo(); activityManager.getMemoryInfo(memoryInfo); return memoryInfo; } 复制代码`
* 

优化布局层次、避免自定义组件中 onDraw 创建大量临时对象。