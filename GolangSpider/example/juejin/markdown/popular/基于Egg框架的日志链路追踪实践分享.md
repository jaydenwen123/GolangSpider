# 基于Egg框架的日志链路追踪实践分享 #

## 快速导航 ##

* ` [Logger-Custom]` [需求背景]( #heading-1 )
* ` [Logger-Custom]` [自定义日志插件开发]( #heading-2 )
* ` [Logger-Custom]` [项目扩展]( #heading-3 )
* ` [Logger-Custom]` [项目应用]( #heading-4 )
* ` [ContextFormatter]` [contextFormatter自定义日志格式]( #heading-5 )
* ` [Logrotator]` [日志切割]( #heading-5 )

## 需求背景 ##

实现全链路日志追踪，便于日志监控、问题排查、接口响应耗时数据统计等，首先 API 接口服务接收到调用方请求，根据调用方传的 traceId，在该次调用链中处理业务时，如需打印日志的，日志信息按照约定的规范进行打印，并记录 traceId，实现日志链路追踪。

* **日志路径约定**

` /var/logs/ ${projectName} /bizLog/ ${projectName} -yyyyMMdd.log 复制代码`

* **日志格式约定**

` 日志时间[]traceId[]服务端IP[]客户端IP[]日志级别[]日志内容 复制代码`

采用 Egg.js 框架 egg-logger 中间件，在实现过程中发现对于按照以上日志格式打印是无法满足需求的（至少目前我还没找到可实现方式），如果要自己实现，可能要自己造轮子了，好在官方的 egg-logger 中间件提供了自定义日志扩展功能，参考 [高级自定义日志]( https://link.juejin.im?target=https%3A%2F%2Feggjs.org%2Fzh-cn%2Fcore%2Flogger.html%23%25E9%25AB%2598%25E7%25BA%25A7%25E8%2587%25AA%25E5%25AE%259A%25E4%25B9%2589%25E6%2597%25A5%25E5%25BF%2597 ) ，本身也提供了日志分割、多进程日志处理等功能。

egg-logger 提供了多种传输通道，我们的需求主要是对请求的业务日志自定义格式存储，主要用到 fileTransport 和 consoleTransport 两个通道，分别打印日志到文件和终端。

## 自定义日志插件开发 ##

基于 egg-logger 定制开发一个插件项目，参考 [插件开发]( https://link.juejin.im?target=https%3A%2F%2Feggjs.org%2Fzh-cn%2Fadvanced%2Fplugin.html ) ，以下以 egg-logger-custom 为项目，展示核心代码编写

* **编写logger.js**

> 
> 
> 
> egg-logger-custom/lib/logger.js
> 
> 

` const moment = require ( 'moment' ); const FileTransport = require ( 'egg-logger' ).FileTransport; const utils = require ( './utils' ); const util = require ( 'util' ); /** * 继承 FileTransport */ class AppTransport extends FileTransport { constructor (options, ctx) { super (options); this.ctx = ctx; // 得到每次请求的上下文 } log(level, args, meta) { // 获取自定义格式消息 const customMsg = this.messageFormat({ level, }); // 针对 Error 消息打印出错误的堆栈 if (args[ 0 ] instanceof Error ) { const err = args[ 0 ] || {}; args[ 0 ] = util.format( '%s: %s\n%s\npid: %s\n' , err.name, err.message, err.stack, process.pid); } else { args[ 0 ] = util.format(customMsg, args[ 0 ]); } // 这个是必须的，否则日志文件不会写入 super.log(level, args, meta); } /** * 自定义消息格式 * 可以根据自己的业务需求自行定义 * @param { String } level */ messageFormat({ level }) { const { ctx } = this ; const params = JSON.stringify( Object.assign({}, ctx.request.query, ctx.body)); return [ moment().format( 'YYYY/MM/DD HH:mm:ss' ), ctx.request.get( 'traceId' ), utils.serviceIPAddress, utils.clientIPAddress(ctx.req), level, ].join(utils.loggerDelimiter) + utils.loggerDelimiter; } } module.exports = AppTransport; 复制代码`

* **工具**

> 
> 
> 
> egg-logger-custom/lib/utils.js
> 
> 

` const interfaces = require ( 'os' ).networkInterfaces(); module.exports = { /** * 日志分隔符 */ loggerDelimiter: '[]' , /** * 获取当前服务器IP */ serviceIPAddress: ( ( ) => { for ( const devName in interfaces) { const iface = interfaces[devName]; for ( let i = 0 ; i < iface.length; i++) { const alias = iface[i]; if (alias.family === 'IPv4' && alias.address !== '127.0.0.1' && !alias.internal) { return alias.address; } } } })(), /** * 获取当前请求客户端IP * 不安全的写法 */ clientIPAddress: req => { const address = req.headers[ 'x-forwarded-for' ] || // 判断是否有反向代理 IP req.connection.remoteAddress || // 判断 connection 的远程 IP req.socket.remoteAddress || // 判断后端的 socket 的 IP req.connection.socket.remoteAddress; return address.replace( /::ffff:/ig , '' ); }, clientIPAddress : ctx => { return ctx.ip; }, } 复制代码`

**注意** ：以上获取当前请求客户端IP的方式，如果你需要对用户的 IP 做限流、防刷限制，请不要使用如上方式，参见 [科普文：如何伪造和获取用户真实 IP ？]( https://link.juejin.im?target=https%3A%2F%2Fwww.yuque.com%2Fegg%2Fnodejs%2Fcoopsc ) ，在 Egg.js 里你也可以通过 ctx.ip 来获取，参考 [前置代理模式]( https://link.juejin.im?target=https%3A%2F%2Feggjs.org%2Fzh-cn%2Ftutorials%2Fproxy.html ) 。

* **初始化 Logger**

` egg-logger-custom/app.js 复制代码` ` const Logger = require ( 'egg-logger' ).Logger; const ConsoleTransport = require ( 'egg-logger' ).ConsoleTransport; const AppTransport = require ( './app/logger' ); module.exports = ( ctx, options ) => { const logger = new Logger(); logger.set( 'file' , new AppTransport({ level : options.fileLoggerLevel || 'INFO' , file : `/var/logs/ ${options.appName} /bizLog/ ${options.appName}.log` , }, ctx)); logger.set( 'console' , new ConsoleTransport({ level : options.consoleLevel || 'INFO' , })); return logger; } 复制代码`

以上对于日志定制格式开发已经好了，如果你有实际业务需要可以根据自己团队的需求，封装为团队内部的一个 npm 中间件来使用。

## 项目扩展 ##

自定义日志中间件封装好之后，在实际项目应用中我们还需要一步操作，Egg 提供了 [框架扩展]( https://link.juejin.im?target=https%3A%2F%2Feggjs.org%2Fzh-cn%2Fbasics%2Fextend.html ) 功能，包含五项：Application、Context、Request、Response、Helper，可以对这几项进行自定义扩展，对于日志因为每次日志记录我们需要记录当前请求携带的 traceId 做一个链路追踪，需要用到 Context（是 Koa 的请求上下文） 扩展项。

新建 ` app/extend/context.js` 文件

` const AppLogger = require ( 'egg-logger-custom' ); // 上面定义的中间件 module.exports = { get logger() { // 名字自定义 也可以是 customLogger return AppLogger( this , { appName : 'test' , // 项目名称 consoleLevel: 'DEBUG' , // 终端日志级别 fileLoggerLevel: 'DEBUG' , // 文件日志级别 }); } } 复制代码`

**建议** ：对于日志级别，可以采用配置中心如 Consul 进行配置，上线时日志级别设置为 INFO，当需要生产问题排查时，可以动态开启 DEBUG 模式。关于 Consul 可以关注我之前写的 [服务注册发现 Consul 系列]( https://link.juejin.im?target=https%3A%2F%2Fwww.nodejs.red%2F%23%2Fmicroservice%2Fconsul )

## 项目应用 ##

错误日志记录，直接会将错误日志完整堆栈信息记录下来，并且输出到 errorLog 中，为了保证异常可追踪，必须保证所有抛出的异常都是 Error 类型，因为只有 Error 类型才会带上堆栈信息，定位到问题。

` const Controller = require ( 'egg' ).Controller; class ExampleController extends Controller { async list() { const { ctx } = this ; ctx.logger.error( new Error ( '程序异常！' )); ctx.logger.debug( '测试' ); ctx.logger.info( '测试' ); } } 复制代码`

最终日志打印格式如下所示：

` 2019/05/30 01:50:21[]d373c38a-344b-4b36-b931-1e8981aef14f[]192.168.1.20[]221.69.245.153[]INFO[]测试 复制代码`

## contextFormatter自定义日志格式 ##

Egg-Logger 最新版本支持通过 contextFormatter 函数自定义日志格式，参见之前 [PR：support contextFormatter #51]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feggjs%2Fegg-logger%2Fpull%2F51 )

应用也很简单，通过配置 contextFormatter 函数即可，以下是简单的应用

` config.logger = { contextFormatter : function ( meta ) { console.log(meta); return [ meta.date, meta.message ].join( '[]' ) }, ... }; 复制代码`

同样的在你的业务里对于需要打印日志的地方，和之前一样

` ctx.logger.info( '这是一个测试数据' ); 复制代码`

输出结果如下所示：

` 2019-06-04 12:20:10,421[]这是一个测试数据 复制代码`

## 日志切割 ##

框架提供了 [egg-logrotator]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feggjs%2Fegg-logrotator ) 中间件，默认切割为按天切割，其它方式可参考官网自行配置。

* **框架默认日志路径**

> 
> 
> 
> egg-logger 模块 lib/egg/config/config.default.js
> 
> 

` config.logger = { dir : path.join(appInfo.root, 'logs' , appInfo.name), ... }; 复制代码`

* **自定义日志目录**

很简单按照我们的需求在项目配置文件重新定义 logger 的 dir 路径

` config.logger = { dir : /var/ logs/test/bizLog/ } 复制代码`

这样是否就可以呢？按照我们上面自定义的日志文件名格式（ ` ${projectName}-yyyyMMdd.log` ），貌似是不行的，在日志分割过程中默认的文件名格式为 `.log.YYYY-MM-DD` ，参考源码

> 
> 
> 
> [github.com/eggjs/egg-l…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feggjs%2Fegg-logrotator%2Fblob%2Fmaster%2Fapp%2Flib%2Fday_rotator.js
> )
> 
> 

` _setFile(srcPath, files) { // don't rotate logPath in filesRotateBySize if ( this.filesRotateBySize.indexOf(srcPath) > -1 ) { return ; } // don't rotate logPath in filesRotateByHour if ( this.filesRotateByHour.indexOf(srcPath) > -1 ) { return ; } if (!files.has(srcPath)) { // allow 2 minutes deviation const targetPath = srcPath + moment() .subtract( 23 , 'hours' ) .subtract( 58 , 'minutes' ) .format( '.YYYY-MM-DD' ); // 日志格式定义 debug( 'set file %s => %s' , srcPath, targetPath); files.set(srcPath, { srcPath, targetPath }); } } 复制代码`

* **日志分割扩展**

中间件 [egg-logrotator]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feggjs%2Fegg-logrotator ) 预留了扩展接口，对于自定义的日志文件名，可以用框架提供的 app.LogRotator 做一个定制。

> 
> 
> 
> app/schedule/custom.js
> 
> 

` const moment = require ( 'moment' ); module.exports = app => { const rotator = getRotator(app); return { schedule : { type : 'worker' , // only one worker run this task cron: '1 0 0 * * *' , // run every day at 00:00 }, async task() { await rotator.rotate(); } }; }; function getRotator ( app ) { class CustomRotator extends app. LogRotator { async getRotateFiles() { const files = new Map (); const srcPath = `/var/logs/test/bizLog/test.log` ; const targetPath = `/var/logs/test/bizLog/test- ${moment().subtract( 1 , 'days' ).format( 'YYYY-MM-DD' )}.log` ; files.set(srcPath, { srcPath, targetPath }); return files; } } return new CustomRotator({ app }); } 复制代码`

经过分割之后文件展示如下：

` $ ls -lh /var/logs/ test /bizLog/ total 188K -rw-r--r-- 1 root root 135K Jun 1 11:00 test -2019-06-01.log -rw-r--r-- 1 root root 912 Jun 2 09:44 test -2019-06-02.log -rw-r--r-- 1 root root 40K Jun 3 11:49 test.log 复制代码`

**扩展** ：基于以上日志格式，可以采用 ELK 做日志搜集、分析、检索。

作者：五月君
链接： [www.imooc.com/article/287…]( https://link.juejin.im?target=https%3A%2F%2Fwww.imooc.com%2Farticle%2F287529 )
来源：慕课网

## 阅读推荐 ##

* 侧重于Nodejs服务端技术栈： [www.nodejs.red]( https://link.juejin.im?target=https%3A%2F%2Fwww.nodejs.red )
* 公众号：Nodejs技术栈