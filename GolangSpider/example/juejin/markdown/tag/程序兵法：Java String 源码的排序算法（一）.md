# 程序兵法：Java String 源码的排序算法（一） #

**摘要: 原创出处 https://www.bysocket.com 「公众号：泥瓦匠BYSocket 」欢迎关注和转载，保留摘要，谢谢！**

这是泥瓦匠的第103篇原创

## 《程序兵法：Java String 源码的排序算法（一）》 ##

文章工程：
* JDK 1.8
* 工程名：algorithm-core-learning # StringComparisonDemo
* 工程地址：https://github.com/JeffLi1993/algorithm-core-learning

## 一、前言 ##

Q：什么是选择问题？
选择问题，是假设一组 N 个数，要确定其中第 K 个最大值者。比如 A 与 B 对象需要哪个更大？又比如：要考虑从一些数组中找出最大项？

解决选择问题，需要对象有个能力，即比较任意两个对象，并确定哪个大，哪个小或者相等。找出最大项问题的解决方法，只要依次用对象的比较（Comparable）能力，循环对象列表，一次就能解决。

那么 JDK 源码如何实现比较（Comparable）能力的呢？

## 二、java.lang.Comparable 接口 ##

![file](https://user-gold-cdn.xitu.io/2019/5/21/16ad854ea13d974f?imageView2/0/w/1280/h/960/ignore-error/1)

Comparable 接口，从 JDK 1.2 版本就有了，历史算悠久。Comparable 接口强制了实现类对象列表的排序。其排序称为自然顺序，其 ` compareTo` 方法，称为自然比较法。

该接口只有一个方法 ` public int compareTo(T o);` ，可以看出

* 入参 T o ：实现该接口类，传入对应的要被比较的对象
* 返回值 int：正数、负数和 0 ，代表大于、小于和等于

对象的集合列表（Collection List）或者数组（arrays） ，也有对应的工具类可以方便的使用：

* java.util.Collections#sort(List) 列表排序
* java.util.Arrays#sort(Object[]) 数组排序

那 String 对象如何被比较的？

## 三、String 源码中的算法 ##

String 源码中可以看到 String JDK 1.0 就有了。那么应该是 JDK 1.2 的时候，String 类实现了 Comparable 接口，并且传入需要被比较的对象是 String。对象如图：

![file](https://user-gold-cdn.xitu.io/2019/5/21/16ad854ea11723f3?imageView2/0/w/1280/h/960/ignore-error/1)

String 是一个 final 类，无法从 String 扩展新的类。从 114 行，可以看出字符串的存储结构是字符（Char）数组。先可以看看一个字符串比较案例，代码如下：

` /** * 字符串比较案例 * * Created by bysocket on 19/5/10. */ public class StringComparisonDemo { public static void main(String[] args) { String foo = "ABC" ; // 前面和后面每个字符完全一样，返回 0 String bar01 = "ABC" ; System.out.println(foo.compareTo(bar01)); // 前面每个字符完全一样，返回：后面就是字符串长度差 String bar02 = "ABCD" ; String bar03 = "ABCDE" ; System.out.println(foo.compareTo(bar02)); // -1 (前面相等,foo 长度小 1) System.out.println(foo.compareTo(bar03)); // -2 (前面相等,foo 长度小 2) // 前面每个字符不完全一样，返回：出现不一样的字符 ASCII 差 String bar04 = "ABD" ; String bar05 = "aABCD" ; System.out.println(foo.compareTo(bar04)); // -1 (foo 的 'C' 字符 ASCII 码值为 67，bar04 的 'D' 字符 ASCII 码值为 68。返回 67 - 68 = -1) System.out.println(foo.compareTo(bar05)); // -32 (foo 的 'A' 字符 ASCII 码值为 65，bar04 的 'a' 字符 ASCII 码值为 97。返回 65 - 97 = -32) String bysocket01 = "泥瓦匠" ; String bysocket02 = "瓦匠" ; System.out.println(bysocket01.compareTo(bysocket02));// -2049 （泥 和 瓦的 Unicode 差值） } } 复制代码`

运行结果如下：

` 0 -1 -2 -1 -32 -2049 复制代码`

可以看出， ` compareTo` 方法是按字典顺序比较两个字符串。具体比较规则可以看代码注释。比较规则如下：

* 字符串的每个字符完全一样，返回 0
* 字符串前面部分的每个字符完全一样，返回：后面就是两个字符串长度差
* 字符串前面部分的每个字符存在不一样，返回：出现不一样的字符 ASCII 码的差值

* 中文比较返回对应的 Unicode 编码值（Unicode 包含 ASCII）
* foo 的 ‘C’ 字符 ASCII 码值为 67
* bar04 的 ‘D’ 字符 ASCII 码值为 68。
* foo.compareTo(bar04)，返回 67 – 68 = -1
* 常见字符 ASCII 码，如图所示

![file](https://user-gold-cdn.xitu.io/2019/5/21/16ad854e9cc54292?imageView2/0/w/1280/h/960/ignore-error/1)

再看看 String 的 ` compareTo` 方法如何实现字典顺序的。源码如图：

![file](https://user-gold-cdn.xitu.io/2019/5/21/16ad854f6ef64c02?imageView2/0/w/1280/h/960/ignore-error/1)

源码解析如下：

* 第 1156 行：获取当前字符串和另一个字符串，长度较小的长度值 lim
* 第 1161 行：如果 lim 大于 0 （较小的字符串非空），则开始比较
* 第 1164 行：当前字符串和另一个字符串，依次字符比较。如果不相等，则返回两字符的 Unicode 编码值的差值
* 第 1169 行：当前字符串和另一个字符串，依次字符比较。如果均相等，则返回两个字符串长度的差值

所以要排序，肯定先有比较能力，即实现 Comparable 接口。然后实现此接口的对象列表（和数组）可以通过 Collections.sort（和 Arrays.sort）进行排序。

还有 TreeSet 使用树结构实现（红黑树），集合中的元素进行排序。其中排序就是实现 Comparable 此接口

另外，如果没有实现 Comparable 接口，使用排序时，会抛出 java.lang.ClassCastException 异常。详细看《Java 集合：三、HashSet，TreeSet 和 LinkedHashSet比较》https://www.bysocket.com/archives/195

## 四、小结 ##

上面也说到，这种比较其实有一定的弊端：

* 默认 compareTo 不忽略字符大小写。如果需要忽略，则重新自定义 compareTo 方法
* 无法进行二维的比较决策。比如判断 2 * 1 矩形和 3 * 3 矩形，哪个更大？
* 比如有些类无法实现该接口。一个 final 类，也无法扩展新的类。其也有解决方案：函数对象（Function Object）

方法参数：定义一个没有数据只有方法的类，并传递该类的实例。一个函数通过将其放在一个对象内部而被传递。这种对象通常叫做函数对象（Funtion Object）

在接口方法设计中， T execute(Callback callback) 参数中使用 callback 类似。比如在 Spring 源码中，可以看出很多设计是：聚合优先于继承或者实现。这样可以减少很多继承或者实现。类似 SpringJdbcTemplate 场景设计，可以考虑到这种 Callback 设计实现。

### 代码示例 ###

本文示例读者可以通过查看下面仓库的中: StringComparisonDemo 字符串比较案例案例：

* Github： [github.com/JeffLi1993/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJeffLi1993%2Falgorithm-core-learning )
* Gitee： [gitee.com/jeff1993/al…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fjeff1993%2Falgorithm-core-learning )

如果您对这些感兴趣，欢迎 star、follow、收藏、转发给予支持！

### 参考资料 ###

* 《数据结构与算法分析：Java语言描述（原书第3版）》
* https://en.wikipedia.org/wiki/Unicode
* https://www.cnblogs.com/vamei/tag/%E7%AE%97%E6%B3%95/
* https://www.bysocket.com/archives/2314/algorithm

### 以下专题教程也许您会有兴趣 ###

* [《程序兵法：算法与数据结构》]( https://link.juejin.im?target=https%3A%2F%2Fwww.bysocket.com%2Farchives%2F2314%2Falgorithm ) https://www.bysocket.com/archives/2314/algorithm
* [《Spring Boot 2.x 系列教程》]( https://link.juejin.im?target=https%3A%2F%2Fwww.bysocket.com%2Fspringboot )
https://www.bysocket.com/springboot
* [《Java 核心系列教程》]( https://link.juejin.im?target=https%3A%2F%2Fwww.bysocket.com%2Farchives%2F2100 )
https://www.bysocket.com/archives/2100

![](https://user-gold-cdn.xitu.io/2019/5/16/16ac0b8cde56ed82?imageView2/0/w/1280/h/960/ignore-error/1)
**（关注微信公众号，领取 Java 精选干货学习资料）**
**（添加我微信：bysocket01。加入纯技术交流群，成长技术）**