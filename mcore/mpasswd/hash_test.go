package mpasswd

import (
	"testing"
)

func TestHashPasswd(t *testing.T) {
	passwd := "demo"

	hashPasswd, _ := GenerateHashPassword(passwd)

	if CompareHashAndPassword(hashPasswd, "demo") != nil {
		t.Error("hashpasswd and provide passwd not equal.")
	}
}
