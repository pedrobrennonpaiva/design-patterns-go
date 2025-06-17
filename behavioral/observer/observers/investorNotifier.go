package observers

type InvestorNotifier struct{}

func NewInvestorNotifier() *InvestorNotifier {
	return &InvestorNotifier{}
}

func (i *InvestorNotifier) Update(price float64) {
	// Notify investors about the price change
	// For simplicity, we'll just print it to the console
	println("Investor notified: Bitcoin price updated to", price)
}
