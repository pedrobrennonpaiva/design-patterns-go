package observers

type NewsPlatform struct{}

func NewNewsPlatform() *NewsPlatform {
	return &NewsPlatform{}
}

func (n *NewsPlatform) Update(price float64) {
	// Notify users about the price change
	// For simplicity, we'll just print it to the console
	println("News Platform: Bitcoin price updated to", price)
}
