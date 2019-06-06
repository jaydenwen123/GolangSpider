# MongoDB入门 #

**数据库服务器** 本质上是一个用于存储数据的服务器；数据库服务器中可以放多个数据库。

### 一、3个概念 ###

**数据库(database)**

` 数据库是一个仓库；可以在仓库中放多个集合。 复制代码`

**集合(collection)**

` 集合类似于数据；在集合中可以存放多个文档。 复制代码`

**文档(document)**

` 文档是数据库中的最小单位，我们存储操作的内容都是文档。 复制代码`

**3个概念的关系**

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2226ede89575c?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 在MongoDB中，数据库和集合都不需要手动创建，当我们创建文档时，如果文档所在的集合或者数据库不存在会自动创建数据库。
> 
> 

### 二、数据库操作(database) ###

**启动数据库命令**

` mongo 复制代码`

**查询所有数据库**

` show dbs; //or show databases; 复制代码`

**进入（创建）数据库**

` //进入（创建）my_test数据库 use my_test; 复制代码`

如果 ` my_test` 数据库存在，则直接 ` 进入` 数据库；

如果 ` my_test` 不存在，则 ` 创建并进入` 数据库。

### 三、集合操作(collection) ###

**新建集合**

` //新建集合：number db.createCollection( 'number' ) 复制代码`

**查看集合List**

` show collections; 复制代码`

**删除集合**

` /** * @collectionName [集合名] */ db.collectionName.drop(); 复制代码`

### 四、1.文档操作——新增 ###

> 
> 
> 
> 友情提示：
> 
> 
> 
> 文档——增、删、改、查的案例是基于如下的基础上：
> 
> 
> 
> ​ 新建一个数据库 ` todos` ，新建一个集合 ` todo` 。
> 
> 
> 
> [增删改查（CRUD）官网API地址](
> https://link.juejin.im?target=https%3A%2F%2Fdocs.mongodb.com%2Fmanual%2Fcrud%2F
> )
> 
> 

**insertOne()**

插入一条数据

` /** * @todo [集合名] * @obj {Object} [要插入的数据] * * @TODO： * 只能插入一条 */ db.<collection>.insertOne(obj); 复制代码`

Example :

` db.todo.insertOne({ name : 'a1' , age : 10 }) //查询todo集合中的所有文档 db.todo.find(); //结果：{ "_id" : ObjectId("5cf64b69743e692169bbdbf7"), "name" : "a1", "age" : 10 } 复制代码`

**insertMany()**

插入多条数据

` /** * @arr {Array|Object} [要插入的多条数据] */ db.<collection>.insertMany(arr) 复制代码`

Example :

` db.todo.insertMany([ { name : 'a2' , age : 11 }, { name : 'a3' , age : 12 } ]); db.todo.find(); /** 结果： { "_id" : ObjectId("5cf64e57743e692169bbdbfa"), "name" : "a1", "age" : 10 } { "_id" : ObjectId("5cf64e5e743e692169bbdbfb"), "name" : "a2", "age" : 11 } { "_id" : ObjectId("5cf64e5e743e692169bbdbfc"), "name" : "a3", "age" : 12 } */ 复制代码`

**insert()**

插入多条 ` or` 插入单条

` /** * @data {Object|Array} [要插入的单条或者多条数据] * * @ TODO: * * 即2个方法的合集 */ db.todo.insert(data); 复制代码`

Example： 同上面 ` insertOne()` 和 ` insertMeny()` 的用法相同。

**关于_id**

> 
> 
> 
> 当插入一个文档没有指定 ` _id` ，则文档会根据时间戳自动生成一个 ` _id` 的值。
> 
> 
> 
> 如果指定了 ` _id` ，则使用 ` _id` 作为ID。
> 
> 

### 四、2.文档操作——查询 ###

**find()**

返回所有符合条件的数据；

` /** * @obj {Object} [需要筛选的条件] * @return {Array} [符合条件值的集合] */ db.todo.find(obj) 复制代码`

Example：

查询 ` name = a1` 的数据

` db.todo.find({ name : 'a1' }); // { "_id" : ObjectId("5cf67ae40f133022c604d5b8"), "name" : "a1", "age" : 10 } 复制代码`

**findOne()**

返回符合条数据中的第一条数据；

Example：

` /** * @obj {Object} [需要筛选的条件] * @return {Object} [符合条件的值] */ db.todo.find(obj) 复制代码`

**count()**

查询符合条件的数据有多少条；

Example：

` /** * @obj {Object} [需要筛选的条件] * @return {Number} [符合条件的值] */ db.todo.find(obj).count(); 复制代码`

### 四、3.文档操作——修改 ###

**update()**

替换 ` 1个` 或者 ` n个` 文档的对象；

修改 ` 1个` 或者 ` n个` 文档的对象的值；

删除 ` 1个` 或者 ` n个` 文档的对象的键；

` /** * @filterObj {Any|Object} [筛选条件] * @newObj {Any} [新的值] * $set {Object} [修改对象的值] * $unset {Object} [删除对象的键] * @config {Object} [配置项] * * multi {Boolean} [是否修改多个对象][default：false] * @ TODO: * * 默认值修改筛选条件的第一个对象 */ db.<collection>.update(filterObj, newObj, config); 复制代码`

Example：

` //将age=10的全部数据的name改为a14 db.todo.update( { //筛选出age=10的数据 age: 10 }, { //将原数据中的name设置为'a14' $set: { name : 'a14' } }, { //修改多个数据 multi: true } ) //删除掉 第一条name=a14的这条数据的name属性 db.todo.update( { name : 'a14' }, { $unset : { name : 'a14' } } ) 复制代码`
> 
> 
> 
> 
> ` $unset` 、 ` $set` 是MongoDB的文档操作符；
> 
> 
> 
> 更多的内容请查阅： [MongoDB官网操作符](
> https://link.juejin.im?target=https%3A%2F%2Fdocs.mongodb.com%2Fmanual%2Freference%2Foperator%2Fquery%2F
> )
> 
> 

**updateOne()**

更新一条数据；

**updateMany()**

更新多条数据；

**replaceOne()**

替换第一条数据；

### 四、4.文档操作——删除 ###

> 
> 
> 
> @TODO：
> 
> 
> 
> ​ 在实际开发中，这个操作并不常用，因为实际上所有的删除操作知识改变了数据的状态，保证不会被输出，但是实际上数据还是存在于数据库中。
> 
> 

**remove()**

删除多个文档或第一个；

**deleteOne()**

删除第一个；

**deleteMany()**

删除多个；

Example All Api ：

` /** * 默认：删除符合条件所有的文档 @removeOne: true; 只删除查询条件的第一个 **/ db.<collection>.remove(query, removeOne); //删除多个 db.<collection>.deleteOne(query); //删除多个 db.<collection>.deleteMany(query); 复制代码`

### 四、5.文档操作——练习题 ###

* 习题的知识点汇总：

**比较操作符：**

` $gt` ：大于

` $gte` ：大于等于

` $lt` ：小于

` $lte` ：小于等于

` $eq` ：等于，必须要全等；

**修改/新增/删除操作符**

` $set` ：修改[value]；新增{[key]: [value]}

` unset` : 删除[key]

**添加操作符**

` $push` ：向数组中添加一个值

` $addToSet` ：向数组中添加一个值（被添加进来的值在这个数组中必须不存在）

**方法**

` limit(n)` ：取出n条数据；

` skip(m)` ：跳过第m条之前的数据，即从 ` m+1` 条开始取。

Example ：

` /** * [分页] * @page {Number} [当前页] * @page_size {Number} [每页条数] * * @return {Object} * @data {Array} [符合条件的数据] * @total {Number} [总数量] */ function getDataList ( page, page_size ) { var skip = ( page - 1 ) + page_size; var data = db.<collection>.find({ /*筛选条件*/ }); return { data : data.limit(page_size).skip(skip), total : data.count() }; } 复制代码`

* 

习题：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25d07a8869a25?imageView2/0/w/1280/h/960/ignore-error/1)

* 

答案：

` //1 use my_test //2 db.user.insert({ username : 'sunwukong' }); //3 db.user.find(); //4 db.user.insert({ username : 'zhubajie' }); //5 db.user.find(); //6 db.user.find().count(); //7 db.user.find({ username : 'sunwukong' }) //8 db.user.update( { username : 'sunwukong' }, { $set : { address : 'huaguoshan' } }, { //修改多个 multi: true } ) //9 db.user.update( { username : 'zhubajie' }, { $set : { username : 'tangseng' } }, { multi : true } ) //10 db.user.update( { username : 'sunwukong' , }, { $unset : { // 这里的值随便填写，主要是给：`$unset`添加一个键，准备删除。 address: 'delete' } }, { multi : true } ) //11 db.user.update( { username : 'sunwukong' }, { $set : { hobby : { cities : [ 'beijing' , 'shanghai' , 'shenzhen' ], movies : [ 'sanguo' , 'hero' ] } } }, { multi : true } ) //12 db.user.update( { username : 'tangseng' }, { $set : { hobby : { movies : [ 'A Chinese Odyssey' , 'King of comedy' ] } } }, { multi : true } ) //13 db.user.find( { 'hobby.movies' : 'hero' } ) //14 db.user.update( { username : 'tangseng' }, { // 操作符：$push可以向数组中添加数据 // 操作符: $addToSet也可以；区别是`$addToSet`对于重复数据不添加；$push则不会。 $push: { 'hobby.movies' : 'Intersterllar' } } ) //15 db.user.remove( { 'hobby.cities' : 'beijing' } ) //16 db.user.drop(); //17.插入20,00条数据 var data = []; for ( var i = 1 ; i <= 1000 * 20 ; i++) { data.push( { num : i } ) } db.numbers.insert(data) //创建numbers集合并插入2w条数据 //18.查询num = 500的数据 numbers.find({ num : 500 }) //or numbers.find({ num : { $eq : 500 } }) //区别： // 前者会匹配数组中值为500的数据； // 后者只匹配num=500的数据； //19.查询num>5000的文档 numbers.find({ num : { $gt : 5000 }) //20.查询num>=19996的文档 numbers.find({ num : { $gte : 19996 }) //21.查询num<30的文档 numbers.find({ num : { $lt : 30 }) //22.查询 40 < num < 50 的文档 numbers.find({ num : { $lt : 50 , $gt : 40 } }) //23.查询集合中前10条数据 numbers.find().limit( 10 ) //24.查询集合中第11-20条数据 numbers.find().limit( 10 ).skip( 10 ) //25.查询第21-30条数据 numbers.find().limit( 10 ).skip( 20 ) 复制代码`

### 五、1.文档间的关系 ###

**一对一** （one to one）

一个人 对应 一个身份证号码

` db.human.insert({ name : 'david' , idcard : { idcard : 123465789102345678 } }) 复制代码`

**一对多** （one to many）/ **多对一** （many to one）

一个账户 对应 多个订单

` /** 用户david有订单 d1、d2，可以使用d1、d2查询到商品信息； 商品通过user_id可以查询到用户； 一个订单只能对应一个用户，一个用户可以拥有多个订单。 */ db.users.insert({ _id : 1 , name : 'david' , //订单 orders: [ 'd1' , 'd2' ] }) db.orders.insert([ { user_id : 1 , id : 'd1' , name : '商品1' }, { user_id : 1 , id : 'd1' , name : '商品1' } ]) 复制代码`

**多对多** （many to many）

多个老师对应多个学生

` /** */ db.teachers.insert([ { tid: 't1', name: 't1', students: ['s1', 's2'] }, { tid: 't2', name: 't2', students: ['s1', 's2', 's3'] }, { tid: 't3', name: 't3', students: ['s1'] } ]) db.students.insert([ { sid: 's1', name: 's1 students: ['t1', 't2'] }, { sid: 's2', name: 's2', students: ['t1', 't2', 't3'] }, { sid: 's3', name: 's3', students: ['t3'] } ]) 复制代码`

### 五、2.文档间的关系——习题 ###

知识点归纳：

` $or` ：或操作符，详细用法见 ` 29` 题答案

` $inc` ：自增操作符，详细用法见 ` 33` 题答案

原始数据：

` //dept.json /* 1 */ { "_id" : ObjectId("5cf7c7c0ffd6ea864b5a22cc"), "deptno" : 10.0, "dname" : "财务部", "loc" : "北京" } /* 2 */ { "_id" : ObjectId("5cf7c7c0ffd6ea864b5a22cd"), "deptno" : 20.0, "dname" : "办公室", "loc" : "上海" } /* 3 */ { "_id" : ObjectId("5cf7c7c0ffd6ea864b5a22ce"), "deptno" : 30.0, "dname" : "销售部", "loc" : "广州" } /* 4 */ { "_id" : ObjectId("5cf7c7c0ffd6ea864b5a22cf"), "deptno" : 40.0, "dname" : "运营部", "loc" : "深圳" } //emp.json /* 1 */ { "_id" : ObjectId("5cf7ca79ffd6ea864b5a22d0"), "empno" : 7369.0, "ename" : "林冲", "job" : "职员", "mgr" : 7902.0, "hiredate" : ISODate("1980-12-16T16:00:00.000Z"), "sal" : 1200.0, "depno" : 20.0 } /* 2 */ { "_id" : ObjectId("5cf7ca79ffd6ea864b5a22d1"), "empno" : 7499.0, "ename" : "孙二娘", "job" : "销售", "mgr" : 7698.0, "hiredate" : ISODate("1981-02-19T16:00:00.000Z"), "sal" : 1600.0, "comm" : 300.0, "depno" : 30.0 } /* 3 */ { "_id" : ObjectId("5cf7ca79ffd6ea864b5a22d2"), "empno" : 7521.0, "ename" : "扈三娘", "job" : "销售", "mgr" : 7698.0, "hiredate" : ISODate("1981-02-19T16:00:00.000Z"), "sal" : 800.0, "comm" : 500.0, "depno" : 30.0 } /* 4 */ { "_id" : ObjectId("5cf7ca79ffd6ea864b5a22d3"), "empno" : 7566.0, "ename" : "卢俊义", "job" : "经理", "mgr" : 7839.0, "hiredate" : ISODate("1981-02-19T16:00:00.000Z"), "sal" : 2975.0, "depno" : 20.0 } /* 5 */ { "_id" : ObjectId("5cf7ca79ffd6ea864b5a22d4"), "empno" : 7654.0, "ename" : "潘金莲", "job" : "销售", "mgr" : 7839.0, "hiredate" : ISODate("1981-02-19T16:00:00.000Z"), "sal" : 2975.0, "depno" : 20.0 } 复制代码`

问题：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b273be62f5de4c?imageView2/0/w/1280/h/960/ignore-error/1)

问题答案：

` var emp = db.getCollection( 'emp' ) var dept = db.getCollection( 'dept' ) //33.为所有工资低于1000的员工增加工资400元 emp.update({ sal : { $lte : 1000 } }, { $inc : { sal : 400 } }) //32.查询所有mgr为7698的所有员工 emp.find({ mgr : 7698 }); //31.查询销售部所有员工 var cwbId = dept.findOne({ dname : '销售部' }).deptno emp.find({ depno : cwbId}) //30.查询财务部的所有员工 var cwbId = dept.findOne({ dname : '财务部' }).deptno emp.find({ depno : cwbId}) //29.工资 >= 2500 或者 <= 1000 emp.find({ $or : [{ sal : { $gte : 2500 }}, { sal : { $lte : 1000 } }] }) //28.工资1000-2000 emp.find({ sal : { $lte : 2000 , $gte : 1000 } }) //27.工资<= 2000 emp.find({ sal : { $lte : 2000 } }); 复制代码`

### 六、排序和投影 ###

**排序**

` sort()` 指定查询文档的筛选条件; ` sort()` 、 ` limit()` 、 ` skip()` 排序部分先后，因为总会先进行排序。

` 1` ：升序排列

` -1` ：降序排列

> 
> 
> 
> 默认：如果不指定 ` sort()` ，则按照 ` _id` 排序，而 ` _id = 时间戳 + 机器码` 组成，所以实际上默认是按照 `
> 创建时间` 排序。
> 
> 

Example ：

` //按照 工资 降序 排列 emp.find().sort({ sal : -1 }) //先按照 工资(sal) 升序；如果工资相同，再按照编号(empno) 降序排列 emp.find().sort({ sal : 1 , empno : -1 }) 复制代码`

**映射**

指定查询时，需要返回的字段。

` 1` ：允许映射;

` 0` ：禁止映射;

` //返回 工资(sal) >= 1250 的员工姓名 emp.find({ sal : { $gte : 1250 } }, { ename : 1 }) /*_id会自动加上 {"_id" : ObjectId("5cf7ca79ffd6ea864b5a22d1"),"ename" : "孙二娘"}, {"_id" : ObjectId("5cf7ca79ffd6ea864b5a22d2"),"ename" : "扈三娘"} {"_id" : ObjectId("5cf7ca79ffd6ea864b5a22d3"),"ename" : "卢俊义"} {"_id" : ObjectId("5cf7ca79ffd6ea864b5a22d4"),"ename" : "潘金莲"} */ //如果想忽略_id，可以尝试： emp.find({ sal : { $gte : 1250 } }, { ename : 1 , _id : 0 }) 复制代码`

### 七、1.操作MongoDB的库——Mongoose ###

Mongoose是一个对象文档模型（ODM）库，他对Node原生的MongoDB模块进行了进一步的优化封装，并提供了更多的功能。

**为什么要使用mongoose?**

* 可以为文档创建一个模式结构（Schema|约束）
* 可以对模型中的对象/文档进行验证
* 数据可以通过类型转换转换为对象模型
* 可以使用中间件来应用业务逻辑挂钩
* 比Node原生的MongoDB驱动更容易

### 七、2.mongoose——3个概念 ###

**Schema（模式对象）**

Schema对象定义约束数据库中的文档结构

**Model**

Model对象作为集合中的所有文档的表示，相当于MongoDB数据库中的集合collection

**Document**

Document表示集合中的具体文档，相当于集合中的一个具体文档。

> 
> 
> 
> 创建顺序：先有Schema约束Model，在用Model操作Document
> 
> 

### 七、3.Mongoose——第一个DEMO ###

连接数据库，并插入第数据

` const mongoose = require ( 'mongoose' ); var Schema = mongoose.Schema; mongoose.connect( 'mongodb://localhost:27017/mongoose_test' , { useNewUrlParser : true }, function ( ) { console.log( '数据库连接成功' ) }); var stuSchema = new Schema({ name : String , age : Number , gender : { type : String , default : 'woman' }, address : String }) var StuModel = mongoose.model( 'student' , stuSchema) StuModel.create({ name : '孙悟空' , age : 18 , gender : 'man' , address : '花果山' }, function ( err ) { if (!err) { console.log( '插入成功' ) } }) StuModel.create({ name : '白骨精' , age : 18 , address : '盘丝洞' }, function ( err ) { if (!err) { console.log( '插入成功' ) } }) 复制代码`

### 七、4.使用mongoose的Model进行增删改查 ###

**Model的增/删/改/查（CRUD）**

`.count()` : 查询数量

`.create()` : 新增

`.find()` : 查询

`.update()` : 修改

`.remove()` : 删除

` const mongoose = require ( 'mongoose' ); const Schema = mongoose.Schema let datas = [ { name : '白骨精' , age : 19 , address : '盘丝洞' }, { name : '孙悟空' , age : 20 , address : '花果山' }, { name : '猪八戒' , age : 21 , address : '高老庄' }, { name : '沙和尚' , age : 22 , address : '流沙河' } ] const userSchema = new Schema({ name : String , age : Number , sex : { type : String , default : 'man' }, address : String }) const userModel = mongoose.model( 'users' , userSchema) //查询数量 userModel.count({}, function ( err, count ) { if (!err) { console.log(count) } }) //写入数据 userModel.create(datas, function ( err ) { if (!err) { console.log( '写入成功' ) } }) /** * [查询] * model.find(conditions, [projection], [options], [callback]) * @conditions {Object} [筛选条件] * @projection {Object|String} [投影] * * { name: 1, _id: 0 } * * 'name -_id' * @options {Object} [查询选项：skip limit] * @callback {Function} [回调函数] * */ userModel.find({}, function ( err, datas ) { if (!err) { console.log( datas ) } }) /** * [修改] * model.update(conditions, docs, [options], [callback]) * @conditions {Object} [筛选条件] * @docs {Object|String} [修改的值] * @options {Object} [查询选项：skip limit] * @callback {Function} [回调函数] * */ userModel.update({ name : '孙悟空' }, { $set : { age : 20 } }, function ( err ) { if (!err) { console.log( '修改成功' ) } }) /** * [删除] * model.remove(conditions, [callback]) * @conditions {Object} [筛选条件] * @callback {Function} [回调函数] * */ userModel.remove({ }, function ( err ) { if (!err) { console.log( '删除成功' ) } }) 复制代码`

### 七、5.使用mongoose的Document进行增删改 ###

Document是Model的一个实例，通过Model查询到的结果都是Document

**Document的增/删/改（CUD）**

` const mongoose = require ( 'mongoose' ); const Schema = mongoose.Schema mongoose.connect( 'mongodb://localhost:27017/mongoose_test' , { useNewUrlParser : true }, function ( err ) { if (!err) { console.log( '数据库连接成功' ) } else { console.log( '数据库连接失败' + err) } }) const userSchema = new Schema({ name : String , age : Number , sex : { type : String , default : 'man' }, address : String }) const userModel = mongoose.model( 'users' , userSchema) let datas = [ { name : '白骨精' , age : 19 , address : '盘丝洞' }, { name : '孙悟空' , age : 20 , address : '花果山' }, { name : '猪八戒' , age : 21 , address : '高老庄' }, { name : '沙和尚' , age : 22 , address : '流沙河' } ] //插入一条数据 var user = new userModel({ name : '玉帝' , age : 10000 , sex : 'woman' , address : '天庭' }) user.save( function ( err, product ) { if (!err) { console.log( '保存成功' ) } }) // 修改数据1 userModel.findOne({}, function ( err, doc ) { if (!err) { doc.update({ $set : { sex : 'woman' } }, function ( err ) { if (!err) { console.log( '修改成功' ) } }) } }) // 修改数据2 userModel.findOne({}, function ( err, doc ) { console.log( doc.toJSON() ) console.log( doc.toObject() ) console.log( doc.toString() ) if (!err) { doc.age = 80 doc.save( function ( err ) { if (!err) { console.log( '修改成功' ) } }) } }) // 删除数据 userModel.findOne({ name : '玉帝' }, function ( err, doc ) { if (!err) { doc.remove( function ( err ) { if (!err) { console.log( '删除成功' ) } }) } }) 复制代码`

关于更多的 ` mongoose` api请参阅 [mongoose官网文档]( https://link.juejin.im?target=https%3A%2F%2Fmongoosejs.com%2Fdocs%2Fapi.html%23model_Model.create )

### 七、6.mongoose模块化处理 ###

**为什么要模块化？**

在没有模块化之前，连接数据库，创建 ` Schame` 、 ` Model` 、数据CRUD都是在 ` index.js` 文件中进行的，但是当一个应用有各种路由的时候，需要在各个页面操纵各种 ` Model` ,这样在每个页面操纵 ` Model` 的时候非常的困难，因此需要模块化解决这个问题。

**实施模块化**

未整理之前的目录：

` /index.js 复制代码`

模块化之后的目录：

` /models /user.js /schame User.js index.js mongoose_init.js 复制代码`

内容如下：

` /schames/User.js`

` const mongoose = require ( 'mongoose' ); const Schema = mongoose.Schema const userSchema = new Schema({ name : String , age : Number , sex : { type : String , default : 'man' }, address : String }) module.exports = userSchema 复制代码`

` /models/user.js`

` const mongoose = require ( 'mongoose' ); const User = require ( '../schames/User' ) const user = mongoose.model( 'users' , User) module.exports = user 复制代码`

` mongoose_init.js`

` const mongoose = require ( 'mongoose' ); mongoose.connect( 'mongodb://localhost:27017/mongoose_test' , { useNewUrlParser : true }, function ( err ) { if (!err) { console.log( '数据库连接成功' ) } else { console.log( '数据库连接失败' + err) } }) 复制代码`

` index.js`

` require ( './mongoose_init' ) const user = require ( './models/user' ) user.find({}, function ( err, data ) { if (!err) { console.log(data) //查询成功 } }) 复制代码`