# Nginx与安全有关的几个配置 #

> 
> 
> 
> 安全无小事，安全防范从nginx配置做起
> 
> 

上一篇文章 [《Nginx的几个常用配置和技巧》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FJyUnN_OtQ2NtXcH0mtCJUg ) 收到了不错的反馈，这里再总结下nginx配置中与安全有关的一些配置

## 隐藏版本号 ##

` http { server_tokens off; } 复制代码`

经常会有针对某个版本的nginx安全漏洞出现，隐藏nginx版本号就成了主要的安全优化手段之一，当然最重要的是及时升级修复漏洞

## 开启HTTPS ##

` server { listen 443; server_name ops-coffee.cn; ssl on; ssl_certificate /etc/nginx/server.crt; ssl_certificate_key /etc/nginx/server.key; ssl_protocols TLSv1 TLSv1.1 TLSv1.2; ssl_ciphers HIGH:!aNULL:!MD5; } 复制代码`

**ssl on：** 开启https

**ssl_certificate：** 配置nginx ssl证书的路径

**ssl_certificate_key：** 配置nginx ssl证书key的路径

**ssl_protocols：** 指定客户端建立连接时使用的ssl协议版本，如果不需要兼容TSLv1，直接去掉即可

**ssl_ciphers：** 指定客户端连接时所使用的加密算法，你可以再这里配置更高安全的算法

## 添加黑白名单 ##

白名单配置

` location /admin/ { allow 192.168.1.0/24; deny all; } 复制代码`

上边表示只允许192.168.1.0/24网段的主机访问，拒绝其他所有

也可以写成黑名单的方式禁止某些地址访问，允许其他所有，例如

` location /ops-coffee/ { deny 192.168.1.0/24; allow all; } 复制代码`

更多的时候客户端请求会经过层层代理，我们需要通过 ` $http_x_forwarded_for` 来进行限制，可以这样写

` set $allow false ; if ( $http_x_forwarded_for = "211.144.204.2" ) { set $allow true ; } if ( $http_x_forwarded_for ~ "108.2.66.[89]" ) { set $allow true ; } if ( $allow = false ) { return 404; } 复制代码`

## 添加账号认证 ##

` server { location / { auth_basic "please input user&passwd" ; auth_basic_user_file key/auth.key; } } 复制代码`

关于账号认证 [《Nginx的几个常用配置和技巧》]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FJyUnN_OtQ2NtXcH0mtCJUg ) 文章中已有详细介绍，这里不赘述

## 限制请求方法 ##

` if ( $request_method !~ ^(GET|POST)$ ) { return 405; } 复制代码`

` $request_method` 能够获取到请求nginx的method

配置只允许GET\POST方法访问，其他的method返回405

## 拒绝User-Agent ##

` if ( $http_user_agent ~* LWP::Simple|BBBike|wget|curl) { return 444; } 复制代码`

可能有一些不法者会利用wget/curl等工具扫描我们的网站，我们可以通过禁止相应的user-agent来简单的防范

Nginx的444状态比较特殊，如果返回444那么客户端将不会收到服务端返回的信息，就像是网站无法连接一样

## 图片防盗链 ##

` location /images/ { valid_referers none blocked www.ops-coffee.cn ops-coffee.cn; if ( $invalid_referer ) { return 403; } } 复制代码`

**valid_referers：** 验证referer，其中 ` none` 允许referer为空， ` blocked` 允许不带协议的请求，除了以上两类外仅允许referer为www.ops-coffee.cn或ops-coffee.cn时访问images下的图片资源，否则返回403

当然你也可以给不符合referer规则的请求重定向到一个默认的图片，比如下边这样

` location /images/ { valid_referers blocked www.ops-coffee.cn ops-coffee.cn if ( $invalid_referer ) { rewrite ^/images/.*\.(gif|jpg|jpeg|png)$ /static/qrcode.jpg last; } } 复制代码`

## 控制并发连接数 ##

可以通过 ` ngx_http_limit_conn_module` 模块限制一个IP的并发连接数

` http { limit_conn_zone $binary_remote_addr zone=ops:10m; server { listen 80; server_name ops-coffee.cn; root /home/project/webapp; index index.html; location / { limit_conn ops 10; } access_log /tmp/nginx_access.log main; } } 复制代码`

**limit_conn_zone：** 设定保存各个键(例如 ` $binary_remote_addr` )状态的共享内存空间的参数，zone=空间名字:大小

大小的计算与变量有关，例如 ` $binary_remote_addr` 变量的大小对于记录IPV4地址是固定的4 bytes，而记录IPV6地址时固定的16 bytes，存储状态在32位平台中占用32或者64 bytes，在64位平台中占用64 bytes。1m的共享内存空间可以保存大约3.2万个32位的状态，1.6万个64位的状态

**limit_conn：** 指定一块已经设定的共享内存空间(例如name为 ` ops` 的空间)，以及每个给定键值的最大连接数

上边的例子表示同一IP同一时间只允许10个连接

当有多个 ` limit_conn` 指令被配置时，所有的连接数限制都会生效

` http { limit_conn_zone $binary_remote_addr zone=ops:10m; limit_conn_zone $server_name zone=coffee:10m; server { listen 80; server_name ops-coffee.cn; root /home/project/webapp; index index.html; location / { limit_conn ops 10; limit_conn coffee 2000; } } } 复制代码`

上边的配置不仅会限制单一IP来源的连接数为10，同时也会限制单一虚拟服务器的总连接数为2000

## 缓冲区溢出攻击 ##

**缓冲区溢出攻击** 是通过将数据写入缓冲区并超出缓冲区边界和重写内存片段来实现的，限制缓冲区大小可有效防止

` client_body_buffer_size 1K; client_header_buffer_size 1k; client_max_body_size 1k; large_client_header_buffers 2 1k; 复制代码`

**client_body_buffer_size：** 默认8k或16k，表示客户端请求body占用缓冲区大小。如果连接请求超过缓存区指定的值，那么这些请求实体的整体或部分将尝试写入一个临时文件。

**client_header_buffer_size：** 表示客户端请求头部的缓冲区大小。绝大多数情况下一个请求头不会大于1k，不过如果有来自于wap客户端的较大的cookie它可能会大于 1k，Nginx将分配给它一个更大的缓冲区，这个值可以在 ` large_client_header_buffers` 里面设置

**client_max_body_size：** 表示客户端请求的最大可接受body大小，它出现在请求头部的Content-Length字段， 如果请求大于指定的值，客户端将收到一个"Request Entity Too Large" (413)错误，通常在上传文件到服务器时会受到限制

**large_client_header_buffers** 表示一些比较大的请求头使用的缓冲区数量和大小，默认一个缓冲区大小为操作系统中分页文件大小，通常是4k或8k，请求字段不能大于一个缓冲区大小，如果客户端发送一个比较大的头，nginx将返回"Request URI too large" (414)，请求的头部最长字段不能大于一个缓冲区，否则服务器将返回"Bad request" (400)

同时需要修改几个超时时间的配置

` client_body_timeout 10; client_header_timeout 10; keepalive_timeout 5 5; send_timeout 10; 复制代码`

**client_body_timeout：** 表示读取请求body的超时时间，如果连接超过这个时间而客户端没有任何响应，Nginx将返回"Request time out" (408)错误

**client_header_timeout：** 表示读取客户端请求头的超时时间，如果连接超过这个时间而客户端没有任何响应，Nginx将返回"Request time out" (408)错误

**keepalive_timeout：** 参数的第一个值表示客户端与服务器长连接的超时时间，超过这个时间，服务器将关闭连接，可选的第二个参数参数表示Response头中Keep-Alive: timeout=time的time值，这个值可以使一些浏览器知道什么时候关闭连接，以便服务器不用重复关闭，如果不指定这个参数，nginx不会在应Response头中发送Keep-Alive信息

**send_timeout：** 表示发送给客户端应答后的超时时间，Timeout是指没有进入完整established状态，只完成了两次握手，如果超过这个时间客户端没有任何响应，nginx将关闭连接

## Header头设置 ##

通过以下设置可有效防止XSS攻击

` add_header X-Frame-Options "SAMEORIGIN" ; add_header X-XSS-Protection "1; mode=block" ; add_header X-Content-Type-Options "nosniff" ; 复制代码`

**X-Frame-Options：** 响应头表示是否允许浏览器加载frame等属性，有三个配置 ` DENY` 禁止任何网页被嵌入, ` SAMEORIGIN` 只允许本网站的嵌套, ` ALLOW-FROM` 允许指定地址的嵌套

**X-XSS-Protection：** 表示启用XSS过滤（禁用过滤为 ` X-XSS-Protection: 0` ）， ` mode=block` 表示若检查到XSS攻击则停止渲染页面

**X-Content-Type-Options：** 响应头用来指定浏览器对未指定或错误指定 ` Content-Type` 资源真正类型的猜测行为，nosniff 表示不允许任何猜测

在通常的请求响应中，浏览器会根据 ` Content-Type` 来分辨响应的类型，但当响应类型未指定或错误指定时，浏览会尝试启用MIME-sniffing来猜测资源的响应类型，这是非常危险的

例如一个.jpg的图片文件被恶意嵌入了可执行的js代码，在开启资源类型猜测的情况下，浏览器将执行嵌入的js代码，可能会有意想不到的后果

另外还有几个关于请求头的安全配置需要注意

**Content-Security-Policy：** 定义页面可以加载哪些资源，

` add_header Content-Security-Policy "default-src 'self'" ; 复制代码`

上边的配置会限制所有的外部资源，都只能从当前域名加载，其中 ` default-src` 定义针对所有类型资源的默认加载策略， ` self` 允许来自相同来源的内容

**Strict-Transport-Security：** 会告诉浏览器用HTTPS协议代替HTTP来访问目标站点

` add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" ; 复制代码`

上边的配置表示当用户第一次访问后，会返回一个包含了 ` Strict-Transport-Security` 响应头的字段，这个字段会告诉浏览器，在接下来的31536000秒内，当前网站的所有请求都使用https协议访问，参数 ` includeSubDomains` 是可选的，表示所有子域名也将采用同样的规则

![](https://user-gold-cdn.xitu.io/2019/6/5/16b277d32977f7d2?imageView2/0/w/1280/h/960/ignore-error/1)

相关文章推荐阅读：

* [我们自研的那些Devops工具]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FHwOi-ARTvvNjGTWrDmZIkQ )
* [Nginx的几个常用配置和技巧]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FJyUnN_OtQ2NtXcH0mtCJUg )