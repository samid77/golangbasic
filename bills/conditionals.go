package main

import "fmt"

func main() {
	// age := 29

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)

	// if age < 30 {
	// 	fmt.Println("Age is less than 30")
	// }

	names := []string{"Dimas", "Septyanto", "Putro", "Samid"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breakin at pos", index)
			break
		}

		fmt.Printf("The value at pos %v is %v\n", index, value)
	}

}
