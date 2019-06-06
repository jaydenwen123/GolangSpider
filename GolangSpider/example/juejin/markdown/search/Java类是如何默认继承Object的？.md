# Java类是如何默认继承Object的？ #

## 前言 ##

学过 ` Java` 的人都知道， ` Object` 是所有类的父类。但是你有没有这样的疑问，我并没有写 ` extends Object` ，它是怎么默认继承Object的呢？

那么今天我们就来看看像Java这种依赖于虚拟机的编程语言是怎样实现默认继承Object的，以及 ` Java编译器` 和 ` JVM` 到底是如何做的？

## 继承自Object验证 ##

首先我们来验证一下Object是不是所有类的父类，随便新建一个Java类，如下图：

![](https://user-gold-cdn.xitu.io/2019/4/1/169d875991c972a0?imageView2/0/w/1280/h/960/ignore-error/1) 从上面的代码可以看出，new MyClass()打点之后可以选择调用的方法有很多，我们定义的MyClass类里面只有一个main方法，那这些方法哪来的，显然是Object里声明的，故MyClass类的父类就是Object，因此，在MyClass中可以使用Object类的public或protected资源。

另外，当A类继承MyClass类时，通过打点也可以调到Object内的方法，这是继承的传递，好比Object是MyClass的“父亲”，MyClass是A类的“父亲”，Object是A类的“爷爷”，间接的继承了Object。

因此，Object是超类，是所有类的父类。

## 推测可能的原因 ##

要了解 ` Java类是如何默认继承Object的？` 的原因其实并不需要知道JVM的实现细节。只需了解一下对于这种虚拟机程序的基本原理即可。一般对于这种靠虚拟机运行的语言（如Java、C#等）会有两种方法处理默认继承问题。

### 编译器处理 ###

在编译源代码时，当一个类没有显式标明继承的父类时，编译器会为其指定一个默认的父类（一般为Object），而交给虚拟机处理这个类时，由于这个类已经有一个默认的父类了，因此，VM仍然会按照常规的方法像处理其他类一样来处理这个类。对于这种情况，从编译后的二进制角度来看，所有的类都会有一个父类（后面可以以此依据来验证）。

### JVM处理 ###

编译器仍然按照实际代码进行编译，并不会做额外的处理，即如果一个类没有显式地继承于其他类时，编译后的代码仍然没有父类。然后由虚拟机运行二进制代码时，当遇到没有父类的类时，就会自动将这个类看成是Object类的子类（一般这类语言的默认父类都是Object）。

## 验证结论 ##

从上面两种情况可以看出，第1种情况是在编译器上做的文章，也就是说，当没有父类时，由编译器在编译时自动为其指定一个父类。第2种情况是在虚拟机上做文章，也就是这个默认的父类是由虚拟机来添加的。

那么Java是属于哪一种情况呢？其实这个答案很好得出。只需要随便找一个反编译工具，将.class文件进行反编译即可得知编译器是如何编译的。

就以上面代码为例，如果是第1种情况，就算MyClass没有父类，但由于编译器已经为MyClass自动添加了一个Object父类，所以，在反编译后得到的源代码中的MyClass类将会继承Object类的。如果不是这种情况，那么就是第2种情况。

那么实际情况是什么样的呢？现在我们就将MyClass.class反编译看看到底如何。

**jd-gui反编：**

![](https://user-gold-cdn.xitu.io/2019/4/2/169dbb785db1ba24?imageView2/0/w/1280/h/960/ignore-error/1) **使用JDK自带的工具（javap）反编译**

CMD命令行下执行： ` javap MyClass>MyClass.txt`

![](https://user-gold-cdn.xitu.io/2019/4/2/169dbaf63b1195a3?imageView2/0/w/1280/h/960/ignore-error/1) 可以看出实际的反编译后的文件中并没有 ` extends Object` ，使用排除法，因此是第2情况。

这样来推导出的结论是第2种情况，但事实真的如此吗？为什么网上还有说反编译后的是有 ` extends Object` 字样？

**JDK版本问题？**

猜想是JDK版本的问题，于是把JDK版本切换到7，使用jd-gui和javap反编译，接果和使用JDK8反编译后的结果一样，也都没有 ` extends Object` 。

继续换版本，昨晚在宿舍准备到Oracle官网下载JDK 6，但是死活下不来，今早到公司后第一件事就是下载，很顺利，安装后把JDK版本切换到JDK 6。

仍然在CMD窗口执行 ` javap MyClass>MyClass.txt` ，得到的TXT文件内容如下：

![](https://user-gold-cdn.xitu.io/2019/4/2/169dbade548201b0?imageView2/0/w/1280/h/960/ignore-error/1) what？竟然有 ` extends Object` ，jd-gui反编译后的依然没有。 即，JDK 6之前使用javap反编译后的MyClass类显式的继承Object，JDK 7以后没有；jd-gui反编译后的不管JDK版本如何始终没有。我们以java自带的工具为准。

## 总结 ##

那么就是说JDK 6之前是 ` 编译器` 处理，JDK 7之后是 ` 虚拟机` 处理。

但是仔细想想我们在 ` 编辑器` 里（IDE）打点时就能列出Object类下的方法，此时还没轮到编译器和jvm，编辑器就已经知道MyClass类的父类是Object类了，这是因为编辑器为我们做了一些智能处理。

【end】

参考文献：

java中 创建一个新的类 怎么默认继承Object类的： [zhidao.baidu.com/question/12…]( https://link.juejin.im?target=https%3A%2F%2Fzhidao.baidu.com%2Fquestion%2F128868315.html )