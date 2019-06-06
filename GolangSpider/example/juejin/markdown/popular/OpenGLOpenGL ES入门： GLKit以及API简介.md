# OpenGL/OpenGL ES入门： GLKit以及API简介 #

> 
> 
> 
> 系列推荐文章：
> [OpenGL/OpenGL ES入门：图形API以及专业名词解析](
> https://juejin.im/post/5cd9816f6fb9a032297b1e1a )
> [OpenGL/OpenGL ES入门：渲染流程以及固定存储着色器](
> https://juejin.im/post/5cdae5486fb9a0323e3ade57 )
> [OpenGL/OpenGL ES入门：图像渲染实现以及渲染问题](
> https://juejin.im/post/5cdd7f0de51d453b44023729 )
> [OpenGL/OpenGL ES入门：基础变换 - 初识向量/矩阵](
> https://juejin.im/post/5ce129a95188256a220236a0 )
> [OpenGL/OpenGL ES入门：纹理初探 - 常用API解析](
> https://juejin.im/post/5cea3a44e51d455d6d53578e )
> [OpenGL/OpenGL ES入门： 纹理应用 - 纹理坐标及案例解析(金字塔)](
> https://juejin.im/post/5cea8f5bf265da1b76387e02 )
> [OpenGL/OpenGL ES入门： 顶点着色器与片元着色器（OpenGL过渡OpenGL ES）](
> https://juejin.im/post/5cf3d29ff265da1b667bc3d0 )
> [OpenGL/OpenGL ES入门： GLKit以及API简介](
> https://juejin.im/post/5cf692e66fb9a07ead59e846 )
> [OpenGL/OpenGL ES入门： GLKit使用以及案例](
> https://juejin.im/post/5cf68ed26fb9a07ede0b30b7 )
> 
> 

## GLKit简介 ##

**` GLKit`** 框架的设计目标是为了简化基于 ` OpenGL/OpenGL ES` 的应用开发。它的出现加快 ` OpenGL` 或 ` OpenGL ES` 应用程序开发。 使用数学库，背景纹理加载，预先创建的着色器效果，以及标准视图和视图控制器来实现渲染循环。

**` GLKit`** 框架提供了功能和类，可以减少创建新的基于着色器的应用程序所需的工作量，或支持依赖早期版本的 ` OpenGL` 或 ` OpenGL ES` 提供的固定函数顶点或片段处理的现有应用程序。

` GLKView` 提供绘制场所（ ` view` ）
` GLKViewController` 扩展于标准的 ` UIKit` 设计模式，用于绘制视图内容的管理与呈现
**苹果弃用 ` OpenGL ES` ，但iOS开发者可以继续使用。**

具体使用请看下一章内容 [OpenGL/OpenGL ES入门： GLKit使用以及案例]( https://juejin.im/post/5cf68ed26fb9a07ede0b30b7 )

## GLKit功能 ##

* 加载纹理
* 提供高性能的数学运算
* 提供常见的着色器
* 提供视图以及视图控制器

## GLKit纹理加载 ##

### GLKTextureInfo创建OpenGL纹理信息。 ###

` name : `OpenGL` 上下⽂文中纹理理名称 target : 纹理理绑定的⽬目标 height : 加载的纹理理⾼高度 width : 加载纹理理的宽度 textureOrigin : 加载纹理理中的原点位置 alphaState: 加载纹理理中alpha分量量状态 containsMipmaps: 布尔值,加载的纹理理是否包含mip贴图 复制代码`

### GLTextureLoader 简化从各种资源文件中加载纹理. ###

* 初始化

` - initWithSharegroup: 初始化⼀个新的纹理加载到对象中 - initWithShareContext: 初始化⼀个新的纹理加载对象 复制代码`

* 从文件中加载处理

` + textureWithContentsOfFile:options:errer: 从⽂件加载2D纹理图像并从数据中 创建新的纹理 - textureWithContentsOfFile:options:queue:completionHandler: 从文件中异步 加载2D纹理图像,并根据数据创建新纹理 复制代码`

### GLTextureLoader 简化从各种资源⽂件中加载纹理. ###

* 从 ` URL` 加载纹理

` - textureWithContentsOfURL:options:error: 从URL加载2D纹理图像并从数据创建新纹理 - textureWithContentsOfURL:options:queue:completionHandler: 从URL异步加载2D纹理图像,并根据数据创建新纹理 复制代码`

* 从内存中表示创建纹理

` + textureWithContentsOfData:options:errer: 从内存空间加载2D纹理图像,并根据数据创建新纹理 - textureWithContentsOfData:options:queue:completionHandler:从内存空间异步加载2D纹理图像,并从数据中创建新纹理 复制代码`

* 从 ` CGImages` 创建纹理

` - textureWithCGImage:options:error: 从Quartz图像 加载2D纹理图像并从数据创建新纹理 - textureWithCGImage:options:queue:completionHandler: 从Quartz图像异步加载2D纹理图像,并根据数据创建新纹理 复制代码`

* 从 ` URL` 加载多维创建纹理

` + cabeMapWithContentsOfURL:options:errer: 从单个URL加载⽴方体贴图纹理图像,并根据数据创建新纹理 - cabeMapWithContentsOfURL:options:queue:completionHandler:从单个URL异步加载⽴方体贴图纹理图像,并根据数据创建新纹理 复制代码`

* 从文件加载多维数据创建纹理

` + cubeMapWithContentsOfFile:options:errer: 从单个文件加载⽴方体贴图纹理对象,并从数据中创建新纹理 - cubeMapWithContentsOfFile:options:queue:completionHandler:从单个文件异步加载⽴方体贴图纹理对象,并从数据中创建新纹理 + cubeMapWithContentsOfFiles:options:errer: 从⼀系列文件中加载⽴方体贴图纹理图像,并从数据总创建新纹理 - cubeMapWithContentsOfFiles:options:options:queue:completionHandler:从⼀系列⽂件异步加载⽴方体贴图纹理图像,并从数据中创建新纹理 复制代码`

### GLKView 使用OpenGL ES 绘制内容的视图默认实现 ###

* 初始化视图

` - initWithFrame:context: 初始化新视图 复制代码`

* 代理

` delegate 视图的代理 复制代码`

* 配置帧缓存区对象

` drawableColorFormat 颜⾊色渲染缓存区格式 drawableDepthFormat 深度渲染缓存区格式 drawableStencilFormat 模板渲染缓存区的格式 drawableMultisample 多重采样缓存区的格式 复制代码`

* 帧缓存区属性

` drawableHeight 底层缓存区对象的高度(以像素为单位) drawableWidth 底层缓存区对象的宽度(以像素为单位) 复制代码`

* 绘制视图的内容

` context 绘制视图内容时使用的OpenGL ES 上下⽂ - bind Drawable 将底层FrameBuffer 对象绑定到OpenGL ES enable SetNeedsDisplay 布尔值,指定视图是否响应使得视图内容无效的消息 - display ⽴即重绘视图内容 snapshot 绘制视图内容并将其作为新图像对象返回 复制代码`

* 删除视图FrameBuffer对象

` - deleteDrawable 删除与视图关联的可绘制对象 复制代码`

### GLKViewDelegate ⽤于GLKView 对象回调⽅法 ###

* 绘制视图的内容

` - glkView:drawInRect: 绘制视图内容 (必须实现代理) 复制代码`

### GLKViewController 管理OpenGL ES 渲染循环的视图控制器 ###

* 更新

` - (void) update 更新视图内容 - (void) glkViewControllerUpdate: 复制代码`

* 配置帧速率

` preferredFramesPerSecond 视图控制器调用视图以及更新视图内容的速率 framesPerSencond 视图控制器调⽤视图以及更新其内容的实际速率 复制代码`

* 配置 ` GLKViewController` 代理

` delegate 视图控制器的代理 复制代码`

* 控制帧更新

` paused 布尔值,渲染循环是否已暂停 pausedOnWillResignActive 布尔值,当前程序重新激活动状态时视图控制器是否⾃动暂停渲染循环 resumeOnDidBecomeActive 布尔值,当前程序变为活动状态时视图控制是否自动恢复呈现循环 复制代码`

* 获取有关 ` View` 更新信息

` frameDisplayed 视图控制器自创建以来发送的帧更新数 timeSinceFirstResume ⾃视图控制器第一次恢复发送更新事件以来经过的时间量 timeSinceLastResume 自上次视图控制器恢复发送更新事件以来更更新的时间量 timeSinceLastUpdate ⾃上次视图控制器调用委托方法以及经过的时间量 glkViewControllerUpdate: timeSinceLastDraw ⾃上次视图控制器调用视图display方法以来经过的时间量 复制代码`

### GLKViewControllerDelegate 渲染循环回调⽅法 ###

* 处理更新事件

` - glkViewControllerUpdate: 在显示每个帧之前调⽤ 复制代码`

* 暂停/恢复通知

` - glkViewController:willPause: 在渲染循环暂停或恢复之前调⽤ 复制代码`

### GLKBaseEffect 一种简单光照/着色系统,⽤于基于着色器OpenGL渲染 ###

* 命名 ` Effect`

` label 给Effect(效果)命名 复制代码`

* 配置模型视图转换

` transform 绑定效果时应⽤于顶点数据的模型视图,投影和纹理变换 复制代码`

* 配置光照效果

` lightingType ⽤于计算每个片段的光照策略,GLKLightingType GLKLightingType GLKLightingTypePerVertex 表示在三角形中每个顶点执行光照计算,然后在三角形进⾏插值 GLKLightingTypePerPixel 表示光照计算的输入在三角形内插入,并且在每个片段执⾏光照计算 复制代码`

* 配置光照

` lightModelTwoSided 布尔值,表示为基元的两侧计算光照 material 计算渲染图元光照使用的材质属性 lightModelAmbientColor 环境颜色,应⽤效果渲染的所有图元. light0 场景中第⼀个光照属性 light1 场景中第二个光照属性 light2 场景中第三个光照属性 复制代码`

* 配置纹理

` texture2d0 第一个纹理属性 texture2d1 第⼆个纹理属性 textureOrder 纹理应用于渲染图元的顺序 复制代码`

* 配置雾化

` fog 应⽤于场景的雾属性 复制代码`

* 配置颜色信息

` colorMaterialEnable 布尔值,表示计算光照与材质交互时是否使用颜⾊顶点属性 useConstantColor 布尔值,指示是否使⽤用常量颜色 constantColor 不提供每个顶点颜⾊数据时使⽤常量颜⾊ 复制代码`

* 准备绘制效果

` - prepareToDraw 准备渲染效果 复制代码`