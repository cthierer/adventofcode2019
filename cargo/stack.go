package cargo

type Crate string

func (c Crate) String() string {
	return string(c)
}

const NoCrate = Crate("")

type Stack struct {
	ID        string
	crates    []Crate
	hasCrates bool
}

func (s *Stack) Push(c ...Crate) error {
	s.crates = append(s.crates, c...)
	s.hasCrates = true
	return nil
}

func (s *Stack) Pop(n int) []Crate {
	if !s.hasCrates {
		return nil
	}

	c := s.crates[len(s.crates)-n:]
	s.crates = s.crates[0 : len(s.crates)-n]

	if len(s.crates) == 0 {
		s.hasCrates = false
	}

	return c
}

func (s *Stack) Peek() Crate {
	if !s.hasCrates {
		return NoCrate
	}
	return s.crates[len(s.crates)-1]
}

func (s *Stack) Values() []Crate {
	if !s.hasCrates {
		return []Crate{}
	}

	v := make([]Crate, len(s.crates))
	for i := 0; i < len(s.crates); i += 1 {
		v[i] = s.crates[i]
	}

	return v
}
