# 为什么越来越多的开发者选择使用Spring Boot？ #

前几篇：

[Spring核心技术原理-（1）-通过Web开发演进过程了解一下为什么要有Spring?]( https://juejin.im/post/5cea202d6fb9a07eb94f67f1 )

[Spring核心技术原理-（2）-通过Web开发演进过程了解一下为什么要有Spring AOP?]( https://juejin.im/post/5cea206d6fb9a07ef710503c )

[Spring核心技术原理-（3）-Spring历史版本变迁和如今的生态帝国]( https://juejin.im/post/5cea20afe51d455a2f2201e4 )

[Spring核心技术原理-（4）-三条路线告诉你如何掌握Spring IoC容器的核心原理]( https://juejin.im/post/5cea2c02e51d4556f76e8001 )

## 一、Web应用开发背景 ##

使用Java做Web应用开发已经有近20年的历史了，从最初的Servlet1.0一步步演化到现在如此多的框架、库以及整个生态系统。经过这么长时间的发展，Java作为一个成熟的语言，也演化出了非常成熟的生态系统，这也是许多公司采用Java作为主流的语言进行服务器端开发的原因，也是为什么Java一直保持着非常活跃的用户群体的原因。

![Java EE开发生态图](https://user-gold-cdn.xitu.io/2019/6/5/16b2694b44904ac2?imageView2/0/w/1280/h/960/ignore-error/1)

最受Java开发者喜好的框架当属Spring，Spring也成为了在Java EE开发中真正意义上的标准，但是随着新技术的发展，脚本语言大行其道的时代（Node JS，Ruby，Groovy，Scala等），Java EE使用Spring逐渐变得笨重起来，大量的XML文件存在与项目中，繁琐的配置，整合第三方框架的配置问题，低下的开发效率和部署效率等等问题。

这些问题在不断的社区反馈下，Spring团队也开发出了相应的框架：Spring Boot。Spring Boot可以说是至少近5年来Spring乃至整个Java社区最有影响力的项目之一，也被人看作是：Java EE开发的颠覆者！

## 二、Spring Boot解决的问题 ##

(1) Spring Boot使编码变简单

(2) Spring Boot使配置变简单

(3) Spring Boot使部署变简单

(4) Spring Boot使监控变简单

(5) Spring的不足

## 三、Spring Boot的优点 ##

官方地址： [spring.io/projects/sp…]( https://link.juejin.im?target=https%3A%2F%2Fspring.io%2Fprojects%2Fspring-boot )

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694b4ae57f11?imageView2/0/w/1280/h/960/ignore-error/1)

Spring Boot继承了Spring的优点，并新增了一些新功能和特性：

（1）SpringBoot是伴随着Spring4.0诞生的，一经推出，引起了巨大的反向； （2）从字面理解，Boot是引导的意思，因此SpringBoot帮助开发者快速搭建Spring框架； （3）SpringBoot帮助开发者快速启动一个Web容器； （4）SpringBoot继承了原有Spring框架的优秀基因； （5）SpringBoot简化了使用Spring的过程； （6）Spring Boot为我们带来了脚本语言开发的效率，但是Spring Boot并没有让我们意外的新技术，都是Java EE开发者常见的额技术。

## 四、Spring Boot主要特性 ##

（1）遵循“习惯优于配置”的原则，使用Spring Boot只需要很少的配置，大部分的时候我们直接使用默认的配置即可； （2）项目快速搭建，可以无需配置的自动整合第三方的框架； （3）可以完全不使用XML配置文件，只需要自动配置和Java Config； （4）内嵌Servlet容器，降低了对环境的要求，可以使用命令直接执行项目，应用可用jar包执行：java -jar； （5）提供了starter POM, 能够非常方便的进行包管理, 很大程度上减少了jar hell或者dependency hell； （6）运行中应用状态的监控； （7）对主流开发框架的无配置集成； （8）与云计算的天然继承；

## 五、Spring Boot的核心功能 ##

**（1）独立运行的Spring项目**

Spring Boot可以以jar包的形式进行独立的运行，使用： ` java -jar xx.jar` 就可以成功的运行项目，或者在应用项目的主程序中运行main函数即可；

**（2）内嵌的Servlet容器**

内嵌容器，使得我们可以执行运行项目的主程序main函数，实现项目的快速运行；

主程序代码SpringbootDemoApplication.java

` package com.springboot.demo.helloworld; import org.springframework.boot.SpringApplication; import org.springframework.boot.autoconfigure.SpringBootApplication; @SpringBootApplication public class SpringBootHelloWorldApplication { public static void main(String[] args) { SpringApplication.run(SpringBootHelloWorldApplication.class, args); } } 复制代码`

**（3）提供starter简化Manen配置**

Spring Boot提供了一系列的starter pom用来简化我们的Maven依赖,下边是创建一个web项目中自动包含的依赖，使用的starter pom依赖为： ` spring-boot-starter-web`

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694a8fd09872?imageView2/0/w/1280/h/960/ignore-error/1)

Spring Boot官网还提供了很多的starter pom，请参考：

[docs.spring.io/spring-boot…]( https://link.juejin.im?target=https%3A%2F%2Fdocs.spring.io%2Fspring-boot%2Fdocs%2F2.0.4.RELEASE%2Freference%2Fhtmlsingle%2F%23using-boot-starter )

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694a9708442b?imageView2/0/w/1280/h/960/ignore-error/1)

**（4）自动配置Spring**

Spring Boot会根据我们项目中类路径的jar包/类，为jar包的类进行自动配置Bean，这样一来就大大的简化了我们的配置。当然，这只是Spring考虑到的大多数的使用场景，在一些特殊情况，我们还需要自定义自动配置；

**（5）应用监控**

注意：以前的版本还支持这个功能，目前使用的2.0.4.RELEASE已经不再支持此功能！

Spring Boot提供了基于http、ssh、telnet对运行时的项目进行监控；这个听起来是不是很炫酷！

示例：以SSH登录为例

1、首先，添加starter pom依赖

` <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-remote-shell</artifactId> </dependency> 复制代码`

2、运行项目,此时在控制台中会出现SSH访问的密码：

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694b681dd260?imageView2/0/w/1280/h/960/ignore-error/1)

3、使用SecureCRT登录到我们的程序，端口为2000，用户为user：

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694b5f12c623?imageView2/0/w/1280/h/960/ignore-error/1)

密码就是刚才的shell access；

但是当我点击连接的时候，出现错误：

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694b7cfc22c4?imageView2/0/w/1280/h/960/ignore-error/1)

显然是SecureCRT的版本不支持，所以就放弃了这个，使用Git Bash：

` ssh -p 2000 user@127.0.0.1 复制代码`

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694b5fc4c7a3?imageView2/0/w/1280/h/960/ignore-error/1)

剩下的事情，大家自己玩吧！

**（6）无代码生成和XML配置**

Spring Boot神奇的地方不是借助于代码生成来实现的，而是通过条件注解的方式来实现的，这也是Spring 4.x的新特性。

## 六、Spring Boot的快速搭建案例 ##

下边使用的是IDEA快速搭建一个Spring Boot项目

（1）File----New---New Project

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694bb16bdad1?imageView2/0/w/1280/h/960/ignore-error/1)

（2）点击Next填写相应的信息

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694bb1d0c446?imageView2/0/w/1280/h/960/ignore-error/1)

（3）点击Next，选择Dependencies，这里创建Web项目选择-----Web：

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694bd5cd5bd3?imageView2/0/w/1280/h/960/ignore-error/1)

（4）点击Next，设置项目名称，这里默认设置，点击Next之后，项目等一下就创建好了

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694bc02be9d9?imageView2/0/w/1280/h/960/ignore-error/1)

找到应用程序的主函数，运行即可：

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694bd840a5ca?imageView2/0/w/1280/h/960/ignore-error/1)

注意，在pom文件里的java版本这个要和你的机子上一致！我的是1.8，默认创建项目的时候为1.8：

` <properties> <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding> <project.reporting.outputEncoding>UTF-8</project.reporting.outputEncoding> <java.version>1.8</java.version> </properties> 复制代码`

## 七、案例代码 ##

GitOS 项目地址：

[gitee.com/xuliugen/sp…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fxuliugen%2Fspring-boot-unofficial-guide%2Ftree%2Fmaster%2Fspring-boot-hello-world )

![这里写图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b2694bd87e7dce?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06d5fb06f8a5f?imageView2/0/w/1280/h/960/format/png/ignore-error/1)

【视频福利】2T免费学习视频， 搜索或扫描上述二维码关注微信公众号：Java后端技术（ID: JavaITWork）,和20万人一起学Java！回复： **1024** ，即可免费获取！内含SSM、Spring全家桶、微服务、MySQL、MyCat、集群、分布式、中间件、Linux、网络、多线程，Jenkins、Nexus、Docker、ELK等等免费学习视频，持续更新！ ( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI1NDQ3MjQxNA%3D%3D%26amp%3Bmid%3D100005643%26amp%3Bidx%3D1%26amp%3Bsn%3Ded08bcf127fc549202ff273abeeb3d1a%26amp%3Bchksm%3D69c5eeba5eb267ac550ed247d72d6c43e11551f25fe80f8ea6a3146682ce2266dad880ce48de%23rd )