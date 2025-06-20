package facades

type Order struct {
	// warn: don't use float64 to represent money in production code.
	Amount float64

	Email     string
	ProductId int
	Quantity  int
	UserId    int
}

type PaymentProcessor struct{}

func (p *PaymentProcessor) ProcessPayment(amount float64) error {
	// Simulate payment processing logic
	return nil
}

type Notifier struct{}

func (n *Notifier) Notify(email string) error {
	// Simulate notification logic
	return nil
}

type InventoryManager struct{}

func (i *InventoryManager) UpdateInventory(productId, quantity int) error {
	// Simulate inventory update logic
	return nil
}

type ShippingService struct{}

func (s *ShippingService) InitiateShipping(order *Order) error {
	// Simulate shipping logic
	return nil
}
