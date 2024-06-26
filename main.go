package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hanifmaliki/alterra-test/calculate"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <target> <maxSum>")
		return
	}

	target, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid target value: %v\n", err)
		return
	}

	maxSum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid maxSum value: %v\n", err)
		return
	}

	result := calculate.CalculateCombinations(target, maxSum)
	fmt.Println(result)
}
