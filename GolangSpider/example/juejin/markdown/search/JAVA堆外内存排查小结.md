# JAVA堆外内存排查小结 #

![timg.jpeg](https://user-gold-cdn.xitu.io/2019/3/31/169d34989fadf0eb?imageView2/0/w/1280/h/960/ignore-error/1)

# 简介 #

> 
> 
> 
> ` JVM` 堆外内存难排查但经常会出现问题，这可能是目前最全的JVM堆外内存排查思路。
> 
> 

通过本文，你应该了解：

* pmap 命令
* gdb 命令
* perf 命令
* 内存 RSS、VSZ的区别
* java NMT

# 起因 #

这几天遇到一个比较奇怪的问题，觉得有必要和大家分享一下。我们的一个服务，运行在 ` docker` 上，在某个版本之后，占用的内存开始增长，直到 ` docker` 分配的内存上限，但是并不会 ` OOM` 。版本的更改如下：

* 升级了基础软件的版本
* 将docker的内存上限由4GB扩展到8GB
* 上上个版本的一项变动是使用了EhCache的Heap缓存
* 没有读文件，也没有mmap操作

使用 ` jps` 查看启动参数，发现分配了大约3GB的堆内存

` [root]$ jps -v 75 Bootstrap -Xmx3000m -Xms3000m -verbose:gc -Xloggc:/home/logs/gc.log -XX:CMSInitiatingOccupancyFraction=80 -XX:+UseCMSCompactAtFullCollection -XX:MaxTenuringThreshold=10 -XX:MaxPermSize=128M -XX:SurvivorRatio=3 -XX:NewRatio=2 -XX:+PrintGCDateStamps -XX:+PrintGCDetails -XX:+UseParNewGC -XX:+UseConcMarkSweepGC 复制代码`

使用ps查看进程使用的内存和虚拟内存 ( Linux内存管理 )。除了虚拟内存比较高达到 ` 17GB` 以外，实际使用的内存RSS也夸张的达到了7GB，远远超过了 ` -Xmx` 的设定。

` [root]$ ps -p 75 -o rss,vsz RSS VSZ 7152568 17485844 复制代码`

# 排查过程 #

明显的，是有堆外内存的使用，不太可能是由于 ` EhCache` 引起的（因为我们使用了heap方式）。了解到基础软件的升级涉及到netty版本升级，netty会用到一些 ` DirectByteBuffer` ，第一轮排查我们采用如下方式：

* ` jmap -dump:format=b,file=75.dump 75` 通过分析堆内存找到 ` DirectByteBuffer` 的引用和大小
* 部署一个升级基础软件之前的版本，持续观察
* 部署另一个版本，更改EhCache限制其大小到1024M
* 考虑到可能由Docker的内存分配机制引起，部署一实例到实体机

结果4个环境中的服务，无一例外的都出现了内存超用的问题。问题很奇怪，宝宝睡不着觉。

## pmap ##

为了进一步分析问题，我们使用pmap查看进程的内存分配，通过RSS升序序排列。结果发现除了地址 ` 000000073c800000` 上分配的3GB堆以外，还有数量非常多的64M一块的内存段，还有巨量小的物理内存块映射到不同的虚拟内存段上。但到现在为止，我们不知道里面的内容是什么，是通过什么产生的。

` [root]$ pmap -x 75 | sort -n -k3 .....省略N行 0000000040626000 55488 55484 55484 rwx-- [ anon ] 00007fa07c000000 65536 55820 55820 rwx-- [ anon ] 00007fa044000000 65536 55896 55896 rwx-- [ anon ] 00007fa0c0000000 65536 56304 56304 rwx-- [ anon ] 00007f9db8000000 65536 56360 56360 rwx-- [ anon ] 00007fa0b8000000 65536 56836 56836 rwx-- [ anon ] 00007fa084000000 65536 57916 57916 rwx-- [ anon ] 00007f9ec4000000 65532 59752 59752 rwx-- [ anon ] 00007fa008000000 65536 60012 60012 rwx-- [ anon ] 00007f9e58000000 65536 61608 61608 rwx-- [ anon ] 00007f9f18000000 65532 61732 61732 rwx-- [ anon ] 00007fa018000000 65532 61928 61928 rwx-- [ anon ] 00007fa088000000 65536 62336 62336 rwx-- [ anon ] 00007fa020000000 65536 62428 62428 rwx-- [ anon ] 00007f9e44000000 65536 64352 64352 rwx-- [ anon ] 00007f9ec0000000 65528 64928 64928 rwx-- [ anon ] 00007fa050000000 65532 65424 65424 rwx-- [ anon ] 00007f9e08000000 65512 65472 65472 rwx-- [ anon ] 00007f9de0000000 65524 65512 65512 rwx-- [ anon ] 00007f9dec000000 65532 65532 65532 rwx-- [ anon ] 00007f9dac000000 65536 65536 65536 rwx-- [ anon ] 00007f9dc8000000 65536 65536 65536 rwx-- [ anon ] 00007f9e30000000 65536 65536 65536 rwx-- [ anon ] 00007f9eb4000000 65536 65536 65536 rwx-- [ anon ] 00007fa030000000 65536 65536 65536 rwx-- [ anon ] 00007fa0b0000000 65536 65536 65536 rwx-- [ anon ] 000000073c800000 3119140 2488596 2487228 rwx-- [ anon ] total kB 17629516 7384476 7377520 复制代码`

通过google，找到以下资料 Linux glibc >= 2.10 (RHEL 6) malloc may show excessive virtual memory usage)

文章指出造成应用程序大量申请64M大内存块的原因是由Glibc的一个版本升级引起的，通过 ` export MALLOC_ARENA_MAX=4` 可以解决VSZ占用过高的问题。虽然这也是一个问题，但却不是我们想要的，因为我们增长的是物理内存，而不是虚拟内存。

## NMT ##

幸运的是 JDK1.8有 ` Native Memory Tracker` 可以帮助定位。通过在启动参数上加入 ` -XX:NativeMemoryTracking=detail` 就可以启用。在命令行执行jcmd可查看内存分配。

` #jcmd 75 VM.native_memory summary Native Memory Tracking: Total: reserved=5074027KB, committed=3798707KB - Java Heap (reserved=3072000KB, committed=3072000KB) (mmap: reserved=3072000KB, committed=3072000KB) - Class (reserved=1075949KB, committed=28973KB) (classes #4819) (malloc=749KB #13158) (mmap: reserved=1075200KB, committed=28224KB) - Thread (reserved=484222KB, committed=484222KB) (thread #470) (stack: reserved=482132KB, committed=482132KB) (malloc=1541KB #2371) (arena=550KB #938) - Code (reserved=253414KB, committed=25070KB) (malloc=3814KB #5593) (mmap: reserved=249600KB, committed=21256KB) - GC (reserved=64102KB, committed=64102KB) (malloc=54094KB #255) (mmap: reserved=10008KB, committed=10008KB) - Compiler (reserved=542KB, committed=542KB) (malloc=411KB #543) (arena=131KB #3) - Internal (reserved=50582KB, committed=50582KB) (malloc=50550KB #13713) (mmap: reserved=32KB, committed=32KB) - Symbol (reserved=6384KB, committed=6384KB) (malloc=4266KB #31727) (arena=2118KB #1) - Native Memory Tracking (reserved=1325KB, committed=1325KB) (malloc=208KB #3083) (tracking overhead=1117KB) - Arena Chunk (reserved=231KB, committed=231KB) (malloc=231KB) - Unknown (reserved=65276KB, committed=65276KB) (mmap: reserved=65276KB, committed=65276KB) 复制代码`

虽然pmap得到的内存地址和 ` NMT` 大体能对的上，但仍然有不少内存去向成谜。虽然是个好工具但问题并不能解决。

## gdb ##

非常好奇64M或者其他小内存块中是什么内容，接下来通过 ` gdb` dump出来。读取 ` /proc` 目录下的maps文件，能精准的知晓目前进程的内存分布。

以下脚本通过传入进程id，能够将所关联的内存全部dump到文件中（会影响服务，慎用）。

` grep rw-p /proc/ $1 /maps | sed -n 's/^\([0-9a-f]*\)-\([0-9a-f]*\) .*$/\1 \2/p' | while read start stop; do gdb --batch --pid $1 -ex "dump memory $1 - $start - $stop.dump 0x $start 0x $stop " ; done 复制代码`

更多时候，推荐之dump一部分内存。(再次提醒操作会影响服务，注意dump的内存块大小，慎用)。

` gdb --batch --pid 75 -ex "dump memory a.dump 0x7f2bceda1000 0x7f2bcef2b000 复制代码` ` [root]$ du -h * dump 4.0K 55-00600000-00601000.dump 400K 55-00eb7000-00f1b000.dump 0 55-704800000-7c0352000.dump 47M 55-7f2840000000-7f2842eb8000.dump 53M 55-7f2848000000-7f284b467000.dump 64M 55-7f284c000000-7f284fffa000.dump 64M 55-7f2854000000-7f2857fff000.dump 64M 55-7f285c000000-7f2860000000.dump 64M 55-7f2864000000-7f2867ffd000.dump 1016K 55-7f286a024000-7f286a122000.dump 1016K 55-7f286a62a000-7f286a728000.dump 1016K 55-7f286d559000-7f286d657000.dump 复制代码`

是时候查看里面的内容了

` [root]$ view 55-7f284c000000-7f284fffa000.dump ^@^@X+^?^@^@^@^@^@d(^?^@^@^@ ÿ^C^@^@^@^@^@ ÿ^C^@^@^@^@^@^@^@^@^@^@^@^@±<97>p^C^@^@^@^@ 8^^Z+^?^@^@ ^@^@d(^?^@^@ 8^^Z+^?^@^@ ^@^@d(^?^@^@ achine ":524993642," timeSecond ":1460272569," inc ":2145712868," new ":false}," device ":{" client ":" android "," uid ":" xxxxx "," version ":881}," device_android ":{" BootSerialno ":" xxxxx "," CpuInfo ":" 0-7 "," MacInfo ":" 2c:5b:b8:b0:d5:10 "," RAMSize ":" 4027212 "," SdcardInfo ":" xxxx "," Serialno ":" xxxx ", " android_id ":" 488aedba19097476 "," buildnumber ":" KTU84P/1416486236 "," device_ip ":" 0.0.0.0 "," mac ":" 2c:5b:b8:b0:d5:10 "," market_source ":" 12 "," model ":" OPPO ...more 复制代码`

纳尼？这些内容不应该在堆里面么？为何还会使用额外的内存进行分配？上面已经排查netty申请directbuffer的原因了，那么还有什么地方在分配堆外内存呢？

## perf ##

传统工具失灵，快到了黔驴技穷的时候了，是时候祭出神器perf了。

使用 ` perf record -g -p 55` 开启监控栈函数调用。运行一段时间后Ctrl+C结束，会生成一个文件perf.data。

执行 ` perf report -i perf.data` 查看报告。

![1.jpeg](https://user-gold-cdn.xitu.io/2019/3/31/169d34989cabd9bc?imageView2/0/w/1280/h/960/ignore-error/1)

如图，进程大量执行bzip相关函数。搜索zip，结果如下：

![2.jpeg](https://user-gold-cdn.xitu.io/2019/3/31/169d34989cc6765d?imageView2/0/w/1280/h/960/ignore-error/1)

-.-!

进程调用了Java_java_util_zip_Inflater_inflatBytes() 申请了内存，仅有一小部分调用Deflater释放内存。与pmap内存地址相比对，确实是bzip在搞鬼。

# 解决 #

java项目搜索zip定位到代码，发现确实有相关bzip压缩解压操作，而且 ` GZIPInputStream` 有个地方没有close。

GZIPInputStream使用Inflater申请堆外内存，Deflater释放内存，调用close()方法来主动释放。如果忘记关闭，Inflater对象的生命会延续到下一次GC。在此过程中，堆外内存会一直增长。

原代码：

` public byte [] decompress ( byte [] input) throws IOException { ByteArrayOutputStream out = new ByteArrayOutputStream(); IOUtils.copy( new GZIPInputStream( new ByteArrayInputStream(input)), out); return out.toByteArray(); } 复制代码`

修改后：

` public byte [] decompress( byte [] input) throws IOException { ByteArrayOutputStream out = new ByteArrayOutputStream(); GZIPInputStream gzip = new GZIPInputStream( new ByteArrayInputStream(input)); IOUtils.copy(gzip, out); gzip.close(); return out.toByteArray(); } 复制代码`

经观察，问题解决。

![](https://user-gold-cdn.xitu.io/2019/3/31/169d34c755ff9124?imageView2/0/w/1280/h/960/ignore-error/1)