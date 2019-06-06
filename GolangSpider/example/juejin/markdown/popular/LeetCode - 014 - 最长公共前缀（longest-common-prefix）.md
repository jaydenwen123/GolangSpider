# LeetCode - 014 - 最长公共前缀（longest-common-prefix） #

> 
> 
> 
> Create by **jsliang** on **2019-06-03 10:13:01**
> Recently revised in **2019-06-03 17:29:09**
> 
> 

为方便小伙伴们的查看，欢迎切换到不同地址~

* [文档库 LeetCode 系列 GitHub 地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2Fother-library%2FLeetCode%2Feasy )
* [文档库 LeetCode 系列 掘金 地址]( https://juejin.im/post/5cf620b4f265da1b8e708bec )
* [文档库 LeetCode 系列 小册 地址]( https://link.juejin.im?target=https%3A%2F%2Fliangjunrong.github.io%2Fother-library%2FLeetCode%2Feasy )
* 公众号地址

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21753063fc780?imageView2/0/w/1280/h/960/ignore-error/1)

## ##

**不折腾的前端，和咸鱼有什么区别**

+--------------------------------+
|              目录              |
+--------------------------------+
| [一 目录]( #chapter-one )      |
|  [二 前言]( #chapter-two       |
| )                              |
|  [三 解题]( #chapter-three     |
| )                              |
|  [3.1 解法 - 暴力破解](        |
| #chapter-three-one )           |
|  [3.2 解法 - 水平扫描](        |
| #chapter-three-two )           |
|  [3.3 解法 - 正则表达式](      |
| #chapter-three-three )         |
|  [3.4 解法 - 水平扫描](        |
| #chapter-three-four )          |
+--------------------------------+

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **难度** ：简单
* **涉及知识** ：字符串
* **题目地址** ： [leetcode-cn.com/problems/lo…]( https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Flongest-common-prefix%2F )
* **题目内容** ：

` 编写一个函数来查找字符串数组中的最长公共前缀。 如果不存在公共前缀，返回空字符串 "" 。 示例 1: 输入: [ "flower" , "flow" , "flight" ] 输出: "fl" 示例 2: 输入: [ "dog" , "racecar" , "car" ] 输出: "" 解释: 输入不存在公共前缀。 说明: 所有输入只包含小写字母 a-z 。 复制代码`

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **官方题解** ： [leetcode-cn.com/problems/lo…]( https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Flongest-common-prefix%2Fsolution%2Fzui-chang-gong-gong-qian-zhui-by-leetcode%2F )

解题千千万，官方独一家，上面是官方使用 Java 进行的题解。

小伙伴可以先自己在本地尝试解题，再看看官方解题，最后再回来看看 **jsliang** 讲解下使用 JavaScript 的解题思路。

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **解题代码** ：

` var longestCommonPrefix = function ( strs ) { if (!strs.length) { return '' ; } let shortStrLength = strs[ 0 ].length; // 最短字符串的长度 let shortStrPosition = 0 ; // 最短字符串的位置 for ( let i = 0 ; i < strs.length; i++){ if (strs[i].length < shortStrLength) { shortStrLength = strs[i].length; shortStrPosition = i; } } let result = []; for ( let i = 0 ; i < shortStrLength; i++) { for ( let j = 0 ; j < strs.length; j++) { if (strs[shortStrPosition][i] != strs[j][i]) { return result.join( '' ); } if (j === strs.length - 1 ) { result[i] = strs[shortStrPosition][i]; } } } return result.join( '' ); }; 复制代码`

* **执行测试 1** ：

* ` strs` ： ` ["flower","flow","flight"]`
* ` return` ：
` "fl" 复制代码`

* **执行测试 2** ：

* ` strs` ： ` ["dog","racecar","car"]`
* ` return` ：
` "" 复制代码`

* **LeetCode Submit** ：

` ✔ Accepted ✔ 118 / 118 cases passed ( 92 ms) ✔ Your runtime beats 86.97 % of javascript submissions ✔ Your memory usage beats 36.33 % of javascript submissions ( 35.1 MB) 复制代码`

* **知识点** ：

* ` join()` ： ` join()` 方法将一个数组（或一个类数组对象）的所有元素连接成一个字符串并返回这个字符串。 ` join()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fjoin.md )

* **解题思路** ：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1ca9e7cedf420?imageView2/0/w/1280/h/960/ignore-error/1)

**首先** ，我们进行非空判断，当它是 ` ['']` 这样子时，我们直接返回 ` ''` 。

**然后** ，我们进行第一次遍历，我们需要获取到最短字符串，因为这样我们就可以进行最短的 ` for` 遍历；我们顺带存储下位置，方便获取这个字符串。

**接着** ，我们进行双重遍历（第二/第三次遍历），将最短字符串的每个字符和其他字符串进行比对，正常情况下，我们找到不相同后，就返回结果。

**最后** ，如果我们第二/第三次遍历没有做到执行，我们返回空字符串即可。

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **解题代码** ：

` var longestCommonPrefix = function ( strs ) { if (strs.length < 2 ) { return !strs.length ? '' : strs[ 0 ]; } var result = strs[ 0 ]; for ( let i = 0 ; i < result.length; i++) { for ( let j = 1 ; j < strs.length; j++) { if (result[i] !== strs[j][i]) { return result.substring( 0 , i); } } } return result; }; 复制代码`

* **执行测试 1** ：

* ` strs` ： ` ["flower","flow","flight"]`
* ` return` ：
` "fl" 复制代码`

* **执行测试 2** ：

* ` strs` ： ` ["dog","racecar","car"]`
* ` return` ：
` "" 复制代码`

* **LeetCode Submit** ：

` ✔ Accepted ✔ 118 / 118 cases passed ( 88 ms) ✔ Your runtime beats 91.9 % of javascript submissions ✔ Your memory usage beats 46.38 % of javascript submissions ( 34.9 MB) 复制代码`

* **知识点** ：

* ` substring()` ： ` substring()` 方法将一个数组（或一个类数组对象）的所有元素连接成一个字符串并返回这个字符串。 ` substring()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fsubstring.md )

* **解题思路** ：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1caa1ad539e41?imageView2/0/w/1280/h/960/ignore-error/1)

该思路和 3.1 解法相似，理解了上面，看这幅图就是 OK 的了。

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **解题代码** ：

` var longestCommonPrefix = function ( strs ) { if (strs.length < 2 ) { return !strs.length ? '' : strs[ 0 ]; } let base = strs.shift(), joinStrs = '@' + strs.join( '@' ), regx = '@' , res = '' ; for ( let i = 0 ; i < base.length; i++){ regx += base.substring(i, i + 1 ); let matchArr = joinStrs.match( new RegExp ( ` ${regx} ` , "g" )) || []; if (matchArr.length === strs.length){ res += base.substring(i, i+ 1 ); } } return res; }; 复制代码`

* **执行测试 1** ：

* ` strs` ： ` ["flower","flow","flight"]`
* ` return` ：
` "fl" 复制代码`

* **执行测试 2** ：

* ` strs` ： ` ["dog","racecar","car"]`
* ` return` ：
` "" 复制代码`

* **LeetCode Submit** ：

` ✔ Accepted ✔ 118 / 118 cases passed ( 108 ms) ✔ Your runtime beats 45.03 % of javascript submissions ✔ Your memory usage beats 16.23 % of javascript submissions ( 35.9 MB) 复制代码`

* **知识点** ：

* ` join()` ： ` join()` 方法将一个数组（或一个类数组对象）的所有元素连接成一个字符串并返回这个字符串。 ` join()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fjoin.md )
* ` substring()` ： ` substring()` 方法将一个数组（或一个类数组对象）的所有元素连接成一个字符串并返回这个字符串。 ` substring()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fsubstring.md )
* ` RegExp` ：构造函数的原型对象。常用语一些便捷操作。 ` RegExp` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FObject%2FRegExp.md )

* **解题思路** ：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1caa4770c1040?imageView2/0/w/1280/h/960/ignore-error/1)

**首先** ，我们跟前两种方法一样进行空数组和数组长度为 1 时的判断。

**然后** ，我们将数组第一个字符串通过 ` shift()` 的形式给裁剪出来。

**接着** ，我们将数组以 ` @` 的形式拼接成字符串。因为下面我们将通过 ` @` 的形式来判断代码是否符合正则校验。

**再然后** ，通过 ` for` 循环和正则表达式判断，如果成立了，下次再判断的时候，就通过字符串拼接的形式，拓展校验规则字段 ` regx` 。

**最后** ，如果匹配返回的长度和数组总长度相等的情况下，我们就通过字符串拼接的形式修改返回值。

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **解题代码** ：

` var longestCommonPrefix = function ( strs ) { if (strs.length < 2 ) { return !strs.length ? '' : strs[ 0 ]; } return strs.reduce( ( prev, next ) => { let i = 0 ; while (prev[i] && next[i] && prev[i] === next[i]) { i++; }; return prev.slice( 0 , i); }); }; 复制代码`

* **执行测试 1** ：

* ` strs` ： ` ["flower","flow","flight"]`
* ` return` ：
` "fl" 复制代码`

* **执行测试 2** ：

* ` strs` ： ` ["dog","racecar","car"]`
* ` return` ：
` "" 复制代码`

* **LeetCode Submit** ：

` ✔ Accepted ✔ 118 / 118 cases passed ( 80 ms) ✔ Your runtime beats 96.32 % of javascript submissions ✔ Your memory usage beats 21.06 % of javascript submissions ( 35.5 MB) 复制代码`

* **知识点** ：

* ` reduce()` ： ` reduce()` 方法对数组中的每个元素执行一个由您提供的reducer函数(升序执行)，将其结果汇总为单个返回值。 ` reduce()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Freduce.md )
* ` slice()` ： ` slice()` 方法提取一个字符串的一部分，并返回一新的字符串。 ` slice()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fslice.md )

* **解题思路** ：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1caa71b9a023b?imageView2/0/w/1280/h/960/ignore-error/1)

**首先** ，这无疑是这四种思路中，写法看起来最简洁的。

**然后** ，通过 ` reduce()` ，我们可以进行一项累加操作：先比较第一项和第二项，然后找到它们共通值后，剪切并 ` return` ；再比较的时候，使用 ` return` 出来的值和第三项进行比较……依次类推

**最后** ，返回最后一次 ` return` 的值。

> 
> 
> 
> **jsliang** 广告推送：
> 也许小伙伴想了解下云服务器
> 或者小伙伴想买一台云服务器
> 或者小伙伴需要续费云服务器
> 欢迎点击 **[云服务器推广](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2Fother-library%2FMonologue%2F%25E7%25A8%25B3%25E9%25A3%259F%25E8%2589%25B0%25E9%259A%25BE.md
> )** 查看！
> 
> 

( https://link.juejin.im?target=https%3A%2F%2Fpromotion.aliyun.com%2Fntms%2Fact%2Fqwbk.html%3FuserCode%3Dw7hismrh )

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1caaa6455bea4?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fpromotion.aliyun.com%2Fntms%2Fact%2Fqwbk.html%3FuserCode%3Dw7hismrh ) ![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cab7cc87f355?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fcloud.tencent.com%2Fredirect.php%3Fredirect%3D1014%26amp%3Bcps_key%3D49f647c99fce1a9f0b4e1eeb1be484c9%26amp%3Bfrom%3Dconsole )

> 
> 
> 
> 知识共享许可协议 (
> https://link.juejin.im?target=http%3A%2F%2Fcreativecommons.org%2Flicenses%2Fby-nc-sa%2F4.0%2F
> )
> jsliang 的文档库 由 [梁峻荣](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library
> ) 采用 [知识共享 署名-非商业性使用-相同方式共享 4.0 国际 许可协议](
> https://link.juejin.im?target=http%3A%2F%2Fcreativecommons.org%2Flicenses%2Fby-nc-sa%2F4.0%2F
> ) 进行许可。
> 基于 [github.com/LiangJunron…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library
> ) 上的作品创作。
> 本许可协议授权之外的使用权限可以从 [creativecommons.org/licenses/by…](
> https://link.juejin.im?target=https%3A%2F%2Fcreativecommons.org%2Flicenses%2Fby-nc-sa%2F2.5%2Fcn%2F
> ) 处获得。
> 
>