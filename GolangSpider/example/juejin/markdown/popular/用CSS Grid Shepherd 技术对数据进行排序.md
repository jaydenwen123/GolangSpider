# 用CSS Grid Shepherd 技术对数据进行排序 #

> 
> 
> 
> 翻译：疯狂的技术宅
> 
> 
> 
> 原文： [css-tricks.com/using-the-g…](
> https://link.juejin.im?target=https%3A%2F%2Fcss-tricks.com%2Fusing-the-grid-shepherd-technique-to-order-data-with-css%2F
> )
> 
> 
> 
> **未经许可，禁止转载！**
> 
> 

牧羊人很擅长照顾他们的羊群，为牧群带来秩序和结构。即使有几百只毛茸茸的动物，牧羊人仍然会在一天结束时将它们悉数带回农场。

而对于程序员来说，当我们在处理数据时，通常不知道这些数据是否已经被正确的过滤或者排序。尤其是当你想要在页面上按照稍微复杂一点的规则显示数据时，这就比较痛苦了。 Grid Shepherd 是一种使用 CSS Grid 帮助定位和排序的技术，完全不需要 JavaScript 的参与。

这就是本文要解决的问题。 Grid Shepherd 技术可以为我们的数据提供所需的顺序和结构，让我们更好地了解它的使用方式和应用场景。

让我们来深入研究一下。

### 用 JavaScript 排序 ###

我们首先针对农场中一系列无序的动物进行排序。想象一下牛和羊在农场中悠闲的样子。我们可以用 ` Array.prototype.sort` ( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FJavaScript%2FReference%2FGlobal_Objects%2FArray%2Fsort ) 方法以编程方式对其排序分组并展示在页面上：

` let animals = [ { name : 'Edna' , animal : 'cow' }, { name : 'Liam' , animal : 'sheep' }, { name : 'Fink' , animal : 'sheep' }, { name : 'Olga' , animal : 'cow' }, ] let sortedAnimals = animals.sort( ( a, b ) => { if (a.animal < b.animal) return -1 if (a.animal > b.animal) return 1 return 0 }) console.log(sortedAnimals) /* Returns: [ { name: 'Elga', animal: 'cow' }, { name: 'Olga', animal: 'cow' }, { name: 'Liam', animal: 'sheep' }, { name: 'Fink', animal: 'sheep' } ] */ 复制代码`

### 认识 Grid Shepherd ###

Grid Shepherd 方法能够在 **不依赖 JavaScript** 的情况下实现对数据的排序，只依靠 CSS Grid 本身就可以做到。

下面的结构与上面的 JavaScript 对象数组完全相同，只不过是在 DOM 节点中表示的。

` < main > < div class = "cow" > Edna </ div > < div class = "sheep" > Liam </ div > < div class = "sheep" > Jenn </ div > < div class = "cow" > Fink </ div > </ main > 复制代码`

CodePen上的演示： [codepen.io/Achilles_2/…]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2FAchilles_2%2Fembed%2FWWypav )

为了放养这些动物，我们必须将它们围在一个公共区域内，这就是我们 ` <main>` 元素要做的事。通过使用 ` display:grid` 设置该栅栏，我们创建了一个 [网格格式化上下文]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fcss-grid-1%2F%23grid-containers ) ，可以在其中定义每种动物应该占据的列（或行）。

`.sheep { grid-column : 1 ; }.cow { grid-column : 2 ; } 复制代码`

![羊在第一列，牛在第二列](https://user-gold-cdn.xitu.io/2019/6/5/16b2749c3458d539?imageView2/0/w/1280/h/960/ignore-error/1)

通过 ` grid-auto-flow:dense` ，每只动物都会让自己进入对应定义区域的第一个可用点。也可以用于任意数量的不同排序规则—— 只需再定义另一个列，数据就会被神奇地引导到其中。

` main display : grid ; grid-auto-flow : dense ; }.sheep { grid-column : 1 ; }.cow { grid-column : 2 ; } 复制代码`

CodePen： [codepen.io/Achilles_2/…]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2FAchilles_2%2Fembed%2FbJKqQG )

### 更专业的使用 Shepherd ###

我们还可以通过 [CSS Counters]( https://link.juejin.im?target=https%3A%2F%2Fcss-tricks.com%2Fcustom-list-number-styling%2F ) 进一步丰富这个例子。这样我们可以计算每一列中有多少只动物，并根据这个数量来有条件地设置它们的样式。

数量查询依赖于某种类型的选择器来计算其数量 —— 这对于伪类表示法 ` :nth-child(An+B [of S\ ]?)` ( https://link.juejin.im?target=https%3A%2F%2Fdrafts.csswg.org%2Fselectors%2F%23ref-for-nth-child-pseudo ) 来说会很好。但它目前仅在 Safari 中可用。这意味着我们必须用 ` :nth-of-type()` 选择器来解决这个问题。

我们需要一些新的元素类型才能实现。这可以通过 [Web 组件]( https://link.juejin.im?target=https%3A%2F%2Fcss-tricks.com%2Fguides%2Fweb-components%2F ) 实现，也可以将 HTML 元素重命名为自定义名称。即使这些元素不在 HTML 规范中，也同样适用，因为浏览器对未定义的标记使用 ` HTMLUnknownElement` ( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FAPI%2FHTMLUnknownElement ) ，这会导致他们的表现很像一个div。该文档现在看起来像这样：

` < fence > < sheep > Lisa </ sheep > < sheep > Bonnie </ sheep > < cow > Olaf </ cow > < sheep > Jenn </ sheep > </ fence > 复制代码`

现在我们可以访问自己的自定义元素类型了。当羊或牛的数量小于等于 10 时应用红色背景。

` sheep :nth-last-of-type(n+10) , sheep :nth-last-of-type(n+10) ~ sheep , cow :nth-last-of-type(n+10) , cow :nth-last-of-type(n+10) ~ cow , { background-color : red; } 复制代码`

可以通过在父元素上使用 ` counter-reset:countsheep countcow;` 并使用 ` before` 选择器来定位每个元素并计数，这样就实现了一个简单的计数器。

` sheep ::before { counter-increment : countsheep; content : counter (countsheep); } 复制代码`

你可以通过下面这个演示观察在不同的排序规则下，对动物进行添加和移除时的效果：

CodePen演示： [codepen.io/Achilles_2/…]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2FAchilles_2%2Fembed%2FYMgrpy )

Grid Shepherd 还可以和任何非有序数据一起使用：

* 根据实时增长的投票数据对选民进行分组和统计;
* 根据人们的地理位置、年龄、身高等进行分组；
* 根据规则创建层次结构。

### Shepherd 和可访问性 ###

` grid-auto-flow:dense` 不会改变网格的 DOM 结构 —— 它只是在视觉上对包含的元素重新排序。最后一个例子中会看到副作用：按字母顺序排序时， ` counter` 的数字被混淆了。更改 DOM 结构不仅会影响使用屏幕阅读器的用户，还会影响对标签遍历的效果。

### 圆满结束！ ###

本文描述了如何将一个功能强大的 CSS 布局工具（如grid）用于不符合传统布局需求的案例。我们可以看到 CSS Grid 的布局优势和 JavaScript 的动态数据处理功能是重叠的，它可以为我们提供更多的选择和功能，是我们能够随心所欲的去渲染数据。

## 欢迎关注前端公众号：前端先锋，获取前端工程化实用工具包。 ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b274aa3531111c?imageView2/0/w/1280/h/960/ignore-error/1)