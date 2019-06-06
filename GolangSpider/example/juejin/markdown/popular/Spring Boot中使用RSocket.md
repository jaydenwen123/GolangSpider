# Spring Boot中使用RSocket #

## 1. 概述 ##

` RSocket` 应用层协议支持 ` Reactive Streams` 语义， 例如：用RSocket作为HTTP的一种替代方案。在本教程中， 我们将看到 ` RSocket` 用在spring boot中，特别是spring boot 如何帮助抽象出更低级别的RSocket API。

## 2. 依赖 ##

让我们从添加 ` spring-boot-starter-rsocket` 依赖开始：

` < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-rsocket </ artifactId > </ dependency > 复制代码`

这个依赖会传递性的拉取 ` RSocket` 相关的依赖，比如： ` rsocket-core` 和 ` rsocket-transport-netty`

## 3.示例的应用程序 ##

现在继续我们的简单应用程序。为了突出 ` RSocket` 提供的交互模式，我打算创建一个交易应用程序， 交易应用程序包括客户端和服务器。

### 3.1. 服务器设置 ###

首先，我们设置由springboot应用程序引导的 ` RSocket server` 服务器。 因为有 ` spring-boot-starter-rsocket dependency` 依赖，所以springboot会自动配置 ` RSocket server` 。 跟平常一样， 可以用属性驱动的方式修改 ` RSocket server` 默认配置值。例如：通过增加如下配置在 ` application.properties` 中，来修改 ` RSocket` 端口

` spring.rsocket.server.port=7000 复制代码`

也可以根据需要进一步修改服务器的其他属性

### 3.2.设置客户端 ###

接下来，我们来设置客户端，也是一个springboot应用程序。虽然springboot自动配置大部分RSocket相关的组件，但还要自定义一些对象来完成设置。

` @Configuration public class ClientConfiguration { @Bean public RSocket rSocket () { return RSocketFactory .connect() .mimeType(MimeTypeUtils.APPLICATION_JSON_VALUE, MimeTypeUtils.APPLICATION_JSON_VALUE) .frameDecoder(PayloadDecoder.ZERO_COPY) .transport(TcpClientTransport.create(7000)) .start() .block(); } @Bean RSocketRequester rSocketRequester(RSocketStrategies rSocketStrategies) { return RSocketRequester.wrap(rSocket(), MimeTypeUtils.APPLICATION_JSON, rSocketStrategies); } } 复制代码`

这儿我们正在创建 ` RSocket` 客户端并且配置TCP端口为：7000。注意： 该服务端口我们在前面已经配置过。 接下来我们定义了一个RSocket的装饰器对象 ` RSocketRequester` 。 这个对象在我们跟 ` RSocket server` 交互时会为我们提供帮助。 定义这些对象配置后，我们还只是有了一个骨架。在接下来，我们将暴露不同的交互模式， 并看看springboot在这个地方提供帮助的。

## 4. ` springboot RSocket` 中的 ` Request/Response` ##

我们从 ` Request/Response` 开始， ` HTTP` 也使用这种通信方式，这也是最常见的、最相似的交互模式。 在这种交互模式里， 由客户端初始化通信并发送一个请求。之后，服务器端执行操作并返回一个响应给客户端--这时通信完成。 在我们的交易应用程序里， 一个客户端询问一个给定的股票的当前的市场数据。 作为回复，服务器会传递请求的数据。

### 4.1.服务器 ###

在服务器这边，我们首先应该创建一个 ` controller` 来持有我们的处理器方法。 我们会使用 ` @MessageMapping` 注解来代替像SpringMVC中的 ` @RequestMapping` 或者 ` @GetMapping` 注解

` @Controller public class MarketDataRSocketController { private final MarketDataRepository marketDataRepository; public MarketDataRSocketController(MarketDataRepository marketDataRepository) { this.marketDataRepository = marketDataRepository; } @MessageMapping( "currentMarketData" ) public Mono<MarketData> currentMarketData(MarketDataRequest marketDataRequest) { return marketDataRepository.getOne(marketDataRequest.getStock()); } } 复制代码`

来研究下我们的控制器。 我们将使用 ` @Controller` 注解来定义一个控制器来处理进入RSocket的请求。 另外，注解 ` @MessageMapping` 让我们定义我们感兴趣的路由和如何响应一个请求。 在这个示例中， 服务器监听路由 ` currentMarketData` ， 并响应一个单一的结果 ` Mono<MarketData>` 给客户端。

### 4.2. 客户端 ###

接下来， 我们的RSocket客户端应该询问一只股票的价格并得到一个单一的响应。 为了初始化请求， 我们该使用 ` RSocketRequester` 类，如下：

` @RestController public class MarketDataRestController { private final RSocketRequester rSocketRequester; public MarketDataRestController(RSocketRequester rSocketRequester) { this.rSocketRequester = rSocketRequester; } @GetMapping(value = "/current/{stock}" ) public Publisher<MarketData> current(@PathVariable( "stock" ) String stock) { return rSocketRequester .route( "currentMarketData" ) .data(new MarketDataRequest(stock)) .retrieveMono(MarketData.class); } } 复制代码`

注意：在示例中， ` RSocket` 客户端也是一个 ` REST` 风格的 ` controller` ，以此来访问我们的 ` RSocket` 服务器。因此，我们使用 ` @RestController` 和 ` @GetMapping` 注解来定义我们的请求/响应端点。 在端点方法中， 我们使用的是类 ` RSocketRequester` 并指定了路由。 事实上，这个是服务器端 ` RSocket` 所期望的路由，然后我们传递请求数据。最后，当调用 ` retrieveMono()` 方法时，springboot会帮我们初始化一个请求/响应交互。

## 5. ` Spring Boot RSocket` 中的 ` Fire And Forget` 模式 ##

接下来我们将查看 ` Fire And Forget` 交互模式。正如名字提示的一样，客户端发送一个请求给服务器，但是不期望服务器的返回响应回来。 在我们的交易程序中， 一些客户端会作为数据资源服务，并且推送市场数据给服务器端。

### 5.1.服务器端 ###

我们来创建另外一个端点在我们的服务器应用程序中，如下：

` @MessageMapping ( "collectMarketData" ) public Mono<Void> collectMarketData (MarketData marketData) { marketDataRepository.add(marketData); return Mono.empty(); } 复制代码`

我们又一次定义了一个新的 ` @MessageMapping` 路由为 ` collectMarketData` 。此外， Spring Boot自动转换传入的负载为一个 ` MarketData` 实例。 但是，这儿最大的不同是我们返回一个 ` Mono<Void>` ，因为客户端不需要服务器的返回。

### 5.2. 客户端 ###

来看看我们如何初始化我们的 ` fire-and-forget` 模式的请求。 我们将创建另外一个REST风格的端点，如下：

` @GetMapping(value = "/collect" ) public Publisher<Void> collect () { return rSocketRequester .route( "collectMarketData" ) .data(getMarketData()) .send(); } 复制代码`

这儿我们指定路由和负载将是一个 ` MarketData` 实例。 由于我们使用 ` send()` 方法来代替 ` retrieveMono()` ，所有交互模式变成了 ` fire-and-forget` 模式。

## 6. ` Spring Boot RSocket` 中的 ` Request Stream` ##

请求流是一种更复杂的交互模式， 这个模式中客户端发送一个请求，但是在一段时间内从服务器端获取到多个响应。 为了模拟这种交互模式， 客户端会询问给定股票的所有市场数据。

### 6.1.服务器端 ###

我们从服务器端开始。 我们将添加另外一个消息映射方法，如下：

` @MessageMapping( "feedMarketData" ) public Flux<MarketData> feedMarketData(MarketDataRequest marketDataRequest) { return marketDataRepository.getAll(marketDataRequest.getStock()); } 复制代码`

正如所见， 这个处理器方法跟其他的处理器方法非常类似。 不同的部分是我们返回一个 ` Flux<MarketData>` 来代替 ` Mono<MarketData>` 。 最后我们的RSocket服务器会返回多个响应给客户端。

### 6.2.客户端 ###

在客户端这边， 我们该创建一个端点来初始化请求/响应通信，如下：

` @GetMapping(value = "/feed/{stock}" , produces = MediaType.TEXT_EVENT_STREAM_VALUE) public Publisher<MarketData> feed(@PathVariable( "stock" ) String stock) { return rSocketRequester .route( "feedMarketData" ) .data(new MarketDataRequest(stock)) .retrieveFlux(MarketData.class); } 复制代码`

我们来研究下RSocket请求。 首先我们定义了路由和请求负载。 然后，我们定义了使用 ` retrieveFlux()` 调用的响应期望。这部分决定了交互模式。 另外注意：由于我们的客户端也是 ` REST` 风格的服务器，客户端也定义了响应媒介类型 ` MediaType.TEXT_EVENT_STREAM_VALUE` 。

### 7.异常的处理 ###

现在让我们看看在服务器程序中，如何以声明式的方式处理异常。 当处理请求/响应式， 我可以简单的使用 ` @MessageExceptionHandler` 注解，如下：

` @MessageExceptionHandler public Mono<MarketData> handleException(Exception e) { return Mono.just(MarketData.fromException(e)); } 复制代码`

这里我们给异常处理方法标记注解为 ` @MessageExceptionHandler` 。作为结果， 这个方法将处理所有类型的异常， 因为 ` Exception` 是所有其他类型的异常的超类。 我们也可以明确地创建更多的不同类型的，不同的异常处理方法。 这当然是请求/响应模式，并且我们返回的是 ` Mono<MarketData>` 。我们期望这里的响应类型跟我们的交互模式的返回类型相匹配。

## 8.总结 ##

在本教程中， 我们介绍了springboot的RSocket支持，并详细列出了RSocket提供的不同交互模式。查看所有示例代码在 [GitHub]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feugenp%2Ftutorials%2Ftree%2Fmaster%2Fspring-5-webflux ) 上。

> 
> 
> 
> 原文链接： [www.baeldung.com/spring-boot…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.baeldung.com%2Fspring-boot-rsocket
> )
> 
> 

> 
> 
> 
> 作者： [baeldung](
> https://link.juejin.im?target=https%3A%2F%2Fwww.baeldung.com%2Fauthor%2Fbaeldung%2F
> )
> 
> 

> 
> 
> 
> 译者：sleeve
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25a1d46727085?imageView2/0/w/1280/h/960/ignore-error/1)