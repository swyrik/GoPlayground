package main

import (
	"fmt"
)

func mapstest() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 10

	fmt.Println("map: ", m)

	fmt.Println("get: ", m["k2"])
	fmt.Println("length: ", len(m))

	delete(m, "k1")
	fmt.Println("map: ", m)

	_, prs := m["k3"]
	fmt.Println("is k3 key available in m: ", prs)

	n := map[string]string{"name": "swyrik", "age": "33"}
	fmt.Println(n)

	fmt.Printf("type of map %T\n", m)

	for k, v := range n {
		fmt.Println(k, v)
	}
}
