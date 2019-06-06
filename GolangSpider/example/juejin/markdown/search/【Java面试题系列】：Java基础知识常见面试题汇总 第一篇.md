# 【Java面试题系列】：Java基础知识常见面试题汇总 第一篇 #

> 
> 
> 
> 文中面试题从茫茫网海中精心筛选，如有错误，欢迎指正！
> 
> 

## 1.前言 ##

​ 参加过社招的同学都了解，进入一家公司面试开发岗位时，填写完个人信息后，一般都会让先做一份笔试题，然后公司会根据笔试题的回答结果，确定要不要继续此次面试，如果答的不好，有些公司可能会直接说“技术经理或者总监在忙，你先回去等通知吧”，有些公司可能会继续面试，了解下你的项目经验等情况。

​ 至少在工作的前5年甚至更久，面试一般不会跳过笔试题这个环节（大牛，个别公司除外），我自己也记不清自己面试过多少家公司，做过多少份面试题了，导致现在有时逛街，总感觉很多地方似曾相识，感觉自己多年前曾经来面过试，一度自嘲，一度也怀疑，自己当年是靠什么在上海坚持下来的，所以说面试题对于求职来说，还是非常重要的。

​ 网上搜索“Java面试题”几个关键字也是有很多很多的文章讲解，为什么我还要自己总结呢？主要有以下几个原因：

* 文章太多，反倒不知道该看哪个（就如一本书中所说太多的资讯等于没有资讯）
* 文章的准确性不高（曾多次发现描述不正确或代码跑不起来的情况）
* 可以加深自己的理解和记忆
* 一劳永逸，下次不用再从网上慢慢筛选，看自己整理的就好了

## 2.提纲 ##

本篇主要整理下Java基础知识的面试题，主要包含以下几点：

### 2.1 Integer与int的区别 ###

### 2.2 ==和equals的区别 ###

### 2.3 String,StringBuilder,StringBuffer的区别 ###

### 2.4 装箱和拆箱 ###

### 2.5Java中的值传递和引用传递 ###

接下来一一讲解。

## 3.Integer与int的区别 ##

### 3.1基本概念区分： ###

* Integer是int的包装类(引用类型)，int是java的一种基本数据类型(值类型)。
* Integer变量必须实例化后才能使用，而int变量不需要。
* Integer实际是对象的引用，当new一个Integer时，实际上是生成一个指针指向此对象；而int则是直接存储数据值。
* Integer的默认值是null，int的默认值是0。

### 3.2Integer与int几种常用的比较场景： ###

1)两个new Integer()变量相比较，永远返回false

` Integer i = new Integer( 100 ); Integer j = new Integer( 100 ); System.out.println(i == j); // false 复制代码`
> 
> 
> 
> 
> 两个通过new生成的Integer变量生成的是两个对象，其内存地址不同
> 
> 

2)非new生成的Integer变量和new Integer()生成的变量相比较，永远返回false

` Integer i = new Integer( 100 ); Integer j = 100 ; System.out.println(i == j); // false 复制代码`
> 
> 
> 
> 
> 非new生成的Integer变量指向的是Java常量池中的对象，而new Integer()生成的变量指向堆中新建的对象，两者在内存中的地址不同
> 
> 

3)两个非new生成的Integer变量比较，如果两个变量的值在区间-128到127 之间，则比较结果为true，如果两个变量的值不在此区间，则比较结果为 false。

` Integer i = 100 ; Integer j = 100 ; System.out.println(i == j); //true Integer i1 = 128 ; Integer j1 = 128 ; System.out.println(i1 == j1); //false 复制代码`

为什么会这样呢，我们来分析下原因：

Integer i = 100; 在编译时，会翻译成 Integer i = Integer.valueOf(100); ,而Java中Integer类的valueOf方法的源码如下：

` public static Integer valueOf ( int i) { if (i >= IntegerCache.low && i <= IntegerCache.high) return IntegerCache.cache[i + (-IntegerCache.low)]; return new Integer(i); } private static class IntegerCache { static final int low = - 128 ; static final int high; static final Integer cache[]; static { // high value may be configured by property int h = 127 ; String integerCacheHighPropValue = sun.misc.VM.getSavedProperty( "java.lang.Integer.IntegerCache.high" ); if (integerCacheHighPropValue != null ) { try { int i = parseInt(integerCacheHighPropValue); i = Math.max(i, 127 ); // Maximum array size is Integer.MAX_VALUE h = Math.min(i, Integer.MAX_VALUE - (-low) - 1 ); } catch ( NumberFormatException nfe) { // If the property cannot be parsed into an int, ignore it. } } high = h; cache = new Integer[(high - low) + 1 ]; int j = low; for ( int k = 0 ; k < cache.length; k++) cache[k] = new Integer(j++); // range [-128, 127] must be interned (JLS7 5.1.7) assert IntegerCache.high >= 127 ; } private IntegerCache () {} } 复制代码`

从源码我们可以看出

> 
> 
> 
> Java对于-128到127之间的数，会进行缓存。 所以 Integer i = 100 时，会将100进行缓存，下次再写Integer j =
> 100时，就会直接从缓存中取，就不会new了。
> 
> 

4)Integer变量和int变量比较时，只要两个变量的值是向等的，则结果为true

` Integer i = new Integer( 100 ); int j = 100 ; System.out.print(i == j); //true 复制代码`
> 
> 
> 
> 
> 因为包装类Integer和基本数据类型int比较时，Java会自动拆包装为int，然后进行比较，实际上就变为两个int变量的比较
> 
> 

## 4.==和equals的区别 ##

### 4.1基本概念区分 ###

1)对于==，比较的是值是否相等

` 如果作用于基本数据类型的变量，则直接比较其存储的 “值”是否相等； 复制代码`

​ 如果作用于引用类型的变量，则比较的是所指向的对象的地址是否相等。

> 
> 
> 
> 其实==比较的不管是基本数据类型，还是引用数据类型的变量，比较的都是值，只是引用类型变量存的值是对象的地址
> 
> 

2)对于equals方法，比较的是是否是同一个对象

​ equals方法不能作用于基本数据类型的变量，equals继承Object类；

​ 如果没有对equals方法进行重写，则比较的是引用类型的变量所指向的对象的地址；

​ 诸如String、Date等类对equals方法进行了重写的话，比较的是所指向的对象的内容。

3)equals()方法存在于Object类中，因为Object类是所有类的直接或间接父类，也就是说所有的类中的equals()方法都继承自Object类，在所有没有重写equals()方法的类中，调用equals()方法其实和使用==的效果一样，也是比较的地址值，不过，Java提供的所有类中，绝大多数类都重写了equals()方法，重写后的equals()方法一般都是比较两个对象的值,比如String类。

Object类equals()方法源码：

` public boolean equals (Object obj) { return ( this == obj); } 复制代码`

String类equals()方法源码：

` public boolean equals (Object anObject) { if ( this == anObject) { return true ; } if (anObject instanceof String) { String anotherString = (String)anObject; int n = value.length; if (n == anotherString.value.length) { char v1[] = value; char v2[] = anotherString.value; int i = 0 ; while (n-- != 0 ) { if (v1[i] != v2[i]) return false ; i++; } return true ; } } return false ; } 复制代码`

### 4.2示例 ###

示例1：

` int x = 10 ; int y = 10 ; String str1 = new String( "abc" ); String str2 = new String( "abc" ); System.out.println(x == y); // true System.out.println(str1 == str2); // false System.out.println(str1.equals(str2)); // true 复制代码`

示例2：

` String str3 = "abc" ; String str4 = "abc" ; System.out.println(str3 == str4); // true 复制代码`
> 
> 
> 
> 
> str3与str4想等的原因是用到了内存中的常量池，当运行到str3创建对象时，如果常量池中没有，就在常量池中创建一个对象"abc",第二次创建的时候，就直接使用，所以两次创建的对象其实是同一个对象，它们的地址值相等。
> 
> 
> 

示例3：

先定义学生Student类

` package com.zwwhnly.springbootdemo; public class Student { private int age; public Student ( int age) { this.age = age; } } 复制代码`

然后创建两个Student实例来比较：

` Student student1 = new Student( 23 ); Student student2 = new Student( 23 ); System.out.println(student1.equals(student2)); // false 复制代码`

此时equals方法调用的是基类Object类的equals()方法，也就是==比较，所以返回false。

然后我们重写下equals()方法，只要两个学生的年龄相同，就认为是同一个学生：

` package com.zwwhnly.springbootdemo; public class Student { private int age; public Student ( int age) { this.age = age; } public boolean equals (Object obj) { Student student = (Student) obj; return this.age == student.age; } } 复制代码`

此时再比较刚刚的两个实例，就返回true：

` Student student1 = new Student( 23 ); Student student2 = new Student( 23 ); System.out.println(student1.equals(student2)); // true 复制代码`

## 5.String,StringBuilder,StringBuffer的区别 ##

### 5.1区别讲解 ###

1)运行速度

运行速度快慢顺序为：StringBuilder > StringBuffer > String

String最慢的原因：

String为字符串常量，而StringBuilder和StringBuffer均为字符串变量，即String对象一旦创建之后该对象是不可以更改的，但后两者的对象是变量，是可以更改的。

2)线程安全

在线程安全上，StringBuilder是线程不安全的，而StringBuffer是线程安全的(很多方法带有synchronized关键字)。

3)使用场景

String：适用于少量的字符串操作的情况。

StringBuilder：适用于单线程下在字符缓冲区进行大量操作的情况。

StringBuffer：适用于多线程下在字符缓冲区进行大量操作的情况。

### 5.2示例 ###

以拼接10000次字符串为例，我们看下三者各自需要的时间：

` String str = "" ; long startTime = System.currentTimeMillis(); for ( int i = 0 ; i < 10000 ; i++) { str = str + i; } long endTime = System.currentTimeMillis(); long time = endTime - startTime; System.out.println( "String消耗时间：" + time); StringBuilder builder = new StringBuilder( "" ); startTime = System.currentTimeMillis(); for ( int j = 0 ; j < 10000 ; j++) { builder.append(j); } endTime = System.currentTimeMillis(); time = endTime - startTime; System.out.println( "StringBuilder消耗时间：" + time); StringBuffer buffer = new StringBuffer( "" ); startTime = System.currentTimeMillis(); for ( int k = 0 ; k < 10000 ; k++) { buffer.append(k); } endTime = System.currentTimeMillis(); time = endTime - startTime; System.out.println( "StringBuffer消耗时间：" + time); 复制代码`

运行结果：

> 
> 
> 
> String消耗时间：258
> 
> 

> 
> 
> 
> StringBuilder消耗时间：0
> 
> 

> 
> 
> 
> StringBuffer消耗时间：1
> 
> 

也验证了上面所说的StringBuilder > StringBuffer > String。

## 6.装箱和拆箱 ##

### 6.1什么是装箱？什么是拆箱？ ###

装箱：自动将基本数据类型转换为包装器类型。

拆箱：自动将包装器类型转换为基本数据类型。

` Integer i = 10 ; // 装箱 int j = i; // 拆箱 复制代码`

### 6.2 装箱和拆箱是如何实现的？ ###

装箱过程是通过调用包装器的valueOf方法实现的，而拆箱过程是通过调用包装器实例的 xxxValue方法实现的。（xxx代表对应的基本数据类型）。

怎么证明这个结论呢，我们新建个类Main,在主方法中添加如下代码：

` package com.zwwhnly.springbootdemo; public class Main { public static void main (String[] args) { Integer i = 100 ; int j = i; } } 复制代码`

然后打开cmd窗口，切换到Main类所在路径，执行命令：javac Main.java，会发现该目录会生成一个Main.class文件，用IDEA打开，会发现编译后的代码如下：

` // // Source code recreated from a .class file by IntelliJ IDEA // (powered by Fernflower decompiler) // package com.zwwhnly.springbootdemo; public class Main { public Main () { } public static void main (String[] var0) { Integer var1 = Integer.valueOf( 100 ); int var2 = var1.intValue(); } } 复制代码`

### 6.3示例 ###

示例1：

` Double i1 = 100.0 ; Double i2 = 100.0 ; Double i3 = 200.0 ; Double i4 = 200.0 ; System.out.println(i1==i2); System.out.println(i3==i4); 复制代码`

输出结果：

> 
> 
> 
> false
> 
> 
> 
> false
> 
> 

为什么都返回false呢，我们看下Double.valueOf()方法，就知晓了：

` private final double value; public Double ( double value) { this.value = value; } public static Double valueOf ( double d) { return new Double(d); } 复制代码`

示例2：

` Boolean i1 = false ; Boolean i2 = false ; Boolean i3 = true ; Boolean i4 = true ; System.out.println(i1==i2); System.out.println(i3==i4); 复制代码`

输出结果：

> 
> 
> 
> true
> 
> 
> 
> true
> 
> 

为什么都返回true呢，我们看下Boolean.valueOf()方法，就知晓了：

` public static final Boolean TRUE = new Boolean( true ); public static final Boolean FALSE = new Boolean( false ); public static Boolean valueOf ( boolean b) { return (b ? TRUE : FALSE); } 复制代码`

## 7.Java中的值传递和引用传递 ##

### 7.1基本概念 ###

值传递：传递对象的一个副本，即使副本被改变，也不会影响源对象，因为值传递的时候，实际上是将实参的值复制一份给形参。

引用传递：传递的并不是实际的对象，而是对象的引用，外部对引用对象的改变也会反映到源对象上，因为引用传递的时候，实际上是将实参的地址值复制一份给形参。

**说明：对象传递（数组、类、接口）是引用传递，原始类型数据（整形、浮点型、字符型、布尔型）传递是值传递。**

### 7.2示例 ###

示例1(值传递)：

` package com.zwwhnly.springbootdemo; public class ArrayListDemo { public static void main (String[] args) { int num1 = 10 ; int num2 = 20 ; swap(num1, num2); System.out.println( "num1 = " + num1); System.out.println( "num2 = " + num2); } public static void swap ( int a, int b) { int temp = a; a = b; b = temp; System.out.println( "a = " + a); System.out.println( "b = " + b); } } 复制代码`

运行结果：

> 
> 
> 
> a = 20
> 
> 
> 
> b = 10
> 
> 
> 
> num1 = 10
> 
> 
> 
> num2 = 20
> 
> 

> 
> 
> 
> 虽然在swap()方法中a,b的值做了交换，但是主方法中num1，num2的值并未改变。
> 
> 

示例2(引用类型传递)：

` package com.zwwhnly.springbootdemo; public class ArrayListDemo { public static void main (String[] args) { int [] arr = { 1 , 2 , 3 , 4 , 5 }; change(arr); System.out.println(arr[ 0 ]); } public static void change ( int [] array) { System.out.println(array[ 0 ]); array[ 0 ] = 0 ; } } 复制代码`

运行结果：

> 
> 
> 
> 1
> 
> 
> 
> 0
> 
> 

> 
> 
> 
> 在change()方法中将数组的第一个元素改为0，主方法中数组的第一个元素也跟着变为0。
> 
> 

示例3(StringBuffer类型)：

` package com.zwwhnly.springbootdemo; public class ArrayListDemo { public static void main (String[] args) { StringBuffer stringBuffer = new StringBuffer( "博客园：周伟伟的博客" ); System.out.println(stringBuffer); changeStringBuffer(stringBuffer); System.out.println(stringBuffer); } public static void changeStringBuffer (StringBuffer stringBuffer) { stringBuffer = new StringBuffer( "掘金：周伟伟的博客" ); stringBuffer.append( ",欢迎大家关注" ); } } 复制代码`

运行结果：

> 
> 
> 
> 博客园：周伟伟的博客
> 
> 
> 
> 博客园：周伟伟的博客
> 
> 

也许你会认为第2次应该输出“掘金：周伟伟的博客,欢迎大家关注”，怎么输出的还是原来的值呢，那是因为在changeStringBuffer中，又new了一个StringBuffer对象，此时stringBuffer对象指向的内存地址已经改变，所以主方法中的stringBuffer变量未受到影响。

如果修改changeStringBuffer()方法的代码为：

` public static void changeStringBuffer (StringBuffer stringBuffer) { stringBuffer.append( ",欢迎大家关注" ); } 复制代码`

则运行结果变为了：

> 
> 
> 
> 博客园：周伟伟的博客
> 
> 
> 
> 博客园：周伟伟的博客,欢迎大家关注
> 
> 

示例4(String类型)：

` package com.zwwhnly.springbootdemo; public class ArrayListDemo { public static void main (String[] args) { String str = new String( "博客园：周伟伟的博客" ); System.out.println(str); changeString(str); System.out.println(str); } public static void changeString (String string) { //string = "掘金：周伟伟的博客"; string = new String( "掘金：周伟伟的博客" ); } } 复制代码`

运行结果：

> 
> 
> 
> 博客园：周伟伟的博客
> 
> 
> 
> 博客园：周伟伟的博客
> 
> 

> 
> 
> 
> 在changeString()方法中不管用 ` string = "掘金：周伟伟的博客";` 还是 ` string = new
> String("掘金：周伟伟的博客");` ，主方法中的str变量都不会受影响，也验证了String创建之后是不可变更的。
> 
> 

示例5(自定义类型)：

` package com.zwwhnly.springbootdemo; public class Person { private String name; public String getName () { return name; } public void setName (String name) { this.name = name; } public Person (String name) { this.name = name; } } 复制代码` ` package com.zwwhnly.springbootdemo; public class ArrayListDemo { public static void main (String[] args) { Person person = new Person( "zhangsan" ); System.out.println(person.getName()); changePerson(person); System.out.println(person.getName()); } public static void changePerson (Person p) { Person person = new Person( "lisi" ); p = person; } } 复制代码`

运行结果：

> 
> 
> 
> zhangsan
> 
> 
> 
> zhangsan
> 
> 

修改changePerson()方法代码为：

` public static void changePerson (Person p) { p.setName( "lisi" ); } 复制代码`

则运行结果为：

> 
> 
> 
> zhangsan
> 
> 
> 
> lisi
> 
> 

## 8.参考链接 ##

[关于==和equals的区别和联系，面试这么回答就可以]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_27471405%2Farticle%2Fdetails%2F81010094 )

[Java中==号与equals()方法的区别]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2FStriverLi%2Farticle%2Fdetails%2F52997927 )

[Java中的String，StringBuilder，StringBuffer三者的区别]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fsu-feng%2Fp%2F6659064.html )

[深入剖析Java中的装箱和拆箱]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3780005.html )

[Integer、new Integer() 和 int 比较的面试题]( https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2Fcxxjohnson%2Fp%2F10504840.html )

[java面试题之int和Integer的区别]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fguodongdidi%2Fp%2F6953217.html )

[最最最常见的Java面试题总结-第一周]( https://juejin.im/post/5b691ae06fb9a04f87522cc3 )