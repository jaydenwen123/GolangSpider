# iOS13-适配夜间模式/深色外观(Dark Mode) #

> 
> 
> 
> 今天的 WWDC 19 上发布了 iOS 13，我们来看下如何适配 DarkMode。
> 
> 

首先我们来看下效果图

![效果图.gif](https://user-gold-cdn.xitu.io/2019/6/4/16b218a0a37d05b0?imageslim)

### DarkMode 主要从两个方面去适配，颜色和图片 ###

#### 颜色适配 ####

iOS 13 下 ` UIColor` 增加了一个初始化方法，我们可以用这个初始化方法完成 DarkMode 的适配

` @available(iOS 13.0, *) public init(dynamicProvider: @escaping (UITraitCollection) -> UIColor) 复制代码`

这个方法要求传一个闭包进去，当系统从 LightMode 和 DarkMode 之间切换的时候就会触发这个回调。

这个闭包返回一个 ` UITraitCollection` 类，我们要用这个类的 ` userInterfaceStyle` 属性。

` userInterfaceStyle` 是一个枚举，声明如下

` @available (iOS 12.0 , *) public enum UIUserInterfaceStyle : Int { case unspecified case light case dark } 复制代码`

这个枚举会告诉我们当前是 LightMode or DarkMode

现在我们创建两个 ` UIColor` 并赋值给 ` view.backgroundColor` 和 ` label` ，代码如下

` let backgroundColor = UIColor { (traitCollection) -> UIColor in switch traitCollection.userInterfaceStyle { case.light: return UIColor.white case.dark: return UIColor.black default : fatalError () } } view.backgroundColor = backgroundColor let labelColor = UIColor { (traitCollection) -> UIColor in switch traitCollection.userInterfaceStyle { case.light: return UIColor.black case.dark: return UIColor.white default : fatalError () } } label.textColor = labelColor 复制代码`

现在，我们做完了背景色和文本颜色的适配，接下来我们看看图片如何适配

#### 图片适配 ####

打开 ` Assets.xcassets` 把图片拖拽进去，我们可以看到这样的页面

![1.jpg](https://user-gold-cdn.xitu.io/2019/6/4/16b218ebd84f46e7?imageView2/0/w/1280/h/960/ignore-error/1)

然后我们在右侧工具栏中点击最后一栏，点击 ` Appearances` 选择 ` Any, Dark` ，如图所示

![2.jpg](https://user-gold-cdn.xitu.io/2019/6/4/16b218f5e8d942fb?imageView2/0/w/1280/h/960/ignore-error/1)

我们把 DarkMode 的图片拖进去，如图所示

![3.jpg](https://user-gold-cdn.xitu.io/2019/6/4/16b218fa6c41c90d?imageView2/0/w/1280/h/960/ignore-error/1)

最后我们加上 ` ImageView` 的代码

` imageView.image = UIImage (named: "icon" ) 复制代码`

**现在我们就已经完成颜色和图片的 DarkMode 适配，是不是很简单呢 (手动滑稽)**

### 如何获取当前样式 (Light or Dark) ###

我们可以看到，不管是颜色还是图片，适配都是系统完成的，我们不用关心现在是什么样的样式。

但是在某些场景下，我们可能会有根据当前样式来做一些其他适配的需求，这时我们就需要知道现在什么样式。

我们可以在 ` UIViewController` 或 ` UIView` 中调用 ` traitCollection.userInterfaceStyle` 来获取 **当前** 样式

代码如下

` switch traitCollection.userInterfaceStyle { case.unspecified: print ( "unspecified" ) case.light: print ( "light" ) case.dark: print ( "dark" ) } 复制代码`

为什么要强调 **当前** 呢，因为默认情况下使用 ` traitCollection.userInterfaceStyle` 属性就能获取到当前系统的样式。

但是我们可以通过 ` overrideUserInterfaceStyle` 属性强制设置 ` UIViewController` 或 ` UIView` 的样式

代码如下

` overrideUserInterfaceStyle = .dark print (traitCollection.userInterfaceStyle) // dark 复制代码`

我们可以看到设置了 ` overrideUserInterfaceStyle` 之后， ` traitCollection.userInterfaceStyle` 就是设置后的样式了。

**注意：** ` overrideUserInterfaceStyle` 默认值为 ` unspecified` ，所以一定要用 ` traitCollection.userInterfaceStyle` 来判断当前样式，而不是用 ` overrideUserInterfaceStyle` 来判断。

**注意：** 以上代码是我自己摸索出来的，在真机上也能达到效果，但是不建议现在就开始做 DarkMode 的适配。毕竟官方关于 DarkMode 适配的 session 还没出，建议等 session 出了之后在做适配，另外如果有和官方有出入我会及时补充修改~