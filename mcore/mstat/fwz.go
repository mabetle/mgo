// Package mstat percent values.
package mstat

// PercentValue gets values percent
func PercentValue(values []float64, percent float64) (r float64) {

	return
}

// Max .
func Max(values []float64) float64 {
	return PercentValue(values, 1)
}

// Min .
func Min(values []float64) float64 {
	return PercentValue(values, 0)
}

// MidValue .
func MidValue(values []float64, percent float64) float64 {
	return PercentValue(values, 0.5)
}

// QuarteValue .
func QuarteValue(values []float64) float64 {
	return PercentValue(values, 0.25)
}

// Quarte3Value .
func Quarte3Value(values []float64) float64 {
	return PercentValue(values, 0.75)
}
