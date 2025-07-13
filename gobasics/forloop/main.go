package main

import (
	"fmt"
)

func main() {

	fmt.Println("We are learning FOR loop")

	for i := 0; i < 4; i++ {
		fmt.Println("Number is :", i)
	}
	//infinite loop
	counter := 0
	for {
		fmt.Println("Counter = ", counter)
		counter++
		if counter == 3 {
			break
		}
	}

	// range keyword example
	numbers := []int{11, 22, 33, 44, 55}
	for index, value := range numbers {
		fmt.Printf("Index of number is %d, and value is %d \n", index, value)
	}

	data := "Hello World Sumeet"
	for index, value := range data {
		fmt.Printf("Index of number is %d, and value is %c \n", index, value)
	}
}
