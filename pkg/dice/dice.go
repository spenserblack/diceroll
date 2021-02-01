// Package dice defines types for rolling dice.
package dice

import "fmt"

// Rollable is the interface implemented by values that can be "rolled," such
// as values that represent dice.
type Rollable interface {
	// Roll returns a random value.
	Roll() int
}

// TooFewSidesError is an error type for when a rollable value has too few sides
// to properly create a random value.
type TooFewSidesError int

func (err TooFewSidesError) Error() string {
	return fmt.Sprintf("diceroll: creation of rollable with <2 sides: %d", int(err))
}
