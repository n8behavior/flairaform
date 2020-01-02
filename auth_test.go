package main

import (
	"fmt"
	"testing"
)

func TestAuthorize(t *testing.T) {
	auth, _ := Authorize()
	fmt.Println(auth.Token)
}
