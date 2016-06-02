package ioc

import (
	"mabetle/libs/hubs"

	"github.com/mabetle/mgo/mdb"
)

var sqlDB *mdb.Sql

func GetSql() *mdb.Sql {
	if sqlDB == nil {
		sqlDB = hubs.GetDemoSql()
	}
	return sqlDB
}
