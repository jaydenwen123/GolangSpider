# 更优雅的获取springboot yml中的复杂数据 #

> 
> 
> 
> 偶然看到国外论坛有人在吐槽同事从配置文件获取值的方式太过冗长和臃肿，便有了这篇文章
> 
> 

github demo地址： [springboot-yml-value]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcaotinging%2Fsimple-demo%2Ftree%2Fmaster%2Fspringboot-yml-value )

### 1.什么是yml文件 ###

application.yml取代application.properties，用来配置数据可读性更强，尤其是当我们已经制定了很多的层次结构配置的时候。yml支持声明map,数组，list，字符串，boolean值，数值，NULL，日期，基本满足开发过程中的所有配置。

下面是一个非常基本的yml文件：

` server: url: http: //localhost myapp: name: MyApplication threadCount: 4... 复制代码`

等同于以下的application.properties文件：

` server.url=http: //localhost server.myapp.name=MyApplication server.myapp.threadCount= 4... 复制代码`

demo中的yml文件如下：

` server: url: http: //myapp.org app: name: MyApplication threadCount: 10 users: - Jacob - James 复制代码`

### 2.yml属性获取配置 ###

访问yml属性的一种方法是使用 ` @Value("$ {property}")` 注释,但是随着配置树形结构以及数量的增加，代码可读性也随之降低，更不利于bean的管理。笔者发现另一种优雅的方法可以确保强类型bean的管理以及更方便的验证我们的程序配置。

为了实现这一点，我们将创建一个 ` @ConfigurationProperties` 类ServerProperties，它映射一组相关的属性：

` import lombok.Data; import org.springframework.boot.context.properties.ConfigurationProperties; import java.util.ArrayList; import java.util.List; /** * @program : simple-demo * @description : 映射属性 (server节点) * @author : CaoTing * @date : 2019/6/3 **/ @Data @ConfigurationProperties ( "server" ) public class ServerProperties { private String url; private final App app = new App(); public App getApp () { return app; } public static class App { private String name; private String threadCount; private List<String> users = new ArrayList<>(); // TODO getter and setter } } 复制代码`

请注意，我们可以创建一个或多个@ConfigurationProperties类。

定义我们的springboot 注册配置类ApplicationConfig：

` import org.springframework.boot.context.properties.EnableConfigurationProperties; import org.springframework.context.annotation.Configuration; /** * @program : simple-demo * @description : 注册所有映射属性类 { }中用逗号分隔即可注册多个属性类 * @author : CaoTing * @date : 2019/6/3 **/ @Configuration @EnableConfigurationProperties ({ServerProperties.class}) public class ApplicationConfig { } 复制代码`

这里已经提到了要在@EnableConfigurationProperties中注册的属性类列表。

### 3.访问yml属性 ###

现在可以通过使用创建的@ConfigurationProperties bean来访问yml属性。可以像任何常规的Spring bean一样注入这些属性bean，测试类如下：

` import com.caotinging.ymldemo.application.YmlValueApplication; import com.caotinging.ymldemo.config.ServerProperties; import org.junit.Test; import org.junit.runner.RunWith; import org.springframework.beans.factory.annotation.Autowired; import org.springframework.boot.test.context.SpringBootTest; import org.springframework.test.context.junit4.SpringJUnit4ClassRunner; /** * @program : simple-demo * @description : 单元测试类 * @author : CaoTing * @date : 2019/6/3 **/ @RunWith (SpringJUnit4ClassRunner.class) @SpringBootTest (classes = YmlValueApplication.class) public class AppYmlValueTest { @Autowired private ServerProperties config; @Test public void printConfigs () { System.out.println( this.config.getUrl()); System.out.println( this.config.getApp().getName()); System.out.println( this.config.getApp().getThreadCount()); System.out.println( this.config.getApp().getUsers()); } } 复制代码`

测试结果如下:

![测试结果](https://user-gold-cdn.xitu.io/2019/6/3/16b1ca848fe8205b?imageView2/0/w/1280/h/960/ignore-error/1)

### 4.补充 ###

因为有小伙伴不太清楚具体用途。笔者补充一下两者的优缺点吧。

Spring Boot通过ConfigurationProperties注解从配置文件中获取属性。从上面的例子可以看出ConfigurationProperties注解可以通过设置prefix指定需要批量导入的数据。支持获取字面值，集合，Map，对象等复杂数据。ConfigurationProperties注解还有其他特点呢？它和Spring的Value注解又有什么区别呢？

#### 一）ConfigurationProperties和@Value优缺点 ####

ConfigurationProperties注解的优缺点

一、可以从配置文件中批量注入属性；

二、支持获取复杂的数据类型；

三、对属性名匹配的要求较低，比如user-name，user_name，userName，USER_NAME都可以取值；

四、支持JAVA的JSR303数据校验；

五、缺点是不支持SpEL表达式；

六、确保强类型bean的管理，更方便的验证程序配置；

Value注解的优缺点正好相反，它只能一个个配置注入值；不支持数组、集合等复杂的数据类型；不支持数据校验；对属性名匹配有严格的要求。最大的特点是支持SpEL表达式，使其拥有更丰富的功能。

#### 回答评论区小伙伴的疑问 ####

第一个属性就是个数组的时候怎么办呢？

yml示例如下

` orgs: - math - chinese - english 复制代码`

只需要如下这样一个bean就能轻松获取哦。

` /** * @program : simple-demo * @description : 映射Org属性 * @author : CaoTing * @date : 2019/6/6 **/ @Data @ConfigurationProperties public class OrgProperties { private List<String> orgs; } 复制代码`

不过别忘记在注册类那里注册一下

` @EnableConfigurationProperties ({ServerProperties.class, OrgProperties.class}) 复制代码`

github仓库也更新了这部分代码，欢迎前往测试哦，地址在文首。

还有map以及其他更复杂的数据结构都可以实现，我就不一个个测试了。

**原创不易，转载请附上原文链接** ，有帮助的话点个赞吧，笔芯。