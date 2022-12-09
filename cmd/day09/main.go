package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func (p position) String() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
}

func distance(p1, p2 position) float64 {
	return math.Max(math.Abs(float64(p2.x-p1.x)), math.Abs(float64(p2.y-p1.y)))
}

type direction func(position) position

func up(p position) position {
	p.y += 1
	return p
}

func down(p position) position {
	p.y -= 1
	return p
}

func left(p position) position {
	p.x -= 1
	return p
}

func right(p position) position {
	p.x += 1
	return p
}

func none(p position) position {
	return p
}

func compose(directions ...direction) direction {
	return func(p position) position {
		for _, d := range directions {
			p = d(p)
		}
		return p
	}
}

func getDirection(head, tail position) direction {
	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	xDir := none
	yDir := none

	if deltaX < 0 {
		xDir = left
	} else if deltaX > 0 {
		xDir = right
	}

	if deltaY < 0 {
		yDir = down
	} else if deltaY > 0 {
		yDir = up
	}

	return compose(xDir, yDir)
}

type rope struct {
	headPositions []position
	tailPositions []position
}

func (r *rope) Head() position {
	if len(r.headPositions) == 0 {
		r.headPositions = append(r.headPositions, position{})
	}
	return r.headPositions[len(r.headPositions)-1]
}

func (r *rope) Tail() position {
	if len(r.tailPositions) == 0 {
		r.tailPositions = append(r.tailPositions, position{})
	}
	return r.tailPositions[len(r.tailPositions)-1]
}

func (r *rope) Move(d direction, n int64) {
	for i := int64(0); i < n; i += 1 {
		head := r.Head()
		tail := r.Tail()

		head = d(head)
		r.headPositions = append(r.headPositions, head)
		if distance(head, tail) <= 1 {
			continue
		}

		td := getDirection(head, tail)
		r.tailPositions = append(r.tailPositions, td(tail))
	}
}

func (r *rope) TailVisited() []position {
	m := make(map[string]position)
	for _, p := range r.tailPositions {
		m[p.String()] = p
	}

	u := make([]position, len(m))
	i := 0
	for _, p := range m {
		u[i] = p
	}

	return u
}

func parseLine(l string) (dir direction, n int64, err error) {
	parts := strings.Split(l, " ")
	if len(parts) < 2 {
		err = errors.New("invalid line format")
		return
	}

	switch parts[0] {
	case "R":
		dir = right
	case "L":
		dir = left
	case "U":
		dir = up
	case "D":
		dir = down
	default:
		err = errors.New("invalid direction")
		return
	}

	n, err = strconv.ParseInt(parts[1], 10, 64)
	return
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	r := rope{}
	lines := strings.Split(string(input), "\n")
	for i, l := range lines {
		dir, n, err := parseLine(l)
		if err != nil {
			log.Fatalf("problem parsing line: %v", err)
		}

		r.Move(dir, n)
		log.Printf("%6d\t%-6v\t%-12v\t%-12v", i+1, l, r.Head(), r.Tail())
	}

	uniqueTailPositions := r.TailVisited()
	fmt.Printf("Unique positions visited by tail: %v\n", len(uniqueTailPositions))
}
