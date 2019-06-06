# 随手记-JVM #

## 1.垃圾回收 ##

#### 1.垃圾回收器 ####

##### 1.新生代 #####

Serial

Parallel New(多线程)

Parallel Scavenge(达到一个可控制的吞吐量Throughput)

> 
> 
> 
> 参数-XX:GCTimeTatio=19(允许最大GC时间就占总时间的5%=1/(1+19),默认是99
> 
> 
> 
> -XX:MaxGCPauseMillis 参数控制最大垃圾收集停顿时间(关注最大停顿时间)
> 
> 
> 
> -XX:+UseAdaptiveSizePolicy GC自适应调节策略,不需要手工指定新生代(-Xmn),Eden 与
> Survivor区的比例(-XX:SurvivorRatio),晋升老年代对象的年龄(-XX:PretenureSizeThreshold)
> 
> 

##### 2.老年代(Tenured generation) #####

Serial Old(标记-压缩算法)

> 
> 
> 
> 可以跟 Serial 、Parallel Scavenge、Parallel New 搭配使用
> 
> 

Parallel Old(标记-压缩算法)

> 
> 
> 
> 只能跟 Parallel scavenge 搭配使用
> 
> 

CMS(标记-清除算法)

> 
> 
> 
> 可以跟 Serial 、Parallel New 搭配使用,和 Serial Old搭配使用
> 
> 
> 
> * 初始标记
> * 并发标记
> * 重新标记
> * 并发清除
> 
> 
> 
> 
> 回收线程的数量=(CPU数量+3)/4
> 
> 
> 
> 两个问题:
> 
> 
>> 
>> 
>> 1.无法处理浮动垃圾(Floating Garbage)
>> 
>> 
>> 
>> 浮动垃圾:当次无法回收的垃圾
>> 
>> 
>> 
>> 可以通过参数-XX:CMSinitiatingOccupancyFraction 控制几时触发回收(默认当老年代使用68%的空间后触发)
>> 
>> 
>> 
>> 2.空间碎片过多
>> 
>> 
>> 
>> -XX:+UseCMSCompactAtFullCollection 当CMS收集器顶不住要进行FullGC时开启内存碎片的合并整理过程(默认开启)
>> 
>> 
>> 
>> 
>> -XX:CMSFullGCsBeforeCompaction 用于设置执行多少次不压缩的Full GC后,跟着来一次带压缩的(默认0)
>> 
>> 
> 
> 

##### 3.G1 #####

横跨新老生代

将整个堆划分为多个大小相等的独立区域(Region),新生代和老年代都是Region(不需要连续)的集合

> 
> 
> 
> * 并行与并发
> * 分代收集
> * 空间整合:整体看是基于 标记-整理算法，从局部(两个Region之间)是基于 复制算法
> * 可预测的停顿
> 
> 
> 

##### 4.Remembered Set #####

每个Region都有对应Remembered Set(避免全堆扫描),虚拟机发现程序在对Reference类型进行写操作时,会产生一个Write Barrier暂时中断写操作,检查Reference引用的对象是否处于不同的Region之中,如果是便通过 CardTable 把相关引用信息记录到被引用对象所属的Region的Remembered Set 之中,当进行内存回收时,在GC根节点的枚举范围中加入Remembered Set 即可保证不对全堆扫描也不会遗漏

##### 5.CardTable #####

> 
> 
> 
> -XX:+UseCondCardMark 来尽量减少写卡表的操作
> 
> 

##### 6.分配担保 #####

##### 7.相关参数 #####

+--------------------------------+---------------------------------------------------------------------------------------------------------+
|              参数              |                                                  描述                                                   |
+--------------------------------+---------------------------------------------------------------------------------------------------------+
| UseSerialGC                    |                                                                                                         |
| UseParNewGC                    | 使用ParNew+Serial                                                                                       |
|                                | Old的收集器组合                                                                                         |
| UseConcMarkSweepGC             | 使用ParNew+CMS+Serial Old                                                                               |
| UseParallelGC                  | 使用Parallel Scavenge + Serial                                                                          |
|                                | Old(PS markSweep)                                                                                       |
| UseParallelOldGC               | 使用Parallel Scavenge+Parallel                                                                          |
|                                | Old                                                                                                     |
| SurvivorRatio                  | 新生代中Eden区域与Survivor区域的容量比值,默认为8，代表Eden:Survivor=8:1                                 |
| PretenureSizeThreshold         | 直接晋升到老年代的对象大小,设置这个参数后，大于这个参数的对象将直接在老年代分配                         |
| MaxTenuringThreshold           | 晋升到老年代的对象年龄，每个对象在坚持过一次Minor                                                       |
|                                | GC 之后，年龄增加1                                                                                      |
| UseAdaptiveSizePolicy          | 动态调整Java堆中各个区域的大小以及进入老年代的年龄(上文提及)                                            |
| HandlePromotionFailure         | 是否允许分配担保失败,即老年代的剩余空间不足以应付新生代的整个Eden和Survivor区的所有对象都存活的极端情况 |
| ParallelGCThreads              | 设置并行GC时进行内存回收的线程数                                                                        |
| GCTimeRatio(仅在Parallel       | GC时间占比总时间的比率,默认值为99，允许1%的GC时间                                                       |
| Scavenge生效)                  |                                                                                                         |
| MaxGCPauseMilis(仅在Parallel   | 设置GC最大停顿时间                                                                                      |
| Scavenge生效)                  |                                                                                                         |
| CMSInitiatingOccupancyFraction | 控制几时触发回收(默认当老年代使用68%的空间后触发)                                                       |
| UseCMSCompactAtFullCollection  | 当CMS收集器顶不住要进行FullGC时开启内存碎片的合并整理过程(默认开启)                                     |
| CMSFullGCsBeforeCompaction     | 用于设置执行多少次不压缩的Full                                                                          |
|                                | GC后,跟着来一次带压缩的(默认0)                                                                          |
+--------------------------------+---------------------------------------------------------------------------------------------------------+

### 2.JVM内存模型(JMM) ###

### 3.运行时内存区域 ###

#### 1.Heap(堆) 公有 ####

#### 2.Method Area(方法区) 公有 ####

存储已被虚拟机加载的类信息、常量、静态变量、即时编译器编译的后的代码数据

##### 1.Runtime Constant Pool(运行时常量池) #####

存放编译期生成的各种 字面量 和 符号引用

注 :Java8中 元数据层(metaspace)

* 

当类元数据的空间占用达到参数MaxMetaspaceSize 设置的值，将会触发堆死亡对象和类加载器的垃圾回收

` -XX:MetaspaceSize -XX:MaxMetaspaceSize -XX:M in MetaspaceFreeRatio -XX:MaxMetaspceFreeRatio 复制代码`

#### 3.VM stack(虚拟机栈) ####

每个方法在执行的同时都会创建一个栈帧(Stack Frame),用于存储 局部变量表 、操作数栈、动态链接、方法出口等信息

#### 4.Program counter Register(程序计数器) ####

字节码的行号指示器

#### 5.Native Method Stack(本地方法栈) ####

### 4.内存布局 ###

#### 1.对象头(Header) ####

##### 1.存储对象自身的运行数据 #####

HashCode、GC分代年龄、锁状态标志、线程持有的锁、偏向线程ID、偏向时间戳

mark Word

+--------------------------------------+--------+------------------+
|               存储内容               | 标志位 |       状态       |
+--------------------------------------+--------+------------------+
| 对象哈希码、对象分代年龄             |     01 | 未锁定           |
| 指向锁记录的指针                     |     00 | 轻量级锁定       |
| 指向重量级锁的指针                   |     10 | 膨胀(重量级锁定) |
| 空、不需要记录信息                   |     11 | GC标记           |
| 偏向线程ID、偏向时间戳、对象分代年龄 |     01 | 可偏向           |
+--------------------------------------+--------+------------------+

##### 2.类型指针 #####

#### 2.实例数据(Instance Data) ####

#### 3.对齐填充(Padding) ####

### 5.线程池 ###

#### 1.参数之间的关系 ####

* 

小于在运行的corePoolSize线程,Executor 会增加一个新的线程

* 

大于或等于corePoolSize,Executor会让提交的请求进入队列

* 

如果请求不能排队,Executor会创建一个新的线程，除非这个线程数大于最大线程数(maximumPoolSize)

* 

超过MaximumPoolSize,就会拒绝

* 

如果当前线程数超过corePoolSize，Executor将会终止超过 KeepAliveTime时间的闲置线程

ThreadPoolExecutor:execute

` /** * Executes the given task sometime in the future. The task * may execute in a new thread or in an existing pooled thread. * * If the task cannot be submitted for execution, either because this * executor has been shutdown or because its capacity has been reached, * the task is handled by the current { @code RejectedExecutionHandler}. * * @param command the task to execute * @throws RejectedExecutionException at discretion of * { @code RejectedExecutionHandler}, if the task * cannot be accepted for execution * @throws NullPointerException if { @code command} is null */ public void execute (Runnable command) { if (command == null ) throw new NullPointerException(); /* * Proceed in 3 steps: * * 1. If fewer than corePoolSize threads are running, try to * start a new thread with the given command as its first * task. The call to addWorker atomically checks runState and * workerCount, and so prevents false alarms that would add * threads when it shouldn't, by returning false. * * 2. If a task can be successfully queued, then we still need * to double-check whether we should have added a thread * (because existing ones died since last checking) or that * the pool shut down since entry into this method. So we * recheck state and if necessary roll back the enqueuing if * stopped, or start a new thread if there are none. * * 3. If we cannot queue task, then we try to add a new * thread. If it fails, we know we are shut down or saturated * and so reject the task. */ int c = ctl.get(); if (workerCountOf(c) < corePoolSize) { if (addWorker(command, true )) return ; c = ctl.get(); } if (isRunning(c) && workQueue.offer(command)) { int recheck = ctl.get(); if (! isRunning(recheck) && remove(command)) reject(command); else if (workerCountOf(recheck) == 0 ) addWorker( null , false ); } else if (!addWorker(command, false )) reject(command); } 复制代码`

### 6.锁 ###

#### 1.synchronized(重入锁) ####

在执行 monitorenter 时,首先尝试获取对象的锁。如果这个对象没被锁定 ，或者当前线程已经拥有了那个对象的锁，把锁的计数器加1;

当执行 monitorexit 指令时会将锁计数器减1，当计数器为0时，锁就会被释放

如果获取对象锁失败，那当前线程就要阻塞等待，直到对象锁被另外一个线程释放位置

#### 2.Lock ####

优势

* 等待可以中断
* 可实现公平锁
* 锁可以绑定多个条件

## 2.类加载 ##

#### 1.双亲委派(Parents Delegation Model) ####

* 一个类加载器收到类加载的请求，首先会交个父类加载器加载;
* 当父类加载器服务完成时，子类加载器才会去尝试加载;
* 双亲委派保证某个类在程序的各个类加载器中都是同一个类
* JVM有三个类加载器:

> 
> 
> 
> * BootStrap classLoader:加载<JAVA_HOME>\lib 目录下或被—Xbootclasspath参数所指定的的路径
> * Extension classLoader:加载lib/ext 目录下,或者被java.ext.dirs 系统变量所指定的路径中所有类库
> * Application classLoader:加载用户类路径(ClassPath)上指定的类库
> 
> 
> 

` //ClassLoader 复制代码`