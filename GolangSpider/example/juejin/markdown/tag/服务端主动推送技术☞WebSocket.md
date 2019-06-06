# 服务端主动推送技术☞WebSocket #

# 服务端主动推送技术☞WebSocket #

[toc]

# 简介 #

* 什么是WebSocket

WebSocket协议是基于TCP的一种新的网络协议。它实现了浏览器与服务器全双工(full-duplex)通信——允许服务器主动发送信息给客户端

* 实用场景

用到服务端主动推送的地方，都会使用WebSocket来实现，如：

弹幕，网页聊天系统，实时监控，股票行情推送等

* 

术语

` 单播(Unicast): 点对点，私信私聊 广播(Broadcast)(所有人): 游戏公告，发布订阅 多播，也叫组播(Multicast)（特地人群）: 多人聊天，发布订阅 复制代码`

* 

webjar

` 1、方便统一管理 2、主要解决前端框架版本不一致，文件混乱等问题 3、把前端资源，打包成jar包，借助maven工具进行管理 复制代码`
> 
> 
> 
> 既然用管理jar的方式管理js，那么这个项目肯定是没有前后端分离的。
> 对于纯前端项目，有其他方式去管理js版本与依赖。就像maven管理jar那样方便。
> 
> 

# 编写基本WebSocket服务端 #

## pom ##

* SpringBoot版本

` < parent > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-parent </ artifactId > < version > 2.1.5.RELEASE </ version > < relativePath /> <!-- lookup parent from repository --> </ parent > 复制代码`

* WebSocket依赖

` < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-websocket </ artifactId > </ dependency > 复制代码`

## 配置类:WebSocketConfig ##

` package com.example.websocket.websocketdemo01.config; import org.springframework.context.annotation.Configuration; import org.springframework.messaging.simp.config.MessageBrokerRegistry; import org.springframework.web.socket.config.annotation.EnableWebSocketMessageBroker; import org.springframework.web.socket.config.annotation.StompEndpointRegistry; import com.example.websocket.websocketdemo01.intecepter.HttpHandShakeIntecepter; import org.springframework.web.socket.config.annotation.WebSocketMessageBrokerConfigurer; @Configuration @EnableWebSocketMessageBroker public class WebSocketConfig implements WebSocketMessageBrokerConfigurer { /** * 注册端点，发布或者订阅消息的时候需要连接此端点 * setAllowedOrigins 非必须，*表示允许其他域进行连接 * withSockJS 表示开始sockejs支持 */ @Override public void registerStompEndpoints (StompEndpointRegistry registry) { registry.addEndpoint( "/endpoint-websocket" ) // .addInterceptors(new HttpHandShakeIntecepter()).setAllowedOrigins( "*" ).withSockJS(); } /** * 配置消息代理(中介) * enableSimpleBroker 服务端推送给客户端的路径前缀 * setApplicationDestinationPrefixes 客户端发送数据给服务器端的一个前缀 */ @Override public void configureMessageBroker (MessageBrokerRegistry registry) { registry.enableSimpleBroker( "/topic" , "/chat" ); registry.setApplicationDestinationPrefixes( "/app" ); } } 复制代码`

## Controller ##

` package com.example.websocket.websocketdemo01.controller.v1; import org.springframework.messaging.handler.annotation.MessageMapping; import org.springframework.messaging.handler.annotation.SendTo; import org.springframework.stereotype.Controller; import com.example.websocket.websocketdemo01.model.InMessage; import com.example.websocket.websocketdemo01.model.OutMessage; @Controller public class GameInfoController { //接收消息 @MessageMapping ( "/v1/chat" ) //发送消息 @SendTo ( "/topic/game_chat" ) public OutMessage gameInfo (InMessage message) { System.out.println( "GameInfoController->gameInfo" ); return new OutMessage(message.getContent()); } } 复制代码`

## 测试 ##

* 管理端

` http://localhost:8080/v1/admin.html`

* 客户端

` http://localhost:8080/v1/index.html`

连接上服务器后，管理端发送的内容会显示在客户端

## 小结 ##

至此，我们的基本的服务端就算编写完毕了

当客户端连接上 ` /endpoint-websocket` 后，可以往 ` /v1/chat` 发送消息，并且监听 ` /topic/game_chat` ，服务端会将消息发往 ` /topic/game_chat`

任何客户端，只要监听了 ` /topic/game_chat` ，就会收到这个推送

# 服务端主动推送消息 #

在上一章中，服务端通过接收前端的WebSocket请求进行响应，其实还是一个请求响应推送，只不过这个过程中链接不断来。

当我们使用WebSocket的时候，更多情况下都是服务端被客户端连接上后进行主动推送，这个时候该怎么做呢？

` @Controller public class GameInfoController { @Autowired private SimpMessagingTemplate template; @GetMapping ( "/v1/chat/http" ) @ResponseBody public OutMessage gameInfoHttp (InMessage message) { System.out.println( "gameInfoHttp" ); OutMessage outMessage = new OutMessage(message.getContent()); template.convertAndSend( "/topic/game_chat" , new OutMessage(message.getContent())); return outMessage; } } 复制代码`

只要使用 ` SimpMessagingTemplate` ，就可以往指定 ` destination` 发送特定的数据，只要监听了这个 ` destination` 的客户端都会收到

## 测试 ##

访问接口： ` http://localhost:8080/v1/chat/http?from=1&to=2&content=哇哈哈`

可以看到 ` http://localhost:8080/v1/index.html` 收到的服务端的推送

# 关于客户端连接地址:ws or http #

因为我们后端使用的是stomp协议，所以此时客户仍旧使用http进行连接

如果要使用 ` ws` 进行连接，那么后端做修改

![](https://user-gold-cdn.xitu.io/2019/5/25/16aef2ce13a291c9?imageView2/0/w/1280/h/960/ignore-error/1) ￼

去掉withSocketJs
前端做修改

![](https://user-gold-cdn.xitu.io/2019/5/25/16aef2d6cc1f6943?imageView2/0/w/1280/h/960/ignore-error/1) ￼
stompClient的构建方式发生了点变化

# ` @SendTo` 与 ` SimpMessagingTemplate` #

` @SendTo` 不够通用，固定发送给指定的订阅者

` SimpMessagingTemplate` 比较灵活

` simpMessagingTemplate.convertAndSend( "/topic/game_chat" , new OutMessage(message.getContent())); 复制代码`

可以动态的指定要发送给谁

# SpringBoot对WebSocket的监听 #

## 连接监听 ##

` package com.example.websocket.websocketdemo01.listener; import org.springframework.context.ApplicationListener; import org.springframework.messaging.simp.stomp.StompHeaderAccessor; import org.springframework.stereotype.Component; import org.springframework.web.socket.messaging.SessionConnectEvent; @Component public class ConnectEventListener implements ApplicationListener < SessionConnectEvent > { @Override public void onApplicationEvent (SessionConnectEvent event) { StompHeaderAccessor headerAccessor = StompHeaderAccessor.wrap(event.getMessage()); System.out.println( "【ConnectEventListener监听器事件 类型】" +headerAccessor.getCommand().getMessageType()); } } 复制代码`

当客户端连接的时候，会触发 ` CONNECT` 事件

## 订阅监听 ##

` package com.example.websocket.websocketdemo01.listener; import org.springframework.context.ApplicationListener; import org.springframework.messaging.simp.stomp.StompHeaderAccessor; import org.springframework.stereotype.Component; import org.springframework.web.socket.messaging.SessionSubscribeEvent; @Component public class SubscribeEventListener implements ApplicationListener < SessionSubscribeEvent > { /** * 在事件触发的时候调用这个方法 * * StompHeaderAccessor 简单消息传递协议中处理消息头的基类， * 通过这个类，可以获取消息类型(例如:发布订阅，建立连接断开连接)，会话id等 * */ @Override public void onApplicationEvent (SessionSubscribeEvent event) { StompHeaderAccessor headerAccessor = StompHeaderAccessor.wrap(event.getMessage()); System.out.println( "【SubscribeEventListener监听器事件 类型】" +headerAccessor.getCommand().getMessageType()); System.out.println( "【SubscribeEventListener监听器事件 sessionId】" +headerAccessor.getSessionAttributes().get( "sessionId" )); } } 复制代码`

当客户端连接的时候，会触发 ` SUBSCRIBE` 事件

取消监听事件： ` SessionUnsubscribeEvent`

## 客户端断开监听 ##

` package com.example.websocket.websocketdemo01.listener; import org.springframework.context.ApplicationListener; import org.springframework.messaging.simp.stomp.StompHeaderAccessor; import org.springframework.stereotype.Component; import org.springframework.web.socket.messaging.SessionDisconnectEvent; import org.springframework.web.socket.messaging.SessionSubscribeEvent; @Component public class DissconnectEventListener implements ApplicationListener < SessionDisconnectEvent > { /** * 在事件触发的时候调用这个方法 * * StompHeaderAccessor 简单消息传递协议中处理消息头的基类， * 通过这个类，可以获取消息类型(例如:发布订阅，建立连接断开连接)，会话id等 * */ @Override public void onApplicationEvent (SessionDisconnectEvent sessionDisconnectEvent) { StompHeaderAccessor headerAccessor = StompHeaderAccessor.wrap(sessionDisconnectEvent.getMessage()); System.out.println( "【SubscribeEventListener监听器事件 类型】" +headerAccessor.getCommand().getMessageType()); System.out.println( "【SubscribeEventListener监听器事件 sessionId】" +headerAccessor.getSessionAttributes().get( "sessionId" )); } } 复制代码`

当客户端连接的时候，会触发 ` DISCONNECT` 事件

## 获取客户端id ##

` @Override public void onApplicationEvent (SessionConnectEvent event) { StompHeaderAccessor headerAccessor = StompHeaderAccessor.wrap(event.getMessage()); System.out.println( "【ConnectEventListener监听器事件 类型】" +headerAccessor.getCommand().getMessageType()); System.out.println( "simpSessionId\t" +headerAccessor.getHeader( "simpSessionId" )); } 复制代码`

这个 ` SimpSessionId` 会在客户端连接、订阅、断线等情况下获取到，可以用于标记客户端

而在上面的监听器中，我们使用

` System.out.println( "【SubscribeEventListener监听器事件 sessionId】" +headerAccessor.getSessionAttributes().get( "sessionId" )); 复制代码`

来获取sessionId，这个好需要我们编写拦截器，将sessionId手动放到这个SessionAttributes，才能取到。

## 拦截器 ##

` package com.example.websocket.websocketdemo01.intecepter; import java.util.Map; import javax.servlet.http.HttpSession; import org.springframework.http.server.ServerHttpRequest; import org.springframework.http.server.ServerHttpResponse; import org.springframework.http.server.ServletServerHttpRequest; import org.springframework.web.socket.WebSocketHandler; import org.springframework.web.socket.server.HandshakeInterceptor; public class HttpHandShakeIntecepter implements HandshakeInterceptor { @Override public boolean beforeHandshake (ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler, Map<String, Object> attributes) throws Exception { System.out.println( "【握手拦截器】beforeHandshake" ); if (request instanceof ServletServerHttpRequest) { ServletServerHttpRequest servletRequest = (ServletServerHttpRequest)request; HttpSession session = servletRequest.getServletRequest().getSession(); String sessionId = session.getId(); System.out.println( "【握手拦截器】beforeHandshake sessionId=" +sessionId); attributes.put( "sessionId" , sessionId); } return true ; } @Override public void afterHandshake (ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler, Exception exception) { System.out.println( "【握手拦截器】afterHandshake" ); if (request instanceof ServletServerHttpRequest) { ServletServerHttpRequest servletRequest = (ServletServerHttpRequest)request; HttpSession session = servletRequest.getServletRequest().getSession(); String sessionId = session.getId(); System.out.println( "【握手拦截器】afterHandshake sessionId=" +sessionId); } } } 复制代码`

* 注册拦截器

![](https://user-gold-cdn.xitu.io/2019/5/25/16aef2d6ccdf29b6?imageView2/0/w/1280/h/960/ignore-error/1) ￼

在拦截器中，我们把sessionId放在了session中，此时WebSocket的监听器就能够获取到sessionId

` headerAccessor.getSessionAttributes().get( "sessionId" ) 复制代码`

# 点对点发送 #

客户端订阅的端，需要唯一，这个需要通过客户端通过参数传递上来

` template.convertAndSend( "/chat/single/" +message.getTo(), new OutMessage(message.getFrom()+ " 发送:" + message.getContent())); 复制代码`

同样的，对于组播来说，只要保证客户端的订阅频道是同一组的就行。

一般对于 ` 组` 与 ` 点` 的定义，都会通过业务来处理

# Nginx反向代理WebSocket #

` http { map $http_upgrade $connection_upgrade { default upgrade; '' close; } upstream websocket { ip_hash; #使用ip固定转发到后端服务器 server localhost:3100; server localhost:3101; server localhost:3102; } server { listen 8020; location / { proxy_pass http://websocket; proxy_http_version 1.1; proxy_set_header Upgrade $http_upgrade ; proxy_set_header Connection $connection_upgrade ; # 声明支持websocket } } } #http/2 nginx conf #server{ # listen 443; # server_name example.com www.example.com; # root /Users/welefen/Develop/git/firekylin/www; # set $node_port 8360; # ssl on; # ssl_certificate %path/ssl/chained.pem; # ssl_certificate_key %path/ssl/domain.key; # ssl_session_timeout 5m; # ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # ssl_ciphers ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA; # ssl_session_cache shared:SSL:50m; # ssl_dhparam %path/ssl/dhparams.pem; # ssl_prefer_server_ciphers on; # index index.js index.html index.htm; # location ^~ /.well-known/acme-challenge/ { # alias %path/ssl/challenges/; # try_files $uri = 404; # } # location / { # proxy_http_version 1.1; # proxy_set_header X-Real-IP $remote_addr; # proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; # proxy_set_header Host $http_host; # proxy_set_header X-NginX-Proxy true; # proxy_set_header Upgrade $http_upgrade; # proxy_set_header Connection "upgrade"; # proxy_pass http://127.0.0.1:$node_port$request_uri; # proxy_redirect off; # } # location = /development.js { # deny all; # } # location = /testing.js { # deny all; # } # location = /production.js { # deny all; # } # location ~ /static/ { # etag on; # expires max; # } #} #server { # listen 80; # server_name example.com www.example.com; # rewrite ^(.*) https://example.com$1 permanent; #} 复制代码`

# 阶段1源码 #

截止当前的所有代码位于： [码云]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Flearn_from%2Fwebsocket%2Ftree%2Fmaster%2Fdemo01-websocket-base )

# 进阶 #

## 使用ServerEndpoint的方式编写 ##

[源码]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Flearn_from%2Fwebsocket%2Ftree%2Fmaster%2Fdemo02-websocket-base )