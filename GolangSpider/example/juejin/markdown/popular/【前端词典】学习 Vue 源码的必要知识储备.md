# 【前端词典】学习 Vue 源码的必要知识储备 #

## 前言 ##

我最近在写 Vue 进阶的内容。在这个过程中，有些人问我看 Vue 源码需要有哪些准备吗？所以也就有了这篇计划之外的文章。

> 
> 
> 
> 当你想学习 Vue 源码的时候，需要有扎实的 JavaScript 基础，下面罗列的只是其中的一部分比较具有代表性的知识点。如果你还不具备
> JavaScript 基础的话，建议不要急着看 Vue 源码，这样你会很容易放弃的。
> 
> 

我会从以下 7 点来展开：

* Flow 基本语法
* 发布/订阅模式
* Object.defineProperty
* ES6+ 语法
* 原型链、闭包
* 函数柯里化
* event loop

![](https://user-gold-cdn.xitu.io/2019/5/29/16b03fc4381323d1?imageView2/0/w/1280/h/960/ignore-error/1)

## 必要知识储备 ##

需要注意的是这篇文章每个点不会讲的特别详细，我这里就是把一些知识点归纳一下。每个详细的点仍需自己花时间学习。

### Flow 基本语法 ###

相信看过 Vue、Vuex 等源码的人都知道它们使用了 Flow 静态类型检查工具。

我们知道 JavaScript 是弱类型的语言，所以我们在写代码的时候容易出现一些始料未及的问题。也正是因为这个问题，才出现了 Flow 这个静态类型检查工具。

> 
> 
> 
> 这个工具可以改变 JavaScript 是弱类型的语言的情况，可以加入类型的限制，提高代码质量。
> 
> 

` // 未使用 Flow 限制 function sum(a, b) { return a + b; } // 使用 Flow 限制 a b 都是 number 类型。 function sum(a: number, b:number) { return a + b; } 复制代码`

#### 基础检测类型 ####

Flow 支持原始数据类型，有如下几种：

` boolean number string null void( 对应 undefined ) 复制代码`

在定义变量的同时在关键的地方声明类型，使用如下：

` let str:string = 'str' ; // 重新赋值 str = 3 // 报错 复制代码`

#### 复杂类型检测 ####

Flow 支持复杂类型检测，有如下几种：

` Object Array Function 自定义的 Class 复制代码`
> 
> 
> 
> 需要注意直接使用 flow.js，JavaScript 是无法在浏览器端运行的，必须借助 babel 插件，vue 源码中使用的是
> babel-preset-flow-vue 这个插件，并且在 babelrc 进行配置。
> 
> 

详细的 Flow 语法可以看以下资料：

**这里推荐两个资料**

* 官方文档： [flow.org/en/]( https://link.juejin.im?target=https%3A%2F%2Fflow.org%2Fen%2F )
* Flow 的使用入门： [zhuanlan.zhihu.com/p/26204569]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F26204569 )

### 发布/订阅模式 ###

我们知道 Vue 是内部是实现了双向绑定机制，使得我们不用再像从前那样还要自己操作 DOM 了。

其实 Vue 的双向绑定机制采用 **数据劫持结合发布/订阅模式** 实现的: 通过 ` Object.defineProperty()` 来劫持各个属性的 ` setter，getter` ，在数据变动时发布消息给订阅者，触发相应的监听回调。

> 
> 
> 
> 我发现有的人把观察者模式和发布/订阅模式混淆一谈，其实订阅模式有一个调度中心，对订阅事件进行统一管理。而观察者模式可以随意注册事件，调用事件。
> 
> 

我画了一个大概的流程图，用来说明观察者模式和发布/订阅模式。如下：

![](https://user-gold-cdn.xitu.io/2019/5/26/16af4afbf1806526?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 这块我会在接下的文章中详细讲到，这里先给出一个概念，感兴趣的可以自己查找资料，也可等我的文章出炉。
> 
> 

其实我们对这种模式再熟悉不过了，但可能你自己也没发现：

` let div = document.getElementById( '#div' ); div.addEventListener( 'click' , () => { console.log( "div 被点击了一下" ) }) 复制代码`

可以思考下上面的事件绑定执行的一个过程，你应该会有共鸣。

### 函数柯里化 ###

### 数据双向绑定基础：Object.defineProperty() ###

#### 一、数据属性 ####

数据属性包含一个数据值的位置。这个位置可以读取和写入值。数据属性有 4 个描述他行为的特性：

+--------------+--------------------------------------------------------------+
|     属性     |                             描述                             |
+--------------+--------------------------------------------------------------+
| Configurable | 能否用 delete                                                |
|              | 删除属性从而重新定义属性。默认为                             |
|              | true                                                         |
| Enumerable   | 能否通过 for-in                                              |
|              | 遍历，即是否可枚举。默认为                                   |
|              | true                                                         |
| Writable     | 是否能修改属性的值。默认为                                   |
|              | true                                                         |
| Value        | 包含这个属性的数据值，读写属性的时候其实就在这里读写。默认为 |
|              | undefined                                                    |
+--------------+--------------------------------------------------------------+

**如果你想要修改上述 4 个默认的数据属性** ，就需要使用 ECMAScript 的 Object.defineProperty() 方法。

> 
> 
> 
> 该方法包含3个参数：属性所在的对象，属性名，描述符对象。 **描述符对象的属性必须在上述 4 个属性中。**
> 
> 

` var person = { name: '' , }; // 不能修改属性的值 Object.defineProperty(person, "name" ,{ writable: false , value: "小生方勤" }); console.log(person.name); // "小生方勤" person.name = "方勤" ; console.log(person.name); // "小生方勤" 复制代码`

#### 二、访问器属性 ####

访问器属性不包含数据值，他们包含一对 ` getter` 和 ` setter` 函数（非必须）。在读写访问器属性的值的时候，会调用相应的 ` getter` 和 ` setter` 函数，而我们的 vue 就是在 ` getter` 和 ` setter` 函数中增加了我们需要的操作。

> 
> 
> 
> 需要注意的是【value 或 writable】一定不能和【get 或 set】共存。
> 
> 

访问器属性有以下 4 个特性：

+--------------+----------------------------------+
|     特性     |               描述               |
+--------------+----------------------------------+
| Configurable | 能否用 delete                    |
|              | 删除属性从而重新定义属性。默认为 |
|              | true                             |
| Enumerable   | 能否通过 for-in                  |
|              | 遍历，即是否可枚举。默认为       |
|              | true                             |
| get          | 读取属性时调用的函数，默认       |
|              | undefined                        |
| set          | 写入属性时调用的函数，默认       |
|              | undefined                        |
+--------------+----------------------------------+

接下来给个例子：

` var person = { _name : "小生方勤" }; Object.defineProperty(person, "name" , { //注意 person 多定义了一个 name 属性 set : function (value){ this._name = "来自 setter : " + value; }, get: function (){ return "来自 getter : " + this._name; } }); console.log( person.name ); // 来自 getter : 小生方勤 person.name = "XSFQ" ; console.log( person._name ); // 来自 setter : XSFQ console.log( person.name ); // 来自 getter : 来自 setter : XSFQ 复制代码`
> 
> 
> 
> 
> 如果之前都不清楚有 Object.defineProperty() 方法，建议你看《JavaScript 高级程序设计》的 139 - 144 页。
> 
> 
> 

#### 额外讲讲 Object.create(null) ####

> 
> 
> 
> 我们在源码随处可以 ` this.set = Object.create(null)` 这样的赋值。为什么这样做呢？这样写的好处就是不需要考虑原型链上的属性，
> **可以真正的创建一个纯净的对象。**
> 
> 

首先 Object.create 可以理解为继承一个对象，它是 ES5 的一个特性，对于旧版浏览器需要做兼容，基本代码如下：

` if (!Object.create) { Object.create = function (o) { function F () {} // 定义了一个隐式的构造函数 F.prototype = o; return new F(); // 其实还是通过new来实现的 }; } 复制代码`

### ES6+ 语法 ###

其实这点应该是默认你需要知道的，不过鉴于之前有人问过我一些相关的问题，我稍微讲一下。

#### ` export default` 和 ` export` 的区别 ####

* 在一个文件或模块中 ` export` 可以有多个，但 ` export default` 仅有一个
* 通过 ` export` 方式导出，在导入时要加 { }，而 ` export default` 则不需要
` 1.export //a.js export const str = "小生方勤" ; //b.js import { str } from 'a' ; // 导入的时候需要花括号 2.export default //a.js const str = "小生方勤" ; export default str; //b.js import str from 'a' ; // 导入的时候无需花括号 复制代码`
> 
> 
> 
> 
> ` export default const a = 1;` 这样写是会报错的哟。
> 
> 

#### 箭头函数 ####

这个一笔带过：

* 箭头函数中的 this 指向是固定不变的，即是在定义函数时的指向
* 而普通函数中的 this 指向时变化的，即是在使用函数时的指向

#### class 继承 ####

Class 可以通过 ` extends` 关键字实现继承，这比 ES5 的通过修改原型链实现继承，要清晰和方便很多。

` class staff { constructor (){ this.company = "ABC" ; this.test = [1,2,3]; } companyName (){ return this.company; } } class employee extends staff { constructor(name,profession){ super(); this.employeeName = name; this.profession = profession; } } // 将父类原型指向子类 let instanceOne = new employee( "Andy" , "A" ); let instanceTwo = new employee( "Rose" , "B" ); instanceOne.test.push(4); // 测试 console.log(instanceTwo.test); // [1,2,3] console.log(instanceOne.companyName()); // ABC // 通过 Object.getPrototypeOf() 方法可以用来从子类上获取父类 console.log(Object.getPrototypeOf(employee) === staff) // 通过 hasOwnProperty() 方法来确定自身属性与其原型属性 console.log(instanceOne.hasOwnProperty( 'test' )) // true // 通过 isPrototypeOf() 方法来确定原型和实例的关系 console.log(staff.prototype.isPrototypeOf(instanceOne)); // true 复制代码`

` super` 关键字，它在这里表示父类的构造函数，用来新建父类的 ` this` 对象。

> 
> * 子类必须在 ` constructor` 方法中调用 ` super` 方法，否则新建实例时会报错。这是因为子类没有自己的 ` this` 对象，而是继承父类的
> ` this` 对象，然后对其进行加工。
> * 只有调用 ` super` 之后，才可以使用 ` this` 关键字，否则会报错。这是因为子类实例的构建，是基于对父类实例加工，只有 `
> super` 方法才能返回父类实例。
> 

> 
> 
> 
> `super` 虽然代表了父类 `A` 的构造函数，但是返回的是子类 `B` 的实例，即` super` 内部的 `this ` 指的是
> `B`，因此 `super()` 在这里相当于 A.prototype.constructor.call(this)
> 
> 

**ES5 和 ES6 实现继承的区别**

ES5 的继承，实质是先创造子类的实例对象 ` this` ，然后再将父类的方法添加到 ` this` 上面（ ` Parent.apply(this)` ）。
ES6 的继承机制完全不同，实质是先创造父类的实例对象 ` this` （所以必须先调用 ` super()` 方法），然后再用子类的构造函数修改 ` this` 。

#### proxy ####

对最新动态了解的人就会知道，在下一个版本的 Vue 中，会使用 ` proxy` 代替 ` Object.defineProperty` 完成数据劫持的工作。

> 
> 
> 
> 尤大说，这个新的方案会使初始化速度加倍，于此同时内存占用减半。
> 
> 

proxy 对象的用法:

` var proxy = new Proxy(target, handler); 复制代码`

new Proxy() 即生成一个 Proxy 实例。target 参数表示所要拦截的目标对象，handler 参数也是一个对象，用来定制拦截行为。

` var proxy = new Proxy({}, { get: function (obj, prop) { console.log( 'get 操作' ) return obj[prop]; }, set : function (obj, prop, value) { console.log( 'set 操作' ) obj[prop] = value; } }); proxy.num = 2; // 设置 set 操作 console.log(proxy.num); // 设置 get 操作 // 2 复制代码`
> 
> 
> 
> 
> 除了 get 和 set 之外，proxy 可以拦截多达 13 种操作。
> 
> 

> 
> 
> 
> 注意，proxy 的最大问题在于浏览器支持度不够，IE 完全不兼容。
> 
> 

倘若你基本不了解 ES6, 推荐下面这个教程：

阮一峰 ECMAScript 6 入门： [es6.ruanyifeng.com/]( https://link.juejin.im?target=http%3A%2F%2Fes6.ruanyifeng.com%2F )

### 原型链、闭包 ###

#### 原型链 ####

因为之前我特意写了一篇文章来解释原型链，所以这里就不在讲述了：

原型链： [juejin.im/post/5c3359…]( https://juejin.im/post/5c335940f265da610e804097 )

#### 闭包 ####

这里我先放一段 Vue 源码中的 once 函数。这就是闭包调用 —— 函数作为返回值：

` /** * Ensure a function is called only once. */ export function once (fn: Function): Function { let called = false return function () { if (!called) { called = true fn.apply(this, arguments) } } } 复制代码`

这个函数的作用就是确保函数只调用一次。

> 
> 
> 
> 为什么只会调用一次呢? 因为函数调用完成之后，其执行上下文环境不会被销毁，所以 called 的值依然在那里。
> 
> 

闭包到底是什么呢。《JavaScript 高级程序设计》的解释是：

> 
> 
> 
> 闭包是指有权访问另一个函数作用域中的变量的函数。创建闭包的常见方式，就是在一个函数内部创建另一个函数。
> 
> 

> 
> 
> 
> 简单讲，闭包就是指有权访问另一个函数作用域中的变量的函数。
> 
> 

给两段代码，如果你知道他们的运行结果，那么说明你是了解闭包的：

` // 第一段 var num = 20; function fun (){ var num = 10; return function con (){ console.log( this.num ) } } var funOne = fun(); funOne(); // 20 // 第二段 var num = 20; function fun (){ var num = 10; return function con (){ console.log( num ) } } var funOne = fun(); funOne(); // 10 复制代码`

### 函数柯里化 ###

> 
> 
> 
> 所谓"柯里化"，就是把一个多参数的函数，转化为单参数函数。
> 
> 

先说说我之前遇到过得一个面试题：

> 
> 
> 
> 如何使 ` add(2)(3)(4)(5)()` 输出 14
> 
> 

在那次面试的时候，我还是不知道柯里化这个概念的，所以当时我没答上。后来我才知道这可以用函数柯里化来解，即：

` function add(num){ var sum=0; sum= sum+num; return function tempFun(numB){ if (arguments.length===0){ return sum; } else { sum= sum+ numB; return tempFun; } } } 复制代码`

那这和 Vue 有什么关系呢？当然是有关系的：

我们是否经常这样写判断呢？

` if ( A ){ // code } else if ( B ){ // code } 复制代码`

这个写法没什么问题，可是在重复的出现这种相同的判断的时候。这个就显得有点不那么智能了。这个时候函数柯里化就可以排上用场了。

因为 Vue 可以在不同平台运行，所以也会存在上面的那种判断。这里利用柯里化的特点，通过 createPatchFunction 方法把一些参数提前保存，以便复用。

` // 这样不用每次调用 patch 的时候都传递 nodeOps 和 modules export function createPatchFunction (backend) { // 省略好多代码 return function patch (oldVnode, vnode, hydrating, removeOnly) { // 省略好多代码 } } 复制代码`

### event loop ###

四个概念：

* 同步任务：即在主线程上排队执行的任务，只有前一个任务执行完毕，才能执行后一个任务。
* 异步任务：指的是不进入主线程，某个异步任务可以执行了，该任务才会进入主线程执行。
* macrotask：主要场景有：主代码块、setTimeout、setInterval等
* microtask：主要场景有：Promise、process.nextTick等。

这一点网上教程已经很多了，再因为篇幅的问题，这里就不详细说了。

推荐一篇文章，说的很细致：

JavaScript 执行机制： [juejin.im/post/59e85e…]( https://juejin.im/post/59e85eebf265da430d571f89#heading-4 )

## 总结 ##

这篇文章讲到这里就结束了。不过有一点我需要在说一篇，这篇文章的定位并不是面面俱到的将所有知识都讲一遍，这不现实我也没这个能力。

我只是希望通过这篇文章告诉大家一个观点，要想看源码，一些必备的 JavaScript 基础知识必须要扎实，否则你会举步维艰。

愿你每天都有进步。

## Vue 相关文章输出计划 ##

最近总有朋友问我 Vue 相关的问题，因此接下来我会输出 9 篇 Vue 相关的文章，希望对大家有一定的帮助。我会保持在 7 到 10 天更新一篇。

* [【前端词典】Vuex 注入 Vue 生命周期的过程（完成）]( https://juejin.im/post/5cb30243e51d456e431ada29 )
* 【前端词典】浅析 Vue 响应式原理
* 【前端词典】新老 VNode 进行 patch 的过程
* 【前端词典】如何开发功能组件并上传 npm
* 【前端词典】从这几个方面优化你的 Vue 项目
* 【前端词典】从 Vue-Router 设计讲前端路由发展
* 【前端词典】在项目中如何正确的使用 Webpack
* 【前端词典】Vue 服务端渲染
* 【前端词典】Axios 与 Fetch 该如何选择

建议你关注我的公众号，第一时间就可以接收最新的文章。

![](https://user-gold-cdn.xitu.io/2019/5/19/16acfa600f92404e?imageView2/0/w/1280/h/960/ignore-error/1)

如果你想加群交流，也可以添加有点智能的机器人，自动拉你进群：

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0120727720670?imageView2/0/w/1280/h/960/ignore-error/1)