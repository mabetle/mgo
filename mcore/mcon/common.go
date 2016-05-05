package mcon

import (
	"fmt"
	"github.com/mabetle/mgo/mcore/mterm"
)

// trace debug info warn error
type ColorPrinter interface {
	PrintBlack(s string) //trace
	PrintWhite(s string) //trace

	PrintRed(s string)    //error
	PrintYellow(s string) //warn
	PrintGreen(s string)  //info

	PrintBlue(s string)    //blue
	PrintMagenta(s string) //Magenta
	PrintCyan(s string)    //debug
}

var (
	c  = NewConsole()
	xc = NewXtermConsole()
)

func GetColorPrinter() ColorPrinter {
	if mterm.IsXterm() {
		return xc
	}
	return c
}

func PrintBlack(s string) { GetColorPrinter().PrintBlack(s) }
func PrintWhite(s string) { GetColorPrinter().PrintWhite(s) }

func PrintRed(s string)    { GetColorPrinter().PrintRed(s) }
func PrintYellow(s string) { GetColorPrinter().PrintYellow(s) }
func PrintGreen(s string)  { GetColorPrinter().PrintGreen(s) }

func PrintBlue(s string)    { GetColorPrinter().PrintBlue(s) }
func PrintMagenta(s string) { GetColorPrinter().PrintMagenta(s) }
func PrintCyan(s string)    { GetColorPrinter().PrintCyan(s) }

func PrintfBlack(f string, args ...interface{}) { PrintBlack(fmt.Sprintf(f, args...)) }
func PrintfWhite(f string, args ...interface{}) { PrintWhite(fmt.Sprintf(f, args...)) }

func PrintfRed(f string, args ...interface{})    { PrintRed(fmt.Sprintf(f, args...)) }
func PrintfYellow(f string, args ...interface{}) { PrintYellow(fmt.Sprintf(f, args...)) }
func PrintfGreen(f string, args ...interface{})  { PrintGreen(fmt.Sprintf(f, args...)) }

// PrintfBlue print msg in blue color
func PrintfBlue(f string, args ...interface{}) { PrintBlue(fmt.Sprintf(f, args...)) }

func PrintfMagenta(f string, args ...interface{}) { PrintMagenta(fmt.Sprintf(f, args...)) }
func PrintfCyan(f string, args ...interface{})    { PrintCyan(fmt.Sprintf(f, args...)) }
