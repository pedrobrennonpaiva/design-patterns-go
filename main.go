package main

import (
	"design-patterns-go/behavioral/observer"
	"design-patterns-go/behavioral/state"
	"design-patterns-go/behavioral/strategy"
	templateMethod "design-patterns-go/behavioral/templateMethod"
	abstractFactory "design-patterns-go/creational/abstractFactory"
	"design-patterns-go/creational/builder"
	factoryMethod "design-patterns-go/creational/factoryMethod"
	"design-patterns-go/creational/prototype"
	"design-patterns-go/creational/singleton"
	"design-patterns-go/structural/decorator"
	"design-patterns-go/structural/proxy"
	"fmt"
)

func main() {
	fmt.Println("Learning Design Patterns in Go")

	fmt.Println("Creational Patterns")

	fmt.Println("\nFactory Method")
	factoryMethod.Run()

	fmt.Println("\nAbstract Factory")
	abstractFactory.Run()

	fmt.Println("\nBuilder")
	builder.Run()

	fmt.Println("\nSingleton")
	singleton.Run()

	fmt.Println("\nPrototype")
	prototype.Run()

	fmt.Println("\nBehavioral Patterns")

	fmt.Println("\nStrategy")
	strategy.Run()

	fmt.Println("\nObserver")
	observer.Run()

	fmt.Println("\nTemplate Method")
	templateMethod.Run()

	fmt.Println("\nState")
	state.Run()

	fmt.Println("\nStructural Patterns")

	fmt.Println("\nDecorator")
	decorator.Run()

	fmt.Println("\nProxy")
	proxy.Run()
}
