# MySQL 快速创建千万级测试数据 #

> 
> 
> 
> 备注： 此文章的数据量在100W，如果想要千万级，调大数量即可,但是不要大量使用rand() 或者uuid() 会导致性能下降
> 
> 

## 背景 ##

在进行查询操作的性能测试或者sql优化时，我们经常需要在线下环境构建大量的基础数据供我们测试，模拟线上的真实环境。

> 
> 
> 
> 废话，总不能让我去线上去测试吧，会被DBA砍死的
> 
> 

## 创建测试数据的方式 ##

` 1. 编写代码，通过代码批量插库（本人使用过，步骤太繁琐，性能不高，不推荐） 2. 编写存储过程和函数执行（本文实现方式1） 3. 临时数据表方式执行 （本文实现方式2，强烈推荐该方式，非常简单，数据插入快速，100W，只需几秒） 4. 一行一行手动插入，（WTF，去死吧） 复制代码`

## 创建基础表结构 ##

> 
> 
> 
> 不管用何种方式，我要插在那张表总要创建的吧
> 
> 

` CREATE TABLE `t_user` ( `id` int(11) NOT NULL AUTO_INCREMENT, `c_user_id` varchar(36) NOT NULL DEFAULT '' , `c_name` varchar(22) NOT NULL DEFAULT '' , `c_province_id` int(11) NOT NULL, `c_city_id` int(11) NOT NULL, `create_time` datetime NOT NULL, PRIMARY KEY (`id`), KEY `idx_user_id` (`c_user_id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; 复制代码`

## 方式1： 采用存储过程和内存表 ##

* 创建内存表

` 利用 MySQL 内存表插入速度快的特点，我们先利用函数和存储过程在内存表中生成数据，然后再从内存表插入普通表中 CREATE TABLE `t_user_memory` ( `id` int(11) NOT NULL AUTO_INCREMENT, `c_user_id` varchar(36) NOT NULL DEFAULT '' , `c_name` varchar(22) NOT NULL DEFAULT '' , `c_province_id` int(11) NOT NULL, `c_city_id` int(11) NOT NULL, `create_time` datetime NOT NULL, PRIMARY KEY (`id`), KEY `idx_user_id` (`c_user_id`) ) ENGINE=MEMORY DEFAULT CHARSET=utf8mb4; 复制代码`

* 创建函数和存储过程

` # 创建随机字符串和随机时间的函数 mysql> delimiter $$ mysql> CREATE DEFINER=`root`@`%` FUNCTION `randStr`(n INT) RETURNS varchar(255) CHARSET utf8mb4 -> DETERMINISTIC -> BEGIN -> DECLARE chars_str varchar(100) DEFAULT 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789' ; -> DECLARE return_str varchar(255) DEFAULT '' ; -> DECLARE i INT DEFAULT 0; -> WHILE i < n DO -> SET return_str = concat(return_str, substring(chars_str, FLOOR(1 + RAND() * 62), 1)); -> SET i = i + 1; -> END WHILE; -> RETURN return_str; -> END$$ Query OK, 0 rows affected (0.00 sec) mysql> CREATE DEFINER=`root`@`%` FUNCTION `randDataTime`(sd DATETIME,ed DATETIME) RETURNS datetime -> DETERMINISTIC -> BEGIN -> DECLARE sub INT DEFAULT 0; -> DECLARE ret DATETIME; -> SET sub = ABS(UNIX_TIMESTAMP(ed)-UNIX_TIMESTAMP(sd)); -> SET ret = DATE_ADD(sd,INTERVAL FLOOR(1+RAND()*(sub-1)) SECOND); -> RETURN ret; -> END $$ mysql> delimiter ; # 创建插入数据存储过程 mysql> CREATE DEFINER=`root`@`%` PROCEDURE `add_t_user_memory`(IN n int) -> BEGIN -> DECLARE i INT DEFAULT 1; -> WHILE (i <= n) DO -> INSERT INTO t_user_memory (c_user_id, c_name, c_province_id,c_city_id, create_time) VALUES (uuid(), randStr(20), FLOOR(RAND() * 1000), FLOOR(RAND() * 100), NOW()); -> SET i = i + 1; -> END WHILE; -> END -> $$ Query OK, 0 rows affected (0.01 sec) 复制代码`

* 调用存储过程

` mysql> CALL add_t_user_memory(1000000); ERROR 1114 (HY000): The table 't_user_memory' is full 出现内存已满时，修改 max_heap_table_size 参数的大小，我使用64M内存，插入了22W数据，看情况改，不过这个值不要太大，默认32M或者64M就好，生产环境不要乱尝试 复制代码`

* 从内存表插入普通表

` mysql> INSERT INTO t_user SELECT * FROM t_user_memory; Query OK, 218953 rows affected (1.70 sec) Records: 218953 Duplicates: 0 Warnings: 0 复制代码`

## 方式2： 采用临时表 ##

* 创建临时数据表tmp_table

` CREATE TABLE tmp_table ( id INT, PRIMARY KEY (id) ); 复制代码`

* 用 python或者bash 生成 100w 记录的数据文件（python瞬间就会生成完）

` python(推荐): python -c "for i in range(1, 1+1000000): print(i)" > base.txt 复制代码`

* 导入数据到临时表tmp_table中

` mysql> load data infile '/Users/LJTjintao/temp/base.txt' replace into table tmp_table; Query OK, 1000000 rows affected (2.55 sec) Records: 1000000 Deleted: 0 Skipped: 0 Warnings: 0 千万级数据 20秒插入完成 复制代码`
> 
> 
> 
> 
> 注意： 导入数据时有可能会报错，原因是mysql默认没有开secure_file_priv（
> 这个参数用来限制数据导入和导出操作的效果，例如执行LOAD DATA、SELECT … INTO
> OUTFILE语句和LOAD_FILE()函数。这些操作需要用户具有FILE权限。 ）
> 
> 

> 
> 
> 
> 解决办法：在mysql的配置文件中（my.ini 或者 my.conf）中添加 secure_file_priv =
> /Users/LJTjintao/temp/`, 然后重启mysql 解决
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/21/16ad8b09e79184d6?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/21/16ad8b72bd618e0d?imageView2/0/w/1280/h/960/ignore-error/1)

* 以临时表为基础数据，插入数据到t_user中，100W数据插入需要10.37s

` mysql> INSERT INTO t_user -> SELECT -> id, -> uuid(), -> CONCAT( 'userNickName' , id), -> FLOOR(Rand() * 1000), -> FLOOR(Rand() * 100), -> NOW() -> FROM -> tmp_table; Query OK, 1000000 rows affected (10.37 sec) Records: 1000000 Duplicates: 0 Warnings: 0 复制代码`

* 更新创建时间字段让插入的数据的创建时间更加随机

` UPDATE t_user SET create_time=date_add(create_time, interval FLOOR(1 + (RAND() * 7)) year); Query OK, 1000000 rows affected (5.21 sec) Rows matched: 1000000 Changed: 1000000 Warnings: 0 mysql> UPDATE t_user SET create_time=date_add(create_time, interval FLOOR(1 + (RAND() * 7)) year); Query OK, 1000000 rows affected (4.77 sec) Rows matched: 1000000 Changed: 1000000 Warnings: 0 复制代码` ` mysql> select * from t_user limit 30; +----+--------------------------------------+----------------+---------------+-----------+---------------------+ | id | c_user_id | c_name | c_province_id | c_city_id | create_time | +----+--------------------------------------+----------------+---------------+-----------+---------------------+ | 1 | bf5e227a-7b84-11e9-9d6e-751d319e85c2 | userNickName1 | 84 | 64 | 2015-11-13 21:13:19 | | 2 | bf5e26f8-7b84-11e9-9d6e-751d319e85c2 | userNickName2 | 967 | 90 | 2019-11-13 20:19:33 | | 3 | bf5e2810-7b84-11e9-9d6e-751d319e85c2 | userNickName3 | 623 | 40 | 2014-11-13 20:57:46 | | 4 | bf5e2888-7b84-11e9-9d6e-751d319e85c2 | userNickName4 | 140 | 49 | 2016-11-13 20:50:11 | | 5 | bf5e28f6-7b84-11e9-9d6e-751d319e85c2 | userNickName5 | 47 | 75 | 2016-11-13 21:17:38 | | 6 | bf5e295a-7b84-11e9-9d6e-751d319e85c2 | userNickName6 | 642 | 94 | 2015-11-13 20:57:36 | | 7 | bf5e29be-7b84-11e9-9d6e-751d319e85c2 | userNickName7 | 780 | 7 | 2015-11-13 20:55:07 | | 8 | bf5e2a4a-7b84-11e9-9d6e-751d319e85c2 | userNickName8 | 39 | 96 | 2017-11-13 21:42:46 | | 9 | bf5e2b58-7b84-11e9-9d6e-751d319e85c2 | userNickName9 | 731 | 74 | 2015-11-13 22:48:30 | | 10 | bf5e2bb2-7b84-11e9-9d6e-751d319e85c2 | userNickName10 | 534 | 43 | 2016-11-13 22:54:10 | | 11 | bf5e2c16-7b84-11e9-9d6e-751d319e85c2 | userNickName11 | 572 | 55 | 2018-11-13 20:05:19 | | 12 | bf5e2c70-7b84-11e9-9d6e-751d319e85c2 | userNickName12 | 71 | 68 | 2014-11-13 20:44:04 | | 13 | bf5e2cca-7b84-11e9-9d6e-751d319e85c2 | userNickName13 | 204 | 97 | 2019-11-13 20:24:23 | | 14 | bf5e2d2e-7b84-11e9-9d6e-751d319e85c2 | userNickName14 | 249 | 32 | 2019-11-13 22:49:43 | | 15 | bf5e2d88-7b84-11e9-9d6e-751d319e85c2 | userNickName15 | 900 | 51 | 2019-11-13 20:55:26 | | 16 | bf5e2dec-7b84-11e9-9d6e-751d319e85c2 | userNickName16 | 854 | 74 | 2018-11-13 22:07:58 | | 17 | bf5e2e50-7b84-11e9-9d6e-751d319e85c2 | userNickName17 | 136 | 46 | 2013-11-13 21:53:34 | | 18 | bf5e2eb4-7b84-11e9-9d6e-751d319e85c2 | userNickName18 | 897 | 10 | 2018-11-13 20:03:55 | | 19 | bf5e2f0e-7b84-11e9-9d6e-751d319e85c2 | userNickName19 | 829 | 83 | 2013-11-13 20:38:54 | | 20 | bf5e2f68-7b84-11e9-9d6e-751d319e85c2 | userNickName20 | 683 | 91 | 2019-11-13 20:02:42 | | 21 | bf5e2fcc-7b84-11e9-9d6e-751d319e85c2 | userNickName21 | 511 | 81 | 2013-11-13 21:16:48 | | 22 | bf5e3026-7b84-11e9-9d6e-751d319e85c2 | userNickName22 | 562 | 35 | 2019-11-13 20:15:52 | | 23 | bf5e3080-7b84-11e9-9d6e-751d319e85c2 | userNickName23 | 91 | 39 | 2016-11-13 20:28:59 | | 24 | bf5e30da-7b84-11e9-9d6e-751d319e85c2 | userNickName24 | 677 | 21 | 2016-11-13 21:37:15 | | 25 | bf5e3134-7b84-11e9-9d6e-751d319e85c2 | userNickName25 | 50 | 60 | 2018-11-13 20:39:20 | | 26 | bf5e318e-7b84-11e9-9d6e-751d319e85c2 | userNickName26 | 856 | 47 | 2018-11-13 21:24:53 | | 27 | bf5e31e8-7b84-11e9-9d6e-751d319e85c2 | userNickName27 | 816 | 65 | 2014-11-13 22:06:26 | | 28 | bf5e324c-7b84-11e9-9d6e-751d319e85c2 | userNickName28 | 806 | 7 | 2019-11-13 20:17:30 | | 29 | bf5e32a6-7b84-11e9-9d6e-751d319e85c2 | userNickName29 | 973 | 63 | 2014-11-13 21:08:09 | | 30 | bf5e3300-7b84-11e9-9d6e-751d319e85c2 | userNickName30 | 237 | 29 | 2018-11-13 21:48:17 | +----+--------------------------------------+----------------+---------------+-----------+---------------------+ 30 rows in set (0.01 sec) 复制代码`