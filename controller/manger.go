package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"日报管理/model"
	"日报管理/sever"
	"日报管理/utils"
)

func ChooseManger(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else{
		UserID:=strconv.FormatInt(int64(sess.UserID),10)
		user:=sever.GetUserByUserID(UserID)
		if user.IDf==1{
			ToManger(w,r)
		}else if user.IDf==2{
			ShowGroupLeader(w,r)
		}else if user.IDf==3{
			ShowMyReport(w,r)
		}
	}
}
func ShowMyReport(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		PageNo:=r.FormValue("PageNo")
		page:=sever.GetPageAllReport(PageNo,"",sess.UserName)
		t:=template.Must(template.ParseFiles("pages/html/管理_个人.html"))
		t.Execute(w,page)
	}

}
func ChangeMyReport(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		ID:=r.FormValue("ID")
		report:=sever.GetReportByID(ID)
		t:=template.Must(template.ParseFiles("pages/html/changeReport.html"))
		t.Execute(w,report)
	}

}
func ChangeReport(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		sever.UpdateReport3(r.FormValue("ID"),r.FormValue("UserName"),r.FormValue("Plan"),r.FormValue("Content"),r.FormValue("Time"),r.FormValue("Summary"),r.FormValue("Algorithm"))
		GetPageAllReport(w,r)
	}

}
func ShowGroupLeader(w http.ResponseWriter,r *http.Request){

	ok,sess:=IsLogin(r)

	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else{
		invitation:=sever.GetGroupByGroupLeaderID(sess.UserID)
		t:=template.Must(template.ParseFiles("pages/html/管理_组长.html"))
		t.Execute(w,invitation)
	}

}
func AddGroup(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		GroupName:=r.FormValue("GroupName")
		invitationCode:=utils.CreateUUID()//生成邀请码
		//将邀请码加入invitationCode表中
		sever.AddInvitationCode(invitationCode,GroupName)
		//创建小组成员列表
		sever.AddGroup(GroupName)
		//把组长加入小组成员列表
		sever.AddPeopleToGroupTable(GroupName,sess.UserName,invitationCode)
		//将组长加入team_and_teammanger表中

		sever.AddGroupLeaderToTable(sess.UserID,sess.UserName,invitationCode)
		w.Write([]byte("你的群邀请码是："+invitationCode))
	}

}
func GetGroupMemberByGroupName(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		GroupName:=r.FormValue("Sub")
		PageNo:=r.FormValue("PageNo")
		UserName:=r.FormValue("UserName")
		Time:=r.FormValue("Time")
		Names:=sever.GetGroupMemberByGroupName(GroupName)
		page:=sever.GetGroupReport(Names,PageNo,UserName,Time)
		SubmitUser:=sever.GetSubmitReportUser(Names)
		NoSubmitUser:=sever.GetNoSubmitReportUser(Names,SubmitUser)
		InvitationCode:=sever.GetInvitationByGroupName(GroupName)
		Group:=&model.Group{
			Page: page,
			SubmitUser: SubmitUser,
			NoSubmitUser: NoSubmitUser,
			InvitationCode: InvitationCode,
			GroupName: GroupName,
		}
		t:=template.Must(template.ParseFiles("pages/html/GroupManger.html"))
		t.Execute(w,Group)
	}

}

func ToUpdateReportByManger(w http.ResponseWriter,r *http.Request){
	//sever.UpdateReport(r.FormValue("ID"),r.FormValue("UserName"),r.FormValue("Plan"),r.FormValue("Content"),r.FormValue("Time"),r.FormValue("Summary"),r.FormValue("Algorithm"))
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		ID:=r.FormValue("ID")
		report:=sever.GetReportByID(ID)
		t:=template.Must(template.ParseFiles("pages/html/UpdateReport.html"))
		t.Execute(w,report)
	}

}
func UpdateReportByManger(w http.ResponseWriter,r *http.Request)  {
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		sever.UpdateReport(r.FormValue("ID"),r.FormValue("UserName"),r.FormValue("Plan"),r.FormValue("Content"),r.FormValue("Time"),r.FormValue("Summary"),r.FormValue("Algorithm"))
		t:=template.Must(template.ParseFiles("pages/html/Update_sess.html"))
		t.Execute(w,"")
	}

}
func DeleteReport(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		sever.DeleteReport(r.FormValue("ID"))
		t:=template.Must(template.ParseFiles("pages/html/Update_sess.html"))
		t.Execute(w,"")
	}

}
func ToUpdateUserName(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		UserName:=r.FormValue("UserName")
		user:=sever.GetUserByUserName(UserName)
		t:=template.Must(template.ParseFiles("pages/html/updateUser.html"))
		t.Execute(w,user)
	}

}
func UpdateUserName(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		sever.UpdateUser(r.FormValue("UserName"),r.FormValue("Password"),r.FormValue("IsLogin"))
		t:=template.Must(template.ParseFiles("pages/html/updateUser.html"))
		t.Execute(w,"")
	}

}

func JoinGroup(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		invitationCode:=r.FormValue("invitationCode")
		Invitation:=sever.GetGroupByInvitationCode(invitationCode)
		sever.AddUserToGroup(invitationCode,sess.UserName,Invitation.GroupName,3,Invitation.ID)
		t:=template.Must(template.ParseFiles("pages/html/Update_sess.html"))
		t.Execute(w,"")
	}

}

func GetUserInformation(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		PageNo:=r.FormValue("PageNo")
		UseName:=r.FormValue("UserName")
		GroupName:=r.FormValue("GroupName")
		//根据用户名查报告
		page:=sever.GetPageAllReport(PageNo,"",UseName)
		UserInformation:=&model.UserInformation{
			Page: page,
			UserName: UseName,
			GroupName: GroupName,
		}
		t:=template.Must(template.ParseFiles("pages/html/个人管理.html"))
		t.Execute(w,UserInformation)
	}

}
func DeleteUserFromGroup(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		GroupName:=r.FormValue("GroupName")
		UserName := r.FormValue("UserName")
		sever.DeleteUserFromGroup(UserName,GroupName)
		t:=template.Must(template.ParseFiles("pages/html/Update_sess.html"))
		t.Execute(w,"")
	}

}
func ToManger(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		PageNo:=r.FormValue("PageNo")
		UserName:=r.FormValue("UserName")
		Time:=r.FormValue("Time")
		page:=sever.GetPageAllReport(PageNo,UserName,Time)
		Names:=sever.GetAllUserName()
		SubmitUser:=sever.GetSubmitReportUser(Names)
		NoSubmitUser:=sever.GetNoSubmitReportUser(Names,SubmitUser)
		Group:=&model.Group{
			Page: page,
			SubmitUser: SubmitUser,
			NoSubmitUser: NoSubmitUser,
		}
		t:=template.Must(template.ParseFiles("pages/html/管理_管理员.html"))
		t.Execute(w,Group)
	}

}

func AddGroupLeaderOrManger(w http.ResponseWriter,r *http.Request){
	ok,sess:=IsLogin(r)
	if !ok{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else if sever.GetUserIdfByUserID(sess.UserID)==sess.UserID{
		t:=template.Must(template.ParseFiles("pages/html/login.html"))
		t.Execute(w,"")
	}else {
		UserName:=r.FormValue("UserName")
		Password:=r.FormValue("Password")
		Idf:=r.FormValue("Idf")
		sever.Regist(UserName,Password,Idf)
		t:=template.Must(template.ParseFiles("pages/html/Update_sess.html"))
		t.Execute(w,"")
	}

}