package main

import (
	"fmt"
	"os"
)

func main() {

	//filename := "baidu.gif"
	//url := "http://s1.bdstatic.com/r/www/cache/mid/static/xueshu/img/logo_4b1971d.gif"
	//err := util.Download(url, filename)
	//if err != nil {
	//	logs.Info("download file failed:", err.Error())
	//} else {
	//	logs.Info("download file success")
	//}
	//hash := md5.New()
	//strtoupper($mv_hash)."kugoumvcloud"
	//06a824a75307983ab83a69a1d80325f6
	//06a824a75307983ab83a69a1d80325f6

	//for {
	//	fmt.Printf("ssafas>\r")
	//	time.Sleep(time.Millisecond*1000)
	//}
	//fmt.Println()
	//fmt.Printf("ssafas>\r\r")
	//fmt.Println()
	//fmt.Printf("ssafas>\r\r\r")
	//fmt.Println()
	//fmt.Printf("ssafas>\b")
	//fmt.Println()
	//fmt.Printf("ssafas>\b\b")
	//fmt.Println()
	//fmt.Printf("ssafas>\b\b\b")
	//fmt.Println()
	//TestFmt()
	//c83b410711be23b67b064077dc5b238a
	//c83b410711be23b67b064077dc5b238a

	//md5加密算法
	//data:=[]byte(strings.ToUpper("51967E53BB70A713E908443F5350A21C")+"kugoumvcloud")
	//val := md5.Sum(data)
	//fmt.Printf("%x",val)
	ProxySendRequest()
}
//http://116.208.55.252:9999
//http://115.53.34.166:9999
//http://106.110.65.109:8118
//http://114.119.116.92:61066
//http://218.73.131.6:9999
//http://125.108.85.13:9000
//http://115.193.100.80:9000
//http://59.62.26.232:9000
//http://117.90.4.204:9000
//http://114.55.103.83:9090
//http://183.129.207.86:13340

type point struct {
	x, y int
}
func TestFmt()  {
	//Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
	p := point{1, 2}
	fmt.Printf("%v\n", p) // {1 2}
	//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
	//需要打印值的类型，使用 %T。
	fmt.Printf("%T\n", p) // main.point
	//格式化布尔值是简单的。
	fmt.Printf("%t\n", true)
	//格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
	fmt.Printf("%d\n", 123)
	//这个输出二进制表示形式。
	fmt.Printf("%b\n", 14)
	//这个输出给定整数的对应字符。
	fmt.Printf("%c\n", 33)
	//%x 提供十六进制编码。
	fmt.Printf("%x\n", 456)
	//对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
	fmt.Printf("%f\n", 78.9)
	//%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	//使用 %s 进行基本的字符串输出。
	fmt.Printf("%s\n", "\"string\"")
	//像 Go 源代码中那样带有双引号的输出，使用 %q。
	fmt.Printf("%q\n", "\"string\"")
	//和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
	fmt.Printf("%x\n", "hex this")
	//要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &p)
	//当输出数字的时候，你将经常想要控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	//你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	//要最对齐，使用 - 标志。
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	//要左对齐，和数字一样，使用 - 标志。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	//到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。Sprintf 则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	//你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
