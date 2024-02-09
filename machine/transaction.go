package machine

import "fmt"

type transaction struct {
	total  int
	pCount int
	nCount int
	dCount int
	qCount int
}

func (t *transaction) deposit() {
	for {
		var amount string

		if t.total > 0 {
			fmt.Printf("%v has been deposited\n", t.total)
		}
		fmt.Print(`
What is being deposited?
	Penny   = p
	Nickle  = n
	Dime    = d
	Quarter = q
	Finish  = f 

Enter: `)

		fmt.Scanln(&amount)

		switch amount {
		case "p":
			t.total += 1
			t.pCount++
		case "n":
			t.total += 5
			t.nCount++
		case "d":
			t.total += 10
			t.dCount++
		case "q":
			t.total += 25
			t.qCount++
		case "f":
			// exit the main loop
			return
		default:
			fmt.Println("invalid coin, select again")
		}

		fmt.Println("----------------------------------")
	}
}

func (t *transaction) cancel() {
	if t.total == 0 {
		return
	}
	output := "\nCancelling...and returning..."
	output += t.getCoinOutput()
	fmt.Println(output)
}

func (t *transaction) getCoinOutput() string {
	output := ""
	if t.pCount > 0 {
		output += fmt.Sprintf("\n%v pennies", t.pCount)
	}
	if t.nCount > 0 {
		output += fmt.Sprintf("\n%v nickles", t.nCount)
	}
	if t.dCount > 0 {
		output += fmt.Sprintf("\n%v dimes", t.dCount)
	}
	if t.qCount > 0 {
		output += fmt.Sprintf("\n%v quarters", t.qCount)
	}
	return output
}

func (t *transaction) recordTransaction(amount int) {
	t.total -= amount
	t.pCount = 0
	t.nCount = 0
	t.dCount = 0
	t.qCount = 0
}

func (t *transaction) returnChange() {
	if t.total == 0 {
		return
	}

	t.qCount, t.total = handleCoin(t.total, 25)
	t.dCount, t.total = handleCoin(t.total, 10)
	t.nCount, t.total = handleCoin(t.total, 5)
	t.pCount, t.total = handleCoin(t.total, 1)

	output := "\nReturning Change..."
	output += t.getCoinOutput()
	fmt.Println(output)
}

func handleCoin(amount, value int) (int, int) {
	count := 0
	for {
		if amount < value {
			return count, amount
		}
		count++
		amount -= value
	}
}
