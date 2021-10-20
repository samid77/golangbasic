package main

import "fmt"

func updateName(x string) {
	x = "Septyanto"
}

func main() {
	name := "Dimas"
	updateName(name)
	fmt.Println(name)
}
