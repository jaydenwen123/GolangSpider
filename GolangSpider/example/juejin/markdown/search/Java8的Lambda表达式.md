# Java8的Lambda表达式 #

![è·ä¸Java8 - äºè§£lambda](https://user-gold-cdn.xitu.io/2019/3/30/169cd384814cce46?imageView2/0/w/1280/h/960/ignore-error/1)

从Java8出现以来 **lambda** 是最重要的特性之一，它可以让我们用简洁流畅的代码完成一个功能。

很长一段时间java被吐槽是冗余和缺乏函数式编程能力的语言，随着函数式编程的流行java8种也引入了这种编程风格。在此之前我们都在写匿名内部类干这些事，但有时候这不是好的做法，本文中将介绍和使用 **lambda** ，带你体验函数式编程的魔力。

## 什么是lambda? ##

lambda表达式是一段可以传递的代码，它的核心思想是将面向对象中的传递数据变成传递行为。
我们回顾一下在使用java8之前要做的事，之前我们编写一个线程时是这样的：

` Runnable r = new Runnable () { @Override public void run () { System.out.println( "do something." ); } } 复制代码`

也有人会写一个类去实现 **Runnable** 接口，这样做没有问题，我们注意这个接口中只有一个run方法，
当把 **Runnable** 对象给 **Thread** 对象作为构造参数时创建一个线程，运行后将输出do something.。
我们使用匿名内部类的方式实现了该方法。

> 
> 这实际上是一个代码即数据的例子，在run方法中是线程要执行的一个任务，但上面的代码中任务内容已经被规定死了。
> 当我们有多个不同的任务时，需要重复编写如上代码。

设计匿名内部类的目的，就是为了方便 Java 程序员将代码作为数据传递。不过，匿名内部 类还是不够简便。
为了执行一个简单的任务逻辑，不得不加上 6 行冗繁的样板代码。那如果是lambda该怎么做?

` Runnable r = () -> System.out.println( "do something." ); 复制代码`

嗯，这代码看起来很酷，你可以看到我们用 **()** 和 **->** 的方式完成了这件事，这是一个没有名字的函数，也没有人和参数，再简单不过了。
使用 **->** 将参数和实现逻辑分离，当运行这个线程的时候执行的是 **->** 之后的代码片段，且编译器帮助我们做了类型推导；
这个代码片段可以是用 **{}** 包含的一段逻辑。下面一起来学习一下lambda的语法。

## 基础语法 ##

在 **lambda** 中我们遵循如下的表达式来编写：

` expression = (variable) -> action 复制代码`

* **variable** : 这是一个变量,一个占位符。像x,y,z,可以是多个变量。
* **action** : 这里我称它为action, 这是我们实现的代码逻辑部分,它可以是一行代码也可以是一个代码片段

可以看到Java中lambda表达式的格式：参数、箭头、以及动作实现，当一个动作实现无法用一行代码完成，可以编写
一段代码用 **{}** 包裹起来。

lambda表达式可以包含多个参数,例如：

` BinaryOperator<Integer> sumFunction = (x, y) -> x + y; int sum = sumFunction.apply(1, 2); 复制代码`

这时候我们应该思考这段代码不是之前的x和y数字相加，而是创建了一个函数，用来计算两个操作数的和。
后面用int类型进行接收，在lambda中为我们省略去了 **return** 。

## 函数式接口 ##

函数式接口是只有一个方法的接口，用作lambda表达式的类型。前面写的例子就是一个函数式接口，来看看jdk中的 **Runnable** 源码

` @FunctionalInterface public interface Runnable { /** * When an object implementing interface <code>Runnable</code> is used * to create a thread, starting the thread causes the object 's * <code>run</code> method to be called in that separately executing * thread. * <p> * The general contract of the method <code>run</code> is that it may * take any action whatsoever. * * @see java.lang.Thread#run() */ public abstract void run(); } 复制代码`

这里只有一个抽象方法run，实际上你不写 **public abstract** 也是可以的，在接口中定义的方法都是 **public abstract** 的。
同时也使用注解 **@FunctionalInterface** 告诉编译器这是一个函数式接口，当然你不这么写也可以，标识后明确了这个函数中
只有一个抽象方法，当你尝试在接口中编写多个方法的时候编译器将不允许这么干。

## 尝试函数式接口 ##

我们来编写一个函数式接口，输入一个年龄，判断这个人是否是成人。

` public class FunctionInterfaceDemo { @FunctionalInterface interface Predicate<T> { boolean test (T t); } /** * 执行Predicate判断 * * @param age 年龄 * @param predicate Predicate函数式接口 * @ return 返回布尔类型结果 */ public static boolean do Predicate(int age, Predicate<Integer> predicate) { return predicate.test(age); } public static void main(String[] args) { boolean isAdult = do Predicate(20, x -> x >= 18); System.out.println(isAdult); } } 复制代码`

从这个例子我们很轻松的完成 **是否是成人** 的动作，其次判断是否是成人，在此之前我们的做法一般是编写一个
判断是否是成人的方法，是无法将 **判断** 共用的。而在本例只，你要做的是将 **行为** (判断是否是成人，或者是判断是否大于30岁)
传递进去，函数式接口告诉你结果是什么。

实际上诸如上述例子中的接口，伟大的jdk设计者为我们准备了 **java.util.function** 包

![](https://user-gold-cdn.xitu.io/2019/3/30/169cd389114a3808?imageView2/0/w/1280/h/960/ignore-error/1)

我们前面写的Predicate函数式接口也是JDK种的一个实现，他们大致分为以下几类：

![](https://user-gold-cdn.xitu.io/2019/3/30/169cd38915e88988?imageView2/0/w/1280/h/960/ignore-error/1)

**消费型接口示例**

` public static void donation(Integer money, Consumer<Integer> consumer){ consumer.accept(money); } public static void main(String[] args) { donation(1000, money -> System.out.println( "好心的麦乐迪为Blade捐赠了" +money+ "元" )) ; } 复制代码`

**供给型接口示例**

` public static List<Integer> supply(Integer num, Supplier<Integer> supplier){ List<Integer> resultList = new ArrayList<Integer>() ; for (int x=0;x<num;x++) resultList.add(supplier.get()); return resultList ; } public static void main(String[] args) { List<Integer> list = supply(10,() -> (int)(Math.random()*100)); list.forEach(System.out::println); } 复制代码`

**函数型接口示例**

转换字符串为 **Integer**

` public static Integer convert(String str, Function<String, Integer> function ) { return function.apply(str); } public static void main(String[] args) { Integer value = convert( "28" , x -> Integer.parseInt(x)); } 复制代码`

**断言型接口示例**

筛选出只有2个字的水果

` public static List<String> filter(List<String> fruit, Predicate<String> predicate){ List<String> f = new ArrayList<>(); for (String s : fruit) { if (predicate.test(s)){ f.add(s); } } return f; } public static void main(String[] args) { List<String> fruit = Arrays.asList( "香蕉" , "哈密瓜" , "榴莲" , "火龙果" , "水蜜桃" ); List<String> newFruit = filter(fruit, (f) -> f.length() == 2); System.out.println(newFruit); } 复制代码`

## 默认方法 ##

在Java语言中，一个接口中定义的方法必须由实现类提供实现。但是当接口中加入新的API时，
实现类按照约定也要修改实现，而Java8的API对现有接口也添加了很多方法，比如List接口中添加了 **sort** 方法。
如果按照之前的做法，那么所有的实现类都要实现sort方法，JDK的编写者们一定非常抓狂。

幸运的是我们使用了Java8，这一问题将得到很好的解决，在Java8种引入新的机制， **支持在接口中声明方法同时提供实现** 。
这令人激动不已，你有两种方式完成 1.在接口内声明静态方法 2.指定一个默认方法。

我们来看看在JDK8中上述 **List** 接口添加方法的问题是如何解决的

` default void sort(Comparator<? super E> c) { Object[] a = this.toArray(); Arrays.sort(a, (Comparator) c); ListIterator<E> i = this.listIterator(); for (Object e : a) { i.next(); i.set((E) e); } } 复制代码`

翻阅 **List** 接口的源码，其中加入一个默认方法 **default void sort(Comparator<? super E> c)** 。
在返回值之前加入 **default** 关键字，有了这个方法我们可以直接调用sort方法进行排序。

` List<Integer> list = Arrays.asList(2, 7, 3, 1, 8, 6, 4); list.sort(Comparator.naturalOrder()); System.out.println(list); 复制代码`

**Comparator.naturalOrder()** 是一个自然排序的实现，这里可以自定义排序方案。
你经常看到使用Java8操作集合的时候可以直接foreach的原因也是在 **Iterable** 接口中也新增了一个默认方法： **forEach** ，
该方法功能和 for 循环类似，但是允许 用户使用一个Lambda表达式作为循环体。

## 福利部分： ##

**[《大数据成神之路》]( https://link.juejin.im/?target=https%3A%2F%2Fshimo.im%2Fdocs%2FjdPhrtFwVCAMkoWv%2F )**

**《几百TJava和大数据资源下载》** ( https://link.juejin.im/?target=https%3A%2F%2Fshimo.im%2Fdocs%2FsSTmRlbM4FUbUwX4%2F )

![](https://user-gold-cdn.xitu.io/2019/3/26/169b5973fef05422?imageView2/0/w/1280/h/960/ignore-error/1)