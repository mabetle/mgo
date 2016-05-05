// Calc percent values.
package mstat

// PercentValue, Percentile
func PercentValue(values []float64, percent float64) (r float64) {

	return
}

func Max(values []float64) {
	return PercentValue(values, 1)
}

func Min(values []float64) {
	return PercentValue(values, 0)
}

func MidValue(values []float64, percent float64) {
	return PercentValue(values, 0.5)
}

func QuarteValue(values []float64) {
	return PercentValue(values, 0.25)
}

func Quarte3Value(values []float64) {
	return PercentValue(values, 0.75)
}
