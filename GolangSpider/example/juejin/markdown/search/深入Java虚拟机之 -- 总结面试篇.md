# 深入Java虚拟机之 -- 总结面试篇 #

> 
> 
> 
> **本文已授权玉刚说公众号**
> 
> 

系列文章：

> 
> 
> 
> [深入Java虚拟机之 -- 总结面试篇](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89239290
> )
> 
> 

> 
> 
> 
> [深入Java虚拟机之 --- JVM的爱恨情仇](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89239290
> )
> 
> 

> 
> 
> 
> [JAVA 垃圾回收机制(一) --- 对象回收与算法初识](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89281473
> )
> 
> 

> 
> 
> 
> [JAVA 垃圾回收机制(二) --- GC回收具体实现](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89284499
> )
> 
> 

> 
> 
> 
> [深入Java虚拟机之 -- 类文件结构(字节码)](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89312932
> )
> 
> 

> 
> 
> 
> [深入Java虚拟机之 -- 类加载机制](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89314894
> )
> 
> 

在学习 JVM 相关知识，怎么让自己有动力看下去，且有思考性呢？笔者认为，开头用一些常用的面试题，来引入读者的兴趣比较好，这样才会有看下去的东西，所以，该篇文章会以面试+总结的方式，希望读者能先思考写出答案，再查看相关知识。

# 一、JVM常见面试题 #

[深入Java虚拟机之 -- 总结面试篇]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89239290 )

* **介绍下 Java 内存区域**
* **Java 对象的创建过程**
* **对象的访问定位有几种**
* **String、StringBuilder、StringBuffer 有什么不同？**

这是一些常见的面试，很多人都看到网上的标准答案，但你知道为什么吗？

### 1.1 介绍下 Java 内存区域 ###

首先看第一个，Java的内存区域，可以看一张编译图：

![](https://user-gold-cdn.xitu.io/2019/5/7/16a8fc166aada1ac?imageView2/0/w/1280/h/960/ignore-error/1) 可以看到Java 的内存区域就是框框里的东西，每一步的大概意思如下，具体细节参考 [深入Java虚拟机之 --- JVM的爱恨情仇]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89239290 ) ： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a7b10d6cbe?imageView2/0/w/1280/h/960/ignore-error/1) 总结，建议读者学习之后，能自己默写这些方法并指导每一步的意思；

### 1.2 Java 对象的创建过程 ###

Java 对象的创建共分为5步，如下图：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a7b12b724d?imageView2/0/w/1280/h/960/ignore-error/1) 然后明白每个步骤做了哪些即可，如下： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a7b155ea95?imageView2/0/w/1280/h/960/ignore-error/1)

## 1.3、对象的访问定位有几种 ##

有两种方式：句柄和直接指针； 创建对象是为了使用对象，虚拟机需要通过栈中的 reference 来获取堆上的对象。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a7b84e2ee4?imageView2/0/w/1280/h/960/ignore-error/1) 优缺点: 使用句柄好处是，当对象发生改变，只需要移动句柄的实例数据指针即可，而直接指针就是速度快。

## 1.4 String、StringBuilder、StringBuffer 有什么不同 ##

参考答案是： String 是用 final 修饰的类，由于它的不可变性，类似拼接、裁剪字符串等，都会产生新的对象。 StringBuffer 解决上面拼接对象而提供一个类，可以通过 append等方法拼接，是线程安全的，由于线程安全，效率也下降 StringBuilder 跟StringBuffer 差不多，只是去掉了线程安全，所以优先使用 StringBuilder

说说String 为什么会产生新的对象？比如 String a = "1" String b = a + "2"，当执行这条指令时，会在常量池中产生一个对象指向a，而创建b时也会重新在常量池中生成b的对象；多次创建容易触发 GC，这也是为什么不建议使用 String 类去拼接的问题。

# 二、Java 回收机制常见面试题 #

[深入Java虚拟机之 --- JVM的爱恨情仇]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89239290 ) [JAVA 垃圾回收机制(二) --- GC回收具体实现]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89284499 )

* **简单的介绍一下强引用、软引用、弱引用、虚引用（虚引用与软引用和弱引用的区别、使用软引用能带来的好处）**
* **谈谈final、finally、finalize 有什么不同**
* **方法区会回收资源吗？**
* **垃圾回收有哪些算法，各自的特点？**

## 2.1 简单的介绍一下强引用、软引用、弱引用、虚引用 ##

首先，在讲解这几个引用之前，先明白虚拟机为什么会由这些引用的说明；我们都知道，对象需要回收，那怎么去判断哪些对象需要回收呢？这就需要一些判断来确定哪些对象是需要回收的，一般有以下几种方法：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a7b94e9334?imageView2/0/w/1280/h/960/ignore-error/1) 无论是 引用计算算法还是可达性分析算法，都是涉及到对象的引用问题，所以，在 JDK1.2 之后，又分为以下几类引用： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a7fde0363f?imageView2/0/w/1280/h/960/ignore-error/1) 通过上面的介绍，知道了" **引用** "是什么关系，这对理解各种引用还是很有必要的，那么使用 软引用的好处也在那里了； 建议一些内存消耗较大的使用软引用，比如 webview。。

## 2.2 谈谈final、finally、finalize 有什么不同 ##

final 和finally 比较好理解。首先 final 用来修饰的对象不可变；finally 则是保证重点代码一定要被执行的一种机制，一般用于 try - catch-finally 语句中。 但finalize 是什么东西呢？在解释标准代码之前，又得回到GC算法中了。 首先，finalize 是 Object 的一个方法，用来特定资源的回收。 上面说到，当 GC Roots 不可达时，认为对象已经不再使用了，但是对象并非是非"死"不可，当 GC Roots 不可达时，系统首先会先判断 对象的 finalize 是否执行，不执行则直接回收；如果可以执行，则放在队列中，由finalize线程去执行它，如果有其他对象关联时，则判断对象不可回收，否则对象回收，finalize 执行一次，如下图：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a813b82ce6?imageView2/0/w/1280/h/960/ignore-error/1) **由于它的不确定性，在 JDK9时，已经标注为deprecated** ，但不影响我们对它的理解。

## 2.3 方法区会回收资源吗？ ##

虽说 Java 堆 可以回收70%~95%的空间，但方法区同样可以回收一些资源，方法区主要回收两个部分 **废弃常量** 和 **无用的类** 。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a829fb9516?imageView2/0/w/1280/h/960/ignore-error/1) 所以，当发生 GC 时，非常常量和无用类是可以被回收，当然这里也是说"可以"，是否像对象一样被回收，还需要对虚拟机的参数配置，这里就不细说了。

## 2.4 垃圾回收有哪些算法，各自的特点？ ##

对象的回收，基于上面讲到的，GC Roots不可达，且判断可以回收。衍生的算法如下图(建议能默认每种算法的理解)：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a80a82af69?imageView2/0/w/1280/h/960/ignore-error/1) 其中，基础是 标记-清除是基础，接下来都是在它的基础上改进，分代算法是主流 Java 虚拟机的主要算法； 其中各个算法特点如下，详细介绍看 [JAVA 垃圾回收机制(一) --- 对象回收与算法初识]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89281473 ) 第四节，垃圾回收篇。 ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a81cab1770?imageView2/0/w/1280/h/960/ignore-error/1) 关系新生代和老年代的问题，参考： [JAVA 垃圾回收机制(二) --- GC回收具体实现]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89284499 )

# 三、类加载的问题 #

[深入Java虚拟机之 -- 类文件结构(字节码)]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89312932 ) [深入Java虚拟机之 -- 类加载机制]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89314894 )

* **类加载过程**
* **写出下列代码打印信息，若将改成System.out.println(Child.c_value);改为System.out.println(Child.value); 如何?**

` public class Parent{ static { System.out.println( "Parent" ); } public static int value = 123; } public class Child extends Parent{ static { System.out.println( "Child" ); } public static int c_value = 123; } //mian 中执行 public static void main(String[] args) { System.out.println(Child.c_value); } 复制代码`

* **说说你对类加载器的理解**
* **什么是双亲委派模型**

## 3.1 类加载的过程 ##

类加载的过程如下图所示(建议能默认每个步骤的理解)：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a81925616b?imageView2/0/w/1280/h/960/ignore-error/1) 也可以成为 加载-连接-初始化 这种叫法。 其中， **加载、验证、准备、初始化和卸载** 的顺序是固定的，而解析则不一定，因为Java是动态语言，它可以在运行时解析，即初始化之后。该阶段解析如下： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a838beebcb?imageView2/0/w/1280/h/960/ignore-error/1)

## 3.2 写出下列代码打印信息，若将改成System.out.println(Child.c_value);改为System.out.println(Child.value); 如何? ##

` public class Parent{ static { System.out.println( "Parent" ); } public static int value = 123; } public class Child extends Parent{ static { System.out.println( "Child" ); } public static int c_value = 123; } //mian 中执行 public static void main(String[] args) { System.out.println(Child.c_value); } 复制代码`

打印信息如：

` Parent Child 123 复制代码`

改为System.out.println(Child.value)时：

` Parent 123 复制代码`

具体看： [深入Java虚拟机之 -- 类加载机制]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011418943%2Farticle%2Fdetails%2F89314894 ) **扩展** ：

` class Parent{ public static int value = 1; static { value = 2; } } class Child extends Parent{ public static int B = value ; } public static void main(String[] args) { System.out.println(Child.B); } 复制代码`

输出什么？

## 3.3 说说你对类加载器的理解 ##

从上面我们知道，类在加载的时候，就是通过一个全限定名去加载这个类的二进制字节流，这个是系统自动完成的。这个动作如果从外部去做，以便于我们去获取所需的类，则我们成为 **类加载器** 。比如通过一个路径获取到一个 class 字节码，然后通过反射，拿到相应的信息。

## 3.4 什么是双亲委派模型 ##

它的工程流程是： 当一个类加载器收到类加载的请求，它首先不会自己去尝试加载这个类，而是委派给她的父类加载器去完成，每一个层次的类加载器都是如此，因此所有的加载器都会传递到父加载器中；只有父加载器无法完成时，子加载器才会尝试自己去加载，它的模型如下：

![类加载双亲委派模型](https://user-gold-cdn.xitu.io/2019/5/6/16a8b0a83b327d7b?imageView2/0/w/1280/h/960/ignore-error/1)