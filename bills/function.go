package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Dimas")
	// sayGreeting("Septyanto")
	// sayBye("Putro")

	cycleNames([]string{"Dimas", "Septyanto", "Putro"}, sayGreeting)
	cycleNames([]string{"Dimas", "Septyanto", "Putro"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Cicrle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)
}
