# 详解go语言的array和slice 【一】 #

本篇会详细讲解go语言中的array和slice，和平时开发中使用时需要注意的地方，以免入坑。

Go语言中array是一组定长的同类型数据集合，并且是连续分配内存空间的。

### 声明一个数组 ###

` var arr [3]int 复制代码`

数组声明后，他包含的类型和长度都是不可变的.如果你需要更多的元素，你只能重新创建一个足够长的数组，并把原来数组的值copy过来。

在Go语言中，初始化一个变量后，默认把变量赋值为指定类型的zero值，如string 的zero值为"" number类型的zero值为0.数组也是一样的，声明完一个数组后，数组中的每一个元素都被初始化为相应的zero值。如上面的声明是一个长度为5的int类型数组。数组中的每一个元素都初始化为int类型的zero值 0

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbc79a7ae9?imageView2/0/w/1280/h/960/ignore-error/1)

可以使用array的字面量来快速创建和初始化一个数组，array的字面量允许你设置array的长度和array中元素的值

` arr := [3]{1, 2, 3} 复制代码`

如果[]中用...来代替具体的长度值，go会根据后面初始化的元素长度来计算array的长度

` arr := [...]{1, 2, 3, 4} 复制代码`

如果你想只给某些元素赋值，可以这样写

` arr := [5]int {1: 5, 3: 200} 复制代码`

上面的语法是创建了一个长度为5的array，并把index为1的元素赋值为0，index为3的元素赋值为200，其他没有初始化的元素设置为他们的zero值

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbc7a04bf0?imageView2/0/w/1280/h/960/ignore-error/1)

### 指针数组 ###

声明一个包含有5个整数指针类型的数组，我们可以在初始化时给相应位置的元素默认值。下面是给索引为0的元素一个新建的的int类型指针(默认为0)，给索引为1的元素指向值v的地址，剩下的没有指定默认值的元素为指针的zero值也就是nil

` var v int = 6 array := [5]*int{0: new(int), 1: &v} fmt.Println(len(array)) fmt.Println(*array[0]) fmt.Println(*array[1]) v2 := 7 array[2] = &v2 fmt.Println( "------------------" ) for i, v := range array { fmt.Printf( "index %d, address %v value is " , i, v) if v != nil { fmt.Print(*v) } else { fmt.Print( "nil" ) } fmt.Println( " " ) } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbc7d87807?imageView2/0/w/1280/h/960/ignore-error/1)

### 数组做为函数参数 ###

比如我个创建一个100万长度的int类型数组，在64位机上需要在内存上占用8M的空间，把他做为一个参数传递到一个方法内，go会复制这个数组，这将导致性能的下降。

` package main import ( "fmt" ) const size int = 1000*1000 func sum(array [size]int) float 64 { total := 0.0 for _, v := range array { total += float 64(v) } return total } func main () { var arr [size]int fmt.Println(sum(arr)) } 复制代码`

当然go也提供了其他的方式，可以用指向数组的指针做为方法的参数，这样在传参的时候会传递array的地址，只需要复制8个字节，

` package main import ( "fmt" ) const size int = 1000*1000 func sum(array *[size]int) float 64 { total := 0.0 for _, v := range array { total += float 64(v) } return total } func main () { var arr [size]int fmt.Println(sum(&arr)) } 复制代码`

### slice ###

slice可以被认为动态数组，在内存中也是连续分配的。他可以动态的调整长度，可以通过内置的方法append来自动的增长slice长度;也可以通过再次切片来减少slice的长度。

slice的内部结构有3个字段，分别是维护的底层数组的指针，长度(元素个数)和容量(元素可增长个数，不足时会增长)，下面我们定义一个有2个长度，容量为5的slice

` func main () { s := make([]int, 2, 5) fmt.Println( "len: " , len(s)) fmt.Println( "cap: " , cap (s)) s = append(s, 2) fmt.Println( "--------" ) fmt.Println( "len: " , len(s)) fmt.Println( "cap: " , cap (s)) fmt.Println( "--------" ) s[0] = 12 for _, v := range s { fmt.Println(v) } } 复制代码`

初始化slice后，他的长度为2，也就是元素个数为2，因为我们没有给任何一个元素赋值，所以为int的zero值，也就是0.可以用len和cap看一下这个slice的长度和容量。

用append给这个slice添加新值，返回一个新的slice，如果容量不够时，go会自动增加容易量，小于一1000个长度时成倍的增长，大于1000个长度时会以1.25或者25%的位数增长。

上面的代码执行完后，slice的结构如下

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbc7d56f41?imageView2/0/w/1280/h/960/ignore-error/1)

### Slice 的声明和初始化 ###

创建和初始化一个slice有几种不同的方式，下面我会一一介绍

#### 使用make声明一个slice ####

` slice1 := make([]int, 3) fmt.Println( "len: " , len(slice1), "cap: " , cap (slice1), "array :" , slice1) slice1 = append(slice1, 1) fmt.Println( "len: " , len(slice1), "cap: " , cap (slice1), "array :" , slice1) 复制代码`

make([]int, 3) 声明了个长度为3的slice，容量也是3。下面的append方法会添加一个新元素到slice里，长度和容量都会发生变化。

输出结果：

` len: 3 cap : 3 array : [0 0 0] len: 4 cap : 6 array : [0 0 0 1] 复制代码`

也可以通过重载方法指定slice的容量，下面：

` slice2 := make([]int, 3, 7) fmt.Println( "len: " , len(slice2), "cap: " , cap (slice2), "array :" , slice2) 复制代码`

输出长度为3，容量为7

### 使用slice字变量 ###

使用字变量来创建Slice,有点像创建一个Array,但是不需要在[]指定长度,这也是Slice和Array的区别。Slice根据初始化的数据来计算度和容量

创建一个长度和容量为5的Slice

` slice3 := []int{1, 2, 3, 4, 5} fmt.Println( "len: " , len(slice3), "cap: " , cap (slice3), "array :" , slice3) 复制代码`

也可以通过索引来指定Slice的长度和容量，

下面创建了一个长度和容量为6的slice

` slice4 := []int{5: 0} fmt.Println( "len: " , len(slice4), "cap: " , cap (slice4), "array :" , slice4) 复制代码`

声明nil Slice，内部结构的指针为nil。可以用append给slice填加新的元素，内部的指针指向一个新的数组

` var slice5 []int fmt.Println( "len: " , len(slice5), "cap: " , cap (slice5), "array :" , slice5) slice5 = append(slice5, 4) fmt.Println( "len: " , len(slice5), "cap: " , cap (slice5), "array :" , slice5) 复制代码`

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbc8634ae0?imageView2/0/w/1280/h/960/ignore-error/1)

创建空Slice有两种方式

` slice6 := []int{} fmt.Println( "len: " , len(slice6), "cap: " , cap (slice6), "array :" , slice6) slice6 = append(slice6, 2) fmt.Println( "len: " , len(slice6), "cap: " , cap (slice6), "array :" , slice6) slice7 := make([]int, 0) fmt.Println( "len: " , len(slice7), "cap: " , cap (slice7), "array :" , slice7) slice7 = append(slice7, 7) fmt.Println( "len: " , len(slice7), "cap: " , cap (slice7), "array :" , slice7) 复制代码`

slice的切片

` // 创建一个容量和长度均为6的slice slice1 := []int{5, 23, 10, 2, 61, 33} // 对slices1进行切片，长度为2容量为4 slice2 := slice1[1:3] fmt.Println( "cap" , cap (slice2)) fmt.Println( "slice2" , slice2) 复制代码`

slice1的底层是一个容量为6的数组，slice2指底层指向slice1的底层数组，但起始位置为array的第一个元素也就是23.因为slices2从索引1开始的，所以无法访问底层数组索引1之前的元素，也无法访问容量之后的元素。可以看下图理解一下。

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbc7f41dae?imageView2/0/w/1280/h/960/ignore-error/1)

新创建的切片长度和容量的计算

对于一个新slice[x:y] 底层数组容量为z，

x: 新切片开始的元素的索引位置，上面的slice1[1:3]中的1就是起始索引

y:新切片希望包含的元素个数，上面的slice1[1:3]，希望包含2个底层数组的元素 1+2=3

容量： z-x 上面的slice1[1:3] 底层数组的容量为6， 6-1=5所以新切片的容量为5

### 修改切片导致的后果 ###

由于新创建的slice2和slice1底层是同一个数组，所以修改任何一个，两个slice共同的指向元素，会导致同时修改的问题

` // 创建一个容量和长度均为6的slice slice1 := []int{5, 23, 10, 2, 61, 33} // 对slices1进行切片，长度为2容量为4 slice2 := slice1[1:3] fmt.Println( "cap" , cap (slice2)) fmt.Println( "slice2" , slice2) //修改一个共同指向的元素 //两个slice的值都会修改 slice2[0] = 11111 fmt.Println( "slice1" , slice1) fmt.Println( "slice2" , slice2) 复制代码`

如图所示

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbed4305c6?imageView2/0/w/1280/h/960/ignore-error/1)

需注意的是，slice只能访问其长度范围内的元素，如果超出长度会报错。

除了修改共同指向元素外，如果新创建的切片长度小于容量，新增元素也会导致原来元素的变动。slice增加新元素使用内置的方法append。append方法会创建一个新切片。

` // 创建一个容量和长度均为6的slice slice1 := []int{5, 23, 10, 2, 61, 33} // 对slices1进行切片，长度为2容量为4 slice2 := slice1[1:3] fmt.Println( "cap" , cap (slice2)) fmt.Println( "slice2" , slice2) //修改一个共同指向的元素 //两个slice的值都会修改 slice2[0] = 11111 fmt.Println( "slice1" , slice1) fmt.Println( "slice2" , slice2) // 增加一个元素 slice2 = append(slice2, 55555) fmt.Println(slice1) fmt.Println(slice2) 复制代码`

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbed7e9c4c?imageView2/0/w/1280/h/960/ignore-error/1)

如果切片的容量足够就把新元素合添加到切片的长度。如果底层的的数组容量不够时，会重新创建一个新的数组并把现有元素复制过去。

` slice3 := []int{1, 2, 3} fmt.Println( "slice2 cap" , cap (slice3)) slice3 = append(slice3, 5) fmt.Println( "slice2 cap" , cap (slice3)) 复制代码`

输出的结果为：

` slice2 cap 3 slice2 cap 6 复制代码`

容量增长了一倍。

### 控制新创建slice的容量 ###

创建一个新的slice的时候可以限制他的容量

` // 创建一个容量和长度均为6的slice slice1 := []int{5, 23, 10, 2, 61, 33} // 对slices1进行切片，长度为2容量为3 slice2 := slice1[1:3:4] fmt.Println( "cap" , cap (slice2)) fmt.Println( "slice2" , slice2) 复制代码`

slice2长度为2容量为3，这也是通过上面的公式算出来的

长度：y-x 3-1=2

容量：z-x 4-1= 3

需要注意的是容量的长度不能大于底层数组的容量

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbef7b3276?imageView2/0/w/1280/h/960/ignore-error/1)

绿颜色表示slice2中的元素，黄颜色表示容量中示使用的元素。但是需要注意的是，我们修改或者增加slice2容量范围内的元素个数依然会修改slice1。

` // 创建一个容量和长度均为6的slice slice1 := []int{5, 23, 10, 2, 61, 33} // 对slices1进行切片，长度为2容量为3 slice2 := slice1[1:3:4] fmt.Println( "cap" , cap (slice2)) fmt.Println( "slice2" , slice2) //修改一个共同指向的元素 //两个slice的值都会修改 slice2[0] = 11111 fmt.Println( "slice1" , slice1) fmt.Println( "slice2" , slice2) // 增加一个元素 slice2 = append(slice2, 55555) fmt.Println(slice1) fmt.Println(slice2) 复制代码`

![](https://user-gold-cdn.xitu.io/2019/3/14/1697a0dbf36ced6e?imageView2/0/w/1280/h/960/ignore-error/1)