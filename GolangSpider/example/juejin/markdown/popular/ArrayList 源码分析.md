# ArrayList 源码分析 #

# ArrayList 源码分析 #

## 前言 ##

ArrayList 算是我们开发中最经常用到的一个集合了，使用起来很方便，对于内部元素的随机访问很快。今天来分析下ArrayList 的源码，本次分析基于 Java1.8 。

## ArrayList 简介 ##

先来看下 ArrayList 的 API 描述：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bf5029d49ba5?imageView2/0/w/1280/h/960/ignore-error/1)

从描述里面来看，ArrayList 是继承于 AbstractList 的，并且实现了 Serializable, Cloneable, Iterable, Collection, List, RandomAccess 这些接口。

* 实现了 Serializable 是序列化接口，因此它支持序列化，能够通过序列化传输。
* 实现了 Cloneable 接口，能被克隆。
* 实现了Iterable 接口，可以被迭代器遍历
* 实现了 Collection ，拥有集合操作的方法
* 实现了 List 接口，拥有增删改查等方法
* 实现了 RandomAccess 随机访问接口，支持快速随机访问，实际上就是通过下标序号进行快速访问。

先大体了解下ArrayList 的特点，然后再从源码的角度去分析：

* ArrayList 底层是一个动态扩容的数组结构,初始容量为 10，每次容量不够的时候，扩容需要增加 1.5 倍的容量（大多数情况下是扩容 1.5 倍的，但是在使用 addAll 的时候，可能有例外。）
* ArrayList 允许存放重复数据，存储顺序按照元素的添加顺序，也允许多个 Null 存在。
* 底层使用 Arrays.copyOf 函数进行扩容，每次扩容都会产生新的数组，和数组中内容的拷贝，所以会耗费性能，所以在多增删的操作的情况可优先考虑 LinkedList。
* ArrayList 并不是一个线程安全的集合。如果集合的增删操作需要保证线程的安全性，可以考虑使用 CopyOnWriteArrayList 或者使Collections.synchronizedList(List l) 函数返回一个线程安全的 ArrayList 类.

## ArrayList 源码分析 ##

### 一些属性 ###

` public class ArrayList < E > extends AbstractList < E > implements List < E >, RandomAccess , Cloneable , java. io. Serializable { // 序列化 ID private static final long serialVersionUID = 8683452581122892189L ; /** * ArrayList 默认的数组容量 */ private static final int DEFAULT_CAPACITY = 10 ; // 一个默认的空数组 private static final Object[] EMPTY_ELEMENTDATA = {}; // 在调用无参构造方法的时候使用该数组 private static final Object[] DEFAULTCAPACITY_EMPTY_ELEMENTDATA = {}; // 存储 ArrayList 元素的数组 // transient 关键字这里简单说一句，被它修饰的成员变量无法被 Serializable 序列化 transient Object[] elementData; // non-private to simplify nested class access // ArrayList 的大小，也就是 elementData 包含的元素个数 private int size; } 复制代码`

### 构造方法 ###

内部几个主要的属性就这些。再来看下构造方法：

` // 指定大小的构造方法，如果传入的是 0 ，直接使用 EMPTY_ELEMENTDATA public ArrayList ( int initialCapacity) { if (initialCapacity > 0 ) { this.elementData = new Object[initialCapacity]; } else if (initialCapacity == 0 ) { this.elementData = EMPTY_ELEMENTDATA; } else { throw new IllegalArgumentException( "Illegal Capacity: " + initialCapacity); } } // 调用该构造方法构造一个默认大小为 10 的数组，但是此时大小未指定， // 还是空的，在第一次 add 的时候指定 public ArrayList () { this.elementData = DEFAULTCAPACITY_EMPTY_ELEMENTDATA; } // 传入一个集合类 // 首先直接利用Collection.toArray()方法得到一个对象数组，并赋值给elementData public ArrayList (Collection<? extends E> c) { elementData = c.toArray(); if ((size = elementData.length) != 0 ) { // c.toArray 出错的时候，使用Arrays.copyOf 生成一个新数组赋值给 elementData if (elementData.getClass() != Object[].class) elementData = Arrays.copyOf(elementData, size, Object[].class); } else { //如果集合c元素数量为0，则将空数组EMPTY_ELEMENTDATA赋值给elementData this.elementData = EMPTY_ELEMENTDATA; } } 复制代码`

可以看到，不管是调用哪个构造方法，都会初始化内部 elementData 。

### add 方法 ###

接下来从最常用的 add 方法看起：

` public boolean add (E e) { ensureCapacityInternal(size + 1 ); // Increments modCount!! elementData[size++] = e; return true ; } 复制代码`

执行 ensureCapacityInternal(size + 1) 确认内部容量

` private void ensureCapacityInternal ( int minCapacity) { // 如果创建 ArrayList 时候，使用的无参的构造方法，那么就取默认容量 10 和最小需要的容量（当前 size + 1 ）中大的一个确定需要的容量。 if (elementData == DEFAULTCAPACITY_EMPTY_ELEMENTDATA) { minCapacity = Math.max(DEFAULT_CAPACITY, minCapacity); } ensureExplicitCapacity(minCapacity); } 复制代码`

其实这里的 size 的默认值是 0 ，所以在使用默认构造方法创建 ArrayList 以后第一次执行 ensureCapacityInternal 的时候，要扩容的容量就是 DEFAULT_CAPACITY = 10；

` private void ensureExplicitCapacity ( int minCapacity) { // 修改 +1 modCount++; // 如果 minCapacity 比当前容量大， 就执行grow 扩容 if (minCapacity - elementData.length > 0 ) grow(minCapacity); } private void grow ( int minCapacity) { // 拿到当前的容量 int oldCapacity = elementData.length; // oldCapacity >> 1 意思就是 oldCapacity/2，所以新容量就是增加 1/2. int newCapacity = oldCapacity + (oldCapacity >> 1 ); // 如果新容量小于，需要最小扩容的容量，以需要最小容量为准扩容 if (newCapacity - minCapacity < 0 ) newCapacity = minCapacity; // 如果新容量大于允许的最大容量，则以 Inerger 的最大值进行扩容 if (newCapacity - MAX_ARRAY_SIZE > 0 ) newCapacity = hugeCapacity(minCapacity); // 使用 Arrays.copyOf 函数进行扩容。 elementData = Arrays.copyOf(elementData, newCapacity); } // 允许的最大容量 private static final int MAX_ARRAY_SIZE = Integer.MAX_VALUE - 8 ; private static int hugeCapacity ( int minCapacity) { if (minCapacity < 0 ) // overflow throw new OutOfMemoryError(); return (minCapacity > MAX_ARRAY_SIZE) ? Integer.MAX_VALUE : MAX_ARRAY_SIZE; } 复制代码`

根据上面的代码可以看出，如果我们默认扩容 1.5 倍的容量比最小需要的容量（minCapacity）还小，那么就使用 minCapacity 进行扩容。所以并不是每次都是以 1.5 倍进行扩容的。

上面讲了扩容，扩容好了以后，就执行

` elementData[size++] = e; return true ; 复制代码`

进行赋值操作，就完成了一次数据的添加。

再来看下在指定位置添加一个元素:

` public void add ( int index, E element) { if (index > size || index < 0 ) throw new IndexOutOfBoundsException(outOfBoundsMsg(index)); ensureCapacityInternal(size + 1 ); // Increments modCount!! System.arraycopy(elementData, index, elementData, index + 1 , size - index); elementData[index] = element; size++; } 复制代码`

先判断传入的位置是够越界。越界就抛出异常

然后确认需不需要扩容，然后再通过 System.arraycopy 方法进行拷贝。

需要注意的是 size - index 表示的是需要移动的元素的数量。也就是 index 后面的元素都要进行移动，这也就是插入效率低的一个原因，在指定位置插入数据，那么这个位置后面的数据都要移动，如果是在第 0 个位置插入，意味着所有的元素都要移动。

上面的 add 方法分析完了，然后再来看下另一个常见的 addAll 方法:

### addAll 方法 ###

先看第一个 addAll

` public boolean addAll (Collection<? extends E> c) { Object[] a = c.toArray(); int numNew = a.length; ensureCapacityInternal(size + numNew); // Increments modCount System.arraycopy(a, 0 , elementData, size, numNew); size += numNew; return numNew != 0 ; } 复制代码`

这里也很简单，先转成数组，拿到长度进行扩容。然后利用 System.arraycopy 函数把传进来的数组拷贝到现有数组里面。

再来看第二个 addAll 方法:

这个是在指定位置添加一个集合。

` public boolean addAll ( int index, Collection<? extends E> c) { if (index > size || index < 0 ) throw new IndexOutOfBoundsException(outOfBoundsMsg(index)); Object[] a = c.toArray(); int numNew = a.length; ensureCapacityInternal(size + numNew); // Increments modCount int numMoved = size - index; if (numMoved > 0 ) System.arraycopy(elementData, index, elementData, index + numNew, numMoved); System.arraycopy(a, 0 , elementData, index, numNew); size += numNew; return numNew != 0 ; } 复制代码`

这里也很简单，基本和使用 add 方法在指定位置添加一个元素差不多。就不在分析了。接下来看看删除相关的。

### remove 方法 ###

看下源码：

删除一个指定位置的元素：

` public E remove ( int index) { if (index >= size) throw new IndexOutOfBoundsException(outOfBoundsMsg(index)); modCount++; E oldValue = (E) elementData[index]; int numMoved = size - index - 1 ; if (numMoved > 0 ) System.arraycopy(elementData, index+ 1 , elementData, index, numMoved); elementData[--size] = null ; // clear to let GC do its work return oldValue; } 复制代码`

很简单，先判断是够越界，越界抛出异常。

然后先把要删除的元素拿出来，存储在 oldValue ，这里看到了一个 numMoved ，也就是删除一个元素需要移动的元素的数量。然后执行 System.arraycopy 进行数组的移动，这里只移动删除的 index 后面的元素，统统向前进一位。然后把数组中最后一个元素置为 null，返回删除的元素。

删除一个指定的元素：

` public boolean remove (Object o) { if (o == null ) { for ( int index = 0 ; index < size; index++) if (elementData[index] == null ) { fastRemove(index); return true ; } } else { for ( int index = 0 ; index < size; index++) if (o.equals(elementData[index])) { fastRemove(index); return true ; } } return false ; } private void fastRemove ( int index) { modCount++; int numMoved = size - index - 1 ; if (numMoved > 0 ) System.arraycopy(elementData, index+ 1 , elementData, index, numMoved); elementData[--size] = null ; // clear to let GC do its work } 复制代码`

这里分两种情况，

* 删除的元素为 null ，根据循环查找到第一个为 null 的元素，然后执行 fastRemove(index) 删除之后，返回 true 删除成功，可以看到这里的 fastRemove 方法和 remove(int index) 是比较类似的，就不讲了。
* 删除的元素不为 null ，和为 null 逻辑差不多，就是对元素的判断不同，这里使用的 ` o.equals(elementData[index])` ，而为 null 的时候，使用 ` elementData[index] == null`

### set 方法 ###

set 方法就是在指定位置改变一个元素的值

` public E set ( int index, E element) { if (index >= size) throw new IndexOutOfBoundsException(outOfBoundsMsg(index)); E oldValue = (E) elementData[index]; elementData[index] = element; return oldValue; } 复制代码`

同样，先判断是否越界，越界抛出异常，没越界直接修改值，把旧值返回。

### get 方法 ###

取某个位置的元素：

` public E get ( int index) { if (index >= size) throw new IndexOutOfBoundsException(outOfBoundsMsg(index)); return (E) elementData[index]; } 复制代码`

同样，先判断是否越界，越界抛出异常，没越界属于数组的操作，直接返回指定位置的值。

### clear 方法 ###

清除数组中的所有元素：

` public void clear () { modCount++; // clear to let GC do its work for ( int i = 0 ; i < size; i++) elementData[i] = null ; size = 0 ; } 复制代码`

可以看到是循环把数组中的每个元素置为 null，可以让 gc 回收，然后再把数组的长度置为 0 。下次 add 的时候，还是直接扩容到长度为 10.

### indexOf 方法 ###

返回元素在集合中的位置

` public int indexOf (Object o) { if (o == null ) { for ( int i = 0 ; i < size; i++) if (elementData[i]== null ) return i; } else { for ( int i = 0 ; i < size; i++) if (o.equals(elementData[i])) return i; } return - 1 ; } 复制代码`

和 remove 的时候类似，分为两种情况处理。饭后返回元素在数组中的位置。

最后元素最后出现的位置

` public int lastIndexOf (Object o) { if (o == null ) { for ( int i = size- 1 ; i >= 0 ; i--) if (elementData[i]== null ) return i; } else { for ( int i = size- 1 ; i >= 0 ; i--) if (o.equals(elementData[i])) return i; } return - 1 ; } 复制代码`

和 indexOf 操作一样，只不过是倒序查找第一个元素出现的位置

### isEmpty 方法 ###

是否为空

` public boolean isEmpty () { return size == 0 ; } 复制代码`

可以看到是根据 size 来判断的，即使你把 ArrayList 中的每个元素置为 null，但是 size 不为 0 的话，isEmpty 依旧返回 false。

## 总结 ##

通过上面的分析可以再次总结下结论：

* ArrayList 底层是一个动态扩容的数组结构,初始容量为 10，每次容量不够的时候，扩容需要增加 1.5 倍的容量
* 增加（add）和删除（remove）操作会改变 modCount，但是查找（get）和修改（set）不会修改
* 从上面可以看出，增加和删除都可能涉及到扩容操作，扩容和删除会移动已有元素的位置，比较低效，但是查找和修改时很高效的。
* 从上面看出，ArrayList 对 null 元素是支持的，并且不会限制数量，也不会限制重复元素的增加
* 全文没见 Synchronized 关键字，也没有其它保证线程安全的操作，所以是线程不安全的，可以使用CopyOnWriteArrayList 或者使Collections.synchronizedList(List l) 函数返回一个线程安全的 ArrayList 类来保证线程安全。

使用建议：

* 如果是修改和获取操作比较多，建议使用 ArrayList ，效率高。
* 如果增加和删除操作较多，建议使用 LinkedList（下篇分析），但是如果增加和删除的操作都在队尾，不涉及到元素的移动，还是建议使用 ArrayList ，毕竟 ArrayList 的查找和修改的效率还是蛮高的。
* 使用的时候，如果确定元素的大小，最好能设置下 ArrayList 的容量，避免扩容浪费空间

这篇就讲到这里，下篇来看下 LinkedList。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bf502a3a6d72?imageView2/0/w/1280/h/960/ignore-error/1)