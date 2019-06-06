# 重学安卓：Activity 的快乐你不懂！ #

> 
> 
> 
> 本文由作者 KunMinX 原创，已授权某公众号独家转载 🔥
> 
> 

## 前言 ##

本文本来是自己复盘 Android 知识梳理用的，没想到在上周部门内部的知识测评中发现，同事们对这些基础知识的掌握参差不齐，甚至可以说是模棱两可。

是网上关于 Activity 的教程太少了吗？不是的，恰恰相反，网上的信息多如牛毛， **却没有一篇愿意费哪怕一丝丝的笔墨** 来介绍 Activity 的起源、它的职责边界、它的存在到底是为了解决什么问题、我们学习它，到底学到什么程度才算掌握。

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af90e287ad6?imageView2/0/w/1280/h/960/ignore-error/1)

正因为对这些最基本而必要的概念模棱两可，使得教程再多、再优秀，也没多少人能消化、能记住。于是我抱着试试看的心态，在经过几番润色过，将自己复盘的结果，在小会上分享出来供同事们享用。想不到原本不屑听的几个同事，在听完这番讲解后，连说真香。

所以如果你因为本文，而对 ` Activity` 乃至向上追溯的 ` View` 、 ` Window` 、 ` WindowManager` 、 ` WindowManagerService` 、 ` Surface` 、 ` Surface Flinger` 各自的起源、职责边界以及相互间的关系有了最基本的感性认识，继而不知不觉地开始有了一丝丝好奇，推动你深入地去探究，那我的愿望也就达到了。

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af90e361352?imageView2/0/w/1280/h/960/ignore-error/1)

## 我是一块板砖 ##

我是一块运行着原始 Android 系统的板砖。我有一块屏幕，人们只要通过硬件抽象层（HAL）的代码对屏幕发起指令，屏幕上就可以显示人们想看到的内容。然而这么做过于原始，也不契合板砖的使用场景。

于是有人考虑在 HAL 之上的运行时层（ART）用 C++ 封装一个服务，该服务的名称就叫 Surface Flinger。

## 我是 Surface Flinger ##

我是 Surface Flinger，我的职责是专门负责 UI 内容的渲染。

人们想要在屏幕上渲染出什么内容，都可以通过我来间接地与屏幕打交道。这就好比你在电脑上排版好的文档，只需通过打印机驱动程序这个中介，就能帮助你将文档内容输出到纸上。

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af90e46535b?imageView2/0/w/1280/h/960/ignore-error/1)

至于内容本身究竟有些什么，这我不管，我只负责统一地、有序地将内容安排成输出设备能理解的方式，来实现输出。

## 我是 Window ##

这块板砖的主人不仅想要渲染 UI，还想要窗口，于是在应用框架层，通过 Java 封装了我。

人如其名，我就是一个窗口，我负责可视化内容的排版，然后将排版结果，通过我的上司 WindowManager，通过进程通信的方式，去与后台服务 WindowManagerService 通信，最终递交到 Surface Flinger 来输出和呈现。

Surface Flinger 为我们每一个 Window 都映射了一块 Surface，来用于管理和渲染屏幕内容。

然而作为一个 Window，我也有我的苦衷。

## 我是一个会套娃的 View ##

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af90e487720?imageView2/0/w/1280/h/960/ignore-error/1)

主人因为经常听 Window 大哥抱怨排版的负担太重，于是用组合模式封装了我。我的 “有容乃大” 版本：ViewGroup，因为组合模式，而能够在自身内部存在更多的 View 或 ViewGroup，这使得我们从结构上来看，就像套娃。

托递归的福，我们的排版工作：Measure、Layout、Draw，可以自己通过如此般的递归，自下往上地完成。然后 Window 大哥就可以直接拿着我们的排版结果，去向上司交差啦。

## …… 所以，Window 成了摸鱼般的存在吗？ ##

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af916b80761?imageslim)

本来 Window 正寻思着，日子过得这般清闲自在，没想到好日子到头 —— 主人不仅要一个窗口，还想要多窗口。这多窗口它就涉及到窗口间的切换、通信等等，甚是麻烦，这些脏活累活要是交给以后的开发者来干，那我不得留下一世骂名、遗臭万年？？！

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af916c4a446?imageView2/0/w/1280/h/960/ignore-error/1)

想到这里我就感觉哆嗦，不行，为了我一世英名，我得向主人进言。

其实早在 20000 多年前，女娲造人的时候，便采用了神级的模板方法模式，将一系列的通用功能都封装好，只暴露一些 DNA 接口，以供后来者随机输入和演变。

换言之，主人只需以模板方法模式的方式将我重新封装，并且编写一套管理窗口的任务和返回栈机制在背地里运筹帷幄，那么未来的开发者就只需继承我，而得到一个简练的配置模板，从而在模板上面输入他们的定制内容，以得到他们想要的结果。

## 于是我改名叫 Activity ##

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3af9d82b8a34?imageView2/0/w/1280/h/960/ignore-error/1)

Window 成了我永恒不变的信仰，存留在我的体内。对于开发者来说，我就是个待继承的 Activity，开发者通过继承我，拿到的就是一个个简练的模板。

对系统来说，我的本质仍是被管理的窗口，系统能够管理我和其他窗口的切换和通信。

对开发者来说，我的本质是视图控制器，开发者通过我可以控制 View 以他们想要的方式进行排版，并且在特殊状况下保存和恢复 View 的排版内容。

## 综上 ##

最开始只有一块运行着原始 Android 系统的板砖。

Surface Flinger 的出现是为了更加方便地完成 UI 渲染。

Window 的出现是为了管理 UI 内容的排版。

Window 不堪重负于是将责任下发到 View 身上。

View 通过组合模式，在递归的帮助下蹭蹭蹭地完成排版工作。

Activity 的出现是为了满足多窗口管理和傻瓜式视图管理的需要。

所以 Activity 的知识边界无非就是生命周期、特殊状况导致的重建、多窗口跳转（启动模式、intent）、视图的加载和优化等等。

这样说，你理解了吗？

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae3b18a95b5e2d?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/24/16aea2340493348e?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> [《重学安卓》系列文章](
> https://link.juejin.im?target=https%3A%2F%2Fxiaozhuanlan.com%2Fkunminx%3Frel%3Dkunminx
> )
> 
> 
> 
> [重学安卓：Activity 的快乐你不懂！](
> https://link.juejin.im?target=https%3A%2F%2Fxiaozhuanlan.com%2Fkunminx%3Frel%3Dkunminx
> )
> 
> 
> 
> [重学安卓：Activity 生命周期的 3 个辟谣](
> https://link.juejin.im?target=https%3A%2F%2Fxiaozhuanlan.com%2Fkunminx%3Frel%3Dkunminx
> )
> 
> 
> 
> [重学安卓：绝不丢失状态的 Activity 重建机制](
> https://link.juejin.im?target=https%3A%2F%2Fxiaozhuanlan.com%2Fkunminx%3Frel%3Dkunminx
> )
> 
>