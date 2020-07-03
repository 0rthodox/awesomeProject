package main

//func GenDisplaceFn(a, v0, s0 float64) func (t float64) float64 {
//	displaceFunc := func(t float64) float64 {
//		return a / 2 * math.Pow(t, 2) + v0 * t + s0
//	}
//	return displaceFunc
//}
//
//func main() {
//	fmt.Println("Please input acceleration, initial velocity,")
//	fmt.Println("initial distance and time:")
//	var a, v0, s0, t float64
//	_, err := fmt.Scan(&a, &v0, &s0, &t)
//	if err != nil {
//		panic(err)
//	}
//	fn := GenDisplaceFn(a, v0, s0)
//	fmt.Print("The answer is ")
//	fmt.Println(fn(t))
//}
