# Lightingcolorfilter 滤镜 #

Paint除了通过 **setColor()/setARGB()/setAlpha()** 等方法设置paint的颜色外，还可以通过设置滤镜ColorFilter来改变绘制出来的颜色。ColorFilter对应以下三个子类。

**颜色矩阵颜色过滤器** ： [ColorMatrixColorFilter]( https://link.juejin.im?target=http%3A%2F%2Fandroiddoc.qiniudn.com%2Freference%2Fandroid%2Fgraphics%2FColorMatrixColorFilter.html )

**光照色彩过滤器** ： [LightingColorFilter]( https://link.juejin.im?target=http%3A%2F%2Fandroiddoc.qiniudn.com%2Freference%2Fandroid%2Fgraphics%2FLightingColorFilter.html )

**混排颜色过滤器滤器** [PorterDuffColorFilter]( https://link.juejin.im?target=http%3A%2F%2Fandroiddoc.qiniudn.com%2Freference%2Fandroid%2Fgraphics%2FPorterDuffColorFilter.html )

## 1 LightingColorFilter ##

先来看看这个ColorFilter的构造器，它接受两个参数，一个是mul，一个是add

` LightingColorFilter( int mul, int add) 复制代码`

比如你现在设置paint的color是#1b8fe6， **然后通过滤镜后想要让其颜色不变，那么可以设置mul为0xffffff，add为0x000000** 。 **前两个ff表示红色的十六进制 中间两个ff表示绿色的十六进制 后边两个ff表示蓝色的十六进制 转换成十进制就是0-255的范围值 值越大 表示对应的颜色就越深** **add为0x000000 这是颜色加深的效果 比如 想要 绿色增强 就可以把 中间两个00改成 **000100** , **000800** ,** **008800** **直到255 最大值也是255  对应十六进制的最大值为** **0x00ff00** **** 其实这个玩意是这样去计算的，我们记paint设置的颜色中三基色的值为R，G，B，mul中的三基色的值为mul.R，mul.G，mul.B，add中的三基色的值为add.R，add.G，add.B，那么最终的三基色为jis

` 经过滤镜后的R = R * mul.R / 0xff + add.R 经过滤镜后的G = G * mul.G / 0xff + add.G 经过滤镜后的B = B * mul.B / 0xff + add.B 复制代码`

### RGBA模型： ###

> 
> 
> 
> RGBA不知道你听过没，黄绿蓝知道了吧，光的三基色，而RAGB则是在此的基础上多了一个透明度！ **R(Red红色)** ， **G(Green绿色)**
> ， **B(Blue蓝色)** ， **A(Alpha透明度)** ；另外要和颜料的三
> 原色区分开来哦，最明显的区别就是颜料的三原色中用黄色替换了光三基色中的绿色！
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/5/16b265c8501a2979?imageView2/0/w/1280/h/960/ignore-error/1)

### 接下来 我们使用 ###

### PorterDuff.Mode 图层混合 和  LightingColorFilter 实现一个滤镜效果,先上效果图 ###

### 正常图片 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26687828390b4?imageView2/0/w/1280/h/960/ignore-error/1)

### 使用滤镜后的图片 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2669a76fcb69e?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2669cab3d7e6a?imageView2/0/w/1280/h/960/ignore-error/1)

### ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2669ebf918141?imageView2/0/w/1280/h/960/ignore-error/1)

` public class PictureView extends View { private LightingColorFilter lightingColorFilter; private Paint mPaint; private int mWidth, mHeight; Bitmap mBitmap; public PictureView(Context context) { this(context, null); } public PictureView(Context context, AttributeSet attrs) { this(context, attrs, 0); } public PictureView(Context context, AttributeSet attrs, int defStyleAttr) { super(context, attrs, defStyleAttr); init(); } private void init () { //初始化画笔 mPaint = new Paint(); mPaint.setColor(Color.RED); mPaint.setStyle(Paint.Style.FILL_AND_STROKE); mBitmap = BitmapFactory.decodeResource(getResources(), R.drawable.buteys).copy(Bitmap.Config.ARGB_8888, true ); } @Override protected void onSizeChanged(int w, int h, int oldw, int oldh) { super.onSizeChanged(w, h, oldw, oldh); mWidth = w; mHeight = h; } @Override protected void onDraw(Canvas canvas) { super.onDraw(canvas); set BackgroundColor(Color.WHITE); //离屏绘制 这里将图像合成的处理放到离屏缓存中进行 int layerId = canvas.saveLayer(0, 0, getWidth(), getHeight(), mPaint, Canvas.ALL_SAVE_FLAG); //目标图 canvas.drawBitmap(mBitmap, mWidth / 2 - mBitmap.getWidth() / 2-120, mHeight / 3 - mBitmap.getHeight() / 4, mPaint); mPaint.setXfermode(new PorterDuffXfermode(PorterDuff.Mode.DST_IN)); canvas.drawBitmap(createCircleBitmap(mWidth, mHeight), 0, 0, mPaint); //清除混合模式 mPaint.setXfermode(null); canvas.restoreToCount(layerId); } //画圆src public Bitmap createCircleBitmap(int width, int height) { Bitmap bitmap = Bitmap.createBitmap(width, height, Bitmap.Config.ARGB_8888); Canvas canvas = new Canvas(bitmap); Paint scrPaint = new Paint(Paint.ANTI_ALIAS_FLAG); scrPaint.setColor(0xFFFFCC44); canvas.drawCircle(width / 2, height / 2, 500, scrPaint); return bitmap; } /** * 获取RGB值 * * @param r * @param g * @param b */ public void changeRGB(int r, int g, int b) { int mr = 0; int mg = 0; int mb = 0; int ar = 0; int ag = 0; int ab = 0; if (r < 0) { mr = r + 255; ar = 0; } else if (r == 0) { mr = 255; ar = 0; } else { mr = 255; ar = r; } if (g < 0) { mg = g + 255; ag = 0; } else if (g == 0) { mg = 255; ag = 0; } else { mg = 255; ag = g; } if (b < 0) { mb = b + 255; ab = 0; } else if (b == 0) { mb = 255; ab = 0; } else { mb = 255; ab = b; } lightingColorFilter = new LightingColorFilter(0xffffff,0x000000); lightingColorFilter = new LightingColorFilter( Integer.valueOf(transString(Integer.toHexString(mr)) + transString(Integer.toHexString(mg)) + transString(Integer.toHexString(mb)), 16), Integer.valueOf(transString(Integer.toHexString(ar)) + transString(Integer.toHexString(ag)) + transString(Integer.toHexString(ab)), 16)); mPaint.setColorFilter(lightingColorFilter); invalidate(); } /** * 处理RGB * * @param s * @ return */ public String transString(String s) { if (s.length() == 1) { return "0" + s; } else { return s; } } } 复制代码`

` <?xml version= "1.0" encoding= "utf-8" ?> <LinearLayout xmlns:android= "http://schemas.android.com/apk/res/android" xmlns:app= "http://schemas.android.com/apk/res-auto" xmlns:tools= "http://schemas.android.com/tools" android:layout_width= "match_parent" android:layout_height= "match_parent" android:orientation= "vertical" tools:context= ".MainActivity" > <com.rx.myapplication.PictureView android:id= "@+id/pv" android:layout_width= "match_parent" android:layout_height= "350dp" /> <LinearLayout android:layout_width= "match_parent" android:layout_height= "match_parent" android:orientation= "vertical" > <LinearLayout android:layout_width= "match_parent" android:layout_height= "0dp" android:layout_weight= "1" android:gravity= "center_vertical" android:orientation= "horizontal" android:paddingRight= "10dp" > <TextView android:layout_width= "wrap_content" android:layout_height= "wrap_content" android:layout_marginLeft= "10dp" android:layout_marginRight= "10dp" android:text= "R" android:textSize= "18sp" /> <SeekBar android:id= "@+id/sb_r" android:layout_width= "0dp" android:layout_height= "wrap_content" android:layout_weight= "8" /> <TextView android:id= "@+id/tv_r" android:layout_width= "0dp" android:layout_height= "wrap_content" android:layout_weight= "1" android:gravity= "center" /> </LinearLayout> <LinearLayout android:layout_width= "match_parent" android:layout_height= "0dp" android:layout_weight= "1" android:gravity= "center_vertical" android:orientation= "horizontal" android:paddingRight= "10dp" > <TextView android:layout_width= "wrap_content" android:layout_height= "wrap_content" android:layout_marginLeft= "10dp" android:layout_marginRight= "10dp" android:text= "G" android:textSize= "18sp" /> <SeekBar android:id= "@+id/sb_g" android:layout_width= "0dp" android:layout_height= "wrap_content" android:layout_weight= "8" /> <TextView android:id= "@+id/tv_g" android:layout_width= "0dp" android:layout_height= "wrap_content" android:layout_weight= "1" android:gravity= "center" /> </LinearLayout> <LinearLayout android:layout_width= "match_parent" android:layout_height= "0dp" android:layout_weight= "1" android:gravity= "center_vertical" android:orientation= "horizontal" android:paddingRight= "10dp" > <TextView android:layout_width= "wrap_content" android:layout_height= "wrap_content" android:layout_marginLeft= "10dp" android:layout_marginRight= "10dp" android:text= "B" android:textSize= "18sp" /> <SeekBar android:id= "@+id/sb_b" android:layout_width= "0dp" android:layout_height= "wrap_content" android:layout_weight= "8" /> <TextView android:id= "@+id/tv_b" android:layout_width= "0dp" android:layout_height= "wrap_content" android:layout_weight= "1" android:gravity= "center" /> </LinearLayout> </LinearLayout> </LinearLayout> 复制代码`

` public class MainActivity extends AppCompatActivity implements SeekBar.OnSeekBarChangeListener { private SeekBar sb_r; private SeekBar sb_g; private SeekBar sb_b; private TextView tv_r; private TextView tv_g; private TextView tv_b; private PictureView pv; @Override protected void onCreate(Bundle savedInstanceState) { super.onCreate(savedInstanceState); set ContentView(R.layout.activity_main); initView(); sb_r.setOnSeekBarChangeListener(this); sb_g.setOnSeekBarChangeListener(this); sb_b.setOnSeekBarChangeListener(this); } private void initView () { pv = (PictureView) findViewById(R.id.pv); sb_r = (SeekBar) findViewById(R.id.sb_r); sb_g = (SeekBar) findViewById(R.id.sb_g); sb_b = (SeekBar) findViewById(R.id.sb_b); tv_r = (TextView) findViewById(R.id.tv_r); tv_g = (TextView) findViewById(R.id.tv_g); tv_b = (TextView) findViewById(R.id.tv_b); sb_r.setMax(510); sb_g.setMax(510); sb_b.setMax(510); sb_r.setProgress(255); sb_g.setProgress(255); sb_b.setProgress(255); } @Override public void onProgressChanged(SeekBar seekBar, int progress, boolean fromUser) { pv.changeRGB(sb_r.getProgress() - 255, sb_g.getProgress() - 255, sb_b.getProgress() - 255); switch (seekBar.getId()) { case R.id.sb_r: tv_r.setText(progress - 255 + "" ); break ; case R.id.sb_g: tv_g.setText(progress - 255 + "" ); break ; case R.id.sb_b: tv_b.setText(progress - 255 + "" ); break ; } } @Override public void onStartTrackingTouch(SeekBar seekBar) { } @Override public void onStopTrackingTouch(SeekBar seekBar) { } } 复制代码`

## 2 PortDuffColorFilter ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b267ddd51cf513?imageView2/0/w/1280/h/960/ignore-error/1)

` PorterDuffColorFilter porterDuffColorFilter=new PorterDuffColorFilter(Color.RED,PorterDuff.Mode.DARKEN); mPaint.setColorFilter(porterDuffColorFilter); invalidate(); 复制代码`

### 最后的效果 红色加深 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b268920e25d7a5?imageView2/0/w/1280/h/960/ignore-error/1)

## 3 ColorMatrixColorFilter ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b268a7eeeb36d9?imageView2/0/w/1280/h/960/ignore-error/1)

### 1、ColorMatrix处理图片原理 ###

ColorMatrix正如我们所翻译的“颜色矩阵”，它是通过矩阵的方式作用与图像的像素来实现的图片的处理。

（1）ColorMAtrix颜色矩阵 颜色矩阵M是一个5*4的矩阵，如图所示。但是在Android中，颜色矩阵M是以一维数组M=[a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t]的方式进行存储的。

（2）颜色的分量矩阵 颜色的分量矩阵分别有R、G、B、A、1组成，用于调整三原色和透明度 （3）变换过程

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26a14858b65b9?imageView2/0/w/1280/h/960/ignore-error/1)

第一行将会影响红色，第二行影响绿色，第三行影响蓝色，最后一行操作的是Alpha值。

ps：默认的ColorMatrix如下面所示，它是不会改变图像的。

1,0,0,0,0

0,1,0,0,0

0,0,1,0,0

0,0,0,1,0

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26a2bb033b944?imageView2/0/w/1280/h/960/ignore-error/1)

### 当使用默认矩阵时 ，图像保持不变, ###

` float [] colorMatrix = { 1,0,0,0,0, //red 0,1,0,0,0, //green 0,0,1,0,0, //blue 0,0,0,1,0 //alpha }; ColorMatrixColorFilter colorMatrixColorFilter = new ColorMatrixColorFilter(colorMatrix); mPaint.setColorFilter(colorMatrixColorFilter); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26a8842b2f0e7?imageView2/0/w/1280/h/960/ignore-error/1)

### 现在我们修改 RGBA值的系数  矩阵的红色变成 2，可以看到红色加深了 ###

` float [] colorMatrix = { 2,0,0,0,0, //red 0,1,0,0,0, //green 0,0,1,0,0, //blue 0,0,0,1,0 //alpha }; ColorMatrixColorFilter colorMatrixColorFilter = new ColorMatrixColorFilter(colorMatrix); mPaint.setColorFilter(colorMatrixColorFilter) 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26ab34125aee2?imageView2/0/w/1280/h/960/ignore-error/1)

### 我们修改 RGBA值的系数 矩阵的绿色变成2，可以看到绿色加深了 ###

` float [] colorMatrix = { 1,0,0,0,0, //red 0,2,0,0,0, //green 0,0,1,0,0, //blue 0,0,0,1,0 //alpha }; ColorMatrixColorFilter colorMatrixColorFilter = new ColorMatrixColorFilter(colorMatrix); mPaint.setColorFilter(colorMatrixColorFilter); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26afd1c91d462?imageView2/0/w/1280/h/960/ignore-error/1)

### 我们也可以修改矩阵的偏移量 来过滤图片，我们让红色和绿色的偏移量都变成60，照片变成黄色系 ###

` float [] colorMatrix = { 1,0,0,0,60, //red 0,1,0,0,60, //green 0,0,1,0,0, //blue 0,0,0,1,0 //alpha };ColorMatrixColorFilter colorMatrixColorFilter = new ColorMatrixColorFilter(colorMatrix); mPaint.setColorFilter(colorMatrixColorFilter); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26b8412b94c54?imageView2/0/w/1280/h/960/ignore-error/1)

` // 胶片效果 复制代码` ` public static final float colormatrix_fanse[] = { -1.0f, 0.0f, 0.0f, 0.0f, 255.0f, 0.0f, -1.0f, 0.0f, 0.0f, 255.0f, 0.0f, 0.0f, -1.0f, 0.0f, 255.0f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f}; mColorMatrixColorFilter = new ColorMatrixColorFilter(colormatrix_fanse); mPaint.setColorFilter(mColorMatrixColorFilter); canvas.drawBitmap(mBitmap, 100, 0, mPaint); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26c4e23839d85?imageView2/0/w/1280/h/960/ignore-error/1)

## 常用效果 ##

` // 黑白 public static final float colormatrix_heibai[] = { 0.8f, 1.6f, 0.2f, 0, -163.9f, 0.8f, 1.6f, 0.2f, 0, -163.9f, 0.8f, 1.6f, 0.2f, 0, -163.9f, 0, 0, 0, 1.0f, 0}; // 怀旧 public static final float colormatrix_huajiu[] = { 0.2f, 0.5f, 0.1f, 0, 40.8f, 0.2f, 0.5f, 0.1f, 0, 40.8f, 0.2f, 0.5f, 0.1f, 0, 40.8f, 0, 0, 0, 1, 0}; // 哥特 public static final float colormatrix_gete[] = { 1.9f, -0.3f, -0.2f, 0, -87.0f, -0.2f, 1.7f, -0.1f, 0, -87.0f, -0.1f, -0.6f, 2.0f, 0, -87.0f, 0, 0, 0, 1.0f, 0}; // 淡雅 public static final float colormatrix_danya[] = { 0.6f, 0.3f, 0.1f, 0, 73.3f, 0.2f, 0.7f, 0.1f, 0, 73.3f, 0.2f, 0.3f, 0.4f, 0, 73.3f, 0, 0, 0, 1.0f, 0}; // 蓝调 public static final float colormatrix_landiao[] = { 2.1f, -1.4f, 0.6f, 0.0f, -71.0f, -0.3f, 2.0f, -0.3f, 0.0f, -71.0f, -1.1f, -0.2f, 2.6f, 0.0f, -71.0f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f}; // 光晕 public static final float colormatrix_guangyun[] = { 0.9f, 0, 0, 0, 64.9f, 0, 0.9f, 0, 0, 64.9f, 0, 0, 0.9f, 0, 64.9f, 0, 0, 0, 1.0f, 0}; // 梦幻 public static final float colormatrix_menghuan[] = { 0.8f, 0.3f, 0.1f, 0.0f, 46.5f, 0.1f, 0.9f, 0.0f, 0.0f, 46.5f, 0.1f, 0.3f, 0.7f, 0.0f, 46.5f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f}; // 酒红 public static final float colormatrix_jiuhong[] = { 1.2f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.9f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.8f, 0.0f, 0.0f, 0, 0, 0, 1.0f, 0}; // 胶片 public static final float colormatrix_fanse[] = { -1.0f, 0.0f, 0.0f, 0.0f, 255.0f, 0.0f, -1.0f, 0.0f, 0.0f, 255.0f, 0.0f, 0.0f, -1.0f, 0.0f, 255.0f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f}; // 湖光掠影 public static final float colormatrix_huguang[] = { 0.8f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.9f, 0.0f, 0.0f, 0, 0, 0, 1.0f, 0}; // 褐片 public static final float colormatrix_hepian[] = { 1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.8f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.8f, 0.0f, 0.0f, 0, 0, 0, 1.0f, 0}; // 复古 public static final float colormatrix_fugu[] = { 0.9f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.8f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.5f, 0.0f, 0.0f, 0, 0, 0, 1.0f, 0}; // 泛黄 public static final float colormatrix_huan_huang[] = { 1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.5f, 0.0f, 0.0f, 0, 0, 0, 1.0f, 0}; // 传统 public static final float colormatrix_chuan_tong[] = { 1.0f, 0.0f, 0.0f, 0, -10f, 0.0f, 1.0f, 0.0f, 0, -10f, 0.0f, 0.0f, 1.0f, 0, -10f, 0, 0, 0, 1, 0}; // 胶片2 public static final float colormatrix_jiao_pian[] = { 0.71f, 0.2f, 0.0f, 0.0f, 60.0f, 0.0f, 0.94f, 0.0f, 0.0f, 60.0f, 0.0f, 0.0f, 0.62f, 0.0f, 60.0f, 0, 0, 0, 1.0f, 0}; // 锐色 public static final float colormatrix_ruise[] = { 4.8f, -1.0f, -0.1f, 0, -388.4f, -0.5f, 4.4f, -0.1f, 0, -388.4f, -0.5f, -1.0f, 5.2f, 0, -388.4f, 0, 0, 0, 1.0f, 0}; // 清宁 public static final float colormatrix_qingning[] = { 0.9f, 0, 0, 0, 0, 0, 1.1f, 0, 0, 0, 0, 0, 0.9f, 0, 0, 0, 0, 0, 1.0f, 0}; // 浪漫 public static final float colormatrix_langman[] = { 0.9f, 0, 0, 0, 63.0f, 0, 0.9f, 0, 0, 63.0f, 0, 0, 0.9f, 0, 63.0f, 0, 0, 0, 1.0f, 0}; // 夜色 public static final float colormatrix_yese[] = { 1.0f, 0.0f, 0.0f, 0.0f, -66.6f, 0.0f, 1.1f, 0.0f, 0.0f, -66.6f, 0.0f, 0.0f, 1.0f, 0.0f, -66.6f, 0.0f, 0.0f, 0.0f, 1.0f, 0.0f}; 复制代码`

## ColorMatrix 类的使用 ##

### 正常图片 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d7d1701843f?imageView2/0/w/1280/h/960/ignore-error/1)

` ColorMatrix cm = new ColorMatrix(); //亮度调节 cm.setScale(1,2,1,1); 绿色高亮 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d83b6b50a8a?imageView2/0/w/1280/h/960/ignore-error/1)

` //饱和度调节0-无色彩 就是黑白图片， 1- 默认效果， >1饱和度加强 //cm.setScale(1,2,1,1); cm.setSaturation(2); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d8992836eb4?imageView2/0/w/1280/h/960/ignore-error/1)

` cm.setSaturation(2); //色调调节 cm.setRotate(0, 45); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26d8d92adaa30?imageView2/0/w/1280/h/960/ignore-error/1)