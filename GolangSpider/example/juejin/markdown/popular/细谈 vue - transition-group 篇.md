# 细谈 vue - transition-group 篇 #

本篇文章是细谈 vue 系列的第四篇，按理说这篇文章是上篇 [《细谈 vue - transition 篇》]( https://juejin.im/post/5cf411d8e51d4550a629b222 ) 中的一个单独的大章节。然鹅，上篇文章篇幅过长，所以不得已将其单独拎出来写成一篇了。对该系列以前的文章感兴趣的可以点击以下链接进行传送

* [《细谈 vue 核心- vdom 篇》]( https://juejin.im/post/5cab347fe51d456e7a303b3d )
* [《细谈 vue - slot 篇》]( https://juejin.im/post/5cced0096fb9a032426510ad )
* [《细谈 vue - transition 篇》]( https://juejin.im/post/5cf411d8e51d4550a629b222 )

书接上文，上篇文章我们主要介绍了 ` <transition>` 组件对 ` props` 和 ` vnode hooks` 的 ` 输入 => 输出` 处理设计，它针对单一元素的 ` enter` 以及 ` leave` 阶段进行了过渡效果的封装处理，使得我们只需关注 ` css` 和 ` js` 钩子函数的业务实现即可。

但是我们在实际开发中，却终究难逃多个元素都需要进行使用过渡效果进行展示，很显然， ` <transition>` 组件并不能实现我的业务需求。这个时候， ` vue` 内部封装了 ` <transition-group>` 这么一个内助组件来满足我们的需要，它很好的帮助我们实现了列表的过渡效果。

## 一、举个例子 ##

老样子，直接先上一个官方的例子

` < template > < div id = "list-demo" > < button v-on:click = "add" > Add </ button > < button v-on:click = "remove" > Remove </ button > < transition-group name = "list" tag = "p" > < span v-for = "item in items" v-bind:key = "item" class = "list-item" > {{ item }} </ span > </ transition-group > </ div > </ template > < script > export default { name : 'home' , data () { return { items : [ 1 , 2 , 3 , 4 , 5 , 6 , 7 , 8 , 9 ], nextNum : 10 } }, methods : { randomIndex : function ( ) { return Math.floor( Math.random() * this.items.length) }, add : function ( ) { this.items.splice( this.randomIndex(), 0 , this.nextNum++) }, remove : function ( ) { this.items.splice( this.randomIndex(), 1 ) } } } </ script > < style lang = "scss" >.list-item { display : inline-block; margin-right : 10px ; }.list-enter-active ,.list-leave-active { transition : all 1s ; }.list-enter ,.list-leave-to { opacity : 0 ; transform : translateY (30px); } </ style > 复制代码`

效果如下图

![](https://user-gold-cdn.xitu.io/2019/6/5/16b274f7ca449abe?imageslim)

接下来，我将带着大家一起探究一下 ` <transition-group>` 组件的设计

## 二、transition-group 实现 ##

和 ` <transition>` 组件相比， ` <transition>` 是一个抽象组件，且只对单个元素生效。而 ` <transition-group>` 组件实现了列表的过渡，并且它会渲染一个真实的元素节点。

但他们的设计理念却是一致的，同样会给我们提供一个 ` props` 和一系列钩子函数给我们当做 ` 输入` 的接口，内部进行 ` 输入 => 输出` 的转换或者说绑定处理

` export default { props, beforeMount () { // ... }, render (h: Function ) { // ... }, updated () { // ... }, methods : { // ... } } 复制代码`

### 1、props & other import ###

` <transition-group>` 的 ` props` 和 ` <transition>` 的 ` props` 基本一致，只是多了一个 ` tag` 和 ` moveClass` 属性，删除了 ` mode` 属性

` // props import { transitionProps, extractTransitionData } from './transition' const props = extend({ tag : String , moveClass : String }, transitionProps) delete props.mode // other import import { warn, extend } from 'core/util/index' import { addClass, removeClass } from '../class-util' import { setActiveInstance } from 'core/instance/lifecycle' import { hasTransition, getTransitionInfo, transitionEndEvent, addTransitionClass, removeTransitionClass } from '../transition-util' 复制代码`

### 2、render ###

首先，我们需要定义一系列变量，方便后续的操作

* ` tag` ：从上面设计的整体脉络我们能看到， ` <transition-group>` 并没有 ` abstract` 属性，即它将渲染一个真实节点，那么节点 ` tag` 则是必须的，其默认值为 ` span` 。
* ` map` ：创建一个空对象
* ` prevChildren` ：用来存储上一次的子节点
* ` rawChildren` ：获取 ` <transition-group>` 包裹的子节点
* ` children` ：用来存储当前的子节点
* ` transitionData` ：获取组件上的渲染数据

` const tag: string = this.tag || this.$vnode.data.tag || 'span' const map: Object = Object.create( null ) const prevChildren: Array <VNode> = this.prevChildren = this.children const rawChildren: Array <VNode> = this.$slots.default || [] const children: Array <VNode> = this.children = [] const transitionData: Object = extractTransitionData( this ) 复制代码`

紧接着是对节点遍历的操作，这里主要对列表中每个节点进行过渡动画的绑定

* 对 ` rawChildren` 进行遍历，并将每个 ` vnode` 节点取出；
* 若节点存在含有 **__vlist** 字符的 ` key` ，则将 ` vnode` 丢到 ` children` 中；
* 随即将提取出来的过渡数据 ` transitionData` 添加到 ` vnode.data.transition` 上，这样便能实现列表中单个元素的过渡动画

` for ( let i = 0 ; i < rawChildren.length; i++) { const c: VNode = rawChildren[i] if (c.tag) { if (c.key != null && String (c.key).indexOf( '__vlist' ) !== 0 ) { children.push(c) map[c.key] = c ;(c.data || (c.data = {})).transition = transitionData } else if (process.env.NODE_ENV !== 'production' ) { const opts: ?VNodeComponentOptions = c.componentOptions const name: string = opts ? (opts.Ctor.options.name || opts.tag || '' ) : c.tag warn( `<transition-group> children must be keyed: < ${name} >` ) } } } 复制代码`

随后对 ` prevChildren` 进行处理

* 如果 ` prevChildren` 存在，则对其进行遍历，将 ` transitionData` 赋值给 ` vnode.data.transition` ，如此之后，当 ` vnode` 子节点 ` enter` 和 ` leave` 阶段存在过渡动画的时候，则会执行对应的过渡动画
* 随即调用原生的 ` getBoundingClientRect` 获取元素的位置信息，将其记录到 ` vnode.data.pos` 中
* 然后判断 ` map` 中是否存在 ` vnode.key` ，若存在，则将 ` vnode` 放到 ` kept` 中，否则丢到 ` removed` 队列中
* 最后将渲染后的元素放到 ` this.kept` 中， ` this.removed` 则用来记录被移除掉的节点

` if (prevChildren) { const kept: Array <VNode> = [] const removed: Array <VNode> = [] for ( let i = 0 ; i < prevChildren.length; i++) { const c: VNode = prevChildren[i] c.data.transition = transitionData c.data.pos = c.elm.getBoundingClientRect() if (map[c.key]) { kept.push(c) } else { removed.push(c) } } this.kept = h(tag, null , kept) this.removed = removed } 复制代码`

最后 ` <transition-group>` 进行渲染

` return h(tag, null , children) 复制代码`

### 3、update & methods ###

上面我们已经在 ` render` 阶段对列表中的每个元素绑定好了 ` transition` 相关的过渡效果，接下来就是每个元素动态变更时，整个列表进行 ` update` 时候的动态过渡了。那具体这块又是如何操作的呢？接下来我们就捋捋这块的逻辑

#### i. 是否需要进行 move 过渡 ####

* 首先在 ` update` 钩子函数里面，会先获取上一次的子节点 ` prevChildren` 和 ` moveClass` ；随后判断 ` children` 是否存在以及 ` children` 是否 **has move** ，若 ` children` 不存在，或者 ` children` 没有 ` move` 状态，那么也没有必要继续进行 ` update` 的 ` move` 过渡了，直接 ` return` 即可

` const children: Array <VNode> = this.prevChildren const moveClass: string = this.moveClass || (( this.name || 'v' ) + '-move' ) if (!children.length || ! this.hasMove(children[ 0 ].elm, moveClass)) { return } 复制代码`

* ` hasMove()` ：该方法主要用来判断 ` el` 节点是否有 ` move` 的状态。
* 当前置 ` return` 条件不符合的情况下，它会先克隆一个 DOM 节点，然后为了避免元素内部已经有了 css 过渡，所以会移除掉克隆节点上的所有的 ` transitionClasses`
* 紧接着，对克隆节点重新加上 ` moveClass` ，并将其 ` display` 设为 ` none` ，然后添加到 ` this.$el` 上
* 接下来通过 ` getTransitionInfo` 获取它的 ` transition` 相关的信息，然后从 ` this.$el` 上将其移除。这个时候我们已经获取到了节点是否有 ` transform` 的信息了

` export const hasTransition = inBrowser && !isIE9 hasMove (el: any, moveClass : string): boolean { // 若不在浏览器中，或者浏览器不支持 transition，直接返回 false 即可 if (!hasTransition) { return false } // 若当前实例上下文的有 _hasMove，直接返回 _hasMove 的值即可 if ( this._hasMove) { return this._hasMove } const clone: HTMLElement = el.cloneNode() if (el._transitionClasses) { el._transitionClasses.forEach( ( cls: string ) => { removeClass(clone, cls) }) } addClass(clone, moveClass) clone.style.display = 'none' this.$el.appendChild(clone) const info: Object = getTransitionInfo(clone) this.$el.removeChild(clone) return ( this._hasMove = info.hasTransform) } 复制代码`

#### ii. move 过渡实现 ####

* 然后对子节点进行一波预处理，这里对子节点的处理使用了三次循环，主要是为了避免每次循环对 DOM 的读写变的混乱，有助于防止布局混乱

` children.forEach(callPendingCbs) children.forEach(recordPosition) children.forEach(applyTranslation) 复制代码`

三个函数的处理分别如下

* ` callPendingCbs()` ：判断每个节点前一帧的过渡动画是否执行完毕，如果没有执行完，则提前执行 ` _moveCb()` 和 ` _enterCb()`
* ` recordPosition()` ：记录每个节点的新位置
* ` applyTranslation()` ：分别获取节点新旧位置，并计算差值，若存在差值，则通过设置节点的 ` transform` 属性将需要移动的节点位置偏移到之前的位置，为列表 ` move` 做准备
` function callPendingCbs ( c: VNode ) { if (c.elm._moveCb) { c.elm._moveCb() } if (c.elm._enterCb) { c.elm._enterCb() } } function recordPosition ( c: VNode ) { c.data.newPos = c.elm.getBoundingClientRect() } function applyTranslation ( c: VNode ) { const oldPos = c.data.pos const newPos = c.data.newPos const dx = oldPos.left - newPos.left const dy = oldPos.top - newPos.top if (dx || dy) { c.data.moved = true const s = c.elm.style s.transform = s.WebkitTransform = `translate( ${dx} px, ${dy} px)` s.transitionDuration = '0s' } } 复制代码`

* 紧接着，对子元素进行遍历实现 ` move` 过渡。遍历前会通过获取 ` document.body.offsetHeight` ，从而发生计算，触发回流，让浏览器进行重绘
* 然后开始对 ` children` 进行遍历，期间若 ` vnode.data.moved` 为 ` true` ，则执行 ` addTransitionClass` 为子节点加上 ` moveClass` ，并将其 ` style.transform` 属性清空，由于我们在子节点预处理中已经将子节点偏移到了之前的旧位置，所以此时它会从旧位置过渡偏移到当前位置，这就是我们要的 ` move` 过渡的效果
* 最后会为节点加上 ` transitionend` 过渡结束的监听事件，在事件里做一些清理的操作

` this._reflow = document.body.offsetHeight children.forEach( ( c: VNode ) => { if (c.data.moved) { const el: any = c.elm const s: any = el.style addTransitionClass(el, moveClass) s.transform = s.WebkitTransform = s.transitionDuration = '' el.addEventListener(transitionEndEvent, el._moveCb = function cb ( e ) { if (e && e.target !== el) { return } if (!e || /transform$/.test(e.propertyName)) { el.removeEventListener(transitionEndEvent, cb) el._moveCb = null removeTransitionClass(el, moveClass) } }) } }) 复制代码`

注：浏览器回流触发条件我稍微做个总结，比如浏览器窗口改变、计算样式、对 DOM 进行元素的添加或者删除、改变元素 class 等

> 
> * 添加或者删除可见的DOM元素
> * 元素位置改变
> * 元素尺寸改变 —— 边距、填充、边框、宽度和高度
> * 内容变化，比如用户在 input 框中输入文字，文本或者图片大小改变而引起的计算值宽度和高度改变
> * 页面渲染初始化
> * 浏览器窗口尺寸改变 —— resize 事件发生时
> * 计算 offsetWidth 和 offsetHeight 属性
> * 设置 style 属性的值
> 

### 4、beforeMount ###

由于 ` VDOM` 在节点 ` diff` 更新的时候是不能保证被移除元素它的一个相对位置。所以这里需要在 ` beforeMount` 钩子函数里面对 ` update` 渲染逻辑重写，来达到我们想要的效果

* 首先获取实例本身的 ` update` 方法，进行缓存
* 从上面我们知道 ` this.kept` 是缓存的上次的节点，并且里面的节点增加了一些 ` transition` 过渡属性。这里首先通过 ` setActiveInstance` 缓存好当前实例，随即对 ` vnode` 进行 ` __patch__` 操作并移除需要被移除掉的 ` vnode` ，然后执行 ` restoreActiveInstance` 将其实例指向恢复
* 随后将 ` this.kept` 赋值给 ` this._vnode` ，使其触发过渡
* 最后执行缓存的 ` update` 渲染节点

` beforeMount () { const update = this._update this._update = ( vnode, hydrating ) => { const restoreActiveInstance = setActiveInstance( this ) // force removing pass this.__patch__( this._vnode, this.kept, false , // hydrating true // removeOnly (!important, avoids unnecessary moves) ) this._vnode = this.kept restoreActiveInstance() update.call( this , vnode, hydrating) } } 复制代码`

* ` setActiveInstance`

` export let activeInstance: any = null export function setActiveInstance ( vm: Component ) { const prevActiveInstance = activeInstance activeInstance = vm return () => { activeInstance = prevActiveInstance } } 复制代码`

## 最后 ##

文章到这就已经差不多了，对 ` transition` 相关的内置组件 ` <transition>` 以及 ` <transition-group>` 的解析也已经是结束了。不同的组件类型，一个抽象组件、一个则会渲染实际节点元素，想要做的事情却是一样的，初始化给用户的 ` 输入` 接口， ` 输入` 后即可得到 ` 输出` 的过渡效果。

前端交流群：731175396，热烈欢迎各位妹纸，汉纸踊跃加入

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2751bde68e552?imageView2/0/w/1280/h/960/ignore-error/1)