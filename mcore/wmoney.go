package mcore

const (
	BAI      = 100
	QIAN     = 1000
	WAN      = 10000
	BAI_WAN  = 100 * 10000
	QIAN_WAN = 1000 * 10000
	YI       = 10000 * 10000
	WAN_YI   = 10000 * 10000 * 10000
)

type Money float64

func (m Money) Divide(beDivided float64) Money {
	r := float64(m) / beDivided
	return Money(r)
}

func (m Money) DivideWan() Money {
	return m.Divide(WAN)
}

func (m Money) DivdeBai() Money {
	return m.Divide(BAI)
}

func (m Money) DivideYi() Money {
	return m.Divide(YI)
}

func (m Money) Multiple(beNum float64) Money {
	r := float64(m) * beNum
	return Money(r)
}

func (m Money) MultipleBai() Money {
	return m.Multiple(BAI)
}

func (m Money) MultipleWan() Money {
	return m.Multiple(WAN)
}

func (m Money) MultipleYi() Money {
	return m.Multiple(YI)
}

func (m Money) Plus(to Money) Money {
	r := float64(m) + float64(to)
	return Money(r)
}

func (m Money) Minus(to Money) Money {
	r := float64(m) - float64(to)
	return Money(r)
}
