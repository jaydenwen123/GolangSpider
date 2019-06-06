# Golang数据库编程之GORM库入门 #

在上一篇文章中我们讲解了使用Go语言的标准库 ` sql/database` 包操作数据库的过程，虽然使用 ` sql/database` 包操作数据也是挺方便的，但是需要自己写每一条SQL语句，因此我们可能会自己再度进行封装，以便更好地使用，而使用现有Go语言开源 ` ORM` 框架则是代替自己封装的一个更好的方式。

> 
> 
> 
> ` ORM` ，即对象关系映射( ` Object Relational Mapping` )，可以简单理解为将 ` 关系型数据库` 中的 `
> 数据表` 映射为编程语言中的具体的数据类型(如 ` struct` )，而 ` GORM` 库就是一个使用Go语言实现的且功能非常完善易使用的 `
> ORM` 框架。
> 
> 

下面一起来探索一下如何使用 ` GORM` 框架吧！

## 特性 ##

* 关联 (Has One, Has Many, Belongs To, Many To Many, 多态)
* 钩子 (在创建/保存/更新/删除/查找之前或之后)
* 预加载
* 事务
* 复合主键
* SQL 生成器
* 数据库自动迁移
* 自定义日志
* 可扩展性, 可基于 GORM 回调编写插件

## 如何安装 ##

安装 ` GORM` 非常简单，使用 ` go get -u` 就可以在 ` GOPATH` 目录下安装最新 ` GROM` 框架。

` go get -u github.com/jinzhu/gorm 复制代码`

安装之后，便可以使用 ` import` 关键字导入 ` GORM` 库，开始使用啦！

` import "github.com/jinzhu/gorm" 复制代码`

## 支持的数据库 ##

` GORM` 框架支持 ` MySQL` , ` SQL Server` , ` Sqlite3` , ` PostgreSQL` 四种数据库驱动，如果我们要连接这些数据库，则需要导入不同的驱动包及定义不同格式的 ` DSN` ( ` Data Source Name` )。

#### MySQL ####

##### 1. 导入 #####

` import _ "github.com/jinzhu/gorm/dialects/mysql" //或者//import _ "github.com/go-sql-driver/mysql" 复制代码`

##### 2. DSN #####

` //user指用户名，password指密码,dbname指数据库名 "user:password@/dbname?charset=utf8&parseTime=True&loc=Local" 复制代码`

#### SQL Server ####

##### 1. 导入 #####

` import _ "github.com/jinzhu/gorm/dialects/mssql" 复制代码`

##### 2. DSN #####

` //username指用户名，password指密码,host指主机地址，port指端口号，database指数据库名 "sqlserver://username:password@host:port?database=dbname" 复制代码`

#### Sqlite3 ####

##### 1. 导入 #####

` import _ "github.com/jinzhu/gorm/dialects/sqlite" 复制代码`

##### 2. DSN #####

连接Sqlite3数据库的DSN只需要指定Sqlite3的数据库文件的路径即可，如：

` //数据库路径 /tmp/gorm.db 复制代码`

#### PostgreSQL ####

##### 1. 导入 #####

` import _ "github.com/jinzhu/gorm/dialects/postgres" 复制代码`

##### 2. DSN #####

` //host指主机地址,port指端口号,user指用户名,dbname指数据库名,password指密码 host=myhost port=myport user=gorm dbname=gorm password=mypassword 复制代码`

## 连接数据库 ##

上面我们定义了连接不同的数据库的DSN，下面演示如果连接数据库，使用 ` gorm.Open()` 方法可以初始化并返回一个 ` gorm.DB` 结构体，这个结构体封装了GORM框架所有的数据库操作方法，下面是 ` gorm.Open()` 方法的定义：

` func Open(dialect string, args ...interface{}) (db *DB, err error) 复制代码`

示例代码：

` package main import "github.com/jinzhu/gorm" import _ "github.com/jinzhu/gorm/dialects/mysql" //导入连接MySQL数据库的驱动包 //DSN const DSN = "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local" //指定驱动 const DRIVER = "mysql" var db *gorm.DB func init () { var err error db,err = gorm.Open(DRIVER,DSN) if err != nil{ panic(err) } } func main (){ defer db.Close()//退出前执行关闭 //调用db执行具体的逻辑 } 复制代码`

在上面的例子中，我们在init方法中初始化 ` gorm.DB` 结构体，这样在下面的例子可以直接使用变量 ` db` 直接进行数据库操作。

## 基本操作 ##

使用 ` gorm.Open()` 函数返回一个 ` gorm.DB` 结构体后，我们可以使用 ` gorm.DB` 结构体提供的方法操作数据库，下面我们演示如何使用 ` gorm.DB` 进行创建、查询、更新、删除等最基本的操作。

其实 ` gorm.DB` 是在Go语言的 ` database/sql` 库中的 ` sql.DB` 结构体上再封装，因为 ` gorm.DB` 提供许多和 ` sql.DB` 一样的方法，如下所示：

` func (s *DB) Exec(sql string, values ...interface{}) *DB func (s *DB) Row() *sql.Row func (s *DB) Rows() (*sql.Rows, error) func (s *DB) Scan(dest interface{}) *DB 复制代码`

另外，使用 ` gorm.DB` 结构体中的 ` DB()` 方法，可以返回一个 ` sql.DB` 对象，如下：

` func (s *DB) DB() *sql.DB 复制代码`

下面演示的是使用 ` gorm.DB` 结构体中一些更简便的方法进行数据库基本操作，不过，在演示之前，我们需要先定义一个模型，如下:

` type User struct { Id int //对应数据表的自增id Username string Password string Email string Phone string } 复制代码`

我们定义了一个名称为 ` User` 的结构体， ` GROM` 支持将结构体按规则映射为某个数据表的一行，结构体的每个字段表示数据表的列，结构体的字段首字母必须是大写的。

#### 创建 ####

使用 ` gorm.DB` 中的 ` Create()` 方法， ` GORM` 会根据传给 ` Create()` 方法的模型，向数据表插入一行。

` func (s *DB) Create(value interface{}) *DB //创建一行 func (s *DB) NewRecord(value interface{}) bool //根据自增id判断主键是否存在 复制代码`

示例

` func main () { defer db.Close() //具体的逻辑 u := &User{Username: "test_one" , Password: "testOne123456" , Email: "test_one@163.com" , Phone: "13711112222" } db.Create(u) if db.NewRecord(u){ fmt.Println( "写入失败" ) } else { fmt.Println( "写入成功" ) } } 复制代码`

#### 查询 ####

` GROM` 框架在 ` sql/database` 包的原生基础上封装了简便的方法，可以直接调用便将数据映射到对应的结构体模型中，用起来非常简单，如下面这几个方法：

` //返回第一条 func (s *DB) First(out interface{}, where...interface{}) *DB //返回最后一条 func (s *DB) Last(out interface{}, where...interface{}) *DB //返回符合条件的内容 func (s *DB) Find(out interface{}, where...interface{}) *DB //返回Count(*)结果 func (s *DB) Count(value interface{}) *DB 复制代码`

示例代码

` //Find方法示例 func find () { var users = make([]*User, 0) db.Model(&User2{}).Find(&users) fmt.Println(users) } //First方法示例 func first () { var user1,user2 User db.First(&user1) fmt.Println(user1) db.First(&user2, "id = ?" ,20) fmt.Println(user2) } //Last方法示例 func last () { var user1,user2 User db.Last(&user1) fmt.Println(user1) db.First(&user2, "id = ?" ,19) fmt.Println(user2) } //Count方法示例 func count () { var count int db.Model(&User{}).Count(&count) fmt.Println(count) } 复制代码`

#### 更新 ####

更新数据可以使用 ` gorm.DB` 的 ` Save()` 或 ` Update()` , ` UpdateColumn()` , ` UpdateColumns()` , ` Updates()` 等方法，后面这四个方法需要与 ` Model()` 方法一起使用。

` func (s *DB) Save(value interface{}) *DB func (s *DB) Model(value interface{}) *DB //下面的方法需要与Model方法一起使用，通过Model方法指定更新数据的条件 func (s *DB) Update(attrs ...interface{}) *DB func (s *DB) UpdateColumn(attrs ...interface{}) *DB func (s *DB) UpdateColumns(values interface{}) *DB func (s *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) *DB 复制代码`

代码示例

` //Save()方法示例 func save (){ u := &User{} db.First(u) u.Email = "test@163.com" db.Save(u) fmt.Println(u) } //Update方法示例 func update () { u := &User{} db.First(u) db.Model(u).Update( "username" , "hello" ) } //Updates方法示例 func updates () { u := &User{} db.First(u) db.Model(&u).Updates(map[string]interface{}{ "username" : "hello2" }) } 复制代码`

#### 删除 ####

使用 ` gorm.DB` 的 ` Delete()` 方法可以很简单地删除满足条件的记录，下面是 ` Delete()` 方法的定义：

` //value如果有主键id，则包含在判断条件内，通过 where 可以指定其他条件 func (s *DB) Delete(value interface{}, where...interface{}) *DB 复制代码`

示例代码

` func delete (){ defer db.Close() u := &User{Id: 16} db.Delete(u)//根据id db.Delete(&User{}, "username = ? " , "test_one" )//根据额外条件删除 } 复制代码`

## 小结 ##

在这篇文章中我们只是讲解使用 ` GROM` 框架如何连接和简单操作数据库而已，其实 ` GROM` 框架还有许多更加高级功能，可以让我们的开发变得更加简洁，在之后的文章中，我们再进行详细讲解吧。

**你的关注，是我写作路上最大的鼓励！**

![](https://user-gold-cdn.xitu.io/2019/5/22/16add5e0a3018453?imageView2/0/w/1280/h/960/ignore-error/1)