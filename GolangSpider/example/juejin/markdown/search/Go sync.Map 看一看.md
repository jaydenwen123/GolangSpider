# Go sync.Map 看一看 #

偶然看见这么篇文章： [一道并发和锁的golang面试题]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_28163175%2Farticle%2Fdetails%2F75287877 ) 。 虽然年代久远，但也稍有兴趣。

正好最近也看到了 sync.Map，所以想试试能不能用 sync.Map 去实现上述的功能。

我还在 gayhub上找到了其他人用 sync.Mutex 的实现方式， [【点击这里】]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FAbelLai%2FWhat-I-Will-Say%2Fissues%2F32 ) 。

## 归结一下 ##

需求是这样的：

> 
> 
> 
> 在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出
> success: 100。
> 
> 

并且给出了原始代码：

` package main import ( "fmt" "time" ) type Ban struct { visitIPs map [ string ]time.Time } func NewBan () * Ban { return &Ban{visitIPs: make ( map [ string ]time.Time)} } func (o *Ban) visit (ip string ) bool { if _, ok := o.visitIPs[ip]; ok { return true } o.visitIPs[ip] = time.Now() return false } func main () { success := 0 ban := NewBan() for i := 0 ; i < 1000 ; i++ { for j := 0 ; j < 100 ; j++ { go func () { ip := fmt.Sprintf( "192.168.1.%d" , j) if !ban.visit(ip) { success++ } }() } } fmt.Println( "success: " , success) } 复制代码`

哦吼，看到源代码我想说，我能只留个 ` package main` 其他都重新写吗？（捂脸）

聪明的你已经发现，这个问题关键就是想让你给 Ban 加一个读写锁罢了。 而且条件中的三分钟根本无伤大雅，因为这程序压根就活不到那天。

## 思路 ##

其实，原始的思路并没有发生改变，还是用一个 BanList 去盛放哪些暂时无法访问的用户 id。 然后每次访问的时候判断一下这个用户是否在这个 List 中。

## 修改 ##

好，那我们现在需要一个结构体，因为我们会并发读取 map，所以我们直接使用 sync.Map：

` type Ban struct { M sync.Map } 复制代码`

如果你点进 sync.Map 你会发现他真正存储数据的是一个 ` atomic.Value` 。 一个具有原子特性的 interface{}。

同时Ban这个结构提还会有一个 ` IsIn` 的方法用来判断用户 id 是否在Map中。

` func (b *Ban) IsIn (user string ) bool { fmt.Printf( "%s 进来了\n" , user) // Load 方法返回两个值，一个是如果能拿到的 key 的 value // 还有一个是否能够拿到这个值的 bool 结果 v, ok := b.M.Load(user) // sync.Map.Load 去查询对应 key 的值 if !ok { // 如果没有，说明可以访问 fmt.Printf( "名单里没有 %s，可以访问\n" , user) // 将用户名存入到 Ban List 中 b.M.Store(ip, time.Now()) return false } // 如果有，则判断用户的时间距离现在是否已经超过了 180 秒，也就是3分钟 if time.Now().Second() - v.(time.Time).Second() > 180 { // 超过则可以继续访问 fmt.Printf( "时间为：%d-%d\n" , v.(time.Time).Second(), time.Now().Second()) // 同时重新存入时间 b.M.Store(ip, time.Now()) return false } // 否则不能访问 fmt.Printf( "名单里有 %s，拒绝访问\n" , user) return true } 复制代码`

下面看看测试的函数：

` func main () { var success int64 = 0 ban := new (Ban) wg := sync.WaitGroup{} // 保证程序运行完成 for i := 0 ; i < 2 ; i++ { // 我们大循环两次，每个user 连续访问两次 for j := 0 ; j < 10 ; j++ { // 人数预先设定为 10 个人 wg.Add( 1 ) go func (c int ) { defer wg.Done() ip := fmt.Sprintf( "%d" , c) if !ban.IsIn(ip) { // 原子操作增加计数器，用来统计我们人数的 atomic.AddInt64(&success, 1 ) } }(j) } } wg.Wait() fmt.Println( "此次访问量：" , success) } 复制代码`

其实测试的函数并没有做大的改动，只不过，因为我们是并发去运行的，需要增加一个 sync.WaitGroup() 保证程序完整运行完毕后才退出。

我特地把运行数值调小一点，以方便测试。 把 ` 1000` 次请求，改为 ` 2` 次。 ` 100` 人改为 ` 10` 人。

所以整个代码应该是这样的：

` package main import ( "fmt" "sync" "sync/atomic" "time" ) type Ban struct { M sync.Map } func (b *Ban) IsIn (user string ) bool { ... } func main () { ... } 复制代码`

运行一下...

诶，似乎不太对哦，发现会出现 10~15 次不等的访问量结果。为什么呢？ 寻思着，其实因为并发导致的，看到这里了吗？

` func (b *Ban) IsIn (user string ) bool { ... v, ok := b.M.Load(user) if !ok { fmt.Printf( "名单里没有 %s，可以访问\n" , user) b.M.Store(ip, time.Now()) return false } ... } 复制代码`

并发发起的 ` sync.Map.Load` 其实并没有与 ` sync.Map.Store` 连接起来形成原子操作。 所以如果有3个 user 同时进来，程序同时查询，三个返回结果都会是 false（不在名单里）。 所以也就增加了访问的数量。

其实 sync.Map 也已经考虑到了这种情况，所以他会有一个 ` LoadOrStore` 的原子方法-- 如果 Load 不出，就直接 Store，如果 Load 出来，那啥也不做。

所以我们小改一下 IsIn 的代码：

` func (b *Ban) IsIn (user string ) bool { ... v, ok := b.M.LoadOrStore(user, time.Now()) if !ok { fmt.Printf( "名单里没有 %s，可以访问\n" , user) // 删除b.M.Store(ip, time.Now()) return false } ... } 复制代码`

然后我们再运行一下，运行几次。 发觉不会再出现 此次访问量大于 10 的情况了。

## 深究一下 ##

到此为止，这个场景下的代码实现我们算是成功了。 但是真正限制用户访问的场景需求可不能这么玩，一般还是配合内存数据库去实现。

那么，如果你只想了解 sync.Map 的应用，就到这里为止了。 然而好奇心驱使我看看 sync.Map 的实现，我们继续吧。

## 制造问题 ##

如果硬是要并发读写一个 go map 会怎么样？ 试一下：

先来个主角 A

` type A map [ string ] int 复制代码`

我们定义成了自己一个类型 A，他骨子里还是 map。

` type A map [ string ] int func main () { // 初始化一个 A m := make (A) m.SetMap( "one" , 1 ) m.SetMap( "two" , 3 ) // 读取 one go m.ReadMap( "one" ) // 设置 two 值为 2 go m.SetMap( "two" , 2 ) time.Sleep( 1 *time.Second) } // A 有个读取某个 Key 的方法 func (a *A) ReadMap (key string ) { fmt.Printf( "Get Key %s: %d" ,key, a[key]) } // A 有个设置某个 Key 的方法 func (a *A) SetMap (key string , value int ) { a[key] = value fmt.Printf( "Set Key %s: %d" ,key, a[key]) // 同协程的读写不会出问题 } 复制代码`

诶，看上去不错，我们给 map A 类型定义了 get， set 方法，如果 golang 不允许并发读写 map 的话，应该会报错吧，我们跑一下。

` > Get Key one: 1 > Set Key two: 2 复制代码`

喵喵喵??? 为什么正常输出了？ 说好的并发读写报错呢？ 好吧，其实原因是上面的 map 读写，虽然我们设置了协程，但是对于计算机来说还是有时间差的。只要一个微小的先后，就不会造成 map 数据的读写异常，所以我们改一下。

` func main () { m := make (A) m[ "one" ] = 1 m[ "two" ] = 3 go func () { for { m.ReadMap( "one" ) } }() go func () { for { m.SetMap( "two" , 2 ) } }() time.Sleep( 1 *time.Second) } 复制代码`

为了让读写能够尽可能碰撞，我们增加了循环。 现在我们可以看到了：

` > fatal error: concurrent map read and map write 复制代码`

*这里之所以为有 panic 是因为在 map 实现中进行了 ` 并发读写的检查` 。

## 解决问题 ##

其实上面的例子和 go 对 sync.Mutex 锁的入门教程很像。 我们证实了 map 并发读写的问题，现在我们尝试来解决。

既然是读写造成的冲突，那我们首先考虑的便是加锁。 我们给读取一个锁，写入一个锁。那么我们现在需要讲单纯的 A map 转换成一个带有锁的结构体：

` type A struct { Value map [ string ] int mu sync.Mutex } 复制代码`

Value 成了真正存放我们值的地方。 我们还要修改下 ` ReadMap` 和 ` SetMap` 两个方法。

` func (a *A) ReadMap (key string ) { a.mu.Lock() fmt.Printf( "Get Key %s: %d" ,key, a.Value[key]) a.mu.Unlock() } func (a *A) SetMap (key string , value int ) { a.mu.Lock() a.Value[key] = value a.mu.Unlock() fmt.Printf( "Set Key %s: %d" ,key, a.Value[key]) } 复制代码`

注意，这里两个方法中，哪一个少了 Lock 和 Unlock 都不行。

我们再跑一下代码，发现可以了，不会报错。

## 到此为止了吗？ ##

我们算是用最简单的方法解决了眼前的问题，但是这样真的没问题吗？ 细心的你会发现，读写我们都加了锁，而且没有任何特殊条件限制，所以当我们要多次读取 map 中的数据的时候，他喵的都会阻塞！就算我压根不想改 map 中的 value... 尽管现在感觉不出来慢，但这对密集读取来说是一个性能坑。

为了避免不必要的锁，我们似乎还要让代码“聪明些”。

## 读写分离 ##

没错，读写分离就是一个十分适用的设计思路。 我们准备一个 Read map，一个 Write map。

但这里的读写分离和我们平时说的不太一样（记住我们的场景永远是并发读写），我们不能实时或者定时让写入的 map 去同步（增删改）到读取的 map 中， 因为...这样和上面的 map 操作没有任何区别，因为读取 map 没有锁，还是会发生并发冲突。

我们要解决的是，不“显示”增删改 map 中的 key 对应的 value。 我们把问题再分类一下：

* 修改（删除）已有的 key 的 value
* 增加不存在的 key 和 value

` 第一个问题：` 我们把 key 的 value 变成指针怎么样？ 相同的 key 指向同一个指针地址，指针地址又指向真正值的地址。

> 
> 
> 
> key -> &地址 -> &真正的值地址
> 
> 

Read 和 Write map 的值同时指向一个 ` &地址` ，不论谁改，大家都会变。 当我们需要修改已有的 key 对应的 value 时，我们修改的是 ` &真正的值地址` 的值，并不会修改 key 对应的 ` &地址` 或值。 同时，通过 ` atomic` 包，我们能够做到对指针修改的原子性。 太棒了，修改已有的 key 问题解决。

` 第二个问题：` 因为并不存在这个 key，所以我们一定会增加新 key， 既然我们有了 Read map & Write map，那我们可以利用起来呀， 我们在 Write map 中加锁并增加这个 key 和对应的值，这样不影响 Read map 的读取。

不过，Read map 我们终究是要更新的，所以我们加一个计数器 ` misses` ，到了一定条件，我们把 Write map 安全地同步到 Read map 中，并且清空 Write map。

Read map 可以看做是 Write map 的一个只读拷贝，不允许自行增加新 key，只能读或者改。

上面的思想其实和 sync.Map 的实现离得很近了。 只不过，sync.Map 把我们的 ` Write map` 叫做了 ` dirty` ，把 Write map ` 同步` 到 Read map 叫做了 ` promote（升级）` 。 又增加了一些结构体封装，和状态标识。

其实 google 一下你就会发现很多分析 sync.Map 源码的文章，都写得很好。我这里也不赘述了，但是我想用我的理解去概括下 sync.Map 中的方法思路。

结合 sync.Map 源码食用味道更佳。

### 读取 Map.Load ###

* Read map 直接读得到吗？
* 么有？好吧，我们上锁，再读一次 Read map
* 还没有？那我只能去读 Dirty map 了
* 读到了，不错，我们记录下这次读取属于 ` 未命中` （misses + 1），顺便看看我们的 dirty 是不是可以升级成 Read 了
* 解锁

*这里2中之所以再上锁，是为了double-checking，防止在极小的时间差内产生脏读（dirty突然升级 Read）。

### 写入 Map.Store ###

* Read map 有没有这个 key ？
* 有，那我们原子操作直接修改值指针呗
* 没有？依旧上锁再看看有没有？
* 还没有，好吧，看看 Dirty map
* 有诶！那就修改 Dirty map 这个 key 的值指针
* 没有？那就要在 Dirty map 新增一个 key 咯，为了方便之后 Dirty map 升级成 Read map，我们还要把原先的 Read map 全复制过来
* 解锁

### 删除 Map.Delete ###

* Read map 有这个 key 吗？
* 有啊，那就把 value 直接改成 nil（防止之后读取没有 key 还要去加锁，影响性能）
* 没有？直接删 dirty 里的这个 key 吧

### 读取或者存 Map.LoadOrStore ###

emmmm......

* Map.Load + Map.Store

## 编不下去了 ##

大致就是这样的思路，我这里再推荐一些正统的源码分析和基准测试，相信看完以后会对 sync.Map 更加清晰。

* [不得不知道的Golang之sync.Map源码分析]( https://juejin.im/post/5b1b3d785188257d45297d0a )
* [Go sync.Map 实现]( https://link.juejin.im?target=http%3A%2F%2Fwudaijun.com%2F2018%2F02%2Fgo-sync-map-implement%2F )
* [Go 1.9 sync.Map揭秘]( https://link.juejin.im?target=https%3A%2F%2Fcolobu.com%2F2017%2F07%2F11%2Fdive-into-sync-Map%2F )

另外，如果你注意到 Map.Store 中第6步的 ` 全部复制` 的话，你就会有预感，sync.Map 的使用场景其实不太适合高并发写的逻辑。 的确，官方说明也明确指出了 sync.Map 适用的场景：

` // Map is like a Go map[interface{}]interface{} but is safe for concurrent use // by multiple goroutines without additional locking or coordination. ... // The Map type is specialized. Most code should use a plain Go map instead, // with separate locking or coordination, for better type safety and to make it // easier to maintain other invariants along with the map content. // // The Map type is optimized for two common use cases: (1) when the entry for a given // key is only ever written once but read many times , as in caches that only grow, // or (2) when multiple goroutines read , write, and overwrite entries for disjoint // sets of keys. In these two cases, use of a Map may significantly reduce lock // contention compared to a Go map paired with a separate Mutex or RWMutex. 复制代码`

sync.Map 只是帮助优化了两个使用场景：

* 多读少写
* 多 goroutine 操作键值

其实 sync.Map 还是在性能和安全之间，找了一个自己觉得合适的平衡点，就如同我们开头的案例一样，其实 sync.Map 也并不适用。 另外，这里有一个 [【sync.Map 的进阶版本】]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Forcaman%2Fconcurrent-map ) 。

## *atomic 和 mutex ##

其实在很久以前翻看 sync Map 源码的时候，我不经会抛出疑问，如果能够用 atomic 来解决并发安全问题，为什么还要 mutex 呢？ 而且，在进行 map.Store 的过程中，还是会直接修改 read 的 key 所对应的值（并且无锁状态），这和普通修改一个 key 的值有什么区别呢？

如果 atomic 可以保证原子性，那和 mutex 有什么区别呢？ 在翻查了一些资料后，我知道了：

> 
> 
> 
> Mutexes are slow, due to the setup and teardown, and due to the fact that
> they block other goroutines for the duration of the lock.
> 
> 

> 
> 
> 
> Atomic operations are fast because they use an atomic CPU instruction,
> rather than relying on external locks to.
> 
> 

互斥锁其实是通过阻塞其他协程起到了原子操作的功能，但是 atomic 是通过控制更底层的 CPU 指令，来达到值操作的原子性的。

所以 atomic 和 mutex 并不是一个层面的东西，而且在专职点上也不尽相同，mutex 很多地方也是通过 atomic 去实现的。

而 sync Map 很巧妙地将两个结合来实现并发安全。

* 

它用一个指针来存储 key 对应的 value，当要修改的时候只是修改 value 的地址（并且是地址值的副本操作），这个可以通过 atomic 的 Pointer 操作来实现，并且不会又冲突。

* 

另外，又使用了读写分离+mutex互斥锁，来控制 增删改查 key 的操作，防止冲突。

其中第一点是很多源码解读中常常一笔带过的，然而萌新我觉得反而是相当重要的技巧（捂脸）。

## *一些疑问 ##

### misses 条件 ###

一直没有明白，为什么从 dirty map 升级成 read map 的条件是 ` misses 次数大于等于 len(m.dirty)` ？

### Go map 为什么不加锁？ ###

我们可以看到下面两篇关于不加锁的叙述：

* [golang.org/doc/faq#ato…]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Ffaq%23atomic_maps )
* [blog.golang.org/go-maps-in-…]( https://link.juejin.im?target=https%3A%2F%2Fblog.golang.org%2Fgo-maps-in-action )

## *参考 ##

* [www.jianshu.com/p/aa0d4808c…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Faa0d4808cbb8 )
* [www.zhihu.com/question/26…]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F266414962 )
* [juejin.im/post/5ae014…]( https://juejin.im/post/5ae01447f265da0ba062d2e8 )
* [juejin.im/post/5b1b3d…]( https://juejin.im/post/5b1b3d785188257d45297d0a )
* [wudaijun.com/2018/02/go-…]( https://link.juejin.im?target=http%3A%2F%2Fwudaijun.com%2F2018%2F02%2Fgo-sync-map-implement%2F )
* [colobu.com/2017/07/11/…]( https://link.juejin.im?target=https%3A%2F%2Fcolobu.com%2F2017%2F07%2F11%2Fdive-into-sync-Map%2F )