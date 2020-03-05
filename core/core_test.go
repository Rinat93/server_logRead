package core

import (
	"testing"
)

func Test_User(t *testing.T) {
	core := new(SCore)
	core.Connect()
	defer core.Net.Close()
	core.SetUser("admin", "930114")
}
