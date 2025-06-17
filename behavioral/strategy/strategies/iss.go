package strategies

type ISSCalculator struct{}

func (c ISSCalculator) Calculate(amount float64) (*float64, error) {
	calc := (amount * 11) / 100
	return &calc, nil
}
