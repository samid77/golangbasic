package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Soup":           4.99,
		"Pie":            7.99,
		"Salad":          6.99,
		"Toffee Pudding": 3.55,
	}

	// fmt.Println(menu)
	// fmt.Println(menu["Pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		1234:  "Dimas",
		6839:  "Septyanto",
		28292: "Putro",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234])

	phonebook[1234] = "Samid"
	fmt.Println(phonebook)

	phonebook[28292] = "Ghisa"
	fmt.Println(phonebook)
}
