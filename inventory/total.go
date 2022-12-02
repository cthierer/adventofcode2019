package inventory

func (inv *Inventory) Total() int64 {
	var total int64
	for _, t := range inv.lineItems {
		total += t.value
	}
	return total
}
