package states

import (
	"errors"
	"fmt"
)

type PreparingState struct{}

func (s *PreparingState) Prepare(order *Order) error {
	return errors.New("order is already being prepared")
}

func (s *PreparingState) StartShipping(order *Order) error {
	fmt.Println("Order will be start ship...")
	order.SetState(&StartShippingState{})
	return nil
}

func (s *PreparingState) FinishShipping(order *Order) error {
	return errors.New("order cannot be finished yet, it is still being prepared")
}
