# 浅谈分布式定时任务之quartz #

# 前言 #

最近一段时间因公司项目需要进行分布式定时任务框架选型，由于资源（人力，时间）有限，所以重点考虑采用开源的一些解决方案，其中重点比较了3款框架：quartz，elastic-job，xxl-job等。由于elastic-job 和xxl-job 实际上也是基于quartz实现的。所以很有必要对quartz 有一定的了解。所以近期阅读了大量quartz相关的一些文档和博客，以及写了一些demo进行验证。本文主要是对近期工作的一个小结，简要的介绍一下quartz 这个分布式定时任务调度框架。

## 目录 ##

* 什么是quartz?它有哪些特性?
* quartz 核心概念和元素介绍
* quartz 常用配置项介绍
* quartz 集群介绍

### 一.什么是quartz?它有哪些特性? ###

` quartz 是一个开源的分布式调度库，它基于java实现。 > 它有着强大的调度功能，支持丰富多样的调度方式,比如简单调度，基于cron表达式的调度等等。 > 支持调度任务的多种持久化方式。比如支持内存存储，数据库存储，Terracotta server 存储。 > 支持分布式和集群能力。 > 采用JDBCJobStore方式存储时，针对事务的处理方式支持全局事务（和业务服务共享同一个事务）和局部事务（quarzt 单独管理自己的事务） > 基于plugin机制以及listener机制支持灵活的扩展。 复制代码`

### 二. quartz 核心概念和元素介绍: ###

* 

Job： 工作任务。它提供了一个接口，只有一个方法void execute(JobExecutionContext context)，开发者实现该接口定义运行任务，JobExecutionContext类提供了调度上下文的各种信息。Job运行时的信息保存在JobDataMap实例中.

* 

JobDetail：工作任务实例.Quartz在每次执行Job时，都重新创建一个Job实例，所以它不直接接受一个Job的实例，相反它接收一个Job实现类，以便运行时通过newInstance()的反射机制实例化Job。因此需要通过一个类来描述Job的实现类及其它相关的静态信息，如Job名字、描述、关联监听器等信息，JobDetail承担了这一角色.

* 

Trigger：触发器,用来描述触发Job执行的时间触发规则。主要有SimpleTrigger和CronTrigger这两个子类。当仅需触发一次或者以固定时间间隔周期执行，SimpleTrigger是最适合的选择；而CronTrigger则可以通过Cron表达式定义出各种复杂时间规则的调度方案：如每早晨9:00执行，周一、周三、周五下午5:00执行等.

* 

Calendar:日历，它表示一些特定时间点的集合.一个Trigger可以和多个Calendar关联，以便排除或包含某些时间点。假设，我们安排每周星期一早上10:00执行任务，但是如果碰到法定的节日，任务则不执行，这时就需要在Trigger触发机制的基础上使用Calendar进行定点排除。

* 

Scheduler：调度器,代表一个Quartz的独立运行容器，Trigger和JobDetail可以注册到Scheduler中，两者在Scheduler中拥有各自的组及名称，组及名称是Scheduler查找定位容器中某一对象的依据，Trigger的组及名称必须唯一，JobDetail的组和名称也必须唯一（但可以和Trigger的组和名称相同，因为它们是不同类型的）。Scheduler定义了多个接口方法，允许外部通过组及名称访问和控制容器中Trigger和JobDetail。

` Scheduler可以将Trigger绑定到某一JobDetail中，这样当Trigger触发时，对应的Job就被执行。一个Job可以对应多个Trigger，但一个Trigger只能对应一个Job。 复制代码`

### 三. quarz 常用配置项介绍 ###

quartz 默认的配置文件是quartz.properties ,里面有很多配置项,比如调度器(schedule)相关配置,事务的配置，监听器，插件的配置，rmi 相关配置，集群相关配置等等。 详细配置信息可以参照官方文档: [www.quartz-scheduler.org/documentati…]( https://link.juejin.im?target=http%3A%2F%2Fwww.quartz-scheduler.org%2Fdocumentation%2F2.4.0-SNAPSHOT%2Fconfiguration.html%25E3%2580%2582 )

### 四. quratz 集群介绍 ###

quartz 集群是依赖于数据库机制实现的。部署多个节点的quartz 数据库配置必须保持一致。quartz集群是通过数据库表来感知其他的应用的，各个节点之间并没有直接的通信。只有使用持久的JobStore才能实现Quartz集群。

quartz 默认提供了11张数据库表:

* 

QRTZ_CALENDARS: 存储Quartz的Calendar信息

* 

QRTZ_CRON_TRIGGERS:存储CronTrigger，包括Cron表达式和时区信息

* 

QRTZ_FIRED_TRIGGERS:存储与已触发的Trigger相关的状态信息，以及相联Job的执行信息

* 

QRTZ_PAUSED_TRIGGER_GRPS:存储已暂停的Trigger组的信息

* 

QRTZ_SCHEDULER_STATE：存储少量的有关Scheduler的状态信息，和别的Scheduler实例

* 

QRTZ_LOCKS：存储程序的悲观锁的信息

* 

QRTZ_JOB_DETAILS:存储每一个已配置的Job的详细信息

* 

QRTZ_SIMPLE_TRIGGERS：存储简单的Trigger，包括重复次数、间隔、以及已触的次数

* 

QRTZ_BLOG_TRIGGERS Trigger作为Blob类型存储

* 

QRTZ_TRIGGERS 存储已配置的Trigger的信息

* 

QRTZ_SIMPROP_TRIGGERS:存储简单触发器相关。

` 要注意的是quartz集群间的相互感知是通过数据库表（QRTZ_SCHEDULER_STATE）实现的。对于水平集群，存在着时间同步问题。节点用时间戳来通知其他实例它自己的最后检入时间。假如节点的时钟被设置为将来的时间，那么运行中的Scheduler将再也意识不到那个结点已经宕掉了。另一方面，如果某个节点的时钟被设置为过去的时间，也许另一节点就会认定那个节点已宕掉并试图接过它的Job重运行。 复制代码`

# 小结 #

本文只是对quartz 分布式定时任务做了一个简单的介绍。quartz 作为一个成熟的分布式定时任务框架，然而好像并没有提供友好的运维管理界面，幸运的是很多的开源框架弥补了这个不足，比如elastic-job ，xxl-job等等，关于quartz 的更多内容，在后面不断的实践中再补充。