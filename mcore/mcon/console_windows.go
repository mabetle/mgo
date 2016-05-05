package mcon

import (
	"fmt"
	"syscall"
)

var (
	BLACK = 0
	WHITE = 15 //info

	YELLOW = 14 //warn
	RED    = 12 //error
	GREEN  = 10 //

	BLUE    = 9
	MAGENTA = 8  //trace
	CYAN    = 13 //debug
)

func (c Console) ColorPrint(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i)) //12 Red light

	fmt.Print(s)

	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7)) //White dark
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

func (c Console) PrintWhite(s string) { c.ColorPrint(s, WHITE) }
func (c Console) PrintBlack(s string) { c.ColorPrint(s, BLACK) }

func (c Console) PrintRed(s string)    { c.ColorPrint(s, RED) }
func (c Console) PrintYellow(s string) { c.ColorPrint(s, YELLOW) }
func (c Console) PrintGreen(s string)  { c.ColorPrint(s, GREEN) }

func (c Console) PrintBlue(s string)    { c.ColorPrint(s, BLUE) }
func (c Console) PrintMagenta(s string) { c.ColorPrint(s, MAGENTA) }
func (c Console) PrintCyan(s string)    { c.ColorPrint(s, CYAN) }
