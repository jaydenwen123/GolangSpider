# 面试那点小事，你从未见过的spring boot面试集锦（附详细答案） #

#### 一， 什么是spring boot？ ####

多年来，随着新功能的增加，spring变得越来越复杂。只需访问页面https://spring.io/projects，我们将看到所有在应用程序中使用的不同功能的spring项目。如果必须启动一个新的spring项目，我们必须添加构建路径或maven依赖项，配置application server，添加spring配置。因此，启动一个新的spring项目需要大量的工作，因为我们目前必须从头开始做所有事情。Spring Boot是这个问题的解决方案。Spring boot构建在现有Spring框架之上。使用spring boot，我们可以避免以前必须执行的所有样板代码和配置。因此，Spring boot帮助我们更健壮地使用现有的Spring功能，并且只需最少的工作量。

#### 二，Spring Boot的优点是什么? ####

* 

减少开发、测试的时间和工作量。

* 

使用JavaConfig有助于避免使用XML。

* 

避免大量maven导入和各种版本冲突。

* 

提供可选的开发方法。

* 

通过提供默认开发方式进行快速开发。

* 

不需要单独的Web服务器。这意味着您不再需要启动Tomcat、Glassfish或其他任何东西。

由于没有web.xml文件，所以需要更少的配置。只需添加带@ configuration注释的类，然后可以添加带@ bean注释的方法，Spring将自动加载对象并像往常一样管理它。您甚至可以将@Autowired添加到bean方法中，使Spring autowire成为bean所需的依赖项。

基于环境的配置——使用这些属性，您可以将其传递到您正在使用的应用程序环境中:- dspring .profile .active={enviorement}。在加载主应用程序属性文件之后，Spring将在(application-{environment}.properties)处加载后续的应用程序属性文件。

#### 三，您使用过哪些构建工具来开发Spring引导应用程序? ####

Spring Boot应用程序可以使用Maven和Gradle开发。

#### 四，什么是JavaConfig? ####

Spring JavaConfig是Spring社区的一个产品，它提供了一种纯java方法来配置Spring IoC容器。因此，它有助于避免使用XML配置。使用JavaConfig的优点是：

> 
> 
> 
> * 
> 
> 面向对象的配置。因为配置在JavaConfig中定义为类，所以用户可以充分利用Java中的面向对象特性。一个配置类可以子类化另一个配置类，覆盖它的@Bean方法，等等。
> 
> 
> 
> * 
> 
> 减少或消除XML配置。已经证明了基于依赖注入原则的外部化配置的好处。然而，许多开发人员不愿意在XML和Java之间来回切换。JavaConfig为开发人员提供了一种纯java方法来配置Spring容器，这种方法在概念上类似于XML配置。从技术上讲，仅使用
> JavaConfig配置类来配置容器是可行的，但是在实践中，许多人发现将JavaConfig与XML混合并匹配是理想的。
> 
> 
> * 
> 
> 类型安全的重构能力。JavaConfig提供了一种类型安全的配置Spring容器的方法。由于Java
> 5.0对泛型的支持，现在可以通过类型而不是名称检索bean，不需要进行任何基于类型转换或字符串的查找。
> 
> 
> 
> 
> 

#### 五，如何在不重启服务器的情况下在Spring引导时重新加载我的更改? ####

这可以通过开发工具来实现。有了这个依赖项，您保存的任何更改都将重新启动嵌入的tomcat。Spring Boot有一个开发人员工具(DevTools)模块，它有助于提高开发人员的工作效率。Java开发人员面临的关键挑战之一是将文件更改自动部署到服务器并自动重启服务器。开发人员可以在Spring引导时重新加载更改，而不必重新启动服务器。这将消除每次手动部署更改的需要。Spring Boot在发布第一个版本时没有这个特性。这是开发人员最需要的特性。DevTools模块完全满足开发人员的需求。此模块将在生产环境中禁用。它还提供了H2-database控制台，以便更好地测试应用程序。使用以下依赖项

` <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-devtools</artifactId> <optional> true </optional></dependency> 复制代码`

下面的 [示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2FSpringBootUsingPagination ) 演示了自动启动和H2 DB控制台的DevTool依赖项用法

#### 六，什么是Spring boot actuator? ####

Spring boot actuator是Spring boot framework的重要特性之一。Spring boot actuator帮助您访问生产环境中正在运行的应用程序的当前状态，在生产环境中必须检查和监视几个指标。甚至一些外部应用程序也可能使用这些服务来触发对相关人员的警报消息。actuator模块公开一组REST端点，这些端点可以作为HTTP URL直接访问，以检查状态。

#### 七，如何将Spring Boot应用程序作为war包部署? ####

Spring Boot WAR部署

#### 八，什么是Docker吗?如何将Spring引导应用程序部署到Docker? ####

Docker是什么 将基于Spring的WAR应用程序部署到Docker 将基于Spring的JAR应用程序部署到Docker

#### 九，如何禁用执行器端点安全在Spring启动? ####

默认情况下，所有敏感的HTTP端点都是安全的，只有具有ACTUATOR角色的用户才能访问它们。安全性是使用标准HttpServletRequest.isUserInRole方法实现的。 我们可以使用-禁用安全性 management.security.enabled = false 建议仅当在防火墙后访问ACTUATOR端点时禁用安全性。

#### 十，如何将Spring引导应用程序运行到自定义端口? ####

要在自定义端口上运行spring引导应用程序，可以在application.properties中指定端口。 server.port = 8090

#### 十一，什么是ELK堆栈?如何与Spring Boot一起使用? ####

ELK堆栈由三个开源产品组成——Elasticsearch、Logstash和Kibana from Elastic。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d5b5d3af6a1?imageView2/0/w/1280/h/960/ignore-error/1) Elasticsearch是一个基于Lucene搜索引擎的NoSQL数据库。     Logstash是一个日志管道工具，它接受来自不同来源的输入，执行不同的转换，并将数据导出到不同的目标。它是一个动态的数据收集管道，具有可扩展的插件生态系统和强大的弹性搜索协同作用     Kibana是一个可视化UI层，工作在Elasticsearch之上。 这三个项目一起用于各种环境中的日志分析。因此Logstash收集和解析日志、弹性搜索索引并存储这些信息，而Kibana提供了一个UI层，提供可操作的可见性。

[Spring Boot + ELK stack]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspringboot-microservice-elk )

#### 十二，您有使用Spring Boot编写测试用例吗? ####

SpringBoot为编写单元测试用例提供了@SpringBootTest

[Spring引导单元测试的简单示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspringboot_testcases )

#### 十三，YAML是什么? ####

YAML是一种人类可读的数据序列化语言。它通常用于配置文件。 与属性文件相比，YAML文件的结构更加结构化，如果我们希望在配置文件中添加复杂的属性，那么它不会造成太大的混乱。可以看到，YAML具有分层的配置数据。

[在Spring引导中使用YAML属性]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fbootyaml )

#### 十四，如何为Spring引导应用程序实现安全性? ####

为了实现Spring Boot的安全性，我们使用Spring - Boot -starter-security依赖项，必须添加安全配置。它只需要很少的代码。Config类必须扩展WebSecurityConfigurerAdapter并覆盖它的方法。

[Spring引导安全性示例和说明]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fsprboot_sec )

#### 十五，您是否集成了Spring Boot和ActiveMQ ? ####

为了集成Spring Boot和ActiveMQ，我们使用Spring - Boot -starter- ActiveMQ依赖项，它只需要很少的配置，没有样板代码。

[Spring引导ActiveMQ说明]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fsprboot_activemq )

#### 十六，您是否集成了Spring Boot和Apache Kafka ? ####

为了集成Spring Boot和Apache Kafka，我们使用Spring - Kafka依赖项。

[Spring Boot + Apache Kafka示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-apache-kafka-hello-world )

#### 十七，如何使用Spring引导实现分页和排序? ####

使用Spring Boot实现分页非常简单。使用Spring Data-JPA，这是通过传递可分页的org.springframe .data.domain来实现的。可分页到存储库方法。

[Spring引导分页说明]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2FSpringBootInterviewQuestions )

#### 十八，什么是Swagger?您是否使用Spring Boot实现了它? ####

Swagger被广泛用于可视化api，Swagger UI为前端开发人员提供在线沙箱环境。在本教程中，我们将使用Swagger 2规范的Springfox实现。Swagger是一种工具、规范和完整的框架实现，用于生成RESTful Web服务的可视化表示。它允许文档以与服务器相同的速度更新。当通过Swagger正确定义时，使用者可以用最少的实现逻辑理解远程服务并与之交互。因此Swagger消除了调用服务时的猜测。

[Spring Boot + Swagger2]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot_swagger )

#### 十九，什么是Spring Profiles?如何使用Spring Boot实现它? ####

Spring Profiles允许用户根据配置文件(dev, test, prod等)注册bean。因此，当应用程序在开发中运行时，只能加载某些bean，当应用程序在生产中运行时，只能加载某些其他bean。假设我们的需求是Swagger文档只对QA环境启用，对所有其他环境禁用。这可以使用配置文件来完成。Spring Boot使得使用配置文件非常容易。

[Spring引导+配置文件]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot_swagger_profile )

#### 二十，什么是Spring Boot Batch?如何使用Spring Boot实现它? ####

Spring Boot Batch提供了处理大量记录所必需的可重用功能，包括日志/跟踪、事务管理、作业处理统计信息、作业重启、作业跳过和资源管理。它还提供了更高级的技术服务和特性，通过优化和分区技术，这些特性将支持极高容量和高性能的批处理作业。无论是简单的还是复杂的，大容量批处理作业都可以以高度可伸缩的方式利用该框架来处理大量信息。

[Spring Boot Batch]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fbootbatch )

#### 二十一，什么是FreeMarker模板?如何使用Spring Boot实现它? ####

答:FreeMarker是一个基于java的模板引擎，最初专注于使用MVC软件架构生成动态web页面。使用Freemarker的主要优势是完全分离了表示层和业务层。程序员可以处理应用程序代码，而设计人员可以处理html页面设计。最后，使用freemarker，这些可以组合在一起，给出最终的输出页面。

[Spring Boot + FreeMarker的例子]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-freemarker-hello-world )

#### 二十二，如何使用Spring Boot实现异常处理? ####

Spring提供了一种非常有用的方法，可以使用ControllerAdvice处理异常。我们将实现一个ControlerAdvice类，它将处理控制器类抛出的所有异常。

[Spring引导异常处理]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot-exception-handling )

#### 二十三，什么是缓存?您在Spring引导中使用过缓存框架吗? ####

缓存是本地内存的一个区域，它保存了频繁访问的数据的副本，否则获取或计算这些数据将非常昂贵。使用Hazelcast进行缓存。

[Spring Boot + Hazelcast示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-hazelcast )

#### 二十四您是否使用Spring Boot公开了SOAP web服务端点? ####

是的。使用Spring Boot公开了要使用的web服务。使用契约优先的方法从wsdl生成类。

[Spring引导+ SOAP Web服务示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspringbootsoapwebservice )

#### 二十五，您如何使用Spring Boot执行数据库操作? ####

[Spring引导教程-Spring Data JPA]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2FSpringBoot_DataJPA )

[Spring引导JDBC示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fbootjdbc )

#### 二十六，如何使用Spring上传文件? ####

[Spring Boot +文件上传的例子]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fbootupload )

#### 二十七，如何用Spring Boot实现拦截器? ####

[使用Spring MVC HandlerInterceptor与Spring引导]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2FbootInterceptor )

#### 二十八，如何在Spring Boot下使用schedulers ? ####

[Spring引导任务调度程序示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2FbootTask )

#### 二十九，您使用过哪些启动器maven依赖项? ####

使用过不同的starter依赖项，如spring-boot-starter-activemq依赖项、spring-boot-starter-security依赖项、spring-boot-starter-web依赖项。 这有助于减少依赖项的数量，并减少版本组合。

[Spring引导安全性示例和说明]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fsprboot_sec )

#### 三十，什么是CSRF攻击?如何启用CSRF对其进行保护? ####

CSRF代表跨站请求伪造。它是一种攻击，迫使最终用户在其当前已经过身份验证的web应用程序上执行不需要的操作。CSRF攻击专门针对状态更改请求，而不是数据窃取，因为攻击者无法看到对伪造请求的响应。

[Spring引导安全性—启用CSRF保护]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot_security_csrf )

#### 三十一，如何使用Spring引导使用表单登录身份验证? ####

[Spring引导表单安全登录Hello World示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot_form_security )

#### 三十二，什么是OAuth2?如何使用Spring Boot实现它? ####

[Spring Boot + OAuth2实现]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-oauth-introduction )

#### 三十三，GZIP是什么?如何使用Spring Boot实现它? ####

gzip是一种文件格式，是一种用于文件压缩和解压缩的软件应用程序。

[Spring引导+ GZIP压缩]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot-zip )

#### 三十四，您在Spring引导中使用过集成框架吗? ####

已将Apache Camel与Spring引导集成。使用Apache Camel Spring启动启动依赖项。

[Spring Boot +Apache Camel]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fbootcamel )

#### 三十五，什么是Apache Freemarker?什么时候使用它而不是JSP?如何与Spring Boot集成? ####

JSP是为网页量身定做的，Freemarker模板是一种更通用的模板语言——它可以用来生成html、纯文本、电子邮件等。

[Spring Boot + FreeMarker的例子]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-freemarker-hello-world )

#### 三十六，你什么时候使用WebSockets?如何使用Spring Boot实现它? ####

WebSocket是一种计算机通信协议，通过单个TCP连接提供全双工通信通道。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d5b5d292438?imageView2/0/w/1280/h/960/ignore-error/1) WebSocket是双向的——使用WebSocket客户端或服务器都可以发起发送消息。 WebSocket是全双工的——客户端和服务器之间的通信是相互独立的。 单个TCP连接——初始连接使用HTTP，然后将此连接升级为基于套接字的连接。然后，这个单一连接将用于未来的所有通信 轻- WebSocket消息数据交换比http轻得多。 [Spring Boot + WebSockets的例子]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot-websocket )

#### 三十七，什么是AOP?如何与Spring Boot一起使用? ####

在软件开发过程中，跨越应用程序多个点的功能称为横切关注点。这些横切关注点不同于应用程序的主要业务逻辑。因此，将这些横切关注点从业务逻辑中分离出来是面向方面编程(AOP)的切入点。

[Spring Boot + AOP示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-aop )

#### 三十八，什么是Apache Kafka?如何与Spring Boot集成? ####

apache Kafka是一个分布式发布-订阅消息传递系统。它是一个可伸缩的、容错的、发布-订阅消息传递系统，使我们能够构建分布式应用程序。这是一个Apache顶级项目。Kafka适用于离线和在线的消息消费。

[Spring Boot + Apache Kafka示例]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-boot-apache-kafka-hello-world )

#### 三十九，我们如何监视所有Spring Boot微服务? ####

Spring Boot提供了actuator 端点来监控单个微服务的指标。这些端点对于获取关于应用程序的信息非常有帮助，比如应用程序是否启动，它们的组件(如数据库等)是否正常工作。但是，使用actuator 接口的一个主要缺点或困难是，我们必须逐个命中这些接口，以了解应用程序的状态或健康状况。假设微服务涉及50个应用程序，管理员将不得不命中所有50个应用程序的actuator 端点。为了帮助我们处理这种情况，我们将使用位于https://github.com/codecentric/springing-boot-admin的开源项目。 它构建在Spring Boot Actuator之上，提供了一个web UI，使我们能够可视化多个应用程序的指标。

[Spring Boot Admin]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fboot-admin )

#### 四十，您在Spring引导中使用过Spring Cloud组件吗? ####

使用过Netflix Eureka等Spring Cloud组件进行服务注册，Ribbon用于负载平衡。 [Spring Boot + Cloud Components ]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspringcloud ) [Spring Cloud interview Questions]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fspring%2Fspring-cloud-interview-questions )

#### 四十一，如何将Spring Boot应用程序部署到Pivotal Cloud Foundry(PCF)? ####

[Deploying Spring Boot Application to PCF]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fpcf%2Fpcf-hello )

#### 四十二，如何将Spring Boot + MySQL应用部署到Pivotal Cloud Foundry(PCF)? ####

[Pivotal Cloud Foundry Tutorial - Deploying Spring Boot + MySQL Application to PCF]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fpcf%2Fpcf-sql )

#### 四十三，如何将Spring Boot + RabbitMQ应用部署到Pivotal Cloud Foundry(PCF)? ####

[Pivotal Cloud Foundry Tutorial - Deploying Spring Boot + RabbitMQ Application to PCF]( https://link.juejin.im?target=https%3A%2F%2Fwww.javainuse.com%2Fpcf%2Fpcf-rabbitmq )