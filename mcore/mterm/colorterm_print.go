package mterm

import (
	"fmt"
)

// Print funcs
func PrintBlack(s string) { fmt.Print(Black(s)) }
func PrintWhite(s string) { fmt.Print(White(s)) }

func PrintRed(s string)    { fmt.Print(Red(s)) }
func PrintGreen(s string)  { fmt.Print(Green(s)) }
func PrintYellow(s string) { fmt.Print(Yellow(s)) }

func PrintBlue(s string)    { fmt.Print(Blue(s)) }
func PrintMagenta(s string) { fmt.Print(Magenta(s)) }
func PrintCyan(s string)    { fmt.Print(Cyan(s)) }

// Printf funcs
func PrintfWhite(format string, args ...interface{}) { PrintWhite(fmt.Sprintf(format, args...)) }
func PrintfBlack(format string, args ...interface{}) { PrintBlack(fmt.Sprintf(format, args...)) }

func PrintfRed(format string, args ...interface{})    { PrintRed(fmt.Sprintf(format, args...)) }
func PrintfGreen(format string, args ...interface{})  { PrintGreen(fmt.Sprintf(format, args...)) }
func PrintfYellow(format string, args ...interface{}) { PrintYellow(fmt.Sprintf(format, args...)) }

func PrintfBlue(format string, args ...interface{})    { PrintBlue(fmt.Sprintf(format, args...)) }
func PrintfMagenta(format string, args ...interface{}) { PrintMagenta(fmt.Sprintf(format, args...)) }
func PrintfCyan(format string, args ...interface{})    { PrintCyan(fmt.Sprintf(format, args...)) }

func ColorPrint(s string, i int) { fmt.Print(ColorText(i, s)) }
