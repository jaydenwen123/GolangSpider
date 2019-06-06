# [译] Part 32: golang 中的 panic 和 recover #

> 
> 
> 
> * 原文地址： [Part 32: Panic and Recover](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Fpanic-and-recover%2F
> )
> * 原文作者： [Naveen R](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Fabout%2F )
> * 译者：咔叽咔叽 转载请注明出处。
> 
> 
> 

## 什么是panic? ##

处理Go中异常情况的惯用方法是使用errors，对于程序中出现的大多数异常情况，errors就足够了。

但是在某些情况下程序不能在异常情况下继续正常执行。在这种情况下，我们使用panic来终止程序。函数遇到panic时将会停止执行，如果有defer的话就执行defer延迟函数，然后返回其调用者。此过程一直持续到当前goroutine的所有函数都返回，然后打印出panic信息，然后是堆栈信息，然后程序终止。待会儿用一个例子来解释，这个概念就会更加清晰一些了。

我们可以使用recover函数恢复被panic终止的程序，将在本教程后面讨论。

panic和recover有点类似于其他语言中的try-catch-finally语句，但是前者使用的比较少，而且使用时更优雅代码也更简洁。

## 什么时候应该用panic? ##

一般情况下我们应该避免使用panic和recover，尽可能使用errors。只有在程序无法继续执行的情况下才应该使用panic和recover。

##### 两个panic典型应用场景 #####

* 

不可恢复的错误，让程序不能继续进行。 比如说Web服务器无法绑定到指定端口。在这种情况下，panic是合理的，因为如果端口绑定失败接下来的逻辑继续也是没有意义的。

* 

coder的人为错误 假设我们有一个接受指针作为参数的方法，然而使用了nil作为参数调用此方法。在这种情况下，我们可以用panic，因为该方法需要一个有效的指针。

## panic示例 ##

panic函数的定义

` func panic ( interface {}) 复制代码`

当程序终止时，参数会传递给panic函数打印出来。看看下面例子的panic是如何使用的。

` package main import ( "fmt" ) func fullName (firstName * string , lastName * string ) { if firstName == nil { panic ( "runtime error: first name cannot be nil" ) } if lastName == nil { panic ( "runtime error: last name cannot be nil" ) } fmt.Printf( "%s %s\n" , *firstName, *lastName) fmt.Println( "returned normally from fullName" ) } func main () { firstName := "Elon" fullName(&firstName, nil ) fmt.Println( "returned normally from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2Fpeylhn6NBMH )

上面这段代码，fullName函数功能是打印一个人的全名。此函数检查firstName和lastName指针是否为nil。如果它为nil，则函数调用panic并显示相应的错误消息。程序终止时将打印此错误消息和错误堆栈信息。

运行此程序将打印以下输出，

` panic: runtime error: last name cannot be nil goroutine 1 [running]: main.fullName(0x1040c128, 0x0) /tmp/sandbox135038844/main.go:12 +0x120 main.main() /tmp/sandbox135038844/main.go:20 +0x80 复制代码`

我们来分析一下这个输出，来了解panic是如何工作以及如何打印堆栈跟踪的。 在第19行，我们将Elon定义给firstName。然后调用fullName函数，其中lastName参数为nil。因此，第11行将触发panic。当触发panic时，程序执行就终止了，然后打印传递给panic的内容，最后打印堆栈跟踪信息。因此14行以后的代码不会被执行。 该程序首先打印传递给panic函数的内容，

` panic: runtime error: last name cannot be nil 复制代码`

然后打印堆栈跟踪信息。 该程序在12行触发panic，因此，

` ain.fullName(0x1040c128, 0x0) /tmp/sandbox135038844/main.go:12 +0x120 复制代码`

将被首先打印。然后将打印堆栈中的下一个内容，

` main.main() /tmp/sandbox135038844/main.go:20 +0x80 复制代码`

现在已经返回到了造成panic的顶层main函数，因此打印结束。

## defer函数 ##

我们回想一下panic的作用。当函数遇到panic时，将会终止panic后面代码的执行，如果函数体包含有defer函数的话会执行完defer函数。然后返回其调用者。此过程一直持续到当前goroutine的所有函数都返回，此时程序打印出panic内容，然后是堆栈跟踪信息，然后终止。

在上面的示例中，我们没有任何defer函数的调用。修改下上面的例子，来看看defer函数的例子吧。

` package main import ( "fmt" ) func fullName (firstName * string , lastName * string ) { defer fmt.Println( "deferred call in fullName" ) if firstName == nil { panic ( "runtime error: first name cannot be nil" ) } if lastName == nil { panic ( "runtime error: last name cannot be nil" ) } fmt.Printf( "%s %s\n" , *firstName, *lastName) fmt.Println( "returned normally from fullName" ) } func main () { defer fmt.Println( "deferred call in main" ) firstName := "Elon" fullName(&firstName, nil ) fmt.Println( "returned normally from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2F0N8J6AvTObI ) 对之前代码所做的唯一更改是在fullName函数和main函数中第一行添加了defer函数调用。 运行的输出，

` deferred call in fullName deferred call in main panic: runtime error: last name cannot be nil goroutine 1 [running]: main.fullName(0x1042bf90, 0x0) /tmp/sandbox060731990/main.go:13 +0x280 main.main() /tmp/sandbox060731990/main.go:22 +0xc0 复制代码`

当发生panic时，首先执行defer函数，然后到下一个defer调用，依此类推，直到达到顶层调用者。

在我们的例子中，defer声明在fullName函数的第一行。首先执行fullName函数。打印

` deferred call in fullName 复制代码`

然后调用返回到main函数的defer，

` deferred call in main 复制代码`

现在调用已返回到顶层函数，然后程序打印panic内容，然后是堆栈跟踪信息，然后终止。

## recover函数 ##

recover是一个内置函数，用于goroutine从panic的中断状况中恢复。 函数定义如下，

` func recover () interface {} 复制代码`

recover只有在defer函数内部调用时才有效。defer函数内通过调用recover可以让panic中断的程序恢复正常执行，调用recover会返回panic的内容。如果在defer函数之外调用recover，它将不会停止panic序列。

修改一下，使用recover来让panic恢复正常执行。

` package main import ( "fmt" ) func recoverName () { if r := recover (); r!= nil { fmt.Println( "recovered from " , r) } } func fullName (firstName * string , lastName * string ) { defer recoverName() if firstName == nil { panic ( "runtime error: first name cannot be nil" ) } if lastName == nil { panic ( "runtime error: last name cannot be nil" ) } fmt.Printf( "%s %s\n" , *firstName, *lastName) fmt.Println( "returned normally from fullName" ) } func main () { defer fmt.Println( "deferred call in main" ) firstName := "Elon" fullName(&firstName, nil ) fmt.Println( "returned normally from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2Fk9oZlqpkVKf ) 第7行调用了recoverName函数。这里打印了recover返回的值， 发现recover返回的是panic的内容。

打印如下，

` recovered from runtime error: last name cannot be nil returned normally from main deferred call in main 复制代码`

程序在19行触发panic，defer函数recoverName通过调用recover来重新控制该goroutine，

` recovered from runtime error: last name cannot be nil 复制代码`

在执行recover之后，panic停止并且返回到调用者，main函数和程序在触发panic之后将继续从第29行执行。然后打印，

` returned normally from main deferred call in main 复制代码`

## Panic, Recover 和 Goroutines ##

recover仅在从同一个goroutine调用时才起作用。从不同的goroutine触发的panic中recover是不可能的。再来一个例子来加深理解。

` package main import ( "fmt" "time" ) func recovery () { if r := recover (); r != nil { fmt.Println( "recovered:" , r) } } func a () { defer recovery() fmt.Println( "Inside A" ) go b() time.Sleep( 1 * time.Second) } func b () { fmt.Println( "Inside B" ) panic ( "oh! B panicked" ) } func main () { a() fmt.Println( "normally returned from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2F-0yB_Ja_UcZ ) 在上面的程序中，函数b在23行触发panic。函数a调用defer函数recovery用于从panic中恢复。函数a的17行用另外一个goroutine执行b函数。Sleep的作用只是为了确保程序在b运行完毕之前不会被终止，当然也可以用sync.WaitGroup来解决。

你认为该段代码的输出是什么？panic会被恢复吗？答案是不可以。panic将无法被恢复。这是因为recover存在于不同的gouroutine中，并且触发panic发生在不同goroutine执行的b函数。因此无法恢复。 运行的输出，

` Inside A Inside B panic: oh! B panicked goroutine 5 [running]: main.b() /tmp/sandbox388039916/main.go:23 +0x80 created by main.a /tmp/sandbox388039916/main.go:17 +0xc0 复制代码`

可以从输出中看到恢复失败了。

如果在同一个goroutine中调用函数b，那么panic就会被恢复。 在第17行把， ` go b()` 换成 ` b()` 那么会输出，

` Inside A Inside B recovered: oh! B panicked normally returned from main 复制代码`

## 运行时的panic ##

panic还可能由运行时的错误引起，例如数组越界访问。这相当于使用由接口类型 [runtime.Error]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fsrc%2Fruntime%2Ferror.go%3Fs%3D267%3A503%23L1 ) 定义的参数调用内置函数panic。 runtime.Error接口的定义如下，

` type Error interface { error // RuntimeError is a no-op function but // serves to distinguish types that are run time // errors from ordinary errors: a type is a // run time error if it has a RuntimeError method. RuntimeError() } 复制代码`

runtime.Error接口满足内置接口类型 [error]( https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Ferror-handling%2F%23errortyperepresentation ) 。

让我们写一个人为的例子来创建运行时panic。

` package main import ( "fmt" ) func a () { n := [] int { 5 , 7 , 4 } fmt.Println(n[ 3 ]) fmt.Println( "normally returned from a" ) } func main () { a() fmt.Println( "normally returned from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2FL8i-JdQz4SR ) 在上面的程序中，第9行我们试图访问n [3]，这是切片中的无效索引。这个会触发panic，输出如下，

` panic: runtime error: index out of range goroutine 1 [running]: main.a() /tmp/sandbox780439659/main.go:9 +0x40 main.main() /tmp/sandbox780439659/main.go:13 +0x20 复制代码`

您可能想知道是否运行中的panic能够被恢复。答案是肯定的。让我们修改上面的程序，让panic恢复过来。

` package main import ( "fmt" ) func r () { if r := recover (); r != nil { fmt.Println( "Recovered" , r) } } func a () { defer r() n := [] int { 5 , 7 , 4 } fmt.Println(n[ 3 ]) fmt.Println( "normally returned from a" ) } func main () { a() fmt.Println( "normally returned from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2Fi-9AA33FiBn ) 执行后输出，

` Recovered runtime error: index out of range normally returned from main 复制代码`

显然可以看到panic被恢复了。

## recover后获取堆栈信息 ##

我们恢复了panic，但是丢失了这次panic的堆栈调用的信息。 有一种方法可以解决这个，就是使用Debug包中的PrintStack函数打印堆栈跟踪信息

` package main import ( "fmt" "runtime/debug" ) func r () { if r := recover (); r != nil { fmt.Println( "Recovered" , r) debug.PrintStack() } } func a () { defer r() n := [] int { 5 , 7 , 4 } fmt.Println(n[ 3 ]) fmt.Println( "normally returned from a" ) } func main () { a() fmt.Println( "normally returned from main" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2FH-GRDqoE1jU ) 在11行调用了debug.PrintStack，可以看到随后输出，

` Recovered runtime error: index out of range goroutine 1 [running]: runtime/debug.Stack(0x1042beb8, 0x2, 0x2, 0x1c) /usr/ local /go/src/runtime/debug/stack.go:24 +0xc0 runtime/debug.PrintStack() /usr/ local /go/src/runtime/debug/stack.go:16 +0x20 main.r() /tmp/sandbox949178097/main.go:11 +0xe0 panic(0xf0a80, 0x17 cd 50) /usr/ local /go/src/runtime/panic.go:491 +0x2c0 main.a() /tmp/sandbox949178097/main.go:18 +0x80 main.main() /tmp/sandbox949178097/main.go:23 +0x20 normally returned from main 复制代码`

从输出中可以知道，首先是panic被恢复然后打印 ` Recovered runtime error: index out of range` ，再然后打印堆栈跟踪信息。最后在panic被恢复后打印 ` normally returned from main`