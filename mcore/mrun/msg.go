package mrun

import (
	"github.com/mabetle/mgo/mmsg"
)

// load i18n message for applicatin.
// Specific a ini file is not convenient when you just want to provide a runable file for
// terminal user.
func init() {
	mmsg.PutMsg("en", "msg-input-which-run", "Input which runner or func to run:")
	mmsg.PutMsg("zh", "msg-input-which-run", "输入要运行的接口或方法:")
}
