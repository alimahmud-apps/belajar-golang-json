package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSteamDecode(t *testing.T) {
	file, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(file)

	user := User{}
	decoder.Decode(&user)
	fmt.Println(user)
}
