package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("\nType 'Ctrl' + 'c' to exit the program")
	for {
		in := bufio.NewReader(os.Stdin)
		fmt.Println("\nInsert a string:")
		s, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		s = strings.ToLower(strings.Trim(s, "\n"))
		matched, _ := regexp.MatchString(`^i.*a.*n$`, s)
		if matched {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
