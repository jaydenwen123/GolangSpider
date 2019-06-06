# 从零实现一个 Webpack Loader #

> 
> 
> 
> 参考：
> 
> 
> 
> * Webpack Book --- Extending with Loaders。
> * Webpack Doc --- Loader Interface
> 
> 
> 

Loader 是 Webpack 几大重要的模块之一。当你需要加载资源，就需要设置对应的 Loader，这样就可以对其源代码进行转换。

由于 Webpack 社区的繁荣，使得大部分的业务场景所使用的资源都有对用的 loader，可以参考官网的 [available loaders]( https://link.juejin.im?target=https%3A%2F%2Fwebpack.js.org%2Floaders%2F ) ，但是由于业务的独特性，也可能没有适用的 loader。

接下来会通过几个示例来让你学会如何开发一个自己的 loader。但在此之前，最好先了解如何单独调试它们。

### 利用 loader-runner 调试 Loaders ###

[loader-runner]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Floader-runner ) 允许你不依靠 webpack 单独运行 loader，首先安装它

` mkdir loader-runner-example npm init npm install loader-runner --save-dev 复制代码`

接下来，创建一个 demo-loader，来进行测试

` mkdir loaders echo "module.exports = input => input + input;" > loaders/demo-loader.js 复制代码`

这个 loader 会将引入模块的内容复制一次并返回。创建所引入的模块

` echo "Hello world" > demo.txt 复制代码`

接下来，通过loader-runner运行加载器：

` // 创建 run-loader.js const fs = require ( "fs" ); const path = require ( "path" ); const { runLoaders } = require ( "loader-runner" ); runLoaders( { resource : "./demo.txt" , loaders : [path.resolve(__dirname, "./loaders/demo-loader" )], readResource : fs.readFile.bind(fs), }, (err, result) => (err ? console.error(err) : console.log(result)) ); 复制代码`

当你运行 ` node run-loader.js` ，会看到终端上 log 出来

` { result : [ 'Hello world\nHello world\n' ], resourceBuffer : < Buffer 48 65 6c 6c 6f 20 77 6f 72 6c 64 0a > , cacheable: true, fileDependencies: [ './demo.txt' ], contextDependencies: [] } 复制代码`

从输出结果中可以看出

* result：loader 完成了，我们赋予它的任务，将目标模块的内容复制了一边；
* resourceBuffer：模块内容被转换为了 Buffer。

如果，需要将转换后的文件输出出来，只需要修改 runLoaders 的第二个参数，如

` runLoaders( { resource : "./demo.txt" , loaders : [path.resolve(__dirname, "./loaders/demo-loader" )], readResource : fs.readFile.bind(fs), }, (err, result) => { if (err) console.error(err) fs.writeFileSync( "./output.txt" , result.result) } ); 复制代码`

### 开发一个异步的 Loader ###

尽管你可以通过上述这种同步式接口（synchronous interface）实现一系列的 loader，但是这种形式并不能适用所有场景，例如将第三方软件包包装为 loader 时就会强制要求你执行此操作。

为了将上述例子调整为异步的形式，我们使用 webpack 提供的 ` this.async()` API。通过调用这个函数可以返回一个遵守 Node 规范的回调函数（error first，result second）。

上述例子可以改写为：

**loaders/demo-loader.js**

` module.exports = function ( input ) { const callback = this.async(); // No callback -> return synchronous results // if (callback) { ... } callback( null , input + input); }; 复制代码`
> 
> 
> 
> 
> webpack 通过 ` this` 进行注入，所以不能使用 （） => {}。
> 
> 

之后运行 ` node run-loader.js` 会在终端上打印出相同的结果。如果你想要在对 loader 执行期间产生的异常进行处理，则可以

` module.exports = function ( input ) { const callback = this.async(); callback( new Error ( "Demo error" )); }; 复制代码`

终端上打印的日志会包含错误：demo error，堆栈跟踪显示错误发生的位置。

### 仅返回输出 ###

loader 也可以用于单独输出代码，可以这样实现

` module.exports = function ( ) { return "foobar" ; }; 复制代码`

为什么要这么做呢？你可以将 webpack 的入口文件传递给 loader。来代替指向预先设定的文件的情况，这样可以动态地生成对应 code 的 loader。

> 
> 
> 
> 如果你想要 return 一个 Buffer 形式的输出，可以设定 module.exports.raw = true，将原有的 string 改为
> buffer。
> 
> 

### 写入文件 ###

有一些 loader，像 file-loader，会生成文件。对此 webpack 提供了一个方法， ` this.emitFile` ，但是 loader-runner 暂时还不支持，所以需要主动实现

` runLoaders( { resource : "./demo.txt" , loaders : [path.resolve(__dirname, "./loaders/demo-loader" )], // 为 this 添加 emitFile method context: { emitFile : () => {}, }, readResource : fs.readFile.bind(fs), }, (err, result) => (err ? console.error(err) : console.log(result)) ); 复制代码`

要实现 file-loader 的基本思想，您必须做两件事：找出文件并返回它的路径。 你可以按如下方式实现：

` const loaderUtils = require ( "loader-utils" ); module.exports = function ( content ) { const url = loaderUtils.interpolateName( this , "[hash].[ext]" , { content, }); this.emitFile(url, content); const path = `__webpack_public_path__ + ${ JSON.stringify(url)} ;` ; return `export default ${path} ` ; }; 复制代码`

Webpack 提供了额外的两个 ` emit` 方法：

* ` this.emitWarning(<string>)`
* ` this.emitError(<string>)`

这些方法都是用来替代控制台。 与 ` this.emitFile` 一样，你必须模拟它们才能使loader-runner工作。

接下来的问题是，如何将文件名传递给 loader。

### 传递配置给 loader ###

为了将所需的配置传递给 loader，我们需要做一些修改

**run-loader.js**

` const fs = require ( "fs" ); const path = require ( "path" ); const { runLoaders } = require ( "loader-runner" ); runLoaders( { resource : "./demo.txt" , loaders : [ { loader : path.resolve(__dirname, "./loaders/demo-loader" ), options : { name : "demo.[ext]" , }, }, ], context : { emitFile : () => {}, }, readResource : fs.readFile.bind(fs), }, (err, result) => (err ? console.error(err) : console.log(result)) ); 复制代码`

可以看到，我们将 loaders 从原有的

` loaders: [path.resolve(__dirname, "./loaders/demo-loader" )] 复制代码`

改为了，从而可以传递 ` options`

` loaders: [ { loader : path.resolve(__dirname, "./loaders/demo-loader" ), options : { name : "demo.[ext]" , }, }, ] 复制代码`

为了能够获取到，我们传递的 options，依然利用 loader-utils 来解析 options。

> 
> 
> 
> 别忘了 npm install loader-utils --save-dev
> 
> 

为了将它与 loader 进行连接

**loaders/demo-loader.js**

` const loaderUtils = require ( "loader-utils" ); module.exports = function ( content ) { // 获取 options const { name } = loaderUtils.getOptions( this ); const url = loaderUtils.interpolateName( this , "[hash].[ext]" , { content, }); const url = loaderUtils.interpolateName( this , name, { content }); ); }; 复制代码`

运行 node run-loader.js，你会发现在终端上打印出了

` { result: [ 'export default __webpack_public_path__ + "f0ef7081e1539ac00ef5b761b4fb01b3.txt";' ], resourceBuffer: <Buffer 48 65 6c 6c 6f 20 77 6f 72 6c 64 0a>, cacheable: true , fileDependencies: [ './demo.txt' ], contextDependencies: [] } 复制代码`

可以看出结果与 loader 应返回的内容一致。 你可以尝试将更多选项传递给 loader 或使用查询参数来查看不同组合会发生什么。

### 连接 webpack 与 自定义 loader ###

为了进行一步地使用 loader，我们需要将它与 webpack 联系起来。在这里，我们采用内联的形式引入自定义 loader

` // webpack.config.js 中引入 resolveLoader: { alias : { "demo-loader" : path.resolve( __dirname, "loaders/demo-loader.js" ), }, }, // 在文件中指定 loader，引入 import "!demo-loader?name=foo!./main.css" 复制代码`

当然你还可以通过规则处理 loader。一旦它足够稳定，就建立一个基于 webpack-defaults 的项目，将逻辑推送到 npm，然后开始将 loader 作为包使用。

> 
> 
> 
> 尽管我们使用 loader-runner 来作为开发、测试 loader 的环境。但是它与 webpack 还是有细微的不同的，所以还需要在
> webpack 上测试一下。
> 
> 

### Pitch Loaders ###

webpack 分为两个阶段来执行 loader：pitching、evaluating。如果你熟悉 web 的事件系统，它与事件的捕获、冒泡很相似。webpack 允许你在 pitching 阶段进行拦截执行。它的顺序是，从左到右pitch，从右到左执行。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1bc7e2b8b7a65?imageView2/0/w/1280/h/960/ignore-error/1)

一个 pitch loader 允许你对请求进行修改，甚至终止它。 例如，创建

**loaders/pitch-loader.js**

` const loaderUtils = require ( "loader-utils" ); module.exports = function ( input ) { const { text } = loaderUtils.getOptions( this ); return input + text; }; module.exports.pitch = function ( remainingReq, precedingReq, input ) { console.log( ` Remaining request: ${remainingReq} Preceding request: ${precedingReq} Input: ${ JSON.stringify(input, null , 2 )} ` ); return "pitched" ; }; 复制代码`

并将其添加到 **run-loader.js** 中，

`... loaders: [ { loader : path.resolve (__dirname, './loaders/demo-loader' ), options : { name : 'demo.[ext]' , }, }, path.resolve(__dirname, "./loaders/pitch-loader" ), ], ... 复制代码`

执行 ` node run-loader.js`

` Remaining request: ./demo.txt Preceding request: .../webpack-demo/loaders/demo-loader?{ "name" : "demo.[ext]" } Input: {} { result : [ 'export default __webpack_public_path__ + "demo.txt";' ], resourceBuffer : null , cacheable : true , fileDependencies : [], contextDependencies : [] } 复制代码`

你会发现 pitch-loader 完成了信息的插入以及执行的拦截。

### 总结 ###

webpack loader 实质上就是在描述一种文件格式如何转换为另一种文件格式。你可以通过研究 API 文档或现有的 loader 来弄清楚如何实现特定的功能。

回顾下：

* ` loader-runner` 是一个非常实用的工具，用来开发、调试 loader；
* webpack loader 是依据输入来生成输出的；
* loader 分为同步、异步两种形式，异步的可以通过 ` this.async` 来编写异步的 loader;
* 可以利用 loader 来为 webpack 动态地生成代码，这种情况下，loader 不必接受输入；
* 使用 ` loader-utils` 能够编译 loader 的配置，还可以通过 ` schema-utils` 进行验证；
* 利用 ` resolveLoader.alias` 来完成局部的自定义 loader 引入，防止影响全局；
* Pitching 阶段允许你对 loader 的输入进行修改或拦截执行顺序。