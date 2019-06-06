# 你是否真的了解全局解析锁(GIL) #

> 
> 
> 
> 关于我
> 编程界的一名小程序猿，目前在一个创业团队任team lead，技术栈涉及Android、Python、Java和Go，这个也是我们团队的主要技术栈。
> 联系：hylinux1024@gmail.com
> 
> 

### 0x00 什么是全局解析锁(GIL) ###

> 
> 
> 
> A global interpreter lock (GIL) is a mechanism used in computer-language
> interpreters to synchronize the execution of threads so that only one
> native thread can execute at a time. -- [引用自wikipedia](
> https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FGlobal_interpreter_lock
> )
> 
> 

从上面的定义可以看出， ` GIL` 是计算机语言解析器用于同步线程执行的一种 **同步锁机制** 。很多编程语言都有 ` GIL` ，例如 ` Python` 、 ` Ruby` 。

### 0x01 为什么会有GIL ###

` Python` 作为一种面向对象的动态类型编程语言，开发者编写的代码是通过解析器顺序解析执行的。 大多数人目前使用的 ` Python` 解析器是 ` CPython` 提供的，而 ` CPython` 的解析器是 **使用引用计数来进行内存管理** ，为了对多线程安全的支持，引用了 ` global intepreter lock` ，只有获取到 ` GIL` 的线程才能执行。如果没有这个锁，在多线程编码中即使是简单的操作也会引起共享变量被多个线程同时修改的问题。例如有两个线程 **同时对同一个对象进行引用时，这两个线程都会将变量的引用计数从0增加为1** ，明显这是不正确的。
可以通过 ` sys` 模块获取一个变量的引用计数

` >>> import sys >>> a = [] >>> sys.getrefcount(a) 2 >>> b = a >>> sys.getrefcount(a) 3 复制代码`

` sys.getrefcount()` 方法中的参数对a的引用也会引起计数的增加。

**是否可以对每个变量都分别使用锁来同步呢？**

如果有多个锁的话，线程同步时就容易出现 **死锁** ，而且编程的复杂度也会上升。当全局只有一个锁时，所有线程都在竞争一把锁，就不会出现相互等待对方锁的情况，编码的实现也更简单。此外只有一把锁时对单线程的影响其实并不是很大。

### 0x02 可以移除GIL吗？ ###

` Python` 核心开发团队以及 ` Python` 社区的技术专家对移除 ` GIL` 也做过多次尝试，然而最后都没有令各方满意的方案。

内存管理技术除了 **引用计数** 外，一些编程语言为了避免引用全局解析锁，内存管理就使用 **垃圾回收** 机制。

当然这也意味着这些使用垃圾回收机制的语言就必须提升其它方面的性能（例如 ` JIT` 编译），来弥补单线程程序的执行性能的损失。
对于 ` Python` 的来说，选择了 **引用计数** 作为内存管理。一方面保证了 **单线程程序执行的性能** ，另一方面 ` GIL` 使得编码也更容易实现。
在 ` Python` 中很多特性是通过 ` C` 库来实现的，而在 ` C` 库中要保证线程安全的话也是依赖于 ` GIL` 。

所以当有人成功移除了 ` GIL` 之后， ` Python` 的程序并没有变得更快，因为大多数人使用的都是单线程场景。

### 0x03 对多线程程序的影响 ###

首先来 ` GIL` 对 ` IO` 密集型程序和 ` CPU` 密集型程序的的区别。 像文件读写、网络请求、数据库访问等操作都是 ` IO` 密集型的，它们的特点 **需要等待 ` IO` 操作的时间** ，然后才进行下一步操作；而像数学计算、图片处理、矩阵运算等操作则是 ` CPU` 密集型的，它们的特点是 **需要大量 ` CPU` 算力来支持** 。

对于 ` IO` 密集型操作，当前拥有锁的线程会先释放锁，然后执行 ` IO` 操作，最后再获取锁。线程在释放锁时会把当前线程状态存在一个全局变量 ` PThreadState` 的数据结构中，当线程获取到锁之后恢复之前的线程状态

用文字描述执行流程

` 保存当前线程的状态到一个全局变量中 释放GIL... 执行IO操作 ... 获取GIL 从全局变量中恢复之前的线程状态 复制代码`

下面这段代码是测试单线程执行500万次消耗的时间

` import time COUNT = 50000000 def countdown (n) : while n > 0 : n -= 1 start = time.time() countdown(COUNT) end = time.time() print( 'Time taken in seconds -' , end - start) # 执行结果 # Time taken in seconds - 2.44541597366333 复制代码`

在我的8核的 ` macbook` 上跑大约是2.4秒，然后再看一个多线程版本

` import time from threading import Thread COUNT = 50000000 def countdown (n) : while n > 0 : n -= 1 t1 = Thread(target=countdown, args=(COUNT // 2 ,)) t2 = Thread(target=countdown, args=(COUNT // 2 ,)) start = time.time() t1.start() t2.start() t1.join() t2.join() end = time.time() print( 'Time taken in seconds -' , end - start) # 执行结果 # Time taken in seconds - 2.4634649753570557 复制代码`

上文代码每个线程都执行250万次，如果线程是并发的，执行时间应该是上面单线程版本的一半时间左右，然而在我电脑中执行时间大约为2.5秒！ 多线程不但没有更高效率，反而还更耗时了。这个例子就说明 ` Python` 中的线程是顺序执行的，只有获取到锁的线程可以获取解析器的执行时间。多线程执行多出来的那点时间就是获取锁和释放锁消耗的时间。

**那如何实现高并发呢？**

答案是使用多进程。 [前面的文章有介绍多进程的使用]( https://juejin.im/post/5cefdc60f265da1bca51c0cf )

` from multiprocessing import Pool import time COUNT = 50000000 def countdown (n) : while n > 0 : n -= 1 if __name__ == '__main__' : pool = Pool(processes= 2 ) start = time.time() r1 = pool.apply_async(countdown, [COUNT // 2 ]) r2 = pool.apply_async(countdown, [COUNT // 2 ]) pool.close() pool.join() end = time.time() print( 'Time taken in seconds -' , end - start) # 执行结果 # Time taken in seconds - 1.2389559745788574 复制代码`

使用多进程，每个进程运行250万次，大约消耗1.2秒的时间。差不多是上面线程版本的一半时间。

当然还可以使用其它 ` Python` 解析器，例如 ` Jython` 、 ` IronPython` 或 ` PyPy` 。

**既然每个线程执行前都要获取锁，那么有一个线程获取到锁一直占用不释放，怎么办？**

` IO` 密集型的程序会主动释放锁，但对于 ` CPU` 密集型的程序或 ` IO` 密集型和 ` CPU` 混合的程序，解析器将会如何工作呢？
早期的做法是 ` Python` 会执行100条指令后就强制线程释放 ` GIL` 让其它线程有可执行的机会。
可以通过以下获取到这个配置

` >>> import sys >>> sys.getcheckinterval() 100 复制代码`

在我的电脑中还打印了下面的输出警告

` Warning (from warnings module): File "__main__" , line 1 DeprecationWarning: sys.getcheckinterval() and sys.setcheckinterval() are deprecated. Use sys.getswitchinterval() instead. 复制代码`

意思是 ` sys.getcheckinterval()` 方法已经废弃，应该使用 ` sys.getswitchinterval()` 方法。 因为传统的实现中每解析100指令的就强制线程释放锁的做法，会导致 ` CPU` 密集型的线程会一直占用 ` GIL` 而 ` IO` 密集型的线程会一直得不到解析的问题。 [于是新的线程切换方案就被提出来了]( https://link.juejin.im?target=https%3A%2F%2Fmail.python.org%2Fpipermail%2Fpython-dev%2F2009-October%2F093321.html )

` >>> sys.getswitchinterval() 0.005 复制代码`

这个方法返回0.05秒，意思是每个线程执行0.05秒后就释放 ` GIL` ，用于线程的切换。

### 0x04 总结 ###

在 ` CPython` 解析器的实现由于 ` global interpreter lock` (全局解释锁)的存在，任何时刻都只有一个线程能执行 ` Python` 的 ` bytecode` (字节码)。
常见的内存管理方案有引用计数和垃圾回收， ` Python` 选择了前者，这保证了单线程的执行效率，同时对编码实现也更加简单。想要移除 ` GIL` 是不容易的，即使成功将 ` GIL` 去除，对 ` Python` 的来说是牺牲了单线程的执行效率。
` Python` 中 ` GIL` 对 ` IO` 密集型程序可以较好的支持多线程并发，然而对 ` CPU` 密集型程序来说就要使用多进程或使用其它不使用 ` GIL` 的解析器。
目前最新的解析器实现中线程每执行0.05秒就会强制释放 ` GIL` ，进行线程的切换。

### 0x05 为了看懂GIL我阅读了下面这些资料 ###

* [docs.python.org/3.7/glossar…]( https://link.juejin.im?target=https%3A%2F%2Fdocs.python.org%2F3.7%2Fglossary.html%23term-global-interpreter-lock )
GIL
* [docs.python.org/3.7/glossar…]( https://link.juejin.im?target=https%3A%2F%2Fdocs.python.org%2F3.7%2Fglossary.html%23term-bytecode )
bytecode字节码
* [github.com/python/cpyt…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpython%2Fcpython )
CPython 源码
* [wiki.python.org/moin/Global…]( https://link.juejin.im?target=https%3A%2F%2Fwiki.python.org%2Fmoin%2FGlobalInterpreterLock )
* [realpython.com/python-gil/]( https://link.juejin.im?target=https%3A%2F%2Frealpython.com%2Fpython-gil%2F )
What is the Python Global Interpreter Lock (GIL)?
* [dabeaz.blogspot.com/2010/01/pyt…]( https://link.juejin.im?target=http%3A%2F%2Fdabeaz.blogspot.com%2F2010%2F01%2Fpython-gil-visualized.html )
* [mail.python.org/pipermail/p…]( https://link.juejin.im?target=https%3A%2F%2Fmail.python.org%2Fpipermail%2Fpython-dev%2F2009-October%2F093321.html )
* [www.youtube.com/watch?v=Obt…]( https://link.juejin.im?target=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DObt-vMVdM8s%26amp%3Bfeature%3Dyoutu.be )
Understanding the Python GIL
* [en.wikipedia.org/wiki/Global…]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FGlobal_interpreter_lock )