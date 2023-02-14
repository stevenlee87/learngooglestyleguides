package main

import "fmt"

// Good:
func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func main() {
	fmt.Println(abs(-2))
}
