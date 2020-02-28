package auth

import (
	"crypto/sha256"
	"fmt"
)

func AddUser() {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x", h.Sum(nil))
}
