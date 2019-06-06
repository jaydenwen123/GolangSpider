# Java动态代理从入门到原理再到实战 #

## 目录 ##

* 前言
* 什么是动态代理，和静态代理有什么区别
* Java动态代理的简单使用
* Java动态代理的原理解读
* 动态代理在Android中的使用

##前言 相信动态代理这个词对于很多Android开发的小伙伴来说既熟悉又陌生，熟悉是应为可能常常会听一些群里，博客上的装B能手挂在嘴边，陌生是因为在日常的Android开发中似乎没有用到过这个东西，也没有自己去学过这个东西（特别是培训班出来的小伙伴们，据我说知大部分Android培训班是不会说这东西的，少部分也只是简单提一下就略过了）。而这种熟悉又陌生的感觉让某些小伙伴觉得动态代理很高级，是那些大佬们才用的知识。而我，今天就帮助小伙伴们把动态代理拉下神坛，下次再有装逼小能手和你说动态代理的时候你就敲着他的脑袋说L（老）Z（子）也会😁。

## 什么是动态代理，和静态代理有什么区别 ##

动态代理和静态代理其实都可以看成是对代理模式的一个使用，什么是代理模式呢？

> 
> 
> 
> 给某个对象提供一个代理对象，并由代理对象提供和控制对原对象的访问。
> 代理模式其实是一个很简单的设计模式，具体的细节小伙伴们可以自己百度，这里就不多做介绍。
> 
> 

静态代理中，需要为被代理的类（委托类）手动编写一个代理类，为需要被代理的每一个方法都写一个新的方法，这部分在编译之前就需要完成了。而动态代理则是可以在运行中动态生成一个代理类，并且不需要去手动的实现每个被代理的方法。简单来说委托类中有1000个方法需要被代理（比如代理的目的就是大家常常用来举例的在每个方法执行前后额外打印一段输出），使用静态代理你需要手动的编写代理类并且实现这1000个方法，而使用静态代理你则需要简单的几行代码就能实现这个问题。

## Java动态代理的简单使用 ##

### 动态代理的相关类 ###

动态代理主要两个相关类：

* Proxy（java.lang.reflect包下的），主要负责管理和创建代理类的工作。
* InvocationHandler 接口，只拥有一个invoke方法，主要负责方法调用部分，是动态代理中我们需要实现的方法

` public Object invoke (Object proxy, Method method, Object[] args) throws Throwable ; //三个参数分别是 代理类proxy 委托类被执行的方法method 方法传入的参数args //返回值则是method方法执行的返回值 复制代码`

下面我们来举个具体的使用例子，假设有下面这么一个接口和两个类来负责对订单的操作

` public interface OrderHandler { void handler (String orderId) ; } public class NormalOrderHandler implements OrderHandler { @Override public void handler (String orderId) { System.out.println( "NormalOrderHandler.handler():orderId:" +orderId); } } public class MaxOrderHandler implements OrderHandler { @Override public void handler (String orderId) { System.out.println( "MaxOrderHandler.handler():orderId:" +orderId); } } 复制代码`

而这个时候要求我们在每个handler方法调用前后都打印一串信息，并且当orderId的长度大于10时截取前10位（不要问我这种需求是哪个RZ提的，反正不是博主，你懂得😁） 这个时候我们可以做如下编码：

` //创建一个处理器类实现InvocationHandler接口并实现invoke方法 public class OrderHandlerProxy implements InvocationHandler { //委托类 在这里就相当于实现了OrderHandler的类的对象 Object target; public Object bind (Object target) { this.target=target; //重点之一，通过Proxy的静态方法创建代理类 第一个参数为委托类的类加载器， //第二个参数为委托类实现的接口集，第三个参数则是处理器类本身 return Proxy.newProxyInstance( this.target.getClass().getClassLoader(), this.target.getClass().getInterfaces(), this ); } //重点之二 @Override public Object invoke (Object proxy, Method method, Object[] args) throws Throwable { //判断执行的方法是否是我们需要代理的handler方法 if (method.getName().equalsIgnoreCase( "handler" )){ System.out.println( "OrderHandlerProxy.invoke.before" ); String orderId= (String) args[ 0 ]; //对orderId的长度做限制 if (orderId.length()>= 10 ){ orderId=orderId.substring( 0 , 10 ); } //重点之三，这个地方通过反射调用委托类的方法 Object invoke = method.invoke(target, orderId); System.out.println( "OrderHandlerProxy.invoke.after" ); return invoke; } else { //当前执行的方法不是我们需要代理的方法时不做操作直接执行委托的相应方法 System.out.println( "Method.name:" +method.getName()); return method.invoke(target,args); } } } 复制代码`

接下来则是具体的使用动态代理的代码

` public class NormalApplicationClass { public void handlerOrder (OrderHandler orderHandler,String orderId) { //创建处理器对象 OrderHandlerProxy proxy= new OrderHandlerProxy(); //为传入的实现了OrderHandler接口的对象创建代理类并实例化对象 OrderHandler handler = (OrderHandler) proxy.bind(orderHandler); handler.handler(orderId); System.out.println(handler.toString()); } public static void main (String[] args) { NormalApplicationClass app= new NormalApplicationClass(); app.handlerOrder( new MaxOrderHandler(), "012345678999" ); } } 复制代码`

具体的返回值如下

` OrderHandlerProxy.invoke.before MaxOrderHandler.handler():orderId:0123456789 OrderHandlerProxy.invoke.after Method.name:toString com.against.javase.rtti.use.MaxOrderHandler@d716361 复制代码`

由上面可见，我们成功的在执行handler方法前后打印了一串东西，并且对orderId进行了长度限制，同时并不会影响对象本身的像toString之类的其他方法调用造成影响。

## 动态代理的原理解读 ##

我们知道动态代理和静态代理的一大区别无法就是创建代理类的过程不一样，而动态代理中创建代理类的操作主要是在Proxy.newProxyInstance()方法中，那我们主要来看下这个方法的源码

` public static Object newProxyInstance (ClassLoader loader, Class<?>[] interfaces, InvocationHandler h) throws IllegalArgumentException { //省略掉部分代码，主要是一些Jvm安全性的验证和权限的验证 /* * 获取或者生成一个代理类的Class对象。为什么是获取或生成呢？ * 这是应为有缓存机制的存在， * 当第二次调用newProxyInstance方法并且上次生成的代理类的缓存还未过期时会直接获取 */ Class<?> cl = getProxyClass0(loader, intfs); /* * 通过反射获取代理类的构造器，并通过构造器生成代理类对象 */ try { //... //通过反射获取代理类的构造器， final Constructor<?> cons = cl.getConstructor(constructorParams); final InvocationHandler ih = h; if (!Modifier.isPublic(cl.getModifiers())) { AccessController.doPrivileged( new PrivilegedAction<Void>() { public Void run () { cons.setAccessible( true ); return null ; } }); } //通过构造器生成代理对象，这里返回的就是给我们使用的代理类的对象了 return cons.newInstance( new Object[]{h}); } //....省略掉一些无关代码 } 复制代码`

通过上面代码不然看出Proxy是先拿到代理类的Class对象，然后通过反射获取构造器，从构造器注入InvocationHandler实例（就是我们自己实现的处理器类）并创建代理类的实例。而获取代理类的Class对象的操作则在Class<?> cl = getProxyClass0(loader, intfs);这里，那么我们继续跟踪，发现里面非常简单，只是做了一些常规的检查并且调用了proxyClassCache.get(loader, interfaces);。我们查看发现

` /** * a cache of proxy classes */ private static final WeakCache<ClassLoader, Class<?>[], Class<?>> proxyClassCache = new WeakCache<>( new KeyFactory(), new ProxyClassFactory()); //看到WeakCache我们就能大概猜到这东西是做什么的了吧？而看ProxyClassFactory的类名我们不难看出创建代理类的Class对象的操作都在这里面 复制代码`

而这个类里面也非常简单，主要方法就是其中的 apply方法。

` @Override public Class<?> apply(ClassLoader loader, Class<?>[] interfaces) { //...前面主要是一些生成proxyName，代理类类名的操作，省略掉 //具体的生成操作就是在这个方法里面 byte [] proxyClassFile = ProxyGenerator.generateProxyClass( proxyName, interfaces, accessFlags); try { //这个方法是一个native方法，通过类加载器和类名以及类的文件的字节流来加载出Class对象。 return defineClass0(loader, proxyName,proxyClassFile, 0 , proxyClassFile.length); } catch (ClassFormatError e) { throw new IllegalArgumentException(e.toString()); } } 复制代码`

接下来已经可以定位了类的生成操作就在ProxyGenerator.generateProxyClass(proxyName, interfaces, accessFlags);中 这部分的代码比较繁杂，我们只是简单的看一个地方。

` this.addProxyMethod(hashCodeMethod, Object.class); this.addProxyMethod(equalsMethod, Object.class); this.addProxyMethod(toStringMethod, Object.class); //通过这个地方我们知道了，为什么我们对代理类的toString等方法的调用也会走代理类的invoke方法，这是应为在生成委托类的时候帮我们重写了这几个方法。 //生成一个带有InvocationHandler参数的构造器 private ProxyGenerator. MethodInfo generateConstructor () throws IOException { ProxyGenerator.MethodInfo var1 = new ProxyGenerator.MethodInfo( "<init>" , "(Ljava/lang/reflect/InvocationHandler;)V" , 1 ); DataOutputStream var2 = new DataOutputStream(var1.code); this.code_aload( 0 , var2); this.code_aload( 1 , var2); var2.writeByte( 183 ); var2.writeShort( this.cp.getMethodRef( "java/lang/reflect/Proxy" , "<init>" , "(Ljava/lang/reflect/InvocationHandler;)V" )); var2.writeByte( 177 ); var1.maxStack = 10 ; var1.maxLocals = 2 ; var1.declaredExceptions = new short [ 0 ]; return var1; } 复制代码`

而接下来的操作无非是根据委托类实现的接口，来生成相应的代理方法，并且生成一个需要传递InvocationHandler对象的构造器，只不过这里生成的是.class文件，而不是我们手动编写的.java文件，而.class文件是已经编译过的文件，jvm可以直接加载进去执行，省去了编译这一步骤。 下面我们看看Proxy为我们生成的代理类是什么样子的：

` public final class $ Proxy0 extends Proxy implements OrderHandler { private static Method m1; private static Method m2; private static Method m3; private static Method m0; public $Proxy0(InvocationHandler var1) throws { super (var1); } public final boolean equals (Object var1) throws { try { return (Boolean) super.h.invoke( this , m1, new Object[]{var1}); } catch (RuntimeException | Error var3) { throw var3; } catch (Throwable var4) { throw new UndeclaredThrowableException(var4); } } public final String toString () throws { try { return (String) super.h.invoke( this , m2, (Object[]) null ); } catch (RuntimeException | Error var2) { throw var2; } catch (Throwable var3) { throw new UndeclaredThrowableException(var3); } } public final void handler (String var1) throws { try { super.h.invoke( this , m3, new Object[]{var1}); } catch (RuntimeException | Error var3) { throw var3; } catch (Throwable var4) { throw new UndeclaredThrowableException(var4); } } public final int hashCode () throws { try { return (Integer) super.h.invoke( this , m0, (Object[]) null ); } catch (RuntimeException | Error var2) { throw var2; } catch (Throwable var3) { throw new UndeclaredThrowableException(var3); } } static { try { m1 = Class.forName( "java.lang.Object" ).getMethod( "equals" , Class.forName( "java.lang.Object" )); m2 = Class.forName( "java.lang.Object" ).getMethod( "toString" ); m3 = Class.forName( "com.against.javase.rtti.use.OrderHandler" ).getMethod( "handler" , Class.forName( "java.lang.String" )); m0 = Class.forName( "java.lang.Object" ).getMethod( "hashCode" ); } catch (NoSuchMethodException var2) { throw new NoSuchMethodError(var2.getMessage()); } catch (ClassNotFoundException var3) { throw new NoClassDefFoundError(var3.getMessage()); } } } 复制代码`

我们可以看到这个代理类本质上是实现了委托类的接口，并且把每一个对委托类实现的接口方法的调用传递到我们编写的处理器类的invoke方法上，一切就清晰明了，动态代理并不神秘哈哈😁。

## 动态代理在Android中的使用 ##

> 
> 
> 
> 纸上得来终觉浅，绝知此事要躬行
> 这句话说得很有道理，博主在学习的过程中深刻感觉到看一千遍不如自己来写一遍，有些知识点看着感觉懂了，不自己写根本发现不了一些问题，不自己写也很难有深刻的映像。现在大部分博客讲到动态代理的时候都会想我上面的一样举一个简单的例子讲一下相关方法就结束了，让人看完感觉什么都懂了，又好像什么都没看懂。
> 
> 
> 

### 实战之简单模仿butterknife（黄油刀）的功能 ###

（前排提示阅读这部分内容需要一些反射的知识） 在选择具体的例子时，我想了一些东西，比如hook SystemManager，但是这部分相比较于动态代理的东西更多的还是和Android系统层面相关的东西，写起来也比较繁琐而且也很难用很少的代码做出一些有意思的操作。最后我想了决定模仿一下黄油刀的简单的功能，说起黄油刀大家应该都不陌生，下面这两个注解肯定都用过：

` @BindView () @OnClick () 复制代码`

今天我们就来实现下这两个注解的功能，下面直接上代码 先定义两个注解

` @Target (ElementType.FIELD) @Retention (RetentionPolicy.RUNTIME) public @interface InjectView { int value () ; } @Target (ElementType.METHOD) @Retention (RetentionPolicy.RUNTIME) public @interface OnClick { int [] value(); } 复制代码`

接着在Activity中使用这两个注解

` public class MainActivity extends AppCompatActivity { @InjectView (R.id.tv) private TextView tv; @InjectView (R.id.iv) private ImageView iv; @Override protected void onCreate (Bundle savedInstanceState) { super.onCreate(savedInstanceState); setContentView(R.layout.activity_main); ViewInject.inject( this ); tv.setText( "inject success!" ); iv.setImageDrawable(getResources().getDrawable(R.mipmap.ic_launcher)); } @OnClick ({R.id.tv,R.id.iv}) public void onClick (View view) { switch (view.getId()){ case R.id.tv: Toast.makeText( this , "onClick,TV" ,Toast.LENGTH_SHORT).show(); break ; case R.id.iv: Toast.makeText( this , "onClick,IV" ,Toast.LENGTH_SHORT).show(); break ; } } } 复制代码`

这部分代码很简单，没有什么好说的，重点在ViewInject.inject(this);这里

` public class ViewInject { public static void inject (Activity activity) { Class<? extends Activity> activityKlazz = activity.getClass(); try { injectView(activity,activityKlazz); proxyClick(activity,activityKlazz); } catch (NoSuchMethodException e) { e.printStackTrace(); } catch (IllegalAccessException e) { e.printStackTrace(); } catch (InvocationTargetException e) { e.printStackTrace(); } } private static void proxyClick (Activity activity,Class<? extends Activity> activityKlazz) throws NoSuchMethodException, InvocationTargetException, IllegalAccessException { Method[] declaredMethods = activityKlazz.getDeclaredMethods(); for (Method declaredMethod : declaredMethods) { //获取标记了OnClick注解的方法 if (declaredMethod.isAnnotationPresent(OnClick.class)){ OnClick annotation = declaredMethod.getAnnotation(OnClick.class); int [] value = annotation.value(); //创建处理器类并且生成代理类，同时将activity中我们标记了OnClick的方法和处理器类绑定起来 OnClickListenerProxy proxy= new OnClickListenerProxy(); Object listener=proxy.bind(activity); proxy.bindEvent(declaredMethod); for ( int viewId : value) { Method findViewByIdMethod = activityKlazz.getMethod( "findViewById" , int.class); findViewByIdMethod.setAccessible( true ); View view = (View) findViewByIdMethod.invoke(activity, viewId); //通过反射把我们的代理类注入到相应view的onClickListener中 Method setOnClickListener = view.getClass().getMethod( "setOnClickListener" , View.OnClickListener.class); setOnClickListener.setAccessible( true ); setOnClickListener.invoke(view,listener); } } } } private static void injectView (Activity activity,Class<? extends Activity> activityKlazz) throws NoSuchMethodException, IllegalAccessException, InvocationTargetException { /** * 注入view其实很简单，通过反射拿到activity中标记了InjectView注解的field。 * 然后通过反射获取到findViewById方法，并且执行这个方法拿到view的实例 * 接着将实例赋值给activity里的field上 */ for (Field field : activityKlazz.getDeclaredFields()) { if (field.isAnnotationPresent(InjectView.class)) { InjectView annotation = field.getAnnotation(InjectView.class); int viewId = annotation.value(); Method findViewByIdMethod = activityKlazz.getMethod( "findViewById" , int.class); findViewByIdMethod.setAccessible( true ); View view = (View) findViewByIdMethod.invoke(activity, viewId); field.setAccessible( true ); field.set(activity, view); } } } } 复制代码`

上面的代码都有注释，下面看下OnClickListenerProxy这个处理器类:

` public class OnClickListenerProxy implements InvocationHandler { static final String TAG= "OnClickListenerProxy" ; //这个其实就是我们相关的activity Object delegate; //这个是activity中我们标记了OnClick的方法，最终的操作就是把对OnClickListener中OnClick方法的调用替换成对这个event方法的调用 Method event; public Object bind (Object delegate) { this.delegate=delegate; //生成代理类，这个没什么好说的了 return Proxy.newProxyInstance( this.delegate.getClass().getClassLoader(), new Class[]{View.OnClickListener.class}, this ); } @Override public Object invoke (Object proxy, Method method, Object[] args) throws Throwable { Log.e(TAG, "invoke" ); Log.e(TAG, "method.name:" +method.getName()+ " args:" +args); //判断调用的是onClick方法的话，替换成对我们event方法的调用。 if ( "onClick".equals(method.getName())){ for (Object arg : args) { Log.e(TAG, "arg:" +arg); } View view= (View) args[ 0 ]; return event.invoke(delegate,view); } return method.invoke(delegate,args); } public void bindEvent (Method declaredMethod) { event=declaredMethod; } } 复制代码`

通过上面的代码实现，我们已经能够将项目运行并且成功实现了我们需要的功能。当然很有很多细节需要处理，不过作为展示学习已经足够用了，而且通过我们自己的方法也实现了黄油刀这种大名鼎鼎的开源库的部分功能是不是很Cool呢？

不过博主并不建议再自己项目中使用这些，大家还是使用黄油刀吧，因为反射和动态代理都会带来一部分的性能损耗，而黄油刀使用的是编译时注解的形式实现上面的功能的，在运行时不会带来性能的损耗，感兴趣的小伙伴们可以去Google相关的文章。