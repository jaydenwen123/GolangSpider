# 力扣LeetCode #

> leetcode中文网址：`https://leetcode-cn.com/`  
> leetcode英文网址：`https://leetcode-cn.com/`  
> 
> ![力扣leetcode](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/home.png)


该实例主要目标是爬取LeetCode中文版上面的所有算法题目数据、题解数据、热门推荐题目数据以及话题分类的详细数据

## 主要工作 ##

### **1.分析网站api接口** ###

> `cookie：Hm_lvt_fa218a3ff7179639febdb15e372f411c=1559096629; _ga=GA1.2.615852490.1559096629; _gid=GA1.2.2143011937.1559096629; gr_user_id=52ad0d83-13df-4f35-a357-5a0f9700644c; a2873925c34ecbd2_gr_session_id=bd05a216-e094-43c7-a35c-d188bcfa9672; a2873925c34ecbd2_gr_session_id_bd05a216-e094-43c7-a35c-d188bcfa9672=true; grwng_uid=15208bb2-3d4e-4372-a32d-ae635d357a1e; csrftoken=4zvge9nnjQAx8E3P2W8RDK9jYtpio2JvqI2PkuxtCezihFcmcopCebw6kReJsz36; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1559096788; _gat_gtag_UA_131851415_1=1`  

**1.获取所有题目的json数据的api**  

> https://leetcode-cn.com/api/problems/algorithms/
> https://leetcode-cn.com/api/problems/database/
> https://leetcode-cn.com/api/problems/shell/


**2.卡片分类api**

> https://leetcode-cn.com/problems/api/card-info/


**3.搜索题目的api**
> https://leetcode-cn.com/problems/api/filter-questions/#filter#

**4.获取热门推荐数据的api**
> https://leetcode-cn.com/api/problems/favorite_lists/hot-100/
> https://leetcode-cn.com/api/problems/favorite_lists/50/
> https://leetcode-cn.com/api/problems/favorite_lists/top/
> https://leetcode-cn.com/problems/api/favorites/


**5.翻译题目的api**
> https://leetcode-cn.com/graphql  
> post请求参数  
> {"operationName":"getQuestionTranslation","variables":{},"query":"query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    questionId\n    __typename\n  }\n}\n"}  

**6.话题分类api**
> https://leetcode-cn.com/problems/api/tags/

**7.获取不同标签的数据**
> https://leetcode-cn.com/graphql  
> post请求参数  
> {"operationName":"getTopicTag","variables":{"slug":"array"},"query":"query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    translatedName\n    questions {\n      status\n      questionId\n      questionFrontendId\n      title\n      titleSlug\n      translatedTitle\n      stats\n      difficulty\n      isPaidOnly\n      topicTags {\n        name\n        translatedName\n        slug\n        __typename\n      }\n      __typename\n    }\n    frequencies\n    __typename\n  }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n      __typename\n    }\n    privateFavorites {\n      ...favoriteFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n    __typename\n  }\n  __typename\n}\n"}


**8.请求单个题目信息的api**
> https://leetcode-cn.com/graphql  
> post请求参数  
> {"operationName":"questionData","variables":{"titleSlug":"two-sum"},"query":"query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"}  


**9.请求所有题解的api**
> https://leetcode-cn.com/graphql  
> post请求参数 
> {"operationName":"questionSolutionArticles","variables":{"questionSlug":"add-two-numbers","first":10,"skip":0,"orderBy":"DEFAULT"},"query":"query questionSolutionArticles($questionSlug: String!, $skip: Int, $first: Int, $orderBy: SolutionArticleOrderBy, $userInput: String) {\n  questionSolutionArticles(questionSlug: $questionSlug, skip: $skip, first: $first, orderBy: $orderBy, userInput: $userInput) {\n    totalNum\n    edges {\n      node {\n        ...article\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment article on SolutionArticleNode {\n  title\n  slug\n  reactedType\n  status\n  identifier\n  canEdit\n  reactions {\n    count\n    reactionType\n    __typename\n  }\n  tags {\n    name\n    nameTranslated\n    slug\n    __typename\n  }\n  createdAt\n  thumbnail\n  author {\n    username\n    profile {\n      userAvatar\n      userSlug\n      realName\n      __typename\n    }\n    __typename\n  }\n  summary\n  topic {\n    id\n    commentCount\n    viewCount\n    __typename\n  }\n  byLeetcode\n  isMyFavorite\n  isMostPopular\n  isEditorsPick\n  upvoteCount\n  upvoted\n  hitCount\n  __typename\n}\n"} 
 
**10.请求单个题解的详细信息的api**
> https://leetcode-cn.com/graphql  
> post请求参数  
> {"operationName":"solutionDetailArticle","variables":{"slug":"liang-shu-xiang-jia-by-leetcode"},"query":"query solutionDetailArticle($slug: String!) {\n  solutionArticle(slug: $slug) {\n    ...article\n    content\n    __typename\n  }\n}\n\nfragment article on SolutionArticleNode {\n  title\n  slug\n  reactedType\n  status\n  identifier\n  canEdit\n  reactions {\n    count\n    reactionType\n    __typename\n  }\n  tags {\n    name\n    nameTranslated\n    slug\n    __typename\n  }\n  createdAt\n  thumbnail\n  author {\n    username\n    profile {\n      userAvatar\n      userSlug\n      realName\n      __typename\n    }\n    __typename\n  }\n  summary\n  topic {\n    id\n    commentCount\n    viewCount\n    __typename\n  }\n  byLeetcode\n  isMyFavorite\n  isMostPopular\n  isEditorsPick\n  upvoteCount\n  upvoted\n  hitCount\n  __typename\n}\n"}  

**11.获取所有题数据的api**
> https://leetcode-cn.com/graphql  
> post请求参数  
> {"operationName":"allQuestions","variables":{},"query":"query allQuestions {\n  allQuestions {\n    ...questionSummaryFields\n    __typename\n  }\n}\n\nfragment questionSummaryFields on QuestionNode {\n  title\n  titleSlug\n  translatedTitle\n  questionId\n  questionFrontendId\n  status\n  difficulty\n  isPaidOnly\n  __typename\n}\n"}

**12.获取某一题的数据**
> https://leetcode-cn.com/graphql  
> post请求参数  
> {"operationName":"questionData","variables":{"titleSlug":"symmetric-tree"},"query":"query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"}


### **2.爬取所有题目列表** ###
> ![爬取所有题目列表](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/all.png)

### **3.爬取每个题目的题解数据** ###
> ![爬取每个题目的题解数据](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/question_answer.png)


### **4.爬取热门推荐所有题目数据** ###
> ![爬取热门推荐所有题目数据](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/recommander.png)


### **5.爬取所有话题分类数据** ###
> ![爬取所有话题分类数据](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/type.png)



## 成果展示 ##

### **0.文件存储目录结构** ###

> **爬取所有数据的存储目录**  
> ![爬取所有数据的存储目录](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/directory.png)  
> 
> **题目信息保存目录**   
> ![题目信息保存目录](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/question_dir.png)  
> 
> **题解保存目录**  
> ![题解保存目录](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/answer_dir.png)

### **1.爬取所有题目列表并存储json文件** ###
> **爬取所有题目列表并存储json文件**  
> ![爬取所有题目列表并存储json文件](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/algorithms.png)  
> 
> **爬取所有题目列表并存储json文件**  
> ![爬取所有题目列表并存储json文件](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/question_list.png)


### **2.爬取单个题目的详细信息** ###
> **爬取所有话题分类数据**  
> ![爬取所有话题分类数据](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/question_detail.png)

### **3.爬取每个题目的题解数据并存储json文件** ###
> **爬取所有题解信息**  
> ![爬取所有题解信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/answers.png) 
>  
> **爬取每个题解详细信息**  
> ![爬取每个题解详细信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/answer_detail.png)  


### **4.爬取热门推荐所有题目数据并存储json文件** ###
> **热题 HOT 100**     
> ![热题 HOT 100](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/hot100.png)  
> 
> **腾讯精选练习（50 题）**  
> ![ 腾讯精选练习（50 题）](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/tencent50.png)  
> 
> **精选TOP面试题**  
> ![ 精选 TOP 面试题](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/top.png)  

### **5.爬取所有话题分类数据并存储json文件** ###
> **爬取所有话题分类数据**  
> ![爬取所有话题分类数据](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/tags.png)  
> 
> **爬取单个话题数据**  
> ![爬取单个话题数据](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/one_tag.png)  
> 
### **6.建立Http服务器在线获取题目列表、单个题目详细信息** ###
> **在线获取单个题目详细信息**  
> ![在线获取单个题目详细信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/web_question_detail.png)  
> 
> **在线获取题目列表信息**  
> ![在线获取题目列表信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/leetcode/images/web_questions_list.png)

## 涉及技术 ##
1. http post请求json数据
2. goroutine & channel并发爬取数据
3. gjson解析json数据
4. golang 单元测试框架Testing使用