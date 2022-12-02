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

	sums := i.SumByIndex().LineItems()
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

func TestSort(t *testing.T) {
	i, err := inventory.ParseInventory(sampleInput)
	if err != nil {
		t.FailNow()
	}

	sums := i.SumByIndex().Sort().LineItems()
	if sums[0].Index() != 3 || sums[0].Value() != 24000 {
		t.Logf("expected fourth elf to carry the most calories")
		t.Fail()
	}

	if sums[1].Index() != 2 || sums[1].Value() != 11000 {
		t.Logf("expected third elf to carry the second most calories")
		t.Fail()
	}

	if sums[2].Index() != 4 || sums[2].Value() != 10000 {
		t.Logf("expected fifth elf too carry the third most calories")
		t.Fail()
	}
}

func TestSlice(t *testing.T) {
	i, err := inventory.ParseInventory(sampleInput)
	if err != nil {
		t.FailNow()
	}

	sums := i.SumByIndex()
	topThree := sums.Sort().Slice(0, 3).LineItems()

	if len(topThree) != 3 {
		t.Logf("expected to only have 3 items in slice")
		t.Fail()
	}

	if topThree[0].Value() != 24000 || topThree[1].Value() != 11000 || topThree[2].Value() != 10000 {
		t.Logf("expected top 3 to be the highest values")
		t.Fail()
	}

}

func TestTotal(t *testing.T) {
	i, err := inventory.ParseInventory(sampleInput)
	if err != nil {
		t.FailNow()
	}

	total := i.SumByIndex().Sort().Slice(0, 3).Total()

	if total != 45000 {
		t.Logf("expected top 3 to total 45000")
		t.Fail()
	}
}
