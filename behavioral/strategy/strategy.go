package strategy

import (
	"design-patterns-go/behavioral/strategy/strategies"
	"fmt"
)

func Run() {

	amount := 100.0
	taxType := "ICMS"

	// tip: use factory pattern to create tax calculators for different tax types.
	switch taxType {
	case "ICMS":
		setTaxType(strategies.ICMSCalculator{})
	case "ISS":
		setTaxType(strategies.ISSCalculator{})
	case "IPI":
		setTaxType(strategies.IPICalculator{})
	default:
		panic("invalid tax type")
	}

	tax, err := calculate(amount)
	if err != nil {
		panic(err)
	}

	println("Tax type:", taxType)
	fmt.Printf("Amount: %.2f\n", amount)
	fmt.Printf("Tax amount: %.2f\n", *tax)
	fmt.Printf("Total amount with tax: %.2f\n", amount+*tax)
	fmt.Printf("Total amount without tax: %.2f\n", amount)
}

var tax strategies.TaxCalculator

// warn: don't use float64 to represent money in production code.
func calculate(amount float64) (*float64, error) {

	calc, err := tax.Calculate(amount)
	if err != nil {
		return nil, err
	}

	return calc, nil
}

func setTaxType(taxType strategies.TaxCalculator) {
	tax = taxType
}
