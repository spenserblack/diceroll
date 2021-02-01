package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

// Die represents a single die (e.g. a d6).
type Die int

var dieRegex = regexp.MustCompile(`^\s*d(\d+)\s*$`)

// Roll returns a random value by rolling the die. Will always be between 1
// and the die value, inclusively.
func (d Die) Roll() int {
	return rand.Intn(int(d)) + 1
}

func (d Die) String() string {
	return fmt.Sprintf("d%v", rune(d))
}

// ParseDie parses a string in the format "dX" into a Rollable type.
func ParseDie(notation string) (d Die, err error) {
	matches := dieRegex.FindStringSubmatch(notation)

	if len(matches) != 2 {
		err = fmt.Errorf("parsing %q: invalid syntax", notation)
	} else {
		var n int
		n, err = strconv.Atoi(matches[1])
		d = Die(n)
	}

	return
}
