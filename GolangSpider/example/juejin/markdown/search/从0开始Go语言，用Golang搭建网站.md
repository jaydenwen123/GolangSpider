# 从0开始Go语言，用Golang搭建网站 #

# 实践是最好的学习方式 #

![](https://user-gold-cdn.xitu.io/2019/5/12/16aaa3a3f5bced85?imageView2/0/w/1280/h/960/ignore-error/1)

零基础通过开发Web服务学习Go语言

本文适合有一定编程基础，但是没有Go语言基础的同学。

**也就是俗称的“骗你”学Go语言系列。**

这是一个适合阅读的系列，我希望您能够在车上、厕所、餐厅都阅读它，涉及代码的部分也是精简而实用的。

## 学习需要动机 ##

Go语言能干什么？为什么要学习Go语言？

本系列文章，将会以编程开发中需求最大、应用最广的Web开发为例，一步一步的学习Go语言。当看完本系列，您能够清晰的了解Go语言Web开发的基本原理，您会惊叹于Go语言的简洁、高效和新鲜。

## 结果反馈才能让你记住 ##

《刻意练习》一书中说，学习需要及时反馈结果，才能提高学习体验。

本系列文章的每一节，都会包含一段可运行的有效代码，跟着内容一步一步操作，你可以在你自己的计算机上体验每一句代码的作用。

## 不要学习不需要的东西 ##

文章围绕范例为核心，介绍知识点。文中不罗列语法和关键字，当您还不知道它们用来干什么时，反而会干扰您的注意力。

希望您在阅读本系列文章后，对Go语言产生更多的学习欲望，成为一名合格的Gopher

> 
> 
> 
> Gopher:原译是囊地鼠，也就是Go语言Logo的那个小可爱；这里特指Go程序员给自己的昵称。
> 
> 

# 如何10分钟搭建Go开发环境 #

## 1.下载Go语言安装文件 ##

**访问Go语言官方网站下载页面:**

[golang.org/dl]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdl )

可以看到官网提供了Microsoft Windows、Apple MacOS、Linux和Source下载。

直接下载对应操作系统的安装包。

![](https://user-gold-cdn.xitu.io/2019/5/13/16aae7fb9941476e?imageView2/0/w/1280/h/960/ignore-error/1)

## 2.和其他软件一样，根据提示安装 ##

## 3.配置环境变量 ##

**在正式使用Go编写代码之前，还有一个重要的“环境变量”需要配置：“$GOPATH”**

> 
> 
> 
> GOPATH环境变量指定工作区的位置。如果没有设置GOPATH，则假定在Unix系统上为 ` $HOME/go` ，在Windows上为 `
> %USERPROFILE%\go` 。如果要将自定义位置用作工作空间，可以设置GOPATH环境变量。
> 
> 

GOPATH环境变量是用于设置Go编译可以执行文件、包源码以及依赖包所必要的工作目录路径，Go1.11后，新的木块管理虽然可以不再依赖 ` $GOPATH/src` ，但是依然需要使用 ` $GOPATH/pkg` 路径来保存依赖包。

首先，创建好一个目录用作GOPATH目录

然后设置环境变量 ` GOPATH` :

Linux & MacOS:

导入环境变量

` $ export GOPATH=$YOUR_PATH/go`

保存环境变量

` $ source ~/.bash_profile`

Windows:

` 控制面板->系统->高级系统设置->高级->环境变量设置`

![GOPATH设置好后，它是一个空目录，当在开发工作中执行go get、go install命令后，](https://juejin.im/equation?tex=GOPATH%E8%AE%BE%E7%BD%AE%E5%A5%BD%E5%90%8E%EF%BC%8C%E5%AE%83%E6%98%AF%E4%B8%80%E4%B8%AA%E7%A9%BA%E7%9B%AE%E5%BD%95%EF%BC%8C%E5%BD%93%E5%9C%A8%E5%BC%80%E5%8F%91%E5%B7%A5%E4%BD%9C%E4%B8%AD%E6%89%A7%E8%A1%8Cgo%20get%E3%80%81go%20install%E5%91%BD%E4%BB%A4%E5%90%8E%EF%BC%8C) GOPATH所指定的目录会生成3个子目录：

* bin：存放 ` go install` 编译的可执行二进制文件
* pkg：存放 ` go install` 编译后的包文件，就会存放在这里
* src：存放 ` go get` 命令下载的源码包文件

## 4.检查环境 ##

打开命令行工具，运行

` $ go env`

如果你看到类似这样的结果，说明Go语言环境安装完成.

` GOARCH="amd64" GOBIN="" GOCACHE="/Users/zeta/Library/Caches/go-build" GOEXE="" GOFLAGS="" GOHOSTARCH="amd64" GOHOSTOS="darwin" GOOS="darwin" GOPATH="/Users/zeta/workspace/go" GOPROXY="https://goproxy.io" GORACE="" GOROOT="/usr/local/go" GOTMPDIR="" GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64" GCCGO="gccgo" CC="clang" CXX="clang++" CGO_ENABLED="1" GOMOD="" CGO_CFLAGS="-g -O2" CGO_CPPFLAGS="" CGO_CXXFLAGS="-g -O2" CGO_FFLAGS="-g -O2" CGO_LDFLAGS="-g -O2" PKG_CONFIG="pkg-config" GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/7v/omg2000000000000019/T/go-build760324613=/tmp/go-build -gno-record-gcc-switches -fno-common" 复制代码`

## 5.选择一款趁手的编辑器或IDE ##

现在很多 **通用的编辑器或IDE** 都支持Go语言比如

* [Atom]( https://link.juejin.im?target=https%3A%2F%2Fatom.io%2F )
* [Visual Studio Code]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2F )
* [Sublime Text2]( https://link.juejin.im?target=https%3A%2F%2Fwww.sublimetext.com%2F2 )
* [ItelliJ Idea]( https://link.juejin.im?target=https%3A%2F%2Fwww.jetbrains.com%2Fidea%2F )

Go语言 **专用的IDE有**

* [LiteIDE]( https://link.juejin.im?target=http%3A%2F%2Fliteide.org%2Fcn%2F )
* [Goland]( https://link.juejin.im?target=https%3A%2F%2Fwww.jetbrains.com%2Fgo%2F%3FfromMenu )

专用的IDE无论是配置和使用都比通用编辑器/IDE的简单许多，但是我还是推荐大家使用通用编辑器/IDE，因为在开发过程中肯定会需要编写一些其他语言的程序或脚本，专用IDE在其他语言编写方面较弱，来回切换不同的编辑器/IDE窗口会很低效。

另外，专用IDE提供很多高效的工具，在编译、调试方面都很方便，但是学习阶段，建议大家手动执行命令编译、调试，有利于掌握Go语言。

# 四行代码的Hello World！所能表达出来的核心 #

**命令行代码仅适用于Linux和MacOS系统，Windows根据说明在视窗下操作即可。**

## 1.创建项目 ##

创建一个文件夹，进入该文件夹

` $ mkdir gowebserver && cd gowebserver`

新建一个文件 main.go

` $ touch main.go`

## 2. 用编辑器打开文件，并输入以下代码： ##

` package main import "fmt" func main () { fmt.Println( "Hello, 世界" ) } 复制代码`

## 3.打开命令行终端，输入以下命令 ##

` $ go run main.go`

看到终端会输出：

` Hello, 世界`

**第一个Go代码就完成了**

这是一个很简单的Hello World，但是包含了Go语言编程的许多核心元素，接下来就详细讲解。

## 解读知识点: 包 与 函数 ##

### ` package` 申明包 & ` import` 导入包 ###

Go程序是由包构成的。

代码的第一行, 申明程序自己的包，用 ` package` 关键字。 ` package` 关键字必须是第一行出现的代码。

范例代码中，申明的本包名 ` main`

在代码中第二行, 导入“fmt”包, 使用 ` import` 关键字。默认情况下，导入包的包名与导入路径的最后一个元素一致，例如 ` import "math/rand"` ，在代码中使用这个包时，直接使用 ` rand` ，例如 ` rand.New()`

导入包的写法可以多行,也可以“分组”, 例如:

` import "fmt" import "math/rand" 复制代码`

或者 分组

` import ( "fmt" "math/rand" ) 复制代码`
> 
> 
> 
> fmt包是Go语言内建的包,作用是输出打印。
> 
> 

## ` func` 关键字：定义函数 ##

` func` 是function的缩写, 在Go语言中是定义函数的关键字。

func定义函数的格式为：

` func 函数名 (参数1 类型,参数2 类型) { 函数体 } 复制代码`

本例中定义了一个main函数。 ` main` 函数没有参数。 然后在 ` main` 函数体里调用 ` fmt` 包的 ` Println` 函数，在控制台输出字符串 “Hello, 世界”

**所有Go语言的程序的入口都是main包下的main函数** ` main.main()` ，所以每一个可执行的Go程序都应该有一个 ` main` 包和一个 ` main函数` 。

**我们已经介绍了九牛一毛中的一毛，接下来正式通过搭建一个简单的Web服务学习Go语言**

# 0依赖，创建一个Web服务 #

## 先从代码开始 ##

打开之前创建好的 ` main.go` 文件，修改代码如下:

` package main import ( "fmt" "net/http" ) func myWeb (w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "这是一个开始" ) } func main () { http.HandleFunc( "/" , myWeb) fmt.Println( "服务器即将开启，访问地址 http://localhost:8080" ) err := http.ListenAndServe( ":8080" , nil ) if err != nil { fmt.Println( "服务器开启错误: " , err) } } 复制代码`

保存文件，然后在命令行工具下输入命令，运行程序

` $ go run main.go`

这时候，你会看到用 ` fmt.Println` 打印出来的提示，在浏览器中访问 ` http://localhost:8080` 你将访问到一个页面，显示 " **这是一个开始** "

## 解读 ##

我们从程序运行的顺序去了解它的工作流程

首先，定义 ` package main` ，然后导入包。

这里，导入了一个新的包 ` net/http` ，这个包是官方的，实现http客户端和服务端的各种功能。Go语言开发Web服务的所有功能就是基于这个包（其他第三方的Go语言Web框架也都基于这个包，没有例外）

### 先看 ` main` 函数里发生了什么 ###

第一句，匹配路由和处理函数

` http.HandleFunc("/", myWeb)`

调用http包的HandleFunc方法，匹配一个路由到一个处理函数 ` myWeb` 。

这句代码的意思是，当通过访问地址 [http://localhost/]( https://link.juejin.im?target=http%3A%2F%2Flocalhost%2F ) 时，就等同于调用了 myWeb 函数。

第二句，用fmt在控制台打印一句话，纯属提示。

第三句，开启服务并且监听端口

` err := http.ListenAndServe( ":8080" , nil ) 复制代码`

在这句，调用了 ` http` 包中的 ` ListenAndServe` 函数，该函数有两个参数，第一个是指定监听的端口号，第二个是指定处理请求的handler，通常这个参数填nil，表示使用默认的ServeMux作为handler。

**什么是nil？**

` nil` 就是其他语言里的 ` null` 。

> 
> 
> 
> 什么是handler?什么是ServeMux？
> ServeMux就是一个HTTP请求多路由复用器。它将每个传入请求的URL与已注册模式的列表进行匹配，并调用与URL最匹配的模式的处理程序。
> 很熟悉吧？还记得前面的 ` http.HandleFunc` 吗？他就是给http包中默认的ServeMux（DefaultServeMux）添加URL与处理函数匹配。
> 通常都是使用http包中的默认ServeMux，所以在 ` http.ListenAndServe` 函数的第二个参数提供nil就可以了
> 
> 

` ListenAndServe` 函数会一直监听，除非强制退出或者出现错误。

如果这句开启监听出现错误，函数会退出监听并会返回一个error类型的对象，因此用 ` err` 变量接收返回对象。紧接着，判断 ` err` 是否为空，打印出错误内容，程序结束。

这里有两个Go语言 **知识点**

#### 1.定义变量 ####

Go语言是静态语言，需要定义变量，定义变量用关键字 ` var`

` var str string = "my string" //^ ^ ^ //关键字 变量名 类型 复制代码`

Go还提了一种简单的变量定义方式 ` :=` ， **自动根据赋值的对象定义变量类型** ，用起来很像脚本语言：

` str := "my string" 复制代码`

#### 2.错误处理 ####

` if err != nil { //处理.... } 复制代码`

在Go语言中，这是很常见的错误处理操作，另一种panic异常，官方建议不要使用或尽量少用，暂不做介绍，先从err开始。

Go语言中规定，如果函数可能出现错误，应该返回一个error对象，这个对象至少包含一个Error()方法错误信息。

因此，在Go中，是看不到try/catch语句的，函数使用error传递错误，用if语句判断错误对象并且处理错误。

#### 3. if 语句 ####

与大多数语言使用方式一样，唯一的区别是，表达式不需要()包起来。

另外，Go语言中的if可以嵌入一个表达式，用;号隔开，例如范例中的代码可以改为：

` if err := http.ListenAndServe( ":8080" , nil ); err != nil { fmt.Println( "服务器开启错误: " , err) } 复制代码`

err这个变量的生命周期只在 ` if` 块中有效。

### 请求处理 myWeb函数 ###

在 ` main` 函数中，用 ` http.HandleFunc` 将 myWeb与路由 ` /` 匹配在一起。

` HandleFunc` 函数定义了两个参数 ` w` , ` r` ，参数类型分别是 ` http.ResponseWriter` 和 ` *http.Request` ， ` w` 是响应留写入器， ` r` 是请求对象的指针。

**响应流写入器 w** : 用来写入http响应数据

**请求对象 * r** : 包含了http请求所有信息，注意，这里使用了指针，在定义参数时用 ` *` 标记类型，说明这个参数需要的是这个类型的对象的 **指针** 。

当有请求路径 ` /` ，请求对象和响应流写入器被传递给 ` myWeb` 函数，并由 ` myWeb` 函数负责处理这次请求。

**Go语言中红的指针** ： 在Go语言中 除了map、slice、chan 其他函数传参都是 **值传递** ，所以，如果需要达到引用传递的效果，通过传递对象的指针实现。在Go语言中，取对象的指针用 ` &` ，取值用 ` *` ，例如：

` mystring := "hi" //取指针 mypointer := &mystring //取值 mystring2 := *mypointer fmt.Println(mystring,mypointer,mystring2) 复制代码`

把这些代码放在 ` main` 函数里， ` $ go run main.go` 运行看看

### myWeb函数体 ###

` fmt.Fprintf(w, "这是一个开始" ) 复制代码`

再一次遇到老熟人 ` fmt` ，这次使用他的 ` Fprintf` 函数将字符串“这是一个开始”,写入到 ` w` 响应流写入器对象。 ` w` 响应流写入器里写入的内容最后会被Response输出到用户浏览器的页面上。

### 总结一下，从编码到运行，你和它都干了些什么： ###

* 定义一个函数myWeb，接收参数 响应流写入器和请求对象两个参数
* 在main函数中，在默认的ServeMux中将路由/与myWeb绑定
* 运行默认的ServeMux监听本地8080端口
* 访问本地8080端口 ` /` 路由
* http将请求对象和响应写入器都传递给myWeb处理
* myWeb向响应流中写入一句话，结束这次请求。

虽然代码很少很少，但是这就是一个最基本的Go语言Web服务程序了。

# Web互动第一步，Go http 获得请求参数 #

## 还是先从代码开始 ##

打开 ` main.go` 文件，修改 ` myWeb` 函数，如下:

` func myWeb (w http.ResponseWriter, r *http.Request) { r.ParseForm() //它还将请求主体解析为表单，获得POST Form表单数据，必须先调用这个函数 for k, v := range r.URL.Query() { fmt.Println( "key:" , k, ", value:" , v[ 0 ]) } for k, v := range r.PostForm { fmt.Println( "key:" , k, ", value:" , v[ 0 ]) } fmt.Fprintln(w, "这是一个开始" ) } 复制代码`

运行程序

` $ go run main.go`

然后用任何工具（推荐Postman）提交一个POST请求，并且带上URL参数，或者在命令行中用cURL提交

` curl --request POST \ --url 'http://localhost:8080/?name=zeta' \ --header 'cache-control: no-cache' \ --header 'content-type: application/x-www-form-urlencoded' \ --data description=hello 复制代码`

页面和终端命令行工具会答应出以下内容：

` key: name , value: zeta key: description , value: hello 复制代码`

## 解读 ##

` http` 请求的所有内容，都保存在 ` http.Request` 对象中，也就是 ` myWeb` 获得的参数 ` r` 。

首先，调用 ` r.ParseForm()` ，作用是填充数据到 ` r.Form` 和 ` r.PostForm`

接下来，分别循环获取遍历打印出 ` r.URL.Query()` 函数返回的值 和 ` r.PostForm` 值里的每一个参数。

` r.URL.Query()` 和 ` r.PostForm` 分别是URL参数对象和表单参数对象 ，它们都是键值对值，键的类型是字符串 ` string` ，值的类型是 ` string` 数组。

> 
> 
> 
> 在http协议中，无论URL和表单，相同名称的参数会组成数组。
> 
> 

**循环遍历：for...range**

Go语言的循环只有 ` for` 关键字，以下是Go中4种 ` for` 循环

` //无限循环，阻塞线程，用不停息，慎用！ for { } //条件循环，如果a<b，循环，否则，退出循环 for a < b{ } //表达式循环，设i为0，i小于10时循环，每轮循环后i增加1 for i:= 0 ; i< 10 ; i++{ } //for...range 遍历objs，objs必须是map、slice、chan类型 for k, v := range objs{ } 复制代码`

前3种，循环你可以看作条件循环的变体（无限循环就是无条件的循环）。

本例种用到的是 ` for...range` 循环，遍历可遍历对象，并且每轮循环都会将键和值分别赋值给变量 ` k` 和 ` v`

我们页面还是只是输出一句“ **这是一个开始** ”。我们需要一个可以见人的页面，这样可以不行

你也许也想到了，是不是可以在输出时，硬编码HTML字符串？当然可以，但是Go http包提供了更好的方式，HTML模版。

**接下来，我们就用HTML模版做一个真正的页面出来**

# 动态响应数据给访客，Go http HTML模版+数据绑定 #

读取HTML模版文件，用数据替换掉对应的标签，生成完整的HTML字符串，响应给浏览器，这是所有Web开发框架的常规操作。Go也是这么干的。

Go html包提供了这样的功能：

" ` html/template` "

## 从代码开始 ##

` main` 函数不变，增加导入 ` html/template` 包，然后修改 ` myWeb` 函数，如下：

` import ( "fmt" "net/http" "text/template" //导入模版包 ) func myWeb (w http.ResponseWriter, r *http.Request) { t := template.New( "index" ) t.Parse( "<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>" ) data := map [ string ] string { "name" : "zeta" , "someStr" : "这是一个开始" , } t.Execute(w, data) // fmt.Fprintln(w, "这是一个开始") } 复制代码`

在命令行中运行 ` $ go run main.go` ，访问 ` http://localhost:8080`

看， ` <div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>` 中的 ` {{.name}}` 和 ` {{.someStr}}` 被替换成了 ` zeta` 和 ` 这是一个开始` 。并且，不再使用 ` fmt.Fprintln` 函数输出数据到Response了

但是...这还是在代码里硬编码HTML字符串啊...

别着急，template包可以解析文件，继续修改代码：

* 根目录下创建一个子目录存放模版文件 templates, 然后进入目录创建一个文件 ` index.html` ，并写入一些HTML代码 (我不是个好前端)
` < html > < head > </ head > < body > < div > Hello {{.name}} </ div > < div > {{.someStr}} </ div > </ body > </ html > 复制代码` * 修改 ` myWeb` 函数
` func myWeb (w http.ResponseWriter, r *http.Request) { //t := template.New("index") //t.Parse("<div>Hi,{{.name}},{{.someStr}}<div>") //将上两句注释掉，用下面一句 t, _ := template.ParseFiles( "./templates/index.html" ) data := map [ string ] string { "name" : "zeta" , "someStr" : "这是一个开始" , } t.Execute(w, data) // fmt.Fprintln(w, "这是一个开始") } 复制代码`

在运行一下看看，页面按照HTML文件的内容输出了，并且{{.name}}和{{.someStr}}也替换了，对吧？

## 解读 ##

可以看到， ` template` 包的核心功能就是将HTML字符串解析暂存起来，然后调用 ` Execute` 的时候，用数据替换掉HTML字符串中的 ` {{}}` 里面的内容

在第一个方式中 ` t:=template.New("index")` 初始化一个template对象变量，然后用调用 ` t.Parse` 函数解析字符串模版。

然后，创建一个map对象，渲染的时候会用到。

最后，调用 ` t.Execute` 函数，不仅用数据渲染模版，还替代了 ` fmt.Fprintln` 函数的工作，将输出到Response数据流写入器中。

第二个方式中，直接调用 ` template` 包的 ` ParseFiles` 函数，直接解析相对路径下的index.html文件并创建对象变量。

### 知识点 ###

本节出现了两个新东西 ` map` 类型 和 赋值给“ ` _` ”

#### map类型 ####

**map类型** : 字典类型（键值对），之前的获取请求参数章节中出现的 url/values类型其实就是从map类型中扩展出来的

` map` 的初始化可以使用 ` make` ：

` var data = make ( map [ string ] string ) data = map [ string ] string {} 复制代码`
> 
> 
> 
> 
> make是内置函数，只能用来初始化 map、slice 和
> chan，并且make函数和另一个内置函数new不同点在于，它返回的并不是指针，而只是一个类型。
> 
> 

map赋值于其他语言的字典对象相同，取值有两种方式，请看下面的代码：

` data[ "name" ]= "zeta" //赋值 name := data[ "name" ] //方式1.普通取值 name,ok := data[ "name" ] //方式2.如果不存在name键，ok为false 复制代码`
> 
> 
> 
> 
> 代码中的变量ok，可以用来判断这一项是否设置过，取值时如果项不存在，是不会异常的，取出来的值为该类型的零值，比如
> int类型的值，不存在的项就为0；string类型的值不存在就为空字符串，所以通过值是否为0值是不能判断该项是否设置过的。 ok，会获得true
> 或者 false，判断该项是否设置过，true为存在，false为不存在于map中。
> 
> 

Go中的map还有几个特点需要了解：

* ` map` 的项的顺序是不固定的，每次遍历排列的顺序都是不同的，所以不能用顺序判断内容
* ` map` 可以用 ` for...range` 遍历
* ` map` 在函数参数中是引用传递（Go语言中，只有map、slice、chan是引用传递，其他都是值传递）

#### 赋值给 “_” ####

Go有一个特点，变量定义后如果没使用，会报错，无法编译。一般情况下没什么问题，但是极少情况下，我们调用函数，但是并不需要使用返回值，但是不使用，又无法编译，怎么办？

" ` _` " 就是用来解决这个问题的， ` _` 用来丢弃函数的返回值。比如本例中， ` template.ParseFiles("./templates/index.html")` 除了返回模版对象外，还会返回一个 ` error` 对象，但是这样简单的例子，出错的可能性极小，所以我不想处理 ` error` 了，将 ` error` 返回值用“ ` _` ”丢弃掉。

> 
> 
> 
> 注意注意注意：在实际项目中，请不要丢弃error，任何意外都是可能出现的，丢弃error会导致当出现罕见的意外情况时，非常难于Debug。所有的error都应该要处理，至少写入到日志或打印到控制台。（切记，不要丢弃
> error ，很多Gopher们在这个问题上有大把的血泪史）
> 
> 

OK，到目前为止，用Go语言搭建一个简单的网页的核心部分就完成了。

### 等等 .js、.css、图片怎么办？ ###

对。例子里的模版全是HTML代码，一个漂亮的网页还必须用到图片、js脚本和css样式文件，可是...和PHP不同，请求路径是通过HandleFunc匹配到处理函数的，难道要把js、css和图片都通过函数输出后，再用HandleFunc和URL路径匹配？

# 处理好js、css和图片，才能做漂亮的网页，Go http静态文件的处理办法 #

以在index.html文件里引用一个index.js文件为例。

## 从代码开始 ##

` func main () { http.HandleFunc( "/" , myWeb) //指定相对路径./static 为文件服务路径 staticHandle := http.FileServer(http.Dir( "./static" )) //将/js/路径下的请求匹配到 ./static/js/下 http.Handle( "/js/" , staticHandle) fmt.Println( "服务器即将开启，访问地址 http://localhost:8080" ) err := http.ListenAndServe( ":8080" , nil ) if err != nil { fmt.Println( "服务器开启错误: " , err) } } 复制代码`

在项目的根目录下创建static目录，进入static目录，创建js目录，然后在js目录里创建一个index.js文件。

` alert( "Javascript running..." ); 复制代码`

打开之前的index.html文件,在后面加上 ` <script src="/js/index.js"></script>`

运行 ` $ go run main.go` ，访问 [http://localhost:8080]( https://link.juejin.im?target=http%3A%2F%2Flocalhost%3A8080 ) ，页面会弹出提示框。

## 解读 ##

页面在浏览器中运行时，当运行到 ` <script src="/js/index.js"></script>` 浏览器会请求 ` /js/index.js` 这个路径

程序检查到第一层路由匹配 ` /js/` ，于是用文件服务处理这次请求，匹配到程序运行的路径下相对路径 `./static/js` 。

匹配的设置是 ` main.go` 文件中这两句

` //指定相对路径./static 为文件服务路径 staticHandle := http.FileServer(http.Dir( "./static" )) //将/js/路径下的请求匹配到 ./static/js/下 http.Handle( "/js/" , staticHandle) 复制代码`

也可以写成一句，更容易理解

` //浏览器访问/js/ 将会以静态文件形式访问目录 ./static/js http.Handle( "/js/" , http.FileServer(http.Dir( "./static" ))) 复制代码`

很简单...但是，可能还是不满足需求，因为, 如果

` http.Handle("/js/", http.FileServer(http.Dir("./static")))` 对应到 ./static/js

` http.Handle("/css/", http.FileServer(http.Dir("./static")))` 对应到 ./static/css

` http.Handle("/img/", http.FileServer(http.Dir("./static")))` 对应到 ./static/img

` http.Handle("/upload/", http.FileServer(http.Dir("./static")))` 对应到 ./static/upload

这样所有请求的路径都必须匹配一个static目录下的子目录。

如果，我就想访问static目录下的文件，或者，js、css、img、upload目录就在项目根目录下怎么办？

http包下，还提供了一个函数 ` http.StripPrefix` 剥开前缀，如下：

` //http.Handle("/js/", http.FileServer(http.Dir("./static"))) //加上http.StripPrefix 改为 ： http.Handle( "/js/" , http.StripPrefix( "/js/" , http.FileServer(http.Dir( "./static" )))) 复制代码`

这样，浏览器中访问/js/时，直接对应到./static目录下，不需要再加一个/js/子目录。

所以，如果需要再根目录添加多个静态目录，并且和URL的路径匹配，可以这样：

` http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))` 对应到 ./js

` http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))` 对应到 ./css

` http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))` 对应到 ./img

` http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload"))))` 对应到 ./upload

到这里，一个从流程上完整的Web服务程序就介绍完了。

整理一下，一个Go语言的Web程序基本的流程：

* 定义请求处理函数
* 用http包的HandleFunc匹配处理函数和路由
* ListenAndServe开启监听

当有http请求时：

* http请求到监听的的端口
* 根据路由将请求对象和响应写入器传递给匹配的处理函数
* 处理函数经过一番操作后，将数据写入到响应写入器
* 响应给请求的浏览器

## 最后编译程序 ##

之前调试都使用的是 ` go run` 命令运行程序。

您会发现，每次运行 ` go run` 都会重新编译源码，如何将程序运行在没有Go环境的计算机上？

使用 ` go build` 命令，它会编译源码，生成可执行的二进制文件。

最简单的 ` go build` 命令什么参数都不用加，它会自动查找目录下的main包下的main()函数，然后依次查找依赖包编译成一个可执行文件。

其他依赖文件的相对路径需要和编译成功后的可执行文件一致，例如范例中的templates文件夹和static文件夹。

默认情况下， ` go build` 会编译为和开发操作系统对应的可执行文件，如果要编译其他操作系统的可执行文件，需要用到交叉编译。

例如将Linux和MacOSX系统编译到windows

` GOOS=windows GOARCH=amd64 go build`

在Windows上需要使用SET命令, 例如在Windows上编译到Linux系统

` SET GOOS=linux SET GOARCH=amd64 go build main.go 复制代码`

# 结语，学到了什么？还要学什么？ #

## 学到了什么？ ##

* 快速简单搭建Go开发环境
* 导入包、申明包
* func 定义函数
* 变量的申明方法
* Go语言的异常处理
* for循环
* map类型
* 用http包，编写一个网站程序

本系列内容很少，很简洁，希望您能对Go多一点点了解，对Go多增加一点点兴趣。

## 没有涉及的其他知识 ##

还有很多内容成为一个合格的Gopher必须要了解的知识

* struct 结构体
* 给struct定义方法
* interface 接口定义和实现
* chan类型
* slice类型
* goroutine
* panic处理

以后的文章中会涉及更多关于Go语言编程的内容

**欢迎关注晓代码公众号，和大家一起学习吧**

![image.png](https://user-gold-cdn.xitu.io/2019/5/12/16aaa35ba7fb5c03?imageView2/0/w/1280/h/960/ignore-error/1)