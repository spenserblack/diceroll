package dice

import (
	"fmt"
	"strconv"
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
	var result int

	for _, set := range d.Dice {
		result += set.Roll()
	}

	return result + d.Modifier
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
func ParseCombo(notation string) (combo Combo, err error) {
	fields := strings.Split(notation, "+")

	// NOTE Most common format would be multiple sets of dice and once modifier.
	if n := len(fields) - 1; n > 0 {
		combo.Dice = make([]Set, 0, n)
	}

	for _, field := range fields {
		if n, err := strconv.Atoi(strings.TrimSpace(field)); err != nil {
			// Not a simple number -- try creating dice set.
			set, err := ParseSet(field)

			if err != nil {
				return combo, fmt.Errorf("parsing %q: %w", notation, err)
			}

			combo.Dice = append(combo.Dice, set)
		} else {
			combo.Modifier += n
		}
	}

	return
}
