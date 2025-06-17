package strategies

type IPICalculator struct{}

func (c IPICalculator) Calculate(amount float64) (*float64, error) {
	calc := (amount * 15) / 100
	return &calc, nil
}
