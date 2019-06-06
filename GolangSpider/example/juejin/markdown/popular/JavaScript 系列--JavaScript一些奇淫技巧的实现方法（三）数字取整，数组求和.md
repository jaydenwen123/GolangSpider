# JavaScript 系列--JavaScript一些奇淫技巧的实现方法（三）数字取整，数组求和 #

## 一、前言 ##

简短的sleep函数，获取时间戳： [www.mwcxs.top/page/746.ht…]( https://link.juejin.im?target=https%3A%2F%2Fwww.mwcxs.top%2Fpage%2F746.html )

数字格式化 1234567890 --> 1,234,567,890；argruments 对象(类数组)转换成数组：

[www.mwcxs.top/page/749.ht…]( https://link.juejin.im?target=https%3A%2F%2Fwww.mwcxs.top%2Fpage%2F749.html )

今天我们来介绍一下数字取整，数组求和。

## 二、数字取整 ##

### 1、普通版 ###

const a = parseInt(2.33333); parseInt()方法是解析一个字符串参数，并返回一个指定基数的整数。这个就是我们最常用的取整的最常用的方式。

parseInt() 函数解析一个字符串参数，并返回一个指定基数的整数 (数学系统的基础)。

parseInt语法：parseInt(string, radix);

string：要被解析的值。如果参数不是一个字符串，则将其转换为字符串(使用  ToString 抽象操作)。字符串开头的空白符将会被忽略。

radix：一个介于2和36之间的整数(数学系统的基础)，表示上述字符串的基数。比如参数"10"表示使用我们通常使用的十进制数值系统。始终指定此参数可以消除阅读该代码时的困惑并且保证转换结果可预测。当未指定基数时，不同的实现会产生不同的结果，通常将值默认为10。

### 2、进阶版 ###

const a = Math.trunc(2.33333) Math.trunc()方法会将数字的小数部分去掉，只保留整数部分（常说的“取整”，不是四舍五入）。

注意：Internet Explorer 不支持这个方法，不过写个 Polyfill 也很简单：

` Math.trunc = Math.trunc || function (x) { if (isNaN(x)) { return NaN; } if (x > 0) { return Math.floor(x); } return Math.ceil(x); }; 复制代码`

数学的事情还是用数学方法来处理比较好。

### 3、~~number ###

这个符号是什么鬼，没有用过，不要紧，慢慢看。这个~~操作符也被称为“双按位非”操作符。你通常可以使用它作为替代Math.trunc()的更快的方法。

` console.log(~~66.11) // 66 console.log(~~12.9999) // 12 console.log(~~6) // 6 console..log(~~-6.9999999999) // -6 console.log(~~[]) // 0 console.log(~~NaN) // 0 console.log(~~null) // 0 复制代码`

失败时返回0,这可能在解决 Math.trunc() 转换错误返回 NaN 时是一个很好的替代。

注意：但是当数字范围超出 ±2^31−1 即：2147483647 时，异常就出现了。

// 异常情况

` console.log(~~2147493647.123) // -> -2147473649 🙁 复制代码`

### 4、按位或 number|0 ###

这个就比较容易理解了。| 是按位或，对每一对比特位执行或（OR）操作。

` console.log(20.15|0); // 20 console.log((-20.15)|0); // -20 复制代码`

注意：但是当数字范围超出 ±2^31−1 即：2147483647 时，异常就出现了。

` console.log(3000000000.15|0); // -1294967296 复制代码`

### 5、按位异或 number^0 ###

^ (按位异或)，对每一对比特位执行异或（XOR）操作。

` console.log(20.15^0); // -> 20 console.log((-20.15)^0); // -> -20 复制代码`

注意：但是当数字范围超出 ±2^31−1 即：2147483647 时，异常就出现了。

console.log(3000000000.15^0); // -> -1294967296

### 6、左移 number<<0 ###

<< (左移) 操作符会将第一个操作数向左移动指定的位数。向左被移出的位被丢弃，右侧用 0 补充。

` console.log(20.15 << 0); // -> 20 console.log((-20.15) << 0); //-20 复制代码`

注意：但是当数字范围超出 ±2^31−1 即：2147483647 时，异常就出现了。

` console.log(3000000000.15 << 0); // -> -1294967296 复制代码`

上面讲的按位运算符方法执行很快，当你执行数比较大的时候非常适用，能看出来区别。

注意：当数字超过±2^31−1（2147483647）的范围，会有一些异常，使用前判断数值的范围。

前端知识点：按位运算

## 三、数组求和 ##

### 1、普通版 ###

` let arr = [1, 2, 3, 4, 5] function sum(arr){ let x = 0 for ( let i = 0; i < arr.length; i++){ x += arr[i] } return x } sum(arr) // 15 复制代码`

优点：通俗易懂，简单粗暴 缺点：没有亮点，太通俗

### 2、优雅版本 ###

` let arr = [1, 2, 3, 4, 5] function sum(arr) { return arr.reduce((prev, item) => prev + item) } sum(arr) //15 复制代码`

优点：简单明了，数组迭代器方式清晰直观 缺点：不兼容 IE 9以下浏览器

### 3、终极版 ###

` let arr = [1, 2, 3, 4, 5] function sum(arr) { return eval (arr.join( "+" )) } sum(arr) //15 复制代码`

优点：让人一时看不懂的就是"好方法"。

缺点：（1）eval 不容易调试，用 chromeDev 等调试工具无法打断点调试，所以麻烦的东西也是不推荐使用的。

（2）性能问题，在旧的浏览器中如果你使用了eval，性能会下降10倍。在现代浏览器中有两种编译模式：fast path和slow path。fast path是编译那些稳定和可预测（stable and predictable）的代码。而明显的，eval 不可预测，所以将会使用 slow path ，所以会慢。

前端知识点：eval的使用细则

【谢谢关注和阅读，后续新的文章首发：sau交流学习社区： [www.mwcxs.top/】]( https://link.juejin.im?target=https%3A%2F%2Fwww.mwcxs.top%2F%25E3%2580%2591 )