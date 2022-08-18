package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) updateEmail(updatedEmail string) {
	fmt.Printf("address of user receiver: %p\n", &u)
	u.email = updatedEmail
}

func main() {
	user := user{
		name:  "Lucy",
		email: "lucy@test.com",
	}

	fmt.Printf("address of user object: %p\n", &user)
	fmt.Println("original email: ", user.email)

	updatedEmail := "hellolucy@test.com"
	user.updateEmail(updatedEmail)
	fmt.Println("updated email: ", user.email)
}

/* output
address of user object: 0xc0000b4000
original email:  lucy@test.com
address of user receiver: 0xc0000b4020
updated email:  lucy@test.com
*/
