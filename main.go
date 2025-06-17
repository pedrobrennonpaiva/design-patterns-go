package main

import (
	"design-patterns-go/behavioral/observer"
	"design-patterns-go/behavioral/strategy"
	abstractFactory "design-patterns-go/creational/abstractFactory"
	"design-patterns-go/creational/builder"
	factoryMethod "design-patterns-go/creational/factoryMethod"
	"design-patterns-go/creational/prototype"
	"design-patterns-go/creational/singleton"
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
}
