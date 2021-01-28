package dice

import (
	"fmt"
	"strings"
)

// DiceRoll represents the complete formula for rolling dice.
type DiceRoll struct {
	// The sets of dice to roll.
	Dice []DiceSet
	// A value to add to the result of the rolled dice.
	Modifier int
}

// Roll returns a random value by rolling the sets of dice and adding an
// optional modifier.
func (d DiceRoll) Roll() int {
	return -1
}

func (d DiceRoll) String() string {
	var builder strings.Builder

	if len(d.Dice) > 0 {
		builder.WriteString(d.Dice[0].String())

		for _, diceSet := range d.Dice[1:] {
			builder.WriteString(" + ")
			builder.WriteString(diceSet.String())
		}
	}

	if d.Modifier != 0 {
		builder.WriteString(fmt.Sprintf(" + %v", d.Modifier))
	}

	return builder.String()
}

// ParseDiceRoll parses a string in the format "XdY +... C" into a Rollable
// type.
func ParseDiceRoll(notation string) (d *DiceRoll, err error) {
	return &DiceRoll{nil, 0}, nil
}
