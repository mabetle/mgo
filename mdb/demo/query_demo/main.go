package main

import "mabetle/libs/hubs"

var (
	sql = hubs.GetDemoSql()
)

func main() {
	sql.NewQuery().Table("demo_table").Print()

	sql.NewQuery().Table("demo_table").Columns("DemoName").Where("DemoAge > 0 ").Order("DemoName").Print()

}
