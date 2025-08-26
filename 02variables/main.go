package main

import "fmt"

func main() {
	/* 
	1. fmt.Print()
		Basic print without adding a newline.
		Works like console.log in JavaScript without the automatic line break.

	2. fmt.Println()
		Same as Print, but adds a space between arguments and ends with a newline.

	3. fmt.Printf()
		Formatted print (like printf in C).
		You provide a format string with placeholders (%s, %d, %f, etc.).
	*/
	var username string = "Aryan"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallfloat float32 = 255.42093798273987298349823
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n", smallfloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)
	var username2 string
	fmt.Println(username2)
	fmt.Printf("variable is of type: %T \n", username2)

	//implicit type:
	var website = "learncode"
	fmt.Println(website)

	//no var style: Walrus operator
	numberOfUser := 300000
	fmt.Println(numberOfUser)
}