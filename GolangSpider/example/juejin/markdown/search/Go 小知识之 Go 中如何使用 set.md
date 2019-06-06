# Go 小知识之 Go 中如何使用 set #

今天来聊一下 Go 如何使用 set，本文将会涉及 set 和 bitset 两种数据结构。

# Go 的数据结构 #

Go 内置的数据结构并不多。工作中，我们最常用的两种数据结构分别是 slice 和 map，即切片和映射。 其实，Go 中也有数组，切片的底层就是数组，只不过因为切片的存在，我们平时很少使用它。

除了 Go 内置的数据结构，还有一些数据结构是由 Go 的官方 container 包提供，如 heap 堆、list 双向链表和ring 回环链表。但今天我们不讲它们，这些数据结构，对于熟手来说，看看文档就会使用了。

我们今天将来聊的是 set 和 bitset。据我所知，其他一些语言，比如 Java，是有这两种数据结构。但 Go 当前还没有以任何形式提供。

# 实现思路 #

先来看一篇文章，访问地址 [2 basic set implementations]( https://link.juejin.im?target=https%3A%2F%2Fyourbasic.org%2Fgolang%2Fimplement-set%2F ) 阅读。文中介绍了两种 go 实现 set 的思路， 分别是 map 和 bitset。

有兴趣可以读读这篇文章，我们接下来具体介绍下。

# map #

我们知道，map 的 key 肯定是唯一的，而这恰好与 set 的特性一致，天然保证 set 中成员的唯一性。而且通过 map 实现 set，在检查是否存在某个元素时可直接使用 _, ok := m[key] 的语法，效率高。

先来看一个简单的实现，如下：

` set := make ( map [ string ] bool ) // New empty set set[ "Foo" ] = true // Add for k := range set { // Loop fmt.Println(k) } delete (set, "Foo" ) // Delete size := len (set) // Size exists := set[ "Foo" ] // Membership 复制代码`

通过创建 map[string]bool 来存储 string 的集合，比较容易理解。但这里还有个问题，map 的 value 是布尔类型，这会导致 set 多占一定内存空间，而 set 不该有这个问题。

怎么解决这个问题？

设置 value 为空结构体，在 Go 中，空结构体不占任何内存。当然，如果不确定，也可以来证明下这个结论。

` unsafe.Sizeof( struct {}{}) // 结果为 0 复制代码`

优化后的代码，如下：

` type void struct{} var member void set := make(map[string]void) // New empty set set [ "Foo" ] = member // Add for k := range set { // Loop fmt.Println(k) } delete( set , "Foo" ) // Delete size := len( set ) // Size _, exists := set [ "Foo" ] // Membership 复制代码`

之前在网上看到有人按这个思路做了封装，还写了 [一篇文章]( https://link.juejin.im?target=https%3A%2F%2Fallenwu.itscoder.com%2Fset-in-go ) ，可以去读一下。

其实，github 上已经有个成熟的包，名为 golang-set，它也是采用这个思路实现的。访问地址 [golang-set]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdeckarep%2Fgolang-set ) ，描述中说 Docker 用的也是它。包中提供了两种 set 实现，线程安全的 set 和非线程安全的 set。

演示一个简单的案例。

` package main import ( "fmt" mapset "github.com/deckarep/golang-set" ) func main () { // 默认创建的线程安全的，如果无需线程安全 // 可以使用 NewThreadUnsafeSet 创建，使用方法都是一样的。 s1 := mapset.NewSet( 1 , 2 , 3 , 4 ) fmt.Println( "s1 contains 3: " , s1.Contains( 3 )) fmt.Println( "s1 contains 5: " , s1.Contains( 5 )) // interface 参数，可以传递任意类型 s1.Add( "poloxue" ) fmt.Println( "s1 contains poloxue: " , s1.Contains( "poloxue" )) s1.Remove( 3 ) fmt.Println( "s1 contains 3: " , s1.Contains( 3 )) s2 := mapset.NewSet( 1 , 3 , 4 , 5 ) // 并集 fmt.Println(s1.Union(s2)) } 复制代码`

输出如下：

` s1 contains 3: true s1 contains 5: false s1 contains poloxue: true s1 contains 3: false Set{4, polxue, 1, 2, 3, 5} 复制代码`

例子中演示了简单的使用方式，如果有不明白的，看下源码，这些数据结构的操作方法名都是很常见的，比如交集 Intersect、差集 Difference 等，一看就懂。

# bitset #

继续聊聊 bitset，bitset 中每个数子用一个 bit 即能表示，对于一个 int8 的数字，我们可以用它表示 8 个数字，能帮助我们大大节省数据的存储空间。

bitset 最常见的应用有 bitmap 和 flag，即位图和标志位。这里，我们先尝试用它表示一些操作的标志位。比如某个场景，我们需要三个 flag 分别表示权限1、权限2和权限3，而且几个权限可以共存。我们可以分别用三个常量 F1、F2、F3 表示位 Mask。

示例代码如下（引用自文章 [Bitmasks, bitsets and flags]( https://link.juejin.im?target=https%3A%2F%2Fyourbasic.org%2Fgolang%2Fbitmask-flag-set-clear%2F ) ）：

` type Bits uint8 const ( F0 Bits = 1 << iota F1 F2 ) func Set (b, flag Bits) Bits { return b | flag } func Clear (b, flag Bits) Bits { return b &^ flag } func Toggle (b, flag Bits) Bits { return b ^ flag } func Has (b, flag Bits) bool { return b&flag != 0 } func main () { var b Bits b = Set(b, F0) b = Toggle(b, F2) for i, flag := range []Bits{F0, F1, F2} { fmt.Println(i, Has(b, flag)) } } 复制代码`

例子中，我们本来需要三个数才能表示这三个标志，但现在通过一个 uint8 就可以。bitset 的一些操作，如设置 Set、清除 Clear、切换 Toggle、检查 Has 通过位运算就可以实现，而且非常高效。

bitset 对集合操作有着天然的优势，直接通过位运算符便可实现。比如交集、并集、和差集，示例如下：

* 交集：a & b
* 并集：a | b
* 差集：a & (~b)

底层的语言、库、框架常会使用这种方式设置标志位。

以上的例子中只展示了少量数据的处理方式，uint8 占 8 bit 空间，只能表示 8 个数字。那大数据场景能否可以使用这套思路呢？

我们可以把 bitset 和 Go 中的切片结合起来，重新定义 Bits 类型，如下：

` type Bitset struct { data []int64 } 复制代码`

但如此也会产生一些问题，设置 bit，我们怎么知道它在哪里呢？仔细想想，这个位置信息包含两部分，即保存该 bit 的数在切片索引位置和该 bit 在数字中的哪位，分别将它们命名为 index 和 position。那怎么获取？

index 可以通过整除获取，比如我们想知道表示 65 的 bit 在切片的哪个 index，通过 65 / 64 即可获得，如果为了高效，也可以用位运算实现，即用移位替换除法，比如 65 >> 6，6 表示移位偏移，即 2^n = 64 的 n。

postion 是除法的余数，我们可以通过模运算获得，比如 65 % 64 = 1，同样为了效率，也有相应的位运算实现，比如 65 & 0b00111111，即 65 & 63。

一个简单例子，如下：

` package main import ( "fmt" ) const ( shift = 6 mask = 0x3f // 即0b00111111 ) type Bitset struct { data [] int64 } func NewBitSet (n int ) * Bitset { // 获取位置信息 index := n >> shift set := &Bitset{ data: make ([] int64 , index+ 1 ), } // 根据 n 设置 bitset set.data[index] |= 1 << uint (n&mask) return set } func (set *Bitset) Contains (n int ) bool { // 获取位置信息 index := n >> shift return set.data[index]&( 1 << uint (n&mask)) != 0 } func main () { set := NewBitSet( 65 ) fmt.Println( "set contains 65" , set.Contains( 65 )) fmt.Println( "set contains 64" , set.Contains( 64 )) } 复制代码`

输出结果

` set contains 65 true set contains 64 false 复制代码`

以上的例子功能很简单，只是为了演示，只有创建 bitset 和 contains 两个功能，其他诸如添加、删除、不同 bitset 间的交、并、差还没有实现。有兴趣的朋友可以继续尝试。

其实，bitset 包也有人实现了，github地址 [bit]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyourbasic%2Fbit ) 。可以读读它的源码，实现思路和上面介绍差不多。

下面是一个使用案例。

` package main import ( "fmt" "github.com/yourbasic/bit" ) func main () { s := bit.New( 2 , 3 , 4 , 65 , 128 ) fmt.Println( "s contains 65" , s.Contains( 65 )) fmt.Println( "s contains 15" , s.Contains( 15 )) s.Add( 15 ) fmt.Println( "s contains 15" , s.Contains( 15 )) fmt.Println( "next 20 is " , s.Next( 20 )) fmt.Println( "prev 20 is " , s.Prev( 20 )) s2 := bit.New( 10 , 22 , 30 ) s3 := s.Or(s2) fmt.Println( "next 20 is " , s3.Next( 20 )) s3.Visit( func (n int ) bool { fmt.Println(n) return false // 返回 true 表示终止遍历 }) } 复制代码`

执行结果：

` s contains 65 true s contains 15 false s contains 15 true next 20 is 65 prev 20 is 15 next 20 is 22 2 3 4 10 15 22 30 65 128 复制代码`

代码的意思很好理解，就是一些增删改查和集合的操作。要注意的是，bitset 和前面的 set 的区别，bitset 的成员只能是 int 整型，没有 set 灵活。平时的使用场景也比较少，主要用在对效率和存储空间要求较高的场景。

# 总结 #

本文介绍了Go 中两种 set 的实现原理，并在此基础介绍了对应于它们的两个包简单使用。我觉得，通过这篇文章，Go 中 set 的使用，基本都可以搞定了。

除这两个包，再补充两个。 [zoumo/goset]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzoumo%2Fgoset ) 和 [github.com/willf/bitse…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fwillf%2Fbitset ) 。