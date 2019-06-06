# GO笔记之GO命令快速体验 #

[上篇文章]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F62922404 ) 利用go run和go build命令分析介绍了GO的编译执行流程。GO提供给我们的命令当然远不止这两个。本文将在所能及的范围内，尽量地介绍GO提供的所有命令，从而实现对它们有个整体的认识。

## 概述 ##

除了gofmt与godoc外，GO中的命令一般都可通过go命令调用，这些命令可理解为go的子命令，查看下命令列表，如下：

` $ go Go is a tool for managing Go source code. Go是管理Golang源码的工具 Usage: 使用方式： go < command > [arguments] go <命令> [参数] The commands are: 涉及的命令包括： bug start a bug report 提交bug报告，执行后会开启浏览器并转到github的issue，当前的配置与环境都会自动填写到issue中 build compile packages and dependencies 编译源码和依赖包 clean remove object files and cached files 清理文件，如测试与编译中生成或缓存的文件 doc show documentation for package or symbol 可用于显示包、接口和函数等的文档 env print Go environment information 打印当前的环境变量信息 fix update packages to use new APIs 可用于go新旧版本之间的代码迁移，修正代码兼容问题 fmt gofmt (reformat) package sources 按规范格式化源码 generate generate Go files by processing source 扫描源码注释，类似//go:generate command argument...实现生成go文件 get download and install packages and dependencies 下载并安装包和依赖 install compile and install packages and dependencies 编译并安装包和依赖 list list packages or modules 列出包或模块的信息 mod module maintenance 用于模块的管理维护 run compile and run Go program 编译与执行Go程序 test test packages 测试包 tool run specified go tool 运行go提供的一些特定工具，比如pprof version print Go version 打印go版本信息 vet report likely mistakes in packages 检查与报告代码包中的错误 ... 复制代码`

输出的介绍大致翻译了下，其中部分命令也作了些介绍。除了go的子命令，go tool下也有些更底层的命令，执行go tool即可查看：

` $ go tool addr2line 可以调用栈的地址转化为文件和行号 asm 和汇编有关的命令，没搞清楚如何使用 buildid 似乎用在编译时，根据文件内容生成 hash cgo 可帮助我们实现在GO中调用C语言代码 compile 用于编译源码生成.o文件 cover 可用于分析测试覆盖率 dist 帮助引导、构建和测试go doc 测试下来似乎和go doc效果一样，都是用于文章管理 fix 用于解决不同版本间代码不兼容问题，和go fix作用一样 link 用于库的链接 nm 可列出如对象文件.o，可执行文件或.a库文件中的函数变量符号等信息 objdump 反汇编命令 pack 似乎是个打包压缩命令 pprof 自带的性能分析工具 test 2json 用于把测试文件转化可读的json格式 tour 启动本地的tour教程，可见GO团队真的很用心 trace 可用于问题诊断与调式的工具 go tool的输出默认没有任何文字说明，这里的介绍是我收集总结出来的，可能有些错误。 复制代码`

总体而言，我把GO中命令分为几个大类：

* 源码编译
* 包的管理
* 代码规范
* 测试相关
* 调试优化
* 其他命令

GO的命令很多，很难在一篇文章中把每个都介绍清楚。接下来只做些简单演示说明，详细的介绍待以后有了具体场景在详细说明。

## 源码编译 ##

go build、go run 两个命令可归于源码编译类的命令，是入门学习首先要掌握的，故而在前篇文章已经做了详细介绍。

go build用于编译可执行文件和库文件，参数可以是源码目录、.go文件。演示如下：

` $ go build # 目录 $ go build main.go # .go文件 $ go build main.go math.go data.go # .go文件列表 复制代码`

go run和build有很多相似点，它们都会编译源码。不同的是，go run只能用于可执行源码（即main包源码）的编译。接收参数为.go文件。演示如下：

` $ go run main.go $ go run main.go math.go data.go 复制代码`

详细了解，可以看下 详解GO的编译执行流程 这篇介绍。

关于编译，go tool还有一些更细节的命令，比如compile、link等。有兴趣可以去了解下。

## 包的管理 ##

GO的包管理是由语言包自带，这点不同于其他语言，如Java、Python、PHP等。从当前我所了解的来看，GO没有提供包管理仓库，而是直接使用的版本管理系统作为仓库，支持如git、svn、mercurial等。而用于包管理的命令有go install、go get、go list等。

go install用于源码包的编译与安装。虽然这里也涉及到编译，但从名字就可以看出，该命令重在强调安装。之前build命令在编译非main包会生成缓存文件，main包会生成执行文件并拷贝到当前目录。而install会将它们安装到指定的文件。假设有名为math的包，执行如下命令：

` $ go install math # 非main包，最终生成文件 $GOPATH/pkg/xxx/xxx/math.a。 $ go install entry # 是main包，最终生成文件 $GOBIN/entry 复制代码`

go get用于从互联网安装更新包和依赖，类似于其他语言包管理器的install。不同于install，它多出网络下载这步，大概可理解为 go get 等价于 git/svn等下载 + git install。我们可以演示下从http://github.com/PuerkitoBio/goquery下载安装goquery的过程，如下：

` $ go get -v github.com/PuerkitoBio/goquery golang.org/x/net/html/atom golang.org/x/net/html github.com/andybalholm/cascadia github.com/PuerkitoBio/goquery 复制代码`

从上面可以看出，go get不仅下载了goquery，还下载了相应的依赖。执行完成后，GOPATH目录下可以找到goquery的源码与编译后的.a库文件。

go list可用于输出包的信息，接口的参数和在源码中import的路径相同，下面演示一些案例：

` $ go list fmt # 查看某包 fmt $ go list fmt net/http # 查看多个包 fmt net/http $ go list --json fmt # 查看包的具体信息 { "Dir" : "/usr/local/go/src/fmt" , "ImportPath" : "fmt" , "Name" : "fmt" , "Doc" : "Package fmt implements formatted I/O with functions analogous to C's printf and scanf." , "Target" : "/usr/local/go/pkg/darwin_amd64/fmt.a" , "Root" : "/usr/local/go" , "Match" : [ "fmt" ], "Goroot" : true , "Standard" : true , "GoFiles" : [ "doc.go" , ... "scan.go" ], "Imports" : [ "errors" , ... "unicode/utf8" ], "Deps" : [ "errors" , ... "unicode/utf8" , "unsafe" ], "TestGoFiles" : [ "export_test.go" ], "XTestGoFiles" : [ "example_test.go" , ... "stringer_test.go" ], "XTestImports" : [ "bufio" , ... "unicode/utf8" ] } 复制代码`

详情信息部分展示内容比较多，包含源码路径、导入路径、依赖了哪些包等一系列信息。

## 代码规范 ##

这类命令可以帮助我们规范代码的格式，减少代码发生错误的几率，其中主要有go fmt、go vet和go fix三个命令。

go fmt的作用是代码的格式化。为了让我们把更多时间花在开发工作上，GO官方制定了标准的代码规范并go fmt实现规范代码。假设有main.go文件内容如下：

` package main func main () { a := x + y } 复制代码`

我们只需执行go fmt main.go文件，然后再次打开main.go文件：

` package main func main () { a := x + y } 复制代码`

格式化已经完成。关于代码格式化还有一个更具体的命令：gofmt，go fmt是它的某个特殊形式：gofmt -l -w。

go vet是一个用于检查GO语言静态语法的工具。GO语言的语法是非常严格的，如不能定义未使用的变量、变量类型必须显式转化等等。示例，假设有main.go文件内容如下：

` package main func main () { a := 1 + 2 b := 1 } 复制代码`

使用go vet执行源码检查，输出结果如下：

` $ go vet main.go # command-line-arguments [command-line-arguments.test]./main.go:4:5: a declared and not used ./main.go:5:5: b declared and not used 复制代码`

我们被告知，变量a和b声明但并没有使用。

go fix主要用于处理代码的兼容性问题，例如go1之前老版本的代码转化到go1。但是有点遗憾的是，没找到该命令的演示案例。我们平时应该很少用到。

## 测试相关 ##

GO也提供与相关的命令，为我们提供了一条方便验证我们代码的途径。与测试有关的命令有go test、go tool cover 和 go tool test2json。

go test可用于运行测试代码，以此验证程序逻辑的正确性能。具体演示下，示例代码包含两部分，分别是功能代码和测试代码。功能代码在math.go文件中，如下：

` package math func Add(x, y int) int { return x + y } 复制代码`

测试用例的代码在math_test.go文件中，如下：

` package math import "testing" func Test_Add(t *testing.T) { r := Add(1, 2) if r != 3 { t.FailNow() } } 复制代码`

接下来我们可以执行go test命令启动测试用例，如下：

` $ go test math_test.go math.go ok command -line-arguments 复制代码`

结果显示，测试执行成功，Add函数功能正常。我们可以把测试代码编译成可执行文件，如下：

` $ go test math_test.go math.go -o math.test 复制代码`

查看下会发现此时目录下多出了编译好的math.test可执行测试文件。

go cover可用于分析测试覆盖率。比如上面的测试案例，我们可以生成覆盖率文件，如下：

` $ go test *.go -coverprofile=coverage.out 复制代码`

go cover 提供多种方式分析测试覆盖率，这里演示下如何用html展示测试结果，如下：

` $ go tool cover --html=size_coverage.out 复制代码`

显示测试覆盖率为100%。我们这里的测试用例比较简单，所以到达了全面覆盖。

go tool test2json可用于将go test测试结果转化为json格式。这里需要使用之前生成的测试执行文件，示例如下：

` $ go tool test 2json ./math.test -test.v { "Action" : "run" , "Test" : "Test_Add" } { "Action" : "output" , "Test" : "Test_Add" , "Output" : "=== RUN Test_Add\n" } { "Action" : "output" , "Test" : "Test_Add" , "Output" : "--- PASS: Test_Add (0.00s)\n" } { "Action" : "pass" , "Test" : "Test_Add" } { "Action" : "output" , "Output" : "PASS\n" } { "Action" : "pass" } 复制代码`

测试结果以json格式打印出来。虽然我对专业的测试不太了解，但是也明白结构化的输出是比较容易程序化分析的。

## 调试优化 ##

完成代码开发后，可能时刻会遇到一些bug或者性能问题。GO提供了go tool pprof、go tool trace、go tool addr2line和go tool nm等一系列命令，可用于代码调试优化。

go tool pprof可用于帮助我们分析程序收集的性能数据，比如CPU、内存等数据。以官方提供的示例为例吧，博客地址在 博客。示例代码在benchgraffiti。

go tool trace 可用于追踪程序执行情况。go tool pprof可以通过cpu和内存数据分析出程序的瓶颈。

go tool addr2line可以将地址转化对应源码的文件和行号，非常方便的便于我们调式问题。

具体的案例就不演示了。这部分的命令稍微有点复杂，待后面有了具体案例再来补充。

## 其他命令 ##

GO中还提供了很多辅助命令。这些命令有go bug、go doc、go tool tourl等。

go bug会直接启动浏览器并进入github的go项目的issue之下，还会把用户当前环境信息自动添加到issue中。如下执行go bug之后跳转的页面：

![](https://user-gold-cdn.xitu.io/2019/5/20/16ad4dd9338f69e4?imageView2/0/w/1280/h/960/ignore-error/1)

由此可见，GO的开发团队真的非常用心，做了很多简化我们工作的事情。

go doc为我们提供了快速查看文档的途径，比如查看fmt文档，我们只需执行go doc fmt，fmt相关的文档便会输出到控制台。我们也可以像官网文档那样用浏览器查看文档，只需执行godoc -http=:6060，便会启动一个本地的web服务。我们访问localhost:6060就能看到一个几乎和官网一样的页面，示例如下：

![](https://user-gold-cdn.xitu.io/2019/5/20/16ad4dde28f59b9d?imageView2/0/w/1280/h/960/ignore-error/1)

go tool tourl是官方提供的本地搭建tour教程的方式。我们只需执行go tool tour便会自动启动浏览器并进入教程首页。

到这里我们可以发现，即使由于一些原因使我们无法访问GO的官网，但有了这些工具，我们也可以愉快地进行GO的学习。

## 总结 ##

本篇文章以GO命令的快速体验为目标，概要式地介绍了几乎所有的命令。在对它们有了基本认识后，在以后遇到问题时，我们才能想到它们，以及更快地掌握和使用它们。