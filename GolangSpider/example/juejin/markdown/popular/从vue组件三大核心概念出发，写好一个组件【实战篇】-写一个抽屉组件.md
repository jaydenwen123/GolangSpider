# 从vue组件三大核心概念出发，写好一个组件【实战篇】-写一个抽屉组件 #

## 前言 ##

上周写了一篇 [如何写好一个 vue 组件]( https://juejin.im/post/5cdacf96e51d453ae110543b ) ，算是理论篇吧，但是知易行难，这周我就自己写一个组件，来实践一下自己的文章

## 项目展示 ##

![项目展示](https://user-gold-cdn.xitu.io/2019/5/23/16ae28269c606e46?imageslim)

本文章以抽屉组件为例子，有不好或者不完善的地方，欢迎在评论区指出或者提 PR， [项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiugq5713%2Fcomponents%2Ftree%2Fmaster%2Fsrc%2Fcomponents%2FDrawer ) ， [项目 demo 地址]( https://link.juejin.im?target=https%3A%2F%2Fso-easy.cc%2Fcomponents%2F%3Fpath%3D%2Fstory%2Fdrawer--without-props )

我们想象一下用户会如何使用我们的组件，它可能需要哪些自定义的功能，比如内容的宽度，控件的位置，抽屉的位置，控件样式自定义等等，可能的交互比如：点击控件/鼠标悬浮打开抽屉，点击抽屉外部收起抽屉等等，接着我们判断一下哪些是需要暴露给外部组件的，哪些是属于组件内部的状态，尽可能的做到这个组件职责单一，且遵循最少知识原则。从这些个角度出发，来编写我们的代码

## 设计 ##

这个组件有个通用的名字，叫抽屉（Drawer），组件结构分为控件和内容两部分。如图：

` +-----------------------+ | | +-------+ | | | | | | | | | content | controls | | | | | | | | | | | +-------+ | | | +-----------------------+ 复制代码`

不以规矩，不成方圆。HTML 有语义化标签，CSS 有 BEM 规范，这些帮助我们写出结构清晰的 HTML 架构（ps:布局部分使用语义化标签还挺适合的，这种局部小组件还是 div 一把梭了）。组件 HTML 结构如下：

` < div class = "drawer-container" > < div class = "drawer" > < div class = "controls__container" ref = "controls__container" > < ul class = "controls" > < li > xxx </ li > </ ul > </ div > < div class = "content" > </ div > </ div > </ div > 复制代码`

## 基本实现 ##

我们拿贴在右侧的抽屉举例(实际代码与它不完全相同)：

### 抽屉的展开收起 ###

我们定义好抽屉的大小，并将其 postion 设置为 fixed，使用 top，right 属性，将其固定在右侧。因为抽屉默认是收起的，然后通过 translate 将其移除可视区。

`.drawer { width : '300px' ; height : '100vh' ; position : fixed; top : 0 ; right : 0 ; transform : 'translate(100%,0)' ; } 复制代码`

抽屉展开的代码也很简单，在通过 translate 将其移回来

`.drawer__container--show.drawer { transform : translate (0); } 复制代码`

### 显示控件 ###

通过负值将控件从抽屉内容区移出来

`.controls__container { position : absolute; left : - 40px ; } 复制代码`

## 基本交互 ##

### 处理抽屉的打开关闭 ###

抽屉组件支持了 mouseover 和 click 事件，开发的时候，遇到一个比较麻烦的问题：当抽屉以 mouseover 触发，鼠标移到控件上的时候，抽屉会很鬼畜的打开收起打开收起。（因为鼠标在控件上，mouseover 事件不断的被触发，导致抽屉的打开和收起）

面对这种情况，我一开始就想到了防抖和节流。但其实直接拿来用是不适合的

> 
> 
> 
> 防抖的原理：你尽管触发事件，但是我一定在事件触发 n 秒后才执行，如果你在一个事件触发的 n 秒内又触发了这个事件，那我就以新的事件的时间为准，n
> 秒后才执行，总之，就是要等你触发完事件 n 秒内不再触发事件，我才执行。
> 
> 

防抖由于是在一个事件触发 n 秒之后才执行，导致组件有一种反应慢的感觉。

> 
> 
> 
> 节流的原理：如果你持续触发事件，每隔一段时间，只执行一次事件。
> 
> 

其执行事件是异步的，那么当我打开抽屉，然后将鼠标移到抽屉外（移到抽屉外会关闭抽屉），因为抽屉的打开和关闭都是由 ` show` 变量控制。如果使用节流，会导致异步执行打开抽屉的函数，导致抽屉关闭之后又开起。

节流一般是指事件在一段时间内执行。我们这里不妨换一种思路，对 ` show` 值进行节流，你也可以把它理解成一种锁。那么当 ` show` 值变化后，我们锁住 ` show` 值，n 秒内不允许修改，n 秒后才可以修改。即控制住了抽屉不会在短时间内迅速开合。我们使用计算属性实现如下：

` // this.lock 初始值为undefine // 开闭抽屉的函数通过对lockedShow进行赋值，不会直接操作show lockedShow: { get() { return this.show; }, set(val) { if ( this.lock) { return ; } else { this.lock = setTimeout( () => { // 200毫秒之后解除锁 this.lock = undefined ; }, 200 ); this.show = val; } } } 复制代码`

### 点击抽屉外部分收起 ###

这里我们通过 [Element.closest()]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FAPI%2FElement%2Fclosest ) 方法用来获取点击的祖先元素（Element.closest：匹配特定选择器且离当前元素最近的祖先元素，也可以是当前元素本身）。如果匹配不到，则返回 null。

` closeSidebar(evt) { const parent = evt.target.closest( ".drawer" ); // 点击抽屉以外部分，即匹配不到，parent值为null if (!parent) { this.show = false ; } } 复制代码`

全局监听点击事件

` window.addEventListener( 'click' , this.closeSidebar) 复制代码`

我一开始的做法是，组件挂载的时候，全局监听点击事件，组件销毁时移除点击事件。但我们可以做的更好，当 controls 被点击时，添加点击事件，收起抽屉的时候，移除点击事件。减少全局监听的 click 事件。

除了点击事件，我们也顺便支持一下 hover 的操作。鼠标移出收起的操作和点击抽屉外部分收起的代码相同。

通过 ` e.type` 判断是点击事件还是鼠标移入事件。

` toggleDrawerShow(e) { if (e.type === "mouseover" && this.triggerEvent === "mouseover" ) { // do some thing } if (e.type === "click" && this.triggerEvent === "click" ) { // do some thing } } 复制代码`

## 优化 ##

### 控件的位置 ###

使得控件完全贴合内容区，不会因为控件的内容变化，比如控件内容为 show 和 hidden，由于切换的时候，两个单词长度不一样，而使得控件显示不完全，或者脱离内容区。

这种情况我们可以使用 JavaScript 动态计算。因为经常用到，还是封装成一个函数吧。还是拿右侧抽屉举例子:

` updateControlLayout() { // 获取控件的宽高 const rect = this.$refs[ 'controls' ].getBoundingClientRect() if ( this.position === 'right' ) { // 重新设置偏移量 this.$refs[ 'controls' ].style[ 'left' ] = `- ${rect.width} px` } } 复制代码`

### 动画 ###

主要是蒙层的显影，以及抽屉的开合。CSS动画 [贝塞尔曲线]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FCSS%2Ftiming-function ) 了解一下，笔者自己也了解不多，感兴趣可以自己去看。

` transition : opacity 0.3s cubic-bezier (0.7 , 0.3 , 0.1 , 1); 复制代码`

### 滚动条 ###

当内容过长的时候，打开抽屉的时候，滚动条还在。因此我们需要在抽屉打开的时候打开滚动条。代码也很好写，给 ` document.body` 添加 ` overflow:hidden` 属性。

这里有一个小小的坑。原先的 css 是置于 scope 里面的，如果想要把这个属性添加到 body 上，是不成功的。把 scoped 去了即可。

` <style> .hidden_scoll_bar{ overflow: hidden; } </style> 复制代码`

## 自定义 ##

### 覆写控件样式 ###

每个人都有自己独特的审美，不然也不会出现那么多的 UI 库了。作为一个组件的设计者，很难预设很多种样式让每一个使用组件的人都满意。不如把自己定义的控件作为插槽的后备内容，用户可以很方便的使用 ` control` 的具名插槽覆写控件。

` < li v-for = "(control,idx) in controlItems" class = "control" :class = "'control-'+idx" :key = "idx" > < template v-if = "show" > // 提供用户自定义插槽所需要的信息（控件是否展示，控件的信息） < slot name = "control" v-bind:drawer = "{drawerShow:show,control}" > {{control.hidden}} </ slot > </ template > < template v-else > < slot name = "control" v-bind:drawer = "{drawerShow:show,control}" > {{control.show}} </ slot > </ template > </ li > 复制代码`

### 支持设置抽屉位置 ###

因为抽屉支持在上下左右四个方向上放置，不同方向上定义的偏移方向都不同。因此需要定义不同的 css 类。通过传入的 position 值，利用 css 的级联特性应用样式

` < div class = "drawer__container" :class = "[positionClass,{'drawer__container--show':show}]" > </ div > 复制代码` ` data() { return { show : false , positionClass : this.position }; }, 复制代码` ` // 定义右侧的drawer，其余方向上的同理 // 通过css的级联，对不同方向上的drawer添加不同的样式 .right .drawer { height: 100vh; width: 100%; transform: translate(100%, 0); top: 0; right: 0; } 复制代码`

### 抽屉开启的钩子 ###

抽屉组件内部的状态没有被暴露出去，用户可能有点击控件，不打开抽屉而去做其他事情的需求。因此我们需要提供一个钩子，通过 prop 将函数 ` openDrawer` 传入， ` openDrawer` 控制是否抽屉被打开。

点击控件，开合抽屉的实现，利用了事件委托，将 click 事件，mouseover 事件直接挂载到了 ` class=controls` 的 ul 元素上，为了方便识别目标 ` li` 元素，给每一个 li 元素添加 ` :class="'control-'+idx"`

` < ul class = "controls" @ click = "toggleDrawerShow" @ mouseover = "toggleDrawerShowByMouseover" > < li v-for = "(control,idx) in controlItems" class = "control" :class = "'control-'+idx" :key = "idx" > <!-- xxx --> </ li > </ ul > 复制代码` ` // 开合抽屉的函数 openDrawerByControl(evt) { const onOpenDraw = this.openDrawer; if (!onOpenDraw) { this.lockedShow = true ; return ; } // 获取到目标阶段指向的函数 const target = evt.target; //获取到代理事件的元素 const currentTarget = evt.currentTarget; // 我们给openDraw传入target，currentTarget两个参数，具体由父组件决定onOpenDraw如何实现 this.lockedShow = onOpenDraw(target, currentTarget); } 复制代码`

父组件传入的函数如下，关于事件委托的知识感觉可以应用在这里，笔者做一个示例，让 ` class='control-0'` 的元素不能点击。

我们使用 [Element.matches]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FAPI%2FElement%2Fmatches ) 匹配 `.control-0` 类，其可以像 CSS 选择器做更加灵活的匹配。但因为 li 元素里面可能还有其他元素，所以需要不断向上寻找其父元素，直到匹配到我们事件委托的元素为止

` openDrawer(target) { let shouldOpen = true ; // 仅遍历到最外层 while (!target.matches( ".controls" )) { // 判断是否匹配到我们所需要的元素上 if (target.matches( ".control-0" )) { shouldOpen = false ; break ; } else { // 向上寻找 target = target.parentNode; } } return shouldOpen; } 复制代码`

## 总结 ##

* 用到了很多 Element 的方法（eg：closest，matches），平时很少接触
* CSS 真难写，作为一个写后台的，不经常写 CSS 的表示好难，这里费了最多的功夫
* 实践了自己之前 [写好一个组件的文章]( https://juejin.im/post/5cdacf96e51d453ae110543b ) ，知易行难，还需努力
* 一开始自己可能很难想全组件需要什么配置，可以文档先行，先想好做什么怎么做

## 参考文章 ##

* [事件委托]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F26536815 )
* [class 与数组语法]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fclass-and-style.html%23%25E6%2595%25B0%25E7%25BB%2584%25E8%25AF%25AD%25E6%25B3%2595 )
* [作用域插槽]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-slots.html%23%25E4%25BD%259C%25E7%2594%25A8%25E5%259F%259F%25E6%258F%2592%25E6%25A7%25BD )
* [JavaScript 专题之跟着 underscore 学防抖]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F22 )
* [JavaScript 专题之跟着 underscore 学节流]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F26 )
* [如何写好一个 vue 组件]( https://juejin.im/post/5cdacf96e51d453ae110543b )

## 关于我 ##

一个一年小前端，关注我的微信公众号，和我一起交流，我会尽我所能，并且看看我能成长成什么样子吧。交个朋友吧！

![](https://user-gold-cdn.xitu.io/2019/5/16/16ac0e7b1f78bede?imageView2/0/w/1280/h/960/ignore-error/1)