package main

import (
	"fmt"
	"time"
)

func switchcase() {
	i := 33
	switch i {
	case 1:
		fmt.Println("It;s One")
	case 2:
		fmt.Println("It's Two")
	default:
		fmt.Print("It's 5\n")

	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("weekend", time.Now().Weekday())
	default:
		fmt.Println("WeekDay", time.Now().Weekday())
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("Good Morning")
	case 12 < t.Hour() && t.Hour() < 24:
		fmt.Println("Good AfterNoon")

	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(3)
	whatAmI("google")
}
