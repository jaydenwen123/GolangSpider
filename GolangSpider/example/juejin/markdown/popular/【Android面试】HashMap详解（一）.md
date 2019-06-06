# 【Android面试】HashMap详解（一） #

### 前言 ###

` HashMap` 是面试中比较常见的问题，这一篇，我们将通过阅读源码，了解其设计原理以及以下问题

* HashMap的实现原理
* 初始容量为什么是2的倍数
* 如何 ` resize`
* 是否线程安全

### 常用参数 ###

` //最大容量 2的30次方 static final int MAXIMUM_CAPACITY = 1 << 30 ; //初始容量为16 扩容时才会触发 static final int DEFAULT_INITIAL_CAPACITY = 1 << 4 ; // aka 16 //默认的加载因子 static final float DEFAULT_LOAD_FACTOR = 0.75f ; //哈希表，存放链表。 长度是2的N次方，或者初始化时为0. transient Node<K,V>[] table; //加载因子，用于计算哈希表元素数量的阈值。 threshold = 哈希表.length * loadFactor; final float loadFactor; //哈希表内元素数量的阈值，当哈希表内元素数量超过阈值时，会发生扩容resize()。 int threshold; 复制代码`

其中， ` table` 称之为哈希表，用于存放 链表 ` Node`

### Node 链表 ###

` static class Node < K , V > implements Map. Entry < K , V > { final int hash; final K key; V value; Node<K,V> next; Node( int hash, K key, V value, Node<K,V> next) { this.hash = hash; this.key = key; this.value = value; this.next = next; } public final K getKey () { return key; } public final V getValue () { return value; } public final String toString () { return key + "=" + value; } public final int hashCode () { return Objects.hashCode(key) ^ Objects.hashCode(value); } public final V setValue (V newValue) { V oldValue = value; value = newValue; return oldValue; } public final boolean equals (Object o) { if (o == this ) return true ; if (o instanceof Map.Entry) { Map.Entry<?,?> e = (Map.Entry<?,?>)o; if (Objects.equals(key, e.getKey()) && Objects.equals(value, e.getValue())) return true ; } return false ; } } 复制代码`

可以看到， ` Node` 的数据结构是 ` Node` 嵌套 ` Node` ，故称之为链表； 整个链表中的某个 ` Node` 称之为节点。

### 构造方法 ###

` public HashMap () { this.loadFactor = DEFAULT_LOAD_FACTOR; // all other fields defaulted } 复制代码` ` public HashMap ( int initialCapacity) { //指定初始化容量的构造函数 this (initialCapacity, DEFAULT_LOAD_FACTOR); } 复制代码` ` public HashMap ( int initialCapacity, float loadFactor) { if (initialCapacity < 0 ) throw new IllegalArgumentException( "Illegal initial capacity: " + initialCapacity); if (initialCapacity > MAXIMUM_CAPACITY) initialCapacity = MAXIMUM_CAPACITY; if (loadFactor <= 0 || Float.isNaN(loadFactor)) throw new IllegalArgumentException( "Illegal load factor: " + loadFactor); this.loadFactor = loadFactor; //设置阈值 初始化容量的 2的n次方 this.threshold = tableSizeFor(initialCapacity); } 复制代码` ` public HashMap (Map<? extends K, ? extends V> m) { this.loadFactor = DEFAULT_LOAD_FACTOR; //将m中元素加入新的哈希表中，同增、改逻辑 putMapEntries(m, false ); } 复制代码`

### 扩容 ` resize` ###

` final Node<K,V>[] resize() { Node<K,V>[] oldTab = table; int oldCap = (oldTab == null ) ? 0 : oldTab.length; int oldThr = threshold; int newCap, newThr = 0 ; if (oldCap > 0 ) { //已经达到上限，不在扩容 if (oldCap >= MAXIMUM_CAPACITY) { threshold = Integer.MAX_VALUE; return oldTab; } else if ((newCap = oldCap << 1 ) < MAXIMUM_CAPACITY && oldCap >= DEFAULT_INITIAL_CAPACITY) //如果旧的容量大于等于默认初始容量16 //新的容量将变为旧的2倍 newThr = oldThr << 1 ; // double threshold } else if (oldThr > 0 ) // initial capacity was placed in threshold newCap = oldThr; else { // zero initial threshold signifies using defaults //旧哈希表为0，则新表为默认值 newCap = DEFAULT_INITIAL_CAPACITY; newThr = ( int )(DEFAULT_LOAD_FACTOR * DEFAULT_INITIAL_CAPACITY); } if (newThr == 0 ) { //新的阈值是0 重新计算 float ft = ( float )newCap * loadFactor; newThr = (newCap < MAXIMUM_CAPACITY && ft < ( float )MAXIMUM_CAPACITY ? ( int )ft : Integer.MAX_VALUE); } threshold = newThr; @SuppressWarnings ({ "rawtypes" , "unchecked" }) Node<K,V>[] newTab = (Node<K,V>[]) new Node[newCap]; table = newTab; if (oldTab != null ) { //遍历旧的哈希表 for ( int j = 0 ; j < oldCap; ++j) { Node<K,V> e; if ((e = oldTab[j]) != null ) { oldTab[j] = null ; if (e.next == null ) //就一个元素，将其放入新哈希表的e.hash & (newCap - 1)位置 newTab[e.hash & (newCap - 1 )] = e; else if (e instanceof TreeNode) ((TreeNode<K,V>)e).split( this , newTab, j, oldCap); else { // preserve order //因为扩容是容量翻倍，所以原哈希表上的每个链表，现在可能存放在原来的下标，即low位， // 或者扩容后的下标，即high位。 high位= low位+原哈希桶容量 Node<K,V> loHead = null , loTail = null ; Node<K,V> hiHead = null , hiTail = null ; Node<K,V> next; do { next = e.next; if ((e.hash & oldCap) == 0 ) { if (loTail == null ) loHead = e; else loTail.next = e; loTail = e; } else { if (hiTail == null ) hiHead = e; else hiTail.next = e; hiTail = e; } } while ((e = next) != null ); if (loTail != null ) { loTail.next = null ; newTab[j] = loHead; } if (hiTail != null ) { hiTail.next = null ; newTab[j + oldCap] = hiHead; } } } } } return newTab; } 复制代码`

简单来说， **如果哈希表为空，则分配为初始容量，否则扩容为原来的2倍，原哈希表的元素可能会在新哈希表的相同位置，也可能会在 ` index+oldCap` 位置**

### 增、改 ` put(K key, V value)` ###

` public V put (K key, V value) { return putVal(hash(key), key, value, false , true ); } 复制代码`

* 1.计算哈希值

` static final int hash (Object key) { int h; return (key == null ) ? 0 : (h = key.hashCode()) ^ (h >>> 16 ); } 复制代码`

哈希值的算法，不是本文的重点，在此先略过

* 2.插入或覆盖值 如果 ` table` 的第 ` (n - 1) & hash` 个位置为空，则在该位置插入新的链表；如不为空，先判断键值是否相等，如相等则直接覆盖值； **如为红黑树，则插入** ；否则，在链表的尾部追加新的节点。

` final V putVal ( int hash, K key, V value, boolean onlyIfAbsent, boolean evict) { Node<K,V>[] tab; Node<K,V> p; int n, i; //如果当前table为null 代表是初始化，直接扩容 if ((tab = table) == null || (n = tab.length) == 0 ) n = (tab = resize()).length; //如果第i（(n - 1) & hash）个为null 直接放入新的节点 if ((p = tab[i = (n - 1 ) & hash]) == null ) tab[i] = newNode(hash, key, value, null ); else { Node<K,V> e; K k; //如果哈希值相等，key也相等，则返回的是需要覆盖的节点 if (p.hash == hash && ((k = p.key) == key || (key != null && key.equals(k)))) e = p; else if (p instanceof TreeNode) //红黑树 e = ((TreeNode<K,V>)p).putTreeVal( this , tab, hash, key, value); else { for ( int binCount = 0 ; ; ++binCount) { if ((e = p.next) == null ) { //遍历到尾部，追加新节点到尾部 p.next = newNode(hash, key, value, null ); //如果追加节点后，链表数量大于等于8，则转化为红黑树 if (binCount >= TREEIFY_THRESHOLD - 1 ) // -1 for 1st treeifyBin(tab, hash); break ; } //如果找到了要覆盖的节点 if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) break ; p = e; } } //有需要覆盖的节点 if (e != null ) { // existing mapping for key V oldValue = e.value; if (!onlyIfAbsent || oldValue == null ) e.value = value; //空的 afterNodeAccess(e); //返回旧的值 return oldValue; } } ++modCount; //更新size，并判断是否需要扩容。此处的size即为map的size if (++size > threshold) resize(); //空的 afterNodeInsertion(evict); return null ; } 复制代码`

流程图如下

![image.png](https://user-gold-cdn.xitu.io/2019/6/5/16b26b093dece152?imageView2/0/w/1280/h/960/ignore-error/1)

### 查 ` get(Object key)` ###

` public V get (Object key) { Node<K,V> e; return (e = getNode(hash(key), key)) == null ? null : e.value; } 复制代码`

* 

1.计算哈希值

* 

2.判断 ` table` 的第 ` (n - 1) & hash` 位置值的情况 如 ` key` 相等，则返回该节点； **如为红黑树，则返回该节点** ；否则，从该节点往后循环查找，直到找到相等的 ` key` 或者空节点。

` final Node<K,V> getNode ( int hash, Object key) { Node<K,V>[] tab; Node<K,V> first, e; int n; K k; if ((tab = table) != null && (n = tab.length) > 0 && (first = tab[(n - 1 ) & hash]) != null ) { if (first.hash == hash && // always check first node ((k = first.key) == key || (key != null && key.equals(k)))) //key相等 则返回该节点 return first; if ((e = first.next) != null ) { //红黑树 if (first instanceof TreeNode) return ((TreeNode<K,V>)first).getTreeNode(hash, key); //从该节点往后循环查找，直到找到相等的key或者空节点 do { if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) return e; } while ((e = e.next) != null ); } } return null ; } 复制代码`

### 删除 ` remove(Object key)` ###

* 1.计算哈希值 ` final Node<K,V> removeNode(int hash, Object key, Object value, boolean matchValue, boolean movable)`
* 2.判断 ` table` 的第 ` (n - 1) & hash` 位置节点的情况

` matchValue` 是 ` true` ，则必须 ` key` 、 ` value` 都相等才删除 ` movable` 参数是 ` false` ，在删除节点时，不移动其他节点

` final Node<K,V> removeNode ( int hash, Object key, Object value, boolean matchValue, boolean movable) { Node<K,V>[] tab; Node<K,V> p; int n, index; if ((tab = table) != null && (n = tab.length) > 0 && (p = tab[index = (n - 1 ) & hash]) != null ) { Node<K,V> node = null , e; K k; V v; if (p.hash == hash && ((k = p.key) == key || (key != null && key.equals(k)))) //key相等 则返回该节点 node = p; else if ((e = p.next) != null ) { if (p instanceof TreeNode) //红黑树 node = ((TreeNode<K,V>)p).getTreeNode(hash, key); else { //从该节点往后循环查找，直到找到相等的key或者空节点 do { if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) { node = e; break ; } p = e; } while ((e = e.next) != null ); } } //有待删的节点且 matchValue为false，或者值也相等 if (node != null && (!matchValue || (v = node.value) == value || (value != null && value.equals(v)))) { if (node instanceof TreeNode) ((TreeNode<K,V>)node).removeTreeNode( this , tab, movable); else if (node == p) //链表头是待删除链表 tab[index] = node.next; else //待删除节点在表中间 p.next = node.next; ++modCount; //修正size --size; afterNodeRemoval(node); return node; } } return null ; } 复制代码`

如 ` key` 相等，则返回该节点； **如为红黑树，则返回该链表** ；否则，从该节点往后循环查找，直到找到节点，然后删除它。

### 总结 ###

` HashMap` 内部是一个 ` Node` 数组即 ` Node<K,V>[] table` ，称之为哈希表，其中存放的 ` Node` 即链表，链表中的每个节点则是我们放入 ` HashMap` 中的元素。

* 是否线程安全

线程不安全，存取过程没有任何锁。

* 为何要扩容

因其数据结构是数组，所以会涉及到扩容。

* 如何扩容

当 ` HashMap` 的容量大于 ` threshold` ( ` length*扩容因子` )值时，就会触发扩容；如果链表为空，则分配为初始容量，否则扩容为原来的2倍，原链表的节点可能会在新链表的相同位置，也可能会在 ` index+oldCap` 位置， **扩容前后，哈希表的长度一定会是2的倍数。**

* 为什么哈希表( ` table` )容量是2的倍数

` HashMap` 存取时，计算 ` index` 即 ` (length- 1) & hash` ，使用 ` &` 运算符（相比 ` %` 效率更高），如果 ` length` 为2的倍数，可以最大程度的确保 ` index` 的均分，简单来说： **如果是2的倍数，就可以用位运算替代取余操作，更加高效**

* 为什么需要需要 ` &` 位运算（取余 ` %` ）

` hashCode()` 是 ` int` 类型，取值范围是40多亿，只要哈希函数映射的比较均匀松散，碰撞几率是很小的。 但是由于 ` HashMap` 的哈希表的长度远比 ` hash` 取值范围小，默认是16，所以当对hash值以表的的长度 ` length` 取余，以找到存放该 ` key` 的下标时，由于取余是通过与操作( ` &` )完成的，会忽略hash值的高位。因此只有 ` hashCode()` 的低位参加运算，发生不同的 ` hash` 值，但是得到的 ` index` 相同的情况的几率会大大增加，这种情况称之为 **hash碰撞** 。

* 怎么解决 **hash碰撞**

扰动函数（ ` hash(Object key)` ）就是为了解决hash碰撞的。它会综合hash值高位和低位的特征，并存放在低位，因此在与运算时，相当于高低位一起参与了运算，以减少hash碰撞的概率。

` static final int hash (Object key) { int h; return (key == null ) ? 0 : (h = key.hashCode()) ^ (h >>> 16 ); } 复制代码`

* ` HashMap` 优化

1.使用合适的初始容量，减少 ` resize` 带来的损失。

2.使用合适的 ` key` ，比如 ` String` 、 ` Integer` ，来减少 ` hash碰撞` ，这样的话，存、取效率最高。

## 觉得有用，就请点个赞，谢谢 ##

参考资料： [面试必备：HashMap源码解析（JDK8）]( https://juejin.im/post/599652796fb9a0249975a318#heading-4 )