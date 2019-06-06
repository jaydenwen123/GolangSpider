# MySQL 的COUNT(x)性能怎么样？ #

> 
> 
> 
> 做一个积极的人
> 
> 
> 
> 编码、改bug、提升自己
> 
> 
> 
> 我有一个乐园，面向编程，春暖花开！
> 
> 

> 
> 
> 
> x 可以代表： 主键id、字段、1、*
> 
> 

## 0 说明 ##

**对于count(主键id)来说**

innodb引擎会遍历整张表，把每一行的id值都取出来，返回给server层，server层判断id值不为空，就按行累加

**对于count(字段)来说**

如果这个字段定义为not null，一行行的从记录里面读出这个字段，判断不为空，则累加值

如果这个字段定义允许为null，那么执行的时候，判断到有可能为null,还要把值取出来在判断一下，不是null才累加

**对于count(1)来说**

innodb引擎遍历整张表，但不取值，返回给server层，server对于返回的每一行，放一个数字1进去，判断是不可能为空的，就按行累加

**对于count(*)**

并不会把全部字段取出来，而是专门做了优化，不取值，count(*)肯定不是null，按行累加

## 1 总结 ##

如果你要统计行数就用 ` count(*)` 或者 ` count(1)` ，推荐前者

如果要统计某个字段不为NULL值的个数就用 ` count(字段)`

在《**高性能MySQL》**中有如下：

* 

当mysql确认括号内的表达式值不可能为空时，实际上就是在统计行数

* 

如果mysql知道某列col不可能为NULL值，那么mysql内部会将count(col)表达式优化为count(*)

也就是说count(主键字段)和count(1)还是要优化到count(*)的。

在 **MySQL 5.7 Reference Manual** 的官方手册中： [dev.mysql.com/doc/refman/…]( https://link.juejin.im?target=https%3A%2F%2Fdev.mysql.com%2Fdoc%2Frefman%2F5.7%2Fen%2Fgroup-by-functions.html%23function_count )

有这么一段：

> 
> 
> 
> InnoDB handles SELECT COUNT(*) and SELECT COUNT(1) operations in the same
> way. There is no performance difference.
> 
> 

翻译： InnoDB以相同的方式处理SELECT COUNT（*）和SELECT COUNT（1）操作。没有性能差异。

所以这几个按照效率排序的话，count(字段)<count(主键id)<count(1)≈count(*)

**所以，尽量使用count(*)**

## 2 拓展 ##

在阿里巴巴的 Mysql数据库 >> ( 三) ) SQL

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b280bd555cf0f3?imageView2/0/w/1280/h/960/ignore-error/1)

**谢谢你的阅读，如果您觉得这篇博文对你有帮助，请点赞或者喜欢，让更多的人看到！祝你每天开心愉快！**

**不管做什么，只要坚持下去就会看到不一样！在路上，不卑不亢!**

**愿你我在人生的路上能都变成最好的自己，能够成为一个独挡一面的人**

![](https://user-gold-cdn.xitu.io/2019/6/5/16b280bd607c37b1?imageView2/0/w/1280/h/960/ignore-error/1)

© 每天都在变得更好的阿飞云