package main

import (
	"fmt"
)

func main() {

	fmt.Println("We are learning switch")

	day := 3

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3, 4, 5:
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Other")
	// }

	switch {
	case day == 1:
		fmt.Println("Monday")
	case day == 2:
		fmt.Println("Tuesday")
	case day == 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other")
	}

}
