# IntelliJ IDEA 新版发布：支持CPU火焰图，新增酷炫主题 #

> 
> 
> 
> JetBrain 是一家伟大的公司，一直致力于为开发者开发世界上最好用的集成开发环境
> 
> 

就在上周，JetBrain 公司发布了 Java 集成开发环境 IntelliJ IDEA 最新版本 ` 2018.3 Beta` ，本篇文章，我将根据官方博客以及自己的理解来为大家解读一下这次更新有哪些重磅的功能。

## 1. 重构类、文件、符号，Action 搜索 ##

IntelliJ IDEA（以下简称 IDEA） 中的搜索可以分为以下几类

* 类搜索，比如 Java，Groovy，Scala 等类文件
* 文件搜索，类文件之外的所有文件
* 符号搜索，包括接口名，类名，函数名，成员变量等
* Action 搜索，找到你的操作
* 字符串搜索及替换

在 IDEA 的世界里，搜索无处不在，你几乎可以瞬间找到你想要找到的任何一行代码甚至任何一个字。新版中，IDEA 更是将类、文件、符号、Action 搜索与双 ` Shift` 键调出来的 ` Search Everywhere` 无缝地结合在一起。

在老的版本中，类、文件、符号、Action 搜索是独立的快捷键，在新版中，任意一种类型的搜索行为被触发，将弹出来以下窗口

![搜索.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c884a4bc92?imageslim)

从以上演示可以看到，我们调出搜索类的窗口，该窗口将首先会展示基于类名搜索的结果，如果你想复用当前输入的字符基于其他的语义（比如文件或者符号）进行搜索，只需要按 Tab 键，结果瞬间就出来了。

## 2. 重新设计的结构搜索/替换对话框 ##

其实，IDEA 里面除了以上五种类型的搜索，还有一种非常强大的搜索叫做 ` 结构化搜索` ，你可以基于一定的代码结构搜到你所需要的结果。

举个栗子：如果我们想搜索所有的 try catch 语句块，在调出结构化搜索框之后，可以输入以下文本

` try { $TryStatement$; } catch ($ExceptionType$ $Exception$) { $CatchStatement$; } 复制代码`

然后，IDEA 就会把所有的 try catch 语句块搜索出来，而新版更是强化了这个功能，下面我用两张动图演示一下这次更新的两个功能

结构化搜索由于输入的文本比较长，所以一般我们会自己预置一些模板，然后给模板命名，然后结构化搜索的时候呢，我们就可以直接基于这个模板名来搜索，新版更新的第一个功能就是，在文本输入框里，按下智能补全键，可以迅速调出模板，按照最近的搜索历史排序，然后再按下回车，文本就自动给你填充上了，你还可以点击左上角的搜索 icon，也会展示你最近的搜索记录，这些记录是以文本的方式展示的

![结构搜索1.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c88500c827?imageslim)

上面的文本就是系统内置的结构化模板 ` try's` ，点击完 ` Find` 按钮之后，所有的 try catch 都会展示出来，我们还可以进一步过滤，比如，我们想要找出 catch 到的 exception 的名字为 ` flash` ，给对应的模板变量加上一个 Text 类型的 filter 即可迅速定位

![结构搜索2.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c885907f61?imageslim)

更多技巧在关注"闪电侠的博客"公众号之后，回复 idea 即可获取。

## 3. 运行一切 ##

你可以双击 ` ctrl` 键，调出 ` Run Anything` 窗口，你可以输入点什么来运行任意可以运行的东西，比如起 tomcat 容器，单元测试，甚至可以运行终端指令，gradle、maven 构建命令

![运行一切.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c8860dac64?imageslim)

另外，你还可以按住 shift 键，那么所有支持 debug 的运行将秒变 debug 模式

## 4. 重构插件中心 ##

IDEA 中很多强大的功能都是通过插件来实现的，随便举个栗子，装个语言插件，IDEA 摇身一变为 nodejs IDE、php IDE、python IDE、scala IDE、go IDE，我自己就安装了 30+ 非常好用的插件。

而在新版的 IDEA 中，JetBrain 更是对插件中心进行全面改版，如下图

![插件中心.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c8861359ad?imageslim)

调出插件配置之后，页面分为三大部分

* ` Marketplace` ： 插件市场，你可以搜索到你想要的插件
* ` Installed` ： 当前安装的所有的插件，你还可以点击左上角搜索小 icon，按类别查看当前已安装的插件，其中的 ` custom` 选项便是自己下载安装的插件
* ` Updates` ：当前安装过的插件如果有更新，都会在这里显示出来
* 最后一个是配置项，你可以自定义你的插件仓库，你可以给配置插件下载的 http 代理（尤其是国外网络访问差的时候），你还可以从本地硬盘中安装插件

## 5. 不断改进的版本控制系统 ##

我个人对于版本控制，是不太喜欢用图形界面的，但是 IDEA 对于版本控制的设计真是太好用了，只能沦陷了，嘿嘿~

### 5.1 GitHub Pull Requests ###

新版中，加入了对 GitHub Pull Requests 的支持，现在你可以直接在电脑上创建或者查看某个项目的 Pull Request 了

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c88647c2d1?imageView2/0/w/1280/h/960/ignore-error/1)

你还可以基于某个 Pull Requests 直接创建一个分支，或者直接在 Github 上查看当前的 Pull Request，这个功能对于开源工作者来说是一件非常幸福的事。

### 5.2 Git 子模块支持 ###

此外，新本 IDEA 对于 Git 子模块的支持也更加友好了。如果你的 Git 项目中包含 Git 子模块，在 clone 代码的时候，也会一并 clone 到本地，另外，项目中任何文件有变更，提交 commit，IDEA 也会智能匹配到外层模块或者子模块，一并提交 commit，进而同时 push 到多个仓库。

### 5.3 Improved Annotate support ###

我们有时候会不经意地格式化自己或者别人写过的代码，这就导致了每次提交代码的时候，即使只更新了一两处代码，最后 diff 出来也会显得很乱，然而其中大部分乱的地方是因为空格导致的。

在新版 IDEA 中，我们在对比文件的时候，可以选择忽略空格

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c8deb877bf?imageView2/0/w/1280/h/960/ignore-error/1)

注意：这个选项默认是打开的

另外，在合并代码的时候，你也可以选择忽略空格

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c8ea823f5b?imageView2/0/w/1280/h/960/ignore-error/1)

这样在解决冲突的时候，你也不会看到空格相关的改动，省下的很多宝贵的注意力。

IDEA 对于版本控制的支持实在是太强大了，更多版本控制神技在关注"闪电侠的博客"公众号之后，回复 idea 即可获取。

## 6. 全新主题 ##

IDEA 终于在这一版新增了一款默认主题，该主题为一款高对比度主题，应该会有很多人会喜欢吧

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c8eff01031?imageView2/0/w/1280/h/960/ignore-error/1)

预计在不久的将来，IDEA 会在主题这方面下功夫，毕竟笔者觉得 VS Code 的主题还是蛮好看的，IDEA 可以吸收过来。

## 7. 编辑器改进 ##

### 7.1 多行 TODO 注释 ###

在 IDEA 中，只要你在注释中添加了 ` todo` 关键词，在边条栏中的 ` todo` 选项卡中就可以看到当前所有待未完成的功能，如下图

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c9171b7f06?imageView2/0/w/1280/h/960/ignore-error/1)

老版本中，是不支持 多行 ` todo` 注释的显示的，而在新版本中，如果 ` todo` 注释有多行，你只需要在下面几行前面再添加一个空格即可

![todo动图.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c956687de2?imageslim)

### 7.2 缩进状态栏 ###

IDEA 现在可以在状态栏中显示当前文件的缩进是几个空格，你可以点击这个状态栏，控制当前文件的缩进风格。

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c967f3651b?imageView2/0/w/1280/h/960/ignore-error/1)

比如，你的项目缩进风格是4个空格，然后某个新人写了个 tab 风格的源文件提交了，你可以直接点击弹出菜单的 ` Configure Indents For Java...` ，然后做一些修改即可

### 7.3 TAB 快速切换源文件 ###

![tab快速切换.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24c96bf1a246?imageslim)

你现在可以使用 Tab+数字，迅速切换到你想要的文件，这比鼠标点击要快一些

### 7.4 多行字符串搜索 ###

在新版 IDEA 中，不仅仅能够搜索字符串，而且能够搜索整个段落

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c9755a9c41?imageView2/0/w/1280/h/960/ignore-error/1)

## 8. JVM 调试器 ##

## 8.1 attach 到任意 Java 进程 ##

IDEA 的 debug 功能无论是对于调试找错还是阅读源码，都发挥了非常重要的作用，新版 IDEA 对 debug 功能进一步加强，现在不仅仅能 debug 当前的应用，而且能够 attach 到任意的 Java 进程，attach 之后，你就可以看到该进程的线程状态，并且使用强大的 Memory View 功能可以看到当前内存的状态。

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24c97e29b662?imageView2/0/w/1280/h/960/ignore-error/1)

## 8.2 远程调试支持异步栈追踪 ##

IDEA 支持远程 debug 几乎和本地 debug 一样，只需要远程端口开启即可。

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24ca091cdfb3?imageView2/0/w/1280/h/960/ignore-error/1)

IDEA 也支持异步线程的调试，断点打在某一行，你不仅可以看到这行对应线程的调用栈，还能看到启动对应线程的外部线程的调用栈。

![异步调试.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24ca0c5a9601?imageslim)

新版中，对远程调试也加入了异步栈的支持，采用以下两个步骤即可

* 拷贝 ` /lib/rt/debugger-agent.jar` 到远程机器
* 添加启动参数 ` -javaagent:debugger-agent.jar` 到远程机器

如何使用 debug 功能来迅速找错，如何通过 debug 闪电般地阅读源码，在关注"闪电侠的博客"公众号之后，回复 idea 即可获取酷炫神技。

## 9. 运行配置 ##

### 9.1 配置宏 ###

我们在运行应用程序的时候，有的时候需要设定不同的启动参数来查看不同的效果，在以前，这些参数都需要你手动敲进去，并且经常会忘记当前启动参数的测试目的，非常麻烦。

现在，你可以提前将参数通过宏的方式输入，调试的时候，通过调整宏，你不用反复修改启动参数文本，通过宏文本还可以一目了然看到当前的启动参数的测试目的是什么。

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24ca0c313836?imageView2/0/w/1280/h/960/ignore-error/1)

### 9.2 使用文本作为控制台输入 ###

有的时候需要在控制台输入一些文本，然后再运行程序，这个对于调试来说非常不便，新版 IDEA 支持指定一个文本文件作为控制台输入，这样，你就可以预先定义好控制台输入，重复利用，提高效率

![image.png](https://user-gold-cdn.xitu.io/2018/10/30/166c24ca0d1a233b?imageView2/0/w/1280/h/960/ignore-error/1)

## 10. JVM Profiler ##

最后一个重磅功能，应该可以说是本次更新最大的亮点，IDEA 现在可以分析 Java 程序的性能分析了，包括如下几个方面

* 火焰图分析 CPU 性能消耗，你可以分析 Java 进程的所有线程的 CPU 消耗火焰图，也可以只选择一个线程来分析
* 方法调用图，可以找到在某个线程中，消耗 cpu 最多的方法
* 方法列表，可以看到每个方法的调用次数，点进去还可以看到详细的调用栈

下面用一章动图来展示一下，具体的细节读者可自行探索

![jvm profiler.gif](https://user-gold-cdn.xitu.io/2018/10/30/166c24ca10d04cdf?imageslim)

有了这个神器之后，你不需要额外的 profiler 工具，就可以直接在 IDEA 里面完成应用程序的性能分析。预计不久的将来，Eclipse MAT 相关的功能可能也会移植到 IDEA 中，届时，Java 应用程序性能分析，堆分析，gc 分析将统统可以在 IDEA 里面运行，真正的 All In One 时代即将到来！

## 11. More…… ##

除此之外，本次更新还有大量的小功能的更新，在你使用新版 IDEA 的时候就会体验到，这里就不一一赘述了，赶紧下载体验吧，下载地址： [www.jetbrains.com/idea/nextve…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jetbrains.com%2Fidea%2Fnextversion%2F%25E3%2580%2582 )

这篇文章更多的是解析本次更新，其实上个版本的更新也有很多重磅的功能，如果你不了解这些，可以参考一下这篇文章： [IntelliJ IDEA 2018.1正式发布]( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI1OTUzMTQyMA%3D%3D%26amp%3Bmid%3D2247484012%26amp%3Bidx%3D1%26amp%3Bsn%3Dfeda939052bee6d28a6fc68e395aef25%26amp%3Bchksm%3Dea76359fdd01bc8913f5b566a0efd652f30f7cfdc395241e76f15e650530943f18ab73fbd079%26amp%3Bscene%3D21%23wechat_redirect ) ，希望能够帮助你

喜欢本文的朋友们，欢迎长按下图关注订阅号 **闪电侠的博客** ，回复 “idea” 立即获取 IntelliJ IDEA 酷炫神技

![image](https://user-gold-cdn.xitu.io/2018/10/30/166c24ca4299b0b2?imageView2/0/w/1280/h/960/ignore-error/1)