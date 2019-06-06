# Kotlin艺术探索之类和接口 #

### 类 ###

#### 1.基本用法 ####

Kotlin的类的普通写法和Java一样

` class Invoice { ... } 复制代码`

如果类里面内容是空的，可以把括号省略

` class Empty 复制代码`

#### 2.类的构造方法 ####

构造方法的参数可以直接写在类名之后

` class InitOrderDemo(name: String) { var firstProperty = "First property: $name ".also (::println) init { println( "First initializer block that prints ${name} " ) } } 复制代码`

可以在init方法里面初始化数据

当然也可以在构造器方法constructor方法里初始化

` class Constructors { init { println( "Init block" ) } constructor(i: Int) { println( "Constructor" ) } } 复制代码`

它们的不同在于，init方法不可以传参数，而constructor可以设置参数。

> 
> 
> 
> 要注意的是，init方法一定是在constructor方法前执行。
> 
> 

#### 3.类的继承 ####

如果某个类需要被继承，那么需要使用open关键字，表示该类可以被继承

` open class Base(p: Int) class Derived(p: Int) : Base(p) 复制代码`

#### 类方法的重写 ####

被继承的父类中方法如果需要被重写，也需要加上open关键字，并且子类重写的方法中要加上override关键字

` open class Base { open fun v () { } fun nv () { } } class Derived() : Base () { override fun v () { } } 复制代码`

#### 4.类属性的重写 ####

类的属性重写和方法相似，属性的赋值可以用get方法

` open class Foo { open val x: Int get() = 1 } class Bar1 : Foo () { override val x: Int get() = 2 } 复制代码`

当然也可以在子类的构造方法中进行重写，这也需要用override关键字

` class Bar2(override val x: Int) : Foo() 复制代码`

#### 5.调用父类的属性和方法 ####

子类可以通过super关键字调用父类的属性和方法

` open class Foo { open fun f () { println( "Foo.f()" ) } open val x: Int get() = 1 } class Bar : Foo () { override fun f () { super.f() println( "Bar.f()" ) } override val x: Int get() = super.x + 1 } 复制代码`

这是针对外部类，如果是内部类调用父类属性和方法除了super还需要借助 @ 符号，类似这样super@Outer

` class Bar : Foo () { override fun f () { /* ... */ } override val x: Int get() = 0 inner class Baz { fun g () { super@Bar.f() // Calls Foo 's implementation of f() println(super@Bar.x) // Uses Foo' s implementation of x 's getter } } } 复制代码`

#### 6.抽象类的表示 ####

抽象类的表示要用abstract关键字

` abstract class abstractClass { abstract fun f() } 复制代码`

抽象类本身和抽象方法不需要加open来表明可以被继承，因为抽象类的本意就是用来被继承的

### 接口 ###

#### 1.接口的基本实现 ####

Kotlin的接口和Java8的特性差不多

它们都可以书写抽象方法的声明和方法的具体实现

定义一个接口 MyInterface

` interface MyInterface { val prop: Int // abstract val propertyWithImplementation: String get() = "foo" fun foo () { print (prop) } } 复制代码`

定义一个类实现MyInterface接口，由于在接口中并没有初始化变量prop，所以实现类中必须要重写该变量

` class Child : MyInterface { override val prop: Int = 29 } 复制代码`

那么在接口中如何初始化变量呢？

可以使用get方法来初始化name这个变量

` interface Named { val name: String } //继承Named接口 interface Person : Named { val firstName: String val lastName: String override val name: String get() = " $firstName $lastName " } 复制代码`

#### 2.解决接口方法冲突 ####

定义两个接口A，B

` interface A { fun foo () { println( "A" ) } fun bar() } interface B { fun foo () { println( "B" ) } fun bar () { println( "B of bar" ) } } 复制代码`

A，B这两个接口有重合的方法，并且类D实现这两个接口，那么如果区分调用这两个接口的方法呢？

` class D: A,B{ override fun foo () { super<A>.foo() super<B>.foo() } override fun bar () { super<B>.bar() } } 复制代码`

接口冲突的解决方法很简单，通过<接口>方式调用需要的接口方法

> 
> 
> 
> 注意：接口方法冲突的解决有两种要求：
> 1.接口方法的签名（方法名和参数）一致且返回值类型相同
> 2.接口方法有默认的实现
> 
>