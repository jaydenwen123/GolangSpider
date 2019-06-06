# 简单聊聊Golang数据库编程的那些事 #

应该说，数据库编程是任何编程语言都有提供的基础功能模块，无论是编程语言内置的支持，还是通过外部库来实现；当然啦，不同编程语言提供的数据库编程API是不尽相同的，而且需要支持的数据库也是多种多样，如常用的 ` MySQL` ， ` SQLServer` , ` Postgres` 等数据库。

抛开其他编程语言不谈，在这篇文章中，我们就来聊一聊Go语言数据库编程的那些事，了解如何使用Go语言提供的标准库，编写通用的操作数据库的代码。

## 数据库连接与驱动 ##

### database/sql和database/sql/driver ###

标准库 ` database/sql` 是Go语言的数据库操作抽象层，具体的数据库操作逻辑实现则由不同的第三方包来做，而标准库 ` database/sql/driver` 提供了这些第三方包实现的标准规范。

所以，Go语言的数据库编程，一般只需要导入标准库 ` database/sql` 包，这个包提供了操作数据库所有必要的结构体、函数与方法，下面是导入这个包的语句：

` import database/sql 复制代码`

Go语言这样做的好处就是，当从一个数据库迁移到另一个数据库时(如 ` SQL Server` 迁移到 ` MySQL` ),则只需要换一个驱动包便可以了。

### Go支持的数据库驱动包 ###

前面我们说Go语言数据操作的由不同第三方包来实现，那么如果我们想要连接MySQL数据库的话，要怎么实现一个这样的包呢？实际上，Go语言标准库 ` database/sql/driver` 定义了实现第三方驱动包的所有接口，我们只导入实现了 ` database/sql/driver` 相关接口驱动包就可以了。

下面是支持Golang的数据库驱动列表：

[github.com/golang/go/w…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgolang%2Fgo%2Fwiki%2FSQLDrivers )

#### 安装第三方驱动包 ####

以MySQL数据库驱动包为例：

` $ go get -u github.com/go-sql-driver/mysql 复制代码`

#### 导入驱动包 ####

` import database/sql import _ "github.com/go-sql-driver/mysql" 复制代码`

### sql.DB结构体 ###

` sql.DB` 结构是 ` sql/database` 包封装的一个数据库操作对象，包含了操作数据库的基本方法。

#### DSN ####

` DSN` 全称为 ` Data Source Name` ，表示数据库连来源，用于定义如何连接数据库，不同数据库的DSN格式是不同的，这取决于数据库驱动的实现，下面是 ` go-sql-driver/sql` 的DSN格式,如下所示：

` //[用户名[:密码]@][协议(数据库服务器地址)]]/数据库名称?参数列表 [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN] 复制代码`

#### 初始化sql.DB ####

` database/sql` 包中有两个函数用于初始化并返回一个 ` *sql.DB` 结构体对象：

` //driverName表示驱动名，如mysql,dataSourceName为上文介绍的DSN func Open(driverName, dataSourceName string) (*DB, error) func OpenDB(c driver.Connector) *DB 复制代码`

下面演示如何使用这两个函数：

一般而言，我们使用 ` Open()` 函数便可初始化并返回一个 ` *sql.DB` 结构体实例，使用 ` Open()` 函数只要传入驱动名称及对应的DSN便可，使用很简单，也很通用，当需要连接不同数据库时，只需要修改驱动名与DSN就可以了。

` import "database/sql" import _ "github.com/go-sql-driver/mysql" //注意前面有_ func open (){ const DRIVER = "mysql" var DSN = "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local" var err error db, err = sql.Open(DRIVER, DSN) if err != nil { panic(err) } if err = db.Ping();err != nil{ panic(err) } } 复制代码`

` OpenDB()` 函数则依赖驱动包实现年 ` sql/database/driver` 包中的 ` Connector` 接口，这种方法并不通用化，不推荐 使用，下面演示使用 ` mysql` 驱动包的 ` driver.Connector` 初始化并返回 ` *sql.DB` 结构体实例：

` import "database/sql" import "github.com/go-sql-driver/mysql" //注意前面没有_ func openConnector () { Connector, err := mysql.NewConnector(&mysql.Config{ User: "root" , Passwd: "123456" , Net: "tcp" , Addr: "localhost:3306" , DBName: "test" , AllowNativePasswords: true , Collation: "utf8_general_ci" , ParseTime: true , Loc:time.Local, }) if err != nil { panic(err) } db = sql.OpenDB(Connector) if err = db.Ping();err != nil{ panic(err) } } 复制代码`

使用前面定义的方法，初始化一个 ` *sql.DB` 指针结构体：

` var db *sql.DB //在init方法初始化`*sql.DB` func init (){ open() //或者 openConnector() } 复制代码`

这里要说一下的是， ` sql.DB` 是 ` sql/database` 包封装的一个结构体，但不表示一个数据库连接对象，实际上，我们可以把 ` sql.DB` 看作一个简单的数据库连接池，我们下面的几个方法设置数据库连接池的相关参数：

` func (db *DB) SetMaxIdleConns(n int)//设置连接池中最大空闲数据库连接数，<=0表示不保留空闲连接，默认值2 func (db *DB) SetMaxOpenConns(n int)//设置连接池最大打开数据库连接数，<=表示不限制打开连接数，默认为0 func (db *DB) SetConnMaxLifetime(d time.Duration)//设置连接超时时间 复制代码`

代码演示

` db.SetMaxOpenConns(100)//设置最多打开100个数据连连接 db.SetMaxIdleConns(0)//设置为0表示 db.SetConnMaxLifetime(time.Second * 5)//5秒超时 复制代码`

## 数据库基本操作 ##

下面我们演示在Go语言中有关数据库的增删改查(CURD)等基本的操作，为此我们创建了一个名为 ` users` 的数据表，其创建的 ` SQL` 语句如下：

` CREATE TABLE users( id INT NOT NULL AUTO_INCREMENT COMMENT 'ID' , username VARCHAR(32) NOT NULL COMMENT '用户名' , moeny INT DEFAULT 0 COMMENT '账户余额' , PRIMARY KEY(id) ); INSERT INTO users VALUES(1, '小明' ,1000); INSERT INTO users VALUES(2, '小红' ,2000); INSERT INTO users VALUES(3, '小刚' ,1400); 复制代码`

### 查询 ###

查询是数据库操作最基本的功能，在Go语言中，可以使用 ` sql.DB` 中的 ` Query()` 或 ` QueryContext()` 方法，这两个方法的定义如下：

` func (db *DB) Query(query string, args ...interface{}) (*Rows, error) func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) 复制代码`

` Query()` 和 ` QueryContext()` 方法返回一个 ` sql.Rows` 结构体，代表一个查询结果集， ` sql.Rows` 的定义及其所包含的方法如下：

` type Rows struct { //contains filtered or unexported fields } func (rs *Rows) Close() error //关闭结果集 func (rs *Rows) ColumnTypes() ([]*ColumnType, error)//返回数据表的列类型 func (rs *Rows) Columns() ([]string, error)//返回数据表列的名称 func (rs *Rows) Err() error//错误集 func (rs *Rows) Next() bool//游标，下一行 func (rs *Rows) NextResultSet() bool func (rs *Rows) Scan(dest ...interface{}) error //扫描结构体 复制代码`

使用 ` sql.Rows` 的 ` Next()` 和 ` Scan` 方法，但可以遍历返回的结果集，下面是示例代码：

` func query () { selectText := "SELECT * FROM users WHERE id = ?" rows, _ := db.Query(selectText, 2) defer rows.Close() for rows. Next () { var ( id int username string money int ) _ = rows.Scan(&id, &username,&money) fmt.Println(id, username,money) } } 复制代码`

还可以用 ` sql.DB` 中的 ` QueryRow()` 或 ` QueryRowContext()` 方法，这两个方法的定义如下所示：

` func (db *DB) QueryRow(query string, args ...interface{}) *Row func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row 复制代码`

` QueryRow` 和 ` QueryRowContext` 返回一个 ` sql.Row` 结构体，代表数据表的一行， ` sql.Row` 的定义如下，可以看到 ` sql.Row` 结构体只有一个 ` Scan()` 方法，用于扫描 ` sql.Row` 结构体中的数据。

` type Row struct{ } func (r *Row) Scan(dest ...interface{}) error 复制代码`

代码演示

` func queryRow (){ selectText := "SELECT * FROM users WHERE id = ?" row := db.QueryRow(selectText, 2) var ( id int username string money int ) _ = row.Scan(&id, &username,&money) fmt.Println(id, username,money) } 复制代码`

另外，使用 ` sql.DB` 中的 ` Prepare()` 或 ` PrepareContext()` 方法，可以返回一个 ` sql.Stmt` 结构体`。

> 
> 
> 
> 注意：sql.Stmt结构体会先把在 ` Prepare()` 或 ` PrepareContext()` 定义的SQL语句发给数据库执行，再将SQL语句中需要的参数发给数据库，再返回处理结果。
> 
> 
> 

` func (db *DB) Prepare(query string) (*Stmt, error) func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error) 复制代码`

` sql.Stmt` 提交了与 ` sql.DB` 用于查询并返回结果集的方法，下面请看示例：

` func queryStmt (){ stmt,err := db.Prepare( "SELECT * FROM users WHERE id = ?" ) if err != nil{ return } defer stmt.Close() rows,err := stmt.Query(2) defer rows.Close() for rows. Next () { var ( id int username string money int ) _ = rows.Scan(&id, &username,&money) fmt.Println(id, username,money) } } 复制代码`

### 添加 ###

添加数据库记录，可以使用 ` sql.DB` 中的 ` Exec()` 或 ` ExecContext()` 方法,这两个方法的定义如下：

` func (db *DB) Exec(query string, args ...interface{}) (Result, error) func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) 复制代码`

代码示例：

` func insert (){ insertText := "INSERT INTO users values(?,?,?)" rs,err := db.Exec(insertText,4, "juejin" ,1000) if err != nil{ fmt.Println(err) return } if id,_ := rs.LastInsertId();id > 0 { fmt.Println( "插入成功" ) } /*也可以这样判断是否插入成功 if n,_ := rs.RowsAffected();n > 0 { fmt.Println( "插入成功" ) } */ } 复制代码`

` Exec()` 或 ` ExecContext()` 方法的第一个返回值为一个实现了 ` sql.Result` 接口的类型， ` sql.Result` 的定义如下：

> 
> 
> 
> 注意 ` LastInsertId()` 方法只有在使用INSERT语句且数据表有自增id时才有返回自增id值，否则返回0。
> 
> 

` type Result interface { LastInsertId() (int64, error)//使用insert向数据插入记录，数据表有自增id时，该函数有返回值 RowsAffected() (int64, error)//表示影响的数据表行数 } 复制代码`

我们可以用 ` sql.Result` 中的 ` LastInsertId()` 方法或 ` RowsAffected()` 来判断 ` SQL` 语句是否执行成功。

除了使用 ` sql.DB` 中的 ` Exec()` 和 ` ExecContext()` 方法外，也可以使用 ` Prepare()` 或 ` PrepareContext()` 返回 ` sql.Stmt` 结构体，再通过 ` sql.Stmt` 中的 ` Exec()` 方法向数据表写入数据。

使用 ` sql.Stmt` 向数据表写入数据的演示：

` func insertStmt (){ stmt,err := db.Prepare( "INSERT INTO users VALUES(?,?,?)" ) defer stmt.Close() if err != nil{ return } rs,err := stmt.Exec(5, "juejin" ,1000) if id,_ := rs.LastInsertId(); id > 0 { fmt.Println( "插入成功" ) } } 复制代码`
> 
> 
> 
> 
> 注意，使用 ` sql.Stmt` 中的 ` Exec()` 或 ` ExecContext()` 执行 ` SQL` 对更新和删除语句同样适合，下面讲更新和删除时不再演示。
> 
> 
> 

### 更新 ###

与往数据表里添加数据一样，可以使用 ` sql.DB` 的 ` Exec()` 或 ` ExecContext()` 方法,不过，使用数据库 ` UPDATE` 语句更新数据时，我们只能通过 ` sql.Result` 结构体中的 ` RowsAffected()` 方法来判断影响的数据行数，进而判断是否执行成功。

` func update () { updateText := "UPDATE users SET username = ? WHERE id = ?" rs,err := db.Exec(updateText, "database" ,2) if err != nil{ fmt.Println(err) return } if n,_ := rs.RowsAffected();n > 0 { fmt.Println( "更新成功" ) } } 复制代码`

### 删除 ###

使用 ` DELETE` 语句删除数据表记录的操作与上面的更新语句是一样的，请看下面的演示：

` func del () { delText := "DELETE FROM users WHERE id = ?" rs,err := db.Exec(delText,1) if err != nil{ fmt.Println(err) return } fmt.Println(rs.RowsAffected()) } 复制代码`

## 事务 ##

在前面的示例中，我们都没有开启事务，如果没有开启事务，那么默认会把提交的每一条 ` SQL` 语句都当作一个事务来处理，如果多条语句一起执行，当其中某个语句执行错误，则前面已经执行的 ` SQL` 语句无法回滚。

对于一些要求比较严格的业务逻辑来说(如订单付款、用户转账等)，应该在同一个事务提交多条 ` SQL` 语句，避免发生执行出错无法回滚事务的情况。

#### 开启事务 ####

如何开启一个新的事务？可以使用 ` sql.DB` 结构体中的 ` Begin()` 或 ` BeginTx()` 方法，这两个方法的定义如下：

` func (db *DB) Begin() (*Tx, error) func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) 复制代码`

` BeginTx()` 方法的第二个参数为 ` TxOptions` ，其定义如下：

` type TxOptions struct { // Isolation is the transaction isolation level. // If zero, the driver or database 's default level is used. Isolation IsolationLevel ReadOnly bool } 复制代码`

` TxOptions` 的 ` Isolation` 字段用于定义事务的隔离级别，其类型为 ` IsolationLevel` ， ` Ioslation` 的取值范围可以为如下常量：

` const ( LevelDefault IsolationLevel = iota LevelReadUncommitted LevelReadCommitted LevelWriteCommitted LevelRepeatableRead LevelSnapshot LevelSerializable LevelLinearizable ) 复制代码`

` Begin()` 和 ` BeginTxt()` 方法返回一个 ` sql.Tx` 结构体，使用 ` sql.Tx` 对数据库进行操作，会在同一个事务中提交，下面演示代码：

` tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}) //Begin方法实际上是调用BeginTx()方法,db.BeginTx(context.Background(), nil) tx, err := db.Begin() 复制代码`

#### sql.Tx支持的基本操作 ####

下面是 ` sql.Tx` 结构体中的基本操作方法，使用方式与我们前面演示的例子一样

` func (tx *Tx) Exec(query string, args ...interface{}) (Result, error) func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error) func (tx *Tx) Query(query string, args ...interface{}) (*Rows, error) func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) func (tx *Tx) QueryRow(query string, args ...interface{}) *Row func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row 复制代码`

#### 事务提交 ####

当使用 ` sql.Tx` 的操作方式操作数据后，需要我们使用 ` sql.Tx` 的 ` Commit()` 方法显式地提交事务，如果出错，则可以使用 ` sql.Tx` 中的 ` Rollback()` 方法回滚事务，保持数据的一致性，下面是这两个方法的定义：

` func (tx *Tx) Commit() error func (tx *Tx) Rollback() error 复制代码`

#### 预编译 ####

` sql.Tx` 结构体中的 ` Stmt()` 和 ` StmtContext()` 可以将 ` sql.Stmt` 封装为支持事务的 ` sql.Stmt` 结构体并返回,这两个方法的定义如下：

` func (tx *Tx) Stmt(stmt *Stmt) *Stmt func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt 复制代码`

而使用 ` sql.Tx` 中的 ` Prepare()` 和 ` PrepareContext()` 方法则可以直接返回一个支持事务的 ` sql.Stmt` 结构体

` func (tx *Tx) Prepare(query string) (*Stmt, error) func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error) 复制代码`

#### 示例 ####

` //修改 func txUpdate (){ tx,_ := db.Begin() rs,err := tx.Exec( "UPDATE users SET username = ? WHERE id = ?" , "sssss" ,2) if err != nil{ panic(err) } err = tx.Commit() if err != nil{ panic(err) } if n,_ := rs.RowsAffected();n > 0{ fmt.Println( "成功" ) } } //使用Stmt修改 func txStmt (){ tx,err := db.Begin() if err != nil{ panic(err) } stmt,err := db.Prepare( "UPDATE users SET username = ? WHERE id = ?" ) stmtTx := tx.Stmt(stmt) defer stmtTx.Close() rs,_ := stmtTx.Exec( "test" ,2) _ = tx.Commit() if n,_ := rs.RowsAffected();n > 0{ fmt.Println( "成功" ) } } 复制代码`

## 相关ORM框架 ##

前面我们介绍是Go语言原生对数据库编程的支持，不过，更方便的是，我们可以直接使用一些开源的 ` ORM` ( ` Object Relational Mapping` )框架, ` ORM` 框架可以封装了底层的 ` SQL` 语句，并直接映射为 ` Struct` , ` Map` 等数据类型，省去我们直接写 ` SQL` 语句的工作，非常简单方便。

下面几个比较常用的ORM框架：

### GORM ###

` GORM` 是一个非常完善的ORM框架，除了基本增加改查的支持，也支持关联包含一个，包含多个，属于，多对多的数据表，另外也可以在创建/保存/更新/删除/查找之前或之后写钩子回调，同时也支持事务。

GORM目前支持数据库驱动 ` MySQL` 、 ` SQLite3` 、 ` SQL Server` 、 ` Postgres` 。

### Xorm ###

` Xorm` 也是一个简单又强大的ORM框架，其功能也 ` GORM` 是类似的，不过支持的数据库驱动比 ` GORM` 多一些，支持 ` MySQL` , ` SQL Server` , ` SQLite3` , ` Postgres` , ` MyMysql` , ` Tidb` , ` Oracle` 等数据库驱动。

### Beego ORM ###

` Beego ORM` 是国人开发的Web框架 ` Beego` 中的一个模块，虽然是 ` Beego` 的一个模块，但可以独立使用，不过目前 ` Beego ORM` 只支持 ` MySQL` 、 ` SQLite3` , ` Postgres` 等数据库驱动。

> 
> 
> 
> 除了上面我们介绍的三个ORM框架，其实还很多很好的ORM的框架，大家有空可以看看。
> 
> 

## 小结 ##

Go语言将对数据库的操作抽象并封装在 ` sql/database` 包中，为我们操作不同数据库提供统一的API，非常实用，我们在这篇文章中讲解了 ` sql/database` 包中的 ` sql.DB` , ` sql.Rows` , ` sql.Stmt` , ` sql.Tx` 等结构体的使用，相信你通过上面的示例，也一定能掌握Go语言操作数据库的方法。

![](https://user-gold-cdn.xitu.io/2019/5/21/16ad9f1d36b8eee8?imageView2/0/w/1280/h/960/ignore-error/1)