# webrtc+canvas+socket.io从零实现一个你画我猜 | 掘金技术征文 #

### 开场白 ###

最近键盘坏了，刚好看到掘金有声网的技术征文，想整个键盘。于是就开始从零开始学习webrtc， 一开始看文档就是个素质三连。这么难啊，这咋整啊，这谁顶的住啊。于是就开始全网找资料，很幸运的在掘金上找到了 [江三疯大佬的webrtc系列]( https://juejin.im/post/5c3acfa56fb9a049f36254be ) ，以及 [WebRTC实时通信系列教程]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fleytton%2Farticle%2Fdetails%2F76696372 ) ,或者英文原版的 [Real time communication with WebRTC]( https://link.juejin.im?target=https%3A%2F%2Fcodelabs.developers.google.com%2Fcodelabs%2Fwebrtc-web%2F ) ，有兴趣的同学也可以去看下，非常棒。既然有这么棒的文章为啥还要再写篇文章呢，那当然是分(zheng)享(ge)经(jian)验(pan)啦。鉴于自己耗时将近三周的学习加项目，项目写着写着就破千行了(枯惹)，虽然中途有事情耽误了一段时间，但是也是花费了我极其大的精力，踩了无数的坑，这里我会尽可能从最基础开始用简答易懂的方式，带领大家完成一个较完整你画我猜。文章可能会很长，可以慢慢看。有些知识点不需要那么详细，为了让你思路更清晰会省略介绍，有兴趣的可以自己去看。

### 项目演示 ###

![演示](https://user-gold-cdn.xitu.io/2019/6/3/16b1b6159e202ada?imageslim)

Github地址： [你画我猜]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyuyuyue%2Fwebrtc-draw-guess )

欢迎Star！

### webrtc ###

> 
> 
> 
> WebRTC (Web Real-Time Communication)是一个可以用在视频聊天，音频聊天或P2P文件分享等Web App中的
> API。
> 
> 

全名叫web的实时通信，从官方文档可以看出来他可以用来视频聊天，音频聊天，端对端(p2p)，数据传输，文件分享的一个api。现在的直播用的就是这个技术

webrtc下有三个重要的api，正好对应三个功能。

* getUserMedia 请求获取用户的媒体信息包括视频流(video)和音频流(audio)
* RTCPeerConnection 代表一个由本地计算机到远端的WebRTC连接,用于实现端对端的连接。该接口提供了创建，保持，监控，关闭连接的方法的实现。
* RTCDataChannel 代表在两者之间建立了一个双向数据通道的连接,是一个数据通道，传输数据

#### **getUserMedia** ####

首先我们先实现一个简单的获取视频和音频并且显示在网页上

` javasrcipt // 获取本地的视频和音频流，{ audio: true , video: true }都是 true 这两个都获取 let local Stream = navigator.mediaDevices.getUserMedia({ audio: true , video: true }) .then((stream) => stream) //找到video标签，用一个video来接受流，并且显示 let video = document.querySelector( "#video" ) // 使用srcObject给video添加流 video.srcObject = local Stream html <video id= "video" autoplay style= "width:600; height:400;" ></video> 复制代码`

因为我们这里只需要获得数据流，这里就不具体的解释api，我们可以去看官方文档 [MDN]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FAPI%2FMediaDevices%2FgetUserMedia ) 。 从这里可以看我们只需要一个简单的api就能获得到本地的视频和音频流，我们最后肯定是需要将这个流发送到其他的客户端的，如何发送流呢，我们通过RTCPeerConnection来进行连接以及流的传输。

> 
> 
> 
> navigator.getUserMedia
> 目前是还是支持的。但是在官方文档中已经不推荐使用，应该使用navigator.MediaDevices上的getUserMedia()，但是该api目前不是所有浏览器都支持，有兼容性问题
> 
> 
> 

为了避免兼容性问题，我们可以用以下代码来进行兼容性适配

` //浏览器不支持navigator.mediaDevices if (navigator.mediaDevices == undefined) { navigator.mediaDevices = {} navigator.mediaDevices.getUserMedia = function (constraints) { //获得旧版的getUserMedia let getUserMedia = navigator.webkitGetUserMedia || navigator.mozGetUserMedia //浏览器就不支持getUserMedia这个api，则返回个错误 if (!getUserMedia) { return Promise.reject(new Error('getUserMedia is can not use in the browser')) } // getUserMedia是异步的，所以用Promise，将返回一个绑定在navigator上的getUserMedia return new Promise((resolve, reject) => { getUserMedia.call(navigator, constraints, resolve, reject) }) } } 复制代码`

#### **RTCPeerConnection** ####

这是实现端对端(既不通过服务器进行数据交换)连接的最重要的api，这也是最难理解的一部分。

端对端的连接第一次是需要借助服务器来连接的，需要服务器来进行中转，当第一次连接上后就不需要再通过服务器了。这里我们使用socket.io,以及一点点koa，这个我们后面再讲。也有其他方式我们这里不讲有兴趣的可以看江三疯大佬的文章。总之第一次是需要服务器来实现两端的连接。

接下来是具体的交换过程

* 创建RTCPeerConnection的实例
* 交换本地和远程的sdp数据描述，使用offer和answer来进行nat穿透，建立p2p
* 交换ice网络信息，用于联网的时候的网络信息交换

创建RTCPeerConnection的实例

` let PeerConnection = window.RTCPeerConnection || window.mozRTCPeerConnection || window.webkitRTCPeerConnection let peer = new PeerConnection(iceServers) 复制代码`

这里有个参数iceServers，参数中存在两个属性，分别是stun和turn。是用于NAT穿透的，具体可以看 [WebRTC in the real world: STUN TURN and signaling]( https://link.juejin.im?target=https%3A%2F%2Fwww.html5rocks.com%2Fen%2Ftutorials%2Fwebrtc%2Finfrastructure%2F )

` { iceServers: [ { url: "stun:stun.l.google.com:19302"}, // 谷歌的公共服务 { url: "turn:***", username: ***, // 用户名 credential: *** // 密码 } ] } 复制代码`

NAT

先说下我们为什么要用NAT穿透技术才能实现p2p的连接。

> 
> 
> 
> NAT全称(Network Address Translation，网络地址转换),是用于网络的地址交换，这会导致我们得不到设备真实的ip地址
> 
> 

由于外网用的是IPV4的地址码，导致地址码的数量不够，于是就将会使用路由之类的NAT设备将外网的ip地址以及端口号都修改并使用IPV6的地址，使得多个内网可以该外网。这样增加了网络连接数量，但是却使得我们无法从内网直接找到对方的内网，所以我们需要进行NAT穿透，来实现端对端的连接。

NAT穿透的大致步骤是如A，B两端，A段向B端发送一条信息，这条信息是会被NAT设备给丢弃，但是会在NAT上留下一个洞，下次信息就可以通过这个洞来传输，同理B也这一发送一条信息，来打通自己的NAT设备。具体实现使用STUN和TURN来进行NAT穿透，该过程是通过STUN Server来进行NAT穿透，如果无法穿透则需要使用TURN Server来进行中转，具体是如何穿透的可以看 [ICE协议下NAT穿越的实现（STUN&TURN）]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F84e8c78ca61d ) ，另外我们可以搭建自己的STUN 和 TURN， [自己动手搭建 WebRTC TURN&STUN 服务器]( https://link.juejin.im?target=https%3A%2F%2Fwww.pressc.cn%2F967.html )

> 
> 
> 
> * STUN（Simple Traversal of User Datagram Protocol through Network Address
> Translators (NATs)，NAT的UDP简单穿越）是一种网络协议
> * TURN的全称为Traversal Using Relay NAT，TURN协议允许NAT或者防火墙后面的对象可以通过TCP或者UDP接收到数据
> 
> 
> 
> 

P2P

现在我们已经了解了NAT穿透，现在让我们用PeerConnection来实现p2p连接。上文中我们已经创建了PeerConnection的实例，我们称他为localPeer，remotePeer。现在我们来交换本地和远程的sdp数据描述，先上代码。

` localPeer.createOffer() .then(offer => localPeer.setLocalDescription(offer)) .then(() => remotePeer.setRemoteDescription(localPeer.localDescription)) .then(() => remotePeer.createAnswer()) .then(answer => remotePeer.setLocalDescription(answer)) .then(() => localPeer.setRemoteDescription(remotePeer.localDescription)) 复制代码`

实现交换本地和远程的sdp数据描述和我们之前的NAT穿透的步骤很像。

* localPeer调用 ` createOffer()` api来创建一个offer类型的sdp，并使用 ` setLocalDescription()` 将其添加到 ` localDescription` ，这里我们只是在本地建立p2p，不需要服务器，来第一次连接
* remotePeer接受到localPeer的 ` localDescription` ,并使用 ` setRemoteDescription` 将其添加到自己的 ` RemoteDescription`
* remotePeer通过 ` createAnswer()` 创建一个answer类型的sdp，并将其添加到自己的 ` LocalDescription`
* localPeer将remotePeer的 ` localDescription` 添加为自己的 ` remoteDescription`

到这里两端的sdp数据交换就已经完成，也就代表了本地的p2p已经连接好了，但是我们这里是在同一个界面创建了两个端，是无法真正的p2p，如果要使用网络的p2p我们就需要使用ice实现网络的对等连接，并且还需要socket.io来建立第一次数据传输

SDP

> 
> 
> 
> SDP（Session Description Protocol,会话描述协议） 它不属于传输协议，
> 但是可以使用多种的传输协议，包括会话通知协议（SAP）、会话初始协议（SIP）、实时流协议（RTSP）、MIME
> 扩展协议的电子邮件以及超文本传输协议（HTTP）。
> 
> 

这是一个具体的sdp，是本地媒体元数据，详情可以去看 [P2P通信标准协议(三)之ICE]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fpannengzhi%2Fp%2F5061674.html )

` v=0 o=- 1877521640243013583 2 IN IP4 127.0.0.1 s=- t=0 0 a=group:BUNDLE 0 1 2 a=msid-semantic: WMS m=audio 9 UDP/TLS/RTP/SAVPF 111 103 104 9 0 8 106 105 13 110 112 113 126 复制代码`

让我们再看下offer

![offer](https://user-gold-cdn.xitu.io/2019/5/13/16aaf341401e0d87?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到offer是一个offer类型的sdp，answer也是同理

ICE

> 
> 
> 
> ICE的全称为Interactive Connectivity
> Establishment,即交互式连接建立。ICE是一个用于在offer/answer模式下的NAT传输协议，主要用于UDP下多媒体会话的建立，使用了STUN协议以及TURN
> 协议
> 
> 

如果我们需要实现网络的p2p就需要进行两端的ice协议连接。这里我们需要用到

* ` RTCPeerConnection.onicecandidate()` api用于监视本地ice网络的变化，如果有了就将其使用socket.io发送出去，
* ` RTCPeerConnection.addIceCandidate()` 用于将收到的ice添加到本地的RTCPeerConnection实例中。

传输stream流 当建立好了p2p后我们可以使用RTCPeerConnection实例中的

* addstream() 添加本地的媒体流，
* onaddstream() 检测本地的媒体流，

> 
> 
> 
> onaddstream()在接送端answer的setRemoteDescription执行完成后会立即执行，也就是说我们不能在p2p创建完成后在使用addstream来添加流。
> 
> 
> 

> 
> 
> 
> addstream()和onaddstream()已经在官方文档中不推荐使用，我们最好使用更新的addTrack()和onaddTrack()，有兴趣可以看MDN
> 
> 
> 

#### **RTCDataChannel** ####

RTCDataChannel用于p2p中的数据通道，我们使用的是RTCPeerConnection中的 ` createDataChannel()` 来创建一个TCDataChannel实例。这里我们假设创建了一个实例叫channel，这里我们需要的api有

* channel.send() channel主动向已连接的通道发送数据
* ondatachannel() 监视是channel是否发生改变，比如打开(onopen)，关闭(onclose)，获得send过来的数据(onmessage)

` //发送数据hello channel.send(JSON.stringify('hello')) // 监听channel的状态 peer.ondatachannel = (event) => { var channel = event.channel channel.binaryType = 'arraybuffer' channel.onopen = (event) => { // 连接成功 console.log('channel onopen') } channel.onclose = function(event) { // 连接关闭 console.log('channel onclose') } channel.onmessage = (event) => { // 收到消息 let data = JSON.parse(event.data) console.log('channel onmessage', data) } } 复制代码`

到这里我们的webrtc基础已经写完了，我们虽然webrtc是一个不需要服务器的p2p，但是我们第一次连接是需要服务器来帮我们找到响应的端的，从而将offer，answer，ice等信息进行交互，建立p2p连接。接下来我们就使用koa和socket.io作为服务器来进行首次的连接，以及一些业务逻辑交互。

### koa&socket.io ###

koa

> 
> 
> 
> koa是一个为一个HTTP服务的中间件框架，极其的轻量级，几乎没有集成，很多功能需要我们安装插件才能使用。并且使用的是es6的语法，使用的是async来实现异步。
> 
> 
> 

我们需要创建一个server.js来部署服务器。

` import Koa from 'koa' import { join } from 'path' import Static from 'koa-static' import Socket from 'socket.io' // 创建一个socket.io const io = new Socket({ options : { pingTimeout: 10000, pingInterval: 5000 } }) // 创建koa const app = new Koa() // socket注入app io.attach(app) // 添加指定静态web文件的Static路径 // Static(root, opts) 这里将public作为根路径 app.use(Static( // join 拼接路径 // __dirname返回被执行文件夹的绝对路径 join( __dirname, './public' ) )) // 服务器端口号，这里两个listen外面的是socket.io的，后面一个是koa的listen,需要将socket监听koa的端口，不然会报错 io.listen(app.listen(3000, () => { console.log( 'server start at port: ' + 3000) })) 复制代码`

socket.io

我们先来介绍下WebSocket网络协议，他是不同于http协议的一种，具体可以看 [websocket]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMjM5MDI3MjA5MQ%3D%3D%26amp%3Bmid%3D2697266556%26amp%3Bidx%3D1%26amp%3Bsn%3D7115ba3d95e9619289287d396b5ce8da%26amp%3Bchksm%3D8376fa48b401735e1d1aad6aa659054991a5956d8c3554aa570a4705dcafbf4098595507bdb2%26amp%3Bmpshare%3D1%26amp%3Bscene%3D1%26amp%3Bsrcid%3D10130ANm40s3xS6OfetJwyGj%26amp%3Bpass_ticket%3DtPRWL )

> 
> 
> 
> socket.io是服务器使用的是WebSocket网络协议，是HTML5新增的一种通信协议，其特点是服务端可以主动向客户端推送信息，客户端也可以主动向服务端发送信息，是真正的双向平等对话，属于服务器推送技术的一种。
> 
> 
> 

这样我们就可以通过两端的主动发送打服务器，以及服务器主动发送到双端，来实现交互。 我们需要使用socket.io的api

* socket.on('event', () => {}) 监听socket触发的事件
* socket.emit('event', () => {}) 主动发送
* socket.join('room', () => {}) 加入房间
* socket.leave('room', () => {}) 离开房间
* socket.to(room | socket.id) | socket.in(room | socket.id) 指定房间，或者服务器

首先客户端和服务器端相互连接。由于服务器端设置了端口号为3000，我们的html页端的socket服务器

` // html // 引入 <script src="https://cdn.bootcss.com/socket.io/2.2.0/socket.io.js"></script> // 连接3000端口 var socket = io('ws://localhost:3000/') // server.js // 监听连接 // io是服务器端的， socket是客户端的 io.on('connection', socket => { ... }) // 监听关闭 io.on('disconnect', socket => {}) 复制代码`

我们通过socket的来实现webrtc的第一次连接

` // A 向 B 的p2p // html // A // user 是全局变量，存在sessionStorage中， 创建时候获取 var user = window.sessionStorage.user || '' // 发给服务器改socket的名称 socket.emit('createUser', 'A') // 兼容性 let PeerConnection = window.RTCPeerConnection || window.mozRTCPeerConnection || window.webkitRTCPeerConnection var peer = new PeerConnection() // 创建A端的offer peer.createOffer() .then(offer => { // 设置A端的本地描述 peer.setLocalDescription(offer, () => { // socket发送offer和房间 socket.emit('offer', {offer: offer, user: 'B'}) }) }) // 监听本地的ice变化，有则发送个B peer.onicecandidate = (event) => { if (event.candidate) { ![](https://user-gold-cdn.xitu.io/2019/6/3/16b1b606f637e98e?w=1829&h=1005&f=gif&s=4145004) // B // user 是全局变量，存在sessionStorage中， 创建时候获取 var user = window.sessionStorage.user || '' // 发给服务器改socket的名称 socket.emit('createUser', 'A') let PeerConnection = window.RTCPeerConnection || window.mozRTCPeerConnection || window.webkitRTCPeerConnection var peer = new PeerConnection() // 接受服务器端发过来的offer辨识的数据 socket.on('offer', date => { // 设置B端的远程offer 描述 peer.setRemoteDescription(data.offer, () => { // 创建B的Answer peer.createAnswer() .then(answer => { // 设置B端的本地描述 peer.setLocalDescription(answer, () => { socket.emit('answer', {answer: answer, user: 'A'}) }) }) }) }) socket.on('ice', data => { // 设置B ICE peer.addIceCandidate(data.candidate); }) socket.emit('createUser', 'B') // server.js // 用于接受客户端的用户名对应的服务器 const sockets = {} // 保存user const users = {} io.on('connection', data => { // 创建账户 socket.on('createUser', data => { let user = new User(data) users[data] = user sockets[data] = socket }) socket.on('offer', data => { // 通过B的socket的id只发送给B socket.to(sockets[data.user].id).emit('offer', data) }) socket.on('answer', data => { // 通过B的socket的id只发送给A socket.to(sockets[data.user].id).emit('answer', data) }) socket.on('ice', data => { // ice发送给B socket.to(sockets[data.user].id).emit('ice', data) }) }) 复制代码`

以上就是通过socket.io来实现p2p的第一次连接。和我们在webrtc基础的过程是一样的，只是通过了server.js来进行中转。在之后的业务逻辑中我们需要对多种不同的服务器群进行广播，这里我们来扩展下socket的广播的种类。

* io.emit() 对连接了服务器的所有客户端进行广播，比如显示房间信息
* io.to(room).emit() 对一个房间中的所有客户端进行广播，用于房间内的通知
* socket.to(room).emit() 发送个房间中除了自己以为的服务器
* socket.emit() 发送给服务器自己
* socket.to(socket.id).emit() 发送给指定的服务器

到这里关于socket.io的我们一些api的使用和使用socket.io来实现p2p我们已经了解了，接下来我们将下关于canvas实现一个画板

### canvas ###

cnavas是html5中的画板，我们可以用它来实现在html上的绘画功能，这里我们的画板也是用这个做的。 实现画板我们用一个类来进行封装，需要实现以下的功能

* 画笔，用来绘制图案
* 橡皮，清除图案
* 回退，回退到上一次绘画
* 前进，前进到下一次绘画
* 清除，清除所有的绘画几率
* 设置线条，用于设置画笔和橡皮的宽度
* 设置颜色，用于设置画笔颜色
* 操作函数，用于根据不同的操作调用不同的函数
* 回调函数，用于将事件进行回调，用于数据的传输，同步画板

所以我们可以写出我们的canvas的绘制类

` // 创建绘图类 class Draw { constructor(canvas, callBack) { this.canvas = canvas this.ctx = canvas.getContext('2d') this.width = this.canvas.width this.height = this.canvas.height this.color = color this.weight = weight this.isMove = false this.option = '' // 保存每次鼠标按下并抬起的所绘制的图片，用于撤回，前进 this.imgData = [] // 记录当前帧 this.index = 0 // 现在的坐标 this.now = [0, 0] // 移动前的坐标 this.last = [0, 0] this.bindMousemove = this.onmousemove.bind(this) this.callBack = callBack || function() {} } // 初始化 init() { } // 监听鼠标按下 onmousedown(event) { } // 监听鼠标移动 onmousemove(event) { } // 监听鼠标抬起 onmouseup() { } //绘制线条 line(last, now, weight, color) { } // 橡皮 eraser(last, now, weight) { } // 回退 back() { } // 前进 go() { } // 清除 clear() { } // 收集每一帧的图片 getImage() { } // 绘制当前帧的图片 putImage() { } // 设置尺寸 setWeight(weight) { } // 设置颜色 setColor(color) { } // 所有的操作的合集 options(option, data) { } } 复制代码`

我们来具体实现下这些方法

操作合集

` options(option, data) { switch (option) { case 'pen' : { this.line(...data) this.callBack( 'pen' , data) break } case 'eraser' : { this.eraser(...data) this.callBack( 'eraser' , data) break } case 'getImage' : { this.callBack( 'getImage' ) this.getImage() break } case 'go' : { this.callBack( 'go' ) this.go() break } case 'back' : { this.callBack( 'back' ) this.back() break } case 'clear' : { this.callBack( 'clear' ) this.clear() break } case 'setWeight' : { this.callBack( 'setWeight' , data) this.setWeight(data) break } case 'setColor' : { this.callBack( 'setColor' , data) this.setColor(data) break } } } 复制代码`

这里我们将所有操作的调用都放在一个方法中，这样有利于代码的重构，但是这样做最主要的目的是为了，当我们将每个操作的回调函数写在option方法中而不写在具体操作的方法中，这样可以避免当我们使用回调函数把参数传递出去的后，接收端使用该方法更新了自己的canvas后又会调用回调导致两端的无限回调。

画笔和橡皮

我们实现画笔的思路是当鼠标按下时，我们监听鼠标的移动，鼠标以移动就将鼠标的位置参数传递给options函数，options函数通过this.option来识别是画笔还是橡皮，调用响应的函数。当鼠标抬起时，结束移动事件的监听，并将当前帧进行保存，并且调用callback函数将保存针的信息传递出去。

` onmousedown(event) { this.last = [event.offsetX, event.offsetY] this.canvas.addEventListener( 'mousemove' , this.bindMousemove) } onmousemove(event) { this.isMove = true this.now = [event.offsetX, event.offsetY] let data = [ this.last, this.now, this.weight, this.color ] this.options( this.option, data) } onmouseup() { this.canvas.removeEventListener( 'mousemove' , this.bindMousemove) if ( this.isMove) { this.isMove = false this.options( 'getImage' ) } } line(last, now, weight, color) { this.ctx.beginPath() this.ctx.lineCap = 'round' this.ctx.lineJoin = 'round' this.ctx.lineWidth = weight this.ctx.strokeStyle = color this.ctx.moveTo(last[ 0 ], last[ 1 ]) this.ctx.lineTo(now[ 0 ], now[ 1 ]) this.ctx.closePath() this.ctx.stroke() this.last = now } eraser(last, now, weight) { this.ctx.save() this.ctx.beginPath() // console.log(now[0] , now[1]) this.ctx.arc(now[ 0 ], now[ 1 ], weight, 0 , 2 * Math.PI) this.ctx.closePath() this.ctx.clip() this.ctx.clearRect( 0 , 0 , this.width, this.height) this.ctx.fillStyle = '#fff' this.ctx.fillRect( 0 , 0 , this.width, this.height) this.ctx.restore() } 复制代码`

画笔的具体实现

* ctx.beginPath()表示开始绘制路径，并且设置下线条的特点，颜色等。
* ctx.moveTo(last[0], last[1])表示将笔的位置移动到一开始的位置，表示画笔的其实位置。
* ctx.lineTo(now[0], now[1])表示画一条从(last[0], last[1])到(now[0], now[1])一条线。
* this.ctx.closePath()关闭路径绘制
* ctx.stroke()使用线条来绘制，而不是填充
* last = now 更新坐标点

橡皮的具体实现

* ctx.save() 保存当前状态
* ctx.beginPath() 开始绘制路径
* ctx.arc(now[0], now[1], weight, 0, 2 * Math.PI) 绘制一个圆形，参数为圆心x，y，半径r，以及开始的角度，结束的角度。这里开始角度为0是从x轴的正轴开始，一圈。就相当于我们以鼠标位移结束位置绘制了一个圆。
* ctx.closePath() 关闭路径绘制
* ctx.clip() 是我们路径绘制的另外一种方法，他将我们绘制的路径进行剪切，使得我们之后的所有操作都会在这个路径绘制区域，使用clip来进行路径绘制，必须是封闭的路径
* ctx.clearRect(0, 0, this.width, this.height) 虽然这里清除整个屏幕，但是由于我们使用了clip来绘制路径，所以我们的所有只会在clip区域内生效，所以我们清除的只是我们绘制的区域，也就是橡皮檫掉的区域
* ctx.fillStyle = '#fff' ctx.fillRect(0, 0, this.width, this.height)将清除的区域填充为白色
* ctx.restore() 将之前的保存的画板重绘，其他地方就不会改变，只有橡皮檫过的地方改变。

更多细节可以看 [canvas绘制形状]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FAPI%2FCanvas_API%2FTutorial%2FDrawing_shapes )

前进和回退

前进和回退的是每当鼠标抬起时我们算一针，通过canvas的

* this.ctx.getImageData(0, 0, this.width, this.height) 参数(x, y, width, height) 这里我们把整个canvas画布进行截图得到图片，并且保存在this.imgData = [] 数组中
* 通过this.index来指定当前帧，前进就index++， 后退相反
* 通过this.ctx.putImageData(this.imgData[this.index], 0, 0)将当前帧的图片放出，使用之前需要清屏

清除，设置参数

* this.imgData = [] 清空图片数组
* this.ctx.clearRect(0, 0, this.width, this.height) 清屏
* this.index = 0 清除指针
* this.getImage() 保存第一针
* this.weight = weight 设置字体宽度
* this.color = color 上传颜色

到这里我们的canvas用到的技术已经介绍完毕

### 一对多，多对多 ###

视频模式有好几种，具体可以去在 [视频模式]( https://juejin.im/post/5cbdc145e51d456e541b4cec ) ，不同的模式处理不同的情况，不过我们这里使用的是p2p多对多的连接。因为是p2p，所以要实现多对多，那就可以变成每个的一对一。就是通过每个端都进行p2p连接。这里我们需要注意添加的顺序问题。这里我们是当有人进入房间时，进入的人和房间每一个进行p2p，已经进入的就只和进入的进行p2p。这样就可以全部都是p2p

` // nat连接方法 function createPeers ( data ) { if (user !== data.joinUser) { let conn = [data.joinUser, user].join( '-' ) if (!peers[conn]) { initPeer(conn) } } else if (data.joinUser === user) { if (data.roomusers.length > 1 ) { data.roomusers.forEach( roomuser => { if (roomuser.name !== user) { let conn = [data.joinUser, roomuser.name].join( '-' ) if (!peers[conn]) { // initPeer和之前差不多，就多了将新建的Peer和channel加入数组 initPeer(conn) } } }) } } } 复制代码`

我们在每个客户端都使用了一个数组来进行存储。通过加入的和现有的user进行标示，来标示不同的p2p。

每个p2p的具体实现

和之前单个的相同，只是我们会通过for循环来遍历数组，将每个房间内的人都会去发送offer

` // 新建对每个已经在房间的offer if (data.joinUser === user) { for ( let conn in peers) { // conn标示 createoffer(conn, peers[conn]) } } function createoffer ( conn, peer ) { peer.createOffer({ offerToReceiveAudio : 1 , offerToReceiveVideo : 1 }) .then( offer => { peer.setLocalDescription(offer, () => { console.log( 'setLocalDescription-offer' , peer.localDescription) socket.emit( 'offer' , { room : room, conn : conn, user : conn.split( '-' )[ 0 ], toUser : conn.split( '-' )[ 1 ], sdp : offer}) }) }) } 复制代码`

而在使用socket.io进行第一个连接的时候，需要通过conn标示来进行对应的传输，我们将conn进行拆分，user是发送者，touser是接受者。

` // 转发offer socket.on('offer', data => { // 通过toUser发送个其对应的socket socket.to(sockets[data.toUser].id).emit('offer', data) }) 复制代码` ` // 接收端收到offer socket.on('offer', (data) => { console.log('setRemoteDescription-offer-sdp', data.conn, data.sdp) var peer = peers[data.conn] peer.setRemoteDescription(data.sdp, () => { peer.createAnswer() .then(answer => { peer.setLocalDescription(answer, () => { console.log('setLocalDescription-answer', data.conn, answer) // 此时将发送者和接受者互换，发送answer socket.emit('answer', {room: room, user: data.toUser, toUser: data.user, conn: data.conn, sdp: answer}) }) }) }) }) 复制代码` ` // 转发answer socket.on('answer', data => { socket.to(sockets[data.toUser].id).emit('answer', data) }) 复制代码` ` // 请求端收到answer socket.on( 'answer' , (data) => { // 呼叫端设置远程 answer 描述 var peer = peers[data.conn] peer.setRemoteDescription(data.sdp, () => { console.log( 'setRemoteDescription-answer-sdp' , data.conn, data.sdp) }) }) 复制代码`

加上ice

` // 监听ICE候选信息 如果收集到，就发送给对方 peer.onicecandidate = (event) => { if (event.candidate) { socket.emit( 'ice' , {room: room, conn: conn, user: conn.split( '-' )[0], toUser: conn.split( '-' )[1], candidate: event.candidate}) } } // 转发iceCandidate socket.on( 'ice' , data => { socket.to(sockets[data.toUser].id).emit( 'ice' , data) }) // 收到Ice socket.on( 'ice' , (data) => { console.log( 'onice' , data.conn, data.candidate) var peer = peers[data.conn] console.log( '------------------------peer' ,peer) peer.addIceCandidate(data.candidate); // 设置远程 ICE }) 复制代码`

到这里我们的p2p就结束了

动态画板效果

这里我们有三种方法：

* 通过socket.io来进行主动的数据传输，不过我们这也是一对多正常的方法， 但是既然我们这次用的是webrtc那我们就不使用这种方法了。
* 通过将canvas变成数据流，并且通过addStream和onAddStream来进行，将流传输并且用video进行接受流，但是这里有个坑，由于这个坑我卡了一星期，由于我们的需求是会更改添加的流对象，但是我们之前说过onaddstream()在接送端answer的setRemoteDescription执行完成后会立即执行，所以我们不能在完成连接后在切换流对象，所以这个方法在我这个需求中是不行的
* 通过RTCDataChannel来实现，这个方法和第一个方法很像，原理就是通过主动发送数据到其他的端，其他端来在自己的canvas上进行绘画，既然我们使用的是这种方法，现在我们介绍下具体的实现流程

前面说过canva类中有个回调函数，当我们进行操作的时候，就会调用回调函数，将参数传递到类外面的sendOther()方法

* sendOther(option， data) 传递两个参数一个是option操作对应不同的方法，data数据对应方法的数据
* channels[conn].send(JSON.stringify(data)) channels[conn] 数组中对应的标示的channel，我们使用for循环就能将已经连接的所有p2p主动发送数据
* 而接收端ondatachannel会去接受发送过来的数据，根据不同option来进行操作

` peer.ondatachannel = ( event ) => { var channel = event.channel channel.binaryType = 'arraybuffer' channel.onopen = ( event ) => { // 连接成功 console.log( 'channel onopen' ) } channel.onclose = function ( event ) { // 连接关闭 console.log( 'channel onclose' ) } channel.onmessage = ( event ) => { // 收到消息 let obj = JSON.parse(event.data) let option = obj.option let data = obj.data // console.log('onmessage----------', data, option, event) if (option === 'text' ) { msgList.push(data) updateMsgList(data) } else { switch (option) { case 'pen' : { draw.line(...data) break } case 'eraser' : { draw.eraser(...data) break } case 'getImage' : { draw.getImage() break } case 'back' : { draw.back() break } case 'go' : { draw.go() break } case 'clear' : { draw.clear() break } case 'setWeight' : { draw.setWeight(...data) break } case 'setColor' : { draw.setColor(...data) break } } } // console.log('channel onmessage', e.data); } } 复制代码`

### 总结 ###

通过这次的项目还是有很多收获的，首先是webrtc领域，如果不是这次项目可能我都不会接触这个领域，也加强了我的canvas和业务逻辑的能力。用原生js写业务是真滴麻烦。 由于这段时间在写小程序，这个项目有些地方还是没有完善的，有些业务逻辑还没写完，不过核心功能已经写完了，没有太大影响。

**[Agora SDK 使用体验征文大赛 | 掘金技术征文，征文活动正在进行中]( https://juejin.im/post/5ca1fa9ff265da30b6219179 )**