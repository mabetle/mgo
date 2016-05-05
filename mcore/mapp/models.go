package mapp

import (
	"fmt"
	"github.com/mabetle/mcore"
	"github.com/mabetle/mdb"
	"strings"
)

var Models = make(map[string]interface{})

func RegModels(models ...interface{}) {
	for _, model := range models {
		pkg := mcore.GetPkgPath(model)
		typ := mcore.GetTypeName(model)
		key := fmt.Sprintf("%s-%s", pkg, typ)
		Models[key] = model
	}
}

func migrate(xorm *mdb.Xorm, k string, v interface{}) {
	logger.Info("migrate model:", k)
	xorm.Migrate(v)
}

func Migrate(xorm *mdb.Xorm) {
	for k, v := range Models {
		migrate(xorm, k, v)
	}
}

func MigratePackage(xorm *mdb.Xorm, pkg string) {
	for k, v := range Models {
		if strings.HasPrefix(k, pkg) {
			migrate(xorm, k, v)
		}
	}
}

func MigrateModel(xorm *mdb.Xorm, name string) {
	for k, v := range Models {
		if strings.HasSuffix(k, name) {
			migrate(xorm, k, v)
		}
	}
}
