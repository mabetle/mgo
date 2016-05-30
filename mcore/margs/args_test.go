package margs

import (
	"fmt"
	"testing"

	"github.com/mabetle/mgo/mcore/mtest"
)

func TestArgs(t *testing.T) {
	mtest.RegTest(t)
	renderMap := map[string]interface{}{
		"name":          "demo",
		"exist":         "false",
		"age":           "20",
		"currentLocale": "zh_CN",
	}
	wa := NewArgs(renderMap, "demo=demo")
	mtest.AssertEqual(wa.GetString("demo", ""), "demo")
	mtest.AssertEqual(wa.GetInt64("age", 0), int64(20))
	fmt.Print(wa)
}
