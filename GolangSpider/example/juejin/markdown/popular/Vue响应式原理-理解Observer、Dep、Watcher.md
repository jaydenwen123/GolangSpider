# Vue响应式原理-理解Observer、Dep、Watcher #

## 开篇 ##

最近在学习Vue的源码，看了网上一些大神的博客，看起来感觉还是蛮吃力的。自己记录一下学习的理解，希望能够达到简单易懂，不看源码也能理解的效果😆

如果有错误，恳求大佬们指点嘿😋

## Object.defineProperty ##

相信很多同学或多或少都了解Vue的响应式原理是通过 ` Object.defineProperty` 实现的。被 ` Object.defineProperty` 绑定过的对象，会变成「 **响应式** 」化。也就是改变这个对象的时候会触发get和set事件。进而触发一些视图更新。举个栗子🌰

` function defineReactive ( obj, key, val ) { Object.defineProperty(obj, key, { enumerable : true , configurable : true , get : () => { console.log( '我被读了，我要不要做点什么好?' ); return val; }, set : newVal => { if (val === newVal) { return ; } val = newVal; console.log( "数据被改变了，我要把新的值渲染到页面上去!" ); } }) } let data = { text : 'hello world' , }; // 对data上的text属性进行绑定 defineReactive(data, 'text' , data.text); console.log(data.text); // 控制台输出 <我被读了，我要不要做点什么好?> data.text = 'hello Vue' ; // 控制台输出 <hello Vue && 数据被改变了，我要把新的值渲染到页面上去!> 复制代码`

## Observer 「响应式」 ##

` Vue` 中用 ` Observer` 类来管理上述响应式化 ` Object.defineProperty` 的过程。我们可以用如下代码来描述，将 ` this.data` 也就是我们在 ` Vue` 代码中定义的 ` data` 属性全部进行「响应式」绑定。

` class Observer { constructor () { // 响应式绑定数据通过方法 observe( this.data); } } export function observe ( data ) { const keys = Object.keys(data); for ( let i = 0 ; i < keys.length; i++) { // 将data中我们定义的每个属性进行响应式绑定 defineReactive(obj, keys[i]); } } 复制代码`

## Dep 「依赖管理」 ##

### 什么是依赖？ ###

相信没有看过源码或者刚接触 ` Dep` 这个词的同学都会比较懵。那 ` Dep` 究竟是用来做什么的呢？ 我们通过 ` defineReactive` 方法将 ` data` 中的数据进行响应式后，虽然可以监听到数据的变化了，那我们怎么处理通知视图就更新呢？

` Dep` 就是帮我们收集【究竟要通知到哪里的】。比如下面的代码案例，我们发现，虽然 ` data` 中有 ` text` 和 ` message` 属性，但是只有 ` message` 被渲染到页面上，至于 ` text` 无论怎么变化都影响不到视图的展示，因此我们仅仅对 ` message` 进行收集即可，可以避免一些无用的工作。

那这个时候 ` message` 的 ` Dep` 就收集到了一个依赖，这个依赖就是用来管理 ` data` 中 ` message` 变化的。

` < div > < p > {{message}} </ p > </ div > 复制代码` ` data: { text : 'hello world' , message : 'hello vue' , } 复制代码`

当使用 ` watch` 属性时，也就是开发者自定义的监听某个data中属性的变化。比如监听 ` message` 的变化， ` message` 变化时我们就要通知到 ` watch` 这个钩子，让它去执行回调函数。

这个时候 ` message` 的 ` Dep` 就收集到了两个依赖，第二个依赖就是用来管理 ` watch` 中 ` message` 变化的。

` watch: { message : function ( val, oldVal ) { console.log( 'new: %s, old: %s' , val, oldVal) }, } 复制代码`

当开发者自定义 ` computed` 计算属性时，如下 ` messageT` 属性，是依赖 ` message` 的变化的。因此 ` message` 变化时我们也要通知到 ` computed` ，让它去执行回调函数。 这个时候 ` message` 的 ` Dep` 就收集到了三个依赖，这个依赖就是用来管理 ` computed` 中 ` message` 变化的。

` computed: { messageT() { return this.message + '!' ; } } 复制代码`

图示如下：一个属性可能有多个依赖，每个响应式数据都有一个 ` Dep` 来管理它的依赖。

![依赖收集](https://user-gold-cdn.xitu.io/2019/6/2/16b1857fd4532ff0?imageView2/0/w/1280/h/960/ignore-error/1)

### 如何收集依赖 ###

我们如何知道 ` data` 中的某个属性被使用了，答案就是 ` Object.defineProperty` ，因为读取某个属性就会触发 ` get` 方法。可以将代码进行如下改造：

` function defineReactive ( obj, key, val ) { let Dep; // 依赖 Object.defineProperty(obj, key, { enumerable : true , configurable : true , get : () => { console.log( '我被读了，我要不要做点什么好?' ); // 被读取了，将这个依赖收集起来 Dep.depend(); // 本次新增 return val; }, set : newVal => { if (val === newVal) { return ; } val = newVal; // 被改变了，通知依赖去更新 Dep.notify(); // 本次新增 console.log( "数据被改变了，我要把新的值渲染到页面上去!" ); } }) } 复制代码`

### 什么是依赖 ###

那所谓的依赖究竟是什么呢？上面的图中已经暴露了答案，就是 ` Watcher` 。

## Watcher 「中介」 ##

` Watcher` 就是类似中介的角色，比如 ` message` 就有三个中介，当 ` message` 变化，就通知这三个中介，他们就去执行各自需要做的变化。

` Watcher` 能够控制自己属于哪个，是 ` data` 中的属性的还是 ` watch` ，或者是 ` computed` ， ` Watcher` 自己有统一的更新入口，只要你通知它，就会执行对应的更新方法。

因此我们可以推测出， ` Watcher` 必须要有的2个方法。一个就是通知变化，另一个就是被收集起来到Dep中去。

` class Watcher { addDep() { // 我这个Watcher要被塞到Dep里去了~~ }, update() { // Dep通知我更新呢~~ }, } 复制代码`

## 总结 ##

回顾一下， ` Vue` 响应式原理的核心就是 ` Observer` 、 ` Dep` 、 ` Watcher` 。

` Observer` 中进行响应式的绑定，在数据被读的时候，触发 ` get` 方法，执行 ` Dep` 来收集依赖，也就是收集 ` Watcher` 。

在数据被改的时候，触发 ` set` 方法，通过对应的所有依赖( ` Watcher` )，去执行更新。比如 ` watch` 和 ` computed` 就执行开发者自定义的回调方法。

本篇文章属于入门篇，能够先简单的理解 ` Observer` 、 ` Dep` 、 ` Watcher` 三者的作用和关系。后面会逐渐详细和深入，循序渐进的理解和学习。

如果你觉得对你有帮助，就点个赞吧~

正在书写的系列~

* [1. Vue响应式原理-理解Observer、Dep、Watcher]( https://juejin.im/post/5cf3cccee51d454fa33b1860 )
* [2. 响应式原理-如何监听Array的变化]( https://juejin.im/post/5cf606d6f265da1b8e708ba6 )

[Github博客 ]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyukiyang0729%2Fblog ) 欢迎交流~