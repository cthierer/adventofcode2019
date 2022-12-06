package cargo_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/cargo"
)

const sampleInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `

func TestCollectionAdd(t *testing.T) {
	c := cargo.NewCollection()
	s := cargo.Stack{ID: "foo"}
	c.Add(&s)
}

func TestCollectionGet(t *testing.T) {
	c := cargo.NewCollection()
	s := cargo.Stack{ID: "foo"}
	c.Add(&s)

	if c.Get("foo") != &s {
		t.Fail()
	}
}

type transferCommand struct {
	fromStack string
	toStack   string
	quantity  int
}

func (c *transferCommand) FromStack() string {
	return c.fromStack
}

func (c *transferCommand) ToStack() string {
	return c.toStack
}

func (c *transferCommand) Quantity() int {
	return c.quantity
}

func TestCollectionTransfer(t *testing.T) {
	c := cargo.NewCollection()
	s1 := cargo.Stack{ID: "1"}
	s2 := cargo.Stack{ID: "2"}

	c.Add(&s1)
	c.Add(&s2)

	s1.Push(cargo.Crate("A"))
	s2.Push(cargo.Crate("B"))

	if err := c.Transfer(&transferCommand{"1", "2", 1}); err != nil {
		t.FailNow()
	}

	if v := s1.Values(); len(v) != 0 {
		t.Logf("expected to transfer all values from stack 1, but stack still has %v", len(v))
		t.Fail()
	}

	if v := s2.Values(); len(v) != 2 {
		t.Logf("expected to transfer all values to stack 2, but stack only has %v", len(v))
		t.Fail()
	}
}

func TestCollectionValues(t *testing.T) {
	c := cargo.NewCollection()
	s1 := cargo.Stack{ID: "1"}
	s2 := cargo.Stack{ID: "2"}

	c.Add(&s1)
	c.Add(&s2)

	if v := c.Values(); len(v) != 2 {
		t.Fail()
	}
}

func TestParseCollectioon(t *testing.T) {
	c, err := cargo.ParseCollection(sampleInput)
	if err != nil {
		t.FailNow()
	}

	v1 := c.Get("1").Values()
	if len(v1) != 2 || v1[0] != cargo.Crate("Z") || v1[1] != cargo.Crate("N") {
		t.Logf("expected stack 1 to have values, but got %v", v1)
		t.FailNow()
	}

	v2 := c.Get("2").Values()
	if len(v2) != 3 || v2[0] != cargo.Crate("M") || v2[1] != cargo.Crate("C") || v2[2] != cargo.Crate("D") {
		t.Logf("expected stack 2 to have values, but got %v", v2)
		t.FailNow()
	}

	v3 := c.Get("3").Values()
	if len(v3) != 1 || v3[0] != cargo.Crate("P") {
		t.Logf("expected stack 3 to have values, but got %v", v3)
		t.FailNow()
	}
}
