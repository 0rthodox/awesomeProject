package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func BubbleSort(sli []int) {
//	for i := 0; i < len(sli); i++ {
//		for j := len(sli) - 1; j > i; j-- {
//			if sli[j] < sli[j - 1] {
//				Swap(sli, j - 1)
//			}
//		}
//	}
//}
//
//func Swap(sli []int, i int) { // Swaps i-th and i + 1 -th elems
//	sli[i], sli[i + 1] = sli[i + 1], sli[i]
//}
//
//func ReadSequence() []int {
//	fmt.Println("Please input up to 10 ints divided by spaces and press Enter:")
//	sequence := make([]int, 0, 10)
//	reader := bufio.NewReader(os.Stdin)
//	if inputString, err := reader.ReadString('\n'); err == nil {
//		stringSequence := strings.Split(inputString, " ")
//		for _, item := range stringSequence  {
//			if integer, err := strconv.Atoi(strings.TrimSpace(item)); err == nil {
//				sequence = append(sequence, integer)
//			} else {
//				panic(err)
//			}
//		}
//	} else {
//		panic(err)
//	}
//	return sequence
//}
//
//func main() {
//	sequence:= ReadSequence()
//	fmt.Println("Input sequence:")
//	fmt.Println(sequence)
//	BubbleSort(sequence)
//	fmt.Println("Sorted sequence:")
//	fmt.Println(sequence)
//}
