package dice

import "testing"

// TestDieParsing creates a Die with the notation "dY" and confirms
// that the notation has been parsed correctly.
func TestDieParsing(t *testing.T) {
	die, err := NewDie("d12")

	if err != nil {
		t.Fatalf(`err = %v, want nil`, err)
	}

	if die != 12 {
		t.Fatalf(`NewDie("d12") = %v, want 12`, die)
	}
}
