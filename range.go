package main

import (
	"fmt"
)

func rangetest() {
	nums := []int{2, 3, 4}
	var sum int = 0
	for _, num := range nums { // the first parameter is the index
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if i == 2 {
			fmt.Println("At index", i, "the value in nums is", num)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}
	// range on string iterates over unicode code points
	for i, c := range "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		fmt.Println(i, c)
	}

	fmt.Println("🔥🔥🔥🔥")

	for i, c := range "😀😁😂🤣😃😄😅😆😉😊😋😎🥲😙😗🥰😘😍" {
		fmt.Println(i, c)
	}
}
