# go-kit微服务：JWT身份认证 #

为了保证系统安全稳定，保护用户数据安全，服务中一般引入身份认证手段，对用户的请求进行安全拦截、校验与过滤。常用的身份认证方式有：

* JWT： JWT提供了一种用于发布接入令牌（Access Token)，并对发布的签名接入令牌进行验证的方法。 令牌（Token）本身包含了一系列声明，应用程序可以根据这些声明限制用户对资源的访问。
* OAuth2：OAuth2是一种授权框架，提供了一套详细的授权机制（指导）。用户或应用可以通过公开的或私有的设置，授权第三方应用访问特定资源。
* Basic：即用户名和密码认证，每次请求都携带，不安全。

## 实战演练 ##

本文将在go-kit微服务中引入jwt验证机制，实现token的签发与验证。关于jwt的原理不再阐述，大家可到最后的参考文献中查阅。简单说下实现思路：

* 新建登录接口，验证用户名和密码。验证通过生成token，返回客户端。
* 为calculate接口增加token验证机制，这里使用go-kit提供的中间件进行封装。
* 本示例使用第三方jwt的go实现 ` dgrijalva/jwt-go`

### Step-1：代码准备 ###

复制目录arithmetic_circuitbreaker_demo，重命名为arithmetic_jwt_demo，重命名register目录为service。

安装依赖的jwt第三方库：

` go get github.com/dgrijalva/jwt-go 复制代码`

### Step-2：创建jwt.go ###

在 ` service` 目录下创建文件 ` jwt.go` 。

* 首先定义生成token时需要的密钥（这个是jwt中最重要的东西，千万不能泄露）。
* 自定义声明。在 ` StandardClaims` 基础上增加了 ` UserId` 、 ` Name` 两个字段，可以根据实际需要扩展其他字段，如角色。
* 定义 ` keyfunc` ，该方法在验证token时作为回调函数使用，后面会有描述。
* 定义生成token的方法 ` Sign` 。这里直接调用jwt第三方库生成，为了演示方便设置token的过期时间为2分钟。

如下为 ` jwt.go` 的全部代码：

` //secret key var secretKey = []byte( "abcd1234!@#$" ) // ArithmeticCustomClaims 自定义声明 type ArithmeticCustomClaims struct { UserId string `json: "userId" ` Name string `json: "name" ` jwt.StandardClaims } // jwtKeyFunc 返回密钥 func jwtKeyFunc(token *jwt.Token) (interface{}, error) { return secretKey, nil } // Sign 生成token func Sign(name, uid string) (string, error) { //为了演示方便，设置两分钟后过期 expAt := time.Now().Add(time.Duration(2) * time.Minute).Unix() // 创建声明 claims := ArithmeticCustomClaims{ UserId: uid, Name: name, StandardClaims: jwt.StandardClaims{ ExpiresAt: expAt, Issuer: "system" , }, } //创建token，指定加密算法为HS256 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //生成token return token.SignedString(secretKey) } 复制代码`

### Step-3：新增登录接口 ###

Service层：新增登录接口，按照go-kit的架构方式，依次：

* 在接口Service中新增Login方法。
* 在ArithmeticService中实现Login方法：验证用户名和密码，调用jwt的Sign方法生成token；
* 在loggingMiddleware中实现Login方法；
* 在metricMiddleware中实现Login方法；

` // Service Define a service interface type Service interface { //…… // HealthCheck Login(name, pwd string) (string, error) } func (s ArithmeticService) Login(name, pwd string) (string, error) { if name == "name" && pwd == "pwd" { token, err := Sign(name, pwd ) return token, err } return "" , errors.New( "Your name or password dismatch" ) } 复制代码`

Endpoint层：新增登录接口所需的请求和响应实体结构，编写创建Endpoint的方法。

` // AuthRequest type AuthRequest struct { Name string `json: "name" ` Pwd string `json: "pwd" ` } // AuthResponse type AuthResponse struct { Success bool `json: "success" ` Token string `json: "token" ` Error string `json: "error" ` } func MakeAuthEndpoint(svc Service) endpoint.Endpoint { return func(ctx context.Context, request interface{}) (response interface{}, err error) { req := request.(AuthRequest) token, err := svc.Login(req.Name, req.Pwd) var resp AuthResponse if err != nil { resp = AuthResponse{ Success: err == nil, Token: token, Error: err.Error(), } } else { resp = AuthResponse{ Success: err == nil, Token: token, } } return resp, nil } } 复制代码`

Transport层：编写decode和encode方法，新增登录接口路由。同时，对calculate接口增加token检测逻辑：在请求处理之前，从HTTP请求头中读取认证信息，若读取成功则加入请求上下文。这里直接使用go-kit提供的HTTPToContext方法。

` func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) { var loginRequest AuthRequest if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil { return nil, err } return loginRequest, nil } func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error { w.Header().Set( "Content-Type" , "application/json;charset=utf-8" ) return json.NewEncoder(w).Encode(response) } 复制代码`

增加http路由：

` r.Methods( "POST" ).Path( "/calculate/{type}/{a}/{b}" ).Handler(kithttp.NewServer( endpoints.ArithmeticEndpoint, decodeArithmeticRequest, encodeArithmeticResponse, //增加了options append(options, kithttp.ServerBefore(kitjwt.HTTPToContext()))..., )) // ... r.Methods( "POST" ).Path( "/login" ).Handler(kithttp.NewServer( endpoints.AuthEndpoint, decodeLoginRequest, encodeLoginResponse, options..., )) 复制代码`

### Step-4：修改main.go ###

扩展原有ArithmeticEndpoints，增加AuthEndpoint；增加AuthEndpoint的创建逻辑代码，为其增加限流、链路追踪等包装。

` //身份认证Endpoint authEndpoint := MakeAuthEndpoint(svc) authEndpoint = NewTokenBucketLimitterWithBuildIn(ratebucket)(authEndpoint) authEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "login-endpoint" )(authEndpoint) //把算术运算Endpoint\健康检查、登录Endpoint封装至ArithmeticEndpoints endpts := ArithmeticEndpoints{ ArithmeticEndpoint: calEndpoint, HealthCheckEndpoint: healthEndpoint, AuthEndpoint: authEndpoint, } 复制代码`

由于我们要求calculate接口只能在token有效的情况下才可访问，所以为calEndpoint增加token校验代码（最后一行代码，直接使用go-kit提供的中间件）：

` calEndpoint := MakeArithmeticEndpoint(svc) calEndpoint = NewTokenBucketLimitterWithBuildIn(ratebucket)(calEndpoint) calEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "calculate-endpoint" )(calEndpoint) calEndpoint = kitjwt.NewParser(jwtKeyFunc, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(calEndpoint) 复制代码`

### Step-5：运行&测试 ###

通过 ` docker-compose` 启动consul、zipkin、hystrix-dashboard；然后启动gateyway（指定consul地址）；最后启动service（指定consul和service地址）。

在Postman中设置POST请求http://localhost:9090/arithmetic/login，Body为以下内容：

` { "name" : "name" , "pwd" : "pwd" } 复制代码`

可看到以下结果：

![](https://user-gold-cdn.xitu.io/2019/3/11/1696b67c640c22ff?imageView2/0/w/1280/h/960/ignore-error/1)

然后，请求calculate接口，并在Header中设置Authorization，结果如下：

![](https://user-gold-cdn.xitu.io/2019/3/11/1696b67e9d828310?imageView2/0/w/1280/h/960/ignore-error/1)

两分钟后，再次测试，会发现返回token过期。

![](https://user-gold-cdn.xitu.io/2019/3/11/1696b680cc2680b1?imageView2/0/w/1280/h/960/ignore-error/1)

## 总结 ##

本文结合实例，在go-kit微服务中引入jwt。新增login接口，使得calculate接口仅在token有效的情况下才可工作。由于jwt的认证特点，登录成功后用户请求的token有效性不再依赖认证中心，相对OAuth2可大大减轻认证中心的压力，使得微服务的水平扩展变得更加容易。

## 参考文献 ##

* [本文示例代码@github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fraysonxin%2Fgokit-article-demo )
* [JSON Web Token 入门教程（阮一峰）]( https://link.juejin.im?target=http%3A%2F%2Fwww.ruanyifeng.com%2Fblog%2F2018%2F07%2Fjson_web_token-tutorial.html )
* [JWT官方网站]( https://link.juejin.im?target=https%3A%2F%2Fjwt.io%2F )
* [dgrijalva/jwt-go（JWT的一个go实现）]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdgrijalva%2Fjwt-go )

本文首发于本人微信公众号【兮一昂吧】，欢迎扫码关注！

![](https://user-gold-cdn.xitu.io/2019/2/21/16910643343d192e?imageView2/0/w/1280/h/960/ignore-error/1)