package inventory

import "sort"

func (inv *Inventory) Sort() *Inventory {
	copy := make([]Tuple, len(inv.lineItems))
	for i, t := range inv.lineItems {
		copy[i] = t
	}

	sort.Slice(copy, func(i, j int) bool {
		return copy[i].value > copy[j].value
	})

	return &Inventory{lineItems: copy}
}
