package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	logJson("ali")
	logJson(1.1)
	logJson(false)
	logJson([]string{"ali", "mahmud"})
}
