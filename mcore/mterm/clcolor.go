package mterm

// modify from https://github.com/xcltapestry/xclpkg/blob/master/clcolor/clcolor.go

import (
	"fmt"
)

// show 8 colors.
// balck red greend yellow blue magenta cyan and white
// xterm should has 256 colors in common.
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func Black(str string) string { return ColorText(TextBlack, str) }
func White(str string) string { return ColorText(TextWhite, str) }

func Red(str string) string    { return ColorText(TextRed, str) }
func Green(str string) string  { return ColorText(TextGreen, str) }
func Yellow(str string) string { return ColorText(TextYellow, str) }

func Blue(str string) string    { return ColorText(TextBlue, str) }
func Magenta(str string) string { return ColorText(TextMagenta, str) }
func Cyan(str string) string    { return ColorText(TextCyan, str) }

func ColorText(color int, str string) string {
	if !IsXterm() {
		return str
	}
	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}
