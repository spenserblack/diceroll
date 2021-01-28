package dice

import "fmt"

// DiceSet is a set of dice (e.g. 3d4) that can be rolled.
type DiceSet struct {
	count int
	// Of is the type of dice being used in the DiceSet. For example, a set "of"
	// d6 dice.
	Of Die
}

// Roll returns a random value by rolling a set of dice.
func (d DiceSet) Roll() int {
	return -1
}

func (d DiceSet) String() string {
	return fmt.Sprintf("%v%v", d.count, d.Of)
}

// ParseDiceSet parses a string in the format "XdY" into a Rollable type.
func ParseDiceSet(notation string) (d *DiceSet, err error) {
	return &DiceSet{0, 0}, nil
}
