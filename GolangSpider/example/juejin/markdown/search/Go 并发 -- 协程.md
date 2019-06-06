# Go 并发 -- 协程 #

> 
> 
> 
> **这是『就要学习 Go 语言』系列的第 21 篇分享文章**
> 
> 

### 并发与并行 ###

提到并发，相信大家还听过另一个概念 -- 并行。我先给大家介绍下这两者之间的区别，再来讲 Go 语言的并发。

**并行** 其实很好理解，就是同时执行的意思，在某一时间点能够执行多个任务。 想达到并行效果，最简单的方式就是借助多线程或多进程，这样才可在同一时刻执行多个任务。单线程是永远无法达到并行状态的。 **并发** 是在某一时间段内可以同时处理多个任务。我们通常会说程序是并发设计的，也就是说它允许多个任务同时执行，这个同时指的就是一段时间内。单线程中多个任务以间隔执行实现并发。 可以说， **多线程或多进程是并行的基础，但单线程也通过协程实现了并发。**

举个常见的例子，一台单核电脑可以下载、听音乐，实际上这两个任务是这样执行的，只不过这两个任务切换时间短，给人的感觉是同时执行的。

![并发](https://user-gold-cdn.xitu.io/2019/4/19/16a32df8ea4ac162?imageView2/0/w/1280/h/960/ignore-error/1) 一台多核电脑的任务执行就是像下面这种图显示的一样： ![并行](https://user-gold-cdn.xitu.io/2019/4/19/16a32df8eb43e47a?imageView2/0/w/1280/h/960/ignore-error/1) 可以看到，同一时刻能执行多个任务。这种任务执行方式才是真正的并行。

Go 通过协程实现并发，协程之间靠信道通信，本篇文章先给大家介绍协程的使用，后面再写信道。

### 协程 ###

协程（Goroutine）可以理解成轻量级的线程，但与线程相比，它的开销非常小。因此，Go 应用程序通常能并发地运行成千上万的协程。 Go 创建一个协程非常简单，只要在方法或函数调用之前加关键字 go 即可。

` func printHello () { fmt.Println( "hello world goroutine" ) } func main () { go printHello() // 创建了协程 fmt.Println( "main goroutine" ) } 复制代码`

输出：

` main goroutine 复制代码`

上面代码，第 6 行使用 go 关键字创建了协程，现在有两个协程，新创建的协程和主协程。printHello() 函数将会独立于主协程并发地执行。 是的，你没有看错，程序的输出就是这样，不信？你可以实际运行下程序。你惊讶的是，printHello() 函数为什么没有输出，到底发生了什么？ **当协程创建完毕之后，主函数立即返回继续执行下一行代码，不像函数调用，需要等函数执行完成** 。主协程执行完毕，程序便退出，printHello 协程随即也退出，便不会有输出。

修改下代码：

` func printHello () { fmt.Println( "hello world goroutine" ) } func main () { go printHello() time.Sleep( 1 *time.Second) fmt.Println( "main goroutine" ) } 复制代码`

协程创建完成之后，main 协程先休眠 1s，预留给 printHello 协程执行的时间，所以这次输出：

` hello world goroutine main goroutine 复制代码`

### 创建多个协程 ###

上一节就提到，可以创建多个协程，来看下例子：

` func printNum () { for i := 1 ; i <= 5 ; i++ { time.Sleep( 20 * time.Millisecond) fmt.Printf( "%d " , i) } } func printChacter () { for i := 'a' ; i <= 'e' ; i++ { time.Sleep( 40 * time.Millisecond) fmt.Printf( "%c " , i) } } func main () { go printNum() go printChacter() time.Sleep( 3 *time.Second) fmt.Println( "main terminated" ) } 复制代码`

上面的代码，除主协程之外，新创建了两个协程：printNum 协程和 printChacter 协程。printNum 协程每隔 20 毫秒输出 5 个数，printChacter 协程每隔 40 毫秒输出字母。主协程创建完这两个协程之后休眠 1s，等待其他协程执行完成。 程序输出（你的输出可能跟我的不一样）：

` 1 a 2 3 b 4 5 c d e main terminated 复制代码`

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/4/19/16a32df8ebd91e58?imageView2/0/w/1280/h/960/ignore-error/1)

话说回来，通过加 time.Sleep() 函数等待协程执行完成是一种“黑科技”。在实际生产环境中，不管我们是否知道其他协程执行完成需要多少时间，都不能在主协程中添加随机睡眠调用等待其他协程执行完成。那怎么办？Go 给我们提供了信道，当协程执行完毕，能够通知到主协程，还能够实现协程间通信。下节课我们来讨论一下。

### 匿名协程 ###

在函数那篇文章讲过，存在匿名函数，通过关键字 go 调用匿名函数就是匿名协程。我们修改之前的例子：

` func main () { go func () { fmt.Println( "hello world goroutine" ) }() time.Sleep( 1 *time.Second) fmt.Println( "main goroutine" ) } 复制代码`

输出结果跟之前的一样。

希望这篇文章给你带来收获，Good Day !

（全文完）
> 
> 
> 
> 原创文章，若需转载请注明出处！
> 欢迎扫码关注公众号「 **Golang来啦** 」或者移步 [seekload.net](
> https://link.juejin.im?target=https%3A%2F%2Fseekload.net ) ，查看更多精彩文章。
> 
> 

**给你准备了学习 Go 语言相关书籍，公号后台回复【电子书】领取！**

![公众号二维码](https://user-gold-cdn.xitu.io/2019/3/27/169be4a300f56486?imageView2/0/w/1280/h/960/ignore-error/1)