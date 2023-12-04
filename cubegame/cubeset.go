package cubegame

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidSegmentFormat = errors.New("segment is malformed")

type CubeSet struct {
	counts map[Color]int
}

func NewCubeSet(red, green, blue int) *CubeSet {
	set := CubeSet{}
	set.counts = make(map[Color]int)

	set.counts[Red] = red
	set.counts[Green] = green
	set.counts[Blue] = blue

	return &set
}

func (set *CubeSet) Colors() []Color {
	colors := make([]Color, len(set.counts))
	idx := 0

	for color := range set.counts {
		colors[idx] = color
		idx += 1
	}

	return colors
}

func (set *CubeSet) GetCount(color Color) int {
	count, ok := set.counts[color]
	if !ok {
		return 0
	}
	return count
}

func (set *CubeSet) SetCount(color Color, count int) {
	if set.counts == nil {
		set.counts = make(map[Color]int)
	}

	set.counts[color] = count
}

func (set *CubeSet) String() string {
	parts := make([]string, len(set.counts))
	idx := 0

	for color, count := range set.counts {
		parts[idx] = fmt.Sprintf("%v %v", count, color)
		idx += 1
	}

	return strings.Join(parts, ", ")
}

func ParseCubeSet(csv string) (*CubeSet, error) {
	parsed := CubeSet{}
	parsed.counts = make(map[Color]int)

	segments := strings.Split(csv, ",")
	for idx, segment := range segments {
		color, count, err := parseCubeSetSegment(segment)
		if err != nil {
			return nil, fmt.Errorf("unable to parse segment at index %v: %w", idx, err)
		}
		parsed.counts[color] += count
	}

	return &parsed, nil
}

func parseCubeSetSegment(value string) (color Color, count int, err error) {
	trimmed := strings.TrimSpace(value)

	parts := strings.SplitN(trimmed, " ", 2)
	if len(parts) < 2 {
		err = ErrInvalidSegmentFormat
		return
	}

	count, err = parseCount(parts[0])
	if err != nil {
		return
	}

	color, err = ParseColor(parts[1])
	if err != nil {
		return
	}

	return
}

func parseCount(value string) (int, error) {
	return strconv.Atoi(value)
}
