# java类加载机制 #

在看java和android的类加载机制，途中有一些疑惑，就先记下来。

### 一些概念的理解 ###

**jdk和jre是什么区别？** JDK就是Java Development Kit.简单的说JDK是面向开发人员使用的SDK，它提供了Java的开发环境和运行环境。SDK是Software Development Kit 一般指软件开发包，可以包括函数库、编译程序等。JRE是Java Runtime Enviroment是指Java的运行环境，是面向Java程序的使用者，而不是开发者。

**rt.jar、dt.jar、tools.jar是什么？** rt.jar这个文件是极为重要的一个文件，rt是runtime的缩写，即运行时的意思。是java程序在运行时必不可少的文件。里面包含了java程序员常用的包，如java.lang，java.util，java.io，java.net,java.applet等；dt.jar是关于运行环境的类库，主要是swing的包，你要用到swing时最好加上；tools.jar 是系统用来编译一个类的时候用到的，也就是javac的时候用到 。

**` lib/` 和 ` jre/lib` 的区别是什么？** JDK下的lib包括java开发环境的jar包，是给JDK用的，例如JDK下有一些工具，可能要用该目录中的文件。例如，编译器等；JDK下的JRE下的lib是开发环境中，运行时需要的jar包。最典型的就是导入的外部驱动jar包。

**什么是android dex文件** 明白什么是 Dex 文件之前，要先了解一下 JVM，Dalvik 和 ART。JVM 是 JAVA 虚拟机，用来运行 JAVA 字节码程序。Dalvik 是 Google 设计的用于 Android平台的运行时环境，适合移动环境下内存和处理器速度有限的系统。ART 即 Android Runtime，是 Google 为了替换 Dalvik 设计的新 Android 运行时环境，在Android 4.4推出。ART 比 Dalvik 的性能更好。Android 程序一般使用 Java 语言开发，但是 Dalvik 虚拟机并不支持直接执行 JAVA 字节码，所以会对编译生成的 .class 文件进行翻译、重构、解释、压缩等处理，这个处理过程是由 dx 进行处理，处理完成后生成的产物会以 .dex 结尾，称为 Dex 文件。Dex 文件格式是专为 Dalvik 设计的一种压缩格式。所以可以简单的理解为：Dex 文件是很多 .class 文件处理后的产物，最终可以在 Android 运行时环境执行。

### jvm中的类加载器 ###

#### 启动类加载器（Bootstrap Class Loader） ####

主要负责加载jdk中的核心类库，比如rt.jar和其它在jre/lib中的核心类。Bootstrap Class Loader是所有classloader的父加载器。它是有native代码实现的

> 
> 
> 
> It’s mainly responsible for loading JDK internal classes, typically rt.jar
> and other core libraries located in $JAVA_HOME/jre/lib directory.
> Additionally, Bootstrap class loader serves as a parent of all the other
> ClassLoader instances.
> 
> 
> 
> This bootstrap class loader is part of the core JVM and is written in
> native code. Different platforms might have different implementations of
> this particular class loader.
> 
> 

#### 扩展类加载器（Extension Class Loader） ####

扩展类加载器是由Sun的ExtClassLoader（sun.misc.Launcher$ExtClassLoader）实现的，它负责将 ` <JAVA_HOME >/lib/ext` 或者由系统变量 ` -Djava.ext.dir` 指定位置中的类库 加载到内存中。开发者可以直接使用标准扩展类加载器。

#### 系统类加载器（System Class Loader） ####

系统类加载器是由 Sun 的 AppClassLoader（sun.misc.Launcher$AppClassLoader）实现的，它负责将用户类路径(java -classpath或-Djava.class.path变量所指的目录，即当前类所在路径及其引用的第三方类库的路径，加载到内存中。开发者可以直接使用系统类加载器。

> 
> 
> 
> The system or application class loader, on the other hand, takes care of
> loading all the application level classes into the JVM. It loads files
> found in the classpath environment variable, -classpath or -cp command
> line option. Also, it’s a child of Extensions classloader.
> 
> 

> 
> 
> 
> The extension class loader is a child of the bootstrap class loader and
> takes care of loading the extensions of the standard core Java classes so
> that it’s available to all applications running on the platform.
> 
> 
> 
> Extension class loader loads from the JDK extensions directory, usually
> $JAVA_HOME/lib/ext directory or any other directory mentioned in the
> java.ext.dirs system property.
> 
> 

#### 举个例子 ####

` public void printClassLoaders () throws ClassNotFoundException { System.out.println( "Classloader of this class:" + PrintClassLoader.class.getClassLoader()); System.out.println( "Classloader of Logging:" + Logging.class.getClassLoader()); System.out.println( "Classloader of ArrayList:" + ArrayList.class.getClassLoader()); } 复制代码`

运行结果是

> 
> 
> 
> Class loader of this class:sun.misc.Launcher$AppClassLoader@18b4aac2
> 
> 
> 
> Class loader of Logging:sun.misc.Launcher$ExtClassLoader@3caeaf62
> 
> 
> 
> Class loader of ArrayList:null
> 
> 

### classloader的类的继承关系 ###

` classDiagram Object <|-- ClassLoader ClassLoader <|-- SecureClassLoader SecureClassLoader <|-- URLClassLoader URLClassLoader <|-- ExtClassLoader URLClassLoader <|-- AppClassLoader BootStrapClassLoader <-- ExtClassLoader ExtClassLoader <-- AppClassLoader AppClassLoader <-- ClassLoaderA AppClassLoader <-- ClassLoaderB 复制代码`

其中，ExtClassLoader属于Extension Class Loader，AppClassLoader属于System Class Loader。

### 双亲委派机制 ###

JVM在加载类时默认采用的是双亲委派机制。通俗的讲，就是某个特定的类加载器在接到加载类的请求时，首先将加载任务委托给父类加载器，依次递归 (本质上就是loadClass函数的递归调用)。因此，所有的加载请求最终都应该传送到顶层的启动类加载器中。如果父类加载器可以完成这个类加载请求，就成功返回；只有当父类加载器无法完成此加载请求时，子加载器才会尝试自己去加载。事实上，大多数情况下，越基础的类由越上层的加载器进行加载。

加载过程如下：

* 源 ClassLoader 先判断该 Class 是否已加载，如果已加载，则直接返回 Class，如果没有则委托给父类加载器。
* 父类加载器判断是否加载过该 Class，如果已加载，则直接返回 Class，如果没有则委托给祖父类加载器。
* 依此类推，直到始祖类加载器（引用类加载器）。
* 始祖类加载器判断是否加载过该 Class，如果已加载，则直接返回 Class，如果没有则尝试从其对应的类路径下寻找 class 字节码文件并载入。如果载入成功，则直接返回 Class，如果载入失败，则委托给始祖类加载器的子类加载器。
* 始祖类加载器的子类加载器尝试从其对应的类路径下寻找 class 字节码文件并载入。如果载入成功，则直接返回 Class，如果载入失败，则委托给始祖类加载器的孙类加载器。
* 依此类推，直到源 ClassLoader。
* 源 ClassLoader 尝试从其对应的类路径下寻找 class 字节码文件并载入。如果载入成功，则直接返回 Class，如果载入失败，源 ClassLoader 不会再委托其子类加载器，而是抛出异常。

### android的class loader ###

` classDiagram Object <|-- ClassLoader ClassLoader <|-- BaseDexClassLoader ClassLoader <|-- SecureClassLoader SecureClassLoader <|-- URLClassLoader BaseDexClassLoader <|-- PathClassLoader BaseDexClassLoader <|-- DexClassLoader BaseDexClassLoader <|-- InMemoryDexClassLoader ClassLoader <|-- BootClassLoader BootClassLoader <-- PathClassLoader 复制代码`

几个知识点

* 在 Android 中，App 安装到手机后，apk 里面的 class.dex 中的 class 均是通过 PathClassLoader 来加载的。
* 对比 PathClassLoader 只能加载已经安装应用的 dex 或 apk 文件，DexClassLoader 则没有此限制，可以从 SD 卡上加载包含 class.dex 的 .jar 和 .apk 文件，这也是插件化和热修复的基础，在不需要安装应用的情况下，完成需要使用的 dex 的加载。 ` A class loader that loads classes from .jar and .apk filescontaining a classes.dex entry. This can be used to execute code notinstalled as part of an application.`
* SecureClassLoader和URLClassLoader和JDK8中的是一样的
* InMemoryDexClassLoader是Android8.0新增的类加载器，继承自BaseDexClassLoader，用于加载内存中的dex文件。
* BootClassLoader是在Zygote进程的入口方法中创建的，PathClassLoader则是在Zygote进程创建SystemServer进程时创建的。

### 疑惑 ###

SystemServer进程中创建了PathClassLoader，Zygote进程中创建了BootClassLoader，因为app进程是靠Zygote进程fork出来的，那么app进程中的PathClassLoader是在哪里创建的？

于是带着这个问题又研究了一下android api-level 为28的源代码，把app中的PathClassLoader的创建过程用时序图描述出来。

![](https://user-gold-cdn.xitu.io/2019/4/5/169ed6f5fdaa5d93?imageView2/0/w/1280/h/960/ignore-error/1)

在ActivityThread中ApplicationThread的bindApplication方法中发送消息BIND_APPLICATION给H，H也是ActivityThread的内部类，它继承了Handler，然后H调用ActivityThread的handleBindApplicationf方法，然后ActivityThread处理该消息时经过层层调用，最后返回了PathClassLoader。

那么ActivityThread中的ApplicationThread的bindApplication是在什么时候被谁调用的呢？调用关系如下描述：

![](https://user-gold-cdn.xitu.io/2019/4/5/169ed6f88828a3e7?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到ActivityThread有个静态main方法，它经过层层调用并且经过IActivityManager这个IPC操作，调用到了AMS的一些方法，然后AMS再经过IApplicationThread这个IPC操作，调用到了ActivityThread的ApplicationThread的bindApplication方法。

IActivityManager对应的是AMS，IApplicationThread对应的是ActivityThread的ApplicationThread。

ActivityThread的这个静态main方法是Zygote fork出来app这个子进程后，子进程调用到的。

### 参考 ###

* [blog.csdn.net/justloveyou…]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fjustloveyou_%2Farticle%2Fdetails%2F72217806 )
* [www.baeldung.com/java-classl…]( https://link.juejin.im?target=https%3A%2F%2Fwww.baeldung.com%2Fjava-classloaders )
* [jaeger.itscoder.com/android/201…]( https://link.juejin.im?target=https%3A%2F%2Fjaeger.itscoder.com%2Fandroid%2F2016%2F08%2F27%2Fandroid-classloader.html )
* [juejin.im/post/5bf22b…]( https://juejin.im/post/5bf22bb5e51d454cdc56cbd5 )
* [www.jianshu.com/p/a1f40b39b…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fa1f40b39b3de )
* [www.jianshu.com/p/fbea00880…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Ffbea00880da1 )
* [www.zhihu.com/question/50…]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F50828920 )