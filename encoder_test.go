package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncode(t *testing.T) {
	write, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(write)

	user := User{
		FirstName:  "Mr",
		MiddleName: "ali",
		Lastname:   "Mahmud",
	}
	encoder.Encode(user)
	fmt.Println(user)
}
