package inventory

func (inv *Inventory) Slice(start, end int) *Inventory {
	selected := inv.lineItems[start:end]
	return &Inventory{lineItems: selected}
}
