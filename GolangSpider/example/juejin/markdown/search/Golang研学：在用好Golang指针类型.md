# Golang研学：在用好Golang指针类型 #

![Golang指针](https://user-gold-cdn.xitu.io/2019/4/26/16a59343554feb15?imageView2/0/w/1280/h/960/ignore-error/1)

在大部分面向对象语言如C++、C#、Java，在函数传参数时除了基础值类型，对象是通过引用方式传递的。

**然而，在Go语言中，除了map、slice和chan，所有类型（包括struct）都是值传递的。**

那么，如何在 **函数外** 使用 **函数内处理后** 的变量呢？只能通过返回新变量吗？

**不，可以使用指针**

大部分面向对象语言都很少有用到指针的场景了，但是在Go语言中有大量的指针应用场景，要成为一名合格的Gopher，必须了解。

## 概念 ##

每一个变量都会分配一块内存，数据保存在内存中，内存有一个地址，就像门牌号，通过这个地址就可以找到里面存储的数据。

指针就是保存这个内存地址的变量。

**在Go语言中，用 ` &` 取得变量的地址**

` //为了说明类型，我采用了显性的变量定义方法，实际开发中更多的是用“:=”自动获取类型变量类型 var mystr string = "Hello!" var mystrP * string = &mystr fmt.Println(mystrP) 复制代码`

将以上代码敲入main函数中， ` go run` ，打印出的内容就是 ` mystr` 的内存地址。 ` mystrP` 就是 ` mystr` 的指针变量。

**用 ` *` 取得指针变量指向的内存地址的值**

在之前的代码的后面增加一句代码：

` fmt.Println(*mystrPointer) 复制代码`

` go run` 运行后，可以看到打印出 ` mystr` 的值“Hello！”

**符号 ` *` 也用做定义指针类型的关键字。**

例如：

` var p * int 复制代码`

## 指针应用场景 ##

在其他OOP语言中，大多数情况是不需要花太多时间操作指针的，如Java、C#，对象的引用操作都已经交给了虚拟机和框架。而Go经常会用到指针。原因主要有3个：

* Go语言中除了map、slice、chan外，其他类型在函数参数中都是值传递
* Go语言不是面向对象的语言，很多时候实现结构体方法时需要用指针类型实现引用结构体对象
* 指针也是一个类型，在实现接口 ` interface` 时，结构体类型和其指针类型对接口的实现是不同的

接下来就分别介绍一下，期间会穿插一些简单的代码片段，您可以创建一个Go文件输入代码，运行体验一下。

### 函数中传递指针参数 ###

Go语言都是值传递，例如：

` package main import "fmt" func main () { i := 0 f(i) fmt.Println(i) } func f (count int ) { fmt.Println(count) count++ } 复制代码`

结果：

` 0 0 复制代码`

` i` 在执行前后没有变化

如果希望被函数调用后， ` i` 的值产生变化， ` f` 函数的参数就应该改为 ` *int` 类型。如：

` func main () { i := 0 f(&i) fmt.Println(i) } func f (count * int ) { fmt.Println(*count) (*count)++ } 复制代码` * f定义参数用 ` *int` 替代 ` int` ，申明参数是一个int类型的指针类型
* 调用函数时，不能直接传递int的变量 ` i` ，而要传递用 ` &` 取得 ` i` 的地址
* f函数内，参数 ` count` 现在是指针了，不能直接打印，需要用 ` *` 取出这个指针指向的地址里保存的值
* ` count` 的取值+1.
* 调用f函数，在主函数 ` main` 里打印 ` i` 。

可以看到结果

` 0 1 复制代码`

` i` 的值改变了。

### Struct结构体指针类型方法 ###

Go语言中给结构体定义方法

` //定义一个结构体类型 type myStruct struct { Name string } //定义这个结构体的改名方法 func (m myStruct) ChangeName (newName string ) { m.Name = newName } func main () { //创建这个结构体变量 mystruct := myStruct{ Name: "zeta" , } //调用改名函数 mystruct.ChangeName( "Chow" ) //没改成功 fmt.Println(mystruct.Name) } 复制代码`

这样的方法不会改掉结构体变量内的字段值。 就算是结构体方法，如果不使用指针，方法内还是传递结构体的值。

现在我们改用指针类型定义结构体方法看看。

只修改 ` ChangeName` 函数，用 ` *myStruct` 类型替代 ` myStruct`

` func (m *myStruct) ChangeName (newName string ) { m.Name = newName } 复制代码`

再运行一次，可以看到打印出来的名字改变了。

**当使用指针类型定义方法后，结构体类型的变量调用方法时会自动取得该结构体的指针类型并传入方法。**

### 指针类型的接口实现 ###

最近在某问答平台回答了一个Gopher的问题，大致内容是问为什么不能用结构体指针类型实现接口方法？

看一下代码

` //定义一个接口 type myInterface interface { ChangeName( string ) SayMyName() } //定义一个结构体类型 type myStruct struct { Name string } //定义这个结构体的改名方法 func (m *myStruct) ChangeName (newName string ) { m.Name = newName } func (m myStruct) SayMyName () { fmt.Println(m.Name) } //一个使用接口作为参数的函数 func SetName (s myInterface, name string ) { s.ChangeName(name) } func main () { //创建这个结构体变量 mystruct := myStruct{ Name: "zeta" , } SetName(mystruct, "Chow" ) mystruct.SayMyName() } 复制代码`

这段代码是无法编译通过的，会提示

` cannot use mystruct (type myStruct) as type myInterface in argument to SetName: myStruct does not implement myInterface (ChangeName method has pointer receiver) 复制代码`

` myStruct` 类型没有实现接口方法 ` ChangeName` ，也就是说 ` func (m *myStruct) ChangeName(newName string)` 并不算实现了接口，因为它是 ` *myStruct` 类型实现的，而不是 ` myStruct` 。

改一改

在调用SetName时，用&mystruct 替代 mystruct:

` SetName(&mystruct, "Chow" ) 复制代码`

编译运行，成功。

为什么结构体类型实现的接口该结构体的指针类型也算实现了，而指针类型实现的接口，不算是该结构体实现了接口呢？

** 原因是，结构体类型定义的方法可以被该结构体的指针类型调用；而结构体类型调用该指针类型的方法时是被转换成指针，不是直接调用。**

所以， ` &mystruct` 直接实现了接口定义的 ` ChangeName` 和 ` SayMyName` 两个方法，而 ` mystruct` 只能实现了 ` SayMyName` ， ` mystruct` 调用 ` ChangeName` 方法其实转换成指针类型后调用的，不算实现了接口。

到此Go语言指针类型的应用介绍差不多了。

## 总结一下： ##

* Go语言中指针非常常用，一定要掌握
* Go语言除了map、slice、chan其他都是值传递，引用传递一定要用指针类型
* 结构体类型定义方法要注意使用指针类型
* 接口实现方法时，用指针类型实现的接口函数只能算是指针类型实现的，用结构体类型实现的方法也作为是指针类型实现。

欢迎大家一起讨论、学习Go语言！！

欢迎关注公众号，和大家一起学习编程吧

![晓代码公众号](https://user-gold-cdn.xitu.io/2019/5/7/16a908cc6e393371?imageView2/0/w/1280/h/960/ignore-error/1)