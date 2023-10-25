package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if (a+b > c) && (a+c > b) && (b+c > a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
