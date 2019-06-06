# MockJS快速入门 #

## 什么是MockJS ##

在前后端分离的开发环境中，前端同学需要等待后端同学给出接口及接口文档之后，才能继续开发。而MockJS可以让前端同学独立于后端同学进行开发，前端同学可以根据业务先梳理出接口文档并使用MockJS模拟后端接口。那么MockJS是如何模拟后端接口的呢？MockJS通过拦截特定的AJAX请求，并生成给定的数据类型的随机数，以此来模拟后端同学提供的接口。

## 准备工作 ##

写在最前面：有的小伙伴可能不太会部署前端环境，这里我将代码上传到 [github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flizijie123%2F2019Study%2Ftree%2Fmaster%2Fnote%2FmockJS ) 中，有需要的小伙伴可以把代码下载下来，对照着代码看下面的教程。

首先安装MockJS，安装axios是为了发送AJAX请求测试模拟的接口，使用其他方式如原生AJAX请求或$.ajax都是可以的，使用其他方式发送AJAX请求无需安装axios。

` npm install mockjs --save npm install axios --save 复制代码`

使用webpack打包工具打包vue单文件。 首先安装css-loader、style-loader、vue、vue-loader、vue-template-compiler。

` npm install css-loader style-loader --save-dev npm install vue vue-loader vue-template-compiler --save-dev 复制代码`

然后在webpack.config.js配置文件中添加对应的loader,完整的配置图如下。

` const path = require( "path" ); const {VueLoaderPlugin} = require( 'vue-loader' ); module.exports = { entry: './webapp/App.js' , output: { filename: 'App.js' , path: path.resolve(__dirname, './dist' ) }, module: { rules: [ { test : /\.css/, use: [ 'style-loader' , 'css-loader' ] }, { test : /\.vue$/, use: 'vue-loader' } ] }, plugins: [ new VueLoaderPlugin() ], mode: "production" } 复制代码`

创建一个mock.js文件，接着在入口文件中(main.js)引入。

` require( './mock.js' ); 复制代码`

在mock.js文件中引入MockJS。

` import Mock from 'mockjs' ; 复制代码`

后面我们将在mock.js中编写模拟接口。

目录结构如下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b257fa58cd9c3d?imageView2/0/w/1280/h/960/ignore-error/1) index.html是主页文件，main.js是入口文件，mock.js是模拟接口文件，webpack.config.js是webpack配置文件，./dist/main.js是打包后的入口文件，./src/axios/api.js是封装axios请求的文件。

## 正式开始 ##

MockJS有两种方式定义模拟接口返回的数据，一种是使用数据模板定义，这种方式自由度大，可以自定义各种随机的数据类型，一种是使用MockJS的Random工具类的方法定义，这种方式自由度小，只能随机出MockJS提供的数据类型。赶时间的小伙伴可以直接跳到2.使用Rndom工具类定义模拟接口返回值。

### 1.数据模板定义模拟接口返回值 ###

先举一个例子

` //mock.js数据模板 { 'string|1-10' : 'string' } //接口返回的生成的数据===> { 'string' : 'stringstringstring' } 复制代码`

数据模板的格式为 '属性名|生成规则':'属性值' ,生成规则决定了生成的数据的属性值。 生成规则一共有7种，分别是：

* 'name|min-max': value
* 'name|count': value
* 'name|min-max.dmin-dmax': value
* 'name|min-max.dcount': value
* 'name|count.dmin-dmax': value
* 'name|count.dcount': value
* 'name|+step': value

对于不同的数据类型，可以使用的生成规则是不同的，属性值的数据类型可以是Number、Boolean、String、Object、Array、Function、Null，不可以是Undefined，下面我将对每一种数据类型分别使用7种生成规则，以此来观察每一种数据类型可以使用哪些生成规则，想要直接看结论的同学直接拉至1.8的表格中。

#### 1.1'name|min-max': value ####

在开始之前大家可以去我的 [github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flizijie123%2F2019Study%2Ftree%2Fmaster%2Fnote%2FmockJS ) 把项目下载运行起来，结合项目更加直观，通过刷新了解每一个生成规则具体随机出了什么样的数据。

在mock.js文件中，定义一个数组data1用于存放随机生成的属性值，然后定义一个对象row1存放生成的属性值，然后定义模拟接口，在App.vue中使用axios发起请求向模拟接口请求数据并显示在表格中。

` const data1=[]; //数据模板 'name|min-max' :value let row1={ 'string|1-9' : 'string' , 'number|1-9' :1, 'boolean|1-9' : false , 'undefined|1-9' :undefined, 'null|1-9' :null, 'object|1-9' :{object1: 'object1' ,object2: 'object2' ,object3: 'object3' }, 'array|1-9' :[ 'array1' , 'array2' ], 'function|1-9' :()=> 'function' }; data1.push(row1); //定义模拟接口只能接收post请求，定义返回的数据为data1 Mock.mock( '/Get/list1' , 'post' ,data1); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b259f1a1371a62?imageView2/0/w/1280/h/960/ignore-error/1) |min-max规则对于字符串是重复min-max次得出新的字符串，图中可以看出重复了4次。

对于数值是随机生成1-9的数值。

对于布尔类型是min/(min+max)概率生成value值，max/(min+max)概率生成!value值，我这里是10%概率生成false,90%概率生成true。

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象，先在min-max中随机生成一个数值value，然后选取该对象的value个属性出来组成一个新的对象，若value大于该对象的属性个数，则将所有属性拿出来，图中可以看到object对象的所有哦属性被拿出来组成了一个新的对象。

对于数组，先在min-max中随机生成一个数值value，然后将数组元素重复value次然后合并为一个数组，图中可以看出随机出来的数值为6，合并了6个数组得出一个新的数组。

对于函数则直接执行函数并返回了函数的值。

#### 1.2'name|count':value ####

` //数据模板 'name|count' :value let row2={ 'string|3' : 'string' , 'number|3' :1, 'boolean|3' : false , 'undefined|3' :undefined, 'null|3' :null, 'object|3' :{object1: 'object1' ,object2: 'object2' ,object3: 'object3' }, 'array|3' :[ 'array1' , 'array2' ], 'function|3' :()=> 'function' }; data1.push(row2); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25ac90321b313?imageView2/0/w/1280/h/960/ignore-error/1) 第二行是通过'name|count':value规则生成的

|count规则对于字符串是重复count次得出新的字符串，图中可以看出重复了3次。

对于数值是生成一个值为count的数据。

对于布尔类型是(count-1)/count概率生成value值，1/count概率生成!value值，我这里是66.7%概率生成false,33.3%概率生成true。

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象，选取该对象的count个属性出来组成一个新的对象，若count大于该对象的属性个数，则将所有属性拿出来，图中可以看到object对象的所有哦属性被拿出来组成了一个新的对象。

对于数组，将数组元素重复count次然后合并为一个数组，图中可以看出随机出来的数值为3，合并了3个数组得出一个新的数组。

对于函数则直接执行函数并返回了函数的值。

#### 1.3'name|min-max.dmin-dmax':value ####

` //数据模板 'name|min-max.dmin-dmax' :value let row3={ 'string|1-9.1-9' : 'string' , 'number|1-9.1-9' :1, //只有数值起作用 'boolean|1-9.1-9' : false , 'undefined|1-9.1-9' :undefined, 'null|1-9.1-9' :null, 'object|1-9.1-9' :{object1: 'object1' ,object2: 'object2' ,object3: 'object3' }, 'array|1-9.1-9' :[ 'array1' , 'array2' ], 'function|1-9.1-9' :()=> 'function' }; data1.push(row3); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25df9a44e28c2?imageView2/0/w/1280/h/960/ignore-error/1) 第三行是通过'name|min-max.dmin-dmax':value规则生成的。

|min-max.dmin-dmax规则对于字符串就是|min-max规则。

对于数值，是生成一个浮点数，浮点数的整数部分是min-max，小数的位数是dmin-dmax。图中生成整数位3小数位数为2位的值3.59。

对于布尔类型就是|min-max规则

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象也是|min-max规则

对于数组也是|min-max规则

对于函数则直接执行函数并返回了函数的值。

#### 1.4'name|min-max.dcount':value ####

` //数据模板 'name|min-max.dcount' :value let row4={ 'string|1-9.3' : 'string' , 'number|1-9.3' :1, //只有数值起作用 'boolean|1-9.3' : false , 'undefined|1-9.3' :undefined, 'null|1-9.3' :null, 'object|1-9.3' :{object1: 'object1' ,object2: 'object2' ,object3: 'object3' }, 'array|1-9.3' :[ 'array1' , 'array2' ], 'function|1-9.3' :()=> 'function' }; data1.push(row4); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25e552002e5d1?imageView2/0/w/1280/h/960/ignore-error/1) 第四行是通过'name|min-max.scount':value规则生成的。

|min-max.scount规则对于字符串就是|min-max规则。

对于数值，是生成一个浮点数，浮点数的整数部分是min-max，小数的位数是scount。图中生成整数位6小数位数为3位的值6.725。

对于布尔类型就是|min-max规则

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象也是|min-max规则

对于数组也是|min-max规则

对于函数则直接执行函数并返回了函数的值。

#### 1.5'name|count.dmin-dmax':value ####

` //数据模板 'name|count.dmin-dmax' :value let row5={ 'string|3.1-9' : 'string' , 'number|3.1-9' :1, //只有数值起作用 'boolean|3.1-9' : false , 'undefined|3.1-9' :undefined, 'null|3.1-9' :null, 'object|3.1-9' :{object1: 'object1' ,object2: 'object2' ,object3: 'object3' }, 'array|3.1-9' :[ 'array1' , 'array2' ], 'function|3.1-9' :()=> 'function' }; data1.push(row5); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25e8768fcdfa6?imageView2/0/w/1280/h/960/ignore-error/1) 第五行是通过'name|count.dmin-dmax':value规则生成的。

|count.dmin-dmax规则对于字符串就是|count规则。

对于数值，是生成一个浮点数，浮点数的整数部分的值是count，小数的位数是dmin-dmax位。图中生成整数位3小数位数为6位的值3.035354。

对于布尔类型就是|count规则

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象也是|count规则

对于数组也是|count规则

对于函数则直接执行函数并返回了函数的值。

#### 1.6'name|count.dcount':value ####

` //数据模板 'name|count.dcount' :value let row6={ 'string|3.3' : 'string' , 'number|3.3' :1, //只有数值起作用 'boolean|3.3' : false , 'undefined|3.3' :undefined, 'null|3.3' :null, 'object|3.3' :{object1: 'object1' ,object2: 'object2' ,object3: 'object3' }, 'array|3.3' :[ 'array1' , 'array2' ], 'function|3.3' :()=> 'function' }; data1.push(row6); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25eb1536a1888?imageView2/0/w/1280/h/960/ignore-error/1) 第六行是通过'name|count.dcount':value规则生成的。

|count.dcount规则对于字符串就是|count规则。

对于数值，是生成一个浮点数，浮点数的整数部分的值是count，小数的位数是dcount位。图中生成整数位3小数位数为3位的值3.425。

对于布尔类型就是|count规则

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象也是|count规则

对于数组也是|count规则

对于函数则直接执行函数并返回了函数的值。

#### 1.7'name|+step':value ####

` //数据模板 'name|+step' :value let row7 = { 'string|+3' : 'string' , 'number|+3' : 1, //只有数值起作用 'boolean|+3' : false , 'undefined|+3' : undefined, 'null|+3' : null, 'object|+3' : { object1: 'object1' , object2: 'object2' , object3: 'object3' }, 'array|+3' : [ 'array1' , 'array2' ], 'function|+3' : () => 'function' }; data1.push(row7); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25fa342f5f243?imageView2/0/w/1280/h/960/ignore-error/1) 第七行是通过'name|+step':value规则生成的。

|+step规则对于字符串就是无作用，直接将字符串返回。

对于数值，初始值为预设的value值1，每重新请求一次时数值会增加一个step值，点击下方按钮后，数值增加3变为4。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25fcd5158ff5f?imageView2/0/w/1280/h/960/ignore-error/1)

对于布尔类型无作用，直接将布尔值返回。

undefined值直接被忽略了，生成的对象中不存在undefined属性名。

对于null生成的值是null。

对于对象也是无作用，直接将对象返回。

对于数组，第一次请求时返回数组下标为0的值，每重新请求一次时，让下标增加一个step，如点击一次按钮后，下标值为0+3=3超出了数组最大下标1，此时将下标值减去数组长度2直到下标值位于数组最大下标之内，3-2=1，故点击一次按钮后返回数组下标为1的值。

对于函数则直接执行函数并返回了函数的值。

#### 1.8总结 ####

+--------------------+-----------------------------------------+----------------------------------------------------------------------+-------------------------------------------------------------+------------------+------------+--------------------------------------------------------------------------------------------------------------------------------------+---------------------------------------------------------------------------------+------------------------------+
|                    |                 STRING                  |                                NUMBER                                |                           BOOLEAN                           |    UNDEFINED     |    NULL    |                                                                OBJECT                                                                |                                      ARRAY                                      |           FUNCTION           |
+--------------------+-----------------------------------------+----------------------------------------------------------------------+-------------------------------------------------------------+------------------+------------+--------------------------------------------------------------------------------------------------------------------------------------+---------------------------------------------------------------------------------+------------------------------+
| |min-max           | 字符串重复min-max次后拼接得出新的字符串 | 随机得到min-max的值                                                  | min/(min+max)概率生成value值，max/(min+max)概率生成!value值 | 当前数据类型无效 | 返回null值 | 先在min-max中随机生成一个数值value，然后选取该对象的value个属性出来组成一个新的对象，若value大于该对象的属性个数，则将所有属性拿出来 | 先在min-max中随机生成一个数值value，然后将数组元素重复value次然后合并为一个数组 | 直接执行函数并返回了函数的值 |
| |count             | 字符串重复count次得出新的字符串         | 生成一个值为count的数值                                              | (count-1)/count概率生成value值，1/count概率生成!value值     | 当前数据类型无效 | 返回null值 | 选取该对象的count个属性出来组成一个新的对象，若count大于该对象的属性个数，则将所有属性拿出来                                         | 将数组元素重复count次然后合并为一个数组                                         | 直接执行函数并返回了函数的值 |
| |min-max.dmin-dmax | 与规则|min-max相同                      | 生成一个浮点数，浮点数的整数部分是min-max，小数的位数是dmin-dmax     | 与规则|min-max相同                                          | 当前数据类型无效 | 返回null值 | 与规则|min-max相同                                                                                                                   | 与规则|min-max相同                                                              | 直接执行函数并返回了函数的值 |
| |min-max.dcount    | 与规则|min-max相同                      | 生成一个浮点数，浮点数的整数部分是min-max，小数的位数是dcount        | 与规则|min-max相同                                          | 当前数据类型无效 | 返回null值 | 与规则|min-max相同                                                                                                                   | 与规则|min-max相同                                                              | 直接执行函数并返回了函数的值 |
| |count.dmin-dmax   | 与|count规则相同                        | 生成一个浮点数，浮点数的整数部分的值是count，小数的位数是dmin-dmax位 | 与|count规则相同                                            | 当前数据类型无效 | 返回null值 | 与|count规则相同                                                                                                                     | 与|count规则相同                                                                | 直接执行函数并返回了函数的值 |
| |count.dcount      | 与|count规则相同                        | 生成一个浮点数，浮点数的整数部分的值是count，小数的位数是dcount位    | 与|count规则相同                                            | 当前数据类型无效 | 返回null值 | 与|count规则相同                                                                                                                     | 与|count规则相同                                                                | 直接执行函数并返回了函数的值 |
| |+step             | 无作用，将value直接返回                 | 初始值为预设的value值，每重新请求一次时数值value会增加一个step值     | 无作用，将value值返回                                       | 当前数据类型无效 | 返回null值 | 无作用，将value值返回                                                                                                                | 初始值为下标是预设的value的值，每重新请求一次时，下标value会增加一个step值      | 直接执行函数并返回了函数的值 |
+--------------------+-----------------------------------------+----------------------------------------------------------------------+-------------------------------------------------------------+------------------+------------+--------------------------------------------------------------------------------------------------------------------------------------+---------------------------------------------------------------------------------+------------------------------+

其中|min-max.dmin-dmax、|min-max.dcount、|count.dmin-dmax、|count.dcount这四个规则主要是给Number类型使用的。

### 2.使用Rndom工具类定义模拟接口返回值 ###

MockJS提供了一组方法让我们快速随机出想要的数据。

` { 'Boolean' : Random.boolean, // 随机生成布尔类型 'Natural' : Random.natural(1, 100), // 随机生成1到100之间自然数 'Integer' : Random.integer(1, 100), // 生成1到100之间的整数 'Float' : Random.float(0, 100, 0, 5), // 生成0到100之间的浮点数,小数点后尾数为0到5位 'Character' : Random.character(), // 生成随机字符串,可加参数定义规则 'String' : Random.string(2, 10), // 生成2到10个字符之间的字符串 'Range' : Random.range(0, 10, 2), // 生成一个数组，数组元素从0开始到10结束，间隔为2 'Date' : Random.date(), // 生成一个随机日期,可加参数定义日期格式，默认yyyy-mm-dd 'Image' : Random.image(Random.size, '#02adea' , 'Hello' ), // Random.size表示将从size数据中任选一个数据，生成Random.size指定大小的，背景为 '#02adea' 的，内容为 'Hello' 的图片 'Color' : Random.color(), // 生成一个颜色随机值 // 'Paragraph' :Random.paragraph(2, 5), //生成2至5个句子的文本 'Name' : Random.name(), // 生成姓名 'Url' : Random.url(), // 生成url地址 'Address' : Random.province() // 生成地址 } 复制代码`

组织好数据之后同样的需要设置模拟接口。

` Mock.mock( '/Get/list2' , 'post' , data2) ; 复制代码`

再通过AJAX访问该接口即可获得模拟数据。

完整代码：

` //mock.js let data2 = [] // 用于接受生成数据的数组 let size = [ '300x250' , '250x250' , '240x400' , '336x280' , '180x150' , '720x300' , '468x60' , '234x60' , '88x31' , '120x90' , '120x60' , '120x240' , '125x125' , '728x90' , '160x600' , '120x600' , '300x600' ] // 定义随机值 for ( let i = 0; i < 10; i ++) { //生成10个对象放到数组中 let template = { 'Boolean' : Random.boolean, // 可以生成基本数据类型 'Natural' : Random.natural(1, 100), // 生成1到100之间自然数 'Integer' : Random.integer(1, 100), // 生成1到100之间的整数 'Float' : Random.float(0, 100, 0, 5), // 生成0到100之间的浮点数,小数点后尾数为0到5位 'Character' : Random.character(), // 生成随机字符串,可加参数定义规则 'String' : Random.string(2, 10), // 生成2到10个字符之间的字符串 'Range' : Random.range(0, 10, 6), // 生成一个随机数组 'Date' : Random.date(), // 生成一个随机日期,可加参数定义日期格式 'Image' : Random.image(Random.size, '#02adea' , 'Hello' ), // Random.size表示将从size数据中任选一个数据 'Color' : Random.color(), // 生成一个颜色随机值 'Paragraph' :Random.paragraph(2, 5), //生成2至5个句子的文本 'Name' : Random.name(), // 生成姓名 'Url' : Random.url(), // 生成web地址 'Address' : Random.province() // 生成地址 } data2.push(template) } Mock.mock( '/Get/list2' , 'post' , data2) // 声明模拟接口 //App.vue axios( '/Get/list2' ).then(res => { this.dataShow2 = res; }); 复制代码`

## 小结 ##

MockJS使前后分离程度更高，同时，我认为最重要的是他使前端人员也开始思考业务。传统开发中，前端人员多是等待后端人员提供的接口及接口文档，不懂得主动梳理接口文档，使用MockJS后前端人员可以从项目整体的角度出发，能更好的参与到项目之中。

## 交流 ##

如果这篇文章帮到你了，觉得不错的话来点个Star吧。 [github.com/lizijie123]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flizijie123%2F2019Study )