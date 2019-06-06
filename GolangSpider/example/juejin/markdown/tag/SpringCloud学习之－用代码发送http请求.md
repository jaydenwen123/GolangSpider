# SpringCloud学习之－用代码发送http请求 #

在spring系列框架中，使用代码向另一个服务发送请求，一般可以用restTemplate或者feign，但是这两种方法的底层原理是啥呢，今天就来探究一下。

首先建立服务端：

` @RestController public class Controller { @PostMapping( "hello" ) public String hello () { return "hello" ; } } 复制代码`

新建一个springboot服务，创建这么一个controller，然后启动服务。

其次使用代码发送请求来调用这个服务：

` public static void main(String[] args) throws IOException { URI uri = URI.create( "http://127.0.0.1:8080/hello" ); URL url = uri.toURL(); HttpURLConnection connection = (HttpURLConnection) url.openConnection(); connection.setRequestMethod(HttpMethod.POST.name()); connection.setDoOutput( true ); connection.connect(); InputStream inputStream = connection.getInputStream(); BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream)); System.out.println(reader.readLine()); } 复制代码`

这样就可以看到控制台会返回一个“hello”。

这就是java代码发送网络请求的基本api，spring框架在此基础上进行了进一步的封装。

下一篇就来拆解spring对jdk的基本api进行了哪些封装。

## **返回目录** ( https://juejin.im/post/5c8a4458f265da2da23d703c ) ##