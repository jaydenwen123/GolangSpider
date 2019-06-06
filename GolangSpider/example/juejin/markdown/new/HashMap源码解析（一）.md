# HashMap源码解析（一） #

## HashMap<K,V>类介绍 ##

HashMap是散列结构，这种结构是支持快速查找的。通过Key计算哈希码，通过哈希码定位到具体的Value(当然具体过程不会这么简单)。在JDK8中HashMap进行了改进，引入了红黑树。JDK8中HashMap是数组+链表+红黑树的复合数据结构。 注意：这一篇文章我们先分析HashMap中和红黑树操作无关的部分。

### HashMap结构 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2176e3b9674a2?imageView2/0/w/1280/h/960/ignore-error/1)

## HashMap中相关字段 ##

` /** * The default initial capacity - MUST be a power of two. */ static final int DEFAULT_INITIAL_CAPACITY = 1 << 4; // aka 16 /** * The maximum capacity, used if a higher value is implicitly specified * by either of the constructors with arguments. * MUST be a power of two <= 1<<30. */ static final int MAXIMUM_CAPACITY = 1 << 30; /** * The load factor used when none specified in constructor. */ static final float DEFAULT_LOAD_FACTOR = 0.75f; /** * The bin count threshold for using a tree rather than list for a * bin. Bins are converted to trees when adding an element to a * bin with at least this many nodes. The value must be greater * than 2 and should be at least 8 to mesh with assumptions in * tree removal about conversion back to plain bins upon * shrinkage. */ static final int TREEIFY_THRESHOLD = 8; /** * The bin count threshold for untreeifying a (split) bin during a * resize operation. Should be less than TREEIFY_THRESHOLD, and at * most 6 to mesh with shrinkage detection under removal. */ static final int UNTREEIFY_THRESHOLD = 6; /** * The smallest table capacity for which bins may be treeified. * (Otherwise the table is resized if too many nodes in a bin.) * Should be at least 4 * TREEIFY_THRESHOLD to avoid conflicts * between resizing and treeification thresholds. */ static final int MIN_TREEIFY_CAPACITY = 64; /* ---------------- Fields -------------- */ /** * The table, initialized on first use, and resized as * necessary. When allocated, length is always a power of two. * (We also tolerate length zero in some operations to allow * bootstrapping mechanics that are currently not needed.) */ transient Node<K,V>[] table; /** * Holds cached entrySet(). Note that AbstractMap fields are used * for keySet() and values(). */ transient Set<Map.Entry<K,V>> entrySet; /** * The number of key-value mappings contained in this map. */ transient int size; /** * The number of times this HashMap has been structurally modified * Structural modifications are those that change the number of mappings in * the HashMap or otherwise modify its internal structure (e.g., * rehash ). This field is used to make iterators on Collection-views of * the HashMap fail-fast. (See ConcurrentModificationException). */ transient int modCount; /** * The next size value at which to resize (capacity * load factor). * * @serial */ // (The javadoc description is true upon serialization. // Additionally, if the table array has not been allocated, this // field holds the initial array capacity, or zero signifying // DEFAULT_INITIAL_CAPACITY.) int threshold; /** * The load factor for the hash table. * * @serial */ final float loadFactor; 复制代码`

从上面注释我们知道：

* DEFAULT_INITIAL_CAPACITY这个字段表示默认table数组的长度，默认是16。
* TREEIFY_THRESHOLD表示桶中链表长度达到8个时，会尝试转化为红黑树，后面会介绍。
* threshold计算出来的阈值，如果HashMap中存储的元素超过这个阈值，会通过resize进行扩容。
* loadFactor加载因子，默认是0.75，我们可以在构造函数中进行动态调整。

### HashMap中链表节点Node结构 ###

` /** * Basic hash bin node, used for most entries. (See below for * TreeNode subclass, and in LinkedHashMap for its Entry subclass.) */ static class Node<K,V> implements Map.Entry<K,V> { final int hash ; final K key; V value; Node<K,V> next; Node(int hash , K key, V value, Node<K,V> next) { this.hash = hash ; this.key = key; this.value = value; this.next = next; } public final K getKey () { return key; } public final V getValue () { return value; } public final String toString () { return key + "=" + value; } public final int hashCode () { return Objects.hashCode(key) ^ Objects.hashCode(value); } public final V set Value(V newValue) { V oldValue = value; value = newValue; return oldValue; } public final boolean equals(Object o) { if (o == this) return true ; if (o instanceof Map.Entry) { Map.Entry<?,?> e = (Map.Entry<?,?>)o; if (Objects.equals(key, e.getKey()) && Objects.equals(value, e.getValue())) return true ; } return false ; } } 复制代码`

从源码中我们看到Node节点的结构还是很清晰的，hash,Key,Value和next链接。所以在HashMap中链表只是单链表，LinkedList中的Node是双向链表。

## HashMap相关构造函数 ##

` /** * Constructs an empty <tt>HashMap</tt> with the specified initial * capacity and load factor. * * @param initialCapacity the initial capacity * @param loadFactor the load factor * @throws IllegalArgumentException if the initial capacity is negative * or the load factor is nonpositive */ public HashMap(int initialCapacity, float loadFactor) { if (initialCapacity < 0) throw new IllegalArgumentException( "Illegal initial capacity: " + initialCapacity); if (initialCapacity > MAXIMUM_CAPACITY) initialCapacity = MAXIMUM_CAPACITY; if (loadFactor <= 0 || Float.isNaN(loadFactor)) throw new IllegalArgumentException( "Illegal load factor: " + loadFactor); this.loadFactor = loadFactor; this.threshold = tableSizeFor(initialCapacity); } /** * Constructs an empty <tt>HashMap</tt> with the specified initial * capacity and the default load factor (0.75). * * @param initialCapacity the initial capacity. * @throws IllegalArgumentException if the initial capacity is negative. */ public HashMap(int initialCapacity) { this(initialCapacity, DEFAULT_LOAD_FACTOR); } /** * Constructs an empty <tt>HashMap</tt> with the default initial capacity * (16) and the default load factor (0.75). */ public HashMap () { this.loadFactor = DEFAULT_LOAD_FACTOR; // all other fields defaulted } /** * Constructs a new <tt>HashMap</tt> with the same mappings as the * specified <tt>Map</tt>. The <tt>HashMap</tt> is created with * default load factor (0.75) and an initial capacity sufficient to * hold the mappings in the specified <tt>Map</tt>. * * @param m the map whose mappings are to be placed in this map * @throws NullPointerException if the specified map is null */ public HashMap(Map<? extends K, ? extends V> m) { this.loadFactor = DEFAULT_LOAD_FACTOR; putMapEntries(m, false ); } 复制代码`

* HashMap默认都会设置加载因子，我们可以设置加载因子。如果大于0.75，那么Map的利用率是提升(阈值变大，扩容延后了)。这样可能会导致Map发生碰撞的几率更高(查找元素会相对慢一些)。如果小于0.75，那么扩容相对频繁写，但是查找元素可能会快一点(Map发生碰撞的几率小了)。0.75是一个平衡，一般无需做修改。
* initialCapacity作为参数传递进来时，会通过tableSizeFor方法计算大于等于输入参数的的最小的2的指数幂 例如参数是15，那么输出16；参数是22，输出32。保证HashMap数组长度是2的幂次方。
* 如果参数是Map的话，直接调用putMapEntries插入元素。下面我们来分析这个函数。

### putMapEntries分析 ###

` /** * Implements Map.putAll and Map constructor * * @param m the map * @param evict false when initially constructing this map, else * true (relayed to method afterNodeInsertion). */ final void putMapEntries(Map<? extends K, ? extends V> m, boolean evict) { int s = m.size(); if (s > 0) { if (table == null) { // pre-size float ft = (( float )s / loadFactor) + 1.0F; int t = ((ft < ( float )MAXIMUM_CAPACITY) ? (int)ft : MAXIMUM_CAPACITY); if (t > threshold) threshold = tableSizeFor(t); } else if (s > threshold) resize(); for (Map.Entry<? extends K, ? extends V> e : m.entrySet()) { K key = e.getKey(); V value = e.getValue(); putVal( hash (key), key, value, false , evict); } } } 复制代码`

* 先根据map的大小来计算需要多大的散列表。如果table为空，那么将threshold(扩容阈值)设置为2的幂次方。
* 如果table不为空，如果map大小超过阈值，那么就先扩容，然后循环调用putValue方法插入元素即可。 我们从上面putMapEntries方法中看到，在插入节点的时候，都需要hash(key)计算出散列值，我们看下源码：

#### hash方法 ####

` /** * Computes key.hashCode() and spreads (XORs) higher bits of hash * to lower. Because the table uses power-of-two masking, sets of * hashes that vary only in bits above the current mask will * always collide. (Among known examples are sets of Float keys * holding consecutive whole numbers in small tables.) So we * apply a transform that spreads the impact of higher bits * downward. There is a tradeoff between speed, utility, and * quality of bit-spreading. Because many common sets of hashes * are already reasonably distributed (so don 't benefit from * spreading), and because we use trees to handle large sets of * collisions in bins, we just XOR some shifted bits in the * cheapest possible way to reduce systematic lossage, as well as * to incorporate impact of the highest bits that would otherwise * never be used in index calculations because of table bounds.' */ static final int hash (Object key) { int h; return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16); } 复制代码`

这个计算hash的方法设计的非常巧妙，下面我们来详细分析下这个过程：

* h=key.hashCode()。通过这一步先计算出Key键值类型自带的哈希函数，返回int类型的散列值。
* 将hashCode的高位参与运算，重新计算hash值。下面会解释为什么需要h>>>16。

#### (n - 1) & hash来计算索引 ####

我们知道HashMap要想获得最好性能，就是计算的Hash值尽可能的均匀分布在每一个桶中。理论上取模运算是一种分布很均匀的算法。但是取模运算性能消耗还是比较大的，在JDK8中，对取模运算进行了优化。通过位与运算来替代取模运算。理论公式是：x mod 2^n = x & (2^n - 1)。我们知道HashMap底层数组的长度总是2的n次方，并且取模运算为“h mod table.length”，对应上面的公式，可以得到该运算等同于“h & (table.length - 1)”。这是HashMap在速度上的优化，因为&比%具有更高的效率。

#### 扰动函数 ####

上面我们在介绍hash计算的时候，看到h>>>16这样的操作。为什么key的hashCode还需要使用高16位进行异或操作呢？

* 我们先假设没有h>>>16这个操作。看看索引位置如何计算的 我们假设HashMap桶的长度是默认值16.现在的索引计算如下：

` 10100101 11000100 00100101 & 00000000 00000000 00001111 ---------------------------------- 00000000 00000000 00000101 //高位全部归零，只保留末四位 复制代码`

但这时候问题就来了，这样就算我的散列值分布再松散，要是只取最后几位的话，碰撞也会很严重。更要命的是如果散列本身做得不好，分布上成等差数列的漏洞，恰好使最后几个低位呈现规律性重复，那么碰撞就会更加严重。

* 扰动函数的作用 上面提到过，只是使用最后几位的话，碰撞会很严重，严重降低Map性能。如果我们将高16位也加入运算，就可以较好的解决问题。如图： ![](https://user-gold-cdn.xitu.io/2019/6/4/16b2270c9ef2cab1?imageView2/0/w/1280/h/960/ignore-error/1) 右位移16位，正好是32bit的一半，自己的高半区和低半区做异或，就是为了混合原始哈希码的高位和低位，以此来加大低位的随机性(减少碰撞)。而且混合后的低位掺杂了高位的部分特征，这样高位的信息也被变相保留下来。

### putVal方法 ###

` /** * Implements Map.put and related methods * * @param hash hash for key * @param key the key * @param value the value to put * @param onlyIfAbsent if true , don 't change existing value * @param evict if false, the table is in creation mode. * @return previous value, or null if none' */ final V putVal(int hash , K key, V value, boolean onlyIfAbsent, boolean evict) { Node<K,V>[] tab; Node<K,V> p; int n, i; if ((tab = table) == null || (n = tab.length) == 0) n = (tab = resize()).length; if ((p = tab[i = (n - 1) & hash ]) == null) tab[i] = newNode( hash , key, value, null); else { Node<K,V> e; K k; if (p.hash == hash && ((k = p.key) == key || (key != null && key.equals(k)))) e = p; else if (p instanceof TreeNode) e = ((TreeNode<K,V>)p).putTreeVal(this, tab, hash , key, value); else { for (int binCount = 0; ; ++binCount) { if ((e = p.next) == null) { p.next = newNode( hash , key, value, null); if (binCount >= TREEIFY_THRESHOLD - 1) // -1 for 1st treeifyB in (tab, hash ); break ; } if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) break ; p = e; } } if (e != null) { // existing mapping for key V oldValue = e.value; if (!onlyIfAbsent || oldValue == null) e.value = value; afterNodeAccess(e); return oldValue; } } ++modCount; if (++size > threshold) resize(); afterNodeInsertion(evict); return null; } 复制代码`

* 首先判断table是否为null或者size==0，如果还未初始化或者集合为空，那么先resize进行扩容。
* 使用p = tab[i = (n - 1) & hash]方法定位到当前的桶在table中的索引。如果当前桶中没有Node节点，那么创建新的Node节点放入桶中即可。否则就说明当前桶中已有元素，需要遍历链表。
* else分支首先判断第一个Node节点是否是符合条件的，如果是，整个查找过程就结束了。如果不是，判断第一个节点是否是树节点(这个桶是否已经树化)。如果已经转化成红黑树，就调用红黑树的插入操作。否则就是普通的链表遍历操作。如果整个遍历下来，都没有找到符合条件的的Node节点的话，就构造一个Node节点，放入链表尾部即可，当然，如果链表的长度超过8，会调用treeifyBin方法将这个链表转化为红黑树。否则将找到的节点赋值给e即可。
* 如果e不为空的话，就说明Map中已有符合条件的节点，新的Value值可以根据参数决定是否覆盖旧值。
* 最后size++，判断节点数量是否超过threshold阈值，超过的话需要扩容。
* 从源码看出，Map中判断一个元素是否相等是否的是==或者equal()方法。所以我们在自定义类中，需要好好设计equal()方法和hashCode方法，因为这关乎到Map中元素的查找与比较。

### resize()扩容方法 ###

` /** * Initializes or doubles table size. If null, allocates in * accord with initial capacity target held in field threshold. * Otherwise, because we are using power-of-two expansion, the * elements from each bin must either stay at same index, or move * with a power of two offset in the new table. * * @ return the table */ final Node<K,V>[] resize () { Node<K,V>[] oldTab = table; int oldCap = (oldTab == null) ? 0 : oldTab.length; int oldThr = threshold; int newCap, newThr = 0; if (oldCap > 0) { if (oldCap >= MAXIMUM_CAPACITY) { threshold = Integer.MAX_VALUE; return oldTab; } else if ((newCap = oldCap << 1) < MAXIMUM_CAPACITY && oldCap >= DEFAULT_INITIAL_CAPACITY) newThr = oldThr << 1; // double threshold } else if (oldThr > 0) // initial capacity was placed in threshold newCap = oldThr; else { // zero initial threshold signifies using defaults newCap = DEFAULT_INITIAL_CAPACITY; newThr = (int)(DEFAULT_LOAD_FACTOR * DEFAULT_INITIAL_CAPACITY); } if (newThr == 0) { float ft = ( float )newCap * loadFactor; newThr = (newCap < MAXIMUM_CAPACITY && ft < ( float )MAXIMUM_CAPACITY ? (int)ft : Integer.MAX_VALUE); } threshold = newThr; @SuppressWarnings({ "rawtypes" , "unchecked" }) Node<K,V>[] newTab = (Node<K,V>[])new Node[newCap]; table = newTab; if (oldTab != null) { for (int j = 0; j < oldCap; ++j) { Node<K,V> e; if ((e = oldTab[j]) != null) { oldTab[j] = null; if (e.next == null) newTab[e.hash & (newCap - 1)] = e; else if (e instanceof TreeNode) ((TreeNode<K,V>)e).split(this, newTab, j, oldCap); else { // preserve order Node<K,V> loHead = null, loTail = null; Node<K,V> hiHead = null, hiTail = null; Node<K,V> next; do { next = e.next; if ((e.hash & oldCap) == 0) { if (loTail == null) loHead = e; else loTail.next = e; loTail = e; } else { if (hiTail == null) hiHead = e; else hiTail.next = e; hiTail = e; } } while ((e = next) != null); if (loTail != null) { loTail.next = null; newTab[j] = loHead; } if (hiTail != null) { hiTail.next = null; newTab[j + oldCap] = hiHead; } } } } } return newTab; } 复制代码`

* 首先根据扩容前的容量oldCap，如果oldCap容量已经到最大值了，那么不进行扩容，只是将阈值设置为Integer.MAX_VALUE。否则就是扩容为原来的2倍。
* 如果oldCap==0，但是oldThr不为空的时候（因为构造HashMap初始容量被放入阈值），会将容量设置为当前的阈值。
* Map的容量和阈值都是0时，是一个空表，Map容量设置为DEFAULT_INITIAL_CAPACITY大小(16)。并计算出阈值。
* 如果新的阈值为0，就根据新的容量和加载因子计算出新的阈值。
* 开始遍历老的Map集合，将里面的Node节点重新定位到新的Map集合中。如果桶中只有一个元素，通过newTab[e.hash & (newCap - 1)] = e;计算出该Node在新Map中的索引即可。
* 否则判断该桶是否已经树化，如果树化，调用树节点的方法进行hash分布。否则就需要将链表数据一个个遍历，重新定位。此处：HashMap的方法设计的非常精妙。通过定义loHead、loTail、hiHead、hiTail来讲一个链表拆分成两个独立的链表。 注意：如果e的hash值与老表的容量进行与运算为0,则扩容后的索引位置跟老表的索引位置一样。所以loHead-->loTail组成的链表在新Map中的索引位置和老Map中是一样的。如果e的hash值与老表的容量进行与运算为1,则扩容后的索引位置为:老表的索引位置＋oldCap。
* 最后将loHead、loTail、hiHead、hiTail组成的两条链表重新定位到新的Map中即可。

### 扩容代码(e.hash & oldCap)是否为0来定位的问题 ###

扩容代码中，使用e节点的hash值跟oldCap进行位与运算，以此决定将节点分布到原索引位置或者原索引+oldCap位置上，为什么可以这样计算，我们来看例子：
假设老表的容量为16，即oldCap=16，则新表容量为16*2=32，假设节点1的hash值为0000 0000 0000 0000 0000 1111 0000 1010，节点2的hash值为0000 0000 0000 0000 0000 1111 0001 1010，则节点1和节点2在老表的索引位置计算如下图计算1，由于老表的长度限制，节点1和节点2的索引位置只取决于节点hash值的最后4位。再看计算2，计算2为新表的索引计算，可以知道如果两个节点在老表的索引位置相同，则新表的索引位置只取决于节点hash值倒数第5位的值，而此位置的值刚好为老表的容量值16，此时节点在新表的索引位置只有两种情况：原索引位置和原索引+oldCap位置（在此例中即为10和10+16=26）。由于结果只取决于节点hash值的倒数第5位，而此位置的值刚好为老表的容量值16，因此此时新表的索引位置的计算可以替换为计算3，直接使用节点的hash值与老表的容量16进行位于运算，如果结果为0则该节点在新表的索引位置为原索引位置，否则该节点在新表的索引位置为原索引+oldCap位置。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b276c286ebff3c?imageView2/0/w/1280/h/960/ignore-error/1)

### treeifyBin方法 ###

` /** * Replaces all linked nodes in bin at index for given hash unless * table is too small, in which case resizes instead. */ final void treeifyB in (Node<K,V>[] tab, int hash ) { int n, index; Node<K,V> e; if (tab == null || (n = tab.length) < MIN_TREEIFY_CAPACITY) resize(); else if ((e = tab[index = (n - 1) & hash ]) != null) { TreeNode<K,V> hd = null, tl = null; do { TreeNode<K,V> p = replacementTreeNode(e, null); if (tl == null) hd = p; else { p.prev = tl; tl.next = p; } tl = p; } while ((e = e.next) != null); if ((tab[index] = hd) != null) hd.treeify(tab); } } 复制代码`

这个代码是将链表转化成红黑树的。我们来看主要的逻辑：

* 如果Map的容量小于MIN_TREEIFY_CAPACITY(64)。是不会讲某一个桶中的链表转化为红黑树的，会对Map进行扩容。这样每一个桶中的链表的长度会减少。
* 如果Map的容量符合要求了，那就将链表转化成红黑树。

### get方法 ###

` /** * Returns the value to which the specified key is mapped, * or {@code null} if this map contains no mapping for the key. * * <p>More formally, if this map contains a mapping from a key * {@code k} to a value {@code v} such that {@code (key==null ? k==null : * key.equals(k))}, then this method returns {@code v}; otherwise * it returns {@code null}. (There can be at most one such mapping.) * * <p>A return value of {@code null} does not <i>necessarily</i> * indicate that the map contains no mapping for the key; it 's also * possible that the map explicitly maps the key to {@code null}. * The {@link #containsKey containsKey} operation may be used to * distinguish these two cases. * * @see #put(Object, Object) */ public V get(Object key) { Node<K,V> e; return (e = getNode(hash(key), key)) == null ? null : e.value; } /** * Implements Map.get and related methods * * @param hash hash for key * @param key the key * @return the node, or null if none */ final Node<K,V> getNode(int hash, Object key) { Node<K,V>[] tab; Node<K,V> first, e; int n; K k; if ((tab = table) != null && (n = tab.length) > 0 && (first = tab[(n - 1) & hash]) != null) { if (first.hash == hash && // always check first node ((k = first.key) == key || (key != null && key.equals(k)))) return first; if ((e = first.next) != null) { if (first instanceof TreeNode) return ((TreeNode<K,V>)first).getTreeNode(hash, key); do { if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) return e; } while ((e = e.next) != null); } } return null; } 复制代码`

HashMap是高效的查找数据结构，所以get方法我们有必要好好分析下。

* 首先计算出key的哈希值，然后使用first = tab[(n - 1) & hash])定位到索引位置。
* 首先判断第一个Node节点是否是符合要求的，符合要求就直接返回即可。否则判断当前桶的结构是树结构还是链表结构，分别使用相关的方法寻找节点。没有找到就返回null。

### remove方法 ###

` /** * Removes the mapping for the specified key from this map if present. * * @param key key whose mapping is to be removed from the map * @ return the previous value associated with <tt>key</tt>, or * <tt>null</tt> if there was no mapping for <tt>key</tt>. * (A <tt>null</tt> return can also indicate that the map * previously associated <tt>null</tt> with <tt>key</tt>.) */ public V remove(Object key) { Node<K,V> e; return (e = removeNode( hash (key), key, null, false , true )) == null ? null : e.value; } /** * Implements Map.remove and related methods * * @param hash hash for key * @param key the key * @param value the value to match if matchValue, else ignored * @param matchValue if true only remove if value is equal * @param movable if false do not move other nodes while removing * @ return the node, or null if none */ final Node<K,V> removeNode(int hash , Object key, Object value, boolean matchValue, boolean movable) { Node<K,V>[] tab; Node<K,V> p; int n, index; if ((tab = table) != null && (n = tab.length) > 0 && (p = tab[index = (n - 1) & hash ]) != null) { Node<K,V> node = null, e; K k; V v; if (p.hash == hash && ((k = p.key) == key || (key != null && key.equals(k)))) node = p; else if ((e = p.next) != null) { if (p instanceof TreeNode) node = ((TreeNode<K,V>)p).getTreeNode( hash , key); else { do { if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) { node = e; break ; } p = e; } while ((e = e.next) != null); } } if (node != null && (!matchValue || (v = node.value) == value || (value != null && value.equals(v)))) { if (node instanceof TreeNode) ((TreeNode<K,V>)node).removeTreeNode(this, tab, movable); else if (node == p) tab[index] = node.next; else p.next = node.next; ++modCount; --size; afterNodeRemoval(node); return node; } } return null; } 复制代码`

* 先根据key找到相应的节点，然后判断需要删除的节点是树结构还是链表结构。如果是树结构调用removeTreeNode方法即可。链表的话只需要重新设置next节点即可。

## 总结 ##

* JDK8中HashMap的结构是数组+链表+红黑树的复合结构。
* HashMap默认大小是16，加载因子0.75，可以根据项目需要自定义加载因子。
* HashMap的扩容操作是比较消耗时间的，如果可以的话最好预估HashMap初始化的容量，以此来避免频繁的扩容操作。
* HashMap中链表转化为红黑树要满足两个条件才行，第一：链表的长度已经达到8个了。第二：HashMap容量大于64即可。
* HashMap中索引的定位和元素的查找，非常依赖key的hashCode和equal方法，我们在自定义的类型的时候需要好好考虑如何比较两个对象。