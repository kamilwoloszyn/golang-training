package main

import "fmt"

func main() {
	baseSlice := make([]int, 6)
	for i := range baseSlice {
		baseSlice[i] = i
	}
	fmt.Printf("baseSlice: %v with len: %d and cap: %d\n", baseSlice, len(baseSlice), cap(baseSlice))
	one := &baseSlice[1]
	fmt.Printf("Borrowed value with addr: %p and value : %d\n", &one, *one)

	for i := 0; i < 20; i++ {
		baseSlice = append(baseSlice, i)
	}
	fmt.Printf("baseSlice expanded. Now len: %d and cap: %d\n", len(baseSlice), cap(baseSlice))
	*one = 100
	fmt.Printf("Changed value of borrowed slice: Now value is %d and value of baseSlice is : %d \naddr: %p != %p", *one, baseSlice[1], &one, &baseSlice[1])
}
