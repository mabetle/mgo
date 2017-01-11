package render

import (
	"io"

	"github.com/mabetle/mgo/mcore/mprint"
)

var NewPrintRender = &PrintRender{}

type PrintRender struct {
}

func (r PrintRender) Rend(value interface{}) {
	PrintRend(value)
}

func PrintRend(value interface{}) {
	mprint.Print(value)
}

func FprintRend(out io.Writer, value interface{}) {
	mprint.Fprint(out, value)
}
