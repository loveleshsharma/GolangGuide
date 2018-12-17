package types

import (
	"testing"
)

func TestNewPersonWithValues(t *testing.T) {
	person := NewPersonWithValues("Lovelesh", "Male", 20)

	if person.Name != "Lovelesh" {
		t.Errorf("Expected name Lovelesh but got %v", person.Name)
	}
}
