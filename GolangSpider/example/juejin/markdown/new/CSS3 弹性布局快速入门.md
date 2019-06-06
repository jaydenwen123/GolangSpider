# CSS3 弹性布局快速入门 #

# 前言 #

弹性布局是新一代的布局方式，传统布局中使用浮动布局会给我们带来不少弊端，如CSS代码高度依赖于HTML代码结构等等，下面我将用几个例子让大家快速学会弹性布局。

PS：弹性布局适用于较简单的场景，过于复杂的场景可以尝试着使用CSS3的Grid布局，弹性布局在PC端中还存在兼容性问题，移动端中无兼容性问题，可以放心使用。

## 1.容器属性 ##

css3为新增的弹性布局提供了多个属性，分别为弹性盒模型的容器属性，以及弹性盒子中子元素的子元素属性。

### 1.1display ###

css3中为display新增了两个属性值，分别为flex、inline-flex

` display:flex; /*将容器声明为一个弹性盒模型且容器表现为块级元素*/ display:inline-flex; /*将容器声明为一个弹性盒模型且容器表现为行内元素*/ 复制代码`

容器display:block;

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c2e2a86b545?imageView2/0/w/1280/h/960/ignore-error/1) 容器display:flex;

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c305392782e?imageView2/0/w/1280/h/960/ignore-error/1) 此时弹性盒模型内的子元素变得类似浮动后的布局，这里要引入弹性盒模型中两条重要的轴线，分别为主轴和垂直轴，如下图所示，弹性盒模型内的子元素默认按照主轴的方向排列。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c8e840652e4?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.2flex-direction ###

flex-direction可以设置主轴的方向，默认值为row。

flex-direction：row | row-reverse | column | column-reverse

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26df3fad7d018?imageView2/0/w/1280/h/960/ignore-error/1) 理解两条轴线至关重要，搞定轴线之后后面就是简单的使用属性了。

### 1.3flex-wrap ###

`.box { width:500px; height:500px; margin:100px auto 0 auto; background: #eee; display: flex; flex-direction: row; } .box-item { width:200px; height:200px; line-height:200px; text-align: center; color: #fff; font-size:20px; } 复制代码`

从上面可以看出容器的宽高都是500px,子元素的宽高都是200px，那如果我们一行放3个元素，元素会像浮动布局那样换行吗？

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e573988d66e?imageView2/0/w/1280/h/960/ignore-error/1) 并没有，同时我们发现了，现在一个子元素的宽度只有166.66px，三个子元素没有换行同时自动等比例缩放至放好可以在容器中放下。 flex-wrap就是控制弹性盒模型的子元素换行方式的，默认值为nowrap。

flex-wrap：nowrap | wrap | wrap-reverse

* flex-wrap：nowrap; /*不换行，等比例缩小*/
* flex-wrap：wrap; /*自动换行*/
* flex-wrap：wrap-reverse; /*自动反方向换行，往下换行变成往上换行*/

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f268f70c62e?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.4justify-content ###

justify-content控制主轴的对齐方式，默认向主轴开始起点位置对齐，值为flex-start。

justify-content：flex-start | flex-end | center | space-between | space-around

* justify-content：flex-start; /*向主轴开始位置对齐*/
* justify-content：flex-end; /*向主轴结束位置对齐*/
* justify-content：center; /*主轴居中对齐*/
* justify-content：space-between; /*等间距对齐，两端不留空*/
* justify-content：space-around; /*等间距对齐，两端留空，每个元素左间距与右间距大小相等,具体见下图*/

![](https://user-gold-cdn.xitu.io/2019/6/5/16b274dcc7087665?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b274dfa130d911?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.5align-items ###

align-items控制垂直轴的对齐方式，默认向主轴开始起点位置对齐，值为flex-start。

align-items：flex-start | flex-end | center | baseline | stretch

* align-items：flex-start; /*向垂直轴开始位置对齐*/
* align-items：flex-end; /*向垂直轴结束位置对齐*/
* align-items：center; /*垂直轴居中对齐*/
* align-items：baseline; /*文本基线对齐，用的不多*/
* align-items：stretch; /*垂直轴方向上的height/width若值为auto，则自动填满，但依然受到min/max-width/height的控制。不设置弹性盒模型时，height默认值为内容区大小，若设置为弹性盒模型且align-items设置为stretch，则高度占满整个父容器*/ ![](https://user-gold-cdn.xitu.io/2019/6/5/16b276232f9043fb?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2762560a0d3a6?imageView2/0/w/1280/h/960/ignore-error/1)

### 1.6align-content ###

刚刚说完了垂直轴只有一个元素的情况，若垂直轴有两个元素时，align-items还能起作用吗？ 为了使垂直轴存在两个元素，我们首先设置自动换行

` flex-wrap:wrap; 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2767045d1fe2d?imageView2/0/w/1280/h/960/ignore-error/1) 从图中可以看到，这不是我们想要的效果，我们想要的效果是垂直轴方向上的两个元素紧贴着的。

这时我们要用align-content。 align-content：flex-start | flex-end | center | space-between | space-around

* align-content：flex-start; /*向主轴开始位置对齐*/
* align-content：flex-end; /*向主轴结束位置对齐*/
* align-content：center; /*主轴居中对齐*/
* align-content：space-between; /*等间距对齐，两端不留空*/
* align-content：space-around; /*等间距对齐，两端留空，每个元素上间距与下间距大小相等,具体见下图*/

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2772fc7670a41?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27733d4316dde?imageView2/0/w/1280/h/960/ignore-error/1) 垂直轴只有一行元素时使用align-items属性，有多行元素时使用align-content属性。

## 2.子元素属性 ##

### 2.1order ###

order属性可用于设置子元素的位置，order的值越小排在越前面，默认值为0，可以设置负值。

` //设置第三个子元素的order为-1 .box-item3 { background: green; order:-1; } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b277ac3978e5a3?imageView2/0/w/1280/h/960/ignore-error/1) 通过为每一个子元素设置order值可以使得布局不依赖于html的结构。

## 3.小结 ##

弹性布局简单易上手，使用熟练之后会发现只需要很少的代码就可以实现以前复杂的布局，弹性布局最重要的就是两条轴线，这个一定要熟练掌握，很多属性的方向都与两条轴线有关。

## 4.交流 ##

如果这篇文章帮到你了，觉得不错的话来点个Star吧。 [github.com/lizijie123]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flizijie123%2F2019Study )