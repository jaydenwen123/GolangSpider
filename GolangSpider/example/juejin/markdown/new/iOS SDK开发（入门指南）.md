# iOS SDK开发（入门指南） #

#### 什么是SDK开发? ####

日常开发中，我们会遇到某些情况不能提供源码，项目组件化等需求，这时候我们就可以使用SDK开发，在OC的开发中，我们涉及到的一般是静态库(.a)或者动态库(.framework)。 ` （注:不是所有的.framework就一定是动态库）`

#### 静态库与动态库的区别？ ####

静态库：链接时完整地拷贝至可执行文件中，被多次使用就有多份冗余拷贝。表现形式为 `.a和.framework` 动态库：链接时不复制，程序运行时由系统动态加载到内存，供程序调用，系统只加载一次，多个程序共用，节省内存。 表现形式为 `.dylib和.framework` ` 注意：动态库只能苹果使用，如果项目中使用了动态库不允许上架(如：jspatch)`

#### a与.framework有什么区别？ ####

.a是一个纯二进制文件，.framework中除了有二进制文件之外还有资源文件。 .a文件不能直接使用，至少要有.h文件配合，.framework文件可以直接使用。 .a + .h + sourceFile = .framework。 建议用.framework.

#### 接下来将以实例帮助大家创建一个自己的 `.framework` ####

首先我们先创建一个 `.workspace`

![workspace](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64aa1deb5fe?imageView2/0/w/1280/h/960/ignore-error/1) 创建完毕后，再创建一个 `.frmawork` ![创建frmawork](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64aa2d3fb44?imageView2/0/w/1280/h/960/ignore-error/1) 将创建好的 ` frmawork` 加入到 ` workspace` ![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64aa3b8ed0e?imageView2/0/w/1280/h/960/ignore-error/1) 在 ` framewrok` 中可以封装入自己需要封装的内容 ![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64aa3e3d9d1?imageView2/0/w/1280/h/960/ignore-error/1) eg: 我在 ` StringUtils` 中加入了一个测试方法

` #import "StringUtils.h" @implementation StringUtils + (NSString *) test String:(NSString *)string { return [@ "MQTestFramework: " stringByAppendingString:string]; } @end 复制代码`

#### 接下来进行项目配置： ####

1、设置Build Setting参数 将Build Active Architecture only设置为NO

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64ab026e792?imageView2/0/w/1280/h/960/ignore-error/1)

2、设置Build Setting参数 Mach-O Type 为Static Library （配置静态、动态）

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64aa43c1846?imageView2/0/w/1280/h/960/ignore-error/1)

3、设置Build Setting参数 在Architectures下增加armv7s

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64ae01b2378?imageView2/0/w/1280/h/960/ignore-error/1)

4、在Build Phases中设置需要公开和需要隐藏的头文件

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64ae7ff37b9?imageView2/0/w/1280/h/960/ignore-error/1)

5、将头文件引入到 ` MQTestFramwork` (自己SDK的头文件)

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64aeb810351?imageView2/0/w/1280/h/960/ignore-error/1) 6、 ` Command + B` 运行项目，在 ` Product` 中找到 ` framework` ![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64af00d2cb7?imageView2/0/w/1280/h/960/ignore-error/1) ####framework使用 将封装好的 `.framework` 拉入需要使用的项目中 ![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64af83ee955?imageView2/0/w/1280/h/960/ignore-error/1) 使用封装好的功能 ![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64b22ed36e2?imageView2/0/w/1280/h/960/ignore-error/1) 运行： ![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b64b2353c908?imageView2/0/w/1280/h/960/ignore-error/1)