# Golang反射技术初始入门 #

反射是Go语言学习中一个比较难的点，需要好好探索一下。

## 什么反射 ##

我们知道，无论是int，float,bool等基础数据类型，亦或是array,slice,map,chan等引用类型，当使用这些类型来定义的变量，在程序编译时，编译器已经知道变量的具体类型和具体值。

但很多时候，当我们使用接口类型(interface{})定义变量时，接口类型的具体类型与具体值，需要在程序运行时才能确定，且可以动态变化，因此需要一种技术来检测变量在程序运行中的具体类型和值。

Go反射技术就是这样一种用来检查未知类型和值的机制与方法。

## 为什么需要反射 ##

当我们需要实现一个通用的函数时，比如实现一个类似fmt.Sprint这样的打印函数时，可以根据不同的数据类型，返回出不同格式的数据，我们也实现一个类似fmt.Sprint()函数但只可以打印一个参数的Sprint函数，代码如下：

` type stringer interface { String() string } func S print (x interface{}) string { switch x := x.( type ) { case stringer: return x.String() case string: return x case int: return strconv.Itoa(x) // 还有int16, uint32,或者更多我们自定义的未知类型. case bool: if x { return "true" } return "false" default: // 默认返回值 return "???" } } 复制代码`

在上面打印函数中，我们只是判断了几种基础类型，但这是不够，还有许多类型没有判断，虽然我们可以在上面的switch结构中继续增加分支判断，但在实际的程序中，还更多自定义的未知类型，因此需要使用反射技术来实现。

## reflect.Type和reflect.Value ##

Go语言反射技术是由reflect包来实现，这个包主要定义了reflect.Type和reflect.Value两个重要的类型。

### reflect.Type ###

reflect.Type是一个接口，代表一个的具体类型，使用reflect.TypeOf()函数，可以返回reflect.Type的实现，reflect.TypeOf()方法可以接收任何类型的参数，如下：

` t := reflect.TypeOf(100) fmt.Println(t)//int 复制代码`

### reflect.Value ###

reflect.Value是一个reflect包中定义的结构体，代表一个类型的具体值，使用reflect.ValueOf()函数可以返回一个reflect.Value值，reflect.ValueOf()可以接收任意类型的参数。

使用reflect.Value中的Type()方法，可以返回对应的reflect.Type。

` v := reflect.ValueOf(10) fmt.Println(v) t := v.Type() fmt.Println(t) 复制代码`

因此，我们可以使用反射来修改上面的Spring函数。

` func S print (vv interface{}) string { v := reflect.ValueOf(vv) switch v. Kind () { case reflect.Invalid: return "invalid" case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64: return strconv.FormatInt(v.Int(), 10) case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: return strconv.FormatUint(v.Uint(), 10) case reflect.Bool: return strconv.FormatBool(v.Bool()) case reflect.String: return strconv.Quote(v.String()) case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map: return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16) default: return v.Type().String() + " value" } } 复制代码`

## 总结 ##

Golang提供的反射reflect包，是用于检测类型、修改类型值、调用类型方法和其他操作的强大技术，在其他的库或框架中都有使用，但还是应该慎用。

除了我们经常使用的fmt包是应用反射实现的之外，encoding/json、encoding/xml等包也是如此。

其实在一些Web框架中，将http请求参数绑定到模型中，在ORM框架中，将数据表查询结果绑定到模型中，应用的都是反射技术。