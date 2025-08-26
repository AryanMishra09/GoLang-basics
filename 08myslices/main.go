package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Banana", "Peach"}
	//Printing list:
	fmt.Printf("Type of fruit list is: %T \n", fruitList)
	fmt.Println("Value of fruit list is: ", fruitList)
	fmt.Println("Value of fruit list is: ", len(fruitList))
	//Appending data into fruitlist:
	fruitList = append(fruitList, "Mango", "Tomato")
	fmt.Println("Value of fruit list is: ", fruitList)
	//Slicing a slice:
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 230
	highScores[1] = 931
	highScores[2] = 432
	highScores[3] = 833
	// highScores[4] = 949   *Error*
	highScores = append(highScores, 555, 666, 888)
	fmt.Printf("Type of high score list is: %T \n", highScores)
	fmt.Println("Value of highscores: ", highScores)
	//soting a slice:
	sort.Ints(highScores)
	fmt.Println("Sorted highScores: ", highScores)
	fmt.Println("Is high scores sorted: ", sort.IntsAreSorted((highScores)))
	//how to remove a value from a slice based on index:
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Value of courses: ", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Course list after removing: ", courses)
}