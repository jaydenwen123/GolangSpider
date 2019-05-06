package response

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type ResponseBody struct {
	RespInfo *ResponseInfo `json:"resp_info"`
	Data interface{}	`json:"data"`
}

//错误回应
func SendErrorResponse(w http.ResponseWriter,er ErrorResponseInfo)  {
	jsonEr, err := json.Marshal(er)
	if err != nil {
		logs.Error("json encoding error in SendErrorResponse",err.Error())
		w.WriteHeader(InternalError.RespCode)
		jsonEr,_=json.Marshal(InternalError)
		_,_=w.Write(jsonEr)
	}else{
		w.WriteHeader(er.RespCode)
		_,err = w.Write(jsonEr)
		if err != nil {
			logs.Error("when write the data to client occurs error",err.Error())
		}
	}
}
//正常的回应
func SendNormalResponse(w http.ResponseWriter,info ResponseInfo)  {
	w.WriteHeader(info.RespCode)
	respData, err:= json.Marshal(info)
	if err!=nil{
		SendErrorResponse(w,InternalError)
		return
	}
	_, err = w.Write(respData)
	if err!=nil{
		SendErrorResponse(w,InternalError)
		return
	}
}
//正常的回应
func SendResponseData(w http.ResponseWriter,body ResponseBody)  {
	w.WriteHeader(body.RespInfo.RespCode)
	respData, err:= json.Marshal(body)
	if err!=nil{
		SendErrorResponse(w,InternalError)
		return
	}
	_, err = w.Write(respData)
	if err!=nil{
		SendErrorResponse(w,InternalError)
		return
	}
}
