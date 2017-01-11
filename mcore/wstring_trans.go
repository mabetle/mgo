package mcore

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func (s String) ToInt() (int, error) {
	v := s.TrimSpace().TrimSuffix("\r").TrimSuffix("\n").TrimSpace().String()
	return strconv.Atoi(v)
}

func (s String) ToIntNoError() int {
	var r int
	var err error
	if r, err = s.ToInt(); err == nil {
		return r
	}
	fmt.Printf("Error: %v\n", err)
	//cast error
	return 0
}

// bool pairs: 1/0 YES/NO TRUE/FALSE ON/OFF
func (s String) ToBool() bool {
	return s.IsInIgnoreCase("1", "T", "TRUE", "YES", "Y", "ON")
}

func (s String) ToFloat(bitSize int) (float64, error) {
	return strconv.ParseFloat(string(s), bitSize)
}

func (s String) ToFloat32() (float32, error) {
	f, e := s.ToFloat(32)
	return float32(f), e
}

func (s String) ToFloat32NoError() float32 {
	f, _ := s.ToFloat32()
	return f
}

// ToFloat64
func (s String) ToFloat64() (float64, error) {
	return s.ToFloat(64)
}

// ToFloat64NoError
func (s String) ToFloat64NoError() float64 {
	f, _ := s.ToFloat64()
	return f
}

func (s String) Quote() string {
	return strconv.Quote(string(s))
}

func (s String) Unquote() (string, error) {
	return strconv.Unquote(string(s))
}

func (s String) QuoteToASCII() string {
	return strconv.QuoteToASCII(string(s))
}

func (s String) ToBytes() []byte {
	return bytes.NewBufferString(string(s)).Bytes()
}

// ToTime
// string format: 1999-01-01 00:00:00
func (s String) ToTime() (time.Time, error) {
	ts := s.String()
	return time.Parse("2006-01-02 15:04:05", ts)
}

// ToExcelTime
// string format: int64 days after 1900.01.01
// FIXME not work properly
func (s String) ToExcelTime() time.Time {
	ds := s.ToIntNoError()
	dua := time.Duration(ds - 2)

	bt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
	du := time.Hour * 24 * dua
	bt = bt.Add(du)
	return bt
}

func (s String) ToExcelTimeString() string {
	ss := FormatTime(s.ToExcelTime())
	return fmt.Sprintf(ss)
}
