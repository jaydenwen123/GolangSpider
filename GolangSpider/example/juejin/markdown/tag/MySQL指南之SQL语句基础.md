# MySQL指南之SQL语句基础 #

> 
> 
> 
> 个人所有文章整理在此篇，将陆续更新收录: [知无涯，行者之路莫言终（我的编程之路）](
> https://juejin.im/post/5c2881c95188252b5627475b )
> 
> 

#### 零、结构化查询语言：SQL(Structured Query Language) ####

` DDL 数据定义语言 管理库，表 DML 数据操作语言 增删改查 DCL 数据控制语言 数据控制，权限访问等 复制代码`

##### 准备活动：创建库和表 #####

` CREATE DATABASE datatype; USE datatype; CREATE TABLE type_number( type CHAR(12), byte TINYINT UNSIGNED, range_singed VARCHAR(20), range_unsinged VARCHAR(20), info VARCHAR(40) ); 复制代码`
> 
> 
> 
> 
> 目前状态：
> 
> 

` mysql> SHOW DATABASES; +--------------------+ | Database | +--------------------+ | datatype | | information_schema | | mycode | | mysql | | performance_schema | | seckill | +--------------------+ mysql> USE datatype; Database changed mysql> SHOW TABLES; +--------------------+ | Tables_in_datatype | +--------------------+ | type_number | +--------------------+ mysql> DESC type_number; +----------------+---------------------+------+-----+---------+-------+ | Field | Type | Null | Key | Default | Extra | +----------------+---------------------+------+-----+---------+-------+ | type | char(12) | YES | | NULL | | | byte | tinyint(3) unsigned | YES | | NULL | | | range_singed | varchar(20) | YES | | NULL | | | range_unsinged | varchar(20) | YES | | NULL | | | info | varchar(40) | YES | | NULL | | +----------------+---------------------+------+-----+---------+-------+ 复制代码`

#### 一、DML 数据库记录操作 ` LEVEL 1` ####

> 
> 
> 
> LEVEL 1先简单掌握一下下面的用法
> 
> 

![MySQL基础操作LEVER1.png](https://user-gold-cdn.xitu.io/2019/3/17/1698b3a5f67773c2?imageView2/0/w/1280/h/960/ignore-error/1)

##### 1、记录的插入操作 #####

> 
> 
> 
> ` INSERT INTO <表名> (属性,...) VALUES (值,...),...;`
> 
> 

` |-- 插入一条数据 INSERT INTO <表名> (属性,...) VALUES (值,...); INSERT INTO type_number( type ,byte,range_singed,range_unsinged,info) VALUES ( 'TINYINT' ,1, '-2⁷ ~ 2⁷-1' , '0 ~ 2⁸-1' , '很小整数' ); |-- 查询所有 SELECT * FROM <表名>; mysql> SELECT * FROM type_number; +---------+------+----------------+----------------+--------------+ | type | byte | range_singed | range_unsinged | info | +---------+------+----------------+----------------+--------------+ | TINYINT | 1 | -2⁷ ~ 2⁷-1 | 0 ~ 2⁸-1 | 很小整数 | +---------+------+----------------+----------------+--------------+ |-- 你也可以一次，插入多条数据 INSERT INTO type_number( type ,byte,range_singed,range_unsinged,info) VALUES ( 'TINYINT' ,1, '-2⁷ ~ 2⁷-1' , '0 ~ 2⁸-1' , '很小整数' ), ( 'SMALLINT' ,2, '-2¹⁶ ~ 2¹⁶-1' , '0 ~ 2¹⁶-1' , '小整数' ), ( 'MEDIUMINT' ,3, '-2²⁴ ~ 2²⁴-1' , '0 ~ 2²⁴-1' , '中等整数' ), ( 'INT' ,4, '-2³² ~ 2³²-1' , '0 ~ 2³²-1' , '标准整数' ), ( 'BIGINT' ,8, '-2⁶⁴ ~ 2⁶⁴-1' , '0 ~ 2⁶⁴-1' , '大整数' ); mysql> SELECT * FROM type_number; +-----------+------+----------------------+----------------+--------------+ | type | byte | range_singed | range_unsinged | info | +-----------+------+----------------------+----------------+--------------+ | TINYINT | 1 | -2⁷ ~ 2⁷-1 | 0 ~ 2⁸-1 | 很小整数 | | TINYINT | 1 | -2⁷ ~ 2⁷-1 | 0 ~ 2⁸-1 | 很小整数 | | SMALLINT | 2 | -2¹⁶ ~ 2¹⁶-1 | 0 ~ 2¹⁶-1 | 小整数 | | MEDIUMINT | 3 | -2²⁴ ~ 2²⁴-1 | 0 ~ 2²⁴-1 | 中等整数 | | INT | 4 | -2³² ~ 2³²-1 | 0 ~ 2³²-1 | 标准整数 | | BIGINT | 8 | -2⁶⁴ ~ 2⁶⁴-1 | 0 ~ 2⁶⁴-1 | 大整数 | +-----------+------+----------------------+----------------+--------------+ 复制代码`

##### 2、记录的更新操作 #####

> 
> 
> 
> ` UPDATE <表名> SET 属性 = 值,... WHERE 条件;`
> 
> 

` UPDATE type_number SET info= '微型整数' WHERE type = 'TINYINT' ; mysql> SELECT * FROM type_number; +-----------+------+----------------------+----------------+--------------+ | type | byte | range_singed | range_unsinged | info | +-----------+------+----------------------+----------------+--------------+ | TINYINT | 1 | -2⁷ ~ 2⁷-1 | 0 ~ 2⁸-1 | 微型整数 | | TINYINT | 1 | -2⁷ ~ 2⁷-1 | 0 ~ 2⁸-1 | 微型整数 | | SMALLINT | 2 | -2¹⁶ ~ 2¹⁶-1 | 0 ~ 2¹⁶-1 | 小整数 | | MEDIUMINT | 3 | -2²⁴ ~ 2²⁴-1 | 0 ~ 2²⁴-1 | 中等整数 | | INT | 4 | -2³² ~ 2³²-1 | 0 ~ 2³²-1 | 标准整数 | | BIGINT | 8 | -2⁶⁴ ~ 2⁶⁴-1 | 0 ~ 2⁶⁴-1 | 大整数 | +-----------+------+----------------------+----------------+--------------+ 复制代码`

##### 3.记录的删除操作 #####

> 
> 
> 
> ` DELETE FROM <表名> WHERE 条件;`
> 
> 

` |--- 删除操作 DELETE FROM type_number WHERE type = 'TINYINT' ; mysql> SELECT * FROM type_number; +-----------+------+----------------------+----------------+--------------+ | type | byte | range_singed | range_unsinged | info | +-----------+------+----------------------+----------------+--------------+ | SMALLINT | 2 | -2¹⁶ ~ 2¹⁶-1 | 0 ~ 2¹⁶-1 | 小整数 | | MEDIUMINT | 3 | -2²⁴ ~ 2²⁴-1 | 0 ~ 2²⁴-1 | 中等整数 | | INT | 4 | -2³² ~ 2³²-1 | 0 ~ 2³²-1 | 标准整数 | | BIGINT | 8 | -2⁶⁴ ~ 2⁶⁴-1 | 0 ~ 2⁶⁴-1 | 大整数 | +-----------+------+----------------------+----------------+--------------+ 复制代码`

##### 4.记录的查询操作 #####

> 
> 
> 
> ` SELECT 属性,... FROM <表名> WHERE 条件;`
> 
> 

` mysql> SELECT type ,range_unsinged FROM type_number WHERE byte>=4; +--------+----------------+ | type | range_unsinged | +--------+----------------+ | INT | 0 ~ 2³²-1 | | BIGINT | 0 ~ 2⁶⁴-1 | +--------+----------------+ 复制代码`

#### 二、图片表pic ` (LEVER 2)` ####

> 
> 
> 
> 这个是用来记录图片信息的表，数据准备过程详见番外篇：
> [[番外]-练习MySQL没素材？来一波字符串操作](
> https://juejin.im/post/5c8d24a4518825493a0b5eb7 )
> 
> 

![MySQL查询LEVER2.png](https://user-gold-cdn.xitu.io/2019/3/17/1698b3a5f6a95969?imageView2/0/w/1280/h/960/ignore-error/1)

##### 1.建表语句 #####

` CREATE TABLE pic( id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, pic_path VARCHAR(120) NOT NULL, pic_length INT UNSIGNED DEFAULT 0, pic_mime TINYINT UNSIGNED, pic_width SMALLINT UNSIGNED, pic_height SMALLINT UNSIGNED ); |--- id 为主键 自增长 |--- pic_path表示名字，不定长度 ，给个VARCHAR 120 吧，差不多够用吧 |--- 图片文件大小不会非常大，给个INT足够了 ， 给个默认值 0 |--- pic_mime 0 表示 image/png 1表示 image/jpeg 给个最小的 |--- pic_width和pic_height也不会非常大，无符号SMALLINT足够 复制代码`

##### 2.查询操作 ` AS` 的作用 #####

` |-- 查询高大于1200像素的记录，使用AS 来 临时更改查询输出的属性名(不会改变实际记录) mysql> SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height>1200; +----------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------+--------+--------+ | 30000X20000.jpg | 30000 | 20000 | | 3000X2000.jpg | 3000 | 2000 | | ecNKedygCmSjTWWF.jpg | 700 | 1352 | | gtQiXnRfkvvTLinw.jpg | 2880 | 2025 | | HXqqASHJETSlvpnc.jpg | 3600 | 2400 | | ndbMXlwKuCpiiVqC.jpg | 1701 | 2268 | | screen.png | 1080 | 1920 | | XQWGrglfjGVuJfzJ.jpg | 1200 | 1696 | +----------------------+--------+--------+ 复制代码`

##### 3.查询是属性可参与运算 #####

` |-- CONCAT函数用于连接字符串 注意：\需要转义 mysql> SELECT CONCAT( 'E:\\SpringBootFiles\\imgs\\' ,pic_path) AS 绝对路径, pic_width * pic_height AS '像素点个数' FROM pic WHERE pic_height>1200; +----------------------------------------------+-----------------+ | 绝对路径 | 像素点个数 | +----------------------------------------------+-----------------+ | E:\SpringBootFiles\imgs\30000X20000.jpg | 600000000 | | E:\SpringBootFiles\imgs\3000X2000.jpg | 6000000 | | E:\SpringBootFiles\imgs\ecNKedygCmSjTWWF.jpg | 946400 | | E:\SpringBootFiles\imgs\gtQiXnRfkvvTLinw.jpg | 5832000 | | E:\SpringBootFiles\imgs\HXqqASHJETSlvpnc.jpg | 8640000 | | E:\SpringBootFiles\imgs\ndbMXlwKuCpiiVqC.jpg | 3857868 | | E:\SpringBootFiles\imgs\screen.png | 2073600 | | E:\SpringBootFiles\imgs\XQWGrglfjGVuJfzJ.jpg | 2035200 | +----------------------------------------------+-----------------+ 复制代码`

##### 4. ` WHERE` 条件的千变万化 #####

![MySQL的WHERE.png](https://user-gold-cdn.xitu.io/2019/3/17/1698b3fe0c93f3b4?imageView2/0/w/1280/h/960/ignore-error/1)

###### 4.1: 条件 ` 与` -- ` AND` 和 ` &&` ######

> 
> 
> 
> 条件必须全部满足
> 
> 

` SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height>1200 AND pic_width > 1500; +----------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------+--------+--------+ | 30000X20000.jpg | 30000 | 20000 | | 3000X2000.jpg | 3000 | 2000 | | gtQiXnRfkvvTLinw.jpg | 2880 | 2025 | | HXqqASHJETSlvpnc.jpg | 3600 | 2400 | | ndbMXlwKuCpiiVqC.jpg | 1701 | 2268 | +----------------------+--------+--------+ |--- AND 效果等于 && SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height>1200 && pic_width > 1500; 复制代码`

###### 4.2: 条件 ` 或` -- ` OR` 和 ` ||` ######

> 
> 
> 
> 条件满足一个即可
> 
> 

` SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height>1200 OR pic_width > 1500; +----------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------+--------+--------+ | 30000X20000.jpg | 30000 | 20000 | | 3000X2000.jpg | 3000 | 2000 | | ecNKedygCmSjTWWF.jpg | 700 | 1352 | | gtQiXnRfkvvTLinw.jpg | 2880 | 2025 | | HXqqASHJETSlvpnc.jpg | 3600 | 2400 | | ndbMXlwKuCpiiVqC.jpg | 1701 | 2268 | | screen.png | 1080 | 1920 | | XQWGrglfjGVuJfzJ.jpg | 1200 | 1696 | +----------------------+--------+--------+ |--- OR 效果等于 || SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height>1200 || pic_width > 1500; 复制代码`

###### 4.3: 条件 ` 非` -- ` NOT` 和 ` !` ######

> 
> 
> 
> 对条件取反
> 
> 

` SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE NOT pic_height < 1200; +----------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------+--------+--------+ | 30000X20000.jpg | 30000 | 20000 | | 3000X2000.jpg | 3000 | 2000 | | ecNKedygCmSjTWWF.jpg | 700 | 1352 | | gtQiXnRfkvvTLinw.jpg | 2880 | 2025 | | HXqqASHJETSlvpnc.jpg | 3600 | 2400 | | ndbMXlwKuCpiiVqC.jpg | 1701 | 2268 | | screen.png | 1080 | 1920 | | XQWGrglfjGVuJfzJ.jpg | 1200 | 1696 | +----------------------+--------+--------+ 复制代码`

###### 4.4: 散点匹配 ` IN(v1,v2,v3,...)` ######

> 
> 
> 
> 符合v1,v2,v3,...之一可匹配
> 
> 

` SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height IN (1696,2268); +----------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------+--------+--------+ | ndbMXlwKuCpiiVqC.jpg | 1701 | 2268 | | XQWGrglfjGVuJfzJ.jpg | 1200 | 1696 | +----------------------+--------+--------+ 复制代码`

###### 4.5: 区间匹配 ` BETWEEN v1 AND v2` ######

> 
> 
> 
> v1,v2之间可匹配
> 
> 

` SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_height BETWEEN 1696 AND 2268; +----------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------+--------+--------+ | 3000X2000.jpg | 3000 | 2000 | | gtQiXnRfkvvTLinw.jpg | 2880 | 2025 | | ndbMXlwKuCpiiVqC.jpg | 1701 | 2268 | | screen.png | 1080 | 1920 | | XQWGrglfjGVuJfzJ.jpg | 1200 | 1696 | +----------------------+--------+--------+ 复制代码`

###### 4.6：模糊查询： ` LIKE` ######

> 
> 
> 
> ` '%'匹配任意多个字符,'_'匹配任意单个字符`
> 
> 

` mysql> SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_path LIKE 'androi%' ; +----------------------------------------------+--------+--------+ | 路径 | 宽/px | 高/px | +----------------------------------------------+--------+--------+ | android\008525ebc2b7d434070e74c00841a30f.png | 544 | 544 | | android\054d98e2d96dc42d9b2b036126fccf49.png | 544 | 544 | | android\05baf2d03651d1110d7a403f14aee877.png | 544 | 544 | | android\0655e07d6717847489 cd 222c9c9e0b1d.png | 500 | 500 | | android\079c4cb46c95b2365b5bc5150e7d5213.png | 544 | 544 | | android\07a4dc9b4b207cb420a71cbf941ad45a.png | 544 | 544 | | android\07abb7972a5638b53afa3b5eb98b19c1.png | 500 | 500 | ...... mysql> SELECT pic_path AS 路径 , pic_width AS '宽/px' , pic_height AS '高/px' FROM pic WHERE pic_path LIKE 'p_em%' ; +--------------------------------------------+--------+--------+ | 路径 | 宽/px | 高/px | +--------------------------------------------+--------+--------+ | poem\世界·绽放.jpg | 1148 | 712 | | poem\我爱你，是火山岩的缄默.jpg | 690 | 397 | | poem\枝·你是树的狂舞.jpg | 500 | 333 | | poem\海与鹿王.jpg | 799 | 499 | | poem\游梦人·诗的诞生.jpg | 800 | 444 | | poem\珊瑚墓地.jpg | 1104 | 719 | +--------------------------------------------+--------+--------+ 复制代码`

###### 4.7:比较符号 ` = != < > <= >=` ######

> 
> 
> 
> 小学生都知道的，就不废话了，查看一下小于10Kb的图片
> 
> 

` mysql> SELECT pic_path AS 路径 , pic_length AS '大小/byte' FROM pic WHERE pic_length < 10*1024; +----------------------------------------------+-------------+ | 路径 | 大小/byte | +----------------------------------------------+-------------+ | 30X20.jpg | 10158 | | android\613f2b8f0eaa8f63bedce9781527c9ab.png | 4001 | | android\94b5c41232f9761403890c09c2b1aae3.png | 4001 | | android\d3fd676f224f0734beb48d0c0d2f4e66.png | 4001 | | udp发送与接收消息_控制台.png | 9184 | +----------------------------------------------+-------------+ 复制代码`

##### 5. ` GROUP BY` 分组查询 #####

> 
> 
> 
> 会先排序，再列出
> 
> 

` |--- GROUP BY SELECT pic_mime AS "类型" , avg(pic_length) AS '平均大小/byte' , count(pic_length) AS '总数量/个' , min(pic_length) AS '最小值/byte' , max(pic_length) AS '最大值/byte' , sum(pic_length) AS '总和/byte' FROM pic GROUP BY pic_mime; +--------+-------------------+---------------+----------------+----------------+-------------+ | 类型 | 平均大小/byte | 总数量/个 | 最小值/byte | 最大值/byte | 总和/byte | +--------+-------------------+---------------+----------------+----------------+-------------+ | 0 | 141518.8734 | 229 | 4001 | 829338 | 32407822 | | 1 | 2133272.8000 | 60 | 10158 | 116342886 | 127996368 | +--------+-------------------+---------------+----------------+----------------+-------------+ 复制代码`

##### 6.结果集筛选： ` HAVING` #####

> 
> 
> 
> 现在查询宽高比在1.1和1.3之间的图片
> 
> 

` |-- 如果用WHERE 来查询 感觉有点不优雅 SELECT pic_path AS 路径 , pic_width/pic_height AS '宽高比' FROM pic WHERE pic_width/pic_height > 1.1 && pic_width/pic_height<1.3; +------------------------------------------------------------------+-----------+ | 路径 | 宽高比 | +------------------------------------------------------------------+-----------+ | dQXbnTRjUdNxhiyl.jpg | 1.2308 | | JsXHWmKqOlziKmeA.jpg | 1.2600 | | logo\android\Android原生绘图之让你了解View的运动.png | 1.2884 | | 洛天依.jpg | 1.1990 | +------------------------------------------------------------------+-----------+ |-- AS 相当于将列取了变量，对结果集再进行筛选用HAVING,用WHERE则报错，找不到列 SELECT pic_path AS 路径 , pic_width/pic_height AS ratio FROM pic HAVING ratio > 1.1 && ratio <1.3; +------------------------------------------------------------------+--------+ | 路径 | ratio | +------------------------------------------------------------------+--------+ | dQXbnTRjUdNxhiyl.jpg | 1.2308 | | JsXHWmKqOlziKmeA.jpg | 1.2600 | | logo\android\Android原生绘图之让你了解View的运动.png | 1.2884 | | 洛天依.jpg | 1.1990 | +------------------------------------------------------------------+--------+ 复制代码`

##### 7.结果排序： ` ORDER BY` #####

> 
> 
> 
> 按照ratio将序排列
> 
> 

` SELECT pic_path AS 路径 , pic_width/pic_height AS ratio FROM pic HAVING ratio > 1.1 && ratio <1.3; ORDER BY ratio DESC +------------------------------------------------------------------+--------+ | 路径 | ratio | +------------------------------------------------------------------+--------+ | dQXbnTRjUdNxhiyl.jpg | 1.2308 | | JsXHWmKqOlziKmeA.jpg | 1.2600 | | logo\android\Android原生绘图之让你了解View的运动.png | 1.2884 | | 洛天依.jpg | 1.1990 | +------------------------------------------------------------------+--------+ 复制代码`

##### 8.控制条目数： ` LIMIT` #####

` |-- 偏移一条，取两条 SELECT pic_path AS 路径 , pic_width/pic_height AS ratio FROM pic HAVING ratio > 1.1 && ratio <1.3 ORDER BY ratio DESC LIMIT 1,2; +----------------------+--------+ | 路径 | ratio | +----------------------+--------+ | JsXHWmKqOlziKmeA.jpg | 1.2600 | | dQXbnTRjUdNxhiyl.jpg | 1.2308 | +----------------------+--------+ 复制代码`

#### 三、子查询 ` (LEVER 3)` ####

##### 1.查询大于平均尺寸的图片 -- ` WHERE` #####

` |--- 出现在其他SQL语句内的SELECT语句 |--- 子查询必须在()内 |--- 增删改查都可以进行子查询,返回：标量，行，列或子查询 复制代码` ` |-- 1-1：查出图片平均大小 SELECT ROUND(AVG(pic_length),2) AS '平均大小' FROM pic; +--------------+ | 平均大小 | +--------------+ | 555031.80 | +--------------+ 1 row in set (0.00 sec) |-- 1-2：在用WHERE 筛选 SELECT pic_path AS 路径 , pic_length AS '大小/byte' FROM pic WHERE pic_length > 555031.80; +----------------------------------------------+-------------+ | 路径 | 大小/byte | +----------------------------------------------+-------------+ | 30000X20000.jpg | 116342886 | | 3000X2000.jpg | 3404969 | | android\12284e5f7197d8be737fa967c8b00fbe.png | 829338 | | android\594665add495ac9da8b6bbee1c63f1b8.png | 598974 | | android\7cc97458727e23f7d161b8a1a7c6b453.png | 559420 | | android\cbb1524f5ab4266698f3a6 fc 2992ccae.png | 829338 | | android\d52539b1b508a594d1f2865037ff50c5.png | 598974 | | android\f07ddfe5a103e4a024e14e2569f1d70e.png | 829338 | | android\f0d1e7713d5557a8f9c74c9904843e09.png | 559420 | | bg.png | 688207 | | gtQiXnRfkvvTLinw.jpg | 771187 | | poem\珊瑚墓地.jpg | 984472 | | XoazFNMQROveEPQn.jpg | 795364 | +----------------------------------------------+-------------+ |--- 也就是将一个语句包在WHERE 条件里 SELECT pic_path AS 路径 , pic_length AS '大小/byte' FROM pic WHERE pic_length > ( SELECT ROUND(AVG(pic_length),2) FROM pic ); 复制代码`

##### 2.查出每种类型的最新插入的图片 -- ` WHERE` #####

` SELECT pic_path AS 路径 , pic_mime AS 类型 FROM pic WHERE id IN ( SELECT max(id) FROM pic GROUP BY pic_mime ); +------------------+--------+ | 路径 | 类型 | +------------------+--------+ | 洛天依.jpg | 1 | | 虚拟机栈.png | 0 | +------------------+--------+ 复制代码`

##### 3.FROM子查询 -- ` FROM` #####

` SELECT id, pic_path AS 路径 , pic_length AS '大小/byte' FROM pic WHERE id>=10&&id<=15 ORDER BY pic_length DESC; +----+----------------------------------------------+-------------+ | id | 路径 | 大小/byte | +----+----------------------------------------------+-------------+ | 15 | android\0f3bf63796ac370a08ee97b056b0587b.png | 178849 | | 14 | android\0951ef0be68f0c498ca34ffcd7 fc 7faa.png | 175842 | | 11 | android\079c4cb46c95b2365b5bc5150e7d5213.png | 86996 | | 10 | android\0655e07d6717847489 cd 222c9c9e0b1d.png | 53764 | | 12 | android\07a4dc9b4b207cb420a71cbf941ad45a.png | 46270 | | 13 | android\07abb7972a5638b53afa3b5eb98b19c1.png | 43360 | +----+----------------------------------------------+-------------+ |--- 将查询结果当做一张表，再查询操作 SELECT id,路径 FROM ( SELECT id, pic_path AS 路径 , pic_length AS '大小/byte' FROM pic WHERE id>=10&&id<=15 ORDER BY pic_length DESC ) AS result WHERE `大小/byte` < 59999; +----+----------------------------------------------+ | id | 路径 | +----+----------------------------------------------+ | 10 | android\0655e07d6717847489 cd 222c9c9e0b1d.png | | 12 | android\07a4dc9b4b207cb420a71cbf941ad45a.png | | 13 | android\07abb7972a5638b53afa3b5eb98b19c1.png | +----+----------------------------------------------+ 复制代码`

#### 四、连接查询 ####

##### 0.创建关联表 #####

> 
> 
> 
> 首先连接查询要多张表，现在建一个 ` mime_type` 的表
> 
> 

` |--- 建表 CREATE TABLE mime_type( mime_id SMALLINT UNSIGNED PRIMARY KEY, mime_info CHAR(24) ); |--- 插入数据 INSERT INTO mime_type(mime_id,mime_info) VALUES (0, 'image/png' ), (1, 'image/jpeg' ), (2, 'image/svg+xml' ), (3, 'video/mp4' ), (4, 'text/plain' ); |--- 效果 mysql> select * from mime_type; +---------+---------------+ | mime_id | mime_info | +---------+---------------+ | 0 | image/png | | 1 | image/jpeg | | 2 | image/svg+xml | | 3 | video/mp4 | | 4 | text/plain | +---------+---------------+ |-- 为了说明问题，pic表添加一条测试数据：pic_mime = 8 也就是 mime_type表找不到时 INSERT INTO pic(pic_path,pic_length,pic_mime,pic_width,pic_height) VALUES( 'test.jpg' ,100,8,300,200); 复制代码`

##### 1.内连接查询 ` INNER JOIN` #####

> 
> 
> 
> ` SELECT 待查属性 FROM 表1 INNER JOIN 表2 ON 条件 WHERE 条件`
> 
> 

` SELECT id, pic_path AS 路径 , mime_type.mime_info AS 类型 , pic_length FROM pic INNER JOIN mime_type ON pic.pic_mime = mime_type.mime_id ORDER BY id DESC LIMIT 4; +-----+------------------+------------+------------+ | id | 路径 | 类型 | pic_length | +-----+------------------+------------+------------+ | 289 | 虚拟机栈.png | image/png | 63723 | | 288 | 统一返回.png | image/png | 29485 | | 287 | 洛天依.jpg | image/jpeg | 42117 | | 286 | 标记整理.png | image/png | 29288 | +-----+------------------+------------+------------+ 复制代码`

##### 2.左连接查询 : ` LEFT JOIN` #####

> 
> 
> 
> 保持左表的记录完整性，右表查不到就摆 NULL
> 
> 

` SELECT id, pic_path AS 路径 , mime_type.mime_info AS 类型 , pic_length FROM pic LEFT JOIN mime_type ON pic.pic_mime = mime_type.mime_id ORDER BY id DESC LIMIT 4; +-----+------------------+------------+------------+ | id | 路径 | 类型 | pic_length | +-----+------------------+------------+------------+ | 290 | test.jpg | NULL | 100 | | 289 | 虚拟机栈.png | image/png | 63723 | | 288 | 统一返回.png | image/png | 29485 | | 287 | 洛天依.jpg | image/jpeg | 42117 | +-----+------------------+------------+------------+ 复制代码`

##### 3. 右(外)连接查询 : ` RIGHT JOIN` #####

> 
> 
> 
> 保持右表的记录完整性，左表查不到就摆 NULL
> 
> 

` SELECT id, pic_path AS 路径 , mime_type.mime_info AS 类型 , pic_length FROM pic RIGHT JOIN mime_type ON pic.pic_mime = mime_type.mime_id ORDER BY id LIMIT 8; +------+--------------------------------------+---------------+------------+ | id | 路径 | 类型 | pic_length | +------+--------------------------------------+---------------+------------+ | NULL | NULL | text/plain | NULL | | NULL | NULL | video/mp4 | NULL | | NULL | NULL | image/svg+xml | NULL | | 1 | 30000X20000.jpg | image/jpeg | 116342886 | | 2 | 3000X2000.jpg | image/jpeg | 3404969 | | 3 | 300X200.jpg | image/jpeg | 99097 | | 4 | 30X20.jpg | image/jpeg | 10158 | | 5 | 6dc9e8455c47d964e1a8a4ef04cf9477.jpg | image/jpeg | 236254 | +------+--------------------------------------+---------------+------------+ 复制代码`

##### 4. 全(外)连接 (伪): ` 使用UNION` #####

> 
> 
> 
> MySQL不支持全外连接，所以只能采取关键字UNION来联合左、右连接的方法 UNION : 将若干条sql的查询结果集合并成一个。 ` UNION
> ALL` 不会覆盖相同结果
> 
> 

` SELECT id, pic_path AS 路径 , mime_type.mime_info AS 类型 , pic_length FROM pic LEFT JOIN mime_type ON pic.pic_mime = mime_type.mime_id UNION( SELECT id, pic_path AS 路径 , mime_type.mime_info AS 类型 , pic_length FROM pic RIGHT JOIN mime_type ON pic.pic_mime = mime_type.mime_id ) ORDER BY id DESC; +------+------------------------------------------------------------------------------------+---------------+------------+ | id | 路径 | 类型 | pic_length | +------+------------------------------------------------------------------------------------+---------------+------------+ | 290 | test.jpg | NULL | 100 | | 289 | 虚拟机栈.png | image/png | 63723 | | 288 | 统一返回.png | image/png | 29485 | | 287 | 洛天依.jpg | image/jpeg | 42117 | ... | 3 | 300X200.jpg | image/jpeg | 99097 | | 2 | 3000X2000.jpg | image/jpeg | 3404969 | | 1 | 30000X20000.jpg | image/jpeg | 116342886 | | NULL | NULL | text/plain | NULL | | NULL | NULL | video/mp4 | NULL | | NULL | NULL | image/svg+xml | NULL | +------+------------------------------------------------------------------------------------+---------------+------------+ 复制代码`

##### 5. UNION小测试 #####

` CREATE TABLE a( id CHAR(4), num INT ); INSERT INTO a(id,num) VALUES ( 'a' ,4),( 'b' ,6),( 'c' ,2),( 'd' ,8); CREATE TABLE b( id CHAR(4), num INT ); INSERT INTO b(id,num) VALUES ( 'b' ,8),( 'c' ,7),( 'd' ,3),( 'e' ,18); mysql> SELECT * FROM a; mysql> SELECT * FROM b; +------+------+ +------+------+ | id | num | | id | num | +------+------+ +------+------+ | a | 4 | | b | 8 | | b | 6 | | c | 7 | | c | 2 | | d | 3 | | d | 8 | | e | 18 | +------+------+ +------+------+ SELECT id,sum(num) FROM (SELECT * FROM a UNION ALL SELECT * FROM b) as temp GROUP BY id; +------+----------+ | id | sum(num) | +------+----------+ | a | 4 | | b | 14 | | c | 9 | | d | 11 | | e | 18 | +------+----------+ 复制代码`

#### 六、DDL 建库/表 ####

##### 1、关于操作数据库 #####

` SHOW DATABASES; # 显示所有的数据库 SHOW CREATE DATABASE <数据库名> # 查看数据库创建信息 USE <数据库名>; # 使用数据库 CREATE DATABASE <数据库名> [CHARACTER SET <字符集>]; # 创建一个将的数据库指定字符集 ALTER DATABASE <数据库名> CHARACTER SET <字符集>; # 修改数据库字符集 DROP DATABASE <数据库名>; # 传说中的删库跑路 SELECT DATABASE(); # 查看当前选中的数据库 复制代码`

##### 2.显示数据库信息 #####

` SHOW TABLES; # 展示当前数据库中的表 SHOW TABLES FROM mysql # 展示指定数据库中的表 DESC <表名>; # 查看表结构 SHOW COLUMNS FROM <表名>; # 查看表结构 复制代码`

##### 3.创建表 #####

` |-- UNSIGNED 无符号 AUTO_INCREMENT 自增长 |-- ZEROFILL 前面自动填 0 , 默认 UNSIGNED CREATE TABLE create_test( id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, code TINYINT(5) ZEROFILL DEFAULT 0 ); INSERT INTO create_test(code) VALUES (5); INSERT INTO create_test VALUES (); #默认值测试 mysql> SELECT * FROM create_test; +----+-------+ | id | code | +----+-------+ | 1 | 00005 | | 2 | 00000 | +----+-------+ 复制代码`

##### 4.为表增加属性 #####

> 
> 
> 
> ` ALTER TABLE <表名> ADD 属性信息 [AFTER 属性] ;`
> 
> 

` |-- 看一下当前表结构 mysql> DESC create_test; +-------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +-------+------------------------------+------+-----+---------+----------------+ | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | +-------+------------------------------+------+-----+---------+----------------+ mysql> ALTER TABLE create_test ADD age SMALLINT UNSIGNED NOT NULL; mysql> DESC create_test; +-------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +-------+------------------------------+------+-----+---------+----------------+ | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | +-------+------------------------------+------+-----+---------+----------------+ |-- AFTER可将属性排在指定属性之后(强迫症专用) |-- ALTER TABLE create_test ADD password VARCHAR(32) AFTER id; mysql> ALTER TABLE create_test ADD password VARCHAR(32) AFTER id; mysql> DESC create_test; +----------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +----------+------------------------------+------+-----+---------+----------------+ | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | password | varchar(32) | YES | | NULL | | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | +----------+------------------------------+------+-----+---------+----------------+ |-- 一次添加多个属性 ALTER TABLE create_test ADD (aaa VARCHAR(32), bbb VARCHAR(32),ccc VARCHAR(32)); mysql> DESC create_test; +----------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +----------+------------------------------+------+-----+---------+----------------+ | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | password | varchar(32) | YES | | NULL | | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | | aaa | varchar(32) | YES | | NULL | | | bbb | varchar(32) | YES | | NULL | | | ccc | varchar(32) | YES | | NULL | | +----------+------------------------------+------+-----+---------+----------------+ 复制代码`

##### 5.为表删除属性 #####

> 
> 
> 
> ` ALTER TABLE <表名> DROP 属性`
> 
> 

` ALTER TABLE create_test DROP aaa,DROP bbb,DROP ccc; mysql> DESC create_test; +----------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +----------+------------------------------+------+-----+---------+----------------+ | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | password | varchar(32) | YES | | NULL | | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | +----------+------------------------------+------+-----+---------+----------------+ 复制代码`

##### 6.修改属性的类型 #####

> 
> 
> 
> ` ALTER TABLE <表名> MODIFY 属性 属性类型 [FIRST];`
> 
> 

` |-- 把password改成VARCHAR(40) ALTER TABLE create_test MODIFY password VARCHAR(40); mysql> DESC create_test; +----------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +----------+------------------------------+------+-----+---------+----------------+ | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | password | varchar(40) | YES | | NULL | | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | +----------+------------------------------+------+-----+---------+----------------+ |-- 将某个属性移到最顶 ALTER TABLE create_test MODIFY password VARCHAR(40) FIRST; +----------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +----------+------------------------------+------+-----+---------+----------------+ | password | varchar(40) | YES | | NULL | | | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | +----------+------------------------------+------+-----+---------+----------------+ 复制代码`

###### 7.修改表的属性名 ######

> 
> 
> 
> ` ALTER TABLE <表名> CHANGE 原属性 新属性 新属性类型;`
> 
> 

` mysql> ALTER TABLE create_test CHANGE password pw varchar(40); mysql> DESC create_test; +-------+------------------------------+------+-----+---------+----------------+ | Field | Type | Null | Key | Default | Extra | +-------+------------------------------+------+-----+---------+----------------+ | pw | varchar(40) | YES | | NULL | | | id | int(10) unsigned | NO | PRI | NULL | auto_increment | | code | tinyint(5) unsigned zerofill | YES | | 00000 | | | age | smallint(5) unsigned | NO | | NULL | | +-------+------------------------------+------+-----+---------+----------------+ 复制代码`

##### 8.修改表名 #####

> 
> 
> 
> 方式一： ` ALTER TABLE 旧表名 RENAME 新表名;` 方式二： ` RENAME TABLE 旧表名 TO 新表名;`
> 
> 

` ALTER TABLE create_test RENAME 阿姆斯特朗回旋加速喷气式阿姆斯特朗炮; mysql> SHOW TABLES; +--------------------------------------------------------+ | Tables_in_datatype | +--------------------------------------------------------+ | 阿姆斯特朗回旋加速喷气式阿姆斯特朗炮 | | a | | b | | mime_type | | pic | | type_number | +--------------------------------------------------------+ RENAME TABLE 阿姆斯特朗回旋加速喷气式阿姆斯特朗炮 TO toly; mysql> SHOW TABLES; +--------------------+ | Tables_in_datatype | +--------------------+ | a | | b | | mime_type | | pic | | toly | | type_number | +--------------------+ 复制代码`
> 
> 
> 
> 
> SQL 的基础就这样 , 下篇见
> 
>