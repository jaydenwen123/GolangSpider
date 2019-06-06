# vue实现录音功能（pc端） #

> 
> 
> 
> 录音功能一般来说在移动端比较常见，但是在pc端也要实现按住说话的功能呢？项目需求：按住说话，时长不超过60秒，生成语音文件并上传，我这里用的是recorder.js
> 
> 
> 

## 1.项目中新建一个recorder.js文件，内容如下，也可在百度上直接搜一个 ##

` // 兼容
window.URL = window.URL || window.webkitURL
navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia

let HZRecorder = function ( stream, config ) {
config = config || {}
config.sampleBits = config.sampleBits || 8 // 采样数位 8, 16
config.sampleRate = config.sampleRate || ( 44100 / 6 ) // 采样率(1/6 44100)
let context = new ( window.webkitAudioContext || window.AudioContext)()
let audioInput = context.createMediaStreamSource(stream)
let createScript = context.createScriptProcessor || context.createJavaScriptNode
let recorder = createScript.apply(context, [ 4096 , 1 , 1 ])
let audioData = {
size : 0 , // 录音文件长度
buffer: [], // 录音缓存
inputSampleRate: context.sampleRate, // 输入采样率
inputSampleBits: 16 , // 输入采样数位 8, 16
outputSampleRate: config.sampleRate, // 输出采样率
oututSampleBits: config.sampleBits, // 输出采样数位 8, 16
input: function ( data ) {
this.buffer.push( new Float32Array (data))
this.size += data.length
},
compress : function ( ) { // 合并压缩
// 合并
let data = new Float32Array ( this.size)
let offset = 0
for ( let i = 0 ; i < this.buffer.length; i++) {
data.set( this.buffer[i], offset)
offset += this.buffer[i].length
}
// 压缩
let compression = parseInt ( this.inputSampleRate / this.outputSampleRate)
let length = data.length / compression
let result = new Float32Array (length)
let index = 0 ; let j = 0
while (index < length) {
result[index] = data[j]
j += compression
index++
}
return result
},
encodeWAV : function ( ) {
let sampleRate = Math.min( this.inputSampleRate, this.outputSampleRate)
let sampleBits = Math.min( this.inputSampleBits, this.oututSampleBits)
let bytes = this.compress()
let dataLength = bytes.length * (sampleBits / 8 )
let buffer = new ArrayBuffer ( 44 + dataLength)
let data = new DataView (buffer)

let channelCount = 1 // 单声道
let offset = 0

let writeString = function ( str ) {
for ( let i = 0 ; i < str.length; i++) {
data.setUint8(offset + i, str.charCodeAt(i))
}
}

// 资源交换文件标识符
writeString( 'RIFF' ); offset += 4
// 下个地址开始到文件尾总字节数,即文件大小-8
data.setUint32(offset, 36 + dataLength, true ); offset += 4
// WAV文件标志
writeString( 'WAVE' ); offset += 4
// 波形格式标志
writeString( 'fmt ' ); offset += 4
// 过滤字节,一般为 0x10 = 16
data.setUint32(offset, 16 , true ); offset += 4
// 格式类别 (PCM形式采样数据)
data.setUint16(offset, 1 , true ); offset += 2
// 通道数
data.setUint16(offset, channelCount, true ); offset += 2
// 采样率,每秒样本数,表示每个通道的播放速度
data.setUint32(offset, sampleRate, true ); offset += 4
// 波形数据传输率 (每秒平均字节数) 单声道×每秒数据位数×每样本数据位/8
data.setUint32(offset, channelCount * sampleRate * (sampleBits / 8 ), true ); offset += 4
// 快数据调整数 采样一次占用字节数 单声道×每样本的数据位数/8
data.setUint16(offset, channelCount * (sampleBits / 8 ), true ); offset += 2
// 每样本数据位数
data.setUint16(offset, sampleBits, true ); offset += 2
// 数据标识符
writeString( 'data' ); offset += 4
// 采样数据总数,即数据总大小-44
data.setUint32(offset, dataLength, true ); offset += 4
// 写入采样数据
if (sampleBits === 8 ) {
for ( let i = 0 ; i < bytes.length; i++ , offset++) {
let s = Math.max( -1 , Math.min( 1 , bytes[i]))
let val = s < 0 ? s * 0x8000 : s * 0x7FFF
val = parseInt ( 255 / ( 65535 / (val + 32768 )))
data.setInt8(offset, val, true )
}
} else {
for ( let i = 0 ; i < bytes.length; i++ , offset += 2 ) {
let s = Math.max( -1 , Math.min( 1 , bytes[i]))
data.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7FFF , true )
}
}

return new Blob([data], { type : 'audio/mp3' })
}
}

// 开始录音
this.start = function ( ) {
audioInput.connect(recorder)
recorder.connect(context.destination)
}

// 停止
this.stop = function ( ) {
recorder.disconnect()
}

// 获取音频文件
this.getBlob = function ( ) {
this.stop()
return audioData.encodeWAV()
}

// 回放
this.play = function ( audio ) {
let downRec = document.getElementById( 'downloadRec' )
downRec.href = window.URL.createObjectURL( this.getBlob())
downRec.download = new Date ().toLocaleString() + '.mp3'
audio.src = window.URL.createObjectURL( this.getBlob())
}

// 上传
this.upload = function ( url, callback ) {
let fd = new FormData()
fd.append( 'audioData' , this.getBlob())
let xhr = new XMLHttpRequest()
/* eslint-disable */
if (callback) {
xhr.upload.addEventListener( 'progress' , function ( e ) {
callback( 'uploading' , e)
}, false )
xhr.addEventListener( 'load' , function ( e ) {
callback( 'ok' , e)
}, false )
xhr.addEventListener( 'error' , function ( e ) {
callback( 'error' , e)
}, false )
xhr.addEventListener( 'abort' , function ( e ) {
callback( 'cancel' , e)
}, false )
}
/* eslint-disable */
xhr.open( 'POST' , url)
xhr.send(fd)
}

// 音频采集
recorder.onaudioprocess = function ( e ) {
audioData.input(e.inputBuffer.getChannelData( 0 ))
// record(e.inputBuffer.getChannelData(0));
}
}
// 抛出异常
HZRecorder.throwError = function ( message ) {
alert(message)
throw new function ( ) { this.toString = function ( ) { return message } }()
}
// 是否支持录音
HZRecorder.canRecording = (navigator.getUserMedia != null )
// 获取录音机
HZRecorder.get = function ( callback, config ) {
if (callback) {
if (navigator.getUserMedia) {
navigator.getUserMedia(
{ audio : true } // 只启用音频
, function ( stream ) {
let rec = new HZRecorder(stream, config)
callback(rec)
}
, function ( error ) {
switch (error.code || error.name) {
case 'PERMISSION_DENIED' :
case 'PermissionDeniedError' :
HZRecorder.throwError( '用户拒绝提供信息。' )
break
case 'NOT_SUPPORTED_ERROR' :
case 'NotSupportedError' :
HZRecorder.throwError( '浏览器不支持硬件设备。' )
break
case 'MANDATORY_UNSATISFIED_ERROR' :
case 'MandatoryUnsatisfiedError' :
HZRecorder.throwError( '无法发现指定的硬件设备。' )
break
default :
HZRecorder.throwError( '无法打开麦克风。异常信息:' + (error.code || error.name))
break
}
})
} else {
HZRecorder.throwErr( '当前浏览器不支持录音功能。' ); return
}
}
}
export default HZRecorder
复制代码`

## 2.页面中使用，具体如下 ##

` < template >
< div class = "wrap" >
< el-form v-model = "form" >
< el-form-item >
< input type = "button" class = "btn-record-voice" @ mousedown.prevent = "mouseStart" @ mouseup.prevent = "mouseEnd" v-model = "form.time" />
< audio v-if = "form.audioUrl" :src = "form.audioUrl" controls = "controls" class = "content-audio" style = "display: block;" > 语音 </ audio >
</ el-form-item >
< el-form >
</ div >
</ template >

< script >
// 引入recorder.js
import recording from '@/js/recorder/recorder.js'
export default {
data() {
return {
form : {
time : '按住说话(60秒)' ,
audioUrl : ''
},
num : 60 , // 按住说话时间
recorder: null ,
interval : '' ,
audioFileList : [], // 上传语音列表
startTime: '' , // 语音开始时间
endTime: '' , // 语音结束
}
},
methods : {
// 清除定时器
clearTimer () {
if ( this.interval) {
this.num = 60
clearInterval( this.interval)
}
},
// 长按说话
mouseStart () {
this.clearTimer()
this.startTime = new Date ().getTime()
recording.get( ( rec ) => {
// 当首次按下时，要获取浏览器的麦克风权限，所以这时要做一个判断处理
if (rec) {
// 首次按下，只调用一次
if ( this.flag) {
this.mouseEnd()
this.flag = false
} else {
this.recorder = rec
this.interval = setInterval( () => {
if ( this.num <= 0 ) {
this.recorder.stop()
this.num = 60
this.clearTimer()
} else {
this.num--
this.time = '松开结束（' + this.num + '秒）'
this.recorder.start()
}
}, 1000 )
}
}
})
},
// 松开时上传语音
mouseEnd () {
this.clearTimer()
this.endTime = new Date ().getTime()
if ( this.recorder) {
this.recorder.stop()
// 重置说话时间
this.num = 60
this.time = '按住说话（' + this.num + '秒）'
// 获取语音二进制文件
let bold = this.recorder.getBlob()
// 将获取的二进制对象转为二进制文件流
let files = new File([bold], 'test.mp3' , { type : 'audio/mp3' , lastModified : Date.now()})
let fd = new FormData()
fd.append( 'file' , files)
fd.append( 'tenantId' , 3 ) // 额外参数，可根据选择填写
// 这里是通过上传语音文件的接口，获取接口返回的路径作为语音路径
this.uploadFile(fd)
}
}
}
}
</ script >

< style scoped >
</ style >
复制代码`

## 3.除了上述代码中的注释外，还有一些地方需要注意 ##

* 上传语音时，一般会有两个参数，一个是语音的路径，一个是语音的时长，路径直接就是 ` this.form.audioUrl` ,不过时长这里需要注意的是，由于我们一开始设置了定时器是有一秒的延迟，所以，要在获取到的时长基础上在减去一秒
* 初次按住说话一定要做判断，不然就会报错啦
* 第三点也是很重要的一点，因为我是在本地项目中测试的，可以实现录音功能，但是打包到测试环境后，就无法访问麦克风，经过多方尝试后，发现是由于我们测试环境的地址是http://***,而在谷歌浏览器中有这样一种安全策略，只允许在localhost下及https下才可以访问 ，因此换一下就完美的解决了这个问题了
* 在使用过程中，针对不同的浏览器可能会有些兼容性的问题，如果遇到了还需自己单独处理下

> 
> 
> 
> 好多东西都是在项目中才学会的，所以要趁着记忆还清晰，赶紧记下来，如果上述有什么不对的地方，还请指正
> 
>