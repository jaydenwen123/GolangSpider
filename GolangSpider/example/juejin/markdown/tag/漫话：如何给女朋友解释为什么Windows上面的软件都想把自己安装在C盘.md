# 漫话：如何给女朋友解释为什么Windows上面的软件都想把自己安装在C盘 #

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0b8e381f0?imageView2/0/w/1280/h/960/ignore-error/1)

周末，我在家里面看电视，女朋友正在旁边鼓捣她的电脑，但是好像并不是很顺利，于是就有了以下对话。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0b8bf9054?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0b8f9d488?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0baf6366b?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0b9082736?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0bae8247a?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0d83d496f?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0db88a7eb?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0dabd219e?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0dc012b31?imageView2/0/w/1280/h/960/ignore-error/1)

计算机存储

我们使用的计算机中，保存信息的介质有两类：

一类是内部存储器，一断电就会把记住的东西丢失。

一类是外部存储器，断了电也能存住。

内部存储器，就是我们通常说的内存，内存的信息存取速度很快，但是通常容量较小，并且依赖电源，断电后其中存储的内容就会丢失。内部存储器包括寄存器、高速缓冲存储器（Cache）和主存储器。

另外一种不依赖电源的外部存储器相对内存来说，容量会大一些，但是存取速度会相对慢一点。常见的外存储器包括磁盘、光盘、U盘等。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0dbfcafcb?imageView2/0/w/1280/h/960/ignore-error/1)
￼
从冯.诺依曼的存储程序工作原理及计算机的组成来说，计算机分为运算器、控制器、存储器和输入/输出设备，这里的存储器就是指内部存储器，而硬盘等外部存储器属于输入/输出设备。

CPU运算所需要的程序代码和数据来自于内存，内存中的东西则来自于磁盘，所以磁盘并不直接与CPU打交道。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0de91f963?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0f36bcfa9?imageView2/0/w/1280/h/960/ignore-error/1)

磁盘

磁盘有软磁盘和硬磁盘两种，就是我们通常说的软盘和硬盘。

根据登上历史舞台的先后顺序我们来见识一下软盘和硬盘

**软盘**

在计算机刚诞生的年代，还没有硬盘，那时数据存储主要靠软盘。

软盘（Floppy Disk）是个人计算机（PC）中最早使用的可移介质。软盘的读写是通过软盘驱动器完成的。

软盘在早期计算机上必备的一个硬件，也是计算机上面最早使用的可移介质。它作为一种可移储存硬件适用于一些需要被物理移动的小文件，软盘的读写是用过软驱也就是软盘驱动器来完成的。

软盘驱动器（floppy disk driver）就是平常所说的“软驱”，它是读取软盘的设备。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0f9ecdc26?imageView2/0/w/1280/h/960/ignore-error/1)
￼
软盘存储在20世纪80至90年代盛行，直至2000年以前，3.5英寸软盘仍是电脑普及设备之一。

所以在早期的DOS计算机上经常能够看到如下信息：

` ·Please insert source disk into drive A:... ·Please insert destination disk into drive A:... ·Please insert source disk into drive A:... 复制代码`

软盘想要被读取到计算机中，就需要映射到计算机中的某一个标识，于是字母“A”就作为第一个盘符被软盘驱动器所占用，而随后更多的计算机开始配备第二个软驱，以满足数据拷贝的需要，所以盘符B也被软驱给占据了。

所以软盘驱动器按照顺序占据了A和B盘符的位置：A盘就是的3.5英寸软盘驱动器、B盘就是的5.25英寸软盘驱动器。

而后来的Windows系统也沿用DOS下分区的设置。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0fbfc97dc?imageView2/0/w/1280/h/960/ignore-error/1)
￼
> 
> 
> 
> A盘的真正含义是“第一软盘驱动器”，并非单指3.5英寸软驱或软盘。实际上，最早的软盘是8英寸软盘，因此，最早期的A盘其实是8英寸软驱。但是，8英寸软盘由于携带不方便等原因，很快被5.25英寸软盘取代，后来出现了一台PC配2个软驱的情况，因此有了A盘和B盘的区分，但这两者都是5.25英寸软驱。后来3.5英寸软盘的推出，3.5英寸和5.25英寸两种软盘开始共存，于是PC上的两个软驱，一个是3.5英寸软驱(通常是A盘)，另一个是5.25英寸软驱(通常是B盘，也有的正好相反)。
> 
> 
> 

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb1000eda15?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb0fc09c0ba?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb110b98aa4?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb11b19ca2b?imageView2/0/w/1280/h/960/ignore-error/1)

**硬盘**

3.5英寸软盘在80至90年代曾盛极一时，1996年时全球有多达50亿只软盘正在使用。直到CD-ROM、USB存储设备出现后，软盘销量逐渐下滑。

1998年苹果推出第一代iMac，是第一台舍弃软式磁盘驱动器的电脑，戴尔2003年推出的Dimension台式机也放弃了软盘支持。之后，标配软驱的新电脑越来越少。

取而代之作为计算机中主要的外部存储器的是硬盘。硬盘是电脑主要的存储媒介之一，由一个或者多个铝制或者玻璃制的碟片组成。碟片外覆盖有铁磁性材料。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb120f0ec2e?imageView2/0/w/1280/h/960/ignore-error/1)
￼

随着硬盘被研发出来，早期的计算机就开始考虑如何兼容硬盘，想要兼容硬盘，最先考虑的就是要给硬盘划分醒的分区。而A和B两个字母命名的分期已经被软盘占用了，所以硬盘只能从C开始。

而随着硬盘技术的发展，一方面软盘逐渐退出历史舞台，另外一方面硬盘开始支持多个分区，于是，就演变成今天我们看到的计算机中有多个分区，从C开始，分别是C、D、E等。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb11c223041?imageView2/0/w/1280/h/960/ignore-error/1)
￼

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb137da8d62?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb12ec8ea6f?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb125f38929?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb137e14af2?imageView2/0/w/1280/h/960/ignore-error/1)

分区

硬盘分区实质上是对硬盘的一种格式化，然后才能使用硬盘保存各种信息。在硬盘中，一般先划分出一个主分区（活动分区），一般来说，这个就是划出的C盘。然后建立扩展分区，在扩展分区内，再建立若干个逻辑分区，这些逻辑分区就是后面的D、E等盘。

所以，很多新买的windows计算机中，至少都会有一个C盘。

因为只要电脑中安装了硬盘，默认情况下都会有C盘，所以软件初始安装位置设定为C盘的话可以避免出现无此分区的情况。

其实，软件安装的时候，默认选择的是系统盘的Program Files目录下（环境变量：%programfiles%），只不过大多数情况下系统盘恰好是C盘而已。

还有另外一个原因，那就是把软件安装在C盘的话，会更加流畅一些。

对于机械硬盘的数据读取，硬盘的主轴的工作方式都是CAV（Constant Angular Velocity,恒定角速度，单位时间内放置的角度一致），所以在相同时间内，读取位于硬盘外圈的数据，比读取硬盘内圈的数据要多。

换句话说，读取相同大小的数据，数据位于硬盘外圈的读取时间比位于内圈的速度时间要短，也就是外圈读取速度快。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb13e7a414a?imageView2/0/w/1280/h/960/ignore-error/1)
￼而按照正常的分区方法，C盘一般位于硬盘外圈，C盘后的D、E、F逐渐向内。所以，C盘的读取速度会相对快一些。

当然，以上只针对机械硬盘，目前已经非常普遍的固态硬盘就没有这种情况了，所以，如果你用的是SSD（固态硬盘），那么就随意吧。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb14624c158?imageView2/0/w/1280/h/960/ignore-error/1)
￼

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb1472824f2?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb156126574?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb1512e9a70?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb157f8bb25?imageView2/0/w/1280/h/960/ignore-error/1)

C盘太满系统会卡？

影响系统速度的原因有很多，硬件上就有两个重要的部分：CPU（处理器）和内存。CPU不用说，相当于大脑，处理所有运算；而内存就是运行程序的场所。

在以前，电脑的配置普遍不太高，CPU计算效率低下，快速运行本就很难，尤其是内存空间还紧张。不过windows系统有个办法，会根据内存情况调用虚拟内存来使用。而C盘恰恰就是虚拟内存的所在地，如果C盘满了，也就没有虚拟内存的空间。内存兄弟只能凭借自己的小身板硬抗，当运行多个程序时，就容易导致电脑卡慢甚至崩溃。

实际上，上面说的情况仅仅是历史遗留问题，现在的电脑在硬件配置上已经足够强大，并且系统会有充足的空间合理分配虚拟内存，所以上述情况基本不存在了。

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb16aaa0dff?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb15c2c97b1?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb1716c2ade?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb17399f0a5?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/3/18/1698edb8b44e54b9?imageView2/0/w/1280/h/960/ignore-error/1)