package main

import "fmt"
import "strings"

func main()  {
	var scanned string
	_, _ = fmt.Scan(&scanned)
	if strings.HasPrefix(scanned, "i") && strings.Contains(
		scanned, "a") && strings.HasSuffix(scanned, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}


