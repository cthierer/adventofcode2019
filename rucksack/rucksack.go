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
	for t := range c.contents {
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

func NewRucksack() *Rucksack {
	compartment1 := NewItemCollection()
	compartment2 := NewItemCollection()
	return &Rucksack{Compartment1: compartment1, Compartment2: compartment2}
}

func addContents(contents string, compartment *ItemCollection) {
	for _, r := range contents {
		compartment.Add(ParseItemType(r))
	}
}

func ParseRucksack(value string) *Rucksack {
	r := NewRucksack()
	halfway := len(value) / 2
	compartment1 := value[0:halfway]
	compartment2 := value[halfway:]

	addContents(compartment1, r.Compartment1)
	addContents(compartment2, r.Compartment2)

	return r
}

func (r *Rucksack) OverlappingItemTypes() *ItemCollection {
	return r.Compartment1.Intersect(r.Compartment2)
}

func (r *Rucksack) AllItems() *ItemCollection {
	allItems := NewItemCollection()
	allItems.Join(r.Compartment1)
	allItems.Join(r.Compartment2)
	return allItems
}

type Group struct {
	rucksacks []*Rucksack
}

func (g *Group) Add(r *Rucksack) {
	g.rucksacks = append(g.rucksacks, r)
}

func (g *Group) OverlappingItemTypes() *ItemCollection {
	itemTypes := g.rucksacks[0].AllItems()
	for i := 1; i < len(g.rucksacks); i += 1 {
		itemTypes = itemTypes.Intersect(g.rucksacks[i].AllItems())
	}
	return itemTypes
}
