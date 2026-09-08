package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Get input from the user
	var input string
	fmt.Print("Enter a string to check if it's a palindrome: ")
	fmt.Scanln(&input)

	// Call the IsPalindrome function and print the result
	result := IsPalindrome(input)
	if result {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}

// IsPalindrome checks if a string is a palindrome.
// A palindrome reads the same backward as forward, ignoring case, spaces, and punctuation.
func IsPalindrome(s string) bool {
	// TODO: Implement this function
	// 1. Clean the string (remove spaces, punctuation, and convert to lowercase)
	// 2. Check if the cleaned string is the same forwards and backwards
	if len(s) == 0 {
	    return true
	}
	new_s := []string{}
	re := regexp.MustCompile(`[a-z0-9]+`)
	text := strings.TrimSpace(strings.ToLower(s))

	for _, s := range text {
		if re.MatchString(string(s)) {
			new_s = append(new_s, string(s))
		}
	}
	
	var flag int
	for i :=0; i < len(new_s); i++ {
	    if new_s[i] == new_s[len(new_s)-1-i] {
	        flag++
	    }
	}
	
	if flag == len(new_s) {
	    return true 
	}
	return false 
}
