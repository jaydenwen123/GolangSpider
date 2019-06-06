# 公益：开放一台Nacos服务端给各位Spring Cloud爱好者 #

之前开放过一台公益 [Eureka Server]( https://link.juejin.im?target=http%3A%2F%2Feureka.didispace.com%2F ) 给大家，以方便大家在阅读我博客中教程时候做实验。由于目前在连载Spring Cloud Alibaba，所以对应的也部署了一台Nacos，并且也开放出来，给大家学习测试之用。

* Nacos控制台

* 地址： [nacos.didispace.com/nacos/index…]( https://link.juejin.im?target=http%3A%2F%2Fnacos.didispace.com%2Fnacos%2Findex.html )
* 账户与密码均为：nacos

* 客户端使用配置

* 使用注册中心服务： ` spring.cloud.nacos.discovery.server-addr=nacos.didispace.com:80`
* 使用配置中心服务： ` spring.cloud.nacos.config.server-addr=nacos.didispace.com:80`

### Spring Cloud Alibaba系列专题 ###

下面是当前已经发布的内容，后续内容也将基于Spring Cloud Alibaba 0.2.2进行。

* [Spring Cloud Alibaba基础教程：使用Nacos实现服务注册与发现]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-1%2F )
* [Spring Cloud Alibaba基础教程：支持的几种服务消费方式（RestTemplate、WebClient、Feign）]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-2%2F )
* [Spring Cloud Alibaba基础教程：使用Nacos作为配置中心]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-3%2F )
* [Spring Cloud Alibaba基础教程：Nacos配置的加载规则详解]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-nacos-config-1%2F )
* [Spring Cloud Alibaba基础教程：Nacos配置的多环境管理]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-nacos-config-2%2F )
* [Spring Cloud Alibaba基础教程：Nacos配置的多文件加载与共享配置]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-nacos-config-3%2F )
* [Spring Cloud Alibaba基础教程：Nacos的数据持久化]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-4%2F )
* [Spring Cloud Alibaba基础教程：Nacos的集群部署]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-5%2F )
* [Spring Cloud Alibaba基础教程：使用Sentinel实现接口限流]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-sentinel-1%2F )
* [Spring Cloud Alibaba基础教程：Sentinel使用Nacos存储规则]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-sentinel-2-1%2F )
* [Spring Cloud Alibaba基础教程：Sentinel使用Apollo存储规则]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-alibaba-sentinel-2-2%2F )
* [查看更多...]( https://link.juejin.im?target=http%3A%2F%2Fblog.didispace.com%2Fspring-cloud-learning%2F )

**示例仓库**

* Github： [github.com/dyc87112/Sp…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdyc87112%2FSpringCloud-Learning%2Ftree%2Fmaster%2F4-Finchley )
* Gitee： [gitee.com/didispace/S…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fdidispace%2FSpringCloud-Learning%2Ftree%2Fmaster%2F4-Finchley )

**如果您对这些感兴趣，欢迎star、follow、收藏、转发给予支持！**

### 后记 ###

后续有更多资源之后，尽量再多部署一些，平时方便大家测试使用。汇总一下当前资源：

* Eureka Server： [eureka.didispace.com/]( https://link.juejin.im?target=http%3A%2F%2Feureka.didispace.com%2F )
* Nacos Server： [nacos.didispace.com/nacos/index…]( https://link.juejin.im?target=http%3A%2F%2Fnacos.didispace.com%2Fnacos%2Findex.html )