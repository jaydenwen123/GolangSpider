# Booster 系列之——性能瓶颈检测 #

> 
> 
> 
> 项目地址： [github.com/didi/booste…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster )
> 
> 

对于一款 APP 来说，卡顿率、ANR 率是衡量这个 APP 质量的两个重要指标，目前已经有很多成熟的 APM 工具和平台来统计 APP 的运行时性能，但是对于实行敏捷开发的产品来说，从 APP 开发，到灰度发布，再到全量，要经历一个漫长的过程，等到收集到上报的卡顿和 ANR，再去修复，又要经历灰度、全量这一漫长的过程。

如果能在上线之前就能发现代码中的性能问题并进行修复，将大大的加速了产品迭代的效率，一般来说，实现的方式可能有以下几种：

* 代码审查
* 代码扫描
* 静态分析

而 Booster 选择了静态分析，之所以采用静态分析的方案，原因是因为前两种方案都无法解决无代码访问权限的情况 [[1]]( #fn1 )

## 性能测量 ##

Android 官方提供了很多 Profiling 工具，尽管这些工具非常强大，但是对于开发者来说，都需要太多的人工介入，而且门槛比较高，如：

* 

[Method Tracing]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fstudio%2Fprofile%2Fgenerate-trace-logs )

启用 Method Tracing 需要在想要测量的代码段中加上这两行代码：

` Debug.startMethodTracing( "booster" ) ... Debug.stopMethodTracing() 复制代码`

而且， Method Tracing 严重损耗运行时性能，如果测量的范围过大，使用起来卡到不能忍受。

* 

[systrace]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fstudio%2Fcommand-line%2Fsystrace )

启用 systrace 需要启动 adb 连上设备进入 debug 模式，并在代码段中加上这两行代码：

` Trace.beginSection( "Activity.onCreate()" ) ... Trace.endSection() 复制代码`

虽然性能开销比 Method Tracing 少了许多，但是测量的范围受 buffer 的限制，只能测量一段代码的性能。

* 

[Android Profiler]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fstudio%2Fprofile%2Fandroid-profiler )

Android Studio 3.0 虽然提供了强大的 Android Profiler 来帮助开发者定位分析问题，但是只有 debug 覆盖到的代码分支才能被检测到，而且范围有限。

* 

[Beanchmark]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fstudio%2Fprofile%2Fbenchmark )

Android 提供的 Jetpack Benchmark Library 可以通过写单元测试来测量代码的性能，对于快速迭代的产品来说，无疑是个摆设。

为了能够在上线之前快速的发现所有代码中潜在的性能问题，我们提出了通过静态分析来检测代码中存在的性能瓶颈。

## 如何确定性能瓶颈？ ##

### 主线程 ###

对 APP 来说，ANR 和卡顿问题的根源在于主线程被阻塞，因此，对于基于 event-loop 的系统来说，任何阻塞主线程的方法调用 [[2]]( #fn2 ) 都可以认为是性能瓶颈。除此之外，还有其它影响运行时性能和稳定性的因素，比如：线程过载 [[3]]( #fn3 ) 、使用 ` finalizer` [[4]]( #fn4 ) 等等。

基于静态分析的性能瓶颈检测的关键在于确定方法运行的线程是否是主线程。几乎所有基于 event-loop 的GUI 系统，操作 UI 都是在主线程/UI 线程中进行，这就意味着：

* 只要能找到跟 UI 相关的方法调用，就可以认为该方法是在主线程中运行；
* 只要一条调用链路中的任意一个方法在主线程中调用，就可以认为该链路是在主线程中运行。

### 主线程入口 ###

经过分析，最终我们确定了如下规则：

* 以 ` Application` 的模板方法为起点的调用链路，详见： [Application Entry Points]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster%2Fblob%2Fmaster%2Fbooster-transform-lint%2Fsrc%2Fmain%2Fkotlin%2Fcom%2Fdidiglobal%2Fbooster%2Ftransform%2Flint%2Fconstants.kt%23L27 )
* 以 ` Activity` 的模板方法为起点的调用链路，详见： [Activity Entry Points]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster%2Fblob%2Fmaster%2Fbooster-transform-lint%2Fsrc%2Fmain%2Fkotlin%2Fcom%2Fdidiglobal%2Fbooster%2Ftransform%2Flint%2Fconstants.kt%23L38 )
* 以 ` Service` 的模板方法为起点的调用链路，说见： [Service Entry Points]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster%2Fblob%2Fmaster%2Fbooster-transform-lint%2Fsrc%2Fmain%2Fkotlin%2Fcom%2Fdidiglobal%2Fbooster%2Ftransform%2Flint%2Fconstants.kt%23L137 )
* 以 ` BroadcastReceiver` 的模板方法为起点的调用链路，详见： [Receiver Entry Points]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster%2Fblob%2Fmaster%2Fbooster-transform-lint%2Fsrc%2Fmain%2Fkotlin%2Fcom%2Fdidiglobal%2Fbooster%2Ftransform%2Flint%2Fconstants.kt%23L153 )
* 以 ` ContentProvider` 的模板方法为起点的调用链路，详见： [Provider Entry Points]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster%2Fblob%2Fmaster%2Fbooster-transform-lint%2Fsrc%2Fmain%2Fkotlin%2Fcom%2Fdidiglobal%2Fbooster%2Ftransform%2Flint%2Fconstants.kt%23L153 )
* 以参数列表及返回值中包含下列类型的方法为起点的调用链路

* ` Fragment`
* ` Dialog`
* ` View`
* ` Widget`
* ` Layout`

* 通过 Main Handler 提交的 Runnable 的 run() 方法

以上规则虽然不能命中所有的主线程入口，但至少解决了 80% 的问题，而且，每个 APP 的架构不一样，如果要做到更加精准，需要针对地性的对 Booster 进行扩展了。

### 方法调用链路 ###

经过前面的分析，我们能够从整个 Call Graph 中分离出所有在主线程中的调用链路了，但是，如何确定哪些调用链路是存在性能瓶颈的呢？

在经过大量的统计分析之后，我们确定了会阻塞主线程的方法列表，由于篇幅原因，以下只列举了一部分 API，详细列表请参见： [LINT_APIS]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Fbooster%2Fblob%2Fmaster%2Fbooster-transform-lint%2Fsrc%2Fmain%2Fkotlin%2Fcom%2Fdidiglobal%2Fbooster%2Ftransform%2Flint%2Fconstants.kt%23L170 ) ：

` "java/lang/Object.wait()V" , "java/lang/Object.wait(J)V" , "java/lang/Object.wait(JI)V" , "java/lang/Thread.start()V" , "java/lang/ClassLoader.getResource(Ljava/lang/String;)Ljava/net/URL;" , "java/lang/ClassLoader.getResources(Ljava/lang/String;)Ljava/util/Enumeration;" , "java/lang/ClassLoader.getResourceAsStream(Ljava/lang/String;)Ljava/io/InputStream;" , "java/lang/ClassLoader.getSystemResource(Ljava/lang/String;)Ljava/net/URL;" , "java/lang/ClassLoader.getSystemResources(Ljava/lang/String;)Ljava/util/Enumeration;" , "java/lang/ClassLoader.getSystemResourceAsStream(Ljava/lang/String;)Ljava/io/InputStream;" , ... "java/util/zip/ZipFile.<init>(Ljava/lang/String;)" , "java/util/zip/ZipFile.getInputStream(Ljava/util/zip/ZipEntry;)" , "java/util/jar/JarFile.<init>(Ljava/lang/String;)" , "java/util/jar/JarFile.getInputStream(Ljava/util/jar/JarEntry;)" , ... "android/content/Context.getSharedPreferences(Ljava/lang/String;I)Landroid/content/SharedPreferences;" , "android/content/SharedPreferences$Editor.apply()V" , "android/content/SharedPreferences$Editor.commit()B" , ... "android/content/res/AssetManager.list(Ljava/lang/String;)[Ljava/lang/String;" , "android/content/res/AssetManager.open(Ljava/lang/String;)Ljava/io/InputStream;" , "android/content/res/AssetManager.open(Ljava/lang/String;I)Ljava/io/InputStream;" , "android/content/res/AssetManager.openFd(Ljava/lang/String;)Landroid/content/res/AssetFileDescriptor;" , "android/content/res/AssetManager.openNonAssetFd(Ljava/lang/String;)Landroid/content/res/AssetFileDescriptor;" , "android/content/res/AssetManager.openNonAssetFd(ILjava/lang/String;)Landroid/content/res/AssetFileDescriptor;" , "android/content/res/AssetManager.openXmlResourceParser(Ljava/lang/String;)Landroid/content/res/XmlResourceParser;" , "android/content/res/AssetManager.openXmlResourceParser(ILjava/lang/String;)Landroid/content/res/XmlResourceParser;" , ... "android/graphics/BitmapFactory.decodeByteArray([BIILandroid/graphics/BitmapFactory$Options;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeByteArray([BII)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeFile(Ljava/lang/String;Landroid/graphics/BitmapFactory$Options;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeFile(Ljava/lang/String;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeFileDescriptor(Ljava/io/FileDescriptor;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeFileDescriptor(Ljava/io/FileDescriptor;Landroid/graphics/Rect;Landroid/graphics/BitmapFactory$Options;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeResource(Landroid/content/res/Resources;I)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeResource(Landroid/content/res/Resources;ILandroid/graphics/BitmapFactory$Options;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeResourceStream(Landroid/content/res/Resources;Landroid/util/TypedValue;Ljava/io/InputStream;Landroid/graphics/Rect;Landroid/graphics/BitmapFactory$Options;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeStream(Ljava/io/InputStream;)Landroid/graphics/Bitmap;" , "android/graphics/BitmapFactory.decodeStream(Ljava/io/InputStream;Landroid/graphics/Rect;Landroid/graphics/BitmapFactory$Options;)Landroid/graphics/Bitmap;" 复制代码`

根据前面得出的结论，我们就可以通过在所有在主线程调用的链路中去匹配上面定义的 API 列表来找出有性能瓶颈的链路了。

## 总结 ##

对于性能瓶颈检测来说，其首要任务是构建 Call Graph [[5]]( #fn5 ) ， Lint Transformer 按如下步骤进行：

* 

解析 AndroidManifest.xml ，得到 ` Application` 以及四大组件的类名；

* 

创建 Globa Call Graph [[6]]( #fn6 ) 和 Lint Call Graph [[7]]( #fn7 ) ，以 ROOT 节点作为所有主线程入口方法的父节点，便于后续分离出主线程的调用链路， Global Call Graph 的结构如下图所示；

![Call Graph](https://user-gold-cdn.xitu.io/2019/6/4/16b20a24bde01eb6?imageView2/0/w/1280/h/960/ignore-error/1)

* 

解析所有的 class 文件，从方法体指令序列中提取 invoke 指令 [[8]]( #fn8 ) ，构建 Edge [[5:1]]( #fn5 ) ，并加入到 Call Graph 中；

* 

以 ROOT 节点的一级子节点为根，开始遍历整个 Call Graph 来匹配前面确定的方法列表，如果匹配成功，则将该链路加到 Lint Call Graph 中

* 

最后将 Lint Call Graph 以入口类单位分成更小的 Call Graph ，生成 [dot]( https://link.juejin.im?target=https%3A%2F%2Fwww.graphviz.org%2Fdoc%2Finfo%2Flang.html ) 格式的报告，转换为 PNG 格式后，如下图所示：

![Main Activity](https://user-gold-cdn.xitu.io/2019/6/4/16b20a24bf17ccd1?imageView2/0/w/1280/h/960/ignore-error/1)

* 

如第三方 SDK，只提供编译好的二进制 [↩︎]( #fnref1 )

* 

如 I/O 操作、网络访问、SQLite 访问、 ` Thread.sleep(...)` 、 ` Object.wait(...)` 等 [↩︎]( #fnref2 )

* 

过多的线程会导致 OOM [↩︎]( #fnref3 )

* 

Android 官方明确声明 ` finalize` 方法的调用时机不可靠 [↩︎]( #fnref4 )

* 

![callgraph = [edge(f, g), …]](https://juejin.im/equation?tex=callgraph%20%3D%20%5Bedge(f%2C%20g)%2C%20%E2%80%A6%5D) ；其中 ![f](https://juejin.im/equation?tex=f) 为 caller , ![g](https://juejin.im/equation?tex=g) 为 callee. [en.wikipedia.org/wiki/Call_g…]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FCall_graph ) [↩︎]( #fnref5 ) [↩︎]( #fnref5:1 )

* 

Global Call Graph 表示整个 APP 的调用关系图 [↩︎]( #fnref6 )

* 

Lint Call Graph 表示存在性能瓶颈的所有链路组成的图 [↩︎]( #fnref7 )

* 

[docs.oracle.com/javase/spec…]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Fspecs%2Fjvms%2Fse7%2Fhtml%2Fjvms-6.html ) [↩︎]( #fnref8 )