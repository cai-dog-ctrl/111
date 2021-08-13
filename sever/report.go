package sever

import (
	"strconv"
	"strings"
	"time"
	"日报管理/dao"
	"日报管理/model"
)
func ChangeTime(Time string)string{
	var times []byte
	for _,v:=range Time{
		if v!=84{
			times=append(times,byte(v))
		}else{
			times=append(times,32)
		}
	}
	t1:=time.Now().Second()
	t:=strconv.FormatInt(int64(t1),10)
	return string(times)+":"+t
}
func GetPageAllReport(pageNo ,Time ,UserName string)*model.Pages{

	if pageNo==""{
		pageNo="1"
	}
	var sTime string
	if Time!=""{
		sTime=ChangeTime(Time)
	}
	return dao.GetPageAllReport(pageNo,sTime,UserName)
}
func AddReport(UserName ,Plan,Summary,Content,Algorithm string)bool{
	Time:=time.Now().Format("2006-01-02 15:04:05")
	Report:=&model.Report{
		UserName: UserName,
		Plan: Plan,
		Summary: Summary,
		Content: Content,
		Algorithm: Algorithm,
		Time: Time,
	}
	if dao.AddReport(Report){
		return true
	}
	return false
}
func GetReportByID(ID string)*model.Report{
	return dao.GetReportByID(ID)
}

func UpdateReport3(ID ,UserName,Plan,Content,Time,Summary,Algorithm string)bool{
	Year:=time.Now().Year()
	Month:=time.Now().Month()
	Day:=time.Now().Day()
	date:=strings.Split(Time,"-")
	year:=date[0]
	month:=date[1]
	day:=strings.Split(date[2]," ")[0]
	sYear,_:=strconv.ParseInt(year,10,0)
	sMonth,_:=strconv.ParseInt(month,10,0)
	sDay,_:=strconv.ParseInt(day,10,0)
	if int64(Year)==sYear&&int64(Month)==sMonth&&int64(Day)==sDay{

		sID,_:=strconv.ParseInt(ID,10,0)
		report:=&model.Report{
			ID: int(sID),
			UserName: UserName,
			Plan: Plan,
			Content: Content,
			Time: time.Now().Format("2006-01-02 15:04:05"),
			Summary: Summary,
			Algorithm: Algorithm,
		}
		dao.UpdateReport(report)
	}else{
		return false
	}
	return false
}
func UpdateReport(ID ,UserName,Plan,Content,Time,Summary,Algorithm string){
	sID,_:=strconv.ParseInt(ID,10,0)
	report:=&model.Report{
		ID: int(sID),
		UserName: UserName,
		Plan: Plan,
		Content: Content,
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Summary: Summary,
		Algorithm: Algorithm,
	}
	dao.UpdateReport(report)
}
func DeleteReport(ReportID string){
	dao.DeleteReport(ReportID)
}