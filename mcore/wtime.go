package mcore

import (
	"fmt"
	"time"
)

type Year int
type Month int
type Day int

type Hour int
type Minute int
type Second int

type WeekDay int
type Quarter int

type Date struct {
	Year
	Month
	Day
}

type Time struct {
	Hour
	Minute
	Second
}

type DateTime struct {
	Year
	Month
	Day
	Hour
	Minute
	Second
}

const (
	Q1 Quarter = 1 + iota
	Q2
	Q3
	Q4
)

const (
	M1 Month = 1 + iota
	M2
	M3
	M4
	M5
	M6
	M7
	M8
	M9
	M10
	M11
	M12
)

const (
	JAN Month = 1 + iota
	FEB
	MAR
	APR
	MAY
	JUN
	AUG
	SEP
	OCT
	NOV
	DEC
)

const (
	SUN WeekDay = iota
	MON
	TUE
	WED
	THU
	FRI
	SAT
)

func (d Date) String() string {
	return fmt.Sprintf("%d-%2d-%2d", d.Year, d.Month, d.Day)
}
func (t Time) String() string {
	return fmt.Sprintf("%d:%d:%d", t.Hour, t.Minute, t.Second)
}
func (t DateTime) String() string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year, t.Month, t.Day, t.Hour, t.Minute, t.Second)
}

var (
	WeekDayArray []string = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
)

func (t WeekDay) String() string {
	n := int(t)
	if n > 5 || n < -1 {
		return "ErrDay"
	}
	return WeekDayArray[n]
}

var (
	MonthArray []string = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
)

func (t Month) String() string {
	n := int(t)
	if n > 12 || n < 0 {
		return "ErrMonth"
	}
	return MonthArray[n-1]
}

func NewDate(y, m, d int) (date Date) {
	date.Year = Year(y)
	date.Month = Month(m)
	date.Day = Day(d)
	return
}

func GetCurrentDate() Date {
	return WrapDate(time.Now())
}

func WrapDate(t time.Time) Date {
	y := t.Year()
	m := int(t.Month())
	d := t.Day()
	return NewDate(y, m, d)
}

func GetCurrentYear() int {
	return int(GetCurrentDate().Year)
}
func GetCurrentMonth() int {
	return int(GetCurrentDate().Month)
}
func GetCurrentDay() int {
	return int(GetCurrentDate().Day)
}

func NewTime(h, m, s int) (t Time) {
	t.Hour = Hour(h)
	t.Minute = Minute(m)
	t.Second = Second(s)
	return
}

func GetCurrentTime() Time {
	return WrapTime(time.Now())
}

func WrapTime(tm time.Time) (t Time) {
	h, m, s := tm.Clock()
	t = NewTime(h, m, s)
	return
}

func GetCurrentDateTime() (dt DateTime) {
	return WrapDateTime(time.Now())
}

func WrapDateTime(tm time.Time) (dt DateTime) {
	d := WrapDate(tm)
	t := WrapTime(tm)

	dt.Year = d.Year
	dt.Month = d.Month
	dt.Day = d.Day

	dt.Hour = t.Hour
	dt.Minute = t.Minute
	dt.Second = t.Second
	return
}

func (dt DateTime) Date() Date {
	return NewDate(int(dt.Year), int(dt.Month), int(dt.Day))
}
func (dt DateTime) Time() Time {
	return NewTime(int(dt.Hour), int(dt.Minute), int(dt.Second))
}

func Parse(layout, v string) (d DateTime, err error) {
	var t time.Time
	t, err = time.Parse(layout, v)
	if err != nil {
		return
	}
	d = WrapDateTime(t)
	return
}

const DefaultTimeFormat = "2006-01-02 15:04:05"

func ParseDefault(v string) (DateTime, error) {
	return Parse(DefaultTimeFormat, v)
}

func GetCurrentQuarter() Quarter {
	m := GetCurrentMonth()
	return GetQuarter(int(m))
}

func GetQuarter(month int) Quarter {
	switch month {
	case 1, 2, 3:
		return Q1
	case 4, 5, 6:
		return Q2
	case 7, 8, 9:
		return Q3
	case 10, 11, 12:
		return Q4
	default:
		return Q1
	}
}

type Week int

func GetCurrentWeek() Week {
	t := time.Now()
	_, w := t.ISOWeek()
	return Week(w)
}

func GetCurrentWeekDay() WeekDay {
	t := time.Now()
	wd := t.Weekday()
	return WeekDay(int(wd))
}

type YearDay int

func GetCurrentYearDay() YearDay {
	t := time.Now()
	yd := t.YearDay()
	return YearDay(yd)
}

func Now() time.Time {
	return time.Now()
}

// FormatTime
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// PrintTime
func PrintTime(t time.Time) {
	ts := FormatTime(t)
	fmt.Printf("%s\n", ts)
}
