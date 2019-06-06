# 通过动图学习 CSS Flex #

> 
> 
> 
> 翻译：疯狂的技术宅
> 
> 
> 
> 原文： [www.freecodecamp.org/news/the-co…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.freecodecamp.org%2Fnews%2Fthe-complete-flex-animated-tutorial%2F
> )
> 
> 
> 
> **声明：未经允许严禁转载**
> 
> 

![CSS Flex – Animated Tutorial](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa0c624947a?imageslim)

> 
> 
> 
> 如果一张图片胜过千言万语 —— **那么动画呢？** Flex 无法通过文字或静态图像有效地完全解释。为了巩固你对flex的了解，我制作了这些 **动画**
> 演示。
> 
> 

![img](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa0b4521c15?imageslim)

注意 **overflow: hidden** 行为类型是默认值，因为 **flex-wrap** 还未设置。

为了获得更好的想法，你可以在这个页面上去尝试一下 **Flex Layout Editor** ( https://link.juejin.im?target=http%3A%2F%2Fwww.csstutorial.org%2Fflex-both.html ) 。

按 **默认** flex不会包装你的内容。它的工作原理很像 **overflow: hidden** 。

可能你在学习 flex 时第一个接触到的就是 **flex-wrap** 。

## Flex Wrap ##

让我们添加 **flex-wrap:wrap** 来看看它是如何改变 flex 项的行为的。

基本上，它只会扩展容器高度并将物品包裹起来。

**注意：**即便是未指定容器得高度（ **auto/unset** ）仍然会这样。

![wrap](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa0d139b37c?imageslim)

wrap

如果你有一些内容大小未知且数量也未知的项目，并且希望在屏幕上全部显示它们时，这是一种常见模式。

可以用 **flex-direction: row-reverse** 来反转项目的实际顺序。

![row-reverse](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa0b24070cd?imageslim)

row-reverse

这可以用于需要 **从右到左** 的顺序阅读的内容。

你可以 " **float:right** " **所有** 与 **flex-end** 在同一行上的项目。

这与 **row-reverse** 不同，因为它 **保留了项目的顺序。**

## Justify Content ##

**justify-content** 属性负责 flex 项目的 **水平对齐** 。

它看起来很像前面的例子……除了项目的顺序。

![flex-end](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa0b0ba5f55?imageslim)

flex-end

在以下示例中（ **justify-content: center** ），所有项目将自然地聚集到父容器的中心 —— 无论其宽度怎样。它与 **position: relative; margin: auto** 相似。

![center](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa0b28805a5?imageslim)

center

**Space between** 意味着所有 内部 项目之间有空格：

![space-between](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa1af2133fb?imageslim)

space-between

下面这个 **似乎** 与上面的完全相同。那是因为它的内容同样是整个字母表。如果用较少的弹性项目，效果会更明显。它们的不同之处是处于 角落的项目的外边距 。

使用 **space-between** 属性（上图） **处于角落的项目没有外边距** 。

使用 **space-around** 属性（下图） **所有项目的边距相同** 。

![space-around](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa20c5aa7f0?imageslim)

space-around 下面这个动画是相同的例子，只不过 middle 元素更宽一些。

![img](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa243710126?imageslim)

space-around

尽管你在前面看到了这些演示，但你仍然需要在自己的环境中去尝试 flex，这样才能是你真正理解布局。这也是我决定制作本教程的原因。这些动画受限于项目大小。你尝试的结果可能会因内容的具体情况而异。

## **对齐内容** ##

上面的所有例子都涉及 **justify-content** 属性。不过即便涉及到 **自动折行** ，你也可以在 flex 中进行 垂直对齐 。

属性 **justify-content** （上面的所有示例）和 **align-content** （下面）采用完全相同的值。它们仅在两个不同的方向上对齐 —— **相对于存储在柔性容器中的项目** 的垂直和水平方向上。

接下来探讨 flex 如何处理 **垂直对齐：**

![align-content:space-evenly](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa1a781e55f?imageslim)

align-content:space-evenly

关于 **space-evenly** 的一些现象：

* Flex **自动分配足够的垂直空间** 。
* **项目行** 与相等的垂直边距空间对齐。

当然，你仍然可以修改父级的高度，并且所有内容仍然可以正确对齐。

## 实际应用中的情况 ##

在实际布局中，你不会有一长串的文字，你将会使用一些独特的内容元素。到目前为止我只简单演示了动画中的 flex 是如何工作的。

当涉及到实际布局时，你可能希望对较少同时更大的项目使用 flex。就像真正网站上的那些内容一样。

我们来看几个想法......

### 均匀排列 ###

对于 **align-content** 和 **justify-content** 使用 **space-evenly** 会对具有5个正方形的一组项目产生以下影响：

![奇数项目的效果不是很好](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa256540767?imageslim)

奇数项目的效果不是很好

当涉及 flex 的 开箱即用的响应 区域时，首先要确保尽可能使项目的宽度保持相同。

请注意，因为此示例中的项目数为 **奇数个** （5），所以这种情况不会产生你想要的那种 理想的 响应效果。使用 **偶数** 数字可以解决这个微妙的问题。

现在，考虑将相同的 flex 属性用在 **偶数** 个项目上：

![以更自然的方式响应偶数个项目](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa2b20d929d?imageslim)

以更自然的方式响应偶数个项目

使用 **偶数** 个项目，你可以实现更清晰的 响应式 缩放，而无需用 **CSS Grid** 或 **JavaScript magic** 。

十多年来，在布局设计中垂直居中的项目已成为一个巨大的问题。

最后用 **flex** 解决了。 （ 呃......你也可以用 **css grid** 来解决 。）

但是在 flex 中，在两个维度中使用 **space-evenly** 值会对内容自动调整， 即使项目的高度可变：

![完美的对多个不同高度的项目垂直对齐](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa2b32909c1?imageslim)

完美的对多个不同高度的项目垂直对齐

以上是对未来10年最常用的响应式 flex 的描述（开个玩笑😆）。

如果你正在学习flex，你会发现这通常是最有用的一组 flex 属性。

最后，下面的动画演示了所有可能会用到的值：

### flex-direction: row; justify-content: [value]; ###

![img](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa313b2ddd5?imageslim)

### flex-direction: column; justify-content: [value]; ###

![img](https://user-gold-cdn.xitu.io/2019/6/6/16b2baa35a6b8d6c?imageslim)

我建议你在 **CSS grid** 中使用这些类型的 flex 项目。 （你用的越多就会越明白 **grid + flex** 。）不过使用 **flex-only** 布局也没有任何问题。

### 要明确指定元素的大小 ###

如果不这样做，一些 flex 缩放将无法正常工作。

相应地使用 **min-width** ， **max-width** 和 **width** / **height** 。

这些属性可以对整个 flex 可伸缩性产生影响。

以上！希望你能够喜欢这篇文章。

## 欢迎关注前端公众号：前端先锋，获取前端工程化实用工具包。 ##

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bab79c121221?imageView2/0/w/1280/h/960/ignore-error/1)