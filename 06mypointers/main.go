package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")
	va := 53
	ptr := &va
	fmt.Println("The address of the pointer variable itself: ", &ptr) 
	fmt.Println("The value of ptr, which is the memory address of va: ", ptr)
	fmt.Println("Value stored at the address ptr is pointing to: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is : ", va)
}