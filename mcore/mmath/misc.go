package mmath

import (
	"math"
)

// Round(3.1556,2)=3.16
// Round(3.1556,0)=3
func Round(val float64, places int)float64{
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}
	return t
}

// RoundInt(3.15)=3
func RoundInt(v float64)(r int64){
	f:=Round(v,0)
	return int64(f)
}

// RoundUp(3.14159, 2)=3.15
func RoundUp(v float64, places int)(r float64){
	if places == 0{
		return float64(math.Ceil(v))
	}
	f := math.Pow10(places)
	x := v * f
	r = math.Ceil(x)/f
	return
}

func RoundUpInt(v float64)(r int64){
	f:=math.Ceil(v)
	return int64(f)
}

// RoundDown(3.14159,3)=3.141
// truncate
func RoundDown(v float64, places int)(r float64){
	if places == 0{
		return float64(math.Floor(v))
	}
	f := math.Pow10(places)
	x := v * f
	r = math.Floor(x)/f
	return
}

func RoundDownInt(v float64)(r int64){
	f:=math.Floor(v)
	return int64(f)
}



