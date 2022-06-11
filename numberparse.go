package main

import (
	"fmt"
	"strconv"
)

func numberParseTest() {

	p := fmt.Printf

	f, _ := strconv.ParseFloat("1.234", 64)
	p("%v %T\n", f, f)

	i, _ := strconv.ParseInt("-90385472323236", 0, 64)
	p("%v %T\n", i, i)

	// ParseInt Can recognize the hex code as well
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p("%v %T\n", d, d)

	u, e := strconv.ParseUint("900", 0, 64)
	p("%v %T\n", u, u)

	if e != nil {
		p("%v\n", e)
	}

	k, _ := strconv.Atoi("135")
	p("%v\n", k)

	// returns an error on bad input
	_, e = strconv.Atoi("wat")
	fmt.Println(e)
}
