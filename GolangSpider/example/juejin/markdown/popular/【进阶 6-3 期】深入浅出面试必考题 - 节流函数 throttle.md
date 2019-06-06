# 【进阶 6-3 期】深入浅出面试必考题 - 节流函数 throttle #

## 引言 ##

上一节我们详细聊了聊高阶函数之柯里化，通过介绍其定义和三种柯里化应用，并在最后实现了一个通用的 currying 函数。这一小节会继续之前的篇幅聊聊函数节流 throttle，给出这种高阶函数的定义、实现原理以及在 underscore 中的实现，欢迎大家拍砖。

有什么想法或者意见都可以在评论区留言，下图是本文的思维导图，高清思维导图和更多文章请看我的 [Github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog ) 。

![66666333](https://user-gold-cdn.xitu.io/2019/5/29/16b0100a03610be6?imageView2/0/w/1280/h/960/ignore-error/1)

## 定义及解读 ##

函数节流指的是某个函数在一定时间间隔内（例如 3 秒）只执行一次，在这 3 秒内 **无视后来产生的函数调用请求** ，也不会延长时间间隔。3 秒间隔结束后第一次遇到新的函数调用会触发执行，然后在这新的 3 秒内依旧无视后来产生的函数调用请求，以此类推。

![img](https://user-gold-cdn.xitu.io/2019/5/29/16b0100a0ae8835c?imageView2/0/w/1280/h/960/ignore-error/1)

举一个小例子，不知道大家小时候有没有养过小金鱼啥的，养金鱼肯定少不了接水，刚开始接水时管道中水流很大，水到半满时开始拧紧水龙头，减少水流的速度变成 3 秒一滴，通过滴水给小金鱼增加氧气。

此时「管道中的水」就是我们频繁操作事件而不断涌入的回调任务，它需要接受「水龙头」安排；「水龙头」就是节流阀，控制水的流速，过滤无效的回调任务；「滴水」就是每隔一段时间执行一次函数，「3 秒」就是间隔时间，它是「水龙头」决定「滴水」的依据。

如果你还无法理解，看下面这张图就清晰多了，另外点击 [这个页面]( https://link.juejin.im?target=http%3A%2F%2Fdemo.nimius.net%2Fdebounce_throttle%2F ) 查看节流和防抖的可视化比较。其中 Regular 是不做任何处理的情况，throttle 是函数节流之后的结果，debounce 是函数防抖之后的结果（下一小节介绍）。

![image-20190525193539745](https://user-gold-cdn.xitu.io/2019/5/29/16b0100a0366ad23?imageView2/0/w/1280/h/960/ignore-error/1)

## 原理及实现 ##

函数节流非常适用于函数被频繁调用的场景，例如：window.onresize() 事件、mousemove 事件、上传进度等情况。使用 throttle API 很简单，那应该如何实现 throttle 这个函数呢？

实现方案有以下两种

* 

第一种是用时间戳来判断是否已到执行时间，记录上次执行的时间戳，然后每次触发事件执行回调，回调中判断当前时间戳距离上次执行时间戳的间隔是否已经达到时间差（Xms） ，如果是则执行，并更新上次执行的时间戳，如此循环。

* 

第二种方法是使用定时器，比如当 scroll 事件刚触发时，打印一个 hello world ，然后设置个 1000ms 的定时器，此后每次触发 scroll 事件触发回调，如果已经存在定时器，则回调不执行方法，直到定时器触发，handler 被清除，然后重新设置定时器。

这里我们采用第一种方案来实现，通过闭包保存一个 previous 变量，每次触发 throttle 函数时判断当前时间和 previous 的时间差，如果这段时间差小于等待时间，那就忽略本次事件触发。如果大于等待时间就把 previous 设置为当前时间并执行函数 fn。

我们来一步步实现，首先实现用闭包保存 previous 变量。

` const throttle = ( fn, wait ) => { // 上一次执行该函数的时间 let previous = 0 return function (...args ) { console.log(previous) ... } } 复制代码`

执行 throttle 函数后会返回一个新的 function，我们命名为 betterFn。

` const betterFn = function (...args ) { console.log(previous) ... } 复制代码`

betterFn 函数中可以获取到 previous 变量值也可以修改，在回调监听或事件触发时就会执行 betterFn，即 ` betterFn()` ，所以在这个新函数内判断当前时间和 previous 的时间差即可。

` const betterFn = function (...args ) { let now = + new Date (); if (now - previous > wait) { previous = now // 执行 fn 函数 fn.apply( this , args) } } 复制代码`

结合上面两段代码就实现了节流函数，所以完整的实现如下。

` // fn 是需要执行的函数 // wait 是时间间隔 const throttle = ( fn, wait = 50 ) => { // 上一次执行 fn 的时间 let previous = 0 // 将 throttle 处理结果当作函数返回 return function (...args ) { // 获取当前时间，转换成时间戳，单位毫秒 let now = + new Date () // 将当前时间和上一次执行函数的时间进行对比 // 大于等待时间就把 previous 设置为当前时间并执行函数 fn if (now - previous > wait) { previous = now fn.apply( this , args) } } } // DEMO // 执行 throttle 函数返回新函数 const betterFn = throttle( () => console.log( 'fn 函数执行了' ), 1000 ) // 每 10 毫秒执行一次 betterFn 函数，但是只有时间差大于 1000 时才会执行 fn setInterval(betterFn, 10 ) 复制代码`

## underscore 源码解读 ##

上述代码实现了一个简单的节流函数，不过 underscore 实现了更高级的功能，即新增了两个功能

* 配置是否需要响应事件刚开始的那次回调（ leading 参数，false 时忽略）
* 配置是否需要响应事件结束后的那次回调（ trailing 参数，false 时忽略）

配置 { leading: false } 时，事件刚开始的那次回调不执行；配置 { trailing: false } 时，事件结束后的那次回调不执行，不过需要注意的是，这两者不能同时配置。

所以在 underscore 中的节流函数有 3 种调用方式，默认的（有头有尾），设置 { leading: false } 的，以及设置 { trailing: false } 的。上面说过实现 throttle 的方案有 2 种，一种是通过时间戳判断，另一种是通过定时器创建和销毁来控制。

第一种方案实现这 3 种调用方式存在一个问题，即事件停止触发时无法响应回调，所以 { trailing: true } 时无法生效。

第二种方案来实现也存在一个问题，因为定时器是延迟执行的，所以事件停止触发时必然会响应回调，所以 { trailing: false } 时无法生效。

underscore 采用的方案是两种方案搭配使用来实现这个功能。

` const throttle = function ( func, wait, options ) { var timeout, context, args, result; // 上一次执行回调的时间戳 var previous = 0 ; // 无传入参数时，初始化 options 为空对象 if (!options) options = {}; var later = function ( ) { // 当设置 { leading: false } 时 // 每次触发回调函数后设置 previous 为 0 // 不然为当前时间 previous = options.leading === false ? 0 : _.now(); // 防止内存泄漏，置为 null 便于后面根据 !timeout 设置新的 timeout timeout = null ; // 执行函数 result = func.apply(context, args); if (!timeout) context = args = null ; }; // 每次触发事件回调都执行这个函数 // 函数内判断是否执行 func // func 才是我们业务层代码想要执行的函数 var throttled = function ( ) { // 记录当前时间 var now = _.now(); // 第一次执行时（此时 previous 为 0，之后为上一次时间戳） // 并且设置了 { leading: false }（表示第一次回调不执行） // 此时设置 previous 为当前值，表示刚执行过，本次就不执行了 if (!previous && options.leading === false ) previous = now; // 距离下次触发 func 还需要等待的时间 var remaining = wait - (now - previous); context = this ; args = arguments ; // 要么是到了间隔时间了，随即触发方法（remaining <= 0） // 要么是没有传入 {leading: false}，且第一次触发回调，即立即触发 // 此时 previous 为 0，wait - (now - previous) 也满足 <= 0 // 之后便会把 previous 值迅速置为 now if (remaining <= 0 || remaining > wait) { if (timeout) { clearTimeout(timeout); // clearTimeout(timeout) 并不会把 timeout 设为 null // 手动设置，便于后续判断 timeout = null ; } // 设置 previous 为当前时间 previous = now; // 执行 func 函数 result = func.apply(context, args); if (!timeout) context = args = null ; } else if (!timeout && options.trailing !== false ) { // 最后一次需要触发的情况 // 如果已经存在一个定时器，则不会进入该 if 分支 // 如果 {trailing: false}，即最后一次不需要触发了，也不会进入这个分支 // 间隔 remaining milliseconds 后触发 later 方法 timeout = setTimeout(later, remaining); } return result; }; // 手动取消 throttled.cancel = function ( ) { clearTimeout(timeout); previous = 0 ; timeout = context = args = null ; }; // 执行 _.throttle 返回 throttled 函数 return throttled; }; 复制代码`

## 小结 ##

* 

函数节流指的是某个函数在一定时间间隔内（例如 3 秒）只执行一次，在这 3 秒内 **无视后来产生的函数调用请求**

* 

节流可以理解为养金鱼时拧紧水龙头放水，3 秒一滴

* 「管道中的水」就是我们频繁操作事件而不断涌入的回调任务，它需要接受「水龙头」安排
* 「水龙头」就是节流阀，控制水的流速，过滤无效的回调任务
* 「滴水」就是每隔一段时间执行一次函数
* 「3 秒」就是间隔时间，它是「水龙头」决定「滴水」的依据

* 

节流实现方案有 2 种

* 第一种是用时间戳来判断是否已到执行时间，记录上次执行的时间戳，然后每次触发事件执行回调，回调中判断当前时间戳距离上次执行时间戳的间隔是否已经达到时间差（Xms） ，如果是则执行，并更新上次执行的时间戳，如此循环。
* 第二种方法是使用定时器，比如当 scroll 事件刚触发时，打印一个 hello world ，然后设置个 1000ms 的定时器，此后每次触发 scroll 事件触发回调，如果已经存在定时器，则回调不执行方法，直到定时器触发，handler 被清除，然后重新设置定时器。

## 参考 ##

> 
> 
> 
> [underscore.js](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjashkenas%2Funderscore%2Fblob%2Fmaster%2Funderscore.js
> )
> 
> 
> 
> [前端性能优化原理与实践](
> https://juejin.im/book/5b936540f265da0a9624b04b?referrer=56dea4aa7664bf00559f002d
> )
> 
> 
> 
> [underscore 函数节流的实现](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flessfish%2Funderscore-analysis%2Fissues%2F22
> )
> 
> 

## 文章穿梭机 ##

* [【进阶 6-2 期】深入高阶函数应用之柯里化]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog%2Fissues%2F37 )
* [【进阶 6-1 期】JavaScript 高阶函数浅析]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog%2Fissues%2F36 )
* [【进阶 5-3 期】深入探究 Function & Object 鸡蛋问题]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog%2Fissues%2F35 )
* [【进阶 5-2 期】图解原型链及其继承优缺点]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog%2Fissues%2F34 )
* [【进阶 5-1 期】重新认识构造函数、原型和原型链]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog%2Fissues%2F32 )

### ❤️ 看完三件事 ###

如果你觉得这篇内容对你挺有启发，我想邀请你帮我三个小忙：

* **点赞** ，让更多的人也能看到这篇内容（ **收藏不点赞，都是耍流氓 -_-** ）
* 关注我的 **GitHub** ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyygmind%2Fblog ) ，让我们成为长期关系
* **关注公众号「高级前端进阶」** ，每周重点攻克一个前端面试重难点，公众号后台回复「资料」 送你精选前端优质资料。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0100b54fae1ab?imageView2/0/w/1280/h/960/ignore-error/1)