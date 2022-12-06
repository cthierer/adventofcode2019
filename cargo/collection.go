package cargo

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Collection struct {
	stacks map[string]*Stack
}

func NewCollection() *Collection {
	stacks := make(map[string]*Stack)
	return &Collection{stacks}
}

func ParseCollection(value string) (*Collection, error) {
	c := NewCollection()
	lines := strings.Split(value, "\n")
	headerRow := lines[len(lines)-1]
	crateRows := lines[0 : len(lines)-1]

	patternCrateID, err := regexp.Compile("\\w+")
	if err != nil {
		return nil, err
	}

	stackIDs := patternCrateID.FindAllString(headerRow, -1)
	if stackIDs == nil {
		return c, nil
	}

	stacks := make([]*Stack, len(stackIDs))
	for i, id := range stackIDs {
		stacks[i] = &Stack{ID: id}
	}

	for i := len(crateRows) - 1; i >= 0; i -= 1 {
		row := crateRows[i]
		for j := 0; j < len(row); j += 4 {
			stackIdx := j / 4
			stack := stacks[stackIdx]

			crateValue := string(row[j+1])
			if crateValue != " " {
				stack.Push(Crate(crateValue))
			}
		}
	}

	for _, s := range stacks {
		c.Add(s)
	}

	return c, nil
}

func (c *Collection) Add(s *Stack) {
	c.stacks[s.ID] = s
}

func (c *Collection) Get(id string) *Stack {
	s, ok := c.stacks[id]
	if !ok {
		return nil
	}
	return s
}

type Command interface {
	FromStack() string
	ToStack() string
	Quantity() int
}

func (c *Collection) Transfer(cmd Command) error {
	fromStack := c.Get(cmd.FromStack())
	if fromStack == nil {
		return fmt.Errorf("invalid stack ID: %v", cmd.FromStack())
	}

	toStack := c.Get(cmd.ToStack())
	if toStack == nil {
		return fmt.Errorf("invalid stack ID: %v", cmd.ToStack())
	}

	for i := 0; i < cmd.Quantity(); i += 1 {
		crate := fromStack.Pop()
		if crate == NoCrate {
			return errors.New("invalid system state: cannot transfer from an empty stack")
		}
		toStack.Push(crate)
	}

	return nil
}

func (c *Collection) Values() []*Stack {
	v := make([]*Stack, len(c.stacks))
	i := 0
	keys := make([]string, len(c.stacks))
	for k := range c.stacks {
		keys[i] = k
		i += 1
	}

	sort.Strings(keys)
	i = 0
	for _, k := range keys {
		v[i] = c.stacks[k]
		i += 1
	}

	return v
}
