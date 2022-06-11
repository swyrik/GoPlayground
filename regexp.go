package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func regexptest() {
	match, _ := regexp.MatchString("p([a-z]+)h", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)h")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))

	// finds all the [start, end] indexs of matching string upto provided integer val
	// if the integer provided is -1 it returns all the matching indexes
	fmt.Println(r.FindAllStringIndex("preech punch paoch preech", 2))
	fmt.Println(r.FindAllStringIndex("preech punch paoch preech", -1))

	fmt.Println(r.FindAllStringSubmatch("peach punch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1))

	fmt.Println(r.Match([]byte("peach")))

	fmt.Println(r.ReplaceAllString("hello preech punch paoch preech", "KSI"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
