package main

import "github.com/jcollins-axway/vending-machine/machine"

func main() {
	// vm := machine.InitVendingMachine("candy.json")
	vm := machine.InitVendingMachine("drinks.json")
	vm.Start()
}
