package main

import (
    "fmt"
    "time"
)

func main() {
    pp := fmt.Println

    // Go doesn’t use yyyy-mm-dd layout to format a time.
    // Instead, you format a special layout parameter: Mon Jan 2 15:04:05 MST 2006
    const (
        layoutISO = "2006-01-02"
        layoutUS  = "January 2, 2006"
    )

    date := "2019-07-20"

    t, _ := time.Parse(layoutISO, date)
    pp(t)
    pp(t.Month())
    pp(t.Weekday())

    pp(t.Format(layoutUS))

    now := time.Now()
    pp(now.Format(layoutUS))

    diff := now.Sub(t)
    pp(int(diff.Hours() / 24))

    // Advance 1 day
    pp(now.Add(time.Hour * 24).Format(layoutUS))

    // Backward 1 day
    pp(now.Add(-time.Hour * 24).Format(layoutUS))
}
