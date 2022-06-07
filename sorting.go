package main

import (
	"fmt"
	"sort"
)

func sortingtest() {
	nums := []int{134, 354, 534, 325, 34, 23, 23, 45, 3, 4, 3}
	sort.Ints(nums)
	fmt.Println(nums)

	// sortin in reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

	words := []string{"java", "python", "c", "c++", "javascript", "scala", "rust", "kotlin"}
	sort.Strings(words)
	fmt.Println(words)

	fmt.Println("are strings sorted: ", sort.StringsAreSorted(words))

	// sorting by functions
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
