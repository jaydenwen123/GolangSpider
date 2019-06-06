# 排序算法 Java实现 #

## 选择排序 ##

### 核心思想 ###

选择最小元素，与第一个元素交换位置；剩下的元素中选择最小元素，与当前剩余元素的最前边的元素交换位置。

### 分析 ###

选择排序的比较次数与序列的初始排序无关， **比较次数都是N(N-1)/2** 。

移动次数最多只有n-1次。

因此，时间复杂度为O(N^2)，无论输入是否有序都是如此，输入的顺序只决定了交换的次数，但是比较的次数不变。

选择排序是不稳定的，比如5 6 5 3的情况。

### 代码 ###

` public class SelectionSort { public void selectionSort ( int [] nums) { if (nums== null ) return ; for ( int i= 0 ;i<nums.length;i++) { int index = i; for ( int j = i; j < nums.length; j++) { if (nums[j] < nums[index]) { index = j; } } swap(nums, i, index); } } } 复制代码`

## 冒泡排序： ##

### 核心思想 ###

从左到右不断交换相邻逆序的元素，这样一趟下来把最大的元素放到了最右侧。不断重复这个过程，知道一次循环中没有发生交换，说明已经有序，退出。

### 分析 ###

* 当原始序列有序，比较次数为 n-1 ，移动次数为0，因此 **最好情况下时间复杂度为 O(N)** 。
* 当逆序排序时，比较次数为 N(N-1)/2，移动次数为 3N(N-1)/2，因此 **最坏情况下时间复杂度为 O(N^2)** 。
* 平均时间复杂度为 O(N^2)。

元素两两交换时，相同元素前后顺序没有改变，因此具有稳定性。

### 代码 ###

` public class BubbleSort { public void bubbleSort ( int [] nums) { for ( int i=nums.length- 1 ;i> 0 ;i--){ boolean sorted= false ; for ( int j= 0 ;j<i;j++){ if (nums[j]>nums[j+ 1 ]){ Sort.swap(nums,j,j+ 1 ); sorted= true ; } } if (!sorted) break ; } } 复制代码`

## 插入排序 ##

### 核心思想 ###

每次将当前元素插入到左侧已经排好序的数组中，使得插入之后左侧数组依然有序。

### 分析 ###

因为插入排序每次只能交换相邻元素，令逆序数量减少1，因此交换次数等于逆序数量。

因此，插入排序的复杂度取决于数组的初始顺序。

* 数组已经有序，需要 N-1 次比较和0次交换，时间复杂度为 O(N)。
* 数组完全逆序，需要 N(N-1)/2 次比较和交换 N(N-1)/2 次，时间复杂度为 O(N^2)
* 平均情况下，时间复杂度为 O(N^2)

插入排序具有稳定性

### 代码 ###

` public class InsertionSort { public void insertionSort ( int [] nums) { for ( int i= 1 ;i<nums.length;i++){ for ( int j=i;j> 0 ;j--){ if (nums[j]<nums[j- 1 ]) swap(nums,j,j- 1 ); else break ; //已经放到正确位置上了 } } } } 复制代码`

## 希尔排序 ##

对于大规模的数组，插入排序很慢，因为它只能交换相邻的元素，每次只能将逆序数量减少1。

### 核心思想 ###

希尔排序为了解决插入排序的局限性，通过交换不相邻的元素，每次将逆序数量减少大于1。希尔排序使用插入排序对间隔为 H 的序列进行排序，不断减少 H 直到 H=1 ，最终使得整个数组是有序的。

### 时间复杂度 ###

希尔排序的时间复杂度难以确定，并且 H 的选择也会改变其时间复杂度。

希尔排序的时间复杂度是低于 O(N^2) 的，高级排序算法只比希尔排序快两倍左右。

### 稳定性 ###

希尔排序不具备稳定性。

### 代码 ###

` public class ShellSort { public void shellSort ( int [] nums) { int N=nums.length; int h= 1 ; while (h<N/ 3 ){ h= 3 *h+ 1 ; } while (h>= 1 ){ for ( int i=h;i<N;i++){ for ( int j=i;j> 0 ;j--){ if (nums[j]<nums[j- 1 ]){ swap(nums,j,j- 1 ); } else { break ; //已经放到正确位置上了 } } } } } } 复制代码`

## 归并排序 ##

### 核心思想 ###

将数组分为两部分，分别进行排序，然后进行归并。

### 归并方法 ###

` public void merge ( int [] nums, int left, int mid, int right) { int p1 = left, p2 = mid + 1 ; int [] tmp = new int [right-left+ 1 ]; int cur= 0 ; //两个指针分别指向左右两个子数组，选择更小者放入辅助数组 while (p1<=mid&&p2<=right){ if (nums[p1]<nums[p2]){ tmp[cur++]=nums[p1++]; } else { tmp[cur++]=nums[p2++]; } } //将还有剩余的数组放入到辅助数组 while (p1<=mid){ tmp[cur++]=nums[p1++]; } while (p2<=right){ tmp[cur++]=nums[p2++]; } //拷贝 for ( int i= 0 ;i<tmp.length;i++){ nums[left+i]=tmp[i]; } } 复制代码`

### 代码实现 ###

#### 递归方法：自顶向下 ####

通过递归调用，自顶向下将一个大数组分成两个小数组进行求解。

` public void up2DownMergeSort ( int [] nums, int left, int right) { if (left==right) return ; int mid=left+(right-left)/ 2 ; mergeSort(nums,left,mid); mergeSort(nums,mid+ 1 ,right); merge(nums,left,mid,right); } 复制代码`

#### 非递归：自底向上 ####

` public void down2UpMergeSort(int[] nums) { int N = nums.length; for (int sz = 1; sz < N; sz += sz) { for (int lo = 0; lo < N - sz; lo += sz + sz) { merge(nums, lo, lo + sz - 1, Math.min(lo + sz + sz - 1, N - 1)); } } } 复制代码`

### 分析 ###

把一个规模为N的问题分解成两个规模分别为 N/2 的子问题，合并的时间复杂度为 O(N)。T(N)=2T(N/2)+O(N)。

得到其时间复杂度为 O(NlogN)，并且在最坏、最好和平均情况下时间复杂度相同。

归并排序需要 O(N) 的空间复杂度。

归并排序具有稳定性。

## 快速排序 ##

### 核心思想 ###

快速排序通过一个切分元素 pivot 将数组分为两个子数组，左子数组小于等于切分元素，右子数组大于等于切分元素，将子数组分别进行排序，最终整个排序。

### partition ###

取 a[l] 作为切分元素，然后从数组的左端向右扫描直到找到第一个大于等于它的元素，再从数组的右端向左扫描找到第一个小于它的元素，交换这两个元素。不断进行这个过程，就可以保证左指针 i 的左侧元素都不大于切分元素，右指针 j 的右侧元素都不小于切分元素。当两个指针相遇时，将切分元素 a[l] 和 a[j] 交换位置。

` private int partition ( int [] nums, int left, int right) { int p1=left,p2=right; int pivot=nums[left]; while (p1<p2){ while (nums[p1++]<pivot&&p1<=right); while (nums[p2--]>pivot&&p2>=left); swap(nums,p1,p2); } swap(nums,left,p2); return p2; } 复制代码`

### 代码实现 ###

` public void sort (T[] nums, int l, int h) { if (h <= l) return ; int j = partition(nums, l, h); sort(nums, l, j - 1 ); sort(nums, j + 1 , h); } 复制代码`

### 分析 ###

最好的情况下，每次都正好将数组对半分，递归调用次数最少，复杂度为 O(NlogN)。

最坏情况下，是有序数组，每次只切分了一个元素，时间复杂度为 O(N^2)。为了防止这种情况，在进行快速排序时需要先随机打乱数组。

不具有稳定性。

### 改进 ###

* 切换到插入排序：递归的子数组规模小时，用插入排序。
* 三数取中：最好的情况下每次取中位数作为切分元素，计算中位数代价比较高，采用取三个元素，将中位数作为切分元素。

### 三路快排 ###

对于有大量重复元素的数组，将数组分为小于、等于、大于三部分，对于有大量重复元素的随机数组可以在线性时间内完成排序。

` public void threeWayQuickSort ( int [] nums, int left, int right) { if (right<=left) return ; int lt=left,cur=left+ 1 ,gt=right; int pivot=nums[left]; while (cur<=gt){ if (nums[cur]<pivot){ swap(nums,lt++,cur++); } else if (nums[cur]>pivot){ swap(nums,cur,gt--); } else { cur++; } } threeWayQuickSort(nums,left,lt- 1 ); threeWayQuickSort(nums,gt+ 1 ,right); } 复制代码`

### 基于 partition 的快速查找 ###

利用 partition() 可以在线性时间复杂度找到数组的第 K 个元素。

假设每次能将数组二分，那么比较的总次数为 (N+N/2+N/4+..)，直到找到第 k 个元素，这个和显然小于 2N。

` public int select ( int [] nums, int k) { int l = 0 , h = nums.length - 1 ; while (h > l) { int j = partition(nums, l, h); if (j == k) { return nums[k]; } else if (j > k) { h = j - 1 ; } else { l = j + 1 ; } } return nums[k]; } 复制代码`

## 堆排序 ##

### 堆 ###

堆可以用数组来表示，这是因为堆是完全二叉树，而完全二叉树很容易就存储在数组中。位置 k 的节点的父节点位置为 k/2，而它的两个子节点的位置分别为 2k 和 2k+1。在这里，从下标为1的索引开始 的位置，是为了更清晰地描述节点的位置关系。

### 上浮和下沉 ###

当一个节点比父节点大，不断交换这两个节点，直到将节点放到位置上，这种操作称为上浮。

` private void shiftUp ( int k) { while (k > 1 && heap[k / 2 ] < heap[k]) { swap(k / 2 , k); k = k / 2 ; } } 复制代码`

当一个节点比子节点小，不断向下进行比较和交换，当一个基点有两个子节点，与最大节点进行交换。这种操作称为下沉。

` private void shiftDown ( int k) { while ( 2 *k<=size){ int j= 2 *k; if (j<size&&heap[j]<heap[j+ 1 ]) j++; if (heap[k]<heap[j]) break ; swap(k,j); k=j; } } 复制代码`

### 堆排序 ###

把最大元素和当前堆中数组的最后一个元素交换位置，并且不删除它，那么就可以得到一个从尾到头的递减序列。

**构建堆** 建立堆最直接的方法是从左到右遍历数组进行上浮操作。一个更高效的方法是从右到左进行下沉操作。叶子节点不需要进行下沉操作，可以忽略，因此只需要遍历一半的元素即可。

**交换堆顶和最坏一个元素，进行下沉操作，维持堆的性质。**

` public class HeapSort { public void sort ( int [] nums) { int N=nums.length- 1 ; for ( int k=N/ 2 ;k>= 1 ;k--){ shiftDown(nums,k,N); } while (N> 1 ){ swap(nums, 1 ,N--); shiftDown(nums, 1 ,N); } System.out.println(Arrays.toString(nums)); } private void shiftDown ( int [] heap, int k, int N) { while ( 2 *k<=N){ int j= 2 *k; if (j<N&&heap[j]<heap[j+ 1 ]) j++; if (heap[k]>=heap[j]) break ; swap(heap,k,j); k=j; } } private void swap ( int [] nums, int i, int j) { int t=nums[i]; nums[i]=nums[j]; nums[j]=t; } } 复制代码`

### 分析 ###

建立堆的时间复杂度是O(N)。

一个堆的高度为 logN, 因此在堆中插入元素和删除最大元素的复杂度都是 logN。

在堆排序中，对N个节点进行下沉操作，复杂度为 O(NlogN)。

现代操作系统很少使用堆排序，因为它无法利用局部性原理进行缓存，也就是数组元素很少和相邻的元素进行比较和交换。

## 比较 ##

+----------+----------------+----------------+----------------+--------------+--------+--------------------------------------------------------------+
| 排序算法 | 最好时间复杂度 | 平均时间复杂度 | 最坏时间复杂度 |  空间复杂度  | 稳定性 |                           适用场景                           |
+----------+----------------+----------------+----------------+--------------+--------+--------------------------------------------------------------+
| 冒泡排序 | O(N)           | O(N^2)         | O(N^2)         | O(1)         | 稳定   |                                                              |
| 选择排序 | O(N)           | O(N^2)         | O(N^2)         | O(1)         | 不稳定 | 运行时间和输入无关，数据移动次数最少，数据量较小的时候适用。 |
| 插入排序 | O(N)           | O(N^2)         | O(N^2)         | O(1)         | 稳定   | 数据量小、大部分已经被排序                                   |
| 希尔排序 | O(N)           | O(N^1.3)       | O(N^2)         | O(1)         | 不稳定 |                                                              |
| 快速排序 | O(NlogN)       | O(NlogN)       | O(N^2)         | O(logN)-O(N) | 不稳定 | 最快的通用排序算法，大多数情况下的最佳选择                   |
| 归并排序 | O(NlogN)       | O(NlogN)       | O(NlogN)       | O(N)         | 稳定   | 需要稳定性，空间不是很重要                                   |
| 堆排序   | O(NlogN)       | O(NlogN)       | O(NlogN)       | O(1)         | O(1)   | 不稳定                                                       |
+----------+----------------+----------------+----------------+--------------+--------+--------------------------------------------------------------+

* 当规模较小，如小于等于50，采用插入或选择排序。
* 当元素基本有序，选择插入、冒泡或随机的快速排序。
* 当规模较大，采用 O(NlogN)排序算法。
* 当待排序的关键字随机分布时，快速排序的平均时间最短。
* 当需要保证稳定性的时候，选用归并排序。

## 非比较排序 ##

之前介绍的算法都是基于比较的排序算法，下边介绍两种不是基于比较的算法。

### 计数排序 ###

已知数据范围 x1 到 x2, 对范围中的元素进行排序。可以使用一个长度为 x2-x1+1 的数组，存储每个数字对应的出现的次数。最终得到排序后的结果。

### 桶排序 ###

桶排序假设待排序的一组数均匀独立的分布在一个范围中，并将这一范围划分成几个桶。然后基于某种映射函数，将待排序的关键字 k 映射到第 i 个桶中。接着将各个桶中的数据有序的合并起来，对每个桶中的元素可以进行排序，然后输出得到一个有序序列。