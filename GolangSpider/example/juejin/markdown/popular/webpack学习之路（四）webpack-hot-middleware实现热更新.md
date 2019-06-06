# webpack学习之路（四）webpack-hot-middleware实现热更新 #

上一节我学习了 ` webpack-dev-middleware` ，但是单独使用它并没有实现热更新，所以这节我要学习搭配使用 ` webpack-hot-middleware` ，来实现热更新功能。

## 创建项目 ##

我们依然创建一个项目

` mkdir hot-middleware && cd hot-middleware npm init -y mkdir src dist cd src && touch index.js block.js cd../dist && touch index.html cd../ touch server.js webpack.dev.config.js 复制代码`

目录：

` ├── dist │   └── index.html ├── package.json ├── server.js ├── src │   ├── block.js │   └── index.js └── webpack.dev.config.js 复制代码`

下载需要的包：

` npm i webpack webpack-cli webpack-dev-middleware webpack-hot-middleware express --save-dev 复制代码`

## 编写内容 ##

### **/dist/index.html** ###

` <!DOCTYPE html> <html lang= "en" > <head> <meta charset= "UTF-8" > <meta name= "viewport" content= "width=device-width, initial-scale=1.0" > <meta http-equiv= "X-UA-Compatible" content= "ie=edge" > <title>webpack-hot-middleware</title> <!-- 引用打包后js文件 --> <script src= "./index.js" ></script> </head> <body> </body> </html> 复制代码`

### **/src/index.js** ###

` 'use strict' import { test } from './block.js' console.log( 'hello world~' ) test () // 接收热更新输出，只有accept才能被更新 if (module.hot) { module.hot.accept(); } 复制代码`

### **/src/block.js** ###

` 'use strict' module.exports = { test : function () { console.log(12345) } } 复制代码`

### **webpack.dev.config.js** ###

` var webpack = require( 'webpack' ); var path = require( 'path' ) module.exports = { mode: 'development' , // 热更新只在开发模式下有用 entry: [ + 'webpack-hot-middleware/client?path=/__webpack_hmr&timeout=20000' , // 必须这么写，这将连接到服务器，以便在包重新构建时接收通知，然后相应地更新客户端 './src/index.js' ], output: { path: path.resolve(__dirname, 'dist' ), publicPath: '/' , // 服务器脚本会用到 filename: 'index.js' }, plugins: [ + new webpack.HotModuleReplacementPlugin(), // 启动HMR + new webpack.NoEmitOnErrorsPlugin() // 在编译出现错误时，使用 NoEmitOnErrorsPlugin 来跳过输出阶段。这样可以确保输出资源不会包含错误。 ], }; 复制代码`

### **server.js** ###

` const express = require( 'express' ); const webpack = require( 'webpack' ); const webpackDevMiddleware = require( 'webpack-dev-middleware' ); const webpackHotMiddleware = require( 'webpack-hot-middleware' ); const app = express(); const config = require( './webpack.dev.config.js' ); // 引入配置文件 const compiler = webpack(config); // 初始化编译器 // 使用webpack-dev-middleware中间件 app.use(webpackDevMiddleware(compiler, { publicPath: config.output.publicPath })); // 使用webpack-hot-middleware中间件，配置在console台输出日志 + app.use(webpackHotMiddleware(compiler, { + log : console.log, path: '/__webpack_hmr' , heartbeat: 10 * 1000 + })); // 使用静态资源目录，才能访问到/dist/idndex.html app.use(express.static(config.output.path)) // Serve the files on port 3000. app.listen(3000, function () { console.log( 'Example app listening on port 3000!\n' ); }); 复制代码`

## 运行 ##

我们增加一个命令运行看看

**package.json 增加一个命令**

` "scripts" : { "test" : "echo \"Error: no test specified\" && exit 1" , + "server" : "node server.js" }, 复制代码` ` npm run server 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b6fd0f5a779e?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b6f6753c6b8c?imageView2/0/w/1280/h/960/ignore-error/1)

浏览器查看 ` http://localhost:3000/`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7af93648a4c?imageView2/0/w/1280/h/960/ignore-error/1)

效果已经出来，我们发现这次比之前单独使用 ` webpack--middleware` 多了一行提示，因为我们配置了热更新日志输出

` [HMR] connected 复制代码`

` HMR- Hot Module Replacement` 即 **热更新** ，这已经很明白地告诉我们热更新已经连接上了，我们来验证下。

修改 ` /src/block.js`

` 'use strict' module.exports = { test : function () { console.log( 'abcd' ) } } 复制代码`

我们发现，请求只是多出来两行，并没有刷新页面

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b790661c0ccf?imageView2/0/w/1280/h/960/ignore-error/1) 控制台也输出了过程

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7968163af3d?imageView2/0/w/1280/h/960/ignore-error/1)

到此，我们就使用 ` webpack-dev-middleware` 和 ` webpack-hot-middleware` 实现了热更新。

详细配置请参考官方文档 [webpack-hot-middleware]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fwebpack-contrib%2Fwebpack-hot-middleware )

## 为什么有了 ` webpack-dev-server` ，还有有 ` webpack-dev-middleware` 搭配 ` webpack-hot-middleware` 的方式呢？ ##

因为 ` webpack-dev-server` 是封装好的，除了 ` webpack.config` 和命令行参数之外，很难去做定制型开发。而 ` webpack-dev-middleware` 是中间件，可以编写自己的后端服务然后使用它，开发更灵活。

I am moving forward.