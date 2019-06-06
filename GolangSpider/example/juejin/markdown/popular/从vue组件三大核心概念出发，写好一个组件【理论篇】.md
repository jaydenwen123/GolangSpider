# 从vue组件三大核心概念出发，写好一个组件【理论篇】 #

> 
> 
> 
> 一个适用性良好的组件，一种是可配置项很多，另一种就是容易覆写，从而扩展功能
> 
> 

Vue 组件的 API 来自三部分——prop、事件和插槽：

* prop 允许外部环境传递数据给组件
* event 允许从组件内触发外部环境的副作用
* slot 允许外部环境将额外的内容组合在组件中

## prop ##

> 
> 
> 
> 组件具有自身状态，当没有相关 porps 传入时，使用自身状态完成渲染和交互逻辑；当该组件被调用时，如果有相关 props
> 传入，那么将会交出控制权，由父组件控制其行为
> 
> 

### 仅一个值传入组件 ###

* 如果该组件设计上支持双向绑定，可使用 ` v-model` 将该参数传入组件，减少记忆成本（毕竟 vue 官方的语法糖，不用白不用）

` < my-component v-model = "foo" /> 复制代码`

* 如果该组件可以独立运行，不依赖父组件时，还是给这个值起个名字吧

` < component-no-sync :childNeed = "foo" /> 复制代码`

### 很多值需要传入组件 ###

比如当一个组件有诸多配置项，且当没有传入配置项取用组件内部默认项的时候，我们原先的父组件写法：

` < child-component :prop1 = "var1" :prop2 = "var2" :prop = "var3"... /> 复制代码`

其实可以在父组件上直接使用 ` v-bind={子组件props集合}`

但是为了方便覆写子组件的内部配置项，不妨使用一个对象将配置收集到一起，但是这种做法不能利用 props 验证对象里面每个的值类型

` <child-component v-model= "text" :setting= "{color:'bule'}" /> // 子组件内部读取配置，通过扩展运算符替换掉默认配置 const setting ={ ...defaultSetting, ...this.setting } 复制代码`

### computed 属性 ###

vue 的 computed 属性默认是只读的，你可以提供一个 ` setter` 。它可以优化我写组件的逻辑，适用于父组件处理的值和子组件处理的值是同一个的情况

` < template > < el-select v-model = "email" > < el-option v-for = "item in adminUserOptions" :key = "item.email" :label = "item.email" :value = "item.email" /> </ el-select > </ template > 复制代码` ` export default { props : { value : {} }, computed : { email : { get() { return this.value }, set(val) { this.$emit( 'input' , val) this.$emit( 'change' , val) } } } } 复制代码`

### 灵活的 prop ###

我们常看到一些优秀的组件库，传入的值既可以是一个 String/Number，也可以是一个函数。

比如 ` ElementUI` 的 ` Table` 组件，当你想要显示树形数据的时候，必须传入 ` row-key` 。看它的介绍就知道是有多灵活：

` row-key` 的作用：行数据的 ` Key` ，用来优化 ` Table` 的渲染；在使用 ` reserve-selection` 功能与显示树形数据时，该属性是必填的。类型为 String 时，支持多层访问： ` user.info.id` ，但不支持 user.info[0].id，此种情况请使用 ` Function`

处理 rowKey 生成 RowIdentity 的函数源码：

` //https://github.com/ElemeFE/element/blob/dev/packages/table/src/util.js export const getRowIdentity = ( row, rowKey ) => { if (!row) throw new Error ( 'row is required when get row identity' ) // 行数据的key if ( typeof rowKey === 'string' ) { if (rowKey.indexOf( '.' ) < 0 ) { return row[rowKey] } // 支持多层访问：user.info.id let key = rowKey.split( '.' ) let current = row for ( let i = 0 ; i < key.length; i++) { current = current[key[i]] } return current // 通过函数自定义 // 我处理过父和子id可能相同的情况，只好通过Function自定义 // 不可以通过时间或者随机字符串生成ID } else if ( typeof rowKey === 'function' ) { return rowKey.call( null , row) } } 复制代码`

由于业务场景多变，组件的设计者很难考虑完全，不妨设计灵活的 prop，由开发者自行定义

### 其他 ###

当组件有 prop 传入的时候，尽量考虑一下，当 prop 变化的时候，组件是否能够响应 prop 的变化。也就是说使用 prop，是否使用了计算属性，或者 watch 了 props 的值。

## 事件 ##

### emit/on ###

读者肯定知道 emit/on 如何使用，我就简单说一下 vue 的 ` v-model` 和 ` sync` 的语法糖，我们可以利用这些语法糖，帮助我们写出简洁的代码（父组件可以少写监听子组件的事件，比如你不用写@input）

#### v-model ####

> 
> 
> 
> 看一下下面的代码示例，就能懂这句话了。v-model 会忽略所有表单元素的 value、checked、selected 特性的初始值而总是将
> Vue 实例的数据作为数据来源。你应该通过 JavaScript 在组件的 data 选项中声明初始值
> 
> 

` < input v-model = "searchText" /> < input v-bind:value = "searchText" v-on:input = "searchText = $event.target.value" /> // 当把v-model用在组件上 < custom-input v-bind:value = "searchText" v-on:input = "searchText = $event" > </ custom-input > 复制代码`

为了让它正常工作，这个组件内的 ` <input>` 必须：将其 value 特性绑定到一个名叫 value 的 prop 上在其 input 事件被触发时，将新的值通过自定义的 input 事件抛出，即 ` this.$emit('input',changedValue)`

#### 自定义 v-model ####

为啥要自定义组件的 v-model 呢，因为数据不符合要求呗。你的输入值不可能总是 value ，你的事件不可能总是 input，具体详见 [文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-custom-events.html%23%25E8%2587%25AA%25E5%25AE%259A%25E4%25B9%2589%25E7%25BB%2584%25E4%25BB%25B6%25E7%259A%2584-v-model )

#### sync（双向绑定语法糖） ####

> 
> 
> 
> vue 真的是方便了开发者很多，站在开发者的角度考虑，很大的提升开发效率
> 
> 

以  update:myPropName  的模式触发事件取代双向绑定 ` this.$emit('update:title', newTitle)` ，具体详见 [文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-custom-events.html%23sync-%25E4%25BF%25AE%25E9%25A5%25B0%25E7%25AC%25A6 )

### Function 通过 prop 传入 ###

> 
> 
> 
> 本来想放在 prop 部分的，但是个人觉得其实它和 emit/on 更有关系一点
> 
> 

有读者可能会问，为什么不能把子组件里面的事件 emit 出来，通过父组件处理？然后传入一个控制子组件的 prop 属性。

我想说的是，可以，但是这样真的很麻烦，子组件内部的状态却要依赖父组件传值。

该组件内部的状态，我们需要把它暴露出来嘛？我觉得不需要，组件内部的状态就让它处于组件内部

但是可以通过传入 function（你可以理解为一个钩子），参与组件状态变更的行为。比如很好用的拖拽库, [Vue.Draggable]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FSortableJS%2FVue.Draggable ) 控制元素是否被拖动的行为。

` Vue.Draggable` 可以传入一个 move 方法，我们看一下它如何处理的。

` onDragMove(evt, originalEvent) { const onMove = this.move; // 如果没有传入move，那么返回true，可以移动 if (!onMove || ! this.realList) { return true ; } const relatedContext = this.getRelatedContextFromMoveEvent(evt); const draggedContext = this.context; const futureIndex = this.computeFutureIndex(relatedContext, evt); Object.assign(draggedContext, { futureIndex }); const sendEvt = Object.assign({}, evt, { relatedContext, draggedContext }); // 组件行为由传入的move函数控制 return onMove(sendEvt, originalEvent); } 复制代码`

这样做的好处，就是组件内部自由一套运行逻辑，但是我可以通过传入 function 来干预。我没有直接修改组件内部状态，而是通过函数（你可以称它为钩子）去触发，方便调试组件，使得组件行为具有可预测性

### 父组件直接操作子组件 ###

很少有这样的骚操作，但是由于数据和操作的复杂性，当数据结构复杂，嵌套过深的情况下，父组件很难对于子组件的数据的精细控制

因此，如果不得已而为之，请在文档里，把子组件可以调用的方法暴露出来，供使用者使用。使用这种组件比较麻烦，得去看文档，没有文档的只好去看源码

` ElementUI` 的 ` tree` 组件 ( https://link.juejin.im?target=https%3A%2F%2Felement.eleme.cn%2F%23%2Fzh-CN%2Fcomponent%2Ftree%23fang-fa ) 提供了很多方法，用于父组件去操作子组件。

eg: ` this.$refs.tree.setCheckedKeys([]);`

## 插槽 ##

> 
> 
> 
> HTML ` <slot>` element (
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fslot
> ) 是 Web Components 技术的一部分，是自定义 web 组件的占位符，vue 里面的 slot 的灵感来自 Web Components
> 规范草案，具体见 [文档](
> https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-slots.html
> )
> 
> 

### [默认插槽]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-slots.html%23%25E6%258F%2592%25E6%25A7%25BD%25E5%2586%2585%25E5%25AE%25B9 ) ###

能用默认插槽就不要使用具名插槽，我真的不想使用你这个组件的时候还去翻看你的插槽叫什么名字

之前我司一个网页模板 三个插槽，header，body，footer，我用的是真的难受，每次都记不得，看似三个单词都挺熟悉的，但是其实 head,content,foot 这些单词也都行啊，谁知道用啥（可能我老了吧，组件如果不是必要尽量不要让人有记忆成本）。

### [后备内容]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-slots.html%23%25E5%2590%258E%25E5%25A4%2587%25E5%2586%2585%25E5%25AE%25B9 ) ###

就是给组件里面的插槽定义默认值，它只会在没有提供内容的时候被渲染。建议用上插槽就给它添加默认内容

## 封装他人组件 ##

有些时候我们可能是对他人的组件进行封装，这里强烈推荐使用 ` v-bind="$attrs" 和 v-on="$listeners"` 。 ` vm.$attrs` 是一个属性，其包含了父作用域中不作为 prop 被识别 (且获取) 的特性绑定 (class 和 style 除外)。这些未识别的属性可以通过 ` v-bind="$attrs"` 传入内部组件。未识别的事件可通过 ` v-on="$listeners"` 传入

举个例子，比如我创建了我的按钮组件 ` myButton` ，封装了 element-ui 的 el-button 组件（其实什么事情都没做），在使用组件 ` <my-button />` 时，就可以直接在组件上使用 el-button 的属性,不被 prop 识别的属性会传入到 el-button 元素上去

` <template> <div> <el-button v-bind="$attrs">导出</el-button> <div> </template> // 父组件使用 <my-button type='primary' size='mini'/> 复制代码`

## 组件命名 ##

> 
> 
> 
> 这里推荐遵循 [vue 官方指南](
> https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fstyle-guide%2F
> ) ，值得一看
> 
> 

我们构建组件的时候通常会将其入口命名为 index.vue ，引入的时候，直接引入该组件的文件夹即可。

但是这样做会有一个问题，当你编辑多个组件的时候，所有的组件入口都叫做 ` index.vue` ，容易糊涂

vscode 显然意识到了这个问题，所以当文件名相同的文件被打开时，它会在文件名旁边显示文件夹名

如何解决呢，我们可以把 index.js 当作一个单纯的入口，不承担任何逻辑。仅仅负责引入 ` component-name-container` 以及 ` export default component-name-container`

` my-app └── src └── components └── component-name ├── component-name.css ├── component-name-container.vue └── index.js 复制代码`

## tips（个人喜好） ##

* [template]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fconditional.html%23%25E5%259C%25A8-lt-template-gt-%25E5%2585%2583%25E7%25B4%25A0%25E4%25B8%258A%25E4%25BD%25BF%25E7%2594%25A8-v-if-%25E6%259D%25A1%25E4%25BB%25B6%25E6%25B8%25B2%25E6%259F%2593%25E5%2588%2586%25E7%25BB%2584 ) ,把一个 ` <template>` 元素当做不可见的包裹元素，并在上面使用 v-if。最终的渲染结果将不包含 ` <template>` 元素
* 能用 computed 计算属性的，尽量就不用 watch
* 模板里面写太多 v-if 会让你的模板很难看， ` v-else-if` 尽量还是别用了吧。一长串的 if else，在模板里面看的很乱

## 关于我 ##

一个一年小前端，关注我的微信公众号，和我一起交流，我会尽我所能，并且看看我能成长成什么样子吧。

![微信公共号](https://user-gold-cdn.xitu.io/2019/5/16/16ac0e7b1f78bede?imageView2/0/w/1280/h/960/ignore-error/1)

你有什么写组件的独特技巧，不妨在评论区告诉我吧