package main

import "fmt"

func main() {

	fmt.Println("We are learning slice")

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice - ", numbers)
	fmt.Println("Slice length - ", len(numbers))

	numbers = append(numbers, 6, 7, 8, 9, 10, 11)
	fmt.Println("Slice - ", numbers)
	fmt.Println("Slice length - ", len(numbers))
	fmt.Println("Slice capacity - ", cap(numbers))

	name := []string{}
	fmt.Println("name - ", name)
	fmt.Println("Slice name length - ", len(name))
	fmt.Println("Slice namecapacity - ", cap(name))

	//make function

	numbers2 := make([]int, 3, 5)
	fmt.Println("Slice numbers2 - ", numbers2)
	fmt.Println("Slice numbers2 length - ", len(numbers2))
	fmt.Println("Slice numbers2 capacity - ", cap(numbers2))

	numbers2 = append(numbers2, 4, 98)
	fmt.Println("Slice numbers2 - ", numbers2)
	fmt.Println("Slice numbers2 length - ", len(numbers2))
	fmt.Println("Slice numbers2 capacity - ", cap(numbers2))

	numbers2 = append(numbers2, 5)
	fmt.Println("Slice numbers2 - ", numbers2)
	fmt.Println("Slice numbers2 length - ", len(numbers2))
	fmt.Println("Slice numbers2 capacity - ", cap(numbers2))

	numbers2 = append(numbers2, 6, 7, 8, 9)
	fmt.Println("Slice numbers2 - ", numbers2)
	fmt.Println("Slice numbers2 length - ", len(numbers2))
	fmt.Println("Slice numbers2 capacity - ", cap(numbers2))

	numbers2 = append(numbers2, 11)
	fmt.Println("Slice numbers2 - ", numbers2)
	fmt.Println("Slice numbers2 length - ", len(numbers2))
	fmt.Println("Slice numbers2 capacity - ", cap(numbers2))
}
