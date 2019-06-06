# 【肥朝】看源码，我为什么推荐IDEA? #

> 
> 
> 
> 本文并不评论Eclipse与IDEA孰好孰坏,但是由于肥朝平时都是使用IDEA开发的,所以推荐IDEA.这个和肥朝平时都是吃粤菜,所以推荐的都是粤菜为主,但是并不是说其他菜不好吃,肥朝不挑食!
> 
> 
> 

## 1.条件断点 ##

看源码的时候,经常遇到这个情况,源码中有个for循环,关键是这个list的size有时候长达数百个.但是我们只想debug一种情况.肥朝就曾经见过,在for循环中打了断点,一直按跳过,按了数十下之后.才找到自己想debug的值.这样效率不高

比如下文这个

` @Test public void testList () throws Exception { List<Integer> list = Arrays.asList( 1 , 2 , 3 , 4 , 5 , 6 , 7 , 8 , 9 , 10 ); for (Integer integer : list) { System.out.println(integer); } } 复制代码`

如果你想debug数字10这种情况,如果你不知道条件断点,那么你可能要一直点9次跳过.我们来看一下条件断点的使用

![](https://user-gold-cdn.xitu.io/2019/4/29/16a667f39fe8fd40?imageView2/0/w/1280/h/960/ignore-error/1)

这样,就只有满足条件的时候才会进入断点了,告别无效的 ` 小手一抖` !

## 2.强制返回值 ##

比如SpringBoot中这个打印Banner的.我们想调试多种情况.就可以利用这个 ` Force Return` ,这样方便我们调试源码中的多种分支流程

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66801481cbee3?imageView2/0/w/1280/h/960/ignore-error/1)

## 3.模拟异常 ##

在做业务开发中,我们有时需要模拟某个方法抛出异常,看看自己的代码是不是像肥朝一样可靠得一逼.但是你每次去写死一个异常,然后再删掉,这种低效的方式有违极客精神.那么我们如果让一个方法抛出异常呢?

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66805f0bb129a?imageView2/0/w/1280/h/960/ignore-error/1)

不过要注意的一点是,这个功能印象中是IDEA 2018年以后的版本才有的功能.

## 4.Evaluate Expression ##

比如我们看源码时遇到这个一个场景,这里有一个 ` byte[]` ,但是我们就想看一下这个的值到底是啥.

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66809694bcb33?imageView2/0/w/1280/h/960/ignore-error/1)

那么我们可以这么操作一波

![](https://user-gold-cdn.xitu.io/2019/4/29/16a6680d5bfc1970?imageView2/0/w/1280/h/960/ignore-error/1)

这个功能的使用场景非常的广,通过这个功能,可以在看源码时,给某个变量赋我们要想的值,从而改变代码的分支走向等等.总之,这个是肥朝看源码中,使用频率最高的功能之一.更多场景,等待 ` 老司机们自己调教!`

## 5.toString的坑(重点) ##

相信看过Dubbo源码的朋友都会遇到过这个一个坑.也就是你把断点打在下面图示的第一个箭头的时候,是无法进入 ` init()` 方法的.但是你把断点打在第二个箭头也就是 ` init()` 方法的时候,是能进入的.曾经也有不少人问过这个问题.

![](https://user-gold-cdn.xitu.io/2019/4/29/16a668119cf9a1f7?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/4/29/16a668137d4d9503?imageView2/0/w/1280/h/960/ignore-error/1)

当然除了这个坑之外,也有类似的坑,如下

![](https://user-gold-cdn.xitu.io/2019/4/29/16a66815ca6fad3e?imageView2/0/w/1280/h/960/ignore-error/1)

所以这个idea的默认设置.建议在一定条件下还是关闭

本文仅为冰山一角， **上百篇原创干货还在路上** ， **扫描下面二维码** 关注肥朝， 让天生就该造火箭的你，不再委屈拧螺丝！

![](https://user-gold-cdn.xitu.io/2019/4/29/16a667ed94034f7e?imageslim)