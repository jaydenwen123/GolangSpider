# 开启 Cocoapods 新选项，加快项目索引速度 #

前几天 Cocoapods 1.7.0 正式版发布了，我最期待的一个功能是 Multiple Pod Projects，昨天顺手就给接入了，项目解析和索引效率有了非常明显的提升，过程中踩了些坑，这次一起把之前 debug 的经验分享一下。

## generate_multiple_pod_projects 选项 ##

之前 Cocoapods 会把每个依赖作为 target 放到 Pods 项目里，但 xcodeproj 本身的编码不太能适应这种情况，在引入几十个 pod 的情况下，项目解析的效率会急剧下降。

以我公司其中一个主项目为例，Pods 项目的大小达到了 5.2 MB（这可都是纯文本），在第一次打开项目，解析项目构建索引时，就能明显听到风扇开始狂转，这个过程会持续好几分钟才会结束。

Cocoapods 这次更新引入了一个 generate_multiple_pod_projects 的选项，可以让每个依赖都作为一个单独的项目引入，大大增加了解析速度：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b09deb23c6f022?imageView2/0/w/1280/h/960/ignore-error/1)

开启的方式很简单，只要在 Podfile 里加入这一行就可以了：

` install! 'cocoapods' , generate_multiple_pod_projects: true 复制代码`

拆分后每个项目的大小都差不多是 40 - 100 kb 左右：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b09df22915c2ba?imageView2/0/w/1280/h/960/ignore-error/1)

这个选项开启之后的效果非常显著，我在 Xcode 里执行了 clean，之后 indexing 的过程在几秒钟里就结束了，而且风扇也没有狂转。

至于为什么这样可以提升项目的解析速度，我大概看了一下 xcodeproj 的编码，所有的 Item 都会按照类别存放到各自的 section 里，最终在项目的结构树里会以引用的形式呈现。

所以文件引用查找的范围是所有 Pod 引用库的文件的集合，而每次索引的构建都至少会遍历一次项目树，这就会导致索引时间的暴增，除此之外单个庞大的项目解析也不利于多线程执行，拆分成多个项目的话就能有效地解决这些问题。

## install! 函数只能调用一次 ##

需要注意 ` install!` 是个用来配置的函数，由于之前我还开启了另一个选项，所以接入时是这么做的：

` install! 'cocoapods' , generate_multiple_pod_projects: true install! 'cocoapods' , disable_input_output_paths: true 复制代码`

但是这么做之后发现不生效，后来才想起来 ` install!` 是一个用来配置的函数，重复调用的话，只会以最后一次的调用为准。所以应该在一次调用里把它们都传入进去：

` install! 'cocoapods' , disable_input_output_paths: true , generate_multiple_pod_projects: true 复制代码`

## Swift 版本控制 ##

另一个坑就是在 ` post_install` 时，为了一些版本的兼容，需要遍历所有 target，调整一部分库的 Swift 版本：

` post_install do |installer| swift_4_0_compatible = [ ... ] swift_4_2_compatible = [ ... ] installer.pod_targets.each do |t| t.build_configurations.each do |c| c.build_settings[ 'SWIFT_VERSION' ] = '4.0' if swift_4_0_compatible. include ? t.name c.build_settings[ 'SWIFT_VERSION' ] = '4.2' if swift_4_2_compatible. include ? t.name end end end 复制代码`

但是如果开启了 ` generate_multiple_pod_projects` 的话，由于项目结构的变化， ` installer.pod_targets` 就没办法获得所有 pods 引入的 target 了。

### Podfile 里的代码如何 debug ###

查了 Xcodeproj 和 Cocoapods 的文档之后我都没有得到很好的解答，所以我就想用 xcodeproj 本身的接口去处理这件事情。

由于 Podfile 本质上是 Ruby 脚本，所以这里我通常会使用 Ruby 的 debugger 去操作，通过 Ruby 强大的自省能力，在 debugger 里进行尝试然后找到我们需要的接口，开始之前我们需要安装一个 Ruby 的工具，步骤如下：

* 首先是安装 debugger ` gem install pry`
* 接着在 Podfile 的开头导入 ` require 'pry'`
* 然后在我们想要插入断点的地方插入 ` binding.pry` 语句就可以了

### 查找能用的接口 ###

我在 post_install 里插入了断点，接着运行 ` pod install` ，就看到断点生效了：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b09e0bde5b0580?imageView2/0/w/1280/h/960/ignore-error/1)

Ruby 的自省能力非常强大，而且 pry 也基于此做了很多实用的功能，在这里我直接输入了 ` installer` 回车，就能看到它所有属性都被递归打印出来。

这里面我找了一下之后，发现一个文档里没有记录的属性，叫做 ` pod_target_subprojects` ，包含了所有 Pods 的项目，似乎可以满足我们的需求：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b09dfe1db06bc1?imageView2/0/w/1280/h/960/ignore-error/1)

接着 Ctrl + d 退出 pry，回到 Podfile 修改即可：

` post_install do |installer| swift_4_0_compatible = [ ... ] swift_4_2_compatible = [ ... ] installer.pod_target_subprojects.flat_map { |p| p.targets }.each do |t| t.build_configurations.each do |c| c.build_settings[ 'SWIFT_VERSION' ] = '4.0' if swift_4_0_compatible. include ? t.name c.build_settings[ 'SWIFT_VERSION' ] = '4.2' if swift_4_2_compatible. include ? t.name end end end 复制代码`

最后 ` pod install` 一下，打开 Xcode 查看对应的 target 的编译设置，确实有效。

这里介绍的 debug 方法在 fastlane 里也适用，非常建议大家在编写复杂脚本时先用 debugger 去提前踩坑。

## 结语 ##

用惯了 Ruby 的 debug 方式之后，回到 LLDB 感觉体验差了很多😂。

> 
> 
> 
> 如果觉得文章还不错的话可以关注一下我的 [博客](
> https://link.juejin.im?target=https%3A%2F%2Fkemchenj.github.io%2F2019-05-31%2F
> ) 。
> 
>