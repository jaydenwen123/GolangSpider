# iOS开发中crash常用处理 #

**
**

## 引言 ##

iOS开发中我们会遇到程序异常退出的情况，如果是在调试的过程中，可能通过设施断点或者打印关键信息的方式来进行调试，但是对于一些复杂模块非必现的异常崩溃，这种方式有时难以定位问题，而且对于已经发布上线的应用，这种方式更是无能为力。

通常我们见到的Crash分为两种，一种是系统内存错误，触发 **EXC** **_** **BAD** **_** **ACCESS** 引起的，程序运行过程中访问了错误的内存地址，另一种是出现了不能被处理的signal异常，导致程序向自身发送了 **SIGABRT** 信号而崩溃，下面是我平时使用较多的crash处理方法，跟大家分享一下。

## Part 1：内存错误引起的crash ##

### ➸ 1.1 产生原因 ###

* 访问了不属于本进程的内存地址。

* 访问已被释放的内存(重复释放已经释放的内存)。

iOS开发中对继承自NSObject的对象，内存管理采用引用计数机制，对于非NSObject对象，引用计数机制不起作用，需要自己管理内存的使用回收。引用计数原理是，当对象被持有一次（retain），它的引用计数（retainCount）+1，被标记释放一次（release），retainCount -1，当retainCount为0，对象被释放。

在使用手动引用计数（MRC）开发时，需要开发者显式的调用retain或release。苹果在iOS5之后推行自动引用计数（ARC），开发者不必显式的调用retain或release了，由编译器来自动添加，在方便开发者的同时，也降低了开发难度，同时内存出现问题的时候也更难定位错误了。

### ➸ 1.2 常用处理方法 ###

****

**1.2.1 添加Xcode全局异常断点**

① 将导航器视图 **切换到断点导航器视图** 下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a5496f6e6?imageView2/0/w/1280/h/960/ignore-error/1)

② 点击左下角的+号， **选择** **Exception Breakpoint** 这一选项。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a54e0e8f3?imageView2/0/w/1280/h/960/ignore-error/1)

③ **异常断点** 可以编辑许多功能，例如执行脚本，输出log，选择只处理objective-c异常等等，功能很丰富。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a54a3318b?imageView2/0/w/1280/h/960/ignore-error/1)

下一次当我们运行程序出现崩溃的时候，程序会自动停止在崩溃的代码处，方便我们查找问题。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a57a63c21?imageView2/0/w/1280/h/960/ignore-error/1)

**1.2.2 僵尸对象调试**

全局异常断点通常情况下都会把崩溃原因定位到具体代码中。但是，如果崩溃不在当前调用栈，系统就仅仅只能把崩溃地址告诉我们，而没办法定位到具体代码，这样我们也没法去修改错误。类似下面这种：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a55aabfc6?imageView2/0/w/1280/h/960/ignore-error/1)

这种情况下我们可以通过Xcode提供的僵尸对象调试（Zombie Objects）来尝试找到问题。

**①** 首先还是打开Xcode **选择屏幕左上角Xcode-> Preferencese** ，在behavior选项卡中，设置一下输出信息，调试的时候输出更多的信息，如下截图，勾上：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a610b1b06?imageView2/0/w/1280/h/960/ignore-error/1)

② 菜单 **Product > Scheme > Edit Scheme** 中，把红色圈里面的三个选项都勾上：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a75696182?imageView2/0/w/1280/h/960/ignore-error/1)

③ 开启该选项后，程序在运行时，如果 **访问了已经释放的对象** ，则会给出较准确的定位信息，可以帮助确定问题所在。

该功能的原理是，在对象释放（retainCount为0)时，使用一个内置的Zombie对象，替代原来被释放的对象。无论向该对象发送什么消息（函数调用），都会触发异常，抛出调试信息。

**注意：记得在问题被修复后，关闭该功能！会引起程序内存占用异常。**

④ 也可以通过系统 **terminal** 打印出调用信息，使用终端的 **mallochistory** 命令，例如"mallochistory 30495 0x60005ef76fd0"，其中30495是该进程的pid，pid可以根据Xcode控制台中的log查看，或者通过活动监视器获得， 根据这个记录，可以大致判断出错误代码的位置。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a89a2011e?imageView2/0/w/1280/h/960/ignore-error/1)

会出现类似以下提示代码，根据一些关键信息，就可以找出错误具体位置。

****

![](https://user-gold-cdn.xitu.io/2019/6/5/16b267ee6575a15f?imageView2/0/w/1280/h/960/ignore-error/1)

**
**

**1.2.3 利用NSSetUncaughtExceptionHandler处理**

之前的两种方式，对于线上的APP可以说无能为力的，还好，iOS提供了异常发生的处理API，NSSetUncaughtExceptionHandler，我们在程序启动的时候可以添加这样的Handler，这样的程序发生异常的时候就可以对这一部分的信息进行必要的处理，适时的反馈给开发者。需要注意的是，利用NSSetUncaughtExceptionHandler可以用来处理异常崩溃，崩溃报告系统会用NSSetUncaughtExceptionHandler方法设置全局的异常处理器。如果自定义NSSetUncaughtExceptionHandler监听事件，会导致第三方监听（如Bugly）失效，已经集成了第三方监听平台的小伙伴需要注意。

① **注册全局处理异常的handler** ，在程序启动或者其他入口注册：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a7cfe0203?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a8fa819ae?imageView2/0/w/1280/h/960/ignore-error/1)

② 当线上程序出现crash时， **代码会执行到之前注册的handle中** ，将错误信息保存在本地。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675a87233111?imageView2/0/w/1280/h/960/ignore-error/1)

③ 通过保存的线上app的 **dSYM符号表** 查找问题 **。**

iOS构建时产生的符号表，它是内存地址与函数名，文件名，行号的映射表。 符号表元素如下所示:

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675aa2befff1?imageView2/0/w/1280/h/960/ignore-error/1)

当应用crash时，我们可以利用crash时的堆栈信息得到对应到源代码的堆栈信息，还能看到出错的代码在多少行，所以能快速定位出错的代码位置，以便快速解决问题。

获取到dSYM符号表和之前的程序崩溃的错误日志，我们就可以定位问题了。

④ 利用 **atos命令** 定位问题。

atos命令来符号化某个特定模块加载地址 atos [-arch 架构名] [-o 符号表] [-l 模块地址] [方法地址]

使用终端计算，首先获得十六进制地址区间。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675ae0e909f6?imageView2/0/w/1280/h/960/ignore-error/1)

终端代码执行：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675aa3a0ff08?imageView2/0/w/1280/h/960/ignore-error/1) 这样，就可以定位出问题代码。

## Part 2：Mach异常和signal信号引发的crash ##

### ➸ 2.1 Mach和signal ###

Mach是Mac OS和iOS操作系统的微内核核心，Mach异常是指最底层的内核级异常，所以当APP中产生异常时，最先能监听到异常的就是Mach。

最先捕获到异常的Mach在接下来会将所有的异常转换为相应的Unix信号，并投递到出错的线程。之后就可以注册想要监听的signal类型，来捕获信号。使用Objective-C的异常处理是不能得到signal的，如果要处理它，我们还要利用unix标准的signal机制，注册SIGABRT, SIGBUS, SIGSEGV等信号发生时的处理函数。该函数中我们可以输出栈信息，版本信息等其他一切我们所想要的。如下，就是监听了SIGSEGV信号，当有SIGSEGV信号产生时，就会回调mySignalHandler方法： **signal** **** (SIGSEGV, mySignalHandler)。

**
**

### ➸ 2.2 signal信号说明 ###

信号默认的处理方法一共有五种，分别用Terminate (terminate process，即结東进程)、Ignore(忽略该信号)、Dump(terminate process and dump core：结束进程并生成 core dump，将进程的内存信息打印出来)，Stop（进程暂停运行，多用于调试）以及 Cont（恢复运行个之前被暂停的进程，多用于调试）来表示。

Signal信号类型：

+----------+-----------+----------------------------+
| 信号名称 | 默认处理  | 说明                       |
| SIGABRT  | Dump      | 程序终止命令               |
| SIGALRM  | Terminate | 程序超时信号               |
| SIGILL   | Dump      | 程序非法指令信号           |
| SIGHUP   | Terminate | 程序终端中止信号           |
| SIGINT   | Terminate | 程序键盘中断信号           |
| SIGKILL  | Terminate | 程序强制结束信号           |
| SIGTERM  | Terminate | 程序终止信号               |
| SIGSTOP  | Stop      | 程序键盘中止信号           |
| SIGSEGV  | Dump      | 程序无效内存中止信号       |
| SIGBUS   | Dump      | 程序内存字节未对齐中止信号 |
| SIGPIPE  | Terminate | 程序Socket发送失败中止信号 |
+----------+-----------+----------------------------+

如果没有为一个信号设置对应的处理函数，就会使用默认的处理函数，否则信号就被进程截获并调用相应的处理函数。在没有处理函数的情况下，程序可以指定两种行为：忽略这个信号 SIG_IGN 或者用默认的处理函数 SIG_DFL 。但是有两个信号是无法被截获并处理的： SIGKILL、SIGSTOP 。

### ➸ 2.3 处理signal信号 ###

① **注册全局处理异常signal的handler** ，在程序启动或者其他入口注册。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675ab35f1c2e?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675ab5743d6f?imageView2/0/w/1280/h/960/ignore-error/1)

有关错误类型可以看上面的说明，SignalExceptionHandler是信号出错时候的回调。当有信号出错的时候，可以回调到这个方法。

② **SignalHandler不要在debug环境下测试** 。因为系统的debug会优先去拦截。我在模拟器上运行一次后，关闭debug状态，然后直接在模拟器上点击我们build上去的app去运行。获得如下的日志：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2675aa7ac8e2f?imageView2/0/w/1280/h/960/ignore-error/1)

## Part 3：总结 ##

对于应用crash，有很多第三方优秀平台(友盟，bugly等)提供日志和打点功能，已经能够满足日常开发需要，但是学习这些常用的crash还是能够帮助我们理解iOS运行机制，以上这些是开发中经常见到的一些crash，实际处理中可能情况复杂，需要多种方式同时使用才能定位问题，灵活使用。