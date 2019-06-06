# 7个有用的Vue开发技巧 #

## 1 状态共享 ##

随着组件的细化，就会遇到多组件状态共享的情况， ` Vuex` 当然可以解决这类问题，不过就像 ` Vuex` 官方文档所说的，如果应用不够大，为避免代码繁琐冗余，最好不要使用它，今天我们介绍的是vue.js 2.6新增加的 [Observable API]( https://link.juejin.im?target=https%3A%2F%2Fvuejs.org%2Fv2%2Fapi%2F%23Vue-observable ) ，通过使用这个api我们可以应对一些简单的跨组件数据状态共享的情况。

如下这个例子，我们将在组件外创建一个 ` store` ，然后在 ` App.vue` 组件里面使用store.js提供的 ` store` 和 ` mutation` 方法，同理其它组件也可以这样使用，从而实现多个组件共享数据状态。

首先创建一个store.js，包含一个 ` store` 和一个 ` mutations` ，分别用来指向数据和处理方法。

` import Vue from "vue" ; export const store = Vue.observable({ count : 0 }); export const mutations = { setCount(count) { store.count = count; } }; 复制代码`

然后在 ` App.vue` 里面引入这个store.js，在组件里面使用引入的数据和方法

` <template> < div id = "app" > < img width = "25%" src = "./assets/logo.png" > < p > count:{{count}} </ p > < button @ click = "setCount(count+1)" > +1 </ button > < button @ click = "setCount(count-1)" > -1 </ button > </ div > </ template > <script> import { store, mutations } from "./store" ; export default { name : "App" , computed : { count() { return store.count; } }, methods : { setCount : mutations.setCount } }; </ script > <style> 复制代码`

你可以点击 [在线DEMO]( https://link.juejin.im?target=https%3A%2F%2Fcodesandbox.io%2Fs%2F4w0mo0kyow ) 查看最终效果

## 2 长列表性能优化 ##

我们应该都知道 ` vue` 会通过 ` object.defineProperty` 对数据进行劫持，来实现视图响应数据的变化，然而有些时候我们的组件就是纯粹的数据展示，不会有任何改变，我们就不需要 ` vue` 来劫持我们的数据，在大量数据展示的情况下，这能够很明显的减少组件初始化的时间，那如何禁止 ` vue` 劫持我们的数据呢？可以通过 ` object.freeze` 方法来冻结一个对象，一旦被冻结的对象就再也不能被修改了。

` export default { data : () => ({ users : {} }), async created() { const users = await axios.get( "/api/users" ); this.users = Object.freeze(users); } }; 复制代码`

另外需要说明的是，这里只是冻结了 ` users` 的值，引用不会被冻结，当我们需要 ` reactive` 数据的时候，我们可以重新给 ` users` 赋值。

` export default { data : () => ({ users : [] }), async created() { const users = await axios.get( "/api/users" ); this.users = Object.freeze(users); }, methods :{ // 改变值不会触发视图响应 this.users[ 0 ] = newValue // 改变引用依然会触发视图响应 this.users = newArray } }; 复制代码`

## 3 去除多余的样式 ##

随着项目越来越大，书写的不注意，不自然的就会产生一些多余的css，小项目还好，一旦项目大了以后，多余的css会越来越多，导致包越来越大，从而影响项目运行性能，所以有必要在正式环境去除掉这些多余的css，这里推荐一个库 [purgecss]( https://link.juejin.im?target=https%3A%2F%2Fwww.purgecss.com%2F ) ，支持CLI、JavascriptApi、Webpack等多种方式使用，通过这个库，我们可以很容易的去除掉多余的css。

我做了一个测试， [在线DEMO]( https://link.juejin.im?target=https%3A%2F%2Fcodesandbox.io%2Fs%2Fzkq258ly4 )

` < h1 > Hello Vanilla! </ h1 > < div > We use Parcel to bundle this sandbox, you can find more info about Parcel < a href = "https://parceljs.org" target = "_blank" rel = "noopener noreferrer" > here </ a >. </ div > 复制代码` ` body { font-family: sans-serif; } a { color: red; } ul { li { list-style: none; } } 复制代码` ` import Purgecss from "purgecss" ; const purgecss = new Purgecss({ content : [ "**/*.html" ], css : [ "**/*.css" ] }); const purgecssResult = purgecss.purge(); 复制代码`

最终产生的 ` purgecssResult` 结果如下，可以看到多余的 ` a` 和 ` ul` 标签的样式都没了

![](https://user-gold-cdn.xitu.io/2019/5/21/16ad97db704d6639?imageView2/0/w/1280/h/960/ignore-error/1)

## 4 作用域插槽 ##

利用好作用域插槽可以做一些很有意思的事情，比如定义一个基础布局组件A，只负责布局，不管数据逻辑，然后另外定义一个组件B负责数据处理，布局组件A需要数据的时候就去B里面去取。假设，某一天我们的布局变了，我们只需要去修改组件A就行，而不用去修改组件B，从而就能充分复用组件B的数据处理逻辑，关于这块我之前写过一篇实际案例，可以点击 [这里]( https://juejin.im/post/5c2d7030f265da613a54236f ) 查看。

这里涉及到的一个最重要的点就是父组件要去获取子组件里面的数据，之前是利用 ` slot-scope` ，自vue 2.6.0起，提供了更好的支持 ` slot` 和 ` slot-scope` 特性的 API 替代方案。

比如，我们定一个名为 ` current-user` 的组件：

` <span> <slot>{{ user.lastName }}</slot> </span> 复制代码`

父组件引用 ` current-user` 的组件，但想用名替代姓（老外名字第一个单词是名，后一个单词是姓）：

` <current-user> {{ user.firstName }} </current-user> 复制代码`

这种方式不会生效，因为 ` user` 对象是子组件的数据，在父组件里面我们获取不到，这个时候我们就可以通过 ` v-slot` 来实现。

首先在子组件里面，将 ` user` 作为一个 ` <slot>` 元素的特性绑定上去：

` < span > < slot v-bind:user = "user" > {{ user.lastName }} </ slot > </ span > 复制代码`

之后，我们就可以在父组件引用的时候，给 ` v-slot` 带一个值来定义我们提供的插槽 prop 的名字：

` < current-user > < template v-slot:default = "slotProps" > {{ slotProps.user.firstName }} </ template > </ current-user > 复制代码`

这种方式还有缩写语法，可以查看 [独占默认插槽的缩写语法]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-slots.html%23%25E7%258B%25AC%25E5%258D%25A0%25E9%25BB%2598%25E8%25AE%25A4%25E6%258F%2592%25E6%25A7%25BD%25E7%259A%2584%25E7%25BC%25A9%25E5%2586%2599%25E8%25AF%25AD%25E6%25B3%2595 ) ，最终我们引用的方式如下：

` <current-user v-slot="slotProps"> {{ slotProps.user.firstName }} </current-user> 复制代码`

相比之前 ` slot-scope` 代码更清晰，更好理解。

## 5 属性事件传递 ##

写过高阶组件的童鞋可能都会碰到过将加工过的属性向下传递的情况，如果碰到属性较多时，需要一个个去传递，非常不友好并且费时，有没有一次性传递的呢（比如react里面的 ` {...this.props}` ）？答案就是 ` v-bind` 和 ` v-on` 。

举个例子，假如有一个基础组件 ` BaseList` ，只有基础的列表展示功能，现在我们想在这基础上增加排序功能，这个时候我们就可以创建一个高阶组件 ` SortList` 。

` <!-- SortList --> < template > < BaseList v-bind = "$props" v-on = "$listeners" > <!-- ... --> </ BaseList > </ template > <script> import BaseList from "./BaseList" ; // 包含了基础的属性定义 import BaseListMixin from "./BaseListMixin" ; // 封装了排序的逻辑 import sort from "./sort.js" ; export default { props : BaseListMixin.props, components : { BaseList } }; </ script > 复制代码`

可以看到传递属性和事件的方便性，而不用一个个去传递

## 6 函数式组件 ##

函数式组件，即无状态，无法实例化，内部没有任何生命周期处理方法，非常轻量，因而渲染性能高，特别适合用来只依赖外部数据传递而变化的组件。

写法如下：

* 在 ` template` 标签里面标明 ` functional`
* 只接受 ` props` 值
* 不需要 ` script` 标签
` <!-- App.vue --> < template > < div id = "app" > < List :items = "['Wonderwoman', 'Ironman']" :item-click = "item => (clicked = item)" /> < p > Clicked hero: {{ clicked }} </ p > </ div > </ template > < script > import List from "./List" ; export default { name : "App" , data : () => ({ clicked : "" }), components : { List } }; </ script > 复制代码` ` <!-- List.vue 函数式组件 --> < template functional > < div > < p v-for = "item in props.items" @ click = "props.itemClick(item);" > {{ item }} </ p > </ div > </ template > 复制代码`

## 7 监听组件的生命周期 ##

比如有父组件 ` Parent` 和子组件 ` Child` ，如果父组件监听到子组件挂载 ` mounted` 就做一些逻辑处理，常规的写法可能如下：

` // Parent.vue <Child @mounted= "doSomething" /> // Child.vue mounted() { this.$emit( "mounted" ); } 复制代码`

这里提供一种特别简单的方式，子组件不需要任何处理，只需要在父组件引用的时候通过 ` @hook` 来监听即可，代码重写如下：

` <Child @hook:mounted= "doSomething" /> 复制代码`

当然这里不仅仅是可以监听 ` mounted` ，其它的生命周期事件，例如： ` created` ， ` updated` 等都可以，是不是特别方便~

参考链接：

* [vueTips]( https://link.juejin.im?target=https%3A%2F%2Fvuedose.tips%2Ftips )
* [vuePost]( https://link.juejin.im?target=https%3A%2F%2Falligator.io%2Fvuejs%2F )