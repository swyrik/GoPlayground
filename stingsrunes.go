package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsrunestest() {

	const s = "స్వైరిక్" // all characters in go by default is UTF-8
	fmt.Println(len(s))  // since strings are raw bytes, this will produce length as18

	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

}
