# 用一句话描述Go语言的命令 #

**Go命令是管理Go资源的工具**

> 
> 有一些命令是非常常用的，比如 **run、build、get、test、get** ，有一些命令在使用IDE后很少会用到，IDE代劳了，比如 **fmt、vet**
> 。
> 把所有命令列出来，了解一下这些命令的用途， **对写代码很有帮助**
> 看看有没有你还没用过的命令吧！

**常规用法:**

` go <命令> [参数] 复制代码`

### 命令: ###

* **bug** ：创建一个bug报告
执行完命令后，会用浏览器访问 [github.com/golang/go]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttp%253A%2F%2Fgithub.com%2Fgolang%2Fgo ) 的issue。自动填写一些内容，引导你如何提交一个bug报告

* **build** ：编译包以及其依赖
最常用的命令之一。默认情况下，会在命令所在目录生成一个当前操作系统对应的可执行文件。安装完整版的Go环境，可以交叉编译其他操作系统的二进制可执行文件

* **clean** ：清空对象文件和缓存文件
前面提到的build命令和下面的test命令会生成一些文件和目录，clean会清理掉这些文件，包括build命令生成可执行文件

* **doc** ：打印包中的文档和标记符
打印出包或指定文件的说明文档，加上-all 参数，可以看到包里的所有函数列表和文档。
创建一个go文件，写入一下代码

` /* 这是一个范例 */ package main import "fmt" //main 主函数 func main () { SayHi() } //SayHi 打印字符串Hello world func SayHi () { fmt.Println( "Hello world!!" ) } 复制代码`

执行命令

` go doc -all -u 复制代码`
![](https://user-gold-cdn.xitu.io/2019/4/8/169fc30b81818306?imageView2/0/w/1280/h/960/ignore-error/1)

* **env** ：打印出你现在的Go环境信息
查看各个go的开发环境参数，忘记GOPATH和GOROOT路径就可以用这个打印出来了

![](https://user-gold-cdn.xitu.io/2019/4/8/169fc30b81a99bc5?imageView2/0/w/1280/h/960/ignore-error/1)

* **fix** ：用go的新版本的API更新

` go fix [packages] 复制代码`

如果你升级了go，担心以前的代码不兼容，那么就可以用 go fix

* **fmt** ：自动格式化代码文件
go的代码格式标准是唯一的，用go fmt可以格式化代码文件，很多IDE就是调用这个命令来在保存文件时调整格式。

* **generate** ：可以执行指令，包括生成和更新go源码文件的指令
查找当前包相关的源代码文件，找出所有包含”//go:generate”的注释，提取并执行该特殊注释后面的命令，类似shell执行命令。

![](https://user-gold-cdn.xitu.io/2019/4/8/169fc30b8357582d?imageView2/0/w/1280/h/960/ignore-error/1)

例子里只是调用了系统的echo指令，打印字符串，实际用途可以用generate生成go的类文件。 **（可能需要写一个小例子来说明其用途，日后的文章中再具体研究吧。）**

* **get** ：下载和安装go包以及其依赖包的命令

` go get <包的路径> 复制代码`

* **install** ：编译和安装包及其依赖包
可执行文件会被安装在$GOPATH/bin目录下。

* **list** ：列出目录下的所有包和模块，每行一个。

* **mod** ：详细内容可以参考文章： [拜拜了，GOPATH君！新版本Golang的包管理入门教程]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F60703832 )

* **run** ： 运行go项目
非常常用。
它会编译包，然后直接运行起来，不会在当前目录生成二进制文件。

* **test** ：运行调试
用于运行_text.go文件中的Test开头并且参数为 *testing.T的函数

![](https://user-gold-cdn.xitu.io/2019/4/8/169fc30b8428c25b?imageView2/0/w/1280/h/960/ignore-error/1)

* **tool** ：运行指定的go工具

* **version** ：查看当前go版本

* **vet** ：查看包中可能出现的错误
例如，给整型%d占位符提供一个字符串参数，就会检查出类型错误，但是这个代码编译是不会报错的。

![](https://user-gold-cdn.xitu.io/2019/4/8/169fc30b85f8cd4b?imageView2/0/w/1280/h/960/ignore-error/1)

**总结**

这些命令大部分使用起来都很简单，想了解更多可以运行go help [命令名]查看详细说明。

也有一些命令使用起来是需要花点时间学习的，比如 **generate、test、mod** ，如果有想要了解更多关于Go语言开发的同学，可以在评论区或私信告诉我们，一起学习一起讨论。

“晓代码”公众号：

![](https://user-gold-cdn.xitu.io/2019/4/8/169fc311ea48fb2b?imageView2/0/w/1280/h/960/ignore-error/1)