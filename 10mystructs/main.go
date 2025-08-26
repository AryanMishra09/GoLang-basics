package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")
	//there is no inheritancein golang
	//there is no concept of super or parent ( because makers beleive when some function inherit somehting from another class, then the developer have to look very much here and there)
	aryan := User{"Aryan", "aryan@go.dev", true, 16}
	fmt.Println(aryan)
	fmt.Printf("Aryan Details are: %+v\n", aryan)
	fmt.Printf("Name is %v, email is %v, ", aryan.Name, aryan.Email)
}

type User struct {
	Name 	string
	Email 	string
	Status	bool
	Age 	int
}