package calculate

import (
	"slices"

	"github.com/hanifmaliki/alterra-test/helper"
)

func CalculateCombinations(target, maxSum int) int {
	numbers := make([]int, target)
	for i := range numbers {
		numbers[i] = 1
	}
	combinations := 1
	helper.PrintIntSlice(numbers, target)
	index := len(numbers) - 1

	for len(numbers) > 1 && index > 0 {
		if numbers[index]+numbers[index-1] <= maxSum {
			numbers[index-1] += numbers[index]
			numbers = slices.Delete(numbers, index, index+1)
			combinations++
			helper.PrintIntSlice(numbers, target)
		}
		index--
	}
	return combinations
}
