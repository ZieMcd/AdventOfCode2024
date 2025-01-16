package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println("Initial slice:", slice)

	sliceA := append(slice, 4, 5)
	fmt.Println("Initial slice:", slice)
	fmt.Println("After appending:", sliceA)

	// slice = append(slice, 6, 7)
	// fmt.Println("After appending again:", slice)
}
