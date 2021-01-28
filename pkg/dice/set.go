package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

// Set is a set of dice (e.g. 3d4) that can be rolled.
type Set struct {
	count int
	// Of is the type of dice being used in the Set. For example, a set "of"
	// d6 dice.
	Of Die
}

var setRegex = regexp.MustCompile(`^\s*(\d+)(.+)$`)

// Roll returns a random value by rolling a set of dice.
func (d Set) Roll() int {
	return -1
}

func (d Set) String() string {
	return fmt.Sprintf("%v%v", d.count, d.Of)
}

// ParseSet parses a string in the format "XdY" into a Rollable type.
func ParseSet(notation string) (set Set, err error) {
	matches := setRegex.FindStringSubmatch(notation)

	if len(matches) != 3 {
		err = fmt.Errorf("parsing %q: invalid syntax", notation)
	} else {
		set.count, err = strconv.Atoi(matches[1])
		if err != nil {
			return set, fmt.Errorf("parsing %q: %w", notation, err)
		}
		set.Of, err = ParseDie(matches[2])
		if err != nil {
			return set, fmt.Errorf("parsing %q: %w", notation, err)
		}
	}
	return
}
