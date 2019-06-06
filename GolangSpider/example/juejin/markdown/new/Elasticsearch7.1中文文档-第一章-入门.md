# Elasticsearch7.1中文文档-第一章-入门 #

## 入门 ##

### 引言 ###

Elasticsearch是一个高度可扩展开源的全文搜索引擎.它搜索几乎是实时的,用ES作为搜索引擎,为复杂搜索功能的需求提供解决方案.

ES的使用场景:

* 

网上商场,搜索商品.

* 

ES配合logstash,kibana,日志分析.

本教程的其他部分,将指导你完成ES的安装,启动,浏览,以及数据的CRUD.如果你完整的完成了本教程,你应该已经对ES有很好的了解了,希望你能从中受到启发.

### 基本概念 ###

ES有几个核心概念，从一开始理解这些概念将对你后面的学习有很大帮助。

**近实时（NRT）**

ES是一个近实时的搜索引擎（平台），代表着从添加数据到能被搜索到只有很少的延迟。（大约是1s）

**集群** ( https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2F%25E9%259B%2586%25E7%25BE%25A4%25E6%258A%2580%25E6%259C%25AF%2F9774443%3Ffr%3Daladdin )

可以将多台ES服务器作为集群使用，可以在任何一台节点上进行搜索。集群有一个默认的名称（可修改），“elasticsearch”，这个集群名称必须是唯一的，因为集群的节点是通过集群名称来加入集群的。

确保在相同环境中不要有相同的集群名称，否则有可能节点会加入到非预期的集群中。

**节点**

节点是作为集群的一部分的单个服务器，存储数据，并且参与集群的索引和搜索功能。与集群一样，节点由一个名称标识，默认情况下，该名称是在启动时分配给节点的随机通用唯一标识符（UUID）。如果不希望使用默认值，则可以定义所需的任何节点名称。此名称对于管理目的很重要，因为您希望确定网络中的哪些服务器对应于ElasticSearch集群中的哪些节点。

**索引**

索引是具有某种相似特性的文档集合。例如，您可以拥有客户数据的索引、产品目录的另一个索引以及订单数据的另一个索引。索引由一个名称（必须全部是小写）标识，当对其中的文档执行索引、搜索、更新和删除操作时，该名称用于引用索引。在单个集群中，您可以定义任意多个索引。

如果你学习过Mysql ，可以将其暂时理解为 MySql中的 database。

**类型**

一个索引可以有多个类型。例如一个索引下可以有文章类型，也可以有用户类型，也可以有评论类型。在一个索引中不能再创建多个类型，在以后的版本中将删除类型的整个概念。

> 
> 
> 
> 在Elasticsearch 7.0.0或更高版本中创建的索引不再接受 ` _default_` 映射。在6.x中创建的索引将继续像以前一样在Elasticsearch
> 6.x中运行。在7.0中的API中不推荐使用类型，对索引创建，放置映射，获取映射，放置模板，获取模板和获取字段映射API进行重大更改。
> 
> 

**文档**

一个文档是一个可被索引的基础信息单元。比如，你可以拥有某一个客户的文档，某一个产品的一个文档，当然，也可以拥有某个订单的一个文档。文档以JSON（Javascript Object Notation）格式来表示，而JSON是一个到处存在的互联网数据交互格式。

在一个index/type里面，你可以存储任意多的文档。注意，尽管一个文档，物理上存在于一个索引之中，文档必须被索引/赋予一个索引的type。

**分片和副本**

索引可能存储大量数据，这些数据可能会c超出单个节点的硬件限制。例如，占用1TB磁盘空间的10亿个文档的单个索引可能不适合单个节点的磁盘，或者速度太慢，无法单独满足单个节点的搜索请求。

为了解决这个问题，ElasticSearch提供了将索引细分为多个片段（称为碎片）的能力。创建索引时，只需定义所需的碎片数量。每个分片（shard）本身就是一个完全功能性和独立的“索引”，可以托管在集群中的任何节点上。

为什么要分片?

* 

它允许您水平拆分/缩放内容量

* 

它允许您跨碎片（可能在多个节点上）分布和并行操作，从而提高性能/吞吐量

如何分配分片以及如何将其文档聚合回搜索请求的机制完全由ElasticSearch管理，并且对作为用户的您是透明的。

在随时可能发生故障的网络/云环境中，非常有用，强烈建议在碎片/节点以某种方式脱机或因任何原因消失时使用故障转移机制。为此，ElasticSearch允许您将索引分片的一个或多个副本复制成所谓的副本分片，简称为副本分片。

为什么要有副本？

* 

当分片/节点发生故障时提供高可用性。因此，需要注意的是，副本分片永远不会分配到复制它的原始/主分片所在的节点上。

* 

允许您扩展搜索量/吞吐量，因为可以在所有副本上并行执行搜索。

总而言之，每个索引可以分割成多个分片。索引也可以零次（意味着没有副本）或多次复制。复制后，每个索引将具有主分片（从中复制的原始分片）和副本分片（主分片的副本）。

可以在创建索引时为每个索引定义分片和副本的数量。创建索引后，您还可以随时动态更改副本的数量。您可以使用收缩和拆分API更改现有索引的分片数量，建议在创建索引时就考虑好分片和副本的数量。

默认情况下，ElasticSearch中的每个索引都分配一个主分片和一个副本，这意味着如果集群中至少有两个节点，则索引将有一个主分片和另一个副本分片（一个完整副本），每个索引总共有两个分片。

> 
> 
> 
> 每个ElasticSearch分片都是一个Lucene索引。在一个Lucene索引中，可以有最多数量的文档。从Lucene-5843起，限制为2147483519（=integer.max_value-128）个文档。您可以使用
> api监视碎片大小（以后会讲到）。
> 
> 

接下来让我们开始有趣的部分吧...

### 安装 ###

二进制文件可从 [www.slastic.co/downloads]( https://link.juejin.im?target=www.slastic.co%2Fdownloads ) 以及过去发布的所有版本中获得。对于每个版本，Windows、Linux和MacOS以及Linux的DEB和RPM软件包以及Windows的MSI安装软件包都提供了与平台相关的存档版本。

**Linux**

简单起见，我们使用tarb包进行安装

下载ElasticSearch 7.1.1 Linux tar。

` curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.1.1-linux-x86_64.tar.gz 复制代码`

解压

` tar -xvf elasticsearch-7.1.1-linux-x86_64.tar.gz 复制代码`

解压完成将在当前目录中创建一组文件和文件夹，然后我们进入bin目录。

` cd elasticsearch-7.1.1/bin 复制代码`

启动节点和单个集群

`./elasticsearch 复制代码`

**Windows**

对于Windows用户，我们建议使用msi安装程序包。该包包含一个图形用户界面（GUI），指导您完成安装过程。

[下载地址]( https://link.juejin.im?target=https%3A%2F%2Fartifacts.elastic.co%2Fdownloads%2Felasticsearch%2Felasticsearch-7.1.1.msi. )

然后双击下载的文件以启动GUI。在第一个屏幕中，选择安装目录：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c8b42c06b6c?imageView2/0/w/1280/h/960/ignore-error/1) 然后选择是作为服务安装，还是根据需要手动启动ElasticSearch。要与Linux示例保持一致，请选择不作为服务安装：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c8b5c68d99c?imageView2/0/w/1280/h/960/ignore-error/1) 对于配置，只需保留默认值：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c8b43542926?imageView2/0/w/1280/h/960/ignore-error/1)

取消选中所有插件以不安装任何插件：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c961db730e8?imageView2/0/w/1280/h/960/ignore-error/1) 单击“安装”按钮后，将安装ElasticSearch：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c8b430168e3?imageView2/0/w/1280/h/960/ignore-error/1) 默认情况下，ElasticSearch将安装在%ProgramFiles%\Elastic\ElasticSearch。进入安装目录，打开命令提示符，输入

`.\elasticsearch.exe 复制代码`

**成功运行节点**

如果安装一切顺利，您将看到下面的一堆消息：

` [2018-09-13T12:20:01,766][INFO ][o.e.e.NodeEnvironment ] [localhost.localdomain] using [1] data paths, mounts [[/home (/dev/mapper/fedora-home)]], net usable_space [335.3gb], net total_space [410.3gb], types [ext4][2018-09-13T12:20:01,772][INFO ][o.e.e.NodeEnvironment ] [localhost.localdomain] heap size [990.7mb], compressed ordinary object pointers [ true ][2018-09-13T12:20:01,774][INFO ][o.e.n.Node ] [localhost.localdomain] node name [localhost.localdomain], node ID [B0aEHNagTiWx7SYj -l 4NTw][2018-09-13T12:20:01,775][INFO ][o.e.n.Node ] [localhost.localdomain] version[7.1.1], pid[13030], build[oss/zip/77 fc 20e/2018-09-13T15:37:57.478402Z], OS[Linux/4.16.11-100.fc26.x86_64/amd64], JVM[ "Oracle Corporation" /OpenJDK 64-Bit Server VM/10/10+46][2018-09-13T12:20:01,775][INFO ][o.e.n.Node ] [localhost.localdomain] JVM arguments [-Xms1g, -Xmx1g, -XX:+UseConcMarkSweepGC, -XX:CMSInitiatingOccupancyFraction=75, -XX:+UseCMSInitiatingOccupancyOnly, -XX:+AlwaysPreTouch, -Xss1m, -Djava.awt.headless= true , -Dfile.encoding=UTF-8, -Djna.nosys= true , -XX:-OmitStackTraceInFastThrow, -Dio.netty.noUnsafe= true , -Dio.netty.noKeySetOptimization= true , -Dio.netty.recycler.maxCapacityPerThread=0, -D log 4j.shutdownHookEnabled= false , -D log 4j2.disable.jmx= true , -Djava.io.tmpdir=/tmp/elasticsearch.LN1ctLCi, -XX:+HeapDumpOnOutOfMemoryError, -XX:HeapDumpPath=data, -XX:ErrorFile=logs/hs_err_pid%p.log, -X log :gc*,gc+age=trace,safepoint:file=logs/gc.log:utctime,pid,tags:filecount=32,filesize=64m, -Djava.locale.providers=COMPAT, -XX:UseAVX=2, -Dio.netty.allocator.type=unpooled, -Des.path.home=/home/manybubbles/Workspaces/Elastic/master/elasticsearch/qa/unconfigured-node-name/build/cluster/integTestCluster node0/elasticsearch-7.0.0-alpha1-SNAPSHOT, -Des.path.conf=/home/manybubbles/Workspaces/Elastic/master/elasticsearch/qa/unconfigured-node-name/build/cluster/integTestCluster node0/elasticsearch-7.0.0-alpha1-SNAPSHOT/config, -Des.distribution.flavor=oss, -Des.distribution.type=zip][2018-09-13T12:20:02,543][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [aggs-matrix-stats][2018-09-13T12:20:02,543][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [analysis-common][2018-09-13T12:20:02,543][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [ingest-common][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [lang-expression][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [lang-mustache][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [lang-painless][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [mapper-extras][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [parent-join][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [percolator][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [rank-eval][2018-09-13T12:20:02,544][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [reindex][2018-09-13T12:20:02,545][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [repository-url][2018-09-13T12:20:02,545][INFO ][o.e.p.PluginsService ] [localhost.localdomain] loaded module [transport-netty4][2018-09-13T12:20:02,545][INFO ][o.e.p.PluginsService ] [localhost.localdomain] no plugins loaded[2018-09-13T12:20:04,657][INFO ][o.e.d.DiscoveryModule ] [localhost.localdomain] using discovery type [zen][2018-09-13T12:20:05,006][INFO ][o.e.n.Node ] [localhost.localdomain] initialized[2018-09-13T12:20:05,007][INFO ][o.e.n.Node ] [localhost.localdomain] starting ...[2018-09-13T12:20:05,202][INFO ][o.e.t.TransportService ] [localhost.localdomain] publish_address {127.0.0.1:9300}, bound_addresses {[::1]:9300}, {127.0.0.1:9300}[2018-09-13T12:20:05,221][WARN ][o.e.b.BootstrapChecks ] [localhost.localdomain] max file descriptors [4096] for elasticsearch process is too low, increase to at least [65535][2018-09-13T12:20:05,221][WARN ][o.e.b.BootstrapChecks ] [localhost.localdomain] max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144][2018-09-13T12:20:08,355][INFO ][o.e.c.s.MasterService ] [localhost.localdomain] elected-as-master ([0] nodes joined)[, ], reason: master node changed {previous [], current [{localhost.localdomain}{B0aEHNagTiWx7SYj -l 4NTw}{hzsQz6CVQMCTpMCVLM4IHg}{127.0.0.1}{127.0.0.1:9300}{testattr= test }]}[2018-09-13T12:20:08,360][INFO ][o.e.c.s.ClusterApplierService] [localhost.localdomain] master node changed {previous [], current [{localhost.localdomain}{B0aEHNagTiWx7SYj -l 4NTw}{hzsQz6CVQMCTpMCVLM4IHg}{127.0.0.1}{127.0.0.1:9300}{testattr= test }]}, reason: apply cluster state (from master [master {localhost.localdomain}{B0aEHNagTiWx7SYj -l 4NTw}{hzsQz6CVQMCTpMCVLM4IHg}{127.0.0.1}{127.0.0.1:9300}{testattr= test } committed version [1] source [elected-as-master ([0] nodes joined)[, ]]])[2018-09-13T12:20:08,384][INFO ][o.e.h.n.Netty4HttpServerTransport] [localhost.localdomain] publish_address {127.0.0.1:9200}, bound_addresses {[::1]:9200}, {127.0.0.1:9200}[2018-09-13T12:20:08,384][INFO ][o.e.n.Node ] [localhost.localdomain] started 复制代码`

我们可以看到名为“6-bjhwl”的节点(在您的示例中是一组不同的字符)已经启动，并将自己选为单个集群中的主机。现在还不用担心master是什么意思。这里最重要的是，我们已经在一个集群中启动了一个节点。

如前所述，我们可以覆盖集群或节点名。当启动ElasticSearch时，可以从命令行执行此操作，如下所示：

`./elasticsearch -Ecluster.name=my_cluster_name -Enode.name=my_node_name 复制代码`

还请注意标记为http的行，其中包含可从中访问节点的http地址（192.168.8.112）和端口（9200）的信息。默认情况下，ElasticSearch使用端口9200提供对其RESTAPI的访问。如有必要，此端口可配置。

### 浏览集群 ###

**使用REST API**

现在我们已经启动并运行了节点（和集群），下一步就是了解如何与之通信。幸运的是，ElasticSearch提供了一个非常全面和强大的RESTAPI，您可以使用它与集群进行交互。可以使用API执行的少数操作如下：

* 

检查集群、节点和索引的运行状况、状态和统计信息。

* 

管理集群、节点和索引数据和元数据。

* 

对索引执行CRUD（创建、读取、更新和删除）和搜索操作。

* 

执行高级搜索操作，如分页、排序、筛选、脚本编写、聚合和许多其他操作。

#### 集群健康 ####

让我们从一个基本的健康检查开始，我们可以使用它来查看集群的运行情况。我们将使用curl来实现这一点，但您可以使用任何允许您进行HTTP/REST调用的工具。假设我们仍然在启动ElasticSearch并打开另一个命令shell窗口的同一个节点上。

为了检查集群的运行状况，我们将使用 ` _cat` API。您可以在Kibana的控制台中运行下面的命令，方法是单击“在控制台中查看”，或者使用curl，方法是单击下面的“复制为curl”链接并将其粘贴到终端中。

` curl -X GET "localhost:9200/_cat/health?v" 复制代码`

响应结果：

` epoch timestamp cluster status node.total node.data shards pri relo init unassign pending_tasks max_task_wait_time active_shards_percent1475247709 17:01:49 elasticsearch green 1 1 0 0 0 0 0 0 - 100.0%​ 复制代码`

我们可以看到名为“elasticsearch”的集群处于绿色状态。每当我们请求集群健康时，我们要么得到绿色、黄色，要么得到红色。

* 

绿色-一切正常（集群功能齐全）

* 

黄色-所有数据都可用，但某些副本尚未分配（群集完全正常工作）

* 

红色-由于任何原因，某些数据不可用（群集部分正常工作）

> 
> 
> 
> 注意：当集群为红色时，它将继续提供来自可用分片的搜索请求，但您可能需要尽快修复它，因为存在未分配的分片。
> 
> 

从上面的响应中，我们可以看到总共1个节点，并且我们有0个碎片，因为我们在其中还没有数据。请注意，由于我们使用的是默认群集名称（ElasticSearch），并且由于ElasticSearch默认情况下发现在同一台计算机上查找其他节点，因此您可能会意外启动计算机上的多个节点，并使它们都加入单个群集。在这个场景中，您可能会在上面的响应中看到多个节点。

我们还可以得到集群中的节点列表:

` curl -X GET "localhost:9200/_cat/nodes?v" 复制代码`

响应结果：

` ip heap.percent ram.percent cpu load_1m load_5m load_15m node.role master name127.0.0.1 10 5 5 4.46 mdi * PB2SGZY 复制代码`

我们可以看到一个名为“pb2sgzy”的节点，它是当前集群中的单个节点。

#### 查看所有索引 ####

现在让我们来看看我们的索引：

` curl -X GET "localhost:9200/_cat/indices?v" 复制代码`

响应结果：

` health status index uuid pri rep docs.count docs.deleted store.size pri.store.size 复制代码`

这就意味着我们在集群中还没有索引。

#### 创建索引 ####

现在，让我们创建一个名为“customer”的索引，然后再次列出所有索引：

` curl -X PUT "localhost:9200/customer?pretty" curl -X GET "localhost:9200/_cat/indices?v" 复制代码`

第一个命令使用put动词创建名为“customer”的索引。我们只需在调用的末尾附加 ` pretty` 命令它漂亮地打印JSON响应（如果有的话）。

响应结果：

` health status index uuid pri rep docs.count docs.deleted store.size pri.store.sizeyellow open customer 95SQ4TSUT7mWBT7VNHH67A 1 1 0 0 260b 260b 复制代码`

第二个命令的结果告诉我们，我们现在有一个名为customer的索引，它有一个主碎片和一个副本（默认值），其中包含零个文档。

您可能还会注意到客户索引中有一个黄色的健康标签。回想我们之前的讨论，黄色意味着一些副本尚未分配。此索引发生这种情况的原因是，默认情况下，ElasticSearch为此索引创建了一个副本。因为目前只有一个节点在运行，所以在另一个节点加入集群之前，还不能分配一个副本（为了高可用性）。一旦该副本分配到第二个节点上，该索引的运行状况将变为绿色。

#### 查询文档 ####

现在我们把一些东西放到客户索引中。我们将在客户索引中索引一个简单的客户文档，其ID为1，如下所示：

` curl -X PUT "localhost:9200/customer/_doc/1?pretty" -H 'Content-Type: application/json' -d '{ "name": "John Doe"}' ​ 复制代码`

响应结果：

` { "_index" : "customer" , "_type" : "_doc" , "_id" : "1" , "_version" : 1, "result" : "created" , "_shards" : { "total" : 2, "successful" : 1, "failed" : 0 }, "_seq_no" : 0, "_primary_term" : 1} 复制代码`

从上面，我们可以看到在客户索引中成功地创建了一个新的客户文档。文档还有一个内部ID 1，我们在索引时指定了它。

需要注意的是，ElasticSearch不要求您在索引文档之前先显式创建索引。在上一个示例中，如果客户索引之前不存在，那么ElasticSearch将自动创建该索引。

现在让我们检索刚才索引的文档：

` curl -X GET "localhost:9200/customer/_doc/1?pretty" 复制代码`

响应结果：

` { "_index" : "customer" , "_type" : "_doc" , "_id" : "1" , "_version" : 1, "_seq_no" : 25, "_primary_term" : 1, "found" : true , "_source" : { "name" : "John Doe" }} 复制代码`

除了一个字段之外 ` found` ，这里没有发现任何异常的地方，说明我们找到了一个具有请求的ID 1的文档和另一个字段 ` _source` ，它返回了我们从上一步索引的完整JSON文档。

#### 删除索引 ####

现在，让我们删除刚刚创建的索引，然后再次列出所有索引：

` curl -X DELETE "localhost:9200/customer?pretty" curl -X GET "localhost:9200/_cat/indices?v" 复制代码`

响应结果：

` health status index uuid pri rep docs.count docs.deleted store.size pri.store.size 复制代码`

这意味着索引被成功地删除了，现在我们又回到了开始时集群中什么都没有的地方。

在我们继续之前，让我们再仔细看看我们迄今为止学到的一些API命令：

` #创建索引curl -X PUT "localhost:9200/customer"#创建文档（添加数据）curl -X PUT "localhost:9200/customer/_doc/1" -H 'Content-Type: application/json' -d'{ "name": "John Doe"}'#查询文档（查询数据）curl -X GET "localhost:9200/customer/_doc/1"#删除文档（删除数据）curl -X DELETE "localhost:9200/customer" 复制代码`

如果我们仔细研究上述命令，我们实际上可以看到在ElasticSearch中如何访问数据的模式。这种模式可以概括如下：

` <HTTP Verb> /<Index>/<Endpoint>/<ID> 复制代码`

这种REST访问模式在所有API命令中都非常普遍，如果您能记住它，那么您将在掌握ElasticSearch方面有一个很好的开端。

### 修改数据 ###

ElasticSearch提供近实时的数据操作和搜索功能。默认情况下，从索引/更新/删除数据到数据出现在搜索结果中，预计一秒钟的延迟（刷新间隔）。这是与其他平台（如SQL）的一个重要区别，在SQL中，数据在事务完成后立即可用。

**创建/替换文档（修改数据）**

我们以前见过如何索引单个文档。让我们再次回忆一下这个命令：

` curl -X PUT "localhost:9200/customer/_doc/1?pretty" -H 'Content-Type: application/json' -d '{ "name": "John Doe"}' 复制代码`

同样，上面将把指定的文档索引到客户索引中，ID为1。如果我们用不同的（或相同的）文档再次执行上述命令，那么ElasticSearch将在现有文档的基础上替换（即重新创建）一个ID为1的新文档：

` curl -X PUT "localhost:9200/customer/_doc/1?pretty" -H 'Content-Type: application/json' -d '{ "name": "Jane Doe"}' 复制代码`

上面将ID为1的文档的名称从“John Doe”更改为“Jane Doe”。另一方面，如果我们使用不同的ID，则会创建新的文档，并且索引中已有的文档将保持不变。

上面的索引是一个ID为2的新文档。

创建时，ID是可填可不填的。如果未指定，ElasticSearch将生成一个随机ID。ElasticSearch生成的实际ID（或在前面的示例中显式指定的任何内容）作为索引API调用的一部分返回。

此示例演示如何索引没有显式ID的文档：

` curl -X POST "localhost:9200/customer/_doc?pretty" -H 'Content-Type: application/json' -d '{ "name": "Jane Doe"}' 复制代码`

注意，在上述情况下，我们使用的是 ` post` 动词而不是 ` put` ，因为我们没有指定ID。

#### 更新数据 ####

除了能够添加和替换文档之外，我们也可以更新文档。请注意，Elasticsearch实际上并没有在底层执行覆盖更新。而是先删除旧文档，再添加一条新文档。

这个例子把原来ID为1的名字修改成了Jane Doe，详情请看下面的例子：

` curl -X POST "localhost:9200/customer/_update/1?pretty" -H 'Content-Type: application/json' -d '{ "doc": { "name": "Jane Doe" }}' 复制代码`

此示例演示如何通过将名称字段更改为“Jane Doe”来更新以前的文档（ID为1），同时向其添加年龄字段：

` curl -X POST "localhost:9200/customer/_update/1?pretty" -H 'Content-Type: application/json' -d '{ "doc": { "name": "Jane Doe", "age": 20 }}' 复制代码`

也可以使用简单的脚本执行更新,此示例使用脚本将年龄增加5：

` curl -X POST "localhost:9200/customer/_update/1?pretty" -H 'Content-Type: application/json' -d '{ "script" : "ctx._source.age += 5"}' 复制代码`

Elasticsearch提供了给定条件下更新多个文档的功能，就像Sql中的 ` updata ... where ...` 后面的章节我们会详细介绍。

#### 删除数据 ####

删除文档相当简单。

此示例显示如何删除ID为2的以前的客户：

` curl -X DELETE "localhost:9200/customer/_doc/2?pretty" 复制代码`

请参阅_delete_by_query API来删除与特定查询匹配的所有文档。值得注意的是，删除整个索引比使用delete By Query API删除所有文档要有效得多。 ` _delete_by_query API` 会在后面详细介绍。

#### 批处理 ####

除了能够索引、更新和删除单个文档之外，Elasticsearch还提供了使用_bulk API批量执行上述操作的能力。此功能非常重要，因为它提供了一种非常有效的机制，可以在尽可能少的网络往返的情况下尽可能快地执行多个操作。

作为一个简单的例子，下面的调用在一个批量操作中索引两个文档(ID 1 - John Doe和ID 2 - Jane Doe):

` curl -X POST "localhost:9200/customer/_bulk?pretty" -H 'Content-Type: application/json' -d '{"index":{"_id":"1"}}{"name": "John Doe" }{"index":{"_id":"2"}}{"name": "Jane Doe" }' 复制代码`

此例更新第一个文档(ID为1)，然后在一次批量操作中删除第二个文档(ID为2):

` curl -X POST "localhost:9200/customer/_bulk?pretty" -H 'Content-Type: application/json' -d '{"update":{"_id":"1"}}{"doc": { "name": "John Doe becomes Jane Doe" } }{"delete":{"_id":"2"}}' 复制代码`

请注意，对于delete操作，它之后没有对应的源文档，因为删除操作只需要删除文档的ID。

批量API不会因为某个操作失败而失败(有错误也会执行下去，最后会返回每个操作的状态)。如果一个操作由于某种原因失败，它将在失败之后继续处理其余的操作。当bulk API返回时，它将为每个操作提供一个状态(与发送操作的顺序相同)，以便检查某个特定操作是否失败。

### 浏览数据 ###

**样本数据**

现在我们已经了解了基本知识，让我们尝试使用更真实的数据集。我准备了一个虚构的JSON客户银行账户信息文档示例。每个文档都有以下内容:

` { "account_number" : 0, "balance" : 16623, "firstname" : "Bradshaw" , "lastname" : "Mckenzie" , "age" : 29, "gender" : "F" , "address" : "244 Columbus Place" , "employer" : "Euron" , "email" : "bradshawmckenzie@euron.com" , "city" : "Hobucken" , "state" : "CO" } 复制代码`

该数据是使用www.json- generator.com/生成的，因此请忽略数据的实际值和语义，因为它们都是随机生成的。

**加载样本数据**

您可以从这里下载示例数据集（accounts.json）。将其提取到当前目录，然后按如下方式将其加载到集群中：

` curl -H "Content-Type: application/json" -XPOST "localhost:9200/bank/_bulk?pretty&refresh" --data-binary "@accounts.json" curl "localhost:9200/_cat/indices?v" 复制代码`

响应结果：

` health status index uuid pri rep docs.count docs.deleted store.size pri.store.sizeyellow open bank l7sSYV2cQXmu6_4rJWVIww 5 1 1000 0 128.6kb 128.6kb 复制代码`

这意味着我们刚刚成功地将索引的1000个文档批量存储到银行索引中。

#### 搜索API ####

现在让我们从一些简单的搜索开始。运行搜索有两种基本方法:

* 

一种是通过REST请求URI发送搜索参数。

* 

另一种是通过REST请求体发送搜索参数。

请求体方法允许您更富表现力，还可以以更可读的JSON格式定义搜索。我们将尝试请求URI方法的一个示例，但是在本教程的其余部分中，我们将只使用请求体方法。

用于搜索的REST API可以从 ` _search` 端点访问。这个例子返回银行索引中的所有文档:

` curl -X GET "localhost:9200/bank/_search?q=*&sort=account_number:asc&pretty" 复制代码`

让我们首先分析一下搜索调用。我们正在银行索引中搜索( ` _search` )， ` q=*` 参数指示Elasticsearch匹配索引中的所有文档。 ` sort=account_number:asc` 参数指示使用每个文档的account_number字段按升序对结果排序。同样， ` pretty` 参数只告诉Elasticsearch返回打印得很漂亮的JSON结果。

相应结果(部分显示)：

` { "took" : 63, "timed_out" : false , "_shards" : { "total" : 5, "successful" : 5, "skipped" : 0, "failed" : 0 }, "hits" : { "total" : { "value" : 1000, "relation" : "eq" }, "max_score" : null, "hits" : [ { "_index" : "bank" , "_type" : "_doc" , "_id" : "0" , "sort" : [0], "_score" : null, "_source" : { "account_number" :0, "balance" :16623, "firstname" : "Bradshaw" , "lastname" : "Mckenzie" , "age" :29, "gender" : "F" , "address" : "244 Columbus Place" , "employer" : "Euron" , "email" : "bradshawmckenzie@euron.com" , "city" : "Hobucken" , "state" : "CO" } }, { "_index" : "bank" , "_type" : "_doc" , "_id" : "1" , "sort" : [1], "_score" : null, "_source" : { "account_number" :1, "balance" :39225, "firstname" : "Amber" , "lastname" : "Duke" , "age" :32, "gender" : "M" , "address" : "880 Holmes Lane" , "employer" : "Pyrami" , "email" : "amberduke@pyrami.com" , "city" : "Brogan" , "state" : "IL" } }, ... ] }} 复制代码`

关于回应，我们可以看到:

* 

` took` Elasticsearch执行搜索所用的时间(以毫秒为单位)

* 

` timed_out` 告诉我们搜索是否超时

* 

` _shards` 告诉我们搜索了多少碎片，以及成功/失败搜索碎片的计数

* 

` hits` 搜索结果

* 

` hits.total` 包含与搜索条件匹配的文档总数相关的信息的对象

* 

` hits.total.value` 总命中数的值。

* 

` hits.total.relation` : ` hits.total.value` 值是准确的命中次数，在这种情况下它等于 ` eq` 或总命中次数的下界(大于或等于)，在这种情况下它等于 ` gte` 。

* 

` hits.hits` 实际的搜索结果数组(默认为前10个文档)

* 

` hits.sort` 结果排序键(如果按分数排序，则丢失)

* 

` hits._score` 和 ` max_score` ——暂时忽略这些字段

` hits.total` 的准确度是由请求参数 ` track_total_hits` 控制，当 ` track_total_hits` 设置为true时，请求将精确地跟踪总命中 ` “relation”:“eq”` 。默认值为10,000，这意味着总命中数可以精确地跟踪到10,000个文档。通过显式地将 ` track_total_hits` 设置为true，可以强制进行准确的计数。有关详细信息,后面章节我们会进行介绍。

这是使用请求体搜索的方式:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_all": {} }, "sort": [ { "account_number": "asc" } ]}' 复制代码`

这里的不同之处在于，我们没有在URI中传递q=*，而是向_search API提供了json风格的查询请求体。我们将在下一节中讨论这个JSON查询。

重要的是要了解，一旦您获得了搜索结果，Elasticsearch就会完全处理请求，并且不会维护任何类型的服务器端资源，也不会在结果中打开游标。这是许多其他平台如SQL形成鲜明对比,你最初可能得到部分的子集查询结果预先然后不断返回到服务器,如果你想获取(或页面)其余的结果使用某种状态的服务器端游标。

#### 引入查询语言 ####

Elasticsearch提供了一种JSON风格的查询语言，您可以使用它来执行查询。这称为Query DSL。查询语言非常全面，乍一看可能有些吓人，但实际上学习它的最佳方法是从几个基本示例开始。

回到上一个例子，我们执行了这个查询：

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_all": {} }}' 复制代码`

仔细分析上面的内容， ` query` 部分告诉我们进行查询操作， ` match_all` 只是我们想要运行的查询类型，match_all只是搜索指定索引中的所有文档。

除了查询参数，我们还可以传递其他参数来影响搜索结果。在上一小节的最后，我们传入了sort，这里我们传入了size：

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_all": {} }, "size": 1}' 复制代码`
> 
> 
> 
> 
> 请注意，如果未指定大小，则默认为10。
> 
> 

下面的例子执行 ` match_all` 并返回文档10到19（from和size可以类比mysql中的limit ？ ？）：

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_all": {} }, "from": 10, "size": 10}' 复制代码`

from参数(基于0)指定从哪个文档索引开始，size参数指定从from参数开始返回多少文档。该特性在实现搜索结果分页时非常有用。

> 
> 
> 
> 注意，如果没有指定from，则默认值为0。
> 
> 

本这个例子执行 ` match_all` 操作，并按帐户余额降序对结果进行排序，并返回前10个(默认大小)文档。

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_all": {} }, "sort": { "balance": { "order": "desc" } }}' 复制代码`

#### 执行搜索 ####

既然我们已经看到了一些基本的搜索参数，那么让我们深入研究一下Query DSL。让我们先看看返回的文档字段。默认情况下，作为所有搜索的一部分,返回完整的JSON文档。这被称为“源”（搜索命中中的 ` _source` 字段）。如果我们不希望返回整个源文档，我们可以只请求从源中返回几个字段。

此示例显示如何从搜索中返回两个字段，即帐号和余额（在 ` !source` 内）：

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_all": {} }, "_source": ["account_number", "balance"]}' 复制代码`

注意，上面的示例只减少了 ` _source` _字段。它仍然只返回一个名为 ` _source` 的字段，但是其中只包含account_number和balance字段。

如果您之前有了解过MySql，那么上面的内容在概念上与SQL SELECT from field list有些类似。

现在让我们进入查询部分。在前面，我们已经了解了如何使用 ` match_all` 查询来匹配所有文档。现在让我们引入一个名为 ` match` 查询的新查询，它可以被看作是基本的字段搜索查询(即针对特定字段或字段集进行的搜索)。

本例返回编号为20的帐户

> 
> 
> 
> 类比mysql match类似于mysql 中的条件查询。
> 
> 

例如返回编号为20的帐户:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match": { "account_number": 20 } }}' 复制代码`

此示例返回地址中包含“mill”的所有帐户：

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match": { "address": "mill" } }}' 复制代码`

此示例返回地址中包含“mill”或“lane”的所有帐户：

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match": { "address": "mill lane" } }}' 复制代码`

**match (match_phrase)的一个变体** ，它返回地址中包含短语“mill lane”的所有帐户:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "match_phrase": { "address": "mill lane" } }}' 复制代码`
> 
> 
> 
> 
> 注意：match 中如果加空格，那么会被认为两个单词，包含任意一个单词将被查询到
> 
> 
> 
> match_parase 将忽略空格，将该字符认为一个整体，会在索引中匹配包含这个整体的文档。
> 
> 

现在让我们介绍bool查询。bool查询允许我们使用布尔逻辑将较小的查询组合成较大的查询。

> 
> 
> 
> 如果您熟悉mysql，那么你就会发现布尔查询其实相当于 and or not...
> 
> 

这个例子包含两个匹配查询，返回地址中包含“mill” **和** “lane”的所有帐户:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "bool": { "must": [ { "match": { "address": "mill" } }, { "match": { "address": "lane" } } ] } }}' 复制代码`

在上面的示例中，bool must子句指定了所有必须为true的查询，则将文档视为匹配。

相反，这个例子包含两个匹配查询，并返回地址中包含“mill” **或** “lane”的所有帐户:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "bool": { "should": [ { "match": { "address": "mill" } }, { "match": { "address": "lane" } } ] } }}' 复制代码`

在上面的示例中，bool should子句指定了一个查询列表，其中任何一个查询必须为真，才能将文档视为匹配。

这个例子包含两个匹配查询，返回地址中既不包含“mill” **也不包含** “lane”的所有帐户:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "bool": { "must_not": [ { "match": { "address": "mill" } }, { "match": { "address": "lane" } } ] } }}' 复制代码`

在上面的例子中，bool must_not子句指定了一个查询列表，其中没有一个查询必须为真，才能将文档视为匹配。

我们可以在bool查询中同时组合must、should和must_not子句。此外，我们可以在这些bool子句中组合bool查询，以模拟任何复杂的多级布尔逻辑。

这个例子返回所有40岁但不居住在ID(aho)的人的账户:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "bool": { "must": [ { "match": { "age": "40" } } ], "must_not": [ { "match": { "state": "ID" } } ] } }}' 复制代码`

#### 执行过滤器 ####

在上一节中，我们跳过了一个名为document score(搜索结果中的_score字段)的小细节。分数是一个数值，它是衡量文档与我们指定的搜索查询匹配程度的一个相对指标。分数越高，文档越相关，分数越低，文档越不相关。

但是查询并不总是需要生成分数，特别是当它们只用于“过滤”文档集时。Elasticsearch会自动优化查询执行，以避免计算无用的分数。

我们在上一节中介绍的bool查询还支持filter子句，它允许我们使用查询来限制将由其他子句匹配的文档，而不改变计算分数的方式。作为一个示例，让我们介绍范围查询，它允许我们根据一系列值筛选文档。这通常用于数字或日期筛选。

本例使用bool查询返回余额在20000到30000之间的所有帐户，包括余额。换句话说，我们希望找到余额大于或等于20000，小于或等于30000的账户。

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "query": { "bool": { "must": { "match_all": {} }, "filter": { "range": { "balance": { "gte": 20000, "lte": 30000 } } } } }}' 复制代码`

通过分析上面的内容，bool查询包含一个 ` match_all` 查询(查询部分)和一个 ` range` 查询(筛选部分)。我们可以将任何其他查询替换到查询和筛选器部分中。在上面的例子中，范围查询非常有意义，因为落入范围的文档都“相等”匹配，没有文档比有更有意义（因为是筛选过的）。

除了 ` match_all、match、bool和range` 查询之外，还有许多其他可用的查询类型，我们在这里不深入讨论它们。既然我们已经对它们的工作原理有了基本的了解，那么在学习和试验其他查询类型时应用这些知识应该不会太难。

#### 执行聚合（类比mysql 聚合函数） ####

聚合提供了对数据进行分组和提取统计信息的能力。考虑聚合最简单的方法是大致将其等同于SQL GROUP by和SQL聚合函数。在Elasticsearch中，您可以执行返回命中的搜索，同时在一个响应中返回与所有命中分离的聚合结果。这是非常强大和高效的，因为您可以运行查询和多个聚合，并一次性获得这两个(或任何一个)操作的结果，从而避免使用简洁和简化的API进行网络往返。

首先，这个示例按状态对所有帐户进行分组，然后返回按count降序排列的前10个(默认)状态(也是默认):

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "size": 0, "aggs": { "group_by_state": { "terms": { "field": "state.keyword" } } }}' 复制代码`

在SQL中，上述聚合在概念上与:

` SELECT state, COUNT(*) FROM bank GROUP BY state ORDER BY COUNT(*) DESC LIMIT 10; 复制代码` ` { "took" : 29, "timed_out" : false , "_shards" : { "total" : 5, "successful" : 5, "skipped" : 0, "failed" : 0 }, "hits" : { "total" : { "value" : 1000, "relation" : "eq" }, "max_score" : null, "hits" : [ ] }, "aggregations" : { "group_by_state" : { "doc_count_error_upper_bound" : 20, "sum_other_doc_count" : 770, "buckets" : [ { "key" : "ID" , "doc_count" : 27 }, { "key" : "TX" , "doc_count" : 27 }, { "key" : "AL" , "doc_count" : 25 }, { "key" : "MD" , "doc_count" : 25 }, { "key" : "TN" , "doc_count" : 23 }, { "key" : "MA" , "doc_count" : 21 }, { "key" : "NC" , "doc_count" : 21 }, { "key" : "ND" , "doc_count" : 21 }, { "key" : "ME" , "doc_count" : 20 }, { "key" : "MO" , "doc_count" : 20 } ] } }} 复制代码`

我们可以看到ID(Idaho)有27个帐户，其次是TX(Texas)的27个帐户，然后是AL(Alabama)的25个帐户，依此类推。

> 
> 
> 
> 注意，我们将size=0设置为不显示搜索结果，因为我们只想看到响应中的聚合结果。
> 
> 

基于前面的汇总，本示例按状态计算平均帐户余额(同样只针对按count降序排列的前10个状态):

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "size": 0, "aggs": { "group_by_state": { "terms": { "field": "state.keyword" }, "aggs": { "average_balance": { "avg": { "field": "balance" } } } } }}' 复制代码`

注意，我们如何将average_balance聚合嵌套在group_by_state聚合中。这是所有聚合的常见模式。您可以在聚合中任意嵌套聚合，以从数据中提取所需的结果。

基于之前的聚合，我们现在按降序对平均余额排序:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "size": 0, "aggs": { "group_by_state": { "terms": { "field": "state.keyword", "order": { "average_balance": "desc" } }, "aggs": { "average_balance": { "avg": { "field": "balance" } } } } }}' 复制代码`

这个例子展示了我们如何按照年龄等级(20-29岁，30-39岁，40-49岁)分组，然后按性别分组，最后得到每个年龄等级，每个性别的平均账户余额:

` curl -X GET "localhost:9200/bank/_search" -H 'Content-Type: application/json' -d '{ "size": 0, "aggs": { "group_by_age": { "range": { "field": "age", "ranges": [ { "from": 20, "to": 30 }, { "from": 30, "to": 40 }, { "from": 40, "to": 50 } ] }, "aggs": { "group_by_gender": { "terms": { "field": "gender.keyword" }, "aggs": { "average_balance": { "avg": { "field": "balance" } } } } } } }}' 复制代码`

还有许多其他聚合功能，我们在这里不会详细讨论。如果您想做进一步的实验，那么聚合参考指南是一个很好的开始。

#### 结论 ####

弹性搜索既是一个简单的产品，也是一个复杂的产品。到目前为止，我们已经了解了它是什么、如何查看它的内部以及如何使用一些RESTapi来使用它。希望本教程能让您更好地理解Elasticsearch是什么，更重要的是，它激发了您进一步试验它的其他优秀特性!