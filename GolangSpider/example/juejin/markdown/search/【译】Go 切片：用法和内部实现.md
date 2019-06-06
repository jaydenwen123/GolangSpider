# 【译】Go 切片：用法和内部实现 #

> 
> 
> 
> * 原文标题： [Go Slices: usage and internals](
> https://link.juejin.im?target=https%3A%2F%2Fblog.golang.org%2Fgo-slices-usage-and-internals
> )
> * 原文作者：Andrew Gerrand
> * 原文时间：2011-01-05
> 
> 
> 

Go 的切片（ ` slice` ）提供了一种方便、高效的处理特定类型数据序列的方法。切片类似于其他语言中的数组，但有些特别的地方。本文讨论切片是什么、以及如何使用它。

## 数组 ##

Go 中切片是基于数组的，因此为了理解切片，首先得理解数组。
一个数组类型包括元素类型和元素个数。例如，类型 ` [4]int` 表示有 4 个整数元素的数组。一个数组的长度是固定的，长度本身也是类型的一部分（ ` [4]int` 和 ` [5]int` 是两个不同的类型）。数组能通过下标访问，因此表达式 ` s[n]` 表示访问下标从 0 开始的第 n 个元素。

` var a [ 4 ] int a[ 0 ] = 1 i := a[ 0 ] // i == 1 复制代码`

数组不需要显示初始化，它会自动初始化为数组零值（the zero value of an array），这个零值中的所有的元素的值都是该元素类型的零值：

` // a[2] == 0, 初始化为 int 类型的零值 复制代码`

类型 ` [4]int` 的内存表示就是 4 个整数顺序摆放：

![内存摆放](https://user-gold-cdn.xitu.io/2019/4/4/169e83fd09b8dfae?imageView2/0/w/1280/h/960/ignore-error/1)
Go 中数组是值类型。一个数组变量表示整个数组，而不是指向数组第一个元素的指针（C 语言是这样的）。这表明对一个数组进行赋值或传递，会复制整个数组内容。（你可以传递数组指针来避免内容复制，但这是一个数组指针，不是数组）一种理解数组的方法是把它看成一种 ` struct` ，通过下标来使用，而不是成员名，这是一种固定大小的复合值。
数组字面量（ ` literal` ）能这样写：

` b := [ 2 ] string { "Penn" , "Teller" } 复制代码`

或者，让编译器计算元素的个数

` b := [...] string { "Penn" , "Teller" } 复制代码`

上面两种情况，变量 ` b` 的类型都是 ` [2]string` 。

## 切片 ##

数组用自己的用武之地，但不灵活，所以并不经常在 Go 代码中出现。与它对比，切片就常用的多。它基于数组，但非常方便和强大。
切片类型表示为 ` []T` ，其中 ` T` 是切片中元素的类型。与数组类型不同，切片类型不用指定长度。
一个切片的字面量和数组类似，除了不能指定元素个数：

` letters := [] string { "a" , "b" , "c" , "d" } 复制代码`

可以使用内建函数 ` make` 来创建切片， ` make` 函数签名如下：

` func make ([]T, len , cap ) [] T 复制代码`

其中 ` T` 表示新切片中元素的类型。 ` make` 参数有：切片元素类型、切片长度、切片容量（可选）。调用 ` make` 时，它会申请一个数组，然后返回一个使用该数组的切片。

` var s [] byte s = make ([] byte , 5 , 5 ) // s == []byte{0, 0, 0, 0} 复制代码`

如果不指定切片容量，默认为与切片长度一样大小。下面是一个更简单的版本：

` s := make ([] byte , 5 ) 复制代码`

使用内建函数 ` len` 和 ` cap` 来获取切片的长度和容量。

` len (s) == 5 cap (s) == 5 复制代码`

接下来两部分讨论切片长度和切片容量的关系。
切片的零值为 ` nil` ，这种情况下 ` len` 和 ` cap` 都返回 0。
对一个已有的数组或切片进行切片操作（译者：注意 ` 切片` 和 ` 切片操作` 的区别），能生成一个新的切片。切片操作通过两个下标中间加个冒号来指定一个半开的区间。比如，表达式 ` b[1:4]` 创建了一个新切片，新切片包括 ` b` 中下标为 1、2、3 的元素（新切片中对应的下标为0、1、2）。

` b := [] byte { 'g' , 'o' , 'l' , 'a' , 'n' , 'g' } // b[1:4] == []byte{'o', 'l', 'a'}, 与b共用同一内存空间 复制代码`

表达式中开始和结束的下标都是可选的，默认分别为 0 和切片的长度。

` // b[:2] == []byte{'g', 'o'} // b[2:] == []byte{'l', 'a', 'n', 'g'} // b[:] == b 复制代码`

下面是通过数组来创建切片：

` x := [ 3 ] string { "Лайка" , "Белка" , "Стрелка" } s := x[:] // 切片 s 使用数组 x 的内存空间 复制代码`

## 切片的内部实现 ##

切片是数组某段的描述符，包括一个 ` 指向数组的指针` ，当前段的 ` 长度（length）` ，还有 ` 容量（capacity）` （这个段能到达的最大长度）。

![切片结构](https://user-gold-cdn.xitu.io/2019/4/4/169e841a14138206?imageView2/0/w/1280/h/960/ignore-error/1)
上面使用 ` make([]byte, 5)` 得到的变量 ` s` ，它的结构如下： ![变量s的结构](https://user-gold-cdn.xitu.io/2019/4/4/169e841a0e773914?imageView2/0/w/1280/h/960/ignore-error/1)
其中的 ` 长度` 表示当前切片中元素的个数。 ` 容量` 表示底层数组的元素个数（该数组的起始地址为切片中的指针值）。下面的几个例子能帮你更佳清晰理解 ` 长度` 和 ` 容量` 的区别。
对上面变量 ` s` 进行切片操作，观察它的结构变化，还有和底下数组的关系：

` s = s[ 2 : 4 ] 复制代码`

![变量s的结构改变](https://user-gold-cdn.xitu.io/2019/4/4/169e841a1d0f6758?imageView2/0/w/1280/h/960/ignore-error/1)
切片操作并不会复制原始切片或数组的数据，而是将新切片的数据指针指向原始数据。这使得切片操作能像操作数组索引一样高效。也因此，修改新切片的元素，原始切片也会被修改：

` d := [] byte { 'r' , 'o' , 'a' , 'd' } e := d[ 2 :] // e == []byte{'a', 'd'} e[ 1 ] = 'm' // e == []byte{'a', 'm'} // d == []byte{'r', 'o', 'a', 'm'} 复制代码`

上面我们对 ` s` 进行了切片操作，使得 ` s` 的长度比其容量小。能再次通过切片操作增加 ` s` 的长度，使得长度和容量相等。

` s = s[: cap (s)] 复制代码`

![恢复s的长度](https://user-gold-cdn.xitu.io/2019/4/4/169e841a1ebb44cd?imageView2/0/w/1280/h/960/ignore-error/1)
切片的长度不能超过其容量，如果尝试这么做会到造成一个运行时错误（runtime panic），就如同数组或切片下标越界一样。同样的，对切片进行切片操作时，参数不能小于0（想要访问底下数组之前的元素）。

## 切片增长（复制和添加元素） ##

想要增加切片的容量，只能创建一个新的，容量更大的切片，然后把原始数据复制过去。其他语言的动态数组的实现也是使用的这种幕后技术。下面代码的操作：创建一个容量为 ` s` 的两倍的新切片 ` t` ，复制 ` s` 的内容到 ` t` ，最后将 ` t` 赋值给 ` s` 。

` t := make ([] byte , len (s), ( cap (s)+ 1 )* 2 ) // +1 防止 cap(s) == 0 的情况 for i := range s { t[i] = s[i] } s = t 复制代码`

内建函数 ` copy` 实现了上面代码中循环的功能。它将要复制的内容从原始切片拷贝到目的切片，返回拷贝元素的个数。

` func copy (dst, src []T) int 复制代码`

` copy` 支持不同长度的切片间的拷贝（拷贝长度为两切片中长度小的那个）。而且，它还能正确处理源切片和目的切片处于同一个数组上的情况。
使用 ` copy` 上面的代码简化为：

` t = make ([] byte , len (s), ( cap (s)+ 1 )* 2 ) copy (t, s) s = t 复制代码`

将数据添加到切片的末尾是很常用的操作。下面这个函数会将一个 ` byte` 元素添加到元素类型为 ` byte` 的切片中。如果有必要，它会增加切片的大小。最后返回添加后的切片。

` func AppendByte (slice [] byte , data ... byte ) [] byte { m := len (slice) n := m + len (data) if n > cap (slice) { //重新申请空间 // 申请两倍的空间，以备后用 newSlice := make ([] byte , (n+ 1 )* 2 ) copy (newSlice, slice) slice = newSlice } slice = slice[ 0 :n] copy (slice[m:n], data) return slice } 复制代码`

` AppendByte` 函数用法如下：

` p := [] byte { 2 , 3 , 5 } p = AppendByte(p, 7 , 11 , 13 ) // p == []byte{2, 3, 5, 7, 11, 13} 复制代码`

类似于 ` AppendByte` 的函数是很有用的，因为它提供了完全掌控切片增长的方式。根据不同程序的不同特性，能调整分配更大或更小的空间，或者设置一个分配空间上限。
但大部分程序不需要这样的完全掌控，所以 Go 提供了一个内建函数实现这个功能。太部分情况下都很好用，函数签名如下：

` func append (s []T, X ...T) [] T 复制代码`

` apppend` 将元素 ` x` 添加到切片 ` s` 的末尾，如有必要，它会增加切片的容量。

` a := make ([] int , 1 ) // a == []int{0} a = append (a, 1 , 2 , 3 ) // a == []int{0, 1, 2, 3} 复制代码`

将一个切片添加到另一个切片，使用操作符 `...` 将第二切片展开成参数列表。

` a := [] string { "John" , "Paul" } b := [] string { "George" , "Ringo" , "pete" } a = append (a, b...) // 等同于 ”append(a, b[0], b[1], b[2])" // a == []string{"John", "Paul", "George", "Ringo", "Pete"} 复制代码`

因为切片的零值（ ` nil` ）有类似于长度为零的切片的属性，因此可以直接声明一个变量，向其添加元素：

` // Filter 函数返回一个包含满足 fn() 的元素的新切片 func Filter (s [] int , fn func ( int ) bool ) [] int { var p [] int // == nil for _, v := range s { if fn(v) { p = append (p, v) } } return p } 复制代码`

## 一个可能的陷阱（A possible "gotcha"） ##

前面提到，对切片进行切片操作不会复制切片结构里面的数组数据。整个数组会一直占用内存空间，直到引用数为零。有些情况下，这会造成一个问题：程序只需要用到一块数据中一小段，却得把整个数据块保留在内存中。
例如，下面的函数加载一个文件进内存，查找第一组连续数字的序列作为新的切片返回。

` var digitRegexp = regexp.MustCompile( "[0-9]+" ) func FindDigits (filename string ) [] byte { b, _ := ioutil.ReadFile(filename) return digitRegexp.Find(b) } 复制代码`

这段代码可以正确运行，但是有个问题：返回的切片使用了包含整个文件的数组。因为返回的切片引用了原始数组，只要切片还在，原始数组就不能被垃圾回收 -- 对文件某一小段的使用使得整个文件都必须占用内存。
为了解决这个问题，可以在返回之前将目标数据复制到一个新切片中：

` func CopyDigits (filename string ) [] byte { b, _ := ioutil.ReadFile(filename) b = digitRegexp.Find(b) c := make ([] byte , len (b)) copy (c, b) return c } 复制代码`

这个函数另外一个更简洁的版本是使用 ` append` ，这是作为一个练习，留个读者完成。

## 更多资料 ##

[Effective Go]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Feffective_go.html ) 包含了对 [切片]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Feffective_go.html%23slices ) 和 [数组]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Feffective_go.html%23arrays ) 的更深入的讨论， [Go language specification]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Fgo_spec.html ) 定义了 [切片]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Fgo_spec.html%23Slice_types ) 以及相关的 [辅助函数]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fdoc%2Fgo_spec.html%23Appending_and_copying_slices ) 。

（原文完）

## 译者总结 ##

只用切片不用数组 : )