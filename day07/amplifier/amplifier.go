package amplifier

import "github.com/cthierer/adventofcode2019/day07/intcode"

// Callback is called when a amplifier is run.
type Callback func(intcode.Scanner, intcode.Writer) error

// Amplifier represents a signal amplifier.
type Amplifier struct {
	PhaseSetting int
	next         *Amplifier
}

// Collection tracks an entier network of amplifiers.
type Collection struct {
	root *Amplifier
	curr *Amplifier
}

// Add inserts a new amplifier into the network.
func (c *Collection) Add(amplifier *Amplifier) {
	if c.root == nil {
		c.root = amplifier
		c.curr = c.root
		return
	}

	c.curr.next = amplifier
	c.curr = c.curr.next
}

// Run runs a program on a collection of amplifiers.
func (c *Collection) Run(signalIn Signal, runProgram Callback) (Signal, error) {
	amp := c.root
	nextSignal := signalIn

	for amp != nil {
		input := valueQueue{Values: []int{amp.PhaseSetting, nextSignal.Value}}

		err := runProgram(&input, &nextSignal)
		if err != nil {
			return Signal{}, err
		}

		amp = amp.next
	}

	return nextSignal, nil
}
