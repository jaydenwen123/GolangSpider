# AI时代验证码的攻与防 #

> 
> 
> 
> 导读：科学技术是第一生产力，技术的演进会很大程度上推动业务的升级，同时业务的演进进一步促进技术往更高的层面发展，智能验证码的技术发展同样类似。
> 
> 

首先，非常高兴和大家来分享下我在验证码攻防方面的一些实践和探索。我主要从事 [反欺诈服务]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2F ) 的相关工作，最近一段时间接触到OpenCV、Tensorflow、Tesseract、ThreeJS等相关技术，在一次分享会上听到Puppeteer的介绍，经过半年研究将这些技术整合起来，在Node中对市面主流验证码做了一轮攻防演练。

## 1、CAPTCHA ##

全自动区分计算机和人类的公开图灵测试（英语：Completely Automated Public Turing test to tell Computers and Humans Apart，简称CAPTCHA），俗称 [验证码]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2Fproduct%2Fcaptcha ) ，是一种区分用户是计算机或人的公共全自动程序。在CAPTCHA测试中，作为服务器的计算机会自动生成一个问题由用户来解答。这个问题可以由计算机生成并评判，但是必须只有人类才能解答。由于计算机无法解答CAPTCHA的问题，所以回答出问题的用户就可以被认为是人类。

## 2、简介 ##

CAPTCHA这个词最早是在2002年由卡内基梅隆大学的路易斯·冯·安、Manuel Blum、Nicholas J.Hopper以及IBM的John Langford所提出。卡内基梅隆大学曾试图申请此词使其成为注册商标， 但该申请于2008年4月21日被拒绝。一种常用的CAPTCHA测试是让用户输入一个扭曲变形的图片上所显示的文字或数字，扭曲变形是为了避免被光学字符识别（OCR, Optical Character Recognition）之类的计算机程序自动识别出图片上的文数字而失去效果。由于这个测试是由计算机来考人类，而不是标准图灵测试中那样由人类来考计算机，人们有时称CAPTCHA是一种反向图灵测试。

为了无法看到图像的身心障碍者，替代的方法是改用语音读出文数字，为了防止语音识别分析声音，声音的内容会有杂音或仍可以被人类接受的变声。

根据CAPTCHA测试的定义，产生验证码图片的算法必须公开，即使该算法可能有专利保护。这样做是证明想破解就需要解决一个不同的人工智能难题，而非仅靠发现原来的（秘密）算法，而后者可以用逆向工程等途径得到。

## 3、作用 ##

防止恶意破解密码、刷票、论坛灌水、刷页。有效防止某个黑客对某一个特定注册用户用特定程序暴力破解方式进行不断的登录尝试，实际上使用验证码是现在很多网站通行的方式（比如招商银行的网上个人银行，百度社区），我们利用比较简易的方式实现了这个功能。虽然登录麻烦一点，但是对网友的密码安全来说这个功能还是很有必要，也很重要。但我们还是 提醒大家要保护好自己的密码 ，尽量使用混杂了数字、字母、符号在内的6位以上密码，不要使用诸如1234之类的简单密码或者与用户名相同、类似的密码 ，免得你的账号被人盗用给自己带来不必要的麻烦。

验证码通常使用一些线条和一些不规则的字符组成，主要作用是为了防止一些黑客把密码数据化盗取。

典型应用场景：

[网站安全]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2Fproduct%2FregisterDefend ) ：垃圾注册、恶意登录、账号盗用

[数据安全]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2Fproduct%2FinterfaceDefend ) ：数据爬取、数据破坏

[运营安全]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2Fproduct%2FmarketingDefend ) ：恶意刷单、虚假秒杀、虚假评论

[交易安全]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2Fproduct%2FtransactionDefend ) ：虚假交易、恶意套现、盗卡支付

## 4、工作原理 ##

同盾智能验证产品主要针对企业不同业务场景，利用生物行为与机器学习方式提供人机验证服务，防范非真实人类流量和恶意程序攻击，帮您降低业务风险。

使用同盾智能验证码，告别烦躁脆弱的验证码，提升用户体验的同时增强安全防御能力。远离垃圾帖、恶意注册、垃圾短信、刷票等交互安全烦恼。

其工作流程如下图所示：

![工作流程](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d859ed71da?imageView2/0/w/1280/h/960/ignore-error/1)

## 5、验证码分类 ##

* 

字符型图片验证码

字符型图片验证码是由阿拉伯数字，英文字母，中文汉字按照一定规律排列，加入干扰噪点之后生产的一张图片。

* 

行为式验证码

[行为式验证码]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2FonlineExperience%2FslidingPuzzle ) 是一种较为流行的验证码。从字面来理解，就是通过用户的操作行为来完成验证，而无需去读懂扭曲的图片文字。常见的有两种：拖动式与点触式。

* 

智能验证码

[智能验证码]( https://link.juejin.im?target=https%3A%2F%2Fx.tongdun.cn%2FonlineExperience%2Fcaptcha ) 是一种基于语言认知的人机区分，考验机器语言认知能力的智能验证码，会是未来一段时间内的重要选择。典型代表有语序点选和空间推理。

### 5.1使用到的相关技术 ###

* OpenCV
* Tensorflow
* Node
* Tesseract
* Python
* node-tesseract & gm & node-cmd 库使用
* ThreeJs & Java 3D

## 6、字符型图片验证码破解 ##

[图像二值化]( https://link.juejin.im?target=https%3A%2F%2Fbaike.baidu.com%2Fitem%2F%25E5%259B%25BE%25E5%2583%258F%25E4%25BA%258C%25E5%2580%25BC%25E5%258C%2596 ) （ Image Binarization）就是将图像上的像素点的灰度值设置为0或255，也就是将整个图像呈现出明显的黑白效果的过程。

在数字图像处理中，二值图像占有非常重要的地位，图像的二值化使图像中数据量大为减少，从而能凸显出目标的轮廓。

![图像二值化](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d90d8d2b83?imageView2/0/w/1280/h/960/ignore-error/1)

### 6.1使用gm去噪处理 ###

主要是去掉图像里的所有干扰信息，比如背景的点，线等。

` /** * 对图片进行阈值处理(默认55%) * thresholdValue，降噪值 */ async function disposeImg(imgPath, newPath, thresholdValue) { return new Promise((resolve, reject) => { gm(imgPath).threshold(thresholdValue || '55%' ).write(newPath, (err) => { if (err) return reject(err); resolve(newPath); }); }); } 复制代码`

### 6.2使用tesseract进行OCR识别 ###

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d85cae28f1?imageView2/0/w/1280/h/960/ignore-error/1)

### 6.3Tesseract介绍 ###

Tesseract(/‘tesərækt/) 这个词的意思是”超立方体”，指的是几何学里的四维标准方体，又称”正八胞体”，是一款被广泛使用的开源 OCR 工具。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d85c1552fb?imageslim)

Tesseract 已经有 30 年历史，开始它是惠普实验室于1985年开始研发的一款专利软件，到1995年一件成为OCR业界内最准确的识别引擎之一。然而，HP不久便决定放弃OCR业务Tesseract从此尘封。数年之后，HP意识到与其将Tesseract束之高阁，还不如贡献给开源，让其重焕新生。在 2005 年，Tesseract由美国内华达州信息技术研究所获得，并求助于Google对Tesseract进行改进、消除Bug、优化工作，并开源，其后一直由 Google 赞助进行后续的开发和维护。因为其免费与较好的效果，许多的个人开发者以及一些较小的团队在使用着 Tesseract ，诸如验证码识别、车牌号识别等应用中，不难见到 Tesseract 的身影。

现在Tesseract托管在Github上，大家有兴趣可以上Github上Star或Frok [该项目]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftesseract-ocr%2Ftesseract ) 。

所谓 OCR(Optical Character Recognition)是指对文本资料进行扫描，然后对图像文件进行分析处理，获取文字和版面信息的过程。OCR是图像识别领域中的一个子领域，该领域专注于对图片中的文字信息进行识别并转换成能被常规文本编辑器编辑的文本。

使用Tesseract对降噪后的图片进行OCR识别，Tesseract支持107种语言，可以通过node-tesseract代理调用，也可以直接通过node-cmd执行Tesseract命令进行调用。 node作为中间层，可以直接调用封装好的node包，或者直接执行底层命令，或者调用Python脚本都可以。

通过node-tesseract进行OCR识别

` async function recognizeImg(imgPath, options) { //tesseract --list-langs，语言选项 options = Object.assign({ psm: 7, //-psm 7 表示告诉tesseract code.jpg图片是一行文本,这个参数可以减少识别错误率，默认为3 }, options); //console.log(options); return new Promise((resolve, reject) => { tesseract.process(imgPath, options, (err, text) => { if (err) return reject(err); resolve(text.replace(/[\r\n\s]/gm, '' )); // 去掉识别结果中的换行回车空格 }); }); } 复制代码`

### 6.4tesseract进行OCR识别 ###

` async function execute(imgPath) { return new Promise((resolve, reject) => { NodeCmd.get(`tesseract ${imgPath} stdout -l chi_sim`, function (err, data) { if (!err) { resolve(data); } else { reject(err); } }, ); }); } 复制代码`

### 6.5打码平台 ###

机器自动识别图片验证码，对简单的情况能有较高的准确率，但对干扰多，变形复杂的图片验证码，其准确率会很差。由于图片验证码重要度增加，复杂的图片验证码被大量使用，导致近年来出现了利用众包力量实现的人工验证码识别平台。 其工作原理图下所示：

![打码平台](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d85e8191cc?imageView2/0/w/1280/h/960/ignore-error/1)

字符型图片验证码基本沦陷，验证码进入行为式验证时代。

## 7、行为式验证码破解 ##

行为式验证的核心思想是利用用户的“行为特征”来做验证安全判别。整个验证框架采用高效的“行为沙盒”主动框架, 这个框架会引导用户在“行为沙盒”内产生特定的行为数据，利用“多重复合行为判别”算法从特指、视觉、思考等多重行为信息中辨识出生物个体的特征， 从而准确快速的提供验证结果。

但是随着puppeteer的出现，行为式验证码的防御不在奏效。Puppeteer 的 Logo 很形象，顾名思义像是一个被操控的傀儡、提线木偶。

![Puppeteer](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d8961b51ab?imageView2/0/w/1280/h/960/ignore-error/1)

Puppeteer 是一个 Node 库，它提供了高级的 API 并通过 DevTools 协议来控制 Chrome(或Chromium)。通俗来说就是一个 headless chrome 浏览器 (也可以配置成有 UI 的，默认是没有的)

![Puppeteer2](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d89972469a?imageView2/0/w/1280/h/960/ignore-error/1)

通过puppetter的page.screenshot进行指定区域截屏，如果页面上未引入jQuery通过page.addScriptTag引入，后面截屏处理需要使用到。利用resemblejs/compareImages进行图片比对取得滑动距离的图片，通过canvas将图片读入内存，取得最终滑动距离，调用puppetter加入人的行为模拟，并最终验证通过。

![slider](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d89dca19a5?imageView2/0/w/1280/h/960/ignore-error/1)

[Puppeteer-无头浏览器简介]( https://link.juejin.im?target=https%3A%2F%2Fpptr.dev%2F )

### 7.1浏览器配置 ###

` const browser = await puppeteer.launch({ args: [ '--start-maximized' ],//讓開啟來的瀏覽器預設最大 executablePath: '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome' , //executablePath: '/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary' , //executablePath: '/Applications/Chromium.app/Contents/MacOS/Chromium' , //devtools: true , headless: false , //这里我设置成 false 主要是为了让大家看到效果，设置为 true 就不会打开浏览器 }); 复制代码`

### 7.2UA相关配置 ###

` // 设置浏览器信息 const UA = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko)' + ' Chrome/69.0.3497.100 Safari/537.36' ; await Promise.all([ page.setUserAgent(UA), // 允许运行js page.setJavaScriptEnabled( true ), // 设置浏览器视窗 page.setViewport({ width: 1280, height: 748, }), ]); 复制代码`

### 7.3Puppeteer过人机检测信息 ###

` await page.evaluateOnNewDocument(() => { Object.defineProperty(navigator, 'webdriver' , { get: () => false , }); Object.defineProperty(navigator, 'plugins' , { get: () => [1, 2, 3, 4, 5], }); const originalQuery = window.navigator.permissions.query; return window.navigator.permissions.query = (parameters) => ( parameters.name === 'notifications' ? Promise.resolve({state: Notification.permission}) : originalQuery(parameters) ); }); 复制代码`

### 7.4注入jQuery脚本 ###

` await page.addScriptTag({url: 'https://cdn.bootcss.com/jquery/3.2.0/jquery.min.js' }); 复制代码`

注入jQuery主要是为了方便Dom操作，在滑动验证的时候进行图片截取处理。

### 7.5截屏处理 ###

` // 截屏 await screenshot(captionPosition, './screenshots/xxx.png' ); await cropImage( './screenshots/xxx.png' , './screenshots/xxx.png' ); 复制代码`

最终通过Puppetter开启一个浏览器，实现自动截屏，自动模拟点击，自动登录功能。

行为式验证码基本沦陷，验证码进入AI时代。

## 8、智能验证码破解 ##

基于语言认知和结构化知识图谱的验证码，可以有效效避免攻击，提高人机识别的准确率。这里引入OpenCV和Tensorflow两个工具，主要用于图片识别和深度学习。

### 8.1Inception V1 ###

GoogLeNet首次出现在2014年ILSVRC 比赛中获得冠军。这次的版本通常称其为Inception V1。Inception V1有22层深，参数量为5M。同一时期的VGGNet性能和Inception V1差不多，但是参数量也是远大于Inception V1。

![V1](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d89e06fa56?imageView2/0/w/1280/h/960/ignore-error/1)

Inception Module是GoogLeNet的核心组成单元，Inception Module基本组成结构有四个成分。1 1卷积，3 3卷积，5 5卷积，3 3最大池化。最后对四个成分运算结果进行通道上组合。这就是Inception Module的核心思想。通过多个卷积核提取图像不同尺度的信息，最后进行融合，可以得到图像更好的表征。结构如下图：

![V5](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d8c730d251?imageView2/0/w/1280/h/960/ignore-error/1)

### 8.2Inception V5图片打标分类&人脸识别 ###

最新的Inception V5训练好的模型大概可与识别1000种类别的图片，通过opencv4nodejs可以调用训练好的模型进行图片分类打标，人脸检测等各种图片识别功能。

Google提供的 [inception5h]( https://link.juejin.im?target=https%3A%2F%2Fstorage.googleapis.com%2Fdownload.tensorflow.org%2Fmodels%2Finception5h.zip ) 模型介绍

### 8.3基于opencv4nodejs完成图片打标和人脸识别 ###

` const cv = require( 'opencv4nodejs' ); const fs = require( 'fs' ); const path = require( 'path' ); const gm = require( 'gm' ); const imageMagick = gm.subClass({imageMagick: true }); const DIRECTION = { NorthWest: 'NorthWest' , North: 'North' , NorthEast: 'NorthEast' , West: 'West' , Center: 'Center' , East: 'East' , SouthWest: 'SouthWest' , South: 'South' , SouthEast: 'SouthEast' , }; if (!cv.xmodules.dnn) { throw new Error( 'exiting: opencv4nodejs compiled without dnn module' ); } // 载入模型 const inceptionModelPath = './models/tf-inception' ; const modelFile = path.resolve(inceptionModelPath, 'tensorflow_inception_graph.pb' ); const classNamesFile = path.resolve(inceptionModelPath, 'imagenet_comp_graph_label_strings_cn.txt' ); /*const inceptionModelPath = './models/flower' ; const modelFile = path.resolve(inceptionModelPath, 'retrained_graph.pb' ); const classNamesFile = path.resolve(inceptionModelPath, 'retrained_labels.txt' );*/ if (!fs.existsSync(modelFile) || !fs.existsSync(classNamesFile)) { console.log( 'exiting: could not find inception model' ); console.log( 'download the model from: https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip' ); return ; } console.log( 'load models:' + inceptionModelPath); // 模型和分类标签 const classNames = fs.readFileSync(classNamesFile).toString().split( '\n' ); //const net = cv.readNetFromTensorflow(modelFile); const net = cv.readNetFromTensorflow(modelFile); const classifyImg = (img) => { // inception model works with 224 x 224 images, so we resize // our input images and pad the image with white pixels to // make the images have the same width and height const maxImgDim = 224; const white = new cv.Vec(255, 255, 255); const imgResized = img.resizeToMax(maxImgDim).padToSquare(white); // network accepts blobs as input const inputBlob = cv.blobFromImage(imgResized); net.setInput(inputBlob); // forward pass input through entire network, will return // classification result as 1xN Mat with confidences of each class const outputBlob = net.forward(); // find all labels with a minimum confidence const minConfidence = 0.05; const locations = outputBlob.threshold(minConfidence, 1, cv.THRESH_BINARY).convertTo(cv.CV_8U).findNonZero(); const result = locations.map(pt => ({ confidence: parseInt(outputBlob.at(0, pt.x) * 100) / 100, className: classNames[pt.x], })) // sort result by confidence .sort((r0, r1) => r1.confidence - r0.confidence).map(res => { return `标签: ${res.className} 、概率: ${res.confidence} 、百分比: ${res.confidence * 100} %`; }); return result; }; const test Data = []; /** * 文件遍历方法 * @param filePath 需要遍历的文件路径 */ function fileDisplay(filePath) { //根据文件路径读取文件，返回文件列表 const tempFileDir = fs.readdirSync(filePath); tempFileDir.forEach( function (filename) { //获取当前文件的绝对路径 const filedir = path.join(filePath, filename); //根据文件路径获取文件信息，返回一个fs.Stats对象 let tempFile = fs.statSync(filedir); if (tempFile.isFile()) { let tempObj = {}; tempObj[ 'image' ] = filedir; tempObj[ 'label' ] = filename; test Data.push(tempObj); } if (tempFile.isDirectory()) { fileDisplay(filedir);//递归，如果是文件夹，就继续遍历该文件夹下面的文件 } }); } fileDisplay(path.resolve( './data' )); test Data.forEach((data) => { const img = cv.imread(data.image); //console.log( '%s,%s: ' , data.image, data.label); //console.log( '%s: ' , data.image); const predictions = classifyImg(img); let tempText = '' ; predictions.forEach(p => { //console.log(JSON.stringify(p)); tempText += p + '\n' ; }); console.log(tempText); imageMagick(data.image).gravity(DIRECTION[ 'NorthWest' ]) //水印的位置 .geometry( '+10+10' ) //距离右下角右边10px下边10px //.fontSize(24) .fill( 'red' ) .pointSize(18) //.font( '/Library/Fonts/Songti.ttc' ) .font( './font/msyh.ttf' )//字体必须正确，否则乱码或者不显示中文 .stroke( 'red' )//文字颜色 .drawText(15, 10, tempText) //15和10是位置信息 最后一个参数是文字信息 .write(data.image.replace( '/data/' , '/process/' ), function (err) { if (err) { return console.error( 'err--------' , err); } console.log( '%s:' , data.image, tempText, '打标处理完成!' ); }); //cv.imshowWait( 'img' , img); console.log( '---------finish---------' ); }); 复制代码`

### 8.4通过迁移训练来定制 TensorFlow 模型 ###

基于Google Inception-V3 模型，在Windows平台通过TensorFlow 利用GTX1080进行并行计算学习，得到自己想要的模型结果。

![模型训练](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d8d00e545f?imageView2/0/w/1280/h/960/ignore-error/1)

python训练脚本

` python retrain.py --bottleneck_dir=C:\Users\***\PycharmProjects\flower_photos\bottlenecks --how_many_training_steps=500 --model_dir=C:\Users\***\PycharmProjects\flower_photos\inception --summaries_dir=C:\Users\***\PycharmProjects\flower_photos\training_summaries\basic --output_graph=C:\Users\***\PycharmProjects\flower_photos\retrained_graph.pb --output_labels=C:\Users\***\PycharmProjects\flower_photos\retrained_labels.txt --image_dir=C:\Users\***\PycharmProjects\flower_photos 复制代码`

使用训练脚本完成图片识别

` const Promise = require( 'bluebird' ); const childProcess = require( 'child_process' ); (async () => new Promise(((resolve, reject) => { /** * child.stdin 获取标准输入 * child.stdout 获取标准输出 * child.stderr 获取标准错误输出 */ childProcess.exec( 'python3 /Users/***/Downloads/python/label_image.py --image /Users/***/Downloads/python/232831_44935.jpg --graph /Users/***/Downloads/python/retrained_graph.pb --labels /Users/***/Downloads/python/retrained_labels.txt' , (error, stdout, stderr) => { // console.log( 'error' , JSON.stringify(error)); // console.log( 'stdout' , stdout); // console.log( 'stderr' , JSON.stringify(stderr)); if (stdout.length > 0) { console.log(stdout); resolve(stdout); } if (error) { console.info(stderr); reject(error); } }); })))(); 复制代码`

## 9、点选、空间推荐验证破解思路 ##

### 9.1语序点选 ###

![语序点选](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d913285544?imageView2/0/w/1280/h/960/ignore-error/1)

### 9.2空间推理验证 ###

![空间推理验证](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d8eaeb9c37?imageView2/0/w/1280/h/960/ignore-error/1)

![AI验证码破解](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d921cfee23?imageView2/0/w/1280/h/960/ignore-error/1)

1、通过Puppetter对目标网站进行数据样本爬取，爬取问题和答案，建立语序点选，图标点选，空间推理的样本库。基于图片学习和深度学习，完成样本库的模型训练之后得到对应的模型数据供后续破解使用。

2、Puppetter访问对应的网站，唤醒拉起验证码，通过Puppetter截图获取问题图片。通过Node服务执行Python脚本完成识图功能，取得对应的答案。基于得到的答案通过Puppetter完成验证，实现目标网站的验证破解。

至此，智能验证码进入AI对抗时代。

## 10、对抗升级 ##

### 10.1 3D验证码 ###

传统验证码大部分基于二位空间，答案就在当前面，这使得Puppetter模拟线性的路径相对容易。从新颖性和趣味性安全性三个方面出发，我们研究了在二维空间模拟三维空间，实现答案的隐藏和解空间难度升级，大大提升了验证码的趣味性和安全性。

![3D文字](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d92f7c8abc?imageslim)

![3D图片](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d94b5cdcae?imageView2/0/w/1280/h/960/ignore-error/1)

3D旋转验证产品特点：

* 问题面由1面提升到6面
* 支持文字，图片等多种形式
* 答案面默认不在当前用户可见面
* 验证码默认有一个旋转速率，增加打码和机器人截取的难度
* 答案在面上随机分布
* 答案面转动的时候，任意角度皆可以通过，只要答案正常

![3D旋转](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d954d6f5df?imageView2/0/w/1280/h/960/ignore-error/1)

3D验证码主要技术点：

* Java服务将face面编码和face面信息随机传递到网页端
* 网页端基于CSS 3D进行3D立方体绘制，结合编面信息将立方体的6个面绘制完成
* 获取空间的相对点位信息用户答案的验证，用户面的转动角度不影响答案的计算

### 10.2空间推理验证 ###

空间推理验证的安全性高于传统的滑动和字符型验证码，为了丰富产品线。我们基于ThreeJS和Java 3D协同完成前后端的验证识别。

![空间推理验证](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d8eaeb9c37?imageView2/0/w/1280/h/960/ignore-error/1)

空间推理验证产品特点：

* Java 3D提供问题和答案建模的数据模型
* 基于ThreeJS 完成问题和答案的生成，生成服务采用离线策略，确保产品的兼容性和速度
* 客户端的验证码基于三维转平面的一张图片，基于视觉视差实现伪3D
* 建模数据里面使用了一些方便人理解，但是机器难以理解的描述语言，增加解空间难度
* 所有校验算法都在后端，前端只负责最基础的展现，安全有保障

![空间推理验证](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5d96ad95ad2?imageView2/0/w/1280/h/960/ignore-error/1)

空间推理验证技术点：

* Java 3D负责ThreeJS 3D建模数据的产生，验证码的问题和答案数据由后端算法动态生成
* 基于ThreeJS 完成3D建模，通过灯光，相机，场景的组合得到三维立体空间展示
* 基于离线生成3D验证码的问题三维空间图片和答案图片，问题图片用户客户端展示，答案图片用于服务端答案验证
* 问题结合了语序，语义，空间推理，物理平衡多种维度算法，解空间难度大大提升，增强的AI破解的难度

随着互联网技术的进步，黑灰产也逐渐形成了团体化、工具化，对企业造成的危害也越发严重。在与黑产的对抗中，同盾科技反欺诈服务，依托大数据分析技术，提供内容安全,业务安全,云安全,人机验证,行为验证码,营销反作弊,注册保护,登录保护,图像鉴黄,文本过滤,敏感词过滤,黑产情报等产品用于黑产对抗，为企业提供安全解决方案。

同盾智能验证码已由单点对抗演进到体系对抗，多维度的混编一体对抗，采用新型的验证形式，提升用户体验的同时增强安全防御能力。安全是互联网公司的生命，同盾验证码将协助用户共同筑起这道安全篱笆。

## 参考资料 ##

[验证码的前世今生（今生篇）]( https://link.juejin.im?target=https%3A%2F%2Fwww.freebuf.com%2Farticles%2Fweb%2F102276.html )

[浅谈Web安全验证码]( https://link.juejin.im?target=http%3A%2F%2Fblog.nsfocus.net%2Fdiscussion-web-security-authentication-code%2F )

[Puppeteer-无头浏览器简介]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F40103840 )

[node识别验证码]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000015134802 )

[固定閾值(threshold)]( https://link.juejin.im?target=http%3A%2F%2Fmonkeycoding.com%2F%3Fp%3D593 )

[利用Tesseract图片文字识别初探]( https://link.juejin.im?target=https%3A%2F%2Ftonydeng.github.io%2F2016%2F07%2F28%2Fon-the-use-of-tesseract-picture-text-recognition%2F )

[tesseract如何限定识别的文字]( https://link.juejin.im?target=https%3A%2F%2Fmy.oschina.net%2Fu%2F2396236%2Fblog%2F1621590 )

[【模型解读】Inception结构，你看懂了吗]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F41691301 )

[通过迁移训练来定制 TensorFlow 模型]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F40ad065c8cc7 )

[复旦大学肖仰华：12306的验证码已不再安全，未来属于智能验证码]( https://link.juejin.im?target=https%3A%2F%2Fwww.leiphone.com%2Fnews%2F201704%2FbJ9OtS2IfrRpyoUT.html )