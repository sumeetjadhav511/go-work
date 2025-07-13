package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple function")
}

func add(a, b int) int {
	return a + b
}

func add2(a, b int) (result int) {
	result = a + b // return a+b
	return
}

func main() {

	fmt.Println("We are learning functions")
	simpleFunction()

	addAns := add(3, 4)
	println("Addition answer is ", addAns)

	addAns2 := add2(3, 4)
	println("Addition answer is ", addAns2)
}
