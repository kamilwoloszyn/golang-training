package mutation

import "fmt"

func run() {

	var testSlice []string
	var lastCapacity = cap(testSlice)
	fmt.Printf("Starting at %p, len %d and capacity\n", &testSlice, len(testSlice))
	for record := 0; record < 1300; record++ {
		testSlice = append(testSlice, "Record %d")
		if lastCapacity != cap(testSlice) {
			capCh := float64((cap(testSlice) - lastCapacity)) / float64(cap(testSlice)) * 100
			fmt.Printf("Current addr: %p, size changed from %d -> %d (%.2f) with elements %d \n", &testSlice, lastCapacity, cap(testSlice), capCh, len(testSlice))
			lastCapacity = cap(testSlice)
		}
	}

	fmt.Println("\nOk, now let borrow some values from original testSlice")

	var anotherSlice []string = testSlice[30:100]
	var lastSlice []string = anotherSlice[1:3:4]
	anotherSliceLen := len(anotherSlice)
	anotherSliceCapacity := cap(anotherSlice)
	fmt.Printf("Starts from %p wchich refers to %p", &anotherSlice[0], &testSlice[30])
	fmt.Println("\nOk, let me expand anotherSlice")
	anotherSlice = append(anotherSlice, "toya")
	fmt.Printf("Len:%d  Cap: %d -> Len: %d Cap: %d", anotherSliceLen, anotherSliceCapacity, len(anotherSlice), cap(anotherSlice))
	fmt.Printf("\nNow create a slice with len %d and cap %d with addr: %p \n", len(lastSlice), cap(lastSlice), &lastSlice)
	lastSlice = append(lastSlice, "Shrek")
	lastSlice = append(lastSlice, "Fiona")
	fmt.Printf("After append addr: %p with len %d, and capacity %d \n", &lastSlice, len(lastSlice), cap(lastSlice))
	fmt.Printf("Compare with original: %p != %p so this is new slice", &lastSlice[0], &anotherSlice[1])
}
