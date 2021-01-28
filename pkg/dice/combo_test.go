package dice

import "testing"

// TestCombo creates a Combo with the notation "XdY +... C" and confirms
// that the notation has been parsed correctly.
func TestComboParsing(t *testing.T) {
	diceRoll, err := ParseCombo("2d6 + 3d4 + 5")

	if err != nil {
		t.Fatalf(`err = %v, want nil`, err)
	}

	if l := len(diceRoll.Dice); l != 2 {
		t.Fatalf(`len(diceRoll.Dice) = %v, want 2`, l)
	}

	for _, ds := range diceRoll.Dice {
		if !((ds.count == 2 && ds.Of == 6) || (ds.count == 3 && ds.Of == 4)) {
			t.Fatalf(`ds = %+v, want 2d6 or 3d4`, ds)
		}
	}

	if modifier := diceRoll.Modifier; modifier != 5 {
		t.Fatalf(`diceRoll.Modifier = %v, want 5`, modifier)
	}
}

// TestComboString tests that a Combo has the expected "XdY +... C"
// notation when converted to a string.
func TestComboString(t *testing.T) {
	diceRoll := Combo{[]Set{Set{3, 4}, Set{1, 6}}, 2}

	want := "3d4 + 1d6 + 2"
	if s := diceRoll.String(); s != want {
		t.Fatalf(`diceRoll.String() = %q, want %q`, s, want)
	}
}
