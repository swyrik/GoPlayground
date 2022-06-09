package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func stringfrmttest() {

	p := point{1, 2}

	fmt.Printf("struct1 : %v\n", p) // prints the struct values

	fmt.Printf("struct2 : %+v\n", p) // struct values along with field names

	fmt.Printf("struct3 : %#v\n", p) // source code snippet the produces the value

	fmt.Printf("type: %T\n", p) // type of the value

	fmt.Printf("bool: %t\n", true) // formatting bool values with %t

	fmt.Printf("int: %d\n", 123) // formatting integers with %d

	fmt.Printf("int: %b\n", 14) // binary number format

	fmt.Printf("char: %c\n", 67) // this prints character corresponding to given ASCII number

	fmt.Printf("hex: %x\n", 95993) // this prints hex code of the given value

	fmt.Printf("float1: %f\n", 78.9) // floating value of a number

	fmt.Printf("float2: %e\n", 123400000.0)

	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"") // this formats the string by escaping the escape characaters

	fmt.Printf("str2: %q\n", "\"string\"") // this returns an unformatted string

	fmt.Printf("str3: %x\n", "hex this") // this returns a hex format string

	fmt.Printf("pointer: %p\n", &p) // prints the pointer of a value

	str := "hello"

	fmt.Printf("Pointer of str: %p\n", &str)

	fmt.Printf("Width1: |%6d|%6d|\n", 12111111, 35) // right justified padded with empty spaces

	fmt.Printf("Width2: |%6.2f|%6.2f|%-6.3f|\n", 1.2, 3.45, 6.76777) // you can specify the width of printed floats, it rounds of to ceil if the float is having more digits

	fmt.Printf("Width3: |%6s|%6s|\n", "rust", "go") // right justified

	fmt.Printf("Width3: |%-6s|%-6s|\n", "rust", "go") // left justified

	s1 := fmt.Sprintf("sprintf: a %s\n", "srting") // it formats and returns the string but never prints
	fmt.Printf(s1)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
