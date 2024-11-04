package main

import {
  "fmt"
  "slice"
}

func main() {
  arr := []int {1,2,3,4,5,6,7,8,9}

  arr2 := slice.DeleteFunc(arr, func (a int) bool {
    if a % 2 == 0 {
      return true
    }
    return false
  })

  fmt.Printf("Original: %v, Output %v", arr, arr2)
}
