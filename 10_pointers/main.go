package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Use * to real value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
