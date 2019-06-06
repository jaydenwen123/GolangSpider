# react-native 滑动吸顶效果 #

> 
> 
> 
> 最近公司开发方向偏向移动端，
> 于是就被调去做RN（react-native），体验还不错，当前有个需求是首页中间吸顶的效果，虽然已经很久没写样式了，不过这种常见样式应该是so-easy，没成想翻车了，网上搜索换了几个方案都不行，最后去github上复制封装好的库来实现，现在把翻车过程记录下来。
> 
> 
> 

### 需求效果 ###

![](https://user-gold-cdn.xitu.io/2019/6/1/16b134016a46203a?imageView2/0/w/1280/h/960/ignore-error/1)

## 翻车过程 ##

### 第一种方案 ` 失败` ###

一开始的思路是这样的，大众思路，我们需要监听页面的滚动状态，当页面滚动到要吸顶元素所处的位置的时候，我们设置它为固定定位，不过很遗憾， RN对于position属性只提供了两种布局方式：absolute和relative，既没有fixed也没有仍处于试验的api：sticky。尴尬了😅

### 第二种方案 ` 失败` ###

不过也不慌，看网上有第二种方案，把图上第二 三块地方作为 ` ScrollView` ，然后 ` ScrollView` 滑动监听距离，把第一块的marginTop设为负值，但是这样第一部分不能滑动，不符合需求，pass

### 第三种方案 ` 完全失败` ###

从网上找到第三种方案，就是一二三部分作为 ` ScrollView` ，

第一部分 ` position` 设为 ` absolute` ，剩下的不设置，默认是relative 第二部分（吸顶部分）marginTop设置（setState）为第一部分高度的state， 添加滑动 **onScroll事件** =》滑动距离y等于第二部分marginTop的state，但是当滑动超过第一部分高度的时候把第二部分（吸顶部分） ` position` 设为 ` absolute` ，并把其marginTop设为0，看起来不错，实际用ios模拟器一跑就无语了😅，效果很奇葩，手指滑动时不吸顶直接划上去隐藏掉大半，一松突然吸顶了。。。

#### 见下图 ####

![](https://user-gold-cdn.xitu.io/2019/6/1/16b12bb45aa7b5fb?imageView2/0/w/1280/h/960/ignore-error/1)

ios的系统，手指在屏幕上滚动时，onScroll一直在触发，如果里面有setState方法，也会不停执行并计算 ` state` ，但是改变react的state是异步的，只要手指不离开屏幕，改变的state就无法生效（触发界面渲染）

## 实现方案 ##

我最终意识到由于ios的机制，react的state机制不能满足需求，RN里面肯定有借助原生渲染的方式，于是github找了现成的代码实现之后，反过来进行研究，大家有RN丰富经验的也可以直接看最下面代码👇

### RN的Animator ###

RN的Animator动画库旨在解决动画问题，由于js桥接过程，动画通常不能很好展现，最好是把动画的 **数据** 和 **变化方法** 一次性发给原生，由原生进行处理，这就是Animator库的核心作用。

记得原来RN的动画一直被吐槽，不过现在效果还挺不错的，可能与近年来手机硬件提升也越来越大也有关系吧。

#### 简单用法 ####

由于Animator内部封装了这四个组件，所以默认可以导出<Animator.View/>,<Animator.Text/>,<Animator.Image/>,<Animator.ScrollView/>

在这几个组件里面想做一些动画处理，数据方面也是react的state，但是赋值要给Animated.Value，如下👇

` this.state = { scrollY: new Animated.Value(0) } 复制代码`

这里虽然使用的还是原生state，但是经过Animated处理，渲染机制完全不一样了

#### 简单原理 ####

经过Animator包装后的组件，会遍历传入的props和自身的state，查找是否有Animated.Value的实例，并绑定进相应的原生操作。 props和自身的state变化时，将Animated.Value值逐个转化为普通数值，再交给原生进行渲染，但是值得注意的是，这里并不会触发react 的 render，更不会有什么domdiff ，是一种特殊处理，类似于Animated.Value改变时每次的shouldUpdateComponent返回都是false（毫秒级的渲染react性能扛不住），shouldUpdateComponent函数里面判断Animated.Value，然后会把数据变化发给原生组件

#### 完整的介绍请移步 [中文官网Animator库介绍]( https://link.juejin.im?target=https%3A%2F%2Freactnative.cn%2Fdocs%2Fanimated%2F ) ####

### 实现思路 ###

既然用了Animator组件了，渲染的问题解决了，下面思路是动态设置吸顶组件的translateY属性。 ` style:{ transform: [{ translateY:translateY }] }`

* 当向下滑动时，不管它
* 向上滑，但是当头部还没有完全隐藏时，也不管它
* 向上滑，头部完全不见了，这时向上再滑一点，那么他的translateY就应该 = 上划总距离 - 头部高度，这样越往上滑，把吸顶组件使劲往下推，这样吸顶组件就牢牢固定在顶部了

下面利用 **插值** 来实现

` const translateY = ScrollY.interpolate({ inputRange : [ -1 , 0 , headerHeight, headerHeight + 1 ], outputRange : [ 0 , 0 , 0 , 1 ], }); 复制代码`

插值 ` interpolate` 略难理解，需要一点基础，这里再细说起来这篇文章就太长了 [官网介绍]( https://link.juejin.im?target=https%3A%2F%2Freactnative.cn%2Fdocs%2Fanimations%2F%23%25E6%258F%2592%25E5%2580%25BC ) ， 如果还不懂可以去网上找找这方面的资料

## 实现源码 ##

> 
> 
> 
> 实现的图中第二部分吸顶功能的核心代码
> 
> 

` import * as React from 'react' ; import { StyleSheet, Animated } from "react-native" ; /** * 滑动吸顶效果组件 * @export * @class StickyHeader */ export default class StickyHeader extends React. Component { static defaultProps = { stickyHeaderY : -1 , stickyScrollY : new Animated.Value( 0 ) } constructor (props) { super (props); this.state = { stickyLayoutY : 0 , }; } // 兼容代码，防止没有传头部高度 _onLayout = ( event ) => { this.setState({ stickyLayoutY : event.nativeEvent.layout.y, }); } render() { const { stickyHeaderY, stickyScrollY, children, style } = this.props const { stickyLayoutY } = this.state let y = stickyHeaderY != -1 ? stickyHeaderY : stickyLayoutY; const translateY = stickyScrollY.interpolate({ inputRange : [ -1 , 0 , y, y + 1 ], outputRange : [ 0 , 0 , 0 , 1 ], }); return ( < Animated.View onLayout = { this._onLayout } style = { [ style , styles.container , { transform: [{ translateY }] } ]} > { children } </ Animated.View > ) } } const styles = StyleSheet.create({ container: { zIndex: 100 }, }); 复制代码`
> 
> 
> 
> 
> 页面里 **实际用法** 如下
> 
> 

` // 在页面constructor里声明state this.state = { scrollY : new Animated.Value( 0 ), headHeight : -1 }; 复制代码` ` < Animated.ScrollView style = {{ flex: 1 }} onScroll = { Animated.event ( [{ nativeEvent: { contentOffset: { y: this.state.scrollY } } // 记录滑动距离 }], { useNativeDriver: true }) // 使用原生动画驱动 } scrollEventThrottle = {1} > < View onLayout = {(e) => { let { height } = e.nativeEvent.layout; this.setState({ headHeight: height }); // 给头部高度赋值 }}> // 里面放入第一部分组件 </ View > < StickyHeader stickyHeaderY = {this.state.headHeight} // 把头部高度传入 stickyScrollY = {this.state.scrollY} // 把滑动距离传入 > // 里面放入第二部分组件 </ StickyHeader > // 这是第三部分的列表组件 < FlatList data = {this.state.dataSource} renderItem = {({item}) => this._createListItem(item)} /> </ Animated.ScrollView > 复制代码`

## 收尾 ##

具体代码就是这样实现了，算是比较完美的方案，特别是照顾了性能，各位可以基于这个封装来实现更复杂的需求，原理大概就是这个原理了，在前端动画领域，自己确实也就刚入门水平，如有问题，请直接指出。

另外，这是我找的那个 **组件** github的 [代码地址：https://github.com/jiasongs/react-native-stickyheader]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjiasongs%2Freact-native-stickyheader ) ，原地址附上，建议如果项目用了给人家一个star