package dice

import "testing"

// TestDiceSetParsing creates a DiceSet with the notation "XdY" and confirms
// that the notation has been parsed correctly.
func TestDiceSetParsing(t *testing.T) {
	diceSet, err := NewDiceSet("3d8")

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
