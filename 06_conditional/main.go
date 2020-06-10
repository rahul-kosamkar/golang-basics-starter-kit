package main

import "fmt"

func main() {
	x := 15
	y := 10

	// If Else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else If
	color := "red"

	if color == "red" {
		fmt.Println("Color is Red")
	} else if color == "blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is not Blue or Red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not Blue or Red")
	}
}
