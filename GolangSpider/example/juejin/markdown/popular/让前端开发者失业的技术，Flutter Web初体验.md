# 让前端开发者失业的技术，Flutter Web初体验 #

> 
> 
> 
> Flutter是一种新型的“客户端”技术。它的最终目标是替代包含几乎所有平台的开发：iOS，Android，Web，桌面；做到了一次编写，多处运行。掌握Flutter
> web可能是Web前端开发者翻盘的唯一机会。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/22/16adede515644cd0?imageView2/0/w/1280/h/960/ignore-error/1)

在前些日子举办的Google IO 2019 年度开发者大会上，Flutter web作为一个很亮眼的技术受到了开发者的追捧。这是继Flutter支持Android、IOS等设备之后，又一个里程碑式的版本，后续还会支持windows、linux、Macos、chroms等其他嵌入式设备。Flutter本身是一个类似于RN、WEEX、hHybrid等多端统一跨平台解决方案，真正做到了一次编写，多处运行，它的发展超出了很多人的想象，值得前端开发者去关注，今天我们来体验一下Flutter Web。

## 概览 ##

先了解一下Flutter， 它是一个由谷歌开发的开源移动应用软件开发工具包，用于为Android和iOS开发应用，同时也将是Google Fuchsia下开发应用的主要工具。自从FLutter 1.5.4版本之后，支持了Web端的开发。它采用Dart语言来进行开发，与JavaScript相比，Dart在 JIT（即时编译）模式下，速度与 JavaScript基本持平。但是当Dart以 AOT模式运行时，Dart性能要高于JavaScript。

Flutter内置了UI界面，与Hybrid App、React Native这些跨平台技术不同，Flutter既没有使用WebView，也没有使用各个平台的原生控件，而是本身实现一个统一接口的渲染引擎来绘制UI，Dart直接编译成了二进制文件，这样做可以保证不同平台UI的一致性。它也可以复用Java、Kotlin、Swift或OC代码，访问Android和iOS上的原生系统功能，比如蓝牙、相机、WiFi等等。我们公司的Now直播、企鹅辅导等项目、阿里的闲鱼等商业化项目已经大量在使用。

## 架构 ##

![Flutter 的 Mobile 架构](https://user-gold-cdn.xitu.io/2019/5/22/16adede8b8b4ec92?imageView2/0/w/1280/h/960/ignore-error/1)

Flutter的顶层是用drat编写的框架，包含Material（Android风格UI）和Cupertino（iOS风格）的UI界面，下面是通用的Widgets（组件），之后是一些动画、绘制、渲染、手势库等。 框架下面是引擎，主要用C / C ++编写，引擎包含三个核心库，Skia是Flutter的2D渲染引擎，它是Google的一个2D图形处理函数库，包含字型、坐标转换，以及点阵图，都有高效能且简洁的表现。Skia是跨平台的，并提供了非常友好的API。第二是Dart 运行时环境以及第三文本渲染布局引擎。 最底层的嵌入层，它所关心的是如何将图片组合到屏幕上，渲染变成像素。这一层的功能是用来解决跨平台的。

了解了FLutter 之后，我来说一下今天的重头戏，Flutter for Web。要想知道Flutter为什么能在web上运行，得先来看看它的架构。

![Flutter 的 web架构](https://user-gold-cdn.xitu.io/2019/5/22/16adedee57be23f0?imageView2/0/w/1280/h/960/ignore-error/1)

通过对比，可以发现，web框架层和mobile的几乎一模一样。因此只需要重新实现一下引擎和嵌入层，不用变动Flutter API就可以完全可以将UI代码从Android / IOS Flutter App移植到Web。Dart能够使用Dart2Js编译器把Dart代码编译成Js代码。大多数原生App元素能够通过DOM实现，DOM实现不了的元素可以通过Canvas来实现。

## 安装 ##

Flutter Web开发环境搭建，以我的windows环境为例进行讲解，其他环境类似，安装环境比较繁琐，需要耐心，有Android开发经验最好。

### 1、在 **Windows平台** 开发的话，官方的环境要求是Windows 7 SP1或更高版本（64位）。 ###

### 2、 **Java环境** ，安装Java 1.8 + 版本之上，并配置环境变量，因为android开发依赖Java环境。 ###

对于Java程序开发而言，主要会使用JDK的两个命令：javac.exe、java.exe。路径：C:\Java\jdk1.8.0_181bin。但是这些命令由于不属于windows自己的命令，所以要想使用，就需要进行路径配置。单击“计算机-属性-高级系统设置”，单击“环境变量”。在“系统变量”栏下单击“新建”，创建新的系统环境变量（或用户变量，等效）。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adedf60530bf6b?imageView2/0/w/1280/h/960/ignore-error/1)

(1)新建->变量名"JAVA_HOME"，变量值"C:\Java\jdk1.8.0_181"（即JDK的安装路径） (2)编辑->变量名"Path"，在原变量值的最后面加上“;%JAVA_HOME%\bin;%JAVA_HOME%\jre\bin” (3)新建->变量名“CLASSPATH”,变量值“.;%JAVA_HOME%\lib;%JAVA_HOME%\lib\dt.jar;%JAVA_HOME%\lib\tools.jar”

### 3、 **Android Studio编辑器** ，安装Android Studio, 3.0或更高版本。我们需要用它来导入Android license和管理Android SDK以及Android虚拟机。（默认安装即可） ###

安装完成之后设置代理，左上角的File-》setting-》搜索proxy，设置公司代理，用来加速下载Android SDK。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adedf9619db405?imageView2/0/w/1280/h/960/ignore-error/1)

之后点击右上角方盒按钮（SDK Manager），用来选择安装SDK版本，最好选Android 9版本，API28，会有一个很长时间的下载过程。SDK是开发必须的代码库。默认情况下，Flutter使用的Android SDK版本是基于你的 adb （Android Debug Bridge，管理连接手机，已打包在SDK）工具版本。 如果您想让Flutter使用不同版本的Android SDK，则必须将该 ANDROID_HOME 环境变量设置为SDK安装目录。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adedfcafe75ca4?imageView2/0/w/1280/h/960/ignore-error/1)

右上角有个小手机类型的按钮（AVD Manager），用来设置Android模拟器，创建一个虚拟机。如果你有一台安卓手机，也可以连接USB接口，替代虚拟机。这个过程是调试必须的。安装完成之后，在 AVD (Android Virtual Device Manager) 中，点击工具栏的 Run。模拟器启动并显示所选操作系统版本或设备的启动画面。代表了正确安装。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adedffed6a5691?imageView2/0/w/1280/h/960/ignore-error/1)

### 4、 **安装Flutter SDK** ###

下载Flutter SDK有多种方法，看看哪种更适合自己： Flutter官网下载最新Beta版本的进行安装： [flutter.dev/docs/develo…]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Ftools%2Fsdk%2Freleases ) 也可Flutter github项目中去下载，地址为： [github.com/flutter/flu…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fflutter%2Fflutter%2Freleases ) 版本越新越好，不要低于1.5.4。

将安装包zip解压到你想安装Flutter SDK的路径（如：C:\src\flutter；注意，不要将flutter安装到需要一些高权限的路径如C:\Program Files\）。记住，之后往环境变量的path中添加；C:\src\flutter\bin，以便于你能在命令行中使用flutter。

使用镜像 由于在国内安装Flutter相关的依赖可能会受到限制，Flutter官方为中国开发者搭建了临时镜像，大家可以将如下环境变量加入到用户环境变量中： ` PUB_HOSTED_URL：https://pub.flutter-io.cn` ` FLUTTER_STORAGE_BASE_URL： https://storage.flutter-io.cn`

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee02bc4a6ef8?imageView2/0/w/1280/h/960/ignore-error/1)

### 5、安装Dart与Pub。安装webdev、stagehand ###

Pub是Dart的包管理工具，类似npm，捆绑安装。 Dart安装版地址： [www.gekorm.com/dart-window…]( https://link.juejin.im?target=http%3A%2F%2Fwww.gekorm.com%2Fdart-windows%2F ) 默认安装即可，安装之后记住Dart的路径，并且配置到环境变量path中，以便于可以在命令行中使用dart与pub，默认的路径是：C:\Program Files\Dart\dart-sdk\bin 先安装stagehand，stagehand是创建项目必须的工具。查看一下 ` C:\Users\chunpengliu\AppData\Roaming\Pub\Cache\bin` 目录下是否包含stagehand和webdev，如果有，添加到环境变量的path里面，如果没有，按下面方法安装：

` pub global activate stagehand 复制代码`

webdev是一个类似于Koa的web服务器，执行以下命令安装

` pub global activate webdev # or flutter packages pub global activate webdev 复制代码`

### 6、配置编辑器安装Flutter和Dart插件 ###

Flutter插件是用来支持Flutter开发工作流 (运行、调试、热重载等)。 Dart插件 提供代码分析 (输入代码时进行验证、代码补全等)。Android Studio的设置在File-》setting-》plugins-》搜索Flutter和Dart，安装之后重启。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee07be21b40a?imageView2/0/w/1280/h/960/ignore-error/1)

VS code的设置在extension-》搜索Flutter和Dart，安装之后重启。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee0f8c28875f?imageView2/0/w/1280/h/960/ignore-error/1)

### 7、运行 flutter doctor ###

打开一个新的命令提示符或PowerShell窗口并运行以下命令以查看是否需要安装任何依赖项来完成安装：

` flutter doctor 复制代码`

这是一个漫长的过程，flutter会检测你的环境，并安装所有的依赖，直至：No issues found！，如果有缺失，会就会再那一项前面打x。你需要一一解决。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee11d69a25f5?imageView2/0/w/1280/h/960/ignore-error/1)

一切就绪！

## 创建应用 ##

### 1、启动 VS Code ###

调用 View>Command Palette…（快捷键ctrl+shift+p） 输入 ‘flutter’, 然后选择 ‘Flutter: New web Project’

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee1371a5f09c?imageView2/0/w/1280/h/960/ignore-error/1)

输入 Project 名称 (如flutterweb), 然后按回车键 指定放置项目的位置，然后按蓝色的确定按钮 等待项目创建继续，并显示main.dart文件。到此，一个Demo创建完成。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee15c8e5c16d?imageView2/0/w/1280/h/960/ignore-error/1)

我们看到了熟悉的HTML文件以及项目入口文件main.dart。 web目录下的index.html是项目的入口文件。main.dart初始化文件，图片相关资源放在此目录。 lib目录下的main.dart，是主程序代码所在的地方。 每个pub包或者Flutter项目都包含一个pubspec.yaml。它包含与此项目相关的依赖项和元数据。 analysis_options.yaml是配置项目的lint规则。 /dart_tool 是项目打包运行编译生成的文件，页面主程序main.dart.js就在其中。

### 2、调试Demo，打开命令行，进入到项目根目录，执行： ###

` webdev flutterweb 复制代码`

编译、打包完成之后，自动启动（或者按F5）默认浏览器，看一下转换后的HTML页面结构：

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee1d0b25c5a4?imageView2/0/w/1280/h/960/ignore-error/1)

lib/main.dart是主程序，源码非常简单，整个页面用widgets堆叠而成，区别于传统的html和css。

` import 'package:flutter_web/material.dart' ; void main() => runApp(MyApp()); class MyApp extends StatelessWidget { @override Widget build(BuildContext context) { return MaterialApp( title: 'Flutter Demo' , theme: ThemeData( primarySwatch: Colors.blue, ), home: MyHomePage(title: 'Flutter Demo Home Page' ), ); } } class MyHomePage extends StatelessWidget { MyHomePage({Key key, this.title}) : super(key: key); final String title; @override Widget build(BuildContext context) { return Scaffold( appBar: AppBar( title: Text(title), ), body: Center( child: Column( mainAxisAlignment: MainAxisAlignment.center, children: <Widget>[ Text( 'Hello, World!' , ), ], ), ), ); } } 复制代码`

区别与flutter App应用，我们导入的是flutter_web/material.dart库而非flutter/material.dart，这是因为目前App的接口并非和Web的完全通用，不过随着谷歌开发的继续，它们最终会被合并到一块。 打开pubspec.yaml（类似于package.json）,可以看到只有两个依赖包flutter_web和flutter_web_ui，这两个都已在github上开源。dev的依赖页非常少，两个编译相关的包，和一个静态文件分析包。

` name: flutterweb description: An app built using Flutter for web environment: # You must be using Flutter >=1.5.0 or Dart >=2.3.0 sdk: '>=2.3.0-dev.0.1 <3.0.0' dependencies: flutter_web: any flutter_web_ui: any dev_dependencies: build_runner: ^1.4.0 build_web_compilers: ^2.0.0 pedantic: ^1.0.0 dependency_overrides: flutter_web: git: url: https://github.com/flutter/flutter_web path: packages/flutter_web flutter_web_ui: git: url: https://github.com/flutter/flutter_web path: packages/flutter_web_ui 复制代码`

## 实战 ##

接下来，我们创建一个具有图文功能的下载，根据实例来学习flutter，我们将实现下图的页面。它是一个上下两栏的布局，下栏又分为左右两栏。

![](https://user-gold-cdn.xitu.io/2019/5/22/16adee1fa411991f?imageView2/0/w/1280/h/960/ignore-error/1)

第一步：更改主应用内容，打开lib/main.dart文件，替换class MyApp，首先是根组件MyApp，它是一个类组件继承自无状态组件，是项目的主题配置，在home属性中调用了Home组件：

` class MyApp extends StatelessWidget { // 应用的根组件 @override Widget build(BuildContext context) { return MaterialApp( title: '腾讯新闻客户端下载页' , //meta 里的titile debugShowCheckedModeBanner: false , // 关闭调试bar theme: ThemeData( primarySwatch: Colors.blue, // 页面主题 Material风格 ), home: Home(), // 启动首页 ); } } 复制代码`

第二步，在Home类中，是我们要渲染的页面顶导，运用了AppBar组件，它包括了一个居中的页面标题和居右的搜索按钮。文本可以像css一样设置外观样式。

` class Home extends StatelessWidget { @override Widget build(BuildContext context) { return Scaffold( backgroundColor: Colors.white, appBar: AppBar( backgroundColor: Colors.white, elevation: 0.0, centerTitle: true , title: Text( // 中心文本 "下载页" , style: TextStyle(color: Colors.black, fontSize: 16.0, fontWeight: FontWeight.w500), ), // 搜索图标及特性 actions: <Widget>[ Padding( padding: const EdgeInsets.symmetric(horizontal: 20.0), child: Icon( Icons.search, color: Colors.black, ), ) ], ), //调用body渲染类，此处可以添加多个方法调用 body: Stack( children: [ Body() ], ), ); } } 复制代码`

第三步，创建页面主体内容，一张图加多个文本，使用了文本组件和图片组件，页面结构采用了flex布局，由于两个Expanded的Flex值均为1，因此将在两个组件之间平均分配空间。SizedBox组件相当于一个空盒子，用来设置margin的距离

` class Body extends StatelessWidget { const Body({Key key}) : super(key: key); @override Widget build(BuildContext context) { return Row( crossAxisAlignment: CrossAxisAlignment.stretch, mainAxisAlignment: MainAxisAlignment.spaceBetween, children: <Widget>[ Expanded( // 左侧 flex: 1, child: Image.asset(// 图片组件 "background-image.jpg" , // 这是一张在web/asserts/下的背景图 fit: BoxFit.contain, ), ), const SizedBox(width: 90.0), Expanded( // 右侧 flex:1, child: Column( mainAxisAlignment: MainAxisAlignment.center, crossAxisAlignment: CrossAxisAlignment.start, children: <Widget>[ Text( // 文本组件 "腾讯新闻" , style: TextStyle( color: Colors.black, fontWeight: FontWeight.w600, fontSize: 50.0, fontFamily: 'Merriweather' ), ), const SizedBox(height: 14.0),// SizedBox用来增加间距 Text( "腾讯新闻是腾讯公司为用户打造的一款全天候、全方位、及时报道的新闻产品，为用户提供高效优质的资讯、视频和直播服务。资讯超新超全，内容独家优质，话题评论互动。" , style: TextStyle( color: Colors.black, fontWeight: FontWeight.w400, fontSize: 24.0, fontFamily: "Microsoft Yahei" ), textAlign: TextAlign.justify, ), const SizedBox(height: 20.0), FlatButton( onPressed: () {}, // 下载按钮的响应事件 color: Color(0xFFCFE8E4), shape: RoundedRectangleBorder( borderRadius: BorderRadius.circular(16.0), ), child: Padding( padding: const EdgeInsets.all(12.0), child: Text( "点击下载" , style: TextStyle(fontFamily: "Open Sans" )), ), ), ], ), ), const SizedBox(width: 100.0), ], ); } } 复制代码`

到此，页面创建结束，保存，运行webdev serve，就可以看到效果了。

## 总结 ##

FLutter web是Flutter 的一个分支，在开发完App之后，UI层面的FLutter代码在不修改的情况下可以直接编译为Web版，基本可以做到代码100%复用，体验还不错。目前Flutter web作为预览版无论从性能上、易用上还是布局上都超出了预期，触摸体验挺好，虽然体验比APP差一些，但是比传统的web要好很多。试想一下 Flutter 开发iOS 和Android的App 还免费赠送一份Web版，并且比传统的web开发出来的体验还好。Write once ，Run anywhere。何乐而不为？

我觉得随着谷歌的持续优化，等到正式版发布之后，开发体验越来越好，Flutter开发者会吃掉H5很大一部分份额。Flutter 可能会给目前客户端的开发模式带来一些变革以及分工的变化， Flutter目前的开发体验不是很好， 但是潜力很大，值得前端人员去学习。

但是目前还是有一部分问题，Flutter web是为客户端开发（尤其是安卓）人员开发准备的，对于前端理解来说学习成本有点高。目前FLutter web和 flutter 还是两个项目，编译环境也是分开的，需要在代码里面修改Flutter相关库的引用为Flutter_web，组件还不能达到完全通用，这个谷歌承诺正在解决中，谷歌的最终目标是Web、移动App、桌面端win mac linux、以及嵌入式版的Flutter代码库之间保持100%的代码可移植性。

个人感觉，开发体验还不太好，还有很多坑要去踩，版本变更很快。还有社区资源稀少的问题，需要一定长期的积累。兼容性问题，代码转换后大量使用了web components，除了chrome之外，兼容性还是有些问题。

## 安利时间 ##

我们在web开发过程中，都见过或者使用过一些奇技淫巧，这种技术我们统称为黑魔法，这些黑魔法散落在各个角落，为了方便大家查阅和学习，我们做了收集、整理和归类，并在github上做了一个项目—— [awesome-blackmargic]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FTnfe%2Fawesome-blackmagic ) ，希望各位爱钻研的开发者能够喜欢，也希望大家可以把自己的独门绝技分享出来，如果有兴趣可以给我们发pr。

如果你对Flutter感兴趣，想进一步了解Flutter，加入我们的QQ群（784383520）吧！