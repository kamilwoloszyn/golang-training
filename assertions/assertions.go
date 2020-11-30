package assertions

import (
	"fmt"
)

type NameInterface interface {
	GetName()
}

type AgeInterface interface {
	GetAge()
}

type DataInterface interface {
	NameInterface
	AgeInterface
}

type Student struct{}
type Person struct{}

func (Student) GetName() {
	fmt.Println("Getting name !")
}
func (Student) GetAge() {
	fmt.Println("Getting age")
}
func (Person) GetName() {
	fmt.Println("Getting name for person")
}
func (Person) GetAge() {
	fmt.Println("Getting Age")
}

func Run() {
	var age AgeInterface
	var name NameInterface
	var ageAndName DataInterface
	ageAndName = Student{}

	age = ageAndName
	//it throws error
	//ageAndName = age
	if st, ok := ageAndName.(Student); ok {
		ageAndName = st
	}

	v := []DataInterface{
		Student{},
		Person{},
	}

	for _, item := range v {
		if _, ok := item.(Person); ok {
			fmt.Println("Yaay, this is it ! ")
		} else {
			fmt.Println("No, it isn't ")
		}
	}

	fmt.Print("%v, %v ", age, name)
}
