package strategies

type ICMSCalculator struct{}

func (c ICMSCalculator) Calculate(amount float64) (*float64, error) {
	calc := (amount * 4) / 100
	return &calc, nil
}
