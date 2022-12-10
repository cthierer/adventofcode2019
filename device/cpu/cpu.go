package cpu

import "fmt"

type Instruction interface {
	Next() bool
	Run(c *Processor)
}

type AddX struct {
	callCount int
	Value     int64
}

func (a *AddX) Next() bool {
	return a.callCount < 2
}

func (a *AddX) Run(c *Processor) {
	a.callCount += 1
	if a.callCount == 2 {
		c.x += a.Value
	}
}

func (a *AddX) String() string {
	return fmt.Sprintf("ADD  %v", a.Value)
}

type Noop struct {
	called bool
}

func (n *Noop) Next() bool {
	return !n.called
}

func (n *Noop) Run(c *Processor) {
	n.called = true
}

func (n *Noop) String() string {
	return "NOOP"
}

type Processor struct {
	cycle int
	x     int64
	i     Instruction
}

func (c *Processor) Cycle() int {
	return c.cycle + 1
}

func (c *Processor) X() int64 {
	if c.cycle == 0 {
		return 1
	}
	return c.x
}

func (c *Processor) Load(i Instruction) {
	if c.cycle == 0 {
		c.x = 1
	}
	c.i = i
}

func (c *Processor) Next() bool {
	if c.i == nil || !c.i.Next() {
		return false
	}

	c.cycle += 1
	c.i.Run(c)
	return true
}

func (c *Processor) String() string {
	return fmt.Sprintf("Cycle=%v, X=%v", c.cycle, c.X())
}
