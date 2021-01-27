package dice

import (
	"fmt"
	"strings"
)

type DiceRoll struct {
	Dice     []DiceSet
	Modifier int
}

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

func NewDiceRoll(notation string) (d *DiceRoll, err error) {
	return &DiceRoll{nil, 0}, nil
}
