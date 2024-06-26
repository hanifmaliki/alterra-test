package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintIntSlice(numbers []int, target int) {
	stringSlice := make([]string, len(numbers))

	for i, num := range numbers {
		stringSlice[i] = strconv.Itoa(num)
	}

	joinedString := strings.Join(stringSlice, ",")
	fmt.Println(joinedString + " = " + strconv.Itoa(target))
}
