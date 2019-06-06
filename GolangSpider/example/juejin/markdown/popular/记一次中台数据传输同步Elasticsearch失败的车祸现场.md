# 记一次中台数据传输同步Elasticsearch失败的车祸现场 #

> 
> 
> 
> 欢迎关注个人微信公众号: **小哈学Java** , 优质文章第一时间推送哟！
> 
> 
> 
> 个人网站: [www.exception.site/essay/elast…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.exception.site%2Fessay%2Felasticsearch-sync-index-read-only-allow-delete
> )
> 
> 

## 目录 ##

* 一、背景
* 二、题外话
* 三、开始排查
* 四、为什么索引处于只读状态呢？
* 五、如何解决

## 一、背景 ##

前几天小哈在钉钉群里收到重庆业务线反馈，说是中台数据传输中间件在同步 Mysql 增量数据到 Elasticsearch 总是失败。

![什么](https://user-gold-cdn.xitu.io/2019/6/2/16b1810e7f6edf30?imageView2/0/w/1280/h/960/ignore-error/1) 什么

## 二、题外话 ##

你说的这个数据传输和阿里云提供的数据传输DTS是一个东西吗？

![阿里数据传输DTS](https://user-gold-cdn.xitu.io/2019/6/2/16b1810e8086da3e?imageView2/0/w/1280/h/960/ignore-error/1) 阿里数据传输DTS

不是！上面说的数据传输是小哈所在的中台研发部自主研发的中间件，目的是为了取代各业务线对阿里DTS同步功能的依赖！

目前来说，数据传输还是要依赖于阿里开源 Canal, 或者阿里 DTS，依赖的目的是实现对 Mysql 数据库 binlog 增量订阅。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1810e822b8244?imageView2/0/w/1280/h/960/ignore-error/1)

以上网络架构示例图中，中台数据传输充当一个 binlog 事件消费者的角色，通过自定义规则映射，数据加工，分发并最终同步到目标源 Elasticsearch 中。

## 三、开始排查 ##

回归正题，出了问题，立马赶紧通过跳板机连上数据传输所在的服务器，开始查看日志：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1810e814ab5d3?imageView2/0/w/1280/h/960/ignore-error/1)

看到日志中存在大量的 ` [FORBIDDEN/12/index read-only / allow delete (api)]` 错误！！

提示错误也很明显： **ES 索引处于只读状态** ！！在和业务组沟通以后，发现需要同步的目标索引有两个，一个商品索引（充当主表），一个商品属性索引(充当商品从表)，从表同步是 ok 的，也就是说商品属性索引非只读状态，写入正常，仅仅是商品索引处于只读状态，最终未能正常同步数据。

## 四、为什么索引处于只读状态呢？ ##

什么原因导致的索引只读的？小哈开始翻阅 [Elasticsearch 官方文档]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fguide%2Fen%2Felasticsearch%2Freference%2Fcurrent%2Fdisk-allocator.html ) , 原文如下：

> 
> 
> 
> Elasticsearch considers the available disk space on a node before deciding
> whether to allocate new shards to that node or to actively relocate shards
> away from that node.
> 
> 

Elasticsearch 在决定是否分配新分片给该节点，或对该节点重新定位分片之前，会先判断该节点存储空间是否足够，如果说你的使用磁盘空间已经超过 95%，ES 会自动将索引 ` index` 置为 ` read-only` 状态。

于是，让运维看下 ES 机器的磁盘空间是否足够，运维反馈说：前两天就是因为磁盘不足告警，刚刚扩的容，肯定是够的！

真相大白了！

前两天磁盘空间不足，那个时候，商品索引刚好有写入的操作，由于 ES 的保护机制，将该索引置为了只读状态。

## 五、如何解决 ##

原因找到了！要如何解决呢？

处于只读状态的索引，只能被查询或者删除。而 ES 还不会自动将索引状态切换回来，就需要我们手动切换了：

` PUT /<yourindex>/_settings { "index.blocks.read_only_allow_delete" : null } 复制代码`

对商品索引执行如上命令后。让业务组再次同步数据，一切正常了。

## 六、参考 ##

* [www.elastic.co/guide/en/el…]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fguide%2Fen%2Felasticsearch%2Freference%2Fcurrent%2Fdisk-allocator.html )
* [stackoverflow.com/questions/4…]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F48032661%2Ftransporterror403-ucluster-block-exception-ublocked-by-forbidden-12-inde )
* [blog.51cto.com/michaelkang…]( https://link.juejin.im?target=https%3A%2F%2Fblog.51cto.com%2Fmichaelkang%2F2164181 )
* [www.aityp.com/%E8%A7%A3%E…]( https://link.juejin.im?target=https%3A%2F%2Fwww.aityp.com%2F%25E8%25A7%25A3%25E5%2586%25B3elasticsearch%25E7%25B4%25A2%25E5%25BC%2595%25E5%258F%25AA%25E8%25AF%25BB%2F )

## 欢迎关注微信公众号: 小哈学Java ##

![关注微信公众号【小哈学Java】,回复【资源】，即可免费无套路领取资源链接哦](https://user-gold-cdn.xitu.io/2019/5/18/16aca6d38b8f78d9?imageView2/0/w/1280/h/960/ignore-error/1) 关注微信公众号【小哈学Java】,回复【资源】，即可免费无套路领取资源链接哦