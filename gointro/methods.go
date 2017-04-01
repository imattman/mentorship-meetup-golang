package main

import "fmt"

type User struct {
	first string
	last  string
	email string
}

func (u User) String() string {
	return fmt.Sprintf("%s, %s - %s", u.last, u.first, u.email)
}

func main() {
	users := []User{
		{"Alice", "Johnson", "alice@example.com"},
		{"Joe", "Smith", "joes@example.com"},
		{"Bob", "Anderson", "bob@example.com"},
	}
	for _, u := range users {
		fmt.Printf("%s\n", u.String())
	}
}
