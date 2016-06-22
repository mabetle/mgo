package mapp

import (
	"fmt"
	"strings"

	"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mdb"
)

var HubModels = make(map[string]interface{})

func RegModels(models ...interface{}) {
	for _, model := range models {
		pkg := mcore.GetPkgPath(model)
		typ := mcore.GetTypeName(model)
		key := fmt.Sprintf("%s-%s", pkg, typ)
		HubModels[key] = model
	}
}

func migrate(xorm *mdb.Xorm, k string, v interface{}) error {
	logger.Info("migrate model:", k)
	return xorm.Migrate(v)
}

func MigratePackage(xorm *mdb.Xorm, pkg string) *mcore.Results {
	rs := mcore.NewResults()
	for k, v := range HubModels {
		if strings.HasPrefix(k, pkg) {
			err := migrate(xorm, k, v)
			rs.Record(err, "migrate ", k)
		}
	}
	return rs
}

func MigrateModel(xorm *mdb.Xorm, name string) *mcore.Results {
	rs := mcore.NewResults()
	for k, v := range HubModels {
		if strings.HasSuffix(k, name) {
			err := migrate(xorm, k, v)
			rs.Record(err, "migrate ", k)
		}
	}
	return rs
}
