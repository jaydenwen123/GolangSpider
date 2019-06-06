# [译]Android 中 Kotlin 与 RecyclerView 高性能列表 #

### 翻译说明: ###

原标题: Kotlin & RecyclerView for High Performance Lists in Android

原文地址: [www.andreasjakl.com]( https://link.juejin.im?target=https%3A%2F%2Fwww.andreasjakl.com%2Fkotlin-recyclerview-for-high-performance-lists-in-android%2F )

原文作者: Andreas Jakl

RecyclerView 是在 Android 上显示滚动列表的最佳方法。它确保了高性能和平滑滚动，同时提供具有灵活布局的列表元素。结合 [Kotlin]( https://link.juejin.im?target=https%3A%2F%2Fwww.kotlincn.net%2F ) 的现代语言功能，与传统的 Java 方法相比，RecyclerView 的代码开销大大降低。

## 示例项目：PartsList - 入门 ##

在本文中，我们将介绍一个示例场景：维护应用程序的滚动列表，列出机器部件：“PartsList”。但是，此方案仅影响我们使用的字符串 - 您可以针对您需要的任何用例复制此方法。

要开始使用，请使用 Android Studio 3+ 创建一个新的 Android 应用。确保启用 Kotlin 支持，并为 MainActivity 选择 “Empty” 模板。或者，如果您不想手动编写以下步骤的代码，请从 [GitHub]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fandijakl%2FPartsList%2F ) 下载完成的 [源码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fandijakl%2FPartsList%2F ) 。

## 什么是 RecyclerView？ ##

Android 中的屏幕列表由多个子视图组成。每个都有一个或多个视图的布局。例如，电子邮件应用会向您显示多封电子邮件; 这些项目中的每一项都包括主题，发件人姓名和一堆其他信息。

解析子视图的 XML 布局并将其扩展到类的实例中是一项消耗性能的操作。快速滑动会对手机性能造成巨大压力。目标是始终坚持每秒60帧。但是，这每帧只剩下不到17毫秒的计算时间。

**RecyclerView** 的主要技巧是重新使用列表中的子视图。一旦子视图 滚出可见区域，它就基本上被放入队列中。迟早，当新的子视图滚动时会再次用到它。然后，子视图的 UI 文本内容将被替换。下图显示了这个（简化的）RecyclerView 原理：

![](https://user-gold-cdn.xitu.io/2019/6/1/16b130fb44c0945d?imageView2/0/w/1280/h/960/ignore-error/1)

## RecyclerView流程 ##

不幸的是，这种有效的方法需要一些架构背景。刚开始看起来可能感到陌生。然而，一旦您拥有所有组件，就可以轻松的进行自定义。

![](https://user-gold-cdn.xitu.io/2019/6/1/16b1355fc67fe56e?imageView2/0/w/1280/h/960/ignore-error/1)

使用RecyclerView需要配置/实现以下组件：

* RecyclerView：管理一切。它主要由Android预先编写。您提供组件和配置。
* Adapter：您将花费大部分时间编写此类。它连接到数据源。在RecyclerView的指示下，它会创建或更新列表中的各个项目。
* ViewHolder：一个简单的类，用于分配/更新视图项中的数据。重新使用视图时，将覆盖先前的数据。
* Data source：您喜欢的任何东西 - 从简单的数组到完整的数据源。您的适配器与之交互。
* LayoutManager：负责将所有单独的视图项放在屏幕上，并确保它们获得所需的屏幕空间。

## 数据源和Kotlin数据类 ##

我们的数据源是一个简单的列表/数组。该列表由自定义类的实例组成。您可以在应用程序中对这些进行硬编码，也可以动态创建它们 - 例如，基于来自在线源的 HTTP REST 请求。

Kotlin有一个很好的功能，可以简化 [数据类]( https://link.juejin.im?target=https%3A%2F%2Fwww.kotlincn.net%2Fdocs%2Freference%2Fdata-classes.html ) 的编写。此方法适用于仅包含数据但不包含任何操作的所有类。我们的示例数据类称为 ` PartData` ，因为它存储有关维护应用程序的计算机部件的信息。

` data class PartData ( val id: Long, val itemName: String) 复制代码`

对于数据类，Kotlin自动生成所有实用程序函数：

* 属性：您不再需要像普通 Java 那样编写 ` getter` 和 ` setter` 。
* 附加功能：例如， ` equals()` ， ` hasCode()` ， ` toString()` 或 ` copy()` 。

[Antonio Leiva]( https://link.juejin.im?target=https%3A%2F%2Fantonioleiva.com%2Fdata-classes-kotlin%2F ) 写了一篇很棒的文章，在相同数据类的情况下将 Java 与 Kotlin 进行比较。您可以节省大约80％的代码并仍然得到相同的结果。

您像任何其他类一样实例化 Kotlin 数据类。请注意，Kotlin 不使用 ` new` 语句，因此我们只使用类名并将值提供给构造函数参数。添加该代码的 ` nCreate()` 您的 MainActivity。

` var partList = ArrayList<PartData>() partList.add(PartData(100411, "LED Green 568 nm, 5mm" )) partList.add(PartData(101119, "Aluminium Capacitor 4.7μF" )) partList.add(PartData(101624, "Potentiometer 500kΩ" )) // ... 复制代码`

## 列表项目布局 ##

RecyclerView 中的每个项目都需要自定义布局。您可以使用任意与 Activity 布局相同的方式创建这个文件：在 Android Studio 的 Android-style 项目视图中，打开 app > res > layout。右键单击 ` "layout"` 文件夹名称，然后选择 New > Layout resource file。创建一个名为 ` part_list_item.xml` 的布局。

对于示例用例，请使用以下属性：

* Layout: ` LinearLayout` (垂直)
* Width: ` match_parent`
* Height: ` wrap_content`
* Padding: ` 16dp`

接下来，将两个子项添加到布局中。我们将使用 TextView的。只要确保你给他们有用的ID，因为我们需要这些来分配Kotlin代码中的文本：

* Id 1: @+id/tv_part_item_name (larger text size)
* Id 2: @+id/tv_part_id (smaller text size)

![](https://user-gold-cdn.xitu.io/2019/6/1/16b1369ce1cb137b?imageView2/0/w/1280/h/960/ignore-error/1)

## Recycler View Dependencies ##

使用 Android 5.0（API级别21）将 [RecyclerView]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fguide%2Ftopics%2Fui%2Flayout%2Frecyclerview.html ) 添加到 Android 系统中。许多应用程序仍然针对 Android 4.4+，因此您通常会通过 [AppCompat / Support]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Ftraining%2Fmaterial%2Fcompatibility.html%23SupportLib ) 库使用 RecyclerView 。将最新的支持库版本（取决于您的编译SDK）包含到应用程序构建 ` build` 的依赖项中 `.gradle` 文件。

` implementation 'com.android.support:appcompat-v7:26.1.0' implementation 'com.android.support:recyclerview-v7:26.1.0' 复制代码`

在我们添加了这个定义之后，我们可以开始编写使用RecyclerView的源代码了。

## ViewHolder：处理列表中的单个项目 ##

首先，创建一个名为 ` PartAdapter` 的类 。将其放在与 ` MainActivity` 相同的目录中 。在这个类中，创建一个名为 ` PartViewHolder` 的 [嵌套类]( https://link.juejin.im?target=https%3A%2F%2Fkotlinlang.org%2Fdocs%2Freference%2Fnested-classes.html ) 。

` class PartViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) { fun bind (part: PartData) { itemView.tv_part_item_name.text = part.itemName itemView.tv_part_id.text = part.id.toString() } } 复制代码`

该 ViewHolder 介绍了项目视图。此外，它还在 RecyclerView 中存储其位置的元数据。

本质上，我们的 [ViewHolder]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Freference%2Fandroid%2Fsupport%2Fv7%2Fwidget%2FRecyclerView.ViewHolder.html ) 实现的主要任务是将当前需要的数据绑定到先前膨胀的UI布局。每当滚动期间新项目可见时，此类确保项目的视图显示我们期望在列表中此位置的内容。

要更新 UI，我们创建自己的方法并将其命名为 ` bind()` 。请注意，这不是基类的重写方法，因此我们可以给它任何名称。作为方法参数，我们只需要在 UI 中显示数据。在 ` bind()` 中，我们只是将提供的数据分配给视图项。

（从技术上讲，我们没有绑定它，我们只是分配数据。如果数据模型发生变化，它将不会自动更新此实现中的可见列表。）

## Kotlin：主要构造函数 ##

如果您是 Kotlin 的新手，您可能想知道：我们如何在 ` bind()` 中访问 ` itemView` ？它从何而来？

魔术被称为主要构造函数。在 Kotlin 中，它可以成为类标题的一部分。前面的构造函数关键字是可选的。

` class PartViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) { ... } 复制代码`

这有什么好处的？ Kotlin 自动将所有参数作为属性提供。因此，我们没有看到属性定义和初始化。相反，我们只是从 ` bind()` 函数中访问该属性 。

## RecyclerView适配器：将它固定在一起 ##

我们正在探索RecyclerView架构中最重要的组件：适配器。它有3个主要任务：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18180050b4698?imageView2/0/w/1280/h/960/ignore-error/1)

让我们看一下大图中的适配器以及这3个任务适合的位置：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1818415802ee0?imageView2/0/w/1280/h/960/ignore-error/1)

* 为了在屏幕上绘制列表，RecyclerView 将询问适配器总共有多少项。我们的适配器在 getItemCount() 中返回此信息 。
* 每当 RecyclerView 决定它需要内存中的另一个视图实例时，它将调用 onCreateViewHolder()。在此方法中，适配器会膨胀并返回我们在上一步中创建的xml布局。
* 每次（重新）使用先前创建的 ViewHolder 时，RecyclerView 都会指示适配器更新其数据。我们通过重写 onBindViewHolder() 来自定义此过程 。

您不必编写代码来手动覆盖这三个函数。只需扩展 PartAdapter 类的定义 即可。让它继承 RecyclerView 的适配器。Android Studio 会自动提示您实施所有必需的成员：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18190b73566f8?imageView2/0/w/1280/h/960/ignore-error/1)

然后，将参数 “partItemList” 添加到 PartAdapter 类中。使用主构造函数。这将允许 MainActivity 在实例化 PartAdapter 时提供数据模型。

和以前一样， partItemList 参数自动作为属性使用。在我们的例子中，属性类型是我们的数据类的简单列表（数组）。

` class PartAdapter (val partItemList: List<PartData>, val clickListener: (PartData) -> Unit) : RecyclerView.Adapter<RecyclerView.ViewHolder>() { 复制代码`

## RecyclerView的适配器实现 ##

Android Studio 已经创建了方法存根。我们只需编写几行实现。此代码根据我们的要求定制适配器，并完成我所描述的3个主要任务。

` override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): RecyclerView.ViewHolder { // LayoutInflater: takes ID from layout defined in XML. // Instantiates the layout XML into corresponding View objects. // Use context from main app -> also supplies theme layout values! val inflater = LayoutInflater.from(parent.context) // Inflate XML. Last parameter: don 't immediately attach new view to the parent view group val view = inflater.inflate(R.layout.part_list_item, parent, false) return PartViewHolder(view) } override fun onBindViewHolder(holder: RecyclerView.ViewHolder, position: Int) { // Populate ViewHolder with data that corresponds to the position in the list // which we are told to load (holder as PartViewHolder).bind(partItemList[position]) } override fun getItemCount() = partItemList.size 复制代码`

* onCreateViewHolder()：扩展视图项的布局XML。 ` inflate()` 方法的最后一个参数 确保新视图不会立即附加到父视图组。相反，它只在缓存中，不应该是可见的。
* onBindViewHolder()：在 RecyclerView 提供所有必要的信息-在 ViewHolder 实例应显示的数据，以及在数据列表中的（新）位置。使用我们的 Adapter 类的 ` partItemList` 属性，我们调用我们添加到 ViewHolder 的自定义 ` bind()` 函数。
* getItemCount()：在我们的例子中简单 - 列表/数组的大小。在 Kotlin中，我们可以进一步将其简化为内联实现; 如果我们直接返回一个值/一个简单的表达式，我们不需要函数的 {} 和 return 语句。

## RecyclerViews的活动布局 ##

我们差不多完成了。接下来，我们创建 RecyclerView。第一步是将其添加到 MainActivity 的 XML 布局定义中。

如果您使用 Android Studio 的新项目向导并创建了一个空活动，则会添加一个 “Hello World” TextView。删除它。

而是为具有ID “rv_parts” 的滚动全屏 RecyclerView 列表添加以下定义：

` <android.support.v7.widget.RecyclerView android:id= "@+id/rv_parts" android:layout_width= "match_parent" android:layout_height= "match_parent" /> 复制代码`

## 实例化RecyclerView ##

在MainActivity的 ` onCreate()` 中只需要三行代码 。

* 分配LayoutManager：它测量和定位项目视图。不同的布局管理器支持各种布局。我们选择最常见的：LinearLayoutManager。默认情况下，它会创建垂直滚动列表布局。
` rv_parts.layoutManager = LinearLayoutManager(this) 复制代码` * 优化：所述 RecyclerView 的大小是不会被影响适配器内容。RecyclerView 的大小仅取决于父级（在我们的示例中，全屏可用）。 因此，我们可以调用 [setHasFixedSize()]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Freference%2Fandroid%2Fsupport%2Fv7%2Fwidget%2FRecyclerView.html%23setHasFixedSize(boolean) ) 来激活 RecyclerView 中的一些优化。在 Kotlin 中，您不需要 “set”-prefix 来访问属性。
` rv_parts.hasFixedSize() 复制代码` * 分配适配器：这最终连接所有内容：我们创建并分配我们编写的Adapter类的新实例（“PartAdapter”）。该适配器需要的数据项，这是我们在本文开头创建的列表。
` rv_parts.adapter = PartAdapter( test Data) 复制代码`

## RecyclerView：工作！ ##

现在按下播放，您的RecyclerView应该可以正常工作！使用我们提供的三个示例数据项，列表有点太短，无法完全理解滚动。要测试它，只需使用一些额外的项目扩展列表。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b181d0f6c69654?imageView2/0/w/1280/h/960/ignore-error/1)

现在您已经掌握了基础知识，可以很容易地扩展项目布局和数据类以及其他信息。

如果某些内容不适合您，请将您的代码与 [GitHub]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fandijakl%2FPartsList%2F ) 上的完成 [解决方案]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fandijakl%2FPartsList%2F ) 进行比较。请注意，解决方案比我们在此阶段更先进 - 例如，它已经包含单击处理程序。

## 下一步：单击“处理程序” ##

缺少一个经常使用的部分：点击处理程序。与 [ListView]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Freference%2Fandroid%2Fapp%2FListActivity.html%23onListItemClick(android.widget.ListView%2C%2520android.view.View%2C%2520int%2C%2520long) ) 相比，这些在 RecyclerViews 中要复杂得多。但是，如果做得好，他们很容易添加 Kotlin。这只是了解一些更先进的 Kotlin 概念的问题。我们将在 [下一部]( https://link.juejin.im?target=https%3A%2F%2Fwww.andreasjakl.com%2Frecyclerview-kotlin-style-click-listener-android%2F ) 分看一下这个！

## 欢迎关注 Kotlin 中文社区！ ##

### 中文官网： [www.kotlincn.net/]( https://link.juejin.im?target=https%3A%2F%2Fwww.kotlincn.net%2F ) ###

### 中文官方博客： [www.kotliner.cn/]( https://link.juejin.im?target=https%3A%2F%2Fwww.kotliner.cn%2F ) ###

### 公众号：Kotlin ###

### 知乎专栏： [Kotlin]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fbennyhuo ) ###

### CSDN： [Kotlin中文社区]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_23626713 ) ###

### 掘金： [Kotlin中文社区]( https://juejin.im/user/5cea6293e51d45775e33f4dd/posts ) ###

### 简书： [Kotlin中文社区]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fu%2Fa324daa6fa19 ) ###