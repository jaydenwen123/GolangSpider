# ES6字符串新增扩展方法以及模板字符串的学习 #

# 扩展方法 #

## includes()、startsWith()、endsWith() ##

ES6之前，我们只能用 ` indexOf` 来判断一个字符串是否包含另外一个字符串。现在ES6提供了三个新方法：

* ` includes` ：返回布尔值，表示是否找到字符串
* ` startsWith` ：返回布尔值，表示参数字符串是否在源字符串的开头
* ` endsWith` ：返回布尔值，表示参数字符串是否在源字符串的开头

` let str = 'i am test string' console.log(str.includes( 't' )) // true console.log(str.startsWith( 't' )) // false console.log(str.startsWith( 'i' )) // true console.log(str.endsWith( 't' )) // false console.log(str.endsWith( 'g' )) // true 复制代码`

这三个方法还接受第二个参数，表示开始搜索的位置, ` endsWith()` 特殊一些，它的第二个参数n，表示针对前n个字符串，一定要注意

` console.log(str.includes( 't' , 12)) // false console.log(str.startsWith( 't' , 5)) // true console.log(str.endsWith( 'g' , 6)) // false 复制代码`

## repeat() ##

` repeat` 方法返回一个新的字符串，将原字符串重复n次。

` 'test'.repeat(2) // "testtest" 复制代码`

特殊情况需要注意一下：

` 'test'.repeat(0) // "" 'test'.repeat(2.3) // "repeat" ，小数会被向下取整 'test'.repeat(-0.9) // "" , 在0到-1之间的小数，则等同于0 'test'.repeat(NaN) // "" 参数NaN等同于0 'test'.repeat( 'test' ) // "" , 字符串会先转换为数字，字符串会被转换成NaN，之后会被解析为0 'test'.repeat(Infinity) // 报错 'test'.repeat(-1) // 报错 // 报错信息 Uncaught RangeError: Invalid count value at String.repeat (<anonymous>) at <anonymous>:1:8 复制代码`

## padStart()、padEnd() ##

ES6引入了字符串补全长度的功能。

` padStart()` 用于头部补全， ` padEnd()` 用于尾部补全

方法接手两个参数，第一个参数用来指定字符串的最小长度，第二参数是用来补全的字符串。

` 'a'.padStart(5, 'b' ) // "bbbba" 'a'.padEnd(3, 'b' ) // "abb" 复制代码`

特殊情况：

### 1.如果原字符串长度等于或者大于指定的最小长度，则返回原字符串 ###

` 'abcdefg'.padStart(2, 'hijk' ) // "abcdefg" 'abcdefg'.padEnd(2, 'hijk' ) // "abcdefg" 复制代码`

### 2.如果用来补全的字符串加上原来的字符串的长度之和超出了指定长度，则会截去超出的补全字符串。 ###

` 'abcdefg'.padStart(10, 'hijklmnopq' ) // "hijabcdefg" 'abcdefg'.padEnd(10, 'hijklmnopq' ) // "abcdefghij" 复制代码`

### 3.如果用来补全的字符串参数为空，则用空格来补全 ###

` 'a'.padStart(10) // " a" 'a'.padEnd(10) // "a " 复制代码`

# 模板字符串（template string） #

传统的模板是通过字符串拼接生成的，比如：

` 'i am <div>' + param + '</div>' 复制代码`

在模板繁杂的时候，大量的拼接使得这种处理相当不便利，ES6引入了 ` 模板字符串` 来解决这个问题。 我们改写下上面的例子

` `i am <div> ${param} </div` 复制代码`

使用反引号(`)标识模板字符串，使用 ` ${}` 来包裹变量， ` ${}` 不仅可以存放变量，也可以放入任意JavaScript表达式，可以进行运算，引用属性对象以及调用方法。

` let param = 'daly' `i am ${param} ` // "i am daly" 复制代码` ` let x = 1 let y = 2 `result is ${x + y} ` // "result is 3" 复制代码` ` let obj = {str: 'daly' } `i am ${obj.str} ` // "i am daly" 复制代码` ` function test (){ return 'daly' } `i am ${test()} ` // "i am daly" 复制代码`

## 标签模板 ##

模板字符串可以紧跟在函数名后面，该函数将被调用来处理这个模板字符串。

这被称为 ` 模板标签（tagged template）`

` console.log `i am daly` // [ "i am daly" , raw: Array(1)] 复制代码`

标签模板是函数调用的一种特殊形式，“标签”指的就是函数，紧跟在后面的模板字符串就是参数。

但是参数会被拆分成多个，第一个是一个数组，包含没有变量的字符串，其他的参数就是变量被替换后的值。

` let a = 'daly' let b = 'xuess' function tag(str,...values) { console.log(str, ...values) } tag `They are ${a} and ${b} ` // [ "They are " , " and " , "" , raw: Array(3)] "daly" "xuess" 复制代码`