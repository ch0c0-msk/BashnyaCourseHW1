package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Scan(&input)
	output := strings.ReplaceAll(input, "1", "one")
	fmt.Println(output)
}
