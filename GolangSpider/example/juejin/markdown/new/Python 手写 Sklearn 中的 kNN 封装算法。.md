# Python 手写 Sklearn 中的 kNN 封装算法。 #

昨天通过一个酒吧猜红酒的故事，介绍了机器学习中最简单的一个算法：kNN （K 近邻算法），并用 Python 一步步实现这个算法。同时为了对比，调用了 Sklearn 中的 kNN 算法包，仅用了 5 行代码。两种方法殊途同归，都正确解决了二分类问题，即新倒的红酒属于赤霞珠。

文章传送门：

[Python 手写机器学习最简单的 kNN 算法]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzA5NDk4NDcwMw%3D%3D%26amp%3Bmid%3D2651387573%26amp%3Bidx%3D1%26amp%3Bsn%3D6b5ddee9cb9d29a0b7eb748d3a78529a%26amp%3Bchksm%3D8bba1625bccd9f330b2e2c86f684036782e600280a29bc356a02016bf54eeba009475743dfc4%26amp%3Btoken%3D2032808936%26amp%3Blang%3Dzh_CN%23rd ) （可点击）

虽然调用 Sklearn 库算法，简单的几行代码就能解决问题，感觉很爽，但其实我们时处于黑箱中的，Sklearn 背后干了些什么我们其实不明白。作为初学者，如果不搞清楚算法原理就直接调包，学的也只是表面功夫，没什么卵用。

所以今天来我们了解一下 Sklearn 是如何封装 kNN 算法的并自己 Python 实现一下。这样，以后我们再调用 Sklearn 算法包时，会有更清晰的认识。

先来回顾昨天 Sklearn 中 kNN 算法的 5 行代码：

` from sklearn.neighbors import KNeighborsClassifier kNN_classifier = KNeighborsClassifier(n_neighbors= 3 ) kNN_classifier.fit(X_train,y_train ) x_test = x_test.reshape( 1 , -1 ) kNN_classifier.predict(x_test)[ 0 ] 复制代码`

代码已解释过，今天用一张图继续加深理解：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2719f3afc3716?imageView2/0/w/1280/h/960/ignore-error/1)

可以说，Sklearn 调用所有的机器学习算法几乎都是按照这样的套路：把训练数据喂给选择的算法进行 fit 拟合，能计算出一个模型，模型有了就把要预测的数据喂给模型，进行预测 predict，最后输出结果，分类和回归算法都是如此。

值得注意的一点是，**kNN 是一个特殊算法，它不需要训练（fit）建立模型，直接拿测试数据在训练集上就可以预测出结果。**这也是为什么说 kNN 算法是最简单的机器学习算法原因之一。

但在上面的 Sklearn 中为什么这里还 fit 拟合这一步操作呢，实际上是可以不用的，不过 Sklearn 的接口很整齐统一，所以为了跟多数算法保持一致把训练集当成模型。

随着之后我们学习更多的算法，会发现每个算法都有一些特点，可以总结对比一下。

把昨天的手写代码整理成一个函数就可以看到没有训练过程：

` import numpy as np from math import sqrt from collections import Counter def kNNClassify (K, X_train, y_train, X_predict) : distances = [sqrt(np.sum((x - X_predict)** 2 )) for x in X_train] sort = np.argsort(distances) topK = [y_train[i] for i in sort[:K]] votes = Counter(topK) y_predict = votes.most_common( 1 )[ 0 ][ 0 ] return y_predict 复制代码`

接下来我们按照上图的思路，把 Sklearn 中封装的 kNN 算法，从底层一步步写出那 5 行代码是如何运行的：

` import numpy as np from math import sqrt from collections import Counter class kNNClassifier : def __init__ (self,k) : self.k =k self._X_train = None self._y_train = None def fit (self,X_train,y_train) : self._X_train = X_train self._y_train = y_train return self 复制代码`

首先，我们需要把之前的函数改写一个名为 kNNClassifier 的 Class 类，因为 Sklearn 中的算法都是面向对象的，使用类更方便。

如果你对类还不熟悉可以参考我以前的一篇文章：

[Python 类 Class 的理解]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzA5NDk4NDcwMw%3D%3D%26amp%3Btempkey%3DMTAxMl9DM1JVakN6MnRPbitNMVRjdk5ELWZyZ1ZfYllZcGNCX0pJSjlGcmJyOVdMSzhEY3ZscmtHYnFDRVdkemsyXzFOeG82MlRZak51ZHZ4YmJ1UkpOdkVHZHQteWVJU2hhYVJyVzJFeEdxSEs3ekRJSW9LNlh3dFg5M2dnSmtjMmY1MGhkOXMwNTd4Zm40LVNWOWxmSlhmMVVBRkNZc1VwSVNrOEpFd1dRfn4%253D%26amp%3Bchksm%3D0bba104d3ccd995bd56d90ffd9b77c7f42745bfdc24b2a769cb1862c6140f4774d9fe2f3b897%23rd ) （可点击）

在 ` __init__` 函数中定义三个初始变量，k 表示我们要选择传进了的 k 个近邻点。

` self._X_train` 和 ` self._y_train` 前面有个 下划线_ ，意思是把它们当成内部私有变量，只在内部运算，外部不能改动。

接着定义一个 fit 函数，这个函数就是用来拟合 kNN 模型，但 kNN 模型并不需要拟合，所以我们就原封不动地把数据集复制一遍，最后返回两个数据集自身。

这里要对输入的变量做一下约束，一个是 X_train 和 y_train 的行数要一样，一个是我们选的 k 近邻点不能是非法数，比如负数或者多于样本点的数， 不然后续计算会出错。用什么做约束呢，可以使用 assert 断言语句：

` def fit (self,X_train,y_train) : assert X_train.shape[ 0 ] == y_train.shape[ 0 ], "添加 assert 断言是为了确保输入正常的数据集和k值，如果不添加一旦输入不正常的值，难找到出错原因" assert self.k <= X_train.shape[ 0 ] self._X_train = X_train self._y_train = y_train return self 复制代码`

接下来我们就要传进待预测的样本点，计算它跟每个样本点之间的距离，对应 Sklearn 中的 predict ，这是算法的核心部分。而这一步代码就是我们之前写的函数，可以直接拿过来用，加几行断言保证输入的变量是合理的。

` def predict (self,X_predict) : assert self._X_train is not None , "要求predict 之前要先运行 fit 这样self._X_train 就不会为空" assert self._y_train is not None assert X_predict.shape[ 1 ] == self._X_train.shape[ 1 ], "要求测试集和预测集的特征数量一致" distances = [sqrt(np.sum((x_train - X_predict)** 2 )) for x_train in self._X_train] sort = np.argsort(distances) topK = [self._y_train[i] for i in sort[:self.k]] votes = Counter(topK) y_predict = votes.most_common( 1 )[ 0 ][ 0 ] return y_predict 复制代码`

到这儿我们就完成了一个简易的 Sklearn kNN 封装算法，保存为 ` kNN_sklearn.py` 文件，然后在 jupyter notebook 运行测试一下：

先获得基础数据：

` # 样本集 X_raw = [[ 13.23 , 5.64 ], [ 13.2 , 4.38 ], [ 13.16 , 4.68 ], [ 13.37 , 4.8 ], [ 13.24 , 4.32 ], [ 12.07 , 2.76 ], [ 12.43 , 3.94 ], [ 11.79 , 3. ], [ 12.37 , 2.12 ], [ 12.04 , 2.6 ]] X_train = np.array(X_raw) # 特征值 y_raw = [ 0 , 0 , 0 , 0 , 0 , 1 , 1 , 1 , 1 , 1 ] y_train = np.array(y_raw) # 待预测值 x_test= np.array([ 12.08 , 3.3 ]) X_predict = x_test.reshape( 1 , -1 ) 复制代码`

**注意：当预测变量只有一个时，一定要 reshape(1,-1) 成二维数组不然会报错。**

在 jupyter notebook 中运行程序可以使用一个魔法命令 %run：

` %run kNN_Euler.py 复制代码`

这样就直接运行好了 kNN_Euler.py 程序，然后就可以调用程序中的 kNNClassifier 类，赋予 k 参数为 3，命名为一个实例 kNN_classify 。

` kNN_classify = kNNClassifier( 3 ) 复制代码`

接着把样本集 X_train，y_train 传给实例 fit ：

` kNN_classify.fit(X_train,y_train) 复制代码`

fit 好后再传入待预测样本 X_predict 进行预测就可以得到分类结果了：

` y_predict = kNN_classify.predict(X_predict) y_predict [out]: 1 复制代码`

答案是 1 和昨天两种方法的结果是一样的。

是不是不难？

再进一步，如果我们一次预测不只一个点，而是多个点，比如要预测下面这两个点属于哪一类：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2719f3b07e549?imageView2/0/w/1280/h/960/ignore-error/1)

那能不能同时给出预测的分类结果呢？答案当然是可以的，我们只需要稍微修改以下上面的封装算法就可以了，把 predict 函数作如下修改：

` def predict (self,X_predict) : y_predict = [self._predict(x) for x in X_predict] # 列表生成是把分类结果都存储到list 中然后返回 return np.array(y_predict) def _predict (self,x) : # _predict私有函数 assert self._X_train is not None assert self._y_train is not None distances = [sqrt(np.sum((x_train - x)** 2 )) for x_train in self._X_train] sort = np.argsort(distances) topK = [self._y_train[i] for i in sort[:self.k]] votes = Counter(topK) y_predict = votes.most_common( 1 )[ 0 ][ 0 ] return y_predict 复制代码`

这里定义了两个函数，predict 用列表生成式来存储多个预测分类值，预测值从哪里来呢，就是利用 ` _predict` 函数计算， ` _predict` 前面的下划线同样表明它是封装的私有函数，只在内部使用，外界不能调用，因为不需要。

算法写好，只需要传入多个预测样本就可以了，这里我们传递两个：

` X_predict = np.array([[ 12.08 , 3.3 ], [ 12.8 , 4.1 ]]) 复制代码`

输出预测结果：

` y_predict = kNN_classify.predict(X_predict) y_predict [out]：array([ 1 , 0 ]) 复制代码`

看，返回了两个值，第一个样本的分类结果是 1 即赤霞珠，第二个样本结果是 0 即黑皮诺。和实际结果一致，很完美。

到这里，我们就按照 Sklearn 算法封装方式写出了 kNN 算法，不过 Sklearn 中的 kNN 算法要比这复杂地多，因为 kNN 算法还有很多要考虑的，比如处理 **kNN 算法的一个缺点：计算耗时。**简单说就是 kNN 算法运行时间高度依赖样本集有和特征值数量的维度，当维度很高时算法运行时间就极速增加，具体原因和改善方法我们后续再说。

现在还有一个重要的问题，我们在全部训练集上实现了 kNN 算法，但它预测的效果和准确率怎么样，我们并不清楚，下一篇文章，来说说怎么衡量 kNN 算法效果的好坏。

本文的 jupyter notebook 代码，可以在我公众号：「 **高级农民工** 」后台回复「 **kNN2** 」得到，加油！

![](https://user-gold-cdn.xitu.io/2019/6/5/16b271b08384fc9a?imageView2/0/w/1280/h/960/ignore-error/1)