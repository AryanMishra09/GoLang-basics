package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma ok // comma error : 

	input, err :=reader.ReadString('\n')
	fmt.Print(err)
	fmt.Println("Thanks for rating: ", input) 
	fmt.Printf("Type of this rating: %T", input)	
}