package mcore

import (
	"fmt"
)

type Float32 float32
type Float64 float64

// FormateFloat cast float to formate string
func FormatFloat(v float64) string {
	return fmt.Sprint("%f", v)
}
