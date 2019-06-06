# Vue 项目里戳中你痛点的问题及解决办法（更新） #

最近要求使用vue进行前后端分离开发微信公众号，不断摸索踩坑之后，总结出如下几点vue项目开发中常见的问题及解决办法。如果你是vue大佬，请忽略小弟的愚见^V^

* 列表进入详情页的传参问题。

* 本地开发环境请求服务器接口跨域的问题

* axios封装和api接口的统一管理
* UI库的按需加载

* 如何优雅的只在当前页面中覆盖ui库中组件的样式
* 定时器问题

* rem文件的导入问题

* Vue-Awesome-Swiper基本能解决你所有的轮播需求

* 打包后生成很大的.map文件的问题

* fastClick的300ms延迟解决方案

* 组件中写选项的顺序

* 路由懒加载（也叫延迟加载）

* 开启gzip压缩代码

* 详情页返回列表页缓存数据和浏览位置、其他页面进入列表页刷洗数据的实践

* css的scoped私有作用域和深度选择器
* hiper打开速度测试
* vue数据的两种获取方式+骨架屏
* 自定义组件（父子组件）的双向数据绑定
* 路由的拆分管理
* mixins混入简化常见操作
* 打包之后文件、图片、背景图资源不存在或者路径错误的问题

* vue插件的开发、发布到github、设置展示地址、发布npm包

===========================这是华丽丽的分割线~~=========================

## 列表进入详情页的传参问题。 ##

例如商品列表页面前往商品详情页面，需要传一个商品id;

` <router-link :to= "{path: 'detail', query: {id: 1}}" >前往detail页面</router-link> 复制代码`

c页面的路径为 ` http://localhost:8080/#/detail?id=1` ，可以看到传了一个参数id=1，并且就算刷新页面id也还会存在。此时在c页面可以通过id来获取对应的详情数据，获取id的方式是 ` this.$route.query.id`

vue传参方式有：query、params+动态路由传参。

说下两者的区别：

1.query通过path切换路由，params通过name切换路由

` // query通过path切换路由 <router-link :to= "{path: 'Detail', query: { id: 1 }}" >前往Detail页面</router-link> // params通过name切换路由 <router-link :to= "{name: 'Detail', params: { id: 1 }}" >前往Detail页面</router-link> 复制代码`

2.query通过 ` this.$route.query` 来接收参数，params通过this.$route.params来接收参数。

` // query通过this. $route.query接收参数 created () { const id = this. $route.query.id; } // params通过this. $route.params来接收参数 created () { const id = this. $route.params.id; } 复制代码`

3.query传参的url展现方式：/detail?id=1&user=123&identity=1&更多参数

params＋动态路由的url方式：/detail/123

4.params动态路由传参，一定要在路由中定义参数，然后在路由跳转的时候必须要加上参数，否则就是空白页面：

` { path: '/detail/:id' , name: 'Detail' , component: Detail }, 复制代码`

注意，params传参时，如果没有在路由中定义参数，也是可以传过去的，同时也能接收到，但是一旦刷新页面，这个参数就不存在了。这对于需要依赖参数进行某些操作的行为是行不通的，因为你总不可能要求用户不能刷新页面吧。 例如：

` // 定义的路由中，只定义一个id参数 { path: 'detail/:id' , name: 'Detail' , components: Detail } // template中的路由传参， // 传了一个id参数和一个token参数 // id是在路由中已经定义的参数，而token没有定义 <router-link :to= "{name: 'Detail', params: { id: 1, token: '123456' }}" >前往Detail页面</router-link> // 在详情页接收 created () { // 以下都可以正常获取到 // 但是页面刷新后，id依然可以获取，而token此时就不存在了 const id = this. $route.params.id; const token = this. $route.params.token; } 复制代码`

## 本地开发环境请求服务器接口跨域的问题 ##

![](https://user-gold-cdn.xitu.io/2018/6/7/163d84c797b26f89?imageView2/0/w/1280/h/960/ignore-error/1)

上面的这个报错大家都不会陌生，报错是说没有访问权限（跨域问题）。本地开发项目请求服务器接口的时候，因为客户端的同源策略，导致了跨域的问题。

下面先演示一个没有配置允许本地跨域的的情况：

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e483b2e9e124?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e4bdd3f14fe0?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e4d03e31bc15?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到，此时我们点击获取数据，浏览器提示我们跨域了。所以我们访问不到数据。

那么接下来我们演示设置允许跨域后的数据获取情况：

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e53a26f6a946?imageView2/0/w/1280/h/960/ignore-error/1)

**注意：配置好后一定要关闭原来的server，重新 ` npm run dev` 启动项目。不然无效。**

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e53da5e7d37a?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e53f181d2722?imageView2/0/w/1280/h/960/ignore-error/1)

我们在1出设置了允许本地跨域，在2处，要注意我们访问接口时，写的是 ` /api` ，此处的 ` /api` 指代的就是我们要请求的接口域名。如果我们不想每次接口都带上 ` /api` ，可以更改axios的默认配置 ` axios.defaults.baseURL = '/api';` 这样，我们请求接口就可以直接 ` this.$axios.get('app.php?m=App&c=Index&a=index')` ，很简单有木有。此时如果你在network中查看xhr请求，你会发现显示的是localhost:8080/api的请求地址。这样没什么大惊小怪的，代理而已：

![](https://user-gold-cdn.xitu.io/2018/7/6/1646e59bebcbb788?imageView2/0/w/1280/h/960/ignore-error/1)

好了，最后附上proxyTable的代码：

` proxyTable: { // 用‘/api’开头，代理所有请求到目标服务器 '/api' : { target: 'http://jsonplaceholder.typicode.com' , // 接口域名 changeOrigin: true , // 是否启用跨域 pathRewrite: { // '^/api' : '' } } } 复制代码`

**注意：配置好后一定要关闭原来的server，重新 ` npm run dev` 启动项目。不然无效。**

**
**

## axios封装和api接口的统一管理 ##

axios的封装，主要是用来帮我们进行请求的拦截和响应的拦截。

在请求的拦截中我们可以携带userToken，post请求头、qs对post提交数据的序列化等。

在响应的拦截中，我们可以进行根据状态码来进行错误的统一处理等等。

axios接口的统一管理，是做项目时必须的流程。这样可以方便我们管理我们的接口，在接口更新时我们不必再返回到我们的业务代码中去修改接口。

由于这里内容稍微多一些，放在另一篇文章， [这里送上链接]( https://juejin.im/post/5b55c118f265da0f6f1aa354 ) 。

## UI库的按需加载: ##

为什么要使用按需加载的方式而不是一次性全部引入，原因就不多说了。这里以vant的按需加载为例，演示vue中ui库怎样进行按需加载：

* 安装： ` cnpm i vant -S`
* 安装babel-plugin-import插件使其按需加载： ` cnpm i babel-plugin-import -D`
* 在 .babelrc文件中中添加插件配置 ：

` libraryDirectory { "plugins" : [ // 这里是原来的代码部分 // ………… // 这里是要我们配置的代码 [ "import" , { "libraryName" : "vant" , "libraryDirectory" : "es" , "style" : true } ] ] } 复制代码`

* 在main.js中按需加载你需要的插件：

` // 按需引入vant组件 import { DatetimePicker, Button, List } from 'vant' ; 复制代码`

* 使用组件：

` // 使用vant组件 Vue.use(DatetimePicker) .use(Button) .use(List); 复制代码`

* 最后在在页面中使用：

` <van-button type = "primary" >按钮</van-button> 复制代码`

ps：出来vant库外，像antiUi、elementUi等，很多ui库都支持按需加载，可以去看文档，上面都会有提到。基本都是通过安装babel-plugin-import插件来支持按需加载的，使用方式与vant的如出一辙，可以去用一下。

## 如何优雅的只在当前页面中覆盖ui库中组件的样式 ##

首先我们vue文件的样式都是写在<style lang="less" scoped></style>标签中的，加scoped是为了使得样式只在当前页面有效。那么问题来了，看图：

![](https://user-gold-cdn.xitu.io/2018/7/6/1646eb217ea4d881?imageView2/0/w/1280/h/960/ignore-error/1)

我们正常写的所有样式，都会被加上[data-v-23d425f8]这个属性（如1所示），但是第三方组件内部的标签并没有编译为附带[data-v-23d425f8]这个属性。所以，我们想修改组件的样式，就没辙了。怎么办呢，有些小伙伴给第三方组件写个class，然后在一个公共的css文件中或者在当前页面再写一个没有socped属性的style标签，然后直接在里面修改第三方组件的样式。这样不失为一个方法，但是存在全局污染和命名冲突的问题。约定特定的命名方式，可以避免命名冲突。但是还是不够优雅。

作为一名优（ **强** ）秀（ **迫** ）的（ **症** ）前（ **患** ）端（ **者** ），怎么能允许这种情况出现呢？好了，下面说下优雅的解决方式：

通过深度选择器解决。例如修改上图中组件里的 ` van-ellipsis` 类的样式，可以这样做：

`.van-tabs /deep/ .van-ellipsis { color: blue}; 复制代码`

编译后的结果就是：

![](https://user-gold-cdn.xitu.io/2018/7/6/1646ec125eb1d457?imageView2/0/w/1280/h/960/ignore-error/1)

这样就不会给van-ellipsis也添加[data-v-23d425f8]属性了。至此你可以愉快的修改第三方组件的样式了。

当然了这里的深度选择器 ` /deep/` 是因为我用的 ` less` 语言，如果你没有使用 ` less/sass` 等，可以用 ` >>>` 符号。

更多的关于深度选择器的内容，在文章后面有介绍。

## 定时器问题: ##

我在a页面写一个定时，让他每秒钟打印一个1，然后跳转到b页面，此时可以看到，定时器依然在执行。这样是非常消耗性能的。如下图所示：

![](https://user-gold-cdn.xitu.io/2018/6/8/163dd3e4f06d86a3?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2018/6/8/163dd407f30494d3?imageView2/0/w/1280/h/960/ignore-error/1)

**解决方法1：**

首先我在data函数里面进行定义定时器名称：

` data () { return { timer: null // 定时器名称 } }, 复制代码`

然后这样使用定时器：

` this.timer = (() => { // 某些操作 }, 1000) 复制代码`

最后在beforeDestroy()生命周期内清除定时器：

` beforeDestroy () { clearInterval(this.timer); this.timer = null; } 复制代码`
方案1有两点不好的地方，引用尤大的话来说就是：

* 它需要在这个组件实例中保存这个 ` timer` ，如果可以的话最好只有生命周期钩子可以访问到它。这并不算严重的问题，但是它可以被视为杂物。
* 我们的建立代码独立于我们的清理代码，这使得我们比较难于程序化的清理我们建立的所有东西。

**解决方案2：**

该方法是通过$once这个事件侦听器器在定义完定时器之后的位置来清除定时器。以下是完整代码：

` const timer = set Interval(() =>{ // 某些定时器操作 }, 500); // 通过 $once 来监听定时器，在beforeDestroy钩子可以被清除。 this. $once ( 'hook:beforeDestroy' , () => { clearInterval(timer); }) 复制代码`

方案2要感谢@ [zzx18023]( https://juejin.im/user/57e62ac6a341310062442254 ) 在评论区提供出的解决方案。类似于其他需要在当前页面使用，离开需要销毁的组件（例如一些第三方库的picker组件等等），都可以使用此方式来解决离开后以后在背后运行的问题。

综合来说，我们更推荐使用 **方案2，使得代码可读性更强，一目了然。** 如果不清楚 ` $once、$on、$off` 的使用，这里送上官网的地址教程， [在程序化的事件侦听器那里]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-edge-cases.html%23%25E7%25A8%258B%25E5%25BA%258F%25E5%258C%2596%25E7%259A%2584%25E4%25BA%258B%25E4%25BB%25B6%25E4%25BE%25A6%25E5%2590%25AC%25E5%2599%25A8 ) 。

## rem文件的导入问题： ##

我们在做手机端时，适配是必须要处理的一个问题。例如，我们处理适配的方案就是通过写一个rem.js，原理很简单，就是根据网页尺寸计算html的font-size大小，基本上小伙伴们都知道，这里直接附上代码，不多做介绍。

` ;( function (c,d){var e=document.documentElement||document.body,a= "orientationchange" in window? "orientationchange" : "resize" ,b= function (){var f=e.clientWidth;e.style.fontSize=(f>=750)? "100px" :100*(f/750)+ "px" };b();c.addEventListener(a,b, false )})(window); 复制代码`

这里说下怎么引入的问题，很简单。在main.js中，直接 ` import './config/rem'` 导入即可。import的路径根据你的文件路径去填写。

# Vue-Awesome-Swiper基本能解决你所有的轮播需求 #

在我们使用的很多ui库（vant、antiUi、elementUi等）中，都有轮播组件，对于普通的轮播效果足够了。但是，某些时候，我们的轮播效果可能比较炫，这时候ui库中的轮播可能就有些力不从心了。当然，如果技术和时间上都还可以的话，可以自己造个比较炫的轮子。

这里我说一下vue-awesome-swiper这个轮播组件，真的非常强大，基本可以满足我们的轮播需求。swiper相信很多人都用过，很好用，也很方便我们二次开发，定制我们需要的轮播效果。vue-awesome-swiper组件实质上基于 ` swiper` 的，或者说就是能在vue中跑的swiper。下面说下怎么使用：

* 安装 ` cnpm install vue-awesome-swiper --save`
* 在组件中使用的方法，全局使用意义不大：

` // 引入组件 import 'swiper/dist/css/swiper.css' import { swiper, swiperSlide } from 'vue-awesome-swiper' // 在components中注册组件 components: { swiper, swiperSlide } // template中使用轮播 // ref是当前轮播 // callback是回调 // 更多参数用法，请参考文档 <swiper :options= "swiperOption" ref= "mySwiper" @someSwiperEvent= "callback" > <!-- slides --> <swiper-slide><div class= "item" >1</div></swiper-slide> <swiper-slide><div class= "item" >2</div></swiper-slide> <swiper-slide><div class= "item" >3</div></swiper-slide> <!-- Optional controls --> <div class= "swiper-pagination" slot= "pagination" ></div> <div class= "swiper-button-prev" slot= "button-prev" ></div> <div class= "swiper-button-next" slot= "button-next" ></div> <div class= "swiper-scrollbar" slot= "scrollbar" ></div> </swiper> 复制代码`

` // 参数要写在data中 data () { return { // swiper轮播的参数 swiperOption: { // 滚动条 scrollbar: { el: '.swiper-scrollbar' , }, // 上一张，下一张 navigation: { nextEl: '.swiper-button-next' , prevEl: '.swiper-button-prev' , }, // 其他参数………… } } }, 复制代码`

swiper需要配置哪些功能需求，自己根据文档进行增加或者删减。附上文档： [npm文档]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fvue-awesome-swiper ) ， [swiper3.0/4.0文档]( https://link.juejin.im?target=http%3A%2F%2Fwww.swiper.com.cn%2Fapi%2Findex.html ) ，更多用法，请参考文档说明。

## 打包后生成很大的.map文件的问题 ##

项目打包后，代码都是经过压缩加密的，如果运行时报错，输出的错误信息无法准确得知是哪里的代码报错。 而生成的.map后缀的文件，就可以像未加密的代码一样，准确的输出是哪一行哪一列有错可以通过设置来不生成该类文件。但是我们在生成环境是不需要.map文件的，所以可以在打包时不生成这些文件：

在config/index.js文件中，设置 ` productionSourceMap: false` ,就可以不生成.map文件

![](https://user-gold-cdn.xitu.io/2018/6/8/163de300c977f7cd?imageView2/0/w/1280/h/960/ignore-error/1)

## fastClick的300ms延迟解决方案 ##

开发移动端项目，点击事件会有300ms延迟的问题。至于为什么会有这个问题，请自行百度即可。这里只说下常见的解决思路，不管vue项目还是jq项目，都可以使用 ` fastClick` 解决。

安装 ` fastClick` :

` cnpm install fastclick -S 复制代码`

在main.js中引入 ` fastClick` 和初始化:

` import FastClick from 'fastclick' ; // 引入插件 FastClick.attach(document.body); // 使用 fastclick 复制代码`

## 组件中写选项的顺序 ##

为什么选项要有统一的书写顺序呢？很简单，就是要将选择和认知成本最小化。

* 

**副作用** (触发组件外的影响)

* ` el`

* 

**全局感知** (要求组件以外的知识)

* ` name`
* ` parent`

* 

**组件类型** (更改组件的类型)

* ` functional`

* 

**模板修改器** (改变模板的编译方式)

* ` delimiters`
* ` comments`

* 

**模板依赖** (模板内使用的资源)

* ` components`
* ` directives`
* ` filters`

* 

**组合** (向选项里合并属性)

* ` extends`
* ` mixins`

* 

**接口** (组件的接口)

* ` inheritAttrs`
* ` model`
* ` props` / ` propsData`

* 

**本地状态** (本地的响应式属性)

* ` data`
* ` computed`

* 

**事件** (通过响应式事件触发的回调)

* ` watch`
* 生命周期钩子 (按照它们被调用的顺序)

* ` beforeCreate`
* ` created`
* ` beforeMount`
* ` mounted`
* ` beforeUpdate`
* ` updated`
* ` activated`
* ` deactivated`
* ` beforeDestroy`
* ` destroyed`

* 

**非响应式的属性** (不依赖响应系统的实例属性)

* ` methods`

* 

**渲染** (组件输出的声明式描述)

* ` template` / ` render`
* ` renderError`

## ##

## ##

## ##

## 查看打包后各文件的体积，帮你快速定位大文件 ##

如果你是vue-cli初始化的项目，会默认安装 ` webpack-bundle-analyzer` 插件，该插件可以帮助我们查看项目的体积结构对比和项目中用到的所有依赖。也可以直观看到各个模块体积在整个项目中的占比。很霸道有木有~~

![](https://user-gold-cdn.xitu.io/2018/6/12/163f310014cc9b59?imageslim)

` npm run build --report // 直接运行，然后在浏览器打开http://127.0.0.1:8888/即可查看 复制代码`

**记得运行的时候先把之前** ` **npm run dev**` **开启的本地关掉**

## 路由懒加载（也叫延迟加载） ##

路由懒加载可以帮我们在进入首屏时不用加载过度的资源，从而减少首屏加载速度。

路由文件中，

非懒加载写法：

` import Index from '@/page/index/index' ; export default new Router({ routes: [ { path: '/' , name: 'Index' , component: Index } ] }) 复制代码`
路由懒加载写法：

` export default new Router({ routes: [ { path: '/' , name: 'Index' , component: resolve => require([ '@/view/index/index' ], resolve) } ] }) 复制代码`

## 开启gzip压缩代码 ##

spa这种单页应用，首屏由于一次性加载所有资源，所有首屏加载速度很慢。解决这个问题非常有效的手段之一就是前后端开启gizp（其他还有缓存、路由懒加载等等）。gizp其实就是帮我们减少文件体积，能压缩到30%左右，即100k的文件gizp后大约只有30k。

vue-cli初始化的项目中，是默认有此配置的，只需要开启即可。但是需要先安装插件：

` // 2.0的版本设置不一样，本文写作时为v1版本。v2需配合vue-cli3cnpm i compression-webpack-plugin@1.1.11 复制代码`

然后在config/index.js中开启即可:

` build: { // 其他代码 ………… productionGzip: true , // false 不开启gizp， true 开启 // 其他代码 } 复制代码`

现在打包的时候，除了会生成之前的文件，还是生成.gz结束的gzip过后的文件。具体实现就是如果客户端支持gzip，那么后台后返回gzip后的文件，如果不支持就返回正常没有gzip的文件。

****注意** ：这里前端进行的打包时的gzip，但是还需要后台服务器的配置。配置是比较简单的，配置几行代码就可以了，一般这个操作可以叫运维小哥哥小姐姐去搞一下，没有运维的让后台去帮忙配置。

## 详情页返回列表页缓存数据和浏览位置、其他页面进入列表页刷新数据的实践 ##

这样一个场景：有三个页面，首页/或者搜索页，商品分类页面，商品详情页。我们希望从首页进入分类页面时，分类页面要刷新数据，从分类进入详情页再返回到分类页面时，我们不希望刷新，我们希望此时的分类页面能够缓存已加载的数据和自动保存用户上次浏览的位置。之前在百度搜索的基本都是keep-alive处理的，但是总有那么一些不完善，所以自己在总结了之后进行了如下的实践。

解决这种场景需求我们可以通过vue提供的keepAlive属性。这里直接送上另一篇 [处理这个问题的传送门]( https://juejin.im/post/5b2ce07ce51d45588a7dbf76 ) 吧

## CSS的coped私有作用域和深度选择器 ##

大家都知道当 ` <style>` 标签有 ` scoped` 属性时，它的 CSS 只作用于当前组件中的元素。那么他是怎么实现的呢，大家看一下编译前后的代码就明白了：

编译前：

` <style scoped> .example { color: red; } </style> 复制代码`

编译后：

` <style> .example[data-v -f 3f3eg9] { color: red; } 复制代码`

看完你肯定就会明白了，其实是在你写的组件的样式，添加了一个属性而已，这样就实现了所谓的私有作用域。但是也会有弊端，考虑到浏览器渲染各种 CSS 选择器的方式，当 ` p { color: red }` 设置了作用域时 (即与特性选择器组合使用时) 会慢很多倍。如果你使用 class 或者 id 取而代之，比如 `.example { color: red }` ，性能影响就会消除。所以，在你的样式里，进来避免直接使用标签，取而代之的你可以给标签起个class名。

如果你希望 ` scoped` 样式中的一个选择器能够作用得“更深”，例如影响子组件，你可以使用 ` >>>` 操作符:

` <style scoped> .parent >>> .child { /* ... */ } </style> 复制代码`

上述代码将会编译成：

`.parent[data-v -f 3f3eg9] .child { /* ... */ } 复制代码`

而对于less或者sass等预编译，是不支持 ` >>>` 操作符的，可以使用 ` /deep/` 来替换 ` >>>` 操作符，例如：.parent /deep/ .child { /* ... */ }

==================================

后面会继续更新：

* axios封装和api接口的统一管理（已更新，在上面的链接）
* hiper打开速度测试
* vue数据的两种获取方式+骨架屏
* 自定义组件（父子组件）的双向数据绑定
* 路由的拆分管理
* mixins混入简化常见操作
* 打包之后文件、图片、背景图资源不存在或者路径错误的问题

* vue插件的开发、发布到github、设置展示地址、发布npm包

------------华丽丽的分割线-------------------------华丽丽的分割线-------------------------华丽丽的分割线-------------------------华丽丽的分割线-------------------------华丽丽的分割线-------------------------华丽丽的分割线-------------------------华丽丽的分割线-------------

## Hiper：一款令人愉悦的性能分析工具 ##

![](https://user-gold-cdn.xitu.io/2018/8/15/1653c6f5c7382367?imageView2/0/w/1280/h/960/ignore-error/1)

如上图，是hiper工具的测试结果，从中我们可以看到DNS查询耗时、TCP连接耗时、第一个Byte到达浏览器的用时、页面下载耗时、DOM Ready之后又继续下载资源的耗时、白屏时间、DOM Ready 耗时、页面加载总耗时。

**在我们的编辑器终端中全局安装：**

` cnpm install hiper -g 复制代码`

**使用：** **终端输入命令：hiper 测试的网址**

****

` # 当我们省略协议头时，默认会在url前添加`https://` # 最简单的用法 hiper baidu.com # 如何url中含有任何参数，请使用双引号括起来 hiper "baidu.com?a=1&b=2" # 加载指定页面100次 hiper -n 100 "baidu.com?a=1&b=2" # 禁用缓存加载指定页面100次 hiper -n 100 "baidu.com?a=1&b=2" --no-cache # 禁JavaScript加载指定页面100次 hiper -n 100 "baidu.com?a=1&b=2" --no-javascript # 使用GUI形式加载指定页面100次 hiper -n 100 "baidu.com?a=1&b=2" -H false # 使用指定useragent加载网页100次 hiper -n 100 "baidu.com?a=1&b=2" -u "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36" 复制代码`

这段用法示例，我直接拷贝的文档说明，具体的可以看下文档， [这里送上链接]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpod4g%2Fhiper%2Fblob%2Fmaster%2FREADME.zh-CN.md ) 。当我们项目打开速度慢时，这个工具可以帮助我们快速定位出到底在哪一步影响的页面加载的速度。

平时我们查看性能的方式，是在performance和network中看数据，记录下几个关键的性能指标，然后刷新几次再看这些性能指标。有时候我们发现，由于样本太少，受当前「网络」、「CPU」、「内存」的繁忙程度的影响很重，有时优化后的项目反而比优化前更慢。

如果有一个工具，一次性地请求N次网页，然后把各个性能指标取出来求平均值，我们就能非常准确地知道这个优化是「正优化」还是「负优化」。

hiper就是解决这个痛点的。

## vue获取数据的两种方式的实践+简单骨架屏实现 ##

在vue中获取数据有两种方式，引入尤大大的话就是：

* 

**导航完成之后获取** ：先完成导航，然后在接下来的组件生命周期钩子中获取数据。在数据获取期间显示“加载中”之类的指示。

* 

**导航完成之前获取** ：导航完成前，在路由进入的守卫中获取数据，在数据获取成功后执行导航。

从技术角度讲，两种方式都不错 —— 就看你想要的用户体验是哪种。那么我们来实践一下这两种获取数据的方式，以及用户体验优化的一点思考。

**一、首先是第一种：导航完成之后获取，** 这种方式是我们大部分都在使用的，（因为可能一开始我们只知道这种方式^V^）。使用这种方式时，我们会马上导航和渲染组件，然后在组件的 ` created` 钩子中获取数据。这让我们有机会在数据获取期间展示一个 loading 状态，还可以在不同视图间展示不同的 loading 状态。获取数据大家都会，这里说下用户体验的一些东西：

* 在数据获取到之前，页面组件已经加载，但是数据没有拿到并渲染，所以在此过程中，我们不能加载页面内展示数据的那块组件，而是要有一个loading的加载中的组件或者骨架屏。

* 当页面数据获取失败，可以理解为请求超时的时候，我们要展示的是断网的组件。

* 如果是列表页，还要考虑到空数据的情况，即为空提示的组件。

那么，我们的页面是要有这基本的三个部分的，放代码：

` <template> <div class= "list" > <!--加载中或者骨架屏--> <div v-if= "loading" > </div> <!--请求失败，即断网的提示组件--> <div v-if= "error" > </div> <!--页面内容--> <div v-if= "requestFinished" class= "content" > <!--页面内容--> <div v-if= "!isEmpty" > <!--例如有个列表，当然肯定还会有其他内容--> <ul></ul> </div> <!--为空提示组件--> <div v-else>空空如也</div> </div> </div> </template> 复制代码`

这种获取数据的情况下，我们进来默认的是展示loading或者骨架屏的内容，然后如果获取数据失败（即请求超时或者断网），则加载error的那个组件，隐藏其他组件。如果数据请求成功，则加载内容的组件，隐藏其他组件。如果是列表页，可能在内容组件中还会有列表和为空提示两块内容，所以这时候也还要根据获取的数据来判断是加载内容还是加载为空提示。

**二、第二种方式：导航完成之前获取**

这种方式是在页面的 beforeRouteEnter 钩子中请求数据，只有在数据获取成功之后才会跳转导航页面。

` beforeRouteEnter (to, from, next) { api.article.articleDetail(to.query.id).then(res=> { next(vm => { vm.info = res.data; vm.loadFinish = true }) }) }, 复制代码`

1. 大家都知道钩子中 ` beforeRouteEnter` 钩子中this还不能使用，所以要想进行赋值操作或者调用方法，我们只能通过在next()方法的回调函数中处理，这个回调函数的第一个参数就代表了this，他会在组件初始化成功后进行操作。

2. 我想，很多时候我们的api或者axios方法都是挂载到vue的原型上的，由于这里使用不了this，所以只能在页面组件内引入api或者我们的axios。

3. 赋值操作也可以写在method方法中，但是调用这个赋值方法还是 ` vm.yourFunction()` 的方式。

4. 为空提示、断网处理等都和第一种方式一样，但是，由于是先获取到数据之后再跳转加载组件的，所以我们不需要在预期的页面内展示骨架屏或者loading组件。可以，我们需要在当前页面进入之前，即在上一个页面的时候有一个加载的提示，比如页面顶部的进度条。这样用户体验就比较友好了，而不至于因为请求的s速度慢一些导致半天没反应而用户又不知道的结果。全局的页面顶部进度条，可以在main.js中通过router.beforeEach(to, from, next) {}来设置，当页面路由变化时，显示页面顶部的进度条，进入新路由后隐藏掉进度条。

![](https://user-gold-cdn.xitu.io/2018/8/15/1653cfc386e0cf1c?imageView2/0/w/1280/h/960/format/png/ignore-error/1)

关于怎么添加进度条，因为在另一篇文章已经写了， [这里直接送上链接吧]( https://juejin.im/post/5b31e07ef265da599c56165a ) ，就不再重复浪费地方了。操作也比较简单，可自行查阅。

其实说到了这里，那么 **骨架屏** 的事情也就顺带已经解决了，一般页面骨架屏也就是一张页面骨架的图片，但是要注意这张图片要尽可能的小。

## 自定义组件（父子组件）的双向数据绑定 ##

说到父子组件的通信，大家一定都不陌生了：父组件通过props向子组件传值，子组件通过emit触发父组件自定义事件。但是这里要说的是父子组件使用v-model实现的通信。相信大家在使用别人的组件库的时候，经常是通过v-model来控制一个组件显示隐藏的效果等，例如弹窗。下面就一步一步解开v-model的神秘面纱。抓~~稳~~喽~~，老司机弯道要踩油门了~~~

提到v-model首先想到的就是我们对于表单用户数据的双向数据绑定，操作起来很简洁很粗暴，例如：

` <input type = "text" v-model= "msg" > data () { return { msg: '' } } 复制代码`

其实v-model是个语法糖，上面这一段代码和下面这一段代码是一样的效果：

` <input type = "text" :value= "msg" @input= "msg = $event.target.value" > data () { return { msg: '' } }, 复制代码`

由此可以看出， ` v-model="msg"` 实则是 ` :value="msg" @input="msg = $event.target.value"` 的语法糖。这里其实就是监听了表单的 ` input` 事件，然后修改 ` :value` 对应的值。除了在输入表单上面可以使用v-model外，在组件上也是可以使用的， [这点官网有提到]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-custom-events.html%23%25E8%2587%25AA%25E5%25AE%259A%25E4%25B9%2589%25E7%25BB%2584%25E4%25BB%25B6%25E7%259A%2584-v-model ) ，但是介绍的不是很详细，导致刚接触的小伙伴会有一种云里雾里不知所云的感觉。既然了解了v-model语法糖本质的用法，那么我们就可以这样实现父子组件的双向数据绑定：

**以上原理实现方法，写法1：**

父组件用法：

` <empty v-model= "msg" ></empty> 复制代码`

子组件写法：

` // 点击该按钮触发父子组件的数据同步 <div class= "share-btn" @click= "confirm" >确定</div> // 接收父组件传递的value值 // 注意，这种实现方法，这里只能使用value属性名 props: { value: { type : Boolean, default: false } }, methods: { confirm () { // 双向数据绑定父组件:value对应的值 // 通过 $emit 触发父组件input事件，第二个参数为传递给父组件的值，这里传递了一个 false 值 // 可以理解为最上面展示的@input= "msg = $event.target.value" 这个事件 // 即触发父组件的input事件，并将传递的值‘ false ’赋值给msg this. $emit ( 'input' , false ) } } 复制代码`

这种方式实现了父子组件见v-model双向数据绑定的操作，例如你可以试一下实现一个全局弹窗组件的操作，通过v-model控制弹窗的显示隐藏，因为你要在页面内进行某些操作将他显示出来，控制其隐藏的代码是写在组件里面的，当组件隐藏了对应的也要父组件对应的值改变。

**以上这种方式实现的父子组件的v-model通信，虽可行，但限制了我们必须popos接收的属性名为value和emit触发的必须为input，这样就容易有冲突，特别是在表单里面。所以，为了更优雅的使用v-model通信而解决冲突的问题，我们可以通过在子组件中使用** ` **model**` **选项，下面演示写法2：**

父组件写法：

` <empty v-model= "msg" ></empty> 复制代码`

子组件写法：

` <div class= "share-btn" @click= "confirm" >确定</div> // model选项用来避免冲突 // prop属性用来指定props属性中的哪个值用来接收父组件v-model传递的值 // 例如这里用props中的show来接收父组件传递的v-model值 // event：为了方便理解，可以简单理解为父组件@input的别名，从而避免冲突 // event的值对应了你emit时要提交的事件名，你可以叫aa，也可以叫bb，但是要命名要有意义哦！！！ model: { prop: 'show' , event: 'changed' }, props: { // 由于model选项中的prop属性指定了，所以show接收的是父组件v-model传递的值 show: { type : Boolean, default: false } }, methods: { confirm () { // 双向数据绑定父组件传递的值 // 第一个参数，对应model选项的event的值，你可以叫aa，bbb，ccc，起名随你 this. $emit ( 'changed' , false ) } } 复制代码`

这种实现父子组件见v-model绑定值的方法，在我们开发中其实是很常用的，特别是你要封装公共组件的时候。

**最后，实现双向数据绑定的方式其实还有** ` **.sync**` **，这个属性一开始是有的，后来由于被认为或破坏单向数据流被删除了，但最后证明他还是有存在意义的，所以在2.3版本又加回来了。**

例如：父组件：

` <empty :oneprop.sync= "msg" ></empty> data () { return { msg: '' } } 复制代码`

子组件：

` <div class= "share-btn" @click= "changeMsg" >改变msg值</div> props: { oneprop: { type : String, default: 'hello world' } }, methods: { changeMsg () { // 双向数据流 this. $emit ( 'update:msg' , 'helow world' ) } } 复制代码`

这样，便可以在子组件更新父组件的数据。由于 ` v-model` 只使用一次，所以当需要双向绑定的值有多个的时候， `.sync` 还是有一定的使用场景的。.sync是下面这种写法的语法糖，旨在简化我们的操作：

` <empty :msg= "message" @update:msg= "message = $event " ></empty> 复制代码`

掌握了组件的v-model写法，在封装一些公共组件的时候就又轻松一些了吧。

这里再提一下：

* ` vm.$emit(event ,[...args])` 这个api，其主要作用就是用来触发当前实例上的事件。附加参数都会传给监听器回调。子组件也属于当前实例。第一个参数：要触发的事件名称。后续的参数可选：即作为参数传递给要触发的事件。 [文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23vm-emit )

* 监听当前实例上的自定义事件，事件可以有$emit触发，也能 **通过hook监听到钩子函数，**

vm.$on( event, callback )：一直监听； [文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23vm-on )

vm.$once( event, callback )：监听一次； [文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23vm-once )

vm.$off( [event, callback] )：移除监听； [文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23vm-off )

监听$emit触发的自定义事件，上面已经有过用法了，监听钩子函数，在上面的定时器那块也有演示到。监听钩子函数的场景使用的不多，但是还是要知道的。

* vm.$attrs：可以获取到父组件传递的除class和style外的所有自定义属性。

* vm.$listeners：可以获取到父组件传递的所有自定义事件

例如：父组件:

` <empty :msg= "message" :title= "articleTitle" @confirm= "func1" @cancel= "func2" ></empty> 复制代码`

就可以在子组件中获取父组件传递的属性和事件，而不用在props中定义。子组件简单演示如下：

` created () { const msg = this. $attrs.msg; // 获取父组件传递的msg this. $listeners.confirm && this. $listeners.confirm(); //若组件传递事件confirm则执行 }, 复制代码`

这在我们写一些高级组件时候，会有用到的。

## 路由拆分管理 ##

这里说的路由拆分指的是将路由的文件，按照模块拆分，这样方便路由的管理，更主要的是方便多人开发。具体要不要拆分，那就要视你的项目情况来定了，如果项目较小的话，也就一二十个路由，那么是拆分是非常没必要的。但倘若你开发一些功能点较多的商城项目，路由可以会有一百甚至几百个，那么此时将路由文件进行拆分是很有必要的。不然，你看着index.js文件中一大长串串串串串串的路由，也是很糟糕的。

![](https://user-gold-cdn.xitu.io/2018/9/7/165b1fab48fba445?imageView2/0/w/1280/h/960/ignore-error/1)

首先我们在router文件夹中创建一个index.js作为路由的入口文件，然后新建一个modules文件夹，里面存放各个模块的路由文件。例如这里储存了一个vote.js投票模块的路由文件和一个公共模块的路由文件。下面直接上index.js吧，而后在简单介绍：

` import Vue from 'vue' import Router from 'vue-router' // 公共页面的路由文件 import PUBLIC from './modules/public' // 投票模块的路由文件 import VOTE from './modules/vote' Vue.use(Router) // 定义路由 const router = new Router({ mode: 'history' , routes: [ ...PUBLIC, ...VOTE, ] }) // 路由变化时 router.beforeEach((to, from, next) => { if (document.title !== to.meta.title) { document.title = to.meta.title; } next() }) // 导出 export default router 复制代码`

首先引入vue和router最后导出，这就不多说了，基本的操作。

这里把 router.beforeEach 的操作写了router的index.js文件中，有些人可能会写在main.js中，这也没有错，只不过，个人而言，既然是路由的操作，还是放在路由文件中管理更好些。这里就顺便演示了，如何在页面切换时，自动修改页面标题的操作。

而后引入你根据路由模块划分的各个js文件，然后在实例化路由的时候，在routes数组中，将导入的各个文件通过结构赋值的方法取出来。最终的结果和正常的写法是一样的。

然后看下我们导入的vote.js吧：

` /** * 投票模块的router列表 */ export default [ // 投票模块首页 { path: '/vote/index' , name: 'VoteIndex' , component: resolve => require([ '@/view/vote/index' ], resolve), meta: { title: '投票' } }, // 详情页 { path: '/vote/detail' , name: 'VoteDetail' , component: resolve => require([ '@/view/vote/detail' ], resolve), meta: { title: '投票详情' } }] 复制代码`

这里就是将投票模块的路由放在一个数组中导出去。整个路由拆分的操作，不是vue的知识，就是一个es6导入导出和结构的语法。具体要不要拆分，还是因项目和环境而异吧。

这里的路由用到了懒加载路由的方式，如果不清楚，文字上面有介绍到。

还有这里的meta元字段中，定义了一个title信息，用来存储当前页面的页面标题，即document.title。

## ##

## mixins混入简化常见操作 ##

我们在开发中经常会遇到金钱保留两位小数，时间戳转换等操作。每次我们会写成一个公共函数，然后在页面里面的filters进行过滤。这种方法每次，但是感觉每次需要用到，都要写一遍在filters，也是比较烦呢！！！但是，我们猿类的极致追究就是懒呀，那这怎么能行~~~

兄弟们，抄家伙！上mixins！！！

` import { u_fixed } from './tool' const mixins = { filters: { // 保留两位小数 mixin_fixed2 (val) { return u_fixed(val) }, // 数字转汉字，16000 => 1.60万 mixin_num2chinese (val) { return val > 9999 ? u_fixed(val/10000) + '万' : val; } }} export default mixins 复制代码`

新建一个mixins.js，把我们需要混入的内容都写在里面，例如这里混入了filters，把常用的几个操作写在了里面，大家可以自行扩展。

这样的话，在我们需要的页面import这个js，然后声明一下混入就好，而后就可以像正常的方式去使用就好了。

![](https://user-gold-cdn.xitu.io/2018/9/7/165b357d3179206d?imageView2/0/w/1280/h/960/ignore-error/1)

例如，我现在可以直接在页面内使用我们的过滤操作 ` {{1000 | mixin_fixed2}}`

## 打包之后文件、图片、背景图资源不存在或者路径错误的问题 ##

![](https://user-gold-cdn.xitu.io/2018/9/7/165b365ffc00827b?imageView2/0/w/1280/h/960/ignore-error/1)

先看下项目的config文件夹下的index.js文件，这个配置选项就好使我们打包后的资源公共路径，默认的值为 ` ‘/’` ，即根路径，所以打包后的资源路径为根目录下的static。由此问题来了，如果你打包后的资源没有放在服务器的根目录，而是在根目录下的mobile等文件夹的话，那么打包后的路径和你代码中的路径就会有冲突了，导致资源找不到。

所以，为了解决这个问题，你可以在打包的时候把上面这个路径由‘/’的根目录，改为 ` ‘./’` 的相对路径。

![](https://user-gold-cdn.xitu.io/2018/9/7/165b3affdd1725d7?imageView2/0/w/1280/h/960/ignore-error/1)

## ##

这样的的话，打包后的图片啊js等路径就是 ` ‘./static/img/asc.jpg’` 这样的相对路径，这就不管你放在哪里，都不会有错了。但是，凡是都有但是~~~~~这里一切正常，但是背景图的路径还是不对。因为此时的相对就变成了 ` static/css/` 文件夹下的 ` static/img/xx.jpg` ，但是实际上 ` static/css/` 文件夹下没有 ` static/img/xx.jpg` ，即 ` static/css/static/img/xx.jpg` 是不存在的。此时相对于的当前的css文件的路径。所以为了解决这个问题，要把我们css中的背景图的加个公共路径 ` ‘../../’` ，即让他往上返回两级到和 ` index.html` 文件同级的位置，那么此时的相对路径 ` static/img/xx.jpg` 就能找到对应的资源了。那么怎么修改背景图的这个公共路径呢，因为背景图是通过loader解析的，所以自然在loader的配置中修改，打开build文件夹下的utils文件，找到exports.cssLoaders的函数，在函数中找到对应下面这些配置： ![](https://user-gold-cdn.xitu.io/2018/9/7/165b3bd6ba562f0a?imageView2/0/w/1280/h/960/ignore-error/1)

找到这个位置，添加一上配置，就是上图红框内的代码，就可以把它的公共路径修改为往上返回两级。这样再打包看下，就ok了！

**最后再郑重说一点，如果你的路由模式是history的，那么打包放在服务器，必须要后台服务器的配合，具体的可以看官方文档，这点很重要。不然你会发现白屏啊等各种莫名其妙的问题。牢记！！！**

**
**

## vue插件的开发、发布到github、设置展示地址、发布npm包 ##

对于平时我们常用的一些组件，我们可以把它封装成插件，然后发布到github上，最后再发布成npm包，这样以后便可以直接从npm安装插件到我们的项目中，省去了我们拷贝的过程了，还能给别人分享呢！

由于插件的这一块内容比较多，我暂且放在另外一篇文章吧， [这里呢就附上链接吧]( https://juejin.im/post/5b96586de51d450e7d0984a6 ) 。