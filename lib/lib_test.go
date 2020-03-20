package trygofuzz

import (
	"testing"
)

func TestChcker(t *testing.T) {
	if !Checker("message") {
		t.Fatal("Checker(message) should be true")
	}
}
