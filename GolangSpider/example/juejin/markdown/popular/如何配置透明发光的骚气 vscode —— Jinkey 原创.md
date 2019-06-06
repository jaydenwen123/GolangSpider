# 如何配置透明发光的骚气 vscode —— Jinkey 原创 #

> 
> 
> 
> 原文链接 [jinkey.ai/post/tech/r…](
> https://link.juejin.im?target=https%3A%2F%2Fjinkey.ai%2Fpost%2Ftech%2Fru-he-pei-zhi-tou-ming-fa-guang-de-sao-qi-vscode
> ) 转载请注明出处
> 
> 

# 1 安装自定义 JS 和 CSS 插件 #

![插件截图](https://user-gold-cdn.xitu.io/2019/5/15/16abb00655612421?imageView2/0/w/1280/h/960/ignore-error/1)

# 2 安装发光主题 #

![插件截图](https://user-gold-cdn.xitu.io/2019/5/15/16abb00655ac3f71?imageView2/0/w/1280/h/960/ignore-error/1)

# 3 添加样式配置文件 #

在 VSCode 安装目录（自己随便选择一个文件夹也可以），放入以下文件。 为了方便下载，文件整理到了 [Github-Jinkeycode]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJinkeycode ) / [vscode-transparent-glow]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJinkeycode%2Fvscode-transparent-glow ) ，欢迎 star。

` enable-electron-vibrancy.js` 开启 electron 透明支持

![](https://user-gold-cdn.xitu.io/2019/5/15/16abb00655cc4077?imageView2/0/w/1280/h/960/ignore-error/1)

` vscode-vibrancy-style.css`

![](https://user-gold-cdn.xitu.io/2019/5/15/16abb00655f55d7a?imageView2/0/w/1280/h/960/ignore-error/1)

` synthwave84.css` 文字发光样式，样式请在 Github 获取。如果要不发光的，可以使用 ` synthwave84-noglow.css` 。可以 watch [github.com/robb0wen/sy…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Frobb0wen%2Fsynthwave-vscode ) 保持更新通知。

` toolbar.css` 引入以上大神的样式配置之后，左边导航栏有部分标题还是未透明状态，我自己修改了配置，引入即可。

![](https://user-gold-cdn.xitu.io/2019/5/15/16abb0065c9c8890?imageView2/0/w/1280/h/960/ignore-error/1)

` terminal.css` 使 vscode 内置的终端透明

![](https://user-gold-cdn.xitu.io/2019/5/16/16abc5b1586903c8?imageView2/0/w/1280/h/960/ignore-error/1)

终端光标颜色修改，由 [@manonloki]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJinkeycode%2Fvscode-transparent-glow%2Fblob%2Fmaster%2Fwww.manonloki.com ) 提供

`.panel.integrated-terminal .xterm-cursor, .xterm-cursor-block { background: rgb(210, 0, 252) !important; } 复制代码`

# 4 修改 VSCode 配置文件 #

![](https://user-gold-cdn.xitu.io/2019/5/15/16abb0065ca3cc4b?imageView2/0/w/1280/h/960/ignore-error/1) 点击上图 ` 在 setting.json 中编辑` ，打开后加入配置( **不需要大括号** ，直接把 key-value 加入原有 json 即可)： ![setting.json](https://user-gold-cdn.xitu.io/2019/5/16/16abc590e4483141?imageView2/0/w/1280/h/960/ignore-error/1)

# 5 重加载 #

按下 Ctrl + Shift + P，运行 "Reload Custom CSS and JS", 重启 vscode 即可。如果提示 ` VSCode 已经损坏` ，选择右上角齿轮“不再提示”即可。

部分电脑提示 reload 失败的，请以管理员模式运动 vscode

` sudo code --user-data-dir= "~/.vscode-root" 复制代码`

# 6 总结 #

成品效果如图，不懂的可以加小助手微信 udujjb 拉你进群询问

![](https://user-gold-cdn.xitu.io/2019/5/15/16abb0068007f824?imageView2/0/w/1280/h/960/ignore-error/1)

以上教程是基于 MacOS 的，Linux 用户如何开启透明请参考；Windows 的electron暂不支持vibrancy模式，可以使用插件 [GlassIt-VSC]( https://link.juejin.im?target=https%3A%2F%2Fmarketplace.visualstudio.com%2Fitems%3FitemName%3Ds-nlf-fh.glassit ) 设置透明。

[Custom CSS and JS Loader 配置]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbe5invis%2Fvscode-custom-css%23getting-started )

[Linux 透明窗口]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsergei-dyshel%2Fvscode%2Fblob%2Fmaster%2FREADME.fork.md )

# Windows 只能透明 + 发光 #

一、安装插件：

> 
> 
> 
> 1.SynthWave '84
> 
> 
> 
> 2.Custom CSS and JS
> 
> 
> 
> 3.GlassIt-VSC
> 
> 

二、下载配置文件(两个方法二选一)：

1.git命令安装 git clone [github.com/Jinkeycode/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FJinkeycode%2Fvscode-transparent-glow.git )

2.浏览器访问 [codeload.github.com/Jinkeycode/…]( https://link.juejin.im?target=https%3A%2F%2Fcodeload.github.com%2FJinkeycode%2Fvscode-transparent-glow%2Fzip%2Fmaster )

三、修改vscode的配置文件setting.json(在setting.json文件中加入下面代码):

` "vscode_custom_css.imports" :[ "file:///C:/Users/Administrator/AppData/Roaming/Code/User/vscode-transparent-glow-master/synthwave84.css" , ] 复制代码`

"C:/Users/Administrator/AppData/Roaming/Code/User/vscode-transparent-glow-master/"这个替换成你步骤二中下载到的文件所在位置

四、使配置生效 按下 Ctrl + Shift + P，运行 "Reload Custom CSS and JS", 重启 vscode 即可。