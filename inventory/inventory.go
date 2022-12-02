package inventory

type Tuple struct {
	index int
	value int64
}

func (t Tuple) Index() int {
	return t.index
}

func (t Tuple) Value() int64 {
	return t.value
}

type Inventory struct {
	lineItems []Tuple
}

func (inv *Inventory) LineItems() []Tuple {
	return inv.lineItems
}
