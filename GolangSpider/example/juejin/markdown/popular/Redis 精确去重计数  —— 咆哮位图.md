# Redis 精确去重计数 —— 咆哮位图 #

如果要统计一篇文章的阅读量，可以直接使用 Redis 的 incr 指令来完成。如果要求阅读量必须按用户去重，那就可以使用 set 来记录阅读了这篇文章的所有用户 id，获取 set 集合的长度就是去重阅读量。但是如果爆款文章阅读量太大，set 会浪费太多存储空间。这时候我们就要使用 Redis 提供的 HyperLogLog 数据结构来代替 set，它只会占用最多 12k 的存储空间就可以完成海量的去重统计。但是它牺牲了准确度，它是模糊计数，误差率约为 0.81%。

那么有没有一种不怎么浪费空间的精确计数方法呢？我们首先想到的就是位图，可以使用位图的一个位来表示一个用户id。如果一个用户id是32字节，那么使用位图就只需要占用 1/256 的空间就可以完成精确计数。但是如何将用户id映射到位图的位置呢？如果用户id是连续的整数这很好办，但是通常用户系统的用户id并不是整数，而是字符串或者是有一定随机性的大整数。

我们可以强行给每个用户id赋予一个整数序列，然后将用户id和整数的对应关系存在redis中。

` $next_user_id = incr user_id_seq set user_id_xxx $next_user_id $next_user_id = incr user_id_seq set user_id_yyy $next_user_id $next_user_id = incr user_id_seq set user_id_zzz $next_user_id 复制代码`

这里你也许会提出疑问，你说是为了节省空间，这里存储用户id和整数的映射关系就不浪费空间了么？这个问题提的很好，但是同时我们也要看到这个映射关系是可以复用的，它可以统计所有文章的阅读量，还可以统计签到用户的日活、月活，还可以用在很多其它的需要用户去重的统计场合中。所谓「功在当代，利在千秋」就是这个意思。

有了这个映射关系，我们就很容易构造出每一篇文章的阅读打点位图，来一个用户，就将相应位图中相应的位置为一。如果位从0变成1，那么就可以给阅读数加1。这样就可以很方便的获得文章的阅读数。

而且我们还可以动态计算阅读了两篇文章的公共用户量有多少？将两个位图做一下 AND 计算，然后统计位图中位 1 的个数。同样，还可以有 OR 计算、XOR 计算等等都是可行的。

问题又来了！Redis 的位图是密集位图，什么意思呢？如果有一个很大的位图，它只有最后一个位是 1，其它都是零，这个位图还是会占用全部的内存空间，这就不是一般的浪费了。你可以想象大部分文章的阅读量都不大，但是它们的占用空间却是很接近的，和哪些爆款文章占据的内存差不多。

看来这个方案行不通，我们需要想想其它方案！这时咆哮位图（RoaringBitmap）来了。

它将整个大位图进行了分块，如果整个块都是零，那么这整个块就不用存了。但是如果位1比较分散，每个块里面都有1，虽然单个块里的1很少，这样只进行分块还是不够的，那该怎么办呢？我们再想想，对于单个块，是不是可以继续优化？如果单个块内部位 1 个数量很少，我们可以只存储所有位1的块内偏移量（整数），也就是存一个整数列表，那么块内的存储也可以降下来。这就是单个块位图的稀疏存储形式 —— 存储偏移量整数列表。只有单块内的位1超过了一个阈值，才会一次性将稀疏存储转换为密集存储。

咆哮位图除了可以大幅节约空间之外，还会降低 AND、OR 等位运算的计算效率。以前需要计算整个位图，现在只需要计算部分块。如果块内非常稀疏，那么只需要对这些小整数列表进行集合的 AND、OR 运算，如是计算量还能继续减轻。

这里既不是用空间换时间，也没有用时间换空间，而是用逻辑的复杂度同时换取了空间和时间。

咆哮位图的位长最大为 2^32，对应的空间为 512M（普通位图），位偏移被分割成高 16 位和低 16 位，高 16 位表示块偏移，低16位表示块内位置，单个块可以表达 64k 的位长，也就是 8K 字节。最多会有64k个块。现代处理器的 L1 缓存普遍要大于 8K，这样可以保证单个块都可以全部放入 L1 Cache，可以显著提升性能。

如果单个块所有的位全是零，那么它就不需要存储。具体某个块是否存在也可以是用位图来表达，当块很少时，用整数列表表示，当块多了就可以转换成普通位图。整数列表占用的空间少，它还有类似于 ArrayList 的动态扩容机制避免反复扩容复制数组内容。当列表中的数字超出4096个时，会立即转变成普通位图。

用来表达块是否存在的数据结构和表达单个块数据的结构可以是同一个，因为块是否存在本质上也是 0 和 1，就是普通的位标志。

但是 Redis 并没有原生支持咆哮位图这个数据结构啊？我们该如何使用呢？

Redis 确实没有原生的，但是咆哮位图的 Redis Module 有。

[github.com/aviggiano/r…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Faviggiano%2Fredis-roaring )

这个项目的 star 数量并不是很多，我们来看看它的官方性能对比

+-------------+--------------+--------------+
|     OP      | TIME/OP (US) | ST DEV  (US) |
+-------------+--------------+--------------+
| R.SETBIT    |        31.89 |        28.49 |
| SETBIT      |        29.98 |        29.23 |
| R.GETBIT    |        29.90 |        14.60 |
| GETBIT      |        28.63 |        14.58 |
| R.BITCOUNT  |        32.13 |         0.10 |
| BITCOUNT    |       192.38 |         0.96 |
| R.BITPOS    |        70.27 |         0.14 |
| BITPOS      |        87.70 |         0.62 |
| R.BITOP NOT |       156.66 |         3.15 |
| BITOP NOT   |       364.46 |         5.62 |
| R.BITOP AND |        81.56 |         0.48 |
| BITOP AND   |       492.97 |         8.32 |
| R.BITOP OR  |       107.03 |         2.44 |
| BITOP OR    |       461.68 |         8.42 |
| R.BITOP XOR |        69.07 |         2.82 |
| BITOP XOR   |       440.75 |         7.90 |
+-------------+--------------+--------------+

很明显这里对比的是稀疏位图，只有稀疏位图才可以呈现出这样好看的数字。如果是密集位图，咆哮位图的性能肯定要稍弱于普通位图，但是通常也不会弱太多。

下面我们来观察一下源代码看看它的内部结构是怎样的

` // 单个块 typedef struct roaring_array_s { int32_t size; int32_t allocation_size; void **containers; // 指向整数数组或者普通位图 uint16_t *keys; uint8_t *typecodes; uint8_t flags; } roaring_array_t ; // 所有块 typedef struct roaring_bitmap_s { roaring_array_t high_low_container; } roaring_bitmap_t ; 复制代码`

很明显可以看到块存在与否和块内数据都是使用同样的数据结构表达的，它们都是 roaring_bitmap_t。这个结构里面有多种编码形式，类型使用 typecodes 字段来表示。

` # define BITSET_CONTAINER_TYPE_CODE 1 # define ARRAY_CONTAINER_TYPE_CODE 2 # define RUN_CONTAINER_TYPE_CODE 3 # define SHARED_CONTAINER_TYPE_CODE 4 复制代码`

看到这里的类型定义，我们发现它不止前面提到的普通位图和数组列表两种形式，还有 RUN 和 SHARED 这两种类型。RUN 形式是位图的压缩形式，比如连续的几个位 101,102,103,104,105,106,107,108,109 表示成 RUN 后就是 101,8（1 后面是 8 个自增的整数），这样在空间上就可以明显压缩不少。在正常情况下咆哮位图内部没有 RUN 类型的块。只有显示调用了咆哮位图的优化 API 才会转换成 RUN 格式，这个 API 是 roaring_bitmap_run_optimize。

而 SHARED 类型用于在多个咆哮位图之间共享块，它还提供了写复制功能。当这个块被修改时将会复制出新的一份。

咆哮位图的计算逻辑还有更多的细节，我们后面有空再继续介绍。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2016b50868aaf?imageView2/0/w/1280/h/960/ignore-error/1)