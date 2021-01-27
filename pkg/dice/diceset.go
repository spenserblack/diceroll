package dice

import "fmt"

type DiceSet struct {
	count int
	Of    Die
}

func (d DiceSet) Roll() int {
	return -1
}

func (d DiceSet) String() string {
	return fmt.Sprintf("%v%v", d.count, d.Of)
}

func NewDiceSet(notation string) (d *DiceSet, err error) {
	return &DiceSet{0, 0}, nil
}
