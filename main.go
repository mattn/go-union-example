package main

import (
	"fmt"
	"time"
)

type Date = time.Time
type Days = int

type DateDaysBase interface {
	priv()
}

func DateOnly(date Date) DateDaysBase {
	return tDateOnly{
		date: date,
	}
}

type tDateOnly struct {
	date Date
}

func (tDateOnly) priv() {}

func DateAndDays(date Date, days Days) DateDaysBase {
	return tDateAndDays{
		date: date,
		days: days,
	}
}

type tDateAndDays struct {
	date Date
	days Days
}

func (tDateAndDays) priv() {}

func main() {
	display := func(d DateDaysBase) {
		switch t := d.(type) {
		case tDateOnly:
			fmt.Println(t.date)
		case tDateAndDays:
			fmt.Println(t.date, t.days)
		}
	}

	d := DateOnly(time.Now())
	display(d)

	d = DateAndDays(time.Now(), 3)
	display(d)
}
