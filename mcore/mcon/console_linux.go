package mcon

import (
	"github.com/mabetle/mgo/mcore/mterm"
)

// implments for ColorPriter
func (c Console) PrintBlack(s string) { mterm.PrintBlack(s) }
func (c Console) PrintWhite(s string) { mterm.PrintWhite(s) }

func (c Console) PrintRed(s string)    { mterm.PrintRed(s) }
func (c Console) PrintYellow(s string) { mterm.PrintYellow(s) }
func (c Console) PrintGreen(s string)  { mterm.PrintGreen(s) }

func (c Console) PrintBlue(s string)    { mterm.PrintBlue(s) }
func (c Console) PrintMagenta(s string) { mterm.PrintMagenta(s) }
func (c Console) PrintCyan(s string)    { mterm.PrintCyan(s) }

func (c Console) ColorPrint(s string, i int) {
	mterm.ColorPrint(s, i)
}
