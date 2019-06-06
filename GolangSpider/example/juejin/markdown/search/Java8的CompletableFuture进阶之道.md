# Java8的CompletableFuture进阶之道 #

# 简介 #

作为Java 8 Concurrency API改进而引入，本文是CompletableFuture类的功能和用例的介绍。同时在Java 9 也有对CompletableFuture有一些改进，之后再进入讲解。

# Future计算 #

Future异步计算很难操作，通常我们希望将任何计算逻辑视为一系列步骤。但是在异步计算的情况下，表示为回调的方法往往分散在代码中或者深深地嵌套在彼此内部。但是当我们需要处理其中一个步骤中可能发生的错误时，情况可能会变得更复杂。

Futrue接口是Java 5中作为异步计算而新增的，但它没有任何方法去进行计算组合或者处理可能出现的错误。

在Java 8中，引入了CompletableFuture类。与Future接口一起，它还实现了CompletionStage接口。此接口定义了可与其他Future组合成异步计算契约。

CompletableFuture同时是一个组合和一个框架，具有大约50种不同的构成，结合，执行异步计算步骤和处理错误。

如此庞大的API可能会令人难以招架，下文将调一些重要的做重点介绍。

# 使用CompletableFuture作为Future实现 #

首先，CompletableFuture类实现Future接口，因此你可以将其用作Future实现，但需要额外的完成实现逻辑。

例如，你可以使用无构参构造函数创建此类的实例，然后使用 ` complete` 方法完成。消费者可以使用get方法来阻塞当前线程，直到 ` get()` 结果。

在下面的示例中，我们有一个创建CompletableFuture实例的方法，然后在另一个线程中计算并立即返回Future。

计算完成后，该方法通过将结果提供给完整方法来完成Future：

` public Future<String> calculateAsync() throws InterruptedException { CompletableFuture<String> completableFuture = new CompletableFuture<>(); Executors.newCachedThreadPool().submit(() -> { Thread.sleep(500); completableFuture.complete( "Hello" ); return null; }); return completableFuture; } 复制代码`

为了分离计算，我们使用了 Executor API ，这种创建和完成 CompletableFuture的方法 可以与任何并发包（包括原始线程）一起使用。

请注意， **该 ` calculateAsync` 方法返回一个 ` Future` 实例。**

我们只是调用方法，接收 Future 实例并在我们准备阻塞结果时调用它的 get 方法。

另请注意， get 方法抛出一些已检查的异常，即 ExecutionException （封装计算期间发生的异常）和 InterruptedException （表示执行方法的线程被中断的异常）：

` Future<String> completableFuture = calculateAsync(); // ... String result = completableFuture.get(); assertEquals( "Hello" , result); 复制代码`

如果你已经知道计算的结果，也可以用变成同步的方式来返回结果。

` Future<String> completableFuture = CompletableFuture.completedFuture( "Hello" ); // ... String result = completableFuture.get(); assertEquals( "Hello" , result); 复制代码`

作为在某些场景中，你可能希望取消Future任务的执行。

假设我们没有找到结果并决定完全取消异步执行任务。这可以通过Future的取消方法完成。此方法 ` mayInterruptIfRunning` ，但在CompletableFuture的情况下，它没有任何效果，因为中断不用于控制CompletableFuture的处理。

这是异步方法的修改版本：

` public Future<String> calculateAsyncWithCancellation() throws InterruptedException { CompletableFuture<String> completableFuture = new CompletableFuture<>(); Executors.newCachedThreadPool().submit(() -> { Thread.sleep(500); completableFuture.cancel( false ); return null; }); return completableFuture; } 复制代码`

当我们使用Future.get()方法阻塞结果时， ` cancel()` 表示取消执行，它将抛出CancellationException：

` Future<String> future = calculateAsyncWithCancellation(); future.get(); // CancellationException 复制代码`

# API介绍 #

## static方法说明 ##

上面的代码很简单，下面介绍几个 **static** 方法，它们使用任务来实例化一个 CompletableFuture 实例。

` CompletableFuture.runAsync(Runnable runnable); CompletableFuture.runAsync(Runnable runnable, Executor executor); CompletableFuture.supplyAsync(Supplier<U> supplier); CompletableFuture.supplyAsync(Supplier<U> supplier, Executor executor) 复制代码`

* runAsync 方法接收的是 Runnable 的实例，但是它没有返回值
* supplyAsync 方法是JDK8函数式接口，无参数，会返回一个结果
* 这两个方法是 executor 的升级，表示让任务在指定的线程池中执行，不指定的话，通常任务是在 ForkJoinPool.commonPool() 线程池中执行的。

## supplyAsync()使用 ##

静态方法 ` runAsync` 和 ` supplyAsync` 允许我们相应地从Runnable和Supplier功能类型中创建CompletableFuture实例。

该Runnable的接口是在线程使用旧的接口，它不允许返回值。

Supplier接口是一个不具有参数，并返回参数化类型的一个值的单个方法的通用功能接口。

这允许将Supplier的实例作为lambda表达式提供，该表达式执行计算并返回结果：

` CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> "Hello" ); // ... assertEquals( "Hello" , future.get()); 复制代码`

## thenRun()使用 ##

在两个任务任务A，任务B中，如果既不需要任务A的值也不想在任务B中引用，那么你可以将Runnable lambda 传递给 ` thenRun()` 方法。在下面的示例中，在调用future.get()方法之后，我们只需在控制台中打印一行：

模板

` CompletableFuture.runAsync(() -> {}).thenRun(() -> {}); CompletableFuture.supplyAsync(() -> "resultA" ).thenRun(() -> {}); 复制代码`

* 第一行用的是 ` thenRun(Runnable runnable)` ，任务 A 执行完执行 B，并且 B 不需要 A 的结果。
* 第二行用的是 ` thenRun(Runnable runnable)` ，任务 A 执行完执行 B，会返回 ` resultA` ，但是 B 不需要 A 的结果。

实战

` CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello" ); CompletableFuture<Void> future = completableFuture .thenRun(() -> System.out.println( "Computation finished." )); future.get(); 复制代码`

## thenAccept()使用 ##

在两个任务任务A，任务B中，如果你不需要在Future中有返回值，则可以用 ` thenAccept` 方法接收将计算结果传递给它。最后的future.get（）调用返回Void类型的实例。

模板

` CompletableFuture.runAsync(() -> {}).thenAccept(resultA -> {}); CompletableFuture.supplyAsync(() -> "resultA" ).thenAccept(resultA -> {}); 复制代码`

* 第一行中， ` runAsync` 不会有返回值，第二个方法 ` thenAccept` ，接收到的resultA值为null，同时任务B也不会有返回结果
* 第二行中， ` supplyAsync` 有返回值，同时任务B不会有返回结果。

实战

` CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello" ); CompletableFuture<Void> future = completableFuture .thenAccept(s -> System.out.println( "Computation returned: " + s)); future.get(); 复制代码`

## thenApply()使用 ##

在两个任务任务A，任务B中，任务B想要任务A计算的结果，可以用 ` thenApply` 方法来接受一个函数实例，用它来处理结果，并返回一个Future函数的返回值：

模板

` CompletableFuture.runAsync(() -> {}).thenApply(resultA -> "resultB" ); CompletableFuture.supplyAsync(() -> "resultA" ).thenApply(resultA -> resultA + " resultB" ); 复制代码`

* 第二行用的是 thenApply(Function fn)，任务 A 执行完执行 B，B 需要 A 的结果，同时任务 B 有返回值。

实战

` CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello" ); CompletableFuture<String> future = completableFuture .thenApply(s -> s + " World" ); assertEquals( "Hello World" , future.get()); 复制代码`

当然，多个任务的情况下，如果任务 B 后面还有任务 C，往下继续调用 .thenXxx() 即可。

## thenCompose()使用 ##

> 
> 
> 
> 接下来会有一个很有趣的设计模式；
> 
> 

CompletableFuture API 的最佳场景是能够在一系列计算步骤中组合CompletableFuture实例。

这种组合结果本身就是CompletableFuture，允许进一步再续组合。这种方法在函数式语言中无处不在，通常被称为 ` monadic设计模式` 。

简单说，Monad就是一种设计模式，表示将一个运算过程，通过函数拆解成互相连接的多个步骤。你只要提供下一步运算所需的函数，整个运算就会自动进行下去。

在下面的示例中，我们使用thenCompose方法按顺序组合两个Futures。

请注意，此方法采用返回CompletableFuture实例的函数。该函数的参数是先前计算步骤的结果。这允许我们在下一个CompletableFuture的lambda中使用这个值：

` CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello" ) .thenCompose(s -> CompletableFuture.supplyAsync(() -> s + " World" )); assertEquals( "Hello World" , completableFuture.get()); 复制代码`

该thenCompose方法连同thenApply一样实现了结果的合并计算。但是他们的内部形式是不一样的，它们与Java 8中可用的Stream和Optional类的map和flatMap方法是有着类似的设计思路在里面的。

两个方法都接收一个CompletableFuture并将其应用于计算结果，但thenCompose（flatMap）方法接收一个函数，该函数返回相同类型的另一个CompletableFuture对象。此功能结构允许将这些类的实例继续进行组合计算。

## thenCombine() ##

> 
> 
> 
> 取两个任务的结果
> 
> 

如果要执行两个独立的任务，并对其结果执行某些操作，可以用Future的thenCombine方法：

模板

` CompletableFuture<String> cfA = CompletableFuture.supplyAsync(() -> "resultA" ); CompletableFuture<String> cfB = CompletableFuture.supplyAsync(() -> "resultB" ); cfA.thenAcceptBoth(cfB, (resultA, resultB) -> {}); cfA.thenCombine(cfB, (resultA, resultB) -> "result A + B" ); 复制代码`

实战

` CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello" ) .thenCombine(CompletableFuture.supplyAsync( () -> " World" ), (s1, s2) -> s1 + s2)); assertEquals( "Hello World" , completableFuture.get()); 复制代码`

更简单的情况是，当你想要使用两个Future结果时，但不需要将任何结果值进行返回时，可以用 ` thenAcceptBoth` ，它表示后续的处理不需要返回值，而 thenCombine 表示需要返回值：

` CompletableFuture future = CompletableFuture.supplyAsync(() -> "Hello" ) .thenAcceptBoth(CompletableFuture.supplyAsync(() -> " World" ), (s1, s2) -> System.out.println(s1 + s2)); 复制代码`

# thenApply()和thenCompose()之间的区别 #

在前面的部分中，我们展示了关于thenApply()和thenCompose()的示例。这两个API都是使用的CompletableFuture调用，但这两个API的使用是不同的。

## thenApply() ##

此方法用于处理先前调用的 **结果** 。但是，要记住的一个关键点是返回类型是转换泛型中的类型，是同一个CompletableFuture。

因此，当我们想要转换CompletableFuture 调用的结果时，效果是这样的 ：

` CompletableFuture<Integer> finalResult = compute().thenApply(s-> s + 1); 复制代码`

## thenCompose() ##

该thenCompose()方法类似于thenApply()在都返回一个新的计算结果。但是，thenCompose()使用前一个Future作为参数。它会直接使结果变新的Future，而不是我们在thenApply()中到的嵌套Future，而是用来连接两个CompletableFuture，是生成一个新的CompletableFuture：

` CompletableFuture<Integer> computeAnother(Integer i){ return CompletableFuture.supplyAsync(() -> 10 + i); } CompletableFuture<Integer> finalResult = compute().thenCompose(this::computeAnother); 复制代码`

因此，如果想要继续嵌套链接 CompletableFuture 方法，那么最好使用 thenCompose() 。

# 并行运行多个任务 #

当我们需要并行执行多个任务时，我们通常希望等待所有它们执行，然后处理它们的组合结果。

该 ` CompletableFuture.allOf` 静态方法允许等待所有的完成任务：

API

` public static CompletableFuture<Void> allOf(CompletableFuture<?>... cfs){...} 复制代码`

实战

` CompletableFuture<String> future1 = CompletableFuture.supplyAsync(() -> "Hello" ); CompletableFuture<String> future2 = CompletableFuture.supplyAsync(() -> "Beautiful" ); CompletableFuture<String> future3 = CompletableFuture.supplyAsync(() -> "World" ); CompletableFuture<Void> combinedFuture = CompletableFuture.allOf(future1, future2, future3); // ... combinedFuture.get(); assertTrue(future1.isDone()); assertTrue(future2.isDone()); assertTrue(future3.isDone()); 复制代码`

请注意，CompletableFuture.allOf()的返回类型是CompletableFuture 。这种方法的局限性在于它不会返回所有任务的综合结果。相反，你必须手动从Futures获取结果。幸运的是，CompletableFuture.join()方法和Java 8 Streams API可以解决：

` String combined = Stream.of(future1, future2, future3) .map(CompletableFuture::join) .collect(Collectors.joining( " " )); assertEquals( "Hello Beautiful World" , combined); 复制代码`

CompletableFuture 提供了 join() 方法，它的功能和 get() 方法是一样的，都是阻塞获取值，它们的区别在于 join() 抛出的是 unchecked Exception。这使得它可以在Stream.map（）方法中用作方法引用。

# 异常处理 #

说到这里，我们顺便来说下 CompletableFuture 的异常处理。这里我们要介绍两个方法：

` public CompletableFuture<T> exceptionally(Function<Throwable, ? extends T> fn); public <U> CompletionStage<U> handle(BiFunction<? super T, Throwable, ? extends U> fn); 复制代码`

看下代码

` CompletableFuture.supplyAsync(() -> "resultA" ) .thenApply(resultA -> resultA + " resultB" ) .thenApply(resultB -> resultB + " resultC" ) .thenApply(resultC -> resultC + " resultD" ); 复制代码`

上面的代码中，任务 A、B、C、D 依次执行，如果任务 A 抛出异常（当然上面的代码不会抛出异常），那么后面的任务都得不到执行。如果任务 C 抛出异常，那么任务 D 得不到执行。

那么我们怎么处理异常呢？看下面的代码，我们在任务 A 中抛出异常，并对其进行处理：

` CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> { throw new RuntimeException(); }) .exceptionally(ex -> "errorResultA" ) .thenApply(resultA -> resultA + " resultB" ) .thenApply(resultB -> resultB + " resultC" ) .thenApply(resultC -> resultC + " resultD" ); System.out.println(future.join()); 复制代码`

上面的代码中，任务 A 抛出异常，然后通过 `.exceptionally()` 方法处理了异常，并返回新的结果，这个新的结果将传递给任务 B。所以最终的输出结果是：

> 
> 
> 
> errorResultA resultB resultC resultD
> 
> 

` String name = null; // ... CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> { if (name == null) { throw new RuntimeException( "Computation error!" ); } return "Hello, " + name; })}).handle((s, t) -> s != null ? s : "Hello, Stranger!" ); assertEquals( "Hello, Stranger!" , completableFuture.get()); 复制代码`

当然，它们也可以都为 null，因为如果它作用的那个 CompletableFuture 实例没有返回值的时候，s 就是 null。

# Async后缀方法 #

CompletableFuture 类中的API的大多数方法都有两个带有 Async 后缀的附加修饰。这些方法表示用于异步线程。

没有 Async 后缀的方法使用调用线程运行下一个执行线程阶段。不带 Async 方法使用 ForkJoinPool.commonPool() 线程池的 fork / join 实现运算任务。带有 Async 方法使用传递式的 Executor 任务去运行。

下面附带一个案例，可以看到有 thenApplyAsync 方法。在程序内部，线程被包装到 ForkJoinTask 实例中。这样可以进一步并行化你的计算并更有效地使用系统资源。

` CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello" ); CompletableFuture<String> future = completableFuture .thenApplyAsync(s -> s + " World" ); assertEquals( "Hello World" , future.get()); 复制代码`

# JDK 9 CompletableFuture API #

在Java 9中， CompletableFuture API通过以下更改得到了进一步增强：

* 新工厂方法增加了
* 支持延迟和超时
* 改进了对子类化的支持。

引入了新的实例API：

* Executor defaultExecutor()
* CompletableFuture newIncompleteFuture()
* CompletableFuture copy()
* CompletionStage minimalCompletionStage()
* CompletableFuture completeAsync(Supplier<? extends T> supplier, Executor executor)
* CompletableFuture completeAsync(Supplier<? extends T> supplier)
* CompletableFuture orTimeout(long timeout, TimeUnit unit)
* CompletableFuture completeOnTimeout(T value, long timeout, TimeUnit unit)

还有一些静态实用方法：

* Executor delayedExecutor(long delay, TimeUnit unit, Executor executor)
* Executor delayedExecutor(long delay, TimeUnit unit)
* CompletionStage completedStage(U value)
* CompletionStage failedStage(Throwable ex)
* CompletableFuture failedFuture(Throwable ex)

最后，为了解决超时问题，Java 9又引入了两个新功能：

* orTimeout()
* completeOnTimeout()

# 结论 #

在本文中，我们描述了CompletableFuture类的方法和典型用例。