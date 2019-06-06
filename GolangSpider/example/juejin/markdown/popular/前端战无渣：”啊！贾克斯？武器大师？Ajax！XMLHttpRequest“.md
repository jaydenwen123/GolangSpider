# 前端战无渣：”啊！贾克斯？武器大师？Ajax！XMLHttpRequest“ #

> 
> 
> 
> 啥都不说了，挺华为就完事了，山姆大叔有点过分了，我相信华为能挺过来
> 
> 

## 现在该我了！ ##

啊，贾克斯？？？现在从事前端的小伙伴不可能不知道这个，如果真不知道这个词，那我觉得你还称不上前端开发🙄

此贾克斯非彼贾克斯，前端说的啊贾克斯是Asynchronous javaScript + XML的简写，Ajax在很大程度上让前端发展加快了脚步，他的出现和使用，可是说是前端史上的里程碑

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0bd66c610234d?imageView2/0/w/1280/h/960/ignore-error/1)

Ajax横空出世，打破了前后端交互的时候需要重新加载页面，就是整页刷新，我们可以通过Ajax技术直接在页面不刷新的情况下，发出请求，获取返回数据，然后通过js操作Dom更改页面内容。

也是Ajax出现了以后，才促成了往后的 **前后端分离** ，不再是前端只负责切图，只是给后端同学提供模板，然后数据由后端同学添加上。我们前端从此也可以不再依赖后端同学模板套数据了，我们自己就可以完成。

那我们是如何实现这个无刷新获取数据呢，这就要基于一个对象了———— **XMLHttpRequest**

## 开打！开打！ ##

Ajax说白了不是新技术，只是一种解决方案，跟jsonp差不多，是基于 **XMLHttpRequest** 对象的一个获取后端数据的方案，那我们来看看怎么写的吧(本文不涉及兼容ie等部分浏览器，走主流)

### get请求 ###

` ( () => { // 获取页面上需要发请求的按钮 let btn = document.querySelector( '.getAjaxBtn' ); // 添加事件 btn.addEventListener( 'click' , createXHR) // 我们ajax的主体函数 function createXHR ( ) { // new一个XMLHttpRequest实例 var xhr = new XMLHttpRequest(); // 第一步使用open方法，接收三个参数 // 第一个参数是请求方法名（get，post等） // 第二个参数是需要请求的接口地址 // 第三个参数是设置请求是否是异步，一般都是都是发送异步请求，同步请求可能会阻塞页面 // 我们先来看同步请求 xhr.open( 'get' , '/api/get' , false ); // open方法只是设置参数，并不会发送请求 // 而请求是由send()方法发送的，并且接收一个参数，就是需要发送到服务端的数据 // 如果没有需要发送到服务端的数据，必须传入null，因为有些浏览器不许要这个参数 xhr.send( null ) } })() 复制代码`

这样写的话，我们就在一个按钮上绑定了一个发送get请求的方法，只要我们点击这个按钮，就会发送请求，那我们来看看是什么效果

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0c11b76815ec2?imageslim)

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0c13e63a3123d?imageView2/0/w/1280/h/960/ignore-error/1)

其实很简单是不是，没错，这样我们确实就已经发送了一个get请求了，那我们发送请求是为了什么呢，当然是为了获取到数据，那我们看到接口已经返回了字符串 ` '恭喜你，你发送了一个get请求，真棒！'` ，那我们怎么能拿到这个数据呢，看下面

### 拿到返回数据 ###

其实我们在收到服务端响应的时候，就已经拿到了数据，响应的数据会自动填充到XHR对象中，但我们什么时候才能知道这个数据已经填充到XHR对象中呢？？这个时候我们就需要一个类似于监听函数的事件———— **readystatechange事件** ，z这个事件当readystate状态改变，就会执行一次。我们还需要一个知道请求发送状态的属性———— **readystate** ，还有一个请求状态码属性———— **status**

状态码就是200、404以及500这种的，在这就不细说了，主要说说readystate有什么状态：

* 0: 未初始化。尚未调用open()方法
* 1: 启动。已经调用open()方法，但尚未调用send()方法
* 2：发送。已经调用send()方法，但尚未接收到响应
* 3：接收。已经接收到 **部分** 相应数据
* 4：完成。已经接收到 **全部** 响应数据，而且已经可以在客户端使用了

看了一下这接个状态，好像我们现在需要的是4这个状态，接收到了全部的数据，然后我们再做一个逻辑处理，上面的请求我们是同步请求，现在我们改改装一下

` ( () => { // 获取页面上需要发请求的按钮 let btn = document.querySelector( '.getAjaxBtn' ); // 添加事件 btn.addEventListener( 'click' , createXHR) // 我们ajax的主体函数 function createXHR ( ) { // new一个XMLHttpRequest实例 var xhr = new XMLHttpRequest(); // 第一步使用open方法，接收三个参数 // 第一个参数是请求方法名（get，post等） // 第二个参数是需要请求的接口地址 // 第三个参数是设置请求是否是异步，一般都是都是发送异步请求，同步请求可能会阻塞页面 // 我们先来看同步请求 xhr.open( 'get' , '/api/get' , true ); // 我们为xhr添加一个readyState值改变就执行的监听事件 // 这个事件还必须写在send()方法前面 xhr.onreadystatechange = function ( ) { // 当readyState值为4就是接收到了全部返回数据，并且http状态码是200多的成功，或者是304的缓存，这时候就判断已经是成功拿到了应该返回的数据 if (xhr.readyState === 4 ) { if (xhr.status >= 200 && xhr.status < 300 || xhr.status === 304 ) { // 输出数据。responseText是作为响应主体被返回的文本 alert(xhr.responseText) } else { alert( '请求不成功' ) } } } // open方法只是设置参数，并不会发送请求 // 而请求是由send()方法发送的，并且接收一个参数，就是需要发送到服务端的数据 // 如果没有需要发送到服务端的数据，必须传入null，因为有些浏览器不许要这个参数 xhr.send( null ) } })() 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0cad506791063?imageView2/0/w/1280/h/960/ignore-error/1)

这样我们就成功的拿到了返回的数据并且弹出来了，我们不仅拿到数据可以弹出，我们就可以进行一些更厉害的逻辑处理或者改变页面内容了。

当然了，也许你get请求的时候需要发送一些数据给服务端，可能是什么页码或者页面需要多少条数据等，我们这个时候只要在open()方法中传入的url上面拼上就行了，类似 ` /api/get?pageNum=1&pageSize=20` 这样的url。

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0cb73ce920cd0?imageView2/0/w/1280/h/960/ignore-error/1)

我们写的服务端就是返回请求中的数据，我们也能看到，我们也收到了返回来的数据了

### post请求 ###

post请求，我们应该接触过form的表单提交，我们post请求传输的数据必须是经过 ` encodeURIComponent` 转码的，而且当我们接收到一个json对象，我们需要将他进行转换

` ( () => { let btn = document.querySelector( '.getAjaxBtn' ); btn.addEventListener( 'click' , createXHR); // 我们需要转换的数据，真实场景我们应该是传进来的，或者是获取到的用户输入的数据 let data = { method : 'post' , a : 'a' , b : 123 , } function createXHR ( ) { var xhr = new XMLHttpRequest(); xhr.open( 'post' , '/api/post' , true ); xhr.onreadystatechange = function ( ) { if (xhr.readyState === 4 ) { if (xhr.status >= 200 && xhr.status < 300 || xhr.status === 304 ) { alert(xhr.responseText) } else { alert( '请求不成功' ) } } } // 模仿表单提交，设置content-type，我们需要重新设置请求头，让服务端能知道我们传的是个什么数据 xhr.setRequestHeader( 'Content-type' , 'application/x-www-form-urlencoded' ); // send()方法需要传入的数据应该经过处理 xhr.send(transformData(data)) } // 处理传入json对象转换格式 /** * 例如这个样的json * { method: 'post', a: 'a', b: 123, } */ // 转换格式为 // method=post&a=a&b=123 // 这样 function transformData ( data ) { let newData = []; for ( let key in data) { newData.push( encodeURIComponent (key) + '=' + encodeURIComponent (data[key])) } return newData.join( '&' ) } })() 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0cda14aaf9946?imageView2/0/w/1280/h/960/ignore-error/1)

上面就是我们进行的post请求，我们能看到我们发送了一了个Form Data，然后服务端拿到了这个数据，并且我们拿到了会返回来的数据，也就是我们发送给服务端的数据

### 还有谁！ ###

当然除了我上面说到的一些属性和方法，XHR还有很多我没有介绍到的方法和属性，这里就说一下几个常用的吧

* **XMLHttpRequest.timeout：** 这个就是可以设置请求超时，设置一个时间，单位为毫秒，如果请求时间超过设置时间，直接抛出异常
* **XMLHttpRequest.withCredentials：** 这个属性是看我们需不需要发送cookie的，因为如果是跨域请求，一般请求是不会带上cookie，需要我们手动设置这个属性
* 
* **XMLHttpRequest.abort()：** 这个方法我们用来终止请求，当然是在请求发送并没有返回的时候
* **XMLHttpRequest.setRequestHeader()：** 这个我们刚才有用到过，就是设置请求头，更改一些我们需要的设置

## 把他们也算上！ ##

如果我们这么写，我们会发现很麻烦，而且如果我们没有考虑周全，这个请求方法会很糟糕，隐藏一些未知问题

技术，向来都是为了提高效率，或者也可以说是为懒人研究的

现在已经有很好的成熟的开源库替我们封装好了ajax请求方法，例如“尚能饭否的jQuery”以及随着框架模块大放异彩的“Axios”，他们的使用方法相对简单很多，也不用我们自己去处理一些逻辑问题，边界问题。

### jQuery & Axios ###

我们来看看这两个是多么简单实现的

` <!DOCTYPE html> < html lang = "en" > < head > < meta charset = "UTF-8" > < meta name = "viewport" content = "width=device-width, initial-scale=1.0" > < meta http-equiv = "X-UA-Compatible" content = "ie=edge" > < script src = "https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js" > </ script > < script src = "https://unpkg.com/axios/dist/axios.min.js" > </ script > < title > Ajax </ title > </ head > < body > < button class = "postAxiosBtn" > 发送axios post请求 </ button > < button class = "getAxiosBtn" > 发送axios get请求 </ button > < button class = "postJqueryBtn" > 发送jQuery post请求 </ button > < button class = "getJqueryBtn" > 发送jQuery get </ button > < script > var postAxiosBtn = $( '.postAxiosBtn' ), getAxiosBtn = $( '.getAxiosBtn' ), postJqueryBtn = $( '.postJqueryBtn' ), getJqueryBtn = $( '.getJqueryBtn' ); postAxiosBtn.click( () => { axios.post( '/api/post' , { data : { method : 'post' , a : 'a' , b : 123 , } }).then( res => { console.log(res) }) }) getAxiosBtn.click( () => { axios.get( '/api/get' , { params : { pageNum : 1 , pageSize : 20 } }).then( res => { console.log(res) }) }) postJqueryBtn.click( () => { $.post( '/api/post' , { method : 'post' , a : 'a' , b : 123 }, function ( data, status ) { console.log(data) console.log(status) }) }) getJqueryBtn.click( () => { $.get( '/api/get' , { pageNum : 1 , pageSize : 20 }, function ( data, status ) { console.log(data) console.log(status) }) }) </ script > </ body > </ html > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ce2aa824e957?imageView2/0/w/1280/h/960/ignore-error/1)

封装库使用起来就是这么简单！！！

## 哼，一个能打的都没有！ ##

### fetch ###

想来大家也都知道fetch这个api了，js新出的一个api，对标的就是XMLHttpRequest，但是相比于XMLHttpRequest，fetch的使用很人性化了，而且自带promise语法，我们不用再用什么回调的方法了，但是fetch再去跟axios这种封装XMLHttpRequest的库来说，使用方法上还有有很大的差距吧。

但我相信以后封装fetch这种的库会越来越多的，现在虽然有，但是我感觉用的人还没有要使用axios那种程度😜

## 总结 ##

现在已经很少会用到纯手写原生Ajax了，都会用已经写好的封装库了，但是面试的时候不乏也会遇到让你手写的，这就很头疼了

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ce80ac54878d?imageView2/0/w/1280/h/960/ignore-error/1)

尤大的回答不无道理，但是能说上来或者手写原生Ajax的人，绝对是对只是有着一种执着

知其然也得知其所以然，写伪代码我感觉也是说的过去的，起码对大的思路要了解，一些api接不住无所谓，js、css那么多属性，方法，要是都记下来肯定是不可能的，但是还是能知道有这个东西的存在，查的时候也好查，是不是🤠

我是前端战五渣，一个前端界的小学生。