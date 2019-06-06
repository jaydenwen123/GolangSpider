# Go 并发 -- Select #

> 
> 
> 
> **这是『就要学习 Go 语言』系列的第 23 篇分享文章**
> 
> 

#### Select 的作用 ####

select 的用法有点类似 switch 语句，但 select 不会有输入值而且只用于信道操作。select 用于从多个发送或接收信道操作中进行选择，语句会阻塞直到其中有信道可以操作，如果有多个信道可以操作，会随机选择其中一个 case 执行。 看下例子：

` func service1 (ch chan string ) { time.Sleep( 2 * time.Second) ch <- "from service1" } func service2 (ch chan string ) { time.Sleep( 1 * time.Second) ch <- "from service2" } func main () { ch1 := make ( chan string ) ch2 := make ( chan string ) go service1(ch1) go service2(ch2) select { // 会发送阻塞 case s1 := <-ch1: fmt.Println(s1) case s2 := <-ch2: fmt.Println(s2) } } 复制代码`

输出：from service2 上面的例子执行到 select 语句的时候回发生阻塞，main 协程等待一个 case 操作可执行，很明显是 service2 先准备好读取的数据（休眠 1s），所以输出 from service2。 看下在两种操作都准备好的情况：

` func service1 (ch chan string ) { //time.Sleep(2 * time.Second) ch <- "from service1" } func service2 (ch chan string ) { //time.Sleep(1 * time.Second) ch <- "from service2" } func main () { ch1 := make ( chan string ) ch2 := make ( chan string ) go service1(ch1) go service2(ch2) time.Sleep( 2 *time.Second) select { case s1 := <-ch1: fmt.Println(s1) case s2 := <-ch2: fmt.Println(s2) } } 复制代码`

我们把函数里的延时注释掉，主函数 select 之前加 2s 的延时以等待两个信道的数据准备好，select 会随机选取其中一个 case 执行，所以输出也是随机的。

#### default case ####

与 switch 语句类似，select 也有 default case，是的 select 语句不在阻塞，如果其他信道操作还没有准备好，将会直接执行 default 分支。

` func service1 (ch chan string ) { ch <- "from service1" } func service2 (ch chan string ) { ch <- "from service2" } func main () { ch1 := make ( chan string ) ch2 := make ( chan string ) go service1(ch1) go service2(ch2) select { // ch1 ch2 都还没有准备好，直接执行 default 分支 case s1 := <-ch1: fmt.Println(s1) case s2 := <-ch2: fmt.Println(s2) default : fmt.Println( "no case ok" ) } } 复制代码`

输出：no case ok 执行到 select 语句的时候，由于信道 ch1、ch2 都没准备好，直接执行 default 语句。

` func service1 (ch chan string ) { ch <- "from service1" } func service2 (ch chan string ) { ch <- "from service2" } func main () { ch1 := make ( chan string ) ch2 := make ( chan string ) go service1(ch1) go service2(ch2) time.Sleep(time.Second) // 延时 1s,等待 ch1 ch2 准备就绪 select { case s1 := <-ch1: fmt.Println(s1) case s2 := <-ch2: fmt.Println(s2) default : fmt.Println( "no case ok" ) } } 复制代码`

在 select 语句之前加了 1s 延时，等待 ch1 ch2 准备就绪。因为两个通道都准备好了，所以不会走 default 语句。随机输出 from service1 或 from service2。

#### nil channel ####

信道的默认值是 nil，不能对 nil 信道进行读写操作。看下面的例子

` func service1 (ch chan string ) { ch <- "from service1" } func main () { var ch chan string go service1(ch) select { case str := <-ch: fmt.Println(str) } } 复制代码`

报错：

` fatal error: all goroutines are asleep - deadlock! goroutine 1 [ select (no cases)]: goroutine 18 [ chan send ( nil chan )]: 复制代码`

有两个错误的地方需要注意： [select (no cases)] case 分支中如果信道是 nil， **该分支就会被忽略** ，那么上面就变成空 select{} 语句，阻塞主协程，调度 service1 协程，在 nil 信道上操作，便报[chan send (nil chan)] 错误。可以使用上面的 default case 避免发生这样的错误。

` func service1 (ch chan string ) { ch <- "from service1" } func main () { var ch chan string go service1(ch) select { case str := <-ch: fmt.Println(str) default : fmt.Println( "I am default" ) } } 复制代码`

输出：I am default

#### 添加超时时间 ####

有时候，我们不希望立即执行 default 语句，而是希望等待一段时间，若这个时间段内还没有可操作的信道，则执行规定的语句。可以在 case 语句后面设置超时时间。

` func service1 (ch chan string ) { time.Sleep( 5 * time.Second) ch <- "from service1" } func service2 (ch chan string ) { time.Sleep( 3 * time.Second) ch <- "from service2" } func main () { ch1 := make ( chan string ) ch2 := make ( chan string ) go service1(ch1) go service2(ch2) select { // 会发送阻塞 case s1 := <-ch1: fmt.Println(s1) case s2 := <-ch2: fmt.Println(s2) case <-time.After( 2 *time.Second): // 等待 2s fmt.Println( "no case ok" ) } } 复制代码`

输出：no case ok 在第三个 case 语句中设置了 2s 的超时时间，这 2s 内如果其他可操作的信道，便会执行该 case。

#### 空 select ####

` package main func main () { select {} } 复制代码`

我们知道 select 语句会发生阻塞，直到有 case 可以操作。但是空 select 语句没有 case 分支，所以便一直阻塞引起死锁。 报错：

` fatal error: all goroutines are asleep - deadlock! goroutine 1 [ select (no cases)] 复制代码`

希望这篇文章能够帮助你，Good day!

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