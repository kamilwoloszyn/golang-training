package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sampleString := "I know, and you not"
	var buffer [utf8.UTFMax]byte
	for i, r := range sampleString {
		rl := utf8.RuneLen(r)
		sliceOffet := i + rl

		copy(buffer[:], sampleString[i:sliceOffet])
		fmt.Printf("Slice Offset: %d and rune : %q %#6x and buffer: %#v\n", sliceOffet, r, r, buffer[:rl])

	}
}
