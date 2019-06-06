# Flutter 高效自学笔记（一） #

[Flutter 高效自学笔记（二）]( https://juejin.im/post/5ceacfb56fb9a07ee9584cf4 )

## 为什么要学 Flutter ##

学就行了，哪那么多废话。

## Get Started ##

中文： [flutterchina.club/get-started…]( https://link.juejin.im?target=https%3A%2F%2Fflutterchina.club%2Fget-started%2Fcodelab%2F ) 英文： [flutter.dev/docs/get-st…]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fget-started%2Fcodelab )

入门教程质量非常高，基本没有阻碍，而且把 VSCode 和 Android Studio 都讲了。

我本人当然是选择 Android Studio 啦，因为我的台式机性能过剩嘛 :)

性能不行的同学可以选择 VSCode。

注意：看教程时你可能会发现网络不是很通畅，那么这时你可以 [设置中国镜像]( https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fcommunity%2Fchina ) ：

` set PUB_HOSTED_URL=https://pub.flutter-io.cn set FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn 或 export PUB_HOSTED_URL=https://pub.flutter-io.cn export FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn 复制代码`

也可以使用 Privoxy + ss 给 Android Studio 设置代理。

如果不会翻墙，就不用学编程了。

可以看出中英文教程还是不太一样的。英文教程只有 4 步，把其他步骤放在了 Part 2；而中文教程直接就是 7 步，没有 Part 2。

其实直接看英文教程里的代码就好，代码周围的废话不用看得太细。

有 TypeScript 基础的前端秒上手，没有 TS 基础的前端滚去学 TS 或者学 Dart 基本语法吧。

## 了解更多 ##

学完 Get Started 教程后，官方给出了一个教程列表。（中文文档的列表内容比英文文档的少）

> 
> 
> 
> * [Building layouts in Flutter](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Fui%2Flayout
> ) tutorial
> * [Add interactivity](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Fui%2Finteractive
> ) tutorial
> * [Introduction to widgets](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fdevelopment%2Fui%2Fwidgets-intro
> )
> * [Flutter for Android developers](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fget-started%2Fflutter-for%2Fandroid-devs
> )
> * [Flutter for React Native developers](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fget-started%2Fflutter-for%2Freact-native-devs
> )
> * [Flutter for web developers](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fget-started%2Fflutter-for%2Fweb-devs
> )
> * [Flutter's Youtube channel](
> https://link.juejin.im?target=https%3A%2F%2Fwww.youtube.com%2Fflutterdev )
> 
> * [Build Native Mobile Apps with Flutter](
> https://link.juejin.im?target=https%3A%2F%2Fwww.udacity.com%2Fcourse%2Fbuild-native-mobile-apps-with-flutter--ud905
> ) (a free Udacity course)
> * [Flutter cookbook](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fcookbook%2F
> )
> * [From Java to Dart](
> https://link.juejin.im?target=https%3A%2F%2Fcodelabs.developers.google.com%2Fcodelabs%2Ffrom-java-to-dart%2F%230
> ) codelab
> * [Bootstrap into Dart: learn more about the language](
> https://link.juejin.im?target=https%3A%2F%2Fflutter.dev%2Fdocs%2Fresources%2Fbootstrap-into-dart
> )
> 
> 
> 

让我来看看哪些有用。

* 第一个教程教我如何使用 Layout，我进去大概扫了一下，就是各个布局控件的用法，先放着，用到再学。
* 第二个教程教我如何使用 setState 改变各种组件的状态，比 Get Started 稍微深入一些，先放着，用到再学。
* 第三个教程教我一个组件的完整使用，需要重点看。
* 第四个教程是给安卓开发者看的，将安卓开发者熟悉的概念映射到 Flutter 中，不适合我看。
* 第五个教程是给 React Native 开发者看的，不适合我看。
* 第六个教程是给 Web 开发者看的，适合我看。进去看一下 * 讲了如何设置样式
* 讲了如何定位、变形、加渐变、加圆角、加阴影
* 讲了如何操作文本
* 写教程的人是不是对 Web 开发者有什么误解，我想看的不是这些
* 也许我应该看给 React Native 开发者看的教程

* 第七个教程是官方的 Youtube 频道，有 30 多个英文视频，果断关掉。
* 第八个教程是 Udacity 的免费课程，点进去看看 * 视频是英文对白，中文字幕
* 教程分两部分，第一部分可以通过看第一个教程和第二个教程搞定，第二部分可以通过看第六个教程搞定
* 我不打算看这个 8 小时的免费课程了。

* 第九个教程是 Cookbook（菜谱），它讲核心概念按首字母顺序排列了出来。显然这只能当字典使用，不适合阅读。
* 第十个教程是教 Dart 语法的，对 Dart 不熟可以先看这个。不过它默认你学过 Java。

* 里面介绍了一个叫做 [DartPad]( https://link.juejin.im?target=https%3A%2F%2Fdartpad.dartlang.org%2F ) 的工具，可以实时练习 Dart，好用。

* 第十一个教程是一个完整的 Dart 教程的索引，内容相当多。用到再看。

扫完这些链接之后，我打算先看第三个教程，然后自己试着做一个 Demo，遇到 Flutter 问题就看上面对应的教程。

未完待续……