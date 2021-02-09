package main

import (
	"crypto/rand"
	"fmt"
)

//APITokenGenerator generates a rondom api key
func APITokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
