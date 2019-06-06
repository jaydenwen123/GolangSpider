# Micro 总结 #

# 常用命令 #

#### 1.创建服务 ####

pd_log 为服务名字 ckp 为项目目录

` micro new ckp/srv/pd_log --type=srv 复制代码`

注: 需先安装micro 在 micro/micro 执行go install 即可

#### 2.编辑服务接口 ####

编写Pdlog接口Insert方法 InsertRequest 为请求参数结构体(字段为string Id) InsertResponse 为返回结构体(字段为 string Msg) 返回参数

` service Pdlog { rpc Insert(InsertRequest) returns (InsertResponse) {} } message InsertRequest { string Id = 1; } message InsertResponse { string Msg = 1; } 复制代码`

#### 3.生成proto (在文件当前目录下生成) ####

protoc --proto_path=. --go_out=. --micro_out=. pd_log.proto

注:需要安装protoc 1.github上下载一个cpp包： [github.com/google/prot…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgoogle%2Fprotobuf%2Freleases ) [github.com/protocolbuf…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprotocolbuffers%2Fprotobuf%2Freleases%2Fdownload%2Fv3.7.1%2Fprotoc-3.7.1-osx-x86_64.zip ) 这个是mac系统的安装包 解压放到环境变量bin目录下即可 2.protoc-gen-go go get -u github.com/golang/protobuf/protoc-gen-go 3.安装protoc-gen-micro go get github.com/micro/protoc-gen-micro

#### 4.在 handler 下编写对应服务的结构体 并设置Client ####

` type PdLog struct{ Client client.Client } 复制代码`

#### 5. 在 main.go 下注册handler ####

初始化服务是加上心跳参数

` micro.RegisterTTL(time.Second*30), micro.RegisterInterval(time.Second*15), 复制代码` ` service := micro.NewService( micro.Name( "go.micro.srv.pd_log" ), micro.Version( "latest" ), micro.RegisterTTL(time.Second*30), micro.RegisterInterval(time.Second*15), ) 复制代码` ` pd_lg.RegisterPdlogHandler(service.Server(), new(handler.PdLog)) 复制代码`

#### 6. 在handler 下实现Insert方法 ####

` func (p *PdLog) Insert(ctx context.Context, req *pd_lg.InsertRequest, rsp *pd_lg.InsertResponse) error { //实现业务 return nil } 复制代码`

#### 7.其他命令 ####

` protoc --proto_path=$GOPATH/src --go_out=. --micro_out=. ckp/api/homepage/proto/homepage/homepage.proto`

homepage.proto 文件可以引用其他目录下proto文件(不建议使用) 执行命令需要在src目录下执行 micro web 启动web页面 可以测试接口

` micro api --handler=rpc 启动api`

API的Handler处理器都是用来接收Http请求，然后根据请求类型进行处理，或向前转发，或触发事件，为了方便，handler的注册名都能匹配http.Handler字样。