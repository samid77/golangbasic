package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 22, 23}

	names := [4]string{"Dimas", "Septyanto", "Putro", "Samid"}
	names[1] = "Vettel"

	fmt.Println(ages)
	fmt.Println(names, len(names))

	var scores = []int{1, 2, 3, 4}
	scores[2] = 12
	scores = append(scores, 14)

	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	// rangeTwo := names[1:]
	// rangeThree := names[:3]
	rangeOne = append(rangeOne, "Del Piero")

	fmt.Println(rangeOne, len(rangeOne))
}
