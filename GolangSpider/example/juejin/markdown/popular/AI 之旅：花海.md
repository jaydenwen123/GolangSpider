# AI 之旅：花海 #

## 概述 ##

正如《AI 之旅：启程》一文所说，“机器学习作为实现 AI 的主要手段，涉及到的知识和领域非常多，而 Google 提供的 TensorFlow 平台让普通人创建或训练机器学习模型成为可能，作为普通人，你无法提出革命性的机器学习理论，甚至无法理解很多数学知识，但是有了 TensorFlow 平台，有了一些训练好的模型，你就可以创建自己想要的，能帮助你自己或者帮助他人解决问题的模型”。 [TensorFlow]( https://link.juejin.im?target=https%3A%2F%2Fwww.tensorflow.org ) 平台提供了很多理论的具体实现，它们就像一个个功能良好的黑匣子，你不需要知道它背后的理论原理和数学公式多么地复杂，你也不需要关心具体的实现细节，你只需要知道它能帮你完成哪些工作就可以了

## TensorFlow Lite ##

TensorFlow 不但提供了电脑端的实现，还提供了针对移动端、Web 端、IoT 以及云端等几乎所有平台的实现。由于移动端、嵌入式设备、IoT 设备本身计算能力和存储空间有限，所以对实现的要求更高，TensorFlow 提供了针对这些平台专门精简优化的 Lite 版本，叫 TensorFlow Lite，有了它，这些设备就可以在本地（on-device）实现低延时低空间占用的机器学习推理了。你可能会问，我可以直接在后台服务器（云端）上实现模型的推理啊，这样就不用考虑硬件的性能限制了，事实上的确可以这样，但是这也会带来一些问题，如受到网络传输速度的影响，从你发送推理网络请求到你接收到请求结果的时间跨度可能会很长，你没法保证低延时，而且一旦数据离开了设备，你没办法保证隐私性和安全性，你也没办法保证设备总能联网，而且网络请求是一个特别耗电的行为，所以你在选择方案时需要慎重权衡一下
TensorFlow Lite 包含两个主要组件，一个是解释器（interpreter），可以在不同的硬件类型（手机、嵌入式 Linux 设备和微控制器等）上运行专门优化的模型。一个是转换器（converter），可以将 TensorFlow 模型转化成供解释器使用的高效形式，以便优化空间占用并提升性能
TensorFlow Lite 目前提供了 5 个训练好（pre-trained）的模型：

* 图像分类（Image classification）：可以识别图像中对象的类别，包括人、动物、植物、地点等
* 对象检测（Object detection）：可以识别出图像中多个对象的类型和位置
* 智能回复（Smart reply）：可以根据聊天内容自动生成回复建议，生成的建议是上下文相关的，且是一键式（one-touch）的，可以帮助用户快速回复到来的消息
* 姿态估计（Pose estimation）：可以通过评估关键身体关节的位置来估计人物在图像或视频中的姿态
* 图像语义分割（Segmentation）：可以将语义标签（如人，狗，猫）分配给图像中的每个像素

这 5 个模型你可以直接使用它，也可以通过迁移学习（transfer learning）的方式重新训练成新的模型，但是如果你想使用其它训练好的 TensorFlow 模型，你可能需要先转换成 TensorFlow Lite 格式才能在 TensorFlow Lite 中使用它
由于开发进度和水平的限制，目前 TensorFlow Lite 只支持 TensorFlow 操作（operations，简称 ops）的有限子集，如果发现有些操作不支持，可以通过一些配置添加缺失的 TensorFlow ops

### 软硬件支持 ###

TensorFlow 的计算量一般非常大，传统 CPU 的计算能力和传统算法的计算能力很难满足需求，所以我们需要针对 AI 算法特殊优化的硬件和软件
从硬件层面上来说，GPU 的计算能力尤其是浮点矩阵运算的能力比 CPU 要强大，所以可以把这部分工作交给 GPU 去做以便加快推理速度提高效率，事实证明，这个速度的提升是非常可观的， Pixel 3 手机上 MobileNet v1 图像分类模型在开启 GPU 加速后至少能快 5.5 倍，而且比 CPU 耗更少的电产生更少的热量。当然，除了 GPU，像 NPU（Neural Networks Processing Unit）、TPU（Tensor Processing Unit）、DSP（Digital Signal Processor）等硬件也能胜任这些工作
从软件层面来说，AI 相关的算法更加丰富，为了从底层就支持 AI 算法，Google 从 Android 8.1（API level 27）开始新增了 NNAPI（Android Neural Networks API），一个为了针对机器学习的计算密集型操作而专门设计的 C 语言 API，开发者一般通过机器学习框架间接使用这些 API，框架会利用 NNAPI 在支持的设备上执行硬件加速下的推理操作

### 转换器 ###

转换器不但是用来把 TensorFlow 模型转换成 TensorFlow Lite 格式（FlatBuffer）的工具，也是最常见的用来优化模型的工具。FlatBuffer 是一个高效的开源跨平台序列化库，相对于 protobuf（protocol buffers）来说更加简单轻量。推荐在 Python API 中使用转换器：

` import tensorflow as tf converter = tf.lite.TFLiteConverter.from_saved_model(saved_model_dir) tflite_model = converter.convert() open( "converted_model.tflite" , "wb" ).write(tflite_model) 复制代码`

![converter](https://user-gold-cdn.xitu.io/2019/6/4/16b2196398d9b5ac?imageView2/0/w/1280/h/960/ignore-error/1) 生成的 .tflite 文件就是供解释器使用的模型，将它和标签文件一同放到 assets 目录下，为了避免被 AAPT 工具压缩，需要添加配置项：

` aaptOptions { noCompress "tflite" } 复制代码`

### 解释器 ###

` implementation 'org.tensorflow:tensorflow-lite:0.0.0-nightly' 复制代码` ` ndk { abiFilters 'armeabi-v7a' , 'arm64-v8a' } 复制代码`

推理的过程很简单，就是先加载模型文件，然后用它实例化一个 ` Interpreter` 解释器，然后运行这个解释器就可以得到结果了
建议预加载模型文件并完成内存映射以便加快加载时间同时减少内存脏页：

` private MappedByteBuffer loadModelFile (Activity activity) throws IOException { AssetFileDescriptor fileDescriptor = activity.getAssets().openFd( "mobilenet_v1_1.0_224.tflite" ); FileInputStream inputStream = new FileInputStream(fileDescriptor.getFileDescriptor()); FileChannel fileChannel = inputStream.getChannel(); long startOffset = fileDescriptor.getStartOffset(); long declaredLength = fileDescriptor.getDeclaredLength(); return fileChannel.map(FileChannel.MapMode.READ_ONLY, startOffset, declaredLength); } 复制代码`

在构造解释器时可以通过类型为 ` Interpreter.Options` 的参数去配置这个解释器，比如使用的线程数、使用 NNAPI、使用 GPU 加速等：

` tfliteModel = loadModelFile(activity); Interpreter.Options options = new Interpreter.Options(); options.setNumThreads( 1 ); tflite = new Interpreter(tfliteModel, options); 复制代码`

以图像分类模型的简单应用为例，我们的输入是个 Bitmap 对象，输出应该是包含标签和对应概率的实体类列表：

![demo](https://user-gold-cdn.xitu.io/2019/6/5/16b265997150b86b?imageslim)

` public List<Recognition> recognizeImage ( final Bitmap bitmap) { ... } 复制代码` ` public static class Recognition { private final String id; private final String title; private final Float confidence; private RectF location; ... } 复制代码`

而为了更高效地处理图片数据，需要把 bitmap 转成 ` ByteBuffer` 格式， ` ByteBuffer` 将图片表示成每个颜色通道 3 个字节的一维数组，我们需要预分配这个内存并调用 ` order(ByteOrder.nativeOrder())` 以保证每一位都以 native order 顺序存储：

` imgData = ByteBuffer.allocateDirect( DIM_BATCH_SIZE * getImageSizeX() * getImageSizeY() * DIM_PIXEL_SIZE * getNumBytesPerChannel()); imgData.order(ByteOrder.nativeOrder()); ... private void convertBitmapToByteBuffer (Bitmap bitmap) { if (imgData == null ) { return ; } imgData.rewind(); bitmap.getPixels(intValues, 0 , bitmap.getWidth(), 0 , 0 , bitmap.getWidth(), bitmap.getHeight()); int pixel = 0 ; for ( int i = 0 ; i < getImageSizeX(); ++i) { for ( int j = 0 ; j < getImageSizeY(); ++j) { final int val = intValues[pixel++]; addPixelValue(val); } } } protected void addPixelValue ( int pixelValue) { imgData.putFloat((((pixelValue >> 16 ) & 0xFF ) - IMAGE_MEAN) / IMAGE_STD); imgData.putFloat((((pixelValue >> 8 ) & 0xFF ) - IMAGE_MEAN) / IMAGE_STD); imgData.putFloat(((pixelValue & 0xFF ) - IMAGE_MEAN) / IMAGE_STD); } 复制代码`

而模型的输出是个二维数组，表示预测的每个标签的概率：

` private List<String> labels; private float [][] labelProbArray; ... labels = loadLabelList(activity); labelProbArray = new float [ 1 ][labels.size()]; ... private List<String> loadLabelList (Activity activity) throws IOException { List<String> labels = new ArrayList<String>(); BufferedReader reader = new BufferedReader( new InputStreamReader(activity.getAssets().open( "labels.txt" ))); String line; while ((line = reader.readLine()) != null ) { labels.add(line); } reader.close(); return labels; } 复制代码`

然后利用这个输入输出运行推理就行了：

` public void runInference () { tflite.run(imgData, labelProbArray); } 复制代码`

最后将这个 labelProbArray 转成 ` List<Recognition>` 就可以用于可视化输出的显示了

### 最佳实践 ###

我们发现 ` mobilenet_v1_1.0_224.tflite` 模型文件的大小有 16.9 MB 这么大，那我们有没有可能把它压缩的更小呢？有，由于模型使用了浮点的权重和激活函数，我们就可以通过量化（Quantization）的方式把模型压缩至少 4 倍，量化有两种方式，一种是训练后量化（post-training quantization），不需要重新训练模型，不过可能会有准确率的损失，如果这个损失超过了可接受的阈值，就只能使用另一种方式，即量化训练（quantized training）了。如果我们选择使用 ` mobilenet_v1_1.0_224_quant.tflite` 模型文件，我们会发现它只有 4.3 MB
另一个比较重要的点是要学会权衡，有些模型虽然很复杂很大但是准确率很高，有些模型虽然准确率不是很高但是更小巧执行更快，你需要根据实际情况选择最适合的模型

![tradeoff](https://user-gold-cdn.xitu.io/2019/6/4/16b2196398ee2ed7?imageView2/0/w/1280/h/960/ignore-error/1) 使用硬件加速并不总是最好的选择，有时候开启硬件加速甚至不如不开，所以你最好做好基准测试（benchmark）

## 参考 ##

* [TensorFlow Lite]( https://link.juejin.im?target=https%3A%2F%2Fwww.tensorflow.org%2Flite%2F )
* [Example]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fshangmingchao%2FAIDemo )