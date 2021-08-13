package dao

import (
	"fmt"
	"日报管理/model"
	"日报管理/utils"
)

func GetUserByUserNameAddPassword(UserName,Password string)*model.User{
	sql:="select * from users where username = ? and password = ?"
	row:=utils.Db.QueryRow(sql,UserName,Password)
	user:=&model.User{}
	err:=row.Scan(&user.ID,&user.UserName,&user.Password,&user.IDf,&user.IsLogin)
	if err!=nil{
		fmt.Println(err)
	}
	return user
}

func AddUser(UserName,Password,Idf string){
	sql:="insert into users (username, password, idf,IsLogin) VALUES (?,?,?,?)"
	utils.Db.Exec(sql,UserName,Password,Idf,1)
}
func GetUserByUserID(ID string)*model.User{
	sql:="select * from users where id = ?"
	row:=utils.Db.QueryRow(sql,ID)
	user:=&model.User{}
	row.Scan(&user.ID,&user.UserName,&user.Password,&user.IDf,&user.IsLogin)
	return user
}
func GetUserByUserName(UserName string)*model.User{
	sql:="select * from users where username = ?"
	row:=utils.Db.QueryRow(sql,UserName)
	user:=&model.User{}
	row.Scan(&user.ID,&user.UserName,&user.Password,&user.IDf,&user.IsLogin)
	return user
}
func UpdateUser(user *model.User){
	sql:="update users set username = ? , password = ? ,IsLogin = ? where username = ?"
	utils.Db.Exec(sql,user.UserName,user.Password,user.IsLogin,user.UserName)
}
func GetAllUserName()[]string{
	sql:="select username from users "
	rows,err:=utils.Db.Query(sql)
	if err!=nil{
		fmt.Println("GetAllUserName:",err)
	}
	var Names []string
	for rows.Next(){
		var name string
		err=rows.Scan(&name)
		if err!=nil{
			fmt.Println("GetAllUserName:",err)
		}
		Names=append(Names,name)
	}
	return Names
}


