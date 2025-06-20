package facades

type OrderFacade struct {
	PaymentProcessor *PaymentProcessor
	Notifier         *Notifier
	InventoryManager *InventoryManager
	ShippingService  *ShippingService
}

func NewOrderFacade(paymentProcessor PaymentProcessor, notifier Notifier, inventoryManager InventoryManager, shippingService ShippingService) *OrderFacade {
	return &OrderFacade{
		PaymentProcessor: &paymentProcessor,
		Notifier:         &notifier,
		InventoryManager: &inventoryManager,
		ShippingService:  &shippingService,
	}
}

func (f *OrderFacade) ProcessOrder(order Order) error {

	if err := f.PaymentProcessor.ProcessPayment(order.Amount); err != nil {
		return err
	}

	if err := f.Notifier.Notify(order.Email); err != nil {
		return err
	}

	if err := f.InventoryManager.UpdateInventory(order.ProductId, order.Quantity); err != nil {
		return err
	}

	if err := f.ShippingService.InitiateShipping(&order); err != nil {
		return err
	}

	return nil
}
