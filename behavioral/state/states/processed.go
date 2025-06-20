package states

import (
	"errors"
	"fmt"
)

type ProcessedState struct{}

func (s *ProcessedState) Prepare(order *Order) error {
	fmt.Println("Order will be prepared...")
	order.SetState(&PreparingState{})
	return nil
}

func (s *ProcessedState) StartShipping(order *Order) error {
	return errors.New("order is not prepared yet")
}

func (s *ProcessedState) FinishShipping(order *Order) error {
	return errors.New("order is not prepared yet")
}
