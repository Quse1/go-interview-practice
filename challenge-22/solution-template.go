package main

import (
	"fmt"
)

func main() {
	denominations := []int{1, 5, 10, 25, 50}
	amounts := []int{87, 42, 11}

	for _, amount := range amounts {
		minCoins := MinCoins(amount, denominations)
		coinCombo := CoinCombination(amount, denominations)

		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

func CoinCombination(amount int, denominations []int) map[int]int {
	coinMap := make(map[int]int)

	for i := len(denominations) - 1; amount >= 0 && i >= 0; i-- {
		num := denominations[i]
		for amount >= num {
			amount -= num
			coinMap[num]++
		}
		if amount == 0 {
			return coinMap
		}
	}
	return map[int]int{} // пустой результат, если не удалось разменять
}

func MinCoins(amount int, denominations []int) int {
	totalCount := 0
	for i := len(denominations) - 1; amount >= 0 && i >= 0; i-- {
		num := denominations[i]
		for amount >= num {
			amount -= num
			totalCount++
		}
	}
	if amount == 0 {
		return totalCount
	}
	return -1
}
