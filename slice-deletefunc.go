package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arr2 := slices.Clone(arr)
	arr2 = slices.DeleteFunc(arr2, func(a int) bool {
		if a%2 == 0 {
			return true
		}
		return false
	})

	fmt.Println("Original slice:", arr)
	fmt.Println("After DeleteFunc:", arr2)

	target := 7
	index, found := slices.BinarySearch(arr2, target)
	if found {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found. It would be at index %d\n", target, index)
	}

	fmt.Println("Min", slices.Min(arr2))
	fmt.Println("Max", slices.Max(arr2))

	arr2 = slices.Delete(arr2, 1, 2)
	fmt.Println("arr2", arr2)
}
