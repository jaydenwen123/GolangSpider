# LeetCode - 020 - 有效的括号（valid-parentheses） #

> 
> 
> 
> Create by **jsliang** on **2019-06-04 11:39:30**
> Recently revised in **2019-06-04 14:41:25**
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

+-----------------------------+
|            目录             |
+-----------------------------+
| [一 目录]( #chapter-one )   |
|  [二 前言]( #chapter-two    |
| )                           |
|  [三 解题]( #chapter-three  |
| )                           |
+-----------------------------+

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **难度** ：简单
* **涉及知识** ：栈、字符串
* **题目地址** ： [leetcode-cn.com/problems/va…]( https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Fvalid-parentheses%2F )
* **题目内容** ：

` 给定一个只包括 '(' ， ')' ， '{' ， '}' ， '[' ， ']' 的字符串，判断字符串是否有效。 有效字符串需满足： 左括号必须用相同类型的右括号闭合。 左括号必须以正确的顺序闭合。 注意空字符串可被认为是有效字符串。 示例 1: 输入: "()" 输出: true 示例 2: 输入: "()[]{}" 输出: true 示例 3: 输入: "(]" 输出: false 示例 4: 输入: "([)]" 输出: false 示例 5: 输入: "{[]}" 输出: true 复制代码`

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **官方题解** ： [leetcode-cn.com/problems/va…]( https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Fvalid-parentheses%2Fsolution%2Fyou-xiao-de-gua-hao-by-leetcode%2F )

解题千千万，官方独一家，上面是官方使用 Java 进行的题解。

小伙伴可以先自己在本地尝试解题，再看看官方解题，最后再回来看看 **jsliang** 讲解下使用 JavaScript 的解题思路。

* **解题代码** ：

` var isValid = function ( s ) { let judge = { '[' : ']' , '(' : ')' , '{' : '}' , } let parameter = s.split( '' ); let stack = []; for ( let i = 0 ; i < parameter.length; i++) { if (judge[stack[stack.length - 1 ]] === parameter[i]) { stack.pop(); } else { stack.push(parameter[i]); } } if (stack.length == 0 ) { return true ; } return false ; }; 复制代码`

* **执行测试** ：

* ` s` ： ` ()[]{}`
* ` return` ： ` true`

* **LeetCode Submit** ：

` ✔ Accepted ✔ 76 / 76 cases passed ( 76 ms) ✔ Your runtime beats 96.48 % of javascript submissions ✔ Your memory usage beats 66.23 % of javascript submissions ( 33.8 MB) 复制代码`

* **知识点** ：

* ` split()` ： ` split()` 方法使用指定的分隔符字符串将一个 String 对象分割成字符串数组，以将字符串分隔为子字符串，以确定每个拆分的位置。 ` split()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fsplit.md )
* ` push()` ： ` push()` 方法将一个或多个元素添加到数组的末尾，并返回该数组的新长度。 ` push()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fpush.md )
* ` pop()` ： ` pop()` 方法从数组中删除最后一个元素，并返回该元素的值。此方法更改数组的长度。 ` pop()` 详细介绍 ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FLiangJunrong%2Fdocument-library%2Fblob%2Fmaster%2FJavaScript-library%2FJavaScript%2FFunction%2Fpop.md )

* **解题思路** ：

**首先** ，我们来想下平时玩的翻牌游戏：

游戏屏幕中有 4 * 4 共 16 张牌。

如果在同一个回合中（一个回合能翻 2 次牌），我们翻出来相同的两张牌，就把它摊开（消掉）；

如果在同一个回合中翻到两张不同的牌，我们就要把它覆盖回去。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b217610fcef8f6?imageView2/0/w/1280/h/960/ignore-error/1)

这时候，机智如你，有没有想过要一个作弊器，用来记录我们翻过的牌？！

就好比：我们先翻到红包，再翻到炸弹，下一次翻到红包的时候，我们就打开一开始翻到红包的地方……

OK，这么讲，小伙伴们大致清楚了。

**现在** ，我们有一个字符串： ` (()[]{})` ，我们要判断它是否是有效的括号，怎么判断呢：

* 这些括号必须成对出现，例如： ` ()` 、 ` []` 、 ` {}` 。
* 这些括号出现的情况不能乱序，例如： ` ([)]` 、 ` {[}]`

**同时** ，我们初始化下数据：

* ` judge` :

` { '[' : ']' , '(' : ')' , '{' : '}' , } 复制代码`

* ` parameter` ： ` [ '(', '(', ')', '[', ']', '{', '}', ')' ]`
* ` stack` ： ` []`

**然后** ，我们是不是该开始使用我们的作弊器了：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b217638c8b1700?imageView2/0/w/1280/h/960/ignore-error/1)

在这个作弊器中，我们将碰到的括号记录起来。

如果碰到下一个括号，是我们想要的类型，那么就消掉最上层的括号；

如果碰到下一个括号，不是我们想要的类型，那么就把它放到数组的最上层。

* 遍历第一次， ` stack` 末尾是空的，所以我们执行 ` push()` 操作， ` stack` ： ` ['(']`
* 遍历第二次， ` stack` 末尾是 ` '('` ，通过 ` judge` 转换就是 ` ')'` ，而在这个位置的 ` parameter[i]` 是 ` '('` ，两者不相同，所以我们还是执行 ` push()` 操作， ` stack` ： ` ['(', '(']`
* 遍历第三次， ` stack` 末尾是 ` '('` ，通过 ` judge` 转换就是 ` ')'` ，而在这个位置的 ` parameter[i]` 是 ` ')'` ，两者相同，所以我们执行 ` pop()` 操作，将数组的末尾给删掉， ` stack` ： ` ['(']`
* ……以此类推，最后我们的 ` stack` 变成 ` []` 空数组。

**最后** ，我们根据 ` stack` 是否为空数组，来进行判断这个字符串是不是有效数组。

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

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21757b8307774?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fpromotion.aliyun.com%2Fntms%2Fact%2Fqwbk.html%3FuserCode%3Dw7hismrh ) ![](https://user-gold-cdn.xitu.io/2019/6/4/16b2175b067e9c2f?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fcloud.tencent.com%2Fredirect.php%3Fredirect%3D1014%26amp%3Bcps_key%3D49f647c99fce1a9f0b4e1eeb1be484c9%26amp%3Bfrom%3Dconsole )

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