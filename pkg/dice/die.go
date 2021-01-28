package dice

import "fmt"

// Die represents a single die (e.g. a d6).
type Die int

// Roll returns a random value by rolling the die. Will always be between 1
// and the die value, inclusively.
func (d Die) Roll() int {
	return -1
}

func (d Die) String() string {
	return fmt.Sprintf("d%v", rune(d))
}

// ParseDie parses a string in the format "dX" into a Rollable type.
func ParseDie(notation string) (d Die, err error) {
	return
}
