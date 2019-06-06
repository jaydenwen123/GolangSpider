# 【译】系统的学习 Android Touch #

原标题: Mastering the Android Touch System

原文地址: [speakerd.s3.amazonaws.com]( https://link.juejin.im?target=https%3A%2F%2Fspeakerd.s3.amazonaws.com%2Fpresentations%2F55e1e8207eaf427a8955e1c6abd5386c%2FMasteringTouch.pdf )

原文作者: Dave Smith

## 主题涵盖 ##

* 触摸系统概述
* 触摸事件框架
* 处理自定义触摸事件
* 系统提供的触摸机制
* 系统提供的手势机制

## 如何处理 Andriod 触摸事件 ##

* 

每个用户的触摸事件都被包装为动作事件

* 

描述用户当前的操作

* ` ACTION_DOWN`
* ` ACTION_UP`
* ` ACTION_MOVE`
* ` ACTION_POINTER_DOWN`
* ` ACTION_POINTER_UP`
* ` ACTION_CANCEL`

* 

包括原数据的事件

* 触摸位置
* 触控点的数量(手指)
* 事件发生的时间

* 

一个”手势”从 ` ACTION_DOWN` 开始到 ` ACTION_UP` 结束

* 

事件是从 ` Activity` 的 ` dispatchTouchEvent()` 开始，通过视图从上而下的形式传递，父视图( ` ViewGroups` )将事件分发给子视图，事件在传递的过程是可以随时拦截的。它会沿着关系链向下传递直到该事件被消费掉。任何未拦截的事件都会传递到 ` Activity` 的 ` onTouchEvent()` 后结束。

* 

` Activity.dispatchTouchEvent()` ,总是被首先调用，然后将事件发送到附加的 ` Window` 根视图中， ` onTouchEvent()` ，如果没有视图消耗事件是调用，总是持续的调用状态中。

* 

` View.dispatchTouchEvent()` ，如果存在，则首先将事件发送给侦听器 ` View.OnTouchListener.onTouch()` ，如果没有消耗则处理事件本身， ` View.onTouchEvent()` 。

* 

` ViewGroup.dispatchTouchEvent()` 中的 ` onInterceptTouchEvent()` ，会检查它是否应该取代子视图。而对于每个子视图，如果事件是相关的（内部视图）则以相反的顺序添加它们。如果 ` child.dispatchTouchEvent()` 以前没有处理则继续传递到下一个视图，直到该事件被消耗（与View相同）。事件拦截（ ` onInterceptTouchEvent()` 返回 ` true` ）会将 ` ACTION_CANCEL` 传递给子 Activity,所有即将的事件都由ViewGroup直接处理。子视图可以调用 ` requestDisallowTouchIntercept()` 来阻止 ` onInterceptTouchEvent()` 继续持有当前手势的时间。 每个新手势( ` ACTION_DOWN` )都会重置 ` fragmework` 的标识。

### 错误的视图案例 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b216df8e603730?imageView2/0/w/1280/h/960/ignore-error/1)

### 有趣的视图案例 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b216f2b78e66dd?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/4/16b216fb25540178?imageView2/0/w/1280/h/960/ignore-error/1)

## 处理自定义事件 ##

处理触摸事件 -Subclass重写onTouchEvent（） - 提供OnTouchListener •消费事件

* 使用ACTION_DOWN返回true以显示兴趣 •即使您对ACTION_DOWN不感兴趣，也请返回true
* 对于其他事件，返回true只会停止进一步处理 •ViewConfiguraCon -getScaledTouchSlop（）中可用的有用常量 •距离移动事件可能会在被视为拖动之前发生变化-getScaledMinimumFlingVelocity（） •系统认为拖动为速度的速度-getScaledPagingTouchSlop（） •触摸用于水平分页手势的slop（即ViewPager） - 为每个设备的密度缩放显示的值

### 处理事件 ###

子类重写 ` onTouchEvent()` 方法，并提供一个 ` OnTouchListener` 。使用 ` ACTION_DOWN` 并返回 ` true` 表示消耗该事件即使您对 ` ACTION_DOWN` 不大算消耗该事件也请返回 ` true` ,对于其他事件，返回 ` true` 会停止事件的进一步处理。

在 ` ViewConfiguration` 中有用的常量：

* ` getScaledTouchSlop()` ：移动距离的事件可能会在其拖动之前就会发生变化
* ` getScaledMinimumFlingVelocity()` :系统认为快速滑动是一种惯性拖拽
* ` getScaledPagingTouchSlop()` :事件池使用一个水平分页手势（i.e. ViewPager）

以上内容是 Mastering the Android Touch System PPT 1-10页的内容总结，文章有些术语及方法释明需要调整校对。剩下的10页会陆续补上。如果有翻译不妥的地方，欢迎大家提出，一起完善。

## 欢迎关注 Kotlin 中文社区！ ##

### 中文官网： [www.kotlincn.net/]( https://link.juejin.im?target=https%3A%2F%2Fwww.kotlincn.net%2F ) ###

### 中文官方博客： [www.kotliner.cn/]( https://link.juejin.im?target=https%3A%2F%2Fwww.kotliner.cn%2F ) ###

### 公众号：Kotlin ###

### 知乎专栏： [Kotlin]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fbennyhuo ) ###

### CSDN： [Kotlin中文社区]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_23626713 ) ###

### 掘金： [Kotlin中文社区]( https://juejin.im/user/5cea6293e51d45775e33f4dd/posts ) ###

### 简书： [Kotlin中文社区]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fu%2Fa324daa6fa19 ) ###