# Flutter基础（四）开发Flutter应用前需要掌握的Basics Widget #

> 
> 
> 
> **本文首发于公众号「刘望舒」**
> 
> 

关联系列
[ReactNative入门系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FReact-Native%25E5%2585%25A5%25E9%2597%25A8%2F )
[React Native组件]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FReact-Native%25E7%25BB%2584%25E4%25BB%25B6%2F )
[Flutter基础系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FFlutter%25E5%259F%25BA%25E7%25A1%2580%2F )

### **前言** ###

学完了Dart语言，接下来就可以学习Widget了，Flutter的UI界面就是由Widget组成的，Widget的数量繁多，因此我会用几篇文章来专门介绍它，本篇就来介绍Basics Widget。

### **1.什么是Widget** ###

Flutter的Widget的设计灵感来自于React，主要目的就是使用Widget构建UI。Widget根据其当前配置和状态来描述视图，当Widget的状态发生更改时，Widget会重建其描述。framework将根据前面的描述进行对比，以确定底层渲染树从一个状态转换到下一个状态所需的最小更改。 在Flutter中，除了Basics 的文本、图片、卡片、输入框这些基础控件，布局方式和动画等也都是由Widget组成的。通过使用不同类型的Widget，就可以实现复杂的界面。 Widget可以翻译为部件，粗略的相当于Android中的View。Widget和View不同的是：Widget具有不同的生命周期：它是不可变的，每当Widget或者其状态发生变化时，Flutter的框架都会创建一个新的Widget实例树。相比之下，Android中的View会被绘制一次，并且在invalidate调用之前不会重绘。

### **2.Widget的分类** ###

Widget的分类有很多类别，每个类别下面又包含很多Widget，主要包括以下几种类别：

* Basics：在构建第一个Flutter应用程序之前，需要知道的Basics Widget。
* Material Components：Material Design风格的Widget。
* Cupertino：iOS风格的Widget。
* Accessibility：辅助功能Widget。
* Animation and Motion：动画和动作Widget。
* Async：Flutter应用程序的异步Widget。
* Input：除了在Material Components和Cupertino中的输入Widget外，还可以接受用户输入的Widget。
* Interaction Models：响应触摸事件并将用户路由到不同的视图中。
* Layout：用于布局的Widget。
* Painting and effects：不改变布局、大小、位置的情况下为子Widget应用视觉效果。
* Scrolling：滚动相关的Widget。
* Styling：主题、填充相关Widget。
* Text：显示文本和文本样式。

Basics有些特殊，它是由Flutter官方从其他的Widget分类中选取的一些Widget组成的，这些Widget是官方建议开发者构建第一个Flutter应用程序之前，需要知道的，目的是让开发者更快的入门。比如Row属于Layout分类，它就被选进了Basics中。本文遵循了Flutter官方的意图，首先介绍Basics(Basics Widget）。

Widget更多的是以组合的形式存在，比如Container是属于Layout中的一个Widget，而Container又由LimitedBox、 ConstrainedBox、Align、 Padding、 DecoratedBox、Transform部件组成。 如果要实现Container的自定义效果，可以组合上面这些Widget以及其他简单的Widget，而不是将Container进行子类化实现。

### **3.Widget的状态分类** ###

在Android中，我们可以通过直接更改View来更新视图。但是在Flutter中，Widget是不可变的并且不会直接更新，而是必须使用Widget的状态。 Widget有两种状态分类分别是无状态的StatelessWidget和有状态的StatefulWidget，StatelessWidget是不可变的，设置以后就不可再变化，所有的值都是最终的设置。StatefulWidget可以保存自己的状态，但是Widget是不可变的，因此需要配合State来保存状态。 State拥有自己的声明周期，如下所示：

+-----------------------+--------------------------------------------------+
|         名称          |                       状态                       |
+-----------------------+--------------------------------------------------+
| initState             | create之后被insert到渲染树时调用的，只会调用一次 |
| didChangeDependencies | state依赖的对象发生变化时调用                    |
| didUpdateWidget       | Widget状态改变时候调用，可能会调用多次           |
| build                 | 构建Widget时调用                                 |
| deactivate            | 当移除渲染树的时调用                             |
| dispose               | Widget即将销毁时调用                             |
+-----------------------+--------------------------------------------------+

### **4.根Widget的种类** ###

` void main() => runApp(MyApp()); class MyApp extends StatelessWidget { @override Widget build(BuildContext context) { return MaterialApp( ... ), ), ); } } 复制代码`

上面的MaterialApp就是一个根Widget，也就是Flutter应用程序的第一个Widget，根Widget有以下几种：

* WidgetsApp： 如果需要自定义风格，可以使用WidgetsApp。
* MaterialApp：Material Design风格的Widget。
* CupertinoApp iOS风格的根Widget。

如果公司没有特殊要求，这里建议使用MaterialApp做为根Widget就可以了。

### **5.Basics Widget** ###

Basics Widget也就是Basics，主要有以下几种：

* Container：一个便利的容器Widget，可以设置Widget的背景、尺寸、定位。
* Row：在水平方向上布置子窗口Widget列表。
* Column：在垂直方向上布置子窗口Widge列表。
* Image：显示图像的Widget
* Text：单一样式的文本。
* Icon：符合Material Design设计规范的图标
* RaisedButton：符合Material Design设计规范的凸起按钮。
* Scaffold：实现Basics 的Material Design布局结构。
* Appbar：Material Design的应用栏。
* FlutterLogo：以Widget形式来展示一个Flutter图标，可以调整样式。
* Placeholder：绘制一个框，为将来添加的Widget的占位。

这里选择一些我们必须要掌握的Basics Widget来进行讲解。

#### **5.1 代码模板和主题** ####

为了更好的理解这些Basics Widget，我们需要写一些例子，这些例子需要一个代码模板，方便测试和学习。

` import 'package:flutter/material.dart' ; void main () => runApp(MyApp()); class MyApp extends StatelessWidget { @override Widget build (BuildContext context) { return MaterialApp( //1 title: 'Welcome to Flutter' , home: Scaffold( //2 appBar: AppBar( //3 title: Text( 'Basics Widget' ), ), body: Padding( padding: EdgeInsets.all( 40.0 ), child: Text( '在这里编写和测试其他Basics Widget' ), ), ), ); } } 复制代码`

上面的代码是稍微改动了官方的Hello World代码，便于测试，具体的代码含义已经在 [Flutter基础（二）Flutter开发环境搭建和Hello World]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Fflutter%2Fprimer%2F2-start.html ) 中讲过了，这里结合本文要讲的内容再说点细节。注释1处的MaterialApp属于Material Components类别中的Widget，MaterialApp中包含了实现Material Design的应用程序所需要的Widget。 注释2和3处的Scaffold和AppBar同样也是Material Components类别中的Widget，Scaffold实现了Material Design布局结构，AppBar是Material Design的应用栏，它们会在下一篇文章介绍Material Components时进行讲解。效果如下图所示：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29861bb86fb2d?imageView2/0/w/1280/h/960/ignore-error/1)

#### **5.2 文本** ####

在4.1小节中已经用了Text，还可以定义样式：

` import 'package:flutter/material.dart' ; void main () => runApp(MyApp()); class MyApp extends StatelessWidget { @override Widget build (BuildContext context) { return MaterialApp( title: 'Welcome to Flutter' , home: Scaffold( appBar: AppBar( title: Text( 'Basics Widget' ), ), body: Padding( padding: EdgeInsets.all( 60.0 ), child: Text( '文本样式' , style: TextStyle( fontSize: 16.0 , color: Colors.indigo, fontStyle: FontStyle.normal, fontWeight: FontWeight.bold, ), ), ), ), ); } } 复制代码`

这次为了便于理解列出了全部的代码， **此后的举例只列出改变的部分** 。 通过TextStyle来定义文本的样式，效果如下：

![VZrQ3V.png](https://user-gold-cdn.xitu.io/2019/6/6/16b29861bb5646b0?imageView2/0/w/1280/h/960/ignore-error/1)

#### **5.3 图片** ####

Image的构造函数有多种：

* new Image：从ImageProvider获取图片
* new Image.asset：使用key从AssetBundle获取图片
* new Image.network：加载网络图片
* new Image.file：从文件中获取图片
* new Image.memory：用于从Uint8List获取图片

Image的属性有很多种，主要的属性为fit，用于表示图片的填充模式，参数类型为BoxFit，BoxFit的取值主要有以下几种，示例图片来自flutter官方。

**contain** 全图显示，保持原比例。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29861abca6029?imageView2/0/w/1280/h/960/ignore-error/1) **cover** 全图充满，可能拉伸也可能被裁剪 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b29861b23cac6a?imageView2/0/w/1280/h/960/ignore-error/1) **fill** 全图显示，通过拉伸来充满目标框 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b29861ad5b8846?imageView2/0/w/1280/h/960/ignore-error/1) **fitHeight** 图片高度充满目标框，可能拉伸也可能被裁剪 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b29861add9c745?imageView2/0/w/1280/h/960/ignore-error/1) **fitWidth** 图片宽度充满目标框，可能拉伸也可能被裁剪 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b29862526145c2?imageView2/0/w/1280/h/960/ignore-error/1) **none** 保持图片的原始大小，剪裁掉位于目标框外的图片部分 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2986282e5dcde?imageView2/0/w/1280/h/960/ignore-error/1) **scaleDown** 与contain缩小图像的方式相同，只不过会在必要时缩小以确保图片完全在目标框内，如果不缩小等同于none。 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b298626599b656?imageView2/0/w/1280/h/960/ignore-error/1)

` child: Image.network( "https://upload-images.jianshu.io/upload_images/1417629-53f7d0902457cbe6.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240" , width: 260 , height: 100 , fit: BoxFit.fill, ), 复制代码`

效果如下图所示：

![VZr8uF.png](https://user-gold-cdn.xitu.io/2019/6/6/16b298627f8a8618?imageView2/0/w/1280/h/960/ignore-error/1)

#### **5.4 凸起按钮** ####

凸起按钮RaisedButton是符合Material Design设计规范的按钮，可以通过onPressed来回调按钮的点击。

` RaisedButton( onPressed: () => print( "onPressed" ), color: Colors.lightBlueAccent, child: Text( 'RaisedButton' , style: TextStyle(fontSize: 10 )), ), 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b298627a63c04f?imageView2/0/w/1280/h/960/ignore-error/1)

除了使用RaisedButton，flutter还提供了其他的按钮，比如FlatButton、IconButton、FloatingActionButton等等，它们的使用方法和RaisedButton大同小异，这里就不再赘述。

#### **5.5 其他Widget** ####

Basics Widget中的还有Row、Column、Container等Widget，这里简单介绍下。

**Row** Row用于在水平方向显示数组中的子元素Widget。

` child: Row( children: <Widget>[ Icon(Icons.access_alarm), Icon(Icons.add_a_photo), Icon(Icons.add_call), ], ), 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b298628e9959a9?imageView2/0/w/1280/h/960/ignore-error/1)

垂直方向显示数组中的子元素Widget用Column，使用方法和Row一样。这里需要提到的是Expanded，可以用Expanded来配合Row和Column使用，用来填充剩余的空间。

` Row( children: <Widget>[ Icon(Icons.access_alarm), Icon(Icons.add_a_photo), Icon(Icons.add_call), Expanded( child: FittedBox( fit: BoxFit.contain, child: const FlutterLogo () , ), flex: 2, ), Expanded ( child: Text( "占剩余部分的三分之一" , ) , flex: 1, ), ], ), 复制代码`

其中的Expanded的作用是在自己的尺寸范围内缩放并且调整child位置，使得child适合其尺寸。FlutterLogo是Basics Widget中的一种，用于展示Flutter图标。使用flex可以调整两个Expanded的占比。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2986304aba2b2?imageView2/0/w/1280/h/960/ignore-error/1)

**Container** 一个便利的容器Widget，可以设置Widget的背景、尺寸、定位。描述起来有些抽象，可以理解它和Android中的ViewGroup差不多。

` Container( decoration:BoxDecoration( color: Colors.lightGreen ), child: Text( 'Container' ), padding: EdgeInsets.all( 36.0 ), margin: EdgeInsets.all( 10.0 ), ), 复制代码`

Container的padding和margin属性和Android中的作用是类似的：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b298631f818b3d?imageView2/0/w/1280/h/960/ignore-error/1)

### **总结** ###

本文主要介绍了什么是Widget、Widget的分类、Basics Widget。因为Widget的数量繁多，官方将Widget进行了分类，并将需要先了解的Widget归入到了Basics Widget中，后续文章会介绍其他的Widge分类。

分享大前端、Android、Java等技术，助力5万程序员成长进阶。

![](https://user-gold-cdn.xitu.io/2018/8/21/1655afa15727cd03?imageView2/0/w/1280/h/960/ignore-error/1)