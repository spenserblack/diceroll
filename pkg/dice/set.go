package dice

import "fmt"

// Set is a set of dice (e.g. 3d4) that can be rolled.
type Set struct {
	count int
	// Of is the type of dice being used in the Set. For example, a set "of"
	// d6 dice.
	Of Die
}

// Roll returns a random value by rolling a set of dice.
func (d Set) Roll() int {
	return -1
}

func (d Set) String() string {
	return fmt.Sprintf("%v%v", d.count, d.Of)
}

// ParseSet parses a string in the format "XdY" into a Rollable type.
func ParseSet(notation string) (set Set, err error) {
	return Set{0, 0}, nil
}
