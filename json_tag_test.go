package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id:       "001",
		Name:     "laptop",
		ImageURL: "http://exampe.com/image.jpg",
	}

	byte, _ := json.Marshal(product)
	fmt.Println(string(byte))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"001","name":"laptop","image_url":"http://exampe.com/image.jpg"}`
	jsonByte := []byte(jsonString)

	product := Product{}
	err := json.Unmarshal(jsonByte, &product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
