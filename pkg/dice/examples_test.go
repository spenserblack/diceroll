package dice_test

import (
	"fmt"
	"github.com/spenserblack/diceroll/pkg/dice"
)

func ExampleParseDie() {
	die, err := dice.ParseDie("d6")

	if err != nil {
		panic(err)
	}

	fmt.Println(int(die))
	// Output: 6
}

func ExampleDie_Roll() {
	die, err := dice.ParseDie("d6")

	if err != nil {
		panic(err)
	}

	roll := die.Roll()
	fmt.Println(1 <= roll && roll <= 6)
	// Output: true
}

func ExampleParseSet() {
	diceSet, err := dice.ParseSet("3d6")

	if err != nil {
		panic(err)
	}

	count, of := diceSet.Count, int(diceSet.Of)

	fmt.Printf("%v %v-sided dice\n", count, of)
	// Output: 3 6-sided dice
}

func ExampleSet_Roll() {
	diceSet, err := dice.ParseSet("3d6")

	if err != nil {
		panic(err)
	}

	roll := diceSet.Roll()
	fmt.Println(3 <= roll && roll <= 18)
	// Output: true
}

func ExampleParseCombo() {
	// Whitespace is ignored
	combo, err := dice.ParseCombo("1d6 + 2d4+3d8     + 5")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%vd%v", combo.Dice[0].Count, int(combo.Dice[0].Of))
	fmt.Printf(" + %vd%v", combo.Dice[1].Count, int(combo.Dice[1].Of))
	fmt.Printf(" + %vd%v", combo.Dice[2].Count, int(combo.Dice[2].Of))
	fmt.Printf(" + %v", combo.Modifier)
	// Output: 1d6 + 2d4 + 3d8 + 5
}

func ExampleCombo_Roll() {
	combo, err := dice.ParseCombo("2d4 + 2")

	if err != nil {
		panic(err)
	}

	roll := combo.Roll()
	fmt.Println(4 <= roll && roll <= 10)
	// Output: true
}
