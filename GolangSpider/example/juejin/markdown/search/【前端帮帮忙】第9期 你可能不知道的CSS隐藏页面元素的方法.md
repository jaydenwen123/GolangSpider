# 【前端帮帮忙】第9期 你可能不知道的CSS隐藏页面元素的方法 #

> 
> 
> 
> 隐藏元素的方法有很多，使用的时候还是要根据实际项目的需求来选择最合适的。今天我们一起来学习一下这方面相关的知识点。
> 
> 

## 1. opacity ##

`.hide-opacity { opacity : 0 ; } 复制代码`

通过下面的gif图，我们可以总结 ` opacity` 隐藏元素有几个特点：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3b1fda612e3?imageslim)

* 只是视觉上的隐藏，隐藏元素的依然占据着空间，影响其他元素的布局
* 依然能够响应用户的交互
* 通过DOM依然可以获取该元素，可以响应DOM事件
* 其子孙元素即使重新设置了 ` opacity: 1` 也无法显示
* 元素和它所有的内容会被读屏软件阅读（没有测试过）

## 2. visibility ##

`.hide-visibility { visibility : hidden; } 复制代码`

通过下面的gif动图，同样我们可以总结出 ` visibility` 隐藏元素的特点：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3b57736c85f?imageslim)

* 隐藏元素的依然占据着空间，影响其他元素的布局
* 不会响应任何用户交互
* 通过DOM依然可以获取该元素，无法响应DOM事件
* 其子孙元素可以通过重新设置 ` visibility: visible` 来显示
* 元素在读屏软件中也会被隐藏（没有测试过）

## 3. display ##

是真正意义上的隐藏元素。

`.hide-display { display : none; } 复制代码`

通过下面的gif动图，我们可以总结出 ` display: none` 隐藏元素的特点：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3b987e6c897?imageslim)

* 真正意义上的隐藏，元素不会占据任何空间
* 用户无法与其进行直接的交互
* 通过DOM依然可以获取到该元素
* 其子孙元素即使重新设置 ` display: block` 也无法显示
* 读屏软件无法读到该元素的内容（没有测试过）
* ` transition` 动画会失效

## 4. hidden ##

HTML5新增的 ` hidden` 属性，可以直接隐藏元素。

` <div hidden> 我是被隐藏的元素 </div> 复制代码`

特点：

**跟 ` display` 表现一致。**

## 5. position ##

`.hide-position { position : absolute; top : - 9999px ; left : - 9999px ; } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3be3c0e9132?imageView2/0/w/1280/h/960/ignore-error/1)

特点：

* 视觉上的隐藏，实际显示在屏幕的可视区之外，不会占据空间，不影响其他元素的布局
* 用户无法与其进行直接的交互
* 元素的内容可以被读屏软件读取（没有测试过）
* 通过DOM依然可以获取到该元素
* 其子孙元素无法通过重新设置对应的属性来显示

## 6. clip-path ##

通过裁剪元素来实现隐藏。

`.hide-clip { clip-path : polygon (0 0, 0 0, 0 0, 0 0); } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3c6ec636b8b?imageView2/0/w/1280/h/960/ignore-error/1)

特点：

* 视觉上的隐藏，隐藏元素的依然占据着空间，影响其他元素的布局
* 用户无法与其进行直接的交互
* 元素的内容可以被读屏软件读取（没有测试过）
* 通过DOM依然可以获取到该元素
* 其子孙元素无法通过重新设置对应的属性来显示

## 7. overflow ##

通过设置元素的宽高为 ` 0` 来隐藏元素。

`.hide-overflow { width: 0; height: 0; overflow: hidden; } 复制代码`

必须加上 ` overflow: hidden` ，否则其子孙元素依然可以显示，下面的动图可以说明：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3ca25c1c0b1?imageslim)

特点：

* 视觉上的隐藏，隐藏元素的不会占据任何空间，不会影响其他元素的布局
* 用户无法与其进行直接的交互
* 元素的内容可以被读屏软件读取（没有测试过）
* 通过DOM依然可以获取到该元素
* 其子孙元素无法通过重新设置对应的属性来显示

## 8. transform ##

`.hide-transform { transform: translate(-9999px, -9999px); } 复制代码`

或者

`.hide-transform { transform: scale(0); } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3cf08454f33?imageslim)

特点：

* 视觉上的隐藏，隐藏元素的依然占据着空间，影响其他元素的布局
* 用户无法与其进行直接的交互
* 元素的内容可以被读屏软件读取（没有测试过）
* 通过DOM依然可以获取到该元素
* 其子孙元素无法通过重新设置对应的属性来显示

## 其他 ##

如果是纯文本的隐藏，可以设置

`.hide-text { text-indent: -9999px; } 复制代码`

或者

`.hide-text { font-size: 0; } 复制代码`

还有一个是 ` 无障碍设计规范` 里面的：

` <div aria-hidden= "true" ></div> 复制代码`

## 差异性 ##

上面简单的罗列了8中隐藏元素的方式，其实给我们视觉上的效果，这些方法都可以让元素不可见（也就是我们所说的隐藏）。然而，屏幕并不是唯一的输出机制，比如说屏幕上看不见的元素（隐藏的元素），其中一些依然能够被读屏软件阅读出来（因为读屏软件依赖于 [可访问性树]( https://link.juejin.im?target=https%3A%2F%2Fallyjs.io%2Fconcepts.html%23Accessibility-tree ) 来阐述）。为了消除它们之间的歧义，我们将其归为三大类：

* 完全隐藏
* 视觉上的隐藏
* 语义上的隐藏

三种类型的隐藏总结下来如下表所示：

+--------------+--------+--------------------------+
|  可见性状态  | 屏幕上 | 可访问性树（读屏软件等） |
+--------------+--------+--------------------------+
| 完全隐藏     | 隐藏   | 隐藏                     |
| 视觉上的隐藏 | 隐藏   | 可见                     |
| 语义上的隐藏 | 可见   | 隐藏                     |
+--------------+--------+--------------------------+

### 完全隐藏 ###

针对上面所列的8种方法，能够实现完全隐藏的只有下面这3种：

* display: none
* visibility: hidden
* HTML5新增的hidden属性

### 视觉上的隐藏 ###

* opacity
* position
* transform
* clip-path
* overflow

### 语义上的隐藏 ###

* aria-hidden="true"

## 其他区别 ##

+------------+----------------+----------------+-----------------+
|  隐藏方式  | 占据原来的空间 | 直接跟用户交互 | 直接响应DOM事件 |
+------------+----------------+----------------+-----------------+
| opacity    | 是             | 是             | 是              |
| visibility | 是             | 否             | 否              |
| display    | 否             | 否             | 否              |
| hidden     | 否             | 否             | 否              |
| position   | 否             | 否             | 否              |
| clip-path  | 是             | 否             | 否              |
| overflow   | 否             | 否             | 否              |
| transform  | 是             | 否             | 否              |
+------------+----------------+----------------+-----------------+

## 最后 ##

感谢您的阅读，希望对你有所帮助。由于本人水平有限，如果文中有不当的地方烦请指正，感激不尽。

## 关注 ##

欢迎大家关注我的公众号 ` 前端帮帮忙` ，一起交流学习，共同进步。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3d653f6c322?imageView2/0/w/1280/h/960/ignore-error/1)

参考：

[【译】用 CSS 隐藏页面元素的 5 种方法]( https://link.juejin.im?target=https%3A%2F%2F75team.com%2Fpost%2Ffive-ways-to-hide-elements-in-css.html )