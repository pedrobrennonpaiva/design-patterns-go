package observers

type BitcoinPriceObserver interface {
	// warn: don't use float64 to represent money in production code.
	Update(price float64)
}
