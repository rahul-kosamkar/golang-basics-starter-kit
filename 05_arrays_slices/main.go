package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Decalre and Assign
	nameArr := [2]string{"Rahul", "Kosamkar"}

	fmt.Println(fruitArr)
	fmt.Println(nameArr)

	// Slices
	comapanyArray := []string{"FlexiLoans", "Epimoney", "TCS", "KPMG"}
	fmt.Println(comapanyArray)
	// Lengh
	fmt.Println(len(comapanyArray))
	// Sub string
	fmt.Println(comapanyArray[1:3])
}
