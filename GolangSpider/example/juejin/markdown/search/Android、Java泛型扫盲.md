# Android、Java泛型扫盲 #

首先我们定义A、B、C、D四个类，他们的关系如下

` class A {} class B extends A {} class C extends B {} class D extends C {} 复制代码`

### 不指明泛型类型 ###

` //以下代码均编译通过 List list = new ArrayList(); //不指明泛型类型，泛型默认为Object类型，故能往里面添加任意实例对象 list.add( new A()); list.add( new B()); list.add( new C()); //取出则默认为Object类型 Object o = list.get( 0 ); 复制代码`

这个好理解，因为所有的类都继承与Object，故能往list里面添加任意实例对象

### 无边界通配符 ` ？` ###

首先我们要明白一个概念，通配符 ` ？` 意义就是它是一个未知的符号，可以是代表任意的类。

` //我们发现，这样写编译不通过，原因很简单，泛型不匹配，虽然B继承A List<A> listA = new ArrayList<B>(); //以下5行代码均编译通过 List<?> list; list = new ArrayList<A>(); list = new ArrayList<B>(); list = new ArrayList<C>(); list = new ArrayList<D>(); Object o = list.get( 0 ); //编译通过 list.add( new A()); //编译不通过 list.add( new B()); //编译不通过 list.add( new C()); //编译不通过 list.add( new D()); //编译不通过 复制代码`

` 知识点`

* 无边界通配符 ` ？` 能取不能存。这个好理解，因为编译器不知道 ` ?` 具体是啥类型，故不能存；但是任意类型都继承于Object，故能取，但取出默认为Object对象。

### 上边界符 ` ？extends` ###

继续上代码

` List<? extends C> listC; listC = new ArrayList<A>(); //编译不通过 listC = new ArrayList<B>(); //编译不通过 listC = new ArrayList<C>(); //编译通过 listC = new ArrayList<D>(); //编译通过 C c = listC.get( 0 ); //编译通过 listC.add( new C()); //编译不通过 listC.add( new D()); //编译不通过 复制代码`

` 知识点` :

* 上边界符 ` ? extends` 只是限定了赋值给它的实例类型(这里为赋值给listC的实例类型)，且边界包括自身。
* 上边界符 ` ? extends` 跟 ` ？` 一样能取不能存，道理是一样的，虽然限定了上边界，但编译器依然不知道 ` ?` 是啥类型，故不能存；但是限定了上边界，故取出来的对象类型默认为上边界的类型

### 下边界符 ` ？super` ###

` List<? super B> listB; listB = new ArrayList<A>(); //编译通过 listB = new ArrayList<B>(); //编译通过 listB = new ArrayList<C>(); //编译不通过 listB = new ArrayList<D>(); //编译不通过 Object o = listB.get( 0 ); //编译通过 listB.add( new A()); //编译不通过 listB.add( new B()); //编译通过 listB.add( new C()); //编译通过 listB.add( new D()); //编译通过 复制代码`

` 知识点`

* 下边界符 ` ？super` ，跟上边界符一样，只是限定了赋值给它的实例类型，也包括边界自身
* 下边界符 ` ？super` 能存能取，因为设定了下边界，故我们能存下边界以下的类型，当然也包括边界自身；然而取得时候编译器依然不知道 ` ?` 具体是什么类型，故取出默认为Object类型。

### 类型擦除 ###

首先我们要明白一点：Java 的泛型在编译期有效，在运行期会被删除 我们来看一段代码

` //这两个方法写在同一个类里 public void list (List<A> listA) {} public void list (List<B> listB) {} 复制代码`

上面的代码会有问题吗？显然是有的，编译器报错，提示如下信息： ` list(List<A>) clashed with list(List<B>) ; both methods have same erasure` 翻译过来就是，在类型擦除后，两个方法具有相同的签名，我们来看看类型擦除后是什么样子

` public void list (List listA) {} public void list (List listB) {} 复制代码`

可以看出，两个方法签名完全一致，故编译不通过。 明白了类型擦除，我们还需要明白一个概念

* 泛型类并没有自己独有的Class类对象

比如并不存在List<A>.class或是List<B>.class，而只有List.class 接下来这个案例就好理解了

` List<A> listA = new ArrayList<A>(); List<B> listB = new ArrayList<B>(); System.out.println(listA.getClass() == listB.getClass()); //输出true 复制代码`

### 泛型传递 ###

现实开发中，我们经常会用到泛型传递，例如我们经常需要对Http请求返回的结果做反序列化操作

` public static <T> T fromJson (String result, Class<T> type) { try { return new Gson().fromJson(result, type); } catch (Exception ignore) { return null ; } } 复制代码`

此时我们传进去是什么类型，就会返回自动该类型的对象

` String result= "xxx" ; A a = fromJson(result, A.class); B b = fromJson(result, B.class); C c = fromJson(result, C.class); D d = fromJson(result, D.class); Integer integer = fromJson(result, Integer.class); String str = fromJson(result, String.class); Boolean boo = fromJson(result, Boolean.class); 复制代码`

那如果我们想返回一个集合呢，如 ` List<A>` ，下面这样明显是不对的。

` //编译报错，前面类型擦除时，我们讲过，不存List<A>.class这种类型 ArrayList<A> list = fromJson(result, ArrayList<A>.class)； 复制代码`

那我们该怎么做呢？首先，我们对 ` fromJson` 改造一下，如下：

` //type为一个数组类型 public static <T> List<T> fromJson (String result, Class<T[]> type) { try { T[] arr = new Gson().fromJson(result, type); //首先拿到数组 return Arrays.asList(arr); //数组转集合 } catch (Exception ignore) { return null ; } } 复制代码`

这个时候我们就可以这么做了

` String result= "xxx" ; List<A> listA = fromJson(result, A[].class); List<B> listB = fromJson(result, B[].class); List<C> listC = fromJson(result, C[].class); List<D> listD = fromJson(result, D[].class); List<Integer> listInt = fromJson(result, Integer[].class); List<String> listStr = fromJson(result, String[].class); List<Boolean> listBoo = fromJson(result, Boolean[].class); 复制代码`

ok，我在再来，相信大多数Http接口返回的数据格式是这样的：

` public class Response < T > { private T data; private int code; private String msg; //省略get/set方法 } 复制代码`

那这种我们又该如何传递呢？显然用前面的两个 ` fromJson` 方法都行不通，我们再来改造一下，如下:

` //这里我们直接传递一个Type类型 public static <T> T fromJson (String result, Type type) { try { return new Gson().fromJson(result, type); } catch (Exception ignore) { return null ; } } 复制代码`

这个Type是什么鬼？点进去看看

` public interface Type { default String getTypeName () { return toString(); } } 复制代码`

哦，原来就是一个接口，并且只有一个方法，我们再来看看它的实现类

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/4/22/16a43a5ab36d5422?imageView2/0/w/1280/h/960/ignore-error/1) 发现有5个实现类，其中4个是接口，另外一个是Class类，我们再来看看Class类的声明

` public final class Class < T > implements java. io. Serializable , GenericDeclaration , Type , AnnotatedElement { //省略内部代码 } 复制代码`

现在有没有明白点，现在我们重点来关注下 ` Type` 接口的其中一个实现接口 ` ParameterizedType` ，我们来看下它的内部代码，里面就只有3个方法

` public interface ParameterizedType extends Type { /** * 例如: * List<String> list; 则返回 {String.class} * Map<String,Long> map; 则返回 {String.class,Long.class} * Map.Entry<String,Long> entry; 则返回 {String.class,Long.class} * * @return 以数组的形式返回所有的泛型类型 */ Type[] getActualTypeArguments(); /** * 例如: * List<String> list; 则返回 List.class * Map<String,Long> map; 则返回 Map.class * Map.Entry<String,Long> entry; 则返回 Entry.class * * @return 返回泛型类的真实类型 */ Type getRawType () ; /** * 例如: * List<String> list; 则返回 null * Map<String,Long> map; 则返回 null * Map.Entry<String,Long> entry; 则返回 Map.class * * @return 返回泛型类持有者的类型，这里可以简单理解为返回外部类的类型，如果没有外部类，则返回null */ Type getOwnerType () ; } 复制代码`

顾名思义， ` ParameterizedType` 代表一个参数化类型。

这个时候我们来自定义一个类，并实现ParameterizedType接口，如下：

` public class ParameterizedTypeImpl implements ParameterizedType { private Type rawType; //真实类型 private Type actualType; //泛型类型 public ParameterizedTypeImpl (Type rawType,Type actualType) { this.rawType = rawType; this.actualType = actualType; } public Type[] getActualTypeArguments() { return new Type[]{actualType}; } public Type getRawType () { return rawType; } public Type getOwnerType () { return null ; } } 复制代码`

我们再次贴出 ` fromJson` 方法

` //这里我们直接传递一个Type类型 public static <T> T fromJson (String result, Type type) { try { return new Gson().fromJson(result, type); } catch (Exception ignore) { return null ; } } 复制代码`

此时我们想得到 ` Response<T>` 对象，就可以这样写

` Response<A> responseA = fromJson(result, new ParameterizedTypeImpl(Response.class, A.class)); Response<B> responseB = fromJson(result, new ParameterizedTypeImpl(Response.class, B.class)); Response<C> responseC = fromJson(result, new ParameterizedTypeImpl(Response.class, C.class)); 复制代码`

想得到 ` List<T>` 对象，也可以通过 ` ParameterizedTypeImpl` 得到，如下:

` List<A> listA = fromJson(result, new ParameterizedTypeImpl(List.class, A.class)); List<B> listB = fromJson(result, new ParameterizedTypeImpl(List.class, B.class)); List<C> listC = fromJson(result, new ParameterizedTypeImpl(List.class, C.class)); 复制代码`

然而，如果我们想得到 ` Response<List<T>>` 对象，又该如何得到呢？ ` ParameterizedTypeImpl` 一样能够实现，如下：

` //第一步，创建List<T>对象对应的Type类型 Type listAType = new ParameterizedTypeImpl(List.class, A.class); Type listBType = new ParameterizedTypeImpl(List.class, B.class); Type listCType = new ParameterizedTypeImpl(List.class, C.class); //第二步，创建Response<List<T>>对象对应的Type类型 Type responseListAType = new ParameterizedTypeImpl(Response.class, listAType); Type responseListBType = new ParameterizedTypeImpl(Response.class, listBType); Type responseListCType = new ParameterizedTypeImpl(Response.class, listCType); //第三步，通过Type对象，获取对应的Response<List<T>>对象 Response<List<A>> responseListA = fromJson(result, responseListAType); Response<List<B>> responseListB = fromJson(result, responseListBType); Response<List<C>> responseListC = fromJson(result, responseListCType); 复制代码`

然后，能不能再简单一点呢？可以，我们对 ` ParameterizedTypeImpl` 改造一下

` /** * User: ljx * Date: 2018/10/23 * Time: 09:36 */ public class ParameterizedTypeImpl implements ParameterizedType { private final Type rawType; private final Type ownerType; private final Type[] actualTypeArguments; //适用于单个泛型参数的类 public ParameterizedTypeImpl (Type rawType, Type actualType) { this ( null , rawType, actualType); } //适用于多个泛型参数的类 public ParameterizedTypeImpl (Type ownerType, Type rawType, Type... actualTypeArguments) { this.rawType = rawType; this.ownerType = ownerType; this.actualTypeArguments = actualTypeArguments; } /** * 本方法仅使用于单个泛型参数的类 * 根据types数组，确定具体的泛型类型 * List<List<String>> 对应 get(List.class, List.class, String.class) * * @param types Type数组 * @return ParameterizedTypeImpl */ public static ParameterizedTypeImpl get (@NonNull Type rawType, @NonNull Type... types) { final int length = types.length; if (length > 1 ) { Type parameterizedType = new ParameterizedTypeImpl(types[length - 2 ], types[length - 1 ]); Type[] newTypes = Arrays.copyOf(types, length - 1 ); newTypes[newTypes.length - 1 ] = parameterizedType; return get(rawType, newTypes); } return new ParameterizedTypeImpl(rawType, types[ 0 ]); } //适用于多个泛型参数的类 public static ParameterizedTypeImpl getParameterized (@NonNull Type rawType, @NonNull Type... actualTypeArguments) { return new ParameterizedTypeImpl( null , rawType, actualTypeArguments); } public final Type[] getActualTypeArguments() { return actualTypeArguments; } public final Type getOwnerType () { return ownerType; } public final Type getRawType () { return rawType; } } 复制代码`

此时，我们就可以这样写

` //第一步，直接创建Response<List<T>>对象对应的Type类型 Type responseListAType = ParameterizedTypeImpl.get(Response.class, List.class, A.class); Type responseListBType = ParameterizedTypeImpl.get(Response.class, List.class, B.class) Type responseListCType = ParameterizedTypeImpl.get(Response.class, List.class, C.class) //第二步，通过Type对象，获取对应的Response<List<T>>对象 Response<List<A>> responseListA = fromJson(result, responseListAType); Response<List<B>> responseListB = fromJson(result, responseListBType); Response<List<C>> responseListC = fromJson(result, responseListCType); 复制代码`

现实开发中，我们还可能遇到这样的数据结构

` { "code" : 0 , "msg" : "" , "data" : { "totalPage" : 0 , "list" : [] } } 复制代码`

此时， ` Response<T>` 里面的泛型传List肯定是不能正常解析的，我们需要再定一个类

` public class PageList < T > { private int totalPage; private List<T> list; //省略get/set方法 } 复制代码`

此时就可以这样解析数据

` //第一步，直接创建Response<PageList<T>>对象对应的Type类型 Type responsePageListAType = ParameterizedTypeImpl.get(Response.class, PageList.class, A.class); Type responsePageListBType = ParameterizedTypeImpl.get(Response.class, PageList.class, B.class) Type responsePageListCType = ParameterizedTypeImpl.get(Response.class, PageList.class, C.class) //第二步，通过Type对象，获取对应的Response<PageList<T>>对象 Response<PageList<A>> responsePageListA = fromJson(result, responsePageListAType); Response<PageList<B>> responsePageListB = fromJson(result, responsePageListBType); Response<PageList<C>> responsePageListC = fromJson(result, responsePageListCType); 复制代码`

注： ` ParameterizedTypeImpl get(Type... types)` 仅仅适用于单个泛型参数的时候，如Map等，有两个泛型参数以上的不要用此方法获取Type类型。如果需要获取Map等两个泛型参数以上的Type类型。可调用 ` getParameterized(@NonNull Type rawType, @NonNull Type... actualTypeArguments)` 构造方法获取，如：

` //获取 Map<String,String> 对应的Type类型 Type mapType = ParameterizedTypeImpl.getParameterized(Map.class, String.classs, String.class) //获取 Map<A,B> 对应的Type类型 Type mapType = ParameterizedTypeImpl.getParameterized(Map.class, A.classs, B.class) 复制代码`

到这，泛型相关知识点讲解完毕，如有疑问，请留言。

感兴趣的同学，可以查看我的另一片文章 [RxHttp 一条链发送请求，新一代Http请求神器（一））]( https://juejin.im/post/5cbd267fe51d456e2b15f623 ) 里面就用到了 ` ParameterizedTypeImpl` 类进行泛型传递。