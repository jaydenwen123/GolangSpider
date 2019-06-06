# 使用go mod结合docker分层缓存进行自动CI/CD #

## 喜大奔的go mod ##

` 官方背书的go mod拯救了我的代码洁癖症! 复制代码`

## 环境 ##

* go v1.12
* docker ce 18.09.0
* gitlab ce latest

### godep ###

写go程序，若是仅仅是你一个人写，或者就是写个小工具玩儿玩儿，依赖管理对你来说可能没那么重要。

但是在商业的工程项目里，多人协同，go的依赖管理就尤为重要了，之前可选的其实不太多，社区提供的实现方式大多差不多的思路，比如我之前使用的 ` godep` 。所以项目中会有一个 ` vendor` 文件夹来存放外部的依赖，这样：

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256c45ce9df?imageView2/0/w/1280/h/960/ignore-error/1)

这样的实现方式，每次更新了外部依赖，其他人就得拉下来一大坨。。。

### go mod ###

来看看使用官方的module来管理依赖的工程结构：

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256cb0f51e5?imageView2/0/w/1280/h/960/ignore-error/1)

是不是，清爽无比，项目也整个瘦身了！

简单的说一下go mod help，至于开启go mod的步骤，其他网文一大堆，就不复制了。毕竟本文是说go工程CI/CD的。

在目前 ` go v1.12` 版本下，命令 ` go mod help` 结果如下：

` The commands are: download download modules to local cache edit edit go.mod from tools or scripts graph print module requirement graph init initialize new module in current directory tidy add missing and remove unused modules vendor make vendored copy of dependencies verify verify dependencies have expected content why explain why packages or modules are needed 复制代码`

后面CI/CD需要用到的是 ` download` 指令。

## dockerfile ##

来看看我这个工程的dockerfile:

` FROM golang: 1.12 as build ENV GOPROXY https://go.likeli.top ENV GO111MODULE on WORKDIR /go/cache ADD go.mod . ADD go.sum . RUN go mod download WORKDIR /go/release ADD. . RUN GOOS=linux CGO_ENABLED=0 go build -ldflags= "-s -w" -installsuffix cgo -o app main.go FROM scratch as prod COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt COPY --from=build /go/release/app / COPY --from=build /go/release/conf.yaml / CMD [ "/app" ] 复制代码`

我这个项目有一些外部依赖，在本地开发的时候都已调整好，并且编译通过，在本地开发环境已经生成了两个文件 ` go.mod` 、 ` go.sum`

在dockerfile的第一步骤中，先启动module模式，且配置代理，因为有些墙外的包服务没有梯子的情况下也是无法下载回来的，这里的代理域名是我自己的，有需要的也可以用。

指令 ` RUN go mod download` 执行的时候，会构建一层缓存，包含了该项所有的依赖。之后再次提交的代码中，若是 ` go.mod` 、 ` go.sum` 没有变化，就会直接使用该缓存，起到加速构建的作用，也 ` 不用重复的去外网下载依赖` 了。若是这两个文件发生了变化，就会重新构建这个缓存分层。

使用缓存构建的效果：

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256c2c31fa7?imageView2/0/w/1280/h/960/ignore-error/1)

这个加速效果是很明显的。

### 减小体积 ###

#### go构建命令使用 ` -ldflags="-s -w"` ####

在官方文档： [Command_Line]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fcmd%2Flink%2F%23hdr-Command_Line ) 里面说名了 ` -s -w` 参数的意义，按需选择即可。

* ` -s` : 省略符号表和调试信息
* ` -w` : 省略DWARF符号表

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256c8e1fca5?imageView2/0/w/1280/h/960/ignore-error/1)

看起来效果不错🙂

#### 使用scratch镜像 ####

使用 ` golang:1.12` 开发镜像构建好应用后，在使用 ` scratch` 来包裹生成二进制程序。

关于 ` 最小基础镜像` ，docker里面有这几类：

* scratch: 空的基础镜像，最小的基础镜像
* busybox: 带一些常用的工具，方便调试， 以及它的一些扩展busybox:glibc
* alpine: 另一个常用的基础镜像，带包管理功能，方便下载其它依赖的包

## 镜像瘦身最终效果 ##

好了，看看最终构建的应用的效果：

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256c5a69d9c?imageView2/0/w/1280/h/960/ignore-error/1)

构建的镜像大小为: ` 16.4MB`

## CI/CD ##

基于gitlab的runner来进行CI/CD，看看我的 `.gitlab-ci.yml` 配置：

` before_script: - if [[ $(whereis docker-compose | wc -l) -eq 0 ]]; then curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && chmod +x /usr/local/bin/docker-compose; fi # ****************************************************************************************************** # ************************************** 测试环境配置 **************************************************** # ****************************************************************************************************** deploy-test-tour: stage: deploy tags: - build only: - release/v2.0 script: - export PRODUCTION=false - docker-compose stop - docker-compose up -d --build # ****************************************************************************************************** # ************************************** 生产环境配置 **************************************************** # ****************************************************************************************************** deploy-prod-tour: stage: deploy tags: - release only: - master script: - export PRODUCTION=true - docker-compose stop - docker-compose up -d --build 复制代码`

我使用 ` docker-compose` 来进行容器控制，所以在 ` before_script` 过程里面增加了这一步，方便新机器的全自动化嘛。

我这个项目做了点儿工程化，所以稍微正规点儿，分出了两个环境，测试和生产环境。分别绑定到不同的分支上。

正主就是下面执行的这三行：

` export PRODUCTION=false docker-compose stop docker-compose up -d --build 复制代码`

* ` export` 控制一下临时环境变量，方便发布不同的环境。
* ` docker-compose stop` 停止旧的容器
* ` docker-compose up -d --build` 编排新的容器并启动，会使用之前的缓存分层镜像，所以除了第一次构建，后面的速度都是杠杠的。

看实际的发布截图:

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256c5f634fd?imageView2/0/w/1280/h/960/ignore-error/1)

**首次执行，总共：1 minute 22 seconds**

![](https://user-gold-cdn.xitu.io/2019/3/13/16975256edd1d7ea?imageView2/0/w/1280/h/960/ignore-error/1)

**使用缓存构建，总共：33 seconds**