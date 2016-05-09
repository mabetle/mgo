package mstat

type IntSlice []int
type Int64Slice []int64

func IntToFloat64Slice(values []int) Float64Slice {
	r := []float64{}
	for _, v := range values {
		r = append(r, float64(v))
	}
	return Float64Slice(r)
}

func Int64ToFloat64Slice(values []int64) Float64Slice {
	r := []float64{}
	for _, v := range values {
		r = append(r, float64(v))
	}
	return Float64Slice(r)
}
