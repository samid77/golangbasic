package main

import "fmt"

var score = 99.5

func main() {

	// sayHello("Samid")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	mybill := newbill("Bill no.447")

	mybill.addItem("Zuppa soup", 4.50)
	mybill.addItem("Tenderloin", 14.50)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
