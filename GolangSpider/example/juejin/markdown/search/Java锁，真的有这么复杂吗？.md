# Java锁，真的有这么复杂吗？ #

### 前言 ###

作者前面也写了几篇关于Java并发编程，以及线程和volatil的基础知识，有兴趣可以阅读作者的原文博客，今天关于Java中的两种锁进行详解，希望对你有所帮助

> 
> 
> 
> 本文受赵sir原创发布，转载请联系原创 [blog.csdn.net/qq_36094018…](
> https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_36094018%2Farticle%2Fdetails%2F90140209
> )
> 
> 

### 为什么使用synchronized ###

在上一章中说了volatile，在多线程下可以保证变量的可见性，但是不能保证原子性，下面一段代码说明：

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bbbf2f3e4?imageView2/0/w/1280/h/960/ignore-error/1)

运行上面代码，会发现输出flag的值不是理想中10000，虽然volatile写入时候会通知其他线程的工作内存值无效，从主内存重写读取。i++是三步操作，读取-赋值-写入不能保证原子性。 **原子性：不能被中断要么成功要么失败。**

**比如此时主内存的flag值10，线程1和线程2读取到自己工作内存都是10，然后线程1在进行赋值的时候，线程2执行了，这时线程2发现自己内存的值和主内存的值一样，并没有修改，然后赋值写入11，此时线程1运行，因为之前读过了，会往下继续运行写入也是11。那么两个线程相当于只增加了一次** 。要想达到理想值，只需要修改 ` public synchronized void increase() { flag++; }` 就行了。

### 什么是synchronized ###

**Java提供的一种原子性性内置锁，Java每个对象都可以把它当做是监视器锁，线程代码执行在进入synchronized代码块时候会自动获取内部锁，这个时候其他线程访问时候会被阻塞到队列，直到进入synchronized中的代码执行完毕或者抛出异常或者调用了wait方法，都会释放锁资源。在进入synchronized会从主内存把变量读取到自己工作内存，在退出的时候会把工作内存的值写入到主内存，保证了原子性。**

### synchronized机制 ###

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bbbe13432?imageView2/0/w/1280/h/960/ignore-error/1)

编译后执行 **javap -v Test.class** 就会发现两条指令。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bbd09794f?imageView2/0/w/1280/h/960/ignore-error/1)

**synchronized是使用一种monitor机制，在进入锁时候先执行monitorenter指令。退出的时候执行monitorexit指令。synchronized是可重入锁，每个对象中都含有一个计数器当前线程再次获取锁，计数器+1，退出时候计算器-1，直到计数器为0才释放锁资源，唤醒其他线程来争抢资源。任意一个对象都拥有自己的监视器，只有在线程获取到监视器锁时才会进入代码中，否则就进入阻塞状态。**

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bbcf6e229?imageView2/0/w/1280/h/960/ignore-error/1)

### synchronized使用场景 ###

* 对于普通方法，锁是当前类实例对象。
* 对于静态方法，锁是当前类对象。
* 对于同步代码块，锁是synchronized括号里的对象。

### synchronized锁升级 ###

**synchronized在1.6以前是重量级锁** ，当前只有一个线程执行，其他线程阻塞。为了减少获得锁和释放锁带来的性能问题，而引入了偏向锁、轻量级锁以及锁的存储过程和升级过程。在1.6后锁分为了无锁、偏向锁、轻量锁、重量锁，锁的状态在多线程竞争的情况下会逐渐升级，只能升级而不能降级，这样是为了提高锁获取和释放的效率。

synchronized的锁是存贮在Java对象头里的，如果对象是数组类型，则虚拟机用3个字宽（Word）存储对象头，如果对象是非数组类型，则用2字宽存储对象头。1个字宽等于4个字节。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bbd1fd7f7?imageView2/0/w/1280/h/960/ignore-error/1) Java对象头中的Mark Word里默认存储了对象是HashCode、分代年龄、和锁标记。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bc7fb5572?imageView2/0/w/1280/h/960/ignore-error/1) 在运行的时候，Mark Word里存储的数据会随着锁标志位的变化而变化，可能会变化为存储以下四种形式。 ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bda4ae528?imageView2/0/w/1280/h/960/ignore-error/1)

### 偏向锁 ###

偏向锁的意思未来只有一个线程使用锁，不会有其他线程来争取。

**获取锁：**

* 

首先检查Mark word中锁的标志是否为01。

* 

如果是01，判断对象头的Mark word记录是否为当前线程ID，如果是执行5，否则执行3.

* 

线程ID并未只指向自己，发送CAS竞争，如果竞争成功，则将Mark Word中线程ID设置为当前线程ID，执行5；如果未成功执行4。

* 

当到达全局安全点（在这个时间点上没有正在执行的字节码）时获得偏向锁的线程被挂起，偏向锁升级为轻量级锁，然后被阻塞在安全点的线程继续往下执行同步代码。

* 

执行同步代码。 撤销锁：偏向锁使用了一种等到竞争出现才释放锁的机制，所以当其他线程尝试竞争偏向锁时，持有偏向锁的线程才会释放锁。需要等待全局安全点，它首先暂停原持有偏向锁的线程，然后检查线程是否还在活着，如果线程处于未活动状态，则释放锁标记，如果处于活动状态则升级为轻量级锁。

### CAS ###

**CAS全称是Compare And Swap 即比较并交换，使用乐观锁机制，包含三个操作数 —— 内存位置（V）、预期原值（A）和新值(B)。 如果内存位置的值与预期原值相匹配，那么才会将该位置值更新为新值 。否则，处理器不做任何操作。**

### 轻量级锁 ###

线程在执行同步代码块之前，JVM会先在当前线程的栈桢中创建用于存储锁记录的空间，并将对象头中的Mark Word复制到锁记录中，官方称为Displaced Mark Word。

**加锁：**

* CAS修改Mark Word，如果成功指向栈中锁记录的指针执行3，如果失败执行2.
* 发生自旋，自旋到一定次数，如果修改成功执行3，否则锁膨胀为重量级锁。
* 执行同步代码块。

解锁： 轻量级解锁时，会使用原子的CAS操作将Displaced Mark Word替换回到对象头，如果成功，则表示没有竞争发生。如果失败，表示当前锁存在竞争，锁就会膨胀成重量级锁。

### 锁的优缺点 ###

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518bdaae1309?imageView2/0/w/1280/h/960/ignore-error/1)

### 彻底搞懂锁升级 ###

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c2d293db7?imageView2/0/w/1280/h/960/ignore-error/1)

### lock ###

它是在1.5之后提供的一个独占锁接口，它的实现类是ReentrantLock，相比较synchronized这种隐式锁（不用手动加锁和释放锁）的便捷性，但是提供了更加锁的可操作性、可中断的获取锁以及超时获取锁等多种synchronized不具备的特性。

### 使用方法 ###

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c18cceadd?imageView2/0/w/1280/h/960/ignore-error/1)

在finally中释放锁，目的保证获取锁最终被释放。不要在获取锁写在try里，因为如果在获取锁时发生了异常，异常抛出的同时，也会导致锁无故释放。

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c001783a3?imageView2/0/w/1280/h/960/ignore-error/1)

### AQS ###

AQS是队列同步器（AbstractQueuedSynchronizer），是用来构建锁或者其他同步器的基础框架，它使用了一个int成员变量表示同步状态，通过内置的FIFO队列来完成资源获取的线程排队工作问题。AQS在内部维护了一个单一的状态信息state，可以通过getState、setState、compareAndSetState（CAS操作）修改此值，对于ReentrantLock来说，state可以用来表示当前线程获取锁的可重入次数。ReentrantLock中当一个线程获取了锁，在AQS的内部会进行compareAndSetState将state变为1，如果再次获取就设置为2，释放锁也会去修改state值，只有当值变为0时，其他线程才能获得锁。

### 锁的介绍 ###

AQS底层维护state和队列来实现独占和共享两种锁。

**独占锁：**每次只能有一个线程能持有锁，如lock、synchronized。 **共享锁：**允许多个线程同时获取锁，并发访问共享资源，如ReadWriteLock。

lock分为公平锁和非公平锁，实现了AQS接口，通过FIFO设置锁的优先级。

**公平锁：**根据线程获取锁的时间来判断，等待时间越久的线程优先被执行。Lock中初始化的时候ReentrantLock（true），默认为false，效率较低因为需要判断线程的等待时间。

**非公平锁：**抢占锁资源，不能保证获取锁的线程优先级，效率较高，因为获取锁是竞争的。

### 两者不同 ###

* synchronized是Java的关键字，lock是提供的类。
* synchronized提供不需要手动加锁和释放的隐式锁，释放锁的条件是代码执行完或者抛出异常自动释放。lock必须手动加锁和释放锁，另外还提供了可中断锁、超时获取锁、判断锁状态。
* synchronized是可重入、不可中断、非公平，lock是可重入、可中断、公平（两者皆可）
* synchronized适合代码量少的同步，lock适合代码量同步多的。**

### Condition接口 ###

还记得在 [Java并发二]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_36094018%2Farticle%2Fdetails%2F90108528 ) 中有一道生产者消费者，使用的是synchronized+wait（notify），lock中也提供了这种等待通知类型的方法await和signal，当前线程调用这些方法时，需要提前获取到Condition对象关联的锁，Condition是依赖于Lock对象，调用lock对象中的newCondition。

老样子还是先定义一个容器：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c8c46d5ec?imageView2/0/w/1280/h/960/ignore-error/1) 生产者：启5个线程往容器里添加数据。

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c5bc5ac2e?imageView2/0/w/1280/h/960/ignore-error/1)

消费者：启10线程消费数据

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c7dc2ae14?imageView2/0/w/1280/h/960/ignore-error/1)

### 最后 ###

注释基本明确，就不多说了。wait和notify是配合synchronized使用，await和signal是配合lock使用，区别在于唤醒时notify不能指定线程唤醒，signal可以唤醒具体的线程，更小的粒度控制锁。

#### [阅读更多]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247487391%26amp%3Bidx%3D1%26amp%3Bsn%3Db69213accad6ddb6caedfc1a3d826c03%26amp%3Bchksm%3Deb476301dc30ea17a96b33c2886f91b2b9cff7f0b116570e4bd4819450b7a302e3ded19c99b7%26amp%3Bscene%3D21%23wechat_redirect ) ####

**金三银四，2019最新面试实战总结** ( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247487513%26amp%3Bidx%3D1%26amp%3Bsn%3D9c9a3feddc59c7bc153daf76fd306253%26amp%3Bchksm%3Deb477c87dc30f5917181c0099020636e50d0e1270615fd4c66f11d3341ecc86f1f888256d32e%26amp%3Bscene%3D21%23wechat_redirect )

**如何通过抓包实战来学习Web协议？** ( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247487946%26amp%3Bidx%3D1%26amp%3Bsn%3Db20fabf4922fa5e5a780692475626b95%26amp%3Bchksm%3Deb477d54dc30f4423156d55f388c90d89f20016ad6395f9ed9d5a26d5fbf22ad83f800097fdc%26amp%3Bscene%3D21%23wechat_redirect )

**动画：一招学会TCP的三次握手和四次挥手** ( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247487654%26amp%3Bidx%3D1%26amp%3Bsn%3Dcd111579a44515ed2138065d24a01f02%26amp%3Bchksm%3Deb477c38dc30f52ede3055544f82830c7b25eb820b0c8fdb27d160d0d51222e0bba6a273e5eb%26amp%3Bscene%3D21%23wechat_redirect )

**上两个月，15家面试，几个offer , 我的面试历程！** ( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247487944%26amp%3Bidx%3D2%26amp%3Bsn%3De8eb7f31e0af297f7dd202f1b06f180b%26amp%3Bchksm%3Deb477d56dc30f4400fa709ae9999359cf9fe66fd956171ee165fb6614bd56eef6c5e2d5612be%26amp%3Bscene%3D21%23wechat_redirect )

#### 相信自己，没有做不到的，只有想不到的 ####

在这里获得的不仅仅是技术！

如果您有什么问题，欢迎阅读上面的文章，关注我微信公众号：终端研发部，一起交流和学习~~

![image](https://user-gold-cdn.xitu.io/2019/5/17/16ac518c800b6b4e?imageView2/0/w/1280/h/960/ignore-error/1)