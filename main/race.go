package main
//
//import "fmt"
//
//func incrementBy(x* int, num* int) {
//	for i := 0; i < *num; i++ {
//		*x += 1
//	}
//}
//
//func main() {
//	fmt.Println("Race condition is a problem when")
//	fmt.Println("the outcome depends on non-deterministic order.")
//	fmt.Println("Here we can see that the same code produces")
//	fmt.Println("3 different results, listed behind:\n")
//	for i := 0; i < 3; i++{
//		x := 0
//		num := 10000000
//		go incrementBy(&x, &num)
//		result := 0
//		for i := 0; i < num; i++ {
//			result += x
//		}
//		fmt.Println("Result", i, "is", result)
//	}
//	fmt.Println("So, the race condition occurs if there is")
//	fmt.Println("a use of a variable by several routines")
//	fmt.Println("and they are not synchronised.")
//	//The results on my machine are:
//	//Result 0 is 48623241226811
//	//Result 1 is 61505493118476
//	//Result 2 is 54452395678935
//}
