package kugou

import "fmt"

//显示操作说明
func ShowHelp() {
	//1.下载排名靠前的多少首歌曲

	//2.查看歌曲信息列表<歌曲名、专辑、时长>
	//3.播放歌曲
	//4.查看帮助文档
	//5.退出程序
	fmt.Println("here is the usage of  ",style,":")
	fmt.Println("\t<gboard> \n\t\t:means download the kugou rank board all song. ")
	fmt.Println("\t<lsong> <first-end>\n\t\t:means show the asc range musics list ")
	fmt.Println("\t<lmv> <first-end>:\n\t\tmeans show the asc range mv list ")
	fmt.Println("\t<gsong> <number> or <first1-end1,first2-end2...>:\n\t\tmeans download \n\t\tone music or download according the range")
	fmt.Println("\t<gmv> <number> or <first1-end1,first2-end2...>:\n\t\tmeans download \n\t\tone mv or download according the range")
	fmt.Println("\t<psong> <number>:\n\t\tmeans play the mv")
	fmt.Println("\t<pmv> <number>:\n\t\tmeans play the song")
	fmt.Println("\t<qsong> <keyword>:\n\t\tmeans to query song by key word")
	fmt.Println("\t<qmv> <keyword>:\n\t\tmeans to query mv by key word")
	fmt.Println("\t<ssong> :\n\t\tmeans show the download song list ")
	fmt.Println("\t<smv> :\n\t\tmeans show the download mv list ")
	fmt.Println("\t<chstyle> <newstyle> or <style> <newstyle>:\n\t\tmeans to change the style...")
	fmt.Println("\t<chdelimiter> <newdelimiter>or <delimiter> <newdelimiter>:\n\t\tmeans to change the delimiter...")
	fmt.Println("\t<mvpath> :\n\t\tmeans to show the save downloaded mv path...")
	fmt.Println("\t<songpath> :\n\t\tmeans to show the save downloaded song path...")
	fmt.Println("\t<chmvpath> <newmvpath>:\n\t\tmeans to change the save downloaded mv path.use the ~ to recovery the default dirctory...")
	fmt.Println("\t<chsongpath> <newsongpath>:\n\t\tmeans to change the save downloaded song path.use the ~ to recovery the default dirctory...")
	fmt.Println("\t<help> or <h>:\n\t\tmeans to show the help information...")
	fmt.Println("\t<quit> or CTRL+C:\n\t\tmeans to quit the program...")
	fmt.Println("\t<exit> or CTRL+C:\n\t\tmeans to exit the program...")
	fmt.Println("\t<cls> or <clear>:\n\t\tmeans to clear the log info...")
}
