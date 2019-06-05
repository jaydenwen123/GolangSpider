package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func ProxySendRequest()  {
	reqeustUrl:="https://weixin.talkedu.cn/Web/Vote/mobile/1/VoteMobileHandler?Action=DoVote "
	param:="app_case_id=7513&option_id=7509"
	header:=map[string]string{
		"User-Agent":`Mozilla/5.0 (Linux; Android 8.0.0; BLN-AL40 Build/HONORBLN-AL40; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044704 Mobile Safari/537.36 MMWEBID/4483 MicroMessenger/7.0.4.1420(0x2700043B) Process/tools NetType/WIFI Language/zh_CN`,
		"Origin":`https://weixin.talkedu.cn`,
		"Content-Type":`application/x-www-form-urlencoded; charset=UTF-8`,
		///x-www-form-urlencoded; charset=UTF-8
		//&sign=f132bb30650e266d2364edf7a1ed5994,f132bb30650e266d2364edf7a1ed5994,4fd7728caa07f9ef099e5953379c0fbf,49b684673477c293b535e0f0e3d13e86,49b684673477c293b535e0f0e3d13e86
		//"Referer":"https://weixin.talkedu.cn/Web/Vote/mobile/1/OptionDescription.aspx?1=1&1=1&app_case_id=7513&option_id=7509&from=timeline",
		"Referer":"https://weixin.talkedu.cn/Web/Vote/mobile/1/OptionDescription.aspx?1=1&1=1&wx_mp_id=6688&app_case_id=7513&option_id=7509&sign=f132bb30650e266d2364edf7a1ed5994,f132bb30650e266d2364edf7a1ed5994,4fd7728caa07f9ef099e5953379c0fbf,49b684673477c293b535e0f0e3d13e86,49b684673477c293b535e0f0e3d13e86&from=timeline",
		"Content-Length":"31",
		"Cookie":"ASP.NET_SessionId=whuxorhj2icugrsr1e4pfg02",
		"X-Requested-With":"XMLHttpRequest",
		//
	}
	//data := common.RequestJsonWithPost(reqeustUrl, header, param)
	//代理请求发送
	proxy,_:= url.Parse("https://112.85.169.44:9999")



	zTransport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	request,_:=http.NewRequest("POST",reqeustUrl,strings.NewReader(param))
	for key,value:=range header{
		request.Header.Add(key,value)
	}
	//resp, err := zTransport.RoundTrip(request)

	//transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: zTransport}

	resp, err := client.Do(request)
	if err!=nil{
		println("出错了",err.Error())
		return
	}
	data,_:=ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(string(data))
}
