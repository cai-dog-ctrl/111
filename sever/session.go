package sever

import (
	"strconv"
	"日报管理/dao"
	"日报管理/model"
)

func DeleteSession(SessionID string){
	dao.DeleteSession(SessionID)
}
func GetSession(cookieValue string)*model.Session{
	return dao.GetSession(cookieValue)
}
func GetUserIdfByUserID(UserID int)int{
	iUserID:=strconv.FormatInt(int64(UserID),10)
	return dao.GetUserIdfByUserID(iUserID)
}
