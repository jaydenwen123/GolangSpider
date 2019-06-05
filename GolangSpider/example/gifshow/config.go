package gifshow

//定义http请求头部信息

var (
	gifUrl=`http://api.gifshow.com/rest/n/feed/hot?app=0&kpf=ANDROID_PHONE&ver=6.5&c=HUAWEI&mod=HUAWEI%28BLN-AL40%29&appver=6.5.1.9253&ftt=&isp=CMCC&kpn=KUAISHOU&lon=0&language=zh-cn&sys=ANDROID_8.0.0&max_memory=384&ud=19523983&country_code=cn&oc=HUAWEI&hotfix_ver=&did_gt=1558500302521&iuid=&extId=84ae6d25936c4eab63503b236a35ad34&net=WIFI&did=ANDROID_d5c03a4945197bbf&lat=0`
	cookie=`token=647e4728fff6462ca87784fa7cced390-19523983`
	user_agent=`kwai-android`
	connection=`keep-alive`
	accept_language=`zh-cn`
	x_requestid= `641013794`
	content_type= `application/x-www-form-urlencoded`
	content_length= `367`
	host=`api.gifshow.com`
	accept_encoding= `gzip`
	headers=map[string]string{
		//"Cookie":cookie,
		"User-Agent":user_agent ,
		//"Connection":connection ,
		//"Accept-Language":accept_language ,
		//"X-REQUESTID":x_requestid ,
		"Content-Type": content_type,
		//"Content-Length": content_length,
		//"Host": host,
		//"Accept-Encoding":accept_encoding,
	}
	bodyData=`type=7&page=1&coldStart=true&count=20&pv=false&id=24&refreshTimes=0&pcursor=&source=1&needInterestTag=false&browseType=1&seid=1277360e-773a-4e23-80f7-c1eab0a921de&volume=0.0&os=android&__NStokensig=58b6fb7c78de2865b0935ed3c62665639872c1e33592476d9c795dfcfe8de554&token=647e4728fff6462ca87784fa7cced390-19523983&sig=791ce6cbae5d700e88fc35033c7d6a3d&client_key=3c2cd3f3`
)
