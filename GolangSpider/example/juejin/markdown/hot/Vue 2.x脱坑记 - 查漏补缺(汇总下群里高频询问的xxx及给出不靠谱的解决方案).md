# Vue 2.x脱坑记 - 查漏补缺(汇总下群里高频询问的xxx及给出不靠谱的解决方案) #

## 前言 ##

文章内容覆盖范围, ` 芝麻绿豆` 的破问题都有,不止于 ` vue` ;

给出的是方案,而非手把手一字一句的给你说十万个为什么!

## 问题汇总 ##

### Q:安装超时( ` install timeout` ) ###

方案有这么些:

* **` cnpm` : 国内对npm的镜像版本**

` /* cnpm website: https://npm.taobao.org/ */ npm install -g cnpm --registry=https: //registry.npm.taobao.org // cnpm 的大多命令跟 npm 的是一致的,比如安装,卸载这些 复制代码`

* 

**` yarn` 和 ` npm` 改源大法**

* 

使用 nrm 模块 : [www.npmjs.com/package/nrm]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fnrm )

* 

npm config : ` npm config set registry https://registry.npm.taobao.org`

* 

yarn config : ` yarn config set registry https://registry.npm.taobao.org`

### Q: 想学习Vue,要先学习脚手架的搭建么 ###

若是你想快速上手，用官方的脚手架即可( ` Vue-Cli 3` ( https://link.juejin.im?target=https%3A%2F%2Fcli.vuejs.org%2Fzh%2F ) )

因为不管是 ` webpack` 还是 ` parcel` ， ` gulp` ，都是一些构建工作流的东东；

学习脚手架的搭建，更多的是要针对项目业务进行定制，调优；

一般入门级的无需太早考虑这方面的，只要专心学好 ` Vue` 的使用姿势便可。

### Q:安装一些需要编译的包:提示没有安装 ` python` 、build失败等 ###

因为一些 ` npm` 的包安装需要编译的环境, ` mac` 和 ` linux` 都还好,

而window 用户依赖 ` visual studio 的一些库` 和 ` python 2+` ,

windows的小伙伴都装上:

* [windows-build-tools]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffelixrieseberg%2Fwindows-build-tools )
* [python 2.x]( https://link.juejin.im?target=https%3A%2F%2Fwww.python.org%2Fdownloads%2F )

### Q: ` can't not find 'xxModule'` - 找不到某些依赖或者模块 ###

这种情况一般报错信息可以看到是哪个包抛出的信息，一般卸载这个模块,安装重新安装下即可。

### Q: ` data functions should return an object` ###

这个问题是 ` Vue` 实例内,单组件的 ` data` 必须返回一个对象;如下

` export default { name : 'page-router-view' , data () { return { tabs : [ { title : '财务信息' , url : '/userinfo' }, { title : '帐号信息' , url : '/userinfo/base' } ] } } } 复制代码`

**为什么要 return 一个数据对象呢?**

官方解释如下: ` data` 必须声明为返回一个初始数据对象的函数，因为组件可能被用来创建多个实例。

如果 ` data` 仍然是一个纯粹的对象，则所有的实例将共享引用同一个数据对象！

简言之,组件复用下,不会造成数据同时指向一处,造出牵一发而动全身的破问题,

### Q:我给组件内的原生控件添加事件,怎么不生效了! ###

` <!--比如用了第三方框架,或者一些封装的内置组件; 然后想绑定事件--> <!--// 错误例子1--> < el-input placeholder = "请输入特定消费金额 " @ mouseover = "test()" > </ el-input > <!--// 错误例子2--> < router-link :to = "item.menuUrl" @ click = "toggleName=''" > < i :class = "['fzicon',item.menuIcon]" > </ i > < span > {{item.menuName}} </ span > </ router-link > <!--上面的两个例子都没法触发事件!--> <!--究其原因,少了一个修饰符 .native--> < router-link :to = "item.menuUrl" @ click.native = "toggleName=''" > < i :class = "['fzicon',item.menuIcon]" > </ i > < span > {{item.menuName}} </ span > </ router-link > <!--明明官方文档有的,一堆人不愿意去看,,Fuck--> <!--https://cn.vuejs.org/v2/guide/components.html#给组件绑定原生事件--> 复制代码`

### Q: ` provide` 和 ` inject` 是什么 ###

` Vue` 在2.2的时候,也提供了该概念。类比 ` ng provider` 和 ` react context` ;

### Q:我用了 ` axios` , 为什么 IE 浏览器不识别(IE9+) ###

那是因为 IE 整个家族都不支持 promise, 解决方案:

` npm install es6-promise // 在 main.js 引入即可 // ES6的polyfill require ( "es6-promise" ).polyfill(); 复制代码`

### Q:我在函数内用了 ` this.xxx=` ,为什么抛出 ` Cannot set property 'xxx' of undefined;` ###

这又是 ` this` 的套路了, ` this` 是和当前运行的上下文绑定的,

一般你在 ` axios` 或者其他 ` promise` , 或者 ` setInterval` 这些默认都是指向最外层的全局钩子.

简单点说:"最外层的上下文就是 ` window` ,vue内则是 Vue 对象而不是实例!";

解决方案:

* 暂存法: 函数内先缓存 ` this` , let that = this;(let是 es6, es5用 var)
* 箭头函数: 会强行关联当前运行区域为 this 的上下文;

` this` 的知识, 读"<<你不知道的 JS 系列>>"最为合适了,里面讲的很清楚

### Q:我看一些Vue教程有这么些写法,是什么意思 ` @click.prevent` , ` v-demo.a.b` ; ###

* ` @click.prevent` : 事件+修饰符 , 作用就是点击但又阻止默认行为
* ` v-demo.a.b` : 自定义指令+修饰符. 具体看你什么指令了,修饰符的作用大多是给事件增加一些确切的拓展功能

比如阻止事件冒泡,阻止默认行为,访问到原生控件,结合键盘快捷键等等

传送门: [事件修饰符]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fevents.html%23%25E4%25BA%258B%25E4%25BB%25B6%25E4%25BF%25AE%25E9%25A5%25B0%25E7%25AC%25A6 ) ;

可以自定义修饰符么?也是可以的,

可以通过全局 ` config.keyCodes` 对象自定义键值修饰符别名：

### Q:为什么我的引入的小图片渲染出来却是 ` data:image/png;base64xxxxxxxx` ###

这个是 webpack 里面的对应插件处理的.

对于小于多少 K 以下的图片(规定的格式)直接转为 base64格式渲染;

具体配置在 ` webpack.base.conf.js` 里面的 rules里面的 ` url-loader`

这样做的好处:在网速不好的时候先于内容加载和减少http的请求次数来减少网站服务器的负担。

### Q: ` Component template shold contain exactly one root element.If you are useing v-if on multiple elements , xxxxx` ###

大体就是说,单组件渲染 DOM 区域必须要有一个根元素,

可以用 ` v-if` 和 ` v-else-if` 指令来控制其他元素达到并存的状态

### Q:跨域问题怎么破! ###

比如 ` No 'Access-Control-Allow-Origin' header is present on the requested resource.`

这种问题老生常谈了,我就不细说了,大体说一下;

1: ` CORS` , 前后端都要对应去配置,IE10+ 2: ` nginx` 反向代理,一劳永逸 <-- 线上环境可以用这个

线下开发模式,比如你用了 ` vue-cli` , 里面的 webpack 有引入了 ` proxyTable` 这么个玩意, 也可以做接口反向代理

` // 在 config 目录下的index.js proxyTable: { "/bp-api" : { target : "http://new.d.st.cn" , changeOrigin : true , // pathRewrite: { // "^/bp-api": "/" // } } } // target : 就是 api 的代理的实际路径 // changeOrigin: 就是是变源,必须是, // pathRewrite : 就是路径重定向,一看就知道 复制代码`

当然还有依旧坚挺的 ` jsonp` 大法!不过局限性比较多,比较适合一些 **特殊** 的信息获取!

### Q:我需要遍历的数组值更新了,值也赋值了,为什么视图不更新! ###

那是因为有局限性啊,官方文档也说的很清楚,

只有一些魔改的之后的方法提供跟原生一样的使用姿势(可以触发视图更新);

一般我们更常用(除了魔改方法)的手段是使用: ` this.$set(obj,item,value)` ;

传送门: [数组更新检测(触发视图更新)]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Flist.html%23%25E6%2595%25B0%25E7%25BB%2584%25E6%259B%25B4%25E6%2596%25B0%25E6%25A3%2580%25E6%25B5%258B )

### Q:为什么我的组件间的样式不能继承或者覆写啊! ###

单组件开发模式下,请确认是否开启了 ` CSS` 模块化功能!!

也就是 ` scoped` (vue-cli 里面配置了,只要加入这个属性就自动启用)

` < style lang = "scss" scoped > </ style > 复制代码`

为什么不能简单的继承或者覆写呢,是因为每个类或者 id 乃至标签都会给自动在css后面添加自定义属性hash!

比如

` // 写的时候是这个 .trangle{} // 编译过后,加上了 hash .trangle[data-v-1ec35ffc]{} 复制代码`

这些都是在 css-loader 里面配置!

### Q:路由模式改为 ` history` 后,除了首次启动首页没报错,刷新访问路由都报错! ###

必须给对应的服务端配置查询的主页面,也可以认为是主路由入口的引导

官方文档也有,为毛总有人不喜欢去看文档,总喜欢做伸手党,FUCK

**传送门** : [Vue-Router history Mode]( https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh-cn%2Fessentials%2Fhistory-mode.html )

### Q:我想拦截页面,或者在页面进来之前做一些事情,可以么? ###

Of course !!

各种路由器的钩子!! 传送门: **导航守卫** ( https://link.juejin.im?target=https%3A%2F%2Frouter.vuejs.org%2Fzh-cn%2Fadvanced%2Fnavigation-guards.html ) ;

当然,记忆滚动的位置也可以做到,详情翻翻里面的文档

### Q: ` TypeError: xxx is not a function` ###

这种问题明显就是写法有问题,能不能动点脑子!!

### Q: 能不能跨级拿到 ` props` ###

这种情况是面向嵌套层次很深的组件，又要拿到上层的父传递的东东，

可以用 ` $attrs` 或者 ` inject + provide` 来实现

### Q: ` Uncaught ReferenceError: xxx is not define` ###

* 实例内的 ` data` 对应的变量没有声明
* 你导入模块报这个错误,那绝逼是导出没写好

### Q: ` Error in render function:"Type Error: Cannot read property 'xxx' of undefined"` ###

这种问题大多都是初始化的姿势不对;

比如引入 ` echart` 这些,仔细去了解下生命周期,再来具体初始化;

vue 组件有时候也会(嵌套组件或者 ` props` 传递初始化),也是基本这个问题

### Q: ` Unexpected token: operator xxxxx` ###

大佬,这个一看就是语法错误啊. 基本都是符号问题. 一般报错会给出哪一行或者哪个组件

### Q: ` npm run build` 之后不能直接访问 ###

大佬!你最起码得在本地搭个服务器才能访问好么!!

### Q: 操作Vue的原型链好么 ###

这个问题需要具体情况具体分析；

我看很多人喜欢把 ` axios` 挂载到 ` Vue.prototype` 上；

这样做有一定的弊端，相当耦合，若是多人维护或者替换其他库的时候有一定困难；

比较好的做法是不挂载，而是单独有服务请求的文件，用函数来封装你所需要的接口聚合；

这样统一暴露函数名，而内部实现可以随便改动

### Q:CSS ` background` 引入图片打包后,访问路径错误 ###

因为打包后图片是在根目录下,你用相对路径肯定报错啊,. 你可以魔改 webpack 的配置文件里面的 ` static` 为 `./static` ,但是不建议

你若是把图片什么丢到 ` assets` 目录下,然后相对路径,打包后是正常的

### Q:安装模块时命令窗口输出 ` unsupported platform xxx` ###

一般两种情况, ` node` 版本不兼容,系统不兼容;

解决方案: 要么不装,要么满足安装要求;

### Q: ` Unexpected tab charater` 这些 ###

一般是你用脚手架初始化的时候开了 eslint ;

要么遵循规则,要么改变规则;

要么直接把 webpack 里面的 eslint 检测给关闭了

### Q: ` Failed to mount component: template or render function not defined` ###

组件挂载失败,问题只有这么几个，组件没有正确引入或挂载点顺序错了了。

### Q: ` Unknown custom element: <xxx> - did you register the component correctly?` ###

组件没有正确引入或者正确使用,依次确认

* 导入对应的组件
* 在 components 内声明
* 在 dom 区域声明标签

### Q: 如何让自定义组件支持 ` Vue.use` 使用呢 ###

只要暴露一个 ` install` 函数即可，大体可以看下以下代码；

` import BtnPopconfirm from './BtnPopconfirm.vue' ; BtnPopconfirm.install = function ( Vue ) { Vue.component(BtnPopconfirm.name, BtnPopconfirm); }; export default BtnPopconfirm; // 然后就支持Vue.use了 复制代码`

### Q: ` axios` 的 ` post` 请求后台接受不到! ###

` axios` 默认是 json 格式提交,确认后台是否做了对应的支持;

若是只能接受传统的表单序列化,就需要自己写一个转义的方法,

当然还有一个更加省事的方案,装一个小模块 ` qs`

` // npm install qs -S // 然后在对应的地方转就行了,单一请求也行,拦截器也行,我是写在拦截器的. // 具体可以看看我 axios 封装那篇文章 //POST传参序列化(添加请求拦截器) Axios.interceptors.request.use( config => { // 在发送请求之前做某件事 if ( config.method === "post" ) { // 序列化 config.data = qs.stringify(config.data); // ***** 这里转义 } // 若是有做鉴权token , 就给头部带上token if (localStorage.token) { config.headers.Authorization = localStorage.token; } return config; }, error => { Message({ // 饿了么的消息弹窗组件,类似toast showClose: true , message : error, type : "error.data.error.message" }); return Promise.reject(error.data.error.message); } ); 复制代码`

### Q: ` Vue` 支持 ` jsx` 的写法么 ###

可以很确定的告诉你，是支持的；

但是和 ` React` 是有所差异的，而非完全等同的，具体可以看官方的支持库( [github.com/vuejs/jsx]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fjsx ) )

### Q: ` Invalid prop: type check failed for prop "xxx". Expected Boolean, got String.` ###

这种问题一般就是组件内的 ` props` 类型已经设置了接受的范围类型, 而你传递的值却又不是它需要的类型,写代码严谨些 OK?

### Q: 过滤器可以用于DOM区域结合指令么? ###

` // 不行,看下面的错误例子 < li v-for = "(item,index) in range | sortByDesc | spliceText" > {{item}} </ li > // `vue2+`的指令只能用语 mustache`{{}}` , 正确姿势如下: < span > {{ message | capitalize }} </ span > 复制代码`

### Q: ` [,Array]` , ` ,mapState` , ` [SOME_MUTATION] (state) {}` , ` increment ({ commit }) {}` 这种写法是什么鬼! ###

出门左拐, ` ES6+(ES2015+)` 的基础去过一遍,

上面依次:数组解构,对象解构,对象风格函数,对象解构赋值传递

### Q: 我的 Vue 网站为什么 UC 访问一片空白亦或者 ` flex` 布局错乱!! ###

来来来,墙角走起,. **UC 号称移动界的 IE 这称号不是白叫的**

* ` flexbox` 布局错乱,一般是你没有把兼容方案写上,就是带各种前缀,复合属性拆分，引入 ` autoprefixer` , 写上兼容范围就好了.
* ` UC访问空白` , 有一种情况绝对会造成,那就是 ES6的代码降级不够彻底. 其他情况可能就是路由配置问题(自己去排除)
* 现在的开发都推荐按需引入,靠 ` babel-preset-env` ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbabel%2Fbabel-preset-env ) 来控制,以达到打包体积减小.
* 但是这样做的后果,有些内核比较老的,嘿嘿,拜拜,
* 所以最好把代码完全 ES5话!!记住有些特性不能乱使用,没有对应的 ` polyfill` ,比如 ES6 的 ` proxy`

### Q: ` this.$set | this.$xxx` 这个 ` $` 是个什么意思?是 ` jQuery` 的么,会冲突么? ###

且看我细细道来.

Vue 的 ` $` 和 jQuery 的 ` $` 并没有半毛钱的关系,就跟 ` javascript` 和 ` java` 一样.

Vue 的 ` $` 是封装了一些 vue 的内建函数,然后导出以 ` $` 开头,这显然并不是 ` jQuery` 的专利;

jQuery 的 ` $` 是选择器!!取得 DOM区域,两者的作用完全不一致!

### Q: ` Module not found: Error : Can't resolve 'xxx-loader' in xxxx` ###

这里问题一般就是webpack的配置文件你改动了或对应的 ` loader` 没有装上

### Q: 父组件可以直接调用子组件的方法么! ###

可以,通过 ` $refs` 或者 ` $chilren` 来拿到对应的实例,从而操作

* 

### Q: ` Error in event handler for "click":"xxx"` ###

这个问题大多都是你写的代码有问题.你的事件触发了. 但是组件内部缺少对应的实现或者变量,所以抛出事件错误.

解决方案:看着报错慢慢排查

### Q: 组件的通讯有哪几种啊! ###

基本最常用的是这三种;

* 父传子: ` props`
* 子传父: ` emit`
* 兄弟通讯:

* ` event bus` : 就是找一个中间组件来作为信息传递中介
* ` vuex` : 信息树

传送门:

* [基本通讯]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents.html )
* [Vuex]( https://link.juejin.im?target=https%3A%2F%2Fvuex.vuejs.org%2Fzh-cn%2Fintro.html )

### Q:既然 ` localStorage` 和 ` sessionStorage` 能做到数据维护,为什么还要引入 ` vuex` ! ###

这个问题问得好, ` Vuex` 的目的用来维护同级组件间的数据通讯,拥有一个共同的状态树;

仅仅活在 ` SPA` 的里面的**伪多页(路由)**内, 这种东东明明然 ` localStorage` 和 ` sessionStorage` 也可以做到,还能做到跨页面数据维护,还不会被浏览器刷新干掉,

为什么还要引入 ` vuex` , 我个人觉得原因只有这么一个,"可维护性"和"易用性"及

怎么理解呢?

* 可维护性: 因为是单向数据流,所有状态是有迹可循的,数据的传递也可以及时分发响应
* 易用性: 它使得我们组件间的通讯变得更强大,而不用借助中间件这类来实现不同组件间的通讯

而且代码量不多,若是你要用 ` ls` 或者 ` ss` ,你必须手动去跟踪维护你的状态表, 虽说可行,但是代码量会多很多,而且可读性很差,

是不是每个项目都需要用到 ` vuex` ? 答案是否定的,小型项目上这个反而是累赘,这东西一般是用在中型项目+的, 因为里面涉及需要维护的数据比较多,同级组件间的通讯比较频繁

若是用到 ` vuex` 的项目记得结合 ` ss` 或者 ` ls` 来达到某些状态持久化!为什么看下面!

### Q: ` vuex` 的用户信息为什么还要存一遍在浏览器里( ` sessionStorage或localStorage` ) ###

因为 ` vuex` 的 store 干不过刷新啊. 保存在浏览器的缓存内,若用户刷新的话,值再取一遍呗;

### Q:"有 Vue + Vue Router + Vuex"或什么"express + vue + mongodb"的项目学习么 ###

Github 一搜一大堆,提这些问题的人动动脑子!.传送门: **Github** ( https://link.juejin.im?target=http%3A%2F%2Fgithub.com%2F )

### Q:我会 Vue 我还需要学习 ` jQuery` 或者原生 ` JS` 么 ###

` jQuery` 还有很多公司在用,源码可以学习的地方很多，框架只是加快开发,提高效率,但不是你在这一行长期立足的根本;

大佬们都是各种设计模式和算法玩的好，才能写出这么优秀的框架。

前端的人不仅需要宽度,也要深度,这样才能走的更远,.

### Q: ` npm run dev` 报端口错误! ` Error: listen EADDRINUSE :::8080` ###

* 自己用 ` webpack` 搭脚手架的都不用我说了;
* ` Vue-cli` 里面的 ` webpack` 配置: ` config/index.js`

` dev: { env : require ( "./dev.env" ), port : 8080 , // 这里这里,若是这个端口已经给系统的其他程序占用了.改我改我!! autoOpenBrowser: true , assetsSubDirectory : "static" , assetsPublicPath : "/" , proxyTable : { "/bp-api" : { target : "http://new.d.st.cn" , changeOrigin : true , // pathRewrite: { // "^/bp-api": "/" // } } }, 复制代码`

### Q: 什么时候用 ` v-if` ,什么用 ` v-show` ! ###

我们先来说说两者的核心差异;

* ` v-if` : DOM 区域没有生成,没有插入文档,等条件成立的时候才动态插入到页面!

* 有些需要遍历的数组对象或者值,最好用这货控制,等到拿到值才处理遍历,不然一些操作过快的情况会报错,比如数据还没请求到!

* ` v-show` : DOM 区域在组件渲染的时候同时渲染了,只是单纯用 css 隐藏了

* 对于下拉菜单,折叠菜单这些数据基本不怎么变动.用这个最合适了,而且可以改善用户体验,因为它不会导致页面的 **重绘** ,DOM 操作会!

简言之: ` DOM结` 构不怎么变化的用 ` v-show` , 数据需要改动很大或者布局改动的用 ` v-if`

### Q: ` <template>` 是什么,html5的标签么? ###

你猜对了, ` html5` 的标签还真有这么一个.传送门 [Can I Use:template]( https://link.juejin.im?target=https%3A%2F%2Fcaniuse.com%2F%23search%3Dtemplate )

不过 ` Vue` 的 ` template` 有点不一样,不是去给浏览器解析的,. 你可以理解为一个临时标签,用来方便你写循环,判断的,. 因为最终 ` template` 不会解析到浏览器的页面,他只是在 ` Vue` 解析的过程充当一个包裹层! 最终我们看到的是内部处理后的组合的 ` DOM` 结构!

### Q: ` Vue` 支持类似 ` React` 的 ` {,props}` 么 ###

` jsx` 的写法肯定是支持的，常规的写法也支持，用 ` v-bind="propsObject"` 会自动展开

### Q: ` Uncaught ReferenceError : Vue is not defined!` ###

依次排除:

* ` Vue` 是否正确引入!
* ` Vue` 是否正确实例化!
* ` Vue` 用的姿势是否正确(比如你直接一个 Vue 的变量!刚好又没定义,,具体问题具体分析吧)

### Q: ` ERROR in static/js/xxxxxxx.js from UglifyJs` ###

我知道其中一种情况会报这种情况,就是你引入的 js,是直接引入压缩版本后的 js( ` xxx.min.js` ); 然后 ` webpack` 内又启用了 ` UglifyJs` (压缩 JS的), 二重压缩大多都会报错!!

解决方案:引入标准未压缩的 JS

### Q: ` props` 不使用 ` :(v-bind)` 可以传递值么! ###

可以,只是默认传递的类型会被解析成字符串! 若是要传递其他类型,该绑定还是绑定!!

### Q: ` Uncaught TypeError : Cannot set property xxx which has only a getter` ###

这个问题就是你要操作的属性只允许 ` getter` ,不允许 ` setter` ;

解决方案? 用了别人的东西就要遵循别人的套路来,不然就只能自己动手丰衣足食了!!

### Q: 单组件中里面的 ` import xxx from '@/components/layout/xxx'` 中的 ` @` 是什么鬼! ###

这是 ` webpack` 方面的知识,看到了也说下吧,

` webpack` 可以配置 ` alias` (也就是路径别名),玩过 ` linux` 或者 ` mac` 都知道

依旧如上,会自己搭脚手架的不用我说了,看看 ` vue-cli` 里面的;

文件名: build -> webpack.base.conf.js

` resolve: { extensions : [ ".js" , ".vue" , ".json" ], // 可以导入的时候忽略的拓展名范围 alias: { vue$ : "vue/dist/vue.esm.js" , "@" : resolve( "src" ), // 这里就是别名了,比如@就代表直接从/src 下开始找起! "~" : resolve( "src/components" ) } }, 复制代码`

### Q: ` SCSS(SASS)` 还是 ` less` , ` stylus` 好!! ###

三者都是预处理器;

` scss` 出现最久,能做的功能比较多,但是若是普通的嵌套写法,继承, ` mixin` 啊.

这三个都差不多,会其中一个其他两个的粗浅用法基本也会了.不过!!

写法有些差异:

* ` scss` : 写法上是向 ` css` 靠齐
* ` sass` : 其实也就是 ` scss` , 只是写法不一样,靠的是缩进
* ` less` : 跟 ` css` 基本靠齐
* ` stylus` : 一样,靠缩进,跟 ` pug(Jade)` 一样

使用环境的差异:

* ` scss` 可以借助 ` ruby` 或者 ` node-sass` 或者 ` dart-sass` 编译
* ` less` 可以用 ` less.js` 或者对应的 ` loader` 解析
* ` stylus` 只能借助 ` loader` 解析,它的出现就是基于 ` node` 的

也有一个后起之秀,主打解耦,插件化的! 那就是 ` PostCSS` ,这个是后处理器! 有兴趣的可以自行去了解,上面的写法都能借助插件实现!

### Q: ` Failed to compile with x errors : This dependency was not found !` ###

编译错误,对应的依赖没找到!

解决如下:

* 知道缺少对应的模块,直接装进去
* 若是一个你已经安装的大模块(比如 ` axios` )里面的子模块(依赖包)出了问题,卸载重装整个大模块.因为你补全不一定有用!

### Q: ` SyntaxError: Unexpected identifier` ###

语法错误,看错误信息去找到对应的页面排查!

### Q: 为什么我的 ` npm` 或者 ` yarn` 安装依赖会生成 ` lock` 文件,有什么用! ###

lock 文件的作用是统一版本号,这对团队协作有很大的作用;

若是没有 lock 锁定,根据 ` package.json` 里面的 ` ^` , ` ~` 这些,

不同人,不同时间安装出来的版本号不一定一致;

有些包甚至有一些 ` breaking change` (破坏性的更新),造成开发很难顺利进行!

### Q: 组件可以缓存么? ###

可以,用 ` keep-alive` ;

不过是有代价的,占有内存会多了,所以无脑的缓存所有组件!别说性能好了,切换几次, 有些硬件 hold不住的,浏览器直接崩溃或者卡死,

所以 ` keep-alive` 一般缓存都是一些列表页,不会有太多的操作,更多的只是结果集的更换,

给路由的组件 ` meta` 增加一个标志位,结合 ` v-if` 就可以按需加上缓存了!

### Q: ` package.json` 里面的 ` dependencies` 和 ` devDependencies` 的差异! ###

其实不严格的话,没有特别的差异; 若是严格,遵循官方的理解;

* ` dependencies` : 存放线上或者业务能访问的核心代码模块,比如 ` vue` , ` vue-router` ;
* ` devDependencies` : 处于开发模式下所依赖的开发模块,也许只是用来解析代码,转义代码,但是不产生额外的代码到生产环境, 比如什么 ` babel-core` 这些

如何把包安装到对应的依赖下呢?

` npm install --save xxxx // dependencies npm install --save-dev xxxx // devDependencies //也能用简易的写法(i:install,-S:save,-D:save-dev) npm i -S xxxx // npm install --save xxxx npm i -D xxxx // npm install --save-dev xxxx 复制代码`

### Q: 安装 ` chromedriver` 报错!!姿势没错啊 ` npm i -D chromedriver` ###

恩,伟大的 GFW,,解决方案:指定国内的源安装就可以了

` npm install --save-dev chromedriver --chromedriver_cdnurl=http://cdn.npm.taobao.org/dist/chromedriver`

### Q: ` Vue ,React, Angular` 学习哪个好?哪个工作比较好找! ###

` Vue` 属于渐进式开发,传统开发过渡 MVVM 模式的小伙伴, ` Vue` 比较好上手,学习成本比较低 基础比较好的,有折腾精神的,可以选择 ` NG5` 或者 ` React 16` ;

NG5需要学习 ` typescript` 和 ` rxjs` ,还用到比较多的新东西,比如装饰器,后端的注入概念. ` ng` 有自己的一整套 MVVM 流程;

而 ` Vue` 和 ` React` 核心只是 ` view` ,可以搭配自己喜欢的

` React` 的写法偏向函数式写法,还有 jsx,官方自己有 ` flow` ,当然也能搭配 ` ts` ,我也没怎么接触,所以也有一定的学习成本;

至于哪个比较好找工作!告诉你,若是只会一个框架,那不是一个合格的前端;

人家要的是动手能力,解决能力!! **技术和待遇是成正比的** !!

颜值和背景,学历,口才可以加分,但是这些条件你必须要有的基础下才能考虑这些!

### Q: 我有个复杂组件需要有新增和编辑的功能同时存在,但是字段要保持不变性怎么破 ###

字段保持不变性怎么理解呢? 就是说比如新增和编辑同时共享一份 ` data` ;

有一种就是路由变了,组件渲染同一个(不引起组件的重新渲染和销毁!),但是功能却不同(新增和编译),

比如从编辑切到新增, ` data` 必须为空白没有赋值的,等待我们去赋值;

这时候有个东西就特别适合了,那就是 [immutable-js]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Fimmutable-js%2F ) ;

这个东西可以模拟数据的唯一性!或者叫做不变性!

### Q:"首屏加载比较慢!!怎么破!打包文件文件比较大" ###

依次排除和确认:

* 减少第三方库的使用,比如 ` jquey` 这些都可以不要了,很少操作 dom,而且原生基本满足开发
* 若是引入 ` moment` 这些,webpack 排除国际化语言包
* webpack 常规压缩js,css, 愿意折腾的还可以引入 dll 这些
* 路由组件采用懒加载
* 加入路由过渡和加载等待效果,虽然不能解决根本,但起码让人等的舒心一点不是么!

整体下来,打包之后一般不会太大;

但是倘若想要更快?那就只能采用服务端渲染(SSR)了,可以避免浏览器去解析模板和指令这些; 直接返回一个 html ,.还能 SEO,

### Vue你们如何做 ` spa` 的模块懒加载呢 ###

` // 推荐这种写法 // 一来可以聚合webpackChunkName名字一样的为一个模块，也是当前版本推荐的加载姿势 const Home = () => import ( /* webpackChunkName: "HomePage" */ "@/views/home/index.vue" ); 复制代码`

### Q: ` Vue SPA` 没法做优化( ` SEO` )!有解决方案么 ###

可以的, ` ssr` (服务端渲染就能满足你的需求),因为请求回来就是一个处理完毕的 ` html`

现在 ` vue` 的服务端开发框架有这么个比较流行,如下

传送门: [Nuxt.js]( https://link.juejin.im?target=https%3A%2F%2Fzh.nuxtjs.org%2F )

也有官方的方案, [ssr 完全指南]( https://link.juejin.im?target=https%3A%2F%2Fssr.vuejs.org%2Fzh%2F )

### Q: ` Vue` 可以写 ` hybird App` 么! ###

当然可以,两个方向.

* [codorva]( https://link.juejin.im?target=https%3A%2F%2Fcordova.apache.org%2F ) + [nativescript]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Frigor789%2Fnativescript-vue )
* [Weex]( https://link.juejin.im?target=https%3A%2F%2Fweex.apache.org%2F )

### Q: Vue 可以写桌面端么? ###

当然可以,有 ` electron` 和 ` node-webkit(nw)` ;

我只了解过 ` electron` ;

* [electron]( https://link.juejin.im?target=https%3A%2F%2Felectron.atom.io%2F )
* [electron-vue]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FSimulatedGREG%2Felectron-vue ) : Vue-cli 针对 electron 的脚手架模板

### Q: Vue开发,项目中还需要 ` jQuery` 么 ###

分情况探讨:

* 若是老项目,只是单纯引入 Vue 简化开发的,依旧用吧,
* 重构项目?或者发起新项目的,真心没必要了.开发思路不一样,很多以前用 DOM 操作的现在基本可以数据驱动实现,而少量迫不得已的DOM 操作原生就能搞定,而且能减小打包体积,速度又快,何乐而不为!

### Q:Vue PC(桌面)端,M(mobile:移动)端,用什么 UI 框架好啊! ###

PC: 推荐的只有两个 ` element UI` 和 ` iview`

mobile : ` Vux` 、 ` Vant`

### Q: Vue可以写微信小程序么,怎么搞起 ###

可以的,社区也有人出了对应的解决方案,比如比较流行的方案 ` wepy` ; ` wepy` 你也可以理解为一个脚手架,让你的写小程序的方式更贴近你用 ` vue-cli` 写 vue 的感觉,

传送门: [wepy]( https://link.juejin.im?target=https%3A%2F%2Fwepyjs.github.io%2Fwepy%2F%23%2F )

### Q: ` the "scope" attribute for scoped slots replaced by "slot-scope" since 2.5` ###

这个问题只出现老项目升级到 vue2.5+的时候, 提示就是 ` scope` 现在要用 ` slot-scope` 来代替, 但是 scope 暂时可以用,以后会移除

### Q: ` Vue` 2.6废除的特性清单 ###

自 2.6.0 起有所更新。已废弃的使用 slot 特性的语法在 [这里]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-slots.html%23%25E5%25BA%259F%25E5%25BC%2583%25E4%25BA%2586%25E7%259A%2584%25E8%25AF%25AD%25E6%25B3%2595 )

官方推荐用 ` v-slot` 来调用插槽

### Q:想要 mock 数据,直接请求 **json文件** 为什么不行! ###

当然不行,浏览器安全机制不允许,JS天生不能越权(NodeJS不能单纯说是JS)

你要 mock 数据,一般都有比较成熟的方案传送门:

* [Mock]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fnuysoft%2FMock )
* [Easy Mock]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feasy-mock%2Feasy-mock )

## Vue 周边库汇总 ##

**Awesome Vue** ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fawesome-vue ) : 里面收集了 ` Vue` 方方面面的热门库!!

## 结语 ##

问题目前就汇总了这么多，有不对之处请留言，会及时修正，谢谢阅读