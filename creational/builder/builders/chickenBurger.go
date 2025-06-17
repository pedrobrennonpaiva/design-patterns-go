package builders

import (
	"design-patterns-go/creational/builder/models"
)

type ChickenBurgerBuilder struct {
	models.Burger
}

func NewChickenBurgerBuilder() *ChickenBurgerBuilder {
	return &ChickenBurgerBuilder{}
}

func (b *ChickenBurgerBuilder) SetProtein() {
	b.Protein = "Chicken"
}

func (b *ChickenBurgerBuilder) SetCarbohydrates() {
	b.Carbohydrates = "Bread"
}

func (b *ChickenBurgerBuilder) SetFat() {
	b.Fat = "Chicken"
}

func (b *ChickenBurgerBuilder) GetBurger() models.Burger {
	return b.Burger
}
