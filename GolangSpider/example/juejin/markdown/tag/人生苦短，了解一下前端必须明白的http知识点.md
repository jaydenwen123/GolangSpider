# 人生苦短，了解一下前端必须明白的http知识点 #

半年了，没有在个人的文章发表任何话题，可能是被喷的多了，也可能是累了，以前的分享都是以边学边写的模式做出的文章，当时也是因为VUE的热点写了一堆看似现在纯小白学的东西。但是言归正传，继续分享我自己所学到的http对于前端需要了解的知识点。

对于http的报文格式就不多细说了，因为做为前端开发，我们需要知道前后端联调时的请求和响应之间请求头和返回头之间的关系和每个字段中的涵意，静态文件资源在加载时我们所观察到可性能优化的点，和一些日常请求报错如何去解决的坑，更重要的是面试的时候如何去从容的应对面试官

### 以下的讲解纯属于个人理解，肯定会有错误和理解不到位的点，请在下方用你们猿族的语言喷起来 ###

> 
> 
> 
> 简单跨域的解决方式
> 
> 

跨域是一个老生常谈的话题，面试官问我如何解决跨域，以前只会和面试官说用 ` webpack` 的 ` proxy` 做代理，叫后端大哥给我本地启一个 ` nginx` 就可以了，那往往在一些特殊的情况下，后端大哥来大姨妈了，进入一个新公司的让你维护一个很老的项目，并没有用到工程化这些东西，而且后端又来了一位新的高不高，低不低的后端工程师，此时对跨域根本性的知识点了解才能解决根本性的问题。

## 猿族前端 VS 猿族后端java ##

` 后端说` : 前端同志，我们先调一个get请求的一个接口，地址我给你， ` http:www.pilishou.com/getname/list`

` 前端操作中。。。`

` fetch( 'http://http:www.pilishou.com/getname/list' , { method: 'GET' }) 复制代码`

写了一个这样的请求，听从后端大哥向服务端发送，此时浏览器报了一个这样的错误 ` Failed to load http://http:www.pilishou.com: No 'Access-Control-Allow-Origin' header is present on the requested resource. Origin 'null' is therefore not allowed access. If an opaque response serves your needs, set the request's mode to 'no-cors' to fetch the resource with CORS disabled.`

` 小白前端会说` : 大哥，你这是什么接口，请求了还尼妈报错，你那里什么鬼！

` 大牛前端会说` : 大哥，帮个忙，你那里忘记设跨域头了。

` 小白后端会说` ：大哥，丢你老母啊，你会不会调接口，报错还找我，我这里postman上面调的一点问题都没有

` 大牛后端会说` : 大哥，等一下，我的跨域头忘记设了，稍等

` 原理讲解` :

在本地向不同域请求的时候，浏览器会做一个 ` Origin` 请求头的验证，如果没有设置，在不同域名下或者本地请求时浏览器会向服务端发送请求，服务端也会客户端发送对应的值，但是浏览器考虑到安全策略，会进行一个关于头信息的报错，此时对于后端来说，需要在response的返回头中加入 ` 'Access-Control-Allow-Origin': '*'` ,来告诉浏览器我允许你进行一个跨域请求，不用报错，把值返回给请求者，这样你就可以安然的拿到数据。同时这样也会导致任何一个域名发送过来的请求，都允许跨域的情况下，可以针对 ` 'Access-Control-Allow-Origin': '此处设置指定的域名'`

> 
> 
> 
> 复杂跨域的解决方式
> 
> 

此时前端唱起来一首抖音网红歌，我知道我对你不仅仅是喜欢！。。。。。

` 后端说` ：小伙，这里有一个接口，需要遵循resutful接口，用PUT方法， ` ·http:www.pilishou.com/getname/update`

` 前端操作中。。。。`

` fetch( 'http://http:www.pilishou.com/getname/list' , { method: 'PUT' }) 复制代码`

继续按部就班的写了一个这样的请求，然后又发现浏览器报了这样一个错误 ` Failed to load http://http:www.pilishou.com: Method PUT is not allowed by Access-Control-Allow-Methods in preflight response`

` 小白前端会说` : 大哥，你接口又怎么了，GET，POST都行，PUT怎么不行，肯定是你的问题，我别的什么都没动啊。

` 大牛前端会说` : 大哥，帮个忙，你把请求头中加一些允许跨域的方法。

` 小白后端会说` ：大哥，丢你老母啊，你不会调接口，报错还找我，我这次postman上面调的还是一点问题都没有

` 大牛后端会说` : 大哥，等一下，我加一些允许跨域的方法，稍等

原理讲解:

` 在简单的跨域请求中 1.请求方法是以下三种方法之一： HEAD GET POST 2.HTTP的头信息不超出以下几种字段： Accept Accept-Language Content-Language Last-Event-ID Content-Type：只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain 复制代码`

如果不超过以上的限制,后端则只需要提供一个允许跨域的Origin就可以了，如果在请求方法超过了以上三种，需要添加'Access-Control-Allow-Methods': 'PUT',同样浏览器为了安全，不允其它请求方法在台端没有设置允许的方法中进行一个跨域请求

同理复杂请求还包函着别的需要后端设置允许一些跨域请求的方式，比如通常会出现的:

* 添加自定义头

` fetch( 'http://127.0.0.1:8887' , { method: 'PUT' , headers: { 'x-header-f' : '1234' , } }) 复制代码`

` 报错信息` Failed to load http://http:www.pilishou.com: Request header field x-header-f is not allowed by Access-Control-Allow-Headers in preflight response.

` 解决方案` 需要服务端加上允许那些自定义头进行一个跨域仿问 'Access-Control-Allow-Headers': 'x-header-f',

* 添加不包括上面三者的请求类型

` fetch( 'http://127.0.0.1:8887' , { method: 'PUT' , headers: { 'x-header-f' : '1234' , 'content-type' : 'json' } }) 复制代码`

` 报错信息` Failed to load http://http:www.pilishou.com: Request header field content-type is not allowed by Access-Control-Allow-Headers in preflight response.

` 解决方案` 需要服务端加上允许那些自定义头进行一个跨域仿问 'Access-Control-Allow-Headers': 'content-type'这个请求头信息

### 复杂的跨域请求中，包括着预请求方案 ###

在非同源的请求情况下，浏览器会首先进行Option请求，所谓的预请求，就是试探性请求，向服务端请求的时候，发现此接口设置了允许对应的请求方法或者请求头，会再次发送真正的请求，分别一共会向后台发送两次请求，拿自己想要的数据，在 ` OPTION` 请求时，服务端也会返回数据，但是在浏览器层被做了屏闭，如果没有检测出对应的跨域设置则会报出对应的错误。

### 减少预请求的认证次数 ###

在本地联调时，每次发送一个非简单请求时都会发送一个预请求，预请求也是一个花费时间和资源的操作，就像一次实名质认证过了，在一定时间内就不用实名质认正了，原理一样，如果当前请求的域名第一次认证通过，则在一定的时间内不需要进行一个二次认证，但是需要进行一次认证的时间控制，通过'Access-Control-Max-Age': '860000'，返回，一旦在这个时间之内再次发送时，直接发送真正的请求，不通要再通过预请求 ` Option` 方法进行一个探测认证。

> 
> 
> 
> cache-control 的使用场景和性能优化
> 
> 

` cache-control` 这个东西就是对服务端拉取的静态资源打上一个缓存标志

对于 ` cache-control` 可以设置几种模式，通常前端工程师又需要知道那几种模式

* max-age = 10000 (以秒为音位，根据需求设定)
* no-cache (每次进行请求时都要向服务端进行验证，需要配合etag,Last-Modified)使用
* no-store (每次请求都需要向服务端拉取新的资源)
* privite (私有的，不经过代理缓存)
* public (公有的，如果本地失效，代理缓存存在的话可以从代理缓存进行通知用过期的资源)

### max-age ###

当加载完资源时，浏览器会自动给我们存储到内存当中，但是浏览内部的失效时间是由内部机制控制的，在用nginx做静态资源的时候，在刷新的时候，浏览会向服务端再次发送是否过期的认证，在资源缓存时间的确定情况下，通过max-age指定强缓存后,浏览器再次加载同样的资源文件时，只需要从memory或者disk上面进行拉取复用

达到以上的功能需要在返回资源的服务端的对返回的资源设置'cache-control': 'max-age=时间（以秒为单位）',当再次刷新页面的时候，在设置的时间之内，刷新页面，不清除缓存的情况下都会重新拉取内存了中的缓存资源。

### no-cache ###

` no-cache` 字面的字意是不缓存的意思，但是很容易迷惑人，但是本质的函意，意味着每次发送请求静态资源时都需要向服务端进行一次过期认证，通常情况下，过期认真证需要配合 ` （etag和Last-Modified）` 进行一个比较，这个话题后继再展开讨论，如果验证并没有过期，则会发送304的状态码，通知浏览进复用浏览器的缓存

### no-store ###

` no-store` 代表每次资源请求都拉取资源服务器的最新资源，就算同时设置max-age , no-store, no-store的优先级则最高，此时max-age则不生效，同样的会从服务端拉取最新的资源

### private vs public ###

在资源请求时，有些情况不会直接到原资源服务器发送请求，中间会经过一些代理服务器，比如说cdn,nginx等一些代理服务器，如果写入public的情况下，所有的代理服务器同样也会进行缓存，比如说s-maxage就是在代理缓存中生效的，如果本地max-age过期了，则会通过代理缓存，代理缓存并没有过期，会告诉浏览器还是可以用本地过期的缓存，但对于private中间代理服务器则不会生效，直接从浏览器端向原服务器进行一个验证。

> 
> 
> 
> 缓存验证 Last-Modified 和 Etag
> 
> 

### Last-Modified ###

最后修改时间，一般在服务端，对文件的修改都会有一个修改时间的记录，在nginx做静态资源时，nginx会返回一个Last-Modified最后修改的时间，在浏览器再次请求的时候，会把对应的If-Modified-Since和 If-UnModified-Since在请求头中再次发送给服务端，告诉服务端上次你给我文件改动的时间，但是Last-Modified只能以秒为单位，在有些情况下，是不够精确的

### Etag ###

是一个更加比较严格的验证，主要通过一些数据签名，每个数据都有自己的唯一签名，一旦数据修改，则会生成另一个唯一的签名，最典型的做法就是对内容做一个hash计算，当浏览器端向服务端再请求的时会带上 IF-Match 或者 If-Non-Match,当服务端接收到后之后会对比服务端的签名和浏览器传过来的签名，这也是弥补了Last-Modified只能以秒为单位，在有些情况下，是不够精确的情况

### Last-Modified和Etag 配合 no-cache 使用 ###

通常只会在 cache-control 在 no-cache的情况下，浏览器也会对资源进行一个缓存， 同时会对服务端进行一个认证过期，一旦服务端返回304状态码，则说明可以复用浏览器的缓存，则会向服务端重新请求数据。

cookie的策略机制

cookie则是一个服务端和用端之间一个像身份证认证一样的东西，一旦后端在返回头中设置了cookie,则在response中会出现设置的cookie数据，同时也会存在浏览器的application/cookie中，当每次发送请求的时候都会在request的头中带上当前域名下的cookie信息

### 健值对方式设置 ###

'Set-Cookie': 'id=1',

### 设置过期时间 ###

通常情况，在不设置过期时间的时候，浏览器关闭的时候，则cookie，则会失效，我们可以通过max-age或者expire进行一个cookie失效时间的设置

### 不可获取的cookie ###

如果在不设置httponly的情况下，可以通过document.cookie进行读取，在不同情况下，考虑安全性，可以通过httponly设置，在document.cookie则获取不到。

### https下的secure cookie ###

如果设置了secure只有在https的服务下才会把字段写入application/cookie中，虽然在response有发送cookie这个字段，但是浏览器在识别不是https服务时，会进行一个乎略

### 二级域名下与二级域名的cookie传输 ###

讲一个例子:

公司的所有内部系统都全要走一个登陆系统。也可能说sso单点登陆，如果登陆是 ` sso.pilishou.com` 的二级域名下，而你自己的开发的时候环境是 ` localhost:9999` 端口，当登陆成功时，此时cookie是设在 ` sso.pilishou.com` 域名下，在本地 ` 127 .0.0.1` 下发送请求，根本拿不到 ` sso.pilishou.com` 下的 ` cookie` 信息，cookie根本不会从request header中带过去，可以通过host的映射，把 ` 127.0.0.1` 映射成 ` web.pilishou.com`

但是问题来了， ` sso` 和 ` web` 都是二级域名，在 ` web` 下同样拿不到 ` sso` 下的 ` cookie` ，此时解决办法，在 ` sso` 登成功后，需要后台配合把 ` cookie` 的信息通过 ` Dioman` 设置到 ` pilishou.com` 的主域下

在 ` web` 二级域名下就可以拿到 ` sso` 下请求成功后设置的 ` cookie` 信息，在不设置httponly情况下，尝试用 ` document.cookie` 可以拿到自己想要的 ` cookie` 信息，但是在发送的时候，发现 ` request头` 中根本没有把 ` cookie信息` 带入请求，在 ` fetch请求` 中我们要设置 ` credentials: 'include'` ，意思代表允许请求时带上跨域cookie,此时就会发现 ` cookie` 带入了 ` request头部`

经历了这么多的设置，在联调的时候，后端同样也需要配合你的行为，需要后台工程师也需要配置在返回头中加入 ` 'Access-Control-Allow-Credentials': 'true'` ，允许进行 ` cookie的跨域` ，

但是问题又来了，真TMD的好多问题，少一步都不行，此时你的浏览器又会报错，在设置跨域 ` cookie` 的时候，不允许 ` response header` 设置 ` Origin` 设置为 ` *` ，只能设置指定的域名进行一个跨域仿问，此时还需要后端工程师配合把前面的 ` *` 改成你指定当前 ` web.pilishou.com` 。

如果讲cookie你就用这么多一套流讲死面试官。

> 
> 
> 
> http长连接与性能优化的各种架构方式
> 
> 

在以前没有打包工具，或者没有应用到打包工具的时候，一个大项目会有一堆js,一堆css，会引起各种问题，会导致引入资源会出现混乱，资源加载慢，有些时候页面呈现了，点击的时候没有任何响应，这个需要从http请求资源时，经过三次握手后创建的TCP连接说起。

因为每个浏览器的执行策略不一样，所以我只针对 ` Chorme` 来说，打开开发者工具，点击 ` network` ,通过右健点击 ` Name` ,有一个 ` connect id` , Chrome可以一次性创建六个并发连接，但是六个并发连接会阻塞后面的资源的请求，如果前六个资源文件很大，后面的资源请求会被一直阻塞着，会进行一个队列的等待请求，当页面在网络不稳的情况下， ` HTML，CSS` ，已经加载好了，也渲染完毕， ` JS` 终于等到请求，但是突然网速变差，用户点击此时是没有任何响应的，因为 ` JS` 根本还没有加载好

为了验证，打开网络资源多的网站，把网速调到 ` 2G模式` ，会发现，一开始只会出现 ` 6次connect连接` ，但是也不是一下子全出来，因为创建 ` TCP连接` 需要经过三次握手，这中间也是需要时间，当6个连接创建完成时，又回到了串行的方式，除非只有 ` connect` 连接请求完成后才会让出连接资源，让下一个队列中的请求，进行复用，不用再创建新的TCP，但是在 ` TCP` 最后的关闭，浏览器会与服务器自行进行一个协商关闭，也可以设置关闭时间，在多长时间没有请求后，才进行一个连接关闭，在观查 ` connect id` 会发现，只会出现6个 ` connect id` ,其余的全会被复用，如果有些资源是复用的其它网站的，会另开新的 ` connect id`

` 解决方案` ：

所以现在的对于 ` spa` 的页面，都采取了，资源合并，把 ` CSS，JS` 进行一个合并，通常在VUE中都打出4个文件， ` vendor.js, app.js, manifest.js, app.css`

能让浏览器充分的刚好利用让一个工程上的主文件一次性全都通过 ` TCP连接` 并行下载下来，无论从性能速上还是解决造成用户无响应的解决方案，那我再简单的讲解一次为什么要分成这四个文件。

1. ` Vendor.js` 一般是 ` node_modules` 文件，不会轻意更改，所以可以通过浏览器缓存，能长期进行一个缓存 2. ` app.js` 一般都是业务代码的文件，业务代码的文件，对于公司来说是业务代码很一个迭代很频繁的事，所以当用户拉取资源的时候，只需要拉需 ` app` 的新资源， ` app.js` 中还可以分每个 ` module` 进行一个资源更新， 3. ` manifest.js` 是一个 ` runtime运行时的文件` ，无论 ` app.js` 或者 ` vendor` 有改动， ` masfiste` 文件则就会改动，所以也进行一个单独更新 4。 ` app.css` 。是一个综合考虑，虽然如果如改动一部分小资源，但是也会重新拉取，但是节省了请求的次数，对于合并自己可能根据项目进行一个着情考虑。

#### 每个文件其实更新并不是通过什么缓存的设置，而是在每个js或者css后面会跟一个文件的hash，这个hash是打包工具给我们做的，一旦文件有改动，就会重新生成一个hash,浏览器在加载资源的时候，发现没有找到对应的缓存文件，则会向服务端进行一个重新请求。 ####

> 
> 
> 
> 多次复用，和单次复用的决择性
> 
> 

上面我们讲了因为浏览器的请求速度影响和，TCP连接的限制，我们采取了以上的方案，但是每个方案是针对不同的场景和架构的，对于后台管理项目，基本上公司都是统一工程化做的，所以的工程方案都是采用一套都或几套，但是采用项目基础文件都是一样的，要升级也会根据项目特定需要才进行升级的，对于公司的内部系统，采用最好的方式就是放弃初次加载的性能，利用缓存进行多项目缓存复用

通常一个vue的项目， ` vue.js vuex.js router.js` 和一些公共的内部js文件都是在项目架构中集成的

#### 举个例子 ####

公司的内部项目一般都有三个环境，加上你本地调试有四个，如果把这些文件全打到vendor中，会产生只要项重发之后，或者切换环境，在这个四个资源环境中就不能行成一个复用，因为域名都是不一样的，所以浏览器找缓存不能共享，往往这些文件在所有项目中，所有环境中都是不会反复变的文件，一次加载，任何环境，任何项目共享利用缓存资源

1。我们可以利用cdn把以上前面提到的文件进行利用 2。也可以把文件放到一个域名下的公共目录下，进行利用。

> 
> 
> 
> from memory cache 和 from disk cache
> 
> 

* ` from memory cache` 从内存中拉取的缓存
* ` from disk cache` 从磁盘上拉取的缓存

在资源拉取过后，这里还是针对的Chrome进行解释，浏览器在拉取资源后，会对资源进行磁盘和内存进行缓存，而 ` css文件` 会缓存到磁盘上， ` html.js,img等文件` 都会在内存和磁盘进行缓存，当刷新页面时，除了在特定的资源中返回头中写入 ` cache-contorl: no-cache` 或者 ` no-store` 的情况下都会直接从缓存中拉取资源，会在size中显示 ` from memory cache` ，而css文件则显示 ` from disk cache` , 但是 ` no-cache` 验证没有过期，则还会返回304进行读取缓存，只是到原服务器进行了一个验证。

> 
> 
> 
> meta http-equiv="Cache-Control" content="no-cache" 设置的备要性
> 
> 

此时前端和后端同学前后端已经联调好了，发到测试环境，让测试同志进行测试。

` 测试说` ：你页面中一个字写错了，改一下，重新发包我再来测，测试关闭浏览器，刷了一会抖音

` 前端一顿操作后` 。。。。这一顿操作猛如虎

` 前端说` ：好了，你测吧。我发上去了，此时也关闭了浏览器，心里想测试要测试默默JJ的，我先刷一会抖音。

` 测试一顿操作后` 。。。打开浏览器，把地址输入了进去，回车后。。。

` 测试说` ： 你到底改没改啊，怎么没有效果。

前端此时也打开浏览器，输入地址，一看，wc什么情况。开始怀疑人生了。。。我明明改了。怎么没效果。然后又是一顿猛如虎的打开文件看了看，又重新发包

` 问题总结` ：

根本原因，进行一个分析，正是因为缓存问题而导致，浏览器对html页会进行一个自动缓存，但是正常刷新情况下，如果用nginx做一个静态资源的情况下，都会进行一个304的重新向服端进行一个资源是否改动的验证，如果没有改动则进行一个304的缓存利用

#### 当关闭浏览器进程的时候，缓存在内存中的资源会随着浏览器的闭毕一起清除，当再次打开浏览器的时候会从磁盘上读取缓存，这时候如果没有设置 ` meta http-equiv="Cache-Control" content="no-cache"` ，当打开浏览器再次仿问的时候，html页面初次会进行浏览器的磁盘上读取就是from disk cache,那此时肯定用的还是原本旧的资源，这就是问题产生的根本，所以在加入每次都从原服务器验证资源，在打开浏览器的时候就不会出来资源没有及时更新的问题。 ####

> 
> 
> 
> redirect 重定向的坑
> 
> 

重定向在response中会有一个location字段进行重定义，比如说返回值/list,需要我们重定向到/list的页面，但是在响应码中，可以返回 302或者301

### 301适合永久重定向 ###

301比较常用的场景是使用域名跳转。比如，我们访问 http://www.baidu.com 会跳转到 https://www.baidu.com，发送请求之后，就会返回301状态码，然后返回一个location，提示新的地址，浏览器就会拿着这个新的地址去访问。

### 302用来做临时跳转 ###

302和301的区别则是设置了302如果再次访问则是从服务端再次拉取资源，然后进行重定向。301则是如果有缓存文件，则直接读缓存文件上响应头上的重定向位置，如果原服务端重定向的位置有变化，则只能通过用户清除缓存进行重新拉取新资源进行再次重定向，所以301的使用需要严谨。

> 
> 
> 
> csp的理解 (cotent Security Policy) 内容安全策略
> 
> 

为了让我们网页更加安全

1。限制资源获取 2。资源获取越权

可以通过设置 default-src 设置全局需要资源的内容，也可以设置资源类型的范围

1。connect-src 我们连接的资源 2。style-src 样式请求的资源 3。script-src 脚本的请求资源 。。。等等

可以通过响应头的返回设置'Content-Security-Policy'进行设置

有些情况一些xss攻击是通过inline scrpit进行注入一些代码进行攻击，可以通过设置进行一个禁用。可以设置'Content-Security-Policy': 'default-src http: https:'对inline scrpit进行一个禁用。设置之后，后报 ` Refused to execute inline script because it violates the following Content Security Policy directive: "default-src http: https:". Either the 'unsafe-inline' keyword, a hash ('sha256-9aPvm9lN9y9aIzoIEagmHYsp/hUxgDFXV185413g/Zc='), or a nonce ('nonce-...') is required to enable inline execution. Note also that 'script-src' was not explicitly set, so 'default-src' is used as a fallback`.错误。

不允许引入外部连接:

可以设置 ''Content-Security-Policy': 'default-src \ ` self\` '' 进行设置，如果引用了外部的资源则会报 ` Refused to load the script 'http://static.ymm56.com/common-lib/jquery/3.1.1/jquery.min.js' because it violates the following Content Security Policy directive: "default-src` self ` ". Note that 'script-src' was not explicitly set, so 'default-src' is used as a fallback.` 错误

如果需要指定外链的地址，则可以，在default-src加入指定的地址

其余的则可以根据Content-Security-Policy' 内容安全策略文档进行设置。