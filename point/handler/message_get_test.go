package handler

import (
	"fmt"
	"testing"
)

func TestTextMsg(t *testing.T) {
	r := TextMsg("sample")
	r_type := fmt.Sprintf("%T", r)
	if r_type != "http.HandlerFunc" {
		t.Fatalf("Wrong type")
	}
}
