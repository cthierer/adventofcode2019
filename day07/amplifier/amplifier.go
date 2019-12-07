package amplifier

import (
	"fmt"
	"sync"

	"github.com/cthierer/adventofcode2019/day07/intcode"
)

// Callback is called when a amplifier is run.
type Callback func(intcode.Scanner, intcode.Writer) error

// Amplifier represents a signal amplifier.
type Amplifier struct {
	PhaseSetting int
	next         *Amplifier
	output       valueQueue
}

// Run a program on this amplifier.
func (a *Amplifier) Run(program *intcode.Program, input *valueQueue) error {
	input.Shift(a.PhaseSetting)
	return program.Execute(input, &a.output)
}

// Collection tracks an entier network of amplifiers.
type Collection struct {
	root *Amplifier
	last *Amplifier
}

// Add inserts a new amplifier into the network.
func (c *Collection) Add(amp *Amplifier) {
	if c.root == nil {
		c.root = amp
		c.last = c.root
		return
	}

	c.last.next = amp
	c.last = c.last.next
}

// Run runs a program on a collection of amplifiers.
func (c *Collection) Run(program *intcode.Program) (int, error) {
	var wg sync.WaitGroup
	amp := c.root
	nextInput := &c.last.output
	nextInput.Shift(0)

	for amp != nil {
		wg.Add(1)
		go func(a *Amplifier, i *valueQueue) {
			err := a.Run(program.Snapshot(), i)
			if err != nil {
				fmt.Printf("amplifier w/ phase %d failed: %v", a.PhaseSetting, err)
			}
			wg.Done()
		}(amp, nextInput)
		nextInput = &amp.output
		amp = amp.next
	}

	wg.Wait()

	return c.last.output.Peek(), nil
}
