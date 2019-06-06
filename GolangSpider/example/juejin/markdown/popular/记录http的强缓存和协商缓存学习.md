# 记录http的强缓存和协商缓存学习 #

一、浏览器缓存能够降低资源重复加载并提高网页的加载速度。

浏览器的缓存分为两种，强缓存和协商缓存。

1.基本原理

* 浏览器在加载资源时，根据请求头的expires和cache-control判断是否命中强缓存，是则直接从缓存读取资源，不会发请求到服务器。
* 如果没有命中强缓存，浏览器一定会发送一个请求到服务器，通过last-modified和etag验证资源是否命中协商缓存，如果命中，服务器会将这个请求返回，但是不会返回这个资源的数据，依然是从缓存中读取资源
* 如果前面两者都没有命中，直接从服务器加载资源

2.异同点

（1）相同点

如果命中，都是从客户端缓存中加载资源，而不是从服务器加载资源数据

（2）不同点

强缓存：直接从本地副本对比读取， **不去请求服务器** ，返回的状态码是 **200**

协商缓存：会去服务器比对，若没改变才直接读取本地缓存，返回的状态码是 **304**

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b75aa355b316?imageView2/0/w/1280/h/960/ignore-error/1)

## 二、强缓存 ##

强缓存主要包括expires和cache-control。

1.expires

expires是HTTP1.0中定义的缓存字段。当我们请求一个资源，服务器返回时，可以在Response Headers中增加expires字段表示资源的过期时间

` expires: Thu, 03 Jan 2019 11:43:04 GMT 复制代码`

它是一个时间戳，当客户端再次请求该资源的时候，会把客户端时间与该时间戳进行对比，如果大于该时间戳则已过期，否则直接使用该缓存资源。

但是，有个大问题，发送请求时使用的是 **客户端时间** 去对比。一是客户端和服务端时间可能快慢不一致，另一方面是客户端的时间是可以自行修改的（比如浏览器是跟随系统时间的，修改系统时间会影响到），所以不一定满足预期。

2.cache-control

正由于上面说的可能存在的问题，HTTP1.1新增了cache-control字段来解决该问题，多以当cache-control和expires都存在时，cache-control优先级更高。该字段是一个时间长度，单位秒（s），表示该资源过了多少秒后失效。当客户端请求资源的时候，发现该资源还在有效时间内则使用该缓存，它不依赖客户端时间。cache-control主要有max-age和s-maxage、public和private、no-cache和no-store等值。

（1）max-age和s-maxage

两者是cache-control的主要字段，它们是一个数组，表示资源过了多少秒之后变为无效。在浏览器中，max-age和s-maxage都起作用，而且s-maxage的优先级高于max-age。在代理服务器中，只有s-maxage起作用。可以通过设置max-age为0表示立马过期来向服务器请求资源。

（2）public和private

public表示该资源可以被所有客户端和代理服务器缓存，而private表示该资源仅能客户端缓存。默认值是private，当设置了s-maxage的时候表示允许代理服务器缓存，相当于public。

（3）no-cache和no-store

no-cache表示的是不直接询问浏览器缓存情况，而是去向服务器验证当前资源是否更新（即协商缓存）。 **no-store则更狠，完全不使用缓存策略，不缓存请求或响应的任何内容，直接向服务器请求最新** 。由于两者都不考虑缓存情况而是直接与服务器交互，所以当no-cache和no-store存在时会直接忽略max-age等。

3.pragma

它的值有no-cache和no-store，表示意思同cacha-control，优先级高于cache-control和expires，即三者同时出现时，先看pragma -> cache-control -> expires。

` pragma: no-cache 复制代码`

## 三、协商缓存 ##

上面的expires和cache-control都会访问本地缓存直接验证看是否过期，如果没过期直接使用本地缓存，并返回200。但如果设置了no-cache则本地缓存会被忽略，会去请求服务器验证资源是否更新，如果没更新才继续使用本地缓存，此时返回的是304，这就是协商缓存。协商缓存主要包括last-modifed和etag。

1.last-modified

last-modified记录资源最后修改的时间。启用后，请求资源之后的响应头会增加一个last-modified字段，如下：

` last-modified : Thu, 20 Dec 2018 11:36:00 GMT 复制代码`

当再次请求该资源时，请求头中会带有if-modified-since字段，值是之前返回的last-modified的值，如： ` if-modified-since:Thu, 20 Dec 2018 11:36:00 GMT` 。服务端会对比该字段和资源的最后修改时间，若一致则证明没有被修改，告知浏览器可直接使用缓存并返回304；若不一致则直接返回修改后的资源，并修改last-modified为新的值。

但last-modified有以下两个缺点：

* 只要编辑了，不管内容是否真的有改变，都会以这最后修改的时间作为判断依据，当成新资源返回，从而导致了没必要的请求响应，而这正是缓存本来的作用即避免没必要的请求。
* 时间的精确度只能到秒，如果在一秒内的修改时检测不到更新的，仍会告知浏览器使用旧的缓存。

2.etag

为了解决last-modified上述问题，有了etag。etag会基于资源的内容编码生成一串唯一的标识字符串，只要内容不同，就会生成不同的etag。启用etag之后，请求资源后的响应返回会增加一个etag字段，如下：

` ETag:W/"14e2-16b26fefeb0"` ``

当再次请求该资源时，请求头会带有if-no-match字段，值是之前返回的etag值，如：ETag:W/"14e2-16b26fefeb0"。服务端会根据该资源当前的内容生成对应的标识字符串和该字段进行对比，若一致则代表未改变可直接使用本地缓存并返回304；若不一致则返回新的资源（状态码200）并修改返回的etag字段为新的值。

可以看出etag比last-modified更加精确地感知了变化，所以etag优先级也更高。不过从上面也可以看出etag存在的问题，就是每次生成表示字符串会增加服务器的开销。所以要如何使用last-modified和etag还需要根据具体需求进行权衡。

## 四、访问刷新分析 ##

我们将访问和刷新分为以下三种情况：

* 标签进入、输入url回车进入
* 按刷新按钮、F5刷新、网页右键“重新加载”
* ctrl+F5强制刷新

假设当前有这么一个index页面，返回的响应信息如下：

` cache-control: max-age=72000 expires: Tue, 20 Nov 2018 20:41:14 GMT last-modified: Tue, 20 Nov 2018 00:41:14 GMT 复制代码`

**1、标签进入、输入url回车进入**

这种情况下会根据实际设计的缓存策略去判断。

1.由于该例没有设置no-cache和no-store，所以默认先走强缓存路线。根据cache-control（expires优先级低）判断缓存是否过期，若没有过期则此时返回200(from cache)。

2.若本地缓存已经过期再走协商缓存路线，根据之前的last-modified值去与服务器比对，若这个时间之后没有改过则去读本都缓存，返回304(not modified)。

3.否则返回新的资源，状态码200(ok)，并更新返回响应的last-modified值。

2、按刷新按钮、F5刷新、网页右键“重新加载”

这种情况下，实际是浏览器将cache-control的max-age直接设置成了0，让缓存立即过期，直接走协商缓存路线。发送的请求头如下：

` cache-control: max-age=0 if -modified-since: Tue, 20 Nov 2018 00:41:14 GMT 复制代码`

3.ctrl+F5强制刷新

浏览器不仅会对本地文件过期，而且不会带上if-modified-since，if-no-match，相当于之前从来没有请求过，返回结果是200