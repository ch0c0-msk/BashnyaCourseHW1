package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()
	numbersList := strings.Fields(inputLine)

	numbersMap := make(map[int]bool)
	uniqueCount := 0

	for _, numStr := range numbersList {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("Wrong input")
		}
		if _, isExist := numbersMap[num]; !isExist {
			uniqueCount++
			numbersMap[num] = true
		}
	}

	fmt.Println(uniqueCount)
}
