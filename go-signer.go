package main

import (
	"crypto/sha512"
	"encoding"
	"fmt"
	//"encoding/base64"
)

func main() {
	hasher := sha512.New()
	hasher.Write([]byte("this is sample text"))

	h := hasher.(encoding.BinaryMarshaler)
	encoding, err := h.MarshalBinary()

	if err != nil {
		fmt.Printf("failed encoding, err: %w\n", err)
	}
	fmt.Println(encoding)

	sha := //base64.URLEncoding.EncodeToString
		(hasher.Sum(nil))
	fmt.Println(sha)
}
