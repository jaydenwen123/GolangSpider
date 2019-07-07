# 在Java中如何优雅地判空 #

### 判空灾难 ###

![NullPointerException](https://user-gold-cdn.xitu.io/2018/11/22/1673927481752990?imageView2/0/w/1280/h/960/ignore-error/1)

作为搬砖党的一族们，我们对判空一定再熟悉不过了，不要跟我说你很少进行判空，除非你喜欢NullPointerException。

不过NullPointerException对于很多猿们来说，也是Exception家族中最亲近的一员了。

![Wowo](https://user-gold-cdn.xitu.io/2018/11/22/1673927481ca39f4?imageView2/0/w/1280/h/960/ignore-error/1)

为了避免NullPointerException来找我们，我们经常会进行如下操作。

` if (data != null ) { do sth. } 复制代码`

如果一个类中多次使用某个对象，那你可能要一顿操作，so:

![1](https://user-gold-cdn.xitu.io/2018/11/22/16739274814959c2?imageView2/0/w/1280/h/960/ignore-error/1)

“世界第九大奇迹”就这样诞生了。Maybe你会想，项目中肯定不止你一个人会这样一顿操作，然后按下Command+Shift+F，真相就在眼前：

![2](https://user-gold-cdn.xitu.io/2018/11/22/16739274816e59b9?imageView2/0/w/1280/h/960/ignore-error/1)

What，我们有接近一万行的代码都是在判空？

![3](https://user-gold-cdn.xitu.io/2018/11/22/1673927481a2e0dc?imageView2/0/w/1280/h/960/ignore-error/1)

好了，接下来，要进入正题了。

### NullObject模式 ###

对于项目中无数次的判空，对代码质量整洁度产生了十分之恶劣的影响，对于这种现象，我们称之为“判空灾难”。

那么，这种现象如何治理呢，你可能听说过NullObject模式，不过这不是我们今天的武器，但是还是需要介绍一下NullObject模式。

什么是NullObject模式呢？

> 
> 
> 
> In object-oriented computer programming, a null object is an object with
> no referenced value or with defined neutral ("null") behavior. The null
> object design pattern describes the uses of such objects and their
> behavior (or lack thereof).
> 
> 

以上解析来自Wikipedia。

NullObject模式首次发表在“ 程序设计模式语言 ”系列丛书中。一般的，在面向对象语言中，对对象的调用前需要使用判空检查，来判断这些对象是否为空，因为在空引用上无法调用所需方法。

空对象模式的一种典型实现方式如下图所示(图片来自网络)：

![4](https://user-gold-cdn.xitu.io/2018/11/22/1673927481958558?imageView2/0/w/1280/h/960/ignore-error/1)

示例代码如下（命名来自网络，哈哈到底是有多懒）：

Nullable是空对象的相关操作接口，用于确定对象是否为空，因为在空对象模式中，对象为空会被包装成一个Object，成为Null Object，该对象会对原有对象的所有方法进行空实现。。

` public interface Nullable { boolean isNull () ; } 复制代码`

这个接口定义了业务对象的行为。

` public interface DependencyBase extends Nullable { void Operation () ; } 复制代码`

这是该对象的真实类，实现了业务行为接口DependencyBase与空对象操作接口Nullable。

` public class Dependency implements DependencyBase , Nullable { @Override public void Operation () { System.out.print( "Test!" ); } @Override public boolean isNull () { return false ; } } 复制代码`

这是空对象，对原有对象的行为进行了空实现。

` public class NullObject implements DependencyBase { @Override public void Operation () { // do nothing } @Override public boolean isNull () { return true ; } } 复制代码`

在使用时，可以通过工厂调用方式来进行空对象的调用，也可以通过其他如反射的方式对对象进行调用（一般多耗时几毫秒）在此不进行详细叙述。

` public class Factory { public static DependencyBase get (Nullable dependencyBase) { if (dependencyBase == null ){ return new NullObject(); } return new Dependency(); } } 复制代码`

这是一个使用范例，通过这种模式，我们不再需要进行对象的判空操作，而是可以直接使用对象，也不必担心NPE（NullPointerException）的问题。

` public class Client { public void test (DependencyBase dependencyBase) { Factory.get(dependencyBase).Operation(); } } 复制代码`

关于空对象模式，更具体的内容大家也可以多找一找资料，上述只是对NullObject的简单介绍，但是，今天我要推荐的是一款协助判空的插件NR Null Object，让我们来优雅地进行判空，不再进行一顿操作来定义繁琐的空对象接口与空独享实现类。

###.NR Null Object ###

NR Null Object是一款适用于Android Studio、IntelliJ IDEA、PhpStorm、WebStorm、PyCharm、RubyMine、AppCode、CLion、GoLand、DataGrip等IDEA的Intellij插件。其可以根据现有对象，便捷快速生成其空对象模式需要的组成成分，其包含功能如下：

* 分析所选类可声明为接口的方法；
* 抽象出公有接口；
* 创建空对象，自动实现公有接口；
* 对部分函数进行可为空声明；
* 可追加函数进行再次生成；
* 自动的函数命名规范

让我们来看一个使用范例：

![5](https://user-gold-cdn.xitu.io/2018/11/22/167392749cc2f3d9?imageslim)

怎么样，看起来是不是非常快速便捷，只需要在原有需要进行多次判空的对象所属类中，右键弹出菜单，选择Generate，并选择NR Null Object即可自动生成相应的空对象组件。

那么如何来获得这款插件呢？

### 安装方式 ###

可以直接通过IDEA的Preferences中的Plugins仓库进行安装。

选择 Preferences → Plugins → Browse repositories

![6](https://user-gold-cdn.xitu.io/2018/11/22/16739274a628e941?imageView2/0/w/1280/h/960/ignore-error/1)

搜索“NR Null Oject”或者“Null Oject”进行模糊查询，点击右侧的Install，restart IDEA即可。

![7](https://user-gold-cdn.xitu.io/2018/11/22/16739274ad65296e?imageView2/0/w/1280/h/960/ignore-error/1)

### Optional ###

感谢评论区小伙伴们的积极补充。   关于优雅地判空，还有一种方式是使用Java8特性/Guava中的Optional来进行优雅地判空，Optional来自官方的介绍如下：

> 
> 
> 
> A container object which may or may not contain a non-null value. If a
> value is present, ` isPresent()` will return ` true` and ` get()` will
> return the value.
> 
> 

一个可能包含也可能不包含非null值的容器对象。 如果存在值，isPresent（）将返回true，get（）将返回该值。

话不多说，举个例子。

![栗子](https://user-gold-cdn.xitu.io/2018/11/24/1674372c23ac07e3?imageView2/0/w/1280/h/960/ignore-error/1)

有如下代码，需要获得Test2中的Info信息，但是参数为Test4，我们要一层层的申请，每一层都获得的对象都可能是空，最后的代码看起来就像这样。

` public String testSimple (Test4 test) { if (test == null ) { return "" ; } if (test.getTest3() == null ) { return "" ; } if (test.getTest3().getTest2() == null ) { return "" ; } if (test.getTest3().getTest2().getInfo() == null ) { return "" ; } return test.getTest3().getTest2().getInfo(); } 复制代码`

但是使用Optional后，整个就都不一样了。

` public String testOptional (Test test) { return Optional.ofNullable(test).flatMap(Test::getTest3) .flatMap(Test3::getTest2) .map(Test2::getInfo) .orElse( "" ); } 复制代码`

1.Optional.ofNullable(test)，如果test为空，则返回一个单例空Optional对象，如果非空则返回一个Optional包装对象，Optional将test包装；

` public static <T> Optional<T> ofNullable (T value) { return value == null ? empty() : of(value); } 复制代码`

2.flatMap(Test::getTest3)判断test是否为空，如果为空，继续返回第一步中的单例Optional对象，否则调用Test的getTest3方法；

` public <U> Optional<U> flatMap (Function<? super T, Optional<U>> mapper) { Objects.requireNonNull(mapper); if (!isPresent()) return empty(); else { return Objects.requireNonNull(mapper.apply(value)); } } 复制代码`

3.flatMap(Test3::getTest2)同上调用Test3的getTest2方法；

4.map(Test2::getInfo)同flatMap类似，但是flatMap要求Test3::getTest2返回值为Optional类型，而map不需要，flatMap不会多层包装，map返回会再次包装Optional；

` public <U> Optional<U> map (Function<? super T, ? extends U> mapper) { Objects.requireNonNull(mapper); if (!isPresent()) return empty(); else { return Optional.ofNullable(mapper.apply(value)); } } 复制代码`

5.orElse("");获得map中的value，不为空则直接返回value，为空则返回传入的参数作为默认值。

` public T orElse (T other) { return value != null ? value : other; } 复制代码`

怎么样，使用Optional后我们的代码是不是瞬间变得非常整洁，或许看到这段代码你会有很多疑问，针对复杂的一长串判空，Optional有它的优势，但是对于简单的判空使用Optional也会增加代码的阅读成本、编码量以及团队新成员的学习成本。毕竟Optional在现在还并没有像RxJava那样流行，它还拥有一定的局限性。

如果直接使用Java8中的Optional，需要保证安卓API级别在24及以上。

![image-20181124085913887.png](https://user-gold-cdn.xitu.io/2018/11/24/1674372c23ad0ab5?imageView2/0/w/1280/h/960/ignore-error/1)

你也可以直接引入Google的Guava。（啥是Guava？来自官方的提示）

> 
> 
> 
> Guava is a set of core libraries that includes new collection types (such
> as multimap and multiset), immutable collections, a graph library,
> functional types, an in-memory cache, and APIs/utilities for concurrency,
> I/O, hashing, primitives, reflection, string processing, and much more!
> 
> 

引用方式，就像这样：

` dependencies { compile 'com.google.guava:guava:27.0-jre' // or, for Android: api 'com.google.guava:guava:27.0-android' } 复制代码`

不过IDEA默认会显示黄色，提示让你将Guava表达式迁移到Java Api上。

![image-3.png](https://user-gold-cdn.xitu.io/2018/11/24/1674372c23dcfedd?imageView2/0/w/1280/h/960/ignore-error/1)

当然，你也可以通过在Preferences搜索"Guava"来Kill掉这个Yellow的提示。

![image-4.png](https://user-gold-cdn.xitu.io/2018/11/24/1674372c23fe15f3?imageView2/0/w/1280/h/960/ignore-error/1)

关于Optional使用还有很多技巧，感兴趣可以查阅Guava和Java8相关书籍和文档。

使用Optional具有如下优点：

* 将防御式编程代码完美包装
* 链式调用
* 有效避免程序代码中的空指针

但是也同样具有一些缺点：

* 流行性不是非常理想，团队新成员需要学习成本
* 安卓中需要引入Guava，需要团队每个人处理IDEA默认提示，或者忍受黄色提示
* 有时候代码阅读看起来可能会如下图所示：

![Duang](https://user-gold-cdn.xitu.io/2018/11/24/1674372c44b3d866?imageView2/0/w/1280/h/960/ignore-error/1)

### Kotlin ###

当然，Kotlin以具有优秀的空安全性为一大特色，并可以与Java很好的混合使用，like this:

` test1?.test2?.test3?.test4 复制代码`

如果你已经开始使用了Kotlin，可以不用再写缭乱的防御判空语句。如果你还没有使用Kotlin，并不推荐为了判空优雅而直接转向Kotlin。