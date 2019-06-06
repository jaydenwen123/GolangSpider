# 【译】WWDC2019之SwiftUI #

### SwiftUI基础教程 ###

> 
> 
> 
> **SwiftUI只支持Xcode 11、iOS 13版本及以上。**
> 
> 

官方文档链接： [developer.apple.com/tutorials/s…]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Ftutorials%2Fswiftui%2Fcreating-and-combining-views )

## 创建组合视图 ##

本篇文章将通过一个构建应用（Landmarks，一个可以发现、分享你喜欢地点的App）示例，来引导大家进行SwiftUI开发。我们将使用SwiftUI框架来构建Landmark详情界面。

Landmarks利用stacks将图片和文本组合起来来进行视图布局。你需要引用MapKit框架头文件来创建一个地图视图。 **你可以通过Xcode新的实时反馈功能，来优化你的视图布局** 。

1. [下载Demo工程。]( https://link.juejin.im?target=https%3A%2F%2Fdocs-assets.developer.apple.com%2Fpublished%2F71844d6561%2FCreatingAndCombiningViews.zip )
2. [下载Xcode11 Beta。]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdownload%2F )

### 创建工程 ###

利用SwiftUI应用模版来创建工程，然后探索了解下SwiftUI的画布。

为了能够体验Xcode 11的view实时预览和交互功能，一定要确保你的mac系统版本是 [macOS 10.15 beta]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fservices-account%2Fdownload%3Fpath%3D%2FWWDC_2019%2FmacOS_10.15_Developer_Beta_Access_Utility%2FmacOSDeveloperBetaAccessUtility.dmg ) 。

#### 第一步 ####

打开 **Xcode->Create a new Xcode project，或者通过File > New > Project** 来创建工程。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c369ec84149?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第二步 ####

在模版选择区域，选择 **iOS->Single View App->Next** 。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c369ec3a08c?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第三步 ####

输入项目名称 **Landmarks->勾选Use SwiftUI->Next** 保存。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c369e0b5436?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第四步 ####

在Xcode导航栏，创建ContentView.swift。通常SwiftUI会声明两个结构体。 第一个结构体继承自View，并且在这儿进行View的布局。第二个结构体声明了一个ContentView 的preview,继承自PreviewProvider 。

**感谢@SoolyChristina基友的友情提示** 。这儿并非是继承的概念，原文的描述如下： The first structure conforms to the View protocol and describes the view’s content and layout. The second structure declares a preview for that view.

所以这儿声明的两个结构体，更像是遵循了View和PreviewProvider协议。

` import SwiftUI struct ContentView: View { var body: some View { Text( "Hello World" ) } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码` ![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c369f5015a6?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第五步 ####

在SwiftUI画布中点击 **Resume** 进行视图预览。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c369f683171?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> Tip：如果画布没有展示出来，可以通过 **Editor > Editor and Canvas** 显示出来。
> 
> 

#### 第六步 ####

把Hello World更改为Hello SwiftUI!

当你修改文案后，SwiftUI会自动更新视图。（这他么不就是热重载嘛 **Hot-Reload** ）

### 自定义Text View ###

你有两种方式来自定义TextView。第一种方式是直接修改view代码，第二种方式是通过inspector检查器来帮助你进行代码编写。

当你构建Landmarks的时候，你可以运用任何一个编辑器来进行编码工作：直接修改源代码、通过画布、通过inspector view检查器。代码并不会关心你用什么工具，它始终能够保持最新状态。

**接下来，你将通过inspector来自定义Text View**

#### 第一步 ####

在preview画布上，按住Command键+点按Text文本框，这时候inspector就会被唤起。

inspector弹出框所展示的属性也会因为不同的UI控件而有所不同。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c36f7912239?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第二步 ####

通过inspector检查器修改Text文本框的属性。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c369fbbb4c8?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第三步 ####

修改文本框字体。

> 
> 
> 
> 修改文本框字体是利用的系统的字体。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c371795eaf9?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第四步 ####

手动修改代码，即添加.color(.green) 把文本修改成绿色。
要自定义SwiftUI视图，你可以调用modifiers方法。Modifiers可以修改视图的属性，并且modifier返回一个新的视图，所以通常会将多个modifiers像链一样垂直堆叠在一起。（ **说白了就是链式编程，每调用一个方法就返回自身** ）。

` import SwiftUI struct ContentView: View { var body: some View { Text( "Turtle Rock" ) .font(.title) .color(.green) } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码`

你编写的代码肯定和view是一一对应的。当你通过inspector修改了view属性之后，Xcode会自动更新你的代码。

#### 第五步 ####

这时候，打开inspector，然后把文本Color属性修改为Inherited。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c377ae8fe73?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第六步 ####

注意一点的就是，Xcode会根据inspector修改自动更新你的代码。

### 利用Stacks组合视图 ###

我们创建了一个文本框用来显示landmark的详情信息，并且把这个文本控件放到头部。
当我们创建SwiftUI视图控件的时候，我们会把控件的内容、布局还有一些行为放在body属性中；然而body属性只返回了一个view。你可以利用stacks嵌入多个view，它可以垂直嵌入、水平嵌入等。

在这个篇幅，我们将使用垂直stack来显示park详情信息。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c37e521d06f?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第一步 ####

Command+点按text初始化方法区域。选择 **Embed in VStack** 。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c38131547f3?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第二步 ####

接下来，我们将拖拽一个text view到stack中。

点击+号，打开Library面板。拖拽一个text view到 **“Turtle Rock”后面** 。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3871272183?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第三步 ####

修改text view文案为 **Joshua Tree National Park** 。

#### 第四步 ####

设置text view的字体。

` import SwiftUI struct ContentView: View { var body: some View { VStack { Text( "Turtle Rock" ) .font(.title) Text( "Joshua Tree National Park" ) .font(.subheadline) } } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码`

#### 第五步 ####

修改VStack对齐方式。

` import SwiftUI struct ContentView: View { var body: some View { VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) Text( "Joshua Tree National Park" ) .font(.subheadline) } } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码`

如果不设置对齐方式，VStack默认是内容垂直居中。

#### 第六步 ####

在面板中，Command+点按 **Joshua Tree National Park** 唤起inspector，选择 **Embed in HStack** 。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c38bcf464c0?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第七步 ####

在location后面添加一个新的文本框，修改文本框文案并设置字体。

` import SwiftUI struct ContentView: View { var body: some View { VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) HStack { Text( "Joshua Tree National Park" ) .font(.subheadline) Text( "California" ) .font(.subheadline) } } } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码`

#### 第八步 ####

可以在两个水平的文本框之间添加Space来适应宽度。
Space把父视图在水平或者垂直方向上全部充满。

` import SwiftUI struct ContentView: View { var body: some View { VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) HStack { Text( "Joshua Tree National Park" ) .font(.subheadline) Spacer() Text( "California" ) .font(.subheadline) } } } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码` ![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3917631219?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第九步 ####

最后，利用padding()来设置边距。

` import SwiftUI struct ContentView: View { var body: some View { VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) HStack { Text( "Joshua Tree National Park" ) .font(.subheadline) Spacer() Text( "California" ) .font(.subheadline) } } .padding() } } struct ContentView_Preview: PreviewProvider { static var previews: some View { ContentView() } } 复制代码`

### 创建一个自定义的图片视图 ###

我们已经把park名称和位置的视图做好了，接下来我们将给park添加个图片。

你不需要添加很多代码，就可以添加一个带mask、border、shadow的图片。

#### 第一步 ####

添加一张图片到asset catalog中。

在Resource文件夹中找到turtlerock.png图片，然后把它拖拽到asset catalog中。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c391a25fb2d?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第二步 ####

选择 **File > New > File** 打开模版选择面板。在 **User Interface** 区域，选择 **SwiftUI View->Next** ，命名为CircleImage.swift。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3982b7ac54?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第三步 ####

把Text构建方法替换成Image。

` import SwiftUI struct CircleImage: View { var body: some View { Image( "turtlerock" ) } } struct CircleImage_Preview: PreviewProvider { static var previews: some View { CircleImage() } } 复制代码`

#### 第四步 ####

调用.clipShape(Circle())方法，创建圆形视图。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c471dae6ba6?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第五步 ####

再创建一个圆圈，用灰色进行填充。并将它作为image的border。

` import SwiftUI struct CircleImage: View { var body: some View { Image( "turtlerock" ) .clipShape(Circle()) .overlay( Circle().stroke(Color.gray, lineWidth: 4)) } } struct CircleImage_Preview: PreviewProvider { static var previews: some View { CircleImage() } } 复制代码`

#### 第六步 ####

添加阴影。

#### 第七步 ####

将边框颜色更改为白色。

` import SwiftUI struct CircleImage: View { var body: some View { Image( "turtlerock" ) .clipShape(Circle()) .overlay( Circle().stroke(Color.white, lineWidth: 4)) .shadow(radius: 10) } } struct CircleImage_Preview: PreviewProvider { static var previews: some View { CircleImage() } } 复制代码`

### UIKit和SwiftUI混合使用 ###

现在我们需要创建一个地图视图。你可以MapKit中的MKMapView类来展示渲染地图界面。

在SwiftUI中要使用UIView或者其子类，你需要让你的view遵循UIViewRepresentable协议。SwiftUI在WatchKit和AppKit同样声明了类似的协议。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3ae87b219c?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第一步 ####

创建新的SwiftUI View来展示MKMapView。 **File > New > File** ，然后创建MapView.swift。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3b46bb531b?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第二步 ####

引入MapKit头文件，并且让MapView遵循UIViewRepresentable协议。

#### 第三步 ####

UIViewRepresentable协议有两个协议方法需要实现。第一是UIView(context:)来创建MKMapView。第二个updateUIView(_:context:)来更新view。

把body属性干掉，然后UIView(context:)协议方法来创建MKMapView。

` import SwiftUI import MapKit struct MapView: UIViewRepresentable { func makeUIView(context: Context) -> MKMapView { MKMapView(frame: .zero) } } struct MapView_Preview: PreviewProvider { static var previews: some View { MapView() } } 复制代码`

#### 第四步 ####

实现updateUIView(_:context:)协议方法，来更新view（设置地图经纬度等）。

` func updateUIView(_ view: MKMapView, context: Context) { let coordinate = CLLocationCoordinate2D( latitude: 34.011286, longitude: -116.166868) let span = MKCoordinateSpan(latitudeDelta: 2.0, longitudeDelta: 2.0) let region = MKCoordinateRegion(center: coordinate, span: span) view.setRegion(region, animated: true ) } 复制代码`

#### 第五步 ####

当在静态模式下进行预览的时候，Xcode只能渲染SwiftUI视图控件。因为MKMapView是UIView子类，所以你需要把模式切换成live模式才能正常预览。

点击 **Live Preview** 切换预览模式。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c471dcd5d46?imageView2/0/w/1280/h/960/ignore-error/1)

### 把上面的子控件组合成一个完成的详情界面 ###

现在我们已经把所有子控件定义实现好了。
利用我们现有的工具，我们可以把这些子控件组合起来，形成完整的landmarks详情界面。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3cc9833d2d?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第一步 ####

在工程导航区，选择ContentView.swift文件。

#### 第二步 ####

在这三个text view控件外面，再嵌入一个VStack视图。

` struct ContentView: View { var body: some View { VStack { VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) HStack(alignment: .top) { Text( "Joshua Tree National Park" ) .font(.subheadline) Spacer() Text( "California" ) .font(.subheadline) } } .padding() } } } 复制代码`

#### 第三步 ####

将你自定义的MapView放在stack的上面。设置MapView的frame。
如果你只设置了Mapview的高度，那么MapView会自动设置其宽度来适应父视图。所以MapView会充满宽度区域。

` struct ContentView: View { var body: some View { VStack { MapView() .frame(height: 300) VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) HStack(alignment: .top) { Text( "Joshua Tree National Park" ) .font(.subheadline) Spacer() Text( "California" ) .font(.subheadline) } } .padding() } } } 复制代码` ![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3db5fb1f1a?imageView2/0/w/1280/h/960/ignore-error/1)

#### 第四步 ####

点击 **Live Preview** 来预览效果。
预览状态下，你可以继续编写view的代码，Live Preview会实时更新视图。

#### 第五步 ####

将CircleImage添加到stack上面。

` struct ContentView: View { var body: some View { VStack { MapView() .frame(height: 300) CircleImage() VStack(alignment: .leading) { Text( "Turtle Rock" ) .font(.title) HStack(alignment: .top) { Text( "Joshua Tree National Park" ) .font(.subheadline) Spacer() Text( "California" ) .font(.subheadline) } } .padding() } } } 复制代码`

#### 第六步 ####

调整一下Image的偏移。

#### 第七步 ####

在VStack的底部添加spacer占位。

#### 第八步 ####

最后设置下 edgesIgnoringSafeArea(.top) 。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20c3dd872c496?imageView2/0/w/1280/h/960/ignore-error/1)