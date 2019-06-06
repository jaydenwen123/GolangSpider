# Flutter全平台！迁移现有Flutter项目到WEB端 #

### 写在前面 ###

Flutter 是 Google推出并开源的移动应用开发框架，主打跨平台、高保真、高性能。开发者可以通过 Dart语言开发 App，一套代码同时运行在 iOS 、Android、web和桌面端。

> 
> 
> 
> Flutter官网： [flutter-io.cn](
> https://link.juejin.im?target=https%3A%2F%2Fflutter-io.cn%2F )
> 
> 

[Flutter_web]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fflutter%2Fflutter_web ) 是Flutter代码兼容web的实现，可以将使用Dart编写的现有Flutter代码编译成可以嵌入浏览器并部署到任何Web服务器的客户端。

> 
> 
> 
> Our goal is to enable building applications for mobile and web
> simultaneously from a single codebase. However, to allow experimentation,
> the tech preview Flutter for web is developed in a separate namespace. So,
> as of today an existing mobile Flutter application will not run on the web
> without changes.
> 
> 
> 
> Flutter的目标是通过单一代码库同时构建移动和Web应用程序。 但是，为了进行实验，Flutter_web是在一个单独的命名空间中开发的。
> 因此，截至目前，现有的移动Flutter应用程序无法在不进行更改的情况下在Web上运行。
> 
> 

简而言之就是Flutter现在还不支持既是移动应用也是Web应用，需要自己进行迁移，但相信日子不会太远。

### 迁移Flutter项目到WEB端 ###

这次我们迁移的项目是 [flutter_challenge_googlemaps]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fflutter-ui-challenges%2Fflutter_challenge_googlemaps ) ，效果图如下：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b17aed14917d07?imageView2/0/w/1280/h/960/ignore-error/1)

#### 怎么做？ ####

大多数Dart代码都是共用的，需要改变的只是一些依赖和配置。

首先是 ` pubspec.yaml` 需要用flutter_web来替换flutter，同时移除asset和字体相关的代码。

` name: flutter_web_challenge_googlemaps environment: # You must be using Flutter >= 1.5.0 or Dart >= 2.3.0 sdk: '>=2.3.0-dev.0.1 <3.0.0' dependencies: flutter_web: any flutter_web_ui: any dev_dependencies: build_runner: ^ 1.4.0 build_web_compilers: ^ 2.0.0 dependency_overrides: flutter_web: git: url: https: //github.com/flutter/flutter_web path: packages/flutter_web flutter_web_ui: git: url: https: //github.com/flutter/flutter_web path: packages/flutter_web_ui flutter_web_test: git: url: https: //github.com/flutter/flutter_web path: packages/flutter_web_test 复制代码`

通过 ` flutter package get` 更新依赖后，需要更新 ` lib` 路径下dart文件中的相关引用。

` //flutter import 'package:flutter/material.dart' ; //flutter web import 'package:flutter_web/material.dart' ; 复制代码`

差别就是将 ` flutter` 替换为 ` flutter_web` 而已，代码基本不用动。

接下来，为了预览网页，我们需要自己创建web目录，并在目录下创建 ` web/index.html` 和 ` web/main.dart` 文件。

**web/index.html**

` <!DOCTYPE html> < html lang = "en" > < head > < meta charset = "UTF-8" > < title > </ title > < script defer src = "main.dart.js" type = "application/javascript" > </ script > </ head > < body > </ body > </ html > 复制代码`

**web/main.dart**

` import 'package:flutter_web_ui/ui.dart' as ui; // 将flutter_web_challenge_googlemaps替换为自己的package import 'package:flutter_web_challenge_googlemaps/main.dart' as app; main() async { await ui.webOnlyInitializePlatform(); app.main(); } 复制代码`

至于资源文件、图片、字体等，和Flutter项目不同，这些都需要放到 ` web\assets` 目录路径下，同时要记得更新代码中的相关引用。

` Image.asset( "assets/logo.ong" ); // 需要更改为 Image.asset( "logo.png" ); 复制代码`

如果你有使用Material Icon的话，你需要在 ` web/assets` 目录下创建 ` FontManifest.json` 文件，并添加相关地址。

` [ { "family" : "MaterialIcons" , "fonts" : [ { "asset" : "https://fonts.gstatic.com/s/materialicons/v42/flUhRq6tzZclQEJ-Vdg-IuiaDsNcIhQ8tQ.woff2" } ] } ] 复制代码`

整个web目录会如下图所示:

![web](https://user-gold-cdn.xitu.io/2019/6/2/16b17aec8c3a6462?imageView2/0/w/1280/h/960/ignore-error/1)

运行项目，可以发现和移动端基本没有区别。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1792a24a53cf9?imageView2/0/w/1280/h/960/ignore-error/1)

效果还是蛮流畅的🤙。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b17aed4fb5917a?imageslim)

### 写在最后 ###

虽然说跨平台的项目很多的，比如weex、RN、Kotlin等等，但是真正让我体会到跨平台高效一体的体验还是Flutter，这也许就是为什么年后我一直在学习和从事Flutter开发的原因之一了。

当然flutter_web还处于早期阶段，一些flutter的功能还没有完全移植过来，比如高斯模糊效果，而且我在safari浏览器里打开是一片空白，应该还没支持safari浏览器，不过Flutter1.0正式版本才到来不久，相信在不久的将来，这些全都会有。

最后附上相关地址： [github.com/flutter-ui-…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fflutter-ui-challenges%2Fflutter_web_challenge_googlemaps ) ，本文是为了方便查看所以新开了一个仓库，实际上只需要新开一个web分支就可以了。

#### 参考文档 ####

flutter_web： [github.com/flutter/flu…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fflutter%2Fflutter_web )

迁移指南： [github.com/flutter/flu…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fflutter%2Fflutter_web%2Fblob%2Fmaster%2Fdocs%2Fmigration_guide.md )

==================== 分割线 ======================

如果你想了解更多关于MVVM、Flutter、响应式编程方面的知识，欢迎关注我。

**你可以在以下地方找到我：**

简书： [www.jianshu.com/u/117f1cf0c…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fu%2F117f1cf0c556 )

掘金： [juejin.im/user/582d60…]( https://link.juejin.im?target=https%3A%2F%2Flinks.jianshu.com%2Fgo%3Fto%3Dhttps%253A%252F%252Fjuejin.im%252Fuser%252F582d601d2e958a0069bbe687 )

Github: [github.com/ditclear]( https://link.juejin.im?target=https%3A%2F%2Flinks.jianshu.com%2Fgo%3Fto%3Dhttps%253A%252F%252Fgithub.com%252Fditclear )

![](https://user-gold-cdn.xitu.io/2019/5/27/16af8e6368738804?imageView2/0/w/1280/h/960/ignore-error/1)