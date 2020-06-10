package main

import "fmt"

func main() {
	// Define Map
	// emails := make(map[string]string)

	// Assign key value
	// emails["Rahul"] = "rahulkosamkar@gmail.com"
	// emails["Rohit"] = "rohitkosamkar97@gmail.com"
	// emails["test"] = "test@gmail.com"

	// Define Map and Assign KV
	emails := map[string]string{"Rahul": "rahulkosamkar@gmail.com", "Rohit": "rohitkosamkar97@gmail.com", "test": "test@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Rahul"])

	// Delete Map
	delete(emails, "test")
	fmt.Println(emails)
}
