# 基于Transform实现更高效的组件化路由框架 #

### 前言 ###

之前通过APT实现了一个 [简易版ARouter框架]( https://juejin.im/post/5cecce216fb9a07f04202904 ) ，碰到的问题是APT在每个module的上下文是不同的，导致需要通过不同的文件来保存映射关系表。因为类文件的不确定，就需要初始化时在dex文件中扫描到指定目录下的class，然后通过反射初始化加载路由关系映射。阿里的做法是直接开启一个异步线程，创建DexFile对象加载dex。这多少会带来一些性能损耗，为了避免这些，我们通过Transform api实现另一种更加高效的路由框架。

### 思路 ###

gradle transform api可以用于android在构建过程的class文件转成dex文件之前，通过自定义插件，进行class字节码处理。有了这个api，我们就可以在apk构建过程找到所有注解标记的class类，然后操作字节码将这些映射关系写到同一个class中。

### 自定义插件 ###

首先我们需要自定义一个gradle插件，在application的模块中使用它。为了能够方便调试，我们取消上传插件环节，直接新建一个名称为buildSrc的library。 删除src/main下的所有文件，build.gradle配置中引入transform api和 [javassist]( https://link.juejin.im?target=http%3A%2F%2Fwww.javassist.org%2F ) （比asm更简便的字节码操作库）

` apply plugin: 'groovy' dependencies { implementation 'com.android.tools.build:gradle:3.1.2' compile 'com.android.tools.build:transform-api:1.5.0' compile 'org.javassist:javassist:3.20.0-GA' compile gradleApi() compile localGroovy() } 复制代码`

然后在src/main下创建groovy文件夹，在此文件夹下创建自己的包，然后新建RouterPlugin.groovy的文件

` package io.github.iamyours import org.gradle.api.Plugin import org.gradle.api.Project class RouterPlugin implements Plugin < Project > { @Override void apply(Project project) { println "=========自定义路由插件=========" } } 复制代码`

然后src下创建resources/META-INF/gradle-plugins目录，在此目录新建一个xxx.properties文件，文件名xxx就表示使用插件时的名称（apply plugin 'xxx'）,里面是具体插件的实现类

` implementation-class=io.github.iamyours.RouterPlugin 复制代码`

整个buildSrc目录如下图

![buildSrc目录](https://user-gold-cdn.xitu.io/2019/6/2/16b169f05a26a827?imageView2/0/w/1280/h/960/ignore-error/1) 然后我们在app下的build.gradle引入插件

` apply plugin: 'RouterPlugin' 复制代码`

然后make app,得到如下结果表明配置成功。

![image.png](https://user-gold-cdn.xitu.io/2019/6/2/16b169f058cb5e6d?imageView2/0/w/1280/h/960/ignore-error/1)

### router-api ###

在使用Transform api之前，创建一个router-api的java module处理路由逻辑。

` ## build.gradle apply plugin: 'java-library' dependencies { implementation fileTree (dir: 'libs' , include : [ '*.jar' ]) compileOnly 'com.google.android:android:4.1.1.4' } sourceCompatibility = "1.7" targetCompatibility = "1.7" 复制代码`

注解类@Route

` @Target ({ElementType.TYPE}) @Retention (RetentionPolicy.CLASS) public @interface Route { String path () ; } 复制代码`

映射类（后面通过插件修改这个class）

` public class RouteMap { void loadInto (Map<String,String> map) { throw new RuntimeException( "加载Router映射错误！" ); } } 复制代码`

ARouter（取名这个是为了方便重构）

` public class ARouter { private static final ARouter instance = new ARouter(); private Map<String, String> routeMap = new HashMap<>(); private ARouter () { } public static ARouter getInstance () { return instance; } public void init () { new RouteMap().loadInto(routeMap); } 复制代码`

因为RouteMap是确定的，直接new创建导入映射，后面只需要修改字节码，替换loadInto方法体即可，如：

` public class RouteMap { void loadInto (Map<String,String> map) { map.put( "/test/test" , "com.xxx.TestActivity" ); map.put( "/test/test2" , "com.xxx.Test2Activity" ); } } 复制代码`

#### RouteTransform ####

新建一个RouteTransform继承自Transform处理class文件，在自定义插件中注册它。

` class RouterPlugin implements Plugin < Project > { @Override void apply(Project project) { project.android.registerTransform( new RouterTransform(project)) } } 复制代码`

在RouteTransform的transform方法中我们遍历一下jar和class，为了测试模块化路由，新建一个news模块，引入library，并且把它加入到app模块。在news模块中，新建一个activity如：

` @Route(path = "/news/news_list" ) class NewsListActivity : AppCompatActivity () {} 复制代码`

然后在通过transform方法中遍历一下jar和class

` @Override void transform(TransformInvocation transformInvocation) throws TransformException, InterruptedException, IOException { def inputs = transformInvocation.inputs for (TransformInput input : inputs) { for (DirectoryInput dirInput : input.directoryInputs) { println( "dir:" +dirInput) } for (JarInput jarInput : input.jarInputs) { println( "jarInput:" +jarInput) } } } 复制代码`

可以得到如下信息

![image.png](https://user-gold-cdn.xitu.io/2019/6/2/16b169f05a331d9b?imageView2/0/w/1280/h/960/ignore-error/1) 通过日志，我们可以得到以下信息：

* app生成的class在directoryInputs下，有两个目录一个是java，一个是kotlin的。
* news和router-api模块的class在jarInputs下，且scopes=SUB_PROJECTS下，是一个jar包
* 其他第三发依赖在EXTERNAL_LIBRARIES下，也是通过jar形式，name和implementation依赖的名称相同。 知道这些信息，遍历查找Route注解生命的activity以及修改RouteMap范围就确定了。我们在directoryInputs中目录中遍历查找app模块的activity，在jarInputs下scopes为SUB_PROJECTS中查找其他模块的activity，然后在name为router-api的jar上修改RouteMap的字节码。

### ASM字节码读取 ###

有了class目录，就可以动手操作字节码了。主要有两种方式，ASM、javassist。两个都可以实现读写操作。ASM是基于指令级别的，性能更好更快，但是写入时你需要知道java虚拟机的一些指令，门槛较高。而javassist操作更佳简便，可以通过字符串写代码，然后转换成对应的字节码。考虑到性能，读取时用ASM，修改RouteMap时用javassist。

##### 读取目录中的class #####

` //从目录中读取class void readClassWithPath(File dir) { def root = dir.absolutePath dir.eachFileRecurse { File file -> def filePath = file.absolutePath if (!filePath.endsWith( ".class" )) return def className = getClassName(root, filePath) addRouteMap(filePath, className) } } /** * 从class中获取Route注解信息 * @param filePath */ void addRouteMap(String filePath, String className) { addRouteMap( new FileInputStream( new File(filePath)), className) } static final ANNOTATION_DESC = "Lio/github/iamyours/router/annotation/Route;" void addRouteMap(InputStream is, String className) { ClassReader reader = new ClassReader(is) ClassNode node = new ClassNode() reader.accept(node, 1 ) def list = node.invisibleAnnotations for (AnnotationNode an : list) { if (ANNOTATION_DESC == an.desc) { def path = an.values[ 1 ] routeMap[path] = className break } } } //获取类名 String getClassName(String root, String classPath) { return classPath.substring(root.length() + 1 , classPath.length() - 6 ) .replaceAll( "/" , "." ) } 复制代码`

通过ASM的ClassReader对象，可以读取一个class的相关信息，包括类信息，注解信息。以下是我通过idea debug得到的ASM相关信息

![ASM读取注解](https://user-gold-cdn.xitu.io/2019/6/2/16b169f05ae15579?imageView2/0/w/1280/h/960/ignore-error/1)

##### 从jar包中读取class #####

读取jar中的class，就需要通过java.util中的JarFile解压读取jar文件，遍历每个JarEntry。

` //从jar中读取class void readClassWithJar(JarInput jarInput) { JarFile jarFile = new JarFile(jarInput.file) Enumeration<JarEntry> enumeration = jarFile.entries() while (enumeration.hasMoreElements()) { JarEntry entry = enumeration.nextElement() String entryName = entry.getName() if (!entryName.endsWith( ".class" )) continue String className = entryName.substring( 0 , entryName.length() - 6 ).replaceAll( "/" , "." ) InputStream is = jarFile.getInputStream(entry) addRouteMap(is, className) } } 复制代码`

至此，我们遍历读取，保存Route注解标记的所有class，在transform最后我们打印routemap,重新make app。

![routeMap信息](https://user-gold-cdn.xitu.io/2019/6/2/16b169f05b1df3ca?imageView2/0/w/1280/h/960/ignore-error/1)

### Javassist修改RouteMap ###

所有的路由信息我们已经通过ASM读取保存了，接下来只要操作RouteMap的字节码，将这些信息保存到loadInto方法中就行了。RouteMap的class文件在route-api下的jar包中，我们通过遍历找到它

` static final ROUTE_NAME = "router-api:" @Override void transform(TransformInvocation transformInvocation) throws TransformException, InterruptedException, IOException { def inputs = transformInvocation.inputs def routeJarInput for (TransformInput input : inputs) { ... for (JarInput jarInput : input.jarInputs) { if (jarInput.name.startsWith(ROUTE_NAME)) { routeJarInput = jarInput } } } insertCodeIntoJar(routeJarInput, transformInvocation.outputProvider) ... } 复制代码`

这里我们新建一个临时文件，拷贝每一项，修改RouteMap，最后覆盖原先的jar。

` /** * 插入代码 * @param jarFile */ void insertCodeIntoJar(JarInput jarInput, TransformOutputProvider out) { File jarFile = jarInput.file def tmp = new File(jarFile.getParent(), jarFile.name + ".tmp" ) if (tmp.exists()) tmp.delete() def file = new JarFile(jarFile) def dest = getDestFile(jarInput, out) Enumeration enumeration = file.entries() JarOutputStream jos = new JarOutputStream(new FileOutputStream(tmp)) while (enumeration.hasMoreElements()) { JarEntry jarEntry = enumeration.nextElement() String entryName = jarEntry.name ZipEntry zipEntry = new ZipEntry(entryName) InputStream is = file.getInputStream(jarEntry) jos.putNextEntry(zipEntry) if (isRouteMapClass(entryName)) { jos.write(hackRouteMap(jarFile)) } else { jos.write(IOUtils.toByteArray(is)) } is.close() jos.closeEntry() } jos.close() file.close() if (jarFile.exists()) jarFile.delete() tmp.renameTo(jarFile) } 复制代码`

具体修改RouteMap的逻辑如下

` private static final String ROUTE_MAP_CLASS_NAME = "io.github.iamyours.router.RouteMap" private static final String ROUTE_MAP_CLASS_FILE_NAME = ROUTE_MAP_CLASS_NAME.replaceAll( "\\." , "/" ) + ".class" private byte [] hackRouteMap(File jarFile) { ClassPool pool = ClassPool.getDefault() pool.insertClassPath(jarFile.absolutePath) CtClass ctClass = pool.get(ROUTE_MAP_CLASS_NAME) CtMethod method = ctClass.getDeclaredMethod( "loadInto" ) StringBuffer code = new StringBuffer( "{" ) for (String key : routeMap.keySet()) { String value = routeMap[key] code.append( "\$1.put(\"" + key + "\",\"" + value + "\");" ) } code.append( "}" ) method.setBody(code.toString()) byte [] bytes = ctClass.toBytecode() ctClass.stopPruning( true ) ctClass.defrost() return bytes } 复制代码`

重新make app,然后使用JD-GUI打开jar包,可以看到RouteMap已经修改。

![RouteMap反编译信息](https://user-gold-cdn.xitu.io/2019/6/2/16b169f0605f0bfb?imageView2/0/w/1280/h/960/ignore-error/1)

### 拷贝class和jar到输出目录 ###

使用Tranform一个重要的步骤就是要把所有的class和jar拷贝至输出目录。

` @Override void transform(TransformInvocation transformInvocation) throws TransformException, InterruptedException, IOException { def sTime = System.currentTimeMillis() def inputs = transformInvocation.inputs def routeJarInput def outputProvider = transformInvocation.outputProvider outputProvider.deleteAll() //删除原有输出目录的文件 for (TransformInput input : inputs) { for (DirectoryInput dirInput : input.directoryInputs) { read ClassWithPath(dirInput.file) File dest = outputProvider.getContentLocation(dirInput.name, dirInput.contentTypes, dirInput.scopes, Format.DIRECTORY) FileUtils.copyDirectory(dirInput.file, dest) } for (JarInput jarInput : input.jarInputs) { ... copyFile(jarInput, outputProvider) } } def eTime = System.currentTimeMillis() println( "route map:" + routeMap) insertCodeIntoJar(routeJarInput, transformInvocation.outputProvider) println( "===========route transform finished:" + (eTime - sTime)) } void copyFile(JarInput jarInput, TransformOutputProvider outputProvider) { def dest = getDestFile(jarInput, outputProvider) FileUtils.copyFile(jarInput.file, dest) } static File getDestFile(JarInput jarInput, TransformOutputProvider outputProvider) { def destName = jarInput.name // 重名名输出文件,因为可能同名,会覆盖 def hexName = DigestUtils.md5Hex(jarInput.file.absolutePath) if (destName.endsWith( ".jar" )) { destName = destName.substring(0, destName.length() - 4) } // 获得输出文件 File dest = outputProvider.getContentLocation(destName + "_" + hexName, jarInput.contentTypes, jarInput.scopes, Format.JAR) return dest } 复制代码`

注意insertCodeIntoJar方法中也要copy。 插件模块至此完成。可以运行一下app，打印一下routeMap

![打印信息](https://user-gold-cdn.xitu.io/2019/6/2/16b169f0955036e8?imageView2/0/w/1280/h/960/ignore-error/1) 而具体的路由跳转就不细说了，具体可以看github的项目源码。

### 项目地址 ###

[github.com/iamyours/Si…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fiamyours%2FSimpleRouter )