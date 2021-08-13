package controller

import (
	"html/template"
	"net/http"
	"日报管理/sever"
)

func GetPageAllReport(w http.ResponseWriter,r *http.Request){

	page:=sever.GetPageAllReport(r.FormValue("pageNo"),r.FormValue("Time"),r.FormValue("UserName"))
	t:=template.Must(template.ParseFiles("pages/html/查看.html"))
	t.Execute(w,page)
}

func SubMitTask(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else{
		sever.AddReport(sess.UserName,r.FormValue("Summary"),r.FormValue("Content"),r.FormValue("Plan"),r.FormValue("Algorithm"))
		GetPageAllReport(w,r)
	}

}

