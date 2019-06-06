# Ant Design中DatePicker设置mode="year"无法获取value及关闭面板的解决方案 #

## 1、前情提要 ##

当初还是 ` antd2.X` 版本时， ` DatePicker` 组件还不支持 ` mode` 属性，不能单独设置为年份选择器。但是公司项目刚好很多地方都有根据年份做筛选的需求，因为 ` antd` 不支持，因此，使用了 ` Select` 组件来实现年份选择。

但是，遭到了客户的强烈吐槽，“你们这个UI风格还是要一致撒”，哈哈😄，官方吐槽最为致命！没办法了，我自己也没法说服自己了，只能照着 ` antd` 的UI风格自己撸一个 ` YearPicker` 咯。（ 时间选择控件 ` YearPicker` 基于 ` React` ， ` antd` [www.cnblogs.com/zyl-Tara/p/…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fzyl-Tara%2Fp%2F9133524.html ) ）但是，老实说，效果不怎么理想，只能说实现了UI风格的一致，以及值的选择。但是，组件值的清除、动画过渡效果等都没有深入处理。

庆幸的是，很快 ` antd3.X` 终于支持了年份选择。设置 ` mode="year"` 便可以使用年份选择器。真是普天同庆！😎

## 2、问题描述 ##

话不多说，赶紧用起来！上代码，

` import React, { Component } from 'react' ; import { DatePicker } from 'antd' ; export default class extends Component { onChange = val => { console.log(val) } render () { return ( <div> <DatePicker placeholder= "请选择年份" mode= "year" onChange={this.onChange} /> </div> ); } } 复制代码`

界面呈现出只有年份的选择器，nice！❤️❤️

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22943ad86d587?imageView2/0/w/1280/h/960/ignore-error/1)

但是，接下来，你就直接懵了。因为不管你怎么点击按钮选择年份都不会起作用， ` onChange` 事件根本不会触发，所以 ` value` 获取不了！

哇，好气哦 ！😡😡

百思不得其解，然后去翻看 ` ant design` 的 ` github` ` issue` 。终于看到一条中肯的 ` comment` 。

[github.com/ant-design/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fant-design%2Fant-design%2Fissues%2F10242%23issuecomment-393055938 )

` <DatePicker mode= "year" onPanelChange={(v) => {console.log(v)}}/> 复制代码`

只需要把 ` onChange` 换成 ` onPanelChange` 就好了。于是可以愉快的获取时间了。

然而，另一个问题出现了，时间虽然是获取了，但是面板并没有关闭。😠😠

继续查找问题，发现当 ` DatePicker` 变为受控后，需要 ` open` 这个属性控制面板的关闭。

这个大概是 ` antd` 团队没考虑到？🤔️🤔️

` import React, { Component } from 'react' ; import { DatePicker } from 'antd' ; export default class extends Component { state = { isopen: false , time: null } render () { const { isopen, time } = this.state return ( <div> <DatePicker value={time} open={isopen} mode= "year" placeholder= "请选择年份" format= "YYYY" onFocus={() => {this.setState({isopen: true })}} onBlur={() => {this.setState({isopen: false })}} onPanelChange={(v) => { console.log(v) this.setState({ time: v, isopen: false }) }}/> </div> ); } } 复制代码`

同时，通过 ` onFocus` 和 ` onBlur` 控制获取焦点和失焦时面板的显隐。

一切似乎很完美，能够获取值，也能正常关闭面板。

然而，快乐的时光总是短暂的，很快，测试便提出缺陷，“这个年份选择器为什么选择完年份后会有闪开闪关的效果，不符合要求哈”

“😂，哦，那我再看看呢”

还真的是会闪，明明记得之前没有这个问题啊，算了继续看看问题在哪儿吧。

## 3、解决方案 ##

查看文档，发现 ` DatePicker` 有个 ` onOpenChange` 方法，是这样描述的：弹出日历和关闭日历的回调， ` function(status)`

所以，可以通过 ` onOpenChange` 方法判断当前的操作是要面板关闭还是打开，来控制面板的显隐。

所以，综上所有的思考，解决思路如下：

1、 ` onChange` 方法无法触发获取到值，需要换成 ` onPanelChange`

2、面板的显示隐藏需要 ` open` 属性进行手动控制

3、 ` onFocus` 、 ` onBlur` 会导致闪开闪关，需要换成 ` onOpenChange`

` import React, { Component } from 'react' ; import { DatePicker } from 'antd' ; export default class extends Component { state = { isopen: false , time: null } render () { const { isopen, time } = this.state return ( <div> <DatePicker value={time} open={isopen} mode= "year" placeholder= "请选择年份" format= "YYYY" onOpenChange={(status) => { if (status){ this.setState({isopen: true }) } else { this.setState({isopen: false }) } }} onPanelChange={(v) => { console.log(v) this.setState({ time: v, isopen: false }) }} /> onChange={() => { this.setState({time: null}) }} </div> ); } } 复制代码`

现在可以正常获取值，并且开关面板流畅，不会出现闪开闪关的效果。当然，细心你可能发现一个秘密，就是我在组件中用到了 ` onChange` 事件，并且做了对值置空的操作。

注意：这里的 ` time` 置空一定要设置为 ` null` 。因为组件接受的是一个对象。

这是为什么呢？

我们都知道 ` DatePicker` 组件有一个 ` allowClear` 属性，让我们可以通过点击输入框中的❌icon来清除选择的值。

但是当我们设置 ` mode=“year”` 后，这个 ` allowClear` 便不起作用了。怎么办呢？

因为 ` onChange` 事件不会在选择值的时候触发，但是点击清除icon 却会触发。因此通过 ` onChange` 事件便可以达到清除 ` value` 的效果。

ok，完美解决～ 🎉🎉