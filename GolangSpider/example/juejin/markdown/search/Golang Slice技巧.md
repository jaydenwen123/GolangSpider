# Golang Slice技巧 #

### 追加元素 ###

` a = append(a, b...) 复制代码`

### 复制 ###

` b = make([]T, len(a)) copy(b,a) // or b = append([]T(nil), a...) // or b = append(a[:0:0], a...) 复制代码`

### 裁剪 ###

` a = append(a[:i], a[j:]...) 复制代码`

### 删除元素 ###

` a = append(a[:i], a[i+1:]...) // or a = a[:i+copy(a[i:], a[i+1:])] 复制代码`

### 删除而不保留顺序 ###

` a[i] = a[len(a)-1] a = a[:len(a)-1] 复制代码`
> 
> 
> 
> 如果元素的类型是指针或带指针字段的结构，需要对其进行垃圾收集，那么上面的Cut和Delete实现有潜在的内存泄漏问题:一些带值的元素仍然被片a引用，因此无法收集。下面的代码可以修复这个问题
> 
> 
> 

> 
> 
> 
> #### 裁剪 ####
> 
> 

` copy(a[i:], a[j:]) for k, n := len(a)-j+i, len(a); k < n; k++ { a[k] = nil // or the zero value of T } a = a[:len(a)-j+i] 复制代码`
> 
> 
> 
> 
> #### 删除 ####
> 
> 

` copy(a[i:], a[i+1:]) a[len(a)-1] = nil // or the zero value of T a = a[:len(a)-1] 复制代码`
> 
> 
> 
> 
> #### 删除不保留顺序 ####
> 
> 

` a[i] = a[len(a)-1] a[len(a)-1] = nil a = a[:len(a)-1] 复制代码`

### 扩大 ###

` a = append(a[:i], append(make([]T, j), a[i:]...)...) 复制代码`

### 延伸 ###

` a = append(a, make([]T, j)...) 复制代码`

### 插入 ###

` a = append(a[:i], append([]T{x}, a[i:]...)...) 复制代码`
> 
> 
> 
> 第二个append创建一个具有自己底层存储的新的切片，并将a[i:]中的元素复制到该切片，然后将这些元素复制回切片a(通过第一个append)。可以使用另一种方法避免创建新的片(以及内存垃圾)和第二个副本
> 
> 
> 

> 
> 
> 
> #### 插入 ####
> 
> 

` s = append(s, 0 /* use the zero value of the element type */) copy(s[i+1:], s[i:]) s[i] = x 复制代码`

### 插入Slice ###

` a = append(a[:i], append(b, a[i:]...)...) 复制代码`

### 增加元素 ###

` a = append(a, x) 复制代码`

### 取出最后一个元素 ###

` x, a = a[len(a)-1], a[:len(a)-1] 复制代码`

### Push Front/Unshift ###

` a = append([]T{x}, a...) 复制代码`

### Pop Front/Shift ###

` x, a = a[0], a[1:] 复制代码`

### 其他技巧 ###

#### 不进行内存分配过滤 ####

让一个切片与原始切片共享底层数组即可

` b := a[:0] for _, x := range a { if f(x) { b = append(b, x) } } 复制代码`

对于必须进行垃圾回收的元素，后面可以包含以下代码:

` for i := len(b); i < len(a); i++ { a[i] = nil // or the zero value of T } 复制代码`

#### 反转 ####

用相同但顺序相反的元素替换切片的内容:

` for i := len(a)/2-1; i >= 0; i-- { opp := len(a)-1-i a[i], a[opp] = a[opp], a[i] } 复制代码` ` for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 { a[left], a[right] = a[right], a[left] } 复制代码`

#### 打乱顺序 ####

` for i := len(a) - 1; i > 0; i-- { j := rand.Intn(i + 1) a[i], a[j] = a[j], a[i] } 复制代码` ` words := strings.Fields( "ink runs from the corners of my mouth" ) rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] }) fmt.Println(words) 复制代码`

#### 批处理 ####

` actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} batchSize := 3 var batches [][]int for batchSize < len(actions) { actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize]) } batches = append(batches, actions) output: [[0 1 2] [3 4 5] [6 7 8] [9]] 复制代码`

#### 去重 ####

` import "sort" in := []int{3,2,1,4,3,2,1,4,1} // any item can be sorted sort.Ints( in ) j := 0 for i := 1; i < len( in ); i++ { if in [j] == in [i] { continue } j++ // preserve the original data // in [i], in [j] = in [j], in [i] // only set what is required in [j] = in [i] } result := in [:j+1] fmt.Println(result) // [1 2 3 4] 复制代码`

翻译自： [github.com/golang/go/w…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgolang%2Fgo%2Fwiki%2FSliceTricks%23filtering-without-allocating )