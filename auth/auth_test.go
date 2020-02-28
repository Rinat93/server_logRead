package auth

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"
)

func Test_User(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x", h.Sum(nil))
	fmt.Println(time.Now())
}
