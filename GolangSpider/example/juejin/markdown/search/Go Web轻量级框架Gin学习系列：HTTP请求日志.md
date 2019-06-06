# Go Web轻量级框架Gin学习系列：HTTP请求日志 #

我们知道，用户向服务器发起的每一次Web请求，都会通过HTTP协议头部或Body携带许多的请求元信息给服务器，如请求的URL地址，请求方法，请求头部和请求IP地址等等诸多原始信息，而在Gin框架中，我们可以使用日志的方式记录和输出这些信息，记录用户的每一次请求行为。

下面是一条Gin框架在控制台中输出的日志：

` [GIN] 2019/05/04 - 22:08:56 | 200 | 5.9997ms | ::1 | GET / test 复制代码`

好了，下面看看要如何输出上面的日志吧！

## 日志中件间 ##

在Gin框架中，要输出用户的http请求日志，最直接简单的方式就是借助日志中间件，下面Gin框架的中间件定义：

` func Logger() HandlerFunc 复制代码`

所以，当我们使用下面的代码创建一个 ` gin.Engine` 时，会在控制台中用户的请求日志：

` router := gin.Default() 复制代码`

而使用下面的代码创建 ` gin.Engine` 时，则不会在控制台输出用户的请求日志：

` router := gin.New() 复制代码`

这是为什么呢？这是由于使用 ` Default()` 函数创建的 ` gin.Engine` 实例默认使用了日志中件间 ` gin.Logger()` ,所以，当我们使用第二种方式创建 ` gin.Engine` 时，可以调用 ` gin.Engine` 中的 ` Use()` 方法调用 ` gin.Logger()` ，如下：

` router := gin.New() router.Use(gin.Logger()) 复制代码`

## 在控制台输出日志 ##

Gin框架请求日志默认是在我们运行程序的控制台中输出，而且输出的日志中有些字体有标颜色，如下图所示：

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8d7ae9c5fd641?imageView2/0/w/1280/h/960/ignore-error/1)

当然，我们可以使用 ` DisableConsoleColor()` 函数禁用控制台日志的颜色输出，代码如下所示

` gin.DisableConsoleColor()//禁用请求日志控制台字体颜色 router := gin.Default() router.GET( "test" ,func(c *gin.Context){ c.JSON(200, "test" ) }) 复制代码`

运行后发出Web请求，在控制台输出日志字体则没有颜色：

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8d7e120164486?imageView2/0/w/1280/h/960/ignore-error/1)

虽然Gin框架默认是开始日志字体颜色的，但可以使用 ` DisableConsoleColor()` 函数来禁用，但当被禁用后，在程序中运行需要重新打开控制台日志的字体颜色输出时，可以使用 ` ForceConsoleColor()` 函数重新开启，使用如下：

` gin.ForceConsoleColor() 复制代码`

## 在文件输出日志 ##

Gin框架的请求日志默认在控制台输出，但更多的时候，尤其上线运行时，我们希望将用户的请求日志保存到日志文件中，以便更好的分析与备份。

#### 1. DefaultWriter ####

在Gin框架中，通过 ` gin.DefaultWriter` 变量可能控制日志的保存方式， ` gin.DefaultWriter` 在Gin框架中的定义如下：

` var DefaultWriter io.Writer = os.Stdout 复制代码`

从上面的定义我们可以看出， ` gin.DefaultWriter` 的类型为 ` io.Writer` ,默认值为 ` os.Stdout` ,即控制台输出，因此我们可以通过修改 ` gin.DefaultWriter` 值来将请求日志保存到日志文件或其他地方(比如数据库)。

` package main import ( "github.com/gin-gonic/gin" "io" "os" ) func main () { gin.DisableConsoleColor()//保存到文件不需要颜色 file, _ := os.Create( "access.log" ) gin.DefaultWriter = file //gin.DefaultWriter = io.MultiWriter(file) 效果是一样的 router := gin.Default() router.GET( "/test" , func(c *gin.Context) { c.String(200, "test" ) }) _ = router.Run( ":8080" ) } 复制代码`

运行后上面的程序，会在程序所在目录创建 ` access.log` 文件，当我们发起Web请求后，请求的日志会保存到 ` access.log` 文件，而不会在控制台输出。

通过下面的代码，也可能让请求日志同行保存到文件和在控制台输出：

` file, _ := os.Create( "access.log" ) gin.DefaultWriter = io.MultiWriter(file,os.Stdout) //同时保存到文件和在控制台中输出 复制代码`

#### 2. LoggerWithWriter ####

另外，我们可以使用 ` gin.LoggerWithWriter` 中间件，其定义如下：

` func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc 复制代码`

示例代码：

` package main import ( "github.com/gin-gonic/gin" "os" ) func main () { gin.DisableConsoleColor() router := gin.New() file, _ := os.Create( "access.log" ) router.Use(gin.LoggerWithWriter(file, "" )) router.GET( "test" , func(c *gin.Context) { c.JSON(200, "test" ) }) _ = router.Run() } 复制代码`

` gin.LoggerWithWriter` 中间件的第二个参数，可以指定哪个请求路径不输出请求日志，例如下面代码， ` /test` 请求不会输出请求日志，而 ` /ping` 请求日志则会输出请求日志。

` router.Use(gin.LoggerWithWriter(file, "/test" ))//指定/ test 请求不输出日志 router.GET( "test" , func(c *gin.Context) { c.JSON(200, "test" ) }) router.GET( "ping" , func(c *gin.Context) { c.JSON(200, "pong" ) }) 复制代码`

## 定制日志格式 ##

#### 1. LogFormatterParams ####

上面的例子，我们都是采用Gin框架默认的日志格式，但默认格式可能并不能满足我们的需求，所以，我们可以使用Gin框架提供的 ` gin.LoggterWithFormatter()` 中间件，定制日志格式， ` gin.LoggterWithFormatter()` 中间件的定义如下：

` func LoggerWithFormatter(f LogFormatter) HandlerFunc 复制代码`

从 ` gin.LoggterWithFormatter()` 中间件的定义可以看到该中间件的接受一个数据类型为 ` LogFormatter` 的参数， ` LogFormatter` 定义如下：

` type LogFormatter func(params LogFormatterParams) string 复制代码`

从 ` LogFormatter` 的定义看到该类型为 ` func(params LogFormatterParams) string` 的函数，其参数是为 ` LogFormatterParams` ,其定义如下：

` type LogFormatterParams struct { Request *http.Request TimeStamp time.Time StatusCode int Latency time.Duration ClientIP string Method string Path string ErrorMessage string BodySize int Keys map[string]interface{} } 复制代码`

定制日志格式示例代码：

` func main () { router := gin.New() router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string { //定制日志格式 return fmt.S printf ( "%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n" , param.ClientIP, param.TimeStamp.Format(time.RFC1123), param.Method, param.Path, param.Request.Proto, param.StatusCode, param.Latency, param.Request.UserAgent(), param.ErrorMessage, ) })) router.Use(gin.Recovery()) router.GET( "/ping" , func(c *gin.Context) { c.String(200, "pong" ) }) _ = router.Run( ":8080" ) } 复制代码`

运行上面的程序后，发起Web请求，控制台会输出以下格式的请求日志：

` ::1 - [Wed, 08 May 2019 21:53:17 CST] "GET /ping HTTP/1.1 200 1.0169ms " Mozilla/5.0 (Windows NT 10.0; W in 64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36 " " 复制代码`

#### 2. LoggerWithConfig ####

在前面的例子中，我们使用 ` gin.Logger()` 开启请求日志、使用 ` gin.LoggerWithWriter` 将日志写到文件中，使用 ` gin.LoggerWithFormatter` 定制日志格式，而实际上，这三个中间件，其底层都是调用 ` gin.LoggerWithConfig` 中间件，也就说，我们使用 ` gin.LoggerWithConfig` 中间件，便可以完成上述中间件所有的功能， ` gin.LoggerWithConfig` 的定义如下：

` func LoggerWithConfig(conf LoggerConfig) HandlerFunc 复制代码`

` gin.LoggerWithConfig` 中间件的参数为 ` LoggerConfig` 结构，该结构体定义如下：

` type LoggerConfig struct { // 设置日志格式 // 可选 默认值为：gin.defaultLogFormatter Formatter LogFormatter // Output用于设置日志将写到哪里去 // 可选. 默认值为：gin.DefaultWriter. Output io.Writer // 可选，SkipPaths切片用于定制哪些请求url不在请求日志中输出. SkipPaths []string } 复制代码`

以下例子演示如何使用 ` gin.LoggerConfig` 达到日志格式、输出日志文件以及忽略某些路径的用法：

` func main () { router := gin.New() file, _ := os.Create( "access.log" ) c := gin.LoggerConfig{ Output:file, SkipPaths:[]string{ "/test" }, Formatter: func(params gin.LogFormatterParams) string { return fmt.S printf ( "%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n" , params.ClientIP, params.TimeStamp.Format(time.RFC1123), params.Method, params.Path, params.Request.Proto, params.StatusCode, params.Latency, params.Request.UserAgent(), params.ErrorMessage, ) }, } router.Use(gin.LoggerWithConfig(c)) router.Use(gin.Recovery()) router.GET( "/ping" , func(c *gin.Context) { c.String(200, "pong" ) }) router.GET( "/test" , func(c *gin.Context) { c.String(200, "test" ) }) _ = router.Run( ":8080" ) } 复制代码`

运行上面的程序后，发起Web请求，控制台会输出以下格式的请求日志：

` ::1 - [Wed, 08 May 2019 22:39:43 CST] "GET /ping HTTP/1.1 200 0s " Mozilla/5.0 (Windows NT 10.0; W in 64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36 " " ::1 - [Wed, 08 May 2019 22:39:46 CST] "GET /ping HTTP/1.1 200 0s " Mozilla/5.0 (Windows NT 10.0; W in 64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36 " " 复制代码`

## 小结 ##

每条HTTP请求日志，都对应一次用户的请求行为，记录每一条用户请求日志，对于我们追踪用户行为，过滤用户非法请求，排查程序运行产生的各种问题至关重要，因此，开发Web应用时一定要记录用户请求行为，并且定时分析过滤。