package mmath

// FloatDivide
func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
