package sever

import (
	"math"
	"strconv"
	"日报管理/dao"
	"日报管理/model"
)

func AddGroup(GroupName string){
	dao.CreatTableForGrouper(GroupName)
}
func AddInvitationCode(InvitationCode,GroupName  string){
	dao.AddInvitationCode(InvitationCode,GroupName)
}
func AddGroupLeaderToTable(UserID int,UserName,invitationCode string){
	dao.AddGroupLeaderToTable(UserID,UserName,invitationCode)
}
func AddPeopleToGroupTable(GroupName,UserName,invitationCode string){
	user:=dao.GetUserByUserName(UserName)
	InvitationCode:=dao.GetInvitationCodeByGroupName(GroupName)
	dao.AddPeopleToGroupTable(GroupName,UserName,invitationCode,user.IDf,InvitationCode.ID)
}
func GetGroupByGroupLeaderID(GroupLeaderID int)[]*model.InvitationCodes{
	sGroupLeaderID:=strconv.FormatInt(int64(GroupLeaderID),10)
	return dao.GetGroupByGroupLeaderID(sGroupLeaderID)
}
func GetGroupMemberByGroupName(GroupName string)[]string{
	return dao.GetGroupMemberByGroupName(GroupName)
}

func GetGroupReport(Names []string,pageNo,UserName,Time string)*model.Pages{
	Reports:=dao.GetGroupReport(UserName,Time)
	Name:=make(map[string]bool)
	for _,v:=range Names{
		Name[v]=true
	}
	var reports []*model.Report
	for _,v:=range Reports{
		if Name[v.UserName]{
			reports=append(reports,v)
		}
	}
	var totalInt int64
	totalInt=int64(len(reports))
	var PageSize int64=10
	var PageNum int64//总共的页数
	if totalInt%PageSize==0{
		PageNum=totalInt/PageSize
	}else{
		PageNum=totalInt/PageSize+1
	}
	ipageNo,_:=strconv.ParseInt(pageNo,10,0)
	S,E:=SlicePage(int(ipageNo),int(PageSize),int(totalInt))
	Page:=&model.Pages{
		Reports: reports[S:E:len(reports)],
		PageNo: ipageNo,
		PageSize: PageSize,
		TotalPageNo:PageNum,
		TotalRecord: totalInt,
	}
	return Page
}
func SlicePage(page, pageSize, nums int) (sliceStart, sliceEnd int) {
	if page < 0 {
		page = 1
	}

	if pageSize < 0 {
		pageSize = 20
	}

	if pageSize > nums {
		return 0, nums
	}

	// 总页数
	pageCount := int(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd
}

func GetSubmitReportUser(Users []string)[]string{
	Names:=dao.GetSubmitUserToday()
	var users []string
	Name:=make(map[string]bool)
	for _,v:=range Users{
		Name[v]=true
	}
	for _,v:=range Names{
		if Name[v]{
			users=append(users,v)
		}
	}
	return users
}
func GetNoSubmitReportUser(Names ,Submit []string)[]string{
	SubmitName:=make(map[string]bool)
	for _,v:=range Submit{
		SubmitName[v]=true
	}
	var NoSubmitName []string
	for _,v:=range Names{
		if SubmitName[v]{
			continue
		}else{
			NoSubmitName=append(NoSubmitName,v)
		}
	}
	return NoSubmitName
}
func GetInvitationByGroupName(Name string)(invitationCode string){
	return dao.GetInvitationByGroupName(Name)
}
func GetGroupByInvitationCode(InvitationCode string)*model.InvitationCodes{
	return dao.GetGroupByInvitationCode(InvitationCode)
}

func AddUserToGroup(invitationCode,UserName,GroupName string,Idf ,GroupID int){
	dao.AddUserToGroup(invitationCode,UserName,GroupName,Idf,GroupID)
}
func DeleteUserFromGroup(UserName , GroupName string){
	dao.DeleteUserFromGroup(UserName,GroupName)
}
