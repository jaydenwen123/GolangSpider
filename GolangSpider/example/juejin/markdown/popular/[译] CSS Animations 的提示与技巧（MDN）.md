# [译] CSS Animations 的提示与技巧（MDN） #

始发于我的博客 [ryougifujino.com]( https://link.juejin.im?target=https%3A%2F%2Fryougifujino.com%2F ) ，欢迎访问留言。

CSS Animations 的出现让你在文档和应用上创建不可思议的动画效果成为了可能。但是，有一些你想实现的东西可能并非如此浅显，或者说你不能轻易的想出一种聪明的方式来完成它们。这篇文章包含了一系列的提示与技巧来使得你的工作更加容易，其中包括了如何让一个已经停止的动画重新运行起来。

## 让动画再次运行起来 ##

CSS Animations 的规范没有提供让动画再次运行起来的方法。并不存在适用于元素的魔法方法 ` requestAnimation()` ，甚至即使你把元素的 ` animation-play-state` 设置为 ` "running"` 也无济于事。你得用一种聪明的技巧来使得已经停止的动画再次回放。

我们将介绍一种我们认为足够稳定可靠的方法给你。

### HTML 内容 ###

首先，让我们创建一个用于执行动画的 ` <div>` 和一个用于播放（或者说回放）动画的按钮。

` < div class = "box" > </ div > < div class = "runButton" > Click me to run the animation </ div > 复制代码`

### CSS 内容 ###

现在让我们用CSS定义动画本身。一些并不重要的CSS（例如播放按钮自身的样式）出于简洁的考虑并不会在这里出现。

` @ keyframes colorchange { 0% { background : yellow } 100% { background : blue } }.box { width : 100px ; height : 100px ; border : 1px solid black; }.changing { animation : colorchange 2s ; } 复制代码`

这里有两个 class。 ` "box"` class 是盒子外观的基础描述，并不包含任何动画信息。动画细节都被包含在 ` "changing"` class 里，它描述了：一个叫 ` colorchange` 的 ` @keyframes` 将会被运用在一个2秒钟的作用于盒子的动画过程之中。

### JavaScript 内容 ###

下一步我们将看看 JavaScript 所做的工作。这项技术的核心位于 ` play()` 函数之中，这个函数将在用户点击“Run”按钮时被回调。

` function play ( ) { document.querySelector( ".box" ).className = "box" ; window.requestAnimationFrame( function ( time ) { window.requestAnimationFrame( function ( time ) { document.querySelector( ".box" ).className = "box changing" ; }); }); } 复制代码`

看起来真的很奇怪，不是吗？这是因为再次播放动画的唯一方法就是删除动画效果，让文档重新计算样式来使得它明白你已经进行了删除，然后再把动画效果添加回元素中。为了让这一切发生，我们必须要有创造性。

这是当 ` play()` 方法被调用时发生了什么：

* 盒子的 ` class` 被重置为了 ` "box"` 。这个动作把其他当前运用在盒子上的 class 都给删除了，包括处理动画的 ` "changing"` class。换而言之，我们正把盒子上的动画效果给移除掉。但是，在样式重新计算完成之前，或反应改变的刷新发生之前，class 列表的变更都不会生效。
* 为了确保样式已经被重新计算，我们使用 ` window.requestAnimationFrame()` 来指定一个回调。这样我们的回调就会在文档下一次重绘之前执行。问题在于，因为是发生重绘前的，所以样式的重计算还没有真的发生！所以……
* 我们的回调聪明地第二次调用了 ` requestAnimationFrame()` ！这一次，回调运行于第二次下一次重绘之前，这时第一次重绘已经完成，所以可以确保样式的重计算已经发生。这个回调把 ` changing` class 添加回了盒子之上，使得重绘后动画将再次开始。

当然，我们还需要给我们的"Run"按钮添加一个事件处理器：

` document.querySelector( ".runButton" ).addEventListener( "click" , play, false ); 复制代码`

## 停止动画 ##

仅仅删除应用于元素的 ` animation-name` 就会让元素跳转到它的下一个状态。如果你想在动画完成时再让它停止，那么你可能需要尝试另一种不同的方法。主要技巧如下：

* 让你的动画越独立越好。意思是你不应该依赖于 ` animation-direction: alternate` 。取而代之，你应该显式地写一个 keyframe 动画来模拟这个过程。
* 在 ` animationiteration` 事件发生时，用 JavaScript 清理正在被使用的动画。

下面的 demo 展示了如何实现上述的 JavaScript 技术：

`.slidein { animation-duration : 5s ; animation-name : slidein; animation-iteration-count : infinite; }.stopped { animation-name : none; } @ keyframes slidein { 0% { margin-left : 0% ; } 50% { margin-left : 50% ; } 100% { margin-left : 0% ; } } 复制代码` ` < h1 id = "watchme" > Click me to stop </ h1 > 复制代码` ` let watchme = document.getElementById( 'watchme' ) watchme.className = 'slidein' const listener = ( e ) => { watchme.className = 'slidein stopped' } watchme.addEventListener( 'click' , () => watchme.addEventListener( 'animationiteration' , listener, false ) ) 复制代码`

[Demo地址]( https://link.juejin.im?target=https%3A%2F%2Fjsfiddle.net%2Fmorenoh149%2F5ty5a4oy%2F )

[原文地址（MDN）]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FCSS%2FCSS_Animations%2FTips ) 更新于Mar 23, 2019, 6:23:51 PM。部分live demo不能链接过来，请前往原文进行查看。