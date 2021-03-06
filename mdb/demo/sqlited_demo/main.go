package main


import (
	"github.com/mabetle/mgo/mdb/sqlited"
	"github.com/mabetle/mgo/mlog"
	"fmt"
)

var(
	logger = mlog.GetLogger("main")
)

func main() {
	sql, err:= sqlited.NewSql("/tmp/demo.db")
	if logger.CheckError(err){
		return	
	}
	n,_:=sql.QueryForInt("select count(*) from demo_bean")
	fmt.Printf("%v\n", n)
	sql.PrintTable("demo_bean")
}
