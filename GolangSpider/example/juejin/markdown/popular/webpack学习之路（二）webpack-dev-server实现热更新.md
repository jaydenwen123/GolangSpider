# webpack学习之路（二）webpack-dev-server实现热更新 #

上一章对 ` webpack` 的配置有了简单的认识。

这一章，我需要学习的是 ` webpack` 热更新，因为在开发过程中，不希望当文件更改时，人肉去编译文件，刷新浏览器。

# webpack热更新 #

## webpack-dev-server 自动刷新 ##

> 
> 
> 
> ` webpack-dev-server` 为你提供了一个简单的 ` web` 服务器，并且能够实时重新加载( ` live reloading` )。
> 
> 
> 

实际操作一下。

我们先创建一个项目

` mkdir dev-erver && cd dev-server npm init -y // 快速创建一个项目配置 npm i webpack webpack-dev-server webpack-cli --save-dev mkdir src // 创建资源目录 mkdir dist // 输出目录 touch webpack.dev.js // 因为是在开发环境需要热更新，所以直接创建dev配置文件 复制代码`

先编写一下配置文件，我们就简单地编写多入口配置

` 'use strict' ; const path = require( 'path' ); module.exports = { entry: './src/index.js' , output: { path: path.resolve(__dirname, 'dist' ), filename: 'index.js' }, mode: 'development' , devServer: { contentBase: path.resolve(__dirname, 'dist' ) } }; 复制代码`

然后我们去 ` src` 创建文件，编写内容

**index.js**

` 'use strict' document.write( 'hello world~' ) 复制代码`

准备就绪，我们就可以启动 ` webpack-dev-server` ，在 ` package.json` 里添加一条命令

` "scripts" : { "test" : "echo \"Error: no test specified\" && exit 1" , + "dev" : "webpack-dev-server --config webpack.dev.js --open" }, 复制代码`

运行一下

` npm run dev 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b266a444537f7b?imageView2/0/w/1280/h/960/ignore-error/1) 我们看到文件已经打包完成了，但是在 ` dist` 目录里并没有看到文件，这是因为 ` WDS` 是把编译好的文件放在缓存中，没有磁盘上的IO，但是我们是可以访问到的

` http://localhost:8080 复制代码`

配置告知 ` webpack-dev-server` ，在 ` localhost:8080` 下建立服务，将 ` dist` 目录下的文件，作为可访问文件，所以我们可以直接输入 ` bundle.js` 的地址查看

` http://localhost:8080/index.js 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25bf496c9f44a?imageView2/0/w/1280/h/960/ignore-error/1) 显然我们想看效果而不是打包后的代码，所以我们在 ` dist` 目录里创建一个 ` html` 文件引入即可

**index.html**

` <script src= "./index.js" ></script> 复制代码`

这个时候我们访问

` http://localhost:8080 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c0eddd12737?imageView2/0/w/1280/h/960/ignore-error/1)

内容出来了，我们接下来修改 ` index.js` 文件，来看下是否可以自动刷新

` 'use strict' document.write( 'hello world~byebye world' ) 复制代码`

web 服务器就会自动重新加载编译后的代码

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c2471aa16ed?imageView2/0/w/1280/h/960/ignore-error/1)

这确实是热更新，但是这种是每一次修改会重新刷新整个页面，大家可以打开控制台查看。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2677ade9556f3?imageView2/0/w/1280/h/960/ignore-error/1)

显然这还是不满足不我们的需求。

## webpack-dev-server搭配HotModuleReplacementPlugin 实现热更新 ##

我们需要的是，更新修改的模块，但是不要刷新页面。这个时候就需要用到 **模块热替换** 。

> 
> 
> 
> 模块热替换( ` Hot Module Replacement` 或 ` HMR` )是 ` webpack` 提供的最有用的功能之一。它允许在运行时更新各种模块，而无需进行完全刷新。
> 
> 
> 

### 特性：
###

模块热替换( ` HMR - Hot Module Replacement` )功能会在应用程序运行过程中替换、添加或删除模块，而无需重新加载整个页面。主要是通过以下几种方式，来显著加快开发速度：

* 保留在完全重新加载页面时丢失的应用程序状态。
* 只更新变更内容，以节省宝贵的开发时间。
* 调整样式更加快速 - 几乎相当于在浏览器调试器中更改样式。

### 启用 ###

启用HMR，其实十分简单，修改下 ` webpack-dev-server` 的配置，和使用 ` webpack` 内置的HMR插件即可。

` 'use strict' ; const path = require( 'path' ); + const webpack = require( 'webpack' ); module.exports = { entry: './src/index.js' , output: { path: path.resolve(__dirname, 'dist' ), filename: 'index.js' }, mode: 'development' , devServer: { contentBase: path.resolve(__dirname, 'dist' ), + hot: true }, module: { rules: [ { test : /\.(html)$/, use: { loader: 'html-loader' } } ] }, + plugins: [ + new webpack.HotModuleReplacementPlugin() + ] }; 复制代码`

我们修改一下文件，形成引用关系

**index.js**

` 'use strict' import { test } from './page1.js' document.write( 'hello world~1234' ) test () 复制代码`

**page1.js**

` 'use strict' module.exports = { test : function () { console.log(11123456) } } 复制代码`

在入口页 **index.js** 面再添加一段

` if (module.hot) { module.hot.accept(); } 复制代码`

OK，接下来执行

` npm run dev 复制代码`

然后我们修改page1.js，会发现页面并没有刷新，只是更新了部分文件

![](https://user-gold-cdn.xitu.io/2019/6/5/16b268ebe962a084?imageView2/0/w/1280/h/960/ignore-error/1)

这样我们的热更新就实现了。

## 原理 ##

整个的过程我们可以简化一下， ` Webpack Compile` 打包文件传输给 ` Bundle Server` ， ` Bundle Server` 就是一个服务器，然后会执行这些编译后的文件，让浏览器可以访问到。当文件产生变化时， ` Webpack Compile` 编译之后会通知到 ` HMR Server` ， ` HMR Server` 就会通知浏览器端的 ` HMR Runtime` 做出修改。

` HMR Runtime` 是会被打包到编译后的js文件内，然后和 ` HMR Server` 建立websocket通信关系，这样就可以实时更新修改。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26a9cc74f76af?imageView2/0/w/1280/h/960/ignore-error/1)