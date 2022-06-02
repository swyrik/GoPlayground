package main

import "fmt"

func structstest() {
	fmt.Println(Person{"swyrik", 27})

	fmt.Println(Person{name: "jhon", age: 41})
	//omitted values are zeroed and we need to provide the keys as well
	fmt.Println(Person{name: "krissy"})

	// pointer to the struct
	fmt.Println(&Person{"pola", 33})

	fmt.Println(newPerson("kush", 34))
}

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name, age}
	p.age = 22
	return &p
}
