package facade

import (
	"design-patterns-go/structural/facade/facades"
	"fmt"
)

func Run() {

	order := facades.Order{
		Amount:    100.0,
		Email:     "test@test.com",
		ProductId: 1,
		Quantity:  2,
		UserId:    123,
	}
	fmt.Println("Processing order:", order)

	//In a real-world scenario, you might want to use interfaces to allow for easier testing
	//We are using concrete types directly for simplicity
	paymentProcessor := facades.PaymentProcessor{}
	notifier := facades.Notifier{}
	inventoryManager := facades.InventoryManager{}
	shippingService := facades.ShippingService{}

	orderFacade := facades.NewOrderFacade(paymentProcessor, notifier, inventoryManager, shippingService)
	if err := orderFacade.ProcessOrder(order); err != nil {
		fmt.Println("Error processing order:", err)
		return
	}

	fmt.Println("Order processed successfully")
}
