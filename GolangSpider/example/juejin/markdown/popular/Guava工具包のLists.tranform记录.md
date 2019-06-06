# Guava工具包のLists.tranform记录 #

![](https://user-gold-cdn.xitu.io/2019/5/25/16aedecc6ae2b6a1?imageView2/0/w/1280/h/960/ignore-error/1)

浪完杰伦演唱会，回来好好学习😝

**[个人博客项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FVip-Augus%2FVip-Augus.github.io )**

希望各位帮忙点个star，给我加个小星星✨

## 引言 ##

Guava工具包是Google推出的Java工具包，想要完整学习的话，推荐去并发编程网(ifeve)进行学习， [传送门]( https://link.juejin.im?target=http%3A%2F%2Fifeve.com%2Fgoogle-guava%2F ) 。

我们在开发过程中，会经常使用到java.util.Collections这个Java自带的工具类，Guava在这个基础上，提供了更多工具方法，而且很多是静态方法。

这篇文章记录的是自己使用Lists.tranform()方法遇到的一个小问题，如有错误请提出~~

## 常用场景和使用示例 ##

在写代码时，经常会要与外部系统进行联调，RPC获取到基础信息列表，例如用户信息List，我们需要在自己系统中根据用户ID列表进行一些信息查询，这样我们需要从Result列表里面拿出ID列表就足够了，不需要其他字段。简单来写，直接用foreach就行了，但使用Lists.tranform这个方法 **可以更加简单和简洁**

举个🌰：

` List<User> users = Lists.newArrayList( new User( 1 , "ceshi1" ), new User( 2 , "ceshi2" )); List<String> names = Lists.transform(users, new Function<User, String>() { @Nullable @Override public String apply (@Nullable User input) { return input.getName(); } }); 复制代码`

通过Lists.tranform这个方法，就能从User列表（用户列表）中取出String列表（姓名列表）。

## 遇到的问题 ##

还是上面那个例子，取出一个新列表后，对原来的list中的name进行了修改，新列表的name同样也会修改，为什么会出现这种情况呢？先来看看Lists.tranform源码注解

` /** * Returns a list that applies { @code function} to each element of { @code * fromList}. The returned list is a transformed view of { @code fromList}; * changes to { @code fromList} will be reflected in the returned list and vice * versa. 返回一个列表，该列表将{@ code函数}应用到{@ code from list}的每个元素。返回的列表是{@代码从列表}转换的视图;将fromList代码的更改将反映在返回的列表中，反之亦然 ··· **/ public static <F, T> List<T> transform ( List<F> fromList, Function<? super F, ? extends T> function) { return (fromList instanceof RandomAccess) ? new TransformingRandomAccessList<>(fromList, function) : new TransformingSequentialList<>(fromList, function); } 复制代码`

翻译了一下注解，大概意思就是，使用这个方法会 **返回一个列表** ，如果对源list或者返回的list进行修改后， **会同时影响到对方** ，同样可以看该静态方法返回的是一个内部静态类对象：

` private static class TransformingSequentialList < F , T > extends AbstractSequentialList < T > implements Serializable { final List<F> fromList; final Function<? super F, ? extends T> function; TransformingSequentialList(List<F> fromList, Function<? super F, ? extends T> function) { this.fromList = checkNotNull(fromList); this.function = checkNotNull(function); } @Override public ListIterator<T> listIterator ( final int index) { return new TransformedListIterator<F, T>(fromList.listIterator(index)) { @Override T transform (F from) { return function.apply(from); } }; } 复制代码`

可以看到最后一个方法 **listIterator（列表迭代器）** ，每次迭代都是从fromList中取出来的，下面放一个Debug动图，可以清楚看出，源list和新list的值，都是指向同一个内存地址， **所以无论是哪一方修改后，对方拿到的引用都是一致的，都是修改后的值** ：

![](https://user-gold-cdn.xitu.io/2019/5/25/16aedecc6afe44b2?imageslim)

所以如果下次对某一方修改后，记得对方列表也会相应变化，避免使用错误的值。