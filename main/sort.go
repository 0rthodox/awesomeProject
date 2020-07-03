package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//	"sync"
//)

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//	"sync"
//)
//
//func sortAndSignal(numbers []int, wg * sync.WaitGroup) {
//	sort.Ints(numbers)
//	fmt.Println(numbers)
//	wg.Done()
//}
//
//func merge(data []int, a, m, b int) {
//	merged := make([]int, 0, b - a)
//	i := a
//	for _, item := range data[m:b] {
//		for item > data[i] && i < m {
//			merged = append(merged, data[i])
//			i++
//		}
//		merged = append(merged, item)
//	}
//	for _, item := range data[i:m] {
//		merged = append(merged, item)
//	}
//	for i, item := range merged {
//		data[i + a] = item
//	}
//}
//
//func main() {
//	fmt.Println("Please input a sequence of numbers\nand press Enter, e.g. 1 2 3 4\\n")
//	reader := bufio.NewReader(os.Stdin)
//	numbers := make([]int, 0, 10)
//	unparsedString, err := reader.ReadString('\n')
//	if err != nil {
//		panic(err)
//	}
//	splitStrings := strings.Split(unparsedString," ")
//	for _, stringedNumber := range splitStrings {
//		if parsedNumber, err := strconv.Atoi(strings.TrimSpace(stringedNumber)); err == nil {
//			numbers = append(numbers, parsedNumber)
//		} else {
//			panic(err)
//		}
//	}
//	fmt.Println("Parsed numbers:")
//	fmt.Println(numbers)
//	sizeOfBlock := len(numbers) / 4
//	var wg sync.WaitGroup
//	wg.Add(3)
//	fmt.Println("Sorted parts:")
//	for i := 0; i < 3; i++ {
//		go sortAndSignal(numbers[i * sizeOfBlock : (i + 1) * sizeOfBlock], &wg)
//	}
//	sort.Ints(numbers[3 * sizeOfBlock :])
//	fmt.Println(numbers[3 * sizeOfBlock :])
//	wg.Wait()
//	merge(numbers, 0, sizeOfBlock, 2 * sizeOfBlock)
//	merge(numbers, 2 * sizeOfBlock, 3 * sizeOfBlock, len(numbers))
//	merge(numbers, 0, 2 * sizeOfBlock, len(numbers))
//	fmt.Println("After merging:")
//	fmt.Println(numbers)
//
//}

