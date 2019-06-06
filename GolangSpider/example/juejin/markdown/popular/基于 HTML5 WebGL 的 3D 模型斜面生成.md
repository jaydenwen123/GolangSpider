# 基于 HTML5 WebGL 的 3D 模型斜面生成 #

## 前言 ##

3D 场景中的面不只有水平面这一个，空间是由无数个面组成的，所以我们有可能会在任意一个面上放置物体，而空间中的面如何确定呢？我们知道，空间中的面可以由一个点和一条法线组成。这个 Demo 左侧为面板，从面板中拖动物体到右侧的 3D 场景中，当然，我鼠标拖动到的位置就是物体放置的点，但是这次我们的重点是如何在斜面上放置模型。

效果图：

![图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b24d45a70f16c7?imageslim)

## 代码生成 ##

#### 创建场景 ####

` dm = new ht.DataModel(); // 数据模型（http://hightopo.com/guide/guide/core/datamodel/ht-datamodel-guide.html） g3d = new ht.graph3d.Graph3dView(dm); // 3D 场景组件（http://hightopo.com/guide/guide/core/3d/ht-3d-guide.html） palette = new ht.widget.Palette(); // 面板组件（http://hightopo.com/guide/guide/plugin/palette/ht-palette-guide.html） splitView = new ht.widget.SplitView(palette, g3d, 'h' , 0.2 ); // 分割组件，第三个参数为分割的方式 h 为左右分，v 为上下分；第四个参数为分割比例，大于 1 的值为绝对宽度，小于 1 则为比例 splitView.addToDOM(); //将分割组件添加进 body 体中 复制代码`

关于这些组件的定义可以到对应的链接里面查看，至于将分割组件添加进 body 体中的 addToDOM 函数有必要解释一下（我每次都提，这个真的很重要！）。

[HT]( https://link.juejin.im?target=http%3A%2F%2Fwww.hightopo.com ) 的组件一般都会嵌入 BorderPane、SplitView 和 TabView 等容器中使用，而最外层的 HT 组件则需要用户手工将 getView() 返回的底层 div 元素添加到页面的 DOM 元素中，这里需要注意的是，当父容器大小变化时，如果父容器是 BorderPane 和 SplitView 等这些HT预定义的容器组件，则 HT 的容器会自动递归调用孩子组件 invalidate 函数通知更新。但如果父容器是原生的 html 元素， 则 HT 组件无法获知需要更新，因此最外层的 HT 组件一般需要监听 window 的窗口大小变化事件，调用最外层组件 invalidate 函数进行更新。

为了最外层组件加载填充满窗口的方便性，HT 的所有组件都有 addToDOM 函数，其实现逻辑如下，其中 iv 是 invalidate 的简写：

` addToDOM = function ( ) { var self = this , view = self.getView(), // 获取组件的底层 div style = view.style; document.body.appendChild(view); // 将组件底层 div 添加进 body 中 style.left = '0' ; // ht 默认将所有的组件的 position 都设置为 absolute 绝对定位 style.right = '0' ; style.top = '0' ; style.bottom = '0' ; window.addEventListener( 'resize' , function ( ) { self.iv(); }, false ); // 窗口大小改变事件，调用刷新函数 } 复制代码`

大家可能注意到了，场景中我添加的斜面实际上就是一个 ht.Node 节点，作为与地平面的参照，在这样的对比下立体感会更强一点。下面是这个节点的定义：

` node = new ht.Node(); node.s3( 1000 , 1 , 1000 ); // 设置节点的大小 node.r3( 0 , 0 , Math.PI/ 4 ); // 设置节点旋转 这个旋转的角度是有学问的，跟下面我们要设置的拖拽放置的位置有关系 node.s( '3d.movable' , false ); // 设置节点在 3d 上不可移动 因为这个节点只是一个参照物，建议是不允许移动 dm.add(node); // 将节点添加进数据容器中 复制代码`

#### 左侧内容构建 ####

![图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b24d45a71877b4?imageView2/0/w/1280/h/960/ignore-error/1) Palette 和 GraphView 类似，由 ht.DataModel 驱动，用 ht.Group 展示分组，ht.Node 展示按钮元素。我将加载 Palette 面板中的图元函数封装为 initPalette，定义如下：

` function initPalette ( ) { // 加载 palette 面板组件中的图元 var arrNode = [ 'displayDevice' , 'cabinetRelative' , 'deskChair' , 'temperature' , 'indoors' , 'monitor' , 'others' ]; var nameArr = [ '展示设施' , '机柜相关' , '桌椅储物' , '温度控制' , '室内' , '视频监控' , '其他' ]; // arrNode 中的 index 与 nameArr 中的一一对应 for ( var i = 0 ; i < arrNode.length; i++) { var name = nameArr[i]; var vName = arrNode[i]; arrNode[i] = new ht.Group(); // palette 面板是将图元都分在“组”里面，然后向“组”中添加图元即可 palette.dm().add(arrNode[i]); // 向 palette 面板组件中添加 group 图元 arrNode[i].setExpanded( true ); // 设置分组为打开的状态 arrNode[i].setName(name); // 设置组的名字 显示在分组上 var imageArr = []; switch (i){ // 根据不同的分组设置每个分组中不同的图元 case 0 : imageArr = [ 'models/机房/展示设施/大屏.png' ]; break ; case 1 : imageArr = [ 'models/机房/机柜相关/配电箱.png' , 'models/机房/机柜相关/室外天线.png' , 'models/机房/机柜相关/机柜1.png' , 'models/机房/机柜相关/机柜2.png' , 'models/机房/机柜相关/机柜3.png' , 'models/机房/机柜相关/机柜4.png' , 'models/机房/机柜相关/电池柜.png' ]; break ; case 2 : imageArr = [ 'models/机房/桌椅储物/储物柜.png' , 'models/机房/桌椅储物/桌子.png' , 'models/机房/桌椅储物/椅子.png' ]; break ; case 3 : imageArr = [ 'models/机房/温度控制/空调精简.png' , 'models/机房/消防设施/消防设备.png' ]; break ; case 4 : imageArr = [ 'models/室内/办公桌简易.png' , 'models/室内/书.png' , 'models/室内/办公桌镜像.png' , 'models/室内/办公椅.png' ]; break ; case 5 : imageArr = [ 'models/机房/视频监控/摄像头方.png' , 'models/机房/视频监控/对讲维护摄像头.png' , 'models/机房/视频监控/微型摄像头.png' ]; break ; default : imageArr = [ 'models/其他/信号塔.png' ]; break ; } setPalNode(imageArr, arrNode[i]); // 创建 palette 上节点及设置名称、显示图片、父子关系 } } 复制代码`

我在 setPalNode 函数中做了一些名称的设置，主要是想要根据上面 initPalette 函数中我传入的路径名称来设置模型的名称以及在不同文件在不同的文件夹下的路径：

` function setPalNode ( imageArr, arr ) { for ( var j = 0 ; j < imageArr.length; j++) { var imageName = imageArr[j]; var jsonUrl = imageName.slice( 0 , imageName.lastIndexOf( '.' )) + '.json' ; // shape3d 中的 json 路径 var name = imageName.slice(imageName.lastIndexOf( '/' )+ 1 , imageName.lastIndexOf( '.' )); // 取最后一个/和.之间的字符串用来设置节点名称 var url = imageName.slice(imageName.indexOf( '/' )+ 1 , imageName.lastIndexOf( '.' )); // 取第一个/和最后一个.之间的字符串用来设置拖拽生成模型 obj 文件的路径 createNode(name, imageName, arr, url, jsonUrl); // 创建节点，这个节点是显示在 palette 面板上 } } 复制代码`

createNode 创建节点的函数比较简单：

` function createNode ( name, image, parent, urlName, jsonUrl ) { // 创建 palette 面板组件上的节点 var node = new ht.Node(); palette.dm().add(node); node.setName(name); // 设置节点名称 palette 面板上显示的文字也是通过这个属性设置名称 node.setImage(image); // 设置节点的图片 node.setParent(parent); // 设置父亲节点 node.s({ 'draggable' : true , // 设置节点可拖拽 'image.stretch' : 'centerUniform' , // 设置节点图片的绘制方式 'label' : '' // 设置节点的 label 为空，这样即使设置了 name 也不会显示在 3d 中的模型下方 }); node.a( 'urlName' , urlName); // a 设置用户自定义属性 node.a( 'jsonUrl' , jsonUrl); return node; } 复制代码`

虽然简单，但是还是要提一下，draggable: true 为设置节点可拖拽，否则节点不可拖拽；还有 node.s 是 HT 默认封装好的样式设置方法，如果用户需要自己添加方法，则可通过 node.a 方法来添加，参数一为用户自定义名称，参数二为用户自定义值，不仅能传常量，也能传变量、对象，还能传函数！又是一个非常强大的功能。

#### 拖拽功能 ####

拖拽基本上就是响应 windows 自带的 dragover 以及 drop 事件，要在放开鼠标的时候创建模型，就要在事件触发时生成模型：

` function dragAndDrop ( ) { // 拖拽功能 g3d.getView().addEventListener( "dragover" , function ( e ) { // 拖拽事件 e.dataTransfer.dropEffect = "copy" ; handleOver(e); }); g3d.getView().addEventListener( "drop" , function ( e ) { // 放开鼠标事件 handleDrop(e); }); } function handleOver ( e ) { e.preventDefault(); // 取消事件的默认动作。 } function handleDrop ( e ) { // 鼠标放开时 e.preventDefault(); // 取消事件的默认动作。 var paletteNode = palette.dm().sm().ld(); // 获取 palette 面板中最后选中的节点 if (paletteNode) { loadObjFunc( 'assets/objs/' + paletteNode.a( 'urlName' ) + '.obj' , 'assets/objs/' + paletteNode.a( 'urlName' ) + '.mtl' , paletteNode.a( 'jsonUrl' ), g3d.getHitPosition(e, [ 0 , 0 , 0 ], [ -1 , 1 , 0 ])); // 加载obj模型 } } 复制代码`

这里完全有必要说明一下，这个 Demo 的重点来了！ loadObjFunc 函数中的最后一个参数为生成模型的 position3d 坐标，g3d.getHitPosition 这个方法总共有三个参数，第一个参数为事件类型，第二和第三个参数如果不设置，则默认为水平面的中心点也就是 [0, 0, 0] 以及法线为 y 轴，也就是 [0, 1, 0]，一条法线和一个点就可以确定一个面，所以我们通过这个方法来设置这个节点所要放置的平面是在哪一个面上，我前面将 node 节点设置为绕 z 轴旋转 45° 角，所以这边的法线也就要好好想想如何设置了，这是数学上的问题，要自己思考了。

#### 加载模型 ####

HT 通过 ht.Default.loadObj 函数来加载模型，但是前提是要有一个节点，然后再在这个节点上加载模型：

` function loadObjFunc ( objUrl, mtlUrl, jsonUrl, p3 ) { // 加载 obj 模型 var node = new ht.Node(); var shape3d = jsonUrl.slice(jsonUrl.lastIndexOf( '/' )+ 1 , jsonUrl.lastIndexOf( '.' )); ht.Default.loadObj(objUrl, mtlUrl, { // HT 通过 loadObj 函数来加载 obj 模型 cube: true , // 是否将模型缩放到单位 1 的尺寸范围内，默认为 false center: true , // 模型是否居中，默认为 false，设置为 true 则会移动模型位置使其内容居中 shape3d: shape3d, // 如果指定了 shape3d 名称，则 HT 将自动将加载解析后的所有材质模型构建成数组的方式，以该名称进行注册 finishFunc: function ( modelMap, array, rawS3 ) { // 用于加载后的回调处理 if (modelMap) { node.s({ // 设置节点样式 'shape3d' : jsonUrl, // jsonUrl 为 obj 模型的 json 文件路径 'label' : '' // 设置 label 为空，label 的优先级高于 name，所以即使设置了 name，节点的下方也不会显示 name名称 }); g3d.dm().add(node); // 将节点添加进数据容器中 node.s3(rawS3); // 设置节点大小 rawS3 模型的原始尺寸 node.p3(p3); // 设置节点的三维坐标 node.setName(shape3d); // 设置节点名称 node.setElevation(node.s3()[ 1 ]/ 2 ); // 控制 Node 图元中心位置所在 3D 坐标系的y轴位置 g3d.sm().ss(node); // 设置选中当前节点 g3d.setFocus(node); // 将焦点设置在当前节点上 return node; } } }); } 复制代码`

代码结束！

## 总结 ##

说实在的这个 Demo 真的是非常容易，难度可能在于空间思维能力了，先确认法线和点，然后根据法线和点找到那个面，这个面按照我的这种方式有个对照还比较能够理解，真幻想的话，可能容易串。这个 Demo 容易主要还是因为封装的 hitPosition 函数简单好用，这个真的是功不可没。