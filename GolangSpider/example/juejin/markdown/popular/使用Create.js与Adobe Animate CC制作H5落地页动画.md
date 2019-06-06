# 使用Create.js与Adobe Animate CC制作H5落地页动画 #

## 一、背景 ##

在移动端，利用H5技术，可实现微杂志,微信邀请函,H5小游戏等营销互动等应用开发，本文将介绍一种Create.js与Adobe Animate CC结合来制作H5落地页的方法。

> 
> 
> 
> [CreateJS 基于H5开发的模块化库和工具,可快捷地开发基于HTML5的游戏、动画和交互应用](
> https://link.juejin.im?target=http%3A%2F%2Fwww.createjs.cc%2F )
> 
> 

> 
> 
> 
> [ADOBE ANIMATE 设计适合游戏、应用程序和 Web 的交互式矢量动画和位图动画工具](
> https://link.juejin.im?target=https%3A%2F%2Fwww.adobe.com%2Fcn%2Fproducts%2Fanimate.html%23x
> )
> 
> 

通过两者结合既充份利用了Create.js的便捷性控制逻辑和整体框架结构，又充份应用了Animate动画制作的灵活性，使设计师与前端工程师分开但可并行工作。
开始之前，先预览下效果：

![预览图](https://user-gold-cdn.xitu.io/2019/6/1/16b12585a8699a35?imageslim)

> 
> 
> 
> 注：本文所提到代码示例使用 ` es5` 进行开发。
> 
> 

## 二、页面框架 ##

#### 1.页面结构 ####

![页面结构](https://user-gold-cdn.xitu.io/2019/6/2/16b1719de2652581?imageView2/0/w/1280/h/960/ignore-error/1) 经过需求分析，从页面结构分层角度，需要将页面结构分层，上从到下，分别为LOADING层，用于显示每个动画加载百分比 ，控制层包括按钮，音乐开关，上下滑动手势控制等，场景层用于加载每个场景用到的动画，通过这种结构，主要的核心点在于，这个需求有8个动画场景，每个动画场景占用的字节比比较大，因此需要按需加载。

#### 2.目录结构 ####

![目录结构](https://user-gold-cdn.xitu.io/2019/6/2/16b171ab07c30724?imageView2/0/w/1280/h/960/ignore-error/1) 从目录结构划分来看，以复用角度出发，将可复用的内容比如CORE，控制动画加载类，交互类放于engine目录，从上到下依次是：

> 
> 
> 
> configs: 存放动画基础配置
> engine: 存放框架引擎
> images: Animate CC制作好的动画资源
> sounds: 声音资源
> 
> 
> 

#### 3.LOADING层 ####

Loading层纯展示，接收外部进度，场景开始加载显示，场景加载完成时消失。浮动于页面最上方，使用非CANVAS技术显示。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b172da50ad0ed1?imageView2/0/w/1280/h/960/ignore-error/1) 加上居中转动的CSS3动画

` body { overflow :hidden; } #_preload_div_ { position :absolute; left : 0 ; top : 0 ; width : 100% ; height : 100% ; background : rgba (0,0,0,0.5); } #animation_container , #_preload_div_.mask { position :absolute; margin :auto; display : inline-block; height : 150px ; width : 120px ; vertical-align :middle; left : 50% ; top : 50% ; transform : translate (-50%,-50%); -webkit-transform : translate (-50%,-50%); text-align : center; } #_preload_div_.img { vertical-align : middle; max-height : 100% } #_preload_div_.img img { animation : rotates 2s linear infinite; -webkit-animation : rotates 2s linear infinite; } #_preload_div_ span { display : inline-block; margin-top : - 5px ; vertical-align : middle; color : #fff ; font-size : 18px ; } #LXB_CONTAINER_SHOW { display : none !important ; } @ keyframes rotates { 0%{ transform : rotate (0); -webkit-transform : rotate (0); } 100%{ transform : rotate (360deg); -webkit-transform : rotate (360deg); } } 复制代码`

#### 4.控制层 ####

从控制层开始，动画效果交由Create.js引擎控制，需要做的是

* 

绑定容器

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1733a5ac3942c?imageView2/0/w/1280/h/960/ignore-error/1)

* 

初始化调用

![](https://user-gold-cdn.xitu.io/2019/6/2/16b17852af59d023?imageView2/0/w/1280/h/960/ignore-error/1)

* 

加载场景

![](https://user-gold-cdn.xitu.io/2019/6/2/16b17941c3934bab?imageView2/0/w/1280/h/960/ignore-error/1)

* 

开始与结束场景处理

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1792c44fd5fc5?imageView2/0/w/1280/h/960/ignore-error/1)

* 

引入各场景动画和控制逻辑

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1735044768c7f?imageView2/0/w/1280/h/960/ignore-error/1)

#### 5.场景层 ####

场景层用于控制整个H5动画交互逻辑组合，它实现了：

* 

子动画场景事件注册

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1744ea53f6bc0?imageView2/0/w/1280/h/960/ignore-error/1) 这里实现了接收每个子动画场景发送过来的动画播放完毕事件，主场景在收到该事件后可进行相关的操作，比如加载下一场景，销毁当前场景等。

* 

加载上下场景逻辑

![](https://user-gold-cdn.xitu.io/2019/6/2/16b174601bea88fb?imageView2/0/w/1280/h/960/ignore-error/1)

* 

场景资源复用、销毁

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1746cee2beab0?imageView2/0/w/1280/h/960/ignore-error/1) 复用的主要逻辑是，将加载的资源在添加到主场景的同时存一份到map中，下次复用再通过场景ID方式获取，在Create.js中通过 ` removeChild()` 移出场景不会占用CPU资源。

#### 6.手势动画支持 ####

手势动画需要引入 ` createjs.Tween` 以及对 ` pressup pressmove mousedown` 三事件综合判断。主要思想：在 ` mousedown` 时记录坐标点,判断Y轴方向在 ` pressup` 时记录的坐标点是否大于原来点，大于表示向上滑动，反之向下滑动。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b174c057288b15?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/2/16b174c89dcd79a9?imageView2/0/w/1280/h/960/ignore-error/1)

#### 7.声音管理 ####

Create.js在设计时考虑到外部资源比如音乐、视频，会存在容量较大情况，因此有提供相应的API来做异步加载。目的是减少主程序字节数，这里的声音管理主要涉及到：

* 背景音乐，循环播放需要考虑加入循环参数 ` {interrupt: createjs.Sound.INTERRUPT_ANY, loop: -1, volume: 0.5}`

* 音效，考虑复用，按需才加载，二次使用则复用原有加载内容

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1751c3a71425d?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/2/16b175239247c071?imageView2/0/w/1280/h/960/ignore-error/1)

#### 8.微信分享 ####

普通的XHR请求使用 ` new createjs.LoadQueue(true)` 要注意的是异步方式调用 ` new createjs.LoadItem().set` 在页面加载完比先要获取签名，然后再设置使用JSDK用到的方法，这方面网络有相关教程，如果想需要有服务端签名自已实现，可参考这篇文章：
[[基于egg.js微信分享API封装]]( https://juejin.im/post/5ce0c44e6fb9a07eba2c0d1c )

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1757e3ce7f040?imageView2/0/w/1280/h/960/ignore-error/1) 同时，分享出去的地址，需要有对应的 ` encodeURIComponent`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b17594e3c285cc?imageView2/0/w/1280/h/960/ignore-error/1)

## 三、动画制作 ##

前面说过，使用这种方式的话可以让设计师和前端工程师并行工作，比如，可以先定好每个动画的FLA文件的模板结构，即：

* 开始帧动作 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b1768b35241c6c?imageView2/0/w/1280/h/960/ignore-error/1)
* 结束帧动作 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b176a0e8917aa3?imageView2/0/w/1280/h/960/ignore-error/1) 剩下的交给设计师过行天马行空创作，那么成稿将有可能是这个样子： ![](https://user-gold-cdn.xitu.io/2019/6/2/16b176ae558bcb81?imageView2/0/w/1280/h/960/ignore-error/1) 像这种复杂的动画，一般CSS3比较难实现，并且效率也不是Animate等创作工具能比拟的，而以上提到的开始和结束，关键点就是该FLA动画被Create.js加载之后能与其交互，接收播开始与完成事件。

## 三、发布 ##

#### 1.作为静态资源与egg.js结合 ####

将制作好的动画需要放在一个服务器上的静态资源目录下使用，这里引用egg.js默认的配置， ` public` 目录,要引入微信分享，则需要有个服务端，借助node.js技术发展，一般小型化服务完全可以由前端一手包办。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1779cd2f7f1d7?imageView2/0/w/1280/h/960/ignore-error/1)

#### 2.可扩展配置 ####

因为引入了服务器，则会遇到有相对路径情况，因此程序中加入了前缀配置，将这些配置抽离，放置独立文件中。比如:

* 静态图片资源目录
* 音乐资源目录
* 避免缓存的版本号 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b177d6733d6bc3?imageView2/0/w/1280/h/960/ignore-error/1)
* 然后，可以这么来使用: ![](https://user-gold-cdn.xitu.io/2019/6/2/16b177e68aacf944?imageView2/0/w/1280/h/960/ignore-error/1)

## 四、小结 ##

通过以上介绍可以发现这是一种综合型工具类的组合，可能大家会有一疑问，H5动画与CSS3不是标配嘛?
我认为大多数情况下是肯定的，但在复杂的动画场景，当需求偏向于制作效率与高还原度场景的时候Animate这类的动画创作工作就派上用场了。当然它选用的是Canvas动画，我们可以截一下它的成品源码，略知一二：

* 图层动画 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b1771bfa5c1548?imageView2/0/w/1280/h/960/ignore-error/1)
* 矢量处理 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b17724601aafdd?imageView2/0/w/1280/h/960/ignore-error/1)
* 使用到的静态资源加载 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b17730bea55b9a?imageView2/0/w/1280/h/960/ignore-error/1)
* 关键的一点，使用自执行函数来接收外部传入的引擎参数，同时导出lib1,这个将对外暴露，可充份利用这一点进行控制。 ![](https://user-gold-cdn.xitu.io/2019/6/2/16b177551e6cc5d7?imageView2/0/w/1280/h/960/ignore-error/1)

本文所指的所有源码包括框架已在github上传，有兴趣的读者可自行下载研究。
[CreateJS-And-Adobe-Animate-CC]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjamesliauw%2FCreateJS-And-Adobe-Animate-CC )