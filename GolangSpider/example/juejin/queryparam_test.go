package juejin

import (
	"github.com/jaydenwen123/go-util"
	"fmt"
	"testing"
)

func TestGenerateQueryParam(t *testing.T)  {
	query:=PostQueryBody{}
	fmt.Println(util.Obj2JsonStr(query))
}