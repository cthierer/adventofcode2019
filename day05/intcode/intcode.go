package intcode

import (
	"fmt"
	"math"
)

// Program is an intcode executable.
type Program struct {
	Instructions []int
	position     int
	halted       bool
}

func (p *Program) parameter(offset int) (int, error) {
	raw := p.Instructions[p.position+offset]

	mode := int(math.Mod(math.Trunc(float64(p.Instructions[p.position])/math.Pow(10, float64(offset+1))), 10))
	switch mode {
	case 0:
		// position
		return p.Instructions[raw], nil
	case 1:
		// immediate
		return raw, nil
	default:
		// error
		return 0, fmt.Errorf("unknown parameter mode: %d", mode)
	}
}

func (p *Program) store(offset int, val int) {
	idx := p.Instructions[p.position+offset]
	p.Instructions[idx] = val
}

func (p *Program) add() {
	val1, _ := p.parameter(1)
	val2, _ := p.parameter(2)
	res := val1 + val2

	p.store(3, res)
	p.position += 4
}

func (p *Program) multiply() {
	val1, _ := p.parameter(1)
	val2, _ := p.parameter(2)
	res := val1 * val2

	p.store(3, res)
	p.position += 4
}

func (p *Program) scan() {
	var val int

	fmt.Print("< ")
	fmt.Scanf("%d", &val)

	p.store(1, val)
	p.position += 2
}

func (p *Program) print() {
	val, _ := p.parameter(1)

	fmt.Printf("> %d\n", val)

	p.position += 2
}

func (p *Program) jumpIfTrue() {
	val, _ := p.parameter(1)
	if val != 0 {
		jumpTo, _ := p.parameter(2)
		p.position = jumpTo
		return
	}

	p.position += 3
}

func (p *Program) jumpIfFalse() {
	val, _ := p.parameter(1)
	if val == 0 {
		jumpTo, _ := p.parameter(2)
		p.position = jumpTo
		return
	}

	p.position += 3
}

func (p *Program) lessThan() {
	val1, _ := p.parameter(1)
	val2, _ := p.parameter(2)
	if val1 < val2 {
		p.store(3, 1)
	} else {
		p.store(3, 0)
	}
	p.position += 4
}

func (p *Program) equals() {
	val1, _ := p.parameter(1)
	val2, _ := p.parameter(2)
	if val1 == val2 {
		p.store(3, 1)
	} else {
		p.store(3, 0)
	}
	p.position += 4
}

func (p *Program) halt() {
	p.halted = true
	p.position++
}

func (p *Program) opcode() int {
	instr := p.Instructions[p.position]
	return instr % 100
}

func (p *Program) execute() error {
	opcode := p.opcode()
	switch opcode {
	case 1:
		p.add()
	case 2:
		p.multiply()
	case 3:
		p.scan()
	case 4:
		p.print()
	case 5:
		p.jumpIfTrue()
	case 6:
		p.jumpIfFalse()
	case 7:
		p.lessThan()
	case 8:
		p.equals()
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
func (p *Program) Next() error {
	err := p.execute()
	if err != nil {
		return err
	}
	return nil
}

// Run executes the provided instructions.
func Run(instructions []int) ([]int, error) {
	p := &Program{Instructions: instructions}

	for p.HasNext() {
		err := p.Next()
		if err != nil {
			return nil, err
		}
	}

	if !p.halted {
		return nil, fmt.Errorf("program failed to halt: failed at index %d", p.position)
	}

	return p.Instructions, nil
}
