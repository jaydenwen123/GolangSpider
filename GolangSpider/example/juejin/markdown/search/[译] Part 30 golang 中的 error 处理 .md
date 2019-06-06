# [译] Part 30: golang 中的 error 处理 #

> 
> 
> 
> * 原文地址： [Part 30: Error Handling](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Ferror-handling%2F%23errortyperepresentation
> )
> * 原文作者： [Naveen R](
> https://link.juejin.im?target=https%3A%2F%2Fgolangbot.com%2Fabout%2F )
> * 译者：咔叽咔叽 转载请注明出处。
> 
> 
> 

## 什么是 Error ##

Error 表示程序中的异常情况。假设我们正在尝试打开文件，文件系统中不存在该文件，那么这是一种异常情况，它就代表一种 ` error` 。 Go 中使用内置的 ` error` 类型表示错误。 就像任何其他的内置类型，如 int，float64，... ` error` 可以存储在变量中，从函数返回等等。

## 例子 ##

用打开了一个不存在的文件的示例程序来解释一下。

` package main import ( "fmt" "os" ) func main () { f, err := os.Open( "/test.txt" ) if err != nil { fmt.Println(err) return } fmt.Println(f.Name(), "opened successfully" ) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2FZZeAGiDCJZQ )

上面的程序中，我们试图在路径/test.txt 中打开文件。 os 包的 Open 函数定义如下，

` func Open(name string) (file *File, err error) 复制代码`

如果文件被成功打开，则 Open 函数将会返回打开的文件，同时 err 为 nil。如果在打开文件时出错，err 将返回非 nil 的错误。

如果一个函数或者方法返回错误，那么按照惯例，函数返回值的最后面那一个值就是 err。所以，Open 函数返回的最后一个值是 err。

在 Go 中处理错误的惯用方法是将返回的错误与 nil 进行比较。 如果返回的 err 是 nil，则表示没有发生错误，非 nil 值则表示存在错误。在我们的例子中，我们在第 10 行检查错误的返回值，如果它不是 nil，我们只需打印错误并从 main 函数返回。

运行上述程序将打印，

` open /test.txt: No such file or directory 复制代码`

完美😃。我们打印了一条错误消息，显示了该文件不存在。

## Error type ##

让我们再深入一点，看看如何定义内置的错误类型。 error 是具有以下定义的接口类型，

` type error interface { Error() string } 复制代码`

它包含一个 Error 方法，实现此接口的 ` Error` 就可以被用作 ` error` 。此方法提供了 ` error` 的描述

当打印 ` error` 时， ` fmt.Println` 函数在内部调用 ` Error` 方法以获取 ` error` 的描述。上述例子的第 11 行就展示了 ` error` 描述的打印。

## 如何从 error 中提取更多的信息 ##

在上面的例子中我们看到打印了错误的描述。如果我们想要导致 ` error` 的文件的实际路径，该怎么办？一种可能的方法是解析错误字符串。这是输出，

` open /test.txt: No such file or directory 复制代码`

我们可以获得错误内容并获取导致错误的文件的文件路径为“/test.txt”，但这不是一种很好的方式。在有些情况下，我们应该捕获更多的错误信息，然后根据不同的错误去做区别处理。（类似其他的语言 catch 多个错误）

有没有办法可靠地获取文件名？答案是肯定的，标准的 Go 库使用不同的方式来提供有关 ` error` 的更多信息。让我们逐一看看它们。

#### 1. 断言结构类型并从结构的字段中获取更多信息 ####

如果仔细阅读 [Open]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fpkg%2Fos%2F%23OpenFile ) 函数的文档，可以看到其 ` error` 返回类型为 ` *PathError` 。 ` PathError` 是一种结构类型，它在标准库中的实现如下，

` type PathError struct { Op string Path string Err error } func (e *PathError) Error () string { return e.Op + " " + e.Path + ": " + e.Err.Error() } 复制代码`

如果有兴趣知道上述源代码的位置，可以在此处找到 [golang.org/src/os/erro…]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fsrc%2Fos%2Ferror.go%3Fs%3D653%3A716%23L11 )

从上面的代码中，可以知道 ` * PathError` 通过声明 ` Error` 方法来实现 ` error` 的接口。此方法返回一个包含路径，实际错误拼接的字符串并返回。因此我们收到了这个错误内容，

` open /test.txt: No such file or directory 复制代码`

` PathError struct` 的 ` Path` 字段包含导致错误的文件的路径。让我们修改上面编写的程序并打印路径。

` package main import ( "fmt" "os" ) func main () { f, err := os.Open( "/test.txt" ) if err, ok := err.(*os.PathError); ok { fmt.Println( "File at path" , err.Path, "failed to open" ) return } fmt.Println(f.Name(), "opened successfully" ) } 复制代码`

[Run in playgroud]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2FYQn7RX7KpPg )

在上面的程序中，在第 10 行我们使用类型断言从结构字段中去获取 ` error` 的基础值。然后我们在 11 行使用 ` err.Path` 打印路径。 输出如下，

` File at path /test.txt failed to open 复制代码`

太棒了。我们成功地使用类型断言从结构中获取到了文件路径。

#### 2. 断言结构类型并从结构的方法中获取更多信息 ####

第二种方法是通过断言结构类型并且通过调用结构类型的方法获取更多信息。

让我们通过一个例子更好地理解这一点。

标准库中的 ` DNSError` 结构类型定义如下，

` type DNSError struct { ... } func (e *DNSError) Error () string { ... } func (e *DNSError) Timeout () bool { ... } func (e *DNSError) Temporary () bool { ... } 复制代码`

从上面的代码中可以看出， ` DNSError` 结构有两个方法 ` Timeout` 和 ` Temporary` ，它们返回一个布尔值，指示 ` error` 是由于超时还是暂时的。

让我们编写一个断言 ` * DNSError` 类型的程序，并调用结构的方法来确定 ` error` 是暂时错误类型还是超时错误类型。

` package main import ( "fmt" "net" ) func main () { addr, err := net.LookupHost( "golangbot123.com" ) if err, ok := err.(*net.DNSError); ok { if err.Timeout() { fmt.Println( "operation timed out" ) } else if err.Temporary() { fmt.Println( "temporary error" ) } else { fmt.Println( "generic error: " , err) } return } fmt.Println(addr) } 复制代码`

注意： ` LookupHost` 在 playground 上不起作用。请在本地计算机上运行此程序。

在上面的程序中，我们在第 9 行试图获取一个无效域名 ` golangbot123.com` 的 IP 地址。在第 10 行，我们通过断言类型 ` * net.DNSError` 来获取结构的基础值。然后我们在第 10 和 13 行分别检查 ` error` 是由于超时还是临时引起的。

在我们的例子中，错误既不是暂时的也不是由于超时，因此程序将打印，

` generic error: lookup golangbot123.com: no such host 复制代码`

如果错误是临时的或由于超时，则执行相应的 if 语句，这样就可以分情况适当地处理它了。

#### 3. 直接比较 ####

获取有关错误的更多详细信息的第三种方法是直接与类型错误的变量进行比较。让我们通过一个例子来理解这一点。

` filepath` 包的 ` Glob` 函数用于返回与正则模式匹配的所有文件的名称。格式错误时，此函数返回错误 ` ErrBadPattern` 。

` ErrBadPattern` 在 ` filepath` 包中定义如下，

` var ErrBadPattern = errors.New( "syntax error in pattern" ) 复制代码`

` errors.New` 用于创建一个新的 ` error` 。我们将在下一个教程中详细讨论这个问题。

当匹配发生格式错误时， ` Glob` 函数返回 ` ErrBadPattern` 。

让我们编写一个小程序来检查这个 ` error` 。

` package main import ( "fmt" "path/filepath" ) func main () { files, error := filepath.Glob( "[" ) if error != nil && error == filepath.ErrBadPattern { fmt.Println(error) return } fmt.Println( "matched files" , files) } 复制代码`

[Run in palygroud]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2FYgX3hgjnFdi )

在上面的程序中，我们匹配带有 ` [` 的文件，这是一个格式错误的方式。我们检查 ` error` 是否为 ` nil` 。要获得有关 ` error` 的更多信息，我们在第 10 行直接将它与 ` filepath.ErrBadPattern` 进行比较。如果条件满足，则是因为格式错误。所以该程序将输出，

` syntax error in pattern 复制代码`

标准库使用上述方法提供有关 ` error` 的更多信息。我们将在下一个教程中使用这些方法来创建自己的自定义错误。

## 不要忽视 Error ##

永远不要忽视 ` error` 。忽略 ` error` 会引发麻烦。让我重写一个示例，该示例忽略了 ` error` 处理。

` package main import ( "fmt" "path/filepath" ) func main () { files, _ := filepath.Glob( "[" ) fmt.Println( "matched files" , files) } 复制代码`

[Run in playground]( https://link.juejin.im?target=http%3A%2F%2Fplay.flysnow.org%2Fp%2Fgf828BB4DEW )

我们从前面的例子中已经知道匹配是无效的。我在第 9 行通过使用 ` _` 标识符忽略了 ` Glob` 函数返回的 ` error` 。然后在第 10 行打印该匹配的文件。程序将打印，

` matched files [] 复制代码`

由于我们忽略了 ` error` ，似乎没有输出匹配文件，但实际上是模式本身的格式不正确，这样就不知道到底是什么导致了该 ` error` 的产生。所以不要忽视 ` error` 。

在本教程中，我们讨论了如何处理程序中发生的 ` error` 以及如何检查 ` error` 以从中获取更多信息。

在下一个教程中，我们将创建自己的自定义 ` error` ，并为标准 ` error` 添加更多上下文。