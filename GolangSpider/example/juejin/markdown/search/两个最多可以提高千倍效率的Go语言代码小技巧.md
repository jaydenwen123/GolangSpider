# 两个最多可以提高千倍效率的Go语言代码小技巧 #

![http://dawngrp.com/gao-xiao-de-goyu-yan-bian-ma-ji-qiao/](https://user-gold-cdn.xitu.io/2019/3/22/169a495325da53cf?imageView2/0/w/1280/h/960/ignore-error/1)

#### 1.不要使用+和fmt.Sprintf操作字符串 ####

+操作字符串很方便，但是真的很慢，在Go语言里使用+会导致你的程序跑得可能比脚本语言还满，不相信的可以自己做个测试，用+操作，迭代十万次，Python、Javascript都比Go快很多（是很多噢，不是一点点）

` func TestStr (t *testing.T) { str := "" for i := 0 ; i < 100000 ; i++ { str += "test" } } 复制代码`

测试结果

> 
> 
> 
> PASS: TestStr (3.32s)
> 
> 

` str= "" for i in range( 100000 ): str+= "test" 复制代码`

测试结果:

> 
> 
> 
> ~/» time python test.py
> 0.03s user 0.03s system 81% cpu 0.078 total
> 
> 

作为静态语言的Go，居然在这么一个段简单的代码上执行效率比Python慢了100倍，不可思议吧？不是Go的问题，而是在Go中使用+处理字符串是很消耗性能的，而Python应该是对+操作字符串进行了重载优化。（Javascript +操作字符串也很快）

##### 最有效的方式是采用buffer #####

` strBuf := bytes.NewBufferString( "" ) for i := 0; i < 100000; i++ { strBuf.WriteString( "test" ) } 复制代码`

结果可以自己测试，会让你很惊讶

有一些需要简单组合两个字符串，用Buffer麻烦了点，比较容易让人想到的就是用fmt.Sprintf()来组合，很多包里的源码也是这么写的。其实fmt的Sprintf也非常慢，如果没有复杂的类型转换输出的情况下，使用strings.Join性能会高很多

` func TestStr(t *testing.T) { a, b := "Hello" , "world" for i := 0; i < 1000000; i++ { fmt.S printf ( "%s%s" , a, b) //strings.Join([]string{a, b}, "" ) } } 复制代码`
> 
> 
> 
> 
> PASS: TestStr (0.29s)
> 
> 

` func TestStr(t *testing.T) { a, b := "Hello" , "world" for i := 0; i < 1000000; i++ { //fmt.S printf ( "%s%s" , a, b) strings.Join([]string{a, b}, "" ) } } 复制代码`
> 
> 
> 
> 
> PASS: TestStr (0.09s)
> 
> 

从结果来看strings.Join 比用Sprint快4倍左右吧。

#### 2.对于固定字段的键值对，用临时Struct，不要用map[string]interface{} ####

举个简单的例子

` func TestData(t *testing.T) { for i := 0; i < 100000000; i++ { var a struct { Name string Age int } a.Name = "Hello" a.Age = 10 } } 复制代码`
> 
> 
> 
> 
> PASS: TestData (0.04s)
> 
> 

` func TestData2(t *testing.T) { for i := 0; i < 100000000; i++ { var a = map[string]interface{}{} a[ "Name" ] = "Hello" a[ "Age" ] = 10 } } 复制代码`
> 
> 
> 
> 
> PASS: TestData2 (38.30s)
> 
> 

相差上千倍的效率呢！ 在能够知道字段的情况下，用临时Struct在运行期间不需要动态分配内容，并且不需要像map那样去检查索引，所以速度会快非常多。