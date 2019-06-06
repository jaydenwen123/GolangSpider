package juejin

import (
	"GolangSpider/GolangSpider/util"
	"fmt"
	"testing"
)

func TestConfigParam(t *testing.T)  {
	fmt.Println(util.Obj2JsonStr(hottestParam))
	fmt.Println(util.Obj2JsonStr(newestParam))
	fmt.Println(util.Obj2JsonStr(searchArticlesParam))
	fmt.Println(util.Obj2JsonStr(searchALLParam))
}
