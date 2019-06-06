# Golang并发之共享内存变量 #

应该说，无论使用哪一种编程语言开发应用程序，并发编程都是复杂的，而Go语言内置的并发支持，则让Go并发编程变得很简单。

## CSP ##

CSP，即顺序通信进程，是Go语言中原生支持的并发模型，一般使用goroutine和channel来实现，CSP的编程思想是“通过通信共享内存，而不是通过共享内存来通信”，因此使用CSP思想来开发并发程序，一般是使用channel串联多个goroutine，最终达到多个goroutine顺序执行的目的。

## 共享内存变量 ##

我们知道，单个的goutine代码是顺序执行，而并发编程时，创建多个goroutine，但我们并不能确定不同的goroutine之间的执行顺序，多个goroutine之间大部分情况是代码交叉执行，在执行过程中，可能会修改或读取共享内存变量，这样就会产生 ` 数据竞争` ,发生一些意外之外的结果。

` package main var balance int //存款 func Deposit(amount int) { balance = balance + amount } //读取余额 func Balance() int { return balance } func main (){ //小王：存600，并读取余额 go func (){ Deposit(600) fmt.Println(Balance()) }() //小张：存500 go func (){ Deposit(500) }() time.Sleep(time.Second) fmt.Println(balance) } 复制代码`

上面的例子叫银行存款问题，是演示并发的经典例子。

一般我们认为，这个例子的运行结果只有三种：

* 小王存600，小王读取余额为600，小张再存500，总金额为1100
* 小张存500，小王存600，小王读取余额为500，总金额为1100
* 小王存600，小张存500，小王读取余额为500，总金额为1100

上面的情况，都是假设存款操作是顺序的，但是，还存在一种情况，也就是小王或小张并发执行存款操作，这时候会发生存款金额丢失的风险。

## 锁 ##

看到上面的例子之后，我们知道数据竞争会产生严重的后果，那如何避免数据竞争呢？有三种：

* 通过channel串联goroutine，达到顺序执行的效果，避免竞争。
* 不在并发程序中修改共享变量，这当然是不太可能的情况。
* 通过使用锁，使用同一时间只有一个goroutine可以修改内存中的变量，也就使用不同goroutine修改变量时发生互斥行为。

### sync.Mutex:互斥锁 ###

我们可以使用Go语言提供的互斥锁来避免上述的数据竞争行为的发生,可以把代码进行相应的修改：

` mu sync.Mutex // 声明一个互斥锁 func Deposit(amount int) { mu.Lock()//获取锁 balance = balance + amount mu.Unlock()//释放锁 } //读取余额 func Balance() int { mu.Lock()//获取锁 return balance mu.Unlock()//释放锁 } 复制代码`

### sync.RWMutex:读写锁 ###

当我们使用Mutex互斥锁的时候，那么无论是读取还是修改，都需要等待其他goroutine释放锁，但是读取相对修改来是，是安全的操作，Go提供了另外一种锁，sync.RWMutex，读写锁，这种锁，多个读取的时候，不会锁，只有修改时候，需要等到所有读取的锁释放，才能修改，所以我们可以把Balance()函数修改为：

` rmu sync.RWMutex func Balance() int { rmu.RLock()//获取读锁 return balance rmu.RUnlock()//释放读锁 } 复制代码`

### 更好地使用锁 ###

上面的例子中，我们都是在函数后面释放锁的，但实际开发中，函数的代码很长，有各种判断，我们无法保证函数能执行到最后，并成功释放锁，如果中发生错误，无法释放锁，就造成其他goroutine的阻塞，因此可以使用defer关键字，让函数无论如何都会释放锁。

` package main import "sync" var balance int mu sync.Mutex // 声明一个互斥锁 rmu sync.RWMutex //存款 func Deposit(amount int) { mu.Lock()//获取锁 balance = balance + amount mu.Unlock()//释放锁 } //读取余额 func Balance() int { rmu.RLock()//获取读锁 return balance rmu.RUnlock()//释放读锁 } func main (){ //小王：存600，并读取余额 go func (){ Deposit(600) fmt.Println(Balance()) }() //小张：存500 go func (){ Deposit(500) }() time.Sleep(time.Second) fmt.Println(balance) } 复制代码`

## 总结 ##

当然，在实际项目并发编程的时候，我们遇到的情况要远比上述例子复杂得多，因此还要多多练习，让自己对并发有更学层次的理解。