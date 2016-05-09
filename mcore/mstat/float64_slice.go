package mstat

type Float64Slice []float64

func (s Float64Slice) Sum() float64 {
	return Sum([]float64(s))
}

func (s Float64Slice) Max() float64 {
	return Max([]float64(s))
}

func (s Float64Slice) Min() float64 {
	return Min([]float64(s))
}

func (s Float64Slice) Avg() float64 {
	return Avg([]float64(s))
}

func (s Float64Slice) Mid() float64 {
	return Mid([]float64(s))
}

func (s Float64Slice) PercentValue(value int) float64 {
	return PercentValue([]float64(s), value)
}

func (s Float64Slice) ValuePercent(value float64) float64 {
	return ValuePercent([]float64(s), value)
}

func (s Float64Slice) Rank(value float64) int {
	return Rank([]float64(s), value)
}

func (s Float64Slice) DescRank(value float64) int {
	return DescRank([]float64(s), value)
}

func (s Float64Slice) Percent10() float64 {
	return Percent10([]float64(s))
}

func (s Float64Slice) Percent25() float64 {
	return Percent25([]float64(s))
}

func (s Float64Slice) Percent50() float64 {
	return Percent50([]float64(s))
}

func (s Float64Slice) Percent75() float64 {
	return Percent75([]float64(s))
}

func (s Float64Slice) Percent90() float64 {
	return Percent90([]float64(s))
}

func (s Float64Slice) Ranks() []int {
	return Ranks([]float64(s))
}

func (s Float64Slice) DescRanks() []int {
	return DescRanks([]float64(s))
}
