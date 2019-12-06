package intcode

import (
	"fmt"
)

// Program is an intcode executable.
type Program struct {
	Instructions []int
	position     int
	halted       bool
}

func (p *Program) clone() *Program {
	cloned := Program{
		Instructions: make([]int, len(p.Instructions)),
		position:     p.position,
	}
	for i, v := range p.Instructions {
		cloned.Instructions[i] = v
	}
	return &cloned
}

func (p *Program) add() {
	idxVal1 := p.Instructions[p.position+1]
	idxVal2 := p.Instructions[p.position+2]
	idxRes := p.Instructions[p.position+3]

	val1 := p.Instructions[idxVal1]
	val2 := p.Instructions[idxVal2]
	res := val1 + val2

	p.Instructions[idxRes] = res
	p.position += 4
}

func (p *Program) multiply() {
	idxVal1 := p.Instructions[p.position+1]
	idxVal2 := p.Instructions[p.position+2]
	idxRes := p.Instructions[p.position+3]

	val1 := p.Instructions[idxVal1]
	val2 := p.Instructions[idxVal2]
	res := val1 * val2

	p.Instructions[idxRes] = res
	p.position += 4
}

func (p *Program) halt() {
	p.halted = true
}

func (p *Program) execute() error {
	opcode := p.Instructions[p.position]
	switch opcode {
	case 1:
		p.add()
	case 2:
		p.multiply()
	case 99:
		p.halt()
	default:
		return fmt.Errorf(fmt.Sprintf("unknown opcode: %d", opcode))
	}

	return nil
}

// HasNext evaluates if the program has another set of instructions.
func (p *Program) HasNext() bool {
	return !p.halted && p.position < len(p.Instructions)
}

// Next executes the next instruction.
func (p *Program) Next() (*Program, error) {
	next := p.clone()
	err := next.execute()
	if err != nil {
		return nil, err
	}
	return next, nil
}

// Run executes the provided instructions.
func Run(instructions []int) ([]int, error) {
	p := &Program{Instructions: instructions}

	for p.HasNext() {
		pn, err := p.Next()
		if err != nil {
			return nil, err
		}
		p = pn
	}

	if !p.halted {
		return nil, fmt.Errorf("program failed to halt: failed at index %d", p.position)
	}

	return p.Instructions, nil
}
