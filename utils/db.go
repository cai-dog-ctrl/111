package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)
var(
	Db *sql.DB
	err error
)
func init(){
	Db,err=sql.Open("mysql", "root:lh1296643805@tcp(127.0.0.1)/go")
	if err!=nil{
		fmt.Println("OPen",err)
	}
	err=Db.Ping()
	if err!=nil{
		fmt.Println("Ping",err)
	}
	//defer Db.Close()
}