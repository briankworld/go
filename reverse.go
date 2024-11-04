package main

import (
  "fmt"
  "reflection"
  "testing"
)

func reverse(arr []int) {
  for i := range len(arr)/2 {
    arr[i], arr[len(arr)-1] =  arr[len(arr)-1], arr[i]
  }
}

func TestReverse(t *testing.T) {
  testCases := []struct {
    input []int
    output[] int
  } {
     { input: []int{1,2,3,4,5}, output: []int{5,4,3,2,1} },
     { input: []int{1,2,3,4}, output: []int{4,3,2,1},
  }

  for tc := range testCases {
    reverse(input)
    if !reflection.DeepEqual(tc.input, t.output) {
      t.Errorf("Reverse(%v), want: %v", tc.intput, tc.output)
    }
  }
}
