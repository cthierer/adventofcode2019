package cargo

import "errors"

type Crate string

func (c Crate) String() string {
	return string(c)
}

const NoCrate = Crate("")

type Stack struct {
	ID        string
	crates    []Crate
	top       int
	hasCrates bool
}

func (s *Stack) Push(c Crate) error {
	if c == NoCrate {
		return errors.New("cannot push an empty crate onto the stack")
	}

	if s.top >= len(s.crates) {
		larger := make([]Crate, len(s.crates)*2+1)
		for i, v := range s.crates {
			larger[i] = v
		}
		s.crates = larger
	}

	s.crates[s.top] = c
	s.top += 1
	s.hasCrates = true
	return nil
}

func (s *Stack) Pop() Crate {
	if !s.hasCrates {
		return NoCrate
	}

	s.top -= 1
	c := s.crates[s.top]

	if s.top == 0 {
		s.hasCrates = false
	}

	return c
}

func (s *Stack) Peek() Crate {
	if !s.hasCrates {
		return NoCrate
	}
	return s.crates[s.top-1]
}

func (s *Stack) Values() []Crate {
	if !s.hasCrates {
		return []Crate{}
	}

	v := make([]Crate, s.top)
	for i := 0; i < s.top; i += 1 {
		v[i] = s.crates[i]
	}

	return v
}
