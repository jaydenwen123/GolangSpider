# 12个常见的实战场景 - [webpack 第二篇] #

## webpack实战系列全目录 ##

* 
* [ webpack 12个常见的实际场景]( https://juejin.im/post/5cda6f67e51d453ccd246501 )
* webpack15个常见的优化策略【敬请期待】
* webpack从0打造兼容ie8的脚手架【敬请期待】
* webpack面试全总结【敬请期待】

本节我们将要说到的实战场景目录：

> 
> * 入门配置
> * 创建单页面应用
> * 接入babel
> * 接入scss
> * 接入vue
> * 分离javascript与css
> * 压缩javascript
> * 压缩css
> * 提取javascript公共代码
> * 加入代码规范检测
> * 搭建本地服务
> * 创建多页面应用
> 

## 一： 入门配置 ##

目录结构如下：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afd343f2cc3cce?imageView2/0/w/1280/h/960/ignore-error/1)

webpack.config.js配置如下：

` const path = require( 'path' ); module.exports = { entry: './main.js' , output: { filename: 'bundle.js' , path: path.resolve(__dirname, 'dist' ) //必须是绝对路径 } }; 复制代码`

以上是一个最基础的配置，我们接下来一步一步的加入更多的功能.

## 二：单页面应用 ##

#### 认识单页面应用 ####

首先，我们通过第一节，可以知道了怎么将一个入口main.js打包成bundle.js了，那么入口文件肯定是要引用在html中啊，那么，怎么创建html文件？怎么创建build以后的html文件，以及html文件如何引用入口js文件？答案就是 html-webpack-plugin。 [点击查看官方文档]( https://link.juejin.im?target=https%3A%2F%2Fwww.webpackjs.com%2Fplugins%2Fhtml-webpack-plugin%2F )

该插件可以实现哪些功能呢？

* 不用我们手动创建html文件，引用该插件以后，build后会自动生成一个index.html
* 生成的index.html不需要手动引入入口js文件，它会自动创建script标签，并且引用生成的bundle.js

#### 创建单页面应用 ####

第一步：安装插件

` npm i html-webpack-plugin --save-dev 复制代码`

第二步：配置插件

` const HtmlWebpackPlugin = require( 'html-webpack-plugin' ); plugins: [ new HtmlWebpackPlugin({ filename: 'index.html' , template: './index.html' }) ] 说明： template：表示模版，即以哪个html文件为模，在dist目录下生成新的html文件 filename: 即编译以后生成的html文件的名字 复制代码`

注意：上面我们说到不用我们手动创建一个html文件，就可以在编译后自动生成一个index.html， 当然，我们也可以创建一个html模版文件，然后通过template属性引入， 这样编译后生成的index.html内容就是我们自己创建的模版html文件的内容)

以下是我们生成的index.html文件：

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0366ee6008778?imageView2/0/w/1280/h/960/ignore-error/1)

## 三： 接入babel ##

#### 1. 认识babel ####

babel完成了两件事情：

> 
> * 它是一个javascript解释器，它可以将es6代码转换为es5代码，让我们在使用语言新特性的时候，不用担心兼容性问题。
> * 它可以通过插件机制，根据需求灵活的拓展。
> 

在 Babel 执行编译的过程中，会从项目根目 录下 的 .babelre 文件中读取配置。 .babelrc 是一个 JSON 格式的文件，内容大致如下:

` { "presets" : [ [ "env" , { "modules" : false , "targets" : { "browsers" : [ "> 1%" , "last 2 versions" , "not ie <= 8" ] } }], "stage-2" , "react" ], "plugins" : [ "transform-vue-jsx" , "transform-runtime" ] } 复制代码`

我们来认识一下这两个属性：

* presets: 告诉babel要转换的源代码用了哪些新的语法特性，例如，react，es6等都可以在presets去声明，表示我们的源代码中使用了react，es6等语法新特性。
* plugins: 用于告诉babel使用了哪些插件，然后babel就可以通过这些插件去控制如何转换代码。常用的的插件是：transform-runtime， 全名是：babel-plugin-transform-runtime，即在前面加上了 babel-plugin-。它的作用是用于减少冗余的代码， 点击查看

#### 2. 接入babal ####

第一步：在webpack中配置babel-loader。

` module.exports = { module: { rules : [ { test : /\.js$/ , use : [ 'babel-loader' ], } //此处还可以在rules下配置其他loader ] } } 复制代码`

第二步：在根目录下创建.babelrc文件，配置presets和plugins

` { "presets" : [ "env" ], "plugins" : [] } 复制代码`

注意：webpack4.x支持在package.json中直接配置babel属性，不用新建.babelrc即可

` "babel" : { "presets" : [ "env" ], "plugins" : [] } 复制代码`

第三步：安装babel相关依赖

` # Webpack 接入 Babel 必须依赖的模块 npm i -D babel-core babel-loader #根据我们的需求选择不同的 Plugins 或 Presets npm i -D babel-preset-env 复制代码`

提示：以上实例代码主要用于说明整个接入webpack的过程，在不同场景下，具体的配置和安装依赖情况也会有所不同。

## 四：SCSS ##

#### 认识scss ####

scss 即为了我们书写css更方便且更有效率，出现了less，sass，scss等css的预处理器，那么，在最后部署上线的时候，依然要转成css，才可以在浏览器端运行。

#### 接入scss ####

第一步：安装依赖

` #安装 Webpack Loader 依赖 npm i -D sass-loader css-loader style-loader # sass-loader 依赖 node-sass npm i 一D node-sass 复制代码`

第二步：在module中配置

` module:{ rules:[ { test : /\.scss$/, use: [ "style-loader" , "css-loader" , "sass-loader" ] } ] } 复制代码`

说明:

* 当有多个loader的时候，编译顺序是从后往前
* sass-loader: 首先通过sass-loader将scss代码转换成css代码，再交给css-loader处理
* css-loader 会找出 css 代码中 eimport 和 url ()这样的导入语句，告诉 Webpack 依赖这些资源 。 同时支持 CSS Modules、压缩 css 等功能 。处理完后再将结果交给 style-loader处理。
* style-loader会将 css代码转换成字符串后，注入 JavaScript代码中，通过 JavaScript 向 DOM 增加样式。

以上是配置wepbakc接入scss的基本配置， 除此之外，我们还可以加入优化配置：

* 压缩css
* css与js代码分离：ExtractTextPlugin 这两点在我们项目开发中，也是需要配置的，之后我们会讲到。

## 五：使用vue框架 ##

#### 认识vue ####

大家可能都知道，vue，react等都是近些年来比较流行的mvvm框架，而我们使用的过程中，更多是依赖vue-cli等自带的脚手架去生成项目模版，不需要我们手动去配置，那么接下来，我们就通过webpack去手动接入vue，这样有利于我们更好的理解vue。

#### 接入vue ####

第一步：安装依赖

` # Vue 框架运行需要的库 npm i -S vue #构建所需的依赖 npm i -D vue-loader css-loader vue-template-compiler 复制代码`

说明：

* vue-loader:解析和转换.vue文件，提取出其中的逻辑代码 script、样式代 码 style及 HTML模板 template，再分别将它们交给对应的 Loader去处理。
* css-loader:加载由 飞rue-loader 提取出 的 css 代码 。
* vue-template-compiler:将 vue-loader 提取出的 HTML 模板编译成对应 的可执行的 JavaScript代码，这和 React中的 JSX语法被编译成 JavaScript代码类 似 。 预先编译好 HTML 模板相对于在浏览器中编译 HTML 模板，性能更好 。

第二步：配置loader

` const VueLoaderPlugin = require( 'vue-loader/lib/plugin' ); module:{ rules:[ { test : /\.vue$/, use: 'vue-loader' } ] }, plugins: [ new VueLoaderPlugin() ], 复制代码`

注意：需要配置VueLoaderPlugin 插件，

![](https://user-gold-cdn.xitu.io/2019/5/28/16afdec239a76b91?imageView2/0/w/1280/h/960/ignore-error/1) 提示：Vue-loader在15.*之后的版本都是 vue-loader的使用都是需要伴生 VueLoaderPlugin的

第三步：引入vue文件

` //此时就和vue-cli提供的模版main.js一样去引入vue和vue相关文件 import Vue from 'vue' ; import App from './app.vue' ; new Vue({ el: '#app' , render: h => h(App) }); 复制代码`

## 六：javascript与css处理 ##

#### 1. 分离javascript与css ####

默认情况下，我们打包entry入口文件的时候，文件中所依赖的css等样式模块也会包含js中，那么，如何分离javascript和css，分别存放在不同的文件中呢？

第一步：安装依赖

说明：Webpack 4.x已经不再支持extract-text-webpack-plugin，推荐使用mini-css-webpack-plugin， 如果想继续使用该插件，请 [参考该文档]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_38526769%2Farticle%2Fdetails%2F82427800 )

此处，我们依然用该插件来实现js与css的代码分离：

` //安装插件 npm i extract-text-webpack-plugin@next --save-dev 注意： 后面的@next必须加上，webpack4.x只支持该版本 或者可以用一个新插件： mini-css-extract-plugin 复制代码`

第二步：增加webpack配置

` const ExtractTextPlugin = require ( 'extract-text-webpack-plugin' ); module.exports = { module:{ rules:[ { test : /\.css$/, use: ExtractTextPlugin.extract({ fallback: "style-loader" , use: "css-loader" }) } ] }, plugins: [ new ExtractTextPlugin({ filename: path.resolve(__dirname, '/style/bundle.css' ) //注意可以指定分离出来的css文件的指定目录 }) ] }; 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b034fcaf8e9601?imageView2/0/w/1280/h/960/ignore-error/1) 此时，我们就可以在build以后的dist文件夹里，看到分离出来的css文件了。

例如：下面是一个文件目录结构：

我们在main.js中引入了main.css这个样式文件，那么，通过我们webpack打包，并且进行javascript与css分离，dist中就分别生成了独立的js文件和css文件。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b223cddc886d72?imageView2/0/w/1280/h/960/ignore-error/1)

#### 2. 压缩css ####

通过上面的操作，我们已经将javascript与css分离开来，打开css文件，我们发现，此刻css文件是位被压缩的，

![](https://user-gold-cdn.xitu.io/2019/6/4/16b223f83b102ab3?imageView2/0/w/1280/h/960/ignore-error/1)

那接下来，我们接入css压缩相关配置： 第一步：安装依赖

` npm i optimize-css-assets-webpack-plugin --save-dev 复制代码`

第二步：添加配置

` const OptimizeCssAssetsPlugin = require( 'optimize-css-assets-webpack-plugin' ); module.export = { plugins: [ //只要配置了该插件，即可对分离出来的css进行压缩 new OptimizeCssAssetsPlugin() ] } 复制代码`

最后，我们再次build，发现dist中css已经是被压缩的了。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b224e4191a12cf?imageView2/0/w/1280/h/960/ignore-error/1)

#### 3. 压缩javascript ####

首先说明一点，可能我们会看到很多教程使用的是UglifyJsPlugin，不过是webpack4.x之前的版本可以使用，webpack4.x已经不支持使用移除webpack.optimize.UglifyJsPlugin 压缩配置了, 推荐使用 optimization.minimize 属性替代。

代码如下：

` module.exports = { optimization: { minimize: true //webpack内置属性，默认为 true ， } } 复制代码`

注意：minimize属性默认为 true，即默认情况下build之后的js文件已经是压缩的了，如果我们不想压缩js，可以设置该属性为false，

我们来看看实际压缩前后的区别：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2261806f07bf4?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/6/4/16b2260b292288da?imageView2/0/w/1280/h/960/ignore-error/1)

#### 4. 提取公共javascript ####

一： 首先，我们说明一下为什么要提取公共代码? ：

大型网站通常由多个页面组成， 每个页面都是一个独立的单页应用。 但由于所有页面都 采用同样的技术枝及同 一套样式代码，就导致这些页面之间有很 多相同的代码。

如果每个页面的代码都将这些公共的部分包含进去，则会造成以下问题 。

* 相同的资源被重复加载，浪费用户的流量和服务器的成本。
* 每个页面需要加载的资源太大，导致网页首屏加载缓慢， 影响用户体验。

所以，在实际开发过程中，我们需要把公共的代码抽离成一个独立的文件，这样用户在首次访问网站是就加载了该独立公共文件，此时这些公共文件就被浏览器缓存起来了，切换到其他页面时，该公共文件就不用在重新请求，而是直接访问缓存接口。

二：实际开发中，都需要我们提取哪些公共文件呢

一说到公共文件，我们可能最先想到是一些公共的工具函数所在的common.js文件，这个公共文件中的工具函数会被多个页面同时使用，其实，还有一个base.js也就是最基础的文件，例如，我们vue开发过程中的vue.js,

所以，此处所说的公共文件包含以下两部分：

* common.js （手动创建）
* base.js （引入的第三方依赖：如vue.js，lodash，jquery等）

注意：此处我们为什么要把base.js也单独提取出来，而不是直接包含在common.js中呢？ 答案：为了长期长期缓存base.js文件，

因为base.js中包含的都是一些基础库，而且是指定版本的，我们平时开发过程中，一般不会升级这些基础库的版本，所有只要基础库版本不升级，文件内容就不会变化，hash值也不会更新，那么之前在浏览器中留下的缓存也不会被更新，从而实现长期缓存。而common.js是依据我们开发的内容来的，所以如果开发过程中发生了变化，那么hash也变化，重新上线以后再次访问，原来的缓存的common.js就无法使用了，会使用新的common.js文件。

三： 最后，我们来所以说一下，如何通过webpack提取公共代码？

首先说明以下：webpack4.x提取公共代码推荐使用：webpack内置的SplitChunksPlugin插件，4.x之前所使用的CommomsChunksPlugin已被淘汰

接下来，我们来通过具体的例子试一下：

第一步：我们创建了两个入口文件：main.js 和 index.js， 还有一个公共的common.js文件，同时再安装一个lodash第三方库，然后两个入口文件中，分别引入common.js和lodash;

` // mian.js和index.js import './assets/js/common' ; import 'lodash' ; 复制代码`

第二步：配置多入口文件：

` module.export = { entry: { main: './src/main.js' , index: './src/index.js' }, output: { filename: '[name].bundle.js' , path: path.resolve(__dirname, 'dist' ) } } 复制代码`

第三步：此时，我们执行npm run build命令，生成的dist文件夹如下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26cd06618aef7?imageView2/0/w/1280/h/960/ignore-error/1)

此时，我们引入的lodash和common.js其实分别在main.bundle.js和index.bundle.js各打包了一份，很显然，这样是不对的，那么如何把common.js和lodash分离打包成独立的文件呢？ 答案就是splitChunks属性

第四步：配置splitChunks属性

` module.exports = { optimization: { splitChunks: { chunks: "async" , } } } 复制代码`

此时，我们执行npm run build发现 dist中的文件没有任何变化，并没有将 公共模块分离出来，原因是为什么呢？

**webpack4.x其实默认会将公共文件抽取出来的，只不过chunks属性默认是async, 顾名思义该词是异步的意思，也就是说默认情况下，webpack4.x只会将异步的公共模块分离成独立的文件** , 而我们手动引入的common.js和lodash是同步引入的，所以没有分离出来，所以我们需要把chunks属性改为‘initial’ 或者‘all’

` module.exports = { optimization: { splitChunks: { chunks: "initial" , } } } 复制代码`

我们再来看一下dist目录：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d59a02c1647?imageView2/0/w/1280/h/960/ignore-error/1) 咦，有了，多了一个vendors~index~main.bundle.js， 也就是说common.js和lodash已经分离出来了，都被存放在 vendors~index~main.bundle.js 中，

但是还差一步：如何让common.js和lodash分别分离成独立的文件呢？

首先，我们分析以下原因，其实是minSize属性导致的，它默认值为30000,也就是说webpack4.x只会将文件大小大于3k的公共文件分离出来，而我们目前的common.js大小还不够，所以没有单独分离成一个文件，很显然，此时只需要修改minSize的值即可。

` module.exports = { optimization: { splitChunks: { chunks: "initial" , minSize: 0 } } } 复制代码`

我们再来看看dist目录：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26db2caa14ca2?imageView2/0/w/1280/h/960/ignore-error/1)

啦啦啦，终于搞定了，lodash等属于nodemules中的公共依赖，都被分离到vendors~index~main.bundle.js中，而common.js被分离到main～index.bundle.js中，至此，如何提取公共js文件，大功告成！

## 七：认识npm script ##

Npm Script Chttps://docs.npm s.com/misc/scripts)是一个任务执行者。 Npm 是在安装 Node.js 时附带的包管理器， Npm Script 则是 Npm 内置的 一个功能，允许在 package.json 文件里 使用 scripts 宇段定义任务:

` { "scripts" : { 'dev' : 'node dev.js' , 'build' : 'node bulid.js' } } 复制代码`

以上代码中的 scripts 字段是 一 个对象，每个属性对应一段脚本，以上代码定义 了两 个任务 dev 和 build。 Npm Script 的底层实现原理是通过调用 Shell 去运行脚本命令， 例如执 行 npm run build 命令等同于执行 node build.j s 命令。

同时npm script可以执行node_modules内置的模块：

例如：

* 我们通过npm i webpack --save-dev 安装了webpack
* 此时，我们就可以在npm script配置build命令，其实执行的就是webpack命令
` { "scripts" : { "build" : "webpack" } } 复制代码`

## 八：代码规范检测 ##

#### 前言 ####

检查代码时主要检查以下几项：

* 代码风格:让项目成员强制遵守统一的代码风格，例如如何缩紧、如何写 注释等， 保障代码的可读性，不将时间浪费在争论如何使代码更好看上。
* 潜在问题 : 分析代码在运行过程中可能出现的潜在 Bug。

#### 检查javascript ####

目前最常用的 JavaScript 检查工具是 ESlint ( [eslint.org]( https://link.juejin.im?target=https%3A%2F%2Feslint.org ) )，它不仅内 置了大量的常用检查 规则，还可以通过插件机制做到灵活扩展。

* 单独使用 第一步：安装eslint
` npm i eslint --save-dev //局部安装 或者 npm i eslint -g //全局安装 复制代码`

第二步：创建.eslintrc

` { //从 eslint :recommended 中继承所有检查规则 "extends" : "eslint:recommended" , // 再自定义一些规则 ”rules”:{ //需要在每行结尾加 ; "semi" :[ "error" ， "always" ] , //需要使用 "" 包裹字符串 "quotes" : [ "error" , "double" ] } } 复制代码`

第三步：运行eslint命令

` eslint yourfile . js 复制代码`

该命令的作用就是检查yourfile.js文件的代码格式是否符合eslint的要求，如果有问题会报警提示：例如：

` 296:13 error Strings must use doublequote quotes 298 :7 error Missing semicolon sem工 复制代码` * webpack接入eslint 第一步：安装eslint-loader
` npm i eslint-loader -D 复制代码`

第二步：添加loader

` module: { rules: [ { test : /\.js$/, use: { loader: 'eslint-loader' , options: { formatter: require( 'eslint-friendly-formatter' ) // 默认的错误提示方式 } }, enforce: 'pre' , // 编译前检查 exclude: /node_modules/, // 不检测的文件 include: [__dirname + '/src' ], // 要检查的目录 } ] } 复制代码`

第三步：新建.eslintrc.js

` 复制代码`

#### 检查css ####

stylelint ( [stylelint.io]( https://link.juejin.im?target=https%3A%2F%2Fstylelint.io ) )是目前最成熟的 css 检查工具，在内置了大量检查规则的 同时，也提供了插件机制让用户自定义扩展 。 stylelint 基于 PostCSS，能检查任何 PostCSS 能解析的代码，例如 scss、 Less 等 。

* 单独使用 第一步：安装stylelint
` npm i stylelint --save-dev 或者 ηpm i -g stylelint 复制代码`

第二步：创建.stylelintrc文件

` { //继承 stylelint-config-standard 中所有的检查规则 "extends" : "stylelint-config-standard" , // 再自定义检查规则 "rules" : { "at-rule-empty-line-before" : null } } 复制代码`

第三步：执行命令

` stylelint ”yourfile.css” 复制代码` * webpack中接入stylelint

## 搭建本地开发环境 ##

第一步. 搭建本地服务webpack-dev-server

` npm install webpack-dev-server --save-dev 复制代码`

第二步：在package.json配置npm script命令

` scripts: { "dev" : 'webpack-dev-server --config webpack.dev.config.js' } 复制代码`

说明：此处我们新建一个测试环境的webpack配置文件，用于区分正式环境和测试环境

第三步：配置devServer

` module.exports = merge(baseWebpackConfig, { mode: 'development' , devServer: { port: 9999, contentBase: './dist' , }, plugins: [ ] }); 复制代码`

此时，我们在浏览器中输入：localhost:9999 就可以访问到我们的页面了。

第四步：实现实时刷新 即，只要代码改动，保存以后，浏览器自动刷新

` module.exports = merge(baseWebpackConfig, { mode: 'development' , devServer: { inline: true , //实时刷新 }, plugins: [ ] }); 复制代码`

第五步：实现热更新

` module.exports = merge(baseWebpackConfig, { mode: 'development' , devServer: { hot: true , //热替换 }, plugins: [ // 热更新插件 new webpack.HotModuleReplacementPlugin() ] }); 复制代码`

注意一下三者的区别：

* 开启本地服务：webpack-dev-server可以启动一个本地服务
* 实时刷新：代码改动，浏览器整个页面会自动刷新
* 热更新：代码改动，浏览器不会整个页面都刷新，而是只会在改动的地方局部更新。

## 十：多页面应用 ##

#### 1. 认识多页面应用 ####

首先说明一下单页面和多页面的区别：

* 单页面：只有一个html文件，页面之前的跳转通过路由机制去控制，平时我们使用vue-cli等脚手架自动生成的模版都是单页面，路由通过vue-router去控制跳转。
* 多页面：多个html文件，页面之间的跳转是通过浏览器原生的机制去控制，比如在没有vue，react等框架之前，基本都是多页面的开发。

#### 2. 创建多页面应用 ####

* 

方式1：创建多个入口文件，同时结合 html-webpack-plugin 创建多个html文件，实现多页面应用

* 

方式2: 使用web-webpack-plugin 插件

说明：此处暂不说明具体配置，大家只需清除单页面和多页面的区别，同时，可以采用一下两种方式去实现多页面应用。

## 总结 ##

通过本节的十几个实战场景，相信大家已经基本了解了webpack整个配置的机制，什么场景下采用什么样的配置，当然，大家不必完全记住具体的配置，但要知道什么场景下，用什么样的loader或者plugin等， 同时，个人觉得最重要的是 理解这些场景下所蕴含的思想，例如为什么要区分出现单页面？为什么要分离javascript与css？为什么压缩文件等等，这些才是最关键的，之后，我们专门用一节去讲述webpack背后所包含的这些思想，大家敬请期待吧！