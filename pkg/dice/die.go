package dice

import "fmt"

type Die int

func (d Die) Roll() int {
	return -1
}

func (d Die) String() string {
	return fmt.Sprintf("d%v", rune(d))
}

func NewDie(notation string) (d Die, err error) {
	return
}
