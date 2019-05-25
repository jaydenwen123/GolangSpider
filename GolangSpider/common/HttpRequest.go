package common

import (
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	USER_AGENT  = "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36"
	RETRY_TIMES = 3
)

func RequestJsonWithRetry(url string, headers map[string]string) string {

	return RequestJsonWithTimes(url, headers, RETRY_TIMES)
}

func RequestJsonWithTimes(url string, headers map[string]string, count int) string {

	for i := 0; i < count; i++ {
		data := RequestJson(url, headers)
		if len(data) > 0 && data != "" {
			return data
		} else {
			logs.Info("there is occurd error. please wait some time.there is now downloading retry...")
			time.Sleep(time.Millisecond * 30)
		}
	}
	logs.Error("there is occurd error.we have trtry for", count, "times.....")
	return ""
}

//通过get发送请求，返回数据
//第一个参数为字节数组，第二个参数为默认编码为utf-8的字符串
func RequestJson(url string, headers map[string]string) string {

	//1.发请求，获取数据
	//如果需要自己设置请求头，则通过http.NewRequest
	//resp, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	//设置请求头
	request.Header.Add("User-Agent", USER_AGENT)
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	//发送请求
	transport := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(120 * time.Second)
			c, err := net.DialTimeout(netw, addr, time.Second*120)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(deadline)
			return c, nil
		},
		IdleConnTimeout:       120 * time.Second,
		TLSHandshakeTimeout:   120 * time.Second,
		ResponseHeaderTimeout: 120 * time.Second,
	}

	client := &http.Client{
		Timeout:   120 * time.Second,
		Transport: transport,
	}
	resp, err := client.Do(request)
	if err != nil {
		logs.Error("http get error:", err.Error())
		//panic(err.Error())
		//resp.Body.Close()
		//defer resp.Body.Close()
		return ""
		//logs.Info("please wait for time.there is now retrying download....")
		//return RequestJson(url,headers)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("ioutil ReadAll error:", err.Error())
		return ""
	}
	if err = resp.Body.Close(); err != nil {
		logs.Error("resp Body Close error:", err.Error())
		return ""
	}
	return string(data)
}

//通过get发送请求，返回数据
//第一个参数为字节数组，第二个参数为默认编码为utf-8的字符串
func Request(url string) ([]byte, string) {

	return RequestWithHeader(url, nil)
}

//通过get发送请求，返回数据
//第一个参数为字节数组，第二个参数为默认编码为utf-8的字符串
func RequestWithHeader(url string, headers map[string]string) ([]byte, string) {

	//1.发请求，获取数据
	//如果需要自己设置请求头，则通过http.NewRequest
	//resp, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	//设置请求头
	request.Header.Add("User-Agent", USER_AGENT)
	if headers!=nil{
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}
	//发送请求
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 120 * time.Second,
		}).Dial,
		IdleConnTimeout:       120 * time.Second,
		TLSHandshakeTimeout:   120 * time.Second,
		ResponseHeaderTimeout: 120 * time.Second,
	}

	client := &http.Client{
		Timeout:   120 * time.Second,
		Transport: transport,
	}
	resp, err := client.Do(request)
	if err != nil {
		logs.Error("http get error:", err.Error())
		panic(err.Error())
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("ioutil ReadAll error:", err.Error())
		return nil, ""
	}
	if err = resp.Body.Close(); err != nil {
		logs.Error("resp Body Close error:", err.Error())
		return nil, ""
	}
	return content, string(content)
}

//通过get发送请求，返回数据
//第一个参数为字节数组，第二个参数为默认编码为utf-8的字符串
func ResponseWithReader(url string) io.Reader {

	//1.发请求，获取数据
	//如果需要自己设置请求头，则通过http.NewRequest
	//resp, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	//设置请求头
	request.Header.Add("User-Agent", USER_AGENT)
	//发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logs.Error("create request error")
		//panic(err.Error())
		return nil
	}
	if err != nil {
		logs.Error("http get error:", err.Error())
		//panic(err.Error())
		return nil
	}
	return resp.Body
}
