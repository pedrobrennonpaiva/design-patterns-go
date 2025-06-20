package states

import (
	"errors"
	"fmt"
)

type StartShippingState struct{}

func (s *StartShippingState) Prepare(order *Order) error {
	return errors.New("order was already prepared")
}

func (s *StartShippingState) StartShipping(order *Order) error {
	return errors.New("order is already being shipped")
}

func (s *StartShippingState) FinishShipping(order *Order) error {
	fmt.Println("Order will be finish ship...")
	order.SetState(&FinishShippingState{})
	return nil
}
