package main

import (
	"fmt"
)

func main() {
	var a, b int
	// Read two integers from standard input
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Printf("Scan error from standart input: %s\n", err)
		return
	}
	sum := Sum(a, b)
	fmt.Println(sum)

	// Call the Sum function and print the result

}

// Sum returns the sum of a and b.
func Sum(a int, b int) int {
	if a != 0 || b != 0 {
		return a + b
	}
	// TODO: Implement the function
	return 0
}
