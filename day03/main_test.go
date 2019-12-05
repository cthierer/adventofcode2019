package main

import "testing"

func TestClosestIntersection(t *testing.T) {
	table := []struct {
		path1    string
		path2    string
		distance float64
	}{
		{
			path1:    "R8,U5,L5,D3",
			path2:    "U7,R6,D4,L4",
			distance: 6,
		},
		{
			path1:    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			path2:    "U62,R66,U55,R34,D71,R55,D58,R83",
			distance: 159,
		},
		{
			path1:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			path2:    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			distance: 135,
		},
	}

	for _, v := range table {
		_, actual := closestIntersection(v.path1, v.path2)
		if v.distance != actual {
			t.Fail()
		}
	}
}

func TestFewestSteps(t *testing.T) {
	table := []struct {
		path1 string
		path2 string
		steps int
	}{
		{
			path1: "R8,U5,L5,D3",
			path2: "U7,R6,D4,L4",
			steps: 30,
		},
		{
			path1: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			path2: "U62,R66,U55,R34,D71,R55,D58,R83",
			steps: 610,
		},
		{
			path1: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			path2: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			steps: 410,
		},
	}

	for _, v := range table {
		_, actual := fewestSteps(v.path1, v.path2)
		if v.steps != actual {
			t.Fail()
		}
	}
}
