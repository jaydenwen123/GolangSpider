# SpringBoot踩坑日记-一个非空校验引发的bug #

## 首先先给出mini版项目 ##

` @Data public class User { @NotNull @Size(min = 1) private List<String> strings; } @RequestMapping( "" ) public User hello(@Validated @RequestBody User user) { user.setStrings(user.getStrings() .stream() .map(String::toUpperCase) .collect(Collectors.toList())); return user; } 复制代码`

User类为数据类，里面有一个list存放String，业务逻辑就是将User类中的String由小写转为大写，返回给前台，很简单吧。

之前的业务代码里只有strings不为空，且长度大于1就可以通过校验了。

但是这样是否是万无一失呢？一开始我也是这么认为的，直到有一天前台发来了这么一个对象

` "strings" :[ "abc" ,null, "def" ] 复制代码`

然后我这儿报了个NullPoint异常。所以这个地方还漏了一个校验，就是集合里的对象，也都不能为空！

集合里的对象不能为空该怎么办呢？如果不想把这部分校验放在业务代码里的话，解决思路有两个:

` 1.在转换的时候，过滤掉null。 2.在校验的时候，校验集合里的对象为非空。 复制代码`

## 1.在转换时候过滤 ##

首先看看如果转换的时候需要过滤该怎么办？直接上代码:

` @Configuration public class GlobalConfiguration { @Bean //向spring容器中注入fastjson消息转换器(我比较喜欢用这个，你们随意) public HttpMessageConverters message (){ FastJsonConfig fastJsonConfig = new FastJsonConfig(); fastJsonConfig.setSerializerFeatures(SerializerFeature.PrettyFormat); FastJsonHttpMessageConverter fastJson = new FastJsonHttpMessageConverter(); fastJson.setFastJsonConfig(fastJsonConfig); return new HttpMessageConverters(fastJson); } } //自定义fastjson反序列化组件 public class ListNotNullDeserializer implements ObjectDeserializer { @Override @SuppressWarnings({ "unchecked" , "rawtypes" }) public <T> T deserialze(DefaultJSONParser parser, Type type , Object fieldName) { if (parser.lexer.token() == JSONToken.NULL) { parser.lexer.nextToken(JSONToken.COMMA); return null; } Collection list = TypeUtils.createCollection( type ); Type itemType = TypeUtils.getCollectionItemType( type ); parser.parseArray(itemType, list, fieldName); //在返回的时候，过滤掉集合里的null对象 return (T) list.stream().filter(Objects::nonNull).collect(Collectors.toList()); } @Override public int getFastMatchToken () { return 0; } } //在需要过滤的集合上面添加相应的转换器标志 @JSONField(deserializeUsing = ListNotNullDeserializer.class) private List<String> strings; 复制代码`

这样的话，我们拿到的list就是已经过滤掉null对象的list了。这个时候在对整个集合进行非空校验和长度校验，就可以确保没问题。

但是讲道理的话，前台传这种数据过来，多半是他们自己的逻辑出现了问题，我们就这样把问题吃了也不是特别合适。所以得通过校验告诉他们哪里有问题。

## 自定义校验器 ##

在springboot中，我并没有找到可以校验集合里所有对象为非空的注解。所以看来得自己实现一个。

` @Retention(RetentionPolicy.RUNTIME) @Target({ElementType.FIELD}) @Inherited @Documented @Constraint(validatedBy = NotNullEleValidator.class) public @interface NotNullElement { String message() default " 集合中有元素为null !" ; Class<?>[] groups() default {}; Class<? extends Payload>[] payload() default {}; } public class NotNullEleValidator implements ConstraintValidator<NotNullElement, List<?>> { @Override public boolean isValid(List<?> value, ConstraintValidatorContext context) { if (value == null) { return true ; } for (Object o : value) { if (o == null) { return false ; } } return true ; } } 复制代码`

自定义校验器规则也很简单校验通过返回true，不通过返回false，这样的话springboot就可以获取到校验信息，来决定是否抛出异常。

需要注意的是，校验器的加载由spring框架完成，也就是校验器可以使用spring容器中的类。这个功能也很有用。

## **返回目录** ( https://juejin.im/post/5c8a4458f265da2da23d703c ) ##