package machine

import (
	"fmt"
	"strings"
)

type item struct {
	Name      string `json:"name"`
	Cost      int    `json:"cost"`
	Inventory int    `json:"inventory"`
}

func (i *item) isAvailable() bool {
	return i.Inventory > 0
}

func (i *item) dispense() {
	fmt.Printf("Dispensing.....%v\n", i.Name)
	i.Inventory--
}

func (i *item) fundsAvailable(avail int) bool {
	return avail > i.Cost
}

func (i *item) isSelected(input string) bool {
	return strings.EqualFold(i.Name, input)
}
