// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	re := regexp.MustCompile(`[-,!?:.\t\n ]`)
	wordsMap := make(map[string]int)
	words := re.Split(strings.ToLower(text), -1)

	if len(words) != 0 {
		for _, word := range words {
			if word != "" {
				if strings.HasSuffix(word, "'s") {
					parts := strings.Split(word, "'s")
					word = parts[0] + "s"
					wordsMap[word]++
				} else {
					wordsMap[word]++
				}
			}
		}
		return wordsMap
	}
	return map[string]int{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		fmt.Println(CountWordFrequency(string(input)))

	}

	if err := scanner.Err(); err != nil {
		return
	}
	//input := "The quick brown fox jumps over the lazy dog."
	//fmt.Println(CountWordFrequency(input))
}
