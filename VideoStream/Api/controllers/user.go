package controllers

import (
	"VideoStream/Api/response"
	"VideoStream/Api/service"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

//处理用户注册的请求
// /user/
func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {

	//读取post提交过来的数据
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Error("read post data failed...-->",err.Error())
		response.SendErrorResponse(w,response.PostParsedError)
		return
	}
	logs.Debug("receive the post data:",string(res))
	userForm:=&service.UserForm{}
	if err := json.Unmarshal(res, userForm);err!=nil{
		logs.Error("json parsed the byte data to userForm failed--->",err.Error())
		response.SendErrorResponse(w,response.InternalError)
		return
	}
	logs.Debug("register user info:",fmt.Sprintf("%#v",userForm))
	//交由serveice处理
	err = service.RegisterUser(userForm)
	if err!=nil{
		response.SendErrorResponse(w,response.ErrorResponseInfo{
				StateCode:"1101",
				ResponseInfo:response.ResponseInfo{
					RespCode:http.StatusServiceUnavailable,
					Msg:err.Error(),
					Action:"/user/",},
		})
		return
	}
	response.SendNormalResponse(w,response.ResponseInfo{
		RespCode:http.StatusOK,
		Msg:"注册用户成功",
		Action:"/user/",
	})
}
//处理用户登录的请求
// /user/:username  post
func LoginIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	//username:=ps.ByName("username")
	resData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.SendErrorResponse(w,response.PostParsedError)
		return
	}
	logs.Debug("recieve the request body data--->",string(resData))
	userInfo:=service.LoginInfo{}
	if err = json.Unmarshal(resData, &userInfo);err!=nil{
		logs.Error("this is json parsed error--->",err.Error())
		response.SendErrorResponse(w,response.InternalError)
		return
	}
	user, err := service.CheckLogin(userInfo.Username,userInfo.Password)
	if err != nil {
		response.SendErrorResponse(w,response.ErrorResponseInfo{
			StateCode:"1102",
			ResponseInfo:response.ResponseInfo{
				RespCode:http.StatusNotFound,
				Msg:"用户或者密码不对",
				Action:"/user/",
			},
		})
		return
	}
	//保存用户信息
	logs.Debug("login user info:",fmt.Sprintf("%#v",user))
	//保存到session
	response.SendNormalResponse(w,response.ResponseInfo{
		Msg:"登录成功",
		RespCode:http.StatusOK,
		Action:"/user/",
	})
}
//获取用户信息
// /user/:username   delete
func GetUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	_, err := io.WriteString(w, "this is handling the Get User Info request")
	if err != nil {
		logs.Error("write to brower data failed!",err.Error())
	}
}
//处理用户注销的请求
// /user/:username
func LoginOut(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	_, err := io.WriteString(w, "this is handling user logout request")
	if err != nil {
		logs.Error("write to brower data failed!",err.Error())
	}
}


