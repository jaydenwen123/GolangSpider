# 从今天开始，学习Webpack，减少对脚手架的依赖(下) #

问：这篇文章适合哪些人？
答：适合没接触过Webpack或者了解不全面的人。

问：这篇文章的目录怎么安排的？
答：先介绍背景，由背景引入Webpack的概念，进一步介绍Webpack基础、核心和一些常用配置案例、优化手段，Webpack的plugin和loader确实非常多，短短2w多字还只是覆盖其中一小部分。

问：这篇文章的出处？
答：此篇文章知识来自付费视频(链接在文章末尾)，文章由自己独立撰写，已获得讲师授权并首发于掘金

上一篇： [从今天开始，学习Webpack，减少对脚手架的依赖(上)]( https://juejin.im/post/5cecd6fdf265da1b827a7ca5 )

如果你觉得写的不错，请给我点一个star，原博客地址： [原文地址]( https://link.juejin.im?target=https%3A%2F%2Fwangtunan.github.io%2Fblog%2Fwebpack%2F )

# PWA配置 #

PWA全称 ` Progressive Web Application` (渐进式应用框架)，它能让我们主动缓存文件，这样用户离线后依然能够使用我们缓存的文件打开网页，而不至于让页面挂掉，实现这种技术需要安装 ` workbox-webpack-plugin` 插件。

> 
> 
> 
> 如果你的谷歌浏览器还没有开启支持PWA，请开启它再进行下面的测试。
> 
> 

## 安装插件 ##

` $ npm install workbox-webpack-plugin -D 复制代码`

## webpack.config.js文件配置 ##

` // PWA只有在线上环境才有效，所以需要在webpack.prod.js文件中进行配置 const WorkboxWebpackPlugin = require('workbox-webpack-plugin'); const prodConfig = { // 其它配置 plugins: [ new MiniCssExtractPlugin({}), new WorkboxWebpackPlugin.GenerateSW({ clientsClaim: true, skipWaiting: true }) ] } module.exports = merge(commonConfig, prodConfig); 复制代码`

以上配置完毕后，让我们使用 ` npm run build` 打包看一看生成了哪些文件， ` dist` 目录的打包结果如下：

` |-- dist | |-- index.html | |-- main.f28cbac9bec3756acdbe.js | |-- main.f28cbac9bec3756acdbe.js.map | |-- precache-manifest.ea54096f38009609a46058419fc7009b.js | |-- service-worker.js 复制代码`

我们可以代码块高亮的部分，多出来了 ` precache-manifest.xxxxx.js` 文件和 ` service-worker.js` ，就是这两个文件能让我们实现PWA。

## 改写index.js ##

需要判断浏览器是否支持PWA，支持的时候我们才进行注册，注册的 `.js` 文件为我们打包后的 ` service-worker.js` 文件。

` console.log( 'hello,world' ); if ( 'serviceWorker' in navigator) { navigator.serviceWorker.register( '/service-worker.js' ).then( ( register ) => { console.log( '注册成功' ); }).catch( error => { console.log( '注册失败' ); }) } 复制代码`

### PWA实际效果 ###

在 ` npm run dev` 后，我们利用 ` webpack-dev-server` 启动了一个小型的服务器，然后我们停掉这个服务器，刷新页面，PWA的实际结果如下图所示

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd50c28056c0f?imageView2/0/w/1280/h/960/ignore-error/1)

### WebpackDevServer请求转发 ###

在这一小节中，我们要学到的技能有：

* 如何进行接口代理配置
* 如何使用接口路径重写
* 其他常见配置的介绍

假设我们现在有这样一个需求：我有一个URL地址( ` http://www.dell-lee.com/react/api/header.json` )，我希望我请求的时候，请求的地址是 ` /react/api/header.json` ，能有一个什么东西能自动帮我把请求转发到 ` http://www.dell-lee.com` 域名下，那么这个问题该如何解决呢？可以使用 Webpack 的 ` webpack-dev-server` 这个插件来解决，其中需要配置 ` proxy` 属性。

#### 如何进行接口代理配置 ####

既然我们要做请求，那么安装 ` axios` 来发请求再合适不过了，使用如下命令安装 ` axios` :

` $ npm install axios --save-dev 复制代码`

因为我们的请求代理只能在开发环境下使用，线上的生产环境，需要走其他的代理配置，所以我们需要在 ` webpack.dev.js` 中进行代理配置

` const devConfig = { // 其它配置 devServer: { contentBase : './dist' , open : false , port : 3000 , hot : true , hotOnly : true , proxy : { '/react/api' : { target : 'http://www.dell-lee.com' } } } } 复制代码`

以上配置完毕后，我们在 ` index.js` 文件中引入 ` axios` 模块，再做请求转发。

` import axios from 'axios' ; axios.get( '/react/api/header.json' ).then( ( res ) => { let {data,status} = res; console.log(data); }) 复制代码`

使用 ` npm run dev` 后， 我们可以在浏览器中看到，我们已经成功请求到了我们的数据。

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd51a81b75efd?imageView2/0/w/1280/h/960/ignore-error/1)

#### 如何使用接口路径重写 ####

现在依然假设有这样一个场景： ` http://www.dell-lee.com/react/api/header.json` 这个后端接口还没有开发完毕，但后端告诉我们可以先使用 ` http://www.dell-lee.com/react/api/demo.json` 这个测试接口，等接口开发完毕后，我们再改回来。解决这个问题最佳办法是，代码中的地址不能变动，我们只在 ` proxy` 代理中处理即可，使用 ` pathRewrite` 属性进行配置。

` const devConfig = { // 其它配置 devServer: { contentBase: './dist', open: false, port: 3000, hot: true, hotOnly: true, proxy: { '/react/api': { target: 'http://www.dell-lee.com', pathRewrite: { 'header.json': 'demo.json' } } } } } 复制代码`

同样，我们打包后在浏览器中可以看到，我们的测试接口的数据已经成功拿到了。

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd51ec8ddfdf0?imageView2/0/w/1280/h/960/ignore-error/1)

#### 其他常见配置的含义 ####

**转发到https：** 一般情况下，不接受运行在 ` https` 上，如果要转发到 ` https` 上，可以使用如下配置

` module.exports = { //其它配置 devServer: { proxy: { '/react/api': { target: 'https://www.dell-lee.com', secure: false } } } } 复制代码`

**跨域：** 有时候，在请求的过程中，由于同源策略的影响，存在跨域问题，我们需要处理这种情况，可以如下进行配置。

` module.exports = { //其它配置 devServer: { proxy: { '/react/api': { target: 'https://www.dell-lee.com', changeOrigin: true, } } } } 复制代码`

**代理多个路径到同一个target：** 代理多个路径到同一个 ` target` ，可以如下进行配置

` module.exports = { //其它配置 devServer: { proxy: [{ context: ['/vue/api', '/react/api'], target: 'http://www.dell-lee.com' }] } } 复制代码`

# 多页打包 #

现在流行的前端框架都推行单页引用(SPA)，但有时候我们不得不兼容一些老的项目，他们是多页的，那么如何进行多页打包配置呢？ 现在我们来思考一个问题：多页运用，即 **多个入口文件+多个对应的html文件** ，那么我们就可以配置 **多个入口+配置多个 ` html-webpack-plugin`** 来进行。

场景：假设现在我们有这样三个页面： ` index.html` , ` list.html` , ` detail.html` ，我们需要配置三个入口文件，新建三个 `.js` 文件。

在 ` webpack.common.js` 中配置多个 ` entry` 并使用 ` html-webpack-plugin` 来生成对应的多个 `.html` 页面。 **HtmlWebpackPlugin参数说明** ：

* ` template` ：代表以哪个HTML页面为模板
* ` filename` ：代表生成页面的文件名
* ` chunks` ：代表需要引用打包后的哪些 `.js` 文件

` module.exports = { // 其它配置 entry: { index : './src/index.js' , list : './src/list.js' , detail : './src/detail.js' , }, plugins : [ new htmlWebpackPlugin({ template : 'src/index.html' , filename : 'index.html' , chunks : [ 'index' ] }), new htmlWebpackPlugin({ template : 'src/index.html' , filename : 'list.html' , chunks : [ 'list' ] }), new htmlWebpackPlugin({ template : 'src/index.html' , filename : 'detail.html' , chunks : [ 'detail' ] }), new cleanWebpackPlugin() ] } 复制代码`

在 ` src` 目录下新建三个 `.js` 文件，名字分别是： ` index.js` ， ` list.js` 和 ` detail.js` ，它们的代码如下：

` // index.js代码 document.getElementById( 'root' ).innerHTML = 'this is index page!' // list.js代码 document.getElementById( 'root' ).innerHTML = 'this is list page!' // detail.js代码 document.getElementById( 'root' ).innerHTML = 'this is detail page!' 复制代码`

运行 ` npm run build` 进行打包：

` $ npm run build 复制代码`

打包后的 ` dist` 目录：

` |-- dist | |-- detail.dae2986ea47c6eceecd6.js | |-- detail.dae2986ea47c6eceecd6.js.map | |-- detail.html | |-- index.ca8e3d1b5e23e645f832.js | |-- index.ca8e3d1b5e23e645f832.js.map | |-- index.html | |-- list.5 f40def0946028db30ed.js | |-- list.5 f40def0946028db30ed.js.map | |-- list.html 复制代码`

随机选择 ` list.html` 在浏览器中运行，结果如下：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd52389ba02ce?imageView2/0/w/1280/h/960/ignore-error/1)

思考：现在只有三个页面，即我们要配置三个入口+三个对应的 ` html` ，如果我们有十个入口，那么我们也要这样做重复的劳动吗？有没有什么东西能帮助我们自动实现呢？答案当然是有的！

我们首先定义一个 ` makeHtmlPlugins` 方法，它接受一个 Webpack 配置项的参数 ` configs` ，返回一个 ` plugins` 数组

` const makeHtmlPlugins = function ( configs ) { const htmlPlugins = [] Object.keys(configs.entry).forEach( key => { htmlPlugins.push( new htmlWebpackPlugin({ template : 'src/index.html' , filename : ` ${key}.html` , chunks : [key] }) ) }) return htmlPlugins } 复制代码`

通过调用 ` makeHtmlPlugins` 方法，它返回一个 ` html` 的 ` plugins` 数组，把它和原有的 ` plugin` 进行合并后再复制给 ` configs`

` configs.plugins = configs.plugins.concat(makeHtmlPlugins(configs)); module.exports = configs; 复制代码`

以上配置完毕后，打包结果依然还是一样的，请自行测试，以下是 ` webpack.commom.js` 完整的代码：

` const path = require ( 'path' ); const webpack = require ( 'webpack' ); const htmlWebpackPlugin = require ( 'html-webpack-plugin' ); const cleanWebpackPlugin = require ( 'clean-webpack-plugin' ); const miniCssExtractPlugin = require ( 'mini-css-extract-plugin' ); const optimizaCssAssetsWebpackPlugin = require ( 'optimize-css-assets-webpack-plugin' ); const configs = { entry : { index : './src/index.js' , list : './src/list.js' , detail : './src/detail.js' }, module : { rules : [ { test : /\.css$/ , use : [ { loader : miniCssExtractPlugin.loader, options : { hmr : true , reloadAll : true } }, 'css-loader' ] }, { test : /\.js$/ , exclude : /node_modules/ , loader : [ { loader : "babel-loader" }, { loader : "imports-loader?this=>window" } ] } ] }, plugins : [ new cleanWebpackPlugin(), new miniCssExtractPlugin({ filename : '[name].css' }), new webpack.ProvidePlugin({ '$' : 'jquery' , '_' : 'lodash' }) ], optimization : { splitChunks : { chunks : 'all' }, minimizer : [ new optimizaCssAssetsWebpackPlugin() ] }, output : { filename : '[name].js' , path : path.resolve(__dirname, '../dist' ) } } const makeHtmlPlugins = function ( configs ) { const htmlPlugins = [] Object.keys(configs.entry).forEach( key => { htmlPlugins.push( new htmlWebpackPlugin({ template : 'src/index.html' , filename : ` ${key}.html` , chunks : [key] }) ) }) return htmlPlugins } configs.plugins = configs.plugins.concat(makeHtmlPlugins(configs)) module.exports = configs 复制代码`

# 如何打包一个库文件(Library) #

在上面所有的 Webpack 配置中，几乎都是针对业务代码的，如果我们要打包发布一个库，让别人使用的话，该怎么配置？在下面的几个小节中，我们将来讲一讲该怎么样打包一个库文件，并让这个库文件在多种场景能够使用。

### 创建一个全新的项目 ###

步骤：

* 创建library项目
* 使用 ` npm init -y` 进行配置 ` package.json`
* 新建 ` src` 目录，创建 ` math.js` 文件、 ` string.js` 文件、 ` index.js` 文件
* 根目录下创建 ` webpack.config.js` 文件
* 安装 ` webpack` 、 ` webpack-cli` :::

按上面的步骤走完后，你的目录大概看起来是这样子的：

` |-- src | |-- index.js | |-- math.js | |-- string.js |-- webpack.config.js |-- package.json 复制代码`

### 初始化package.json ###

` // 初始化后，改写package.json { "name": "library", "version": "1.0.0", "description": "", "main": "index.js", "scripts": { "build": "webpack" }, "keywords": [], "author": "", "license": "MIT" } 复制代码`

### 创建src目录，并添加文件 ###

在 ` src` 目录下新建 ` math.js` ，它的代码是四则混合运算的方法，如下：

` export function add ( a, b ) { return a + b; } export function minus ( a, b ) { return a - b; } export function multiply ( a, b ) { return a * b; } export function division ( a, b ) { return a / b; } 复制代码`

在 ` src` 目录下新建 ` string.js` ，它有一个 ` join` 方法，如下：

` export function join ( a, b ) { return a + '' + b; } 复制代码`

在 ` src` 目录下新建 ` index.js` 文件，它引用 ` math.js` 和 ` string.js` 并导出，如下：

` import * as math from './math' ; import * as string from './string' ; export default { math, string }; 复制代码`

### 添加webpack.config.js ###

因为我们是要打包一个库文件，所以 ` mode` 只配置为生产环境( ` production` )即可。

在以上文件添加完毕后，我们来配置一下 ` webpack.config.js` 文件，它的代码非常简单，如下：

` const path = require ( 'path' ); module.exports = { mode : 'production' , entry : './src/index.js' , output : { filename : 'library.js' , path : path.resolve(__dirname, 'dist' ) } } 复制代码`

### 安装Webpack ###

因为涉及到 Webpack 打包，所以我们需要使用 ` npm instll` 进行安装：

` $ npm install webpack webpack-cli -D 复制代码`

### 进行第一次打包 ###

使用 ` npm run build` 进行第一次打包，在 ` dist` 目录下会生成一个叫 ` library.js` 的文件，我们要测试这个文件的话，需要在 ` dist` 目录下新建 ` index.html`

` $ npm run build $ cd dist $ touch index.html 复制代码`

在 ` index.html` 中引入 ` library.js` 文件：

` < script src = "./library.js" > </ script > 复制代码`

至此，我们已经基本把项目目录搭建完毕，现在我们来考虑一下，可以在哪些情况下使用我们打包的文件：

* 使用 ` ES Module` 语法引入，例如 ` import library from 'library'`
* 使用 ` CommonJS` 语法引入，例如 ` const library = require('library')`
* 使用 ` AMD` 、 ` CMD` 语法引入，例如 ` require(['library'], function() {// todo})`
* 使用 ` script` 标签引入，例如 ` <script src="library.js"></script>`

> 
> 
> 
> 针对以上几种使用场景，我们可以在output中配置library和libraryTarget属性(注意：这里的library和libraryTarget和我们的库名字library.js没有任何关系，前者是Webpack固有的配置项，后者只是我们随意取的一个名字)
> 
> 
> 

` const path = require('path'); module.exports = { mode: 'production', entry: './src/index.js', output: { filename: '[name].js', path: path.resolve(__dirname, 'dist'), library: 'library', libraryTarget: 'umd' } } 复制代码`

**配置属性说明：**

* **` library`** ：这个属性指，我们库的全局变量是什么，类似于 ` jquery` 中的 ` $` 符号
* **` libraryTarget`** : 这个属性指，我们库应该支持的模块引入方案， ` umd` 代表支持 ` ES Module` 、 ` CommomJS` 、 ` AMD` 以及 ` CMD`

在配置完毕后，我们再使用 ` npm run build` 进行打包，并在浏览器中运行 ` index.html` ，在 ` console` 控制台输出 ` library` 这个全局变量，结果如下图所示:

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd526a52c245b?imageView2/0/w/1280/h/960/ignore-error/1)

以上我们所写的库非常简单，在实际的库开发过程中，往往需要使用到一些 **第三方库** ，如果我们不做其他配置的话，第三方库会直接打包进我们的库文件中。

如果用户在使用我们的库文件时，也引入了这个第三方库，就造成了重复引用的问题，那么如何解决这个问题呢？可以在 ` webpack.config.js` 文件中配置 ` externals` 属性。

在 ` string.js` 文件的 ` join` 方法中，我们使用第三方库 ` lodash` 中的 ` _join()` 方法来进行字符串的拼接。

` import _ from 'lodash' ; export function join ( a, b ) { return _.join([a, b], ' ' ); } 复制代码`

在修改完毕 ` string.js` 文件后，使用 ` npm run build` 进行打包，发现 ` lodash` 直接打包进了我们的库文件，造成库文件积极臃肿，有70.8kb。

` $ npm run build Built at: 2019-04-05 00:47:25 Asset Size Chunks Chunk Names library.js 70.8 KiB 0 [emitted] main 复制代码`

针对以上问题，我们可以在 ` webpack.config.js` 中配置 ` externals` 属性，更多 ` externals` 的用法请点击 [externals]( https://link.juejin.im?target=https%3A%2F%2Fwebpack.js.org%2Fconfiguration%2Fexternals%2F%23root )

` const path = require('path'); module.exports = { mode: 'production', entry: './src/index.js', externals: ['lodash'], output: { filename: 'library.js', path: path.resolve(__dirname, 'dist'), library: 'library', libraryTarget: 'umd' } } 复制代码`

配置完 ` externals` 后，我们再进行打包，它的打包结果如下，我们可以看到我们的库文件又变回原来的大小了，证明我们的配置起作用了。

` $ npm run build Built at: 2019-04-05 00:51:22 Asset Size Chunks Chunk Names library.js 1.63 KiB 0 [emitted] main 复制代码`

### 如何发布并使用我们的库文件 ###

在打包完毕后，我们如何发布我们的库文件呢，以下是 **发布的步骤** ：

* 注册 ` npm` 账号
* 修改 ` package.json` 文件的入口，修改为： ` "main": "./dist/library.js"`
* 运行 ` npm adduser` 添加账户名称
* 运行 ` npm publish` 命令进行发布
* 运行 ` npm install xxx` 来进行安装

> 
> 
> 
> 为了维护npm仓库的干净，我们并未实际运行npm
> publish命令，因为我们的库是无意义的，发布上去属于垃圾代码，所以请自行尝试发布。另外自己包的名字不能和npm仓库中已有的包名字重复，所以需要在package.json中给name属性起一个特殊一点的名字才行，例如"name":
> "why-library-2019"
> 
> 

# TypeScript配置 #

随着 ` TypeScript` 的不断发展，相信未来使用 ` TypeScript` 来编写 JS 代码将变成主流形式，那么如何在 Webpack 中配置支持 ` TypeScript` 呢？可以安装 ` ts-loader` 和 ` typescript` 来解决这个问题。

### 新建一个项目webpack-typescript ###

新创建一个项目，命名为 ` webpack-typescript` ，并按如下步骤处理：

* 使用 ` npm init -y` 初始化 ` package.json` 文件，并在其中添加 ` build` Webpack打包命令
* 新建 ` webpack.config.js` 文件，并做一些简单配置，例如 ` entry` 、 ` output` 等
* 新建 ` src` 目录，并在 ` src` 目录下新建 ` index.ts` 文件
* 新建 ` tsconfig.json` 文件，并做一些配置
* 安装 ` webpack` 和 ` webpack-cli`
* 安装 ` ts-loader` 和 ` typescript`

按以上步骤完成后，项目目录大概如下所示：

` |-- src | |-- index.ts |-- tsconfig.json |-- webpack.config.js |-- package.json 复制代码`

在 ` package.json` 中添加好打包命令命令：

` "scripts" : { "build" : "webpack" }, 复制代码`

接下来我们需要对 ` webpack.config.js` 做一下配置：

` const path = require('path'); module.exports = { mode: 'production', module: { rules: [ { test: /\.(ts|tsx)?$/, use: { loader: 'ts-loader' } } ] }, entry: { main: './src/index.ts' }, output: { filename: '[name].js', path: path.resolve(__dirname, 'dist') } } 复制代码`

在 ` tsconfig.json` 里面进行 ` typescript` 的相关配置，配置项的说明如下

* ` module` : 表示我们使用 ` ES6` 模块
* ` target` : 表示我们转换成 ` ES5` 代码
* ` allowJs` : 允许我们在 `.ts` 文件中通过 ` import` 语法引入其他 `.js` 文件

` { "compilerOptions" : { "module" : "ES6" , "target" : "ES5" , "allowJs" : true } } 复制代码`

在 ` src/index.ts` 文件中书写 ` TypeScript` 代码，像下面这样

` class Greeter { greeting: string constructor (message: string) { this.greeting = message; } greet() { return 'hello, ' + this.greeting; } } let greeter = new Greeter( 'why' ); console.log(greeter.greet()); 复制代码`

### 打包测试 ###

* 运行 ` npm run build` 进行打包
* 在生成 ` dist` 目录下，新建 ` index.html` ，并引入打包后的 ` main.js` 文件
* 在浏览器中运行 ` index.html`

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd52986b2b155?imageView2/0/w/1280/h/960/ignore-error/1)

### 使用其他模块的类型定义文件 ###

> 
> 
> 
> 如果我们要使用lodash库，必须安装其对应的类型定义文件，格式为@types/xxx
> 
> 

安装 ` lodash` 对应的 ` typescript` 类型文件：

` $ npm install lodash @types/lodash -D 复制代码`

安装完毕后，我们在 ` index.ts` 中引用 ` lodash` ，并使用里面的方法：

` import * as _ from 'lodash' class Greeter { greeting: string constructor(message: string) { this.greeting = message; } greet() { return _.join(['hello', this.greeting], '**'); } } let greeter = new Greeter('why'); console.log(greeter.greet()); 复制代码`

### 打包测试 ###

使用 ` npm run build` ，在浏览器中运行 ` index.html` ，结果如下：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd52c0ad4099f?imageView2/0/w/1280/h/960/ignore-error/1)

# Webpack性能优化 #

## 打包分析 ##

在进行 Webpack 性能优化之前，如果我们知道我们每一个打包的文件有多大，打包时间是多少，它对于我们进行性能优化是很有帮助的，这里我们使用 ` webpack-bundle-analyzer` 来帮助我们解决这个问题。

首先需要使用如下命令去安装这个插件：

` $ npm install webpack-bundle-analyzer --save-dev 复制代码`

安装完毕后，我们需要在 ` webpack.prod.js` 文件中做一点小小的改动：

` const BundleAnalyzerPlugin = require ( 'webpack-bundle-analyzer' ).BundleAnalyzerPlugin; const prodConfig = { // 其它配置项 mode: 'production' , plugins : [ new BundleAnalyzerPlugin() ] } 复制代码`

配置完毕后，我们运行 ` npm run build` 命令来查看打包分析结果，以下打包结果仅供参考：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd52f056d0a96?imageView2/0/w/1280/h/960/ignore-error/1)

## 缩小文件的搜索范围 ##

首先我们要弄明白 Webpack 的一个配置参数( ` Resolve` )的作用：它告诉了 Webpack 怎么去搜索文件，它同样有几个属性需要我们去理解：

* ` extensions` ：它告诉了 Webpack 当我们在导入模块，但没有写模块的后缀时应该如何去查找模块。
* ` mainFields` ：它告诉了 Webpack 当我们在导入模块，但并没有写模块的具体名字时，应该如何去查找这个模块。
* ` alias` ：当我们有一些不得不引用的第三方库或者模块的时候，可以通过配置别名，直接引入它的 `.min.js` 文件，这样可以库内的直接解析
* 其它 ` include` 、 ` exclude` 、 ` test` 来配合loader进行限制文件的搜索范围

### extensions参数 ###

就像上面所说的那样， ` extensions` 它告诉了 Webpack 当我们在导入模块，但没有写模块的后缀时，应该如何去查找模块。这种情况在我们开发中是很常见的，一个情形可能如下所示：

` // 书写了模块后缀 import main from 'main.js' // 没有书写模块后缀 import main from 'main' 复制代码`

像上面那样，我们不写 ` main.js` 的 `.js` 后缀，是因为 Webpack 会默认帮我们去查找一些文件，我们也可以去配置自己的文件后缀配置：

> 
> 
> 
> extensions参数应尽可能只配置主要的文件类型，不可为了图方便写很多不必要的，因为每多一个，底层都会走一遍文件查找的工作，会损耗一定的性能。
> 
> 

` module.exports = { // 其它配置 resolve: { extensions : [ '.js' , '.json' , '.vue' ] } } 复制代码`

如果我们像上面配置后，我们可以在代码中这样写：

` // 省略 .vue文件扩展 import BaseHeader from '@/components/base-header' ; // 省略 .json文件扩展 import CityJson from '@/static/city' ; 复制代码`

### mainFields参数 ###

` mainFields` 参数主要应用场景是，我们可以不写具体的模块名称，由 Webpack 去查找，一个可能的情形如下:

` // 省略具体模块名称 import BaseHeader from '@components/base-header/' ; // 以上相当于这一段代码 import BaseHeader from '@components/base-header/index.vue' ; // 或者这一段 import BaseHeader from '@components/base-header/main.vue' ; 复制代码`

我们也可以去配置自己的 ` mainFields` 参数：

> 
> 
> 
> 同extensions参数类似，我们也不建议过多的配置mainFields的值，原因如上。
> 
> 

` module.exports = { // 其它配置 resolve: { extensions : [ '.js' , '.json' , '.vue' ], mainFields : [ 'main' , 'index' ] } } 复制代码`

### alias参数 ###

` alias` 参数更像一个别名，如果你有一个目录很深、文件名很长的模块，为了方便，配置一个别名这是很有用的；对于一个庞大的第三方库，直接引入 `.min.js` 而不是从 ` node_modules` 中引入也是一个极好的方案，一个可能得情形如下：

> 
> 
> 
> 通过别名配置的模块，会影响Tree Shaking，建议只对整体性比较强的库使用，像lodash库不建议通过别名引入，因为lodash使用Tree
> Shaking更合适。
> 
> 

` // 没有配置别名之前 import main from 'src/a/b/c/main.js' ; import React from 'react' ; // 配置别名之后 import main from 'main.js' ; import React from 'react' ; 复制代码` ` // 别名配置 const path = require ( 'path' ); module.exports = { // 其它配置 resolve: { extensions : [ '.js' , '.json' , '.vue' ], mainFields : [ 'main' , 'index' ], alias : { main : path.resolve(__dirname, 'src/a/b/c' ), react : path.resolve(__dirname, './node_modules/react/dist/react.min.js' ) } } } 复制代码`

## Tree Shaking去掉冗余的代码 ##

` Tree Shaking` 配置我们已经在上面讲过，配置 ` Tree Shaking` 也很简单。

` module.exports = { // 其它配置 optimization: { usedExports : true } } 复制代码`

如果你对 ` Tree Shaking` 还不是特别理解，请点击 [Tree Shaking]( https://link.juejin.im?target=%2Fwebpack%2F%23tree-shaking ) 阅读更多。

## DllPlugin减少第三方库的编译次数 ##

对于有些固定的第三方库，因为它是固定的，我们每次打包，Webpack 都会对它们的代码进行分析，然后打包。那么有没有什么办法，让我们只打包一次，后面的打包直接使用第一次的分析结果就行。答案当然是有的，我们可以使用 Webpack 内置的 ` DllPlugin` 来解决这个问题，解决这个问题可以分如下的步骤进行：

* 把第三方库单独打包在一个 ` xxx.dll.js` 文件中
* 在 ` index.html` 中使用 ` xxx.dll.js` 文件
* 生成第三方库的打包分析结果保存在 ` xxx.manifest.json` 文件中
* 当 ` npm run build` 时，引入已经打包好的第三方库的分析结果
* 优化

### 单独打包第三方库 ###

为了单独打包第三方库，我们需要进行如下步骤：

* 根目录下生成 ` dll` 文件夹
* 在 ` build` 目录下生成一个 ` webpack.dll.js` 的配置文件，并进行配置。
* 在 ` package.json` 文件中，配置 ` build:dll` 命令
* 使用 ` npm run build:dll` 进行打包

生成 ` dll` 文件夹：

` $ mkdir dll 复制代码`

在 ` build` 文件夹下生成 ` webpack.dll.js` :

` $ cd build $ touch webpack.dll.js 复制代码`

创建完毕后，需要在 ` webpack.dll.js` 文件中添加如下代码：

` const path = require ( 'path' ); module.exports = { mode : 'production' , entry : { vendors : [ 'lodash' , 'jquery' ] }, output : { filename : '[name].dll.js' , path : path.resolve(__dirname, '../dll' ), library : '[name]' } } 复制代码`

最后需要在 ` package.json` 文件中添加新的打包命令：

` { // 其它配置 "scripts": { "dev": "webpack-dev-server --config ./build/webpack.dev.js", "build": "webpack --config ./build/webpack.prod.js", "build:dll": "webpack --config ./build/webpack.dll.js" } } 复制代码`

使用 ` npm run build:dll` 打包结果，你的打包结果看起来是下面这样的：

` |-- build | |-- webpack.common.js | |-- webpack.dev.js | |-- webpack.dll.js | |-- webpack.prod.js |-- dll | |-- vendors.dll.js |-- src | |-- index.html | |-- index.js |-- package.json 复制代码`

### 引用 ` xxx.dll.js` 文件 ###

在上一小节中我们成功拿到了 ` xxx.dll.js` 文件，那么如何在 ` index.html` 中引入这个文件呢？答案是需要安装 ` add-asset-html-webpack-plugin` 插件：

` $ npm install add-asset-html-webpack-plugin -D 复制代码`

在 ` webpack.common.js` 中使用 ` add-asset-html-webpack-plugin` 插件：

` const addAssetHtmlWebpackPlugin = require ( 'add-asset-html-webpack-plugin' ); const configs = { // 其它配置 plugins: [ new addAssetHtmlWebpackPlugin({ filepath : path.resolve(__dirname, '../dll/vendors.dll.js' ) }) ] } module.exports = configs; 复制代码`

我们将第三方库全局暴露了一个 ` vendors` 变量，现引入 ` xxx.dll.js` 文件结果如下所示：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd5158e15270e?imageView2/0/w/1280/h/960/ignore-error/1)

### 生成打包分析文件 ###

在 ` webpack.dll.js` 中使用 Webpack 内置的 ` DllPlugin` 插件，进行打包分析：

` const path = require('path'); const webpack = require('webpack'); module.exports = { mode: 'production', entry: { vendors: ['lodash', 'jquery'] }, output: { filename: '[name].dll.js', path: path.resolve(__dirname, '../dll'), library: '[name]' }, plugins: [ new webpack.DllPlugin({ name: '[name]', path: path.resolve(__dirname, '../dll/[name].manifest.json') }) ] } 复制代码`

### 引用打包分析文件 ###

在 ` webpack.common.js` 中使用 Webpack 内置的 ` DllReferencePlugin` 插件来引用打包分析文件：

` const htmlWebpackPlugin = require('html-webpack-plugin'); const cleanWebpackPlugin = require('clean-webpack-plugin'); const addAssetHtmlWebpackPlugin = require('add-asset-html-webpack-plugin'); const webpack = require('webpack'); const path = require('path'); module.exports = { // 其它配置 plugins: [ new cleanWebpackPlugin(), new htmlWebpackPlugin({ template: 'src/index.html' }), new addAssetHtmlWebpackPlugin({ filepath: path.resolve(__dirname, '../dll/vendors.dll.js') }), new webpack.DllReferencePlugin({ manifest: path.resolve(__dirname, '../dll/vendors.manifest.json') }) ] } 复制代码`

### 优化 ###

现在我们思考一个问题，我们目前是把 ` lodash` 和 ` jquery` 全部打包到了 ` vendors` 文件中，那么如果我们要拆分怎么办，拆分后又该如何去配置引入？一个可能的拆分结果如下：

` const path = require('path'); const webpack = require('webpack'); module.exports = { mode: 'production', entry: { vendors: ['lodash'], jquery: ['jquery'] }, output: { filename: '[name].dll.js', path: path.resolve(__dirname, '../dll'), library: '[name]' }, plugins: [ new webpack.DllPlugin({ name: '[name]', path: path.resolve(__dirname, '../dll/[name].manifest.json') }) ] } 复制代码`

根据上面的拆分结果，我们需要在 ` webpack.common.js` 中进行如下的引用配置：

` const htmlWebpackPlugin = require('html-webpack-plugin'); const cleanWebpackPlugin = require('clean-webpack-plugin'); const addAssetHtmlWebpackPlugin = require('add-asset-html-webpack-plugin'); const path = require('path'); const configs = { // ... 其他配置 plugins: [ new cleanWebpackPlugin(), new htmlWebpackPlugin({ template: 'src/index.html' }), new addAssetHtmlWebpackPlugin({ filepath: path.resolve(__dirname, '../dll/vendors.dll.js') }), new addAssetHtmlWebpackPlugin({ filepath: path.resolve(__dirname, '../dll/jquery.dll.js') }), new webpack.DllReferencePlugin({ manifest: path.resolve(__dirname, '../dll/vendors.manifest.json') }), new webpack.DllReferencePlugin({ manifest: path.resolve(__dirname, '../dll/jquery.manifest.json') }) ] } module.exports = configs; 复制代码`

我们可以发现：随着我们引入的第三方模块越来越多，我们不断的要进行 Webpack 配置文件的修改。对于这个问题，我们可以使用 ` Node` 的核心模块 ` fs` 来分析 ` dll` 文件夹下的文件，进行动态的引入，根据这个思路我们新建一个 ` makePlugins` 方法，它返回一个 Webpack 的一个 ` plugins` 数组：

` const makePlugins = function ( ) { const plugins = [ new cleanWebpackPlugin(), new htmlWebpackPlugin({ template : 'src/index.html' }), ]; // 动态分析文件 const files = fs.readdirSync(path.resolve(__dirname, '../dll' )); files.forEach( file => { // 如果是xxx.dll.js文件 if ( /.*\.dll.js/.test(file)) { plugins.push( new addAssetHtmlWebpackPlugin({ filepath : path.resolve(__dirname, '../dll' , file) }) ) } // 如果是xxx.manifest.json文件 if ( /.*\.manifest.json/.test(file)) { plugins.push( new webpack.DllReferencePlugin({ manifest : path.resolve(__dirname, '../dll' , file) }) ) } }) return plugins; } configs.plugins = makePlugins(configs); module.exports = configs; 复制代码`

使用 ` npm run build:dll` 进行打包第三方库，再使用 ` npm run build` 打包，打包结果如下:

> 
> 
> 
> 本次试验，第一次打包时间为1100ms+，后面的打包稳定在800ms+，说明我们的 Webpack性能优化已经生效。
> 
> 

` |-- build | |-- webpack.common.js | |-- webpack.dev.js | |-- webpack.dll.js | |-- webpack.prod.js |-- dist | |-- index.html | |-- jquery.dll.js | |-- main.1158 fa9f961c50aaea21.js | |-- main.1158 fa9f961c50aaea21.js.map |-- dll | |-- jquery.dll.js | |-- jquery.manifest.json | |-- vendors.dll.js | |-- vendors.manifest.json |-- src | |-- index.html | |-- index.js |-- package.json |-- postcss.config.js 复制代码`

**小结** ：Webpack 性能优化是一个长久的话题，本章也仅仅只是浅尝辄止，后续会有关于 Webpack 更加深入的解读博客，敬请期待(立个flag)。

# 编写自己的Loader #

在我们使用 Webpack 的过程中，我们使用了很多的 ` loader` ，那么那些 ` loader` 是哪里来的？我们能不能写自己的 ` loader` 然后使用？ 答案当然是可以的，Webpack 为我们提供了一些 ` loader` 的API，通过这些API我们能够编写出自己的 ` loader` 并使用。

## 如何编写及使用自己的Loader ##

场景: 我们需要把 `.js` 文件中，所有出现 ` Webpack is good!` ，改成 ` Webpack is very good!` 。实际上我们需要编写自己的 ` loader` ，所以我们有如下的步骤需要处理：

* 新建 ` webpack-loader` 项目
* 使用 ` npm init -y` 命令生成 ` package.json` 文件
* 创建 ` webpack.config.js` 文件
* 创建 ` src` 目录，并在 ` src` 目录下新建 ` index.js`
* 创建 ` loaders` 目录，并在 ` loader` 目录下新建 ` replaceLoader.js`
* 安装 ` webpack` 、 ` webpack-cli`

按上面的步骤新建后的项目目录如下：

` |-- loaders | | -- replaceLoader.js |-- src | | -- index.js |-- webpack.config.js |-- package.json 复制代码`

首先需要在 ` webpack.config.js` 中添加下面的代码：

` const path = require ( 'path' ); module.exports = { mode : 'development' , entry : './src/index.js' , module : { rules : [ { test : /\.js$/ , use : [path.resolve(__dirname, './loaders/replaceLoader.js' )] } ] }, output : { filename : '[name].js' , path : path.resolve(__dirname, 'dist' ) } } 复制代码`

随后在 ` package.json` 文件添加 ` build` 打包命令：

` // 其它配置 "scripts": { "build": "webpack" } 复制代码`

接下来在 ` src/index.js` 文件中添加一行代码：这个文件使用最简单的例子，只是打印一句话。

` console.log( 'Webpack is good!' ); 复制代码`

最后就是在 ` loader/replaceLoader.js` 编写我们自己 ` loader` 文件中的代码：

* 编写 ` loader` 时， ` module.exports` 是固定写法，并且它只能是一个普通函数，不能写箭头函数(因为需要 ` this` 指向自身)
* ` source` 是打包文件的源文件内容

` const loaderUtils = require ( 'loader-utils' ); module.exports = function ( source ) { return source.replace( 'good' , 'very good' ); } 复制代码`

使用我们的 ` loader` : 要使用我们的 ` loader` ，则需要在 ` modules` 中写 ` loader` ， ` resolveLoader` 它告诉了 Webpack 使用 ` loader` 时，应该去哪些目录下去找，默认是 ` node_modules` ，做了此项配置后，我们就不用去显示的填写其路径了，因为它会自动去 ` loaders` 文件夹下面去找。

` const path = require('path'); module.exports = { mode: 'development', entry: './src/index.js', resolveLoader: { modules: ['node_modules', './loaders'] }, module: { rules: [ { test: /\.js$/, use: [{ loader: 'replaceLoader', options: { name: 'wanghuayu' } }] } ] }, output: { filename: '[name].js', path: path.resolve(__dirname, 'dist') } } 复制代码`

最后我们运行 ` npm run build` ，在生成的 ` dist` 目录下打开 ` main.js` 文件，可以看到文件内容已经成功替换了，说明我们的 ` loader` 已经使用成功了。

` /***/ "./src/index.js" : /*!**********************!*\ !*** ./src/index.js ***! \**********************/ /*! no static exports found */ /***/ ( function ( module, exports ) { eval ( "console.log('Webpack is very good!');\n\n//# sourceURL=webpack:///./src/index.js?" ); /***/ }) /******/ }); 复制代码`

## 如何向Loader传参及返回多个值 ##

问题：

* 我们如何返回多个值？
* 我们如何向自己的Loader传递参数？

### 如何返回多个值 ###

Webpack 的 API允许我们使用 ` callback(error, result, sourceMap?, meta?)` 返回多个值，它有四个参数：

* ` Error || Null` ：错误类型， 没有错误传递 ` null`
* ` result` ：转换后的结果
* ` sourceMap` ：可选参数，处理分析后的 ` sourceMap`
* ` meta` : 可选参数，元信息

返回多个值，可能有如下情况：

` // 第三，第四个参数是可选的。 this.callback( null , result); 复制代码`

### 如何传递参数 ###

我们知道在使用 ` loader` 的时候，可以写成如下的形式：

` // options里面可以传递一些参数 { test : /\.js$/ , use : [{ loader : 'replaceLoader' , options : { word : 'very good' } }] } 复制代码`

再使用 ` options` 传递参数后，我们可以使用官方提供的 [loader-utils]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fwebpack%2Floader-utils ) 来获取 ` options` 参数，可以像下面这样写：

` const loaderUtils = require ( 'loader-utils' ); module.exports = function ( source ) { var options = loaderUtils.getOptions( this ); return source.replace( 'good' , options.word) } 复制代码`

## 如何在Loader中写异步代码 ##

在上面的例子中，我们都是使用了同步的代码，那么如果我们有必须异步的场景，该如何实现呢？我们不妨做这样的假设，先写一个 ` setTimeout` ：

` const loaderUtils = require ( 'loader-utils' ); module.exports = function ( source ) { var options = loaderUtils.getOptions( this ); setTimeout( () => { var result = source.replace( 'World' , options.name); return this.callback( null , result); }, 0 ); } 复制代码`

如果你运行了 ` npm run build` 进行打包，那么一定会报错，解决办法是：使用 ` this.async()` 主动标识有异步代码：

` const loaderUtils = require('loader-utils'); module.exports = function(source) { var options = loaderUtils.getOptions(this); var callback = this.async(); setTimeout(() => { var result = source.replace('World', options.name); callback(null, result); }, 0); } 复制代码`

至此，我们已经掌握了如何编写、如何引用、如何传递参数以及如何写异步代码，在下一小节当中我们将学习如何编写自己的 ` plugin` 。

# 编写自己的Plugin #

与 ` loader` 一样，我们在使用 Webpack 的过程中，也经常使用 ` plugin` ，那么我们学习如何编写自己的 ` plugin` 是十分有必要的。 场景：编写我们自己的 ` plugin` 的场景是在打包后的 ` dist` 目录下生成一个 ` copyright.txt` 文件

## plugin基础 ##

` plugin` 基础讲述了怎么编写自己的 ` plugin` 以及如何使用，与创建自己的 ` loader` 相似，我们需要创建如下的项目目录结构：

` |-- plugins | -- copyWebpackPlugin.js |-- src | -- index.js |-- webpack.config.js |-- package.json 复制代码`

` copyWebpackPlugins.js` 中的代码：使用 ` npm run build` 进行打包时，我们会看到控制台会输出 ` hello, my plugin` 这段话。

> 
> 
> 
> plugin与loader不同，plugin需要我们提供的是一个类，这也就解释了我们必须在使用插件时，为什么要进行new操作了。
> 
> 

` class copyWebpackPlugin { constructor () { console.log( 'hello, my plugin' ); } apply(compiler) { } } module.exports = copyWebpackPlugin; 复制代码`

` webpack.config.js` 中的代码：

` const path = require ( 'path' ); // 引用自己的插件 const copyWebpackPlugin = require ( './plugins/copyWebpackPlugin.js' ); module.exports = { mode : 'development' , entry : './src/index.js' , output : { filename : '[name].js' , path : path.resolve(__dirname, 'dist' ) }, plugins : [ // new自己的插件 new copyWebpackPlugin() ] } 复制代码`

## 如何传递参数 ##

在使用其他 ` plugin` 插件时，我们经常需要传递一些参数进去，那么我们如何在自己的插件中传递参数呢？在哪里接受呢？
其实，插件传参跟其他插件传参是一样的，都是在构造函数中传递一个对象，插件传参如下所示：

` const path = require('path'); const copyWebpackPlugin = require('./plugins/copyWebpackPlugin.js'); module.exports = { mode: 'development', entry: './src/index.js', output: { filename: '[name].js', path: path.resolve(__dirname, 'dist') }, plugins: [ // 向我们的插件传递参数 new copyWebpackPlugin({ name: 'why' }) ] } 复制代码`

在 ` plugin` 的构造函数中调用：使用 ` npm run build` 进行打包，在控制台可以打印出我们传递的参数值 ` why`

` class copyWebpackPlugin { constructor(options) { console.log(options.name); } apply(compiler) { } } module.exports = copyWebpackPlugin; 复制代码`

## 如何编写及使用自己的Plugin ##

* ` apply` 函数是我们插件在调用时，需要执行的函数
* ` apply` 的参数，指的是 Webpack 的实例
* ` compilation.assets` 打包的文件信息

我们现在有这样一个需求：使用自己的插件，在打包目录下生成一个 ` copyright.txt` 版权文件，那么该如何编写这样的插件呢？ 首先我们需要知道 ` plugin` 的钩子函数，符合我们规则钩子函数叫： ` emit` ，它的用法如下：

` class CopyWebpackPlugin { constructor () { } apply(compiler) { compiler.hooks.emit.tapAsync( 'CopyWebpackPlugin' , (compilation, cb) => { var copyrightText = 'copyright by why' ; compilation.assets[ 'copyright.txt' ] = { source : function ( ) { return copyrightText }, size : function ( ) { return copyrightText.length; } } cb(); }) } } module.exports = CopyWebpackPlugin; 复制代码`

使用 ` npm run build` 命名打包后，我们可以看到 ` dist` 目录下，确实生成了我们的 ` copyright.txt` 文件。

` |-- dist | |-- copyright.txt | |-- main.js |-- plugins | |-- copyWebpackPlugin.js |-- src | |-- index.js |-- webpack.config.js |-- package.json 复制代码`

我们打开 ` copyright.txt` 文件，它的内容如下：

` copyright by why 复制代码`

本篇博客由慕课网视频 [从基础到实战手把手带你掌握新版Webpack4.0]( https://link.juejin.im?target=https%3A%2F%2Fcoding.imooc.com%2Fclass%2F316.html ) 阅读整理而来，观看视频请支持正版。