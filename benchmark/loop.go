package benchmark

func IterHorizontalOrder(a [size][size]int) {
	var dummy int = 10
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			dummy = a[i][j]
			dummy--
		}
	}
}

func IterVerticalOrder(a [size][size]int) {
	var dummy int = 10
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			dummy = a[j][i]
			dummy--
		}
	}
}

func ModifyHorizontalOrder(a [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			a[i][j]++
		}
	}
}

func ModifyVerticalOrder(a [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			a[j][i]++
		}
	}
}
