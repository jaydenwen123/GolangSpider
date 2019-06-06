# [译] part 15: golang 指针 pointers #

> 
> 
> 
> * 原文地址： [Part 15: Pointers](
> https://link.juejin.im?target=%255Bhttps%3A%2F%2Fgolangbot.com%2Fpointers%2F
> )
> * 原文作者： [Naveen R](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Fabout%2F )
> * 译者：咔叽咔叽 转载请注明出处。
> 
> 
> 

## 什么是指针 ##

指针是存储另一个变量的内存地址的变量。

![](https://user-gold-cdn.xitu.io/2019/4/7/169f5976fedf48df?imageView2/0/w/1280/h/960/ignore-error/1)

在上面的图示中，变量 ` b` 值为 156 并存储在内存地址 0x1040a124 处。变量 ` a` 保存了 ` b` 的地址，那么 ` a` 就是指针并指向 ` b` 。

## 声明指针 ##

` * T` 是指针变量的类型，它指向类型为 ` T` 的值。

看段代码吧，

` package main import ( "fmt" ) func main () { b := 255 var a * int = &b fmt.Printf( "Type of a is %T\n" , a) fmt.Println( "address of b is" , a) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FA4vmlgxAy8 )

` ＆` 运算符用于获取变量的地址。上面程序的第 9 行，我们将 ` b` 的地址分配给其类型为 ` * int` 的 ` a` 。现在可以说 ` a` 指向 ` b` 。当我们打印 ` a``的值时就是` b`的地址。输出，

` Type of a is *int address of b is 0x1040a124 复制代码`

你可能会获得不同的地址，因为 ` b` 可以在内存中的任何位置。

## 指针的零值 ##

指针的零值是 ` nil`

` package main import ( "fmt" ) func main () { a := 25 var b * int if b == nil { fmt.Println( "b is" , b) b = &a fmt.Println( "b after initialization is" , b) } } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FyAeGhzgQE1 )

` b` 在上述程序中最初为 ` nil` ，然后将 ` a` 的地址赋值给 ` b` 。输出，

` b is <nil> b after initialisation is 0x1040a124 复制代码`

## 指针解引用 ##

解引用指针意味着访问指针指向的变量的值。 ` * a` 是解引用的语法。

看看是如何执行的，

` package main import ( "fmt" ) func main () { b := 255 a := &b fmt.Println( "address of b is" , a) fmt.Println( "value of b is" , *a) } 复制代码`

[Run in playground]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2Fm5pNbgFwbM )

在上述程序的第 10 行，我们解引用指针 ` a` 并打印它的值。正如预期的那样，它打印出 ` b` 的值。该程序的输出是

` address of b is 0x1040a124 value of b is 255 复制代码`

再写一个程序，我们用指针改变 b 中的值。

` package main import ( "fmt" ) func main () { b := 255 a := &b fmt.Println( "address of b is" , a) fmt.Println( "value of b is" , *a) *a++ fmt.Println( "new value of b is" , b) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FcdmvlpBNmb )

上面程序的第 12 行，我们将 ` a` 指向的值增加 1，它将改变 ` b` 的值，因为 ` a` 指向 ` b` 。因此， ` b` 的值变为 256。程序的输出是

` address of b is 0x1040a124 value of b is 255 new value of b is 256 复制代码`

## 将指针传递给函数 ##

` package main import ( "fmt" ) func change (val * int ) { *val = 55 } func main () { a := 58 fmt.Println( "value of a before function call is" ,a) b := &a change(b) fmt.Println( "value of a after function call is" , a) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2F3n2nHRJJqn )

在上面程序的第 14 行，我们传递指针变量 ` b` 给 ` change` 函数。在 ` change` 函数内部，使用解引用来修改 ` a` 的值。此程序输出，

` value of a before function call is 58 value of a after function call is 55 复制代码`

## 不要将指向数组的指针作为函数的参数，应该改用切片 ##

我们假设想在函数内部对数组进行一些修改，并且数组所做的修改应该对调用者可见。一种方法是将指向数组的指针作为函数的参数传递。

` package main import ( "fmt" ) func modify (arr *[3] int ) { (*arr)[ 0 ] = 90 } func main () { a := [ 3 ] int { 89 , 90 , 91 } modify(&a) fmt.Println(a) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FlOIznCbcvs )

在上面的程序的第 13 行，我们将数组 ` a` 的地址传递给 ` modify` 函数。在 ` modify` 函数中，我们解除引用 ` arr` 并将 90 分配给数组的第一个元素。该程序输出[90 90 91]

` a [x]` 是 ` (* a)[x]` 的简写。所以上面程序中的 ` (* arr)[0]` 可以用 ` arr [0]` 代替。让我们用这个语法重写上面的程序。

` package main import ( "fmt" ) func modify (arr *[3] int ) { arr[ 0 ] = 90 } func main () { a := [ 3 ] int { 89 , 90 , 91 } modify(&a) fmt.Println(a) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2Fk7YR0EUE1G ) 该程序也会输出 ` [90 90 91]`

虽然这种将指向数组的指针作为函数的参数传递并对其进行修改的方式有效，但这并不是 Go 中的惯用方法。我们有切片 [slice]( https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Farrays-and-slices%2F ) )。

我们使用切片重新上述代码，

` package main import ( "fmt" ) func modify (sls [] int ) { sls[ 0 ] = 90 } func main () { a := [ 3 ] int { 89 , 90 , 91 } modify(a[:]) fmt.Println(a) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FrRvbvuI67W )

在上面程序的第 13 行中，我们将一个切片传递给 ` modify` 函数。切片的第一个元素在 ` modify` 函数内被修改为 90。该程序也输出 ` [90 90 91]` 。所以忘记将指针传递给数组吧，使用切片更干净，是惯用的 Go :)。 译者注： ` But even this style isn't idiomatic Go. Use slices instead.` 这句话是 Go 官方文档推荐的，其实重要的一点是，在 Go 中数组是定长的，所以看到按引用传递的定义 ` func modify(arr *[3]int)` 是这个样子。如果我的数组要扩容还得修改入参，非常不灵活。当然还有其他的弊病，如果没有确定数组能干什么，那就按官方文档的建议来吧。

## Go 不支持指针算数运算 ##

Go 不支持指针算运算，这点和像 C 这样的其他语言不同。

` package main func main () { b := [...] int { 109 , 110 , 111 } p := &b p++ } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FWRaj4pkqRD ) )

上面的程序将抛出编译错误 ` main.go:6: invalid operation: p++ (non-numeric type *[3]int)`

我在 github 中创建了一个 [程序]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgolangbot%2Fpointers%2Fblob%2Fmaster%2Fpointers.go ) ，它涵盖了我们这一节讨论过的所有内容。