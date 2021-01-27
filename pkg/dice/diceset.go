package dice

type DiceSet struct {
	count int
	Of    Die
}

func (d *DiceSet) Roll() int {
	return -1
}

func NewDiceSet(notation string) (d *DiceSet, err error) {
	return &DiceSet{0, 0}, nil
}
