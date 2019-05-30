package leetcode

import (
	"os"
)

//定义http请求头部信息
var (
	cookie= `Hm_lvt_fa218a3ff7179639febdb15e372f411c=1559096629; _ga=GA1.2.615852490.1559096629; _gid=GA1.2.2143011937.1559096629; gr_user_id=52ad0d83-13df-4f35-a357-5a0f9700644c; a2873925c34ecbd2_gr_session_id=bd05a216-e094-43c7-a35c-d188bcfa9672; a2873925c34ecbd2_gr_session_id_bd05a216-e094-43c7-a35c-d188bcfa9672=true; grwng_uid=15208bb2-3d4e-4372-a32d-ae635d357a1e; csrftoken=4zvge9nnjQAx8E3P2W8RDK9jYtpio2JvqI2PkuxtCezihFcmcopCebw6kReJsz36; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1559096788; _gat_gtag_UA_131851415_1=1`
	//x_csrftoken=`4zvge9nnjQAx8E3P2W8RDK9jYtpio2JvqI2PkuxtCezihFcmcopCebw6kReJsz36`
	user_agent=`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36`
	content_type="application/json"
	//accept_encoding="deflate, br"
	headers=map[string]string{"cookie":cookie,"user_agent":user_agent,"Content-Type":content_type,}
	)

//cookie:Hm_lvt_fa218a3ff7179639febdb15e372f411c=1559096629; _ga=GA1.2.615852490.1559096629; _gid=GA1.2.2143011937.1559096629; gr_user_id=52ad0d83-13df-4f35-a357-5a0f9700644c; a2873925c34ecbd2_gr_session_id=bd05a216-e094-43c7-a35c-d188bcfa9672; a2873925c34ecbd2_gr_session_id_bd05a216-e094-43c7-a35c-d188bcfa9672=true; grwng_uid=15208bb2-3d4e-4372-a32d-ae635d357a1e; csrftoken=DL7Zremi769ib0dPysOXVqK33dyPAIkAb44eXQ0eyW6neMaYFvr6pDJZ9ZLqHm9x; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1559103584; _gat_gtag_UA_131851415_1=1
//Hm_lvt_fa218a3ff7179639febdb15e372f411c=1559096629; _ga=GA1.2.615852490.1559096629; _gid=GA1.2.2143011937.1559096629; gr_user_id=52ad0d83-13df-4f35-a357-5a0f9700644c; grwng_uid=15208bb2-3d4e-4372-a32d-ae635d357a1e; a2873925c34ecbd2_gr_session_id=a9c6c927-d59d-4ac7-8f46-2f7c6919eb8a; a2873925c34ecbd2_gr_session_id_a9c6c927-d59d-4ac7-8f46-2f7c6919eb8a=true; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=a9c6c927-d59d-4ac7-8f46-2f7c6919eb8a; a2873925c34ecbd2_gr_last_sent_cs1=wenxiaofeicode; _gat_gtag_UA_131851415_1=1; a2873925c34ecbd2_gr_cs1=wenxiaofeicode; csrftoken=VCzrZ7if5lAIQL5YDfxZZB8iaAJ1PpwIDbZNXZG99i8GFEkCNGa4qhTHBG6hpfIE; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1559112500
//x-csrftoken: DL7Zremi769ib0dPysOXVqK33dyPAIkAb44eXQ0eyW6neMaYFvr6pDJZ9ZLqHm9x

//获取所有题目的json数据的url
var (
	//获取所有算法题的链接
	algorithmsUrl=`https://leetcode-cn.com/api/problems/algorithms/`
	//获取所有数据库题的链接
	databaseUrl=`https://leetcode-cn.com/api/problems/database/`
	//获取所有shell题的链接
	shellUrl=`https://leetcode-cn.com/api/problems/shell/`
	//卡片分类
	//cardInfoUrl=`https://leetcode-cn.com/problems/api/card-info/`
	//课题分类
	tagUrl=`https://leetcode-cn.com/problems/api/tags/`

	//获取热门推荐的数据
	recommanderHotUrl=`https://leetcode-cn.com/api/problems/favorite_lists/hot-100/`
	recommanderTencentUrl=`https://leetcode-cn.com/api/problems/favorite_lists/50/`
	recommanderTopUrl=`https://leetcode-cn.com/api/problems/favorite_lists/top/`
	recommanderFavoriateUrl=`https://leetcode-cn.com/problems/api/favorites/`


)

//题目相关的url
var (

	//
	commonUrl=`https://leetcode-cn.com/graphql`
	//翻译题目链接 post请求

	transQuestionsParamTemplate=`{"operationName":"getQuestionTranslation","variables":{},"query":"query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    questionId\n    __typename\n  }\n}\n"}`

	//请求单个题目的信息	post请求
	//请求题目信息：
	questionParam="#titleSlum#"
	questionParamTemplate=`{"operationName":"questionData","variables":{"titleSlug":"#titleSlum#"},"query":"query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"}`

	//请求所有题解：
	answersParam="#questionSlug#"
	answersParamTemplate=`{"operationName":"questionSolutionArticles","variables":{"questionSlug":"#questionSlug#","first":10,"skip":0,"orderBy":"DEFAULT"},"query":"query questionSolutionArticles($questionSlug: String!, $skip: Int, $first: Int, $orderBy: SolutionArticleOrderBy, $userInput: String) {\n  questionSolutionArticles(questionSlug: $questionSlug, skip: $skip, first: $first, orderBy: $orderBy, userInput: $userInput) {\n    totalNum\n    edges {\n      node {\n        ...article\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment article on SolutionArticleNode {\n  title\n  slug\n  reactedType\n  status\n  identifier\n  canEdit\n  reactions {\n    count\n    reactionType\n    __typename\n  }\n  tags {\n    name\n    nameTranslated\n    slug\n    __typename\n  }\n  createdAt\n  thumbnail\n  author {\n    username\n    profile {\n      userAvatar\n      userSlug\n      realName\n      __typename\n    }\n    __typename\n  }\n  summary\n  topic {\n    id\n    commentCount\n    viewCount\n    __typename\n  }\n  byLeetcode\n  isMyFavorite\n  isMostPopular\n  isEditorsPick\n  upvoteCount\n  upvoted\n  hitCount\n  __typename\n}\n"}`
	//请求单个题解的详细信息：
	//https://leetcode-cn.com/graphql POST
	answerDetailParam="#articleSlug#"
	answerDetailParamTemplate=`{"operationName":"solutionDetailArticle","variables":{"slug":"#articleSlug#"},"query":"query solutionDetailArticle($slug: String!) {\n  solutionArticle(slug: $slug) {\n    ...article\n    content\n    __typename\n  }\n}\n\nfragment article on SolutionArticleNode {\n  title\n  slug\n  reactedType\n  status\n  identifier\n  canEdit\n  reactions {\n    count\n    reactionType\n    __typename\n  }\n  tags {\n    name\n    nameTranslated\n    slug\n    __typename\n  }\n  createdAt\n  thumbnail\n  author {\n    username\n    profile {\n      userAvatar\n      userSlug\n      realName\n      __typename\n    }\n    __typename\n  }\n  summary\n  topic {\n    id\n    commentCount\n    viewCount\n    __typename\n  }\n  byLeetcode\n  isMyFavorite\n  isMostPopular\n  isEditorsPick\n  upvoteCount\n  upvoted\n  hitCount\n  __typename\n}\n"}`

	//获取不同分类标签的数据
	tagData="#tagSlug#"
	tagDataTemplate=`{"operationName":"getTopicTag","variables":{"slug":"#tagSlug#"},"query":"query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    translatedName\n    questions {\n      status\n      questionId\n      questionFrontendId\n      title\n      titleSlug\n      translatedTitle\n      stats\n      difficulty\n      isPaidOnly\n      topicTags {\n        name\n        translatedName\n        slug\n        __typename\n      }\n      __typename\n    }\n    frequencies\n    __typename\n  }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n      __typename\n    }\n    privateFavorites {\n      ...favoriteFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n    __typename\n  }\n  __typename\n}\n"}`
)


//定义目录和文件名
var (
	dataDir="data"
	questionDir="data\\questions"
	answersDir="answers"
	tagDir="data\\tags"
	recommandDir="data\\recommand"
	//eachAnswerJsonfile=questionDir+string(os.PathSeparator)+"#question#"+string(os.PathSeparator)+answersDir+string(os.PathSeparator)+"#answer#.json"
	//answersJsonfile=questionDir+string(os.PathSeparator)+"#question#"+string(os.PathSeparator)+answersDir+string(os.PathSeparator)+"#answer#.json"
	eachQuestionDir=questionDir+string(os.PathSeparator)+"#filename#"
	eachTagJsonfile=tagDir+string(os.PathSeparator)+"#filename#.json"
	algorithmsJsonfile=dataDir+string(os.PathSeparator)+"algorithms.json"
	transAlgorithmsJsonfile=dataDir+string(os.PathSeparator)+"translateAlgorithms.json"
	questionsJsonfile=questionDir+string(os.PathSeparator)+"questions.json"
	databaseJsonfile=dataDir+string(os.PathSeparator)+"database.json"
	shellJsonfile=dataDir+string(os.PathSeparator)+"shell.json"
	topJsonfile=recommandDir+string(os.PathSeparator)+"top.json"
	hot100Jsonfile=recommandDir+string(os.PathSeparator)+"hot100.json"
	tencent50Jsonfile=recommandDir+string(os.PathSeparator)+"tencent50.json"
	favoriatesJsonfile=recommandDir+string(os.PathSeparator)+"favoriates.json"
	tagsJsonfile=dataDir+string(os.PathSeparator)+"tags.json"
)


var (
	questionIdPath=`stat_status_pairs.#.stat.question_id`
	questionTitlePath=`stat_status_pairs.#.stat.question__title`
	questionTitleSlugPath=`stat_status_pairs.#.stat.question__title_slug`
	questionTotalAcsPath=`stat_status_pairs.#.stat.total_acs`
	questionTotalSubmittedPath=`stat_status_pairs.#.stat.total_submitted`
	questionTotalArticlesPath=`stat_status_pairs.#.stat.total_column_articles`
	questionDifficultyLevelPath=`stat_status_pairs.#.difficulty.level`

	difficlutLevel=map[int]string{1:"简单",2:"中等",3:"困难"}

	transQuestionTitlePath=`data.translations.#.title`

)


