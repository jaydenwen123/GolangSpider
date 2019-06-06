# 关于webpack4的14个知识点,童叟无欺 #

![](https://user-gold-cdn.xitu.io/2019/5/8/16a94b981baecfa7?imageView2/0/w/1280/h/960/ignore-error/1) 没有什么比时间更具有说服力了，因为时间无需通知我们就可以改变一切。 ![](https://user-gold-cdn.xitu.io/2019/5/8/16a94b981b0702d0?imageView2/0/w/1280/h/960/ignore-error/1)

最近工作中用到了nuxt,才发现,如果webpack学的6,nuxt基本不需要学习,没什么学习成本的,因此,这篇重新记录下webpack4的一些基础知识点, **下一篇将会配置一个优化到极致的react脚手架** ,也希望大家能够持续关注,配置webpack就是优化优化再优化,哈哈~

> 
> 
> 
> 酒壮怂人胆,我学这个的办法基本就分3步：
> 
> 

* **首先,将这些必要的配置,以及某些loader,某些插件,像语文课文一样默读,并背诵(这一步最重要)**
* **动手去实践,去试错**
* **理解其原理**

好了,正式开始

# 前言 #

Webpack可以看做是模块打包机：它做的事情是，分析你的项目结构，找到JavaScript模块以及其它的一些浏览器不能直接运行的拓展语言（Scss，TypeScript等），并将其打包为合适的格式以供浏览器使用。

#### WebPack和Grunt以及Gulp相比有什么特性 ####

其实Webpack和另外两个并没有太多的可比性，Gulp/Grunt是一种能够优化前端的开发流程的工具，而WebPack是一种模块化的解决方案，不过Webpack的优点使得Webpack在很多场景下可以替代Gulp/Grunt类的工具。

* Entry：入口，Webpack 执行构建的第一步将从 Entry 开始，可抽象成输入。
* Module：模块，在 Webpack 里一切皆模块，一个模块对应着一个文件。Webpack 会从配置的 Entry 开始递归找出所有依赖的模块。
* Chunk：代码块，一个 Chunk 由多个模块组合而成，用于代码合并与分割。
* Loader：模块转换器，用于把模块原内容按照需求转换成新内容。
* Plugin：扩展插件，在 Webpack 构建流程中的特定时机注入扩展逻辑来改变构建结果或做你想要的事情。
* Output：输出结果，在 Webpack 经过一系列处理并得出最终想要的代码后输出结果。

# 1. 从0开始配置结构 #

* 初始化项目结构

![](https://user-gold-cdn.xitu.io/2019/5/26/16af28bf49f5c2d9?imageView2/0/w/1280/h/960/ignore-error/1)

# 2. 配置webpack.config.js #

* 在项目根目录新建webpack.config.js

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2cd0e47b7b1b?imageView2/0/w/1280/h/960/ignore-error/1)

# 3. 配置开发服务器 #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2cdb4cb7be3b?imageView2/0/w/1280/h/960/ignore-error/1)

# 4. 打包js #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2ce28d674199?imageView2/0/w/1280/h/960/ignore-error/1)

# 5. 支持ES6,react,vue #

![](https://user-gold-cdn.xitu.io/2019/6/4/16b1fdf739fe273a?imageView2/0/w/1280/h/960/ignore-error/1)

# 6. 处理css,sass,以及css3属性前缀 #

## 处理css ##

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d01b1f62296?imageView2/0/w/1280/h/960/ignore-error/1)

## 动态卸载和加载 ` CSS` ##

> 
> 
> 
> style-loader为 css 对象提供了use()和unuse()两种方法可以用来加载和卸载css
> 
> 

比如实现一个点击切换颜色的需求，修改index.js

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d116212a246?imageView2/0/w/1280/h/960/ignore-error/1)

## 处理sass ##

![](https://user-gold-cdn.xitu.io/2019/5/27/16af6f99cb26d258?imageView2/0/w/1280/h/960/ignore-error/1)

## 提取css文件为单独文件 ##

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d382319f886?imageView2/0/w/1280/h/960/ignore-error/1)

# 7.产出html #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d43624a2681?imageView2/0/w/1280/h/960/ignore-error/1)

# 8. 处理引用的第三方库,暴露全局变量 #

webpack.ProvidePlugin参数是键值对形式，键就是我们项目中使用的变量名，值就是键所指向的库

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d4e9e2b3585?imageView2/0/w/1280/h/960/ignore-error/1)

# 9. code splitting、懒加载(按需加载) #

说白了就是在需要的时候在进行加载，比如一个场景，点击按钮才加载某个js.

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d5c4a1100e1?imageView2/0/w/1280/h/960/ignore-error/1)

# 10. JS Tree Shaking #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d6e651e68e6?imageView2/0/w/1280/h/960/ignore-error/1)

# 11. 图片处理 #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d781ecf6876?imageView2/0/w/1280/h/960/ignore-error/1)

# 12. Clean Plugin and Watch Mode #

清空目录，文件有改动就重新打包

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d80bbde0581?imageView2/0/w/1280/h/960/ignore-error/1)

# 13. 区分环境变量 #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d889c307745?imageView2/0/w/1280/h/960/ignore-error/1)

# 14. 开发模式与webpack-dev-server,proxy #

![](https://user-gold-cdn.xitu.io/2019/5/26/16af2d8fd11ec8c2?imageView2/0/w/1280/h/960/ignore-error/1)

到这里基本就结束了,觉得有帮助,不妨点个 **赞** ,不足之处，还望斧正~