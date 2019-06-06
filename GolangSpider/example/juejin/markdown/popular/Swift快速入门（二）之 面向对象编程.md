# Swift快速入门（二）之 面向对象编程 #

# Swift快速入门（二）之 面向对象编程 #

[Swift快速入门（一）之基础语法]( https://juejin.im/post/5ced5b8b5188252db706f5cc/ )

本文为第二篇《Swift面向对象编程》主要涵盖以下内容

* 函数
* 闭包
* 枚举
* 结构体和类
* 属性
* 初始化

## 函数 ##

函数（function）相当于java中的方法。函数的的声明使用 **func** 关键字。

#### 函数声明示例 ####

` var title = "Hello,函数" //声明函数 func sayHell () { print (title) } //函数的调用 sayHell() 复制代码`

#### 声明带参数的函数 ####

` func sayHello (name:String) { print ( "Hello \(name) " ) } sayHello(name: "Taoto" ) 复制代码`

#### 函数-外部参数 ####

**外部参数** 让函数体外可见的参数名不同于内部的参数名。也就是说调用函数的时候用一个参数名字，在函数体内用另一个名字。

` //外部调用使用参数名 为 to,内部使用参数名为 name func sayHello (to name:String) { print ( "外部参数的使用：Hello \(name) " ) } sayHello(to : "jake" ) ```swift #### 函数-可变参数 1. 可变参数可接受 0 个或多个值 2. 函数只能一个可变参数，且可变参数需作为最后一个参数 ```swift func sayHelloVariadic (to names : String...) { print ( "Hello everyons: \(names) " ) } sayHelloVariadic(to: "jake" , "jerry" , "kobe" , "jordan" ) sayHelloVariadic() //可变参数可接受0个值 复制代码`

#### 函数-默认参数 ####

* 默认参数可接受默认参数值；
* 默认参数值应该放在函数参数列表的末尾；
* 如果形参有默认值，则调用时可以省略实参；
` func sayHelloDefault (name:String = "sir" ) { print ( "Hello \(name) " ) } func hi (age:Int32, name:String = "sir" ) { print ( "hi \(name) ,your age is \(age) " ) } sayHelloDefault() //省略实参 sayHelloDefault(name: "god" ) hi(age: 12 ) hi(age: 22 ,name: "mimo" ) 复制代码`

#### 函数-inout参数 ####

* inout参数能让函数影响函数体以外的变量
* 传递inout参数的变量名前需在前面加一个&
` var erro = "401" func appendErroCode (erroStr : inout String) { erroStr += " erro code" } appendErroCode(erroStr: &erro) //传递inout参数的变量名前需在前面加一个& print (erro) 复制代码`

#### 函数-返回值 ####

返回值语法： -> returnType

` //声明一个加法函数，接收两个整数的参数，并返回它们的和 func addDescriptionFor (a:Int,b:Int) -> Int { return a+b; } var sumNum = addDescriptionFor(a: 10 , b: 18 ) print ( "求和结果 \(sumNum) " ) 复制代码`

#### 嵌套函数和作用域 ####

嵌套函数在另一个函数内部声明并实现，并可以访问该函数内部的成员变量

` //示例：通过底和高计算面积 func areaOfTraiangleWith (base:Double, height:Double) -> Double { let numberator = base * height func divider () -> Double { return numberator/ 2 ; } return divider(); } print (areaOfTraiangleWith(base: 9 , height: 10 )) 复制代码`

#### 函数-多个返回值 ####

函数可以返回不止一个值，可以是元祖，或复杂对象

` //使用元祖作为返回值 //示例代码，传入两个参数，分别将这两个参数乘2，再组装成一个元祖返回 func copyTwo (d1:Double ,d2:Double) ->(n1: Double ,n2 : Double ) { return (d1* 2 ,d2* 2 ) } var result = copyTwo(d1: 3 , d2: 4 ) print (result) 复制代码`

#### 函数-可空的返回值类型 ####

` func funcNil (targetStr:String?) -> String ? { return targetStr } print ( "可空的返回值类型 \(funcNil(targetStr: nil ) )" ) 复制代码`

#### 提前退出函数 ####

**guard** 语句 中文 “卫语句”，可提前退出函数，防止不当条件下运行的方式

` func sayHello (name : String, age:Int) { if age< 0 || age> 100 { return } print ( "hello \(name) ,age is \(age) " ) } sayHello(name: "leo" , age: 18 ) 复制代码`

## 闭包 closure ##

闭包是完成特定任务的功能组，可以理解为无名字的函数。

* swift中闭包可以作为参数，也可以作为返回值
* 闭包和函数能捕获在其作用域中定义的变量
* 闭包是引用类型

==ps==:闭包和java中lambda表达式匿名函数类似，甚至可以理解为同一个东西。

#### 闭包语法表达式 ####

` {(parameter) -> return type in //表达式 } 复制代码`

通过一个对整数数组进行排序对闭包特性进行学习

` var ages = [ 1 , 99 , 78 , 24 , 88 , 61 , 56 , 12 , 44 , 39 ] func sortAscending ( _ i:Int , _ j:Int) -> Bool { return i < j } var agesSorted = ages.sorted(by: sortAscending) //print(agesSorted) //闭包表达式对排序进行简化 agesSorted = ages.sorted(by: { (i: Int , j: Int ) -> Bool in return i<j }) //利用类型推断的闭包语法 agesSorted = ages.sorted(by:{ i,j in i<j }) //利用参数的快捷语法 //使用$0表示第一个参数的值，$1表示第二个参数的值 //现在内联闭包表达式利用了快捷参数语法，就不需要像之前申明i和j那样显示声明参数了 agesSorted = ages.sorted(by: {$ 0 < $ 1 }) 复制代码`

#### 函数作为返回值 ####

` //声明一个getSum函数，其返回值为 为一个 //（接收两个Int参数，返回一个Int返回值的函数） func getSum () -> ( Int , Int ) -> Int { func sumTotal (a:Int , b:Int) -> Int { return a + b; } return sumTotal } var sumF = getSum() print (sumF( 8 , 9 )) 复制代码`

#### 函数作为参数 ####

` //contion为 接受一个Int参数返回Bool类型 的闭包 func isBiggerThanZero (num:Int, condition : (Int) -> Bool ) -> String { if (condition(num)) { return "大于0的整数！！！" } else { return "小于0的整数" } } var result = isBiggerThanZero(num: 9 , condition: { (i : Int ) -> Bool in return i> 0 } ) print (result) //利用类型推断进行简化 闭包 result = isBiggerThanZero(num: - 9 , condition: { i in i> 0 }) print (result) //利用快捷表达式再次进行简化 result = isBiggerThanZero(num: 10 , condition: {$ 0 > 0 }) print (result) 复制代码`

### 函数式编程--高阶函数 ###

#### map 变换数组的内容 ####

把数组的内容从一个值变成另外一个值，并把这些新值放进一个新数组返回

` let projectPopulations = [ 1024 , 2048 , 4096 ] let projectedPopulations = projectPopulations. map { (population : Int ) -> Int in return population * 2 } print (projectedPopulations) //闭包的简写 let projectedPopulations3 = projectPopulations. map { i in i* 3 } //将整数数组变为每个元素依次乘以4 的数组 print ( "闭包的简写 \(projectedPopulations3) " ) let projectedPopulations4 = projectPopulations. map { $ 0 * 4 } //将整数数组变为Bool类型数组 print ( "4闭包的简写 \(projectedPopulations4) " ) let projectedPopulations2 = projectPopulations. map { (i : Int ) -> Bool in return i> 10 } print (projectedPopulations2) 复制代码`

#### filter 过滤数组 ####

结果数组会包含原数组满足条件的值

` var bigProjectedPopulations = projectPopulations. filter { (population: Int ) -> Bool in return population> 2000 //当population>2000时 返回true } //print(bigProjectedPopulations) bigProjectedPopulations = projectPopulations. filter ({ (i: Int ) -> Bool in return i > 10 }) print ( "bigProjectedPopulations2 : \(bigProjectedPopulations) " ) //简写,利用条件推断 bigProjectedPopulations = projectPopulations. filter ({ i in i > 0 }) //print("bigProjectedPopulations3 :\(bigProjectedPopulations)") //利用参数的快捷语法 bigProjectedPopulations = projectPopulations. filter ({ $ 0 > 0 }) print ( "bigProjectedPopulations4 : \(bigProjectedPopulations) " ) //尾部闭包 bigProjectedPopulations = projectPopulations. filter { $ 0 > 0 } print ( "bigProjectedPopulations5 尾部闭包: \(bigProjectedPopulations) " ) 复制代码`

#### reduce 求和 ####

对数组中各个元素遍历求和

` //let projectPopulations = [1024,2048,4096] var numberSum = projectPopulations. reduce ( 0 , { x, y in x + y }) print (numberSum) numberSum = projectPopulations. reduce ( 0 , {$ 0 + $ 1 }) print ( "j \(numberSum) " ) 复制代码`

## 枚举 ##

#### 声明枚举 ####

` enum TextAlignment { case left case right case center } 复制代码`

#### 创建一个枚举实例 ####

` //第一次创建枚举变量时，必须指定枚举名和值 var alignment = TextAlignment.center alignment = . left print (alignment) 复制代码`

#### 比较枚举时 ####

` if (alignment == . left ) { print ( "The alignment is \(alignment) " ) } 复制代码`

#### 枚举在switch的使用 ####

==注意==：枚举作为switc语句case时，建议不使用default，编译器会检测case是否全部覆盖，若覆盖不全则编译报错

` switch alignment{ case. left : print ( "left align" ) case.center: print ( "center align" ) case. right : print ( "right align" ) } 复制代码`

#### 原始枚举值 rawValue ####

目前支持：Int、Float、Double、String

` //若原始值未赋值，则从0..1依次递增 enum TextAlignmentInt : Int { case left = 10 case center = 20 case right = 30 } print ( TextAlignmentInt. left.rawValue) 复制代码`

#### 原始值转回枚举 ####

` let myRawValue = 20 if let myAlignment = TextAlignmentInt (rawValue: myRawValue) { print ( "successful converted \(myRawValue) into \(myAlignment) " ) } else { print ( "failure converted \(myRawValue) into TextAlignmentInt" ) } 复制代码`

#### 创建带字符串原始值的枚举 ####

` enum ProgramingLanguage : String { case swift = "swift" case java = "java" case other //若未赋值则rawValue = others } var language = ProgramingLanguage.swift print (language.rawValue) language = .other print (language.rawValue) 复制代码`

## 结构体 和 类 ##

结构体（strut）是把相关数据块组合在一起的一种类型；结构体是值类型。

#### 声明结构体 strut ####

` struct Town { } 复制代码`

#### 结构体内添加属性 ####

` struct Town { //存储属性，可以有默认值 var name = "温泉小镇" //名称 var population = 2300 //人口 } 复制代码`

#### 结构体实例方法 ####

` struct Town { var name = "温泉小镇" //名称 var population = 2300 //人口 func sayHello () { print ( "Hello! My name is \(name) ,population is \(population) " ) } } 复制代码`

#### 结构体 mutating 方法 ####

结构体是值类型，需要在修改实例属性的方法前加上 **mutating** 关键字。

` struct Town { //存储属性，可以有默认值 var name = "温泉小镇" //名称 var population = 2300 //人口 //由于stuct是值类型，修改其属性内容需使用 mutating mutating func addPopulation (amount : Int) { population += amount } 复制代码`

#### 结构体 类型方法（static方法） ####

类型方法用static关键字进行修饰，可通过类型本身进行调用(即java中的静态方法)

` struct Town { //stuct的 类型方法static修饰 static func numberOfSteet () -> Int { return 10 } } //类型方法的调用 Town.numberOfSteet() 复制代码`

#### 创建结构体实例 ####

` var myTown = Town () print ( "Town name is \(myTown.name) ,population is \(myTown.population) " ) myTown.sayHello() myTown.addPopulation(amount: 1000 ) myTown.sayHello() print ( Town.numberOfSteet()) 复制代码`

### 类 class ###

类和结构体类似，但类是引用类型。

#### 类的声明 ####

示例：类的声明，包含属性、对象方法

* 父类中普通对象方法可以被子类重写override
* 父类中使用final修饰的对象方法子类不能重写override
` class Monster { var town: Town ? var name = "Monster" func terrorizeTown () { if town != nil { print ( " \(name) terrorizing a town: \(town?.name) !" ) } else { print ( " \(name) hasn't found a town to terrorize yet..." ) } } 复制代码`

#### 类的 类方法 ####

* 类方法可以用class或static修饰
* class修饰的类方法子类可以override
* static修饰的类方法子类不可重写（和java一致）

#### 类的继承 子类方法override ####

子类继承父类的属性和方法。

` import Foundation //声明类 怪物 class Monster { var town: Town ? var name = "Monster" func terrorizeTown () { if town != nil { print ( " \(name) terrorizing a town: \(town?.name) !" ) } else { print ( " \(name) hasn't found a town to terrorize yet..." ) } } //final方法 ，子类中无法ovveride final func hello (name :String) { print ( "hello \(name) " ) } //类的 类型方法class修饰,即java中的静态方法，子类可以继承 class func numberOfMonster () -> Int { return 10 } //类的 类型方法class修饰,即java中的静态方法，子类可以继承 //当类型方法 前添加 final或static修饰时，则子类不可override final class func makeNoise () -> String { return "Ping Pong" } } //声明类 僵尸继承自 Monster class Zombie : Monster { var walkWithLimp = true //复写父类方法 override func terrorizeTown () { town?.addPopulation(amount: - 10 ) //人口减10 super.terrorizeTown() } //在swift中子类可以override 父类的 类型方法 override class func numberOfMonster () -> Int { return 1 } } 复制代码`

## 属性 class ##

属性将值跟特定的类、结构或枚举关联。 属性可分为存储属性和计算属性:

* 存储属性：存储常量或变量作为实例的一部分，用于类和结构体
* 计算属性：计算一个值，用于类和结构体及枚举

#### 声明存储属性 ####

` class People { //存储属性可以是常量或变量，可以有默认初始值 let no = 0 var age: Int var name: String } 复制代码`

#### 声明惰性存储属性 ####

惰性存储属性当第一次访问的时候被调用，且只会被计算一次

` //声明懒加载属性 //当第一次访问的时候被调用 //只会被计算一次 lazy var personType: Classification = { switch age{ case 0... 12 : return Classification. Child case 13... 18 : return Classification. Teenagers case let age where (age> 18 ): return Classification. Adult default : return Classification. Ohter } }() 复制代码`

#### 嵌套类型 ####

嵌套类型是定义在一个类型内部的类型，类似为Java中的内部类

` // 属性的使用 class People { //存储属性可以是常量或变量，可以有默认初始值 let no = 0 var age: Int var name: String init (age: Int ,name: String ) { self.age = age self.name = name } //内部类型 enum Classification { case Child case Teenagers case Adult case Ohter } } 复制代码`

#### 只读属性 ####

只读属性不允许进行赋值操作

` //声明只读属性 var victimPool: Int { get { return 9 } } 复制代码`

#### 声明计算属性 ####

计算属性不能直接设置变量值，只能根据读取方法中定义的计算获取一个返回值，每次获取都是动态实时的。

` //声明计算属性:不能直接设置变量值，只能根据读取方法中定义的计算获取一个返回值 var personTypeDynamic: Classification { get { switch age{ case 0... 12 : return Classification. Child case 13... 18 : return Classification. Teenagers case let age where (age> 18 ): return Classification. Adult default : return Classification. Ohter } } } 复制代码`

## 类型属性 ##

类型属性 类似java中的静态属性，但区别在于：

* strut中 类型属性使用static关键进行声明
* class中，类型属性可使用static和class进行声明；static声明的类型属性子类不可以override，而class声明的类型属性可以被子类override
` class People { //static声明的属性变量 static let height = 10 static var height2 = 11 class var width : Int { return 12 } } class Man : People { //class声明的类型属性可以被override override class var width : Int { return 33 } } 复制代码`

## 初始化 initializer ##

Swift中初始化方法相当于java中的构造函数 1.在类和结构体中编译器会默认提供空的初始化方法 2.当用户新增带参数初始化方法后，编译器则不提供默认午餐初始化方法，如需使用要自行实现

` struct Zoo { var name = "" var area: Double = 100 //初始化方法、结构体 init (name : String ) { //支持初始化方法的重载 self. init (name: name, area: 90 ) } init (name : String ,area : Double ) { self.area = area // self.init(name: name) self.name = name } } 复制代码`

#### 委托初始化 ####

对类型其他初始化方法进行调用，相当于java中的重载构造函数

` //初始化方法、结构体 init (name : String ) { //支持初始化方法的重载，对其他初始化方法进行调用 self. init (name: name, area: 90 ) } init (name : String ,area : Double ) { self.area = area // self.init(name: name) self.name = name } 复制代码`

### 类初始化 ###

类初始化通用语法和值类型类似，但由于类支持继承，会在值类型初始化的基础上新增了 指定( **designated** )初始化方法和便捷( **convenience** )初始化，类初始化方法一定是二者之一。

* 指定初始化需要确保初始化完成前所有的属性都有值，以便实例可用。
* 便捷初始化方法是指定初始化方法的补充，通过调用所在类的指定初始化方法来实现，主要作用是为某种特殊目的创建实例。

` class Animal { var name: String var age: Int //指定初始化方法，不加任何修饰符 //1.需要确保初始化完成前所有的属性都有值，以便实例可用 init (name: String ,age: Int ) { self.name = name self.age = age } //初始化、构造函数 //△convenience 便捷初始化 //是指定初始化方法的补充，通过调用所在类的指定初始化方法来实现 convenience init (name: String ) { self. init (name: name, age: 18 ) } } 复制代码`

#### 初始化方法自动继承 ####

只有满足以下两种场景，子类会继承父类的初始化方法

* 如果子类没有定义任何指定初始化方法，就会继承父类的指定初始化方法。
* 如果子类实现了父类的所有指定初始化方法（无论通过显示还是隐式继承），就会继承父类的所有便捷初始化方法。

#### 指定初始化方法创建实例 ####

` class Panda : Animal { //子类没有定义任何指定初始化方法，就会继承父类的指定初始化方法 } //调用如下 var p1: Panda = Panda. init (name: "tuantuan" , age: 2 ) print (p1.name) 复制代码`

子类新增指定初始化方法，在确保所有属性都有初始值时，则父类的初始化方法不在可以调用

` class Panda : Animal { var weight : Float //子类新增指定初始化方法，在确保所有属性都有初始值时，则父类的初始化方法不在可以调用 init (name: String , age: Int ,weight: Float ) { self.weight = weight super. init (name:name,age:age) } } //调用如下 var p1: Panda = Panda. init (name: "tuantuan" , age: 2 ,weight: 19 ) print (p1.name) 复制代码`

#### 便捷初始化方法创建实例 ####

由于便捷初始化方法必需要调用到指定初始化方法，一个类的便捷初始化方法和指定初始化方法会形成一个路径，类的存储属性通过这条路劲收到初始值 swift

` class Panda: Animal { var weight :Float init(name: String, age: Int,weight:Float) { self.weight = weight super.init(name:name,age:age) } //声明一个便捷初始化方法 convenience init(weight:Float) { self.init(name: "Panda_007" , age: 2, weight: weight) } } //使用便捷初始化方法进行实例的初始化 let p2 = Panda.init(weight: 99) print ( "name is \(p2.name),weight is \(p2.weight)" ) 复制代码`

#### 类的必需初始化方法 required ####

使用required关键字声明的初始化方法，表示所有子类都必需提供这个初始化方法。跟覆盖继承自父类的其他方法不同，必需初始化方法不需要用override关键字标记，required标记已经隐含了覆盖的意思。

` class Animal { var name: String var age: Int //必需子类实现的初始化方法 required init (newname: String , age: Int ) { self.age = age self.name = newname } } class Panda : Animal { required init (newname: String , age: Int ) { name = newname self.age = age } } 复制代码`

#### 反初始化 deinit ####

即在类实例没用之后将其清除内存的过程；反初始化使用 **deinit** 表示，没有参数。

` class Panda : Animal { //声明反初始化方法 deinit { print ( "Panda is no longer with us" ) } } //调用：当实例为空，反初始化被触发调用 var p3: Panda ? = Panda. init (newname: "mimi" , age: 2 ) print ( "name is \(p3?.name) ,weight is \(p3?.age) " ) p3 = nil //实例为空，反初始化被触发调用 复制代码`

#### 可失败的初始化方法 init? ####

可失败的初始化方法会返回可空实例

` class Pig : Animal { var color: String //声明失败的初始化方法 init ?(name: String , age: Int ,color: String ) { guard color != "" else { //f当不满足条件则返回nil,表示初始化失败 return nil } self.color = color super. init (name: name, age: age) } } //调用 //可失败的初始化方法调用， //当调用失败，初始化值为nil var pig = Pig. init (name: "pig_9" , age: 1 , color: "" ) print (pig) 复制代码`