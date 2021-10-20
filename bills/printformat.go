package main

import "fmt"

func main() {
	age := 29
	name := "Samid"

	// fmt.Print("Hello, ")
	// fmt.Print("World! \n")

	// fmt.Println("my age is", age, "and my name is ", name)

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T\n", age)
}
