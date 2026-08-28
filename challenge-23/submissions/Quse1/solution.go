package main

import (
	"fmt"
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

func NaivePatternMatch(text, pattern string) []int {
	n, m := len(text), len(pattern)
	founded := []int{}

	if n == 0 || m == 0 || m > n {
		return []int{}
	}

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			founded = append(founded, i)
		}
	}
	if len(founded) == 0 {
		return []int{}
	}
	return founded
}

func ComputePrefix(pattern string) []int {
	length := 0
	m := len(pattern)
	lps := make([]int, m)
	lps[0] = 0

	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	if len(lps) == 0 {
		return []int{}
	}
	return lps
}

// KMPSearch implements the Knuth-Morris-Pratt algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
	n, m := len(text), len(pattern)
	if n == 0 || m == 0 || m > n {
		return []int{}
	}
	result := []int{}
	lps := ComputePrefix(pattern)

	i := 0
	j := 0

	for i < n {
		if text[i] == pattern[j] {
			i++
			j++
			if j == m {
				result = append(result, i-j)
				j = lps[j-1]
			}
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	if len(result) == 0 {
		return []int{}
	}
	return result
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

	if n == 0 || m == 0 || m > n {
		return []int{}
	}

	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}
	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(text[i])) % q
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
			t = (d*(t-int(text[i])*h) + int(text[i+m])) % q
			if t < 0 {
				t += q
			}
		}
	}

	if len(founded) == 0 {
		return []int{}
	}

	return founded
}
