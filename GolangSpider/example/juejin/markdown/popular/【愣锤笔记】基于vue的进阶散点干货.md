# 【愣锤笔记】基于vue的进阶散点干货 #

vue的开发如日中天，越来越多的开发者投入了vue开发大军中。希望本文中的一些vue散点，能在实际项目中灵活运用，切实地为我们解决一些难点问题。

### 插槽 ###

` // 组件调用时使用插槽 <todo-list-item v-for= "(item, index) in list" :key= "index" > <template v-slot:theNameOfSlot= "theSlotProps" > <span>{{item}}-{{theSlotProps.checked}}</span> </template> </todo-list-item> // todo-list-item组件的定义 <li> <input type = "checkbox" v-model= "checked" > // name定义插槽名 // :checked= "checked" 通过 bind 属性的方法，使得父组件在使用插槽时可以读取到现在分发出去的数，例如这样父组件就可以读取到这个分发出的checked值 <slot name= "theNameOfSlot" :checked= "checked" ></slot> </li> 复制代码`

注意点：

* 父组件的作用域是父组件的，如果想获取自组件的插槽数据，则需要自组件的插件将数据分发出去
* 2.6版本之后绑定插槽名称，只能在template模板上， ` v-slot:插槽名` , ` v-slot:theNameOfSlot="theSlotProps"` 属性值为slot分发给父组件的数据

### 依赖注入 ###

` // 在一个组件上设置注入的属性，可以是对象，也可以是函数返回一个对象 provide: { parentProvide: { msg: 'hello world' } }, // 在其任意层级的子节点可以获取到父节点注入的属性 inject: [ 'parentProvide' ] 复制代码`

依赖注入的属性是无法修改的，如果需要在祖孙组件中监听注入的属性变化，需要在祖宗组件中的注入属性为 ` this` , 即把祖宗属性作为注入属性往下传递。

` // 注意这里注入时使用的是函数返回的对象 provide () { return { parentProvide: this } }, // 接收注入的属性并可以直接修改，修改后祖宗的这个属性值也会变化 inject: [ 'parentProvide' ], methods: { updataParentMsg () { this.parentProvide.msg = '重置了' } }, 复制代码`

依赖注入很好的解决了在跨层级组件直接的通信问题，在封装高级组件的时候会很常用。

### 实现简易的vuex ###

` // 封装 import Vue from 'vue' const Store = function (options = {}) { const {state = {}, mutations = {}} = options this._vm = new Vue({ data: { $ $state : state } }) this._mutations = mutations } Store.prototype.commit = function ( type , payload) { if (this._mutations[ type ]) { this._mutations[ type ](this.state, payload) } } Object.defineProperties(Store.prototype, { state: { get () { return this._vm._data.$ $state } } }) export default { Store } // main.js，使用 // 首先导入我们封装的vuex import Vuex from './min-store/index' // 简易挂载 Vue.prototype. $store = new Vuex.Store({ state: { count: 1 }, getters: { getterCount: state => state.count }, mutations: { updateCount (state) { state.count ++ } } }) // 页面使用 computed: { count () { return this. $store.state.count } }, methods: { addCount () { this. $store.commit( 'updateCount' ) } }, 复制代码`

### vuex-getter的注意点 ###

` // getter第一个参数是state，第二个参数是其他getters，模块中的getter第三个参数是根状态 const getters = { count: state => state.count, // 例如可以返回getters.count * 2 otherCount: (state, getters) => getters.count * 2, // 跟状态 otherCount: (state, getters, rootState) => rootState.someState, } // 辅助函数 import { mapGetters } from 'vuex' computed: { ...mapGetters([ 'count' , 'otherCount' ]) }, // 模块加命名空间之后，car是当前getter所在的文件名 // 如果父级也有命名空间，则需要加上父级的命名空间，例如`parentName/car`: computed: { ...mapGetters( 'car' , [ 'count' , 'otherCount' ]) }, // 如果mapGetters中的值来自于多个模块，可以用对象的形式分别定义: ...mapGetters({ 'count' : 'car/count' , 'otherCount' : 'car/otherCount' , 'userName' : 'account/userName' }) // 也可以写多个mapGetters computed: { ...mapGetters( 'account' , { 'userName' : 'userName' }), ...mapGetters( 'car' , [ 'count' , 'otherCount' ]) } 复制代码`
> 
> 
> 
> 
> mapGetter的参数用数组的形式，书写更简洁方便，但是在需要重新命名getters等情况下则无法实现，此时可以换成对象的书写方式。
> 
> 

### vuex-mutations注意点 ###

推荐使用常量替代 mutation 事件类型：在store文件夹中新建mutation-types.js文件，将所有的mutation事件类型以常量的形式定义好。

` // mutation-types.js export const SOME_MUTATION = 'SOME_MUTATION' // 引用时，通过es6的计算属性命名的方式引入事件类型 import * as types from '../mutation-types' const mutations = { [types.UPDATE_USERINFO] (state, userInfo) { state.count = userInfo } } // 使用mapMutations import { mapMutations } from 'vuex' // 常量+命名空间的mutation貌似没法通过像mapGetters的一样用法，只能通过this. $store.commit的方式提交。 this[`account/ ${types.UPDATE_USERINFO} `]( 'xiaoming' ) // 所以个人觉得这种情况，还是用action去触发mutation 复制代码`
> 
> 
> 
> 
> 常量放在单独的文件中可以让你的代码合作者对整个 app 包含的 mutation 一目了然
> 
> 

### vetur ###

生成简易vue模板的快捷键：scaffold

### 灵活的路由配置 ###

` import Vue from 'vue' import Router from 'vue-router' Vue.use(Router) export default new Router({ mode: 'history' , routes: [ { path: '/user' , title: '个人中心' , component: { render: h => h( 'router-view' ) }, children: [ { path: '/user/index' , name: 'user-dashboard' , title: '个人中心' , meta: { // 其他meta信息 }, component: () => import(/* webpackChunkName: user */ '@/views/dashboard/index' ) }, { path: '/user/car' , name: 'user-car' , title: '我的汽车' , meta: { // 其他meta信息 }, component: () => import(/* webpackChunkName: user */ '@/views/car/index' ) } ] } ] }) 复制代码`

* ` () => import()` 自动代码分割的异步组件
* ` () => import(/* webpackChunkName: user */ '@/views/dashboard/index')` 将文件打包到一个模块，例如这里的写法可以将dashboard的index打包到user模块中
* ` render: h => h('router-view')` 可以用render函数灵活的创建user模块的入口文件，省去了入口文件的编写。

### spa中的页面刷新 ###

spa应用的刷新我们不能采取reload的方式。所以需要另辟蹊径。

方案一：当路由当query部分变化时，配router-view的key属性，路由是会重新刷新的。

` <router-view :key= " $route.path" > this. $router.replace({ path: this. $route.fullPath, query: { timestamp: Date.now() } }) 复制代码`
> 
> 
> 
> 
> 此种方法的弊端是url无缘无故多了一个参数，不是很好看。
> 
> 

方案二：新建一个redirect空页面，刷新就从当前页面跳转到redirect页面，然后在redirect页面理解replace跳转回来。

` // redirect页面 <script> export default { name: 'redirect' , // 在路由进入redirect前立即返回原页面，从而达到刷新原页面的一个效果 beforeRouteEnter (to, from, next) { next(vm => { vm. $router.replace({ path: to.params.redirect }) }) }, // 渲染一个空内容 render(h) { return h() } } </script> // 跳转方法，我们可以封装在utils公共函数中 // 在utils/index.js文件中： import router from '../router' /** * 刷新当前路由 */ export const refreshCurrentRoute = () => { const { fullPath } = router.currentRoute router.replace({ name: 'redirect' , params: { redirect: fullPath } }) } 复制代码`

### 自定义指令实现最细粒度的权限控制（组件/元素级别的） ###

为了扩展，我们在src下新建directives指令文件夹，然后新建index.js、account.js

src/directives/index.js：

` import account from './account' const directives = [ account ] const install = Vue => directives.forEach(e => e(Vue)) export default { install } 复制代码`

src/directives/account.js

` import { checkAuthorization } from '../utils/index' /** * 导出和权限相关的自定义指令 */ export default Vue => { /** * 自定义权限指令 * @description 当用户有权限才显示当前元素 */ Vue.directive( 'auth' , { inserted (el, binding) { if (!checkAuthorization(binding.value)) { el.parentNode && el.parentNode.removeChild(el) } } }) } 复制代码`

最后简单附上工具函数，具体的根据项目实际情况而定:

` /** * 检测用户是否拥有当前的权限 * @param { Array } auth 待检查的权限 */ export const checkAuthorization = (auth = []) => { if (!Array.isArray(auth)) { throw TypeError( '请检查参数类型：Excepted Array，got ' + Object.prototype.toString.call(auth).slice(8, -1)) } return store.state.account.authorization.some(e => auth.includes(e)) } 复制代码`

### 自定义权限组件实现最细粒度的权限控制（组件/元素级别的） ###

` // 自定义权限组件 <script> import { check } from "../utils/auth" ; export default { // 由于该组件笔记简单，不需要状态/周期等，可以将其设置为函数式组件 // 即无状态/无响应式数据/无实例/无this functional: true , // 接收的用于权限验证的参数 props: { authority: { type : Array, required: true } }, // 由于函数式组件缺少实例，故引入第二个参数作为上下文 render(h, context) { const { props, scopedSlots } = context; return check(props.authority) ? scopedSlots.default() : null; } }; </script> // main.js中注册为全局组件 import globalAuth from './components/authorization' Vue.component( 'global-auth' , globalAuth) // 使用 <global-auth :authorization= "['user']" >asdasdas</global-auth> 复制代码`

### 监听某个元素的尺寸变化，例如div ###

` // 安装resize-detector cnpm i --save resize-detector // 引入 import { addListener, removeListener } from 'resize-detector' // 使用 addListener(this. $refs [ 'dashboard' ], this.resize) // 监听 removeListener(this. $refs [ 'dashboard' ], this.resize) // 取消监听 // 一般我们会对回调函数进行去抖 methods: { // 这里用了lodash的debounce resize: _.debounce( function (e) { console.log( '123' ) }, 200) } 复制代码`
> 
> 
> 
> 
> [resize-detector传送门](
> https://link.juejin.im?target=https%3A%2F%2Fnpm.taobao.org%2Fpackage%2Fresize-detector
> )
> 
> 

### 提升编译速度 ###

可以通过在本地环境使用同步加载组件，生产环境异步加载组件

` // 安装插件 cnpm i babel-plugin-dynamic-import-node --save-dev // .bablerc文件的dev增加配置 "env" : { // 新增插件配置 "development" : { "plugins" : [ "dynamic-import-node" ] } // 其他的内容 …… } // 然后路由文件的引入依旧可以使用之前的异步加载的方式 component: () => import( 'xxx/xxx' ) // 通过注释可以使多个模块打包到一起 component: () => import(/* user */ 'xxx/xxx' ) 复制代码`

该方式修改本地环境和生产环境的文件加载方式，对代码的侵入性最小，在不需要的时候直接删去.bablerc文件中的配置就好，而不需要修改文件代码 [github传送门]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fairbnb%2Fbabel-plugin-dynamic-import-node )

### 参考 ###

> 
> 
> 
> 1.《Vue开发实战》视频，作者：唐金州 地址： [传送门](
> https://link.juejin.im?target=https%3A%2F%2Ftime.geekbang.org%2Fcourse%2Fintro%2F163%3Futm_term%3Dzeus0PMQ2%26amp%3Butm_source%3Dzhihu%26amp%3Butm_medium%3Dgeektime%26amp%3Butm_campaign%3D163-presell%26amp%3Butm_content%3Darticle%26amp%3Bgk_activity%3D0
> )
> 2. D2admin基于vue的中后台框架, [传送门](
> https://link.juejin.im?target=https%3A%2F%2Fdoc.d2admin.fairyever.com%2Fzh%2F
> )
> 
> 

### END ###

> 
> 
> 
> 百尺竿头、日进一步
> 我是愣锤，一名前端爱好者
> 欢迎交流、批评
> 
>