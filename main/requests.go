package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//type Animal interface {
//	Eat()
//	Move()
//	Speak()
//}
//
//type Cow string
//func (cow Cow) Eat() {
//	fmt.Println("grass")
//}
//func (cow Cow) Move() {
//	fmt.Println("walk")
//}
//func (cow Cow) Speak() {
//	fmt.Println("moo")
//}
//
//type Snake string
//func (snake Snake) Eat() {
//	fmt.Println("mice")
//}
//func (snake Snake) Move() {
//	fmt.Println("slither")
//}
//func (snake Snake) Speak() {
//	fmt.Println("hsss")
//}
//
//type Bird string
//func (bird Bird) Eat() {
//	fmt.Println("worms")
//}
//func (bird Bird) Move() {
//	fmt.Println("fly")
//}
//func (bird Bird) Speak() {
//	fmt.Println("peep")
//}
//
//func main() {
//	animals := make(map[string]Animal)
//	for {
//		fmt.Println("Please input a command: > ")
//		reader := bufio.NewReader(os.Stdin)
//		query, err := reader.ReadString('\n')
//		if err != nil {
//			panic(err)
//		}
//		commands := strings.Split(query, " ")
//		for i := range commands {
//			commands[i] = strings.TrimSpace(commands[i])
//		}
//		if commands[0] == "newanimal" {
//			fmt.Printf("Creating %s...\n", commands[1])
//			switch commands[2] {
//			case "cow":
//				cow := Cow(commands[1])
//				animals[commands[1]] = cow
//			case "bird":
//				bird := Bird(commands[1])
//				animals[commands[1]] = bird
//			case "snake":
//				snake := Snake(commands[1])
//				animals[commands[1]] = snake
//			default:
//				println("Sorry, the animal " + commands[2] + " is unknown :(")
//			}
//			fmt.Println("Created it!")
//		} else if commands[0] == "query" {
//			fmt.Printf("Querying %s...\n", commands[2])
//			animal, presence := animals[commands[1]]
//			if !presence {
//				println("Sorry, " + commands[1] + " has not been created yet :(")
//			} else {
//				switch commands[2] {
//				case "eat": animal.Eat()
//				case "move": animal.Move()
//				case "speak": animal.Speak()
//				default: println("Sorry, action " + commands[2] + " is unknown :(")
//				}
//			}
//		}
//	}
//}
