# git commit 规范校验配置和版本发布配置 #

## 一、 快速配置和版本发布流程 ##

该章节主要是对下文内容的归纳方便往后的查阅，如果需要了解细节部分请从第二章节开始阅读

## 1.1 依赖包安装 ##

` # husky 包安装 npm install husky --save-dev # commitlint 所需包安装 npm install @commitlint/config-angular @commitlint/cli --save-dev # commitizen 包安装 npm install commitizen --save-dev npm install commitizen -g # standard-version 包安装 npm install standard-version --save-dev 复制代码`

## 1.2 配置 commitlint 和 commitizen ##

` # 生成 commitlint 配置文件 echo "module.exports = {extends: ['@commitlint/config-angular']};" > commitlint.config.js # commitizen 初始化 commitizen init cz-conventional-changelog --save-dev --save-exact 复制代码`

## 1.3 更新 package.json ##

` { "scripts": { + "commit": "git-cz", + "release": "standard-version" }, + "husky": { + "hooks": { + "commit-msg": "commitlint -E HUSKY_GIT_PARAMS" + } + } } 复制代码`

## 1.4 commit 方式 ##

* 全局安装 commitizen 情况下可使用 ` git cz` 或者 ` npm run commit` 来提交代码
* 未全局安装 commitizen 情况下可使用 ` npm run commit` 来提交代码

## 1.5 版本发布流程 ##

` git checkout master git pull origin master git fetch origin --prune # 1.0.0 表示当前要发布的版本 npm run release -- --release-as 1.0.0 git push --follow-tags origin master 复制代码`

## 1.6 为远程仓库添加 releases ##

![打开releases](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea645ac89374?imageView2/0/w/1280/h/960/ignore-error/1)

![创建releases](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea6b43c79348?imageView2/0/w/1280/h/960/ignore-error/1)

![编辑releases](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea705081bfd3?imageView2/0/w/1280/h/960/ignore-error/1)

## 二、 使用 husky + commitlint 校验 commit 信息 ##

### 2.1 npm 包 husky 简介 ###

[husky]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftypicode%2Fhusky%2Fblob%2Fmaster%2FDOCS.md ) 主用功能是为 git 添加 [git 钩子]( https://link.juejin.im?target=https%3A%2F%2Fgit-scm.com%2Fdocs%2Fgithooks ) ，它允许我们在使用 git 中在一些重要动作发生时触发自定义脚本(npm script), 比如我们可以在 git push 之前执行特定的自定义脚本对代码进行单元测试、又或者在 git commit 之前执行 eslint 校验，当然本文主要介绍如何借用 husky 为 git 添加 commit-msg 钩子并对 commit 进行校验。

* npm 包 husky 安装

` npm install husky --save-dev 复制代码`

* 更新 package.json

* 添加 pre-commit 钩子，在 git commit 之前前将会在终端输出 ` 我要提交代码啦`
* 添加 commit-msg 钩子，在 git 校验 commit 时会在终端输出 git 所有参数和输入， husky 通过环境变量 HUSKY_GIT_PARAMS HUSKY_GIT_STDIN 是 husky 返回 git 参数和输入
* 添加 pre-push 钩子， 在 git push 之前将在终端输出 ` 提交代码前需要先进行单元测试` 并执行 ` npm test`
` { + "husky": { + "hooks": { + "pre-commit": "echo 我要提交代码啦", + "commit-msg": "echo $HUSKY_GIT_PARAMS $HUSKY_GIT_STDIN", + "pre-push": "echo 提交代码前需要先进行单元测试 && npn test" + } + } } 复制代码`

### 2.2 npm 包 commitlint 简介 ###

commitlint 用于检查您的提交消息是否符合规定提交格式，一般和 husky 包一起使用，用于对 git commit 信息的格式进行校验，当 commit 信息不符合规定格式的情况下将会抛出错误。

#### 2.2.1 commitlint 默认格式 ####

[commitlint]( https://link.juejin.im?target=https%3A%2F%2Fcommitlint.js.org%2F%23%2F ) 的默认格式为：

` # 注意：冒号前面是需要一个空格的, 带 ？ 表示非必填信息 type(scope?): subject body? footer? 复制代码`

* scope 指 commit 的范围（哪些模块进行了修改）
* subject 指 commit 的简短描述
* body 指 commit 主体内容（长描述）
* footer 指 commit footer 信息
* type 指当前 commit 类型，一般有下面几种可选类型：

+----------+--------------------------------------------------------------------------------------+
|   类型   |                                         描述                                         |
+----------+--------------------------------------------------------------------------------------+
| build    | 主要目的是修改项目构建系统(例如                                                      |
|          | glup，webpack，rollup                                                                |
|          | 的配置等)的提交                                                                      |
| ci       | 主要目的是修改项目继续集成流程(例如                                                  |
|          | Travis，Jenkins，GitLab                                                              |
|          | CI，Circle等)的提交                                                                  |
| docs     | 文档更新                                                                             |
| feat     | 新增功能                                                                             |
| merge    | 分支合并 Merge branch ? of ?                                                         |
| fix      | bug 修复                                                                             |
| perf     | 性能, 体验优化                                                                       |
| refactor | 重构代码(既没有新增功能，也没有修复                                                  |
|          | bug)                                                                                 |
| style    | 不影响程序逻辑的代码修改(修改空白字符，格式缩进，补全缺失的分号等，没有改变代码逻辑) |
| test     | 新增测试用例或是更新现有测试                                                         |
| revert   | 回滚某个更早之前的提交                                                               |
| chore    | 不属于以上类型的其他类型                                                             |
+----------+--------------------------------------------------------------------------------------+

#### 2.2.2 配合 husky 包进行使用 ####

* 安装依赖包

` npm install --save-dev @commitlint/config-angular @commitlint/cli 复制代码`

* 修改 package.json: 在 husky 配置下使用 commitlint 配置 commit-msg 钩子脚本

` { "husky": { "hooks": { + "commit-msg": "commitlint -E HUSKY_GIT_PARAMS" } } } 复制代码`

* 添加 commitlint 配置文件

项目下新增 commitlint.config.js 文件，并针对 commitlint 进行简单配置

**配置说明:** 规则由键值和配置数组组成，如：'name: [0, 'always', 72]'，数组中第一位为 level（等级），可选 0, 1, 2，0 为 disable（禁用），1 为 warning（警告），2 为 error（错误），第二位为该规则是否被应用，可选 always | never， 第三位为该规则允许值。

` module.exports = { // 继承默认配置 extends: [ "@commitlint/config-angular" ], // 自定义规则 rules: { 'type-enum' : [ 2 , 'always' , [ 'upd' , 'feat' , 'fix' , 'refactor' , 'docs' , 'chore' , 'style' , 'revert' , ]], 'header-max-length' : [ 0 , 'always' , 72 ] } }; 复制代码`

* 下面是几种不同 commit 信息的校验情况

![测试](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea7e111a95c5?imageView2/0/w/1280/h/960/ignore-error/1)

## 三、 在提交代码时使用 commitizen 来提供一个交互界面 ##

在上文我们对 git commit 信息添加了校验，但有个缺点就是需要我们手动编辑 commit 信息，这样就显得麻烦了很多，这时我们可以引入 commitizen，在使用 commitizen 提交我们的代码时，会在终端给出一个交互界面，该界面会提示我们在提交时所需要填写的所有字段, 我们只需要在交互界面中按照顺序填写相应信息 commitizen 会自动帮我们合成信息，并发起 commit。

### 3.1 配置 commitizen ###

* 安装 commitizen cli 工具

` npm install commitizen -g 复制代码`

* 初始化当前项目

` # 下面面的命令为你做了下面几件事: # # 安装 cz-conventional-changelog npm 模块 # # 将添加 config.commitizen 配置 commitizen init cz-conventional-changelog --save-dev --save-exact 复制代码`

* 使用 git cz 代替 git commit 提交代码

![测试](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea83031387c0?imageView2/0/w/1280/h/960/ignore-error/1)

### 3.2 通过 npm 脚本来代替 git commit ###

其实到此 commitizen 配置可以算是基本完成，但如果只是配置到这里的话用户在使用 git cz 来提交代码时则需要全局安装 commitizen 否则将会报错，要想用户无障碍使用 commitizen 进行提交代码则可以在当前项目下安装 commitizen 并添加 npm 脚本,用户在未全局安装 commitizen 的情况下就可以通过该 npm 脚本来提交代码。

* 添加 commitizen 依赖

` npm install commitizen -D 复制代码`

* 添加 npm 脚本

` "scripts": { + "commit": "git-cz" }, 复制代码`

* 当用户未全局安装 commitizen 可通过 ` npm run commit` 代替 git commit 来提交代码

![测试](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea87f743cc02?imageView2/0/w/1280/h/960/ignore-error/1)

## 四、 自动生成 change log 并进行版本发布 ##

在使用上文 commit 规范的情况下， 可以通过 standard-version 自动生成 change log，并更新项目的版本信息添加 git tag, change log 中将收录所有 type 为 feat 和 fix 的 commit

### 4.1 standard-version 安装和配置 ###

* 安装 standard-version

` npm i --save-dev standard-version 复制代码`

* 配置 npm 脚本命令

` { "scripts": { + "release": "standard-version" } } 复制代码`

### 4.2 版本发布流程 ###

* 切换至 master 分支

` git checkout master 复制代码`

* 拉取远程分支

` git pull origin master 复制代码`

* 拉取远程信息

` git fetch origin --prune 复制代码`

* 自动生成 changelog 并更新版本为 1.0.0

` # 下面面的命令为你做了下面几件事: # # 修改 package.json 中的版本号 # # 使用 legacy -changelog 更新 CHANGELOG.md # # 提交 package.json 和 CHANGELOG.md # # 添加 tag npm run release -- --release-as 1.0.0 复制代码`

* 更新本地 tag 到远程分支

` git push --follow-tags origin master 复制代码`

### 4.3 为远程仓库添加 releases ###

![打开releases](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea645ac89374?imageView2/0/w/1280/h/960/ignore-error/1)

![创建releases](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea6b43c79348?imageView2/0/w/1280/h/960/ignore-error/1)

![编辑releases](https://user-gold-cdn.xitu.io/2019/6/1/16b0ea705081bfd3?imageView2/0/w/1280/h/960/ignore-error/1)