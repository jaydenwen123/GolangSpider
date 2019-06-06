# iOS使用Stripe集成支付宝Alipay支付 #

最近在一个海外项目中需要对接支付宝、微信支付，商务对接的同事（产品）要求使用海外支付集成平台Stripe，Stripe的授权支付流程与我们项目中订单支付逻辑并不是十分吻合，还是遇到了不少麻烦，这篇文章主要从前端角度讲解一下Stripe使用 ` Sources` 进行第三方支付的授权流程。

我们以支付宝为例，构建一个支付流程

## 一、创建对应的 ` Source` ##

* 创建一个 ` Source` 需要提供一个Source参数
` let sourceParams = STPSourceParams.alipayParams(withAmount: UInt (aliMoney), currency: "eur" , returnURL: "bentobus://stripe-redirect" ) //以下可以通过metadata为source添加自定义的数据 if sourceParams.metadata == nil { sourceParams.metadata = [ "orderNo" : 3321 ] } else { sourceParams.metadata?[ "orderNo" ] = 12121 } 复制代码` * 创建 ` Source`
` STPAPIClient.shared().createSource(with: source Params) { ( source , error) in if let s = source , s.flow == .redirect { //:TODO } } 复制代码`

## 二、根据创建 ` Source` 后回调的数据，进行下一步操作（支付宝授权支付） ##

创建 ` Source` 以后，根据以上回调属性我们需要进行重定向授权(支付)操作

` let redirectContext = STPRedirectContext( source : s, completion: { ( source Id, clientSecret, error) in // 接收授权（支付）回调 }) redirectContext.startRedirectFlow(from: self)//安装支付宝的情况下会唤起支付宝支付 //redirectContext.startSafariViewControllerRedirectFlowFromViewControlle(self) //redirectContext.startSafariAppRedirectFlow 复制代码`

## 三、接受授权支付后的回调 ##

在第二步重定向授权支付后，后端会收到通知并更改 ` Source` 的状态，接下来我们可以根据 ` Source` 最终的状态来校验支付结果

` STPAPIClient.shared().retrieveSource(withId: source Id, clientSecret: clientSecret, completion: { ( source , error) in switch source !.status { case.chargeable, .consumed : break case.pending : break case.canceled : break case.failed: break default: break } } 复制代码`

以上是支付宝支付的整个流程： 创建 ` Source` ~> 重定向认证支付宝支付 ~> (后端)从支付宝获取支付款项 ~> 获取 ` Source` 最新状态 ~> 支付成功\失败

最后提一下微信支付，截止19年6月初，官方文档中提到Stripe是支持微信支付的，但是客户端原生SDK中并没有提供直接创建微信支付 ` Sources` 的Api，遂弃。 但是，可以看到文档中web端创建 ` Sources` 的方法：

` stripe.createSource({ type : 'wechat' , amount: 1099, currency: 'usd' , }).then( function (result) { // handle result.error or result.source }); 复制代码`

这个方法与我们在客户端SDK中创建Alipay ` Source` 的方法类似，只是客户端并没有提供 ` WeChat` 这样一个 ` type` ：

![屏幕快照2019-06-0616.15.12.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2c085b0b7bb41?imageView2/0/w/1280/h/960/ignore-error/1)

对此，我们可以根据Alipay相关类方法的实现，使用 ` STPSourceParams` 实例方法进行自定义构建：

` STPSourceParams *params = [self new]; params.type = STPSourceTypeUnknown; params.currency = currency; params.redirect = @{ @ "return_url" : return URL }; params.usage = STPSourceUsageReusable; 复制代码`

` STPSourceType` 枚举没有 ` WeChat` 选项，不过根据web端的文档可以使用其中的 ` rawTypeString` 进行设置

以上微信支付只是一个思路，没有经过验证，仅供参考