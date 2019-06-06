# 方法 - Go 语言学习笔记 #

### 什么是方法？ ###

Go 语言中同时有函数和方法。方法就是一个包含了接收者的函数，接收者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。

在 Go 中，（接收者）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收者来建立。 方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

### 声明方法 ###

语法格式如下:

` func (variable_receiver_name receiver_type) function_name([parameter list]) [return_type]{ /* 函数体*/ } 复制代码`

* variable_receiver_name 接收者必须有一个显式的名字，这个名字必须在方法中被使用。官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名。例如，Socket 类型的接收器变量应该命名为 s，Connector 类型的接收器变量应该命名为 c 等。
* receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。
* 方法名、参数列表、返回参数：格式与函数定义一致。

` package main import ( "fmt" ) // 定义结构体 type Circle struct { radius float 64 } func main () { var c1 Circle c1.radius = 10.00 fmt.Println( "圆的面积 = " , c1.getArea()) } // 该 method 属于 Circle 类型对象中的方法 func (c Circle) getArea() float 64 { // c.radius 即为 Circle 类型对象中的属性 return 3.14 * c.radius * c.radius } // 执行结果为：的面积 = 314 复制代码`

### 接收者 ###

接收者有两种，一种是值接收者，一种是指针接收者。顾名思义，值接收者，是接收者的类型是一个值，是一个副本，方法内部无法对其真正的接收者做更改；指针接收者，接收者的类型是一个指针，是接收者的引用，对这个引用的修改之间影响真正的接收者。像上面一样定义方法，将 user 改成 *user 就是指针接收者。

### 调用方法 ###

struct的方法调用

方法调用相当于普通函数调用的语法糖。Value方法的调用m.Value()等价于func Value(m M) 即把对象实例m作为函数调用的第一个实参压栈，这时m称为receiver。通过实例或实例的指针其实都可以调用所有方法，区别是复制给函数的receiver不同。

如下，通过实例m调用Value时，以及通过指针p调用Value时，receiver是m和*p，即复制的是m实例本身。因此receiver是m实例的副本，他们地址不同。通过实例m调用Pointer时，以及通过指针p调用Pointer时，复制的是都是&m和p，即复制的都是指向m的指针，返回的都是m实例的地址。

` type M struct { a int } func (m M) Value() string { return fmt.S printf ( "Value: %p\n" , &m)} func (m *M) Pointer() string { return fmt.S printf ( "Pointer: %p\n" , m)} var m M p := &m // p is address of m 0x2101ef018 m.Value() // value(m) return 0x2101ef028 m.Pointer() // value(&m) return 0x2101ef018 p.Value() // value(*p) return 0x2101ef030 p.Pointer() // value(p) return 0x2101ef018 复制代码`

### 方法用法 ###

#### 1. 基于指针对象的方法 ####

当接受者变量本身比较大时，可以用其指针而不是对象来声明方法，这样可以节省内存空间的占用。

` package main import ( "fmt" "math" ) // 方法声明 type Point struct { X,Y float 64 } func (p *Point) Distance(q *Point) float 64 { return math.Hypot(q.X - p.X, q.Y - p.Y) } func main () { p := &Point{3,5} fmt.Println(p.Distance(&Point{5, 6})) } // 执行结果为：2.23606797749979 复制代码`

#### 2. 将nil作为接收器 ####

` package main import "fmt" // 将nil作为接收器 type IntNode struct { Value int Next *IntNode } func (node *IntNode) Sum() int { if node == nil { return 0 } return node.Value + node.Next.Sum() } func main () { node1 := IntNode{30, nil} node2 := IntNode{12, nil} node3 := IntNode{43, nil} node1.Next = &node2 node2.Next = &node3 fmt.Println(node1.Sum()) fmt.Println(node2.Sum()) node := &IntNode{3, nil} node = nil fmt.Println(node.Sum()) } // 执行结果为： // 85 // 55 // 0 复制代码`

#### 3. 通过嵌入结构体拓展类型 ####

示例：

` package main import ( "fmt" ) // 基础颜色 type BasicColor struct { // 红、绿、蓝三种颜色分量 R, G, B float 32 } // 完整颜色定义 type Color struct { // 将基本颜色作为成员 BasicColor // 透明度 Alpha float 32 } func main () { // 设置基本颜色分量 var c Color c.R = 1 c.G = 1 c.B = 0 // 设置透明度 c.Alpha = 1 // 显示整个结构体内容 fmt.Printf( "%+v" , c) } 复制代码`

代码输出如下： {Basic:{R:1 G:1 B:0} Alpha:1}

Go语言的结构体内嵌特性:

* 内嵌的结构体可以直接访问其成员变量
嵌入结构体的成员，可以通过外部结构体的实例直接访问。如果结构体有多层嵌入结构体，结构体实例访问任意一级的嵌入结构体成员时都只用给出字段名，而无须像传统结构体字段一样，通过一层层的结构体字段访问到最终的字段。例如，ins.a.b.c的访问可以简化为ins.c。
* 内嵌结构体的字段名是它的类型名
内嵌结构体字段仍然可以使用详细的字段进行一层层访问，内嵌结构体的字段名就是它的类型名，代码如下：
` var c Color c.BasicColor.R = 1 c.BasicColor.G = 1 c.BasicColor.B = 0 复制代码`

一个结构体只能嵌入一个同类型的成员，无须担心结构体重名和错误赋值的情况，编译器在发现可能的赋值歧义时会报错。