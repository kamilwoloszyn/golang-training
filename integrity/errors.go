package integrity

import (
	"errors"
	"fmt"
)

var (
	movedPermanently error = errors.New("Moved permanently")
	badRequest       error = errors.New("Bad request")
)

func WebRequest() error {
	//return movedPermanently
	return badRequest
}

func Run() {

	//basic handling
	if err := WebRequest(); err != nil {
		switch err {
		case movedPermanently:
			fmt.Println("Access you tried to access is moved premanently.")
			break
		case badRequest:
			fmt.Println("Bad request!")
			break
		}
	} else {
		fmt.Println("Executed successfuly!")
	}

}
