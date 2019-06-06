# 不可思议的纯 CSS 滚动进度条效果 #

问题先行，如何使用 CSS 实现下述滚动条效果？

![scrollbar](https://user-gold-cdn.xitu.io/2019/1/9/168314f06c5dee8a?imageslim)

就是顶部黄色的滚动进度条，随着页面的滚动进度而变化长短。

在继续阅读下文之前，你可以先缓一缓。尝试思考一下上面的效果或者动手尝试一下，不借助 JS ，能否巧妙的实现上述效果。

OK，继续。这个效果是我在业务开发的过程中遇到的一个类似的小问题。其实即便让我借助 Javascript ，我的第一反应也是，感觉很麻烦啊。所以我一直在想，有没有可能只使用 CSS 完成这个效果呢？

![image](https://user-gold-cdn.xitu.io/2019/1/9/168314f06bcc0eba?imageView2/0/w/1280/h/960/ignore-error/1)

## 分析需求 ##

第一眼看到这个效果，感觉这个跟随滚动动画，仅靠 CSS 是不可能完成的，因为这里涉及了页面滚动距离的计算。

如果想只用 CSS 实现，只能另辟蹊径，使用一些讨巧的方法。

好，下面就借助一些奇技淫巧，使用 CSS 一步一步完成这个效果。分析一下难点：

* **如何得知用户当前滚动页面的距离并且通知顶部进度条？**

正常分析应该是这样的，但是这就陷入了传统的思维。进度条就只是进度条，接收页面滚动距离，改变宽度。如果页面滚动和进度条是一个整体呢？

## 实现需求 ##

不卖关子了，下面我们运用 **线性渐变** 来实现这个功能。

假设我们的页面被包裹在 ` <body>` 中，可以滚动的是整个 body，给它添加这样一个从左下到到右上角的线性渐变：

` body { background-image : linear-gradient (to right top, #ffcc00 50%, #eee 50%); background-repeat : no-repeat; } 复制代码`

那么，我们可以得到一个这样的效果：

![scrollbar2](https://user-gold-cdn.xitu.io/2019/1/9/168314f06c416417?imageslim)

Wow，黄色块的颜色变化其实已经很能表达整体的进度了。其实到这里，聪明的同学应该已经知道下面该怎么做了。

我们运用一个伪元素，把多出来的部分遮住：

` body ::after { content : "" ; position : fixed; top : 5px ; left : 0 ; bottom : 0 ; right : 0 ; background : #fff ; z-index : - 1 ; } 复制代码`

为了方便演示，我把上面白色底改成了黑色透明底，：

![scrollbar3](https://user-gold-cdn.xitu.io/2019/1/9/168314f06c66d095?imageslim)

实际效果达成了这样：

![scrollbar4](https://user-gold-cdn.xitu.io/2019/1/9/168314f06e5f9eb2?imageslim)

眼尖的同学可能会发现，这样之后，滑到底的时候，进度条并没有到底：

![image](https://user-gold-cdn.xitu.io/2019/1/9/168314f06c1757ac?imageView2/0/w/1280/h/960/ignore-error/1)

究其原因，是因为 ` body` 的线性渐变高度设置了整个 body 的大小，我们调整一下渐变的高度：

` body { background-image : linear-gradient (to right top, #ffcc00 50%, #eee 50%); background-size : 100% calc (100% - 100vh + 5px); background-repeat : no-repeat; } 复制代码`

这里使用了 ` calc` 进行了运算，减去了 ` 100vh` ，也就是减去一个屏幕的高度，这样渐变刚好在滑动到底部的时候与右上角贴合。

而 ` + 5px` 则是滚动进度条的高度，预留出 ` 5px` 的高度。再看看效果，完美：

![scrollbar](https://user-gold-cdn.xitu.io/2019/1/9/168314f06c5dee8a?imageslim)

至此，这个需求就完美实现拉，算是一个不错的小技巧，完整的 Demo：

[CodePen Demo -- 使用线性渐变实现滚动进度条]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2FChokcoco%2Fpen%2FKbBXQM%3Feditors%3D1100 )

![image](https://user-gold-cdn.xitu.io/2019/1/9/168314f15ff853c8?imageView2/0/w/1280/h/960/ignore-error/1)

别人写过的东西通常我都不会再写，这个技巧很早以前就有看到，中午在业务中刚好用到这个技巧就写下了本文，没有去考证最先发明这个技巧的是谁。不知道已经有过类似的文章，所以各位也可以看看下面这篇：

[ W3C -- 纯CSS实现Scroll Indicator(滚动指示器)]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3cplus.com%2Fcss%2Fpure-css-create-scroll-indicator.html )

## 最后 ##

其实这只是非常牛逼的 **渐变** 非常小的一个技巧。更多你想都想不到的有趣的 CSS 你可以来这里瞧瞧：

[CSS-Inspiration -- CSS灵感]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fchokcoco%2FCSS-Inspiration )

更多精彩 CSS 技术文章汇总在我的 [Github -- iCSS]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fchokcoco%2FiCSS ) ，持续更新，欢迎点个 star 订阅收藏。

好了，本文到此结束，希望对你有帮助 :)

如果还有什么疑问或者建议，可以多多交流，原创文章，文笔有限，才疏学浅，文中若有不正之处，万望告知。