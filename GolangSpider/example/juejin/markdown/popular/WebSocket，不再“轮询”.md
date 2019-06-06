# WebSocket，不再“轮询” #

## 1. 前言 ##

本文先讲解WebSocket的应用场景和特点，然后通过前后端示例代码讲解，展示在实际的开发中的应用。

### 1.1. 应用场景 ###

WebSocket是一种在单个TCP连接上进行全双工通信的协议, 是为了满足基于 Web 的日益增长的实时通信需求而产生的。我们平时接触的大多数是HTTP的接口，但是在有些业务场景中满足不了我们的需求，这时候就需要用到WebSocket。简单举两个例子：

（1） 页面地图上要实时显示在线人员坐标：传统基于HTTP接口的处理方式是轮询，每次轮询更新最新的坐标信息。

（2）手机的付款码页面，在外界设备扫描付款码支付成功后，手机付款码页面提示“支付成功”并自动关闭：传统方式还是轮询，付款码页面一直调用接口，直到从服务器获取成功支付的状态后，手机提示“支付成功”并关闭付款码页面。

HTTP 协议有一个缺陷：通信只能由客户端发起。这种单向请求的特点，注定了如果服务器有连续的状态变化，客户端要获知就非常麻烦。我们只能使用"轮询"：每隔一段时候，就发出一个询问，了解服务器有没有新的信息。但这种方式即浪费带宽（HTTP HEAD 是比较大的），又消耗服务器 CPU 占用（没有信息也要接受请求）。

在WebSocket API尚未被众多浏览器实现和发布的时期，开发者在开发需要接收来自服务器的实时通知应用程序时，不得不求助于一些“hacks”来模拟实时连接以实现实时通信，最流行的一种方式是长轮询 。 长轮询主要是发出一个HTTP请求到服务器，然后保持连接打开以允许服务器在稍后的时间响应（由服务器确定）。为了这个连接有效地工作，许多技术需要被用于确保消息不错过，如需要在服务器端缓存和记录多个的连接信息（每个客户）。虽然长轮询是可以解决这一问题的，但它会耗费更多的资源，如CPU、内存和带宽等，要想很好的解决实时通信问题就需要设计和发布一种新的协议

### 1.2. WebSocket定义 ###

WebSocket是一种协议，是一种与HTTP 同等的网络协议，两者都是应用层协议，都基于 TCP 协议。但是 WebSocket 是一种双向通信协议，在建立连接之后，WebSocket 的 server 与 client 都能主动向对方发送或接收数据。同时，WebSocket在建立连接时需要借助 HTTP 协议，连接建立好了之后 client 与 server 之间的双向通信就与 HTTP 无关了。

相比于传统HTTP 的每次“请求-应答”都要client 与 server 建立连接的模式，WebSocket 是一种长连接的模式。就是一旦WebSocket 连接建立后，除非client 或者 server 中有一端主动断开连接，否则每次数据传输之前都不需要HTTP 那样请求数据。

WebSocket 对象提供了一组 API，用于创建和管理 WebSocket 连接，以及通过连接发送和接收数据。浏览器提供的WebSocket API很简洁，调用示例如下：

` var ws = new WebSocket( 'wss://example.com/socket' ); // 创建安全WebSocket 连接（wss） ws.onerror = function (error) { ... } // 错误处理 ws.onclose = function () { ... } // 关闭时调用 ws.onopen = function () { ws.send( "Connection established. Hello server!" );} // 连接建立时调用向服务端发送消息 ws.onmessage = function (msg) { ... }// 接收服务端发送的消息 复制代码`

HTTP、WebSocket 等应用层协议，都是基于 TCP 协议来传输数据的。我们可以把这些高级协议理解成对 TCP 的封装。既然大家都使用 TCP 协议，那么大家的连接和断开，都要遵循 TCP 协议中的三次握手和四次握手 ，只是在连接之后发送的内容不同，或者是断开的时间不同。对于 WebSocket 来说，它必须依赖 HTTP 协议进行一次握手 ，握手成功后，数据就直接从 TCP 通道传输，与 HTTP 无关了。

## 2. 后端WebSocket服务（SpringBoot） ##

pom.xml

` <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-websocket</artifactId> </dependency> 复制代码`

在配置类@Configuration下注册ServerEndpointExporter的Bean，这个bean会自动注册使用了@ServerEndpoint注解声明的Websocket endpoint

` @Bean public ServerEndpointExporter serverEndpointExporter (){ return new ServerEndpointExporter(); } 复制代码`

创建WebSocket的工具类WebSocket.java

` import org.springframework.stereotype.Component; import javax.websocket.OnClose; import javax.websocket.OnMessage; import javax.websocket.OnOpen; import javax.websocket.Session; import javax.websocket.server.PathParam; import javax.websocket.server.ServerEndpoint; import java.util.HashMap; import java.util.Map; import java.util.concurrent.CopyOnWriteArraySet; @Component @ServerEndpoint( "/websocket/{userId}" ) public class WebSocket { private Session session; private static CopyOnWriteArraySet<WebSocket> webSockets =new CopyOnWriteArraySet<>(); private static Map<String,Session> sessionPool = new HashMap<String,Session>(); @OnOpen public void onOpen(Session session, @PathParam(value= "userId" )String userId) { this.session = session; webSockets.add(this); sessionPool.put(userId, session); System.out.println( "【websocket消息】有新的连接，总数为:" +webSockets.size()); } @OnClose public void onClose () { webSockets.remove(this); System.out.println( "【websocket消息】连接断开，总数为:" +webSockets.size()); } @OnMessage public void onMessage(String message) { System.out.println( "【websocket消息】收到客户端消息:" +message); } // 此为广播消息 public void sendAllMessage(String message) { for (WebSocket webSocket : webSockets) { System.out.println( "【websocket消息】广播消息:" +message); try { webSocket.session.getAsyncRemote().sendText(message); } catch (Exception e) { e.printStackTrace(); } } } // 此为单点消息 public void sendOneMessage(String userId, String message) { Session session = sessionPool.get(userId); if (session != null) { try { session.getAsyncRemote().sendText(message); } catch (Exception e) { e.printStackTrace(); } } } } 复制代码`

到此WebSocket的代码就结束了，运行该SpringBoot项目，对应的WebSocket地址为：ws://127.0.0.1:port/websocket/{userId}

可以在 [WebSocket在线测试网站]( https://link.juejin.im?target=http%3A%2F%2Fwww.blue-zero.com%2FWebSocket%2F ) 上测试后端接口。

## 3. 前端WebSocket调用（Angular） ##

### 3.1. npm依赖 ###

安装 rxjs 的依赖库

` npm install rxjs@6.3.3 --save //示例代码，也可以装其他版本 复制代码`

安装websocket 依赖库

` npm install ws --save 复制代码`

安装类型定义文件

` npm install @types/ws --save-dev 复制代码`

### 3.2. WebSocket Service ###

创建 WebSocket 的Service文件

` ng g service websocket 复制代码`

上述命令生成了websocket.service.ts文件，示例代码为：

` import { Injectable } from '@angular/core' ; import {Subject, Observer, Observable} from 'rxjs' ; import { ThrowStmt } from '@angular/compiler' ; @Injectable({ providedIn: 'root' }) export class WebsocketService { ws:WebSocket; constructor () { } createObservableSocket(url:string):Observable<any>{ this.ws=new WebSocket(url); return new Observable( observer=>{ this.ws.onmessage=(event)=>observer.next(event.data); this.ws.onerror=(event)=>observer.error(event); this.ws.onclose=(event)=> observer.complete(); } ); } sendMessage(message:string){ this.ws.send(message); } } 复制代码`

### 3.3. Demo演示 ###

简单做个demo页面，有留言板和输入框。同时开多个浏览器页面，只要在任意一个页面的输入框中输入文字，所有页面的留言板上都会实时显示出来。

![avatar](https://user-gold-cdn.xitu.io/2019/6/5/16b25cdb64a6dbe1?imageView2/0/w/1280/h/960/ignore-error/1)

示例的代码提供，app.component.html

` <div class= "output-window" > <div *ngFor= "let mv of messageValue;let i=index" > {{i+1}}. {{mv}} </div> </div> <nz-input-group class= "input-window" nzSearch nzSize= "large" [nzAddOnAfter]= "suffixButton" > <input type = "text" nz-input placeholder= "发送消息..." [(ngModel)]= "sendValue" (keyup.enter)= "sendMessage()" /> </nz-input-group> <ng-template #suffixButton> <button nz-button nzType= "primary" nzSize= "large" nzSearch (click)= "sendMessage()" >发送</button> </ng-template> 复制代码`

app.component.ts

` import { Component, OnInit, Input } from '@angular/core' ; import { WebsocketService } from './websocket.service' ; @Component({ selector: 'app-root' , templateUrl: './app.component.html' , styleUrls: [ './app.component.css' ] }) export class AppComponent implements OnInit{ title = 'websocket-demo-web' ; messageValue=[]; sendValue= "" ; constructor(private websocketService:WebsocketService){} ngOnInit (){ this.websocketService.createObservableSocket( "ws://127.0.0.1:port/websocket/kerry" ) .subscribe( data=>{console.log(data);this.messageValue.push(data) ;}, error=>console.log(error), ()=>console.log( "已结束" ) ); } sendMessage (){ this.websocketService.sendMessage(this.sendValue); this.sendValue=null; } } 复制代码`

app.component.css

`.output-window{ width: 60%; height: 250px; background-color: black; color: chartreuse; margin-left: 20%; } .input-window{ width: 60%; height: 30px; background-color: #c4c4c40d; margin-left: 20%; margin-top: 20px; } .input-button{ height: 30px; width: 60px; margin-left: 10px; } 复制代码`

本人创业团队产品MadPecker，主要做BUG管理、测试管理、应用分发，网址:www.madpecker.com，有需要的朋友欢迎试用、体验！

本文为MadPecker团队技术人员编写，转载请标明出处