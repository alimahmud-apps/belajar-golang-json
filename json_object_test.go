package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	City       string
	PostalCode string
}
type User struct {
	FirstName  string
	MiddleName string
	Lastname   string
	Age        int
	Merried    bool
	Hobbies    []string
	Address    []Address
}

func TestObject(t *testing.T) {
	user := User{
		FirstName:  "Moh",
		MiddleName: "Nur",
		Lastname:   "Alimahmud",
		Age:        28,
		Merried:    true,
	}

	byte, _ := json.Marshal(user)

	fmt.Println(string(byte))
}
