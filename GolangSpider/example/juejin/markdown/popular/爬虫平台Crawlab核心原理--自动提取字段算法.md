# 爬虫平台Crawlab核心原理--自动提取字段算法 #

## 背景 ##

实际的大型爬虫开发项目中，爬虫工程师会被要求抓取监控几十上百个网站。一般来说这些网站的结构大同小异，不同的主要是被抓取项的提取规则。传统方式是让爬虫工程师写一个通用框架，然后将各网站的提取规则做成可配置的，然后将配置工作交给更初级的工程师或外包出去。这样做将爬虫开发流水线化，提高了部分生产效率。但是，配置的工作还是一个苦力活儿，还是非常消耗人力。因此， **自动提取字段** 应运而生。

自动提取字段是Crawlab在 [版本v0.2.2]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftikazyq%2Fcrawlab%2Freleases%2Ftag%2Fv0.2.2 ) 中在 [可配置爬虫]( https://juejin.im/post/5ceb4342f265da1bc8540660 ) 基础上开发的新功能。它让用户不用做任何繁琐的提取规则配置，就可以自动提取出可能的要抓取的列表项，做到真正的“一键抓取”，顺利的话，开发一个网站的爬虫可以半分钟内完成。市面上有利用机器学习的方法来实现自动抓取要提取的抓取规则，有一些可以做到精准提取，但遗憾的是平台要收取高额的费用，个人开发者或小型公司一般承担不起。

Crawlab的自动提取字段是根据人为抓取的模式来模拟的，因此不用经过任何训练就可以使用。而且，Crawlab的自动提取字段功能不会向用户收取费用，因为Crawlab本身就是免费的。

## 算法介绍 ##

算法的核心来自于人的行为本身，通过查找网页中看起来像列表的元素来定位列表及抓取项。一般我们查找列表项是怎样的一个过程呢？有人说：这还不容易吗，一看就知道那个是各列表呀！兄弟，拜托... 咱们是在程序的角度谈这个的，它只理解HTML、CSS、JS这些代码，并不像你那样智能。

我们识别一个列表，首先要看它是不是有很多类似的子项；其次，这些列表通常来说看起来比较“复杂”，含有很多看得见的元素；最后，我们还要关注分页，分页按钮一般叫做“下一页”、“下页”、“Next”、“Next Page”等等。

用程序可以理解的语言，我们把以上规则总结如下：

**列表项**

* 从根节点自上而下遍历标签；
* 对于每一个标签，如果包含多个同样的子标签，判断为列表标签候选；
* 取子标签（递归）个数最多的列表标签候选为列表标签；

**列表子项**

* 对以上规则提取的列表标签，对每个子标签（递归）进行遍历
* 将有href的a标签为加入目标字段；
* 将有text的标签为加入目标字段。

**分页**

* 对于每一个标签，如果标签文本为特定文本（“下一页”、“下页”、“next page”、“next”），选取该标签为目标标签。

这样，我们就设计好了自动提取列表项、列表子项、分页的规则。剩下的就是写代码了。我知道这样的设计过于简单，也过于理想，没有考虑到一些特殊情况。后面我们将通过在一些知名网站上测试看看我们的算法表现如何。

## 算法实现 ##

算法实现很简单。为了更好的操作HTML标签，我们选择了 ` lxml` 库作为HTML的操作库。 ` lxml` 是python的一个解析库，支持HTML和XML的解析，支持XPath、CSS解析方式，而且解析效率非常高。

自上而下的遍历语法是 ` sel.iter()` 。 ` sel` 是 ` etree.Element` ，而 ` iter` 会从根节点自上而下遍历各个元素，直到遍历完所有元素。它是一个 ` generator` 。

#### 构造解析树 ####

在获取到页面的HTML之后，我们需要调用 ` lxml` 中的 ` etree.HTML` 方法构造解析树。代码很简单如下，其中 ` r` 为 ` requests.get` 的 ` Response`

` # get html parse tree sel = etree.HTML(r.content) 复制代码`

这段带代码在 ` SpiderApi._get_html` 方法里。源码请见 [这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftikazyq%2Fcrawlab%2Fblob%2Fmaster%2Fcrawlab%2Froutes%2Fspiders.py ) 。

#### 辅助函数 ####

在开始构建算法之前，我们需要实现一些辅助函数。所有函数是封装在 ` SpiderApi` 类中的，所以写法与类方法一样。

` @staticmethod def _get_children (sel) : # 获取所有不包含comments的子节点 return [tag for tag in sel.getchildren() if type(tag) != etree._Comment] 复制代码` ` @staticmethod def _get_text_child_tags (sel) : # 递归获取所有文本子节点（根节点） tags = [] for tag in sel.iter(): if type(tag) != etree._Comment and tag.text is not None and tag.text.strip() != '' : tags.append(tag) return tags 复制代码` ` @staticmethod def _get_a_child_tags (sel) : # 递归获取所有超链接子节点（根节点） tags = [] for tag in sel.iter(): if tag.tag == 'a' : if tag.get( 'href' ) is not None and not tag.get( 'href' ).startswith( '#' ) and not tag.get( 'href' ).startswith( 'javascript' ): tags.append(tag) return tags 复制代码`

#### 获取列表项 ####

下面是核心中的核心！同学们请集中注意力。

我们来编写获取列表项的代码。以下是获得列表标签候选列表 ` list_tag_list` 的代码。看起来稍稍有些复杂，但其实逻辑很简单：对于每一个节点，我们获得所有子节点（一级），过滤出高于阈值（默认10）的节点，然后过滤出节点的子标签类别唯一的节点。这样候选列表就得到了。

` list_tag_list = [] threshold = spider.get( 'item_threshold' ) or 10 # iterate all child nodes in a top-down direction for tag in sel.iter(): # get child tags child_tags = self._get_children(tag) if len(child_tags) < threshold: # if number of child tags is below threshold, skip continue else : # have one or more child tags child_tags_set = set(map( lambda x: x.tag, child_tags)) # if there are more than 1 tag names, skip if len(child_tags_set) > 1 : continue # add as list tag list_tag_list.append(tag) 复制代码`

接下来我们将从候选列表中筛选出包含最多文本子节点的节点。听起来有些拗口，打个比方：一个电商网站的列表子项，也就是产品项，一定是有许多例如价格、产品名、卖家等信息的，因此会包含很多文本节点。我们就是通过这种方式过滤掉文本信息不多的列表（例如菜单列表、类别列表等等），得到最终的列表。在代码里我们存为 ` max_tag` 。

` # find the list tag with the most child text tags max_tag = None max_num = 0 for tag in list_tag_list: _child_text_tags = self._get_text_child_tags(self._get_children(tag)[ 0 ]) if len(_child_text_tags) > max_num: max_tag = tag max_num = len(_child_text_tags) 复制代码`

下面，我们将生成列表项的CSS选择器。以下代码实现的逻辑主要就是根据上面得到的目标标签根据其 ` id` 或 ` class` 属性来生成CSS选择器。

` # get list item selector item_selector = None if max_tag.get( 'id' ) is not None : item_selector = f'# {max_tag.get( "id" )} > {self._get_children(max_tag)[ 0 ].tag} ' elif max_tag.get( 'class' ) is not None : cls_str = '.'.join([x for x in max_tag.get( "class" ).split( ' ' ) if x != '' ]) if len(sel.cssselect( f'. {cls_str} ' )) == 1 : item_selector = f'. {cls_str} > {self._get_children(max_tag)[ 0 ].tag} ' 复制代码`

找到目标列表项之后，我们需要做的就是将它下面的文本标签和超链接标签提取出来。代码如下，就不细讲了。感兴趣的读者可以看 [源码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftikazyq%2Fcrawlab%2Fblob%2Fmaster%2Fcrawlab%2Froutes%2Fspiders.py ) 来理解。

` # get list fields fields = [] if item_selector is not None : first_tag = self._get_children(max_tag)[ 0 ] for i, tag in enumerate(self._get_text_child_tags(first_tag)): if len(first_tag.cssselect( f' {tag.tag} ' )) == 1 : fields.append({ 'name' : f'field {i + 1 } ' , 'type' : 'css' , 'extract_type' : 'text' , 'query' : f' {tag.tag} ' , }) elif tag.get( 'class' ) is not None : cls_str = '.'.join([x for x in tag.get( "class" ).split( ' ' ) if x != '' ]) if len(tag.cssselect( f' {tag.tag}. {cls_str} ' )) == 1 : fields.append({ 'name' : f'field {i + 1 } ' , 'type' : 'css' , 'extract_type' : 'text' , 'query' : f' {tag.tag}. {cls_str} ' , }) for i, tag in enumerate(self._get_a_child_tags(self._get_children(max_tag)[ 0 ])): # if the tag is <a...></a>, extract its href if tag.get( 'class' ) is not None : cls_str = '.'.join([x for x in tag.get( "class" ).split( ' ' ) if x != '' ]) fields.append({ 'name' : f'field {i + 1 } _url' , 'type' : 'css' , 'extract_type' : 'attribute' , 'attribute' : 'href' , 'query' : f' {tag.tag}. {cls_str} ' , }) 复制代码`

分页的代码很简单，实现也很容易，就不多说了，大家感兴趣的可以看 [源码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftikazyq%2Fcrawlab%2Fblob%2Fmaster%2Fcrawlab%2Froutes%2Fspiders.py )

这样我们就实现了提取列表项以及列表子项的算法。

## 使用方法 ##

要使用自动提取字段，首先得安装Crawlab。如何安装请查看 [Github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftikazyq%2Fcrawlab ) 。

Crawlab安装完毕运行起来后，得创建一个 **可配置爬虫** ，详细步骤请参考 [[爬虫手记] 我是如何在3分钟内开发完一个爬虫的
]( https://juejin.im/post/5ceb4342f265da1bc8540660 ) 。

创建完毕后，我们来到创建好的可配置爬虫的爬虫详情的 **配置** 标签，输入 **开始URL** ，点击 **提取字段** 按钮，Crawlab将从开始URL中提取列表字段。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b231eb17118f89?imageView2/0/w/1280/h/960/ignore-error/1)

接下来，点击预览看看这些字段是否为有效字段，可以适当增删改。可以的话点击运行，爬虫就开始爬数据了。

好了，你需要做的就是这几步，其余的交给Crawlab来做就可以了。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b23218d09873fa?imageView2/0/w/1280/h/960/ignore-error/1)

## 测试结果 ##

本文在对排名前10的电商网站上进行了测试，仅有3个网站不能识别（分别是因为“动态内容”、“列表没有id/class”、“lxml定位元素问题”），成功率为70%。读者们可以尝试用Crawlab自动提取字段功能对你们自己感兴趣的网站进行测试，看看是否符合预期。结果的详细列表如下。

+--------------+----------+------------------+
|     网站     | 成功提取 |       原因       |
+--------------+----------+------------------+
| 淘宝         | N        | 动态内容         |
| 京东         | Y        |                  |
| 阿里巴巴1688 | Y        |                  |
| 搜了网       | Y        |                  |
| 苏宁易购     | Y        |                  |
| 糯米网       | Y        |                  |
| 买购网       | N        | 列表没有id/class |
| 天猫         | Y        |                  |
| 当当网       | N        | lxml定位元素问题 |
+--------------+----------+------------------+

Crawlab的算法当然还需要改进，例如考虑动态内容和列表没有id/class等定位点的时候。也欢迎各位前来试用，甚至贡献该项目。

**Github** : [tikazyq/crawlab]( https://link.juejin.im?target=http%3A%2F%2Fgithub.com%2Ftikazyq%2Fcrawlab )

如果您觉得Crawlab对您的日常开发或公司有帮助，请加作者微信拉入开发交流群，大家一起交流关于Crawlab的使用和开发。

![](https://user-gold-cdn.xitu.io/2019/3/15/169814cbd5e600e9?imageView2/0/w/1280/h/960/ignore-error/1)