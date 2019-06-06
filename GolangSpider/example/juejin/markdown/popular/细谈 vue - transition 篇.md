# 细谈 vue - transition 篇 #

本篇文章是细谈 vue 系列的第三篇，这篇文章主要会介绍一下 ` vue` 的内置组件 ` transition` 。前几篇链接如下

* 

[《细谈 vue 核心- vdom 篇》]( https://juejin.im/post/5cab347fe51d456e7a303b3d )

* 

[《细谈 vue - slot 篇》]( https://juejin.im/post/5cced0096fb9a032426510ad )

开始之前，我们先看下官方对 ` transition` 的定义

* 自动嗅探目标元素是否使用了 CSS 过渡或动画，如果使用，会在合适的时机添加/移除 CSS 过渡 class。
* 如果过渡组件设置了 [JavaScript 钩子函数]( https://link.juejin.im?target=https%3A%2F%2Fvue.docschina.org%2Fv2%2Fguide%2Ftransitions.html%23JavaScript-Hooks ) ，这些钩子函数将在合适的时机调用。
* 如果没有检测到 CSS 过渡/动画，并且也没有设置 JavaScript 钩子函数，插入和/或删除 DOM 的操作会在下一帧中立即执行

## 一、抽象组件 ##

在 ` vue` 中，对于 ` options` 有一条属性叫 ` abstract` ，类型为 ` boolean` ，他代表组件是否为抽象组件。但找寻了有关 ` options` 的类型定义，却只找到这么一条

` static options: Object ; 复制代码`

但实际上， ` vue` 还是有对 ` abstract` 属性进行处理的，在 ` lifecycle.js` 的 ` initLifecycle()` 方法中有这么一段逻辑

` let parent = options.parent if (parent && !options.abstract) { while (parent.$options.abstract && parent.$parent) { parent = parent.$parent } parent.$children.push(vm) } vm.$parent = parent 复制代码`

从这段逻辑我们能看出， ` vue` 会一层一层往上找父节点是否拥有 ` abstract` 属性，找到之后则直接将 ` vm` 丢到其父节点的 ` $children` 中，其本身的父子关系是直接被忽略的

然后在 ` create-component.js` 的 ` createComponent()` 方法中对其处理如下

` if (isTrue(Ctor.options.abstract)) { // abstract components do not keep anything // other than props & listeners & slot // work around flow const slot = data.slot data = {} if (slot) { data.slot = slot } } 复制代码`

## 二、举个例子 ##

分析之前，先看个官方 ` transition` 的例子

` < template > < div id = "example" > < button @ click = "show = !show" > Toggle render </ button > < transition name = "slide-fade" > < p v-if = "show" > hello </ p > </ transition > </ div > </ template > < script > export default { data () { return { show : true } } } </ script > < style lang = "scss" >.slide-fade-enter-active { transition : all . 3s ease; }.slide-fade-leave-active { transition : all . 8s cubic-bezier (1.0, 0.5, 0.8, 1.0); }.slide-fade-enter ,.slide-fade-leave-to { transform : translateX (10px); opacity : 0 ; } </ style > 复制代码`

当点击按钮切换显示状态时，被 ` transition` 包裹的元素会有一个 css 过渡效果。接下来，我们将分析具体它是如何做到这种效果的。

## 三、transition 实现 ##

从上面的用法中，我们应该能感知出，对于 ` <transition>` ， ` vue` 给我们提供了一整套 css 以及 js 的钩子，这些钩子本身都已经帮我定义好并绑定在实例上面，剩下的就是需要我们自己去重写 css 或者 js 钩子函数。其实这和 ` vue` 本身的 ` hooks` 实现非常类似，都会在不同的钩子帮我们做好对应的不同的事情。

这种编程思维非常像函数式编程中的思想，或者说它就是。

> 
> 
> 
> 函数式编程关心数据的映射，函数式编程中的 lambda
> 可以看成是两个类型之间的关系，一个输入类型和一个输出类型。你给它一个输入类型的值，则可以得到一个输出类型的值。
> 
> 

其实这个反映到这块，在其不同的钩子里面进行不同程度的重写，即可得到不同的输出结果，当然这里的输出其实就是过渡效果，而这种方式不恰好是 ` 输入 => 输出` 的方式吗？接下来我们就讲一下这块的具体实现。

` <transtion>` 组件的实现是 ` export` 出 一个对象，从这里我们能看出它会将预先设定好的 ` props` 绑定到 ` transition` 上，而后我们只需对其中的点进行输入即可得到对应的样式输出，具体如何进行的过程，就不需要我们去考虑了

` export default { name : 'transition' , props : transitionProps, abstract : true , render (h: Function ) { // 具体 render 等会再看 } } 复制代码`

其中支持的 ` props` 如下，你可以对 ` enter` 和 ` leave` 的样式进行任意形式的重写

` export const transitionProps = { name : String , appear : Boolean , css : Boolean , mode : String , type : String , enterClass : String , leaveClass : String , enterToClass : String , leaveToClass : String , enterActiveClass : String , leaveActiveClass : String , appearClass : String , appearActiveClass : String , appearToClass : String , duration : [ Number , String , Object ] } 复制代码`

接下来我们来看看 ` render` 里面的内容是如何对我们预先绑定好的东西进行 ` 输入 => 输出` 的一个转换的

* ` children` 处理逻辑：首先从默认插槽中获取到 ` <transition>` 包裹的子节点，随后对其进行 ` filter` 过滤，将文本节点以及空格给过滤掉。如果不存在子节点则直接 ` return` ，如果子节点为多个，则报错，因为 ` <transition>` 组件只能有一个子节点

` let children: any = this.$slots.default if (!children) { return } children = children.filter(isNotTextNode) if (!children.length) { return } if (process.env.NODE_ENV !== 'production' && children.length > 1 ) { warn( '<transition> can only be used on a single element. Use ' + '<transition-group> for lists.' , this.$parent ) } 复制代码`

* ` model` 处理逻辑：判断 ` mode` 是否为 ` in-out` 和 ` out-in` 两种模式，如果不是，直接报错

` const mode: string = this.mode if ( process.env.NODE_ENV !== 'production' && mode && mode !== 'in-out' && mode !== 'out-in' ) { warn( 'invalid <transition> mode: ' + mode, this.$parent ) } 复制代码`

* ` rawChild & child` 处理逻辑： ` rawChild` 为 ` <transition>` 组件包裹的第一个 ` vnode` 子节点，紧接着判断其父容器是否也为 ` transition` ，如果是则直接返回 ` rawChild` ；随后执行 ` getRealChild()` 方法获取组件的真实子节点，不存在则返回 ` rawChild`

` const rawChild: VNode = children[ 0 ] if (hasParentTransition( this.$vnode)) { return rawChild } const child: ?VNode = getRealChild(rawChild) if (!child) { return rawChild } 复制代码`

* ` hasParentTransition()` : 往上一层一层寻找父节点是否也拥有 ` transition` 属性

` function hasParentTransition ( vnode: VNode ): ? boolean { while ((vnode = vnode.parent)) { if (vnode.data.transition) { return true } } } 复制代码`

* ` getRealChild()` ：递归获取第一个非抽象子节点并返回

` function getRealChild ( vnode: ?VNode ): ? VNode { const compOptions: ?VNodeComponentOptions = vnode && vnode.componentOptions if (compOptions && compOptions.Ctor.options.abstract) { return getRealChild(getFirstComponentChild(compOptions.children)) } else { return vnode } } 复制代码`

* 紧接着的就是对 ` child.key` 和 ` data` 的处理了。首先通过 ` id` 和一系列条件对 ` child.key` 进行赋值操作，然后使用 ` extractTransitionData` 从组件实例上获取过渡需要的 ` data` 数据

` const id: string = `__transition- ${ this._uid} -` // 使用 this._uid 进行拼接 child.key = child.key == null // 若 child.key 为 null ? child.isComment // child 为注释节点 ? id + 'comment' : id + child.tag : isPrimitive(child.key) // 若 child.key 为 原始值 ? ( String (child.key).indexOf(id) === 0 ? child.key : id + child.key) : child.key // 获取过渡需要的数据 const data: Object = (child.data || (child.data = {})).transition = extractTransitionData( this ) if (child.data.directives && child.data.directives.some(isVShowDirective)) { child.data.show = true } 复制代码`

* ` extractTransitionData()` ：这块整体就是对输入值 ` props` 以及内部 ` events` 的一个转换处理，该方法参数为 ` this` ，即当前组件。首先遍历当前组件的 ` options.propsData` 并赋值给 ` data` ，然后遍历父组件的事件并将其赋值给 ` data` 。

` export function extractTransitionData ( comp: Component ): Object { const data = {} const options: ComponentOptions = comp.$options // props for ( const key in options.propsData) { data[key] = comp[key] } // events. // extract listeners and pass them directly to the transition methods const listeners: ? Object = options._parentListeners for ( const key in listeners) { data[camelize(key)] = listeners[key] } return data } 复制代码`

根据目前逻辑，我上面所举的例子，获取到的 ` child.data` 如下

` { transition : { name : 'slide-fade' } } 复制代码`

* 最后则是对新旧 ` child` 进行比较，并对部分钩子函数进行 ` hook merge` 等操作

` const oldRawChild: VNode = this._vnode const oldChild: VNode = getRealChild(oldRawChild) if ( oldChild && oldChild.data && !isSameChild(child, oldChild) && !isAsyncPlaceholder(oldChild) && !(oldChild.componentInstance && oldChild.componentInstance._vnode.isComment) ) { // 将 oldData 更为为当前的 data const oldData: Object = oldChild.data.transition = extend({}, data) // 当 transition mode 为 out-in 的情况，返回一个 vnode 占位，执行完 watch 进行更新 if (mode === 'out-in' ) { this._leaving = true mergeVNodeHook(oldData, 'afterLeave' , () => { this._leaving = false this.$forceUpdate() }) return placeholder(h, rawChild) // 当 transition mode 为 in-out 的情况 } else if (mode === 'in-out' ) { if (isAsyncPlaceholder(child)) { return oldRawChild } let delayedLeave const performLeave = () => { delayedLeave() } mergeVNodeHook(data, 'afterEnter' , performLeave) mergeVNodeHook(data, 'enterCancelled' , performLeave) mergeVNodeHook(oldData, 'delayLeave' , leave => { delayedLeave = leave }) } } 复制代码`

这里就看下 ` mergeVNodeHook()` 、 ` placeholder()` 和 ` $forceUpdate()` 的逻辑，这里就提一下 ` mergeVNodeHook` 的逻辑：它会将 ` hook` 函数合并到 ` def.data.hook[hookKey]` 中，生成一个新的 ` invoker`

` // mergeVNodeHook export function mergeVNodeHook ( def: Object, hookKey: string, hook: Function ) { if (def instanceof VNode) { def = def.data.hook || (def.data.hook = {}) } let invoker // 获取已有 hook 赋值给 oldHook const oldHook = def[hookKey] function wrappedHook ( ) { hook.apply( this , arguments ) // 删除合并的钩子确保其只被调用一次，这样能防止内存泄漏 remove(invoker.fns, wrappedHook) } // 如果 oldHook 不存在，则直接创建一个 invoker if (isUndef(oldHook)) { invoker = createFnInvoker([wrappedHook]) } else { // oldHook 已经存在，则将 invoker 赋值为 oldHook if (isDef(oldHook.fns) && isTrue(oldHook.merged)) { invoker = oldHook invoker.fns.push(wrappedHook) } else { // 现有的普通钩子 invoker = createFnInvoker([oldHook, wrappedHook]) } } invoker.merged = true def[hookKey] = invoker } // placeholder function placeholder ( h: Function, rawChild: VNode ): ? VNode { if ( /\d-keep-alive$/.test(rawChild.tag)) { return h( 'keep-alive' , { props : rawChild.componentOptions.propsData }) } } // $forceUpdate Vue.prototype.$forceUpdate = function ( ) { const vm: Component = this if (vm._watcher) { vm._watcher.update() } } 复制代码`

## 四、enter & leave ##

介绍完 ` <transition>` 组件的实现，我们了解到其在 ` render` 阶段对于 ` 输入值` ` props` 以及部分 js 钩子进行 ` 输出` 处理的过程。

或者我们可以这么理解， ` <transition>` 在 ` render` 阶段会获取节点上的数据、在不同的 ` mode` 下绑定了对应的钩子函数以及其每个钩子需要用到的 ` data` 数据、且同时也会返回了 ` rawChild` ` vnode` 节点。

但是，目前为止它却并没有在动画这块进行任何设计。so，接下来，我们将详细谈一下 ` transition` 在 ` enter` 以及 ` leave` 这块是如何进行 ` 输入 => 输出` 的。

首先在 ` src/platforms/web/modules/transition.js` 中，有这么一段代码

` function _enter ( _: any, vnode: VNodeWithData ) { if (vnode.data.show !== true ) { enter(vnode) } } export default inBrowser ? { create : _enter, activate : _enter, remove (vnode: VNode, rm : Function ) { if (vnode.data.show !== true ) { leave(vnode, rm) } else { rm() } } } : {} 复制代码`

从上面的代码我们能看出，在动画处理这块，它设定了两个时机，分别是：

* 在 ` create` 和 ` activate` 的时候执行 ` enter()`
* ` remove` 的时候执行 ` leave()`

### 1、enter ###

分析前，我们看下 ` enter()` 设计脉络

` export function enter ( vnode: VNodeWithData, toggleDisplay: ?( ) => void ) { const el: any = vnode.elm // 如果 el 中存在 _leaveCb，则立即执行 _leaveCb() if (isDef(el._leaveCb)) { el._leaveCb.cancelled = true el._leaveCb() } // 一系列处理，这里忽略，后面会具体分析 } 复制代码`

* 首先我们看下 ` enter` 对于过渡数据是如何处理的：它会将 ` vnode.data.transition` 传入给 ` resolveTransition()` 当参数，并对其进行解析进而拿到 ` data` 数据

` const data = resolveTransition(vnode.data.transition) // 如果 data 不存在，则直接返回 if (isUndef(data)) { return } // 如果 el 中存在 _enterCb 或者 el 不是元素节点，则直接返回 if (isDef(el._enterCb) || el.nodeType !== 1 ) { return } const { css, type, enterClass, enterToClass, enterActiveClass, appearClass, appearToClass, appearActiveClass, beforeEnter, enter, afterEnter, enterCancelled, beforeAppear, appear, afterAppear, appearCancelled, duration } = data 复制代码`

然后我们在看看 ` resolveTransition()` ：该方法有一个参数为 ` vnode.data.transition` ，它通过 ` autoCssTransition()` 方法处理 ` name` 属性，并拓展到 ` vnode.data.transition` 上并进行返回

` export function resolveTransition ( def?: string | Object ): ? Object { if (!def) { return } if ( typeof def === 'object' ) { const res = {} if (def.css !== false ) { extend(res, autoCssTransition(def.name || 'v' )) } extend(res, def) return res } else if ( typeof def === 'string' ) { return autoCssTransition(def) } } 复制代码`

其中 ` autoCssTransition()` 具体逻辑如下：获取到参数 ` name` 后返回与 ` name` 相关的 ` css class`

` const autoCssTransition: ( name: string ) => Object = cached( name => { return { enterClass : ` ${name} -enter` , enterToClass : ` ${name} -enter-to` , enterActiveClass : ` ${name} -enter-active` , leaveClass : ` ${name} -leave` , leaveToClass : ` ${name} -leave-to` , leaveActiveClass : ` ${name} -leave-active` } }) 复制代码`

* 随即，对 ` <transition>` 是子组件的根节点的边界情况进行处理，我们需要对其父组件进行 ` appear check`

` let context = activeInstance let transitionNode = activeInstance.$vnode // 往上查找出 <transition> 是子组件的根节点的边界情况，进行赋值 while (transitionNode && transitionNode.parent) { context = transitionNode.context transitionNode = transitionNode.parent } // 上下文实例没有 mounted 或者 vnode 不是根节点插入的 const isAppear = !context._isMounted || !vnode.isRootInsert if (isAppear && !appear && appear !== '' ) { return } 复制代码`

* 对过渡的 ` class` 、钩子函数进行处理

` // 过渡 class 处理 const startClass = isAppear && appearClass ? appearClass : enterClass const activeClass = isAppear && appearActiveClass ? appearActiveClass : enterActiveClass const toClass = isAppear && appearToClass ? appearToClass : enterToClass // 钩子函数处理 const beforeEnterHook = isAppear ? (beforeAppear || beforeEnter) : beforeEnter const enterHook = isAppear ? ( typeof appear === 'function' ? appear : enter) : enter const afterEnterHook = isAppear ? (afterAppear || afterEnter) : afterEnter const enterCancelledHook = isAppear ? (appearCancelled || enterCancelled) : enterCancelled 复制代码`

* 获取其他配置

` const explicitEnterDuration: any = toNumber( isObject(duration) ? duration.enter : duration ) // 获取 enter 动画执行时间 if (process.env.NODE_ENV !== 'production' && explicitEnterDuration != null ) { checkDuration(explicitEnterDuration, 'enter' , vnode) } // 过渡动画是否受 css 影响 const expectsCSS = css !== false && !isIE9 // 用户是否想介入控制 css 动画 const userWantsControl = getHookArgumentsLength(enterHook) 复制代码`

* 对 ` insert` 钩子函数进行合并

` if (!vnode.data.show) { // remove pending leave element on enter by injecting an insert hook mergeVNodeHook(vnode, 'insert' , () => { const parent = el.parentNode const pendingNode = parent && parent._pending && parent._pending[vnode.key] if (pendingNode && pendingNode.tag === vnode.tag && pendingNode.elm._leaveCb ) { pendingNode.elm._leaveCb() } enterHook && enterHook(el, cb) }) } 复制代码`

* 过渡动画钩子函数执行时机：先执行 ` beforeEnterHook` 钩子，并将 DOM 节点 ` el` 传入，随即判断是否希望通过 css 来控制动画，如果是 ` true` 则，执行 ` addTransitionClass()` 方法为节点加上 ` startClass` 和 ` activeClass` 。
* 然后执行 ` nextFrame()` 进入下一帧，下一帧主要是移除掉上一帧增加好的 ` class` ；随即判断过渡是否取消，如未取消，则加上 ` toClass` 过渡类；之后如果用户没有通过 ` enterHook` 钩子函数来控制动画，此时若用户指定了 ` duration` 时间，则执行 ` setTimeout` 进行 ` duration` 时长的延时，否则执行 ` whenTransitionEnds` 决定 ` cb` 的执行时机

` beforeEnterHook && beforeEnterHook(el) if (expectsCSS) { addTransitionClass(el, startClass) addTransitionClass(el, activeClass) nextFrame( () => { removeTransitionClass(el, startClass) if (!cb.cancelled) { addTransitionClass(el, toClass) if (!userWantsControl) { if (isValidDuration(explicitEnterDuration)) { setTimeout(cb, explicitEnterDuration) } else { whenTransitionEnds(el, type, cb) } } } }) } 复制代码`

* 最后执行事先定义好 ` cb` 函数：若动画受 ` css` 控制，则移除掉 ` toClass` 和 ` activeClass` ；随即判定过渡是否取消，若取消了，直接移除 ` startClass` 并执行 ` enterCancelledHook` ，否则继续执行 ` afterEnterHook`

` // enter 执行后不同动机对应不同的 cb 回调处理 const cb = el._enterCb = once( () => { if (expectsCSS) { removeTransitionClass(el, toClass) removeTransitionClass(el, activeClass) } if (cb.cancelled) { if (expectsCSS) { removeTransitionClass(el, startClass) } enterCancelledHook && enterCancelledHook(el) } else { afterEnterHook && afterEnterHook(el) } el._enterCb = null }) 复制代码`

上面牵扯到的函数的代码具体如下：

* ` nextFrame` ：简易版 ` requestAnimationFrame` ，参数为 ` fn` ，即下一帧需要执行的方法
* ` addTransitionClass` ：为当前元素 ` el` 增加指定的 ` class`
* ` removeTransitionClass` ：移除当前元素 ` el` 指定的 ` class`
* ` whenTransitionEnds` ：通过 ` getTransitionInfo` 获取到 ` transition` 的一些信息，如 ` type` 、 ` timeout` 、 ` propCount` ，并为 ` el` 元素绑定上 ` onEnd` 。随后不停执行下一帧，更新 ` ended` 值，直到动画结束。则将 ` el` 元素的 ` onEnd` 移除，并执行 ` cb` 函数
` // nextFrame const raf = inBrowser ? window.requestAnimationFrame ? window.requestAnimationFrame.bind( window ) : setTimeout : fn => fn() export function nextFrame ( fn: Function ) { raf( () => { raf(fn) }) } // addTransitionClass export function addTransitionClass ( el: any, cls: string ) { const transitionClasses = el._transitionClasses || (el._transitionClasses = []) if (transitionClasses.indexOf(cls) < 0 ) { transitionClasses.push(cls) addClass(el, cls) } } // removeTransitionClass export function removeTransitionClass ( el: any, cls: string ) { if (el._transitionClasses) { remove(el._transitionClasses, cls) } removeClass(el, cls) } // whenTransitionEnds export function whenTransitionEnds ( el: Element, expectedType: ?string, cb: Function ) { const { type, timeout, propCount } = getTransitionInfo(el, expectedType) if (!type) return cb() const event: string = type === TRANSITION ? transitionEndEvent : animationEndEvent let ended = 0 const end = () => { el.removeEventListener(event, onEnd) cb() } const onEnd = e => { if (e.target === el) { if (++ended >= propCount) { end() } } } setTimeout( () => { if (ended < propCount) { end() } }, timeout + 1 ) el.addEventListener(event, onEnd) } 复制代码`

### 2、leave ###

上面我们讲完了 ` enter` 这块 ` 输入 => 输出` 的处理，它主要发生在组件插入后。接下来就是与之对应 ` leave` 阶段的 ` 输入 => 输出` 处理了，它主要发生在组件销毁前。 ` leave` 和 ` enter` 阶段的处理非常类似，所以这里我就不再赘述，大家自行阅读即可。下面我稍微谈下其中的 ` delayLeave` 对于延时执行 ` leave` 过渡动画是如何设计的

` export function leave ( vnode: VNodeWithData, rm: Function ) { const el: any = vnode.elm // call enter callback now if (isDef(el._enterCb)) { el._enterCb.cancelled = true el._enterCb() } const data = resolveTransition(vnode.data.transition) if (isUndef(data) || el.nodeType !== 1 ) { return rm() } if (isDef(el._leaveCb)) { return } const { css, type, leaveClass, leaveToClass, leaveActiveClass, beforeLeave, leave, afterLeave, leaveCancelled, delayLeave, duration } = data const expectsCSS = css !== false && !isIE9 const userWantsControl = getHookArgumentsLength(leave) const explicitLeaveDuration: any = toNumber( isObject(duration) ? duration.leave : duration ) if (process.env.NODE_ENV !== 'production' && isDef(explicitLeaveDuration)) { checkDuration(explicitLeaveDuration, 'leave' , vnode) } const cb = el._leaveCb = once( () => { if (el.parentNode && el.parentNode._pending) { el.parentNode._pending[vnode.key] = null } if (expectsCSS) { removeTransitionClass(el, leaveToClass) removeTransitionClass(el, leaveActiveClass) } if (cb.cancelled) { if (expectsCSS) { removeTransitionClass(el, leaveClass) } leaveCancelled && leaveCancelled(el) } else { rm() afterLeave && afterLeave(el) } el._leaveCb = null }) if (delayLeave) { delayLeave(performLeave) } else { performLeave() } function performLeave ( ) { // the delayed leave may have already been cancelled if (cb.cancelled) { return } // record leaving element if (!vnode.data.show && el.parentNode) { (el.parentNode._pending || (el.parentNode._pending = {}))[(vnode.key: any)] = vnode } beforeLeave && beforeLeave(el) if (expectsCSS) { addTransitionClass(el, leaveClass) addTransitionClass(el, leaveActiveClass) nextFrame( () => { removeTransitionClass(el, leaveClass) if (!cb.cancelled) { addTransitionClass(el, leaveToClass) if (!userWantsControl) { if (isValidDuration(explicitLeaveDuration)) { setTimeout(cb, explicitLeaveDuration) } else { whenTransitionEnds(el, type, cb) } } } }) } leave && leave(el, cb) if (!expectsCSS && !userWantsControl) { cb() } } } 复制代码`

` delayLeave` 是通过 ` resolveTransition(vnode.data.transition)` 获取到的函数，如果存在，则执行 ` delayLeave` ，否则直接执行 ` performLeave`

` if (delayLeave) { delayLeave(performLeave) } else { performLeave() } 复制代码`

能看出来 ` delayLeave` 是一个函数，它本身是不做任何操作的，唯一要做的事情就是将 ` performLeave` 作为回调参数暴露给用户去自行调用。

根据这个思路，我们简单改造一下上面官方的例子，具体如下

` < template > < div id = "example" > < button @ click = "show = !show" > Toggle render </ button > < transition name = "slide-fade" @ delay-leave = "handleDelay" > < p v-if = "show" > hello </ p > </ transition > </ div > </ template > < script > export default { data () { return { show : true } }, methods : { handleDelay (done) { setTimeout( () => { done() }, 2000 ) } } } </ script > < style lang = "scss" >.slide-fade-enter-active { transition : all . 3s ease; }.slide-fade-leave-active { transition : all . 8s cubic-bezier (1.0, 0.5, 0.8, 1.0); }.slide-fade-enter ,.slide-fade-leave-to { transform : translateX (10px); opacity : 0 ; } </ style > 复制代码`

最终实现的效果就是，我的 ` leave` 过渡效果将在 2S 后执行

## 总结 ##

文章到这，对于内置 ` <transition>` 组件的分析就介绍了。文章开篇到结尾，我都穿插了一个 ` 输入 => 输出` 的概念进去。 ` <transition>` 组件或者说 ` vue` 本身的设计也正是如此，它在内部帮我们做好了 ` 输入 => 输出` 的处理，让我们本身可以不用去管其内部，只需关注本身业务逻辑即可。

所以，目前看来，要把 ` <transition>` 用的溜，还需要你自己 ` css` 溜起来先

![](https://user-gold-cdn.xitu.io/2019/6/3/16b19672eda86ac3?imageView2/0/w/1280/h/960/ignore-error/1)

最后的最后，鄙人建了一个前端交流群：731175396

欢迎各位妹纸（算了汉纸也行吧 ~）加入，大家一起吹牛逼，聊人生，聊技术