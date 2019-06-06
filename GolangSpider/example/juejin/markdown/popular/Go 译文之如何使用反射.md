# Go 译文之如何使用反射 #

作者：Jon Bodner

地址： [Learning to Use Go Reflection]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fcapital-one-tech%2Flearning-to-use-go-reflection-822a0aed74b7 )

# 什么是反射 #

多数情况下，Go 中的变量、类型和函数的使用都是非常简单的。

当你需要一个类型，定义如下：

` type Foo struct { A int B string } 复制代码`

当你需要一个变量，定义如下：

` var x Foo 复制代码`

当你需要一个函数，定义如下：

` func DoSomething (f Foo) { fmt.Println(f.A, f.B) } 复制代码`

但有时候，你想使用的变量依赖于运行时信息，它们在编程时并不存在。比如数据来源于文件，或来源于网络，你想把它映射到一个变量，而它们可能是不同的类型。在这类场景下，你就需要用到反射。反射让你可以在运行时检查类型，创建、更新、检查变量以及组织结构。

Go 中的反射主要围绕着三个概念：类型（Types）、类别（Kinds）和值（Values）。反射的实现源码位于 Go 标准库 reflection 包中。

# 检查类型 #

首先，让我们来看看类型（Types）。你可以通过 reflect.TypeOf(var) 形式的函数调用获取变量的类型，它会返回一个类型为 reflect.Type 的变量，reflect.Type 中的操作方法涉及了定义该类型变量的各类信息。

我们要看的第一个方法是 Name()，它返回的是类型的名称。有些类型，比如 slice 或 指针，没有类型名称，那么将会返回空字符串。

下一个介绍方法是 Kind()，我的观点，这是第一个真正有用的方法。Kind，即类别，比如切片 slice、映射 map、指针 pointer、结构体 struct、接口 interface、字符串 string、数组 array、函数 function、整型 int、或其他的基本类型。type 和 kind 是区别不是那么容易理清楚，但是可以这么想：

当你定义一个名称为 Foo 的结构体，那么它的 kind 是 struct，而它的 type 是 Foo。

当使用反射时，我们必须要意识到：在使用 reflect 包时，会假设你清楚的知道自己在做什么，如果使用不当，将会产生 panic。举个例子，你在 int 类型上调用 struct 结构体类型上才用的方法，你的代码就会产生 panic。我们时刻要记住，什么类型有有什么方法可以使用，从而避免产生 panic。

如果一个变量是指针、映射、切片、管道、或者数组类型，那么这个变量的类型就可以调用方法 varType.Elem()。

如果一个变量是结构体，那么你就可以使用反射去得到它的字段个数，并且可以得到每个字段的信息，这些信息包含在 reflect.StructField 结构体中。reflect.StructField 包含字段的名称、排序、类型、标签。

前言万语也不如一行代码看的明白，下面的这个例子输出了不同变量所属类型的信息。

` type Foo struct { A int `tag1:"First Tag" tag2:"Second Tag"` B string } func main () { sl := [] int { 1 , 2 , 3 } greeting := "hello" greetingPtr := &greeting f := Foo{A: 10 , B: "Salutations" } fp := &f slType := reflect.TypeOf(sl) gType := reflect.TypeOf(greeting) grpType := reflect.TypeOf(greetingPtr) fType := reflect.TypeOf(f) fpType := reflect.TypeOf(fp) examiner(slType, 0 ) examiner(gType, 0 ) examiner(grpType, 0 ) examiner(fType, 0 ) examiner(fpType, 0 ) } func examiner (t reflect.Type, depth int ) { fmt.Println(strings.Repeat( "\t" , depth), "Type is" , t.Name(), "and kind is" , t.Kind()) switch t.Kind() { case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice: fmt.Println(strings.Repeat( "\t" , depth+ 1 ), "Contained type:" ) examiner(t.Elem(), depth+ 1 ) case reflect.Struct: for i := 0 ; i < t.NumField(); i++ { f := t.Field(i) fmt.Println(strings.Repeat( "\t" , depth+ 1 ), "Field" , i+ 1 , "name is" , f.Name, "type is" , f.Type.Name(), "and kind is" , f.Type.Kind()) if f.Tag != "" { fmt.Println(strings.Repeat( "\t" , depth+ 2 ), "Tag is" , f.Tag) fmt.Println(strings.Repeat( "\t" , depth+ 2 ), "tag1 is" , f.Tag.Get( "tag1" ), "tag2 is" , f.Tag.Get( "tag2" )) } } } } 复制代码`

输出如下：

` Type is and kind is slice Contained type : Type is int and kind is int Type is string and kind is string Type is and kind is ptr Contained type : Type is string and kind is string Type is Foo and kind is struct Field 1 name is A type is int and kind is int Tag is tag1: "First Tag" tag2: "Second Tag" tag1 is First Tag tag2 is Second Tag Field 2 name is B type is string and kind is string Type is and kind is ptr Contained type : Type is Foo and kind is struct Field 1 name is A type is int and kind is int Tag is tag1: "First Tag" tag2: "Second Tag" tag1 is First Tag tag2 is Second Tag Field 2 name is B type is string and kind is string 复制代码`

[运行示例]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FlZ97yAUHxX )

# 创建实例 #

除了检查变量的类型外，你还可以利用来获取、设置和创建变量。首先，通过 refVal := reflect.ValueOf(var) 创建类型为 reflect.Value 的实例。如果你想通过反射来更新值，那么必须要获取到变量的指针 refPtrVal := reflect.ValueOf(&var)，如果不这么做，那么你只能读取值，而不能设置值。

一旦得到变量的 reflect.Value，你就可以通过 Value 的 Type 属性获取变量的 reflect.Type 类型信息。

如果想更新值，记住要通过指针，而且在设置时，要先取消引用，通过 refPtrVal.Elem().Set(newRefVal) 更新其中的值，传递给 Set 的参数也必须要是 reflect.Value 类型。

如果想创建一个新的变量，可以通过 reflect.New(varType) 实现，传递的参数是 reflect.Type 类型，该方法将会返回一个指针，如前面介绍的那样，你可以通过使用 Elem().Set() 来设置它的值。

最终，通过 Interface() 方法，你就得到一个正常的变量。Go 中没有泛型，变量的类型将会丢失，Interface() 方法将会返回一个类型为 interface{} 的变量。如果你为了能更新值，创建的是一个指针，那么需要使用 Elem().Interface() 来获取变量。但无论是上面的哪种情况，你都需要把 interface{} 类型变量转化为实际的类型，如此才能使用。

下面是一些代码，实现了这些概念。

` type Foo struct { A int `tag1:"First Tag" tag2:"Second Tag"` B string } func main () { greeting := "hello" f := Foo{A: 10 , B: "Salutations" } gVal := reflect.ValueOf(greeting) // not a pointer so all we can do is read it fmt.Println(gVal.Interface()) gpVal := reflect.ValueOf(&greeting) // it’s a pointer, so we can change it, and it changes the underlying variable gpVal.Elem().SetString( "goodbye" ) fmt.Println(greeting) fType := reflect.TypeOf(f) fVal := reflect.New(fType) fVal.Elem().Field( 0 ).SetInt( 20 ) fVal.Elem().Field( 1 ).SetString( "Greetings" ) f2 := fVal.Elem().Interface().(Foo) fmt.Printf( "%+v, %d, %s\n" , f2, f2.A, f2.B) } 复制代码`

输出如下：

` hello goodbye {A: 20 B:Greetings}, 20 , Greetings 复制代码`

[运行示例]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FPFcEYfZqZ8 )

# 无 make 的创建实例 #

对于像 slice、map、channel类型，它们需要用 make 创建实例，你也可以使用反射实现。slice 使用 reflect.MakeSlice，map 使用 reflect.MakeMap，channel 使用 reflect.MakeChan，你需要提供将创建变量的类型，即 reflect.Type，传递给这些函数。成功调用后，你将得到一个类型为 reflect.Value 的变量，你可以通过反射操作这个变量，操作完成后，就 可以将它转化为正常的变量。

` func main () { // declaring these vars, so I can make a reflect.Type intSlice := make ([] int , 0 ) mapStringInt := make ( map [ string ] int ) // here are the reflect.Types sliceType := reflect.TypeOf(intSlice) mapType := reflect.TypeOf(mapStringInt) // and here are the new values that we are making intSliceReflect := reflect.MakeSlice(sliceType, 0 , 0 ) mapReflect := reflect.MakeMap(mapType) // and here we are using them v := 10 rv := reflect.ValueOf(v) intSliceReflect = reflect.Append(intSliceReflect, rv) intSlice2 := intSliceReflect.Interface().([] int ) fmt.Println(intSlice2) k := "hello" rk := reflect.ValueOf(k) mapReflect.SetMapIndex(rk, rv) mapStringInt2 := mapReflect.Interface().( map [ string ] int ) fmt.Println(mapStringInt2) } 复制代码`

输出如下：

` [10] map[hello:10] 复制代码`

[运行示例]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2Fz4tnyEf6bH )

# 创建函数 #

你不仅经可以通过反射创建空间存储数据，还可以通过反射提供的函数 reflect.MakeFunc 来创建新的函数。这个函数期待接收参数有两个，一个是 reflect.Type 类型，并且 Kind 为 Function，另外一个是闭包函数，它的输入参数类型是 []reflect.Value，输出参数是 []reflect.Value。

下面是一个快速体验示例，可为任何函数在外层包裹一个记录执行时间的函数。

` func MakeTimedFunction (f interface {}) interface {} { rf := reflect.TypeOf(f) if rf.Kind() != reflect.Func { panic ( "expects a function" ) } vf := reflect.ValueOf(f) wrapperF := reflect.MakeFunc(rf, func (in []reflect.Value) [] reflect. Value { start := time.Now() out := vf.Call(in) end := time.Now() fmt.Printf( "calling %s took %v\n" , runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start)) return out }) return wrapperF.Interface() } func timeMe () { fmt.Println( "starting" ) time.Sleep( 1 * time.Second) fmt.Println( "ending" ) } func timeMeToo (a int ) int { fmt.Println( "starting" ) time.Sleep(time.Duration(a) * time.Second) result := a * 2 fmt.Println( "ending" ) return result } func main () { timed := MakeTimedFunction(timeMe).( func () ) timed () timedToo := MakeTimedFunction (timeMeToo). ( func ( int ) int ) fmt. Println (timedToo(2) ) } 复制代码`

输出如下：

` starting ending calling main.timeMe took 1s starting ending calling main.timeMeToo took 2s 4 复制代码`

[运行示例]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FQZ8ttFZzGx )

# 创建一个新的结构 #

Go 中，反射还可以在运行时创建一个全新的结构体，你可以通过传递一个 reflect.StructField 的 slice 给 reflect.StructOf 函数来实现。是不是听起来挺荒诞的，我们创建的一个新的类型，但是这个类型没有名字，因此也就无法将它转化为正常的变量。你可以通过它创建实例，用 Interface() 把它的值转给类型为 interface{} 的变量，但是如果要设置它的值，必须来反射来做。

` func MakeStruct (vals ... interface {}) interface {} { var sfs []reflect.StructField for k, v := range vals { t := reflect.TypeOf(v) sf := reflect.StructField{ Name: fmt.Sprintf( "F%d" , (k + 1 )), Type: t, } sfs = append (sfs, sf) } st := reflect.StructOf(sfs) so := reflect.New(st) return so.Interface() } func main () { s := MakeStruct( 0 , "" , [] int {}) // this returned a pointer to a struct with 3 fields: // an int, a string, and a slice of ints // but you can’t actually use any of these fields // directly in the code; you have to reflect them sr := reflect.ValueOf(s) // getting and setting the int field fmt.Println(sr.Elem().Field( 0 ).Interface()) sr.Elem().Field( 0 ).SetInt( 20 ) fmt.Println(sr.Elem().Field( 0 ).Interface()) // getting and setting the string field fmt.Println(sr.Elem().Field( 1 ).Interface()) sr.Elem().Field( 1 ).SetString( "reflect me" ) fmt.Println(sr.Elem().Field( 1 ).Interface()) // getting and setting the []int field fmt.Println(sr.Elem().Field( 2 ).Interface()) v := [] int { 1 , 2 , 3 } rv := reflect.ValueOf(v) sr.Elem().Field( 2 ).Set(rv) fmt.Println(sr.Elem().Field( 2 ).Interface()) } 复制代码`

输出如下：

` 0 20 reflect me [] [1 2 3] 复制代码`

[运行示例]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FlJiTP6vYYN )

# 反射的限制 #

反射有一个大的限制。虽然运行时可以通过反射创建新的函数，但无法用反射创建新的方法，这也就意味着你不能在运行时用反射实现一个接口，用反射创建的结构体使用起来很支离破碎。而且，通过反射创建的结构体，无法实现 GO 的一个特性 —— 通过匿名字段实现委托模式。

看一个通过结构体实现委托模式的例子，通常情况下，结构体的字段都会定义名称。在这例子中，我们定义了两个类型，Foo 和 Bar：

` type Foo struct { A int } func (f Foo) Double () int { return f.A * 2 } type Bar struct { Foo B int } type Doubler interface { Double() int } func DoDouble (d Doubler) { fmt.Println(d.Double()) } func main () { f := Foo{ 10 } b := Bar{Foo: f, B: 20 } DoDouble(f) // passed in an instance of Foo; it meets the interface, so no surprise here DoDouble(b) // passed in an instance of Bar; it works! } 复制代码`

[运行示例]( https://link.juejin.im?target=https%3A%2F%2Fplay.golang.org%2Fp%2FaeroNQ7bEI )

代码中显示，Bar 中的 Foo 字段并没有名称，这使它成了一个匿名或内嵌的字段。Bar 也是满足 Double 接口的，虽然只有 Foo 实现了 Double 方法，这种能力被称为委托。在编译时，Go 会自动为 Bar 生成 Foo 中的方法。这不是继承，如果你尝试给一个只接收 Foo 的函数传递 Bar，编译将不会通过。

如果你用反射去创建一个内嵌字段，并且尝试去访问它的方法，将会产生一些非常奇怪的行为。最好的方式就是，我们不要用它。关于这个问题，可以看下 github 的两个 issue， [issue/15924]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgolang%2Fgo%2Fissues%2F15924 ) 和 [issues/16522]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgolang%2Fgo%2Fissues%2F16522 ) 。不幸的是，它们还没有任何的进展。

那么，这会有什么问题呢？如果支持动态的接口，我们可以实现什么功能？如前面介绍，我们能通过 Go 的反射创建函数，实现包裹函数，通过 interface 也可以实现。在 Java 中，这叫做动态代理。当把它和注解结合，将能得到一个非常强大的能力，实现从命令式编程方式到声明式编程的切换，一个例子 [JDBI]( https://link.juejin.im?target=http%3A%2F%2Fjdbi.org%2F ) ，这个 Java 库让你可以在 DAO 层定义一个接口，它的 SQL 查询通过注解定义。所有数据操作的代码都是在运行时动态生成，就是如此的强大。

# 有什么意义 #

即使有这个限制，反射依然一个很强大的工具，每位 Go 开发者都应该掌握这项技能。但我们如何利用好它呢， [下一篇]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fcapital-one-tech%2Flearning-to-use-go-reflection-part-2-c91657395066 ) 博文再介绍，我将会通过一些库来探索反射的使用，并将利用它实现一些功能。

中文待续...