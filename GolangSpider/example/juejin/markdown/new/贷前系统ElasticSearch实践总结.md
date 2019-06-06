# 贷前系统ElasticSearch实践总结 #

贷前系统负责从进件到放款前所有业务流程的实现，其中涉及一些数据量较大、条件多样且复杂的综合查询，引入ElasticSearch主要是为了提高查询效率，并希望基于ElasticSearch快速实现一个简易的数据仓库，提供一些OLAP相关功能。本文将介绍贷前系统ElasticSearch的实践经验。

## 一、索引 ##

描述：为快速定位数据而设计的某种数据结构。

索引好比是一本书前面的目录，能加快数据库的查询速度。了解索引的构造及使用，对理解ES的工作模式有非常大的帮助。

常用索引：

* 

位图索引

* 

哈希索引

* 

BTREE索引

* 

倒排索引

### 1.1 位图索引（BitMap） ###

位图索引适用于字段值为可枚举的有限个数值的情况。

位图索引使用二进制的数字串（bitMap）标识数据是否存在，1标识当前位置（序号）存在数据，0则表示当前位置没有数据。

下图1 为用户表，存储了性别和婚姻状况两个字段；

图2中分别为性别和婚姻状态建立了两个位图索引。

例如：性别->男 对应索引为：101110011，表示第1、3、4、5、8、9个用户为男性。其他属性以此类推。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac225517fe8f?imageView2/0/w/1280/h/960/ignore-error/1)

使用位图索引查询：

* 

男性 并且已婚 的记录 = 101110011 & 11010010 = 100100010，即第1、4、8个用户为已婚男性。

* 

女性 或者未婚的记录 = 010001100 | 001010100 = 011011100, 即第2、3、5、6、7个用户为女性或者未婚。

### 1.2 哈希索引 ###

顾名思义，是指使用某种哈希函数实现key->value 映射的索引结构。

哈希索引适用于等值检索，通过一次哈希计算即可定位数据的位置。

下图3 展示了哈希索引的结构，与JAVA中HashMap的实现类似，是用冲突表的方式解决哈希冲突的。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac2256be4046?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.3 BTREE索引 ###

BTREE索引是关系型数据库最常用的索引结构，方便了数据的查询操作。

BTREE: 有序平衡N阶树, 每个节点有N个键值和N+1个指针, 指向N+1个子节点。

一棵BTREE的简单结构如下图4所示，为一棵2层的3叉树，有7条数据：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac2256cbb1a1?imageView2/0/w/1280/h/960/ignore-error/1)

以Mysql最常用的InnoDB引擎为例，描述下BTREE索引的应用。

Innodb下的表都是以索引组织表形式存储的，也就是整个数据表的存储都是B+tree结构的，如图5所示。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac2256ca3f81?imageView2/0/w/1280/h/960/ignore-error/1)

主键索引为图5的左半部分（如果没有显式定义自主主键，就用不为空的唯一索引来做聚簇索引，如果也没有唯一索引，则innodb内部会自动生成6字节的隐藏主键来做聚簇索引），叶子节点存储了完整的数据行信息（以主键 + row_data形式存储）。

二级索引也是以B+tree的形式进行存储，图5右半部分，与主键不同的是二级索引的叶子节点存储的不是行数据，而是索引键值和对应的主键值，由此可以推断出，二级索引查询多了一步查找数据主键的过程。

维护一颗有序平衡N叉树，比较复杂的就是当插入节点时节点位置的调整，尤其是插入的节点是随机无序的情况；而插入有序的节点，节点的调整只发生了整个树的局部，影响范围较小，效率较高。

可以参考红黑树的节点的插入算法：

[en.wikipedia.org/wiki/Red–bl…]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FRed%25E2%2580%2593black_tree )

因此如果innodb表有自增主键，则数据写入是有序写入的，效率会很高；如果innodb表没有自增的主键，插入随机的主键值，将导致B+tree的大量的变动操作，效率较低。这也是为什么会建议innodb表要有无业务意义的自增主键，可以大大提高数据插入效率。

注：

* 

Mysql Innodb使用自增主键的插入效率高。

* 

使用类似Snowflake的ID生成算法，生成的ID是趋势递增的，插入效率也比较高。

### 1.4 倒排索引（反向索引） ###

倒排索引也叫反向索引，可以相对于正向索引进行比较理解。

正向索引反映了一篇文档与文档中关键词之间的对应关系；给定文档标识，可以获取当前文档的关键词、词频以及该词在文档中出现的位置信息，如图6 所示，左侧是文档，右侧是索引。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac225707bdfc?imageView2/0/w/1280/h/960/ignore-error/1)

反向索引则是指某关键词和该词所在的文档之间的对应关系；给定了关键词标识，可以获取关键词所在的所有文档列表，同时包含词频、位置等信息，如图7所示。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac22571208fc?imageView2/0/w/1280/h/960/ignore-error/1)

反向索引（倒排索引）的单词的集合和文档的集合就组成了如图8所示的”单词-文档矩阵“，打钩的单元格表示存在该单词和文档的映射关系。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac227af84cb6?imageView2/0/w/1280/h/960/ignore-error/1)

倒排索引的存储结构可以参考图9。其中词典是存放的内存里的，词典就是整个文档集合中解析出的所有单词的列表集合；每个单词又指向了其对应的倒排列表，倒排列表的集合组成了倒排文件，倒排文件存放在磁盘上，其中的倒排列表内记录了对应单词在文档中信息，即前面提到的词频、位置等信息。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac227a547b50?imageView2/0/w/1280/h/960/ignore-error/1)

下面以一个具体的例子来描述下，如何从一个文档集合中生成倒排索引。

如图10，共存在5个文档，第一列为文档编号，第二列为文档的文本内容。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac227a1be3d1?imageView2/0/w/1280/h/960/ignore-error/1)

将上述文档集合进行分词解析，其中发现的10个单词为：[谷歌，地图，之父，跳槽，Facebook，加盟，创始人，拉斯，离开，与]，以第一个单词”谷歌“为例：首先为其赋予一个唯一标识 ”单词ID“， 值为1，统计出文档频率为5，即5个文档都有出现，除了在第3个文档中出现2次外，其余文档都出现一次，于是就有了图11所示的倒排索引。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac227a0528e7?imageView2/0/w/1280/h/960/ignore-error/1)

#### 1.4.1 单词词典查询优化 ####

对于一个规模很大的文档集合来说，可能包含几十万甚至上百万的不同单词，能否快速定位某个单词，这直接影响搜索时的响应速度，其中的优化方案就是为单词词典建立索引，有以下几种方案可供参考：

* 词典Hash索引

Hash索引简单直接，查询某个单词，通过计算哈希函数，如果哈希表命中则表示存在该数据，否则直接返回空就可以；适合于完全匹配，等值查询。如图12，相同hash值的单词会放在一个冲突表中。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac227ce8e1b5?imageView2/0/w/1280/h/960/ignore-error/1)

* 词典BTREE索引

类似于Innodb的二级索引，将单词按照一定的规则排序，生成一个BTree索引，数据节点为指向倒排索引的指针。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac2284686ee7?imageView2/0/w/1280/h/960/ignore-error/1)

* 二分查找

同样将单词按照一定的规则排序，建立一个有序单词数组，在查找时使用二分查找法；二分查找法可以映射为一个有序平衡二叉树，如图14这样的结构。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac22ab21bfe3?imageView2/0/w/1280/h/960/ignore-error/1)

* FST（Finite State Transducers ）实现

FST为一种有限状态转移机，FST有两个优点：1）空间占用小。通过对词典中单词前缀和后缀的重复利用，压缩了存储空间；2）查询速度快。O(len(str))的查询时间复杂度。

以插入“cat”、 “deep”、 “do”、 “dog” 、“dogs”这5个单词为例构建FST（注：必须已排序）。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac22a7757416?imageView2/0/w/1280/h/960/ignore-error/1)

如图15 最终我们得到了如上一个有向无环图。利用该结构可以很方便的进行查询，如给定一个词 “dog”，我们可以通过上述结构很方便的查询存不存在，甚至我们在构建过程中可以将单词与某一数字、单词进行关联，从而实现key-value的映射。

当然还有其他的优化方式，如使用Skip List、Trie、Double Array Trie等结构进行优化，不再一一赘述。

## 二、ElasticSearch使用心得 ##

下面结合贷前系统具体的使用案例，介绍ES的一些心得总结。

### 2.1 概况 ###

目前使用的ES版本：5.6

官网地址： [www.elastic.co/products/el…]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fproducts%2Felasticsearch )

ES一句话介绍：The Heart of the Elastic Stack（摘自官网）

ES的一些关键信息：

* 

2010年2月首次发布

* 

Elasticsearch Store, Search, and Analyze

* 

丰富的Restful接口

### 2.2 基本概念 ###

* 索引（index）

ES的索引，也就是Index，和前面提到的索引并不是一个概念，这里是指所有文档的集合，可以类比为RDB中的一个数据库。

* 文档（document）

即写入ES的一条记录，一般是JSON形式的。

* 映射（Mapping）

文档数据结构的元数据描述，一般是JSON schema形式，可动态生成或提前预定义。

* 类型（type）

由于理解和使用上的错误，type已不推荐使用，目前我们使用的ES中一个索引只建立了一个默认type。

* 节点

一个ES的服务实例，称为一个服务节点。为了实现数据的安全可靠，并且提高数据的查询性能，ES一般采用集群模式进行部署。

* 集群

多个ES节点相互通信，共同分担数据的存储及查询，这样就构成了一个集群。

* 分片

分片主要是为解决大量数据的存储，将数据分割为若干部分，分片一般是均匀分布在各ES节点上的。需要注意：分片数量无法修改。

* 副本

分片数据的一份完全的复制，一般一个分片会有一个副本，副本可以提供数据查询，集群环境下可以提高查询性能。

### 2.3 安装部署 ###

* 

JDK版本： JDK1.8

* 

安装过程比较简单，可参考官网：下载安装包 -> 解压 -> 运行

* 

安装过程遇到的坑:

ES启动占用的系统资源比较多，需要调整诸如文件句柄数、线程数、内存等系统参数，可参考下面的文档。

[www.cnblogs.com/sloveling/p…]( https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2Fsloveling%2Fp%2Felasticsearch.html )

### 2.4 实例讲解 ###

下面以一些具体的操作介绍ES的使用：

#### 2.4.1 初始化索引 ####

初始化索引，主要是在ES中新建一个索引并初始化一些参数，包括索引名、文档映射（Mapping）、索引别名、分片数（默认：5）、副本数（默认：1）等，其中分片数和副本数在数据量不大的情况下直接使用默认值即可，无需配置。

下面举两个初始化索引的方式，一个使用基于Dynamic Template（动态模板） 的Dynamic Mapping（动态映射），一个使用显式预定义映射。

1） 动态模板 (Dynamic Template)

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" >curl -X PUT http://ip:9200/loan_idx -H 'content-type: application/json' <br> -d '{"mappings":{ "order_info":{ "dynamic_date_formats":["yyyy-MM-dd HH:mm:ss||yyyy-MM-dd],<br> "dynamic_templates":[<br> {"orderId2":{<br> "match_mapping_type":"string",<br> "match_pattern":"regex",<br> "match":"^orderId$",<br> "mapping":{<br> "type":"long"<br> }<br> }<br> },<br> {"strings_as_keywords":{<br> "match_mapping_type":"string",<br> "mapping":{<br> "type":"keyword",<br> "norms":false<br> }<br> }<br> }<br> ]<br> }<br>},<br>"aliases":{<br> "loan_alias":{}<br>}}' <br></span></p> 复制代码`

上面的JSON串就是我们用到的动态模板，其中定义了日期格式：dynamic_date_formats 字段；定义了规则orderId2：凡是遇到orderId这个字段，则将其转换为long型；定义了规则strings_as_keywords：凡是遇到string类型的字段都映射为keyword类型，norms属性为false；关于keyword类型和norms关键字，将在下面的数据类型小节介绍。

2）预定义映射

预定义映射和上面的区别就是预先把所有已知的字段类型描述写到mapping里，下图截取了一部分作为示例：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac229b2c5f3d?imageView2/0/w/1280/h/960/ignore-error/1)

图16中JSON结构的上半部分与动态模板相同，红框中内容内容为预先定义的属性：apply.applyInfo.appSubmissionTime, apply.applyInfo.applyId, apply.applyInfo.applyInputSource等字段，type表明了该字段的类型，映射定义完成后，再插入的数据必须符合字段定义，否则ES将返回异常。

#### 2.4.2 常用数据类型 ####

常用的数据类型有text, keyword, date, long, double, boolean, ip

实际使用中，将字符串类型定义为keyword而不是text，主要原因是text类型的数据会被当做文本进行语法分析，做一些分词、过滤等操作，而keyword类型则是当做一个完整数据存储起来，省去了多余的操作，提高索引性能。

配合keyword使用的还有一个关键词norm，置为false表示当前字段不参与评分；所谓评分是指根据单词的TF/IDF或其他一些规则，对查询出的结果赋予一个分值，供展示搜索结果时进行排序， 而一般的业务场景并不需要这样的排序操作（都有明确的排序字段），从而进一步优化查询效率。

#### 2.4.3 索引名无法修改 ####

初始化一个索引，都要在URL中明确指定一个索引名，一旦指定则无法修改，所以一般建立索引都要指定一个默认的别名(alias):

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" > "aliases" :{ "loan_alias" :{ }<br> }<br></span></p> 复制代码`

别名和索引名是多对多的关系，也就是一个索引可以有多个别名，一个别名也可以映射多个索引；在一对一这种模式下，所有用到索引名的地方都可以用别名进行替换；别名的好处就是可以随时的变动，非常灵活。

#### 2.4.4 Mapping中已存在的字段无法更新 ####

如果一个字段已经初始化完毕（动态映射通过插入数据，预定义通过设置字段类型），那就确定了该字段的类型，插入不兼容的数据则会报错，比如定义了一个long类型字段，如果写入一个非数字类型的数据，ES则会返回数据类型错误的提示。

这种情况下可能就需要重建索引，上面讲到的别名就派上了用场；一般分3步完成：

* 新建一个索引将格式错误的字段指定为正确格式;
* 2）使用ES的Reindex API将数据从旧索引迁移到新索引;
* 3）使用Aliases API将旧索引的别名添加到新索引上，删除旧索引和别名的关联。

上述步骤适合于离线迁移，如果要实现不停机实时迁移步骤会稍微复杂些。

#### 2.4.5 API ####

基本的操作就是增删改查，可以参考ES的官方文档：

[www.elastic.co/guide/en/el…]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fguide%2Fen%2Felasticsearch%2Freference%2Fcurrent%2Fdocs.html )

一些比较复杂的操作需要用到ES Script，一般使用类Groovy的painless script，这种脚本支持一些常用的JAVA API（ES安装使用的是JDK8，所以支持一些JDK8的API），还支持Joda time等。

举个比较复杂的更新的例子，说明painless script如何使用：

**需求描述**

appSubmissionTime表示进件时间，lenssonStartDate表示开课时间，expectLoanDate表示放款时间。要求2018年9月10日的进件，如果进件时间 与 开课时间的日期差小于2天，则将放款时间设置为进件时间。

Painless Script如下:

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" >POST loan_idx/_update_by_query<br> { "script" :{ "source" : "long getDayDiff(def dateStr1, def dateStr2){ <br> LocalDateTime date1= toLocalDate(dateStr1); LocalDateTime date2= toLocalDate(dateStr2); ChronoUnit.DAYS.between(date1, date2);<br> }<br> LocalDateTime toLocalDate(def dateStr)<br> { <br> DateTimeFormatter formatter = DateTimeFormatter.ofPattern(\"yyyy-MM-dd HH:mm:ss\"); LocalDateTime.parse(dateStr, formatter);<br> }<br> if(getDayDiff(ctx._source.appSubmissionTime, ctx._source.lenssonStartDate) < 2)<br> { <br> ctx._source.expectLoanDate=ctx._source.appSubmissionTime<br> }" , "lang" : "painless" <br> }<br> , "query" :<br> { "bool" :{ "filter" :[<br> { "bool" :{ "must" :[<br> { "range" :{ <br> "appSubmissionTime" :<br> {<br> "from" : "2018-09-10 00:00:00" , "to" : "2018-09-10 23:59:59" , "include_lower" : true , "include_upper" : true <br> }<br> }<br> }<br> ]<br> }<br> }<br> ]<br> }<br> }<br>}<br></span></p> 复制代码`

解释：整个文本分两部分，下半部分query关键字表示一个按范围时间查询（2018年9月10号），上半部分script表示对匹配到的记录进行的操作，是一段类Groovy代码（有Java基础很容易读懂），格式化后如下， 其中定义了两个方法getDayDiff()和toLocalDate()，if语句里包含了具体的操作：

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" >long getDayDiff(def dateStr1, def dateStr2){<br> LocalDateTime date1= toLocalDate(dateStr1);<br> LocalDateTime date2= toLocalDate(dateStr2);<br> ChronoUnit.DAYS.between(date1, date2);<br>}<br>LocalDateTime toLocalDate(def dateStr){<br> DateTimeFormatter formatter = DateTimeFormatter.ofPattern( "yyyy-MM-dd HH:mm:ss" );<br> LocalDateTime.parse(dateStr, formatter);<br>} if (getDayDiff(ctx._source.appSubmissionTime, ctx._source.lenssonStartDate) < 2){<br> ctx._source.expectLoanDate=ctx._source.appSubmissionTime<br>}<br></span></p> 复制代码`

然后提交该POST请求，完成数据修改。

#### 2.4.6 查询数据 ####

这里重点推荐一个ES的插件ES-SQL:

[github.com/NLPchina/el…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FNLPchina%2Felasticsearch-sql%2Fwiki%2FBasic-Queries-And-Conditions )

这个插件提供了比较丰富的SQL查询语法，让我们可以使用熟悉的SQL语句进行数据查询。其中，有几个需要注意的点：

* 

ES-SQL使用Http GET方式发送情况，所以SQL的长度是受限制的(4kb)，可以通过以下参数进行修改：http.max_initial_line_length: "8k"

* 

计算总和、平均值这些数字操作，如果字段被设置为非数值类型，直接使用ESQL会报错，可改用painless脚本。

* 

使用Select as语法查询出的结果和一般的查询结果，数据的位置结构是不同的，需要单独处理。

* 

NRT（Near Real Time）：准实时

向ES中插入一条记录，然后再查询出来，一般都能查出最新的记录，ES给人的感觉就是一个实时的搜索引擎，这也是我们所期望的，然而实际情况却并非总是如此，这跟ES的写入机制有关，做个简单介绍：

* Lucene 索引段 -> ES 索引

写入ES的数据，首先是写入到Lucene索引段中的，然后才写入ES的索引中，在写入ES索引前查到的都是旧数据。

* commit：原子写操作

索引段中的数据会以原子写的方式写入到ES索引中，所以提交到ES的一条记录，能够保证完全写入成功，而不用担心只写入了一部分，而另一部分写入失败。

* refresh：刷新操作，可以保证最新的提交被搜索到

索引段提交后还有最后一个步骤：refresh，这步完成后才能保证新索引的数据能被搜索到。

出于性能考虑，Lucene推迟了耗时的刷新，因此它不会在每次新增一个文档的时候刷新，默认每秒刷新一次。这种刷新已经非常频繁了，然而有很多应用却需要更快的刷新频率。如果碰到这种状况，要么使用其他技术，要么审视需求是否合理。

不过，ES给我们提供了方便的实时查询接口，使用该接口查询出的数据总是最新的，调用方式描述如下：

GET [http://IP]( https://link.juejin.im?target=http%3A%2F%2FIP ) :PORT/index_name/type_name/id

上述接口使用了HTTP GET方法，基于数据主键（id）进行查询，这种查询方式会同时查找ES索引和Lucene索引段中的数据，并进行合并，所以最终结果总是最新的。但有个副作用：每次执行完这个操作，ES就会强制执行refresh操作，导致一次IO，如果使用频繁，对ES性能也会有影响。

#### 2.4.7 数组处理 ####

数组的处理比较特殊，拿出来单独讲一下。

1）表示方式就是普通的JSON数组格式，如：

[1, 2, 3]、 [“a”, “b”]、 [ { "first" : "John", "last" : "Smith" },{"first" : "Alice", "last" : "White"} ]

2）需要注意ES中并不存在数组类型，最终会被转换为object，keyword等类型。

3）普通数组对象查询的问题。

普通数组对象的存储，会把数据打平后将字段单独存储，如：

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" >{ "user" :[<br> { "first" : "John" , "last" : "Smith" <br> },<br> { "first" : "Alice" , "last" : "White" <br> }<br> ]<br>}<br></span></p> 复制代码`

会转化为下面的文本

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" >{ "user.first" :[ "John" , "Alice" <br> ], "user.last" :[ "Smith" , "White" <br> ]<br>}<br></span></p> 复制代码`

将原来文本之间的关联打破了，图17展示了这条数据从进入索引到查询出来的简略过程：

* 

组装数据，一个JSONArray结构的文本。

* 

写入ES后，默认类型置为object。

* 

查询user.first为Alice并且user.last为Smith的文档（实际并不存在同时满足这两个条件的）。

* 

返回了和预期不符的结果。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac22afa778d4?imageView2/0/w/1280/h/960/ignore-error/1)

4）嵌套（Nested）数组对象查询

嵌套数组对象可以解决上面查询不符的问题，ES的解决方案就是为数组中的每个对象单独建立一个文档，独立于原始文档。如图18所示，将数据声明为nested后，再进行相同的查询，返回的是空，因为确实不存在user.first为Alice并且user.last为Smith的文档。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ac22b4ddceee?imageView2/0/w/1280/h/960/ignore-error/1)

5）一般对数组的修改是全量的，如果需要单独修改某个字段，需要借助painless script，参考： [www.elastic.co/guide/en/el…]( https://link.juejin.im?target=https%3A%2F%2Fwww.elastic.co%2Fguide%2Fen%2Felasticsearch%2Freference%2F5.6%2Fdocs-update.html )

### 2.5 安全 ###

数据安全是至关重要的环节，主要通过以下三点提供数据的访问安全控制：

* XPACK

XPACK提供了Security插件，可以提供基于用户名密码的访问控制，可以提供一个月的免费试用期，过后收取一定的费用换取一个license。

* IP白名单

是指在ES服务器开启防火墙，配置只有内网中若干服务器可以直接连接本服务。

* 代理

一般不允许业务系统直连ES服务进行查询，需要对ES接口做一层包装，这个工作就需要代理去完成；并且代理服务器可以做一些安全认证工作，即使不适用XPACK也可以实现安全控制。

### 2.6 网络 ###

ElasticSearch服务器默认需要开通9200、9300 这两个端口。

下面主要介绍一个和网络相关的错误，如果大家遇到类似的错误，可以做个借鉴。

* 

引出异常前，先介绍一个网络相关的关键词，keepalive ：

* 

Http keep-alive和Tcp keepalive。

HTTP1.1中默认启用"Connection: Keep-Alive"，表示这个HTTP连接可以复用，下次的HTTP请求就可以直接使用当前连接，从而提高性能，一般HTTP连接池实现都用到keep-alive；

TCP的keepalive的作用和HTTP中的不同，TPC中主要用来实现连接保活，相关配置主要是net.ipv4.tcp_keepalive_time这个参数，表示如果经过多长时间（默认2小时）一个TCP连接没有交换数据，就发送一个心跳包，探测下当前链接是否有效，正常情况下会收到对方的ack包，表示这个连接可用。

下面介绍具体异常信息，描述如下：

两台业务服务器，用restClient（基于HTTPClient，实现了长连接）连接的ES集群（集群有三台机器），与ES服务器分别部署在不同的网段，有个异常会有规律的出现：

每天9点左右会发生异常Connection reset by peer. 而且是连续有三个Connection reset by peer

` <p style= "line-height: 2em;" ><span style= "font-size: 14px;" >Caused by: java.io.IOException: Connection reset by peer <br> at sun.nio.ch.FileDispatcherImpl.read0(Native Method) <br> at sun.nio.ch.SocketDispatcher.read(SocketDispatcher.java:39) <br> at sun.nio.ch.IOUtil.readIntoNativeBuffer(IOUtil.java:223) <br> at sun.nio.ch.IOUtil.read(IOUtil.java:197)<br></span></p> 复制代码`

为了解决这个问题，我们尝试了多种方案，查官方文档、比对代码、抓包。。。经过若干天的努力，最终发现这个异常是和上面提到keepalive关键词相关（多亏运维组的同事帮忙）。

实际线上环境，业务服务器和ES集群之间有一道防火墙，而防火墙策略定义空闲连接超时时间为例如为1小时，与上面提到的linux服务器默认的例如为2小时不一致。由于我们当前系统晚上访问量较少，导致某些连接超过2小时没有使用，在其中1小时后防火墙自动就终止了当前连接，到了2小时后服务器尝试发送心跳保活连接，直接被防火墙拦截，若干次尝试后服务端发送RST中断了链接，而此时的客户端并不知情；当第二天早上使用这个失效的链接请求时，服务端直接返回RST，客户端报错Connection reset by peer，尝试了集群中的三台服务器都返回同样错误，所以连续报了3个相同的异常。解决方案也比较简单，修改服务端keepalive超时配置，小于防火墙的1小时即可。

## 参考 ##

《深入理解ElasticSearch》

[www.cnblogs.com/Creator/p/3…]( https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2FCreator%2Fp%2F3722408.html )

[yq.aliyun.com/articles/10…]( https://link.juejin.im?target=https%3A%2F%2Fyq.aliyun.com%2Farticles%2F108048 )

[www.cnblogs.com/LBSer/p/411…]( https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2FLBSer%2Fp%2F4119841.html )

[www.cnblogs.com/yjf512/p/53…]( https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2Fyjf512%2Fp%2F5354055.html )

作者：综合信贷雷鹏

来源：宜信技术学院