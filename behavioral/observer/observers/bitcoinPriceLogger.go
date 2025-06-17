package observers

type BitcoinPriceLogger struct{}

func NewBitcoinPriceLogger() *BitcoinPriceLogger {
	return &BitcoinPriceLogger{}
}

func (b *BitcoinPriceLogger) Update(price float64) {
	// Log the price to a file or console
	// For simplicity, we'll just print it to the console
	println("Bitcoin price updated:", price)
}
