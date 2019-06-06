# 我就是不看好jpa #

知乎看到问题《SpringBoot开发使用Mybatis还是Spring Data JPA??》，顺手一答，讨论激烈。我实在搞不懂spring data jpa为啥选了hibernate作为它的实现，是“Gavin King”的裙带关系么？ DAO层搞来搞去，从jdbc到hibernate，从toplink到jdo，到现在MyBatis胜出，是有原因的。

目前，一些狗屁培训公司，还有一些网络课程，包括一些洋课程，为了做项目简单，很多会使用jpa。但到了公司发现根本就不是这样，这就是理论和现实的区别。

顶着被骂的风险，我整理发出来。这可不是争论php好还是java好性质了。

> 
> 
> 
> 忠告：精力有限，一定要先学MyBatis啊。jpa那一套，表面简单而已。
> 
> 

##### 以下是原始回答 #####

如果你经历过多个公司的产品迭代，尤其是复杂项目，你就会发现 ` Spring Data JPA` 这种东西是多么的不讨好。说实话， ` Mybatis` 的功能有时候都嫌多，有些纯粹是画蛇添足。

` jpa` 虽然是规范，但和 ` hibernate` 这种 ` ORM` 是长得比较像的（不说一些特殊场景的用法）。

**Spring Data JPA 是个玩具，只适合一些简单的映射关系** 。也不得不提一下被人吹捧的 ` querydsl` ，感觉有些自作聪明的讨巧。实在不想为了一个简单的DAO层，要学一些费力不讨好的东西。

` List<Person> persons = queryFactory.selectFrom(person) .where(person.children.size().eq( JPAExpressions.select(parent.children.size().max()) .from(parent))) .fetch(); 复制代码`

看看上面的查询语句，完全不如普通 ` SQL` 表达的清晰。要是紧急排查个问题，妈蛋...

jpa虽然有很多好处，比如和底层的SQL无关。但我觉得Spring Data JPA有以下坏处：

1、 屏蔽了SQL的优雅，发明了一种自己的查询方式。这种查询方式并不能够覆盖所有的SQL场景。

2、 增加了代码的复杂度，需要花更多的时间来理解DAO

3、DAO操作变的特别的分散，分散到多个java文件中，或者注解中（虽然也支持XML）。如果进行一些扫描，或者优化，重构成本大

4、不支持复杂的SQL，DBA流程不好切入

Mybatis虽然也有一些问题，但你更像是在写确切的SQL。它比Spring Data JPA更加轻量级。

你只要在其他地方调试好了SQL，只需要写到配置文件里，起个名字就可以用了，少了很多烧脑的转换过程。

Mybatis完全能够应对工业级的复杂SQL，甚至存储过程（不推荐）。个人认为Mybatis依然是搞复杂了，它其中还加了一些类似if else的编程语言特性。

你的公司如果有DBA，是不允许你乱用SQL的。用Mybatis更能切入到公司的流程上。

所以我认为：玩具项目或者快速开发，使用Spring Boot JPA。反之，Mybatis是首选。

# 一些有用的评论 #

你说的if else是指的SQL拼接吧，这可是MyBatis的重要功能之一，学起来一点儿也不复杂好吗？

最近也在研究持久层，可以充分利用这个jpa这个玩具，两者结合是不错的选择，jpa基本的单表操作，mybatis做复杂查询，开发效率高，降低sql维护成本，也为优化留下空间，当然这需要对spring-data-jpa做一些扩展

查询直接sql，其他的还是orm方便

mybatis主要是原生sql，对于其他没学习过jpa的开发人员而言降低了学习维护门槛，而且说真的jpa写了个锅你去追其实还是挺头疼的...

mybatis-plus整合之后基本curd不用纠结了，很多对对象操作然后直接save就好。复杂场景、联表就直接用原生sql就好，至于性能问题可以用sqlAdvice优化下。

jdbc template+代码生成器，更简单更高效

jpa这玩意，写一个简单的数据库操作，就比如说单表的操作，很好用。如果是多表，那就算了吧

spring boot 推荐jpa，知道为什么吗

native=true 想用原生查询也没人拦着你啊

你好像不是很懂hibernate和jpa…

不知道你是否清楚 jpa，hibernate，spring data jpa，还有querydsl 之间的关系。

总有一天你会知道数据库优先和程序优先的区别

jpa还有一个好处，那就是帅锅啊

# END #

来，越年轻越狂妄的家伙们，来喷我啊。

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0d3b2fb538421?imageView2/0/w/1280/h/960/ignore-error/1)