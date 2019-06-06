# webpack学习之路（一）基础配置 #

## webpack是什么？ ##

引自官网：

> 
> 
> 
> 本质上，webpack 是一个现代 JavaScript 应用程序的静态模块打包器(module bundler)。当 webpack
> 处理应用程序时，它会递归地构建一个依赖关系图(dependency graph)，其中包含应用程序需要的每个模块，然后将所有这些模块打包成一个或多个
> bundle。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/4/16b216222f208ff2?imageView2/0/w/1280/h/960/ignore-error/1)

## 目前版本的主要特性 ##

目前最新的版本是v4.32.2，webpack4升级之后，增加了很多新特性

* 不在支持Node.js 4
* **移除CommonChunkPlugin，增加optimization**
* 支持WebAssembly
* 支持多种模块类型
* **增加mode配置**
* **零配置模块打包**
* 更快的构建时间，速度提升了98%

## 配置 ##

### 1.entry ###

webpack是个模块打包机，无论什么资源都会被打包成模块，模块之间是有引用关系的，所以会构建一个关系依赖图，那么我们需要一个入口文件，从这个入口开始一步一步查找依赖关系，加载模块。

我们需要在 ` webpack.config.js` 配置

` module.exports = { entry: './path/to/my/entry/file.js' }; 复制代码`

上面这个例子是一个单入口，也就是单页应用所使用的配置， 那么如果是多页应用，配置就需要改一下

` module.exports = { entry: { index: './path/to/my/entry/list.js' , list: './path/to/my/entry/list.js' } }; 复制代码`

多入口，需要定义成Object对象，key之后还会被使用

### 2.output ###

> 
> 
> 
> ` output` 属性告诉 ` webpack` 在哪里输出它所创建的 ` bundles` ，以及如何命名这些文件，默认值为 `./dist` 。
> 
> 
> 

无论是单入口还是多入口，只能指定一个output

* ` filename` 用于输出文件的文件名。
* 目标输出目录 ` path` 的绝对路径。

单入口的配置如下

` const config = { output: { filename: 'bundle.js' , path: '/home/proj/public/assets' } }; 复制代码`

多入口配置

` { entry: { app: './src/app.js' , search: './src/search.js' }, output: { filename: '[name].js' , path: __dirname + '/dist' } } 复制代码`

` filename` 应该使用占位符来确保每个文件具有唯一的名称，占位符取值就是 ` entry` 里的 ` key`

### 3.mode ###

` mode` 是webpack4新增的特性。 提供两个选择：

+----------------+------------------------------+
|      选项      |             描述             |
+----------------+------------------------------+
| ` development` | 会将 `                       |
|                | process.env.NODE_ENV`        |
|                | 的值设为 `                   |
|                | development` 。启用 `        |
|                | NamedChunksPlugin` 和 `      |
|                | NamedModulesPlugin` 。       |
| ` production`  | 会将 ` process.env.NODE_ENV` |
|                | 的值设为 `                   |
|                | production` 。启用 `         |
|                | FlagDependencyUsagePlugin`   |
|                | , `                          |
|                | FlagIncludedChunksPlugin` ,  |
|                | ` ModuleConcatenationPlugin` |
|                | , ` NoEmitOnErrorsPlugin` ,  |
|                | ` OccurrenceOrderPlugin` , ` |
|                | SideEffectsFlagPlugin` 和 `  |
|                | UglifyJsPlugin` 。           |
+----------------+------------------------------+

综合来说就是会默认启用对当前环境设置的默认插件，方便开发或者有利于打包输出。 可以直接在webpack.config.js里配置，也可以在命令里直接添加参数

` module.exports = { mode: 'production' }; 复制代码` ` webpack --mode=production 复制代码`

### 4.loaders ###

` webpack` 开箱即用的只有 ` js` 和 ` json` 两种文件类型，想要支持其他文件类型的源代码转换就需要 ` loader` 。比如 ` es6` 转换需要 ` babel-loader` , ` css` 转换需要 ` css-loader` , ` typeScript` 转换需要 ` ts-loader` 。

` loader` 本身是一个函数，接收源文件作为参数，返回转换的结果。

推荐配置，在 ` webpack.config.js` 文件中指定 ` loader` ：

` module: { rules: [ { test : /\.css$/, use: [ { loader: 'style-loader' }, { loader: 'css-loader' , options: { modules: true } } ] } ] } 复制代码`

注意：多个loader的情况下，是按照从右到左的顺序执行，要注意书写顺序。

### 5.plugins ###

插件是 ` webpack` 的支柱功能，是对webpack功能的增强。 可以做打包文件的优化压缩，资源的管理，环境变量注入等 ` loader` 无法实现的事情。

> 
> 
> 
> webpack 插件是一个具有 apply 属性的 JavaScript 对象。apply 属性会被 webpack compiler 调用，并且
> compiler 对象可在整个编译生命周期访问。
> 
> 

换句话说， ` plugins` 可以作用于整个构建过程。

由于插件可以携带参数/选项，你必须在 webpack 配置中，向 plugins 属性传入 new 实例。

**webpack.config.js:**

` const HtmlWebpackPlugin = require( 'html-webpack-plugin' ); //通过 npm 安装 const webpack = require( 'webpack' ); //访问内置的插件 const path = require( 'path' ); const config = { entry: './path/to/my/entry/file.js' , output: { filename: 'my-first-webpack.bundle.js' , path: path.resolve(__dirname, 'dist' ) }, module: { rules: [ { test : /\.(js|jsx)$/, use: 'babel-loader' } ] }, plugins: [ new webpack.optimize.UglifyJsPlugin(), new HtmlWebpackPlugin({template: './src/index.html' }) ] }; module.exports = config; 复制代码`