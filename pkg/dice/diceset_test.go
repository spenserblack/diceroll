package dice

import "testing"

// TestDiceSetParsing creates a DiceSet with the notation "XdY" and confirms
// that the notation has been parsed correctly.
func TestDiceSetParsing(t *testing.T) {
	diceSet, err := ParseDiceSet("3d8")

	if err != nil {
		t.Fatalf(`err = %v, want nil`, err)
	}

	if c := diceSet.count; c != 3 {
		t.Fatalf(`diceSet.count = %v, want 3`, c)
	}

	if die := diceSet.Of; die != 8 {
		t.Fatalf(`diceSet.Of = %v, want 8`, die)
	}
}

// TestDiceSetString tests that a DiceSet has the expected "XdY" notation when
// converted to a string.
func TestDiceSetString(t *testing.T) {
	want := "1d20"

	if s := (DiceSet{1, 20}).String(); s != want {
		t.Fatalf(`(DiceSet{1, 20}).String() = %q, want %q`, s, want)
	}
}
