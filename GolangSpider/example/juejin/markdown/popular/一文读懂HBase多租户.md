# 一文读懂HBase多租户 #

> 
> 本文从三个方面介绍了HBase的多租户实现。
> 上篇文章回顾： [HDFS短路读详解](
> https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUxMDQxMDMyNg%3D%3D%26amp%3Bmid%3D2247485652%26amp%3Bidx%3D1%26amp%3Bsn%3D7859175e1c0cba7290bcd61a662cadc6%26amp%3Bchksm%3Df90223edce75aafb4c60c969c9a74781c4696fe8fb6f450f133f8e1b3dd6ed847f5e94bf5c30%26amp%3Bscene%3D21%23wechat_redirect
> )

多租户(multi-tenancy technology)，参考维基百科定义，它是在探讨与实现如何于多用户的环境下共享相同的系统或程序，并且仍可确保各用户间数据的隔离性。随着云计算时代的到来，多租户对于云上服务显得更加重要。所以HBase也有许多多租户相关的功能，其为多个用户共享同一个HBase集群，提供了资源隔离的能力。本文将从Namespace&ACL，Quota，RSGroup三个方面来进行介绍。

## Namespace&ACL ##

在HBase中，创建namespace是一个很轻量的操作，将不同业务的表隔离在不同的namespace是一个最简单的资源隔离的方法。同时，ACL、quota、 rsgroup等常用的资源隔离方式都支持设置在namespace上。

ACL，全称Access Control Lists，用于限制不同的用户对不同的资源的操作或访问权限。

使用ACL需要添加如下配置：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068ce15a48c?imageView2/0/w/1280/h/960/ignore-error/1)

### 1、ACL的几个概念 ###

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068d24c14c3?imageView2/0/w/1280/h/960/ignore-error/1)

User分为普通user和super user。super user包括启动HBase服务的用户和hbase.superuser配置的用户，可以对集群进行管理操作。普通用户需要授权后，才能访问或操作HBase。Scope可以理解为资源的粒度。

HBase的各种操作需要的Action可以在HBase的官方文档中查看：http://hbase.apache.org/book.html#appendix_acl_matrix

结合用户的访问或操作需求，将user在合理的scope上设置合理的action，是实现用户权限控制的最佳方式。

### 2、设置或取消权限 ###

在HBase shell中或调用HBase API，设置或取消权限。shell中的操作如图：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068d10171f5?imageView2/0/w/1280/h/960/ignore-error/1)

设置namespace的权限需要加@前缀：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068d08a1896?imageView2/0/w/1280/h/960/ignore-error/1)

设置Cell的权限：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068d09d9d36?imageView2/0/w/1280/h/960/ignore-error/1)

### 3、权限的存储 ###

存储在hbase:acl表中，rowkey是根据scope计算出来的。acl表结构如下表：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068f14a23c1?imageView2/0/w/1280/h/960/ignore-error/1)

Cell权限使用tags of HFile v3存储。

### 4、鉴定权限 ###

鉴定权限是指判断某个用户是否拥有某个操作的权限。这个过程是在AccessController中完成的，AccessController是一个实现了MasterObserver、RegionServerObserver、RegionObserver等的coprocessor，在master、regionserver、region等操作的hook中检查权限。由于每台RS上都维护了完整的PermissionCache，检查PermissionCache中是否包含了所需的权限，如果权限不足，则抛出AccessDeniedException。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c0691497967d?imageView2/0/w/1280/h/960/ignore-error/1)

### 5、添加/删除权限 ###

添加/删除授予的过程如下图所示：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c0690ec0c3ad?imageView2/0/w/1280/h/960/ignore-error/1)

（1）client向有acl region的region server发出grant或revoke请求；

（2）收到请求的region server，将新的权限put或者delete到acl表中；

（3）AccessController在region的postPut和postDelete的hook中，如果操作的是acl region，则将更新的权限从acl table中读出，并写入到zk上；

（4）通过zk的监听机制，通知master和regionserver更新PermissionCache，实现权限在master和其他regionserver中的同步。

### 6、基于Procedure的添加/删除权限 ###

为了使用Procedure实现权限的同步，需要首先将grant/revoke请求发送到master处理， 参考HBASE-21739。然后在添加/删除权限阶段，主要有两个关键的步骤，一是记录权限到acl table中，二是将更新后的权限同步到全部的RegionServer上。设计了UpdatePermissionProcedure来实现这个操作，参考HBASE-22271(目前还没有合并到社区版的master分支)。在UpdatePermissionStorage阶段，更新acl表及zk，master上的PermissionCache，在UpdatePermissionCacheOnRS阶段，发起UpdatePermissionRemoteProcedure，更新RS的PermissionCache。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068f4a250c8?imageView2/0/w/1280/h/960/ignore-error/1)

UpdatePermissionProcedure需要解决五种权限同步的case：

**Grant** ：添加权限

**Revoke：** 删除权限

**Delete Namespace：** 删除namespace的全部权限

**Delete Table：** 删除table的全部权限

**Reload：** 重新获取全部的Permission。

在新的方案中，zk不用于通知RS更新PermissionCache，只用于acl的存储。因为当RS或Master启动时，acl table不一定online，此时，需要从zk上load permission。当acl表中的权限与zk上的权限不一致时，应该以acl表中的权限为准。因此，当master启动且acl table online后，发起类型为Reload的UpdatePermissionProcedure，更新zk上的permission，并更新RS上的PermissionCache。

## Quota&Throttle ##

由于集群的资源及服务能力是有上限的，Quota用于限制各个资源的数据量的大小及访问速度。

需要如下配置开启HBase的quota功能：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06903a8c1d6?imageView2/0/w/1280/h/960/ignore-error/1)

HBase中关于Quota的几个概念及其相互关系如下图所示：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c068f59abd6d?imageView2/0/w/1280/h/960/ignore-error/1)

### 1、Throttle Quota ###

Throttle限制单位时间内，访问资源的次数或数据量。

* 

支持的时间单位包括sec, min, hour, day。

* 

使用req限制请求的次数；

* 

使用B, K, M, G, T, P限制请求的数据量的大小；

* 

使用CU限制请求的读/写容量单位，一个读/写容量单位是指一次读出/写入数据量小于1KB的请求，如果一个请求读出了2.5K的数据，则需要消耗3个容量单位。可以通过hbase.quota.read.capacity.unit或hbase.quota.write.capacity.unit配置一个容量单位的数据量。

* 

Machine scope代表throttle额度配置在单台RS上。Cluster代表throttle配额被集群的全部RS共享。如果不指定QuotaScope的话，默认为Machine。

设置Throttle的shell命令如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06916fb25b1?imageView2/0/w/1280/h/960/ignore-error/1)

设置RegionServer的throttle(目前只支持使用all关键字代表全部的RegionServer，不支持对指定的RegionServer设置Quota)，一般来说，RS的quota代表该RS的服务上限，推荐以秒为时间单位设置：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06926fd1984?imageView2/0/w/1280/h/960/ignore-error/1)

设置Cluster scope的quota：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c0693b7b62cc?imageView2/0/w/1280/h/960/ignore-error/1)

Cluster scope的quota是如何分配到各个RS上的：

* 

对于table的quota，TableMachineLimit = ClusterLimit / TotalTableRegionNum * MachineTableRegionNum；

* 

对于namespace的quota，NamespaceMachineLimit = ClusterLimit / RsNum，需要注意的是，这里没有考虑RSGroup，如果把namespace隔离到某个RSGroup，分配到RS上的throttle limit是偏小的，后续需要改进这个计算方式。

GlobalBypass在全局范围内，跳过throttle，配置在用户上。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c069239f2ce4?imageView2/0/w/1280/h/960/ignore-error/1)

### 2、Space Quota ###

Space用于限制资源的数据量大小，配置在namespace或者table上。当数据量达到限额时，执行配置的违反策略，包括：

**Disable：** disable table/ the tables of namespace

**NoInserts：** 禁止除Delete以外的Mutation操作，允许Compaction

**NoWrites：** 禁止Mutation操作，允许Compaction

**NoWritesCompactions：** 禁止Mutation操作，禁止Compaction

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06944aa2b91?imageView2/0/w/1280/h/960/ignore-error/1)

看当前Space quota的快照(这里的快照并不是HBase中的快照)，而是指当前表的空间大小，配置的limit，触发的策略的状态：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c0693c2e1359?imageView2/0/w/1280/h/960/ignore-error/1)

限制namespace的table或region数量：

` hbase.namespace.quota.maxtables/hbase.namespace.quota.maxregions 复制代码`

如果超出限制的话，会抛出QuotaExceededException。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06947d79617?imageView2/0/w/1280/h/960/ignore-error/1)

Space quota的实现原理是：

（1）RS周期的把Region size信息发送给master：RegionSizeReportingChoreMaster

（2）统计表的size及触发的策略并存到quota表：QuotaObserverChoreRS

（3）周期的读quota表，执行policy：SpaceQuotaRefresherChore

### 3、Soft limit ###

配置throttle limit为soft limit，也就是在集群资源富余的情况下，允许超发，使用如下命令打开或关闭超发：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c069485cb0fe?imageView2/0/w/1280/h/960/ignore-error/1)

注意，超发是指允许用户在RS的quota有富余的情况下，允许请求超出配置的user/namespace/table的quota，因此，必须首先设置RS的quota，才能打开超发功能。RS的quota推荐设置的时间单位为秒，因为使用其他时间单位的话，一旦RS的quota被其它用户的请求先消耗的话，恢复quota需要较长的时间，可能会影响后续的请求，即使这些后来的请求并没有超出其配置的user/namespace/table quota。

### 4、Quota存储 ###

quota相关的信息存储在hbase:quota表中。

row key主要有以下几种：

**n.namespace：** namespace的quota

**t.table：** table的quota

**u.user：** user的quota

**r.all：** RegionServer的quota

**exceedThrottleQuota：** 是否允许超发

Throttle相关的quota存储在q CF中，Space相关的quota存储在u CF中。

Throttle是否打开存储在/hbase/rpc-throttle的zk节点上，值为true或者false。因为打开或关闭Throttle是实时生效的，而其它quota配置是通过RS定期的读quota表，是延迟生效的。

### 5、Throttle ###

设置throttle分为2步：

（1）client向master发送set quota请求，master把quota存入hbase:quota表中；

（2）RS每五分钟，从quota表中加载最新的quota值并更新QuotaCache。因此，对于新设置的quota，最多五分钟后生效（可以通过hbase.quota.refresh.period配置时间间隔）。

当读写请求到达RS上时，限流过程如下图所示：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06960ea63bd?imageView2/0/w/1280/h/960/ignore-error/1)

其中，在读数据前，会首先预估本次请求将要消耗的quota数目，目前社区的代码是按照一个get或mutate预计消耗100字节，一个scan预计消耗1000字节，这里应该是可以优化的，可以根据上次请求后读出的数据量来动态的调整预估的字节数。

Throttle limit是设置在某个时间单位上的，会随着时间的推移逐渐恢复，主要有两种恢复方式：

（1）Average Interval Refill(默认)：根据当前和上一次的恢复时间，恢复出这段时间内的quota，但最大不能超出quota配置的limit。

比如，配置了100资源/秒，100ms后，恢复出10个资源。2s后，恢复出100资源，而不是200资源。

（2）Fixed Interval Refill：经过固定的时间间隔，恢复出全部quota。

比如，配置了100资源/秒，如果上次quota恢复的时间是10:10:10,100，则下次恢复时间为10:10:11,100，并记录本次恢复时间，如果在10:10:11,099访问，此时可用资源依然为0。

打开或关闭限流：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06948414512?imageView2/0/w/1280/h/960/ignore-error/1)

关闭限流时，配置的throttle将不会进行限流，即使集群开启了quota功能。

## RSGroup ##

RSGroup，是把RS分配到不同的组中，之后，将namespace或者table分配到某个RSGroup中，从而实现隔离的目的，可以形象的理解为每个RSGroup组成了一个小集群。 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06959bada69?imageView2/0/w/1280/h/960/ignore-error/1)

使用RSGroup，需要添加如下配置：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c069621a38cd?imageView2/0/w/1280/h/960/ignore-error/1)

当开启RSGroup后，所有的RS默认在default这个group中。

创建新的group后，必须首先移入RS到这个group中，之后才能把namespace或者table移动到这个group中。

添加新的RSGroup：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06963d75489?imageView2/0/w/1280/h/960/ignore-error/1)

先将RS移动到这个group中，再将namespace移动到这个group中：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06966369f53?imageView2/0/w/1280/h/960/ignore-error/1)

RSGroup的功能主要在RSGroupAdminEndpoint中实现，它是一个实现了MasterObserver的Endponit，在master操作的hook中，将table的region移动到对应的RSGroup中。

RSGroup的信息存储在hbase:rsgroup表中。同时，RSGroup的信息也在zk中存储，当集群启动时，rsgroup表还没有online时，从zk中读出RSGroup的信息。

综上，就是HBase中多租户相关功能的介绍，希望大家在生产环境中多多使用，并向社区反馈改进建议，共同推动HBase多租户功能的进一步完善。

### 关于作者 ###

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c06988a0d257?imageView2/0/w/1280/h/960/ignore-error/1)

梅祎，小米最年轻的美女HBase Committer，梅祎所在的HBase生态组，团队技术氛围浓厚，累计培养了9位HBase Committer和2位PMC，还有多位开源项目Contributer。欢迎开源爱好者及HBase爱好者加入我们一起成长。

PS:他们团队还在招人哦，戳 [招募令|寻找优秀工程师]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzUxMDQxMDMyNg%3D%3D%26amp%3Bmid%3D2247485626%26amp%3Bidx%3D2%26amp%3Bsn%3D2192a101c1104176e053474cc02a9830%26amp%3Bchksm%3Df9022383ce75aa95786f7c1cd93dfea4f6e52c5d90804757848c8e48a504e8eb2d15b720e38a%26amp%3Bscene%3D21%23wechat_redirect ) 了解详情。

本文首发于公众号“小米云技术”， [点击查看原文]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FT3d3VY0NketO8m8ogJG2wQ ) 。