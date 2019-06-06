# iOS 项目优化 #

### 前言 ###

近期正处于一段工作空白区，也想着学习学习一下项目优化，所以就自己的项目出手，一步一步地优化项目。

### 一、项目结构与应用包瘦身 ###

#### 项目结构 ####

项目本身首先划分功能区以 ` Page` 、 ` Core` 、 ` App` 划分

* ` Page` 存储应用的模块，包含 ` 首页` 、 ` 个人中心等` ，每个模块下再以 ` Controller` 、 ` View` 、 ` Model` 划分
* ` Core` 存储着一些与项目业务、界面无关的类，包括 ` 分类` 、 ` 宏定义` 、 ` 封装的请求基类` 等
* ` App` 则存储着一些与项目相关的类，包括 ` API` 、 ` Base基类` 等

#### 应用包瘦身 ####

##### 1、首先找出项目中未使用的图片： #####

这里使用了 ` python` 脚本 [脚本地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flefex%2FTCZLocalizableTool%2Fblob%2Fmaster%2FLocalToos%2FunUseImage.py ) 找出项目中未使用的图片，结果不是非常准确，但是可以自行判断 ` (存在图片使用是根据服务端返回显示的情况等)`

##### 使用方法： #####

* 下载链接的 ` unUseImage.py` 文件
* 修改项目路径以及输出路径
* 终端执行 ` python unUseImage.py` 使用后会在路径里输出文件： ![未使用图片](https://user-gold-cdn.xitu.io/2019/6/6/16b2be869168220d?imageView2/0/w/1280/h/960/ignore-error/1)

#### 后面发现 [这个工具]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftinymind%2FLSUnusedResources ) 您可以试试 ####

#### 2、图片无损压缩 ####

这里使用工具 [ImageOptim]( https://link.juejin.im?target=https%3A%2F%2Fimageoptim.com%2Fhowto.html ) ，点击链接下载 将项目图片拖入优化即可 优化结果:

![图片压缩](https://user-gold-cdn.xitu.io/2019/6/6/16b2be8683680ac4?imageView2/0/w/1280/h/960/ignore-error/1)

#### 3、代码瘦身 ####

这里推荐使用 [LinkMap]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhuanxsd%2FLinkMap ) ，可以知道项目中各个类的大小，以权衡是否有替换方案

![LinkMap](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86920250d3?imageView2/0/w/1280/h/960/ignore-error/1)

#### 4、 使用 fui 找到应用中未使用的类 ####

安装fui

` sudo gem install fui -n /usr/ local /bin 复制代码`

到项目中使用

` fui find 复制代码`

即可找出未引用的类，自行判断删除即可 #####优化前与优化后的 ` ipa` 包大小对比

![优化前](https://user-gold-cdn.xitu.io/2019/6/6/16b2be8692dd0fbe?imageView2/0/w/1280/h/960/ignore-error/1)

![优化后](https://user-gold-cdn.xitu.io/2019/6/6/16b2be8692b6935a?imageView2/0/w/1280/h/960/ignore-error/1)

### 二、项目性能优化 ###

### 内存 ###

使用 ` instruments leaks` 检测 打开：

* Product -> Profile -> Leaks
* 点击 ` CellTree` 、下方筛选
* 定位泄漏代码，修改 ![Leaks](https://user-gold-cdn.xitu.io/2019/6/6/16b2be8692daac50?imageView2/0/w/1280/h/960/ignore-error/1)

![内存泄漏](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86c4fe55e1?imageView2/0/w/1280/h/960/ignore-error/1) 工具只是辅助你寻找泄漏的地方，具体是否泄漏还需要自行判断

### 卡顿 ###

在性能优化中一个最具参考价值的属性是FPS:全称Frames Per Second,其实就是屏幕刷新率，苹果的iphone推荐的刷新率是60Hz，也就是说GPU每秒钟刷新屏幕60次，这每刷新一次就是一帧frame，FPS也就是每秒钟刷新多少帧画面。静止不变的页面FPS值是0，这个值是没有参考意义的，只有当页面在执行动画或者滑动的时候，FPS值才具有参考价值，FPS值的大小体现了页面的流畅程度高低，当低于45的时候卡顿会比较明显。

##### 这里使用 ` Core Animation` 来检测，注:需使用真机 #####

打开： 1、Product -> Profile -> Core Animation 2、启动应用 3、滑动查看FPS值

![FPS](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86c6288de2?imageView2/0/w/1280/h/960/ignore-error/1) 卡顿优化：

##### 1、Color Blended Layers (图层混合) #####

这个选项是检测哪里发生了图层混合，先介绍一下什么是图层混合？很多情况下，界面都是会出现多个UI控件叠加的情况，如果有透明或者半透明的控件，那么GPU会去计算这些这些layer最终的显示的颜色，也就是我们肉眼所看到的效果。例如一个上层Veiw颜色是绿色RGB(0,255,0)，下层又放了一个View颜色是红色RGB(0,0,255)，透明度是50%，那么最终显示到我们眼前的颜色是蓝色RGB(0,127.5,127.5)。这个计算过程会消耗一定的GPU资源损耗性能。如果我们把上层的绿色View改为不透明， 那么GPU就不用耗费资源计算，直接显示绿色。 如果出现图层混合了，打开Color Blended Layers选项，那块区域会显示红色，所以我们调试的目的就是将红色区域消减的越少越好。那么如何减少红色区域的出现呢？只要设置控件不透明即可。

* 设置opaque 属性为true。
* 给View设置一个不透明的颜色，没有特殊需要设置白色即可。

### eg: ###

运行应用 在模拟器中找到：

![Color Blended Layers](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86c64b81fe?imageView2/0/w/1280/h/960/ignore-error/1) 显示运行后的界面： ![混合图层](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86d72734ef?imageView2/0/w/1280/h/960/ignore-error/1) 出现了许多混合图层，我们就是要消除这些红色

##### TIP：当UILabel的内容是中文，需要添加一句 ` label.layer.masksToBounds = YES` ，因为当UILabel的内容为中文时，label实际渲染区域要大于label的size，最外层多了一个sublayer，如果不设置第二行label的边缘外层灰出现图层混合的红色，因此需要在label内容是中文的情况下加第二句。单独使用 ` label.layer.masksToBounds = YES` 是不会发生离屏渲染的 #####

注：xib 也可以直接设置 ` masksToBounds` 在控件的：

![xib.masksToBounds](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86f062ede3?imageView2/0/w/1280/h/960/ignore-error/1) 点击 ` +` 号添加 ` layer.masksToBounds` 打钩即可 优化后的界面： ![Color Blended Layers 优化后](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86f522aba0?imageView2/0/w/1280/h/960/ignore-error/1) 相比之前，清爽了不少，图片还需美工提供的图片为无透明的图片。

##### 2、Color Misaligned Images(图片大小) #####

这个选项可以帮助我们查看图片大小是否正确显示。如果image size和imageView size不匹配，image会出现黄色。要尽可能的减少黄色的出现，因为image size与imageView size不匹配，会消耗资源压缩图片。 选择：

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2be86fa57d9e9?imageView2/0/w/1280/h/960/ignore-error/1) 如果图片出现黄色，可自行将图片压缩至ImageView大小

##### 3、Color Offscreen-Rendered Yellow（离屏渲染） #####

离屏渲染Off-Screen Rendering 指的是GPU在当前屏幕缓冲区以外新开辟一个缓冲区进行渲染操作。还有另外一种屏幕渲染方式-当前屏幕渲染On-Screen Rendering ，指的是GPU的渲染操作是在当前用于显示的屏幕缓冲区中进行。 离屏渲染会先在屏幕外创建新缓冲区，离屏渲染结束后，再从离屏切到当前屏幕， 把离屏的渲染结果显示到当前屏幕上，这个上下文切换的过程是非常消耗性能的，实际开发中尽可能避免离屏渲染。 触发离屏渲染Offscreen rendering的行为：

* drawRect:方法
* layer.shadow
* layer.allowsGroupOpacity or layer.allowsEdgeAntialiasing
* layer.shouldRasterize
* layer.mask
* layer.masksToBounds && layer.cornerRadius

### eg: ###

运行项目，打开：

![Color Off-screen Rendered](https://user-gold-cdn.xitu.io/2019/6/6/16b2be871437c489?imageView2/0/w/1280/h/960/ignore-error/1) 项目中，此处显示黄色： ![离屏渲染](https://user-gold-cdn.xitu.io/2019/6/6/16b2be871be8e509?imageView2/0/w/1280/h/960/ignore-error/1)

在切圆角时，使用了 ` layer.masksToBounds && layer.cornerRadius` 这里我们换一种实现方式，使用贝塞尔曲线画一个边框Layer覆盖在上面即解决了离屏渲染

` UIBezierPath *maskPath = [UIBezierPath bezierPathWithRoundedRect:rect byRoundingCorners:UIRectCornerAllCorners cornerRadii:CGSizeMake(cornerRadius, cornerRadius)]; CAShapeLayer *strokeLayer = [CAShapeLayer layer]; strokeLayer.path = maskPath.CGPath; strokeLayer.fillColor = [UIColor clearColor].CGColor; //内容填充的颜色设置为clear strokeLayer.strokeColor = kWhiteColor.CGColor; //边色 strokeLayer.lineWidth = 1; // 边宽 [self.layer addSublayer:strokeLayer]; 复制代码`

网上还有许多方式可以设置圆角:

* ` UIBezierPath + Core Graphics` 切圆角
* ` UIBezierPath + Core Graphics` 覆盖镂空图片在四角 等等 处理后的效果： ![处理离屏渲染后](https://user-gold-cdn.xitu.io/2019/6/6/16b2be87271406f3?imageView2/0/w/1280/h/960/ignore-error/1)

##### 优化后的FPS #####

![优化后FPS](https://user-gold-cdn.xitu.io/2019/6/6/16b2be87359f8dcf?imageView2/0/w/1280/h/960/ignore-error/1)

### 三、使用RunTime 尽量避免 crash ###

程序运行中难免会出现崩溃，这里我们可以使用 ` runtime` 尽量避免一些常见的崩溃错误： #####eg: 给NSArray 替换 ` objectAtIndex:` 方法

` + (void)load { [NSClassFromString(@ "__NSArrayI" ) swapMethod:@selector(objectAtIndex:) currentMethod:@selector(mq_objectAtIndex:)]; } - (id)mq_objectAtIndex:(NSUInteger)index { if (index >= [self count]) { return nil; } return [self mq_objectAtIndex:index]; } + (void)swapMethod:(SEL)originMethod currentMethod:(SEL)currentMethod; { Method firstMethod = class_getInstanceMethod(self, originMethod); Method secondMethod = class_getInstanceMethod(self, currentMethod); method_exchangeImplementations(firstMethod, secondMethod); } 复制代码`

Load方法替换 ` objectAtIndex` 为 ` mq_objectAtIndex` ，当调用 ` objectAtIndex` 时会走到 ` mq_objectAtIndex` ，判断是否越界，以此来预防数组越界的crash 其他类像 ` NSDictionary、NSString` 也可以自行添加

### 结语 ###

iOS项目优化还有挺多方面的，包括电池优化、启动优化等等，笔者这里就先优化到这里，如果有需求需要优化的话，会再进行更新，谢谢支持！