# Java 笔记《异常》 #

* 为什么需要异常？
* 程序中可能会出现那些错误和问题？
* 异常有哪些种类？
* 常见的异常处理方式有哪些?
* 使用异常机制的技巧有哪些？

## 1. 为什么需要异常？ ##

用户在遇到错误时会感觉到很不爽，如果一个用户在运行程序期间，由于程序的一些错误或者外部环境的影响造成了用户数据的丢失，用户就有可能不在使用这个程序了。为了避免此类事情发生，至少应该做到：

* 向用户报告错误
* 保存所有的工作结果
* 允许用户以妥善的形式退出程序
* 返回到一种安全状态，并能够让用户执行一些其他的操作

Java提供的异常捕获机制来改善这种情况。

> 
> 
> 
> **某个方法不能采用正常的途径完成它的任务，就可能通过另一个路径退出该方法。** 这种情况下，方法并不返回任何值，而是抛出（ ` throw` ）一个封装了错误信息的对象。这个方法会
> 立刻退出，并不返回任何值，调用这个方法的代码也将无法执行，而是异常处理机制开始搜索能够处理这种异常状况的异常处理器（exception
> handler）。
> 
> 

## 2.程序中可能会出现那些错误和问题？ ##

* 用户输入错
* 设备错误
* 物理限制
* 代码错img

## 3.异常有哪些种类？ ##

![Java异常层次结构](https://user-gold-cdn.xitu.io/2019/4/3/169e346364d99a5e?imageView2/0/w/1280/h/960/ignore-error/1)

### 3.1 ` Exception` ###

[JavaAPI]( https://link.juejin.im?target=http%3A%2F%2Fdocs.oracle.com%2Fjavase%2F8%2Fdocs%2Fapi%2F ) 里是这样描述的：

* extends Throwable

The class Exception and its subclasses are a form of Throwable that indicates conditions that a reasonable application might want to catch.

The class Exception and any subclasses that are not also subclasses of RuntimeException are **checked exceptions**. Checked exceptions **need to** be declared in a method or constructor's throws clause if they can be thrown by the execution of the method or constructor and propagate outside the method or constructor boundary.

### 3.2 ` RuntimeException` 派生出 ###

* ` NullPointerException` - 空指针引用异常。
* ` ClassCastException` - 类型强制转换异常。
* ` IllegalArgumentException` - 传递非法参数异常。
* ` ArithmeticException` - 算术运算异常
* ` ArrayStoreException` - 向数组中存放与声明类型不兼容对象异常
* ` IndexOutOfBoundsException` - 下标越界异常
* ` NegativeArraySizeException` - 创建一个大小为负数的数组错误异常
* ` NumberFormatException` - 数字格式异常
* ` SecurityException` - 安全异常
* ` UnsupportedOperationException` - 不支持的操作异常

[JavaAPI]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F8%2Fdocs%2Fapi%2F ) 里是这样描述的：

* extends Exception

RuntimeException is the superclass of those exceptions that can be thrown during the normal operation of the Java Virtual Machine.

RuntimeException and its subclasses are **unchecked exceptions**. Unchecked exceptions **do not need to** be declared in a method or constructor's throws clause if they can be thrown by the execution of the method or constructor and propagate outside the method or constructor boundary.

### 3.3 ` IOException` 有 ###

* 试图打开一个不存在的文件。
* 试图在文件尾部读数据。
* 试图根据给的的字符串查找 ` Class` 对象，而这个字符串表示的类并不存在。

Java把所有的 ` Error` ＆ ` RuntimeException` 称为 **未检查(unchecked) **异常，其他的异常称为** 已检查异常(checked)** ，编译器会检查是否未所以的已检查异常提供了异常处理器。

> 
> 
> 
> 如果出现了 ` RuntimeException` ，那么一定是我的问题
> 
> 

## 4.常见的异常处理方式有哪些? ##

### 4.1 声明异常 ###

如标准类库提供的 ` FileInputStream` 类的一个构造器的声明：

` public FileInputStream (String name) throws FileNotFoundException 复制代码`

### 4.2 抛出异常 ###

` String readData (Scanner in) throws EOFException { ... while (...) { if (!in.hasNext()) /*EOF encountered*/ { if (n < len) throw new EOFException(); } ... } return s; } 复制代码`

其中为了更加细致的描述这个异常， ` EOFException` 类提供了一个含义字符串类型参数的构造器：

` String gripe = "Conten - length: " + len + ", Received: " + n; throw new EOFException(gripe); 复制代码`

### 4.3 捕获异常 ###

` try { code more code } catch (Exception) { handler for this type } 复制代码`

有时候需要捕获多个异常：

` try { code that might throw exceptions } catch (FileNotFoundException | UnknowHostException) { emergency action for missing file and unkonw hosts } catch (IOException) { emergency action for all other I/O problems } 复制代码`
> 
> 
> 
> 
> 注意这里两个异常有相同的处理动作时可以合并，合并采用的 **逻辑或( ` |` )** 而不是短路或（ ` ||` ）
> 
> 

### 4.4 再次抛出异常与异常链 ###

在catch中可以再抛出一个异常，这样做的目的是改变异常的类型。比如执行servlet的代码可能不想知道发生错误的细节原因，但希望明确的知道servlet是否又问题。

` try { access the database } catch (SQLExeption e) { Throwable se = new ServletException(); se.initCause(e); throw se; } 复制代码`

这样做在捕获到异常时就可以使用:

` Throwable e = se.getCause(); 复制代码`

重新得到原始异常而补丢失细节。

### 4.5 finally子句 ###

` @Test public void myFinally () { try { InputStream in = new FileInputStream( "/home/lighk/test.txt" ); try { in.read(); /*take care! There is might throw other Exception*/ } finally { in.close(); } } catch (IOException e) { e.printStackTrace(); } } 复制代码`

* 内层的try只有一个职责：却倒关闭输入流。
* 外层的try也只有一个职责：确保出现的错误。

这里又一个问题，如果内层的try只是抛出 ` IOException` 那么这种结构很适合。 但是如果内层的try抛出了非IOException，这些异常只有这个方法的调用者才能处理。 执行这个finally语句块，并调用close。而close()方法本身也可能抛出IOException。 当这种情况出现，原始异常会丢失，转而抛出close()方法的异常。然而第一个异常可能会更有意思。 若嵌套多层try语句也可以解决这个问题，但这会让代码变得很繁琐。 Java 7 未这种代码模式提供了一个很又用的快捷方式－－带资源的try语句。

### 4.6 带资源的try语句 ###

` try (Scanner in = new Scanner( new FileInputStream( "/home/lighk/test.txt" )); PrintWriter out = New PrintWriter ( "out.txt" ) ) { while (in.hasNext()) out.println(in.next().toUpperCase()); } 复制代码`

前提是这个资源必须实现了AutoCloseable接口，或者其子接口Closeable.

不论这这个方法如何退出，in & out 都会关闭。若try抛出一个异常，close（）抛出一个异常，close() 的异常会被｀抑制｀。这些异常会被自动捕获，并由addSuppressed方法增加到原来的异常。若对这个异常感兴趣， 可以调用getSuppressed方法获取close抛出并被抑制的异常列表。

## 5.使用异常机制的技巧有哪些？ ##

* 异常不能代替简单的测试
` if (! s.empty()) s.pop(); 复制代码` ` try { s.pop(); } catch (EmptyStackException e) { /*&emsp;do something */ } 复制代码`

前者的性能要比后者高的多，因此只在异常情况下使用异常机制。

* 不要过分的细化异常,将正常处理与错误处理分开
* 利用异常的层次结构
* 不要压制异常
* 检测错误时，｀苛刻｀比放任更好 在无效的参数调用一个方法时，返回一个虚拟的数字还是抛出一个异常？ eg. when stack is empty, we should return a null or throw a EmptyStackException? Horstmann recommanded :

> 
> 
> 
> Throw a EmptyStackException is better than throw a NullPointerException in
> the following code to the error.
> 
> 

* 不要羞于传递异常 ５和６可以归结为 **早抛出，晚捕获**