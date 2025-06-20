package states

type OrderState interface {
	Prepare(order *Order) error
	StartShipping(order *Order) error
	FinishShipping(order *Order) error
}

type Order struct {
	state OrderState
}

func NewOrder() *Order {
	return &Order{
		state: &ProcessedState{}, // Default state
	}
}

func (order *Order) SetState(state OrderState) {
	order.state = state
}

func (order *Order) GetStateName() string {
	switch order.state.(type) {
	case *ProcessedState:
		return "Processed"
	case *PreparingState:
		return "Preparing"
	case *StartShippingState:
		return "Start Shipping"
	case *FinishShippingState:
		return "Finish Shipping"
	default:
		return "Unknown State"
	}
}

func (order *Order) Prepare() error {
	return order.state.Prepare(order)
}

func (order *Order) StartShipping() error {
	return order.state.StartShipping(order)
}

func (order *Order) FinishShipping() error {
	return order.state.FinishShipping(order)
}
