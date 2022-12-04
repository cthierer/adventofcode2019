package rucksack_test

import (
	"testing"

	"github.com/cthierer/advent-of-code/rucksack"
)

func TestRucksackIntersection(t *testing.T) {
	scenarios := []struct {
		compartment1 string
		compartment2 string
		expected     []rune
	}{
		{
			compartment1: "vJrwpWtwJgWr",
			compartment2: "hcsFMMfFFhFp",
			expected:     []rune{'p'},
		},
		{
			compartment1: "jqHRNqRjqzjGDLGL",
			compartment2: "rsFMfFZSrLrFZsSL",
			expected:     []rune{'L'},
		},
		{
			compartment1: "PmmdzqPrV",
			compartment2: "vPwwTWBwg",
			expected:     []rune{'P'},
		},
		{
			compartment1: "wMqvLMZHhHMvwLH",
			compartment2: "jbvcjnnSBnvTQFn",
			expected:     []rune{'v'},
		},
		{
			compartment1: "ttgJtRGJ",
			compartment2: "QctTZtZT",
			expected:     []rune{'t'},
		},
		{
			compartment1: "CrZsJsPPZsGz",
			compartment2: "wwsLwLmpwMDw",
			expected:     []rune{'s'},
		},
	}

	for _, s := range scenarios {
		r := rucksack.NewRucksack()

		for _, itemType := range s.compartment1 {
			r.Compartment1.Add(rucksack.ParseItemType(itemType))
		}

		for _, itemType := range s.compartment2 {
			r.Compartment2.Add(rucksack.ParseItemType(itemType))
		}

		actual := r.OverlappingItemTypes().UniqueValues()

		if len(actual) != len(s.expected) {
			t.Logf("expected %v overlapping types, but got %v", len(s.expected), len(actual))
			t.FailNow()
		}

		for _, expectedItemType := range s.expected {
			found := false
			for _, v := range actual {
				found = v.Value() == expectedItemType
				if found {
					break
				}
			}
			if !found {
				t.Logf("expected to find %v, but did not", expectedItemType)
				t.Fail()
			}
		}
	}
}

func TestItemCollectionSum(t *testing.T) {
	scenarios := []struct {
		values   []rune
		expected int
	}{
		{
			values:   []rune{'p'},
			expected: 16,
		},
		{
			values:   []rune{'L'},
			expected: 38,
		},
		{
			values:   []rune{'P'},
			expected: 42,
		},
		{
			values:   []rune{'v'},
			expected: 22,
		},
		{
			values:   []rune{'t'},
			expected: 20,
		},
		{
			values:   []rune{'s'},
			expected: 19,
		},
	}

	for _, s := range scenarios {
		collection := rucksack.NewItemCollection()
		for _, r := range s.values {
			collection.Add(rucksack.ParseItemType(r))
		}

		actual := collection.Sum()
		if actual != s.expected {
			t.Logf("expected sum to be %v, but got %v", s.expected, actual)
			t.Fail()
		}
	}
}

func TestItemCollectionJoin(t *testing.T) {
	collections := []struct {
		values []rune
	}{
		{
			values: []rune{'p'},
		},
		{
			values: []rune{'L'},
		},
		{
			values: []rune{'P'},
		},
		{
			values: []rune{'v'},
		},
		{
			values: []rune{'t'},
		},
		{
			values: []rune{'s'},
		},
	}

	actual := rucksack.NewItemCollection()
	for _, s := range collections {
		collection := rucksack.NewItemCollection()
		for _, r := range s.values {
			collection.Add(rucksack.ParseItemType(r))
		}

		actual.Join(collection)
	}

	if actual.Sum() != 157 {
		t.Fail()
	}
}

func TestGroupOverlappingItemTypes(t *testing.T) {
	g1 := rucksack.Group{}
	g1.Add(rucksack.ParseRucksack("vJrwpWtwJgWrhcsFMMfFFhFp"))
	g1.Add(rucksack.ParseRucksack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
	g1.Add(rucksack.ParseRucksack("PmmdzqPrVvPwwTWBwg"))

	g1UniqueTypes := g1.OverlappingItemTypes().UniqueValues()
	if len(g1UniqueTypes) != 1 {
		t.Logf("expected only one unique type, but got %v", len(g1UniqueTypes))
		t.Fail()
	}

	if g1UniqueTypes[0].String() != "r" {
		t.Logf("expected unique type to be r, but got %v", g1UniqueTypes[0].String())
		t.Fail()
	}

	g2 := rucksack.Group{}
	g2.Add(rucksack.ParseRucksack("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"))
	g2.Add(rucksack.ParseRucksack("ttgJtRGJQctTZtZT"))
	g2.Add(rucksack.ParseRucksack("CrZsJsPPZsGzwwsLwLmpwMDw"))

	g2UniqueTypes := g2.OverlappingItemTypes().UniqueValues()
	if len(g2UniqueTypes) != 1 {
		t.Logf("expected only one unique type, but got %v", len(g2UniqueTypes))
		t.Fail()
	}

	if g2UniqueTypes[0].String() != "Z" {
		t.Logf("expected unique types to be Z, but got %v", g2UniqueTypes[0].String())
		t.Fail()
	}
}
