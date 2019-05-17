package kugou

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"os"
	"strings"
)

const (
	LIST_SONG="lsong"
	LIST_MV="lmv"
	DOWNLOAD="get"
	PLAY_SONG="psong"
	PLAY_MV="pmv"
	SEARCH_SONG="qsong"
	SEARCH_MV="qmv"
	HELP="help"
	QUIT="quit"
	EXIT="exit"
)

//定义命令显示
type Command struct {
	Action string
	Arguements []string
}

func NewCommand(action string, arguements []string) *Command {
	return &Command{Action: action, Arguements: arguements}
}




//处理用户的操作
func DispatcherOperation(results []gjson.Result) {
	var operation []byte
	var err error
	ShowHelp()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("command>")
		if operation, _,err = reader.ReadLine();err!=nil{
			logs.Error("receive action error.please retry again.")
			continue
		}
		HandleOperation(string(operation))
	}
}
//执行不同的操作
func HandleOperation(operation string) {
	//解析参数
	op := strings.Split(operation, " ")
	len:=len(op)
	if len==0{
		fmt.Println("the operation is error.please show the help doc.")
	}else if len==1{
		switch op[0] {
		case HELP:
			ShowHelp()
		case "h":
			ShowHelp()
		case QUIT:
			os.Exit(0)
		case EXIT:
			os.Exit(0)
		}
	}else{
		switch op[0] {
		case SEARCH_MV:

		case SEARCH_SONG:

		case LIST_MV:

		case LIST_SONG:

		case PLAY_MV:

		case PLAY_SONG:

		case DOWNLOAD:

		}
	}
}

//显示操作说明
func ShowHelp() {
	//1.下载排名靠前的多少首歌曲
	//2.查看歌曲信息列表<歌曲名、专辑、时长>
	//3.播放歌曲
	//4.查看帮助文档
	//5.退出程序
	fmt.Println("here is the usage of  action :")
	fmt.Println("\t<lsong> <first-end>:means show the asc range musics list ")
	fmt.Println("\t<get> <number> or <first1-end1,first2-end2...>:means download \none music or download according the range")
	fmt.Println("\t<lmv> <first-end>:means show the asc range mv list ")
	fmt.Println("\t<psong> <number>:means play the mv")
	fmt.Println("\t<pmv> <number>:means play the song")
	fmt.Println("\t<qsong> <keyword>:means to query song by key word")
	fmt.Println("\t<qmv> <keyword>:means to query mv by key word")
	fmt.Println("\t<help> or <h>:means to show the help information...")
	fmt.Println("\t<quit> or CTRL+C:means to quit the program...")
}