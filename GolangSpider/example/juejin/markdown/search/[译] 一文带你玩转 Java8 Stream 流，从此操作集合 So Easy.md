# [译] 一文带你玩转 Java8 Stream 流，从此操作集合 So Easy #

> 
> 
> 
> 本文翻译自 [winterbe.com/posts/2014/…](
> https://link.juejin.im?target=https%3A%2F%2Fwinterbe.com%2Fposts%2F2014%2F07%2F31%2Fjava8-stream-tutorial-examples%2F
> )
> 
> 
> 
> 作者: @Winterbe
> 
> 
> 
> 欢迎关注个人微信公众号: **小哈学Java** ，即可免费无套路领取10G面试学习资料哦，文末资料截图。
> 
> 
> 
> 个人网站: [www.exception.site/java8/java8…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.exception.site%2Fjava8%2Fjava8-stream-tutorial
> )
> 
> 

` Stream` 流可以说是 Java8 新特性中用起来最爽的一个功能了，有了它，从此操作集合告别繁琐的 ` for` 循环。但是还有很多小伙伴对 Stream 流不是很了解。今天就通过这篇 @Winterbe 的译文，一起深入了解下如何使用它吧。

## 目录 ##

一、Stream 流是如何工作的？

二、不同类型的 Stream 流

三、Stream 流的处理顺序

四、中间操作顺序这么重要？

五、数据流复用问题

六、高级操作

* 6.1 Collect
* 6.2 FlatMap
* 6.3 Reduce

七、并行流

八、结语

当我第一次阅读 Java8 中的 Stream API 时，说实话，我非常困惑，因为它的名字听起来与 Java I0 框架中的 ` InputStream` 和 ` OutputStream` 非常类似。但是实际上，它们完全是不同的东西。

Java8 Stream 使用的是函数式编程模式，如同它的名字一样，它可以被用来对集合进行链状流式的操作。

本文就将带着你如何使用 Java 8 不同类型的 Stream 操作。同时您还将了解流的处理顺序，以及不同顺序的流操作是如何影响运行时性能的。

我们还将学习终端操作 API ` reduce` ， ` collect` 以及 ` flatMap` 的详细介绍，最后我们再来深入的探讨一下 Java8 并行流。

> 
> 
> 
> 注意：如果您还不熟悉 Java 8 lambda 表达式，函数式接口以及方法引用，您可以先阅读一下小哈的另一篇译文 [《Java8 新特性教程》](
> https://link.juejin.im?target=https%3A%2F%2Fwww.exception.site%2Fjava8%2Fjava8-new-features
> )
> 
> 

接下来，就让我们进入正题吧！

## 一、Stream 流是如何工作的？ ##

流表示包含着一系列元素的集合，我们可以对其做不同类型的操作，用来对这些元素执行计算。听上去可能有点拗口，让我们用代码说话：

` List<String> myList = Arrays.asList( "a1" , "a2" , "b1" , "c2" , "c1" ); myList .stream() // 创建流.filter(s -> s.startsWith( "c" )) // 执行过滤，过滤出以 c 为前缀的字符串.map(String::toUpperCase) // 转换成大写.sorted() // 排序.forEach(System.out::println); // for 循环打印 // C1 // C2 复制代码`

我们可以对流进行中间操作或者终端操作。小伙伴们可能会疑问？ **什么是中间操作？什么又是终端操作？**

![Stream中间操作，终端操作](https://user-gold-cdn.xitu.io/2019/4/25/16a52771d4587e64?imageView2/0/w/1280/h/960/ignore-error/1) Stream中间操作，终端操作

* **①** ：中间操作会再次返回一个流，所以，我们可以链接多个中间操作，注意这里是不用加分号的。上图中的 ` filter` 过滤， ` map` 对象转换， ` sorted` 排序，就属于中间操作。
* **②** ：终端操作是对流操作的一个结束动作，一般返回 ` void` 或者一个非流的结果。上图中的 ` forEach` 循环 就是一个终止操作。

看完上面的操作，感觉是不是很像一个流水线式操作呢。

实际上，大部分流操作都支持 lambda 表达式作为参数，正确理解，应该说是接受一个函数式接口的实现作为参数。

## 二、不同类型的 Stream 流 ##

我们可以从各种数据源中创建 Stream 流，其中以 Collection 集合最为常见。如 ` List` 和 ` Set` 均支持 ` stream()` 方法来创建顺序流或者是并行流。

并行流是通过多线程的方式来执行的，它能够充分发挥多核 CPU 的优势来提升性能。本文在最后再来介绍并行流，我们先讨论顺序流：

` Arrays.asList( "a1" , "a2" , "a3" ) .stream() // 创建流.findFirst() // 找到第一个元素.ifPresent(System.out::println); // 如果存在，即输出 // a1 复制代码`

在集合上调用 ` stream()` 方法会返回一个普通的 Stream 流。但是, 您大可不必刻意地创建一个集合，再通过集合来获取 Stream 流，您还可以通过如下这种方式：

` Stream.of( "a1" , "a2" , "a3" ) .findFirst() .ifPresent(System.out::println); // a1 复制代码`

例如上面这样，我们可以通过 ` Stream.of()` 从一堆对象中创建 Stream 流。

除了常规对象流之外，Java 8还附带了一些特殊类型的流，用于处理原始数据类型 ` int` ， ` long` 以及 ` double` 。说道这里，你可能已经猜到了它们就是 ` IntStream` ， ` LongStream` 还有 ` DoubleStream` 。

其中， ` IntStreams.range()` 方法还可以被用来取代常规的 ` for` 循环, 如下所示：

` IntStream.range( 1 , 4 ) .forEach(System.out::println); // 相当于 for (int i = 1; i < 4; i++) {} // 1 // 2 // 3 复制代码`

上面这些原始类型流的工作方式与常规对象流基本是一样的，但还是略微存在一些区别：

* 

原始类型流使用其独有的函数式接口，例如 ` IntFunction` 代替 ` Function` ， ` IntPredicate` 代替 ` Predicate` 。

* 

原始类型流支持额外的终端聚合操作， ` sum()` 以及 ` average()` ，如下所示：

` Arrays.stream( new int [] { 1 , 2 , 3 }) .map(n -> 2 * n + 1 ) // 对数值中的每个对象执行 2*n + 1 操作.average() // 求平均值.ifPresent(System.out::println); // 如果值不为空，则输出 // 5.0 复制代码`

但是，偶尔我们也有这种需求，需要将常规对象流转换为原始类型流，这个时候，中间操作 ` mapToInt()` ， ` mapToLong()` 以及 ` mapToDouble` 就派上用场了：

` Stream.of( "a1" , "a2" , "a3" ) .map(s -> s.substring( 1 )) // 对每个字符串元素从下标1位置开始截取.mapToInt(Integer::parseInt) // 转成 int 基础类型类型流.max() // 取最大值.ifPresent(System.out::println); // 不为空则输出 // 3 复制代码`

如果说，您需要将原始类型流装换成对象流，您可以使用 ` mapToObj()` 来达到目的：

` IntStream.range( 1 , 4 ) .mapToObj(i -> "a" + i) // for 循环 1->4, 拼接前缀 a.forEach(System.out::println); // for 循环打印 // a1 // a2 // a3 复制代码`

下面是一个组合示例，我们将双精度流首先转换成 ` int` 类型流，然后再将其装换成对象流：

` Stream.of( 1.0 , 2.0 , 3.0 ) .mapToInt(Double::intValue) // double 类型转 int.mapToObj(i -> "a" + i) // 对值拼接前缀 a.forEach(System.out::println); // for 循环打印 // a1 // a2 // a3 复制代码`

## 三、Stream 流的处理顺序 ##

上小节中，我们已经学会了如何创建不同类型的 Stream 流，接下来我们再深入了解下数据流的执行顺序。

在讨论处理顺序之前，您需要明确一点，那就是中间操作的有个重要特性 —— **延迟性** 。观察下面这个没有终端操作的示例代码：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .filter(s -> { System.out.println( "filter: " + s); return true ; }); 复制代码`

执行此代码段时，您可能会认为，将依次打印 "d2", "a2", "b1", "b3", "c" 元素。然而当你实际去执行的时候，它不会打印任何内容。

**为什么呢？**

原因是：当且仅当存在终端操作时，中间操作操作才会被执行。

是不是不信？接下来，对上面的代码添加 ` forEach` 终端操作：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .filter(s -> { System.out.println( "filter: " + s); return true ; }) .forEach(s -> System.out.println( "forEach: " + s)); 复制代码`

再次执行，我们会看到输出如下：

` filter: d2 forEach: d2 filter: a2 forEach: a2 filter: b1 forEach: b1 filter: b3 forEach: b3 filter: c forEach: c 复制代码`

输出的顺序可能会让你很惊讶！你脑海里肯定会想，应该是先将所有 ` filter` 前缀的字符串打印出来，接着才会打印 ` forEach` 前缀的字符串。

事实上，输出的结果却是随着链条垂直移动的。比如说，当 Stream 开始处理 d2 元素时，它实际上会在执行完 filter 操作后，再执行 forEach 操作，接着才会处理第二个元素。

是不是很神奇？为什么要设计成这样呢？

原因是出于性能的考虑。这样设计可以减少对每个元素的实际操作数，看完下面代码你就明白了：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .map(s -> { System.out.println( "map: " + s); return s.toUpperCase(); // 转大写 }) .anyMatch(s -> { System.out.println( "anyMatch: " + s); return s.startsWith( "A" ); // 过滤出以 A 为前缀的元素 }); // map: d2 // anyMatch: D2 // map: a2 // anyMatch: A2 复制代码`

终端操作 ` anyMatch()` 表示任何一个元素以 A 为前缀，返回为 ` true` ，就停止循环。所以它会从 ` d2` 开始匹配，接着循环到 ` a2` 的时候，返回为 ` true` ，于是停止循环。

由于数据流的链式调用是垂直执行的， ` map` 这里只需要执行两次。相对于水平执行来说， ` map` 会执行尽可能少的次数，而不是把所有元素都 ` map` 转换一遍。

## 四、中间操作顺序这么重要？ ##

下面的例子由两个中间操作 ` map` 和 ` filter` ，以及一个终端操作 ` forEach` 组成。让我们再来看看这些操作是如何执行的：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .map(s -> { System.out.println( "map: " + s); return s.toUpperCase(); // 转大写 }) .filter(s -> { System.out.println( "filter: " + s); return s.startsWith( "A" ); // 过滤出以 A 为前缀的元素 }) .forEach(s -> System.out.println( "forEach: " + s)); // for 循环输出 // map: d2 // filter: D2 // map: a2 // filter: A2 // forEach: A2 // map: b1 // filter: B1 // map: b3 // filter: B3 // map: c // filter: C 复制代码`

学习了上面一小节，您应该已经知道了， ` map` 和 ` filter` 会对集合中的每个字符串调用五次，而 ` forEach` 却只会调用一次，因为只有 "a2" 满足过滤条件。

如果我们改变中间操作的顺序，将 ` filter` 移动到链头的最开始，就可以大大减少实际的执行次数：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .filter(s -> { System.out.println( "filter: " + s) return s.startsWith( "a" ); // 过滤出以 a 为前缀的元素 }) .map(s -> { System.out.println( "map: " + s); return s.toUpperCase(); // 转大写 }) .forEach(s -> System.out.println( "forEach: " + s)); // for 循环输出 // filter: d2 // filter: a2 // map: a2 // forEach: A2 // filter: b1 // filter: b3 // filter: c 复制代码`

现在， ` map` 仅仅只需调用一次，性能得到了提升，这种小技巧对于流中存在大量元素来说，是非常很有用的。

接下来，让我们对上面的代码再添加一个中间操作 ` sorted` ：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .sorted((s1, s2) -> { System.out.printf( "sort: %s; %s\n" , s1, s2); return s1.compareTo(s2); // 排序 }) .filter(s -> { System.out.println( "filter: " + s); return s.startsWith( "a" ); // 过滤出以 a 为前缀的元素 }) .map(s -> { System.out.println( "map: " + s); return s.toUpperCase(); // 转大写 }) .forEach(s -> System.out.println( "forEach: " + s)); // for 循环输出 复制代码`

` sorted` 是一个有状态的操作，因为它需要在处理的过程中，保存状态以对集合中的元素进行排序。

执行上面代码，输出如下：

` sort: a2; d2 sort: b1; a2 sort: b1; d2 sort: b1; a2 sort: b3; b1 sort: b3; d2 sort: c; b3 sort: c; d2 filter: a2 map: a2 forEach: A2 filter: b1 filter: b3 filter: c filter: d2 复制代码`

咦咦咦？这次怎么又不是垂直执行了。你需要知道的是， ` sorted` 是水平执行的。因此，在这种情况下， ` sorted` 会对集合中的元素组合调用八次。这里，我们也可以利用上面说道的优化技巧，将 filter 过滤中间操作移动到开头部分：

` Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .filter(s -> { System.out.println( "filter: " + s); return s.startsWith( "a" ); }) .sorted((s1, s2) -> { System.out.printf( "sort: %s; %s\n" , s1, s2); return s1.compareTo(s2); }) .map(s -> { System.out.println( "map: " + s); return s.toUpperCase(); }) .forEach(s -> System.out.println( "forEach: " + s)); // filter: d2 // filter: a2 // filter: b1 // filter: b3 // filter: c // map: a2 // forEach: A2 复制代码`

从上面的输出中，我们看到了 ` sorted` 从未被调用过，因为经过 ` filter` 过后的元素已经减少到只有一个，这种情况下，是不用执行排序操作的。因此性能被大大提高了。

## 五、数据流复用问题 ##

Java8 Stream 流是不能被复用的，一旦你调用任何终端操作，流就会关闭：

` Stream<String> stream = Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .filter(s -> s.startsWith( "a" )); stream.anyMatch(s -> true ); // ok stream.noneMatch(s -> true ); // exception 复制代码`

当我们对 stream 调用了 ` anyMatch` 终端操作以后，流即关闭了，再调用 ` noneMatch` 就会抛出异常：

` java.lang.IllegalStateException: stream has already been operated upon or closed at java.util.stream.AbstractPipeline.evaluate(AbstractPipeline.java: 229 ) at java.util.stream.ReferencePipeline.noneMatch(ReferencePipeline.java: 459 ) at com.winterbe.java8.Streams5.test7(Streams5.java: 38 ) at com.winterbe.java8.Streams5.main(Streams5.java: 28 ) 复制代码`

为了克服这个限制，我们必须为我们想要执行的每个终端操作创建一个新的流链，例如，我们可以通过 ` Supplier` 来包装一下流，通过 ` get()` 方法来构建一个新的 ` Stream` 流，如下所示：

` Supplier<Stream<String>> streamSupplier = () -> Stream.of( "d2" , "a2" , "b1" , "b3" , "c" ) .filter(s -> s.startsWith( "a" )); streamSupplier.get().anyMatch(s -> true ); // ok streamSupplier.get().noneMatch(s -> true ); // ok 复制代码`

通过构造一个新的流，来避开流不能被复用的限制, 这也是取巧的一种方式。

## 六、高级操作 ##

` Streams` 支持的操作很丰富，除了上面介绍的这些比较常用的中间操作，如 ` filter` 或 ` map` （参见 [Stream Javadoc]( https://link.juejin.im?target=http%3A%2F%2Fdocs.oracle.com%2Fjavase%2F8%2Fdocs%2Fapi%2Fjava%2Futil%2Fstream%2FStream.html ) ）外。还有一些更复杂的操作，如 ` collect` ， ` flatMap` 以及 ` reduce` 。接下来，就让我们学习一下：

本小节中的大多数代码示例均会使用以下 ` List<Person>` 进行演示：

` class Person { String name; int age; Person(String name, int age) { this.name = name; this.age = age; } @Override public String toString () { return name; } } // 构建一个 Person 集合 List<Person> persons = Arrays.asList( new Person( "Max" , 18 ), new Person( "Peter" , 23 ), new Person( "Pamela" , 23 ), new Person( "David" , 12 )); 复制代码`

### 6.1 Collect ###

collect 是一个非常有用的终端操作，它可以将流中的元素转变成另外一个不同的对象，例如一个 ` List` ， ` Set` 或 ` Map` 。collect 接受入参为 ` Collector` （收集器），它由四个不同的操作组成：供应器（supplier）、累加器（accumulator）、组合器（combiner）和终止器（finisher）。

这些都是个啥？别慌，看上去非常复杂的样子，但好在大多数情况下，您并不需要自己去实现收集器。因为 Java 8通过 ` Collectors` 类内置了各种常用的收集器，你直接拿来用就行了。

让我们先从一个非常常见的用例开始：

` List<Person> filtered = persons .stream() // 构建流.filter(p -> p.name.startsWith( "P" )) // 过滤出名字以 P 开头的.collect(Collectors.toList()); // 生成一个新的 List System.out.println(filtered); // [Peter, Pamela] 复制代码`

你也看到了，从流中构造一个 ` List` 异常简单。如果说你需要构造一个 ` Set` 集合，只需要使用 ` Collectors.toSet()` 就可以了。

接下来这个示例，将会按年龄对所有人进行分组：

` Map<Integer, List<Person>> personsByAge = persons .stream() .collect(Collectors.groupingBy(p -> p.age)); // 以年龄为 key,进行分组 personsByAge .forEach((age, p) -> System.out.format( "age %s: %s\n" , age, p)); // age 18: [Max] // age 23: [Peter, Pamela] // age 12: [David] 复制代码`

除了上面这些操作。您还可以在流上执行聚合操作，例如，计算所有人的平均年龄：

` Double averageAge = persons .stream() .collect(Collectors.averagingInt(p -> p.age)); // 聚合出平均年龄 System.out.println(averageAge); // 19.0 复制代码`

如果您还想得到一个更全面的统计信息，摘要收集器可以返回一个特殊的内置统计对象。通过它，我们可以简单地计算出最小年龄、最大年龄、平均年龄、总和以及总数量。

` IntSummaryStatistics ageSummary = persons .stream() .collect(Collectors.summarizingInt(p -> p.age)); // 生成摘要统计 System.out.println(ageSummary); // IntSummaryStatistics{count=4, sum=76, min=12, average=19.000000, max=23} 复制代码`

下一个这个示例，可以将所有人名连接成一个字符串：

` String phrase = persons .stream() .filter(p -> p.age >= 18 ) // 过滤出年龄大于等于18的.map(p -> p.name) // 提取名字.collect(Collectors.joining( " and " , "In Germany " , " are of legal age." )); // 以 In Germany 开头，and 连接各元素，再以 are of legal age. 结束 System.out.println(phrase); // In Germany Max and Peter and Pamela are of legal age. 复制代码`

连接收集器的入参接受分隔符，以及可选的前缀以及后缀。

对于如何将流转换为 ` Map` 集合，我们必须指定 ` Map` 的键和值。这里需要注意， ` Map` 的键必须是唯一的，否则会抛出 ` IllegalStateException` 异常。

你可以选择传递一个合并函数作为额外的参数来避免发生这个异常:

` Map<Integer, String> map = persons .stream() .collect(Collectors.toMap( p -> p.age, p -> p.name, (name1, name2) -> name1 + ";" + name2)); // 对于同样 key 的，将值拼接 System.out.println(map); // {18=Max, 23=Peter;Pamela, 12=David} 复制代码`

既然我们已经知道了这些强大的内置收集器，接下来就让我们尝试构建自定义收集器吧。

比如说，我们希望将流中的所有人转换成一个字符串，包含所有大写的名称，并以 ` |` 分割。为了达到这种效果，我们需要通过 ` Collector.of()` 创建一个新的收集器。同时，我们还需要传入收集器的四个组成部分：供应器、累加器、组合器和终止器。

` Collector<Person, StringJoiner, String> personNameCollector = Collector.of( () -> new StringJoiner( " | " ), // supplier 供应器 (j, p) -> j.add(p.name.toUpperCase()), // accumulator 累加器 (j1, j2) -> j1.merge(j2), // combiner 组合器 StringJoiner::toString); // finisher 终止器 String names = persons .stream() .collect(personNameCollector); // 传入自定义的收集器 System.out.println(names); // MAX | PETER | PAMELA | DAVID 复制代码`

由于Java 中的字符串是 final 类型的，我们需要借助辅助类 ` StringJoiner` ，来帮我们构造字符串。

最开始供应器使用分隔符构造了一个 ` StringJointer` 。

累加器用于将每个人的人名转大写，然后加到 ` StringJointer` 中。

组合器将两个 ` StringJointer` 合并为一个。

最终，终结器从 ` StringJointer` 构造出预期的字符串。

### 6.2 FlatMap ###

上面我们已经学会了如通过 ` map` 操作, 将流中的对象转换为另一种类型。但是， ` Map` 只能将每个对象映射到另一个对象。

如果说，我们想要将一个对象转换为多个其他对象或者根本不做转换操作呢？这个时候， ` flatMap` 就派上用场了。

` FlatMap` 能够将流的每个元素, 转换为其他对象的流。因此，每个对象可以被转换为零个，一个或多个其他对象，并以流的方式返回。之后，这些流的内容会被放入 ` flatMap` 返回的流中。

在学习如何实际操作 ` flatMap` 之前，我们先新建两个类，用来测试：

` class Foo { String name; List<Bar> bars = new ArrayList<>(); Foo(String name) { this.name = name; } } class Bar { String name; Bar(String name) { this.name = name; } } 复制代码`

接下来，通过我们上面学习到的流知识，来实例化一些对象：

` List<Foo> foos = new ArrayList<>(); // 创建 foos 集合 IntStream .range( 1 , 4 ) .forEach(i -> foos.add( new Foo( "Foo" + i))); // 创建 bars 集合 foos.forEach(f -> IntStream .range( 1 , 4 ) .forEach(i -> f.bars.add( new Bar( "Bar" + i + " <- " + f.name)))); 复制代码`

我们创建了包含三个 ` foo` 的集合，每个 ` foo` 中又包含三个 ` bar` 。

` flatMap` 的入参接受一个返回对象流的函数。为了处理每个 ` foo` 中的 ` bar` ，我们需要传入相应 stream 流：

` foos.stream() .flatMap(f -> f.bars.stream()) .forEach(b -> System.out.println(b.name)); // Bar1 <- Foo1 // Bar2 <- Foo1 // Bar3 <- Foo1 // Bar1 <- Foo2 // Bar2 <- Foo2 // Bar3 <- Foo2 // Bar1 <- Foo3 // Bar2 <- Foo3 // Bar3 <- Foo3 复制代码`

如上所示，我们已成功将三个 ` foo` 对象的流转换为九个 ` bar` 对象的流。

最后，上面的这段代码可以简化为单一的流式操作：

` IntStream.range( 1 , 4 ) .mapToObj(i -> new Foo( "Foo" + i)) .peek(f -> IntStream.range( 1 , 4 ) .mapToObj(i -> new Bar( "Bar" + i + " <- " f.name)) .forEach(f.bars::add)) .flatMap(f -> f.bars.stream()) .forEach(b -> System.out.println(b.name)); 复制代码`

` flatMap` 也可用于Java8引入的 ` Optional` 类。 ` Optional` 的 ` flatMap` 操作返回一个 ` Optional` 或其他类型的对象。所以它可以用于避免繁琐的 ` null` 检查。

接下来，让我们创建层次更深的对象：

` class Outer { Nested nested; } class Nested { Inner inner; } class Inner { String foo; } 复制代码`

为了处理从 Outer 对象中获取最底层的 foo 字符串，你需要添加多个 ` null` 检查来避免可能发生的 ` NullPointerException` ，如下所示：

` Outer outer = new Outer(); if (outer != null && outer.nested != null && outer.nested.inner != null ) { System.out.println(outer.nested.inner.foo); } 复制代码`

我们还可以使用 ` Optional` 的 ` flatMap` 操作，来完成上述相同功能的判断，且更加优雅：

` Optional.of( new Outer()) .flatMap(o -> Optional.ofNullable(o.nested)) .flatMap(n -> Optional.ofNullable(n.inner)) .flatMap(i -> Optional.ofNullable(i.foo)) .ifPresent(System.out::println); 复制代码`

如果不为空的话，每个 ` flatMap` 的调用都会返回预期对象的 ` Optional` 包装，否则返回为 ` null` 的 ` Optional` 包装类。

> 
> 
> 
> 笔者补充：关于 Optional 可参见我另一篇译文 [《Java8 新特性如何防止空指针异常》](
> https://link.juejin.im?target=https%3A%2F%2Fwww.exception.site%2Fjava8%2Fjava8-avoid-null-check
> )
> 
> 

### 6.3 Reduce ###

规约操作可以将流的所有元素组合成一个结果。Java 8 支持三种不同的 ` reduce` 方法。第一种将流中的元素规约成流中的一个元素。

让我们看看如何使用这种方法，来筛选出年龄最大的那个人：

` persons .stream() .reduce((p1, p2) -> p1.age > p2.age ? p1 : p2) .ifPresent(System.out::println); // Pamela 复制代码`

` reduce` 方法接受 ` BinaryOperator` 积累函数。该函数实际上是两个操作数类型相同的 ` BiFunction` 。 ` BiFunction` 功能和 ` Function` 一样，但是它接受两个参数。示例代码中，我们比较两个人的年龄，来返回年龄较大的人。

第二种 ` reduce` 方法接受标识值和 ` BinaryOperator` 累加器。此方法可用于构造一个新的 ` Person` ，其中包含来自流中所有其他人的聚合名称和年龄：

` Person result = persons .stream() .reduce( new Person( "" , 0 ), (p1, p2) -> { p1.age += p2.age; p1.name += p2.name; return p1; }); System.out.format( "name=%s; age=%s" , result.name, result.age); // name=MaxPeterPamelaDavid; age=76 复制代码`

第三种 ` reduce` 方法接受三个参数：标识值， ` BiFunction` 累加器和类型的组合器函数 ` BinaryOperator` 。由于初始值的类型不一定为 ` Person` ，我们可以使用这个归约函数来计算所有人的年龄总和：

` Integer ageSum = persons .stream() .reduce( 0 , (sum, p) -> sum += p.age, (sum1, sum2) -> sum1 + sum2); System.out.println(ageSum); // 76 复制代码`

结果为 76 ，但是内部究竟发生了什么呢？让我们再打印一些调试日志：

` Integer ageSum = persons .stream() .reduce( 0 , (sum, p) -> { System.out.format( "accumulator: sum=%s; person=%s\n" , sum, p); return sum += p.age; }, (sum1, sum2) -> { System.out.format( "combiner: sum1=%s; sum2=%s\n" , sum1, sum2); return sum1 + sum2; }); // accumulator: sum=0; person=Max // accumulator: sum=18; person=Peter // accumulator: sum=41; person=Pamela // accumulator: sum=64; person=David 复制代码`

你可以看到，累加器函数完成了所有工作。它首先使用初始值 ` 0` 和第一个人年龄相加。接下来的三步中 ` sum` 会持续增加，直到76。

等等？好像哪里不太对！组合器从来都没有调用过啊？

我们以并行流的方式运行上面的代码，看看日志输出：

` Integer ageSum = persons .parallelStream() .reduce( 0 , (sum, p) -> { System.out.format( "accumulator: sum=%s; person=%s\n" , sum, p); return sum += p.age; }, (sum1, sum2) -> { System.out.format( "combiner: sum1=%s; sum2=%s\n" , sum1, sum2); return sum1 + sum2; }); // accumulator: sum=0; person=Pamela // accumulator: sum=0; person=David // accumulator: sum=0; person=Max // accumulator: sum=0; person=Peter // combiner: sum1=18; sum2=23 // combiner: sum1=23; sum2=12 // combiner: sum1=41; sum2=35 复制代码`

并行流的执行方式完全不同。这里组合器被调用了。实际上，由于累加器被并行调用，组合器需要被用于计算部分累加值的总和。

让我们在下一章深入探讨并行流。

## 七、并行流 ##

流是可以并行执行的，当流中存在大量元素时，可以显著提升性能。并行流底层使用的 ` ForkJoinPool` , 它由 ` ForkJoinPool.commonPool()` 方法提供。底层线程池的大小最多为五个 - 具体取决于 CPU 可用核心数：

` ForkJoinPool commonPool = ForkJoinPool.commonPool(); System.out.println(commonPool.getParallelism()); // 3 复制代码`

在我的机器上，公共池初始化默认值为 3。你也可以通过设置以下JVM参数可以减小或增加此值：

` -Djava.util.concurrent.ForkJoinPool.common.parallelism=5 复制代码`

集合支持 ` parallelStream()` 方法来创建元素的并行流。或者你可以在已存在的数据流上调用中间方法 ` parallel()` ，将串行流转换为并行流，这也是可以的。

为了详细了解并行流的执行行为，我们在下面的示例代码中，打印当前线程的信息：

` Arrays.asList( "a1" , "a2" , "b1" , "c2" , "c1" ) .parallelStream() .filter(s -> { System.out.format( "filter: %s [%s]\n" , s, Thread.currentThread().getName()); return true ; }) .map(s -> { System.out.format( "map: %s [%s]\n" , s, Thread.currentThread().getName()); return s.toUpperCase(); }) .forEach(s -> System.out.format( "forEach: %s [%s]\n" , s, Thread.currentThread().getName())); 复制代码`

通过日志输出，我们可以对哪个线程被用于执行流式操作，有个更深入的理解：

` filter: b1 [main] filter: a2 [ForkJoinPool.commonPool-worker- 1 ] map: a2 [ForkJoinPool.commonPool-worker- 1 ] filter: c2 [ForkJoinPool.commonPool-worker- 3 ] map: c2 [ForkJoinPool.commonPool-worker- 3 ] filter: c1 [ForkJoinPool.commonPool-worker- 2 ] map: c1 [ForkJoinPool.commonPool-worker- 2 ] forEach: C2 [ForkJoinPool.commonPool-worker- 3 ] forEach: A2 [ForkJoinPool.commonPool-worker- 1 ] map: b1 [main] forEach: B1 [main] filter: a1 [ForkJoinPool.commonPool-worker- 3 ] map: a1 [ForkJoinPool.commonPool-worker- 3 ] forEach: A1 [ForkJoinPool.commonPool-worker- 3 ] forEach: C1 [ForkJoinPool.commonPool-worker- 2 ] 复制代码`

如您所见，并行流使用了所有的 ` ForkJoinPool` 中的可用线程来执行流式操作。在持续的运行中，输出结果可能有所不同，因为所使用的特定线程是非特定的。

让我们通过添加中间操作 ` sort` 来扩展上面示例：

` Arrays.asList( "a1" , "a2" , "b1" , "c2" , "c1" ) .parallelStream() .filter(s -> { System.out.format( "filter: %s [%s]\n" , s, Thread.currentThread().getName()); return true ; }) .map(s -> { System.out.format( "map: %s [%s]\n" , s, Thread.currentThread().getName()); return s.toUpperCase(); }) .sorted((s1, s2) -> { System.out.format( "sort: %s <> %s [%s]\n" , s1, s2, Thread.currentThread().getName()); return s1.compareTo(s2); }) .forEach(s -> System.out.format( "forEach: %s [%s]\n" , s, Thread.currentThread().getName())); 复制代码`

运行代码，输出结果看上去有些奇怪：

` filter: c2 [ForkJoinPool.commonPool-worker-3] filter: c1 [ForkJoinPool.commonPool-worker-2] map: c1 [ForkJoinPool.commonPool-worker-2] filter: a2 [ForkJoinPool.commonPool-worker-1] map: a2 [ForkJoinPool.commonPool-worker-1] filter: b1 [main] map: b1 [main] filter: a1 [ForkJoinPool.commonPool-worker-2] map: a1 [ForkJoinPool.commonPool-worker-2] map: c2 [ForkJoinPool.commonPool-worker-3] sort: A2 <> A1 [main] sort: B1 <> A2 [main] sort: C2 <> B1 [main] sort: C1 <> C2 [main] sort: C1 <> B1 [main] sort: C1 <> C2 [main] forEach: A1 [ForkJoinPool.commonPool-worker-1] forEach: C2 [ForkJoinPool.commonPool-worker-3] forEach: B1 [main] forEach: A2 [ForkJoinPool.commonPool-worker-2] forEach: C1 [ForkJoinPool.commonPool-worker-1] 复制代码`

貌似 ` sort` 只在主线程上串行执行。但是实际上，并行流中的 ` sort` 在底层使用了Java8中新的方法 ` Arrays.parallelSort()` 。如 [javadoc]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F8%2Fdocs%2Fapi%2Fjava%2Futil%2FArrays.html%23parallelSort-T%3AA- ) 官方文档解释的，这个方法会按照数据长度来决定以串行方式，或者以并行的方式来执行。

> 
> 
> 
> 如果指定数据的长度小于最小数值，它则使用相应的 ` Arrays.sort` 方法来进行排序。
> 
> 

回到上小节 ` reduce` 的例子。我们已经发现了组合器函数只在并行流中调用，而不不会在串行流中被调用。

让我们来实际观察一下涉及到哪个线程：

` List<Person> persons = Arrays.asList( new Person( "Max" , 18), new Person( "Peter" , 23), new Person( "Pamela" , 23), new Person( "David" , 12)); persons .parallelStream() .reduce(0, (sum, p) -> { System.out.format( "accumulator: sum=%s; person=%s [%s]\n" , sum, p, Thread.currentThread().getName()); return sum += p.age; }, (sum1, sum2) -> { System.out.format( "combiner: sum1=%s; sum2=%s [%s]\n" , sum1, sum2, Thread.currentThread().getName()); return sum1 + sum2; }); 复制代码`

通过控制台日志输出，累加器和组合器均在所有可用的线程上并行执行：

` accumulator: sum=0; person=Pamela; [main] accumulator: sum=0; person=Max; [ForkJoinPool.commonPool-worker-3] accumulator: sum=0; person=David; [ForkJoinPool.commonPool-worker-2] accumulator: sum=0; person=Peter; [ForkJoinPool.commonPool-worker-1] combiner: sum1=18; sum2=23; [ForkJoinPool.commonPool-worker-1] combiner: sum1=23; sum2=12; [ForkJoinPool.commonPool-worker-2] combiner: sum1=41; sum2=35; [ForkJoinPool.commonPool-worker-2] 复制代码`

总之，你需要记住的是，并行流对含有大量元素的数据流提升性能极大。但是你也需要记住并行流的一些操作，例如 ` reduce` 和 ` collect` 操作，需要额外的计算（如组合操作），这在串行执行时是并不需要。

此外，我们也了解了，所有并行流操作都共享相同的 JVM 相关的公共 ` ForkJoinPool` 。所以你可能需要避免写出一些又慢又卡的流式操作，这很有可能会拖慢你应用中，严重依赖并行流的其它部分代码的性能。

## 八、结语 ##

Java8 Stream 流编程指南到这里就结束了。如果您有兴趣了解更多有关 Java 8 Stream 流的相关信息，我建议您使用 [Stream Javadoc]( https://link.juejin.im?target=http%3A%2F%2Fdocs.oracle.com%2Fjavase%2F8%2Fdocs%2Fapi%2Fjava%2Futil%2Fstream%2Fpackage-summary.html%23NonInterference ) 阅读官方文档。如果您想了解有关底层机制的更多信息，您也可以阅读 Martin Fowlers 关于 [Collection Pipelines]( https://link.juejin.im?target=http%3A%2F%2Fmartinfowler.com%2Farticles%2Fcollection-pipeline%2F ) 的文章。

最后，祝您学习愉快！

## 赠送 10G 面试&学习福利资源 ##

获取方式: 关注微信公众号: **小哈学Java** , 后台回复" **666** "，既可 **免费无套路获取资源链接** ，下面是目录以及部分截图：

![关注微信公众号【小哈学Java】,回复“666”，即可免费无套路领取哦](https://user-gold-cdn.xitu.io/2019/4/27/16a5dc18d3543937?imageView2/0/w/1280/h/960/ignore-error/1) 关注微信公众号【小哈学Java】,回复“666”，即可免费无套路领取哦

## 欢迎关注微信公众号: 小哈学Java ##

![小哈学Java，关注领取10G面试学习资料哦](https://user-gold-cdn.xitu.io/2019/4/27/16a5dc18e4c739ff?imageView2/0/w/1280/h/960/ignore-error/1) 小哈学Java，关注领取10G面试学习资料哦