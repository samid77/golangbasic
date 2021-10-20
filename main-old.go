package main

import (
	"fmt"
	// "strings"
	"sort"
)

func main() {

	// greeting := "Hello there friends"

	// fmt.Println(strings.Contains(greeting, "Hello"))

	ages := []int{45, 55, 60, 30, 21, 88, 5}

	sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 88)
	// fmt.Println(index)

	names := []string{"Dimas", "Septyanto", "Putro", "Samid", "Xabi", "Alonso"}
	// sort.Strings(names)
	// fmt.Println(names)

	// x := 0 

	// for x <= 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	// for i:=0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("The value at indeex %v is %v \n", index, value)
	// }

	for index, value := range names {
		if index == 1 {
			fmt.Println(" continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println(" breaking at pos", index)
			break
		}

		fmt.Printf(" the value at position %v is %v \n", index, value)
	}
}
