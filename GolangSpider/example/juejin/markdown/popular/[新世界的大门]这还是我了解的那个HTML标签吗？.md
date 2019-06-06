# [新世界的大门]这还是我了解的那个HTML标签吗？ #

## 写在前面 ##

由于 IE 这种老古董的存在和浏览器厂家标准不一的情况，前端开发者们大多更愿意写 JavaScript，而对 HTML 和 CSS 嗤之以鼻。但随着 win10 的普及和 Edge 转投 chromium 的怀抱，大家是不是可以对前端开发的未来更多了一些期待，其实现在 HTML 和 CSS 标准推出的也是很快了，为何不分出一点时间给它们呢。下面整理了一些大家可能不熟悉或不经常用的标签，不知道会不会刷新你对 HTML 的认识？

## 想要给汉字加上拼音 ##

你有没有遇到过在页面里想要给一些生僻字加上注释拼音的场景，现在好了，有了 ` ruby` ，你可以很方便地的给汉字加上拼音了。

![ruby-normal](https://user-gold-cdn.xitu.io/2019/5/21/16ad8401f16e4dc6?imageView2/0/w/1280/h/960/ignore-error/1)

` < ruby > 快狗打车 < rt > kuaigoudache </ rt > </ ruby > 复制代码`

这时候你可能想给拼音加上声调，细心的同学还可能注意到了拼音和汉字有时候并不会完全地对齐，怎么做？搜狗输入法切换软键盘可以添加声调，或者点开下面的 /在线中文拼音转换/ 输入汉字可得拼音；对齐的终极大法就是单独给每个字加拼音， ` rp` 是 浏览器不支持 ` ruby` 标签时的降级方案。

` < ruby > 快 < rp > ( </ rp > < rt > kuài </ rt > < rp > ) </ rp > 狗 < rp > ( </ rp > < rt > gǒu </ rt > < rp > ) </ rp > 打 < rp > ( </ rp > < rt > dǎ </ rt > < rp > ) </ rp > 车 < rp > ( </ rp > < rt > chē </ rt > < rp > ) </ rp > </ ruby > 复制代码`

![ruby](https://user-gold-cdn.xitu.io/2019/5/21/16ad841f0b093efe?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> [codepen](
> https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fsgiyy%2Fpen%2FbyrPZz
> ) > [在线中文拼音转换](
> https://link.juejin.im?target=http%3A%2F%2Fxh.5156edu.com%2Fconversion.html
> ) > [ruby| MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fruby
> )
> 
> 

## ` details` 挂件 ##

` details` 创建一个挂件，默认处于关闭状态，只有在打开的状态才会显示其中隐藏的内容。

![details](https://user-gold-cdn.xitu.io/2019/5/21/16ad84dbf0069e5f?imageView2/0/w/1280/h/960/ignore-error/1)

` < details > < summary > 总结 </ summary > <!-- 非必填，默认显示“详细信息” --> 我是内容，我是内容。我是内容，我是内容。我是内容，我是内容。 </ details > 复制代码`
> 
> 
> 
> 
> [codepen](
> https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fsgiyy%2Fpen%2FdEzxBa
> ) > [details | MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fdetails
> )
> 
> 

## ` meter` 计量器 ##

` meter` 用来展示确定范围内的标量值。 Min (fix: 大小写统一) 最小值，如果设置必须要比最大值小，不设置时默认为 0；max 最大值，如果设置必须要比最小值大，不设置时默认为 1。 low 和 high 分别定义了低值区间上限值和高值区间的下限值，浏览器会根据当前值所在的区间渲染不同的效果。 注意区别于 ` progress` ，如果只是用来表示百分比和进度，建议用 ` progress` 。

![meter](https://user-gold-cdn.xitu.io/2019/5/21/16ad84c1f21866e1?imageView2/0/w/1280/h/960/ignore-error/1)

` <!-- 计量器 --> < p > 油体积： < meter min = "0" max = "100" value = "60" high = "85" > 60L </ meter > </ p > < p > 油体积： < meter min = "0" max = "100" value = "90" high = "85" > 85L </ meter > </ p > 复制代码`
> 
> 
> 
> 
> [codepen](
> https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fsgiyy%2Fpen%2FzQEOLe
> ) > [meter | MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fmeter
> )
> 
> 

## ` progress` 进度条 ##

` progress` 用来显示任务的完成进度，外观如何显示可以自己定义，通常显示成进度条的形式。

![progress](https://user-gold-cdn.xitu.io/2019/5/21/16ad84e7095bd715?imageView2/0/w/1280/h/960/ignore-error/1)

` < label for = "file" > 进度: </ label > < progress id = "file" max = "100" value = "60" > 60% </ progress > 复制代码`
> 
> 
> 
> 
> [progress | MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fprogress
> )
> 
> 

## ` picture` 图片元素 ##

` picture` 元素包含一个或多个 ` source` 元素和一个 ` img` 元素，包含的 ` img` 元素更像是一个备选方案，浏览器加载的时候会先检查每一个 ` source` 元素的 ` srcset` 、 ` media` 、 ` type` 等属性，并选用最合适的一个，如果没有合适的，就会显示 ` img` 元素。

` < picture > <!-- 屏幕尺寸大于800px才显示，否则显示 下面的 img --> < source srcset = "https://cn.vuejs.org/images/logo.png" media = "(min-width: 800px)" /> < img src = "https://static.daojia.com/assets/project/tosimple-pic/iview-logo_1558354678823.png" /> </ picture > 复制代码`
> 
> 
> 
> 
> [picture | MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fpicture
> )
> 
> 

## output 显示表单元素的计算结果 ##

可以组成一个简易计算器。

![output](https://user-gold-cdn.xitu.io/2019/5/21/16ad84c1f3fdc535?imageView2/0/w/1280/h/960/ignore-error/1)

` < form oninput = "result.value=parseInt(a.value)+parseInt(b.value)" > < input type = "number" name = "b" value = "40" /> + < input type = "number" name = "a" value = "10" /> = < output name = "result" > 50 </ output > </ form > 复制代码`
> 
> 
> 
> 
> [codepen](
> https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fsgiyy%2Fpen%2FWBZKJa
> ) > [output | MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Foutput
> )
> 
> 

## ` track` 媒体元素 ##

HTML 5 新增标签，当作 ` audio` 和 ` video` 的子元素，一般是用来处理字幕等功能。 ` track` 添加的数据的类型通过 ` kind` 属性指定，属性值可以是： subtitles | captions | descriptions | chapters | metadata 。

` < video controls poster = "/images/sample.gif" > < source src = "sample.mp4" type = "video/mp4" /> < track kind = "subtitles" src = "sampleSubtitles_en.vtt" srclang = "en" /> Sorry, your browser doesn't support embedded videos. </ video > 复制代码`
> 
> 
> 
> 
> [track | MDN](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Ftrack
> )
> 
> 

## 结尾 ##

诚然，有一些标签会有一定的兼容性，比如 ` picture` 。但如果你的目标用户使用 IE 的份额很少，为了更好的用户体验和网站性能，可以评估这样是不是也值得呢。

> 
> 
> 
> [Can I use… Support tables for HTML5, CSS3, etc](
> https://link.juejin.im?target=https%3A%2F%2Fcaniuse.com%2F )
> 
> 

## 关于我们 ##

快狗打车前端团队专注前端技术分享，定期推送高质量文章，欢迎关注点赞。