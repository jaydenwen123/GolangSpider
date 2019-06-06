# Go 实现简易 RPC 框架 #

本文旨在讲述 RPC 框架设计中的几个核心问题及其解决方法，并基于 Golang 反射技术，构建了一个简易的 RPC 框架。

项目地址： [Tiny-RPC]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJiahonzheng%2FTiny-RPC )

## RPC ##

[RPC]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FRemote_procedure_call ) （Remote Procedure Call），即远程过程调用，可以理解成，服务 A 想调用不在同一内存空间的服务 B 的函数，由于不在一个内存空间，不能直接调用，需要通过网络来表达调用的语义和传达调用的数据。

## 服务端 ##

RPC 服务端需要解决 2 个问题：

* 由于客户端传送的是 RPC 函数名，服务端如何维护 函数名 与 函数实体 之间的映射
* 服务端如何根据 函数名 实现对应的 函数实体 的调用

### 核心流程 ###

* 维护函数名到函数的映射
* 在接收到来自客户端的函数名、参数列表后，解析参数列表为反射值，并执行对应函数
* 对函数执行结果进行编码，并返回给客户端

### 方法注册 ###

服务端需要维护 RPC 函数名到 RPC 函数实体的映射，我们可以使用 ` map` 数据结构来维护映射关系。

` type Server struct { addr string funcs map [ string ]reflect.Value } // Register a method via name func (s *Server) Register (name string , f interface {}) { if _, ok := s.funcs[name]; ok { return } s.funcs[name] = reflect.ValueOf(f) } 复制代码`

### 执行调用 ###

一般来说，客户端在调用 RPC 时，会将 函数名 和 参数列表 作为请求数据，发送给服务端。

由于我们使用了 ` map[string]reflect.Value` 来维护函数名与函数实体之间的映射，则我们可以通过 ` Value.Call()` 来调用与函数名相对应的函数。

代码地址：https://play.golang.org/p/jaPHviCbe5K

` package main import ( "fmt" "reflect" ) func main () { // Register methods funcs := make ( map [ string ]reflect.Value) funcs[ "add" ] = reflect.ValueOf(add) // When receives client's request req := []reflect.Value{reflect.ValueOf( 1 ), reflect.ValueOf( 2 )} vals := funcs[ "add" ].Call(req) var rsp [] interface {} for _, val := range vals { rsp = append (rsp, val.Interface()) } fmt.Println(rsp) } func add (a, b int ) ( int , error) { return a + b, nil } 复制代码`

### 具体实现 ###

由于篇幅的限制，此处没有贴出服务端实现的具体代码，细节请查看 [项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJiahonzheng%2FTiny-RPC ) 。

## 客户端 ##

RPC 客户端需要解决 1 个问题：

* 由于函数的具体实现在服务端，客户端只有函数的原型，客户端如何通过 函数原型 调用其 函数实体

### 核心流程 ###

* 对调用者传入的函数参数进行编码，并传送给服务端
* 对服务端响应数据进行解码，并返回给调用者

### 生成调用 ###

我们可以通过 ` reflect.MakeFunc` 为指定的函数原型绑定一个函数实体。

代码地址： https://play.golang.org/p/AaedlW9U-6n

` package main import ( "fmt" "reflect" ) func main () { add := func (args []reflect.Value) [] reflect. Value { result := args[ 0 ].Interface().( int ) + args[ 1 ].Interface().( int ) return []reflect.Value{reflect.ValueOf(result)} } var addptr func ( int , int ) int container := reflect. ValueOf (&addptr). Elem () v := reflect. MakeFunc (container.Type() , add ) container. Set (v) fmt. Println (addptr(1, 2) ) } 复制代码`

### 具体实现 ###

由于篇幅的限制，此处没有贴出客户端实现的具体代码，细节请查看 [项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJiahonzheng%2FTiny-RPC ) 。

## 数据传输格式 ##

我们需要定义服务端与客户端交互的数据格式。

` type Data struct { Name string // service name Args [] interface {} // request's or response's body except error Err string // remote server error } 复制代码`

与交互数据相对应的编码与解码函数。

` func encode (data Data) ([] byte , error) { var buf bytes.Buffer encoder := gob.NewEncoder(&buf) if err := encoder.Encode(data); err != nil { return nil , err } return buf.Bytes(), nil } func decode (b [] byte ) (Data, error) { buf := bytes.NewBuffer(b) decoder := gob.NewDecoder(buf) var data Data if err := decoder.Decode(&data); err != nil { return Data{}, err } return data, nil } 复制代码`

同时，我们需要定义简单的 [TLV]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FType-length-value ) 协议（固定长度消息头 + 变长消息体），规范数据的传输。

` // Transport struct type Transport struct { conn net.Conn } // NewTransport creates a transport func NewTransport (conn net.Conn) * Transport { return &Transport{conn} } // Send data func (t *Transport) Send (req Data) error { b, err := encode(req) // Encode req into bytes if err != nil { return err } buf := make ([] byte , 4 + len (b)) binary.BigEndian.PutUint32(buf[: 4 ], uint32 ( len (b))) // Set Header field copy (buf[ 4 :], b) // Set Data field _, err = t.conn.Write(buf) return err } // Receive data func (t *Transport) Receive () (Data, error) { header := make ([] byte , 4 ) _, err := io.ReadFull(t.conn, header) if err != nil { return Data{}, err } dataLen := binary.BigEndian.Uint32(header) // Read Header filed data := make ([] byte , dataLen) // Read Data Field _, err = io.ReadFull(t.conn, data) if err != nil { return Data{}, err } rsp, err := decode(data) // Decode rsp from bytes return rsp, err } 复制代码`

## 相关资料 ##

* 项目地址： [Tiny-RPC]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJiahonzheng%2FTiny-RPC )
* [go rpc 源码分析]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000013532622 )