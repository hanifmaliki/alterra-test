package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintNumbers(intSlice []int) {
	strSlice := make([]string, len(intSlice))

	for i, num := range intSlice {
		strSlice[i] = strconv.Itoa(num)
	}

	result := strings.Join(strSlice, ",")
	fmt.Println(result)
}
