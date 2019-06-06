# 【译】Java8官方教程：Java技术概述 #

> 
> 
> 
> 原文地址: [docs.oracle.com/javase/tuto…](
> https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Ftutorial%2FgetStarted%2Fintro%2Findex.html
> )
> 
> 

# 课程：Java技术概述 #

关于Java技术的讨论似乎无处不在，但它究竟是什么呢？下面几节将解释Java技术是怎样同时作为编程语言和平台的，并且提供这项技术能为你做些什么的概述。

* 

Java技术到底是什么？

* 

Java技术能做什么？

* 

Java技术将如何改变我们的生活？

# Java技术到底是什么？ #

Java技术既是一门编程语言，同时又是一个平台。

## Java编程语言 ##

Java编程语言是一门高级语言，可以用以下的所有流行词汇来描述它：
简单
面向对象
分布式
多线程
动态的
体系结构中立
可移植
高效
健壮
安全

前面的每个术语都在James Gosling和Henry McGilton撰写的白皮书-《The Java Language Environment》中进行了解释。
在Java编程语言中，所有的源代码都是用.java拓展名的纯文本文件编写的，这些源文件通过javac编译器编译成.class文件。.class文件中包含的不是与本地机器相关的机器码，而是可被Java虚拟机(Java VM)执行的字节码，Java启动工具使用Java虚拟机实例运行你的程序。

![](https://user-gold-cdn.xitu.io/2019/5/7/16a920bb02e507f9?imageView2/0/w/1280/h/960/ignore-error/1) 因为Java VM可以在不同的操作系统上使用，因此.class文件也可以在Microsoft Windows,Solaris OS, Linux, 或者 Mac OS上运行。某些虚拟机(如HotSpot)在运行时执行额外的步骤来提高应用程序的性能。其中包括寻找性能瓶颈、重编译(编译成机器码)热点代码。

![](https://user-gold-cdn.xitu.io/2019/5/7/16a92130add670d6?imageView2/0/w/1280/h/960/ignore-error/1)

## Java平台 ##

平台是程序运行的硬件或软件环境，我们已经提过一些流行的平台，例如:Microsoft Windows, Linux, Solaris OS, 和 Mac OS。大多数的平台可以描述为操作系统和底层硬件的组合，Java平台与大多数其他平台的不同之处在于:它是一个运行在其他基于硬件的平台之上的纯软件平台.

Java平台包含两个组件:
1、Java虚拟机
2、Java API(Application Programming Interface)
你已经对Java虚拟机有了一定了解；它是Java平台的基础，并可被移植到各种基于硬件的平台上。 API是大量现成的软件组件的集合，提供了许多有用的功能。相关的类和接口被分到不同的库；这些库称为包(package)。下一节 将突显API提供的一些功能。

![](https://user-gold-cdn.xitu.io/2019/5/7/16a921f3a440b0e5?imageView2/0/w/1280/h/960/ignore-error/1) 作为一个独立于具体平台的环境，Java平台可能比本地机器码要慢一些，但是随着编译器和虚拟机技术的进步使得性能接近于原生代码，并具有良好的可移植性。
术语"Java Virtual Machine"和"JVM"指的是Java平台中的Java虚拟机。

# Java技术能做什么？ #

Java技术提供一个功能强大的软件平台，Java平台的每个完整实现都提供了以下特性：

* 

开发工具：开发工具提供了编译、运行、监视、调试和注释应用程序所需的一切，作为一个新开发人员，你主要使用的工具将是javac编译器，java启动器，javadoc文档工具

* 

API：API提供了Java编程语言的核心功能。它提供了大量有用的类，可以在您自己的应用程序中使用。它涵盖了从基本对象、到网络和安全、到XML生成和数据库访问等所有方面，核心API非常庞大；要获得它所包含内容的概述，请参考 [Java Platform Standard Edition 8 Documentation]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F8%2Fdocs%2Findex.html )

* 

部署技术:JDK软件提供了标准的机制，比如Java Web Start软件和Java插件软件，用于将应用程序部署到最终用户

* 

用户界面工具包：JavaFX、Swing和Java 2D工具包使创建复杂的图形用户界面(GUI)成为可能

* 

集成库:诸如Java IDL API、JDBC API、Java命名和目录接口(JNDI) API、Java RMI、Java RMI-IIOP。

# Java技术将如何改变我们的生活 #

我们不能保证你通过学习Java编程语言从而拥有名望、财富、或者是一份工作。但是，与其他语言相比，它使得你的程序更好并且节省你的精力。我们相信Java技术能够帮助你完成以下的工作:

* 

简单易学：尽管Java编程语言是一种强大的面向对象语言，但它很容易学习，尤其是对于已经熟悉C或C++的程序员们来说。

* 

代码简洁：对程序指标(类数、方法数等)的比较表明:用Java编程语言编写的程序可能比用C++编写的相同程序小四倍

* 

代码优美：Java编程语言鼓励良好的编码实践，并且自动垃圾收集机制帮助您避免内存泄漏。它的面向对象、JavaBeans™组件体系结构和广泛的、易于扩展的API允许重用现有的、经过测试的代码并引入更少的bug

* 

快速开发：Java编程语言比c++简单，因此，在用它编写代码时，您的开发时间可能比c++快两倍您的程序也将需要更少的代码行。

* 

可移植性良好: 您可以通过避免使用其他语言编写的库来保持程序的可移植性。一次编写，到处运行:因为用Java编程语言编写的应用程序被编译成与机器无关的字节码，所以它们可以在任何Java平台上一致地运行。

* 

易发布性:使用Java Web Start软件，用户只需单击鼠标就可以启动应用程序。启动时的自动版本检查确保用户始终与软件的最新版本保持同步。如果有更新可用，Java Web Start软件将自动更新它们的安装