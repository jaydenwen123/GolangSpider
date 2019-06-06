# Go 语言编程规范 #

#### 1. gofmt 命令 ####

大部分的格式问题可以通过 gofmt 来解决，gofmt 自动格式化代码，保证所有的 go 代码与官方推荐的格式保持一致，所有格式有关问题，都以gofmt的结果为准。所以，建议在提交代码库之前先运行一下这个命令。

#### 2. 行长 ####

一行最长不超过80个字符，超过的使用换行展示，尽量保持格式优雅。

#### 3. 注释 ####

在编码阶段应该同步写好 变量、函数、包 的注释，最后可以利用 godoc 命令导出文档。注释必须是完整的句子，句子的结尾应该用句号作为结尾（英文句号）。注释推荐用英文，可以在写代码过程中锻炼英文的阅读和书写能力。而且用英文不会出现各种编码的问题。

每个包都应该有一个包注释，一个位于 package 子句之前的块注释或行注释。包如果有多个 go 文件，只需要出现在一个 go 文件中即可。

` // ping包实现了常用的ping相关的函数 package ping 复制代码`

#### 4. 命名 ####

需要注释来补充的命名就不算是好命名。 使用可搜索的名称：单字母名称和数字常量很难从一大堆文字中搜索出来。 单字母名称仅适用于短方法中的本地变量，名称长短应与其作用域相对应。 若变量或常量可能在代码中多处使用，则应赋其以便于搜索的名称。

做有意义的区分：Product 和 ProductInfo 和 ProductData 没有区别， NameString 和 Name 没有区别，要区分名称，就要以读者能鉴别不同之处的方式来区分 。

` 函数命名规则：驼峰式命名，名字可以长但是得把功能，必要的参数描述清楚， 函数名应当是动词或动词短语，如 postPayment、deletePage、save。 并依 Javabean 标准加上 get、set、is前缀。 例如：xxx + With + 需要的参数名 + And + 需要的参数名 + ….. 结构体命名规则：结构体名应该是名词或名词短语， 如 Custome、WikiPage、Account、AddressParser，避免使用Manager、Processor、Data、Info、 这样的类名，类名不应当是动词。 包名命名规则：包名应该为小写单词，不要使用下划线或者混合大小写。 接口命名规则：单个函数的接口名以”er”作为后缀，如 Reader,Writer。接口的实现则去掉“er”。 复制代码 复制代码` ` type Reader interface { Read(p []byte) (n int, err error) } // 多个函数接口 type WriteFlusher interface { Write([]byte) (int, error) Flush() error } 复制代码`

#### 5. 常量 ####

常量均需使用全部大写字母组成，并使用下划线分词：

` const APP_VER = "1.0" // 如果是枚举类型的常量，需要先创建相应类型： type Scheme string const ( HTTP Scheme = "http" HTTPS Scheme = "https" ) 复制代码`

#### 6. 变量 ####

变量命名基本上遵循相应的英文表达或简写，在相对简单的环境（对象数量少、针对性强）中， 可以将一些名称由完整单词简写为单个字母，例如：

` user 可以简写为 u userID 可以简写 uid // 若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头： var isExist bool var hasConflict bool var canManage bool var allowGitHook bool 复制代码`

#### 7. 变量命名惯例 ####

变量名称一般遵循驼峰法，但遇到特有名词时，需要遵循以下规则：

如果变量为私有，且特有名词为首个单词，则使用小写，

如：apiClient其它情况都应当使用该名词原有的写法，

如 APIClient、repoID、UserID 错误示例：UrlArray，应该写成 urlArray 或者 URLArray

` //下面列举了一些常见的特有名词： "API" ， "ASCII" ， "CPU" ， "CSS" ， "DNS" ， "EOF" ，GUID "，" HTML "，" HTTP "， " HTTPS "，" ID "," IP "，" JSON "，" LHS "，" QPS "，" RAM "，" RHS " " RPC ", " SLA "， " SMTP "，" SSH "，" TLS "，" TTL "," UI "，" UID "，" UUID "，" URI "，" URL "， " UTF8 "， " VM "，" XML "，" XSRF "，" XSS " 复制代码`

#### 8. struct规范 ####

` struct申明和初始化格式采用多行，定义如下： type User struct{ Username string Email string } 初始化如下： u := User{ Username: "test", Email: "test@gmail.com", } 复制代码`

#### 9. panic ####

尽量不要使用panic，除非你知道你在做什么

#### 10. import ####

对 import 的包进行分组管理，用换行符分割，而且标准库作为分组的第一组。 如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包， 建议采用如下方式进行组织你的包

` 复制代码 package main import ( "fmt" "os" "kmg/a" "kmg/b" "code.google.com/a" "github.com/b" ) 复制代码 goimports 会自动帮你格式化 复制代码`

#### 11. 参数传递 ####

对于少量数据，不要传递指针 对于大量数据的 struct 可以考虑使用指针 传入的参数是 map，slice，chan 不要传递指针， 因为 map，slice，chan 是引用类型，不需要传递指针的指针

#### 12. 单元测试 ####

` 单元测试文件名命名规范： example_test.go 测试用例的函数名称必须以 Test 开头，例如： func TestExample 复制代码`