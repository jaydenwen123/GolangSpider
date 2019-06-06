# OpenGL/OpenGL ES入门： GLKit使用以及案例 #

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

## 案例一： 图片渲染 ##

第一个案例，我们创建一个继承 ` GLKViewController` 的类，并导入 ` #import <GLKit/GLKit.h>` 头文件， `.h` 文件如下

` #import <UIKit/UIKit.h> #import <GLKit/GLKit.h> @interface ViewController : GLKViewController @end 复制代码`

`.m` 文件导入下面两个文件 ` #import <OpenGLES/ES3/gl.h>` 、 ` #import <OpenGLES/ES3/glext.h>`

下面看具体步骤：

### 初始化上下文&设置当前上下文 ###

` // 1.初始化上下文&设置当前上下文 _context = [[EAGLContext alloc]initWithAPI:kEAGLRenderingAPIOpenGLES3]; if (!_context) { NSLog(@ "Create ES context Failed" ); } // 设置当前上下文 [EAGLContext set CurrentContext:_context]; // 2.获取GLKView & 设置context GLKView *view = [[GLKView alloc]initWithFrame:self.view.bounds context:_context]; view.backgroundColor = [UIColor clearColor]; view.delegate = self; [self.view addSubview:view]; 复制代码`

在初始化 ` context` 时，我们需要选择我们使用的 ` OpenGL ES` 的版本，分别有如下几种：

` kEAGLRenderingAPIOpenGLES1 = 1, 固定管线 kEAGLRenderingAPIOpenGLES2 = 2, kEAGLRenderingAPIOpenGLES3 = 3, 复制代码`

` EAGLContext` 是苹果iOS平台下实现 ` OpenGL ES` 渲染层。其中 ` kEAGLRenderingAPIOpenGLES1` 类似前几篇文章说的固定管线。而 ` kEAGLRenderingAPIOpenGLES2` 和 ` kEAGLRenderingAPIOpenGLES3` ，在这里使用哪个，区别不大，只是版本不同而已。

下面需要配置视图创建的渲染缓冲区，以及设置背景颜色

` view.drawableColorFormat = GLKViewDrawableColorFormatRGBA8888; view.drawableDepthFormat = GLKViewDrawableDepthFormat24; glClearColor(1, 1, 1, 1); 复制代码`
> 
> 
> 
> 
> **` drawableColorFormat`** : 颜色缓冲区格式
> 在 ` OpenGL ES` 中有一个缓存区，它用以存储在屏幕中显示的颜色。你可以使用其属性设置缓冲区中的每个像素的颜色格式。
> ` GLKViewDrawableColorFormatRGBA8888` = 0
> 默认值，缓冲区的每个像素的最小组成部分(RGBA)使用8个bit，(所以每个像素4个字节， 4*8个bit)
> ` GLKViewDrawableColorFormatRGB565`
> 如果你的APP允许更小范围的颜色，即可设置这个。会让你的APP消耗更小的资源(内存和处理时间)
> 
> 

> 
> 
> 
> **` drawableDepthFormat`** : 深度缓存区格式
> 
> 
> 
> ` GLKViewDrawableDepthFormatNone` = 0, 意味着完全没有深度缓冲区
> ` GLKViewDrawableDepthFormat16` ,
> ` GLKViewDrawableDepthFormat24` ,
> 如果你要使用这个属性（一般用于3D游戏），你应该选择 ` GLKViewDrawableDepthFormat16` 或 `
> GLKViewDrawableDepthFormat24` 。这里的差别是使用 ` GLKViewDrawableDepthFormat16` 将消耗更少的资源
> 
> 
> 

### 加载纹理数据(使用 ` GLBaseEffect` ) ###

使用本地图片作为纹理，代码如下：

` //1.获取纹理图片路径 NSString *filePath = [[NSBundle mainBundle]pathForResource:@ "kunkun" ofType:@ "jpg" ]; //2.设置纹理参数 NSDictionary *options = [NSDictionary dictionaryWithObjectsAndKeys:@(1),GLKTextureLoaderOriginBottomLeft, nil]; GLKTextureInfo *textureInfo = [GLKTextureLoader textureWithContentsOfFile:filePath options:options error:nil]; //3.使用苹果GLKit 提供GLKBaseEffect 完成着色器工作(顶点/片元) cEffect = [[GLKBaseEffect alloc]init]; cEffect.texture2d0.enabled = GL_TRUE; cEffect.texture2d0.name = textureInfo.name; 复制代码`

在设置纹理参数的时，需要注意一点，纹理坐标默认左下角为原点(0,0)，而iOS坐标系中，默认手机屏幕的左上角为原点(0,0)，所以需要特别注意这一点，需要设置 ` GLKTextureLoaderOriginBottomLeft` ，不然图片显示的时候，会发生翻转。

### 加载顶点&纹理数据 ###

先看代码

` // 设置顶点数组 GL float vertexData[] = { 1, -0.5, 0.0f, 1.0f, 0.0f, //右下 1, 0.5, -0.0f, 1.0f, 1.0f, //右上 -1, 0.5, 0.0f, 0.0f, 1.0f, //左上 1, -0.5, 0.0f, 1.0f, 0.0f, //右下 -1, 0.5, 0.0f, 0.0f, 1.0f, //左上 -1, -0.5, 0.0f, 0.0f, 0.0f, //左下 }; // 开辟顶点缓存区 // 1.创建顶点缓存区标识符ID GLuint bufferID; glGenBuffers(1, &bufferID); // 2.绑定顶点缓存区.(明确作用) glBindBuffer(GL_ARRAY_BUFFER, bufferID); // 3.将顶点数组的数据copy到顶点缓存区中(GPU显存中) glBufferData(GL_ARRAY_BUFFER, sizeof(vertexData), vertexData, GL_STATIC_DRAW); // 顶点坐标数据 glEnableVertexAttribArray(GLKVertexAttribPosition); glVertexAttribPointer(GLKVertexAttribPosition, 3, GL_FLOAT, GL_FALSE, sizeof(GL float ) * 5, (GL float *)NULL + 0); // 纹理坐标数据 glEnableVertexAttribArray(GLKVertexAttribTexCoord0); glVertexAttribPointer(GLKVertexAttribTexCoord0, 2, GL_FLOAT, GL_FALSE, sizeof(GL float ) * 5, (GL float *)NULL + 3); 复制代码`

这个案例中，我们把顶点坐标、纹理坐标放在数组中，后面说的案例则使用了联合体的方式存放，主要是为了让大家认识不同方式而已。 也可以把顶点坐标和纹理坐标分开写到不同的数组中，但是为了简化代码量，所以选择放在一起写。

顶点数组，开发者可以选择设定函数指针，在调用绘制方法的时候，直接由内存传入顶点数据，也就是说这部分数据之前是存储在内存中的。而后面要做的是开辟顶点缓冲区，主要目的是为了追求性能更高，提前分配一块显存，将顶点数据预先传入到显存中，而关于顶点相关的计算都是在GPU中进行的，所以这样性能有了很大的提高。

> 
> 
> 
> **注意点：**
> 在iOS中，默认情况下，出于性能考虑，所有顶点着色器的属性( ` Attribute` )变量都是关闭的，意味着，顶点数据在着色器（服务端）是不可用的。即使你已经使用
> ` glBufferData` 方法，将顶点数据从内存拷贝到顶点缓冲区中(GPU显存中); 所以，必须由 `
> glEnableVertexAttribArray` 方法打开通道，指定访问属性，才能让顶点着色器能够访问到从CPU复制到GPU的数据。
> **注意：**
> 数据在GPU端是否可见，即着色器能否读取到数据，由是否启用了对应的属性决定，这就是 ` glEnableVertexAttribArray` 的功能，允许顶点着色器读取GPU（服务端）数据。
> 
> GLKVertexAttribPosition, // 顶点
> GLKVertexAttribNormal, // 法线
> GLKVertexAttribColor, // 颜色
> GLKVertexAttribTexCoord0, // 纹理
> GLKVertexAttribTexCoord1
> 
> 

方法 ` glVertexAttribPointer` 表示上传顶点数据到显存的方法，即设置合适的方式从 ` buffer` 里读取数据。

> 
> 
> 
> 参数1: 指定要修改的属性的索引值
> 参数2: 每次读取数量(如 ` position` 是由3个（x,y,z）组成，而颜色是4个（r,g,b,a）,纹理则是2个.)
> 参数3: 指定数组中每个组件的数据类型。可用的符号常量有 ` GL_BYTE` , ` GL_UNSIGNED_BYTE` , ` GL_SHORT`
> , ` GL_UNSIGNED_SHORT` , ` GL_FIXED` , 和 ` GL_FLOAT` ，初始值为 ` GL_FLOAT` 。
> 参数4: 指定当被访问时，固定点数据值是否应该被归一化（ ` GL_TRUE` ）或者直接转换为固定点值（ ` GL_FALSE` ）
> 参数5: 指定连续顶点属性之间的偏移量。如果为0，那么顶点属性会被理解为：它们是紧密排列在一起的。初始值为0
> 参数6: 指定一个指针，指向数组中第一个顶点属性的第一个组件。初始值为0
> 
> 

上述代码中，对于顶点来说，每次读取3个，数据类型为浮点型，偏移量即读取下一组顶点数据，需要移动5个位置，顶点读取位置对应从数组的一开始读取，所以最后一个参数写成 ` (GLfloat *)NULL + 0` ，这里这样的写法是为了方便读者理解，也可以直接写成 ` NULL` 。所以对于纹理的最后一个参数写成 ` (GLfloat *)NULL + 3` 也不难理解。

### 绘制视图的内容 ###

` GLKView` 对象使 ` OpenGL ES` 上下文成为当前上下文，并将其 ` framebuffer` 绑定为 ` OpenGL ES` 呈现命令的目标。然后，委托方法应该绘制视图的内容。看下面代码

` - (void)glkView:(GLKView *)view drawInRect:(CGRect)rect { //1.清除颜色缓存区 glClear(GL_COLOR_BUFFER_BIT); //2.准备绘制 [cEffect prepareToDraw]; //3.开始绘制 glDrawArrays(GL_TRIANGLES, 0, 6); } 复制代码`

OK，案例一就这么多内容，下面看实现效果（其实就是我们平时用 ` UIImageView` 加载一张图片而已）

![](https://user-gold-cdn.xitu.io/2019/6/4/16b23265cdd2a7b9?imageView2/0/w/1280/h/960/ignore-error/1)

## 案例二 绘制立方体 ##

案例二中，我们不再创建继承 ` GLKViewController` 的类，而是直接在 ` ViewController` 中写代码，直接在 `.m` 文件中导入 ` #import <GLKit/GLKit.h>` 头文件。

然后声明一些需要的属性：

` typedef struct { GLKVector3 positionCoord; // 顶点坐标 GLKVector2 textureCoord; // 纹理坐标 GLKVector3 normal; // 法线坐标 } ZBVertex; // 顶点个数 static NSInteger const KCoordCount = 36; @interface ViewController () <GLKViewDelegate> @property (nonatomic, strong) GLKView *glkView; @property (nonatomic, strong) GLKBaseEffect *baseEffect; @property (nonatomic, assign) ZBVertex *vertices; // 计时器 @property (nonatomic, strong) CADisplayLink *displayLink; // 弧度 @property (nonatomic, assign) NSInteger angle; // 顶点缓存区标识符ID @property (nonatomic, assign) GLuint vertexBuffer; @end 复制代码`

### 初始化上下文&设置当前上下文 ###

` // 1.创建context EAGLContext *context = [[EAGLContext alloc]initWithAPI:kEAGLRenderingAPIOpenGLES3]; [EAGLContext set CurrentContext:context]; // 2.创建GLKView并设置代理 CGRect frame = CGRectMake(0, 100, self.view.frame.size.width, self.view.frame.size.width); self.glkView = [[GLKView alloc]initWithFrame:frame context:context]; self.glkView.backgroundColor = [UIColor redColor]; self.glkView.delegate = self; // 3. 使用深度缓冲区 self.glkView.drawableDepthFormat = GLKViewDrawableDepthFormat24; self.glkView.drawableColorFormat = GLKViewDrawableColorFormatRGBA8888; // 默认为（0，1），这里用于翻转z轴，是正方形朝屏幕外 // glDepthRangef(1, 0); [self.view addSubview:self.glkView]; NSString *imagePath = [[[NSBundle mainBundle] resourcePath] stringByAppendingPathComponent:@ "timg.jpeg" ]; UIImage *image = [UIImage imageWithContentsOfFile:imagePath]; NSDictionary *options = @{GLKTextureLoaderOriginBottomLeft : @(YES)}; GLKTextureInfo *textureInfo = [GLKTextureLoader textureWithCGImage:[image CGImage] options:options error:nil]; // 使用苹果GLKit提供GLKBaseEffect完成着色器工作（顶点/片元） self.baseEffect = [[GLKBaseEffect alloc]init]; self.baseEffect.texture2d0.name = textureInfo.name; self.baseEffect.texture2d0.target = textureInfo.target; 复制代码`

上面初始化代码和案例一中的几乎是一样的，同时把加载纹理数据的代码也放在了这里，所以这里不再多说。

### 加载顶点&纹理数据 ###

` // 开辟顶点数据空间（数据结构SenceVertex 大小 * 顶点个数kCoordCount） self.vertices = malloc(sizeof(ZBVertex) * KCoordCount); // 前面 self.vertices[0] = (ZBVertex){{-0.5, 0.5, 0.5}, {0, 1}}; self.vertices[1] = (ZBVertex){{-0.5, -0.5, 0.5}, {0, 0}}; self.vertices[2] = (ZBVertex){{0.5, 0.5, 0.5}, {1, 1}}; self.vertices[3] = (ZBVertex){{-0.5, -0.5, 0.5}, {0, 0}}; self.vertices[4] = (ZBVertex){{0.5, 0.5, 0.5}, {1, 1}}; self.vertices[5] = (ZBVertex){{0.5, -0.5, 0.5}, {1, 0}}; // 上面 self.vertices[6] = (ZBVertex){{0.5, 0.5, 0.5}, {1, 1}}; self.vertices[7] = (ZBVertex){{-0.5, 0.5, 0.5}, {0, 1}}; self.vertices[8] = (ZBVertex){{0.5, 0.5, -0.5}, {1, 0}}; self.vertices[9] = (ZBVertex){{-0.5, 0.5, 0.5}, {0, 1}}; self.vertices[10] = (ZBVertex){{0.5, 0.5, -0.5}, {1, 0}}; self.vertices[11] = (ZBVertex){{-0.5, 0.5, -0.5}, {0, 0}}; // 下面 self.vertices[12] = (ZBVertex){{0.5, -0.5, 0.5}, {1, 1}}; self.vertices[13] = (ZBVertex){{-0.5, -0.5, 0.5}, {0, 1}}; self.vertices[14] = (ZBVertex){{0.5, -0.5, -0.5}, {1, 0}}; self.vertices[15] = (ZBVertex){{-0.5, -0.5, 0.5}, {0, 1}}; self.vertices[16] = (ZBVertex){{0.5, -0.5, -0.5}, {1, 0}}; self.vertices[17] = (ZBVertex){{-0.5, -0.5, -0.5}, {0, 0}}; // 左面 self.vertices[18] = (ZBVertex){{-0.5, 0.5, 0.5}, {1, 1}}; self.vertices[19] = (ZBVertex){{-0.5, -0.5, 0.5}, {0, 1}}; self.vertices[20] = (ZBVertex){{-0.5, 0.5, -0.5}, {1, 0}}; self.vertices[21] = (ZBVertex){{-0.5, -0.5, 0.5}, {0, 1}}; self.vertices[22] = (ZBVertex){{-0.5, 0.5, -0.5}, {1, 0}}; self.vertices[23] = (ZBVertex){{-0.5, -0.5, -0.5}, {0, 0}}; // 右面 self.vertices[24] = (ZBVertex){{0.5, 0.5, 0.5}, {1, 1}}; self.vertices[25] = (ZBVertex){{0.5, -0.5, 0.5}, {0, 1}}; self.vertices[26] = (ZBVertex){{0.5, 0.5, -0.5}, {1, 0}}; self.vertices[27] = (ZBVertex){{0.5, -0.5, 0.5}, {0, 1}}; self.vertices[28] = (ZBVertex){{0.5, 0.5, -0.5}, {1, 0}}; self.vertices[29] = (ZBVertex){{0.5, -0.5, -0.5}, {0, 0}}; // 后面 self.vertices[30] = (ZBVertex){{-0.5, 0.5, -0.5}, {0, 1}}; self.vertices[31] = (ZBVertex){{-0.5, -0.5, -0.5}, {0, 0}}; self.vertices[32] = (ZBVertex){{0.5, 0.5, -0.5}, {1, 1}}; self.vertices[33] = (ZBVertex){{-0.5, -0.5, -0.5}, {0, 0}}; self.vertices[34] = (ZBVertex){{0.5, 0.5, -0.5}, {1, 1}}; self.vertices[35] = (ZBVertex){{0.5, -0.5, -0.5}, {1, 0}}; // 开辟顶点缓存区 glGenBuffers(1, &_vertexBuffer); glBindBuffer(GL_ARRAY_BUFFER, _vertexBuffer); glBufferData(GL_ARRAY_BUFFER, sizeof(ZBVertex) * KCoordCount, self.vertices, GL_STATIC_DRAW); // 顶点数据 glEnableVertexAttribArray(GLKVertexAttribPosition); glVertexAttribPointer(GLKVertexAttribPosition, 3, GL_FLOAT, GL_FALSE, sizeof(ZBVertex), NULL + offsetof(ZBVertex, positionCoord)); // 纹理数据 glEnableVertexAttribArray(GLKVertexAttribTexCoord0); glVertexAttribPointer(GLKVertexAttribTexCoord0, 2, GL_FLOAT, GL_FALSE, sizeof(ZBVertex), NULL + offsetof(ZBVertex, textureCoord)); 复制代码`

咔咔咔，这么一大串，基本上都是在设置顶点坐标和纹理坐标，这里使用了联合体存放数据，一开始就开辟了相应的空间 ` self.vertices = malloc(sizeof(ZBVertex) * KCoordCount);` ，后面代码也和案例一一样。

### 创建一个循环 ###

` self.angle = 0; self.displayLink = [CADisplayLink displayLinkWithTarget:self selector:@selector(reDisplay)]; [self.displayLink addToRunLoop:[NSRunLoop mainRunLoop] for Mode:NSRunLoopCommonModes]; 复制代码`

上面代码创建了一个循环，来执行方法 ` reDisplay` 。 ` CADisplayLink` 类似定时器，提供一个周期性调用，属于 ` QuartzCore.framework` 中。
具体可以参考该博客 [www.cnblogs.com/panyangjun/…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fpanyangjun%2Fp%2F4421904.html )

` - (void)reDisplay { // 计算旋转度数 self.angle = (self.angle + 1) % 360; // 修改baseEffect.transform.modelviewMatrix self.baseEffect.transform.modelviewMatrix = GLKMatrix4MakeRotation(GLKMathDegreesToRadians(self.angle), 0.3, 1, -0.7); // 重新渲染 [self.glkView display]; } 复制代码`

### dealloc ###

` - (void)dealloc { if ([EAGLContext currentContext] == self.glkView.context) { [EAGLContext set CurrentContext:nil]; } if (_vertices) { free(_vertices); _vertexBuffer = 0; } //displayLink 失效 [self.displayLink invalidate]; } 复制代码`

最终效果：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2328b09b75a94?imageslim)