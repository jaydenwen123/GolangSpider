# Java 如何优雅的使用注解 #

## 什么是注解 ##

Java注解可以想象成代码是具有生命的，注解就是对于代码中的某些鲜活的个体贴上一张标签。 ` 简单的说，注解就如同一张标签` 。

## 元注解 ##

元注解是可以注解到注解上的注解，或者说元注解是一种基本注解，但是它能够应用到其它的注解上面。

其实说白了， ` 元注解也是一张标签，但是它是一张特殊的标签，它的作用和目的就是给其他普通的标签进行解释说明的` 。

### 元注解的类型 ###

* 

@Documented

> 
> 
> 
> 如果使用 ` @Documented` 修饰 ` Annotation` ，则表示它可以出现在javadoc中。 定义Annotation时， `
> @Documented可有可无` ；若没有定义，则Annotation不会出现在javadoc中。
> 
> 

* 

@Retention

> 
> 
> 
> Retention 的英文意为保留期的意思。当 @Retention
> 应用到一个注解上的时候，它解释说明了这个注解的的存活时间。定义Annotation时， ` @Retention可有可无` 。若没有@Retention，则@Retention的默认取值是
> ` RetentionPolicy.CLASS` 。
> 
> 
> 
> 它的取值如下：
> 
> 
> 
> * RetentionPolicy.SOURCE 注解只在源码阶段保留，在编译器进行编译时它将被丢弃并忽视。
> * RetentionPolicy.CLASS 注解只被保留到编译进行的时候，在加载到 JVM 之前进行丢弃并忽略，即不会加载到JVM中。。
> * RetentionPolicy.RUNTIME 注解可以保留到程序运行的时候，它会被加载进入到 JVM 中，所以在程序运行时可以获取到它们。
> 
> 
> 

* 

@Target

> 
> 
> 
> Target 是目标的意思，@Target 指定了注解运用的地方。 定义Annotation时， ` @Target可有可无` 。若有@Target，则该Annotation只能用于它所指定的地方；若没有@Target，则该Annotation可以用于任何地方。
> 
> 
> 
> 
> 你可以这样理解，当一个注解被 @Target 注解时，这个注解就被限定了运用的场景。
> 
> 
> 
> @Target 有下面的取值
> 
> 
> 
> * ElementType.PACKAGE 可以给一个包进行注解
> * ElementType.ANNOTATION_TYPE 可以给一个注解进行注解
> * ElementType.TYPE 可以给一个类型进行注解，比如类、接口、枚举
> * ElementType.CONSTRUCTOR 可以给构造方法进行注解
> * ElementType.FIELD 可以给属性进行注解
> * ElementType.METHOD 可以给方法进行注解
> * ElementType.PARAMETER 可以给一个方法内的参数进行注解
> * ElementType.LOCAL_VARIABLE 可以给局部变量进行注解
> 
> 
> 

* 

@Repeatable

> 
> 
> 
> Repeatable 自然是可重复的意思.是指由另一个注解来存储重复注解，在使用的时候，用存储注解来扩展重复注解。
> 
> ![](https://user-gold-cdn.xitu.io/2019/4/4/169e7c90f6e93803?imageView2/0/w/1280/h/960/ignore-error/1)
> 创建重复注解Authority时，加上@Repeatable,指向存储注解Authorities，在使用时候，直接可以重复使用Authority注解。
> 
> 
> 
> 
> 

* 

@Inherited

> 
> 
> 
> @Inherited
> 元注解是一个标记注解，@Inherited阐述了某个被标注的类型是被继承的。如果一个使用了@Inherited修饰的annotation类型被用于一个class，则这个annotation将被用于该class的子类。
> 
> 
> 
> 
> 注意：@Inherited
> annotation类型是被标注过的class的子类所继承。类并不从它所实现的接口继承annotation，方法并不从它所重载的方法继承annotation。
> 
> 
> 
> 
> 当@Inherited
> annotation类型标注的annotation的Retention是RetentionPolicy.RUNTIME，则反射API增强了这种继承性。如果我们使用java.lang.reflect去查询一个@Inherited
> annotation类型的annotation时，反射代码检查将展开工作：检查class和其父类，直到发现指定的annotation类型被发现，或者到达类继承结构的顶层。
> 
> 
> 

## 注解语法 ##

因为平常开发少见使用注解，导致有不少人认为注解的地位不高。其实 ` Annotation` 同 ` classs` 和 ` interface` 一样，注解也是一种类型，只不过它是在 Java SE 5.0 版本中才开始引入的概念。

一个 ` Annotation` 和一个 ` @Retention中的RetentionPolicy` 关联，即每个 ` Annotation` ，都会有唯一的 ` RetentionPolicy` 的属性。

1个 ` Annotation` 和 ` 1~n个@Target中的ElementType` 关联，即对于每 ` 1个Annotation对象` ，可以有 ` 若干个@Target的ElementType` 属性值。

示例：

` @Documented @Retention(RetentionPolicy.RUNTIME) @Target(value = {ElementType.ANNOTATION_TYPE, ElementType.TYPE, ElementType.METHOD, ElementType.FIELD}) public @interface AnnotationTest { } 复制代码`
> 
> 
> 
> 
> 上面的示例，使用 ` @Documented` 说明可以出现在Java文档中 `
> ，@Retention(RetentionPolicy.RUNTIME)` 说明AnnotationTest可以加载到JVM
> 中，可以在程序运行的时候的时候通过反射获取到。@Target(value = {ElementType.ANNOTATION_TYPE,
> ElementType.TYPE, ElementType.METHOD,
> ElementType.FIELD})说明AnnotationTest可以用于注解，类和接口，方法，属性中。
> 
> 

### 注解的定义 ###

注解通过 ` @interface` 关键字进行定义。自动继承了 ` java.lang.annotation.Annotation` 接口，由编译程序自动完成其他细节

### 定义注解时的注意事项 ###

* 在定义注解时， ` 不能继承其他的注解或接口` 。
* ` @interface` 用来声明一个注解，其中的每一个方法实际上相当于一个 ` 配置参数` 。 ` 方法的名称就是参数的名称` ， ` 返回值类型就是参数的类型` 。
* 可以通过 ` default` 来声明参数的默认值。

### 注解参数支持数据类型： ###

* 所有基本数据类型（int,float,boolean,byte,double,char,long,short)
* String类型
* Class类型
* enum类型
* Annotation类型
* 以上所有类型的数组

### 注解参数的设定: ###

第一,只能用 ` public` 或 ` default` 这两个访问权修饰.例如, ` String value();` 这里把方法设为defaul默认类型；

第二,参数成员只能使用注解参数支持的数据类型

第三,如果只有一个参数成员,最好把参数名称设为 ` value` , ` 后加小括号`.

### 注解元素的默认值： ###

` 注解元素必须有确定的值` ，要么在定义注解的默认值中指定，要么在使用注解时指定，非基本类型的注解元素的值不可为null。因此, 使用空字符串或0作为默认值是一种常用的做法。这个约束使得处理器很难表现一个元素的存在或缺失的状态，因为每个注解的声明中，所有元素都存在，并且都具有相应的值，为了绕开这个约束，我们只能定义一些特殊的值，例如空字符串或者负数，一次表示某个元素不存在，在定义注解时，这已经成为一个习惯用法。

定义了注解，并在需要的时候给相关类，类属性加上注解信息，如果没有相应的注解信息处理流程，注解可以说是没有实用价值。如何让注解真真的发挥作用，主要就在于注解处理方法。

如果没有用来读取注解的方法和工作，那么注解也就不会比注释更有用处了。使用注解的过程中，很重要的一部分就是创建于使用注解处理器。Java SE5扩展了反射机制的API，以帮助程序员快速的构造自定义注解处理器。

简单的自定义注解和使用注解实例：

` @Documented @Retention(RetentionPolicy.RUNTIME) @Target(value = {ElementType.ANNOTATION_TYPE, ElementType.TYPE, ElementType.METHOD, ElementType.FIELD}) public @interface AnnotationTest { /*** * 实体默认firstLevelCache属性为 false * @ return boolean */ boolean firstLevelCache() default false ; /*** * 实体默认secondLevelCache属性为 false * @ return boolean */ boolean secondLevelCache() default true ; /*** * 表名默认为空 * @ return String */ String tableName() default "" ; /*** * 默认以 "" 分割注解 */ String split() default "" ; } 复制代码`

它的形式跟接口很类似，只不过前面多了一个 ` @` 符号。通过上面的语句，就可以创建一个名为 ` AnnotationTest` 的注解。

### 注解的使用 ###

` @AnnotationTest public class Test { } 复制代码`

上面的代码，创建一个类 ` Test,` 然后在类定义的地方加上 ` @AnnotationTest` 就可以用 ` AnnotationTest` 注解这个类了。可以简单的理解为将 ` AnnotationTest` 这张标签贴到 Test 这个类上面。

## Java预置的注解 ##

### @Deprecated ###

这个元素是 ` 用来标记过时的元素` 。编译器在编译阶段遇到这个注解时会发出提醒警告，告诉开发者正在调用一个过时的元素比如过时的方法、过时的类、过时的成员变量。

使用如下

` public class Hero { @Deprecated public void say (){ System.out.println( "Noting has to say!" ); } public void speak (){ System.out.println( "I have a dream!" ); } } 复制代码`
> 
> 
> 
> 
> 在IDE中调用hero.say()时，say()方法的上面会有方法过时的提醒。
> 
> 

### @Override ###

这个大家应该很熟悉了，提示子类要复写父类中被 @Override 修饰的方法

### @SuppressWarnings ###

阻止警告的意思。之前说过调用被 @Deprecated 注解的方法后，编译器会警告提醒，而有时候开发者会忽略这种警告，他们可以在调用的地方通过 @SuppressWarnings 达到目的。

### @SafeVarargs ###

参数安全类型注解。它的目的是提醒开发者不要用参数做一些不安全的操作,它的存在会阻止编译器产生 unchecked 这样的警告。

` @SafeVarargs // Not actually safe! static void m(List<String>... stringLists) { Object[] array = stringLists; List<Integer> tmpList = Arrays.asList(42); array[0] = tmpList; // Semantically invalid, but compiles without warnings String s = stringLists[0].get(0); // Oh no, ClassCastException at runtime! } 复制代码`

### @FunctionalInterface ###

函数式接口 (Functional Interface) 就是一个只具有一个方法的普通接口。

例如：

` @FunctionalInterface public interface Runnable { /** * When an object implementing interface <code>Runnable</code> is used * to create a thread, starting the thread causes the object 's * <code>run</code> method to be called in that separately executing * thread. * <p> * The general contract of the method <code>run</code> is that it may * take any action whatsoever. * * @see java.lang.Thread#run() */ public abstract void run(); } 复制代码`

可能有人会疑惑，函数式接口标记有什么用，这个原因是函数式接口可以很容易转换为 Lambda 表达式

## 注解的提取 ##

Java使用Annotation接口来代表程序元素前面的注解，该接口是所有Annotation类型的父接口。除此之外，Java在java.lang.reflect 包下新增了 ` AnnotatedElement` 接口，该接口代表程序中可以接受注解的程序元素，该接口主要有如下几个实现类：

* Class：类定义
* Constructor：构造器定义
* Field：类的成员变量定义
* Method：类的方法定义
* Package：类的包定义
* 

java.lang.reflect 包下主要包含一些实现反射功能的工具类，实际上，java.lang.reflect 包所有提供的反射API扩充了运行时读取Annotation信息的能力。当一个Annotation类型被定义为运行时的Annotation后，该注解才能是运行时可见，当class文件被装载时被保存在class文件中的Annotation才会被虚拟机读取。

` AnnotatedElement` 接口是所有程序元素（Class、Method和Constructor）的父接口，所以程序通过反射获取了某个类的AnnotatedElement对象之后，程序就可以调用该对象的如下四个个方法来访问Annotation信息：

* 方法1： T getAnnotation(Class annotationClass): 返回该程序元素上存在的、指定类型的注解，如果该类型注解不存在，则返回null。
* 方法2：Annotation[] getAnnotations():返回该程序元素上存在的所有注解。
* 方法3：boolean is AnnotationPresent(Class<?extends Annotation> annotationClass):判断该程序元素上是否包含指定类型的注解，存在则返回true，否则返回false.
* 方法4：Annotation[] getDeclaredAnnotations()：返回直接存在于此元素上的所有注释。与此接口中的其他方法不同，该方法将忽略继承的注释。（如果没有注释直接存在于此元素上，则返回长度为零的一个数组。）该方法的调用者可以随意修改返回的数组；这不会对其他调用者返回的数组产生任何影响。

### 注解与反射 ###

注解通过反射获取。

* 首先可以通过 Class 对象的 isAnnotationPresent() 方法判断它是否应用了某个注解
` public boolean isAnnotationPresent(Class<? extends Annotation> annotationClass) {} 复制代码` * 然后通过 getAnnotation() 方法来获取 Annotation 对象。
` public <A extends Annotation> A getAnnotation(Class<A> annotationClass) {} 复制代码`

或者是 getAnnotations() 方法

` public Annotation[] getAnnotations () {} 复制代码`
> 
> 
> 
> 前一种方法返回指定类型的注解，后一种方法返回注解到这个元素上的所有注解。
> 
> 

使用方法

` @TestAnnotation() public class Test { public static void main(String[] args) { boolean hasAnnotation = Test.class.isAnnotationPresent(AnnotationTest.class); if ( hasAnnotation ) { TestAnnotation test Annotation = Test.class.getAnnotation(AnnotationTest.class); System.out.println( "id:" + test Annotation.firstLevelCache()); System.out.println( "msg:" + test Annotation.tableName()); } } } 复制代码`

## 注解的使用场景 ##

> 
> 
> 
> 注解是一系列元数据，它提供数据用来解释程序代码，但是注解并非是所解释的代码本身的一部分。注解对于代码的运行效果没有直接影响。
> 
> 

注解用处主要如下：

* 提供信息给编译器
> 
> 
> 
> 编译器可以利用注解来探测错误和警告信息
> 
> 

* 编译阶段时的处理
> 
> 
> 
> 软件工具可以用来利用注解信息来生成代码、Html文档或者做其它相应处理。
> 
> 

* 运行时的处理
> 
> 
> 
> 某些注解可以在程序运行的时候接受代码的提取
> 
>