package main

import (
	"crypto/rand"
	"log"

	// "encoding/base64"
	// "encoding/hex"
	"fmt"
	// ...
)

func Key() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		log.Fatalf("Out of randomness, should never happen: %v", err)
	}
	return fmt.Sprintf("%x", buf)
	// or hex.EncodeToString(buf)
	// or base64.StdEncoding.EncodeToString(buf)
}

func main() {
	fmt.Println(Key())
}
