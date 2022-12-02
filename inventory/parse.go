package inventory

import (
	"strconv"
	"strings"
)

func ParseInventory(input string) (*Inventory, error) {
	var items []Tuple
	i := 0

	for _, v := range strings.Split(strings.TrimSpace(input), "\n") {
		if v == "" {
			i += 1
			continue
		}

		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}

		items = append(items, Tuple{index: i, value: value})
	}

	return &Inventory{lineItems: items}, nil
}
