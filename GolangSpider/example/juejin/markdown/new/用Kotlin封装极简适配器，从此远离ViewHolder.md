# 用Kotlin封装极简适配器，从此远离ViewHolder #

作为一名Android开发者，用过ListView或者RecycleView后想必对ViewHolder再熟悉不过了。ViewHolder 一开始并不是 Android 原生提供的，而是在ListView中作为减少频繁调用 ` findViewById` 而引入的，再到后来推出了更好的 RecycleView，直接内置了ViewHolder。不过我们总归逃脱不了在写适配器时写 ` ViewHolder` 或者 ` findViewById` 的命运。

而现在Kotlin的出现似乎轻而易举地解决了这个问题。你可能还记得，引入Kotlin后，Activity中可以直接用布局文件的Id来使用view，原理可以看下以前写过的一篇文章 [Kotlin直接使用控件ID原理解析]( https://juejin.im/post/5c161975f265da612b137e83 ) ，本文就是用这个特性来封装一个极简地不需要自己创建ViewHolder的通用RecycleView适配器。

### 通用ViewHolder ###

首先Kotlin上述特性在普通View中默认是关闭，打开app的 ` build.gradle` ,启用实验性功能：

` android { ... } androidExtensions { experimental = true } 复制代码`

然后创建一个通用的ViewHolder，很简单只有一行代码：

` class CommonViewHolder (itemView: View) : RecyclerView.ViewHolder(itemView), LayoutContainer { override val containerView: View = itemView } 复制代码`

### 极简适配器 ###

我们直接上代码：

` open class BaseRecyclerAdapter < M > ( @LayoutRes val itemLayoutId: Int , list: Collection<M>? = null , bind: (BaseRecyclerAdapter<M>.() -> Unit )? = null ) : RecyclerView.Adapter<BaseRecyclerAdapter.CommonViewHolder>() { init { if (bind != null ) { apply(bind) } } private var dataList = mutableListOf<M>() private var mOnItemClickListener: ((v: View, position: Int ) -> Unit )? = null private var mOnItemLongClickListener: ((v: View, position: Int ) -> Boolean ) = { _, _ -> false } private var onBindViewHolder: ((holder: CommonViewHolder, position: Int ) -> Unit )? = null fun onBindViewHolder (onBindViewHolder: (( holder : CommonViewHolder , position: Int ) -> Unit )) { this.onBindViewHolder = onBindViewHolder } /** * 填充数据,此操作会清除原来的数据 * * @param list 要填充的数据 * @return true:填充成功并调用刷新数据 */ fun setData (list: Collection < M >?) : Boolean { var result = false dataList.clear() if (list != null ) { result = dataList.addAll(list) } return result } /** * 根据位置获取一条数据 * * @param position View的位置 * @return 数据 */ fun getItem (position: Int ) = dataList[position] override fun onCreateViewHolder (parent: ViewGroup , viewType: Int ) : CommonViewHolder { val itemView = LayoutInflater.from(parent.context).inflate(itemLayoutId, parent, false ) val viewHolder = CommonViewHolder(itemView) val position = viewHolder.adapterPosition itemView.setOnClickListener { mOnItemClickListener?.invoke(it, position) } itemView.setOnLongClickListener { return @setOnLongClickListener mOnItemLongClickListener.invoke(it, position) } return viewHolder } override fun getItemCount () = dataList.size override fun onBindViewHolder (holder: CommonViewHolder , position: Int ) { if (onBindViewHolder != null ) { onBindViewHolder!!.invoke(holder, position) } else { bindData(holder, position) } } open fun bindData (holder: CommonViewHolder , position: Int ) {} class CommonViewHolder (itemView: View) : RecyclerView.ViewHolder(itemView), LayoutContainer { override val containerView: View = itemView } } 复制代码`

### 使用 ###

首先创建一个简单的布局 ` item_textview.xml` ：

` <?xml version= "1.0" encoding= "utf-8" ?> <TextView xmlns:android= "http://schemas.android.com/apk/res/android" android:id= "@+id/textview" android:layout_width= "match_parent" android:layout_height= "50dp" > </TextView> 复制代码`

我们可以通过继承这个适配器来创建：

` class StringAdapter : BaseRecyclerAdapter < String > (R.layout.item_textview) { override fun onBindViewHolder (holder: CommonViewHolder , position: Int ) { super.onBindViewHolder(holder, position) holder.textview.text = getItem(position) } } 复制代码`

也可以采用类似DSL的形式直接创建：

` val adapter = BaseRecyclerAdapter<String>(R.layout.item_textview) { onBindViewHolder { holder, position -> holder.textview.text = getItem(position) } } 复制代码`