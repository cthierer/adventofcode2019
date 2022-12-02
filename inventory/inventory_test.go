package inventory_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/inventory"
)

const sampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestParseInventory(t *testing.T) {
	_, err := inventory.ParseInventory(sampleInput)
	if err != nil {
		t.Logf("failed to parse inventory: %v", err)
		t.FailNow()
	}
}

func TestSumByIndex(t *testing.T) {
	i, err := inventory.ParseInventory(sampleInput)
	if err != nil {
		t.FailNow()
	}

	sums := i.SumByIndex()
	actual := make([]int64, len(sums))
	for _, tuple := range sums {
		actual[tuple.Index()] = tuple.Value()
	}

	if actual[0] != 6000 {
		t.Logf("expected first elf to carry 6000 calories")
		t.Fail()
	}

	if actual[1] != 4000 {
		t.Logf("expected second elf to carry 4000 calories")
		t.Fail()
	}

	if actual[2] != 11000 {
		t.Logf("expected third elf to carry 11000 calories")
		t.Fail()
	}

	if actual[3] != 24000 {
		t.Logf("expected fourth elf to carry 24000 calories")
		t.Fail()
	}

	if actual[4] != 10000 {
		t.Logf("expected fifth elf to carry 10000 calories")
		t.Fail()
	}
}
