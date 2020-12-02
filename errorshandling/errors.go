package errorshandling

import (
	"fmt"

	"github.com/pkg/errors"
)

type AppErr struct {
	id int
}

func (a *AppErr) Error() string {
	return fmt.Sprintf("App error id: %d", a.id)
}

func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return errors.Wrapf(err, "second call")
	}
	return nil
}

func secondCall(i int) error {
	if err := thirdCall(); err != nil {
		return errors.Wrapf(err, "thirdCall")
	}
	return nil
}

func thirdCall() error {
	return &AppErr{99}
}

// Run () Run this example
func Run() {

	if e := firstCall(1); e != nil {
		switch v := errors.Cause(e).(type) {
		case *AppErr:
			fmt.Printf("Got Err: %d", v.id)
		default:
			fmt.Printf("Default Error")
		}
		fmt.Printf("\nStack trace:\n")
		fmt.Printf("%+v", e)
	}

}
