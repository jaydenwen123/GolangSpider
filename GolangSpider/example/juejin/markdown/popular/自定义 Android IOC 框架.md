# 自定义 Android IOC 框架 #

### 概述 ###

#### 什么是 IOC ####

> 
> 
> 
> Inversion of Control,英文缩写为 IOC，意思为控制反转。
> 
> 

具体什么含义呢？

假设一个类中有很多的成员变量，如果你需要用到里面的成员变量，传统做法是 new 出来进行使用，但是在 IOC 的原则中，我们不要 new，因为这样的耦合度太高，我们可以在需要注入（new）的成员变量上添加注解,等待加载这个类的时候，则进行注入。

那么怎么进行注入呢？

简单的说，就是通过反射的方式，将字符串类路径变为类。

#### 什么是反射 ####

JAVA 并不是一种动态变成语言，为了使语言更加灵活，JAVA 引入了反射机制。JAVA 反射机制是在运行过程中，对于任意一个类，都能知道这个类的所有属性和方法；对于任意一个属性，都能够调用它的任意一个方法；这种动态获取信息以及动态调用对象方法的功能成为 JAVA 语言的反射机制。

#### 什么是注解 ####

JAVA 1.5 之后引入的注解和反射，注解的实现依赖于反射。JAVA 中的注解是一种继承自接口 java.lang.annotation.Annotation 的特殊接口。
那么接口怎么能够设置属性呢？
简单来说就是 JAVA 通过动态代理的方式为你生成了一个实现了接口 Annotation 的实例，然后对该代理实例的属性赋值，这样就可以在程序运行时（如果将注解设置为运行时可见的话）通过反射获取到注解的配置信息。说的通俗一点，注解相当于一种标记，在程序中加了注解就等于为程序打上了某种标记。程序可以利用JAVA的反射机制来了解你的类及各种元素上有无何种标记，针对不同的标记，就去做相应的事件。标记可以加在包，类，方法，方法的参数以及成员变量上。

### 实现 ###

#### 定义注解 ####

* 布局注解
` /** * Created by Keven on 2019/6/3. * * 布局注解 */ //RUNTIME 运行时检测，CLASS 编译时检测 SOURCE 源码资源时检测 @Retention(RetentionPolicy.RUNTIME) //TYPE 用在类上 FIELD 注解只能放在属性上 METHOD 用在方法上 CONSTRUCTOR 构造方法上 @Target(ElementType.TYPE) public @interface KevenContentViewInject { int value();//代表可以 Int 类型,取注解里面的参数 } 复制代码` * 属性组件注解
` /** * Created by Keven on 2019/6/3. * * 组件注解 */ @Retention(RetentionPolicy.RUNTIME) @Target(ElementType.FIELD)//用在属性字段上 public @interface KevenViewInject { int value(); } 复制代码` * 事件注解
` /** * Created by zhengjian on 2019/6/3. * * 事件注解 */ @Retention(RetentionPolicy.RUNTIME) @Target(ElementType.METHOD)//使用在方法上 public @interface KevenOnClickInject { //会有很多个点击事件，所以使用数组 int[] value(); } 复制代码`

#### 实现注入工具类 ####

` /** * Created by Keven on 2019/6/3. * <p> * InjectUtils 注入工具类 */ public class InjectUtils { //注入方法 Activity public static void inject(Activity activity) { injectLayout(activity); injectViews(new ViewFinder(activity), activity); injectEvents(new ViewFinder(activity), activity); } //注入方法 View public static void inject(View view, Activity activity) { injectViews(new ViewFinder(view), activity); injectEvents(new ViewFinder(view), activity); } //注入方法 Fragment public static void inject(View view, Object object) { injectViews(new ViewFinder(view), object); injectEvents(new ViewFinder(view), object); } /** * 事件注入 */ private static void injectEvents(ViewFinder viewFinder, Object object) { // 1.获取所有方法 Class<?> clazz = object.getClass(); Method[] methods = clazz.getDeclaredMethods(); // 2.获取方法上面的所有id for (Method method : methods) { KevenOnClickInject onClick = method.getAnnotation(KevenOnClickInject.class); if (onClick != null) { int[] viewIds = onClick.value(); if (viewIds.length > 0) { for (int viewId : viewIds) { // 3.遍历所有的id 先findViewById然后 set OnClickListener View view = viewFinder.findViewById(viewId); if (view != null) { view.setOnClickListener(new DeclaredOnClickListener(method, object)); } } } } } } private static class DeclaredOnClickListener implements View.OnClickListener { private Method mMethod; private Object mHandlerType; public DeclaredOnClickListener(Method method, Object handlerType) { mMethod = method; mHandlerType = handlerType; } @Override public void onClick(View v) { // 4.反射执行方法 mMethod.setAccessible( true ); try { mMethod.invoke(mHandlerType, v); } catch (Exception e) { e.printStackTrace(); try { mMethod.invoke(mHandlerType, null); } catch (Exception e1) { e1.printStackTrace(); } } } } //控件注入 private static void injectViews(ViewFinder viewFinder, Object object) { //获取每一个属性上的注解 Class<?> myClass = object.getClass(); Field[] myFields = myClass.getDeclaredFields();//先拿到所有的成员变量 for (Field field : myFields) { KevenViewInject myView = field.getAnnotation(KevenViewInject.class); if (myView != null) { int value = myView.value();//拿到属性id View view = viewFinder.findViewById(value); //将view 赋值给类里面的属性 try { field.setAccessible( true );//为了防止其是私有的，设置允许访问 field.set(object, view); } catch (IllegalAccessException e) { e.printStackTrace(); } } } } private static void injectLayout(Activity activity) { //获取我们自定义类KevenContentViewInject 上面的注解 Class<?> myClass = activity.getClass(); KevenContentViewInject myContentView = myClass.getAnnotation(KevenContentViewInject.class); if (myContentView!=null){ int myLayoutResId = myContentView.value(); activity.setContentView(myLayoutResId); } } } 复制代码`

#### 定义 ViewFinder 类 ####

用于注入工具类中的 findViewById

` /** * Created by Keven on 2019/6/3. */ final class ViewFinder { private View view; private Activity activity; public ViewFinder(View view) { this.view = view; } public ViewFinder(Activity activity) { this.activity = activity; } public View findViewById(int id) { if (view != null) return view.findViewById(id); if (activity != null) return activity.findViewById(id); return null; } public View findViewById(int id, int pid) { View pView = null; if (pid > 0) { pView = this.findViewById(pid); } View view = null; if (pView != null) { view = pView.findViewById(id); } else { view = this.findViewById(id); } return view; } /*public Context getContext () { if (view != null) return view.getContext(); if (activity != null) return activity; return null; }*/ } 复制代码`

### 使用 IOC 框架 ###

#### 布局文件 ####

` <?xml version= "1.0" encoding= "utf-8" ?> <android.support.constraint.ConstraintLayout xmlns:android= "http://schemas.android.com/apk/res/android" xmlns:app= "http://schemas.android.com/apk/res-auto" xmlns:tools= "http://schemas.android.com/tools" android:layout_width= "match_parent" android:layout_height= "match_parent" tools:context= ".ioc.IocActivity" > <TextView android:id= "@+id/tv_title" android:layout_width= "wrap_content" android:layout_height= "wrap_content" android:layout_marginTop= "@dimen/dimen_35dp" android:text= "你好，IOC" android:textSize= "25sp" app:layout_constraintLeft_toLeftOf= "parent" app:layout_constraintRight_toRightOf= "parent" app:layout_constraintTop_toTopOf= "parent" /> <Button android:id= "@+id/bt_pop" android:layout_width= "wrap_content" android:layout_height= "wrap_content" android:layout_marginTop= "@dimen/dimen_35dp" android:text= "弹窗" app:layout_constraintLeft_toLeftOf= "parent" app:layout_constraintRight_toRightOf= "parent" app:layout_constraintTop_toBottomOf= "@+id/tv_title" /> </android.support.constraint.ConstraintLayout> 复制代码`

#### Activity 代码 ####

` //布局文件注入 @KevenContentViewInject(R.layout.activity_ioc) public class IocActivity extends AppCompatActivity { //属性控件注入 @KevenViewInject(R.id.tv_title) private TextView tv_title; @KevenViewInject(R.id.bt_pop) private Button bt_pop; @Override public void onCreate(Bundle savedInstanceState) { super.onCreate(savedInstanceState); //注入工具绑定 InjectUtils.inject(this); } //点击事件注入 @KevenOnClickInject(R.id.bt_pop) public void change (){ tv_title.setText( "hello IOC" ); Toast.makeText(this, "Hello IOC" ,Toast.LENGTH_SHORT).show(); } } 复制代码`

当我们点击弹窗按钮时，上方 TextView 内容会改变，并且有 Toast 弹出。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1ca345d2b1a9a?imageView2/0/w/1280/h/960/ignore-error/1)