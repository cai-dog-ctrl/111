package dao

import (
	"fmt"
	"日报管理/model"
	"日报管理/utils"
)

func AddSession(session *model.Session){
	sqlStr:="insert into sessions(session_id, username, user_id) values (?,?,?)"
	_,err:=utils.Db.Exec(sqlStr,session.ID,session.UserName,session.UserID)
	if err!=nil{
		fmt.Println("AddSession:",err)
	}
}
func GetSession(cookieValue string)*model.Session{
	sqlstr:="select session_id ,username,user_id from sessions where session_id = ?"
	row:=utils.Db.QueryRow(sqlstr,cookieValue)
	session:=&model.Session{}
	row.Scan(&session.ID,&session.UserName,&session.UserID)
	return session
}
func DeleteSession(SessionID string){
	sql:="delete from sessions where session_id = ?"
	utils.Db.Exec(sql,SessionID)
}

func GetUserIdfByUserID(UserID string)int{
	sql:="select idf from sessions,users where sessions.user_id=users.id and sessions.user_id = ?"
	row:=utils.Db.QueryRow(sql,UserID)
	var Idf int
	row.Scan(&Idf)
	return Idf
}

