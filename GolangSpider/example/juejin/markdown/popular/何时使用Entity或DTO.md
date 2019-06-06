# 何时使用Entity或DTO #

关注公众号： 锅外的大佬

每日推送国外优秀的技术翻译文章，励志帮助国内的开发者更好地成长！

` JPA` 和 ` Hibernate` 允许你在 ` JPQL` 和 ` Criteria` 查询中使用 ` DTO` 和 ` Entity` 作为映射。当我在我的在线培训或研讨会上讨论 ` Hibernate` 性能时，我经常被问到，选择使用适当的映射是否是重要的？ 答案是：是的！为你的用例选择正确的映射会对性能产生巨大影响。我只选择你需要的数据。很明显，选择不必要的信息不会为你带来任何性能优势。

## 1.DTO与Entity之间的主要区别 ##

` Entity` 和 ` DTO` 之间常被忽略的区别是—— ` Entity` 被持久上下文(persistence context)所管理。当你想要更新 ` Entity` 时，只需要调用 ` setter` 方法设置新值。 ` Hibernate` 将处理所需的SQL语句并将更改写入数据库。

天下没有免费的午餐。 ` Hibernate` 必须对所有托管实体(managed entities)执行脏检查(dirty checks)，以确定是否需要在数据库中保存变更。这很耗时，当你只想向客户端发送少量信息时，这完全没有必要。

你还需要记住， ` Hibernate` 和任何其他 ` JPA` 实现都将所有托管实体存储在一级缓存中。这似乎是一件好事。它可以防止执行重复查询，这是Hibernate写入优化所必需的。但是，需要时间来管理一级缓存，如果查询数百或数千个实体，甚至可能发生问题。

使用 ` Entity` 会产生开销，而你可以在使用 ` DTO` 时避免这种开销。但这是否意味着不应该使用 ` Entity` ？显然不是。

## 2.写操作投影 ##

实体投影(Entity Projections)适用于所有写操作。 ` Hibernate` 以及其他 ` JPA` 实现管理实体的状态，并创建所需的SQL语句以在数据库中保存更改。这使得大多数创建，更新和删除操作的实现变得非常简单和有效。

` EntityManager em = emf.createEntityManager(); em.getTransaction().begin(); Author a = em.find(Author.class, 1L ); a.setFirstName( "Thorben" ); em.getTransaction().commit(); em.close(); 复制代码`

## 3.读操作投影 ##

但是只读(read-only)操作要用不同方式处理。如果想从数据库中读取数据，那么 ` Hibernate` 就不会管理状态或执行脏检查。 因此，从理论上说，对于读取数据， ` DTO` 投影是更好的选择。但真的有什么不同吗？我做了一个小的性能测试来回答这个问题。

### 3.1.测试设置 ###

我使用以下领域模型进行测试。它由 ` Author` 和 ` Book` 实体组成，使用多对一关联（many-to-one）。所以，每本书都是由一位作者撰写。

` @Entity public class Author { @Id @GeneratedValue (strategy = GenerationType.AUTO) @Column (name = "id" , updatable = false , nullable = false ) private Long id; @Version private int version; private String firstName; private String lastName; @OneToMany (mappedBy = "author" ) private List bookList = new ArrayList();     ... } 复制代码`

要确保 ` Hibernate` 不获取任何额外的数据，我设置了 ` @ManyToOne` 的 ` FetchType` 为 ` LAZH` 。你可以阅读 [ Introduction to JPA FetchTypes]( https://link.juejin.im?target=https%3A%2F%2Fthoughts-on-java.org%2Fentity-mappings-introduction-jpa-fetchtypes%2F ) 获取不同 ` FetchType` 及其效果的更多信息。

` @Entity public class Book { @Id @GeneratedValue (strategy = GenerationType.AUTO) @Column (name = "id" , updatable = false , nullable = false ) private Long id; @Version private int version; private String title; @ManyToOne (fetch = FetchType.LAZY) @JoinColumn (name = "fk_author" ) private Author author;     ... } 复制代码`

我用10个作者创建了一个测试数据库，他们每人写了10 本书，所以数据库总共包含100 本书。在每个测试中，我将使用不同的投影来查询100 本书并测量执行查询和事务所需的时间。为了减少任何副作用的影响，我这样做1000次并测量平均时间。 OK，让我们开始吧。

### 3.2.查询实体 ###

在大多数应用程序中，实体投影(Entity Projection)是最受欢迎的。有了 ` Entity` ， ` JPA` 可以很容易地将它们用作投影。 运行这个小测试用例并测量检索100个 ` Book` 实体所需的时间。

` long timeTx = 0 ; long timeQuery = 0 ; long iterations = 1000 ; // Perform 1000 iterations for ( int i = 0 ; i < iterations; i++) {     EntityManager em = emf.createEntityManager(); long startTx = System.currentTimeMillis();     em.getTransaction().begin(); // Execute Query long startQuery = System.currentTimeMillis();     List<Book> books = em.createQuery( "SELECT b FROM Book b" ).getResultList(); long endQuery = System.currentTimeMillis();     timeQuery += endQuery - startQuery;     em.getTransaction().commit(); long endTx = System.currentTimeMillis();     em.close();     timeTx += endTx - startTx; } System.out.println( "Transaction: total " + timeTx + " per iteration " + timeTx / ( double )iterations); System.out.println( "Query: total " + timeQuery + " per iteration " + timeQuery / ( double )iterations); 复制代码`

平均而言，执行查询、检索结果并将其映射到100个 ` Book` 实体需要2ms。如果包含事务处理，则为2.89ms。对于小型且不那么新的笔记本电脑来说也不错。

` Transaction: total 2890 per iteration 2.89 Query: total 2000 per iteration 2.0 复制代码`

### 3.3.默认FetchType对To-One关联的影响 ###

当我向你展示Book实体时，我指出我将 [FetchType]( https://link.juejin.im?target=https%3A%2F%2Fthoughts-on-java.org%2Fentity-mappings-introduction-jpa-fetchtypes%2F ) 设置为 ` LAZY` 以避免其他查询。默认情况下， ` To-one` 关联的 ` FetchtType` 是 ` EAGER` ，它告诉 ` Hibernate` 立即初始化关联。

这需要额外的查询，如果你的查询选择多个实体，则会产生巨大的性能影响。让我们更改 ` Book` 实体以使用默认的 ` FetchType` 并执行相同的测试。

` @Entity public class Book { @ManyToOne @JoinColumn (name = "fk_author" ) private Author author;     ... } 复制代码`

这个小小的变化使测试用例的执行时间增加了两倍多。现在花了7.797ms执行查询并映射结果，而不是2毫秒。每笔交易的时间上升到8.681毫秒而不是2.89毫秒。

` Transaction: total 8681 per iteration 8.681 Query: total 7797 per iteration 7.797 复制代码`

因此，最好确保 ` To-one` 关联设置 ` FetchType` 为 ` LAZY` 。

### 3.4.选择@Immutable实体 ###

Joao Charnet在评论中告诉我要在测试中添加一个不可变的实体(Immutable Entity)。有趣的问题是：返回使用 ` @Immutable` 注解的实体，查询性能会更好吗？

` Hibernate` 不必对这些实体执行任何脏检查，因为它们是不可变的。这可能会带来更好的表现。所以，让我们试一试。

我在测试中添加了以下 ` ImmutableBook` 实体。

` @Entity @Table (name = "book" ) @Immutable public class ImmutableBook { @Id @GeneratedValue (strategy = GenerationType.AUTO) @Column (name = "id" , updatable = false , nullable = false ) private Long id; @Version private int version; private String title; @ManyToOne (fetch = FetchType.LAZY) @JoinColumn (name = "fk_author" ) private Author author;     ... } 复制代码`

它是 ` Book` 实体的副本，带有2个附加注解。 ` @Immutable` 注解告诉 ` Hibernate` ，这个实体是不可变得。并且 ` @Table（name =“book”）` 将实体映射到 ` book` 表。因此，我们可以使用与以前相同的数据运行相同的测试。

` long timeTx = 0 ; long timeQuery = 0 ; long iterations = 1000 ; // Perform 1000 iterations for ( int i = 0 ; i < iterations; i++) {     EntityManager em = emf.createEntityManager(); long startTx = System.currentTimeMillis();     em.getTransaction().begin(); // Execute Query long startQuery = System.currentTimeMillis();     List<Book> books = em.createQuery( "SELECT b FROM ImmutableBook b" )             .getResultList(); long endQuery = System.currentTimeMillis();     timeQuery += endQuery - startQuery;     em.getTransaction().commit(); long endTx = System.currentTimeMillis();     em.close();     timeTx += endTx - startTx; } System.out.println( "Transaction: total " + timeTx + " per iteration " + timeTx / ( double )iterations); System.out.println( "Query: total " + timeQuery + " per iteration " + timeQuery / ( double )iterations); 复制代码`

有趣的是，实体是否是不可变的，对查询没有任何区别。测量的事务和查询的平均执行时间几乎与先前的测试相同。

` Transaction: total 2879 per iteration 2.879 Query: total 2047 per iteration 2.047 复制代码`

### 3.5.使用QueryHints.HINT_READONLY查询Entity ###

Andrew Bourgeois建议在测试中包含只读查询。所以，请看这里。

此测试使用我在文章开头向你展示的 ` Book` 实体。但它需要测试用例进行修改。

` JPA` 和 ` Hibernate` 支持一组查询提示(hits)，允许你提供有关查询及其执行方式的其他信息。查询提示 ` QueryHints.HINT_READONLY` 告诉 ` Hibernate` 以只读模式查询实体。因此， ` Hibernate` 不需要对它们执行任何脏检查，也可以应用其他优化。

你可以通过在 ` Query` 接口上调用 ` setHint` 方法来设置此提示。

` long timeTx = 0 ; long timeQuery = 0 ; long iterations = 1000 ; // Perform 1000 iterations for ( int i = 0 ; i < iterations; i++) {     EntityManager em = emf.createEntityManager(); long startTx = System.currentTimeMillis();     em.getTransaction().begin(); // Execute Query long startQuery = System.currentTimeMillis();     Query query = em.createQuery( "SELECT b FROM Book b" );     query.setHint(QueryHints.HINT_READONLY, true );     query.getResultList(); long endQuery = System.currentTimeMillis();     timeQuery += endQuery - startQuery;     em.getTransaction().commit(); long endTx = System.currentTimeMillis();     em.close();     timeTx += endTx - startTx; } System.out.println( "Transaction: total " + timeTx + " per iteration " + timeTx / ( double )iterations); System.out.println( "Query: total " + timeQuery + " per iteration " + timeQuery / ( double )iterations); 复制代码`

你可能希望将查询设置为只读来让性能显著的提升—— ` Hibernate` 执行了更少的工作，因此应该更快。

但正如你在下面看到的，执行时间几乎与之前的测试相同。至少在此测试场景中，将 ` QueryHints.HINT_READONLY` 设置为 ` true` 不会提高性能。

` Transaction: total 2842 per iteration 2.842 Query: total 2006 per iteration 2.006 复制代码`

### 3.6.查询DTO ###

加载100 本书实体大约需要2ms。让我们看看在 ` JPQL` 查询中使用构造函数表达式获取相同的数据是否表现更好。

当然，你也可以在 ` Criteria` 查询中使用构造函数表达式。

` long timeTx = 0 ; long timeQuery = 0 ; long iterations = 1000 ; // Perform 1000 iterations for ( int i = 0 ; i < iterations; i++) {     EntityManager em = emf.createEntityManager(); long startTx = System.currentTimeMillis();     em.getTransaction().begin(); // Execute the query long startQuery = System.currentTimeMillis();     List<BookValue> books = em.createQuery( "SELECT new org.thoughts.on.java.model.BookValue(b.id, b.title) FROM Book b" ).getResultList(); long endQuery = System.currentTimeMillis();     timeQuery += endQuery - startQuery;     em.getTransaction().commit(); long endTx = System.currentTimeMillis();     em.close();     timeTx += endTx - startTx; } System.out.println( "Transaction: total " + timeTx + " per iteration " + timeTx / ( double )iterations); System.out.println( "Query: total " + timeQuery + " per iteration " + timeQuery / ( double )iterations); 复制代码`

正如所料， ` DTO` 投影比 ` 实体(Entity)` 投影表现更好。

` Transaction: total 1678 per iteration 1.678 Query: total 1143 per iteration 1.143 复制代码`

平均而言，执行查询需要1.143ms，执行事务需要1.678ms。查询的性能提升43％，事务的性能提高约42％。

对于一个花费一分钟实现的小改动而言，这已经很不错了。

在大多数项目中， ` DTO` 投影的性能提升将更高。它允许你选择用例所需的数据，而不仅仅是实体映射的所有属性。选择较少的数据几乎总能带来更好的性能。

## 4.摘要 ##

为你的用例选择正确的投影比你想象的更容易也更重要。

如果要实现写入操作，则应使用实体(Entity)作为投影。 ` Hibernate` 将管理其状态，你只需在业务逻辑中更新其属性。然后 ` Hibernate` 会处理剩下的事情。

你已经看到了我的小型性能测试的结果。我的笔记本电脑可能不是运行这些测试的最佳环境，它肯定比生产环境慢。但是性能的提升是如此之大，很明显你应该使用哪种投影。

![file](https://user-gold-cdn.xitu.io/2019/6/4/16b1fa93dea02925?imageView2/0/w/1280/h/960/ignore-error/1)

使用 ` DTO` 投影的查询比选择实体的查询快约40％。因此，最好花费额外的精力为你的只读操作创建 ` DTO` 并将其用作投影。

此外，还应确保对所有关联使用 ` FetchType.LAZY` 。正如在测试中看到的那样，即使是一个热切获取 ` to-one` 的关联操作，也可能会将查询的执行时间增加两倍。因此，最好使用 ` FetchType.LAZY` 并初始化你的用例所需的 [关系]( https://link.juejin.im?target=https%3A%2F%2Fthoughts-on-java.org%2F5-ways-to-initialize-lazy-relations-and-when-to-use-them%2F ) 。

> 
> 
> 
> 原文链接： [thoughts-on-java.org/entities-dt…](
> https://link.juejin.im?target=https%3A%2F%2Fthoughts-on-java.org%2Fentities-dtos-use-projection%2F
> )
> 
> 

> 
> 
> 
> 作者： Thorben Janssen
> 
> 

> 
> 
> 
> 译者：Yunooa
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/4/16b1fabf791ad0c9?imageView2/0/w/1280/h/960/ignore-error/1)