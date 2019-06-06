# 详细讲解：零知识证明 之 ZCash 完整的匿名交易流程 #

> 
> 
> 
> 作者：林冠宏 / 指尖下的幽灵
> 
> 

> 
> 
> 
> 掘金： [juejin.im/user/587f0d…](
> https://juejin.im/user/587f0dfe128fe100570ce2d8 )
> 
> 

> 
> 
> 
> 博客： [www.cnblogs.com/linguanh/](
> https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2Flinguanh%2F )
> 
> 
> 

> 
> 
> 
> GitHub ： [github.com/af913337456…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Faf913337456%2F )
> 
> 

> 
> 
> 
> 腾讯云专栏： [cloud.tencent.com/developer/u…](
> https://link.juejin.im?target=https%3A%2F%2Fcloud.tencent.com%2Fdeveloper%2Fuser%2F1148436%2Factivities
> )
> 
> 

> 
> 
> 
> 虫洞区块链专栏： [www.chongdongshequ.com/article/153…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.chongdongshequ.com%2Farticle%2F1536563643883.html
> )
> 
> 

## 目录 ##

* 前序
* 交易体的结构 note
* commitment 和 nullifier
* ZCash 1.0 的公私钥机制
* 转账人发出交易 note
* 收款人如何获取 note 的使用权
* 零知识自证
* 后记

### 前序 ###

在这篇文章中，我将承接上一篇文章 [详细讲解：零知识证明 之 zk-SNARK 开篇]( https://juejin.im/post/5ce02116e51d4510744c23d6 ) (开篇中介绍了什么是零知识证明及其它术语) 来 ` 从一个完整的交易流程` 讲解 ` ZCash` 是如何利用 ` 零知识证明` 的 ` zk-SNARK` 实现匿名交易的。

其中第 ` 六` 部分 ` 收款人如何获取 note 的使用权` 是目前国内网上所有的介绍 ` ZCash` 的文章都没有谈及的， ` 造成了读者只知道交易的发出` ，而 ` 不知道交易是凭借什么机制让收款人有权限使用的` 。

此外， ` "现在关于 ZCash 的文章和回答，很多都不准确，甚至是有误导性的！"` 此话---引自 [woodstock]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fpeople%2Fkunxian%2Fpins )

文章不从源码分析的角度去展开，那样的写作和阅读成本太高。

### 交易体的结构 note ###

首先 ` ZCash` 在交易的整体模式上，参考了 ` BTC` 的 ` UTXO` 模型，拥有交易输入和交易输出的概念，对于 ` UTXO` 的讲解，可以自行网上搜索文章进行阅读，目前介绍 ` UTXO` 的优秀文章还是很多的。

` UTXO` 是一种模型，模型是可以被以不同的形式展现出的。在 ` ZCash` 中，交易原始的输入输出结构体被形象成了代码中的 ` note` 结构体。

一个完整的 ` note` 包含有如下的变量：

* 持有者的公钥: a_pk，又称收款人地址
* 面额: value，又被简称为 v，代表这笔 note 的代币数值
* 随机数: rho, 是每一条 note 的唯一标识，当一条 note 被消费了之后，这个值会被放置到 ` nullifier` 表中，代表这条 note 已经被消费了，再次进行消费同一条 note的时候，会触发 ` 双花` 错误，即交易双花防护机制。
* 随机数: r

用向量组代表上面的 note，可以表示为： ` note = <a_pk , v , r , rho>`

### commitment 和 nullifier ###

在 ` ZCash` 中，存在两种表格，分别是： ` commitment` 和 ` nullifier` ，下图取自于 ` ZCash` 的官方文档 [How Transactions Between Shielded Addresses Work]( https://link.juejin.im?target=https%3A%2F%2Fz.cash%2Fblog%2Fzcash-private-transactions%2F ) 中， ` 提示` ：该文章内部并没有指出 note 的收款人是如何对一笔 note 有使用权限的，看完也会有很多疑问。

图中显示出了 ` commitment` 和 ` nullifier` 表格的大致结构：

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae2c20ac9746f0?imageView2/0/w/1280/h/960/ignore-error/1)

左边的 ` hashed notes` 就是 ` commitment` 列表，右边的 ` nullifier set` 就是 ` nullifier` 列表。

` commitment` 列表存储的是所有的，注意！是所有。存储所有 note 经过不可逆 hash 函数后生成的 hash 值 ` Hx` x∈ (1,2,3,4,5,6...N）

` nullifier` 列表存储的是 ` 已经被消费` 的 note 中的随机数 ` r` 生成的 hash 值 ` nfx` 。r 就是 note 结构中的 r。 ` nullifier` 中文意思： ` 作废` ， ` nullifier set` 作废集合。

` 注意一点` ：对于两个不同的 note，他们的 commitment hash 值一定不相同，从 hash 值又无法推测出其背后的 note.。

如上图所示，在右边我们可以看到 r2 对应的 note2 的 nf1 已经被记录到了 ` nullifier` 列表中，这个 nf1 就是结构体中的 rho。被记录到了这里，代表着 note2 不再是 ` UTXO` ，不再是没被花费的输出，它已经被消费了。

一条 ` 被花费的输出 output` 会导致一条 ` 新产生的交易输入 input` 。继续以上图为例子，note2 作为被消费了的输出，据表可以， note3 应该是它所产生的交易的输入，此时 note3 还没被消费，因为它的 nf3 还没被记录到 ` nullifier` 列表。note2 相对于 note3 来说，note2 是 note3 的输入。note3 作为一条新的交易输入，还没输出给其它的 note。

### ZCash 1.0 的公私钥机制 ###

在认识交易被发出前的操作，我们现来认识下 ` ZCash 1.0` 的公私钥机制，下图取自于官方文档：

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae2d2f78f555d3?imageView2/0/w/1280/h/960/ignore-error/1)

下面我将讲解下这图的主要要表达的意思，上面我列举的 note ，是一个整体的讲解，在实际的 ` ZCash` 源码中，note 其实还分为两种，分别是： ` SproutNote` 和 ` SaplingNote` 。目前 ` ZCash` 使用的是第一种，本文所谈的也是 ` SproutNote` 。 ` ZCash` 的发展将会慢慢向第二种 ` SaplingNote` 迁移。

因为 note 结构中有一个 a_pk 字段，在 ` SproutNote` 和 ` SaplingNote` 中，内部的字段组成是不同的，源码的定义如下图所示：

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae2d6361860d77?imageView2/0/w/1280/h/960/ignore-error/1)

` SaplingNote` 中，明显 a_pk 不见了，多了其它的。回到下图，我们主要看左边的 ` Sprout` ，其中：

* 双竖线表示相同
* 箭头表示生成关系，如果 A 指向B, 那么表示A可以生成B。

下图的 ` Sprout` 中 a_sk 代表的是私钥，由 a_sk 可以生成第 ` 一个公钥 a_pk` 和 ` 私钥 sk_enc` ，由 sk_enc 可以生成第二个公钥 ` pk_enc` 。意味着：

> 
> 
> 
> 在ZCash 1.0 中，一个钱包地址里面包含由两个公钥：a_pk , pk_enc
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae2d2f78f555d3?imageView2/0/w/1280/h/960/ignore-error/1)

### 转账人发出交易 note ###

现在进行到转账人发出交易 note，假设转账人是A，收款人是B，A 要转 5 个币个B。

那么 A 组装 note 的过程如下：

* 首先 A 找到自己的一条或多条没消费的 note，即是 ` UTXO` 输出，每条 note 中有对应的 value，我们假设一条就足够转出，多条的情况是如果一条 note 无法满足目标转出 value，才会凑多条 note 作为输出。
* A 找出了 note 1 ，使用自己的 ` 私钥 sk_enc` 解密 note 1，获取 note 1 中的 value 和其它数据，假设 value 是8，此时 ` 8 > 5` 。
* A 先新建两条 note，分别是 note4 和 note5，note4 内部的 value 设置为 5，代表是要给 B 的。note 5 的 value 是 (8-5=3)，代表给自己找零，不找零将会损失掉这3个币，相关的 ` 找零` 解析见 ` UTXO` 模型介绍。
* A 为 note 4 和 note 5 分别生成随机数 r4 和 r5
* A 将 B 的 a_pk 公钥设置到 note 4 里面去， 代表收款人是 B。再将自己的 a_pk 公钥设置到 note 5 里面去，代表收款人是自己。
* 使用 hash 函数生成 note 4 和 note 5 的 rho。 ` PS: ( rho = nf = HASH (r) )`
* 此时 note 4 和 note 5 分别是：

> 
> 
> 
> note4 = <B的a_pk ，v=5，r4，rho4>
> 
> 

> 
> 
> 
> note5 = <A的a_pk ，v=3，r5，rho5>
> 
> 

* 与此同时，A 还要将 note 1 的 nf2 ( ` nf2=HASH (r1)` ) 发往公链的节点网络，即 note 1 的 rho，此时节点在收到 nf2 后会判断是否已经在 ` nullifier` 列表存在 nf2 了，是的话，那么判断 note 1 被双花了。否则，就将 note 1 的 nf2 记录下来了。
* A 此时使用 B 的 pk_enc 签名 note 4，和自己的 pk_nec 签名 note 5。这里 B 的 pk_enc 是公开的，注意！ZCash 1.0 中，一个地址的 a_pk 和 pk_enc 都是公开的。
* A 将 note 4 通过秘密通道发给 B，自己的 note 5 便自己保存，同时将 note 4 和 note 5 的 hash 值 h4、h5 发给所有链上的节点。
* 以上便是一个 ` 发起交易` 的流程。

#### 对上述流程可以提出下面细节。 ####

* 节点能够知道的只有 note1 的 nf2 和 note4 的 h4 和 note5 的 h5。他们对收款人地址，金额是多少都一无所知。
* 此时链上节点们维护的 ` commitment` 和 ` nullifier` 表变成了如下的样子。

+-------------------------+-----------------+
| NOTE HASHS(COMMITMENTS) |  NULLIFIER SET  |
+-------------------------+-----------------+
| h1=HASH (note1)         | nf1 = HASH (r2) |
| h2=HASH (note2)         | nf2 = HASH (r1) |
| h3=HASH (note3)         |                 |
| h4=HASH (note4)         |                 |
| h5=HASH (note5)         |                 |
+-------------------------+-----------------+

* A 发给 B note4 的秘密通道。这里我们不展开说，它的方式有很多，可以在用加密的电子邮件传送，或者可以面对面传递小纸条。具体见官方完整的文档： [完整文档]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzcash%2Fzips%2Fblob%2Fmaster%2Fprotocol%2Fprotocol.pdf )
* B 手上有经过了自己的 sk_enc 签名了的 note4

### 收款人如何获取 note 的使用权 ###

其实在上面小节中，读者应该能理会到 B 是可以使用自己的原始私钥 a_sk 对自己获取到的 note4 数据进行解密的，进而获取到里面的 <a_pk，v，r，rho>，到了这一步，B 保存好 note4 的 rho，那么他就能够向 ` nullifier` 表发送 note4 的 rho ，达到消费 note4 的目的。

` 至此，ZCash 的匿名交易流程形成了闭环` 。

那么为什么 ` a_sk` 能对 ` sk_enc` 签名的数据，进行解密呢？因为在 ZCash 1.0 中，由地址的公私钥生成规则，可知原始私钥 a_sk 可以导出 sk_enc。在 ` ZCash 1.0 的公私钥机制` 小节中也做了说明。

### 零知识自证 ###

* 节点在校验完了一个 note 的 rho 后，如何判断发送者是真正对 note 拥有使用权的？

答：对于 note 的所有权拥有者A 来说，好像除了公布 note 里面的内容外，好像没其它手段来自我证明？这个时候 ` 零知识证明` 就排上用场了，note 的拥有者在发布使用该 note 的时候还要向节点出示称为 ` Π` 的零知识证明凭据，根据 Π ，节点们作为验证者，能够验证 note 的使用权的确属于A。ZCash 在这里应用到了 ` 零知识证明` ，它的代码是根据 ` zk-SNARK` 理论完成的，同时也参考了 ` Zerocash` 。

* PS：如何统计一个地址的余额？

答：不在本文的讨论范围内，详细可以见官方完整的文档： [完整文档]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzcash%2Fzips%2Fblob%2Fmaster%2Fprotocol%2Fprotocol.pdf ) 的 ` 第4章` ，关于 Balance 的描述。

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3489db0e1643?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae348f90fcd728?imageView2/0/w/1280/h/960/ignore-error/1)

## 后记 ##

为了理清 ZCash 的匿名交易的最后一部分，也即是收款人是如何获得 note 的所有权的，使整个流程 ` 形成闭环` 。

我查阅了很多文章都、文档和咨询了一位网友已经查看了部分源码。目前网上的其它文章都没有讲到 ` 收款人如何获取 note 的使用权` 这一部分。稍微较好的是对官方文档 [在隐藏地址之间如何进行交易]( https://link.juejin.im?target=https%3A%2F%2Fz.cash%2Fzh%2Fblog%2Fzcash-private-transactions%2F ) 的直接做了翻译，但是由于官方的这篇文章是个简版，也没有对 ` 收款人如何获取 note 的使用权` 做出解析，所以，几乎所有的翻译文章都是没答案的，而且大部分文章，本身还有一些错误，可能作者自己也一知半解。

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae36e0877a603b?imageView2/0/w/1280/h/960/ignore-error/1)

感谢下面的人和文给予了我有用的信息引导：

人： [woodstock]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fpeople%2Fkunxian%2Factivities )

文：

* [z.cash/zh/blog/zca…]( https://link.juejin.im?target=https%3A%2F%2Fz.cash%2Fzh%2Fblog%2Fzcash-private-transactions%2F )
* [github.com/zcash/zips/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzcash%2Fzips%2Fblob%2Fmaster%2Fprotocol%2Fprotocol.pdf )
* [zhuanlan.zhihu.com/p/25168970]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F25168970 )
* [zhuanlan.zhihu.com/p/58006716]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F58006716 )
* [blog.csdn.net/lvbin2012/a…]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Flvbin2012%2Farticle%2Fdetails%2F87339907 )