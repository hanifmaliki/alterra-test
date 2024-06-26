package main

import (
	"fmt"
	"slices"

	"github.com/hanifmaliki/alterra-test/helper"
)

func CalculatePossibility(target int, posibility int) int {
	list := []int{}
	for a := 0; a < target; a++ {
		list = append(list, 1)
	}
	result := 1
	helper.PrintNumbers(list)
	cursor := len(list) - 1
	if len(list) > 1 {
		for {
			if list[cursor]+list[cursor-1] <= posibility {
				list[cursor-1] = list[cursor] + list[cursor-1]
				list = slices.Delete(list, cursor, cursor+1)
				result++
				helper.PrintNumbers(list)
			}
			cursor -= 1
			if cursor == 0 {
				break
			}
		}
	}
	return result
}

func main() {
	result := CalculatePossibility(6, 3)
	fmt.Println(result)
}
