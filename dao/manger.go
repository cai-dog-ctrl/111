package dao

import (
	"fmt"
	"time"
	"日报管理/model"
	"日报管理/utils"
)

func CreatTableForGrouper(GroupName string){
	sql:="create table "+GroupName+" (ID int primary key auto_increment,UserName varchar(100) not null ,invitationCode varchar(100) not null ,Idf int not null ,GroupID int not null )"
	_,err:=utils.Db.Exec(sql)
	if err!=nil{
		fmt.Println("CreatTableForGrouper:",err)
	}
}
func AddInvitationCode(InvitationCode,GroupName string){
	sql:="insert into invitationcode (invitation_Code, group_Name) VALUES (?,?)"
	_,err:=utils.Db.Exec(sql,InvitationCode,GroupName)
	if err!=nil{
		fmt.Println("AddInvitationCode:",err)
	}
}
func AddGroupLeaderToTable(UserID int,UserName,invitationCode string){
	sql:="insert into team_and_teammanger (groupLeaderID, groupLeaderName, invitation_Code) VALUES (?,?,?)"
	_,err:=utils.Db.Exec(sql,UserID,UserName,invitationCode)
	if err!=nil{
		fmt.Println("AddGroupLeaderToTable:",err)
	}
}
func GetInvitationCodeByGroupName(GroupName string)*model.InvitationCodes{
	sql:="select * from invitationcode where group_Name = ?"
	row:=utils.Db.QueryRow(sql,GroupName)
	InvitationCode:=&model.InvitationCodes{}
	row.Scan(&InvitationCode.ID,&InvitationCode.InvitationCode,&InvitationCode.GroupName)
	return InvitationCode
}
func AddPeopleToGroupTable(GroupName,UserName,invitationCode string,Idf ,GroupID int){
	sql:="insert into " +GroupName+ " (UserName,invitationCode,Idf,GroupID) values (?,?,?,?)"
	_,err:=utils.Db.Exec(sql,UserName,invitationCode,Idf,GroupID)
	if err!=nil{
		fmt.Println("AddPeopleToGroupTable:",err)
	}
}

func GetGroupByGroupLeaderID(GroupLeaderID string)[]*model.InvitationCodes{
	sql:="select invitationcode.ID,invitationcode.group_Name from invitationcode,team_and_teammanger where invitationcode.invitation_Code=team_and_teammanger.invitation_Code and  team_and_teammanger.groupLeaderID = ?"
	rows,err:=utils.Db.Query(sql,GroupLeaderID)
	if err!=nil{
		fmt.Println("GetGroupMemberByGroupName:",err)
	}
	var Invitations []*model.InvitationCodes
	for rows.Next(){
		invitation:=&model.InvitationCodes{}
		err=rows.Scan(&invitation.ID,&invitation.GroupName)
		if err!=nil{
			fmt.Println("GetGroupMemberByGroupName:",err)
		}
		Invitations=append(Invitations,invitation)
	}
	return Invitations
}
func GetGroupMemberByGroupName(GroupName string)[]string{
	sql:="select UserName from juan"
	var Names []string
	rows,_:=utils.Db.Query(sql)
	for rows.Next(){
		var name string
		rows.Scan(&name)
		Names=append(Names,name)
	}
	return Names
}

func GetSubmitUserToday()[]string{
	day:=time.Now().Day()
	month:=time.Now().Month()
	year:=time.Now().Year()
	sql:="select username from report where year(time) = ? and month(time) = ? and day(time) = ?"
	row,err:=utils.Db.Query(sql,year,month,day)
	if err!=nil{
		fmt.Println("GetSubmitUserToday:",err)
		return nil
	}
	var Name []string
	for row.Next(){
		var name string
		row.Scan(&name)
		Name=append(Name,name)
	}
	return Name
}

func GetInvitationByGroupName(Name string)(invitationCode string){
	sql:="select  invitation_Code from invitationcode where group_Name = ?"
	row:=utils.Db.QueryRow(sql,Name)
	row.Scan(&invitationCode)
	return invitationCode
}
func GetGroupByInvitationCode(InvitationCode string)*model.InvitationCodes{
	sql:="select * from invitationcode where invitation_Code = ?"
	row:=utils.Db.QueryRow(sql,InvitationCode)
	invitationCode:=&model.InvitationCodes{}
	err:=row.Scan(&invitationCode.ID,&invitationCode.InvitationCode,&invitationCode.GroupName)
	if err!=nil{
		fmt.Println("GetGroupByInvitationCode:",err)
		return nil
	}
	return invitationCode
}

func AddUserToGroup(invitationCode,UserName,GroupName string,Idf ,GroupID int){
	sql:="insert into "+GroupName+" (UserName, invitationCode, Idf, GroupID) VALUES (?,?,?,?)"
	utils.Db.Exec(sql,UserName,invitationCode,Idf,GroupID)
}

func DeleteUserFromGroup(UserName , GroupName string){
	sql:="delete from "+GroupName+" where UserName = ?"
	utils.Db.Exec(sql,UserName)
}
