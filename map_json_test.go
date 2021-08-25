package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "001",
		"nama":  "laptop",
		"price": 100000,
	}

	byte, _ := json.Marshal(product)
	fmt.Println(string(byte))
}

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id":"001","nama":"laptop","price":100000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonByte, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
}
