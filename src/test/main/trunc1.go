package main

import (
  "fmt"
  "strconv"
)

func main() {
  var input string
  print("Enter a floating point number: ")
  fmt.Scan(&input)
  floatValue, err := strconv.ParseFloat(input, 64)

  if err == nil {
    integerValue := int(floatValue)
    println("Result: ", integerValue)
  } else {
    fmt.Printf("%q does not look like a floating point number\n", input)
  }

}

