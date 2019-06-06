# 好程序员web前端教程：字符串 #

好程序员web前端教程：字符串，观察某宝网商品数据，有一个东西叫服务器>>>>js的作用重要作用之一>>>> **交互** >>>>人机交互(事件)>>>>服务器交互(ajax)；

服务器交互，数据处理方式json>>>>>要把它转化成字符串操作。

字符串操作重要性不言而喻。

## 什么是字符串? ##

字符串就是一串字符，由单（双）引号括起来，字符串是JavaScript的一种基本类型。

● "undefined"——如果这个值未定义；

● "boolean"——如果这个值是布尔值；

● "string"——如果这个值是字符串；

● "number"——如果这个值是数值；

● "object"——如果这个值是对象或null；

字符串的操作 >>>>> 从1+1=2到1+1=11又怎样的区别那？(小复习)

字符串的声明：

var str="亲"； 基本类型 定义一个字符串变量str，内容为‘亲'

var str = new String(“hello”); 引用类型 定义一个字符串变量str，内容为hello， 注意此刻str为object(对象)类型 用new产生的变量都是引用类型的变量，也叫对象。

JavaScript特性之一>>>>>>万事万物皆对象；

基本类型值指的是简单的数据段，而引用类型是一个指向，指向javascript的内部对象。

字符串与html

1.当把html编译成字符串插入到页面中的时候 ， JavaScript解析器会直接将字符串解析成代码。 比如:document.write('<strong>我是加粗的文字/strong>')

写在页面上是什么样子那？for循环和字符串拼接。(练习)

big() 用大号字体显示字符串

bold() 使用粗体显示字符串

fixed() 以打字机文本显示字符串

strike() 使用删除线来显示字符串

fontcolor() 使用指定颜色来显示字符串

fontsize() 使用指定尺寸来显示字符串

link() 将字符串显示为链接

sub() 把字符串显示为下标

sup() 把字符串显示为上标

//上述方法，都返回一个增加了标签的字符串,但是不对字符串本身进行操作；

没有html代码的商品列表页面

**![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b6d6e98ad5a1?imageView2/0/w/1280/h/960/ignore-error/1)**

**![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b6d6e9ec8632?imageView2/0/w/1280/h/960/ignore-error/1)**

两种声明字符串的方式差别。类型不同（原因）。字符串的下标length

字符串操作>>>>

1.查询操作

1)indexOf("abc") 查找字符串第一次出现的位置 ;

2)lastIndexOf("abc") 查找字符串最后一次出现的位置 如果没找到 返回-1

3)replace() 替换字符串//返回一个修改后的字符串不对原字符串进行操作

replace 替换字符串

如： var str="how are you";

alert(str.replace("are","old are"));

2.获取操作

charAt(3) //获取下标为3的字符

charCodeAt(3) //获取下标为3的字符的Unicode码

Unicode（统一码、万国码、单一码）是计算机科学领域里的一项业界标准,包括字符集、编码方案等。>>>>>翻译官思密达

String.fromCharCode(94) //编码转换成字符；

由于fromCharCode( )是String对象中的方法 ，所以在使用的时候要加上前缀String；

**![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b6d6ea1f7481?imageView2/0/w/1280/h/960/ignore-error/1)**

substring(start,end)//截取字符串，从第start位开始，到end位停止。

字符集

GBK、GB2312、GB18030、BIG5（繁体中文）

Unicode-8 UTF-8 Unicode-16

split(separator, howmany) >>>>>> 根据分隔符、拆分成数组;

separator (字符串);//根据什么进行拆分

howmany(可以指定返回的数组的最大长度) ;

【注】如果空字符串(“”)用作separator,那么stringObject中的每个字符之间都会被分割。

3.拼接操作

concat() 连接字符串 //最没用的方法

4.大小写操作

toLowerCase（）

toUpperCase（）

字符串操作练习