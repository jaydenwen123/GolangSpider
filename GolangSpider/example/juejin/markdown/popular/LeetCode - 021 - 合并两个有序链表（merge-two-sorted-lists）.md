# LeetCode - 021 - 合并两个有序链表（merge-two-sorted-lists） #

> 
> 
> 
> Create by **jsliang** on **2019-06-05 08:37:00**
> Recently revised in **2019-06-05 17:31:37**
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
|  [3.1 官方题解](               |
| #chapter-three-one )           |
|  [3.2 解题代码](               |
| #chapter-three-two )           |
|  [3.3 执行测试](               |
| #chapter-three-three )         |
|  [3.4 LeetCode Submit](        |
| #chapter-three-four )          |
|  [3.5 知识补充](               |
| #chapter-three-five )          |
|  [3.6 解题思路](               |
| #chapter-three-six )           |
|  [四 总结]( #chapter-four      |
| )                              |
+--------------------------------+

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **难度** ：简单
* **涉及知识** ：链表
* **题目地址** ： [leetcode-cn.com/problems/me…]( https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Fmerge-two-sorted-lists%2F )
* **题目内容** ：

` 将两个有序链表合并为一个新的有序链表并返回。 新链表是通过拼接给定的两个链表的所有节点组成的。 示例： 输入：1->2->4, 1->3->4 输出：1->1->2->3->4->4 复制代码`

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

[leetcode-cn.com/problems/me…]( https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Fmerge-two-sorted-lists%2Fsolution%2Fhe-bing-liang-ge-you-xu-lian-biao-by-leetcode%2F )

解题千千万，官方独一家，上面是官方使用 Java 进行的题解。

小伙伴可以先自己在本地尝试解题，再看看官方解题，最后再回来看看 **jsliang** 讲解下使用 JavaScript 的解题思路。

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

` var mergeTwoLists = function ( l1, l2 ) { var mergedHead = { val : -1 , next : null }, crt = mergedHead; while (l1 && l2) { if (l1.val > l2.val) { crt.next = l2; l2 = l2.next; } else { crt.next = l1; l1 = l1.next; } crt = crt.next; } crt.next = l1 || l2; return mergedHead.next; }; 复制代码`

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

* **参数 1** ：

` l1 = { val : 1 , next : { val : 2 , next : { val : 4 , next : null } } } 复制代码`

* **参数 2** ：

` l2 = { val : 1 , next : { val : 3 , next : { val : 5 , next : null } } } 复制代码`

* **返回值** ：

` { val : 1 , next : { val : 1 , next : { val : 2 , next : { val : 3 , next : { val : 4 , next : { val : 5 , next : null , } } } } } } 复制代码`

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

` ✔ Accepted ✔ 208 / 208 cases passed ( 84 ms) ✔ Your runtime beats 98.22 % of javascript submissions ✔ Your memory usage beats 87.2 % of javascript submissions ( 35 MB) 复制代码`

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

在这道题的解题知识中，存在一个知识点：链表。

然后因为 **jsliang** 跟小伙伴们一样，也是 **算法与数据结构** 的菜鸟，所以网上找了篇文章：

js实现链表： [www.cnblogs.com/EganZhang/p…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2FEganZhang%2Fp%2F6594830.html )

**jsliang** 碰到这种结构的时候，也是挠头抓耳，然后百度找到这篇文章，基本的结构看懂了，所以在这推荐给小伙伴们。

当然， **jsliang** 立马下单了一本书 《学习 JavaScript 数据结构与算法》，有没用不知道，先买了再说，后面 **jsliang** 会补上相关的知识点，小伙伴们先看下上面大佬写的文章咯~

### ###

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26fddee838b7b?imageView2/0/w/1280/h/960/ignore-error/1)

## ##

> 
> 
> 
> [返回目录]( #chapter-one )
> 
> 

这样，我们就完成了 21 题的题解，感觉理解出来的话，其实是挺容易实现的。

如果小伙伴们还是有点懵，最好多打印几个 ` console.log` ，就清楚它是怎么运行的了！

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

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26fe1addeb165?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fpromotion.aliyun.com%2Fntms%2Fact%2Fqwbk.html%3FuserCode%3Dw7hismrh ) ![](https://user-gold-cdn.xitu.io/2019/6/5/16b26fe47168d3a5?imageView2/0/w/1280/h/960/ignore-error/1) ( https://link.juejin.im?target=https%3A%2F%2Fcloud.tencent.com%2Fredirect.php%3Fredirect%3D1014%26amp%3Bcps_key%3D49f647c99fce1a9f0b4e1eeb1be484c9%26amp%3Bfrom%3Dconsole )

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