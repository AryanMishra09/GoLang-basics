package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Values of languages: ", languages)
	fmt.Println("Js shorts for: ", languages["JS"])

	//delete some values
	delete(languages, "RB")
	fmt.Println("Values of languages: ", languages)
	//loops are interesting in golang:
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}