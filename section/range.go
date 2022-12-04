package section

import (
	"errors"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	stop  int64
}

func NewRange(start, stop int64) Range {
	return Range{start, stop}
}

func ParseRange(value string) (Range, error) {
	parts := strings.Split(value, "-")
	if len(parts) < 2 {
		return Range{}, errors.New("unexpected range format")
	}

	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return Range{}, err
	}

	stop, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return Range{}, err
	}

	return NewRange(start, stop), nil
}

func (r Range) Start() int64 {
	return r.start
}

func (r Range) Stop() int64 {
	return r.stop
}

func (r Range) Contains(other Range) bool {
	return r.start <= other.start && r.stop >= other.stop
}

func (r Range) Overlaps(other Range) bool {
	return r.Contains(other) || (r.start <= other.stop && r.start >= other.start) || (r.stop >= other.start && r.stop <= other.stop)
}
