# 当 Go 遇上了 Lua #

在 GitHub 玩耍时，偶然发现了 [gopher-lua]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyuin%2Fgopher-lua ) ，这是一个纯 Golang 实现的 Lua 虚拟机。我们知道 Golang 是静态语言，而 Lua 是动态语言，Golang 的性能和效率各语言中表现得非常不错，但在动态能力上，肯定是无法与 Lua 相比。那么如果我们能够将二者结合起来，就能综合二者各自的长处了（手动滑稽。

在项目 [Wiki]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyuin%2Fgopher-lua%2Fwiki%2FBenchmarks ) 中，我们可以知道 gopher-lua 的执行效率和性能仅比 C 实现的 bindings 差。因此从性能方面考虑，这应该是一款非常不错的虚拟机方案。

## Hello World ##

这里给出了一个简单的 Hello World 程序。我们先是新建了一个虚拟机，随后对其进行了 ` DoString(...)` 解释执行 lua 代码的操作，最后将虚拟机关闭。执行程序，我们将在命令行看到 "Hello World" 的字符串。

` package main import ( "github.com/yuin/gopher-lua" ) func main () { l := lua.NewState() defer l.Close() if err := l.DoString( `print("Hello World")` ); err != nil { panic (err) } } // Hello World 复制代码`

## 提前编译 ##

在查看上述 ` DoString(...)` 方法的调用链后，我们发现每执行一次 ` DoString(...)` 或 ` DoFile(...)` ，都会各执行一次 parse 和 compile 。

` func (ls *LState) DoString (source string ) error { if fn, err := ls.LoadString(source); err != nil { return err } else { ls.Push(fn) return ls.PCall( 0 , MultRet, nil ) } } func (ls *LState) LoadString (source string ) (*LFunction, error) { return ls.Load(strings.NewReader(source), "<string>" ) } func (ls *LState) Load (reader io.Reader, name string ) (*LFunction, error) { chunk, err := parse.Parse(reader, name) // ... proto, err := Compile(chunk, name) // ... } 复制代码`

从这一点考虑，在同份 Lua 代码将被执行多次（如 **在 http server 中，每次请求将执行相同 Lua 代码** ）的场景下，如果我们能够对代码进行提前编译，那么应该能够减少 parse 和 compile 的开销（如果这属于 hotpath 代码）。根据 Benchmark 结果，提前编译确实能够减少不必要的开销。

` package glua_test import ( "bufio" "os" "strings" lua "github.com/yuin/gopher-lua" "github.com/yuin/gopher-lua/parse" ) // 编译 lua 代码字段 func CompileString (source string ) (*lua.FunctionProto, error) { reader := strings.NewReader(source) chunk, err := parse.Parse(reader, source) if err != nil { return nil , err } proto, err := lua.Compile(chunk, source) if err != nil { return nil , err } return proto, nil } // 编译 lua 代码文件 func CompileFile (filePath string ) (*lua.FunctionProto, error) { file, err := os.Open(filePath) defer file.Close() if err != nil { return nil , err } reader := bufio.NewReader(file) chunk, err := parse.Parse(reader, filePath) if err != nil { return nil , err } proto, err := lua.Compile(chunk, filePath) if err != nil { return nil , err } return proto, nil } func BenchmarkRunWithoutPreCompiling (b *testing.B) { l := lua.NewState() for i := 0 ; i < b.N; i++ { _ = l.DoString( `a = 1 + 1` ) } l.Close() } func BenchmarkRunWithPreCompiling (b *testing.B) { l := lua.NewState() proto, _ := CompileString( `a = 1 + 1` ) lfunc := l.NewFunctionFromProto(proto) for i := 0 ; i < b.N; i++ { l.Push(lfunc) _ = l.PCall( 0 , lua.MultRet, nil ) } l.Close() } // goos: darwin // goarch: amd64 // pkg: glua // BenchmarkRunWithoutPreCompiling-8 100000 19392 ns/op 85626 B/op 67 allocs/op // BenchmarkRunWithPreCompiling-8 1000000 1162 ns/op 2752 B/op 8 allocs/op // PASS // ok glua 3.328s 复制代码`

## 虚拟机实例池 ##

在同份 Lua 代码被执行的场景下，除了可使用提前编译优化性能外，我们还可以引入虚拟机实例池。

因为新建一个 Lua 虚拟机会涉及到大量的内存分配操作，如果采用每次运行都重新创建和销毁的方式的话，将消耗大量的资源。引入虚拟机实例池，能够复用虚拟机，减少不必要的开销。

` func BenchmarkRunWithoutPool (b *testing.B) { for i := 0 ; i < b.N; i++ { l := lua.NewState() _ = l.DoString( `a = 1 + 1` ) l.Close() } } func BenchmarkRunWithPool (b *testing.B) { pool := newVMPool( nil , 100 ) for i := 0 ; i < b.N; i++ { l := pool.get() _ = l.DoString( `a = 1 + 1` ) pool.put(l) } } // goos: darwin // goarch: amd64 // pkg: glua // BenchmarkRunWithoutPool-8 10000 129557 ns/op 262599 B/op 826 allocs/op // BenchmarkRunWithPool-8 100000 19320 ns/op 85626 B/op 67 allocs/op // PASS // ok glua 3.467s 复制代码`

Benchmark 结果显示，虚拟机实例池的确能够减少很多内存分配操作。

下面给出了 README 提供的实例池实现，但注意到该实现在初始状态时， **并未创建足够多的虚拟机实例** （初始时，实例数为0），以及存在 **slice 的动态扩容问题** ，这都是值得改进的地方。

` type lStatePool struct { m sync.Mutex saved []*lua.LState } func (pl *lStatePool) Get () * lua. LState { pl.m.Lock() defer pl.m.Unlock() n := len (pl.saved) if n == 0 { return pl.New() } x := pl.saved[n -1 ] pl.saved = pl.saved[ 0 : n -1 ] return x } func (pl *lStatePool) New () * lua. LState { L := lua.NewState() // setting the L up here. // load scripts, set global variables, share channels, etc... return L } func (pl *lStatePool) Put (L *lua.LState) { pl.m.Lock() defer pl.m.Unlock() pl.saved = append (pl.saved, L) } func (pl *lStatePool) Shutdown () { for _, L := range pl.saved { L.Close() } } // Global LState pool var luaPool = &lStatePool{ saved: make ([]*lua.LState, 0 , 4 ), } 复制代码`

## 模块调用 ##

gopher-lua 支持 Lua 调用 Go 模块，个人觉得，这是一个非常令人振奋的功能点，因为在 Golang 程序开发中，我们可能设计出许多常用的模块，这种跨语言调用的机制，使得我们能够对代码、工具进行复用。

当然，除此之外，也存在 Go 调用 Lua 模块，但个人感觉后者是没啥必要的，所以在这里并没有涉及后者的内容。

` package main import ( "fmt" lua "github.com/yuin/gopher-lua" ) const source = ` local m = require("gomodule") m.goFunc() print(m.name) ` func main () { L := lua.NewState() defer L.Close() L.PreloadModule( "gomodule" , load) if err := L.DoString(source); err != nil { panic (err) } } func load (L *lua.LState) int { mod := L.SetFuncs(L.NewTable(), exports) L.SetField(mod, "name" , lua.LString( "gomodule" )) L.Push(mod) return 1 } var exports = map [ string ]lua.LGFunction{ "goFunc" : goFunc, } func goFunc (L *lua.LState) int { fmt.Println( "golang" ) return 0 } // golang // gomodule 复制代码`

## 变量污染 ##

当我们使用实例池减少开销时，会引入另一个棘手的问题：由于同一个虚拟机可能会被多次执行同样的 Lua 代码，进而变动了其中的全局变量。如果代码逻辑依赖于全局变量，那么可能会出现难以预测的运行结果（这有点数据库隔离性中的“不可重复读”的味道）。

### 全局变量 ###

如果我们需要限制 Lua 代码只能使用局部变量，那么站在这个出发点上，我们需要对全局变量做出限制。那问题来了，该如何实现呢？

我们知道，Lua 是编译成字节码，再被解释执行的。那么，我们可以在编译字节码的阶段中，对全局变量的使用作出限制。在查阅完 Lua 虚拟机指令后，发现涉及到全局变量的指令有两条：GETGLOBAL（Opcode 5）和 SETGLOBAL（Opcode 7）。

到这里，已经有了大致的思路：我们可通过判断字节码是否含有 GETGLOBAL 和 SETGLOBAL 进而限制代码的全局变量的使用。至于字节码的获取，可通过调用 ` CompileString(...)` 和 ` CompileFile(...)` ，得到 Lua 代码的 FunctionProto ，而其中的 Code 属性即为字节码 slice，类型为 ` []uint32` 。

在虚拟机 [实现代码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyuin%2Fgopher-lua%2Fblob%2Fmaster%2Fopcode.go ) 中，我们可以找到一个根据字节码输出对应 OpCode 的工具函数。

` // 获取对应指令的 OpCode func opGetOpCode (inst uint32 ) int { return int (inst >> 26 ) } 复制代码`

有了这个工具函数，我们即可实现对全局变量的检查。

` package main // ... func CheckGlobal (proto *lua.FunctionProto) error { for _, code := range proto.Code { switch opGetOpCode(code) { case lua.OP_GETGLOBAL: return errors.New( "not allow to access global" ) case lua.OP_SETGLOBAL: return errors.New( "not allow to set global" ) } } // 对嵌套函数进行全局变量的检查 for _, nestedProto := range proto.FunctionPrototypes { if err := CheckGlobal(nestedProto); err != nil { return err } } return nil } func TestCheckGetGlobal (t *testing.T) { l := lua.NewState() proto, _ := CompileString( `print(_G)` ) if err := CheckGlobal(proto); err == nil { t.Fail() } l.Close() } func TestCheckSetGlobal (t *testing.T) { l := lua.NewState() proto, _ := CompileString( `_G = {}` ) if err := CheckGlobal(proto); err == nil { t.Fail() } l.Close() } 复制代码`

### 模块 ###

除变量可能被污染外，导入的 Go 模块也有可能在运行期间被篡改。因此，我们需要一种机制，确保导入到虚拟机的模块不被篡改，即导入的对象是 **只读** 的。

在查阅 [相关博客]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fnellson%2Farticle%2Fdetails%2F10928063 ) 后，我们可以对 Table 的 ` __newindex` 方法的修改，将模块设置为只读模式。

` package main import ( "fmt" "github.com/yuin/gopher-lua" ) // 设置表为只读 func SetReadOnly (l *lua.LState, table *lua.LTable) * lua. LUserData { ud := l.NewUserData() mt := l.NewTable() // 设置表中域的指向为 table l.SetField(mt, "__index" , table) // 限制对表的更新操作 l.SetField(mt, "__newindex" , l.NewFunction( func (state *lua.LState) int { state.RaiseError( "not allow to modify table" ) return 0 })) ud.Metatable = mt return ud } func load (l *lua.LState) int { mod := l.SetFuncs(l.NewTable(), exports) l.SetField(mod, "name" , lua.LString( "gomodule" )) // 设置只读 l.Push(SetReadOnly(l, mod)) return 1 } var exports = map [ string ]lua.LGFunction{ "goFunc" : goFunc, } func goFunc (l *lua.LState) int { fmt.Println( "golang" ) return 0 } func main () { l := lua.NewState() l.PreloadModule( "gomodule" , load) // 尝试修改导入的模块 if err := l.DoString( `local m = require("gomodule");m.name = "hello world"` ); err != nil { fmt.Println(err) } l.Close() } // <string>:1: not allow to modify table 复制代码`

## 写在最后 ##

Golang 和 Lua 的融合，开阔了我的视野：原来静态语言和动态语言还能这么融合，静态语言的运行高效率，配合动态语言的开发高效率，想想都兴奋（逃。

在网上找了很久，发现并没有关于 Go-Lua 的技术分享，只找到了一篇稍微有点联系的文章（ [京东三级列表页持续架构优化 — Golang + Lua (OpenResty) 最佳实践]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzIwODA4NjMwNA%3D%3D%26amp%3Bmid%3D2652898074%26amp%3Bidx%3D1%26amp%3Bsn%3Dd9770ca7b5f5707b041cb1568d8a19a4%26amp%3Bchksm%3D8cdcd755bbab5e43d52d938644e0b754545c84d718eaafe2330ea2f6778c0b416bc5836ad2b3%26amp%3Bmpshare%3D1%26amp%3Bscene%3D1%26amp%3Bsrcid%3D1121RtxVlJ77tySDxb7hxgax%26amp%3Bfrom%3Dgroupmessage%23wechat_redirect ) ），且在这篇文章中， Lua 还是跑在 C 上的。由于信息的缺乏以及本人（学生党）开发经验不足的原因，并不能很好地评价该方案在实际生产中的可行性。因此，本篇文章也只能当作“闲文”了，哈哈。

## 参考资料 ##

* [深入浅出Lua虚拟机]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000014342487 )
* [A No-Frills Introduction to Lua 5.1 VM Instructions]( https://link.juejin.im?target=http%3A%2F%2Fluaforge.net%2Fdocman%2F83%2F98%2FANoFrillsIntroToLua51VMInstructions.pdf )
* [cocos2d-lua disable unexpected global variable]( https://link.juejin.im?target=https%3A%2F%2Fgist.github.com%2Fdualface%2F798d922023d773902287 )
* [lua中设置只读table]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fnellson%2Farticle%2Fdetails%2F10928063 )
* [MetableEvents]( https://link.juejin.im?target=http%3A%2F%2Flua-users.org%2Fwiki%2FMetatableEvents )
* [github.com/zhu327/glua…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzhu327%2Fglualor )