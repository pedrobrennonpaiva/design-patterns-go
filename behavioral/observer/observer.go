package observer

import (
	"design-patterns-go/behavioral/observer/observers"
)

func Run() {

	bitcoin := NewBitcoin()

	binanceApi := NewBinanceApi()
	newPrice := binanceApi.GetLastPrice()

	bitcoin.AddObserver(observers.NewBitcoinPriceLogger())
	bitcoin.AddObserver(observers.NewInvestorNotifier())
	bitcoin.AddObserver(observers.NewNewsPlatform())

	bitcoin.SetPrice(newPrice)
}

type Bitcoin struct {
	price     float64
	observers []observers.BitcoinPriceObserver
}

func NewBitcoin() *Bitcoin {
	return &Bitcoin{
		price: 0.0,
	}
}

func (b *Bitcoin) SetPrice(newPrice float64) {

	if newPrice != b.price {
		b.price = newPrice
		b.notifyObservers()
	}
}

func (b *Bitcoin) GetPrice() float64 {
	return b.price
}

func (b *Bitcoin) AddObserver(observer observers.BitcoinPriceObserver) {
	b.observers = append(b.observers, observer)
}

func (b *Bitcoin) notifyObservers() {
	for _, observer := range b.observers {
		observer.Update(b.price)
	}
}

type BinanceApi struct{}

func NewBinanceApi() *BinanceApi {
	return &BinanceApi{}
}

func (b *BinanceApi) GetLastPrice() float64 {
	// Simulating an API call to get the last price of Bitcoin.
	// In a real-world scenario, this would involve making an HTTP request to a cryptocurrency exchange API.
	return 50000.0 // Example price
}
