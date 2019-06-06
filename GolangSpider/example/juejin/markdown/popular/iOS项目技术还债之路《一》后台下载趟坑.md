# iOS项目技术还债之路《一》后台下载趟坑 #

## 前言 ##

去年底我在公司开始接手几个迭代了五六年的iOS老项目的技术优化工作。互联网公司的闭源N手业务老代码，经过了若干年和若干波人的轮番洗礼，再加上若干个deadline的赶工加持，已经是千疮百孔，改点东西如履薄冰。往好处想想，前人埋的坑越多，后人才有发挥空间不是。于是我愉快的开始了趟坑之旅。

从提高效率和质量入手，从无数的问题里最先识别出最大的痛点——打包时间过长，开始着手优化编译时间，到梳理团队编码规范及分支管理，再到老模块重构、IAP掉单优化、网络优化等等，一路到刚完成的后台下载优化，有一些想记录和沉淀的东西。

这个系列所涉及的一些技术，大多都不是什么新技术，你可以在网上轻松搜到各种文章。这个系列更关注的是技术怎么帮助业务解决用户的问题，毕竟工作中最让我兴奋的就是自己的方案能落地并真正为用户和公司带来价值。因此这个系列的文章大致都会按照下面的思路来写：

* 业务侧有什么样的 **痛点**
* 当前的技术方案是怎样的
* 技术侧如何分析并识别出问题的根源
* 技术侧如何在给定的 **资源** 下， **选择** 和 **评估** 合适的方案
* 技术侧如何将方案最终 **落地**
* 线上效果如何

作为系列的第一篇，打算先讲下刚刚完成的后台下载优化。

## 目录 ##

* 背景和痛点
* iOS后台机制概述
* 趟坑过程
* 小结

## 一. 背景和痛点 ##

就像所有视频网站都提供移动端视频缓存服务一样，我所在公司的移动端产品也有类似的资源离线缓存服务。缓存服务基本已经是每个提供内容服务App的标配了，有很成熟的技术和各种参考文档。按理来讲照着文档敲一遍代码，这块应该没什么疑问的。但偏偏最近业务侧梳理的用户反馈中，文件下载类反馈成了用户最大的槽点。用户给的反馈普遍比较含糊， ` "下着下着就停了"` 通常是最多的说辞。要根治问题，首先需要挖掘出真正的问题所在，从而对症下药。

回想资源缓存服务也已经上线很久，以前反馈并没这么频繁，应该跟最近的两个改动相关：

* 产品侧增加了批量缓存功能
* 技术侧对资源本身做了优化和技术方案更新

### 产品侧改动 ###

批量缓存使得下载任务的完成时间延长，出错几率更高，确实可能造成用户报障增多。

### 技术侧改动 ###

这里需要简单再介绍下背景，除了资源的缓存服务，我们还提供资源的在线播放服务。所以在这次技术改动之前，资源在服务端一直是存在两份的：

* 一份是原始的资源文件夹，里面包含各种小文件，有 ` png` 、 ` txt` 、 ` json` 、 ` ts` 、 ` wav` 等各种格式的一个富媒体合集，供在线播放使用。播放时直接在线加载小文件地址。
* 另一份是原始资源的压缩包，供离线缓存使用。移动端通过压缩包地址缓存完成后再解压。

这次调整目的是为了节省服务端资源，如果能将冗余的压缩包去掉，仅保留原始资源文件夹的话，可以节省接近50%的空间。不过客户端也需要做一次大重构，离线缓存方式从 ` 压缩包下载` 改为 ` 分片下载` ，即分别下载资源包里面的每个小文件。其实以前的 ` 压缩包下载` 方式对客户端来讲是更友好一些，实现难度更低，下载速度理论上更快。最终全盘考虑和调研下来， ` 分片下载` 和 ` 压缩包下载` 实测下载速度差不多，因此决定客户端先做一次让步，采用 ` 分片下载` 的方式，这里就不多展开了。

结合产品和技术侧的改动，用户在批量下载多个资源时，每个资源里面又有一系列小文件需要下载，复杂度已经远高于单文件下载，这块如果处理的有问题，尤其在后台运行限制那么多的iOS设备上，确实会很影响用户体验。

## 二. iOS后台机制概述 ##

先来捋一下iOS后台机制相关的内容，毕竟是iOS7时代的产物，到现在已经比较生疏了。

这么多年你可能接触过很多后台相关的技术，这些名词或API对你来说一定不陌生：

* ` Background Fetch` , ` Background Task` , ` Background Modes` , ` Background Execution` , ` Background Download`
* ` beginBackgroundTaskWithName:expirationHandler:` , ` endBackgroundTask:`
* ` application:performFetchWithCompletionHandler:`
* ` application:didReceiveRemoteNotification:fetchCompletionHandler:`
* ` application:handleEventsForBackgroundURLSession:completionHandler:`

后台下载( ` Background Download` )是什么，后台运行( ` Bakcground Execution` )又是什么，本文要解决的问题又和哪几块内容有关。带着这些问题，我们先把概念理一理。

照着 [官方文档]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Flibrary%2Farchive%2Fdocumentation%2FiPhone%2FConceptual%2FiPhoneOSProgrammingGuide%2FBackgroundExecution%2FBackgroundExecution.html ) 画了下后台相关的技术全景图，如下所示：

![后台运行](https://user-gold-cdn.xitu.io/2019/6/6/16b286e173e17d6c?imageView2/0/w/1280/h/960/ignore-error/1)

### 后台运行 ( ` Bakcground Execution` ) ###

所有这些概念都隶属于 ` Bakcground Execution` 的范畴，它是App在后台运行任务的统称。

### 后台任务 ( ` Background Task` ) ###

这个概念听着像是所有后台相关任务的统称（要和 ` Background Execution` 区分开），实际它是特指某一类任务：有时App在进入后台后还有任务没执行完，还需要运行一小段时间，那么可以用 ` Background Task` 相关API向系统申请运行权限，运行完了再通知系统可以挂起App了。

### 后台模式 ( ` Background Modes` ) ###

需要后台长时间运行任务的App都需要显式向系统申请权限。这里以 ` Background Audio` 和 ` Background Fetch` 为例。前者允许App在后台播放音频，像QQ音乐这种；后者允许App时不时被唤醒来更新一些数据。

### 后台下载 ( ` Background Download` ) ###

专指由配置了 ` backgroundSessionConfiguration` 的 ` NSURLSession` 管理的下载过程。由系统进程接管App数据的下载，因此即便App被系统挂起，甚至杀死或崩溃了，也能继续下载。下载完成后App会被唤醒，处理一些状态更新和回调。

简单理完相关概念，然后回到我们要解决的问题，先锁定 ` Background Task` 和 ` Background Download` ，初步怀疑导致后台下载效果不好的原因如下：

* App中有其他 ` Background Task` 执行超时导致App过早被杀，后台活跃下载时间过短
* 后台下载没有正确实现
* 我们的下载场景比较特殊，后台下载hold不住

脑子里还有一些别的疑问，比如：

* App进入后台以后到底还能活多久
* App被系统强杀后还能后台下载么
* App被唤醒是一种什么体验
* 保证系统内存充足的话，让App活更久，后台下载是不是就能更持久

带着这么多吃不准的问题，我们直接来看实际效果。

## 三. 趟坑过程 ##

我理想中的后台下载体验是这样的：

* 批量下载一批资源
* 睡一觉
* 醒来发现都下载好了

然鹅现实是骨感的：

* 批量下载一批资源
* 洗个澡
* 发现App被杀，第一个资源都只下了不到一半

后台下载基本处于无效状态，这种体验用户能不吐槽么。

### 第一阶段：确保后台下载可用 ###

后台相关问题调试起来比较麻烦，因为Xcode的调试器会阻止App被系统挂起，没法模拟真正的App后台行为，因此不能用Debug。同时由于模拟器也不一定能准确模拟App行为，最好能在真机上测试。

可以选择打日志，真机连上Mac，用Mac上的Console程序看日志。

为了让日志看起来方便，我在 ` BackgroundDownloader` 里的如下位置加了一些带特征值( ` hello` 开头)的日志：

` @implementation BackgroundDownloader - ( void )addURLs:( NSArray < NSURL *> *)urls { [urls enumerateObjectsUsingBlock:^( NSURL *url, NSUInteger idx, BOOL * _Nonnull stop) { NSURLSessionTask *task = [ self.urlSession downloadTaskWithURL:url]; [task resume]; NSLog ( @"hello 添加下载任务: %@" , url.absoluteString); }]; } - ( void )handleEventsForBackgroundURLSession:( NSString *)aBackgroundURLSessionIdentifier completionHandler:( void (^)())aCompletionHandler { NSLog ( @"hello handleEventsForBackgroundURLSession:completionHandler: 调用" ); // Do other stuff } - ( void )URLSession:( NSURLSession *)session downloadTask:( NSURLSessionDownloadTask *)downloadTask didFinishDownloadingToURL:( NSURL *)location { NSLog ( @"hello 分片下载完成: %@" , downloadTask.originalRequest.URL.absoluteString); } - ( void )URLSessionDidFinishEventsForBackgroundURLSession:( NSURLSession *)session { NSLog ( @"hello URLSessionDidFinishEventsForBackgroundURLSession: 调用" ); // Do other stuff } @end 复制代码`

这样在Console中通过筛选 ` 信息 -> hello` 就能排除各种干扰日志，只保留你想看的：

![console_hello.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e17a6d3dcf?imageView2/0/w/1280/h/960/ignore-error/1)

为了能更好地理解日志里方法的调用顺序，再简单补充下资源的下载逻辑：

> 
> 
> 
> 上文有提到过公司资源走的都是 ` 分片下载` ，即每个资源都是由一组分片组成的，一个分片就是一个待下载的小文件。对于分片采用并发下载，最大并发量是4，同一个资源里的所有分片下载完成会开始下一个资源里分片的下载。
> 
> 
> 

> 
> 
> 
> 假设批量开启了3个资源A、B、C的下载任务，分别包含100、200、300个分片。一开始资源A里的四个分片下载任务同时被启动，每当收到一个下载完成回调，就开启一个新分片的下载。直到A中所有分片都下载完成，再从B中同时开启四个分片的下载，如此循环，直到所有资源都下载完成。从上图的日志里也可以观察到类似的过程。
> 
> 
> 

我们看看App在下载中退到后台会发生什么，为了能同时看到系统日志，Console中筛选 ` 任一 -> BackgroundDemo` ，并且在 ` 18:02:20` 进入后台：

![console_enter_background_wrong.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e17b38538a?imageView2/0/w/1280/h/960/ignore-error/1)

随后App继续下载了将近三分钟，于 ` 18:05:17` 停止了下载：

![console_enter_background_end.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e17b5273a1?imageView2/0/w/1280/h/960/ignore-error/1)

我们看到，这里同时还调用了 ` URLSessionDidFinishEventsForBackgroundURLSession:` ，表示url session里的事件都处理完了。

代码中 ` BackgroundDownloader` 在进入后台后确实开启了 ` Background Task` ，这里和 ` Background Task` 进入后台能执行180s的说法是一致的：

` - ( void )didEnterBackground:( NSNotification *)notification { self.backgroundTask = [[ UIApplication sharedApplication] beginBackgroundTaskWithName: NSStringFromClass ([ self class ]) expirationHandler:^{ [[ UIApplication sharedApplication] endBackgroundTask: self.backgroundTask]; self.backgroundTask = UIBackgroundTaskInvalid ; }]; } 复制代码`
> 
> 
> 
> 
> 题外话，同时试了下在不开启 ` Background Task` 的情况下，退到后台下载任务几秒后就停了
> 
> 

三分钟的后台执行时间对于批量下载是远远不够的，三分钟后发生了什么，我们看下日志，这里摘取了一些关键日志如下：

` 18:05:17 assertiond [BackgroundDemo:12636] Setting up BG permission check timer for 30s 18:05:17 BackgroundDemo hello 下载进度: 0.0751051 18:05:17 BackgroundDemo hello 添加下载任务: https://test.com/class_ocs/slice/923373023660117946/raw/ts/669b14bdd49d7b1704b5c2241c492b1b/a5c175b811175d2005c9cf96 cd 2f830b.ts 18:05:17 BackgroundDemo hello URLSessionDidFinishEventsForBackgroundURLSession: 调用 18:05:43 assertiond [BackgroundDemo:12636] Sending background permission expiration warning! 18:05:48 assertiond [BackgroundDemo:12636] Forcing crash report with description: BackgroundDemo:12636 has active assertions beyond permitted time: <BKProcessAssertion: 0x104273920; "com.apple.nsurlsessiond.handlesession com.mycompany.backgrounddemo.background1" (backgroundDownload:30s); id:…AEB425907B3C> (owner: nsurlsessiond:94) 18:05:48 assertiond [BackgroundDemo:12636] Finished crash reporting. 复制代码`

可以看到：

* ` 18:05:17` 系统发出了后台执行最后通牒，30秒内不结束就要给颜色看了！
* ` 18:05:43` 系统发出最后一次警告
* ` 18:05:48` 系统强制杀死App，并生成了App崩溃报告

我们打开崩溃日志看下：

![crash_18.05.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e17b287c28?imageView2/0/w/1280/h/960/ignore-error/1)

看到了著名的错误码 ` 0x8badf00d` ，"ate bad food"：

> 
> 
> 
> The exception code 0x8badf00d indicates that an application has been
> terminated by iOS because a watchdog timeout occurred. The application
> took too long to launch, terminate, or respond to system events. One
> common cause of this is doing synchronous networking on the main thread.
> Whatever operation is on Thread 0 needs to be moved to a background
> thread, or processed differently, so that it does not block the main
> thread.
> 
> 

[官网]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Flibrary%2Farchive%2Ftechnotes%2Ftn2151%2F_index.html ) 提到是由于主线程卡住时间太长，系统的watchdog将App强杀了。MrPeak的 [这篇]( https://link.juejin.im?target=http%3A%2F%2Fmrpeak.cn%2Fblog%2Fios-background-task%2F ) 和官网论坛里的 [这篇]( https://link.juejin.im?target=https%3A%2F%2Fforums.developer.apple.com%2Fthread%2F88529 ) 也提到了另外的可能性，即App中出现了 ` leaked Background Task` ，后台任务没有被正确end。究竟是哪种情况，可以进一步分析下崩溃时主线程的Stack：

` Thread 0 Crashed: 0 libsystem_kernel.dylib 0x0000000196505c60 mach_msg_trap + 8 1 CoreFoundation 0x000000019690de10 __CFRunLoopServiceMachPort + 240 2 CoreFoundation 0x0000000196908ab4 __CFRunLoopRun + 1344 3 CoreFoundation 0x0000000196908254 CFRunLoopRunSpecific + 452 4 GraphicsServices 0x0000000198b47d8c GSEventRunModal + 108 5 UIKitCore 0x00000001c3c504c0 UIApplicationMain + 216 6 BackgroundDemo 0x00000001047025c4 0x104688000 + 501188 7 libdyld.dylib 0x00000001963c4fd8 start + 4 复制代码`

和MrPeak文中提到的Stack非常类似：

> 
> 
> 
> 这个 stack 很经典，经常会看到，不需要 symbolicate 也能知道是干啥，这是 UI 线程 runloop 处于 idle 状态的
> stack，在等待 kernel 的 message。表示 UI 线程此时处于闲置状态，这种状态下的系统强杀大概率是由于 leaked
> Background Task 导致的。
> 
> 

基本可以断定 ` 18:05:48` 的这次崩溃是由 ` leaked Background Task` 导致。

所以罪魁祸首是后台崩溃么？并不见得，因为后台下载在App崩溃、被挂起或杀死的情况下仍然有效。只有当用户手动杀死App，后台下载才会失效。可以参考 [官方回复]( https://link.juejin.im?target=https%3A%2F%2Fforums.developer.apple.com%2Fthread%2F14855 ) ：

> 
> 
> 
> The behaviour of background sessions after a force quit has changed over
> time:
> 
> 

> 
> 
> 
> * In iOS 7 the background session and all its tasks just ‘disappear’.
> * In iOS 8 and later the background session persists but all of the tasks
> are cancelled.
> 
> 
> 

> 
> 
> 
> When you terminate an app via the multitasking UI (force quit), the system
> interprets that as a strong indication that the app should do no more work
> in the background, and that includes NSURLSession background tasks.
> 
> 

` objc.io` 上的 [这篇]( https://link.juejin.im?target=https%3A%2F%2Fwww.objc.io%2Fissues%2F5-ios7%2Fmultitasking%2F ) 也提到了：

> 
> 
> 
> Tasks added to a background session are run in an external process and
> continue even if your app is suspended, crashes, or is killed.
> 
> 

(经验证，App崩溃或者在后台被强杀以后确实还能继续下载，不过有一些注意点和坑，可以参考 [iOS原生级别后台下载详解]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fda565e14ef88 ) ，写的非常详细。)

那么我们再次从接下来的日志里找找后台下载的痕迹。继续截取一些关键日志：

` 18:05:48 symptomsd Entry, display name com.mycompany.backgrounddemo uuid (null) pid 12636 isFront 0 18:11:28 nsurlsessiond [C3794 Hostname #aca88929:443 tcp, bundle id: com.mycompany.backgrounddemo, url hash: 8f857432, traffic class: 100, tls, indefinite] start 18:11:30 nsurlsessiond [C3794 Hostname #aca88929:443 tcp, bundle id: com.mycompany.backgrounddemo, url hash: 8f857432, traffic class: 100, tls, indefinite] cancel 18:11:30 nsurlsessiond [C3794 Hostname #aca88929:443 tcp, bundle id: com.mycompany.backgrounddemo, url hash: 8f857432, traffic class: 100, tls, indefinite] cancelled 复制代码`

` 18:05:48` 是之前后台崩溃完了以后的最后一条，随后有大概5分30秒的时间内没有一条日志，直到 ` 18:11:28` 一个名为 ` nsurlsessiond` 的进程打印了几条日志。 ` nsurlsessiond` 实际上就是接管了App后台下载的系统 ` daemon` 进程。奇怪的是，它 ` start` 后立马就 ` cancel` 了，非常诡异的行为。暂时不管，接着往下看日志：

` 18:14:29 nsurlsessiond [C3795 Hostname #aca88929:443 tcp, bundle id: com.mycompany.backgrounddemo, url hash: fae4ed2d, traffic class: 100, tls, indefinite] start 18:14:31 nsurlsessiond [FBSSystemService][0xc652] Sending request to open "com.mycompany.backgrounddemo" 18:14:31 SpringBoard [FBSystemService][0xc652] Received request to open "com.mycompany.backgrounddemo" from nsurlsessiond:94. 18:14:31 SpringBoard Received trusted open application request for "com.mycompany.backgrounddemo" from <FBProcess: 0x28396e370; nsurlsessiond; pid: 94>. 18:14:31 SpringBoard Executing request: <SBMainWorkspaceTransitionRequest: 0x28245b380; eventLabel: OpenApplication(com.mycompany.backgrounddemo)ForRequester(nsurlsessiond.94); display: Main; source : FBSystemService> 18:14:31 SpringBoard Executing suspended-activation immediately: OpenApplication(com.mycompany.backgrounddemo)ForRequester(nsurlsessiond.94) 18:14:31 SpringBoard Bootstrapping com.mycompany.backgrounddemo with intent background 18:14:31 SpringBoard Application process state changed for com.mycompany.backgrounddemo: <SBApplicationProcessState: 0x280ba7ae0; pid: 12648; taskState: Running; visibility: Unknown> 复制代码`

又过了三分钟，到 ` 18:14:29` 后 ` nsurlsessiond` 又一次 ` start` ，这次它挺住了，向系统申请打开我们的App。随后 ` SpringBoard` 同意并帮它从后台成功启动了App，将App状态置为 ` running` 。接下来：

` 18:14:31 assertiond [BackgroundDemo:12648] Add assertion: <BKProcessAssertion: 0x1046029b0; id: 94-679EF71C-A8E3-42E4-B679-DE0D01BDBB0F; name: "com.apple.nsurlsessiond.handlesession com.mycompany.backgrounddemo.background1" ; state: active; reason: backgroundDownload; duration: 30.0s> 复制代码`

系统宣布， ` nsurlsessiond` 中任务都处理完了，即将调用App中的 ` application:handleEventsForBackgroundURLSession:completionHandler:` 方法了，并且传入 ` sessionIdentifier` 为 ` com.mycompany.backgrounddemo.background1` ，并且给了App 最多30秒的后台执行时间。接下来：

` 18:14:31 assertiond [BackgroundDemo:12648] Mutating assertion reason from finishTask to finishTaskAfterBackgroundDownload 18:14:31 assertiond [BackgroundDemo:12648] Add assertion: <BKProcessAssertion: 0x1042246a0; id: 12648-B0511F94-D84E-479C-AD8F-CA13B7CA55F6; name: "Shared Background Assertion 1 for com.mycompany.backgrounddemo" ; state: active; reason: finishTaskAfterBackgroundDownload; duration: 40.0s> 复制代码`

紧接着系统改主意了，给了App最多40秒的后台执行时间。具体原因不明。接下来系统果不其然，在大约40秒后又一次警告并杀死了App：

` 18:15:07 assertiond [BackgroundDemo:12648] Sending background permission expiration warning! 18:15:12 assertiond [BackgroundDemo:12648] Forcing crash report with description: BackgroundDemo:12648 has active assertions beyond permitted time 复制代码`

崩溃错误码同样是 ` 0x8badf00d` 。再然后，App就再不产生任何日志了，可以认为这一波的App后台行为全部结束了。

打开App后发现下载进度还停留在App后台满三分钟被杀死那一刻的进度，说明由 ` nsurlsessiond` 接管的后台下载根本没有生效。

hmm，一定是哪里有问题，经过review代码，发现一个低级错误，可能是不同团队间沟通问题导致：在 ` AppDelegate` 中虽然实现了 ` application:handleEventsForBackgroundURLSession:completionHandler:` ，但是没有正确把参数传入 ` BackgroundDownloader` ，导致后台下载实现不完整。上面日志中一些奇怪的点说不定就是这个引起的。OK，修复完bug再跑一遍，看看还有没有别的坑。

为了能更直观地看清整个过程，我手撸了下面这张图，可以看到App进入后台后，随着时间推移，都发生了些什么事。横轴是时间，单位为秒：

![sequence_4.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e17b8699f1?imageView2/0/w/1280/h/960/ignore-error/1)

大致总结一下就是：进入后台 ` ->` 系统进程干活 ` ->` 系统唤醒App ` ->` App添加新任务 ` ->` App通知系统处理完成并挂起 ` ->` 系统进程干活 ` ->` 系统唤醒App ` ->` App添加新任务 ` ->` App通知系统处理完成并挂起 ` ->`....... 如此循环往复直到系统觉得累了并不再唤醒App（1290秒以后就没有任何动静了，不再唤醒的原因会在后面解释）。

这里简单提一下，当系统在后台下载这一批添加到四个分片的时候，App已经被 ` suspend` 了，甚至已经挂了，因此这时候是不会执行App里任何回调的。只有当系统完成了这一批添加的所有四个分片任务时，才会唤醒App，然后App趁这个机会再添加新一批的四个分片任务。

Anyway，后台下载算是生效了。

### 第二阶段：提高后台下载效率 ###

让我们看看能不能继续优化。

从上面的时序图里我们可以看到：

* 每次系统起来干活就只能完成4个分片的下载。原因是每次App被唤醒后就添加了4个分片下载任务，然后被挂起。随后系统抽空开始新一轮的干活。
* 每次系统起来干活时间跨度很长，大概都需要三五分钟才能干完。原因是后台下载什么时候开始是不固定的，是系统根据运行环境和可用资源动态调配的，系统可能隔好几分钟才开始分片的下载，也可能下一会停一会，或者在蜂窝网络下停止下载。
* 系统在经历五六次循环以后，就不再工作了，导致后台下载仅仅完成了20+个分片的下载，连一个完整资源都不能下完，更别提批量下多个资源了。原因在 [官方文档]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdocumentation%2Ffoundation%2Furl_loading_system%2Fdownloading_files_in_the_background ) 里有解释：

> 
> 
> 
> When the system resumes or relaunches your app, it uses a rate limiter to
> prevent abuse of background downloads. When your app starts a new download
> task while in the background, the task doesn't begin until the delay
> expires. The delay increases each time the system resumes or relaunches
> your app. As a result, if your app starts a single background download,
> gets resumed when the download completes, and then starts a new download,
> it will greatly increase the delay.
> 
> 

简言之，后台下载会随着循环次数增加而推迟，直到最后系统罢工。解决方案 [官方文档]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdocumentation%2Ffoundation%2Furl_loading_system%2Fdownloading_files_in_the_background ) 里也提到了：

> 
> 
> 
> Instead, use a small number of background sessions — ideally just one —
> and use these sessions to start many download tasks at once. This allows
> the system to perform multiple downloads at once, and resume your app when
> they have completed.
> 
> 

也就是说用一个 ` NSURLSession` 开启尽可能多的任务数，说不定可以有所改善。不过文档后面也补了一句：

> 
> 
> 
> Keep in mind, though, that each task has its own overhead. If you find you
> need to launch thousands of download tasks, change your design to perform
> fewer, larger transfers.
> 
> 

每一个小分片任务的开启都是需要代价的，一个资源里有几百个小分片，同时都开启官方并不推荐。官方也更支持 ` 压缩包下载` 的方式。之前也提过， ` 分片下载` 是客户端和服务端权衡下来的结果，米已成粥，我们只能想办法优化。

先试试看一个资源里所有的分片同时开启下载会怎么样。之前控制并发的方式是一个时间段内只允许最多 ` resume` 4个 ` task` ，每下好一个分片再 ` resume` 一个新的。现在改为一次性把所有的 ` task` 都 ` resume` 。同时控制并发数为4：

` sessionConfiguration.HTTPMaximumConnectionsPerHost = 4 ; 复制代码`

这样一来，可以避免并发量过大，也可以保证所有的分片任务都在 ` NSURLSession` 的队列中。又跑了一遍，整个过程如下：

![sequence_max.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1a83dc789?imageView2/0/w/1280/h/960/ignore-error/1)

和之前的过程类似，区别在于后台下载每个循环从完成4个分片增加到了完成资源中所有分片的下载。最终总共完成了4个完整资源的下载，比之前提高了不少。

虽然不是官方推荐的方式，但毕竟离目标又近了一步。

顺带提下，可以通过看Console中是否有类似日志来判断后台下载是否正在进行：

![console_nsurlsessiond_working.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1a01f6cbc?imageView2/0/w/1280/h/960/ignore-error/1)

### 第三阶段：业界方案调研 ###

我们当然不满足于此，毕竟批量下载30个资源，才完成4个，是无论如何都说不过去的。并且整个时间跨度比较长，受限于系统的后台策略，没法马不停蹄地进行下载。

从之前的测试我们可以得出结论：要尽量一次性把任务都添加到 ` NSURLSession` 里，这样后台下载才能持久。

我们大胆的想一下，如果我们在批量添加资源的时候，把所有资源的所有分片一次性给 ` NSURLSession` ，那理论上应该是可以全部都下完的。

不过这么做有两个问题：

* 更进一步违背了官方所倡导的最佳实践，有可能带来性能问题和一些意想不到的其他问题
* 现有代码改造成本很大

感觉没了方向的时候，看看业界有没有类似的问题，以及他们是怎么做的。

同事之前研究过腾讯视频mac版，发现下载的都是ts小文件，因此推测腾讯视频iOS端应该也是采用类似 ` 分片下载` 的方式。

我试了试腾讯视频和爱奇艺的后台下载，发现效果都很好，批量添加的任务都可以顺利全部下完。对于批量下载大文件来说不稀奇，但腾讯视频对于 ` 分片下载` 也能有这么好的效果，我决定研究一番。

手机连上mac，用Console观察运行日志，发现：

* 无论腾讯视频在前台还是后台，Console日志都差不多
* 无论在后台放多久，Console日志几乎和之前没区别
* 找不到一些后台下载的特征log，比如本文前面列出的一些log

难道苹果跟鹅厂关系比较好所以... 打住，这么想就太low了。（看了下爱奇艺的运行日志，也类似）

种种迹象似乎都表明，腾讯视频并没有真正被系统挂起，而是在后台仍然处于active状态。那么，它是怎么做到的。

有没有一些种类的App是可以在后台保活的。有，类似导航App和音乐类App。

如果一个App拥有 ` Background Audio` 权限，在后台播放音乐，系统肯定没理由挂起它。

那么如果这段音乐是一段没有声音的空音频，系统应该也没有办法知道。

那么腾讯视频和爱奇艺是不是通过类似手段来实现后台下载的呢，我们来探一探。

我们先用PP助手获取到IPA，解开看看。

#### 腾讯视频 ####

我们发现资源包里有一个与众不同的音频文件 ` sound_drag_refresh.wav` ：

![tengxun_ipa.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1b287d0b7?imageView2/0/w/1280/h/960/ignore-error/1)

用QuickTimePlayer打开，选择 ` 编辑` -> ` 修剪` ，可以看到其波形图：

![tengxun_sound_compare.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1cfb928c5?imageView2/0/w/1280/h/960/ignore-error/1)

上面是正常音频的波形图，下面是 ` sound_drag_refresh.wav` 的，可以看到这应该是一段空音频。为什么要放一个空音频在bundle里，并且起一个看似不相关的名字，令人浮想联翩。

再用Hopper打开，搜索 ` backgroundaudio` ，可以找到一个名为 ` startInfiniteBackgroundAudioTask:` 的方法：

![hopper_tengxun_backgroundaudio.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1b32e1799?imageView2/0/w/1280/h/960/ignore-error/1)

这个方法名字比较可疑，不过看实现没有直接证据表明和刚刚的空音频有关系。

继续搜索 ` sound_drag_refresh` ，可以看到有个叫 ` restartPlayer` 的方法引用了它：

![hopper_tengxun_sound.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1ae22e3d5?imageView2/0/w/1280/h/960/ignore-error/1)

看看 ` restartPlayer` 方法的实现：

![hopper_tengxun_restartplayer.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1d230c5f6?imageView2/0/w/1280/h/960/ignore-error/1)

基本和网上这篇讲 [iOS保活机制]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fb2abd7afdc21 ) 的文章里的代码有些类似之处：

` //静音文件 NSString *filePath = [[ NSBundle mainBundle] pathForResource: @"音频文件+文件名" ofType: @"mp3" ]; NSURL *fileURL = [[ NSURL alloc] initFileURLWithPath:filePath]; self.playerBack = [[ AVAudioPlayer alloc] initWithContentsOfURL:fileURL error: nil ]; [ self.playerBack prepareToPlay]; // 0.0~1.0,默认为1.0 self.playerBack.volume = 0.01 ; // 循环播放 self.playerBack.numberOfLoops = -1 ; 复制代码`

注意这段：

` mov.w r2, #0xffffffff movt r0, #0x1e9 ; @selector(setNumberOfLoops:), :upper16:(0x21f2f88 - 0x361cee) add r0, pc ; @selector( set NumberOfLoops:) ldr r1, [r0] ; @selector( set NumberOfLoops:), "setNumberOfLoops:" 复制代码`

` 0xffffffff` 应该就是 ` -1` ，放在 ` r2` 寄存器中，作为参数传递给 ` [setNumberOfLoops:]` ，表示循环播放。

题外话： ` armv7` 体系下，函数前四参数用 ` r0` 到 ` r3` 来传递， ` r0` 放 ` self` ， ` r1` 放 ` _cmd` 。

基本上破案了。至于内部是怎么调用的， ` startInfiniteBackgroundAudioTask:` 和 ` restartPlayer` 怎么串联起来的，调研起来太费精力，且不是本文重点，这里就忽略了。有兴趣的读者可以研究下。

#### 爱奇艺 ####

用同样的方法看下爱奇艺。

我们也发现资源包里有一个空音频文件 ` JM.wav` （后续同事猜测为 ` 静默` 的拼音 2333）：

![iqiyi_sound.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1e0ec416f?imageView2/0/w/1280/h/960/ignore-error/1)

接着用Hopper打开，搜索 ` jm` ，可以找到下面几个方法：

![hopper_iqiyi_search_jm.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1e0fbe20e?imageView2/0/w/1280/h/960/ignore-error/1)

我们看下 ` [QYOfflineBaseModelUtil isOpenJMAudio]` 的实现：

![hopper_iqiyi_isopenjm.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1e3ed7b5b?imageView2/0/w/1280/h/960/ignore-error/1)

大概意思用伪代码描述下：

` + ( BOOL )isOpenJMAudio { if ([ self isUseURLSession]) { // 如果用URLSession，则不要开启后台保活（这里可能是做了开关，可以在原生后台下载方式和播放静默音频保活间切换） return NO ; } else { // 否则，如果有任务在下载，则开启后台保活，不然不开启 return [[QYDownloadTaskManager sharedInstance] isAnyTaskCanBeDownloaded]; } } 复制代码`

再看下 ` [AppDelegate turnOnJM]` 的实现：

![hopper_iqiyi_turnonjm.png](https://user-gold-cdn.xitu.io/2019/6/6/16b286e1e5559d4f?imageView2/0/w/1280/h/960/ignore-error/1)

又看到了熟悉的 ` [setNumberOfLoops:]` 和 ` 0xffffffff` ，以及刚刚的 ` JM.wav` 。

好，又破案了。至此，业界做法基本搞清楚了：简单粗暴，通过后台保活机制，使得App在后台的行为像在前台一样。

幸运的是我们的App原来已经有了 ` Background Audio` 的权限，可以依葫芦画瓢增加后台保活机制，同时把 ` NSURLSession` 相关后台实现去掉。

实测下来耗电量并没有增加多少，批量下载也和腾讯视频一样顺畅了，都可以顺利全部下完。

上线一个月以来，收到的相关用户报障基本没了。至此，趟坑之旅告一段落。

## 四. 小结 ##

简单对比下 ` NSURLSession` 和后台保活两种机制：

+----------------+------------------------------------------+--------------------------+
|      指标      |               NSURLSESSION               |         后台保活         |
+----------------+------------------------------------------+--------------------------+
| 耗电量         | 少，系统会做优化                         | 多，但实测下来增加有限   |
| 速度           | 慢                                       | 快                       |
| 大文件批量下载 | 可以全部下完                             | 可以全部下完             |
| 分片批量下载   | 实现成本高，官方不推荐，不一定能全部下完 | 实现成本低，可以全部下完 |
| 遇到崩溃       | 可以继续下载                             | 下载过程停止             |
| 权限           | 无须申请额外权限                         | 需要申请 `               |
|                |                                          | Background Audio`        |
|                |                                          | 权限                     |
+----------------+------------------------------------------+--------------------------+

综上，可以根据自己公司的业务诉求，可以采用不同的策略实现iOS后台下载，或者尝试下两者的结合，应对崩溃的情况。（之前测了下两种机制都开启好像会有问题，有兴趣的可以调研下）

这个系列下一篇准备讲一讲 ` IAP` 掉单的优化。

完。

### 参考链接： ###

* [Downloading Files in the Background]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdocumentation%2Ffoundation%2Furl_loading_system%2Fdownloading_files_in_the_background )
* [Background Execution]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Flibrary%2Farchive%2Fdocumentation%2FiPhone%2FConceptual%2FiPhoneOSProgrammingGuide%2FBackgroundExecution%2FBackgroundExecution.html )
* [Preparing Your App to Run in the Background]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdocumentation%2Fuikit%2Fcore_app%2Fmanaging_your_app_s_life_cycle%2Fpreparing_your_app_to_run_in_the_background )
* [Testing Background Session Code]( https://link.juejin.im?target=https%3A%2F%2Fforums.developer.apple.com%2Fthread%2F14855 )
* [iOS11 watchdog timeout crashes (0x8badf00d) but code not on stack]( https://link.juejin.im?target=https%3A%2F%2Fforums.developer.apple.com%2Fthread%2F88529 )
* [0x8badf00d - "ate bad food" ]( https://link.juejin.im?target=https%3A%2F%2Fgeek-is-stupid.github.io%2F2018-10-15-0x8badf00d-ate-bad-food%2F )
* [Understanding and Analyzing Application Crash Reports]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Flibrary%2Farchive%2Ftechnotes%2Ftn2151%2F_index.html )
* [iOS Background Download with silent notification iOS 11]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40pratheeshdhayadev%2Fios-background-download-with-silent-notification-ios-11-77af4b1ab09d )
* [Multitasking in iOS 7]( https://link.juejin.im?target=https%3A%2F%2Fwww.objc.io%2Fissues%2F5-ios7%2Fmultitasking%2F )
* [iOS 后台下载及管理库]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F2ccb34c460fd )
* [iOS原生级别后台下载详解]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fda565e14ef88 )
* [iOS程序后台保活]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fb2abd7afdc21 )
* [Working with NSURLSession: Part 4]( https://link.juejin.im?target=https%3A%2F%2Fcode.tutsplus.com%2Ftutorials%2Fworking-with-nsurlsession-part-4--mobile-22545 )
* [Download multiple files using iOS background transfer service]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F22842933%2Fdownload-multiple-files-using-ios-background-transfer-service )
* [Background Transfer Service in iOS 7 SDK: How To Download File in Background
]( https://link.juejin.im?target=https%3A%2F%2Fwww.appcoda.com%2Fbackground-transfer-service-ios7%2F )
* [iOS Swift download lots of small files in background]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F40403953%2Fios-swift-download-lots-of-small-files-in-background )
* [iOS 7 NSURLSession Download multiple files in Background]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F25762034%2Fios-7-nsurlsession-download-multiple-files-in-background )
* [Downloading a list of 100 files one by one using NSURLSession in Background]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F23356118%2Fdownloading-a-list-of-100-files-one-by-one-using-nsurlsession-in-background )