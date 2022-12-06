package crane

import (
	"strconv"
	"strings"
)

type Command struct {
	fromStack string
	toStack   string
	quantity  int64
}

func (c *Command) FromStack() string {
	return c.fromStack
}

func (c *Command) ToStack() string {
	return c.toStack
}

func (c *Command) Quantity() int {
	return int(c.quantity)
}

func ParseCommands(value string) ([]*Command, error) {
	lines := strings.Split(value, "\n")
	commands := make([]*Command, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, " ")
		c := Command{}
		for j := 0; j < len(parts); j += 2 {
			value := parts[j+1]
			switch parts[j] {
			case "move":
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, err
				}

				c.quantity = intValue
			case "from":
				c.fromStack = value
			case "to":
				c.toStack = value
			}
		}
		commands[i] = &c
	}
	return commands, nil
}
