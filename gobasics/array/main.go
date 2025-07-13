package main

import "fmt"

func main() {

	fmt.Println("We are learning array")
	var name [5]string
	name[0] = "sumeet"
	name[2] = "manjali"
	fmt.Println("Name array is : ", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers are : ", numbers)
	fmt.Println("numbers array length : ", len(numbers))

	var price [5]int
	fmt.Println(price)

	var pricestring [5]string
	pricestring[0] = "sumeet"
	pricestring[2] = "manjali"
	fmt.Printf("Price string %q\n", pricestring)
}
