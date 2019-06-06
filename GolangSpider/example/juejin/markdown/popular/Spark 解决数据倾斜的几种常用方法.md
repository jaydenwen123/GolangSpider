# Spark 解决数据倾斜的几种常用方法 #

数据倾斜是大数据计算中一个最棘手的问题，出现数据倾斜后，Spark 作业的性能会比期望值差很多。数据倾斜的调优，就是利用各种技术方案解决不同类型的数据倾斜问题，保证 Spark 作业的性能。

### 一，数据倾斜原理 ###

一个 Spark 作业，会根据其内部的 Action 操作划分成多个 job，每个 job 内部又会根据 shuffle 操作划分成多个 stage，然后每个 stage 会分配多个 task 去执行任务。每个 task 会领取一个 partition 的数据处理。

同一个 stage 内的 task 是可以并行处理数据的，具有依赖关系的不同 stage 之间是串行处理的。由于这个处理机制，假设某个 Spark 作业中的某个 job 有两个 stage，分别为 stage0 和 stage1，那么 stage1 必须要等待 stage0 处理结束才能进行。如果 stage0 内部分配了 n 个 task 进行计算任务，其中有一个 task 领取的 partition 数据过大，执行了 1 个小时还没结束，其余的 n-1 个 task 在半小时内就执行结束了，都在等这最后一个 task 执行结束才能进入下一个 stage。这种由于某个 stage 内部的 task 领取的数据量过大的现象就是数据倾斜。

下图就是一个例子：hello 这个 key 对应了 7 条数据，映射到了同一个 task 去处理了，剩余的 2 个 task 分别只处理了一个数据。

![数据倾斜](https://user-gold-cdn.xitu.io/2019/6/6/16b2bdb2d719cb2a?imageView2/0/w/1280/h/960/ignore-error/1)

### 二，数据倾斜发生的现象 ###

1，绝大多数task执行得都非常快，但个别task执行极慢。比如，总共有1000个task，997个task都在1分钟之内执行完了，但是剩余两三个task却要一两个小时。这种情况很常见。

2，原本能够正常执行的Spark作业，某天突然报出OOM（内存溢出）异常，观察异常栈，是我们写的业务代码造成的。这种情况比较少见。

### 三，如何定位数据倾斜的代码 ###

数据倾斜只会发生在 shuffle 过程中。常用并且可能会触发 shuffle 操作的算子有：distinct，groupByKey，reduceByKey，aggregateByKey，join，cogroup，repartition 等。出现数据倾斜，很有可能就是使用了这些算子中的某一个导致的。

如果我们是 yarn-client 模式提交，我们可以在本地直接查看 log，在 log 中定位到当前运行到了哪个 stage；如果用的 yarn-cluster 模式提交的话，我们可以通过 spark web UI 来查看当前运行到了哪个 stage。无论用的哪种模式我们都可以在 spark web UI 上面查看到当前这个 stage 的各个 task 的数据量和运行时间，从而能够进一步确定是不是 task 的数据分配不均导致的数据倾斜。

当确定了发生数据倾斜的 stage 后，我们可以找出会触发 shuffle 的算子，推算出发生倾斜的那个 stage 对应代码。触发 shuffle 操作的除了上面提到的那些算子外，还要注意使用 spark sql 的某些 sql 语句，比如 group by 等。

### 四，解决方法 ###

解决数据倾斜的思路就是保证每个 stage 内部的 task 领取的数据足够均匀。一个是想办法让数据源在 Spark 内部计算粒度的这个维度上划分足够均匀，如果做不到这个，就要相办法将读取的数据源进行加工，尽量保证均匀。大致有以下几种方案。

#### 1，聚合源数据和过滤导致倾斜的 key ####

**a，聚合源数据**

假设我们的某个 Spark 作业的数据源是每天 ETL 存储到 Hive 中的数据，这些数据主要是电商平台每天用户的操作日志。Spark 作业中分析的粒度是 session，那么我们可以在往 Hive 中写数据的时候就保证每条数据对应一个 session 的所有信息，也就是以 session 为粒度的将数据写入到 Hive 中。

这样我们可以保证一点，我们 Spark 作业中就没有必要做一些 groupByKey + map 的操作了，可以直接对每个 key 的 value 进行 map 操作，计算我们需要的数据。省去了 shuffle 操作，避免了 shuffle 时的数据倾斜。

但是，当我们 Spark 作业中分析粒度不止一个粒度，比如除了 session 这个粒度外，还有 date 的粒度，userId 的粒度等等。这时候是无法保证这些所有粒度的数据都能聚合到一条数据上的。这时候我们可以做个妥协，选择一个相对比较大的粒度，进行聚合数据。比如我们按照原来的存储方式可能有 100W 条数据，但按照某个粒度，比如 date 这个粒度，进行聚合后存储，这样的话我们的数据可以降到 50W 条，可以做到减轻数据倾斜的现象。

**b，过滤导致倾斜的 key**

比如说我们的 Hive 中数据，共有 100W 个 key，其中有 5 个 key 对应的数据量非常的大，可能有几十万条数据(这种情况在电商平台上发生恶意刷单时候会出现)，其它的 key 对应数据量都只有几十。如果我们业务上面能够接受这 5 个 key 对应的数据可以舍弃。这种情况下我们可以在用 sql 从 Hive 中取数据时过滤掉这个 5 个 key。从而避免了数据倾斜。

#### 2，shuffle 操作提高 reduce 端的并行度 ####

Spark 在做 shuffle 操作时，默认使用的是 HashPartitioner 对数据进行分区。如果 reduce 端的并行度设置的不合适，很可能造成大量不同的 key 被分配到同一个 task 上去，造成某个 task 上处理的数据量大于其他 task，造成数据倾斜。

如果调整 reduce 端的并行度，可以让 reduce 端的每个 task 处理的数据减少，从而缓解数据倾斜。

设置方法：在对 RDD 执行 shuffle 算子时，给 shuffle 算子传入一个参数，比如 reduceByKey(1000) ，该参数就设置了这个 shuffle 算子执行时 shuffle read task 的数量。对于 Spark SQL 中的 shuffle 类语句，比如 group by、join 等，需要设置一个参数，即 spark.sql.shuffle.partitions ，该参数代表了 shuffle read task 的并行度，该值默认是 200 ，对于很多场景来说都有点过小。

#### 3，使用随机 key 进行双重聚合 ####

在 Spark 中使用 groupByKey 和 reduceByKey 这两个算子会进行 shuffle 操作。这时候如果 map 端的文件每个 key 的数据量偏差很大，很容易会造成数据倾斜。

我们可以先对需要操作的数据中的 key 拼接上随机数进行打散分组，这样原来是一个 key 的数据可能会被分到多个 key 上，然后进行一次聚合，聚合完之后将原来拼在 key 上的随机数去掉，再进行聚合，这样对数据倾斜会有比较好的效果。

具体可以看下图：

![随机 key 双重聚合](https://user-gold-cdn.xitu.io/2019/6/6/16b2bdb2d6264d29?imageView2/0/w/1280/h/960/ignore-error/1)

这种方案对 RDD 执行 reduceByKey 等聚合类 shuffle 算子或者在 Spark SQL 中使用 group by 语句进行分组聚合时，比较适用这种方案。

#### 4， 将 reduce-side join 转换成 map-side join ####

两个 RDD 在进行 join 时会有 shuffle 操作，如果每个 key 对应的数据分布不均匀也会有数据倾斜发生。

这种情况下，如果两个 RDD 中某个 RDD 的数据量不大，可以将该 RDD 的数据提取出来，然后做成广播变量，将数据量大的那个 RDD 做 map 算子操作，然后在 map 算子内和广播变量进行 join，这样可以避免了 join 过程中的 shuffle，也就避免了 shuffle 过程中可能会出现的数据倾斜现象。

#### 5，采样倾斜 key 并拆分 join 操作 ####

当碰到这种情况：两个 RDD 进行 join 的时候，其中某个 RDD 中少数的几个 key 对应的数据量很大，导致了数据倾斜，而另外一个 RDD 数据相对分布均匀。这时候我们可以采用这种方法。

1，对包含少数几个数据量过大的 key 的那个 RDD，通过 sample 算子采样出一份样本来，然后统计一下每个 key 的数量，计算出来数据量最大的是哪几个 key。

2，然后将这几个 key 对应的数据从原来的 RDD 中拆分出来，形成一个单独的 RDD，并给每个 key 都打上 n 以内的随机数作为前缀，而不会导致倾斜的大部分 key 形成另外一个 RDD。

3，接着将需要 join 的另一个 RDD，也过滤出来那几个倾斜 key 对应的数据并形成一个单独的 RDD，将每条数据膨胀成 n 条数据，这 n 条数据都按顺序附加一个 0~n 的前缀，不会导致倾斜的大部分 key 也形成另外一个 RDD。

4，再将附加了随机前缀的独立 RDD 与另一个膨胀 n 倍的独立 RDD 进行 join，此时就可以将原先相同的 key 打散成 n 份，分散到多个 task 中去进行 join 了。

5， 而另外两个普通的 RDD 就照常 join 即可。 最后将两次 join 的结果使用 union 算子合并起来即可，就是最终的 join 结果。

具体可以看下图：

![采样倾斜 key 分别 join](https://user-gold-cdn.xitu.io/2019/6/6/16b2bdb2d813fa26?imageView2/0/w/1280/h/960/ignore-error/1)

#### 6，使用随机前缀和扩容 RDD 进行 join ####

如果在进行 join 操作时，RDD 中有大量的 key 导致数据倾斜，那么进行分拆 key 也没什么意义，这时候可以采取这种方案。

该方案的实现思路基本和上一种类似，首先查看 RDD 中数据分布情况，找到那个造成数据倾斜的 RDD，比如有多个 key 映射数据都很大。然后我们将该 RDD 没调数据都打上一个 n 以内的随机前缀。同时对另一个 RDD 进行扩容，将其没调数据扩容成 n 条数据，扩容出来的每条数据依次打上 0~n 的前缀。然后将这两个 RDD 进行 join。

这种方案和上一种比，少了取样的过程，因为上一种是针对某个 RDD 中只有非常少的几个 key 发生数据倾斜，需要针对这几个 key 特殊处理。而这个方案是针对某个 RDD 有大量的 key 发生数据倾斜，这时候就没必要取样了。