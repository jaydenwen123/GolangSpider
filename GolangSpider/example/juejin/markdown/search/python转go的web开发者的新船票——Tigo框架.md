# python转go的web开发者的新船票——Tigo框架 #

# 引言 #

Go语言目前比较火的web框架有gin、mux等。但是gin、mux等的代码风格让很多曾经使用Tornado框架的开发人员感觉不适应。这里给大家带来了一个新的框架——Tigo框架，让Python/Tornado转Go的开发者又多了一条可选择的道路。

首先我们看一下Tornado框架的一个Demo：

` # -*- coding: utf-8 -*- import tornado.ioloop import tornado.web class MainHandler (tornado.web.RequestHandler) : def get (self) : self.write( "Hello, Demo!" ) urls = [ ( r"/" , MainHandler), ] if __name__ == "__main__" : app = tornado.web.Application(urls) app.listen( 8888 ) tornado.ioloop.IOLoop.current().start() 复制代码`

接下来再看一下Tigo的Demo：

` package main import "github.com/karldoenitz/Tigo/TigoWeb" type DemoHandler struct { TigoWeb.BaseHandler } func (demoHandler *DemoHandler) Get () { demoHandler.ResponseAsText( "Hello, Demo!" ) } var urls = []TigoWeb.Router{ { "/demo" , &DemoHandler{}, nil }, } func main () { application := TigoWeb.Application{IPAddress: "0.0.0.0" , Port: 8888 , UrlRouters: urls} application.Run() } 复制代码`

二者的代码风格还是比较相近的。

## Tigo的基本信息及安装 ##

Tigo是一款Go（Golang）语言开发的web应用框架，主要设计灵感来源于Tornado框架，结合了Go本身的一些特性（interface、struct嵌入等）而设计的一个框架。
Tigo框架的首页： [点击此处]( https://link.juejin.im?target=https%3A%2F%2Fkarldoenitz.github.io%2FTigo%2F )
Tigo框架的项目地址： [点击此处]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fkarldoenitz%2FTigo )
Tigo框架的API文档： [点击此处]( https://link.juejin.im?target=https%3A%2F%2Fgodoc.org%2Fgithub.com%2Fkarldoenitz%2FTigo )
安装方法：

` go get github.com/karldoenitz/Tigo/... 复制代码`

Tigo主要包含 ` TigoWeb` 、 ` request` 、 ` binding` 、 ` logger` 四个包，其中 ` TigoWeb` 是Tigo框架的核心包，搭建服务主要依靠此包； ` request` 包是一个httpclient工具包，用来发送http请求； ` binding` 包是用来对json以及form进行校验的工具包； ` logger` 包则是用来记录log的工具包。

# Tigo工具包使用简介 #

## TigoWeb包的使用 ##

首先来看一下Tigo一个服务所包含的基本代码块： ` Application` 、 ` Handler` 、 ` UrlRouter` ;
一个 ` Application` 实例就是一个服务，一个 ` Handler` 就是一个Controller， ` UrlRouter` 将url与handler进行绑定。

### 服务搭建 ###

我们写一个简单的示例：
定义一个handler，在handler中实现一个get方法，用来响应HTTP的get请求方式；

` import "github.com/karldoenitz/Tigo/TigoWeb" type DemoHandler struct { TigoWeb.BaseHandler } func (demoHandler *DemoHandler) Get () { demoHandler.ResponseAsText( "Hello, Demo!" ) } 复制代码`

接下来写一个路由映射；

` var urls = []TigoWeb.Router{ { "/demo" , &DemoHandler{}, nil }, } 复制代码`

最后我们写一个main方法，初始化一个Application实例，并运行该实例；

` func main () { application := TigoWeb.Application{IPAddress: "0.0.0.0" , Port: 8888 , UrlRouters: urls} application.Run() } 复制代码`

### BaseHandler ###

` BaseHandler` 是所有handler的基类，cookie操作、http header操作、http上下文操作等方法都在此结构体中实现，开发者只要继承这个handler，就可以使用这里面的方法。

## binding包的使用 ##

目前binding支持json和form的实例化校验，后续将推出url参数的实例化校验。

### json及form的校验 ###

我们定义一个结构体 ` UserInfo` ，具体如下：

` type Person struct { Name string `json:"name" required:"true"` Age int `json:"age" required:"true" default:"18"` Mobile string `json:"mobile" required:"true" regex:"^1([38][0-9]|14[57]|5[^4])\\d{8}$"` Info string `json:"info" required:"false"` } 复制代码`

当tag中 ` required` 为true的时候，将会对此字段进行校验，否则略过； ` default` 表示默认值，当此字段必须，并且json/form中没有值的时候，就会采用此默认值； ` regex` 则表示正则匹配，如果该字段的值满足正则匹配，则认为该字段合法，否则校验失败。
例如：

` // 我们向服务发送一个json，数据格式如下所示 { "name": "张三", "mobile": "13746588129" } 复制代码`

校验结果如下：

` // 我们将json校验之后的结果转为json打印出来 { "name": "张三", "age": 18, "mobile": "13746588129", "info": "" } 复制代码`

当然， ` TigoWeb.Basehandler` 中封装了 ` CheckJsonBinding` 、 ` CheckFormBinding` 、 ` CheckParamBinding` 三个内置方法进行json或form的校验。

## request包的使用 ##

` request` 包是用来发送http请求的工具包，

### 发送http请求 ###

使用 ` request` 包发送http请求非常简单，如果你使用过Python的requests模块，那么Tigo的 ` request` 包极易上手。
示例1：

` // 发送Post请求示例 import "github.com/karldoenitz/Tigo/request" func main () { headers := map [ string ] string { "Content-Type" : "application/x-www-form-urlencoded" , } postData := map [ string ] interface {}{ "chlid" : "news_news_bj" , } response, err := request.Post( "https://test.hosts.com/api/get_info_list?cachedCount=0" , postData, headers) if err != nil { fmt.Println(err.Error()) } contentStr := response.ToContentStr() fmt.Println(contentStr) } 复制代码`

示例2：

` // 发送get请求 import "github.com/karldoenitz/Tigo/request" func main () { response, err := request.Get( "https://demo.host.com/api/detail?id=773947310848622080" ) if err != nil { fmt.Println(err.Error()) } contentStr := response.ToContentStr() result := struct { Code int `json:"code"` Msg string `json:"msg"` }{} json.Unmarshal(response.Content, &result) fmt.Println(result.Code) fmt.Println(result.Msg) fmt.Println(contentStr) } 复制代码`

## logger包的使用 ##

logger的模块的使用非常简单，可以通过json文件或者yaml文件配置，示例如下：

` { "cookie": "TencentCode", "ip": "0.0.0.0", "port": 8080, "log": { "trace": "/Users/karllee/Desktop/trace.log", // 此文件存储trace跟踪日志 "info": "/Users/karllee/Desktop/run-info.log", // 此文件存储info日志 "warning": "/Users/karllee/Desktop/run.log", // 将warning和error日志都写入run.log "error": "/Users/karllee/Desktop/run.log", "time_roll": "H*2" // 每2小时切分一次log日志 } } 复制代码`

配置文件写好之后，只需要在初始化 ` Application` 实例的时候指定配置文件地址即可，例如：

` application := TigoWeb.Application{ ... ConfigPath: "./configuration.json" , // 配置文件的绝对路径或相对路径 } 复制代码`

# 结语 #

Tigo更详细的功能及文档请查看 [github项目主页]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fkarldoenitz%2FTigo ) 。
Demo可以 [点击此处]( https://link.juejin.im?target=https%3A%2F%2Fkarldoenitz.github.io%2FTigoOld%2Fsource%2FDemos.zip ) 下载。