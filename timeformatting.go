package main

import (
	"fmt"
	"time"
)

func timeformattest() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

	p(t1)
	// for fomatting we must always use "Jan 2 15:04:05 MST 2006" as an example
	p(t.Format("3:04PM"))
	p(t.Format("Mon 01/JAN/2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
