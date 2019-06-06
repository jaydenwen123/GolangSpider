# Java volatile关键字解析 #

![image](https://user-gold-cdn.xitu.io/2019/3/19/169950c461712466?imageView2/0/w/1280/h/960/ignore-error/1)

## volatile简介 ##

` volatile` 被称为 **轻量级的synchronized** ，运行时开销比 ` synchronized` 更小，在多线程并发编程中发挥着 **同步共享变量** 、 **禁止处理器重排序** 的重要作用。建议在学习 ` volatie` 之前，先看一下Java内存模型 [《什么是Java内存模型？》]( https://link.juejin.im?target=https%3A%2F%2Fddnd.cn%2F2019%2F03%2F11%2Fjava-memory-model%2F ) ，因为 ` volatile` 和Java内存模型有着莫大的关系。

## Java内存模型 ##

在学习 ` volatie` 之前，需要补充下Java内存模型的相关(JMM)知识，我们知道Java线程的所有操作都是在工作区进行的，那么工作区和主存之间的变量是怎么进行交互的呢，可以用下面的图来表示。

![](https://user-gold-cdn.xitu.io/2019/3/19/1699456ae3425c46?imageView2/0/w/1280/h/960/ignore-error/1) Java通过几种原子操作完成 **工作区内存** 和 **主存** 的交互

* lock：作用于主存，把变量标识为线程独占状态。
* unlock：作用于主存，解除变量的独占状态。
* read：作用于主存，把一个变量的值通过主存传输到线程的工作区内存。
* load：作用于工作区内存，把 ` read` 操作传过来的变量值储存到工作区内存的变量副本中。
* use：作用于工作内存，把工作区内存的变量副本传给执行引擎。
* assign：作用于工作区内存，把从执行引擎传过来的值赋值给工作区内存的变量副本。
* store：作用于工作区内存，把工作区内存的变量副本传给主存。
* write：作用于主存，把 ` store` 操作传过来的值赋值给主存变量。

**这 ` 8` 个操作每个操作都是原子性的，但是几个操作连着一起就不是原子性了！**

## volatile原理 ##

上面介绍了Java模型的 ` 8` 个操作，那么这 ` 8` 个操作和 ` volatile` 又有着什么关系呢。

### volatile的可见性 ###

什么是 **可见性** ，用一个例子来解释，先看一段代码，加入线程 ` 1` 先执行，线程 ` 2` 再执行

` //线程1 boolean stop = false ; while (!stop) { do (); } //线程2 stop = true ; 复制代码`

线程 ` 1` 执行后会进入到一个死循环中，当线程 ` 2` 执行后，线程 ` 1` 的死循环就一定会马上结束吗？答案是不一定，因为线程 ` 2` 执行完 ` stop = true` 后，并不会马上将变量 ` stop` 的值 ` true` 写回主存中，也就是上图中的 ` assign` 执行完成之后， ` store` 和 ` write` 并不会随着执行， **线程 ` 1` 没有立即将修改后的变量的值更新到主存中** ，即使线程 ` 2` 及时将变量 ` stop` 的值写回主存中了， **线程 ` 1` 也没有了解到变量 ` stop` 的值已被修改而去主存中重新获取** ，也就是线程 ` 1` 的 ` load` 、 ` read` 操作并不会马上执行造成线程 ` 1` 的工作区内存中的变量副本不是最新的。这两个原因造成了线程 ` 1` 的死循环也就不会马上结束。
那么如何避免上诉的问题呢？我们可以使用 ` volatile` 关键字修饰变量 ` stop` ，如下

` //线程1 volatile boolean stop = false ; while (!stop) { do (); } //线程2 stop = true ; 复制代码`

这样线程 ` 1` 每次读取变量 ` stop` 的时候都会先去主存中获取变量 ` stop` 最新的值，线程 ` 2` 每次修改变量 ` stop` 的值之后都会马上将变量的值写回主存中，这样也就不会出现上述的问题了。

那么关键字 ` volatie` 是如何做到的呢？ ` volatie` 规定了上述 ` 8` 个操作的规则

* 只有当线程对变量执行的 **前一个操作** 是 ` load` 时，线程才能对变量执行 ` use` 操作；只有线程的后一个操作是 ` use` 时，线程才能对变量执行 ` load` 操作。即规定了 ` use` 、 ` load` 、 ` read` 三个操作之间的约束关系， **规定这三个操作必须连续的出现，保证了线程每次读取变量的值前都必须去主存获取最新的值** 。
* 只有当前程对变量执行的 **前一个操作** 是 ` assign` 时，线程才能对变量执行 ` store` 操作；只有线程的后一个操作是 ` store` 时，线程才能对变量执行 ` assign` 操作，即规定了 ` assign` 、 ` store` 、 ` write` 三个操作之间的约束关系， **规定了这三个操作必须连续的出现，保证线程每次修改变量后都必须将变量的值写回主存** 。

` volatile` 的这两个规则，也正是保证了 **共享变量的可见性** 。

### volatile的有序性 ###

有序性即程序执行的顺序按照代码的先后顺序执行，Java内存模型(JMM)允许编译器和处理器对指令进行重排序，但是规定了 ` as-if-serial` 语义，即保证 **单线程** 情况下不管怎么重排序，程序的结果不能改变，如

` double pi = 3.14; //A double r = 1; //B double s = pi * r * r; //C 复制代码`

上面的代码可能按照 ` A->B->C` 顺序执行，也有可能按照 ` B->A->C` 顺序执行，这两种顺序都不会影响程序的结果。但是不会以 ` C->A(B)->B(A)` 的顺序去执行，因为 ` C` 语句是依赖于 ` A` 和 ` B` 的，如果按照这样的顺序去执行就不能保证结果不变了(违背了 ` as-if-serial` )。

上面介绍的是单线程的执行，不管指令怎么重排序都不会影响结果，但是在多线程下就会出现问题了。
下面看个例子

` double pi = 3.14; double r = 0; double s = 0; boolean start = false ; //线程1 r = 10; //A start = true ; //B //线程2 if (start) { //C s = pi * r * r; //D } 复制代码`

线程 ` 1` 和线程 ` 2` 同时执行，线程 ` 1` 的 ` A` 和 ` B` 的执行顺序可能是 ` A->B` 或者 ` B->A` (因为A和B之间没有依赖关系，可以指令重排序)。如果线程 ` 1` 按照 ` A->B` 的顺序执行，那么线程 ` 2` 执行后的结果s就是我们想要的正确结果，如果线程 ` 1` 按照 ` B->A` 的顺序执行，那么线程 ` 2` 执行后的结果s可能就不是我们想要的结果了，因为线程 ` 1` 将变量 ` stop` 的值修改为 ` true` 后，线程 ` 2` 马上获取到 ` stop` 为 ` true` 然后执行 ` C` 语句，然后执行 ` D` 语句即 ` s = 3.14 * 0 * 0` ，然后线程 ` 1` 再执行 ` B` 语句，那么结果就是有问题了。

那么为了解决这个问题，我们可以在变量 ` true` 加上关键字 ` volatile`

` double pi = 3.14; double r = 0; double s = 0; volatile boolean start = false ; //线程1 r = 10; //A start = true ; //B //线程2 if (start) { //C s = pi * r * r; //D } 复制代码`

这样线程 ` 1` 的执行顺序就只能是 ` A->B` 了，因为关键字 **发挥了禁止处理器指令重排序的作用** ，所以线程 ` 2` 的执行结果就不会有问题了。

那么 ` volatile` 是怎么实现禁止处理器重排序的呢？
**编译器会在编译生成字节码的时候，在加有 ` volatile` 关键字的变量的指令进行插入内存屏障来禁止特定类型的处理器重排序**
我们先看 **内存屏障** 有哪些及发挥的作用

![image](https://user-gold-cdn.xitu.io/2019/3/20/16998f6f4ab7bb24?imageView2/0/w/1280/h/960/ignore-error/1)

* ` StoreStore` 屏障：禁止屏障上面变量的写和下面所有进行写的变量进行处理器重排序。
* ` StoreLoad` 屏障：禁止屏障上面变量的写和下面所有进行读的变量进行处理器重排序。
* ` LoadLoad` 屏障：禁止屏障上面变量的读和下面所有进行读的变量进行处理器重排序。
* ` LoadStore` 屏障：禁止屏障上面变量的读和下面所有进行写的变量进行处理器重排序。

再看 ` volatile` 是怎么插入屏障的

* 在每个 ` volatile` 变量的写 **前面** 插入一个 ` StoreStore` 屏障。
* 在每个 ` volatile` 变量的写 **后面** 插入一个 ` StoreLoad` 屏障。
* 在每个 ` volatile` 变量的读 **后面** 插入一个 ` LoadLoad` 屏障。
* 在每个 ` volatile` 变量的读 **后面** 插入一个 ` LoadStore` 屏障。

> 
> 
> 
> 注意：写操作是在 ` volatile` **前后** 插入一个内存屏障，而读操作是在 **后面** 插入两个内存屏障。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/3/19/16994e005f6e4dc2?imageView2/0/w/1280/h/960/ignore-error/1)

**` volatile` 变量通过插入内存屏障禁止了处理器重排序，从而解决了多线程环境下处理器重排序的问题** 。

### volatile有没有原子性？ ###

上面分别介绍了 ` volatile` 的可见性和有序性，那么 ` volatile` 有原子性吗？我们先看一段代码

` public class Test { public volatile int inc = 0; public void increase () { inc++; } public static void main(String[] args) { final Test test = new Test(); for (int i=0;i<10;i++){ new Thread (){ public void run () { for (int j=0;j<1000;j++) test.increase(); }; }.start(); } while (Thread.activeCount()>1) //保证前面的线程都执行完 Thread.yield(); System.out.println(test.inc); } } 复制代码`

我们开启 ` 10` 个线程对 ` volatile` 变量进行自增操作，每个线程对 ` volatile` 变量执行 ` 1000` 次自增操作，那结果变量 ` inc` 会是 ` 10000` 吗？答案是，变量 ` inc` 的值基本都是小于 ` 10000` 。
可能你会有疑问， ` volatile` 变量 ` inc` 不是保证了共享变量的可见性了吗，每次线程读取到的都是最新的值，是的没错， **但是线程每次将值写回主存的时候并不能保证主存中的值没有被其他的线程修过过** 。

![](https://user-gold-cdn.xitu.io/2019/3/19/16994f4d97616f74?imageView2/0/w/1280/h/960/ignore-error/1)

如果所示：线程 ` 1` 在主存中获取了 ` i` 的最新值(i=1)，线程 ` 2` 也在主存中获取了 ` i` 的最新值(i=1，注意这时候线程 ` 1` 并未对变量 ` i` 进行修改，所以 ` i` 的值还是 ` 1` )），然后线程 ` 2` 将i自增后写回主存，这时候主存中 ` i=2` ，到这里还没有问题，然后线程 ` 1` 又对i进行了自增写回了主存，这时候主存中 ` i=2` ，也就是对i做了2次自增操作，结果i的结果只自增了1，问题就出来了这里。

为什么会有这个问题呢，前面我们提到了Java内存模型和主存之间交互的 ` 8` 个操作都是原子性的，但是他们的操作连在一起就不是原子性了，而 ` volatile` 关键字也只是保证了 ` use` 、 ` load` 、 ` read` 三个操作连在一起时候的原子性，还有 ` assign` 、 ` store` 、 ` write` 这三个操作连在一起时候的原子性，也就是 ` volatile` 关键字 **保证了变量读操作的原子性和写操作的原子性，而变量的自增过程需要对变量进行读和写两个过程，而这两个过程连在一起就不是原子性操作了。**

所以说 ` volatile` 变量对于变量的单独写操作/读操作是保证了原子性的，而常说的原子性包括读写操作连在一起，所以说对于 ` volatile` 不保证原子性的。那么如何解决上面程序的问题呢？只能给 ` increase` 方法加锁，让在多线程情况下只有一个线程能执行 ` increase` 方法，也就是保证了一个线程对变量的读写是原子性的。 **当然还有个更优的方案，就是利用读写都为原子性的 ` CAS` ，利用 ` CAS` 对 ` volatile` 进行操作，既解决了 ` volatile` 不保证原子性的问题，同时消耗也没加锁的方式大**

## volatile和CAS ##

学完 ` volatile` 之后，是不是觉得 ` volatile` 和 ` CAS` 有种似曾相识的感觉？那它们之间有什么关系或者区别呢。

* ` volatile` 只能保证共享变量的读和写操作单个操作的原子性，而 ` CAS` 保证了共享变量的读和写两个操作一起的原子性(即CAS是原子性操作的)。
* ` volatile` 的实现基于 ` JMM` ，而 ` CAS` 的实现基于硬件。

## 参考 ##

[Java并发编程：volatile关键字解析]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fdolphin0520%2Fp%2F3920373.html )
[JAVA并发六：彻底理解volatile]( https://link.juejin.im?target=https%3A%2F%2Fwww.javazhiyin.com%2F887.html )
[Java内存模型与volatile]( https://link.juejin.im?target=https%3A%2F%2Fjiangzhengjun.iteye.com%2Fblog%2F652532 )
[Java面试官最爱问的volatile关键字]( https://link.juejin.im?target=http%3A%2F%2Fwww.techug.com%2Fpost%2Fjava-volatile-keyword.html )

原文地址： [ddnd.cn/2019/03/19/…]( https://link.juejin.im?target=https%3A%2F%2Fddnd.cn%2F2019%2F03%2F19%2Fjava-volatile%2F )