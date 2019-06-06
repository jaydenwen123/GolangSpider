# Java泛型 #

java泛型在平时开发中或者阅读项目源码的时候都见过他，我们虽然知道它，但是大多数我们也是对他并不太了解。这个星期我花了点时间重新复习了一下泛型的一些内容，这篇文章是对复习笔记的简单整理，里面内容只是整理一些我们经常忽视或者有很模糊的知识点。

* 概述
* 类型擦除
* 泛型晋级使用
* 通配符
* 其他

### 1. 概述 ###

泛型指的是可以将类型作为参数进行传递，其本质上就是类型参数化。比如:我们平时定义一个方法的时候，常会指定要传入一个具体类对象作为参数。而如果使用泛型，那么这个具体传入类的对象，就可以指定为某个类型，而不必指定具体的类。也就是我们将某个类型作为参数进行传递了。

` //普通方法 public void test Value(String s) {} //泛型方法 public <T> void test Value(T t) {} 复制代码`

#### 他与使用Object有什么区别？ ####

如果我们使用Object，就要将传入的类型强制转换成我们需要的类型，如果传入的类型不匹配将会导致程序包 ` ClassCastException` 异常。比如下面的代码,testObj()传入的是int类型的值，程序在执行的时候将会出错:

` public void test Obj(Object o){ String name= (String) o; } 复制代码`

我们可以通过泛型将来实现这样的需求:

` public <O extends String> void test Obj(O o) { String name = o; } 复制代码`

#### 使用泛型有哪些好处？ ####

* 它可以避免类型强制转换，而引起的程序异常。
* 可以是代码更加简洁易度。
* 是代码更加灵活，可定制型强。

### 2. 类型擦除 ###

泛型值存在于编译期，代码在进入虚拟机后泛型就会会被擦除掉，这个者特性就叫做类型擦除。当泛型被擦除后，他有两种转换方式，第一种是如果泛型没有设置类型上限，那么将泛型转化成Object类型，第二种是如果设置了类型上限，那么将泛型转化成他的类型上限。

` //未指定上限 public class Test1<T> { T t; public T getValue () { return t; } public void set Vale(T t) { this.t = t; } } //指定上限 public class Test2<T extends String> { T t; public T getT () { return t; } public void set T(T t) { this.t = t; } } //通过反射调用获取他们的属性类型 @Test public void testType1 () { Test1<String> test 1 = new Test1<>(); test 1.setVale( "11111" ); Class<? extends Test1> aClass = test 1.getClass(); for (Field field : aClass.getDeclaredFields()) { System.out.println( "Test1属性:" + field.getName() + "的类型为：" + field.getType().getName()); } Test2 test 2 = new Test2(); test 2.setT( "2222" ); Class<? extends Test2> aClass2 = test 2.getClass(); for (Field field : aClass2.getDeclaredFields()) { System.out.println( "test2属性：" + field.getName() + "的类型为：" + field.getType().getName()); } } 复制代码`

上面方法打印的结果:

` Test1属性:t的类型为：java.lang.Object Test2属性：t的类型为：java.lang.String 复制代码`

### 3. 泛型晋级使用 ###

#### 继承关系 ####

即设置泛型上限，传入的泛型必须是String类型或者是他的子类

> 
> 
> 
> 这里有一个小小的坑，感谢以为热心网友的反馈。如果读者看到这段请想一想String的特性。这个问题在文章末尾的评论去有答案。
> 
> 

` public <T extends String> void test Type(T t) {} 复制代码`

#### 依赖关系的使用 ####

泛型间可以存在依赖关系，比如下面的C是继承自E。即传入的类型是E类型或者是E类型的子类

` public <E, C extends E> void test Dependys(E e, C c) {} 复制代码`

### 4. 通配符 ###

当我们不知道或者不关心实际操作类型的时候我们可以使用 ` 无限通配符` ，当我们不指定或者不关心操作类型，但是又想进行一定范围限制的时候，我们可以通过添加 ` 上限` 或 ` 下限` 来起到限制作用。

#### <?>无限通配符 ####

无限通配符表示的是未知类型，表示不关心或者不能确定实际操作的类型，一般配合容器类使用。

` public void test V(List<?> list) {} 复制代码`

需要注意的是: 无限通配符只能读的能力，没有写的能力。

` public void test V(List<?> list) { Object o = list.get(0); //编译器不允许该操作 // list.add( "jaljal" ); } 复制代码`

上面的List<?>为无限通配符，他只能使用get()获取元素，但不能使用add()方法添加元素。（即使修改元素也不被允许）

#### <? extends T> ####

定义了上限，期只有读的能力。此方式表示参数化的类型可能是所 ` 指定的类型` ，或者是 ` 此类型的子类` 。

` //t1要么是Test2，要么是Test2的子类 public void test C(Test1<? extends Test2> t1) { Test2 value = t1.getValue(); System.out.println( "testC中的：" + value.getT()); } 复制代码`

#### <? super T> ####

定义了下限，有读的能力以及部分写的能力，子类可以写入父类。此方式表示参数化的类型可能是 ` 指定的类型` ，或者是 ` 此类型的父类`

` //t1要么是Test5,要么是Test5的父类 public void test B(Test1<? super Test5> t1) { //子类代替父类 Test2 value = (Test2) t1.getValue(); System.out.println(value.getT()); } 复制代码`

#### 通配符不能用作返回值 ####

如果返回值依赖类型参数，不能使用通配符作为返回值。可以使用类型参数返回方式：

` public <T> T test A(T t, Test1<T> test 1) { System.out.println( "这是传入的T:" + t); t = test 1.t; System.out.println( "这是赋值后的T:" + t); return t; } 复制代码`

* 要从泛型类取数据时，用extends；
* 要往泛型类写数据时，用super；
* 既要取又要写，就不用通配符（即extends与super都不用）。

> 
> 
> 
> 泛型中只有通配符可以使用super关键字，类型参数不支持 这种写法
> 
> 

### 5. 其他 ###

#### 什么时候使用通配符 ####

* 通配符形式和类型参数经常 ` 配合使用`
* ` 类型参数` 的形式都可以 ` 替代` 通配符的形式
* ` 能用通配符的就用通配符` ,因为通配符形式上往往更为 ` 简单` 、 ` 可读性也更好` 。
* 类型参数之间有 ` 依赖关系` 、 ` 返回值` 依赖类型参数或者需要 ` 写操作` ，则只能用 ` 类型参数` 。

#### 查看源码使用 ####

如果想查找源码中的相关使用可以 ` Collections` 类的的下面这些方法:

` public static <T extends Comparable<? super T>> void sort(List<T> list) public static <T> void sort(List<T> list, Comparator<? super T> c) public static <T> void copy(List<? super T> dest, List<? extends T> src) public static <T> T max(Collection<? extends T> coll, Comparator<? super T> comp) 复制代码`

## 参考 ##

[java 泛型，你了解类型擦除吗？]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fbriblue%2Farticle%2Fdetails%2F76736356 )

[深入理解 Java 泛型]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011240877%2Farticle%2Fdetails%2F53545041 )

[(36) 泛型 (中) - 解析通配符 / 计算机程序的思维逻辑]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2Fte9K3alu8P8jRUUU2AkO3g )