# Core Java 52 问（含答案） #

上篇文章 [4.9k Star 安卓面试知识点，请收下！]( https://juejin.im/post/5caacb2af265da24d320bca9 ) 翻译了 ` Mindorks` 的一份超强面试题，今天带来的是其中 ` Core Java` 部分 52 道题目的答案。题目的质量还是比较高的，基本涵盖了 Java 基础知识点，面向对象、集合、基本数据类型、并发、Java 内存模型、GC、异常等等都有涉及。整理答案的过程中才发现自己也有一些知识点记不太清了，一边回忆学习，一边整理答案。52 道题，可以代码验证的都经过我的验证，保证答案准确。

文章比较长，翻到文末可以直接获取 ` Core Java 52 问` pdf 文档。

下面就进入提问！

## Core Java ##

### 面向对象 ###

#### 1. 什么是 OOP ？ ####

别说我还真的被问到过这个问题，记得当时我第一句话就是 “万物皆对象”。当然答案很开放，说说你对面向对象的理解就行了。下面是从 [维基百科]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2F%25E9%259D%25A2%25E5%2590%2591%25E5%25AF%25B9%25E8%25B1%25A1%25E7%25A8%258B%25E5%25BA%258F%25E8%25AE%25BE%25E8%25AE%25A1 ) 总结的答案：

` Object-oriented programming` ，面向对象程序设计，是种具有对象概念的程序编程典范，同时也是一种程序开发的抽象方针。

它可能包含数据 、属性 、代码与方法。对象则指的是类的实例。它将对象作为程序的基本单元，将程序和数据封装其中，以提高软件的重用性、灵活性和扩展性，

对象里的程序可以访问及经常修改对象相关连的数据。在面向对象程序编程里，计算机程序会被设计成彼此相关的对象。

OOP 一般具有以下特征：

* 类与对象

类定义了一件事物的抽象特点。类的定义包含了数据的形式以及对数据的操作。举例来说， ` 狗` 这个类会包含狗的一切基础特征，即所有 ` 狗` 都共有的特征或行为，例如它的孕育、毛皮颜色和吠叫的能力。 对象就是类的实例。

* 封装

封装（Encapsulation）是指 OOP 隐藏了某一方法的具体运行步骤和实现细节，限制只有特定类的对象可以访问这一特定类的成员，通常暴露接口来供调用。每个人都知道怎么访问它，但却不必考虑它的内部实现细节。

举例来说， ` 狗` 这个类有 ` 吠叫()` 的方法，这一方法定义了狗具体该通过什么方法吠叫。但是，调用者并不知道它到底是如何吠叫的。

* 继承

继承性（Inheritance）是指，在某种情况下，一个类会有 ` 子类` 。子类比原本的类（称为父类）要更加具体化。例如， ` 狗` 这个类可能会有它的子类 ` 牧羊犬` 和 ` 吉娃娃犬` 。

子类会继承父类的属性和行为，并且也可包含它们自己的。这意味着程序员只需要将相同的代码写一次。

* 多态

多态（Polymorphism）是指由继承而产生的相关的不同的类，其对象对同一消息会做出不同的响应。例如，狗和鸡都有 ` 叫()` 这一方法，但是调用狗的 ` 叫()` ，狗会吠叫；调用鸡的 ` 叫()` ，鸡则会啼叫。

除了继承，接口实现，同一类中进行方法重载也是多态的体现。

#### 2. 抽象类和接口的区别 ？ ####

* 抽象类可以有默认的方法实现。接口在 jdk1.8 之前没有方法实现，1.8 之后可以使用 ` default` 关键字定义方法实现
* 抽象类可以有构造函数，接口不可以
* 子类使用 ` extends` 关键字来继承抽象类。如果子类不是抽象类的话，它需要提供抽象类中所有声明的方法的实现。子类使用关键字 ` implements` 来实现接口。它需要提供接口中所有声明的方法的实现
* 抽象方法可以有 ` public` 、 ` protected` 和 ` default` 这些修饰符，接口方法默认是 ` public` 的，可以缺省
* 抽象类的字段可以使用任何修饰符。接口中字段默认是 ` public final` 的
* 单继承 多实现
* 抽象类： ` is-a` 的关系，体现的是一种关系的延续。 接口： ` like-a` 体现的是一种功能的扩展关系

#### 3. Iterator 和 Enumeration 的区别 ？ ####

* 

函数接口不同 Enumeration 只有 2 个函数接口。通过 Enumeration，我们只能读取集合的数据，而不能对数据进行修改。

` Iterator 有 3 个函数接口。Iterator 除了能读取集合的数据之外，也能数据进行删除操作。 复制代码`
* 

Iterator 支持 fail-fast 机制，而 Enumeration 不支持。 Enumeration 是 JDK 1.0 添加的接口。使用到它的函数包括 Vector 、Hashtable 等类，这些类都是 JDK 1.0 中加入的，Enumeration 存在的目的就是为它们提供遍历接口。Enumeration 本身并没有支持同步，而在 Vector 、Hashtable 实现 Enumeration 时，添加了同步。

` 而 Iterator 是 JDK 1.2 才添加的接口，它也是为了 HashMap 、ArrayList 等集合提供遍历接口。Iterator 是支持 fail-fast 机制的：当多个线程对同一个集合的内容进行操作时，就可能会产生 fail-fast 事件。 所以 Enumeration 比 Iterator 的遍历速度更快。 复制代码`

#### 4. 你同意 组合优先于继承 吗 ？ ####

继承的功能非常强大，但是也存在诸多问题，因为它违背了封装原则 。 只 有当子类和超类之间确实存在子类型关系时，使用继承才是恰当的 。 即使如此，如果子 类和超类处在不同的包中，并且超类并不是为了继承而设计的，那么继承将会导致脆弱性 ( fragility ） 。 为了避免这种脆弱性，可以用复合和转发机制来代替继承，尤其是当存在适当 的接口可以实现包装类的时候 。 包装类不仅比子类更加健壮，而且功能也更加强大。（也就是装饰者模式）。

具体见 Effective Java 18条 复合优先于继承

#### 5. 方法重载和方法重写的区别 ？ ####

同一个类中，方法名称相同但是参数类型不同，称为方法重载。 重载的方法在编译过程中即可完成识别。具体到每一个方法调用，Java 编译器会根据所传入参数的声明类型（注意与实际类型区分）来选取重载方法。

如果子类中定义了与父类中非私有方法同名的方法，而且这两个方法参数类型不同，那么在子类中，这两个方法同样构成了重载。反之，如果方法参数类型相同， 这时候要区分是否是静态方法。如果是静态方法，那么子类中的方法会隐藏父类的方法。如果不是静态方法，就是子类重写了父类的方法、

对重载方法的区分在编译阶段已经完成，重载也被称为静态绑定，或者编译时多态。重写被称为为动态绑定。

#### 6. 你知道哪些访问修饰符 ？ 它们分别的作用 ？ ####

+----------+--------------------+------+------+------+----------+
| 访问级别 |   访问控制修饰符   | 同类 | 同包 | 子类 | 不同的包 |
+----------+--------------------+------+------+------+----------+
| 公开     | public             | √   | √   | √   | √       |
| 受保护   | protected          | √   | √   | √   | --       |
| 默认     | 没有访问控制修饰符 | √   | √   | --   | --       |
| 私有     | private            | √   | --   | --   | --       |
+----------+--------------------+------+------+------+----------+

#### 7. 一个接口可以实现另一个接口吗 ？ ####

可以，但是不是 ` implements` , 而是 ` extends` 。一个接口可以继承一个或多个接口。

#### 8. 什么是多态 ？什么是继承 ？ ####

在 java 中多态有编译期多态（静态绑定）和运行时多态（动态绑定）。方法重载是编译期多态的一种形式。方法重写是运行时多态的一种形式。

多态的另一个重要例子是父类引用子类实例。事实上，满足 is-a 关系的对象都可以看出多态。 例如， ` Cat` 类 是 ` Animal` 类的子类，所以 ` Cat is Animal` ，这就满足了 is-a 关系。

继承性（Inheritance）是指，在某种情况下，一个类会有“子类”。子类比原本的类（称为父类）要更加具体化。例如，“狗”这个类可能会有它的子类“牧羊犬”和“吉娃娃犬”。 子类会继承父类的属性和行为，并且也可包含它们自己的。这意味着程序员只需要将相同的代码写一次。

#### 9. Java 中类和接口的多继承 ####

在 java 中一个类不可以继承多个类，但是接口可以继承多个接口。

#### 10. 什么是设计模式？ ####

设计模式就不在这里展开说了。推荐一个 github 项目 [java-design-patterns]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Filuwatar%2Fjava-design-patterns ) 。 后面有机会单独写一写设计模式。

### 集合和泛型 ###

#### 11. Arrays vs ArrayLists ####

` Arrays` 是一个工具类，提供了许多操作，排序，查找数组的静态方法。

` ArrayList` 是一个动态数组队列，实现了 Collection 和 List 接口，提供了数据的增加，删除，获取等方法。

#### 12. HashSet vs TreeSet ####

` HashSet` 与 ` TreeSet` 都是基于 ` Set` 接口的实现类。其中 ` TreeSet` 是 ` Set` 的子接口 ` SortedSet` 的实现类。

` HashSet` 基于哈希表实现，它不保证集合的迭代顺序，特别是它不保证该顺序恒久不变。允许 null 值。不支持同步。

` TreeSet` 基于二叉树实现，它的元素自动排序，按照自然顺序或者提供的比较器进行排序，所以 ` TreeSet` 中元素要实现 ` Comparable` 接口。不允许 null 值。

#### 13. HashMap vs HashSet ####

+------------------------------------+---------------------------------------------------------------+
|              HASHMAP               |                            HASHSET                            |
+------------------------------------+---------------------------------------------------------------+
| 实现了 Map 接口                    | 实现了 Set 接口                                               |
| 存储键值对                         | 仅存储对象                                                    |
| 调用 put() 向 map 中添加元素       | 调用 add() 方法向 Set中                                       |
|                                    | 添加元素                                                      |
| 使用键对象来计算 hashcode 值       | 使用成员对象来计算 hashcode 值，对于两个对象来说              |
|                                    | hashcode 可能相同，所以 equals()                              |
|                                    | 方法用来判断对象的相等性，如果两个对象不同的话，那么返回false |
| HashMap 相对于 HashSet             | HashSet 较 HashMap 来说比较慢                                 |
| 较快，因为它是使用唯一的键获取对象 |                                                               |
+------------------------------------+---------------------------------------------------------------+

#### 14. Stack vs Queue ####

队列是一种基于先进先出（FIFO）策略的集合类型。队列在保存元素的同时保存它们的相对顺序：使它们入列顺序和出列顺序相同。队列在生活和编程中极其常见，就像排队，先进入队伍的总是先出去。

栈是一种基于后进先出（LIFO）策略的集合类型，当使用 foreach 语句遍历栈中的元素时，元素的处理顺序和它们被压入的顺序正好相反。就像我们的邮箱，后进来的邮件总是会先看到。

#### 15. 解释 java 中的泛型 ####

#### 16. String 类是如何实现的？它为什么被设计成不可变类 ？ ####

String 类是使用 char 数组实现的，jdk 9 中改为使用 byte 数组实现。 不可变类好处：

* 不可变类比较简单。
* 不可变对象本质上是线程安全的，它们不要求同步。不可变对象可以被自由地共享。
* 不仅可以共享不可变对象，甚至可以共享它们的内部信息。
* 不可变对象为其他对象提供了大量的构建。
* 不可变类真正唯一的缺点是，对于每个不同的值都需要一个单独的对象。

[走进 JDK 之 String]( https://juejin.im/post/5ca30c31f265da30c1724a04 )

### 对象和基本类型 ###

#### 17. 为什么说 String 不可变 ？ ####

* ` String` 是 ` final` 类，不可以被扩展
* ` private final char value[]` ，不可变
* 没有对外提供任何修改 ` value[]` 的方法

参见我的文章 [String 为什么不可变 ？]( https://juejin.im/post/59cef72b518825276f49fe40 )

#### 18. 什么是 String.intern() ？ 何时使用？ 为什么使用 ？ ####

如果常量池中存在当前字符串, 就会直接返回当前字符串. 如果常量池中没有此字符串, 会将此字符串放入常量池中后, 再返回。

将运行时需要大量使用的字符串放入常量池。

[深入解析 String.intern()]( https://link.juejin.im?target=https%3A%2F%2Ftech.meituan.com%2F2014%2F03%2F06%2Fin-depth-understanding-string-intern.html )

#### 19. 列举 8 种基本类型 ####

+----------+---------+-------------------------+--------------------------+-----------+--------------+
| 基本类型 |  大小   |         最大值          |          最小值          |  包装类   | 虚拟机中符号 |
+----------+---------+-------------------------+--------------------------+-----------+--------------+
| boolean  | -       | -                       | -                        | Boolean   | Z            |
| char     | 16 bits |                   65536 |                        0 | Character | C            |
| byte     | 8 bits  |                     127 |                     -128 | Byte      | B            |
| short    | 16 bits |                       2 | - 2                      | Short     | S            |
|          |         |                      15 |                       15 |           |              |
|          |         |                      -1 |                          |           |              |
| int      | 32 bits |                       2 |                        2 | Integer   | I            |
|          |         |                      31 |                       31 |           |              |
|          |         |                      -1 |                          |           |              |
| long     | 64 bits |                       2 |                       -2 | Long      | J            |
|          |         |                      63 |                       63 |           |              |
|          |         |                      -1 |                          |           |              |
| float    | 32 bits | 3.4028235e+38f          | -3.4028235e+38f          | Float     | F            |
| double   | 64 bits | 1.7976931348623157e+308 | -1.7976931348623157e+308 | Double    | D            |
+----------+---------+-------------------------+--------------------------+-----------+--------------+

#### 20. int 和 Integer 区别 ####

` int` 是基本数据类型，一般直接存储在栈中，更加高效

` Integer` 是包装类型，new 出来的对象存储在堆中，比较耗费资源

#### 21. 什么是自动装箱拆箱 ？ ####

把基本数据类型转换成包装类的过程叫做装箱。

把包装类转换成基本数据类型的过程叫做拆箱。

在Java 1.5之前，要手动进行装箱，

` Integer i = new Integer( 10 ); 复制代码`

java 1.5 中，提供了自动拆箱与自动装箱功能。需要拆箱和装箱的时候，会自动进行转换。

` Integer i = 10 ; //自动装箱 int b= i; //自动拆箱 复制代码`

自动装箱都是通过Integer.valueOf()方法来实现的，Integer的自动拆箱都是通过integer.intValue来实现的。

关于 Java 基本类型可以看我的一篇总结文章： [走进 JDK 之谈谈基本类型]( https://juejin.im/post/5c9cd777f265da60f96f9ccf )

#### 22. Java 中的类型转换 ####

赋值和方法调用转换规则：从低位类型到高位类型自动转换；从高位类型到低位类型需要强制类型转换：

* 布尔型和其它基本数据类型之间不能相互转换；
* ` byte` 型可以转换为 ` short` 、 ` int` 、 ` long` 、 ` float` 和 ` double`
* ` short` 可转换为 ` int` 、 ` long` 、 ` float` 和 ` double`
* ` char` 可转换为 ` int` 、 ` long` 、 ` float` 和 ` double`
* ` int` 可转换为 ` long` 、 ` float` 和 ` double`
* ` long` 可转换为 ` float` 和 ` double`
* ` float` 可转换为 ` double`

基本类型 与 对应包装类 可自动转换，这是自动装箱和折箱的原理。

两个引用类型间转换：

* 

子类能直接转换为父类 或 接口类型

* 

父类转换为子类要强制类型转换，且在运行时若实际不是对应的对象，会抛出 ` ClassCastException` 运行时异常；

#### 23. Java 值传递还是引用传递 ？ ####

值传递。

值传递（pass by value）是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

引用传递（pass by reference）是指在调用函数时将实际参数的地址直接传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

Java 调用方法传递的是实参引用的副本。

[为什么说Java中只有值传递。]( https://link.juejin.im?target=https%3A%2F%2Fwww.hollischuang.com%2Farchives%2F2275 )

#### 24. 对象实例化和初始化之间的区别 ？ ####

Initialization(实例化) 是创建新对象并且分配内存的过程。新创建的变量必须显示赋值，否则它将使用存储在该内存区域上的上一个变量包含的值。为了避免这个问题，Java 会给不同的数据类型赋予默认值：

* boolean defaults to false;
* byte defaults to 0;
* short defaults to 0;
* int defaults to 0;
* long defaults to 0L;
* char defaults to \u0000;
* float defaults to 0.0f;
* double defaults to 0.0d;
* object defaults to null.

Instantiation(初始化)是给已经声明的变量显示赋值的过程。

` int j; // Initialized variable (int defaults to 0 right after) j = 10 ; // Instantiated variable 复制代码`

#### 25. 局部变量、实例变量以及类变量之间的区别？ ####

局部变量仅仅存在于创建它的方法中，他们被保存在栈内存，在方法外无法获得它们的引用。Java 的方法执行不是依赖寄存器的，而是栈帧，每个方法的执行和结束都伴随着栈帧的入栈和出栈，也伴随着局部变量的创建和释放。

实例变量也就是成员变量，声明在类中，依赖类实例而存在，不同类实例中变量值也可能不同。

类变量也就是静态变量，在所有类实例中只有一个值，在一个地方改变它的值将会改变所有类实例中的值。

### Java 内存模型和垃圾收集器 ###

#### 26. 什么是垃圾收集器 ？ 它是如何工作的 ？ ####

> 
> 
> 
> Java 和 C++ 之前有一堵由内存动态分配和垃圾收集技术所围成的高墙，墙外面的人想进去，墙里面的人想出去。
> 
> 

垃圾收集器主要用来回收堆上的无用对象，Java 开发者只管创建和使用对象，JVM 来为你自动分配和回收内存。

JVM 通过可达性分析算法来判定对象是否存活。这个算法的基本思路就是通过一系列的称为 ` GC Roots` 的对象作为起始点，从这些节点开始向下搜索，搜索所走过的路径称为引用链（Reference Chain），当一个对象到 ` GC Roots` 没有任何引用链相连（用图论的话来说，就是从 GC Roots 到这个对象不可达）时，则证明此对象是不可用的。

即使在可达性分析算法中不可达的对象，也并非是 非死不可 的，这时候他们暂时处于缓刑阶段，要真正宣告一个对象死亡，至少要经历两次标记过程。

更多详细内容可以阅读 ` 《深入理解 Java 虚拟机》` 第三章 ` 垃圾收集器与内存分配策略` 。

#### 27. 什么是 java 内存模型？ 它遵循了什么原则？它的堆栈是如何组织的 ？ ####

Java 虚拟机规范中试图定义一种 Java 内存模型（Java Memory Model，JMM）来屏蔽掉各种硬件和操作系统的内存访问差异，以实现让 Java 程序在各种平台下都能达到一致的内存访问效果。JMM 是语言级的内存模型，它确保在不同的编译器和不同的处理器平台上，通过禁止特定类型的编译器重排序和处理器重排序，为程序员提供一致的内存可见性保证。

JMM 内存模型的抽象表示如下：

![](https://user-gold-cdn.xitu.io/2019/4/8/169fd6645a3b9133?imageView2/0/w/1280/h/960/ignore-error/1)

结合上图，在 Java 中，所有实例域、静态域和数组元素都存储在堆内存中，堆内存在线程之间共享。局部变量、方法定义参数和异常处理器参数在栈中，不会在线程之间共享，它们不会有内存可见性问题，也不会受内存模型影响。

Java 线程之间的通信由 JMM 控制，JMM 决定一个线程对共享变量的写入何时对另一个线程可见。JMM 通过控制主内存与每个线程的本地内存之间的交互，来为 Java 程序员提供内存可见性保证。

更多详细内容可以阅读 ` 《Java 并发编程的艺术》` 。

#### 28. 什么是 内存泄漏，java 如何处理它 ？ ####

内存泄露就是不会再被使用的对象无法被 GC 回收，即这些对象在可达性分析中是可达的，但在程序中的确不会再被使用。比如长生命周期的对象引用了短生命周期的对象，导致短生命周期对象不能被回收。

Java 应该不会处理内存泄漏，我们能做的更多是防患于未然，以及使用合理手段监测，比如 Android 里常用的 ` LeakCanary` ，详细原理可以看我之前的一篇文章 [LeakCanary 源码解析]( https://juejin.im/post/5a9d46d2f265da237d0280a3 ) 。

#### 29. 什么是 强引用，软引用，弱引用，虚引用 ？ ####

在 JDK 1.2 之后，Java 对引用的概念进行了扩充，将引用分为 强引用、软引用、弱引用、虚引用 4 种，这 4 种引用强度依次逐渐减弱。

* 

强引用就是指程序代码之中普遍存在的，类似 ` Object obj = new Object()` 这类的引用，只要强引用还存在，GC 永远不会回收掉被引用的对象。

* 

软引用是用来描述一些还有用但并非必须的对象。对于软引用关联着的对象，在系统将要发生内存溢出异常之前，将会把这些对象列进回收范围之中进行第二次回收。如果这次回收还没有足够的内存，才会抛出内存溢出异常。在 JDK 1.2 之后，提供了 ` SoftReference` 类来实现软引用。

* 

弱引用也是用来描述非必须对象的，但它的强度比软引用要弱一些，被弱引用关联的对象只能生存到下一次 GC 发生之前。当 GC 工作时，无法当前内存是否足够，都会回收掉只被弱引用关联的对象。在 JDK 1.2 之后，提供了 ` WeakReference` 类来实现弱引用。

* 

虚引用也称为幽灵引用或者幻影引用，它是最弱的一种引用关系。一个对象是否有虚引用的存在，完全不会对其生存时间构成影响，也无法通过虚引用来取得一个对象实例。为一个对象设置虚引用关联的唯一目的就是能在这个对象被 GC 回收时收到一个系统通知。在 JDK 1.2 之后，提供了 ` PhantomReference` 类来实现虚引用。

### 并发 ###

#### 30. 关键字 synchronized 的作用 ？ ####

关键字 synchronized 可以修饰方法或者以同步块的形式来进行使用，它主要确保多个线程在同一时刻，只能有一个线程处于方法或同步块中，它保证了线程对变量访问的可见性和排他性。

* 对于普通同步方法，锁是当前实例对象
* 对于静态同步方法，锁是当前类的 Class 对象
* 对于同步方法块，锁是 Synchonized 括号里配置的对象

#### 31. ThreadPoolExecutor 作用 ？ ####

在 Java 中，使用线程来执行异步任务。Java 线程的创建与销毁需要一定的开销，如果我们为每一个任务都创建一个新线程来执行，这些线程的创建与销毁将消耗大量的计算资源。同时，为每一个任务创建一个新线程来执行，这种策略可能会使处于高负荷的应用最终崩溃。

关于线程池的详细介绍，推荐一篇文章 [Java并发编程：线程池的使用]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3932921.html )

#### 32. 关键字 volatile 的作用 ？ ####

volatile 是轻量级的 synchronized，它在多处理器开发中保证了共享变量的 ` 可见性` 。volatile 用来修饰字段(成员变量),就是告知程序任何对该变量的访问均需从共享内存中获取，而对它的改变必须同步刷新回共享内存，它能保证所有线程对变量访问的可见性。

但是，过多的使用 volatile 是不必要的，因为它会降低程序执行的效率。

### 异常 ###

#### 33. try{} catch{} finally{} 是如何工作的 ? ####

` try` 代码块用来标记需要进行异常监控的代码

` catch` 代码块跟在 ` try` 代码块之后，用来捕获 try 代码块中触发的某种指定类型的异常。除了声明所捕获的异常类型之外，catch 代码块还定义了针对该异常类型的异常处理器。在 Java 中，try 代码块后面可以跟着多个 catch 代码块，来捕获不同的异常。Java 虚拟机会从上至下匹配异常处理器。因此，前面的 catch 代码块所捕获的异常类型不能覆盖后面的，否则编译器会报错。

` finally` 代码块跟在 ` try` 代码块和 ` catch` 代码块之后，用来声明一段必定运行的代码。它的设计初衷是为了避免跳过某些关键的清理代码，例如关闭已打开的系统资源。

在编译生成的字节码中，每个方法都附带一个异常表。异常表中的每一个条目代表一个异常处理器，并且由 from 指针，to 指针， target 指针以及所捕获的异常类型构成。这些指针的值是字节码索引（bytecode index，bci），用以定位字节码。

其中，from 指针和 to 指针标示了该异常处理器所监控的范围，例如 try 代码块所覆盖的范围。target 指针则指向异常处理器的起始位置，比如 catch 代码块的起始位置。

finally 代码块的编译比较复杂。当前版本 Java 编译器的做法，是复制 finally 代码块的内容，分别放在try-catch 代码块所有正常执行路径以及异常执行路径的出口中。

以上内容来自极客时间专栏 ` 深入拆解 Java 虚拟机` 。

#### 34. Checked Exception 和 Un-Checked Exception 区别 ？ ####

在 Java 中，所有异常都是 Throwable 类或者其子类的实例。Throwable 有两大直接子类。一个是 ` Error` ，涵盖程序不应捕获的异常。当 Error 发生时，它的执行状态已经无法恢复，需要终止线程甚至虚拟机。第二个子类是 Exception，涵盖程序可能需要捕获并且处理的异常。

Exception 有一个特殊的子类 RuntimeException，运行时异常，用来表示 “程序虽然无法继续执行，但还能抢救一下” 的情况。

RuntimeException 和 Error 属于 Java 里的非检查异常（unchecked exception）。其他异常则属于检查异常（checked exception）。在 Java 语法中，所有的检查异常都需要程序显式地捕获，或者在方法声明中用 throws 关键字标注。通常情况下，程序中自定义的异常应为检查异常，以便最大化利用 Java 编译器的编译时检查。

以上内容来自极客时间专栏 ` 深入拆解 Java 虚拟机` 。

### 其他 ###

#### 35. 什么是序列化？如何实现 ？ ####

序列化是将对象转换成字节流以便持久化存储的过程。它可以保存对象的状态和数据，方便在特定时刻重新构建该对象。在 Android 中，一般使用 ` Serializable` , ` Externalizable` (implements Serializable) 或者 ` Parcelable` 接口。

` Serializable` 最容易实现，直接实现接口即可。 ` Externalizable` 可以在序列化的过程中插入一些自己的逻辑代码，考虑到它是 Java 早期版本的遗留物，现在基本已经没人再使用它。在 Android 中推荐使用 ` Parcelable` ，它就是为 Android 而实现，性能是 ` Serializable` 的十倍，因为 ` Serializable` 使用了反射。反射不仅慢，还会创建大量临时对象，导致频繁 GC。

例子：

` /** * Implementing the Serializeable interface is all that is required */ public class User implements Serializable { private String name; private String email; public User () { } public String getName () { return name; } public void setName ( final String name) { this.name = name; } public String getEmail () { return email; } public void setEmail ( final String email) { this.email = email; } } 复制代码`

Parcelable 需要多一些工作：

` public class User implements Parcelable { private String name; private String email; /** * Interface that must be implemented and provided as a public CREATOR field * that generates instances of your Parcelable class from a Parcel. */ public static final Creator<User> CREATOR = new Creator<User>() { /** * Creates a new USer object from the Parcel. This is the reason why * the constructor that takes a Parcel is needed. */ @Override public User createFromParcel (Parcel in) { return new User(in); } /** * Create a new array of the Parcelable class. * @return an array of the Parcelable class, * with every entry initialized to null. */ @Override public User[] newArray( int size) { return new User[size]; } }; public User () { } /** * Parcel overloaded constructor required for * Parcelable implementation used in the CREATOR */ private User (Parcel in) { name = in.readString(); email = in.readString(); } public String getName () { return name; } public void setName ( final String name) { this.name = name; } public String getEmail () { return email; } public void setEmail ( final String email) { this.email = email; } @Override public int describeContents () { return 0 ; } /** * This is where the parcel is performed. */ @Override public void writeToParcel ( final Parcel parcel, final int i) { parcel.writeString(name); parcel.writeString(email); } } 复制代码`

#### 36. 关键字 transient 的作用 ？ ####

` transient` 很简单，它的作用就是让被其修饰的成员变量在序列化的过程中不被序列化。

#### 37. 什么是匿名内部类 ？ ####

匿名内部类是唯一一种没有构造器的类。正因为其没有构造器，所以匿名内部类的使用范围非常有限，大部分匿名内部类用于接口回调。匿名内部类在编译的时候由系统自动起名为 ` Outter$1.class` 。一般来说，匿名内部类用于继承其他类或是实现接口，并不需要增加额外的方法，只是对继承方法的实现或是重写。

Android 中应用最常见的就是各种点击事件。

#### 38. 对象的 == 和 .equals 区别 ？ ####

对于对象而言， ` ==` 永远比较的都是其内存地址。而 ` equals()` 则要看该对象是否重写了 ` equals()` 方法，如果没有则会调用父类的 ` equals()` 方法，如果父类也没有实现的话，就不断向上追溯，直至 ` Object` 类。看一下 ` Object.java` 中的 ` equals()` 方法:

` public boolean equals (Object obj) { return ( this == obj); } 复制代码`

在 ` Object` 中， ` equals` 等同于 ` ==` ，都是比较内存地址。再看一下不是比较内存地址的， ` String.equals()` ：

` public boolean equals (Object anObject) { if ( this == anObject) { return true ; } if (anObject instanceof String) { String anotherString = (String)anObject; int n = value.length; if (n == anotherString.value.length) { char v1[] = value; char v2[] = anotherString.value; int i = 0 ; while (n-- != 0 ) { if (v1[i] != v2[i]) return false ; i++; } return true ; } } return false ; } 复制代码`

这里比较的就不再是内存地址，而是其实际的值。

#### 39. hashCode() 和 equals() 用处 ？ ####

` equals()` 用于判断两个对象是否相等，未被重写的话就是判断内存地址，和 ` ==` 语义一致。重写了的话，就按照重写的逻辑进行判断。

` hashCode()` 用于计算对象的哈希码，默认实现是将对象的内存地址作为哈希码返回，可以保证不同对象的返回值不同。理论上， ` hashCode` 也可以用来比较对象是否相等。 ` hashCode()` 主要用在哈希表中，比如 ` HashMap` 、 ` HashSet` 等。

当我们向哈希表(如HashSet、HashMap等)中添加对象object时，首先调用hashCode()方法计算object的哈希码，通过哈希码可以直接定位object在哈希表中的位置(一般是哈希码对哈希表大小取余)。如果该位置没有对象，可以直接将object插入该位置；如果该位置有对象(可能有多个，通过链表实现)，则调用equals()方法比较这些对象与object是否相等，如果相等，则不需要保存object；如果不相等，则将该对象加入到链表中。

` equals()` 相等， ` hashCode` 必然相等。反之则不然， ` hashCode` 相等， ` equals()` 不能保证一定相等。

#### 40. 构造函数中为什么不能调用抽象方法 ？ ####

构造函数中不能调用抽象方法，说的更严谨一点， **构造函数中不能调用可被覆盖的方法** 。

先看这样一个例子：

` public abstract class Super { Super(){ overrideMe(); } abstract void overrideMe () ; } public class Sub extends Super { private final Instant instant; public Sub () { instant = Instant.now(); } @Override void overrideMe () { System.out.println(instant); } public static void main (String[] args) { Sub sub= new Sub(); sub.overrideMe(); } } 复制代码`

最后的打印结果：

` null 2019-04-01T02:42:13.947Z 复制代码`

第一次打印出的是 ` null` ，因为 overrideMe 方法被 Super 构造器调用的时候，构造器 Sub 还没有机会初始化 instant 域 。 注意，这个程序观察到的 final 域处于两种不同的状态 。

超类的构造器在子类的构造器之前运行，所以，子类中覆盖版本的方法将会在子类的构造器运行之前先被 调用 。 如果该覆盖版本的方法依赖于子类构造器所执行的任何初始化工作，该方法将不会如预期般执行 。

#### 41. 你什么时候会使用 final 关键字 ？ ####

对于一个 ` final` 变量，如果是基本数据类型的变量，则其数值一旦在初始化之后便不能更改；如果是引用类型的变量，则在对其初始化之后便不能再让其指向另一个对象。

另外，匿名内部类中使用的外部局部变量只能是 ` final` 变量。

当 ` final` 变量是基本数据类型以及 String 类型时，如果在编译期间能知道它的确切值，则编译器会把它当做编译期常量使用。

用 ` final` 修饰方法参数也是为了强调参数不可改变。

用 ` final` 修饰类表示类不可被继承。

[浅析 Java 的 final 关键字]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3736238.html )

#### 42. final, finally 和 finalize 的区别 ？ ####

final 和 finally 就不再说了。重点看看 finalize。

如果类中重写了 ` finalize` 方法，当该类对象被回收时， ` finalize` 方法有可能会被触发。 ` Effective Java` 中明确说明 ` 终结方法（finalize）通常是不可预测的，也是很危险的，一般情况下是不必要的。`

JVM 不仅不保证 ` finalize` 方法可以被及时执行，而且根本就不保证它们会被执行。所以不要依赖 ` finalize` 方法来做一些例如 释放资源的操作。可能会延时对象的回收，造成性能损失。

#### 43. Java 中 static 关键字的含义 ？ ####

` static` 就是为了方便在没有创建对象的情况下来进行调用（方法/变量）。

* 

static 方法一般称作静态方法，由于静态方法不依赖于任何对象就可以进行访问，因此对于静态方法来说，是没有 this 的，因为它不依附于任何对象，既然都没有对象，就谈不上 this 了。并且由于这个特性，在静态方法中不能访问类的非静态成员变量和非静态成员方法，因为非静态成员方法/变量都是必须依赖具体的对象才能够被调用。

* 

static 变量也称作静态变量，静态变量和非静态变量的区别是：静态变量被所有的对象所共享，在内存中只有一个副本，它当且仅当在类初次加载时会被初始化。而非静态变量是对象所拥有的，在创建对象的时候被初始化，存在多个副本，各个对象拥有的副本互不影响。

* 

static 关键字还有一个比较关键的作用就是 用来形成静态代码块以优化程序性能。static 块可以置于类中的任何地方，类中可以有多个 static 块。在类初次被加载的时候，会按照 static 块的顺序来执行每个 static 块，只会在类加载的时候执行一次。

static 成员变量的初始化顺序按照定义的顺序进行初始化。

#### 44. 静态方法可以重写吗 ? ####

你可以重写，但这并不是多态的体现，并不是真正意义上的重写。子类的静态方法会隐藏父类的静态方法，这两个方法并没有什么关系，具体调用哪一个方法是看调用者是哪个对象的引用，并不存在多态。只有普通的方法调用才可以是多态的。

#### 45. 静态代码块如何运行 ？ ####

静态代码块随着类的加载而执行，而且只执行一次。

静态代码块经过编译后是放在 ` <clinit>` 中, ` <clinit>` 在jvm第一次加载class文件时调用，包括静态变量初始化语句和静态块的执行。

#### 46. 什么是反射 ？ ####

反射 (Reflection) 是 Java 的特征之一，它允许运行中的 Java 程序获取自身的信息，并且可以操作类或对象的内部属性。

Oracle 官方对反射的解释是：

> 
> 
> 
> Reflection enables Java code to discover information about the fields,
> methods and constructors of loaded classes, and to use reflected fields,
> methods, and constructors to operate on their underlying counterparts,
> within security restrictions. The API accommodates applications that need
> access to either the public members of a target object (based on its
> runtime class) or the members declared by a given class. It also allows
> programs to suppress default reflective access control.
> 
> 

简而言之，通过反射，我们可以在运行时获得程序或程序集中每一个类型的成员和成员的信息。程序中一般的对象的类型都是在编译期就确定下来的，而 Java 反射机制可以动态地创建对象并调用其属性，这样的对象的类型在编译期是未知的。所以我们可以通过反射机制直接创建对象，即使这个对象的类型在编译期是未知的。

反射的核心是 JVM 在运行时才动态加载类或调用方法/访问属性，它不需要事先（写代码的时候或编译期）知道运行对象是谁。

Java 反射主要提供以下功能：

在运行时判断任意一个对象所属的类； 在运行时构造任意一个类的对象； 在运行时判断任意一个类所具有的成员变量和方法（通过反射甚至可以调用private方法）； 在运行时调用任意一个对象的方法

[深入 Java 反射]( https://link.juejin.im?target=https%3A%2F%2Fwww.sczyh30.com%2Fposts%2FJava%2Fjava-reflection-1%2F )

IDE 的智能提示就是利用反射。

#### 47. 什么是依赖注入 ？列举几个库 ？你使用过吗 ？ ####

理解的还不够透彻，放上来一篇网上的写的不错的文章： [轻松理解 Java开发中的依赖注入(DI)和控制反转(IOC)]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F506dcd94d4f9 )

#### 48. StringBuilder 如何避免不可变类 String 的分配问题？ ####

` StringBuilder` 内部维护了一个可变长的 ` char[]` ，用来存储和拼接字符串，从而避免了因 String 是不可变类带来的频繁创建 String 对象的问题。

#### 49. StringBuffer 和 StringBuilder 区别 ？ ####

` StringBuffer` 和 ` StringBuilder` 在使用上基本没有区别。 ` StringBuffer` 通过 synchronized 关键字保证了线程安全，而 ` StringBuilder` 没有任何同步操作。所以在确定无线程同步问题时，使用 ` StringBuilder` 效率更高。

#### 50. Enumeration and an Iterator 区别 ？ ####

重复了，见第 3 题。

#### 51. fail-fast and fail-safe 区别 ？ ####

` fail-fast` 机制在遍历一个集合时，当集合结构被修改，会抛出 ` Concurrent Modification Exception` 。迭代器在遍历过程中是直接访问内部数据的，因此内部的数据在遍历的过程中无法被修。 为了保证不被修改，迭代器内部维护了一个标记 ` “mode”` ，当集合结构改变（添加删除或者修改），标记 ` "mode"` 会被修改， 而迭代器每次的 ` hasNext()` 和 ` next()` 方法都会检查该 ` "mode"` 是否被改变，当检测到被修改时，抛出 ` Concurrent Modification Exception` 。

` fail-safe` 任何对集合结构的修改都会在一个复制的集合上进行修改，因此不会抛出 ` ConcurrentModificationException` 。

` fail-safe` 机制有两个问题:

* 

需要复制集合，产生大量的无效对象，开销大

* 

无法保证读取的数据是目前原始数据结构中的数据

#### 52. 什么是 NIO ？ ####

在 JDK 1. 4 中 新 加入 了 NIO( New Input/ Output) 类, 引入了一种基于通道和缓冲区的 I/O 方式， 它可以使用 Native 函数库直接分配堆外内存，然后通过一个存储在 Java 堆的 DirectByteBuffer 对象作为这块内存的引用进行操作， 避免了在 Java 堆和 Native 堆中来回复制数据。

NIO 是一种同步非阻塞的 IO 模型。同步是指线程不断轮询 IO 事件是否就绪，非阻塞是指线程在等待 IO 的时候，可以同时做其他任务。 同步的核心就是 Selector，Selector 代替了线程本身轮询 IO 事件，避免了阻塞同时减少了不必要的线程消耗；非阻塞的核心就是通道和缓冲区， 当 IO 事件就绪时，可以通过写道缓冲区，保证 IO 的成功，而无需线程阻塞式地等待。

[深入理解 Java NIO]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fgeason%2Fp%2F5774096.html )

## End ##

> 
> 
> 
> 文章首发于微信公众号： **` 秉心说`** ， 专注 Java 、 Android 原创知识分享，LeetCode 题解，欢迎关注！
> 
> 

![](https://user-gold-cdn.xitu.io/2019/3/30/169cf046d9579e78?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 微信搜索 **` 秉心说`** , 或者扫码关注，回复 ` Core Java` 即可领取所有回答 pdf 文档 。
> 
>