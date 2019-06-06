# JVM学习（三）JVM常用命令 #

**[个人博客项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FVip-Augus%2FVip-Augus.github.io )**

希望各位帮忙点个star，给我加个小星星✨

本篇记录JVM常用的指令，通过Java的bin目录下强大的工具就能进行查看。

其中很多命令参考option参数，自己要多敲几遍才能记住。

## JVM常用命令 ##

其中[]方括号内的参数，表示可有可无。

### **jps** ###

JVM Process Status Tool,显示指定系统内所有的HotSpot虚拟机进程。

**命令格式**

` jps [option] [hostid] 复制代码`

**option参数**

` -l : 输出主类全名或jar路径 -q : 只输出LVMID -m : 输出JVM启动时传递给main()的参数 -v : 输出JVM启动时显示指定的JVM参数 复制代码`

示例：

` [root@VM_247_254_centos ~]#jps -lm 26176 org.apache.zookeeper.server.quorum.QuorumPeerMain /usr/local/zookeeper-3.4.10/bin/../conf/zoo.cfg 25044 /usr/local/apache-activemq-5.14.5//bin/activemq.jar start 23732 sun.tools.jps.Jps -lm 25446 org.apache.catalina.startup.Bootstrap start 复制代码`

最前面数字表示PID，后面有用到。

## **jstat** ##

jstat(JVM statistics Monitoring)是用于监视虚拟机运行时状态信息的命令，它可以显示出虚拟机进程中的类装载、内存、垃圾收集、JIT编译等运行数据。

**命令格式**

` jstat [option] LVMID [interval] [count] [option] : 操作参数 LVMID : 本地虚拟机进程ID [interval] : 连续输出的时间间隔 [count] : 连续输出的次数 复制代码`

**option参数**

+------------------+------------------------------------------------------------------------+
|      OPTION      |                                  解释                                  |
+------------------+------------------------------------------------------------------------+
| class            | class                                                                  |
|                  | loader的行为统计。Statistics                                           |
|                  | on the behavior of the class                                           |
|                  | loader.                                                                |
| compiler         | HotSpt                                                                 |
|                  | JIT编译器行为统计。Statistics                                          |
|                  | of the behavior of the HotSpot                                         |
|                  | Just-in-Time compiler.                                                 |
| gc               | 垃圾回收堆的行为统计。Statistics                                       |
|                  | of the behavior of the garbage                                         |
|                  | collected heap.                                                        |
| gccapacity       | 各个垃圾回收代容量(young,old,perm)和他们相应的空间统计。Statistics     |
|                  | of the capacities of the generations and their corresponding           |
|                  | spaces.                                                                |
| gcutil           | 垃圾回收统计概述。Summary of                                           |
|                  | garbage collection statistics.                                         |
| gccause          | 垃圾收集统计概述（同-gcutil），附加最近两次垃圾回收事件的原因。Summary |
|                  | of garbage collection statistics (same as -gcutil), with the cause of  |
|                  | the last and                                                           |
| gcnew            | 新生代行为统计。Statistics                                             |
|                  | of the behavior of the new                                             |
|                  | generation.                                                            |
| gcnewcapacity    | 新生代与其相应的内存空间的统计。Statistics                             |
|                  | of the sizes of the new generations and                                |
|                  | its corresponding spaces.                                              |
| gcold            | 年老代和永生代行为统计。Statistics                                     |
|                  | of the behavior of the old and                                         |
|                  | permanent generations.                                                 |
| gcoldcapacity    | 年老代行为统计。Statistics                                             |
|                  | of the sizes of the old                                                |
|                  | generation.                                                            |
| gcpermcapacity   | 永生代行为统计。Statistics                                             |
|                  | of the sizes of the permanent                                          |
|                  | generation.                                                            |
| printcompilation | HotSpot编译方法统计。HotSpot                                           |
|                  | compilation method statistics.                                         |
+------------------+------------------------------------------------------------------------+

**例如查看垃圾回收堆的行为统计**

` [root@VM_247_254_centos ~] # jstat -gc 25446 S0C S1C S0U S1U EC EU OC OU MC MU CCSC CCSU YGC YGCT FGC FGCT GCT 6400.0 6400.0 0.0 1601.4 51712.0 50837.0 128808.0 88450.0 67584.0 66167.7 7936.0 7630.5 401 5.939 10 1.247 7.186 复制代码`

顺便介绍一下参数意义： **C：Capacity表示的是容量 U：Used表示的是已使用**

` S0C : survivor0区的总容量 S1C : survivor1区的总容量 S0U : survivor0区已使用的容量 S1C : survivor1区已使用的容量 EC : Eden区的总容量 EU : Eden区已使用的容量 OC : Old区的总容量 OU : Old区已使用的容量 PC 当前perm的容量 (KB) PU perm的使用 (KB) YGC : 新生代垃圾回收次数 YGCT : 新生代垃圾回收时间 FGC : 老年代垃圾回收次数 FGCT : 老年代垃圾回收时间 GCT : 垃圾回收总消耗时间 复制代码`

上面有很多个option参数，可以一个个敲过去看看具体的作用。

## **jmap** ##

jmap(JVM Memory Map)命令用于生成heap dump文件，如果不使用这个命令，可以使用-XX:+HeapDumpOnOutOfMemoryError参数来让虚拟机出现OOM的时候，自动生成dump文件。 jmap不仅能生成dump文件，还可以查询finalize执行队列、Java堆和永久代的详细信息，如当前使用率、当前使用的是哪种收集器等。

**命令格式**

` jmap [option] LVMID 复制代码`

**option参数**

` dump : 生成堆转储快照 finalizerinfo : 显示在F-Queue队列等待Finalizer线程执行finalizer方法的对象 heap : 显示Java堆详细信息 histo : 显示堆中对象的统计信息 permstat : to print permanent generation statistics F : 当-dump没有响应时，强制生成dump快照 复制代码`

**举个🌰**

* **-dump**

` jmap -dump:format=b,file=dump.dprof 25446 Dumping heap to /home/gem/dump.dprof ... Heap dump file created 复制代码`

**输出.dprof文件后，使用MAT分析工具进行分析**

* **-heap** 打印heap的概要信息，GC使用的算法，heap的配置及wise heap的使用情况,可以用此来判断内存目前的使用情况以及垃圾回收情况

` jmap -heap 25446 Attaching to process ID 25446, please wait... Debugger attached successfully. Server compiler detected. JVM version is 25.121-b13 using thread-local object allocation. Mark Sweep Compact GC //GC 方式 ,该次是标记-清理算法，上一篇有记录哦 //堆内存初始化配置 Heap Configuration: //对应jvm启动参数-XX:M in HeapFreeRatio设置JVM堆最小空闲比率(default 40) M in HeapFreeRatio = 40 //对应jvm启动参数 -XX:MaxHeapFreeRatio设置JVM堆最大空闲比率(default 70) MaxHeapFreeRatio = 70 //对应jvm启动参数-XX:MaxHeapSize=设置JVM堆的最大大小 MaxHeapSize = 262144000 (250.0MB) //对应jvm启动参数-XX:NewSize=设置JVM堆的‘新生代’的默认大小 NewSize = 5570560 (5.3125MB) //对应jvm启动参数-XX:MaxNewSize=设置JVM堆的‘新生代’（YG）的最大大小 MaxNewSize = 87359488 (83.3125MB) //对应jvm启动参数-XX:OldSize=<value>:设置JVM堆的‘老生代’（OG）的大小 OldSize = 11206656 (10.6875MB) //对应jvm启动参数-XX:NewRatio=:‘新生代’和‘老年代’的大小比率 NewRatio = 2 //对应jvm启动参数-XX:SurvivorRatio=设置年轻代中Eden区与Survivor区的大小比值 SurvivorRatio = 8 //元空间大小，对应-XX:MetaspaceSize，初始空间大小 //达到该值就会触发垃圾收集进行类型卸载，同时GC会对该值进行调整 //JDK 8 中永久代向元空间的转换 MetaspaceSize = 21807104 (20.796875MB) //只有当-XX:+UseCompressedClassPointers开启了才有效 //通过java -XX:+PrintFlagsInitial | grep UseCompressedClassPointers //发现bool UseCompressedClassPointers= false ，是没有启用的 CompressedClassSpaceSize = 1073741824 (1024.0MB) //对应启动参数-XX:MaxMetaspaceSize对应元空间最大大小 MaxMetaspaceSize = 17592186044415 MB //当使用G1收集器时，设置java堆被分割的大小。这个大小范围在1M到32M之间。 //可能我这个JVM没有启用G1收集器，所以为0 G1HeapRegionSize = 0 (0.0MB) //堆内存使用情况 Heap Usage: //新的复制算法，一个伊甸区+Survivor区 New Generation (Eden + 1 Survivor Space): capacity = 59506688 (56.75MB) used = 17941224 (17.110084533691406MB) free = 41565464 (39.639915466308594MB) 30.14992869372935% used //Eden区内存分布 Eden Space: capacity = 52953088 (50.5MB) used = 17935840 (17.104949951171875MB) free = 35017248 (33.395050048828125MB) 33.871188022122524% used //其中一个Survivor区的内存分布 From Space: capacity = 6553600 (6.25MB) used = 5384 (0.00513458251953125MB) free = 6548216 (6.244865417480469MB) 0.0821533203125% used //另一个Survivor区的内存分布 To Space: capacity = 6553600 (6.25MB) used = 0 (0.0MB) free = 6553600 (6.25MB) 0.0% used //‘当前’老年代内存分布 tenured generation: capacity = 131899392 (125.7890625MB) used = 94610832 (90.22792053222656MB) free = 37288560 (35.56114196777344MB) 71.7295436812931% used 27405 interned Strings occupying 3101144 bytes. 复制代码`

## **jhat** ##

jhat(JVM Heap Analysis Tool)命令是与jmap搭配使用，用来分析jmap生成的dump，jhat内置了一个微型的HTTP/HTML服务器，生成dump的分析结果后，可以在浏览器中查看。在此要注意，一般不会直接在服务器上进行分析，因为jhat是一个耗时并且耗费硬件资源的过程，一般把服务器生成的dump文件复制到本地或其他机器上进行分析。

因为常用分析方法是用各平台通用的MAT进行分析，就不具体在远程服务器操作展示效果了。

## **jstack** ##

jstack用于生成java虚拟机当前时刻的线程快照。 线程快照是当前java虚拟机内每一条线程正在执行的方法堆栈的集合，生成线程快照的主要目的是定位线程出现长时间停顿的原因，如线程间死锁、死循环、请求外部资源导致的长时间等待等。 线程出现停顿的时候通过jstack来查看各个线程的调用堆栈，就可以知道没有响应的线程到底在后台做什么事情，或者等待什么资源。 如果java程序崩溃生成core文件，jstack工具可以用来获得core文件的java stack和native stack的信息，从而可以轻松地知道java程序是如何崩溃和在程序何处发生问题。另外，jstack工具还可以附属到正在运行的java程序中，看到当时运行的java程序的java stack和native stack的信息, 如果现在运行的java程序呈现hung的状态，jstack是非常有用的。

**命令格式**

` jstack [option] LVMID 复制代码`

**option参数**

` -F : 当正常输出请求不被响应时，强制输出线程堆栈 -l : 除堆栈外，显示关于锁的附加信息 -m : 如果调用到本地方法的话，可以显示C/C++的堆栈 复制代码`

**举个例子**

` jstack -l 25446 | more 2018-01-25 21:18:22 Full thread dump Java HotSpot(TM) 64-Bit Server VM (25.121-b13 mixed mode): "main-EventThread" #174 daemon prio=5 os_prio=0 tid=0x00007feb692b7000 nid=0x6502 waiting on condition [0x00007feb32bb1000] //等待 java.lang.Thread.State: WAITING (parking) at sun.misc.Unsafe.park(Native Method) - parking to wait for <0x00000000f0a00860> (a java.util.concurrent.locks.AbstractQueuedSynchronizer $ConditionObject ) at java.util.concurrent.locks.LockSupport.park(LockSupport.java:175) at java.util.concurrent.locks.AbstractQueuedSynchronizer $ConditionObject.await(AbstractQueuedSynchronizer.java:2039) at java.util.concurrent.LinkedBlockingQueue.take(LinkedBlockingQueue.java:442) at org.apache.zookeeper.ClientCnxn $EventThread.run(ClientCnxn.java:501) Locked ownable synchronizers: - None "main-SendThread(115.159.192.69:2181)" #173 daemon prio=5 os_prio=0 tid=0x00007feb694bd800 nid=0x6501 runnable [0x00007feb363ce000] //运行 java.lang.Thread.State: RUNNABLE at sun.nio.ch.EPollArrayWrapper.epollWait(Native Method) at sun.nio.ch.EPollArrayWrapper.poll(EPollArrayWrapper.java:269) at sun.nio.ch.EPollSelectorImpl.doSelect(EPollSelectorImpl.java:93) at sun.nio.ch.SelectorImpl.lockAndDoSelect(SelectorImpl.java:86) - locked <0x00000000f09ef268> (a sun.nio.ch.Util $3 ) - locked <0x00000000f09ef258> (a java.util.Collections $UnmodifiableSet ) - locked <0x00000000f09ef140> (a sun.nio.ch.EPollSelectorImpl) at sun.nio.ch.SelectorImpl.select(SelectorImpl.java:97) at org.apache.zookeeper.ClientCnxnSocketNIO.doTransport(ClientCnxnSocketNIO.java:349) at org.apache.zookeeper.ClientCnxn $SendThread.run(ClientCnxn.java:1141) Locked ownable synchronizers: - None 复制代码`

## **jinfo** ##

jinfo(JVM Configuration info)这个命令作用是实时查看和调整虚拟机运行参数。 之前的jps -v口令只能查看到显示指定的参数，如果想要查看未被显示指定的参数的值就要使用jinfo口令

一般我常用它来看JVM启动时的参数

**命令格式**

` jinfo [option] [args] LVMID 复制代码`

**optin参数**

` -flag : 输出指定args参数的值 -flags : 不需要args参数，输出所有JVM参数的值 -sysprops : 输出系统属性，等同于System.getProperties() 复制代码`

**举个🌰**

` [root@VM_247_254_centos ~] # jinfo -flags 25446 Attaching to process ID 25446, please wait... Debugger attached successfully. Server compiler detected. JVM version is 25.121-b13 Non-default VM flags: -XX:CICompilerCount=2 -XX:InitialHeapSize=16777216 -XX:MaxHeapSize=262144000 -XX:MaxNewSize=87359488 -XX:M in HeapDeltaBytes=196608 -XX:NewSize=5570560 -XX:OldSize=11206656 -XX:+UseCompressedClassPointers -XX:+UseCompressedOops Command line: -Djava.util.logging.config.file=/usr/ local /tomcat/conf/logging.properties -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager -Djdk.tls.ephemeralDHKeySize=2048 -Djava.protocol.handler.pkgs=org.apache.catalina.webresources -Dcatalina.base=/usr/ local /tomcat -Dcatalina.home=/usr/ local /tomcat -Djava.io.tmpdir=/usr/ local /tomcat/temp 复制代码`

从上面可以看到，jinfo打印出来的参数，下一篇原本想写MAT的使用，但是自己手动制造过异常，异常信息集中在Java Objects，只要联想到跟上一次改动过的代码，就能发现大对象可能出现的地方，所以具体分析的话，等之后拿到比较复杂的dump文件再具体学习。

## 参考资料 ##

1、 [《jvm系列(四):jvm调优-命令大全（jps jstat jmap jhat jstack jinfo）》]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fityouknow%2Fp%2F5714703.html )

2、 [Java8内存模型—永久代(PermGen)和元空间(Metaspace)]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fpaddix%2Fp%2F5309550.html )

3、周志明的《深入理解JVM》