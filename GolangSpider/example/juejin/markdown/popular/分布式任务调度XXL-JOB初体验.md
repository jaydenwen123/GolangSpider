# 分布式任务调度XXL-JOB初体验 #

原文： [gcdd1993.github.io/分布式任务调度XXL-…]( https://link.juejin.im?target=https%3A%2F%2Fgcdd1993.github.io%2F%25E5%2588%2586%25E5%25B8%2583%25E5%25BC%258F%25E4%25BB%25BB%25E5%258A%25A1%25E8%25B0%2583%25E5%25BA%25A6XXL-JOB%25E5%2588%259D%25E4%25BD%2593%25E9%25AA%258C )

# 简介 #

[XXL-JOB]( https://link.juejin.im?target=http%3A%2F%2Fwww.xuxueli.com%2Fxxl-job%2F%23%2F ) 是一个轻量级分布式任务调度平台，其核心设计目标是开发迅速、学习简单、轻量级、易扩展。现已开放源代码并接入多家公司线上产品线，开箱即用。

官方文档很完善，不多赘述。本文主要是搭建 ` XXL-JOB` 和简单使用的记录。

# 搭建xxl-job-admin管理端 #

## 运行环境 ##

* Ubuntu 16.04 64位
* Mysql 5.7

### 安装Mysql ###

` $ sudo apt-get update $ sudo apt-get install mysql-server ## 设置mysql，主要是安全方面的，密码策略等 $ mysql_secure_installation ## 配置远程访问 $ sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf bind -address = 0.0.0.0 $ sudo service mysql restart $ sudo service mysql status ● mysql.service - MySQL Community Server Loaded: loaded (/lib/systemd/system/mysql.service; enabled; vendor preset: enabled) Active: active (running) since Wed 2019-06-05 13:23:41 HKT; 45s ago ... 复制代码`

### 创建数据库 ###

` $ mysql -u root -p mysql> CREATE database if NOT EXISTS `xxl-job` default character set utf8 collate utf8_general_ci; 复制代码`

### 创建用户 ###

` $ mysql -u root -p mysql> CREATE USER 'xxl-job' @ '%' IDENTIFIED BY 'xxlJob2019@' ; mysql> GRANT ALL PRIVILEGES ON `xxl-job`.* TO 'xxl-job' @ '%' ; 复制代码`

## 本地测试xxl-job-admin ##

### 拉取最新源码 ###

` $ git clone git@github.com:xuxueli/xxl-job.git $ cd xxl-job 复制代码`

### 导入项目 ###

我比较熟悉 ` Idea` 开发工具，所以这里使用 ` Idea` 的 ` Gradle` 项目进行演示。

打开 ` xxl-job` ，项目结构如下

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0d6e6c8590?imageView2/0/w/1280/h/960/ignore-error/1)

### 测试项目 ###

打开 ` xxl-job-admin/resources/application.properties` ，修改mysql连接信息

` ### xxl-job, datasource spring.datasource.url=jdbc:mysql://192.168.32.129:3306/xxl-job?Unicode=true&characterEncoding=UTF-8 spring.datasource.username=xxl-job spring.datasource.password=xxlJob2019@ 复制代码`

使用 ` /xxl-job/doc/db/tables_xxl_job.sql` 初始化数据库，初始化完应该如下图

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0d6cc15b94?imageView2/0/w/1280/h/960/ignore-error/1)

准备就绪后，就可以启动项目了，然后打开地址http://localhost:8080/xxl-job-admin将会看到首页

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0dabff2fc0?imageView2/0/w/1280/h/960/ignore-error/1)

## 部署 ##

### 打包调度中心 ###

` $ cd /xxl-job $ mvn install ... [INFO] xxl-job ............................................ SUCCESS [ 0.513 s] [INFO] xxl-job-core ....................................... SUCCESS [ 4.258 s] [INFO] xxl-job-admin ...................................... SUCCESS [ 5.525 s] [INFO] xxl-job-executor-samples ........................... SUCCESS [ 0.016 s] [INFO] xxl-job-executor-sample-spring ..................... SUCCESS [ 2.188 s] [INFO] xxl-job-executor-sample-springboot ................. SUCCESS [ 0.892 s] [INFO] xxl-job-executor-sample-jfinal ..................... SUCCESS [ 1.753 s] [INFO] xxl-job-executor-sample-nutz ....................... SUCCESS [ 1.316 s] [INFO] xxl-job-executor-sample-frameless .................. SUCCESS [ 0.358 s] [INFO] xxl-job-executor-sample-jboot ...................... SUCCESS [ 1.279 s] [INFO] ------------------------------------------------------------------------ [INFO] BUILD SUCCESS [INFO] ------------------------------------------------------------------------ [INFO] Total time: 18.549 s [INFO] Finished at: 2019-06-05T14:40:25+08:00 [INFO] ------------------------------------------------------------------------ 复制代码`

看到以上信息，说明我们打包成功了，在 ` /xxl-job/xxl-job-admin` 目录下会存在jar文件： ` xxl-job-admin-2.1.0-SNAPSHOT.jar`

### 部署到服务器 ###

` $ sudo apt install openjdk-8-jdk $ java -version openjdk version "1.8.0_212" OpenJDK Runtime Environment (build 1.8.0_212-8u212-b03-0ubuntu1.16.04.1-b03) OpenJDK 64-Bit Server VM (build 25.212-b03, mixed mode) $ sudo mkdir -p /data/xxl-job $ sudo cd /data/xxl-job ## 上传我们打包好的jar至此目录，并添加软连接 $ sudo ln -s xxl-job-admin-2.1.0-SNAPSHOT.jar current.jar ## 注册为system服务，可以达到异常重启，开机自启等目的 $ sudo vim /etc/systemd/system/xxl-job.service Description=xxl-job Service Daemon After=mysql.service [Service] Type=simple Environment= "JAVA_OPTS= -Xmx1024m -Xms1024m -XX:+UseG1GC -XX:MaxGCPauseMillis=200 -XX:NewRatio=3" ExecStart=java -jar /data/xxl-job/current.jar Restart=always WorkingDirectory=/data/xxl-job/ [Install] WantedBy=multi-user.target $ sudo systemctl enable xxl-job.service $ sudo service xxl-job start $ sudo service xxl-job status ● xxl-job.service Loaded: loaded (/etc/systemd/system/xxl-job.service; enabled; vendor preset: enabled) Active: active (running) since Wed 2019-06-05 15:30:08 HKT; 2min 34s ago ... 复制代码`

我们访问一下http://192.168.32.129:8080/xxl-job-admin：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0d6e798ce2?imageView2/0/w/1280/h/960/ignore-error/1)

# 测试任务调度 #

以上，我们的任务调度管理端已经搭建完成，接下来，让我们测试下任务调度。

直接使用自带的 ` SpringBoot` 测试项目 ` xxl-job-executor-sample-springboot` 进行测试，修改配置文件

` xxl-job-executor-sample-springboot=http://192.168.32.129:8080/xxl-job-admin 复制代码`

## 自定义任务 ##

编写一个简单的任务，打印100次当前序列

` package com.xxl.job.executor.service.jobhandler; import com.xxl.job.core.biz.model.ReturnT; import com.xxl.job.core.handler.IJobHandler; import com.xxl.job.core.handler.annotation.JobHandler; import com.xxl.job.core.log.XxlJobLogger; import org.springframework.stereotype.Component; import java.util.concurrent.TimeUnit; /** * TODO * * @author gaochen * @date 2019/6/5 */ @JobHandler (value= "gcddJobHandler" ) @Component public class GcddJobHandler extends IJobHandler { @Override public ReturnT<String> execute (String param) throws Exception { for ( int i = 0 ; i < 100 ; i++) { XxlJobLogger.log( "XXL-JOB, print " + i); TimeUnit.SECONDS.sleep( 1 ); } return SUCCESS; } } 复制代码`

## 启动执行器 ##

然后启动执行器，启动完成后，我们会发现管理页面的执行器列表会多出我们刚才启动的执行器

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0d724ebb3c?imageView2/0/w/1280/h/960/ignore-error/1)

## 添加任务 ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0d72cd3192?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c0dbe3d0401?imageView2/0/w/1280/h/960/ignore-error/1)

## 查看任务执行日志 ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c15590cf3b7?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到，任务已经按照我们的规划执行成功了，非常的方便。

# 结语 #

想要了解更详细的内容，请访问 [xxl-job官网]( https://link.juejin.im?target=http%3A%2F%2Fwww.xuxueli.com%2Fxxl-job%2F%23%2F )