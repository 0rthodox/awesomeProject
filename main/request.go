package main

//import "fmt"
//
//type Animal struct {
//	food, locomotion, noise string
//}
//
//func (animal * Animal) Eat() {
//	fmt.Println(animal.food)
//}
//
//func (animal * Animal) Move() {
//	fmt.Println(animal.locomotion)
//}
//
//func (animal * Animal) Speak() {
//	fmt.Println(animal.noise)
//}
//
//func main()  {
//	var cow = Animal{ "grass", "walk", "moo" }
//	var bird = Animal{ "worms", "fly", "peep" }
//	var snake = Animal{ "mice", "slither", "hsss" }
//	for {
//		var animalName, actionName string
//		var animal Animal
//		fmt.Println("Please input an animal and an action: > ")
//		_, err := fmt.Scanln(&animalName, &actionName)
//		if err != nil {
//			panic(err)
//		}
//		switch animalName {
//		case "cow": animal = cow
//		case "bird": animal = bird
//		case "snake": animal = snake
//		default: fmt.Println("unknown animal: " + animalName)
//		}
//		switch actionName {
//		case "eat": animal.Eat()
//		case "move": animal.Move()
//		case "speak": animal.Speak()
//		default: fmt.Println("unknown action: " + actionName)
//		}
//	}
//}
