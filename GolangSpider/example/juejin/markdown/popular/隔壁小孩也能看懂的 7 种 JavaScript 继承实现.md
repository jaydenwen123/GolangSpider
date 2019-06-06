# 隔壁小孩也能看懂的 7 种 JavaScript 继承实现 #

# JavaScript 没有类，原型链来实现继承 #

因为我在学校接触的第一门语言是cpp，是一个静态类型语言，并且实现面向对象直接就有class关键字，而且只讲了面向对象一种设计思想，导致我一直很难理解javascript语言的继承机制。

JavaScript没有”子类“和”父类“的概念，也没有”类“（class）和”实例“（instance）的区分，全靠” **原型链** “（prototype chain）实现继承。

学的时候就很想吐槽，费了这么大的劲去模拟类，那js干嘛不一开始就设计class关键字而是最开始仅将class作为保留字呢？（ES6之后有了class关键字，是原型的语法糖）

当时我一直怀疑，“js没有class是一种设计缺陷吗？”

原来，JavaScript设计之初，设计里面所有的数据类型都是 **对象** （object），最开始，JavaScript只想被设计成一种简易的脚本语言，设计者JavaScript里面都是对象，必须要有一种机制将所有对象联系起来，但如果引入“类”（class）的概念，那么就太“正式”了，增加了上手难度。

要实现继承，但又不想用类，那该怎么办呢？

JavaScript 的设计者Brendan Eich发现，可以像c++和Java语言中使用 **new** 命令生成实例。

于是new命令被引入到JavaScript，用来从原型对象生成一个实例对象。但是JavaScript没有“类”，原型对象该如何表示呢？

这时，他想到c++和java使用new命令时，都会调用“类”的 **构造函数** （constructor），于是他做了个简化设计，在JavaScript中，new命令后面跟的不是类而是构造函数。

用构造函数生成实例对象，有一个缺点就是无法共享属性和方法。

每一个实例对象，都有自己的属性和方法的副本。这不仅无法做到数据共享，也是极大的资源浪费。

考虑到这一点，brendan Eich决定为构造函数设置一个 **prototype属性** 。

这个属性包含一个 **prototype对象** （是的，prototype属性的值是prototype对象），所有的实例对象需要共享的属性和方法，都放在这个对象里面，那些不需要共享的属性和方法，就放在构造函数里。

实例对象一旦创建，将自动引用prototype对象的属性和方法，也就是说，实例对象的属性和方法，分成两种，一种是本地的，另一种是引用的。

由于所有的实例对象共享同一个prototype对象，那么从外界看起来，prototype对象就好像是实例对象的原型，而实例对象则好像"继承"了prototype对象一样。

如果没了解过c++、java或者其他的编程语言，我相信你看完上面这段内容应该会看睡着了吧！好的，我们还是直接来看看代码吧~

# 原型链继承 #

` //原型链继承 // 父类 // 拥有属性 name function parents ( ) { this.name = "JoseyDong" ; } // 在父类的原型对象上添加一个getName方法 parents.prototype.getName = function ( ) { console.log( this.name); } //子类 function child ( ) { } //子类的原型对象 指向 父类的实例对象 child.prototype = new parents() // 创建一个子类的实例对象，如果它有父类的属性和方法，那么就证明继承实现了 let child1 = new child(); child1.getName(); // => JoseyDong 复制代码`

在只有一个 子类实例对象的时候，我们貌似看不出什么问题。然而在实际场景中，我们会创建很多实例对象来继承父类，毕竟继承得越多，被复写的代码量就越多嘛~

` //原型链继承 // 父类 // 拥有属性 name function parents ( ) { this.name = [ "JoseyDong" ]; } // 在父类的原型对象上添加一个getName方法 parents.prototype.getName = function ( ) { console.log( this.name); } //子类 function child ( ) { } //子类的原型对象 指向 父类的实例对象 child.prototype = new parents() // 创建一个子类的实例对象，如果它有父类的属性和方法，那么就证明继承实现了 let child1 = new child(); child1.getName(); // => ["JoseyDong"] // 创建一个子类的实例对象，在child1修改name前实现继承 let child2 = new child(); // 修改子类的实例对象child1的name属性 child1.name.push( "xixi" ); // 创建子类的另一个实例对象，在child1修改name后实现继承 let child3 = new child(); child1.getName(); // => ["JoseyDong", "xixi"] child2.getName(); // => ["JoseyDong", "xixi"] child3.getName(); // => ["JoseyDong", "xixi"] 复制代码`

当很多时候，我们的实例对象里的值是会虽具体场景而改变的。比如这个时候，我们的child1除了joseydong以外，她的朋友又给她取了个新名字xixi，我们改变了child1的name值。而child1、child2、child3是三个独立的个体，但是最后发现三个孩子都有了新名字！

这就表示，原型链继承里面，使用的都是同一个内存里的值，这样修改该内存里的值，其他继承的子类实例里的值都会变化。

这可不是我们想要的效果，毕竟只有child1被赋予了新名字。并且，如果我想通过子类实例对象传递参数给父类，也是做不到的。

# 借用构造函数 #

` // 构造函数继承 function parents ( ) { this.name = [ "JoseyDong" ]; } // 在子类中，使用call方法构造函数，实现继承 function child ( ) { parents.call( this ); } let child1 = new child(); let child2 = new child(); child1.name.push( "xixi" ); let child3 = new child(); console.log(child1.name); // => ["JoseyDong", "xixi"] console.log(child2.name); // => ["JoseyDong"] console.log(child3.name); // => ["JoseyDong"] 复制代码`

我们使用构造函数的方法，就只修改了child1的名字，而child2和child3的name属性并没有受影响~

同时，由于call()支持传递参数，我们也可以在child中向parent传参啦~

` // 构造函数实现继承 //子类向父类传参 function parents ( name ) { this.name = name; } //call方法支持传递参数 function child ( name ) { parents.call( this ,name) } let child1 = new child( "I am child1" ); let child2 = new child( "I am child2" ); console.log(child1.name); // => I am child1 console.log(child2.name); // => I am child2 复制代码`

好了，现在我们通过构造函数实现继承弥补了用原型链实现继承的缺点，同时也是通过构造函数实现继承的优点：

1.避免了引用类型的属性被所有实例共享

2.可以在child中向parent传参

但是，这种方式也有缺点，因为方法都在构造函数中定义，每次创建实例都会创建一遍方法。

# 组合继承 #

我们发现，通过原型链实现的继承，都是复用同一个属性和方法；通过构造函数实现的继承，都是独立的属性和方法。于是我们大打算利用这一点，将两种方式组合起来：通过在原型上定义方法实现对 **函数的复用** ，通过构造函数的方式保证每个实例都有它 **自己的属性** 。

下面我再举个栗子，让大家感受下组合继承的好处~

` //组合继承 // 偶像练习生大赛开始报名了 // 初赛，我们找了一类练习生 // 这类练习生都有名字这个属性，但名字的值不同，并且都有爱好，而爱好是相同的 // 只有会唱跳rap的练习生才可进入初赛 function student ( name ) { this.name = name; this.hobbies = [ "sing" , "dance" , "rap" ]; } // 我们在student那类里面找到更特殊的一类进入复赛 // 当然，我们已经知道初赛时有了name属性了，而不同练习生名字的值不同，所以使用构造函数方法继承 // 同时，我们想再让练习生们再介绍下自己的年龄，每个子类还可以自己新增属性 // 当然啦，具体的名字年龄就由每个练习生实例来定 // 类只告诉你，有这个属性 function greatStudent ( name,age ) { student.call( this ,name); this.age = age; } // 而大家的爱好值都相同，这个时候用原型链继承就好啦 // 每个对象都有构造函数，原型对象也是对象，也有构造函数，这里简单的把构造函数理解为谁的构造函数就要指向谁 // 第一句将子类的原型对象指向父类的实例对象时，同时也把子类的构造函数指向了父类 // 我们需要手动的将子类原型对象的构造函数指回子类 greatStudent.prototype = new student(); greatStudent.prototype.constructor = greatStudent; // 决赛 kunkun和假kunkun进入了决赛 let kunkun = new greatStudent( 'kunkun' , '18' ); let fakekun = new greatStudent( 'fakekun' , '28' ); // 有请两位选手介绍下自己的属性值 console.log(kunkun.name,kunkun.age,kunkun.hobbies) // => kunkun 18 ["sing", "dance", "rap"] console.log(fakekun.name,fakekun.age,fakekun.hobbies) // => fakekunkun 28 ["sing", "dance", "rap"] // 这个时候，kunkun选手说自己还有个隐藏技能是打篮球 kunkun.hobbies.push( "basketball" ); console.log(kunkun.name,kunkun.age,kunkun.hobbies) // => kunkun 18 ["sing", "dance", "rap", "basketball"] console.log(fakekun.name,fakekun.age,fakekun.hobbies) // => fakekun 28 ["sing", "dance", "rap"] // 我们可以看到，假kunkun并没有抄袭到kunkun的打篮球技能 // 并且如果这个时候新来一位选手，从初赛复赛闯进来的一匹黑马 // 可以看到黑马并没有学习到kunkun的隐藏技能 let heima = new greatStudent( 'heima' , '20' ) console.log(heima.name,heima.age,heima.hobbies) // => heima 20 ["sing", "dance", "rap"] 复制代码`

可以看到，组合继承避开了原型链继承和构造函数继承的缺点，结合了两者的优点，成为了javascript中最常用的继承方式。

# 原型式继承 #

这种继承的思想是将传入的对象作为创建的对象的原型。

` function createObj ( o ) { function F ( ) {}; F.prototype = o; return new F(); } 复制代码`

我们来实现下原型式继承，看看会不会有什么问题

` // 原型式继承 function createObj ( o ) { function F ( ) {}; F.prototype = o; return new F(); } let person = { name : 'JoseyDong' , hobbies :[ 'sing' , 'dance' , 'rap' ] } let person1 = createObj(person); let person2 = createObj(person); console.log(person1.name,person1.hobbies) // => JoseyDong ["sing", "dance", "rap"] console.log(person2.name,person2.hobbies) // => JoseyDong ["sing", "dance", "rap"] person1.name = "xixi" ; person1.hobbies.push( "basketball" ); console.log(person1.name,person1.hobbies) //xixi ["sing", "dance", "rap", "basketball"] console.log(person2.name,person2.hobbies) //JoseyDong ["sing", "dance", "rap", "basketball"] 复制代码`

这个时候我们发现，修改了person1的hobbies的值，person2的hobbies的值也变了。

这是因为包含引用类型的属性值始终会共享相应的值，这点跟原型链继承一样~

而修改了person1.name的值，person2.name的值并未发生改变，并不是因为person1和person2有独立的name值，而是因为person1.name = "xixi"这条语句是给person1实例对象添加了一个name属性，而它的原型对象上name值并没有被修改，所以person2的name没有变化。因为我们找对象上的属性时，总是先找实例对象，没有找到的话再找原型对象上的属性。实例对象和原型对象上如果有同名属性，总是先取实例对象上的值。

ESMAScript5新增了Object.create()方法规范化了原型式继承~

# 寄生式继承 #

创建一个仅用于封装继承过程的函数，该函数在内部以某种形式来做增强对象，最后返回对象。

` //寄生式继承 function createObj ( o ) { let clone = Object.create(o); clone.sayName = function ( ) { console.log( 'hi' ); } return clone } let person = { name : "JoseyDong" , hobbies :[ "sing" , "dance" , "rap" ] } let anotherPerson = createObj(person); anotherPerson.sayName(); // => hi 复制代码`

当然，用寄生式继承来为对象添加函数，和借用构造函数模式一样，每次创建对象都会创建一遍方法。

# 寄生组合式继承 #

前面我们说了，组合继承是javascript最常用的继承模式。这里我们先来回顾下组合式继承的代码：

` //组合继承 function student ( name ) { this.name = name; this.hobbies = [ "sing" , "dance" , "rap" ]; } function greatStudent ( name,age ) { student.call( this ,name); this.age = age; } greatStudent.prototype = new student(); greatStudent.prototype.constructor = greatStudent; let kunkun = new greatStudent( 'kunkun' , '18' ); 复制代码`

组合继承最大的缺点是最调用两次父构造函数

一次是设置子类实例的原型的时候：

` greatStudent.prototype = new student(); 复制代码`

一次是在创建子类型实例的时候：

` let kunkun = new greatStudent( 'kunkun' , '18' ); 复制代码`

在这个例子中，如果我们打印一下kunkun这个对象，我们就会发现greatStudent.prototype和kunkun都有一个属性为hobbies。

![](https://user-gold-cdn.xitu.io/2019/5/27/16af7149f9b69d79?imageView2/0/w/1280/h/960/ignore-error/1)

这其实就是实例对象和原型对象上的属性值重复了，而再找属性值的时候，在实例对象上找到了属性值就不会在原型对象上找了，而这部分原型对象上的值就实打实的浪费了存储空间。

那么我们该如何精益求精，避免这一次重复调用呢？

如果我们不使用greatStudent.prototype = new student()，而是直接让greatStudent.prototype访问到student.prototype呢？

看看如何实现：

` // 寄生组合式继承 function student ( name ) { this.name = name; this.hobbies = [ "sing" , "dance" , "rap" ]; } function greatStudent ( name,age ) { student.call( this ,name); this.age = age; } //关键的三步 实现继承 // 使用F空函数当子类和父类的媒介 是为了防止修改子类的原型对象影响到父类的原型对象 let F = function ( ) {}; F.prototype = student.prototype; greatStudent.prototype = new F(); let kunkun = new greatStudent( 'kunkun' , '18' ); console.log(kunkun); 复制代码`

打印结果：

![](https://user-gold-cdn.xitu.io/2019/5/27/16af7155c693add6?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到，kunkun实例的原型对象上不再有hobbies属性了。

最后，我们封装下这个继承方法：

` function object ( o ) { function F ( ) {} F.prototype = o; return new F(); } function prototype ( child, parent ) { let prototype = object(parent.prototype); prototype.constructor = child; child.prototype = prototype; } // 当我们使用的时候： prototype(Child, Parent); 复制代码`

引用《JavaScript高级程序设计》中对寄生组合式继承的夸赞就是：

这种方式的高效率体现它只调用了一次 Parent 构造函数，并且因此避免了在 Parent.prototype 上面创建不必要的、多余的属性。与此同时，原型链还能保持不变；因此，还能够正常使用 instanceof 和 isPrototypeOf。开发人员普遍认为寄生组合式继承是引用类型最理想的继承范式。

总而言之就是，这种js实现继承的方式是最佳的。

# ES6实现继承 #

然而，ES6之后通过extends关键字实现了继承。

` // ES6 class parents { constructor (){ this.grandmather = 'rose' ; this.grandfather = 'jack' ; } } class children extends parents{ constructor(mather,father){ //super 关键字，它在这里表示父类的构造函数，用来新建父类的 this 对象。 super(); this.mather = mather; this.father = father; } } let child = new children( 'mama' , 'baba' ); console.log(child) // => // father: "baba" // grandfather: "jack" // grandmather: "rose" // mather: "mama" 复制代码`

子类必须在 constructor 方法中调用 super方法，否则新建实例时会报错。这是因为子类没有自己的this 对象，而是继承父类的 this 对象，然后对其进行加工。

只有调用 super 之后，才可以使用 this 关键字，否则会报错。这是因为子类实例的构建，是基于对父类实例加工，只有 super 方法才能返回父类实例。

ES5 的继承实质是先创造子类的实例对象 this，然后再将父类的方法添加到 this 上面（Parent.call(this)）。

ES6 的继承机制实质是先创造父类的实例对象 this （所以必须先调用 super() 方法），然后再用子类的构造函数修改 this。

es6实现继承的核心代码如下：

` function _inherits(subType, superType) { subType.prototype = Object.create(superType && superType.prototype, { constructor: { value: subType, enumerable: false , writable: true , configurable: true } }); if (superType) { Object.setPrototypeOf ? Object.setPrototypeOf(subType, superType) : subType.__proto__ = superType; } } 复制代码`

子类的 **proto** 属性：表示构造函数的继承，总是指向父类。 子类 prototype 属性的 **proto** 属性：表示方法的继承，总是指向父类的 prototype 属性。

除此之外，ES6 可以自定义原生数据结构（比如Array、String等）的子类，这是 ES5 无法做到的。