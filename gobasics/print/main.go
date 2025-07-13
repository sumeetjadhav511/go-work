package main

import "fmt"

func main(){
	age :=25
	name :="sumeet"
	height :=5.4123456

	fmt.Println("age: ",age,"height: ",height,"name: ",name)

	fmt.Println("-------------------------------")

	fmt.Printf(" Age is %d \n Height is %.3f \n",age,height)
	fmt.Printf(" \n Height is %T variable \n",height)

	fmt.Printf("name is %s              ",name)
	//fmt.Printf("name is %s \n",name)

	
}