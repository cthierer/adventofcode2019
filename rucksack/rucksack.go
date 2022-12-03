package rucksack

import (
	"strings"
)

const priorityLookup = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type ItemType interface {
	Priority() int
	Value() rune
	String() string
}

type itemType rune

func ParseItemType(value rune) ItemType {
	return itemType(value)
}

func (t itemType) Priority() int {
	priority := strings.Index(priorityLookup, string(t)) + 1
	return priority
}

func (t itemType) Value() rune {
	return rune(t)
}

func (t itemType) String() string {
	return string(t)
}

type ItemCollection struct {
	contents map[ItemType]int
}

func NewItemCollection() *ItemCollection {
	contents := make(map[ItemType]int)
	return &ItemCollection{contents: contents}
}

func (c *ItemCollection) Add(i ItemType) {
	c.contents[i] += 1
}

func (c *ItemCollection) Sum() int {
	total := 0
	for t, count := range c.contents {
		total += t.Priority() * count
	}
	return total
}

func (c *ItemCollection) Join(other *ItemCollection) {
	for itemType, count := range other.contents {
		c.contents[itemType] += count
	}
}

func (c *ItemCollection) UniqueValues() []ItemType {
	values := make([]ItemType, len(c.contents))
	i := 0
	for t := range c.contents {
		values[i] = t
		i += 1
	}
	return values
}

func (c *ItemCollection) Intersect(other *ItemCollection) *ItemCollection {
	intersection := make(map[ItemType]int)
	for t, _ := range c.contents {
		if _, ok := other.contents[t]; ok {
			intersection[t] = 1
		}
	}
	return &ItemCollection{contents: intersection}
}

type Rucksack struct {
	Compartment1 *ItemCollection
	Compartment2 *ItemCollection
}

func (r *Rucksack) OverlappingItemTypes() *ItemCollection {
	return r.Compartment1.Intersect(r.Compartment2)
}

func NewRucksack() *Rucksack {
	compartment1 := NewItemCollection()
	compartment2 := NewItemCollection()
	return &Rucksack{Compartment1: compartment1, Compartment2: compartment2}
}
