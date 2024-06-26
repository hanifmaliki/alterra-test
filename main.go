package main

import (
	"fmt"

	"github.com/hanifmaliki/alterra-test/calculate"
)

func main() {
	result := calculate.CalculateCombinations(6, 3)
	fmt.Println(result)
}
