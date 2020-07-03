package main

import "fmt"

func main() {
	var number float64
	fmt.Printf("Please enter a floating-point number:\n")
	_, _ = fmt.Scan(&number)
	var truncatedNumber = int64(number)
	fmt.Printf("Truncated integer: %d", truncatedNumber)
}
