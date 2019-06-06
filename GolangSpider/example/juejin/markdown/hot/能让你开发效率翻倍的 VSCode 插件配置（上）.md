# 能让你开发效率翻倍的 VSCode 插件配置（上） #

工欲善其事必先利其器，软件工程师每天打交道最多的可能就是编辑器了。入行几年来，先后折腾过的编辑器有 EditPlus、UltraEdit、Visual Studio、EClipse、WebStorm、Vim、SublimeText、Atom、VSCode，现在仍高频使用的就是 [VSCode]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com ) 和 [Vim]( https://link.juejin.im?target=http%3A%2F%2Fwww.vim.org ) 了。实际上我在 VSCode 里面安装了 Vim 插件，用 Vim 的按键模式编码，因为自从发现双手不离键盘带来的效率提升之后，就尽可能的不去摸鼠标。

折腾过 Atom 的我，首次试用 VSCode 就有种 Vim 的轻量感，试用之后果断弃坑 Atom。Atom 之前，我还使用过 SublimeText，但它在保存文件时会不时弹出购买授权的弹窗，实在是令人烦恼。

每每上手新的编辑器，我都会根据自己的开发习惯把它调较到理想状态，加上熟悉编辑器各种特性，这个过程通常需要几周的时间。接下来，我就从外观配置、风格检查、编码效率、功能增强等 4 方面来侃侃怎么配置 VSCode 来提高工作幸福感。

## 外观配置 ##

外观是最先考虑的部分，从配置的角度，无非是配色、图标、字体等，俗话说萝卜白菜各有所爱，我目前的配色、图标、字体从下图基本都能看出来，供大家参考：

![](https://user-gold-cdn.xitu.io/2017/11/13/34b7cbfbdba9ab93373e40a5d1540083?imageView2/0/w/1280/h/960/ignore-error/1)

* 配色： [Solarized Dark]( https://link.juejin.im?target=http%3A%2F%2Fethanschoonover.com%2Fsolarized ) ，VSCode 已经内置，使用了至少 5 年以上的主题，Vim 下的配置完全相同；
* 图标： [VSCode Great Icons]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Demmanuelbeziat.vscode-great-icons ) ，给不同类型的文件配置不同的图标，非常直观；
* 字体： [Fira Code]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftonsky%2FFiraCode%2Fwiki%2FVS-Code-Instructions ) ，自从发现并开始使用 [Fira Code]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftonsky%2FFiraCode ) ，我就再也没多看自其它字体一眼，字体如果比较优雅，尤其是对数学运算符的处理，写代码时你真的会感觉在写诗，哈哈，Fira Code 的安装过程稍微复杂点，但是效果绝对是值当的；

配色、图标、字体以及其他外观配置项具体如下（注意，如果不安装上述插件，部分配置项如果直接使用是无效的）：

` { "editor.cursorStyle" : "block" , "editor.fontFamily" : "Fira Code" , "editor.fontLigatures" : true , "editor.fontSize" : 16 , "editor.lineHeight" : 24 , "editor.lineNumbers" : "on" , "editor.minimap.enabled" : false , "editor.renderIndentGuides" : false , "editor.rulers" : [ 120 ], "workbench.colorTheme" : "Solarized Dark" , "workbench.iconTheme" : "vscode-great-icons" } 复制代码`

## 风格检查 ##

之前我写过一篇在 Git 提交环节保障代码风格的文章： [《使用 husky 和 lint-staged 打造超溜的代码检查工作流》]( https://juejin.im/post/592615580ce463006bf19aa0 ) 。如果编辑器在编码时实时给出反馈，对开发者个人而言才是最高效的，在提交时做强制检查只是从团队的视角保证编码风格的规范性和一致性。前端工程师会书写的代码无非是：HTML、CSS、Javascript、Markdown、TypeScript、JSON，对应的 Lint 工具就显而易见：

* [ESLint]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Ddbaeumer.vscode-eslint ) ：插件式架构，有多种主流的编码风格规则集可供选择，典型的有 [Airbnb]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Feslint-config-airbnb ) 、 [Google]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgoogle%2Feslint-config-google ) 等，你甚至可以攒个自己的，按下不表；
* [StyleLint]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dshinnn.stylelint ) ，同样插件式架构的样式检查工具，不过我在配置其检查 [react-native]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact-native ) 中 [styled-components]( https://link.juejin.im?target=https%3A%2F%2Fstyled-components.com ) 组件样式时确实费了不小的功夫，可以单独写篇文章了；
* [TSLint]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Deg2.tslint ) ：TypeScript 目前不是我的主要编程语言，但也早早的准备好了；
* [MarkdownLint]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3DDavidAnson.vscode-markdownlint ) ：Markdown 如果不合法，可能在某些场合导致解析器异常，因为 Markdown 有好几套标准，在不同标准间部分语法支持可能是不兼容的；

除上面列的 Lint 工具之外，非常值得拥有的还有 [EditorConfig]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3DEditorConfig.EditorConfig ) 插件，几乎所有主流 IDE 都有支持，我们可以通过简单的配置文件在不同团队成员、不同 IDE、不同平台下约定好文件的缩进方式、编码格式，避免出现混乱，下面是我常用的配置：

` [*] end_of_line = lf charset = utf- 8 trim_trailing_whitespace = false insert_final_newline = true indent_style = space indent_size = 2 [{*.yml,*.json}] indent_style = space indent_size = 2 复制代码`

有了风格检查，自然就会产生按配置好的风格规则做文件格式化的需求，格式化的工具试用了好多，现在还在用的如下：

* [Prettier]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Desbenp.prettier-vscode ) ，实际上已经是代码格式化的 [工具标准]( https://link.juejin.im?target=https%3A%2F%2Fprettier.io ) ，支持格式化几乎所有的前端代码，并且类似于 EditorConfig 支持用文件来配置格式规则；
* [Vetur]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Doctref.vetur ) ，格式化 .vue 文件，包括里面的 CSS、JS，至于模板即 HTML 部分，官方维护者说没有比较好的工具支持，默认是不格式化的；

## 编码效率 ##

说到编码效率，连续六年几乎每天都编码的我目前最大的感受是：击键的速度越来越跟不上思维的速度，这种情况下，就需要在编码时设置适当的快捷键，组合使用智能建议、代码片段、自动补全来达到速度的最大化。

VSCode 内置的智能建议已经非常强大，不过我对默认的配置做了如下修改，以达到类似于在 Vim 中那样在任何地方都启用智能提示（尤其是注释和字符串里面）：

` { "editor.quickSuggestions" : { "other" : true , "comments" : true , "strings" : true }, } 复制代码`

接下来，重点说说代码片段和自动补全两个效率提升利器。

### 代码片段 ###

英文叫做 [Snippets]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdocs%2Feditor%2Fuserdefinedsnippets ) ，市面上主流的编辑器也都支持，其基本思想就是把常见的代码模式抽出来，通过 2~3 个键就能展开 N 行代码，代码片段的积累一方面是根据个人习惯，另一方面是学习社区里面积累出来的好的编码模式，如果觉得不适合你，可以改（找个现有的插件依葫芦画瓢），我常用的代码片段插件如下：

* [HTML Snippets]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dabusaidm.html-snippets ) ，各种 HTML 标签片段，如果你 Emmet 玩的熟，完全可以忽略这个；
* [Javascript (ES6) Code Snippets]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dxabikos.JavaScriptSnippets ) ，常用的类声明、ES 模块声明、CMD 模块导入等，支持的缩写不下 20 种；
* [Javascript Patterns Snippets]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dnikhilkumar80.js-patterns-snippets ) ，常见的编码模式，比如 IIFE；

### 自动补全 ###

自动补全本质上和代码片段类似，不过是在特殊场合下以你的键入做为启发式信息提供最有可能要输入的建议，我常用的自动补全工具有：

* [Auto Close Tag]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dformulahendry.auto-close-tag ) ，适用于 JSX、Vue、HTML，在打开标签并且键入 ` </` 的时候，能自动补全要闭合的标签；
* [Auto Rename Tag]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dformulahendry.auto-rename-tag ) ，适用于 JSX、Vue、HTML，在修改标签名时，能在你修改开始（结束）标签的时候修改对应的结束（开始）标签，帮你减少 50% 的击键；
* [Path Intellisense]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dchristian-kohler.path-intellisense ) ，文件路径补全，在你用任何方式引入文件系统中的路径时提供智能提示和自动完成；
* [NPM Intellisense]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dchristian-kohler.npm-intellisense ) ，NPM 依赖补全，在你引入任何 node_modules 里面的依赖包时提供智能提示和自动完成；
* [IntelliSense for CSS class names]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3DZignd.html-css-class-completion ) ，CSS 类名补全，会自动扫描整个项目里面的 CSS 类名并在你输入类名时做智能提示；
* [Emmet]( https://link.juejin.im?target=https%3A%2F%2Femmet.io ) ，以前叫做 Zen Coding，我发现后，也是爱不释手，可以把类 CSS 选择符的字符串展开成 HTML 标签，VSCode 已经内置，官方介绍文档 [参见]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdocs%2Feditor%2Femmet ) ，你需要做的就是熟悉他的语法，并勤加练习；

当然，如果你还用 VSCode 编写其他语言的代码，比如 PHP，就去市场上搜索 “PHP Intellisense” 好了。

## 功能增强 ##

在效率提升方面除了上面的代码片段、自动补全之外，我还安装了下面几个插件，方便快速的浏览和理解代码，并且在不同项目之间切换。

* [Color Highlight]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dnaumovs.color-highlight ) ，识别代码中的颜色，包括各种颜色格式；
* [Bracket Pair Colorizer]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3DCoenraadS.bracket-pair-colorizer ) ，识别代码中的各种括号，并且标记上不同的颜色，方便你扫视到匹配的括号，在括号使用非常多的情况下能环节眼部压力，编辑器快捷键固然好用，但是在临近嵌套多的情况下却有些力不从心；
* [Project Manager]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Dalefragnani.project-manager ) ，项目管理，让我们方便的在命令面板中切换项目文件夹，当然，你也可以直接打开包含多个项目的父级文件夹，但这样可能会让 VSCode 变慢；

## 结语 ##

说了这么多，相信读到这里的你也期望用工具来提高自己的效率。

提高效率有没有法门？是有的，简单的事情重复化，重复的事情标准化，标准的事情自动化，发现一个痛点，用插件解决一个痛点，你的效率自然就上来了。

你都用了哪些插件呢？欢迎留言交流！

## 题外话 ##

就用上面列出来的 VSCode 配置我录制了 3 门前端短视频教程，你能在这些教程里看到我 VSCode 的实操效果，如果你有兴趣，欢迎点击下面链接：

* [与时俱进：React 16 新特性尝鲜教程]( https://juejin.im/post/59f271a35188255a6a0d47cb ) ，6 小节，18 分钟，用实例演示如何运用新特性提高代码稳定性，优化设计；
* [组件化必杀技：styled-components 简明教程]( https://juejin.im/post/59f7c12b5188255a6a0d5497 ) ，8 小节，23 分钟，让你快速入门现象级的 CSS-IN-JS 解决方案；
* [玩转异步JS：async/await 简明教程]( https://juejin.im/post/5a04f36d6fb9a045293636e0 ) ，8 小节，20 分钟，通过实战让你把 async/await 用的飞起；

另外，以后每周会放出新的短视频教程，如果你想接到推送，欢迎关注《前端周刊》微信公众号：fullstackacademy，关注即得高清视频云盘地址。