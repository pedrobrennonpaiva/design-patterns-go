package strategies

type TaxCalculator interface {
	Calculate(amount float64) (*float64, error)
}
