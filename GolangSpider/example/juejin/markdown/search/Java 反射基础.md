# Java 反射基础 #

## 一、什么是反射 ##

反射 (Reflection) 是 Java 的特征之一，它允许运行中的 Java 程序获取自身的信息，并且可以操作类或对象的内部属性。

总而言之，通过反射，我们可以在 ` 运行时获得程序或程序集中每一个类型的成员和成员的信息` 。程序中一般的对象的类型都是在编译期就确定下来的，而 ` Java 反射机制可以动态地创建对象并调用其属性` ，这样的对象的类型在编译期是未知的。

` 反射的核心是 JVM 在运行时才动态加载类或调用方法/访问属性，它不需要事先（写代码的时候或编译期）知道运行对象是谁。`

### 1.1 反射机制主要提供以下功能 ###

* 在 ` 运行时` 判断任意一个对象所属的类。
* 在 ` 运行时` 构造任意一个类的对象。
* 在 ` 运行时` 判断任意一个类所具有的成员变量和方法。
* 在 ` 运行时` 调用任意一个对象的方法

当我们的程序在运行时，需要动态的加载一些类，这些类可能之前用不到所以不用加载到jvm，而是在运行时根据需要才加载。

> 
> 
> 
> 例如我们的项目底层有时是用mysql，有时用oracle，需要动态地根据实际情况加载驱动类，这个时候反射就有用了，假设
> com.java.dbtest.mySqlConnection，com.java.dbtest.oracleConnection这两个类我们要用，这时候我们的程序就写得比较动态化，通过Class
> tc =
> Class.forName("com.java.dbtest.mySqlConnection");通过类的全类名让jvm在服务器中找到并加载这个类，而如果是oracle则传入的参数就变成另一个了。这时候就可以看到反射的好处了，这个动态性就体现出java的特性了！
> 
> 
> 

## 二、反射的主要用途 ##

很多人都认为反射在实际的 Java 开发应用中并不广泛，其实不然。当我们在使用 IDE时，我们输入一个对象或类并想调用它的属性或方法时，一按点号，编译器就会自动列出它的属性或方法，这里就会用到反射。

` 反射最重要的用途就是开发各种通用框架` 。很多框架（比如 Spring）都是配置化的（比如通过 XML 文件配置 Bean），为了保证框架的通用性，它们可能需要根据配置文件加载不同的对象或类，调用不同的方法，这个时候就必须用到反射，运行时动态加载需要加载的对象。

## 三、反射的基本运用 ##

上面我们提到了反射可以用于判断任意对象所属的类，获得Class对象，构造任意一个对象以及调用一个对象。

### 3.1 Java 的反射机制的实现借助于4个类：class，Constructor，Field，Method; ###

其中class代表的是类对象，Constructor－类的构造器对象，Field－类的属性对象，Method－类的方法对象，通过这四个对象我们可以粗略的看到一个类的各个组成部分。其中最核心的就是Class类，它是实现反射的基础，它包含的方法我们在第一部分已经进行了基本的阐述。应用反射时我们最关心的一般是一个类的构造器、属性和方法，下面我们主要介绍Class类中针对这三个元素的方法:

### 3.2 获得Class对象 ###

Class类的实例表示Java应用运行时的类（class and enum）或接口（interface and annotation）（每个Java类运行时都在JVM里表现为一个Class对象，可通过类名.class，类型.getClass()，Class.forName("类名")等方法获取Class对象）。基本类型boolean，byte，char，short，int，long，float，double和关键字void同样表现为Class对象。

声明普通的Class对象，在编译器并不会检查Class对象的确切类型是否符合要求，如果存在错误只有在运行时才得以暴露出来。但是通过泛型声明指明类型的Class对象，编译器在编译期将对带泛型的类进行额外的类型检查，确保在编译期就能保证类型的正确性。

#### 使用 Class 类的 forName 静态方法: ####

` public static Class<?> for Name(String className) 比如在 JDBC 开发中常用此方法加载数据库驱动: Class.forName(driver); 复制代码`

#### 直接获取某一个对象的 class ####

` Class<?> klass = int.class; Class<?> classInt = Integer.TYPE; 复制代码`

#### 调用某个对象的 getClass() 方法 ####

` StringBuilder str = new StringBuilder( "123" ); Class<?> klass = str.getClass(); 复制代码`

### 3.3 判断是否为某个类的实例 ###

一般地，我们用 ` instanceof` 关键字来判断是否为某个类的实例。同时我们也可以借助反射中 ` Class` 对象的 isInstance() 方法来判断是否为某个类的实例，它是一个 native 方法：

` public native boolean isInstance(Object obj); 复制代码`

### 3.4 创建实例 ###

#### 使用Class对象的newInstance()方法来创建Class对象对应类的实例 ####

> 
> 
> 
> ` classObj.newInstance() 只能够调用public类型的无参构造函数，此方法是过时的` 。
> 
> 

` Class<?> c = String.class; Object str = c.newInstance(); 复制代码`

#### 先通过Class对象获取指定的Constructor对象，再调用Constructor对象的newInstance()方法来创建实例。 ####

> 
> 
> 
> ` 可以根据传入的参数，调用任意构造函数，在特定情况下，可以调用私有的构造函数，此方法是推荐使用的` 。
> 
> 

` //获取String所对应的Class对象 Class<?> c = String.class; //获取String类带一个String参数的构造器 Constructor constructor = c.getConstructor(String.class); //根据构造器创建实例 Object obj = constructor.newInstance( "23333" ); System.out.println(obj); 复制代码`

### 3.5 获取方法 ###

#### getMethods() ####

> 
> 
> 
> 返回某个类的所有 ` public` 方法， **包括自己声明和从父类继承的** 。
> 
> 

#### getDeclaredMethods() ####

> 
> 
> 
> 获取所有本类自己的方法，不问访问权限，不包括从父类继承的方法
> 
> 

#### getMethod(String name, Class<?>... parameterTypes) ####

> 
> 
> 
> 方法返回一个特定的方法，其中第一个参数为方法名称，后面的参数为方法的参数对应Class的对象。
> 
> 

#### Method getDeclaredMethod(String name, Class<?>... params) ####

> 
> 
> 
> 方法返回一个特定的方法，其中第一个参数为方法名称，后面的参数为方法的参数对应Class的对象
> 
> 

#### 操作私有方法 ####

` /** * 访问对象的私有方法 * 为简洁代码，在方法上抛出总的异常，实际开发别这样 */ private static void getPrivateMethod() throws Exception{ //1. 获取 Class 类实例 TestClass test Class = new TestClass(); Class mClass = test Class.getClass(); //2. 获取私有方法 //第一个参数为要获取的私有方法的名称 //第二个为要获取方法的参数的类型，参数为 Class...，没有参数就是null //方法参数也可这么写 ：new Class[]{String.class , int.class} Method privateMethod = mClass.getDeclaredMethod( "privateMethod" , String.class, int.class); //3. 开始操作方法 if (privateMethod != null) { //获取私有方法的访问权 //只是获取访问权，并不是修改实际权限 privateMethod.setAccessible( true ); //使用 invoke 反射调用私有方法 //privateMethod 是获取到的私有方法 // test Class 要操作的对象 //后面两个参数传实参 privateMethod.invoke( test Class, "Java Reflect " , 666); } } 复制代码`

### 3.6 获取构造函数 ###

#### Constructor[] getConstructors() ####

> 
> 
> 
> 获得类的所有公共构造函数
> 
> 

#### Constructor getConstructor(Class[] params) ####

> 
> 
> 
> 获得使用特殊的参数类型的公共构造函数，
> 
> 

#### Constructor[] getDeclaredConstructors() ####

> 
> 
> 
> 获得类的所有构造函数
> 
> 

#### Constructor getDeclaredConstructor(Class[] params) ####

> 
> 
> 
> 获得使用特定参数类型的构造函数
> 
> 

### 3.7 获取成员变量字段 ###

#### Field[] getFields() ####

> 
> 
> 
> 获得类的所有公共字段
> 
> 

#### Field getField(String name) ####

> 
> 
> 
> 获得命名的公共字段
> 
> 

#### Field[] getDeclaredFields() ####

> 
> 
> 
> 获得类声明的所有字段
> 
> 

#### Field getDeclaredField(String name) ####

> 
> 
> 
> 获得类声明的命名的字段
> 
> 

#### 修改私有变量 ####

` ** * 修改对象私有变量的值 * 为简洁代码，在方法上抛出总的异常 */ private static void modifyPrivateFiled() throws Exception { //1. 获取 Class 类实例 TestClass test Class = new TestClass(); Class mClass = test Class.getClass(); //2. 获取私有变量 Field privateField = mClass.getDeclaredField( "MSG" ); //3. 操作私有变量 if (privateField != null) { //获取私有变量的访问权 privateField.setAccessible( true ); //修改私有变量，并输出以测试 System.out.println( "Before Modify：MSG = " + test Class.getMsg()); //调用 set (object , value) 修改变量的值 //privateField 是获取到的私有变量 // test Class 要操作的对象 // "Modified" 为要修改成的值 privateField.set( test Class, "Modified" ); System.out.println( "After Modify：MSG = " + test Class.getMsg()); } } 复制代码`

### 3.8 调用方法 ###

当我们从类中获取了一个方法后，我们就可以用 invoke() 方法来调用这个方法。invoke 方法的原型为:

` public Object invoke(Object obj, Object... args) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException 复制代码`

下面是一个实例

` public class test 1 { public static void main(String[] args) throws IllegalAccessException, InstantiationException, NoSuchMethodException, InvocationTargetException { Class<?> klass = methodClass.class; //创建methodClass的实例 Object obj = klass.newInstance(); //获取methodClass类的add方法 Method method = klass.getMethod( "add" ,int.class,int.class); //调用method对应的方法 => add(1,4) Object result = method.invoke(obj,1,4); System.out.println(result); } } class methodClass { public final int fuck = 3; public int add(int a,int b) { return a+b; } public int sub(int a,int b) { return a+b; } } 复制代码`

### 3.9 利用反射创建数组 ###

数组在Java里是比较特殊的一种类型，它可以赋值给一个Object Reference。下面我们看一看利用反射创建数组的例子：

` public static void test Array() throws ClassNotFoundException { Class<?> cls = Class.forName( "java.lang.String" ); Object array = Array.newInstance(cls,25); //往数组里添加内容 Array.set(array,0, "hello" ); Array.set(array,1, "Java" ); Array.set(array,2, "fuck" ); Array.set(array,3, "Scala" ); Array.set(array,4, "Clojure" ); //获取某一项的内容 System.out.println(Array.get(array,3)); } 复制代码`

其中的Array类为java.lang.reflect.Array类。我们通过Array.newInstance()创建数组对象，它的原型是:

` public static Object newInstance(Class<?> componentType, int length) throws NegativeArraySizeException { return newArray(componentType, length); } 复制代码`

而 newArray 方法是一个 native 方法

` private static native Object newArray(Class<?> componentType, int length) throws NegativeArraySizeException; 复制代码`

### 3.10 反射修改常量值 ###

![判断可不可以修改常量](https://user-gold-cdn.xitu.io/2019/3/30/169ced4b7300f833?imageView2/0/w/1280/h/960/ignore-error/1)