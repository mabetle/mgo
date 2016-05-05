package mcore

import (
	"os"
)

func IsHasArg(arg string)bool{
	return String(arg).IsInArrayIgnoreCase(os.Args)
}

func IsHasDevArg()bool{
	return IsHasArg("dev")
}

func IsHasProdArg()bool{
	return IsHasArg("prod")
}

func IsHasTestArg()bool{
	return IsHasArg("test")
}
