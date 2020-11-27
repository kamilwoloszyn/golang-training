package methods

import (
	"fmt"
	"time"
)

type User struct {
	email    string
	password string //should be encrypted, but for this example we don't care
	name     string
	surname  string
	blocked  bool
	age      uint8
}

//ShowDetalis() with value semantics
func (u User) ShowDetails() {
	fmt.Printf("Email: %s, Name: %s, Surname: %s, Age: %d\n", u.email, u.name, u.surname, u.age)
}

// ChangePassword with pointer semantics

func (u *User) ChangePassword(newPassword string) {
	u.password = newPassword
}

//Nothing change, so we use value semantics
func (u User) ShowPassword() {
	fmt.Printf("Password is: %s", u.password)
}

func run() {

	kamil := User{
		"kamilos@example.com",
		"pass123",
		"Kamil",
		"Woloszyn",
		false,
		23,
	}
	time.Now()
	kamil.ShowDetails()
	kamil.ChangePassword("jablko")
	kamil.ShowPassword()
}
