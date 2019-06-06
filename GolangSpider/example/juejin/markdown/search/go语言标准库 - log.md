# go语言标准库 - log #

# 一、目录 #

* [一、目录]( #%E4%B8%80%E7%9B%AE%E5%BD%95 )
* [二、log]( #%E4%BA%8Clog )

* [1. 简介]( #1-%E7%AE%80%E4%BB%8B )
* [2. 直接使用标准Logger]( #2-%E7%9B%B4%E6%8E%A5%E4%BD%BF%E7%94%A8%E6%A0%87%E5%87%86logger )
* [3. 使用自定义Logger]( #3-%E4%BD%BF%E7%94%A8%E8%87%AA%E5%AE%9A%E4%B9%89logger )
* [4. 将log写入指定文件中]( #4-%E5%B0%86log%E5%86%99%E5%85%A5%E6%8C%87%E5%AE%9A%E6%96%87%E4%BB%B6%E4%B8%AD )

* [三、log/syslog]( #%E4%BA%8Clogsyslog )

* [1. 简介]( #1-%E7%AE%80%E4%BB%8B-1 )
* [2. TCP监听端口，等待log写入]( #2-tcp%E7%9B%91%E5%90%AC%E7%AB%AF%E5%8F%A3%E7%AD%89%E5%BE%85log%E5%86%99%E5%85%A5 )
* [3. 通过Dial函数，生成Writer实例，向TCP服务写入log]( #3-%E9%80%9A%E8%BF%87dial%E5%87%BD%E6%95%B0%E7%94%9F%E6%88%90writer%E5%AE%9E%E4%BE%8B%E5%90%91tcp%E6%9C%8D%E5%8A%A1%E5%86%99%E5%85%A5log )

# 二、log #

## 1. 简介 ##

该包定义了一个结构体Logger，这个结构体上挂载了以下函数：
1.1 初始化(New)
1.2 设置log风格（SetFlags）
1.3 设置log记录的载体（SetOutput）
1.4 设置log前缀（SetPrefix）
1.5 普通记录log（Print[f|ln]）
1.6 以致命错误退出程序并记录log(Fatal[f|ln])
1.7 以panic形式退出程序并记录log(Panic[f|ln])
1.8 打印对应函数调用栈帧的log(Output)。
为了使用方便，标准库中还预实现了一个标准Logger，方便我们直接函数调用。

## 2. 直接使用标准Logger ##

` package main import ( "log" "os" ) func main () { // 可通过SetOutput设置记录载体（载体只需实现io.Writer接口）。我们这里选择标准输出。 log.SetOutput(os.Stdout) // 设置打印前缀 log.SetPrefix( "My log:" ) // 设置打印风格 /* const ( Ldate = 1 << iota // the date in the local time zone: 2009/01/23 Ltime // the time in the local time zone: 01:23:23 Lmicroseconds // microsecond resolution: 01:23:23.123123. assumes Ltime. Llongfile // full file name and line number: /a/b/c/d.go:23 Lshortfile // final file name element and line number: d.go:23. overrides LUTC // if Ldate or Ltime is set, use UTC rather than the local time zone LstdFlags = Ldate | Ltime // initial values for the standard logger ) */ log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile) // 往载体中输入日志 log.Println( "begin log" ) // 打印风格、前缀 log.Print(log.Flags()) log.Printf( "%s\n" , log.Prefix()) // 打印当前记录日志的地方，和调用main函数的地方 log.Output( 0 , "I'm caller." ) log.Output( 1 , "I'm caller's caller" ) log.Output( 2 , "I'm caller's caller's caller" ) // 以致命错误退出，并记录信息 log.Fatalln( "I'm exit by Fatal." ) } 复制代码`

## 3. 使用自定义Logger ##

` package main import ( "bytes" "fmt" "log" ) func main () { var ( // 定义一个载体 buf bytes.Buffer // logger初始化，参数分别是载体、前缀 和 风格 logger = log.New(&buf, "INFO: " , log.Lshortfile) // 定义Output闭包函数 infof = func (info string ) { logger.Output( 2 , info) } ) infof( "Hello world" ) // 打印载体、前缀、风格 logger.Print(logger.Flags()) logger.Printf( "%s\n" , logger.Prefix()) // 打印载体内容 fmt.Print(&buf) // 也可以写成 fmt.Print(logger.Writer()) // 以panic形式退出 logger.Panic( "I'm exit by Panic" ) } 复制代码`

## 4. 将log写入指定文件中 ##

` package main import ( "log" "os" ) func main () { f, err := os.OpenFile( "test.log" , os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644 ) if nil != err { log.Fatal(err) } defer f.Close() log.SetOutput(f) log.Println( "works" ) } 复制代码`

# 三、log/syslog #

## 1. 简介 ##

可通过unix套接字、UDP、TCP去发送日志信息。

## 2. TCP监听端口，等待log写入 ##

` package main import ( "fmt" "io/ioutil" "net" ) func main () { localAddress, _ := net.ResolveTCPAddr( "tcp4" , "127.0.0.1:8080" ) var tcpListener, err = net.ListenTCP( "tcp" , localAddress) if nil != err { fmt.Println( "监听出错：" , err) return } defer func () { tcpListener.Close() }() fmt.Println( "正在等待连接..." ) var conn, err2 = tcpListener.AcceptTCP() if nil != err2 { fmt.Println( "接受连接失败：" , err2) return } var remoteAddr = conn.RemoteAddr() fmt.Println( "接收到一个链接：" , remoteAddr) fmt.Println( "正在读取消息..." ) var bys, _ = ioutil.ReadAll(conn) fmt.Println( "接受到客户端的消息：" , string (bys)) // conn.Write([]byte("hello, Nice to meet you, my name is Dawn")) conn.Close() } 复制代码`

## 3. 通过Dial函数，生成Writer实例，向TCP服务写入log ##

` package main import ( "log" "log/syslog" ) func main () { sysLog, err := syslog.Dial( "tcp" , "localhost:8080" , syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag" ) if nil != err { log.Fatal(err) } defer sysLog.Close() sysLog.Write([] byte ( "a log message.\n" )) sysLog.Info( "an info message.\n" ) sysLog.Notice( "a notice message.\n" ) sysLog.Emerg( "an emergency message.\n" ) sysLog.Warning( "a warning message.\n" ) sysLog.Err( "a err message.\n" ) sysLog.Debug( "a debug message.\n" ) sysLog.Crit( "a crit message.\n" ) sysLog.Alert( "a alert message.\n" ) } 复制代码`