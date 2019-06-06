# Java Web 扫盲行动 #

## 前言 ##

这次分享讲一下 Java Web 相关的基础知识，主要就是 ` servlet` 部分的知识。涉及到的知识点比较的多，如果同学们来不及看，可以先收藏起来，有空的时候再慢慢看哦！下面我们步入正题。

## 1. HTTP 协议 ##

* ` 协议` 就是一套约定好的规则，只要我们遵循其中的规则就能很好的进行沟通与协作。 ` HTTP` 协议也一样，HTTP 协议严格规定了 HTTP 请求和 HTTP 响应的数据格式，只要 HTTP 服务器与客户程序之间的交换数据都遵守 HTTP 协议，双方就能看得懂对方发送的数据，从而顺利交流。
* ` HTTP` 协议位于 ` 应用层` ，建立在 ` TCP/IP` 协议的基础上。HTTP 协议使用可靠的 ` TCP` 连接，默认端口 ` 80` 端口。

### 1.1 HTTP 响应格式 ###

HTTP 响应也由3部分构成：

* HTTP 协议的版本、状态码和描述。
* 响应头（Response Header）。
* 响应正文（Response Content）。

以下是一些常见的状态码：

* ` 200` ： **表示服务器已经成功地处理了客户端发出的请求** 。
* ` 400` ：错误的请求，客户发送的HTTP请求不正确。
* ` 404` ：文件不存在。
* ` 405` ：服务器不支持客户的请求方式。
* ` 500` ：服务器内部错误。

### 1.2 正文部分的 MIME 类型 ###

HTTP 请求及响应的正文部分可以是任意格式的数据，如何保证接收方能看得懂发送方发送的 ` 正文数据` 呢？HTTP协议采用 MIME 协议来规范正文的数据格式。

遵循 MIME 协议的数据类型统称为 MIME 类型。在 HTTP 请求头和 HTTP 响应头中都有一个 ` Content-Type` 项，用来指定 ` 请求正文部分` 或 ` 响应正文` 部分的 MIME 类型。下标列出了常见的 MIME 类型与文件扩展名之间的对应关系。

+----------------------------------+--------------------------+
|            文件扩展名            |         MIME类型         |
+----------------------------------+--------------------------+
| 未知的数据类型或不可识别的扩展名 | content/unknown          |
| .bin、.exe、.o、.a、.z           | application/octet-stream |
| .pdf                             | application/pdf          |
| .zip                             | application/zip          |
| .tar                             | application/x-tar        |
| .gif                             | image/gif                |
| .jpg、.jpeg                      | image/jpeg               |
| .htm ,html                       | text/html                |
| .text .c .h .txt .java           | text/plain               |
| .mpg .mpeg                       | video/mpeg               |
| .xml                             | application/xml          |
+----------------------------------+--------------------------+

## 2. Web 服务器 ##

什么是Web服务器呢？ ` Web服务器` 是由专门的服务器开发商创建 ，用来发布和运行Web应用的。Web 服务器如何能动态执行由第三方创建的Web应用中的程序代码呢？所以我们需要一个 ` 中介方` 制定 ` Web应用` 与 ` Web 服务器` 进行协作的标准接口， ` Servlet` 就是其中最主要的一个接口。中介方规定：

* Web 服务器可以访问任意一个 Web 应用中实现 Servlet 接口的类。
* Web 应用中用于被 Web 服务器动态调用的程序代码位于 Servlet 接口的实现类中。

![3B49D892-27A6-4C31-AA06-066C45CF2BA3](https://user-gold-cdn.xitu.io/2019/4/26/16a572ae0e61eb8e?imageView2/0/w/1280/h/960/ignore-error/1)

Servlet 规范把能够发布和运行 ` JavaWeb` 应用的 ` Web服务器` 称为 ` Servlet 容器` ，它的最主要的特征是动态执行 JavaWeb 实现类中的程序代码。常见的 Servlet 容器有 Tomcat、 Jetty、WebLogic、WebSphere、JBoss 等。

### 2.1 Tomcat 服务器 ###

**Tomca 作为 Servlet 容器的基本功能** 如下图所示， ` Tomcat` 作为运行 Servlet 的容器，其基本功能是 ` 负责接收和解析来自客户的请求` ，同时把客户的请求 传送给相应的 ` Servlet` ，并把 Servlet 的 ` 响应结果` 返回给客户。

![B5915359-62D4-4A78-AC89-C6DEBA70980D](https://user-gold-cdn.xitu.io/2019/4/26/16a572ae13a17471?imageView2/0/w/1280/h/960/ignore-error/1)

Servlet 容器响应客户请求访问特定 Servlet 的流程如下：

* 客户发出要求访问特定 Servlet 的请求。
* Servlet 容器接收到客户请求，对其解析。
* Servlet 容器创建一个 ServletRequest 对象，在 ServletRequest 对象中包含了客户请求信息及其他关于客户的信息，如请求头、请求正文，以及客户机的IP地址等。
* Servlet 容器创建一个 ServletResponse 对象。
* Servlet 容器调用客户所请求的 Servlet 的 service() 服务方法，并且把 ServletRequest对象和 ServletResponse 对象作为参数传给该服务方法。
* Servlet 从 ServletRequest 对象中可获取客户的请求信息。
* Servlet 利用 ServletResponse 对象来生成响应结果。
* Servlet 容器把 Servlet 生成的响应结果发送给客户。 ![C566E1C3-8F21-48F2-95AC-5E811A2AA3DF](https://user-gold-cdn.xitu.io/2019/4/26/16a572af8a3d96e3?imageView2/0/w/1280/h/960/ignore-error/1)

## 3. JavaWeb 应用 ##

为了让 Servlet 容器能顺利地找到 JavaWeb 应用中的各个组件，Servlet 规范规定，JavaWeb 应用必须采用固定的目录结构。Servlet 规范还规定，JavaWeb 应用的配置信息存放在 WEB-INF/web.xml 文件中，Servlet 容器从该文件中读取配置信息。

### 3.1 JavaWeb 应用的目录结构 ###

假定开发一个名为 helloapp 的 JavaWeb 应用，首先，应该创建这个Web应用的目录结构，如下表所示：

+---------------------------+-------------------------------------------------------------+
|           目录            |                            描述                             |
+---------------------------+-------------------------------------------------------------+
| /helloapp                 | Web应用的根目录。                                           |
| /helloapp/WEB-INF         | 存放Web应用的配置文件web.xml                                |
| /helloapp/WEB-INF/classes | 存放各种.class文件，Servlet                                 |
|                           | 类的 .class 文件也放于此目录下                              |
| /helloapp/WEB-INF/lib     | 存放Web应用所需的各种JAR文件。例如，JDBC驱动程序的JAR文件。 |
+---------------------------+-------------------------------------------------------------+

可以看出 Servlet 容器并不关心你的源代码放在哪里，它只关心 .class 文件，因为加载类只需要用到 .class 文件。在 WEB-INF 目录的 classes及 lib 子目录下，都可以存放 Java 类文件。在运行时，Servlet 容器的类加载器先加载 classes 目录下的类，再加载 lib 目录下的 JAR 文件（Java类库的打包文件）中的类。因此，如果两个目录下存在同名的类，classes 目录下的类具有优先权。

## 4. Servlet 技术 ##

大佬终于出场了，大家掌声欢迎！本节主要介绍的就是 Servlet 中经常需要使用到的 "神器“。

### 4.1 Servlet 常用对象 ###

* ` 请求对象` （ServletRequest 和 HttpServletRequest）：Servlet 从该对象中获取来自客户端的请求信息。
* ` 响应对象` （ServletResponse 和 HttpServletResponse）：Servlet 通过该对象来生成响应结果。
* ` Servlet配置对象` （ServletConfig）：当容器初始化一个 Servlet 对象时，会向 Servlet 提供一个 ServletConfig 对象，Servlet 通过该对象来获取初始化参数信息及 ServletContext 对象。
* ` Servlet上下文对象` （ServletContext）：Servlet 通过该对象来访问容器为当前 Web 应用提供的各种资源。

### 4.2 Servlet API ###

Servlet API 主要由两个 Java 包组成： ` javax.servlet` 和 ` javax.servlet.http` 。

* ` javax.servlet` ：包中定义了 Servlet 接口及相关的通用接口和类；
* ` javax.servlet.http` 包中主要定义了与 HTTP 协议相关的 HTTPServlet 类、HTTPServletRequest 接口和 HTTPServletResponse 接口。

### 4.3 Servlet 接口 ###

Servlet API 的核心是 ` javax.servlet.Servlet` 接口，所有的 Servlet 类都必须实现这一接口。在 Servlet 接口中定义了5个方法，其中3个方法都由 Servlet 容器来调用，容器会在 Servlet的生命周期的不同阶段调用特定的方法。

* ` init(ServletConfig config)` ：负责初始化 Servlet 对象。容器在创建好 Servlet 对象后，就会调用该方法。
* ` service(ServletRequest req, ServletResponse res)` ：负责响应客户的请求，为客户提供相应服务。当容器接收到客户端要求访问特定 Servlet 对象的请求时，就会调用该 Servlet 对象的 service() 方法。
* ` destroy()` ：负责释放Servlet对象占用的资源。当 Servlet 对象结束生命周期时，容器会调用此方法。

Servlet 接口还定义了以下两个返回 Servlet 的相关信息的方法。JavaWeb 应用中的程序代码可以访问 Servlet 的这两个方法，从而获得 Servlet 的配置信息及其他相关信息。

* ` getServletConfig()` ：返回一个 ServletConfig 对象，在该对象中包含了 Servlet 的初始化参数信息。
* ` getServletInfo()` ：返回一个字符串，在该字符串中包含了Servlet的创建者、版本和版权等信息。

在 Servlet API 中， ` java.servlet.GenericServlet` 抽象类实现了 Servlet 接口，而 ` javax.servlet.http.HttpServlet` 抽象类是 ` GenericServlet` 类的子类。当用户开发自己的 Servlet 类时，可以选择扩展 ` GenericServlet` 类或者 ` HTTPServlet` 类。

### 4.4 GenericServlet 抽象类 ###

` GenericServlet` 抽象类为 Servlet 接口提供了通用实现，它与任何 ` 网络应用层协议无关` 。GenericServlet 类除了实现 Servlet 接口，还实现了 ServletConfig 接口和 Serializable 接口。

从GenericServlet 类的源代码可以看出，GenericServlet 类实现了 Servlet 接口中的 ` init(ServletConfig config)` 初始化方法。 ` GenericServlet` 类有一个 ` ServletConfig 类型的私有实例变量config` ，当 ` Sevlet容器` 调用 ` GenericServlet` 的 ` init(ServletConfig config)` 方法时，该方法使得私有实例变量 ` config` 引用由 ` 容器传入` 的 ServletConfig 对象，即使得 GenericServlet 对象与一个 ServletConfig 对象关联。

` public void init (ServletConfig config) throws ServletException { this.config = config; this.init(); } 复制代码`

GenericServlet类还自定义了一个不带参数的 init() 方法， ` init(ServletConfig config)` 方法会调用此方法。对于GenericServlet类的子类，如果希望覆盖父类的初始化行为，有以下两种办法：

* 覆盖父类的不带参数的init()方法：
` public void init () { // 子类具体的初始化行为 } 复制代码` * 覆盖父类的带参数的 ` init(ServletConfig config)` 方法。如果希望当前 Servlet 对象与 ServletConfig 对象关联，应该现在该方法中调用 super.init(config) 方法：
` public void init (ServletConfig config) { // 调用父类的init(config)方法 super.init(config); // 子类具体的初始化行为 } 复制代码`

GenericServlet 类没有实现 Servlet 接口中的 service() 方法。service() 方法是 GenericServlet 类中唯一的抽象方法，GenericServlet 类的具体子类必须实现该方法，从而为特定的客户请求提供具体的服务。

此外，GenericServlet 类实现了 ServletConfig 接口中的所有方法。因此，GenericServlet 类的子类可以直接调用在 ServletConfig 接口中定义的 getServletContext()、getInitParameter() 和 getInitParameterNames() 等方法。

GenericServlet 类实现了 Servlet 接口和 ServletConfig 接口。GenericServlet 类的主要身份是 Servlet，此外，它还运用 ` 装饰设计模式` ，为自己附加了 ServletConfig 装饰身份。在具体实现中，GenericServlet 类包装了一个 ServletConfig 接口的实例，通过该实例来实现 ServletConfig 接口中的方法。

### 4.5 HttpServlet 抽象类 ###

HttpServlet 类是 GenericServlet 类的子类。HttpServlet 类为 Servlet 接口提供了与 ` HTTP协议` 相关的通用实现，也就是说，HttpServlet 对象适合运行在客户端采用 HTTP 协议通信的 Servlet 容器或者 Web 服务器中。在开发 JavaWeb 应用时，自定义的 Servlet 类一般都扩展 HttpServlet 类。

HTTP 协议把客户请求分为 GET、POST、PUT 和 DELETE 等多种方式。HttpServlet 类针对每一种请求方式都提供了相应的服务方法，如doGet()、doPost()、doPut和doDelete()等方法。

从 HttpServlet 的源码可以看出，HttpServlet 类实现了 Servlet 接口中的 service(ServletRequest req,ServletResponse res) 方法，该方法实际上调用的是它的重载方法：

> 
> 
> 
> service(HttpServletRequest req, HttpServletResponse resp)
> 
> 

在以上重载 service() 方法中，首先调用 HttpServletRequest 类型的 req 参数的 ` getMethod()` 方法，从而获得客户端的请求方式，然后依据该请求方式来调用匹配的服务方法。如果为 GET 方法，则调用 doGet() 方法；如果为 POST 方式，则调用 doPost() 方法，依次类推。

### 4.6 ServletRequest接口 ###

在 Servlet 接口的 ` service(ServletRequest req, ServletResponse res)` 方法中有一个 ServletRequest 类型的参数。ServletRequest 类表示来自客户端的请求。当 Servlet 容器接收到客户端要求访问特定 Servlet 的请求时， ` 容器先解析客户端的原始请求数据，把它包装成一个ServletRequest 对象` 。当容器调用 Servlet 对象的 service() 方法时，就可以把 ServletRequest 对象作为参数传给 service() 方法。

ServletRequest 接口提供了一系列用于读取客户端的请求数据的方法：

* 

` getContentLength()` ：返回请求正文的长度。如果请求正文的长度未知，则返回-1。

* 

` getContentType()` ：获得请求正文的MIME类型。如果请求正文的类型未知，则返回null。

* 

` getInputStream()` ：返回用于读取请求正文的输入流。

* 

` getParameter(String name)` ：根据给定的请求参数名，返回来自客户端请求中的匹配的请求参数值。

* 

` getReader()` ：返回用于读取字符串形式的请求正文的BufferedReader对象。等等

此外，在 ServletRequest 接口中还定义了一组用于在请求范围内存取共享数据的方法：

* 

` getAttribute(String name,Object object)` ：在请求范围内保存一个属性，参数 name 表示属性名，参数 object 表示属性值。

* 

` getAttribute(String name)` ：根据 name 参数给定的属性名，返回请求范围内的匹配的属性值。

* 

` removeAttribute(String name)` ：从请求范围内删除一个属性。

### 4.7 HttpServletRequest 接口 ###

HttpServletRequest 接口是 ServletRequest 接口的子接口。HttpServlet 类的重载 service() 方法及 doGet() 和doPost() 等方法都有一个 HttpServletRequest 类型的参数。

HttpServletRequest 接口提供了用于读取HTTP请求中的相关信息的方法：

* 

` getContextPath()` ：返回客户端所请求访问的Web应用的URL入口。例如，客户端访问的URL为 ` http://localhost:8080/helloapp/info` ，那么该方法返回 ”/helloapp"

* 

` getCookies()` ：返回HTTP请求中的所有 Cookie。

* 

` getHeader(String name)` ：返回 HTTP 请求头部的特定项。

* 

` getHeaderNames()` ：返回一个 Enumeration 对象，它包含了 HTTP 请求头部的所有项目名。

* 

` getMethod()` ：返回 HTTP 请求方式。

* 

` getRequestURI()` ：返回HTTP请求的头部的第一行中的 URI。

* 

` getQueryString()` ：返回HTTP请求中的查询字符串，即URL中的 ` ?` 后面的内容。

### 4.8 ServletResponse接口 ###

在 Servlet 接口的 ` service(ServletRequest req, ServletResponse res)` 方法中有一个 ServletResponse 类型的参数。Servlet 通过 ServletResponse 对象来生成响应结果。当 Servlet 容器接收到客户端要求访问特定 Servlet 的请求时容器会创建一个 ServletResponse 对象，并把它作为参数传给 Servlet 的 service() 方法。

在 ServletResponse 接口中定义了一系列与生成响应结果相关的方法。

* 

` getOutputStream()` ：返回一个 ServletOutputStream 对象，Servlet 用它来输出二进制的正文数据。

* 

` getWriter()` ：返回一个PrintWriter对象，Servlet 用它来输出字符串像是的正文数据。

` ServletResponse` 中响应正文的默认 MIME 类型为 ` text/plain` ，即纯文本类型。而 ` HttpServletResponse` 中响应正文的默认 MIME 类型为 ` text/html` ，即HTML文档类型。

为了提高输出数据的效率，ServletOutputStream 和 PrintWriter 先把数据写到缓冲区内。当缓冲区内的数据被提交给客户后，ServletResponse 的 isCommitted() 方法返回 true。在以下几种情况下，缓冲区内的数据会被提交给客户，即数据被发送到客户端：

* 

当缓冲区内的数据 ` 已满` 时，ServletOutputStream 或 PrintWriter 会自动把缓冲区内的数据发送给客户端，并且清空缓冲区。

* 

Servlet 调用 ` ServletResponse` 对象的 ` flushBuffer()` 方法。

* 

Servlet 调用 ` ServletOutputStream` 或 ` PrintWriter` 对象的 ` flush()` 方法或 ` close()` 方法。

为了确保 ServletOutputStream 或 PrintWriter 输出的所有数据都会被提交给客户，比较安全的做法是在所有数据都输出完毕后，调用 ServletOutputStream 或 PrintWriter 的 ` close()` 方法。

在 Tomcat 的实现中，如果 Servlet 的 service() 方法没有调用 ServletOutputStream 或 PrintWriter 的 close() 方法，那么 ` Tomcat` 在调用完 Servlet 的 service() 方法后，会 ` 关闭` ServletOutputStream 或 PrintWriter，从而确保 Servlet 输出的所有数据被提交给客户。

值得注意的是，如果要设置响应正文的MIME类型和字符编码，必须 ` 先调用` ServletResponse 对象的 ` setContentType()` 和 ` setCharacterEncoding()` 方法，然后 ` 再调用` ServletResponse 的 getOutputStream() 或 getWriter() 方法，或者提交缓冲区内的正文数据。只有满足这样的操作顺序，所做的设置才能生效。

### 4.9 HttpServletResponse 接口 ###

HttpServletResponse 接口是 ServletResponse 的子接口，HttpServlet 类的重载 service() 方法及 doGet 和 doPost() 等方法都有一个 HttpServletResponse 类型的参数。

HttpServletResponse接口提供了与HTTP协议相关的一些方法，Servlet可通过这些方法来设置HTTP响应头或向客户端写Cookie。

* 

` addHeader(String name,String value)` ：向HTTP响应头中加入一项内容。

* 

` setHeader(String name,String msg)` ：设置HTTP响应头中的一项内容。如果在响应头中已经存在这项内容，那么原先所做的设置将被覆盖。

* 

` sendError(int sc)` ：向客户端发送一个代表特定错误的HTTP响应状态代码。

* 

` sendError(int sc,String msg)` ：向客户端发送一个代表特定错误的HTTP响应状态代码，并且发送具体的错误消息。

* 

` setStatus(int sc)` ：设置HTTP响应的状态代码。

* 

` addCookie(Cookie cookie)` ：向HTTP响应中加入一个Cookie。

以下3种方式都能设置 HTTP 响应正文的 MIME 类型及字符编码：

` // 方式一 response.setContentType(“text/html;charset=utf- 8 ”); // 方式二 response.setContentType(“text/html”); response.setCharacterEncoding(“utf- 8 ”); // 方式三 response.setHeader(“Content-type”,”text/html;charset=utf- 8 ”); 复制代码`

### 4.10 ServletConfig 接口 ###

Servlet接口的 ` init(ServletConfig congfig)` 方法有一个 ServletConfig 类型的参数。当 Servlet 容器初始化一个 Servlet 对象时，会为这个 Servlet 对象创建一个 ServletConfig 对象。在 ` ServletConfig` 对象中包含了 Servlet 的 ` 初始化参数信息` ，此外，ServletConfig 对象还与当前Web应用的 ServletContext 对象关联。

ServletConfig 接口中定义了以下方法。

* 

getInitParameter(String name)：根据给定的初始化参数名，返回匹配的初始化参数值。

* 

getInitParameterNames()：返回一个Enumeration对象，里面包含了所有的初始化参数名。

* 

getServletContext()：返回一个ServletContext对象。

每个初始化参数包括一对参数名和参数值。在web.xml文件中配置一个Servlet时，可以通过 元素来设置初始化参数。 元素的 子元素设定参数名， 子元素设定参数值。

HttpServlet 类继承 GenericServle t类，而 GenericServlet 实现了 ServletConfig 接口，因此在 HttpServlet 或GenericServlet 类及子类中都可以直接调用 ServletConfig 接口中的方法。

### 4.11 ServletContext 接口 ###

` ServletContext` 是 ` Servlet` 与 ` Servlet 容器` 之间 ` 直接通信` 的接口。Servlet容器在启动一个Web应用时，会为它创建一个ServletContext对象。每个Web应用都有 ` 唯一` 的ServletContext对象，可以把 ` ServletContext` 对象形象地理解为Web应用的总管家， ` 同一个Web应用中的所有Servlet对象都共享一个总管家` (敲黑板了，所有Servlet对象共享一个 ServletContext 对象) ，Servlet对象们可通过这个总管家来访问容器中的各种资源。

Servlet 容器在 ` 启动一个Web应用时` ，会为它创建 ` 唯一` 的 ServletContext 对象。当 Servlet 容器终止一个Web应用时，就会销毁它的ServletContext对象。由此可见，ServletContext 对象与Web应用具有 ` 相同的生命周期` 。

下面展示 ServletContext 接口经常使用的几个方法：

* 

用于 ` Web应用范围内` 存取 ` 共享数据` 的方法。

* 

` setAttribute(String name,Object object)` ：把一个Java对象与一个属性名绑定，并把它存入到ServletContext中。

* 

` getAttribute(String name)` ：根据参数给定的属性名，返回一个Object类型的对象，它表示ServletContext中与属性名匹配的属性值。

* 

` getAttributeNames()` ：返回一个Enumeration对象，该对象包含了所有存放在ServletContext中的属性名。

* 

` removeAttribute(String name)` ：根据参数指定的属性名，从ServletContext中删除匹配的属性。

* 

访问当前Web应用的资源。

* 

` getContextPath()` ：返回当前Web应用的URL入口。

* 

` getInitParameter(String name)` ：根据给定的参数名，返回Web应用范围内的匹配的初始化参数值。在web.xml文件中，直接在 根元素下定义的 元素表示应用范围内的初始化参数。

* 

` getInitParameterNames()` ：返回一个Enumeration对象，它包含了Web应用范围内所有初始化参数名。

* 

` getServletContextName()` ：返回Web应用的名字，即web.xml文件中<display-name>元素的值。

* 

` getRequestDispatcher(String path)` ：返回一个用于向其他Web组件转发请求的RequestDispatcher对象。

在 ServletConfig 接口中定义了 getServletContext() 方法。HttpServlet 类继承 GenericServlet类，而GenericServlet 类实现了 ServletConfig 接口，因此在 HttpServlet 类或 GenericServlet 类及子类中都可以直接调用 getServletContext() 方法，从而得到当前Web应用的 ServletContext 对象。

## 5. JavaWeb 应用的生命周期 ##

JavaWeb 应用的生命周期是由 Servlet 容器来控制的。归纳起来，JavaWeb 应用的生命周期包括 3 个阶段。

* 

启动阶段：加载Web应用的有关数据，创建ServletContext对象 ，对Filter和一些Servlet进行初始化。

* 

运行时阶段：为客户端服务。

* 

终止阶段：释放Web应用所占用的各种资源。

### 5.1 启动阶段 ###

Servlet 容器在启动 JavaWeb 应用时，会完成以下操作：

* 

把 ` web.xml` 文件中的数据加载到内存中。

* 

为 JavaWeb 应用创建一个 ` ServletContext` 对象。

* 

对所有的 ` Filter` 进行初始化。

* 

对那些需要在 Web 应用 ` 启动时就被初始化` 的 ` Servlet` 进行初始化。

### 5.2 运行时阶段 ###

这是 JavaWeb 应用最主要的生命阶段。在这个阶段，它的所有 Servlet 都处于待命状态，随时可以响应客户端的特定请求，提供相应的服务。加入客户端请求的 Servlet 还不存在，Servlet 容器会先初始化Servlet ，然后再调用它的 service() 服务。

### 5.3 终止阶段 ###

Servlet 容器在终止 JavaWeb 应用时，会完成以下操作：

* 

销毁 JavaWeb 应用中所有处于运行时状态的 Servlet。

* 

销毁 JavaWeb 应用中所有处于运行时状态的 Filter。

* 

销毁所有与 JavaWeb 应用相关的对象，如果 ServletContext 对象等，并且释放 Web 应用所占用的相关资源。

## 6. Servlet 的生命周期 ##

JavaWeb 应用的生命周期由 Servlet 容器来控制，而 Servlet 作为 JavaWeb 应用的最核心的组件，其生命周期也由 Servlet 容器来控制。Servlet 的生命周期可以分为3个阶段：初始化阶段、运行时阶段和销毁阶段。在javax.servlet.servlet接口中定义了3个方法：init()、service() 和 destroy()，它们将分别在 Servlet 的不同阶段被 Servlet 容器调用。

### 6.1 初始化阶段 ###

Servlet 的初始化阶段包括 4 个步骤：

* 

Servlet 容器加载 Servlet 类，把它的 .class 文件中的数据读入到内存中。

* 

Servlet 容器创建 ServletConfig 对象。 ` ServletConfig` 对象包含了 ` 特定` Servlet 的 ` 初始化配置信息` ，如 Servlet 的初始化参数。此外， Servlet 容器还会使得 ServletConfig 对象与当前Web应用的 ` ServletContext` 对象 ` 关联` 。

* 

Servlet 容器创建 Servlet 对象。

* 

Servlet 容器调用 Servlet 对象的 init(ServletConfig config) 方法。

以上初始化步骤创建了 Servlet 对象和 ServletConfig 对象，并且 Servlet 对象与 ServletConfig 对象关联，而 ServletConfig 对象又与当前 Web 应用的 ServletContext 对象关联。当 Servlet 容器初始化完 Servlet 后，Servlet 对象只要通过 getServletContext() 方法就能得到当前Web应用的 ServletContext 对象。

在下列情况之一，Servlet会进入初始化阶段。

* 

当前Web应用处于运行时阶段，特定 Servlet 被客户端首次请求访问。多数 Servlet 都会在这种情况下被 Servlet 容器初始化阶段。

* 

如果在 web.xml 文件中为一个 Servlet 设置了 元素，那么当 Servlet 容器启动 Servlet 所属的Web应用时，就会初始化这个Servlet。假如 Servlet1 和 Servlet2 的 的值分别为1和2，因此Servlet容器启动当前Web应用时，Servlet1第一个初始化，Servlet2被第二个初始化。而没有配置 元素的Servlet，当Servlet容器启动当前Web应用时将不会被初始化，只有当客户端首次请求访问该Servlet时，它才会被初始化。

* 

当Web应用被重新启动时，Web应用中的所有Servlet都会在特定的时刻被重新初始化。

### 6.2 运行时阶段 ###

这是 Servlet 的声明周期中的最重要阶段。在这个阶段，Servlet 可以随时响应客户端的请求。当 Servlet 容器接收到要求访问特定 Servlet 的客户请求时，Servlet 容器会创建针对于这个请求的 ServletRequest 对象和 ServletResponse 对象，然后调用相应 Servlet 对象的 service() 方法。service() 方法从 ServletRequest 对象中获得客户请求信息并处理该请求，在通过 ServletResponse 对象生成响应结果。

当 Servlet 容器把 Servlet 生成的 ` 响应结果发送` 给了客户，Servlet 容器就会 ` 销毁` ServletRequest 对象和 ServletResponse 对象。

### 6.3 销毁阶段 ###

` 当Web应用被终止时` ，Servlet 容器会先调用Web应用中 ` 所有Servlet` 对象的 ` destroy()` 方法，然后 ` 再销毁` 这些Servlet对象。在destroy()方法的实现中，可以释放Servlet所占用的资源。

此外，容器还会销毁与对象关联的 ` ServletConfig` 对象。

## 7. 使用 ServletContextListener 监听器 ##

在 Servlet API 中有一个 ` ServletContextListener` 接口，它能够监听 ServletContext 对象的生命周期，实际上就是监听 Web应用 的生命周期。

当Servlet容器启动或终止Web应用时，它能够监听 ServletContextEvent 事件，该事件由ServletContextListener来处理。在ServletContextListener接口中定义了处理ServletContextEvent事件的两个方法。

* 

` contextInitialized(ServletContextEvent sec)` ：当Servlet容器 ` 启动` Web应用时调用该方法。在调用完该方法 ` 之后` ，容器再对 ` Filter初始化` ，并且对那些在Web应用启动时就需要被初始化的 Servlet 进行初始化。

* 

` contextDestroyed(ServletContextEvent sec)` ：当Servlet容器 ` 终止` Web应用调用该方法。在调用该方法 ` 之前` ，容器会先 ` 销毁` 所有的 ` Servlet` 和 ` Filter` 过滤器。

可以看出，在Web应用的生命周期中，ServletContext 对象最早被创建，最晚被销毁。

用户自定义的 ServletContextListener 监听器只有先向 Servlet 容器注册，Servlet 容器在启动或终止 Web 应用时，才会调用该监听器的相关方法。在 web.xml 为文件中， 元素用于向容器注册监听器

` < listener > < listener-class > LinstnerClass </ listener-class > </ listener > 复制代码`

## 8. Cookie ##

Cookie的英文原意是“点心”，它是在客户端访问Web服务器时， ` 服务器` 在 ` 客户端硬盘上存放的信息` 。当客户端 ` 首次` 访问服务器时，服务器现在客户端存放包含该客户的相关信息的Cookie，以后客户端每次请求访问服务器时，都会在 HTTP 请求数据中包含 Cookie，服务器解析 HTTP 请求中的 Cookie，就能由此获得关于用户的相关信息。

Cookie的运行机制是由HTTP协议规定的，多数 Web 服务器和浏览器都支持 Cookie。Web服务器为了支持Cookie，需要具备以下功能。

* 

在HTTP响应结果中添加Cookie数据。

* 

解析HTTP请求中的Cookie数据。

浏览器为了支持Cookie，需要具备以下功能。

* 

解析HTTP响应结果中的Cookie数据。

* 

把Cookie数据保存到本地硬盘。读取本地硬盘上的Cookie数据，把它添加到HTTP请求中。

Tomcat 作为Web服务器，对Cookie提供了良好的支持。那么，运行在Tomcat中的Servlet该如何访问Cookie呢？Java Servlet API为Servlet访问Cookie提供了简单易用的接口，Cookie 用 ` javax.servlet.http.Cookie` 类来表示，每个 Cookie 对象包含一个 Cookie 名字和 Cookie 值。

下面代码创建了一个 Cookie 对象，然后调用 HttpServletResponse 的 ` addCookie()` 方法，把 Cookie 添加到 HTTP 响应结果中：

` Cookie theCookie = new Cookie(“username”,”Tom”); response.addCookie(theCookie); 复制代码`

如果Servlet想读取来自客户端的Cookie，那么可以通过以下方式从HTTP请求中取得所有的Cookie：

` Cookie[] cookies = request.getCookies(); 复制代码`

对于每个 Cookie 对象，可调用 ` getName()` 方法来获得 Cookie 的名字，调用 ` getValue()` 方法来获得 Cookie 的值。

当 Servlet 向客户端写 Cookie 时，还可以通过 Cookie 类的 ` setMaxAge(int expiry)` 方法来设置 Cookie 的有效期。参数 expiry 以秒为单位，它具有以下含义：

* 

如果 expiry 大于零，就指示浏览器在客户端硬盘上保存 Cookie 的时间为 expiry 秒，有效期内，其他浏览器也能访问这个 Cookie。

* 

如果 expiry 等于零，就指示浏览器删除当前 Cookie。

* 

如果 expiry 小于零，就指示浏览器不要把 Cookie 保存到客户端硬盘。Cookie 仅仅存在于当前浏览器进程中，当浏览器进程关闭，Cookie 也就消失。

Cookie默认的有效期为-1。

服务器对客户端进行读写 Cookie 操作，会给客户端带来安全隐患。服务器可能会向客户端发送包含恶意代码的Cookie 数据，此外，服务器可能会依据客户端的 Cookie 来窃取用户的保密信息。因此出于安全起见，多数浏览器可以设置是否启用 Cookie。

## 9. Session ##

当客户端访问 Web 应用时，在许多情况下，Web 服务器必须能够跟踪客户的状态。

Web服务器跟踪用户的状态通常有4种方法：

* 

在HTML表单中加入隐藏字段，它包含用于跟踪用户状态的数据。

* 

重写URL，使它包含用于跟踪客户状态的数据。

* 

用Cookie来传送用于跟踪客户状态的数据。

* 

使用会话（Session）机制 。

### 9.1 会话简介 ###

HTTP 是无状态的协议。在Web开发领域，会话机制是用于跟踪客户状态的普遍解决方案。会话指的是在一段时间内，单个客户与Web应用的一连串相关的交互过程。在一个会话中，客户可能多次请求访问Web应用的同一个网页，也有可能请求访问同一个Web应用中的多个网页。

Servlet 规范制定了基于Java的会话的具体运作机制。在Servlet API中定义了代表会话的 ` javax.servlet.http.HttpSession` 接口，Servlet容器必须实现这一接口。当一个会话开始时，Servlet容器将创建一个 HttpSession 对象，在该对象中可以存放表示客户状态的信息。Servlet容器为每个 HttpSession 对象分配一个唯一标志符，符号位 ` Session ID` 。

**key** ：客户端 Cookie 只负责存 Session ID，而 Session 对象是存储在服务器上的。

下面以一个名为 bookstore 应用为例，介绍会话的运作流程：

* 

一个浏览器进程第一次请求访问bookstore应用中的任意一个支持会话的网页，Servlet 容器试图寻找 HTTP 请求中表示 Session ID 的 Cookie ，由于还不存在这样的 Cookie ，因此就认为一个新的会话开始了，于是创建一个 HttpSession 对象，为它分配唯一的 Session ID ，然后把 Session ID 作为 Cookie 添加到 HTTP 响应结果中。当浏览器接收到 HTTP 响应结果后，会把其中表示 Session ID 的 Cookie 保存在客户端。

* 

浏览器进程继续请求访问 bookstore 应用中的任意一个支持会话的网页，在本次 HTTP 请求中会包含表示Session ID 的 Cookie。Servlet 容器试图寻找 HTTP 请求中表示 Session ID 的 Cookie，由于能得到这样的 Cookie 。因此认为本次请求已经处于一个会话中了，Servlet容器不再创建新的 HttpSession 对象，而是从 Cookie 中获取 Session ID ，然后根据 Session ID 找到内存中对应的 HttpSession 对象。

* 

浏览器进程重复步骤二，直到当前会话被销毁，HttpSession对象就会结束生命周期。

表示Session ID的Cookie的有效期为-1，这意味着该Cookie也就消失，本次会话也会结束。在两个浏览器进程中显式的Session ID的值不一样，因为两个浏览器进程分别对应不同的会话，而每个会话都有唯一的Session ID。

### 9.2 HttpSession 的生命周期及会话范围 ###

会话范围是指浏览器端与一个Web应用进行一次会话的过程。在具体实现上，会话范围与 HttpSession 对象的生命周期对应。因此，Web组件只要共享同一个 HttpSession 对象，也就能共享会话范围内的共享数据。

HttpSession接口中的方法描述如下表，Web应用中的JSP或Servlet组件可通过这些方法来访问会话。

+--------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
|              方法              |                                                                                                                        描述                                                                                                                         |
+--------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| getId()                        | 返回 Session ID。                                                                                                                                                                                                                                   |
| invalidate()                   | 销毁当前的会话，Servlet容器会释放HttpSession对象占用的资源。                                                                                                                                                                                        |
| setAttribute(String name,      | 将一对name/value属性保存在HttpSession对象中。                                                                                                                                                                                                       |
| Object value)                  |                                                                                                                                                                                                                                                     |
| getAttribute(String name)      | 根据name参数返回保存在HttpSession对象中的属性值。                                                                                                                                                                                                   |
| getAttributeNames()            | 以数组的方式返回HttpSession对象中所有属性名。                                                                                                                                                                                                       |
| removeAttribute(String name)   | 从HttpSession对象中删除name参数的指定属性。                                                                                                                                                                                                         |
| isNew()                        | 判断是否是新创建的会话                                                                                                                                                                                                                              |
| setMaxInactiveInterval(int     | 设定一个会话可以处于不活动状态的最长时间，以秒为单位。如果超过这个时间，Servlet容器会自动销毁会话。如果把参数interval设置为负数，表示不限制会话处于不活动状态的时间，即会话永远不会过期。Tomcat为会话设定的默认的保持不活动状态的最长时间为1800秒。 |
| interval)                      |                                                                                                                                                                                                                                                     |
| getMaxInactiveInterval()       | 读取当前会话可以处于不活动状态的最长时间。                                                                                                                                                                                                          |
+--------------------------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+

在以下情况下，会开始一个新的会话，即Servlet容器会创建一个新的 HttpSession 对象：

* 

一个浏览器进程第一访问Web应用中的支持会话的任意一个网页。

* 

当浏览器进程与Web应用的一次会话已经被销毁后，浏览器进程再次访问Web应用中的支持会话的任意一个网页。

在以下情况下，会话被销毁，即Servlet容器使HttpSession对象结束生命周期，并且存放在会话范围内的共享数据也都被销毁：

* 

浏览器进程终止。

* 

服务器端执行 HttpSession 对象的 invalidate() 方法。

* 

会话过期。

当Tomcat中的Web应用被终止时，它的会话不会被销毁，而是被Tomcat持久化到永久性存储设备中，当Web应用重启后，Tomcat会重新加载这些会话。

当一个会话开始后，如果浏览器进程突然关闭，Servlet容器端无法立即知道浏览器进程已经被关闭，因此Servlet容器端的HttpSession对象不会立即结束生命周期。不过，当浏览器进程关闭后，这个会话就进入不活动状态，等到超过了 ` setMaxInactiveInterval(int interval)` 方法设置的时间，会话就会因为过期而被 Servlet 容器销毁。

**key** ：是不是属于同一个 session，看 sessionId 是否相同就可以了。而 sessionId 又保存在浏览器的 cookie 中，显然不同浏览器保存的cookie是不一样的。

### 9.3 获取 Session ###

可以通过 HttpServletRequest 对象来获得 HttpSession 对象。

在 HttpServletRequest 接口中提供了两个与会话有关的方法：

* 

` getSession()` ：是的当前 HttpServlet 支持会话。假如会话已经存在，就返回相应的 HttpSession 对象，否则就创建一个新会话，并返回新建的 HttpSession 对象。该方法等价于调用 HttpServletRequest 的 ` getSession(true)` 方法。

* 

` getSession(boolean create)` ：如果参数 create 为 ` true` ，等价于调用 HttpServletRequest 的 getSession() 方法；如果参数 create 为 false ，那么假如会话已经存在，就返回响应的HttpSession对象，否则就返回 null。

**key** ：一个常见的误解是以为s ession 在有客户端访问时就被创建，然而事实是直到某 server 端程序调用HttpServletRequest.getSession(true)这样的语句时才被创建。

## 10. session 与 cookie 的区别 ##

* 

cookie数据存放在客户的浏览器上，session数据放在服务器上。

* 

cookie不是很安全，别人可以分析存放在本地的COOKIE并进行COOKIE欺骗，考虑到安全应当使用session。

* 

session会在一定时间内保存在服务器上。当访问增多，会比较占用你服务器的性能，考虑到减轻服务器性能方面，应当使用COOKIE。

* 

单个cookie保存的数据不能超过4K，很多浏览器都限制一个站点最多保存20个cookie。

## 11. 请求转发 ##

转发可以将请求转发给同一Web应用的组件。

` // 把请求转发给OutputServlet ServletContext context = getServletContext(); RequestDispatcher dispatcher = context.getRequestDispatcher( "/OutputServlet" ); // ok // RequestDispatcher dispatcher = context.getRequestDispatcher("outputServlet");//worng // RequestDispatcher dispatcher = req.getRequestDispatcher("outputServlet");//ok PrintWriter out = res.getWriter(); out.println( "Output from CheckServlet before forwading request." ); System.out.println( "Output from CheckServlet before forwading request." ); // throw IllegalArgumentException:Cannot forward after response has been committed //out.close(); dispatcher.forward(req, res); out.println( "Output from CheckServlet after forwading request." ); System.out.println( "Output from CheckServlet after forwading request." ); 复制代码` ` public class OutputServlet extends GenericServlet { @Override public void service (ServletRequest req, ServletResponse res) throws ServletException, IOException { // 读取CheckServlet存放在请求范围内的消息 String message = (String) req.getAttribute( "msg" ); PrintWriter out = res.getWriter(); out.println(message); out.close(); } } 复制代码`

控制台打印结果：

> 
> 
> 
> Output from CheckServlet before forwading request.
> 
> 
> 
> Output from CheckServlet after forwading request.
> 
> 

以上 ` dispatcher.forward(request,response)` 方法的处理流程如下：

* 

清空用于存放响应正文数据的缓冲区。

* 

如果目标组件为 Servlet 或 JSP，就调用它们的 service() 方法，把该方法产生的响应结果发送到客户端；如果目标组件为文件系统中的静态 HTML 文档，就读取文档中的数据并把它发送到客户端。

从 ` dispatcher.forward(request,response)` 方法的处理流程可以看出，请求转发具有以下特点：

* 

由于 forward() 方法先清空用于存放响应正文数据的缓冲区，因此 Servlet 源组件生成的响应结果不会被发送到客户端，只有目标组件生成的响应结果才会被发送到客户端。

* 

如果源组件在进行请求转发 ` 之前` ，已经提交了响应结果（例如调用ServletResponse的flushBuffer()方法，或者调用与ServletResponse关联的输出流的close()方法），那么 forward() 方法会抛出 ` lllegalStateException` 。为了避免该异常，不应该在源组件中提交响应结果。

* 

在 Servlet 源组件中调用 ` dispatcher.forward(request,response)` 方法之后的代码也会被执行。同理调用out.close() 之后的代码也会执行。只是 out.close() 之后的代码不会再通过 response 返回到客户端。

## 12. 包含 ##

下面代码把 header.html 的内容，GreetServlet 生成的响应内容以及 foot.html 的内容都包含到自己的响应结果中。也就是说，下面的Servlet返回给客户的HTML文档是由自身、header.htm、GreetServlet，以及foot.htm共同产生的。

` ServletContext context = getServletContext(); RequestDispatcher headDispatcher = context.getRequestDispatcher( "/header.htm" ); RequestDispatcher greetDispatcher = context.getRequestDispatcher( "/greet" ); RequestDispatcher footDispatcher = context.getRequestDispatcher( "/footer.htm" ); headDispatcher.include(request, response); greetDispatcher.include(request, response); footDispatcher.include(request, response); 复制代码`

RequestDispatcher对象的include()方法的处理流程如下。

* 

如果目标组件为Servlet或JSP，就调用它们的相应的service()方法，把该方法产生的响应正文添加到源组件的响应结果中；如果目标组件为HTML文档，就直接把文档的内容添加到源组件的响应结果中。

* 

返回到源组件的服务方法中，继续执行后续代码块。

包含与请求转发相比，前者有以下特点：

* 

源组件与被包含的目标组件的输出数据都会被添加到响应结果中。

* 

在目标组件中对响应状态代码或者响应头所做的修改都会被忽略。

当源组件和目标组件之间为请求转发关系或者包含关系时，对于每一次客户请求，它们都共享同一个ServletRequest对象及ServletResponse对象，因此源组件和目标组件能共享请求范围内的共享数据。

## 13. 重定向 ##

HTTP 协议规定了一种重定向机制，重定向的运作流程如下：

* 

用户在浏览器输入特定URL，请求访问服务器端的某个组件。

* 

服务器端的组件返回一个状态码为302的响应结果，响应结果的含义为：让浏览器端再请求访问另一个Web组件。在响应结果中提供了另一个Web组件的URL。另一个Web组件有可能在同一个Web服务器上，也可能不在同一个Web服务器上。

* 

当浏览器端接收到这种响应结果后，再立即自动请求访问另一个Web组件。

* 

浏览器端接收到来自另一个Web组件的响应结果。

在Java Servlet API中，HttpServletResponse接口的 ` sendRedirect(String location)` 方法用于重定向。

重定向之后，导航栏的 URL 会变成目标组件的 URL 地址。 ` response.sendRedirect(String location)` 方法具有以下特点：

* 

Servlet源组件生成的响应结果不会被发送到客户端。 ` response.sendRedirect(String location)` 方法一律返回状态码为302的响应结果，浏览器接收到这种响应结果后，再立即自动请求访问重定向的目标Web组件，客户端最后接收到的是目标Web组件的响应结果。

* 

如果源组件在进行重定向之前，已经提交了响应结果，那么 sendRedirect() 方法会抛出异常。为了避免该异常，不应该在源组件中提交响应结果。

* 

在Servlet源组件中调用 ` response.sendRedirect(String location)` 方法之后的代码也会被执行。

* 

源组件和目标组件不共享同一个 ServletRequest 对象，因此不共享请求范围内的共享数据。

* 

对于 ` response.sendRedirect(String location)` 方法中的参数location，如果以”/“开头，表示相对于当前服务器根路径的URL，如果以”http://“开头，表示一个完整的URL。

* 

目标组件不必是同一个服务器上的同一个Web应用中的组件，它可以是Internet上的任意一个有效的网页。

sendRedirect() 方法是在 HttpServletResponse 接口中定义的，而在 ServletResponse 接口中没有 sendRedirect() 方法，因为重定向机制是由 HTTP 协议规定的。

## 14. 转发与重定向的区别 ##

* 

forward 方法只能转发给同一个web站点的资源，而 sendRedirect 方法还可以定位到同一个web站点的其他应用。

* 

forward 转发后，浏览器 URL 地址不变，sendRedirect 重定向后，浏览器url地址变为目的 URL 地址。

* 

forward 转发的过程，在一个 servlet 中调用了同一个web应用的另一个 servlet 的 service() 方法，所以还是属于一次请求响应。sendRedirect，浏览器先向目的Servlet发送一次请求，Servlet 看到 sendRedirect 将目的 URL 返回到浏览器，浏览器再去请求目的 URL ,目的 URL 再返回response到浏览器。浏览器和服务器两次请求响应。

* 

forward 方法的调用者与被调用者之间共享 Request 和 Response。 sendRedirect 方法由于两次浏览器服务器请求，所以有两个 Request 和 Response。

## 15. 过滤器 ##

各个 Web 组件中的 ` 相同操作` 可以放到一个 ` 过滤器` 中来完成，这样就能减少重复编码。过滤器能够对一部分客户请求先进行预处理操作，然后再把请求转发给响应的 Web 组件，等到 Web 组件生产了响应结果后，过滤器还能对响应结果进行检查和修改，然后再把修改后的响应结果发送给客户。

过滤器负责过滤的 Web 组件可以是 Servlet、JSP 或 HTML 文件，过滤器的过滤过程如下图所示。

![1](https://user-gold-cdn.xitu.io/2019/4/26/16a572ae12b63e6b?imageView2/0/w/1280/h/960/ignore-error/1)

### 15.1 创建过滤器 ###

所有自定义的过滤器都必须实现 ` javax.servlet.Filter` 接口，这个接口含以下 3 个必须实现的方法。

* 

` init(FilterConfig filterConfig)` ：这是过滤器的初始化方法。在 ` Web 应用启动时` ， Servlet 容器先创建包含了过滤器配置信息的 ` FilterConfig` 对象，然后创建 ` Filter` 对象，接着调用 Filter 对象的 ` init(FilterConfig filterConfig)` 方法，在这个方法中可通过 ` config` 参数来读取 ` web.xml` 文件中为过滤器配置的初始化参数。

* 

` doFilter(ServletRequest request, ServletResponse response,FilterChain chain)` ：这个方法完成实际的过滤操作。当客户请求访问的 URL 与为过滤器映射的 URL 匹配时， Servlet 容器将先调用过滤器的 ` doFilter()` 方法。 ` FilterChain` 参数用于访问 ` 后续过滤器` 或者 ` Web组件` 。

* 

` destroy()` ：Servlet 容器在销毁过滤器对象前调用该方法，在这个方法中可以释放过滤器占用的资源。

### 15.2 Filter 中的 doFilter()方法 ###

doFilter()是整个过滤器中最为关键的一个方法，servlet 容器调用该方法完成实际的过滤操作

* 

该方法接收两个参数，分别为 request、response 参数，所以我们可以对 request 进行一些预处理之后再继续传递下去，也可以对返回的 response 结果进行一些额外的操作，再返回到客户端（前提是response没有被 close）。

* 

在 ` Filter.doFilter()` 方法中不能直接调用 Servlet 的 ` service` 方法，而是调用 ` FilterChain.doFilter` 方法来激活目标 Servlet 的 ` service` 方法，FilterChain 对象时通过 ` Filter.doFilter` 方法的参数传递进来的。

* 

在一个 Web 应用程序中可以注册 ` 多个` ` Filter` 程序，如果有多个 ` Filter` 程序都可以对某个 ` Servlet` 程序的访问过程进行拦截，当针对该 ` Servlet` 的访问请求到达时，Web 容器将把这多个 ` Filter` 程序组合成一个 ` Filter` 链（也叫过滤器链）。

* 

Filter 链中的各个 Filter 的拦截顺序与它们在 web.xml 文件中的映射顺序一致，上一个 Filter.doFilter 方法中调用 FilterChain.doFilter 方法将激活下一个 Filter的doFilter 方法，最后一个 Filter.doFilter 方法中调用的 FilterChain.doFilter 方法将激活目标 Servlet的service 方法。

* 

只要 Filter 链中任意一个 Filter 没有调用 FilterChain.doFilter 方法，则目标 Servlet 的 service 方法都不会被执行。

### 15.3 实例 ###

下面通过一个简单的例子演示基本的过滤器使用。

**定义过滤器**

该过滤器主要的功能就是在请求之前打印一下初始化参数，请求之后拼接一下返回结果：

` public class LogFilter implements Filter { private String ip; public void init (FilterConfig filterConfig) throws ServletException { System.out.println( "执行LogFilter初始化方法" ); // 获取Filter初始化参数 this.ip = filterConfig.getInitParameter( "ip" ); } public void doFilter (ServletRequest request, ServletResponse response, FilterChain chain) throws IOException, ServletException { System.out.println( "正在执行 LogFilter.doFilter()方法，LogFilter初始化参数ip=" + ip); System.out.println( "调用LogFilter中的chain.doFilter()之前" ); chain.doFilter(request, response); System.out.println( "调用LogFilter chain.doFilter()之后" ); PrintWriter out = response.getWriter(); // 拼接返回结果 out.println( "This msg is from LogFilter" ); out.flush(); } public void destroy () { System.out.println( "正在执行 LogFilter 的销毁方法" ); } } 复制代码`

**TestFilter servlet**

` public class TestFilter extends HttpServlet { protected void doGet (HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException { System.out.println( "正在执行 TestFilter 的 doGet() 方法" ); PrintWriter out = response.getWriter(); out.println( "This msg is from HelloServlet doGet()" ); // 如果执行 close 操作，后续过滤器中对 response 的操作 将无法返回到客户端 // out.close(); out.flush(); } } 复制代码`

**在 web.xml 中配置过滤器**

` <filter> <filter-name>LogFilter</filter-name> <filter-class>controller.LogFilter</filter-class> <init-param> <param-name>ip</param-name> <param-value>127.0.0.1</param-value> </init-param> </filter> <filter-mapping> <filter-name>LogFilter</filter-name> <url-pattern>/ test Filter</url-pattern> </filter-mapping> 复制代码`

控制台打印结果如下：

> 
> 
> 
> 正在执行 LogFilter.doFilter()方法，LogFilter初始化参数ip=127.0.0.1
> 
> 
> 
> 调用LogFilter中的chain.doFilter()之前
> 
> 
> 
> 正在执行 TestFilter 的 doGet() 方法
> 
> 
> 
> 调用LogFilter chain.doFilter()之后
> 
> 

http响应正文如下：

> 
> 
> 
> This msg is from HelloServlet doGet()
> 
> 
> 
> This msg is from LogFilter
> 
> 

### 15.4 串联过滤器 ###

多个过滤器可以串联起来协同工作，Servlet 容器将根据它们在 web.xml 中定义的先后顺序，一次调用它们的doFilter() 方法。假定有两个过滤器串联起来，它们的 doFilter() 方法均采用以下结构：

` Code1; //表示chain.doFilter()前的代码 chain.doFilter(); Code2; //表示chain.doFilter()后的代码 复制代码`

假定这两个过滤器都会为同一个 Servlet 预处理客户请求。当客户请求访问这个 Servlet 时，这两个过滤器及 Servlet 的工作流程如下图所示。

![681945C9-9ACE-4542-8621-5D10C988A858](https://user-gold-cdn.xitu.io/2019/4/26/16a572ae133730cb?imageView2/0/w/1280/h/960/ignore-error/1)

由于篇幅的问题，这里就不再展示串联过滤器的用法了，感兴趣的同学可以自行写个小demo尝试一下。只需要自己新建一个过滤器，然后早web.xml文件中添加一些相关配置即可。

## 16. 合理决定在 Servlet 中定义的变量的作用域类型 ##

在 Java 语言中，局部变量和实例变量有着不同的作用域，它们的区别如下：局部变量在一个方法中定义，每个线程都拥有自己的局部变量。实例变量在类中定义。类的每一个实例都拥有自己的实例变量，如果一个实例结束生命周期，那么属于它的实例变量也就结束生命周期。如果有多个线程同时执行一个实例的方法，而这个方法会访问一个实例变量，那么这些线程访问的是同一个实例变量。

## 最后 ##

觉得对您有帮助，欢迎评论转发点赞哦！

参考资料：

* Tomcat与Java Web开发技术详解 [ 孙卫琴 ]
* [Filter、FilterChain、FilterConfig 介绍]( https://link.juejin.im?target=http%3A%2F%2Fwww.runoob.com%2Fw3cnote%2Ffilter-filterchain-filterconfig-intro.html )
* [cookie 和 session 的区别详解]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fshiyangxt%2Farticles%2F1305506.html )