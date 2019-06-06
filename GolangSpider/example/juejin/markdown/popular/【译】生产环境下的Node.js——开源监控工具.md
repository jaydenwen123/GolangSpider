# 【译】生产环境下的Node.js——开源监控工具 #

你认为Node.js应用程序可以拥有的最重要的功能是什么？ 是花哨的全文模糊匹配搜索，还是用socket进行实时聊天呢？ 你能告诉我可以添加到Node.js应用中的最高级，最惊人和最吸引人的功能是什么么？

想知道我的么？ **高性能和不间断服务** 。高性能应用程序需要做好以下三点：

* 最短的停机时间；
* 可预测的资源使用率；
* 根据负载有效扩展

在第1部分， [Node.js要监控的关键指标]( https://link.juejin.im?target=https%3A%2F%2Fsematext.com%2Fblog%2Ftop-nodejs-metrics-to-watch%2F ) 中，我们讨论了您应该监控的关键Node.js指标，以便了解应用程序的运行状况。 我还解释了你应该避免的Node.js中的错误做法，例如阻塞线程和造成内存泄漏，还有一些巧妙的技巧可以用来提高应用程序的性能，比如使用集群模块创建工作进程和将长时间运行的任务从主线程分离开来用独立线程运行。

在本文中，我将描述如何使用5种不同的开源工具监控Node.js应用程序。它们可能没有像 ` Sematext` 或 ` Datadog` 那样功能全面，但它们是开源产品，可以完全由自己控制。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cec2dfd0e848?imageView2/0/w/1280/h/960/ignore-error/1)

## Appmetrics ##

Node应用指标监控看板显示了运行中的Node.js应用程序的性能数据。这是一个简单的模块，在Node.js入口文件的顶部应用并初始化。你可以通过在终端中运行以下命令从npm安装。

` $ npm install appmetrics-dash 复制代码`

Appmetrics提供了一个非常易于使用的Web仪表板。为了获得所有由应用程序创建的HTTP服务的仪表板，你需要做的是在app.js(或者以其他命名的入口文件)文件中添加以下代码段。

` // Before all other 'require' statements require( 'appmetrics-dash' ).attach() 复制代码`

之后你将通过这个请求路径 ` /appmetrics-dash` 中看到大量有用的指标。

* CPU Profiling
* HTTP传入请求
* HTTP吞吐量
* 平均响应时间（前5名）
* CPU
* 内存
* 堆（Heap)
* 事件循环时间（Event Loop Times）
* 环境
* 其他请求
* HTTP出站请求

此工具不仅显示指标。它允许您直接从仪表板生成Node.js报告和堆快照(Heap Snapshots)。 除此之外，您还可以使用Flame Graphs,非常酷的开源工具。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cec809663eaf?imageView2/0/w/1280/h/960/ignore-error/1)

## Express Status Monitor ##

Express.js是当前Node.js开发人员的的首选框架。 [Express Status Monitor]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FRafalWilinski%2Fexpress-status-monitor ) 是一个非常简单的独立模块，您可以将其添加到Express应用。它公开了一个 ` /status` 路由，在Socket.io和Chart.js的帮助下报告实时服务器指标。

从npm安装即可。

` $ npm install express-status-monitor 复制代码`

安装完这个模块之后，你需要在其他中间件或者路由之前添加它。

` app.use( require ( 'express-status-monitor' )()) 复制代码`

之后一旦你运行你的应用，你就可以通过 ` /status` 路由检查你的Node.js指标。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cecb81bd2b92?imageView2/0/w/1280/h/960/ignore-error/1)

## Prometheus ##

除非你生活在原始时代，不然你一定听说过 [Prometheus]( https://link.juejin.im?target=https%3A%2F%2Fprometheus.io%2F ) 。这是目前我们能使用的最著名的开源监控工具。Prometheus 100%开源并由社区驱动。所有的组件在遵从Apache 2 License开源协议并可以从GitHub下载。它是由 [CNCF]( https://link.juejin.im?target=https%3A%2F%2Fcncf.io%2F ) （Cloud Native Computing Foundation）管理并已经毕业 [成员项目]( https://link.juejin.im?target=https%3A%2F%2Fwww.cncf.io%2Fprojects%2F ) 之一，跟它同样的成员项目包括 ` Kubernetes` 和 ` Fluentd` 等。

要开始使用Prometheus进行监控，您需要下载最新版本并进行安装。

` $ tar xvfz prometheus-\*.tar.gz $ cd prometheus-\* 复制代码`

然后通过运行可执行文件启动它，但在运行此命令之前，需要创建一个 ` prometheus.yml` 文件。 它是一个配置文件，用于配置在哪些targets上，通过抓取HTTP端点数据监控哪些指标。

` # prometheus.yml scrape_configs: - job_name: 'prometheus' scrape_interval: 1 s static_configs: - targets: ['127.0.0.1:3000'] labels: service: 'test-prom' group: 'production' 复制代码`

现在你可以使用Prometheus了。

` $ ./prometheus --config.file=prometheus.yml 复制代码`

但是，我很懒，而且我非常喜欢Docker。 所以我的做法是运行官方的 [Prometheus Docker镜像]( https://link.juejin.im?target=https%3A%2F%2Fhub.docker.com%2Fr%2Fprom%2Fprometheus%2F ) ，避免下载它的所有麻烦。

## Prometheus and Docker ##

首先，进到Node.js应用程序的根目录。在这里，创建一个 ` prometheus-data` 目录并将 ` prometheus.yml` 文件放入其中。完成此操作后，运行Prometheus Docker容器。

获取正式的Prometheus Docker镜像并使用docker run命令运行该镜像。

` $ docker run -d \ --name prometheus \ --network= "host" \ -v " $(pwd) /prometheus-data" :/prometheus-data \ prom/prometheus \ --config.file=/prometheus-data/prometheus.yml 复制代码`

我选择使用-network =“host”运行容器，让Prometheus容器可以通过本机localhost地址访问，并且这样做，Node.js应用程序的也能通过本机HTTP端口访问到。否则，如果你将Prometheus和Node.js分别运行在容器内，则需要在两者之间建立一个 [网络]( https://link.juejin.im?target=https%3A%2F%2Fdocs.docker.com%2Fnetwork%2Fbridge%2F ) ，以便彼此之间只能相互访问到。

` -v` 选项用于将 ` prometheus-data` 目录从主机映射到容器内的同名目录。

在Prometheus容器运行后，需要在Node.js应用程序中添加配置的代码以暴露一个监控数据接口。 首先需要从npm安装适用于Node.js的 [Prometheus客户端]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fprom-client ) 。

` $ npm install prom-client 复制代码`

接在添加相关Prometheus相关配置代码

` // after all 'require' statements const client = require ( 'prom-client' ) const collectDefaultMetrics = client.collectDefaultMetrics collectDefaultMetrics({ timeout : 1000 }) app.get( '/metrics' , (req, res) => { res.set( 'Content-Type' , client.register.contentType) res.end(client.register.metrics()) }) 复制代码`

接下来你只需运行Node.js应用之后，通过 ` http://localhost:9090/graph` 就可以看到Prometheus图表

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cecfc9095b31?imageView2/0/w/1280/h/960/ignore-error/1)

## Clinic.js ##

Clinic.js包含三个工具，可帮助诊断和查明Node.js性能问题。它的使用非常简单。你需要做的就是从npm安装模块并运行它。它将为您生成报告，使故障排除变得更加容易。

使用如下命令安装Clinic.js

` $ npm install clinic 复制代码`

一旦安装完毕，就可以选择要生成的报告类型了。你可以选择以下三种报告类型。

* Doctor * 通过注入探针来收集指标
* 评估健康和启发式
* 提供修复推荐

* Bubbleprof- 一种全新的，完全独特的方法来分析Node.js代码 * 使用async_hooks收集指标
* 跟踪操作之间的延迟
* 创建气泡图

* Flame - 使用火焰图揭示代码中的瓶颈和热路径 * 通过CPU采样收集指标
* 跟踪栈顶频率
* 创建火焰图

让我们从运行Doctor并测试node.js应用程序开始。

` $ clinic doctor -- node app.js 复制代码`

在程序运行时，使用压测工具运行负载测试。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b206c53ccf6ef9?imageView2/0/w/1280/h/960/ignore-error/1)

` $ loadtest -n 1000 -c 100 [http://localhost:3000/api](http://localhost:3000/api) 复制代码`

一旦完成运行，停止服务器和Clinic.js Doctor将打开您可以查看的报告。

使用相同的方法，您可以运行Bubbleprof或Flame并获取相应工具的图形报告。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b206c8fa518628?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/4/16b206caee2bf04b?imageView2/0/w/1280/h/960/ignore-error/1)

## PM2 ##

使用PM2在生产中运行Node.js应用程序变得更加容易。 它是一个进程管理器，可以轻松地让您以集群模式运行应用程序。通俗来说，它将为您的主机每个CPU核心都生成一个进程。

首先安装 [PM2]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fpm2 )

` $ npm install pm2 -g 复制代码`

安装完成后，如果您的主源文件是app.js，则通过在终端中运行此命令来生成PM2守护程序。

` $ pm2 start app.js -i 0 复制代码`

` -i 0` 标志实例个数。这将以集群模式运行Node.js应用程序，其中数字0表示CPU核心数。你可以手动输入你想要的任何数字，但让PM2计算核心个数并自动产生相应个数的工作进程更简单些。

使用PM2查看 [Node.js监控数据]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FUnitech%2Fpm2 ) 也很容易

` $ pm2 monit 复制代码`

此命令将在终端中打开仪表板。在这里，您可以监视进程，日志，循环延迟，进程内存和CPU。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b206cde5caa007?imageView2/0/w/1280/h/960/ignore-error/1)

## 使用开源工具将监控Node.js方案进行包装 ##

性能指标对于让用户满意至关重要。在本文中，我向您展示了如何使用5种不同的开源工具向Node.js应用程序添加监视。 在了解了本系列第1部分 [Node.js要监控的关键指标]( https://link.juejin.im?target=https%3A%2F%2Fsematext.com%2Fblog%2Ftop-nodejs-metrics-to-watch%2F ) 之后，添加工具来监控现实生活中的应用程序是自然的学习进程。 本系列的最后一部分将介绍 [使用Sematext进行生产环境下Node.js监控]( https://link.juejin.im?target=https%3A%2F%2Fsematext.com%2Fblog%2Fnodejs-monitoring-made-easy-with-sematext%2F ) 。

如果你想查看示例代码，这里是一个包含所有的 [实例代码 repo]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fadnanrahic%2Fnodejs-monitoring-sematext%2Ftree%2Fdevelop ) 。你还可以克隆下来并选择任何工具打开。

如果你需要更多软件的全栈可观察性，请查看 [Sematext]( https://link.juejin.im?target=https%3A%2F%2Fsematext.com%2F ) 。 我们正在推动 [开源我们的产品]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsematext ) 并产生影响。

原文： [Node.js Open-Source Monitoring Tools]( https://link.juejin.im?target=https%3A%2F%2Fdev.to%2Fsematext%2Fnode-js-open-source-monitoring-tools-440a )

## 参考资料 ##

* [Prometheus操作指南]( https://link.juejin.im?target=https%3A%2F%2Fwww.ctolib.com%2Fdocs%2Fsfile%2Fprometheus-book%2Findex.html )
* [Prometheus 入门与实践]( https://link.juejin.im?target=https%3A%2F%2Fwww.ibm.com%2Fdeveloperworks%2Fcn%2Fcloud%2Flibrary%2Fcl-lo-prometheus-getting-started-and-practice%2F )