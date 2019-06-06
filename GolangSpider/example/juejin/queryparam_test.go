package juejin

import (
	"GolangSpider/GolangSpider/util"
	"fmt"
	"testing"
)

func TestGenerateQueryParam(t *testing.T)  {
	query:=PostQueryBody{}
	fmt.Println(util.Obj2JsonStr(query))
}