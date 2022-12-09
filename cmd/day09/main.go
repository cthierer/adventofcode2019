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
	knots [][]position
}

func (r *rope) Knot(i int) position {
	positions := r.knots[i]
	if len(positions) == 0 {
		positions = append(positions, position{})
		r.knots[i] = positions
	}
	return positions[len(positions)-1]
}

func (r *rope) Head() position {
	return r.Knot(0)
}

func (r *rope) Tail() position {
	return r.Knot(len(r.knots) - 1)
}

func (r *rope) MoveKnot(i int, d direction) position {
	next := d(r.Knot(i))
	r.knots[i] = append(r.knots[i], next)
	return next
}

func (r *rope) MoveHead(d direction) position {
	return r.MoveKnot(0, d)
}

func (r *rope) Move(d direction, n int64) {
	for i := int64(0); i < n; i += 1 {
		r.MoveHead(d)

		for j := 1; j < len(r.knots); j += 1 {
			last := r.Knot(j - 1)
			curr := r.Knot(j)
			if distance(last, curr) <= 1 {
				continue
			}

			move := getDirection(last, curr)
			r.MoveKnot(j, move)
		}
	}
}

func (r *rope) TailVisited() []position {
	m := make(map[string]position)
	for _, p := range r.knots[len(r.knots)-1] {
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

	lines := strings.Split(string(input), "\n")

	log.Printf("--- PART 1 ---\n")
	r1 := rope{knots: make([][]position, 2)}
	for i, l := range lines {
		dir, n, err := parseLine(l)
		if err != nil {
			log.Fatalf("problem parsing line: %v", err)
		}

		r1.Move(dir, n)
		log.Printf("%6d\t%-6v\t%-12v\t%-12v", i+1, l, r1.Head(), r1.Tail())
	}

	log.Printf("--- PART 2 ---\n")
	r2 := rope{knots: make([][]position, 10)}
	for i, l := range lines {
		dir, n, err := parseLine(l)
		if err != nil {
			log.Fatalf("problem parsing line: %v", err)
		}

		r2.Move(dir, n)
		log.Printf("%6d\t%-6v\t%-12v\t%-12v", i+1, l, r2.Head(), r2.Tail())
	}

	fmt.Printf("Unique positions visited by tail with 2 knots (part 1): %v\n", len(r1.TailVisited()))
	fmt.Printf("Unique positions visited by tail with 10 knots (part 2): %v\n", len(r2.TailVisited()))
}
