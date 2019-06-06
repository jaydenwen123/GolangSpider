# JavaScript 系列--JavaScript一些奇淫技巧的实现方法（二）数字格式化；类数组转数组 #

## 一、前言 ##

之前写了一篇文章：JavaScript 系列--JavaScript一些奇淫技巧的实现方法（一）简短的sleep函数，获取时间戳

[www.mwcxs.top/page/746.ht…]( https://link.juejin.im?target=https%3A%2F%2Fwww.mwcxs.top%2Fpage%2F746.html )

介绍了sleep函数和获取时间戳的方法。接下来我们来介绍数字格式化1234567890 --> 1,234,567,890

## 二、数字格式化 1234567890 --> 1,234,567,890 ##

### 1、普通版 ###

` // 数字格式化 1234567890 --> 1,234,567,890 function formatNumber(str){ var arr = []; var count = str.length; while (count>=3){ arr.unshift(str.slice(count - 3, count)); count -= 3; } // 如果是不是3的倍数就另外追加到上去 str.length % 3 && arr.unshift(str.slice(0, str.length % 3)); return arr.toString(); } formatNumber( '1234567890' ) 复制代码`

优点：自我感觉比网上写的一堆 for循环 还有 if-else 判断的逻辑更加清晰直白。 缺点：太普通

### 2、进阶版 ###

` // 2、进阶版 function formatNumber(str){ return str.split( "" ).reverse().reduce((prev,next,index) => { return ((index%3)? next: (next+ ',' )) + prev; }) } formatNumber( "1234567890" ); 复制代码`

优点：把 JS 的 API 玩的了如指掌 缺点：不好理解

### 3、正则版 ###

` function formatNumber(str){ return str.replace(/(?!^)(?=(\d{3})+$)/g, ',' ) } formatNumber( "1234567890" ); 复制代码`

我们来看看正则的分析：

（1）/(?!^)(?=(\d{3})+\b)/g：匹配的这个位置不能是开头(?!^)

（2）(\d{3})+：必须是1个或者多个的3个连续数字

优点：代码少，简洁。

缺点：正则表达式的位置匹配很重要，可以参考这个： [www.mwcxs.top/page/587.ht…]( https://link.juejin.im?target=https%3A%2F%2Fwww.mwcxs.top%2Fpage%2F587.html )

### 4、API版本 ###

` (1234567890).toLocaleString( 'en-us' ); (1234567890).toLocaleString(); 1234567890..toLocaleString(); 复制代码`

你可能还不知道 JavaScript 的 toLocaleString 还可以这么玩。

` 123456789..toLocaleString( 'zh-hans-cn-u-nu-hanidec' ,{useGrouping: false }); // "一二三四五六七八九" 123456789..toLocaleString( 'zh-hans-cn-u-nu-hanidec' ,{useGrouping: false }); // "一二三，四五六，七八九" new Date().toLocaleString( 'zh-hans-cn-u-nu-hanidec' ); // "二〇一九/五/二九 下午三:一五:四〇" 复制代码`

还可以使用Intl对象，

Intl 对象是 ECMAScript 国际化 API 的一个命名空间，它提供了精确的字符串对比，数字格式化，日期和时间格式化。Collator，NumberFormat 和 DateTimeFormat 对象的构造函数是 Intl 对象的属性。

new Intl.NumberFormat().format(1234567890) // 1,234,567,890 优点：简单粗暴，直接调用api

缺点：Intl兼容性不太好，不过 toLocaleString的话 IE6 都支持

前端知识点：Intl对象 和 toLocaleString的方法。

## 三、argruments 对象(类数组)转换成数组 ##

那什么是类数组？就是跟数组很像，但是他是对象，格式像数组所以叫类数组。比如：{0:a,1:b,2:c,length:3}，按照数组下标排序的对象，还有一个length的属性，有时候我们需要这种对象能调用数组下的一个方法，这时候就需要把把类数组转化成真正的数组。

### 1、普通版 ###

` var makeArray = function (arr){ var res = []; if (arr != null){ var i = arr.length; if (i == null || typeof arr == "string" ) res[0] = arr; else while (i){res[--i] = arr[i];} } return res; }; var obj = {0: 'a' ,1: 'b' ,2: 'c' ,length:3}; makeArray(obj); 复制代码`

优点：通用版本，没有任何兼容性问题 缺点：暂时没有啥缺点

### 2、进阶版 ###

` // 2、进阶版 var arr = Array.prototype.slice.call(arguments); 复制代码`

大家用过最常用的方法，至于为什么可以这么用，很多人估计也是一知半解，要搞清为什么里面的原因，我们还是从规范和源码说起。

slice.call 的作用原理就是，利用 call，将 slice 的方法作用于 arrayLike，slice 的两个参数为空，slice 内部解析使得 arguments.lengt 等于0的时候 相当于处理 slice(0) ： 即选择整个数组，slice 方法内部没有强制判断必须是 Array 类型，slice 返回的是新建的数组（使用循环取值）”，所以这样就实现了类数组到数组的转化，call 这个神奇的方法、slice 的处理缺一不可。

直接看 slice 怎么实现的吧。其实就是将 array-like 对象通过下标操作放进了新的 Array 里面，下面是源码

` // This will work for genuine arrays, array-like objects, // NamedNodeMap (attributes, entities, notations), // NodeList (e.g., getElementsByTagName), HTMLCollection (e.g., childNodes), // and will not fail on other DOM objects (as do DOM elements in IE < 9) Array.prototype.slice = function (begin, end) { // IE < 9 gets unhappy with an undefined end argument end = (typeof end !== 'undefined' ) ? end : this.length; // For native Array objects, we use the native slice function if (Object.prototype.toString.call(this) === '[object Array]' ){ return _slice.call(this, begin, end); } // For array like object we handle it ourselves. var i, cloned = [], size, len = this.length; // Handle negative value for "begin" var start = begin || 0; start = (start >= 0) ? start : Math.max(0, len + start); // Handle negative value for "end" var upTo = (typeof end == 'number' ) ? Math.min(end, len) : len; if (end < 0) { upTo = len + end; } // Actual expected size of the slice size = upTo - start; if (size > 0) { cloned = new Array(size); if (this.charAt) { for (i = 0; i < size; i++) { cloned[i] = this.charAt(start + i); } } else { for (i = 0; i < size; i++) { cloned[i] = this[start + i]; } } } return cloned; }; 复制代码`

优点：最常用的版本，兼容性强。

缺点：ie 低版本，无法处理 dom 集合的 slice call 转数组。

### 3、ES6版 ###

使用 Array.from, 值需要对象有 length 属性, 就可以转换成数组。

` var arr = Array.from(arguments); 复制代码`

扩展运算符

` var args = [...arguments]; 复制代码`

ES6 中的扩展运算符...也能将某些数据结构转换成数组，这种数据结构必须有便利器接口。

优点：直接使用内置 API，简单易维护 缺点：兼容性，使用 babel 的 profill 转化可能使代码变多，文件包变大

前端知识点：slice 方法的具体原理

【注：我是saucxs，也叫songEagle，松宝写代码，文章首发于sau交流学习社区（ [www.mwcxs.top]( https://link.juejin.im?target=https%3A%2F%2Fwww.mwcxs.top ) ），关注我们每天阅读更多精彩内容】