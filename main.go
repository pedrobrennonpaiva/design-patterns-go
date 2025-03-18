package main

import (
	"design-patterns-go/patterns/creational"
	"fmt"
)

func main() {
	fmt.Println("Learning Design Patterns in Go")

	// Abstract Factory
	emailFactory, err := creational.GetChannelFactory("email")
	if err != nil {
		fmt.Println(err)
	}

	email := emailFactory.MakeMessage()
	fmt.Println(email.GetTitle())
	fmt.Println(email.GetBody())
	fmt.Println(email.GetRecipient())
	fmt.Println(email.GetSender())

	err = emailFactory.SendMessage()
	if err != nil {
		fmt.Println(err)
	}

	// Factory Method
	car, _ := creational.GetCar("Prius")
	car2, _ := creational.GetCar("Corolla")

	fmt.Println("Car 1: ", car.GetName())
	fmt.Println("Car 1: ", car.GetColor())
	fmt.Println("Car 1: ", car.GetEngine())

	fmt.Println("Car 2: ", car2.GetName())
	fmt.Println("Car 2: ", car2.GetColor())
	fmt.Println("Car 2: ", car2.GetEngine())

	// Builder
	cheeseBurgerBuilder := creational.GetBuilder("cheese")
	veggieBurgerBuilder := creational.GetBuilder("veggie")
	chickenBurgerBuilder := creational.GetBuilder("chicken")

	chef := creational.NewChef(cheeseBurgerBuilder)
	cheeseBurger := chef.BuildBurger()

	fmt.Println("Cheese Burger: ", cheeseBurger.Protein)
	fmt.Println("Cheese Burger: ", cheeseBurger.Carbohydrates)
	fmt.Println("Cheese Burger: ", cheeseBurger.Fat)

	chef.SetBuilder(veggieBurgerBuilder)
	veggieBurger := chef.BuildBurger()

	fmt.Println("Veggie Burger: ", veggieBurger.Protein)
	fmt.Println("Veggie Burger: ", veggieBurger.Carbohydrates)
	fmt.Println("Veggie Burger: ", veggieBurger.Fat)

	chef.SetBuilder(chickenBurgerBuilder)
	chickenBurger := chef.BuildBurger()

	fmt.Println("Chicken Burger: ", chickenBurger.Protein)
	fmt.Println("Chicken Burger: ", chickenBurger.Carbohydrates)
	fmt.Println("Chicken Burger: ", chickenBurger.Fat)

	// Prototype
	npc := creational.Npc{Name: "NPC 1"}
	npc2 := npc.Clone()

	fmt.Println("NPC 1: ", npc.Name)
	npc2.Print("NPC 2: ")

	// Singleton
	for i := 0; i < 10; i++ {
		go creational.GetInstance()
	}
	fmt.Scanln()
}
