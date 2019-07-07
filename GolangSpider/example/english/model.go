package english

type Content struct {
	Text string
	File string
	Name string
	IsVideo bool
}
type ArticleDetail struct {
	Name string
	Link string
	Times string
	Date string
	Content *Content
}
type Article struct {
	Name string
	Link string
	Times string
	Date string
	ArticleDetails []*ArticleDetail
}

type Channel struct {
	Name string
	Link string
	Articles []*Article
}

type Category struct {
	Name string//一级栏目名称
	Channels []*Channel
}