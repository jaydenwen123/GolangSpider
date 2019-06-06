# 前后端分离 Vue + NodeJS(Koa) + MongoDB，从产品到开发，全栈实践 #

## 写在前面 ##

闲来无事，试了一下 Koa，第一次搞感觉还不错，这个项目比较基础但还是比较完整了，还是有一定的参考价值

## 项目简介 ##

以下是项目地址，希望给个 ` star` ，鼓励一下：

前端 ` gitHub` 地址 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffrankxjkuang%2Fdaike-client )

后端 ` gitHub` 地址 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffrankxjkuang%2Fdaike-server )

**PS：** 数据库我放在了后端项目的 ` db-daike` 目录下

项目名：《代课》

介绍：大学期间加入了两个比较大的社团，虽然已经毕业多年（这个夏天刚好一年哈哈），社团的群里有很多学弟学妹经常发一些帮忙代课的信息，并且也会附带一些好处等等...都是这么过来的，确实比较了解有的课程老师就爱点名，三次就挂科（我没有说《毛概》），好吧扯远了！大概需求就是这样的...我写这个项目的原因就是为了实现这样一个目的...

效果预览：

![one-c.gif](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87a4e1ac0f?imageslim)

![two-c.gif](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87a4f4f6f1?imageslim)

这个就是我们最终做出来的样子，那么开始干正事之前我们还是先捋一捋技术栈吧，说一说都用到了那些东西

## 技术栈 ##

主要还是分为三大块：前端（Vue） + 后端（NodeJS - Koa）+ 数据库（MongoDB）

### 前端 ###

主要是基于 ` Vue` 全家桶， [Vuex]( https://link.juejin.im?target=https%3A%2F%2Fvuex.vuejs.org%2F ) + [axios]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Faxios%2Faxios )

UI 框架随便选了一个，对！就是这么随意: [Vant]( https://link.juejin.im?target=https%3A%2F%2Fyouzan.github.io%2Fvant%2F%23%2Fzh-CN%2Fintro )

css采用 ` scss`

登录页的logo也是用的一个网站在线制作的，不好意思，网址我给忘了

项目结构（大致）如下：

` ├── axios // 对 axios进行 二次封装 │ └── interface // api 文件目录 ├── src │ ├── router // 路由配置 │ └── views // 路由页面 └── vuex // 全局的状态 └── views // 按路由模块进行状态分组 复制代码`

**Tip：** 结构中很多部分我省略了，一部分是属于 ` vue` 全家桶的就没必要赘述了，另一部分比如请求接口和 ` Vuex` 的模块文件，之后会有讲到

### 后端 ###

后端采用 ` NodeJS` ，框架采用的是 [Koa]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fkoajs%2Fkoa%23readme )

其它就是在项目中使用的第三方库，这里理列一下吧：

` // ... // 在项目中使用的文件中，我都有写这些库的npm或git地址方便学习 { "dependencies" : { // 加密用户密码（数据库没有存明文密码） "bcrypt" : "^3.0.0" , // 解析前段请求参数 "koa-bodyparser" : "^4.2.1" , // 路由 "koa-router" : "^7.4.0" , // 解决跨域 "koa2-cors" : "^2.0.6" , // 操作 mongoDB 数据库 "mongoose" : "^5.2.7" , // 生成唯一 id "uuid" : "^3.3.2" } } // ... 复制代码`

后端的项目结构有必要说一些，这个是我参考一些比较规范的项目自己搞的，也是比较随意了，哈哈（我也是第一次这么搞）：

` ├── app // 对 axios进行 二次封装 │ └── controllers // 控制器文件目录，用来操作数据库 │ │ └── ... // 对应操作的表，这里就省略了 │ ├── middleware// 自定义中间件目录 │ ├── models // 定义的表结构 │ │ └── ... // 对应的表，这里就省略了 │ └── utils // 工具模块目录 │ │ └── ... // 工具模块，这里就省略了 ├── rotes // 路由文件 │ ├── router // 路由配置 │ └── views // 路由页面 └── vuex // 全局的状态 └── views // 按路由模块进行状态分组 ├── app.js // 项目入口文件 └── config.js // 配置文件 复制代码`

### 数据库 ###

依稀还得大学的时候学过 SQL，不行了毕业太久忘了，所以这里使用 [MongoDB]( https://link.juejin.im?target=https%3A%2F%2Fwww.mongodb.com%2F ) ，也不做过多介绍，也没啥好说了，安装好了，增删改查...剩下的就是提高了！这里补充一下我用的可视化工具是 [Robo 3T]( https://link.juejin.im?target=https%3A%2F%2Frobomongo.org%2F )

如果你还不了解 ` MongoDB` 的话，我这里简单写了一下如何 [安装使用 MongoDB]( https://link.juejin.im?target=http%3A%2F%2Fwww.kxjun.top%2F2018%2F07%2F31%2Fmongo-start%2F )

这里还是看一下几张主要的表都长啥样吧：

` const CourseSchema = new Schema({ id : { type : String , unique : true , required : true }, status : { type : String }, publisher : { type : String , required : true }, publisherHeader : { type : String }, publisherName : { type : String }, studentId : { type : String }, schoolId : { required : true , type : String }, school : { type : String }, phone : { type : String }, publishTime : { type : String }, closeTime : { type : String }, remark : { type : String }, receiver : { type : String }, receiverName : { type : String }, province : { type : Number }, college : { type : String }, major : { type : String }, courseName : { type : String }, courseTime : { required : true , type : String }, courseClass : { type : String }, coursePlace : { required : true , type : String }, reward : { type : Number }, hasName : { type : Boolean }, hasStuId : { type : Boolean }, hasPhone : { type : Boolean }, hasReward : { type : Boolean } }, { collection : 'courses' , versionKey : false }); 复制代码`

**Tip:** 算了算了，有点占地方，这里就看一张表吧，其它的在项目的 ` models` 文件目录下有

工具和项目结构咱们都搞完了，就开始写代码吧

## 打通前后端 ##

咋们这个项目采用前后端分离的方式进行开发，为了开发的顺畅进行，我们先调试一下：前端发个请求，后端接收消息，并从数据库中拿到数据响应给前端(我们以post和get方法为例写两个接口)，get请求获取数据，post请求插入一条数据；

想一下这里主要的问题应该就是跨域的问题了！再仔细一想，跨域也不能算什么问题吧...哈哈（强行有问题）废话不多说，开始：

* 连接数据库

启动一个 ` Node` 服务连接数据库，后续的操作都是基于数据库的:

` const Koa = require ( 'koa' ); // 这里是一些常量的配置文件 const config = require ( './config' ); const mongoose = require ( 'mongoose' ); const app = new Koa(); mongoose.connect(config.db, { useNewUrlParser : true }, err => { if (err) { console.error( 'Failed to connect to database' ); } else { console.log( 'Connecting database successfully' ); } }); app.listen(config.port); 复制代码`

**Tip:** 启动后服务 ` node app.js` 之后看到如图所示的打印就说明数据库连接成功了:

![connect.png](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87a4d8f8fc?imageView2/0/w/1280/h/960/ignore-error/1)

* 新建一张表，就叫 ` example` 吧

定义一下表结构，为了演示，我们就定义为只有一个类型为 ` String` 类型的字段：

在后端项目 ` models` 目录下新建一个 ` example.js` 文件来定义表结构；

` const mongoose = require ( 'mongoose' ); // 这里的流程官网上有，讲的很清楚，每一步是干什么的 const Schema = mongoose.Schema; const exampleSchema = new Schema({ msg : { type : String , required : true }, }, { collection : 'example' , // 这里是为了避免新建的表会带上 s 后缀 versionKey: false // 不需要 __v 字段，默认是加上的 }); module.exports = mongoose.model( 'example' , exampleSchema); 复制代码`

这里我们先插入一条数据吧，这里为了方便，我直接使用前面提到的可视化工具 ` Robo 3T` 插入一条 'Hello World' 数据：

![insert.png](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87a4ecdb50?imageView2/0/w/1280/h/960/ignore-error/1)

* 编写对应 ` example` 表的控制器，用来暴露接口

在 ` controllers` 目录下新建一个 ` example_controller.js` ：

` // 引入刚才定义的表 const Example_col = require ( './../models/example' ); // get 请求返回所有数据 const getExample = async (ctx, next) => { const req = ctx.request.body; const examples = await Example_col.find({}, { _id : 0 }); ctx.status = 200 ; ctx.body = { msg : 'get request!!' , data : { data : req, examples, } } } // post 带一个 msg 参数，并插入数据库 const postExample = async (ctx, next) => { const req = ctx.request.body; ctx.status = 200 ; if (!req.msg || typeof req.msg != 'string' ) { ctx.status = 401 ; ctx.body = { msg : 'post request!!' , desc : `parameter error！！msg: ${req.msg} ` , data : req } return ; } const result = await Example_col.create({ msg : req.msg}); ctx.body = { msg : 'post request!!' , desc : 'insert success!' , data : result } } // 暴露出这两个方法，在路由中使用 module.exports = { getExample, postExample } 复制代码` * 编写对应的路由模块

在 ` routes/api` 目录下新建一个 ` example_router.js` 文件，主要的作用就是定义接口的请求路径和方式：

` // 引入路由模块并实例化 const Router = require ( 'koa-router' ); const router = new Router(); // 导如对应的控制器 const example_controller = require ( './../../app/controllers/example_controller' ); // 为控制器的方法定义请求路径和请求方式 router.get( '/example/get' , example_controller.getExample); router.post( '/example/post' , example_controller.postExample); module.exports = router; 复制代码` * 作为中间件在入口文件中使用

这个概念我们就不多说了，这里把 [Koa]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fkoajs%2Fkoa%23readme ) 的地址放在这吧

在入口文件 ` app.js` 中增加两句话：

` const example_router = require( './routes/api/example_router' ); app.use(example_router.routes()).use(example_router.allowedMethods()); 复制代码`

重新启动服务 ` node app.js`

* 前端请求接口

这里就不多说了，有兴趣的可以直接看仓库的代码就好了，大致是两个接口有兴趣的可以看看我前面写的文章 [二次封装axios]( https://juejin.im/post/5ae432aaf265da0b9c1063c8 ) :

` const getExample = params => { return axios({ url : '/example/get' , method : 'get' , params }) } const postExample = data => { return axios({ url : '/example/post' , method : 'post' , data }) } 复制代码`

为了方便展示结果，我直接在 前端的入口文件 ` App.vue` 的生命周期钩子中使用：

` // ... 此处代码省略 mounted() { this.$http.getExample({ name : 'frank' }); } 复制代码`

启动前端项目： ` npm start`

![acao.png](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87a748f393?imageView2/0/w/1280/h/960/ignore-error/1)

哎哟！报错了，不要慌仔细看看原来是跨域了呀，还记得前面说的强行有问题吗？我前面有说到在依赖的三方库里有一个叫 ` koa2-cors` ，是时候该它上场了，我们在 ` app.js` 中作为中间件使用它：

* 后端解决跨域

解决跨域的方式有很多，但这不是我们现在讨论的重点，这里我使用上述的 ` koa2-cors` ，只需要将其作为中间件使用就好了，在 ` app.js` 中添加：

` const cors = require ( 'koa2-cors' ); app.use(cors()); 复制代码`

**PS：** 这里要注意一下，js 是单线程语言，中间件是有执行先后顺序的，所以 ` app.use(cors());` 的使用必须在 ` router` 之前，不然就无法解决跨域的问题哦！

好了重新启动服务： ` node app.js`

![getSuccess.png](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87a7d66b3a?imageView2/0/w/1280/h/960/ignore-error/1)

` examples` 这个字段就是我们刚刚在表里插入的数据， ` get` 请求成功了，再来使用 ` post` 请求向数据库里插入一条数据吧：

我们还是在 ` App.vue` 中写：

` mounted() { // this.$http.getExample({name: 'frank'}); this.$http.postExample({ msg : 'test post request!' }); } 复制代码`

![postSuccess.png](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87c9a89e63?imageView2/0/w/1280/h/960/ignore-error/1)

` post` 请求也成功了，好了现在来看看数据库的 ` example` 表：

![example.png](https://user-gold-cdn.xitu.io/2018/8/25/1656fd87cf18ccf6?imageView2/0/w/1280/h/960/ignore-error/1)

没毛病，目前！我们已经成功实现前后端分离，并且已经打通了前后端的交互，后续的开发无非就是依葫芦画瓢拓展开发了。

## 产品详述 ##

文章开头我做了一下产品简介，正式开发之前就需要了解一下到底要做什么，做成什么样！画个流程图，或者做个 ` PRD` 什么的，当然这些对于我来说就算了，全靠 YY，我就大致说一下吧：

上面的预览图看得出来大致的结构，我们主要分为五个模块（其实应该不算登录就四个）：

### 登录 ###

登录模块兼注册，这里参考一下用户表（后端项目 ` models/user.js` 文件），用户的密码我单独写了一个 ` password` 表，用户id作为关联，密码是经过加密的，未使用明文。用户注册的时候不需要详细信息，但是发布课程就需要用户完善个人信息了（比如学校总得有吧），比如用户可以收藏课程，那么增加一个 ` collections` 字段（类型为数组）用来存放课程的 id

### 代课 ###

代课模块主要就是展示课程的信息，用户点击之后可以查看详细的信息，当然部分信息是发布者希望让你看到的才会展示，该模块包括后续的发布模块，课程模块都依赖于 ` courses` 表，这里也不多赘述了...

### 发布 ###

发布课程需要用户完善必要的信息才能发布，发布课程也需要一些课程相关的必要信息（比如没有时间地点怎么上课？），当然为了一些其他的 PY 交易，发布者可以选择提供一些额外的信息或者备注。

### 课程 ###

课程模块，分为三个 tab，分别为我发布的、我代课的和我收藏的，点击之后会展示详细信息

### 我的 ###

我的模块，主要就是个人信息的展示，支持个人信息的修改，对应的表也就是用户表

之后就是正式开发了，我也就不重复的讲代码了，文章的开头结尾我会放上 ` git` 仓库的地址，希望大家点个 ` start` 这将是我这个单身狗最大的快乐（题外话扯多了）

## 基础数据来源 ##

最大的问题就是数据来源，用户可以选择自己的学校（大学），因此需要全国的大学信息作为数据基础， [全国高校名单]( https://link.juejin.im?target=http%3A%2F%2Fwww.huaue.com%2Fgxmd.htm ) 可以在这个网站查到，但是需要自己做一些调整，而且不够完善，本来打算用爬虫搞数据下来，但是我在网上看到某位大佬有全国高校的数据，所以...感谢大佬，让我节约了很多时间！（实在不好意思，我目前不知道是在哪看到的了，不能贴出大佬的地址，再次表示感谢），好了，基础数据有了剩下就是添砖加瓦，下面还是看一下学校的数据结构（单条数据）：

` { "_id" : ObjectId( "5b7648965675dd2687a5b680" ), "id" : "3500" , "name" : "四川大学" , "website" : "http://www.scu.edu.cn/" , "provinceId" : 23 , "level" : "本科" , "abbreviation" : "scu" , "city" : "成都市" } 复制代码`

此外，还有一些全国城市的信息，省的信息等，这些信息我都和课程、用户做了关联方便后续开发的扩展，比如可以统计展示一些可视化图表等。

## 总结 ##

项目比较简单，但是也算是一次比较全的实践了，所用的框架技术（Vue、Koa、mongodb）等都是目前比较火的，对于初学者还是比较有意义的（本人也是第一次用，这篇文章算是学习笔记了）。

不足之处：在 ` 代课` 模块应该区分一下用户应该默认获取同校的课程，当然可以加上查询其它或所有学校的课程，获取列表应该做分页，前段最好还是下拉刷新...诸如此类的问题就不多说了！如果正在看文章的你有兴趣的话可以在此基础上优化。

此外，后端的日志，参数等统一处理也没有做...诸如此类的问题还有...

最后总结一下，学习才能进步...

希望你不吝赐教，可以的话给颗 ` star` 鼓励一下吧：

前端 ` gitHub` 地址 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffrankxjkuang%2Fdaike-client )

后端 ` gitHub` 地址 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffrankxjkuang%2Fdaike-server )

**PS：** 数据库导出的 json 文件我放在了后端仓库的 ` db-daike` 目录下