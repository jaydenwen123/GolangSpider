# [译] JavaScript 线性代数：向量 #

> 
> 
> 
> * 原文地址： [Linear Algebra: Vectors](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40geekrodion%2Flinear-algebra-vectors-f7610e9a0f23
> )
> * 原文作者： [Rodion Chachura](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40geekrodion )
> * 译文出自： [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> )
> * 本文永久链接： [github.com/xitu/gold-m…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%2Fblob%2Fmaster%2FTODO1%2Flinear-algebra-vectors.md
> )
> * 译者： [lsvih](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flsvih )
> * 校对者： [Endone](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FEndone )
> 
> 
> 

本文是“ [JavaScript 线性代数]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40geekrodion%2Flinear-algebra-with-javascript-46c289178c0 ) ”教程的一部分。

**向量** 是用于精确表示空间中方向的方法。向量由一系列数值构成，每维数值都是向量的一个 **分量** 。在下图中，你可以看到一个由两个分量组成的、在 2 维空间内的向量。在 3 维空间内，向量会由 3 个分量组成。

![the vector in 2D space](https://user-gold-cdn.xitu.io/2019/6/4/16b215d286a3c481?imageView2/0/w/1280/h/960/ignore-error/1)

我们可以为 2 维空间的向量创建一个 **Vector2D** 类，然后为 3 维空间的向量创建一个 **Vector3D** 类。但是这么做有一个问题：向量并不仅用于表示物理空间中的方向。比如，我们可能需要将颜色（RGBA）表示为向量，那么它会有 4 个分量：红色、绿色、蓝色和 alpha 通道。或者，我们要用向量来表示有不同占比的 **n** 种选择（比如表示 5 匹马赛马，每匹马赢得比赛的概率的向量）。因此，我们会创建一个不指定维度的类，并像这样使用它：

` class Vector { constructor (...components) { this.components = components } } const direction2d = new Vector( 1 , 2 ) const direction3d = new Vector( 1 , 2 , 3 ) const color = new Vector( 0.5 , 0.4 , 0.7 , 0.15 ) const probabilities = new Vector( 0.1 , 0.3 , 0.15 , 0.25 , 0.2 ) 复制代码`

## 向量运算 ##

考虑有两个向量的情况，可以对它们定义以下运算：

![basic vector operations](https://user-gold-cdn.xitu.io/2019/6/4/16b215d28b284a80?imageView2/0/w/1280/h/960/ignore-error/1)

其中， **α ∈ R** 为任意常数。

我们对除了叉积之外的运算进行了可视化，你可以在 [此处]( https://link.juejin.im?target=https%3A%2F%2Frodionchachura.github.io%2Flinear-algebra%2F ) 找到相关示例。 [此 GitHub 仓库]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FRodionChachura%2Flinear-algebra ) 里有用来创建这些可视化示例的 React 项目和相关的库。如果你想知道如何使用 React 和 SVG 来制作这些二维可视化示例，请参考 [本文]( https://juejin.im/post/5cefbc37f265da1bd260d129 ) 。

### 加法与减法 ###

与数值运算类似，你可以对向量进行加法与减法运算。对向量进行算术运算时，可以直接对向量各自的分量进行数值运算得到结果：

![vectors addition](https://user-gold-cdn.xitu.io/2019/6/4/16b215d28a82f9e2?imageView2/0/w/1280/h/960/ignore-error/1)

![vectors subtraction](https://user-gold-cdn.xitu.io/2019/6/4/16b215d2914afdef?imageView2/0/w/1280/h/960/ignore-error/1)

加法函数接收另一个向量作为参数，并将对应的向量分量相加，返回得出的新向量。减法函数与之类似，不过会将加法换成减法：

` class Vector { constructor (...components) { this.components = components } add({ components }) { return new Vector( ...components.map( ( component, index ) => this.components[index] + component) ) } subtract({ components }) { return new Vector( ...components.map( ( component, index ) => this.components[index] - component) ) } } const one = new Vector( 2 , 3 ) const other = new Vector( 2 , 1 ) console.log(one.add(other)) // Vector { components: [ 4, 4 ] } console.log(one.subtract(other)) // Vector { components: [ 0, 2 ] } 复制代码`

## 缩放 ##

我们可以对一个向量进行缩放，缩放比例可为任意数值 **α ∈ R** 。缩放时，对所有向量分量都乘以缩放因子 **α** 。当 **α > 1** 时，向量会变得更长；当 **0 ≤ α < 1** 时，向量会变得更短。如果 **α** 是负数，缩放后的向量将会指向原向量的反方向。

![scaling vector](https://user-gold-cdn.xitu.io/2019/6/4/16b215d28bbe7766?imageView2/0/w/1280/h/960/ignore-error/1)

在 **scaleBy** 方法中，我们对所有的向量分量都乘上传入参数的数值，得到新的向量并返回：

` class Vector { constructor (...components) { this.components = components } // ... scaleBy(number) { return new Vector( ...this.components.map( component => component * number) ) } } const vector = new Vector( 1 , 2 ) console.log(vector.scaleBy( 2 )) // Vector { components: [ 2, 4 ] } console.log(vector.scaleBy( 0.5 )) // Vector { components: [ 0.5, 1 ] } console.log(vector.scaleBy( -1 )) // Vector { components: [ -1, -2 ] } 复制代码`

## 长度 ##

向量长度可由勾股定理导出：

![vectors length](https://user-gold-cdn.xitu.io/2019/6/4/16b215d34a753889?imageView2/0/w/1280/h/960/ignore-error/1)

由于在 JavaScript 内置的 Math 对象中有现成的函数，因此计算长度的方法非常简单：

` class Vector { constructor (...components) { this.components = components } // ... length() { return Math.hypot(...this.components) } } const vector = new Vector( 2 , 3 ) console.log(vector.length()) // 3.6055512754639896 复制代码`

## 点积 ##

点积可以计算出两个向量的相似程度。点积方法接收两个向量作为输入，并输出一个数值。两个向量的点积等于它们各自对应分量的乘积之和。

![dot product](https://user-gold-cdn.xitu.io/2019/6/4/16b215d354c267ec?imageView2/0/w/1280/h/960/ignore-error/1)

在 **dotProduct** 方法中，接收另一个向量作为参数，通过 reduce 方法来计算对应分量的乘积之和：

` class Vector { constructor (...components) { this.components = components } // ... dotProduct({ components }) { return components.reduce( ( acc, component, index ) => acc + component * this.components[index], 0 ) } } const one = new Vector( 1 , 4 ) const other = new Vector( 2 , 2 ) console.log(one.dotProduct(other)) // 10 复制代码`

在我们观察几个向量间的方向关系前，需要先实现一种将向量长度归一化为 1 的方法。这种归一化后的向量在许多情景中都会用到。比如说当我们需要在空间中指定一个方向时，就需要用一个归一化后的向量来表示这个方向。

` class Vector { constructor (...components) { this.components = components } // ... normalize() { return this.scaleBy( 1 / this.length()) } } const vector = new Vector( 2 , 4 ) const normalized = vector.normalize() console.log(normalized) // Vector { components: [ 0.4472135954999579, 0.8944271909999159 ] } console.log(normalized.length()) // 1 复制代码`

![using dot product](https://user-gold-cdn.xitu.io/2019/6/4/16b215d293a1ca15?imageView2/0/w/1280/h/960/ignore-error/1)

如果两个归一化后的向量的点积结果等于 1，则意味着这两个向量的方向相同。我们创建了 **areEqual** 函数用来比较两个浮点数：

` const EPSILON = 0.00000001 const areEqual = ( one, other, epsilon = EPSILON ) => Math.abs(one - other) < epsilon class Vector { constructor (...components) { this.components = components } // ... haveSameDirectionWith(other) { const dotProduct = this.normalize().dotProduct(other.normalize()) return areEqual(dotProduct, 1 ) } } const one = new Vector( 2 , 4 ) const other = new Vector( 4 , 8 ) console.log(one.haveSameDirectionWith(other)) // true 复制代码`

如果两个归一化后的向量点积结果等于 -1，则表示它们的方向完全相反：

` class Vector { constructor (...components) { this.components = components } // ... haveOppositeDirectionTo(other) { const dotProduct = this.normalize().dotProduct(other.normalize()) return areEqual(dotProduct, -1 ) } } const one = new Vector( 2 , 4 ) const other = new Vector( -4 , -8 ) console.log(one.haveOppositeDirectionTo(other)) // true 复制代码`

如果两个归一化后的向量的点积结果为 0，则表示这两个向量是相互垂直的：

` class Vector { constructor (...components) { this.components = components } // ... isPerpendicularTo(other) { const dotProduct = this.normalize().dotProduct(other.normalize()) return areEqual(dotProduct, 0 ) } } const one = new Vector( -2 , 2 ) const other = new Vector( 2 , 2 ) console.log(one.isPerpendicularTo(other)) // true 复制代码`

## 叉积 ##

叉积仅对三维向量适用，它会产生垂直于两个输入向量的向量：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b215d354b69fb9?imageView2/0/w/1280/h/960/ignore-error/1)

我们实现叉积时，假定它只用于计算三维空间内的向量。

` class Vector { constructor (...components) { this.components = components } // ... // 只适用于 3 维向量 crossProduct({ components }) { return new Vector( this.components[ 1 ] * components[ 2 ] - this.components[ 2 ] * components[ 1 ], this.components[ 2 ] * components[ 0 ] - this.components[ 0 ] * components[ 2 ], this.components[ 0 ] * components[ 1 ] - this.components[ 1 ] * components[ 0 ] ) } } const one = new Vector( 2 , 1 , 1 ) const other = new Vector( 1 , 2 , 2 ) console.log(one.crossProduct(other)) // Vector { components: [ 0, -3, 3 ] } console.log(other.crossProduct(one)) // Vector { components: [ 0, 3, -3 ] } 复制代码`

## 其它常用方法 ##

在现实生活的应用中，上述方法是远远不够的。比如说，我们有时需要找到两个向量的夹角、将一个向量反向，或者计算一个向量在另一个向量上的投影等。

在开始编写上面说的方法前，需要先写下面两个函数，用于在角度与弧度间相互转换：

` const toDegrees = radians => (radians * 180 ) / Math.PI const toRadians = degrees => (degrees * Math.PI) / 180 复制代码`

### 夹角 ###

` class Vector { constructor (...components) { this.components = components } // ... angleBetween(other) { return toDegrees( Math.acos( this.dotProduct(other) / ( this.length() * other.length()) ) ) } } const one = new Vector( 0 , 4 ) const other = new Vector( 4 , 4 ) console.log(one.angleBetween(other)) // 45.00000000000001 复制代码`

### 反向 ###

当需要将一个向量的方向指向反向时，我们可以对这个向量进行 -1 缩放：

` class Vector { constructor (...components) { this.components = components } // ... negate() { return this.scaleBy( -1 ) } } const vector = new Vector( 2 , 2 ) console.log(vector.negate()) // Vector { components: [ -2, -2 ] } 复制代码`

## 投影 ##

![project v on d](https://user-gold-cdn.xitu.io/2019/6/4/16b215d36d1d1f5d?imageView2/0/w/1280/h/960/ignore-error/1)

` class Vector { constructor (...components) { this.components = components } // ... projectOn(other) { const normalized = other.normalize() return normalized.scaleBy( this.dotProduct(normalized)) } } const one = new Vector( 8 , 4 ) const other = new Vector( 4 , 7 ) console.log(other.projectOn(one)) // Vector { components: [ 6, 3 ] } 复制代码`

### 设定长度 ###

当需要给向量指定一个长度时，可以使用如下方法：

` class Vector { constructor (...components) { this.components = components } // ... withLength(newLength) { return this.normalize().scaleBy(newLength) } } const one = new Vector( 2 , 3 ) console.log(one.length()) // 3.6055512754639896 const modified = one.withLength( 10 ) // 10 console.log(modified.length()) 复制代码`

### 判断相等 ###

为了判断两个向量是否相等，可以对它们对应的分量使用 **areEqual** 函数：

` class Vector { constructor (...components) { this.components = components } // ... equalTo({ components }) { return components.every( ( component, index ) => areEqual(component, this.components[index])) } } const one = new Vector( 1 , 2 ) const other = new Vector( 1 , 2 ) console.log(one.equalTo(other)) // true const another = new Vector( 2 , 1 ) console.log(one.equalTo(another)) // false 复制代码`

## 单位向量与基底 ##

我们可以将一个向量看做是“在 x 轴上走 ![v_x](https://juejin.im/equation?tex=v_x) 的距离、在 y 轴上走 ![v_y](https://juejin.im/equation?tex=v_y) 的距离、在 z 轴上走 ![v_z](https://juejin.im/equation?tex=v_z) 的距离”。我们可以使用 ![\hat { \imath }](https://juejin.im/equation?tex=%5Chat%20%7B%20%5Cimath%20%7D) 、 ![\hat { \jmath }](https://juejin.im/equation?tex=%5Chat%20%7B%20%5Cjmath%20%7D) 和 ![\hat { k }](https://juejin.im/equation?tex=%5Chat%20%7B%20k%20%7D) 分别乘上一个值更清晰地表示上述内容。下图分别是 ![x](https://juejin.im/equation?tex=x) 、 ![y](https://juejin.im/equation?tex=y) 、 ![z](https://juejin.im/equation?tex=z) 轴上的 **单位向量** ：

![\hat { \imath } = ( 1,0,0 ) \quad \hat { \jmath } = ( 0,1,0 ) \quad \hat { k } = ( 0,0,1 )](https://juejin.im/equation?tex=%5Chat%20%7B%20%5Cimath%20%7D%20%3D%20(%201%2C0%2C0%20)%20%5Cquad%20%5Chat%20%7B%20%5Cjmath%20%7D%20%3D%20(%200%2C1%2C0%20)%20%5Cquad%20%5Chat%20%7B%20k%20%7D%20%3D%20(%200%2C0%2C1%20))

任何数值乘以 ![\hat { \imath }](https://juejin.im/equation?tex=%5Chat%20%7B%20%5Cimath%20%7D) 向量，都可以得到一个第一维分量等于该数值的向量。例如：

![2 \hat { \imath } = ( 2,0,0 ) \quad 3 \hat { \jmath } = ( 0,3,0 ) \quad 5 \hat { K } = ( 0,0,5 )](https://juejin.im/equation?tex=2%20%5Chat%20%7B%20%5Cimath%20%7D%20%3D%20(%202%2C0%2C0%20)%20%5Cquad%203%20%5Chat%20%7B%20%5Cjmath%20%7D%20%3D%20(%200%2C3%2C0%20)%20%5Cquad%205%20%5Chat%20%7B%20K%20%7D%20%3D%20(%200%2C0%2C5%20))

向量中最重要的一个概念是 **基底** 。设有一个 3 维向量 ![\mathbb{R}^3](https://juejin.im/equation?tex=%5Cmathbb%7BR%7D%5E3) ，它的基底是一组向量： ![\{\hat{e}_1,\hat{e}_2,\hat{e}_3\}](https://juejin.im/equation?tex=%5C%7B%5Chat%7Be%7D_1%2C%5Chat%7Be%7D_2%2C%5Chat%7Be%7D_3%5C%7D) ，这组向量也可以作为 ![\mathbb{R}^3](https://juejin.im/equation?tex=%5Cmathbb%7BR%7D%5E3) 的坐标系统。如果 ![\{\hat{e}_1,\hat{e}_2,\hat{e}_3\}](https://juejin.im/equation?tex=%5C%7B%5Chat%7Be%7D_1%2C%5Chat%7Be%7D_2%2C%5Chat%7Be%7D_3%5C%7D) 是一组基底，则可以将任何向量 ![\vec{v} \in \mathbb{R}^3](https://juejin.im/equation?tex=%5Cvec%7Bv%7D%20%5Cin%20%5Cmathbb%7BR%7D%5E3) 表示为该基底的系数 ![(v_1,v_2,v_3)](https://juejin.im/equation?tex=(v_1%2Cv_2%2Cv_3)) ：

![\vec{v} = v_1 \hat{e}_1 + v_2 \hat{e}_2 + v_3 \hat{e}_3](https://juejin.im/equation?tex=%5Cvec%7Bv%7D%20%3D%20v_1%20%5Chat%7Be%7D_1%20%2B%20v_2%20%5Chat%7Be%7D_2%20%2B%20v_3%20%5Chat%7Be%7D_3)

向量 ![\vec{v}](https://juejin.im/equation?tex=%5Cvec%7Bv%7D) 是通过在 ![\hat{e}_1](https://juejin.im/equation?tex=%5Chat%7Be%7D_1) 方向上测量 ![v_2](https://juejin.im/equation?tex=v_2) 的距离、在 ![\hat{e}_2](https://juejin.im/equation?tex=%5Chat%7Be%7D_2) 方向上测量 ![v_1](https://juejin.im/equation?tex=v_1) 的距离、在 ![\hat{e}_3](https://juejin.im/equation?tex=%5Chat%7Be%7D_3) 方向上测量 ![v_3](https://juejin.im/equation?tex=v_3) 的距离得出的。

在不知道一个向量的基底前，向量的系数三元组并没有什么意义。只有知道向量的基底，才能将类似于 ![(a,b,c)](https://juejin.im/equation?tex=(a%2Cb%2Cc)) 三元组的数学对象转化为现实世界中的概念（比如颜色、概率、位置等）。

> 
> 
> 
> 如果发现译文存在错误或其他需要改进的地方，欢迎到 [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 对译文进行修改并 PR，也可获得相应奖励积分。文章开头的 **本文永久链接** 即为本文在 GitHub 上的 MarkDown 链接。
> 
> 

> 
> 
> 
> [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 是一个翻译优质互联网技术文章的社区，文章来源为 [掘金]( https://juejin.im ) 上的英文分享文章。内容覆盖 [Android](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23android
> ) 、 [iOS](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23ios
> ) 、 [前端](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E5%2589%258D%25E7%25AB%25AF
> ) 、 [后端](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E5%2590%258E%25E7%25AB%25AF
> ) 、 [区块链](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E5%258C%25BA%25E5%259D%2597%25E9%2593%25BE
> ) 、 [产品](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E4%25BA%25A7%25E5%2593%2581
> ) 、 [设计](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E8%25AE%25BE%25E8%25AE%25A1
> ) 、 [人工智能](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E4%25BA%25BA%25E5%25B7%25A5%25E6%2599%25BA%25E8%2583%25BD
> ) 等领域，想要查看更多优质译文请持续关注 [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 、 [官方微博](
> https://link.juejin.im?target=http%3A%2F%2Fweibo.com%2Fjuejinfanyi ) 、 [知乎专栏](
> https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fjuejinfanyi
> ) 。
> 
>