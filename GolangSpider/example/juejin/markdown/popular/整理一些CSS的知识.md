# 整理一些CSS的知识 #

最近在翻《CSS权威指南》，一些零散的知识点平时不太注意，这里记录一下。

### CSS属性display ###

display指定了元素的显示类型，它包含两类基础特征，用于指定元素怎样生成盒模型——外部显示类型定义了元素怎样参与流式布局的处理，内部显示类型定义了元素内子元素的布局方式。

` <display-outside>（指定了元素的外部显示类型，实际上就是其在流式布局中的角色） = block | inline | run-in <display-inside> （指定了元素的内部显示类型，它们定义了元素内部内容的格式化上下文的类型）= flow | flow-root | table | flex | grid | ruby <display-listitem> （元素的外部显示类型变为 block 盒，并将内部显示类型变为多个 list-item inline 盒）= <display-outside>? && [ flow | flow-root ]? && list-item <display-internal> （用来定义这些“内部”显示类型，只有在特定的布局模型中才有意义）= table-row-group | table-header-group | table-footer-group | table-row | table-cell | table-column-group | table-column | table-caption | ruby-base | ruby-text | ruby-base-container | ruby-text-container <display-box> （是否完全生成显示盒）= contents | none <display-legacy> （CSS 2 对于 display 属性使用单关键字语法）= inline-block | inline-list-item | inline-table | inline-flex | inline-grid 复制代码`

#### 候选样式表 ####

` < link rel = "stylesheet" href = "day.css" title = "Default Day" > < link rel = "alternate stylesheet" href = "night.css" title = "Night" > 复制代码`

rel（关系）中可以指定候选样式表，默认使用第一个样式表。刚好看到这篇文章 [link rel=alternate网站换肤功能最佳实现]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhangxinxu.com%2Fwordpress%2F2019%2F02%2Flink-rel-alternate-website-skin ) ，自己实现了一下：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20b293c6df04f?imageslim)

#### 属性选择符 ####

* [ foo|='bar' ] 选择的元素有foo属性，且其值以bar和一个英文破折号开头，或者值就是bar本身
* [ foo~='bar' ] 选择的元素有foo属性，且其值是包含bar这个词的一组词
* [ foo*='bar' ] 选择的元素有foo属性，且其值包含子串bar
* [ foo^='bar' ] 选择的元素有foo属性，且其值以bar开头
* [ foo$='bar' ] 选择的元素有foo属性，且其值以bar结尾

这里的应用在于，如果我们在开发一个CSS框架或者模式库，定义一个类 ` 'btn btn-small btn-arrow btn-active'` 显得冗余，我们可以直接使用 ` 'btn-small-arrow-active'`

` < button class = "btn-small-arrow-active" > </ button > *[class|="btn"][class*="-arrow"]:after {content: '▼'} 复制代码`

#### 特指度 ####

一个元素可能被两个或多个规则匹配，其中只有一个规则能胜出，特指度能够解决冲突。一个特指度由四部分构成，比如 ` 0, 0, 0, 0`

* 选择符的每个ID属性值加 ` 0, 1, 0, 0`
* 选择符的每个类属性值、属性选择或伪类加 ` 0, 0, 1, 0`
* 选择符中的每个元素和伪元素加 ` 0, 0, 0, 1`
* 连结符和通用选择符不增加特指度

` h1 {} /* 0, 0, 0, 1 */ p em {} /* 0, 0, 0, 2 */.grape {} /* 0, 0, 1, 0 */ p.bright em.dark {} /* 0, 0, 2, 2 */ li #answer {} /* 0, 1, 0, 1 */ 复制代码`

* ` !important` 重要规则始终胜出
* 声明冲突，按照特指度排序
* 权重、来源和特指度一致，按照位置靠后权重更大

#### 伪元素选择器 ####

` el ::first-letter {} /*装饰首字母*/ el ::first-line {} /*装饰首行*/ el ::before {} /*前置内容元素*/ el ::after {} /*后置内容元素*/ 复制代码`

#### 超链接伪类 ####

` a :link {} a :visited {} a :focus {} a :hover {} a :active {} 复制代码`

* 注意顺序，因为这里选择符的特指度一致，所以最后一个匹配的规则将胜出。

#### 属性值 ####

可以使用样式获取的元素上的HTML属性值（实际兼容浏览器很少），例如：

` <!--css--> p::before {content: "[" attr(id) "]"} <!--html--> < p id = "leadoff" > This is the first paragraph </ p > <!--显示结果--> [leadoff]This is the first paragraph 复制代码`

#### 角度单位 ####

* deg 度数，完整圆周是360度
* grad 百分度，完整圆周是400百分度
* rad 弧度，完整圆周是2π
* turn 全周，一个完整的圆周是一圈，在旋转动画中最有用

#### 自定义值 ####

` :root { /* 或者html */ --base-color : #639 ; } h1 { color : var (--base-color); } 复制代码`

#### 文本 ####

* text-indent： ` <length>` | ` <percentage>`

* 文本缩进，用于块级元素，缩进将沿着行内方向展开

* line-height： ` <number>` | ` <length>` | ` <percentage>` | normal

* em, ex, 百分比是相对与元素的font-size值计算
* 从父元素继承时根据父元素的字号计算，因此最好使用纯数字进行系数换算

* vertical-align：baseline | sub | super | top | text-top | middle | bottom | text-bottom | ` <length>` | ` <percentage>`

* 纵向对齐文本
* 适用于行内元素和单元格

* text-transform：uppercase | lowercase | capitalize | none
* text-decoration：none | underline | overline | line-through | blink

#### 边框图像属性 ####

* border-image-source
* border-image-slice
* border-image-width
* border-image-outset
* border-image-repeat

#### 背景属性 ####

* background-clip: border-box | padding-box | content-box | text ` < div class = "bg-img" > HELLO </ div > <!--css-->.bg-img { width: 1500px; height: 400px; color: transparent; font-size: 300px; background-image: url('bg.jpg'); background-size: contain; -webkit-background-clip: text; background-clip: text; } 复制代码`

![结果](https://user-gold-cdn.xitu.io/2019/6/3/16b1e0ad71f3f4d7?imageView2/0/w/1280/h/960/ignore-error/1)

* background-repeat：repeat-x | repeat-y | [repeat | space | round]

* space：确定沿某一轴能完全重复多少次，然后均匀排列图像
* round：为了放下整个图像，有时会调整图像尺寸，利用round有时能实现一些有趣的效果，比如下面的平铺效果 ![](https://user-gold-cdn.xitu.io/2019/6/4/16b1e1e90fcc10e8?imageView2/0/w/1280/h/960/ignore-error/1)

* background-attachment: scroll | fixed | local

* fixed 固定在视区
* scroll 随文档滚动
* local 随内容滚动

* background-size：length | percentage | cover | contain | auto

* auto 计算相应轴的固有尺寸
* cover 自动覆盖背景，保持固有高宽比 （无需考虑高宽）
* contain 保持固有高宽比，相当于 ` 100% auto` （如果元素高度比宽度大，反之为 ` auto 100%` ）

未完待续