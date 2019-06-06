# 一行命令同时修改maven项目中多个mudule的版本号 #

Maven，是一个Java开发比较常用的项目管理工具，可以对 Java 项目进行构建、依赖管理。

对于很多Java程序员来说，分层架构都是不陌生的，至少MVC三层架构都是不陌生的，甚至有人说："Any problem in computer science can be solved by anther layer of indirection."

想要在代码中进行分层，比较好的做法就是创建多module的项目

` maven-parent (Maven Project) |- maven-dao (Maven Module) |- pom.xml |- maven-service (Maven Module) |- pom.xml |- maven-view (Maven Module) |- pom.xml |- pom.xml 复制代码`

以上项目，主要有三个模块，一般通过Maven进行模块间关系的管理。如：

最外层的pom.xml中，定义以下内容：

` <modelVersion>4.0.0</modelVersion> <artifactId>hollis-test</artifactId> <groupId>com.hollis.lab</groupId> <name>test</name> <packaging>pom</packaging> <version>1.0.0</version> <modules> <module>maven-dao</module> <module>maven-service</module> <module>maven-view</module> </modules> 复制代码`

然后在每个子模块中定义以下内容：

` <modelVersion>4.0.0</modelVersion> <parent> <artifactId>hollis-test</artifactId> <groupId>com.hollis.lab</groupId> <version>1.0.0</version> </parent> <artifactId>maven-service</artifactId> <name>com.hollis.lab</name> <packaging>jar</packaging> 复制代码`

这样，就形成了一个父子模块的关系。

但是，这样的项目，在版本升级的时候就会比较麻烦，因为要遍历的修改所有pom中的版本号。比

如要把1.0.0升级到1.0.1，那么就需要把所有的pom中的version都改掉。

这个人肉修改的过程既繁琐又容易出错，那么有没有什么办法可以代替人肉修改呢？

答案是有的。

### 一行命令修改所有版本号 ###

maven之所以强大，是因为他有一个牛X的插件机制。我们可以借助一个插件来实现这个功能。

这个插件就是versions-maven-plugin。使用方法也很简单，就是在最外层的pom文件中，增加以下插件配置：

` </plugins> <plugin> <groupId>org.codehaus.mojo</groupId> <artifactId>versions-maven-plugin</artifactId> <version>2.7</version> <configuration> <generateBackupPoms>false</generateBackupPoms> </configuration> </plugin> </plugins> 复制代码`

generateBackupPoms用于配置是否生成备份Pom，用于版本回滚。

配置好插件后，执行命令

` mvn versions:set -DnewVersion=1.0.1 复制代码`

即可降以上例子中的所有版本号修改成1.0.1。

为了方便使用，还可以在linux上设置别名，如：

` alias mvs='mvs() { mvn versions:set -DnewVersion=$1 }; mvs' 复制代码`

即可使用 ` mvs` 命令一键修改版本号。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1af32f142d564?imageView2/0/w/1280/h/960/ignore-error/1)