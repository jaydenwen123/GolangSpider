# ES6的一些小技巧，代替lodash #

Javascript ES6标准实行后，Lodash或者Ramada中的一些功能我们就不需要了，可以使用ES6的语法来实现

## 获取Object中指定键值 ##

我们现在可以使用解包的方法快速获取对象中指定键值的值

` const obj = { a : 1 , b : 2 , c : 3 , d : 4 }; // 获取obj中a与b的值 const {a,b} = obj; // 也可以给他们取别名 const { a :A, b :B} = obj; 复制代码`

这个小技巧非常的方便，也是最基础的使用方法

## 排除Object中不需要的键值 ##

既然我们可以获取到想要的对象键值，那么也可以排除掉不想要的键值，使用方法就要用到ES6的rest新特性

` const obj = { a : 1 , b : 2 , c : 3 , d : 4 } // 我们想要获取除了a之外的所有属性 const {a, ...other} = obj 复制代码`

我们只要指定那些排除掉的属性，剩下的就是需要的属性，这样可以非常快速的排除不需要的属性

## 对象快速求和 ##

有时候我们需要对一组对象数组中的某一个属性求总和，以前我们可以使用 ` forEach` 或者 ` for` 这样的循环遍历的方法来计算，现在我们可以使用 ` reduce` 方法来快速实现

` const objs = [ { name : 'lilei' , score : 98 }, { name : 'hanmeimei' , score : 95 }, { name : 'polo' , score : 85 }, ... ] const scoreTotal = objs.reduce( ( total, obj ) => { return obj.score + total; }, 0 /*第二个参数是total的初始值*/ ) 复制代码`

使用 ` reduce` 就能快速的实现对某一个属性的总和计算

## map也能异步遍历 ##

是不是觉得只有 ` for` 能够进行异步操作不方便，其实 ` map` 也能进行异步操作，不过需要结合 ` Promise` 的新方法一起使用

` const arr = [ 1 , 2 , 3 , 4 ,...] const queue = arr.map( async item => { return item + 1 ; }) Promise.all(queue).then( newArr => console.log(newArr)) 复制代码`

这样一来我们在 ` map` 中也能使用异步操作了