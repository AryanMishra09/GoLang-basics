package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Orange"
	fruitlist[3] = "Tomato"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist))

	var vegList = [3]string{"Potato", "Beans", "Mushrooms"}
	fmt.Println("Vegy List is: ", vegList)
	fmt.Println("Vegy List is: ", len(vegList))
}