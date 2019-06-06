# Java代码的编译与反编译那些事儿 #

GitHub 2.5k Star 的 [Java工程师成神之路]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhollischuang%2FtoBeTopJavaer ) ，不来了解一下吗?

GitHub 2.5k Star 的 [Java工程师成神之路]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhollischuang%2FtoBeTopJavaer ) ，真的不来了解一下吗?

GitHub 2.5k Star 的 [Java工程师成神之路]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhollischuang%2FtoBeTopJavaer ) ，真的确定不来了解一下吗?

### 编程语言 ###

在介绍编译和反编译之前，我们先来简单介绍下编程语言（Programming Language）。编程语言（Programming Language）分为低级语言（Low-level Language）和高级语言（High-level Language）。

机器语言（Machine Language）和汇编语言（Assembly Language）属于低级语言，直接用计算机指令编写程序。

而C、C++、Java、Python等属于高级语言，用语句（Statement）编写程序，语句是计算机指令的抽象表示。

举个例子，同样一个语句用C语言、汇编语言和机器语言分别表示如下：

![](https://user-gold-cdn.xitu.io/2019/5/9/16a9a61bbc074e85?imageView2/0/w/1280/h/960/ignore-error/1)

计算机只能对数字做运算，符号、声音、图像在计算机内部都要用数字表示，指令也不例外，上表中的机器语言完全由十六进制数字组成。最早的程序员都是直接用机器语言编程，但是很麻烦，需要查大量的表格来确定每个数字表示什么意思，编写出来的程序很不直观，而且容易出错，于是有了汇编语言，把机器语言中一组一组的数字用助记符（Mnemonic）表示，直接用这些助记符写出汇编程序，然后让汇编器（Assembler）去查表把助记符替换成数字，也就把汇编语言翻译成了机器语言。

但是，汇编语言用起来同样比较复杂，后面，就衍生出了Java、C、C++等高级语言。

### 什么是编译 ###

上面提到语言有两种，一种低级语言，一种高级语言。可以这样简单的理解：低级语言是计算机认识的语言、高级语言是程序员认识的语言。

那么如何从高级语言转换成低级语言呢？这个过程其实就是编译。

从上面的例子还可以看出，C语言的语句和低级语言的指令之间不是简单的一一对应关系，一条 ` a=b+1` ;语句要翻译成三条汇编或机器指令，这个过程称为编译（Compile），由编译器（Compiler）来完成，显然编译器的功能比汇编器要复杂得多。用C语言编写的程序必须经过编译转成机器指令才能被计算机执行，编译需要花一些时间，这是用高级语言编程的一个缺点，然而更多的是优点。首先，用C语言编程更容易，写出来的代码更紧凑，可读性更强，出了错也更容易改正。

**将便于人编写、阅读、维护的高级计算机语言所写作的源代码程序，翻译为计算机能解读、运行的低阶机器语言的程序的过程就是编译。负责这一过程的处理的工具叫做编译器**

现在我们知道了什么是编译，也知道了什么是编译器。不同的语言都有自己的编译器，Java语言中负责编译的编译器是一个命令： ` javac`

> 
> 
> 
> javac是收录于JDK中的Java语言编译器。该工具可以将后缀名为.java的源文件编译为后缀名为.class的可以运行于Java虚拟机的字节码。
> 
> 
> 

**当我们写完一个 ` HelloWorld.java` 文件后，我们可以使用 ` javac HelloWorld.java` 命令来生成 ` HelloWorld.class` 文件，这个 ` class` 类型的文件是JVM可以识别的文件。通常我们认为这个过程叫做Java语言的编译。其实， ` class` 文件仍然不是机器能够识别的语言，因为机器只能识别机器语言，还需要JVM再将这种 ` class` 文件类型字节码转换成机器可以识别的机器语言。**

### 什么是反编译 ###

反编译的过程与编译刚好相反，就是将已编译好的编程语言还原到未编译的状态，也就是找出程序语言的源代码。就是将机器看得懂的语言转换成程序员可以看得懂的语言。Java语言中的反编译一般指将 ` class` 文件转换成 ` java` 文件。

有了反编译工具，我们可以做很多事情，最主要的功能就是有了反编译工具，我们就能读得懂Java编译器生成的字节码。如果你想问读懂字节码有啥用，那么我可以很负责任的告诉你，好处大大的。比如我的博文几篇典型的原理性文章，都是通过反编译工具得到反编译后的代码分析得到的。如深入理解多线程（一）——Synchronized的实现原理、深度分析Java的枚举类型—-枚举的线程安全性及序列化问题、Java中的Switch对整型、字符型、字符串型的具体实现细节、Java的类型擦除等。我最近在GitChat写了一篇关于Java语法糖的文章，其中大部分内容都用到反编译工具来洞悉语法糖背后的原理。

### Java反编译工具 ###

本文主要介绍3个Java的反编译工具： **javap** 、 **jad** 和 **cfr**

#### javap ####

` javap` 是jdk自带的一个工具，可以对代码反编译，也可以查看java编译器生成的字节码。 ` javap` 和其他两个反编译工具最大的区别是他生成的文件并不是 ` java` 文件，也不像其他两个工具生成代码那样更容易理解。拿一段简单的代码举例，如我们想分析Java 7中的 ` switch` 是如何支持 ` String` 的，我们先有以下可以编译通过的源代码：

` public class switchDemoString { public static void main(String[] args) { String str = "world"; switch (str) { case "hello": System.out.println("hello"); break; case "world": System.out.println("world"); break; default: break; } } } 复制代码`

执行以下两个命令：

` javac switchDemoString.java javap -c switchDemoString.class 复制代码`

生成代码如下：

` public class com.hollis.suguar.switchDemoString { public com.hollis.suguar.switchDemoString(); Code: 0: aload_0 1: invokespecial #1 // Method java/lang/Object."<init>":()V 4: return public static void main(java.lang.String[]); Code: 0: ldc #2 // String world 2: astore_1 3: aload_1 4: astore_2 5: iconst_m1 6: istore_3 7: aload_2 8: invokevirtual #3 // Method java/lang/String.hashCode:()I 11: lookupswitch { // 2 99162322: 36 113318802: 50 default: 61 } 36: aload_2 37: ldc #4 // String hello 39: invokevirtual #5 // Method java/lang/String.equals:(Ljava/lang/Object;)Z 42: ifeq 61 45: iconst_0 46: istore_3 47: goto 61 50: aload_2 51: ldc #2 // String world 53: invokevirtual #5 // Method java/lang/String.equals:(Ljava/lang/Object;)Z 56: ifeq 61 59: iconst_1 60: istore_3 61: iload_3 62: lookupswitch { // 2 0: 88 1: 99 default: 110 } 88: getstatic #6 // Field java/lang/System.out:Ljava/io/PrintStream; 91: ldc #4 // String hello 93: invokevirtual #7 // Method java/io/PrintStream.println:(Ljava/lang/String;)V 96: goto 110 99: getstatic #6 // Field java/lang/System.out:Ljava/io/PrintStream; 102: ldc #2 // String world 104: invokevirtual #7 // Method java/io/PrintStream.println:(Ljava/lang/String;)V 107: goto 110 110: return } 复制代码`

我个人的理解， ` javap` 并没有将字节码反编译成 ` java` 文件，而是生成了一种我们可以看得懂字节码。其实javap生成的文件仍然是字节码，只是程序员可以稍微看得懂一些。如果你对字节码有所掌握，还是可以看得懂以上的代码的。其实就是把String转成hashcode，然后进行比较。

个人认为，一般情况下我们会用到 ` javap` 命令的时候不多，一般只有在真的需要看字节码的时候才会用到。但是字节码中间暴露的东西是最全的，你肯定有机会用到，比如我在分析 ` synchronized` 的原理的时候就有是用到 ` javap` 。通过 ` javap` 生成的字节码，我发现 ` synchronized` 底层依赖了 ` ACC_SYNCHRONIZED` 标记和 ` monitorenter` 、 ` monitorexit` 两个指令来实现同步。

### jad ###

jad是一个比较不错的反编译工具，只要下载一个执行工具，就可以实现对 ` class` 文件的反编译了。还是上面的源代码，使用jad反编译后内容如下：

命令： ` jad switchDemoString.class`

` public class switchDemoString { public switchDemoString() { } public static void main(String args[]) { String str = "world"; String s; switch((s = str).hashCode()) { default: break; case 99162322: if(s.equals("hello")) System.out.println("hello"); break; case 113318802: if(s.equals("world")) System.out.println("world"); break; } } } 复制代码`

看，这个代码你肯定看的懂，因为这不就是标准的java的源代码么。这个就很清楚的可以看到原来 **字符串的switch是通过 ` equals()` 和 ` hashCode()` 方法来实现的** 。

但是，jad已经很久不更新了，在对Java7生成的字节码进行反编译时，偶尔会出现不支持的问题，在对Java 8的lambda表达式反编译时就彻底失败。

### CFR ###

jad很好用，但是无奈的是很久没更新了，所以只能用一款新的工具替代他，CFR是一个不错的选择，相比jad来说，他的语法可能会稍微复杂一些，但是好在他可以work。

如，我们使用cfr对刚刚的代码进行反编译。执行一下命令：

` java -jar cfr_0_125.jar switchDemoString.class --decodestringswitch false 复制代码`

得到以下代码：

` public class switchDemoString { public static void main(String[] arrstring) { String string; String string2 = string = "world"; int n = -1; switch (string2.hashCode()) { case 99162322: { if (!string2.equals("hello")) break; n = 0; break; } case 113318802: { if (!string2.equals("world")) break; n = 1; } } switch (n) { case 0: { System.out.println("hello"); break; } case 1: { System.out.println("world"); break; } } } } 复制代码`

通过这段代码也能得到字符串的switch是通过 ` equals()` 和 ` hashCode()` 方法来实现的结论。

相比Jad来说，CFR有很多参数，还是刚刚的代码，如果我们使用以下命令，输出结果就会不同：

` java -jar cfr_0_125.jar switchDemoString.class public class switchDemoString { public static void main(String[] arrstring) { String string; switch (string = "world") { case "hello": { System.out.println("hello"); break; } case "world": { System.out.println("world"); break; } } } } 复制代码`

所以 ` --decodestringswitch` 表示对于switch支持string的细节进行解码。类似的还有 ` --decodeenumswitch` 、 ` --decodefinally` 、 ` --decodelambdas` 等。在我的关于语法糖的文章中，我使用 ` --decodelambdas` 对lambda表达式警进行了反编译。 源码：

` public static void main(String... args) { List<String> strList = ImmutableList.of("Hollis", "公众号：Hollis", "博客：www.hollischuang.com"); strList.forEach( s -> { System.out.println(s); } ); } 复制代码`

` java -jar cfr_0_125.jar lambdaDemo.class --decodelambdas false` 反编译后代码：

` public static /* varargs */ void main(String ... args) { ImmutableList strList = ImmutableList.of((Object)"Hollis", (Object)"\u516c\u4f17\u53f7\uff1aHollis", (Object)"\u535a\u5ba2\uff1awww.hollischuang.com"); strList.forEach((Consumer<String>)LambdaMetafactory.metafactory(null, null, null, (Ljava/lang/Object;)V, lambda$main$0(java.lang.String ), (Ljava/lang/String;)V)()); } private static /* synthetic */ void lambda$main$0(String s) { System.out.println(s); } 复制代码`

CFR还有很多其他参数，均用于不同场景，读者可以使用 ` java -jar cfr_0_125.jar --help` 进行了解。这里不逐一介绍了。

### 如何防止反编译 ###

由于我们有工具可以对 ` Class` 文件进行反编译，所以，对开发人员来说，如何保护Java程序就变成了一个非常重要的挑战。但是，魔高一尺、道高一丈。当然有对应的技术可以应对反编译咯。但是，这里还是要说明一点，和网络安全的防护一样，无论做出多少努力，其实都只是提高攻击者的成本而已。无法彻底防治。

典型的应对策略有以下几种：

* 隔离Java程序

* 让用户接触不到你的Class文件

* 对Class文件进行加密

* 提到破解难度

* 代码混淆

* 将代码转换成功能上等价，但是难于阅读和理解的形式

![](https://user-gold-cdn.xitu.io/2019/5/8/16a95c40250fdc90?imageView2/0/w/1280/h/960/ignore-error/1)