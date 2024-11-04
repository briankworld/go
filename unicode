package main

import (
  "fmt"
  "unicode"
)

func main() {
  str := "a1b2c3"

  for _, s := range str {
    if unicode.IsDigit(s) {
      fmt.Printf("Digit: %c", s)
    } else if unicode.IsLetter(s) {
      fmt.Printf("Letter: %c", s)
    } else {
      fmt.Println("Not a digit or letter")
    }
  }

}
