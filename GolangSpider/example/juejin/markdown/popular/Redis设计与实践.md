# Redis设计与实践 #

## Memcache和Redis有哪些区别？ ##

1、不支持持久化 VS 支持持久化，RDB、AOF，同时配置时，从AOF文件加载持久化文件。

2、简单key-value形式 VS 多种数据结构

3、多线程+锁 VS 单线程+事件轮询机制

4、伪分布式，客户端分发 VS P2P模式，没有中心节点，在redis cluster架构下，每个redis要放开两个端口号，比如一个是6379，另外一个就是加10000的端口号16379。16379端口号是用来进行节点间通信的，也就是cluster bus。cluster bus的通信用来进行故障检测，配置更新，故障转移授权。

5、Slab Allocation，预先分配一系列大小固定的组，然后根据数据大小选择最合适的块存储。避免了内存碎片。缺点是不能变长，浪费了一定空间，memcached默认情况下下一个slab的最大值为前一个的1.25倍。 VS malloc/free，由于malloc首先以链表的方式搜索已管理的内存中可用的空间分配，导致内存碎片比较多。

6、哈希环算法 VS 哈希槽算法，无损伸缩。

## Redis如何设计？ ##

单线程 -> CPU并非性能瓶颈，不需要考虑锁操作，避免上下为切换。

多路复用 -> 多路复用可以处理并发的连接。非阻塞IO 内部实现采用epoll，采用了epoll+自己实现的简单的事件框架。epoll中的读、写、关闭、连接都转化成了事件，然后利用epoll的多路复用特性，绝不在io上浪费一点时间。

RDB -> fork+写时复制

lru过期键回收 -> Redis 默认会每秒进行十次过期扫描，过期扫描不会遍历过期字典中所有的 key，而是 采用了一种简单的贪心策略。1、从过期字典中随机 20 个 key; 2、删除这 20 个 key 中已经过期的 key; 3、如果过期的 key 比率超过 1/4，那就重复步骤 1; 同时，为了保证过期扫描不会出现循环过度，导致线程卡死现象，算法还增加了扫描时间的上限，默认不会超过 25ms。

内存不足回收机制 -> 当实际内存超过配置的maxmemory，可选策略valatile-lru、volatile-ttl、volatile-random、allkeys-lru、allkeys-random。

同步策略 -> 主从增量同步，主从刚刚连接的时候，进行全量同步；全同步结束后，进行增量同步。

## 常用数据结构 ##

string、list、set、zset、hash、hyperloglog

## 主从配置 ##

slaveof ip port

## 哨兵如何配置？ ##

配置文件sentinel monitor mymaster 127.0.0.1 6379 1 启动命令redis-sentinal redis.sentinal.conf

## 主机挂了如何恢复？ ##

主机挂-等待-哨兵投票-从机变主机-主机恢复，变为从机

## 主机挂了如何选定主机？ ##

优先级、同步偏移量、结点id

## Redis Cluster如何进行动态扩容？ ##

先增加结点到集群redis-trib.rb add-node ... 然后重新分配slot到该节点redis-trib.rb reshard... 然后增加从结点、然后复制增加的主节点cluster replicate <主节点的ID>

## Redis分布式锁如何实现？缺点是什么？ ##

tryLock(){
SET Key UniqId Seconds } release(){
EVAL( //LuaScript if redis.call("get",KEYS[1]) == ARGV[1] then return redis.call("del",KEYS[1]) else return 0 end ) } 由于Redis集群数据同步为异步，假设在Master节点获取到锁后未完成数据同步情况下Master节点crash，此时在新的Master节点依然可以获取锁，所以多个Client同时获取到了锁。

## Redis集群如何进行在线迁移？ ##

唯品会 Redis-Migrate-Tool

参考：《Redis设计与实现》《Redis深度历险》