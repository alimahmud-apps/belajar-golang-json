package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{"firstname":"Mohs","MiddleName":"Nur","Lastname":"Alimahmud","Age":0,"Merried":false}`
	jsonByte := []byte(jsonString)

	user := User{}
	err := json.Unmarshal(jsonByte, &user)

	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(user.FirstName)
}
