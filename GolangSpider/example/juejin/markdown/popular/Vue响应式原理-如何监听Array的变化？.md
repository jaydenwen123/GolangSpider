# Vue响应式原理-如何监听Array的变化？ #

## 回忆 ##

在上一篇 [Vue响应式原理-理解Observer、Dep、Watcher]( https://juejin.im/post/5cf3cccee51d454fa33b1860 ) 简单讲解了 ` Observer` 、 ` Dep` 、 ` Watcher` 三者的关系。

在 ` Observer` 的伪代码中我们模拟了如下代码:

` class Observer { constructor () { // 响应式绑定数据通过方法 observe( this.data); } } export function observe ( data ) { const keys = Object.keys(data); for ( let i = 0 ; i < keys.length; i++) { // 将data中我们定义的每个属性进行响应式绑定 defineReactive(obj, keys[i]); } } export function defineReactive ( ) { // ...省略 Object.defineProperty get-set } 复制代码`

今天我们就进一步了解 ` Observer` 里还做了什么事。

## Array的变化如何监听？ ##

` data` 中的数据如果是一个数组怎么办？我们发现 ` Object.defineProperty` 对数组进行响应式化是有缺陷的。

虽然我们可以监听到索引的改变。

` function defineReactive ( obj, key, val ) { Object.defineProperty(obj, key, { enumerable : true , configurable : true , get : () => { console.log( '我被读了，我要不要做点什么好?' ); return val; }, set : newVal => { if (val === newVal) { return ; } val = newVal; console.log( "数据被改变了，我要渲染到页面上去!" ); } }) } let data = [ 1 ]; // 对数组key进行监听 defineReactive(data, 0 , 1 ); console.log(data[ 0 ]); // 我被读了，我要不要做点什么好? data[ 0 ] = 2 ; // 数据被改变了，我要渲染到页面上去! 复制代码`

但是 ` defineProperty` 不能检测到数组长度的变化，准确的说是 **通过改变length** 而增加的长度不能监测到。这种情况无法触发任何改变。

` data.length = 0 ; // 控制台没有任何输出 复制代码`

而且监听数组所有索引的的代价也比较高，综合一些其他因素，Vue用了另一个方案来处理。

首先我们的 ` observe` 需要改造一下，单独加一个数组的处理。

` // 将data中我们定义的每个属性进行响应式绑定 export function observe ( data ) { const keys = Object.keys(data); for ( let i = 0 ; i < keys.length; i++) { // 如果是数组 if ( Array.isArray(keys[i])) { observeArray(keys[i]); } else { // 如果是对象 defineReactive(obj, keys[i]); } } } // 数组的处理 export function observeArray ( ) { // ...省略 } 复制代码`

那接下来我们就应该考虑下 ` Array` 变化如何监听？

` Vue` 中对这个数组问题的解决方案非常的简单粗暴，就是对能够改变数组的方法做了一些手脚。

我们知道，改变数组的方法有很多，举个例子比如说 ` push` 方法吧。 ` push` 存在 ` Array.prototype` 上的，如果我们能

能拦截到原型上的 ` push` 方法，是不是就可以做一些事情呢？

## Object.defineProperty ##

对象里目前存在的属性描述符有两种主要形式： **数据描述符** 和 **存取描述符** 。 **存取描述符** 是由getter-setter函数对描述的属性，也就是我们用来给对象做响应式绑定的。 [Object.defineProperty-MDN]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FJavaScript%2FReference%2FGlobal_Objects%2FObject%2FdefineProperty )

虽然我们无法使用 ` Object.defineProperty` 将数组进行响应式的处理，也就是 ` getter-setter` ，但是还有其他的功能可以供我们使用。就是 **数据描述符** ， **数据描述符** 是一个具有值的属性，该值可能是可写的，也可能不是可写的。

### value ###

> 
> 
> 
> 该属性对应的值。可以是任何有效的 JavaScript 值（数值，对象，函数等）。 **默认为 undefined** 。
> 
> 

### writable ###

> 
> 
> 
> 当且仅当该属性的 ` writable` 为 ` true` 时， ` value` 才能被 [赋值运算符](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FJavaScript%2FReference%2FOperators%2FAssignment_Operators
> ) 改变。 **默认为 false** 。
> 
> 

因此我们只要把原型上的方法，进行 ` value` 的重新赋值。

如下代码，在重新赋值的过程中，我们可以获取到方法名和所有参数。

` function def ( obj, key ) { Object.defineProperty(obj, key, { writable : true , enumerable : true , configurable : true , value : function (...args ) { console.log( 'key' , key); console.log( 'args' , args); } }); } // 重写的数组方法 let obj = { push() {} } // 数组方法的绑定 def(obj, 'push' ); obj.push([ 1 , 2 ], 7 , 'hello!' ); // 控制台输出 key push // 控制台输出 args [Array(2), 7, "hello!"] 复制代码`

通过如上代码我们就可以知道，用户使用了数组上原型的方法以及参数我们都可以拦截到，这个拦截的过程就可以做一些变化的通知。

## Vue监听Array三步曲 ##

接下来，就看看 ` Vue` 是如何实现的吧~

第一步：先获取原生 ` Array` 的原型方法，因为拦截后还是需要原生的方法帮我们实现数组的变化。

第二步：对 ` Array` 的原型方法使用 ` Object.defineProperty` 做一些拦截操作。

第三步：把需要被拦截的 ` Array` 类型的数据原型指向改造后原型。

我们将代码进行下改造，拦截的过程中还是要将开发者的参数传给原生的方法，保证数组按照开发者的想法被改变，然后我们再去做视图的更新等操作。

` const arrayProto = Array.prototype // 获取Array的原型 function def ( obj, key ) { Object.defineProperty(obj, key, { enumerable : true , configurable : true , value : function (...args ) { console.log(key); // 控制台输出 push console.log(args); // 控制台输出 [Array(2), 7, "hello!"] // 获取原生的方法 let original = arrayProto[key]; // 将开发者的参数传给原生的方法，保证数组按照开发者的想法被改变 const result = original.apply( this , args); // do something 比如通知Vue视图进行更新 console.log( '我的数据被改变了，视图该更新啦' ); this.text = 'hello Vue' ; return result; } }); } // 新的原型 let obj = { push() {} } // 重写赋值 def(obj, 'push' ); let arr = [ 0 ]; // 原型的指向重写 arr.__proto__ = obj; // 执行push arr.push([ 1 , 2 ], 7 , 'hello!' ); console.log(arr); 复制代码`

被改变后的 ` arr` 。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2017973a7931a?imageView2/0/w/1280/h/960/ignore-error/1)

## Vue源码解析 ##

#### array.js ####

` Vue` 在 ` array.js` 中重写了 ` methodsToPatch` 中七个方法，并将重写后的原型暴露出去。

` // Object.defineProperty的封装 import { def } from '../util/index' // 获得原型上的方法 const arrayProto = Array.prototype // Vue拦截的方法 const methodsToPatch = [ 'push' , 'pop' , 'shift' , 'unshift' , 'splice' , 'sort' , 'reverse' ]; // 将上面的方法重写 methodsToPatch.forEach( function ( method ) { def(arrayMethods, method, function mutator (...args ) { console.log( 'method' , method); // 获取方法 console.log( 'args' , args); // 获取参数 // ...功能如上述，监听到某个方法执行后，做一些对应的操作 // 1、将开发者的参数传给原生的方法，保证数组按照开发者的想法被改变 // 2、视图更新等 }) }) export const arrayMethods = Object.create(arrayProto); 复制代码`

## observer ##

在进行数据 ` observer` 绑定的时候，我们先判断是否 ` hasProto` ，如果存在 ` __proto__` ，就直接将 ` value` 的 ` __proto__` 指向重写过后的原型。如果不能使用 ` __proto__` ，貌似有些浏览器厂商没有实现。那就直接循环 ` arrayMethods` 把它身上的这些方法直接装到 ` value` 身上好了。毕竟调用某个方法是先去自身查找，当自身找不到这关方法的时候，才去原型上查找。

` // 判断是否有__proto__，因为部分浏览器是没有__proto__ const hasProto = '__proto__' in {} // 重写后的原型 import { arrayMethods } from './array' // 方法名 const arrayKeys = Object.getOwnPropertyNames(arrayMethods); // 数组的处理 export function observeArray ( value ) { // 如果有__proto__，直接覆盖 if (hasProto) { protoAugment(value, arrayMethods); } else { // 没有__proto__就把方法加到属性自身上 copyAugment(value, arrayMethods, ) } } // 原型的赋值 function protoAugment ( target, src ) { target.__proto__ = src; } // 复制 function copyAugment ( target, src, keys ) { for ( let i = 0 , l = keys.length; i < l; i++) { const key = keys[i] def(target, key, src[key]); } } 复制代码`

通过上面的代码我们发现，没有直接修改 ` Array.prototype` ，而是直接把 ` arrayMenthods` 赋值给 ` value` 的 ` __proto__` 。因为这样不会污染全局的Array， ` arrayMenthods` 只对 ` data` 中的 ` Array` 生效。

## 总结 ##

因为监听的数组带来的代价和一些问题， ` Vue` 使用了重写原型的方案代替。拦截了数组的一些方法，在这个过程中再去做通知变化等操作。

本文的一些代码均是 ` Vue` 源码简化后的，为了方便大家理解。思想理解了，源码就容易看懂了。

正在书写的系列~

* [1. Vue响应式原理-理解Observer、Dep、Watcher]( https://juejin.im/post/5cf3cccee51d454fa33b1860 )
* [2. 响应式原理-如何监听Array的变化]( https://juejin.im/post/5cf606d6f265da1b8e708ba6 )

[Github博客 ]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyukiyang0729%2Fblog ) 欢迎交流~