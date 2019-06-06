# 微信小程序中解决iOS中new Date() 时间格式不兼容 #

本周写小程序，遇到的一个bug，在chrome上显示得好好的时间，一到Safari/iPhone 就报错 “invalid date”，时间格式为“2019.06.06 13:12:49”，然后利用new Date() 转换时间戳时，使用微信开发工具、安手机开发版、安手机体验版都没问题，ios中无法展示。

猜想，会不会是Safari不支持yyyy-mm-dd / yyyy.mm.dd 这种格式，于是在 safari 浏览器测试一波，顺便也测试了 “2018-12-10”格式的：

safari 浏览器报错：2018.12.10 11:11:11日期格式

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2bfd12b001d97?imageView2/0/w/1280/h/960/ignore-error/1)

safari 浏览器报错：2018-12-10 11:11:11 日期格式

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2bfd3ef7a5a7a?imageView2/0/w/1280/h/960/ignore-error/1)

于是就replace正则替换

` let dateStr1 = '2018.12.10 11:11:11' ; let dateStr2 = '2018-12-10 11:11:11' ; /* 利用正则表达式替换时间中的”-或者.”为”/”即可 */ dateToTimestamp(dateStr) { if (!dateStr) { return '' } let newDataStr = dateStr.replace(/\.|\-/g, '/' ) let date = new Date(newDataStr); let timestamp = date.getTime(); return timestamp } this.dateToTimestamp(dateStr1) this.dateToTimestamp(dateStr2) 复制代码`

后来为了验证自己的想法，上stackoverflow上查查，看到了几个类似的问题，这里挑一个有代表性的给大家看看：

[Safari JS cannot parse YYYY-MM-DD date format?]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F3085937%2Fsafari-js-cannot-parse-yyyy-mm-dd-date-format%3Frq%3D1 )

大概的意思是说，在执行 ` new Date( string )` 的时候，不同浏览器会采用不同的parse，目前 ` chrome` 两种格式都支持，而 ` Safari` 只支持yyyy/mm/dd。

PS：最近在开始做移动端开发，后面应该会遇到了不少兼容性问题，不断总结，希望以后少踩坑！