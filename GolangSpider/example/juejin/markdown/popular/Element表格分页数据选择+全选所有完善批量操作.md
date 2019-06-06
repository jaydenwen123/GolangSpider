# Element表格分页数据选择+全选所有完善批量操作 #

> 
> 后台管理系统中的列表页面，一般都会有对列表数据进行批量操作的功能，例如：批量删除、批量删除等。

之前项目中只是简单的用到Element框架中常规的属性、事件。在一次机缘巧合下，了解到一个公司内部的框架是基于Element框架内部实现了一些插件功能，对于表格这一块完善了很多功能，当时没有把握住机会去看源码是怎么实现的，现在有点小后悔呢，嘤嘤嘤~~~~没关系，自己慢慢一点一点实现。

实现的功能有：

* 分页数据选择
* 全选所有数据（不是element框架自带的全选本页哦！）

1、分页数据选择
一开始以为不就是分页的时候把之前选中的数据存储在一个list里面嘛，然后选择的时候map一下。等到自己写代码的时候，会发现没有那么简单，百度后，发现有两个属性被忽视了

* row-key
* reserve-selection ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7823be4f904?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b783ee166966?imageView2/0/w/1280/h/960/ignore-error/1)

代码截图：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7a3ea795b78?imageView2/0/w/1280/h/960/ignore-error/1)

事件代码：

` getRowKeys (row) { return row.execId } 复制代码`

这样通过 selectionChange 方法就能获取页面中选中数据的改变，将选中的数据保存到list中

` selectionChange (rows) { this.checkList = rows } 复制代码`

2、全选所有数据

> 
> element框架中有select-all事件，全选本页所有数据，但是项目中，经常会遇到说对所有的进行操作，例如批量删除（删除所有数据，这个权限有点大）
> 
> 

实现思路：

* 一个全选所有复选框，当选中时，前端传递一个参数Flag:1给后台，后台就会知道这是对所有数据进行操作，同时前后台之间都不用进行庞大的数据传输

` <el-checkbox v-model= "allCheck" @change= "allCheckEvent" >全选所有</el-checkbox> 复制代码`

* 选中全选所有复选框，当前页数据需全部是选中状态，翻页到另一页，这一页的数据也是全部选中状态 （监听当前页中数据）

` allCheckEvent () { if (this.allCheck) { this.testList.forEach(row => { this. $refs.recordTable.toggleRowSelection(row, true ) }) } else { this. $refs.recordTable.clearSelection() } } 复制代码`

` watch: { test List: { handler (value) { if (this.allCheck) { let that = this let len = that.checkList.length value.forEach(row => { for ( let i = 0; i < len; i++) { if (row.execId === that.checkList[i].execId) { that. $refs.recordTable.toggleRowSelection(row, false ) break } else { that. $refs.recordTable.toggleRowSelection(row, true ) } } }) } }, deep: true } } 复制代码`

* 选中全选所有复选框，同时，已经翻页了两页，选中的数据是两页数据，若取消其中一行数据的选中状态，此时，全选所有取消，当前选中的数据应是：已翻页的两页数据-取消的那一行数据

` selectOne () { if (this.allCheck) { this.allCheck = false } } 复制代码`

实现的表格：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2baf77c193d01?imageView2/0/w/1280/h/960/ignore-error/1)

走了不少弯路才注意到的问题：

* 若从第一页翻选到第二页，然后又回到第一页，选中的数据理应是1+2两页的数据，现实是1+2+1这三页数据，在展现形式上是看不出来问题，而且前面说了，全选所有的时候，我向后台传的参数只是一个flag，而不是这些选中数据。但是若在第一页取消一行数据，此时全选所有数据框已取消，本条数据也不是选中状态，翻到第二页，在回到第二页，Duang~那条数据又回到了选中状态！因为选中数据中该条数据是两条啊，你取消了一个，另一个还在呀，当然你再取消一次，再回来，是取消状态，bug，bug，bug！
* 想到的就是数据要去重，首先想到的是从结果去重，在selectionChange方法中去重，悲剧了，根本不起作用，理清思路后发现：当选择项发生改变时，调用selectionChange方法获取选中的所有数据，此时我们用forEach遍历数据，用toggleRowSelection方法将页面中的数据选中，此时toggleRowSelection一次，selectionChange方法执行一次
* 那就在监听数据时，如果数据ID相同，不在执行toggleRowSelection方法

最后说一句：当你累了，不开心了，去看RM吧，你会发现，世间还有这么美好的一群人，在开心的笑，努力的生活！

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bb47ca162264?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bb4c14356fdd?imageView2/0/w/1280/h/960/ignore-error/1)