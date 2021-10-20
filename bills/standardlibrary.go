package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello, there friends!"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.Split(greeting, " "))

	ages := []int{35, 20, 23, 2, 44}
	sort.Ints(ages)
	// fmt.Println(ages)

	index := sort.SearchInts(ages, 23)
	fmt.Println(index)

	names := []string{"Dimas", "Septyanto", "Putro"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Putro"))

}
