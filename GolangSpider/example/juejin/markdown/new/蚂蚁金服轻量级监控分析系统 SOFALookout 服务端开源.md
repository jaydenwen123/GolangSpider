# 蚂蚁金服轻量级监控分析系统 SOFALookout 服务端开源 #

> 
> 
> 
> **SOFA** Stack
> 
> 
> 
> 
> **S** calable **O** pen **F** inancial **A** rchitecture Stack
> 是蚂蚁金服自主研发的金融级分布式架构，包含了构建金融级云原生架构所需的各个组件，是在金融场景里锤炼出来的最佳实践。
> 
> 
> 
> SOFALookout 是蚂蚁金服在 SOFAStack 体系内研发开源的一款解决系统的度量和监控问题的轻量级中间件服务。本文给大家介绍下
> SOFALookout 服务器端主要提供的特性以及使用方式。
> 
> SOFALookout： [github.com/sofastack/s…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsofastack%2Fsofa-lookout
> )

## 前言 ##

容器，K8S，微服务，Mesh 以及 Serverless 这些新技术方向正在根本的变革我们运行软件的方式。我们构建的系统更加分布式化，另外由于容器，系统的生命周期更加短，变得易逝。针对这些变化，SOFALookout 希望提供一套轻量级解决方案。之前 SOFALookout 已经开源客户端的能力。今天，SOFALookout 服务器端 Metrics 部分的代码终于正式开源啦！本文给大家介绍下 SOFALookout 服务器端的主要特性以及使用方法。

## 什么是 SOFALookout ##

SOFALookout 是蚂蚁金服开源的一款解决系统的度量和监控问题的轻量级中间件服务。它提供的服务包括：Metrics 的埋点、收集、加工、存储与查询等。该开源项目包括了两个独立部分，分别是客户端与服务器端服务。

SOFALookout 目标是打造一套轻量级 Observability 实时工具平台，帮助用户解决基础设施、应用和服务等的监控和分析的问题。SOFALookout（目前已开源部分） 是一个利用多维度的 metrics 对目标系统进行度量和监控的项目。SOFALookout 的多维度 metrics 参考 [Metrics2.0]( https://link.juejin.im?target=http%3A%2F%2Fmetrics20.org%2F ) 标准。

SOFALookout ： [github.com/sofastack/s…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsofastack%2Fsofa-lookout )

SOFALookout 安装文档： [www.sofastack.tech/sofa-lookou…]( https://link.juejin.im?target=https%3A%2F%2Fwww.sofastack.tech%2Fsofa-lookout%2Fdocs%2Fquickstart-metrics-server )

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa6e394734fb?imageView2/0/w/1280/h/960/ignore-error/1)

SOFALookout 服务器端的主要特性:

* 适配社区主要 Metrics 数据源协议写入（比如: [Prometheus]( https://link.juejin.im?target=https%3A%2F%2Fprometheus.io%2F ) ， [Metricbeat]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fguide%2Fen%2Fbeats%2Fmetricbeat%2F6.4%2Findex.html ) 等）；
* 数据的存储支持扩展，暂时开源版默认支持 [Elasticsearch]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fcn%2Fproducts%2Felasticsearch ) ， 并且透明和自动化了相关运维操作；
* 遵循 Prometheus 查询 API 的标准以及支持 [PromQL]( https://link.juejin.im?target=https%3A%2F%2Fprometheus.io%2Fdocs%2Fprometheus%2Flatest%2Fquerying%2Fbasics%2F ) ，并进行了适当改进；
* 自带数据查询的控制台，并支持 [Grafana]( https://link.juejin.im?target=https%3A%2F%2Fgrafana.com%2F ) 进行数据可视化；
* 使用简单，支持单一进程运行整个服务器端模块。

随着 SOFALookout （metrics）服务器端代码开源，metrics 数据的处理已经形成闭环。后续我们将会进一步开源 Trace 和 Event 相关的服务能力，敬请期待。

## SOFALookout 项目结构 ##

服务器端代码分别包括两部分：Gateway 模块和 Server 模块。如下图所示（展示了 SOFALookout 源码项目的模块概要结构）

` ├── boot ├── client ├── gateway └── server 复制代码`

项目中的 boot 模块作用是方便集成和运行服务端的模块，既可以单独运行 Gateway 和 Server 的服务，也可以借助 SOFAArk 完成（Gateway 和 Server）的 All in One 的合并为单一进程运行。

## SOFALookout 工作机制 ##

下图完整展示了 SOFALookout 如何从 metrics 数据采集、上报、存储到最终展示的完整流程路径。

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa6e3ac68e18?imageView2/0/w/1280/h/960/ignore-error/1)

目前 SOFALookout 支持灵活的 metrics 数据存储选型。但开源版本我们暂时只支持了 Elasticsearch 作为存储的方案（后续可能继续支持 Cassandra,InfluxDB...），其他存储地适配我们希望更多同学能参与共建和支持。优先支持 Elasticsearch 是因为我们考虑到了 ELK 解决方案在业界已经广泛使用，尤其是日志数据。

为了开箱即用，同时考虑到不熟悉 Elasticsearch 的同学的使用，SOFALookout已经内置了关于 metrics 数据存储的自动化运维工具，可以免除大家自己建 Index，和日常维护 ES Index 的麻烦，更多细节后续单独讲解。

## 本次新增开源模块 ##

### 一、SOFALookout Gateway 模块 ###

SOFALookout Gateway 轻量的数据管道，它提供丰富的协议接入支持，包括自有SDK（SOFALookout Client）上报协议，还支持 Prometheus 的数据协议（推模式和拉模式），Metricbeat 协议（版本是6）， [OpenTSDB]( https://link.juejin.im?target=http%3A%2F%2Fopentsdb.net%2F ) 写入协议。每种数据来源对应于一个 Importer 的概念。

SOFALookout Gateway 对于远程（推模式）上报提供本地硬盘缓冲的支持。Gateway 总体设计是围绕数据加工的Pipeline 形式，包括前置后置的数据过滤器方便进行开发者数据加工。 另外 Gateway 可以支持自定义 Exporter，默认提供了 Elasticsearch Exporter，Standard Exporter(用于 Gateway 间数据中继)，开发者也可以自定义其他存储的 或 Kafka 等各式各样 Exporter。

### 二、SOFALookout Server 模块 ###

SOFALookout Server 兼容和增强了 Prometheus 的数据及元数据查询的 RESTful API。同样对应 PromQL 我们也基本实现了兼容和增强（不包括 Alert 相关语法），SOFALookout 的 promQL 相关解析逻辑是从 Prometheus 移植而来，做了一些优化和改进， 感谢 Prometheus 开源了如此易用和强大的 golang 版本的 QL 实现。

为了方便方便开发者做数据探索和试验，我们也提供了自有 Web-UI 的支持，能够满足基本功能使用。

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa6e3a22906b?imageView2/0/w/1280/h/960/ignore-error/1)

我们还是推荐大家使用 Grafana 进行数据展示。Grafana 集成 SOFALookout 很简单，只需要选择 Prometheus 作为数据源协议即可（SOFALookout默认查询端口也是: 9090）。下图展示 Grafana 新增数据源配置：

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa6e3c493439?imageView2/0/w/1280/h/960/ignore-error/1)

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa6e3bb8a166?imageView2/0/w/1280/h/960/ignore-error/1)

## ##

## 近期计划 ##

下图是近期的 Roadmap：

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa6e3c3cb57d?imageView2/0/w/1280/h/960/ignore-error/1)

非常欢迎更多同学参与 SOFALookout 共建，尤其是支持更多的 Metrics 存储库。

### 公众号：金融级分布式架构（Antfin_SOFA） ###