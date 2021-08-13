package main

import (
	"net/http"
	"日报管理/controller"
)

func main(){
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("pages"))))
	http.HandleFunc("/login",controller.Login)
	http.HandleFunc("/regist",controller.Register)
	http.HandleFunc("/subMitTask",controller.SubMitTask)
	http.HandleFunc("/main",controller.GetPageAllReport)
	http.HandleFunc("/toChaKan",controller.ToChanKan)
	http.HandleFunc("/toReg",controller.ToReg)
	http.HandleFunc("/toLogin",controller.ToLogin)
	http.HandleFunc("/showMyReport",controller.ShowMyReport)
	http.HandleFunc("/chooseManger",controller.ChooseManger)
	http.HandleFunc("/changeMyReport",controller.ChangeMyReport)
	http.HandleFunc("/changeReport",controller.ChangeReport)
	http.HandleFunc("/addGroup",controller.AddGroup)
	http.HandleFunc("/mangerGroup",controller.GetGroupMemberByGroupName)
	http.HandleFunc("/toUpdateReportByManger",controller.ToUpdateReportByManger)
	http.HandleFunc("/updateReportByManger",controller.UpdateReportByManger)
	http.HandleFunc("/deleteReport",controller.DeleteReport)
	http.HandleFunc("/toUpdateUser",controller.ToUpdateUserName)
	http.HandleFunc("/UpdateUser",controller.UpdateUserName)
	http.HandleFunc("/joinGroup",controller.JoinGroup)
	http.HandleFunc("/toUserInformation",controller.GetUserInformation)
	http.HandleFunc("/logout",controller.Logout)
	http.HandleFunc("/deleteFromGroup",controller.DeleteUserFromGroup)
	http.HandleFunc("/addGroupLeaderOrManger",controller.AddGroupLeaderOrManger)
	http.ListenAndServe(":8080",nil)
}
