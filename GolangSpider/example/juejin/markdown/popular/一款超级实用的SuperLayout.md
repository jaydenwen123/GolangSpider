# 一款超级实用的SuperLayout #

## 前言 ##

项目中会经常用到横向的图文布局，比如下面这些：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26147bba2a12f?imageView2/0/w/1280/h/960/ignore-error/1)

还有这样的布局，经常出现在设置性的界面：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26147bbb5944a?imageView2/0/w/1280/h/960/ignore-error/1)

这些布局虽然不难，但是长相类似，频繁出现，每次都要手写这么多代码还是很累的。完全可以自定义一个布局，兼容这些常见的场景，通过对外暴漏一些属性来设置里面的内容。

## 实现 ##

于是我自定义了SuperLayout，实现了常见的场景。使用时只需要写一个布局，设置一些属性即可实现效果，能将原先手写的布局代码减少80%左右，大大提高布局的编写效率。

在xml中这样使用：

` < com.lxj.androidktx.widget.SuperLayout android:layout_marginLeft = "20dp" android:layout_marginRight = "20dp" android:paddingLeft = "14dp" android:paddingRight = "15dp" android:paddingTop = "10dp" android:paddingBottom = "10dp" android:layout_marginTop = "15dp" app:sl_leftImageSrc = "@drawable/avatar" app:sl_leftText = "头像" app:sl_leftTextColor = "#222" app:sl_leftTextSize = "18sp" app:sl_rightImageSrc = "@mipmap/jt" app:sl_solid = "#3EDCE9" app:sl_corner = "15dp" app:sl_strokeWidth = "2dp" app:sl_stroke = "#f00" android:layout_width = "match_parent" android:layout_height = "wrap_content" /> 复制代码`

所有可以设置的属性如下：

` < declare-styleable name = "SuperLayout" > < attr name = "sl_leftImageSrc" format = "reference" /> < attr name = "sl_leftImageSize" format = "dimension" /> < attr name = "sl_leftText" format = "string" /> < attr name = "sl_leftTextColor" format = "color" /> < attr name = "sl_leftTextSize" format = "dimension" /> < attr name = "sl_leftTextMarginLeft" format = "dimension" /> < attr name = "sl_leftTextMarginTop" format = "dimension" /> < attr name = "sl_leftTextMarginRight" format = "dimension" /> < attr name = "sl_leftTextMarginBottom" format = "dimension" /> < attr name = "sl_leftSubText" format = "string" /> < attr name = "sl_leftSubTextColor" format = "color" /> < attr name = "sl_leftSubTextSize" format = "dimension" /> < attr name = "sl_centerText" format = "string" /> < attr name = "sl_centerTextColor" format = "color" /> < attr name = "sl_centerTextGravity" format = "integer" /> < attr name = "sl_centerTextBg" format = "reference" /> < attr name = "sl_centerTextSize" format = "dimension" /> < attr name = "sl_rightText" format = "string" /> < attr name = "sl_rightTextColor" format = "color" /> < attr name = "sl_rightTextSize" format = "dimension" /> < attr name = "sl_rightTextBg" format = "reference" /> < attr name = "sl_rightTextBgColor" format = "color" /> < attr name = "sl_rightTextVerticalPadding" format = "dimension" /> < attr name = "sl_rightTextHorizontalPadding" format = "dimension" /> < attr name = "sl_rightImageSrc" format = "reference" /> < attr name = "sl_rightImageMarginLeft" format = "dimension" /> < attr name = "sl_rightImageSize" format = "dimension" /> < attr name = "sl_rightImage2Src" format = "reference" /> < attr name = "sl_rightImage2MarginLeft" format = "dimension" /> < attr name = "sl_rightImage2Size" format = "dimension" /> < attr name = "sl_solid" format = "color" /> < attr name = "sl_corner" format = "dimension" /> < attr name = "sl_stroke" format = "color" /> < attr name = "sl_strokeWidth" format = "dimension" /> < attr name = "sl_topLineColor" format = "color" /> < attr name = "sl_bottomLineColor" format = "color" /> < attr name = "sl_enableRipple" format = "boolean" /> < attr name = "sl_rippleColor" format = "color" /> </ declare-styleable > 复制代码`

## 添加依赖 ##

在使用之前需要先添加依赖，这个类放在了 ` AndroidKTX` 类库中，需要依赖一下：

` implementation 'com.lxj:androidktx:1.2.0' //for androidx implementation 'com.lxj:androidktx:1.2.0-x' 复制代码`

AndroidKTX包含了一些列非常有用的Kotlin扩展和通用控件，代码不多，个个都很实用。如果你用Kotlin开发，一定不要错过。它的地址是： [github.com/li-xiaojun/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fli-xiaojun%2FAndroidKTX )

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26147bbb89a5e?imageView2/0/w/1280/h/960/ignore-error/1)