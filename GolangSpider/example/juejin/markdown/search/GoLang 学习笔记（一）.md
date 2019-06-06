# GoLang 学习笔记（一） #

## 包 ##

main 包中的main函数是程序的入口；

#### 包的两种导入方式 ####

1 逐个导入

` import "fmt" import "math" 复制代码`

2 分组导入

` import ( "fmt" "math" ) 复制代码`

官方建议使用分组导入的方式

#### 导出名 ####

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的；在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

例如math 包中的Pi是导出的，pi是未导出的，执行程序时math.pi会报错，而math.Pi可以正常运行

## 函数 ##

函数可以没有参数或接受多个参数；

` func add(x int, y int) int { return x + y } func main() { fmt.Println(add(42, 13)) } 复制代码`

注意：当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

` func add(x, y int) int { return x + y } 复制代码`

上面的函数中

` x int, y int 复制代码`

被缩写为

` x, y int 复制代码`

#### 多值返回 ####

函数可以返回任意数量的返回值。

` func swap(x, y string) (string, string) { return y, x } 复制代码`

## 变量 ##

var 语句用于声明一个变量或一个变量列表。

` var c bool var python, java bool 复制代码`

#### 短变量 ####

只在函数内有效的变量类似于Java中的局部变量，使用:= 符号声明。

` func main() { k := 3 } 复制代码`

注意：:= 结构不能在函数外使用。

## 基本类型 ##

bool 布尔

string 字符串

int int8 int16 int32 int64 整形

uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名，表示一个 Unicode 码点

float32 float64 浮点

complex64 complex128

注意：int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 **当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型**

#### 类型转换 ####

可以使表达式 T(v) 将值 v 转换为类型 T。例如：

` var i int = 42 var f float64 = float64(i) 复制代码`

注意：Go 在不同类型的项之间赋值时必须要显式转换。

## 常量 ##

常量的声明与变量类似，使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 := 语法声明。

` const Pi = 3.14 复制代码`

## 循环结构 ##

**Go 只有一种循环结构：for 循环**

` for i := 0; i < 10; i++ { sum += i } 复制代码`

i := 0 初始化语句 i < 10 条件表达式 i++ 后置语句

和Java非常像了。

初始化语句和后置语句也是可选的

` for ; sum < 1000; { sum += sum } 复制代码`

#### 无限循环 ####

` for { } 复制代码`

## if ##

if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。

` if x < 0 { return sqrt(-x) + "i" } 复制代码`

if 语句可以在条件表达式前执行一个简单的语句，该语句声明的变量作用域仅在 if 之内。

` if v := math.Pow(x, n); v < lim { return v } 复制代码`

## switch ##

` switch os := runtime.GOOS; os { case "darwin": fmt.Println("OS X.") case "linux": fmt.Println("Linux.") default: fmt.Printf("%s.\n", os) } 复制代码`

case 无需为常量，且取值不必为整数。

#### 没有条件的 switch ####

没有条件的 switch 同 switch true 一样,通常用来代替较多的if-then-else。

` func main() { t := time.Now() switch { case t.Hour() < 12: fmt.Println("Good morning!") case t.Hour() < 17: fmt.Println("Good afternoon.") default: fmt.Println("Good evening.") } } 复制代码`

## defer ##

defer 语句会将函数推迟到外层函数返回之后执行。例如：

` func main() { defer fmt.Println("world") fmt.Println("hello") } 复制代码`

结果为：

` hello world 复制代码`

原理：推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

## 指针 ##

Go 拥有指针。指针保存了值的内存地址。类型 *T 是指向 T 类型值的指针。其零值为 nil。

var p *int

& 操作符会生成一个指向其操作数的指针。

i := 42

p = &i

星号操作符表示指针指向的底层值。

fmt.Println(*p) // 通过指针 p 读取 i 结果为42

*p = 21 // 通过指针 p 设置 i

fmt.Println(*p) //结果为21

## 结构体 ##

类似于Java中的实体，一个结构体（struct）就是一组字段（field）。

` type Vertex struct { X int Y int } func main() { fmt.Println(Vertex{1, 2}) } 复制代码`

注意：结构体中的字段也可以通过指针来访问。

## 数组 ##

表达式：[n]T 表示拥有 n 个 T 类型的值的数组。

` var a [10]int 复制代码`

注意：数组的长度是其类型的一部分，因此数组不能改变大小。

## 切片 ##

切片为数组元素提供动态大小的、灵活的视角，在实践中，切片比数组更常用。

类型 []T 表示一个元素类型为 T 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

a[low : high]

它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

a[1:4]

**切片并不存储任何数据，它只是描述了底层数组中的一段。**

**更改切片的元素会修改其底层数组中对应的元素。**

**与它共享底层数组的切片都会观测到这些修改。**

对于数组

var a [10]int 来说，以下切片是等价的：

` a[0:10] a[:10] a[0:] a[:] 复制代码`

切片的长度就是它所包含的元素个数，可以通过len(s)获取 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数，可以通过cap(s)获取。

#### nil 切片 ####

切片的零值是 nil。

nil 切片的长度和容量为 0 且没有底层数组。

切片可以用内建函数 make 来创建，这也是创建动态数组的方式。

切片中可以包含其他切片。

## 函数的闭包？ ##

## 方法 ##

Go 没有类。不过你可以为结构体类型定义方法。

方法就是一类带特殊的 接收者 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。

` type Vertex struct { X, Y float64 } func (v Vertex) Abs() float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) } func main() { v := Vertex{3, 4} fmt.Println(v.Abs()) } 复制代码`

注意：方法只是个带接收者参数的函数。

非结构体也可定义方法，例如：

` type MyFloat float64 func (f MyFloat) Abs() float64 { if f < 0 { return float64(-f) } return float64(f) } 复制代码`

## 接口类型 ##

接口类型 是由一组方法签名定义的集合,接口类型的变量可以保存任何实现了这些方法的值。

## Go 程 ##

Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

go f(x, y, z)会启动一个新的 Go 程并执行

f(x, y, z)f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。

## 信道 ##

信道是带有类型的管道，适合在各个 Go 程间进行通信。你可以通过它用信道操作符 <- 来发送或者接收值。

` ch <- v // 将 v 发送至信道 ch。 v := <-ch // 从 ch 接收值并赋予 v。（“箭头”就是数据流的方向。） 复制代码`

和映射与切片一样，信道在使用前必须创建：

` ch := make(chan int) 复制代码`

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

#### 带缓冲的信道 ####

信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

` ch := make(chan int, 100) 复制代码`

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。