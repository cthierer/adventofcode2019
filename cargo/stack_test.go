package cargo_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/cargo"
)

func TestStackPush(t *testing.T) {
	s := cargo.Stack{ID: "1"}
	c := cargo.Crate("foo")
	err := s.Push(c)
	if err != nil {
		t.FailNow()
	}
}

func TestStackPop(t *testing.T) {
	s := cargo.Stack{ID: "1"}
	c1 := cargo.Crate("foo")
	c2 := cargo.Crate("bar")

	if err := s.Push(c1); err != nil {
		t.FailNow()
	}

	if err := s.Push(c2); err != nil {
		t.FailNow()
	}

	if v := s.Pop(1); v[0] != c2 {
		t.Logf("expected first pop to yield %s, but got %s", c2, v)
		t.Fail()
	}

	if v := s.Pop(1); v[0] != c1 {
		t.Logf("expected second pop to yield %s, but got %s", c1, v)
		t.Fail()
	}
}

func TestStackPeep(t *testing.T) {
	s := cargo.Stack{ID: "1"}
	c1 := cargo.Crate("foo")
	c2 := cargo.Crate("bar")

	if err := s.Push(c1); err != nil {
		t.FailNow()
	}

	if err := s.Push(c2); err != nil {
		t.FailNow()
	}

	if v := s.Peek(); v != c2 {
		t.Logf("expected first peek to yield %s, but got %s", c2, v)
		t.Fail()
	}

	if v := s.Peek(); v != c2 {
		t.Logf("expected second peek to yield %s, but got %s", c2, v)
		t.Fail()
	}
}

func TestStackValues(t *testing.T) {
	s := cargo.Stack{ID: "1"}
	c1 := cargo.Crate("foo")
	c2 := cargo.Crate("bar")

	if err := s.Push(c1); err != nil {
		t.FailNow()
	}

	if err := s.Push(c2); err != nil {
		t.FailNow()
	}

	v := s.Values()
	if v[0] != c1 || v[1] != c2 {
		t.Fail()
	}
}
