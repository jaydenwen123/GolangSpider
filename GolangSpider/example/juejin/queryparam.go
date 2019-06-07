package juejin

import "github.com/jaydenwen123/go-util"
//以下构造post请求中请求体内部的参数
type Variable struct {
	Type string	`json:"type"`
	Query string	`json:"query"`
	After string	`json:"after"`
	Period	string	`json:"period"`
	First int	`json:"first"`
	//只有在热榜、最新里面用到该关键词
	Order string	`json:"order,omitempty"`
}

type Query struct {
	Id string	`json:"id"`
}

type Extension struct {
	Query Query	`json:"query"`
}

type PostQueryBody struct {
	OperationName string`json:"operationName"`
	Query string	`json:"query"`
	Variables Variable`json:"variables"`
	Extensions Extension`json:"extensions"`
}

//搜索参数
//{"operationName":"","query":"","variables":{"type":"ALL","query":"golang","after":"","period":"M3","first":20},"extensions":{"query":{"id":"d9997080c3d67a02bfdae094729fed3b"}}}
func NewSearchQueryParam(id string,first int,after string,period string,typ string,query string) *PostQueryBody {
	return &PostQueryBody{Variables:Variable{
		Type:typ,
		Query:query,
		After:after,
		Period:period,
		First:first,
	},
	Extensions:Extension{Query:Query{Id:id}}}
}

//推荐参数（热门、最新、热榜）
//{"operationName":"","query":"","variables":{"first":20,"after":"","order":"POPULAR"},"extensions":{"query":{"id":"21207e9ddb1de777adeaca7a2fb38030"}}}
//{"operationName":"","query":"","variables":{"first":20,"after":"","order":"THREE_DAYS_HOTTEST"},"extensions":{"query":{"id":"21207e9ddb1de777adeaca7a2fb38030"}}}
//{"operationName":"","query":"","variables":{"first":20,"after":"","order":"NEWEST"},"extensions":{"query":{"id":"21207e9ddb1de777adeaca7a2fb38030"}}}
func NewRecommandQueryParam(id string,first int,after string,order string) *PostQueryBody {
	return &PostQueryBody{Variables:Variable{
		First:first,
		After:after,
		Order:order,
	},
	Extensions:Extension{Query:Query{id}}}
}

func (p *PostQueryBody) String() string {
	return util.Obj2JsonStr(p)
}


