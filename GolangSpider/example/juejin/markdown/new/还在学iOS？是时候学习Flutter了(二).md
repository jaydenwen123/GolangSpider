# 还在学iOS？是时候学习Flutter了(二) #

### 概述 ###

本文承接上文，是Flutter For iOS 的第二篇文章，通过阅读本文你将获取如下信息：

* 线程和异步
* 项目结构与本地化
* 视图控制器
* 布局
* 手势
* 表单
* 列表
* 其他

### 线程和异步 ###

#### 如何写异步代码 ####

Dart拥有单线程执行模型，同时也支 ` Isolate` (一种将Dart代码执行在另一个线程的方式)、事件循环和异步编程。除非你创建一个 ` Isolate` ，你的Dart代码将一直在主UI线程中执行，并由事件循环驱动。Flutter的事件循环相当于iOS中的主循环，也就是说 ` Looper` 绑定在主线程上。

Dart的单线程模型并不意味着你必须将一切代码作为一个导致UI卡顿的阻塞块来执行。相反，你可以使用Dart提供的异步功能比如说： ` async/awiat` 来执行异步任务。

比如说，你可以使用 ` asyn/await` 执行网络代码和繁重的工作而避免UI卡顿。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1da3317de6c?imageView2/0/w/1280/h/960/ignore-error/1)

一旦网络请求结束，通过调用 ` setState()` 更新UI，触发当前widget的子树和更新数据。

下面例子异步加载数据并展示在 ` ListView` s上：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1e17d5b588f?imageView2/0/w/1280/h/960/ignore-error/1)

参考下一节了解如何在后台线程执行任务，与iOS有何不同。

#### 如何将任务放到后台线程 ####

由于Flutter的单线程模型和事件循环，你不用担心线程管理或者开启后台线程。你可以放心的使用 ` async/await` 方法执行I/O操作，比如访问磁环或者请求网络。另一方面，如何你想执行复杂的计算而使CPU持续的处于繁忙状态，你可以将任务已到 ` Isolate` 而避免阻塞事件循环。

对于iOS操作，将方法声明为 ` async` 方法，使用 ` await` 等待耗时任务完成。

` loadData() async { String dataURL = "https://jsonplaceholder.typicode.com/posts" ; http.Response response = await http. get (dataURL); setState(() { widgets = json.decode(response.body); }); } 复制代码`

这是对常的I/O操作如网络请求，访问数据库的常规操作。

但是，当你处理大量数据的时候这仍然可能会导致UI挂起。在Flutter中，使用 ` Isolate` 来使用CPU多核的优势来执行耗时任务或者计算密集型任务。

` Isolates` 是分离线程，它不和主线程共享任何堆内存，这也就意味着，你不能访问主线程中的变脸，或者直接调用 ` setState()` 更新主线程。 ` Isolates` 正如其名，不能共享内存。

下面代码展示了一个简单的 ` isolate` ， 如何将数据返回到主线程并更新UI的。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1e44bc5ccce?imageView2/0/w/1280/h/960/ignore-error/1)

上面代码中， ` dataLoader()` 是 ` Isolate` ，它在一个独立的线程中执行。在这个isolate中你可以执行CPU密集型任务如解析JSON，或者执行浮躁的数学计算任务，如加密或者信号处理。

你可以执行完整代码，如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1e6c6a51dd6?imageView2/0/w/1280/h/960/ignore-error/1)

#### 如何发生网络请求 ####

在Flutter中使用流行的第三方库 ` http` package ( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Fhttp ) 来请求网络是非常简单的。它抽象了大量的本需要你自己实现的操作，使得发送请求非常简单。

为了使用 ` http` 这个框架，你需要在 ` pubspec.yaml` 中增加依赖。

` dependencies: ... http: ^0.11.3+16 复制代码`

为了发起网络请求，在 ` async` 方法 ` http.get()` 前添加 ` await` 。

` import 'dart:convert' ; import 'package:flutter/material.dart' ; import 'package:http/http.dart' as http; [...] loadData() async { String dataURL = "https://jsonplaceholder.typicode.com/posts" ; http.Response response = await http. get (dataURL); setState(() { widgets = json.decode(response.body); }); } } 复制代码`

#### 如何展示耗时任务的进度 ####

在iOS中，当在后台执行一个耗时任务的时候，你通过会使用 ` UIProgressView` 展示进度。

在Flutter中，使用 ` ProgressIndicator` 组件。通过给它传递一个布尔标识来控制它的展示，告诉Flutter去更新它的状态在耗时任务执行之前和执行结束之后隐藏掉它。

在下面的例子中，build方法被分割为三个不同方法。如果 ` showLoadingDialog()` 是true，那就渲染 ` ProgressIndicator` 否则使用网络返回的数据渲染 ` ListView` 。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1eec9c77840?imageView2/0/w/1280/h/960/ignore-error/1)

### 项目结构、本地化、依赖和资源管理 ###

#### 如何在Flutter中管理图片，如何放置多种分辨率的图片 ####

与iOS将图片和资源作为不同的类型来处理不同的是Flutter中只有一种assets。iOS中资源被放在Image.xcassert中文件中，而Flutter中放在assets文件中。与iOS一样，assets是许多类型的文件，不仅仅是图片，比如说你可以将json文件放到my-assets文件夹中。

` my-assets/data.json 复制代码`

在 ` pubspec.yaml` 文件中声明：

` assets: - my-assets/data.json 复制代码`

然后就可以在代码中使用 ` AssetBunlde` 访问：

` import 'dart:async' show Future; import 'package:flutter/services.dart' show rootBundle; Future< String > loadAsset() async { return await rootBundle.loadString( 'my-assets/data.json' ); } 复制代码`

对于图片，Flutter和iOS的格式一样，图片可以是1倍图，2倍图，3倍图或者其他任何倍数。这些所谓的 ` devicePixelRatio` ( https://link.juejin.im?target=https%3A%2F%2Fapi.flutter.dev%2Fflutter%2Fdart-ui%2FWindow%2FdevicePixelRatio.html ) 表示的是物理像素到单个逻辑像素的比率。

Assets可以被放到任何类型的文件夹中，Flutter中没有事先预定义文件的结构。在 ` pubSpec.yaml` 文件中声明assets，然后Flutter就能识别出来。

比如说：将 ` my_icon.png` 放置到Flutter项目中，你可能把存储的文件夹叫作images。把相关系数的图片放在不同的子文件家中，如下：a

` images/my_icon.png // Base: 1.0x image images/ 2.0 x/my_icon.png // 2.0x image images/ 3.0 x/my_icon.png // 3.0x image 复制代码`

接下来在 ` pubspec.yaml` 中声明图片

` assets: - images/my_icon.png 复制代码`

你现在就可以使用AssetImage返回图片

` return AssetImage( "images/a_dot_burr.jpeg" ); 复制代码`

或者直接使用Image组件

` @override Widget build(BuildContext context) { return Image.asset( "images/my_image.png" ); } 复制代码`

更多细节参考 [Adding Assets and Images in Flutter]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Fui%2Fassets-and-images ) 。

#### 如何存放字符串，如何管理本地化 ####

iOS中，我们使用 ` Localizable.strings` 文件管理本地化字符串，而Flutter中没有专门的模块处理本地化字符串，所以最好的办法就是将字符串统一放到一个类中，以静态字段的形式存储。如下：

` class Strings { static String welcomeMessage = "Welcome To Flutter" ; } 复制代码`

访问方式如下：

` Text(Strings.welcomeMessage) 复制代码`

默认情况下，Flutter只支持英文字符串，如果你想支持其他语言，可以通过引入 ` flutter_localizations` 库。 同时你需要将Dart的 ` intl` 包以便支持 i10n 机制，比如日期/时间格式化。

` dependencies: # ... flutter_localizations: sdk: flutter intl: "^0.15.6" 复制代码`

为了使用 ` flutter_localizations` ，需要在App widget上指定 ` localizationsDelegates` 和 ` supportedLocales` 属性。

` import 'package:flutter_localizations/flutter_localizations.dart' ; MaterialApp( localizationsDelegates: [ // Add app-specific localization delegate[s] here GlobalMaterialLocalizations.delegate, GlobalWidgetsLocalizations.delegate, ], supportedLocales: [ const Locale( 'en' , 'US' ), // English const Locale( 'he' , 'IL' ), // Hebrew // ... other locales the app supports ], // ... ) 复制代码`

代理中包含了实际的本地化值， ` supportedLocales` 定义了要支持那些语言的本地化。上面的例子使用的是 ` MaterialApp` , 它既有针对基本Widget的本地化值 ` GlobalWidgetsLocalizations` ，也有针对Material widget的 ` MaterialWidgetsLocalizations` 本地化。如果你的App使用的是 ` WidgetApp` ，那么后者就不需要了。值得注意的是这两个代理都包含默认值，但如果你想让你的App本地化，你扔需要提供一个或者多个代理作为你的App本地化副本。

当初始化完成的时候， ` WidgetsApp` 或者 ` MaterialApp` 使用你指定的代理为你创建了一个 ` Localizations` widget。你可从 ` Localizations` Widget中随时访问当前设备的本地化信息，或者使用 ` window.locale` 。

为了访问本地化资源，使用 ` Localizations.of()` 方法访问有给定的delegate提供的特有的本地化类。使用 ` intl_translation` 取出翻译副本到 [arb]( https://link.juejin.im?target=https%3A%2F%2Fcode.google.com%2Fp%2Farb%2Fwiki%2FApplicationResourceBundleSpecification ) 文件中。将它们引入App中，并用 ` intl` 来使用它们。

更多国际化和本地化的内容参考： [internationalization guide]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Faccessibility-and-localization%2Finternationalization ) ，它包含了不使用 ` intl` 示例代码。

需要注意的是：Flutter1.0 beta2 之前 fullter中定义的资源文件不能被原生访问，同时原生定义的资源不能被flutter访问，因为它们存储在不能的文件目录下。

#### 如何管理依赖 ####

在iOS中，我们将依赖添加到 ` Podfile` 文件中，Flutter使用的是Dart语言构建的系统和 ` Pub` 包管理器操作依赖。这些工具将原生 Android 和 iOS 包装应用程序的构建委派给相应的构建系统。

如果在你的Flutter项目中iOS目录下包含Podfile，只需要使用它添加iOS原生的依赖。使用 ` pubspec.yaml` 声明Flutter 中的外部依赖。 [Pub]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fflutter%2Fpackages%2F ) 网站可以找到一些比较好用的第三方依赖。

### 视图控制器 ###

#### Flutter中与ViewControllers相等的元素是什么？ ####

在iOS中， ` ViewController` 表示用户界面的一部分，通常表示一个屏幕或者部分屏幕。多个ViewController组合在一起构造复杂的用户界面，并帮助你规整应用的UI部分。在Flutter中，这项工作落在了Widget头上，正如导航那一个章节提到的，屏幕由Widget所表示，因"一切都是Widget"。使用 ` Navigator` 在不同的路由间切换表示不同的屏幕或者页面或者表示不同的状态或者渲染相同的数据。

#### 如何监听iOS的生命周期事件 ####

在iOS中，你可以重写 ` ViewController` 中的方法来捕获视图的生命周期，或者在 ` AppDelegate` 中注册生命周期的回调。在Flutter中没有这两个概念，但是我们可以通过hook ` WidgetsBinding` 并在 ` didChangeAppLifecycleState()` 方法中监听生命周期事件。

能够监听到的生命周期事件如下：

* Inactive — 应用程序处于不活跃状态，不能相应用户输入。该事件只在iOS中有效。
* paused — 应用程序当前不可用，不响应用户输入，但是还在后台运行。
* resumed — 应用程序可用，并能响应用户输入。
* suspending — 应用程序暂时被挂起。该事件只在Android系统上有效。

更多细节参考： [AppLifecycleStatus documentation]( https://link.juejin.im?target=https%3A%2F%2Fapi.flutter.dev%2Fflutter%2Fdart-ui%2FAppLifecycleState-class.html ) 。

### 布局 ###

#### Flutter中的 ` UITableView` 和 ` UICollectionView` ####

Flutter中使用 ` ListView` 实现iOS中的 ` UITableView` 和 ` UICollectionView` 。实现代码如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1f5a5e75384?imageView2/0/w/1280/h/960/ignore-error/1)

#### 如何知道那个cell被点击 ####

在iOS中，通过实现 ` tableView:didSelectRowAtIndexPath:` 方法来相应cell的点击事件，在Flutter中，使用所包含的widget本身提供的事件来处理相应。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1f7226b3be5?imageView2/0/w/1280/h/960/ignore-error/1)

#### 如何动态更新ListView ####

在iOS中，我们使用 ` reloadData` 来刷新表格视图。

在Flutter中，如果更新setState()中的小部件列表，你会发现列表数据没有发生变化。这是因为当调用setState()时，Flutter呈现引擎会查看widget树以查看是否有任何更改。当它到达ListView时，它执行==检查，并确定两个ListView是相同的。没有任何改变，因此不需要更新。

在setState()方法内创建一个新List是更新ListView的一个简单的方法。并将旧列表中的数据复制到新列表中。虽然这种方法很简单，但不建议用于大型数据集，如下一个示例所示。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1fb312c3db6?imageView2/0/w/1280/h/960/ignore-error/1)

我们推荐使用 ` ListView.Builder` 来构建列表，它比较高效。当你的列表包含大量数据的列表时，此方法非常有用。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a1fcd38d4992?imageView2/0/w/1280/h/960/ignore-error/1)

与创建一个 ` ListView` 不同的是，创建 ` ListView.builder` 携带两个参数：列表的初始长度和 ` ItemBuilder` 方法。

` ItemBuilder` 方法和iOS中的table或者collection的 ` cellForItemAt` 代理相似，一样的携带一个位置，并返回该位置需要渲染的cell。

最后也是最重要的，onTap方法并没有重新创建一个list，而是 `.add` 了一个Widget。

#### 如何使用类似ScrollView的功能 ####

` @override Widget build(BuildContext context) { return ListView( children: <Widget>[ Text( 'Row One' ), Text( 'Row Two' ), Text( 'Row Three' ), Text( 'Row Four' ), ], ); } 复制代码`

### 手势检测和触摸事件处理 ###

#### 如何向widget添加一个事件监听 ####

如果widget支持事件处理，如RaisedButton，可以直接将相应方法传递给对应的属性，如RaisedButton的onPressed。

` Widget build(BuildContext context) { return RaisedButton( onPressed: () { print ( "click" ); }, child: Text( "Button" ), ); } 复制代码`

如果widget不支持事件处理，可以使用 ` GestureDetector` 包裹一下，然后给 ` onTap` 属性传递一个方法。

` Widget build(BuildContext context) { return Scaffold( appBar: AppBar( title: Text( 'Sample App' ), ), body: Center( child: GestureDetector( child: FlutterLogo( size: 200 , ), onTap: () { print ( 'taped' ); }, ), )); } 复制代码`

#### 如何处理widget上的其他类型的事件 ####

我们可以使用 ` GestureDetector` 来实现如下事件的监听：

* 单击

* ` onTapDown` — 按下手势事件
* ` onTapUp` — 抬起事件
* ` onTap` — 点击事件
* ` onTapCancel` — 取消点击事件， ` onTapDown` 发生，但onTap没有发生。

* 双击

* ` onDoubleTap` — 双击事件

* 长按

* ` onLongPress` — 长按事件

* 垂直拖动

* ` onVerticalDragStart` —开始垂直移动
* ` onVerticalDragUpdate` — 垂直移动进行中。
* ` onVerticalDragEnd` — 垂直移动结束。

* 水平拖动

* ` onHorizontalDragStart` — 开始水平移动。
* ` onHorizontalDragUpdate` — 水平移动进行中。
* ` onHorizontalDragEnd` — 水平移动结束。

下面代码展示了使用 ` GestureDetector` 实现双击事件：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a20460ad1d20?imageView2/0/w/1280/h/960/ignore-error/1)

运行效果：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a20628277c27?imageslim)

### 主题和文本 ###

#### 如何为应用程序设置主题 ####

Flutter提供了一套完美符合Material Design的主题，它帮你处理了大多数需要你自己处理的样式和主题。

为了在你的App中充分发挥Material组件的优势，在顶层组件上声明MaterialApp，作为你的应用的入口。MaterialApp 是一个便利的组件，它包含了许多App通常需要的Materail Desigin风格的组件。它通过由给WidgetsApp增加MD功能实现的。

同时 Flutter 足够地灵活和富有表现力来实现任何其他的设计语言。在 iOS 上，你可以用 [Cupertino library]( https://link.juejin.im?target=https%3A%2F%2Fdocs.flutter.io%2Fflutter%2Fcupertino%2Fcupertino-library.html ) 来制作遵守 [Human Interface Guidelines]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fios%2Fhuman-interface-guidelines%2Foverview%2Fthemes%2F ) 的界面。查看这些 widget 的集合，请参阅 [Cupertino widgets gallery]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fwidgets%2Fcupertino%2F ) 。

你也可以在你的 App 中使用 WidgetApp，它提供了许多相似的功能，但不如 ` MaterialApp` 那样丰富。

对任何子组件定义颜色和样式，可以给 ` MaterialApp` widget 传递一个 ` ThemeData` 对象。举个例子，在下面的代码中，primary swatch 被设置为蓝色，并且文字的选中颜色是红色：

` class SampleApp extends StatelessWidget { @override Widget build(BuildContext context) { return MaterialApp( title: 'Sample App' , theme: ThemeData( primarySwatch: Colors.blue, textSelectionColor: Colors.red ), home: SampleAppPage(), ); } } 复制代码`

#### 如何在Text widget上使用自定义字体 ####

在 iOS 中，你在项目中引入任意的 ` ttf` 文件，并在 ` info.plist` 中设置引用。在 Flutter 中，在文件夹中放置字体文件，并在 ` pubspec.yaml` 中引用它，就像添加图片那样。

` fonts: - family: MyCustomFont fonts: - asset: fonts/MyCustomFont.ttf - style: italic 复制代码`

然后在你的 ` Text` widget 中指定字体：

` @override Widget build(BuildContext context) { return Scaffold( appBar: AppBar( title: Text( "Sample App" ), ), body: Center( child: Text( 'This is a custom font text' , style: TextStyle(fontFamily: 'MyCustomFont' ), ), ), ); } 复制代码`

#### 如何设置Text widget的样式 ####

除了字体以外，你也可以给 Text widget 的样式元素设置自定义值。 ` Text` widget 接受一个 ` TextStyle` 对象，你可以指定许多参数，如下：

* ` color`
* ` decoration`
* ` decorationColor`
* ` decorationStyle`
* ` fontFamily`
* ` fontSize`
* ` fontStyle`
* ` fontWeight`
* ` hashCode`
* ` height`
* ` inherit`
* ` letterSpacing`
* ` textBaseline`
* ` wordSpacing`

### 表单输入 ###

#### 表单在Flutter中如何工作的，如何取回用户输入的值 ####

在iOS中，我们通常在用户提交的时候获取组件上的内容，对于具有使用独立状态的不可变组件的Flutter来讲，你可能会好奇如何获取用户输入内容。

对于表单操作而言，与其他功能一样也是通过特定的Widget实现的。通过使用 ` TextField` 或者 ` TextFormField` 可以通过 ` TextEditingController` ( https://link.juejin.im?target=https%3A%2F%2Fapi.flutter.dev%2Fflutter%2Fwidgets%2FTextEditingController-class.html ) 取回输入内容。

示例代码如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a20a8ad02217?imageView2/0/w/1280/h/960/ignore-error/1)

运行效果

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a20e3dac82e5?imageslim)

更多信息参考： [Flutter Cookbook]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fcookbook ) 的 [Retrieve the value of a text field]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fcookbook%2Fforms%2Fretrieve-input )

#### 如何实现类似文本输入框占位符的功能 ####

通过给decoration属性传递一个 ` InputDecoration` 对象来给TextField实现占位符的功能。

` body: Center( child: TextField( decoration: InputDecoration(hintText: "This is a hint" ), ), ) 复制代码`

#### 如何展示验收错误信息 ####

与上面代码一样，只不过是再添加一个 ` errorText` 字段，通过state控制错误信息的提示。

示例代码如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a21061b07fd7?imageView2/0/w/1280/h/960/ignore-error/1)

运行效果：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a213ad6881bf?imageslim)

### 与硬件、第三方服务和平台的交互 ###

#### 如何与平台和平台原生代码交互 ####

Flutter不是在直接在平台下运行代码的，相反，由Dart语言构建的FlutterApp在设备本机运行，"回避"平台提供的SDK。比如说：在Dart中发送一个网络请求，它是直接在Dart上下文中执行的，而不适用我们在写原生App的时候所使用的Android或者iOSAPI。我们的FlutterApp仍然被原生app的ViewController当做一个View所持有，但我们不用直接访问ViewController或者原生框架。

这并不意味着Flutter应用不能与原生API或者其他你写的原生代码交互。Flutter提供了 [platform channels]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Fplatform-integration%2Fplatform-channels ) ，它可以与持有你Flutter视图的VIewController通信或者交换数据。 [platform channels]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fplatform-channels%2F ) 本质上是一个异步通信机制，桥接了Dart代码和其宿主ViewController，iOS框架。比如说。你可以用platform channels执行一个原生的函数，或者是从设备的传感器中获取数据。

除了直接使用 [platform channels]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fplatform-channels%2F ) 之外，你还可以使用一系列预先制作好的 [plugins]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fusing-packages%2F ) 。例如，你可以直接使用插件来访问相机胶卷或是设备的摄像头，而不必编写你自己的集成层代码。你可以在 [Pub]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2F ) 上找到插件，这是一个 Dart 和 Flutter 的开源包仓库。其中一些包可能会支持集成 iOS 或 Android，或两者均可。

如果你在 Pub 上找不到符合你需求的插件，你可以 [自己编写]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fdeveloping-packages%2F ) ，并且 [发布在 Pub 上]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fdeveloping-packages%2F%23publish ) 。

#### 如何访问GPS传感器 ####

使用 [geolocator]( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Fgeolocator )

#### 如何访问相机 ####

使用 [image_picker]( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Fimage_picker )

#### 如何使用FaceBook登陆 ####

使用 [flutter_facebook_login]( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Fflutter_facebook_login )

#### 如何使用Firebase ####

大多数 Firebase 特性被 [first party plugins]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fflutter%2Fpackages%3Fq%3Dfirebase ) 包含了。这些第一方插件由 Flutter 团队维护：

* [firebase_admob]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_admob ) : Firebase AdMob
* [firebase_analytics]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_analytics ) : Firebase Analytics
* [firebase_auth]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_auth ) : Firebase Auth
* [firebase_core]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_core ) : Firebase’s Core package
* [firebase_database]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_database ) : Firebase RTDB
* [firebase_storage]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_storage ) : Firebase Cloud Storage
* [firebase_messaging]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Ffirebase_messaging ) : Firebase Messaging (FCM)
* [cloud_firestore]( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Fcloud_firestore ) : Firebase Cloud Firestore 你也可以在 Pub 上找到 Firebase 的第三方插件。

#### 如何创建原生集成层代码 ####

如果有一些 Flutter 和社区插件遗漏的平台相关的特性，可以根据 [developing packages and plugins]( https://link.juejin.im?target=https%3A%2F%2Fflutter.io%2Fdeveloping-packages%2F ) 页面构建自己的插件。 Flutter 的插件结构，简要来说，就像 Android 中的 Event bus。你发送一个消息，并让接受者处理并反馈结果给你。在这种情况下，接受者就是在 Android 或 iOS 上的原生代码。

### 数据库和本地存储 ###

#### 如何在Flutter中使用UserDefaults ####

在iOS中，我们可以使用 ` UserDefaults` 来存储键值对集合，在Flutter中，可以使用 [Shared Preferences plugin]( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Fshared_preferences ) 插件来显示类似的功能。 这个插件包装了 ` UserDefaults` 和Android 上的 ` SharedPreferences` 。

### Flutter中和Coredata相等的功能。 ###

可以使用 [SQFlite]( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Fsqflite ) 插件实现iOS中CoreData相关的功能。

### 通知 ###

#### 如何设置推送通知 ####

在iOS，你需要在开发者网站上注册app以便获取推送权限。在Flutter中使用 ` firebase_messaging` 插件可以实现推送。 更多关于使用 ` Firebase Cloud Messaging API` 的文档请参考： ` firebase_messaging` ( https://link.juejin.im?target=https%3A%2F%2Fpub.dev%2Fpackages%2Ffirebase_messaging )

### 参考 ###

本文主要参考Flutter官方文档，Flutter中文网。 由于排版原因，文中我使用了图片的形式展示代码，如果你需要源码，可以关注我的公众号，回复关键字"flutter"获取相关代码。

**本文首发自微信公众号：RiverLi**