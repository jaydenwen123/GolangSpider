# 原生 JS 撸一个轮播图（支持拖拽切屏） #

> 
> 
> 
> ## FoxSlider.js 称不上库的库 ##
> 
> 

## 1、简述 ##

用惯了各种各样的组件，并没有真正意义上的封装一个可以拖拽切屏的轮播图，经过一番编写，发现写这样一个轮播图要想写的好，还真的是挺有难度，不同设备的不同事件完备性？事件触发时机的把控？如何更好的去封装？自适应窗口后的事件重置？等等...，看看swiper这个库的源码，就知道小事情也可以不简单。
现在写的这个基本的需求是可以满足的，可以通过拖拽切换也可以点击切换。

> 
> 
> 
> [github 传送门（想你来一起完（wan）善(shua)！！Fork & Star ^_^一下你就会很美](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fforrestyuan%2FFoxSlider.js
> )
> 
> 

> 
> 
> 
> 传送点不了的可以复制链接： [github.com/forrestyuan…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fforrestyuan%2FFoxSlider.js
> )
> 
> 

> 
> 
> 
> **原生撸码一时爽，一直原生一直爽**
> 
> 

### 1.1、特性 ###

* 面向手机、平板，PC等终端。

### 1.2、缺点 ###

* 封装简陋，功能亟需扩充
* 语义化不够强
* 用户配置能力弱

## 2、调用实例 ##

> 
> 
> 
> html 结构代码
> 
> 

` <!-- 引入js文件 --> < script src = "js/base.js" > </ script > <!-- 主要dom结构 --> < div class = "slider-container" > < div class = "slide-bar" > < div class = "slider " > < img src = "assets/slider1.jpg" alt = "" > </ div > < div class = "slider" > < img src = "assets/slider2.jpg" alt = "" > </ div > < div class = "slider" > < img src = "assets/slider3.jpg" alt = "" > </ div > </ div > < div class = "slider-pin" > < span class = "pin on" > </ span > < span class = "pin" > </ span > < span class = "pin" > </ span > </ div > </ div > 复制代码`
> 
> 
> 
> 
> js 代码
> 
> 

` //实例化TouchPlugin，传入参数 var tp = new TouchPlugin({ sliderContainer : '.slider-container' , slider : '.slider' , slidePin : '.slider-pin' , sliderBar : '.slide-bar' , pinClassName : 'on' , pin : '.pin' , callback : function ( e, dir, distance ) { console.log(dir, distance); } }); 复制代码`
> 
> 
> 
> 
> 运行效果
> 
> 

![运行效果图](https://user-gold-cdn.xitu.io/2019/6/5/16b27b7957e1305d?imageslim)

## 3、 ` base.js` 内主要方法 ##

> 
> 
> 
> ### init() ###
> 
> 

初始化函数

**Kind** : global function

### refreshParam(totalMoveLen, spinIndex) ###

刷新参数

**Kind** : global function

+--------------+-----------+------------------+
|    PARAM     |   TYPE    |   DESCRIPTION    |
+--------------+-----------+------------------+
| totalMoveLen | ` number` | 滚动位移         |
| spinIndex    | ` number` | 轮播指标高亮下标 |
+--------------+-----------+------------------+

### setTranslate(domNode, conf, moveLen) ###

设置指定对象移动样式 （transform）

**Kind** : global function

+---------+-----------+------------------------------------+
|  PARAM  |   TYPE    |            DESCRIPTION             |
+---------+-----------+------------------------------------+
| domNode | ` Object` | 应用移动样式的对象                 |
| conf    | ` Object` | 配置对象（animateStyle:            |
|         |           | ease-in-out                        |
| moveLen | ` number` | 轮播图移动距离（切屏通过控制位移） |
+---------+-----------+------------------------------------+

### resize() ###

改变屏幕尺寸重置参数

**Kind** : global function

### autoRun(time, initStep) ###

自动轮播

**Kind** : global function

+----------+-----------+-------------------+
|  PARAM   |   TYPE    |    DESCRIPTION    |
+----------+-----------+-------------------+
| time     | ` number` | 轮播时间          |
| initStep | ` number` | spin下标 和下一屏 |
+----------+-----------+-------------------+

### judgeDir(curX, preX) ###

判断鼠标或触摸移动的方向

**Kind** : global function

+-------+-----------+---------------------------------+
| PARAM |   TYPE    |           DESCRIPTION           |
+-------+-----------+---------------------------------+
| curX  | ` number` | 鼠标点击或开始触摸屏幕时的坐标X |
| preX  | ` numer`  | 鼠标移动或触摸移动时的坐标X     |
+-------+-----------+---------------------------------+

### testTouchEvent() ###

检测当前设备支持的事件（鼠标点击移动和手触摸移动）

**Kind** : global function

### mouseX(event) ###

获取鼠标横坐标位置

**Kind** : global function

+-------+----------+-------------+
| PARAM |   TYPE   | DESCRIPTION |
+-------+----------+-------------+
| event | ` Event` | 事件对象    |
+-------+----------+-------------+

### cancelBind(domNode) ###

取消绑定触摸或鼠标点击移动事件

**Kind** : global function

+---------+-----------+----------------------------+
|  PARAM  |   TYPE    |        DESCRIPTION         |
+---------+-----------+----------------------------+
| domNode | ` Object` | 需要被取消绑定事件节点对象 |
+---------+-----------+----------------------------+

### reBindTouchEvent(domNode, callback, isUnbind) ###

重新绑定触摸事件

**Kind** : global function

+----------+-------------+------------------------+
|  PARAM   |    TYPE     |      DESCRIPTION       |
+----------+-------------+------------------------+
| domNode  | ` Object`   | 需要被绑定事件节点对象 |
| callback | ` function` | 回调方法               |
| isUnbind | ` boolean`  | 是否取消绑定           |
+----------+-------------+------------------------+

### removeClsName(nodeList, clsName) ###

移除节点的样式类名

**Kind** : global function

+----------+-----------+----------------------+
|  PARAM   |   TYPE    |     DESCRIPTION      |
+----------+-----------+----------------------+
| nodeList | ` Array`  | 被移除样式的节点数组 |
| clsName  | ` string` | 移除的样式类名称     |
+----------+-----------+----------------------+

### setClsName(node, clsName) ###

添加样式

**Kind** : global function

+---------+-----------+----------------+
|  PARAM  |   TYPE    |  DESCRIPTION   |
+---------+-----------+----------------+
| node    | ` Object` | 添加类名的节点 |
| clsName | ` string` | 样式类名       |
+---------+-----------+----------------+

### bindSpinClick() ###

点击轮播spin 切换屏

**Kind** : global function

### checkTargetByCls(domNode, clsName) ###

通过检测dom节点是否包含某个样式名来判断是否属于目标target

**Kind** : global function

+---------+-----------+
|  PARAM  |   TYPE    |
+---------+-----------+
| domNode | ` Object` |
| clsName | ` string` |
+---------+-----------+

### bindTouchEvent(domNode, callback, isUnbind) ###

**Kind** : global function

+----------+-------------+--------------------+
|  PARAM   |    TYPE     |    DESCRIPTION     |
+----------+-------------+--------------------+
| domNode  | ` Object`   | 绑定事件的代理对象 |
| callback | ` function` | 回调方法           |
| isUnbind | ` boolean`  | 是否取消绑定       |
+----------+-------------+--------------------+