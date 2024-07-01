package main

import "fmt"

func main() {
	u := newAdmin("some@email.com", "somepass")
	u.print()
	u.clear()
	u.print()

}

func newUser(firstName string, lastName string) user {
	return user{firstName, lastName}
}

func newAdmin(email, password string) admin {
	adm := admin{}
	adm.email = email
	adm.password = password
	adm.firstName = "admin"
	adm.lastName = "admin"

	return adm
}

type user struct {
	firstName string
	lastName  string
}

type admin struct {
	email    string
	password string
	user
}

func (u *user) clear() {
	u.firstName = ""
	u.lastName = ""
	fmt.Println("Cleared")
}

func (u *user) print() {
	fmt.Printf("Printing: %s %s\n", u.firstName, u.lastName)
}
