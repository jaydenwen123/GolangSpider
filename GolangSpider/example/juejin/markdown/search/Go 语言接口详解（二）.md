# Go 语言接口详解（二） #

> 
> 
> 
> **这是『就要学习 Go 语言』系列的第 20 篇分享文章**
> 
> 

提醒：文末给大家留了小练习，可以先看文章，再做练习，检验自己的学习成果！

我们接着上一篇，继续讲接口的其他用法。

### 实现多个接口 ###

一种类型可以实现多个接口，来看下例子：

` type Shape interface { Area() float32 } type Object interface { Perimeter() float32 } type Circle struct { radius float32 } func (c Circle) Area () float32 { return math.Pi * (c.radius * c.radius) } func (c Circle) Perimeter () float32 { return 2 * math.Pi * c.radius } func main () { c := Circle{ 3 } var s Shape = c var p Object = c fmt.Println( "area: " , s.Area()) fmt.Println( "perimeter: " , p.Perimeter()) } 复制代码`

[输出]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2Fmi756xNLlF7 ) ：

` area: 28.274334 perimeter: 18.849556 复制代码`

上面的代码，结构体 Circle 分别实现了 Shape 接口和 Object 接口，所以可以将结构体变量 c 赋给变量 s 和 p，此时 s 和 p 具有相同的动态类型和动态值，分别调用各自实现的方法 Area() 和 Perimeter()。 我们修改下程序：

` fmt.Println( "area: " , p.Area()) fmt.Println( "perimeter: " , s.Perimeter()) 复制代码`

编译会出错：

` p.Area undefined ( type Object has no field or method Area) s.Perimeter undefined ( type Shape has no field or method Perimeter) 复制代码`

为什么？因为 s 的静态类型是 Shape，而 p 的静态类型是 Object。那有什么解决办法吗？有的，我们接着看下一节

### 类型断言 ###

类型断言可以用来获取接口的底层值，通常的语法： **i.(Type)** ，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。

` type Shape interface { Area() float32 } type Object interface { Perimeter() float32 } type Circle struct { radius float32 } func (c Circle) Area () float32 { return math.Pi * (c.radius * c.radius) } func (c Circle) Perimeter () float32 { return 2 * math.Pi * c.radius } func main () { var s Shape = Circle{ 3 } c := s.(Circle) fmt.Printf( "%T\n" ,c) fmt.Printf( "%v\n" ,c) fmt.Println( "area: " , c.Area()) fmt.Println( "perimeter: " , c.Perimeter()) } 复制代码`

[输出]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FdHv0KVE92vG ) ：

` main.Circle { 3 } area: 28.274334 perimeter: 18.849556 复制代码`

上面的代码，我们可以通过 c 访问接口 s 的底层值，也可以通过 c 分别调用方法 Area() 和 Perimeter()，这就解决了上面遇到的问题。 在语法 i.(Type) 中，如果 Type 没有实现 i 所属的接口，编译的时候会报错；或者 i 的动态值不是 Type，则会报 panic 错误。怎么解决呢？可以使用下面的语法：

` value, ok := i.(Type) 复制代码`

使用上面的语法，Go 会自动检测上面提到的两种情况，我们只需要通过变量 ok 判断结果是否正确即可。如果正确，ok 为 true，否则为 false，value 为 Type 对应的零值。

### 类型选择 ###

类型选择用于将接口的具体类型与各种 case 语句中指定的多种类型进行匹配比较，有点类似于 switch case 语句，不同的是 case 中指定是类型。 类型选择的语法有点类似于类型断言的语法：i.(type)，其中 i 是接口，type 是固定关键字，使用这个可以获得接口的具体类型而不是值，每一个 case 中的类型必须实现了 i 接口。

` func switchType (i interface {}) { switch i.( type ) { case string : fmt.Printf( "string and value is %s\n" , i.( string )) case int : fmt.Printf( "int and value is %d\n" , i.( int )) default : fmt.Printf( "Unknown type\n" ) } } func main () { switchType( "Seekload" ) switchType( 27 ) switchType( true ) } 复制代码`

输出：

` string and value is Seekload int and value is 27 Unknown type 复制代码`

上面的代码应该很好理解，i 的类型匹配到哪个 case ，就会执行相应的输出语句。 **注意：只有接口类型才可以进行类型选择** 。其他类型，例如 int、string等是不能的：

` i := 1 switch i.( type ) { case int : println ( "int type" ) default : println ( "unknown type" ) } 复制代码`

报错：

` cannot type switch on non- interface value i ( type int ) 复制代码`

### 接口嵌套 ###

Go 语言中，接口不能去实现别的接口也不能继承，但是可以通过嵌套接口创建新接口。

` type Math interface { Shape Object } type Shape interface { Area() float32 } type Object interface { Perimeter() float32 } type Circle struct { radius float32 } func (c Circle) Area () float32 { return math.Pi * (c.radius * c.radius) } func (c Circle) Perimeter () float32 { return 2 * math.Pi * c.radius } func main () { c := Circle{ 3 } var m Math = c fmt.Printf( "%T\n" , m ) fmt.Println( "area: " , m.Area()) fmt.Println( "perimeter: " , m.Perimeter()) } 复制代码`

[输出]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FZmpHYxwxxhh ) ：

` main.Circle area: 28.274334 perimeter: 18.849556 复制代码`

上面的代码，通过嵌套接口 Shape 和 Object，创建了新的接口 Math。任何类型如果实现了接口 Shape 和 Object 定义的方法，则说类型也实现了接口 Math，例如我们创建的结构体 Circle。 主函数里面，定义了接口类型的变量 m，动态类型是结构体 Circle，注意下方法 Area 和 Perimeter 的调用方式，类似与访问嵌套结构体的成员。

### 使用指针接收者和值接收者实现接口 ###

在前面我们都是通过值接收者去实现接口的，其实还可以通过指针接收者实现接口。实现过程中还是有需要注意的地方，我们来看下：

` type Shape interface { Area() float32 } type Circle struct { radius float32 } type Square struct { side float32 } func (c Circle) Area () float32 { return math.Pi * (c.radius * c.radius) } func (s *Square) Area () float32 { return s.side * s.side } func main () { var s Shape c1 := Circle{ 3 } s = c1 fmt.Printf( "%v\n" ,s.Area()) c2 := Circle{ 4 } s = &c2 fmt.Printf( "%v\n" ,s.Area()) c3 := Square{ 3 } //s = c3 s = &c3 fmt.Printf( "%v\n" ,s.Area()) } 复制代码`

[输出]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2F2J25PMh7bgn ) ：

` 28.274334 50.265484 9 复制代码`

上面的代码，结构体 Circle 通过值接收者实现了接口 Shape。我们在方法那篇文章中已经讨论过了，值接收者的方法可以使用值或者指针调用，所以上面的 c1 和 c2 的调用方式是合法的。

结构体 Square 通过指针接收者实现了接口 Shape。如果将上方注释部分打开的话，编译就会出错：

` cannot use c3 ( type Square) as type Shape in assignment: Square does not implement Shape (Area method has pointer receiver) 复制代码`

从报错提示信息可以清楚看出，此时我们尝试将值类型 c3 分配给 s，但 c3 并没有实现接口 Shape。这可能会令我们有点惊讶，因为在方法中，我们可以直接通过值类型或者指针类型调用指针接收者方法。 记住一点： **对于指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的** 。但接口存储的具体值是不可寻址的，对于编译器无法自动获取 c3 的地址，于是程序报错。

关于接口的使用方法总结到这，希望这两篇文章能够给你带来帮助！

**作业：** 文章提到的类型断言： **i.(Type)** ，其中 i 是接口，Type 可以是类型或接口，如果 Type 是接口的话，表达式是什么意思呢？下面的程序输出什么？

` type Shape interface { Area() float32 } type Object interface { Perimeter() float32 } type Circle struct { radius float32 } func (c Circle) Area () float32 { return math.Pi * (c.radius * c.radius) } func main () { var s Shape = Circle{ 3 } value1,ok1 := s.(Shape) value2,ok2 := s.(Object) fmt.Println(value1,ok1) fmt.Println(value2,ok2) } 复制代码`

欢迎大家留言讨论！

（全文完）
> 
> 
> 
> 原创文章，若需转载请注明出处！
> 欢迎扫码关注公众号「 **Golang来啦** 」或者移步 [seekload.net](
> https://link.juejin.im?target=https%3A%2F%2Fseekload.net ) ，查看更多精彩文章。
> 
> 

**公众号「Golang来啦」给你准备了一份神秘学习大礼包，后台回复【电子书】领取！**

![公众号二维码](https://user-gold-cdn.xitu.io/2019/3/27/169be4a300f56486?imageView2/0/w/1280/h/960/ignore-error/1)