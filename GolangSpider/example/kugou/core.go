package kugou

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"os/exec"
	"strconv"
	"strings"
)


const (
	SHOW_SONG	="ssong"
	SHOW_MV		="smv"
	LIST_SONG   = "lsong"
	LIST_MV     = "lmv"
	DOWNLOAD_BOARD	="gboard"
	DOWNLOAD_SONG    = "gsong"
	DOWNLOAD_MV    = "gmv"
	PLAY_SONG   = "psong"
	PLAY_MV     = "pmv"
	SEARCH_SONG = "qsong"
	SEARCH_MV   = "qmv"
	PATH_MV = "mvpath"
	PATH_SONG   = "songpath"
	CHANGE_PATH_MV = "chmvpath"
	CHANGE_PATH_SONG   = "chsongpath"

	CHSTYLE   = "chstyle"
	STYLE   = "style"
	CHDELIMITER   = "chdelimiter"
	DELIMITER   = "delimiter"
	HELP        = "help"
	QUIT        = "quit"
	EXIT        = "exit"
	CLEAR = "cls"
)

var (
	style="somusic"
	delimiter=">"
)

//处理用户的操作
func DispatcherOperation() {
	var operation []byte
	var err error
	ShowHelp()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s%s",style,delimiter)
		if operation, _, err = reader.ReadLine(); err != nil {
			logs.Error("receive action error.please retry again.")
			continue
		}
		HandleOperation(string(operation))
	}
}



//执行不同的操作
func HandleOperation(operation string) {
	mp := MusicPlayer{}
	//解析参数
	op := strings.SplitN(operation, " ", 2)
	length := len(op)
	//fmt.Println(length)
	if   length == 0 {
		fmt.Println("the operation is error.please show the help doc.")
		return
	} else if length == 1 {
		if len(op[0])==0{
			return
		}
		cmd:=&Command{Action:op[0]}
		switch op[0] {
		case HELP:
			ShowHelp()
		case "h":
			ShowHelp()
		case QUIT:
			os.Exit(0)
		case EXIT:
			os.Exit(0)
		case CLEAR :
			clearLog()
		case "clear":
			clearLog()
		case LIST_SONG:
			cmd.Arguement = "1-5"
			cmd.Arguements = SplitBlockArguements(cmd.Arguement)
			mp.ListSong(cmd)
		case LIST_MV:
			cmd.Arguement = "1-5"
			cmd.Arguements = SplitBlockArguements(cmd.Arguement)
			mp.ListMV(cmd)
		case SHOW_SONG:
			mp.ShowSong(cmd)
		case SHOW_MV:
			mp.ShowMV(cmd)
		case PATH_SONG:
			mp.ShowSongPath(cmd)
		case PATH_MV:
			mp.ShowMVPath(cmd)
		case DOWNLOAD_BOARD:
			mp.DownloadBoardMusic(cmd)
		default:
			fmt.Printf("%s\n", "the operation is error.please show the help doc.")
			//ShowHelp()
			return
		}
	} else {
		op[1] = strings.TrimRight(strings.TrimLeft(op[1], " "), " ")
		if len(op[1]) == 0 || op[1] == "" {
			fmt.Printf("%s%s", op[0], "the arguements is error.")
		}
		cmd := &Command{Action: op[0], Arguement: op[1]}
		switch op[0] {
		case SEARCH_MV:
			//准备参数
			mp.SearchMV(cmd)
		case SEARCH_SONG:
			//准备参数
			mp.SearchSong(cmd)
		case LIST_MV:
			if len(cmd.Arguement)==0{
				cmd.Arguement="5"
			}
			//准备参数
			cmd.Arguement = "1-" + cmd.Arguement
			cmd.Arguements = SplitBlockArguements(cmd.Arguement)
			mp.ListMV(cmd)
		case LIST_SONG:
			//准备参数
			if len(cmd.Arguement)==0{
				cmd.Arguement="5"
			}
			cmd.Arguement = "1-" + cmd.Arguement
			cmd.Arguements = SplitBlockArguements(cmd.Arguement)
			mp.ListSong(cmd)
		case PLAY_MV:
			//准备参数
			cmd.Arguements = []string{cmd.Arguement}
			mp.PlayMV(cmd)
		case PLAY_SONG:
			//准备参数
			cmd.Arguements = []string{cmd.Arguement}
			mp.PlaySong(cmd)
		case DOWNLOAD_SONG:
			//准备参数
			// 1,3,5    1-5,6-10
			if strings.Contains(cmd.Arguement, ",") || strings.Contains(cmd.Arguement, "-") {
				arr := SplitArguemens(cmd.Arguement)
				if arr == nil || len(arr) == 0 {
					fmt.Println("there is no download queue...")
				}
				cmd.Arguements = arr
			} else {
				cmd.Arguements = []string{cmd.Arguement}
			}
			mp.DownloadSong(cmd)
		case DOWNLOAD_MV:
			//准备参数
			// 1,3,5    1-5,6-10
			if strings.Contains(cmd.Arguement, ",") || strings.Contains(cmd.Arguement, "-") {
				arr := SplitArguemens(cmd.Arguement)
				if arr == nil || len(arr) == 0 {
					fmt.Println("there is no download queue...")
				}
				cmd.Arguements = arr
			} else {
				cmd.Arguements = []string{cmd.Arguement}
			}
			mp.DownloadMV(cmd)
		case DELIMITER:
			delimiter=op[1]
		case CHDELIMITER:
			delimiter=op[1]
		case STYLE:
			style=op[1]
		case CHSTYLE:
			style=op[1]
		case CHANGE_PATH_SONG:
			mp.ChangeSongPath(cmd)
		case CHANGE_PATH_MV:
			mp.ChangeMVPath(cmd)
		default:
			fmt.Println("unsupported action,please select operation again..")
			//ShowHelp()
		}
	}
}
//检查输入的命令后的参数是否为空
func checkIsNull(arg string) bool{
	if arg=="" || len(arg)==0{
		fmt.Println("the arguement<",arg,">is must not null,you should retry again...")
		return false
	}
	return true
}

func clearLog() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//分割参数
//1,3,5   1-3,2-7
func SplitArguemens(args string) [] string {
	arguements := make([]string, 0)
	if strings.Contains(args, ",") && strings.Contains(args, "-") {
		splits := strings.Split(args, ",")
		for _, item := range splits {
			if strings.Contains(item, "-") {
				for _, each := range SplitBlockArguements(item) {
					arguements = append(arguements, each)
				}
			} else {
				arguements = append(arguements, item)
			}
		}
	} else if strings.Contains(args, ",") {
		arguements = strings.Split(args, ",")
	} else if strings.Contains(args, "-") {
		arguements = SplitBlockArguements(args)
	}
	return arguements
}

//分割-结合的参数
func SplitBlockArguements(args string) []string {
	arguements := make([]string, 0)
	arrange := strings.SplitN(args, "-", 2)
	start, err := strconv.Atoi(arrange[0])
	if err != nil {
		fmt.Println("the start index:", arrange[0], " is not an number.")
		return nil
	}
	end, err := strconv.Atoi(arrange[1])
	if err != nil {
		fmt.Println("the end index:", arrange[1], " is not an number.")
		return nil
	}
	if start < 0 || start >= end {
		fmt.Println("the start index", arrange[0], " and end index:", arrange[1], " is wrong.")
		return nil
	}
	for ; start <= end; start++ {
		arguements = append(arguements, fmt.Sprintf("%d", start))
	}
	return arguements
}

