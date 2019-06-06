# 【跟着我们学Golang】之异常处理 #

Java中的异常分为Error和Exception来处理，这里也以错误和异常两种，来分别讲一讲Go的异常处理。

> 
> 
> 
> Go 语言没有类似 Java 或 .NET 中的异常处理机制，虽然可以使用 defer、panic、recover 模拟，但官方并不主张这样做。Go
> 语言的设计者认为其他语言的异常机制已被过度使用，上层逻辑需要为函数发生的异常付出太多的资源。同时，如果函数使用者觉得错误处理很麻烦而忽略错误，那么程序将在不可预知的时刻崩溃。
> Go 语言希望开发者将错误处理视为正常开发必须实现的环节，正确地处理每一个可能发生错误的函数。同时，Go
> 语言使用返回值返回错误的机制，也能大幅降低编译器、运行时处理错误的复杂度，让开发者真正地掌握错误的处理。 -- 摘自： [C语言中文网](
> https://link.juejin.im?target=http%3A%2F%2Fc.biancheng.net%2Fview%2F62.html
> )
> 
> 

## error接口 ##

Go处理错误的思想

通过返回error接口的方式来处理函数的错误，在调用之后进行错误的检查。如果调用该函数出现错误，就返回error接口的实现，指出错误的具体内容，如果成功，则返回nil作为error接口的实现。

error接口声明了一个 ` Error() string` 的函数，实际使用时使用相应的接口实现，由函数返回error信息，函数的调用之后进行错误的判断从而进行处理。

` // The error built-in interface type is the conventional interface for // representing an error condition, with the nil value representing no error. type error interface { Error() string } 复制代码`

Error() 方法返回错误的具体描述，使用者可以通过这个字符串知道发生了什么错误。下面看一个例子。

` package main import ( "errors" "fmt" ) func main () { sources := []string{ "hello" , "world" , "souyunku" , "gostack" } fmt.Println(getN(0, sources))//直接调用，会打印两项内容，字符串元素以及error空对象 fmt.Println(getN(1, sources)) fmt.Println(getN(2, sources)) fmt.Println(getN(3, sources)) target, err := getN(4, sources)//将返回结果赋值 if err != nil {//常见的错误处理，如果error不为nil，则进行错误处理 fmt.Println(err) return } fmt.Println(target) } //定义函数获取第N个元素，正常返回元素以及为nil的error，异常返回空元素以及error func getN(n int, sources []string) (string, error) { if n > len(sources)-1 { return "" , fmt.Errorf( "%d, out of index range %d" , n, len(sources) - 1) } return sources[n], nil } /* 打印内容： hello <nil> world <nil> souyunku <nil> gostack <nil> 4, out of index range 3 */ 复制代码`

常见的错误处理就是在函数调用结束之后进行error的判断，确定是否出现错误，如果出现错误则进行相应的错误处理；没有错误就继续执行下面的逻辑。

遇到多个函数都带有error返回的时候，都需要进行error的判断，着实会让人感到非常的苦恼，但是它的作用是很好的，其鲁棒性也要比其他静态语言要好的多。

### 自定义error ###

身为一个接口，任何定义实现了 ` Error() string` 函数，都可以认为是error接口的实现。所以可以自己定义具体的接口实现来满足业务的需求。

error接口的实现有很多，各大项目也都喜欢自己实现error接口供自己使用。最常用的是官方的error包下的 ` errorString` 实现。

` // Copyright 2011 The Go Authors. All rights reserved. // Use of this source code is governed by a BSD-style // license that can be found in the LICENSE file. // Package errors implements functions to manipulate errors. package errors // New returns an error that formats as the given text. func New(text string) error { return &errorString{text} } // errorString is a trivial implementation of error. type errorString struct { s string } func (e *errorString) Error() string { return e.s } 复制代码`

可以看到，官方error包通过定义了 ` errorString` 来实现了error接口，在使用的时候通过 ` New(text string) error` 这个函数进行调用从而返回error接口内容（该函数在返回的时候是一个 ` errorString` 类型的指针，但是定义的返回内容是error接口类型，这也举例说明了上节讲到的接口的内容）下面看例子。

` package main import ( "errors" "fmt" ) func main () { //直接使用errors.New来定义错误消息 notFound := errors.New( "404 not found" ) fmt.Println(notFound) //也可以使用fmt包中包装的Errorf来添加 fmt.Println(fmt.Errorf( "404: page %v is not found" , "index.html" )) } /* 打印内容 404 not found 404: page index.html is not found */ 复制代码`

自己试着实现一个404notfound的异常

` type NOTFoundError struct { name string } func (e *NOTFoundError) Error() string { return fmt.S printf ( "%s is not found, please new again" , e.name) } func NewNotFoundError(name string) error{ return &NOTFoundError{name} } func runDIYError () { err := NewNotFoundError( "your girl" ) // 根据switch，确定是哪种error switch err.( type ) { case *NOTFoundError: fmt.Printf( "error : %v \n" ,err) default: // 其他类型的错误 fmt.Println( "other error" ) } } /**调用runDIYError()结果 error : your girl is not found, please new again */ 复制代码`

自己定义异常NotFoundError只是简单的实现 ` Error() string` 函数，并在出错的时候提示内容找不到，不支持太多的功能，如果业务需要，还是可以继续扩展。

## defer ##

在将panic和recover之前插播一下defer这个关键字，这个关键字在panic和recover中也会用到。

defer的作用就是指定某个函数在执行return之前在执行，而不是立即执行。下面是defer的语法

` defer func (){}() 复制代码`

defer指定要执行的函数，或者直接声明一个匿名的函数并直接执行。这个还是结合实例进行了解比较合适。

` func runDefer (){ defer func () { fmt.Println( "3" ) }()//括号表示定义 function 之后直接执行 fmt.Println( "1" ) defer func(index string) { fmt.Println(index) }( "2" )//括号表示定义 function 之后直接执行，如果定义的 function 包含参数，括号中也要进行相应的赋值操作 } /** 执行结果： 1 2 3 */ 复制代码`

执行该函数能看到顺序打印出了123三个数字，这就是defer的执行过程。其特点就是LIFO，先进后出，先指定的函数总是在后面执行，是一个逆序的执行过程。

defer在Go中也是经常被用到的，而且设计的极其巧妙，举个例子

` file.Open() defer file.Close()//该语句紧跟着file.Open()被指定 file.Lock() defer file.Unclock()// 该语句紧跟着file.Lock()被指定 复制代码`

像这样需要开关或者其他操作必须执行的操作都可以在相邻的行进行执行指定，可以说很好的解决了那些忘记执行Close操作的痛苦。

### defer面试题 ###

` package main import ( "fmt" ) func main () { defer_call() } func defer_call () { defer func () { fmt.Println( "打印前" ) }() defer func () { fmt.Println( "打印中" ) }() defer func () { fmt.Println( "打印后" ) }() panic( "触发异常" ) } 考点：defer执行顺序 解答： defer 是后进先出。 panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。 结果： 打印后 打印中 打印前 panic: 触发异常 --- //摘自：https://blog.csdn.net/weiyuefei/article/details/77963810 复制代码` ` func calc(index string, a, b int) int { ret := a + b fmt.Println(index, a, b, ret) return ret } func main () { a := 1 b := 2 defer calc( "1" , a, calc( "10" , a, b)) a = 0 defer calc( "2" , a, calc( "20" , a, b)) b = 1 } 考点：defer执行顺序 解答： 这道题类似第1题 需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用calc( "10" ,1,2)==>10,1,2,3 执行index:2时,与之前一样，需要先调用calc( "20" ,0,2)==>20,0,2,2 执行到b=1时候开始调用，index:2==>calc( "2" ,0,2)==>2,0,2,2 最后执行index:1==>calc( "1" ,1,3)==>1,1,3,4 结果： 10 1 2 3 20 0 2 2 2 0 2 2 1 1 3 4 --- 摘自： https://blog.csdn.net/weiyuefei/article/details/77963810 复制代码`

defer 虽然是基础知识，其调用过程也非常好理解，但是往往在面试的过程中会出现一些比较绕的题目，这时候不要惊慌，只需要好好思考其执行的过程还是可以解出来的。

## panic & recover ##

panic英文直译是 恐慌 ，在Go中意为程序出现了崩溃。recover直译是 恢复 ，其目的就是恢复恐慌。

> 
> 
> 
> 在其他语言里，宕机往往以异常的形式存在。底层抛出异常，上层逻辑通过 try/catch
> 机制捕获异常，没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行。 Go 没有异常系统，其使用 panic
> 触发宕机类似于其他语言的抛出异常，那么 recover 的宕机恢复机制就对应 try/catch 机制。-- 摘自： [C语言中文网](
> https://link.juejin.im?target=http%3A%2F%2Fc.biancheng.net%2Fview%2F62.html
> )
> 
> 

### panic ###

程序崩溃就像遇到电脑蓝屏时一样，大家都不希望遇到这样的情况。但有时程序崩溃也能终止一些不可控的情况，以此来做出防范。出于学习的目的，咱们简单了解一下panic造成的崩溃，以及如何处理。先看一下panic的定义。

` // The panic built-in function stops normal execution of the current // goroutine. When a function F calls panic, normal execution of F stops // immediately. Any functions whose execution was deferred by F are run in // the usual way, and then F returns to its caller. To the caller G, the // invocation of F then behaves like a call to panic, terminating G 's // execution and running any deferred functions. This continues until all // functions in the executing goroutine have stopped, in reverse order. At // that point, the program is terminated and the error condition is reported, // including the value of the argument to panic. This termination sequence // is called panicking and can be controlled by the built-in function // recover. func panic(v interface{}) 复制代码`

从定义中可以了解到，panic可以接收任何类型的数据。而接收的数据可以通过recover进行获取，这个后面recover中进行讲解。

从分类上来说，panic的触发可以分为两类，主动触发和被动触发。

在程序运行期间，主动执行panic可以提前中止程序继续向下执行，避免造成更恶劣的影响。同时还能根据打印的信息进行问题的定位。

` func runSimplePanic (){ defer func () { fmt.Println( "before panic" ) }() panic( "simple panic" ) } /** 调用runSimplePanic()函数结果： before panic panic: simple panic goroutine 1 [running]: main.runSimplePanic() /Users/fyy/go/src/github.com/souyunkutech/gosample/chapter6/main.go:102 +0x55 main.main() /Users/fyy/go/src/github.com/souyunkutech/gosample/chapter6/main.go:18 +0x22 */ 复制代码`

从运行结果中能看到，panic执行后先执行了defer中定义的函数，再打印的panic的信息，同时还给出了执行panic的具体行（行数需要针对具体代码进行定论），可以方便的进行检查造成panic的原因。

还有在程序中不可估计的panic，这个可以称之为被动的panic，往往由于空指针和数组下标越界等问题造成。

` func runBePanic (){ fmt.Println(ss[100])//ss集合中没有下标为100的值，会造成panic异常。 } /** 调用runBePanic()函数结果： panic: runtime error: index out of range goroutine 1 [running]: main.runBePanic(...) /Users/fyy/go/src/github.com/souyunkutech/gosample/chapter6/main.go:106 main.main() /Users/fyy/go/src/github.com/souyunkutech/gosample/chapter6/main.go:21 +0x10f */ 复制代码`

从运行结果中看到，数组下标越界，直接导致panic，panic信息也是有Go系统运行时runtime所提供的信息。

### recover ###

先来简单看一下recover的注释。

` // The recover built-in function allows a program to manage behavior of a // panicking goroutine. Executing a call to recover inside a deferred // function (but not any function called by it) stops the panicking sequence // by restoring normal execution and retrieves the error value passed to the // call of panic. If recover is called outside the deferred function it will // not stop a panicking sequence. In this case , or when the goroutine is not // panicking, or if the argument supplied to panic was nil, recover returns // nil. Thus the return value from recover reports whether the goroutine is // panicking. func recover() interface{} 复制代码`

注释指明recover可以管理panic，通过defer定义在panic之前的函数中的recover，可以正确的捕获panic造成的异常。

结合panic来看一下recover捕获异常，并继续程序处理的简单实现。

` import "fmt" func main() { runError() fmt.Println("---------------------------") runPanicError() } type Student struct { Chinese int Math int English int } var ss = []Student{{100, 90, 89}, {80, 80, 80}, {70, 80, 80}, {70, 80, 60}, {90, 80, 59}, {90, 40, 59}, {190, 40, 59}, {80, 75, 66}, } func runError() { i := 0 for ; i < len(ss); i++ { flag, err := checkStudent(&ss[i]) if err != nil { fmt.Println(err) return }//遇到异常数据就会立即返回，不能处理剩余的数据 //而且，正常逻辑中参杂异常处理，使得程序并不是那么优雅 fmt.Printf("student %#v,及格? ：%t \n", ss[i], flag) } } func checkStudent(s *Student) (bool, error) { if s.Chinese > 100 || s.Math > 100 || s.English > 100 { return false, fmt.Errorf("student %#v, something error", s) } if s.Chinese > 60 && s.Math > 60 && s.English > 60 { return true, nil } return false, nil } func runPanicError() { i := 0 defer func() { if err := recover(); err != nil { fmt.Println(err) } i ++//跳过异常的数据，继续处理剩余的数据 for ; i < len(ss); i ++ { fmt.Printf("student %#v,及格? ：%t \n", ss[i], checkStudentS(&ss[i])) } }() for ; i < len(ss); i++ { fmt.Printf("student %#v,及格? ：%t \n", ss[i], checkStudentS(&ss[i])) } } func checkStudentS(s *Student) bool { if s.Chinese > 100 || s.Math > 100 || s.English > 100 { panic(fmt.Errorf("student %#v, something error", s)) } if s.Chinese > 60 && s.Math > 60 && s.English > 60 { return true } return false } 结果： student main.Student{Chinese:100, Math:90, English:89},及格? ：true student main.Student{Chinese:80, Math:80, English:80},及格? ：true student main.Student{Chinese:70, Math:80, English:80},及格? ：true student main.Student{Chinese:70, Math:80, English:60},及格? ：false student main.Student{Chinese:90, Math:80, English:59},及格? ：false student main.Student{Chinese:90, Math:40, English:59},及格? ：false student &main.Student{Chinese:190, Math:40, English:59}, something error --------------------------- student main.Student{Chinese:100, Math:90, English:89},及格? ：true student main.Student{Chinese:80, Math:80, English:80},及格? ：true student main.Student{Chinese:70, Math:80, English:80},及格? ：true student main.Student{Chinese:70, Math:80, English:60},及格? ：false student main.Student{Chinese:90, Math:80, English:59},及格? ：false student main.Student{Chinese:90, Math:40, English:59},及格? ：false student &main.Student{Chinese:190, Math:40, English:59}, something error student main.Student{Chinese:80, Math:75, English:66},及格? ：true 复制代码`

从结果中可以看出runPanicError函数将全部正常的数据都输出了，并给出了是否及格的判断，runError并没有全部将数据输出，而是遇到错误就中止了后续的执行，导致了执行的不够彻底。

panic和recover的用法虽然简单，但是一般程序中用到的却很少，除非你对panic有着很深的了解。但也可以通过Panic来很好的美化自己的代码，从程序上看，runPanicError中的异常处理与正常逻辑区分开，也使得程序看起来非常的舒畅-_-!

> 
> 
> 
> 相对于那些对panic和recover掌握非常好的人来说，panic和recover能随便用，真的可以御剑飞行那种；但是如果掌握不好的话，还是尽可能的使用相对简单但不失高效又能很好的解决问题的error来处理就好了，以此来避免过度的使用从而造成的意外影响。毕竟我们的经验甚少，复杂的事物还是交给真正的大佬比较合适。
> 
> 
> 

## 总结 ##

Go中的异常处理相对比Java这些有着相对完善的错误处理机制的语言来说，还是显得非常的低级的，这也是Go一直被大家诟病的一点，但Go的更新计划中也有针对异常处理的改善，相信用不了多久就能看到不一样的错误处理机制。

源码可以通过'github.com/souyunkutech/gosample'获取。

## 关注我们的「微信公众号」 ##

![](https://user-gold-cdn.xitu.io/2019/5/28/16afcc22b47206b9?imageView2/0/w/1280/h/960/ignore-error/1)

首发微信公众号：Go技术栈，ID：GoStack

版权归作者所有，任何形式转载请联系作者。

作者：搜云库技术团队

出处： [gostack.souyunku.com/2019/05/27/…]( https://link.juejin.im?target=https%3A%2F%2Fgostack.souyunku.com%2F2019%2F05%2F27%2Ferror )