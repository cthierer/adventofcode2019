package inventory

func (inv *Inventory) SumByIndex() *Inventory {
	totals := make(map[int]int64)
	for _, t := range inv.lineItems {
		totals[t.index] += t.value
	}

	results := make([]Tuple, len(totals))
	i := 0
	for idx, sum := range totals {
		results[i] = Tuple{index: idx, value: sum}
		i += 1
	}

	return &Inventory{lineItems: results}
}
