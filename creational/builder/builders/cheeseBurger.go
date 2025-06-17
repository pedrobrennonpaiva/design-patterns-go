package builders

import "design-patterns-go/creational/builder/models"

type CheeseBurgerBuilder struct {
	models.Burger
}

func NewCheeseBurgerBuilder() *CheeseBurgerBuilder {
	return &CheeseBurgerBuilder{}
}

func (b *CheeseBurgerBuilder) SetProtein() {
	b.Protein = "Cheese"
}

func (b *CheeseBurgerBuilder) SetCarbohydrates() {
	b.Carbohydrates = "Bread"
}

func (b *CheeseBurgerBuilder) SetFat() {
	b.Fat = "Cheese"
}

func (b *CheeseBurgerBuilder) GetBurger() models.Burger {
	return b.Burger
}
