package main

import "fmt"

func main() {
	ids := []int{34, 45, 6, 56, 77, 88, 56, 43, 34}

	// Loop through Ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not Using Index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add Ids togther
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with Map
	emails := map[string]string{"Rahul": "rahulkosamkar@gmail.com", "Rohit": "rohitkosamkar97@gmail.com", "test": "test@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
