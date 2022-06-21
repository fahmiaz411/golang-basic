package encoding

import (
	"encoding/base64"
	"fmt"
	"log"
)

func Base64StdEncoding() {
	sample := "a@"
	encodedString := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Println(encodedString) // YUA=

	originalStringBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes)) // a@
}

func Base64URLEncoding() {
	sample := "�"

	encodedStringURL := base64.URLEncoding.EncodeToString([]byte(sample))
	fmt.Printf("URL Encoding: %s\n", encodedStringURL) // 77-9
	encodedStringSTD := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Printf("STD Encoding: %s\n", encodedStringSTD) // 77+9

	originalStringBytes, err := base64.URLEncoding.DecodeString(encodedStringURL)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes)) // �
}