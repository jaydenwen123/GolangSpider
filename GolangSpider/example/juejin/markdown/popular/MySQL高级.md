# MySQL高级 #

> 
> 
> 
> 个人技术博客 [www.zhenganwen.top](
> https://link.juejin.im?target=http%3A%2F%2Fwww.zhenganwen.top )
> 
> 

## 本文大纲 ##

![SQL高级](https://user-gold-cdn.xitu.io/2019/6/3/16b19039f9c1d785?imageView2/0/w/1280/h/960/ignore-error/1)

## 环境 ##

* 

win10-64

* 

MySQL Community Server 5.7.1

` mysqld –version` 可查看版本

* 

[官方文档]( https://link.juejin.im?target=https%3A%2F%2Fdev.mysql.com%2Fdoc%2Frefman%2F5.7%2Fen%2F )

## SQL执行顺序 ##

### 手写顺序 ###

我们可以将手写SQL时遵循的格式归结如下：

` select distinct <select_list> from <left_table> <join_type> join <right_table> on <join_condition> where <where_condition> group by <group_by_list> having <having_condition> order by <order_by_condition> limit < offset >,< rows > 复制代码`

* ` distinct` ，用于对查询出的结果集去重（若查出各列值相同的多条结果则只算一条）
* ` join` ，关联表查询，若将两个表看成两个集合，则能有7种不同的查询效果（将在下节介绍）。
* ` group by` ，通常与合计函数结合使用，将结果集按一个或多个列值分组后再合计
* ` having` ，通常与合计函数结合使用，弥补 ` where` 条件中无法使用函数
* ` order by` ，按某个标准排序，结合 ` asc/desc` 实现升序降序
* ` limit` ，如果跟一个整数 ` n` 则表示返回前 ` n` 条结果；如果跟两个整数 ` m,n` 则表示返回第 ` m` 条结果之后的 ` n` 条结果（不包括第 ` m` 条结果）

### MySQL引擎解析顺序 ###

而我们将SQL语句发给MySQL服务时，其解析执行的顺序一般是下面这样：

` from <left_table> on <join_condition> <join_type> join <right_table> where <where_condition> group by <group_by_list> having <having_condition> select <select_list> order by <order_by_condition> limit offset,rows 复制代码`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b19039f9a37a0e?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 了解这个对于后续分析SQL执行计划提供依据。
> 
> 

## 七种Join方式 ##

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b19039f9b283df?imageView2/0/w/1280/h/960/ignore-error/1)

下面我们创建部门表 ` tbl_dept` 和员工表 ` tbl_emp` 对上述7种方式进行逐一实现：

* 部门表：主键 ` id` 、部门名称 ` deptName` ，部门楼层 ` locAdd`

` mysql> CREATE TABLE `tbl_dept` ( -> `id` INT(11) NOT NULL AUTO_INCREMENT, -> `deptName` VARCHAR(30) DEFAULT NULL, -> `locAdd` VARCHAR(40) DEFAULT NULL, -> PRIMARY KEY (`id`) -> ) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8; 复制代码`

* 员工表：主键 ` id` ，姓名 ` name` 、所属部门 ` deptId`

` mysql> CREATE TABLE `tbl_emp` ( -> `id` INT(11) NOT NULL AUTO_INCREMENT, -> `name` VARCHAR(20) DEFAULT NULL, -> `deptId` INT(11) DEFAULT NULL, -> PRIMARY KEY (`id`), -> KEY `fk_dept_id` (`deptId`) -> #CONSTRAINT `fk_dept_id` FOREIGN KEY (`deptId`) REFERENCES `tbl_dept` (`id`) -> ) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8; 复制代码`

插入一些测试数据：

` mysql> INSERT INTO tbl_dept(deptName,locAdd) VALUES('技术部',11); Query OK, 1 row affected (0.07 sec) mysql> INSERT INTO tbl_dept(deptName,locAdd) VALUES('美工部',12); Query OK, 1 row affected (0.08 sec) mysql> INSERT INTO tbl_dept(deptName,locAdd) VALUES('总裁办',13); Query OK, 1 row affected (0.06 sec) mysql> INSERT INTO tbl_dept(deptName,locAdd) VALUES('人力资源',14); Query OK, 1 row affected (0.11 sec) mysql> INSERT INTO tbl_dept(deptName,locAdd) VALUES('后勤组',15); Query OK, 1 row affected (0.10 sec) mysql> insert into tbl_emp(name,deptId) values('jack',1); Query OK, 1 row affected (0.11 sec) mysql> insert into tbl_emp(name,deptId) values('tom',1); Query OK, 1 row affected (0.08 sec) mysql> insert into tbl_emp(name,deptId) values('alice',2); Query OK, 1 row affected (0.08 sec) mysql> insert into tbl_emp(name,deptId) values('john',3); Query OK, 1 row affected (0.13 sec) mysql> insert into tbl_emp(name,deptId) values('faker',4); Query OK, 1 row affected (0.10 sec) mysql> insert into tbl_emp(name) values('mlxg'); Query OK, 1 row affected (0.13 sec) mysql> select * from tbl_dept; +----+----------+--------+ | id | deptName | locAdd | +----+----------+--------+ | 1 | 技术部 | 11 | | 2 | 美工部 | 12 | | 3 | 总裁办 | 13 | | 4 | 人力资源 | 14 | | 5 | 后勤组 | 15 | +----+----------+--------+ 5 rows in set (0.00 sec) mysql> select * from tbl_emp; +----+-------+--------+ | id | name | deptId | +----+-------+--------+ | 1 | jack | 1 | | 2 | tom | 1 | | 3 | alice | 2 | | 4 | john | 3 | | 5 | faker | 4 | | 7 | ning | NULL | | 8 | mlxg | NULL | +----+-------+--------+ 7 rows in set (0.00 sec) 复制代码`

两表的关联关系如图所示：

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b19039fc67f609?imageView2/0/w/1280/h/960/ignore-error/1)

### 1、左连接（A独有+AB共有） ###

查询所有部门以及各部门的员工数：

` mysql> select t1.id,t1.deptName,count(t2.name) as emps from tbl_dept t1 left join tbl_emp t2 on t2.deptId=t1.id group by deptName order by id; +----+----------+------+ | id | deptName | emps | +----+----------+------+ | 1 | 技术部 | 2 | | 2 | 美工部 | 1 | | 3 | 总裁办 | 1 | | 4 | 人力资源 | 1 | | 5 | 后勤组 | 0 | +----+----------+------+ 5 rows in set (0.00 sec) 复制代码`

### 2、右连接（B独有+AB共有） ###

查询所有员工及其所属部门：

` mysql> select t2.id,t2.name,t1.deptName from tbl_dept t1 right join tbl_emp t2 on t2.deptId=t1.id; +----+-------+----------+ | id | name | deptName | +----+-------+----------+ | 1 | jack | 技术部 | | 2 | tom | 技术部 | | 3 | alice | 美工部 | | 4 | john | 总裁办 | | 5 | faker | 人力资源 | | 7 | ning | NULL | | 8 | mlxg | NULL | +----+-------+----------+ 7 rows in set (0.04 sec) 复制代码`

### 3、内连接（AB共有） ###

查询两表共有的数据：

` mysql> select deptName,t2.name empName from tbl_dept t1 inner join tbl_emp t2 on t1.id=t2.deptId; +----------+---------+ | deptName | empName | +----------+---------+ | 技术部 | jack | | 技术部 | tom | | 美工部 | alice | | 总裁办 | john | | 人力资源 | faker | +----------+---------+ 复制代码`

### 4、A独有 ###

即在（A独有+AB共有）的基础之上排除B即可（通过 ` b.id is null` 即可实现）：

` mysql> select a.deptName,b.name empName from tbl_dept a left join tbl_emp b on a.id=b.deptId where b.id is null; +----------+---------+ | deptName | empName | +----------+---------+ | 后勤组 | NULL | +----------+---------+ 复制代码`

### 5、B独有 ###

与（A独有）同理：

` mysql> select a.name empName,b.deptName from tbl_emp a left join tbl_dept b on a.deptId=b.id where b.id is null; +---------+----------+ | empName | deptName | +---------+----------+ | ning | NULL | | mlxg | NULL | +---------+----------+ 复制代码`

### 6、A独有+B独有 ###

使用 ` union` 将（A独有）和（B独有）联合在一起：

` mysql> select a.deptName,b.name empName from tbl_dept a left join tbl_emp b on a.id=b.deptId where b.id is null union select b.deptName,a.name emptName from tbl_emp a left join tbl_dept b on a.deptId=b.id where b.id is null; +----------+---------+ | deptName | empName | +----------+---------+ | 后勤组 | NULL | | NULL | ning | | NULL | mlxg | +----------+---------+ 复制代码`

### 7、A独有+AB公共+B独有 ###

使用 ` union` （可去重）联合（A独有+AB公共）和（B独有+AB公共）

` mysql> select a.deptName,b.name empName from tbl_dept a left join tbl_emp b on a.id=b.deptId union select a.deptName,b.name empName from tbl_dept a right join tbl_emp b on a.id=b.deptId; +----------+---------+ | deptName | empName | +----------+---------+ | 技术部 | jack | | 技术部 | tom | | 美工部 | alice | | 总裁办 | john | | 人力资源 | faker | | 后勤组 | NULL | | NULL | ning | | NULL | mlxg | +----------+---------+ 复制代码`

## 索引与数据处理 ##

### 什么是索引？ ###

索引是一种数据结构，在插入一条记录时，它从记录中提取（建立了索引的字段的）字段值作为该数据结构的元素，该数据结构中的元素被有序组织，因此在建立了索引的字段上搜索记录时能够借助二分查找提高搜索效率；此外每个元素还有一个指向它所属记录（数据库表记录一般保存在磁盘上）的指针，因此索引与数据库表的关系可类比于字典中目录与正文的关系，且目录的篇幅（索引所占的存储空间存储空间）很小。

数据库中，常用的索引数据结构是BTree（也称B-Tree，即Balance Tree，多路平衡查找树。Binary Search Tree平衡搜索二叉树是其中的一个特例）。

### 建立索引之后为什么快？ ###

索引是大文本数据的摘要，数据体积小，且能二分查找。这样我们在根据建立了索引的字段搜索时：其一，由表数据变为了索引数据（要查找的数据量显著减小）；其二，索引数据是有序组织的，搜索时间复杂度由线性的 ` O(N)` 变成了 ` O(logN)` （这是很可观的，意味着线性的 ` 2^32` 次操作被优化成了 ` 32` 次操作）。

### MySQL常用索引类型 ###

* 主键索引（ ` primary key` ），只能作用于一个字段（列），字段值不能为 ` null` 且不能重复。
* 唯一索引（ ` unique key` ），只能作用于一个字段，字段值可以为 ` null` 但不能重复
* 普通索引（ ` key` ），可以作用于一个或多个字段，对字段值没有限制。为一个字段建立索引时称为单值索引，为多个字段同时建立索引时称为复合索引（提取多个字段值组合而成）。

测试唯一索引的不可重复性和可为 ` null` ：

` mysql> create table `student` ( -> `id` int(10) not null auto_increment, -> `stuId` int(32) default null, -> `name` varchar(100) default null, -> primary key(`id`), -> unique key(`stuId`) -> ) engine=innodb auto_increment=1 default charset=utf8; mysql> insert into student(stuId,name) values('123456789','jack'); Query OK, 1 row affected (0.10 sec) mysql> insert into student(stuId,name) values('123456789','tom'); ERROR 1062 (23000): Duplicate entry '123456789' for key 'stuId' mysql> insert into student(stuId,name) values(null,'tom'); Query OK, 1 row affected (0.11 sec) 复制代码`

### 索引管理 ###

#### 创建索引 ####

* 

创建表（DDL）时创建索引

` mysql> create table `student` ( -> `id` int(10) not null auto_increment, -> `stuId` int(32) default null, -> `name` varchar(100) default null, -> primary key(`id`), -> unique key(`stuId`) -> ) engine=innodb auto_increment=1 default charset=utf8; 复制代码`
* 

创建索引语句： ` create [unique] index <index_name> on <table_name>(<col1>,<col2>...)`

` mysql> create index idx_name on student(name); Query OK, 0 rows affected (0.44 sec) Records: 0 Duplicates: 0 Warnings: 0 复制代码`
* 

更改表结构语句： ` alter table <table_name> add [unique] index <index_name> on (<col1>,<col2>....)`

` mysql> drop index idx_name on student; Query OK, 0 rows affected (0.27 sec) Records: 0 Duplicates: 0 Warnings: 0 mysql> alter table student add index idx_name(name); Query OK, 0 rows affected (0.32 sec) Records: 0 Duplicates: 0 Warnings: 0 复制代码`

#### 删除索引 ####

` drop index <index_name> on <table_name>`

#### 查看索引 ####

` SHOW INDEX FROM <table_name>`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b19039faaf7ac2?imageView2/0/w/1280/h/960/ignore-error/1)

## SQL执行计划——Explain ##

使用 ` EXPLAIN` 关键字可以模拟优化器执行SQL查询语句，从而知道MySQL是如何处理你的SQL语句的。分析你的查询语句或是表结构的性能瓶颈。

### 能干嘛 ###

通过 ` EXPLAIN` 分析某条SQL语句执行时的如下特征：

* 表的读取顺序（涉及到多张表时）
* 数据读取操作的操作类型
* 哪些索引可以使用
* 哪些索引被实际使用
* 表之间的引用
* 每张表有多少行被优化器查询

### 怎么玩 ###

格式为： ` explain <SQL语句>` ：

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b19039f9deb01e?imageView2/0/w/1280/h/960/ignore-error/1)

### 表头解析 ###

#### id ####

select查询的序列号，包含一组数字，表示查询中执行select子句或操作表的顺序。根据 ` id` 是否相同可以分为下列三种情况：

* 

所有表项的 ` id` 相同，如：

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a31fb42a8?imageView2/0/w/1280/h/960/ignore-error/1)

则上表中的3个表项按照从上到下的顺序执行，如读表顺序为 ` t1,t3,t2` 。由第一节提到的SQL解析顺序也可验证，首先 ` from t1,t2,t3` 表明此次查询设计到的表，由于没有 ` join` ，接着解析 ` where` 时开始读表，值得注意的是并不是按照 ` where` 书写的顺序，而是逆序，即先解析 ` t1.other_column=''` 于是读表 ` t1` ，然后 ` t1.id=t3.id` 读表 ` t3` ，最后 ` t1.id=t2.id` 读表 ` t2` 。解析顺序如下：

` from t1,t2,t3 where t1.other_column='', t1.id=t3.id, t1.id=t2.id select t2.* 复制代码`
* 

所有表项的 ` id` 不同：嵌套查询，id的序号会递增，id值越大优先级越高，越先被执行。如：

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a4ae35982?imageView2/0/w/1280/h/960/ignore-error/1)

对于多层嵌套的查询，执行顺序由内而外。解析顺序：

` from t2 where t2.id= from t1 where t1.id= from t3 where t3.other_column='' select t3.id select t1.id select t2.* 复制代码`

由第 ` 12,8,4` 行可知查表顺序为 ` t3,t1,t2` 。

* 

有的表项 ` id` 相同，有的则不同。 ` id` 相同的表项遵循结论1，不同的则遵循结论2

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a3e84d9b6?imageView2/0/w/1280/h/960/ignore-error/1)

解析顺序：

` from ( from t3 where t3.other_column='' select t3.id ) s1, t2 #s1是衍生表 where s1.id=t2.id select t2.* 复制代码`

由第 ` 6,11` 两行可以看出读表顺序为 ` t3,s1,t2`

#### select_type ####

该列常出现的值如下：

* 

SIMPLE，表示此SQL是简单的 ` select` 查询，查询中不包含子查询或者 ` union`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a33065a22?imageView2/0/w/1280/h/960/ignore-error/1)

* 

PRIMARY，查询中若包含任何复杂的子部分，最外层查询被标记为 ` PRIMARY`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a4e42f03c?imageView2/0/w/1280/h/960/ignore-error/1)

* 

SUBQUERY，在 ` select` 或 ` where` 列表中包含的子查询

* 

DERIVED，在 ` from` 子句中的子查询被标记为 ` DERIVED` （衍生）。MySQL会递归执行这些子查询, 把结果放在临时表里

* 

UNION， ` union` 右侧的 ` select`

* 

UNION RESULT， ` union` 的结果

#### table ####

表名，表示该表项是关于哪张表的，也可以是如形式：

* ` <derivedN>` ，表示该表是表项id为 ` N` 的衍生表
* ` <unionM,N` >，表示该表是表项id为 ` M` 和 ` N` 两者 ` union` 之后的结果

#### partition ####

如果启用了表分区策略，则该字段显示可能匹配查询的记录所在的分区

#### type ####

type显示的是访问类型，是较为重要的一个指标，结果值从最好到最坏依次是： system > const > eq_ref > ref > fulltext > ref_or_null > index_merge > unique_subquery > index_subquery > range > index > ALL 。

* 

` system` ， **表只有一行记录** （等于系统表），这是const类型的特列，平时不会出现，这个也可以忽略不计

* 

` const` ， **表示通过索引一次就找到了** ， ` const` 用于比较 ` primary key` 或者 ` unique key` 。因为只匹配一行数据，所以很快。若将主键置于where列表中，MySQL就能将该查询转换为一个常量

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a554e1d4f?imageView2/0/w/1280/h/960/ignore-error/1)

` mysql> select * from student; +----+-----------+------+ | id | stuId | name | +----+-----------+------+ | 1 | 123456789 | jack | | 3 | NULL | tom | +----+-----------+------+ 复制代码`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a5de4ef06?imageView2/0/w/1280/h/960/ignore-error/1)

* 

` eq_ref` ， **唯一性索引扫描** ，对于每个索引键，表中只有一条记录与之匹配。 **常见于主键或唯一索引扫描**

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a5ff3f1bc?imageView2/0/w/1280/h/960/ignore-error/1)

对于b中的每一条数据，从a的主键索引中查找id和其相等的

* 

` ref` ， **非唯一性索引扫描** ， **返回匹配某个单独值的所有行** 。本质上也是一种索引访问，它返回所有匹配某个单独值的行，然而，它可能会找到多个符合条件的行，所以他应该属于查找和扫描的混合体。（ **查找是基于有序性的能利用二分，而扫描则是线性的** ）

` mysql> create table `person` ( -> `id` int(32) not null auto_increment, -> `firstName` varchar(30) default null, -> `lastName` varchar(30) default null, -> primary key(`id`), -> index idx_name (firstName,lastName) -> ) engine=innodb auto_increment=1 default charset=utf8; 复制代码`

查询姓张的人：

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a863fa8c6?imageView2/0/w/1280/h/960/ignore-error/1)

* 

` range` ，根据索引的有序性检索特定范围内的行，通常出现在 ` between、<、>、in` 等范围检索中

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a67e10d62?imageView2/0/w/1280/h/960/ignore-error/1)

* 

` index` ，在索引中扫描，只需读取索引数据。

![1559474816109](图片格式不对)

由于复合索引 ` idx_name` 是基于（firstName，lastName）的，这种索引只能保证在整体上是按定义时的第一列（即firstName）有序的，当firstName相同时，再按lastName排序，如果不只两列则以此类推。也就是说在根据lastName查找时是无法利用二分的，只能做全索引扫描。

* 

` all` ，全表扫描，需要从磁盘上读取表数据。

> 
> 
> 
> 备注：一般来说，得保证查询至少达到 ` range` 级别，最好能达到 ` ref` 。
> 
> 

#### possible_keys ####

MySQL可以利用以快速检索行的索引。

#### key ####

MySQL执行时实际使用的索引。

#### key_len ####

* 

表示索引中每个元素最大字节数，可通过该列计算查询中使用的索引的长度（如何计算稍后详细结束）。

> 
> 
> 
> 在不损失精确性的情况下，长度越短越好。
> 
> 

* 

key_len显示的值为索引字段的最大可能长度，并非实际使用长度，即key_len是根据表定义计算而得，不是通过表内检索出的。

如何计算？首先我们要了解MySQL各数据类型所占空间：

* 

数值类型

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a88a90cde?imageView2/0/w/1280/h/960/ignore-error/1)

* 

日期类型（ ` datetime` 类型在MySQL5.6中字段长度是5个字节，在5.5中字段长度是8个字节）

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a88a90cde?imageView2/0/w/1280/h/960/ignore-error/1)

* 

字符串类型

![1559475897645](图片格式不对)

> 
> 
> 
> ` latin1` 编码的字符占1个字节， ` gbk` 编码的字符占2个字节， ` utf8` 编码的字符占3个字节。
> 
> 
> 
> ` c1 char(10)` 表示每行记录的 ` c1` 字段固定占用10个字节；而 ` c2 varchar(10)` 则不一定，如果某数据行的 `
> c2` 字段值只占3个字节，那么该数据行的 ` c2` 字段实际占5个字节，因为该类型字段所占空间大小是可变的，所以需要额外2个字节来保存字段值的长度，并且因为
> ` varchar` 最大字节数为65535，因此字段值最多占65533个字节。
> 
> 
> 
> 因此，
> 
> 
> 
> * 如果事先知道某字段存储的数据都是固定个数的字符则优先使用 ` char` 以节省存储空间。
> * 尽量设置 ` not null` 并将默认值设为 ` ‘’` 或 ` 0`
> 
> 
> 

以字符串类型字段的索引演示 ` key_len` 的计算过程（以 ` utf8` 编码为例）：

* 

索引字段为 ` char` 类型 + ` not null` ： ` key_len` = 字段申明字符个数 * 3（utf8编码的每个字符占3个字节）

` mysql> create table test( -> id int(10) not null auto_increment, -> primary key(id) -> ) engine=innodb auto_increment=1 default charset=utf8; mysql> alter table test add c1 char(10) not null; mysql> create index idx_c1 on test(c1); 复制代码`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a804965ee?imageView2/0/w/1280/h/960/ignore-error/1)

* 

索引字段为 ` char` 类型 + 可以为 ` null` ： ` key_len` = 字段申明字符个数 * 3 + 1（单独用一个字节表示字段值是否为 ` null` ）

` mysql> alter table test add c2 char(10) default null; mysql> create index idx_c2 on test(c2); 复制代码`

![1559477148209](图片格式不对)

* 

索引字段为 ` varchar` + ` not null` ， ` key_len` = 字段申明字符个数 * 3 + 2（用来保存字段值所占字节数）

` mysql> alter table test add c3 varchar(10) not null; mysql> create index idx_c3 on test(c3); 复制代码`

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a8da9c0ec?imageView2/0/w/1280/h/960/ignore-error/1)

* 

` varchar` + 可以为 ` null` ， ` key_len` = 字段申明字符个数 * 3 + 2 + 1（用来标识字段值是否为 ` null` ）

> 
> 
> 
> 根据这个值，就可以判断索引使用情况，特别是在使用复合索引时判断组成该复合索引的多个字段是否都能被查询用到。
> 
> 
> 
> 如：
> 
> ` mysql> desc person;
> +-----------+-------------+------+-----+---------+----------------+ |
> Field | Type | Null | Key | Default | Extra |
> +-----------+-------------+------+-----+---------+----------------+ | id |
> int(32) | NO | PRI | NULL | auto_increment | | firstName | varchar(30) |
> YES | MUL | NULL | | | lastName | varchar(30) | YES | | NULL | |
> +-----------+-------------+------+-----+---------+----------------+ 复制代码`
> 
> 
> 
> ![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a91ff3cd8?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 
> 
> 
> 
> ![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903a9383584e?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 
> 
> 前者使用了部分复合索引，而后者使用了全部，这在索引类型一节中也提到过，是由最左前缀（定义复合索引时的第一列 ）有序这一特性决定的。
> 
> 

#### ref ####

显示哪一列或常量被拿来与索引列进行比较以从表中检索行。

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903abb5e3347?imageView2/0/w/1280/h/960/ignore-error/1)

如上我们使用 ` ‘’` 到索引中检索行。

#### rows ####

根据表统计信息及索引选用情况，大致估算出找到所需的记录所需要读取的行数

#### Extra ####

包含不适合在其他列中显示但十分重要的额外信息：

* 

` Using filesort` ：说明mysql会对数据使用一个外部的索引排序，而不是按照表内的索引顺序进行读取。MySQL中无法利用索引完成的排序操作称为“文件排序”

` mysql> explain select * from person order by lastName\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: person partitions: NULL type: index possible_keys: NULL key: idx_name key_len: 186 ref: NULL rows: 1 filtered: 100.00 Extra: Using index; Using filesort 复制代码`
> 
> 
> 
> 
> 使用 ` \G` 代替 ` ;` 结尾可以使执行计划垂直显示。
> 
> 

` mysql> explain select * from person order by firstName,lastName\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: person partitions: NULL type: index possible_keys: NULL key: idx_name key_len: 186 ref: NULL rows: 1 filtered: 100.00 Extra: Using index 复制代码`
* 

` Using temporary` ：使用了临时表保存中间结果。 **MySQL在对查询结果聚合时使用临时表** 。常见于排序 ` order by` 和分组查询 ` group by` 。

` mysql> insert into person(firstName,lastName) values('张','三'); mysql> insert into person(firstName,lastName) values('李','三'); mysql> insert into person(firstName,lastName) values('王','三'); mysql> insert into person(firstName,lastName) values('李','明'); mysql> select lastName,count(lastName) from person group by lastName; +----------+-----------------+ | lastName | count(lastName) | +----------+-----------------+ | 三 | 3 | | 明 | 1 | +----------+-----------------+ mysql> explain select lastName,count(lastName) from person group by lastName\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: person partitions: NULL type: index possible_keys: idx_name key: idx_name key_len: 186 ref: NULL rows: 4 filtered: 100.00 Extra: Using index; Using temporary; Using filesort 复制代码`
* 

` Using index` ：表示相应的select操作中使用了 **覆盖索引(Covering Index)** ，避免访问了表的数据行（需要读磁盘），效率不错！如果同时出现 ` Using where` ，表明索引被用来执行索引键值的查找；如果没有同时出现 ` Using where` ，表明索引用来读取数据而非执行查找动作。

> 
> 
> 
> 索引覆盖：就是 ` select` 的数据列只用从索引中就能够取得，不必读取数据行，MySQL可以利用索引返回 ` select` 列表中的字段，而不必根据索引再次读取数据文件，换句话说查询列要被所建的索引覆盖。
> 
> 
> 
> 
> 如果要使用覆盖索引，一定要注意 ` select` 列表中只取出需要的列，不可 ` select *` ，因为如果将所有字段一起做索引会导致索引文件过大，查询性能下降。
> 
> 
> 

* 

` Using where` ：查询使用到了 ` where` 语句

* 

` Using join buffer` ：使用了连接缓存

* 

` Impossible where` ： ` where` 子句的值总是 ` false` ，如

` select * from person where id = 1 and id = 2 ; 复制代码`

## 索引失效 ##

如果使用 ` explain` 分析SQL的执行计划时发现访问类型 ` type` 为 ` ALL` 或实际使用到的索引 ` key` 为 ` NULL` ，则说明该查询没有利用索引而导致了全表扫描，这是我们需要避免的。以下总结了利用索引的一些原则：

### 1、全值匹配我最爱 ###

根据常量在索引字段上检索时一定能够利用到索引。

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903ac0028de3?imageView2/0/w/1280/h/960/ignore-error/1)

这种方式

### 2、最佳左前缀法则 ###

对于复合索引检索时一定要遵循左前缀列在前的原则。

` mysql> alter table test add c5 varchar(10) default null, add c6 varchar(10) default null, add c7 varchar(10) default null; mysql> create index idx_c5_c6_c7 on test(c5,c6,c7); 复制代码`

如果没有左前缀列则不会利用索引：

` mysql> explain select * from test where c6=''\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: test partitions: NULL type: ALL possible_keys: NULL key: NULL key_len: NULL ref: NULL rows: 1 filtered: 100.00 Extra: Using where mysql> explain select * from test where c6='' and c7=''\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: test partitions: NULL type: ALL possible_keys: NULL key: NULL key_len: NULL ref: NULL rows: 1 filtered: 100.00 Extra: Using where 复制代码`

而只要最左前缀列在前，其他列可以不按顺序也可以不要，但最好不要那么做（按照定义复合索引时的列顺序能达到最佳效率）：

` mysql> explain select * from test where c5='' and c7=''\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: test partitions: NULL type: ref possible_keys: idx_c5_c6_c7 key: idx_c5_c6_c7 key_len: 33 ref: const rows: 1 filtered: 100.00 Extra: Using index condition 1 row in set, 1 warning (0.00 sec) mysql> explain select * from test where c5='' and c7='' and c6=''\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: test partitions: NULL type: ref possible_keys: idx_c5_c6_c7 key: idx_c5_c6_c7 key_len: 99 ref: const,const,const rows: 1 filtered: 100.00 Extra: NULL 1 row in set, 1 warning (0.00 sec) 复制代码`

最优的做法是：

` mysql> explain select * from test where c5=''\G mysql> explain select * from test where c5='' and c6=''\G mysql> explain select * from test where c5='' and c6='' and c7=''\G 复制代码`

### 3、不在列名上添加任何操作 ###

有时我们会在列名上进行计算、函数运算、自动/手动类型转换，这会直接导致索引失效。

` mysql> explain select * from person where left(firstName,1)='张'\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: person partitions: NULL type: index possible_keys: NULL key: idx_name key_len: 186 ref: NULL rows: 4 filtered: 100.00 Extra: Using where; Using index mysql> explain select * from person where firstName='张'\G *************************** 1. row *************************** id: 1 select_type: SIMPLE table: person partitions: NULL type: ref possible_keys: idx_name key: idx_name key_len: 93 ref: const rows: 1 filtered: 100.00 Extra: Using index 复制代码`

上面两条SQL同样是实现查找姓张的人，但在列名 ` firstName` 上使用了 ` left` 函数使得访问类型 ` type` 从 ` ref` （非唯一性索引扫描）降低到了 ` index` （全索引扫描）

### 4、存储引擎无法使用索引中范围条件右边的列 ###

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903ac33dadd8?imageView2/0/w/1280/h/960/ignore-error/1)

由上图可知 ` c6 > ‘a’` 右侧的列 ` c7` 虽然也在复合索引 ` idx_c5_c6_c7` 中，但由 ` key_len:66` 可知其并未被利用上。通常索引利用率越高，查找效率越高。

### 5、尽量使用索引覆盖 ###

尽量使查询列和索引列保持一致，这样就能避免访问数据行而直接返回索引数据。避免使用 ` select *` 除非表数据很少，因为 ` select *` 很大概率访问数据行。

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903ac577be38?imageView2/0/w/1280/h/960/ignore-error/1)

` Using index` 表示发生了索引覆盖

### 6、使用 != 或 <> 时可能会导致索引失效 ###

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903acf5ef16f?imageView2/0/w/1280/h/960/ignore-error/1)

### 7、not null对索引也有影响 ###

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903ae783acd8?imageView2/0/w/1280/h/960/ignore-error/1)

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903ae2bba451?imageView2/0/w/1280/h/960/ignore-error/1)

若 ` name` 的定义不是 ` not null` 则不会有索引未利用的情况。

### 8、like以通配符开头会导致索引失效 ###

` like` 语句以通配符 ` %` 开头无法利用索引会导致全索引扫描，而只以通配符结尾则不会。

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903af93d2e21?imageView2/0/w/1280/h/960/ignore-error/1)

### 9、join on的列只要有一个没索引则全表扫描 ###

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903b08b4d28e?imageView2/0/w/1280/h/960/ignore-error/1)

### 10、or两侧的列只要有一个没索引则全表扫描 ###

![image](https://user-gold-cdn.xitu.io/2019/6/3/16b1903b73112d98?imageView2/0/w/1280/h/960/ignore-error/1)

### 11、字符串不加单引号索引失效 ###

` mysql> explain select * from staff where name=123; 复制代码`
> 
> 
> 
> 打油诗：
> 
> 
> 
> 全值匹配我最爱，最左前缀要遵循。
> 
> 
> 
> 带头大哥不能死，中间兄弟不能断。
> 
> 
> 
> 索引列上少计算，范围之后全失效。
> 
> 
> 
> LIKE百分比最右，覆盖索引不写*。
> 
> 
> 
> 不等空值还有OR，ON的右侧要注意。
> 
> 
> 
> VAR引号不能丢，SQL优化有诀窍。
> 
>