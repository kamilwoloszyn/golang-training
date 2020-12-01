package context

import (
	"errors"
	"fmt"
)

func IsString(value interface{}) error {
	if _, ok := value.(string); ok {
		return nil
	} else {
		return errors.New("Isn't string")
	}
}

func Run() {
	fmt.Println("Let's check whether is it string or not")
	var myString string = "test string"
	var myNumber int = 10

	if err := IsString(myString); err != nil {
		fmt.Printf("String not found")
	}

	if err := IsString(myNumber); err != nil {
		fmt.Printf("String not found")
	}
}
