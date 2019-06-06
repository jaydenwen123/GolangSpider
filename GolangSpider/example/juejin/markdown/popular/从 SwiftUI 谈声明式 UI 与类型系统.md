# 从 SwiftUI 谈声明式 UI 与类型系统 #

# Hello, SwiftUI #

Apple 在 WWDC19 上正式发布了 **Project Catalyst** ( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fipad-apps-for-mac%2F ) （原 Marzipan），使得开发者能够将 iPadOS app 移植到 macOS 上。同时 **SwiftUI** ( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fxcode%2Fswiftui%2F ) 也压轴亮相，正式统一了 Apple 全平台的 UI 开发解决方案。恰逢前些时候，Google 在其 I/O 大会上亮相了 Jetpack Compose —— 一个全新的 Android 原生 UI 开发框架，标志着两大移动操作系统阵营全面拥抱声明式 UI 开发模式。

# 声明式 UI 的前世今生 #

其实声明式 UI 并不是什么新技术，早在 2006 年，微软就已经发布了其新一代界面开发框架 **WPF** ，其采用了 XAML 标记语言，支持双向数据绑定、可复用模板等特性。

2010 年，由诺基亚领导的 Qt 团队也正式发布了其下一代界面解决方案 **Qt Quick** ，同样也是声明式，甚至 Qt Quick 起初的名字就是 Qt Declarative。QML 语言同样支持数据绑定、模块化等特性，此外还支持内置 JavaScript，开发者只用 QML 就可以开发出简单的带交互的原型应用。

声明式 UI 框架近年来飞速发展，并且被 Web 开发带向高潮。 **React** 更是为声明式 UI 奠定了坚实基础并一直引领其未来的发展。随后 **Flutter** 的发布也将声明式 UI 的思想成功带到移动端开发领域...

# 声明式到底是什么 #

想象我们要实现下面这个界面：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ba22c59ee9fa?imageView2/0/w/1280/h/960/ignore-error/1)

打开开关就让下面的 label 显示 on，反之显示 off。如果我们要用非声明式的方式实现，即命令式，那么需要：

* 创建一个 ` UISwitch` ，设置它的 change 事件 handler
* 创建一个 ` UILabel`
* 创建一个 ` UIStackView` ，设置方向为垂直
* 将 1、2 创建的两个视图添加到 ` UIStackView` 中
* change 事件触发时读取开关的当前状态，设置相应字符串到 label 中

这样做面对一个状态，我们尚且能够正确处理，但随着应用日渐复杂，状态也越来越多并且错综复杂，状态变化的顺序甚至也能影响应用逻辑的正确性，因为我们对每个事件的处理都是对界面的增量修改。一旦前一个状态有错误，后面就会错上加错，接下来多线程混入，然后 boom，你的应用可能就 crash 了。

声明式的意思就是让我们描述我们需要一个什么样的界面，而不是告诉计算机一步一步干什么。那么上面的例子用声明式就是这样：

> 
> 
> 
> “我需要一个界面，它是一个 VStack（垂直布局），里面有一个开关，开关的值与 switchValue 的布尔值绑定，VStack 里接下来是一个
> Text，它的值当 switchValue 为 true 时是 foo，否则是 bar”
> 
> 

我们可以发现，全文没有命令，都是在描述界面是怎样的。 ` switchValue` 我们称之为 “The Source of Truth”，Toggle 的状态、Text 的文本内容都与它相绑定。状态变化时，界面按照先前描述的重新“渲染”即可得到状态绝对正确的界面。这正是声明式的优势所在， **降低状态增加时界面维护的复杂度** 。

# SwiftUI 与其他框架的异同 #

SwiftUI 自亮相以来，全网就在讨论其与 React、Flutter 之间的关系云云。经过这两天的研究，我想简单谈谈我的观点： （免责声明：没有看过源码，也没有参与现场 Lab，一切都是个人想法）

首先是与 Flutter 的对比，Flutter 的思路是从 0 开始，即语言、基础库、渲染引擎、排版引擎即框架本身全部由自己实现，其渲染引擎 Skia 只需要操作系统为止提供一个 GL Context 便可以完成所有图形渲染，这使得其跨平台性变得十分强大，到目前为止 Windows、Linux、macOS、Fuchsia 都已经得到了 Flutter 官方的支持。

这种做法我认为有利有弊，首先好处是所有平台下行为一致，不管是滚动视图、Material Design 控件还是模糊效果这些在其他平台没有的都得到了全平台的支持，开发者并不需要为这些去做平台间的适配，反观 React Native... 当然缺点也是存在的，Flutter 这种做法类似于游戏引擎，平台提供的 UI 特性它一概不用，因此 Flutter View 与原生视图的交互就没有那么容易了，同时新的 Dart 语言貌似也不是非常受社区和开发者喜爱。

SwiftUI 没有像 Flutter 那样从头再来，这个全新的框架依旧使用了 UIKit、AppKit 等作为基础。但它并不是一个 UIKit 的声明式封装，通过 Xcode 的调试视图可以看出这一点：

![Xcode View Debugger](https://user-gold-cdn.xitu.io/2019/6/6/16b2b945b863f2a7?imageView2/0/w/1280/h/960/ignore-error/1)

许多基础组件，像 Text、Button 等都并不是直接使用 ` UILabel` 、 ` UIButton` 而是一个名为 ` DisplayList.ViewUpdater.Platform.CGDrawingView` 的 ` UIView` 子类。它们使用了自定义绘制，但又承载于 UIKit 的环境中，因此我猜测 SwiftUI 只提供了组件的自定义渲染和布局引擎，它使用到的底层技术还是 Core Animation、Core Graphics、Core Text 等。使用自定义绘制去实现组件可以理解成为跨平台提供便利，毕竟一个按钮还要区分 ` UIButton` 、 ` NSButton` 来实现未免有些麻烦。但是部分复杂的控件还是采用了 UIKit 中已有的类，比如 ` UISwitch` 等。由于未脱离 UIKit 体系，嵌入一个 UIView 非常容易，你不需要搞什么外部纹理（Flutter 需要），因为它们的上下文是同一个，坐标系也是同一个。

所以我认为 SwiftUI 更加类似 React Native，使用系统框架提供的组件，只不过绘制和布局可以自己来实现，这在 SwiftUI 之前也有相关的框架这样实践的，比如 **Yoga** 、 **ComponentKit** 等。

# SwiftUI 的类型系统 #

Flutter、React 的类型系统并不是强约束，一个界面里有一个 Text 和有两个 Text 类型是一样的，React 使用 JavaScript 更是无类型。SwiftUI 与它们不同，它使用了强类型约束。举个例子：

` VStack { Text ( "Hello" ) } 复制代码`

与

` VStack { Text ( "Hello" ) Text ( "World" ) } 复制代码`

与

` VStack { Text ( "Hello" ) .color( Color.red) } 复制代码`

类型都是不同的。首先上面这种语法叫做 **Function Builders** ，是 Apple “私自”夹带到 Swift 里的私货。上面这些表达式最后都会得到一个实现了 ` View` 协议的 **具体类型** ，SwiftUI 里基本使用的都是 **具体类型** ，而不是 **协议类型** ，首先 ` VStack` 是一个 struct 同时也是一个具体类型，它的构造方法里接受一个闭包，这个闭包使用了通过 ` @functionBuilder` 修饰的 ` ViewBuilder` 结构体作为 builder，因此上面的第二段代码在编译时会被转化成：

` VStack { let v1 = ViewBuilder.buildExpression( Text ( "Hello" )) let v2 = ViewBuilder.buildExpression( Text ( "World" )) return ViewBuilder.buildBlock(v1, v2) } 复制代码`

然后我们看一下上面这个 ` ViewBuilder.buildBlock` 重载的签名：

` static func buildBlock<C0, C1>(_ c0: C0, _ c1: C1) -> TupleView<(C0, C1)> where C0 : View, C1 : View 复制代码`

所以一个 Text 和两个 Text，它们的父容器 VStack 的类型都是不同的！另外提一下， ` buildBlock` 的范型参数最多有 10 个：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bc2dc598431a?imageView2/0/w/1280/h/960/ignore-error/1)

**划重点：也就是你的一个视图层级（目前）不能有超过 10 个子视图。** 且超过后编译器的错误提示丝毫不会体现这一点，了解这个将会非常节约你的时间！

不同的状态对应的视图也不同，但是它们的类型是相同的，这意味着什么呢？那就是， **不需要 Diff-Patch 了** 。

我们想象下面的场景：

` VStack { if something { Text ( "something is true" ) } Text ( "something else" ) if !something { Text ( "something is not true" ) } } 复制代码`

当 ` something` 变化时，视图应该怎么变化？对于 React、Flutter 来说，它们没有类型的概念，每次只能拿到两个快照（一个当前状态的，一个新状态的）。它们有两个选择去完成界面的更新：

* 把老的视图全部移除，重新添加新视图
* 找出它们的差异，根据差异去修改视图

第一种方法最简单，但是性能很差，且不能保存视图自身的状态。第二种方法需要高效的算法加持，看起来能解决我们的问题，但是 **它不是必要的** 。

SwiftUI 的做法是根据类型来更新界面，上面这段代码的类型是：

` VStack < TupleView < Text ?, Text , Text ?>> 复制代码`

有了类型框架就能做静态优化，这类似前端框架 Svelte 和 Vue.js 3.0 所做的一些优化，可以称之为 AOT。

在没有类型的情况下，每次状态变化，界面中都只有两个 Text，只不过内容不一样，这时候框架通过 diff 认为界面中的 Text 控件本身没变，只是内容变了，于是给它们设置了新的内容。

但事实并不是这样， ` something` 变化时，界面显示的 Text 是不同的，中间的 Text 始终显示 “something else”，变化的是它上下两个相邻的 Text。框架拿到新视图时就可以按范型参数的顺序去检查他们的差异：

` Before update: VStack(TupleView(Text(...), Text(...), nil)) After update: VStack(TupleView(nil, Text(...), Text(...))) 复制代码`

它们的相对位置写在了类型中，这样就能避免中间的视图被修改，没有类型信息或其他元信息，这点是绝对做不到的。

SwiftUI 对于类型做得其实更多，所有的字体调整、位置调整等操作在 SwiftUI 中都是通过 ` ViewModifier` 实现的，调整后的视图类型为 ` View.Modified<some ViewModifier>` ，因此有无这些参数调整的视图，类型也是不同的，这些都将有助于框架去做一些静态优化。

关于 SwiftUI 的详细使用方面，我之后可能还会再更新文章，本文就是简单谈谈我对框架宏观层面的理解。祝大家 WWDC 周玩得开心～

References:

* [forums.swift.org/t/pitch-fun…]( https://link.juejin.im?target=https%3A%2F%2Fforums.swift.org%2Ft%2Fpitch-function-builders%2F25167 )
* [github.com/apple/swift…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fapple%2Fswift-evolution%2Fblob%2F9992cf3c11c2d5e0ea20bee98657d93902d5b174%2Fproposals%2FXXXX-function-builders.md )
* [github.com/apple/swift…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fapple%2Fswift-evolution%2Fblob%2Fmaster%2Fproposals%2F0258-property-delegates.md )
* [github.com/apple/swift…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fapple%2Fswift-evolution%2Fblob%2Fmaster%2Fproposals%2F0244-opaque-result-types.md )
* [forums.swift.org/t/important…]( https://link.juejin.im?target=https%3A%2F%2Fforums.swift.org%2Ft%2Fimportant-evolution-discussion-of-the-new-dsl-feature-behind-swiftui%2F25168 )
* [developer.apple.com/wwdc19/402]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fwwdc19%2F402 )
* [svelte.dev/]( https://link.juejin.im?target=https%3A%2F%2Fsvelte.dev%2F )
* [twitter.com/unixzii/sta…]( https://link.juejin.im?target=https%3A%2F%2Ftwitter.com%2Funixzii%2Fstatus%2F1136330564582092800%3Fs%3D20 )