package cubegame

import "errors"

type Color string

func (color Color) String() string {
	return string(color)
}

var ErrUnknownColor = errors.New("unknown color")

const (
	Blue  = Color("blue")
	Red   = Color("red")
	Green = Color("green")
)

func ParseColor(value string) (Color, error) {
	parsed := Color(value)

	switch parsed {
	case Blue:
		return Blue, nil
	case Red:
		return Red, nil
	case Green:
		return Green, nil
	}

	return parsed, ErrUnknownColor
}
