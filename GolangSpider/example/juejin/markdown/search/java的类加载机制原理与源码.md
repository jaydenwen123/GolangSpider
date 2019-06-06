# java的类加载机制原理与源码 #

编写的java程序编译后会放在以 `.class` 结尾的字节码文件当中，这些字节码文件都放在磁盘上，毫无疑问jvm运行的时候需要从磁盘上读取到对应的字节码文件，那这个过程是怎样的呢？

## class文件的格式 ##

class文件格式采用类似于C的结构体的方式来存储数据

` ClassFile { u4 magic; u2 minor_version; u2 major_version; u2 constant_pool_count; cp_info constant_pool[constant_pool_count-1]; u2 access_flags; u2 this_class; u2 super_class; u2 interfaces_count; u2 interfaces[interfaces_count]; u2 fields_count; field_info fields[fields_count]; u2 methods_count; method_info methods[methods_count]; u2 attributes_count; attribute_info attributes[attributes_count]; } 复制代码`
> 
> 
> 
> 
> [类文件格式](
> https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Fspecs%2Fjvms%2Fse7%2Fhtml%2Fjvms-4.html%23jvms-4.1
> ) 信息存储在方法区中：
> 
> 
> 
> * u4 这中结构表示magic信息占据4字节，类似这种都是class格式中的基本类型
> * cp_info 由多个class格式中的基本类型构成的复合数据类型，当然也可以包含其它定义的复合类型，类似这种结构都是复合类型
> 
> 
> 

所有的class文件中的字节都按照这样的约定紧密的排列，不能出现任何的改动

## class文件中标明的constant_pool ##

constant_pool中主要包含两大类常量：字面量和符号引用。通过一个字节来区分类型

* 

字面量如 CONSTANT_Integer_info

` CONSTANT_Integer_info { u1 tag; u4 bytes; } 复制代码`

如果读到的tag是3表示这个结构就是CONSTANT_Integer_info，接下来的4字节就表示这个int的取值

* 

符号引用如 CONSTANT_Class_info

` CONSTANT_Class_info { u1 tag; u2 name_index; } 复制代码`

如果读到的tag是7表示这个结构是CONSTANT_Class_info，接下来的2字节必须是这个类的constant_pool中的一个有效的索引位置。比如取一个class的字节码 ` 07-》class 00 16->class完` ,07标识着这个常量结构是 CONSTANT_Class，0016标识在这个常量池的第22个下标存储着对应的值

` Constant pool: #3 = Class #22 // main/domain/A... #22 = Utf8 main/domain/A 复制代码`
> 
> 
> 
> 
> * javap -v A.class 可看
> * 符号引用：编译的时候肯定是不知道代码具体在内存中的位置的，只能使用符号，比如 ` main/domain/A` 来事先标识内存地址，直接引用则必定是已经加载到内存
> [符号引用于直接引用的区别](
> https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fshinubi%2Farticles%2F6116993.html
> )
> 
> 
> 

jvm自身会持有所有的常量池，在创建类或者接口的时候，则使用它来构建运行时常量池

> 
> 
> 
> jvm 运行时用术语 run-time constant pool 来表示的class文件中的constant_pool
> 
> 

## classFileParser ##

以hotspot虚拟机为例，在 ` classFileParser.cpp的instanceKlassHandle ClassFileParser::parseClassFile` 中可以看到解析类文件

` instanceKlassHandle ClassFileParser::parseClassFile(Symbol* name, Handle class_loader, Handle protection_domain, KlassHandle host_klass, GrowableArray<Handle>* cp_patches, TempNewSymbol& parsed_name, bool verify, TRAPS) { ... ClassFileStream* cfs = stream(); ... // Magic value u4 magic = cfs->get_u4_fast(); ... // Version numbers u2 minor_version = cfs->get_u2_fast(); u2 major_version = cfs->get_u2_fast(); if (!is_supported_version(major_version, minor_version)){ ... throw... "Unsupported major.minor version %u.%u"... } ... // Constant pool constantPoolHandle cp = parse_constant_pool(CHECK_(nullHandle)); ... } 复制代码`

常量池解析如下

` constantPoolHandle ClassFileParser::parse_constant_pool(TRAPS) { ClassFileStream* cfs = stream(); constantPoolHandle nullHandle; ... u2 length = cfs->get_u2_fast(); ... constantPoolOop constant_pool = oopFactory::new_constantPool(length, oopDesc::IsSafeConc, CHECK_(nullHandle)); constantPoolHandle cp (THREAD, constant_pool); cp->set_partially_loaded(); // Enables heap verify to work on partial constantPoolOops ConstantPoolCleaner cp_in_error(cp); // set constant pool to be cleaned up. // parsing constant pool entries parse_constant_pool_entries(cp, length, CHECK_(nullHandle)); ... } 复制代码`

parse_constant_pool_entries解析常量池如下

` void ClassFileParser::parse_constant_pool_entries(constantPoolHandle cp, int length, TRAPS) { ... // Used for batching symbol allocations. const char* names[SymbolTable::symbol_alloc_batch_size]; int lengths[SymbolTable::symbol_alloc_batch_size]; int indices[SymbolTable::symbol_alloc_batch_size]; unsigned int hash Values[SymbolTable::symbol_alloc_batch_size]; int names_count = 0; // parsing Index 0 is unused for (int index = 1; index < length; index++) { // Each of the following case guarantees one more byte in the stream // for the following tag or the access_flags following constant pool, // so we do not need bounds-check for reading tag. //读取常量池中的第一个字节 u1 tag = cfs->get_u1_fast(); //判断类型做不同的处理 switch (tag) { case JVM_CONSTANT_Class : { cfs->guarantee_more(3, CHECK); // name_index, tag/access_flags u2 name_index = cfs->get_u2_fast(); cp->klass_index_at_put(index, name_index); } break ; ... case JVM_CONSTANT_Utf8 : { ... unsigned int hash ; //查找符号链表中是否已经存在这个符号 Symbol* result = SymbolTable::lookup_only((char*)utf8_buffer, utf8_length, hash ); if (result == NULL) { names[names_count] = (char*)utf8_buffer; lengths[names_count] = utf8_length; indices[names_count] = index; hash Values[names_count++] = hash ; if (names_count == SymbolTable::symbol_alloc_batch_size) { //如果已经达到一定的符号量,取值是8个，就一次性存起来,放入SymbolTable中 SymbolTable::new_symbols(cp, names_count, names, lengths, indices, hash Values, CHECK); names_count = 0; } } else { cp->symbol_at_put(index, result); } } ... } if (names_count > 0) { SymbolTable::new_symbols(cp, names_count, names, lengths, indices, hash Values, CHECK); } } 复制代码`
> 
> 
> 
> 
> [classFileParser由RednaxelaFX在此博客中提及](
> https://link.juejin.im?target=https%3A%2F%2Fhllvm-group.iteye.com%2Fgroup%2Ftopic%2F35385
> )
> 
> 

## 类或接口的创建 ##

触发类或接口C的创建时机包括

* 另一个类或接口D的运行时常量池中包含了对当前类或接口C的引用
* 另一个类或接口D调用了一些特定类库方法，比如反射

给定一个名字N代表要创建的类或接口C

* 如果N不是数组，那么会使用以下两种方式中的一种：如果触发C创建的D是由Bootstrap class loader加载的，就由Bootstrap class loader来加载C。如果触发C创建的D是由用户自定义的加载器加载的，那么C初始加载也会由同一个用户自定义的加载器加载
* N是数组，则直接由JVM创建

### 非数组类加载 ###

使用Bootstrap class loader有如下步骤

* 首先，JVM会看N指定的类或接口是否已经由Bootstrap class loader记录为类的加载器，如果是，就不在创建class
* 没有记录，N会被当做参数传递给BootStrap class loader,执行从class文件提取Class的流程，没有找到，抛出 ClassNotFoundException
* 找到了，从class文件中提取出Class对象

> 
> 
> 
> todo bootstrap 的加载代码
> 
> 

使用用户自定义的类加载器L步骤如下

* 首先，看L是否已经是N的加载器，是就不再创建
* 不是，JVM会调用L的loadClass(N)方法，执行成功则将L标记为C的初始加载器。

自定义类加载器注意事项: 自定义的L必须要实现以下两种方式之一：

* L自己读取到字节数组后，将字节数组传递给ClassLoader的defineClass方法，defineClass则会执行从class字节数组提取Class对象的流程
* L仅作为代理，把加载C的细节交给另一个类加载器L'，一般对应的也是loadClass方法，返回结果就是C的Class对象

根据这种约定，可以自定义类加载器如下

` ClassLoader myLoader = new ClassLoader () { @Override public Class<?> loadClass(String name) throws ClassNotFoundException { try{ String fName=name.substring(name.lastIndexOf( "." ))+ ".class" ; //todo 补充这里的含义 InputStream is=getClass().getResourceAsStream(fName); if (is == null){ return super.loadClass(name); } byte[] b = new byte[is.available()]; is.read(b); return defineClass(name,b,0,b.length); }catch (Exception e){ throw new ClassNotFoundException(name); } } }; 复制代码`

### 数组类加载 ###

数组的组件类型仍然会由类加载器L加载，这里的L即可能是BootstrapClassLoader，也可能是用户自定义。流程与非数组类加载类似，只不过当组件需要加载切加载完成时，JVM会自己创建一个对应的数组

### 从class文件提取Class的方式 ###

假设需要加载的类或接口C使用N唯一标识，加载器为L。提取或经历如下步骤

* JVM查找L是不是已经被标记成了N的初始加载器，是，这个步骤非法，抛出异常LinkageError
* 否则，JVM会尝试去解析，这个过程可能会出现如下异常：如果要解析的文件不符合ClassFile的规范，抛出ClassFormatError;如果ClassFile的major或者minor版本不支持，抛出 UnsupportedClassVersionError;另外如果提供的数据如果并不是C，那么扔出 NoClassDefFoundError
* 如果C拥有父类，解析出其父类的符号引用，如果得到的父类并不是类，是接口，抛出IncompatibleClassChangeError，如果是它自己，抛出ClassCircularityError
> 
> 
> 

* 如果C实现了接口，处理类似父类
* JVM标记L是定义了它的类加载器，并记录L是C的初始加载器

### 常用的类加载器 ###

从JVM的角度来讲，只有两种不同的类加载器：Bootstrap ClassLoader和其它继承了ClassLoader的加载器。绝大部分java程序使用的类加载器如下

* Bootstrap ClassLoader:负责将 java_home\lib 目录下或者是 -Xbootclasspath且虚拟机识别的类库加载到JVM中， 它无法被java程序直接引用
> 
> 
> 
> 
> 这意味着即使是自己写的库放到 java_home\lib 下面也不会被加载
> 
> 

* Extension ClassLoader:负责加载 java_home\lib\ext 目录下的或者被 java.ext.dirs 所指定的路径中的所有类库， 开发者可以使用
* Application ClassLoader:负责加载 classpath 上所指定的类库， 开发者可以使用
> 
> 
> 
> 通过ClassLoader的getSystemClassLoader方法默认返回的就是Application ClassLoader
> 
> 

Application ClassLoader和Extension ClassLoader实现在 ` sun.misc.Launcher` ,Launcher类本身是‘系统’用来启动主应用程序，当它初始化时

` public Launcher () { // Create the extension class loader ClassLoader extcl; try { //它的parent字段设置为null extcl = ExtClassLoader.getExtClassLoader(); } catch (IOException e) { throw new InternalError( "Could not create extension class loader" ); } // Now create the class loader to use to launch the application try { //将extClassLoader作为AppClassLoader的parent字段存储 loader = AppClassLoader.getAppClassLoader(extcl); } catch (IOException e) { throw new InternalError( "Could not create application class loader" ); } // Also set the context class loader for the primordial thread. Thread.currentThread().setContextClassLoader(loader); ... } 复制代码`

由此可见通过组合的方式联系起了 AppClassLoader 和 ExtClassLoader。从JVM的角度来看，AppClassLoader 和 ExtClassLoader均是自定义类加载器，用他们来加载类时，执行loadClass方法，在ClassLoader核心实现如下

` protected Class<?> loadClass(String name, boolean resolve) throws ClassNotFoundException { synchronized (getClassLoadingLock(name)) { // 首先查看类是否已经加载了，如果加载了什么都不做 Class c = findLoadedClass(name); if (c == null) { long t0 = System.nanoTime(); try { if (parent != null) { //没有加载过，以AppClassLoader为例，则是先使用ExtClassLoader去执行加载 c = parent.loadClass(name, false ); } else { //没有加载过，以ExtClassLoader为例，先使用BootstrapClassLoader来加载 c = findBootstrapClassOrNull(name); } } catch (ClassNotFoundException e) { // ClassNotFoundException thrown if class not found // from the non-null parent class loader } if (c == null) { // If still not found, then invoke findClass in order // to find the class. long t1 = System.nanoTime(); //“parent”加载失败，则自己加载 c = findClass(name); // this is the defining class loader; record the stats sun.misc.PerfCounter.getParentDelegationTime().addTime(t1 - t0); sun.misc.PerfCounter.getFindClassTime().addElapsedTimeFrom(t1); sun.misc.PerfCounter.getFindClasses().increment(); } } if (resolve) { resolveClass(c); } return c; } } 复制代码`

从加载的逻辑可以看到，这种方式优先使用BootstrapClassLoader来加载，然后再是ExtClassLoader,最后是ApplicationClassLoader,这种类加载模式又称为 ` 双亲委派模型` 。

## 双亲委派模型 ##

类在加载时，优先交给自己“上游”（“父”）的加载器去执行，他们找不到才自己去加载。优势在于：

* 比如要加载 java.lang.Object ，无论使用什么加载器，这种模式下加载的都是同一个Class。同时从loadClass的约定加载模式来看，就算是写了一个一模一样的类，也不会加载，避免了混乱

> 
> 
> 
> 由loadClass的实现方式来看，如果 ` 要利用双亲委派模型的优势，则自定义类加载器实现findClass是最佳的选择`
> 
> 

当然有些场景需要从父加载器使用子加载器去加载，比如 JNDI，JNDI服务本身由BootstrapClassLoader去加载，但是JNDI自己干的事儿就是要去调用ClassPath下的JNDI接口提供者，它是无法被BootstrapClassLoader加载的，类似还有JDBC，这就是通过 ` Thread.currentThread().setContextClassLoader` 来实现父加载器去用子加载器实现加载

> 
> 
> 
> 另外还有OSGI，在平级类加载器中进行
> 
> 

## BootstrapClassLoader加载机制 ##

ClassLoader最终会优先在BootstrapClassLoader中加载，它是native的实现，最终会由JVM本身去实现加载，对应的方法实现则是在JVM代码中，以hotspot为例，加载的地方为 ` jvm.cpp中的JVM_FindClassFromBootLoader`

` JVM_ENTRY(jclass, JVM_FindClassFromBootLoader(JNIEnv* env, const char* name)) ... //首先从符号表里面去查找，看看能不能找到 TempNewSymbol h_name = SymbolTable::new_symbol(name, CHECK_NULL); //根据查找的结果去解析符号链接 klassOop k = SystemDictionary::resolve_or_null(h_name, CHECK_NULL); if (k == NULL) { return NULL; } ... //转成java对象 return (jclass) JNIHandles::make_local(env, Klass::cast(k)->java_mirror()); JVM_END 复制代码`
> 
> 
> 
> 
> 调用这个方法则是通过动态连接的方式，比如 ` dlsym` 函数
> 
> 

### SymbolTable ###

在 ` symbolTable.hpp` 中可以看到 ` SymbolTable` 的定义

` class SymbolTable : public Hashtable<Symbol*> 复制代码`

它实际上就是一个hashtable，由类文件解析过程可知，它里面装的其实就是已经加载过的类信息

### SystemDictionary ###

用来记录所有已加载的 (类型名, 类加载器) -> 类型 的映射关系

` // Forwards to resolve_instance_class_or_null klassOop SystemDictionary::resolve_or_null(Symbol* class_name, Handle class_loader, Handle protection_domain, TRAPS) { assert(!THREAD->is_Compiler_thread(), "Can not load classes with the Compiler thread" ); if (FieldType::is_array(class_name)) { //数组类型 return resolve_array_class_or_null(class_name, class_loader, protection_domain, CHECK_NULL); } else if (FieldType::is_obj(class_name)) { //对象类型 ResourceMark rm(THREAD); // Ignore wrapping L and ;. TempNewSymbol name = SymbolTable::new_symbol(class_name->as_C_string() + 1, class_name->utf8_length() - 2, CHECK_NULL); return resolve_instance_class_or_null(name, class_loader, protection_domain, CHECK_NULL); } else { return resolve_instance_class_or_null(class_name, class_loader, protection_domain, CHECK_NULL); } } 复制代码`

对实例的处理 ` resolve_instance_class_or_null` :

` klassOop SystemDictionary::resolve_instance_class_or_null(Symbol* name, Handle class_loader, Handle protection_domain, TRAPS) { ... klassOop check = find_class(d_index, d_hash, name, class_loader); if (check != NULL) { // Klass is already loaded, so just return it class_has_been_loaded = true ; k = instanceKlassHandle(THREAD, check); } else { //没有加载过则在PlaceholderTable中查找,如果目标的类加载器和类名一样，就找到 placeholder = placeholders()->get_entry(p_index, p_hash, name, class_loader); if (placeholder && placeholder->super_load_in_progress()) { super_load_in_progress = true ; if (placeholder->havesupername() == true ) { superclassname = placeholder->supername(); havesupername = true ; } } } ... return k(); } 复制代码`

当然在 ` resolve_instance_class_or_null` 没有加载这个名字的类文件的时候，就会由 ` classLoader.cpp中的load_classfile` 来执行加载

` stringStream st; // st.print() uses too much stack space while handling a StackOverflowError // st.print( "%s.class" , h_name->as_utf8()); st.print_raw(h_name->as_utf8()); st.print_raw( ".class" ); char* name = st.as_string(); // Lookup stream for parsing .class file ClassFileStream* stream = NULL; int classpath_index = 0; { PerfClassTraceTime vmtimer(perf_sys_class_lookup_time(), ((JavaThread*) THREAD)->get_thread_stat()->perf_timers_addr(), PerfClassTraceTime::CLASS_LOAD); ClassPathEntry* e = _first_entry; while (e != NULL) { //根据名字找到文件流 stream = e->open_stream(name); if (stream != NULL) { break ; } e = e->next(); ++classpath_index; } } instanceKlassHandle h(THREAD, klassOop(NULL)); if (stream != NULL) { // 找到类文件开始执行解析等等 ClassFileParser parser(stream); ... } 复制代码`

终于这里开始去找文件了

` ClassFileStream* ClassPathDirEntry::open_stream(const char* name) { // construct full path name char path[JVM_MAXPATHLEN]; //将名字和默认的目录匹配到路径 if (jio_snprintf(path, sizeof(path), "%s%s%s" , _dir, os::file_separator(), name) == -1) { return NULL; } // check if file exists struct stat st; if (os:: stat (path, &st) == 0) { // 调用操作系统的方法来找 int file_handle = os::open(path, 0, 0); ... return new ClassFileStream(buffer, st.st_size, _dir); // Resource allocated } } } return NULL; } 复制代码`

此处唯一不确定的是 _dir 即目录的位置是哪儿，也就是说bootstrapLoader在哪个目录加载的， ` _dir` 本身是一个目录对象，它最终由系统设定，追踪可以在os.cpp总看到

` const char* home = Arguments::get_java_home(); ... // Any modification to the JAR-file list, for the boot classpath must be // aligned with install/install/make/common/Pack.gmk. Note: boot class // path class JARs, are stripped for StackMapTable to reduce download size. static const char classpath_format[] = "%/lib/resources.jar:" "%/lib/rt.jar:" "%/lib/sunrsasign.jar:" "%/lib/jsse.jar:" "%/lib/jce.jar:" "%/lib/charsets.jar:" "%/lib/jfr.jar:" #ifdef __APPLE__ "%/lib/JObjC.jar:" #endif "%/classes" ; char* sysclasspath = format_boot_path(classpath_format, home, home_len, fileSep, pathSep); if (sysclasspath == NULL) return false ; Arguments::set_sysclasspath(sysclasspath); 复制代码`

这些也就是bootstrapclassloader加载的东西