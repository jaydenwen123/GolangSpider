# 前端打包混编压缩js代码，如何不重新打包，修改js文件内部配置参数？ #

> 
> 
> 
> 利用worker多线程 实现基于vue打包后外置配置化操作 实际就是vue build打包文件都混编了 但是worker多线程 实现外部配置。
> 
> 

> 
> 
> 
> 前端项目在build后，项目的代码通常进行混编、压缩等处理，我们的js代码最终会成为无序的js模块文件。若修改项目中业务的配置参数，通常可以通过接口服务来传达，但是有时候也需要外部的配置文件来传达，如项目已经到生产环境，在不重新打包发版本的基础上，修改其代码内部参数。
> 
> 
> 

### 1.多线程方法 ###

` /** * @description worker.js * @author trsoliu * @date 2019-01-27 * @params url 需要执行的线程 */ const worker = { setWorker : ( url ) => { if ( typeof (Worker) !== "undefined" ) { return new Worker(url); } } } export default worker; 复制代码`

#### 2.配置外部配置文件 ####

` /** * @description config.js 特别说明一下，config.js需要放在根目录static文件夹下，如下图 * @author trsoliu * @date 2019-01-27 */ postMessage({ params1 : 1111 , params2 : 2222 }) 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21bf4f4c668e1?imageView2/0/w/1280/h/960/ignore-error/1)

#### 3.内部js利用worker调用调用配置文件 ####

` //import worker from "./../../../../assets/js/libs/worker.js" //先引入调用方法 worker.setWorker( "./static/js/config/config.js" ).onmessage = ( event ) => { let paramsData = event.data; console.log(paramsData); //console.log结果为：{params1:1111, params2:2222} }; 复制代码`

#### 4.部署后，生产环境config.js文件位置 ####

` 控制台=>Sources=>Page=>config.js，如下图 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21bfa77f21d42?imageView2/0/w/1280/h/960/ignore-error/1)

#### 5.使用 ####

> 
> 
> 
> 按照上述使用后打包，在打包文件中会有一个没有被import合并压缩的config.js文件，此文件为后续版本配置参数修改的文件。单独修改更新服务器上此文件，就可以在无需重新打包的情况下修改全局配置参数。
> 
> 
> 

> 
> 
> 
> 有建议或问题可以加群qq交流 ` 535798405`
> 
>