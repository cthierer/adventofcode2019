package amplifier

type valueQueue struct {
	Values   []int
	position int
}

func (q *valueQueue) ScanInt() int {
	val := q.Values[q.position]
	q.position++
	return val
}
