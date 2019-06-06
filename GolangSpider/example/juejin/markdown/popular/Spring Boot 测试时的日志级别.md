# Spring Boot 测试时的日志级别 #

## 1.概览 ##

该教程中，我将向你展示：如何在测试时设置spring boot 日志级别。虽然我们可以在测试通过时忽略日志，但是如果需要诊断失败的测试，选择正确的日志级别是非常重要的。

## 2.日志级别的重要性 ##

正确设置日志级别可以节省我们许多时间。 举例来说，如果测试在CI服务器上失败，但在开发服务器上时却通过了。我们将无法诊断失败的测试，除非有足够的日志输出。 为了获取正确数量的详细信息，我们可以微调应用程序的日志级别，如果发现某个java包对我们的测试更加重要，可以给它一个更低的日志级别，比如DEBUG。类似地，为了避免日志中有太多干扰，我们可以为那些不太重要的包配置更高级别的日志级别，例如INFO或者ERROR。 一起来探索设置日志级别的各种方法吧！

## 3. application.properties中的日志设置 ##

如果想要修改测试中的日志级别，我们可以在 ` src/test/resources/application.properties` 设置属性：

` logging.level.com.baeldung.testloglevel=DEBUG 复制代码`

该属性将会为指定的包 ` com.baeldung.testloglevel` 设置日志级别。 同样地，我们可以通过设置 ` root` 日志等级，更改所有包的日志级别

` logging.level.root=INFO 复制代码`

现在通过添加REST端点写入日志，来尝试下日志设置。

` @RestController public class TestLogLevelController { private static final Logger LOG = LoggerFactory.getLogger(TestLogLevelController.class); @Autowired private OtherComponent otherComponent; @GetMapping ( "/testLogLevel" ) public String testLogLevel () { LOG.trace( "This is a TRACE log" ); LOG.debug( "This is a DEBUG log" ); LOG.info( "This is an INFO log" ); LOG.error( "This is an ERROR log" ); otherComponent.processData(); return "Added some log output to console..." ; } } 复制代码`

正如所料，如果我们在测试中调用这个端点，我们将可以看到来自 ` TestLogLevelController` 的调试日志。

` 2019-04-01 14:08:27.545 DEBUG 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController : This is a DEBUG log 2019-04-01 14:08:27.545 INFO 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController : This is an INFO log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController : This is an ERROR log 2019-04-01 14:08:27.546 INFO 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent : This is an INFO log from another package 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent : This is an ERROR log from another package 复制代码`

这样设置日志级别十分简单，如果测试用 ` @SpringBootTest` 注解，那么我们肯定应该这样做。但是，如果不使用该注解，则必须以另一种方式配置日志级别。

### 3.1 基于Profile的日志设置 ###

尽管将配置放在 ` src\test\application.properties` 在大多数场景下好用，但在某些情况下，我们可能希望为一个或一组测试设置不同的配置。 在这种情况下，我们可以使用 ` @ActiveProfiles` 注解向测试添加一个 ` Spring Profile` :

` @RunWith (SpringRunner.class) @SpringBootTest (webEnvironment = WebEnvironment.RANDOM_PORT, classes = TestLogLevelApplication.class) @EnableAutoConfiguration (exclude = SecurityAutoConfiguration.class) @ActiveProfiles ( "logging-test" ) public class TestLogLevelWithProfileIntegrationTest { // ... } 复制代码`

日志设置将会存在 ` src/test/resources` 目录下的 ` application-logging-test.properties` 中：

` logging.level.com.baeldung.testloglevel=TRACE logging.level.root=ERROR 复制代码`

如果使用描述的设置调用 ` TestLogLevelCcontroller` ，将看到controller中打印的 ` TRACE` 级别日志，并且不会看到其他包出现INFO级别以上的日志。

` 2019-04-01 14:08:27.545 DEBUG 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController : This is a DEBUG log 2019-04-01 14:08:27.545 INFO 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController : This is an INFO log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController : This is an ERROR log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent : This is an ERROR log from another package 复制代码`

# 4.配置Logback #

如果使用Spring Boot默认的 ` Logback` ,可以在 ` src/test/resources` 目录下的 ` logback-text.xml` 文件中设置日志级别：

` < configuration > < include resource = "/org/springframework/boot/logging/logback/base.xml" /> < appender name = "STDOUT" class = "ch.qos.logback.core.ConsoleAppender" > < encoder > < pattern > %d{HH:mm:ss.SSS} [%thread] %-5level %logger{36} - %msg%n </ pattern > </ encoder > </ appender > < root level = "error" > < appender-ref ref = "STDOUT" /> </ root > < logger name = "com.baeldung.testloglevel" level = "debug" /> </ configuration > 复制代码`

以上例子如何在测试中为Logback配置日志级别。 root日志级别设置为INFO， ` com.baeldung.testloglevel` 包的日志级别设置为DEBUG。 再来一次，看看提交以上配置后的日志输出情况

` 2019-04-01 14:08:27.545 DEBUG 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is a DEBUG log 2019-04-01 14:08:27.545  INFO 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an INFO log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an ERROR log 2019-04-01 14:08:27.546  INFO 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an INFO log from another package 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an ERROR log from another package 复制代码`

### 4.1 基于Profile配置Logback ###

另一种配置指定Profile文件的方式就是在 ` application.properties` 文件中设置 ` logging.config` 属性：

` logging.config=classpath:logback-testloglevel.xml 复制代码`

或者，如果想在classpath只有一个的Logback配置，可以在 ` logbacl.xml` 使用 ` springProfile` 属性。

` < configuration > < include resource = "/org/springframework/boot/logging/logback/base.xml" /> < appender name = "STDOUT" class = "ch.qos.logback.core.ConsoleAppender" > < encoder > < pattern > %d{HH:mm:ss.SSS} [%thread] %-5level %logger{36} - %msg%n </ pattern > </ encoder > </ appender > < root level = "error" > < appender-ref ref = "STDOUT" /> </ root > < springProfile name = "logback-test1" > < logger name = "com.baeldung.testloglevel" level = "info" /> </ springProfile > < springProfile name = "logback-test2" > < logger name = "com.baeldung.testloglevel" level = "trace" /> </ springProfile > </ configuration > 复制代码`

现在使用 ` logback-test1` 配置文件调用 ` TestLogLevelController` ，将会获得如下输出：

` 2019-04-01 14:08:27.545  INFO 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an INFO log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an ERROR log 2019-04-01 14:08:27.546  INFO 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an INFO log from another package 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an ERROR log from another package 复制代码`

另一方面，如果更改配置为 ` logback-test2` ，输出将变成如下：

` 2019-04-01 14:08:27.545 DEBUG 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is a DEBUG log 2019-04-01 14:08:27.545  INFO 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an INFO log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an ERROR log 2019-04-01 14:08:27.546  INFO 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an INFO log from another package 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an ERROR log from another package 复制代码`

## 5.可选的Log4J ##

另外，如果我们使用 ` Log4J2` ，我们可以在 ` src\main\resources` 目录下的 ` log4j2-spring.xml` 文件中配置日志等级。

` < Configuration > < Appenders > < Console name = "Console" target = "SYSTEM_OUT" > < PatternLayout pattern = "%d{HH:mm:ss.SSS} [%thread] %-5level %logger{36} - %msg%n" /> </ Console > </ Appenders > < Loggers > < Logger name = "com.baeldung.testloglevel" level = "debug" /> < Root level = "info" > < AppenderRef ref = "Console" /> </ Root > </ Loggers > </ Configuration > 复制代码`

我们可以通过 ` application.properties` 中的 ` logging.config` 属性来设置Log4J 配置的路径。

` logging.config=classpath:log4j-testloglevel.xml 复制代码`

最后，查看使用以上配置后的输出：

` 2019-04-01 14:08:27.545 DEBUG 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is a DEBUG log 2019-04-01 14:08:27.545  INFO 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an INFO log 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.testloglevel.TestLogLevelController  : This is an ERROR log 2019-04-01 14:08:27.546  INFO 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an INFO log from another package 2019-04-01 14:08:27.546 ERROR 56585 --- [nio-8080-exec-1] c.b.component.OtherComponent  : This is an ERROR log from another package 复制代码`

## 6.结论 ##

在本文中，我们学习了如何在Spring Boot测试应用程序时设置日志级别，并探索了许多不同的配置方法。在 ` Spring Boot` 应用程序中使用 ` application.properties` 设置日志级别是最简便的，尤其是当我们使用@SpringBootTest注解时。 与往常一样，这些示例的源代码都在 [GitHub]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feugenp%2Ftutorials%2Ftree%2Fmaster%2Fspring-boot-testing ) 上。

> 
> 
> 
> 原文链接： [www.baeldung.com/spring-boot…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.baeldung.com%2Fspring-boot-testing-log-level
> )
> 
> 

> 
> 
> 
> 作者：baeldung
> 
> 

> 
> 
> 
> 译者：Leesen
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bb921b2daafc?imageView2/0/w/1280/h/960/ignore-error/1)