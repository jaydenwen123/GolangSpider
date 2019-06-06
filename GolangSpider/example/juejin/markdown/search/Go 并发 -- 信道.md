# Go 并发 -- 信道 #

> 
> 
> 
> **这是『就要学习 Go 语言』系列的第 22 篇分享文章**
> 
> 

上篇文章讲了关于协程的一些用法，比如如何创建协程、匿名协程等。这篇文章我们来讲讲信道。 **信道是协程之间通信的管道，从一端发送数据，另一端接收数据。**

### 信道声明 ###

使用信道之前需要声明，有两种方式：

` var c chan int // 方式一 c := make ( chan int ) // 方式二 复制代码`

使用关键字 chan 创建信道，声明时有类型，表明信道只允许该类型的数据传输。信道的零值为 nil。方式一就声明了 nil 信道。nil 信道没什么作用，既不能发送数据也不能接受数据。方式二使用 make 函数创建了可用的信道 c。

` func main () { c := make ( chan int ) fmt.Printf( "c Type is %T\n" ,c) fmt.Printf( "c Value is %v\n" ,c) } 复制代码`

输出：

` c Type is chan int c Value is 0xc000060060 复制代码`

上面的代码创建了信道 c，而且只允许 int 型数据传输。一般将信道作为参数传递给函数或者方法实现两个协程之间通信，有没有注意到信道 c 的值是一个地址，传参的时候直接使用 c 的值就可以，而不用取址。

### 信道的使用 ###

#### 读写数据 ####

Go 提供了语法方便我们操作信道：

` c := make ( chan int ) // 写数据 c <- data // 读数据 variable <- c // 方式一 <- c // 方式二 复制代码`

读写数据注意信道的位置， **信道在箭头的左边是写数据，在右边是从信道读数据** 。上面的方式二读数据是合理的，读出来的数据丢弃不使用。 注意： **信道操作默认是阻塞的，往信道里写数据之后当前协程便阻塞，直到其他协程将数据读出。一个协程被信道操作阻塞后，Go 调度器会去调用其他可用的协程，这样程序就不会一直阻塞** 。信道的这种特性非常有用，接下来我们就可以看到。 我们来温习下上篇文章的一个例子：

` func printHello () { fmt.Println( "hello world goroutine" ) } func main () { go printHello() time.Sleep( 1 *time.Second) fmt.Println( "main goroutine" ) } 复制代码`

这个例子 main() 协程使用了 time.Sleep() 函数休眠了 1s 等待 printHello() 执行完成。很黑科技，在生产环境绝对不可以这样用。我们使用信道修改下：

` func printHello (c chan bool ) { fmt.Println( "hello world goroutine" ) <- c // 读取信道的数据 } func main () { c := make ( chan bool ) go printHello(c) c <- true // main 协程阻塞 fmt.Println( "main goroutine" ) } 复制代码`

输出：

` hello world goroutine main goroutine 复制代码`

上面的例子，main 协程创建完 printHello 协程之后，第 8 行往信道 c 写数据，main 协程阻塞，Go 调度器调度可使用 printHello 协程，从信道 c 读出数据，main 协程接触阻塞继续运行。注意：读取操作没有阻塞是因为信道 c 已有可读的数据，否则，读取操作会阻塞。

#### 死锁 ####

前面提到过，读/写数据的时候信道会阻塞，调度器会去调度其他可用的协程。问题来了，如果没有其他可用的协程会发生什么情况？没错，就会发生著名的 **死锁** 。最简单的情况就是，只往信道写数据。

` func main () { c := make ( chan bool ) c <- true // 只写不读 fmt.Println( "main goroutine" ) } 复制代码`

报错：

` fatal error: all goroutines are asleep - deadlock! 复制代码`

同理，只读不写也会报同样的错误。

#### 关闭信道与 for loop ####

发送数据的信道有能力选择关闭信道，数据就不能传输。数据接收的时候可以返回一个状态判断该信道是否关闭：

` val, ok := <- channel 复制代码`

val 是接收的值，ok 标识信道是否关闭。为 true 的话，该信道还可以进行读写操作；为 false 则标识信道关闭，数据不能传输。 使用内置函数 close() 关闭信道。

` func printNums (ch chan int ) { for i := 0 ; i < 4 ; i++ { ch <- i } close (ch) } func main () { ch := make ( chan int ) go printNums(ch) for { v, ok := <-ch if ok == false { // 通过 ok 判断信道是否关闭 fmt.Println(v, ok) break } fmt.Println(v, ok) } } 复制代码`

输出：

` 0 true 1 true 2 true 3 true 0 false 复制代码`

printNums 协程写完数据之后关闭了信道，在 main 协程里对 ok 判断，若为 false 说明信道关闭，退出 for 循坏。从关闭的信道读出来的值是对应类型的零值，上面最后一行的输出值是 int 类型的零值 0。

使用 for 循环，需要手动判断信道有没有关闭。如果嫌烦的话，那就使用 for range 读取信道吧，信道关闭，for range 自动退出。

` func printNums (ch chan int ) { for i := 0 ; i < 4 ; i++ { ch <- i } close (ch) } func main () { ch := make ( chan int ) go printNums(ch) for v := range ch { fmt.Println(v) } } 复制代码`

输出：

` 0 1 2 3 复制代码`

提一点，使用 for range 一个信道，发送完毕之后必须 close() 信道，不然发生死锁。

#### 缓冲信道和信道容量 ####

之前创建的信道是无缓冲的，读写信道会立马阻塞当前协程。对于缓冲信道，写不会阻塞当前信道直到信道满了，同理，读操作也不会阻塞当前信道除非信道没数据。创建带缓冲的信道：

` ch := make ( chan type , capacity) 复制代码`

capacity 是缓冲大小，必须大于 0。 内置函数 len()、cap() 可以计算信道的长度和容量。

` func main () { ch := make ( chan int , 3 ) ch <- 7 ch <- 8 ch <- 9 //ch <- 10 // 注释打开的话，协程阻塞，发生死锁 会发生死锁：信道已满且没有其他可用信道读取数据 fmt.Println( "main stopped" ) } 复制代码`

输出：main stopped 创建了缓冲为 3 的信道，写入 3 个数据时信道不会阻塞。如果将第 7 行代码注释打开的话，此时信道已满，协程阻塞，又没有其他可用协程读数据，便发生死锁。 再来看个例子：

` func printNums (ch chan int ) { ch <- 7 ch <- 8 ch <- 9 fmt.Printf( "channel len:%d,capacity:%d\n" , len (ch), cap (ch)) fmt.Println( "blocking..." ) ch <- 10 // 阻塞 close (ch) } func main () { ch := make ( chan int , 3 ) go printNums(ch) // 休眠 2s time.Sleep( 2 *time.Second) for v := range ch { fmt.Println(v) } fmt.Println( "main stopped" ) } 复制代码`

输出：

` channel len : 3 ,capacity: 3 blocking... 7 8 9 10 main stopped 复制代码`

休眠 2s 的目的是让信道写满数据发生阻塞，从打印结果可以看出。2s 之后，主协程从信道读取数据，信道容量有余阻塞便解除，继续写数据。

如果缓冲信道是关闭状态但有数据，仍然可以读取数据：

` func main () { ch := make ( chan int , 3 ) ch <- 7 ch <- 8 //ch <- 9 close (ch) for v := range ch { fmt.Println(v) } fmt.Println( "main stopped" ) } 复制代码`

输出：

` 7 8 main stopped 复制代码`

#### 单向信道 ####

之前创建的都是双向信道，既能发送数据也能接收数据。我们还可以创建单向信道，只发送或者只接收数据。 语法：

` sch := make ( chan <- int ) rch := make (<- chan int ) 复制代码`

sch 是只发送信道，rch 是只接受信道。 这种单向信道有什么用呢？我们总不能只发不接或只接不发吧。这种信道主要用在信道作为参数传递的时候，Go 提供了自动转化，双向转单向。 重写之前的例子：

` func printNums (ch chan <- int ) { for i := 0 ; i < 4 ; i++ { ch <- i } close (ch) } func main () { ch := make ( chan int ) go printNums(ch) for v := range ch { fmt.Println(v) } } 复制代码`

输出：

` 0 1 2 3 复制代码`

main 协程中 ch 是一个双向信道，printNums() 在接收参数的时候将 ch 自动转成了单向信道，只发不收。但在 main 协程中，ch 仍然可以接收数据。 **使用单向通道主要是可以提高程序的类型安全性，程序不容易出错。**

#### 信道数据类型 ####

信道是一类值，类似于 int、string 等，可以像其他值一样在任何地方使用，比如作为结构体成员、函数参数、函数返回值，甚至作为另一个通道的类型。我们来看下 **使用通道作为另一个通道的数据类型** 。

` func printWord (ch chan string ) { fmt.Println( "Hello " + <-ch) } func productCh (ch chan chan string ) { c := make ( chan string ) // 创建 string type 信道 ch <- c // 传输信道 } func main () { // 创建 chan string 类型的信道 ch := make ( chan chan string ) go productCh(ch) // c 是 string type 的信道 c := <-ch go printWord(c) c <- "world" fmt.Println( "main stopped" ) } 复制代码`

输出：

` Hello world main stopped 复制代码`

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