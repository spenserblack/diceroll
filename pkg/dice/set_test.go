package dice

import "testing"

// TestSetParsing creates a Set with the notation "XdY" and confirms
// that the notation has been parsed correctly.
func TestSetParsing(t *testing.T) {
	diceSet, err := ParseSet("3d8")

	if err != nil {
		t.Fatalf(`err = %v, want nil`, err)
	}

	if c := diceSet.Count; c != 3 {
		t.Fatalf(`diceSet.Count = %v, want 3`, c)
	}

	if die := diceSet.Of; die != 8 {
		t.Fatalf(`diceSet.Of = %v, want 8`, die)
	}
}

// TestSetString tests that a Set has the expected "XdY" notation when
// converted to a string.
func TestSetString(t *testing.T) {
	want := "1d20"

	if s := (Set{1, 20}).String(); s != want {
		t.Fatalf(`(Set{1, 20}).String() = %q, want %q`, s, want)
	}
}

// TestSetParsingErr cretes a Set of dice without enough sides (e.g. a 3d1) and
// checks that an error is returned.
func TestSetParsingErr(t *testing.T) {
	want := TooFewSidesError(1)
	_, err := ParseSet("3d1")

	if err != want {
		t.Fatalf(`err = %v, want %v`, err, want)
	}
}
