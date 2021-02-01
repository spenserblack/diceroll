package main

import (
	"fmt"
	"github.com/spenserblack/diceroll/pkg/dice"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args[1:]) == 0 || os.Args[1] == "help" {
		println("Usage: diceroll <dice notation>")
		return
	}

	var rollable dice.Rollable
	var err error

	allDice := strings.Join(os.Args[1:], "")
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
