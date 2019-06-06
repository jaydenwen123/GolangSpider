# 前端进阶之什么是BFC？BFC的原理是什么？如何创建BFC？ #

> 
> 
> 
> * 作者：陈大鱼头
> * github： [KRISACHAN](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FKRISACHAN )
> * 链接： [github.com/YvetteLau/S…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F15%23issuecomment-496791363
> )
> * 背景：最近高级前端工程师 **刘小夕** 在 **github** 上开了个每个工作日布一个前端相关题的 **repo** ，怀着学习的心态我也参与其中，以下为我的回答，如果有不对的地方，非常欢迎各位指出。
> 
> 
> 
> 

## 盒模型 ##

> 
> 
> 
> The CSS box model describes the rectangular boxes that are generated for
> elements in the document tree and laid out according to the visual
> formatting model.
> 
> 
> 
> CSS盒模型描述了通过 **文档树中的元素** 以及相应的 **视觉格式化模型(visual formatting model)** 所生成的矩形盒子。
> 
> 
> 

### 基础盒模型(CSS basic box model) ###

当浏览器对一个 **render tree** 进行渲染时，浏览器的渲染引擎就会根据 **基础盒模型(CSS basic box model)** ，将所有元素划分为一个个矩形的盒子，这些盒子的外观，属性由 ` CSS` 来决定。

我们在浏览器控制台输入如下代码就可以看到页面的每一个元素都是由一个矩形来包裹的，这些就是 **盒子** 。

` $$( '*' ).forEach( e => { e.style.border = '1px solid' ; }) 复制代码`

图示如下：

![](https://user-gold-cdn.xitu.io/2019/5/29/16b021a4682f916a?imageView2/0/w/1280/h/960/ignore-error/1)

### 视觉格式化模型(visual formatting model) ###

> 
> 
> 
> **CSS** 的 **视觉格式化模型(visual formatting model)** 是根据 **基础盒模型(CSS basic box
> model)** 将 **文档(doucment)** 中的元素转换一个个盒子的实际算法。
> 
> 
> 
> **官方说法就是：** 它规定了用户端在媒介中如何处理文档树( document tree )。
> 
> 

每个盒子的布局由以下因素决定：

* 盒子的尺寸
* 盒子的类型： **行内盒子 (inline)** 、 **行内级盒子 (inline-level)** 、 **原子行内级盒子 (atomic inline-level)** 、 **块级盒子 (block-level)**
* 定位： **正常流** 、 **浮动** 、 **绝对定位**
* 文档树中当前盒子的 **子元素** 或 **兄弟元素**
* **视口(viewport)** 的 **尺寸** 和 **位置**
* 盒子内部图片的 **尺寸**
* 其他某些外部因素

**视觉格式化模型(visual formatting model)** 的计算，都取决于一个矩形的边界，这个矩形，被称作是 **包含块( containing block )** 。 一般来说，(元素)生成的框会扮演它子孙元素包含块的角色；我们称之为：一个(元素的)框为它的子孙节点建造了包含块。包含块是一个相对的概念。

例子如下：

` < div > < table > < tr > < td > hi </ td > </ tr > </ table > </ div > 复制代码`

以上代码为例， ` div` 和 ` table` 都是包含块。 ` div` 是 ` table` 的包含块，同时 ` table` 又是 ` td` 的包含块，不是绝对的。

图示：(图片来自 [w3help]( https://link.juejin.im?target=http%3A%2F%2Fw3help.org%2Fzh-cn%2Fkb%2F008%2F ) )：

![](https://user-gold-cdn.xitu.io/2019/5/29/16b021a4647aaa27?imageView2/0/w/1280/h/960/ignore-error/1)

## 盒子的生成 ##

> 
> 
> 
> 盒子的生成是 **CSS视觉格式化模型** 的一部分，用于从文档元素生成盒子。盒子的类型取决于 ` CSS display` 属性。
> 
> 
> 
> **格式化上下文(formatting context)** 是定义 **盒子环境** 的规则，不同 **格式化上下文(formatting
> context)** 下的盒子有不同的表现。
> 
> 

**以下是盒子相关的概念定义：**

* 

块级元素

* 当元素的 ` display` 为 ` block` 、 ` list-item` 或 ` table` 时，它就是块级元素。

* 

块级盒子

* 块级盒子用于描述它与父、兄弟元素之间的关系。
* 每个块级盒子都会参与 **块格式化上下文（block formatting context）** 的创建。
* 每个块级元素都会至少生成一个块级盒子，即 **主块级盒子（principal block-level box）**
* 主块级盒子包含由后代元素生成的盒子以及内容，同时它也会参与定位方案。
* 一个同时是块容器盒子的块级盒子称为 **块盒子（block box）** 。

* 

匿名盒子

* 某些情况下需要进行视觉格式化时，需要添加一些增补性的盒子，这些盒子不能被 ` CSS 选择器` 选中，也就是所有可继承的 CSS 属性值都为 ` inherit` ，而所有不可继承的 CSS 属性值都为 ` initial` 。因此称为 **匿名盒子(anonymous boxes)** 。

* 

行内元素

* 当元素的 ` display` 为 ` inline` 、 ` inline-block` 或 ` inline-table` 时，它就是行内级元素。
* 显示时可以与其他行内级内容一起显示为多行。

* 

行内盒子

* 行内级元素会生成行内级盒子，该盒子同时会参与 ` 行内格式化上下文（inline formatting context）` 的创建。

* 

匿名行内盒子

* 类似于块盒子，CSS引擎有时候也会自动创建一些行内盒子。这些行内盒子无法被选择符选中，因此是匿名的，它们从父元素那里继承那些可继承的属性，其他属性保持默认值 ` initial` 。

* 

行盒子

* 行盒子由行内格式化上下文创建，用来显示一行文本。在块盒子内部，行盒子总是从块盒子的一边延伸到另一边（译注：即占据整个块盒子的宽度）。当有浮动元素时，行盒子会从向左浮动的元素的右边缘延伸到向右浮动的元素的左边缘。

* 

run-in 盒子（在CSS 2.1的标准中移除了）

* run-in盒子可以通过 ` display: run-in` 来设置，它既可以是块盒子，又可以是行内盒子，这取决于它后面的盒子的类型。

## BFC(Block formatting contexts) ##

> 
> 
> 
> **BFC** 这个概念来自于 **视觉格式化模型(visual formatting model)** 中的 **正常流(Normal
> flow)** 。
> 
> 

### 定义 ###

浮动、绝对定位元素、块容器（例如inline-blocks、table-cells、and table-captions）以及溢出而非可视的元素（除非该值已经传播到了视口）都是建立 **BFC(Block formatting contexts)** 的条件。

### 表现 ###

在 **BFC(Block formatting contexts) **中，在** 包含块** 内一个盒子一个盒子不重叠地垂直排列，两个兄弟盒子直接的垂直距离由 ` margin` 决定。 **浮动** 也是如此（虽然有可能两个盒子的距离会因为 ` floats` 而变小），除非该盒子再创建一个新的 **BFC** 。

**鱼头注：简单来说，BFC就是一个独立不干扰外界也不受外界干扰的盒子啊( /ω＼ )。**

## 块级相关的计算 ##

### 正常流中的块级与非替换元素 ###

> 
> 
> 
> ['margin-left'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fbox.html%23propdef-margin-left
> ) + ['border-left-width'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fbox.html%23propdef-border-left-width
> ) + ['padding-left'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fbox.html%23propdef-padding-left
> ) + ['width'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fvisudet.html%23propdef-width
> ) + ['padding-right'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fbox.html%23propdef-padding-right
> ) + ['border-right-width'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fbox.html%23propdef-border-right-width
> ) + ['margin-right'](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fbox.html%23propdef-margin-right
> ) = 包含块的宽度
> 
> 

上面的计算法则是基于 ** ` writing-mode: ltr` **而言，如果是别的书写顺序，则按照该顺序来计算。

如果宽度不是 **auto** 或者 **'border-left-width'+'padding-left'+'width'+'padding-right'+'border-right-width'** 的结果大于包含块的宽度，对于以下规则，被视为零。

如果只有一个值指定为'auto'，则其使用的值来自相等。

如果宽度设置为 **auto** ，则任何其他 **auto** 值变为 **0** ，并且宽度会跟着所以盒子的情况铺满。

如果 **'margin-left'** 跟 **'margin-right'** 都为 **auto** ，则会使元素相对于包含块的边缘水平居中。

### 浮动与非替换元素 ###

如果 **'margin-left'** 跟 **'margin-right'** 都为 **auto** ，则它们的具体值为 **0** 。

如果宽度为 **auto** ，则使用 **shrink-to-fit** 的宽度计算方式（ **CSS 2.2没有定义精确的算法** ）。

然后 **shrink-to-fit** 大概的计算方式则是： **min(max(preferred minimum width, available width), preferred width)** 。

### 绝对定位与非替换元素 ###

> 
> 
> 
> 'left' + 'margin-left' + 'border-left-width' + 'padding-left' + 'width' +
> 'padding-right' + 'border-right-width' + 'margin-right' + 'right' = 包含块的宽度
> 
> 
> 

如果 **'left'** ， **'width'** 和 **'right'** 都是 **'auto'** ，则首先将 **'margin-left'** 和 **'margin-right'** 的 **'auto'** 值设置为 **0** 。

如果 **'left'** ， **'width'** 和 **'right'** 都不是 **'auto'** ，则按照实际值来算。

如果 **'margin-left'** 跟 **'margin-right'** 都为 **0** ，则根据 **'left'** ， **'width'** 和 **'right'** 的值是否是 **'auto'** 来计算。 如果 **一个方向值** ， **'width'** 的值是 **'auto'** ，而 **'另一个一个方向值'** 不是，则宽度使用 **shrink-to-fit** 算法计算。如果一个值为**'auto'**而另外两个值不算，则该值使用 **shrink-to-fit** 来计算。

上面的计算法则是基于 **` writing-mode: ltr`** 而言，如果是别的书写顺序，则按照该顺序来计算。

**鱼头注：这里特别说明一点，在 [MDN]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FGuide%2FCSS%2FBlock_formatting_context ) 中依然把flexbox跟gridbox 算在 BFC中，但在最新的规范里，它们已经从BFC中分离了出去，成为独立的一个CSS模块，内容如下：**

* [CSS Flexible Box Layout Module Level 1]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fcss-flexbox-1%2F )
* [CSS Grid Layout Module Level 2]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fcss-grid-2%2F )

如果你、喜欢探讨技术，或者对本文有任何的意见或建议，你可以扫描下方二维码，关注微信公众号“ **鱼头的Web海洋** ”，随时与鱼头互动。欢迎！衷心希望可以遇见你。

![](https://user-gold-cdn.xitu.io/2019/5/18/16aca8a03abb7c13?imageView2/0/w/1280/h/960/ignore-error/1)