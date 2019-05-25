package util

import (
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)
const (
	DOWNLOAD_TETRY_TIMES=3
	TIMEOUT=120
)

func DownloadWithRetry(url string,filename string)  error{
	return DownloadWithTimes(url,filename,DOWNLOAD_TETRY_TIMES)
}

//设置下载失败时的请求次数
func DownloadWithTimes(url string,filename string,times int)  error{
	if times<=0{
		times=DOWNLOAD_TETRY_TIMES
	}
	for i:=0;i<times;i++{
		err := Download(url, filename)
		if err==nil{
			if i!=0{
				logs.Info("there is retrying",i,"times to download  success...")
			}
			return err
		}else{
			logs.Info("there is happened download error.now is ",i+1,"time retrying...")
			time.Sleep(time.Millisecond*30)
		}
	}
	return errors.New("there is retry with"+strconv.Itoa(times)+"times.but still is download error.please retry again after several minutes.")
}

//根据url下载图片、视频等二进制文件
func Download(url string, filename string) error {
	if len(url)==0 || url==""{
		return errors.New("the url's format is errored.")
	}

	//1.发请求，获取数据
	//如果需要自己设置请求头，则通过http.NewRequest
	//resp, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	//发送请求
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: TIMEOUT * time.Second,
		}).Dial,
		IdleConnTimeout:       TIMEOUT * time.Second,
		TLSHandshakeTimeout:   TIMEOUT* time.Second,
		ResponseHeaderTimeout: TIMEOUT * time.Second,
	}

	client := &http.Client{
		Timeout:   TIMEOUT * time.Second,
		Transport: transport,
	}
	resp, err := client.Do(request)
	if err != nil {
		logs.Error("http get error:", err.Error())
		return err
	}
	//
	//
	//resp, err := http.Get(url)
	//if err != nil {
	//	panic(err.Error())
	//}
	defer resp.Body.Close()
	//读取响应体的内容
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("download file error!", err.Error())
		return err
	}
	//请求下来的内容为空，返回自定义的错误，然后重新下载
	if len(data)==0{
		logs.Error("there is network error,which cause not download content......")
		return errors.New("there is network error,which cause not download content.....")
	}
	//写入到文件中
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		logs.Error("download file error!", err.Error())
		return err
	}
	//解决Golang"Connection reset"&"EOF"问题
	request.Close=true
	return nil
}
