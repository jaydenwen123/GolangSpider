# 《你不知道的Javascript--中卷 学习总结》（类型、值） #

### 类型 ###

类型是 ` 值` 的内部特征，它定义了 ` 值的行为` ，以使其区别于其他值。

#### 内置类型 ####

1、七种内置类型:

* 空值（null）
* 未定义（undefined）
* 布尔值（boolean）
* 数字（number）
* 字符串（string）
* 对象（object）
* 符号（symbol,ES6中新增）

2、我们可以用 ` typeof` 运算符来查看值的类型，它返回的是类型的 ` 字符串值` 。

3、因为typeof null === 'object',所以检测null值的类型需要下面这段复合条件

` var a = null; (!a && typeof a == "object" ); // true 复制代码`

4、函数实际上是object的一个 ` "子类型"` 。具体的说，函数是"可调用对象"，它有一个内部属性 ` [[call]]` ，该属性使其可以被调用。

5、函数对象的 ` length属性` 是 ` 其声明的参数的个数`

#### 值和类型 ####

1、Javascript中的 ` 变量是没有类型的` ，只有 ` 值` 才有。变量可以随时持有任何类型的值。

2、已在作用域中 ` 声明` 但还 ` 没有赋值` 的变量,是 ` undefined` 。

3、还 ` 没有在作用域中声明过得变量` ,是 ` undeclared` 。

` var a; typeof a // "undefined" typeof b // "undefined" 没有声明还是undefined 这主要是因为typeof的安全防范机制 复制代码`

#### typeof安全防范机制 ####

* 全局变量的检验

` // 这样是安全的 if (typeof DEBUG ! == "undefined" ){ console.log( 'xxxxx' ) } 复制代码`

* 非全局变量是否定义

` ( function (){ function FeatureXYZ (){}; function doSomethingCool (){ var helper = (typeof FeatureXYZ !== "undefined" ) ? FeatureXYZ : function (){ // default feature } var a = helper(); } do SomethingCool() })() 复制代码`

* 函数传递参数检测

` function do SomethingCool(FeatureXYZ){ var helper = FeatureXYZ || function (){} var val = helper(); // ... } 复制代码`

### 值 ###

#### 数组 ####

1、数组可以容纳任何类型的值，可以是 ` 字符串` 、 ` 数字` 、 ` 对象` ，甚至是 ` 其他数组` 。

2、使用delete运算符可以将单元从数组中删除，但是请注意，单元删除后，数组的length属性 ` 并不会发生变化` 。

3、在创建"稀疏"数组（即含有空白或空缺单元的数组）。a[1]的值为undefined,但这与将其显示赋值为undefined(a[1] = undefined)还是有所区别。

` var a = []; a[0] = 1; a[2] = [3]; a[1] // undefined a.length // 3 复制代码`

4、数组通过 ` 数字进行索引` ,但是他们也是对象，所以也可以包含字符串 ` 键值` 和 ` 属性` （但这些并 ` 不计算在数组长度内` ）。

5、如果 ` 字符串键值能够被强制类型转换为十进制数字的话` ,它就会被当做数字索引来处理

` var a = []; a[0] = 1; a[ 'foo' ] = 2; a.length // 1 a[ 'foobar' ] //2 a.foobar //2 var a = []; a[ '13' ] = 42; a.length // 14 复制代码`

6、类数组=>数组的方式

* slice

` function foo (){ var arr = Array.prototype.slice.call(arguments); arr.push( 'bam' ); console.log(arr) } foo( 'bar' , 'baz' ) // [ 'bar' , 'baz' , 'bam' ] 复制代码`

* Array.from() (ES6)

` var arr = Array.from(arguments) 复制代码`

#### 字符串 ####

1、Javascript中字符串是不可变的（指字符串的成员函数 ` 不会改变其原始值` ，而是创建并返回一个 ` 新的字符串` ）,而数组是可变的。

2、字符串没有reverse方法，变通方法可以先转换为数组然后转换为字符串。

` var a = 'foo' var c = a.split( "" ).reverse().join( "" ) c; // "oof" 复制代码`

#### 数字 ####

1、Javascript中的数字类型是基于 ` IEEE754标准` 来实现的，该标准通常也被称为"浮点数"。Javascript使用的是"双精度"格式。

2、toFixed() 方法可以指定 ` 小数部分的显示位数` ，如果指定的小数部分的显示位数多于实际位数就用0补齐。

3、toPrecision() 方法用来指定 ` 有效数位` 的显示位数。

4、.运算符需要给与特别的注意，因为它是一个有效的 ` 数字字符` ，会被优先识别为 ` 数字字面量的一部分` ，然后才是 ` 对象属性访问运算符` 。

` // 无效语法 42.toFixed(3) // SyntaxError // 下面的语法都有效 (42).toFixed(3) // 42.000 0.42.toFixed(3) // 0.420 42..toFixed(3) // 42.000 复制代码`

5、二进制浮点数最大的问题（所有遵循IEEE754规范的语言都是如此）就是小数的精度不准。（例如0.1+0.2===0.3 等于false的经典问题）

` // 判断两个小数相等 function numberCloseEnoughToEqual(n1,n2){ return Math.abs(n1-n2) < Number.EPSILON; } var a = 0.1 + 0.2; var b = 0.3; numberCloseEnoughToEqual(a,b) // true 复制代码`

6、可以使用Number.isInteger()来检测一个值是否是整数。

` Number.isInteger(42) // true Number.isInteger(42.000) // true Number.isInteger(42.3) // false // polyfill if (!Number.isInteger){ Number.isInteger = function (num){ return typeof num === 'number' && num % 1 == 0 } } 复制代码`

7、a | 0 可以将变量a中的数值转换为32位有符号整数，因为数位运算符|只适用于32位整数（它只关心32位以内的值，其他的数位将被忽略）。

8、undefined类型只有一个值，即 ` undefined` 。null类型也只有一个值，即 ` null` 。他们的名称既是类型也是值。 ` undefined指从未赋值` ， ` null只曾赋过值，但是目前没有值`

9、null是一个特殊 ` 关键字` ，不是 ` 标识符` ,我们不能将其当做变量来使用和赋值。然而undefined却是 ` 一个标识符` ,可以被当作变量来使用和赋值。

10、在非严格模式下，我们可以为全局标识符undefined赋值

` function foo (){ undefined = 2; // 非常糟糕的做法 } foo(); function foo (){ "use strict" ; undefined = 2; // TypeError } foo() 复制代码`

11、我们可以通过 ` void运算符` 即可得到undefined。表达式void __ 没有返回值，因此返回结果是undefined。void并 ` 不改变表达式的结果` ，只是 ` 让表达式不返回值` 。通常我们用void 0 来获得undefined。

` var a = 42; console.log(void a,a) // undefined 42 复制代码`

12、不是数字的数字(NaN)

* NaN是 ` 数字` 类型的

` var a = 2 / "foo" ; // NaN typeof a === 'number' // true 复制代码`

* 检测NaN

` // 全局函数isNaN判断。(但是有一个缺陷，就是当传递一个非数字的时候，isNaN也返回 true ) var a = 2 / 'foo' ; var b = 'foo' ; window.isNaN(a) // true window.isNaN(b) // true // ES6 的Number.isNaN if (!Number.isNaN){ Number.isNaN = function (n){ return ( typeof n === "number" && window.isNaN(n) ) } } // 最简单的方法，NaN是Javascript中唯一一个不等于自身的值 if (!Number.isNaN){ Number.isNaN = function (){ return n !== n; } } 复制代码`

13、零值

* Javascript中有一个常规的0和一个-0
* 对 ` 负零` 进行 ` 字符串化` 会返回'0'

` var a = 0 / -3; a; // -0 a.toString() // '0' a+ '' // '0' String(a) // '0' 复制代码`

* 将 ` 字符串` 转换为 ` 数字` ，得到的结果是准确的。

` + "-0" // -0 Number( "-0" ) // -0 JSON.parse( "-0" ) // -0 复制代码`

* 判断-0

` function isNegZero(n){ n = Number(n); return (n===0) && (1/n===-Infinity) } 复制代码`

14、ES6新加入了一个工具方法 ` Object.is()` 来判断两个值是否绝对相等（加入了NaN和-0的判断）

` if (!Object.is){ Object.is = function (v1,v2) { // 判断是否为-0 if (v1===0 && v2===0){ return 1/v1 === 1/v2 } //判断是否为NaN if (v1!==v1){ return v2!==v2; } // 其他情况 return v1 === v2; } } 复制代码`