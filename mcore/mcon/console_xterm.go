package mcon

import (
	"github.com/mabetle/mcore/mterm"
)

// XtermConsole implements ColorPrinter
type XtermConsole struct {
	*mterm.ColorTerm
}

func NewXtermConsole() *XtermConsole {
	return &XtermConsole{ColorTerm: mterm.NewColorTerm()}
}
