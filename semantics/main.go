package main

import "fmt"

type User struct {
	name  string
	email string
}

//go:noinline
func (u User) CreateUserByValueSemantics(name string, email string) User {

	u = User{
		name:  name,
		email: email,
	}
	return u //allocated on the stack
}

//go:noinline
func (u *User) CreateUserByPointerSemantics(name string, email string) *User {
	var user User = User{
		name:  name,
		email: email,
	}
	return &user //allocated on the heap
}

func (u User) changeUserNameByValueSemantics(name string) {
	u.name = name

}
func (u *User) changeUserNameByPointerSemantics(name string) {
	u.name = name
}

func (u User) displayName() {
	fmt.Println(u.name)
}

//go:noinline
func main() {
	var kamil User
	var kamila *User
	kamil.CreateUserByValueSemantics("Kamil", "kamil@example.com")
	kamila.CreateUserByPointerSemantics("Kamila", "kamila@example.com")

	f1 := kamil.changeUserNameByValueSemantics
	fn2 := kamil.changeUserNameByPointerSemantics

	fn2("Marysia")
	kamil.displayName()

	f1("Zosia")
	kamil.displayName()
}
