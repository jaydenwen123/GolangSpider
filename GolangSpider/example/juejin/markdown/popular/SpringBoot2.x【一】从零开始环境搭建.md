# SpringBoot2.x【一】从零开始环境搭建 #

对于之前的Spring框架的使用，各种配置文件XML、properties一旦出错之后错误难寻，这也是为什么SpringBoot被推上主流的原因，SpringBoot的配置简单，说5分钟能从框架的搭建到运行也不为过.
现在更是微服务当道，所以在此总结下SpringBoot的一些知识，新手教程.
Gradle是一个基于Apache Ant和Apache Maven概念的项目自动化构建开源工具,它使用一种基于Groovy语言来声明项目设置.也就是和Maven差不多的项目构建工具.

### 1. Maven 与 Gradle 对比 ###

maven要引入依赖 pom.xml

` <!-- https://mvnrepository.com/artifact/org.springframework.boot/spring-boot-starter-web --> < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-web </ artifactId > < version > 2.1.5.RELEASE </ version > </ dependency > 复制代码`

而Gradle引入 build.gradle

` implementation 'org.springframework.boot:spring-boot-starter-web' 复制代码`

Gradle本地安装教程
windows ： [www.cnblogs.com/linkstar/p/…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Flinkstar%2Fp%2F7899191.html ) Mac_OS ： [www.jianshu.com/p/e9d035f30…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Flinkstar%2Fp%2F7899191.html )

优点: Gradle 相当于 Maven 与 Ant 的合体
缺点: 对于微服务多项目的子类引用,不如 Maven

### 2.在官网快速创建SpringBoot项目 ###

下面开始进入正题：

进入 [start.spring.io/]( https://link.juejin.im?target=https%3A%2F%2Fstart.spring.io%2F ) 生成一个初始项目

![快速开始](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f12c13111?imageView2/0/w/1280/h/960/ignore-error/1)

这里会下载一个zip的项目压缩包

### 3. 使用Gradle导入SpringBoot项目 ###

demo.zip解压之后记得复制下demo文件夹放的路径
在此用的开发工具是IntelliJ IDEA
下面是导入流程： IDEA里点击File -> Open -> 粘贴刚刚的demo文件夹路径 -> 找到build.gradle双击
-> Open as Peoject -> 等待Gradle加载完就好,看不明白看下图

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f12fb4185?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f12b1b409?imageView2/0/w/1280/h/960/ignore-error/1)

(可选) 更改项目名
修改 settings.gradle

` rootProject.name = 'SpringBoot-demo' 复制代码`

去文件夹把项目文件夹名称改了
重新导入, 到此, 更改项目名结束

打开之后Gradle加载下载的特别慢，要换成国内源，打开build.gradle配置文件用下面的替换

**build.gradle**

` /** buildscript中的声明是gradle脚本自身需要使用的资源。 * 可以声明的资源包括依赖项、第三方插件、maven仓库地址等 */ plugins { id 'org.springframework.boot' version '2.1.5.RELEASE' id 'java' } apply plugin: 'io.spring.dependency-management' group = 'com.example' version = '0.0.1-SNAPSHOT' sourceCompatibility = '1.8' //让工程支持IDEA的导入 apply plugin: 'idea' repositories { //使用国内源下载依赖 maven { url 'http://maven.aliyun.com/nexus/content/groups/public/' } mavenCentral() } dependencies { implementation 'org.springframework.boot:spring-boot-starter-web' implementation 'com.alibaba:druid:1.1.11' testImplementation 'org.springframework.boot:spring-boot-starter-test' } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f12dfca0d?imageView2/0/w/1280/h/960/ignore-error/1)

### 4. SpringBoot项目启动 ###

启动前准备
依据下图把 DemoApplication 启动类 移到包最外层
启动类相当于管理项目的负责人，你把他扔到与控制层同级肯定出错不是;

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f1302a1d1?imageView2/0/w/1280/h/960/ignore-error/1)

** TestController.java **

` package com.example.controller; import org.springframework.web.bind.annotation.GetMapping; import org.springframework.web.bind.annotation.RestController; /** * 这里的 @RestController 相当于 @ResponseBody + @Controller * 使用 @RestController 相当于使每个方法都加上了 @ResponseBody 注解 * created by cfa 2018-11-06 下午 11:30 **/ @RestController public class TestController { /** * 这里的 @GetMapping 相当于 @RequestMapping (value = "/hello", method = RequestMethod.GET) * created by cfa 2018-11-06 下午 11:29 **/ @GetMapping ( "hello" ) public String test () { return "i love java" ; } } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f13be46c9?imageView2/0/w/1280/h/960/ignore-error/1)

启动成功之后访问 [http://localhost:8080/hello]( https://link.juejin.im?target=http%3A%2F%2Flocalhost%3A8080%2Fhello )

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f3e808c88?imageView2/0/w/1280/h/960/ignore-error/1) 上图成功代表项目可以访问了

### 5.配置application.yml ###

什么是yml? YML文件格式是YAML (YAML Aint Markup Language)编写的文件格式，YAML是一种直观的能够被电脑识别的的数据数据序列化格式，并且容易被人类阅读，容易和脚本语言交互的，可以被支持YAML库的不同的编程语言程序导入，比如： C/C++, Ruby, Python, Java, Perl, C#, PHP等。

听不懂吧，其实我也看不明白
就是相当于xml，properties的配置文件,看的更直观，上代码吧还是

` # 下述properties spring.resources.locations= classpath:/templates # 改为yml格式之后 spring: resources: static-locations: classpath:/templates 复制代码`

yml需要注意，冒号(:)后面要跟空格，第二级和第一级要在上下行用一个Tab的距离

** application.yml **

` server: port: 8080 spring: datasource: type : com.alibaba.druid.pool.DruidDataSource driver-class-name: com.mysql.jdbc.Driver url: jdbc:mysql://127.0.0.1:3306/dovis?characterEncoding=utf-8 username: root password: root mvc: view: suffix: .html resources: static-locations: classpath:/templates 复制代码`

欢迎关注微信公众号

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25c7f42ed2129?imageView2/0/w/1280/h/960/ignore-error/1)