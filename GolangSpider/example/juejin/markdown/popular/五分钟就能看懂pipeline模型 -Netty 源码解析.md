# 五分钟就能看懂pipeline模型 -Netty 源码解析 #

## netty源码解析系列 ##

> 
> 
> 
> * [Netty 源码解析系列-服务端启动流程解析](
> https://juejin.im/post/5ce53460e51d4556db694979 )
> * [Netty 源码解析系列-客户端连接接入及读I/O解析](
> https://juejin.im/post/5cecfd09f265da1ba647ccf2 )
> * [五分钟就能看懂pipeline模型 -Netty 源码解析](
> https://juejin.im/post/5cf522dff265da1bab299967 )
> 
> 
> 

## 一、pipeline介绍 ##

### 1. 什么是pipeline ###

**pipeline** 有管道，流水线的意思，最早使用在 **Unix** 操作系统中，可以让不同功能的程序相互通讯，使软件更加”高内聚，低耦合”，它以一种”链式模型”来串起不同的程序或组件，使它们组成一条直线的工作流。

### 2. Netty的ChannelPipeline ###

**ChannelPipeline** 是处理或拦截 **channel** 的进站事件和出站事件的双向链表，事件在 **ChannelPipeline** 中流动和传递，可以增加或删除 **ChannelHandler** 来实现对不同业务逻辑的处理。通俗的说， **ChannelPipeline** 是工厂里的流水线， **ChannelHandler** 是流水线上的工人。
**ChannelPipeline** 在创建 **Channel** 时会自动创建，每个 **Channel** 都拥有自己的 **ChannelPipeline** 。

### 3. Netty I/O事件的处理过程 ###

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1db3633e9e214?imageView2/0/w/1280/h/960/ignore-error/1) 如图所示，入站事件是由 **I/O** 线程被动触发，由入站处理器按自下而上的方向处理，在中途可以被拦截丢弃，出站事件由用户 **handler** 主动触发，由出站处理器按自上而下的方向处理
![](https://user-gold-cdn.xitu.io/2019/6/3/16b1db3055f4b042?imageView2/0/w/1280/h/960/ignore-error/1)

## 二、ChannelHandlerContext ##

### 1. 什么是ChannelHandlerContext ###

**ChannelHandlerContext** 是将 **ChannelHandler** 和 **ChannelPipeline** 关联起来的上下文环境，每添加一个 **handler** 都会创建 **ChannelHandlerContext** 实例，管理 **ChannelHandler** 在 **ChannelPipeline** 中的传播流向。

### 2. ChannelHandlerContext和ChannelPipeline以及ChannelHandler之间的关系 ###

**ChannelPipeline** 依赖于 **Channel** 的创建而自动创建，保存了 **channel** ，将所有 **handler** 组织起来，相当于工厂的流水线。
**ChannelHandler** 拥有独立功能逻辑，可以注册到多个 **ChannelPipeline** ，是不保存 **channel** 的，相当于工厂的工人。
**ChannelHandlerContext** 是关联 **ChannelHandler** 和 **ChannelPipeline** 的上下文环境，保存了 **ChannelPipeline** ，控制 **ChannelHandler** 在 **ChannelPipeline** 中的传播流向，相当于流水线上的小组长。

## 三、传播Inbound事件 ##

### 1. Inbound事件有哪些？ ###

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1db26117c494d?imageView2/0/w/1280/h/960/ignore-error/1) **(1) channelRegistered** 注册事件， **channel** 注册到 **EventLoop** 上后调用，例如服务岗启动时， **pipeline.fireChannelRegistered()** ;
**(2) channelUnregistered** 注销事件， **channel** 从 **EventLoop** 上注销后调用，例如关闭连接成功后， **pipeline.fireChannelUnregistered();** **(3) channelActive** 激活事件，绑定端口成功后调用， **pipeline.fireChannelActive();**
**(4) channelInactive** 非激活事件，连接关闭后调用， **pipeline.fireChannelInactive();** **(5) channelRead** 读事件， **channel** 有数据时调用， **pipeline.fireChannelRead();**
**(6) channelReadComplete** 读完事件， **channel** 读完之后调用， **pipeline.fireChannelReadComplete();**
**(7) channelWritabilityChanged** 可写状态变更事件，当一个 **Channel** 的可写的状态发生改变的时候执行，可以保证写的操作不要太快，防止 **OOM** ， **pipeline.fireChannelWritabilityChanged();**
**(8) userEventTriggered** 用户事件触发，例如心跳检测， **ctx.fireUserEventTriggered(evt);**
**(9) exceptionCaught** 异常事件 说明:我们可以看出， **Inbound** 事件都是由 **I/O** 线程触发，用户实现部分关注的事件被动调用
**说明** : 我们可以看出， **Inbound** 事件都是由 **I/O** 线程触发，用户实现部分关注的事件被动调用

### 2. 添加读事件 ###

从前面 [《Netty 源码解析-服务端启动流程解析》]( https://juejin.im/post/5ce53460e51d4556db694979 ) 和 [《Netty 源码解析-客户端连接接入及读I/O解析》]( https://juejin.im/post/5cecfd09f265da1ba647ccf2 ) 我们知道，当有新连接接入时，我们执行注册流程，注册成功后，会调用 **channelRegistered** ，我们从这个方法开始

` public final void channelRegistered(ChannelHandlerContext ctx) throws Exception { initChannel((C) ctx.channel()); ctx.pipeline().remove(this); ctx.fireChannelRegistered(); } 复制代码`

**initChannel** 是在服务启动时配置的参数 **childHandler** 重写了父类方法

` private class IOChannelInitialize extends ChannelInitializer<SocketChannel> { @Override protected void initChannel(SocketChannel ch) throws Exception { System.out.println( "initChannel" ); ch.pipeline().addLast(new IdleStateHandler(1000, 0, 0)); ch.pipeline().addLast(new IOHandler()); } } 复制代码`

我们回忆一下， **pipeline** 是在哪里创建的

` protected AbstractChannel(Channel parent) { this.parent = parent; unsafe = newUnsafe(); pipeline = new DefaultChannelPipeline(this); } 复制代码`

当创建 **channel** 时会自动创建 **pipeline**

` public DefaultChannelPipeline(AbstractChannel channel) { if (channel == null) { throw new NullPointerException( "channel" ); } this.channel = channel; tail = new TailContext(this); head = new HeadContext(this); head.next = tail; tail.prev = head; } 复制代码`

在这里会创建两个默认的 **handler** ，一个 **InboundHandler --> TailContext** ，一个 **OutboundHandler --> HeadContext**
再看 **addLast** 方法

` @Override public ChannelPipeline addLast(ChannelHandler... handlers) { return addLast(null, handlers); } 复制代码`

在这里生成一个 **handler** 名字，生成规则由 **handler** 类名加 **”#0”**

` @Override public ChannelPipeline addLast(EventExecutorGroup executor, ChannelHandler... handlers) { … for (ChannelHandler h: handlers) { if (h == null) { break ; } addLast(executor, generateName(h), h); } return this; } 复制代码` ` @Override public ChannelPipeline addLast(EventExecutorGroup group, final String name, ChannelHandler handler) { synchronized (this) { checkDuplicateName(name); AbstractChannelHandlerContext newCtx = new DefaultChannelHandlerContext(this, group, name, handler); addLast0(name, newCtx); } return this; } 复制代码`

由于 **pipeline** 是线程非安全的，通过加锁来保证并发访问的安全，进行 **handler** 名称重复性校验，将 **handler** 包装成 **DefaultChannelHandlerContext** ，最后再添加到 **pipeline**

` private void addLast0(final String name, AbstractChannelHandlerContext newCtx) { checkMultiplicity(newCtx); AbstractChannelHandlerContext prev = tail.prev; newCtx.prev = prev; newCtx.next = tail; prev.next = newCtx; tail.prev = newCtx; name2ctx.put(name, newCtx); callHandlerAdded(newCtx); } 复制代码`

这里分三步
**(1)** 对 **DefaultChannelHandlerContext** 进行重复性校验，如果 **DefaultChannelHandlerContext** 不是可以在多个 **pipeline** 中共享的，且已经被添加到 **pipeline** 中，则抛出异常
**(2)** 修改 **pipeline** 中的指针
添加 **IdleStateHandler** 之前
**HeadContext --> IOChannelInitialize --> TailContext**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1db191e616115?imageView2/0/w/1280/h/960/ignore-error/1)

添加 **IdleStateHandler** 之后
**HeadContext --> IOChannelInitialize --> IdleStateHandler --> TailContext**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1db11c769cde4?imageView2/0/w/1280/h/960/ignore-error/1)

**(3)** 将 **handler** 名和 **DefaultChannelHandlerContext** 建立映射关系
**(4)** 回调 **handler** 添加完成监听事件
最后删除 **IOChannelInitialize**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1daf1d5d5c3af?imageView2/0/w/1280/h/960/ignore-error/1) 最后事件链上的顺序为:
**HeadContext --> IdleStateHandler --> IOHandler --> TailContext**

### 3. pipeline.fireChannelRead()事件解析 ###

在这里我们选一个比较典型的读事件解析，其他事件流程基本类似

` private static void processSelectedKey(SelectionKey k, AbstractNioChannel ch) { … if ((readyOps & (SelectionKey.OP_READ | SelectionKey.OP_ACCEPT)) != 0 || readyOps == 0) { unsafe.read(); } … } 复制代码`

当 **boss** 线程监听到读事件，会调用**unsafe.read()**方法

` @Override public final void read () { … pipeline.fireChannelRead(byteBuf); … } 复制代码`

入站事件从 **head** 开始， **tail** 结束

` @Override public ChannelPipeline fireChannelRead(Object msg) { head.fireChannelRead(msg); return this; } 复制代码` ` @Override public ChannelHandlerContext fireChannelRead(final Object msg) { if (msg == null) { throw new NullPointerException( "msg" ); } final AbstractChannelHandlerContext next = findContextInbound(); EventExecutor executor = next.executor(); if (executor.inEventLoop()) { next.invokeChannelRead(msg); } else { executor.execute(new OneTimeTask () { @Override public void run () { next.invokeChannelRead(msg); } }); } return this; } 复制代码`

查找 **pipeline** 中下一个 **Inbound** 事件

` private AbstractChannelHandlerContext findContextInbound () { AbstractChannelHandlerContext ctx = this; do { ctx = ctx.next; } while (!ctx.inbound); return ctx; } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1dae592a06bef?imageView2/0/w/1280/h/960/ignore-error/1) **HeadContext** 的下一个 **Inbound** 事件是 **IdleStateHandler**

` private void invokeChannelRead(Object msg) { try { ((ChannelInboundHandler) handler()).channelRead(this, msg); } catch (Throwable t) { notifyHandlerException(t); } } 复制代码` ` @Override public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception { if (readerIdleTimeNanos > 0 || allIdleTimeNanos > 0) { reading = true ; firstReaderIdleEvent = firstAllIdleEvent = true ; } ctx.fireChannelRead(msg); } 复制代码`

将这个 **channel** 读事件标识为 **true** ，并传到下一个 **handler**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1dad769bb39a3?imageView2/0/w/1280/h/960/ignore-error/1)

` @Override public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception { super.channelRead(ctx, msg); System.out.println(msg.toString()); } 复制代码`

这里执行 **IOHandler** 重写的 **channelRead() **方法,并调用父类** channelRead** 方法

` @Override public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception { ctx.fireChannelRead(msg); } 复制代码`

继续调用事件链上的下一个 **handler**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1dacf3717f352?imageView2/0/w/1280/h/960/ignore-error/1)

` @Override public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception { try { logger.debug( "Discarded inbound message {} that reached at the tail of the pipeline. " + "Please check your pipeline configuration." , msg); } finally { ReferenceCountUtil.release(msg); } } 复制代码`

这里会调用 **TailContext** 的 **Read** 方法，释放 **msg** 缓存
**总结: **传播** Inbound** 事件是从 **HeadContext** 节点往上传播，一直到 **TailContext** 节点结束

## 四、传播Outbound事件 ##

### 1. Outbound事件有哪些？ ###

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1dac6420632aa?imageView2/0/w/1280/h/960/ignore-error/1) **(1) bind** 事件,绑定端口
**(2) close** 事件，关闭channel
**(3) connect** 事件，用于客户端，连接一个远程机器
**(4) disconnect** 事件，用于客户端，关闭远程连接
**(5) deregister** 事件，用于客户端，在执行断开连接 **disconnect** 操作后调用，将 **channel** 从 **EventLoop** 中注销
**(6) read** 事件，用于新接入连接时，注册成功多路复用器上后，修改监听为 **OP_READ** 操作位
**(7) write** 事件，向通道写数据
**(8) flush** 事件，将通道排队的数据刷新到远程机器上

### 2. 解析write事件 ###

` ByteBuf resp = Unpooled.copiedBuffer( "hello".getBytes()); ctx.channel().write(resp); 复制代码`

我们在项目中像上面这样直接调用 **write** 写数据，并不能直接写进 **channel** ，而是写到缓冲区，还要调用 **flush** 方法才能将数据刷进 **channel** ，或者直接调用 **writeAndFlush** 。
在这里我们选择比较典型的 **write** 事件来解析 **Outbound** 流程，其他事件流程类似

` @Override public ChannelFuture write(Object msg) { return pipeline.write(msg); } 复制代码`

通过上下文绑定的 **channel** 直接调用 **write** 方法，调用 **channel** 相对应的事件链上的 **handler**

` @Override public ChannelFuture write(Object msg) { return tail.write(msg); } 复制代码`

写事件是从 **tail** 向 **head** 调用，和读事件刚好相反

` @Override public ChannelFuture write(Object msg) { return write(msg, newPromise()); } 复制代码` ` @Override public ChannelFuture write(final Object msg, final ChannelPromise promise) { ... write(msg, false , promise); ... } 复制代码` ` private void write(Object msg, boolean flush, ChannelPromise promise) { AbstractChannelHandlerContext next = findContextOutbound(); EventExecutor executor = next.executor(); if (executor.inEventLoop()) { next.invokeWrite(msg, promise); if (flush) { next.invokeFlush(); } ... } ... } 复制代码`

经过多次跳转，获取上一个 **Ounbound** 事件链的 **handler**

` private AbstractChannelHandlerContext findContextOutbound () { AbstractChannelHandlerContext ctx = this; do { ctx = ctx.prev; } while (!ctx.outbound); return ctx; } 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1dab8627f6a32?imageView2/0/w/1280/h/960/ignore-error/1) **IdleStateHandler** 既是 **Inbound** 事件，又是 **Outbound** 事件
继续跳转到上一个 **handler**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1daacc14091c3?imageView2/0/w/1280/h/960/ignore-error/1) 上一个是 **HeadContext** 处理

` @Override public void write(ChannelHandlerContext ctx, Object msg, ChannelPromise promise) throws Exception { unsafe.write(msg, promise); } 复制代码` ` @Override public final void write(Object msg, ChannelPromise promise) { ChannelOutboundBuffer outboundBuffer = this.outboundBuffer; ... outboundBuffer.addMessage(msg, size, promise); ... } 复制代码`

从这里我们看到，最终是把数据丢到了缓冲区，自此 **netty** 的 **pipeline** 模型我们解析完毕
有关 **inbound** 事件和 **outbound** 事件的传输, 可通过下图进行归纳:

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1da87e399add2?imageView2/0/w/1280/h/960/ignore-error/1)

觉得对您有帮助请点 **"赞"**