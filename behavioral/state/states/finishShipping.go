package states

import "errors"

type FinishShippingState struct{}

func (s *FinishShippingState) Prepare(order *Order) error {
	return errors.New("order was already shipped")
}

func (s *FinishShippingState) StartShipping(order *Order) error {
	return errors.New("order was already shipped")
}

func (s *FinishShippingState) FinishShipping(order *Order) error {
	return errors.New("order was already shipped")
}
