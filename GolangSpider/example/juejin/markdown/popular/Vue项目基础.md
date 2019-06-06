# Vue项目基础 #

# 一. **Vue** #

## **1. 概念** ##

### 1). Vue是一个MVVM库 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f5234dcca63?imageView2/0/w/1280/h/960/ignore-error/1)

MVC（Model-View-Controller）：

M指的是从后台获取到的数据，

V指的是显示动态数据的html页面，

C是指响应用户操作、经过业务逻辑处理后去更新视图的过程，在此过程中会导致对view层的引用。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f5445a0a63c?imageView2/0/w/1280/h/960/ignore-error/1)

MVVM (Model-View-ViewModel)：
MVVM是MVC的一个衍生模型，这里的 ViewModel把业务逻辑处理、用户输入验证等跟视图更新操作分离开了。MVVM是数据驱动的，我们只需要关心数据的处理逻辑即可，它会通过模板渲染去单独处理视图的更新而不需要我们亲自去操作Dom元素。

### 2). Vue.js是一套构建用户界面的渐进式框架 ###

Vue的核心的功能，是一个视图模板引擎，但这不是说Vue就不能成为一个框架。如下图所示，这里包含了Vue的所有部件，在声明式渲染（视图模板引擎）的基础上，我们可以通过添加组件系统、客户端路由、大规模状态管理来构建一个完整的框架。更重要的是，这些功能相互独立，你可以在核心功能的基础上任意选用其他的部件，不一定要全部整合在一起。可以看到，所说的“渐进式”，其实就是Vue的使用方式，同时也体现了Vue的设计的理念。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f58ff999d99?imageView2/0/w/1280/h/960/ignore-error/1)

渐进式代表的含义是：主张最少。

每个框架都不可避免会有自己的一些特点，从而会对使用者有一定的要求，这些要求就是主张，主张有强有弱，它的强势程度会影响在业务开发中的使用方式。

比如说，Angular，它两个版本都是强主张的，如果你用它，必须接受以下东西：

* 必须使用它的模块机制- 必须使用它的依赖注入
* 必须使用它的特殊形式定义组件（这一点每个视图框架都有，难以避免）

所以Angular是带有比较强的排它性的，如果你的应用不是从头开始，而是要不断考虑是否跟其他东西集成，这些主张会带来一些困扰。

比如React，它也有一定程度的主张，它的主张主要是函数式编程的理念，比如说，你需要知道什么是副作用，什么是纯函数，如何隔离副作用。它的侵入性看似没有Angular那么强，主要因为它是软性侵入。

Vue可能有些方面是不如React，不如Angular，但它是渐进的，没有强主张，你可以在原有大系统的上面，把一两个组件改用它实现，当jQuery用；也可以整个用它全家桶开发，当Angular用；还可以用它的视图，搭配你自己设计的整个下层用。你可以在底层数据逻辑的地方用OO和设计模式的那套理念，也可以函数式，都可以，它只是个轻量视图而已，只做了自己该做的事，没有做不该做的事，仅此而已。

渐进式的含义，没有多做职责之外的事。

### 3). 声明式渲染也叫作响应式渲染 ###

* 声明式：告诉“机器”你想要的是什么(what)，让机器想出如何去做(how)。例如：Vue、React。

> 
> 
> 
> 在Vue中声明了message变量，当message变量的值改变时，文本对象的值会自动改变，这是声明式。通过Vue不仅可以通过声明的方式修改文本对象，还可以修改所有dom对象，包括四种类型（document、Node、Element、Text、Attributes）的对象。
> 
> 
> 

* 命令式：命令“机器”如何去做事情(how)，这样不管你想要的是什么(what)，它都会按照你的命令实现。例如：jquery、JavaScript。

> 
> 
> 
> 通过javascript的方式修改文本对象的值，我们需要先获取文本对象，然后给文本对象赋值，这是命令式。
> 
> 

### 4). 双向绑定/数据驱动 ###

双向中的两方分别是：Vue.js中的数据和DOM的对象。 VUE实现双向数据绑定的原理就是通过数据劫持结合发布者-订阅者模式的方式来实现的。即利用的是ES5的Object.defineProperty和存储器属性: getter和setter（所以只兼容IE9及以上版本），可称为基于依赖收集的观测机制。核心是VM，即ViewModel，保证数据和视图的一致性。 当修改Vue中数据时，dom的对象会自动修改，当修改dom的对象时，Vue的数据也会自动修改。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f5f63a19e2b?imageView2/0/w/1280/h/960/ignore-error/1)

当你把一个普通的 JavaScript 对象传入 Vue 实例作为 data 选项，Vue将遍历此对象所有的属性，并使用 [Object.defineProperty]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FJavaScript%2FReference%2FGlobal_Objects%2FObject%2FdefineProperty ) 把这些属性全部转为 getter/setter。Object.defineProperty 是ES5 中一个无法 shim 的特性，这也就是 Vue 不支持 IE8 以及更低版本浏览器的原因。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f650525aa0b?imageView2/0/w/1280/h/960/ignore-error/1) 这些 getter/setter 对用户来说是不可见的，但是在内部它们让 Vue能够追踪依赖，在属性被访问和修改时通知变更。每个组件实例都对应一个 **watcher** 实例，它会在组件渲染的过程中把“接触”过的数据属性记录为依赖。之后当依赖项的setter 触发时，会通知 watcher，从而使它关联的组件重新渲染。

受现代 JavaScript 的限制(而且 Object.observe 也已经被废弃)，Vue 无法检测到对象属性的添加或删除。由于 Vue会在初始化实例时对属性执行 getter/setter转化，所以属性必须在 data 对象上存在才能让 Vue 将它转换为响应式的。由于 Vue不允许动态添加根级响应式属性，所以你必须在初始化实例前声明所有根级响应式属性，哪怕只是一个空值。如果你未在 data 选项中声明 message，Vue将警告你渲染函数正在试图访问不存在的属性。

### 5). 异步更新队列 ###

Vue 在更新 DOM 时是异步执行的。只要侦听到数据变化，Vue将开启一个队列，并缓冲在同一事件循环中发生的所有数据变更。如果同一个 watcher被多次触发，只会被推入到队列中一次。当你设置 vm.someData = 'new value'，该组件不会立即重新渲染。当刷新队列时，组件会在下一个事件循环“tick”中更新。 为了在数据变化之后等待 Vue 完成更新DOM，可以在数据变化之后立即使用 Vue.nextTick(callback)。这样回调函数将在 DOM更新完成后被调用。$nextTick() 返回是一个 Promise 对象。

` methods: { updateMessage: async function () { this.message = '已更新' console.log(this.\ $el.textContent) // =\> '未更新' await this.\ $nextTick () console.log(this.\ $el.textContent) // =\> '已更新' } } 复制代码`

### 6). 组件化 ###

实现了扩展HTML元素，封装可重用的代码。每一个组件都对应一个ViewModel。页面上每个独立的可视/可交互区域都可以视为一个组件。每个组件对应一个工程目录，组件所需要的各种资源在这个目录下就进维护。页面是组件的容器，组件可以嵌套自由组合形成完整的页面。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f73d65a0d9c?imageView2/0/w/1280/h/960/ignore-error/1)

## **2. 生命周期** ##

vue框架的入口就是Vue实例，每个Vue实例在被创建之前都要经过一系列的初始化过程,这个过程就是vue的生命周期。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f7663dced43?imageView2/0/w/1280/h/960/ignore-error/1)

在vue一整个的生命周期中会有很多钩子函数提供给我们在vue生命周期不同的时刻进行操作。钩子函数类似于回调事件，当vue的实例进行到某处时会检查是否有对应的钩子函数，有则执行回调函数。

### 1). beforeCreate ###

> 
> 
> 
> 进行初始化事件，进行数据的观测，可以看到在created的时候数据已经和data属性进行绑定（放在data中的属性当值发生改变的同时，视图也会改变），但是没有el选项。
> 
> 
> 

### 2). created ###

> 
> 
> 
> 判断对象是否有el选项。如果有的话就继续向下编译，如果没有el选项，则停止编译，也就意味着停止了生命周期，直到在该vue实例上调用vm.$mount(el)。
> 
> 
> 
> 
> * 如果vue实例对象中有template参数选项，则将其作为模板编译成render函数。
> * 如果没有template选项，则将外部HTML作为模板编译。
> * 可以看到template中的模板优先级要高于outer HTML的优先级。
> 
> 
> 

### 3). beforeMount ###

> 
> 
> 
> 给vue实例对象添加$el成员，并且替换掉挂在的DOM元素。
> 
> 

### 4). mounted ###

> 
> 
> 
> 通过虚拟DOM渲染页面元素内容。
> 
> 

### 5). beforeUpdate ###

> 
> 
> 
> 当vue发现data中的数据发生了改变，会触发对应组件的重新渲染，渲染前的钩子函数。
> 
> 

### 6). updated ###

> 
> 
> 
> 当vue发现data中的数据发生了改变，会触发对应组件的重新渲染，渲染后的钩子函数。
> 
> 

### 7). beforeDestroy ###

> 
> 
> 
> 钩子函数在实例销毁之前调用。在这一步，实例仍然完全可用。
> 
> 

### 8). destroyed ###

> 
> 
> 
> 钩子函数在Vue 实例销毁后调用。调用后，Vue实例指示的所有东西都会解绑定，所有的事件监听器会被移除，所有的子实例也会被销毁。
> 
> 

## **3. 语法** ##

页面绑定的指令都可以是变量、对象的属性也可以是表达式，但只能是单个语句表达式，比如三元运算符。

以v-开头的都被叫做指令。当表达式的值改变时，将其产生的连带影响，响应式地作用于DOM。

### {{ msg }} ###

> 
> 
> 
> 纯文本。元素节点如果绑定 v-once指令，能执行一次性地插值，当数据改变时，插值处的内容不会更新。
> 
> 

### v-text="message" ###

> 
> 
> 
> 元素内绑定字符串
> 
> 

### v-html=”html” ###

> 
> 
> 
> 元素内传递HTML标签
> 
> 

### v-bind ###

可以简写成: 。用于绑定页面DOM元素上的属性值，比如 :src 、 :type

> 
> 
> 
> :style 支持样式数组、对象和直接写样式
> 
> 

> 
> 
> 
> :class支持数组、对象和三元表达式对其赋值，动态的切换class
> 
> 

### v-model ###

v-model 指令用来表单元素的双向绑定，v-model

> 
> 
> 
> 会根据控件类型自动选取正确的方法来更新元素。在input、select、textarea、checkbox、radio等表单控件元素上创建双向数据绑定，根据表单>
> 上的值，自动更新绑定的元素的值。
> 
> 

v-model.lazy

默认情况下， v-model 在 input 事件中同步输入框的值与数据，但你可以添加一个修饰符lazy ，从而转变为在 change 事件中同步，即只有input输入框失去焦点的时候才同步

v-model.number

> 
> 
> 
> 自动将用户的输入值转为 Number 类型（如果原值的转换结果为 NaN 则返回原值）
> 
> 

v-model.trim

> 
> 
> 
> 自动过滤用户输入的首尾空格
> 
> 

### v-on ###

可以简写成@。用来绑定事件监听器，事件类型由参数指定。表达式可以是一个方法的名字或一个内联语句。用在普通元素上时，只能监听原生DOM 事件。在监听原生 DOM事件时，方法以事件为唯一的参数。如果使用内联语句进行传参，语句可以访问一个 $event 属性。 v-on可以以绑定修饰符。修饰符是以半角句号 . 指明的特殊后缀，用于指出一个指令应该以特殊方式绑定。

> 
> 
> 
> 事件修饰符： .stop阻止事件冒泡 .once只能响应一次 …
> 
> 

> 
> 
> 
> 按键修饰符 .enter 捕获回车键 .delete捕获 "删除" 和 "退格" .space 捕获空格键 ...
> 
> 

### v-pre ###

> 
> 
> 
> 跳过某元素和他的子元素的编译，可以用来显示原始Mustache标签
> 
> 

### v-cloak ###

> 
> 
> 
> 使某元素保持某种指定行为，直到与该元素相关的实例编译结束。隐藏未编译的标签，通常用来解决{{}}表达式闪烁问题
> 
> 

### v-once ###

> 
> 
> 
> 只渲染元素和组件一次，之后重新渲染，该元素和组件均会被视作静态内容跳过
> 
> 

### v-for ###

> 
> 
> 
> 遍历数组或者对象，基于数据源多次渲染元素或模板块。 v-for 属性值是一个 item inexpression 结构的表达式，其中 item 代表
> expression 运算结果的每一项。除了item in items 这种形式还可以用 （value，key） in Objects
> 或者（value，key,index） in Objects来循环遍历。也支持对整数的循环。
> 
> 

### v-if、v-else-if、v-else ###

> 
> 
> 
> 条件渲染指令，动态插入或删除元素。通过表达式结果的真假来插入和删除元素。 v-if可以单独使用，而 v-else-if 、 v-else 必须和
> v-if 一起使用。
> 
> 

### v-show ###

> 
> 
> 
> 条件渲染指令，动态显示或隐藏元素。符合表达式结果判断的元素内容将被显示，不符合结果判断的元素将被隐藏
> 
> 

### filters ###

` {{ message \| capitalize }} v-bind:id="rawId \| formatId"`

> 
> 
> 
> 过滤器可以用在两个地方：mustache 插值和 v-bind表达式。允许自定义过滤器，对常见的文本进行格式化。过滤器是
> JavaScript函数，因此可以接受参数。
> 
> 

### computed ###

> 
> 
> 
> 计算属性。computed是基于它的依赖缓存，只有相关依赖发生改变时才会重新取值。而使用 methods，在重新渲染的时候，函数总会重新调用执行。即
> computed会缓存数据，当改变时才会去重新取值。而写在methods 中的则不会被缓存。
> 
> 

### $watch ###

> 
> 
> 
> 监听属性。Watch通过对象中创建的方式去实时监听数据变化并改变自身的值。
> 
> 

### Vue.component(tagName, options) ###

> 
> 
> 
> tagName 为组件名，options为配置选项。当注册成功后，再页面用<tagName></tagName>方式调用。
> 
> 

prop

> 
> 
> 
> prop 是父组件用来传递数据的一个自定义属性。 父组件的数据需要通过 props 把数据传给子组件，子组件需要显式地用 props
> 选项声明"prop"。用v-bind 指令可以绑定props，但是prop是单向绑定的：当父组件的属性变化时，将传导给子组件，但是不会反过来。
> 子组件声明时候可以验证prop的类型。可以使String、Number、Boolean、Function、Object和Array。
> 子组件在methods中定义的事件，可以让父组件使用v-on监听事件，也可以用$emit触发事件。
> 子组件的data，不能是一个对象。这样可以不会影响到其他实例。
> 
> 

### directives ###

> 
> 
> 
> Vue支持自定义注册全局与局部指令。
> 
> 

> 
> 
> 
> 指令定义函数提供了几个钩子函数（可选）：
> 
> 

> 
> 
> 
> * bind: 只调用一次，指令第一次绑定到元素时调用，用这个钩子函数可以定义一个在绑定时执行一次的初始化动作。
> * inserted: 被绑定元素插入父节点时调用（父节点存在即可调用，不必存在于 document 中）。
> * update:
> 被绑定元素所在的模板更新时调用，而不论绑定值是否变化。通过比较更新前后的绑定值，可以忽略不必要的模板更新（详细的钩子函数参数见下）。
> * componentUpdated: 被绑定元素所在模板完成一次更新周期时调用。
> * unbind: 只调用一次， 指令与元素解绑时调用。
> 
> 
> 

> 
> 
> 
> 钩子函数的参数有：
> 
> 

> 
> 
> 
> * el: 指令所绑定的元素，可以用来直接操作 DOM 。
> * binding: 一个对象，包含以下属性：
> 
> * name: 指令名，不包括 v- 前缀。
> 
> 
> 
> * value: 指令的绑定值， 例如： v-my-directive="1 + 1", value 的值是 2。
> * oldValue: 指令绑定的前一个值，仅在 update 和 componentUpdated 钩子中可用。无论值是否改变都可用。
> * expression: 绑定值的表达式或变量名。 例如 v-my-directive="1 + 1" ， expression 的值是 "1 +
> 1"。
> * arg: 传给指令的参数。例如 v-my-directive:foo， arg 的值是 "foo"。
> * modifiers: 一个包含修饰符的对象。 例如： v-my-directive.foo.bar, 修饰符对象 modifiers 的值是 {
> foo: true, bar: true }。
> * vnode: Vue 编译生成的虚拟节点。
> * oldVnode: 上一个虚拟节点，仅在 update 和 componentUpdated 钩子中可用。
> 
> 
> 

# 二. **Vue** #

## **1. 概念** ##

> 
> 
> 
> Vuex 是一个专为
> Vue.js应用程序开发的状态管理模式。它采用集中式存储管理应用的所有组件的状态即全局状态管理，并以相应的规则保证状态以一种可预测的方式发生变化。
> 当多个不是父子关系的组件需要共享状态时，就需要用到vuex。即把需要共享的变量存储在组件之外的对象里面，全局单例的状态管理树中的变量。
> 
> 

Vuex 和单纯的全局对象有以下两点不同：

* 

Vuex 的状态存储是响应式的。当 Vue 组件从 store 中读取状态的时候，若 store中的状态发生变化，那么相应的组件也会相应地得到高效更新。

* 

你不能直接改变 store 中的状态。改变 store 中的状态的唯一途径就是显式地提交(commit)mutation。这样使得我们可以方便地跟踪每一个状态的变化，从而让我们能够实现一些工具帮助我们更好地了解我们的应用。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f820244a9e0?imageView2/0/w/1280/h/960/ignore-error/1)

## **2. Store** ##

> 
> 
> 
> “store”基本上就是一个容器，它包含着你的应用中大部分的状态 。理解为存储于内存中的数据仓库。
> 
> 

## **3. State** ##

> 
> 
> 
> 单一状态树中的变量，用来存储状态。用一个对象就包含了全部的应用层级状态。即state、mutation、action、getter和嵌套子模块。
> 
> 

## **4. Getter** ##

> 
> 
> 
> computed计算属性，getter的返回值会根据它的依赖被缓存起来，且只有当它的依赖值发生了改变才会被重新计算。类似于get、set中的get。有两个可选参数：state、getters分别可以获取state中的变量和其他的getters。
> 调用方式：store.getters.XXX
> 
> 

## **5. Mutations** ##

> 
> 
> 
> 提交状态修改。类似于get、set中的set。也是vuex中唯一修改state的方式，只支持同步操作。第一个参数默认是state。
> 
> 

> 
> 
> 
> 调用方式：store.commit('SET_XXX', XXX)
> 
> 

## **6. Actions** ##

> 
> 
> 
> Action类似于mutations，不同的是不能直接变更状态。actions支持异步操作。第一个参数默认是和store具有相同参数属性的对象。
> 
> 

> 
> 
> 
> 调用方式：store.dispatch(‘XXX’)
> 
> 

## **7. Modules** ##

> 
> 
> 
> store的分割子模块，内容就相当于是store的一个实例。每个模块拥有自己的state、mutation、action、getter、甚至是嵌套子模块。
> 可添加namespaced:true来添加独立的命名空间。例如：store.dispatch(‘moduleName/XXX’)
> 调用子模块的方式：store.module.state.xxx
> 
> 

# 三. **VueRouter** #

## **1. 概念** ##

> 
> 
> 
> vue-router是 [Vue.js](
> https://link.juejin.im?target=http%3A%2F%2Fcn.vuejs.org%2F ) 官方的路由管理器。具有模块化的、基于组件的路由配置、细粒度的导航控制等等特点。它的作用是通过映射配置渲染对应的组件。
> 
> 
> 

> 
> 
> 
> 当两个路由渲染同一组件时，原来的组件会被复用，虽然效率提高但组件的生命周期钩子函数也不会被调用。
> 
> 

> 
> 
> 
> route是一条路由，url和函数的映射。
> 
> 

> 
> 
> 
> routes，是一组路由。
> 
> 

> 
> 
> 
> router管理一组route，它可以理解为一个容器，或者说一种机制。
> 
> 

> 
> 
> 
> 路径匹配规则支持正则表达式。
> 
> 

> 
> 
> 
> 以 / 开头的嵌套路径会被当作根路径。
> 
> 

> 
> 
> 
> 如果有设置嵌套路由，父路由不会渲染任何组件，只有设置空的子路由或者重定向才可以定位到父路由。
> 
> 

> 
> 
> 
> 同一个路径可以匹配多个路由，谁先定义谁的优先级最高。
> 
> 

## **2. 导航守卫** ##

> 
> 
> 
> 导航守卫是指通过发生改变或者取消的方式守卫导航。在路由导航改变的过程中植入导航守卫。参数或查询的改变并不会触发进入/离开的导航守卫。
> 
> 

解析流程：

* 

导航被触发。

* 

在失活的组件里调用离开守卫。

* 

调用全局的 beforeEach 守卫。

* 

在重用的组件里调用 beforeRouteUpdate 守卫。

* 

在路由配置里调用 beforeEnter。

* 

解析异步路由组件。

* 

在被激活的组件里调用 beforeRouteEnter。

* 

调用全局的 beforeResolve 守卫。

* 

导航被确认。

* 

调用全局的 afterEach 钩子。

* 

触发 DOM 更新。

* 

用创建好的实例调用 beforeRouteEnter 守卫中传给 next 的回调函数。

使用 router.beforeEach 注册一个全局前置守卫。当一个导航触发时，全局前置守卫按照创建顺序调用。守卫是异步解析执行，此时导航在所有守卫 resolve 完之前一直处于 等待中。

每个守卫方法接收三个参数：

* to: Route:

> 
> 
> 
> 即将要进入的目标 [路由对象](
> https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh%2Fapi%2F%23%25E8%25B7%25AF%25E7%2594%25B1%25E5%25AF%25B9%25E8%25B1%25A1
> )
> 
> 

* from: Route:

> 
> 
> 
> 当前导航正要离开的路由
> 
> 

* next: Function:

> 
> 
> 
> 一定要调用该方法来 resolve 这个钩子。执行效果依赖 next 方法的调用参数。一定要调用 next 方法，否则钩子就不会被
> resolved。
> 
> 

> 
> 
> 
> * next(): 进行管道中的下一个钩子。如果全部钩子执行完了，则导航的状态就是 confirmed (确认的)。
> * next(false): 中断当前的导航。如果浏览器的 URL 改变了 (可能是用户手动或者浏览器后退按钮)，那么 URL地址会重置到 from
> 路由对应的地址。
> * next('/') 或者 next({ path: '/' }): 跳转到一个不同的地址。当前的导航被中断，然后进行一个新的导航。你可以向
> next 传递任意位置对象，且允许设置诸如 replace: true、name:'home' 之类的选项以及任何用在 [router-link 的
> to prop](
> https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh%2Fapi%2F%23to
> ) 或 [router.push](
> https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh%2Fapi%2F%23router-push
> ) 中的选项。
> * next(error): (2.4.0+) 如果传入 next 的参数是一个 Error 实例，则导航会被终止且该错误会被传递给 [router.onError()](
> https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh%2Fapi%2F%23router-onerror
> ) 注册过的回调。
> 
> 
> 

## **3. API** ##

### <router-link> ###

> 
> 
> 
> 通过 to 属性指定目标地址，默认渲染成带有正确链接的 <a> 标签，可以通过配置
> tag属性生成别的标签.。另外，当目标路由成功激活时，链接元素自动设置一个表示激活的CSS 类名。
> 
> 

to：

> 
> 
> 
> 目标路由的链接。值可以是一个字符串或者是描述目标位置的对象
> 
> 

replace：

> 
> 
> 
> 导航后不会留下 history 记录
> 
> 

append：

> 
> 
> 
> 在当前 (相对) 路径前添加基路径。例如：/a append b = /a/b
> 
> 

tag：

> 
> 
> 
> 把a标签渲染成定义的标签
> 
> 

active-class：

> 
> 
> 
> 链接激活时使用的 CSS 类名
> 
> 

### <router-view> ###

> 
> 
> 
> 组件是一个 functional 组件，渲染路径匹配到的视图组件。<router-view>渲染的组件还可以内嵌自己的
> <router-view>，根据嵌套路径，渲染嵌套组件。可配合 <transition> 和 <keep-alive> 使用，缓存组件。
> 
> 

name：

> 
> 
> 
> 渲染对应的路由配置中 components 下的相应组件。
> 
> 

### Router参数 ###

mode：

> 
> 
> 
> 值为history/ hash。默认hash，使用 URL 的 hash
> 来模拟一个完整的URL浏览器url链接中会带有#，#后面的hash值的变化不会影响浏览器发出新请求。每次变化都会触发hasChange变化，监听变化更新页面。
> URL改变时，页面不会重新加载。history模式,就是直接匹配的/,这种模式充分利用 history.pushState
> API来完成URL跳转而无需重新加载页面。
> 
> 

> 
> 
> 
> 当使用history模式时，需要后台支持：当url匹配不到静态资源时，返回index.html。
> 
> 

name：

> 
> 
> 
> 当前路由的名称。
> 
> 

path:

> 
> 
> 
> 字符串，对应当前路由的路径。
> 
> 

redirect：

> 
> 
> 
> 重定向的path路径。
> 
> 

Component：

> 
> 
> 
> path路径对应的组件。
> 
> 

params:

> 
> 
> 
> key/value 对象，包含参数。类似post传参，地址栏无显示。
> 
> 

query：

> 
> 
> 
> key/value 对象，包含参数。类似get传参，地址栏显示参数。
> 
> 

fullPath：

> 
> 
> 
> 完成解析后包含查询参数和 hash 的完整路径的 URL。
> 
> 

Children：

> 
> 
> 
> 嵌套路由，传入的是数组格式的route。
> 
> 

### Router方法 ###

push:

> 
> 
> 
> 导航到不同的 URL。这个方法会向 history 栈添加一个新的记录，所以，当用户点击浏览器后退按钮时，则回到之前的
> URL。方法的参数可以是一个字符串路径，或者一个描述地址的对象。类似于html5中的window. [history.pushState()](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FAPI%2FHistory%2FpushState
> )
> 
> 

> 
> 
> 
> 如果提供了 path，params 会被忽略。path和query连用，name和params连用。
> 
> 

> 
> 
> 
> 调用方式：this.$router.push
> 
> 

> 
> 
> 
> <router-link :to="...">声明式 === router.push(...)编程式
> 
> 

replace:

> 
> 
> 
> 导航到不同的 URL。这个方法不会向 history 栈添加一个新的记录，而是替换掉当前记录。类似于html5中的window. [history.replaceState()](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FAPI%2FHistory%2FpushState
> )
> 
> 

> 
> 
> 
> <router-link :to="..." replace>声明式 === router.replace(...)编程式
> 
> 

.go(n):

> 
> 
> 
> history 记录中向前或者后退多少步，参数是一个整数值。router.go(n) === window.history.go(n)
> 
> 

# 四. **参考文档** #

vueAPI： [cn.vuejs.org/v2/api/#met…]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23methods )

vuexAPI： [vuex.vuejs.org/zh/api/]( https://link.juejin.im?target=https%3A%2F%2Fvuex.vuejs.org%2Fzh%2Fapi%2F )

vueRouterAPI： [router.vuejs.org/zh/api/]( https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh%2Fapi%2F )