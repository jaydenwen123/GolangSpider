# Go 语言类型系统详解 #

> 
> 
> 
> **这是『就要学习 Go 语言』系列的第 18 篇分享文章**
> 
> 

### 什么是类型 ###

不同的编程语言之间，类型的概念有所不同，可以用许多不同的方式来表达，但大体上都有一些相同的地方。

* 类型是一组值；
* 相同类型的值之间可以进行哪些操作，例如：int 类型可以执行 + 和 - 等运算，而对于字符类型，可以执行连接、空检查等操作；

> 
> 
> 
> 因此，语言类型系统指定哪些运算符对哪些类型有效。
> 
> 

### Go 语言的类型系统 ###

boolean、numeric 和 string 是 Go 的基础数据类型，也称为 **预声明类型** （pre-declared type），可用来构造其他的类型，例如字面量类型。

**字面量类型** （type literal）：由预声明类型组合而成（没有用 type 关键字定义），例如：[3]int 、chan int、map[string] string、* int 等。

由字面量类型可构成 **复合类型** ，如：array、struct、map、slice、channel、func、interface 等。

#### 命名类型和未命名类型 ####

具有名称的类型：例如 int、int64、float32、string、bool 等预先声明类型。另外，使用 type 关键字声明的任意类型也称为命名类型。

` var i int // named type type myInt int // named type var b bool // named type 复制代码`

未命名类型：上面提到的复合类型，包括 array、struct、pointer、function、interface、slice、Map 和 channel，都是未命名类型。它们没有名称，但是有关于如何组成的字面量描述符。

` [] string // unnamed type map [ string ] string // unnamed type [ 10 ] int // unnamed type 复制代码`

#### 底层类型 ####

每种类型都有底层类型，如果 T 是预声明类型或字面量类型，则底层类型就是 T 本身；否则，T 的底层类型是 T 在定义时引用的类型的底层类型。

` type A string // string type B A // string type M map [ string ] int // map[string]int type N M // map[string]int type P *N // *N type S string // string type T map [S] int // map[S]int type U T // map[S]int 复制代码`

第 1、 6 行，预声明的字符串类型，因此底层类型是 T 本身，即字符串；

第 3 、5 行，是字面量类型，因此底层类型就是 T 本身，即 map[string]int 和 指针 *N。注意：字面量类型也是未命名类型；

第 2 、4、8 行，T 的底层类型是 T 在其定义时引用的类型的底层类型，例如：B 引用了 A，所以 B 的底层类型是字符串类型，其他情况同理；

我们再来看下第 7 行的例子：type T map[S]int ，由于 S 的底层类型是 string，难道此时 T 的底层类型不应该是 map[string]int 而不是 map[S]int 吗？因为我们在谈论 map[S]int 的底层未命名类型，所以向下追溯到未命名类型，正如 Go 语言规范上写的一样：如果 T 是字面量类型，则对应的底层类型就是 T 本身。

#### 可赋值性 ####

关于变量的可赋值性在 Go 语言的文档中已经讲得很清楚了，我们来看其中比较重要的一条：当变量 a 可以赋值给类型 T 的变量时， **两者都应该具有相同的底层类型，并且至少其中一个不是命名类型** 。

看下代码

` package main type aInt int func main () { var i int = 10 var ai aInt = 100 i = ai printAiType(i) } func printAiType (ai aInt) { print (ai) } 复制代码`

上面的代码编译不通过，编译时报错：

` 8 : 4 : cannot use ai ( type aInt) as type int in assignment 9 : 13 : cannot use i ( type int ) as type aInt in argument to printAiType 复制代码`

因为 i 是命名类型 int，而 ai 是命名类型 aInt，虽然它们的底层类型相同，都是 int。

` package main type MyMap map [ int ] int func main () { m := make ( map [ int ] int ) var mMap MyMap mMap = m printMyMapType(mMap) print (m) } func printMyMapType (mMap MyMap) { print (mMap) } 复制代码`

上面这段代码编译通过，因为 m 是未命名类型并且 m 和 mMap 的底层类型相同。

#### 类型转化 ####

看下类型转化的规范

![类型转化规范](https://user-gold-cdn.xitu.io/2019/3/18/169900c595eb6379?imageView2/0/w/1280/h/960/ignore-error/1)

` package main type Meter int64 type Centimeter int32 func main () { var cm Centimeter = 1000 var m Meter m = Meter(cm) print (m) cm = Centimeter(m) print (cm) } 复制代码`

上面的代码可以编译通过，因为 **Meter** 和 **Centimeter** 都是整型，并且它们的底层类型可以相互转化。

#### 类型一致性 ####

两种类型要么相同要么不同。

**已定义类型与其他任意类型总是不同** 。因此，即使预先声明的命名类型 int、int64 等也是不相同的。

来看下结构体的一条转化规则：

> 
> 
> 
> x 赋值给 T 时，不考虑结构体标签，x 和 T 应具有相同的底层类型
> 
> 

` package main type Meter struct { value int64 } type Centimeter struct { value int32 } func main () { cm := Centimeter{ value: 1000 , } var m Meter m = Meter(cm) print (m.value) cm = Centimeter(m) print (cm.value) } 复制代码`

记住一点： **相同的底层类型** 。由于成员 Meter.value 的底层类型是 int64，而成员 Centimeter.value 的底层类型是 int32，所以它们不相同，因为 **已定义类型与其他任意类型总是不同** 。所以上面的代码片段编译会出错。

` package main type Meter struct { value int64 } type Centimeter struct { value int64 } func main () { cm := Centimeter{ value: 1000 , } var m Meter m = Meter(cm) print (m.value) cm = Centimeter(m) print (cm.value) } 复制代码`

成员 Meter.value 和 Centimeter.value 的底层类型都是 int64，所以它们相同，编译可以通过。

ps：我在 GCTT（Go 中国翻译组） 翻译的第一篇文章就是关于 Go 语言的类型系统的，今天整理处理给大家看下，希望这篇文章对你理解 Go 类型系统有所帮助！

（全文完）
> 
> 
> 
> 原创文章，若需转载请注明出处！
> 欢迎扫码关注公众号「 **Golang来啦** 」或者移步 [seekload.net](
> https://link.juejin.im?target=https%3A%2F%2Fseekload.net ) ，查看更多精彩文章。
> 
> 

**公众号「Golang来啦」给你准备了一份神秘学习大礼包，后台回复【电子书】领取！**

![公众号二维码](https://user-gold-cdn.xitu.io/2019/3/27/169be4a300f56486?imageView2/0/w/1280/h/960/ignore-error/1)