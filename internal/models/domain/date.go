package domain

import "time"

type Date struct {
	t time.Time
}

func NewDate(y int, m time.Month, d int) Date {
	return Date{
		t: time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
	}
}

func NewDateFromTime(t time.Time) Date {
	return NewDate(t.Year(), t.Month(), t.Day())
}

func (d Date) Before(other Date) bool {
	return d.t.Before(other.t)
}
