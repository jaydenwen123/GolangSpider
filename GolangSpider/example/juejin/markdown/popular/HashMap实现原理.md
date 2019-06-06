# HashMap实现原理 #

**[个人博客项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FVip-Augus%2FVip-Augus.github.io )**

希望各位帮忙点个star，给我加个小星星✨

**HashMap** 在编程开发中经常使用到，用来存储key-value，但是一直没深入学习它的实现原理，这次学习了记录一下。

## **HashMap类** ##

` public class HashMap < K , V > extends AbstractMap < K , V > implements Map < K , V >, Cloneable , Serializable 复制代码`

HashMap继承自 **AbstractMap** ，AbstractMap是Map接口的骨干实现，AbstractMap中实现了Map中最重要最常用和方法，这样HashMap继承AbstractMap就不需要实现Map的所有方法，让HashMap减少了大量的工作。

### **HashMap源码定义的变量和常量:** ###

` //默认初始化容量 static final int DEFAULT_INITIAL_CAPACITY = 16 ; //最大值容量 static final int MAXIMUM_CAPACITY = 1 << 30 ; //默认加载因子 static final float DEFAULT_LOAD_FACTOR = 0.75f ; //Entry类型的数组,HashMap用来维护内部的数据结构,长度由容量决定 transient Entry<K,V>[] table; //HashMap的大小 transient int size; //HashMap的极限容量,容量乘以加载因子,是扩容的临界点 int threshold; //手动设置的加载因子 final float loadFactor; //修改次数 transient int modCount; 复制代码`

### **HashMap的构造函数** ###

` public HashMap () ：构造一个具有默认初始容量 ( 16 ) 和默认加载因子 ( 0.75 ) 的空 HashMap public HashMap ( int initialCapacity) ：构造一个带指定初始容量和默认加载因子 ( 0.75 ) 的空 HashMap public HashMap ( int initialCapacity, float loadFactor) ：构造一个带指定初始容量和加载因子的空 HashMap public HashMap (Map< ? extends K, ? extends V> m) ：构造一个映射关系与指定 Map 相同的新 HashMap 复制代码`

两个很重要的参数： **initialCapacity（初始容量）** 、 **loadFactor（加载因子）** ，JDK中的是这样解释的：   HashMap 的实例有两个参数影响其性能： **初始容量** 和 **加载因子** 。   容量 ：是哈希表中桶的数量，初始容量只是哈希表在创建时的容量，实际上就是Entry< K,V>[] table的容量   加载因子 ：是哈希表在其容量自动增加之前可以达到多满的一种尺度。它衡量的是一个散列表的空间的使用程度，负载因子越大表示散列表的装填程度越高，反之愈小。对于使用链表法的散列表来说，查找一个元素的平均时间是O(1+a)，因此如果负载因子越大，对空间的利用更充分，然而后果是查找效率的降低；如果负载因子太小，那么散列表的数据将过于稀疏，对空间造成严重浪费。   当哈希表中的条目数超出了加载因子与当前容量的乘积时，则要对该哈希表进行 **rehash** 操作（即重建内部数据结构），从而哈希表将具有大约两倍的桶数。

## **HashMap的数据结构** ##

![HashMap数据结构](https://user-gold-cdn.xitu.io/2019/5/25/16aedeb26ae639fe?imageView2/0/w/1280/h/960/ignore-error/1)

底层实现是用 **数组** 和 **模拟指针** 实现的链表散列

**HashMap构造函数** 的源码：

` public HashMap ( int initialCapacity, float loadFactor) { //如果传入初始容量小于0，报错 if (initialCapacity < 0 ) throw new IllegalArgumentException( "Illegal initial capacity: " + initialCapacity); //限定最大容量 if (initialCapacity > MAXIMUM_CAPACITY) initialCapacity = MAXIMUM_CAPACITY; //加载因子要传大于零的数字，不然会报错 if (loadFactor <= 0 || Float.isNaN(loadFactor)) throw new IllegalArgumentException( "Illegal load factor: " + loadFactor); // 找到一个大于初始化容量最近的2的n次方 int capacity = 1 ; while (capacity < initialCapacity) capacity <<= 1 ; this.loadFactor = loadFactor; //计算极限容量 threshold = ( int )Math.min(capacity * loadFactor, MAXIMUM_CAPACITY + 1 ); //创建table数组 table = new Entry[capacity]; useAltHashing = sun.misc.VM.isBooted() && (capacity >= Holder.ALTERNATIVE_HASHING_THRESHOLD); init(); } 复制代码`

初始化函数主要做了三件事情：

* 对传入的初始化容量和加载因子做了校验
* 设置HashMap的极限容量
* 计算出一个离初始化容量最近的2的n次方，用来创建table数组

Table数组的成员 **Entry<K, V>** :

` static class Entry < K , V > implements Map. Entry < K , V > { final K key; //键 V value; //值 Entry<K,V> next; //next引用（该引用指向当前table的位置的链表） int hash; //用来确定每一个Entry链表在table中的位置 /** * Creates new entry. */ Entry( int h, K k, V v, Entry<K,V> n) { value = v; next = n; key = k; hash = h; } ··· 复制代码`

Entry是HashMap的一个 **内部类** ，它也是维护着一个 **key-value** 映射关系，除了key和value，还有 **next引用** （该引用指向当前table位置的链表）， **hash值** （用来确定每一个Entry链表在table中位置）

## **HashMap的存储实现Put(K, V)** ##

` public V put (K key, V value) { //key为的null的情况下 if (key == null ) return putForNullKey(value); //计算key值对应的hash值 int hash = hash(key); //根据hash值找到数组对应的下标i int i = indexFor(hash, table.length); //如果在数组中该链表中存在该key，且hash值也相等，用新值替换旧值 for (Entry<K,V> e = table[i]; e != null ; e = e.next) { Object k; if (e.hash == hash && ((k = e.key) == key || key.equals(k))) { V oldValue = e.value; e.value = value; e.recordAccess( this ); //返回结束 return oldValue; } } //修改次数加一 modCount++; //真正存入的方法 addEntry(hash, key, value, i); return null ; } 复制代码`

存储有两种情况（ **key** 为null和非null） **1.** key为null的情况： 调用putForNullKey(value):

` private V putForNullKey (V value) { //遍历链表，查看是否存在key为null for (Entry<K,V> e = table[ 0 ]; e != null ; e = e.next) { if (e.key == null ) { V oldValue = e.value; e.value = value; e.recordAccess( this ); return oldValue; } } modCount++; //如果找不到，就把null键插入到链表中 addEntry( 0 , null , value, 0 ); return null ; } 复制代码`

先找是否有null的键，如果有的话替换旧值，没有的话才进行新增。

**2.** key不为null： 如果存在key，则在原来链表中更换旧值；如果不存在key，将key-value添加进table数组中。 **addEntry(0, null, value, 0):**

` void addEntry ( int hash, K key, V value, int bucketIndex) { //如果初始容量超过极限容量，需要扩容 if ((size >= threshold) && ( null != table[bucketIndex])) { resize( 2 * table.length); //这一步就是对null的处理，如果key为null，hash值为0，也就是会插入到哈希表的表头table[0]的位置 hash = ( null != key) ? hash(key) : 0 ; bucketIndex = indexFor(hash, table.length); } createEntry(hash, key, value, bucketIndex); } 复制代码`

### **hash值的计算** ###

` final int hash (Object k) { int h = 0 ; if (useAltHashing) { if (k instanceof String) { return sun.misc.Hashing.stringHash32((String) k); } h = hashSeed; } h ^= k.hashCode(); // This function ensures that hashCodes that differ only by // constant multiples at each bit position have a bounded // number of collisions (approximately 8 at default load factor). h ^= (h >>> 20 ) ^ (h >>> 12 ); return h ^ (h >>> 7 ) ^ (h >>> 4 ); } 复制代码`

hash算法的作用是为了让hashMap中的元素尽量 **分散** ，尽量做到每一个位置上面只有一个元素，当计算出key对应的hash值，马上就能得到该位置上的value就是我们所希望得到的。

其中计算hash值是用异或运算，有可能想到为什么不将hashcode与 **数组长度** 做取模运算，取模预算的话得到的位置应该更加均匀。看到的文章说：因为模运算的消耗比较大（可能是计算机最终执行的还是二进制，所以直接用异或会比高级语言更快-- **大雾，不确定** ），所以用了异或，消耗更小。

还有一个根据hash值得到table数组下标i的运算 **indexFor(hash)** :

` static int indexFor ( int h, int length) { return h & (length- 1 ); } 复制代码`

下标算法中，先将数组长度-1，然后跟hash值进行**&（与）运算**，数组长度上面已经计算过了，是2的n次方，先给大家看个图：

![哈希值运算](https://user-gold-cdn.xitu.io/2019/5/25/16aedeb26af1bbf6?imageView2/0/w/1280/h/960/ignore-error/1)

为什么数组的长度要设定在 **2的n次方** 呢： ​
假设一个数组长度为16，另一个数组长度为15，length-1之后换成二进制，刚好两个key计算出来的hash值为8、9，进行与运算，得到的下标值如上图，数组长度为15的下标都为8，这样就会造成不同的key，存放在同一个链表中，造成 **碰撞** 的几率增大，如果查询的时候还需要 **循环** 这个链表，造成查询慢。

还有一个问题，就是末位为零，进行与运算，得到的结果是 **末位永远没有1** ，这样就会造成某些下标无用，消耗了table的空间。

### **真正创建新节点** ###

` void createEntry ( int hash, K key, V value, int bucketIndex) { Entry<K,V> e = table[bucketIndex]; table[bucketIndex] = new Entry<>(hash, key, value, e); size++; } 复制代码`

它会先获得链表中的头节点，然后将新节点放进链表头部，旧节点放在新节点后面。

### **数组扩容问题** ###

` void resize ( int newCapacity) { Entry[] oldTable = table; int oldCapacity = oldTable.length; if (oldCapacity == MAXIMUM_CAPACITY) { threshold = Integer.MAX_VALUE; return ; } Entry[] newTable = new Entry[newCapacity]; boolean oldAltHashing = useAltHashing; useAltHashing |= sun.misc.VM.isBooted() && (newCapacity >= Holder.ALTERNATIVE_HASHING_THRESHOLD); boolean rehash = oldAltHashing ^ useAltHashing; //这里是个新数组哦，所以需要将旧数据移到新数组中 transfer(newTable, rehash); table = newTable; threshold = ( int )Math.min(newCapacity * loadFactor, MAXIMUM_CAPACITY + 1 ); } 复制代码`

插入新值的时候，计算极限容量是否已经到达，如果达到了，就调用上面的方法进行扩容，将桶容量增大到两倍。

## **HashMap的get（K）** ##

` public V get (Object key) { if (key == null ) return getForNullKey(); Entry<K,V> entry = getEntry(key); return null == entry ? null : entry.getValue(); } 复制代码`

真正调用的方法：

` final Entry<K,V> getEntry (Object key) { //判断key是否为null int hash = (key == null ) ? 0 : hash(key); //循环链表，找到key键对应的Entry for (Entry<K,V> e = table[indexFor(hash, table.length)]; e != null ; e = e.next) { Object k; if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) return e; } return null ; } 复制代码`

get获取实现比较简单，就是计算出key键的hash值，找到在数组中对应的链表位置，然后循环链表，比较key是否一致。

## **为什么HashMap是无序的** ##

下面是HashMap的迭代器实现：

` private abstract class HashIterator < E > implements Iterator < E > { Entry<K,V> next; // next entry to return int expectedModCount; // For fast-fail int index; // current slot Entry<K,V> current; // current entry HashIterator() { //判断线程是否安全 expectedModCount = modCount; if (size > 0 ) { // advance to first entry Entry[] t = table; //从哈希表数组从上到下，查找第一个不为null的节点，并把next引用指向该节点 while (index < t.length && (next = t[index++]) == null ) ; } } public final boolean hasNext () { return next != null ; } final Entry<K,V> nextEntry () { if (modCount != expectedModCount) throw new ConcurrentModificationException(); Entry<K,V> e = next; if (e == null ) throw new NoSuchElementException(); if ((next = e.next) == null ) { Entry[] t = table; //如果当前节点的下一个节点为null，从节点处往下查找哈希表，找到第一个不为null的节点 while (index < t.length && (next = t[index++]) == null ) ; } current = e; return e; } public void remove () { if (current == null ) throw new IllegalStateException(); if (modCount != expectedModCount) throw new ConcurrentModificationException(); Object k = current.key; current = null ; HashMap. this.removeEntryForKey(k); //判断线程是否安全 expectedModCount = modCount; } } 复制代码`

从上面的代码可以看出，遍历HashMap，都是根据链表的index从上往下进行遍历，由于HashMap的存储规则，后来新加的值有可能在链表的头部，所以遍历HashMap是 **无序** 的。

## **HashMap不是线程安全的** ##

因为看到有 **modCount** 这个字段，查询了资料，发现这个代表这修改次数，对HashMap内容修改都将增加这个值，在迭代器初始化过程中会将这个值赋给迭代器的expectedModCount，在迭代过程中，判断modCount跟expectedModCount是否相等，如果不相等，表示已经有其它线程需改了Map。

有个同步机制： 如果多个线程同时访问一个哈希映射，而其中至少一个线程从结构上修改了该映射，则它必须保持外部同步。

` Map m = Collections.synchronizedMap( new HashMap(...)); 复制代码`

## 参考资料 ##

1、 [Java容器（四）：HashMap（Java 7）的实现原理]( https://link.juejin.im?target=http%3A%2F%2Fblog.csdn.net%2Fjeffleo%2Farticle%2Fdetails%2F54946424 )

2、 [深入理解HashMap（及hash函数的真正巧妙之处）]( https://link.juejin.im?target=http%3A%2F%2Fblog.csdn.net%2Fjiary5201314%2Farticle%2Fdetails%2F51439982 )

3、 [modCount到底是干什么的呢]( https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%2Fnulisaonian%2Fp%2F5946382.html )