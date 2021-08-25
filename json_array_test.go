package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayEncode(t *testing.T) {
	user := User{
		FirstName: "ali",
		Lastname:  "Mahmud",
		Hobbies:   []string{"gaming", "swimming"},
	}

	byte, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))

}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"ali","MiddleName":"","Lastname":"Mahmud","Age":0,"Merried":false,"Hobbies":["gaming","swimming"]}`
	jsonByte := []byte(jsonString)

	user := User{}
	err := json.Unmarshal(jsonByte, &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(user.FirstName)
	fmt.Println(user.Hobbies)

}

func TestJsonArrayComplexEncode(t *testing.T) {
	user := User{
		FirstName: "alimahmud",
		Address: []Address{
			{
				Street:     "jl raya",
				City:       "tangerang",
				PostalCode: "1234",
			},
			{
				Street:     "jl raya 1",
				City:       "tangerang 1",
				PostalCode: "1234 1",
			},
		},
	}

	byte, _ := json.Marshal(user)
	fmt.Println(string(byte))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"alimahmud","MiddleName":"","Lastname":"","Age":0,"Merried":false,"Hobbies":null,"Address":[{"Street":"jl raya","City":"tangerang","PostalCode":"1234"},{"Street":"jl raya 1","City":"tangerang 1","PostalCode":"1234 1"}]}`
	jsonByte := []byte(jsonString)

	user := User{}
	err := json.Unmarshal(jsonByte, &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(user.FirstName)
	fmt.Println(user.Address)

}
