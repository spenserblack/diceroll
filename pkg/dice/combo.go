package dice

import (
	"fmt"
	"strings"
)

// Combo represents a combination of sets of dice to roll.
type Combo struct {
	// The sets of dice to roll.
	Dice []Set
	// A value to add to the result of the rolled dice.
	Modifier int
}

// Roll returns a random value by rolling the sets of dice and adding an
// optional modifier.
func (d Combo) Roll() int {
	return -1
}

func (d Combo) String() string {
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

// ParseCombo parses a string in the format "XdY +... C" into a Rollable
// type.
func ParseCombo(notation string) (d *Combo, err error) {
	return &Combo{nil, 0}, nil
}
