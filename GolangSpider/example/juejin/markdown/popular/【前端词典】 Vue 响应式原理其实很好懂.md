# 【前端词典】 Vue 响应式原理其实很好懂 #

## 前言 ##

这是十篇 Vue 系列文章的第三篇，这篇文章我们讲讲 Vue 最核心的功能之一 —— 响应式原理。

## 如何理解响应式 ##

可以这样理解：当一个状态改变之后，与这个状态相关的事务也立即随之改变，从前端来看就是数据状态改变后相关 DOM 也随之改变。数据模型仅仅是普通的 JavaScript 对象。而当你修改它们时，视图会进行更新。

## 抛个问题 ##

我们先看看我们在 Vue 中常见的写法：

` <div id= "app" @click= "changeNum" > {{ num }} </div> var app = new Vue({ el: '#app' , data: { num: 1 }, methods: { changeNum () { this.num = 2 } } }) 复制代码`
> 
> 
> 
> 
> 这种写法很常见，不过你考虑过当为什么执行 ` this.num = 2` 后视图为什么会更新呢？通过这篇文章我力争把这个点讲清楚。
> 
> 

## 如果不使用 Vue，我们应该怎么实现？ ##

我的第一想法是像下面这样实现：

` let data = { num: 1 }; Object.defineProperty(data, 'num' ,{ value: value, set : function ( newVal ){ document.getElementById( 'app' ).value = newVal; } }); input.addEventListener( 'input' , function (){ data.num = 2; }); 复制代码`

这样可以粗略的实现点击元素，自动更新视图。

> 
> 
> 
> 这里我们需要通过 Object.defineProperty 来操作对象的访问器属性。监听到数据变化的时候，操作相关 DOM。
> 
> 

而这里用到了一个常见模式 —— 发布/订阅模式。

我画了一个大概的流程图，用来说明观察者模式和发布/订阅模式。如下：

![](https://user-gold-cdn.xitu.io/2019/5/26/16af4afbf1806526?imageView2/0/w/1280/h/960/ignore-error/1)

仔细的同学会发现，我这个粗略的过程和使用 Vue 的不同的地方就是需要我自己操作 DOM 重新渲染。

> 
> 
> 
> 如果我们使用 Vue 的话，这一步就是 Vue 内部的代码来处理的。这也是我们为什么在使用 Vue 的时候无需手动操作 DOM 的原因。
> 
> 

关于 ` Object.defineProperty` 我在上一篇文章已经提及，这里就不再复述。

## Vue 是如何实现响应式的 ##

我们知道对象可以通过 ` Object.defineProperty` 操作其访问器属性，即对象拥有了 ` getter` 和 ` setter` 方法。这就是实现响应式的基石。

先看一张很直观的流程图：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22c8fe2629e68?imageView2/0/w/1280/h/960/ignore-error/1)

### initData 方法 ###

在 Vue 的初始化的时候，其 ` _init()` 方法会调用执行 ` initState(vm)` 方法。 ` initState` 方法主要是对 ` props` 、 ` methods` 、 ` data` 、 ` computed` 和 ` wathcer` 等属性做了初始化操作。

这里我们就对 ` data` 初始化的过程做一个比较详细的分析。

` function initData (vm: Component) { let data = vm. $options.data data = vm._data = typeof data === 'function' ? getData(data, vm) : data || {} if (!isPlainObject(data)) { ...... } // proxy data on instance const keys = Object.keys(data) const props = vm. $options.props const methods = vm. $options.methods let i = keys.length while (i--) { const key = keys[i] ...... // 省略部分兼容代码，但不影响理解 if (props && hasOwn(props, key)) { ...... } else if (!isReserved(key)) { proxy(vm, `_data`, key) } } // observe data observe(data, true /* asRootData */) } 复制代码`

` initData` 初始化 data 的主要过程也是做两件事:

* 通过 ` proxy` 把每一个值 ` vm._data.[key]` 都代理到 ` vm.[key]` 上；
* 调用 ` observe` 方法观测整个 data 的变化，把 data 也变成响应式（可观察），可以通过 ` vm._data.[key]` 访问到定义 data 返回函数中对应的属性。

### 数据劫持 — ` Observe` ###

通过这个方法将 data 下面的所有属性变成响应式（可观察）。

` // 给对象的属性添加 getter 和 setter，用于依赖收集和发布更新 export class Observer { value: any; dep: Dep; vmCount: number; constructor (value: any) { this.value = value // 实例化 Dep 对象 this.dep = new Dep() this.vmCount = 0 // 把自身实例添加到数据对象 value 的 __ob__ 属性上 def(value, '__ob__' , this) // value 是否为数组的不同调用 if (Array.isArray(value)) { const augment = hasProto ? protoAugment : copyAugment augment(value, arrayMethods, arrayKeys) this.observeArray(value) } else { this.walk(value) } } // 取出所有属性遍历 walk (obj: Object) { const keys = Object.keys(obj) for ( let i = 0; i < keys.length; i++) { defineReactive(obj, keys[i]) } } observeArray (items: Array<any>) { for ( let i = 0, l = items.length; i < l; i++) { observe(items[i]) } } } 复制代码`

` def` 函数内封装了 ` Object.defineProperty` ，所以你 console.log(data) ，会发现多了一个 ` __ob__` 的属性。

### defineReactive 方法遍历所有属性 ###

` // 定义一个响应式对象的具体实现 export function defineReactive ( obj: Object, key: string, val: any, customSetter?: ?Function, shallow?: boolean ) { const dep = new Dep() ..... // 省略部分兼容代码，但不影响理解 let childOb = !shallow && observe(val) Object.defineProperty(obj, key, { enumerable: true , configurable: true , get: function reactiveGetter () { const value = getter ? getter.call(obj) : val if (Dep.target) { // 进行依赖收集 dep.depend() if (childOb) { childOb.dep.depend() if (Array.isArray(value)) { dependArray(value) } } } return value }, set : function reactiveSetter (newVal) { const value = getter ? getter.call(obj) : val ..... // 省略部分兼容代码，但不影响理解 if (setter) { setter.call(obj, newVal) } else { val = newVal } // 对新的值进行监听 childOb = !shallow && observe(newVal) // 通知所有订阅者，内部调用 watcher 的 update 方法 dep.notify() } }) } 复制代码`

` defineReactive` 方法最开始初始化 Dep 对象的实例，然后通过对子对象递归调用 ` observe` 方法，使所有子属性也能变成响应式的对象。并且在 ` Object.defineProperty` 的 ` getter` 和 ` setter` 方法中调用 ` dep` 的相关方法。

即：

* ` getter` 方法完成的工作就是依赖收集 —— ` dep.depend()`
* ` setter` 方法完成的工作就是发布更新 —— ` dep.notify()`

我们发现这里都和 Dep 对象有着不可忽略的关系。接下来我们就看看 Dep 对象。这个 Dep

### 调度中心作用的 Dep ###

前文中我们提到发布/订阅模式，在发布者和订阅者之前有一个调度中心。这里的 Dep 扮演的角色就是调度中心，主要的作用就是:

* 收集订阅者 Watcher 并添加到观察者列表 subs
* 接收发布者的事件
* 通知订阅者目标更新，让订阅者执行自己的 update 方法

详细代码如下：

` // Dep 构造函数 export default class Dep { static target: ?Watcher; id: number; subs: Array<Watcher>; constructor () { this.id = uid++ this.subs = [] } // 向 dep 的观察者列表 subs 添加 Watcher addSub (sub: Watcher) { this.subs.push(sub) } // 从 dep 的观察者列表 subs 移除 Watcher removeSub (sub: Watcher) { remove(this.subs, sub) } // 进行依赖收集 depend () { if (Dep.target) { Dep.target.addDep(this) } } // 通知所有订阅者，内部调用 watcher 的 update 方法 notify () { const subs = this.subs.slice() for ( let i = 0, l = subs.length; i < l; i++) { subs[i].update() } } } // Dep.target 是全局唯一的观察者，因为在任何时候只有一个观察者被处理。 Dep.target = null // 待处理的观察者队列 const targetStack = [] export function pushTarget (_target: ?Watcher) { if (Dep.target) targetStack.push(Dep.target) Dep.target = _target } export function popTarget () { Dep.target = targetStack.pop() } 复制代码`

Dep 可以理解成是对 ` Watcher` 的一种管理，Dep 和 ` Watcher` 是紧密相关的。所以我们必须看一看 ` Watcher` 的实现。

### 订阅者 —— Watcher ###

` Watcher` 中定义了许多原型方法，这里我只粗略的讲 ` update` 和 ` get` 这三个方法。

` // 为了方便理解，部分兼容代码已被我省去 get () { // 设置需要处理的观察者 pushTarget(this) const vm = this.vm let value = this.getter.call(vm, vm) // deep 是否为 true 的处理逻辑 if (this.deep) { traverse(value) } // 将 Dep.target 指向栈顶的观察者，并将他从待处理的观察者队列中移除 popTarget() // 执行依赖清空动作 this.cleanupDeps() return value } update () { if (this.computed) { ... } else if (this.sync) { // 标记为同步 this.run() } else { // 一般都是走这里，即异步批量更新：nextTick queueWatcher(this) } } 复制代码`

Vue 的响应式过程大概就是这样了。感兴趣的可以看看源码。

最后我们在通过这个流程图来复习一遍：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b287f7a78e5cab?imageView2/0/w/1280/h/960/ignore-error/1)

## Vue 相关文章输出计划 ##

最近总有朋友问我 Vue 相关的问题，因此接下来我会输出 9 篇 Vue 相关的文章，希望对大家有一定的帮助。我会保持在 7 到 10 天更新一篇。

* [【前端词典】Vuex 注入 Vue 生命周期的过程（完成）]( https://juejin.im/post/5cb30243e51d456e431ada29 )
* [【前端词典】学习 Vue 源码的必要知识储备（完成）]( https://juejin.im/post/5ce5565d6fb9a07ed2244513 )
* 
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