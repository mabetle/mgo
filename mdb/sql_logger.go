package mdb

import (
	"fmt"

	"github.com/mabetle/mgo/mcore"
)

// SqlLogger Log
func (s Sql) Log(query string, args ...interface{}) {
	if !s.ShowSql {
		return
	}
	fmt.Printf("Host:%s Database:%s\n", s.Host, s.Schema)
	fmt.Printf("Sql :%s\n", query)
	if len(args) > 0 {
		argStr := mcore.SepJoin(",", args...)
		fmt.Printf("Args:%s\n", argStr)
	}
}

func (s *Sql) SetShowSql(b bool) {
	s.ShowSql = b
}
