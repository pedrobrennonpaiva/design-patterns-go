package state

import (
	"design-patterns-go/behavioral/state/states"
	"fmt"
)

func Run() {

	order := states.NewOrder()

	if err := order.Prepare(); err != nil {
		fmt.Println(err)
		return
	}

	if err := order.StartShipping(); err != nil {
		fmt.Println(err)
		return
	}

	if err := order.FinishShipping(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Final state order:", order.GetStateName())
}
