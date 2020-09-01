package structures

import "fmt"

// B struct
//Struct which include value and pointer to the next struct
type B struct {
	value byte
	next  *B
}

const (
	rows = 4 * 1024
	cols = 4 * 1024
)

var matrix [rows][cols]byte

//IterOverLinkedList (list *B) uint32
// It takes pointer to the list and return counter<uint32>
func IterOverLinkedList(list *B) uint32 {
	var ctr uint32
	var helperPointer *B
	if list != nil {
		ctr++
		helperPointer = list.next
		for helperPointer != nil {
			helperPointer = helperPointer.next
			ctr++
		}
	}
	return ctr
}

// IterOverMatrix () uint32
// It takes matrix, number of rows and number of cols
func IterOverMatrix() uint32 {
	var ctr uint32

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

func main() {

	var myStruct *B
	var lastItemPointer *B
	for row := 0; row < rows; row++ {
		for column := 0; column < cols; column++ {
			matrix[row][column] = 0xFF
			if *&myStruct != nil {
				lastItemPointer.next = &B{
					value: 0xFF,
					next:  nil,
				}
				lastItemPointer = lastItemPointer.next

			} else {
				myStruct = &B{
					value: 0xFF,
					next:  nil,
				}
				lastItemPointer = *&myStruct
			}

		}
	}
	fmt.Printf("LinkedList elemnts: %d and Matrix elements: %d", IterOverLinkedList(*&myStruct), IterOverMatrix())

}
