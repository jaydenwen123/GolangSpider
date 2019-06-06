# spring源码之IOC容器资源定位过程 #

### 铺垫 ###

平时我们在使用spring进行项目实践的时候，对于底层的某些实现逻辑大部分都被忽略过，但是对底层实现的了解却往往是我们在实践中解决出现问题最有利的帮助。以下是我阅读源码进行并整理的一些学习笔记。

### 准备工作 ###

在开发环境中构建自己的spring源码环境，本人是使用IDEA开发工具，构建步骤网上一搜一大把，下面就附上一个在IDEA搭建spring源码环境的链接：

[blog.csdn.net/u011976388/…]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu011976388%2Farticle%2Fdetails%2F80356808 )

### 正文 ###

搭建好环境之后，在读源码之前，我觉得有必要先了解下spring容器类的实现和继承结构，下图是我找到的最全的一张图（可能也是最好看的），有必要的话可以收藏起来。

![](https://user-gold-cdn.xitu.io/2019/6/1/16b10f4caeb26822?imageView2/0/w/1280/h/960/ignore-error/1)

今天我们以FileSystemXmlApplicationContext容器类为例，分析容器初始化的调用流程。

简单来说 spring IoC 容器的初始化的开始是由一个名为refresh()方法来启动的，这个方法标志着IoC容器的正式启动。具体来说，这个启动主要有三个步骤（可以记下）：

* **BeanDefinition的Resource的定位**
* **BeanDefinition的载入**

* **BeanDefinition的注册**

需要提醒的是，spring把这三个过程分开，并使用不同的模块来完成，这样的设计方式，是可以让用户更加灵活的对这三个过程进行裁剪和拓展，定义出最合适自己的IoC容器的初始化过程。

今天我们先分析第一个过程是Resource定位过程，后续文章继续记录剩余的两个加载过程，可以看到初始化的三个过程都是和BeanDefinition这个东西有关联，先来说说BeanDefinition是啥

**BeanDefinition** 是指将我们定义的Bean抽象化，是让容器起作用的主要数据类型，对IoC容器来说，BeanDefinition就是对依赖反转模式中管理的对象依赖关系的抽象数据，也是容器实现反转功能的核心数据结构，依赖反转功能都是围绕这个BeanDefinition的处理来完成的。

以下是我整理出来的资源定位的调用流程图（选择流程图是因为由于继承结构复杂，方法调用过程大部分多态实现，这样定位方法可能比较麻烦，所以使用流程图）

![](https://user-gold-cdn.xitu.io/2019/6/1/16b111ebdbd191cf?imageView2/0/w/1280/h/960/ignore-error/1)

可以根据这个流程图，到前面已经搭建好的spring源码环境中看看具体的操作，不过我图中已经在备注出了每个调用环节的大致内容。

可以看到FileSystemXmlApplicationContext继承了DefaultResourceLoader并重写了getResourceByPath，进而在AbstractBeanDefinitionReader类中获取FileSystemResource实例对象。这个对象代表着我们定义好的配置信息资源，通过这个对象，可以进行相关的IO操作完成BeanDefinition的定位

如果是其他的容器实现类，那么会生成对应种类的Resource，比如ClassPathResource、ServletContextResource等，下图是Resource的继承关系

![](https://user-gold-cdn.xitu.io/2019/6/1/16b1124ab842ea5c?imageView2/0/w/1280/h/960/ignore-error/1)

**总结**

如果我们把IOC容器看成一个水桶，那么可以说BeanFactory就定义了可以作为水桶的基本功能（如至少能装水等），在spring提供的基本IoC容器的接口定义和实现的基础上，spring通过定义BeanDefinition来管理Spring应用中的各种对象，以及对象之间的相互依赖关系，BeanDefinition将我们定义的Bean抽象化，这些BeanDefinition就像容器里装的水，有了这些基本数据，容器就能发挥作用。

我们从前面的实现原理的分析，了解到了Resource定位问题的解决方案，即以FileSystem方式存在的Resource的定位问题，在BeanDefinition定位完成的基础上，就可以通过返回的Resource对象来进行BeanDefinition的载入了，在定位完成以后，为后面的BeanDefinition的载入创造了I/O操作的条件。就相当于我们需要用一个桶去装水，水源的位置已经确定了。

下面是我自己开的一个公众号，主要是记录自己的学习笔记，后期的更新文章也会在对应的公众号同步更新，有兴趣了解的可以关注下。

![](https://user-gold-cdn.xitu.io/2019/6/1/16b1128d76cd9222?imageView2/0/w/1280/h/960/ignore-error/1)