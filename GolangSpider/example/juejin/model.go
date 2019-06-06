package juejin

type Article struct {
	Title string
	Id string
	OriginalUrl string
	CommentCount	int
	LikeCount	int
}

//获取下一页的信息
type PageInfo struct {
	HasNextPage bool
	EndCursor string
}