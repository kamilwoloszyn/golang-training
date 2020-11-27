package basics

import "fmt"

func InspectSlice(s []string) {

	fmt.Printf("Starting with address: %p\nIt contains %d elements od %d (len,capacity)", &s, len(s), cap(s))
	for i, v := range s {
		fmt.Printf("Element [%d]: %p with element addr: %p \n", i, &s[i], &v)
	}
}

func run() {

	sampleSlice := make([]string, 10, 20)
	sampleSlice[0] = "somestring1"
	sampleSlice[1] = "somestring2"
	sampleSlice[2] = "somestring3"
	sampleSlice[3] = "somestring4"
	InspectSlice(sampleSlice)

}
