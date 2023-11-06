package main

import (
	"fmt"
)

func main() {
	var (
		inputSlice []int
		n          int
		temp       int
	)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		inputSlice = append(inputSlice, temp)
	}

	outputSlice := append(inputSlice[n-1:n], inputSlice[0:n-1]...)
	for _, elem := range outputSlice {
		fmt.Printf("%d ", elem)
	}
}
