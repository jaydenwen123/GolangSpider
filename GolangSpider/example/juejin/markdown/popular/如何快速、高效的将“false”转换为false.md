# 如何快速、高效的将“false”转换为false #

## 问题提出 ##

> 
> 
> 
> 背景： 从cookie中读的键值为'false'，当把字符串作为判断条件时，问题就来了
> 
> 

在JS的世界里， 0、-0、null、""、false、undefined 或 NaN，这些都可以自动转化为布尔的 false，那么字符串的"false"却不等于false呢，if("false") 来判断的话，是等于true的。

根据W3C的解释：

` var myBoolean= new Boolean (); //下面的所有的代码行均会创建初始值为 false 的 Boolean 对象。 var myBoolean= new Boolean (); var myBoolean= new Boolean ( 0 ); var myBoolean= new Boolean ( null ); var myBoolean= new Boolean ( "" ); var myBoolean= new Boolean ( false ); //不带单引号的是false var myBoolean= new Boolean ( NaN ); //下面的所有的代码行均会创初始值为 true 的 Boolean 对象： var myBoolean= new Boolean ( 1 ); var myBoolean= new Boolean ( true ); var myBoolean= new Boolean ( "true" ); var myBoolean= new Boolean ( "false" ); //带单引号的字符串false最终等于true var myBoolean= new Boolean ( "Bill Gates" ); 复制代码`

因此，如果从cookie或者后台返回数据为字符串的'false'时，是不能直接作为判断条件的。

## 解决方案 ##

* 三目运算符转换

` let bool = result === 'false' ? false : true 复制代码`

* 索引

` let bool = { 'true' : true , 'false' : false }; bool[ v ] !== undefined ? bool[ v ] : false ; 复制代码`

* JSON.parese

这个应该是最方便、简洁的方法了

` JSON.parse( 'false' ) // false JSON.parse( 'true' ) //true 复制代码`

* 后端配合 ，不要使用字符串的'ture'后者'false'，使用0,1代替

` let target = '0' let bool = !!+target 复制代码`

【讨论向】这是目前找到的较好的解决方案，如果有更好、更优雅的方式，欢迎贴过来

## 参考资料 ##

[连接地址]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fq%2F1010000013892022%2Fa-1020000013893195 )