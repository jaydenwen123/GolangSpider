# Go 语言指针符号 *和& #

先看一段代码，人工运行一下，看看自己能做对几题？

` package main import "fmt" func main () { var a int = 1 var b *int = &a var c **int = &b var x int = *b fmt.Println( "a = " ,a) fmt.Println( "&a = " ,&a) fmt.Println( "*&a = " ,*&a) fmt.Println( "b = " ,b) fmt.Println( "&b = " ,&b) fmt.Println( "*&b = " ,*&b) fmt.Println( "*b = " ,*b) fmt.Println( "c = " ,c) fmt.Println( "*c = " ,*c) fmt.Println( "&c = " ,&c) fmt.Println( "*&c = " ,*&c) fmt.Println( "**c = " ,**c) fmt.Println( "***&*&*&*&c = " ,***&*&*&*&*&c) fmt.Println( "x = " ,x) } 复制代码`

#### 符号 & 的意思是对变量取地址 ####

如：变量a的地址是&a

#### 符号 * 的意思是对指针取值， ####

如:*&a，就是a变量所在地址的值，当然也就是a的值了

#### 简单的解释 ####

` *和& 可以互相抵消 但是注意 【 *& 】可以抵消掉，但【 &* 】是不可以抵消的 a和 *&a 是一样的,都是a的值，值为1 (因为*&互相抵消掉了) a和 *&*&*&*&a是一样的，都是1 (因为4个*&互相抵消掉了) 因为有 var b *int = &a 所以 a和*&a和*b是一样的，都是a的值，值为1 (把b当做&a看) 因为有 var c **int = &b 所以 **c和**&b是一样的，把&约去后 会发现**c和b是一样的 (从这里也不难看出，c和b也是一样的) 又因为上面得到的&a和b是一样的 所以**c和&a是一样的， 再次把*&约去后**c和a`是一样的，都是1 复制代码`

你也试着运行一下吧 ~

` ## 运行结果 $ go run main.go a = 1 &a = 0xc200000018 *&a = 1 b = 0xc200000018 &b = 0xc200000020 *&b = 0xc200000018 *b = 1 c = 0xc200000020 *c = 0xc200000018 &c = 0xc200000028 *&c = 0xc200000020 **c = 1 ***&*&*&*&c = 1 x = 1 复制代码`

两个符号抵消顺序

` *&可以在任何时间抵消掉，但&*不可以被抵消的，因为顺序不对 fmt.Println( "*&a\t=\t" ,*&a) //成功抵消掉，打印出1，即a的值 fmt.Println( "&*a\t=\t" ,&*a) //无法抵消，会报错 复制代码`