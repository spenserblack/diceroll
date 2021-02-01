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

	if dice := diceRoll.Dice[0]; !(dice.Count == 2 && dice.Of == 6) {
		t.Fatalf(`diceRoll.Dice[0] = %+v, want 2d6`, dice)
	}

	if dice := diceRoll.Dice[1]; !(dice.Count == 3 && dice.Of == 4) {
		t.Fatalf(`diceRoll.Dice[1] = %+v, want 3d4`, dice)
	}

	if modifier := diceRoll.Modifier; modifier != 5 {
		t.Fatalf(`diceRoll.Modifier = %v, want 5`, modifier)
	}
}

// TestComboString tests that a Combo has the expected "XdY +... C"
// notation when converted to a string.
func TestComboString(t *testing.T) {
	diceRoll := Combo{[]Set{{3, 4}, {1, 6}}, 2}

	want := "3d4 + 1d6 + 2"
	if s := diceRoll.String(); s != want {
		t.Fatalf(`diceRoll.String() = %q, want %q`, s, want)
	}
}

// TestComboParsingErr cretes a dice Combo with a set of dice with a die without
// enough sides (e.g. a d1) and checks that an error is returned.
func TestComboParsingErr(t *testing.T) {
	want := TooFewSidesError(1)
	_, err := ParseCombo("3d2 + 2d1 + 2")

	if err != want {
		t.Fatalf(`err = %v, want %v`, err, want)
	}
}
