# AS插件利器，实现依赖库自动补全，支持变量提取和JetPack #

# 背景 #

最近博主在关注 ` JetPack` 相关内容，发现从 ` support` 转到 ` androidx` 后，很多依赖库的名称变化有点大，每次想添加一个依赖库，都都得扒一扒官网查看路径和版本，确实有点难受。不过在 ` jetbrains` 插件世界里，有一款可以自动补全的插件 [GradleDependenciesHelperPlugin]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsiosio%2FGradleDependenciesHelperPlugin ) ，它只支持从 ` mavenCentral()` 搜索并不支持 ` google()` 的仓库，所以 ` android-dependencies-completion` 应运而生，这是一款尝试对 ` Android开发` **友好的** ` dependencies` 补全插件。

# 功能特色 #

* 支持gradle依赖库名称自动补全，包括 ` Jetpack` 相关的软件包
* 支持版本号提取生成变量和整个路径提取生成变量
* 变量提取功能可以单独使用
* 简洁明了的用户界面

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0239b1d4b3618?imageView2/0/w/1280/h/960/ignore-error/1)

# 如何获取 #

## AS中安装 ##

在Android Studio->Setting ->Plugins中搜索关键字： ` android-dependencies-completion` 或者 ` 12479` ：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a736bba12c6?imageView2/0/w/1280/h/960/ignore-error/1)

## jar包安装 ##

` 插件网站` : [plugins.jetbrains.com/plugin/1247…]( https://link.juejin.im?target=https%3A%2F%2Fplugins.jetbrains.com%2Fplugin%2F12479 )

` 下载地址` ： [github.com/HitenDev/an…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FHitenDev%2Fandroid-dependencies-completion%2Freleases )

您可以在上面两个地址下载jar包，AS安装本地jar插件的方法比较简单，这里就不啰嗦了；

# 如何使用 #

## 掌握快捷键 ##

由于插件依赖 Code Completion->SmartType Completion，所以使用时务必保证 ` SmartType Completion` 是开启的，而且包装 ` SmartType Completion` 的快捷键不和其它快捷键冲突；

![](https://user-gold-cdn.xitu.io/2019/5/29/16b01bc59c8ffd3d?imageView2/0/w/1280/h/960/ignore-error/1)

**默认快捷键：**

* MacOS ` ^(control) + ⇧(shift) + space`
* Windows ` ctrl + alt + space`
* Linux ` ctrl + shift + space`

据悉 ` Windows` 快捷键和系统的快捷键冲突，请使用的小伙伴耐心解决一下冲突，千万不要因此而放弃；

## 基本使用 ##

在项目 ` gradle` 文件中输入字符串时，如果需要补全，请按下 ` 快捷键`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b021d453eef7c0?imageslim)

使用场景不限制 ` build.gradle` ，也不限制 ` dependencies` 下，更不限制是 ` implementation` 还是 ` compile` ；

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06becbba141b9?imageView2/0/w/1280/h/960/ignore-error/1)

通常Anroid开发者喜欢把依赖库统一放置，不见得定义在 ` build.gradle` 中，所有这种场景还是得支持；

## 生成变量 ##

Android开发习惯把gradle依赖库提取成变量，这种场景也是考虑在内，操作方式是在输入字符串的尾部添加 ` #` 符号；

* 添加一个 ` #` ，表示需要提取版本号
* 添加两个 ` ##` ，表示需要把整个字符串都提取出来；

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06ce55cbc65bd?imageslim)

**常见场景：**

由于依赖库字符串是由 ` group:artifact:version` 三部分组成，而用户输入关键字时大部分都不会是完整的三段式，所以 ` #` 支持在缺少的状态下完成；

* 

` 关键字+#`

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06c8e31d10907?imageView2/0/w/1280/h/960/ignore-error/1)

* 

` group:artifact:+#`

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06c98e7545cf5?imageView2/0/w/1280/h/960/ignore-error/1)

### 直接转 ###

* ` group:artifact:version+#` 当字符串是由完整的格式加 ` #` 时，会直接提取变量并完成替换 ![](https://user-gold-cdn.xitu.io/2019/5/29/16b031d2139637d9?imageslim)

## 变量规则 ##

### 命名规则 ###

* 使用 ` #` 生成的版本变量，命名规则是ver_$artifact

` //before implementation 'com.google.code.gson:gson:2.8.5#' //after ext.ver_gson = '2.8.5' //please move this code to a unified place. implementation "com.google.code.gson:gson: $ver_gson " 复制代码`

* ` #` 号后面追加字符串xxx，则命名为ver_xxx

` //before implementation 'com.google.code.gson:gson:2.8.5#hiten' //after ext.ver_hiten = '2.8.5' //please move this code to a unified place. implementation "com.google.code.gson:gson: $ver_hiten " 复制代码`

* 使用 ` ##` 生成的全路径变量，命名规则是dep_$artifact

` //before implementation 'com.google.code.gson:gson:2.8.5##' //after ext.dep_gson = 'com.google.code.gson:gson:2.8.5' //please move this code to a unified place. implementation " $dep_gson " 复制代码`

* ` ##` 号后面追加字符串xxx，则命名为dep_xxx

` //before implementation 'com.google.code.gson:gson:2.8.5##hiten' //after ext.dep_hiten = 'com.google.code.gson:gson:2.8.5' //please move this code to a unified place. implementation " $dep_hiten " 复制代码`

### 插入规则 ###

变量生成的代码，会在当前光标上一行插入，并和当前行左对齐，理论上这行代码放在此处不讲究，所以通常还需要作者把这行代码移动到项目的指定位置；

` ext.ver_gson = '2.8.5' //please move this code to a unified place. 复制代码`

# 其他 #

**单/双引号不设限**

gralde字符串中可以使用 ` $` 引用变量，前提是字符串必须是双引号 ` "` 包裹，我此处做了特殊转换，用户输入时不用在意是单引号还是双引号，只管使用 ` #` 生成就行；

## 再次提示:自动补全不会自动触发，需要用快捷键触发，请读者不要困惑于此。 ##

# 关注该项目 #

如果您对这个功能感兴趣，可以加入一起完善：

项目地址： [github.com/HitenDev/an…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FHitenDev%2Fandroid-dependencies-completion )

您有新的想法，欢迎私聊我或者在github上添加issues；

# 联系我 #

* ` 昵称` : HitenDev
* ` 邮箱` : zzdxit@gmail.com
* ` gayhub` : [github.com/HitenDev]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FHitenDev )
* ` 掘金` : [juejin.im/user/595a16…]( https://juejin.im/user/595a16125188250d944c6997 ) ![](https://user-gold-cdn.xitu.io/2019/5/30/16b071a5ebe62733?imageView2/0/w/1280/h/960/ignore-error/1)