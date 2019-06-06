# 使用CSS定义页面元素的外观样式 #

` CSS` 在网页里的作用，我分成两块：一是定义页面元素的外观样式，二是定义页面元素的排版布局。本篇就通过例子来说明：如何使用 ` CSS` 来定义页面元素的样式？

回到我们的按钮例子：页面上有一个按钮，目前是没有任何自定义样式的，只有一个默认的样式。

` < html > < head > < title > </ title > < script > function buttonHandler () { alert( 'Hello' ) } </ script > </ head > < body > < button onclick = "buttonHandler()" > 按钮 </ button > </ body > </ html > 复制代码`

### 1、设置按钮的宽度和高度 ###

` < html > < head > < title > </ title > < style > /* 元素通过选择器与样式关联 */.button { /* 设置按钮高度 */ height : 32px ; /* 设置按钮内边距 */ padding : 0 15px ; } </ style > < script > function buttonHandler () { alert( 'Hello' ) } </ script > </ head > < body > < button class = "button" onclick = "buttonHandler()" > 按钮 </ button > </ body > </ html > 复制代码`

这里没有直接设置按钮的宽度，而是设置了按钮内边距，这样按钮的宽度会根据文字的长度而撑开。

### 2、让这个按钮变靓 ###

` < html > < head > < title > </ title > < style > /* 元素通过选择器与样式关联 */.button { /* 设置按钮高度 */ height : 32px ; /* 设置按钮内边距 */ padding : 0 15px ; /* 设置背景颜色 */ background-color : #1890ff ; /* 设置边框颜色 */ border-color : #1890ff ; /* 设置文字的颜色 */ color : #fff ; /* 设置字号 */ font-size : 14px ; /* 设置圆角 */ border-radius : 4px ; /* 给文字加上阴影 */ text-shadow : 0 - 1px 0 rgba (0,0,0,0.12); /* 给边框加上阴影 */ box-shadow : 0 2px 0 rgba (0,0,0,0.045); } </ style > < script > function buttonHandler () { alert( 'Hello' ) } </ script > </ head > < body > < button class = "button" onclick = "buttonHandler()" > 按钮 </ button > </ body > </ html > 复制代码`

现在这个按钮就长得和 ` ant-design` 一样了，我填加了注释来说明不同属性的作用。通过设置不同属性的值， ` CSS` 支持你把页面元素定义成自己想要的任何外观！其它的属性可以查阅 ` CSS` 文档。

### 让按钮的样式响应交互 ###

` < html > < head > < title > </ title > < style > /* 元素通过选择器与样式关联 */.button { /* 设置按钮高度 */ height : 32px ; /* 设置按钮内边距 */ padding : 0 15px ; /* 设置背景颜色 */ background-color : #1890ff ; /* 设置边框颜色 */ border-color : #1890ff ; /* 设置文字的颜色 */ color : #fff ; /* 设置字号 */ font-size : 14px ; /* 设置圆角 */ border-radius : 4px ; /* 给文字加上阴影 */ text-shadow : 0 - 1px 0 rgba (0,0,0,0.12); /* 给边框加上阴影 */ box-shadow : 0 2px 0 rgba (0,0,0,0.045); } /* CSS伪类 */.button :hover { /* 设置透明度 */ opacity : 0.8 ; } </ style > < script > function buttonHandler () { alert( 'Hello' ) } </ script > </ head > < body > < button class = "button" onclick = "buttonHandler()" > 按钮 </ button > </ body > </ html > 复制代码`

当鼠标经过按钮的时候，按钮会变成半透明。

### 给按钮添加（过渡）动效 ###

` < html > < head > < title > </ title > < style > /* 元素通过选择器与样式关联 */.button { /* 设置按钮高度 */ height : 32px ; /* 设置按钮内边距 */ padding : 0 15px ; /* 设置背景颜色 */ background-color : #1890ff ; /* 设置边框颜色 */ border-color : #1890ff ; /* 设置文字的颜色 */ color : #fff ; /* 设置字号 */ font-size : 14px ; /* 设置圆角 */ border-radius : 4px ; /* 给文字加上阴影 */ text-shadow : 0 - 1px 0 rgba (0,0,0,0.12); /* 给边框加上阴影 */ box-shadow : 0 2px 0 rgba (0,0,0,0.045); /* 设置透明度变化时的过渡效果 */ transition : opacity 0.5s ; } /* CSS伪类 */.button :hover { /* 设置透明度 */ opacity : 0.2 ; } </ style > < script > function buttonHandler () { alert( 'Hello' ) } </ script > </ head > < body > < button class = "button" onclick = "buttonHandler()" > 按钮 </ button > </ body > </ html > 复制代码`

为了让效果看起来明显，我将鼠标经过时的透明度设置成了0.2。设置 ` transition` 属性后，按钮的透明度变化就会有一个过渡效果，而不是直接从1变成0.2。

## 小结 ##

还想增加一个动画的例子，但是为了避免代码越加越长，必须来个总结了。

本篇涉及到的问题有点多：

* 例子里讲CSS样式写在页面的style标签里，还有其它的方式？
* 什么是选择器？除了通过 `.` 开头的类选择器，还有哪些？
* 设置高度的时候，用到的单位px是什么意思，还支持哪些单位？
* 按钮的背景可以是一个颜色，也可以是一张背景图？
* 除了可以响应鼠标经过，还支持哪些伪类？
* 过渡效果和动画？
* ...

每一行代码都可以挖掘问题！下一篇例子关于如何使用CSS进行页面布局。