package main

import (
	"fmt"
)

// PrintSlice will print all the values of the slice to standard output.
func PrintSlice(x []int) {
	for _, y := range x {
		fmt.Printf("%d, ", y)
	}
	fmt.Println("\n----")
}

// Add attempts to append a value to the slice.
func Add(x []int, y int) {
	x = append(x, y)
}

func main() {
	x := []int{1}
	Add(x, 2)
	PrintSlice(x)

	// expand the slice to its full capacity.
	x = x[:cap(x)]
	PrintSlice(x)

	// restart, this time using make and a pre-chosen capacity.
	x = make([]int, 0, 2)
	x = append(x, 1)
	Add(x, 2)
	PrintSlice(x)

	// expand the slice to its full capacity.
	x = x[:cap(x)]
	PrintSlice(x)
}
