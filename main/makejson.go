//package main
//
//import (
//	"bufio"
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//func main()  {
//	fmt.Println("Please enter your name, then press Enter:")
//	var name, address string
//	_, _ = fmt.Scanln(&name)
//	fmt.Println("Please enter your address:")
//	_, _ = fmt.Scanln(&address)
//	jsonObj, _ := json.Marshal(map[string]string {"name": name,
//		"address": address})
//	fmt.Println("JSON Object:")
//	var out bytes.Buffer
//	_ = json.Indent(&out, jsonObj, "", "\t")
//	_, _ = out.WriteTo(os.Stdout)
//}

package main

//import (
//"bufio"
//"encoding/json"
//"fmt"
//"os"
//)
//
//type person struct {
//	name    string
//	address string
//}
//
//var name1 string
//var address1 string
//var p1 person
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	idmap := make(map[string]string)
//	fmt.Print("enter name   ")
//	scanner.Scan()
//	name1 = scanner.Text()
//	fmt.Println("enter address   ")
//	scanner.Scan()
//	address1 = scanner.Text()
//	p1.name = name1
//	p1.address = address1
//	idmap["name"] = p1.name
//	idmap["address"] = p1.address
//	barr, err := json.Marshal(idmap)
//	if err != nil {
//		fmt.Print("error encountered ", err)
//	}
//	fmt.Print("jason object:  ")
//	os.Stdout.Write(barr)
//}