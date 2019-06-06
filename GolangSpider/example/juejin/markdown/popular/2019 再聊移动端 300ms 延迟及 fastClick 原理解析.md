# 2019 再聊移动端 300ms 延迟及 fastClick 原理解析 #

## 前言 ##

最近公司新开了一条业务线，有幸和大佬们一起从头开始构建一套适合新业务的框架。俗话说得好呀，适合自己的才是最好的 😎。在新项目的 CodeReview 的时候，被大哥提到有没有添加 fastClick 解决移动端 300ms 延迟的问题。以下就带你追溯移动端延迟的 ` 前世` ` 今生` 。

## 介绍 ##

### 前世 - 诞生的因 ###

国外有一篇关于 300ms 延迟的文章： [What Exactly Is..... The 300ms Click Delay]( https://link.juejin.im?target=https%3A%2F%2Fwww.telerik.com%2Fblogs%2Fwhat-exactly-is.....-the-300ms-click-delay )

世间万物皆有因果，网页兴起于桌面端，那时候有谁会想到手机等移动设备的风靡？犹记得上大学那会儿，手机访问学校网站的时候都是通过手指缩放来控制的 🙃，心里真的是一万头草泥马奔腾而过，后来为了解决移动端适配的问题，提出了 ` viewport` 的解决方案，基于 [无障碍(accessibility)]( https://link.juejin.im?target=https%3A%2F%2Fsupport.google.com%2Faccessibility%2Fandroid%2Fanswer%2F6006949%3Fhl%3Dzh-Hans ) （需要代理）交互设计师为了更好的用户体验，特地提供了 双击缩放 的手势支持。殊不知这正是一切祸乱的根源。

### 今生 - 消逝的果 ###

谷歌有开发者文档： [300ms tap delay, gone away]( https://link.juejin.im?target=https%3A%2F%2Fdevelopers.google.com%2Fweb%2Fupdates%2F2013%2F12%2F300ms-tap-delay-gone-away ) （需要代理）

以下是原文的部分引用

> 
> 
> 
> For many years, mobile browsers applied a 300-350ms delay between touchend
> and click while they waited to see if this was going to be a double-tap or
> not, since double-tap was a gesture to zoom into text.
> 
> 

大致是说，移动浏览器 会在 ` touchend` 和 ` click` 事件之间，等待 300 - 350 ms，判断用户是否会进行双击手势用以缩放文字。

> 
> 
> 
> Ever since the first release of Chrome for Android, this delay was removed
> if pinch-zoom was also disabled. However, pinch zoom is an important
> accessibility feature. As of Chrome 32 (back in 2014) this delay is gone
> for mobile-optimized sites, without removing pinch-zooming! Firefox and
> IE/Edge did the same shortly afterwards, and in March 2016 a similar fix
> landed in iOS 9.3.
> 
> 

从上面我们可以获取到几个非常重要的信息：首先，谷歌就开始吹啦，自打我们移动版 Chrome 发布以来，只要你把缩放禁用掉，这个延迟就不会出现。不得不吹一波 Google，真的是甩开 Apple 几条街，fastClick 源码大部分都是用来解决 iOS 各个版本各种奇奇怪怪的 BUG。说实话，有些源码我也不是很理解，但是咱啥也不敢说，啥也不敢问啊 😂。其次，Chrome 32 对移动端进行了优化，可以不禁用缩放，也能解决延迟的问题。接着 Firefox 和 IE/Edge 紧随其后也修复了这个 BUG，最后，就是 iOS 9.3 也同样修复这个 BUG （亲测的确修复了）。

## 解决方案 ##

以下可以通过 hack 技巧，不添加 fastClick 也能修复延迟的问题

### 禁用缩放 ###

* Chrome on Android (all versions)
* iOS 9.3

` < meta name = "viewport" content = "user-scalable=no" /> 复制代码`

或者

` html { touch-action : manipulation; } 复制代码`

* IE on Windows Phone

` html { touch-action: manipulation; // IE11+ -ms-touch-action: manipulation; // IE10 } 复制代码`

### 不禁用缩放 ###

* Chrome 32+ on Android
* iOS 9.3

` < meta name = "viewport" content = "width=device-width" /> 复制代码`

经测试，如果不添加 ` width=device-width` 不管是 Android 还是 iOS 在已修复的版本中仍然会出现延时的问题。

## WebView ##

上面说了这么多，都是针对移动端浏览器的，既然是提到移动端，WebView 当然不得不说啦。

### Android WebView ###

[如何设计一个优雅健壮的 Android WebView？]( https://juejin.im/post/5a94f9d15188257a63113a74 )

Android WebView 中 300ms 的延迟问题和移动端浏览器解决思路一致。

### iOS WebView ###

[UIWebView]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdocumentation%2Fuikit%2Fuiwebview )

> 
> 
> 
> In apps that run in iOS 8 and later, use the WKWebView class instead of
> using UIWebView. Additionally, consider setting the WKPreferences property
> javaScriptEnabled to false if you render files that are not supposed to
> run JavaScript.
> 
> 

iOS WebView 就有点让人头疼了。因为 iOS 8 之前一直都是 UIWebView，iOS 8 出了个新秀 WKWebView，那么 iOS 9.3 300ms 延迟的 BUG 修复到底干了啥呢？在客户端 iOS 小姐姐的帮助下，最终的测试结果是 UIWebView 300ms 延迟的问题到现在一直存在，哪怕是最新的 iOS 版本（这大概这就是为什么老外推荐使用 WKWebView 而非 UIWebView，估计是不想修 BUG 了吧 😂），但是 WKWebView 在 iOS 9.3 的时候将这个问题给修复了。也就是说 iOS 9.3 之前 WKWebView 仍然是存在 300ms 延迟的问题的（忙活了半天，总算把所有的都给理清楚了 🙄）。

## FastClick 原理解析 ##

这部分可能有点烂大街了，网上一搜一大把，再说也没啥意思，我就挑点个人觉得有意思的说一下 😁。

### 原理 ###

首先，讲一下 fastClick 的实现原理吧，MDN 上 [同时支持触屏事件和鼠标事件]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FAPI%2FTouch_events%2FSupporting_both_TouchEvent_and_MouseEvent ) 也有提到。

移动端，当用户点击屏幕时，会依次触发 ` touchstart` ， ` touchmove` (0 次或多次)， ` touchend` ， ` mousemove` ， ` mousedown` ， ` mouseup` ， ` click` 。 ` touchmove` 。只有当手指在屏幕发生移动的时候才会触发 ` touchmove` 事件。在 ` touchstart` ， ` touchmove` 或者 ` touchend` 事件中的任意一个调用 ` event.preventDefault` ， ` mouse` 事件 以及 ` click` 事件将不会触发。

fastClick 在 ` touchend` 阶段 调用 ` event.preventDefault` ，然后通过 ` document.createEvent` 创建一个 ` MouseEvents` ，然后 通过 ` event​Target​.dispatch​Event` 触发对应目标元素上绑定的 ` click` 事件。

### 你不知道的 ` JavaScript` (Maybe) ###

首先，我们需要明确一个问题，300ms 的延迟只有在移动端才会出现，PC 端是没有的。fastClick 中又有个一 ` notNeeded` 的函数是用来判断有没有必要使用 fastClick。刚开始的时候，刚开始我阅读完代码表示对没有进行移动端和 PC 端的区分表示不满。不过后来一段不起眼的代码改变了我的看法。

` // Devices that don't support touch don't need FastClick if ( typeof window.ontouchstart === 'undefined' ) { return true ; } 复制代码`

PC 端是没有 ` touch` 事件的因此 ` window.ontouchstart` 返回 ` undefined` ，移动端如果没有绑定事件则返回 ` null` 。果然只能证明我还是太过年轻 🤣。

阅读源码期间，无意中发现使用事件委托时，Safari 手机版会有一个 bug，当点击事件不是绑定在交互式的元素上（比如说 HTML 的 div），并且也没有直接的事件监听器绑定在他们自身。不会触发 ` click` 事件。具体可以参考 [click 浏览器兼容性]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FAPI%2FElement%2Fclick_event%23%25E6%25B5%258F%25E8%25A7%2588%25E5%2599%25A8%25E5%2585%25BC%25E5%25AE%25B9%25E6%2580%25A7 )

解决方法如下：(请原谅我厚颜无耻地直接搬过来了 😝)

* 为其元素或者祖先元素，添加 ` cursor: pointer` 的样式，使元素具有交互式点击
* 为需要交互式点击的元素添加 ` onclick="void(0)"` 的属性，但并不包括 body 元素
* 使用可点击元素如 ` <a>` ，代替不可交互式元素如 div
* 不使用 click 的事件委托。

` event.stopPropagation` 只会阻止相同类型(event.type 相同)事件传播，上面有提到过 移动端 触摸事件触发的顺序问题，假如 我在 ` touchstart` 中调用了 ` event.stopPropagation` 只会 阻止后续 event flow 上其他 ` touchstart` 事件，并不会阻止 ` touchmove` ， ` touchend` 等 mouseEvent 事件的发生。

` event.stopPropagation` ， ` event​.stop​Immediate​Propagation` 的区别你真的知道吗 🧐， ` event.stopPropagation` 阻止捕获和冒泡阶段中当前事件的进一步传播。如果有多个相同类型事件的事件监听函数绑定到同一个元素，当该类型的事件触发时，它们会按照被添加的顺序执行。如果其中某个监听函数执行 ` event.stopImmediatePropagation` 方法，则当前元素剩下的监听函数将不会被执行。

` event​Target​.dispatch​Event` 仍然会触发完整的 event flow，而不仅仅触发 event​Target​ 本身注册的事件。

总的来说，阅读源码的过程是一次自我修炼的过程，是对过去某些不足的完善，其实就是发现自己很菜 😂，前方道险且长，同志们仍需努力呀 🤜。

### 不足 ###

个人觉得阅读优秀的源码是一件很幸福的事，因为它能潜移默化的提升你的 ` 审美` 能力。但同时我们也要带有挑剔的眼光，找出当中存在的不足之处

fastClick 中 ` notNeeded` 函数总的来说，已经相当不错了，但是美中不足的是，对于 iOS 9.3 以上 使用 WKWebView 的用户来说，引入 fastClick 无疑是多此一举，还有可能导致某些潜在的问题。对于处女座的我来说，这一点是不能忍受的。不过单纯的通过 ` UA` 是无法区分 ` UIWebView` 和 ` WKWebView` 的。不过如果页面是在自己 App 中话，可以通过在 ` UA` 中携带 ` WebView` 的信息来决定是否加载

## 延迟检测 code ##

` <!DOCTYPE html> < html lang = "en" > < head > < meta charset = "UTF-8" /> < meta name = "viewport" content = "width=device-width, initial-scale=1.0" /> < meta http-equiv = "X-UA-Compatible" content = "ie=edge" /> < title > Document </ title > < style > </ style > </ head > < body > < div > < label for = "userAgent" > userAgent: </ label > < span id = "userAgent" > </ span > </ div > < div > < label for = "touchstart" > touchstart: </ label > < span id = "touchstart" > </ span > </ div > < div > < label for = "touchend" > touchend: </ label > < span id = "touchend" > </ span > </ div > < div > < label for = "click" > click: </ label > < span id = "click" > </ span > </ div > < div > < label for = "diffClickTouchend" > diff click - touchend: </ label > < span id = "diffClickTouchend" > </ span > </ div > < div > < div id = "test" > test </ div > < div id = "diff" > diff </ div > </ div > < script > var userAgent = document.getElementById( 'userAgent' ); userAgent.innerText = window.navigator.userAgent; var test = document.getElementById( 'test' ); var diff = document.getElementById( 'diff' ); var touchstart = document.getElementById( 'touchstart' ); var touchend = document.getElementById( 'touchend' ); var click = document.getElementById( 'click' ); var diffClickTouchend = document.getElementById( 'diffClickTouchend' ); test.addEventListener( 'touchstart' , function ( e ) { touchstart.innerText = Date.now(); }); test.addEventListener( 'touchend' , function ( e ) { touchend.innerText = Date.now(); }); test.addEventListener( 'click' , function ( e ) { click.innerText = Date.now(); }); diff.addEventListener( 'click' , function ( ) { diffClickTouchend.innerText = click.innerText - touchend.innerText; }); </ script > </ body > </ html > 复制代码`

## 结束语 ##

回眸历史，不可否认 fastClick 在解决移动端 300ms 延迟的问题上的确作出杰出的贡献，不过 9102 的今天，是否仍然有必要使用呢，回到开始，我说过，适合自己的才是最好的，因此，如果你的业务需求，是只需要对 iOS 9.3 以上的 WKWebView 做适配，那么强烈建议你不去使用，毕竟减少了文件请求大小，引入风险的概率。

最后，引用一句名言 ` 老兵不死，只是凋零` 向 fastClick 致敬。

## 参考 (References) ##

* [What Exactly Is..... The 300ms Click Delay]( https://link.juejin.im?target=https%3A%2F%2Fwww.telerik.com%2Fblogs%2Fwhat-exactly-is.....-the-300ms-click-delay )
* [300ms tap delay, gone away]( https://link.juejin.im?target=https%3A%2F%2Fdevelopers.google.com%2Fweb%2Fupdates%2F2013%2F12%2F300ms-tap-delay-gone-away )
* [如何设计一个优雅健壮的 Android WebView？]( https://juejin.im/post/5a94f9d15188257a63113a74 )
* [UIWebView]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.apple.com%2Fdocumentation%2Fuikit%2Fuiwebview )
* [同时支持触屏事件和鼠标事件]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FAPI%2FTouch_events%2FSupporting_both_TouchEvent_and_MouseEvent )
* [click 浏览器兼容性]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FAPI%2FElement%2Fclick_event%23%25E6%25B5%258F%25E8%25A7%2588%25E5%2599%25A8%25E5%2585%25BC%25E5%25AE%25B9%25E6%2580%25A7 )

> 
> 
> 
> 文 / [lastSeries]( https://juejin.im/user/5987d2b6f265da3e292a63d4 )
> 
> 
> 
> 作者也在掘金哦，快关注他吧！
> 
> 
> 
> 编 / [荧声]( https://juejin.im/user/56d15d581ea493005c05f71a )
> 
> 

本文由创宇前端作者授权发布，版权属于作者，创宇前端出品。 欢迎注明出处转载本文。文章链接： [juejin.im/post/5cdf84…]( https://juejin.im/post/5cdf84e3f265da1bb564c79a )

![](https://user-gold-cdn.xitu.io/2018/9/29/16623743b60a2b71?imageView2/0/w/1280/h/960/ignore-error/1)

**本文是创宇前端相关账号最后一次由荧声负责。终于，我们一起走到了一个故事的结束。**

**[都知欢聚最难得，难奈别离多]( https://link.juejin.im?target=https%3A%2F%2Fy.qq.com%2Fn%2Fyqq%2Fsong%2F003eXdik1VJHED.html ) 。**

感谢您的阅读，以及长期以来的支持。

再见啦。