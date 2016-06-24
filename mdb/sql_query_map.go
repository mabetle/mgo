package mdb

import "github.com/mabetle/mgo/mcore"

// QueryForMaps query for maps
// rowMaps holds rows key and value
// columns holds include columns
// good ways to store query result.
// because map has no sequence concepts, so columns is needed when access
// columns.
func (s Sql) QueryForMaps(sql string, args ...interface{}) (rowMaps []map[string]interface{}, columns []string, err error) {
	rows, errQ := s.Query(sql, args...)
	if errQ != nil {
		err = errQ
		return
	}
	defer rows.Close()
	columns, err = rows.Columns()
	if err != nil {
		return
	}
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scanArgs...)
		rowMap := make(map[string]interface{})
		for i, value := range values {
			sv := GetString(value)
			if fv, err := mcore.NewString(sv).ToFloat64(); err == nil {
				rowMap[columns[i]] = fv
				continue
			}
			rowMap[columns[i]] = sv
		}
		rowMaps = append(rowMaps, rowMap)
	}
	return
}
