# Go之底层利器-AST遍历 #

原文出处： [dpjeep.com/gozhi-di-ce…]( https://link.juejin.im?target=https%3A%2F%2Fdpjeep.com%2Fgozhi-di-ceng-li-qi-astbian-li%2F )

## 背景 ##

最近需要基于AST来做一些自动化工具，遂也需要针对这个神兵利器进行一下了解研究。本篇文章也准备只是简单的讲解一下以下两个部分：

* 通过AST解析一个Go程序
* 然后通过Go的标准库来对这个AST进行分析

## AST ##

什么是 [AST]( https://link.juejin.im?target=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2F%25E6%258A%25BD%25E8%25B1%25A1%25E8%25AA%259E%25E6%25B3%2595%25E6%25A8%25B9 ) ，其实就是抽象语法树Abstract Syntax Tree的简称。它以树状的形式表现编程语言的语法结构，树上的每个节点都表示源代码中的一种结构。之所以说语法是“抽象”的，是因为这里的语法并不会表示出真实语法中出现的每个细节。

## 主菜 ##

### 开胃提示语 ###

以下内容有点长，要不先去买点瓜子，边磕边看？

### 编译过程 ###

要讲解相关AST部分，先简单说一下我们知道的编译过程：

* 词法分析
* 语法分析
* 语义分析和中间代码产生
* 编译器优化
* 目标代码生成 而我们现在要利用的正是Google所为我们准备的一套非常友好的词法分析和语法分析工具链，有了它我们就可以造车了。

### 代码示例 ###

在Golang官方文档中已经提供实例，本处就不把 [文档源码]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fsrc%2Fgo%2Fast%2Fexample_test.go ) 贴出来了，只放出部分用例

` // This example shows what an AST looks like when printed for debugging. func ExamplePrint () { // src is the input for which we want to print the AST. src := ` package main func main() { println("Hello, World!") } ` // Create the AST by parsing src. fset := token.NewFileSet() // positions are relative to fset f, err := parser.ParseFile(fset, "" , src, 0 ) if err != nil { panic (err) } // Print the AST. ast.Print(fset, f) // Output: // 0 *ast.File { // 1 . Package: 2:1 // 2 . Name: *ast.Ident { // 3 . . NamePos: 2:9 // 4 . . Name: "main" // 5 . } // 6 . Decls: []ast.Decl (len = 1) { // 7 . . 0: *ast.FuncDecl { // 8 . . . Name: *ast.Ident { // 9 . . . . NamePos: 3:6 // 10 . . . . Name: "main" // 11 . . . . Obj: *ast.Object { // 12 . . . . . Kind: func // 13 . . . . . Name: "main" // 14 . . . . . Decl: *(obj @ 7) // 15 . . . . } // 16 . . . } // 17 . . . Type: *ast.FuncType { // 18 . . . . Func: 3:1 // 19 . . . . Params: *ast.FieldList { // 20 . . . . . Opening: 3:10 // 21 . . . . . Closing: 3:11 // 22 . . . . } // 23 . . . } // 24 . . . Body: *ast.BlockStmt { // 25 . . . . Lbrace: 3:13 // 26 . . . . List: []ast.Stmt (len = 1) { // 27 . . . . . 0: *ast.ExprStmt { // 28 . . . . . . X: *ast.CallExpr { // 29 . . . . . . . Fun: *ast.Ident { // 30 . . . . . . . . NamePos: 4:2 // 31 . . . . . . . . Name: "println" // 32 . . . . . . . } // 33 . . . . . . . Lparen: 4:9 // 34 . . . . . . . Args: []ast.Expr (len = 1) { // 35 . . . . . . . . 0: *ast.BasicLit { // 36 . . . . . . . . . ValuePos: 4:10 // 37 . . . . . . . . . Kind: STRING // 38 . . . . . . . . . Value: "\"Hello, World!\"" // 39 . . . . . . . . } // 40 . . . . . . . } // 41 . . . . . . . Ellipsis: - // 42 . . . . . . . Rparen: 4:25 // 43 . . . . . . } // 44 . . . . . } // 45 . . . . } // 46 . . . . Rbrace: 5:1 // 47 . . . } // 48 . . } // 49 . } // 50 . Scope: *ast.Scope { // 51 . . Objects: map[string]*ast.Object (len = 1) { // 52 . . . "main": *(obj @ 11) // 53 . . } // 54 . } // 55 . Unresolved: []*ast.Ident (len = 1) { // 56 . . 0: *(obj @ 29) // 57 . } // 58 } } 复制代码`

一看到上面的打印是不是有点头晕？哈哈，我也是。没想到一个简单的hello world就能打印出这么多东西，里面其实隐藏了很多有趣的元素，比如函数、变量、评论、imports等等，那我们要如何才能从中提取出我们想要的数据呢？为达这个目的，我们需要用到Golang所为我们提供的 ` go/parser` 包：

` // Create the AST by parsing src. fset := token.NewFileSet() // positions are relative to fset f, err := parser.ParseFile(fset, "" , src, 0 ) if err != nil { panic (err) } 复制代码`

第一行引用了 ` go/token` 包，用来创建一个新的用于解析的源文件FileSet。
然后我们使用的 ` parser.ParseFile` 返回的是一个 ` ast.File` 类型结构体（ [原始文档]( https://link.juejin.im?target=https%3A%2F%2Fgolang.org%2Fpkg%2Fgo%2Fast%2F%23File ) ）,然后回头查看上面的日志打印，每个字段元素的含义你也许已经霍然开朗了，结构体定义如下：

` type File struct { Doc *CommentGroup // associated documentation; or nil Package token.Pos // position of "package" keyword Name *Ident // package name Decls []Decl // top-level declarations; or nil Scope *Scope // package scope (this file only) Imports []*ImportSpec // imports in this file Unresolved []*Ident // unresolved identifiers in this file Comments []*CommentGroup // list of all comments in the source file } 复制代码`

好了，目前我们就是要利用这个结构体做一下小的代码示例，我们就来解析下面的这个文件 ` ast_traversal.go` ：

` package ast_demo import "fmt" type Example1 struct { // Foo Comments Foo string `json:"foo"` } type Example2 struct { // Aoo Comments Aoo int `json:"aoo"` } // print Hello World func PrintHello () { fmt.Println( "Hello World" ) } 复制代码`

我们已经可以利用上面说到的 ` ast.File` 结构体去解析这个文件了，比如利用 ` f.Imports` 列出所引用的包：

` for _, i := range f.Imports { t.Logf( "import: %s" , i.Path.Value) } 复制代码`

同样的，我们可以过滤出其中的评论、函数等，如：

` for _, i := range f.Comments { t.Logf( "comment: %s" , i.Text()) } for _, i := range f.Decls { fn, ok := i.(*ast.FuncDecl) if !ok { continue } t.Logf( "function: %s" , fn.Name.Name) } 复制代码`

上面，获取comment的方式和import类似，直接就能使用，而对于函数，则采用了 ` *ast.FucDecl` 的方式，此时，移步至本文最上层，查看AST树的打印，你就发现了 ` Decls: []ast.Decl` 是以数组形式存放，且其中存放了多种类型的node，此处通过强制类型转换的方式，检测某个类型是否存在，存在的话则按照该类型中的结构进行打印。上面的方式已能满足我们的基本需求，针对某种类型可以进行具体解析。
但是，凡是还是有个但是，哈哈，通过上面的方式来一个一个解析是不是有点麻烦？没事，谷歌老爹通过 ` go/ast` 包给我们又提供了一个方便快捷的方法：

` // Inspect traverses an AST in depth-first order: It starts by calling // f(node); node must not be nil. If f returns true, Inspect invokes f // recursively for each of the non-nil children of node, followed by a // call of f(nil). // func Inspect (node Node, f func (Node) bool ) { Walk(inspector(f), node) } 复制代码`

这个方法的大概用法就是：通过深度优先的方式，把整个传递进去的AST进行了解析，它通过调用f(node) 开始；节点不能为零。如果 f 返回 true，Inspect 会为节点的每个非零子节点递归调用f，然后调用 f(nil)。相关用例如下：

` ast.Inspect(f, func (n ast.Node) bool { // Find Return Statements ret, ok := n.(*ast.ReturnStmt) if ok { t.Logf( "return statement found on line %d:\n\t" , fset.Position(ret.Pos()).Line) printer.Fprint(os.Stdout, fset, ret) return true } // Find Functions fn, ok := n.(*ast.FuncDecl) if ok { var exported string if fn.Name.IsExported() { exported = "exported " } t.Logf( "%sfunction declaration found on line %d: %s" , exported, fset.Position(fn.Pos()).Line, fn.Name.Name) return true } return true }) 复制代码`

## 后记 ##

至此，你手中的瓜子可能已经嗑完了，AST用处颇多，上面我们所讲到的也只是AST其中的一小部分，很多底层相关分析工具都是基于它来进行语法分析进行，工具在手，然后要制造什么艺术品就得看各位手艺人了。后续会陆续更新部分基于Go AST的小工具出来，希望自己能早日实现吧，哈哈😆。
以下为上文中所用到的测试用例及使用AST针对结构体进行字段解析的源码，我已提交至Github，如有兴趣可以去看看

* [ast-demo]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpingdai%2Fast-demo )
* [parseStruct]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpingdai%2FparseStruct )