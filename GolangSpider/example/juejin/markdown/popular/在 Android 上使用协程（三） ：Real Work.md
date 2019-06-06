# 在 Android 上使用协程（三） ：Real Work #

这里是关于在 Android 上使用协程的一系列文章。本篇文章将着重于介绍使用协程来解决实际问题。

该系列其他文章：

> 
> 
> 
> [在 Android 上使用协程（一）：Getting The Background](
> https://juejin.im/post/5cea3ee0f265da1bca51b841 )
> 
> 
> 
> [在 Android 上使用协程（二）：Getting started](
> https://juejin.im/post/5cee800051882544171c5a2c )
> 
> 

## 使用协程解决现实问题 ##

系列前两篇文章着重于介绍协程如何简化代码，在 Android 上提供主线程安全，避免泄露任务。以此为背景，对于在 Android 中处理后台任务和简化回调代码，这都是一个很好的解决方案。

到目前为止，我们了解了什么是协程以及如何管理它们。在这篇文章中，我们将看一下如何使用协程来完成真实的任务。协程是一种通用的编程语言特性，和函数同一级别，所以你可以使用协程来实现任何对象或函数可以完成的工作。然而，对下面这两种真实代码中经常出现的任务来说，协程是一个很好的解决方案。

* **一次性请求 ：** 调用一次执行一次，它们总是在结果准备好之后才结束执行。
* **流式请求 ：** 观察变化并反馈给调用者，它们直到第一个结果返回才会结束执行。

协程很好的解决了上面这些任务。这篇文章中，我们会深入一次性请求，探讨在 Android 上如何实现它。

## 一次性请求 ##

一次性请求每调用一次就会执行一次，结果一旦准备好就会结束执行。这和普通函数调用是一样的模式 —— 调用，做一些工作，返回。由于其和函数调用的相似性，它比流式请求更容易理解。

> 
> 
> 
> 一次性请求每次调用时执行，结果一旦准备好就会停止执行。
> 
> 

举个一次性请求的例子，想象一下你的浏览器是如何加载网页的。当你点击链接时，向服务器发送了一个网络请求来加载网页。一旦数据传输到了你的浏览器，它就停止与后端的交互了，此时它已经拥有了需要的所有数据。如果服务器修改了数据，新的修改不会在浏览器展示，你必须刷新页面。

所以，即使一次性请求缺少流式请求的实时推送功能，但它仍然很强大。在 Android 上，你可以使用一次性请求做很多事情，例如查询，存储或者更新数据。对于列表排序来说，它也是一种好方案。

## 问题：展示有序列表 ##

让我们通过展示有序列表来探索一次性请求。为了使例子更加具体，我们编写一个产品库存应用给商店的员工使用。它被用于根据最后一次进货的时间来查询货物。货物既可以升序排列，也可以降序排列。这儿的货物太多了以至于排序花费了几乎一秒，让我们使用协程来避免阻塞主线程。

这个 App 中的所有产品都存储在数据库 Room 中。这是一个很好的例子，因为我们不需要进行网络请求，这样我们就可以专注于设计模式。由于无需网络请求使得这个例子很简单，尽管这样，但是它仍然展示了实现一次性请求所使用的模式。

为了使用协程实现这个请求，你需要把协程引入 ` ViewModel` 、 ` Repository` 和 ` Dao` 。让我们逐个看看它们是如何与协程结合在一起的。

` class ProductsViewModel ( val productsRepository: ProductsRepository) : ViewModel() { private val _sortedProducts = MultableLiveData<List<ProductListing>>() val sortedProducts: LiveData<List<ProductListing>> = _sortedProducts /** * Called by the UI when the user clicks the appropriate sort button */ fun onSortAscending () = sortPricesBy(ascending = true ) fun onSortDescending () = sortPricesBy(ascending = false ) private fun sortPricesBy (ascending: Boolean ) { viewModelScope.launch { // suspend and resume make this database request main-safe // so our ViewModel doesn't need to worry about threading _sortedProducts.value = productsRepository.loadSortedProducts(ascending) } } } 复制代码`

` ProductsViewModel` 负责接收用户层事件，然后请求 repository 更新数据。它使用 ` LiveData` 存储要在 UI 中进行展示的当前有序列表。当接收到一个新的事件， ` sortPricesBy` 方法会开启一个新的协程来排序集合，当结果可用时更新 ` LiveData` 。由于 ` ViewModel` 可以在 ` onCleared` 回调中取消协程，所以它是这个架构中启动协程的好位置。当用户离开界面的时候，就无需再继续未完成的任务了。

如果你不是很了解 LiveData，这里有一篇介绍 LiveData 如何为 UI 层存储数据的好文章，作者是 [CeruleanOtter]( https://link.juejin.im?target=https%3A%2F%2Ftwitter.com%2FCeruleanOtter ) 。

> 
> 
> 
> [ViewModels: A Simple Example](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fandroiddevelopers%2Fviewmodels-a-simple-example-ed5ac416317e
> )
> 
> 

这是在 Android 上使用协程的通用模式。由于 Android Framework 无法调用 suspend 函数，你需要配合一个协程来响应 UI 事件。最简单的方法就是当事件发生时启动一个新的协程，最适合的地方就是 ` ViewModel` 了。

> 
> 
> 
> 在 ViewModel 中启动协程是一个通用的设计模式。
> 
> 

` ViewModel` 实际上通过 ` ProductsRepository` 来获取数据。让我们来看一下代码：

` class ProductsRepository ( val productsDao: ProductsDao) { /** * This is a "regular" suspending function, which means the caller must * be in a coroutine. The repository is not responsible for starting or * stoppong coroutines since it doesn't have a natural lifecycle to cancel * unnecssary work. * * This *may* be called from Dispatchers.Main abd is main-safe because * Room will take care of main-safety for us. */ suspend fun loadSortedProducts (ascending: Boolean ) : List<ProductListing> { return if (ascending) { productsDao.loadProductsByDateStockedAscending() } else { productsDao.loadProductsByDateStockedDescending() } } } 复制代码`

` ProductsRepository` 为商品数据的交互提供了合理的接口。在这个 App 中，由于所有数据都是存储在 Room 数据库中，它提供了具有两个针对不同排序的方法的 ` Dao` 。

Repository 是 Android Architecture Components 架构的可选部分。如果你在 app 中使用了 repository 或者相似作用的层级，它更偏向于使用挂起函数。由于 repository 没有生命周期，它仅仅只是一个对象，所有它没有办法做资源清理工作。在 repository 中启动的协程将有可能泄露。

使用挂起函数，除了避免泄露以外，在不同上下文中也可以重复使用 repository 。任何知道如何创建协程的都可以调用 ` loadSortedProducts` ，例如 ` WorkManager` 库启动了后台任务。

> 
> 
> 
> Repository 应该使用挂起函数来保证主线程安全。
> 
> 

> 
> 
> 
> 注意： 当用户离开界面时，一些后台执行的保存操作可能想继续运行，这种情况下，脱离生命周期运行是有意义的。在大多数情况下， `
> viewModelScope` 都是一个好选择。
> 
> 

再来看看 ` ProductsDao` ：

` @Dao interface ProductsDao { // Because this is marked suspend, Room will use it's own dispatcher // to run this query in a main-safe way, @Query( "select * from ProductListing ORDER BY dataStocked ASC" ) suspend fun loadProductsByDateStockedAsceding () : List<ProductListing> // Because this is marked suspend, Room will use it's own dispatcher // to run this query in a main-safe way, @Query( "select * from ProductListing ORDER BY dataStocked DESC" ) suspend fun loadProductsByDateStockedDesceding () : List<ProductListing> } 复制代码`

` ProductsDao` 是一个 Room [Dao]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Freference%2Fandroid%2Farch%2Fpersistence%2Froom%2FDao ) ，它对外提供了两个挂起函数。由于函数由 suspend 修饰，Room 会确保它们主线程安全。这就意味着你可以直接在 ` Dispatchers.Main` 中调用它们。

如果你没有在 Room 中使用过协程，阅读一下 [FMuntenescu]( https://link.juejin.im?target=https%3A%2F%2Ftwitter.com%2FFMuntenescu ) 的这篇文章:

> 
> 
> 
> [Room && Coroutines](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fandroiddevelopers%2Froom-coroutines-422b786dc4c5
> )
> 
> 

不过需要注意这一点，调用它的协程将运行在主线程。所以如果你要对结果进行一些昂贵的操作，例如转换成集合，你要确保不会阻塞主线程。

> 
> 
> 
> 注意：Room 使用自己的调度器在后台线程进行查询操作。你不应该再使用 ` withContext(Dispatchers.IO)` 来调用 Room
> 的 suspend 查询，这只会让你的代码运行的更慢。
> 
> 

> 
> 
> 
> Room 中的挂起函数是主线程安全的，它运行在自定义的调度器中。
> 
> 

## 一次性请求模式 ##

这就是在 Android Architecture Components 中使用协程进行一次性请求的完整模式。我们将协程添加到 ` ViewModel` 、 ` Repository` 和 ` Room` 中，每一层都有不同的责任。

* ViewModel 在主线程启动协程，一旦有了结果就结束。
* Repository 提供挂起函数并保证它们主线程安全。
* 数据库和网络层提供挂起函数并保证它们主线程安全。

` ViewModel` 负责启动协程，保证用户离开界面时取消协程。它本身不做昂贵的操作，而是依赖其他层来做。一旦有了结果，就使用 ` LiveData` 发送给 UI 界面。也正因为 ` ViewModel` 不做昂贵的操作，所以它在主线程启动协程。通过在主线程启动，当结果可用它可以更快的响应用户事件（例如内存缓存）。

` Repository` 提供挂起函数来访问数据。它通常不会启动长生命周期的协程，因为它没有办法取消它们。无论何时 ` Repository` 需要做昂贵的操作（集合转换等），它都需要使用 ` withContext` 来提供主线程安全的接口。

` 数据层` （网络或者数据库）总是提供挂起函数。使用 Kotlin 协程的时候需要保证这些挂起函数是主线程安全的，Room 和 Retrofit 都遵循了这一原则。

在一次性请求中，数据层只提供挂起函数。如果想要获取新值，就必须再调用一次。这就像浏览器中的刷新按钮。

花点时间让你明白一次性请求的模式是值得的。这在 Android 协程中是通用的模式，你也会一直使用它。

## 第一个 Bug Report ##

在测试过该解决方案之后，你将其用到生产环境，几周内都运行良好，直到你收到了一个非常奇怪的错误报告：

> 
> 
> 
> Subject: 🐞 — 排序错误！
> 
> 
> 
> Report: 当我非常非常非常非常快速点击排序按钮时，排序偶尔是错误的。这并不是每次都会发生。
> 
> 

你看了看，挠挠头，哪里可能发生错误了呢？这个逻辑看起来相当简单：

* 开始用户请求的排序
* 在 Room 调度器中开始排序
* 展示排序结果

你正准备关闭这个 bug，关闭理由是 “不予处理 —— 不要快速点击按钮”，但是你又担心的确是哪里出了什么问题。在添加了日志以及编写测试用例来测试一次性发起许多排序请求，你最终找到了原因。

最后获得的结果实际上并不是 “排序的结果”，而是 “完成最后一次排序时” 的结果。当用户狂点按钮时，同时发起了多次排序，可能以任意顺序结束。（译者注：可以想象成 Java 中的多线程并发）

> 
> 
> 
> 当启动一个新协程来响应用户事件时，要考虑到用户在该协程未结束之前又启动一个协程会发生什么。
> 
> 

这是一个并发导致的 bug，实际上它和协程并没有什么关系。当我们以同样的方式使用回调，Rx，甚至 ` ExecutorService` ，都可能会有这样的 bug。让我们探索一下下面这些方案是如何保证一次性请求按用户所期望的顺序执行的。

## 最佳方案：禁用按钮 ##

核心问题就是我们如何进行两次排序。我们可以让它仅仅只进行一次排序！最简单的方法就是禁用排序按钮，停止发送新事件。

这似乎是一个很简单的方案，但它的确是个好主意。代码实现也很简单，易于测试。

要禁用按钮，可以通知 UI ` sortPricesBy` 中正在进行一次排序请求，如下所示：

` // Solution 0: Disable the sort buttons when any sort is running class ProductsViewModel ( val productsRepository: ProductsRepository): ViewModel() { private val _sortedProducts = MutableLiveData<List<ProductListing>>() val sortedProducts: LiveData<List<ProductListing>> = _sortedProducts private val _sortButtonsEnabled = MutableLiveData< Boolean >() val sortButtonsEnabled: LiveData< Boolean > = _sortButtonsEnabled init { _sortButtonsEnabled.value = true } /** * Called by the UI when the user clicks the appropriate sort button */ fun onSortAscending () = sortPricesBy(ascending = true ) fun onSortDescending () = sortPricesBy(ascending = false ) private fun sortPricesBy (ascending: Boolean ) { viewModelScope.launch { // disable the sort buttons whenever a sort is running _sortButtonsEnabled.value = false try { _sortedProducts.value = productsRepository.loadSortedProducts(ascending) } finally { // re-enable the sort buttons after the sort is complete _sortButtonsEnabled.value = true } } } } 复制代码`

这看起来还不赖。只需在调用 repository 时在 ` sortPricesBy` 内部禁用按钮。

大多数情况下，这都是解决问题的好方案。但是我们想在按钮可用的情况下来解决这个 bug 呢？这有一点点困难，我们将在本文剩余部分来看几种方式。

> 
> 
> 
> Important ：This code shows a major advantage of starting on main — the
> buttons disable instantly in response to a click. If you switched
> dispatchers, a fast-fingered user on a slow phone could send more than one
> click!
> 
> 

## 并发模式 ##

下面几节将探讨一些高级话题。如果你才刚刚开始使用协程，你不必完全理解。简单的禁用按钮就是你遇到的大部分问题的良好解决方案。

在本文的剩余部分，我们将讨论在不禁用按钮的前提下，如何去保证一次性请求正常运行。我们可以通过控制协程何时运行（或者不运行）来避免意外的并发情况。

下面有三种模式，你可以在一次性请求中使用它们来确保同一时间只进行一次请求。

* 在启动更多协程之前先取消上一个。
* 将下一个任务放入等待队列，直到前一个请求执行完成在开始另一个。
* 如果已经有一个请求在运行，那么就返回该请求，而不是启动另一个请求。

想一下这些解决方案，你会发现它们的实现相对都比较复杂。为了专注于设计模式而不是实现细节，我创建了 [gist]( https://link.juejin.im?target=https%3A%2F%2Fgist.github.com%2Fobjcode%2F7ab4e7b1df8acd88696cb0ccecad16f7%23file-concurrencyhelpers-kt-L19 ) 来提供这三种模式的实现作为可用抽象。（这里可以大概浏览一下 gist 中的代码实现）

### 方案一 : 取消前一个任务 ###

在排序的情况下，从用户那获取了一个新的事件，就意味着你可以取消上一个请求了。毕竟，用户已经不想知道上一个任务的结果了，继续下去还有什么意义呢？

为了取消上一个请求，我们首先要以某种方式追踪它。 中的 ` cancelPreviousThenRun` 函数就是这么做的。

让我们看看它是如何被用来修复 bug 的：

` // Solution #1: Cancel previous work // This is a great solution for tasks like sorting and filtering that // can be cancelled if a new request comes in. class ProductsRepository ( val productsDao: ProductsDao, val productsApi: ProductsService) { var controlledRunner = ControlledRunner<List<ProductListing>>() suspend fun loadSortedProducts (ascending: Boolean ) : List<ProductListing> { // cancel the previous sorts before starting a new one return controlledRunner.cancelPreviousThenRun { if (ascending) { productsDao.loadProductsByDateStockedAscending() } else { productsDao.loadProductsByDateStockedDescending() } } } } 复制代码`

看一下 gist 中 ` cancelPreviousThenRun` 中的 [实现]( https://link.juejin.im?target=https%3A%2F%2Fgist.github.com%2Fobjcode%2F7ab4e7b1df8acd88696cb0ccecad16f7%23file-concurrencyhelpers-kt-L91 ) ，你可以了解到它是如何追踪正在工作的任务的。

` // see the complete implementation at // https://gist.github.com/objcode/7ab4e7b1df8acd88696cb0ccecad16f7 suspend fun cancelPreviousThenRun (block: suspend () -> T): T { // If there is an activeTask, cancel it because it's result is no longer needed activeTask?.cancelAndJoin() // ... 复制代码`

简而言之，它总是追踪成员变量 ` activeTask` 中的当前排序。无论何时开始一次新的排序，都会立即 [cancelAndJoin]( https://link.juejin.im?target=https%3A%2F%2Fkotlin.github.io%2Fkotlinx.coroutines%2Fkotlinx-coroutines-core%2Fkotlinx.coroutines%2Fcancel-and-join.html ) ` activeTask` 中的所有内容。这会造成的影响就是，在开启一次新的排序之前会取消所有正在进行的排序。

使用类似 ` ControlledRunner<T>` 的抽象实现来封装逻辑是个好方法，而不是将并发性和程序逻辑混杂在一起。

> 
> 
> 
> 重要：这个模式不适合在全局单例中使用，因为不相关的调用者不应该互相取消。
> 
> 

### 方案二 ：将下一个任务入队 ###

这里有一个对于并发 bug 总是有效的解决方案。

只需要将请求排队，这样同时只会进行一个请求。就像商店中排队一样，请求将按它们排队的顺序依次执行。

对于这种特定的排队问题，取消可能比排队更好。但值得一提的是它总是可以保证正常工作。

` // Solution #2: Add a Mutex // Note: This is not optimal for the specific use case of sorting // or filtering but is a good pattern for network saves. class ProductsRepository ( val productsDao: ProductsDao, val productsApi: ProductsService) { val singleRunner = SingleRunner() suspend fun loadSortedProducts (ascending: Boolean ) : List<ProductListing> { // wait for the previous sort to complete before starting a new one return singleRunner.afterPrevious { if (ascending) { productsDao.loadProductsByDateStockedAscending() } else { productsDao.loadProductsByDateStockedDescending() } } } } 复制代码`

无论何时进行一次新的排序，它使用一个 ` SingleRunner` 实例来确保同时只进行一个排序任务。

它使用了 [Mutex]( https://link.juejin.im?target=https%3A%2F%2Fgist.github.com%2Fobjcode%2F7ab4e7b1df8acd88696cb0ccecad16f7%23file-concurrencyhelpers-kt-L49 ) ，Mutex（互斥锁） 是一个单程票，或者说锁，协程必须获取锁才能进入代码块。如果一个协程在运行时另一个协程尝试进入，它将挂起自己直到所有等待的协程都完成。

> 
> 
> 
> Mutex 保证同时只有一个协程运行，并且它们将按启动的顺序结束。
> 
> 

### 方案三 ：加入前一个任务 ###

第三种解决方案是加入前一个任务。如果新请求可以重复使用已经存在的，已经完成了一半的相同的任务，这会是一个好主意。

这种模式对于排序功能来说并没有太大意义，但是对于网络请求来说是很适用的。

对于我们的产品库存应用，用户需要一种方式来从服务器获取最新的产品库存数据。我们提供了一个刷新按钮，用户可以点击来发起一次新的网络请求。

就和排序按钮一样，当请求正在进行的时候，禁用按钮就可以解决问题。但是如果我们不想这样，或者不能这样，我们可以选择加入已经存在的请求。

查看 gist 中使用 [joinPreviousOrRun]( https://link.juejin.im?target=https%3A%2F%2Fgist.github.com%2Fobjcode%2F7ab4e7b1df8acd88696cb0ccecad16f7%23file-concurrencyhelpers-kt-L124 ) 的代码，看看它是如何工作的：

` class ProductsRepository ( val productsDao: ProductsDao, val productsApi: ProductsService) { var controlledRunner = ControlledRunner<List<ProductListing>>() suspend fun fetchProductsFromBackend () : List<ProductListing> { // if there's already a request running, return the result from the // existing request. If not, start a new request by running the block. return controlledRunner.joinPreviousOrRun { val result = productsApi.getProducts() productsDao.insertAll(result) result } } } 复制代码`

这与 ` cancelPreviousAndRun` 的行为相反。 ` cancelPreviousAndRun` 会通过取消直接放弃前一个请求，而 ` joinPreviousOrRun` 将会放弃新请求。如果已经存在正在运行的请求，它将会等待执行结果并返回，而不是发起一次新的请求。只有在没有正在运行的请求时才会执行代码块。

在下面的代码中你可以看到 ` joinPreviousOrRun` 中的任务是如何工作的。它仅仅只是当 ` activeTask` 中存在任务的时候，直接返回前一个请求的结果。

` // see the complete implementation at // https://gist.github.com/objcode/7ab4e7b1df8acd88696cb0ccecad16f7#file-concurrencyhelpers-kt-L124 suspend fun joinPreviousOrRun (block: suspend () -> T): T { // if there is an activeTask, return it's result and don't run the block activeTask?.let { return it.await() } // ... 复制代码`

这个模式很适合通过 id 查询产品的请求。你可以使用 map 来保存 ` id` 到 ` Deferred` 的映射关系，然后使用相同的 join 逻辑来追踪同一个产品的之前的请求。

> 
> 
> 
> 加入前面的任务可以有效避免重复的网络请求。
> 
> 

## What's next ?

在这篇文章中，我们探讨了如何使用 Kotlin 协程来实现一次性请求。首先我们通过在 ` ViewModel` 中启动协程，通过 ` Repository` 和 Room ` Dao` 提供公开的挂起函数来实现了一个完整的设计模式。

对于大多数任务，为了在 Android 上使用 Kotlin 协程，这就是全部你所需要做的。这个模式可以应用在许多场景，就像上面说过的排序。你也可以使用它查询，保存，更新网络数据。

然后我们看了一个可能出现的 bug 及其解决方案。最简单的（经常也是最好的）方案就是从 UI 上修改，当一个排序正在运行时直接禁用排序按钮。

最后，我们研究了一些高级并发模式，以及如何在 Kotlin 协程中实现。 [代码]( https://link.juejin.im?target=https%3A%2F%2Fgist.github.com%2Fobjcode%2F7ab4e7b1df8acd88696cb0ccecad16f7%23file-concurrencyhelpers-kt-L158 ) 有点复杂，但为一些高级协程方面的话题提供了很好的介绍。

下一篇中，让我们进入流式请求，以及如何使用 ` liveData` 构建器 ！

> 
> 
> 
> 文章首发微信公众号： **` 秉心说`** ， 专注 Java 、 Android 原创知识分享，LeetCode 题解。
> 
> 
> 
> 更多原创文章，扫码关注我吧！
> 
> 

![](https://user-gold-cdn.xitu.io/2019/4/27/16a5f352eab602c4?imageView2/0/w/1280/h/960/ignore-error/1)