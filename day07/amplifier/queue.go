package amplifier

import "sync"

type valueQueue struct {
	Values   []int
	position int
	mux      sync.Mutex
}

func (q *valueQueue) next() (bool, int) {
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.position >= len(q.Values) {
		return false, 0
	}
	val := q.Values[q.position]
	q.position++
	return true, val
}

func (q *valueQueue) ScanInt(c chan int) {
	var hasVal bool
	var val int
	for !hasVal {
		hasVal, val = q.next()
	}
	c <- val
}

func (q *valueQueue) WriteInt(val int) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.Values = append(q.Values, val)
}

func (q *valueQueue) Shift(val int) {
	q.mux.Lock()
	defer q.mux.Unlock()
	newValues := make([]int, 1, len(q.Values)+1)
	newValues[0] = val
	newValues = append(newValues, q.Values...)
	q.Values = newValues
}

func (q *valueQueue) Peek() int {
	q.mux.Lock()
	defer q.mux.Unlock()
	return q.Values[len(q.Values)-1]
}
