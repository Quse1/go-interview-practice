package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Sample texts and patterns
	testCases := []struct {
		text    string
		pattern string
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB"},
		{"AABAACAADAABAABA", "AABA"},
		{"GEEKSFORGEEKS", "GEEK"},
		{"AAAAAA", "AA"},
	}

	// Test each pattern matching algorithm
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Text: %s\n", tc.text)
		fmt.Printf("Pattern: %s\n", tc.pattern)

		// Test naive pattern matching
		naiveResults := NaivePatternMatch(tc.text, tc.pattern)
		fmt.Printf("Naive Pattern Match: %v\n", naiveResults)

		// Test KMP algorithm
		kmpResults := KMPSearch(tc.text, tc.pattern)
		fmt.Printf("KMP Search: %v\n", kmpResults)

		// Test Rabin-Karp algorithm
		rkResults := RabinKarpSearch(tc.text, tc.pattern)
		fmt.Printf("Rabin-Karp Search: %v\n", rkResults)

		fmt.Println("------------------------------")
	}
}

// NaivePatternMatch performs a brute force search for pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func NaivePatternMatch(text, pattern string) []int {
	n, m := len(text), len(pattern)
	//founded := make([]int, m)
	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++

		}
		if j == m {
			return []int{i}
		}
	}
	// TODO: Implement this function
	return nil
}

func ComputePrefix(pattern string) []int {
	m := len(pattern)
	pi := make([]int, m)
	for i := 1; i < m; i++ {
		j := pi[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

// KMPSearch implements the Knuth-Morris-Pratt algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
	pi := ComputePrefix(pattern)
	m := len(text)
	matches := []int{}
	j := 0
	for i := 0; i < m; i++ {
		if j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			matches = append(matches, i-len(pattern)+1)
			j = pi[j-1]
		}
	}
	// TODO: Implement this function
	return matches
}

// RabinKarpSearch implements the Rabin-Karp algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func RabinKarpSearch(text, pattern string) []int {
	// TODO: Implement this function
	d := 256
	q := 101
	m := len(pattern)
	n := len(text)
	p := 0
	t := 0
	h := 1
	match := false
	founded := []int{}

	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}
	for i := 0; i < m; i++ {
		pat, err := strconv.Atoi(string(pattern[i]))
		if err != nil {
			fmt.Printf("From str to int: %s\n", err)
		}
		txti, err := strconv.Atoi(string(text[i]))
		if err != nil {
			fmt.Printf("From str to int: %s\n", err)
		}
		p = (d*p + pat) % q
		t = (d*t + txti) % q
	}
	for i := 0; i < n-m+1; i++ {
		if p == t {
			match = true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				founded = append(founded, i)
			}
		}
		if i < n-m {
			txti, err := strconv.Atoi(string(text[i]))
			if err != nil {
				fmt.Printf("From str to int: %s\n", err)
			}
			txtim, err := strconv.Atoi(string(text[i+m]))
			if err != nil {
				fmt.Printf("From str to int: %s\n", err)
			}
			t = (d*(t-txti*h) + txtim) % q
			if t < 0 {
				t += q
			}
		}
	}
	return founded
}
