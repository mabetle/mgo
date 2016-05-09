package mstat

import (
	"sort"

	"github.com/mabetle/mgo/mcore/mmath"
)

// PercentValue gets values percent
func PercentValue(values []float64, percent int) (r float64) {
	fv := sort.Float64Slice(values)
	sort.Sort(fv)
	if percent > 100 {
		percent = 100
	}
	if percent < 0 {
		percent = 0
	}
	if percent == 0 {
		return Min(values)
	}
	if percent == 100 {
		return Max(values)
	}

	n := len(values)
	jg := (float64(n) - 1.0) / 100.0
	theJG := float64(percent) * jg
	nCell := mmath.RoundDownInt(theJG)
	nFloat := fv[nCell]
	nextFloat := fv[nCell+1]
	nAmount := (nextFloat - nFloat) * (theJG - float64(nCell))
	return nFloat + nAmount
}

func ValuePercent(values []float64, value float64) float64 {
	fv := sort.Float64Slice(values)
	sort.Sort(fv)
	if value <= fv[0] {
		return 0.0
	}
	if value >= fv[len(fv)-1] {
		return 100.0
	}
	n := len(fv)
	jg := float64(100) / (float64(n) - 1.0)
	var theJG, amountJG float64 = 0.0, 0.0
	for k, v := range fv {
		// v < value, skip
		if v < value {
			continue
		}
		// v == value
		if v == value {
			theJG = float64(k) * jg
			break
		}
		theJG = float64(k-1) * jg
		// v > value
		pValue := fv[k-1]
		amountJG = (value - pValue) / (v - pValue)
		break
	}
	logger.Infof("JG:%v,TheJG:%v Amount:%v", jg, theJG, amountJG)
	return theJG + amountJG
}

func Ranks(values []float64) []int {
	r := []int{}
	for _, v := range values {
		r = append(r, Rank(values, v))
	}
	return r
}

func Rank(values []float64, value float64) int {
	fv := sort.Float64Slice(values)
	sort.Sort(fv)
	for k, v := range values {
		if v == value {
			return k + 1
		}
	}
	// not found
	return 0
}

func DescRank(values []float64, value float64) int {
	fv := sort.Float64Slice(values)
	sort.Reverse(fv)
	for k, v := range values {
		if v == value {
			return k + 1
		}
	}
	// not found
	return 0
}

func DescRanks(values []float64) []int {
	r := []int{}
	for _, v := range values {
		r = append(r, DescRank(values, v))
	}
	return r
}

// Max .
func Max(values []float64) float64 {
	fv := sort.Float64Slice(values)
	sort.Sort(fv)
	return fv[len(fv)-1]
}

// Min .
func Min(values []float64) float64 {
	fv := sort.Float64Slice(values)
	sort.Sort(fv)
	return fv[0]
}

func Sum(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum = sum + v
	}
	return sum
}

func Avg(values []float64) float64 {
	n := len(values)
	if n == 0 {
		return 0.0
	}
	sum := Sum(values)
	avg := sum / float64(n)
	return avg
}

func Mid(values []float64) float64 {
	return Percent50(values)
}

// Percent10 .
func Percent10(values []float64) float64 {
	return PercentValue(values, 10)
}

// Percent25 .
func Percent25(values []float64) float64 {
	return PercentValue(values, 25)
}

// Percent50 .
func Percent50(values []float64) float64 {
	return PercentValue(values, 50)
}

// Percent75 .
func Percent75(values []float64) float64 {
	return PercentValue(values, 75)
}

// Percent90
func Percent90(values []float64) float64 {
	return PercentValue(values, 90)
}

func Percent100(values []float64) float64 {
	return Max(values)
}

func Percent0(values []float64) float64 {
	return Min(values)
}
