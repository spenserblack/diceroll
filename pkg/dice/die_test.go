package dice

import "testing"

// TestDieParsing creates a Die with the notation "dY" and confirms
// that the notation has been parsed correctly.
func TestDieParsing(t *testing.T) {
	die, err := ParseDie("d12")

	if err != nil {
		t.Fatalf(`err = %v, want nil`, err)
	}

	if die != 12 {
		t.Fatalf(`ParseDie("d12") = %v, want 12`, die)
	}
}

// TestDieString tests that a Die has the expected "dY" notation when converted
// to a string.
func TestDieString(t *testing.T) {
	if s := Die(6).String(); s != "d6" {
		t.Fatalf(`Die(6).String() = %q, want "d6"`, s)
	}
}

// TestDieParsingErr cretes a Die without enough sides (e.g. a d1) and checks
// that an error is returned.
func TestDieParsingErr(t *testing.T) {
	want := TooFewSidesError(1)
	_, err := ParseDie("d1")

	if err != want {
		t.Fatalf(`err = %v, want %v`, err, want)
	}
}
