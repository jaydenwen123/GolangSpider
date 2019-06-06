# Golang 深入 slice 实现原理及使用技巧 #

slice 使用总结，持续更新于我的 [Github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flvgithub%2Fgo_blog%2Fblob%2Fmaster%2FSlice%2Fslice.md )

介绍

我们都知道array是固定长度的数组, slice是对array的扩展,本质上是基于数组实现的,主要特点是定义完一个slice变量之后，不需要为它的容量而担心。 本文记录直接深入slice的底层实现原理，不再介绍slice的基本使用。

slice 结构

* slice中 array 是一个指针，它指向的是一个array
* len 代表的是这个slice中的元素长度
* cap 是slice的容量
* [参考 Golang slice 源码]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fsrc%2Fruntime%2Fslice.go%3Fh%3Dtype%2520slice%2520struct%23L13 ) ` type slice struct { array unsafe.Pointer len int cap int } 复制代码`

slice 扩容

` s := []int{1,2,3,4,5,6} s = append(s, 6) 复制代码`

* 如果新的slice大小是当前大小2倍以上，则大小增长为新大小
* 如果当前slice cap 小于1024，按每次2倍增长，否则每次按当前大小1/4增长。直到增长的大小超过或等于新大小
* append的实现是在内存中将slice的array值赋值到新申请的array上
* [扩容源码实现]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fsrc%2Fruntime%2Fslice.go%3Fh%3Dgrowslice%23L77 )

性能

* 通过上面我们知道slice的扩容涉及到内存的拷贝，这样带来的好处是数据存储在连续内存上，比随机访问快很多，最直接的性能提升就是缓存命中率会高很多,这也就是为什么slice不采用动态链表实现的原因吧
* 我们知道拷贝内存数据是有开销的， 而其中最大的开销不在 memmove 数据上，而是在开辟一块新内存malloc及之后的GC压力
* 拷贝连续内存是很快的，随着cap变大，拷贝总成本还是 O(N) ,只是常数大了
* 假如不想发生拷贝，那你就没有连续内存。此时随机访问开销会是：链表 O(N), 2倍增长块链 O(LogN), 二级表一个常数很大的 O(1)
* 当你能大致知道所需的最大空间（在大部分时候都是的）时，在make的时候预留相应的 cap 就好
* 如果需要的空间很大，而且每次都不确定，那就要在浪费内存和耗 CPU 在 malloc + gc 上做权衡
* 链表的查找操作是从第一个元素开始，所以相对数组要耗时间的多，因为采用这样的结构对读的性能有很大的提高

选择

* slice是很灵活的,大部分情况都能表现的很好
* 但也有特殊情况,slice的容量超大并且需要频繁的更改slice的内容时,改用list更合适

注意点

如果你理解了上面内容，那下面这段代码的输出结果你就不意外了

` s := [] byte { 1 , 23 , 4 , 5 , 67 , 7 } s1 := s[ 2 : 3 ] s1[ 0 ] = 100 fmt.Printf( "s:%+v\n" , s) // s:[1 23 100 5 67 7] 复制代码`

没错，切片s 第三位的值4被替换为了100,这是因为 切片s1 的底层array指针指向 切片s 的第三位，因此操作s1会影响切片s