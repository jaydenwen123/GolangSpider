# 翻译：测量vue应用运行时的性能！ #

## 前言 ##

为了提高英文水平，尝试着翻译一些英文技术文章，首先就从这个Vue的小技巧文章开始，目前英文版一共22篇。计划用时2~3个月翻译完成。

> 
> 
> 
> 目前进度[2/22]
> 
> 

## 原文 ##

[Measure runtime performance in Vue.js apps]( https://link.juejin.im?target=https%3A%2F%2Fvuedose.tips%2Ftips%2Fmeasure-runtime-performance-in-vue-js-apps%2F )

## 译文 ##

在 [上一篇]( https://juejin.im/post/5ceddc6e6fb9a07ed657b4cf ) 文章中，我们讨论了如何提高大型数据的性能。但是我们还没有测量它提高了多少。

我们可以使用Chrome DevTools 的性能选项来实现这一点。但是为了获取准确数据，我们必须在Vue上激活性能模式。

我们可以在 ` main.js` 或者插件中设置全局变量，代码如下：

` Vue.config.performance = true ; 复制代码`

如果你设置了正确的 NODE_ENV 环境变量，那么可以使用非生产环境做判断。

` const isDev = process.env.NODE_ENV !== "production" ; Vue.config.performance = isDev; 复制代码`

这将在Vue内部激活标记组件性能的 [User Timing API
]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FAPI%2FUser_Timing_API ) 。

上一篇文章内容，我已经在 [codesandbox]( https://link.juejin.im?target=https%3A%2F%2F0ql846q66w.codesandbox.io%2F ) 上创建了代码。打开 Chrome DevTools 里的 performance 选项并且点击重新加载按钮。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b204335288bdb9?imageView2/0/w/1280/h/960/ignore-error/1)

这将记录页面加载性能。同时，感谢你在 ` main.js` 中的 ` Vue.config.performance` 设置，这个设置会使你在统计资料能够看到 ` User Timing` 部分。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20552b4cb68a8?imageView2/0/w/1280/h/960/ignore-error/1)

在哪里，你会发现3个指标：

* **Init** ：创建组件实例需要的时间
* **Render** ：创建VDom结构需要的时间
* **Patch** ：把VDom应用到实际Dom的时间

回到上一篇文章好奇（性能提高了多少）的地方，结果是：正常的组件需要417毫秒初始化：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2067596133a76?imageView2/0/w/1280/h/960/ignore-error/1)

而使用 ` Object.freeze` 阻止了默认反应则只需要3.9毫秒：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20fc59bec4df6?imageView2/0/w/1280/h/960/ignore-error/1)

当然，每次运行的结果都会有小的变化，但是，仍然有非常巨大的性能差别。由于在创建组件的时候会有默认反应的问题，你可以通过 ` Init` （初始化指标）看到阻止了默认反应和没有阻止的差异。

就是这样！

你可以在线阅读文章 [tip online]( https://link.juejin.im?target=https%3A%2F%2Fvuedose.tips%2Ftips%2Fimprove-performance-on-large-lists-in-vue-js ) （可以 复制/粘贴 代码），但是请你记住，如果你喜欢，要和所有同事分享 [VueDose]( https://link.juejin.im?target=https%3A%2F%2Fvuedose.tips ) 。

下周见。

## 我的理解 ##

vue项目，我们可以通过在全局main.js设置 ` Vue.config.performance` 为 ` true` 来开启性能检测，可以通过环境变量来区分是否需要开启，然后就可以通过Chrome DevTools里的 performance 选项去看统计的性能数据。

## 结尾 ##

水平有限，难免有错漏之处，望各位大大轻喷的同时能够指出，跪谢！

## 其它翻译 ##

1、 [翻译：提高vue.js中大型数据的性能]( https://juejin.im/post/5ceddc6e6fb9a07ed657b4cf )
2、 翻译：测量vue应用运行时的性能！