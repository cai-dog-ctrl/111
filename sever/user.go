package sever

import (
	"strconv"
	"日报管理/dao"
	"日报管理/model"
)

func Login(UserName ,Password string)*model.User{
	user:=dao.GetUserByUserNameAddPassword(UserName,Password)
	if user!=nil{

		return user
	}else {
		return nil
	}
}
func Regist(UserName,Password,Idf string){
	dao.AddUser(UserName,Password,Idf)
}


func GetUserByUserID(ID string)*model.User{
	return dao.GetUserByUserID(ID)
}

func GetUserByUserName(Name string)*model.User{
	return dao.GetUserByUserName(Name)
}
func UpdateUser(UserName ,Password ,IsLogin string){
	iIsLogin,_:=strconv.ParseInt(IsLogin,10,0)
	user:=&model.User{
		UserName: UserName,
		Password: Password,
		IsLogin: int(iIsLogin),
	}
	dao.UpdateUser(user)
}
func GetAllUserName()[]string{
	return dao.GetAllUserName()
}


