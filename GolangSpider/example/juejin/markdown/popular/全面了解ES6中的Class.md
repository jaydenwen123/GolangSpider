# 全面了解ES6中的Class #

## 写在前面 ##

在ES5及之前的版本中，并没有类的概念，那时候创建类的方法是大家都很熟悉，如下是一个最简单的例子：

` // 一般我们约定以大写字母开头来表示一个构造函数 function Person (name) { this.name = name; } // 在构造函数的原型链上定义不同的方法 Person.prototype.getName = function () { return this.name; } // 实例化一个对象，拥有构造函数原型链上的方法 var people = new Person( 'lilei' ); people.getName(); // lilei 复制代码`

但是这种写法和传统的面向对象语言（例如java）有较大区别，不太好理解。所以在ES6中引入了class关键字，它是一个让对象原型的写法更像面向对象编程语法的一个语法糖。

在node开发以及一些组件的开发中，会经常用到class语法，本文将先介绍class的一些基础用法，然后深入剖析class的继承，以及原型链相关的知识。总结一下自己学习class的过程，也希望通过一些示例能帮助到部分同学，由浅及深，逐步理解class的用法以及原理。

## 基础用法 ##

### 写法示例 ###

我们先来看一个class的常规写法：

` class Person { _myName = 'initial-name' ; constructor (name = '' ) { this._myName = name; } getName () { return this._myName; } static toUpperCaseName () { return this.name.toUpperCase(); } get name () { return this._myName; } set name (name) { this._myName = name; } } let people = new Person( 'lilei' ); console.log(people.name); // lilei console.log(people.getName()); // lilei console.log(Person.toUpperCaseName()); // PERSON people.name = 'lilei2' ; console.log(people.name); // lilei2 复制代码`

下面我们把这个示例拆分开来，逐一了解class的基础用法。

### 构造方法（constructor） ###

1、class在定义时， **必须定义一个构造方法** （constructor），如果不写，Javascript引擎会自动添加一个空的constructor方法。

` class Person {} // 其实是 class Person { constructor () {} } 复制代码`

2、constructor方法是class的构造方法，在class实例化对象时（即new一个对象时）会 **自动调用构造方法** ，并且默认返回实例对象（this），当然你完全可以返回另外一个对象。

` class Person { constructor (name) { console.log(name); } } new Person( 'lilei' ); // lilei 复制代码`

### 静态方法 ###

1、定义class的非静态方法不用加上function关键字，也不需要用逗号隔开，用了反而会报错。而这些普通方法在实例化时，就会被实例继承。 2、定义class的静态方法，则是在方法名前加static关键字，这样定义的方法，就 **不会被实例继承** ，但是它 **可以直接被类调用** ，如下：

` class Person { func1 () { console.log( 'func1' ); } static func2 () { console.log( 'func2' ); } } let people = new Person(); people.func1(); // func1 people.func2(); // people.func2 is not a function Person.func1(); // Person.func1 is not a function Person.func2(); // func2 复制代码`

3、静态方法可以与非静态方法重名；静态方法中，this指向的是类，不是实例。

` class Person { static func1 () { console.log(this.prototype.func2()); console.log(this.func2()); } static func2 () { console.log( 'static' ); } func2 () { console.log( 'function' ); } } Person.func1(); // function // static 复制代码`

4、非静态方法中，不能直接使用this关键字来访问静态方法，而是要用类名来调用。 **如果静态方法包含this关键字，这个this指的是类，而不是实例。**

` class Person { constructor () { console.log(Person.func1()); console.log(this.constructor.func1()); } static func1 () { console.log( 'static' ); } } 复制代码`

5、另外提一点，class的方法名可以用变量来命名，如下：

` let funcName = 'getName' ; class Person { [funcName] () {} } 复制代码`

6、类的所有方法都定义在类的prototype属性上面

### 属性的定义方法 ###

1、在ES6中，属性定义的常规写法是定义在类的constructor方法里的this上，示例见本文基础用法中第一个代码段，这里就不再写了。 2、在ES7中定义了一种新的属性定义规范，可以将属性定义在类的最顶层，这样使代码看上去更加整洁，很清晰的看到类中定义的属性。

` class Person { _name = '' ; constructor (name) { this._name = name; } } 复制代码`

### get与set ###

1、在class里还可以给属性添加get和set关键字，拦截属性的存取行为。写法见基础用法中第一个代码段。 2、 **当一个属性只有getter没有setter的时候，我们是无法进行赋值操作的** ，第一次初始化也不行。如果把变量定义在类的外面，就可以只使用getter不使用setter。

` let data = {}; class Person{ constructor () { this.name = 'lilei' ; this.age = 20; } get age (){ return this._age; } // 如果没有这个 set ，则会报错：Cannot set property width of #<GetSet> which has only a getter set age(age){ this._age = age; } get data (){ return data; } } let people = new Person(); console.log(people.age); people.age = 30; console.log(people.age); console.log(people._age); console.log(people.data); 复制代码`

3、不能在set方法中设置自己的值，不然会陷入无限递归，导致栈溢出。

` set age (v) { this.age = v; } // Uncaught RangeError: Maximum call stack size exceeded 复制代码`

### class表达式 ###

1、类可以用表达式的方式来定义，需要注意的是P只能在class内部使用，用于指代当前的类；Person只能在外部使用。

` const Person = class P { getName () { return P.name // p } } Person.name // P 复制代码`

2、和let一样，class不存在变量提升，即不能先使用后声明。 3、class拥有name属性，返回的是 **class关键字后面的名字** ，见上面的例子。

## 深入了解 ##

### 属性及方法定义的位置 ###

实例的属性及方法除非显式定义在其本身（即定义在this对象上），否则都是定义在原型上（即定义在class上）。ES7中约定的新的定义属性的规范也是定义在实例上的。

` class P { a = 1; constructor(b) { this.b = b; } c () { console.log(this.a, this.b); } } var p1 = new P(2); p1.hasOwnProperty( 'a' ) // true p1.hasOwnProperty( 'b' ) // true p1.hasOwnProperty( 'c' ) // false p1.__proto__.hasOwnProperty( 'c' ) // true 复制代码`

### class语法存在屏蔽方法的问题 ###

我们来看这样一段代码：

` class P { constructor(id) { this.id = id; } id () { console.log( "Id: " + this.id ); } } var p1 = new P( "p1" ); p1.id(); // TypeError -- p1.id 现在是字符串 "p1" 复制代码`

在构造函数中定义的属性会 **屏蔽** 同名的方法。注意这里是屏蔽，不是覆盖、前面说到通过this定义的属性是定义在实例上的，而方法是定义在类上的，所以实例在读取id的时候会率先读取到定义在实例上的属性id，也就读不到方法id了。

## 继承 ##

class引入目的就是让类定义起来更直观接单，通过extends关键字实现继承的写法也更便于理解。其中class A extends B 中A称为派生类，派生类是指继承自其它类的新类。

### super ###

1、派生类没有自己的this对象，而是继承父类的this对象，所以需要在构造函数中先调用super()初始化this对象，之后才能在构造函数中访问this。ES6 要求，子类的构造函数必须执行一次super函数。 2、super既可以当函数使用，也可以当对象使用。 1）当函数使用时，它代表的是父类的构造函数，并且只能在子类的构造函数中使用，不然就会报错。 2）当对象使用时，如果是在普通方法中，指向对象的原型对象；如果是在静态方法中，指向的是父类。

` class Person { static getMethod () { console.log( 'static' ); } getMethod () { console.log( 'instance' ); } } class Man extends Person { static getMethod () { super.getMethod();   } getMethod(msg) { super.getMethod(); } } Man.getMethod(); // static var lilei = new Man(); lilei.getMethod(); // instance 复制代码`

3、super指向的是父类的原型对象，所以定义在父类实例上的方法或属性，是无法通过super调用的。

` class A { constructor () { this.p = 2; } } class B extends A { get m () { return super.p; } } let b = new B(); b.m; // undefined 复制代码`

4、通过super调用父类的方法时，super会绑定子类的this

` class A { constructor () { this.color = 'red' ; } draw () { console.log(this.color); } } class B extends A { constructor () { super(); this.color = 'yellow' ; } draw () { super.draw(); // 这里是会绑定B的this，也就是super.draw.call(this) } } let b = new B(); b.draw() // yellow 复制代码`

### 派生自表达式的类 ###

因为class继承的逻辑是先新建一个父类的this对象，然后在子类的构造函数中对this进行加工，所以父类的所有行为都可以被继承。也就是 **class可以继承原生构造函数来定义子类** 。

` class MyArray extends Array {} let arr = new MyArray(1,2); console.log(arr); // [1, 2] 复制代码`

## 参考 ##

[ECMAScript 6 入门]( https://link.juejin.im?target=http%3A%2F%2Fes6.ruanyifeng.com%2F )