package mterm

// wrap package method to object.
// convivent when you want to extern and some methods.
type ColorTerm struct{}

func NewColorTerm() *ColorTerm {
	return &ColorTerm{}
}

func (c ColorTerm) Black(str string) string { return Black(str) }
func (c ColorTerm) White(str string) string { return White(str) }

func (c ColorTerm) Red(str string) string    { return Red(str) }
func (c ColorTerm) Green(str string) string  { return Green(str) }
func (c ColorTerm) Yellow(str string) string { return Yellow(str) }

func (c ColorTerm) Blue(str string) string    { return Blue(str) }
func (c ColorTerm) Magenta(str string) string { return Magenta(str) }
func (c ColorTerm) Cyan(str string) string    { return Cyan(str) }

// Print funcs
func (c ColorTerm) PrintBlack(s string) { PrintBlack(s) }
func (c ColorTerm) PrintWhite(s string) { PrintWhite(s) }

func (c ColorTerm) PrintRed(s string)    { PrintRed(s) }
func (c ColorTerm) PrintGreen(s string)  { PrintGreen(s) }
func (c ColorTerm) PrintYellow(s string) { PrintYellow(s) }

func (c ColorTerm) PrintBlue(s string)    { PrintBlue(s) }
func (c ColorTerm) PrintMagenta(s string) { PrintMagenta(s) }
func (c ColorTerm) PrintCyan(s string)    { PrintCyan(s) }

// Printf funcs
func (c ColorTerm) PrintfBlack(format string, args ...interface{}) { PrintfBlack(format, args...) }
func (c ColorTerm) PrintfWhite(format string, args ...interface{}) { PrintfWhite(format, args...) }

func (c ColorTerm) PrintfRed(format string, args ...interface{})    { PrintfRed(format, args...) }
func (c ColorTerm) PrintfGreen(format string, args ...interface{})  { PrintfGreen(format, args...) }
func (c ColorTerm) PrintfYellow(format string, args ...interface{}) { PrintfYellow(format, args...) }

func (c ColorTerm) PrintfBlue(format string, args ...interface{})    { PrintfBlue(format, args...) }
func (c ColorTerm) PrintfMagenta(format string, args ...interface{}) { PrintfMagenta(format, args...) }
func (c ColorTerm) PrintfCyan(format string, args ...interface{})    { PrintfCyan(format, args...) }
