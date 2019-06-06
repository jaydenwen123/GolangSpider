# 面试官：自己搭建过vue开发环境吗？ #

## 开篇 ##

> 
> 
> 
> 原文地址： [www.ccode.live/lentoo/list…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.ccode.live%2Flentoo%2Flist%2F33
> )
> 
> 

前段时间，看到群里一些小伙伴面试的时候被面试官问到这类题目。平时大家开发vue项目的时候，相信大部分人都是使用 ` vue-cli` 脚手架生成的项目架构，然后 ` npm run install` 安装依赖， ` npm run serve` 启动项目然后就开始写业务代码了。

但是对项目里的 ` webpack` 封装和配置了解的不清楚，容易导致出问题不知如何解决，或者不会通过 ` webpack` 去扩展新功能。

该篇文章主要是想告诉小伙伴们，如何一步一步的通过 ` webpack4` 来搭建自己的 ` vue` 开发环境

首先我们要知道 ` vue-cli` 生成的项目，帮我们配置好了哪些功能？

* ` ES6` 代码转换成 ` ES5` 代码
* ` scss/sass/less/stylus` 转 ` css`
* `.vue` 文件转换成 ` js` 文件
* 使用 ` jpg` 、 ` png` ， ` font` 等资源文件
* 自动添加css各浏览器产商的前缀
* 代码热更新
* 资源预加载
* 每次构建代码清除之前生成的代码
* 定义环境变量
* 区分开发环境打包跟生产环境打包
* ....

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66cc9a0ca6c9f?imageView2/0/w/1280/h/960/ignore-error/1)

## 1. 搭建 ` webpack` 基本环境 ##

该篇文章并不会细讲 ` webpack` 是什么东西，如果还不是很清楚的话，可以先去看看 [webpack官网]( https://link.juejin.im?target=https%3A%2F%2Fwebpack.github.io%2F )

简单的说， ` webpack` 是一个模块打包机，可以分析你的项目依赖的模块以及一些浏览器不能直接运行的语言 ` jsx` 、 ` vue` 等转换成 ` js` 、 ` css` 文件等，供浏览器使用。

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66d345acde608?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.1 初始化项目 ###

在命令行中执行 ` npm init` 然后一路回车就行了，主要是生成一些项目基本信息。最后会生成一个 ` package.json` 文件

` npm init 复制代码`

### 1.2 安装 ` webpack` ###

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66e40ecf19e95?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.3 写点小代码测试一下 ` webpack` 是否安装成功了 ###

新建一个 ` src` 文件夹，然后再建一个 ` main.js` 文件

` // src/main.js console.log( 'hello webpack' ) 复制代码`

然后在 package.json 下面加一个脚本命令

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66f9c4569e58c?imageView2/0/w/1280/h/960/ignore-error/1)

然后运行该命令

` npm run serve 复制代码`

如果在 dist 目录下生成了一个 ` main.js` 文件，则表示 ` webpack` 工作正常

## 2. 开始配置功能 ##

* 新建一个 ` build` 文件夹，用来存放 ` webpack` 配置相关的文件
* 在 ` build` 文件夹下新建一个 ` webpack.config.js` ，配置 ` webpack` 的基本配置
* 修改 ` webpack.config.js` 配置

![](https://user-gold-cdn.xitu.io/2019/4/29/16a67c1cb8f525d0?imageView2/0/w/1280/h/960/ignore-error/1)

* 修改 ` package.json` 文件，将之前添加的 ` serve` 修改为

` "serve" : "webpack ./src/main.js --config ./build/webpack.config.js" 复制代码`

### 2.1 配置 ` ES6/7/8` 转 ` ES5` 代码 ###

* 安装相关依赖

` npm install babel-loader @babel/core @babel/preset-env 复制代码`

* 修改 ` webpack.config.js` 配置

![](https://user-gold-cdn.xitu.io/2019/5/7/16a8fe90c3c9384b?imageView2/0/w/1280/h/960/ignore-error/1)

* 在项目根目录添加一个 ` babel.config.js` 文件

![](https://user-gold-cdn.xitu.io/2019/4/29/16a672d7527a9097?imageView2/0/w/1280/h/960/ignore-error/1)

* 然后执行 ` npm run serve` 命令，可以看到 ES6代码被转成了ES5代码了

#### 2.1.1 ` ES6/7/8 Api` 转 ` es5` ####

` babel-loader` 只会将 ES6/7/8语法转换为ES5语法，但是对新api并不会转换。

我们可以通过 babel-polyfill 对一些不支持新语法的客户端提供新语法的实现

* 安装

` npm install @babel/polyfill 复制代码`

* 修改 ` webpack.config.js` 配置

在 ` entry` 中添加 ` @babel-polyfill`

![](https://user-gold-cdn.xitu.io/2019/4/30/16a6bcbe53ecb912?imageView2/0/w/1280/h/960/ignore-error/1)

#### 2.1.2 按需引入 ` polyfill` ####

2.1.2 和 2.1.1 只需要配置一个就行

> 
> 
> 
> 修改时间 2019-05-05、 来自评论区 **兮漫天** 的提醒
> 
> 

* 安装相关依赖

` npm install core-js@2 @babel/runtime-corejs2 -S 复制代码`

* 修改 babel-config.js

![](https://user-gold-cdn.xitu.io/2019/5/5/16a8593e4b0e0179?imageView2/0/w/1280/h/960/ignore-error/1)

配置了按需引入 ` polyfill` 后，用到 ` es6` 以上的函数， ` babel` 会自动导入相关的 ` polyfill` ，这样能大大减少 打包编译后的体积

### 2.2 配置 ` scss` 转 ` css` ###

在没配置 ` css` 相关的 ` loader` 时，引入 ` scss` 、 ` css` 相关文件打包的话，会报错

* 安装相关依赖

` npm install sass-loader dart-sass css-loader style-loader -D 复制代码`

` sass-loader` , ` dart-sass` 主要是将 scss/sass 语法转为css

` css-loader` 主要是解析 css 文件

` style-loader` 主要是将 css 解析到 ` html` 页面 的 ` style` 上

* 修改 ` webpack.config.js` 配置

![](https://user-gold-cdn.xitu.io/2019/4/29/16a67a226b004f44?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.3 配置 postcss 实现自动添加css3前缀 ###

* 安装相关依赖

` npm install postcss-loader autoprefixer -D 复制代码`

* 修改 ` webpack.config.js` 配置

![](https://user-gold-cdn.xitu.io/2019/4/29/16a67ec74576da4b?imageView2/0/w/1280/h/960/ignore-error/1)

* 在项目根目录下新建一个 ` postcss.config.js`

![](https://user-gold-cdn.xitu.io/2019/4/29/16a68300d9a47ae1?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.3 使用 ` html-webpack-plugin` 来创建html页面 ###

使用 ` html-webpack-plugin` 来创建html页面，并自动引入打包生成的 ` js` 文件

* 安装依赖

` npm install html-webpack-plugin -D 复制代码`

* 新建一个 public/index.html 页面

` <!DOCTYPE html> < html lang = "en" > < head > < meta charset = "UTF-8" > < meta name = "viewport" content = "width=device-width, initial-scale=1.0" > < meta http-equiv = "X-UA-Compatible" content = "ie=edge" > < title > Document </ title > </ head > < body > < div id = "app" > </ div > </ body > </ html > 复制代码`

* 修改 ` webpack-config.js` 配置 ![](https://user-gold-cdn.xitu.io/2019/4/29/16a67ae8b4fde570?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.4 配置 devServer 热更新功能 ###

通过代码的热更新功能，我们可以实现不刷新页面的情况下，更新我们的页面

* 安装依赖

` npm install webpack-dev-server -D 复制代码`

* 修改 ` webpack.config.js` 配置

通过配置 ` devServer` 和 ` HotModuleReplacementPlugin` 插件来实现热更新

![](https://user-gold-cdn.xitu.io/2019/4/29/16a67ce41df0920b?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.5 配置 webpack 打包 图片、媒体、字体等文件 ###

* 安装依赖

` npm install file-loader url-loader -D 复制代码`

` file-loader` 解析文件url，并将文件复制到输出的目录中

` url-loader` 功能与 ` file-loader` 类似，如果文件小于限制的大小。则会返回 ` base64` 编码，否则使用 ` file-loader` 将文件复制到输出的目录中

* 修改 ` webpack-config.js` 配置 添加 ` rules` 配置，分别对 图片，媒体，字体文件进行配置

` // build/webpack.config.js const path = require ( 'path' ) const HtmlWebpackPlugin = require ( 'html-webpack-plugin' ) const webpack = require ( 'webpack' ) module.exports = { // 省略其它配置 ... module : { rules : [ // ... { test : /\.(jpe?g|png|gif)$/i , use : [ { loader : 'url-loader' , options : { limit : 4096 , fallback : { loader : 'file-loader' , options : { name : 'img/[name].[hash:8].[ext]' } } } } ] }, { test : /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/ , use : [ { loader : 'url-loader' , options : { limit : 4096 , fallback : { loader : 'file-loader' , options : { name : 'media/[name].[hash:8].[ext]' } } } } ] }, { test : /\.(woff2?|eot|ttf|otf)(\?.*)?$/i , use : [ { loader : 'url-loader' , options : { limit : 4096 , fallback : { loader : 'file-loader' , options : { name : 'fonts/[name].[hash:8].[ext]' } } } } ] }, ] }, plugins : [ // ... ] } 复制代码`

## 3. 让 ` webpack` 识别 `.vue` 文件 ##

* 安装需要的依赖文件

` npm install vue-loader vue-template-compiler cache-loader thread-loader -D npm install vue -S 复制代码`

` vue-loader` 用于解析 `.vue` 文件

` vue-template-compiler` 用于编译模板

` cache-loader` 用于缓存 ` loader` 编译的结果

` thread-loader` 使用 ` worker` 池来运行 ` loader` ，每个 ` worker` 都是一个 ` node.js` 进程。

* 修改 ` webpack.config.js` 配置

` // build/webpack.config.js const path = require ( 'path' ) const webpack = require ( 'webpack' ) const HtmlWebpackPlugin = require ( 'html-webpack-plugin' ) const VueLoaderPlugin = require ( 'vue-loader/lib/plugin' ) module.exports = { // 指定打包模式 mode: 'development' , entry : { // ... }, output : { // ... }, devServer : { // ... }, resolve : { alias : { vue$ : 'vue/dist/vue.runtime.esm.js' }, }, module : { rules : [ { test : /\.vue$/ , use : [ { loader : 'cache-loader' }, { loader : 'thread-loader' }, { loader : 'vue-loader' , options : { compilerOptions : { preserveWhitespace : false }, } } ] }, { test : /\.jsx?$/ , use : [ { loader : 'cache-loader' }, { loader : 'thread-loader' }, { loader : 'babel-loader' } ] }, // ... ] }, plugins : [ // ... new VueLoaderPlugin() ] } 复制代码`

* 测试一下

* 在 src 新建一个 App.vue
` // src/App.vue <template> <div class= "App" > Hello World </div> </template> <script> export default { name: 'App' , data () { return {}; } }; </script> <style lang= "scss" scoped> .App { color: skyblue; } </style> 复制代码` * 修改 ` main.js`
` import Vue from 'vue' import App from './App.vue' new Vue({ render: h => h(App) }). $mount ( '#app' ) 复制代码` * 运行一下

` npm run serve`

## 4. 定义环境变量 ##

通过 ` webpack` 提供的 ` DefinePlugin` 插件，可以很方便的定义环境变量

` plugins: [ new webpack.DefinePlugin({ 'process.env' : { VUE_APP_BASE_URL : JSON.stringify( 'http://localhost:3000' ) } }), ] 复制代码`

## 5. 区分生产环境和开发环境 ##

新建两个文件

* 

` webpack.dev.js` 开发环境使用

* 

` webpack.prod.js` 生产环境使用

* 

` webpack.config.js` 公用配置

* 

开发环境与生产环境的不同

### 5.1 开发环境 ###

* 不需要压缩代码
* 需要热更新
* css不需要提取到css文件
* sourceMap
* ...

### 5.2 生产环境 ###

* 压缩代码
* 不需要热更新
* 提取css，压缩css文件
* sourceMap
* 构建前清除上一次构建的内容
* ...

* 安装所需依赖

` npm i @intervolga/optimize-cssnano-plugin mini-css-extract-plugin clean-webpack-plugin webpack-merge copy-webpack-plugin -D 复制代码` * ` @intervolga/optimize-cssnano-plugin` 用于压缩css代码
* ` mini-css-extract-plugin` 用于提取css到文件中
* ` clean-webpack-plugin` 用于删除上次构建的文件
* ` webpack-merge` 合并 ` webpack` 配置
* ` copy-webpack-plugin` 用户拷贝静态资源

### 5.3 开发环境配置 ###

* build/webpack.dev.js

` // build/webpack.dev.js const merge = require ( 'webpack-merge' ) const webpackConfig = require ( './webpack.config' ) const webpack = require ( 'webpack' ) module.exports = merge(webpackConfig, { mode : 'development' , devtool : 'cheap-module-eval-source-map' , module : { rules : [ { test : /\.(scss|sass)$/ , use : [ { loader : 'style-loader' }, { loader : 'css-loader' , options : { importLoaders : 2 } }, { loader : 'sass-loader' , options : { implementation : require ( 'dart-sass' ) } }, { loader : 'postcss-loader' } ] }, ] }, plugins : [ new webpack.DefinePlugin({ 'process.env' : { NODE_ENV : JSON.stringify( 'development' ) } }), ] }) 复制代码`

* webpack.config.js

` // build/webpack.config.js const path = require ( 'path' ) const webpack = require ( 'webpack' ) const HtmlWebpackPlugin = require ( 'html-webpack-plugin' ) const VueLoaderPlugin = require ( 'vue-loader/lib/plugin' ) module.exports = { entry : { // 配置入口文件 main: path.resolve(__dirname, '../src/main.js' ) }, output : { // 配置打包文件输出的目录 path: path.resolve(__dirname, '../dist' ), // 生成的 js 文件名称 filename: 'js/[name].[hash:8].js' , // 生成的 chunk 名称 chunkFilename: 'js/[name].[hash:8].js' , // 资源引用的路径 publicPath: '/' }, devServer : { hot : true , port : 3000 , contentBase : './dist' }, resolve : { alias : { vue$ : 'vue/dist/vue.runtime.esm.js' }, extensions : [ '.js' , '.vue' ] }, module : { rules : [ { test : /\.vue$/ , use : [ { loader : 'cache-loader' }, { loader : 'vue-loader' , options : { compilerOptions : { preserveWhitespace : false }, } } ] }, { test : /\.jsx?$/ , loader : 'babel-loader' }, { test : /\.(jpe?g|png|gif)$/ , use : [ { loader : 'url-loader' , options : { limit : 4096 , fallback : { loader : 'file-loader' , options : { name : 'img/[name].[hash:8].[ext]' } } } } ] }, { test : /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/ , use : [ { loader : 'url-loader' , options : { limit : 4096 , fallback : { loader : 'file-loader' , options : { name : 'media/[name].[hash:8].[ext]' } } } } ] }, { test : /\.(woff2?|eot|ttf|otf)(\?.*)?$/i , use : [ { loader : 'url-loader' , options : { limit : 4096 , fallback : { loader : 'file-loader' , options : { name : 'fonts/[name].[hash:8].[ext]' } } } } ] }, ] }, plugins : [ new VueLoaderPlugin(), new HtmlWebpackPlugin({ template : path.resolve(__dirname, '../public/index.html' ) }), new webpack.NamedModulesPlugin(), new webpack.HotModuleReplacementPlugin(), ] } 复制代码`

### 5.4 生产环境配置 ###

` const path = require ( 'path' ) const merge = require ( 'webpack-merge' ) const webpack = require ( 'webpack' ) const webpackConfig = require ( './webpack.config' ) const MiniCssExtractPlugin = require ( 'mini-css-extract-plugin' ) const OptimizeCssnanoPlugin = require ( '@intervolga/optimize-cssnano-plugin' ); const CleanWebpackPlugin = require ( 'clean-webpack-plugin' ) const CopyWebpackPlugin = require ( 'copy-webpack-plugin' ) module.exports = merge(webpackConfig, { mode : 'production' , devtool : '#source-map' , optimization : { splitChunks : { cacheGroups : { vendors : { name : 'chunk-vendors' , test : /[\\\/]node_modules[\\\/]/ , priority : -10 , chunks : 'initial' }, common : { name : 'chunk-common' , minChunks : 2 , priority : -20 , chunks : 'initial' , reuseExistingChunk : true } } } }, module : { rules : [ { test : /\.(scss|sass)$/ , use : [ { loader : MiniCssExtractPlugin.loader }, { loader : 'css-loader' , options : { importLoaders : 2 } }, { loader : 'sass-loader' , options : { implementation : require ( 'dart-sass' ) } }, { loader : 'postcss-loader' } ] }, ] }, plugins : [ new webpack.DefinePlugin({ 'process.env' : { NODE_ENV : 'production' } }), new MiniCssExtractPlugin({ filename : 'css/[name].[contenthash:8].css' , chunkFilename : 'css/[name].[contenthash:8].css' }), new OptimizeCssnanoPlugin({ sourceMap : true , cssnanoOptions : { preset : [ 'default' , { mergeLonghand : false , cssDeclarationSorter : false } ] } }), new CopyWebpackPlugin([ { from : path.resolve(__dirname, '../public' ), to : path.resolve(__dirname, '../dist' ) } ]), new CleanWebpackPlugin() ] }) 复制代码`

### 5.5 修改package.json ###

` "scripts" : { "serve" : "webpack-dev-server --config ./build/webpack.dev.js" , "build" : "webpack --config ./build/webpack.prod.js" }, 复制代码`

## 6 打包分析 ##

有的时候，我们需要看一下webpack打包完成后，到底打包了什么东西，

这时候就需要用到这个模块分析工具了 ` webpack-bundle-analyzer`

* 安装依赖

` npm install --save-dev webpack-bundle-analyzer 复制代码`

* 修改 ` webpack-prod.js` 配置，在 ` plugins` 属性中新增一个插件

在开发环境中，我们是没必要进行模块打包分析的，所以我们将插件配置在了生产环境的配置项中

![](https://user-gold-cdn.xitu.io/2019/5/15/16ab98888065fcbb?imageView2/0/w/1280/h/960/ignore-error/1)

* 运行打包命令

` npm run build 复制代码`

执行成功后会自动打开这个页面

![](https://user-gold-cdn.xitu.io/2019/5/15/16ab9897ac9dfea3?imageView2/0/w/1280/h/960/ignore-error/1)

## 7. 集成 ` VueRouter` , ` Vuex` ##

* 首先是安装相关依赖
` npm install vue-router vuex --save 复制代码`

### 7.1 集成 ` Vue-Router` ###

* 新增视图组件 在 ` src` 目录下新增两个视图组件 ` src/views/Home.vue` 和 ` src/views/About.vue`

` // src/views/Home.vue < template > < div class = "Home" > < h2 > Home </ h2 > </ div > </ template > < script > export default { name : 'Home' , data() { return {}; } }; </ script > < style lang = "scss" scoped > </ style > 复制代码`

` About.vue` 内容跟 ` Home.vue` 差不多，将里面的 ` Home` 换成 ` About` 就OK了

* 新增路由配置文件

在 ` src` 目录下新增一个 ` router/index.js` 文件

` // src/router/index.js import Vue from 'vue' import VueRouter from "vue-router" ; import Home from '../views/Home' ; import About from '../views/About' ; Vue.use(VueRouter) export default new VueRouter({ mode : 'hash' , routes : [ { path : '/Home' , component : Home }, { path : '/About' , component : About }, { path : '*' , redirect : '/Home' } ] }) 复制代码`

* 修改 ` main.js` 文件

` // main.js import Vue from 'vue' import App from './App.vue' import router from './router' new Vue({ router, render : h => h(App) }).$mount( '#app' ) 复制代码`

* 修改 ` App.vue` 组件

` // App.vue // 在 template 中添加 // src/App.vue <template> <div class= "App" > Hello World </div> <div> // router-link 组件 用来导航到哪个路由 <router-link to= "/Home" >go Home</router-link> <router-link to= "/About" >go About</router-link> </div> <div> // 用于展示匹配到的路由视图组件 <router-view></router-view> </div> </template> <script> export default { name: 'App' , data () { return {}; } }; </script> <style lang= "scss" scoped> .App { color: skyblue; } </style> 复制代码`

运行 ` npm run serve` 命令，如没配置错误，是可以看到点击不同的路由，会切换到不同的路由视图

### 7.2 配置路由懒加载 ###

在没配置路由懒加载的情况下，我们的路由组件在打包的时候，都会打包到同一个 ` js` 文件去，当我们的视图组件越来越多的时候，就会导致这个 ` js` 文件越来越大。然后就会导致请求这个文件的时间变长，最终影响用户体验

* 安装依赖
` npm install @babel/plugin-syntax-dynamic-import --save-dev 复制代码` * 修改 ` babel.config.js`
` module.exports = { presets : [ [ "@babel/preset-env" , { useBuiltIns : "usage" } ] ], plugins : [ // 添加这个 '@babel/plugin-syntax-dynamic-import' ] } 复制代码` * 修改 ` router/index.js` 路由配置文件
` import Vue from 'vue' import VueRouter from "vue-router" ; Vue.use(VueRouter) export default new VueRouter({ mode : 'hash' , routes : [ { path : '/Home' , component : () => import ( /* webpackChunkName: "Home" */ '../views/Home.vue' ) // component: Home }, { path : '/About' , component : () => import ( /* webpackChunkName: "About" */ '../views/About.vue' ) // component: About }, { path : '*' , redirect : '/Home' } ] }) 复制代码` * 运行命令 ` npm run build` 查看是否生成了 ` Home...js` 文件 和 ` About...js` 文件

### 7.3 集成 Vuex ###

* 在 ` src` 目录下新建一个 ` store/index.js` 文件
` // store/index.js import Vue from 'vue' import Vuex from 'vuex' Vue.use(Vuex) const state = { counter : 0 } const actions = { add : ( {commit} ) => { return commit( 'add' ) } } const mutations = { add : ( state ) => { state.counter++ } } const getters = { getCounter (state) { return state.counter } } export default new Vuex.Store({ state, actions, mutations, getters }) 复制代码` * 修改 ` main.js` 文件 导入 ` vuex`
` // main.js import Vue from 'vue' import App from './App.vue' import router from './router' import store from './store' // ++ new Vue({ router, store, // ++ render: h => h(App) }).$mount( '#app' ) 复制代码` * 修改 ` App.vue` ，查看 vuex 配置效果
` // App.vue < template > < div class = "App" > < div > < router-link to = "/Home" > go Home </ router-link > < router-link to = "/About" > go About </ router-link > </ div > < div > < p > {{getCounter}} </ p > < button @ click = "add" > add </ button > </ div > < div > < router-view > </ router-view > </ div > </ div > </ template > < script > import { mapActions, mapGetters } from 'vuex' export default { name : 'App' , data() { return {}; }, computed : { ...mapGetters([ 'getCounter' ]) }, methods : { ...mapActions([ 'add' ]) } }; </ script > < style lang = "scss" scoped >.App { text-align : center; color : skyblue; font-size : 28px ; } </ style > 复制代码` * 运行命令 ` npm run serve`

当点击按钮的时候，可以看到我们的 ` getCounter` 一直在增加

## 8 总结 ##

到目前为止，我们已经成功的自己搭建了一个 ` vue` 开发环境，不过还是有一些功能欠缺的，有兴趣的小伙伴可以交流交流。在搭建过程中，还是会踩很多坑的。

如果还不熟悉 webpack 的话，建议自己搭建一次。可以让自己能深入的理解 ` vue-cli` 替我们做了什么

![](https://user-gold-cdn.xitu.io/2019/4/29/16a687ab79eaeedd?imageView2/0/w/1280/h/960/ignore-error/1)

## 推荐阅读 ##

* 

[使用 webpack 的各种插件提升你的开发效率]( https://juejin.im/post/5c8852f95188257a323f5cee )

* 

[vue-cli3 项目从搭建优化到docker部署]( https://juejin.im/post/5c4a6fcd518825469414e062 )

* 

[Event Loop 原来是这么回事]( https://juejin.im/post/5c36b3b0f265da611f07e409 )

* 

[通过vue-cli3构建一个SSR应用程序]( https://juejin.im/post/5b98e5875188255c8320f88a )

## 欢迎关注 ##

欢迎关注公众号“ **码上开发** ”，每天分享最新技术资讯

![image](https://user-gold-cdn.xitu.io/2018/12/24/167ddc2c7f13cdf5?imageView2/0/w/1280/h/960/ignore-error/1)