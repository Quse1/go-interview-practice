package main

import (
	"fmt"
	"math"
)

func main() {
	// Example slice for testing
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	// Test FindMax
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)

	// Test RemoveDuplicates
	unique := RemoveDuplicates(numbers)
	fmt.Printf("After removing duplicates: %v\n", unique)

	// Test ReverseSlice
	reversed := ReverseSlice(numbers)
	fmt.Printf("Reversed: %v\n", reversed)

	// Test FilterEven
	evenOnly := FilterEven(numbers)
	fmt.Printf("Even numbers only: %v\n", evenOnly)
}

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
	// TODO: Implement this function
	if len(numbers) == 0 {
		return 0
	}

	max := math.MinInt64
	for _, num := range numbers {
		if max < num {
			max = num
		}
	}
	return max
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
	// TODO: Implement this function
	if len(numbers) == 0 {
		return []int{}
	}

	numMap := make(map[int]struct{})
	nums := []int{}
	for _, num := range numbers {
		if _, inMap := numMap[num]; !inMap {
			numMap[num] = struct{}{}
			nums = append(nums, num)
		}
	}
	return nums
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
	// TODO: Implement this function
	if len(slice) == 0 {
		return []int{}
	}

	reverse := []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		reverse = append(reverse, slice[i])
	}
	return reverse
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
	// TODO: Implement this function
	if len(numbers) == 0 {
		return []int{}
	}

	even_slice := []int{}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			even_slice = append(even_slice, numbers[i])
		}
	}
	return even_slice
}
