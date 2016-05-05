package mmodel

import (
	"fmt"
	"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mcore/mtag"
	"github.com/mabetle/mgo/govendors/uuid"
	"reflect"
)

// GetModelTableName
// Model Should has TableName() method.
// If TableName() function not present, cast model name
// to table name.
func GetModelTableName(model interface{}) (result string) {
	// reflect model method TableName() error, try TypeName
	defer func() {
		err := recover()
		if err != nil {
			//typ := reflect.TypeOf(model)
			//typeName := fmt.Sprintf("%v", typ)
			typeName := mcore.GetTypeName(model)
			//type name to table name
			result = mcore.ToTableName(typeName)
			logger.Debugf("Model not implements TableName() function. Cast model name to table name. Model:%s, Table: %s", typeName, result)
		}
	}()

	in := make([]reflect.Value, 0)
	method := reflect.ValueOf(model).MethodByName("TableName")
	r := method.Call(in)
	result = fmt.Sprintf("%v", r[0])
	return
}

// AddUuid
func AddModelUuid(model interface{}) interface{} {
	id := uuid.New()
	logger.Tracef("Add new UUID: %v", id)
	return mcore.SetFieldValue(model, "Id", id)
}

// GetModelId
func GetModelId(model interface{}) string {
	return mcore.GetFieldValue(model, "Id")
}

//
func PrintModel(model interface{}) {
	fmt.Printf("%+v\n", model)
}

// PrintModelWithLabel
func PrintModelWithLabel(model interface{}, locale string) {
	fields := mcore.GetFields(model)
	for _, field := range fields {
		l := mtag.GetLocaleLabel(model, field, locale)
		v := mcore.GetFieldValue(model, field)
		fmt.Printf("%s:%v\n", l, v)
	}
}
