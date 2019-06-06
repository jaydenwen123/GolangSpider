# CSS 多行省略失效 (-webkit-box-orient 失效) #

## 背景 ##

scss文件中，设置多行省略（如下），代码未生效

` overflow: hidden; text-overflow: ellipsis; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; 复制代码`

## 现象 ##

* 

代码未生效：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bc69984656e1?imageView2/0/w/1280/h/960/ignore-error/1)

* 

查看style,-webkit-box-orient 编译之后消失 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bc63347d527c?imageView2/0/w/1280/h/960/ignore-error/1)

## 原因 ##

postcss-loader 默认编译的时候，会过滤 -webkit-box-orient: vertical;

## 解决方案 ##

方案一： 对scss文件没有特殊要求，可以换成.less等其他格式文件

方案二: autoprefixer（推荐使用）

` //单个属性 text-overflow: ellipsis; display: -webkit-box; -webkit-line-clamp: 2; /* autoprefixer: off */ -webkit-box-orient: vertical; /* autoprefixer: on */ 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bc7e98110718?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bc8092f14cd7?imageView2/0/w/1280/h/960/ignore-error/1)

原因：

1.autoprefixer会帮你删除老式过时的代码
2. autoprefixer也会帮你增加新前缀

**方案二优化** ：在webpack中配置postcss的autoprefixer： [www.webpackjs.com/loaders/pos…]( https://link.juejin.im?target=https%3A%2F%2Fwww.webpackjs.com%2Floaders%2Fpostcss-loader%2F )

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bc91b53c4bf7?imageView2/0/w/1280/h/960/ignore-error/1)

参考：

[github.com/postcss/aut…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpostcss%2Fautoprefixer%2Fissues%2F776 )

[PostCSS 配置指北]( https://juejin.im/entry/57e67ac28ac247005bc9562a )

[webpack4之使用postcss-loader和autoprefixer浏览器兼容]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Faa3e52be303e )