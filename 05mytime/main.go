package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //Standard to use this particular time to format as per documentation
	/* 
	Because the digits are non-repeating and unique, so they clearly map to a specific time element.

	Component	Value		Meaning
	01			Month		January
	02			Day			2nd
	2006		Year		2006
	15			Hour		3 PM (24h clock)
	04			Minute		04
	05			Second		05
	PM			Period		PM/AM
	MST			Zone		Timezone
	-0700		Offset		Timezone offset
	*/

	createdDate := time.Date(2025, time.May, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}