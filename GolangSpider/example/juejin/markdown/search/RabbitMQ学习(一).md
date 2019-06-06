# RabbitMQ学习(一) #

# 1.RabbitMQ简介 #

## 1.1 什么是RabbitMQ ##

* 

RabbitMQ 是一个 **消息中间件** , 一个由 **Erlang 语言** 开发的 **AMQP** 的开源实现。

* 

AMQP : Advanced Message Queue Protocol，高级消息队列协议。它是应用层协议的一个开放标准，为面向消息的中间件设计，基于此协议的客户端与消息中间件可传递消息，并不受产品、开发语言等条件的限制。

## 1.2 RabbitMQ的特点 ##

RabbitMQ 最初起源于金融系统，用于在分布式系统中存储转发消息，在易用性、扩展性、高可用性等方面表现不俗。具体特点如下：

* **可靠性（Reliability） :** RabbitMQ 使用一些机制来保证可靠性，如持久化、传输确认、发布确认。
* **灵活的路由（Flexible Routing） :** 在消息进入队列之前，通过 Exchange 来路由消息的。对于典型的路由功能，RabbitMQ 已经提供了一些内置的 Exchange 来实现。针对更复杂的路由功能，可以将多个 Exchange 绑定在一起，也通过插件机制实现自己的 Exchange 。
* **消息集群（Clustering） :** 多个 RabbitMQ 服务器可以组成一个集群，形成一个逻辑 Broker 。
* **高可用（Highly Available Queues） :** 队列可以在集群中的机器上进行镜像，使得在部分节点出问题的情况下队列仍然可用。
* **多种协议（Multi-protocol） :** RabbitMQ 支持多种消息队列协议，比如 STOMP、MQTT 等等。
* **多语言客户端（Many Clients） :** RabbitMQ 几乎支持所有常用语言，比如 Java、.NET、Ruby 等等。
* **管理界面（Management UI）:** RabbitMQ 提供了一个易用的用户界面，使得用户可以监控和管理消息 Broker 的许多方面。
* **跟踪机制（Tracing）:** 如果消息异常，RabbitMQ 提供了消息跟踪机制，使用者可以找出发生了什么。
* **插件机制（Plugin System）:** RabbitMQ 提供了许多插件，来从多方面进行扩展，也可以编写自己的插件

## 1.3 架构图与主要概念 ##

### 1.3.1 架构图 ###

![clipboard (1).png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a921251c96?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.3.2 主要概念 ###

* **RabbitMQ Server ：** 也叫 broker server ，它是一种传输服务。 他的角色就是维护一条从Producer 到 Consumer 的路线，保证数据能够按照指定的方式进行传输。
* **Producer ：** 消息生产者，如架构图中的 A、B、C，数据的发送方。消息生产者连接 RabbitMQ 服务器然后将消息投递到 Exchange。
* **Consumer ：**消息消费者，如架构图图 1、2、3，数据的接收方。消息消费者订阅队列，RabbitMQ 将 Queue 中的消息发送到消息消费者。
* **Exchange(交换器) ：**生产者将消息发送到 Exchange（交换器），由 Exchange 将消息路由到一个或多个 Queue 中（或者丢弃）。Exchange 并不存储消息。RabbitMQ 中的 Exchange 有 direct、fanout、topic、headers 四种类型，每种类型对应不同的路由规则。
* **Queue(队列) :** 是 RabbitMQ 的内部对象，用于存储消息。消息消费者就是通过订阅队列来获取消息的，RabbitMQ 中的消息都只能存储在 Queue 中，生产者生产消息并最终投递到Queue 中，消费者可以从 Queue 中获取消息并消费。 **多个消费者可以订阅同一个 Queue，这时 Queue 中的消息会被平均分摊给多个消费者进行处理，而不是每个消费者都收到所有的消息并处理。**
* **RoutingKey(路由规则)：**生产者在将消息发送给 Exchange 的时候，一般会指定一个 routing key，来指定这个消息的路由规则，而这个 routing key 需要与 Exchange Type 及binding key 联合使用才能最终生效。在 Exchange Type 与 binding key 固定的情况下（在正常使用时一般这些内容都是固定配置好的），我们的生产者就可以在发送消息给 Exchange 时，通过指定 routing key 来决定消息流向哪里。RabbitMQ 为 routing key 设定的长度限制为 255bytes。
* **Connection(连接) ：**Producer 和 Consumer 都是通过 **TCP** 连接到 RabbitMQ Server的。以后我们可以看到，程序的起始处就是建立这个 TCP 连接。
* **Channels(信道) :** 它建立在上述的 TCP 连接中。数据流动都是在 Channel 中进行的。也就是说，一般情况是程序起始建立 TCP 连接，第二步就是建立这个 Channel。
* **VirtualHost ：**权限控制的基本单位，一个 VirtualHost 里面有若干 Exchange 和MessageQueue，以及指定被哪些 user 使用

# 2. RabbitMQ的安装 #

## 2.1 Windows下RabbitMQ的安装 ##

**(1) 下载并安装Erlang**

下载 **otp_win64_20.2.exe** , 并以管理员的身份运行安装 ;

**(2) 下载并安装rabbitMQ Server**

下载地址 : [www.rabbitmq.com/install-win…]( https://link.juejin.im?target=http%3A%2F%2Fwww.rabbitmq.com%2Finstall-windows-manual.html )

下载完成为 : **rabbitmq-server-3.7.4.exe**

**注意 :** 安装路径不能存在中文与空格 , 安装完整后Windows服务中就已经是存在rabbitMQ了 , 而且是启动状态

**(3) 安装管理界面**

进入 rabbitMQ 安装目录的 sbin 目录，输入以下命令

` rabbitmq-plugins enable rabbitmq_management 复制代码`

![clipboard.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9250b5627?imageView2/0/w/1280/h/960/ignore-error/1)

**(4) 重启rabbitmq 服务**

![clipboard.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a924f6e685?imageView2/0/w/1280/h/960/ignore-error/1)

**(5) 测试**

打开浏览器地址栏输入http://127.0.0.1:15672(默认端口为15671) ,即可看到管理界面的登陆页如下 , 用户名和密码均为 ` guest` , 进入主界面

![Snipaste_2019-05-08_21-11-44.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a925021b26?imageView2/0/w/1280/h/960/ignore-error/1)

![clipboard.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a928e5c264?imageView2/0/w/1280/h/960/ignore-error/1)

## 2.2 docker下安装rabbitMQ ##

**(1) 镜像下载 :**

` docker pull rabbitmq:management 复制代码`

**(2) 创建容器 :**

rabbitmq 需要有映射以下端口: 5671 5672 4369 15671 15672 25672 , 启动命令中以多个 -p 映射端口号 , 其中各个端口的作用如下 :

* 15672 (if management plugin is enabled)
* 15671 management 监听端口
* 5672, 5671 (AMQP 0-9-1 without and with TLS)
* 4369 (epmd) epmd 代表 Erlang 端口映射守护进程
* 25672 (Erlang distribution)

` docker run -di --name=rabbitmq_demo -p 5671:5617 -p 5672:5672 -p 4369:4369 -p 15671:15671 -p 15672:15672 -p 25672:25672 rabbitmq:management 复制代码`

**(3) 测试 :**

在浏览器中访问如下url(ip地址根据实际情况更换 , 这里我安装在个人虚拟上) :

[http://192.168.66.133:15672]( https://link.juejin.im?target=http%3A%2F%2F192.168.66.133%3A15672 )

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a927f80bf9?imageView2/0/w/1280/h/960/ignore-error/1)

# 3. RabbitMQ的Exchange模式详述 #

以下RabbitMQ代码实例均使用Spring对Rabbit进行整合 , 执行相关操作

## 3.1 直接模式（Direct） ##

### 3.1.1 使用场景 ###

我们需要 **将消息发给唯一一个节点(队列** )时使用这种模式，这是最简单的一种形式。

### **3.1.2 直接模式的流程图** ###

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a952487544?imageView2/0/w/1280/h/960/ignore-error/1)

* 任何发送到 Direct Exchange 的消息都会被转发到 RouteKey 中指定的 Queue。如上所示 , routing key 指定的队列为"KEY" ;
* 一般情况可以使用 rabbitMQ **自带的 Exchange** ： ` ""` (该 Exchange 的名字为空字符串，下文称其为 **default Exchange** )。这种模式下不需要将 Exchange 进行任何绑定(binding)操作 ;
* 消息传递时需要一个Routingkey , 可以简单理解为要消息要发送到的目的队列的名字 ;
* 如果 vhost 中不存在 RouteKey 中指定的队列名，则该消息会被抛弃。

### 3.1.3 直接模式实例 ###

**(1) 创建队列**

创建一个叫 ` test_queue01` 的队列

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a953268136?imageView2/0/w/1280/h/960/ignore-error/1)

` Durability` ：是否做持久化 Durable（持久） transient（临时）

` Auto delete` : 是否自动删除

**(2) 环境准备**

创建maven工程(module) , 在pom.xml中引入AMQP起步依赖以及其他 :

` <? xml version= "1.0" encoding= "UTF-8" ?> < project xmlns = "http://maven.apache.org/POM/4.0.0" xmlns:xsi = "http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation = "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd" > < modelVersion > 4.0.0 </ modelVersion > < groupId > com.wk </ groupId > < artifactId > rabbitmq_demo </ artifactId > < version > 1.0-SNAPSHOT </ version > < parent > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-parent </ artifactId > < version > 2.0.1.RELEASE </ version > < relativePath /> </ parent > < properties > < project.build.sourceEncoding > UTF-8 </ project.build.sourceEncoding > < project.reporting.outputEncoding > UTF-8 </ project.reporting.outputEncoding > < java.version > 1.8 </ java.version > </ properties > < dependencies > <!--RabbitMQ--> < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-amqp </ artifactId > </ dependency > <!--Spring Boot 整合 Junit--> < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-test </ artifactId > < scope > test </ scope > </ dependency > </ dependencies > </ project > 复制代码`

编写spring boot的application.yml配置文件 :

` server: port: 9201 spring: rabbitmq: host: 192.168.66.133 username: guest # 默认值,可以不写 password: guest # 默认值,可以不写 复制代码`

编写SpringBoot启动类 :

` import org.springframework.boot.SpringApplication; import org.springframework.boot.autoconfigure.SpringBootApplication; @SpringBootApplication public class RabbitmqApplication { public static void main (String[] args) { SpringApplication.run(RabbitmqApplication.class, args); } } 复制代码`

**(3) 代码实现消息生产者与消费者**

编写消息 **生产者** 测试类 , 发送消息到上面已创建好的queue :

关键对象 : RabbitMessagingTemplate

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9540acdac?imageView2/0/w/1280/h/960/ignore-error/1)

编写消息 **消费者** 测试类 :

关键注解 : ` @RabbitListener` , ` @Component` , ` @RabbitHandler`

代码编写完毕后 , 运行工程 , 消费者自动生效

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a957859250?imageView2/0/w/1280/h/960/ignore-error/1)

工程启动完毕 , 消费者自动接收指定queue的消息并处理 , 可以看到控制台中显示如下内容:

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a95f32f2ac?imageView2/0/w/1280/h/960/ignore-error/1)

**(4) 多个消费者同时监听一个queue测试**

这里使用IDEA进行测试时会遇到一个问题, 就是IDEA默认是不允许启动两个相同的application的, 需要进行相关设置 , 如下 :

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a967bdfac5?imageView2/0/w/1280/h/960/ignore-error/1)

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a983170b7a?imageView2/0/w/1280/h/960/ignore-error/1)

**注意 :** 上述设置修改完毕后 , 若要再次启动同一个应用 , 需要在application.yml中 重新指定不同的端口号 , 默认是8080 , 上面已经启动了一个端口号为9201的程序 , 这里修改端口号为9202并再次启动 ;

` server: port: 9202 # <----修改端口号为9202 spring: rabbitmq: host: 192.168.66.133 username: guest # 默认值,可以不写,密码同 password: guest 复制代码`

再次启动时发现控制台运行着两个同样的程序 , 端口号不相同

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a982f835f7?imageView2/0/w/1280/h/960/ignore-error/1)

连续4次运行生产者程序 , 两个消息消费者接收消息的情况如下 , 表明多个消费者同时监听一个queue的时候 , 消息的分发情况是遵循 **轮询** 算法的

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a98ce8e0b2?imageView2/0/w/1280/h/960/ignore-error/1)

## 3.2 分列模式（Fanout） ##

### 3.2.1 使用场景 ###

当我们需要将消息一次发送给多个队列时 , 需要使用该模式

**3.2.2 分列模式流程图**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9995bc367?imageView2/0/w/1280/h/960/ignore-error/1)

* 任何发送到 Fanout Exchange 的消息都会被转发到与该 Exchange 绑定(Binding)的所有Queue 上。
* 该模式可以理解为路由表模式;
* 这种模式不需要routing key
* 这种模式需要提前将Exchange与Queue进行绑定 , 一个Exchange可以绑定多个Queue, 同时一个Queue可以与多个Exchange进行绑定;
* 如果接受到消息的 Exchange 没有与任何 Queue 绑定，则消息会被抛弃 。

### 3.2.3 分列模式实例 ###

**(1) 新建两个queue**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9adf26d14?imageView2/0/w/1280/h/960/ignore-error/1)

**(2) 新建一个exchange**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a999d6a303?imageView2/0/w/1280/h/960/ignore-error/1)

**(3) 将** **test_queue02 和 test_queue03 绑定到 test_fanout_exchange 上**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9bf3191aa?imageView2/0/w/1280/h/960/ignore-error/1)

绑定完成如下 :

![1557322434836.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9c7d19687?imageView2/0/w/1280/h/960/ignore-error/1)

**(4) 代码实现消息生产者**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9e1f17d38?imageView2/0/w/1280/h/960/ignore-error/1)

**(5) 代码实现消息消费者**

消费者1 监听 test_queue02

` @RabbitListener (queues = "test_queue02" ) @Component public class TestFanoutConsumer1 { @RabbitHandler public void handle (String msg) { System.out.println( "consumer1接收消息为:" + msg); } } 复制代码`

消费者2 监听 test_queue03

` @RabbitListener (queues = "test_queue03" ) @Component public class TestFanoutConsumer2 { @RabbitHandler public void handle (String msg) { System.out.println( "consumer2接收消息为:" + msg); } } 复制代码`

执行生产者代码 , 消息发送到exchange 中 , 然后启动消费者 , 消费者接收并处理消息内容 , 打印到控制台上 , 如下

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9f9635f9a?imageView2/0/w/1280/h/960/ignore-error/1)

## 3.3 主题模式(Topic) ##

### 3.3.1 使用场景 ###

需要通过exchange将消息转发到关心该消息routingkey 的queue 上时 , 可以使用主题模式

### 3.3.2 主题模式详解 ###

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9ce7866e4?imageView2/0/w/1280/h/960/ignore-error/1)

实质是通过exchange 实现 类似activeMQ的 **主题/订阅模式** , 指定exchange 将消息发送到 **其绑定的 ,** 且 **bindkey与生产者发送消息到exchange时指定的routingkey相匹配的** queue中 , 说白一点 , 就是消息的 **routingkey** 与 **bindkey** 进行 **模糊匹配**.

这里需要了解上图中的一些概念 :

* **routingkey :** 生产者发送消息时 , 指定该消息的路由规则 , 如下 :

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282a9fd40daaf?imageView2/0/w/1280/h/960/ignore-error/1)

* **bindkey :** exchange与queue进行绑定的时 , 指定bindkey , 如下 :

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aa0d7f176a?imageView2/0/w/1280/h/960/ignore-error/1)

* **通配符 :** 在绑定时指定的bindkey是可以使用通配符的 , 通配符有 ` #` 和 ` *` 两种 , 其中 ` #` **匹配一个或多个单词** , 每个单词以 `.` 分隔 , **而 ` *` 仅匹配一个单词** , 例如 : 当queue的bindkey为 ` usa.#` ,消息的routingke为 ` usa.new.weather` , 该消息是会转发到该queue中的 , 但如果queuebindkey为 ` usa.*` , 则该消息不会转发到queue中 ;

**主题模式的特点 :**

* 这种模式较为复杂，简单来说，就是每个队列都有其关心的主题，所有的消息都带有一个"标题"(RouteKey)，Exchange 会将消息转发到所有关注主题能与 RouteKey 模糊匹配的队列。
* 这种模式需要 RouteKey，也需要提前绑定 Exchange 与 Queue。
* 同样，如果 Exchange 没有发现能够与 RouteKey 匹配的 Queue，则会抛弃此消息

### 3.3.3 主题模式实例 ###

**(1) 新建exchange**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aa1f507b4b?imageView2/0/w/1280/h/960/ignore-error/1)

**(2) 新建queue**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aa4e799006?imageView2/0/w/1280/h/960/ignore-error/1)

**(3) 绑定exchange与queue , 并指定routingkey**

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aa7ffd4c82?imageView2/0/w/1280/h/960/ignore-error/1)

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aa6c90a734?imageView2/0/w/1280/h/960/ignore-error/1)

**(4) 代码实现消息消费者**

消费者1 :

` @RabbitListener (queues = "test_topic_queue01" ) @Component public class TestTopicConsumer01 { @RabbitHandler public void handle (String msg) { System.out.println( "TopicConsumer01 - routingkey:aaa.#" ); System.out.println( "接收消息为:" + msg); System.out.println( "------------------------------------" ); } } 复制代码`

消费者2 :

` @RabbitListener (queues = "test_topic_queue02" ) @Component public class TestTopicConsumer02 { @RabbitHandler public void handle (String msg) { System.out.println( "TopicConsumer02 - routingkey:aaa.*" ); System.out.println( "接收消息为:" + msg); System.out.println( "------------------------------------" ); } } 复制代码`

消费者3 :

` @RabbitListener (queues = "test_topic_queue03" ) @Component public class TestTopicConsumer03 { @RabbitHandler public void handle (String msg) { System.out.println( "TopicConsumer03 - routingkey:bbb.aaa" ); System.out.println( "接收消息为:" + msg); System.out.println( "------------------------------------" ); } } 复制代码`

启动程序监听queue ;

**(5) 代码实现消息生产者**

` /** * topic模式 */ @Test public void testTopicSend () { rabbitTemplate.convertAndSend( "test_topic_exchange" , "aaa.bbb" , "test topic 1111" ); //rabbitTemplate.convertAndSend("test_topic_exchange","aaa.bbb.ccc","test topic 2222"); //rabbitTemplate.convertAndSend("test_topic_exchange","bbb.aaa","test topic 3333"); } 复制代码`

依次执行6 7 8行代码 , 消费者接收消息打印到控制台的情况如下 :

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aae5f2c5d2?imageView2/0/w/1280/h/960/ignore-error/1)

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282aaec86e50c?imageView2/0/w/1280/h/960/ignore-error/1)

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b282af86abf7eb?imageView2/0/w/1280/h/960/ignore-error/1)

**这里需要注意两点 :**

* 多个消费者的操作是线程不安全的 , 上述第一个图出现了以下打印情况
` TopicConsumer02 - routingkey:aaa.* 接收消息为: test topic 1111 TopicConsumer01 - routingkey:aaa. # 接收消息为: test topic 1111 ------------------------------------ ------------------------------------ 复制代码` * 执行消费者代码时 , 如果同时取消注释 6 7 行代码 , 按道理应该是出现以下打印情况 :
` TopicConsumer01 - routingkey:aaa. # 接收消息为: test topic 1111 ------------------------------------ TopicConsumer01 - routingkey:aaa. # 接收消息为: test topic 2222 ------------------------------------ TopicConsumer02 - routingkey:aaa.* 接收消息为: test topic 1111 ------------------------------------ 复制代码`

但是实际打印结果是 :

` TopicConsumer01 - routingkey:aaa. # 接收消息为: test topic 1111 ------------------------------------ TopicConsumer02 - routingkey:aaa.* 接收消息为: test topic 1111 ------------------------------------ 复制代码`

这个问题尚待思考 ;

参考链接:

[blog.csdn.net/m0_37383637…]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fm0_37383637%2Farticle%2Fdetails%2F79264767 )

[blog.csdn.net/u013952133/…]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu013952133%2Farticle%2Fdetails%2F79435783 )

[blog.csdn.net/mingongge/a…]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fmingongge%2Farticle%2Fdetails%2F81132632 )