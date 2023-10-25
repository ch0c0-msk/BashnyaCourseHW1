package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	inputSlice := make([][]int, n)
	for i := range inputSlice {
		inputSlice[i] = make([]int, n)
	}
	var temp int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&temp)
			inputSlice[i][j] = temp
		}
	}

	isSimmetric := true
	for i := 0; i < (n - 1); i++ {
		for j := i + 1; j < n; j++ {
			if inputSlice[i][j] != inputSlice[j][i] {
				isSimmetric = false
				break
			}
		}
	}
	if isSimmetric {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
