package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch case in golang")
	// random := rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.Intn(6)
	fmt.Println("value of dice is: ", diceNumber)
	switch(diceNumber){
	case 1: 
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can make 2 spot") 
	case 3:
		fmt.Println("You can make 3 spot")
	case 4:
		fmt.Println("You can make 4 spot")
	case 5:
		fmt.Println("You can make 5 spot")
	case 6:
		fmt.Println("You can make 6 spot")
	default: 
		fmt.Println("What was that?")
	}
}