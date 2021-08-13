package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"日报管理/model"
	"日报管理/utils"
)

func AddReport(report *model.Report)bool{
	sql:="insert into report (username, plan, content, time, summary,algorithm) VALUES (?,?,?,?,?,?)"
	_,err:=utils.Db.Exec(sql,report.UserName,report.Plan,report.Content,report.Time,report.Summary,report.Algorithm)
	if err!=nil{
		return false
	}
	return true
}

func GetAllReport()[]*model.Report{
	sql:="select * from report "
	var Reports []*model.Report
	rows,err:=utils.Db.Query(sql)
	if err!=nil{
		fmt.Println("GetAllReport:",err)
		return nil
	}
	for rows.Next(){
		report:=&model.Report{}
		err=rows.Scan(&report.ID,&report.UserName,&report.Plan,&report.Content,&report.Time,&report.Summary,&report.Algorithm)
		if err!=nil{
			fmt.Println("GetAllReport:",err)
			return nil
		}
		Reports=append(Reports,report)
	}
	return Reports
}
func GetPageAllReport(pageNo string,Time string,UserName string)*model.Pages{
	var month string
	var day string
	if Time!=""{
		t:=strings.Split(Time,"-")
		month=t[1]
		day=t[2]
	}
	var totalInt int64
	var PageSize int64=10
	var PageNum int64//总共的页数
	ipageNo,_:=strconv.ParseInt(pageNo,10,0)
	sqlstr1:="select count(*) from report where day(time)= ? and month(time) = ?"//只有行个数
	sqlstr2:="select count(*) from report where username = ? "//只有列个数
	sqlstr3:="select count(*) from report where day(time)= ? and month(time) = ? and  username = ? "//有行有列个数
	sqlstr4:="select count(*) from report "
	sqlStr1:="select * from report where  day(time)= ? and month(time) = ? limit ?,?"//只有行
	sqlStr2:="select * from report where username = ? limit ?,?"//只有列
	sqlStr3:="select * from report where day(time)= ? and month(time) = ? and  username = ? limit ?,?"//有行有列
	sqlStr4:="select * from report limit ?,?"
	var row *sql.Row
	var rows *sql.Rows
	var err error
	if Time!=""&&UserName==""{
		row=utils.Db.QueryRow(sqlstr1,day,month)
		row.Scan(&totalInt)
		if totalInt%PageSize==0{
			PageNum=totalInt/PageSize
		}else{
			PageNum=totalInt/PageSize+1
		}
		rows,err=utils.Db.Query(sqlStr1,day,month,(ipageNo-1)*PageSize,PageSize)
		if err!=nil{
			fmt.Println("GetPageAllReport:",err)
			return nil
		}
	}
	if Time==""&&UserName!=""{
		row=utils.Db.QueryRow(sqlstr2,UserName)
		row.Scan(&totalInt)
		if totalInt%PageSize==0{
			PageNum=totalInt/PageSize
		}else{
			PageNum=totalInt/PageSize+1
		}
		rows,err=utils.Db.Query(sqlStr2,UserName,(ipageNo-1)*PageSize,PageSize)
		if err!=nil{
			fmt.Println("GetPageAllReport:",err)
			return nil
		}
	}
	if Time!=""&&UserName!=""{
		row=utils.Db.QueryRow(sqlstr3,day,month,UserName)
		row.Scan(&totalInt)
		if totalInt%PageSize==0{
			PageNum=totalInt/PageSize
		}else{
			PageNum=totalInt/PageSize+1
		}
		rows,err=utils.Db.Query(sqlStr3,day,month,UserName,(ipageNo-1)*PageSize,PageSize)
		if err!=nil{
			fmt.Println("GetPageAllReport:",err)
			return nil
		}

	}
	if Time==""&&UserName==""{
		row=utils.Db.QueryRow(sqlstr4)
		row.Scan(&totalInt)
		if totalInt%PageSize==0{
			PageNum=totalInt/PageSize
		}else{
			PageNum=totalInt/PageSize+1
		}
		rows,err=utils.Db.Query(sqlStr4,(ipageNo-1)*PageSize,PageSize)
		if err!=nil{
			fmt.Println("GetPageAllReport:",err)
			return nil
		}
	}
	var Reports []*model.Report
	for rows.Next() {
		report:=&model.Report{}
		rows.Scan(&report.ID,&report.UserName,&report.Plan,&report.Content,&report.Time,&report.Summary,&report.Algorithm)
		Reports=append(Reports,report)
	}
	Page:=&model.Pages{
		Reports: Reports,
		PageNo: ipageNo,
		PageSize: PageSize,
		TotalPageNo:PageNum,
		TotalRecord: totalInt,
		Time: Time,
		UserName: UserName,
	}
	return Page

}
func GetReportByID(ID string)*model.Report{
	sql:="select * from report where id = ?"
	row:=utils.Db.QueryRow(sql,ID)
	report:=&model.Report{}
	row.Scan(&report.ID,&report.UserName,&report.Plan,&report.Content,&report.Time,&report.Summary,&report.Algorithm)
	return report
}

func UpdateReport(report * model.Report){
	sql:="update report set plan = ? ,content = ? , time = ? ,summary = ? ,algorithm = ? where id = ?"
	utils.Db.Exec(sql,report.Plan,report.Content,report.Time,report.Summary,report.Algorithm,report.ID)
}
func GetGroupReport(UserName,Time string)[]*model.Report{
	var month string
	var day string
	if Time!=""{
		t:=strings.Split(Time,"-")
		month=t[1]
		day=t[2]
	}
	sqlStr1:="select * from report where  day(time)= ? and month(time) = ? "//只有行
	sqlStr2:="select * from report where username = ? "//只有列
	sqlStr3:="select * from report where day(time)= ? and month(time) = ? and  username = ? "//有行有列
	sqlStr4:="select * from report "
	var rows *sql.Rows
	var err error
	if Time!=""&&UserName==""{
		rows,err=utils.Db.Query(sqlStr1,day,month)
		if err!=nil{
			fmt.Println("GetGroupReport:",err)
			return nil
		}
	}
	if Time==""&&UserName!=""{
		rows,err=utils.Db.Query(sqlStr2,UserName)
		if err!=nil{
			fmt.Println("GetGroupReport:",err)
			return nil
		}
	}
	if Time!=""&&UserName!=""{
		rows,err=utils.Db.Query(sqlStr3,day,month,UserName)
		if err!=nil{
			fmt.Println("GetGroupReport:",err)
			return nil
		}
	}
	if Time==""&&UserName==""{

		rows,err=utils.Db.Query(sqlStr4)
		if err!=nil{
			fmt.Println("GetGroupReport:",err)
			return nil
		}
	}
	var Reports []*model.Report
	for rows.Next() {
		report:=&model.Report{}
		rows.Scan(&report.ID,&report.UserName,&report.Plan,&report.Content,&report.Time,&report.Summary,&report.Algorithm)
		Reports=append(Reports,report)
	}

	return Reports
}

func DeleteReport(ReportID string){
	sql:="delete from report where id = ?"
	utils.Db.Exec(sql,ReportID)
}
