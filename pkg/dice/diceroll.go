package dice

type DiceRoll struct {
	Dice     []DiceSet
	Modifier int
}

func (d DiceRoll) Roll() int {
	return -1
}

func NewDiceRoll(notation string) (d *DiceRoll, err error) {
	return &DiceRoll{nil, 0}, nil
}
