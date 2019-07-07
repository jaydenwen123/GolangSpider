# 深入浅出 Node.js Cluster #

> 
> 
> 
> 本文首发于 [猫眼前端团队公众号](
> https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FXm_c841UdKA06s76rJ_4nw
> ) ，转载请注明出处。
> 
> 

## 前言 ##

如果大家用 PM2 管理 Node.js 进程，会发现它支持一种 cluster mode。开启 cluster mode 后，支持给 Node.js 创建多个进程。 如果将 cluster mode 下的 instances 设置为 max 的话，它还会根据服务器的 CPU 核心数，来创建对应数量的 Node 进程。

![](https://user-gold-cdn.xitu.io/2019/3/12/169712524145986b?imageView2/0/w/1280/h/960/ignore-error/1)

PM2 其实利用的是 Node.js Cluster 模块来实现的，这个模块的出现就是为了解决 Node.js 实例单线程运行，无法利用多核 CPU 的优势而出现的。那么，Cluster 内部又是如何工作的呢？多个进程间是如何通信的？多个进程是如何监听同一个端口的？Node.js 是如何将请求分发到各个进程上的？如果你对上述问题还不清楚，不妨接着往下看。

## 核心原理 ##

Node.js worker 进程由 ` child_process.fork()` ( https://link.juejin.im?target=http%3A%2F%2Fnodejs.cn%2Fs%2FVDCJMa ) 方法创建，这也意味存在着父进程和多个子进程。代码大致是这样：

` const cluster = require( 'cluster' ); const os = require( 'os' ); if (cluster.isMaster) { for (var i = 0, n = os.cpus().length; i < n; i += 1) { cluster.fork(); } } else { // 启动程序 } 复制代码`

学过操作系统的同学，应该对 fork() 这个系统调用不陌生，调用它的进程为父进程，fork 出来的都是子进程。子进程和父进程具有相同的代码段、数据段、堆栈，但是它们的内存空间不共享。父进程（即 master 进程）负责监听端口，接收到新的请求后将其分发给下面的 worker 进程。这里涉及三个问题：父子进程通信、负载均衡策略以及多进程的端口监听。

备注：Linux 上 fork() 支持写时复制，只有进程空间的各段的内容要发生变化时，才会将父进程的内容复制一份给子进程。因此子进程和父进程一开始是共享相同的内存空间。

#### 进程通信 ####

master 进程通过 process.fork() 创建子进程，他们之间通过 IPC (内部进程通信)通道实现通信。操作系统的进程间通信方式主要有以下几种：

* 共享内存 不同进程共享同一段内存空间。通常还需要引入信号量机制，来实现同步与互斥。
* 消息传递 这种模式下，进程间通过发送、接收消息来实现信息的同步。
* 信号量 信号量简单说就是系统赋予进程的一个状态值，未得到控制权的进程会在特定地方被强迫停下来，等待可以继续进行的信号到来。如果信号量只有 0 或者 1 两个值的话，又被称作“互斥锁”。这个机制也被广泛用于各种编程模式中。
* 管道 管道本身也是一个进程，它用于连接两个进程，将一个进程的输出作为另一个进程的输入。可以用 pipe 系统调用来创建管道。我们经常用的“ | ”命令行就是利用了管道机制。

Node.js 为父子进程的通信提供了事件机制来传递消息。下面的例子实现了父进程将 TCP server 对象句柄传给子进程。

` const subprocess = require( 'child_process' ).fork( 'subprocess.js' ); // 开启 server 对象，并发送该句柄。 const server = require( 'net' ).createServer(); server.on( 'connection' , (socket) => { socket.end( '被父进程处理' ); }); server.listen(1337, () => { subprocess.send( 'server' , server); }); 复制代码` ` process.on( 'message' , (m, server) => { if (m === 'server' ) { server.on( 'connection' , (socket) => { socket.end( '被子进程处理' ); }); } }); 复制代码`

那么问题又来了，如果进程间没有父子关系，换句话说，我们应该如何实现任意进程间的通信呢？大家可以去看看这篇文章： [进程间通信的另类实现]( https://link.juejin.im?target=http%3A%2F%2Ftaobaofed.org%2Fblog%2F2016%2F01%2F27%2Fnodejs-ipc%2F )

#### 负载均衡策略 ####

前面提到，所有请求是通过 master 进程分配的，要保证服务器负载比较均衡的分配到各个 worker 进程上，这就涉及到负载均衡策略了。Node.js 默认采用的策略是 **round-robin** 时间片轮转法。

round-robin 是一种很常见的负载均衡算法，Nginx 上也采用了它作为负载均衡策略之一。它的原理很简单，每一次把来自用户的请求轮流分配给各个进程，从 1 开始，直到 N(worker 进程个数)，然后重新开始循环。这个算法的问题在于，它是假定各个进程或者说各个服务器的处理性能是一样的，但是如果请求处理间隔较长，就容易导致出现负载不均衡。因此我们通常在 Nginx 上采用另一种算法： **WRR** ，加权轮转法。通过给各个服务器分配一定的权重，每次选出权重最大的，给其权重减 1，直到权重全部为 0 后，按照此时生成的序列轮询。

可以通过设置 NODE_CLUSTER_SCHED_POLICY 环境变量，或者通过 cluster.setupMaster(options) 来修改负载均衡策略。读到这里大家可以发现，我们可以 Nginx 做多机器集群上的负载均衡，然后用 Node.js Cluster 来实现单机多进程上的负载均衡。

#### 多进程的端口监听 ####

最初的 Node.js 上，多个进程监听同一个端口，它们相互竞争新 accept 过来的连接。这样会导致各个进程的负载很不均衡，于是后来使用了上文提到的 round-robin 策略。具体思路是，master 进程创建 socket，绑定地址并进行监听。该 socket 的 fd 不传递到各个 worker 进程。当 master 进程获取到新的连接时，再决定将 accept 到的客户端连接分发给指定的 worker 处理。简单说就是，master 进程监听端口，然后将连接通过某种分发策略（比如 round-robin），转发给 worker 进程。这样由于只有 master 进程接收客户端连接，就解决了竞争导致的负载不均衡的问题。但是这样设计就要求 master 进程的稳定性足够好了。

## 总结 ##

本文以 PM2 的 Cluster Mode 作为切入点，向大家介绍了 Node.js Cluster 实现多进程的核心原理。重点讲了进程通信、负载均衡以及多进程端口监听三个方面。通过研究 cluster 模块可以发现，很多底层原理或者是算法，其实都是通用的。比如 round-robin 算法，它在操作系统底层的进程调度中也有使用；比如 master-worker 这种架构，是不是在 Nginx 的多进程架构中也似曾相识；比如信号量、管道这些机制，也可以在各种编程模式中见到它们的身影。当下市面上各种新技术层出不穷，但核心其实是 **万变不离其宗** ，理解了这些最基础的知识，剩下的也可以触类旁通了。

参考链接：

* [当我们谈论 cluster 时我们在谈论什么（下）]( https://link.juejin.im?target=http%3A%2F%2Ftaobaofed.org%2Fblog%2F2015%2F11%2F10%2Fnodejs-cluster-2%2F )
* [Node.js进阶：cluster模块深入剖析]( https://juejin.im/entry/5ad3eb536fb9a028d375db4e )
* [进程间通信的另类实现]( https://link.juejin.im?target=http%3A%2F%2Ftaobaofed.org%2Fblog%2F2016%2F01%2F27%2Fnodejs-ipc%2F )