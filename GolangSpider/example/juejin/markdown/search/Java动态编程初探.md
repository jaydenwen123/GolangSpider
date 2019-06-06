# Java动态编程初探 #

> 
> 
> 
> 作者简介
> 
> 
> 
> 传恒，一个喜欢摄影和旅游的软件工程师，先后从事饿了么物流蜂鸟自配送和蜂鸟众包的开发，现在转战 Java，目前负责物流策略组分流相关业务的开发。
> 
> 

# 什么是动态编程 #

动态编程是相对于静态编程而言的，平时我们讨论比较多的静态编程语言例如Java， 与动态编程语言例如JavaScript相比，二者有什么明显的区别呢？ 简单的说就是在静态编程中，类型检查是在编译时完成的，而动态编程中类型检查是在运行时完成的， 所谓动态编程就是绕过编译过程在运行时进行操作的技术。

# 动态编程使用场景 #

* 通过配置生成代码，减少重复编码，降低维护成本。
* AOP的一种实现方式，方便实现性能监控和分析，日志，事务，权限校验等。
* 实现新语言的语义，例如Groovy使用ASM生成字节码。
* 单元测试中动态mock测试依赖。

# 在Java中有如下几种方式实现动态编程： #

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c8aeb8e5b1a3?imageView2/0/w/1280/h/960/ignore-error/1)

## 反射 ##

我们常用到的动态特性主要是反射，在运行时查找对象的属性和方法，修改作用域，通过方法名称调用方法等。在线的应用不建议频繁使用反射，因为反射的性能开销较大。

## 动态代理 ##

在java的java.lang.reflect包下提供了一个Proxy类和一个InvocationHandler接口，通过这个类和这个接口可以生成JDK动态代理类和动态代理对象。

## 动态编译 ##

动态编译是从Java 6开始支持的，主要是通过一个JavaCompiler接口来完成的。通过这种方式我们可以直接编译一个已经存在的java文件，也可以在内存中动态生成Java代码，动态编译执行。

## 调用Java Script引擎 ##

Java 6加入了对Script(JSR223)的支持。这是一个脚本框架，提供了让脚本语言来访问Java内部的方法。你可以在运行的时候找到脚本引擎，然后调用这个引擎去执行脚本，这个脚本API允许你为脚本语言提供Java支持。

## 动态生成字节码 ##

操作java字节码的工具有BECL/ASM/CGLIB/Javassist，其中有两个比较流行的，一个是ASM，一个是Javassist。 ASM直接操作字节码指令，执行效率高，要求使用者掌握Java类字节码文件格式及指令，对使用者的要求比较高。 Javassist提供了更高级的API，执行效率相对较差，但无需掌握字节码指令的知识，对使用者要求较低，所以接下来我们重点讲讲Javassist。

# Javassist #

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c8a512c45978?imageView2/0/w/1280/h/960/ignore-error/1) Javassist是一个开源的分析、编辑和创建Java字节码的类库。 它是由东京工业大学的数学和计算机科学系的 Shigeru Chiba (千叶滋) 所创建的，目前已经加入到开放源代码JBoss应用服务器项目，JBoss通过使用Javassist对字节码进行操作，实现动态AOP框架。

Javassist(Java Programming Assistant) 使对Java字节码的操作变得简单，它使Java程序能够在运行时定义新类，并且可以在JVM加载时修改类文件。 与其它类似的字节码编辑器不同，它提供两个级别的API：源级别和字节码级别。 如果用户使用源级别API，他们可以在不知道Java字节码规范的情况下编辑类文件。整个API仅使用Java语言的词汇表进行设计，你甚至可以使用Java源代码的方式插入字节码。 另外，用户也可以使用字节码级别的API去直接编辑类文件。

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c8b40a3c2109?imageView2/0/w/1280/h/960/ignore-error/1)

` // ClassPool 是 CtClass 对象的容器，存储着CtClass的Hash表。它按需读取类文件来构造CtClass对象，并且保存CtClass对象以便之后使用 ClassPool classPool = ClassPool.getDefault(); // CtClass 表示一个class文件，一个 GtClass(compile-time class)对象用来处理一个class文件，下面是从classpath中查找该类 CtClass ctClass = classPool.get( "test.config.ConfigHandle" ); // 通知编辑器去寻找对应的包 classPool.importPackage( "org.mockito.Mockito" ); classPool.importPackage( "test.adapter.ext.IDowngrade" ); classPool.importPackage( "test.utils.property.IProperties" ); // 使用removeField() removeMethod() 去删除对应的属性和方法 ctClass.removeField(ctClass.getDeclaredField( "serviceHandle" )); ctClass.removeField(ctClass.getDeclaredField( "switchHandle" )); ctClass.removeField(ctClass.getDeclaredField( "configHandle" )); // CtMethod 和 CtConstructor 提供了 setBody() 方法去修改方法体 CtConstructor ctConstructor = ctClass.getDeclaredConstructors()[ 0 ]; ctConstructor.setBody( "{this.mySwitch = Mockito.mock(IDowngrade.class);\n" + " this.myConfig = Mockito.mock(IProperties.class);}" ); // toClass() 请求当前线程的 ClassLoader 去加载 CtClass 所代表的类文件 ctClass.toClass(); //输出成二进制格式 //byte[] b = ctClass.toBytecode(); //输出class文件到目录中 //ctClass.writeFile("/tmp"); 复制代码`

ClassPool是CtClass对象的容器，因为编译器在编译引用CtClass代表的Java类的源代码时，可能会引用CtClass对象，所以一旦一个CtClass被创建，它就被保存在ClassPool中。

如果事先知道要修改哪些类，修改类的最简单方法如下：

* * 调用 ClassPool.get() 获取 CtClass 对象

* * 修改对象

* * 调用 CtClass 对象的 writeFile() 或者 toBytecode() 获得修改过的类文件。

如果需要定义一个新类，只需要

` ClassPool pool = ClassPool.getDefault(); CtClass cc = pool.makeClass( "HelloWorld" ); 复制代码`

## 冻结classes ##

如果一个 CtClass 对象通过 writeFile(), toClass(), toBytecode()被转换成一个类文件，该CtClass对象会被冻结起来，不允许再修改，因为一个类只能被JVM加载一次。

` CtClasss cc = ...; : cc.writeFile(); cc.defrost(); cc.setSuperclass(...); // 类已经被解冻 复制代码`

## Class 搜索路径： ##

通过 ClassPool.getDefault() 获取的ClassPool默认使用JVM的类搜索路径。如果程序运行在JBoss或者Tomcat等Web服务器上，ClassPool可能无法找到用户自己定义的类，因为这种Web服务器使用多个类加载器作为系统类加载器。在这种情况下，ClassPool必须添加额外的类搜索路径。

` pool.insertClassPath( new ClassClassPath( this.getClass())); // 当前的类使用的类路径，注册到类搜索路径 pool.insertClassPath( "/usr/local/javalib" ); // 添加目录 /usr/local/javalib 到类搜索路径 ClassPath cp = new URLClassPath( "www.javassist.org" , 80 , "/java/" , "org.javassist." ); pool.insertClassPath(cp); // 注册URL到搜索路径 复制代码`

在Java中，多个类加载器是可以共存的。每个类加载器创建了自己的命名空间，不同的类加载器可以加载具有相同类名的不同类文件，被加载的类也会被视为不同的类。此功能使我们能够在单个JVM上面运行多个应用程序，即使这些程序包含具有相同名称的类。

注意，JVM不允许动态重新加载类，一旦类加载器加载了一个类，就不能再在运行时重新加载该类的其它版本。因此，在JVM加载类之后，就不能再更改该类的定义。 但是，JPDA（Java平台调试器架构）提供有限的重新加载类的能力，如果相同的类文件由两个不同的类加载器加载，则JVM内会创建两个具有相同名称但是定义的不同的类。由于两个类不相同，所以一个类的实例不能被分配给另一个类的变量，两个类之间的转换操作也会失败并且抛出一个ClassCastException异常。

## 总结 ##

Javassist比我们在本文中所讨论的功能要丰富得多，作为jboss的一个子项目，其主要的优点在于简单和快速，可以直接使用java编码的形式，而不需要了解虚拟机指令，就能动态改变类的结构，或者动态生成类。如果你不是很了解虚拟机指令，可以采用javassist。

## 参考文档： ##

* [www.javassist.org/tutorial/tu…]( https://link.juejin.im?target=https%3A%2F%2Fwww.javassist.org%2Ftutorial%2Ftutorial.html )
* [en.wikipedia.org/wiki/Javass…]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FJavassist )

阅读博客还不过瘾？

> 
> 
> 
> 欢迎大家扫二维码通过添加群助手，加入交流群，讨论和博客有关的技术问题，还可以和博主有更多互动
> 
> ![](https://user-gold-cdn.xitu.io/2018/12/26/167e9cc24048932b?imageView2/0/w/1280/h/960/ignore-error/1)
> 博客转载、线下活动及合作等问题请邮件至 shadowfly_zyl@hotmail.com 进行沟通
> 
> 
> 
>