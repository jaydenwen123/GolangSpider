# VUE防抖与节流的最佳解决方案——函数式组件 #

## 前言 ##

> 
> 
> 
> 有echarts使用经验的同学可能遇到过这样的场景，在window.onresize事件回调里触发echartsBox.resize()方法来达到重绘的目的，resize事件是连续触发的这意味着echarts实例会连续的重绘这是非常耗性能的。还有一个常见的场景在input标签的input事件里请求后端接口，input事件也是连续触发的，假设我输入了“12”就会请求两次接口参数分别是“1”和“12”，比浪费网络资源更要命的是如果参数为“1”的请求返回数据的时间晚于参数为“12”的接口，那么我们得到的数据是和期望不符的。当然基于axios可以做很多封装可以取消上一个请求或者通过拦截做处理，但还是从防抖入手比较简单。
> 
> 
> 

## 防抖和节流到底是啥 ##

### 函数防抖（debounce） ###

解释：当持续触发某事件时，一定时间间隔内没有再触发事件时，事件处理函数才会执行一次，如果设定的时间间隔到来之前，又一次触发了事件，就重新开始延时。
案例：持续触发scroll事件时，并不立即执行handle函数，当1000毫秒内没有触发scroll事件时，才会延时触发一次handle函数。

` function debounce(fn, wait ) { let timeout = null return function () { if (timeout !== null) clearTimeout(timeout) timeout = set Timeout(fn, wait ); } } function handle () { console.log(Math.random()) } window.addEventListener( 'scroll' , debounce(handle, 1000)) 复制代码`

**addEventListener的第二个参数实际上是debounce函数里return回的方法， ` let timeout = null` 这行代码只在addEventListener的时候执行了一次 触发事件的时候不会执行，那么每次触发scroll事件的时候都会清除上次的延时器同时记录一个新的延时器，当scroll事件停止触发后最后一次记录的延时器不会被清除可以延时执行，这是debounce函数的原理**

### 函数节流（throttle） ###

解释：当持续触发事件时，有规律的每隔一个时间间隔执行一次事件处理函数。
案例：持续触发scroll事件时，并不立即执行handle函数，每隔1000毫秒才会执行一次handle函数。

` function throttle(fn, delay) { var prev = Date.now() return function () { var now = Date.now() if (now - prev > delay) { fn() prev = Date.now() } } } function handle () { console.log(Math.random()) } window.addEventListener( 'scroll' , throttle(handle, 1000)) 复制代码`

**原理和防抖类似，每次执行fn函数都会更新prev用来记录本次执行的时间，下一次事件触发时判断时间间隔是否到达预先的设定，重复上述操作。**

防抖和节流都可以用于 mousemove、scroll、resize、input等事件，他们的区别在于防抖只会在连续的事件周期结束时执行一次，而节流会在事件周期内按间隔时间有规律的执行多次。

![](https://user-gold-cdn.xitu.io/2019/5/16/16ac0fba56ec1726?imageslim)

## 在vue中的实践 ##

在vue中实现防抖无非下面这两种方法

* 封装utils工具
* 封装组件

### 封装utils工具 ###

把上面的案例改造一下就能封装一个简单的utils工具

` utils.js let timeout = null function debounce(fn, wait ) { if (timeout !== null) clearTimeout(timeout) timeout = set Timeout(fn, wait ) } export default debounce app.js <input type = "text" @input= "debounceInput( $event )" > import debounce from './utils' export default { methods: { debounceInput(E){ debounce(() => { console.log(E.target.value) }, 1000) } } } 复制代码`

### 封装组件 ###

至于组件的封装我们要用到 ` $listeners、$attrs` 这两个属性，他俩都是vue2.4新增的内容，官网的介绍比较晦涩，我们来看他俩到底是干啥的：

` $listeners:` 父组件在绑定子组件的时候会在子组件上绑定很多属性，然后在子组件里通过props注册使用，那么没有被props注册的就会放在 ` $listeners` 里，当然不包括class和style，并且可以通过 ` v-bind="$attrs"` 传入子组件的内部组件。

` $listeners:` 父组件在子组件上绑定的不含.native修饰器的事件会放在 ` $listeners里` ，它可以通过 ` v-on="$listeners"` 传入内部组件。

简单来说 ` $listeners、$attrs` 他俩是做属性和事件的承接，这在对组件做二次封装的时候非常有用。

我们以element-ui的el-input组件为例封装一个带防抖的debounce-input组件

` debounce-input.vue <template> <el-input v-bind= " $attrs " @input= "debounceInput" /> </template> <script> export default { data () { return { timeout: null } }, methods: { debounceInput(value){ if (this.timeout !== null) clearTimeout(this.timeout) this.timeout = set Timeout(() => { this. $emit ( 'input' , value) }, 1000) } } } </script> app.vue <template> <debounce-input placeholder= "防抖" prefix-icon= "el-icon-search" @input= "inputEve" ></debounce-input> </template> <script> import debounceInput from './debounce-input' export default { methods: { inputEve(value){ console.log(value) } }, components: { debounceInput } } </script> 复制代码`

上面组件的封装用了$attrs，虽然不需要开发者关注属性的传递，但是在使用上还是不方便的，因为把el-input封装在了内部这样对样式的限定也比较局限。有接触过react高阶组件的同学可能有了解，react高阶组件本质上是一个函数通过包裹被传入的React组件，经过一系列处理，最终返回一个相对增强的React组件。那么在vue中可以借鉴这种思路吗，我们来了解一下vue的函数式组件。

## 关于vue函数式组件 ##

### 什么是函数式组件？ ###

函数式组件是指用一个Function来渲染一个vue组件，这个组件只接受一些 prop，我们可以将这类组件标记为 functional，这意味着它无状态 (没有响应式数据)，也没有实例 (没有this上下文)。

一个函数式组件大概向下面这样：

` export default () => { functional: true , props: { // Props 是可选的 }, // 为了弥补缺少的实例, 提供第二个参数作为上下文 render: function (createElement, context) { return vNode } } 复制代码`
> 
> 
> 
> 
> 注意：在 2.3.0 之前的版本中，如果一个函数式组件想要接收 prop，则 props 选项是必须的。在 2.3.0 或以上的版本中，你可以省略
> props 选项，所有组件上的特性都会被自动隐式解析为 prop。但是你一旦注册了 prop 那么只有被注册的 prop 会出现在
> context.prop 里。
> 
> 

render函数的第二个参数context用来代替上下文this他是一个包含如下字段的对象：

* props：提供所有 prop 的对象
* children: VNode 子节点的数组
* slots: 一个函数，返回了包含所有插槽的对象
* scopedSlots: (2.6.0+) 一个暴露传入的作用域插槽的对象。也以函数形式暴露普通插槽。
* data：传递给组件的整个数据对象，作为 createElement 的第二个参数传入组件
* parent：对父组件的引用
* listeners: (2.3.0+) 一个包含了所有父组件为当前组件注册的事件监听器的对象。这是 data.on 的一个别名。
* injections: (2.3.0+) 如果使用了 inject 选项，则该对象包含了应当被注入的属性。

### vm.$slots API 里面是什么 ###

slots用来访问被插槽分发的内容。每个具名插槽 有其相应的属性 (例如：v-slot:foo 中的内容将会在 vm.$slots.foo 中被找到)。default 属性包括了所有没有被包含在具名插槽中的节点，或 v-slot:default 的内容。

### slots() 和 children 对比 ###

你可能想知道为什么同时需要 slots() 和 children。slots().default 不是和 children 类似的吗？在一些场景中，是这样——但如果是如下的带有子节点的函数式组件呢？

` <my-functional-component> <p v-slot:foo> first </p> <p>second</p> </my-functional-component> 复制代码`

对于这个组件，children 会给你两个段落标签，而 slots().default 只会传递第二个匿名段落标签，slots().foo 会传递第一个具名段落标签。同时拥有 children 和 slots()，因此你可以选择让组件感知某个插槽机制，还是简单地通过传递 children，移交给其它组件去处理。

### 一个函数式组件的使用场景 ###

假设有一个a组件，引入了 a1,a2,a3 三个组件，a组件的父组件给a组件传入了一个type属性根据type的值a组件来决定显示 a1,a2,a3 中的那个组件。这样的场景a组件用函数式组件是非常方便的。那么为什么要用函数式组件呢？一句话：渲染开销低，因为函数式组件只是函数。

## 用函数式组件的方式来实现防抖 ##

因为业务关系该防抖组件的封装同时支持 input、button、el-input、el-button 的使用，如果是input类组件对input事件做防抖处理，如果是button类组件对click事件做防抖处理。

` const debounce = (fun, delay = 500, before) => { let timer = null return (params) => { timer && window.clearTimeout(timer) before && before(params) timer = window.setTimeout(() => { // click事件fun是Function input事件fun是Array if (!Array.isArray(fun)) { fun = [fun] } for ( let i in fun) { fun[i](params) } timer = null }, parseInt(delay)) } } export default { name: 'Debounce' , functional: true , // 静态组件 当不声明functional时该组件同样拥有上下文以及生命周期函数 render(createElement, context) { const before = context.props.before const time = context.props.time const vnodeList = context.slots().default if (vnodeList === undefined){ console.warn( '<debounce> 组件必须要有子元素' ) return null } const vnode = vnodeList[0] || null // 获取子元素虚拟dom if (vnode.tag === 'input' ) { const defaultFun = vnode.data.on.input const debounceFun = debounce(defaultFun, time, before) // 获取节流函数 vnode.data.on.input = debounceFun } else if (vnode.tag === 'button' ) { const defaultFun = vnode.data.on.click const debounceFun = debounce(defaultFun, time, before) // 获取节流函数 vnode.data.on.click = debounceFun } else if (vnode.componentOptions && vnode.componentOptions.tag === 'el-input' ) { const defaultFun = vnode.componentOptions.listeners.input const debounceFun = debounce(defaultFun, time, before) // 获取节流函数 vnode.componentOptions.listeners.input = debounceFun } else if (vnode.componentOptions && vnode.componentOptions.tag === 'el-button' ) { const defaultFun = vnode.componentOptions.listeners.click const debounceFun = debounce(defaultFun, time, before) // 获取节流函数 vnode.componentOptions.listeners.click = debounceFun } else { console.warn( '<debounce> 组件内只能出现下面组件的任意一个且唯一 el-button、el-input、button、input' ) return vnode } return vnode } } 复制代码` ` <template> <debounce time= "300" :before= "beforeFun" > <input type = "text" v-model= "inpModel" @input= "inputChange" /> </debounce> </template> <script> import debounce from './debounce' export default { data () { return { inpModel: 1 } }, methods: { inputChange(e){ console.log(e.target.value, '防抖' ) }, beforeFun(e){ console.log(e.target.value, '不防抖' ) } }, components: { debounce } } </script> 复制代码`

原理也很简单就是在vNode中拦截on下面的click、input事件做防抖处理，这样在使用上就非常简单了。

## 自定义指令 directive ##

我们来思考一个问题，函数式组件封装防抖的关节是获取vNode，那么我们通过自定义指令同样可以拿到vNode，甚至还可以得到原生的Dom，这样用自定义指令来处理会更加方便。。。。。。

## 相关阅读 ##

` $attrs or $listeners` [cn.vuejs.org/v2/api/#vm-…]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23vm-attrs )
` 函数式组件` [cn.vuejs.org/v2/guide/re…]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Frender-function.html%23%25E5%2587%25BD%25E6%2595%25B0%25E5%25BC%258F%25E7%25BB%2584%25E4%25BB%25B6 )
` 自定义指令` [cn.vuejs.org/v2/guide/cu…]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcustom-directive.html )