# 从源码分析HashMap #

## HashMap简介 ##

如果你了解过数据结构，就应该理解散列表的概念，类似于数学中函数的概念，通过一个自变量映射到一个因变量上。如果把键值当作自变量，对应的值当作因变量，这样我们就得到了一些键值对，保存这些键值对的数据结构我们就叫做散列表，在Java中，拥有一个Map接口来提供操作这种数据结构的方法

Java的数据结构接口分为两个大类，Collection和Map，我们这里要讲的HashMap就是Map分支的一个实例对象

## 源码分析 ##

#### HashMap的结构 ####

我们点进HashMap的源码，可以看到HashMap继承了AbstractMap抽象类，实现了Map接口，同时还实现了Cloneable和Serializable接口以提供浅复制和序列化

![hashmap](https://user-gold-cdn.xitu.io/2019/6/4/16b2286d1360dcb2?imageView2/0/w/1280/h/960/ignore-error/1)

我们的重点是HashMap，所以把重心放在HashMap本身上，只需要知道Map和AbstractMap提供了一些方法接口和默认实现即可 类似于ArrayList的数组，HashMap也有存储数据的对象，是一个Node类型的数组

` transient Node<K,V>[] table; 复制代码`

这个Node类型是什么呢，我们点进去，发现是一个HashMap的静态内部类

` static class Node < K , V > implements Map. Entry < K , V > { final int hash; final K key; V value; Node<K,V> next; 复制代码`

有四个属性，我们依次解释

* hash：键的hash值。注意，这里的hash值要和键的hashcode值区分开，下面还会提到
* key：键值对的键
* value：键值对的值
* next：链表指针

这里的next指针可能有些人不了解，这里先解释一下Hash冲突

计算机不像人类一样可以清晰地分辨出哪些对象属于一类，在HashMap中，通过一个key的hash值来唯一确定一个key在数组中的索引位置。但是Java并不能保证任意两个逻辑意义不同的对象一定拥有不同的hash值，所以就有可能发生有两个在我们看来完全不同的对象，却有着相同的hash值，但是一个索引位置只能放一个元素，这时候就发生了 **哈希冲突**

这里我只说一下HashMap中的解决办法，即 **数组-链表** 法，将发生冲突的节点（拥有不同key值但是相同hash值的节点）以链表的形式挂在数组的某一个索引位置处，如下图

![](https://user-gold-cdn.xitu.io/2019/6/4/16b228663c31700b?imageView2/0/w/1280/h/960/ignore-error/1) 回来继续看HashMap的结构，还剩下三个常量

` /** 默认初始容量 */ static final int DEFAULT_INITIAL_CAPACITY = 1 << 4 ; // aka 16 /** 最大容量 */ static final int MAXIMUM_CAPACITY = 1 << 30 ; /** 负载因子 */ static final float DEFAULT_LOAD_FACTOR = 0.75f ; 复制代码`

初始容量和最大容量很好理解，都是为了控制数组长度的常量，负载因子如果学过数据结构应该也很好理解，我们在存放KV时，并不能将数组占满，而是最多只能达到 **数组长度 * 负载因子** 的长度，这样可以尽量避免频繁的哈希冲突

还有其余一些属性，比如临界值threshold，我们遇到了再说

### HashMap的操作 ###

##### 扩容操作 #####

因为底层数据结构是一个数组，并不能弹性地自动扩容，所以HashMap提供了一个扩容方法：resize。这个方法返回一个Node对象的数组，也就是扩容后的新数组

方法很长，我们一点一点看

` Node<K,V>[] oldTab = table; int oldCap = (oldTab == null ) ? 0 : oldTab.length; int oldThr = threshold; int newCap, newThr = 0 ; 复制代码`

首先是定义了

* oldTab：扩容前的数组
* oldCap：扩容前数组的长度
* oldThr：扩容前的临界值（超过临界值会进行扩容）（计算公式：数组长度*负载因子）
* newCap：扩容后的长度
* newThr：扩容后的临界值

其次是一些边界条件判断：

` if (oldCap > 0 ) { if (oldCap >= MAXIMUM_CAPACITY) { threshold = Integer.MAX_VALUE; return oldTab; } else if ((newCap = oldCap << 1 ) < MAXIMUM_CAPACITY && oldCap >= DEFAULT_INITIAL_CAPACITY) newThr = oldThr << 1 ; // double threshold } else if (oldThr > 0 ) // initial capacity was placed in threshold newCap = oldThr; else { // zero initial threshold signifies using defaults newCap = DEFAULT_INITIAL_CAPACITY; newThr = ( int )(DEFAULT_LOAD_FACTOR * DEFAULT_INITIAL_CAPACITY); } if (newThr == 0 ) { float ft = ( float )newCap * loadFactor; newThr = (newCap < MAXIMUM_CAPACITY && ft < ( float )MAXIMUM_CAPACITY ? ( int )ft : Integer.MAX_VALUE); } 复制代码`

这几个条件，具体表示为以下含义：

* 如果之前数组长度已经达到最大容量，则取消负载因子的限制（将临界值设为最大值）
* 如果之前数组长度扩大到两倍后（每次扩容到之前的两倍），没有超过最大限制，新数组的长度为之前的两倍
* 如果之前数组长度为0，但是临界值大于0，则让新数组的长度等于临界值
* 如果之前的数组长度为0，且临界值也为0，则把新数组的长度和临界值设为默认值

然后是赋值部分

` threshold = newThr; @SuppressWarnings ({ "rawtypes" , "unchecked" }) Node<K,V>[] newTab = (Node<K,V>[]) new Node[newCap]; table = newTab; 复制代码`

把新的临界值和新的空数组赋值过去

最后就是核心部分，即将原数组的内容拷贝到新数组去。整段代码逻辑复杂，我们一部分一部分的来看

` if (oldTab != null ) { for ( int j = 0 ; j < oldCap; ++j) { Node<K,V> e; 复制代码`

首先进行了一次非空判断，最外层是一个循环，循环次数是扩容前数组的长度，即它的目的是遍历一遍老数组

然后是接着是两个终止条件（单轮次的终止而不是整个循环的终止）

` if ((e = oldTab[j]) != null ) { oldTab[j] = null ; if (e.next == null ) newTab[e.hash & (newCap - 1 )] = e; 复制代码`

* 如果当前索引位置为空，跳过该节点，进入下一次循环
* 如果当前索引位置节点的next节点为空，即索引位置只有一个节点，则直接将该节点的值赋到新数组对应的位置上，然后进入下一次循环

这里需要再解释一下，键的hash值并不是对应的索引位置，而是通过和 **数组长度-1** 作 **与运算** 得到最终的索引位置

接下来，判断是否为TreeNode节点

` else if (e instanceof TreeNode) ((TreeNode<K,V>)e).split( this , newTab, j, oldCap); 复制代码`

TreeNode节点也是一个内部类，表示一棵红黑树（不理解红黑树的就当成二叉树即可）。因为链表长度如果过长，那么索引的效率就会降低，所以将其链表转换为红黑树。这里如果发现当前索引节点是树节点，则通过对应的方法进行转移

接下来我们只看这一段核心代码

` else { // preserve order Node<K,V> loHead = null , loTail = null ; Node<K,V> hiHead = null , hiTail = null ; Node<K,V> next; do { next = e.next; if ((e.hash & oldCap) == 0 ) { if (loTail == null ) loHead = e; else loTail.next = e; loTail = e; } else { if (hiTail == null ) hiHead = e; else hiTail.next = e; hiTail = e; } } while ((e = next) != null ); if (loTail != null ) { loTail.next = null ; newTab[j] = loHead; } if (hiTail != null ) { hiTail.next = null ; newTab[j + oldCap] = hiHead; } } 复制代码`

刚看可能大家都是一头雾水，主要是因为有一个点没有理解，我简要讲解一下：

* 采用e.hash & **oldCap** 而不是 **oldCap-1** ，是因为这里不是为了计算索引位置，而是判断是否需要进行移动

这样就很清晰了，这段代码让loHead=>loTail这条链连接所有新老数组中索引位置一样的元素，而hiHead=>hiTail这条链连接所有新老数组中位置改变的数组

不改变位置的链上元素还是移动到新数组对应位置上，而改变位置的链上元素则移动到 **j+oldCap** 的索引位置

最后返回新数组作为结果

##### 取值方法 #####

在HashMap中，通过调用get()方法即可获取某个key值对应的value，在HashMap内部，这个方法具体实现如下

` public V get (Object key) { Node<K,V> e; return (e = getNode(hash(key), key)) == null ? null : e.value; } 复制代码`

get()方法内部调用getNode方法获取key值对应的节点，这个方法需要传入两个参数，key和key的哈希值，这个hash()方法的具体实现如下

` static final int hash (Object key) { int h; return (key == null ) ? 0 : (h = key.hashCode()) ^ (h >>> 16 ); } 复制代码`

显然，并不是直接调用key的hashCode()方法，而是将key的hashCode的高16位和低16位做异或运算，来充分运用高位和低位的信息，毕竟只取低位或只取高位的话会发生大量的碰撞情况

好的继续说getNode方法，方法很简练，如下

` final Node<K,V> getNode ( int hash, Object key) { Node<K,V>[] tab; Node<K,V> first, e; int n; K k; // 临界判断，需要保证数组非空以及索引位置非空 if ((tab = table) != null && (n = tab.length) > 0 && (first = tab[(n - 1 ) & hash]) != null ) { if (first.hash == hash && // 检查索引位置的第一个节点的key值是否为所求 ((k = first.key) == key || (key != null && key.equals(k)))) return first; if ((e = first.next) != null ) { // 如果是树节点，则通过对应的方法来获取 if (first instanceof TreeNode) return ((TreeNode<K,V>)first).getTreeNode(hash, key); // 否则就扫描整条链表 do { if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) return e; } while ((e = e.next) != null ); } } return null ; } 复制代码`

本身不难理解，获取节点的过程可分为以下几步

* 临界判断，如果是数组或索引位置为空就直接返回，否则进入下一步
* 检查第一个节点是否为所求节点，如果是，则直接返回，否则进入下一步
* 如果索引位置是树节点，通过红黑树对应的方法来查找节点，否则遍历链表来寻找

##### 插入方法 #####

和get方法类似，put方法也是调用了其他方法来进行插入，如下

` public V put (K key, V value) { return putVal(hash(key), key, value, false , true ); } 复制代码`

下面是具体putVal方法：

` final V putVal ( int hash, K key, V value, boolean onlyIfAbsent, boolean evict) { Node<K,V>[] tab; Node<K,V> p; int n, i; if ((tab = table) == null || (n = tab.length) == 0 ) n = (tab = resize()).length; if ((p = tab[i = (n - 1 ) & hash]) == null ) tab[i] = newNode(hash, key, value, null ); else { Node<K,V> e; K k; if (p.hash == hash && ((k = p.key) == key || (key != null && key.equals(k)))) e = p; else if (p instanceof TreeNode) e = ((TreeNode<K,V>)p).putTreeVal( this , tab, hash, key, value); else { for ( int binCount = 0 ; ; ++binCount) { if ((e = p.next) == null ) { p.next = newNode(hash, key, value, null ); if (binCount >= TREEIFY_THRESHOLD - 1 ) // -1 for 1st treeifyBin(tab, hash); break ; } if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) break ; p = e; } } if (e != null ) { // existing mapping for key V oldValue = e.value; if (!onlyIfAbsent || oldValue == null ) e.value = value; afterNodeAccess(e); return oldValue; } } ++modCount; if (++size > threshold) resize(); afterNodeInsertion(evict); return null ; } 复制代码`

前三个参数很好理解，问题是后面两个参数，点进putVal方法，可以看到这两个参数的含义依次是：

* onlyIfAbsent：如果为true，则不覆盖原值
* evict：如果为false，则table处于创建模式

onlyIfAbsent很好理解，主要是evict这个变量很奇怪，其实我们接下来看源码就会发现，这个变量仅仅触发了插入之后的一个回调方法，这个方法体为空，也就是这个变量没有实际意义

把重心放在putVal方法上，一开始还是两个临界判断，如下

` if ((tab = table) == null || (n = tab.length) == 0 ) n = (tab = resize()).length; if ((p = tab[i = (n - 1 ) & hash]) == null ) tab[i] = newNode(hash, key, value, null ); 复制代码`

* 如果数组为空，触发扩容方法
* 如果索引位置为空，直接在该位置插入

至于我为什么把插入语句当成临界判断条件，是因为putVal方法的核心不在这里，我们接着看

` else { Node<K,V> e; K k; if (p.hash == hash && ((k = p.key) == key || (key != null && key.equals(k)))) e = p; else if (p instanceof TreeNode) e = ((TreeNode<K,V>)p).putTreeVal( this , tab, hash, key, value); else { for ( int binCount = 0 ; ; ++binCount) { if ((e = p.next) == null ) { p.next = newNode(hash, key, value, null ); if (binCount >= TREEIFY_THRESHOLD - 1 ) // -1 for 1st treeifyBin(tab, hash); break ; } if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) break ; p = e; } } // 说明存在拥有相同key值的节点 if (e != null ) { V oldValue = e.value; // 如果我们设置了onlyIfAbsent为true或原value为空 if (!onlyIfAbsent || oldValue == null ) e.value = value; // 一个空回调方法 afterNodeAccess(e); return oldValue; } } 复制代码`

如果上述两个条件不满足，则说明需要在链表/红黑树节点上进行插入，我们一部分一部分的看

` else if (p instanceof TreeNode) e = ((TreeNode<K,V>)p).putTreeVal( this , tab, hash, key, value); 复制代码`

p节点是当前数组索引位置的节点，假如这个节点是树节点，会调用putTreeVal方法进行插入

否则说明是链表结构，则需要找到链表尾进行插入，如下

` for ( int binCount = 0 ; ; ++binCount) { // 如果找到链表尾就直接插入 if ((e = p.next) == null ) { p.next = newNode(hash, key, value, null ); if (binCount >= TREEIFY_THRESHOLD - 1 ) // -1 for 1st // 如果大于链表长度阈值，就将其转换为红黑树 treeifyBin(tab, hash); break ; } // 判断条件：hash值相等，且key值（物理/逻辑）相等 if (e.hash == hash && ((k = e.key) == key || (key != null && key.equals(k)))) break ; p = e; } 复制代码`

这里提醒一点，TREEIFY_THRESHOLD是链表长度的阈值，也就是说链表长度最多只能为TREEIFY_THRESHOLD，超过了这个值就会被转换成红黑树结构，这个值固定为8

最后，判断是否需要进行扩容

` ++modCount; if (++size > threshold) resize(); // 空回调函数 afterNodeInsertion(evict); return null ; 复制代码`

注意，只要不触发原值覆盖条件，putVal方法一定只会返回null