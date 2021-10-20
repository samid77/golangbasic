package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func propmtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		fmt.Println(name, price)
	case "t":

		tip, _ := getInput("Enter Tip (Rp.): ", reader)
		fmt.Println(tip)

	case "s":
		fmt.Println("choose s")
	default:
		fmt.Println("that was not a valid option")
		propmtOptions(b)
	}

}

func main() {

	// sayHello("Samid")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// mybill := newbill("Bill no.447")

	// mybill.addItem("Zuppa soup", 4.50)
	// mybill.addItem("Tenderloin", 14.50)
	// mybill.updateTip(10)

	// fmt.Println(mybill.format())

	mybill := createBill()
	propmtOptions(mybill)

}
