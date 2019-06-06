# 深入理解Java I/O模型 #

### 前言 ###

我们在开发过程中经常跟I/O打交道，很多人在学习I/O模型过程和进行I/O编程过程中，对很多概念可能不明朗，特别是像Java这样的高级语言，它对底层操作系统的各种I/O模型进行了封装，使得我们可以很轻松的进行开发，但是在方便之余你是否对Java中各种I/O模型，以及它们和操作系统之间的关联是否有过了解？

### 什么是I/O? ###

I/O 在计算机中指Input/Output,即输入输出。以一次文件读取为例，我们需要将磁盘上的数据读取到用户空间，那么这次数据转移操作其实就是一次I/O操作，也就是一次文件I/O；我们每天都浏览着各种各样的网页，在我们每请求一个网页，服务器通过网络将一个个的分组数据发送给我们，应用程序从TCP缓冲区将数据复制到用户空间的过程也是一次I/O,即一次网络I/O。可以发现I/O如此重要，它时刻都在。

## Liunx 网络 I/O模型 ##

根据UNIX网络编程对I/O模型的分类，UNIX提供了5中I/O模型分别如下：

### 阻塞I/O模型 ###

这是最传统的I/O模型，即在读写数据过程中会阻塞，我们通过图可以看到，在应用进程调用recvfrom,系统调用直到数据从内核从复制到用户空间，应用进程在这一段时间内一直是被阻塞的。 **这种模型适合并发量较小的对时延不敏感的系统** 。

![](https://user-gold-cdn.xitu.io/2019/5/5/16a870b9a955f492?imageView2/0/w/1280/h/960/ignore-error/1)

### 非阻塞I/O模型 ###

应用进程不停的通过recvfrom调用不停的和内核交互直到数据被被准备好，将他复制到用户空间中，如果recvfrom调用没有数据可以返回时返回一个EWOULDBLOCK错误， **我们将这样的操作称作轮询，这么做往往需要耗费大量的CPU时间** 。

![](https://user-gold-cdn.xitu.io/2019/5/5/16a8726d68af8356?imageView2/0/w/1280/h/960/ignore-error/1)

### I/O复用模型 ###

在Liunx中为我们提供了select/poll,也就是管道，我们就可以将调用它们阻塞在这两个系统调用中的一个上，而不是阻塞在真正的I/O调用上，我们阻塞select调用当数据返回可读条件时，通过recvfrom调用将数据复制到应用程序缓冲区。 **多路I/O复用本质上并不是非阻塞的，对比阻塞I/O模型它并没有什么优势，事实上使用select需要两个系统而不是当个调用，I/O复用其实稍有劣势，它只是能处理更多的连接（等待多个I/O就绪）**

![](https://user-gold-cdn.xitu.io/2019/5/5/16a87b415ed485de?imageView2/0/w/1280/h/960/ignore-error/1)

### 信号驱动式I/O模型 ###

我们首先开启套接字的信号驱动I/O功能，通过sigaction系统调用安装一个信号处理函数，系统调用立即返回，进程继续工作，当数据包准备好时内核产生一个SIGIO信号通知，我们通过recvfrom调用读取数据报。 **信号驱动式I/O模型的优点是我们在数据报到达期间进程不会被阻塞，我们只要等待信号处理函数的通知即可**

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8b4049b4d0e83?imageView2/0/w/1280/h/960/ignore-error/1)

### 异步I/O模型 ###

告知内核启动某个操作(包括将数据从内核复制到自己的缓冲区)之后通知我们。 **信号驱动模型是内核通知我们何时启动一个I/O操作，而异步I/O模型是由内核通知我们I/O何时完成**

![](https://user-gold-cdn.xitu.io/2019/5/7/16a8fe9a87035c0f?imageView2/0/w/1280/h/960/ignore-error/1)

### 同步I/O和异步I/O对比 ###

同步I/O操作：导致请求进程阻塞，直到I/O操作完成

异步I/O操作：不导致请求进程阻塞

![](https://user-gold-cdn.xitu.io/2019/5/7/16a8ff5d8f734fa3?imageView2/0/w/1280/h/960/ignore-error/1) 综上 **阻塞式I/O模型、非阻塞式I/O模型、I/O复用模型和信号驱动模型都是同步I/O模型，他们真正的I/O操作将进程阻塞，只有异步I/O模型是异步I/O操作**

## Java I/O 模型 ##

### Java I/O历史 ###

在JDK 1.4之前，基于Java的所有Socket通信都使用了同步阻塞模式（Blocking I/O），这种一请求一应答的通信模型简化了上层开发，但性能可靠性存在巨大瓶颈，对高并发和低时延支持不好

在JDK 1.4之后，提供了新的NIO(New I/O)类库，Java也可以支持非阻塞I/O了，新增了java.nio包，提供了很多异步I/O开发的API和类库。

JDK 1.7发布后，将原来的NIO类库进行了升级，提供了AIO功能，支持基于文件的异步I/O操作和针对套接字的异步I/O操作等功能

### BIO 编程 ###

使用BIO通信模型的服务端，通常通过一个独立的Acceptor线程负责监听客户端的连接，监听到客户端连接请求后为每一个客户端创建一个新的线程链路进行处理，处理完成通过输出流回应客户端，线程消耗，这就是典型一对一答模型，下面我们通过代码对BIO模式进行具体分析，我们实现客户端发送消息服务端将消息回传我们的功能。

服务端：

` int port = 3000 ; try (ServerSocket serverSocket = new ServerSocket(port)) { Socket socket = null ; while ( true ) { //主程序阻塞在accept操作上 socket = serverSocket.accept(); new Thread( new BioExampleServerHandle(socket)).start(); } } catch (Exception e) { e.printStackTrace(); } 复制代码` ` private Socket socket; public BioExampleServerHandle (Socket socket) { this.socket = socket; } @Override public void run () { try (BufferedReader reader = new BufferedReader( new InputStreamReader(socket.getInputStream())); PrintWriter writer = new PrintWriter(socket.getOutputStream(), true )) { String message = reader.readLine(); System.out.println( "收到客户端消息：" + message); writer.println( "answer: " + message); } catch (Exception e) { e.printStackTrace(); } } 复制代码`

客户端：

` String host = "127.0.0.1" ; int port = 3000; try(Socket socket = new Socket(host, port); BufferedReader reader = new BufferedReader(new InputStreamReader(socket.getInputStream())); PrintWriter writer = new PrintWriter(socket.getOutputStream(), true )) { Scanner input = new Scanner(System.in); System.out.println( "输入你想说的话：" ); String message = input.nextLine(); writer.println(message); String answer = reader.readLine(); System.out.println(answer); } catch (Exception e) { e.printStackTrace(); } 复制代码`

运行结果如下：

客户端：

![](https://user-gold-cdn.xitu.io/2019/5/7/16a9071867ca9fa9?imageView2/0/w/1280/h/960/ignore-error/1) 服务端： ![](https://user-gold-cdn.xitu.io/2019/5/7/16a907002e137122?imageView2/0/w/1280/h/960/ignore-error/1) 通过代码我们可以发现BIO的主要问题在于， **每当一个连接接入时我们都需要new一个线程进行处理** ，这显然是不合适的，因为一个线程只能处理一个连接， **如果在高并发的情况下，我们的程序肯定无法满足性能需求，同时我们对线程创建也缺乏管理** 。为了改进这种模型我们可以通过消息队列和线程池技术对他加以优化，我们称它为伪异步I/O，代码如下：

` int port = 3000 ; ThreadPoolExecutor socketPool = null ; try (ServerSocket serverSocket = new ServerSocket(port)) { Socket socket = null ; int cpuNum = Runtime.getRuntime().availableProcessors(); socketPool = new ThreadPoolExecutor(cpuNum, cpuNum * 2 , 1000 , TimeUnit.SECONDS, new ArrayBlockingQueue<Runnable>( 1000 )); while ( true ) { socket = serverSocket.accept(); socketPool.submit( new BioExampleServerHandle(socket)); } } catch (Exception e) { e.printStackTrace(); } finally { socketPool.shutdown(); } 复制代码`

可以看到每当有新连接接入，我们都将他投递给线程池进行处理，由于我们设置了线程池大小和阻塞队列大小，因此在并发情况下都不会导致服务崩溃，但是如果并发数大于阻塞队列大小，或服务端处理连接缓慢时，阻塞队列无法继续处理，会导致客户端连接超时，影响用户体验。

### NIO 编程 ###

NIO 弥补了同步阻塞I/O的不足，它提供了高速、面向块的I/O，我们对一些概念介绍一下：

**Buffer** : Buffer用于和NIO通道进行交互。数据从通道读入缓冲区，从缓冲区写入到通道中,它的主要作用就是和Channel进行交互。

**Channel** : Channel是一个通道，可以通过它读取和写入数据，通道是双向的，通道可以用于读、写或者同时读写。

**Selector** : Selector会不断的轮询注册在它上面的Channe,如果Channel上面有新的连接读写事件的时候就会被轮询出来，一个Selector可以注册多个Channel， **只需要一个线程负责Selector轮询，就可以支持成千上万的连接，可以说为高并发服务器的开发提供了很好的支撑** 。

我们通过实际代码演示NIO的使用：

服务端代码：

` int port = 3000 ; ServerSocketChannel socketChannel = null ; Selector selector = null ; try { selector = Selector.open(); socketChannel = ServerSocketChannel.open(); //设置连接模式为非阻塞模式 socketChannel.configureBlocking( false ); socketChannel.socket().bind( new InetSocketAddress(port)); //在selector上注册通道，监听连接事件 socketChannel.register(selector, SelectionKey.OP_ACCEPT); while ( true ) { //设置selector 每隔一秒扫描所有channel selector.select( 1000 ); Set<SelectionKey> selectionKeys = selector.selectedKeys(); Iterator<SelectionKey> iterable = selectionKeys.iterator(); SelectionKey key = null ; while (iterable.hasNext()) { key = iterable.next(); //对key进行处理 try { handlerKey(key, selector); } catch (Exception e) { if ( null != key) { key.cancel(); if ( null != key.channel()) { key.channel().close(); } } } } } } catch (Exception e) { e.printStackTrace(); } finally { try { if ( null != selector) { selector.close(); } if ( null != socketChannel) { socketChannel.close(); } } catch (Exception e) { throw new RuntimeException(e); } } 复制代码`

handlerKey代码如下：

` private void handlerKey (SelectionKey key, Selector selector) throws IOException { if (key.isValid()) { //判断是否是连接请求，对所有连接请求进行处理 if (key.isAcceptable()) { ServerSocketChannel serverSocketChannel = (ServerSocketChannel) key.channel(); SocketChannel channel = serverSocketChannel.accept(); channel.configureBlocking( false ); //在selector上注册通道，监听读事件 channel.register(selector, SelectionKey.OP_READ); } else if (key.isReadable()) { SocketChannel channel = (SocketChannel) key.channel(); //分配一个1024字节的缓冲区 ByteBuffer byteBuffer = ByteBuffer.allocate( 1024 ); int readBytes = channel.read(byteBuffer); if (readBytes > 0 ) { //从写模式切换到读模式 byteBuffer.flip(); byte [] bytes = new byte [byteBuffer.remaining()]; byteBuffer.get(bytes); String message = new String(bytes, "UTF-8" ); System.out.println( "收到客户端消息: " + message); //回复客户端 message = "answer: " + message; byte [] responseByte = message.getBytes(); ByteBuffer writeBuffer = ByteBuffer.allocate(responseByte.length); writeBuffer.put(responseByte); writeBuffer.flip(); channel.write(writeBuffer); } } } } 复制代码`

客户端代码：

` int port = 3000 ; String host = "127.0.0.1" ; SocketChannel channel = null ; Selector selector = null ; try { selector = Selector.open(); channel = SocketChannel.open(); channel.configureBlocking( false ); if (channel.connect( new InetSocketAddress(host, port))) { channel.register(selector, SelectionKey.OP_READ); write(channel); } else { channel.register(selector, SelectionKey.OP_CONNECT); } while ( true ) { selector.select( 1000 ); Set<SelectionKey> selectionKeys = selector.selectedKeys(); Iterator<SelectionKey> iterator = selectionKeys.iterator(); SelectionKey key = null ; while (iterator.hasNext()) { try { key = iterator.next(); handle(key, selector); } catch (Exception e) { e.printStackTrace(); if ( null != key.channel()) { key.channel().close(); } if ( null != key) { key.cancel(); } } } } } catch (Exception e) { e.printStackTrace(); } finally { try { if ( null != channel) { channel.close(); } if ( null != selector) { selector.close(); } } catch (Exception e) { throw new RuntimeException(e); } } 复制代码`

write 方法：

` private void write (SocketChannel channel) throws IOException { Scanner in = new Scanner(System.in); System.out.println( "输入你想说的话：" ); String message = in.next(); byte [] bytes = message.getBytes(); ByteBuffer byteBuffer = ByteBuffer.allocate(bytes.length); byteBuffer.put(bytes); byteBuffer.flip(); channel.write(byteBuffer); } 复制代码`

handle 方法：

` private void handle (SelectionKey key, Selector selector) throws IOException { if (key.isValid()) { SocketChannel channel = (SocketChannel) key.channel(); if (key.isConnectable()) { if (channel.finishConnect()) { channel.register(selector, SelectionKey.OP_READ); write(channel); } } else if (key.isReadable()) { ByteBuffer byteBuffer = ByteBuffer.allocate( 1024 ); int readBytes = channel.read(byteBuffer); if (readBytes > 0 ) { byteBuffer.flip(); byte [] bytes = new byte [byteBuffer.remaining()]; byteBuffer.get(bytes); String message = new String(bytes, "UTF-8" ); System.out.println(message); } else if (readBytes < 0 ) { key.cancel(); channel.close(); } } } } 复制代码`

通过代码我们发现NIO比BIO复杂太多，这个代码量也是刷刷的增长啊，虽然复杂但是NIO的优点也值得我们去尝试，比起 **BIO客户端连接操作是异步的，我们可以注册OP_CONNECT事件等待结果而不用像那样被同步阻塞** ， **Channel的读写操作都是异步的，没有等待数据它不会等待直接返回，比起BIO我们不需要频繁的创建线程来处理客户端连接，我们通过一个Selector处理多个客户端连接，而且性能也可以得到保障，适合做高性能服务器开发**

### AIO 编程 ###

NIO2.0 引入了异步通道的概念，提供了异步文件通道和异步套接字通道的实现，我们可以通过Future类来表示异步操作结果，也可以在执行异步操作的时候传入一个Channels,实现CompletionHandler接口为操作回调。示例代码如下

服务端：

` int port = 3000 ; AsynchronousServerSocketChannel socketChannel = null ; try { socketChannel = AsynchronousServerSocketChannel.open(); socketChannel.bind( new InetSocketAddress(port)); //接收客户端连接，传入AcceptCompletionHandler作为回调来接收连接消息 socketChannel.accept(socketChannel, new AcceptCompletionHandler()); Thread.currentThread().join(); } catch (Exception e) { e.printStackTrace(); } finally { try { if ( null != socketChannel) { socketChannel.close(); } } catch (Exception e1) { throw new RuntimeException(e1); } } 复制代码`

AcceptCompletionHandler 类：

` public class AcceptCompletionHandler implements CompletionHandler < AsynchronousSocketChannel , AsynchronousServerSocketChannel > { @Override public void completed (AsynchronousSocketChannel result, AsynchronousServerSocketChannel attachment) { //继续接受其他客户端的连接请求，形成一个循环 attachment.accept(attachment, this ); ByteBuffer byteBuffer = ByteBuffer.allocate( 1024 ); //调用read操作进行异步读取操作，传入ReadCompletionHandler作为回调 result.read(byteBuffer, byteBuffer, new ReadCompletionHandler(result)); } @Override public void failed (Throwable exc, AsynchronousServerSocketChannel attachment) { //异常失败处理在这里 } } 复制代码`

ReadCompletionHandler 类

` public class ReadCompletionHandler implements CompletionHandler < Integer , ByteBuffer > { private AsynchronousSocketChannel channel; public ReadCompletionHandler (AsynchronousSocketChannel channel) { this.channel = channel; } @Override public void completed (Integer result, ByteBuffer byteBuffer) { try { byteBuffer.flip(); byte [] bytes = new byte [byteBuffer.remaining()]; byteBuffer.get(bytes); String message = new String(bytes, "UTF-8" ); System.out.println( "收到客户端消息：: " + message); write(message); } catch (UnsupportedEncodingException e) { e.printStackTrace(); } } @Override public void failed (Throwable exc, ByteBuffer attachment) { try { channel.close(); } catch (Exception e) { throw new RuntimeException(e); } } private void write (String message) { message = "answer: " + message; byte [] bytes = message.getBytes(); ByteBuffer byteBuffer = ByteBuffer.allocate(bytes.length); byteBuffer.put(bytes); byteBuffer.flip(); channel.write(byteBuffer, byteBuffer, new WriteCompletionHandler(channel)); } } 复制代码`

客户端：

` int port = 3000 ; String host = "127.0.0.1" ; AsynchronousSocketChannel channel = null ; try { channel = AsynchronousSocketChannel.open(); channel.connect( new InetSocketAddress(host, port), channel, new AioClientHandler()); Thread.currentThread().join(); } catch (Exception e) { e.printStackTrace(); } finally { try { if ( null != channel) { channel.close(); } } catch (Exception e) { throw new RuntimeException(e); } } 复制代码`

AioClientHandler 类(由于客户端比较简单我这里使用了嵌套类部类)：

` public class AioClientHandler implements CompletionHandler < Void , AsynchronousSocketChannel > { @Override public void completed (Void result, AsynchronousSocketChannel channel) { Scanner in = new Scanner(System.in); System.out.println( "输入你想说的话：" ); String message = in.next(); byte [] bytes = message.getBytes(); ByteBuffer byteBuffer = ByteBuffer.allocate(bytes.length); byteBuffer.put(bytes); byteBuffer.flip(); channel.write(byteBuffer, byteBuffer, new CompletionHandler<Integer, ByteBuffer>() { @Override public void completed (Integer result, ByteBuffer buffer) { //判断是否写完如果没有继续写 if (buffer.hasRemaining()) { channel.write(buffer, buffer, this ); } else { ByteBuffer readBuffer = ByteBuffer.allocate( 1024 ); channel.read(readBuffer, readBuffer, new CompletionHandler<Integer, ByteBuffer>() { @Override public void completed (Integer result, ByteBuffer attachment) { try { attachment.flip(); byte [] bytes1 = new byte [attachment.remaining()]; attachment.get(bytes1); String message = new String(bytes1, "UTF-8" ); System.out.println(message); System.exit( 1 ); } catch (UnsupportedEncodingException e) { e.printStackTrace(); } } @Override public void failed (Throwable exc, ByteBuffer attachment) { try { channel.close(); } catch (Exception e) { throw new RuntimeException(e); } } }); } } @Override public void failed (Throwable exc, ByteBuffer attachment) { try { channel.close(); } catch (Exception e) { throw new RuntimeException(e); } } }); } @Override public void failed (Throwable exc, AsynchronousSocketChannel attachment) { } 复制代码`

通过对比代码我们发现AIO比BIO简单，这是因为我们不需要创建一个独立的I/O线程来来处理读写操作， AsynchronousSocketChannel、AsynchronousServerSocketChannel由JDK底层线程池负责回调驱动读写操作。

### 对比 ###

+----------------+------------------+-----------+----------------+--------------+
|                | 同步阻塞I/O(BIO) | 伪异步I/O | 非阻塞I/O(NIO) | 异步I/O(AIO) |
+----------------+------------------+-----------+----------------+--------------+
| 是否阻塞       | 是               | 是        | 否             | 否           |
| 是否同步       | 是               | 是        | 是             | 否(异步)     |
| 程序员友好程度 | 简单             | 简单      | 非常难         | 比较难       |
| 可靠性         | 非常差           | 差        | 高             | 高           |
| 吞吐量         | 低               | 中        | 高             | 高           |
+----------------+------------------+-----------+----------------+--------------+

### 总结 ###

通过学习Lunix底层I/O模型和JavaI/O模型我们发现上层只是对底层的抽象和封装，BIO其实是对阻塞I/O模型的实现，NIO是对I/O复用模型的实现，AIO是对信号驱动I/O的实现，理解了底层I/O模型，在实际开发中应该可以很自如。如果你觉得不错的话就点个赞吧，如果有bug也您请批评指正，您的赞赏和批评是进步路上的好伙伴。