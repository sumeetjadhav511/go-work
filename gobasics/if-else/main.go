package main

import "fmt"

func main() {

	fmt.Println("We are learning if-else")

	x := 6
	if x > 5 {
		fmt.Println("x > 5")
	} else {
		fmt.Println("x < 5")
	}

	if x < 7 {
		fmt.Println("7> x > 5")
	} else if x > 7 {
		fmt.Println("x > 7")
	} else {
		fmt.Println("x < 5")
	}
}
