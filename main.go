package main

import (
	"fmt"

	"github.com/hanifmaliki/alterra-test/helper"
)

func CalculateCombinations(target, maxSum int) int {
	numbers := make([]int, target)
	for i := range numbers {
		numbers[i] = 1
	}
	combinations := 1
	helper.PrintIntSlice(numbers)
	index := len(numbers) - 1

	for len(numbers) > 1 {
		if index > 0 && numbers[index]+numbers[index-1] <= maxSum {
			numbers[index-1] += numbers[index]
			numbers = numbers[:index] // effectively delete the last element
			combinations++
			helper.PrintIntSlice(numbers)
		}
		index--
		if index == 0 {
			index = len(numbers) - 1
		}
	}
	return combinations
}

func main() {
	result := CalculateCombinations(6, 3)
	fmt.Println(result)
}
