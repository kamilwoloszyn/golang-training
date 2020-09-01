package main

import "fmt"

func main() {
	var array [5]string = [5]string{"banan", "kot", "pies", "gruszka", "japko"}

	for i, v := range array {
		fmt.Printf("%s, Address i: %p, addres of one element:%p  and address array : %p \n", v, &i, &v, &array[i])
	}
}
