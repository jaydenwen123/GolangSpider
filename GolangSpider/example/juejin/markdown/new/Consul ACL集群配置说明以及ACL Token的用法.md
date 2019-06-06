# Consul ACL集群配置说明以及ACL Token的用法 #

在上一篇文章里面，我们讲了如何搭建带有Acl控制的Consul集群。 这一篇文章主要讲述一下上一篇文章那一大串配置文件的含义。

# 1.配置说明 #

# #1.1 勘误 #

上一篇文章关于机器规划方面，consul client agent的端口写的有误。这里再贴一下正确的机器规划。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286eb834ceee3?imageView2/0/w/1280/h/960/ignore-error/1)

## 1.2 我们先来看一下consul server agent的配置。 ##

上一节中，提供了三个配置文件，consul-server1.json, consul-server2.json以及consul-server3.json。 其中consul-server1.json参数最多，这里就以它来说明各个配置的含义：

` { "datacenter" : "dc1" , "primary_datacenter" : "dc1" , "bootstrap_expect" :1, "start_join" :[ "10.211.55.25" , "10.211.55.26" ], "retry_join" :[ "10.211.55.25" , "10.211.55.26" ], "advertise_addr" : "10.211.55.28" , "bind_addr" : "10.211.55.28" , "server" : true , "connect" :{ "enabled" : true }, "node_name" : "consul-server1" , "data_dir" : "/opt/consul/data/" , "enable_script_checks" : false , "enable_local_script_checks" : true , "log_file" : "/opt/consul/log/" , "log_level" : "info" , "log_rotate_bytes" :100000000, "log_rotate_duration" : "24h" , "encrypt" : "krCysDJnrQ8dtA7AbJav8g==" , "acl" :{ "enabled" : true , "default_policy" : "deny" , "enable_token_persistence" : true , "tokens" :{ "master" : "cd76a0f7-5535-40cc-8696-073462acc6c7" , "agent" : "deaa315d-98c5-b9f6-6519-4c8f6574a551" } } } 复制代码`

* datacenter 此标志表示代理运行的数据中心。如果未提供，则默认为“dc1”。 Consul拥有对多个数据中心的一流支持，但它依赖于正确的配置。同一数据中心中的节点应在同一个局域网内。
* primary_datacenter: 这指定了对ACL信息具有权威性的数据中心。必须提供它才能启用ACL。
* bootstrap_expect: Consul将等待指定数量的服务器可用，然后才会引导群集。这允许自动选择初始领导者。
* start_join: 一个字符串数组，指定是其他的consul server agent的地址。这里这样配置，会在启动时，尝试将consul-server2，consul-server3这两个节点加进来，形成一个集群。
* retry_join: 允许start_join时失败时，继续重新连接。重试的时间间隔，可以用retry_interval设置，默认是30s;重试的最大次数，可以用retry_max设置，默认是0，也就是无限次重试。关于retry_interval和retry_max，这里都是用的默认值。
* bind_addr: 内部群集通信绑定的地址。这是群集中所有其他节点都应该可以访问的IP地址。默认情况下，这是“0.0.0.0”，这意味着Consul将绑定到本地计算机上的所有地址，并将第一个可用的私有IPv4地址通告给群集的其余部分。如果有多个私有IPv4地址可用，Consul将在启动时退出并显示错误。如果指定“[::]”，Consul将通告第一个可用的公共IPv6地址。如果有多个可用的公共IPv6地址，Consul将在启动时退出并显示错误。 Consul同时使用TCP和UDP，并且两者使用相同的端口。如果您有防火墙，请务必同时允许这两种协议。
* advertise_addr: 更改我们向群集中其他节点通告的地址。默认情况下，会使用-bind参数指定的地址.
* server: 是否是server agent节点。
* connect.enabled: 是否启动 [Consul Connect]( https://link.juejin.im?target=https%3A%2F%2Fwww.consul.io%2Fdocs%2Fconnect%2Findex.html ) ，这里是启用的。
* node_name：节点名称。
* data_dir: agent存储状态的目录。
* enable_script_checks： 是否在此代理上启用执行脚本的健康检查。有安全漏洞，默认值就是false，这里单独提示下。
* enable_local_script_checks: 与enable_script_checks类似，但只有在本地配置文件中定义它们时才启用它们。仍然不允许在HTTP API注册中定义的脚本检查。
* log-file: 将所有Consul Agent日志消息重定向到文件。这里指定的是/opt/consul/log/目录。
* log_rotate_bytes：指定在需要轮换之前应写入日志的字节数。除非指定，否则可以写入日志文件的字节数没有限制
* log_rotate_duration：指定在需要旋转日志之前应写入日志的最长持续时间。除非另有说明，否则日志会每天轮换（24小时。单位可以是"ns", "us" (or "µs"), "ms", "s", "m", "h"， 比如设置值为24h
* encrypt：用于加密Consul Gossip 协议交换的数据。在启动各个server之前，配置成同一个UUID值就行，或者你用命令行consul keygen 命令来生成也可以。
* acl.enabled: 是否启用acl.
* acl.default_policy: “allow”或“deny”; 默认为“allow”，但这将在未来的主要版本中更改。当没有匹配规则时，默认策略控制令牌的行为。在“allow”模式下，ACL是黑名单：允许任何未明确禁止的操作。在“deny”模式下，ACL是白名单：阻止任何未明确允许的操作.
* acl.enable_token_persistence: 可能值为true或者false。值为true时，API使用的令牌集合将被保存到磁盘，并且当代理重新启动时会重新加载。
* acl.tokens.master: 具有全局管理的权限，也就是最大的权限。它允许操作员使用众所周知的令牌密钥ID来引导ACL系统。需要在所有的server agent上设置同一个值，可以设置为一个随机的UUID。这个值权限最大，注意保管好。
* acl.tokens.agent: 用于客户端和服务器执行内部操作.比如catalog api的更新，反熵同步等。

## 1.3 再来说下consul-client1的相关配置。 ##

我再贴一下配置信息。

` { "datacenter" : "dc1" , "primary_datacenter" : "dc1" , "advertise_addr" : "10.211.55.27" , "start_join" :[ "10.211.55.25" , "10.211.55.26" , "10.211.55.28" ], "retry_join" :[ "10.211.55.25" , "10.211.55.26" , "10.211.55.28" ], "bind_addr" : "10.211.55.27" , "node_name" : "consul-client1" , "client_addr" : "0.0.0.0" , "connect" :{ "enabled" : true }, "data_dir" : "/opt/consul/data/" , "log_file" : "/opt/consul/log/" , "log_level" : "info" , "log_rotate_bytes" :100000000, "log_rotate_duration" : "24h" , "encrypt" : "krCysDJnrQ8dtA7AbJav8g==" , "ui" : true , "enable_script_checks" : false , "enable_local_script_checks" : true , "disable_remote_exec" : true , "ports" :{ "http" :7110 }, "acl" :{ "enabled" : true , "default_policy" : "deny" , "enable_token_persistence" : true , "tokens" :{ "agent" : "deaa315d-98c5-b9f6-6519-4c8f6574a551" } } } 复制代码`

这里，start_join, retry_join都是指定的server agent的地址。 另外还没有提过的配置就是client_addr， ui, ports.http . 下面依次说明：

* client_addr: Consul将绑定客户端接口的地址，包括HTTP和DNS服务器
* ui: 启用内置Web UI服务器和对应所需的HTTP路由
* ports.http： 更改默认的http端口。

# 2. ACL Token的用法 #

## 2.1 ACL Token 有什么用呢？ ##

可以有人会说，你上面让我又是搭建环境，又是看配置说明，我建好了一个这么一个带ACL控制的Consul集群有什么用呢？

ACL 全称 Access Control List，也就是访问控制列表的意思，现在我们生成了带有ACL控制的集群，就意味不是谁都能来向我注册的，也不是谁都能像我获取服务列表-- 也就是你想对Consul执行任何操作，你得对应的令牌，也就是ACL Token。

## 2.2 不带Token行不行？ ##

为了模拟一般的Http请求，我这里下载一个 [Postman]( https://link.juejin.im?target=https%3A%2F%2Fwww.getpostman.com%2F ) , 是的这里没有用命令行curl。 我们现在postman输入 [http://127.0.0.1:7110/v1/catalog/nodes]( https://link.juejin.im?target=http%3A%2F%2F127.0.0.1%3A7110%2Fv1%2Fcatalog%2Fnodes ) 会发现一个节点都拿不到：

![without token](https://user-gold-cdn.xitu.io/2019/6/6/16b286eb8352e04f?imageView2/0/w/1280/h/960/ignore-error/1) 此时如果加上master token, 也就是访问 [http://127.0.0.1:7110/v1/catalog/nodes?token=cd76a0f7-5535-40cc-8696-073462acc6c7]( https://link.juejin.im?target=http%3A%2F%2F127.0.0.1%3A7110%2Fv1%2Fcatalog%2Fnodes%3Ftoken%3Dcd76a0f7-5535-40cc-8696-073462acc6c7 ) 会发现可以拿到所有节点的数据（下图只截取一部分） ![with token](https://user-gold-cdn.xitu.io/2019/6/6/16b286eb835a4f33?imageView2/0/w/1280/h/960/ignore-error/1)

## 2.3 不带token是不行，那能不能带权限小点的token呢？ ##

前面说过master token是权限最大的token，假如这样给出去，各个部分都拿来用。如果两个不同的部分注册名称一样的服务该怎么办，取消注册了其他部门的服务又该怎么办。总之，权限能不能给小点，答案是可以的。 首先说一下目标： 1.不同部门的服务必须要有自己的前缀，比如deptA表示部门A，比如deptB表示部门B 2.不同部门只能更改自己的服务。

### 2.3.1 先给两个部门各注册一个服务 ###

注册服务deptA-pingbaidu1, 注意这里选择的PUT方法。

` PUT http://127.0.0.1:7110/v1/agent/service/register?token= cd 76a0f7-5535-40cc-8696-073462acc6c7 { "ID" : "deptA-pingbaidu1" , "Name" : "deptA-pingbaidu" , "Tags" : [ "primary" , "v1" ], "Address" : "127.0.0.1" , "Port" : 8000, "Meta" : { "my_version" : "4.0" }, "EnableTagOverride" : false , "Check" : { "DeregisterCriticalServiceAfter" : "90m" , "HTTP" : "http://www.baidu.com/" , "Interval" : "10s" } }&emsp; 复制代码`

在截个图，当返回status为200时，表示成功注册

![注册成功](https://user-gold-cdn.xitu.io/2019/6/6/16b286eb8361170b?imageView2/0/w/1280/h/960/ignore-error/1) 此时可以在consul web ui中进行查看，打开consul-client1所在的机器，在浏览器中，输入http://127.0.0.1:7110/ui/dc1/services，（ 注意在此之前你需要先设置consul web ui的token，上一篇文章末尾已经提及），此时会看到 ![service-on-web-ui](https://user-gold-cdn.xitu.io/2019/6/6/16b286eb8948e6a5?imageView2/0/w/1280/h/960/ignore-error/1) 类似地，在注册个deptA-pingMe1的服务

` PUT http://127.0.0.1:7110/v1/agent/service/register?token= cd 76a0f7-5535-40cc-8696-073462acc6c7 { "ID" : "deptB-pingMe1" , "Name" : "deptB-pingMe" , "Tags" : [ "primary" , "v1" ], "Address" : "127.0.0.1" , "Port" : 7000, "Meta" : { "my_version" : "4.0" }, "EnableTagOverride" : false , "Check" : { "DeregisterCriticalServiceAfter" : "90m" , "HTTP" : "https://blog.csdn.net/yellowstar5" , "Interval" : "10s" } }&emsp; 复制代码`

### 2.3.2. 生成两个token，让部门A，B各自管理自己的服务 ###

首先我们来生成部门A的policy， 意思度所有节点具有写权限（写权限包括读），并且只能写deptA开头的服务。

#### 2.3.2.1先生成部门A的token ####

` node_prefix "" { policy = "write" } service_prefix "deptA" { policy = "write" } 复制代码`

下面是具体的生成过程 1.新建policy并保存

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286eb8957a183?imageView2/0/w/1280/h/960/ignore-error/1) 2.生成token ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebba4b02fe?imageView2/0/w/1280/h/960/ignore-error/1) ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebba68a234?imageView2/0/w/1280/h/960/ignore-error/1) 3.查看token列表，并点击deptA-policy那一项查看并复制token ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebc2ddbcbf?imageView2/0/w/1280/h/960/ignore-error/1) ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebc569da95?imageView2/0/w/1280/h/960/ignore-error/1)

#### 2.3.2. 2.在生成部门B的token ####

只需要把policy稍作修改就可以，其他部门和部门A的类似，这里就不贴图了。

` node_prefix "" { policy = "write" } service_prefix "deptB" { policy = "write" } 复制代码`

## 2.3.用不同的token来获取服务列表 ##

最后我们拿到部门A和部门B的token，以及master的token 我这里deptA的token是：8764c083-0acb-e11e-433d-8d8803db9bd2 deptB的token是： 052f467f-9581-cc7c-a8a5-84d8df51dc9d master token是: d76a0f7-5535-40cc-8696-073462acc6c7 下面用postman测试一下，看看不同token返回的服务列表

deptA-token [http://127.0.0.1:7110/v1/agent/services?token=8764c083-0acb-e11e-433d-8d8803db9bd2]( https://link.juejin.im?target=http%3A%2F%2F127.0.0.1%3A7110%2Fv1%2Fagent%2Fservices%3Ftoken%3D8764c083-0acb-e11e-433d-8d8803db9bd2 )

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebc2ea0a89?imageView2/0/w/1280/h/960/ignore-error/1) deptB-token [http://127.0.0.1:7110/v1/agent/services?token=052f467f-9581-cc7c-a8a5-84d8df51dc9d]( https://link.juejin.im?target=http%3A%2F%2F127.0.0.1%3A7110%2Fv1%2Fagent%2Fservices%3Ftoken%3D052f467f-9581-cc7c-a8a5-84d8df51dc9d ) ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebe563dbe3?imageView2/0/w/1280/h/960/ignore-error/1) master-token [http://127.0.0.1:7110/v1/agent/services?token=cd76a0f7-5535-40cc-8696-073462acc6c7]( https://link.juejin.im?target=http%3A%2F%2F127.0.0.1%3A7110%2Fv1%2Fagent%2Fservices%3Ftoken%3Dcd76a0f7-5535-40cc-8696-073462acc6c7 ) ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ebf312d4bb?imageView2/0/w/1280/h/960/ignore-error/1) 可以发现deptA-token只能看到部门A的服务，deptB-token只能看到部门B的服务，master-token可以看到所以的。 另外取消注册，注册之类的验证大家可以自己试试，都与上面的方式差不多，只不过要使用不同的Http API ， [www.consul.io/api/agent/s…]( https://link.juejin.im?target=https%3A%2F%2Fwww.consul.io%2Fapi%2Fagent%2Fservice.html%25EF%25BC%258C%25E5%259C%25A8%25E8%25BF%2599%25E4%25B8%25AA%25E5%259C%25B0%25E5%259D%2580%25E4%25BD%25A0%25E5%258F%25AF%25E4%25BB%25A5%25E7%259C%258B%25E5%2588%25B0consul ) 提供的关于服务相关的api。

参考：

[www.consul.io/api/agent/s…]( https://link.juejin.im?target=https%3A%2F%2Fwww.consul.io%2Fapi%2Fagent%2Fservice.html ) [www.consul.io/docs/agent/…]( https://link.juejin.im?target=https%3A%2F%2Fwww.consul.io%2Fdocs%2Fagent%2Foptions.html ) 文章之后会第一时间发于微信， 欢迎关注我的微信公众号，大家一起交流学习

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/6/16b286ec00ef7c65?imageView2/0/w/1280/h/960/ignore-error/1)