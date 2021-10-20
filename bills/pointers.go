package main

import "fmt"

func updateName(n *string) {
	*n = "John"
}

func main() {
	name := "Doe"

	m := &name

	updateName(m)

	fmt.Println(*m)
}
