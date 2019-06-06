# [译] part 21: golang goroutines #

> 
> 
> 
> * 原文地址： [Part 21: Goroutines](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Fgoroutines%2F
> )
> * 原文作者： [Naveen R](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Fabout%2F )
> * 译者：咔叽咔叽 转载请注明出处。
> 
> 
> 

在前面的教程中，我们讨论了并发以及它与并行的不同之处。在本教程中，我们将讨论如何使用 ` Goroutines` 在 Go 中实现并发性。

## 什么是 ` Goroutines` ##

` Goroutines` 是与其他函数或方法同时运行的函数或方法。 ` Goroutines` 可以被认为是轻量级线程。与线程相比，创建 ` Goroutine` 的成本很小。因此，Go 应用程序通常可以轻松运行数千个 ` Goroutines` 。

## ` Goroutines` 相较线程的优势 ##

* 与线程相比， ` Goroutines` 非常轻量化。它们的堆栈大小只有几 kb，堆栈可以根据应用程序的需要而伸缩，而在线程的情况下，堆栈必须固定指定大小。
* ` Goroutines` 被复用到较少数量的系统线程。程序中可能一个线程有数千个 ` Goroutines` 。如果该线程中的任何 ` Goroutine` 阻塞，则创建另一个系统线程，并将剩余的 ` Goroutines` 移动到新的线程。所有这些都由运行时处理，Go 从这些复杂的细节中抽象出来一个简洁的 API 来原生支持并发。
* ` Goroutines` 使用 ` channel` 进行通信。 ` Goroutines` 使用 ` channel` 通信可以避免因访问共享内存而发生竞态条件。 ` channel` 可以被认为是 ` Goroutines` 通信的管道。我们将在下一个教程中详细讨论 ` channel` 。

## 如何启动 ` Goroutines` ##

使用关键字 ` go` 对函数或方法调用进行前缀修饰，就可以运行新的 ` Goroutine` 了。

来创建一个 ` Goroutine` 吧:)

` package main import ( "fmt" ) func hello () { fmt.Println( "Hello world goroutine" ) } func main () { go hello() fmt.Println( "main function" ) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2FKUvIpVzepjs )

在第 11 行， ` go hello()` 启动了一个新的 ` Goroutine` 。现在 ` hello` 函数将与 ` main` 函数一起并发运行。 ` main` 在其自己的 ` Goroutine` 中运行，并把 ` main` 函数执行的 ` Goroutine` 为 ` main Goroutine` 主协程。

运行这个程序，你会有一个惊喜！

该程序仅输出了 ` main` 函数的文本。我们启动的 ` Goroutine` 怎么了？我们需要了解 ` go` 协程的两个主要属性，就知道为什么会发生这种情况了。

* 当一个新的 ` Goroutine` 启动时， ` Goroutine` 调用立即返回。与函数不同，控制器不会等待 ` Goroutine` 完成执行。在 ` Goroutine` 调用之后，控制器立即返回到下一行代码，并忽略了 ` Goroutine` 的任何返回值。
* ` main Goroutine` 控制该进程的任何其他 ` Goroutines` 运行。如果 ` main Goroutine` 终止，那么程序将被终止，其他 ` Goroutine` 将可能得不到运行。

我猜你能够理解为什么我们的 ` Goroutine` 没有被执行。在第 11 行调用 ` go hello()` ，控制器立即执行下一行代码并不等待 ` hello goroutine` 执行，然后打印了 ` main function` 之后 ` main Goroutine` 终止。没有等待时间，因此你的 ` Goroutine` 没有时间执行。

我们简单解决一下这个问题。

` package main import ( "fmt" "time" ) func hello () { fmt.Println( "Hello world goroutine" ) } func main () { go hello() time.Sleep( 1 * time.Second) fmt.Println( "main function" ) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2F4nT_Q6CuGrp )

在上面程序的第 13 行，我们调用了 ` time` 包的 ` Sleep` 方法，该方法让正在执行它的 ` go` 协程休眠。在这种情况下， ` main Goroutine` 进入休眠状态 1 秒钟。现在调用 ` go hello()` 有足够的时间在 ` main Goroutine` 终止之前执行。该程序首先打印 ` Hello world goroutine` ，然后等待 1 秒然后打印 ` main function` 。

这种在 ` mian Goroutine` 中使用 ` sleep` 等待其他 ` Goroutines` 完成执行的方式只是我们用来理解 ` Goroutines` 如何工作的，正常情况下肯定不能这么做。 ` channels` 可用于阻塞 ` main Goroutine` ，直到所有其他 ` Goroutines` 完成执行。我们将在下一个教程中讨论 ` channel` 。

## 启动多个 ` Goroutines` ##

再写一个程序，启动多个 ` Goroutines` 以更好地理解它。

` package main import ( "fmt" "time" ) func numbers () { for i := 1 ; i <= 5 ; i++ { time.Sleep( 250 * time.Millisecond) fmt.Printf( "%d " , i) } } func alphabets () { for i := 'a' ; i <= 'e' ; i++ { time.Sleep( 400 * time.Millisecond) fmt.Printf( "%c " , i) } } func main () { go numbers() go alphabets() time.Sleep( 3000 * time.Millisecond) fmt.Println( "main terminated" ) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2Foltn5nw0w3 )

上面的程序在 21 和 22 行分别启动了两个 ` Goroutines` ，这两个 ` Goroutines` 同时运行。 ` numbers Goroutine` 最初睡眠 250 毫秒然后打印 1，然后再次睡眠并打印 2，并且循环直到它打印 5。类似地， ` alphabets Goroutine` 从 a 到 e 打印字母并且是 400 毫秒的睡眠时间。 ` main Goroutine` 在启动 ` alphabets` 和 ` numbers Goroutines` 后睡眠 3000 毫秒，然后终止。

输出，

` 1 a 2 3 b 4 c 5 d e main terminated 复制代码`

下图描绘了该程序的工作原理。请在新标签页中打开图片以获得更好的可视性 :)

![](https://user-gold-cdn.xitu.io/2019/4/2/169d9a94421ae4bf?imageView2/0/w/1280/h/960/ignore-error/1)

蓝色图像的第一部分代表 ` numbers Goroutine` ，栗色的第二部分代表 ` alphabets Goroutine` ，绿色的第三部分代表 ` main Goroutine` ，黑色的合并了上述三个并向我们展示如何程序如何执行的。每个框顶部的 0 ms，250 ms 等字符串表示以毫秒为单位的时间，输出在每个框的底部，例如 1, 2, 3 等等。蓝色框告诉我们在 250 毫秒时打印 1，在 500 毫秒时打印 2，依此类推。黑色框底部的值为 1 a 2 3 b 4 c 5 d e 然后 ` main` 终止，这也是程序的输出。希望能通过这个图理解该程序的工作原理。