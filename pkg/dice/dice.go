// Package dice defines types for rolling dice.
package dice

// Rollable is the interface implemented by values that can be "rolled," such
// as values that represent dice.
type Rollable interface {
	// Roll returns a random value.
	Roll() int
}
