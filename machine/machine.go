package machine

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type VendingMachine interface {
	Start()
}

type vend struct {
	items []*item
}

func InitVendingMachine(itemsFile string) VendingMachine {
	// open file
	f, err := os.Open(itemsFile)
	if err != nil {
		panic("no items loaded")
	}

	// read file
	data, err := io.ReadAll(f)
	if err != nil {
		panic("could not data from file")
	}

	// convert json data to array
	items := []*item{}
	err = json.Unmarshal(data, &items)
	if err != nil {
		panic("could not read data as json")
	}

	return &vend{
		items: items,
	}
}

func (v *vend) Start() {
	for {
		var action string

		fmt.Print(`
What would you like to do?
	Check Inventory  = i
	Buy Item         = b
	Quit             = q

Enter: `)

		fmt.Scanln(&action)

		switch action {
		case "i":
			v.inventoryCheck()
		case "b":
			v.purchaseItem()
		case "q":
			// exit the main loop
			return
		default:
			fmt.Println("invalid option, select again")
		}

		fmt.Println("----------------------------------")
	}
}

func (v *vend) inventoryCheck() {
	fmt.Println("\nInventory of Items")
	v.displayOpts()
}

func (v *vend) displayOpts() {
	output := ""
	for _, i := range v.items {
		if i.isAvailable() {
			output += fmt.Sprintf("\n  %v - %v cents - %v available", i.Name, i.Cost, i.Inventory)
		}
	}
	output += "\n"
	fmt.Println(output)
}

func (v *vend) purchaseItem() {
	// start of transaction
	t := transaction{}

	t.deposit()

	fmt.Printf("\nA total of %v has been deposited\n", t.total)

	fmt.Println("\nType an items name to select it")
	v.displayOpts()
	fmt.Println("type in q to cancel the transaction")
	fmt.Print("Enter: ")

	var selection string
	fmt.Scanln(&selection)
	if selection == "q" {
		// revert transaction
		t.cancel()
		return
	}

	// find the item selected and dispense
	var selectedItem *item
	for _, i := range v.items {
		if i.isSelected(selection) {
			selectedItem = i
			break
		}
	}

	if !selectedItem.isAvailable() {
		fmt.Println("Selected item is not available")
		t.cancel()
		return
	}

	if !selectedItem.fundsAvailable(t.total) {
		fmt.Println("Not enough deposited")
		t.cancel()
		return
	}

	selectedItem.dispense()
	t.recordTransaction(selectedItem.Cost)

	// return change
	t.returnChange()
}
