# 跨端兼容 npm script #

本人用的是 Mac，所以上面几章 npm script 都顺利运行（前半生一路开挂），但是还是有使用 Windows 系统的开发小伙伴（多句嘴，经济允许还是建议购置一台 Mac，懂得投资自己）。

那下面就说说在 Windows 平台下，我们的 npm script 还如何顺利地运行。

> 
> 
> 
> 注：Windows 自带的 cmd 是个不靠谱的家伙，建议使用 git bash 来代替 cmd 运行 npm script，别说我没告诉过你。
> 
> 

## 常用命令的替代者 ##

前面涉及到文件操作的命令有 **文件和目录创建、删除和复制等** 操作，npm 社区对于这些也提供了跨平台的兼容包，一看来看看吧

* 目录新增 [make-dir-cli]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fmake-dir-cli ) ，同等 ` mkdir -p` ;
* 文件或目录的删除 [rimraf]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fisaacs%2Frimraf ) 或 ` del-cli` ，同等 ` rm -rf` ;
* 文件或目录的复制 [cpr]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fcpr ) ，同等 ` cp -r` ;
* 跨平台的变量引用，变量写法统一， [cross-var]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fcross-var ) ，比如 Windows 写法 ` %npm_package_name%` ，Linux 写法 ` $npm_package_name` ;

### 添加依赖包 ###

` npm install make-dir-cli rimraf cpr cross- var -D // 或 yarn add make-dir-cli rimraf cpr cross- var -D 复制代码`

### 命令改写 ###

1.文件 package.json

` // 不兼容 Windows "scripts" : { ..., "//" : "# 不兼容 Windows" , "cover:cleanup" : "# 清理覆盖率报告- \n rm -rf coverage && rm -rf .nyc_output" , "cover" : "# 生成覆盖率报告 \n nyc --reporter=html npm run test" , "cover:archive" : "# 归档覆盖率报告 \n mkdir -p coverage_archive/$npm_package_version && cp -r coverage/* coverage_archive/$npm_package_version" , "cover:server" : "http-server coverage_archive/$npm_package_version -p $npm_package_config_port" , "cover:open" : "open http://localhost:$npm_package_config_port" , "postcover" : "# 执行并打开覆盖率报告 \n npm-run-all cover:archive cover:cleanup --parallel cover:server cover:open" } // 兼容 Windows "scripts" : { ..., "//" : "# 兼容 Windows" , "cover:cleanup" : "# 清理覆盖率报告- \n rimraf coverage && rimraf .nyc_output" , "cover" : "# 生成覆盖率报告 \n nyc --reporter=html npm run test" , "cover:archive" : "# 归档覆盖率报告 \n cross-var \"make-dir coverage_archive/$npm_package_version && cpr coverage coverage_archive/$npm_package_version -o\"" , "cover:server" : "cross-var http-server coverage_archive/$npm_package_version -p $npm_package_config_port" , "cover:open" : "cross-var open http://localhost:$npm_package_config_port" , "precover" : "npm run cover:cleanup" , "postcover" : "# 执行并打开覆盖率报告 \n npm-run-all cover:archive --parallel cover:server cover:open" } 复制代码`

2.执行

` npm run cover 复制代码`

### 命令剖析 ###

* 文件或目录的复制，Linux 写法是 ` cp -r coverage/* coverage_archive/$npm_package_version` ，Windows 写法 ` cpr coverage coverage_archive/$npm_package_version -o` 。细节有两点注意：第一是参数位置，Linux 平台 ` -p` 紧跟在 ` cp` 后面，Windows 平台 ` -o` 在 ` cpr` 末尾；第二是复制内容路径写法，Linux 平台 ` coverage/*` ，Windows 平台 ` coverage` ;
* 把 ` cover:cleanup` 从 ` postcover` 移出来放入 ` precover` 里执行，避免 ` cpr` 没归档完就被清空的情况；
* 使用了变量的命令，整个命令用双引号包起来，双引号前记得加 ` \` 转义，然后开头记得加上 ` cross-var` ；

## 环境变量跨平台设置 ##

开门见山，使用 [cross-env]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fcross-env ) 来实现。

### 添加依赖包 ###

` npm i cross-env -D // 或 yarn add cross-env -D 复制代码`

### 命令改写 ###

` "scripts" : { "test" : "cross-env NODE_ENV=test mocha test/" } 复制代码`

## You can ##

[上一章：环境变量的使用于 npm script]( https://juejin.im/post/5cee32476fb9a07ee4634596 )

[下一章：命令补全的实现于 npm script]( https://juejin.im/post/5cf47f105188253a2b01c981 )