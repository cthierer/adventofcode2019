package rps

type Rock struct{}

func (r Rock) Score() int {
	return 1
}

func (r Rock) Compare(other interface{}) int {
	switch other.(type) {
	case Scissors:
		return 1
	case Paper:
		return -1
	default:
		return 0
	}
}

type Paper struct{}

func (p Paper) Score() int {
	return 2
}

func (p Paper) Compare(other interface{}) int {
	switch other.(type) {
	case Scissors:
		return -1
	case Rock:
		return 1
	default:
		return 0
	}
}

type Scissors struct{}

func (s Scissors) Score() int {
	return 3
}

func (s Scissors) Compare(other interface{}) int {
	switch other.(type) {
	case Paper:
		return 1
	case Rock:
		return -1
	default:
		return 0
	}
}
