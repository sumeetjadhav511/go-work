package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be 0")
	}
	return a / b, nil
}

func divide2(a, b float64) (float64, string) {
	if b == 0 {
		return 0, ("denominator must not be 0")
	}
	return a / b, "nil"
}

func main() {

	fmt.Println("We are learning error handling")

	// divAns, err := divide(10, 0)
	// if err != nil {
	// 	println("Error handling ")
	// }
	// println("Addition answer is ", divAns)

	divAns, _ := divide(10, 4)
	println("Addition answer is ", divAns)

	divAns, _ = divide2(10, 4)
	println("Addition answer is ", divAns)

}
