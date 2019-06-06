# Golang环境变量设置详解 #

无论你是使用Windows,Linux还是Mac OS操作系统来开发Go应用程序，在安装好Go安装语言开发工具之后，都必须配置好Go语言开发所要求的 环境变量，才算初步完成Go开发环境的搭建。

但对于一些初学者来说，可能不太明白Go语言中常用的 ` 环境变量(Environment variables)` 的作用以及如何设置环境变量，今天我们来讲讲。

> 
> 
> 
> 注意：Go提供的Windows操作系统安装包(.msi后缀)安装完成后，会自动配置几个常用的环境变量。
> 
> 

## 常用环境变量 ##

Go语言中可以设置的环境变量有很多，每个环境变量都有其作用，不过很多我们可能都用不到，一般需要了解也是最重要的有以下几个：

` $ go env //打印Go所有默认环境变量 $ go env GOPATH //打印某个环境变量的值 复制代码`

#### GOROOT ####

环境变量 ` GOROOT` 表示Go语言的安装目录。

在 ` Windows` 中， ` GOROOT` 的默认值是 ` C:/go` ，而在 ` Mac OS` 或 ` Linux` 中 ` GOROOT` 的默认值是 ` usr/loca/go` ，如果将Go安装在其他目录中，而需要将GOROOT的值修改为对应的目录。

另外， ` GOROOT/bin` 则包含Go为我们提供的工具链，因此，应该将 ` GOROOT/bin` 配置到环境变量PATH中，方便我们在全局中使用Go工具链。

##### Linux设置GOROOT演示 #####

` export GOROOT=~/go export PATH= $PATH : $GOROOT /bin 复制代码`

#### GOPATH ####

> 
> 
> 
> 注意， ` GOPATH` 的值不能与 ` GOROOT` 相同。
> 
> 

环境变量 ` GOPATH` 用于指定我们的开发工作区(workspace),是存放源代码、测试文件、库静态文件、可执行文件的工作。

在 ` 类Unix` (Mac OS或Linux)操作系统中 ` GOPATH` 的默认值是$home/go。而在Windows中GOPATH的默认值则为%USERPROFILE%\go(比如在Admin用户，其值为C:\Users\Admin\go)。

当然，我们可以通过修改GOPATH来更换工作区，比如将工作设置 ` opt/go` 方式如下：

##### Linux设置GOPATH演示 #####

` export GOPATH=/opt/go 复制代码`

还有，可以在GOPATH中设置多个工作区，如：

` export GOPATH=/opt/go; $home /go 复制代码`

##### GOPATH的子目录 #####

上面的代码表示我们指定两个工作区，不过当我们使用 ` go get` 命令去获取远程库的时候，一般会安装到第一个工作区当中。

按照Go开发规范，GOPATH目录下的每个工作一般分为三个子目录: ` src` , ` pkg` , ` bin` ，所以我们看到的每个工作区是这样子的：

` bin/ hello # 可执行文件 outyet # 可执行文件 src/ github.com/golang/example/ .git/ hello/ hello.go # 命令行代码 outyet/ main.go # 命令行代码 main_test.go # 测试代码 stringutil/ reverse.go # 库文件 reverse_test.go # 库文件 golang.org/x/image/ .git/ bmp/ reader.go # 库文件 writer.go # 库文件 复制代码`

` src` 目录放的是我们开发的源代码文件，其下面对应的目录称为 ` 包` , ` pkg` 放的是编译后的库静态文件， ` bin` 放的是源代码编译后台的可执行文件。

#### GOBIN ####

环境变量 ` GOBIN` 表示我们开发程序编译后二进制命令的安装目录。

当我们使用 ` go install` 命令编译和打包应用程序时，该命令会将编译后二进制程序打包GOBIN目录，一般我们将GOBIN设置为 ` GOPATH/bin` 目录。

##### Linux设置GOBIN演示 #####

` export GOBIN= $GOPATH /bin 复制代码`

上面的代码中，我们都是使用export命令设置环境变量的，这样设置只能在当前shell中有效，如果想一直有效，如在Linux中，则应该将环境变量添加到 ` /etc/profile` 等文件当中。

## 交叉编译 ##

什么是交叉编译？所谓的交叉编译，是指在一个平台上就能生成可以在另一个平台运行的代码，例如，我们可以32位的Windows操作系统开发环境上，生成可以在64位Linux操作系统上运行的二进制程序。

在其他编程语言中进行交叉编译可能要借助第三方工具，但在Go语言进行交叉编译非常简单，最简单只需要设置GOOS和GOARCH这两个环境变量就可以了。

#### GOOS与GOARCH ####

GOOS的默认值是我们当前的操作系统， 如果windows，linux,注意mac os操作的上的值是darwin。 GOARCH则表示CPU架构，如386，amd64,arm等。

##### 获取GOOS和GOARCH的值 #####

我们可以使用 ` go env` 命令获取当前GOOS和GOARCH的值。

` $ go env GOOS GOARCH 复制代码`

##### GOOS和GOARCH的取值范围 #####

GOOS和GOARCH的值成对出现，而且只能是下面列表对应的值。

` $GOOS $GOARCH android arm darwin 386 darwin amd64 darwin arm darwin arm64 dragonfly amd64 freebsd 386 freebsd amd64 freebsd arm linux 386 linux amd64 linux arm linux arm64 linux ppc64 linux ppc64le linux mips linux mipsle linux mips64 linux mips64le linux s390x netbsd 386 netbsd amd64 netbsd arm openbsd 386 openbsd amd64 openbsd arm plan9 386 plan9 amd64 solaris amd64 windows 386 windows amd64 复制代码`

#### 示例 ####

##### 编译在64位Linux操作系统上运行的目标程序 #####

` $ GOOS=linux GOARCH=amd64 go build main.go 复制代码`

##### 编译arm架构Android操作上的目标程序 #####

` $ GOOS=android GOARCH=arm GOARM=7 go build main.go 复制代码`

## 环境变量列表 ##

虽然我们一般虽然配置的环境变量就那么几个，但其实Go语言是提供了非常多的环境变量，让我们可以自由地定制开发和编译器行为。

下面是Go提供的所有的环境变量列表，一般可以划分为下面几大类，大概了解一下就可以了，因为有些环境变量我们可以永远都不会用到。

#### 通过环境变量 ####

` GCCGO GOARCH GOBIN GOCACHE GOFLAGS GOOS GOPATH GOPROXY GORACE GOROOT GOTMPDIR 复制代码`

#### 和cgo一起使用的环境变量 ####

` CC CGO_ENABLED CGO_CFLAGS CGO_CFLAGS_ALLOW CGO_CFLAGS_DISALLOW CGO_CPPFLAGS, CGO_CPPFLAGS_ALLOW, CGO_CPPFLAGS_DISALLOW CGO_CXXFLAGS, CGO_CXXFLAGS_ALLOW, CGO_CXXFLAGS_DISALLOW CGO_FFLAGS, CGO_FFLAGS_ALLOW, CGO_FFLAGS_DISALLOW CGO_LDFLAGS, CGO_LDFLAGS_ALLOW, CGO_LDFLAGS_DISALLOW CXX PKG_CONFIG AR 复制代码`

#### 与系统架构体系相关的环境变量 ####

` GOARM GO386 GOMIPS GOMIPS64 复制代码`

#### 专用的环境变量 ####

` GCCGOTOOLDIR GOROOT_FINAL GO_EXTLINK_ENABLED GIT_ALLOW_PROTOCOL 复制代码`

#### 其他环境变量 ####

` GOEXE GOHOSTARCH GOHOSTOS GOMOD GOTOOLDIR 复制代码`

## 小结 ##

环境变量的设置，可以影响我们开发和编译项目的过程与结果，所以还是很有必要了解一下的。