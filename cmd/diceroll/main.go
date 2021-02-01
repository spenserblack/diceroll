package main

import (
	"fmt"
	"github.com/spenserblack/diceroll/pkg/dice"
	"math/rand"
	"os"
	"strings"
	"time"
)

var allDice = strings.Join(os.Args[1:], "")

func main() {
	var rollable dice.Rollable
	var err error

	if rollable, err = dice.ParseDie(allDice); err == nil {
	} else if rollable, err = dice.ParseSet(allDice); err == nil {
	} else if rollable, err = dice.ParseCombo(allDice); err == nil {
	} else {
		panic(fmt.Errorf("Couldn't parse to a rollable type: %q", allDice))
	}

	fmt.Println(rollable.Roll())
}

func init() {
	rand.Seed(time.Now().Unix())
}
