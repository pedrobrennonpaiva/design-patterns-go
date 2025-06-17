package builder

import (
	"design-patterns-go/creational/builder/builders"
	"design-patterns-go/creational/builder/models"
	"fmt"
)

type IBuilder interface {
	SetProtein()
	SetCarbohydrates()
	SetFat()
	GetBurger() models.Burger
}

func Run() {
	cheeseBurgerBuilder := getBuilder("cheese")
	veggieBurgerBuilder := getBuilder("veggie")
	chickenBurgerBuilder := getBuilder("chicken")

	chef := newChef(cheeseBurgerBuilder)
	cheeseBurger := chef.buildBurger()

	fmt.Println("Cheese Burger: ", cheeseBurger.Protein)
	fmt.Println("Cheese Burger: ", cheeseBurger.Carbohydrates)
	fmt.Println("Cheese Burger: ", cheeseBurger.Fat)

	chef.setBuilder(veggieBurgerBuilder)
	veggieBurger := chef.buildBurger()

	fmt.Println("Veggie Burger: ", veggieBurger.Protein)
	fmt.Println("Veggie Burger: ", veggieBurger.Carbohydrates)
	fmt.Println("Veggie Burger: ", veggieBurger.Fat)

	chef.setBuilder(chickenBurgerBuilder)
	chickenBurger := chef.buildBurger()

	fmt.Println("Chicken Burger: ", chickenBurger.Protein)
	fmt.Println("Chicken Burger: ", chickenBurger.Carbohydrates)
	fmt.Println("Chicken Burger: ", chickenBurger.Fat)
}

func getBuilder(builderType string) IBuilder {
	switch builderType {
	case "cheese":
		return builders.NewCheeseBurgerBuilder()
	case "veggie":
		return builders.NewVeggieBurgerBuilder()
	case "chicken":
		return builders.NewChickenBurgerBuilder()
	}

	return nil
}

type Chef struct {
	builder IBuilder
}

func newChef(b IBuilder) *Chef {
	return &Chef{
		builder: b,
	}
}

func (c *Chef) setBuilder(b IBuilder) {
	c.builder = b
}

func (c *Chef) buildBurger() models.Burger {
	c.builder.SetProtein()
	c.builder.SetCarbohydrates()
	c.builder.SetFat()
	return c.builder.GetBurger()
}
