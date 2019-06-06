# next + koa + mongodb 从搭建到部署（一） #

## 前言 ##

看了很多写服务端渲染（next.js）的文章，但发现搜索的服务端渲染的案例都是一些比较简单，或者不太系统的例子，如果做一个演示或者demo还是可以的，但是在实际应用上不利于分工协作，部署上线。明明可以做大型项目却不知怎么去应用，所以带着这些痛点和疑惑，决定自己做一套next + koa + mongodb 可以应用到应用级项目的项目框架（此项目还在更新中），项目还在不断优化中，希望大家可以多多指点。

话不多说，上图先开始介绍下项目的规划和想法

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc5e605d32e84?imageView2/0/w/1280/h/960/ignore-error/1)

## 项目介绍 ##

这个是项目的首页，应该大家都能看出来了，是一个模仿掘金的的个人博客项目。包含了基本的登录，注册，写文章，展示文章。。。后续还会继续添加新的功能。目标是发布上线，不断迭代，最终做一个成熟且完整的项目。

## 项目结构 ##

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc7bd73698afc?imageView2/0/w/1280/h/960/ignore-error/1) 跟传统的ssr项目一样，分为View层，和sever层，不熟悉的next.js，和koa.js的可以先去脑补一下 [next文档]( https://link.juejin.im?target=http%3A%2F%2Fnextjs.frontendx.cn ) && [koa文档]( https://link.juejin.im?target=https%3A%2F%2Fkoa.bootcss.com%2F ) ，view层后面会简单说下，咱们重点说一下server层的构建，sever层主要分为apps.js（入口文件 app.js已废弃），controller(接口定义，渲染的页面配置)，middleware（中间件），router（路由），token（登录注册时token定义）

* 入口文件apps.js

` const Koa = require( 'koa' ) const next = require( 'next' ) const koaRoute = require( './router' ) const bodyParser = require( 'koa-bodyparser' ); const middleware = require( './middleware' ) const cors = require( 'koa2-cors' ); const port = parseInt(process.env.PORT, 10) || 3000 const dev = process.env.NODE_ENV !== 'production' const app = next({ dev }) const handle = app.getRequestHandler() app.prepare() .then(() => { const server = new Koa() //注入中间件 middleware(server) server.use(bodyParser()) //注入路由 koaRoute(server,app,handle) server.listen(port, () => { console.log(`> Ready on http://localhost: ${port} `) }) }) 复制代码`

## 1.middleware ##

apps.js入口文件比较简单，因为主要逻辑封装到组件中，先从middleware说起

` const bodyParser = require( 'koa-bodyparser' ); const logger = () => { return async (ctx, next) => { const start = Date.now() bodyParser() await next() const responseTime = (Date.now() - start) console.log(`响应时间为: ${responseTime / 1000} s`) } } module.exports = (app) => { app.use(logger()) } 复制代码`

如果看过koa文档会发现，不管使用路由和插件，都需要new Koa()以后再用use去调用，这里只用到一个响应时间方法，以后可能会用到更多中间间，如果引入一个就要在入口引入一次会比较繁琐，所以封装通用方法，可以继续添加，只需在入口引入一次就可以了。

## 2. controller ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b264480084d55a?imageView2/0/w/1280/h/960/ignore-error/1) controller里面主要是koa里的路由管理，分为接口管理（api）和视图路由管理（view）以及数据库管理（db.js），以及入口文件（index.js）。

api里的getListInfor.js：

` const DB = require( '../db' ) const loginInfor = (app) =>{ return async (ctx, next) => { await DB.insert(ctx.request.body).then(res =>{ ctx.response.body = { infor: 'ok' } }) } } module.exports = loginInfor 复制代码`

view里的home.js

` const home = (app) =>{ return async (ctx, next) => { await app.render(ctx.req, ctx.res, '/home' , ctx.query) ctx.respond = false } } module.exports = home 复制代码`

api.js

` const Monk = require( 'monk' ) const url = 'mongodb://localhost:27017/home' ; // Connection URL const db = Monk(url) const dbName = 'col' const collection = db.get(dbName) module.exports = collection //本地用mongdb搭建的数据库，调用方法用的monk插件，不了解的可以去github搜索 复制代码`

index.js:

` //VIEW const index = require( './view/index' ) const home = require( './view/home' ) const editText = require( './view/editText' ) const essay = require( './view/essay' ) const myself = require( './view/myself' ) //API const getListInfor = require( './api/getListInfor' ) const loginInfor = require( './api/loginInfor' ) const POST = 'post' const GET = 'get' module.exports = { view:{// 不需要请求方式 index, home, editText, essay, myself, }, api:{ getListInfor:{ method:GET, getListInfor }, loginInfor:{ method:POST, loginInfor } } } 复制代码`

## 3.router ##

` const router = require( './node_modules/koa-router' )() const Controller = require( '../controller' ) const koaRoute = (app,handle) =>{ //把view层和api层挂载到router // view const {view,api} = Controller for (item in view){ let _name = null; let _moudle = null if (item == 'index' ){ _name = '/' ; _moudle = view[ 'index' ] } else { _name = '/' + item; _moudle = view[item] } router.get(_name,_moudle(app)) } //api for (item in api){ let _method = api[item].method let _name = '/' + item; let _moudle = api[item][item] router[_method](_name,_moudle(app)) } router.get( '*' , async ctx => { await handle(ctx.req, ctx.res) ctx.respond = false }) return router.routes() //启动路由 } module.exports = (server,app,handle) =>{ server.use(koaRoute(app,handle)) } 复制代码`

这时候可以再看一下apps.js是怎么引入的：

` const app = next({ dev }) const handle = app.getRequestHandler() const koaRoute = require( './router' ) app.prepare() .then(() => { const server = new Koa() //注入路由 koaRoute(server,app,handle) // 相当于koa里面app.use(router.routes()) 启动路由...... 复制代码`

未完待续。。。