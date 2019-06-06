# OkHTTP、Retrofit 中文乱码解决方法 #

### 1. 乱码出现的原因是什么？ ###

出现乱码的根本原因是客户端、服务端两端编码格式不一致导致的。

### 2. 两端的编码格式一般是什么？ ###

客户端：多数情况下，客户端的编码格式是 UTF-8。
服务端：服务端会根据不同的请求方法使用不同的编码格式。如：请求方法为 POST 时，编码格式为 UTF-8；请求方法为 GET 时，编码格式为 ISO8859-1。

### 3. 如何解决乱码问题？ ###

当请求方法为 POST 时，客户端和服务端两边的编码格式一致，所以不存在乱码问题。因此此处着重看下如何解决当请求方法为 GET 时的乱码问题。

解决方法倒也简单，只不过需要客户端和服务端配合：

#### 3.1 客户端需要做什么？ ####

在向 URL 添加参数之前，先对目标参数进行两次 encode，如 UTF-8：

` String username = "xxx" ; username = URLEncoder.encode(username, "UTF-8" ); username = URLEncoder.encode(username, "UTF-8" ); String url = xxxxx.xxx + "?username=" + username + "&password=" + password; 复制代码`

#### 3.2 服务端需要做什么？ ####

服务器在收到数据之后，只需将数据进行一次跟客户端编码格式一样的 decode，如 UTF-8：

` String username = URLDecoder.decode(username, "UTF-8" ) 复制代码`

这样处理之后，两边就不会再出现乱码了。

### 4. 为什么对 URL 中的参数进行两次 encode 之后，就可以解决乱码问题了？ ###

#### 4.1 原理解析 ####

通过上面的分析可知，乱码产生的主要原因是客户端、服务器两边编码不一致造成的，即发送 GET 请求时，客户端使用的是 UTF-8 编码格式对 URL 中的参数进行编码，而服务器在接收数据的时候，使用的是 ISO8859-1（解析 POST 请求时，服务器使用的编码格式是 UTF-8 编码格式）编码格式对 URL 中的参数进行解码。

ISO8859-1 跟 ASCII 码一样，都是单字节编码，ISO8859-1 是从 ASCII 扩展而来的。ISO8859-1 将 ASCII 一个字节中剩余的最后一位用了起来，也就是说，它比 ASCII 多了 128 个字符。另外，因为 ISO8859-1 是从 ASCII 扩展而来的，所以，ISO8859-1 兼容 ASCII。

原数据：

` 极速 复制代码`

客户端第一次编码，URLDecoder.decode(username, "UTF-8") 编码之后：

` %E6%9E%81%E9%80%9F 复制代码`

客户端第二次编码，URLDecoder.decode(username, "UTF-8") 编码之后：

` %25E6%259E%2581%25E9%2580%259F 复制代码`

客户端发出的 URL：

` http://192.168.31.148:8080/OkHttpServer/login?username=%25E6%259E%2581%25E9%2580%259F&password=123456 复制代码`

服务器接收的 URL：

` http://192.168.31.148:8080/OkHttpServer/login?username=%25E6%259E%2581%25E9%2580%259F&password=123456 复制代码`

服务器第一次解码，服务器接收到 GET 请求之后，默认会用 ISO8859-1 编码格式解码，解码之后得到：

` http://192.168.31.148:8080/OkHttpServer/login?username=%E6%9E%81%E9%80%9F&password=123456 复制代码`

需要注意的是，服务器用 ISO8859-1 编码格式解码 URL 中的参数是自动完成的。
因为客户端第一次用 URLDecoder.decode(username, "UTF-8") 编码 URL 中参数之后，得到的是 ASCII 码，且 UTF-8 和 ISO8859-1 对 ASCII 的编码结果是一致的，所以，客户端第二次用 URLDecoder.decode(username, "UTF-8") 之后的结果可以直接用 ISO8859-1 编码格式解码。
由于服务器解码之后的 URL 中的参数是用 UTF-8 编码格式编码的，所以，此时需要服务器再用 UTF-8 编码格式解码一次。

服务器第二次解码，服务器用 UTF-8 编码格式解码之后得到：

` http://192.168.31.148:8080/OkHttpServer/login?username=极速&password=123456 复制代码`

#### 4.2 实际应用 ####

如果客户端程序员没有显式用 UTF-8 编码格式编码 URL 中的参数，服务端要如何处理才能获取到原数据？

首先，分析下如果客户端没有用 UTF-8 编码格式编码 URL 中的参数，程序是如何执行的：

网络请求框架会对 URL 中的参数进行一次 UTF-8 编码：

` URLDecoder.encode(username, "UTF-8" ) 复制代码`

服务器会对 URL 中的参数进行一次 ISO8859-1 编码：

` URLDecoder.decode(username, "ISO8859-1" ) 复制代码`

明白了执行流程之后，如何解决自然也就显而易见了：
先转回 ISO8859-1 解码（decode）之前的结果，再转会 UTF-8 编码（encode）之前的结果。

具体操作步骤：

` //1. 先转回 ISO8859-1 解码（decode）之前的结果 String temp = URLDecoder.encode(username, "ISO8859-1" )； //2. 再转会 UTF-8 编码（encode）之前的结果 temp = URLDecoder.decode(username, "UTF-8" ) 复制代码`

#### 4.3 为什么 URL 中的参数经 UTF-8 编码格式编码之后不能通过 ISO8859-1 编码格式直接解码呢？ ####

因为 URL 中的参数经 UTF-8 编码格式编码之后得到的结果在 ISO8859-1 字符集可能一样也可能根本表示不了，这也是为什么 ASCII 码经 UTF-8 编码格式编码之后的结果可以用 ISO8859-1 编码格式解码。如，在 Unicode 字符集中，第 20013 个字符是“中”，而在 ISO8859-1 字符集中，一共才有 256 个字符。字符“中”经 UTF-8 编码之后的结果再经 ISO8859-1 解码，无论如何也得不到正确答案的。