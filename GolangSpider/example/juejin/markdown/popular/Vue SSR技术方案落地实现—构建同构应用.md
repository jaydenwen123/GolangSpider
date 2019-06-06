# Vue SSR技术方案落地实现—构建同构应用 #

## Vue SSR技术方案落地实现—构建同构应用 ##

### 一、基本知识扫盲 ###

#### 1、何为服务器端渲染？ ####

**1)、服务器端渲染** ：这种技术方案在前端领域处于蛮荒时代就已出现，当时的解决方案主要是后台开发通过模板引擎来设计（如：Java Web的JSP）；简而言之，就是模板页面在后台获取到数据并填充，然后响应返回模板页面html字符串给浏览器渲染；

#### 2、为什么要使用服务器端渲染？ ####

**1)、优化SEO** ：它主要解决搜索引擎SEO优化（Ajax异步请求是SEO优化的一大阻力；例如：去超市买A商品（数据），正好A商品没有，需要等待一段时间去仓库调取；作为消费者（搜索引擎）你能等吗？）；

**2)、优化单页应用的首屏加载时间** ：现今单页面应用大行其道，单页面应用解决了页面无感加载，但是带来了首屏加载缓慢；通过服务器端渲染机制可以很好的解决首屏页面加载问题。当然这不是唯一的解决方案（合理拆分成多页面应用也可以解决）；

#### 3、实现服务器端渲染技术方案有哪些？ ####

**1)、纯后台技术实现** ：利用后台语言模板引擎进行服务器端渲染方案落地。对于前端来说，可以利用node作为中间件，然后利用node的ejs模板引擎负责数据填充，最后通过node路由响应机制输出html字符串给客户端浏览器进行渲染；

**2)、构建同构应用** ： 同构应用就是可以同时运行在客户端和服务器端的Web应用； 这种一般采用webpack构建工具和开源工具进行实现；以下会以Vue来介绍同构应用；这种实现方式相对于上面的方案更复杂，开发难度大；但是可以享受到Vue框架带来的便利（响应式数据，路由无感切换等便利）。前提需要“客户端激活”。

> 
> 
> 
> **客户端激活：官方术语，可以理解为服务器端渲染成html字符串给浏览器之后，需要引入客户端的bundleClient文件，这个环节就交给客户端处理了；**
> 
> 
> 

#### 4、怎样利用Vue构建同构应用？ ####

**1)、准备什么？**

* Vue的运行环境：对于vue的运行环境不建议直接用vue-cli脚手架构建，因为脚手架工具继承度比较高，不好对配置进行扩展；利用webpack构建工具从零搭建vue运行环境。
* node后台服务环境：这里可以利用expres或者koa框架对后台部分进行构建。这里使用koa框架搭建的后台服务环境；
* 核心包vue-server-renderer：该包主要是将Vue组件渲染成html字符串，是同构应用的核心包。主要用于node后台服务环境；

**2)、注意什么？**

* 生命周期不同：Vue同构应用在客户端和服务器端的生命周期不一致，在客户端，可以使用全部Vue生命周期方法。在服务器端，只有beforeCreate和created这两个生命周期方法会执行，其余的方法不会执行；
* 实例方式不同：Vue同构应用在客户端只需要实例化一次，在服务器端需要根据每次请求重新实例化。主要是为了解决后台服务占用同一进程，容易导致状态污染问题（交叉请求）；
* 异步数据获取方式不同：在客户端，可以在mounted或created钩子方法中初始化请求数据；由于服务器端拥有created钩子方法，所以在created初始化请求数据是有效果的，但是不建议这么干，服务器端没有destory钩子方法。这样就导致资源得不到释放，很耗内存；

### 二、通过Demo深入学习 ###

**1、事先准备** ：工欲善其事必先安装包，话不多说；构建同步应用所需要的package.json包。这里不解释包的作用；

` "dependencies": { "@babel/core": "^7.4.5", "@babel/preset-env": "^7.4.5", "autoprefixer": "^9.5.1", "babel-loader": "^8.0.6", "babel-plugin-dynamic-import-webpack": "^1.1.0", "css-loader": "^2.1.1", "extract-text-webpack-plugin": "^3.0.2", "html-webpack-plugin": "^3.2.0", "koa": "^2.7.0", "koa-router": "^7.4.0", "koa-static": "^5.0.0", "mini-css-extract-plugin": "^0.7.0", "postcss-loader": "^3.0.0", "url-loader": "^1.1.2", "vue": "^2.6.10", "vue-loader": "^15.7.0", "vue-router": "^3.0.6", "vue-server-renderer": "^2.6.10", "vue-style-loader": "^4.1.2", "vue-template-compiler": "^2.6.10", "vuex": "^3.1.1", "vuex-router-sync": "^5.0.0", "webpack": "^4.32.2", "webpack-cli": "^3.3.2", "webpack-dev-server": "^3.4.1", "webpack-merge": "^4.2.1", "webpack-node-externals": "^1.7.2" } 复制代码`

在项目根目录下创建.babelrc文件

` { "presets": [ "@babel/preset-env" ], "plugins": [ // 支持路由动态加载的写法 const Foo = () => import('../components/Foo.vue') "dynamic-import-webpack" ] } 复制代码`

**2、webpack配置** ：由于同构应用同时支持客户端和服务器端，对于webpack配置要根据平台不同而做不同的配置（建议将公共配置抽离处理）。这里分了3个webpack配置文件：webpack.base.conf.js、webpack.client.conf.js、webpack.server.conf.js；

**webpack.base.conf.js**

` const path= require ( "path" ); const VueLoaderPlugin = require ( 'vue-loader/lib/plugin' ); //在webpack4.x版本中mini-css-extract-plugin插件代替extract-text-webpack-plugin插件 const MiniCssExtractPlugin = require ( "mini-css-extract-plugin" ); module.exports={ mode : "development" , output :{ path :path.resolve(__dirname, "../dist" ), filename : "[name].bundle.js" }, resolve : { extensions : [ '.js' , '.vue' ] }, module :{ rules :[ { test : /\.js$/ , use : 'babel-loader' },{ test : /\.vue$/ , use : 'vue-loader' },{ test : /\.(jpg|jpeg|png|gif|svg)$/ , use :{ loader : 'url-loader' , options : { limit : 20000 } } },{ test : /\.css$/ , use :[ { loader : MiniCssExtractPlugin.loader, options :{ publicPath :path.resolve(__dirname, "../dist" ) } }, "css-loader" , "postcss-loader" ] } ] }, plugins :[ new VueLoaderPlugin(), new MiniCssExtractPlugin({ filename : "[name].client.css" , chunkFilename : "[id].client.css" }) ] } 复制代码`

**webpack.client.conf.js**

` const path= require ( "path" ); const merge = require ( 'webpack-merge' ); const HtmlWebpackPlugin = require ( 'html-webpack-plugin' ); const base = require ( './webpack.base.config' ); const VueSSRClientPlugin = require ( "vue-server-renderer/client-plugin" ); const config=merge(base,{ entry :{ client :path.resolve(__dirname, "../src/entry-client.js" ) }, output :{ path :path.resolve(__dirname, "../dist" ) }, plugins :[ new VueSSRClientPlugin(), new HtmlWebpackPlugin({ template :path.resolve(__dirname, "../index.html" ), filename : "index.client.html" }) ] }); module.exports=config; 复制代码`

**webpack.server.conf.js**

` const path= require ( "path" ); const merge = require ( 'webpack-merge' ); const HtmlWebpackPlugin = require ( 'html-webpack-plugin' ); const base = require ( './webpack.base.config' ); const nodeExternals = require ( 'webpack-node-externals' ); const VueSSRServerPlugin = require ( 'vue-server-renderer/server-plugin' ); module.exports=merge(base,{ target : "node" , output : { path :path.resolve(__dirname, "../dist" ), libraryTarget : 'commonjs2' }, devtool : '#source-map' , externals :[nodeExternals()], //排除node_modules entry:path.resolve(__dirname, "../src/entry-server.js" ), plugins :[ new VueSSRServerPlugin(), new HtmlWebpackPlugin({ template :path.resolve(__dirname, "../index.ssr.html" ), filename : "index.ssr.html" , excludeChunks : [ 'main' , 'client' ] }) ] }); 复制代码`

**3、node后台服务搭建** ：这里主要利用koa框架搭建node后台服务层，实现过程采用了koa、koa-router和koa-static依赖包；文件命名为server.js

` const fs = require ( 'fs' ); const path = require ( 'path' ); const Koa = require ( 'koa' ); const Router = require ( 'koa-router' ); const serve = require ( 'koa-static' ); const appOne = new Koa(); const appTwo = new Koa(); const routerOne = new Router(); const routerTwo = new Router(); // 后端Server routerOne.get( '/index' , (ctx, next) => { ctx.type = 'html' ; ctx.status = 200 ; //先这样展示，后续需要服务器端渲染代码 ctx.body = '<h1>服务器端渲染机制</h1>' ; }); appOne.use(serve(path.resolve(__dirname, '../dist' ))); appOne.use(routerOne.routes()) .use(routerOne.allowedMethods()); //处理跨域问题 appOne.listen( 3001 , () => { console.log( '服务器端渲染地址： http://localhost:3001' ); }); // 前端Server routerTwo.get( '/index' , (ctx, next) => { let html = fs.readFileSync(path.resolve(__dirname, '../dist/index.client.html' ), 'utf-8' ); ctx.type = 'html' ; ctx.status = 200 ; ctx.body = html; }); appTwo.use(serve(path.resolve(__dirname, '../dist' ))); appTwo.use(routerTwo.routes()) .use(routerTwo.allowedMethods()); //处理跨域问题 appTwo.listen( 3002 , () => { console.log( '浏览器端渲染地址： http://localhost:3002' ); }); 复制代码`

**4、使用vue-server-renderer包** ：上述提到过该包一般在node后台服务端使用，也就是继承到server.js代码中。它主要解决将Vue组件渲染成html字符串；需要注意点，将Vue渲染成html字符串时间不确定，所以需要使用async/await关键字等待渲染完成，才能返回响应。不然这里还没渲染完，响应就开始了；

` const fs = require ( 'fs' ); const path = require ( 'path' ); const Koa = require ( 'koa' ); const Router = require ( 'koa-router' ); const serve = require ( 'koa-static' ); const appOne = new Koa(); const appTwo = new Koa(); const routerOne = new Router(); const routerTwo = new Router(); //不同点AAAAAAAAA——start const serverBundle = require (path.resolve(__dirname, '../dist/vue-ssr-server-bundle.json' )); const clientManifest = require (path.resolve(__dirname, '../dist/vue-ssr-client-manifest.json' )); const template = fs.readFileSync(path.resolve(__dirname, '../dist/index.ssr.html' ), 'utf-8' ); const renderer = require ( 'vue-server-renderer' ).createBundleRenderer(serverBundle, { runInNewContext : false , template : template, clientManifest : clientManifest }); //不同点AAAAAAAAA——end // 后端Server //不同点BBBBBBBBBBB——start routerOne.get( '/index' , async (ctx, next) => { //vue-server-renderer标记点—— try { //由于渲染时间不确定，所以需要async/await关键字等待渲染完成 let html = await new Promise ( ( resolve, reject ) => { renderer.renderToString( ( err, html ) => { if (err) { reject(err); } else { resolve(html); } }); }); ctx.type = 'html' ; ctx.status = 200 ; ctx.body = html; } catch (err) { console.log(err); ctx.status = 500 ; ctx.body = '服务器内部错误' ; } //不同点BBBBBBBBBBB——end }); appOne.use(serve(path.resolve(__dirname, '../dist' ))); appOne.use(routerOne.routes()) .use(routerOne.allowedMethods()); //处理跨域问题 appOne.listen( 3001 , () => { console.log( '服务器端渲染地址： http://localhost:3001' ); }); // 前端Server routerTwo.get( '/index' , (ctx, next) => { let html = fs.readFileSync(path.resolve(__dirname, '../dist/index.client.html' ), 'utf-8' ); ctx.type = 'html' ; ctx.status = 200 ; ctx.body = html; }); appTwo.use(serve(path.resolve(__dirname, '../dist' ))); appTwo.use(routerTwo.routes()) .use(routerTwo.allowedMethods()); //处理跨域问题 appTwo.listen( 3002 , () => { console.log( '浏览器端渲染地址： http://localhost:3002' ); }); 复制代码`

**5、业务代码设计** ：在Vue-cli创建的项目中，main.js文件为主入口文件。对于同构应用，需要将实例Vue封装到函数中（app.js）给不同平台设置不同主入口文件（entry-client.js和entry-server.js）。记住它们的区别：vue的生命周期不一样和初始化Vue实例不一样；

**普通的vue实例化（main.js）**

` //main.js文件 import Vue from "vue" ; import App from "./App.vue" ; import {initRouter} from "./router" ; import {initStore} from "./store" ; new Vue({ el : "#app" , router, store, render : h => h(App) }); 复制代码`

**同构应用Vue实例化（app.js、entry-client.js、entry-server.js）**

` //app.js文件 import Vue from "vue" ; import App from "./App.vue" ; import {initRouter} from "./router" ; import {initStore} from "./store" ; export function initVue ( ) { const {router}=initRouter(); const {store}=initStore(); const app= new Vue({ router, store, render : h => h(App) }); return {app,store,router,App}; } 复制代码` ` //entry-client.js文件 import {initVue} from "./app" ; const {app,store,router}=initVue(); router.onReady( () => { app.$mount( "#app" ); }); 复制代码` ` //entry-server.js文件 import {initVue} from "./app" ; export default context => { const { app } = initVue(); return new Promise ( ( resolve, reject ) => { const { app, store, router, App } = initVue(); let components = App.components; //判断组件是否有asyncData方法，执行asyncData方法 Object.values(components).forEach( ( component ) => { if (component.asyncData) { component.asyncData({ store }); } }); resolve(app); }); } 复制代码`
> 
> 
> 
> 
> 实例项目地址： [github.com/song199210/…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsong199210%2FDemoPro%2Ftree%2Fmaster%2FVueSRR
> )
> 
>