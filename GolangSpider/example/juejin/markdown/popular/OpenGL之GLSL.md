# OpenGL之GLSL #

` GLSL` 是OpenGL Shader的编程语言，为了更好的进行视频编辑和特效开发，抽空学习了其语法和特性，并做此记录，留作备忘查询。

## 类型 ##

### 类型概述 ###

+-------------------------+-------------------+---------------------------------------------------------------------------+
|        变量类别         |     变量类型      |                                   备注                                    |
+-------------------------+-------------------+---------------------------------------------------------------------------+
| 空                      | void              | 用于标识无参函数或者无返回值函数：void                                    |
|                         |                   | main(void);                                                               |
| 标量                    | bool/int/float    | 布尔类型、整型、浮点型                                                    |
| 布尔型向量              | bvec2/bvec3/bvec4 | 其中 ` b`                                                                 |
|                         |                   | 表示向量类型，数字表示向量的分量数                                        |
| 整型向量                | ivec2/ivec3/ivec4 | 其中 ` i`                                                                 |
|                         |                   | 表示向量类型，数字表示向量的分量数                                        |
| 浮点型向量              | vec2/vec3/vec4    | 默认情况下是 ` float`                                                     |
|                         |                   | 类型，数字表示向量的分量数                                                |
| 浮点型矩阵              | mat2/mat3/mat4    | 数字表示矩阵的列数，行数和列数相同                                        |
| 2D texture              | sampler2D         | 2D纹理，仅能作为uniform变量                                               |
| Cubemap(立方体) texture | samplerCube       | 立方体纹理，仅能作为uniform变量                                           |
| 结构体                  | struct            | 类似于C语言结构体，把多个变量聚合在一起                                   |
| 数组                    | array             | GLSL只支持1维数组，数据类型可以是标量类型、向量类型、矩阵类型、结构体类型 |
+-------------------------+-------------------+---------------------------------------------------------------------------+

向量构造器有两种参数构造方式：

* 单标量参数：向量中的所有分量都会初始化为该标量值。
* 多标量参数、向量参数、或者标量和向量混合参数：按照参数顺序初始化向量的所有分量，需要保证参数个数（向量参数的分量拆分开）不少于向量分量个数。
` vec4 aVec4 = vec4 ( 1.0 , 1.0 , 1.0 , 1.0 ); // 全部分量都为1.0 vec4 bVec4 = vec4 ( 1.0 ); vec3 aVec3 = vec3 ( 1.0 ); // 通过向量和标量初始化4维向量 vec4 cVec4 = vec4 (aVec3, 1.0 ); 复制代码`

类型转换： ` GLSL` 不支持隐式类型转换，必须进行显示类型转换，向量也支持直接类型转换。即只有类型一致时，变量才能完成赋值或其它操作。

` float floatParam = 1.0 ; // float -> bool bool boolParam = bool (floatParam); vec4 fVec4 = vec4 ( 1.0 , 1.0 , 1.0 , 1.0 ); // 4维float向量 -> 4维int向量 ivec4 iVec4 = ivec4 (fVec4); // 4维float向量 -> 3维int向量 ivec3 iVec3 = ivec3 (fVec4); 复制代码`

矩阵构造器也有两种参数构造方式：

* 单标量参数：用于主对角线上分量的初始化，其他分量皆为0.0。
* 多标量参数、向量参数、或者标量和向量混合参数：按照参数顺序初始化矩阵的所有分量（列优先），需要保证参数个数（向量参数的分量拆分开）不少于矩阵分量个数。
` // 通过多个标量为矩阵的各个分量赋值 mat3 aMat3 = mat3 ( 1.0 , 0.0 , 0.0 , // 第一列 0.0 , 1.0 , 0.0 , // 第二列 0.0 , 0.0 , 1.0 ); // 第三列 // 单个标量用于主对角线上分量的初始化，其他分量皆为0.0 mat3 bMat3 = mat3 ( 1.0 ); vec3 aVec3 = vec3 ( 1.0 , 0.0 , 0.0 ); vec3 bVec3 = vec3 ( 0.0 , 1.0 , 0.0 ); mat3 cMat3 = mat3 (aVec3, // 通过向量初始化第一列 bVec3, // 通过向量初始化第二列 0.0 , 0.0 , 1.0 ); // 通过多个标量初始化第三列 复制代码`

#### 向量和矩阵操作 ####

向量和矩阵都是由分量构成的，我们可以通过多种方式访问分量。 针对向量，可以通过 `.` 和数组下标访问，也可以通过 ` xyzw` 、 ` rgba` 或者 ` strq` 来访问，其中， ` xyzw` 通常用于位置相关向量， ` rgba` 通常用于颜色相关向量， ` strq` 通常用于纹理坐标相关向量。 ` x` 、 ` r` 、 ` s` 分别表示向量的第一个分量。上述三种方式不能混用，比如： ` xgr` 是不允许的。

针对矩阵，可以认为是向量的组合，一个mat3可以当做三个vec3，所以可以通过数组下标获取某一列的向量，然后再通过上述方法访问向量分量。

` // 向量访问 vec3 aVec3 = vec3 ( 0.0 , 1.0 , 2.0 ); vec3 temp; temp = aVec3.xyz; // temp = {0.0, 1.0, 2.0} temp = aVec3.xxx; // temp = {0.0, 0.0, 0.0} temp = aVec3.zyx; // temp = {2.0, 1.0, 0.0} // 矩阵访问 mat4 aMat4 = mat4 ( 1.0 ); vec4 column0 = aMat4[ 0 ]; // column0 = {1.0, 0.0, 0.0,0.0} float column1_row1 = aMat4[ 1 ][ 1 ]; // column1_row1 = 1.0 float column2_row2 = myMat4[ 2 ].z; // column2_row2 = 1.0 复制代码`

### 结构体 ###

与C语言类似， ` glsl` 支持定义结构体，同时会自动为结构体创建构造器（与结构体同名），用于生成结构体实例。

` // Leon是新的结构体类型，leon是结构体变量，是可选的。 struct Leon { vec4 aVec4; mat4 bMat4; float cF; } leon; // 通过结构体构造器初始化结构体参数，构造器的参数和结构体的元素必须有精确的对应关系 leon = Leon( vec4 ( 1.0 ), mat4 ( 1.0 ), 1.0 ); // 通过.操作符访问成员变量 vec4 dVec4 = leon.aVec4; mat4 eMat4 = leon.bMat4; 复制代码`

### 数组 ###

与C语言类似， ` glsl` 支持定义数组，数组类型可以是任意基本类型和结构体类型。 ` glsl` 有两点特殊之处：

* 除了 ` uniform` 变量之外，数组的索引必须是编译期常量。
* 数组的定义和初始化是分开的，即数组中的元素必须在数组定义之后逐个初始化，且数组不能使用const限定符。
` // 定义数组 float fArray[ 2 ]; vec2 vec2Array[ 2 ]; mat2 mat2Array[ 2 ]; Leon leonArray[ 2 ]; // 数组元素必须逐一被初始化 fArray[ 0 ] = 0.0 ; fArray[ 1 ] = 1.0 ; vec2Array[ 0 ] = vec2 ( 0.0 ); vec2Array[ 1 ] = vec2 ( 1.0 ); mat2Array[ 0 ] = mat2 ( 0.0 ); mat2Array[ 1 ] = mat2 ( 1.0 ); leonArray[ 0 ] = Leon( vec4 ( 1.0 ), mat4 ( 1.0 ), 1.0 ); leonArray[ 1 ] = Leon( vec4 ( 1.0 ), mat4 ( 1.0 ), 1.0 ); 复制代码`

## 语句 ##

### 流程控制语句 ###

流程控制语句与C语言非常相似。

#### 循环语句 ####

` GLSL` 支持通过 ` continue` 跳过单次循环， ` break` 退出整个循环。

` // 通过计数器控制的循环 for ( int i = 0 ; i <= 99 ; i++) { function(); } // 通过前置（循环）条件控制的循环 int i = 0 ; while (i <= 99 ) { function(); i++; } // 通过后置（循环）条件控制的循环 do { function(); i++; } while (i <= 99 ); 复制代码`

#### 条件语句 ####

与C语言类似

` if (condition){ functionA(); } else { functionB(); } 复制代码`

#### discard语句 ####

` discard` 关键字一般用于片元着色器中，用于丢弃当前片元，即：立即退出当前片元的着色程序，当前片元最终是默认色。

#### Start语句 ####

` // main函数是顶点和片元着色器入口，相当于C语言的主函数 void main( void ) { function(); } 复制代码`

### 运算符 ###

GLSL运算符与C语言类似，此处不再赘述，只说不同点：

* GLSL要求运算符两侧的变量必须具有相同的数据类型。
* 对于二目运算符（*，/，+，-），操作数必须为浮点型或者整型。
* 比较运算符只能用于标量，有专门的内置函数用于向量比较，后续内置函数会介绍。

## 限定符 ##

### 存储限定符 ###

声明变量时，可以使用存储限定符进行修饰，类似C语言中的说明符。

#### const ####

` const` 用于修饰编译时常量（声明时必须赋初值），或者只读的函数参数。结构体成员不能被声明为常量，但是结构体变量可以被声明为常量，并且需要在初始化时使用构造器初始化其值。

#### attribute ####

` attribute` 变量是应用程序传输给顶点着色器的逐顶点数据，即每个顶点都需要一份数据，通常用于表示顶点坐标、顶点颜色、纹理坐标等数据。 ` attribute` 变量在 ` shader` 中是只读的，不能在顶点着色器中进行修改。此外， ` attribute` 变量必须由执行着色器的应用程序传输数据进行赋值。

可使用的最大 ` attribute` 数量是有上限的，可以使用内置函数 ` glGetIntegerv` 来查询 ` GL_MAX_VERTEX_ATTRIBS` 。OpenGL ES 2.0至少支持8个 ` attribute` 变量。

#### uniform ####

` uniform` 可用在顶点和片元着色器中，用于修饰全局变量，即它们不限于某个顶点或者某个像素。通常用于表示变换矩阵、纹理等数据。 ` uniform` 变量在 ` shader` 中是只读的，不能在着色器中进行修改。此外， ` uniform` 变量必须由执行着色器的应用程序传输数据进行赋值。

着色器可以使用的 ` uniform` 个数是有限的，可以使用内置函数 ` glGetIntegerv` 来查询 ` GL_MAX_VERTEX_UNIFORM_VECTORS` 和 ` GL_MAX_FRAGMENT_UNIFORM_VECTORS` 。OpenGL ES 2.0至少支持128个顶点uniform以及16个片元uniform。

#### varying ####

` varying` 用于声明顶点着色器和片元着色器之间共享的变量。在顶点和片元着色器中，必须有相同的 ` varying` 变量声明。首先，顶点着色器为每个顶点的 ` varying` 变量赋值，然后光栅化阶段会对 ` varying` 变量进行插值计算（光栅化后的每个片元都会得到一个对应的 ` varying` 变量值），最后把插值结果交给片元着色器进行着色。 ` varying` 只能用于修饰浮点标量、浮点向量、浮点矩阵以及包含这些类型的数组。

varying数量也存在限制，可以使用 ` glGetIntegerv` 来查询 ` GL_MAX_VARYING_VECTORS` 。OpenGL ES 2.0至少支持8个varying变量。

> 
> 
> 
> 本地变量和函数参数只能使用 ` const` 限定符，函数返回值和结构体成员不能使用限定符。 ` attribute` 和 ` uniform` 变量不能在初始化时赋值，这些变量必须由应用程序主动进行赋值。
> 
> 
> 

### 参数限定符 ###

GLSL提供了参数限定符用于修饰函数参数，表示参数是否可读写。

#### in（只读参数） ####

` in` 是默认修饰符，表示参数是值传递，并且在函数内不能修改该值。

#### out（只写参数） ####

` out` 表示参数是引用传递，但是参数没有被初始化，所以在函数内不可读，只可写，并且函数结束后，修改依然生效。

#### inout（可读写参数） ####

` inout` 表示参数是引用传递，在函数内可以对参数进行读写。并且函数结束后，修改依然生效（类似于C++中的引用）。

` int newFunction( in bvec4 aBvec4, // read-only out vec3 aVec3, // write-only inout int aInt); // read-write 复制代码`

### 精度限定符 ###

精度限定符明确指定了着色器变量使用的精度范围，包括： ` lowp` 、 ` mediump` 和 ` highp` 三种精度。当使用低精度时，OpenGL ES可以更快速和低功耗地运行着色器，效率的提高得益于精度的舍弃，如果精度选择不合理，着色器运行的结果会出现失真。所以需要测试不同的精度才能确定最合适的配置。

+------------+------------------------+
| 精度限定符 |          描述          |
+------------+------------------------+
| highp      | 最高精度               |
| lowp       | 最低精度               |
| mediump    | 中间精度，介于两者之间 |
+------------+------------------------+

精度限定符可以修饰整型或者浮点型标量、向量和矩阵。在顶点着色器中，精度限定符是可选的，假如没有指定限定符，那么默认使用最高精度。在片元着色器中，精度限定符是必须的，要么在定义变量时直接指定精度，要么指定类型的默认精度。

` // 为类型指定精度，若定义float变量时，没有指定精度限定符，那么默认是highp精度 precision highp float ; precision mediump int ; // 为变量指定精度 highp vec4 position; varying lowp vec4 color; 复制代码`

在OpenGL ES中，精度的定义及范围取决于具体的实现，不是完全一致。精度限定符指定了存储这些变量时，必须满足的最小范围和精度。具体实现可能会使用比要求更大的范围和精度，但绝对不会比要求少。

不同着色器支持的不同精度限定符的范围与精度可以使下面的函数查询：

` void GL_APIENTRY glGetShaderPrecisionFormat (GLenum shadertype, GLenum precisiontype, GLint *range, GLint * precision ); 复制代码`

其中，shadertype必须是 ` VERTEX_SHADER` 或 ` FRAGMENT_SHADER` ；precisiontype必须是 ` GL_LOW_FLOAT` 、 ` GL_MEDIUM_FLOAT` 、 ` GL_HIGH_FLOAT` 、 ` GL_LOW_INT` 、 ` GL_MEDIUM_INT` 或 ` GL_HIGH_INT` ；range是指向长度为2的整数数组的指针，这两个整数表示数值范围：

` range[ 0 ] = log2 (| min |) range[ 1 ] = log2 (| max |) 复制代码`

precision是指向整数的指针，该整数是对应的精度位数（用 ` log2` 取对数值）。

> 
> 
> 
> 当表达式中出现多种精度时，选择精度最高的那个；当一个精度都没找到时，则使用默认或最大精度。
> 
> 

### 限定符的顺序 ###

当出现多种限定度时，需要遵循一定的顺序：

* 一般变量：存储限定符 -> 精度限定符
* 函数参数：存储限定符 -> 参数限定符 -> 精度限定符

` // 变量 varying highp float color; // storage -> precision void fun( const in lowp float size){ // storage -> parameter -> precision } 复制代码`

## 预处理 ##

GLSL中的预处理指令与C语言类似，都是以 ` #` 开头，每个指令独占一行。常见的预处理指令如下所示：

` // 定义和取消定义 #define #undef // 条件判断 #if #ifdef #ifndef #else #elif #endif #pragma #extension #version #line 复制代码`

` #pragma` 表示编译指示，用来控制编译器行为。

` #pragma debug(on) #pragma debug(off) 复制代码`

开发和调试时可以打开debug选项，以便获取更多的调试信息，默认为off。

` #version` 指定使用哪个GLSL版本编译着色器，必须写在编译单元的最前面，其前面只能有注释或空白，不能有其他字符。

` #extension` 用来启用某些扩展功能，每个显卡驱动厂商都可以定义自己的OpenGL扩展，如： ` GL_OES_EGL_image_external` 。启用扩展的命令如下所示：

` // 为某个扩展设置行为 #extension extension_name: behavior // 为所有扩展设置行为 #extension all : behavior // 启用OES纹理 #extension GL_OES_EGL_image_external : require 复制代码`

` extension_name` 表示具体某个扩展， ` all` 表示编译器支持的所有扩展， ` behavior` 表示对扩展的具体操作，比如启用、禁用等，如下所示：

+----------+------------------------------------------------------------------------------------------+
| BEHAVIOR |                                           描述                                           |
+----------+------------------------------------------------------------------------------------------+
| require  | 启用某扩展，如果不支持，则报错，如果扩展参数为                                           |
|          | ` all` ，则一定会出错                                                                    |
| enable   | 启用某扩展，如果不支持，则会警告，如果扩展参数为                                         |
|          | ` all` ，则一定会出错                                                                    |
| warn     | 除非因为该扩展被其它处于启用状态的扩展所需要，否则使用该扩展时会发出警告，如果扩展参数为 |
|          | ` all` ，则一定会抛出警告                                                                |
| disable  | 禁用某扩展，如果使用该扩展则会抛出错误，如果扩展参数为                                   |
|          | ` all` （默认设置），则不允许使用任何扩展                                                |
+----------+------------------------------------------------------------------------------------------+

对于每一个被支持的扩展，都有一个对应的宏定义，我们可以用它来判断编译器是否支持该扩展。

` #ifdef OES_extension_name #extension OES_extension_name : enable #else #endif 复制代码`

除此之外，还有一些预定义的系统变量：

` __LINE__ // int类型，当前源码中的行号. __FILE__ // int类型，当前Source String的唯一ID标识 __VERSION__ // int类型，当前的glsl版本，比如100（100 = v1.00） GL_ES // 如果当前是在OpenGL ES环境，为1，否则为0 复制代码`

## 内置变量 ##

OpenGL着色语言包含一些内置变量，这里主要介绍顶点和片元着色器的内置变量。

### 顶点着色器内置变量 ###

#### gl_Position ####

` gl_Position` 是输出变量，用来保存顶点位置的齐次坐标。该值用作图元装配、裁剪以及其他固定管道操作。如果顶点着色器中没有对 ` gl_Position` 赋值，那么在后续阶段它的值是不确定的。 ` gl_Position` 可以被写入多次，后续步骤以最后一次写入值为准。

#### gl_PointSize ####

` gl_PointSize` 是输出变量，是着色器用来控制被栅格化点的大小，以像素为单位。如果顶点着色器中没有对 ` gl_PointSize` 赋值，那么在后续阶段它的值是不确定的。

### 片元着色器内置变量 ###

#### gl_FragColor ####

` gl_FragColor` 是输出变量，定义了后续管线中片元的颜色值。 ` gl_FragColor` 可以被写入多次，后续步骤以最后一次写入值为准。如果执行了 ` discard` 操作，则片元会被丢弃， ` gl_FragColor` 将不再有意义。

#### gl_FragCoord ####

` gl_FragCoord` 是只读变量，保存当前片元的窗口坐标(x, y, z, 1/w)，该值是图元装配阶段对图元插值计算所得， ` z` 分量表示当前片元的深度值。

#### gl_FragDepth ####

` gl_FragDepth` 是输出变量，用于改变片元的深度值（替代 ` gl_FragCoord.z` 深度值）。如果在任何地方写入了它的值，那么务必在所有执行路径中都写入它的值，否则未写入的路径上它的值有可能不明确。

#### gl_FrontFacing ####

` gl_FrontFacing` 是只读变量，如果片元属于面朝前的图元，那么它的值为true。该变量可以选取顶点着色器计算出的两个颜色之一以模拟双面光照。

#### gl_PointCoord ####

` gl_PointCoord` 是只读变量，表示当前片元在点图元中的二维坐标，范围是0.0到1.0。如果当前图元不是点，那么 ` gl_PointCoord` 读取的值将是不明确的。

## 函数 ##

### 自定义函数 ###

` GLSL` 支持自定义函数，但是函数不能嵌套，不支持递归调用，必须声明函数返回值类型（无返回值时声明为void）。如果一个函数在定义前被调用，则需要先声明其函数原型。

` vec4 getPosition(){ vec4 v4 = vec4 ( 0. , 0. , 0. , 1. ); return v4; } 复制代码`

### 内嵌函数 ###

#### 三角形函数 ####

` ParamType` 可以是 ` float` 、 ` vec2` 、 ` vec3` 、 ` vec4` ，参数类型和返回类型是一致的。

+--------------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------+
|            函数原型            |                                                                        描述                                                                        |
+--------------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------+
| ParamType radians(ParamType    | 把角度转换为弧度                                                                                                                                   |
| degrees)                       |                                                                                                                                                    |
| ParamType degrees(ParamType    | 把弧度转换为角度                                                                                                                                   |
| radians)                       |                                                                                                                                                    |
| ParamType sin(ParamType        | 以弧度为单位，计算正弦值                                                                                                                           |
| radians)                       |                                                                                                                                                    |
| ParamType cos(ParamType        | 以弧度为单位，计算余弦值                                                                                                                           |
| radians)                       |                                                                                                                                                    |
| ParamType tan(ParamType        | 以弧度为单位，计算正切值                                                                                                                           |
| radians)                       |                                                                                                                                                    |
| ParamType asin(ParamType       | sin的反函数，计算给定值的弧度，value的绝对值                                                                                                       |
| value)                         | <= 1，返回的弧度范围是[-π/2,π/2]                                                                                                                 |
| ParamType acos(ParamType       | cos的反函数，计算给定值的弧度，value的绝对值                                                                                                       |
| value)                         | <= 1，返回的弧度范围是[0,π]                                                                                                                       |
| ParamType atan(ParamType       | tan的反函数，计算给定值的弧度，返回的弧度范围是[-π/2,π/2]                                                                                        |
| y_over_x)                      |                                                                                                                                                    |
| ParamType atan(ParamType y,    | tan的反函数，也可以称作atan2，返回一个正切值为y/x的弧度，x和y的符号用来确定角在哪个象限。返回的弧度范围是(−π,π)。如果x和y都为0，则结果是未定义的 |
| ParamType x)                   |                                                                                                                                                    |
+--------------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------+

#### 指数函数 ####

` ParamType` 可以是 ` float` 、 ` vec2` 、 ` vec3` 、 ` vec4` ，参数类型和返回类型是一致的。

+--------------------------------+--------------------------+
|            函数原型            |           描述           |
+--------------------------------+--------------------------+
| ParamType pow(ParamType x,     | 幂函数返回x的y次方       |
| ParamType y)                   |                          |
| ParamType exp(ParamType x)     | exp函数返回常数e的x次方  |
| ParamType exp2(ParamType x)    | exp2函数返回常数2的x次方 |
| ParamType log(ParamType x)     | 以常数e为底x的对数函数   |
| ParamType log2(ParamType x)    | 以常数2为底x的对数函数   |
| ParamType sqrt(ParamType x)    | 返回X的平方根            |
| ParamType                      | 返回x的平方根的倒数      |
| inversesqrt(ParamType x)       |                          |
+--------------------------------+--------------------------+

#### 通用函数 ####

` ParamType` 可以是 ` float` 、 ` vec2` 、 ` vec3` 、 ` vec4` ，参数类型和返回类型是一致的。

+--------------------------------+-------------------------------------------------------------------------------------------------------------+
|            函数原型            |                                                    描述                                                     |
+--------------------------------+-------------------------------------------------------------------------------------------------------------+
| ParamType abs(ParamType x)     | 返回x的绝对值                                                                                               |
| ParamType sign(ParamType x)    | x为正时返回1.0,x为零时返回0.0,x为负时返回-1.0                                                               |
| ParamType floor(ParamType x)   | 返回小于或等于x的最大整数                                                                                   |
| ParamType ceil(ParamType x)    | 返回大于或等于x的最小值                                                                                     |
| ParamType fract(ParamType x)   | 返回x的小数部分，即: x -                                                                                    |
|                                | floor(x)。                                                                                                  |
| ParamType mod(ParamType x,     | x – y * floor(x /                                                                                          |
| ParamType y)                   | y)，如果x和y是整数，返回值是x除以y的余数                                                                    |
| ParamType mod(ParamType x,     | 上面mod的变体：y总是float，对于浮点向量，向量的每个分量除以y                                                |
| float y)                       |                                                                                                             |
| ParamType min(ParamType x,     | 返回两个参数的较小值。对于浮点向量，操作是按分量进行比较的                                                  |
| ParamType y)                   |                                                                                                             |
| ParamType min(ParamType x,     | 上面min的变体：y总是float，对于浮点向量，向量的每个分量与y进行比较                                          |
| float y)                       |                                                                                                             |
| ParamType max(ParamType x,     | 返回两个参数的较大值。对于浮点向量，操作是按分量进行比较的                                                  |
| ParamType y)                   |                                                                                                             |
| ParamType max(ParamType x,     | 上面max的变体：y总是float，对于浮点向量，向量的每个分量与y进行比较                                          |
| float y)                       |                                                                                                             |
| ParamType clamp(ParamType x,   | 如果x大于minVal，小于maxVal，则返回x。如果x小于minVal，则返回minVal。如果x大于maxVal，则返回maxVal          |
| ParamType minVal, ParamType    |                                                                                                             |
| maxVal)                        |                                                                                                             |
| ParamType clamp(ParamType x,   | 上面clamp的变体，minVal和maxVal总是float，对于浮点向量，向量的每个分量与minVal和maxVal进行比较              |
| float minVal, float maxVal)    |                                                                                                             |
| ParamType mix(ParamType x,     | 返回x和y的线性混合，即                                                                                      |
| ParamType y, ParamType a)      | ` x * (1 - a) + y * a`                                                                                      |
| ParamType mix(ParamType x,     | 上面mix的变体，a总是float，对于浮点向量，向量的每个分量与a计算                                              |
| ParamType y, float a)          |                                                                                                             |
| ParamType step(ParamType edge, | 如果x比edge小，则返回0.0，否则返回1.0                                                                       |
| ParamType x)                   |                                                                                                             |
| ParamType step(float edge,     | 上面step的变体，edge总是float，对于浮点向量，向量的每个分量与edge计算                                       |
| ParamType x)                   |                                                                                                             |
| ParamType smoothstep(ParamType | 如果x小于edge0，则返回0.0；如果x大于edge1，则返回1.0。否则，返回值将使用Hermite多项式在0.0到1.0之间进行插值 |
| edge0, ParamType edge1,        |                                                                                                             |
| ParamType x)                   |                                                                                                             |
| ParamType smoothstep(float     | 上面smoothstep的变体，edge0和edge1总是float，对于浮点向量，向量的每个分量与edge0和edge1比较                 |
| edge0, float edge1, ParamType  |                                                                                                             |
| x)                             |                                                                                                             |
+--------------------------------+-------------------------------------------------------------------------------------------------------------+

#### 几何函数 ####

` ParamType` 可以是 ` float` 、 ` vec2` 、 ` vec3` 、 ` vec4` 。

+--------------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------+
|            函数原型            |                                                                        描述                                                                        |
+--------------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------+
| float length(ParamType x)      | 返回向量长度，即各分量平方和的平方根，对于浮点标量则返回其绝对值                                                                                   |
| float distance(ParamType p0,   | 返回两点之间的距离，两点的距离就是向量d(p0 - p1,                                                                                                   |
| ParamType p1)                  | 从p1开始，指向p0)的长度，对于浮点标量则返回其绝对值                                                                                                |
| float distance(ParamType p0,   | 返回两点之间的距离，两点的距离就是向量d(p0 - p1,                                                                                                   |
| ParamType p1)                  | 从p1开始，指向p0)的长度，对于浮点标量则返回其绝对值                                                                                                |
| float dot(ParamType x,         | 返回两个向量的点积，即两个向量各个分量乘积的和，对于浮点标量则返回x和y的乘积                                                                       |
| ParamType y)                   |                                                                                                                                                    |
| vec3 cross(vec3 x, vec3 y)     | 返回两个向量的叉乘向量，更为熟知的叫法是法向量，该向量垂直于x和y向量构成的平面，其长度等于x和y构成的平行四边形的面积。输入参数只能是三分量浮点向量 |
| ParamType normalize(ParamType  | 返回向量x的单位向量，即                                                                                                                            |
| x)                             | ` x / \|x\|`                                                                                                                                       |
|                                | ，对于浮点标量直接返回1.0                                                                                                                          |
| ParamType                      | 如果dot(Nref, I) <                                                                                                                                 |
| faceforward(ParamType N,       | 0那么返回N向量，否则返回–N向量                                                                                                                    |
| ParamType I, ParamType Nref)   |                                                                                                                                                    |
| ParamType reflect(ParamType I, | I：the incident vector(入射向量),                                                                                                                  |
| ParamType N)                   | N：the normal vector of the reflecting                                                                                                             |
|                                | surface(反射面的法向量)，返回反射方向:                                                                                                             |
|                                | I – 2 * dot(N,I) * N，N一般是单位向量                                                                                                             |
| ParamType refract(ParamType I, | I：the incident vector(入射向量), N：the normal vector of the refracting                                                                           |
| ParamType N, float eta)        | surface(折射面法向量)，eta：折射率。返回入射向量I关于法向量N的折射向量，折射率为eta，I和N一般是单位向量                                            |
+--------------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------+

#### 矩阵函数 ####

不同于传统的向量相乘，这里 ` matrixCompMult` 返回一个向量，该向量的各分量等于x和y向量各个分量的乘积。

` mat2 matrixCompMult ( mat2 x, mat2 y) mat3 matrixCompMult ( mat3 x, mat3 y) mat4 matrixCompMult ( mat4 x, mat4 y) // 结果向量z z[i][j] = x[i][j] * y[i][j] 复制代码`

#### 矢量（向量）关系函数 ####

矢量（向量）关系函数主要包含：<, <=, >, >=, ==, !=），函数参数是整型或者浮点型向量，函数返回值是参数向量各个分量比较产生的一个布尔型向量。 下面用 ` bvec` 表示bvec2、bvec3、bvec4； ` ivec` 表示ivec2、ivec3、ivec4； ` vec` 表示 vec2、vec3、vec4。参数向量和返回值向量的大小必须一致。

+--------------------------------+------------------------------------------------+
|            函数原型            |                      描述                      |
+--------------------------------+------------------------------------------------+
| bvec lessThan(vec x, vec y)    | 结果向量result[i] = x[i] <                     |
| bvec lessThan(ivec x, ivec y)  | y[i]                                           |
| bvec lessThanEqual(vec x, vec  | 结果向量result[i] = x[i] <=                    |
| y)  bvec lessThanEqual(ivec x, | y[i]                                           |
| ivec y)                        |                                                |
| bvec greaterThan(vec x, vec y) | 结果向量result[i] = x[i] >                     |
|  bvec greaterThan(ivec x, ivec | y[i]                                           |
| y)                             |                                                |
| bvec greaterThanEqual(vec      | 结果向量result[i] = x[i] >=                    |
| x, vec y)  bvec                | y[i]                                           |
| greaterThanEqual(ivec x, ivec  |                                                |
| y)                             |                                                |
| bvec equal(vec x, vec y)       | 结果向量result[i] = x[i] ==                    |
| bvec equal(ivec x, ivec y)     | y[i]                                           |
| bvec notEqual(vec x, vec y)    | 结果向量result[i] = x[i] !=                    |
| bvec notEqual(ivec x, ivec y)  | y[i]                                           |
| bool any(bvec x)               | 假如参数向量的任意一个分量为true，那么返回true |
| bool all(bvec x)               | 假如参数向量的所有分量为true，那么返回true     |
| bvec not(bvec x)               | 对布尔向量取反                                 |
+--------------------------------+------------------------------------------------+

#### 纹理查找函数 ####

纹理查询的目的是从纹理中提取指定坐标的颜色信息。OpenGL中纹理有两种：

* 2D纹理（sampler2D）
* 3D纹理（samplerCube）

顶点着色器和片元着色器中都可以使用纹理查找函数。但是顶点着色器中不会计算细节级别（level of detail），所以二者的纹理查找函数略有不同。

以下函数只在顶点着色器中可用：

` vec4 texture2DLod ( sampler2D sampler, vec2 coord, float lod); vec4 texture2DProjLod ( sampler2D sampler, vec3 coord, float lod); vec4 texture2DProjLod ( sampler2D sampler, vec4 coord, float lod); vec4 textureCubeLod ( samplerCube sampler, vec3 coord, float lod); 复制代码`

以下函数只在片元着色器中可用：

` vec4 texture2D ( sampler2D sampler, vec2 coord, float bias); vec4 texture2DProj ( sampler2D sampler, vec3 coord, float bias); vec4 texture2DProj ( sampler2D sampler, vec4 coord, float bias); vec4 textureCube ( samplerCube sampler, vec3 coord, float bias); 复制代码`

float类型参数bias表示：使用mipmaps计算纹理的适当细节级别之后，在执行实际的纹理查找操作之前添加的偏差。

以下函数在顶点着色器和片元着色器中都可用:

` vec4 texture2D ( sampler2D sampler, vec2 coord); vec4 texture2DProj ( sampler2D sampler, vec3 coord); vec4 texture2DProj ( sampler2D sampler, vec4 coord); vec4 textureCube ( samplerCube sampler, vec3 coord); 复制代码`