package controller

import (
	"html/template"
	"net/http"
	"日报管理/dao"
	"日报管理/model"
	"日报管理/sever"
	"日报管理/utils"
)
func IsLogin(r *http.Request)(bool,*model.Session){
	cookie,err:=r.Cookie("user")
	if err!=nil{

	}
	if cookie!=nil{
		cookieValue:=cookie.Value
		session:=sever.GetSession(cookieValue)
		if session.UserID>0{
			return true,session
		}
	}

	return false,nil
}
func Login(w http.ResponseWriter,r *http.Request){
	user:=sever.Login(r.FormValue("UserName"),r.FormValue("Password"))
	if user!=nil&&user.IsLogin==1{
		uuid:=utils.CreateUUID()
		sess:=&model.Session{
			ID: uuid,
			UserName: user.UserName,
			UserID: user.ID,
		}
		//加入数据库
		dao.AddSession(sess)
		//创建一个Cookie
		cookie:=http.Cookie{
			Name: "user",
			Value: uuid,
			HttpOnly: true,
		}
		http.SetCookie(w,&cookie)
		t:=template.Must(template.ParseFiles("pages/html/提交.html"))
		t.Execute(w,"")
	}else{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"用户名或密码错误")
	}
}

func Register(w http.ResponseWriter,r *http.Request){
	sever.Regist(r.FormValue("UserName"),r.FormValue("Password"),"3")
	ToLogin(w,r)
}


func ToChanKan(w http.ResponseWriter,r *http.Request){
	t:=template.Must(template.ParseFiles("pages/html/提交.html"))
	t.Execute(w,"")
}
func ToLogin(w http.ResponseWriter,r *http.Request){
	t:=template.Must(template.ParseFiles("pages/html/login.html"))
	t.Execute(w,"")
}
func ToReg(w http.ResponseWriter,r *http.Request){
	t:=template.Must(template.ParseFiles("pages/html/regist.html"))
	t.Execute(w,"")
}

func Logout(w http.ResponseWriter,r *http.Request){
	cookie,_:=r.Cookie("user")
	if cookie!=nil{
		cookieValue:=cookie.Value
		dao.DeleteSession(cookieValue)
		cookie.MaxAge=-1
		http.SetCookie(w,cookie)
	}
	GetPageAllReport(w,r)
}

