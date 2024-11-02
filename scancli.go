package main

import (
  "bufio"
  "fmt"
  "os"
  "scanner"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  if scanner.Scan() {
    in := scanner.Text()
    num, err := strconv.Atoi(in)

    if err != nil {
			fmt.Println("Invalid integer:", err)
			return
		}
		fmt.Println("You entered the number:", number)
	} else {
		fmt.Println("Failed to read input")
	}
}
