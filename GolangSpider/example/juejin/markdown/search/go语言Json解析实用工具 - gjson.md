# go语言Json解析实用工具 - gjson #

# 一、目录 #

* [一、目录]( #%E4%B8%80%E7%9B%AE%E5%BD%95 )
* [二、简介]( #%E4%BA%8C%E7%AE%80%E4%BB%8B )

* [1. 为什么要使用gjson?]( #1-%E4%B8%BA%E4%BB%80%E4%B9%88%E8%A6%81%E4%BD%BF%E7%94%A8gjson )
* [2. Result结构体]( #2-result%E7%BB%93%E6%9E%84%E4%BD%93 )
* [3. 属于Result的函数]( #3-%E5%B1%9E%E4%BA%8Eresult%E7%9A%84%E5%87%BD%E6%95%B0 )
* [4. 初始化函数]( #4-%E5%88%9D%E5%A7%8B%E5%8C%96%E5%87%BD%E6%95%B0 )
* [5. 判断Json是否合法]( #5-%E5%88%A4%E6%96%ADjson%E6%98%AF%E5%90%A6%E5%90%88%E6%B3%95 )

* [三、实际操作]( #%E4%B8%89%E5%AE%9E%E9%99%85%E6%93%8D%E4%BD%9C )

* [1. 使用]( #1-%E4%BD%BF%E7%94%A8 )

# 二、简介 #

## 1. 为什么要使用gjson? ##

golang初学者肯定会觉得Json的解析十分麻烦。其实是要转换思维，我们不能像PHP或JS一样把Json直接转化为对象。
所以我们定义一系列的函数去获取Json里面的值。
gjson(github.com/tidwall/gjson) 很好的支持了各种Json操作。使用它可以方便地解析Json，并进行判断、取值。

## 2. Result结构体 ##

` // 首先定义一个Result结构体，它是所有数据的抽象 type Result struct { Type Type // 该结构体在Json中的类型 Raw string // 原json串 Str string // 字符串 Num float64 // 浮点数 Index int // 索引 } 复制代码`

## 3. 属于Result的函数 ##

` func (t Result) Exists () bool // 判断某值是否存在 func (t Result) Value () interface {} // func (t Result) Int () int64 func (t Result) Uint () uint64 func (t Result) Float () float64 func (t Result) String () string func (t Result) Bool () bool func (t Result) Time () time. Time func (t Result) Array () [] gjson. Result func (t Result) Map () map [ string ] gjson. Result func (t Result) Get (path string ) Result func (t Result) ForEach (iterator func (key, value Result) bool ) // 可传闭包函数 func (t Result) Less (token Result, caseSensitive bool ) bool 复制代码`

## 4. 初始化函数 ##

` gjson.Parse(json).Get( "name" ).Get( "last" ) gjson.Get(json, "name" ).Get( "last" ) gjson.Get(json, "name.last" ) 复制代码`

## 5. 判断Json是否合法 ##

` if !gjson.Valid(json) { return errors.New( "invalid json" ) } 复制代码`

# 三、实际操作 #

## 1. 使用 ##

` package main import ( "fmt" "log" "strings" "github.com/tidwall/gjson" ) const json = `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Dale","last":"Murphy","age":44},{"first":"Roger","last":"Craig","age":68},{"first":"Jane","last":"Murphy","age":47}]}` func main () { // 首先我们判断该json是否合法 if !gjson.Valid(json) { log.Fatalf( "%s" , "invalid json" ) } // 获取Json中的age age := gjson.Get(json, `age` ).Int() fmt.Printf( "%T, %+v\n" , age, age) // 获取lastname lastname := gjson.Get(json, `name.last` ).String() fmt.Printf( "%T, %+v\n" , lastname, lastname) // 获取children数组 for _, v := range gjson.Get(json, `children` ).Array() { fmt.Printf( "%q " , v.String()) } fmt.Println() // 获取第二个孩子 fmt.Printf( "%q\n" , gjson.Get(json, `children.1` ).String()) fmt.Printf( "%q\n" , gjson.Get(json, `children|1` ).String()) // 通配符获取第三个孩子 fmt.Printf( "%q\n" , gjson.Get(json, `child*.2` ).String()) // 反转数组函数 fmt.Printf( "%q\n" , gjson.Get(json, `children|@reverse` ).Array()) // 自定义函数 - 全转大写 gjson.AddModifier( "case" , func (json, arg string ) string { if arg == "upper" { return strings.ToUpper(json) } return json }) fmt.Printf( "%+v\n" , gjson.Get(json, `children|@case:upper` ).Array()) // 直接解析为map jsonMap := gjson.Parse(json).Map() fmt.Printf( "%+v\n" , jsonMap) for _, v := range jsonMap { fmt.Printf( "%T, %+v\n" , v, v) } } 复制代码`